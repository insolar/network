package endpoints

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock"
	"github.com/insolar/insolar/longbits"
	"github.com/insolar/insolar/network/consensus/common/cryptkit"
)

// InboundMock implements Inbound
type InboundMock struct {
	t minimock.Tester

	funcAsByteString          func() (b1 longbits.ByteString)
	inspectFuncAsByteString   func()
	afterAsByteStringCounter  uint64
	beforeAsByteStringCounter uint64
	AsByteStringMock          mInboundMockAsByteString

	funcGetNameAddress          func() (n1 Name)
	inspectFuncGetNameAddress   func()
	afterGetNameAddressCounter  uint64
	beforeGetNameAddressCounter uint64
	GetNameAddressMock          mInboundMockGetNameAddress

	funcGetTransportCert          func() (c1 cryptkit.CertificateHolder)
	inspectFuncGetTransportCert   func()
	afterGetTransportCertCounter  uint64
	beforeGetTransportCertCounter uint64
	GetTransportCertMock          mInboundMockGetTransportCert

	funcGetTransportKey          func() (s1 cryptkit.SignatureKeyHolder)
	inspectFuncGetTransportKey   func()
	afterGetTransportKeyCounter  uint64
	beforeGetTransportKeyCounter uint64
	GetTransportKeyMock          mInboundMockGetTransportKey
}

// NewInboundMock returns a mock for Inbound
func NewInboundMock(t minimock.Tester) *InboundMock {
	m := &InboundMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.AsByteStringMock = mInboundMockAsByteString{mock: m}

	m.GetNameAddressMock = mInboundMockGetNameAddress{mock: m}

	m.GetTransportCertMock = mInboundMockGetTransportCert{mock: m}

	m.GetTransportKeyMock = mInboundMockGetTransportKey{mock: m}

	return m
}

type mInboundMockAsByteString struct {
	mock               *InboundMock
	defaultExpectation *InboundMockAsByteStringExpectation
	expectations       []*InboundMockAsByteStringExpectation
}

// InboundMockAsByteStringExpectation specifies expectation struct of the Inbound.AsByteString
type InboundMockAsByteStringExpectation struct {
	mock *InboundMock

	results *InboundMockAsByteStringResults
	Counter uint64
}

// InboundMockAsByteStringResults contains results of the Inbound.AsByteString
type InboundMockAsByteStringResults struct {
	b1 longbits.ByteString
}

// Expect sets up expected params for Inbound.AsByteString
func (mmAsByteString *mInboundMockAsByteString) Expect() *mInboundMockAsByteString {
	if mmAsByteString.mock.funcAsByteString != nil {
		mmAsByteString.mock.t.Fatalf("InboundMock.AsByteString mock is already set by Set")
	}

	if mmAsByteString.defaultExpectation == nil {
		mmAsByteString.defaultExpectation = &InboundMockAsByteStringExpectation{}
	}

	return mmAsByteString
}

// Inspect accepts an inspector function that has same arguments as the Inbound.AsByteString
func (mmAsByteString *mInboundMockAsByteString) Inspect(f func()) *mInboundMockAsByteString {
	if mmAsByteString.mock.inspectFuncAsByteString != nil {
		mmAsByteString.mock.t.Fatalf("Inspect function is already set for InboundMock.AsByteString")
	}

	mmAsByteString.mock.inspectFuncAsByteString = f

	return mmAsByteString
}

// Return sets up results that will be returned by Inbound.AsByteString
func (mmAsByteString *mInboundMockAsByteString) Return(b1 longbits.ByteString) *InboundMock {
	if mmAsByteString.mock.funcAsByteString != nil {
		mmAsByteString.mock.t.Fatalf("InboundMock.AsByteString mock is already set by Set")
	}

	if mmAsByteString.defaultExpectation == nil {
		mmAsByteString.defaultExpectation = &InboundMockAsByteStringExpectation{mock: mmAsByteString.mock}
	}
	mmAsByteString.defaultExpectation.results = &InboundMockAsByteStringResults{b1}
	return mmAsByteString.mock
}

