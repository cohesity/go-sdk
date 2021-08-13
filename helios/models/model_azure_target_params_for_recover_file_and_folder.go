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

// AzureTargetParamsForRecoverFileAndFolder Specifies the parameters for an Azure files and folders recovery target.
type AzureTargetParamsForRecoverFileAndFolder struct {
	// Specifies whether to recover to the original target. If true, originalTargetConfig must be specified. If false, newTargetConfig must be specified.
	RecoverToOriginalTarget NullableBool `json:"recoverToOriginalTarget"`
	// Specifies the configuration for recovering to the original target.
	OriginalTargetConfig NullableAzureRecoverFilesOriginalTargetConfig `json:"originalTargetConfig,omitempty"`
	// Specifies the configuration for recovering to a new target.
	NewTargetConfig NullableAzureRecoverFilesNewTargetConfig `json:"newTargetConfig,omitempty"`
	// Specifies whether to overwrite the existing files. Default is true.
	OverwriteExisting NullableBool `json:"overwriteExisting,omitempty"`
	// Specifies whether to preserve original file/folder attributes. Default is true.
	PreserveAttributes NullableBool `json:"preserveAttributes,omitempty"`
	// Specifies whether to continue recovering other files if one of the objects encounters an error. Default is false.
	ContinueOnError NullableBool `json:"continueOnError,omitempty"`
	// Specifies VLAN Params associated with the recovered files and folders. If this is not specified, then the VLAN settings will be automatically selected from one of the below options: a. If VLANs are configured on Cohesity, then the VLAN host/VIP will be automatically based on the client's (e.g. ESXI host) IP address. b. If VLANs are not configured on Cohesity, then the partition hostname or VIPs will be used for Recovery.
	VlanConfig NullableRecoveryVlanConfig `json:"vlanConfig,omitempty"`
}

// NewAzureTargetParamsForRecoverFileAndFolder instantiates a new AzureTargetParamsForRecoverFileAndFolder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureTargetParamsForRecoverFileAndFolder(recoverToOriginalTarget NullableBool) *AzureTargetParamsForRecoverFileAndFolder {
	this := AzureTargetParamsForRecoverFileAndFolder{}
	this.RecoverToOriginalTarget = recoverToOriginalTarget
	return &this
}

// NewAzureTargetParamsForRecoverFileAndFolderWithDefaults instantiates a new AzureTargetParamsForRecoverFileAndFolder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureTargetParamsForRecoverFileAndFolderWithDefaults() *AzureTargetParamsForRecoverFileAndFolder {
	this := AzureTargetParamsForRecoverFileAndFolder{}
	return &this
}

// GetRecoverToOriginalTarget returns the RecoverToOriginalTarget field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *AzureTargetParamsForRecoverFileAndFolder) GetRecoverToOriginalTarget() bool {
	if o == nil || o.RecoverToOriginalTarget.Get() == nil {
		var ret bool
		return ret
	}

	return *o.RecoverToOriginalTarget.Get()
}

// GetRecoverToOriginalTargetOk returns a tuple with the RecoverToOriginalTarget field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureTargetParamsForRecoverFileAndFolder) GetRecoverToOriginalTargetOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RecoverToOriginalTarget.Get(), o.RecoverToOriginalTarget.IsSet()
}

// SetRecoverToOriginalTarget sets field value
func (o *AzureTargetParamsForRecoverFileAndFolder) SetRecoverToOriginalTarget(v bool) {
	o.RecoverToOriginalTarget.Set(&v)
}

// GetOriginalTargetConfig returns the OriginalTargetConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureTargetParamsForRecoverFileAndFolder) GetOriginalTargetConfig() AzureRecoverFilesOriginalTargetConfig {
	if o == nil || o.OriginalTargetConfig.Get() == nil {
		var ret AzureRecoverFilesOriginalTargetConfig
		return ret
	}
	return *o.OriginalTargetConfig.Get()
}

// GetOriginalTargetConfigOk returns a tuple with the OriginalTargetConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureTargetParamsForRecoverFileAndFolder) GetOriginalTargetConfigOk() (*AzureRecoverFilesOriginalTargetConfig, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OriginalTargetConfig.Get(), o.OriginalTargetConfig.IsSet()
}

// HasOriginalTargetConfig returns a boolean if a field has been set.
func (o *AzureTargetParamsForRecoverFileAndFolder) HasOriginalTargetConfig() bool {
	if o != nil && o.OriginalTargetConfig.IsSet() {
		return true
	}

	return false
}

// SetOriginalTargetConfig gets a reference to the given NullableAzureRecoverFilesOriginalTargetConfig and assigns it to the OriginalTargetConfig field.
func (o *AzureTargetParamsForRecoverFileAndFolder) SetOriginalTargetConfig(v AzureRecoverFilesOriginalTargetConfig) {
	o.OriginalTargetConfig.Set(&v)
}
// SetOriginalTargetConfigNil sets the value for OriginalTargetConfig to be an explicit nil
func (o *AzureTargetParamsForRecoverFileAndFolder) SetOriginalTargetConfigNil() {
	o.OriginalTargetConfig.Set(nil)
}

