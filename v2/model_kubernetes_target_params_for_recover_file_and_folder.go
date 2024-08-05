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

// checks if the KubernetesTargetParamsForRecoverFileAndFolder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesTargetParamsForRecoverFileAndFolder{}

// KubernetesTargetParamsForRecoverFileAndFolder Specifies the parameters for a Kubernetes recovery target.
type KubernetesTargetParamsForRecoverFileAndFolder struct {
	// Specifies whether to continue recovering other files if one of files or folders failed to recover. Default value is false.
	ContinueOnError NullableBool `json:"continueOnError,omitempty"`
	NewTargetConfig NullableKubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig `json:"newTargetConfig,omitempty"`
	OriginalTargetConfig NullableKubernetesTargetParamsForRecoverFileAndFolderOriginalTargetConfig `json:"originalTargetConfig,omitempty"`
	// Specifies whether to overwrite the existing files. Default is true.
	OverwriteExisting NullableBool `json:"overwriteExisting,omitempty"`
	// Specifies whether to preserve original attributes. Default is true.
	PreserveAttributes NullableBool `json:"preserveAttributes,omitempty"`
	// Specifies whether to recover to the original target. If true, originalTargetConfig must be specified. If false, newTargetConfig must be specified.
	RecoverToOriginalTarget NullableBool `json:"recoverToOriginalTarget"`
	VlanConfig NullableAcropolisTargetParamsForRecoverFileAndFolderVlanConfig `json:"vlanConfig,omitempty"`
}

type _KubernetesTargetParamsForRecoverFileAndFolder KubernetesTargetParamsForRecoverFileAndFolder

// NewKubernetesTargetParamsForRecoverFileAndFolder instantiates a new KubernetesTargetParamsForRecoverFileAndFolder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesTargetParamsForRecoverFileAndFolder(recoverToOriginalTarget NullableBool) *KubernetesTargetParamsForRecoverFileAndFolder {
	this := KubernetesTargetParamsForRecoverFileAndFolder{}
	this.RecoverToOriginalTarget = recoverToOriginalTarget
	return &this
}

// NewKubernetesTargetParamsForRecoverFileAndFolderWithDefaults instantiates a new KubernetesTargetParamsForRecoverFileAndFolder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesTargetParamsForRecoverFileAndFolderWithDefaults() *KubernetesTargetParamsForRecoverFileAndFolder {
	this := KubernetesTargetParamsForRecoverFileAndFolder{}
	return &this
}

// GetContinueOnError returns the ContinueOnError field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesTargetParamsForRecoverFileAndFolder) GetContinueOnError() bool {
	if o == nil || IsNil(o.ContinueOnError.Get()) {
		var ret bool
		return ret
	}
	return *o.ContinueOnError.Get()
}

// GetContinueOnErrorOk returns a tuple with the ContinueOnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesTargetParamsForRecoverFileAndFolder) GetContinueOnErrorOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContinueOnError.Get(), o.ContinueOnError.IsSet()
}

// HasContinueOnError returns a boolean if a field has been set.
func (o *KubernetesTargetParamsForRecoverFileAndFolder) HasContinueOnError() bool {
	if o != nil && o.ContinueOnError.IsSet() {
		return true
	}

	return false
}

// SetContinueOnError gets a reference to the given NullableBool and assigns it to the ContinueOnError field.
func (o *KubernetesTargetParamsForRecoverFileAndFolder) SetContinueOnError(v bool) {
	o.ContinueOnError.Set(&v)
}
// SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil
func (o *KubernetesTargetParamsForRecoverFileAndFolder) SetContinueOnErrorNil() {
	o.ContinueOnError.Set(nil)
}

// UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
func (o *KubernetesTargetParamsForRecoverFileAndFolder) UnsetContinueOnError() {
	o.ContinueOnError.Unset()
}

// GetNewTargetConfig returns the NewTargetConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesTargetParamsForRecoverFileAndFolder) GetNewTargetConfig() KubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig {
	if o == nil || IsNil(o.NewTargetConfig.Get()) {
		var ret KubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig
		return ret
	}
	return *o.NewTargetConfig.Get()
}