//Set uses given function f to mock the Inbound.AsByteString method
func (mmAsByteString *mInboundMockAsByteString) Set(f func() (b1 longbits.ByteString)) *InboundMock {
	if mmAsByteString.defaultExpectation != nil {
		mmAsByteString.mock.t.Fatalf("Default expectation is already set for the Inbound.AsByteString method")
	}

	if len(mmAsByteString.expectations) > 0 {
		mmAsByteString.mock.t.Fatalf("Some expectations are already set for the Inbound.AsByteString method")
	}

	mmAsByteString.mock.funcAsByteString = f
	return mmAsByteString.mock
}

// AsByteString implements Inbound
func (mmAsByteString *InboundMock) AsByteString() (b1 longbits.ByteString) {
	mm_atomic.AddUint64(&mmAsByteString.beforeAsByteStringCounter, 1)
	defer mm_atomic.AddUint64(&mmAsByteString.afterAsByteStringCounter, 1)

	if mmAsByteString.inspectFuncAsByteString != nil {
		mmAsByteString.inspectFuncAsByteString()
	}

	if mmAsByteString.AsByteStringMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmAsByteString.AsByteStringMock.defaultExpectation.Counter, 1)

		results := mmAsByteString.AsByteStringMock.defaultExpectation.results
		if results == nil {
			mmAsByteString.t.Fatal("No results are set for the InboundMock.AsByteString")
		}
		return (*results).b1
	}
	if mmAsByteString.funcAsByteString != nil {
		return mmAsByteString.funcAsByteString()
	}
	mmAsByteString.t.Fatalf("Unexpected call to InboundMock.AsByteString.")
	return
}

// AsByteStringAfterCounter returns a count of finished InboundMock.AsByteString invocations
func (mmAsByteString *InboundMock) AsByteStringAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAsByteString.afterAsByteStringCounter)
}

// AsByteStringBeforeCounter returns a count of InboundMock.AsByteString invocations
func (mmAsByteString *InboundMock) AsByteStringBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAsByteString.beforeAsByteStringCounter)
}

// MinimockAsByteStringDone returns true if the count of the AsByteString invocations corresponds
// the number of defined expectations
func (m *InboundMock) MinimockAsByteStringDone() bool {
	for _, e := range m.AsByteStringMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.AsByteStringMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterAsByteStringCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAsByteString != nil && mm_atomic.LoadUint64(&m.afterAsByteStringCounter) < 1 {
		return false
	}
	return true
}

// MinimockAsByteStringInspect logs each unmet expectation
func (m *InboundMock) MinimockAsByteStringInspect() {
	for _, e := range m.AsByteStringMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to InboundMock.AsByteString")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.AsByteStringMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterAsByteStringCounter) < 1 {
		m.t.Error("Expected call to InboundMock.AsByteString")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAsByteString != nil && mm_atomic.LoadUint64(&m.afterAsByteStringCounter) < 1 {
		m.t.Error("Expected call to InboundMock.AsByteString")
	}
}

type mInboundMockGetNameAddress struct {
	mock               *InboundMock
	defaultExpectation *InboundMockGetNameAddressExpectation
	expectations       []*InboundMockGetNameAddressExpectation
}

// InboundMockGetNameAddressExpectation specifies expectation struct of the Inbound.GetNameAddress
type InboundMockGetNameAddressExpectation struct {
	mock *InboundMock

	results *InboundMockGetNameAddressResults
	Counter uint64
}

// InboundMockGetNameAddressResults contains results of the Inbound.GetNameAddress
type InboundMockGetNameAddressResults struct {
	n1 Name
}

// Expect sets up expected params for Inbound.GetNameAddress
func (mmGetNameAddress *mInboundMockGetNameAddress) Expect() *mInboundMockGetNameAddress {
	if mmGetNameAddress.mock.funcGetNameAddress != nil {
		mmGetNameAddress.mock.t.Fatalf("InboundMock.GetNameAddress mock is already set by Set")
	}

	if mmGetNameAddress.defaultExpectation == nil {
		mmGetNameAddress.defaultExpectation = &InboundMockGetNameAddressExpectation{}
	}

	return mmGetNameAddress
}

