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

// checks if the RecoverNetappNasVolumeParamsGpfsTargetParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecoverNetappNasVolumeParamsGpfsTargetParams{}

// RecoverNetappNasVolumeParamsGpfsTargetParams Specifies the params for a GPFS recovery target.
type RecoverNetappNasVolumeParamsGpfsTargetParams struct {
	// Specifies whether to continue recovering other volumes if one of the volumes fails to recover. Default value is false.
	ContinueOnError NullableBool `json:"continueOnError,omitempty"`
	// Specifies whether encryption should be enabled during recovery.
	EncryptionEnabled NullableBool `json:"encryptionEnabled,omitempty"`
	FilterIpConfig *FilterIpConfig `json:"filterIpConfig,omitempty"`
	// Specifies whether to overwrite existing file/folder during recovery.
	OverwriteExistingFile NullableBool `json:"overwriteExistingFile,omitempty"`
	ParentSource *RecoverOtherNasToElastifileFilesTargetParamsParentSource `json:"parentSource,omitempty"`
	// Specifies whether to preserve file/folder attributes during recovery.
	PreserveFileAttributes NullableBool `json:"preserveFileAttributes,omitempty"`
	VlanConfig *RecoveryVlanConfig `json:"vlanConfig,omitempty"`
	Volume RecoverOtherNasToElastifileVolumeTargetParamsVolume `json:"volume"`
}

type _RecoverNetappNasVolumeParamsGpfsTargetParams RecoverNetappNasVolumeParamsGpfsTargetParams

// NewRecoverNetappNasVolumeParamsGpfsTargetParams instantiates a new RecoverNetappNasVolumeParamsGpfsTargetParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverNetappNasVolumeParamsGpfsTargetParams(volume RecoverOtherNasToElastifileVolumeTargetParamsVolume) *RecoverNetappNasVolumeParamsGpfsTargetParams {
	this := RecoverNetappNasVolumeParamsGpfsTargetParams{}
	this.Volume = volume
	return &this
}

// NewRecoverNetappNasVolumeParamsGpfsTargetParamsWithDefaults instantiates a new RecoverNetappNasVolumeParamsGpfsTargetParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverNetappNasVolumeParamsGpfsTargetParamsWithDefaults() *RecoverNetappNasVolumeParamsGpfsTargetParams {
	this := RecoverNetappNasVolumeParamsGpfsTargetParams{}
	return &this
}

// GetContinueOnError returns the ContinueOnError field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverNetappNasVolumeParamsGpfsTargetParams) GetContinueOnError() bool {
	if o == nil || IsNil(o.ContinueOnError.Get()) {
		var ret bool
		return ret
	}
	return *o.ContinueOnError.Get()
}

// GetContinueOnErrorOk returns a tuple with the ContinueOnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverNetappNasVolumeParamsGpfsTargetParams) GetContinueOnErrorOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContinueOnError.Get(), o.ContinueOnError.IsSet()
}

// HasContinueOnError returns a boolean if a field has been set.
func (o *RecoverNetappNasVolumeParamsGpfsTargetParams) HasContinueOnError() bool {
	if o != nil && o.ContinueOnError.IsSet() {
		return true
	}

	return false
}

// SetContinueOnError gets a reference to the given NullableBool and assigns it to the ContinueOnError field.
func (o *RecoverNetappNasVolumeParamsGpfsTargetParams) SetContinueOnError(v bool) {
	o.ContinueOnError.Set(&v)
}
// SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil
func (o *RecoverNetappNasVolumeParamsGpfsTargetParams) SetContinueOnErrorNil() {
	o.ContinueOnError.Set(nil)
}

// UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
func (o *RecoverNetappNasVolumeParamsGpfsTargetParams) UnsetContinueOnError() {
	o.ContinueOnError.Unset()
}

// GetEncryptionEnabled returns the EncryptionEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverNetappNasVolumeParamsGpfsTargetParams) GetEncryptionEnabled() bool {
	if o == nil || IsNil(o.EncryptionEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.EncryptionEnabled.Get()
}

// GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverNetappNasVolumeParamsGpfsTargetParams) GetEncryptionEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EncryptionEnabled.Get(), o.EncryptionEnabled.IsSet()
}

// HasEncryptionEnabled returns a boolean if a field has been set.
func (o *RecoverNetappNasVolumeParamsGpfsTargetParams) HasEncryptionEnabled() bool {
	if o != nil && o.EncryptionEnabled.IsSet() {
		return true
	}

	return false
}

// SetEncryptionEnabled gets a reference to the given NullableBool and assigns it to the EncryptionEnabled field.
func (o *RecoverNetappNasVolumeParamsGpfsTargetParams) SetEncryptionEnabled(v bool) {
	o.EncryptionEnabled.Set(&v)
}
// SetEncryptionEnabledNil sets the value for EncryptionEnabled to be an explicit nil
func (o *RecoverNetappNasVolumeParamsGpfsTargetParams) SetEncryptionEnabledNil() {
	o.EncryptionEnabled.Set(nil)
}

// UnsetEncryptionEnabled ensures that no value is present for EncryptionEnabled, not even an explicit nil
func (o *RecoverNetappNasVolumeParamsGpfsTargetParams) UnsetEncryptionEnabled() {
	o.EncryptionEnabled.Unset()
}

// GetFilterIpConfig returns the FilterIpConfig field value if set, zero value otherwise.
func (o *RecoverNetappNasVolumeParamsGpfsTargetParams) GetFilterIpConfig() FilterIpConfig {
	if o == nil || IsNil(o.FilterIpConfig) {
		var ret FilterIpConfig
		return ret
	}
	return *o.FilterIpConfig
}

// GetFilterIpConfigOk returns a tuple with the FilterIpConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoverNetappNasVolumeParamsGpfsTargetParams) GetFilterIpConfigOk() (*FilterIpConfig, bool) {
	if o == nil || IsNil(o.FilterIpConfig) {
		return nil, false
	}
	return o.FilterIpConfig, true
}

// HasFilterIpConfig returns a boolean if a field has been set.
func (o *RecoverNetappNasVolumeParamsGpfsTargetParams) HasFilterIpConfig() bool {
	if o != nil && !IsNil(o.FilterIpConfig) {
		return true
	}

	return false
}

// SetFilterIpConfig gets a reference to the given FilterIpConfig and assigns it to the FilterIpConfig field.
func (o *RecoverNetappNasVolumeParamsGpfsTargetParams) SetFilterIpConfig(v FilterIpConfig) {
	o.FilterIpConfig = &v
}

// GetOverwriteExistingFile returns the OverwriteExistingFile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverNetappNasVolumeParamsGpfsTargetParams) GetOverwriteExistingFile() bool {
	if o == nil || IsNil(o.OverwriteExistingFile.Get()) {
		var ret bool
		return ret
	}
	return *o.OverwriteExistingFile.Get()
}

// GetOverwriteExistingFileOk returns a tuple with the OverwriteExistingFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverNetappNasVolumeParamsGpfsTargetParams) GetOverwriteExistingFileOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.OverwriteExistingFile.Get(), o.OverwriteExistingFile.IsSet()
}

// HasOverwriteExistingFile returns a boolean if a field has been set.
func (o *RecoverNetappNasVolumeParamsGpfsTargetParams) HasOverwriteExistingFile() bool {
	if o != nil && o.OverwriteExistingFile.IsSet() {
		return true
	}

	return false
}

// SetOverwriteExistingFile gets a reference to the given NullableBool and assigns it to the OverwriteExistingFile field.
func (o *RecoverNetappNasVolumeParamsGpfsTargetParams) SetOverwriteExistingFile(v bool) {
	o.OverwriteExistingFile.Set(&v)
}
// SetOverwriteExistingFileNil sets the value for OverwriteExistingFile to be an explicit nil
func (o *RecoverNetappNasVolumeParamsGpfsTargetParams) SetOverwriteExistingFileNil() {
	o.OverwriteExistingFile.Set(nil)
}

