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

// checks if the UpdateIpmiUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateIpmiUser{}

// UpdateIpmiUser Specifies the params for updating ipmi user.
type UpdateIpmiUser struct {
	// Specifies the node id of the node for which ipmi user info needs to be added/updated. This parameter is incompatible with 'nodeIp'.
	NodeId NullableString `json:"nodeId,omitempty"`
	// Specifies the IP Address of the node for which ipmi user needs to be added/updated. This parameter is incompatible with 'nodeId'.
	NodeIp NullableString `json:"nodeIp,omitempty"`
	// Specifies the password to be updated for the ipmi user. 
	Password NullableString `json:"password,omitempty"`
	// Specifies the ipmi username to be added/updated for given node. 
	Username NullableString `json:"username,omitempty"`
}

// NewUpdateIpmiUser instantiates a new UpdateIpmiUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateIpmiUser() *UpdateIpmiUser {
	this := UpdateIpmiUser{}
	return &this
}

// NewUpdateIpmiUserWithDefaults instantiates a new UpdateIpmiUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateIpmiUserWithDefaults() *UpdateIpmiUser {
	this := UpdateIpmiUser{}
	return &this
}

// GetNodeId returns the NodeId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateIpmiUser) GetNodeId() string {
	if o == nil || IsNil(o.NodeId.Get()) {
		var ret string
		return ret
	}
	return *o.NodeId.Get()
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateIpmiUser) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NodeId.Get(), o.NodeId.IsSet()
}

// HasNodeId returns a boolean if a field has been set.
func (o *UpdateIpmiUser) HasNodeId() bool {
	if o != nil && o.NodeId.IsSet() {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given NullableString and assigns it to the NodeId field.
func (o *UpdateIpmiUser) SetNodeId(v string) {
	o.NodeId.Set(&v)
}
// SetNodeIdNil sets the value for NodeId to be an explicit nil
func (o *UpdateIpmiUser) SetNodeIdNil() {
	o.NodeId.Set(nil)
}

// UnsetNodeId ensures that no value is present for NodeId, not even an explicit nil
func (o *UpdateIpmiUser) UnsetNodeId() {
	o.NodeId.Unset()
}

// GetNodeIp returns the NodeIp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateIpmiUser) GetNodeIp() string {
	if o == nil || IsNil(o.NodeIp.Get()) {
		var ret string
		return ret
	}
	return *o.NodeIp.Get()
}

// GetNodeIpOk returns a tuple with the NodeIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateIpmiUser) GetNodeIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NodeIp.Get(), o.NodeIp.IsSet()
}

// HasNodeIp returns a boolean if a field has been set.
func (o *UpdateIpmiUser) HasNodeIp() bool {
	if o != nil && o.NodeIp.IsSet() {
		return true
	}

	return false
}

// SetNodeIp gets a reference to the given NullableString and assigns it to the NodeIp field.
func (o *UpdateIpmiUser) SetNodeIp(v string) {
	o.NodeIp.Set(&v)
}
// SetNodeIpNil sets the value for NodeIp to be an explicit nil
func (o *UpdateIpmiUser) SetNodeIpNil() {
	o.NodeIp.Set(nil)
}

// UnsetNodeIp ensures that no value is present for NodeIp, not even an explicit nil
func (o *UpdateIpmiUser) UnsetNodeIp() {
	o.NodeIp.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateIpmiUser) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateIpmiUser) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *UpdateIpmiUser) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *UpdateIpmiUser) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *UpdateIpmiUser) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *UpdateIpmiUser) UnsetPassword() {
	o.Password.Unset()
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateIpmiUser) GetUsername() string {
	if o == nil || IsNil(o.Username.Get()) {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateIpmiUser) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *UpdateIpmiUser) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *UpdateIpmiUser) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *UpdateIpmiUser) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *UpdateIpmiUser) UnsetUsername() {
	o.Username.Unset()
}

func (o UpdateIpmiUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateIpmiUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.NodeId.IsSet() {
		toSerialize["nodeId"] = o.NodeId.Get()
	}
	if o.NodeIp.IsSet() {
		toSerialize["nodeIp"] = o.NodeIp.Get()
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	return toSerialize, nil
}

type NullableUpdateIpmiUser struct {
	value *UpdateIpmiUser
	isSet bool
}

func (v NullableUpdateIpmiUser) Get() *UpdateIpmiUser {
	return v.value
}

func (v *NullableUpdateIpmiUser) Set(val *UpdateIpmiUser) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateIpmiUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateIpmiUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateIpmiUser(val *UpdateIpmiUser) *NullableUpdateIpmiUser {
	return &NullableUpdateIpmiUser{value: val, isSet: true}
}

func (v NullableUpdateIpmiUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateIpmiUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


