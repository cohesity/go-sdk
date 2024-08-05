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

// checks if the SmbFileAccessTypes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmbFileAccessTypes{}

// SmbFileAccessTypes Specifies the SMB file access types.
type SmbFileAccessTypes struct {
	// Specifies the list of SMB file access types.
	Value *string `json:"value,omitempty"`
}

// NewSmbFileAccessTypes instantiates a new SmbFileAccessTypes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmbFileAccessTypes() *SmbFileAccessTypes {
	this := SmbFileAccessTypes{}
	return &this
}

// NewSmbFileAccessTypesWithDefaults instantiates a new SmbFileAccessTypes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmbFileAccessTypesWithDefaults() *SmbFileAccessTypes {
	this := SmbFileAccessTypes{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SmbFileAccessTypes) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmbFileAccessTypes) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SmbFileAccessTypes) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *SmbFileAccessTypes) SetValue(v string) {
	o.Value = &v
}

func (o SmbFileAccessTypes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmbFileAccessTypes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableSmbFileAccessTypes struct {
	value *SmbFileAccessTypes
	isSet bool
}

func (v NullableSmbFileAccessTypes) Get() *SmbFileAccessTypes {
	return v.value
}

func (v *NullableSmbFileAccessTypes) Set(val *SmbFileAccessTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableSmbFileAccessTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableSmbFileAccessTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmbFileAccessTypes(val *SmbFileAccessTypes) *NullableSmbFileAccessTypes {
	return &NullableSmbFileAccessTypes{value: val, isSet: true}
}

func (v NullableSmbFileAccessTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmbFileAccessTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


