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

// ArchivalGcpExternalTargetParams Specifies the parameters which are specific to GCP related External Targets of archival purpose type.
type ArchivalGcpExternalTargetParams struct {
	// Specifies the bucket name of the external target.
	BucketName NullableString `json:"bucketName"`
	// Specifies the project Id of the external target.
	ProjectId NullableString `json:"projectId"`
	// Specifies the client email address of the external target.
	ClientEmailAddress NullableString `json:"clientEmailAddress"`
	// Specifies the client private key of the external target.
	ClientPrivateKey NullableString `json:"clientPrivateKey"`
	// Specifies the GCP External Target storage class.
	StorageClass NullableString `json:"storageClass"`
	// Specifies the Source Side Deduplication setting for the GCP external target
	SourceSideDeduplication NullableBool `json:"sourceSideDeduplication,omitempty"`
	// Specifies if Incremental Archival setting is enabled or not.
	IsIncrementalArchivalEnabled NullableBool `json:"isIncrementalArchivalEnabled,omitempty"`
	// Specifies if Forever Incremental Archival setting is enabled or not.
	IsForeverIncrementalArchivalEnabled NullableBool `json:"isForeverIncrementalArchivalEnabled,omitempty"`
}

// NewArchivalGcpExternalTargetParams instantiates a new ArchivalGcpExternalTargetParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArchivalGcpExternalTargetParams(bucketName NullableString, projectId NullableString, clientEmailAddress NullableString, clientPrivateKey NullableString, storageClass NullableString) *ArchivalGcpExternalTargetParams {
	this := ArchivalGcpExternalTargetParams{}
	this.BucketName = bucketName
	this.ProjectId = projectId
	this.ClientEmailAddress = clientEmailAddress
	this.ClientPrivateKey = clientPrivateKey
	this.StorageClass = storageClass
	return &this
}

// NewArchivalGcpExternalTargetParamsWithDefaults instantiates a new ArchivalGcpExternalTargetParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArchivalGcpExternalTargetParamsWithDefaults() *ArchivalGcpExternalTargetParams {
	this := ArchivalGcpExternalTargetParams{}
	return &this
}

// GetBucketName returns the BucketName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ArchivalGcpExternalTargetParams) GetBucketName() string {
	if o == nil || o.BucketName.Get() == nil {
		var ret string
		return ret
	}

	return *o.BucketName.Get()
}

// GetBucketNameOk returns a tuple with the BucketName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArchivalGcpExternalTargetParams) GetBucketNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BucketName.Get(), o.BucketName.IsSet()
}

// SetBucketName sets field value
func (o *ArchivalGcpExternalTargetParams) SetBucketName(v string) {
	o.BucketName.Set(&v)
}

// GetProjectId returns the ProjectId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ArchivalGcpExternalTargetParams) GetProjectId() string {
	if o == nil || o.ProjectId.Get() == nil {
		var ret string
		return ret
	}

	return *o.ProjectId.Get()
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArchivalGcpExternalTargetParams) GetProjectIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ProjectId.Get(), o.ProjectId.IsSet()
}

// SetProjectId sets field value
func (o *ArchivalGcpExternalTargetParams) SetProjectId(v string) {
	o.ProjectId.Set(&v)
}

// GetClientEmailAddress returns the ClientEmailAddress field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ArchivalGcpExternalTargetParams) GetClientEmailAddress() string {
	if o == nil || o.ClientEmailAddress.Get() == nil {
		var ret string
		return ret
	}

	return *o.ClientEmailAddress.Get()
}

// GetClientEmailAddressOk returns a tuple with the ClientEmailAddress field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArchivalGcpExternalTargetParams) GetClientEmailAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ClientEmailAddress.Get(), o.ClientEmailAddress.IsSet()
}

// SetClientEmailAddress sets field value
func (o *ArchivalGcpExternalTargetParams) SetClientEmailAddress(v string) {
	o.ClientEmailAddress.Set(&v)
}