// GetNewTargetConfigOk returns a tuple with the NewTargetConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesTargetParamsForRecoverFileAndFolder) GetNewTargetConfigOk() (*KubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.NewTargetConfig.Get(), o.NewTargetConfig.IsSet()
}

// HasNewTargetConfig returns a boolean if a field has been set.
func (o *KubernetesTargetParamsForRecoverFileAndFolder) HasNewTargetConfig() bool {
	if o != nil && o.NewTargetConfig.IsSet() {
		return true
	}

	return false
}

// SetNewTargetConfig gets a reference to the given NullableKubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig and assigns it to the NewTargetConfig field.
func (o *KubernetesTargetParamsForRecoverFileAndFolder) SetNewTargetConfig(v KubernetesTargetParamsForRecoverFileAndFolderNewTargetConfig) {
	o.NewTargetConfig.Set(&v)
}
// SetNewTargetConfigNil sets the value for NewTargetConfig to be an explicit nil
func (o *KubernetesTargetParamsForRecoverFileAndFolder) SetNewTargetConfigNil() {
	o.NewTargetConfig.Set(nil)
}

// UnsetNewTargetConfig ensures that no value is present for NewTargetConfig, not even an explicit nil
func (o *KubernetesTargetParamsForRecoverFileAndFolder) UnsetNewTargetConfig() {
	o.NewTargetConfig.Unset()
}

// GetOriginalTargetConfig returns the OriginalTargetConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesTargetParamsForRecoverFileAndFolder) GetOriginalTargetConfig() KubernetesTargetParamsForRecoverFileAndFolderOriginalTargetConfig {
	if o == nil || IsNil(o.OriginalTargetConfig.Get()) {
		var ret KubernetesTargetParamsForRecoverFileAndFolderOriginalTargetConfig
		return ret
	}
	return *o.OriginalTargetConfig.Get()
}

// GetOriginalTargetConfigOk returns a tuple with the OriginalTargetConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesTargetParamsForRecoverFileAndFolder) GetOriginalTargetConfigOk() (*KubernetesTargetParamsForRecoverFileAndFolderOriginalTargetConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.OriginalTargetConfig.Get(), o.OriginalTargetConfig.IsSet()
}

// HasOriginalTargetConfig returns a boolean if a field has been set.
func (o *KubernetesTargetParamsForRecoverFileAndFolder) HasOriginalTargetConfig() bool {
	if o != nil && o.OriginalTargetConfig.IsSet() {
		return true
	}

	return false
}

// SetOriginalTargetConfig gets a reference to the given NullableKubernetesTargetParamsForRecoverFileAndFolderOriginalTargetConfig and assigns it to the OriginalTargetConfig field.
func (o *KubernetesTargetParamsForRecoverFileAndFolder) SetOriginalTargetConfig(v KubernetesTargetParamsForRecoverFileAndFolderOriginalTargetConfig) {
	o.OriginalTargetConfig.Set(&v)
}
// SetOriginalTargetConfigNil sets the value for OriginalTargetConfig to be an explicit nil
func (o *KubernetesTargetParamsForRecoverFileAndFolder) SetOriginalTargetConfigNil() {
	o.OriginalTargetConfig.Set(nil)
}

// UnsetOriginalTargetConfig ensures that no value is present for OriginalTargetConfig, not even an explicit nil
func (o *KubernetesTargetParamsForRecoverFileAndFolder) UnsetOriginalTargetConfig() {
	o.OriginalTargetConfig.Unset()
}

// GetOverwriteExisting returns the OverwriteExisting field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesTargetParamsForRecoverFileAndFolder) GetOverwriteExisting() bool {
	if o == nil || IsNil(o.OverwriteExisting.Get()) {
		var ret bool
		return ret
	}
	return *o.OverwriteExisting.Get()
}

