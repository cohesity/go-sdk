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

// checks if the Objects type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Objects{}

// Objects Specifies the list of objects.
type Objects struct {
	// Specifies the list of objects.
	Objects []Object `json:"objects,omitempty"`
}

// NewObjects instantiates a new Objects object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjects() *Objects {
	this := Objects{}
	return &this
}

// NewObjectsWithDefaults instantiates a new Objects object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectsWithDefaults() *Objects {
	this := Objects{}
	return &this
}

// GetObjects returns the Objects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Objects) GetObjects() []Object {
	if o == nil {
		var ret []Object
		return ret
	}
	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Objects) GetObjectsOk() ([]Object, bool) {
	if o == nil || IsNil(o.Objects) {
		return nil, false
	}
	return o.Objects, true
}

// HasObjects returns a boolean if a field has been set.
func (o *Objects) HasObjects() bool {
	if o != nil && !IsNil(o.Objects) {
		return true
	}

	return false
}

// SetObjects gets a reference to the given []Object and assigns it to the Objects field.
func (o *Objects) SetObjects(v []Object) {
	o.Objects = v
}

func (o Objects) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Objects) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Objects != nil {
		toSerialize["objects"] = o.Objects
	}
	return toSerialize, nil
}

type NullableObjects struct {
	value *Objects
	isSet bool
}

func (v NullableObjects) Get() *Objects {
	return v.value
}

func (v *NullableObjects) Set(val *Objects) {
	v.value = val
	v.isSet = true
}

func (v NullableObjects) IsSet() bool {
	return v.isSet
}

func (v *NullableObjects) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjects(val *Objects) *NullableObjects {
	return &NullableObjects{value: val, isSet: true}
}

func (v NullableObjects) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjects) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


