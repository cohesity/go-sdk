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

// PhysicalTargetParamsForRecoverFileAndFolder Specifies the parameters for a Physical recovery target.
type PhysicalTargetParamsForRecoverFileAndFolder struct {
	// Specifies the target entity where the volumes are being mounted.
	RecoverTarget RecoverTarget `json:"recoverTarget"`
	// If this is true, then files will be restored to original paths.
	RestoreToOriginalPaths NullableBool `json:"restoreToOriginalPaths,omitempty"`
	// Specifies whether to overwrite existing file/folder during recovery.
	OverwriteExisting NullableBool `json:"overwriteExisting,omitempty"`
	// Specifies the directory path where restore should happen if restore_to_original_paths is set to false.
	AlternateRestoreDirectory NullableString `json:"alternateRestoreDirectory,omitempty"`
	// Specifies whether to preserve file/folder attributes during recovery.
	PreserveAttributes NullableBool `json:"preserveAttributes,omitempty"`
	// Whether to preserve the original time stamps.
	PreserveTimestamps NullableBool `json:"preserveTimestamps,omitempty"`
	// Whether to preserve the ACLs of the original file.
	PreserveAcls NullableBool `json:"preserveAcls,omitempty"`
	// Specifies whether to continue recovering other volumes if one of the volumes fails to recover. Default value is false.
	ContinueOnError NullableBool `json:"continueOnError,omitempty"`
	// Specifies VLAN Params associated with the recovered. If this is not specified, then the VLAN settings will be automatically selected from one of the below options: a. If VLANs are configured on Cohesity, then the VLAN host/VIP will be automatically based on the client's (e.g. ESXI host) IP address. b. If VLANs are not configured on Cohesity, then the partition hostname or VIPs will be used for Recovery.
	VlanConfig NullableRecoveryVlanConfig `json:"vlanConfig,omitempty"`
}

// NewPhysicalTargetParamsForRecoverFileAndFolder instantiates a new PhysicalTargetParamsForRecoverFileAndFolder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhysicalTargetParamsForRecoverFileAndFolder(recoverTarget RecoverTarget) *PhysicalTargetParamsForRecoverFileAndFolder {
	this := PhysicalTargetParamsForRecoverFileAndFolder{}
	this.RecoverTarget = recoverTarget
	return &this
}

// NewPhysicalTargetParamsForRecoverFileAndFolderWithDefaults instantiates a new PhysicalTargetParamsForRecoverFileAndFolder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhysicalTargetParamsForRecoverFileAndFolderWithDefaults() *PhysicalTargetParamsForRecoverFileAndFolder {
	this := PhysicalTargetParamsForRecoverFileAndFolder{}
	return &this
}

// GetRecoverTarget returns the RecoverTarget field value
func (o *PhysicalTargetParamsForRecoverFileAndFolder) GetRecoverTarget() RecoverTarget {
	if o == nil {
		var ret RecoverTarget
		return ret
	}

	return o.RecoverTarget
}

// GetRecoverTargetOk returns a tuple with the RecoverTarget field value
// and a boolean to check if the value has been set.
func (o *PhysicalTargetParamsForRecoverFileAndFolder) GetRecoverTargetOk() (*RecoverTarget, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RecoverTarget, true
}

// SetRecoverTarget sets field value
func (o *PhysicalTargetParamsForRecoverFileAndFolder) SetRecoverTarget(v RecoverTarget) {
	o.RecoverTarget = v
}

// GetRestoreToOriginalPaths returns the RestoreToOriginalPaths field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PhysicalTargetParamsForRecoverFileAndFolder) GetRestoreToOriginalPaths() bool {
	if o == nil || o.RestoreToOriginalPaths.Get() == nil {
		var ret bool
		return ret
	}
	return *o.RestoreToOriginalPaths.Get()
}

// GetRestoreToOriginalPathsOk returns a tuple with the RestoreToOriginalPaths field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PhysicalTargetParamsForRecoverFileAndFolder) GetRestoreToOriginalPathsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RestoreToOriginalPaths.Get(), o.RestoreToOriginalPaths.IsSet()
}

// HasRestoreToOriginalPaths returns a boolean if a field has been set.
func (o *PhysicalTargetParamsForRecoverFileAndFolder) HasRestoreToOriginalPaths() bool {
	if o != nil && o.RestoreToOriginalPaths.IsSet() {
		return true
	}

	return false
}

// SetRestoreToOriginalPaths gets a reference to the given NullableBool and assigns it to the RestoreToOriginalPaths field.
func (o *PhysicalTargetParamsForRecoverFileAndFolder) SetRestoreToOriginalPaths(v bool) {
	o.RestoreToOriginalPaths.Set(&v)
}
// SetRestoreToOriginalPathsNil sets the value for RestoreToOriginalPaths to be an explicit nil
func (o *PhysicalTargetParamsForRecoverFileAndFolder) SetRestoreToOriginalPathsNil() {
	o.RestoreToOriginalPaths.Set(nil)
}

