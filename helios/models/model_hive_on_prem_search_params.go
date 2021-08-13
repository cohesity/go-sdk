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

// HiveOnPremSearchParams Parameters required to search Hive on a cluster.
type HiveOnPremSearchParams struct {
	// Specifies the search string to search the Hive Objects
	SearchString NullableString `json:"searchString"`
	// Specifies one or more Hive object types be searched.
	HiveObjectTypes []string `json:"hiveObjectTypes"`
	// Specifies a list of source ids. Only files found in these sources will be returned.
	SourceIds []int64 `json:"sourceIds,omitempty"`
}

// NewHiveOnPremSearchParams instantiates a new HiveOnPremSearchParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHiveOnPremSearchParams(searchString NullableString, hiveObjectTypes []string) *HiveOnPremSearchParams {
	this := HiveOnPremSearchParams{}
	this.SearchString = searchString
	this.HiveObjectTypes = hiveObjectTypes
	return &this
}

// NewHiveOnPremSearchParamsWithDefaults instantiates a new HiveOnPremSearchParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHiveOnPremSearchParamsWithDefaults() *HiveOnPremSearchParams {
	this := HiveOnPremSearchParams{}
	return &this
}

// GetSearchString returns the SearchString field value
// If the value is explicit nil, the zero value for string will be returned
func (o *HiveOnPremSearchParams) GetSearchString() string {
	if o == nil || o.SearchString.Get() == nil {
		var ret string
		return ret
	}

	return *o.SearchString.Get()
}

// GetSearchStringOk returns a tuple with the SearchString field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HiveOnPremSearchParams) GetSearchStringOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SearchString.Get(), o.SearchString.IsSet()
}

// SetSearchString sets field value
func (o *HiveOnPremSearchParams) SetSearchString(v string) {
	o.SearchString.Set(&v)
}

// GetHiveObjectTypes returns the HiveObjectTypes field value
func (o *HiveOnPremSearchParams) GetHiveObjectTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.HiveObjectTypes
}

// GetHiveObjectTypesOk returns a tuple with the HiveObjectTypes field value
// and a boolean to check if the value has been set.
func (o *HiveOnPremSearchParams) GetHiveObjectTypesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HiveObjectTypes, true
}

// SetHiveObjectTypes sets field value
func (o *HiveOnPremSearchParams) SetHiveObjectTypes(v []string) {
	o.HiveObjectTypes = v
}

// GetSourceIds returns the SourceIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HiveOnPremSearchParams) GetSourceIds() []int64 {
	if o == nil  {
		var ret []int64
		return ret
	}
	return o.SourceIds
}

// GetSourceIdsOk returns a tuple with the SourceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HiveOnPremSearchParams) GetSourceIdsOk() (*[]int64, bool) {
	if o == nil || o.SourceIds == nil {
		return nil, false
	}
	return &o.SourceIds, true
}

// HasSourceIds returns a boolean if a field has been set.
func (o *HiveOnPremSearchParams) HasSourceIds() bool {
	if o != nil && o.SourceIds != nil {
		return true
	}

	return false
}

// SetSourceIds gets a reference to the given []int64 and assigns it to the SourceIds field.
func (o *HiveOnPremSearchParams) SetSourceIds(v []int64) {
	o.SourceIds = v
}

func (o HiveOnPremSearchParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["searchString"] = o.SearchString.Get()
	}
	if true {
		toSerialize["hiveObjectTypes"] = o.HiveObjectTypes
	}
	if o.SourceIds != nil {
		toSerialize["sourceIds"] = o.SourceIds
	}
	return json.Marshal(toSerialize)
}

type NullableHiveOnPremSearchParams struct {
	value *HiveOnPremSearchParams
	isSet bool
}

func (v NullableHiveOnPremSearchParams) Get() *HiveOnPremSearchParams {
	return v.value
}

func (v *NullableHiveOnPremSearchParams) Set(val *HiveOnPremSearchParams) {
	v.value = val
	v.isSet = true
}

func (v NullableHiveOnPremSearchParams) IsSet() bool {
	return v.isSet
}

func (v *NullableHiveOnPremSearchParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHiveOnPremSearchParams(val *HiveOnPremSearchParams) *NullableHiveOnPremSearchParams {
	return &NullableHiveOnPremSearchParams{value: val, isSet: true}
}

func (v NullableHiveOnPremSearchParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHiveOnPremSearchParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


