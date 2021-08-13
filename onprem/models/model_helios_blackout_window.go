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

// HeliosBlackoutWindow List of Blackout Windows. If specified, this field defines blackout periods when backups are not triggered..
type HeliosBlackoutWindow struct {
	// Specifies a day in the week when no new Protection Group Runs should be started such as 'Sunday'. Specifies a day in a week such as 'Sunday', 'Monday', etc.
	Day NullableString `json:"day"`
	StartTime TimeOfDay `json:"startTime"`
	EndTime TimeOfDay `json:"endTime"`
	// Specifies the unique identifier for the blackout getting added. This field should only be set if policy is getting updated.
	ConfigId NullableString `json:"configId,omitempty"`
}

// NewHeliosBlackoutWindow instantiates a new HeliosBlackoutWindow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeliosBlackoutWindow(day NullableString, startTime TimeOfDay, endTime TimeOfDay) *HeliosBlackoutWindow {
	this := HeliosBlackoutWindow{}
	this.Day = day
	this.StartTime = startTime
	this.EndTime = endTime
	return &this
}

// NewHeliosBlackoutWindowWithDefaults instantiates a new HeliosBlackoutWindow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeliosBlackoutWindowWithDefaults() *HeliosBlackoutWindow {
	this := HeliosBlackoutWindow{}
	return &this
}

// GetDay returns the Day field value
// If the value is explicit nil, the zero value for string will be returned
func (o *HeliosBlackoutWindow) GetDay() string {
	if o == nil || o.Day.Get() == nil {
		var ret string
		return ret
	}

	return *o.Day.Get()
}

// GetDayOk returns a tuple with the Day field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeliosBlackoutWindow) GetDayOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Day.Get(), o.Day.IsSet()
}

// SetDay sets field value
func (o *HeliosBlackoutWindow) SetDay(v string) {
	o.Day.Set(&v)
}

// GetStartTime returns the StartTime field value
func (o *HeliosBlackoutWindow) GetStartTime() TimeOfDay {
	if o == nil {
		var ret TimeOfDay
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *HeliosBlackoutWindow) GetStartTimeOk() (*TimeOfDay, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *HeliosBlackoutWindow) SetStartTime(v TimeOfDay) {
	o.StartTime = v
}

// GetEndTime returns the EndTime field value
func (o *HeliosBlackoutWindow) GetEndTime() TimeOfDay {
	if o == nil {
		var ret TimeOfDay
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *HeliosBlackoutWindow) GetEndTimeOk() (*TimeOfDay, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *HeliosBlackoutWindow) SetEndTime(v TimeOfDay) {
	o.EndTime = v
}

// GetConfigId returns the ConfigId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeliosBlackoutWindow) GetConfigId() string {
	if o == nil || o.ConfigId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConfigId.Get()
}

// GetConfigIdOk returns a tuple with the ConfigId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeliosBlackoutWindow) GetConfigIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ConfigId.Get(), o.ConfigId.IsSet()
}

// HasConfigId returns a boolean if a field has been set.
func (o *HeliosBlackoutWindow) HasConfigId() bool {
	if o != nil && o.ConfigId.IsSet() {
		return true
	}

	return false
}

// SetConfigId gets a reference to the given NullableString and assigns it to the ConfigId field.
func (o *HeliosBlackoutWindow) SetConfigId(v string) {
	o.ConfigId.Set(&v)
}
// SetConfigIdNil sets the value for ConfigId to be an explicit nil
func (o *HeliosBlackoutWindow) SetConfigIdNil() {
	o.ConfigId.Set(nil)
}

// UnsetConfigId ensures that no value is present for ConfigId, not even an explicit nil
func (o *HeliosBlackoutWindow) UnsetConfigId() {
	o.ConfigId.Unset()
}

func (o HeliosBlackoutWindow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["day"] = o.Day.Get()
	}
	if true {
		toSerialize["startTime"] = o.StartTime
	}
	if true {
		toSerialize["endTime"] = o.EndTime
	}
	if o.ConfigId.IsSet() {
		toSerialize["configId"] = o.ConfigId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableHeliosBlackoutWindow struct {
	value *HeliosBlackoutWindow
	isSet bool
}

func (v NullableHeliosBlackoutWindow) Get() *HeliosBlackoutWindow {
	return v.value
}

func (v *NullableHeliosBlackoutWindow) Set(val *HeliosBlackoutWindow) {
	v.value = val
	v.isSet = true
}

func (v NullableHeliosBlackoutWindow) IsSet() bool {
	return v.isSet
}

func (v *NullableHeliosBlackoutWindow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHeliosBlackoutWindow(val *HeliosBlackoutWindow) *NullableHeliosBlackoutWindow {
	return &NullableHeliosBlackoutWindow{value: val, isSet: true}
}

func (v NullableHeliosBlackoutWindow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHeliosBlackoutWindow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o HeliosBlackoutWindow) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}