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

// PhysicalObjectEntityParams Specifies the common parameters for physical objects.
type PhysicalObjectEntityParams struct {
	// Specifies if system backup was enabled for the source in a particular run.
	EnableSystemBackup NullableBool `json:"enableSystemBackup,omitempty"`
}

// NewPhysicalObjectEntityParams instantiates a new PhysicalObjectEntityParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhysicalObjectEntityParams() *PhysicalObjectEntityParams {
	this := PhysicalObjectEntityParams{}
	return &this
}

// NewPhysicalObjectEntityParamsWithDefaults instantiates a new PhysicalObjectEntityParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhysicalObjectEntityParamsWithDefaults() *PhysicalObjectEntityParams {
	this := PhysicalObjectEntityParams{}
	return &this
}

// GetEnableSystemBackup returns the EnableSystemBackup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PhysicalObjectEntityParams) GetEnableSystemBackup() bool {
	if o == nil || o.EnableSystemBackup.Get() == nil {
		var ret bool
		return ret
	}
	return *o.EnableSystemBackup.Get()
}

// GetEnableSystemBackupOk returns a tuple with the EnableSystemBackup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PhysicalObjectEntityParams) GetEnableSystemBackupOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EnableSystemBackup.Get(), o.EnableSystemBackup.IsSet()
}

// HasEnableSystemBackup returns a boolean if a field has been set.
func (o *PhysicalObjectEntityParams) HasEnableSystemBackup() bool {
	if o != nil && o.EnableSystemBackup.IsSet() {
		return true
	}

	return false
}

// SetEnableSystemBackup gets a reference to the given NullableBool and assigns it to the EnableSystemBackup field.
func (o *PhysicalObjectEntityParams) SetEnableSystemBackup(v bool) {
	o.EnableSystemBackup.Set(&v)
}
// SetEnableSystemBackupNil sets the value for EnableSystemBackup to be an explicit nil
func (o *PhysicalObjectEntityParams) SetEnableSystemBackupNil() {
	o.EnableSystemBackup.Set(nil)
}

// UnsetEnableSystemBackup ensures that no value is present for EnableSystemBackup, not even an explicit nil
func (o *PhysicalObjectEntityParams) UnsetEnableSystemBackup() {
	o.EnableSystemBackup.Unset()
}

func (o PhysicalObjectEntityParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EnableSystemBackup.IsSet() {
		toSerialize["enableSystemBackup"] = o.EnableSystemBackup.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePhysicalObjectEntityParams struct {
	value *PhysicalObjectEntityParams
	isSet bool
}

func (v NullablePhysicalObjectEntityParams) Get() *PhysicalObjectEntityParams {
	return v.value
}

func (v *NullablePhysicalObjectEntityParams) Set(val *PhysicalObjectEntityParams) {
	v.value = val
	v.isSet = true
}

func (v NullablePhysicalObjectEntityParams) IsSet() bool {
	return v.isSet
}

func (v *NullablePhysicalObjectEntityParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhysicalObjectEntityParams(val *PhysicalObjectEntityParams) *NullablePhysicalObjectEntityParams {
	return &NullablePhysicalObjectEntityParams{value: val, isSet: true}
}

func (v NullablePhysicalObjectEntityParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhysicalObjectEntityParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o PhysicalObjectEntityParams) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}