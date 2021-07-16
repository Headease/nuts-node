// Code generated by MockGen. DO NOT EDIT.
// Source: core/engine.go

// Package core is a generated GoMock package.
package core

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRoutable is a mock of Routable interface.
type MockRoutable struct {
	ctrl     *gomock.Controller
	recorder *MockRoutableMockRecorder
}

// MockRoutableMockRecorder is the mock recorder for MockRoutable.
type MockRoutableMockRecorder struct {
	mock *MockRoutable
}

// NewMockRoutable creates a new mock instance.
func NewMockRoutable(ctrl *gomock.Controller) *MockRoutable {
	mock := &MockRoutable{ctrl: ctrl}
	mock.recorder = &MockRoutableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoutable) EXPECT() *MockRoutableMockRecorder {
	return m.recorder
}

// Routes mocks base method.
func (m *MockRoutable) Routes(router EchoRouter) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Routes", router)
}

// Routes indicates an expected call of Routes.
func (mr *MockRoutableMockRecorder) Routes(router interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Routes", reflect.TypeOf((*MockRoutable)(nil).Routes), router)
}

// MockRunnable is a mock of Runnable interface.
type MockRunnable struct {
	ctrl     *gomock.Controller
	recorder *MockRunnableMockRecorder
}

// MockRunnableMockRecorder is the mock recorder for MockRunnable.
type MockRunnableMockRecorder struct {
	mock *MockRunnable
}

// NewMockRunnable creates a new mock instance.
func NewMockRunnable(ctrl *gomock.Controller) *MockRunnable {
	mock := &MockRunnable{ctrl: ctrl}
	mock.recorder = &MockRunnableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRunnable) EXPECT() *MockRunnableMockRecorder {
	return m.recorder
}

// Shutdown mocks base method.
func (m *MockRunnable) Shutdown() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shutdown")
	ret0, _ := ret[0].(error)
	return ret0
}

// Shutdown indicates an expected call of Shutdown.
func (mr *MockRunnableMockRecorder) Shutdown() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockRunnable)(nil).Shutdown))
}

// Start mocks base method.
func (m *MockRunnable) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockRunnableMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockRunnable)(nil).Start))
}

// MockConfigurable is a mock of Configurable interface.
type MockConfigurable struct {
	ctrl     *gomock.Controller
	recorder *MockConfigurableMockRecorder
}

// MockConfigurableMockRecorder is the mock recorder for MockConfigurable.
type MockConfigurableMockRecorder struct {
	mock *MockConfigurable
}

// NewMockConfigurable creates a new mock instance.
func NewMockConfigurable(ctrl *gomock.Controller) *MockConfigurable {
	mock := &MockConfigurable{ctrl: ctrl}
	mock.recorder = &MockConfigurableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfigurable) EXPECT() *MockConfigurableMockRecorder {
	return m.recorder
}

// Configure mocks base method.
func (m *MockConfigurable) Configure(config ServerConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Configure", config)
	ret0, _ := ret[0].(error)
	return ret0
}

// Configure indicates an expected call of Configure.
func (mr *MockConfigurableMockRecorder) Configure(config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Configure", reflect.TypeOf((*MockConfigurable)(nil).Configure), config)
}

// MockViewableDiagnostics is a mock of ViewableDiagnostics interface.
type MockViewableDiagnostics struct {
	ctrl     *gomock.Controller
	recorder *MockViewableDiagnosticsMockRecorder
}

// MockViewableDiagnosticsMockRecorder is the mock recorder for MockViewableDiagnostics.
type MockViewableDiagnosticsMockRecorder struct {
	mock *MockViewableDiagnostics
}

// NewMockViewableDiagnostics creates a new mock instance.
func NewMockViewableDiagnostics(ctrl *gomock.Controller) *MockViewableDiagnostics {
	mock := &MockViewableDiagnostics{ctrl: ctrl}
	mock.recorder = &MockViewableDiagnosticsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockViewableDiagnostics) EXPECT() *MockViewableDiagnosticsMockRecorder {
	return m.recorder
}

