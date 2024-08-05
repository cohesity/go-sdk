/*
Cohesity REST API

Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
)

// checks if the DataTieringAnalysisGroupRuns type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataTieringAnalysisGroupRuns{}

// DataTieringAnalysisGroupRuns Specifies the runs of a data tiering analysis group.
type DataTieringAnalysisGroupRuns struct {
	// Indicates whether the result is truncated due to hitting maximum size limit
	IsResponseTruncated NullableBool `json:"isResponseTruncated,omitempty"`
	// Specifies the data tiering analysis group runs.
	Runs []DataTieringAnalysisGroupRun `json:"runs,omitempty"`
}

// NewDataTieringAnalysisGroupRuns instantiates a new DataTieringAnalysisGroupRuns object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataTieringAnalysisGroupRuns() *DataTieringAnalysisGroupRuns {
	this := DataTieringAnalysisGroupRuns{}
	return &this
}

// NewDataTieringAnalysisGroupRunsWithDefaults instantiates a new DataTieringAnalysisGroupRuns object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataTieringAnalysisGroupRunsWithDefaults() *DataTieringAnalysisGroupRuns {
	this := DataTieringAnalysisGroupRuns{}
	return &this
}

// GetIsResponseTruncated returns the IsResponseTruncated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataTieringAnalysisGroupRuns) GetIsResponseTruncated() bool {
	if o == nil || IsNil(o.IsResponseTruncated.Get()) {
		var ret bool
		return ret
	}
	return *o.IsResponseTruncated.Get()
}

// GetIsResponseTruncatedOk returns a tuple with the IsResponseTruncated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataTieringAnalysisGroupRuns) GetIsResponseTruncatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsResponseTruncated.Get(), o.IsResponseTruncated.IsSet()
}

// HasIsResponseTruncated returns a boolean if a field has been set.
func (o *DataTieringAnalysisGroupRuns) HasIsResponseTruncated() bool {
	if o != nil && o.IsResponseTruncated.IsSet() {
		return true
	}

	return false
}

// SetIsResponseTruncated gets a reference to the given NullableBool and assigns it to the IsResponseTruncated field.
func (o *DataTieringAnalysisGroupRuns) SetIsResponseTruncated(v bool) {
	o.IsResponseTruncated.Set(&v)
}
// SetIsResponseTruncatedNil sets the value for IsResponseTruncated to be an explicit nil
func (o *DataTieringAnalysisGroupRuns) SetIsResponseTruncatedNil() {
	o.IsResponseTruncated.Set(nil)
}

// UnsetIsResponseTruncated ensures that no value is present for IsResponseTruncated, not even an explicit nil
func (o *DataTieringAnalysisGroupRuns) UnsetIsResponseTruncated() {
	o.IsResponseTruncated.Unset()
}

// GetRuns returns the Runs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataTieringAnalysisGroupRuns) GetRuns() []DataTieringAnalysisGroupRun {
	if o == nil {
		var ret []DataTieringAnalysisGroupRun
		return ret
	}
	return o.Runs
}

// GetRunsOk returns a tuple with the Runs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataTieringAnalysisGroupRuns) GetRunsOk() ([]DataTieringAnalysisGroupRun, bool) {
	if o == nil || IsNil(o.Runs) {
		return nil, false
	}
	return o.Runs, true
}

// HasRuns returns a boolean if a field has been set.
func (o *DataTieringAnalysisGroupRuns) HasRuns() bool {
	if o != nil && !IsNil(o.Runs) {
		return true
	}

	return false
}

// SetRuns gets a reference to the given []DataTieringAnalysisGroupRun and assigns it to the Runs field.
func (o *DataTieringAnalysisGroupRuns) SetRuns(v []DataTieringAnalysisGroupRun) {
	o.Runs = v
}

func (o DataTieringAnalysisGroupRuns) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataTieringAnalysisGroupRuns) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.IsResponseTruncated.IsSet() {
		toSerialize["isResponseTruncated"] = o.IsResponseTruncated.Get()
	}
	if o.Runs != nil {
		toSerialize["runs"] = o.Runs
	}
	return toSerialize, nil
}

type NullableDataTieringAnalysisGroupRuns struct {
	value *DataTieringAnalysisGroupRuns
	isSet bool
}

func (v NullableDataTieringAnalysisGroupRuns) Get() *DataTieringAnalysisGroupRuns {
	return v.value
}

func (v *NullableDataTieringAnalysisGroupRuns) Set(val *DataTieringAnalysisGroupRuns) {
	v.value = val
	v.isSet = true
}

func (v NullableDataTieringAnalysisGroupRuns) IsSet() bool {
	return v.isSet
}

func (v *NullableDataTieringAnalysisGroupRuns) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataTieringAnalysisGroupRuns(val *DataTieringAnalysisGroupRuns) *NullableDataTieringAnalysisGroupRuns {
	return &NullableDataTieringAnalysisGroupRuns{value: val, isSet: true}
}

func (v NullableDataTieringAnalysisGroupRuns) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataTieringAnalysisGroupRuns) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


