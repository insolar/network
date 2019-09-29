package artifacts

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock"
	"github.com/insolar/insolar/insolar"
)

// ObjectDescriptorMock implements ObjectDescriptor
type ObjectDescriptorMock struct {
	t minimock.Tester

	funcEarliestRequestID          func() (ip1 *insolar.ID)
	inspectFuncEarliestRequestID   func()
	afterEarliestRequestIDCounter  uint64
	beforeEarliestRequestIDCounter uint64
	EarliestRequestIDMock          mObjectDescriptorMockEarliestRequestID

	funcHeadRef          func() (rp1 *insolar.Reference)
	inspectFuncHeadRef   func()
	afterHeadRefCounter  uint64
	beforeHeadRefCounter uint64
	HeadRefMock          mObjectDescriptorMockHeadRef

	funcMemory          func() (ba1 []byte)
	inspectFuncMemory   func()
	afterMemoryCounter  uint64
	beforeMemoryCounter uint64
	MemoryMock          mObjectDescriptorMockMemory

	funcParent          func() (rp1 *insolar.Reference)
	inspectFuncParent   func()
	afterParentCounter  uint64
	beforeParentCounter uint64
	ParentMock          mObjectDescriptorMockParent

	funcPrototype          func() (rp1 *insolar.Reference, err error)
	inspectFuncPrototype   func()
	afterPrototypeCounter  uint64
	beforePrototypeCounter uint64
	PrototypeMock          mObjectDescriptorMockPrototype

	funcStateID          func() (ip1 *insolar.ID)
	inspectFuncStateID   func()
	afterStateIDCounter  uint64
	beforeStateIDCounter uint64
	StateIDMock          mObjectDescriptorMockStateID
}

// NewObjectDescriptorMock returns a mock for ObjectDescriptor
func NewObjectDescriptorMock(t minimock.Tester) *ObjectDescriptorMock {
	m := &ObjectDescriptorMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.EarliestRequestIDMock = mObjectDescriptorMockEarliestRequestID{mock: m}

	m.HeadRefMock = mObjectDescriptorMockHeadRef{mock: m}

	m.MemoryMock = mObjectDescriptorMockMemory{mock: m}

	m.ParentMock = mObjectDescriptorMockParent{mock: m}

	m.PrototypeMock = mObjectDescriptorMockPrototype{mock: m}

	m.StateIDMock = mObjectDescriptorMockStateID{mock: m}

	return m
}

type mObjectDescriptorMockEarliestRequestID struct {
	mock               *ObjectDescriptorMock
	defaultExpectation *ObjectDescriptorMockEarliestRequestIDExpectation
	expectations       []*ObjectDescriptorMockEarliestRequestIDExpectation
}

// ObjectDescriptorMockEarliestRequestIDExpectation specifies expectation struct of the ObjectDescriptor.EarliestRequestID
type ObjectDescriptorMockEarliestRequestIDExpectation struct {
	mock *ObjectDescriptorMock

	results *ObjectDescriptorMockEarliestRequestIDResults
	Counter uint64
}

// ObjectDescriptorMockEarliestRequestIDResults contains results of the ObjectDescriptor.EarliestRequestID
type ObjectDescriptorMockEarliestRequestIDResults struct {
	ip1 *insolar.ID
}

// Expect sets up expected params for ObjectDescriptor.EarliestRequestID
func (mmEarliestRequestID *mObjectDescriptorMockEarliestRequestID) Expect() *mObjectDescriptorMockEarliestRequestID {
	if mmEarliestRequestID.mock.funcEarliestRequestID != nil {
		mmEarliestRequestID.mock.t.Fatalf("ObjectDescriptorMock.EarliestRequestID mock is already set by Set")
	}

	if mmEarliestRequestID.defaultExpectation == nil {
		mmEarliestRequestID.defaultExpectation = &ObjectDescriptorMockEarliestRequestIDExpectation{}
	}

	return mmEarliestRequestID
}

// Inspect accepts an inspector function that has same arguments as the ObjectDescriptor.EarliestRequestID
func (mmEarliestRequestID *mObjectDescriptorMockEarliestRequestID) Inspect(f func()) *mObjectDescriptorMockEarliestRequestID {
	if mmEarliestRequestID.mock.inspectFuncEarliestRequestID != nil {
		mmEarliestRequestID.mock.t.Fatalf("Inspect function is already set for ObjectDescriptorMock.EarliestRequestID")
	}

	mmEarliestRequestID.mock.inspectFuncEarliestRequestID = f

	return mmEarliestRequestID
}

