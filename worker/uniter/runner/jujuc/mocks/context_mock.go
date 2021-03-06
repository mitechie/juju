// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/worker/uniter/runner/jujuc (interfaces: Context)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	params "github.com/juju/juju/apiserver/params"
	application "github.com/juju/juju/core/application"
	network "github.com/juju/juju/core/network"
	jujuc "github.com/juju/juju/worker/uniter/runner/jujuc"
	charm_v6 "gopkg.in/juju/charm.v6"
	names_v3 "gopkg.in/juju/names.v3"
	reflect "reflect"
	time "time"
)

// MockContext is a mock of Context interface
type MockContext struct {
	ctrl     *gomock.Controller
	recorder *MockContextMockRecorder
}

// MockContextMockRecorder is the mock recorder for MockContext
type MockContextMockRecorder struct {
	mock *MockContext
}

// NewMockContext creates a new mock instance
func NewMockContext(ctrl *gomock.Controller) *MockContext {
	mock := &MockContext{ctrl: ctrl}
	mock.recorder = &MockContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockContext) EXPECT() *MockContextMockRecorder {
	return m.recorder
}

// ActionParams mocks base method
func (m *MockContext) ActionParams() (map[string]interface{}, error) {
	ret := m.ctrl.Call(m, "ActionParams")
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ActionParams indicates an expected call of ActionParams
func (mr *MockContextMockRecorder) ActionParams() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActionParams", reflect.TypeOf((*MockContext)(nil).ActionParams))
}

