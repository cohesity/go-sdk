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

// VcdConnectionParams Specifies the parameters to connect to a seed node and fetch information from its config file.
type VcdConnectionParams struct {
	// IP or hostname of the source.
	Host NullableString `json:"host"`
	// Specifies the username to access target entity.
	Username string `json:"username"`
	// Specifies the password to access target entity.
	Password string `json:"password"`
}

// NewVcdConnectionParams instantiates a new VcdConnectionParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcdConnectionParams(host NullableString, username string, password string) *VcdConnectionParams {
	this := VcdConnectionParams{}
	this.Host = host
	this.Username = username
	this.Password = password
	return &this
}

// NewVcdConnectionParamsWithDefaults instantiates a new VcdConnectionParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcdConnectionParamsWithDefaults() *VcdConnectionParams {
	this := VcdConnectionParams{}
	return &this
}

// GetHost returns the Host field value
// If the value is explicit nil, the zero value for string will be returned
func (o *VcdConnectionParams) GetHost() string {
	if o == nil || o.Host.Get() == nil {
		var ret string
		return ret
	}

	return *o.Host.Get()
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VcdConnectionParams) GetHostOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Host.Get(), o.Host.IsSet()
}

// SetHost sets field value
func (o *VcdConnectionParams) SetHost(v string) {
	o.Host.Set(&v)
}

// GetUsername returns the Username field value
func (o *VcdConnectionParams) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *VcdConnectionParams) GetUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *VcdConnectionParams) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value
func (o *VcdConnectionParams) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *VcdConnectionParams) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *VcdConnectionParams) SetPassword(v string) {
	o.Password = v
}

func (o VcdConnectionParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["host"] = o.Host.Get()
	}
	if true {
		toSerialize["username"] = o.Username
	}
	if true {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableVcdConnectionParams struct {
	value *VcdConnectionParams
	isSet bool
}

func (v NullableVcdConnectionParams) Get() *VcdConnectionParams {
	return v.value
}

func (v *NullableVcdConnectionParams) Set(val *VcdConnectionParams) {
	v.value = val
	v.isSet = true
}

func (v NullableVcdConnectionParams) IsSet() bool {
	return v.isSet
}

func (v *NullableVcdConnectionParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcdConnectionParams(val *VcdConnectionParams) *NullableVcdConnectionParams {
	return &NullableVcdConnectionParams{value: val, isSet: true}
}

func (v NullableVcdConnectionParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcdConnectionParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