// Inspect accepts an inspector function that has same arguments as the Inbound.GetNameAddress
func (mmGetNameAddress *mInboundMockGetNameAddress) Inspect(f func()) *mInboundMockGetNameAddress {
	if mmGetNameAddress.mock.inspectFuncGetNameAddress != nil {
		mmGetNameAddress.mock.t.Fatalf("Inspect function is already set for InboundMock.GetNameAddress")
	}

	mmGetNameAddress.mock.inspectFuncGetNameAddress = f

	return mmGetNameAddress
}

// Return sets up results that will be returned by Inbound.GetNameAddress
func (mmGetNameAddress *mInboundMockGetNameAddress) Return(n1 Name) *InboundMock {
	if mmGetNameAddress.mock.funcGetNameAddress != nil {
		mmGetNameAddress.mock.t.Fatalf("InboundMock.GetNameAddress mock is already set by Set")
	}

	if mmGetNameAddress.defaultExpectation == nil {
		mmGetNameAddress.defaultExpectation = &InboundMockGetNameAddressExpectation{mock: mmGetNameAddress.mock}
	}
	mmGetNameAddress.defaultExpectation.results = &InboundMockGetNameAddressResults{n1}
	return mmGetNameAddress.mock
}

//Set uses given function f to mock the Inbound.GetNameAddress method
func (mmGetNameAddress *mInboundMockGetNameAddress) Set(f func() (n1 Name)) *InboundMock {
	if mmGetNameAddress.defaultExpectation != nil {
		mmGetNameAddress.mock.t.Fatalf("Default expectation is already set for the Inbound.GetNameAddress method")
	}

	if len(mmGetNameAddress.expectations) > 0 {
		mmGetNameAddress.mock.t.Fatalf("Some expectations are already set for the Inbound.GetNameAddress method")
	}

	mmGetNameAddress.mock.funcGetNameAddress = f
	return mmGetNameAddress.mock
}

// GetNameAddress implements Inbound
func (mmGetNameAddress *InboundMock) GetNameAddress() (n1 Name) {
	mm_atomic.AddUint64(&mmGetNameAddress.beforeGetNameAddressCounter, 1)
	defer mm_atomic.AddUint64(&mmGetNameAddress.afterGetNameAddressCounter, 1)

	if mmGetNameAddress.inspectFuncGetNameAddress != nil {
		mmGetNameAddress.inspectFuncGetNameAddress()
	}

	if mmGetNameAddress.GetNameAddressMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetNameAddress.GetNameAddressMock.defaultExpectation.Counter, 1)

		results := mmGetNameAddress.GetNameAddressMock.defaultExpectation.results
		if results == nil {
			mmGetNameAddress.t.Fatal("No results are set for the InboundMock.GetNameAddress")
		}
		return (*results).n1
	}
	if mmGetNameAddress.funcGetNameAddress != nil {
		return mmGetNameAddress.funcGetNameAddress()
	}
	mmGetNameAddress.t.Fatalf("Unexpected call to InboundMock.GetNameAddress.")
	return
}

// GetNameAddressAfterCounter returns a count of finished InboundMock.GetNameAddress invocations
func (mmGetNameAddress *InboundMock) GetNameAddressAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetNameAddress.afterGetNameAddressCounter)
}

// GetNameAddressBeforeCounter returns a count of InboundMock.GetNameAddress invocations
func (mmGetNameAddress *InboundMock) GetNameAddressBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetNameAddress.beforeGetNameAddressCounter)
}

