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

// checks if the AwsTargetParamsForRecoverS3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsTargetParamsForRecoverS3{}

// AwsTargetParamsForRecoverS3 Specifies the parameters for an AWS recovery target.
type AwsTargetParamsForRecoverS3 struct {
	// Specifies whether to continue restore on receiving error or not. Default is true.
	ContinueOnError NullableBool `json:"continueOnError,omitempty"`
	NewTargetConfig NullableAwsTargetParamsForRecoverS3NewTargetConfig `json:"newTargetConfig,omitempty"`
	// Specifies the prefix to be added to all the objects being recovered.
	ObjectPrefix NullableString `json:"objectPrefix,omitempty"`
	// Specifies whether to override the existing objects. Default is false.
	OverwriteExisting NullableBool `json:"overwriteExisting,omitempty"`
	// Specifies whether to preserve the objects attributes at the time of restore. Default is true.
	PreserveAttributes NullableBool `json:"preserveAttributes,omitempty"`
	// Specifies whether to recover to the original target. If true, originalTargetConfig must be specified. If false, newTargetConfig must be specified.
	RecoverToOriginalTarget NullableBool `json:"recoverToOriginalTarget"`
}

type _AwsTargetParamsForRecoverS3 AwsTargetParamsForRecoverS3

// NewAwsTargetParamsForRecoverS3 instantiates a new AwsTargetParamsForRecoverS3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsTargetParamsForRecoverS3(recoverToOriginalTarget NullableBool) *AwsTargetParamsForRecoverS3 {
	this := AwsTargetParamsForRecoverS3{}
	this.RecoverToOriginalTarget = recoverToOriginalTarget
	return &this
}

// NewAwsTargetParamsForRecoverS3WithDefaults instantiates a new AwsTargetParamsForRecoverS3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsTargetParamsForRecoverS3WithDefaults() *AwsTargetParamsForRecoverS3 {
	this := AwsTargetParamsForRecoverS3{}
	return &this
}

// GetContinueOnError returns the ContinueOnError field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsTargetParamsForRecoverS3) GetContinueOnError() bool {
	if o == nil || IsNil(o.ContinueOnError.Get()) {
		var ret bool
		return ret
	}
	return *o.ContinueOnError.Get()
}

// GetContinueOnErrorOk returns a tuple with the ContinueOnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsTargetParamsForRecoverS3) GetContinueOnErrorOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContinueOnError.Get(), o.ContinueOnError.IsSet()
}

// HasContinueOnError returns a boolean if a field has been set.
func (o *AwsTargetParamsForRecoverS3) HasContinueOnError() bool {
	if o != nil && o.ContinueOnError.IsSet() {
		return true
	}

	return false
}

// SetContinueOnError gets a reference to the given NullableBool and assigns it to the ContinueOnError field.
func (o *AwsTargetParamsForRecoverS3) SetContinueOnError(v bool) {
	o.ContinueOnError.Set(&v)
}
// SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil
func (o *AwsTargetParamsForRecoverS3) SetContinueOnErrorNil() {
	o.ContinueOnError.Set(nil)
}

// UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
func (o *AwsTargetParamsForRecoverS3) UnsetContinueOnError() {
	o.ContinueOnError.Unset()
}

// GetNewTargetConfig returns the NewTargetConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsTargetParamsForRecoverS3) GetNewTargetConfig() AwsTargetParamsForRecoverS3NewTargetConfig {
	if o == nil || IsNil(o.NewTargetConfig.Get()) {
		var ret AwsTargetParamsForRecoverS3NewTargetConfig
		return ret
	}
	return *o.NewTargetConfig.Get()
}

// GetNewTargetConfigOk returns a tuple with the NewTargetConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsTargetParamsForRecoverS3) GetNewTargetConfigOk() (*AwsTargetParamsForRecoverS3NewTargetConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.NewTargetConfig.Get(), o.NewTargetConfig.IsSet()
}

// HasNewTargetConfig returns a boolean if a field has been set.
func (o *AwsTargetParamsForRecoverS3) HasNewTargetConfig() bool {
	if o != nil && o.NewTargetConfig.IsSet() {
		return true
	}

	return false
}

