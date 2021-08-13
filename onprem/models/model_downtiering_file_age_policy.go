/*
 * Cohesity REST API
 *
 * Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

// package onprem
package models

import (
	"encoding/json"
	. "github.com/cohesity/go-sdk/onprem/utils"
	"fmt"
)

var _ = NullableBool{}

// DowntieringFileAgePolicy Specifies the file's selection rule by file age for down tiering data tiering task eg. 1. select files older than 10 days. 2. select files last accessed 2 weeks ago. 3. select files last modified 1 month ago.
type DowntieringFileAgePolicy struct {
	// Specifies the condition for the file age.
	Condition NullableString `json:"condition,omitempty"`
	// Specifies the number of msecs used for file selection.
	AgeMsecs NullableInt64 `json:"ageMsecs,omitempty"`
}

// NewDowntieringFileAgePolicy instantiates a new DowntieringFileAgePolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDowntieringFileAgePolicy() *DowntieringFileAgePolicy {
	this := DowntieringFileAgePolicy{}
	return &this
}

// NewDowntieringFileAgePolicyWithDefaults instantiates a new DowntieringFileAgePolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDowntieringFileAgePolicyWithDefaults() *DowntieringFileAgePolicy {
	this := DowntieringFileAgePolicy{}
	return &this
}

// GetCondition returns the Condition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DowntieringFileAgePolicy) GetCondition() string {
	if o == nil || o.Condition.Get() == nil {
		var ret string
		return ret
	}
	return *o.Condition.Get()
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DowntieringFileAgePolicy) GetConditionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Condition.Get(), o.Condition.IsSet()
}

// HasCondition returns a boolean if a field has been set.
func (o *DowntieringFileAgePolicy) HasCondition() bool {
	if o != nil && o.Condition.IsSet() {
		return true
	}

	return false
}

// SetCondition gets a reference to the given NullableString and assigns it to the Condition field.
func (o *DowntieringFileAgePolicy) SetCondition(v string) {
	o.Condition.Set(&v)
}
// SetConditionNil sets the value for Condition to be an explicit nil
func (o *DowntieringFileAgePolicy) SetConditionNil() {
	o.Condition.Set(nil)
}

// UnsetCondition ensures that no value is present for Condition, not even an explicit nil
func (o *DowntieringFileAgePolicy) UnsetCondition() {
	o.Condition.Unset()
}

// GetAgeMsecs returns the AgeMsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DowntieringFileAgePolicy) GetAgeMsecs() int64 {
	if o == nil || o.AgeMsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.AgeMsecs.Get()
}

// GetAgeMsecsOk returns a tuple with the AgeMsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DowntieringFileAgePolicy) GetAgeMsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AgeMsecs.Get(), o.AgeMsecs.IsSet()
}

// HasAgeMsecs returns a boolean if a field has been set.
func (o *DowntieringFileAgePolicy) HasAgeMsecs() bool {
	if o != nil && o.AgeMsecs.IsSet() {
		return true
	}

	return false
}

// SetAgeMsecs gets a reference to the given NullableInt64 and assigns it to the AgeMsecs field.
func (o *DowntieringFileAgePolicy) SetAgeMsecs(v int64) {
	o.AgeMsecs.Set(&v)
}
// SetAgeMsecsNil sets the value for AgeMsecs to be an explicit nil
func (o *DowntieringFileAgePolicy) SetAgeMsecsNil() {
	o.AgeMsecs.Set(nil)
}

// UnsetAgeMsecs ensures that no value is present for AgeMsecs, not even an explicit nil
func (o *DowntieringFileAgePolicy) UnsetAgeMsecs() {
	o.AgeMsecs.Unset()
}

func (o DowntieringFileAgePolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Condition.IsSet() {
		toSerialize["condition"] = o.Condition.Get()
	}
	if o.AgeMsecs.IsSet() {
		toSerialize["ageMsecs"] = o.AgeMsecs.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableDowntieringFileAgePolicy struct {
	value *DowntieringFileAgePolicy
	isSet bool
}

func (v NullableDowntieringFileAgePolicy) Get() *DowntieringFileAgePolicy {
	return v.value
}

func (v *NullableDowntieringFileAgePolicy) Set(val *DowntieringFileAgePolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableDowntieringFileAgePolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableDowntieringFileAgePolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDowntieringFileAgePolicy(val *DowntieringFileAgePolicy) *NullableDowntieringFileAgePolicy {
	return &NullableDowntieringFileAgePolicy{value: val, isSet: true}
}

func (v NullableDowntieringFileAgePolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDowntieringFileAgePolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o DowntieringFileAgePolicy) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}