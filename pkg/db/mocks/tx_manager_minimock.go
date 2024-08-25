// Code generated by http://github.com/gojuno/minimock (v3.3.14). DO NOT EDIT.

package mocks

//go:generate minimock -i github.com/ukrainskykirill/chat-server/internal/client/db.TxManager -o tx_manager_minimock.go -n TxManagerMock -p mocks

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"

	mm_db "github.com/ukrainskykirill/platform_common/pkg/db"
)

// TxManagerMock implements db.TxManager
type TxManagerMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcReadCommitted          func(ctx context.Context, f mm_db.Handler) (err error)
	inspectFuncReadCommitted   func(ctx context.Context, f mm_db.Handler)
	afterReadCommittedCounter  uint64
	beforeReadCommittedCounter uint64
	ReadCommittedMock          mTxManagerMockReadCommitted
}

// NewTxManagerMock returns a mock for db.TxManager
func NewTxManagerMock(t minimock.Tester) *TxManagerMock {
	m := &TxManagerMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.ReadCommittedMock = mTxManagerMockReadCommitted{mock: m}
	m.ReadCommittedMock.callArgs = []*TxManagerMockReadCommittedParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mTxManagerMockReadCommitted struct {
	optional           bool
	mock               *TxManagerMock
	defaultExpectation *TxManagerMockReadCommittedExpectation
	expectations       []*TxManagerMockReadCommittedExpectation

	callArgs []*TxManagerMockReadCommittedParams
	mutex    sync.RWMutex

	expectedInvocations uint64
}

// TxManagerMockReadCommittedExpectation specifies expectation struct of the TxManager.ReadCommitted
type TxManagerMockReadCommittedExpectation struct {
	mock      *TxManagerMock
	params    *TxManagerMockReadCommittedParams
	paramPtrs *TxManagerMockReadCommittedParamPtrs
	results   *TxManagerMockReadCommittedResults
	Counter   uint64
}

// TxManagerMockReadCommittedParams contains parameters of the TxManager.ReadCommitted
type TxManagerMockReadCommittedParams struct {
	ctx context.Context
	f   mm_db.Handler
}

// TxManagerMockReadCommittedParamPtrs contains pointers to parameters of the TxManager.ReadCommitted
type TxManagerMockReadCommittedParamPtrs struct {
	ctx *context.Context
	f   *mm_db.Handler
}

// TxManagerMockReadCommittedResults contains results of the TxManager.ReadCommitted
type TxManagerMockReadCommittedResults struct {
	err error
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmReadCommitted *mTxManagerMockReadCommitted) Optional() *mTxManagerMockReadCommitted {
	mmReadCommitted.optional = true
	return mmReadCommitted
}

// Expect sets up expected params for TxManager.ReadCommitted
func (mmReadCommitted *mTxManagerMockReadCommitted) Expect(ctx context.Context, f mm_db.Handler) *mTxManagerMockReadCommitted {
	if mmReadCommitted.mock.funcReadCommitted != nil {
		mmReadCommitted.mock.t.Fatalf("TxManagerMock.ReadCommitted mock is already set by Set")
	}

	if mmReadCommitted.defaultExpectation == nil {
		mmReadCommitted.defaultExpectation = &TxManagerMockReadCommittedExpectation{}
	}

	if mmReadCommitted.defaultExpectation.paramPtrs != nil {
		mmReadCommitted.mock.t.Fatalf("TxManagerMock.ReadCommitted mock is already set by ExpectParams functions")
	}

	mmReadCommitted.defaultExpectation.params = &TxManagerMockReadCommittedParams{ctx, f}
	for _, e := range mmReadCommitted.expectations {
		if minimock.Equal(e.params, mmReadCommitted.defaultExpectation.params) {
			mmReadCommitted.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmReadCommitted.defaultExpectation.params)
		}
	}

	return mmReadCommitted
}

// ExpectCtxParam1 sets up expected param ctx for TxManager.ReadCommitted
func (mmReadCommitted *mTxManagerMockReadCommitted) ExpectCtxParam1(ctx context.Context) *mTxManagerMockReadCommitted {
	if mmReadCommitted.mock.funcReadCommitted != nil {
		mmReadCommitted.mock.t.Fatalf("TxManagerMock.ReadCommitted mock is already set by Set")
	}

	if mmReadCommitted.defaultExpectation == nil {
		mmReadCommitted.defaultExpectation = &TxManagerMockReadCommittedExpectation{}
	}

	if mmReadCommitted.defaultExpectation.params != nil {
		mmReadCommitted.mock.t.Fatalf("TxManagerMock.ReadCommitted mock is already set by Expect")
	}

	if mmReadCommitted.defaultExpectation.paramPtrs == nil {
		mmReadCommitted.defaultExpectation.paramPtrs = &TxManagerMockReadCommittedParamPtrs{}
	}
	mmReadCommitted.defaultExpectation.paramPtrs.ctx = &ctx

	return mmReadCommitted
}

