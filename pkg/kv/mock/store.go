// Code generated by MockGen. DO NOT EDIT.
// Source: store.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	kv "github.com/treeverse/lakefs/pkg/kv"
	params "github.com/treeverse/lakefs/pkg/kv/params"
)

// MockDriver is a mock of Driver interface.
type MockDriver struct {
	ctrl     *gomock.Controller
	recorder *MockDriverMockRecorder
}

// MockDriverMockRecorder is the mock recorder for MockDriver.
type MockDriverMockRecorder struct {
	mock *MockDriver
}

// NewMockDriver creates a new mock instance.
func NewMockDriver(ctrl *gomock.Controller) *MockDriver {
	mock := &MockDriver{ctrl: ctrl}
	mock.recorder = &MockDriverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDriver) EXPECT() *MockDriverMockRecorder {
	return m.recorder
}

// Open mocks base method.
func (m *MockDriver) Open(ctx context.Context, params params.Config) (kv.Store, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Open", ctx, params)
	ret0, _ := ret[0].(kv.Store)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Open indicates an expected call of Open.
func (mr *MockDriverMockRecorder) Open(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockDriver)(nil).Open), ctx, params)
}

// MockPredicate is a mock of Predicate interface.
type MockPredicate struct {
	ctrl     *gomock.Controller
	recorder *MockPredicateMockRecorder
}

// MockPredicateMockRecorder is the mock recorder for MockPredicate.
type MockPredicateMockRecorder struct {
	mock *MockPredicate
}

// NewMockPredicate creates a new mock instance.
func NewMockPredicate(ctrl *gomock.Controller) *MockPredicate {
	mock := &MockPredicate{ctrl: ctrl}
	mock.recorder = &MockPredicateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPredicate) EXPECT() *MockPredicateMockRecorder {
	return m.recorder
}

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockStore) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockStoreMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStore)(nil).Close))
}

// Delete mocks base method.
func (m *MockStore) Delete(ctx context.Context, partitionKey, key []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, partitionKey, key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockStoreMockRecorder) Delete(ctx, partitionKey, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockStore)(nil).Delete), ctx, partitionKey, key)
}

// Get mocks base method.
func (m *MockStore) Get(ctx context.Context, partitionKey, key []byte) (*kv.ValueWithPredicate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, partitionKey, key)
	ret0, _ := ret[0].(*kv.ValueWithPredicate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockStoreMockRecorder) Get(ctx, partitionKey, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockStore)(nil).Get), ctx, partitionKey, key)
}

// Scan mocks base method.
func (m *MockStore) Scan(ctx context.Context, partitionKey []byte, options kv.ScanOptions) (kv.EntriesIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Scan", ctx, partitionKey, options)
	ret0, _ := ret[0].(kv.EntriesIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Scan indicates an expected call of Scan.
func (mr *MockStoreMockRecorder) Scan(ctx, partitionKey, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scan", reflect.TypeOf((*MockStore)(nil).Scan), ctx, partitionKey, options)
}

// Set mocks base method.
func (m *MockStore) Set(ctx context.Context, partitionKey, key, value []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", ctx, partitionKey, key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockStoreMockRecorder) Set(ctx, partitionKey, key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockStore)(nil).Set), ctx, partitionKey, key, value)
}

// SetIf mocks base method.
func (m *MockStore) SetIf(ctx context.Context, partitionKey, key, value []byte, valuePredicate kv.Predicate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetIf", ctx, partitionKey, key, value, valuePredicate)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetIf indicates an expected call of SetIf.
func (mr *MockStoreMockRecorder) SetIf(ctx, partitionKey, key, value, valuePredicate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetIf", reflect.TypeOf((*MockStore)(nil).SetIf), ctx, partitionKey, key, value, valuePredicate)
}

// MockEntriesIterator is a mock of EntriesIterator interface.
type MockEntriesIterator struct {
	ctrl     *gomock.Controller
	recorder *MockEntriesIteratorMockRecorder
}

// MockEntriesIteratorMockRecorder is the mock recorder for MockEntriesIterator.
type MockEntriesIteratorMockRecorder struct {
	mock *MockEntriesIterator
}

// NewMockEntriesIterator creates a new mock instance.
func NewMockEntriesIterator(ctrl *gomock.Controller) *MockEntriesIterator {
	mock := &MockEntriesIterator{ctrl: ctrl}
	mock.recorder = &MockEntriesIteratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEntriesIterator) EXPECT() *MockEntriesIteratorMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockEntriesIterator) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockEntriesIteratorMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockEntriesIterator)(nil).Close))
}

// Entry mocks base method.
func (m *MockEntriesIterator) Entry() *kv.Entry {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Entry")
	ret0, _ := ret[0].(*kv.Entry)
	return ret0
}

// Entry indicates an expected call of Entry.
func (mr *MockEntriesIteratorMockRecorder) Entry() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Entry", reflect.TypeOf((*MockEntriesIterator)(nil).Entry))
}

// Err mocks base method.
func (m *MockEntriesIterator) Err() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err.
func (mr *MockEntriesIteratorMockRecorder) Err() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockEntriesIterator)(nil).Err))
}

// Next mocks base method.
func (m *MockEntriesIterator) Next() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next.
func (mr *MockEntriesIteratorMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockEntriesIterator)(nil).Next))
}