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

// checks if the UdaRecoveryParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UdaRecoveryParams{}

// UdaRecoveryParams Parameters to customize recovery workflow.
type UdaRecoveryParams struct {
	DynamicForm NullableUdaDynamicFormParams `json:"dynamicForm,omitempty"`
	PrimaryFields NullableUdaRecoveryParamsPrimaryFields `json:"primaryFields,omitempty"`
}

// NewUdaRecoveryParams instantiates a new UdaRecoveryParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUdaRecoveryParams() *UdaRecoveryParams {
	this := UdaRecoveryParams{}
	return &this
}

// NewUdaRecoveryParamsWithDefaults instantiates a new UdaRecoveryParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUdaRecoveryParamsWithDefaults() *UdaRecoveryParams {
	this := UdaRecoveryParams{}
	return &this
}

// GetDynamicForm returns the DynamicForm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UdaRecoveryParams) GetDynamicForm() UdaDynamicFormParams {
	if o == nil || IsNil(o.DynamicForm.Get()) {
		var ret UdaDynamicFormParams
		return ret
	}
	return *o.DynamicForm.Get()
}

// GetDynamicFormOk returns a tuple with the DynamicForm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UdaRecoveryParams) GetDynamicFormOk() (*UdaDynamicFormParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.DynamicForm.Get(), o.DynamicForm.IsSet()
}

// HasDynamicForm returns a boolean if a field has been set.
func (o *UdaRecoveryParams) HasDynamicForm() bool {
	if o != nil && o.DynamicForm.IsSet() {
		return true
	}

	return false
}

// SetDynamicForm gets a reference to the given NullableUdaDynamicFormParams and assigns it to the DynamicForm field.
func (o *UdaRecoveryParams) SetDynamicForm(v UdaDynamicFormParams) {
	o.DynamicForm.Set(&v)
}
// SetDynamicFormNil sets the value for DynamicForm to be an explicit nil
func (o *UdaRecoveryParams) SetDynamicFormNil() {
	o.DynamicForm.Set(nil)
}

// UnsetDynamicForm ensures that no value is present for DynamicForm, not even an explicit nil
func (o *UdaRecoveryParams) UnsetDynamicForm() {
	o.DynamicForm.Unset()
}

// GetPrimaryFields returns the PrimaryFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UdaRecoveryParams) GetPrimaryFields() UdaRecoveryParamsPrimaryFields {
	if o == nil || IsNil(o.PrimaryFields.Get()) {
		var ret UdaRecoveryParamsPrimaryFields
		return ret
	}
	return *o.PrimaryFields.Get()
}

// GetPrimaryFieldsOk returns a tuple with the PrimaryFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UdaRecoveryParams) GetPrimaryFieldsOk() (*UdaRecoveryParamsPrimaryFields, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryFields.Get(), o.PrimaryFields.IsSet()
}

// HasPrimaryFields returns a boolean if a field has been set.
func (o *UdaRecoveryParams) HasPrimaryFields() bool {
	if o != nil && o.PrimaryFields.IsSet() {
		return true
	}

	return false
}

// SetPrimaryFields gets a reference to the given NullableUdaRecoveryParamsPrimaryFields and assigns it to the PrimaryFields field.
func (o *UdaRecoveryParams) SetPrimaryFields(v UdaRecoveryParamsPrimaryFields) {
	o.PrimaryFields.Set(&v)
}
// SetPrimaryFieldsNil sets the value for PrimaryFields to be an explicit nil
func (o *UdaRecoveryParams) SetPrimaryFieldsNil() {
	o.PrimaryFields.Set(nil)
}

// UnsetPrimaryFields ensures that no value is present for PrimaryFields, not even an explicit nil
func (o *UdaRecoveryParams) UnsetPrimaryFields() {
	o.PrimaryFields.Unset()
}

func (o UdaRecoveryParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UdaRecoveryParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DynamicForm.IsSet() {
		toSerialize["dynamicForm"] = o.DynamicForm.Get()
	}
	if o.PrimaryFields.IsSet() {
		toSerialize["primaryFields"] = o.PrimaryFields.Get()
	}
	return toSerialize, nil
}

type NullableUdaRecoveryParams struct {
	value *UdaRecoveryParams
	isSet bool
}

func (v NullableUdaRecoveryParams) Get() *UdaRecoveryParams {
	return v.value
}

func (v *NullableUdaRecoveryParams) Set(val *UdaRecoveryParams) {
	v.value = val
	v.isSet = true
}

func (v NullableUdaRecoveryParams) IsSet() bool {
	return v.isSet
}

func (v *NullableUdaRecoveryParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUdaRecoveryParams(val *UdaRecoveryParams) *NullableUdaRecoveryParams {
	return &NullableUdaRecoveryParams{value: val, isSet: true}
}

func (v NullableUdaRecoveryParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUdaRecoveryParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