// UnsetOriginalTargetConfig ensures that no value is present for OriginalTargetConfig, not even an explicit nil
func (o *AzureTargetParamsForRecoverFileAndFolder) UnsetOriginalTargetConfig() {
	o.OriginalTargetConfig.Unset()
}

// GetNewTargetConfig returns the NewTargetConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureTargetParamsForRecoverFileAndFolder) GetNewTargetConfig() AzureRecoverFilesNewTargetConfig {
	if o == nil || o.NewTargetConfig.Get() == nil {
		var ret AzureRecoverFilesNewTargetConfig
		return ret
	}
	return *o.NewTargetConfig.Get()
}

// GetNewTargetConfigOk returns a tuple with the NewTargetConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureTargetParamsForRecoverFileAndFolder) GetNewTargetConfigOk() (*AzureRecoverFilesNewTargetConfig, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NewTargetConfig.Get(), o.NewTargetConfig.IsSet()
}

// HasNewTargetConfig returns a boolean if a field has been set.
func (o *AzureTargetParamsForRecoverFileAndFolder) HasNewTargetConfig() bool {
	if o != nil && o.NewTargetConfig.IsSet() {
		return true
	}

	return false
}

// SetNewTargetConfig gets a reference to the given NullableAzureRecoverFilesNewTargetConfig and assigns it to the NewTargetConfig field.
func (o *AzureTargetParamsForRecoverFileAndFolder) SetNewTargetConfig(v AzureRecoverFilesNewTargetConfig) {
	o.NewTargetConfig.Set(&v)
}
// SetNewTargetConfigNil sets the value for NewTargetConfig to be an explicit nil
func (o *AzureTargetParamsForRecoverFileAndFolder) SetNewTargetConfigNil() {
	o.NewTargetConfig.Set(nil)
}

// UnsetNewTargetConfig ensures that no value is present for NewTargetConfig, not even an explicit nil
func (o *AzureTargetParamsForRecoverFileAndFolder) UnsetNewTargetConfig() {
	o.NewTargetConfig.Unset()
}

// GetOverwriteExisting returns the OverwriteExisting field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureTargetParamsForRecoverFileAndFolder) GetOverwriteExisting() bool {
	if o == nil || o.OverwriteExisting.Get() == nil {
		var ret bool
		return ret
	}
	return *o.OverwriteExisting.Get()
}

// GetOverwriteExistingOk returns a tuple with the OverwriteExisting field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureTargetParamsForRecoverFileAndFolder) GetOverwriteExistingOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OverwriteExisting.Get(), o.OverwriteExisting.IsSet()
}

// HasOverwriteExisting returns a boolean if a field has been set.
func (o *AzureTargetParamsForRecoverFileAndFolder) HasOverwriteExisting() bool {
	if o != nil && o.OverwriteExisting.IsSet() {
		return true
	}

	return false
}

// SetOverwriteExisting gets a reference to the given NullableBool and assigns it to the OverwriteExisting field.
func (o *AzureTargetParamsForRecoverFileAndFolder) SetOverwriteExisting(v bool) {
	o.OverwriteExisting.Set(&v)
}
// SetOverwriteExistingNil sets the value for OverwriteExisting to be an explicit nil
func (o *AzureTargetParamsForRecoverFileAndFolder) SetOverwriteExistingNil() {
	o.OverwriteExisting.Set(nil)
}

// UnsetOverwriteExisting ensures that no value is present for OverwriteExisting, not even an explicit nil
func (o *AzureTargetParamsForRecoverFileAndFolder) UnsetOverwriteExisting() {
	o.OverwriteExisting.Unset()
}

// GetPreserveAttributes returns the PreserveAttributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureTargetParamsForRecoverFileAndFolder) GetPreserveAttributes() bool {
	if o == nil || o.PreserveAttributes.Get() == nil {
		var ret bool
		return ret
	}
	return *o.PreserveAttributes.Get()
}

// GetPreserveAttributesOk returns a tuple with the PreserveAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureTargetParamsForRecoverFileAndFolder) GetPreserveAttributesOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PreserveAttributes.Get(), o.PreserveAttributes.IsSet()
}

// HasPreserveAttributes returns a boolean if a field has been set.
func (o *AzureTargetParamsForRecoverFileAndFolder) HasPreserveAttributes() bool {
	if o != nil && o.PreserveAttributes.IsSet() {
		return true
	}

	return false
}

// SetPreserveAttributes gets a reference to the given NullableBool and assigns it to the PreserveAttributes field.
func (o *AzureTargetParamsForRecoverFileAndFolder) SetPreserveAttributes(v bool) {
	o.PreserveAttributes.Set(&v)
}
// SetPreserveAttributesNil sets the value for PreserveAttributes to be an explicit nil
func (o *AzureTargetParamsForRecoverFileAndFolder) SetPreserveAttributesNil() {
	o.PreserveAttributes.Set(nil)
}

