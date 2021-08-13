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

// TearDownStatus Tear Down Status
type TearDownStatus struct {
	// Specifies the tear down status.
	TearDownStatus *string `json:"tearDownStatus,omitempty"`
}

// NewTearDownStatus instantiates a new TearDownStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTearDownStatus() *TearDownStatus {
	this := TearDownStatus{}
	return &this
}

// NewTearDownStatusWithDefaults instantiates a new TearDownStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTearDownStatusWithDefaults() *TearDownStatus {
	this := TearDownStatus{}
	return &this
}

// GetTearDownStatus returns the TearDownStatus field value if set, zero value otherwise.
func (o *TearDownStatus) GetTearDownStatus() string {
	if o == nil || o.TearDownStatus == nil {
		var ret string
		return ret
	}
	return *o.TearDownStatus
}

// GetTearDownStatusOk returns a tuple with the TearDownStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TearDownStatus) GetTearDownStatusOk() (*string, bool) {
	if o == nil || o.TearDownStatus == nil {
		return nil, false
	}
	return o.TearDownStatus, true
}

// HasTearDownStatus returns a boolean if a field has been set.
func (o *TearDownStatus) HasTearDownStatus() bool {
	if o != nil && o.TearDownStatus != nil {
		return true
	}

	return false
}

// SetTearDownStatus gets a reference to the given string and assigns it to the TearDownStatus field.
func (o *TearDownStatus) SetTearDownStatus(v string) {
	o.TearDownStatus = &v
}

func (o TearDownStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TearDownStatus != nil {
		toSerialize["tearDownStatus"] = o.TearDownStatus
	}
	return json.Marshal(toSerialize)
}

type NullableTearDownStatus struct {
	value *TearDownStatus
	isSet bool
}

func (v NullableTearDownStatus) Get() *TearDownStatus {
	return v.value
}

func (v *NullableTearDownStatus) Set(val *TearDownStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableTearDownStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableTearDownStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTearDownStatus(val *TearDownStatus) *NullableTearDownStatus {
	return &NullableTearDownStatus{value: val, isSet: true}
}

func (v NullableTearDownStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTearDownStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o TearDownStatus) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}