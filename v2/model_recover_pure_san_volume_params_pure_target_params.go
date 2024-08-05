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

// checks if the RecoverPureSanVolumeParamsPureTargetParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecoverPureSanVolumeParamsPureTargetParams{}

// RecoverPureSanVolumeParamsPureTargetParams Specifies the parameters of the Pure SAN volume to recover to.
type RecoverPureSanVolumeParamsPureTargetParams struct {
	NewSourceConfig NullableRecoverPureVolumeTargetParamsNewSourceConfig `json:"newSourceConfig,omitempty"`
	OriginalSourceConfig NullableRecoverPureVolumeTargetParamsOriginalSourceConfig `json:"originalSourceConfig,omitempty"`
	// Specifies whether to recover to a new source.
	RecoverToNewSource bool `json:"recoverToNewSource"`
	// Specifies whether to use thin clone to restore storage array snapshots.
	UseThinClone NullableBool `json:"useThinClone,omitempty"`
}

type _RecoverPureSanVolumeParamsPureTargetParams RecoverPureSanVolumeParamsPureTargetParams

// NewRecoverPureSanVolumeParamsPureTargetParams instantiates a new RecoverPureSanVolumeParamsPureTargetParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverPureSanVolumeParamsPureTargetParams(recoverToNewSource bool) *RecoverPureSanVolumeParamsPureTargetParams {
	this := RecoverPureSanVolumeParamsPureTargetParams{}
	this.RecoverToNewSource = recoverToNewSource
	return &this
}

// NewRecoverPureSanVolumeParamsPureTargetParamsWithDefaults instantiates a new RecoverPureSanVolumeParamsPureTargetParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverPureSanVolumeParamsPureTargetParamsWithDefaults() *RecoverPureSanVolumeParamsPureTargetParams {
	this := RecoverPureSanVolumeParamsPureTargetParams{}
	return &this
}

// GetNewSourceConfig returns the NewSourceConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverPureSanVolumeParamsPureTargetParams) GetNewSourceConfig() RecoverPureVolumeTargetParamsNewSourceConfig {
	if o == nil || IsNil(o.NewSourceConfig.Get()) {
		var ret RecoverPureVolumeTargetParamsNewSourceConfig
		return ret
	}
	return *o.NewSourceConfig.Get()
}

// GetNewSourceConfigOk returns a tuple with the NewSourceConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverPureSanVolumeParamsPureTargetParams) GetNewSourceConfigOk() (*RecoverPureVolumeTargetParamsNewSourceConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.NewSourceConfig.Get(), o.NewSourceConfig.IsSet()
}

// HasNewSourceConfig returns a boolean if a field has been set.
func (o *RecoverPureSanVolumeParamsPureTargetParams) HasNewSourceConfig() bool {
	if o != nil && o.NewSourceConfig.IsSet() {
		return true
	}

	return false
}

// SetNewSourceConfig gets a reference to the given NullableRecoverPureVolumeTargetParamsNewSourceConfig and assigns it to the NewSourceConfig field.
func (o *RecoverPureSanVolumeParamsPureTargetParams) SetNewSourceConfig(v RecoverPureVolumeTargetParamsNewSourceConfig) {
	o.NewSourceConfig.Set(&v)
}
// SetNewSourceConfigNil sets the value for NewSourceConfig to be an explicit nil
func (o *RecoverPureSanVolumeParamsPureTargetParams) SetNewSourceConfigNil() {
	o.NewSourceConfig.Set(nil)
}

// UnsetNewSourceConfig ensures that no value is present for NewSourceConfig, not even an explicit nil
func (o *RecoverPureSanVolumeParamsPureTargetParams) UnsetNewSourceConfig() {
	o.NewSourceConfig.Unset()
}

// GetOriginalSourceConfig returns the OriginalSourceConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverPureSanVolumeParamsPureTargetParams) GetOriginalSourceConfig() RecoverPureVolumeTargetParamsOriginalSourceConfig {
	if o == nil || IsNil(o.OriginalSourceConfig.Get()) {
		var ret RecoverPureVolumeTargetParamsOriginalSourceConfig
		return ret
	}
	return *o.OriginalSourceConfig.Get()
}

// GetOriginalSourceConfigOk returns a tuple with the OriginalSourceConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverPureSanVolumeParamsPureTargetParams) GetOriginalSourceConfigOk() (*RecoverPureVolumeTargetParamsOriginalSourceConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.OriginalSourceConfig.Get(), o.OriginalSourceConfig.IsSet()
}

// HasOriginalSourceConfig returns a boolean if a field has been set.
func (o *RecoverPureSanVolumeParamsPureTargetParams) HasOriginalSourceConfig() bool {
	if o != nil && o.OriginalSourceConfig.IsSet() {
		return true
	}

	return false
}

