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

// checks if the NonCurrentVersionExpirationAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NonCurrentVersionExpirationAction{}

// NonCurrentVersionExpirationAction Specifies the Lifecycle Non-current Version Expiration Action.
type NonCurrentVersionExpirationAction struct {
	// Specifies the number of days an object is non-current before performing the associated action.
	Days NullableInt64 `json:"days,omitempty"`
}

// NewNonCurrentVersionExpirationAction instantiates a new NonCurrentVersionExpirationAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNonCurrentVersionExpirationAction() *NonCurrentVersionExpirationAction {
	this := NonCurrentVersionExpirationAction{}
	return &this
}

// NewNonCurrentVersionExpirationActionWithDefaults instantiates a new NonCurrentVersionExpirationAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNonCurrentVersionExpirationActionWithDefaults() *NonCurrentVersionExpirationAction {
	this := NonCurrentVersionExpirationAction{}
	return &this
}

// GetDays returns the Days field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NonCurrentVersionExpirationAction) GetDays() int64 {
	if o == nil || IsNil(o.Days.Get()) {
		var ret int64
		return ret
	}
	return *o.Days.Get()
}

// GetDaysOk returns a tuple with the Days field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NonCurrentVersionExpirationAction) GetDaysOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Days.Get(), o.Days.IsSet()
}

// HasDays returns a boolean if a field has been set.
func (o *NonCurrentVersionExpirationAction) HasDays() bool {
	if o != nil && o.Days.IsSet() {
		return true
	}

	return false
}

// SetDays gets a reference to the given NullableInt64 and assigns it to the Days field.
func (o *NonCurrentVersionExpirationAction) SetDays(v int64) {
	o.Days.Set(&v)
}
// SetDaysNil sets the value for Days to be an explicit nil
func (o *NonCurrentVersionExpirationAction) SetDaysNil() {
	o.Days.Set(nil)
}

// UnsetDays ensures that no value is present for Days, not even an explicit nil
func (o *NonCurrentVersionExpirationAction) UnsetDays() {
	o.Days.Unset()
}

func (o NonCurrentVersionExpirationAction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NonCurrentVersionExpirationAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Days.IsSet() {
		toSerialize["days"] = o.Days.Get()
	}
	return toSerialize, nil
}

type NullableNonCurrentVersionExpirationAction struct {
	value *NonCurrentVersionExpirationAction
	isSet bool
}

func (v NullableNonCurrentVersionExpirationAction) Get() *NonCurrentVersionExpirationAction {
	return v.value
}

func (v *NullableNonCurrentVersionExpirationAction) Set(val *NonCurrentVersionExpirationAction) {
	v.value = val
	v.isSet = true
}

func (v NullableNonCurrentVersionExpirationAction) IsSet() bool {
	return v.isSet
}

func (v *NullableNonCurrentVersionExpirationAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonCurrentVersionExpirationAction(val *NonCurrentVersionExpirationAction) *NullableNonCurrentVersionExpirationAction {
	return &NullableNonCurrentVersionExpirationAction{value: val, isSet: true}
}

func (v NullableNonCurrentVersionExpirationAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonCurrentVersionExpirationAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