// Diagnostics mocks base method.
func (m *MockViewableDiagnostics) Diagnostics() []DiagnosticResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Diagnostics")
	ret0, _ := ret[0].([]DiagnosticResult)
	return ret0
}

// Diagnostics indicates an expected call of Diagnostics.
func (mr *MockViewableDiagnosticsMockRecorder) Diagnostics() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Diagnostics", reflect.TypeOf((*MockViewableDiagnostics)(nil).Diagnostics))
}

// Name mocks base method.
func (m *MockViewableDiagnostics) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockViewableDiagnosticsMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockViewableDiagnostics)(nil).Name))
}

// MockDiagnosable is a mock of Diagnosable interface.
type MockDiagnosable struct {
	ctrl     *gomock.Controller
	recorder *MockDiagnosableMockRecorder
}

// MockDiagnosableMockRecorder is the mock recorder for MockDiagnosable.
type MockDiagnosableMockRecorder struct {
	mock *MockDiagnosable
}

// NewMockDiagnosable creates a new mock instance.
func NewMockDiagnosable(ctrl *gomock.Controller) *MockDiagnosable {
	mock := &MockDiagnosable{ctrl: ctrl}
	mock.recorder = &MockDiagnosableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDiagnosable) EXPECT() *MockDiagnosableMockRecorder {
	return m.recorder
}

// Diagnostics mocks base method.
func (m *MockDiagnosable) Diagnostics() []DiagnosticResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Diagnostics")
	ret0, _ := ret[0].([]DiagnosticResult)
	return ret0
}

// Diagnostics indicates an expected call of Diagnostics.
func (mr *MockDiagnosableMockRecorder) Diagnostics() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Diagnostics", reflect.TypeOf((*MockDiagnosable)(nil).Diagnostics))
}

// MockEngine is a mock of Engine interface.
type MockEngine struct {
	ctrl     *gomock.Controller
	recorder *MockEngineMockRecorder
}

// MockEngineMockRecorder is the mock recorder for MockEngine.
type MockEngineMockRecorder struct {
	mock *MockEngine
}

// NewMockEngine creates a new mock instance.
func NewMockEngine(ctrl *gomock.Controller) *MockEngine {
	mock := &MockEngine{ctrl: ctrl}
	mock.recorder = &MockEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEngine) EXPECT() *MockEngineMockRecorder {
	return m.recorder
}

// MockNamed is a mock of Named interface.
type MockNamed struct {
	ctrl     *gomock.Controller
	recorder *MockNamedMockRecorder
}

// MockNamedMockRecorder is the mock recorder for MockNamed.
type MockNamedMockRecorder struct {
	mock *MockNamed
}

// NewMockNamed creates a new mock instance.
func NewMockNamed(ctrl *gomock.Controller) *MockNamed {
	mock := &MockNamed{ctrl: ctrl}
	mock.recorder = &MockNamedMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNamed) EXPECT() *MockNamedMockRecorder {
	return m.recorder
}

// Name mocks base method.
func (m *MockNamed) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockNamedMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockNamed)(nil).Name))
}

// MockInjectable is a mock of Injectable interface.
type MockInjectable struct {
	ctrl     *gomock.Controller
	recorder *MockInjectableMockRecorder
}

// MockInjectableMockRecorder is the mock recorder for MockInjectable.
type MockInjectableMockRecorder struct {
	mock *MockInjectable
}

// NewMockInjectable creates a new mock instance.
func NewMockInjectable(ctrl *gomock.Controller) *MockInjectable {
	mock := &MockInjectable{ctrl: ctrl}
	mock.recorder = &MockInjectableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInjectable) EXPECT() *MockInjectableMockRecorder {
	return m.recorder
}

// Config mocks base method.
func (m *MockInjectable) Config() interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Config indicates an expected call of Config.
func (mr *MockInjectableMockRecorder) Config() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockInjectable)(nil).Config))
}

// Name mocks base method.
func (m *MockInjectable) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockInjectableMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockInjectable)(nil).Name))
}
