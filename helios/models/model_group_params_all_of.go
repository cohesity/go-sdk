/*
 * Cohesity REST API
 *
 * Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

// package helios
package models

import (
	"encoding/json"
	. "github.com/cohesity/go-sdk/helios/utils"
)

var _ = NullableBool{}

// GroupParamsAllOf struct for GroupParamsAllOf
type GroupParamsAllOf struct {
	// Specifies the sid of the Group.
	Sid NullableString `json:"sid,omitempty"`
	// Specifies the epoch time in milliseconds when the Group was created.
	CreatedTimeMsecs NullableInt64 `json:"createdTimeMsecs,omitempty"`
	// Specifies the epoch time in milliseconds when the Group was last modified.
	LastUpdatedTimeMsecs NullableInt64 `json:"lastUpdatedTimeMsecs,omitempty"`
}

// NewGroupParamsAllOf instantiates a new GroupParamsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupParamsAllOf() *GroupParamsAllOf {
	this := GroupParamsAllOf{}
	return &this
}

// NewGroupParamsAllOfWithDefaults instantiates a new GroupParamsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupParamsAllOfWithDefaults() *GroupParamsAllOf {
	this := GroupParamsAllOf{}
	return &this
}

// GetSid returns the Sid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupParamsAllOf) GetSid() string {
	if o == nil || o.Sid.Get() == nil {
		var ret string
		return ret
	}
	return *o.Sid.Get()
}

// GetSidOk returns a tuple with the Sid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupParamsAllOf) GetSidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Sid.Get(), o.Sid.IsSet()
}

// HasSid returns a boolean if a field has been set.
func (o *GroupParamsAllOf) HasSid() bool {
	if o != nil && o.Sid.IsSet() {
		return true
	}

	return false
}

// SetSid gets a reference to the given NullableString and assigns it to the Sid field.
func (o *GroupParamsAllOf) SetSid(v string) {
	o.Sid.Set(&v)
}
// SetSidNil sets the value for Sid to be an explicit nil
func (o *GroupParamsAllOf) SetSidNil() {
	o.Sid.Set(nil)
}

// UnsetSid ensures that no value is present for Sid, not even an explicit nil
func (o *GroupParamsAllOf) UnsetSid() {
	o.Sid.Unset()
}

// GetCreatedTimeMsecs returns the CreatedTimeMsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupParamsAllOf) GetCreatedTimeMsecs() int64 {
	if o == nil || o.CreatedTimeMsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CreatedTimeMsecs.Get()
}

// GetCreatedTimeMsecsOk returns a tuple with the CreatedTimeMsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupParamsAllOf) GetCreatedTimeMsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedTimeMsecs.Get(), o.CreatedTimeMsecs.IsSet()
}

// HasCreatedTimeMsecs returns a boolean if a field has been set.
func (o *GroupParamsAllOf) HasCreatedTimeMsecs() bool {
	if o != nil && o.CreatedTimeMsecs.IsSet() {
		return true
	}

	return false
}

// SetCreatedTimeMsecs gets a reference to the given NullableInt64 and assigns it to the CreatedTimeMsecs field.
func (o *GroupParamsAllOf) SetCreatedTimeMsecs(v int64) {
	o.CreatedTimeMsecs.Set(&v)
}
// SetCreatedTimeMsecsNil sets the value for CreatedTimeMsecs to be an explicit nil
func (o *GroupParamsAllOf) SetCreatedTimeMsecsNil() {
	o.CreatedTimeMsecs.Set(nil)
}

// UnsetCreatedTimeMsecs ensures that no value is present for CreatedTimeMsecs, not even an explicit nil
func (o *GroupParamsAllOf) UnsetCreatedTimeMsecs() {
	o.CreatedTimeMsecs.Unset()
}

// GetLastUpdatedTimeMsecs returns the LastUpdatedTimeMsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupParamsAllOf) GetLastUpdatedTimeMsecs() int64 {
	if o == nil || o.LastUpdatedTimeMsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.LastUpdatedTimeMsecs.Get()
}

// GetLastUpdatedTimeMsecsOk returns a tuple with the LastUpdatedTimeMsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupParamsAllOf) GetLastUpdatedTimeMsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastUpdatedTimeMsecs.Get(), o.LastUpdatedTimeMsecs.IsSet()
}

// HasLastUpdatedTimeMsecs returns a boolean if a field has been set.
func (o *GroupParamsAllOf) HasLastUpdatedTimeMsecs() bool {
	if o != nil && o.LastUpdatedTimeMsecs.IsSet() {
		return true
	}

	return false
}

// SetLastUpdatedTimeMsecs gets a reference to the given NullableInt64 and assigns it to the LastUpdatedTimeMsecs field.
func (o *GroupParamsAllOf) SetLastUpdatedTimeMsecs(v int64) {
	o.LastUpdatedTimeMsecs.Set(&v)
}
// SetLastUpdatedTimeMsecsNil sets the value for LastUpdatedTimeMsecs to be an explicit nil
func (o *GroupParamsAllOf) SetLastUpdatedTimeMsecsNil() {
	o.LastUpdatedTimeMsecs.Set(nil)
}

// UnsetLastUpdatedTimeMsecs ensures that no value is present for LastUpdatedTimeMsecs, not even an explicit nil
func (o *GroupParamsAllOf) UnsetLastUpdatedTimeMsecs() {
	o.LastUpdatedTimeMsecs.Unset()
}

func (o GroupParamsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Sid.IsSet() {
		toSerialize["sid"] = o.Sid.Get()
	}
	if o.CreatedTimeMsecs.IsSet() {
		toSerialize["createdTimeMsecs"] = o.CreatedTimeMsecs.Get()
	}
	if o.LastUpdatedTimeMsecs.IsSet() {
		toSerialize["lastUpdatedTimeMsecs"] = o.LastUpdatedTimeMsecs.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableGroupParamsAllOf struct {
	value *GroupParamsAllOf
	isSet bool
}

func (v NullableGroupParamsAllOf) Get() *GroupParamsAllOf {
	return v.value
}

func (v *NullableGroupParamsAllOf) Set(val *GroupParamsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupParamsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupParamsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupParamsAllOf(val *GroupParamsAllOf) *NullableGroupParamsAllOf {
	return &NullableGroupParamsAllOf{value: val, isSet: true}
}

func (v NullableGroupParamsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupParamsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


