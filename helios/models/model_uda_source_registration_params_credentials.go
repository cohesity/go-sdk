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

// UdaSourceRegistrationParamsCredentials Specifies credentials to access the Universal Data Adapter source. For e.g.: To perform backup and recovery tasks with Oracle Recovery Manager (RMAN), specify credentials for a user having 'SYSDBA' or 'SYSBACKUP' administrative privilege.
type UdaSourceRegistrationParamsCredentials struct {
	// Specifies the password to access target entity.
	Password NullableString `json:"password,omitempty"`
	// Specifies the username to access target entity.
	Username NullableString `json:"username,omitempty"`
}

// NewUdaSourceRegistrationParamsCredentials instantiates a new UdaSourceRegistrationParamsCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUdaSourceRegistrationParamsCredentials() *UdaSourceRegistrationParamsCredentials {
	this := UdaSourceRegistrationParamsCredentials{}
	return &this
}

// NewUdaSourceRegistrationParamsCredentialsWithDefaults instantiates a new UdaSourceRegistrationParamsCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUdaSourceRegistrationParamsCredentialsWithDefaults() *UdaSourceRegistrationParamsCredentials {
	this := UdaSourceRegistrationParamsCredentials{}
	return &this
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UdaSourceRegistrationParamsCredentials) GetPassword() string {
	if o == nil || o.Password.Get() == nil {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UdaSourceRegistrationParamsCredentials) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *UdaSourceRegistrationParamsCredentials) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *UdaSourceRegistrationParamsCredentials) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *UdaSourceRegistrationParamsCredentials) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *UdaSourceRegistrationParamsCredentials) UnsetPassword() {
	o.Password.Unset()
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UdaSourceRegistrationParamsCredentials) GetUsername() string {
	if o == nil || o.Username.Get() == nil {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UdaSourceRegistrationParamsCredentials) GetUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *UdaSourceRegistrationParamsCredentials) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *UdaSourceRegistrationParamsCredentials) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *UdaSourceRegistrationParamsCredentials) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *UdaSourceRegistrationParamsCredentials) UnsetUsername() {
	o.Username.Unset()
}

func (o UdaSourceRegistrationParamsCredentials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableUdaSourceRegistrationParamsCredentials struct {
	value *UdaSourceRegistrationParamsCredentials
	isSet bool
}

func (v NullableUdaSourceRegistrationParamsCredentials) Get() *UdaSourceRegistrationParamsCredentials {
	return v.value
}

func (v *NullableUdaSourceRegistrationParamsCredentials) Set(val *UdaSourceRegistrationParamsCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableUdaSourceRegistrationParamsCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableUdaSourceRegistrationParamsCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUdaSourceRegistrationParamsCredentials(val *UdaSourceRegistrationParamsCredentials) *NullableUdaSourceRegistrationParamsCredentials {
	return &NullableUdaSourceRegistrationParamsCredentials{value: val, isSet: true}
}

func (v NullableUdaSourceRegistrationParamsCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUdaSourceRegistrationParamsCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


