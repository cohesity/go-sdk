/*
 * Cohesity REST API
 *
 * This API list provides operations for interfacing with the Cohesity Cluster.
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cohesitysdk

import (
	"encoding/json"
)

// SetupRestoreDiskTaskInfoProto Each available extension is listed below along with the location of the proto file (relative to magneto/connectors) where it is defined.  SetupRestoreDiskTaskInfoProto extension, extension_number Location ============================================================================= vmware::SetupRestoreDiskTaskInfo vmware_setup_restore_disk_task_info, 100 connectors/vmware/vmware_setup_restore_disks.proto.proto  AgentSetupRestoreDiskTaskInfo agent_setup_restore_disk_task_info, 101 base/agent.proto  app_file::SetupRestoreTaskInfo app_file_setup_restore_task_info, 102 connectors/app_file/app_file_setup_restore.proto  hyperv::SetupRestoreDiskTaskInfo hyperv_setup_restore_disk_task_info, 103 connectors/hyperv/hyperv_setup_restore_disks.proto =============================================================================
type SetupRestoreDiskTaskInfoProto struct {
	Entity *EntityProto `json:"entity,omitempty"`
	// The path to the progress monitor root task if any.
	ProgressMonitorRootTaskPath NullableString `json:"progressMonitorRootTaskPath,omitempty"`
	RootEntity *EntityProto `json:"rootEntity,omitempty"`
	// The source view which contains the backups for the 'entity'.
	SourceViewName NullableString `json:"sourceViewName,omitempty"`
	// The id of the associated task.
	TaskId NullableInt64 `json:"taskId,omitempty"`
	// The view box id containing the backups for the 'entity'.
	ViewBoxId NullableInt64 `json:"viewBoxId,omitempty"`
	// Destination view into which the files will be cloned.
	ViewName NullableString `json:"viewName,omitempty"`
}

// NewSetupRestoreDiskTaskInfoProto instantiates a new SetupRestoreDiskTaskInfoProto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetupRestoreDiskTaskInfoProto() *SetupRestoreDiskTaskInfoProto {
	this := SetupRestoreDiskTaskInfoProto{}
	return &this
}

// NewSetupRestoreDiskTaskInfoProtoWithDefaults instantiates a new SetupRestoreDiskTaskInfoProto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetupRestoreDiskTaskInfoProtoWithDefaults() *SetupRestoreDiskTaskInfoProto {
	this := SetupRestoreDiskTaskInfoProto{}
	return &this
}

// GetEntity returns the Entity field value if set, zero value otherwise.
func (o *SetupRestoreDiskTaskInfoProto) GetEntity() EntityProto {
	if o == nil || o.Entity == nil {
		var ret EntityProto
		return ret
	}
	return *o.Entity
}

// GetEntityOk returns a tuple with the Entity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetupRestoreDiskTaskInfoProto) GetEntityOk() (*EntityProto, bool) {
	if o == nil || o.Entity == nil {
		return nil, false
	}
	return o.Entity, true
}

// HasEntity returns a boolean if a field has been set.
func (o *SetupRestoreDiskTaskInfoProto) HasEntity() bool {
	if o != nil && o.Entity != nil {
		return true
	}

	return false
}

// SetEntity gets a reference to the given EntityProto and assigns it to the Entity field.
func (o *SetupRestoreDiskTaskInfoProto) SetEntity(v EntityProto) {
	o.Entity = &v
}

// GetProgressMonitorRootTaskPath returns the ProgressMonitorRootTaskPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SetupRestoreDiskTaskInfoProto) GetProgressMonitorRootTaskPath() string {
	if o == nil || o.ProgressMonitorRootTaskPath.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProgressMonitorRootTaskPath.Get()
}

// GetProgressMonitorRootTaskPathOk returns a tuple with the ProgressMonitorRootTaskPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SetupRestoreDiskTaskInfoProto) GetProgressMonitorRootTaskPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ProgressMonitorRootTaskPath.Get(), o.ProgressMonitorRootTaskPath.IsSet()
}

// HasProgressMonitorRootTaskPath returns a boolean if a field has been set.
func (o *SetupRestoreDiskTaskInfoProto) HasProgressMonitorRootTaskPath() bool {
	if o != nil && o.ProgressMonitorRootTaskPath.IsSet() {
		return true
	}

	return false
}

// SetProgressMonitorRootTaskPath gets a reference to the given NullableString and assigns it to the ProgressMonitorRootTaskPath field.
func (o *SetupRestoreDiskTaskInfoProto) SetProgressMonitorRootTaskPath(v string) {
	o.ProgressMonitorRootTaskPath.Set(&v)
}
// SetProgressMonitorRootTaskPathNil sets the value for ProgressMonitorRootTaskPath to be an explicit nil
func (o *SetupRestoreDiskTaskInfoProto) SetProgressMonitorRootTaskPathNil() {
	o.ProgressMonitorRootTaskPath.Set(nil)
}

// UnsetProgressMonitorRootTaskPath ensures that no value is present for ProgressMonitorRootTaskPath, not even an explicit nil
func (o *SetupRestoreDiskTaskInfoProto) UnsetProgressMonitorRootTaskPath() {
	o.ProgressMonitorRootTaskPath.Unset()
}

// GetRootEntity returns the RootEntity field value if set, zero value otherwise.
func (o *SetupRestoreDiskTaskInfoProto) GetRootEntity() EntityProto {
	if o == nil || o.RootEntity == nil {
		var ret EntityProto
		return ret
	}
	return *o.RootEntity
}

// GetRootEntityOk returns a tuple with the RootEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetupRestoreDiskTaskInfoProto) GetRootEntityOk() (*EntityProto, bool) {
	if o == nil || o.RootEntity == nil {
		return nil, false
	}
	return o.RootEntity, true
}

// HasRootEntity returns a boolean if a field has been set.
func (o *SetupRestoreDiskTaskInfoProto) HasRootEntity() bool {
	if o != nil && o.RootEntity != nil {
		return true
	}

	return false
}

// SetRootEntity gets a reference to the given EntityProto and assigns it to the RootEntity field.
func (o *SetupRestoreDiskTaskInfoProto) SetRootEntity(v EntityProto) {
	o.RootEntity = &v
}

// GetSourceViewName returns the SourceViewName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SetupRestoreDiskTaskInfoProto) GetSourceViewName() string {
	if o == nil || o.SourceViewName.Get() == nil {
		var ret string
		return ret
	}
	return *o.SourceViewName.Get()
}

// GetSourceViewNameOk returns a tuple with the SourceViewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SetupRestoreDiskTaskInfoProto) GetSourceViewNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SourceViewName.Get(), o.SourceViewName.IsSet()
}

// HasSourceViewName returns a boolean if a field has been set.
func (o *SetupRestoreDiskTaskInfoProto) HasSourceViewName() bool {
	if o != nil && o.SourceViewName.IsSet() {
		return true
	}

	return false
}

// SetSourceViewName gets a reference to the given NullableString and assigns it to the SourceViewName field.
func (o *SetupRestoreDiskTaskInfoProto) SetSourceViewName(v string) {
	o.SourceViewName.Set(&v)
}
// SetSourceViewNameNil sets the value for SourceViewName to be an explicit nil
func (o *SetupRestoreDiskTaskInfoProto) SetSourceViewNameNil() {
	o.SourceViewName.Set(nil)
}

// UnsetSourceViewName ensures that no value is present for SourceViewName, not even an explicit nil
func (o *SetupRestoreDiskTaskInfoProto) UnsetSourceViewName() {
	o.SourceViewName.Unset()
}

// GetTaskId returns the TaskId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SetupRestoreDiskTaskInfoProto) GetTaskId() int64 {
	if o == nil || o.TaskId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.TaskId.Get()
}

// GetTaskIdOk returns a tuple with the TaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SetupRestoreDiskTaskInfoProto) GetTaskIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TaskId.Get(), o.TaskId.IsSet()
}

// HasTaskId returns a boolean if a field has been set.
func (o *SetupRestoreDiskTaskInfoProto) HasTaskId() bool {
	if o != nil && o.TaskId.IsSet() {
		return true
	}

	return false
}

// SetTaskId gets a reference to the given NullableInt64 and assigns it to the TaskId field.
func (o *SetupRestoreDiskTaskInfoProto) SetTaskId(v int64) {
	o.TaskId.Set(&v)
}
// SetTaskIdNil sets the value for TaskId to be an explicit nil
func (o *SetupRestoreDiskTaskInfoProto) SetTaskIdNil() {
	o.TaskId.Set(nil)
}

// UnsetTaskId ensures that no value is present for TaskId, not even an explicit nil
func (o *SetupRestoreDiskTaskInfoProto) UnsetTaskId() {
	o.TaskId.Unset()
}

// GetViewBoxId returns the ViewBoxId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SetupRestoreDiskTaskInfoProto) GetViewBoxId() int64 {
	if o == nil || o.ViewBoxId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ViewBoxId.Get()
}

// GetViewBoxIdOk returns a tuple with the ViewBoxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SetupRestoreDiskTaskInfoProto) GetViewBoxIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ViewBoxId.Get(), o.ViewBoxId.IsSet()
}

// HasViewBoxId returns a boolean if a field has been set.
func (o *SetupRestoreDiskTaskInfoProto) HasViewBoxId() bool {
	if o != nil && o.ViewBoxId.IsSet() {
		return true
	}

	return false
}

// SetViewBoxId gets a reference to the given NullableInt64 and assigns it to the ViewBoxId field.
func (o *SetupRestoreDiskTaskInfoProto) SetViewBoxId(v int64) {
	o.ViewBoxId.Set(&v)
}
// SetViewBoxIdNil sets the value for ViewBoxId to be an explicit nil
func (o *SetupRestoreDiskTaskInfoProto) SetViewBoxIdNil() {
	o.ViewBoxId.Set(nil)
}

// UnsetViewBoxId ensures that no value is present for ViewBoxId, not even an explicit nil
func (o *SetupRestoreDiskTaskInfoProto) UnsetViewBoxId() {
	o.ViewBoxId.Unset()
}

// GetViewName returns the ViewName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SetupRestoreDiskTaskInfoProto) GetViewName() string {
	if o == nil || o.ViewName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ViewName.Get()
}

// GetViewNameOk returns a tuple with the ViewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SetupRestoreDiskTaskInfoProto) GetViewNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ViewName.Get(), o.ViewName.IsSet()
}

// HasViewName returns a boolean if a field has been set.
func (o *SetupRestoreDiskTaskInfoProto) HasViewName() bool {
	if o != nil && o.ViewName.IsSet() {
		return true
	}

	return false
}

// SetViewName gets a reference to the given NullableString and assigns it to the ViewName field.
func (o *SetupRestoreDiskTaskInfoProto) SetViewName(v string) {
	o.ViewName.Set(&v)
}
// SetViewNameNil sets the value for ViewName to be an explicit nil
func (o *SetupRestoreDiskTaskInfoProto) SetViewNameNil() {
	o.ViewName.Set(nil)
}

// UnsetViewName ensures that no value is present for ViewName, not even an explicit nil
func (o *SetupRestoreDiskTaskInfoProto) UnsetViewName() {
	o.ViewName.Unset()
}

func (o SetupRestoreDiskTaskInfoProto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Entity != nil {
		toSerialize["entity"] = o.Entity
	}
	if o.ProgressMonitorRootTaskPath.IsSet() {
		toSerialize["progressMonitorRootTaskPath"] = o.ProgressMonitorRootTaskPath.Get()
	}
	if o.RootEntity != nil {
		toSerialize["rootEntity"] = o.RootEntity
	}
	if o.SourceViewName.IsSet() {
		toSerialize["sourceViewName"] = o.SourceViewName.Get()
	}
	if o.TaskId.IsSet() {
		toSerialize["taskId"] = o.TaskId.Get()
	}
	if o.ViewBoxId.IsSet() {
		toSerialize["viewBoxId"] = o.ViewBoxId.Get()
	}
	if o.ViewName.IsSet() {
		toSerialize["viewName"] = o.ViewName.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSetupRestoreDiskTaskInfoProto struct {
	value *SetupRestoreDiskTaskInfoProto
	isSet bool
}

func (v NullableSetupRestoreDiskTaskInfoProto) Get() *SetupRestoreDiskTaskInfoProto {
	return v.value
}

func (v *NullableSetupRestoreDiskTaskInfoProto) Set(val *SetupRestoreDiskTaskInfoProto) {
	v.value = val
	v.isSet = true
}

func (v NullableSetupRestoreDiskTaskInfoProto) IsSet() bool {
	return v.isSet
}

func (v *NullableSetupRestoreDiskTaskInfoProto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetupRestoreDiskTaskInfoProto(val *SetupRestoreDiskTaskInfoProto) *NullableSetupRestoreDiskTaskInfoProto {
	return &NullableSetupRestoreDiskTaskInfoProto{value: val, isSet: true}
}

func (v NullableSetupRestoreDiskTaskInfoProto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetupRestoreDiskTaskInfoProto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


