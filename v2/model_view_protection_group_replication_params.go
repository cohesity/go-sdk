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

// checks if the ViewProtectionGroupReplicationParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ViewProtectionGroupReplicationParams{}

// ViewProtectionGroupReplicationParams Specifies the parameters for view replication.
type ViewProtectionGroupReplicationParams struct {
	// Specifies the list of remote view names for the protected views in the Protection Group. By default the names will be the same as the name of the protected view.
	ViewNameConfigList []ReplicatedViewNameConfig `json:"viewNameConfigList,omitempty"`
}

// NewViewProtectionGroupReplicationParams instantiates a new ViewProtectionGroupReplicationParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViewProtectionGroupReplicationParams() *ViewProtectionGroupReplicationParams {
	this := ViewProtectionGroupReplicationParams{}
	return &this
}

// NewViewProtectionGroupReplicationParamsWithDefaults instantiates a new ViewProtectionGroupReplicationParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewProtectionGroupReplicationParamsWithDefaults() *ViewProtectionGroupReplicationParams {
	this := ViewProtectionGroupReplicationParams{}
	return &this
}

// GetViewNameConfigList returns the ViewNameConfigList field value if set, zero value otherwise.
func (o *ViewProtectionGroupReplicationParams) GetViewNameConfigList() []ReplicatedViewNameConfig {
	if o == nil || IsNil(o.ViewNameConfigList) {
		var ret []ReplicatedViewNameConfig
		return ret
	}
	return o.ViewNameConfigList
}

// GetViewNameConfigListOk returns a tuple with the ViewNameConfigList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewProtectionGroupReplicationParams) GetViewNameConfigListOk() ([]ReplicatedViewNameConfig, bool) {
	if o == nil || IsNil(o.ViewNameConfigList) {
		return nil, false
	}
	return o.ViewNameConfigList, true
}

// HasViewNameConfigList returns a boolean if a field has been set.
func (o *ViewProtectionGroupReplicationParams) HasViewNameConfigList() bool {
	if o != nil && !IsNil(o.ViewNameConfigList) {
		return true
	}

	return false
}

// SetViewNameConfigList gets a reference to the given []ReplicatedViewNameConfig and assigns it to the ViewNameConfigList field.
func (o *ViewProtectionGroupReplicationParams) SetViewNameConfigList(v []ReplicatedViewNameConfig) {
	o.ViewNameConfigList = v
}

func (o ViewProtectionGroupReplicationParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ViewProtectionGroupReplicationParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ViewNameConfigList) {
		toSerialize["viewNameConfigList"] = o.ViewNameConfigList
	}
	return toSerialize, nil
}

type NullableViewProtectionGroupReplicationParams struct {
	value *ViewProtectionGroupReplicationParams
	isSet bool
}

func (v NullableViewProtectionGroupReplicationParams) Get() *ViewProtectionGroupReplicationParams {
	return v.value
}

func (v *NullableViewProtectionGroupReplicationParams) Set(val *ViewProtectionGroupReplicationParams) {
	v.value = val
	v.isSet = true
}

func (v NullableViewProtectionGroupReplicationParams) IsSet() bool {
	return v.isSet
}

func (v *NullableViewProtectionGroupReplicationParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViewProtectionGroupReplicationParams(val *ViewProtectionGroupReplicationParams) *NullableViewProtectionGroupReplicationParams {
	return &NullableViewProtectionGroupReplicationParams{value: val, isSet: true}
}

func (v NullableViewProtectionGroupReplicationParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViewProtectionGroupReplicationParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


