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

// Code generated by insgocc. DO NOT EDIT.
// source template in logicrunner/preprocessor/templates

package wallet

import (
	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/logicrunner/builtin/foundation"
	"github.com/insolar/insolar/logicrunner/common"
	"github.com/pkg/errors"
)

func INS_META_INFO() []map[string]string {
	result := make([]map[string]string, 0)

	return result
}

func INSMETHOD_GetCode(object []byte, data []byte) ([]byte, []byte, error) {
	ph := common.CurrentProxyCtx
	self := new(Wallet)

	if len(object) == 0 {
		return nil, nil, &foundation.Error{S: "[ Fake GetCode ] ( Generated Method ) Object is nil"}
	}

	err := ph.Deserialize(object, self)
	if err != nil {
		e := &foundation.Error{S: "[ Fake GetCode ] ( Generated Method ) Can't deserialize args.Data: " + err.Error()}
		return nil, nil, e
	}

	state := []byte{}
	err = ph.Serialize(self, &state)
	if err != nil {
		return nil, nil, err
	}

	ret := []byte{}
	err = ph.Serialize([]interface{}{self.GetCode().Bytes()}, &ret)

	return state, ret, err
}

func INSMETHOD_GetPrototype(object []byte, data []byte) ([]byte, []byte, error) {
	ph := common.CurrentProxyCtx
	self := new(Wallet)

	if len(object) == 0 {
		return nil, nil, &foundation.Error{S: "[ Fake GetPrototype ] ( Generated Method ) Object is nil"}
	}

	err := ph.Deserialize(object, self)
	if err != nil {
		e := &foundation.Error{S: "[ Fake GetPrototype ] ( Generated Method ) Can't deserialize args.Data: " + err.Error()}
		return nil, nil, e
	}

	state := []byte{}
	err = ph.Serialize(self, &state)
	if err != nil {
		return nil, nil, err
	}

	ret := []byte{}
	err = ph.Serialize([]interface{}{self.GetPrototype().Bytes()}, &ret)

	return state, ret, err
}

func INSMETHOD_GetAccount(object []byte, data []byte) (newState []byte, result []byte, err error) {
	ph := common.CurrentProxyCtx
	ph.SetSystemError(nil)

	self := new(Wallet)

	if len(object) == 0 {
		err = &foundation.Error{S: "[ FakeGetAccount ] ( INSMETHOD_* ) ( Generated Method ) Object is nil"}
		return
	}

	err = ph.Deserialize(object, self)
	if err != nil {
		err = &foundation.Error{S: "[ FakeGetAccount ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Data: " + err.Error()}
		return
	}

	args := make([]interface{}, 1)
	var args0 string
	args[0] = &args0

	err = ph.Deserialize(data, &args)
	if err != nil {
		err = &foundation.Error{S: "[ FakeGetAccount ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Arguments: " + err.Error()}
		return
	}

	var ret0 *insolar.Reference
	var ret1 error

	serializeResults := func() error {
		return ph.Serialize(
			foundation.Result{Returns: []interface{}{ret0, ret1}},
			&result,
		)
	}

	needRecover := true
	defer func() {
		if !needRecover {
			return
		}
		if r := recover(); r != nil {
			recoveredError := errors.Wrap(errors.Errorf("%v", r), "Failed to execute method (panic)")
			ret1 = ph.MakeErrorSerializable(recoveredError)

			newState = object
			err = serializeResults()
		}
	}()

	ret0, ret1 = self.GetAccount(args0)

	needRecover = false

	if ph.GetSystemError() != nil {
		return nil, nil, ph.GetSystemError()
	}

	err = ph.Serialize(self, &newState)
	if err != nil {
		return nil, nil, err
	}

	ret1 = ph.MakeErrorSerializable(ret1)

	err = serializeResults()
	if err != nil {
		return
	}

	return
}