// Return sets up results that will be returned by ObjectDescriptor.EarliestRequestID
func (mmEarliestRequestID *mObjectDescriptorMockEarliestRequestID) Return(ip1 *insolar.ID) *ObjectDescriptorMock {
	if mmEarliestRequestID.mock.funcEarliestRequestID != nil {
		mmEarliestRequestID.mock.t.Fatalf("ObjectDescriptorMock.EarliestRequestID mock is already set by Set")
	}

	if mmEarliestRequestID.defaultExpectation == nil {
		mmEarliestRequestID.defaultExpectation = &ObjectDescriptorMockEarliestRequestIDExpectation{mock: mmEarliestRequestID.mock}
	}
	mmEarliestRequestID.defaultExpectation.results = &ObjectDescriptorMockEarliestRequestIDResults{ip1}
	return mmEarliestRequestID.mock
}

//Set uses given function f to mock the ObjectDescriptor.EarliestRequestID method
func (mmEarliestRequestID *mObjectDescriptorMockEarliestRequestID) Set(f func() (ip1 *insolar.ID)) *ObjectDescriptorMock {
	if mmEarliestRequestID.defaultExpectation != nil {
		mmEarliestRequestID.mock.t.Fatalf("Default expectation is already set for the ObjectDescriptor.EarliestRequestID method")
	}

	if len(mmEarliestRequestID.expectations) > 0 {
		mmEarliestRequestID.mock.t.Fatalf("Some expectations are already set for the ObjectDescriptor.EarliestRequestID method")
	}

	mmEarliestRequestID.mock.funcEarliestRequestID = f
	return mmEarliestRequestID.mock
}

// EarliestRequestID implements ObjectDescriptor
func (mmEarliestRequestID *ObjectDescriptorMock) EarliestRequestID() (ip1 *insolar.ID) {
	mm_atomic.AddUint64(&mmEarliestRequestID.beforeEarliestRequestIDCounter, 1)
	defer mm_atomic.AddUint64(&mmEarliestRequestID.afterEarliestRequestIDCounter, 1)

	if mmEarliestRequestID.inspectFuncEarliestRequestID != nil {
		mmEarliestRequestID.inspectFuncEarliestRequestID()
	}

	if mmEarliestRequestID.EarliestRequestIDMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmEarliestRequestID.EarliestRequestIDMock.defaultExpectation.Counter, 1)

		results := mmEarliestRequestID.EarliestRequestIDMock.defaultExpectation.results
		if results == nil {
			mmEarliestRequestID.t.Fatal("No results are set for the ObjectDescriptorMock.EarliestRequestID")
		}
		return (*results).ip1
	}
	if mmEarliestRequestID.funcEarliestRequestID != nil {
		return mmEarliestRequestID.funcEarliestRequestID()
	}
	mmEarliestRequestID.t.Fatalf("Unexpected call to ObjectDescriptorMock.EarliestRequestID.")
	return
}

// EarliestRequestIDAfterCounter returns a count of finished ObjectDescriptorMock.EarliestRequestID invocations
func (mmEarliestRequestID *ObjectDescriptorMock) EarliestRequestIDAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmEarliestRequestID.afterEarliestRequestIDCounter)
}

// EarliestRequestIDBeforeCounter returns a count of ObjectDescriptorMock.EarliestRequestID invocations
func (mmEarliestRequestID *ObjectDescriptorMock) EarliestRequestIDBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmEarliestRequestID.beforeEarliestRequestIDCounter)
}

// MinimockEarliestRequestIDDone returns true if the count of the EarliestRequestID invocations corresponds
// the number of defined expectations
func (m *ObjectDescriptorMock) MinimockEarliestRequestIDDone() bool {
	for _, e := range m.EarliestRequestIDMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.EarliestRequestIDMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterEarliestRequestIDCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcEarliestRequestID != nil && mm_atomic.LoadUint64(&m.afterEarliestRequestIDCounter) < 1 {
		return false
	}
	return true
}

// MinimockEarliestRequestIDInspect logs each unmet expectation
func (m *ObjectDescriptorMock) MinimockEarliestRequestIDInspect() {
	for _, e := range m.EarliestRequestIDMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to ObjectDescriptorMock.EarliestRequestID")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.EarliestRequestIDMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterEarliestRequestIDCounter) < 1 {
		m.t.Error("Expected call to ObjectDescriptorMock.EarliestRequestID")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcEarliestRequestID != nil && mm_atomic.LoadUint64(&m.afterEarliestRequestIDCounter) < 1 {
		m.t.Error("Expected call to ObjectDescriptorMock.EarliestRequestID")
	}
}

