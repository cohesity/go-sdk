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

// HeliosAwsRequest Specifies a AWS registration request.
type HeliosAwsRequest struct {
	// Specifies the registration token sent from Subscription manager.
	Token *string `json:"token,omitempty"`
}

// NewHeliosAwsRequest instantiates a new HeliosAwsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeliosAwsRequest() *HeliosAwsRequest {
	this := HeliosAwsRequest{}
	return &this
}

// NewHeliosAwsRequestWithDefaults instantiates a new HeliosAwsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeliosAwsRequestWithDefaults() *HeliosAwsRequest {
	this := HeliosAwsRequest{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *HeliosAwsRequest) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeliosAwsRequest) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *HeliosAwsRequest) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *HeliosAwsRequest) SetToken(v string) {
	o.Token = &v
}

func (o HeliosAwsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullableHeliosAwsRequest struct {
	value *HeliosAwsRequest
	isSet bool
}

func (v NullableHeliosAwsRequest) Get() *HeliosAwsRequest {
	return v.value
}

func (v *NullableHeliosAwsRequest) Set(val *HeliosAwsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableHeliosAwsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableHeliosAwsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHeliosAwsRequest(val *HeliosAwsRequest) *NullableHeliosAwsRequest {
	return &NullableHeliosAwsRequest{value: val, isSet: true}
}

func (v NullableHeliosAwsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHeliosAwsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


