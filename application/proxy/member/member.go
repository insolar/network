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

package member

import (
	"github.com/insolar/insolar/application/proxy/deposit"
	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/logicrunner/goplugin/foundation"
	"github.com/insolar/insolar/logicrunner/goplugin/proxyctx"
)

// PrototypeReference to prototype of this contract
// error checking hides in generator
var PrototypeReference, _ = insolar.NewReferenceFromBase58("11112hAiqZPxGpUKi6u1EAd8vYf5JYBoMcRcMAGuTND.11111111111111111111111111111111")

// Member holds proxy type
type Member struct {
	Reference insolar.Reference
	Prototype insolar.Reference
	Code      insolar.Reference
}

// ContractConstructorHolder holds logic with object construction
type ContractConstructorHolder struct {
	constructorName string
	argsSerialized  []byte
}

// AsChild saves object as child
func (r *ContractConstructorHolder) AsChild(objRef insolar.Reference) (*Member, error) {
	ref, err := proxyctx.Current.SaveAsChild(objRef, *PrototypeReference, r.constructorName, r.argsSerialized)
	if err != nil {
		return nil, err
	}
	return &Member{Reference: ref}, nil
}

// AsDelegate saves object as delegate
func (r *ContractConstructorHolder) AsDelegate(objRef insolar.Reference) (*Member, error) {
	ref, err := proxyctx.Current.SaveAsDelegate(objRef, *PrototypeReference, r.constructorName, r.argsSerialized)
	if err != nil {
		return nil, err
	}
	return &Member{Reference: ref}, nil
}

// GetObject returns proxy object
func GetObject(ref insolar.Reference) (r *Member) {
	return &Member{Reference: ref}
}

// GetPrototype returns reference to the prototype
func GetPrototype() insolar.Reference {
	return *PrototypeReference
}

// GetImplementationFrom returns proxy to delegate of given type
func GetImplementationFrom(object insolar.Reference) (*Member, error) {
	ref, err := proxyctx.Current.GetDelegate(object, *PrototypeReference)
	if err != nil {
		return nil, err
	}
	return GetObject(ref), nil
}

// New is constructor
func New(ethAddr string, key string) *ContractConstructorHolder {
	var args [2]interface{}
	args[0] = ethAddr
	args[1] = key

	var argsSerialized []byte
	err := proxyctx.Current.Serialize(args, &argsSerialized)
	if err != nil {
		panic(err)
	}

	return &ContractConstructorHolder{constructorName: "New", argsSerialized: argsSerialized}
}

// NewOracleMember is constructor
func NewOracleMember(name string, key string) *ContractConstructorHolder {
	var args [2]interface{}
	args[0] = name
	args[1] = key

	var argsSerialized []byte
	err := proxyctx.Current.Serialize(args, &argsSerialized)
	if err != nil {
		panic(err)
	}

	return &ContractConstructorHolder{constructorName: "NewOracleMember", argsSerialized: argsSerialized}
}

// GetReference returns reference of the object
func (r *Member) GetReference() insolar.Reference {
	return r.Reference
}

// GetPrototype returns reference to the code
func (r *Member) GetPrototype() (insolar.Reference, error) {
	if r.Prototype.IsEmpty() {
		ret := [2]interface{}{}
		var ret0 insolar.Reference
		ret[0] = &ret0
		var ret1 *foundation.Error
		ret[1] = &ret1

		res, err := proxyctx.Current.RouteCall(r.Reference, true, false, "GetPrototype", make([]byte, 0), *PrototypeReference)
		if err != nil {
			return ret0, err
		}

		err = proxyctx.Current.Deserialize(res, &ret)
		if err != nil {
			return ret0, err
		}

		if ret1 != nil {
			return ret0, ret1
		}

		r.Prototype = ret0
	}

	return r.Prototype, nil

}

// GetCode returns reference to the code
func (r *Member) GetCode() (insolar.Reference, error) {
	if r.Code.IsEmpty() {
		ret := [2]interface{}{}
		var ret0 insolar.Reference
		ret[0] = &ret0
		var ret1 *foundation.Error
		ret[1] = &ret1

		res, err := proxyctx.Current.RouteCall(r.Reference, true, false, "GetCode", make([]byte, 0), *PrototypeReference)
		if err != nil {
			return ret0, err
		}

		err = proxyctx.Current.Deserialize(res, &ret)
		if err != nil {
			return ret0, err
		}

		if ret1 != nil {
			return ret0, ret1
		}

		r.Code = ret0
	}

	return r.Code, nil
}

