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

// PerformActionOnProtectionGroupRunRequest Specifies the request to perform actions on protection runs.
type PerformActionOnProtectionGroupRunRequest struct {
	// Specifies the type of the action which will be performed on protection runs.
	Action NullableString `json:"action"`
	// Specifies the pause action params for a protection run.
	PauseParams *[]PauseProtectionRunActionParams `json:"pauseParams,omitempty"`
	// Specifies the resume action params for a protection run.
	ResumeParams *[]ResumeProtectionRunActionParams `json:"resumeParams,omitempty"`
	// Specifies the cancel action params for a protection run.
	CancelParams *[]CancelProtectionGroupRunRequest `json:"cancelParams,omitempty"`
}

// NewPerformActionOnProtectionGroupRunRequest instantiates a new PerformActionOnProtectionGroupRunRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerformActionOnProtectionGroupRunRequest(action NullableString) *PerformActionOnProtectionGroupRunRequest {
	this := PerformActionOnProtectionGroupRunRequest{}
	this.Action = action
	return &this
}

// NewPerformActionOnProtectionGroupRunRequestWithDefaults instantiates a new PerformActionOnProtectionGroupRunRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPerformActionOnProtectionGroupRunRequestWithDefaults() *PerformActionOnProtectionGroupRunRequest {
	this := PerformActionOnProtectionGroupRunRequest{}
	return &this
}

// GetAction returns the Action field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PerformActionOnProtectionGroupRunRequest) GetAction() string {
	if o == nil || o.Action.Get() == nil {
		var ret string
		return ret
	}

	return *o.Action.Get()
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PerformActionOnProtectionGroupRunRequest) GetActionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Action.Get(), o.Action.IsSet()
}

// SetAction sets field value
func (o *PerformActionOnProtectionGroupRunRequest) SetAction(v string) {
	o.Action.Set(&v)
}

// GetPauseParams returns the PauseParams field value if set, zero value otherwise.
func (o *PerformActionOnProtectionGroupRunRequest) GetPauseParams() []PauseProtectionRunActionParams {
	if o == nil || o.PauseParams == nil {
		var ret []PauseProtectionRunActionParams
		return ret
	}
	return *o.PauseParams
}

// GetPauseParamsOk returns a tuple with the PauseParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformActionOnProtectionGroupRunRequest) GetPauseParamsOk() (*[]PauseProtectionRunActionParams, bool) {
	if o == nil || o.PauseParams == nil {
		return nil, false
	}
	return o.PauseParams, true
}

// HasPauseParams returns a boolean if a field has been set.
func (o *PerformActionOnProtectionGroupRunRequest) HasPauseParams() bool {
	if o != nil && o.PauseParams != nil {
		return true
	}

	return false
}

// SetPauseParams gets a reference to the given []PauseProtectionRunActionParams and assigns it to the PauseParams field.
func (o *PerformActionOnProtectionGroupRunRequest) SetPauseParams(v []PauseProtectionRunActionParams) {
	o.PauseParams = &v
}

// GetResumeParams returns the ResumeParams field value if set, zero value otherwise.
func (o *PerformActionOnProtectionGroupRunRequest) GetResumeParams() []ResumeProtectionRunActionParams {
	if o == nil || o.ResumeParams == nil {
		var ret []ResumeProtectionRunActionParams
		return ret
	}
	return *o.ResumeParams
}

// GetResumeParamsOk returns a tuple with the ResumeParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformActionOnProtectionGroupRunRequest) GetResumeParamsOk() (*[]ResumeProtectionRunActionParams, bool) {
	if o == nil || o.ResumeParams == nil {
		return nil, false
	}
	return o.ResumeParams, true
}

// HasResumeParams returns a boolean if a field has been set.
func (o *PerformActionOnProtectionGroupRunRequest) HasResumeParams() bool {
	if o != nil && o.ResumeParams != nil {
		return true
	}

	return false
}

// SetResumeParams gets a reference to the given []ResumeProtectionRunActionParams and assigns it to the ResumeParams field.
func (o *PerformActionOnProtectionGroupRunRequest) SetResumeParams(v []ResumeProtectionRunActionParams) {
	o.ResumeParams = &v
}

// GetCancelParams returns the CancelParams field value if set, zero value otherwise.
func (o *PerformActionOnProtectionGroupRunRequest) GetCancelParams() []CancelProtectionGroupRunRequest {
	if o == nil || o.CancelParams == nil {
		var ret []CancelProtectionGroupRunRequest
		return ret
	}
	return *o.CancelParams
}

// GetCancelParamsOk returns a tuple with the CancelParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformActionOnProtectionGroupRunRequest) GetCancelParamsOk() (*[]CancelProtectionGroupRunRequest, bool) {
	if o == nil || o.CancelParams == nil {
		return nil, false
	}
	return o.CancelParams, true
}

// HasCancelParams returns a boolean if a field has been set.
func (o *PerformActionOnProtectionGroupRunRequest) HasCancelParams() bool {
	if o != nil && o.CancelParams != nil {
		return true
	}

	return false
}

// SetCancelParams gets a reference to the given []CancelProtectionGroupRunRequest and assigns it to the CancelParams field.
func (o *PerformActionOnProtectionGroupRunRequest) SetCancelParams(v []CancelProtectionGroupRunRequest) {
	o.CancelParams = &v
}

func (o PerformActionOnProtectionGroupRunRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["action"] = o.Action.Get()
	}
	if o.PauseParams != nil {
		toSerialize["pauseParams"] = o.PauseParams
	}
	if o.ResumeParams != nil {
		toSerialize["resumeParams"] = o.ResumeParams
	}
	if o.CancelParams != nil {
		toSerialize["cancelParams"] = o.CancelParams
	}
	return json.Marshal(toSerialize)
}

type NullablePerformActionOnProtectionGroupRunRequest struct {
	value *PerformActionOnProtectionGroupRunRequest
	isSet bool
}

func (v NullablePerformActionOnProtectionGroupRunRequest) Get() *PerformActionOnProtectionGroupRunRequest {
	return v.value
}

func (v *NullablePerformActionOnProtectionGroupRunRequest) Set(val *PerformActionOnProtectionGroupRunRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePerformActionOnProtectionGroupRunRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePerformActionOnProtectionGroupRunRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerformActionOnProtectionGroupRunRequest(val *PerformActionOnProtectionGroupRunRequest) *NullablePerformActionOnProtectionGroupRunRequest {
	return &NullablePerformActionOnProtectionGroupRunRequest{value: val, isSet: true}
}

func (v NullablePerformActionOnProtectionGroupRunRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePerformActionOnProtectionGroupRunRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