// GetClientPrivateKey returns the ClientPrivateKey field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ArchivalGcpExternalTargetParams) GetClientPrivateKey() string {
	if o == nil || o.ClientPrivateKey.Get() == nil {
		var ret string
		return ret
	}

	return *o.ClientPrivateKey.Get()
}

// GetClientPrivateKeyOk returns a tuple with the ClientPrivateKey field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArchivalGcpExternalTargetParams) GetClientPrivateKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ClientPrivateKey.Get(), o.ClientPrivateKey.IsSet()
}

// SetClientPrivateKey sets field value
func (o *ArchivalGcpExternalTargetParams) SetClientPrivateKey(v string) {
	o.ClientPrivateKey.Set(&v)
}

// GetStorageClass returns the StorageClass field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ArchivalGcpExternalTargetParams) GetStorageClass() string {
	if o == nil || o.StorageClass.Get() == nil {
		var ret string
		return ret
	}

	return *o.StorageClass.Get()
}

// GetStorageClassOk returns a tuple with the StorageClass field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArchivalGcpExternalTargetParams) GetStorageClassOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StorageClass.Get(), o.StorageClass.IsSet()
}

// SetStorageClass sets field value
func (o *ArchivalGcpExternalTargetParams) SetStorageClass(v string) {
	o.StorageClass.Set(&v)
}

// GetSourceSideDeduplication returns the SourceSideDeduplication field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArchivalGcpExternalTargetParams) GetSourceSideDeduplication() bool {
	if o == nil || o.SourceSideDeduplication.Get() == nil {
		var ret bool
		return ret
	}
	return *o.SourceSideDeduplication.Get()
}

// GetSourceSideDeduplicationOk returns a tuple with the SourceSideDeduplication field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArchivalGcpExternalTargetParams) GetSourceSideDeduplicationOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SourceSideDeduplication.Get(), o.SourceSideDeduplication.IsSet()
}

// HasSourceSideDeduplication returns a boolean if a field has been set.
func (o *ArchivalGcpExternalTargetParams) HasSourceSideDeduplication() bool {
	if o != nil && o.SourceSideDeduplication.IsSet() {
		return true
	}

	return false
}

// SetSourceSideDeduplication gets a reference to the given NullableBool and assigns it to the SourceSideDeduplication field.
func (o *ArchivalGcpExternalTargetParams) SetSourceSideDeduplication(v bool) {
	o.SourceSideDeduplication.Set(&v)
}
// SetSourceSideDeduplicationNil sets the value for SourceSideDeduplication to be an explicit nil
func (o *ArchivalGcpExternalTargetParams) SetSourceSideDeduplicationNil() {
	o.SourceSideDeduplication.Set(nil)
}

// UnsetSourceSideDeduplication ensures that no value is present for SourceSideDeduplication, not even an explicit nil
func (o *ArchivalGcpExternalTargetParams) UnsetSourceSideDeduplication() {
	o.SourceSideDeduplication.Unset()
}

// GetIsIncrementalArchivalEnabled returns the IsIncrementalArchivalEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArchivalGcpExternalTargetParams) GetIsIncrementalArchivalEnabled() bool {
	if o == nil || o.IsIncrementalArchivalEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsIncrementalArchivalEnabled.Get()
}

// GetIsIncrementalArchivalEnabledOk returns a tuple with the IsIncrementalArchivalEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArchivalGcpExternalTargetParams) GetIsIncrementalArchivalEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsIncrementalArchivalEnabled.Get(), o.IsIncrementalArchivalEnabled.IsSet()
}

// HasIsIncrementalArchivalEnabled returns a boolean if a field has been set.
func (o *ArchivalGcpExternalTargetParams) HasIsIncrementalArchivalEnabled() bool {
	if o != nil && o.IsIncrementalArchivalEnabled.IsSet() {
		return true
	}

	return false
}

