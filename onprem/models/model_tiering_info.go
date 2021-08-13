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

// TieringInfo Specifies the data tiering task details.
type TieringInfo struct {
	Progress *ProgressSummary `json:"progress,omitempty"`
	// Status of the analysis run. 'Running' indicates that the run is still running. 'Canceled' indicates that the run has been canceled. 'Canceling' indicates that the run is in the process of being  canceled. 'Failed' indicates that the run has failed. 'Missed' indicates that the run was unable to take place at the  scheduled time because the previous run was still happening. 'Succeeded' indicates that the run has finished successfully. 'SucceededWithWarning' indicates that the run finished  successfully, but there were some warning messages. 'OnHold' indicates that the run has On hold.
	Status NullableString `json:"status,omitempty"`
	Stats *DataTieringTaskStats `json:"stats,omitempty"`
}

// NewTieringInfo instantiates a new TieringInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTieringInfo() *TieringInfo {
	this := TieringInfo{}
	return &this
}

// NewTieringInfoWithDefaults instantiates a new TieringInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTieringInfoWithDefaults() *TieringInfo {
	this := TieringInfo{}
	return &this
}

// GetProgress returns the Progress field value if set, zero value otherwise.
func (o *TieringInfo) GetProgress() ProgressSummary {
	if o == nil || o.Progress == nil {
		var ret ProgressSummary
		return ret
	}
	return *o.Progress
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TieringInfo) GetProgressOk() (*ProgressSummary, bool) {
	if o == nil || o.Progress == nil {
		return nil, false
	}
	return o.Progress, true
}

// HasProgress returns a boolean if a field has been set.
func (o *TieringInfo) HasProgress() bool {
	if o != nil && o.Progress != nil {
		return true
	}

	return false
}

// SetProgress gets a reference to the given ProgressSummary and assigns it to the Progress field.
func (o *TieringInfo) SetProgress(v ProgressSummary) {
	o.Progress = &v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TieringInfo) GetStatus() string {
	if o == nil || o.Status.Get() == nil {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TieringInfo) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *TieringInfo) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *TieringInfo) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *TieringInfo) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *TieringInfo) UnsetStatus() {
	o.Status.Unset()
}

// GetStats returns the Stats field value if set, zero value otherwise.
func (o *TieringInfo) GetStats() DataTieringTaskStats {
	if o == nil || o.Stats == nil {
		var ret DataTieringTaskStats
		return ret
	}
	return *o.Stats
}

// GetStatsOk returns a tuple with the Stats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TieringInfo) GetStatsOk() (*DataTieringTaskStats, bool) {
	if o == nil || o.Stats == nil {
		return nil, false
	}
	return o.Stats, true
}

// HasStats returns a boolean if a field has been set.
func (o *TieringInfo) HasStats() bool {
	if o != nil && o.Stats != nil {
		return true
	}

	return false
}

// SetStats gets a reference to the given DataTieringTaskStats and assigns it to the Stats field.
func (o *TieringInfo) SetStats(v DataTieringTaskStats) {
	o.Stats = &v
}

func (o TieringInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Progress != nil {
		toSerialize["progress"] = o.Progress
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.Stats != nil {
		toSerialize["stats"] = o.Stats
	}
	return json.Marshal(toSerialize)
}

type NullableTieringInfo struct {
	value *TieringInfo
	isSet bool
}

func (v NullableTieringInfo) Get() *TieringInfo {
	return v.value
}

func (v *NullableTieringInfo) Set(val *TieringInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTieringInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTieringInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTieringInfo(val *TieringInfo) *NullableTieringInfo {
	return &NullableTieringInfo{value: val, isSet: true}
}

func (v NullableTieringInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTieringInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o TieringInfo) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}