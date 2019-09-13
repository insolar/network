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

package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/pkg/errors"

	"github.com/insolar/insolar/api/sdk"
	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/insolar/backoff"
)

type createMembersScenario struct {
	name           string
	concurrent     int
	repetitions    int
	out            io.Writer
	totalTime      int64
	goroutineTimes []time.Duration
	successes      uint32
	errors         uint32
	timeouts       uint32
	insSDK         *sdk.SDK
	penRetries     int32
}

func (s *createMembersScenario) getOperationsNumber() int {
	return s.concurrent * s.repetitions
}

func (s *createMembersScenario) getAverageOperationDuration() time.Duration {
	return time.Duration(s.totalTime / int64(s.getOperationsNumber()))
}

func (s *createMembersScenario) getOperationPerSecond() float64 {
	if len(s.goroutineTimes) == 0 {
		return 0
	}

	max := s.goroutineTimes[0]
	for _, t := range s.goroutineTimes {
		if max < t {
			max = t
		}
	}
	elapsedInSeconds := float64(max) / float64(time.Second)
	return float64(s.getOperationsNumber()-int(s.timeouts)) / elapsedInSeconds
}

func (s *createMembersScenario) getName() string {
	return s.name
}

func (s *createMembersScenario) getOut() io.Writer {
	return s.out
}

func (s *createMembersScenario) prepare() {

}

func (s *createMembersScenario) canBeStarted() error {
	writeToOutput(s.getOut(), fmt.Sprint("canBeStarted\n"))
	return nil
}

func (s *createMembersScenario) start(ctx context.Context) {
	var wg sync.WaitGroup
	for i := 0; i < s.concurrent*2; i += 2 {
		wg.Add(1)
		go s.startMember(ctx, i, &wg)
	}
	wg.Wait()
}

func (s *createMembersScenario) checkResult() {

}

func (s *createMembersScenario) startMember(ctx context.Context, index int, wg *sync.WaitGroup) {
	defer wg.Done()
	goroutineTime := time.Duration(0)
	for j := 0; j < s.repetitions; j++ {
		select {
		case <-ctx.Done():
			return
		default:
		}

		var start time.Time
		var stop time.Duration
		var traceID string
		var err error

		bof := backoff.Backoff{Min: 500 * time.Millisecond, Max: 20 * time.Second}

		retry := true
		for retry && bof.Attempt() < backoffAttemptsCount {
			start = time.Now()
			_, traceID, err = s.insSDK.CreateMember()
			stop = time.Since(start)

			if err != nil && strings.Contains(err.Error(), insolar.ErrTooManyPendingRequests.Error()) {
				time.Sleep(bof.Duration())
				atomic.AddInt32(&s.penRetries, 1)
			} else {
				retry = false
			}
		}

		if err == nil {
			atomic.AddUint32(&s.successes, 1)
			atomic.AddInt64(&s.totalTime, int64(stop))
			goroutineTime += stop
		} else if netErr, ok := errors.Cause(err).(net.Error); ok && netErr.Timeout() {
			atomic.AddUint32(&s.timeouts, 1)
			writeToOutput(s.out, fmt.Sprintf("[Member №%d] Create member error. Timeout. Error: %s \n", index, err.Error()))
		} else {
			atomic.AddUint32(&s.errors, 1)
			atomic.AddInt64(&s.totalTime, int64(stop))
			goroutineTime += stop

			if strings.Contains(err.Error(), "invalid state record") {
				writeToOutput(s.out, fmt.Sprintf("[ OK ] Invalid state record.    Trace: %s.\n", traceID))
			} else {
				writeToOutput(s.out, fmt.Sprintf("[Member №%d] Create member error with traceID: %s. Response: %s.\n", index, traceID, err.Error()))
			}
		}
	}
	s.goroutineTimes = append(s.goroutineTimes, goroutineTime)
}

func (s *createMembersScenario) printResult() {
	writeToOutput(s.out, fmt.Sprintf("Scenario result:\n\tSuccesses: %d\n\tErrors: %d\n\tTimeouts: %d\n\tPending retries: %d\n", s.successes, s.errors, s.timeouts, s.penRetries))
}