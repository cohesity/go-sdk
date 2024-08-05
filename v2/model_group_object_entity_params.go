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

// checks if the GroupObjectEntityParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupObjectEntityParams{}

// GroupObjectEntityParams Specifies the common parameters for Office365 Group objects.
type GroupObjectEntityParams struct {
	// Specifies whether the Group is mail enabled. Mail enabled groups are used within Microsoft to distribute messages.
	IsMailEnabled NullableBool `json:"isMailEnabled,omitempty"`
	// Specifies whether the Group is security enabled. Security enabled groups are used to grant access permissions to resources in Exchange and Active Directory.
	IsSecurityEnabled NullableBool `json:"isSecurityEnabled,omitempty"`
	// Specifies the count of members within the Group.
	MemberCount NullableInt64 `json:"memberCount,omitempty"`
}

// NewGroupObjectEntityParams instantiates a new GroupObjectEntityParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupObjectEntityParams() *GroupObjectEntityParams {
	this := GroupObjectEntityParams{}
	return &this
}

// NewGroupObjectEntityParamsWithDefaults instantiates a new GroupObjectEntityParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupObjectEntityParamsWithDefaults() *GroupObjectEntityParams {
	this := GroupObjectEntityParams{}
	return &this
}

// GetIsMailEnabled returns the IsMailEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupObjectEntityParams) GetIsMailEnabled() bool {
	if o == nil || IsNil(o.IsMailEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.IsMailEnabled.Get()
}

// GetIsMailEnabledOk returns a tuple with the IsMailEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupObjectEntityParams) GetIsMailEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsMailEnabled.Get(), o.IsMailEnabled.IsSet()
}

// HasIsMailEnabled returns a boolean if a field has been set.
func (o *GroupObjectEntityParams) HasIsMailEnabled() bool {
	if o != nil && o.IsMailEnabled.IsSet() {
		return true
	}

	return false
}

// SetIsMailEnabled gets a reference to the given NullableBool and assigns it to the IsMailEnabled field.
func (o *GroupObjectEntityParams) SetIsMailEnabled(v bool) {
	o.IsMailEnabled.Set(&v)
}
// SetIsMailEnabledNil sets the value for IsMailEnabled to be an explicit nil
func (o *GroupObjectEntityParams) SetIsMailEnabledNil() {
	o.IsMailEnabled.Set(nil)
}

// UnsetIsMailEnabled ensures that no value is present for IsMailEnabled, not even an explicit nil
func (o *GroupObjectEntityParams) UnsetIsMailEnabled() {
	o.IsMailEnabled.Unset()
}

// GetIsSecurityEnabled returns the IsSecurityEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupObjectEntityParams) GetIsSecurityEnabled() bool {
	if o == nil || IsNil(o.IsSecurityEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.IsSecurityEnabled.Get()
}

// GetIsSecurityEnabledOk returns a tuple with the IsSecurityEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupObjectEntityParams) GetIsSecurityEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsSecurityEnabled.Get(), o.IsSecurityEnabled.IsSet()
}

// HasIsSecurityEnabled returns a boolean if a field has been set.
func (o *GroupObjectEntityParams) HasIsSecurityEnabled() bool {
	if o != nil && o.IsSecurityEnabled.IsSet() {
		return true
	}

	return false
}

// SetIsSecurityEnabled gets a reference to the given NullableBool and assigns it to the IsSecurityEnabled field.
func (o *GroupObjectEntityParams) SetIsSecurityEnabled(v bool) {
	o.IsSecurityEnabled.Set(&v)
}
// SetIsSecurityEnabledNil sets the value for IsSecurityEnabled to be an explicit nil
func (o *GroupObjectEntityParams) SetIsSecurityEnabledNil() {
	o.IsSecurityEnabled.Set(nil)
}

// UnsetIsSecurityEnabled ensures that no value is present for IsSecurityEnabled, not even an explicit nil
func (o *GroupObjectEntityParams) UnsetIsSecurityEnabled() {
	o.IsSecurityEnabled.Unset()
}

// GetMemberCount returns the MemberCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupObjectEntityParams) GetMemberCount() int64 {
	if o == nil || IsNil(o.MemberCount.Get()) {
		var ret int64
		return ret
	}
	return *o.MemberCount.Get()
}

// GetMemberCountOk returns a tuple with the MemberCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupObjectEntityParams) GetMemberCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MemberCount.Get(), o.MemberCount.IsSet()
}

// HasMemberCount returns a boolean if a field has been set.
func (o *GroupObjectEntityParams) HasMemberCount() bool {
	if o != nil && o.MemberCount.IsSet() {
		return true
	}

	return false
}

// SetMemberCount gets a reference to the given NullableInt64 and assigns it to the MemberCount field.
func (o *GroupObjectEntityParams) SetMemberCount(v int64) {
	o.MemberCount.Set(&v)
}
// SetMemberCountNil sets the value for MemberCount to be an explicit nil
func (o *GroupObjectEntityParams) SetMemberCountNil() {
	o.MemberCount.Set(nil)
}

// UnsetMemberCount ensures that no value is present for MemberCount, not even an explicit nil
func (o *GroupObjectEntityParams) UnsetMemberCount() {
	o.MemberCount.Unset()
}

func (o GroupObjectEntityParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupObjectEntityParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.IsMailEnabled.IsSet() {
		toSerialize["isMailEnabled"] = o.IsMailEnabled.Get()
	}
	if o.IsSecurityEnabled.IsSet() {
		toSerialize["isSecurityEnabled"] = o.IsSecurityEnabled.Get()
	}
	if o.MemberCount.IsSet() {
		toSerialize["memberCount"] = o.MemberCount.Get()
	}
	return toSerialize, nil
}

type NullableGroupObjectEntityParams struct {
	value *GroupObjectEntityParams
	isSet bool
}

func (v NullableGroupObjectEntityParams) Get() *GroupObjectEntityParams {
	return v.value
}

func (v *NullableGroupObjectEntityParams) Set(val *GroupObjectEntityParams) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupObjectEntityParams) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupObjectEntityParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupObjectEntityParams(val *GroupObjectEntityParams) *NullableGroupObjectEntityParams {
	return &NullableGroupObjectEntityParams{value: val, isSet: true}
}

func (v NullableGroupObjectEntityParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupObjectEntityParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


