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

// HeliosLoginConfiguration Returns the Helios Configuration for Salesforce login
type HeliosLoginConfiguration struct {
	// Specifies the consumer key configured in Salesforce
	ClientId *string `json:"clientId,omitempty"`
	// Specifies the salesforce URL to be redirected to
	PortalUrl *string `json:"portalUrl,omitempty"`
	// Redirect URL after successful authentication
	RedirectUrl *string `json:"redirectUrl,omitempty"`
}

// NewHeliosLoginConfiguration instantiates a new HeliosLoginConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeliosLoginConfiguration() *HeliosLoginConfiguration {
	this := HeliosLoginConfiguration{}
	return &this
}

// NewHeliosLoginConfigurationWithDefaults instantiates a new HeliosLoginConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeliosLoginConfigurationWithDefaults() *HeliosLoginConfiguration {
	this := HeliosLoginConfiguration{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *HeliosLoginConfiguration) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeliosLoginConfiguration) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *HeliosLoginConfiguration) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *HeliosLoginConfiguration) SetClientId(v string) {
	o.ClientId = &v
}

// GetPortalUrl returns the PortalUrl field value if set, zero value otherwise.
func (o *HeliosLoginConfiguration) GetPortalUrl() string {
	if o == nil || o.PortalUrl == nil {
		var ret string
		return ret
	}
	return *o.PortalUrl
}

// GetPortalUrlOk returns a tuple with the PortalUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeliosLoginConfiguration) GetPortalUrlOk() (*string, bool) {
	if o == nil || o.PortalUrl == nil {
		return nil, false
	}
	return o.PortalUrl, true
}

// HasPortalUrl returns a boolean if a field has been set.
func (o *HeliosLoginConfiguration) HasPortalUrl() bool {
	if o != nil && o.PortalUrl != nil {
		return true
	}

	return false
}

// SetPortalUrl gets a reference to the given string and assigns it to the PortalUrl field.
func (o *HeliosLoginConfiguration) SetPortalUrl(v string) {
	o.PortalUrl = &v
}

// GetRedirectUrl returns the RedirectUrl field value if set, zero value otherwise.
func (o *HeliosLoginConfiguration) GetRedirectUrl() string {
	if o == nil || o.RedirectUrl == nil {
		var ret string
		return ret
	}
	return *o.RedirectUrl
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeliosLoginConfiguration) GetRedirectUrlOk() (*string, bool) {
	if o == nil || o.RedirectUrl == nil {
		return nil, false
	}
	return o.RedirectUrl, true
}

// HasRedirectUrl returns a boolean if a field has been set.
func (o *HeliosLoginConfiguration) HasRedirectUrl() bool {
	if o != nil && o.RedirectUrl != nil {
		return true
	}

	return false
}

// SetRedirectUrl gets a reference to the given string and assigns it to the RedirectUrl field.
func (o *HeliosLoginConfiguration) SetRedirectUrl(v string) {
	o.RedirectUrl = &v
}

func (o HeliosLoginConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["clientId"] = o.ClientId
	}
	if o.PortalUrl != nil {
		toSerialize["portalUrl"] = o.PortalUrl
	}
	if o.RedirectUrl != nil {
		toSerialize["redirectUrl"] = o.RedirectUrl
	}
	return json.Marshal(toSerialize)
}

type NullableHeliosLoginConfiguration struct {
	value *HeliosLoginConfiguration
	isSet bool
}

func (v NullableHeliosLoginConfiguration) Get() *HeliosLoginConfiguration {
	return v.value
}

func (v *NullableHeliosLoginConfiguration) Set(val *HeliosLoginConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableHeliosLoginConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableHeliosLoginConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHeliosLoginConfiguration(val *HeliosLoginConfiguration) *NullableHeliosLoginConfiguration {
	return &NullableHeliosLoginConfiguration{value: val, isSet: true}
}

func (v NullableHeliosLoginConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHeliosLoginConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


