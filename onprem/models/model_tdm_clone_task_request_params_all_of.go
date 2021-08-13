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

// TdmCloneTaskRequestParamsAllOf struct for TdmCloneTaskRequestParamsAllOf
type TdmCloneTaskRequestParamsAllOf struct {
	OracleParams *OracleCloneTask `json:"oracleParams,omitempty"`
}

// NewTdmCloneTaskRequestParamsAllOf instantiates a new TdmCloneTaskRequestParamsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTdmCloneTaskRequestParamsAllOf() *TdmCloneTaskRequestParamsAllOf {
	this := TdmCloneTaskRequestParamsAllOf{}
	return &this
}

// NewTdmCloneTaskRequestParamsAllOfWithDefaults instantiates a new TdmCloneTaskRequestParamsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTdmCloneTaskRequestParamsAllOfWithDefaults() *TdmCloneTaskRequestParamsAllOf {
	this := TdmCloneTaskRequestParamsAllOf{}
	return &this
}

// GetOracleParams returns the OracleParams field value if set, zero value otherwise.
func (o *TdmCloneTaskRequestParamsAllOf) GetOracleParams() OracleCloneTask {
	if o == nil || o.OracleParams == nil {
		var ret OracleCloneTask
		return ret
	}
	return *o.OracleParams
}

// GetOracleParamsOk returns a tuple with the OracleParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TdmCloneTaskRequestParamsAllOf) GetOracleParamsOk() (*OracleCloneTask, bool) {
	if o == nil || o.OracleParams == nil {
		return nil, false
	}
	return o.OracleParams, true
}

// HasOracleParams returns a boolean if a field has been set.
func (o *TdmCloneTaskRequestParamsAllOf) HasOracleParams() bool {
	if o != nil && o.OracleParams != nil {
		return true
	}

	return false
}

// SetOracleParams gets a reference to the given OracleCloneTask and assigns it to the OracleParams field.
func (o *TdmCloneTaskRequestParamsAllOf) SetOracleParams(v OracleCloneTask) {
	o.OracleParams = &v
}

func (o TdmCloneTaskRequestParamsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OracleParams != nil {
		toSerialize["oracleParams"] = o.OracleParams
	}
	return json.Marshal(toSerialize)
}

type NullableTdmCloneTaskRequestParamsAllOf struct {
	value *TdmCloneTaskRequestParamsAllOf
	isSet bool
}

func (v NullableTdmCloneTaskRequestParamsAllOf) Get() *TdmCloneTaskRequestParamsAllOf {
	return v.value
}

func (v *NullableTdmCloneTaskRequestParamsAllOf) Set(val *TdmCloneTaskRequestParamsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTdmCloneTaskRequestParamsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTdmCloneTaskRequestParamsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTdmCloneTaskRequestParamsAllOf(val *TdmCloneTaskRequestParamsAllOf) *NullableTdmCloneTaskRequestParamsAllOf {
	return &NullableTdmCloneTaskRequestParamsAllOf{value: val, isSet: true}
}

func (v NullableTdmCloneTaskRequestParamsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTdmCloneTaskRequestParamsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o TdmCloneTaskRequestParamsAllOf) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}