// UnsetRestoreToOriginalPaths ensures that no value is present for RestoreToOriginalPaths, not even an explicit nil
func (o *PhysicalTargetParamsForRecoverFileAndFolder) UnsetRestoreToOriginalPaths() {
	o.RestoreToOriginalPaths.Unset()
}

// GetOverwriteExisting returns the OverwriteExisting field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PhysicalTargetParamsForRecoverFileAndFolder) GetOverwriteExisting() bool {
	if o == nil || o.OverwriteExisting.Get() == nil {
		var ret bool
		return ret
	}
	return *o.OverwriteExisting.Get()
}

// GetOverwriteExistingOk returns a tuple with the OverwriteExisting field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PhysicalTargetParamsForRecoverFileAndFolder) GetOverwriteExistingOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OverwriteExisting.Get(), o.OverwriteExisting.IsSet()
}

// HasOverwriteExisting returns a boolean if a field has been set.
func (o *PhysicalTargetParamsForRecoverFileAndFolder) HasOverwriteExisting() bool {
	if o != nil && o.OverwriteExisting.IsSet() {
		return true
	}

	return false
}

// SetOverwriteExisting gets a reference to the given NullableBool and assigns it to the OverwriteExisting field.
func (o *PhysicalTargetParamsForRecoverFileAndFolder) SetOverwriteExisting(v bool) {
	o.OverwriteExisting.Set(&v)
}
// SetOverwriteExistingNil sets the value for OverwriteExisting to be an explicit nil
func (o *PhysicalTargetParamsForRecoverFileAndFolder) SetOverwriteExistingNil() {
	o.OverwriteExisting.Set(nil)
}

// UnsetOverwriteExisting ensures that no value is present for OverwriteExisting, not even an explicit nil
func (o *PhysicalTargetParamsForRecoverFileAndFolder) UnsetOverwriteExisting() {
	o.OverwriteExisting.Unset()
}

// GetAlternateRestoreDirectory returns the AlternateRestoreDirectory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PhysicalTargetParamsForRecoverFileAndFolder) GetAlternateRestoreDirectory() string {
	if o == nil || o.AlternateRestoreDirectory.Get() == nil {
		var ret string
		return ret
	}
	return *o.AlternateRestoreDirectory.Get()
}

// GetAlternateRestoreDirectoryOk returns a tuple with the AlternateRestoreDirectory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PhysicalTargetParamsForRecoverFileAndFolder) GetAlternateRestoreDirectoryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AlternateRestoreDirectory.Get(), o.AlternateRestoreDirectory.IsSet()
}

// HasAlternateRestoreDirectory returns a boolean if a field has been set.
func (o *PhysicalTargetParamsForRecoverFileAndFolder) HasAlternateRestoreDirectory() bool {
	if o != nil && o.AlternateRestoreDirectory.IsSet() {
		return true
	}

	return false
}

// SetAlternateRestoreDirectory gets a reference to the given NullableString and assigns it to the AlternateRestoreDirectory field.
func (o *PhysicalTargetParamsForRecoverFileAndFolder) SetAlternateRestoreDirectory(v string) {
	o.AlternateRestoreDirectory.Set(&v)
}
// SetAlternateRestoreDirectoryNil sets the value for AlternateRestoreDirectory to be an explicit nil
func (o *PhysicalTargetParamsForRecoverFileAndFolder) SetAlternateRestoreDirectoryNil() {
	o.AlternateRestoreDirectory.Set(nil)
}

// UnsetAlternateRestoreDirectory ensures that no value is present for AlternateRestoreDirectory, not even an explicit nil
func (o *PhysicalTargetParamsForRecoverFileAndFolder) UnsetAlternateRestoreDirectory() {
	o.AlternateRestoreDirectory.Unset()
}

// GetPreserveAttributes returns the PreserveAttributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PhysicalTargetParamsForRecoverFileAndFolder) GetPreserveAttributes() bool {
	if o == nil || o.PreserveAttributes.Get() == nil {
		var ret bool
		return ret
	}
	return *o.PreserveAttributes.Get()
}

// GetPreserveAttributesOk returns a tuple with the PreserveAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PhysicalTargetParamsForRecoverFileAndFolder) GetPreserveAttributesOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PreserveAttributes.Get(), o.PreserveAttributes.IsSet()
}

// HasPreserveAttributes returns a boolean if a field has been set.
func (o *PhysicalTargetParamsForRecoverFileAndFolder) HasPreserveAttributes() bool {
	if o != nil && o.PreserveAttributes.IsSet() {
		return true
	}

	return false
}

