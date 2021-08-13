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

// ArchivalTargetTierInfo Specifies the tier info for archival.
type ArchivalTargetTierInfo struct {
	// Specifies the cloud platform to enable tiering.
	CloudPlatform NullableString `json:"cloudPlatform"`
	AwsTiering *AWSTiers `json:"awsTiering,omitempty"`
	AzureTiering *AzureTiers `json:"azureTiering,omitempty"`
	GoogleTiering *GoogleTiers `json:"googleTiering,omitempty"`
	OracleTiering *OracleTiers `json:"oracleTiering,omitempty"`
	// Specifies the type of the current tier where the snapshot resides. This will be specified if the run is a CAD run.
	CurrentTierType NullableString `json:"currentTierType,omitempty"`
}

// NewArchivalTargetTierInfo instantiates a new ArchivalTargetTierInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArchivalTargetTierInfo(cloudPlatform NullableString) *ArchivalTargetTierInfo {
	this := ArchivalTargetTierInfo{}
	this.CloudPlatform = cloudPlatform
	return &this
}

// NewArchivalTargetTierInfoWithDefaults instantiates a new ArchivalTargetTierInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArchivalTargetTierInfoWithDefaults() *ArchivalTargetTierInfo {
	this := ArchivalTargetTierInfo{}
	return &this
}

// GetCloudPlatform returns the CloudPlatform field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ArchivalTargetTierInfo) GetCloudPlatform() string {
	if o == nil || o.CloudPlatform.Get() == nil {
		var ret string
		return ret
	}

	return *o.CloudPlatform.Get()
}

// GetCloudPlatformOk returns a tuple with the CloudPlatform field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArchivalTargetTierInfo) GetCloudPlatformOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CloudPlatform.Get(), o.CloudPlatform.IsSet()
}

// SetCloudPlatform sets field value
func (o *ArchivalTargetTierInfo) SetCloudPlatform(v string) {
	o.CloudPlatform.Set(&v)
}

// GetAwsTiering returns the AwsTiering field value if set, zero value otherwise.
func (o *ArchivalTargetTierInfo) GetAwsTiering() AWSTiers {
	if o == nil || o.AwsTiering == nil {
		var ret AWSTiers
		return ret
	}
	return *o.AwsTiering
}

// GetAwsTieringOk returns a tuple with the AwsTiering field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchivalTargetTierInfo) GetAwsTieringOk() (*AWSTiers, bool) {
	if o == nil || o.AwsTiering == nil {
		return nil, false
	}
	return o.AwsTiering, true
}

// HasAwsTiering returns a boolean if a field has been set.
func (o *ArchivalTargetTierInfo) HasAwsTiering() bool {
	if o != nil && o.AwsTiering != nil {
		return true
	}

	return false
}

// SetAwsTiering gets a reference to the given AWSTiers and assigns it to the AwsTiering field.
func (o *ArchivalTargetTierInfo) SetAwsTiering(v AWSTiers) {
	o.AwsTiering = &v
}

// GetAzureTiering returns the AzureTiering field value if set, zero value otherwise.
func (o *ArchivalTargetTierInfo) GetAzureTiering() AzureTiers {
	if o == nil || o.AzureTiering == nil {
		var ret AzureTiers
		return ret
	}
	return *o.AzureTiering
}

// GetAzureTieringOk returns a tuple with the AzureTiering field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchivalTargetTierInfo) GetAzureTieringOk() (*AzureTiers, bool) {
	if o == nil || o.AzureTiering == nil {
		return nil, false
	}
	return o.AzureTiering, true
}

// HasAzureTiering returns a boolean if a field has been set.
func (o *ArchivalTargetTierInfo) HasAzureTiering() bool {
	if o != nil && o.AzureTiering != nil {
		return true
	}

	return false
}

// SetAzureTiering gets a reference to the given AzureTiers and assigns it to the AzureTiering field.
func (o *ArchivalTargetTierInfo) SetAzureTiering(v AzureTiers) {
	o.AzureTiering = &v
}

// GetGoogleTiering returns the GoogleTiering field value if set, zero value otherwise.
func (o *ArchivalTargetTierInfo) GetGoogleTiering() GoogleTiers {
	if o == nil || o.GoogleTiering == nil {
		var ret GoogleTiers
		return ret
	}
	return *o.GoogleTiering
}

