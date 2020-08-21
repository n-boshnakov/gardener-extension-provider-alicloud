// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gardener/gardener-extension-provider-alicloud/pkg/alicloud/client (interfaces: ClientFactory,ECS,STS,SLB,VPC,Storage)

// Package client is a generated GoMock package.
package client

import (
	context "context"
	reflect "reflect"

	vpc "github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"
	client "github.com/gardener/gardener-extension-provider-alicloud/pkg/alicloud/client"
	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
	client0 "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockClientFactory is a mock of ClientFactory interface.
type MockClientFactory struct {
	ctrl     *gomock.Controller
	recorder *MockClientFactoryMockRecorder
}

// MockClientFactoryMockRecorder is the mock recorder for MockClientFactory.
type MockClientFactoryMockRecorder struct {
	mock *MockClientFactory
}

// NewMockClientFactory creates a new mock instance.
func NewMockClientFactory(ctrl *gomock.Controller) *MockClientFactory {
	mock := &MockClientFactory{ctrl: ctrl}
	mock.recorder = &MockClientFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientFactory) EXPECT() *MockClientFactoryMockRecorder {
	return m.recorder
}

// NewECSClient mocks base method.
func (m *MockClientFactory) NewECSClient(arg0, arg1, arg2 string) (client.ECS, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewECSClient", arg0, arg1, arg2)
	ret0, _ := ret[0].(client.ECS)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewECSClient indicates an expected call of NewECSClient.
func (mr *MockClientFactoryMockRecorder) NewECSClient(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewECSClient", reflect.TypeOf((*MockClientFactory)(nil).NewECSClient), arg0, arg1, arg2)
}

// NewSLBClient mocks base method.
func (m *MockClientFactory) NewSLBClient(arg0, arg1, arg2 string) (client.SLB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewSLBClient", arg0, arg1, arg2)
	ret0, _ := ret[0].(client.SLB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewSLBClient indicates an expected call of NewSLBClient.
func (mr *MockClientFactoryMockRecorder) NewSLBClient(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewSLBClient", reflect.TypeOf((*MockClientFactory)(nil).NewSLBClient), arg0, arg1, arg2)
}

// NewSTSClient mocks base method.
func (m *MockClientFactory) NewSTSClient(arg0, arg1, arg2 string) (client.STS, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewSTSClient", arg0, arg1, arg2)
	ret0, _ := ret[0].(client.STS)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewSTSClient indicates an expected call of NewSTSClient.
func (mr *MockClientFactoryMockRecorder) NewSTSClient(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewSTSClient", reflect.TypeOf((*MockClientFactory)(nil).NewSTSClient), arg0, arg1, arg2)
}

// NewStorageClientFromSecretRef mocks base method.
func (m *MockClientFactory) NewStorageClientFromSecretRef(arg0 context.Context, arg1 client0.Client, arg2 *v1.SecretReference, arg3 string) (client.Storage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewStorageClientFromSecretRef", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(client.Storage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewStorageClientFromSecretRef indicates an expected call of NewStorageClientFromSecretRef.
func (mr *MockClientFactoryMockRecorder) NewStorageClientFromSecretRef(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewStorageClientFromSecretRef", reflect.TypeOf((*MockClientFactory)(nil).NewStorageClientFromSecretRef), arg0, arg1, arg2, arg3)
}

// NewVPCClient mocks base method.
func (m *MockClientFactory) NewVPCClient(arg0, arg1, arg2 string) (client.VPC, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewVPCClient", arg0, arg1, arg2)
	ret0, _ := ret[0].(client.VPC)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewVPCClient indicates an expected call of NewVPCClient.
func (mr *MockClientFactoryMockRecorder) NewVPCClient(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewVPCClient", reflect.TypeOf((*MockClientFactory)(nil).NewVPCClient), arg0, arg1, arg2)
}

// MockECS is a mock of ECS interface.
type MockECS struct {
	ctrl     *gomock.Controller
	recorder *MockECSMockRecorder
}

// MockECSMockRecorder is the mock recorder for MockECS.
type MockECSMockRecorder struct {
	mock *MockECS
}

// NewMockECS creates a new mock instance.
func NewMockECS(ctrl *gomock.Controller) *MockECS {
	mock := &MockECS{ctrl: ctrl}
	mock.recorder = &MockECSMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockECS) EXPECT() *MockECSMockRecorder {
	return m.recorder
}

// CheckIfImageExists mocks base method.
func (m *MockECS) CheckIfImageExists(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckIfImageExists", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckIfImageExists indicates an expected call of CheckIfImageExists.
func (mr *MockECSMockRecorder) CheckIfImageExists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckIfImageExists", reflect.TypeOf((*MockECS)(nil).CheckIfImageExists), arg0, arg1)
}

// ShareImageToAccount mocks base method.
func (m *MockECS) ShareImageToAccount(arg0 context.Context, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShareImageToAccount", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// ShareImageToAccount indicates an expected call of ShareImageToAccount.
func (mr *MockECSMockRecorder) ShareImageToAccount(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShareImageToAccount", reflect.TypeOf((*MockECS)(nil).ShareImageToAccount), arg0, arg1, arg2, arg3)
}

// MockSTS is a mock of STS interface.
type MockSTS struct {
	ctrl     *gomock.Controller
	recorder *MockSTSMockRecorder
}

// MockSTSMockRecorder is the mock recorder for MockSTS.
type MockSTSMockRecorder struct {
	mock *MockSTS
}

// NewMockSTS creates a new mock instance.
func NewMockSTS(ctrl *gomock.Controller) *MockSTS {
	mock := &MockSTS{ctrl: ctrl}
	mock.recorder = &MockSTSMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSTS) EXPECT() *MockSTSMockRecorder {
	return m.recorder
}

// GetAccountIDFromCallerIdentity mocks base method.
func (m *MockSTS) GetAccountIDFromCallerIdentity(arg0 context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountIDFromCallerIdentity", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountIDFromCallerIdentity indicates an expected call of GetAccountIDFromCallerIdentity.
func (mr *MockSTSMockRecorder) GetAccountIDFromCallerIdentity(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountIDFromCallerIdentity", reflect.TypeOf((*MockSTS)(nil).GetAccountIDFromCallerIdentity), arg0)
}

// MockSLB is a mock of SLB interface.
type MockSLB struct {
	ctrl     *gomock.Controller
	recorder *MockSLBMockRecorder
}

// MockSLBMockRecorder is the mock recorder for MockSLB.
type MockSLBMockRecorder struct {
	mock *MockSLB
}

// NewMockSLB creates a new mock instance.
func NewMockSLB(ctrl *gomock.Controller) *MockSLB {
	mock := &MockSLB{ctrl: ctrl}
	mock.recorder = &MockSLBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSLB) EXPECT() *MockSLBMockRecorder {
	return m.recorder
}

// DeleteLoadBalancer mocks base method.
func (m *MockSLB) DeleteLoadBalancer(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLoadBalancer", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteLoadBalancer indicates an expected call of DeleteLoadBalancer.
func (mr *MockSLBMockRecorder) DeleteLoadBalancer(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLoadBalancer", reflect.TypeOf((*MockSLB)(nil).DeleteLoadBalancer), arg0, arg1, arg2)
}

// GetFirstVServerGroupName mocks base method.
func (m *MockSLB) GetFirstVServerGroupName(arg0 context.Context, arg1, arg2 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFirstVServerGroupName", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFirstVServerGroupName indicates an expected call of GetFirstVServerGroupName.
func (mr *MockSLBMockRecorder) GetFirstVServerGroupName(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFirstVServerGroupName", reflect.TypeOf((*MockSLB)(nil).GetFirstVServerGroupName), arg0, arg1, arg2)
}

// GetLoadBalancerIDs mocks base method.
func (m *MockSLB) GetLoadBalancerIDs(arg0 context.Context, arg1 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLoadBalancerIDs", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLoadBalancerIDs indicates an expected call of GetLoadBalancerIDs.
func (mr *MockSLBMockRecorder) GetLoadBalancerIDs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLoadBalancerIDs", reflect.TypeOf((*MockSLB)(nil).GetLoadBalancerIDs), arg0, arg1)
}

// MockVPC is a mock of VPC interface.
type MockVPC struct {
	ctrl     *gomock.Controller
	recorder *MockVPCMockRecorder
}

// MockVPCMockRecorder is the mock recorder for MockVPC.
type MockVPCMockRecorder struct {
	mock *MockVPC
}

// NewMockVPC creates a new mock instance.
func NewMockVPC(ctrl *gomock.Controller) *MockVPC {
	mock := &MockVPC{ctrl: ctrl}
	mock.recorder = &MockVPCMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVPC) EXPECT() *MockVPCMockRecorder {
	return m.recorder
}

// DescribeEipAddresses mocks base method.
func (m *MockVPC) DescribeEipAddresses(arg0 *vpc.DescribeEipAddressesRequest) (*vpc.DescribeEipAddressesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeEipAddresses", arg0)
	ret0, _ := ret[0].(*vpc.DescribeEipAddressesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeEipAddresses indicates an expected call of DescribeEipAddresses.
func (mr *MockVPCMockRecorder) DescribeEipAddresses(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeEipAddresses", reflect.TypeOf((*MockVPC)(nil).DescribeEipAddresses), arg0)
}

// DescribeNatGateways mocks base method.
func (m *MockVPC) DescribeNatGateways(arg0 *vpc.DescribeNatGatewaysRequest) (*vpc.DescribeNatGatewaysResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeNatGateways", arg0)
	ret0, _ := ret[0].(*vpc.DescribeNatGatewaysResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeNatGateways indicates an expected call of DescribeNatGateways.
func (mr *MockVPCMockRecorder) DescribeNatGateways(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeNatGateways", reflect.TypeOf((*MockVPC)(nil).DescribeNatGateways), arg0)
}

// DescribeVpcs mocks base method.
func (m *MockVPC) DescribeVpcs(arg0 *vpc.DescribeVpcsRequest) (*vpc.DescribeVpcsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeVpcs", arg0)
	ret0, _ := ret[0].(*vpc.DescribeVpcsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeVpcs indicates an expected call of DescribeVpcs.
func (mr *MockVPCMockRecorder) DescribeVpcs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeVpcs", reflect.TypeOf((*MockVPC)(nil).DescribeVpcs), arg0)
}

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// CreateBucketIfNotExists mocks base method.
func (m *MockStorage) CreateBucketIfNotExists(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBucketIfNotExists", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateBucketIfNotExists indicates an expected call of CreateBucketIfNotExists.
func (mr *MockStorageMockRecorder) CreateBucketIfNotExists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBucketIfNotExists", reflect.TypeOf((*MockStorage)(nil).CreateBucketIfNotExists), arg0, arg1)
}

// DeleteBucketIfExists mocks base method.
func (m *MockStorage) DeleteBucketIfExists(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBucketIfExists", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBucketIfExists indicates an expected call of DeleteBucketIfExists.
func (mr *MockStorageMockRecorder) DeleteBucketIfExists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBucketIfExists", reflect.TypeOf((*MockStorage)(nil).DeleteBucketIfExists), arg0, arg1)
}

// DeleteObjectsWithPrefix mocks base method.
func (m *MockStorage) DeleteObjectsWithPrefix(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteObjectsWithPrefix", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteObjectsWithPrefix indicates an expected call of DeleteObjectsWithPrefix.
func (mr *MockStorageMockRecorder) DeleteObjectsWithPrefix(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteObjectsWithPrefix", reflect.TypeOf((*MockStorage)(nil).DeleteObjectsWithPrefix), arg0, arg1, arg2)
}