func INSMETHOD_Transfer(object []byte, data []byte) (newState []byte, result []byte, err error) {
	ph := common.CurrentProxyCtx
	ph.SetSystemError(nil)

	self := new(Wallet)

	if len(object) == 0 {
		err = &foundation.Error{S: "[ FakeTransfer ] ( INSMETHOD_* ) ( Generated Method ) Object is nil"}
		return
	}

	err = ph.Deserialize(object, self)
	if err != nil {
		err = &foundation.Error{S: "[ FakeTransfer ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Data: " + err.Error()}
		return
	}

	args := make([]interface{}, 5)
	var args0 string
	args[0] = &args0
	var args1 string
	args[1] = &args1
	var args2 *insolar.Reference
	args[2] = &args2
	var args3 insolar.Reference
	args[3] = &args3
	var args4 insolar.Reference
	args[4] = &args4

	err = ph.Deserialize(data, &args)
	if err != nil {
		err = &foundation.Error{S: "[ FakeTransfer ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Arguments: " + err.Error()}
		return
	}

	var ret0 interface{}
	var ret1 error

	serializeResults := func() error {
		return ph.Serialize(
			foundation.Result{Returns: []interface{}{ret0, ret1}},
			&result,
		)
	}

	needRecover := true
	defer func() {
		if !needRecover {
			return
		}
		if r := recover(); r != nil {
			recoveredError := errors.Wrap(errors.Errorf("%v", r), "Failed to execute method (panic)")
			ret1 = ph.MakeErrorSerializable(recoveredError)

			newState = object
			err = serializeResults()
		}
	}()

	ret0, ret1 = self.Transfer(args0, args1, args2, args3, args4)

	needRecover = false

	if ph.GetSystemError() != nil {
		return nil, nil, ph.GetSystemError()
	}

	err = ph.Serialize(self, &newState)
	if err != nil {
		return nil, nil, err
	}

	ret1 = ph.MakeErrorSerializable(ret1)

	err = serializeResults()
	if err != nil {
		return
	}

	return
}

func INSMETHOD_GetBalance(object []byte, data []byte) (newState []byte, result []byte, err error) {
	ph := common.CurrentProxyCtx
	ph.SetSystemError(nil)

	self := new(Wallet)

	if len(object) == 0 {
		err = &foundation.Error{S: "[ FakeGetBalance ] ( INSMETHOD_* ) ( Generated Method ) Object is nil"}
		return
	}

	err = ph.Deserialize(object, self)
	if err != nil {
		err = &foundation.Error{S: "[ FakeGetBalance ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Data: " + err.Error()}
		return
	}

	args := make([]interface{}, 1)
	var args0 string
	args[0] = &args0

	err = ph.Deserialize(data, &args)
	if err != nil {
		err = &foundation.Error{S: "[ FakeGetBalance ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Arguments: " + err.Error()}
		return
	}

	var ret0 string
	var ret1 error

	serializeResults := func() error {
		return ph.Serialize(
			foundation.Result{Returns: []interface{}{ret0, ret1}},
			&result,
		)
	}

	needRecover := true
	defer func() {
		if !needRecover {
			return
		}
		if r := recover(); r != nil {
			recoveredError := errors.Wrap(errors.Errorf("%v", r), "Failed to execute method (panic)")
			ret1 = ph.MakeErrorSerializable(recoveredError)

			newState = object
			err = serializeResults()
		}
	}()

	ret0, ret1 = self.GetBalance(args0)

	needRecover = false

	if ph.GetSystemError() != nil {
		return nil, nil, ph.GetSystemError()
	}

	err = ph.Serialize(self, &newState)
	if err != nil {
		return nil, nil, err
	}

	ret1 = ph.MakeErrorSerializable(ret1)

	err = serializeResults()
	if err != nil {
		return
	}

	return
}

