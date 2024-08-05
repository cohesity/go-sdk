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

// checks if the RecoverNetappToNetappVolumeTargetParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecoverNetappToNetappVolumeTargetParams{}

// RecoverNetappToNetappVolumeTargetParams Specifies the params of the Netapp recovery target.
type RecoverNetappToNetappVolumeTargetParams struct {
	NewSourceConfig NullableRecoverNetappToNetappVolumeTargetParamsNewSourceConfig `json:"newSourceConfig,omitempty"`
	OriginalSourceConfig NullableRecoverNetappToNetappVolumeTargetParamsOriginalSourceConfig `json:"originalSourceConfig,omitempty"`
	// Specifies the parameter whether the recovery should be performed to a new or the original Netapp target.
	RecoverToNewSource bool `json:"recoverToNewSource"`
}

type _RecoverNetappToNetappVolumeTargetParams RecoverNetappToNetappVolumeTargetParams

// NewRecoverNetappToNetappVolumeTargetParams instantiates a new RecoverNetappToNetappVolumeTargetParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverNetappToNetappVolumeTargetParams(recoverToNewSource bool) *RecoverNetappToNetappVolumeTargetParams {
	this := RecoverNetappToNetappVolumeTargetParams{}
	this.RecoverToNewSource = recoverToNewSource
	return &this
}

// NewRecoverNetappToNetappVolumeTargetParamsWithDefaults instantiates a new RecoverNetappToNetappVolumeTargetParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverNetappToNetappVolumeTargetParamsWithDefaults() *RecoverNetappToNetappVolumeTargetParams {
	this := RecoverNetappToNetappVolumeTargetParams{}
	return &this
}

// GetNewSourceConfig returns the NewSourceConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverNetappToNetappVolumeTargetParams) GetNewSourceConfig() RecoverNetappToNetappVolumeTargetParamsNewSourceConfig {
	if o == nil || IsNil(o.NewSourceConfig.Get()) {
		var ret RecoverNetappToNetappVolumeTargetParamsNewSourceConfig
		return ret
	}
	return *o.NewSourceConfig.Get()
}

// GetNewSourceConfigOk returns a tuple with the NewSourceConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverNetappToNetappVolumeTargetParams) GetNewSourceConfigOk() (*RecoverNetappToNetappVolumeTargetParamsNewSourceConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.NewSourceConfig.Get(), o.NewSourceConfig.IsSet()
}

// HasNewSourceConfig returns a boolean if a field has been set.
func (o *RecoverNetappToNetappVolumeTargetParams) HasNewSourceConfig() bool {
	if o != nil && o.NewSourceConfig.IsSet() {
		return true
	}

	return false
}

// SetNewSourceConfig gets a reference to the given NullableRecoverNetappToNetappVolumeTargetParamsNewSourceConfig and assigns it to the NewSourceConfig field.
func (o *RecoverNetappToNetappVolumeTargetParams) SetNewSourceConfig(v RecoverNetappToNetappVolumeTargetParamsNewSourceConfig) {
	o.NewSourceConfig.Set(&v)
}
// SetNewSourceConfigNil sets the value for NewSourceConfig to be an explicit nil
func (o *RecoverNetappToNetappVolumeTargetParams) SetNewSourceConfigNil() {
	o.NewSourceConfig.Set(nil)
}

// UnsetNewSourceConfig ensures that no value is present for NewSourceConfig, not even an explicit nil
func (o *RecoverNetappToNetappVolumeTargetParams) UnsetNewSourceConfig() {
	o.NewSourceConfig.Unset()
}

// GetOriginalSourceConfig returns the OriginalSourceConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverNetappToNetappVolumeTargetParams) GetOriginalSourceConfig() RecoverNetappToNetappVolumeTargetParamsOriginalSourceConfig {
	if o == nil || IsNil(o.OriginalSourceConfig.Get()) {
		var ret RecoverNetappToNetappVolumeTargetParamsOriginalSourceConfig
		return ret
	}
	return *o.OriginalSourceConfig.Get()
}