// GetGoogleTieringOk returns a tuple with the GoogleTiering field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchivalTargetTierInfo) GetGoogleTieringOk() (*GoogleTiers, bool) {
	if o == nil || o.GoogleTiering == nil {
		return nil, false
	}
	return o.GoogleTiering, true
}

// HasGoogleTiering returns a boolean if a field has been set.
func (o *ArchivalTargetTierInfo) HasGoogleTiering() bool {
	if o != nil && o.GoogleTiering != nil {
		return true
	}

	return false
}

// SetGoogleTiering gets a reference to the given GoogleTiers and assigns it to the GoogleTiering field.
func (o *ArchivalTargetTierInfo) SetGoogleTiering(v GoogleTiers) {
	o.GoogleTiering = &v
}

// GetOracleTiering returns the OracleTiering field value if set, zero value otherwise.
func (o *ArchivalTargetTierInfo) GetOracleTiering() OracleTiers {
	if o == nil || o.OracleTiering == nil {
		var ret OracleTiers
		return ret
	}
	return *o.OracleTiering
}

// GetOracleTieringOk returns a tuple with the OracleTiering field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchivalTargetTierInfo) GetOracleTieringOk() (*OracleTiers, bool) {
	if o == nil || o.OracleTiering == nil {
		return nil, false
	}
	return o.OracleTiering, true
}

// HasOracleTiering returns a boolean if a field has been set.
func (o *ArchivalTargetTierInfo) HasOracleTiering() bool {
	if o != nil && o.OracleTiering != nil {
		return true
	}

	return false
}

// SetOracleTiering gets a reference to the given OracleTiers and assigns it to the OracleTiering field.
func (o *ArchivalTargetTierInfo) SetOracleTiering(v OracleTiers) {
	o.OracleTiering = &v
}

// GetCurrentTierType returns the CurrentTierType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArchivalTargetTierInfo) GetCurrentTierType() string {
	if o == nil || o.CurrentTierType.Get() == nil {
		var ret string
		return ret
	}
	return *o.CurrentTierType.Get()
}

// GetCurrentTierTypeOk returns a tuple with the CurrentTierType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArchivalTargetTierInfo) GetCurrentTierTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CurrentTierType.Get(), o.CurrentTierType.IsSet()
}

// HasCurrentTierType returns a boolean if a field has been set.
func (o *ArchivalTargetTierInfo) HasCurrentTierType() bool {
	if o != nil && o.CurrentTierType.IsSet() {
		return true
	}

	return false
}

// SetCurrentTierType gets a reference to the given NullableString and assigns it to the CurrentTierType field.
func (o *ArchivalTargetTierInfo) SetCurrentTierType(v string) {
	o.CurrentTierType.Set(&v)
}
// SetCurrentTierTypeNil sets the value for CurrentTierType to be an explicit nil
func (o *ArchivalTargetTierInfo) SetCurrentTierTypeNil() {
	o.CurrentTierType.Set(nil)
}

// UnsetCurrentTierType ensures that no value is present for CurrentTierType, not even an explicit nil
func (o *ArchivalTargetTierInfo) UnsetCurrentTierType() {
	o.CurrentTierType.Unset()
}

func (o ArchivalTargetTierInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cloudPlatform"] = o.CloudPlatform.Get()
	}
	if o.AwsTiering != nil {
		toSerialize["awsTiering"] = o.AwsTiering
	}
	if o.AzureTiering != nil {
		toSerialize["azureTiering"] = o.AzureTiering
	}
	if o.GoogleTiering != nil {
		toSerialize["googleTiering"] = o.GoogleTiering
	}
	if o.OracleTiering != nil {
		toSerialize["oracleTiering"] = o.OracleTiering
	}
	if o.CurrentTierType.IsSet() {
		toSerialize["currentTierType"] = o.CurrentTierType.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableArchivalTargetTierInfo struct {
	value *ArchivalTargetTierInfo
	isSet bool
}

func (v NullableArchivalTargetTierInfo) Get() *ArchivalTargetTierInfo {
	return v.value
}

func (v *NullableArchivalTargetTierInfo) Set(val *ArchivalTargetTierInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableArchivalTargetTierInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableArchivalTargetTierInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArchivalTargetTierInfo(val *ArchivalTargetTierInfo) *NullableArchivalTargetTierInfo {
	return &NullableArchivalTargetTierInfo{value: val, isSet: true}
}

func (v NullableArchivalTargetTierInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArchivalTargetTierInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