func INSMETHOD_AddDeposit(object []byte, data []byte) (newState []byte, result []byte, err error) {
	ph := common.CurrentProxyCtx
	ph.SetSystemError(nil)

	self := new(Wallet)

	if len(object) == 0 {
		err = &foundation.Error{S: "[ FakeAddDeposit ] ( INSMETHOD_* ) ( Generated Method ) Object is nil"}
		return
	}

	err = ph.Deserialize(object, self)
	if err != nil {
		err = &foundation.Error{S: "[ FakeAddDeposit ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Data: " + err.Error()}
		return
	}

	args := make([]interface{}, 2)
	var args0 string
	args[0] = &args0
	var args1 insolar.Reference
	args[1] = &args1

	err = ph.Deserialize(data, &args)
	if err != nil {
		err = &foundation.Error{S: "[ FakeAddDeposit ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Arguments: " + err.Error()}
		return
	}

	var ret0 error

	serializeResults := func() error {
		return ph.Serialize(
			foundation.Result{Returns: []interface{}{ret0}},
			&result,
		)
	}

	needRecover := true
	defer func() {
		if !needRecover {
			return
		}
		if r := recover(); r != nil {
			recoveredError := errors.Wrap(errors.Errorf("%v", r), "Failed to execute method (panic)")
			ret0 = ph.MakeErrorSerializable(recoveredError)

			newState = object
			err = serializeResults()
		}
	}()

	ret0 = self.AddDeposit(args0, args1)

	needRecover = false

	if ph.GetSystemError() != nil {
		return nil, nil, ph.GetSystemError()
	}

	err = ph.Serialize(self, &newState)
	if err != nil {
		return nil, nil, err
	}

	ret0 = ph.MakeErrorSerializable(ret0)

	err = serializeResults()
	if err != nil {
		return
	}

	return
}

func INSMETHOD_GetDeposits(object []byte, data []byte) (newState []byte, result []byte, err error) {
	ph := common.CurrentProxyCtx
	ph.SetSystemError(nil)

	self := new(Wallet)

	if len(object) == 0 {
		err = &foundation.Error{S: "[ FakeGetDeposits ] ( INSMETHOD_* ) ( Generated Method ) Object is nil"}
		return
	}

	err = ph.Deserialize(object, self)
	if err != nil {
		err = &foundation.Error{S: "[ FakeGetDeposits ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Data: " + err.Error()}
		return
	}

	args := []interface{}{}

	err = ph.Deserialize(data, &args)
	if err != nil {
		err = &foundation.Error{S: "[ FakeGetDeposits ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Arguments: " + err.Error()}
		return
	}

	var ret0 []interface{}
	var ret1 error

	serializeResults := func() error {
		return ph.Serialize(
			foundation.Result{Returns: []interface{}{ret0, ret1}},
			&result,
		)
	}

	needRecover := true
	defer func() {
		if !needRecover {
			return
		}
		if r := recover(); r != nil {
			recoveredError := errors.Wrap(errors.Errorf("%v", r), "Failed to execute method (panic)")
			ret1 = ph.MakeErrorSerializable(recoveredError)

			newState = object
			err = serializeResults()
		}
	}()

	ret0, ret1 = self.GetDeposits()

	needRecover = false

	if ph.GetSystemError() != nil {
		return nil, nil, ph.GetSystemError()
	}

	err = ph.Serialize(self, &newState)
	if err != nil {
		return nil, nil, err
	}

	ret1 = ph.MakeErrorSerializable(ret1)

	err = serializeResults()
	if err != nil {
		return
	}

	return
}

