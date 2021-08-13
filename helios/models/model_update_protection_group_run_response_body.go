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

// UpdateProtectionGroupRunResponseBody Specifies the response of update Protection Group Runs request.
type UpdateProtectionGroupRunResponseBody struct {
	// Specifies a list of ids of Protection Group Runs that are successfully updated.
	SuccessfulRunIds []string `json:"successfulRunIds,omitempty"`
	// Specifies a list of ids of Protection Group Runs that failed to update along with error details.
	FailedRuns []FailedRunDetails `json:"failedRuns,omitempty"`
}

// NewUpdateProtectionGroupRunResponseBody instantiates a new UpdateProtectionGroupRunResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateProtectionGroupRunResponseBody() *UpdateProtectionGroupRunResponseBody {
	this := UpdateProtectionGroupRunResponseBody{}
	return &this
}

// NewUpdateProtectionGroupRunResponseBodyWithDefaults instantiates a new UpdateProtectionGroupRunResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateProtectionGroupRunResponseBodyWithDefaults() *UpdateProtectionGroupRunResponseBody {
	this := UpdateProtectionGroupRunResponseBody{}
	return &this
}

// GetSuccessfulRunIds returns the SuccessfulRunIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateProtectionGroupRunResponseBody) GetSuccessfulRunIds() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.SuccessfulRunIds
}

// GetSuccessfulRunIdsOk returns a tuple with the SuccessfulRunIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateProtectionGroupRunResponseBody) GetSuccessfulRunIdsOk() (*[]string, bool) {
	if o == nil || o.SuccessfulRunIds == nil {
		return nil, false
	}
	return &o.SuccessfulRunIds, true
}

// HasSuccessfulRunIds returns a boolean if a field has been set.
func (o *UpdateProtectionGroupRunResponseBody) HasSuccessfulRunIds() bool {
	if o != nil && o.SuccessfulRunIds != nil {
		return true
	}

	return false
}

// SetSuccessfulRunIds gets a reference to the given []string and assigns it to the SuccessfulRunIds field.
func (o *UpdateProtectionGroupRunResponseBody) SetSuccessfulRunIds(v []string) {
	o.SuccessfulRunIds = v
}

// GetFailedRuns returns the FailedRuns field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateProtectionGroupRunResponseBody) GetFailedRuns() []FailedRunDetails {
	if o == nil  {
		var ret []FailedRunDetails
		return ret
	}
	return o.FailedRuns
}

// GetFailedRunsOk returns a tuple with the FailedRuns field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateProtectionGroupRunResponseBody) GetFailedRunsOk() (*[]FailedRunDetails, bool) {
	if o == nil || o.FailedRuns == nil {
		return nil, false
	}
	return &o.FailedRuns, true
}

// HasFailedRuns returns a boolean if a field has been set.
func (o *UpdateProtectionGroupRunResponseBody) HasFailedRuns() bool {
	if o != nil && o.FailedRuns != nil {
		return true
	}

	return false
}

// SetFailedRuns gets a reference to the given []FailedRunDetails and assigns it to the FailedRuns field.
func (o *UpdateProtectionGroupRunResponseBody) SetFailedRuns(v []FailedRunDetails) {
	o.FailedRuns = v
}

func (o UpdateProtectionGroupRunResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SuccessfulRunIds != nil {
		toSerialize["successfulRunIds"] = o.SuccessfulRunIds
	}
	if o.FailedRuns != nil {
		toSerialize["failedRuns"] = o.FailedRuns
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateProtectionGroupRunResponseBody struct {
	value *UpdateProtectionGroupRunResponseBody
	isSet bool
}

func (v NullableUpdateProtectionGroupRunResponseBody) Get() *UpdateProtectionGroupRunResponseBody {
	return v.value
}

func (v *NullableUpdateProtectionGroupRunResponseBody) Set(val *UpdateProtectionGroupRunResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateProtectionGroupRunResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateProtectionGroupRunResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateProtectionGroupRunResponseBody(val *UpdateProtectionGroupRunResponseBody) *NullableUpdateProtectionGroupRunResponseBody {
	return &NullableUpdateProtectionGroupRunResponseBody{value: val, isSet: true}
}

func (v NullableUpdateProtectionGroupRunResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateProtectionGroupRunResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


