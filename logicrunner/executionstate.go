//
// Copyright 2019 Insolar Technologies GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package logicrunner

import (
	"context"
	"sync"

	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/insolar/message"
	"github.com/insolar/insolar/instrumentation/inslogger"
	"github.com/insolar/insolar/logicrunner/artifacts"
)

type ExecutionState struct {
	sync.Mutex

	Ref Ref // Object reference

	ObjectDescriptor artifacts.ObjectDescriptor

	Broker                *ExecutionBroker
	CurrentList           *CurrentExecutionList
	LedgerHasMoreRequests bool
	getLedgerPendingMutex sync.Mutex

	// TODO not using in validation, need separate ObjectState.ExecutionState and ObjectState.Validation from ExecutionState struct
	pending              message.PendingState
	PendingConfirmed     bool
	HasPendingCheckMutex sync.Mutex
}

func NewExecutionState(ref insolar.Reference) *ExecutionState {
	es := &ExecutionState{
		Ref:     ref,
		pending: message.PendingUnknown,
		Broker:  NewExecutionBroker(nil),
	}
	return es
}

type ExecutionStateMethods struct {
	ledgerChecked sync.Once
	lr            *LogicRunner
	es            *ExecutionState
}

func NewExecutionStateMethods(lr *LogicRunner, es *ExecutionState) *ExecutionStateMethods {
	return &ExecutionStateMethods{
		ledgerChecked: sync.Once{},
		lr:            lr,
		es:            es,
	}
}

func (m *ExecutionStateMethods) Check(ctx context.Context) bool {
	es := m.es
	logger := inslogger.FromContext(ctx)

	// check pending state of execution (whether we can process task or not)
	es.Lock()
	if es.pending == message.PendingUnknown {
		logger.Debug("One shouldn't call ExecuteTranscript in case when pending state is unknown")
		es.Unlock()
		return false
	} else if es.pending == message.InPending {
		logger.Debug("Object in pending, wont start queue processor")
		es.Unlock()
		return false
	}
	es.Unlock()

	return true
}
func (m *ExecutionStateMethods) Execute(ctx context.Context, t *Transcript) bool {
	es := m.es

	logger := inslogger.FromContext(ctx)

	pub := m.lr.publisher

	if readyToExecute := m.Check(ctx); !readyToExecute {
		// we can get only "ErrRetryLater" here, so we'll pass it up and our
		// caller will find some way to process it
		return false
	}

	// Ask ledger kindly to give us next pending task and continue execution
	// note: should be done only once
	m.ledgerChecked.Do(func() {
		wmMessage := makeWMMessage(ctx, es.Ref.Bytes(), getLedgerPendingRequestMsg)
		if err := pub.Publish(InnerMsgTopic, wmMessage); err != nil {
			logger.Warnf("can't send processExecutionQueueMsg: ", err)
		}
	})

	re, err := m.lr.RequestsExecutor.ExecuteAndSave(ctx, t)
	if err != nil {
		inslogger.FromContext(ctx).Warn("contract execution error: ", err)
	}

	go m.lr.RequestsExecutor.SendReply(t.Context, t, re, err)

	if t.FromLedger {
		// we've already told ledger that we've processed it's task;
		// trying to take another one
		wmMessage := makeWMMessage(ctx, es.Ref.Bytes(), getLedgerPendingRequestMsg)
		if err := pub.Publish(InnerMsgTopic, wmMessage); err != nil {
			logger.Warnf("can't send processExecutionQueueMsg: ", err)
		}
	}

	// we're checking here that pulse was changed and we should send
	// a message that we've finished processing task
	// note: ideally we should tell here that we've stopped executing
	//       but we only hoped that OnPulse had already told us that
	//       pulse changed and we should stop execution
	m.lr.finishPendingIfNeeded(ctx, es)

	return true
}

func (es *ExecutionState) RegisterLogicRunner(lr *LogicRunner) {
	es.Broker.methods = NewExecutionStateMethods(lr, es)
}

// PendingNotConfirmed checks that we were in pending and waiting
// for previous executor to notify us that he still executes it
//
// Used in OnPulse to tell next executor, that it's time to continue
// work on this object
func (es *ExecutionState) InPendingNotConfirmed(ctx context.Context) bool {
	return es.pending == message.InPending && !es.PendingConfirmed
}

func (es *ExecutionState) OnPulse(ctx context.Context, meNext bool) []insolar.Message {
	if es == nil {
		return nil
	}

	logger := inslogger.FromContext(ctx)

	messages := make([]insolar.Message, 0)
	ref := es.Ref

	// if we are executor again we just continue working
	// without sending data on next executor (because we are next executor)
	if !meNext {
		sendExecResults := false

		if !es.Broker.currentList.Empty() { // nolint: gocritics
			es.pending = message.InPending
			sendExecResults = true

			// TODO: this should return delegation token to continue execution of the pending
			messages = append(
				messages,
				&message.StillExecuting{
					Reference: ref,
				},
			)
		} else if es.InPendingNotConfirmed(ctx) {
			logger.Warn("looks like pending executor died, continuing execution on next executor")
			es.pending = message.NotPending
			sendExecResults = true
			es.LedgerHasMoreRequests = true
		} else if es.Broker.finished.Len() > 0 {
			sendExecResults = true
		}

		// rotation results also contain finished requests
		rotationResults := es.Broker.Rotate(maxQueueLength)
		if len(rotationResults.Requests) > 0 || sendExecResults {
			// TODO: we also should send when executed something for validation
			// TODO: now validation is disabled
			messagesQueue := convertQueueToMessageQueue(ctx, rotationResults.Requests)

			messages = append(
				messages,
				&message.ExecutorResults{
					RecordRef:             ref,
					Pending:               es.pending,
					Queue:                 messagesQueue,
					LedgerHasMoreRequests: es.LedgerHasMoreRequests || rotationResults.LedgerHasMoreRequests,
				},
			)
		}
	} else {
		if !es.Broker.currentList.Empty() {
			// no pending should be as we are executing
			if es.pending == message.InPending {
				logger.Warn("we are executing ATM, but ES marked as pending, shouldn't be")
				es.pending = message.NotPending
			}
		} else if es.InPendingNotConfirmed(ctx) {
			logger.Warn("looks like pending executor died, re-starting execution")
			es.pending = message.NotPending
			es.LedgerHasMoreRequests = true
		}
		es.PendingConfirmed = false
	}

	return messages
}