func INSMETHOD_FindDeposit(object []byte, data []byte) (newState []byte, result []byte, err error) {
	ph := common.CurrentProxyCtx
	ph.SetSystemError(nil)

	self := new(Wallet)

	if len(object) == 0 {
		err = &foundation.Error{S: "[ FakeFindDeposit ] ( INSMETHOD_* ) ( Generated Method ) Object is nil"}
		return
	}

	err = ph.Deserialize(object, self)
	if err != nil {
		err = &foundation.Error{S: "[ FakeFindDeposit ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Data: " + err.Error()}
		return
	}

	args := make([]interface{}, 1)
	var args0 string
	args[0] = &args0

	err = ph.Deserialize(data, &args)
	if err != nil {
		err = &foundation.Error{S: "[ FakeFindDeposit ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Arguments: " + err.Error()}
		return
	}

	var ret0 bool
	var ret1 *insolar.Reference
	var ret2 error

	serializeResults := func() error {
		return ph.Serialize(
			foundation.Result{Returns: []interface{}{ret0, ret1, ret2}},
			&result,
		)
	}

	needRecover := true
	defer func() {
		if !needRecover {
			return
		}
		if r := recover(); r != nil {
			recoveredError := errors.Wrap(errors.Errorf("%v", r), "Failed to execute method (panic)")
			ret2 = ph.MakeErrorSerializable(recoveredError)

			newState = object
			err = serializeResults()
		}
	}()

	ret0, ret1, ret2 = self.FindDeposit(args0)

	needRecover = false

	if ph.GetSystemError() != nil {
		return nil, nil, ph.GetSystemError()
	}

	err = ph.Serialize(self, &newState)
	if err != nil {
		return nil, nil, err
	}

	ret2 = ph.MakeErrorSerializable(ret2)

	err = serializeResults()
	if err != nil {
		return
	}

	return
}

func INSCONSTRUCTOR_New(ref insolar.Reference, data []byte) (state []byte, result []byte, err error) {
	ph := common.CurrentProxyCtx
	ph.SetSystemError(nil)

	args := make([]interface{}, 1)
	var args0 insolar.Reference
	args[0] = &args0

	err = ph.Deserialize(data, &args)
	if err != nil {
		err = &foundation.Error{S: "[ FakeNew ] ( INSCONSTRUCTOR_* ) ( Generated Method ) Can't deserialize args.Arguments: " + err.Error()}
		return
	}

	var ret0 *Wallet
	var ret1 error

	serializeResults := func() error {
		return ph.Serialize(
			foundation.Result{Returns: []interface{}{ref, ret1}},
			&result,
		)
	}

	needRecover := true
	defer func() {
		if !needRecover {
			return
		}
		if r := recover(); r != nil {
			recoveredError := errors.Wrap(errors.Errorf("%v", r), "Failed to execute method (panic)")
			ret1 = ph.MakeErrorSerializable(recoveredError)

			state = data
			err = serializeResults()
		}
	}()

	ret0, ret1 = New(args0)

	needRecover = false

	ret1 = ph.MakeErrorSerializable(ret1)
	if ret0 == nil && ret1 == nil {
		ret1 = &foundation.Error{S: "constructor returned nil"}
	}

	if ph.GetSystemError() != nil {
		err = ph.GetSystemError()
		return
	}

	err = serializeResults()
	if err != nil {
		return
	}

	if ret1 != nil {
		// logical error, the result should be registered with type RequestSideEffectNone
		state = nil
		return
	}

	err = ph.Serialize(ret0, &state)
	if err != nil {
		return
	}

	return
}

func Initialize() insolar.ContractWrapper {
	return insolar.ContractWrapper{
		GetCode:      INSMETHOD_GetCode,
		GetPrototype: INSMETHOD_GetPrototype,
		Methods: insolar.ContractMethods{
			"GetAccount":  INSMETHOD_GetAccount,
			"Transfer":    INSMETHOD_Transfer,
			"GetBalance":  INSMETHOD_GetBalance,
			"AddDeposit":  INSMETHOD_AddDeposit,
			"GetDeposits": INSMETHOD_GetDeposits,
			"FindDeposit": INSMETHOD_FindDeposit,
		},
		Constructors: insolar.ContractConstructors{
			"New": INSCONSTRUCTOR_New,
		},
	}
}
