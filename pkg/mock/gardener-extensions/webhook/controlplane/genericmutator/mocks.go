// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gardener/gardener-extensions/pkg/webhook/controlplane/genericmutator (interfaces: Ensurer)

// Package genericmutator is a generated GoMock package.
package genericmutator

import (
	context "context"
	unit "github.com/coreos/go-systemd/unit"
	controller "github.com/gardener/gardener-extensions/pkg/controller"
	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/apps/v1"
	v10 "k8s.io/api/core/v1"
	v1beta1 "k8s.io/kubelet/config/v1beta1"
	reflect "reflect"
)

// MockEnsurer is a mock of Ensurer interface
type MockEnsurer struct {
	ctrl     *gomock.Controller
	recorder *MockEnsurerMockRecorder
}

// MockEnsurerMockRecorder is the mock recorder for MockEnsurer
type MockEnsurerMockRecorder struct {
	mock *MockEnsurer
}

// NewMockEnsurer creates a new mock instance
func NewMockEnsurer(ctrl *gomock.Controller) *MockEnsurer {
	mock := &MockEnsurer{ctrl: ctrl}
	mock.recorder = &MockEnsurerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEnsurer) EXPECT() *MockEnsurerMockRecorder {
	return m.recorder
}

// EnsureETCDStatefulSet mocks base method
func (m *MockEnsurer) EnsureETCDStatefulSet(arg0 context.Context, arg1 *v1.StatefulSet, arg2 *controller.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureETCDStatefulSet", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureETCDStatefulSet indicates an expected call of EnsureETCDStatefulSet
func (mr *MockEnsurerMockRecorder) EnsureETCDStatefulSet(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureETCDStatefulSet", reflect.TypeOf((*MockEnsurer)(nil).EnsureETCDStatefulSet), arg0, arg1, arg2)
}

// EnsureKubeAPIServerDeployment mocks base method
func (m *MockEnsurer) EnsureKubeAPIServerDeployment(arg0 context.Context, arg1 *v1.Deployment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureKubeAPIServerDeployment", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureKubeAPIServerDeployment indicates an expected call of EnsureKubeAPIServerDeployment
func (mr *MockEnsurerMockRecorder) EnsureKubeAPIServerDeployment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureKubeAPIServerDeployment", reflect.TypeOf((*MockEnsurer)(nil).EnsureKubeAPIServerDeployment), arg0, arg1)
}

// EnsureKubeAPIServerService mocks base method
func (m *MockEnsurer) EnsureKubeAPIServerService(arg0 context.Context, arg1 *v10.Service) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureKubeAPIServerService", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureKubeAPIServerService indicates an expected call of EnsureKubeAPIServerService
func (mr *MockEnsurerMockRecorder) EnsureKubeAPIServerService(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureKubeAPIServerService", reflect.TypeOf((*MockEnsurer)(nil).EnsureKubeAPIServerService), arg0, arg1)
}

// EnsureKubeControllerManagerDeployment mocks base method
func (m *MockEnsurer) EnsureKubeControllerManagerDeployment(arg0 context.Context, arg1 *v1.Deployment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureKubeControllerManagerDeployment", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureKubeControllerManagerDeployment indicates an expected call of EnsureKubeControllerManagerDeployment
func (mr *MockEnsurerMockRecorder) EnsureKubeControllerManagerDeployment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureKubeControllerManagerDeployment", reflect.TypeOf((*MockEnsurer)(nil).EnsureKubeControllerManagerDeployment), arg0, arg1)
}

// EnsureKubeSchedulerDeployment mocks base method
func (m *MockEnsurer) EnsureKubeSchedulerDeployment(arg0 context.Context, arg1 *v1.Deployment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureKubeSchedulerDeployment", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureKubeSchedulerDeployment indicates an expected call of EnsureKubeSchedulerDeployment
func (mr *MockEnsurerMockRecorder) EnsureKubeSchedulerDeployment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureKubeSchedulerDeployment", reflect.TypeOf((*MockEnsurer)(nil).EnsureKubeSchedulerDeployment), arg0, arg1)
}

// EnsureKubeletCloudProviderConfig mocks base method
func (m *MockEnsurer) EnsureKubeletCloudProviderConfig(arg0 context.Context, arg1 *string, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureKubeletCloudProviderConfig", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureKubeletCloudProviderConfig indicates an expected call of EnsureKubeletCloudProviderConfig
func (mr *MockEnsurerMockRecorder) EnsureKubeletCloudProviderConfig(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureKubeletCloudProviderConfig", reflect.TypeOf((*MockEnsurer)(nil).EnsureKubeletCloudProviderConfig), arg0, arg1, arg2)
}

// EnsureKubeletConfiguration mocks base method
func (m *MockEnsurer) EnsureKubeletConfiguration(arg0 context.Context, arg1 *v1beta1.KubeletConfiguration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureKubeletConfiguration", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureKubeletConfiguration indicates an expected call of EnsureKubeletConfiguration
func (mr *MockEnsurerMockRecorder) EnsureKubeletConfiguration(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureKubeletConfiguration", reflect.TypeOf((*MockEnsurer)(nil).EnsureKubeletConfiguration), arg0, arg1)
}

// EnsureKubeletServiceUnitOptions mocks base method
func (m *MockEnsurer) EnsureKubeletServiceUnitOptions(arg0 context.Context, arg1 []*unit.UnitOption) ([]*unit.UnitOption, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureKubeletServiceUnitOptions", arg0, arg1)
	ret0, _ := ret[0].([]*unit.UnitOption)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnsureKubeletServiceUnitOptions indicates an expected call of EnsureKubeletServiceUnitOptions
func (mr *MockEnsurerMockRecorder) EnsureKubeletServiceUnitOptions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureKubeletServiceUnitOptions", reflect.TypeOf((*MockEnsurer)(nil).EnsureKubeletServiceUnitOptions), arg0, arg1)
}

// EnsureKubernetesGeneralConfiguration mocks base method
func (m *MockEnsurer) EnsureKubernetesGeneralConfiguration(arg0 context.Context, arg1 *string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureKubernetesGeneralConfiguration", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureKubernetesGeneralConfiguration indicates an expected call of EnsureKubernetesGeneralConfiguration
func (mr *MockEnsurerMockRecorder) EnsureKubernetesGeneralConfiguration(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureKubernetesGeneralConfiguration", reflect.TypeOf((*MockEnsurer)(nil).EnsureKubernetesGeneralConfiguration), arg0, arg1)
}

// ShouldProvisionKubeletCloudProviderConfig mocks base method
func (m *MockEnsurer) ShouldProvisionKubeletCloudProviderConfig() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShouldProvisionKubeletCloudProviderConfig")
	ret0, _ := ret[0].(bool)
	return ret0
}

// ShouldProvisionKubeletCloudProviderConfig indicates an expected call of ShouldProvisionKubeletCloudProviderConfig
func (mr *MockEnsurerMockRecorder) ShouldProvisionKubeletCloudProviderConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShouldProvisionKubeletCloudProviderConfig", reflect.TypeOf((*MockEnsurer)(nil).ShouldProvisionKubeletCloudProviderConfig))
}
