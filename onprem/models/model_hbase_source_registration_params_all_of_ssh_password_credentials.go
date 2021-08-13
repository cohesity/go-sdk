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

// HbaseSourceRegistrationParamsAllOfSshPasswordCredentials SSH username + password required for reading configuration file. Either 'sshPasswordCredentials' or 'sshPrivateKeyCredentials' are required.
type HbaseSourceRegistrationParamsAllOfSshPasswordCredentials struct {
	// SSH password.
	Password string `json:"password"`
	// SSH username.
	Username string `json:"username"`
}

// NewHbaseSourceRegistrationParamsAllOfSshPasswordCredentials instantiates a new HbaseSourceRegistrationParamsAllOfSshPasswordCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHbaseSourceRegistrationParamsAllOfSshPasswordCredentials(password string, username string) *HbaseSourceRegistrationParamsAllOfSshPasswordCredentials {
	this := HbaseSourceRegistrationParamsAllOfSshPasswordCredentials{}
	this.Password = password
	this.Username = username
	return &this
}

// NewHbaseSourceRegistrationParamsAllOfSshPasswordCredentialsWithDefaults instantiates a new HbaseSourceRegistrationParamsAllOfSshPasswordCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHbaseSourceRegistrationParamsAllOfSshPasswordCredentialsWithDefaults() *HbaseSourceRegistrationParamsAllOfSshPasswordCredentials {
	this := HbaseSourceRegistrationParamsAllOfSshPasswordCredentials{}
	return &this
}

// GetPassword returns the Password field value
func (o *HbaseSourceRegistrationParamsAllOfSshPasswordCredentials) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *HbaseSourceRegistrationParamsAllOfSshPasswordCredentials) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *HbaseSourceRegistrationParamsAllOfSshPasswordCredentials) SetPassword(v string) {
	o.Password = v
}

// GetUsername returns the Username field value
func (o *HbaseSourceRegistrationParamsAllOfSshPasswordCredentials) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *HbaseSourceRegistrationParamsAllOfSshPasswordCredentials) GetUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *HbaseSourceRegistrationParamsAllOfSshPasswordCredentials) SetUsername(v string) {
	o.Username = v
}

func (o HbaseSourceRegistrationParamsAllOfSshPasswordCredentials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["password"] = o.Password
	}
	if true {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableHbaseSourceRegistrationParamsAllOfSshPasswordCredentials struct {
	value *HbaseSourceRegistrationParamsAllOfSshPasswordCredentials
	isSet bool
}

func (v NullableHbaseSourceRegistrationParamsAllOfSshPasswordCredentials) Get() *HbaseSourceRegistrationParamsAllOfSshPasswordCredentials {
	return v.value
}

func (v *NullableHbaseSourceRegistrationParamsAllOfSshPasswordCredentials) Set(val *HbaseSourceRegistrationParamsAllOfSshPasswordCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableHbaseSourceRegistrationParamsAllOfSshPasswordCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableHbaseSourceRegistrationParamsAllOfSshPasswordCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHbaseSourceRegistrationParamsAllOfSshPasswordCredentials(val *HbaseSourceRegistrationParamsAllOfSshPasswordCredentials) *NullableHbaseSourceRegistrationParamsAllOfSshPasswordCredentials {
	return &NullableHbaseSourceRegistrationParamsAllOfSshPasswordCredentials{value: val, isSet: true}
}

func (v NullableHbaseSourceRegistrationParamsAllOfSshPasswordCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHbaseSourceRegistrationParamsAllOfSshPasswordCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o HbaseSourceRegistrationParamsAllOfSshPasswordCredentials) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}