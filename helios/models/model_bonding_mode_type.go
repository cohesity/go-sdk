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

// BondingModeType Type of bonding mode.
type BondingModeType struct {
	// Specifies the bonding mode type.
	BondingModeType *string `json:"bondingModeType,omitempty"`
}

// NewBondingModeType instantiates a new BondingModeType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBondingModeType() *BondingModeType {
	this := BondingModeType{}
	return &this
}

// NewBondingModeTypeWithDefaults instantiates a new BondingModeType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBondingModeTypeWithDefaults() *BondingModeType {
	this := BondingModeType{}
	return &this
}

// GetBondingModeType returns the BondingModeType field value if set, zero value otherwise.
func (o *BondingModeType) GetBondingModeType() string {
	if o == nil || o.BondingModeType == nil {
		var ret string
		return ret
	}
	return *o.BondingModeType
}

// GetBondingModeTypeOk returns a tuple with the BondingModeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BondingModeType) GetBondingModeTypeOk() (*string, bool) {
	if o == nil || o.BondingModeType == nil {
		return nil, false
	}
	return o.BondingModeType, true
}

// HasBondingModeType returns a boolean if a field has been set.
func (o *BondingModeType) HasBondingModeType() bool {
	if o != nil && o.BondingModeType != nil {
		return true
	}

	return false
}

// SetBondingModeType gets a reference to the given string and assigns it to the BondingModeType field.
func (o *BondingModeType) SetBondingModeType(v string) {
	o.BondingModeType = &v
}

func (o BondingModeType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BondingModeType != nil {
		toSerialize["bondingModeType"] = o.BondingModeType
	}
	return json.Marshal(toSerialize)
}

type NullableBondingModeType struct {
	value *BondingModeType
	isSet bool
}

func (v NullableBondingModeType) Get() *BondingModeType {
	return v.value
}

func (v *NullableBondingModeType) Set(val *BondingModeType) {
	v.value = val
	v.isSet = true
}

func (v NullableBondingModeType) IsSet() bool {
	return v.isSet
}

func (v *NullableBondingModeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBondingModeType(val *BondingModeType) *NullableBondingModeType {
	return &NullableBondingModeType{value: val, isSet: true}
}

func (v NullableBondingModeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBondingModeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