// MinimockGetNameAddressDone returns true if the count of the GetNameAddress invocations corresponds
// the number of defined expectations
func (m *InboundMock) MinimockGetNameAddressDone() bool {
	for _, e := range m.GetNameAddressMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetNameAddressMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetNameAddressCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetNameAddress != nil && mm_atomic.LoadUint64(&m.afterGetNameAddressCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetNameAddressInspect logs each unmet expectation
func (m *InboundMock) MinimockGetNameAddressInspect() {
	for _, e := range m.GetNameAddressMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to InboundMock.GetNameAddress")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetNameAddressMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetNameAddressCounter) < 1 {
		m.t.Error("Expected call to InboundMock.GetNameAddress")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetNameAddress != nil && mm_atomic.LoadUint64(&m.afterGetNameAddressCounter) < 1 {
		m.t.Error("Expected call to InboundMock.GetNameAddress")
	}
}

type mInboundMockGetTransportCert struct {
	mock               *InboundMock
	defaultExpectation *InboundMockGetTransportCertExpectation
	expectations       []*InboundMockGetTransportCertExpectation
}

// InboundMockGetTransportCertExpectation specifies expectation struct of the Inbound.GetTransportCert
type InboundMockGetTransportCertExpectation struct {
	mock *InboundMock

	results *InboundMockGetTransportCertResults
	Counter uint64
}

// InboundMockGetTransportCertResults contains results of the Inbound.GetTransportCert
type InboundMockGetTransportCertResults struct {
	c1 cryptkit.CertificateHolder
}

// Expect sets up expected params for Inbound.GetTransportCert
func (mmGetTransportCert *mInboundMockGetTransportCert) Expect() *mInboundMockGetTransportCert {
	if mmGetTransportCert.mock.funcGetTransportCert != nil {
		mmGetTransportCert.mock.t.Fatalf("InboundMock.GetTransportCert mock is already set by Set")
	}

	if mmGetTransportCert.defaultExpectation == nil {
		mmGetTransportCert.defaultExpectation = &InboundMockGetTransportCertExpectation{}
	}

	return mmGetTransportCert
}

// Inspect accepts an inspector function that has same arguments as the Inbound.GetTransportCert
func (mmGetTransportCert *mInboundMockGetTransportCert) Inspect(f func()) *mInboundMockGetTransportCert {
	if mmGetTransportCert.mock.inspectFuncGetTransportCert != nil {
		mmGetTransportCert.mock.t.Fatalf("Inspect function is already set for InboundMock.GetTransportCert")
	}

	mmGetTransportCert.mock.inspectFuncGetTransportCert = f

	return mmGetTransportCert
}

// Return sets up results that will be returned by Inbound.GetTransportCert
func (mmGetTransportCert *mInboundMockGetTransportCert) Return(c1 cryptkit.CertificateHolder) *InboundMock {
	if mmGetTransportCert.mock.funcGetTransportCert != nil {
		mmGetTransportCert.mock.t.Fatalf("InboundMock.GetTransportCert mock is already set by Set")
	}

	if mmGetTransportCert.defaultExpectation == nil {
		mmGetTransportCert.defaultExpectation = &InboundMockGetTransportCertExpectation{mock: mmGetTransportCert.mock}
	}
	mmGetTransportCert.defaultExpectation.results = &InboundMockGetTransportCertResults{c1}
	return mmGetTransportCert.mock
}

//Set uses given function f to mock the Inbound.GetTransportCert method
func (mmGetTransportCert *mInboundMockGetTransportCert) Set(f func() (c1 cryptkit.CertificateHolder)) *InboundMock {
	if mmGetTransportCert.defaultExpectation != nil {
		mmGetTransportCert.mock.t.Fatalf("Default expectation is already set for the Inbound.GetTransportCert method")
	}

	if len(mmGetTransportCert.expectations) > 0 {
		mmGetTransportCert.mock.t.Fatalf("Some expectations are already set for the Inbound.GetTransportCert method")
	}

	mmGetTransportCert.mock.funcGetTransportCert = f
	return mmGetTransportCert.mock
}

// GetTransportCert implements Inbound
func (mmGetTransportCert *InboundMock) GetTransportCert() (c1 cryptkit.CertificateHolder) {
	mm_atomic.AddUint64(&mmGetTransportCert.beforeGetTransportCertCounter, 1)
	defer mm_atomic.AddUint64(&mmGetTransportCert.afterGetTransportCertCounter, 1)

	if mmGetTransportCert.inspectFuncGetTransportCert != nil {
		mmGetTransportCert.inspectFuncGetTransportCert()
	}

	if mmGetTransportCert.GetTransportCertMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetTransportCert.GetTransportCertMock.defaultExpectation.Counter, 1)

		results := mmGetTransportCert.GetTransportCertMock.defaultExpectation.results
		if results == nil {
			mmGetTransportCert.t.Fatal("No results are set for the InboundMock.GetTransportCert")
		}
		return (*results).c1
	}
	if mmGetTransportCert.funcGetTransportCert != nil {
		return mmGetTransportCert.funcGetTransportCert()
	}
	mmGetTransportCert.t.Fatalf("Unexpected call to InboundMock.GetTransportCert.")
	return
}