// AddMetric mocks base method
func (m *MockContext) AddMetric(arg0, arg1 string, arg2 time.Time) error {
	ret := m.ctrl.Call(m, "AddMetric", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddMetric indicates an expected call of AddMetric
func (mr *MockContextMockRecorder) AddMetric(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMetric", reflect.TypeOf((*MockContext)(nil).AddMetric), arg0, arg1, arg2)
}

// AddMetricLabels mocks base method
func (m *MockContext) AddMetricLabels(arg0, arg1 string, arg2 time.Time, arg3 map[string]string) error {
	ret := m.ctrl.Call(m, "AddMetricLabels", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddMetricLabels indicates an expected call of AddMetricLabels
func (mr *MockContextMockRecorder) AddMetricLabels(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMetricLabels", reflect.TypeOf((*MockContext)(nil).AddMetricLabels), arg0, arg1, arg2, arg3)
}

// AddUnitStorage mocks base method
func (m *MockContext) AddUnitStorage(arg0 map[string]params.StorageConstraints) error {
	ret := m.ctrl.Call(m, "AddUnitStorage", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUnitStorage indicates an expected call of AddUnitStorage
func (mr *MockContextMockRecorder) AddUnitStorage(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUnitStorage", reflect.TypeOf((*MockContext)(nil).AddUnitStorage), arg0)
}

// ApplicationStatus mocks base method
func (m *MockContext) ApplicationStatus() (jujuc.ApplicationStatusInfo, error) {
	ret := m.ctrl.Call(m, "ApplicationStatus")
	ret0, _ := ret[0].(jujuc.ApplicationStatusInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplicationStatus indicates an expected call of ApplicationStatus
func (mr *MockContextMockRecorder) ApplicationStatus() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplicationStatus", reflect.TypeOf((*MockContext)(nil).ApplicationStatus))
}

// AvailabilityZone mocks base method
func (m *MockContext) AvailabilityZone() (string, error) {
	ret := m.ctrl.Call(m, "AvailabilityZone")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AvailabilityZone indicates an expected call of AvailabilityZone
func (mr *MockContextMockRecorder) AvailabilityZone() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AvailabilityZone", reflect.TypeOf((*MockContext)(nil).AvailabilityZone))
}

// ClosePorts mocks base method
func (m *MockContext) ClosePorts(arg0 string, arg1, arg2 int) error {
	ret := m.ctrl.Call(m, "ClosePorts", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClosePorts indicates an expected call of ClosePorts
func (mr *MockContextMockRecorder) ClosePorts(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClosePorts", reflect.TypeOf((*MockContext)(nil).ClosePorts), arg0, arg1, arg2)
}

// CloudSpec mocks base method
func (m *MockContext) CloudSpec() (*params.CloudSpec, error) {
	ret := m.ctrl.Call(m, "CloudSpec")
	ret0, _ := ret[0].(*params.CloudSpec)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloudSpec indicates an expected call of CloudSpec
func (mr *MockContextMockRecorder) CloudSpec() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudSpec", reflect.TypeOf((*MockContext)(nil).CloudSpec))
}

// Component mocks base method
func (m *MockContext) Component(arg0 string) (jujuc.ContextComponent, error) {
	ret := m.ctrl.Call(m, "Component", arg0)
	ret0, _ := ret[0].(jujuc.ContextComponent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Component indicates an expected call of Component
func (mr *MockContextMockRecorder) Component(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Component", reflect.TypeOf((*MockContext)(nil).Component), arg0)
}

// ConfigSettings mocks base method
func (m *MockContext) ConfigSettings() (charm_v6.Settings, error) {
	ret := m.ctrl.Call(m, "ConfigSettings")
	ret0, _ := ret[0].(charm_v6.Settings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConfigSettings indicates an expected call of ConfigSettings
func (mr *MockContextMockRecorder) ConfigSettings() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigSettings", reflect.TypeOf((*MockContext)(nil).ConfigSettings))
}

// DeleteCacheValue mocks base method
func (m *MockContext) DeleteCacheValue(arg0 string) error {
	ret := m.ctrl.Call(m, "DeleteCacheValue", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCacheValue indicates an expected call of DeleteCacheValue
func (mr *MockContextMockRecorder) DeleteCacheValue(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCacheValue", reflect.TypeOf((*MockContext)(nil).DeleteCacheValue), arg0)
}

// GetCache mocks base method
func (m *MockContext) GetCache() (map[string]string, error) {
	ret := m.ctrl.Call(m, "GetCache")
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCache indicates an expected call of GetCache
func (mr *MockContextMockRecorder) GetCache() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCache", reflect.TypeOf((*MockContext)(nil).GetCache))
}

// GetPodSpec mocks base method
func (m *MockContext) GetPodSpec() (string, error) {
	ret := m.ctrl.Call(m, "GetPodSpec")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPodSpec indicates an expected call of GetPodSpec
func (mr *MockContextMockRecorder) GetPodSpec() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPodSpec", reflect.TypeOf((*MockContext)(nil).GetPodSpec))
}

// GetSingleCacheValue mocks base method
func (m *MockContext) GetSingleCacheValue(arg0 string) (string, error) {
	ret := m.ctrl.Call(m, "GetSingleCacheValue", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSingleCacheValue indicates an expected call of GetSingleCacheValue
func (mr *MockContextMockRecorder) GetSingleCacheValue(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSingleCacheValue", reflect.TypeOf((*MockContext)(nil).GetSingleCacheValue), arg0)
}

// GoalState mocks base method
func (m *MockContext) GoalState() (*application.GoalState, error) {
	ret := m.ctrl.Call(m, "GoalState")
	ret0, _ := ret[0].(*application.GoalState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GoalState indicates an expected call of GoalState
func (mr *MockContextMockRecorder) GoalState() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GoalState", reflect.TypeOf((*MockContext)(nil).GoalState))
}

// HookRelation mocks base method
func (m *MockContext) HookRelation() (jujuc.ContextRelation, error) {
	ret := m.ctrl.Call(m, "HookRelation")
	ret0, _ := ret[0].(jujuc.ContextRelation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HookRelation indicates an expected call of HookRelation
func (mr *MockContextMockRecorder) HookRelation() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HookRelation", reflect.TypeOf((*MockContext)(nil).HookRelation))
}

// HookStorage mocks base method
func (m *MockContext) HookStorage() (jujuc.ContextStorageAttachment, error) {
	ret := m.ctrl.Call(m, "HookStorage")
	ret0, _ := ret[0].(jujuc.ContextStorageAttachment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HookStorage indicates an expected call of HookStorage
func (mr *MockContextMockRecorder) HookStorage() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HookStorage", reflect.TypeOf((*MockContext)(nil).HookStorage))
}

// IsLeader mocks base method
func (m *MockContext) IsLeader() (bool, error) {
	ret := m.ctrl.Call(m, "IsLeader")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsLeader indicates an expected call of IsLeader
func (mr *MockContextMockRecorder) IsLeader() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsLeader", reflect.TypeOf((*MockContext)(nil).IsLeader))
}

// LeaderSettings mocks base method
func (m *MockContext) LeaderSettings() (map[string]string, error) {
	ret := m.ctrl.Call(m, "LeaderSettings")
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LeaderSettings indicates an expected call of LeaderSettings
func (mr *MockContextMockRecorder) LeaderSettings() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LeaderSettings", reflect.TypeOf((*MockContext)(nil).LeaderSettings))
}

// LogActionMessage mocks base method
func (m *MockContext) LogActionMessage(arg0 string) error {
	ret := m.ctrl.Call(m, "LogActionMessage", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// LogActionMessage indicates an expected call of LogActionMessage
func (mr *MockContextMockRecorder) LogActionMessage(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogActionMessage", reflect.TypeOf((*MockContext)(nil).LogActionMessage), arg0)
}

// NetworkInfo mocks base method
func (m *MockContext) NetworkInfo(arg0 []string, arg1 int) (map[string]params.NetworkInfoResult, error) {
	ret := m.ctrl.Call(m, "NetworkInfo", arg0, arg1)
	ret0, _ := ret[0].(map[string]params.NetworkInfoResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NetworkInfo indicates an expected call of NetworkInfo
func (mr *MockContextMockRecorder) NetworkInfo(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetworkInfo", reflect.TypeOf((*MockContext)(nil).NetworkInfo), arg0, arg1)
}

// OpenPorts mocks base method
func (m *MockContext) OpenPorts(arg0 string, arg1, arg2 int) error {
	ret := m.ctrl.Call(m, "OpenPorts", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// OpenPorts indicates an expected call of OpenPorts
func (mr *MockContextMockRecorder) OpenPorts(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenPorts", reflect.TypeOf((*MockContext)(nil).OpenPorts), arg0, arg1, arg2)
}

// OpenedPorts mocks base method
func (m *MockContext) OpenedPorts() []network.PortRange {
	ret := m.ctrl.Call(m, "OpenedPorts")
	ret0, _ := ret[0].([]network.PortRange)
	return ret0
}

// OpenedPorts indicates an expected call of OpenedPorts
func (mr *MockContextMockRecorder) OpenedPorts() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenedPorts", reflect.TypeOf((*MockContext)(nil).OpenedPorts))
}

// PrivateAddress mocks base method
func (m *MockContext) PrivateAddress() (string, error) {
	ret := m.ctrl.Call(m, "PrivateAddress")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrivateAddress indicates an expected call of PrivateAddress
func (mr *MockContextMockRecorder) PrivateAddress() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrivateAddress", reflect.TypeOf((*MockContext)(nil).PrivateAddress))
}

// PublicAddress mocks base method
func (m *MockContext) PublicAddress() (string, error) {
	ret := m.ctrl.Call(m, "PublicAddress")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PublicAddress indicates an expected call of PublicAddress
func (mr *MockContextMockRecorder) PublicAddress() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublicAddress", reflect.TypeOf((*MockContext)(nil).PublicAddress))
}

// Relation mocks base method
func (m *MockContext) Relation(arg0 int) (jujuc.ContextRelation, error) {
	ret := m.ctrl.Call(m, "Relation", arg0)
	ret0, _ := ret[0].(jujuc.ContextRelation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Relation indicates an expected call of Relation
func (mr *MockContextMockRecorder) Relation(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Relation", reflect.TypeOf((*MockContext)(nil).Relation), arg0)
}

// RelationIds mocks base method
func (m *MockContext) RelationIds() ([]int, error) {
	ret := m.ctrl.Call(m, "RelationIds")
	ret0, _ := ret[0].([]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RelationIds indicates an expected call of RelationIds
func (mr *MockContextMockRecorder) RelationIds() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelationIds", reflect.TypeOf((*MockContext)(nil).RelationIds))
}

// RemoteApplicationName mocks base method
func (m *MockContext) RemoteApplicationName() (string, error) {
	ret := m.ctrl.Call(m, "RemoteApplicationName")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoteApplicationName indicates an expected call of RemoteApplicationName
func (mr *MockContextMockRecorder) RemoteApplicationName() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteApplicationName", reflect.TypeOf((*MockContext)(nil).RemoteApplicationName))
}

// RemoteUnitName mocks base method
func (m *MockContext) RemoteUnitName() (string, error) {
	ret := m.ctrl.Call(m, "RemoteUnitName")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoteUnitName indicates an expected call of RemoteUnitName
func (mr *MockContextMockRecorder) RemoteUnitName() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteUnitName", reflect.TypeOf((*MockContext)(nil).RemoteUnitName))
}

// RequestReboot mocks base method
func (m *MockContext) RequestReboot(arg0 jujuc.RebootPriority) error {
	ret := m.ctrl.Call(m, "RequestReboot", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RequestReboot indicates an expected call of RequestReboot
func (mr *MockContextMockRecorder) RequestReboot(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestReboot", reflect.TypeOf((*MockContext)(nil).RequestReboot), arg0)
}

// SetActionFailed mocks base method
func (m *MockContext) SetActionFailed() error {
	ret := m.ctrl.Call(m, "SetActionFailed")
	ret0, _ := ret[0].(error)
	return ret0
}

// SetActionFailed indicates an expected call of SetActionFailed
func (mr *MockContextMockRecorder) SetActionFailed() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetActionFailed", reflect.TypeOf((*MockContext)(nil).SetActionFailed))
}

// SetActionMessage mocks base method
func (m *MockContext) SetActionMessage(arg0 string) error {
	ret := m.ctrl.Call(m, "SetActionMessage", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetActionMessage indicates an expected call of SetActionMessage
func (mr *MockContextMockRecorder) SetActionMessage(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetActionMessage", reflect.TypeOf((*MockContext)(nil).SetActionMessage), arg0)
}

// SetApplicationStatus mocks base method
func (m *MockContext) SetApplicationStatus(arg0 jujuc.StatusInfo) error {
	ret := m.ctrl.Call(m, "SetApplicationStatus", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetApplicationStatus indicates an expected call of SetApplicationStatus
func (mr *MockContextMockRecorder) SetApplicationStatus(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetApplicationStatus", reflect.TypeOf((*MockContext)(nil).SetApplicationStatus), arg0)
}

// SetCacheValue mocks base method
func (m *MockContext) SetCacheValue(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "SetCacheValue", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetCacheValue indicates an expected call of SetCacheValue
func (mr *MockContextMockRecorder) SetCacheValue(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCacheValue", reflect.TypeOf((*MockContext)(nil).SetCacheValue), arg0, arg1)
}

// SetPodSpec mocks base method
func (m *MockContext) SetPodSpec(arg0 string) error {
	ret := m.ctrl.Call(m, "SetPodSpec", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetPodSpec indicates an expected call of SetPodSpec
func (mr *MockContextMockRecorder) SetPodSpec(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPodSpec", reflect.TypeOf((*MockContext)(nil).SetPodSpec), arg0)
}

// SetUnitStatus mocks base method
func (m *MockContext) SetUnitStatus(arg0 jujuc.StatusInfo) error {
	ret := m.ctrl.Call(m, "SetUnitStatus", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetUnitStatus indicates an expected call of SetUnitStatus
func (mr *MockContextMockRecorder) SetUnitStatus(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUnitStatus", reflect.TypeOf((*MockContext)(nil).SetUnitStatus), arg0)
}

// SetUnitWorkloadVersion mocks base method
func (m *MockContext) SetUnitWorkloadVersion(arg0 string) error {
	ret := m.ctrl.Call(m, "SetUnitWorkloadVersion", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetUnitWorkloadVersion indicates an expected call of SetUnitWorkloadVersion
func (mr *MockContextMockRecorder) SetUnitWorkloadVersion(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUnitWorkloadVersion", reflect.TypeOf((*MockContext)(nil).SetUnitWorkloadVersion), arg0)
}

// Storage mocks base method
func (m *MockContext) Storage(arg0 names_v3.StorageTag) (jujuc.ContextStorageAttachment, error) {
	ret := m.ctrl.Call(m, "Storage", arg0)
	ret0, _ := ret[0].(jujuc.ContextStorageAttachment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Storage indicates an expected call of Storage
func (mr *MockContextMockRecorder) Storage(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Storage", reflect.TypeOf((*MockContext)(nil).Storage), arg0)
}

// StorageTags mocks base method
func (m *MockContext) StorageTags() ([]names_v3.StorageTag, error) {
	ret := m.ctrl.Call(m, "StorageTags")
	ret0, _ := ret[0].([]names_v3.StorageTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StorageTags indicates an expected call of StorageTags
func (mr *MockContextMockRecorder) StorageTags() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageTags", reflect.TypeOf((*MockContext)(nil).StorageTags))
}

// UnitName mocks base method
func (m *MockContext) UnitName() string {
	ret := m.ctrl.Call(m, "UnitName")
	ret0, _ := ret[0].(string)
	return ret0
}

// UnitName indicates an expected call of UnitName
func (mr *MockContextMockRecorder) UnitName() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnitName", reflect.TypeOf((*MockContext)(nil).UnitName))
}

// UnitStatus mocks base method
func (m *MockContext) UnitStatus() (*jujuc.StatusInfo, error) {
	ret := m.ctrl.Call(m, "UnitStatus")
	ret0, _ := ret[0].(*jujuc.StatusInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnitStatus indicates an expected call of UnitStatus
func (mr *MockContextMockRecorder) UnitStatus() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnitStatus", reflect.TypeOf((*MockContext)(nil).UnitStatus))
}

// UnitWorkloadVersion mocks base method
func (m *MockContext) UnitWorkloadVersion() (string, error) {
	ret := m.ctrl.Call(m, "UnitWorkloadVersion")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnitWorkloadVersion indicates an expected call of UnitWorkloadVersion
func (mr *MockContextMockRecorder) UnitWorkloadVersion() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnitWorkloadVersion", reflect.TypeOf((*MockContext)(nil).UnitWorkloadVersion))
}

// UpdateActionResults mocks base method
func (m *MockContext) UpdateActionResults(arg0 []string, arg1 string) error {
	ret := m.ctrl.Call(m, "UpdateActionResults", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateActionResults indicates an expected call of UpdateActionResults
func (mr *MockContextMockRecorder) UpdateActionResults(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateActionResults", reflect.TypeOf((*MockContext)(nil).UpdateActionResults), arg0, arg1)
}

// WriteLeaderSettings mocks base method
func (m *MockContext) WriteLeaderSettings(arg0 map[string]string) error {
	ret := m.ctrl.Call(m, "WriteLeaderSettings", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteLeaderSettings indicates an expected call of WriteLeaderSettings
func (mr *MockContextMockRecorder) WriteLeaderSettings(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteLeaderSettings", reflect.TypeOf((*MockContext)(nil).WriteLeaderSettings), arg0)
}
