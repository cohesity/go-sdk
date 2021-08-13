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

// ScheduledBackup Scheduled Backup type.
type ScheduledBackup struct {
	// Specifies Scheduled Backup type.
	ScheduledBackup *string `json:"scheduledBackup,omitempty"`
}

// NewScheduledBackup instantiates a new ScheduledBackup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduledBackup() *ScheduledBackup {
	this := ScheduledBackup{}
	return &this
}

// NewScheduledBackupWithDefaults instantiates a new ScheduledBackup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduledBackupWithDefaults() *ScheduledBackup {
	this := ScheduledBackup{}
	return &this
}

// GetScheduledBackup returns the ScheduledBackup field value if set, zero value otherwise.
func (o *ScheduledBackup) GetScheduledBackup() string {
	if o == nil || o.ScheduledBackup == nil {
		var ret string
		return ret
	}
	return *o.ScheduledBackup
}

// GetScheduledBackupOk returns a tuple with the ScheduledBackup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledBackup) GetScheduledBackupOk() (*string, bool) {
	if o == nil || o.ScheduledBackup == nil {
		return nil, false
	}
	return o.ScheduledBackup, true
}

// HasScheduledBackup returns a boolean if a field has been set.
func (o *ScheduledBackup) HasScheduledBackup() bool {
	if o != nil && o.ScheduledBackup != nil {
		return true
	}

	return false
}

// SetScheduledBackup gets a reference to the given string and assigns it to the ScheduledBackup field.
func (o *ScheduledBackup) SetScheduledBackup(v string) {
	o.ScheduledBackup = &v
}

func (o ScheduledBackup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ScheduledBackup != nil {
		toSerialize["scheduledBackup"] = o.ScheduledBackup
	}
	return json.Marshal(toSerialize)
}

type NullableScheduledBackup struct {
	value *ScheduledBackup
	isSet bool
}

func (v NullableScheduledBackup) Get() *ScheduledBackup {
	return v.value
}

func (v *NullableScheduledBackup) Set(val *ScheduledBackup) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduledBackup) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduledBackup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduledBackup(val *ScheduledBackup) *NullableScheduledBackup {
	return &NullableScheduledBackup{value: val, isSet: true}
}

func (v NullableScheduledBackup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduledBackup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