// GetTransportCertAfterCounter returns a count of finished InboundMock.GetTransportCert invocations
func (mmGetTransportCert *InboundMock) GetTransportCertAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetTransportCert.afterGetTransportCertCounter)
}

// GetTransportCertBeforeCounter returns a count of InboundMock.GetTransportCert invocations
func (mmGetTransportCert *InboundMock) GetTransportCertBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetTransportCert.beforeGetTransportCertCounter)
}

// MinimockGetTransportCertDone returns true if the count of the GetTransportCert invocations corresponds
// the number of defined expectations
func (m *InboundMock) MinimockGetTransportCertDone() bool {
	for _, e := range m.GetTransportCertMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetTransportCertMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetTransportCertCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetTransportCert != nil && mm_atomic.LoadUint64(&m.afterGetTransportCertCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetTransportCertInspect logs each unmet expectation
func (m *InboundMock) MinimockGetTransportCertInspect() {
	for _, e := range m.GetTransportCertMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to InboundMock.GetTransportCert")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetTransportCertMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetTransportCertCounter) < 1 {
		m.t.Error("Expected call to InboundMock.GetTransportCert")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetTransportCert != nil && mm_atomic.LoadUint64(&m.afterGetTransportCertCounter) < 1 {
		m.t.Error("Expected call to InboundMock.GetTransportCert")
	}
}

type mInboundMockGetTransportKey struct {
	mock               *InboundMock
	defaultExpectation *InboundMockGetTransportKeyExpectation
	expectations       []*InboundMockGetTransportKeyExpectation
}

// InboundMockGetTransportKeyExpectation specifies expectation struct of the Inbound.GetTransportKey
type InboundMockGetTransportKeyExpectation struct {
	mock *InboundMock

	results *InboundMockGetTransportKeyResults
	Counter uint64
}

// InboundMockGetTransportKeyResults contains results of the Inbound.GetTransportKey
type InboundMockGetTransportKeyResults struct {
	s1 cryptkit.SignatureKeyHolder
}

// Expect sets up expected params for Inbound.GetTransportKey
func (mmGetTransportKey *mInboundMockGetTransportKey) Expect() *mInboundMockGetTransportKey {
	if mmGetTransportKey.mock.funcGetTransportKey != nil {
		mmGetTransportKey.mock.t.Fatalf("InboundMock.GetTransportKey mock is already set by Set")
	}

	if mmGetTransportKey.defaultExpectation == nil {
		mmGetTransportKey.defaultExpectation = &InboundMockGetTransportKeyExpectation{}
	}

	return mmGetTransportKey
}

// Inspect accepts an inspector function that has same arguments as the Inbound.GetTransportKey
func (mmGetTransportKey *mInboundMockGetTransportKey) Inspect(f func()) *mInboundMockGetTransportKey {
	if mmGetTransportKey.mock.inspectFuncGetTransportKey != nil {
		mmGetTransportKey.mock.t.Fatalf("Inspect function is already set for InboundMock.GetTransportKey")
	}

	mmGetTransportKey.mock.inspectFuncGetTransportKey = f

	return mmGetTransportKey
}

// Return sets up results that will be returned by Inbound.GetTransportKey
func (mmGetTransportKey *mInboundMockGetTransportKey) Return(s1 cryptkit.SignatureKeyHolder) *InboundMock {
	if mmGetTransportKey.mock.funcGetTransportKey != nil {
		mmGetTransportKey.mock.t.Fatalf("InboundMock.GetTransportKey mock is already set by Set")
	}

	if mmGetTransportKey.defaultExpectation == nil {
		mmGetTransportKey.defaultExpectation = &InboundMockGetTransportKeyExpectation{mock: mmGetTransportKey.mock}
	}
	mmGetTransportKey.defaultExpectation.results = &InboundMockGetTransportKeyResults{s1}
	return mmGetTransportKey.mock
}

//Set uses given function f to mock the Inbound.GetTransportKey method
func (mmGetTransportKey *mInboundMockGetTransportKey) Set(f func() (s1 cryptkit.SignatureKeyHolder)) *InboundMock {
	if mmGetTransportKey.defaultExpectation != nil {
		mmGetTransportKey.mock.t.Fatalf("Default expectation is already set for the Inbound.GetTransportKey method")
	}

	if len(mmGetTransportKey.expectations) > 0 {
		mmGetTransportKey.mock.t.Fatalf("Some expectations are already set for the Inbound.GetTransportKey method")
	}

	mmGetTransportKey.mock.funcGetTransportKey = f
	return mmGetTransportKey.mock
}

// GetTransportKey implements Inbound
func (mmGetTransportKey *InboundMock) GetTransportKey() (s1 cryptkit.SignatureKeyHolder) {
	mm_atomic.AddUint64(&mmGetTransportKey.beforeGetTransportKeyCounter, 1)
	defer mm_atomic.AddUint64(&mmGetTransportKey.afterGetTransportKeyCounter, 1)

	if mmGetTransportKey.inspectFuncGetTransportKey != nil {
		mmGetTransportKey.inspectFuncGetTransportKey()
	}

	if mmGetTransportKey.GetTransportKeyMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetTransportKey.GetTransportKeyMock.defaultExpectation.Counter, 1)

		results := mmGetTransportKey.GetTransportKeyMock.defaultExpectation.results
		if results == nil {
			mmGetTransportKey.t.Fatal("No results are set for the InboundMock.GetTransportKey")
		}
		return (*results).s1
	}
	if mmGetTransportKey.funcGetTransportKey != nil {
		return mmGetTransportKey.funcGetTransportKey()
	}
	mmGetTransportKey.t.Fatalf("Unexpected call to InboundMock.GetTransportKey.")
	return
}

