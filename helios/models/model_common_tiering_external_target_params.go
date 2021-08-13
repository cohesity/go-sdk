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

// CommonTieringExternalTargetParams Specifies the common parameters which are specific to Tiering purpose type External Targets.
type CommonTieringExternalTargetParams struct {
	// Specifies the Storage type of the External Target.
	StorageType NullableString `json:"storageType"`
	// Specifies the type of encryption for the Setting.
	EncryptionLevel NullableString `json:"encryptionLevel"`
}

// NewCommonTieringExternalTargetParams instantiates a new CommonTieringExternalTargetParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonTieringExternalTargetParams(storageType NullableString, encryptionLevel NullableString) *CommonTieringExternalTargetParams {
	this := CommonTieringExternalTargetParams{}
	this.StorageType = storageType
	this.EncryptionLevel = encryptionLevel
	return &this
}

// NewCommonTieringExternalTargetParamsWithDefaults instantiates a new CommonTieringExternalTargetParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonTieringExternalTargetParamsWithDefaults() *CommonTieringExternalTargetParams {
	this := CommonTieringExternalTargetParams{}
	return &this
}

// GetStorageType returns the StorageType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CommonTieringExternalTargetParams) GetStorageType() string {
	if o == nil || o.StorageType.Get() == nil {
		var ret string
		return ret
	}

	return *o.StorageType.Get()
}

// GetStorageTypeOk returns a tuple with the StorageType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonTieringExternalTargetParams) GetStorageTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StorageType.Get(), o.StorageType.IsSet()
}

// SetStorageType sets field value
func (o *CommonTieringExternalTargetParams) SetStorageType(v string) {
	o.StorageType.Set(&v)
}

// GetEncryptionLevel returns the EncryptionLevel field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CommonTieringExternalTargetParams) GetEncryptionLevel() string {
	if o == nil || o.EncryptionLevel.Get() == nil {
		var ret string
		return ret
	}

	return *o.EncryptionLevel.Get()
}

// GetEncryptionLevelOk returns a tuple with the EncryptionLevel field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonTieringExternalTargetParams) GetEncryptionLevelOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EncryptionLevel.Get(), o.EncryptionLevel.IsSet()
}

// SetEncryptionLevel sets field value
func (o *CommonTieringExternalTargetParams) SetEncryptionLevel(v string) {
	o.EncryptionLevel.Set(&v)
}

func (o CommonTieringExternalTargetParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["storageType"] = o.StorageType.Get()
	}
	if true {
		toSerialize["encryptionLevel"] = o.EncryptionLevel.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCommonTieringExternalTargetParams struct {
	value *CommonTieringExternalTargetParams
	isSet bool
}

func (v NullableCommonTieringExternalTargetParams) Get() *CommonTieringExternalTargetParams {
	return v.value
}

func (v *NullableCommonTieringExternalTargetParams) Set(val *CommonTieringExternalTargetParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonTieringExternalTargetParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonTieringExternalTargetParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonTieringExternalTargetParams(val *CommonTieringExternalTargetParams) *NullableCommonTieringExternalTargetParams {
	return &NullableCommonTieringExternalTargetParams{value: val, isSet: true}
}

func (v NullableCommonTieringExternalTargetParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonTieringExternalTargetParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


