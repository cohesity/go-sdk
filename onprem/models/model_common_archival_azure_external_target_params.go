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

// CommonArchivalAzureExternalTargetParams Specifies the common parameters which are specific to Azure related External Targets of archival purpose type.
type CommonArchivalAzureExternalTargetParams struct {
	// Specifies the container name of the external target.
	ContainerName NullableString `json:"containerName"`
	// Specifies the storage account name of the external target.
	StorageAccountName NullableString `json:"storageAccountName"`
	// Specifies the storage access key of the external target.
	StorageAccessKey NullableString `json:"storageAccessKey"`
	// Specifies the Azure External Target storage class.
	StorageClass NullableString `json:"storageClass"`
	// Specifies the Source Side Deduplication setting for the Azure external target
	SourceSideDeduplication NullableBool `json:"sourceSideDeduplication,omitempty"`
	// Specifies if Incremental Archival setting is enabled or not.
	IsIncrementalArchivalEnabled NullableBool `json:"isIncrementalArchivalEnabled,omitempty"`
	// Specifies if Forever Incremental Archival setting is enabled or not.
	IsForeverIncrementalArchivalEnabled NullableBool `json:"isForeverIncrementalArchivalEnabled,omitempty"`
}

// NewCommonArchivalAzureExternalTargetParams instantiates a new CommonArchivalAzureExternalTargetParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonArchivalAzureExternalTargetParams(containerName NullableString, storageAccountName NullableString, storageAccessKey NullableString, storageClass NullableString) *CommonArchivalAzureExternalTargetParams {
	this := CommonArchivalAzureExternalTargetParams{}
	this.ContainerName = containerName
	this.StorageAccountName = storageAccountName
	this.StorageAccessKey = storageAccessKey
	this.StorageClass = storageClass
	return &this
}

// NewCommonArchivalAzureExternalTargetParamsWithDefaults instantiates a new CommonArchivalAzureExternalTargetParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonArchivalAzureExternalTargetParamsWithDefaults() *CommonArchivalAzureExternalTargetParams {
	this := CommonArchivalAzureExternalTargetParams{}
	return &this
}

// GetContainerName returns the ContainerName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CommonArchivalAzureExternalTargetParams) GetContainerName() string {
	if o == nil || o.ContainerName.Get() == nil {
		var ret string
		return ret
	}

	return *o.ContainerName.Get()
}

// GetContainerNameOk returns a tuple with the ContainerName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonArchivalAzureExternalTargetParams) GetContainerNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContainerName.Get(), o.ContainerName.IsSet()
}

// SetContainerName sets field value
func (o *CommonArchivalAzureExternalTargetParams) SetContainerName(v string) {
	o.ContainerName.Set(&v)
}

// GetStorageAccountName returns the StorageAccountName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CommonArchivalAzureExternalTargetParams) GetStorageAccountName() string {
	if o == nil || o.StorageAccountName.Get() == nil {
		var ret string
		return ret
	}

	return *o.StorageAccountName.Get()
}

// GetStorageAccountNameOk returns a tuple with the StorageAccountName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonArchivalAzureExternalTargetParams) GetStorageAccountNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StorageAccountName.Get(), o.StorageAccountName.IsSet()
}

// SetStorageAccountName sets field value
func (o *CommonArchivalAzureExternalTargetParams) SetStorageAccountName(v string) {
	o.StorageAccountName.Set(&v)
}

// GetStorageAccessKey returns the StorageAccessKey field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CommonArchivalAzureExternalTargetParams) GetStorageAccessKey() string {
	if o == nil || o.StorageAccessKey.Get() == nil {
		var ret string
		return ret
	}

	return *o.StorageAccessKey.Get()
}

// GetStorageAccessKeyOk returns a tuple with the StorageAccessKey field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonArchivalAzureExternalTargetParams) GetStorageAccessKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StorageAccessKey.Get(), o.StorageAccessKey.IsSet()
}

