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

// checks if the SnmpOperationType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SnmpOperationType{}

// SnmpOperationType SNMP Operation Type
type SnmpOperationType struct {
	// Specifies the SNMP Operation Type.
	Type *string `json:"type,omitempty"`
}

// NewSnmpOperationType instantiates a new SnmpOperationType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnmpOperationType() *SnmpOperationType {
	this := SnmpOperationType{}
	return &this
}

// NewSnmpOperationTypeWithDefaults instantiates a new SnmpOperationType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnmpOperationTypeWithDefaults() *SnmpOperationType {
	this := SnmpOperationType{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SnmpOperationType) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpOperationType) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SnmpOperationType) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SnmpOperationType) SetType(v string) {
	o.Type = &v
}

func (o SnmpOperationType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SnmpOperationType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableSnmpOperationType struct {
	value *SnmpOperationType
	isSet bool
}

func (v NullableSnmpOperationType) Get() *SnmpOperationType {
	return v.value
}

func (v *NullableSnmpOperationType) Set(val *SnmpOperationType) {
	v.value = val
	v.isSet = true
}

func (v NullableSnmpOperationType) IsSet() bool {
	return v.isSet
}

func (v *NullableSnmpOperationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnmpOperationType(val *SnmpOperationType) *NullableSnmpOperationType {
	return &NullableSnmpOperationType{value: val, isSet: true}
}

func (v NullableSnmpOperationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnmpOperationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


