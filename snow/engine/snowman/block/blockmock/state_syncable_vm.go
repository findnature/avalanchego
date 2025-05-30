// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ava-labs/avalanchego/snow/engine/snowman/block (interfaces: StateSyncableVM)
//
// Generated by this command:
//
//	mockgen -package=blockmock -destination=blockmock/state_syncable_vm.go -mock_names=StateSyncableVM=StateSyncableVM . StateSyncableVM
//

// Package blockmock is a generated GoMock package.
package blockmock

import (
	context "context"
	reflect "reflect"

	block "github.com/ava-labs/avalanchego/snow/engine/snowman/block"
	gomock "go.uber.org/mock/gomock"
)

// StateSyncableVM is a mock of StateSyncableVM interface.
type StateSyncableVM struct {
	ctrl     *gomock.Controller
	recorder *StateSyncableVMMockRecorder
	isgomock struct{}
}

// StateSyncableVMMockRecorder is the mock recorder for StateSyncableVM.
type StateSyncableVMMockRecorder struct {
	mock *StateSyncableVM
}

// NewStateSyncableVM creates a new mock instance.
func NewStateSyncableVM(ctrl *gomock.Controller) *StateSyncableVM {
	mock := &StateSyncableVM{ctrl: ctrl}
	mock.recorder = &StateSyncableVMMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *StateSyncableVM) EXPECT() *StateSyncableVMMockRecorder {
	return m.recorder
}

// GetLastStateSummary mocks base method.
func (m *StateSyncableVM) GetLastStateSummary(arg0 context.Context) (block.StateSummary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastStateSummary", arg0)
	ret0, _ := ret[0].(block.StateSummary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLastStateSummary indicates an expected call of GetLastStateSummary.
func (mr *StateSyncableVMMockRecorder) GetLastStateSummary(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastStateSummary", reflect.TypeOf((*StateSyncableVM)(nil).GetLastStateSummary), arg0)
}

// GetOngoingSyncStateSummary mocks base method.
func (m *StateSyncableVM) GetOngoingSyncStateSummary(arg0 context.Context) (block.StateSummary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOngoingSyncStateSummary", arg0)
	ret0, _ := ret[0].(block.StateSummary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOngoingSyncStateSummary indicates an expected call of GetOngoingSyncStateSummary.
func (mr *StateSyncableVMMockRecorder) GetOngoingSyncStateSummary(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOngoingSyncStateSummary", reflect.TypeOf((*StateSyncableVM)(nil).GetOngoingSyncStateSummary), arg0)
}

// GetStateSummary mocks base method.
func (m *StateSyncableVM) GetStateSummary(ctx context.Context, summaryHeight uint64) (block.StateSummary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStateSummary", ctx, summaryHeight)
	ret0, _ := ret[0].(block.StateSummary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStateSummary indicates an expected call of GetStateSummary.
func (mr *StateSyncableVMMockRecorder) GetStateSummary(ctx, summaryHeight any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStateSummary", reflect.TypeOf((*StateSyncableVM)(nil).GetStateSummary), ctx, summaryHeight)
}

// ParseStateSummary mocks base method.
func (m *StateSyncableVM) ParseStateSummary(ctx context.Context, summaryBytes []byte) (block.StateSummary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseStateSummary", ctx, summaryBytes)
	ret0, _ := ret[0].(block.StateSummary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseStateSummary indicates an expected call of ParseStateSummary.
func (mr *StateSyncableVMMockRecorder) ParseStateSummary(ctx, summaryBytes any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseStateSummary", reflect.TypeOf((*StateSyncableVM)(nil).ParseStateSummary), ctx, summaryBytes)
}

// StateSyncEnabled mocks base method.
func (m *StateSyncableVM) StateSyncEnabled(arg0 context.Context) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateSyncEnabled", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateSyncEnabled indicates an expected call of StateSyncEnabled.
func (mr *StateSyncableVMMockRecorder) StateSyncEnabled(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateSyncEnabled", reflect.TypeOf((*StateSyncableVM)(nil).StateSyncEnabled), arg0)
}
