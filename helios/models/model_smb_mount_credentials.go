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

// SmbMountCredentials Specifies the credentials to mount a view.
type SmbMountCredentials struct {
	// Specifies the username to access target entity.
	Username string `json:"username"`
	// Specifies the password to access target entity.
	Password string `json:"password"`
}

// NewSmbMountCredentials instantiates a new SmbMountCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmbMountCredentials(username string, password string) *SmbMountCredentials {
	this := SmbMountCredentials{}
	this.Username = username
	this.Password = password
	return &this
}

// NewSmbMountCredentialsWithDefaults instantiates a new SmbMountCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmbMountCredentialsWithDefaults() *SmbMountCredentials {
	this := SmbMountCredentials{}
	return &this
}

// GetUsername returns the Username field value
func (o *SmbMountCredentials) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *SmbMountCredentials) GetUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *SmbMountCredentials) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value
func (o *SmbMountCredentials) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *SmbMountCredentials) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *SmbMountCredentials) SetPassword(v string) {
	o.Password = v
}

func (o SmbMountCredentials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["username"] = o.Username
	}
	if true {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableSmbMountCredentials struct {
	value *SmbMountCredentials
	isSet bool
}

func (v NullableSmbMountCredentials) Get() *SmbMountCredentials {
	return v.value
}

func (v *NullableSmbMountCredentials) Set(val *SmbMountCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableSmbMountCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableSmbMountCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmbMountCredentials(val *SmbMountCredentials) *NullableSmbMountCredentials {
	return &NullableSmbMountCredentials{value: val, isSet: true}
}

func (v NullableSmbMountCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmbMountCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


