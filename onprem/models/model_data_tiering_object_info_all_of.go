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

// DataTieringObjectInfoAllOf struct for DataTieringObjectInfoAllOf
type DataTieringObjectInfoAllOf struct {
	AnalysisInfo *DataTieringObjectAnalysisInfo `json:"analysisInfo,omitempty"`
}

// NewDataTieringObjectInfoAllOf instantiates a new DataTieringObjectInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataTieringObjectInfoAllOf() *DataTieringObjectInfoAllOf {
	this := DataTieringObjectInfoAllOf{}
	return &this
}

// NewDataTieringObjectInfoAllOfWithDefaults instantiates a new DataTieringObjectInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataTieringObjectInfoAllOfWithDefaults() *DataTieringObjectInfoAllOf {
	this := DataTieringObjectInfoAllOf{}
	return &this
}

// GetAnalysisInfo returns the AnalysisInfo field value if set, zero value otherwise.
func (o *DataTieringObjectInfoAllOf) GetAnalysisInfo() DataTieringObjectAnalysisInfo {
	if o == nil || o.AnalysisInfo == nil {
		var ret DataTieringObjectAnalysisInfo
		return ret
	}
	return *o.AnalysisInfo
}

// GetAnalysisInfoOk returns a tuple with the AnalysisInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataTieringObjectInfoAllOf) GetAnalysisInfoOk() (*DataTieringObjectAnalysisInfo, bool) {
	if o == nil || o.AnalysisInfo == nil {
		return nil, false
	}
	return o.AnalysisInfo, true
}

// HasAnalysisInfo returns a boolean if a field has been set.
func (o *DataTieringObjectInfoAllOf) HasAnalysisInfo() bool {
	if o != nil && o.AnalysisInfo != nil {
		return true
	}

	return false
}

// SetAnalysisInfo gets a reference to the given DataTieringObjectAnalysisInfo and assigns it to the AnalysisInfo field.
func (o *DataTieringObjectInfoAllOf) SetAnalysisInfo(v DataTieringObjectAnalysisInfo) {
	o.AnalysisInfo = &v
}

func (o DataTieringObjectInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AnalysisInfo != nil {
		toSerialize["analysisInfo"] = o.AnalysisInfo
	}
	return json.Marshal(toSerialize)
}

type NullableDataTieringObjectInfoAllOf struct {
	value *DataTieringObjectInfoAllOf
	isSet bool
}

func (v NullableDataTieringObjectInfoAllOf) Get() *DataTieringObjectInfoAllOf {
	return v.value
}

func (v *NullableDataTieringObjectInfoAllOf) Set(val *DataTieringObjectInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDataTieringObjectInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDataTieringObjectInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataTieringObjectInfoAllOf(val *DataTieringObjectInfoAllOf) *NullableDataTieringObjectInfoAllOf {
	return &NullableDataTieringObjectInfoAllOf{value: val, isSet: true}
}

func (v NullableDataTieringObjectInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataTieringObjectInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o DataTieringObjectInfoAllOf) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}