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

// CipherSuite Cipher Suite
type CipherSuite struct {
	// Specifies the cipher suite used for TLS handshake.
	Cipher *string `json:"cipher,omitempty"`
}

// NewCipherSuite instantiates a new CipherSuite object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCipherSuite() *CipherSuite {
	this := CipherSuite{}
	return &this
}

// NewCipherSuiteWithDefaults instantiates a new CipherSuite object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCipherSuiteWithDefaults() *CipherSuite {
	this := CipherSuite{}
	return &this
}

// GetCipher returns the Cipher field value if set, zero value otherwise.
func (o *CipherSuite) GetCipher() string {
	if o == nil || o.Cipher == nil {
		var ret string
		return ret
	}
	return *o.Cipher
}

// GetCipherOk returns a tuple with the Cipher field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CipherSuite) GetCipherOk() (*string, bool) {
	if o == nil || o.Cipher == nil {
		return nil, false
	}
	return o.Cipher, true
}

// HasCipher returns a boolean if a field has been set.
func (o *CipherSuite) HasCipher() bool {
	if o != nil && o.Cipher != nil {
		return true
	}

	return false
}

// SetCipher gets a reference to the given string and assigns it to the Cipher field.
func (o *CipherSuite) SetCipher(v string) {
	o.Cipher = &v
}

func (o CipherSuite) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Cipher != nil {
		toSerialize["cipher"] = o.Cipher
	}
	return json.Marshal(toSerialize)
}

type NullableCipherSuite struct {
	value *CipherSuite
	isSet bool
}

func (v NullableCipherSuite) Get() *CipherSuite {
	return v.value
}

func (v *NullableCipherSuite) Set(val *CipherSuite) {
	v.value = val
	v.isSet = true
}

func (v NullableCipherSuite) IsSet() bool {
	return v.isSet
}

func (v *NullableCipherSuite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCipherSuite(val *CipherSuite) *NullableCipherSuite {
	return &NullableCipherSuite{value: val, isSet: true}
}

func (v NullableCipherSuite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCipherSuite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