// GetName is proxy generated method
func (r *Member) GetName() (string, error) {
	var args [0]interface{}

	var argsSerialized []byte

	ret := [2]interface{}{}
	var ret0 string
	ret[0] = &ret0
	var ret1 *foundation.Error
	ret[1] = &ret1

	err := proxyctx.Current.Serialize(args, &argsSerialized)
	if err != nil {
		return ret0, err
	}

	res, err := proxyctx.Current.RouteCall(r.Reference, true, false, "GetName", argsSerialized, *PrototypeReference)
	if err != nil {
		return ret0, err
	}

	err = proxyctx.Current.Deserialize(res, &ret)
	if err != nil {
		return ret0, err
	}

	if ret1 != nil {
		return ret0, ret1
	}
	return ret0, nil
}

// GetNameNoWait is proxy generated method
func (r *Member) GetNameNoWait() error {
	var args [0]interface{}

	var argsSerialized []byte

	err := proxyctx.Current.Serialize(args, &argsSerialized)
	if err != nil {
		return err
	}

	_, err = proxyctx.Current.RouteCall(r.Reference, false, false, "GetName", argsSerialized, *PrototypeReference)
	if err != nil {
		return err
	}

	return nil
}

// GetNameAsImmutable is proxy generated method
func (r *Member) GetNameAsImmutable() (string, error) {
	var args [0]interface{}

	var argsSerialized []byte

	ret := [2]interface{}{}
	var ret0 string
	ret[0] = &ret0
	var ret1 *foundation.Error
	ret[1] = &ret1

	err := proxyctx.Current.Serialize(args, &argsSerialized)
	if err != nil {
		return ret0, err
	}

	res, err := proxyctx.Current.RouteCall(r.Reference, true, true, "GetName", argsSerialized, *PrototypeReference)
	if err != nil {
		return ret0, err
	}

	err = proxyctx.Current.Deserialize(res, &ret)
	if err != nil {
		return ret0, err
	}

	if ret1 != nil {
		return ret0, ret1
	}
	return ret0, nil
}

// GetEthAddr is proxy generated method
func (r *Member) GetEthAddr() (string, error) {
	var args [0]interface{}

	var argsSerialized []byte

	ret := [2]interface{}{}
	var ret0 string
	ret[0] = &ret0
	var ret1 *foundation.Error
	ret[1] = &ret1

	err := proxyctx.Current.Serialize(args, &argsSerialized)
	if err != nil {
		return ret0, err
	}

	res, err := proxyctx.Current.RouteCall(r.Reference, true, false, "GetEthAddr", argsSerialized, *PrototypeReference)
	if err != nil {
		return ret0, err
	}

	err = proxyctx.Current.Deserialize(res, &ret)
	if err != nil {
		return ret0, err
	}

	if ret1 != nil {
		return ret0, ret1
	}
	return ret0, nil
}

// GetEthAddrNoWait is proxy generated method
func (r *Member) GetEthAddrNoWait() error {
	var args [0]interface{}

	var argsSerialized []byte

	err := proxyctx.Current.Serialize(args, &argsSerialized)
	if err != nil {
		return err
	}

	_, err = proxyctx.Current.RouteCall(r.Reference, false, false, "GetEthAddr", argsSerialized, *PrototypeReference)
	if err != nil {
		return err
	}

	return nil
}

// GetEthAddrAsImmutable is proxy generated method
func (r *Member) GetEthAddrAsImmutable() (string, error) {
	var args [0]interface{}

	var argsSerialized []byte

	ret := [2]interface{}{}
	var ret0 string
	ret[0] = &ret0
	var ret1 *foundation.Error
	ret[1] = &ret1

	err := proxyctx.Current.Serialize(args, &argsSerialized)
	if err != nil {
		return ret0, err
	}

	res, err := proxyctx.Current.RouteCall(r.Reference, true, true, "GetEthAddr", argsSerialized, *PrototypeReference)
	if err != nil {
		return ret0, err
	}

	err = proxyctx.Current.Deserialize(res, &ret)
	if err != nil {
		return ret0, err
	}

	if ret1 != nil {
		return ret0, ret1
	}
	return ret0, nil
}

