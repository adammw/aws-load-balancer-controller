// Code generated by MockGen. DO NOT EDIT.
// Source: sigs.k8s.io/aws-load-balancer-controller/pkg/k8s (interfaces: PodInfoRepo)

// Package mock_k8s is a generated GoMock package.
package mock_k8s

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	types "k8s.io/apimachinery/pkg/types"
	reflect "reflect"
	k8s "sigs.k8s.io/aws-load-balancer-controller/pkg/k8s"
)

// MockPodInfoRepo is a mock of PodInfoRepo interface
type MockPodInfoRepo struct {
	ctrl     *gomock.Controller
	recorder *MockPodInfoRepoMockRecorder
}

// MockPodInfoRepoMockRecorder is the mock recorder for MockPodInfoRepo
type MockPodInfoRepoMockRecorder struct {
	mock *MockPodInfoRepo
}

// NewMockPodInfoRepo creates a new mock instance
func NewMockPodInfoRepo(ctrl *gomock.Controller) *MockPodInfoRepo {
	mock := &MockPodInfoRepo{ctrl: ctrl}
	mock.recorder = &MockPodInfoRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPodInfoRepo) EXPECT() *MockPodInfoRepoMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockPodInfoRepo) Get(arg0 context.Context, arg1 types.NamespacedName) (k8s.PodInfo, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(k8s.PodInfo)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get
func (mr *MockPodInfoRepoMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPodInfoRepo)(nil).Get), arg0, arg1)
}

// ListKeys mocks base method
func (m *MockPodInfoRepo) ListKeys(arg0 context.Context) []types.NamespacedName {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListKeys", arg0)
	ret0, _ := ret[0].([]types.NamespacedName)
	return ret0
}

// ListKeys indicates an expected call of ListKeys
func (mr *MockPodInfoRepoMockRecorder) ListKeys(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListKeys", reflect.TypeOf((*MockPodInfoRepo)(nil).ListKeys), arg0)
}