type mObjectDescriptorMockHeadRef struct {
	mock               *ObjectDescriptorMock
	defaultExpectation *ObjectDescriptorMockHeadRefExpectation
	expectations       []*ObjectDescriptorMockHeadRefExpectation
}

// ObjectDescriptorMockHeadRefExpectation specifies expectation struct of the ObjectDescriptor.HeadRef
type ObjectDescriptorMockHeadRefExpectation struct {
	mock *ObjectDescriptorMock

	results *ObjectDescriptorMockHeadRefResults
	Counter uint64
}

// ObjectDescriptorMockHeadRefResults contains results of the ObjectDescriptor.HeadRef
type ObjectDescriptorMockHeadRefResults struct {
	rp1 *insolar.Reference
}

// Expect sets up expected params for ObjectDescriptor.HeadRef
func (mmHeadRef *mObjectDescriptorMockHeadRef) Expect() *mObjectDescriptorMockHeadRef {
	if mmHeadRef.mock.funcHeadRef != nil {
		mmHeadRef.mock.t.Fatalf("ObjectDescriptorMock.HeadRef mock is already set by Set")
	}

	if mmHeadRef.defaultExpectation == nil {
		mmHeadRef.defaultExpectation = &ObjectDescriptorMockHeadRefExpectation{}
	}

	return mmHeadRef
}

// Inspect accepts an inspector function that has same arguments as the ObjectDescriptor.HeadRef
func (mmHeadRef *mObjectDescriptorMockHeadRef) Inspect(f func()) *mObjectDescriptorMockHeadRef {
	if mmHeadRef.mock.inspectFuncHeadRef != nil {
		mmHeadRef.mock.t.Fatalf("Inspect function is already set for ObjectDescriptorMock.HeadRef")
	}

	mmHeadRef.mock.inspectFuncHeadRef = f

	return mmHeadRef
}

// Return sets up results that will be returned by ObjectDescriptor.HeadRef
func (mmHeadRef *mObjectDescriptorMockHeadRef) Return(rp1 *insolar.Reference) *ObjectDescriptorMock {
	if mmHeadRef.mock.funcHeadRef != nil {
		mmHeadRef.mock.t.Fatalf("ObjectDescriptorMock.HeadRef mock is already set by Set")
	}

	if mmHeadRef.defaultExpectation == nil {
		mmHeadRef.defaultExpectation = &ObjectDescriptorMockHeadRefExpectation{mock: mmHeadRef.mock}
	}
	mmHeadRef.defaultExpectation.results = &ObjectDescriptorMockHeadRefResults{rp1}
	return mmHeadRef.mock
}

//Set uses given function f to mock the ObjectDescriptor.HeadRef method
func (mmHeadRef *mObjectDescriptorMockHeadRef) Set(f func() (rp1 *insolar.Reference)) *ObjectDescriptorMock {
	if mmHeadRef.defaultExpectation != nil {
		mmHeadRef.mock.t.Fatalf("Default expectation is already set for the ObjectDescriptor.HeadRef method")
	}

	if len(mmHeadRef.expectations) > 0 {
		mmHeadRef.mock.t.Fatalf("Some expectations are already set for the ObjectDescriptor.HeadRef method")
	}

	mmHeadRef.mock.funcHeadRef = f
	return mmHeadRef.mock
}

// HeadRef implements ObjectDescriptor
func (mmHeadRef *ObjectDescriptorMock) HeadRef() (rp1 *insolar.Reference) {
	mm_atomic.AddUint64(&mmHeadRef.beforeHeadRefCounter, 1)
	defer mm_atomic.AddUint64(&mmHeadRef.afterHeadRefCounter, 1)

	if mmHeadRef.inspectFuncHeadRef != nil {
		mmHeadRef.inspectFuncHeadRef()
	}

	if mmHeadRef.HeadRefMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmHeadRef.HeadRefMock.defaultExpectation.Counter, 1)

		results := mmHeadRef.HeadRefMock.defaultExpectation.results
		if results == nil {
			mmHeadRef.t.Fatal("No results are set for the ObjectDescriptorMock.HeadRef")
		}
		return (*results).rp1
	}
	if mmHeadRef.funcHeadRef != nil {
		return mmHeadRef.funcHeadRef()
	}
	mmHeadRef.t.Fatalf("Unexpected call to ObjectDescriptorMock.HeadRef.")
	return
}

// HeadRefAfterCounter returns a count of finished ObjectDescriptorMock.HeadRef invocations
func (mmHeadRef *ObjectDescriptorMock) HeadRefAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmHeadRef.afterHeadRefCounter)
}

