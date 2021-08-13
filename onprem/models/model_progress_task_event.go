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

// ProgressTaskEvent Specifies the details about the various events which are created during the execution of Progress Task.
type ProgressTaskEvent struct {
	// Specifies the log message describing the current event.
	Message NullableString `json:"message,omitempty"`
	// Specifies the time of the event occurance in Unix epoch Timestamp(in microseconds).
	OccuredAtUsecs NullableInt64 `json:"occuredAtUsecs,omitempty"`
}

// NewProgressTaskEvent instantiates a new ProgressTaskEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProgressTaskEvent() *ProgressTaskEvent {
	this := ProgressTaskEvent{}
	return &this
}

// NewProgressTaskEventWithDefaults instantiates a new ProgressTaskEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProgressTaskEventWithDefaults() *ProgressTaskEvent {
	this := ProgressTaskEvent{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProgressTaskEvent) GetMessage() string {
	if o == nil || o.Message.Get() == nil {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProgressTaskEvent) GetMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *ProgressTaskEvent) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *ProgressTaskEvent) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *ProgressTaskEvent) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *ProgressTaskEvent) UnsetMessage() {
	o.Message.Unset()
}

// GetOccuredAtUsecs returns the OccuredAtUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProgressTaskEvent) GetOccuredAtUsecs() int64 {
	if o == nil || o.OccuredAtUsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.OccuredAtUsecs.Get()
}

// GetOccuredAtUsecsOk returns a tuple with the OccuredAtUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProgressTaskEvent) GetOccuredAtUsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OccuredAtUsecs.Get(), o.OccuredAtUsecs.IsSet()
}

// HasOccuredAtUsecs returns a boolean if a field has been set.
func (o *ProgressTaskEvent) HasOccuredAtUsecs() bool {
	if o != nil && o.OccuredAtUsecs.IsSet() {
		return true
	}

	return false
}

// SetOccuredAtUsecs gets a reference to the given NullableInt64 and assigns it to the OccuredAtUsecs field.
func (o *ProgressTaskEvent) SetOccuredAtUsecs(v int64) {
	o.OccuredAtUsecs.Set(&v)
}
// SetOccuredAtUsecsNil sets the value for OccuredAtUsecs to be an explicit nil
func (o *ProgressTaskEvent) SetOccuredAtUsecsNil() {
	o.OccuredAtUsecs.Set(nil)
}

// UnsetOccuredAtUsecs ensures that no value is present for OccuredAtUsecs, not even an explicit nil
func (o *ProgressTaskEvent) UnsetOccuredAtUsecs() {
	o.OccuredAtUsecs.Unset()
}

func (o ProgressTaskEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	if o.OccuredAtUsecs.IsSet() {
		toSerialize["occuredAtUsecs"] = o.OccuredAtUsecs.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableProgressTaskEvent struct {
	value *ProgressTaskEvent
	isSet bool
}

func (v NullableProgressTaskEvent) Get() *ProgressTaskEvent {
	return v.value
}

func (v *NullableProgressTaskEvent) Set(val *ProgressTaskEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableProgressTaskEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableProgressTaskEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProgressTaskEvent(val *ProgressTaskEvent) *NullableProgressTaskEvent {
	return &NullableProgressTaskEvent{value: val, isSet: true}
}

func (v NullableProgressTaskEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProgressTaskEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o ProgressTaskEvent) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}