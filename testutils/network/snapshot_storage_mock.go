package network

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/network/node"
)

// SnapshotStorageMock implements storage.SnapshotStorage
type SnapshotStorageMock struct {
	t minimock.Tester

	funcAppend          func(pulse insolar.PulseNumber, snapshot *node.Snapshot) (err error)
	inspectFuncAppend   func(pulse insolar.PulseNumber, snapshot *node.Snapshot)
	afterAppendCounter  uint64
	beforeAppendCounter uint64
	AppendMock          mSnapshotStorageMockAppend

	funcForPulseNumber          func(p1 insolar.PulseNumber) (sp1 *node.Snapshot, err error)
	inspectFuncForPulseNumber   func(p1 insolar.PulseNumber)
	afterForPulseNumberCounter  uint64
	beforeForPulseNumberCounter uint64
	ForPulseNumberMock          mSnapshotStorageMockForPulseNumber
}

// NewSnapshotStorageMock returns a mock for storage.SnapshotStorage
func NewSnapshotStorageMock(t minimock.Tester) *SnapshotStorageMock {
	m := &SnapshotStorageMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.AppendMock = mSnapshotStorageMockAppend{mock: m}
	m.AppendMock.callArgs = []*SnapshotStorageMockAppendParams{}

	m.ForPulseNumberMock = mSnapshotStorageMockForPulseNumber{mock: m}
	m.ForPulseNumberMock.callArgs = []*SnapshotStorageMockForPulseNumberParams{}

	return m
}

type mSnapshotStorageMockAppend struct {
	mock               *SnapshotStorageMock
	defaultExpectation *SnapshotStorageMockAppendExpectation
	expectations       []*SnapshotStorageMockAppendExpectation

	callArgs []*SnapshotStorageMockAppendParams
	mutex    sync.RWMutex
}

// SnapshotStorageMockAppendExpectation specifies expectation struct of the SnapshotStorage.Append
type SnapshotStorageMockAppendExpectation struct {
	mock    *SnapshotStorageMock
	params  *SnapshotStorageMockAppendParams
	results *SnapshotStorageMockAppendResults
	Counter uint64
}

// SnapshotStorageMockAppendParams contains parameters of the SnapshotStorage.Append
type SnapshotStorageMockAppendParams struct {
	pulse    insolar.PulseNumber
	snapshot *node.Snapshot
}

// SnapshotStorageMockAppendResults contains results of the SnapshotStorage.Append
type SnapshotStorageMockAppendResults struct {
	err error
}

// Expect sets up expected params for SnapshotStorage.Append
func (mmAppend *mSnapshotStorageMockAppend) Expect(pulse insolar.PulseNumber, snapshot *node.Snapshot) *mSnapshotStorageMockAppend {
	if mmAppend.mock.funcAppend != nil {
		mmAppend.mock.t.Fatalf("SnapshotStorageMock.Append mock is already set by Set")
	}

	if mmAppend.defaultExpectation == nil {
		mmAppend.defaultExpectation = &SnapshotStorageMockAppendExpectation{}
	}

	mmAppend.defaultExpectation.params = &SnapshotStorageMockAppendParams{pulse, snapshot}
	for _, e := range mmAppend.expectations {
		if minimock.Equal(e.params, mmAppend.defaultExpectation.params) {
			mmAppend.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmAppend.defaultExpectation.params)
		}
	}

	return mmAppend
}

// Inspect accepts an inspector function that has same arguments as the SnapshotStorage.Append
func (mmAppend *mSnapshotStorageMockAppend) Inspect(f func(pulse insolar.PulseNumber, snapshot *node.Snapshot)) *mSnapshotStorageMockAppend {
	if mmAppend.mock.inspectFuncAppend != nil {
		mmAppend.mock.t.Fatalf("Inspect function is already set for SnapshotStorageMock.Append")
	}

	mmAppend.mock.inspectFuncAppend = f

	return mmAppend
}

