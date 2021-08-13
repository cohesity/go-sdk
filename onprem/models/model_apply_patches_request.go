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

// ApplyPatchesRequest Specifies the request to apply a patch.
type ApplyPatchesRequest struct {
	// Specifies the name of the service whose patch should be applied.
	Service NullableString `json:"service,omitempty"`
	// Specifies all the available patches should be applied.
	All NullableBool `json:"all,omitempty"`
}

// NewApplyPatchesRequest instantiates a new ApplyPatchesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplyPatchesRequest() *ApplyPatchesRequest {
	this := ApplyPatchesRequest{}
	return &this
}

// NewApplyPatchesRequestWithDefaults instantiates a new ApplyPatchesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplyPatchesRequestWithDefaults() *ApplyPatchesRequest {
	this := ApplyPatchesRequest{}
	return &this
}

// GetService returns the Service field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplyPatchesRequest) GetService() string {
	if o == nil || o.Service.Get() == nil {
		var ret string
		return ret
	}
	return *o.Service.Get()
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplyPatchesRequest) GetServiceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Service.Get(), o.Service.IsSet()
}

// HasService returns a boolean if a field has been set.
func (o *ApplyPatchesRequest) HasService() bool {
	if o != nil && o.Service.IsSet() {
		return true
	}

	return false
}

// SetService gets a reference to the given NullableString and assigns it to the Service field.
func (o *ApplyPatchesRequest) SetService(v string) {
	o.Service.Set(&v)
}
// SetServiceNil sets the value for Service to be an explicit nil
func (o *ApplyPatchesRequest) SetServiceNil() {
	o.Service.Set(nil)
}

// UnsetService ensures that no value is present for Service, not even an explicit nil
func (o *ApplyPatchesRequest) UnsetService() {
	o.Service.Unset()
}

// GetAll returns the All field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplyPatchesRequest) GetAll() bool {
	if o == nil || o.All.Get() == nil {
		var ret bool
		return ret
	}
	return *o.All.Get()
}

// GetAllOk returns a tuple with the All field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplyPatchesRequest) GetAllOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.All.Get(), o.All.IsSet()
}

// HasAll returns a boolean if a field has been set.
func (o *ApplyPatchesRequest) HasAll() bool {
	if o != nil && o.All.IsSet() {
		return true
	}

	return false
}

// SetAll gets a reference to the given NullableBool and assigns it to the All field.
func (o *ApplyPatchesRequest) SetAll(v bool) {
	o.All.Set(&v)
}
// SetAllNil sets the value for All to be an explicit nil
func (o *ApplyPatchesRequest) SetAllNil() {
	o.All.Set(nil)
}

// UnsetAll ensures that no value is present for All, not even an explicit nil
func (o *ApplyPatchesRequest) UnsetAll() {
	o.All.Unset()
}

func (o ApplyPatchesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Service.IsSet() {
		toSerialize["service"] = o.Service.Get()
	}
	if o.All.IsSet() {
		toSerialize["all"] = o.All.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableApplyPatchesRequest struct {
	value *ApplyPatchesRequest
	isSet bool
}

func (v NullableApplyPatchesRequest) Get() *ApplyPatchesRequest {
	return v.value
}

func (v *NullableApplyPatchesRequest) Set(val *ApplyPatchesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApplyPatchesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApplyPatchesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplyPatchesRequest(val *ApplyPatchesRequest) *NullableApplyPatchesRequest {
	return &NullableApplyPatchesRequest{value: val, isSet: true}
}

func (v NullableApplyPatchesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplyPatchesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o ApplyPatchesRequest) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}