// SetStorageAccessKey sets field value
func (o *CommonArchivalAzureExternalTargetParams) SetStorageAccessKey(v string) {
	o.StorageAccessKey.Set(&v)
}

// GetStorageClass returns the StorageClass field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CommonArchivalAzureExternalTargetParams) GetStorageClass() string {
	if o == nil || o.StorageClass.Get() == nil {
		var ret string
		return ret
	}

	return *o.StorageClass.Get()
}

// GetStorageClassOk returns a tuple with the StorageClass field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonArchivalAzureExternalTargetParams) GetStorageClassOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StorageClass.Get(), o.StorageClass.IsSet()
}

// SetStorageClass sets field value
func (o *CommonArchivalAzureExternalTargetParams) SetStorageClass(v string) {
	o.StorageClass.Set(&v)
}

// GetSourceSideDeduplication returns the SourceSideDeduplication field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonArchivalAzureExternalTargetParams) GetSourceSideDeduplication() bool {
	if o == nil || o.SourceSideDeduplication.Get() == nil {
		var ret bool
		return ret
	}
	return *o.SourceSideDeduplication.Get()
}

// GetSourceSideDeduplicationOk returns a tuple with the SourceSideDeduplication field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonArchivalAzureExternalTargetParams) GetSourceSideDeduplicationOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SourceSideDeduplication.Get(), o.SourceSideDeduplication.IsSet()
}

// HasSourceSideDeduplication returns a boolean if a field has been set.
func (o *CommonArchivalAzureExternalTargetParams) HasSourceSideDeduplication() bool {
	if o != nil && o.SourceSideDeduplication.IsSet() {
		return true
	}

	return false
}

// SetSourceSideDeduplication gets a reference to the given NullableBool and assigns it to the SourceSideDeduplication field.
func (o *CommonArchivalAzureExternalTargetParams) SetSourceSideDeduplication(v bool) {
	o.SourceSideDeduplication.Set(&v)
}
// SetSourceSideDeduplicationNil sets the value for SourceSideDeduplication to be an explicit nil
func (o *CommonArchivalAzureExternalTargetParams) SetSourceSideDeduplicationNil() {
	o.SourceSideDeduplication.Set(nil)
}

// UnsetSourceSideDeduplication ensures that no value is present for SourceSideDeduplication, not even an explicit nil
func (o *CommonArchivalAzureExternalTargetParams) UnsetSourceSideDeduplication() {
	o.SourceSideDeduplication.Unset()
}

// GetIsIncrementalArchivalEnabled returns the IsIncrementalArchivalEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonArchivalAzureExternalTargetParams) GetIsIncrementalArchivalEnabled() bool {
	if o == nil || o.IsIncrementalArchivalEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsIncrementalArchivalEnabled.Get()
}

// GetIsIncrementalArchivalEnabledOk returns a tuple with the IsIncrementalArchivalEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonArchivalAzureExternalTargetParams) GetIsIncrementalArchivalEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsIncrementalArchivalEnabled.Get(), o.IsIncrementalArchivalEnabled.IsSet()
}

// HasIsIncrementalArchivalEnabled returns a boolean if a field has been set.
func (o *CommonArchivalAzureExternalTargetParams) HasIsIncrementalArchivalEnabled() bool {
	if o != nil && o.IsIncrementalArchivalEnabled.IsSet() {
		return true
	}

	return false
}

// SetIsIncrementalArchivalEnabled gets a reference to the given NullableBool and assigns it to the IsIncrementalArchivalEnabled field.
func (o *CommonArchivalAzureExternalTargetParams) SetIsIncrementalArchivalEnabled(v bool) {
	o.IsIncrementalArchivalEnabled.Set(&v)
}
// SetIsIncrementalArchivalEnabledNil sets the value for IsIncrementalArchivalEnabled to be an explicit nil
func (o *CommonArchivalAzureExternalTargetParams) SetIsIncrementalArchivalEnabledNil() {
	o.IsIncrementalArchivalEnabled.Set(nil)
}

