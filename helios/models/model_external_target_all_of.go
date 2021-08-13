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

// ExternalTargetAllOf struct for ExternalTargetAllOf
type ExternalTargetAllOf struct {
	ArchivalParams *ArchivalExternalTargetParams `json:"archivalParams,omitempty"`
	TieringParams *TieringExternalTargetParams `json:"tieringParams,omitempty"`
}

// NewExternalTargetAllOf instantiates a new ExternalTargetAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalTargetAllOf() *ExternalTargetAllOf {
	this := ExternalTargetAllOf{}
	return &this
}

// NewExternalTargetAllOfWithDefaults instantiates a new ExternalTargetAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalTargetAllOfWithDefaults() *ExternalTargetAllOf {
	this := ExternalTargetAllOf{}
	return &this
}

// GetArchivalParams returns the ArchivalParams field value if set, zero value otherwise.
func (o *ExternalTargetAllOf) GetArchivalParams() ArchivalExternalTargetParams {
	if o == nil || o.ArchivalParams == nil {
		var ret ArchivalExternalTargetParams
		return ret
	}
	return *o.ArchivalParams
}

// GetArchivalParamsOk returns a tuple with the ArchivalParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalTargetAllOf) GetArchivalParamsOk() (*ArchivalExternalTargetParams, bool) {
	if o == nil || o.ArchivalParams == nil {
		return nil, false
	}
	return o.ArchivalParams, true
}

// HasArchivalParams returns a boolean if a field has been set.
func (o *ExternalTargetAllOf) HasArchivalParams() bool {
	if o != nil && o.ArchivalParams != nil {
		return true
	}

	return false
}

// SetArchivalParams gets a reference to the given ArchivalExternalTargetParams and assigns it to the ArchivalParams field.
func (o *ExternalTargetAllOf) SetArchivalParams(v ArchivalExternalTargetParams) {
	o.ArchivalParams = &v
}

// GetTieringParams returns the TieringParams field value if set, zero value otherwise.
func (o *ExternalTargetAllOf) GetTieringParams() TieringExternalTargetParams {
	if o == nil || o.TieringParams == nil {
		var ret TieringExternalTargetParams
		return ret
	}
	return *o.TieringParams
}

// GetTieringParamsOk returns a tuple with the TieringParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalTargetAllOf) GetTieringParamsOk() (*TieringExternalTargetParams, bool) {
	if o == nil || o.TieringParams == nil {
		return nil, false
	}
	return o.TieringParams, true
}

// HasTieringParams returns a boolean if a field has been set.
func (o *ExternalTargetAllOf) HasTieringParams() bool {
	if o != nil && o.TieringParams != nil {
		return true
	}

	return false
}

// SetTieringParams gets a reference to the given TieringExternalTargetParams and assigns it to the TieringParams field.
func (o *ExternalTargetAllOf) SetTieringParams(v TieringExternalTargetParams) {
	o.TieringParams = &v
}

func (o ExternalTargetAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ArchivalParams != nil {
		toSerialize["archivalParams"] = o.ArchivalParams
	}
	if o.TieringParams != nil {
		toSerialize["tieringParams"] = o.TieringParams
	}
	return json.Marshal(toSerialize)
}

type NullableExternalTargetAllOf struct {
	value *ExternalTargetAllOf
	isSet bool
}

func (v NullableExternalTargetAllOf) Get() *ExternalTargetAllOf {
	return v.value
}

func (v *NullableExternalTargetAllOf) Set(val *ExternalTargetAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalTargetAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalTargetAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalTargetAllOf(val *ExternalTargetAllOf) *NullableExternalTargetAllOf {
	return &NullableExternalTargetAllOf{value: val, isSet: true}
}

func (v NullableExternalTargetAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalTargetAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