// HeadRefBeforeCounter returns a count of ObjectDescriptorMock.HeadRef invocations
func (mmHeadRef *ObjectDescriptorMock) HeadRefBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmHeadRef.beforeHeadRefCounter)
}

// MinimockHeadRefDone returns true if the count of the HeadRef invocations corresponds
// the number of defined expectations
func (m *ObjectDescriptorMock) MinimockHeadRefDone() bool {
	for _, e := range m.HeadRefMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.HeadRefMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterHeadRefCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcHeadRef != nil && mm_atomic.LoadUint64(&m.afterHeadRefCounter) < 1 {
		return false
	}
	return true
}

// MinimockHeadRefInspect logs each unmet expectation
func (m *ObjectDescriptorMock) MinimockHeadRefInspect() {
	for _, e := range m.HeadRefMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to ObjectDescriptorMock.HeadRef")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.HeadRefMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterHeadRefCounter) < 1 {
		m.t.Error("Expected call to ObjectDescriptorMock.HeadRef")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcHeadRef != nil && mm_atomic.LoadUint64(&m.afterHeadRefCounter) < 1 {
		m.t.Error("Expected call to ObjectDescriptorMock.HeadRef")
	}
}

type mObjectDescriptorMockMemory struct {
	mock               *ObjectDescriptorMock
	defaultExpectation *ObjectDescriptorMockMemoryExpectation
	expectations       []*ObjectDescriptorMockMemoryExpectation
}

// ObjectDescriptorMockMemoryExpectation specifies expectation struct of the ObjectDescriptor.Memory
type ObjectDescriptorMockMemoryExpectation struct {
	mock *ObjectDescriptorMock

	results *ObjectDescriptorMockMemoryResults
	Counter uint64
}

// ObjectDescriptorMockMemoryResults contains results of the ObjectDescriptor.Memory
type ObjectDescriptorMockMemoryResults struct {
	ba1 []byte
}

// Expect sets up expected params for ObjectDescriptor.Memory
func (mmMemory *mObjectDescriptorMockMemory) Expect() *mObjectDescriptorMockMemory {
	if mmMemory.mock.funcMemory != nil {
		mmMemory.mock.t.Fatalf("ObjectDescriptorMock.Memory mock is already set by Set")
	}

	if mmMemory.defaultExpectation == nil {
		mmMemory.defaultExpectation = &ObjectDescriptorMockMemoryExpectation{}
	}

	return mmMemory
}

// Inspect accepts an inspector function that has same arguments as the ObjectDescriptor.Memory
func (mmMemory *mObjectDescriptorMockMemory) Inspect(f func()) *mObjectDescriptorMockMemory {
	if mmMemory.mock.inspectFuncMemory != nil {
		mmMemory.mock.t.Fatalf("Inspect function is already set for ObjectDescriptorMock.Memory")
	}

	mmMemory.mock.inspectFuncMemory = f

	return mmMemory
}

// Return sets up results that will be returned by ObjectDescriptor.Memory
func (mmMemory *mObjectDescriptorMockMemory) Return(ba1 []byte) *ObjectDescriptorMock {
	if mmMemory.mock.funcMemory != nil {
		mmMemory.mock.t.Fatalf("ObjectDescriptorMock.Memory mock is already set by Set")
	}

	if mmMemory.defaultExpectation == nil {
		mmMemory.defaultExpectation = &ObjectDescriptorMockMemoryExpectation{mock: mmMemory.mock}
	}
	mmMemory.defaultExpectation.results = &ObjectDescriptorMockMemoryResults{ba1}
	return mmMemory.mock
}

//Set uses given function f to mock the ObjectDescriptor.Memory method
func (mmMemory *mObjectDescriptorMockMemory) Set(f func() (ba1 []byte)) *ObjectDescriptorMock {
	if mmMemory.defaultExpectation != nil {
		mmMemory.mock.t.Fatalf("Default expectation is already set for the ObjectDescriptor.Memory method")
	}

	if len(mmMemory.expectations) > 0 {
		mmMemory.mock.t.Fatalf("Some expectations are already set for the ObjectDescriptor.Memory method")
	}

	mmMemory.mock.funcMemory = f
	return mmMemory.mock
}