// Return sets up results that will be returned by SnapshotStorage.Append
func (mmAppend *mSnapshotStorageMockAppend) Return(err error) *SnapshotStorageMock {
	if mmAppend.mock.funcAppend != nil {
		mmAppend.mock.t.Fatalf("SnapshotStorageMock.Append mock is already set by Set")
	}

	if mmAppend.defaultExpectation == nil {
		mmAppend.defaultExpectation = &SnapshotStorageMockAppendExpectation{mock: mmAppend.mock}
	}
	mmAppend.defaultExpectation.results = &SnapshotStorageMockAppendResults{err}
	return mmAppend.mock
}

//Set uses given function f to mock the SnapshotStorage.Append method
func (mmAppend *mSnapshotStorageMockAppend) Set(f func(pulse insolar.PulseNumber, snapshot *node.Snapshot) (err error)) *SnapshotStorageMock {
	if mmAppend.defaultExpectation != nil {
		mmAppend.mock.t.Fatalf("Default expectation is already set for the SnapshotStorage.Append method")
	}

	if len(mmAppend.expectations) > 0 {
		mmAppend.mock.t.Fatalf("Some expectations are already set for the SnapshotStorage.Append method")
	}

	mmAppend.mock.funcAppend = f
	return mmAppend.mock
}

// When sets expectation for the SnapshotStorage.Append which will trigger the result defined by the following
// Then helper
func (mmAppend *mSnapshotStorageMockAppend) When(pulse insolar.PulseNumber, snapshot *node.Snapshot) *SnapshotStorageMockAppendExpectation {
	if mmAppend.mock.funcAppend != nil {
		mmAppend.mock.t.Fatalf("SnapshotStorageMock.Append mock is already set by Set")
	}

	expectation := &SnapshotStorageMockAppendExpectation{
		mock:   mmAppend.mock,
		params: &SnapshotStorageMockAppendParams{pulse, snapshot},
	}
	mmAppend.expectations = append(mmAppend.expectations, expectation)
	return expectation
}

// Then sets up SnapshotStorage.Append return parameters for the expectation previously defined by the When method
func (e *SnapshotStorageMockAppendExpectation) Then(err error) *SnapshotStorageMock {
	e.results = &SnapshotStorageMockAppendResults{err}
	return e.mock
}

// Append implements storage.SnapshotStorage
func (mmAppend *SnapshotStorageMock) Append(pulse insolar.PulseNumber, snapshot *node.Snapshot) (err error) {
	mm_atomic.AddUint64(&mmAppend.beforeAppendCounter, 1)
	defer mm_atomic.AddUint64(&mmAppend.afterAppendCounter, 1)

	if mmAppend.inspectFuncAppend != nil {
		mmAppend.inspectFuncAppend(pulse, snapshot)
	}

	mm_params := &SnapshotStorageMockAppendParams{pulse, snapshot}

	// Record call args
	mmAppend.AppendMock.mutex.Lock()
	mmAppend.AppendMock.callArgs = append(mmAppend.AppendMock.callArgs, mm_params)
	mmAppend.AppendMock.mutex.Unlock()

	for _, e := range mmAppend.AppendMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmAppend.AppendMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmAppend.AppendMock.defaultExpectation.Counter, 1)
		mm_want := mmAppend.AppendMock.defaultExpectation.params
		mm_got := SnapshotStorageMockAppendParams{pulse, snapshot}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmAppend.t.Errorf("SnapshotStorageMock.Append got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmAppend.AppendMock.defaultExpectation.results
		if mm_results == nil {
			mmAppend.t.Fatal("No results are set for the SnapshotStorageMock.Append")
		}
		return (*mm_results).err
	}
	if mmAppend.funcAppend != nil {
		return mmAppend.funcAppend(pulse, snapshot)
	}
	mmAppend.t.Fatalf("Unexpected call to SnapshotStorageMock.Append. %v %v", pulse, snapshot)
	return
}

// AppendAfterCounter returns a count of finished SnapshotStorageMock.Append invocations
func (mmAppend *SnapshotStorageMock) AppendAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAppend.afterAppendCounter)
}