// ExpectFParam2 sets up expected param f for TxManager.ReadCommitted
func (mmReadCommitted *mTxManagerMockReadCommitted) ExpectFParam2(f mm_db.Handler) *mTxManagerMockReadCommitted {
	if mmReadCommitted.mock.funcReadCommitted != nil {
		mmReadCommitted.mock.t.Fatalf("TxManagerMock.ReadCommitted mock is already set by Set")
	}

	if mmReadCommitted.defaultExpectation == nil {
		mmReadCommitted.defaultExpectation = &TxManagerMockReadCommittedExpectation{}
	}

	if mmReadCommitted.defaultExpectation.params != nil {
		mmReadCommitted.mock.t.Fatalf("TxManagerMock.ReadCommitted mock is already set by Expect")
	}

	if mmReadCommitted.defaultExpectation.paramPtrs == nil {
		mmReadCommitted.defaultExpectation.paramPtrs = &TxManagerMockReadCommittedParamPtrs{}
	}
	mmReadCommitted.defaultExpectation.paramPtrs.f = &f

	return mmReadCommitted
}

// Inspect accepts an inspector function that has same arguments as the TxManager.ReadCommitted
func (mmReadCommitted *mTxManagerMockReadCommitted) Inspect(f func(ctx context.Context, f mm_db.Handler)) *mTxManagerMockReadCommitted {
	if mmReadCommitted.mock.inspectFuncReadCommitted != nil {
		mmReadCommitted.mock.t.Fatalf("Inspect function is already set for TxManagerMock.ReadCommitted")
	}

	mmReadCommitted.mock.inspectFuncReadCommitted = f

	return mmReadCommitted
}

// Return sets up results that will be returned by TxManager.ReadCommitted
func (mmReadCommitted *mTxManagerMockReadCommitted) Return(err error) *TxManagerMock {
	if mmReadCommitted.mock.funcReadCommitted != nil {
		mmReadCommitted.mock.t.Fatalf("TxManagerMock.ReadCommitted mock is already set by Set")
	}

	if mmReadCommitted.defaultExpectation == nil {
		mmReadCommitted.defaultExpectation = &TxManagerMockReadCommittedExpectation{mock: mmReadCommitted.mock}
	}
	mmReadCommitted.defaultExpectation.results = &TxManagerMockReadCommittedResults{err}
	return mmReadCommitted.mock
}

// Set uses given function f to mock the TxManager.ReadCommitted method
func (mmReadCommitted *mTxManagerMockReadCommitted) Set(f func(ctx context.Context, f mm_db.Handler) (err error)) *TxManagerMock {
	if mmReadCommitted.defaultExpectation != nil {
		mmReadCommitted.mock.t.Fatalf("Default expectation is already set for the TxManager.ReadCommitted method")
	}

	if len(mmReadCommitted.expectations) > 0 {
		mmReadCommitted.mock.t.Fatalf("Some expectations are already set for the TxManager.ReadCommitted method")
	}

	mmReadCommitted.mock.funcReadCommitted = f
	return mmReadCommitted.mock
}

// When sets expectation for the TxManager.ReadCommitted which will trigger the result defined by the following
// Then helper
func (mmReadCommitted *mTxManagerMockReadCommitted) When(ctx context.Context, f mm_db.Handler) *TxManagerMockReadCommittedExpectation {
	if mmReadCommitted.mock.funcReadCommitted != nil {
		mmReadCommitted.mock.t.Fatalf("TxManagerMock.ReadCommitted mock is already set by Set")
	}

	expectation := &TxManagerMockReadCommittedExpectation{
		mock:   mmReadCommitted.mock,
		params: &TxManagerMockReadCommittedParams{ctx, f},
	}
	mmReadCommitted.expectations = append(mmReadCommitted.expectations, expectation)
	return expectation
}

// Then sets up TxManager.ReadCommitted return parameters for the expectation previously defined by the When method
func (e *TxManagerMockReadCommittedExpectation) Then(err error) *TxManagerMock {
	e.results = &TxManagerMockReadCommittedResults{err}
	return e.mock
}

