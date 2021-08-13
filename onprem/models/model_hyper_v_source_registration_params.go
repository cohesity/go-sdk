/*
 * Cohesity REST API
 *
 * Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

// package onprem
package models

import (
	"encoding/json"
	. "github.com/cohesity/go-sdk/onprem/utils"
	"fmt"
)

var _ = NullableBool{}

// HyperVSourceRegistrationParams Specifies the paramaters to register a HyperV source.
type HyperVSourceRegistrationParams struct {
	// Specifies the HyperV Source type.
	Type NullableString `json:"type"`
	ScvmmParams *ScvmmRegistrationParams `json:"scvmmParams,omitempty"`
	StandaloneHostParams *StandaloneHostRegistrationParams `json:"standaloneHostParams,omitempty"`
	StandaloneClusterParams *StandaloneClusterRegistrationParams `json:"standaloneClusterParams,omitempty"`
}

// NewHyperVSourceRegistrationParams instantiates a new HyperVSourceRegistrationParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperVSourceRegistrationParams(type_ NullableString) *HyperVSourceRegistrationParams {
	this := HyperVSourceRegistrationParams{}
	this.Type = type_
	return &this
}

// NewHyperVSourceRegistrationParamsWithDefaults instantiates a new HyperVSourceRegistrationParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperVSourceRegistrationParamsWithDefaults() *HyperVSourceRegistrationParams {
	this := HyperVSourceRegistrationParams{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *HyperVSourceRegistrationParams) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}

	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperVSourceRegistrationParams) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// SetType sets field value
func (o *HyperVSourceRegistrationParams) SetType(v string) {
	o.Type.Set(&v)
}

// GetScvmmParams returns the ScvmmParams field value if set, zero value otherwise.
func (o *HyperVSourceRegistrationParams) GetScvmmParams() ScvmmRegistrationParams {
	if o == nil || o.ScvmmParams == nil {
		var ret ScvmmRegistrationParams
		return ret
	}
	return *o.ScvmmParams
}

// GetScvmmParamsOk returns a tuple with the ScvmmParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperVSourceRegistrationParams) GetScvmmParamsOk() (*ScvmmRegistrationParams, bool) {
	if o == nil || o.ScvmmParams == nil {
		return nil, false
	}
	return o.ScvmmParams, true
}

// HasScvmmParams returns a boolean if a field has been set.
func (o *HyperVSourceRegistrationParams) HasScvmmParams() bool {
	if o != nil && o.ScvmmParams != nil {
		return true
	}

	return false
}

// SetScvmmParams gets a reference to the given ScvmmRegistrationParams and assigns it to the ScvmmParams field.
func (o *HyperVSourceRegistrationParams) SetScvmmParams(v ScvmmRegistrationParams) {
	o.ScvmmParams = &v
}

// GetStandaloneHostParams returns the StandaloneHostParams field value if set, zero value otherwise.
func (o *HyperVSourceRegistrationParams) GetStandaloneHostParams() StandaloneHostRegistrationParams {
	if o == nil || o.StandaloneHostParams == nil {
		var ret StandaloneHostRegistrationParams
		return ret
	}
	return *o.StandaloneHostParams
}

// GetStandaloneHostParamsOk returns a tuple with the StandaloneHostParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperVSourceRegistrationParams) GetStandaloneHostParamsOk() (*StandaloneHostRegistrationParams, bool) {
	if o == nil || o.StandaloneHostParams == nil {
		return nil, false
	}
	return o.StandaloneHostParams, true
}

// HasStandaloneHostParams returns a boolean if a field has been set.
func (o *HyperVSourceRegistrationParams) HasStandaloneHostParams() bool {
	if o != nil && o.StandaloneHostParams != nil {
		return true
	}

	return false
}

// SetStandaloneHostParams gets a reference to the given StandaloneHostRegistrationParams and assigns it to the StandaloneHostParams field.
func (o *HyperVSourceRegistrationParams) SetStandaloneHostParams(v StandaloneHostRegistrationParams) {
	o.StandaloneHostParams = &v
}

// GetStandaloneClusterParams returns the StandaloneClusterParams field value if set, zero value otherwise.
func (o *HyperVSourceRegistrationParams) GetStandaloneClusterParams() StandaloneClusterRegistrationParams {
	if o == nil || o.StandaloneClusterParams == nil {
		var ret StandaloneClusterRegistrationParams
		return ret
	}
	return *o.StandaloneClusterParams
}

// GetStandaloneClusterParamsOk returns a tuple with the StandaloneClusterParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperVSourceRegistrationParams) GetStandaloneClusterParamsOk() (*StandaloneClusterRegistrationParams, bool) {
	if o == nil || o.StandaloneClusterParams == nil {
		return nil, false
	}
	return o.StandaloneClusterParams, true
}

// HasStandaloneClusterParams returns a boolean if a field has been set.
func (o *HyperVSourceRegistrationParams) HasStandaloneClusterParams() bool {
	if o != nil && o.StandaloneClusterParams != nil {
		return true
	}

	return false
}

// SetStandaloneClusterParams gets a reference to the given StandaloneClusterRegistrationParams and assigns it to the StandaloneClusterParams field.
func (o *HyperVSourceRegistrationParams) SetStandaloneClusterParams(v StandaloneClusterRegistrationParams) {
	o.StandaloneClusterParams = &v
}

func (o HyperVSourceRegistrationParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type.Get()
	}
	if o.ScvmmParams != nil {
		toSerialize["scvmmParams"] = o.ScvmmParams
	}
	if o.StandaloneHostParams != nil {
		toSerialize["standaloneHostParams"] = o.StandaloneHostParams
	}
	if o.StandaloneClusterParams != nil {
		toSerialize["standaloneClusterParams"] = o.StandaloneClusterParams
	}
	return json.Marshal(toSerialize)
}

type NullableHyperVSourceRegistrationParams struct {
	value *HyperVSourceRegistrationParams
	isSet bool
}

func (v NullableHyperVSourceRegistrationParams) Get() *HyperVSourceRegistrationParams {
	return v.value
}

func (v *NullableHyperVSourceRegistrationParams) Set(val *HyperVSourceRegistrationParams) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperVSourceRegistrationParams) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperVSourceRegistrationParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperVSourceRegistrationParams(val *HyperVSourceRegistrationParams) *NullableHyperVSourceRegistrationParams {
	return &NullableHyperVSourceRegistrationParams{value: val, isSet: true}
}

func (v NullableHyperVSourceRegistrationParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperVSourceRegistrationParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o HyperVSourceRegistrationParams) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}