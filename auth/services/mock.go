// Code generated by MockGen. DO NOT EDIT.
// Source: auth/services/services.go

// Package services is a generated GoMock package.
package services

import (
	http "net/http"
	url "net/url"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	did "github.com/nuts-foundation/go-did/did"
	contract "github.com/nuts-foundation/nuts-node/auth/contract"
)

// MockOAuthClient is a mock of OAuthClient interface.
type MockOAuthClient struct {
	ctrl     *gomock.Controller
	recorder *MockOAuthClientMockRecorder
}

// MockOAuthClientMockRecorder is the mock recorder for MockOAuthClient.
type MockOAuthClientMockRecorder struct {
	mock *MockOAuthClient
}

// NewMockOAuthClient creates a new mock instance.
func NewMockOAuthClient(ctrl *gomock.Controller) *MockOAuthClient {
	mock := &MockOAuthClient{ctrl: ctrl}
	mock.recorder = &MockOAuthClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOAuthClient) EXPECT() *MockOAuthClientMockRecorder {
	return m.recorder
}

// Configure mocks base method.
func (m *MockOAuthClient) Configure() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Configure")
	ret0, _ := ret[0].(error)
	return ret0
}

// Configure indicates an expected call of Configure.
func (mr *MockOAuthClientMockRecorder) Configure() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Configure", reflect.TypeOf((*MockOAuthClient)(nil).Configure))
}