// UnsetOverwriteExistingFile ensures that no value is present for OverwriteExistingFile, not even an explicit nil
func (o *RecoverNetappNasVolumeParamsGpfsTargetParams) UnsetOverwriteExistingFile() {
	o.OverwriteExistingFile.Unset()
}

// GetParentSource returns the ParentSource field value if set, zero value otherwise.
func (o *RecoverNetappNasVolumeParamsGpfsTargetParams) GetParentSource() RecoverOtherNasToElastifileFilesTargetParamsParentSource {
	if o == nil || IsNil(o.ParentSource) {
		var ret RecoverOtherNasToElastifileFilesTargetParamsParentSource
		return ret
	}
	return *o.ParentSource
}

// GetParentSourceOk returns a tuple with the ParentSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoverNetappNasVolumeParamsGpfsTargetParams) GetParentSourceOk() (*RecoverOtherNasToElastifileFilesTargetParamsParentSource, bool) {
	if o == nil || IsNil(o.ParentSource) {
		return nil, false
	}
	return o.ParentSource, true
}

// HasParentSource returns a boolean if a field has been set.
func (o *RecoverNetappNasVolumeParamsGpfsTargetParams) HasParentSource() bool {
	if o != nil && !IsNil(o.ParentSource) {
		return true
	}

	return false
}

// SetParentSource gets a reference to the given RecoverOtherNasToElastifileFilesTargetParamsParentSource and assigns it to the ParentSource field.
func (o *RecoverNetappNasVolumeParamsGpfsTargetParams) SetParentSource(v RecoverOtherNasToElastifileFilesTargetParamsParentSource) {
	o.ParentSource = &v
}

// GetPreserveFileAttributes returns the PreserveFileAttributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverNetappNasVolumeParamsGpfsTargetParams) GetPreserveFileAttributes() bool {
	if o == nil || IsNil(o.PreserveFileAttributes.Get()) {
		var ret bool
		return ret
	}
	return *o.PreserveFileAttributes.Get()
}

// GetPreserveFileAttributesOk returns a tuple with the PreserveFileAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverNetappNasVolumeParamsGpfsTargetParams) GetPreserveFileAttributesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreserveFileAttributes.Get(), o.PreserveFileAttributes.IsSet()
}

// HasPreserveFileAttributes returns a boolean if a field has been set.
func (o *RecoverNetappNasVolumeParamsGpfsTargetParams) HasPreserveFileAttributes() bool {
	if o != nil && o.PreserveFileAttributes.IsSet() {
		return true
	}

	return false
}

// SetPreserveFileAttributes gets a reference to the given NullableBool and assigns it to the PreserveFileAttributes field.
func (o *RecoverNetappNasVolumeParamsGpfsTargetParams) SetPreserveFileAttributes(v bool) {
	o.PreserveFileAttributes.Set(&v)
}
// SetPreserveFileAttributesNil sets the value for PreserveFileAttributes to be an explicit nil
func (o *RecoverNetappNasVolumeParamsGpfsTargetParams) SetPreserveFileAttributesNil() {
	o.PreserveFileAttributes.Set(nil)
}

// UnsetPreserveFileAttributes ensures that no value is present for PreserveFileAttributes, not even an explicit nil
func (o *RecoverNetappNasVolumeParamsGpfsTargetParams) UnsetPreserveFileAttributes() {
	o.PreserveFileAttributes.Unset()
}

// GetVlanConfig returns the VlanConfig field value if set, zero value otherwise.
func (o *RecoverNetappNasVolumeParamsGpfsTargetParams) GetVlanConfig() RecoveryVlanConfig {
	if o == nil || IsNil(o.VlanConfig) {
		var ret RecoveryVlanConfig
		return ret
	}
	return *o.VlanConfig
}

// GetVlanConfigOk returns a tuple with the VlanConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoverNetappNasVolumeParamsGpfsTargetParams) GetVlanConfigOk() (*RecoveryVlanConfig, bool) {
	if o == nil || IsNil(o.VlanConfig) {
		return nil, false
	}
	return o.VlanConfig, true
}

