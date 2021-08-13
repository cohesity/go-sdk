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

// OracleTierAllOf struct for OracleTierAllOf
type OracleTierAllOf struct {
	// Specifies the Oracle tier types.
	TierType NullableString `json:"tierType"`
}

// NewOracleTierAllOf instantiates a new OracleTierAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOracleTierAllOf(tierType NullableString) *OracleTierAllOf {
	this := OracleTierAllOf{}
	this.TierType = tierType
	return &this
}

// NewOracleTierAllOfWithDefaults instantiates a new OracleTierAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOracleTierAllOfWithDefaults() *OracleTierAllOf {
	this := OracleTierAllOf{}
	return &this
}

// GetTierType returns the TierType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *OracleTierAllOf) GetTierType() string {
	if o == nil || o.TierType.Get() == nil {
		var ret string
		return ret
	}

	return *o.TierType.Get()
}

// GetTierTypeOk returns a tuple with the TierType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OracleTierAllOf) GetTierTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TierType.Get(), o.TierType.IsSet()
}

// SetTierType sets field value
func (o *OracleTierAllOf) SetTierType(v string) {
	o.TierType.Set(&v)
}

func (o OracleTierAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["tierType"] = o.TierType.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableOracleTierAllOf struct {
	value *OracleTierAllOf
	isSet bool
}

func (v NullableOracleTierAllOf) Get() *OracleTierAllOf {
	return v.value
}

func (v *NullableOracleTierAllOf) Set(val *OracleTierAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOracleTierAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOracleTierAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOracleTierAllOf(val *OracleTierAllOf) *NullableOracleTierAllOf {
	return &NullableOracleTierAllOf{value: val, isSet: true}
}

func (v NullableOracleTierAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOracleTierAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