// SetEthAddr is proxy generated method
func (r *Member) SetEthAddr(ethAddr string) error {
	var args [1]interface{}
	args[0] = ethAddr

	var argsSerialized []byte

	ret := [1]interface{}{}
	var ret0 *foundation.Error
	ret[0] = &ret0

	err := proxyctx.Current.Serialize(args, &argsSerialized)
	if err != nil {
		return err
	}

	res, err := proxyctx.Current.RouteCall(r.Reference, true, false, "SetEthAddr", argsSerialized, *PrototypeReference)
	if err != nil {
		return err
	}

	err = proxyctx.Current.Deserialize(res, &ret)
	if err != nil {
		return err
	}

	if ret0 != nil {
		return ret0
	}
	return nil
}

// SetEthAddrNoWait is proxy generated method
func (r *Member) SetEthAddrNoWait(ethAddr string) error {
	var args [1]interface{}
	args[0] = ethAddr

	var argsSerialized []byte

	err := proxyctx.Current.Serialize(args, &argsSerialized)
	if err != nil {
		return err
	}

	_, err = proxyctx.Current.RouteCall(r.Reference, false, false, "SetEthAddr", argsSerialized, *PrototypeReference)
	if err != nil {
		return err
	}

	return nil
}

// SetEthAddrAsImmutable is proxy generated method
func (r *Member) SetEthAddrAsImmutable(ethAddr string) error {
	var args [1]interface{}
	args[0] = ethAddr

	var argsSerialized []byte

	ret := [1]interface{}{}
	var ret0 *foundation.Error
	ret[0] = &ret0

	err := proxyctx.Current.Serialize(args, &argsSerialized)
	if err != nil {
		return err
	}

	res, err := proxyctx.Current.RouteCall(r.Reference, true, true, "SetEthAddr", argsSerialized, *PrototypeReference)
	if err != nil {
		return err
	}

	err = proxyctx.Current.Deserialize(res, &ret)
	if err != nil {
		return err
	}

	if ret0 != nil {
		return ret0
	}
	return nil
}

// GetPublicKey is proxy generated method
func (r *Member) GetPublicKey() (string, error) {
	var args [0]interface{}

	var argsSerialized []byte

	ret := [2]interface{}{}
	var ret0 string
	ret[0] = &ret0
	var ret1 *foundation.Error
	ret[1] = &ret1

	err := proxyctx.Current.Serialize(args, &argsSerialized)
	if err != nil {
		return ret0, err
	}

	res, err := proxyctx.Current.RouteCall(r.Reference, true, false, "GetPublicKey", argsSerialized, *PrototypeReference)
	if err != nil {
		return ret0, err
	}

	err = proxyctx.Current.Deserialize(res, &ret)
	if err != nil {
		return ret0, err
	}

	if ret1 != nil {
		return ret0, ret1
	}
	return ret0, nil
}

// GetPublicKeyNoWait is proxy generated method
func (r *Member) GetPublicKeyNoWait() error {
	var args [0]interface{}

	var argsSerialized []byte

	err := proxyctx.Current.Serialize(args, &argsSerialized)
	if err != nil {
		return err
	}

	_, err = proxyctx.Current.RouteCall(r.Reference, false, false, "GetPublicKey", argsSerialized, *PrototypeReference)
	if err != nil {
		return err
	}

	return nil
}

// GetPublicKeyAsImmutable is proxy generated method
func (r *Member) GetPublicKeyAsImmutable() (string, error) {
	var args [0]interface{}

	var argsSerialized []byte

	ret := [2]interface{}{}
	var ret0 string
	ret[0] = &ret0
	var ret1 *foundation.Error
	ret[1] = &ret1

	err := proxyctx.Current.Serialize(args, &argsSerialized)
	if err != nil {
		return ret0, err
	}

	res, err := proxyctx.Current.RouteCall(r.Reference, true, true, "GetPublicKey", argsSerialized, *PrototypeReference)
	if err != nil {
		return ret0, err
	}

	err = proxyctx.Current.Deserialize(res, &ret)
	if err != nil {
		return ret0, err
	}

	if ret1 != nil {
		return ret0, ret1
	}
	return ret0, nil
}

