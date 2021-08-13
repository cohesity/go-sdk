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

// MSSQLProtectionGroupParams Specifies the parameters specific to MSSQL Protection Group.
type MSSQLProtectionGroupParams struct {
	// Specifies the MSSQL Protection Group type.
	ProtectionType string `json:"protectionType"`
	FileProtectionTypeParams *MSSQLFileProtectionGroupParams `json:"fileProtectionTypeParams,omitempty"`
	VolumeProtectionTypeParams *MSSQLVolumeProtectionGroupParams `json:"volumeProtectionTypeParams,omitempty"`
	NativeProtectionTypeParams *MSSQLNativeProtectionGroupParams `json:"nativeProtectionTypeParams,omitempty"`
}

// NewMSSQLProtectionGroupParams instantiates a new MSSQLProtectionGroupParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMSSQLProtectionGroupParams(protectionType string) *MSSQLProtectionGroupParams {
	this := MSSQLProtectionGroupParams{}
	this.ProtectionType = protectionType
	return &this
}

// NewMSSQLProtectionGroupParamsWithDefaults instantiates a new MSSQLProtectionGroupParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMSSQLProtectionGroupParamsWithDefaults() *MSSQLProtectionGroupParams {
	this := MSSQLProtectionGroupParams{}
	return &this
}

// GetProtectionType returns the ProtectionType field value
func (o *MSSQLProtectionGroupParams) GetProtectionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProtectionType
}

// GetProtectionTypeOk returns a tuple with the ProtectionType field value
// and a boolean to check if the value has been set.
func (o *MSSQLProtectionGroupParams) GetProtectionTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProtectionType, true
}

// SetProtectionType sets field value
func (o *MSSQLProtectionGroupParams) SetProtectionType(v string) {
	o.ProtectionType = v
}

// GetFileProtectionTypeParams returns the FileProtectionTypeParams field value if set, zero value otherwise.
func (o *MSSQLProtectionGroupParams) GetFileProtectionTypeParams() MSSQLFileProtectionGroupParams {
	if o == nil || o.FileProtectionTypeParams == nil {
		var ret MSSQLFileProtectionGroupParams
		return ret
	}
	return *o.FileProtectionTypeParams
}

// GetFileProtectionTypeParamsOk returns a tuple with the FileProtectionTypeParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MSSQLProtectionGroupParams) GetFileProtectionTypeParamsOk() (*MSSQLFileProtectionGroupParams, bool) {
	if o == nil || o.FileProtectionTypeParams == nil {
		return nil, false
	}
	return o.FileProtectionTypeParams, true
}

// HasFileProtectionTypeParams returns a boolean if a field has been set.
func (o *MSSQLProtectionGroupParams) HasFileProtectionTypeParams() bool {
	if o != nil && o.FileProtectionTypeParams != nil {
		return true
	}

	return false
}

// SetFileProtectionTypeParams gets a reference to the given MSSQLFileProtectionGroupParams and assigns it to the FileProtectionTypeParams field.
func (o *MSSQLProtectionGroupParams) SetFileProtectionTypeParams(v MSSQLFileProtectionGroupParams) {
	o.FileProtectionTypeParams = &v
}

// GetVolumeProtectionTypeParams returns the VolumeProtectionTypeParams field value if set, zero value otherwise.
func (o *MSSQLProtectionGroupParams) GetVolumeProtectionTypeParams() MSSQLVolumeProtectionGroupParams {
	if o == nil || o.VolumeProtectionTypeParams == nil {
		var ret MSSQLVolumeProtectionGroupParams
		return ret
	}
	return *o.VolumeProtectionTypeParams
}

// GetVolumeProtectionTypeParamsOk returns a tuple with the VolumeProtectionTypeParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MSSQLProtectionGroupParams) GetVolumeProtectionTypeParamsOk() (*MSSQLVolumeProtectionGroupParams, bool) {
	if o == nil || o.VolumeProtectionTypeParams == nil {
		return nil, false
	}
	return o.VolumeProtectionTypeParams, true
}

// HasVolumeProtectionTypeParams returns a boolean if a field has been set.
func (o *MSSQLProtectionGroupParams) HasVolumeProtectionTypeParams() bool {
	if o != nil && o.VolumeProtectionTypeParams != nil {
		return true
	}

	return false
}

// SetVolumeProtectionTypeParams gets a reference to the given MSSQLVolumeProtectionGroupParams and assigns it to the VolumeProtectionTypeParams field.
func (o *MSSQLProtectionGroupParams) SetVolumeProtectionTypeParams(v MSSQLVolumeProtectionGroupParams) {
	o.VolumeProtectionTypeParams = &v
}

// GetNativeProtectionTypeParams returns the NativeProtectionTypeParams field value if set, zero value otherwise.
func (o *MSSQLProtectionGroupParams) GetNativeProtectionTypeParams() MSSQLNativeProtectionGroupParams {
	if o == nil || o.NativeProtectionTypeParams == nil {
		var ret MSSQLNativeProtectionGroupParams
		return ret
	}
	return *o.NativeProtectionTypeParams
}

// GetNativeProtectionTypeParamsOk returns a tuple with the NativeProtectionTypeParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MSSQLProtectionGroupParams) GetNativeProtectionTypeParamsOk() (*MSSQLNativeProtectionGroupParams, bool) {
	if o == nil || o.NativeProtectionTypeParams == nil {
		return nil, false
	}
	return o.NativeProtectionTypeParams, true
}

// HasNativeProtectionTypeParams returns a boolean if a field has been set.
func (o *MSSQLProtectionGroupParams) HasNativeProtectionTypeParams() bool {
	if o != nil && o.NativeProtectionTypeParams != nil {
		return true
	}

	return false
}

// SetNativeProtectionTypeParams gets a reference to the given MSSQLNativeProtectionGroupParams and assigns it to the NativeProtectionTypeParams field.
func (o *MSSQLProtectionGroupParams) SetNativeProtectionTypeParams(v MSSQLNativeProtectionGroupParams) {
	o.NativeProtectionTypeParams = &v
}

func (o MSSQLProtectionGroupParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["protectionType"] = o.ProtectionType
	}
	if o.FileProtectionTypeParams != nil {
		toSerialize["fileProtectionTypeParams"] = o.FileProtectionTypeParams
	}
	if o.VolumeProtectionTypeParams != nil {
		toSerialize["volumeProtectionTypeParams"] = o.VolumeProtectionTypeParams
	}
	if o.NativeProtectionTypeParams != nil {
		toSerialize["nativeProtectionTypeParams"] = o.NativeProtectionTypeParams
	}
	return json.Marshal(toSerialize)
}

type NullableMSSQLProtectionGroupParams struct {
	value *MSSQLProtectionGroupParams
	isSet bool
}

func (v NullableMSSQLProtectionGroupParams) Get() *MSSQLProtectionGroupParams {
	return v.value
}

func (v *NullableMSSQLProtectionGroupParams) Set(val *MSSQLProtectionGroupParams) {
	v.value = val
	v.isSet = true
}

func (v NullableMSSQLProtectionGroupParams) IsSet() bool {
	return v.isSet
}

func (v *NullableMSSQLProtectionGroupParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMSSQLProtectionGroupParams(val *MSSQLProtectionGroupParams) *NullableMSSQLProtectionGroupParams {
	return &NullableMSSQLProtectionGroupParams{value: val, isSet: true}
}

func (v NullableMSSQLProtectionGroupParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMSSQLProtectionGroupParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o MSSQLProtectionGroupParams) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}