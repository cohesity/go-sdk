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

// VerifyTotpRequest Holds the Totp code to be verified.
type VerifyTotpRequest struct {
	// Specifies the Totp code.
	TotpCode NullableString `json:"totpCode,omitempty"`
}

// NewVerifyTotpRequest instantiates a new VerifyTotpRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifyTotpRequest() *VerifyTotpRequest {
	this := VerifyTotpRequest{}
	return &this
}

// NewVerifyTotpRequestWithDefaults instantiates a new VerifyTotpRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifyTotpRequestWithDefaults() *VerifyTotpRequest {
	this := VerifyTotpRequest{}
	return &this
}

// GetTotpCode returns the TotpCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VerifyTotpRequest) GetTotpCode() string {
	if o == nil || o.TotpCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.TotpCode.Get()
}

// GetTotpCodeOk returns a tuple with the TotpCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VerifyTotpRequest) GetTotpCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TotpCode.Get(), o.TotpCode.IsSet()
}

// HasTotpCode returns a boolean if a field has been set.
func (o *VerifyTotpRequest) HasTotpCode() bool {
	if o != nil && o.TotpCode.IsSet() {
		return true
	}

	return false
}

// SetTotpCode gets a reference to the given NullableString and assigns it to the TotpCode field.
func (o *VerifyTotpRequest) SetTotpCode(v string) {
	o.TotpCode.Set(&v)
}
// SetTotpCodeNil sets the value for TotpCode to be an explicit nil
func (o *VerifyTotpRequest) SetTotpCodeNil() {
	o.TotpCode.Set(nil)
}

// UnsetTotpCode ensures that no value is present for TotpCode, not even an explicit nil
func (o *VerifyTotpRequest) UnsetTotpCode() {
	o.TotpCode.Unset()
}

func (o VerifyTotpRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TotpCode.IsSet() {
		toSerialize["totpCode"] = o.TotpCode.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableVerifyTotpRequest struct {
	value *VerifyTotpRequest
	isSet bool
}

func (v NullableVerifyTotpRequest) Get() *VerifyTotpRequest {
	return v.value
}

func (v *NullableVerifyTotpRequest) Set(val *VerifyTotpRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifyTotpRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifyTotpRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifyTotpRequest(val *VerifyTotpRequest) *NullableVerifyTotpRequest {
	return &NullableVerifyTotpRequest{value: val, isSet: true}
}

func (v NullableVerifyTotpRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifyTotpRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


