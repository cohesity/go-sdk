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

// checks if the WorkloadStatsSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkloadStatsSummary{}

// WorkloadStatsSummary Specifies the Workload Stats Summary Schema.
type WorkloadStatsSummary struct {
	// Specifies the Workload types.
	Workloads []WorkloadStatsSchema `json:"workloads,omitempty"`
}

// NewWorkloadStatsSummary instantiates a new WorkloadStatsSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkloadStatsSummary() *WorkloadStatsSummary {
	this := WorkloadStatsSummary{}
	return &this
}

// NewWorkloadStatsSummaryWithDefaults instantiates a new WorkloadStatsSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkloadStatsSummaryWithDefaults() *WorkloadStatsSummary {
	this := WorkloadStatsSummary{}
	return &this
}

// GetWorkloads returns the Workloads field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkloadStatsSummary) GetWorkloads() []WorkloadStatsSchema {
	if o == nil {
		var ret []WorkloadStatsSchema
		return ret
	}
	return o.Workloads
}

// GetWorkloadsOk returns a tuple with the Workloads field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkloadStatsSummary) GetWorkloadsOk() ([]WorkloadStatsSchema, bool) {
	if o == nil || IsNil(o.Workloads) {
		return nil, false
	}
	return o.Workloads, true
}

// HasWorkloads returns a boolean if a field has been set.
func (o *WorkloadStatsSummary) HasWorkloads() bool {
	if o != nil && !IsNil(o.Workloads) {
		return true
	}

	return false
}

// SetWorkloads gets a reference to the given []WorkloadStatsSchema and assigns it to the Workloads field.
func (o *WorkloadStatsSummary) SetWorkloads(v []WorkloadStatsSchema) {
	o.Workloads = v
}

func (o WorkloadStatsSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkloadStatsSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Workloads != nil {
		toSerialize["workloads"] = o.Workloads
	}
	return toSerialize, nil
}

type NullableWorkloadStatsSummary struct {
	value *WorkloadStatsSummary
	isSet bool
}

func (v NullableWorkloadStatsSummary) Get() *WorkloadStatsSummary {
	return v.value
}

func (v *NullableWorkloadStatsSummary) Set(val *WorkloadStatsSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkloadStatsSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkloadStatsSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkloadStatsSummary(val *WorkloadStatsSummary) *NullableWorkloadStatsSummary {
	return &NullableWorkloadStatsSummary{value: val, isSet: true}
}

func (v NullableWorkloadStatsSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkloadStatsSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