// SetPreserveAttributes gets a reference to the given NullableBool and assigns it to the PreserveAttributes field.
func (o *PhysicalTargetParamsForRecoverFileAndFolder) SetPreserveAttributes(v bool) {
	o.PreserveAttributes.Set(&v)
}
// SetPreserveAttributesNil sets the value for PreserveAttributes to be an explicit nil
func (o *PhysicalTargetParamsForRecoverFileAndFolder) SetPreserveAttributesNil() {
	o.PreserveAttributes.Set(nil)
}

// UnsetPreserveAttributes ensures that no value is present for PreserveAttributes, not even an explicit nil
func (o *PhysicalTargetParamsForRecoverFileAndFolder) UnsetPreserveAttributes() {
	o.PreserveAttributes.Unset()
}

// GetPreserveTimestamps returns the PreserveTimestamps field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PhysicalTargetParamsForRecoverFileAndFolder) GetPreserveTimestamps() bool {
	if o == nil || o.PreserveTimestamps.Get() == nil {
		var ret bool
		return ret
	}
	return *o.PreserveTimestamps.Get()
}

// GetPreserveTimestampsOk returns a tuple with the PreserveTimestamps field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PhysicalTargetParamsForRecoverFileAndFolder) GetPreserveTimestampsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PreserveTimestamps.Get(), o.PreserveTimestamps.IsSet()
}

// HasPreserveTimestamps returns a boolean if a field has been set.
func (o *PhysicalTargetParamsForRecoverFileAndFolder) HasPreserveTimestamps() bool {
	if o != nil && o.PreserveTimestamps.IsSet() {
		return true
	}

	return false
}

// SetPreserveTimestamps gets a reference to the given NullableBool and assigns it to the PreserveTimestamps field.
func (o *PhysicalTargetParamsForRecoverFileAndFolder) SetPreserveTimestamps(v bool) {
	o.PreserveTimestamps.Set(&v)
}
// SetPreserveTimestampsNil sets the value for PreserveTimestamps to be an explicit nil
func (o *PhysicalTargetParamsForRecoverFileAndFolder) SetPreserveTimestampsNil() {
	o.PreserveTimestamps.Set(nil)
}

// UnsetPreserveTimestamps ensures that no value is present for PreserveTimestamps, not even an explicit nil
func (o *PhysicalTargetParamsForRecoverFileAndFolder) UnsetPreserveTimestamps() {
	o.PreserveTimestamps.Unset()
}

// GetPreserveAcls returns the PreserveAcls field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PhysicalTargetParamsForRecoverFileAndFolder) GetPreserveAcls() bool {
	if o == nil || o.PreserveAcls.Get() == nil {
		var ret bool
		return ret
	}
	return *o.PreserveAcls.Get()
}

// GetPreserveAclsOk returns a tuple with the PreserveAcls field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PhysicalTargetParamsForRecoverFileAndFolder) GetPreserveAclsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PreserveAcls.Get(), o.PreserveAcls.IsSet()
}

// HasPreserveAcls returns a boolean if a field has been set.
func (o *PhysicalTargetParamsForRecoverFileAndFolder) HasPreserveAcls() bool {
	if o != nil && o.PreserveAcls.IsSet() {
		return true
	}

	return false
}

// SetPreserveAcls gets a reference to the given NullableBool and assigns it to the PreserveAcls field.
func (o *PhysicalTargetParamsForRecoverFileAndFolder) SetPreserveAcls(v bool) {
	o.PreserveAcls.Set(&v)
}
// SetPreserveAclsNil sets the value for PreserveAcls to be an explicit nil
func (o *PhysicalTargetParamsForRecoverFileAndFolder) SetPreserveAclsNil() {
	o.PreserveAcls.Set(nil)
}

// UnsetPreserveAcls ensures that no value is present for PreserveAcls, not even an explicit nil
func (o *PhysicalTargetParamsForRecoverFileAndFolder) UnsetPreserveAcls() {
	o.PreserveAcls.Unset()
}

// GetContinueOnError returns the ContinueOnError field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PhysicalTargetParamsForRecoverFileAndFolder) GetContinueOnError() bool {
	if o == nil || o.ContinueOnError.Get() == nil {
		var ret bool
		return ret
	}
	return *o.ContinueOnError.Get()
}

// GetContinueOnErrorOk returns a tuple with the ContinueOnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PhysicalTargetParamsForRecoverFileAndFolder) GetContinueOnErrorOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContinueOnError.Get(), o.ContinueOnError.IsSet()
}

// HasContinueOnError returns a boolean if a field has been set.
func (o *PhysicalTargetParamsForRecoverFileAndFolder) HasContinueOnError() bool {
	if o != nil && o.ContinueOnError.IsSet() {
		return true
	}

	return false
}

