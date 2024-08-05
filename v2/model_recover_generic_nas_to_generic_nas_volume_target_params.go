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

// checks if the RecoverGenericNasToGenericNasVolumeTargetParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecoverGenericNasToGenericNasVolumeTargetParams{}

// RecoverGenericNasToGenericNasVolumeTargetParams Specifies the params of the Generic NAS recovery target.
type RecoverGenericNasToGenericNasVolumeTargetParams struct {
	NewSourceConfig NullableRecoverGenericNasToGenericNasVolumeTargetParamsNewSourceConfig `json:"newSourceConfig,omitempty"`
	OriginalSourceConfig NullableRecoverGenericNasToGenericNasVolumeTargetParamsOriginalSourceConfig `json:"originalSourceConfig,omitempty"`
	// Specifies the parameter whether the recovery should be performed to a new or the original Generic NAS target.
	RecoverToNewSource bool `json:"recoverToNewSource"`
}

type _RecoverGenericNasToGenericNasVolumeTargetParams RecoverGenericNasToGenericNasVolumeTargetParams

// NewRecoverGenericNasToGenericNasVolumeTargetParams instantiates a new RecoverGenericNasToGenericNasVolumeTargetParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverGenericNasToGenericNasVolumeTargetParams(recoverToNewSource bool) *RecoverGenericNasToGenericNasVolumeTargetParams {
	this := RecoverGenericNasToGenericNasVolumeTargetParams{}
	this.RecoverToNewSource = recoverToNewSource
	return &this
}

// NewRecoverGenericNasToGenericNasVolumeTargetParamsWithDefaults instantiates a new RecoverGenericNasToGenericNasVolumeTargetParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverGenericNasToGenericNasVolumeTargetParamsWithDefaults() *RecoverGenericNasToGenericNasVolumeTargetParams {
	this := RecoverGenericNasToGenericNasVolumeTargetParams{}
	return &this
}

// GetNewSourceConfig returns the NewSourceConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverGenericNasToGenericNasVolumeTargetParams) GetNewSourceConfig() RecoverGenericNasToGenericNasVolumeTargetParamsNewSourceConfig {
	if o == nil || IsNil(o.NewSourceConfig.Get()) {
		var ret RecoverGenericNasToGenericNasVolumeTargetParamsNewSourceConfig
		return ret
	}
	return *o.NewSourceConfig.Get()
}

// GetNewSourceConfigOk returns a tuple with the NewSourceConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverGenericNasToGenericNasVolumeTargetParams) GetNewSourceConfigOk() (*RecoverGenericNasToGenericNasVolumeTargetParamsNewSourceConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.NewSourceConfig.Get(), o.NewSourceConfig.IsSet()
}

// HasNewSourceConfig returns a boolean if a field has been set.
func (o *RecoverGenericNasToGenericNasVolumeTargetParams) HasNewSourceConfig() bool {
	if o != nil && o.NewSourceConfig.IsSet() {
		return true
	}

	return false
}

// SetNewSourceConfig gets a reference to the given NullableRecoverGenericNasToGenericNasVolumeTargetParamsNewSourceConfig and assigns it to the NewSourceConfig field.
func (o *RecoverGenericNasToGenericNasVolumeTargetParams) SetNewSourceConfig(v RecoverGenericNasToGenericNasVolumeTargetParamsNewSourceConfig) {
	o.NewSourceConfig.Set(&v)
}
// SetNewSourceConfigNil sets the value for NewSourceConfig to be an explicit nil
func (o *RecoverGenericNasToGenericNasVolumeTargetParams) SetNewSourceConfigNil() {
	o.NewSourceConfig.Set(nil)
}

// UnsetNewSourceConfig ensures that no value is present for NewSourceConfig, not even an explicit nil
func (o *RecoverGenericNasToGenericNasVolumeTargetParams) UnsetNewSourceConfig() {
	o.NewSourceConfig.Unset()
}

// GetOriginalSourceConfig returns the OriginalSourceConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverGenericNasToGenericNasVolumeTargetParams) GetOriginalSourceConfig() RecoverGenericNasToGenericNasVolumeTargetParamsOriginalSourceConfig {
	if o == nil || IsNil(o.OriginalSourceConfig.Get()) {
		var ret RecoverGenericNasToGenericNasVolumeTargetParamsOriginalSourceConfig
		return ret
	}
	return *o.OriginalSourceConfig.Get()
}

