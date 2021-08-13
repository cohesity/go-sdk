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

// ArchivalRunSummary Specifies summary information about archival run.
type ArchivalRunSummary struct {
	// Archival results for each archival target.
	ArchivalTargetResults *[]ArchivalTargetResult `json:"archivalTargetResults,omitempty"`
}

// NewArchivalRunSummary instantiates a new ArchivalRunSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArchivalRunSummary() *ArchivalRunSummary {
	this := ArchivalRunSummary{}
	return &this
}

// NewArchivalRunSummaryWithDefaults instantiates a new ArchivalRunSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArchivalRunSummaryWithDefaults() *ArchivalRunSummary {
	this := ArchivalRunSummary{}
	return &this
}

// GetArchivalTargetResults returns the ArchivalTargetResults field value if set, zero value otherwise.
func (o *ArchivalRunSummary) GetArchivalTargetResults() []ArchivalTargetResult {
	if o == nil || o.ArchivalTargetResults == nil {
		var ret []ArchivalTargetResult
		return ret
	}
	return *o.ArchivalTargetResults
}

// GetArchivalTargetResultsOk returns a tuple with the ArchivalTargetResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchivalRunSummary) GetArchivalTargetResultsOk() (*[]ArchivalTargetResult, bool) {
	if o == nil || o.ArchivalTargetResults == nil {
		return nil, false
	}
	return o.ArchivalTargetResults, true
}

// HasArchivalTargetResults returns a boolean if a field has been set.
func (o *ArchivalRunSummary) HasArchivalTargetResults() bool {
	if o != nil && o.ArchivalTargetResults != nil {
		return true
	}

	return false
}

// SetArchivalTargetResults gets a reference to the given []ArchivalTargetResult and assigns it to the ArchivalTargetResults field.
func (o *ArchivalRunSummary) SetArchivalTargetResults(v []ArchivalTargetResult) {
	o.ArchivalTargetResults = &v
}

func (o ArchivalRunSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ArchivalTargetResults != nil {
		toSerialize["archivalTargetResults"] = o.ArchivalTargetResults
	}
	return json.Marshal(toSerialize)
}

type NullableArchivalRunSummary struct {
	value *ArchivalRunSummary
	isSet bool
}

func (v NullableArchivalRunSummary) Get() *ArchivalRunSummary {
	return v.value
}

func (v *NullableArchivalRunSummary) Set(val *ArchivalRunSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableArchivalRunSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableArchivalRunSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArchivalRunSummary(val *ArchivalRunSummary) *NullableArchivalRunSummary {
	return &NullableArchivalRunSummary{value: val, isSet: true}
}

func (v NullableArchivalRunSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArchivalRunSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o ArchivalRunSummary) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}