// UnsetPreserveAttributes ensures that no value is present for PreserveAttributes, not even an explicit nil
func (o *AzureTargetParamsForRecoverFileAndFolder) UnsetPreserveAttributes() {
	o.PreserveAttributes.Unset()
}

// GetContinueOnError returns the ContinueOnError field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureTargetParamsForRecoverFileAndFolder) GetContinueOnError() bool {
	if o == nil || o.ContinueOnError.Get() == nil {
		var ret bool
		return ret
	}
	return *o.ContinueOnError.Get()
}

// GetContinueOnErrorOk returns a tuple with the ContinueOnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureTargetParamsForRecoverFileAndFolder) GetContinueOnErrorOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContinueOnError.Get(), o.ContinueOnError.IsSet()
}

// HasContinueOnError returns a boolean if a field has been set.
func (o *AzureTargetParamsForRecoverFileAndFolder) HasContinueOnError() bool {
	if o != nil && o.ContinueOnError.IsSet() {
		return true
	}

	return false
}

// SetContinueOnError gets a reference to the given NullableBool and assigns it to the ContinueOnError field.
func (o *AzureTargetParamsForRecoverFileAndFolder) SetContinueOnError(v bool) {
	o.ContinueOnError.Set(&v)
}
// SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil
func (o *AzureTargetParamsForRecoverFileAndFolder) SetContinueOnErrorNil() {
	o.ContinueOnError.Set(nil)
}

// UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
func (o *AzureTargetParamsForRecoverFileAndFolder) UnsetContinueOnError() {
	o.ContinueOnError.Unset()
}

// GetVlanConfig returns the VlanConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureTargetParamsForRecoverFileAndFolder) GetVlanConfig() RecoveryVlanConfig {
	if o == nil || o.VlanConfig.Get() == nil {
		var ret RecoveryVlanConfig
		return ret
	}
	return *o.VlanConfig.Get()
}

// GetVlanConfigOk returns a tuple with the VlanConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureTargetParamsForRecoverFileAndFolder) GetVlanConfigOk() (*RecoveryVlanConfig, bool) {
	if o == nil  {
		return nil, false
	}
	return o.VlanConfig.Get(), o.VlanConfig.IsSet()
}

// HasVlanConfig returns a boolean if a field has been set.
func (o *AzureTargetParamsForRecoverFileAndFolder) HasVlanConfig() bool {
	if o != nil && o.VlanConfig.IsSet() {
		return true
	}

	return false
}

// SetVlanConfig gets a reference to the given NullableRecoveryVlanConfig and assigns it to the VlanConfig field.
func (o *AzureTargetParamsForRecoverFileAndFolder) SetVlanConfig(v RecoveryVlanConfig) {
	o.VlanConfig.Set(&v)
}
// SetVlanConfigNil sets the value for VlanConfig to be an explicit nil
func (o *AzureTargetParamsForRecoverFileAndFolder) SetVlanConfigNil() {
	o.VlanConfig.Set(nil)
}

// UnsetVlanConfig ensures that no value is present for VlanConfig, not even an explicit nil
func (o *AzureTargetParamsForRecoverFileAndFolder) UnsetVlanConfig() {
	o.VlanConfig.Unset()
}

func (o AzureTargetParamsForRecoverFileAndFolder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["recoverToOriginalTarget"] = o.RecoverToOriginalTarget.Get()
	}
	if o.OriginalTargetConfig.IsSet() {
		toSerialize["originalTargetConfig"] = o.OriginalTargetConfig.Get()
	}
	if o.NewTargetConfig.IsSet() {
		toSerialize["newTargetConfig"] = o.NewTargetConfig.Get()
	}
	if o.OverwriteExisting.IsSet() {
		toSerialize["overwriteExisting"] = o.OverwriteExisting.Get()
	}
	if o.PreserveAttributes.IsSet() {
		toSerialize["preserveAttributes"] = o.PreserveAttributes.Get()
	}
	if o.ContinueOnError.IsSet() {
		toSerialize["continueOnError"] = o.ContinueOnError.Get()
	}
	if o.VlanConfig.IsSet() {
		toSerialize["vlanConfig"] = o.VlanConfig.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAzureTargetParamsForRecoverFileAndFolder struct {
	value *AzureTargetParamsForRecoverFileAndFolder
	isSet bool
}

func (v NullableAzureTargetParamsForRecoverFileAndFolder) Get() *AzureTargetParamsForRecoverFileAndFolder {
	return v.value
}

func (v *NullableAzureTargetParamsForRecoverFileAndFolder) Set(val *AzureTargetParamsForRecoverFileAndFolder) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureTargetParamsForRecoverFileAndFolder) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureTargetParamsForRecoverFileAndFolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureTargetParamsForRecoverFileAndFolder(val *AzureTargetParamsForRecoverFileAndFolder) *NullableAzureTargetParamsForRecoverFileAndFolder {
	return &NullableAzureTargetParamsForRecoverFileAndFolder{value: val, isSet: true}
}

func (v NullableAzureTargetParamsForRecoverFileAndFolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureTargetParamsForRecoverFileAndFolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