// CreateAccessToken mocks base method.
func (m *MockOAuthClient) CreateAccessToken(request CreateAccessTokenRequest) (*AccessTokenResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccessToken", request)
	ret0, _ := ret[0].(*AccessTokenResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccessToken indicates an expected call of CreateAccessToken.
func (mr *MockOAuthClientMockRecorder) CreateAccessToken(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccessToken", reflect.TypeOf((*MockOAuthClient)(nil).CreateAccessToken), request)
}

// CreateJwtGrant mocks base method.
func (m *MockOAuthClient) CreateJwtGrant(request CreateJwtGrantRequest) (*JwtBearerTokenResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateJwtGrant", request)
	ret0, _ := ret[0].(*JwtBearerTokenResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateJwtGrant indicates an expected call of CreateJwtGrant.
func (mr *MockOAuthClientMockRecorder) CreateJwtGrant(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateJwtGrant", reflect.TypeOf((*MockOAuthClient)(nil).CreateJwtGrant), request)
}

// GetOAuthEndpointURL mocks base method.
func (m *MockOAuthClient) GetOAuthEndpointURL(service string, authorizer did.DID) (url.URL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOAuthEndpointURL", service, authorizer)
	ret0, _ := ret[0].(url.URL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOAuthEndpointURL indicates an expected call of GetOAuthEndpointURL.
func (mr *MockOAuthClientMockRecorder) GetOAuthEndpointURL(service, authorizer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOAuthEndpointURL", reflect.TypeOf((*MockOAuthClient)(nil).GetOAuthEndpointURL), service, authorizer)
}

// IntrospectAccessToken mocks base method.
func (m *MockOAuthClient) IntrospectAccessToken(token string) (*NutsAccessToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IntrospectAccessToken", token)
	ret0, _ := ret[0].(*NutsAccessToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IntrospectAccessToken indicates an expected call of IntrospectAccessToken.
func (mr *MockOAuthClientMockRecorder) IntrospectAccessToken(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IntrospectAccessToken", reflect.TypeOf((*MockOAuthClient)(nil).IntrospectAccessToken), token)
}

// MockSignedToken is a mock of SignedToken interface.
type MockSignedToken struct {
	ctrl     *gomock.Controller
	recorder *MockSignedTokenMockRecorder
}

// MockSignedTokenMockRecorder is the mock recorder for MockSignedToken.
type MockSignedTokenMockRecorder struct {
	mock *MockSignedToken
}

// NewMockSignedToken creates a new mock instance.
func NewMockSignedToken(ctrl *gomock.Controller) *MockSignedToken {
	mock := &MockSignedToken{ctrl: ctrl}
	mock.recorder = &MockSignedTokenMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSignedToken) EXPECT() *MockSignedTokenMockRecorder {
	return m.recorder
}

// Contract mocks base method.
func (m *MockSignedToken) Contract() contract.Contract {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Contract")
	ret0, _ := ret[0].(contract.Contract)
	return ret0
}

// Contract indicates an expected call of Contract.
func (mr *MockSignedTokenMockRecorder) Contract() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Contract", reflect.TypeOf((*MockSignedToken)(nil).Contract))
}

// SignerAttributes mocks base method.
func (m *MockSignedToken) SignerAttributes() (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignerAttributes")
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignerAttributes indicates an expected call of SignerAttributes.
func (mr *MockSignedTokenMockRecorder) SignerAttributes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignerAttributes", reflect.TypeOf((*MockSignedToken)(nil).SignerAttributes))
}

// MockVPProofValueParser is a mock of VPProofValueParser interface.
type MockVPProofValueParser struct {
	ctrl     *gomock.Controller
	recorder *MockVPProofValueParserMockRecorder
}

// MockVPProofValueParserMockRecorder is the mock recorder for MockVPProofValueParser.
type MockVPProofValueParserMockRecorder struct {
	mock *MockVPProofValueParser
}

// NewMockVPProofValueParser creates a new mock instance.
func NewMockVPProofValueParser(ctrl *gomock.Controller) *MockVPProofValueParser {
	mock := &MockVPProofValueParser{ctrl: ctrl}
	mock.recorder = &MockVPProofValueParserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVPProofValueParser) EXPECT() *MockVPProofValueParserMockRecorder {
	return m.recorder
}

// Parse mocks base method.
func (m *MockVPProofValueParser) Parse(rawAuthToken string) (SignedToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parse", rawAuthToken)
	ret0, _ := ret[0].(SignedToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Parse indicates an expected call of Parse.
func (mr *MockVPProofValueParserMockRecorder) Parse(rawAuthToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parse", reflect.TypeOf((*MockVPProofValueParser)(nil).Parse), rawAuthToken)
}

// Verify mocks base method.
func (m *MockVPProofValueParser) Verify(token SignedToken) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", token)
	ret0, _ := ret[0].(error)
	return ret0
}

// Verify indicates an expected call of Verify.
func (mr *MockVPProofValueParserMockRecorder) Verify(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockVPProofValueParser)(nil).Verify), token)
}

// MockContractNotary is a mock of ContractNotary interface.
type MockContractNotary struct {
	ctrl     *gomock.Controller
	recorder *MockContractNotaryMockRecorder
}

// MockContractNotaryMockRecorder is the mock recorder for MockContractNotary.
type MockContractNotaryMockRecorder struct {
	mock *MockContractNotary
}

// NewMockContractNotary creates a new mock instance.
func NewMockContractNotary(ctrl *gomock.Controller) *MockContractNotary {
	mock := &MockContractNotary{ctrl: ctrl}
	mock.recorder = &MockContractNotaryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContractNotary) EXPECT() *MockContractNotaryMockRecorder {
	return m.recorder
}

// DrawUpContract mocks base method.
func (m *MockContractNotary) DrawUpContract(template contract.Template, orgID did.DID, validFrom time.Time, validDuration time.Duration) (*contract.Contract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DrawUpContract", template, orgID, validFrom, validDuration)
	ret0, _ := ret[0].(*contract.Contract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DrawUpContract indicates an expected call of DrawUpContract.
func (mr *MockContractNotaryMockRecorder) DrawUpContract(template, orgID, validFrom, validDuration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DrawUpContract", reflect.TypeOf((*MockContractNotary)(nil).DrawUpContract), template, orgID, validFrom, validDuration)
}

// MockContractClient is a mock of ContractClient interface.
type MockContractClient struct {
	ctrl     *gomock.Controller
	recorder *MockContractClientMockRecorder
}

// MockContractClientMockRecorder is the mock recorder for MockContractClient.
type MockContractClientMockRecorder struct {
	mock *MockContractClient
}

// NewMockContractClient creates a new mock instance.
func NewMockContractClient(ctrl *gomock.Controller) *MockContractClient {
	mock := &MockContractClient{ctrl: ctrl}
	mock.recorder = &MockContractClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContractClient) EXPECT() *MockContractClientMockRecorder {
	return m.recorder
}

// Configure mocks base method.
func (m *MockContractClient) Configure() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Configure")
	ret0, _ := ret[0].(error)
	return ret0
}

// Configure indicates an expected call of Configure.
func (mr *MockContractClientMockRecorder) Configure() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Configure", reflect.TypeOf((*MockContractClient)(nil).Configure))
}

// CreateSigningSession mocks base method.
func (m *MockContractClient) CreateSigningSession(sessionRequest CreateSessionRequest) (contract.SessionPointer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSigningSession", sessionRequest)
	ret0, _ := ret[0].(contract.SessionPointer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSigningSession indicates an expected call of CreateSigningSession.
func (mr *MockContractClientMockRecorder) CreateSigningSession(sessionRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSigningSession", reflect.TypeOf((*MockContractClient)(nil).CreateSigningSession), sessionRequest)
}

// HandlerFunc mocks base method.
func (m *MockContractClient) HandlerFunc() http.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandlerFunc")
	ret0, _ := ret[0].(http.HandlerFunc)
	return ret0
}

// HandlerFunc indicates an expected call of HandlerFunc.
func (mr *MockContractClientMockRecorder) HandlerFunc() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandlerFunc", reflect.TypeOf((*MockContractClient)(nil).HandlerFunc))
}

// SigningSessionStatus mocks base method.
func (m *MockContractClient) SigningSessionStatus(sessionID string) (contract.SigningSessionResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SigningSessionStatus", sessionID)
	ret0, _ := ret[0].(contract.SigningSessionResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SigningSessionStatus indicates an expected call of SigningSessionStatus.
func (mr *MockContractClientMockRecorder) SigningSessionStatus(sessionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SigningSessionStatus", reflect.TypeOf((*MockContractClient)(nil).SigningSessionStatus), sessionID)
}

// VerifyVP mocks base method.
func (m *MockContractClient) VerifyVP(rawVerifiablePresentation []byte, checkTime *time.Time) (*contract.VPVerificationResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyVP", rawVerifiablePresentation, checkTime)
	ret0, _ := ret[0].(*contract.VPVerificationResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyVP indicates an expected call of VerifyVP.
func (mr *MockContractClientMockRecorder) VerifyVP(rawVerifiablePresentation, checkTime interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyVP", reflect.TypeOf((*MockContractClient)(nil).VerifyVP), rawVerifiablePresentation, checkTime)
}

// MockCompoundServiceClient is a mock of CompoundServiceClient interface.
type MockCompoundServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockCompoundServiceClientMockRecorder
}

// MockCompoundServiceClientMockRecorder is the mock recorder for MockCompoundServiceClient.
type MockCompoundServiceClientMockRecorder struct {
	mock *MockCompoundServiceClient
}

// NewMockCompoundServiceClient creates a new mock instance.
func NewMockCompoundServiceClient(ctrl *gomock.Controller) *MockCompoundServiceClient {
	mock := &MockCompoundServiceClient{ctrl: ctrl}
	mock.recorder = &MockCompoundServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCompoundServiceClient) EXPECT() *MockCompoundServiceClientMockRecorder {
	return m.recorder
}

// GetCompoundService mocks base method.
func (m *MockCompoundServiceClient) GetCompoundService(id did.DID, serviceType string) (*did.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCompoundService", id, serviceType)
	ret0, _ := ret[0].(*did.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCompoundService indicates an expected call of GetCompoundService.
func (mr *MockCompoundServiceClientMockRecorder) GetCompoundService(id, serviceType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCompoundService", reflect.TypeOf((*MockCompoundServiceClient)(nil).GetCompoundService), id, serviceType)
}
