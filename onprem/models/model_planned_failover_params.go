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

// PlannedFailoverParams Specifies parameters of a planned failover.
type PlannedFailoverParams struct {
	// Spcifies the planned failover type.<br> 'Prepare' indicates this is a preparation for failover.<br> 'Finalize' indicates this is finalization of failover. After this is done, the view can be used as source view.
	Type NullableString `json:"type"`
	// Specifies parameters of preparation of a planned failover.
	PreparePlannedFailverParams *PreparePlannedFailverParams `json:"preparePlannedFailverParams,omitempty"`
}

// NewPlannedFailoverParams instantiates a new PlannedFailoverParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlannedFailoverParams(type_ NullableString) *PlannedFailoverParams {
	this := PlannedFailoverParams{}
	this.Type = type_
	return &this
}

// NewPlannedFailoverParamsWithDefaults instantiates a new PlannedFailoverParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlannedFailoverParamsWithDefaults() *PlannedFailoverParams {
	this := PlannedFailoverParams{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PlannedFailoverParams) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}

	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlannedFailoverParams) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// SetType sets field value
func (o *PlannedFailoverParams) SetType(v string) {
	o.Type.Set(&v)
}

// GetPreparePlannedFailverParams returns the PreparePlannedFailverParams field value if set, zero value otherwise.
func (o *PlannedFailoverParams) GetPreparePlannedFailverParams() PreparePlannedFailverParams {
	if o == nil || o.PreparePlannedFailverParams == nil {
		var ret PreparePlannedFailverParams
		return ret
	}
	return *o.PreparePlannedFailverParams
}

// GetPreparePlannedFailverParamsOk returns a tuple with the PreparePlannedFailverParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlannedFailoverParams) GetPreparePlannedFailverParamsOk() (*PreparePlannedFailverParams, bool) {
	if o == nil || o.PreparePlannedFailverParams == nil {
		return nil, false
	}
	return o.PreparePlannedFailverParams, true
}

// HasPreparePlannedFailverParams returns a boolean if a field has been set.
func (o *PlannedFailoverParams) HasPreparePlannedFailverParams() bool {
	if o != nil && o.PreparePlannedFailverParams != nil {
		return true
	}

	return false
}

// SetPreparePlannedFailverParams gets a reference to the given PreparePlannedFailverParams and assigns it to the PreparePlannedFailverParams field.
func (o *PlannedFailoverParams) SetPreparePlannedFailverParams(v PreparePlannedFailverParams) {
	o.PreparePlannedFailverParams = &v
}

func (o PlannedFailoverParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type.Get()
	}
	if o.PreparePlannedFailverParams != nil {
		toSerialize["preparePlannedFailverParams"] = o.PreparePlannedFailverParams
	}
	return json.Marshal(toSerialize)
}

type NullablePlannedFailoverParams struct {
	value *PlannedFailoverParams
	isSet bool
}

func (v NullablePlannedFailoverParams) Get() *PlannedFailoverParams {
	return v.value
}

func (v *NullablePlannedFailoverParams) Set(val *PlannedFailoverParams) {
	v.value = val
	v.isSet = true
}

func (v NullablePlannedFailoverParams) IsSet() bool {
	return v.isSet
}

func (v *NullablePlannedFailoverParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlannedFailoverParams(val *PlannedFailoverParams) *NullablePlannedFailoverParams {
	return &NullablePlannedFailoverParams{value: val, isSet: true}
}

func (v NullablePlannedFailoverParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlannedFailoverParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o PlannedFailoverParams) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}