// SetNewTargetConfig gets a reference to the given NullableAwsTargetParamsForRecoverS3NewTargetConfig and assigns it to the NewTargetConfig field.
func (o *AwsTargetParamsForRecoverS3) SetNewTargetConfig(v AwsTargetParamsForRecoverS3NewTargetConfig) {
	o.NewTargetConfig.Set(&v)
}
// SetNewTargetConfigNil sets the value for NewTargetConfig to be an explicit nil
func (o *AwsTargetParamsForRecoverS3) SetNewTargetConfigNil() {
	o.NewTargetConfig.Set(nil)
}

// UnsetNewTargetConfig ensures that no value is present for NewTargetConfig, not even an explicit nil
func (o *AwsTargetParamsForRecoverS3) UnsetNewTargetConfig() {
	o.NewTargetConfig.Unset()
}

// GetObjectPrefix returns the ObjectPrefix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsTargetParamsForRecoverS3) GetObjectPrefix() string {
	if o == nil || IsNil(o.ObjectPrefix.Get()) {
		var ret string
		return ret
	}
	return *o.ObjectPrefix.Get()
}

// GetObjectPrefixOk returns a tuple with the ObjectPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsTargetParamsForRecoverS3) GetObjectPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ObjectPrefix.Get(), o.ObjectPrefix.IsSet()
}

// HasObjectPrefix returns a boolean if a field has been set.
func (o *AwsTargetParamsForRecoverS3) HasObjectPrefix() bool {
	if o != nil && o.ObjectPrefix.IsSet() {
		return true
	}

	return false
}

// SetObjectPrefix gets a reference to the given NullableString and assigns it to the ObjectPrefix field.
func (o *AwsTargetParamsForRecoverS3) SetObjectPrefix(v string) {
	o.ObjectPrefix.Set(&v)
}
// SetObjectPrefixNil sets the value for ObjectPrefix to be an explicit nil
func (o *AwsTargetParamsForRecoverS3) SetObjectPrefixNil() {
	o.ObjectPrefix.Set(nil)
}

// UnsetObjectPrefix ensures that no value is present for ObjectPrefix, not even an explicit nil
func (o *AwsTargetParamsForRecoverS3) UnsetObjectPrefix() {
	o.ObjectPrefix.Unset()
}

// GetOverwriteExisting returns the OverwriteExisting field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsTargetParamsForRecoverS3) GetOverwriteExisting() bool {
	if o == nil || IsNil(o.OverwriteExisting.Get()) {
		var ret bool
		return ret
	}
	return *o.OverwriteExisting.Get()
}

// GetOverwriteExistingOk returns a tuple with the OverwriteExisting field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsTargetParamsForRecoverS3) GetOverwriteExistingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.OverwriteExisting.Get(), o.OverwriteExisting.IsSet()
}

// HasOverwriteExisting returns a boolean if a field has been set.
func (o *AwsTargetParamsForRecoverS3) HasOverwriteExisting() bool {
	if o != nil && o.OverwriteExisting.IsSet() {
		return true
	}

	return false
}

// SetOverwriteExisting gets a reference to the given NullableBool and assigns it to the OverwriteExisting field.
func (o *AwsTargetParamsForRecoverS3) SetOverwriteExisting(v bool) {
	o.OverwriteExisting.Set(&v)
}
// SetOverwriteExistingNil sets the value for OverwriteExisting to be an explicit nil
func (o *AwsTargetParamsForRecoverS3) SetOverwriteExistingNil() {
	o.OverwriteExisting.Set(nil)
}

// UnsetOverwriteExisting ensures that no value is present for OverwriteExisting, not even an explicit nil
func (o *AwsTargetParamsForRecoverS3) UnsetOverwriteExisting() {
	o.OverwriteExisting.Unset()
}

// GetPreserveAttributes returns the PreserveAttributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsTargetParamsForRecoverS3) GetPreserveAttributes() bool {
	if o == nil || IsNil(o.PreserveAttributes.Get()) {
		var ret bool
		return ret
	}
	return *o.PreserveAttributes.Get()
}

