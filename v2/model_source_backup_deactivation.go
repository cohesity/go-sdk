/*
Cohesity REST API

Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
)

// checks if the SourceBackupDeactivation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourceBackupDeactivation{}

// SourceBackupDeactivation Specifies the request parmeters to deactivate the backup of failover entities on source cluster.
type SourceBackupDeactivation struct {
	// If this is set to true then objects will not be removed from protection group. If this is set to false, then all objects which are being failed over will be removed from the protection group. If protection group left with zero entities then it will be paused automatically.
	KeepFailoverObjects NullableBool `json:"keepFailoverObjects,omitempty"`
	// Specifies the list of all local entity ids of all the objects being failed from the source cluster. Backup will be deactiaved for all given objects.
	Objects []FailoverObject `json:"objects,omitempty"`
	// Specifies the protection group id of the source cluster from where the objects being failed over. If this is not specified then it will be infer from the list of objects being failed over.
	ProtectionGroupId NullableString `json:"protectionGroupId,omitempty"`
	// Specifies the replication cluster Id involved in failover operation.
	ReplicationClusterId NullableInt64 `json:"replicationClusterId,omitempty"`
	// If failover is initiated by view based orchastrator, then this field specifies the local view id of source cluster which is being failed over. Backup will be deactivated for view object.
	ViewId NullableString `json:"viewId,omitempty"`
}

// NewSourceBackupDeactivation instantiates a new SourceBackupDeactivation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceBackupDeactivation() *SourceBackupDeactivation {
	this := SourceBackupDeactivation{}
	return &this
}

// NewSourceBackupDeactivationWithDefaults instantiates a new SourceBackupDeactivation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceBackupDeactivationWithDefaults() *SourceBackupDeactivation {
	this := SourceBackupDeactivation{}
	return &this
}

// GetKeepFailoverObjects returns the KeepFailoverObjects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourceBackupDeactivation) GetKeepFailoverObjects() bool {
	if o == nil || IsNil(o.KeepFailoverObjects.Get()) {
		var ret bool
		return ret
	}
	return *o.KeepFailoverObjects.Get()
}

// GetKeepFailoverObjectsOk returns a tuple with the KeepFailoverObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SourceBackupDeactivation) GetKeepFailoverObjectsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.KeepFailoverObjects.Get(), o.KeepFailoverObjects.IsSet()
}

// HasKeepFailoverObjects returns a boolean if a field has been set.
func (o *SourceBackupDeactivation) HasKeepFailoverObjects() bool {
	if o != nil && o.KeepFailoverObjects.IsSet() {
		return true
	}

	return false
}

// SetKeepFailoverObjects gets a reference to the given NullableBool and assigns it to the KeepFailoverObjects field.
func (o *SourceBackupDeactivation) SetKeepFailoverObjects(v bool) {
	o.KeepFailoverObjects.Set(&v)
}
// SetKeepFailoverObjectsNil sets the value for KeepFailoverObjects to be an explicit nil
func (o *SourceBackupDeactivation) SetKeepFailoverObjectsNil() {
	o.KeepFailoverObjects.Set(nil)
}

// UnsetKeepFailoverObjects ensures that no value is present for KeepFailoverObjects, not even an explicit nil
func (o *SourceBackupDeactivation) UnsetKeepFailoverObjects() {
	o.KeepFailoverObjects.Unset()
}

// GetObjects returns the Objects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourceBackupDeactivation) GetObjects() []FailoverObject {
	if o == nil {
		var ret []FailoverObject
		return ret
	}
	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SourceBackupDeactivation) GetObjectsOk() ([]FailoverObject, bool) {
	if o == nil || IsNil(o.Objects) {
		return nil, false
	}
	return o.Objects, true
}

// HasObjects returns a boolean if a field has been set.
func (o *SourceBackupDeactivation) HasObjects() bool {
	if o != nil && !IsNil(o.Objects) {
		return true
	}

	return false
}

// SetObjects gets a reference to the given []FailoverObject and assigns it to the Objects field.
func (o *SourceBackupDeactivation) SetObjects(v []FailoverObject) {
	o.Objects = v
}

// GetProtectionGroupId returns the ProtectionGroupId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourceBackupDeactivation) GetProtectionGroupId() string {
	if o == nil || IsNil(o.ProtectionGroupId.Get()) {
		var ret string
		return ret
	}
	return *o.ProtectionGroupId.Get()
}

// GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SourceBackupDeactivation) GetProtectionGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProtectionGroupId.Get(), o.ProtectionGroupId.IsSet()
}

// HasProtectionGroupId returns a boolean if a field has been set.
func (o *SourceBackupDeactivation) HasProtectionGroupId() bool {
	if o != nil && o.ProtectionGroupId.IsSet() {
		return true
	}

	return false
}

// SetProtectionGroupId gets a reference to the given NullableString and assigns it to the ProtectionGroupId field.
func (o *SourceBackupDeactivation) SetProtectionGroupId(v string) {
	o.ProtectionGroupId.Set(&v)
}
// SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil
func (o *SourceBackupDeactivation) SetProtectionGroupIdNil() {
	o.ProtectionGroupId.Set(nil)
}

// UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil
func (o *SourceBackupDeactivation) UnsetProtectionGroupId() {
	o.ProtectionGroupId.Unset()
}

// GetReplicationClusterId returns the ReplicationClusterId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourceBackupDeactivation) GetReplicationClusterId() int64 {
	if o == nil || IsNil(o.ReplicationClusterId.Get()) {
		var ret int64
		return ret
	}
	return *o.ReplicationClusterId.Get()
}

// GetReplicationClusterIdOk returns a tuple with the ReplicationClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SourceBackupDeactivation) GetReplicationClusterIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReplicationClusterId.Get(), o.ReplicationClusterId.IsSet()
}

// HasReplicationClusterId returns a boolean if a field has been set.
func (o *SourceBackupDeactivation) HasReplicationClusterId() bool {
	if o != nil && o.ReplicationClusterId.IsSet() {
		return true
	}

	return false
}

// SetReplicationClusterId gets a reference to the given NullableInt64 and assigns it to the ReplicationClusterId field.
func (o *SourceBackupDeactivation) SetReplicationClusterId(v int64) {
	o.ReplicationClusterId.Set(&v)
}
// SetReplicationClusterIdNil sets the value for ReplicationClusterId to be an explicit nil
func (o *SourceBackupDeactivation) SetReplicationClusterIdNil() {
	o.ReplicationClusterId.Set(nil)
}

// UnsetReplicationClusterId ensures that no value is present for ReplicationClusterId, not even an explicit nil
func (o *SourceBackupDeactivation) UnsetReplicationClusterId() {
	o.ReplicationClusterId.Unset()
}

// GetViewId returns the ViewId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourceBackupDeactivation) GetViewId() string {
	if o == nil || IsNil(o.ViewId.Get()) {
		var ret string
		return ret
	}
	return *o.ViewId.Get()
}

// GetViewIdOk returns a tuple with the ViewId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SourceBackupDeactivation) GetViewIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ViewId.Get(), o.ViewId.IsSet()
}

// HasViewId returns a boolean if a field has been set.
func (o *SourceBackupDeactivation) HasViewId() bool {
	if o != nil && o.ViewId.IsSet() {
		return true
	}

	return false
}

// SetViewId gets a reference to the given NullableString and assigns it to the ViewId field.
func (o *SourceBackupDeactivation) SetViewId(v string) {
	o.ViewId.Set(&v)
}
// SetViewIdNil sets the value for ViewId to be an explicit nil
func (o *SourceBackupDeactivation) SetViewIdNil() {
	o.ViewId.Set(nil)
}

// UnsetViewId ensures that no value is present for ViewId, not even an explicit nil
func (o *SourceBackupDeactivation) UnsetViewId() {
	o.ViewId.Unset()
}

func (o SourceBackupDeactivation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceBackupDeactivation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.KeepFailoverObjects.IsSet() {
		toSerialize["keepFailoverObjects"] = o.KeepFailoverObjects.Get()
	}
	if o.Objects != nil {
		toSerialize["objects"] = o.Objects
	}
	if o.ProtectionGroupId.IsSet() {
		toSerialize["protectionGroupId"] = o.ProtectionGroupId.Get()
	}
	if o.ReplicationClusterId.IsSet() {
		toSerialize["replicationClusterId"] = o.ReplicationClusterId.Get()
	}
	if o.ViewId.IsSet() {
		toSerialize["viewId"] = o.ViewId.Get()
	}
	return toSerialize, nil
}

type NullableSourceBackupDeactivation struct {
	value *SourceBackupDeactivation
	isSet bool
}

func (v NullableSourceBackupDeactivation) Get() *SourceBackupDeactivation {
	return v.value
}

func (v *NullableSourceBackupDeactivation) Set(val *SourceBackupDeactivation) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceBackupDeactivation) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceBackupDeactivation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceBackupDeactivation(val *SourceBackupDeactivation) *NullableSourceBackupDeactivation {
	return &NullableSourceBackupDeactivation{value: val, isSet: true}
}

func (v NullableSourceBackupDeactivation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceBackupDeactivation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