// Times sets number of times TxManager.ReadCommitted should be invoked
func (mmReadCommitted *mTxManagerMockReadCommitted) Times(n uint64) *mTxManagerMockReadCommitted {
	if n == 0 {
		mmReadCommitted.mock.t.Fatalf("Times of TxManagerMock.ReadCommitted mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmReadCommitted.expectedInvocations, n)
	return mmReadCommitted
}

func (mmReadCommitted *mTxManagerMockReadCommitted) invocationsDone() bool {
	if len(mmReadCommitted.expectations) == 0 && mmReadCommitted.defaultExpectation == nil && mmReadCommitted.mock.funcReadCommitted == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmReadCommitted.mock.afterReadCommittedCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmReadCommitted.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// ReadCommitted implements db.TxManager
func (mmReadCommitted *TxManagerMock) ReadCommitted(ctx context.Context, f mm_db.Handler) (err error) {
	mm_atomic.AddUint64(&mmReadCommitted.beforeReadCommittedCounter, 1)
	defer mm_atomic.AddUint64(&mmReadCommitted.afterReadCommittedCounter, 1)

	if mmReadCommitted.inspectFuncReadCommitted != nil {
		mmReadCommitted.inspectFuncReadCommitted(ctx, f)
	}

	mm_params := TxManagerMockReadCommittedParams{ctx, f}

	// Record call args
	mmReadCommitted.ReadCommittedMock.mutex.Lock()
	mmReadCommitted.ReadCommittedMock.callArgs = append(mmReadCommitted.ReadCommittedMock.callArgs, &mm_params)
	mmReadCommitted.ReadCommittedMock.mutex.Unlock()

	for _, e := range mmReadCommitted.ReadCommittedMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmReadCommitted.ReadCommittedMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmReadCommitted.ReadCommittedMock.defaultExpectation.Counter, 1)
		mm_want := mmReadCommitted.ReadCommittedMock.defaultExpectation.params
		mm_want_ptrs := mmReadCommitted.ReadCommittedMock.defaultExpectation.paramPtrs

		mm_got := TxManagerMockReadCommittedParams{ctx, f}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmReadCommitted.t.Errorf("TxManagerMock.ReadCommitted got unexpected parameter ctx, want: %#v, got: %#v%s\n", *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.f != nil && !minimock.Equal(*mm_want_ptrs.f, mm_got.f) {
				mmReadCommitted.t.Errorf("TxManagerMock.ReadCommitted got unexpected parameter f, want: %#v, got: %#v%s\n", *mm_want_ptrs.f, mm_got.f, minimock.Diff(*mm_want_ptrs.f, mm_got.f))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmReadCommitted.t.Errorf("TxManagerMock.ReadCommitted got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmReadCommitted.ReadCommittedMock.defaultExpectation.results
		if mm_results == nil {
			mmReadCommitted.t.Fatal("No results are set for the TxManagerMock.ReadCommitted")
		}
		return (*mm_results).err
	}
	if mmReadCommitted.funcReadCommitted != nil {
		return mmReadCommitted.funcReadCommitted(ctx, f)
	}
	mmReadCommitted.t.Fatalf("Unexpected call to TxManagerMock.ReadCommitted. %v %v", ctx, f)
	return
}

// ReadCommittedAfterCounter returns a count of finished TxManagerMock.ReadCommitted invocations
func (mmReadCommitted *TxManagerMock) ReadCommittedAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmReadCommitted.afterReadCommittedCounter)
}

// ReadCommittedBeforeCounter returns a count of TxManagerMock.ReadCommitted invocations
func (mmReadCommitted *TxManagerMock) ReadCommittedBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmReadCommitted.beforeReadCommittedCounter)
}

// Calls returns a list of arguments used in each call to TxManagerMock.ReadCommitted.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmReadCommitted *mTxManagerMockReadCommitted) Calls() []*TxManagerMockReadCommittedParams {
	mmReadCommitted.mutex.RLock()

	argCopy := make([]*TxManagerMockReadCommittedParams, len(mmReadCommitted.callArgs))
	copy(argCopy, mmReadCommitted.callArgs)

	mmReadCommitted.mutex.RUnlock()

	return argCopy
}

// MinimockReadCommittedDone returns true if the count of the ReadCommitted invocations corresponds
// the number of defined expectations
func (m *TxManagerMock) MinimockReadCommittedDone() bool {
	if m.ReadCommittedMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.ReadCommittedMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.ReadCommittedMock.invocationsDone()
}

// MinimockReadCommittedInspect logs each unmet expectation
func (m *TxManagerMock) MinimockReadCommittedInspect() {
	for _, e := range m.ReadCommittedMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to TxManagerMock.ReadCommitted with params: %#v", *e.params)
		}
	}

	afterReadCommittedCounter := mm_atomic.LoadUint64(&m.afterReadCommittedCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.ReadCommittedMock.defaultExpectation != nil && afterReadCommittedCounter < 1 {
		if m.ReadCommittedMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to TxManagerMock.ReadCommitted")
		} else {
			m.t.Errorf("Expected call to TxManagerMock.ReadCommitted with params: %#v", *m.ReadCommittedMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcReadCommitted != nil && afterReadCommittedCounter < 1 {
		m.t.Error("Expected call to TxManagerMock.ReadCommitted")
	}

	if !m.ReadCommittedMock.invocationsDone() && afterReadCommittedCounter > 0 {
		m.t.Errorf("Expected %d calls to TxManagerMock.ReadCommitted but found %d calls",
			mm_atomic.LoadUint64(&m.ReadCommittedMock.expectedInvocations), afterReadCommittedCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *TxManagerMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockReadCommittedInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *TxManagerMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *TxManagerMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockReadCommittedDone()
}