// Memory implements ObjectDescriptor
func (mmMemory *ObjectDescriptorMock) Memory() (ba1 []byte) {
	mm_atomic.AddUint64(&mmMemory.beforeMemoryCounter, 1)
	defer mm_atomic.AddUint64(&mmMemory.afterMemoryCounter, 1)

	if mmMemory.inspectFuncMemory != nil {
		mmMemory.inspectFuncMemory()
	}

	if mmMemory.MemoryMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmMemory.MemoryMock.defaultExpectation.Counter, 1)

		results := mmMemory.MemoryMock.defaultExpectation.results
		if results == nil {
			mmMemory.t.Fatal("No results are set for the ObjectDescriptorMock.Memory")
		}
		return (*results).ba1
	}
	if mmMemory.funcMemory != nil {
		return mmMemory.funcMemory()
	}
	mmMemory.t.Fatalf("Unexpected call to ObjectDescriptorMock.Memory.")
	return
}

// MemoryAfterCounter returns a count of finished ObjectDescriptorMock.Memory invocations
func (mmMemory *ObjectDescriptorMock) MemoryAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmMemory.afterMemoryCounter)
}

// MemoryBeforeCounter returns a count of ObjectDescriptorMock.Memory invocations
func (mmMemory *ObjectDescriptorMock) MemoryBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmMemory.beforeMemoryCounter)
}

// MinimockMemoryDone returns true if the count of the Memory invocations corresponds
// the number of defined expectations
func (m *ObjectDescriptorMock) MinimockMemoryDone() bool {
	for _, e := range m.MemoryMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.MemoryMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterMemoryCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcMemory != nil && mm_atomic.LoadUint64(&m.afterMemoryCounter) < 1 {
		return false
	}
	return true
}

// MinimockMemoryInspect logs each unmet expectation
func (m *ObjectDescriptorMock) MinimockMemoryInspect() {
	for _, e := range m.MemoryMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to ObjectDescriptorMock.Memory")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.MemoryMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterMemoryCounter) < 1 {
		m.t.Error("Expected call to ObjectDescriptorMock.Memory")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcMemory != nil && mm_atomic.LoadUint64(&m.afterMemoryCounter) < 1 {
		m.t.Error("Expected call to ObjectDescriptorMock.Memory")
	}
}

type mObjectDescriptorMockParent struct {
	mock               *ObjectDescriptorMock
	defaultExpectation *ObjectDescriptorMockParentExpectation
	expectations       []*ObjectDescriptorMockParentExpectation
}

// ObjectDescriptorMockParentExpectation specifies expectation struct of the ObjectDescriptor.Parent
type ObjectDescriptorMockParentExpectation struct {
	mock *ObjectDescriptorMock

	results *ObjectDescriptorMockParentResults
	Counter uint64
}

// ObjectDescriptorMockParentResults contains results of the ObjectDescriptor.Parent
type ObjectDescriptorMockParentResults struct {
	rp1 *insolar.Reference
}

// Expect sets up expected params for ObjectDescriptor.Parent
func (mmParent *mObjectDescriptorMockParent) Expect() *mObjectDescriptorMockParent {
	if mmParent.mock.funcParent != nil {
		mmParent.mock.t.Fatalf("ObjectDescriptorMock.Parent mock is already set by Set")
	}

	if mmParent.defaultExpectation == nil {
		mmParent.defaultExpectation = &ObjectDescriptorMockParentExpectation{}
	}

	return mmParent
}

// Inspect accepts an inspector function that has same arguments as the ObjectDescriptor.Parent
func (mmParent *mObjectDescriptorMockParent) Inspect(f func()) *mObjectDescriptorMockParent {
	if mmParent.mock.inspectFuncParent != nil {
		mmParent.mock.t.Fatalf("Inspect function is already set for ObjectDescriptorMock.Parent")
	}

	mmParent.mock.inspectFuncParent = f

	return mmParent
}

// Return sets up results that will be returned by ObjectDescriptor.Parent
func (mmParent *mObjectDescriptorMockParent) Return(rp1 *insolar.Reference) *ObjectDescriptorMock {
	if mmParent.mock.funcParent != nil {
		mmParent.mock.t.Fatalf("ObjectDescriptorMock.Parent mock is already set by Set")
	}

	if mmParent.defaultExpectation == nil {
		mmParent.defaultExpectation = &ObjectDescriptorMockParentExpectation{mock: mmParent.mock}
	}
	mmParent.defaultExpectation.results = &ObjectDescriptorMockParentResults{rp1}
	return mmParent.mock
}

//Set uses given function f to mock the ObjectDescriptor.Parent method
func (mmParent *mObjectDescriptorMockParent) Set(f func() (rp1 *insolar.Reference)) *ObjectDescriptorMock {
	if mmParent.defaultExpectation != nil {
		mmParent.mock.t.Fatalf("Default expectation is already set for the ObjectDescriptor.Parent method")
	}

	if len(mmParent.expectations) > 0 {
		mmParent.mock.t.Fatalf("Some expectations are already set for the ObjectDescriptor.Parent method")
	}

	mmParent.mock.funcParent = f
	return mmParent.mock
}

