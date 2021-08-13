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

// Office365AppCredentials Specifies credentials for office365 azure registered applications, used for office 365 source registration.
type Office365AppCredentials struct {
	// Specifies the application ID that the registration portal (apps.dev.microsoft.com) assigned.
	ClientId NullableString `json:"clientId,omitempty"`
	// Specifies the application secret that was created in app registration portal.
	ClientSecret NullableString `json:"clientSecret,omitempty"`
}

// NewOffice365AppCredentials instantiates a new Office365AppCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOffice365AppCredentials() *Office365AppCredentials {
	this := Office365AppCredentials{}
	return &this
}

// NewOffice365AppCredentialsWithDefaults instantiates a new Office365AppCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOffice365AppCredentialsWithDefaults() *Office365AppCredentials {
	this := Office365AppCredentials{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Office365AppCredentials) GetClientId() string {
	if o == nil || o.ClientId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClientId.Get()
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Office365AppCredentials) GetClientIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ClientId.Get(), o.ClientId.IsSet()
}

// HasClientId returns a boolean if a field has been set.
func (o *Office365AppCredentials) HasClientId() bool {
	if o != nil && o.ClientId.IsSet() {
		return true
	}

	return false
}

// SetClientId gets a reference to the given NullableString and assigns it to the ClientId field.
func (o *Office365AppCredentials) SetClientId(v string) {
	o.ClientId.Set(&v)
}
// SetClientIdNil sets the value for ClientId to be an explicit nil
func (o *Office365AppCredentials) SetClientIdNil() {
	o.ClientId.Set(nil)
}

// UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
func (o *Office365AppCredentials) UnsetClientId() {
	o.ClientId.Unset()
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Office365AppCredentials) GetClientSecret() string {
	if o == nil || o.ClientSecret.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClientSecret.Get()
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Office365AppCredentials) GetClientSecretOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ClientSecret.Get(), o.ClientSecret.IsSet()
}

// HasClientSecret returns a boolean if a field has been set.
func (o *Office365AppCredentials) HasClientSecret() bool {
	if o != nil && o.ClientSecret.IsSet() {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given NullableString and assigns it to the ClientSecret field.
func (o *Office365AppCredentials) SetClientSecret(v string) {
	o.ClientSecret.Set(&v)
}
// SetClientSecretNil sets the value for ClientSecret to be an explicit nil
func (o *Office365AppCredentials) SetClientSecretNil() {
	o.ClientSecret.Set(nil)
}

// UnsetClientSecret ensures that no value is present for ClientSecret, not even an explicit nil
func (o *Office365AppCredentials) UnsetClientSecret() {
	o.ClientSecret.Unset()
}

func (o Office365AppCredentials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId.IsSet() {
		toSerialize["clientId"] = o.ClientId.Get()
	}
	if o.ClientSecret.IsSet() {
		toSerialize["clientSecret"] = o.ClientSecret.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableOffice365AppCredentials struct {
	value *Office365AppCredentials
	isSet bool
}

func (v NullableOffice365AppCredentials) Get() *Office365AppCredentials {
	return v.value
}

func (v *NullableOffice365AppCredentials) Set(val *Office365AppCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableOffice365AppCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableOffice365AppCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOffice365AppCredentials(val *Office365AppCredentials) *NullableOffice365AppCredentials {
	return &NullableOffice365AppCredentials{value: val, isSet: true}
}

func (v NullableOffice365AppCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOffice365AppCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o Office365AppCredentials) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}