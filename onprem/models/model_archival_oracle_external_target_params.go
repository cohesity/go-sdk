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

// ArchivalOracleExternalTargetParams Specifies the common parameters which are specific to Oracle related External Targets.
type ArchivalOracleExternalTargetParams struct {
	// Specifies the bucket name of the external target.
	BucketName NullableString `json:"bucketName"`
	// Specifies the region of the external target.
	Region NullableString `json:"region"`
	// Specifies the access key id of the external target.
	AccessKeyId NullableString `json:"accessKeyId"`
	// Specifies the storage access key of the external target.
	StorageAccessKey NullableString `json:"storageAccessKey"`
	// Specifies the tenancy of the external target.
	Tenancy NullableString `json:"tenancy"`
	// Specifies the Oracle External Target storage class.
	StorageClass NullableString `json:"storageClass"`
	// Specifies the Source Side Deduplication setting for the Oracle external target
	SourceSideDeduplication NullableBool `json:"sourceSideDeduplication,omitempty"`
	// Specifies if Incremental Archival setting is enabled or not.
	IsIncrementalArchivalEnabled NullableBool `json:"isIncrementalArchivalEnabled,omitempty"`
	// Specifies if Forever Incremental Archival setting is enabled or not.
	IsForeverIncrementalArchivalEnabled NullableBool `json:"isForeverIncrementalArchivalEnabled,omitempty"`
}

// NewArchivalOracleExternalTargetParams instantiates a new ArchivalOracleExternalTargetParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArchivalOracleExternalTargetParams(bucketName NullableString, region NullableString, accessKeyId NullableString, storageAccessKey NullableString, tenancy NullableString, storageClass NullableString) *ArchivalOracleExternalTargetParams {
	this := ArchivalOracleExternalTargetParams{}
	this.BucketName = bucketName
	this.Region = region
	this.AccessKeyId = accessKeyId
	this.StorageAccessKey = storageAccessKey
	this.Tenancy = tenancy
	this.StorageClass = storageClass
	return &this
}

// NewArchivalOracleExternalTargetParamsWithDefaults instantiates a new ArchivalOracleExternalTargetParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArchivalOracleExternalTargetParamsWithDefaults() *ArchivalOracleExternalTargetParams {
	this := ArchivalOracleExternalTargetParams{}
	return &this
}

// GetBucketName returns the BucketName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ArchivalOracleExternalTargetParams) GetBucketName() string {
	if o == nil || o.BucketName.Get() == nil {
		var ret string
		return ret
	}

	return *o.BucketName.Get()
}

// GetBucketNameOk returns a tuple with the BucketName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArchivalOracleExternalTargetParams) GetBucketNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BucketName.Get(), o.BucketName.IsSet()
}

// SetBucketName sets field value
func (o *ArchivalOracleExternalTargetParams) SetBucketName(v string) {
	o.BucketName.Set(&v)
}

// GetRegion returns the Region field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ArchivalOracleExternalTargetParams) GetRegion() string {
	if o == nil || o.Region.Get() == nil {
		var ret string
		return ret
	}

	return *o.Region.Get()
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArchivalOracleExternalTargetParams) GetRegionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Region.Get(), o.Region.IsSet()
}

// SetRegion sets field value
func (o *ArchivalOracleExternalTargetParams) SetRegion(v string) {
	o.Region.Set(&v)
}

// GetAccessKeyId returns the AccessKeyId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ArchivalOracleExternalTargetParams) GetAccessKeyId() string {
	if o == nil || o.AccessKeyId.Get() == nil {
		var ret string
		return ret
	}

	return *o.AccessKeyId.Get()
}

// GetAccessKeyIdOk returns a tuple with the AccessKeyId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArchivalOracleExternalTargetParams) GetAccessKeyIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccessKeyId.Get(), o.AccessKeyId.IsSet()
}

// SetAccessKeyId sets field value
func (o *ArchivalOracleExternalTargetParams) SetAccessKeyId(v string) {
	o.AccessKeyId.Set(&v)
}

// GetStorageAccessKey returns the StorageAccessKey field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ArchivalOracleExternalTargetParams) GetStorageAccessKey() string {
	if o == nil || o.StorageAccessKey.Get() == nil {
		var ret string
		return ret
	}

	return *o.StorageAccessKey.Get()
}

