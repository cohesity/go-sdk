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

// ElastifileRegistrationParams Specifies parameters to register an Elastifile Source.
type ElastifileRegistrationParams struct {
	// Specifies the Hostname or IP Address Endpoint for the Elastifile Source.
	Endpoint NullableString `json:"endpoint"`
	Credentials Credentials `json:"credentials"`
	ThrottlingConfig *NasThrottlingConfig `json:"throttlingConfig,omitempty"`
}

// NewElastifileRegistrationParams instantiates a new ElastifileRegistrationParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewElastifileRegistrationParams(endpoint NullableString, credentials Credentials) *ElastifileRegistrationParams {
	this := ElastifileRegistrationParams{}
	this.Endpoint = endpoint
	this.Credentials = credentials
	return &this
}

// NewElastifileRegistrationParamsWithDefaults instantiates a new ElastifileRegistrationParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewElastifileRegistrationParamsWithDefaults() *ElastifileRegistrationParams {
	this := ElastifileRegistrationParams{}
	return &this
}

// GetEndpoint returns the Endpoint field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ElastifileRegistrationParams) GetEndpoint() string {
	if o == nil || o.Endpoint.Get() == nil {
		var ret string
		return ret
	}

	return *o.Endpoint.Get()
}

// GetEndpointOk returns a tuple with the Endpoint field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ElastifileRegistrationParams) GetEndpointOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Endpoint.Get(), o.Endpoint.IsSet()
}

// SetEndpoint sets field value
func (o *ElastifileRegistrationParams) SetEndpoint(v string) {
	o.Endpoint.Set(&v)
}

// GetCredentials returns the Credentials field value
func (o *ElastifileRegistrationParams) GetCredentials() Credentials {
	if o == nil {
		var ret Credentials
		return ret
	}

	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
func (o *ElastifileRegistrationParams) GetCredentialsOk() (*Credentials, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Credentials, true
}

// SetCredentials sets field value
func (o *ElastifileRegistrationParams) SetCredentials(v Credentials) {
	o.Credentials = v
}

// GetThrottlingConfig returns the ThrottlingConfig field value if set, zero value otherwise.
func (o *ElastifileRegistrationParams) GetThrottlingConfig() NasThrottlingConfig {
	if o == nil || o.ThrottlingConfig == nil {
		var ret NasThrottlingConfig
		return ret
	}
	return *o.ThrottlingConfig
}

// GetThrottlingConfigOk returns a tuple with the ThrottlingConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElastifileRegistrationParams) GetThrottlingConfigOk() (*NasThrottlingConfig, bool) {
	if o == nil || o.ThrottlingConfig == nil {
		return nil, false
	}
	return o.ThrottlingConfig, true
}

// HasThrottlingConfig returns a boolean if a field has been set.
func (o *ElastifileRegistrationParams) HasThrottlingConfig() bool {
	if o != nil && o.ThrottlingConfig != nil {
		return true
	}

	return false
}

// SetThrottlingConfig gets a reference to the given NasThrottlingConfig and assigns it to the ThrottlingConfig field.
func (o *ElastifileRegistrationParams) SetThrottlingConfig(v NasThrottlingConfig) {
	o.ThrottlingConfig = &v
}

func (o ElastifileRegistrationParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["endpoint"] = o.Endpoint.Get()
	}
	if true {
		toSerialize["credentials"] = o.Credentials
	}
	if o.ThrottlingConfig != nil {
		toSerialize["throttlingConfig"] = o.ThrottlingConfig
	}
	return json.Marshal(toSerialize)
}

type NullableElastifileRegistrationParams struct {
	value *ElastifileRegistrationParams
	isSet bool
}

func (v NullableElastifileRegistrationParams) Get() *ElastifileRegistrationParams {
	return v.value
}

func (v *NullableElastifileRegistrationParams) Set(val *ElastifileRegistrationParams) {
	v.value = val
	v.isSet = true
}

func (v NullableElastifileRegistrationParams) IsSet() bool {
	return v.isSet
}

func (v *NullableElastifileRegistrationParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableElastifileRegistrationParams(val *ElastifileRegistrationParams) *NullableElastifileRegistrationParams {
	return &NullableElastifileRegistrationParams{value: val, isSet: true}
}

func (v NullableElastifileRegistrationParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableElastifileRegistrationParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o ElastifileRegistrationParams) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}