// Call is proxy generated method
func (r *Member) Call(rootDomainRef insolar.Reference, method string, params []byte, seed []byte, sign []byte) (interface{}, error) {
	var args [5]interface{}
	args[0] = rootDomainRef
	args[1] = method
	args[2] = params
	args[3] = seed
	args[4] = sign

	var argsSerialized []byte

	ret := [2]interface{}{}
	var ret0 interface{}
	ret[0] = &ret0
	var ret1 *foundation.Error
	ret[1] = &ret1

	err := proxyctx.Current.Serialize(args, &argsSerialized)
	if err != nil {
		return ret0, err
	}

	res, err := proxyctx.Current.RouteCall(r.Reference, true, false, "Call", argsSerialized, *PrototypeReference)
	if err != nil {
		return ret0, err
	}

	err = proxyctx.Current.Deserialize(res, &ret)
	if err != nil {
		return ret0, err
	}

	if ret1 != nil {
		return ret0, ret1
	}
	return ret0, nil
}

// CallNoWait is proxy generated method
func (r *Member) CallNoWait(rootDomainRef insolar.Reference, method string, params []byte, seed []byte, sign []byte) error {
	var args [5]interface{}
	args[0] = rootDomainRef
	args[1] = method
	args[2] = params
	args[3] = seed
	args[4] = sign

	var argsSerialized []byte

	err := proxyctx.Current.Serialize(args, &argsSerialized)
	if err != nil {
		return err
	}

	_, err = proxyctx.Current.RouteCall(r.Reference, false, false, "Call", argsSerialized, *PrototypeReference)
	if err != nil {
		return err
	}

	return nil
}

// CallAsImmutable is proxy generated method
func (r *Member) CallAsImmutable(rootDomainRef insolar.Reference, method string, params []byte, seed []byte, sign []byte) (interface{}, error) {
	var args [5]interface{}
	args[0] = rootDomainRef
	args[1] = method
	args[2] = params
	args[3] = seed
	args[4] = sign

	var argsSerialized []byte

	ret := [2]interface{}{}
	var ret0 interface{}
	ret[0] = &ret0
	var ret1 *foundation.Error
	ret[1] = &ret1

	err := proxyctx.Current.Serialize(args, &argsSerialized)
	if err != nil {
		return ret0, err
	}

	res, err := proxyctx.Current.RouteCall(r.Reference, true, true, "Call", argsSerialized, *PrototypeReference)
	if err != nil {
		return ret0, err
	}

	err = proxyctx.Current.Deserialize(res, &ret)
	if err != nil {
		return ret0, err
	}

	if ret1 != nil {
		return ret0, ret1
	}
	return ret0, nil
}

// FindDeposit is proxy generated method
func (r *Member) FindDeposit(txHash string, amount uint) (bool, deposit.Deposit, error) {
	var args [2]interface{}
	args[0] = txHash
	args[1] = amount

	var argsSerialized []byte

	ret := [3]interface{}{}
	var ret0 bool
	ret[0] = &ret0
	var ret1 deposit.Deposit
	ret[1] = &ret1
	var ret2 *foundation.Error
	ret[2] = &ret2

	err := proxyctx.Current.Serialize(args, &argsSerialized)
	if err != nil {
		return ret0, ret1, err
	}

	res, err := proxyctx.Current.RouteCall(r.Reference, true, false, "FindDeposit", argsSerialized, *PrototypeReference)
	if err != nil {
		return ret0, ret1, err
	}

	err = proxyctx.Current.Deserialize(res, &ret)
	if err != nil {
		return ret0, ret1, err
	}

	if ret2 != nil {
		return ret0, ret1, ret2
	}
	return ret0, ret1, nil
}

// FindDepositNoWait is proxy generated method
func (r *Member) FindDepositNoWait(txHash string, amount uint) error {
	var args [2]interface{}
	args[0] = txHash
	args[1] = amount

	var argsSerialized []byte

	err := proxyctx.Current.Serialize(args, &argsSerialized)
	if err != nil {
		return err
	}

	_, err = proxyctx.Current.RouteCall(r.Reference, false, false, "FindDeposit", argsSerialized, *PrototypeReference)
	if err != nil {
		return err
	}

	return nil
}