// Parent implements ObjectDescriptor
func (mmParent *ObjectDescriptorMock) Parent() (rp1 *insolar.Reference) {
	mm_atomic.AddUint64(&mmParent.beforeParentCounter, 1)
	defer mm_atomic.AddUint64(&mmParent.afterParentCounter, 1)

	if mmParent.inspectFuncParent != nil {
		mmParent.inspectFuncParent()
	}

	if mmParent.ParentMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmParent.ParentMock.defaultExpectation.Counter, 1)

		results := mmParent.ParentMock.defaultExpectation.results
		if results == nil {
			mmParent.t.Fatal("No results are set for the ObjectDescriptorMock.Parent")
		}
		return (*results).rp1
	}
	if mmParent.funcParent != nil {
		return mmParent.funcParent()
	}
	mmParent.t.Fatalf("Unexpected call to ObjectDescriptorMock.Parent.")
	return
}

// ParentAfterCounter returns a count of finished ObjectDescriptorMock.Parent invocations
func (mmParent *ObjectDescriptorMock) ParentAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmParent.afterParentCounter)
}

// ParentBeforeCounter returns a count of ObjectDescriptorMock.Parent invocations
func (mmParent *ObjectDescriptorMock) ParentBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmParent.beforeParentCounter)
}

// MinimockParentDone returns true if the count of the Parent invocations corresponds
// the number of defined expectations
func (m *ObjectDescriptorMock) MinimockParentDone() bool {
	for _, e := range m.ParentMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ParentMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterParentCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcParent != nil && mm_atomic.LoadUint64(&m.afterParentCounter) < 1 {
		return false
	}
	return true
}

// MinimockParentInspect logs each unmet expectation
func (m *ObjectDescriptorMock) MinimockParentInspect() {
	for _, e := range m.ParentMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to ObjectDescriptorMock.Parent")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ParentMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterParentCounter) < 1 {
		m.t.Error("Expected call to ObjectDescriptorMock.Parent")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcParent != nil && mm_atomic.LoadUint64(&m.afterParentCounter) < 1 {
		m.t.Error("Expected call to ObjectDescriptorMock.Parent")
	}
}

type mObjectDescriptorMockPrototype struct {
	mock               *ObjectDescriptorMock
	defaultExpectation *ObjectDescriptorMockPrototypeExpectation
	expectations       []*ObjectDescriptorMockPrototypeExpectation
}

// ObjectDescriptorMockPrototypeExpectation specifies expectation struct of the ObjectDescriptor.Prototype
type ObjectDescriptorMockPrototypeExpectation struct {
	mock *ObjectDescriptorMock

	results *ObjectDescriptorMockPrototypeResults
	Counter uint64
}

// ObjectDescriptorMockPrototypeResults contains results of the ObjectDescriptor.Prototype
type ObjectDescriptorMockPrototypeResults struct {
	rp1 *insolar.Reference
	err error
}

// Expect sets up expected params for ObjectDescriptor.Prototype
func (mmPrototype *mObjectDescriptorMockPrototype) Expect() *mObjectDescriptorMockPrototype {
	if mmPrototype.mock.funcPrototype != nil {
		mmPrototype.mock.t.Fatalf("ObjectDescriptorMock.Prototype mock is already set by Set")
	}

	if mmPrototype.defaultExpectation == nil {
		mmPrototype.defaultExpectation = &ObjectDescriptorMockPrototypeExpectation{}
	}

	return mmPrototype
}

// Inspect accepts an inspector function that has same arguments as the ObjectDescriptor.Prototype
func (mmPrototype *mObjectDescriptorMockPrototype) Inspect(f func()) *mObjectDescriptorMockPrototype {
	if mmPrototype.mock.inspectFuncPrototype != nil {
		mmPrototype.mock.t.Fatalf("Inspect function is already set for ObjectDescriptorMock.Prototype")
	}

	mmPrototype.mock.inspectFuncPrototype = f

	return mmPrototype
}

// Return sets up results that will be returned by ObjectDescriptor.Prototype
func (mmPrototype *mObjectDescriptorMockPrototype) Return(rp1 *insolar.Reference, err error) *ObjectDescriptorMock {
	if mmPrototype.mock.funcPrototype != nil {
		mmPrototype.mock.t.Fatalf("ObjectDescriptorMock.Prototype mock is already set by Set")
	}

	if mmPrototype.defaultExpectation == nil {
		mmPrototype.defaultExpectation = &ObjectDescriptorMockPrototypeExpectation{mock: mmPrototype.mock}
	}
	mmPrototype.defaultExpectation.results = &ObjectDescriptorMockPrototypeResults{rp1, err}
	return mmPrototype.mock
}