// GetOriginalSourceConfigOk returns a tuple with the OriginalSourceConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverGenericNasToGenericNasVolumeTargetParams) GetOriginalSourceConfigOk() (*RecoverGenericNasToGenericNasVolumeTargetParamsOriginalSourceConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.OriginalSourceConfig.Get(), o.OriginalSourceConfig.IsSet()
}

// HasOriginalSourceConfig returns a boolean if a field has been set.
func (o *RecoverGenericNasToGenericNasVolumeTargetParams) HasOriginalSourceConfig() bool {
	if o != nil && o.OriginalSourceConfig.IsSet() {
		return true
	}

	return false
}

// SetOriginalSourceConfig gets a reference to the given NullableRecoverGenericNasToGenericNasVolumeTargetParamsOriginalSourceConfig and assigns it to the OriginalSourceConfig field.
func (o *RecoverGenericNasToGenericNasVolumeTargetParams) SetOriginalSourceConfig(v RecoverGenericNasToGenericNasVolumeTargetParamsOriginalSourceConfig) {
	o.OriginalSourceConfig.Set(&v)
}
// SetOriginalSourceConfigNil sets the value for OriginalSourceConfig to be an explicit nil
func (o *RecoverGenericNasToGenericNasVolumeTargetParams) SetOriginalSourceConfigNil() {
	o.OriginalSourceConfig.Set(nil)
}

// UnsetOriginalSourceConfig ensures that no value is present for OriginalSourceConfig, not even an explicit nil
func (o *RecoverGenericNasToGenericNasVolumeTargetParams) UnsetOriginalSourceConfig() {
	o.OriginalSourceConfig.Unset()
}

// GetRecoverToNewSource returns the RecoverToNewSource field value
func (o *RecoverGenericNasToGenericNasVolumeTargetParams) GetRecoverToNewSource() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RecoverToNewSource
}

// GetRecoverToNewSourceOk returns a tuple with the RecoverToNewSource field value
// and a boolean to check if the value has been set.
func (o *RecoverGenericNasToGenericNasVolumeTargetParams) GetRecoverToNewSourceOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecoverToNewSource, true
}

// SetRecoverToNewSource sets field value
func (o *RecoverGenericNasToGenericNasVolumeTargetParams) SetRecoverToNewSource(v bool) {
	o.RecoverToNewSource = v
}

func (o RecoverGenericNasToGenericNasVolumeTargetParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecoverGenericNasToGenericNasVolumeTargetParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.NewSourceConfig.IsSet() {
		toSerialize["newSourceConfig"] = o.NewSourceConfig.Get()
	}
	if o.OriginalSourceConfig.IsSet() {
		toSerialize["originalSourceConfig"] = o.OriginalSourceConfig.Get()
	}
	toSerialize["recoverToNewSource"] = o.RecoverToNewSource
	return toSerialize, nil
}

func (o *RecoverGenericNasToGenericNasVolumeTargetParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"recoverToNewSource",
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

	varRecoverGenericNasToGenericNasVolumeTargetParams := _RecoverGenericNasToGenericNasVolumeTargetParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRecoverGenericNasToGenericNasVolumeTargetParams)

	if err != nil {
		return err
	}

	*o = RecoverGenericNasToGenericNasVolumeTargetParams(varRecoverGenericNasToGenericNasVolumeTargetParams)

	return err
}

type NullableRecoverGenericNasToGenericNasVolumeTargetParams struct {
	value *RecoverGenericNasToGenericNasVolumeTargetParams
	isSet bool
}

func (v NullableRecoverGenericNasToGenericNasVolumeTargetParams) Get() *RecoverGenericNasToGenericNasVolumeTargetParams {
	return v.value
}

func (v *NullableRecoverGenericNasToGenericNasVolumeTargetParams) Set(val *RecoverGenericNasToGenericNasVolumeTargetParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverGenericNasToGenericNasVolumeTargetParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverGenericNasToGenericNasVolumeTargetParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverGenericNasToGenericNasVolumeTargetParams(val *RecoverGenericNasToGenericNasVolumeTargetParams) *NullableRecoverGenericNasToGenericNasVolumeTargetParams {
	return &NullableRecoverGenericNasToGenericNasVolumeTargetParams{value: val, isSet: true}
}

func (v NullableRecoverGenericNasToGenericNasVolumeTargetParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverGenericNasToGenericNasVolumeTargetParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