// AppendBeforeCounter returns a count of SnapshotStorageMock.Append invocations
func (mmAppend *SnapshotStorageMock) AppendBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAppend.beforeAppendCounter)
}

// Calls returns a list of arguments used in each call to SnapshotStorageMock.Append.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmAppend *mSnapshotStorageMockAppend) Calls() []*SnapshotStorageMockAppendParams {
	mmAppend.mutex.RLock()

	argCopy := make([]*SnapshotStorageMockAppendParams, len(mmAppend.callArgs))
	copy(argCopy, mmAppend.callArgs)

	mmAppend.mutex.RUnlock()

	return argCopy
}

// MinimockAppendDone returns true if the count of the Append invocations corresponds
// the number of defined expectations
func (m *SnapshotStorageMock) MinimockAppendDone() bool {
	for _, e := range m.AppendMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.AppendMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterAppendCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAppend != nil && mm_atomic.LoadUint64(&m.afterAppendCounter) < 1 {
		return false
	}
	return true
}

// MinimockAppendInspect logs each unmet expectation
func (m *SnapshotStorageMock) MinimockAppendInspect() {
	for _, e := range m.AppendMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to SnapshotStorageMock.Append with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.AppendMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterAppendCounter) < 1 {
		if m.AppendMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to SnapshotStorageMock.Append")
		} else {
			m.t.Errorf("Expected call to SnapshotStorageMock.Append with params: %#v", *m.AppendMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAppend != nil && mm_atomic.LoadUint64(&m.afterAppendCounter) < 1 {
		m.t.Error("Expected call to SnapshotStorageMock.Append")
	}
}

type mSnapshotStorageMockForPulseNumber struct {
	mock               *SnapshotStorageMock
	defaultExpectation *SnapshotStorageMockForPulseNumberExpectation
	expectations       []*SnapshotStorageMockForPulseNumberExpectation

	callArgs []*SnapshotStorageMockForPulseNumberParams
	mutex    sync.RWMutex
}

// SnapshotStorageMockForPulseNumberExpectation specifies expectation struct of the SnapshotStorage.ForPulseNumber
type SnapshotStorageMockForPulseNumberExpectation struct {
	mock    *SnapshotStorageMock
	params  *SnapshotStorageMockForPulseNumberParams
	results *SnapshotStorageMockForPulseNumberResults
	Counter uint64
}

// SnapshotStorageMockForPulseNumberParams contains parameters of the SnapshotStorage.ForPulseNumber
type SnapshotStorageMockForPulseNumberParams struct {
	p1 insolar.PulseNumber
}

// SnapshotStorageMockForPulseNumberResults contains results of the SnapshotStorage.ForPulseNumber
type SnapshotStorageMockForPulseNumberResults struct {
	sp1 *node.Snapshot
	err error
}

// Expect sets up expected params for SnapshotStorage.ForPulseNumber
func (mmForPulseNumber *mSnapshotStorageMockForPulseNumber) Expect(p1 insolar.PulseNumber) *mSnapshotStorageMockForPulseNumber {
	if mmForPulseNumber.mock.funcForPulseNumber != nil {
		mmForPulseNumber.mock.t.Fatalf("SnapshotStorageMock.ForPulseNumber mock is already set by Set")
	}

	if mmForPulseNumber.defaultExpectation == nil {
		mmForPulseNumber.defaultExpectation = &SnapshotStorageMockForPulseNumberExpectation{}
	}

	mmForPulseNumber.defaultExpectation.params = &SnapshotStorageMockForPulseNumberParams{p1}
	for _, e := range mmForPulseNumber.expectations {
		if minimock.Equal(e.params, mmForPulseNumber.defaultExpectation.params) {
			mmForPulseNumber.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmForPulseNumber.defaultExpectation.params)
		}
	}

	return mmForPulseNumber
}

