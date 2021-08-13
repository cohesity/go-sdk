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

// CommonArchivalExternalTargetParams Specifies the common parameters which are specific to Archival purpose type External Targets.
type CommonArchivalExternalTargetParams struct {
	// Specifies the Storage type of the External Target.
	StorageType NullableString `json:"storageType"`
	Encryption EncryptionSettings `json:"encryption"`
	TargetBandwidthThrottlings *TargetBandwidthThrottlings `json:"targetBandwidthThrottlings,omitempty"`
}

// NewCommonArchivalExternalTargetParams instantiates a new CommonArchivalExternalTargetParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonArchivalExternalTargetParams(storageType NullableString, encryption EncryptionSettings) *CommonArchivalExternalTargetParams {
	this := CommonArchivalExternalTargetParams{}
	this.StorageType = storageType
	this.Encryption = encryption
	return &this
}

// NewCommonArchivalExternalTargetParamsWithDefaults instantiates a new CommonArchivalExternalTargetParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonArchivalExternalTargetParamsWithDefaults() *CommonArchivalExternalTargetParams {
	this := CommonArchivalExternalTargetParams{}
	return &this
}

// GetStorageType returns the StorageType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CommonArchivalExternalTargetParams) GetStorageType() string {
	if o == nil || o.StorageType.Get() == nil {
		var ret string
		return ret
	}

	return *o.StorageType.Get()
}

// GetStorageTypeOk returns a tuple with the StorageType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonArchivalExternalTargetParams) GetStorageTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StorageType.Get(), o.StorageType.IsSet()
}

// SetStorageType sets field value
func (o *CommonArchivalExternalTargetParams) SetStorageType(v string) {
	o.StorageType.Set(&v)
}

// GetEncryption returns the Encryption field value
func (o *CommonArchivalExternalTargetParams) GetEncryption() EncryptionSettings {
	if o == nil {
		var ret EncryptionSettings
		return ret
	}

	return o.Encryption
}

// GetEncryptionOk returns a tuple with the Encryption field value
// and a boolean to check if the value has been set.
func (o *CommonArchivalExternalTargetParams) GetEncryptionOk() (*EncryptionSettings, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Encryption, true
}

// SetEncryption sets field value
func (o *CommonArchivalExternalTargetParams) SetEncryption(v EncryptionSettings) {
	o.Encryption = v
}

// GetTargetBandwidthThrottlings returns the TargetBandwidthThrottlings field value if set, zero value otherwise.
func (o *CommonArchivalExternalTargetParams) GetTargetBandwidthThrottlings() TargetBandwidthThrottlings {
	if o == nil || o.TargetBandwidthThrottlings == nil {
		var ret TargetBandwidthThrottlings
		return ret
	}
	return *o.TargetBandwidthThrottlings
}

// GetTargetBandwidthThrottlingsOk returns a tuple with the TargetBandwidthThrottlings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonArchivalExternalTargetParams) GetTargetBandwidthThrottlingsOk() (*TargetBandwidthThrottlings, bool) {
	if o == nil || o.TargetBandwidthThrottlings == nil {
		return nil, false
	}
	return o.TargetBandwidthThrottlings, true
}

// HasTargetBandwidthThrottlings returns a boolean if a field has been set.
func (o *CommonArchivalExternalTargetParams) HasTargetBandwidthThrottlings() bool {
	if o != nil && o.TargetBandwidthThrottlings != nil {
		return true
	}

	return false
}

// SetTargetBandwidthThrottlings gets a reference to the given TargetBandwidthThrottlings and assigns it to the TargetBandwidthThrottlings field.
func (o *CommonArchivalExternalTargetParams) SetTargetBandwidthThrottlings(v TargetBandwidthThrottlings) {
	o.TargetBandwidthThrottlings = &v
}

func (o CommonArchivalExternalTargetParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["storageType"] = o.StorageType.Get()
	}
	if true {
		toSerialize["encryption"] = o.Encryption
	}
	if o.TargetBandwidthThrottlings != nil {
		toSerialize["targetBandwidthThrottlings"] = o.TargetBandwidthThrottlings
	}
	return json.Marshal(toSerialize)
}

type NullableCommonArchivalExternalTargetParams struct {
	value *CommonArchivalExternalTargetParams
	isSet bool
}

func (v NullableCommonArchivalExternalTargetParams) Get() *CommonArchivalExternalTargetParams {
	return v.value
}

func (v *NullableCommonArchivalExternalTargetParams) Set(val *CommonArchivalExternalTargetParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonArchivalExternalTargetParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonArchivalExternalTargetParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonArchivalExternalTargetParams(val *CommonArchivalExternalTargetParams) *NullableCommonArchivalExternalTargetParams {
	return &NullableCommonArchivalExternalTargetParams{value: val, isSet: true}
}

func (v NullableCommonArchivalExternalTargetParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonArchivalExternalTargetParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o CommonArchivalExternalTargetParams) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}