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

// checks if the DeleteInfectedObjects type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteInfectedObjects{}

// DeleteInfectedObjects Specifies a list of infected objects.
type DeleteInfectedObjects struct {
	// Specifies the list of infected objects that failed deletion.
	DeleteFailedInfectedObjects []InfectedObject `json:"deleteFailedInfectedObjects,omitempty"`
	// Specifies the list of infected objects that are successfully deleted.
	DeleteSucceededInfectedObjects []InfectedObject `json:"deleteSucceededInfectedObjects,omitempty"`
}

// NewDeleteInfectedObjects instantiates a new DeleteInfectedObjects object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteInfectedObjects() *DeleteInfectedObjects {
	this := DeleteInfectedObjects{}
	return &this
}

// NewDeleteInfectedObjectsWithDefaults instantiates a new DeleteInfectedObjects object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteInfectedObjectsWithDefaults() *DeleteInfectedObjects {
	this := DeleteInfectedObjects{}
	return &this
}

// GetDeleteFailedInfectedObjects returns the DeleteFailedInfectedObjects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeleteInfectedObjects) GetDeleteFailedInfectedObjects() []InfectedObject {
	if o == nil {
		var ret []InfectedObject
		return ret
	}
	return o.DeleteFailedInfectedObjects
}

// GetDeleteFailedInfectedObjectsOk returns a tuple with the DeleteFailedInfectedObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeleteInfectedObjects) GetDeleteFailedInfectedObjectsOk() ([]InfectedObject, bool) {
	if o == nil || IsNil(o.DeleteFailedInfectedObjects) {
		return nil, false
	}
	return o.DeleteFailedInfectedObjects, true
}

// HasDeleteFailedInfectedObjects returns a boolean if a field has been set.
func (o *DeleteInfectedObjects) HasDeleteFailedInfectedObjects() bool {
	if o != nil && !IsNil(o.DeleteFailedInfectedObjects) {
		return true
	}

	return false
}

// SetDeleteFailedInfectedObjects gets a reference to the given []InfectedObject and assigns it to the DeleteFailedInfectedObjects field.
func (o *DeleteInfectedObjects) SetDeleteFailedInfectedObjects(v []InfectedObject) {
	o.DeleteFailedInfectedObjects = v
}

// GetDeleteSucceededInfectedObjects returns the DeleteSucceededInfectedObjects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeleteInfectedObjects) GetDeleteSucceededInfectedObjects() []InfectedObject {
	if o == nil {
		var ret []InfectedObject
		return ret
	}
	return o.DeleteSucceededInfectedObjects
}

// GetDeleteSucceededInfectedObjectsOk returns a tuple with the DeleteSucceededInfectedObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeleteInfectedObjects) GetDeleteSucceededInfectedObjectsOk() ([]InfectedObject, bool) {
	if o == nil || IsNil(o.DeleteSucceededInfectedObjects) {
		return nil, false
	}
	return o.DeleteSucceededInfectedObjects, true
}

// HasDeleteSucceededInfectedObjects returns a boolean if a field has been set.
func (o *DeleteInfectedObjects) HasDeleteSucceededInfectedObjects() bool {
	if o != nil && !IsNil(o.DeleteSucceededInfectedObjects) {
		return true
	}

	return false
}

// SetDeleteSucceededInfectedObjects gets a reference to the given []InfectedObject and assigns it to the DeleteSucceededInfectedObjects field.
func (o *DeleteInfectedObjects) SetDeleteSucceededInfectedObjects(v []InfectedObject) {
	o.DeleteSucceededInfectedObjects = v
}

func (o DeleteInfectedObjects) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteInfectedObjects) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DeleteFailedInfectedObjects != nil {
		toSerialize["deleteFailedInfectedObjects"] = o.DeleteFailedInfectedObjects
	}
	if o.DeleteSucceededInfectedObjects != nil {
		toSerialize["deleteSucceededInfectedObjects"] = o.DeleteSucceededInfectedObjects
	}
	return toSerialize, nil
}

type NullableDeleteInfectedObjects struct {
	value *DeleteInfectedObjects
	isSet bool
}

func (v NullableDeleteInfectedObjects) Get() *DeleteInfectedObjects {
	return v.value
}

func (v *NullableDeleteInfectedObjects) Set(val *DeleteInfectedObjects) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteInfectedObjects) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteInfectedObjects) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteInfectedObjects(val *DeleteInfectedObjects) *NullableDeleteInfectedObjects {
	return &NullableDeleteInfectedObjects{value: val, isSet: true}
}

func (v NullableDeleteInfectedObjects) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteInfectedObjects) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


