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

// checks if the ExternallyTriggeredClientType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternallyTriggeredClientType{}

// ExternallyTriggeredClientType Client type of an externally triggered backup.
type ExternallyTriggeredClientType struct {
	// Specifies the client type of an externally triggered backup.
	Type *string `json:"type,omitempty"`
}

// NewExternallyTriggeredClientType instantiates a new ExternallyTriggeredClientType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternallyTriggeredClientType() *ExternallyTriggeredClientType {
	this := ExternallyTriggeredClientType{}
	return &this
}

// NewExternallyTriggeredClientTypeWithDefaults instantiates a new ExternallyTriggeredClientType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternallyTriggeredClientTypeWithDefaults() *ExternallyTriggeredClientType {
	this := ExternallyTriggeredClientType{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ExternallyTriggeredClientType) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternallyTriggeredClientType) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ExternallyTriggeredClientType) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ExternallyTriggeredClientType) SetType(v string) {
	o.Type = &v
}

func (o ExternallyTriggeredClientType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternallyTriggeredClientType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableExternallyTriggeredClientType struct {
	value *ExternallyTriggeredClientType
	isSet bool
}

func (v NullableExternallyTriggeredClientType) Get() *ExternallyTriggeredClientType {
	return v.value
}

func (v *NullableExternallyTriggeredClientType) Set(val *ExternallyTriggeredClientType) {
	v.value = val
	v.isSet = true
}

func (v NullableExternallyTriggeredClientType) IsSet() bool {
	return v.isSet
}

func (v *NullableExternallyTriggeredClientType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternallyTriggeredClientType(val *ExternallyTriggeredClientType) *NullableExternallyTriggeredClientType {
	return &NullableExternallyTriggeredClientType{value: val, isSet: true}
}

func (v NullableExternallyTriggeredClientType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternallyTriggeredClientType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