//Set uses given function f to mock the ObjectDescriptor.Prototype method
func (mmPrototype *mObjectDescriptorMockPrototype) Set(f func() (rp1 *insolar.Reference, err error)) *ObjectDescriptorMock {
	if mmPrototype.defaultExpectation != nil {
		mmPrototype.mock.t.Fatalf("Default expectation is already set for the ObjectDescriptor.Prototype method")
	}

	if len(mmPrototype.expectations) > 0 {
		mmPrototype.mock.t.Fatalf("Some expectations are already set for the ObjectDescriptor.Prototype method")
	}

	mmPrototype.mock.funcPrototype = f
	return mmPrototype.mock
}

// Prototype implements ObjectDescriptor
func (mmPrototype *ObjectDescriptorMock) Prototype() (rp1 *insolar.Reference, err error) {
	mm_atomic.AddUint64(&mmPrototype.beforePrototypeCounter, 1)
	defer mm_atomic.AddUint64(&mmPrototype.afterPrototypeCounter, 1)

	if mmPrototype.inspectFuncPrototype != nil {
		mmPrototype.inspectFuncPrototype()
	}

	if mmPrototype.PrototypeMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmPrototype.PrototypeMock.defaultExpectation.Counter, 1)

		results := mmPrototype.PrototypeMock.defaultExpectation.results
		if results == nil {
			mmPrototype.t.Fatal("No results are set for the ObjectDescriptorMock.Prototype")
		}
		return (*results).rp1, (*results).err
	}
	if mmPrototype.funcPrototype != nil {
		return mmPrototype.funcPrototype()
	}
	mmPrototype.t.Fatalf("Unexpected call to ObjectDescriptorMock.Prototype.")
	return
}

// PrototypeAfterCounter returns a count of finished ObjectDescriptorMock.Prototype invocations
func (mmPrototype *ObjectDescriptorMock) PrototypeAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmPrototype.afterPrototypeCounter)
}

// PrototypeBeforeCounter returns a count of ObjectDescriptorMock.Prototype invocations
func (mmPrototype *ObjectDescriptorMock) PrototypeBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmPrototype.beforePrototypeCounter)
}

// MinimockPrototypeDone returns true if the count of the Prototype invocations corresponds
// the number of defined expectations
func (m *ObjectDescriptorMock) MinimockPrototypeDone() bool {
	for _, e := range m.PrototypeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.PrototypeMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterPrototypeCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcPrototype != nil && mm_atomic.LoadUint64(&m.afterPrototypeCounter) < 1 {
		return false
	}
	return true
}

// MinimockPrototypeInspect logs each unmet expectation
func (m *ObjectDescriptorMock) MinimockPrototypeInspect() {
	for _, e := range m.PrototypeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to ObjectDescriptorMock.Prototype")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.PrototypeMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterPrototypeCounter) < 1 {
		m.t.Error("Expected call to ObjectDescriptorMock.Prototype")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcPrototype != nil && mm_atomic.LoadUint64(&m.afterPrototypeCounter) < 1 {
		m.t.Error("Expected call to ObjectDescriptorMock.Prototype")
	}
}

type mObjectDescriptorMockStateID struct {
	mock               *ObjectDescriptorMock
	defaultExpectation *ObjectDescriptorMockStateIDExpectation
	expectations       []*ObjectDescriptorMockStateIDExpectation
}

// ObjectDescriptorMockStateIDExpectation specifies expectation struct of the ObjectDescriptor.StateID
type ObjectDescriptorMockStateIDExpectation struct {
	mock *ObjectDescriptorMock

	results *ObjectDescriptorMockStateIDResults
	Counter uint64
}

// ObjectDescriptorMockStateIDResults contains results of the ObjectDescriptor.StateID
type ObjectDescriptorMockStateIDResults struct {
	ip1 *insolar.ID
}

// Expect sets up expected params for ObjectDescriptor.StateID
func (mmStateID *mObjectDescriptorMockStateID) Expect() *mObjectDescriptorMockStateID {
	if mmStateID.mock.funcStateID != nil {
		mmStateID.mock.t.Fatalf("ObjectDescriptorMock.StateID mock is already set by Set")
	}

	if mmStateID.defaultExpectation == nil {
		mmStateID.defaultExpectation = &ObjectDescriptorMockStateIDExpectation{}
	}

	return mmStateID
}