// Inspect accepts an inspector function that has same arguments as the SnapshotStorage.ForPulseNumber
func (mmForPulseNumber *mSnapshotStorageMockForPulseNumber) Inspect(f func(p1 insolar.PulseNumber)) *mSnapshotStorageMockForPulseNumber {
	if mmForPulseNumber.mock.inspectFuncForPulseNumber != nil {
		mmForPulseNumber.mock.t.Fatalf("Inspect function is already set for SnapshotStorageMock.ForPulseNumber")
	}

	mmForPulseNumber.mock.inspectFuncForPulseNumber = f

	return mmForPulseNumber
}

// Return sets up results that will be returned by SnapshotStorage.ForPulseNumber
func (mmForPulseNumber *mSnapshotStorageMockForPulseNumber) Return(sp1 *node.Snapshot, err error) *SnapshotStorageMock {
	if mmForPulseNumber.mock.funcForPulseNumber != nil {
		mmForPulseNumber.mock.t.Fatalf("SnapshotStorageMock.ForPulseNumber mock is already set by Set")
	}

	if mmForPulseNumber.defaultExpectation == nil {
		mmForPulseNumber.defaultExpectation = &SnapshotStorageMockForPulseNumberExpectation{mock: mmForPulseNumber.mock}
	}
	mmForPulseNumber.defaultExpectation.results = &SnapshotStorageMockForPulseNumberResults{sp1, err}
	return mmForPulseNumber.mock
}

//Set uses given function f to mock the SnapshotStorage.ForPulseNumber method
func (mmForPulseNumber *mSnapshotStorageMockForPulseNumber) Set(f func(p1 insolar.PulseNumber) (sp1 *node.Snapshot, err error)) *SnapshotStorageMock {
	if mmForPulseNumber.defaultExpectation != nil {
		mmForPulseNumber.mock.t.Fatalf("Default expectation is already set for the SnapshotStorage.ForPulseNumber method")
	}

	if len(mmForPulseNumber.expectations) > 0 {
		mmForPulseNumber.mock.t.Fatalf("Some expectations are already set for the SnapshotStorage.ForPulseNumber method")
	}

	mmForPulseNumber.mock.funcForPulseNumber = f
	return mmForPulseNumber.mock
}

// When sets expectation for the SnapshotStorage.ForPulseNumber which will trigger the result defined by the following
// Then helper
func (mmForPulseNumber *mSnapshotStorageMockForPulseNumber) When(p1 insolar.PulseNumber) *SnapshotStorageMockForPulseNumberExpectation {
	if mmForPulseNumber.mock.funcForPulseNumber != nil {
		mmForPulseNumber.mock.t.Fatalf("SnapshotStorageMock.ForPulseNumber mock is already set by Set")
	}

	expectation := &SnapshotStorageMockForPulseNumberExpectation{
		mock:   mmForPulseNumber.mock,
		params: &SnapshotStorageMockForPulseNumberParams{p1},
	}
	mmForPulseNumber.expectations = append(mmForPulseNumber.expectations, expectation)
	return expectation
}

// Then sets up SnapshotStorage.ForPulseNumber return parameters for the expectation previously defined by the When method
func (e *SnapshotStorageMockForPulseNumberExpectation) Then(sp1 *node.Snapshot, err error) *SnapshotStorageMock {
	e.results = &SnapshotStorageMockForPulseNumberResults{sp1, err}
	return e.mock
}