// FindDepositAsImmutable is proxy generated method
func (r *Member) FindDepositAsImmutable(txHash string, amount uint) (bool, deposit.Deposit, error) {
	var args [2]interface{}
	args[0] = txHash
	args[1] = amount

	var argsSerialized []byte

	ret := [3]interface{}{}
	var ret0 bool
	ret[0] = &ret0
	var ret1 deposit.Deposit
	ret[1] = &ret1
	var ret2 *foundation.Error
	ret[2] = &ret2

	err := proxyctx.Current.Serialize(args, &argsSerialized)
	if err != nil {
		return ret0, ret1, err
	}

	res, err := proxyctx.Current.RouteCall(r.Reference, true, true, "FindDeposit", argsSerialized, *PrototypeReference)
	if err != nil {
		return ret0, ret1, err
	}

	err = proxyctx.Current.Deserialize(res, &ret)
	if err != nil {
		return ret0, ret1, err
	}

	if ret2 != nil {
		return ret0, ret1, ret2
	}
	return ret0, ret1, nil
}

// DumpUserInfo is proxy generated method
func (r *Member) DumpUserInfo(rdRef insolar.Reference, userRef insolar.Reference) ([]byte, error) {
	var args [2]interface{}
	args[0] = rdRef
	args[1] = userRef

	var argsSerialized []byte

	ret := [2]interface{}{}
	var ret0 []byte
	ret[0] = &ret0
	var ret1 *foundation.Error
	ret[1] = &ret1

	err := proxyctx.Current.Serialize(args, &argsSerialized)
	if err != nil {
		return ret0, err
	}

	res, err := proxyctx.Current.RouteCall(r.Reference, true, false, "DumpUserInfo", argsSerialized, *PrototypeReference)
	if err != nil {
		return ret0, err
	}

	err = proxyctx.Current.Deserialize(res, &ret)
	if err != nil {
		return ret0, err
	}

	if ret1 != nil {
		return ret0, ret1
	}
	return ret0, nil
}

// DumpUserInfoNoWait is proxy generated method
func (r *Member) DumpUserInfoNoWait(rdRef insolar.Reference, userRef insolar.Reference) error {
	var args [2]interface{}
	args[0] = rdRef
	args[1] = userRef

	var argsSerialized []byte

	err := proxyctx.Current.Serialize(args, &argsSerialized)
	if err != nil {
		return err
	}

	_, err = proxyctx.Current.RouteCall(r.Reference, false, false, "DumpUserInfo", argsSerialized, *PrototypeReference)
	if err != nil {
		return err
	}

	return nil
}

// DumpUserInfoAsImmutable is proxy generated method
func (r *Member) DumpUserInfoAsImmutable(rdRef insolar.Reference, userRef insolar.Reference) ([]byte, error) {
	var args [2]interface{}
	args[0] = rdRef
	args[1] = userRef

	var argsSerialized []byte

	ret := [2]interface{}{}
	var ret0 []byte
	ret[0] = &ret0
	var ret1 *foundation.Error
	ret[1] = &ret1

	err := proxyctx.Current.Serialize(args, &argsSerialized)
	if err != nil {
		return ret0, err
	}

	res, err := proxyctx.Current.RouteCall(r.Reference, true, true, "DumpUserInfo", argsSerialized, *PrototypeReference)
	if err != nil {
		return ret0, err
	}

	err = proxyctx.Current.Deserialize(res, &ret)
	if err != nil {
		return ret0, err
	}

	if ret1 != nil {
		return ret0, ret1
	}
	return ret0, nil
}

// DumpAllUsers is proxy generated method
func (r *Member) DumpAllUsers(rdRef insolar.Reference) ([]byte, error) {
	var args [1]interface{}
	args[0] = rdRef

	var argsSerialized []byte

	ret := [2]interface{}{}
	var ret0 []byte
	ret[0] = &ret0
	var ret1 *foundation.Error
	ret[1] = &ret1

	err := proxyctx.Current.Serialize(args, &argsSerialized)
	if err != nil {
		return ret0, err
	}

	res, err := proxyctx.Current.RouteCall(r.Reference, true, false, "DumpAllUsers", argsSerialized, *PrototypeReference)
	if err != nil {
		return ret0, err
	}

	err = proxyctx.Current.Deserialize(res, &ret)
	if err != nil {
		return ret0, err
	}

	if ret1 != nil {
		return ret0, ret1
	}
	return ret0, nil
}

