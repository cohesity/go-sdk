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

// VmwareRecoverTargetSourceDiskParams Specifies disk specific parameters for performing a disk recovery.
type VmwareRecoverTargetSourceDiskParams struct {
	// Specifies the UUID of the source disk being recovered.
	DiskUuid NullableString `json:"diskUuid"`
	// Specifies the ID of the datastore on which the specified disk will be spun up.
	DatastoreId NullableInt64 `json:"datastoreId"`
}

// NewVmwareRecoverTargetSourceDiskParams instantiates a new VmwareRecoverTargetSourceDiskParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVmwareRecoverTargetSourceDiskParams(diskUuid NullableString, datastoreId NullableInt64) *VmwareRecoverTargetSourceDiskParams {
	this := VmwareRecoverTargetSourceDiskParams{}
	this.DiskUuid = diskUuid
	this.DatastoreId = datastoreId
	return &this
}

// NewVmwareRecoverTargetSourceDiskParamsWithDefaults instantiates a new VmwareRecoverTargetSourceDiskParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmwareRecoverTargetSourceDiskParamsWithDefaults() *VmwareRecoverTargetSourceDiskParams {
	this := VmwareRecoverTargetSourceDiskParams{}
	return &this
}

// GetDiskUuid returns the DiskUuid field value
// If the value is explicit nil, the zero value for string will be returned
func (o *VmwareRecoverTargetSourceDiskParams) GetDiskUuid() string {
	if o == nil || o.DiskUuid.Get() == nil {
		var ret string
		return ret
	}

	return *o.DiskUuid.Get()
}

// GetDiskUuidOk returns a tuple with the DiskUuid field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmwareRecoverTargetSourceDiskParams) GetDiskUuidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DiskUuid.Get(), o.DiskUuid.IsSet()
}

// SetDiskUuid sets field value
func (o *VmwareRecoverTargetSourceDiskParams) SetDiskUuid(v string) {
	o.DiskUuid.Set(&v)
}

// GetDatastoreId returns the DatastoreId field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *VmwareRecoverTargetSourceDiskParams) GetDatastoreId() int64 {
	if o == nil || o.DatastoreId.Get() == nil {
		var ret int64
		return ret
	}

	return *o.DatastoreId.Get()
}

// GetDatastoreIdOk returns a tuple with the DatastoreId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmwareRecoverTargetSourceDiskParams) GetDatastoreIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DatastoreId.Get(), o.DatastoreId.IsSet()
}

// SetDatastoreId sets field value
func (o *VmwareRecoverTargetSourceDiskParams) SetDatastoreId(v int64) {
	o.DatastoreId.Set(&v)
}

func (o VmwareRecoverTargetSourceDiskParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["diskUuid"] = o.DiskUuid.Get()
	}
	if true {
		toSerialize["datastoreId"] = o.DatastoreId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableVmwareRecoverTargetSourceDiskParams struct {
	value *VmwareRecoverTargetSourceDiskParams
	isSet bool
}

func (v NullableVmwareRecoverTargetSourceDiskParams) Get() *VmwareRecoverTargetSourceDiskParams {
	return v.value
}

func (v *NullableVmwareRecoverTargetSourceDiskParams) Set(val *VmwareRecoverTargetSourceDiskParams) {
	v.value = val
	v.isSet = true
}

func (v NullableVmwareRecoverTargetSourceDiskParams) IsSet() bool {
	return v.isSet
}

func (v *NullableVmwareRecoverTargetSourceDiskParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVmwareRecoverTargetSourceDiskParams(val *VmwareRecoverTargetSourceDiskParams) *NullableVmwareRecoverTargetSourceDiskParams {
	return &NullableVmwareRecoverTargetSourceDiskParams{value: val, isSet: true}
}

func (v NullableVmwareRecoverTargetSourceDiskParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVmwareRecoverTargetSourceDiskParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o VmwareRecoverTargetSourceDiskParams) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}