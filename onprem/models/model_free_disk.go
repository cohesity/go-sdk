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

// FreeDisk Specifies the details of a free disk.
type FreeDisk struct {
	// Specifies the location of disk.
	Location NullableString `json:"location,omitempty"`
	// Specifies serial number of disk.
	SerialNumber NullableString `json:"serialNumber"`
	// Specifies path of disk.
	Path NullableString `json:"path,omitempty"`
	// Size of disk.
	SizeInBytes NullableInt64 `json:"sizeInBytes,omitempty"`
}

// NewFreeDisk instantiates a new FreeDisk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFreeDisk(serialNumber NullableString) *FreeDisk {
	this := FreeDisk{}
	this.SerialNumber = serialNumber
	return &this
}

// NewFreeDiskWithDefaults instantiates a new FreeDisk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFreeDiskWithDefaults() *FreeDisk {
	this := FreeDisk{}
	return &this
}

// GetLocation returns the Location field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FreeDisk) GetLocation() string {
	if o == nil || o.Location.Get() == nil {
		var ret string
		return ret
	}
	return *o.Location.Get()
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FreeDisk) GetLocationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Location.Get(), o.Location.IsSet()
}

// HasLocation returns a boolean if a field has been set.
func (o *FreeDisk) HasLocation() bool {
	if o != nil && o.Location.IsSet() {
		return true
	}

	return false
}

// SetLocation gets a reference to the given NullableString and assigns it to the Location field.
func (o *FreeDisk) SetLocation(v string) {
	o.Location.Set(&v)
}
// SetLocationNil sets the value for Location to be an explicit nil
func (o *FreeDisk) SetLocationNil() {
	o.Location.Set(nil)
}

// UnsetLocation ensures that no value is present for Location, not even an explicit nil
func (o *FreeDisk) UnsetLocation() {
	o.Location.Unset()
}

// GetSerialNumber returns the SerialNumber field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FreeDisk) GetSerialNumber() string {
	if o == nil || o.SerialNumber.Get() == nil {
		var ret string
		return ret
	}

	return *o.SerialNumber.Get()
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FreeDisk) GetSerialNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SerialNumber.Get(), o.SerialNumber.IsSet()
}

// SetSerialNumber sets field value
func (o *FreeDisk) SetSerialNumber(v string) {
	o.SerialNumber.Set(&v)
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FreeDisk) GetPath() string {
	if o == nil || o.Path.Get() == nil {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FreeDisk) GetPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *FreeDisk) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *FreeDisk) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *FreeDisk) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *FreeDisk) UnsetPath() {
	o.Path.Unset()
}

// GetSizeInBytes returns the SizeInBytes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FreeDisk) GetSizeInBytes() int64 {
	if o == nil || o.SizeInBytes.Get() == nil {
		var ret int64
		return ret
	}
	return *o.SizeInBytes.Get()
}

// GetSizeInBytesOk returns a tuple with the SizeInBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FreeDisk) GetSizeInBytesOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SizeInBytes.Get(), o.SizeInBytes.IsSet()
}

// HasSizeInBytes returns a boolean if a field has been set.
func (o *FreeDisk) HasSizeInBytes() bool {
	if o != nil && o.SizeInBytes.IsSet() {
		return true
	}

	return false
}

// SetSizeInBytes gets a reference to the given NullableInt64 and assigns it to the SizeInBytes field.
func (o *FreeDisk) SetSizeInBytes(v int64) {
	o.SizeInBytes.Set(&v)
}
// SetSizeInBytesNil sets the value for SizeInBytes to be an explicit nil
func (o *FreeDisk) SetSizeInBytesNil() {
	o.SizeInBytes.Set(nil)
}

// UnsetSizeInBytes ensures that no value is present for SizeInBytes, not even an explicit nil
func (o *FreeDisk) UnsetSizeInBytes() {
	o.SizeInBytes.Unset()
}

func (o FreeDisk) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Location.IsSet() {
		toSerialize["location"] = o.Location.Get()
	}
	if true {
		toSerialize["serialNumber"] = o.SerialNumber.Get()
	}
	if o.Path.IsSet() {
		toSerialize["path"] = o.Path.Get()
	}
	if o.SizeInBytes.IsSet() {
		toSerialize["sizeInBytes"] = o.SizeInBytes.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableFreeDisk struct {
	value *FreeDisk
	isSet bool
}

func (v NullableFreeDisk) Get() *FreeDisk {
	return v.value
}

func (v *NullableFreeDisk) Set(val *FreeDisk) {
	v.value = val
	v.isSet = true
}

func (v NullableFreeDisk) IsSet() bool {
	return v.isSet
}

func (v *NullableFreeDisk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFreeDisk(val *FreeDisk) *NullableFreeDisk {
	return &NullableFreeDisk{value: val, isSet: true}
}

func (v NullableFreeDisk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFreeDisk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o FreeDisk) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}