// ForPulseNumber implements storage.SnapshotStorage
func (mmForPulseNumber *SnapshotStorageMock) ForPulseNumber(p1 insolar.PulseNumber) (sp1 *node.Snapshot, err error) {
	mm_atomic.AddUint64(&mmForPulseNumber.beforeForPulseNumberCounter, 1)
	defer mm_atomic.AddUint64(&mmForPulseNumber.afterForPulseNumberCounter, 1)

	if mmForPulseNumber.inspectFuncForPulseNumber != nil {
		mmForPulseNumber.inspectFuncForPulseNumber(p1)
	}

	mm_params := &SnapshotStorageMockForPulseNumberParams{p1}

	// Record call args
	mmForPulseNumber.ForPulseNumberMock.mutex.Lock()
	mmForPulseNumber.ForPulseNumberMock.callArgs = append(mmForPulseNumber.ForPulseNumberMock.callArgs, mm_params)
	mmForPulseNumber.ForPulseNumberMock.mutex.Unlock()

	for _, e := range mmForPulseNumber.ForPulseNumberMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.sp1, e.results.err
		}
	}

	if mmForPulseNumber.ForPulseNumberMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmForPulseNumber.ForPulseNumberMock.defaultExpectation.Counter, 1)
		mm_want := mmForPulseNumber.ForPulseNumberMock.defaultExpectation.params
		mm_got := SnapshotStorageMockForPulseNumberParams{p1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmForPulseNumber.t.Errorf("SnapshotStorageMock.ForPulseNumber got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmForPulseNumber.ForPulseNumberMock.defaultExpectation.results
		if mm_results == nil {
			mmForPulseNumber.t.Fatal("No results are set for the SnapshotStorageMock.ForPulseNumber")
		}
		return (*mm_results).sp1, (*mm_results).err
	}
	if mmForPulseNumber.funcForPulseNumber != nil {
		return mmForPulseNumber.funcForPulseNumber(p1)
	}
	mmForPulseNumber.t.Fatalf("Unexpected call to SnapshotStorageMock.ForPulseNumber. %v", p1)
	return
}

// ForPulseNumberAfterCounter returns a count of finished SnapshotStorageMock.ForPulseNumber invocations
func (mmForPulseNumber *SnapshotStorageMock) ForPulseNumberAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmForPulseNumber.afterForPulseNumberCounter)
}

// ForPulseNumberBeforeCounter returns a count of SnapshotStorageMock.ForPulseNumber invocations
func (mmForPulseNumber *SnapshotStorageMock) ForPulseNumberBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmForPulseNumber.beforeForPulseNumberCounter)
}

// Calls returns a list of arguments used in each call to SnapshotStorageMock.ForPulseNumber.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmForPulseNumber *mSnapshotStorageMockForPulseNumber) Calls() []*SnapshotStorageMockForPulseNumberParams {
	mmForPulseNumber.mutex.RLock()

	argCopy := make([]*SnapshotStorageMockForPulseNumberParams, len(mmForPulseNumber.callArgs))
	copy(argCopy, mmForPulseNumber.callArgs)

	mmForPulseNumber.mutex.RUnlock()

	return argCopy
}

// MinimockForPulseNumberDone returns true if the count of the ForPulseNumber invocations corresponds
// the number of defined expectations
func (m *SnapshotStorageMock) MinimockForPulseNumberDone() bool {
	for _, e := range m.ForPulseNumberMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ForPulseNumberMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterForPulseNumberCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcForPulseNumber != nil && mm_atomic.LoadUint64(&m.afterForPulseNumberCounter) < 1 {
		return false
	}
	return true
}

// MinimockForPulseNumberInspect logs each unmet expectation
func (m *SnapshotStorageMock) MinimockForPulseNumberInspect() {
	for _, e := range m.ForPulseNumberMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to SnapshotStorageMock.ForPulseNumber with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ForPulseNumberMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterForPulseNumberCounter) < 1 {
		if m.ForPulseNumberMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to SnapshotStorageMock.ForPulseNumber")
		} else {
			m.t.Errorf("Expected call to SnapshotStorageMock.ForPulseNumber with params: %#v", *m.ForPulseNumberMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcForPulseNumber != nil && mm_atomic.LoadUint64(&m.afterForPulseNumberCounter) < 1 {
		m.t.Error("Expected call to SnapshotStorageMock.ForPulseNumber")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *SnapshotStorageMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockAppendInspect()

		m.MinimockForPulseNumberInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *SnapshotStorageMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *SnapshotStorageMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockAppendDone() &&
		m.MinimockForPulseNumberDone()
}
