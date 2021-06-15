// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openservicemesh/osm/pkg/configurator (interfaces: Configurator)

// Package configurator is a generated GoMock package.
package configurator

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/openservicemesh/osm/pkg/apis/config/v1alpha1"
	auth "github.com/openservicemesh/osm/pkg/auth"
	v1 "k8s.io/api/core/v1"
)

// MockConfigurator is a mock of Configurator interface
type MockConfigurator struct {
	ctrl     *gomock.Controller
	recorder *MockConfiguratorMockRecorder
}

// MockConfiguratorMockRecorder is the mock recorder for MockConfigurator
type MockConfiguratorMockRecorder struct {
	mock *MockConfigurator
}

// NewMockConfigurator creates a new mock instance
func NewMockConfigurator(ctrl *gomock.Controller) *MockConfigurator {
	mock := &MockConfigurator{ctrl: ctrl}
	mock.recorder = &MockConfiguratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConfigurator) EXPECT() *MockConfiguratorMockRecorder {
	return m.recorder
}

// GetClusterId mocks base method
func (m *MockConfigurator) GetClusterId() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterId")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetClusterId indicates an expected call of GetClusterId
func (mr *MockConfiguratorMockRecorder) GetClusterId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterId", reflect.TypeOf((*MockConfigurator)(nil).GetClusterId))
}

// GetConfigResyncInterval mocks base method
func (m *MockConfigurator) GetConfigResyncInterval() time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfigResyncInterval")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// GetConfigResyncInterval indicates an expected call of GetConfigResyncInterval
func (mr *MockConfiguratorMockRecorder) GetConfigResyncInterval() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfigResyncInterval", reflect.TypeOf((*MockConfigurator)(nil).GetConfigResyncInterval))
}

// GetEnvoyImage mocks base method
func (m *MockConfigurator) GetEnvoyImage() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEnvoyImage")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetEnvoyImage indicates an expected call of GetEnvoyImage
func (mr *MockConfiguratorMockRecorder) GetEnvoyImage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnvoyImage", reflect.TypeOf((*MockConfigurator)(nil).GetEnvoyImage))
}

// GetEnvoyLogLevel mocks base method
func (m *MockConfigurator) GetEnvoyLogLevel() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEnvoyLogLevel")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetEnvoyLogLevel indicates an expected call of GetEnvoyLogLevel
func (mr *MockConfiguratorMockRecorder) GetEnvoyLogLevel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnvoyLogLevel", reflect.TypeOf((*MockConfigurator)(nil).GetEnvoyLogLevel))
}

// GetFeatureFlags mocks base method
func (m *MockConfigurator) GetFeatureFlags() v1alpha1.FeatureFlags {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFeatureFlags")
	ret0, _ := ret[0].(v1alpha1.FeatureFlags)
	return ret0
}

// GetFeatureFlags indicates an expected call of GetFeatureFlags
func (mr *MockConfiguratorMockRecorder) GetFeatureFlags() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFeatureFlags", reflect.TypeOf((*MockConfigurator)(nil).GetFeatureFlags))
}

// GetInboundExternalAuthConfig mocks base method
func (m *MockConfigurator) GetInboundExternalAuthConfig() auth.ExtAuthConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInboundExternalAuthConfig")
	ret0, _ := ret[0].(auth.ExtAuthConfig)
	return ret0
}

// GetInboundExternalAuthConfig indicates an expected call of GetInboundExternalAuthConfig
func (mr *MockConfiguratorMockRecorder) GetInboundExternalAuthConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInboundExternalAuthConfig", reflect.TypeOf((*MockConfigurator)(nil).GetInboundExternalAuthConfig))
}

// GetInboundPortExclusionList mocks base method
func (m *MockConfigurator) GetInboundPortExclusionList() []int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInboundPortExclusionList")
	ret0, _ := ret[0].([]int)
	return ret0
}

// GetInboundPortExclusionList indicates an expected call of GetInboundPortExclusionList
func (mr *MockConfiguratorMockRecorder) GetInboundPortExclusionList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInboundPortExclusionList", reflect.TypeOf((*MockConfigurator)(nil).GetInboundPortExclusionList))
}

// GetInitContainerImage mocks base method
func (m *MockConfigurator) GetInitContainerImage() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInitContainerImage")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetInitContainerImage indicates an expected call of GetInitContainerImage
func (mr *MockConfiguratorMockRecorder) GetInitContainerImage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInitContainerImage", reflect.TypeOf((*MockConfigurator)(nil).GetInitContainerImage))
}

// GetMaxDataPlaneConnections mocks base method
func (m *MockConfigurator) GetMaxDataPlaneConnections() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMaxDataPlaneConnections")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetMaxDataPlaneConnections indicates an expected call of GetMaxDataPlaneConnections
func (mr *MockConfiguratorMockRecorder) GetMaxDataPlaneConnections() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMaxDataPlaneConnections", reflect.TypeOf((*MockConfigurator)(nil).GetMaxDataPlaneConnections))
}

// GetMeshConfigJSON mocks base method
func (m *MockConfigurator) GetMeshConfigJSON() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMeshConfigJSON")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMeshConfigJSON indicates an expected call of GetMeshConfigJSON
func (mr *MockConfiguratorMockRecorder) GetMeshConfigJSON() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMeshConfigJSON", reflect.TypeOf((*MockConfigurator)(nil).GetMeshConfigJSON))
}

// GetOSMNamespace mocks base method
func (m *MockConfigurator) GetOSMNamespace() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOSMNamespace")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetOSMNamespace indicates an expected call of GetOSMNamespace
func (mr *MockConfiguratorMockRecorder) GetOSMNamespace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOSMNamespace", reflect.TypeOf((*MockConfigurator)(nil).GetOSMNamespace))
}