// Inspect accepts an inspector function that has same arguments as the ObjectDescriptor.StateID
func (mmStateID *mObjectDescriptorMockStateID) Inspect(f func()) *mObjectDescriptorMockStateID {
	if mmStateID.mock.inspectFuncStateID != nil {
		mmStateID.mock.t.Fatalf("Inspect function is already set for ObjectDescriptorMock.StateID")
	}

	mmStateID.mock.inspectFuncStateID = f

	return mmStateID
}

// Return sets up results that will be returned by ObjectDescriptor.StateID
func (mmStateID *mObjectDescriptorMockStateID) Return(ip1 *insolar.ID) *ObjectDescriptorMock {
	if mmStateID.mock.funcStateID != nil {
		mmStateID.mock.t.Fatalf("ObjectDescriptorMock.StateID mock is already set by Set")
	}

	if mmStateID.defaultExpectation == nil {
		mmStateID.defaultExpectation = &ObjectDescriptorMockStateIDExpectation{mock: mmStateID.mock}
	}
	mmStateID.defaultExpectation.results = &ObjectDescriptorMockStateIDResults{ip1}
	return mmStateID.mock
}

//Set uses given function f to mock the ObjectDescriptor.StateID method
func (mmStateID *mObjectDescriptorMockStateID) Set(f func() (ip1 *insolar.ID)) *ObjectDescriptorMock {
	if mmStateID.defaultExpectation != nil {
		mmStateID.mock.t.Fatalf("Default expectation is already set for the ObjectDescriptor.StateID method")
	}

	if len(mmStateID.expectations) > 0 {
		mmStateID.mock.t.Fatalf("Some expectations are already set for the ObjectDescriptor.StateID method")
	}

	mmStateID.mock.funcStateID = f
	return mmStateID.mock
}

// StateID implements ObjectDescriptor
func (mmStateID *ObjectDescriptorMock) StateID() (ip1 *insolar.ID) {
	mm_atomic.AddUint64(&mmStateID.beforeStateIDCounter, 1)
	defer mm_atomic.AddUint64(&mmStateID.afterStateIDCounter, 1)

	if mmStateID.inspectFuncStateID != nil {
		mmStateID.inspectFuncStateID()
	}

	if mmStateID.StateIDMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmStateID.StateIDMock.defaultExpectation.Counter, 1)

		results := mmStateID.StateIDMock.defaultExpectation.results
		if results == nil {
			mmStateID.t.Fatal("No results are set for the ObjectDescriptorMock.StateID")
		}
		return (*results).ip1
	}
	if mmStateID.funcStateID != nil {
		return mmStateID.funcStateID()
	}
	mmStateID.t.Fatalf("Unexpected call to ObjectDescriptorMock.StateID.")
	return
}

// StateIDAfterCounter returns a count of finished ObjectDescriptorMock.StateID invocations
func (mmStateID *ObjectDescriptorMock) StateIDAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmStateID.afterStateIDCounter)
}

// StateIDBeforeCounter returns a count of ObjectDescriptorMock.StateID invocations
func (mmStateID *ObjectDescriptorMock) StateIDBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmStateID.beforeStateIDCounter)
}

// MinimockStateIDDone returns true if the count of the StateID invocations corresponds
// the number of defined expectations
func (m *ObjectDescriptorMock) MinimockStateIDDone() bool {
	for _, e := range m.StateIDMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.StateIDMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterStateIDCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcStateID != nil && mm_atomic.LoadUint64(&m.afterStateIDCounter) < 1 {
		return false
	}
	return true
}

// MinimockStateIDInspect logs each unmet expectation
func (m *ObjectDescriptorMock) MinimockStateIDInspect() {
	for _, e := range m.StateIDMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to ObjectDescriptorMock.StateID")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.StateIDMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterStateIDCounter) < 1 {
		m.t.Error("Expected call to ObjectDescriptorMock.StateID")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcStateID != nil && mm_atomic.LoadUint64(&m.afterStateIDCounter) < 1 {
		m.t.Error("Expected call to ObjectDescriptorMock.StateID")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *ObjectDescriptorMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockEarliestRequestIDInspect()

		m.MinimockHeadRefInspect()

		m.MinimockMemoryInspect()

		m.MinimockParentInspect()

		m.MinimockPrototypeInspect()

		m.MinimockStateIDInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *ObjectDescriptorMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *ObjectDescriptorMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockEarliestRequestIDDone() &&
		m.MinimockHeadRefDone() &&
		m.MinimockMemoryDone() &&
		m.MinimockParentDone() &&
		m.MinimockPrototypeDone() &&
		m.MinimockStateIDDone()
}