// SetOriginalSourceConfig gets a reference to the given NullableRecoverPureVolumeTargetParamsOriginalSourceConfig and assigns it to the OriginalSourceConfig field.
func (o *RecoverPureSanVolumeParamsPureTargetParams) SetOriginalSourceConfig(v RecoverPureVolumeTargetParamsOriginalSourceConfig) {
	o.OriginalSourceConfig.Set(&v)
}
// SetOriginalSourceConfigNil sets the value for OriginalSourceConfig to be an explicit nil
func (o *RecoverPureSanVolumeParamsPureTargetParams) SetOriginalSourceConfigNil() {
	o.OriginalSourceConfig.Set(nil)
}

// UnsetOriginalSourceConfig ensures that no value is present for OriginalSourceConfig, not even an explicit nil
func (o *RecoverPureSanVolumeParamsPureTargetParams) UnsetOriginalSourceConfig() {
	o.OriginalSourceConfig.Unset()
}

// GetRecoverToNewSource returns the RecoverToNewSource field value
func (o *RecoverPureSanVolumeParamsPureTargetParams) GetRecoverToNewSource() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RecoverToNewSource
}

// GetRecoverToNewSourceOk returns a tuple with the RecoverToNewSource field value
// and a boolean to check if the value has been set.
func (o *RecoverPureSanVolumeParamsPureTargetParams) GetRecoverToNewSourceOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecoverToNewSource, true
}

// SetRecoverToNewSource sets field value
func (o *RecoverPureSanVolumeParamsPureTargetParams) SetRecoverToNewSource(v bool) {
	o.RecoverToNewSource = v
}

// GetUseThinClone returns the UseThinClone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverPureSanVolumeParamsPureTargetParams) GetUseThinClone() bool {
	if o == nil || IsNil(o.UseThinClone.Get()) {
		var ret bool
		return ret
	}
	return *o.UseThinClone.Get()
}

// GetUseThinCloneOk returns a tuple with the UseThinClone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverPureSanVolumeParamsPureTargetParams) GetUseThinCloneOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.UseThinClone.Get(), o.UseThinClone.IsSet()
}

// HasUseThinClone returns a boolean if a field has been set.
func (o *RecoverPureSanVolumeParamsPureTargetParams) HasUseThinClone() bool {
	if o != nil && o.UseThinClone.IsSet() {
		return true
	}

	return false
}

// SetUseThinClone gets a reference to the given NullableBool and assigns it to the UseThinClone field.
func (o *RecoverPureSanVolumeParamsPureTargetParams) SetUseThinClone(v bool) {
	o.UseThinClone.Set(&v)
}
// SetUseThinCloneNil sets the value for UseThinClone to be an explicit nil
func (o *RecoverPureSanVolumeParamsPureTargetParams) SetUseThinCloneNil() {
	o.UseThinClone.Set(nil)
}

// UnsetUseThinClone ensures that no value is present for UseThinClone, not even an explicit nil
func (o *RecoverPureSanVolumeParamsPureTargetParams) UnsetUseThinClone() {
	o.UseThinClone.Unset()
}

func (o RecoverPureSanVolumeParamsPureTargetParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecoverPureSanVolumeParamsPureTargetParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.NewSourceConfig.IsSet() {
		toSerialize["newSourceConfig"] = o.NewSourceConfig.Get()
	}
	if o.OriginalSourceConfig.IsSet() {
		toSerialize["originalSourceConfig"] = o.OriginalSourceConfig.Get()
	}
	toSerialize["recoverToNewSource"] = o.RecoverToNewSource
	if o.UseThinClone.IsSet() {
		toSerialize["useThinClone"] = o.UseThinClone.Get()
	}
	return toSerialize, nil
}

func (o *RecoverPureSanVolumeParamsPureTargetParams) UnmarshalJSON(data []byte) (err error) {
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

	varRecoverPureSanVolumeParamsPureTargetParams := _RecoverPureSanVolumeParamsPureTargetParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRecoverPureSanVolumeParamsPureTargetParams)

	if err != nil {
		return err
	}

	*o = RecoverPureSanVolumeParamsPureTargetParams(varRecoverPureSanVolumeParamsPureTargetParams)

	return err
}

type NullableRecoverPureSanVolumeParamsPureTargetParams struct {
	value *RecoverPureSanVolumeParamsPureTargetParams
	isSet bool
}

func (v NullableRecoverPureSanVolumeParamsPureTargetParams) Get() *RecoverPureSanVolumeParamsPureTargetParams {
	return v.value
}

func (v *NullableRecoverPureSanVolumeParamsPureTargetParams) Set(val *RecoverPureSanVolumeParamsPureTargetParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverPureSanVolumeParamsPureTargetParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverPureSanVolumeParamsPureTargetParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverPureSanVolumeParamsPureTargetParams(val *RecoverPureSanVolumeParamsPureTargetParams) *NullableRecoverPureSanVolumeParamsPureTargetParams {
	return &NullableRecoverPureSanVolumeParamsPureTargetParams{value: val, isSet: true}
}

func (v NullableRecoverPureSanVolumeParamsPureTargetParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverPureSanVolumeParamsPureTargetParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


