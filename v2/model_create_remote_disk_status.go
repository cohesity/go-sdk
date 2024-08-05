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

// checks if the CreateRemoteDiskStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRemoteDiskStatus{}

// CreateRemoteDiskStatus Specifies the status of creating remote disk.
type CreateRemoteDiskStatus struct {
	// Specifies the error message when creating remote disk fails.
	Message NullableString `json:"message,omitempty"`
	// Specifies the NFS mount path of the remote disk.
	MountPath NullableString `json:"mountPath,omitempty"`
	// Specifies the node id of the disk. If not specified, the disk will be evenly distributed across all the nodes.
	NodeId NullableInt64 `json:"nodeId,omitempty"`
	// Specifies the creating status of this disk.
	Status NullableString `json:"status,omitempty"`
	// Specifies the tier of the disk
	Tier NullableString `json:"tier,omitempty"`
}

// NewCreateRemoteDiskStatus instantiates a new CreateRemoteDiskStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRemoteDiskStatus() *CreateRemoteDiskStatus {
	this := CreateRemoteDiskStatus{}
	return &this
}

// NewCreateRemoteDiskStatusWithDefaults instantiates a new CreateRemoteDiskStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRemoteDiskStatusWithDefaults() *CreateRemoteDiskStatus {
	this := CreateRemoteDiskStatus{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateRemoteDiskStatus) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateRemoteDiskStatus) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *CreateRemoteDiskStatus) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *CreateRemoteDiskStatus) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *CreateRemoteDiskStatus) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *CreateRemoteDiskStatus) UnsetMessage() {
	o.Message.Unset()
}

// GetMountPath returns the MountPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateRemoteDiskStatus) GetMountPath() string {
	if o == nil || IsNil(o.MountPath.Get()) {
		var ret string
		return ret
	}
	return *o.MountPath.Get()
}

// GetMountPathOk returns a tuple with the MountPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateRemoteDiskStatus) GetMountPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MountPath.Get(), o.MountPath.IsSet()
}

// HasMountPath returns a boolean if a field has been set.
func (o *CreateRemoteDiskStatus) HasMountPath() bool {
	if o != nil && o.MountPath.IsSet() {
		return true
	}

	return false
}

// SetMountPath gets a reference to the given NullableString and assigns it to the MountPath field.
func (o *CreateRemoteDiskStatus) SetMountPath(v string) {
	o.MountPath.Set(&v)
}
// SetMountPathNil sets the value for MountPath to be an explicit nil
func (o *CreateRemoteDiskStatus) SetMountPathNil() {
	o.MountPath.Set(nil)
}

// UnsetMountPath ensures that no value is present for MountPath, not even an explicit nil
func (o *CreateRemoteDiskStatus) UnsetMountPath() {
	o.MountPath.Unset()
}

// GetNodeId returns the NodeId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateRemoteDiskStatus) GetNodeId() int64 {
	if o == nil || IsNil(o.NodeId.Get()) {
		var ret int64
		return ret
	}
	return *o.NodeId.Get()
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateRemoteDiskStatus) GetNodeIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NodeId.Get(), o.NodeId.IsSet()
}

// HasNodeId returns a boolean if a field has been set.
func (o *CreateRemoteDiskStatus) HasNodeId() bool {
	if o != nil && o.NodeId.IsSet() {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given NullableInt64 and assigns it to the NodeId field.
func (o *CreateRemoteDiskStatus) SetNodeId(v int64) {
	o.NodeId.Set(&v)
}
// SetNodeIdNil sets the value for NodeId to be an explicit nil
func (o *CreateRemoteDiskStatus) SetNodeIdNil() {
	o.NodeId.Set(nil)
}

// UnsetNodeId ensures that no value is present for NodeId, not even an explicit nil
func (o *CreateRemoteDiskStatus) UnsetNodeId() {
	o.NodeId.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateRemoteDiskStatus) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateRemoteDiskStatus) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *CreateRemoteDiskStatus) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *CreateRemoteDiskStatus) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *CreateRemoteDiskStatus) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *CreateRemoteDiskStatus) UnsetStatus() {
	o.Status.Unset()
}

// GetTier returns the Tier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateRemoteDiskStatus) GetTier() string {
	if o == nil || IsNil(o.Tier.Get()) {
		var ret string
		return ret
	}
	return *o.Tier.Get()
}

// GetTierOk returns a tuple with the Tier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateRemoteDiskStatus) GetTierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tier.Get(), o.Tier.IsSet()
}

// HasTier returns a boolean if a field has been set.
func (o *CreateRemoteDiskStatus) HasTier() bool {
	if o != nil && o.Tier.IsSet() {
		return true
	}

	return false
}

// SetTier gets a reference to the given NullableString and assigns it to the Tier field.
func (o *CreateRemoteDiskStatus) SetTier(v string) {
	o.Tier.Set(&v)
}
// SetTierNil sets the value for Tier to be an explicit nil
func (o *CreateRemoteDiskStatus) SetTierNil() {
	o.Tier.Set(nil)
}

// UnsetTier ensures that no value is present for Tier, not even an explicit nil
func (o *CreateRemoteDiskStatus) UnsetTier() {
	o.Tier.Unset()
}

func (o CreateRemoteDiskStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRemoteDiskStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	if o.MountPath.IsSet() {
		toSerialize["mountPath"] = o.MountPath.Get()
	}
	if o.NodeId.IsSet() {
		toSerialize["nodeId"] = o.NodeId.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.Tier.IsSet() {
		toSerialize["tier"] = o.Tier.Get()
	}
	return toSerialize, nil
}

type NullableCreateRemoteDiskStatus struct {
	value *CreateRemoteDiskStatus
	isSet bool
}

func (v NullableCreateRemoteDiskStatus) Get() *CreateRemoteDiskStatus {
	return v.value
}

func (v *NullableCreateRemoteDiskStatus) Set(val *CreateRemoteDiskStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRemoteDiskStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRemoteDiskStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRemoteDiskStatus(val *CreateRemoteDiskStatus) *NullableCreateRemoteDiskStatus {
	return &NullableCreateRemoteDiskStatus{value: val, isSet: true}
}

func (v NullableCreateRemoteDiskStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRemoteDiskStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