// GetTransportKeyAfterCounter returns a count of finished InboundMock.GetTransportKey invocations
func (mmGetTransportKey *InboundMock) GetTransportKeyAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetTransportKey.afterGetTransportKeyCounter)
}

// GetTransportKeyBeforeCounter returns a count of InboundMock.GetTransportKey invocations
func (mmGetTransportKey *InboundMock) GetTransportKeyBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetTransportKey.beforeGetTransportKeyCounter)
}

// MinimockGetTransportKeyDone returns true if the count of the GetTransportKey invocations corresponds
// the number of defined expectations
func (m *InboundMock) MinimockGetTransportKeyDone() bool {
	for _, e := range m.GetTransportKeyMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetTransportKeyMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetTransportKeyCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetTransportKey != nil && mm_atomic.LoadUint64(&m.afterGetTransportKeyCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetTransportKeyInspect logs each unmet expectation
func (m *InboundMock) MinimockGetTransportKeyInspect() {
	for _, e := range m.GetTransportKeyMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to InboundMock.GetTransportKey")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetTransportKeyMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetTransportKeyCounter) < 1 {
		m.t.Error("Expected call to InboundMock.GetTransportKey")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetTransportKey != nil && mm_atomic.LoadUint64(&m.afterGetTransportKeyCounter) < 1 {
		m.t.Error("Expected call to InboundMock.GetTransportKey")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *InboundMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockAsByteStringInspect()

		m.MinimockGetNameAddressInspect()

		m.MinimockGetTransportCertInspect()

		m.MinimockGetTransportKeyInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *InboundMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *InboundMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockAsByteStringDone() &&
		m.MinimockGetNameAddressDone() &&
		m.MinimockGetTransportCertDone() &&
		m.MinimockGetTransportKeyDone()
}