// GetOutboundIPRangeExclusionList mocks base method
func (m *MockConfigurator) GetOutboundIPRangeExclusionList() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOutboundIPRangeExclusionList")
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetOutboundIPRangeExclusionList indicates an expected call of GetOutboundIPRangeExclusionList
func (mr *MockConfiguratorMockRecorder) GetOutboundIPRangeExclusionList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOutboundIPRangeExclusionList", reflect.TypeOf((*MockConfigurator)(nil).GetOutboundIPRangeExclusionList))
}

// GetOutboundPortExclusionList mocks base method
func (m *MockConfigurator) GetOutboundPortExclusionList() []int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOutboundPortExclusionList")
	ret0, _ := ret[0].([]int)
	return ret0
}

// GetOutboundPortExclusionList indicates an expected call of GetOutboundPortExclusionList
func (mr *MockConfiguratorMockRecorder) GetOutboundPortExclusionList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOutboundPortExclusionList", reflect.TypeOf((*MockConfigurator)(nil).GetOutboundPortExclusionList))
}

// GetProxyResources mocks base method
func (m *MockConfigurator) GetProxyResources() v1.ResourceRequirements {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProxyResources")
	ret0, _ := ret[0].(v1.ResourceRequirements)
	return ret0
}

// GetProxyResources indicates an expected call of GetProxyResources
func (mr *MockConfiguratorMockRecorder) GetProxyResources() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProxyResources", reflect.TypeOf((*MockConfigurator)(nil).GetProxyResources))
}

// GetServiceCertValidityPeriod mocks base method
func (m *MockConfigurator) GetServiceCertValidityPeriod() time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceCertValidityPeriod")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// GetServiceCertValidityPeriod indicates an expected call of GetServiceCertValidityPeriod
func (mr *MockConfiguratorMockRecorder) GetServiceCertValidityPeriod() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceCertValidityPeriod", reflect.TypeOf((*MockConfigurator)(nil).GetServiceCertValidityPeriod))
}

// GetTracingEndpoint mocks base method
func (m *MockConfigurator) GetTracingEndpoint() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTracingEndpoint")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetTracingEndpoint indicates an expected call of GetTracingEndpoint
func (mr *MockConfiguratorMockRecorder) GetTracingEndpoint() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTracingEndpoint", reflect.TypeOf((*MockConfigurator)(nil).GetTracingEndpoint))
}

// GetTracingHost mocks base method
func (m *MockConfigurator) GetTracingHost() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTracingHost")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetTracingHost indicates an expected call of GetTracingHost
func (mr *MockConfiguratorMockRecorder) GetTracingHost() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTracingHost", reflect.TypeOf((*MockConfigurator)(nil).GetTracingHost))
}

// GetTracingPort mocks base method
func (m *MockConfigurator) GetTracingPort() uint32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTracingPort")
	ret0, _ := ret[0].(uint32)
	return ret0
}

// GetTracingPort indicates an expected call of GetTracingPort
func (mr *MockConfiguratorMockRecorder) GetTracingPort() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTracingPort", reflect.TypeOf((*MockConfigurator)(nil).GetTracingPort))
}

// IsDebugServerEnabled mocks base method
func (m *MockConfigurator) IsDebugServerEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsDebugServerEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsDebugServerEnabled indicates an expected call of IsDebugServerEnabled
func (mr *MockConfiguratorMockRecorder) IsDebugServerEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsDebugServerEnabled", reflect.TypeOf((*MockConfigurator)(nil).IsDebugServerEnabled))
}

// IsEgressEnabled mocks base method
func (m *MockConfigurator) IsEgressEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsEgressEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsEgressEnabled indicates an expected call of IsEgressEnabled
func (mr *MockConfiguratorMockRecorder) IsEgressEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsEgressEnabled", reflect.TypeOf((*MockConfigurator)(nil).IsEgressEnabled))
}

// IsPermissiveTrafficPolicyMode mocks base method
func (m *MockConfigurator) IsPermissiveTrafficPolicyMode() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsPermissiveTrafficPolicyMode")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsPermissiveTrafficPolicyMode indicates an expected call of IsPermissiveTrafficPolicyMode
func (mr *MockConfiguratorMockRecorder) IsPermissiveTrafficPolicyMode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsPermissiveTrafficPolicyMode", reflect.TypeOf((*MockConfigurator)(nil).IsPermissiveTrafficPolicyMode))
}

// IsPrivilegedInitContainer mocks base method
func (m *MockConfigurator) IsPrivilegedInitContainer() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsPrivilegedInitContainer")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsPrivilegedInitContainer indicates an expected call of IsPrivilegedInitContainer
func (mr *MockConfiguratorMockRecorder) IsPrivilegedInitContainer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsPrivilegedInitContainer", reflect.TypeOf((*MockConfigurator)(nil).IsPrivilegedInitContainer))
}

// IsTracingEnabled mocks base method
func (m *MockConfigurator) IsTracingEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsTracingEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsTracingEnabled indicates an expected call of IsTracingEnabled
func (mr *MockConfiguratorMockRecorder) IsTracingEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsTracingEnabled", reflect.TypeOf((*MockConfigurator)(nil).IsTracingEnabled))
}

// UseHTTPSIngress mocks base method
func (m *MockConfigurator) UseHTTPSIngress() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UseHTTPSIngress")
	ret0, _ := ret[0].(bool)
	return ret0
}

// UseHTTPSIngress indicates an expected call of UseHTTPSIngress
func (mr *MockConfiguratorMockRecorder) UseHTTPSIngress() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UseHTTPSIngress", reflect.TypeOf((*MockConfigurator)(nil).UseHTTPSIngress))
}
