/*
Cohesity REST API

Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CommonArchivalAzureExternalTargetParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonArchivalAzureExternalTargetParams{}

// CommonArchivalAzureExternalTargetParams Specifies the common parameters which are specific to Azure related External Targets of archival purpose type.
type CommonArchivalAzureExternalTargetParams struct {
	// Specifies the client id of the managed identity assigned to the cluster This is used only for clusters running as Azure VMs where authentication is done using AD.
	ClientId NullableString `json:"clientId,omitempty"`
	// Specifies the container name of the external target.
	ContainerName NullableString `json:"containerName"`
	// Specifies region of the External Target. This is only populated for FortKnox vaults.
	Region NullableString `json:"region,omitempty"`
	// Specifies the storage access key of the external target.
	StorageAccessKey NullableString `json:"storageAccessKey,omitempty"`
	// Specifies the storage account name of the external target.
	StorageAccountName NullableString `json:"storageAccountName"`
	// Specifies if Forever Incremental Archival setting is enabled or not.
	IsForeverIncrementalArchivalEnabled NullableBool `json:"isForeverIncrementalArchivalEnabled,omitempty"`
	// Specifies if Incremental Archival setting is enabled or not.
	IsIncrementalArchivalEnabled NullableBool `json:"isIncrementalArchivalEnabled,omitempty"`
	// Specifies whether write once read many (WORM) protection is enabled for the Azure container or not.
	IsWormEnabled NullableBool `json:"isWormEnabled,omitempty"`
	// Specifies the Source Side Deduplication setting for the Azure external target
	SourceSideDeduplication NullableBool `json:"sourceSideDeduplication,omitempty"`
	// Specifies the Azure External Target storage class.
	StorageClass NullableString `json:"storageClass"`
	WormSpecificTargetParams *WormSpecificTargetParams `json:"wormSpecificTargetParams,omitempty"`
}

type _CommonArchivalAzureExternalTargetParams CommonArchivalAzureExternalTargetParams

// NewCommonArchivalAzureExternalTargetParams instantiates a new CommonArchivalAzureExternalTargetParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonArchivalAzureExternalTargetParams(containerName NullableString, storageAccountName NullableString, storageClass NullableString) *CommonArchivalAzureExternalTargetParams {
	this := CommonArchivalAzureExternalTargetParams{}
	this.ContainerName = containerName
	this.StorageAccountName = storageAccountName
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

// GetClientId returns the ClientId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonArchivalAzureExternalTargetParams) GetClientId() string {
	if o == nil || IsNil(o.ClientId.Get()) {
		var ret string
		return ret
	}
	return *o.ClientId.Get()
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonArchivalAzureExternalTargetParams) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientId.Get(), o.ClientId.IsSet()
}

// HasClientId returns a boolean if a field has been set.
func (o *CommonArchivalAzureExternalTargetParams) HasClientId() bool {
	if o != nil && o.ClientId.IsSet() {
		return true
	}

	return false
}

// SetClientId gets a reference to the given NullableString and assigns it to the ClientId field.
func (o *CommonArchivalAzureExternalTargetParams) SetClientId(v string) {
	o.ClientId.Set(&v)
}
// SetClientIdNil sets the value for ClientId to be an explicit nil
func (o *CommonArchivalAzureExternalTargetParams) SetClientIdNil() {
	o.ClientId.Set(nil)
}

// UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
func (o *CommonArchivalAzureExternalTargetParams) UnsetClientId() {
	o.ClientId.Unset()
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
	if o == nil {
		return nil, false
	}
	return o.ContainerName.Get(), o.ContainerName.IsSet()
}

// SetContainerName sets field value
func (o *CommonArchivalAzureExternalTargetParams) SetContainerName(v string) {
	o.ContainerName.Set(&v)
}

// GetRegion returns the Region field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonArchivalAzureExternalTargetParams) GetRegion() string {
	if o == nil || IsNil(o.Region.Get()) {
		var ret string
		return ret
	}
	return *o.Region.Get()
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonArchivalAzureExternalTargetParams) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Region.Get(), o.Region.IsSet()
}

// HasRegion returns a boolean if a field has been set.
func (o *CommonArchivalAzureExternalTargetParams) HasRegion() bool {
	if o != nil && o.Region.IsSet() {
		return true
	}

	return false
}

// SetRegion gets a reference to the given NullableString and assigns it to the Region field.
func (o *CommonArchivalAzureExternalTargetParams) SetRegion(v string) {
	o.Region.Set(&v)
}
// SetRegionNil sets the value for Region to be an explicit nil
func (o *CommonArchivalAzureExternalTargetParams) SetRegionNil() {
	o.Region.Set(nil)
}

// UnsetRegion ensures that no value is present for Region, not even an explicit nil
func (o *CommonArchivalAzureExternalTargetParams) UnsetRegion() {
	o.Region.Unset()
}

// GetStorageAccessKey returns the StorageAccessKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonArchivalAzureExternalTargetParams) GetStorageAccessKey() string {
	if o == nil || IsNil(o.StorageAccessKey.Get()) {
		var ret string
		return ret
	}
	return *o.StorageAccessKey.Get()
}

// GetStorageAccessKeyOk returns a tuple with the StorageAccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonArchivalAzureExternalTargetParams) GetStorageAccessKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StorageAccessKey.Get(), o.StorageAccessKey.IsSet()
}

// HasStorageAccessKey returns a boolean if a field has been set.
func (o *CommonArchivalAzureExternalTargetParams) HasStorageAccessKey() bool {
	if o != nil && o.StorageAccessKey.IsSet() {
		return true
	}

	return false
}

// SetStorageAccessKey gets a reference to the given NullableString and assigns it to the StorageAccessKey field.
func (o *CommonArchivalAzureExternalTargetParams) SetStorageAccessKey(v string) {
	o.StorageAccessKey.Set(&v)
}
// SetStorageAccessKeyNil sets the value for StorageAccessKey to be an explicit nil
func (o *CommonArchivalAzureExternalTargetParams) SetStorageAccessKeyNil() {
	o.StorageAccessKey.Set(nil)
}

// UnsetStorageAccessKey ensures that no value is present for StorageAccessKey, not even an explicit nil
func (o *CommonArchivalAzureExternalTargetParams) UnsetStorageAccessKey() {
	o.StorageAccessKey.Unset()
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
	if o == nil {
		return nil, false
	}
	return o.StorageAccountName.Get(), o.StorageAccountName.IsSet()
}

// SetStorageAccountName sets field value
func (o *CommonArchivalAzureExternalTargetParams) SetStorageAccountName(v string) {
	o.StorageAccountName.Set(&v)
}

// GetIsForeverIncrementalArchivalEnabled returns the IsForeverIncrementalArchivalEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonArchivalAzureExternalTargetParams) GetIsForeverIncrementalArchivalEnabled() bool {
	if o == nil || IsNil(o.IsForeverIncrementalArchivalEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.IsForeverIncrementalArchivalEnabled.Get()
}

// GetIsForeverIncrementalArchivalEnabledOk returns a tuple with the IsForeverIncrementalArchivalEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonArchivalAzureExternalTargetParams) GetIsForeverIncrementalArchivalEnabledOk() (*bool, bool) {
	if o == nil {
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

// GetIsIncrementalArchivalEnabled returns the IsIncrementalArchivalEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonArchivalAzureExternalTargetParams) GetIsIncrementalArchivalEnabled() bool {
	if o == nil || IsNil(o.IsIncrementalArchivalEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.IsIncrementalArchivalEnabled.Get()
}

// GetIsIncrementalArchivalEnabledOk returns a tuple with the IsIncrementalArchivalEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonArchivalAzureExternalTargetParams) GetIsIncrementalArchivalEnabledOk() (*bool, bool) {
	if o == nil {
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

// GetIsWormEnabled returns the IsWormEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonArchivalAzureExternalTargetParams) GetIsWormEnabled() bool {
	if o == nil || IsNil(o.IsWormEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.IsWormEnabled.Get()
}

// GetIsWormEnabledOk returns a tuple with the IsWormEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonArchivalAzureExternalTargetParams) GetIsWormEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsWormEnabled.Get(), o.IsWormEnabled.IsSet()
}

// HasIsWormEnabled returns a boolean if a field has been set.
func (o *CommonArchivalAzureExternalTargetParams) HasIsWormEnabled() bool {
	if o != nil && o.IsWormEnabled.IsSet() {
		return true
	}

	return false
}

// SetIsWormEnabled gets a reference to the given NullableBool and assigns it to the IsWormEnabled field.
func (o *CommonArchivalAzureExternalTargetParams) SetIsWormEnabled(v bool) {
	o.IsWormEnabled.Set(&v)
}
// SetIsWormEnabledNil sets the value for IsWormEnabled to be an explicit nil
func (o *CommonArchivalAzureExternalTargetParams) SetIsWormEnabledNil() {
	o.IsWormEnabled.Set(nil)
}

// UnsetIsWormEnabled ensures that no value is present for IsWormEnabled, not even an explicit nil
func (o *CommonArchivalAzureExternalTargetParams) UnsetIsWormEnabled() {
	o.IsWormEnabled.Unset()
}

// GetSourceSideDeduplication returns the SourceSideDeduplication field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonArchivalAzureExternalTargetParams) GetSourceSideDeduplication() bool {
	if o == nil || IsNil(o.SourceSideDeduplication.Get()) {
		var ret bool
		return ret
	}
	return *o.SourceSideDeduplication.Get()
}

// GetSourceSideDeduplicationOk returns a tuple with the SourceSideDeduplication field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonArchivalAzureExternalTargetParams) GetSourceSideDeduplicationOk() (*bool, bool) {
	if o == nil {
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
	if o == nil {
		return nil, false
	}
	return o.StorageClass.Get(), o.StorageClass.IsSet()
}

// SetStorageClass sets field value
func (o *CommonArchivalAzureExternalTargetParams) SetStorageClass(v string) {
	o.StorageClass.Set(&v)
}

// GetWormSpecificTargetParams returns the WormSpecificTargetParams field value if set, zero value otherwise.
func (o *CommonArchivalAzureExternalTargetParams) GetWormSpecificTargetParams() WormSpecificTargetParams {
	if o == nil || IsNil(o.WormSpecificTargetParams) {
		var ret WormSpecificTargetParams
		return ret
	}
	return *o.WormSpecificTargetParams
}

// GetWormSpecificTargetParamsOk returns a tuple with the WormSpecificTargetParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonArchivalAzureExternalTargetParams) GetWormSpecificTargetParamsOk() (*WormSpecificTargetParams, bool) {
	if o == nil || IsNil(o.WormSpecificTargetParams) {
		return nil, false
	}
	return o.WormSpecificTargetParams, true
}

// HasWormSpecificTargetParams returns a boolean if a field has been set.
func (o *CommonArchivalAzureExternalTargetParams) HasWormSpecificTargetParams() bool {
	if o != nil && !IsNil(o.WormSpecificTargetParams) {
		return true
	}

	return false
}

// SetWormSpecificTargetParams gets a reference to the given WormSpecificTargetParams and assigns it to the WormSpecificTargetParams field.
func (o *CommonArchivalAzureExternalTargetParams) SetWormSpecificTargetParams(v WormSpecificTargetParams) {
	o.WormSpecificTargetParams = &v
}

func (o CommonArchivalAzureExternalTargetParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonArchivalAzureExternalTargetParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId.IsSet() {
		toSerialize["clientId"] = o.ClientId.Get()
	}
	toSerialize["containerName"] = o.ContainerName.Get()
	if o.Region.IsSet() {
		toSerialize["region"] = o.Region.Get()
	}
	if o.StorageAccessKey.IsSet() {
		toSerialize["storageAccessKey"] = o.StorageAccessKey.Get()
	}
	toSerialize["storageAccountName"] = o.StorageAccountName.Get()
	if o.IsForeverIncrementalArchivalEnabled.IsSet() {
		toSerialize["isForeverIncrementalArchivalEnabled"] = o.IsForeverIncrementalArchivalEnabled.Get()
	}
	if o.IsIncrementalArchivalEnabled.IsSet() {
		toSerialize["isIncrementalArchivalEnabled"] = o.IsIncrementalArchivalEnabled.Get()
	}
	if o.IsWormEnabled.IsSet() {
		toSerialize["isWormEnabled"] = o.IsWormEnabled.Get()
	}
	if o.SourceSideDeduplication.IsSet() {
		toSerialize["sourceSideDeduplication"] = o.SourceSideDeduplication.Get()
	}
	toSerialize["storageClass"] = o.StorageClass.Get()
	if !IsNil(o.WormSpecificTargetParams) {
		toSerialize["wormSpecificTargetParams"] = o.WormSpecificTargetParams
	}
	return toSerialize, nil
}

func (o *CommonArchivalAzureExternalTargetParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"containerName",
		"storageAccountName",
		"storageClass",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCommonArchivalAzureExternalTargetParams := _CommonArchivalAzureExternalTargetParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCommonArchivalAzureExternalTargetParams)

	if err != nil {
		return err
	}

	*o = CommonArchivalAzureExternalTargetParams(varCommonArchivalAzureExternalTargetParams)

	return err
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


