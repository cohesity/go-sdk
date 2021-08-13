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

// McmGetProtectionLastRunStatsResponseBody Specifies the last Protection Run stats across objects.
type McmGetProtectionLastRunStatsResponseBody struct {
	// Specifies the aggregated status across all adapters for respective last run.
	AggregatedStats *[]McmLastRunStatusStats `json:"aggregatedStats,omitempty"`
	// Specifies the enviournment specific last run status stats.
	EnvStats *[]McmLastRunEnvSpecificStatusStats `json:"envStats,omitempty"`
}

// NewMcmGetProtectionLastRunStatsResponseBody instantiates a new McmGetProtectionLastRunStatsResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMcmGetProtectionLastRunStatsResponseBody() *McmGetProtectionLastRunStatsResponseBody {
	this := McmGetProtectionLastRunStatsResponseBody{}
	return &this
}

// NewMcmGetProtectionLastRunStatsResponseBodyWithDefaults instantiates a new McmGetProtectionLastRunStatsResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMcmGetProtectionLastRunStatsResponseBodyWithDefaults() *McmGetProtectionLastRunStatsResponseBody {
	this := McmGetProtectionLastRunStatsResponseBody{}
	return &this
}

// GetAggregatedStats returns the AggregatedStats field value if set, zero value otherwise.
func (o *McmGetProtectionLastRunStatsResponseBody) GetAggregatedStats() []McmLastRunStatusStats {
	if o == nil || o.AggregatedStats == nil {
		var ret []McmLastRunStatusStats
		return ret
	}
	return *o.AggregatedStats
}

// GetAggregatedStatsOk returns a tuple with the AggregatedStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *McmGetProtectionLastRunStatsResponseBody) GetAggregatedStatsOk() (*[]McmLastRunStatusStats, bool) {
	if o == nil || o.AggregatedStats == nil {
		return nil, false
	}
	return o.AggregatedStats, true
}

// HasAggregatedStats returns a boolean if a field has been set.
func (o *McmGetProtectionLastRunStatsResponseBody) HasAggregatedStats() bool {
	if o != nil && o.AggregatedStats != nil {
		return true
	}

	return false
}

// SetAggregatedStats gets a reference to the given []McmLastRunStatusStats and assigns it to the AggregatedStats field.
func (o *McmGetProtectionLastRunStatsResponseBody) SetAggregatedStats(v []McmLastRunStatusStats) {
	o.AggregatedStats = &v
}

// GetEnvStats returns the EnvStats field value if set, zero value otherwise.
func (o *McmGetProtectionLastRunStatsResponseBody) GetEnvStats() []McmLastRunEnvSpecificStatusStats {
	if o == nil || o.EnvStats == nil {
		var ret []McmLastRunEnvSpecificStatusStats
		return ret
	}
	return *o.EnvStats
}

// GetEnvStatsOk returns a tuple with the EnvStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *McmGetProtectionLastRunStatsResponseBody) GetEnvStatsOk() (*[]McmLastRunEnvSpecificStatusStats, bool) {
	if o == nil || o.EnvStats == nil {
		return nil, false
	}
	return o.EnvStats, true
}

// HasEnvStats returns a boolean if a field has been set.
func (o *McmGetProtectionLastRunStatsResponseBody) HasEnvStats() bool {
	if o != nil && o.EnvStats != nil {
		return true
	}

	return false
}

// SetEnvStats gets a reference to the given []McmLastRunEnvSpecificStatusStats and assigns it to the EnvStats field.
func (o *McmGetProtectionLastRunStatsResponseBody) SetEnvStats(v []McmLastRunEnvSpecificStatusStats) {
	o.EnvStats = &v
}

func (o McmGetProtectionLastRunStatsResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AggregatedStats != nil {
		toSerialize["aggregatedStats"] = o.AggregatedStats
	}
	if o.EnvStats != nil {
		toSerialize["envStats"] = o.EnvStats
	}
	return json.Marshal(toSerialize)
}

type NullableMcmGetProtectionLastRunStatsResponseBody struct {
	value *McmGetProtectionLastRunStatsResponseBody
	isSet bool
}

func (v NullableMcmGetProtectionLastRunStatsResponseBody) Get() *McmGetProtectionLastRunStatsResponseBody {
	return v.value
}

func (v *NullableMcmGetProtectionLastRunStatsResponseBody) Set(val *McmGetProtectionLastRunStatsResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableMcmGetProtectionLastRunStatsResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableMcmGetProtectionLastRunStatsResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMcmGetProtectionLastRunStatsResponseBody(val *McmGetProtectionLastRunStatsResponseBody) *NullableMcmGetProtectionLastRunStatsResponseBody {
	return &NullableMcmGetProtectionLastRunStatsResponseBody{value: val, isSet: true}
}

func (v NullableMcmGetProtectionLastRunStatsResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMcmGetProtectionLastRunStatsResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