// GetOriginalSourceConfigOk returns a tuple with the OriginalSourceConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverNetappToNetappVolumeTargetParams) GetOriginalSourceConfigOk() (*RecoverNetappToNetappVolumeTargetParamsOriginalSourceConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.OriginalSourceConfig.Get(), o.OriginalSourceConfig.IsSet()
}

// HasOriginalSourceConfig returns a boolean if a field has been set.
func (o *RecoverNetappToNetappVolumeTargetParams) HasOriginalSourceConfig() bool {
	if o != nil && o.OriginalSourceConfig.IsSet() {
		return true
	}

	return false
}

// SetOriginalSourceConfig gets a reference to the given NullableRecoverNetappToNetappVolumeTargetParamsOriginalSourceConfig and assigns it to the OriginalSourceConfig field.
func (o *RecoverNetappToNetappVolumeTargetParams) SetOriginalSourceConfig(v RecoverNetappToNetappVolumeTargetParamsOriginalSourceConfig) {
	o.OriginalSourceConfig.Set(&v)
}
// SetOriginalSourceConfigNil sets the value for OriginalSourceConfig to be an explicit nil
func (o *RecoverNetappToNetappVolumeTargetParams) SetOriginalSourceConfigNil() {
	o.OriginalSourceConfig.Set(nil)
}

// UnsetOriginalSourceConfig ensures that no value is present for OriginalSourceConfig, not even an explicit nil
func (o *RecoverNetappToNetappVolumeTargetParams) UnsetOriginalSourceConfig() {
	o.OriginalSourceConfig.Unset()
}

// GetRecoverToNewSource returns the RecoverToNewSource field value
func (o *RecoverNetappToNetappVolumeTargetParams) GetRecoverToNewSource() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RecoverToNewSource
}

// GetRecoverToNewSourceOk returns a tuple with the RecoverToNewSource field value
// and a boolean to check if the value has been set.
func (o *RecoverNetappToNetappVolumeTargetParams) GetRecoverToNewSourceOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecoverToNewSource, true
}

// SetRecoverToNewSource sets field value
func (o *RecoverNetappToNetappVolumeTargetParams) SetRecoverToNewSource(v bool) {
	o.RecoverToNewSource = v
}

func (o RecoverNetappToNetappVolumeTargetParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecoverNetappToNetappVolumeTargetParams) ToMap() (map[string]interface{}, error) {
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

func (o *RecoverNetappToNetappVolumeTargetParams) UnmarshalJSON(data []byte) (err error) {
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

	varRecoverNetappToNetappVolumeTargetParams := _RecoverNetappToNetappVolumeTargetParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRecoverNetappToNetappVolumeTargetParams)

	if err != nil {
		return err
	}

	*o = RecoverNetappToNetappVolumeTargetParams(varRecoverNetappToNetappVolumeTargetParams)

	return err
}

type NullableRecoverNetappToNetappVolumeTargetParams struct {
	value *RecoverNetappToNetappVolumeTargetParams
	isSet bool
}

func (v NullableRecoverNetappToNetappVolumeTargetParams) Get() *RecoverNetappToNetappVolumeTargetParams {
	return v.value
}

func (v *NullableRecoverNetappToNetappVolumeTargetParams) Set(val *RecoverNetappToNetappVolumeTargetParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverNetappToNetappVolumeTargetParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverNetappToNetappVolumeTargetParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverNetappToNetappVolumeTargetParams(val *RecoverNetappToNetappVolumeTargetParams) *NullableRecoverNetappToNetappVolumeTargetParams {
	return &NullableRecoverNetappToNetappVolumeTargetParams{value: val, isSet: true}
}

func (v NullableRecoverNetappToNetappVolumeTargetParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverNetappToNetappVolumeTargetParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