// UnsetIsIncrementalArchivalEnabled ensures that no value is present for IsIncrementalArchivalEnabled, not even an explicit nil
func (o *CommonArchivalAzureExternalTargetParams) UnsetIsIncrementalArchivalEnabled() {
	o.IsIncrementalArchivalEnabled.Unset()
}

// GetIsForeverIncrementalArchivalEnabled returns the IsForeverIncrementalArchivalEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonArchivalAzureExternalTargetParams) GetIsForeverIncrementalArchivalEnabled() bool {
	if o == nil || o.IsForeverIncrementalArchivalEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsForeverIncrementalArchivalEnabled.Get()
}

// GetIsForeverIncrementalArchivalEnabledOk returns a tuple with the IsForeverIncrementalArchivalEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonArchivalAzureExternalTargetParams) GetIsForeverIncrementalArchivalEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsForeverIncrementalArchivalEnabled.Get(), o.IsForeverIncrementalArchivalEnabled.IsSet()
}

// HasIsForeverIncrementalArchivalEnabled returns a boolean if a field has been set.
func (o *CommonArchivalAzureExternalTargetParams) HasIsForeverIncrementalArchivalEnabled() bool {
	if o != nil && o.IsForeverIncrementalArchivalEnabled.IsSet() {
		return true
	}

	return false
}

// SetIsForeverIncrementalArchivalEnabled gets a reference to the given NullableBool and assigns it to the IsForeverIncrementalArchivalEnabled field.
func (o *CommonArchivalAzureExternalTargetParams) SetIsForeverIncrementalArchivalEnabled(v bool) {
	o.IsForeverIncrementalArchivalEnabled.Set(&v)
}
// SetIsForeverIncrementalArchivalEnabledNil sets the value for IsForeverIncrementalArchivalEnabled to be an explicit nil
func (o *CommonArchivalAzureExternalTargetParams) SetIsForeverIncrementalArchivalEnabledNil() {
	o.IsForeverIncrementalArchivalEnabled.Set(nil)
}

// UnsetIsForeverIncrementalArchivalEnabled ensures that no value is present for IsForeverIncrementalArchivalEnabled, not even an explicit nil
func (o *CommonArchivalAzureExternalTargetParams) UnsetIsForeverIncrementalArchivalEnabled() {
	o.IsForeverIncrementalArchivalEnabled.Unset()
}

func (o CommonArchivalAzureExternalTargetParams) MarshalJSON() ([]byte, error) {
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
	if o.SourceSideDeduplication.IsSet() {
		toSerialize["sourceSideDeduplication"] = o.SourceSideDeduplication.Get()
	}
	if o.IsIncrementalArchivalEnabled.IsSet() {
		toSerialize["isIncrementalArchivalEnabled"] = o.IsIncrementalArchivalEnabled.Get()
	}
	if o.IsForeverIncrementalArchivalEnabled.IsSet() {
		toSerialize["isForeverIncrementalArchivalEnabled"] = o.IsForeverIncrementalArchivalEnabled.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCommonArchivalAzureExternalTargetParams struct {
	value *CommonArchivalAzureExternalTargetParams
	isSet bool
}

func (v NullableCommonArchivalAzureExternalTargetParams) Get() *CommonArchivalAzureExternalTargetParams {
	return v.value
}

func (v *NullableCommonArchivalAzureExternalTargetParams) Set(val *CommonArchivalAzureExternalTargetParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonArchivalAzureExternalTargetParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonArchivalAzureExternalTargetParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonArchivalAzureExternalTargetParams(val *CommonArchivalAzureExternalTargetParams) *NullableCommonArchivalAzureExternalTargetParams {
	return &NullableCommonArchivalAzureExternalTargetParams{value: val, isSet: true}
}

func (v NullableCommonArchivalAzureExternalTargetParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonArchivalAzureExternalTargetParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o CommonArchivalAzureExternalTargetParams) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}