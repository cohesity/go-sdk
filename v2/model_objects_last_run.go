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

// checks if the ObjectsLastRun type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectsLastRun{}

// ObjectsLastRun Last protection run info of objects.
type ObjectsLastRun struct {
	// Specifies a list of last protection runs of objects.
	ObjectLastRuns []ObjectLastRun `json:"objectLastRuns,omitempty"`
}

// NewObjectsLastRun instantiates a new ObjectsLastRun object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectsLastRun() *ObjectsLastRun {
	this := ObjectsLastRun{}
	return &this
}

// NewObjectsLastRunWithDefaults instantiates a new ObjectsLastRun object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectsLastRunWithDefaults() *ObjectsLastRun {
	this := ObjectsLastRun{}
	return &this
}

// GetObjectLastRuns returns the ObjectLastRuns field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ObjectsLastRun) GetObjectLastRuns() []ObjectLastRun {
	if o == nil {
		var ret []ObjectLastRun
		return ret
	}
	return o.ObjectLastRuns
}

// GetObjectLastRunsOk returns a tuple with the ObjectLastRuns field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ObjectsLastRun) GetObjectLastRunsOk() ([]ObjectLastRun, bool) {
	if o == nil || IsNil(o.ObjectLastRuns) {
		return nil, false
	}
	return o.ObjectLastRuns, true
}

// HasObjectLastRuns returns a boolean if a field has been set.
func (o *ObjectsLastRun) HasObjectLastRuns() bool {
	if o != nil && !IsNil(o.ObjectLastRuns) {
		return true
	}

	return false
}

// SetObjectLastRuns gets a reference to the given []ObjectLastRun and assigns it to the ObjectLastRuns field.
func (o *ObjectsLastRun) SetObjectLastRuns(v []ObjectLastRun) {
	o.ObjectLastRuns = v
}

func (o ObjectsLastRun) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectsLastRun) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ObjectLastRuns != nil {
		toSerialize["objectLastRuns"] = o.ObjectLastRuns
	}
	return toSerialize, nil
}

type NullableObjectsLastRun struct {
	value *ObjectsLastRun
	isSet bool
}

func (v NullableObjectsLastRun) Get() *ObjectsLastRun {
	return v.value
}

func (v *NullableObjectsLastRun) Set(val *ObjectsLastRun) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectsLastRun) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectsLastRun) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectsLastRun(val *ObjectsLastRun) *NullableObjectsLastRun {
	return &NullableObjectsLastRun{value: val, isSet: true}
}

func (v NullableObjectsLastRun) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectsLastRun) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