// SetIsIncrementalArchivalEnabled gets a reference to the given NullableBool and assigns it to the IsIncrementalArchivalEnabled field.
func (o *ArchivalGcpExternalTargetParams) SetIsIncrementalArchivalEnabled(v bool) {
	o.IsIncrementalArchivalEnabled.Set(&v)
}
// SetIsIncrementalArchivalEnabledNil sets the value for IsIncrementalArchivalEnabled to be an explicit nil
func (o *ArchivalGcpExternalTargetParams) SetIsIncrementalArchivalEnabledNil() {
	o.IsIncrementalArchivalEnabled.Set(nil)
}

// UnsetIsIncrementalArchivalEnabled ensures that no value is present for IsIncrementalArchivalEnabled, not even an explicit nil
func (o *ArchivalGcpExternalTargetParams) UnsetIsIncrementalArchivalEnabled() {
	o.IsIncrementalArchivalEnabled.Unset()
}

// GetIsForeverIncrementalArchivalEnabled returns the IsForeverIncrementalArchivalEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArchivalGcpExternalTargetParams) GetIsForeverIncrementalArchivalEnabled() bool {
	if o == nil || o.IsForeverIncrementalArchivalEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsForeverIncrementalArchivalEnabled.Get()
}

// GetIsForeverIncrementalArchivalEnabledOk returns a tuple with the IsForeverIncrementalArchivalEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArchivalGcpExternalTargetParams) GetIsForeverIncrementalArchivalEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsForeverIncrementalArchivalEnabled.Get(), o.IsForeverIncrementalArchivalEnabled.IsSet()
}

// HasIsForeverIncrementalArchivalEnabled returns a boolean if a field has been set.
func (o *ArchivalGcpExternalTargetParams) HasIsForeverIncrementalArchivalEnabled() bool {
	if o != nil && o.IsForeverIncrementalArchivalEnabled.IsSet() {
		return true
	}

	return false
}

// SetIsForeverIncrementalArchivalEnabled gets a reference to the given NullableBool and assigns it to the IsForeverIncrementalArchivalEnabled field.
func (o *ArchivalGcpExternalTargetParams) SetIsForeverIncrementalArchivalEnabled(v bool) {
	o.IsForeverIncrementalArchivalEnabled.Set(&v)
}
// SetIsForeverIncrementalArchivalEnabledNil sets the value for IsForeverIncrementalArchivalEnabled to be an explicit nil
func (o *ArchivalGcpExternalTargetParams) SetIsForeverIncrementalArchivalEnabledNil() {
	o.IsForeverIncrementalArchivalEnabled.Set(nil)
}

// UnsetIsForeverIncrementalArchivalEnabled ensures that no value is present for IsForeverIncrementalArchivalEnabled, not even an explicit nil
func (o *ArchivalGcpExternalTargetParams) UnsetIsForeverIncrementalArchivalEnabled() {
	o.IsForeverIncrementalArchivalEnabled.Unset()
}

func (o ArchivalGcpExternalTargetParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bucketName"] = o.BucketName.Get()
	}
	if true {
		toSerialize["projectId"] = o.ProjectId.Get()
	}
	if true {
		toSerialize["clientEmailAddress"] = o.ClientEmailAddress.Get()
	}
	if true {
		toSerialize["clientPrivateKey"] = o.ClientPrivateKey.Get()
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

type NullableArchivalGcpExternalTargetParams struct {
	value *ArchivalGcpExternalTargetParams
	isSet bool
}

func (v NullableArchivalGcpExternalTargetParams) Get() *ArchivalGcpExternalTargetParams {
	return v.value
}

func (v *NullableArchivalGcpExternalTargetParams) Set(val *ArchivalGcpExternalTargetParams) {
	v.value = val
	v.isSet = true
}

func (v NullableArchivalGcpExternalTargetParams) IsSet() bool {
	return v.isSet
}

func (v *NullableArchivalGcpExternalTargetParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArchivalGcpExternalTargetParams(val *ArchivalGcpExternalTargetParams) *NullableArchivalGcpExternalTargetParams {
	return &NullableArchivalGcpExternalTargetParams{value: val, isSet: true}
}

func (v NullableArchivalGcpExternalTargetParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArchivalGcpExternalTargetParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


