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

// FailoverRunsResponse Specifies the response upon creating a special run during failover workflow.
type FailoverRunsResponse struct {
	// Specifies the list of planned runs created during various planeed failover workflows. Each planned run is uniqely identified by falioverId and runId.
	FailoverPlannedRuns []PlannedRunPollStatus `json:"failoverPlannedRuns,omitempty"`
}

// NewFailoverRunsResponse instantiates a new FailoverRunsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFailoverRunsResponse() *FailoverRunsResponse {
	this := FailoverRunsResponse{}
	return &this
}

// NewFailoverRunsResponseWithDefaults instantiates a new FailoverRunsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFailoverRunsResponseWithDefaults() *FailoverRunsResponse {
	this := FailoverRunsResponse{}
	return &this
}

// GetFailoverPlannedRuns returns the FailoverPlannedRuns field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FailoverRunsResponse) GetFailoverPlannedRuns() []PlannedRunPollStatus {
	if o == nil  {
		var ret []PlannedRunPollStatus
		return ret
	}
	return o.FailoverPlannedRuns
}

// GetFailoverPlannedRunsOk returns a tuple with the FailoverPlannedRuns field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FailoverRunsResponse) GetFailoverPlannedRunsOk() (*[]PlannedRunPollStatus, bool) {
	if o == nil || o.FailoverPlannedRuns == nil {
		return nil, false
	}
	return &o.FailoverPlannedRuns, true
}

// HasFailoverPlannedRuns returns a boolean if a field has been set.
func (o *FailoverRunsResponse) HasFailoverPlannedRuns() bool {
	if o != nil && o.FailoverPlannedRuns != nil {
		return true
	}

	return false
}

// SetFailoverPlannedRuns gets a reference to the given []PlannedRunPollStatus and assigns it to the FailoverPlannedRuns field.
func (o *FailoverRunsResponse) SetFailoverPlannedRuns(v []PlannedRunPollStatus) {
	o.FailoverPlannedRuns = v
}

func (o FailoverRunsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FailoverPlannedRuns != nil {
		toSerialize["failoverPlannedRuns"] = o.FailoverPlannedRuns
	}
	return json.Marshal(toSerialize)
}

type NullableFailoverRunsResponse struct {
	value *FailoverRunsResponse
	isSet bool
}

func (v NullableFailoverRunsResponse) Get() *FailoverRunsResponse {
	return v.value
}

func (v *NullableFailoverRunsResponse) Set(val *FailoverRunsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFailoverRunsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFailoverRunsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFailoverRunsResponse(val *FailoverRunsResponse) *NullableFailoverRunsResponse {
	return &NullableFailoverRunsResponse{value: val, isSet: true}
}

func (v NullableFailoverRunsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFailoverRunsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o FailoverRunsResponse) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}