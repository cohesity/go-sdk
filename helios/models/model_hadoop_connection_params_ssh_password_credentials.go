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

// HadoopConnectionParamsSshPasswordCredentials SSH username + password required for reading configuration file and for scp backup.Either 'sshPasswordCredential' or 'sshPrivateKeyCredential' are required.
type HadoopConnectionParamsSshPasswordCredentials struct {
	// SSH password.
	Password string `json:"password"`
	// SSH username.
	Username string `json:"username"`
}

// NewHadoopConnectionParamsSshPasswordCredentials instantiates a new HadoopConnectionParamsSshPasswordCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHadoopConnectionParamsSshPasswordCredentials(password string, username string) *HadoopConnectionParamsSshPasswordCredentials {
	this := HadoopConnectionParamsSshPasswordCredentials{}
	this.Password = password
	this.Username = username
	return &this
}

// NewHadoopConnectionParamsSshPasswordCredentialsWithDefaults instantiates a new HadoopConnectionParamsSshPasswordCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHadoopConnectionParamsSshPasswordCredentialsWithDefaults() *HadoopConnectionParamsSshPasswordCredentials {
	this := HadoopConnectionParamsSshPasswordCredentials{}
	return &this
}

// GetPassword returns the Password field value
func (o *HadoopConnectionParamsSshPasswordCredentials) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *HadoopConnectionParamsSshPasswordCredentials) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *HadoopConnectionParamsSshPasswordCredentials) SetPassword(v string) {
	o.Password = v
}

// GetUsername returns the Username field value
func (o *HadoopConnectionParamsSshPasswordCredentials) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *HadoopConnectionParamsSshPasswordCredentials) GetUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *HadoopConnectionParamsSshPasswordCredentials) SetUsername(v string) {
	o.Username = v
}

func (o HadoopConnectionParamsSshPasswordCredentials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["password"] = o.Password
	}
	if true {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableHadoopConnectionParamsSshPasswordCredentials struct {
	value *HadoopConnectionParamsSshPasswordCredentials
	isSet bool
}

func (v NullableHadoopConnectionParamsSshPasswordCredentials) Get() *HadoopConnectionParamsSshPasswordCredentials {
	return v.value
}

func (v *NullableHadoopConnectionParamsSshPasswordCredentials) Set(val *HadoopConnectionParamsSshPasswordCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableHadoopConnectionParamsSshPasswordCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableHadoopConnectionParamsSshPasswordCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHadoopConnectionParamsSshPasswordCredentials(val *HadoopConnectionParamsSshPasswordCredentials) *NullableHadoopConnectionParamsSshPasswordCredentials {
	return &NullableHadoopConnectionParamsSshPasswordCredentials{value: val, isSet: true}
}

func (v NullableHadoopConnectionParamsSshPasswordCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHadoopConnectionParamsSshPasswordCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


