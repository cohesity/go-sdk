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

// CreateViewFailoverRequest Specifies the request parameters to create a view failover task.
type CreateViewFailoverRequest struct {
	// Specifies the failover type.<br> 'Planned' indicates this is a planned failover.<br> 'Unplanned' indicates this is an unplanned failover, which is used when source cluster is down.
	Type NullableString `json:"type"`
	// Specifies parameters to create a planned failover.
	PlannedFailoverParams *PlannedFailoverParams `json:"plannedFailoverParams,omitempty"`
	// Specifies parameters to create an unplanned failover.
	UnplannedFailoverParams *UnplannedFailoverParams `json:"unplannedFailoverParams,omitempty"`
}

// NewCreateViewFailoverRequest instantiates a new CreateViewFailoverRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateViewFailoverRequest(type_ NullableString) *CreateViewFailoverRequest {
	this := CreateViewFailoverRequest{}
	this.Type = type_
	return &this
}

// NewCreateViewFailoverRequestWithDefaults instantiates a new CreateViewFailoverRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateViewFailoverRequestWithDefaults() *CreateViewFailoverRequest {
	this := CreateViewFailoverRequest{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreateViewFailoverRequest) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}

	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateViewFailoverRequest) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// SetType sets field value
func (o *CreateViewFailoverRequest) SetType(v string) {
	o.Type.Set(&v)
}

// GetPlannedFailoverParams returns the PlannedFailoverParams field value if set, zero value otherwise.
func (o *CreateViewFailoverRequest) GetPlannedFailoverParams() PlannedFailoverParams {
	if o == nil || o.PlannedFailoverParams == nil {
		var ret PlannedFailoverParams
		return ret
	}
	return *o.PlannedFailoverParams
}

// GetPlannedFailoverParamsOk returns a tuple with the PlannedFailoverParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateViewFailoverRequest) GetPlannedFailoverParamsOk() (*PlannedFailoverParams, bool) {
	if o == nil || o.PlannedFailoverParams == nil {
		return nil, false
	}
	return o.PlannedFailoverParams, true
}

// HasPlannedFailoverParams returns a boolean if a field has been set.
func (o *CreateViewFailoverRequest) HasPlannedFailoverParams() bool {
	if o != nil && o.PlannedFailoverParams != nil {
		return true
	}

	return false
}

// SetPlannedFailoverParams gets a reference to the given PlannedFailoverParams and assigns it to the PlannedFailoverParams field.
func (o *CreateViewFailoverRequest) SetPlannedFailoverParams(v PlannedFailoverParams) {
	o.PlannedFailoverParams = &v
}

// GetUnplannedFailoverParams returns the UnplannedFailoverParams field value if set, zero value otherwise.
func (o *CreateViewFailoverRequest) GetUnplannedFailoverParams() UnplannedFailoverParams {
	if o == nil || o.UnplannedFailoverParams == nil {
		var ret UnplannedFailoverParams
		return ret
	}
	return *o.UnplannedFailoverParams
}

// GetUnplannedFailoverParamsOk returns a tuple with the UnplannedFailoverParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateViewFailoverRequest) GetUnplannedFailoverParamsOk() (*UnplannedFailoverParams, bool) {
	if o == nil || o.UnplannedFailoverParams == nil {
		return nil, false
	}
	return o.UnplannedFailoverParams, true
}

// HasUnplannedFailoverParams returns a boolean if a field has been set.
func (o *CreateViewFailoverRequest) HasUnplannedFailoverParams() bool {
	if o != nil && o.UnplannedFailoverParams != nil {
		return true
	}

	return false
}

// SetUnplannedFailoverParams gets a reference to the given UnplannedFailoverParams and assigns it to the UnplannedFailoverParams field.
func (o *CreateViewFailoverRequest) SetUnplannedFailoverParams(v UnplannedFailoverParams) {
	o.UnplannedFailoverParams = &v
}

func (o CreateViewFailoverRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type.Get()
	}
	if o.PlannedFailoverParams != nil {
		toSerialize["plannedFailoverParams"] = o.PlannedFailoverParams
	}
	if o.UnplannedFailoverParams != nil {
		toSerialize["unplannedFailoverParams"] = o.UnplannedFailoverParams
	}
	return json.Marshal(toSerialize)
}

type NullableCreateViewFailoverRequest struct {
	value *CreateViewFailoverRequest
	isSet bool
}

func (v NullableCreateViewFailoverRequest) Get() *CreateViewFailoverRequest {
	return v.value
}

func (v *NullableCreateViewFailoverRequest) Set(val *CreateViewFailoverRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateViewFailoverRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateViewFailoverRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateViewFailoverRequest(val *CreateViewFailoverRequest) *NullableCreateViewFailoverRequest {
	return &NullableCreateViewFailoverRequest{value: val, isSet: true}
}

func (v NullableCreateViewFailoverRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateViewFailoverRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