// DumpAllUsersNoWait is proxy generated method
func (r *Member) DumpAllUsersNoWait(rdRef insolar.Reference) error {
	var args [1]interface{}
	args[0] = rdRef

	var argsSerialized []byte

	err := proxyctx.Current.Serialize(args, &argsSerialized)
	if err != nil {
		return err
	}

	_, err = proxyctx.Current.RouteCall(r.Reference, false, false, "DumpAllUsers", argsSerialized, *PrototypeReference)
	if err != nil {
		return err
	}

	return nil
}

// DumpAllUsersAsImmutable is proxy generated method
func (r *Member) DumpAllUsersAsImmutable(rdRef insolar.Reference) ([]byte, error) {
	var args [1]interface{}
	args[0] = rdRef

	var argsSerialized []byte

	ret := [2]interface{}{}
	var ret0 []byte
	ret[0] = &ret0
	var ret1 *foundation.Error
	ret[1] = &ret1

	err := proxyctx.Current.Serialize(args, &argsSerialized)
	if err != nil {
		return ret0, err
	}

	res, err := proxyctx.Current.RouteCall(r.Reference, true, true, "DumpAllUsers", argsSerialized, *PrototypeReference)
	if err != nil {
		return ret0, err
	}

	err = proxyctx.Current.Deserialize(res, &ret)
	if err != nil {
		return ret0, err
	}

	if ret1 != nil {
		return ret0, ret1
	}
	return ret0, nil
}

// AddBurnAddressCall is proxy generated method
func (r *Member) AddBurnAddressCall(rdRef insolar.Reference, params []byte) (interface{}, error) {
	var args [2]interface{}
	args[0] = rdRef
	args[1] = params

	var argsSerialized []byte

	ret := [2]interface{}{}
	var ret0 interface{}
	ret[0] = &ret0
	var ret1 *foundation.Error
	ret[1] = &ret1

	err := proxyctx.Current.Serialize(args, &argsSerialized)
	if err != nil {
		return ret0, err
	}

	res, err := proxyctx.Current.RouteCall(r.Reference, true, false, "AddBurnAddressCall", argsSerialized, *PrototypeReference)
	if err != nil {
		return ret0, err
	}

	err = proxyctx.Current.Deserialize(res, &ret)
	if err != nil {
		return ret0, err
	}

	if ret1 != nil {
		return ret0, ret1
	}
	return ret0, nil
}

// AddBurnAddressCallNoWait is proxy generated method
func (r *Member) AddBurnAddressCallNoWait(rdRef insolar.Reference, params []byte) error {
	var args [2]interface{}
	args[0] = rdRef
	args[1] = params

	var argsSerialized []byte

	err := proxyctx.Current.Serialize(args, &argsSerialized)
	if err != nil {
		return err
	}

	_, err = proxyctx.Current.RouteCall(r.Reference, false, false, "AddBurnAddressCall", argsSerialized, *PrototypeReference)
	if err != nil {
		return err
	}

	return nil
}

// AddBurnAddressCallAsImmutable is proxy generated method
func (r *Member) AddBurnAddressCallAsImmutable(rdRef insolar.Reference, params []byte) (interface{}, error) {
	var args [2]interface{}
	args[0] = rdRef
	args[1] = params

	var argsSerialized []byte

	ret := [2]interface{}{}
	var ret0 interface{}
	ret[0] = &ret0
	var ret1 *foundation.Error
	ret[1] = &ret1

	err := proxyctx.Current.Serialize(args, &argsSerialized)
	if err != nil {
		return ret0, err
	}

	res, err := proxyctx.Current.RouteCall(r.Reference, true, true, "AddBurnAddressCall", argsSerialized, *PrototypeReference)
	if err != nil {
		return ret0, err
	}

	err = proxyctx.Current.Deserialize(res, &ret)
	if err != nil {
		return ret0, err
	}

	if ret1 != nil {
		return ret0, ret1
	}
	return ret0, nil
}
