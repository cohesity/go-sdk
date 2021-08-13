/*
 * Cohesity REST API
 *
 * Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

// package helios
package models

import (
	"encoding/json"
	. "github.com/cohesity/go-sdk/helios/utils"
)

var _ = NullableBool{}

// RedoLogGroupConfig Specifies Redo log group configuration
type RedoLogGroupConfig struct {
	// Specifies no. of redo log groups.
	NumGroups NullableInt32 `json:"numGroups,omitempty"`
	// Specifies Log member name prefix.
	MemberPrefix NullableString `json:"memberPrefix,omitempty"`
	// Specifies Size of the member in MB.
	SizeMBytes NullableInt32 `json:"sizeMBytes,omitempty"`
	// Specifies list of members of this redo log group.
	GroupMembers []string `json:"groupMembers,omitempty"`
}

// NewRedoLogGroupConfig instantiates a new RedoLogGroupConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRedoLogGroupConfig() *RedoLogGroupConfig {
	this := RedoLogGroupConfig{}
	return &this
}

// NewRedoLogGroupConfigWithDefaults instantiates a new RedoLogGroupConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRedoLogGroupConfigWithDefaults() *RedoLogGroupConfig {
	this := RedoLogGroupConfig{}
	return &this
}

// GetNumGroups returns the NumGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RedoLogGroupConfig) GetNumGroups() int32 {
	if o == nil || o.NumGroups.Get() == nil {
		var ret int32
		return ret
	}
	return *o.NumGroups.Get()
}

// GetNumGroupsOk returns a tuple with the NumGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RedoLogGroupConfig) GetNumGroupsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NumGroups.Get(), o.NumGroups.IsSet()
}

// HasNumGroups returns a boolean if a field has been set.
func (o *RedoLogGroupConfig) HasNumGroups() bool {
	if o != nil && o.NumGroups.IsSet() {
		return true
	}

	return false
}

// SetNumGroups gets a reference to the given NullableInt32 and assigns it to the NumGroups field.
func (o *RedoLogGroupConfig) SetNumGroups(v int32) {
	o.NumGroups.Set(&v)
}
// SetNumGroupsNil sets the value for NumGroups to be an explicit nil
func (o *RedoLogGroupConfig) SetNumGroupsNil() {
	o.NumGroups.Set(nil)
}

// UnsetNumGroups ensures that no value is present for NumGroups, not even an explicit nil
func (o *RedoLogGroupConfig) UnsetNumGroups() {
	o.NumGroups.Unset()
}

// GetMemberPrefix returns the MemberPrefix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RedoLogGroupConfig) GetMemberPrefix() string {
	if o == nil || o.MemberPrefix.Get() == nil {
		var ret string
		return ret
	}
	return *o.MemberPrefix.Get()
}

// GetMemberPrefixOk returns a tuple with the MemberPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RedoLogGroupConfig) GetMemberPrefixOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MemberPrefix.Get(), o.MemberPrefix.IsSet()
}

// HasMemberPrefix returns a boolean if a field has been set.
func (o *RedoLogGroupConfig) HasMemberPrefix() bool {
	if o != nil && o.MemberPrefix.IsSet() {
		return true
	}

	return false
}

// SetMemberPrefix gets a reference to the given NullableString and assigns it to the MemberPrefix field.
func (o *RedoLogGroupConfig) SetMemberPrefix(v string) {
	o.MemberPrefix.Set(&v)
}
// SetMemberPrefixNil sets the value for MemberPrefix to be an explicit nil
func (o *RedoLogGroupConfig) SetMemberPrefixNil() {
	o.MemberPrefix.Set(nil)
}

// UnsetMemberPrefix ensures that no value is present for MemberPrefix, not even an explicit nil
func (o *RedoLogGroupConfig) UnsetMemberPrefix() {
	o.MemberPrefix.Unset()
}

// GetSizeMBytes returns the SizeMBytes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RedoLogGroupConfig) GetSizeMBytes() int32 {
	if o == nil || o.SizeMBytes.Get() == nil {
		var ret int32
		return ret
	}
	return *o.SizeMBytes.Get()
}

// GetSizeMBytesOk returns a tuple with the SizeMBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RedoLogGroupConfig) GetSizeMBytesOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SizeMBytes.Get(), o.SizeMBytes.IsSet()
}

// HasSizeMBytes returns a boolean if a field has been set.
func (o *RedoLogGroupConfig) HasSizeMBytes() bool {
	if o != nil && o.SizeMBytes.IsSet() {
		return true
	}

	return false
}

// SetSizeMBytes gets a reference to the given NullableInt32 and assigns it to the SizeMBytes field.
func (o *RedoLogGroupConfig) SetSizeMBytes(v int32) {
	o.SizeMBytes.Set(&v)
}
// SetSizeMBytesNil sets the value for SizeMBytes to be an explicit nil
func (o *RedoLogGroupConfig) SetSizeMBytesNil() {
	o.SizeMBytes.Set(nil)
}

// UnsetSizeMBytes ensures that no value is present for SizeMBytes, not even an explicit nil
func (o *RedoLogGroupConfig) UnsetSizeMBytes() {
	o.SizeMBytes.Unset()
}

// GetGroupMembers returns the GroupMembers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RedoLogGroupConfig) GetGroupMembers() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.GroupMembers
}

// GetGroupMembersOk returns a tuple with the GroupMembers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RedoLogGroupConfig) GetGroupMembersOk() (*[]string, bool) {
	if o == nil || o.GroupMembers == nil {
		return nil, false
	}
	return &o.GroupMembers, true
}

// HasGroupMembers returns a boolean if a field has been set.
func (o *RedoLogGroupConfig) HasGroupMembers() bool {
	if o != nil && o.GroupMembers != nil {
		return true
	}

	return false
}

// SetGroupMembers gets a reference to the given []string and assigns it to the GroupMembers field.
func (o *RedoLogGroupConfig) SetGroupMembers(v []string) {
	o.GroupMembers = v
}

func (o RedoLogGroupConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NumGroups.IsSet() {
		toSerialize["numGroups"] = o.NumGroups.Get()
	}
	if o.MemberPrefix.IsSet() {
		toSerialize["memberPrefix"] = o.MemberPrefix.Get()
	}
	if o.SizeMBytes.IsSet() {
		toSerialize["sizeMBytes"] = o.SizeMBytes.Get()
	}
	if o.GroupMembers != nil {
		toSerialize["groupMembers"] = o.GroupMembers
	}
	return json.Marshal(toSerialize)
}

type NullableRedoLogGroupConfig struct {
	value *RedoLogGroupConfig
	isSet bool
}

func (v NullableRedoLogGroupConfig) Get() *RedoLogGroupConfig {
	return v.value
}

func (v *NullableRedoLogGroupConfig) Set(val *RedoLogGroupConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableRedoLogGroupConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableRedoLogGroupConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedoLogGroupConfig(val *RedoLogGroupConfig) *NullableRedoLogGroupConfig {
	return &NullableRedoLogGroupConfig{value: val, isSet: true}
}

func (v NullableRedoLogGroupConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedoLogGroupConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


