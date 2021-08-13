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

// PrivateKeyCredentials Specifies the credentials for certificate based authentication.
type PrivateKeyCredentials struct {
	// Specifies the userId to access target entity.
	UserId string `json:"userId"`
	// Specifies the private key to access target entity.
	PrivateKey string `json:"privateKey"`
	// Specifies the passphrase to access target entity.
	Passphrase *string `json:"passphrase,omitempty"`
}

// NewPrivateKeyCredentials instantiates a new PrivateKeyCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivateKeyCredentials(userId string, privateKey string) *PrivateKeyCredentials {
	this := PrivateKeyCredentials{}
	this.UserId = userId
	this.PrivateKey = privateKey
	return &this
}

// NewPrivateKeyCredentialsWithDefaults instantiates a new PrivateKeyCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivateKeyCredentialsWithDefaults() *PrivateKeyCredentials {
	this := PrivateKeyCredentials{}
	return &this
}

// GetUserId returns the UserId field value
func (o *PrivateKeyCredentials) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *PrivateKeyCredentials) GetUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *PrivateKeyCredentials) SetUserId(v string) {
	o.UserId = v
}

// GetPrivateKey returns the PrivateKey field value
func (o *PrivateKeyCredentials) GetPrivateKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value
// and a boolean to check if the value has been set.
func (o *PrivateKeyCredentials) GetPrivateKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PrivateKey, true
}

// SetPrivateKey sets field value
func (o *PrivateKeyCredentials) SetPrivateKey(v string) {
	o.PrivateKey = v
}

// GetPassphrase returns the Passphrase field value if set, zero value otherwise.
func (o *PrivateKeyCredentials) GetPassphrase() string {
	if o == nil || o.Passphrase == nil {
		var ret string
		return ret
	}
	return *o.Passphrase
}

// GetPassphraseOk returns a tuple with the Passphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateKeyCredentials) GetPassphraseOk() (*string, bool) {
	if o == nil || o.Passphrase == nil {
		return nil, false
	}
	return o.Passphrase, true
}

// HasPassphrase returns a boolean if a field has been set.
func (o *PrivateKeyCredentials) HasPassphrase() bool {
	if o != nil && o.Passphrase != nil {
		return true
	}

	return false
}

// SetPassphrase gets a reference to the given string and assigns it to the Passphrase field.
func (o *PrivateKeyCredentials) SetPassphrase(v string) {
	o.Passphrase = &v
}

func (o PrivateKeyCredentials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["userId"] = o.UserId
	}
	if true {
		toSerialize["privateKey"] = o.PrivateKey
	}
	if o.Passphrase != nil {
		toSerialize["passphrase"] = o.Passphrase
	}
	return json.Marshal(toSerialize)
}

type NullablePrivateKeyCredentials struct {
	value *PrivateKeyCredentials
	isSet bool
}

func (v NullablePrivateKeyCredentials) Get() *PrivateKeyCredentials {
	return v.value
}

func (v *NullablePrivateKeyCredentials) Set(val *PrivateKeyCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivateKeyCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivateKeyCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivateKeyCredentials(val *PrivateKeyCredentials) *NullablePrivateKeyCredentials {
	return &NullablePrivateKeyCredentials{value: val, isSet: true}
}

func (v NullablePrivateKeyCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivateKeyCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o PrivateKeyCredentials) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}