// HasVlanConfig returns a boolean if a field has been set.
func (o *RecoverNetappNasVolumeParamsGpfsTargetParams) HasVlanConfig() bool {
	if o != nil && !IsNil(o.VlanConfig) {
		return true
	}

	return false
}

// SetVlanConfig gets a reference to the given RecoveryVlanConfig and assigns it to the VlanConfig field.
func (o *RecoverNetappNasVolumeParamsGpfsTargetParams) SetVlanConfig(v RecoveryVlanConfig) {
	o.VlanConfig = &v
}

// GetVolume returns the Volume field value
func (o *RecoverNetappNasVolumeParamsGpfsTargetParams) GetVolume() RecoverOtherNasToElastifileVolumeTargetParamsVolume {
	if o == nil {
		var ret RecoverOtherNasToElastifileVolumeTargetParamsVolume
		return ret
	}

	return o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value
// and a boolean to check if the value has been set.
func (o *RecoverNetappNasVolumeParamsGpfsTargetParams) GetVolumeOk() (*RecoverOtherNasToElastifileVolumeTargetParamsVolume, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Volume, true
}

// SetVolume sets field value
func (o *RecoverNetappNasVolumeParamsGpfsTargetParams) SetVolume(v RecoverOtherNasToElastifileVolumeTargetParamsVolume) {
	o.Volume = v
}

func (o RecoverNetappNasVolumeParamsGpfsTargetParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecoverNetappNasVolumeParamsGpfsTargetParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ContinueOnError.IsSet() {
		toSerialize["continueOnError"] = o.ContinueOnError.Get()
	}
	if o.EncryptionEnabled.IsSet() {
		toSerialize["encryptionEnabled"] = o.EncryptionEnabled.Get()
	}
	if !IsNil(o.FilterIpConfig) {
		toSerialize["filterIpConfig"] = o.FilterIpConfig
	}
	if o.OverwriteExistingFile.IsSet() {
		toSerialize["overwriteExistingFile"] = o.OverwriteExistingFile.Get()
	}
	if !IsNil(o.ParentSource) {
		toSerialize["parentSource"] = o.ParentSource
	}
	if o.PreserveFileAttributes.IsSet() {
		toSerialize["preserveFileAttributes"] = o.PreserveFileAttributes.Get()
	}
	if !IsNil(o.VlanConfig) {
		toSerialize["vlanConfig"] = o.VlanConfig
	}
	toSerialize["volume"] = o.Volume
	return toSerialize, nil
}

func (o *RecoverNetappNasVolumeParamsGpfsTargetParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"volume",
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

	varRecoverNetappNasVolumeParamsGpfsTargetParams := _RecoverNetappNasVolumeParamsGpfsTargetParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRecoverNetappNasVolumeParamsGpfsTargetParams)

	if err != nil {
		return err
	}

	*o = RecoverNetappNasVolumeParamsGpfsTargetParams(varRecoverNetappNasVolumeParamsGpfsTargetParams)

	return err
}

type NullableRecoverNetappNasVolumeParamsGpfsTargetParams struct {
	value *RecoverNetappNasVolumeParamsGpfsTargetParams
	isSet bool
}

func (v NullableRecoverNetappNasVolumeParamsGpfsTargetParams) Get() *RecoverNetappNasVolumeParamsGpfsTargetParams {
	return v.value
}

func (v *NullableRecoverNetappNasVolumeParamsGpfsTargetParams) Set(val *RecoverNetappNasVolumeParamsGpfsTargetParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverNetappNasVolumeParamsGpfsTargetParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverNetappNasVolumeParamsGpfsTargetParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverNetappNasVolumeParamsGpfsTargetParams(val *RecoverNetappNasVolumeParamsGpfsTargetParams) *NullableRecoverNetappNasVolumeParamsGpfsTargetParams {
	return &NullableRecoverNetappNasVolumeParamsGpfsTargetParams{value: val, isSet: true}
}

func (v NullableRecoverNetappNasVolumeParamsGpfsTargetParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverNetappNasVolumeParamsGpfsTargetParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