// GetOverwriteExistingOk returns a tuple with the OverwriteExisting field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesTargetParamsForRecoverFileAndFolder) GetOverwriteExistingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.OverwriteExisting.Get(), o.OverwriteExisting.IsSet()
}

// HasOverwriteExisting returns a boolean if a field has been set.
func (o *KubernetesTargetParamsForRecoverFileAndFolder) HasOverwriteExisting() bool {
	if o != nil && o.OverwriteExisting.IsSet() {
		return true
	}

	return false
}

// SetOverwriteExisting gets a reference to the given NullableBool and assigns it to the OverwriteExisting field.
func (o *KubernetesTargetParamsForRecoverFileAndFolder) SetOverwriteExisting(v bool) {
	o.OverwriteExisting.Set(&v)
}
// SetOverwriteExistingNil sets the value for OverwriteExisting to be an explicit nil
func (o *KubernetesTargetParamsForRecoverFileAndFolder) SetOverwriteExistingNil() {
	o.OverwriteExisting.Set(nil)
}

// UnsetOverwriteExisting ensures that no value is present for OverwriteExisting, not even an explicit nil
func (o *KubernetesTargetParamsForRecoverFileAndFolder) UnsetOverwriteExisting() {
	o.OverwriteExisting.Unset()
}

// GetPreserveAttributes returns the PreserveAttributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesTargetParamsForRecoverFileAndFolder) GetPreserveAttributes() bool {
	if o == nil || IsNil(o.PreserveAttributes.Get()) {
		var ret bool
		return ret
	}
	return *o.PreserveAttributes.Get()
}

// GetPreserveAttributesOk returns a tuple with the PreserveAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesTargetParamsForRecoverFileAndFolder) GetPreserveAttributesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreserveAttributes.Get(), o.PreserveAttributes.IsSet()
}

// HasPreserveAttributes returns a boolean if a field has been set.
func (o *KubernetesTargetParamsForRecoverFileAndFolder) HasPreserveAttributes() bool {
	if o != nil && o.PreserveAttributes.IsSet() {
		return true
	}

	return false
}

// SetPreserveAttributes gets a reference to the given NullableBool and assigns it to the PreserveAttributes field.
func (o *KubernetesTargetParamsForRecoverFileAndFolder) SetPreserveAttributes(v bool) {
	o.PreserveAttributes.Set(&v)
}
// SetPreserveAttributesNil sets the value for PreserveAttributes to be an explicit nil
func (o *KubernetesTargetParamsForRecoverFileAndFolder) SetPreserveAttributesNil() {
	o.PreserveAttributes.Set(nil)
}

// UnsetPreserveAttributes ensures that no value is present for PreserveAttributes, not even an explicit nil
func (o *KubernetesTargetParamsForRecoverFileAndFolder) UnsetPreserveAttributes() {
	o.PreserveAttributes.Unset()
}

// GetRecoverToOriginalTarget returns the RecoverToOriginalTarget field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *KubernetesTargetParamsForRecoverFileAndFolder) GetRecoverToOriginalTarget() bool {
	if o == nil || o.RecoverToOriginalTarget.Get() == nil {
		var ret bool
		return ret
	}

	return *o.RecoverToOriginalTarget.Get()
}

// GetRecoverToOriginalTargetOk returns a tuple with the RecoverToOriginalTarget field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesTargetParamsForRecoverFileAndFolder) GetRecoverToOriginalTargetOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecoverToOriginalTarget.Get(), o.RecoverToOriginalTarget.IsSet()
}

// SetRecoverToOriginalTarget sets field value
func (o *KubernetesTargetParamsForRecoverFileAndFolder) SetRecoverToOriginalTarget(v bool) {
	o.RecoverToOriginalTarget.Set(&v)
}

// GetVlanConfig returns the VlanConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesTargetParamsForRecoverFileAndFolder) GetVlanConfig() AcropolisTargetParamsForRecoverFileAndFolderVlanConfig {
	if o == nil || IsNil(o.VlanConfig.Get()) {
		var ret AcropolisTargetParamsForRecoverFileAndFolderVlanConfig
		return ret
	}
	return *o.VlanConfig.Get()
}

