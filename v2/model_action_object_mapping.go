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

// checks if the ActionObjectMapping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActionObjectMapping{}

// ActionObjectMapping Specifies the object paring for performing action on list of objects.
type ActionObjectMapping struct {
	// Specifies the destination object id.
	DestinationObjectId NullableInt64 `json:"destinationObjectId,omitempty"`
	// Specifies the source object id.
	SourceObjectId NullableInt64 `json:"sourceObjectId,omitempty"`
}

// NewActionObjectMapping instantiates a new ActionObjectMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionObjectMapping() *ActionObjectMapping {
	this := ActionObjectMapping{}
	return &this
}

// NewActionObjectMappingWithDefaults instantiates a new ActionObjectMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionObjectMappingWithDefaults() *ActionObjectMapping {
	this := ActionObjectMapping{}
	return &this
}

// GetDestinationObjectId returns the DestinationObjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActionObjectMapping) GetDestinationObjectId() int64 {
	if o == nil || IsNil(o.DestinationObjectId.Get()) {
		var ret int64
		return ret
	}
	return *o.DestinationObjectId.Get()
}

// GetDestinationObjectIdOk returns a tuple with the DestinationObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActionObjectMapping) GetDestinationObjectIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DestinationObjectId.Get(), o.DestinationObjectId.IsSet()
}

// HasDestinationObjectId returns a boolean if a field has been set.
func (o *ActionObjectMapping) HasDestinationObjectId() bool {
	if o != nil && o.DestinationObjectId.IsSet() {
		return true
	}

	return false
}

// SetDestinationObjectId gets a reference to the given NullableInt64 and assigns it to the DestinationObjectId field.
func (o *ActionObjectMapping) SetDestinationObjectId(v int64) {
	o.DestinationObjectId.Set(&v)
}
// SetDestinationObjectIdNil sets the value for DestinationObjectId to be an explicit nil
func (o *ActionObjectMapping) SetDestinationObjectIdNil() {
	o.DestinationObjectId.Set(nil)
}

// UnsetDestinationObjectId ensures that no value is present for DestinationObjectId, not even an explicit nil
func (o *ActionObjectMapping) UnsetDestinationObjectId() {
	o.DestinationObjectId.Unset()
}

// GetSourceObjectId returns the SourceObjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActionObjectMapping) GetSourceObjectId() int64 {
	if o == nil || IsNil(o.SourceObjectId.Get()) {
		var ret int64
		return ret
	}
	return *o.SourceObjectId.Get()
}

// GetSourceObjectIdOk returns a tuple with the SourceObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActionObjectMapping) GetSourceObjectIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceObjectId.Get(), o.SourceObjectId.IsSet()
}

// HasSourceObjectId returns a boolean if a field has been set.
func (o *ActionObjectMapping) HasSourceObjectId() bool {
	if o != nil && o.SourceObjectId.IsSet() {
		return true
	}

	return false
}

// SetSourceObjectId gets a reference to the given NullableInt64 and assigns it to the SourceObjectId field.
func (o *ActionObjectMapping) SetSourceObjectId(v int64) {
	o.SourceObjectId.Set(&v)
}
// SetSourceObjectIdNil sets the value for SourceObjectId to be an explicit nil
func (o *ActionObjectMapping) SetSourceObjectIdNil() {
	o.SourceObjectId.Set(nil)
}

// UnsetSourceObjectId ensures that no value is present for SourceObjectId, not even an explicit nil
func (o *ActionObjectMapping) UnsetSourceObjectId() {
	o.SourceObjectId.Unset()
}

func (o ActionObjectMapping) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActionObjectMapping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DestinationObjectId.IsSet() {
		toSerialize["destinationObjectId"] = o.DestinationObjectId.Get()
	}
	if o.SourceObjectId.IsSet() {
		toSerialize["sourceObjectId"] = o.SourceObjectId.Get()
	}
	return toSerialize, nil
}

type NullableActionObjectMapping struct {
	value *ActionObjectMapping
	isSet bool
}

func (v NullableActionObjectMapping) Get() *ActionObjectMapping {
	return v.value
}

func (v *NullableActionObjectMapping) Set(val *ActionObjectMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableActionObjectMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableActionObjectMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionObjectMapping(val *ActionObjectMapping) *NullableActionObjectMapping {
	return &NullableActionObjectMapping{value: val, isSet: true}
}

func (v NullableActionObjectMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionObjectMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


