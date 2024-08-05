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

// checks if the RecoverNetappToNetappFilesTargetParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecoverNetappToNetappFilesTargetParams{}

// RecoverNetappToNetappFilesTargetParams Specifies the params of the Netapp recovery target.
type RecoverNetappToNetappFilesTargetParams struct {
	NewSourceConfig NullableRecoverNetappToNetappFilesTargetParamsNewSourceConfig `json:"newSourceConfig,omitempty"`
	OriginalSourceConfig NullableRecoverNetappToNetappFilesTargetParamsOriginalSourceConfig `json:"originalSourceConfig,omitempty"`
	// Specifies the parameter whether the recovery should be performed to a new or the original Netapp target.
	RecoverToNewSource bool `json:"recoverToNewSource"`
}

type _RecoverNetappToNetappFilesTargetParams RecoverNetappToNetappFilesTargetParams

// NewRecoverNetappToNetappFilesTargetParams instantiates a new RecoverNetappToNetappFilesTargetParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverNetappToNetappFilesTargetParams(recoverToNewSource bool) *RecoverNetappToNetappFilesTargetParams {
	this := RecoverNetappToNetappFilesTargetParams{}
	this.RecoverToNewSource = recoverToNewSource
	return &this
}

// NewRecoverNetappToNetappFilesTargetParamsWithDefaults instantiates a new RecoverNetappToNetappFilesTargetParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverNetappToNetappFilesTargetParamsWithDefaults() *RecoverNetappToNetappFilesTargetParams {
	this := RecoverNetappToNetappFilesTargetParams{}
	return &this
}

// GetNewSourceConfig returns the NewSourceConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverNetappToNetappFilesTargetParams) GetNewSourceConfig() RecoverNetappToNetappFilesTargetParamsNewSourceConfig {
	if o == nil || IsNil(o.NewSourceConfig.Get()) {
		var ret RecoverNetappToNetappFilesTargetParamsNewSourceConfig
		return ret
	}
	return *o.NewSourceConfig.Get()
}

// GetNewSourceConfigOk returns a tuple with the NewSourceConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverNetappToNetappFilesTargetParams) GetNewSourceConfigOk() (*RecoverNetappToNetappFilesTargetParamsNewSourceConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.NewSourceConfig.Get(), o.NewSourceConfig.IsSet()
}

// HasNewSourceConfig returns a boolean if a field has been set.
func (o *RecoverNetappToNetappFilesTargetParams) HasNewSourceConfig() bool {
	if o != nil && o.NewSourceConfig.IsSet() {
		return true
	}

	return false
}

// SetNewSourceConfig gets a reference to the given NullableRecoverNetappToNetappFilesTargetParamsNewSourceConfig and assigns it to the NewSourceConfig field.
func (o *RecoverNetappToNetappFilesTargetParams) SetNewSourceConfig(v RecoverNetappToNetappFilesTargetParamsNewSourceConfig) {
	o.NewSourceConfig.Set(&v)
}
// SetNewSourceConfigNil sets the value for NewSourceConfig to be an explicit nil
func (o *RecoverNetappToNetappFilesTargetParams) SetNewSourceConfigNil() {
	o.NewSourceConfig.Set(nil)
}

// UnsetNewSourceConfig ensures that no value is present for NewSourceConfig, not even an explicit nil
func (o *RecoverNetappToNetappFilesTargetParams) UnsetNewSourceConfig() {
	o.NewSourceConfig.Unset()
}

// GetOriginalSourceConfig returns the OriginalSourceConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverNetappToNetappFilesTargetParams) GetOriginalSourceConfig() RecoverNetappToNetappFilesTargetParamsOriginalSourceConfig {
	if o == nil || IsNil(o.OriginalSourceConfig.Get()) {
		var ret RecoverNetappToNetappFilesTargetParamsOriginalSourceConfig
		return ret
	}
	return *o.OriginalSourceConfig.Get()
}

// GetOriginalSourceConfigOk returns a tuple with the OriginalSourceConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverNetappToNetappFilesTargetParams) GetOriginalSourceConfigOk() (*RecoverNetappToNetappFilesTargetParamsOriginalSourceConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.OriginalSourceConfig.Get(), o.OriginalSourceConfig.IsSet()
}

// HasOriginalSourceConfig returns a boolean if a field has been set.
func (o *RecoverNetappToNetappFilesTargetParams) HasOriginalSourceConfig() bool {
	if o != nil && o.OriginalSourceConfig.IsSet() {
		return true
	}

	return false
}

// SetOriginalSourceConfig gets a reference to the given NullableRecoverNetappToNetappFilesTargetParamsOriginalSourceConfig and assigns it to the OriginalSourceConfig field.
func (o *RecoverNetappToNetappFilesTargetParams) SetOriginalSourceConfig(v RecoverNetappToNetappFilesTargetParamsOriginalSourceConfig) {
	o.OriginalSourceConfig.Set(&v)
}
// SetOriginalSourceConfigNil sets the value for OriginalSourceConfig to be an explicit nil
func (o *RecoverNetappToNetappFilesTargetParams) SetOriginalSourceConfigNil() {
	o.OriginalSourceConfig.Set(nil)
}

// UnsetOriginalSourceConfig ensures that no value is present for OriginalSourceConfig, not even an explicit nil
func (o *RecoverNetappToNetappFilesTargetParams) UnsetOriginalSourceConfig() {
	o.OriginalSourceConfig.Unset()
}

// GetRecoverToNewSource returns the RecoverToNewSource field value
func (o *RecoverNetappToNetappFilesTargetParams) GetRecoverToNewSource() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RecoverToNewSource
}

// GetRecoverToNewSourceOk returns a tuple with the RecoverToNewSource field value
// and a boolean to check if the value has been set.
func (o *RecoverNetappToNetappFilesTargetParams) GetRecoverToNewSourceOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecoverToNewSource, true
}

// SetRecoverToNewSource sets field value
func (o *RecoverNetappToNetappFilesTargetParams) SetRecoverToNewSource(v bool) {
	o.RecoverToNewSource = v
}

func (o RecoverNetappToNetappFilesTargetParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecoverNetappToNetappFilesTargetParams) ToMap() (map[string]interface{}, error) {
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

func (o *RecoverNetappToNetappFilesTargetParams) UnmarshalJSON(data []byte) (err error) {
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

	varRecoverNetappToNetappFilesTargetParams := _RecoverNetappToNetappFilesTargetParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRecoverNetappToNetappFilesTargetParams)

	if err != nil {
		return err
	}

	*o = RecoverNetappToNetappFilesTargetParams(varRecoverNetappToNetappFilesTargetParams)

	return err
}

type NullableRecoverNetappToNetappFilesTargetParams struct {
	value *RecoverNetappToNetappFilesTargetParams
	isSet bool
}

func (v NullableRecoverNetappToNetappFilesTargetParams) Get() *RecoverNetappToNetappFilesTargetParams {
	return v.value
}

func (v *NullableRecoverNetappToNetappFilesTargetParams) Set(val *RecoverNetappToNetappFilesTargetParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverNetappToNetappFilesTargetParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverNetappToNetappFilesTargetParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverNetappToNetappFilesTargetParams(val *RecoverNetappToNetappFilesTargetParams) *NullableRecoverNetappToNetappFilesTargetParams {
	return &NullableRecoverNetappToNetappFilesTargetParams{value: val, isSet: true}
}

func (v NullableRecoverNetappToNetappFilesTargetParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverNetappToNetappFilesTargetParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