// GetStorageAccessKeyOk returns a tuple with the StorageAccessKey field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArchivalOracleExternalTargetParams) GetStorageAccessKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StorageAccessKey.Get(), o.StorageAccessKey.IsSet()
}

// SetStorageAccessKey sets field value
func (o *ArchivalOracleExternalTargetParams) SetStorageAccessKey(v string) {
	o.StorageAccessKey.Set(&v)
}

// GetTenancy returns the Tenancy field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ArchivalOracleExternalTargetParams) GetTenancy() string {
	if o == nil || o.Tenancy.Get() == nil {
		var ret string
		return ret
	}

	return *o.Tenancy.Get()
}

// GetTenancyOk returns a tuple with the Tenancy field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArchivalOracleExternalTargetParams) GetTenancyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Tenancy.Get(), o.Tenancy.IsSet()
}

// SetTenancy sets field value
func (o *ArchivalOracleExternalTargetParams) SetTenancy(v string) {
	o.Tenancy.Set(&v)
}

// GetStorageClass returns the StorageClass field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ArchivalOracleExternalTargetParams) GetStorageClass() string {
	if o == nil || o.StorageClass.Get() == nil {
		var ret string
		return ret
	}

	return *o.StorageClass.Get()
}

// GetStorageClassOk returns a tuple with the StorageClass field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArchivalOracleExternalTargetParams) GetStorageClassOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StorageClass.Get(), o.StorageClass.IsSet()
}

// SetStorageClass sets field value
func (o *ArchivalOracleExternalTargetParams) SetStorageClass(v string) {
	o.StorageClass.Set(&v)
}

// GetSourceSideDeduplication returns the SourceSideDeduplication field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArchivalOracleExternalTargetParams) GetSourceSideDeduplication() bool {
	if o == nil || o.SourceSideDeduplication.Get() == nil {
		var ret bool
		return ret
	}
	return *o.SourceSideDeduplication.Get()
}

// GetSourceSideDeduplicationOk returns a tuple with the SourceSideDeduplication field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArchivalOracleExternalTargetParams) GetSourceSideDeduplicationOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SourceSideDeduplication.Get(), o.SourceSideDeduplication.IsSet()
}

// HasSourceSideDeduplication returns a boolean if a field has been set.
func (o *ArchivalOracleExternalTargetParams) HasSourceSideDeduplication() bool {
	if o != nil && o.SourceSideDeduplication.IsSet() {
		return true
	}

	return false
}

// SetSourceSideDeduplication gets a reference to the given NullableBool and assigns it to the SourceSideDeduplication field.
func (o *ArchivalOracleExternalTargetParams) SetSourceSideDeduplication(v bool) {
	o.SourceSideDeduplication.Set(&v)
}
// SetSourceSideDeduplicationNil sets the value for SourceSideDeduplication to be an explicit nil
func (o *ArchivalOracleExternalTargetParams) SetSourceSideDeduplicationNil() {
	o.SourceSideDeduplication.Set(nil)
}

// UnsetSourceSideDeduplication ensures that no value is present for SourceSideDeduplication, not even an explicit nil
func (o *ArchivalOracleExternalTargetParams) UnsetSourceSideDeduplication() {
	o.SourceSideDeduplication.Unset()
}

// GetIsIncrementalArchivalEnabled returns the IsIncrementalArchivalEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArchivalOracleExternalTargetParams) GetIsIncrementalArchivalEnabled() bool {
	if o == nil || o.IsIncrementalArchivalEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsIncrementalArchivalEnabled.Get()
}

// GetIsIncrementalArchivalEnabledOk returns a tuple with the IsIncrementalArchivalEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArchivalOracleExternalTargetParams) GetIsIncrementalArchivalEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsIncrementalArchivalEnabled.Get(), o.IsIncrementalArchivalEnabled.IsSet()
}

// HasIsIncrementalArchivalEnabled returns a boolean if a field has been set.
func (o *ArchivalOracleExternalTargetParams) HasIsIncrementalArchivalEnabled() bool {
	if o != nil && o.IsIncrementalArchivalEnabled.IsSet() {
		return true
	}

	return false
}

