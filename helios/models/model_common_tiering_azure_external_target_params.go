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

// CommonTieringAzureExternalTargetParams Specifies the common parameters which are specific to Azure related External Targets of tiering purpose type.
type CommonTieringAzureExternalTargetParams struct {
	// Specifies the container name of the external target.
	ContainerName NullableString `json:"containerName"`
	// Specifies the storage account name of the external target.
	StorageAccountName NullableString `json:"storageAccountName"`
	// Specifies the storage access key of the external target.
	StorageAccessKey NullableString `json:"storageAccessKey"`
	// Specifies the Azure External Target class.
	StorageClass NullableString `json:"storageClass"`
}

// NewCommonTieringAzureExternalTargetParams instantiates a new CommonTieringAzureExternalTargetParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonTieringAzureExternalTargetParams(containerName NullableString, storageAccountName NullableString, storageAccessKey NullableString, storageClass NullableString) *CommonTieringAzureExternalTargetParams {
	this := CommonTieringAzureExternalTargetParams{}
	this.ContainerName = containerName
	this.StorageAccountName = storageAccountName
	this.StorageAccessKey = storageAccessKey
	this.StorageClass = storageClass
	return &this
}

// NewCommonTieringAzureExternalTargetParamsWithDefaults instantiates a new CommonTieringAzureExternalTargetParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonTieringAzureExternalTargetParamsWithDefaults() *CommonTieringAzureExternalTargetParams {
	this := CommonTieringAzureExternalTargetParams{}
	return &this
}

// GetContainerName returns the ContainerName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CommonTieringAzureExternalTargetParams) GetContainerName() string {
	if o == nil || o.ContainerName.Get() == nil {
		var ret string
		return ret
	}

	return *o.ContainerName.Get()
}

// GetContainerNameOk returns a tuple with the ContainerName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonTieringAzureExternalTargetParams) GetContainerNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContainerName.Get(), o.ContainerName.IsSet()
}

// SetContainerName sets field value
func (o *CommonTieringAzureExternalTargetParams) SetContainerName(v string) {
	o.ContainerName.Set(&v)
}

// GetStorageAccountName returns the StorageAccountName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CommonTieringAzureExternalTargetParams) GetStorageAccountName() string {
	if o == nil || o.StorageAccountName.Get() == nil {
		var ret string
		return ret
	}

	return *o.StorageAccountName.Get()
}

// GetStorageAccountNameOk returns a tuple with the StorageAccountName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonTieringAzureExternalTargetParams) GetStorageAccountNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StorageAccountName.Get(), o.StorageAccountName.IsSet()
}

// SetStorageAccountName sets field value
func (o *CommonTieringAzureExternalTargetParams) SetStorageAccountName(v string) {
	o.StorageAccountName.Set(&v)
}

// GetStorageAccessKey returns the StorageAccessKey field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CommonTieringAzureExternalTargetParams) GetStorageAccessKey() string {
	if o == nil || o.StorageAccessKey.Get() == nil {
		var ret string
		return ret
	}

	return *o.StorageAccessKey.Get()
}

// GetStorageAccessKeyOk returns a tuple with the StorageAccessKey field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonTieringAzureExternalTargetParams) GetStorageAccessKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StorageAccessKey.Get(), o.StorageAccessKey.IsSet()
}

// SetStorageAccessKey sets field value
func (o *CommonTieringAzureExternalTargetParams) SetStorageAccessKey(v string) {
	o.StorageAccessKey.Set(&v)
}

// GetStorageClass returns the StorageClass field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CommonTieringAzureExternalTargetParams) GetStorageClass() string {
	if o == nil || o.StorageClass.Get() == nil {
		var ret string
		return ret
	}

	return *o.StorageClass.Get()
}

// GetStorageClassOk returns a tuple with the StorageClass field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonTieringAzureExternalTargetParams) GetStorageClassOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StorageClass.Get(), o.StorageClass.IsSet()
}

// SetStorageClass sets field value
func (o *CommonTieringAzureExternalTargetParams) SetStorageClass(v string) {
	o.StorageClass.Set(&v)
}

func (o CommonTieringAzureExternalTargetParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["containerName"] = o.ContainerName.Get()
	}
	if true {
		toSerialize["storageAccountName"] = o.StorageAccountName.Get()
	}
	if true {
		toSerialize["storageAccessKey"] = o.StorageAccessKey.Get()
	}
	if true {
		toSerialize["storageClass"] = o.StorageClass.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCommonTieringAzureExternalTargetParams struct {
	value *CommonTieringAzureExternalTargetParams
	isSet bool
}

func (v NullableCommonTieringAzureExternalTargetParams) Get() *CommonTieringAzureExternalTargetParams {
	return v.value
}

func (v *NullableCommonTieringAzureExternalTargetParams) Set(val *CommonTieringAzureExternalTargetParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonTieringAzureExternalTargetParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonTieringAzureExternalTargetParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonTieringAzureExternalTargetParams(val *CommonTieringAzureExternalTargetParams) *NullableCommonTieringAzureExternalTargetParams {
	return &NullableCommonTieringAzureExternalTargetParams{value: val, isSet: true}
}

func (v NullableCommonTieringAzureExternalTargetParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonTieringAzureExternalTargetParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


