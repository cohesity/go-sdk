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

// LdapConnectionStatus Specifies the LDAP connection status.
type LdapConnectionStatus struct {
	// Specifies the LDAP connection status.
	Type *string `json:"type,omitempty"`
}

// NewLdapConnectionStatus instantiates a new LdapConnectionStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLdapConnectionStatus() *LdapConnectionStatus {
	this := LdapConnectionStatus{}
	return &this
}

// NewLdapConnectionStatusWithDefaults instantiates a new LdapConnectionStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLdapConnectionStatusWithDefaults() *LdapConnectionStatus {
	this := LdapConnectionStatus{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LdapConnectionStatus) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapConnectionStatus) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LdapConnectionStatus) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *LdapConnectionStatus) SetType(v string) {
	o.Type = &v
}

func (o LdapConnectionStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableLdapConnectionStatus struct {
	value *LdapConnectionStatus
	isSet bool
}

func (v NullableLdapConnectionStatus) Get() *LdapConnectionStatus {
	return v.value
}

func (v *NullableLdapConnectionStatus) Set(val *LdapConnectionStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableLdapConnectionStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableLdapConnectionStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLdapConnectionStatus(val *LdapConnectionStatus) *NullableLdapConnectionStatus {
	return &NullableLdapConnectionStatus{value: val, isSet: true}
}

func (v NullableLdapConnectionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLdapConnectionStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


