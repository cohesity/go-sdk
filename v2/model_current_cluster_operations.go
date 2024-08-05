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

// checks if the CurrentClusterOperations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CurrentClusterOperations{}

// CurrentClusterOperations Specifies the current operations in a cluster.
type CurrentClusterOperations struct {
	// Specifies the list of cluster operations.
	Value *string `json:"value,omitempty"`
}

// NewCurrentClusterOperations instantiates a new CurrentClusterOperations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCurrentClusterOperations() *CurrentClusterOperations {
	this := CurrentClusterOperations{}
	return &this
}

// NewCurrentClusterOperationsWithDefaults instantiates a new CurrentClusterOperations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCurrentClusterOperationsWithDefaults() *CurrentClusterOperations {
	this := CurrentClusterOperations{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CurrentClusterOperations) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentClusterOperations) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CurrentClusterOperations) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *CurrentClusterOperations) SetValue(v string) {
	o.Value = &v
}

func (o CurrentClusterOperations) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CurrentClusterOperations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableCurrentClusterOperations struct {
	value *CurrentClusterOperations
	isSet bool
}

func (v NullableCurrentClusterOperations) Get() *CurrentClusterOperations {
	return v.value
}

func (v *NullableCurrentClusterOperations) Set(val *CurrentClusterOperations) {
	v.value = val
	v.isSet = true
}

func (v NullableCurrentClusterOperations) IsSet() bool {
	return v.isSet
}

func (v *NullableCurrentClusterOperations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCurrentClusterOperations(val *CurrentClusterOperations) *NullableCurrentClusterOperations {
	return &NullableCurrentClusterOperations{value: val, isSet: true}
}

func (v NullableCurrentClusterOperations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCurrentClusterOperations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