// SetContinueOnError gets a reference to the given NullableBool and assigns it to the ContinueOnError field.
func (o *PhysicalTargetParamsForRecoverFileAndFolder) SetContinueOnError(v bool) {
	o.ContinueOnError.Set(&v)
}
// SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil
func (o *PhysicalTargetParamsForRecoverFileAndFolder) SetContinueOnErrorNil() {
	o.ContinueOnError.Set(nil)
}

// UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
func (o *PhysicalTargetParamsForRecoverFileAndFolder) UnsetContinueOnError() {
	o.ContinueOnError.Unset()
}

// GetVlanConfig returns the VlanConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PhysicalTargetParamsForRecoverFileAndFolder) GetVlanConfig() RecoveryVlanConfig {
	if o == nil || o.VlanConfig.Get() == nil {
		var ret RecoveryVlanConfig
		return ret
	}
	return *o.VlanConfig.Get()
}

// GetVlanConfigOk returns a tuple with the VlanConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PhysicalTargetParamsForRecoverFileAndFolder) GetVlanConfigOk() (*RecoveryVlanConfig, bool) {
	if o == nil  {
		return nil, false
	}
	return o.VlanConfig.Get(), o.VlanConfig.IsSet()
}

// HasVlanConfig returns a boolean if a field has been set.
func (o *PhysicalTargetParamsForRecoverFileAndFolder) HasVlanConfig() bool {
	if o != nil && o.VlanConfig.IsSet() {
		return true
	}

	return false
}

// SetVlanConfig gets a reference to the given NullableRecoveryVlanConfig and assigns it to the VlanConfig field.
func (o *PhysicalTargetParamsForRecoverFileAndFolder) SetVlanConfig(v RecoveryVlanConfig) {
	o.VlanConfig.Set(&v)
}
// SetVlanConfigNil sets the value for VlanConfig to be an explicit nil
func (o *PhysicalTargetParamsForRecoverFileAndFolder) SetVlanConfigNil() {
	o.VlanConfig.Set(nil)
}

// UnsetVlanConfig ensures that no value is present for VlanConfig, not even an explicit nil
func (o *PhysicalTargetParamsForRecoverFileAndFolder) UnsetVlanConfig() {
	o.VlanConfig.Unset()
}

func (o PhysicalTargetParamsForRecoverFileAndFolder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["recoverTarget"] = o.RecoverTarget
	}
	if o.RestoreToOriginalPaths.IsSet() {
		toSerialize["restoreToOriginalPaths"] = o.RestoreToOriginalPaths.Get()
	}
	if o.OverwriteExisting.IsSet() {
		toSerialize["overwriteExisting"] = o.OverwriteExisting.Get()
	}
	if o.AlternateRestoreDirectory.IsSet() {
		toSerialize["alternateRestoreDirectory"] = o.AlternateRestoreDirectory.Get()
	}
	if o.PreserveAttributes.IsSet() {
		toSerialize["preserveAttributes"] = o.PreserveAttributes.Get()
	}
	if o.PreserveTimestamps.IsSet() {
		toSerialize["preserveTimestamps"] = o.PreserveTimestamps.Get()
	}
	if o.PreserveAcls.IsSet() {
		toSerialize["preserveAcls"] = o.PreserveAcls.Get()
	}
	if o.ContinueOnError.IsSet() {
		toSerialize["continueOnError"] = o.ContinueOnError.Get()
	}
	if o.VlanConfig.IsSet() {
		toSerialize["vlanConfig"] = o.VlanConfig.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePhysicalTargetParamsForRecoverFileAndFolder struct {
	value *PhysicalTargetParamsForRecoverFileAndFolder
	isSet bool
}

func (v NullablePhysicalTargetParamsForRecoverFileAndFolder) Get() *PhysicalTargetParamsForRecoverFileAndFolder {
	return v.value
}

func (v *NullablePhysicalTargetParamsForRecoverFileAndFolder) Set(val *PhysicalTargetParamsForRecoverFileAndFolder) {
	v.value = val
	v.isSet = true
}

func (v NullablePhysicalTargetParamsForRecoverFileAndFolder) IsSet() bool {
	return v.isSet
}

func (v *NullablePhysicalTargetParamsForRecoverFileAndFolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhysicalTargetParamsForRecoverFileAndFolder(val *PhysicalTargetParamsForRecoverFileAndFolder) *NullablePhysicalTargetParamsForRecoverFileAndFolder {
	return &NullablePhysicalTargetParamsForRecoverFileAndFolder{value: val, isSet: true}
}

func (v NullablePhysicalTargetParamsForRecoverFileAndFolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhysicalTargetParamsForRecoverFileAndFolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