// SetIsIncrementalArchivalEnabled gets a reference to the given NullableBool and assigns it to the IsIncrementalArchivalEnabled field.
func (o *ArchivalOracleExternalTargetParams) SetIsIncrementalArchivalEnabled(v bool) {
	o.IsIncrementalArchivalEnabled.Set(&v)
}
// SetIsIncrementalArchivalEnabledNil sets the value for IsIncrementalArchivalEnabled to be an explicit nil
func (o *ArchivalOracleExternalTargetParams) SetIsIncrementalArchivalEnabledNil() {
	o.IsIncrementalArchivalEnabled.Set(nil)
}

// UnsetIsIncrementalArchivalEnabled ensures that no value is present for IsIncrementalArchivalEnabled, not even an explicit nil
func (o *ArchivalOracleExternalTargetParams) UnsetIsIncrementalArchivalEnabled() {
	o.IsIncrementalArchivalEnabled.Unset()
}

// GetIsForeverIncrementalArchivalEnabled returns the IsForeverIncrementalArchivalEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArchivalOracleExternalTargetParams) GetIsForeverIncrementalArchivalEnabled() bool {
	if o == nil || o.IsForeverIncrementalArchivalEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsForeverIncrementalArchivalEnabled.Get()
}

// GetIsForeverIncrementalArchivalEnabledOk returns a tuple with the IsForeverIncrementalArchivalEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArchivalOracleExternalTargetParams) GetIsForeverIncrementalArchivalEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsForeverIncrementalArchivalEnabled.Get(), o.IsForeverIncrementalArchivalEnabled.IsSet()
}

// HasIsForeverIncrementalArchivalEnabled returns a boolean if a field has been set.
func (o *ArchivalOracleExternalTargetParams) HasIsForeverIncrementalArchivalEnabled() bool {
	if o != nil && o.IsForeverIncrementalArchivalEnabled.IsSet() {
		return true
	}

	return false
}

// SetIsForeverIncrementalArchivalEnabled gets a reference to the given NullableBool and assigns it to the IsForeverIncrementalArchivalEnabled field.
func (o *ArchivalOracleExternalTargetParams) SetIsForeverIncrementalArchivalEnabled(v bool) {
	o.IsForeverIncrementalArchivalEnabled.Set(&v)
}
// SetIsForeverIncrementalArchivalEnabledNil sets the value for IsForeverIncrementalArchivalEnabled to be an explicit nil
func (o *ArchivalOracleExternalTargetParams) SetIsForeverIncrementalArchivalEnabledNil() {
	o.IsForeverIncrementalArchivalEnabled.Set(nil)
}

// UnsetIsForeverIncrementalArchivalEnabled ensures that no value is present for IsForeverIncrementalArchivalEnabled, not even an explicit nil
func (o *ArchivalOracleExternalTargetParams) UnsetIsForeverIncrementalArchivalEnabled() {
	o.IsForeverIncrementalArchivalEnabled.Unset()
}

func (o ArchivalOracleExternalTargetParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bucketName"] = o.BucketName.Get()
	}
	if true {
		toSerialize["region"] = o.Region.Get()
	}
	if true {
		toSerialize["accessKeyId"] = o.AccessKeyId.Get()
	}
	if true {
		toSerialize["storageAccessKey"] = o.StorageAccessKey.Get()
	}
	if true {
		toSerialize["tenancy"] = o.Tenancy.Get()
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

type NullableArchivalOracleExternalTargetParams struct {
	value *ArchivalOracleExternalTargetParams
	isSet bool
}

func (v NullableArchivalOracleExternalTargetParams) Get() *ArchivalOracleExternalTargetParams {
	return v.value
}

func (v *NullableArchivalOracleExternalTargetParams) Set(val *ArchivalOracleExternalTargetParams) {
	v.value = val
	v.isSet = true
}

func (v NullableArchivalOracleExternalTargetParams) IsSet() bool {
	return v.isSet
}

func (v *NullableArchivalOracleExternalTargetParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArchivalOracleExternalTargetParams(val *ArchivalOracleExternalTargetParams) *NullableArchivalOracleExternalTargetParams {
	return &NullableArchivalOracleExternalTargetParams{value: val, isSet: true}
}

func (v NullableArchivalOracleExternalTargetParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArchivalOracleExternalTargetParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o ArchivalOracleExternalTargetParams) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}