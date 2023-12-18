// Code generated by MockGen. DO NOT EDIT.
// Source: vcr/verifier/interface.go
//
// Generated by this command:
//
//	mockgen -destination=vcr/verifier/mock.go -package=verifier -source=vcr/verifier/interface.go
//
// Package verifier is a generated GoMock package.
package verifier

import (
	reflect "reflect"
	time "time"

	ssi "github.com/nuts-foundation/go-did"
	vc "github.com/nuts-foundation/go-did/vc"
	core "github.com/nuts-foundation/nuts-node/core"
	credential "github.com/nuts-foundation/nuts-node/vcr/credential"
	gomock "go.uber.org/mock/gomock"
)

// MockVerifier is a mock of Verifier interface.
type MockVerifier struct {
	ctrl     *gomock.Controller
	recorder *MockVerifierMockRecorder
}

// MockVerifierMockRecorder is the mock recorder for MockVerifier.
type MockVerifierMockRecorder struct {
	mock *MockVerifier
}

// NewMockVerifier creates a new mock instance.
func NewMockVerifier(ctrl *gomock.Controller) *MockVerifier {
	mock := &MockVerifier{ctrl: ctrl}
	mock.recorder = &MockVerifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVerifier) EXPECT() *MockVerifierMockRecorder {
	return m.recorder
}

// GetRevocation mocks base method.
func (m *MockVerifier) GetRevocation(id ssi.URI) (*credential.Revocation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRevocation", id)
	ret0, _ := ret[0].(*credential.Revocation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRevocation indicates an expected call of GetRevocation.
func (mr *MockVerifierMockRecorder) GetRevocation(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRevocation", reflect.TypeOf((*MockVerifier)(nil).GetRevocation), id)
}

// IsRevoked mocks base method.
func (m *MockVerifier) IsRevoked(credentialID ssi.URI) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsRevoked", credentialID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsRevoked indicates an expected call of IsRevoked.
func (mr *MockVerifierMockRecorder) IsRevoked(credentialID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsRevoked", reflect.TypeOf((*MockVerifier)(nil).IsRevoked), credentialID)
}

// RegisterRevocation mocks base method.
func (m *MockVerifier) RegisterRevocation(revocation credential.Revocation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterRevocation", revocation)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterRevocation indicates an expected call of RegisterRevocation.
func (mr *MockVerifierMockRecorder) RegisterRevocation(revocation any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterRevocation", reflect.TypeOf((*MockVerifier)(nil).RegisterRevocation), revocation)
}

// Verify mocks base method.
func (m *MockVerifier) Verify(credential vc.VerifiableCredential, allowUntrusted, checkSignature bool, validAt *time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", credential, allowUntrusted, checkSignature, validAt)
	ret0, _ := ret[0].(error)
	return ret0
}

// Verify indicates an expected call of Verify.
func (mr *MockVerifierMockRecorder) Verify(credential, allowUntrusted, checkSignature, validAt any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockVerifier)(nil).Verify), credential, allowUntrusted, checkSignature, validAt)
}

// VerifySignature mocks base method.
func (m *MockVerifier) VerifySignature(credentialToVerify vc.VerifiableCredential, at *time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifySignature", credentialToVerify, at)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifySignature indicates an expected call of VerifySignature.
func (mr *MockVerifierMockRecorder) VerifySignature(credentialToVerify, at any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifySignature", reflect.TypeOf((*MockVerifier)(nil).VerifySignature), credentialToVerify, at)
}

// VerifyVP mocks base method.
func (m *MockVerifier) VerifyVP(presentation vc.VerifiablePresentation, verifyVCs, allowUntrustedVCs bool, validAt *time.Time) ([]vc.VerifiableCredential, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyVP", presentation, verifyVCs, allowUntrustedVCs, validAt)
	ret0, _ := ret[0].([]vc.VerifiableCredential)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyVP indicates an expected call of VerifyVP.
func (mr *MockVerifierMockRecorder) VerifyVP(presentation, verifyVCs, allowUntrustedVCs, validAt any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyVP", reflect.TypeOf((*MockVerifier)(nil).VerifyVP), presentation, verifyVCs, allowUntrustedVCs, validAt)
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
func (m *MockStore) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockStoreMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStore)(nil).Close))
}

// Diagnostics mocks base method.
func (m *MockStore) Diagnostics() []core.DiagnosticResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Diagnostics")
	ret0, _ := ret[0].([]core.DiagnosticResult)
	return ret0
}

// Diagnostics indicates an expected call of Diagnostics.
func (mr *MockStoreMockRecorder) Diagnostics() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Diagnostics", reflect.TypeOf((*MockStore)(nil).Diagnostics))
}

// GetRevocations mocks base method.
func (m *MockStore) GetRevocations(id ssi.URI) ([]*credential.Revocation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRevocations", id)
	ret0, _ := ret[0].([]*credential.Revocation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRevocations indicates an expected call of GetRevocations.
func (mr *MockStoreMockRecorder) GetRevocations(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRevocations", reflect.TypeOf((*MockStore)(nil).GetRevocations), id)
}

// StoreRevocation mocks base method.
func (m *MockStore) StoreRevocation(r credential.Revocation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreRevocation", r)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreRevocation indicates an expected call of StoreRevocation.
func (mr *MockStoreMockRecorder) StoreRevocation(r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreRevocation", reflect.TypeOf((*MockStore)(nil).StoreRevocation), r)
}
