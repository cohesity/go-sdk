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

// GetProtectionRunsStatusResponseBody Specifies a list of protection runs stats taken at different time.
type GetProtectionRunsStatusResponseBody struct {
	// Specifies a list of protection runs stats taken at different time.
	ProtectionRunsStatsList []ProtectionRunsStatsList `json:"protectionRunsStatsList,omitempty"`
}

// NewGetProtectionRunsStatusResponseBody instantiates a new GetProtectionRunsStatusResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProtectionRunsStatusResponseBody() *GetProtectionRunsStatusResponseBody {
	this := GetProtectionRunsStatusResponseBody{}
	return &this
}

// NewGetProtectionRunsStatusResponseBodyWithDefaults instantiates a new GetProtectionRunsStatusResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProtectionRunsStatusResponseBodyWithDefaults() *GetProtectionRunsStatusResponseBody {
	this := GetProtectionRunsStatusResponseBody{}
	return &this
}

// GetProtectionRunsStatsList returns the ProtectionRunsStatsList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetProtectionRunsStatusResponseBody) GetProtectionRunsStatsList() []ProtectionRunsStatsList {
	if o == nil  {
		var ret []ProtectionRunsStatsList
		return ret
	}
	return o.ProtectionRunsStatsList
}

// GetProtectionRunsStatsListOk returns a tuple with the ProtectionRunsStatsList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetProtectionRunsStatusResponseBody) GetProtectionRunsStatsListOk() (*[]ProtectionRunsStatsList, bool) {
	if o == nil || o.ProtectionRunsStatsList == nil {
		return nil, false
	}
	return &o.ProtectionRunsStatsList, true
}

// HasProtectionRunsStatsList returns a boolean if a field has been set.
func (o *GetProtectionRunsStatusResponseBody) HasProtectionRunsStatsList() bool {
	if o != nil && o.ProtectionRunsStatsList != nil {
		return true
	}

	return false
}

// SetProtectionRunsStatsList gets a reference to the given []ProtectionRunsStatsList and assigns it to the ProtectionRunsStatsList field.
func (o *GetProtectionRunsStatusResponseBody) SetProtectionRunsStatsList(v []ProtectionRunsStatsList) {
	o.ProtectionRunsStatsList = v
}

func (o GetProtectionRunsStatusResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProtectionRunsStatsList != nil {
		toSerialize["protectionRunsStatsList"] = o.ProtectionRunsStatsList
	}
	return json.Marshal(toSerialize)
}

type NullableGetProtectionRunsStatusResponseBody struct {
	value *GetProtectionRunsStatusResponseBody
	isSet bool
}

func (v NullableGetProtectionRunsStatusResponseBody) Get() *GetProtectionRunsStatusResponseBody {
	return v.value
}

func (v *NullableGetProtectionRunsStatusResponseBody) Set(val *GetProtectionRunsStatusResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProtectionRunsStatusResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProtectionRunsStatusResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProtectionRunsStatusResponseBody(val *GetProtectionRunsStatusResponseBody) *NullableGetProtectionRunsStatusResponseBody {
	return &NullableGetProtectionRunsStatusResponseBody{value: val, isSet: true}
}

func (v NullableGetProtectionRunsStatusResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProtectionRunsStatusResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o GetProtectionRunsStatusResponseBody) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}