// GetPreserveAttributesOk returns a tuple with the PreserveAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsTargetParamsForRecoverS3) GetPreserveAttributesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreserveAttributes.Get(), o.PreserveAttributes.IsSet()
}

// HasPreserveAttributes returns a boolean if a field has been set.
func (o *AwsTargetParamsForRecoverS3) HasPreserveAttributes() bool {
	if o != nil && o.PreserveAttributes.IsSet() {
		return true
	}

	return false
}

// SetPreserveAttributes gets a reference to the given NullableBool and assigns it to the PreserveAttributes field.
func (o *AwsTargetParamsForRecoverS3) SetPreserveAttributes(v bool) {
	o.PreserveAttributes.Set(&v)
}
// SetPreserveAttributesNil sets the value for PreserveAttributes to be an explicit nil
func (o *AwsTargetParamsForRecoverS3) SetPreserveAttributesNil() {
	o.PreserveAttributes.Set(nil)
}

// UnsetPreserveAttributes ensures that no value is present for PreserveAttributes, not even an explicit nil
func (o *AwsTargetParamsForRecoverS3) UnsetPreserveAttributes() {
	o.PreserveAttributes.Unset()
}

// GetRecoverToOriginalTarget returns the RecoverToOriginalTarget field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *AwsTargetParamsForRecoverS3) GetRecoverToOriginalTarget() bool {
	if o == nil || o.RecoverToOriginalTarget.Get() == nil {
		var ret bool
		return ret
	}

	return *o.RecoverToOriginalTarget.Get()
}

// GetRecoverToOriginalTargetOk returns a tuple with the RecoverToOriginalTarget field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsTargetParamsForRecoverS3) GetRecoverToOriginalTargetOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecoverToOriginalTarget.Get(), o.RecoverToOriginalTarget.IsSet()
}

// SetRecoverToOriginalTarget sets field value
func (o *AwsTargetParamsForRecoverS3) SetRecoverToOriginalTarget(v bool) {
	o.RecoverToOriginalTarget.Set(&v)
}

func (o AwsTargetParamsForRecoverS3) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsTargetParamsForRecoverS3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ContinueOnError.IsSet() {
		toSerialize["continueOnError"] = o.ContinueOnError.Get()
	}
	if o.NewTargetConfig.IsSet() {
		toSerialize["newTargetConfig"] = o.NewTargetConfig.Get()
	}
	if o.ObjectPrefix.IsSet() {
		toSerialize["objectPrefix"] = o.ObjectPrefix.Get()
	}
	if o.OverwriteExisting.IsSet() {
		toSerialize["overwriteExisting"] = o.OverwriteExisting.Get()
	}
	if o.PreserveAttributes.IsSet() {
		toSerialize["preserveAttributes"] = o.PreserveAttributes.Get()
	}
	toSerialize["recoverToOriginalTarget"] = o.RecoverToOriginalTarget.Get()
	return toSerialize, nil
}

func (o *AwsTargetParamsForRecoverS3) UnmarshalJSON(data []byte) (err error) {
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

	varAwsTargetParamsForRecoverS3 := _AwsTargetParamsForRecoverS3{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAwsTargetParamsForRecoverS3)

	if err != nil {
		return err
	}

	*o = AwsTargetParamsForRecoverS3(varAwsTargetParamsForRecoverS3)

	return err
}

type NullableAwsTargetParamsForRecoverS3 struct {
	value *AwsTargetParamsForRecoverS3
	isSet bool
}

func (v NullableAwsTargetParamsForRecoverS3) Get() *AwsTargetParamsForRecoverS3 {
	return v.value
}

func (v *NullableAwsTargetParamsForRecoverS3) Set(val *AwsTargetParamsForRecoverS3) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsTargetParamsForRecoverS3) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsTargetParamsForRecoverS3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsTargetParamsForRecoverS3(val *AwsTargetParamsForRecoverS3) *NullableAwsTargetParamsForRecoverS3 {
	return &NullableAwsTargetParamsForRecoverS3{value: val, isSet: true}
}

func (v NullableAwsTargetParamsForRecoverS3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsTargetParamsForRecoverS3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


