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

// ICAPUriStatus Specifies the ICAP Uri connection status.
type ICAPUriStatus struct {
	// Specifies the ICAP Uri connection status.
	Type *string `json:"type,omitempty"`
}

// NewICAPUriStatus instantiates a new ICAPUriStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewICAPUriStatus() *ICAPUriStatus {
	this := ICAPUriStatus{}
	return &this
}

// NewICAPUriStatusWithDefaults instantiates a new ICAPUriStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewICAPUriStatusWithDefaults() *ICAPUriStatus {
	this := ICAPUriStatus{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ICAPUriStatus) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ICAPUriStatus) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ICAPUriStatus) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ICAPUriStatus) SetType(v string) {
	o.Type = &v
}

func (o ICAPUriStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableICAPUriStatus struct {
	value *ICAPUriStatus
	isSet bool
}

func (v NullableICAPUriStatus) Get() *ICAPUriStatus {
	return v.value
}

func (v *NullableICAPUriStatus) Set(val *ICAPUriStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableICAPUriStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableICAPUriStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableICAPUriStatus(val *ICAPUriStatus) *NullableICAPUriStatus {
	return &NullableICAPUriStatus{value: val, isSet: true}
}

func (v NullableICAPUriStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableICAPUriStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