// GetVlanConfigOk returns a tuple with the VlanConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesTargetParamsForRecoverFileAndFolder) GetVlanConfigOk() (*AcropolisTargetParamsForRecoverFileAndFolderVlanConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.VlanConfig.Get(), o.VlanConfig.IsSet()
}

// HasVlanConfig returns a boolean if a field has been set.
func (o *KubernetesTargetParamsForRecoverFileAndFolder) HasVlanConfig() bool {
	if o != nil && o.VlanConfig.IsSet() {
		return true
	}

	return false
}

// SetVlanConfig gets a reference to the given NullableAcropolisTargetParamsForRecoverFileAndFolderVlanConfig and assigns it to the VlanConfig field.
func (o *KubernetesTargetParamsForRecoverFileAndFolder) SetVlanConfig(v AcropolisTargetParamsForRecoverFileAndFolderVlanConfig) {
	o.VlanConfig.Set(&v)
}
// SetVlanConfigNil sets the value for VlanConfig to be an explicit nil
func (o *KubernetesTargetParamsForRecoverFileAndFolder) SetVlanConfigNil() {
	o.VlanConfig.Set(nil)
}

// UnsetVlanConfig ensures that no value is present for VlanConfig, not even an explicit nil
func (o *KubernetesTargetParamsForRecoverFileAndFolder) UnsetVlanConfig() {
	o.VlanConfig.Unset()
}

func (o KubernetesTargetParamsForRecoverFileAndFolder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesTargetParamsForRecoverFileAndFolder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ContinueOnError.IsSet() {
		toSerialize["continueOnError"] = o.ContinueOnError.Get()
	}
	if o.NewTargetConfig.IsSet() {
		toSerialize["newTargetConfig"] = o.NewTargetConfig.Get()
	}
	if o.OriginalTargetConfig.IsSet() {
		toSerialize["originalTargetConfig"] = o.OriginalTargetConfig.Get()
	}
	if o.OverwriteExisting.IsSet() {
		toSerialize["overwriteExisting"] = o.OverwriteExisting.Get()
	}
	if o.PreserveAttributes.IsSet() {
		toSerialize["preserveAttributes"] = o.PreserveAttributes.Get()
	}
	toSerialize["recoverToOriginalTarget"] = o.RecoverToOriginalTarget.Get()
	if o.VlanConfig.IsSet() {
		toSerialize["vlanConfig"] = o.VlanConfig.Get()
	}
	return toSerialize, nil
}

func (o *KubernetesTargetParamsForRecoverFileAndFolder) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"recoverToOriginalTarget",
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

	varKubernetesTargetParamsForRecoverFileAndFolder := _KubernetesTargetParamsForRecoverFileAndFolder{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKubernetesTargetParamsForRecoverFileAndFolder)

	if err != nil {
		return err
	}

	*o = KubernetesTargetParamsForRecoverFileAndFolder(varKubernetesTargetParamsForRecoverFileAndFolder)

	return err
}

type NullableKubernetesTargetParamsForRecoverFileAndFolder struct {
	value *KubernetesTargetParamsForRecoverFileAndFolder
	isSet bool
}

func (v NullableKubernetesTargetParamsForRecoverFileAndFolder) Get() *KubernetesTargetParamsForRecoverFileAndFolder {
	return v.value
}

func (v *NullableKubernetesTargetParamsForRecoverFileAndFolder) Set(val *KubernetesTargetParamsForRecoverFileAndFolder) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesTargetParamsForRecoverFileAndFolder) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesTargetParamsForRecoverFileAndFolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesTargetParamsForRecoverFileAndFolder(val *KubernetesTargetParamsForRecoverFileAndFolder) *NullableKubernetesTargetParamsForRecoverFileAndFolder {
	return &NullableKubernetesTargetParamsForRecoverFileAndFolder{value: val, isSet: true}
}

func (v NullableKubernetesTargetParamsForRecoverFileAndFolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesTargetParamsForRecoverFileAndFolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


