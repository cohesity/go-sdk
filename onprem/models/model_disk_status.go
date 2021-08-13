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

// DiskStatus Disk Status
type DiskStatus struct {
	// Disk Status
	DiskStatus *string `json:"diskStatus,omitempty"`
}

// NewDiskStatus instantiates a new DiskStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiskStatus() *DiskStatus {
	this := DiskStatus{}
	return &this
}

// NewDiskStatusWithDefaults instantiates a new DiskStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiskStatusWithDefaults() *DiskStatus {
	this := DiskStatus{}
	return &this
}

// GetDiskStatus returns the DiskStatus field value if set, zero value otherwise.
func (o *DiskStatus) GetDiskStatus() string {
	if o == nil || o.DiskStatus == nil {
		var ret string
		return ret
	}
	return *o.DiskStatus
}

// GetDiskStatusOk returns a tuple with the DiskStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskStatus) GetDiskStatusOk() (*string, bool) {
	if o == nil || o.DiskStatus == nil {
		return nil, false
	}
	return o.DiskStatus, true
}

// HasDiskStatus returns a boolean if a field has been set.
func (o *DiskStatus) HasDiskStatus() bool {
	if o != nil && o.DiskStatus != nil {
		return true
	}

	return false
}

// SetDiskStatus gets a reference to the given string and assigns it to the DiskStatus field.
func (o *DiskStatus) SetDiskStatus(v string) {
	o.DiskStatus = &v
}

func (o DiskStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DiskStatus != nil {
		toSerialize["diskStatus"] = o.DiskStatus
	}
	return json.Marshal(toSerialize)
}

type NullableDiskStatus struct {
	value *DiskStatus
	isSet bool
}

func (v NullableDiskStatus) Get() *DiskStatus {
	return v.value
}

func (v *NullableDiskStatus) Set(val *DiskStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableDiskStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableDiskStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiskStatus(val *DiskStatus) *NullableDiskStatus {
	return &NullableDiskStatus{value: val, isSet: true}
}

func (v NullableDiskStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiskStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o DiskStatus) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}