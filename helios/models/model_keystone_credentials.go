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

// KeystoneCredentials Specifies user credentials of a Keystone server.
type KeystoneCredentials struct {
	// Specifies parameters related to Keystone administrator.
	AdminCreds KeystoneAdminParams `json:"adminCreds"`
	// Specifies parameters related to Keystone scope.
	Scope KeystoneScopeParams `json:"scope"`
}

// NewKeystoneCredentials instantiates a new KeystoneCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeystoneCredentials(adminCreds KeystoneAdminParams, scope KeystoneScopeParams) *KeystoneCredentials {
	this := KeystoneCredentials{}
	this.AdminCreds = adminCreds
	this.Scope = scope
	return &this
}

// NewKeystoneCredentialsWithDefaults instantiates a new KeystoneCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeystoneCredentialsWithDefaults() *KeystoneCredentials {
	this := KeystoneCredentials{}
	return &this
}

// GetAdminCreds returns the AdminCreds field value
func (o *KeystoneCredentials) GetAdminCreds() KeystoneAdminParams {
	if o == nil {
		var ret KeystoneAdminParams
		return ret
	}

	return o.AdminCreds
}

// GetAdminCredsOk returns a tuple with the AdminCreds field value
// and a boolean to check if the value has been set.
func (o *KeystoneCredentials) GetAdminCredsOk() (*KeystoneAdminParams, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AdminCreds, true
}

// SetAdminCreds sets field value
func (o *KeystoneCredentials) SetAdminCreds(v KeystoneAdminParams) {
	o.AdminCreds = v
}

// GetScope returns the Scope field value
func (o *KeystoneCredentials) GetScope() KeystoneScopeParams {
	if o == nil {
		var ret KeystoneScopeParams
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *KeystoneCredentials) GetScopeOk() (*KeystoneScopeParams, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *KeystoneCredentials) SetScope(v KeystoneScopeParams) {
	o.Scope = v
}

func (o KeystoneCredentials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["adminCreds"] = o.AdminCreds
	}
	if true {
		toSerialize["scope"] = o.Scope
	}
	return json.Marshal(toSerialize)
}

type NullableKeystoneCredentials struct {
	value *KeystoneCredentials
	isSet bool
}

func (v NullableKeystoneCredentials) Get() *KeystoneCredentials {
	return v.value
}

func (v *NullableKeystoneCredentials) Set(val *KeystoneCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableKeystoneCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableKeystoneCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeystoneCredentials(val *KeystoneCredentials) *NullableKeystoneCredentials {
	return &NullableKeystoneCredentials{value: val, isSet: true}
}

func (v NullableKeystoneCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeystoneCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


