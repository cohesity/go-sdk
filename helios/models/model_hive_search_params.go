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

// HiveSearchParams Specifies the parameters which are specific for searching Hive objects.
type HiveSearchParams struct {
	// Specifies the search string to search the Hive Objects
	SearchString NullableString `json:"searchString"`
	// Specifies one or more Hive object types be searched.
	HiveObjectTypes []string `json:"hiveObjectTypes"`
}

// NewHiveSearchParams instantiates a new HiveSearchParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHiveSearchParams(searchString NullableString, hiveObjectTypes []string) *HiveSearchParams {
	this := HiveSearchParams{}
	this.SearchString = searchString
	this.HiveObjectTypes = hiveObjectTypes
	return &this
}

// NewHiveSearchParamsWithDefaults instantiates a new HiveSearchParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHiveSearchParamsWithDefaults() *HiveSearchParams {
	this := HiveSearchParams{}
	return &this
}

// GetSearchString returns the SearchString field value
// If the value is explicit nil, the zero value for string will be returned
func (o *HiveSearchParams) GetSearchString() string {
	if o == nil || o.SearchString.Get() == nil {
		var ret string
		return ret
	}

	return *o.SearchString.Get()
}

// GetSearchStringOk returns a tuple with the SearchString field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HiveSearchParams) GetSearchStringOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SearchString.Get(), o.SearchString.IsSet()
}

// SetSearchString sets field value
func (o *HiveSearchParams) SetSearchString(v string) {
	o.SearchString.Set(&v)
}

// GetHiveObjectTypes returns the HiveObjectTypes field value
func (o *HiveSearchParams) GetHiveObjectTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.HiveObjectTypes
}

// GetHiveObjectTypesOk returns a tuple with the HiveObjectTypes field value
// and a boolean to check if the value has been set.
func (o *HiveSearchParams) GetHiveObjectTypesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HiveObjectTypes, true
}

// SetHiveObjectTypes sets field value
func (o *HiveSearchParams) SetHiveObjectTypes(v []string) {
	o.HiveObjectTypes = v
}

func (o HiveSearchParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["searchString"] = o.SearchString.Get()
	}
	if true {
		toSerialize["hiveObjectTypes"] = o.HiveObjectTypes
	}
	return json.Marshal(toSerialize)
}

type NullableHiveSearchParams struct {
	value *HiveSearchParams
	isSet bool
}

func (v NullableHiveSearchParams) Get() *HiveSearchParams {
	return v.value
}

func (v *NullableHiveSearchParams) Set(val *HiveSearchParams) {
	v.value = val
	v.isSet = true
}

func (v NullableHiveSearchParams) IsSet() bool {
	return v.isSet
}

func (v *NullableHiveSearchParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHiveSearchParams(val *HiveSearchParams) *NullableHiveSearchParams {
	return &NullableHiveSearchParams{value: val, isSet: true}
}

func (v NullableHiveSearchParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHiveSearchParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


