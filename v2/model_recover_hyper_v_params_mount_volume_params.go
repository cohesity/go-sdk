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

// checks if the RecoverHyperVParamsMountVolumeParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecoverHyperVParamsMountVolumeParams{}

// RecoverHyperVParamsMountVolumeParams Specifies the parameters to mount HyperV Volumes.
type RecoverHyperVParamsMountVolumeParams struct {
	HypervTargetParams NullableMountHyperVVolumeParamsHypervTargetParams `json:"hypervTargetParams,omitempty"`
	// Specifies the environment of the recovery target. The corresponding params below must be filled out.
	TargetEnvironment string `json:"targetEnvironment"`
}

type _RecoverHyperVParamsMountVolumeParams RecoverHyperVParamsMountVolumeParams

// NewRecoverHyperVParamsMountVolumeParams instantiates a new RecoverHyperVParamsMountVolumeParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverHyperVParamsMountVolumeParams(targetEnvironment string) *RecoverHyperVParamsMountVolumeParams {
	this := RecoverHyperVParamsMountVolumeParams{}
	this.TargetEnvironment = targetEnvironment
	return &this
}

// NewRecoverHyperVParamsMountVolumeParamsWithDefaults instantiates a new RecoverHyperVParamsMountVolumeParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverHyperVParamsMountVolumeParamsWithDefaults() *RecoverHyperVParamsMountVolumeParams {
	this := RecoverHyperVParamsMountVolumeParams{}
	return &this
}

// GetHypervTargetParams returns the HypervTargetParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverHyperVParamsMountVolumeParams) GetHypervTargetParams() MountHyperVVolumeParamsHypervTargetParams {
	if o == nil || IsNil(o.HypervTargetParams.Get()) {
		var ret MountHyperVVolumeParamsHypervTargetParams
		return ret
	}
	return *o.HypervTargetParams.Get()
}

// GetHypervTargetParamsOk returns a tuple with the HypervTargetParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverHyperVParamsMountVolumeParams) GetHypervTargetParamsOk() (*MountHyperVVolumeParamsHypervTargetParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.HypervTargetParams.Get(), o.HypervTargetParams.IsSet()
}

// HasHypervTargetParams returns a boolean if a field has been set.
func (o *RecoverHyperVParamsMountVolumeParams) HasHypervTargetParams() bool {
	if o != nil && o.HypervTargetParams.IsSet() {
		return true
	}

	return false
}

// SetHypervTargetParams gets a reference to the given NullableMountHyperVVolumeParamsHypervTargetParams and assigns it to the HypervTargetParams field.
func (o *RecoverHyperVParamsMountVolumeParams) SetHypervTargetParams(v MountHyperVVolumeParamsHypervTargetParams) {
	o.HypervTargetParams.Set(&v)
}
// SetHypervTargetParamsNil sets the value for HypervTargetParams to be an explicit nil
func (o *RecoverHyperVParamsMountVolumeParams) SetHypervTargetParamsNil() {
	o.HypervTargetParams.Set(nil)
}

// UnsetHypervTargetParams ensures that no value is present for HypervTargetParams, not even an explicit nil
func (o *RecoverHyperVParamsMountVolumeParams) UnsetHypervTargetParams() {
	o.HypervTargetParams.Unset()
}

// GetTargetEnvironment returns the TargetEnvironment field value
func (o *RecoverHyperVParamsMountVolumeParams) GetTargetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetEnvironment
}

// GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field value
// and a boolean to check if the value has been set.
func (o *RecoverHyperVParamsMountVolumeParams) GetTargetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetEnvironment, true
}

// SetTargetEnvironment sets field value
func (o *RecoverHyperVParamsMountVolumeParams) SetTargetEnvironment(v string) {
	o.TargetEnvironment = v
}

func (o RecoverHyperVParamsMountVolumeParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecoverHyperVParamsMountVolumeParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.HypervTargetParams.IsSet() {
		toSerialize["hypervTargetParams"] = o.HypervTargetParams.Get()
	}
	toSerialize["targetEnvironment"] = o.TargetEnvironment
	return toSerialize, nil
}

func (o *RecoverHyperVParamsMountVolumeParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"targetEnvironment",
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

	varRecoverHyperVParamsMountVolumeParams := _RecoverHyperVParamsMountVolumeParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRecoverHyperVParamsMountVolumeParams)

	if err != nil {
		return err
	}

	*o = RecoverHyperVParamsMountVolumeParams(varRecoverHyperVParamsMountVolumeParams)

	return err
}

type NullableRecoverHyperVParamsMountVolumeParams struct {
	value *RecoverHyperVParamsMountVolumeParams
	isSet bool
}

func (v NullableRecoverHyperVParamsMountVolumeParams) Get() *RecoverHyperVParamsMountVolumeParams {
	return v.value
}

func (v *NullableRecoverHyperVParamsMountVolumeParams) Set(val *RecoverHyperVParamsMountVolumeParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverHyperVParamsMountVolumeParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverHyperVParamsMountVolumeParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverHyperVParamsMountVolumeParams(val *RecoverHyperVParamsMountVolumeParams) *NullableRecoverHyperVParamsMountVolumeParams {
	return &NullableRecoverHyperVParamsMountVolumeParams{value: val, isSet: true}
}

func (v NullableRecoverHyperVParamsMountVolumeParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverHyperVParamsMountVolumeParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


