/*
Cohesity REST API

Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the PrivateKeyCredentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrivateKeyCredentials{}

// PrivateKeyCredentials Specifies the credentials for certificate based authentication.
type PrivateKeyCredentials struct {
	// Specifies the passphrase to access target entity.
	Passphrase *string `json:"passphrase,omitempty"`
	// Specifies the private key to access target entity.
	PrivateKey string `json:"privateKey"`
	// Specifies the userId to access target entity.
	UserId string `json:"userId"`
}

type _PrivateKeyCredentials PrivateKeyCredentials

// NewPrivateKeyCredentials instantiates a new PrivateKeyCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivateKeyCredentials(privateKey string, userId string) *PrivateKeyCredentials {
	this := PrivateKeyCredentials{}
	this.PrivateKey = privateKey
	this.UserId = userId
	return &this
}

// NewPrivateKeyCredentialsWithDefaults instantiates a new PrivateKeyCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivateKeyCredentialsWithDefaults() *PrivateKeyCredentials {
	this := PrivateKeyCredentials{}
	return &this
}

// GetPassphrase returns the Passphrase field value if set, zero value otherwise.
func (o *PrivateKeyCredentials) GetPassphrase() string {
	if o == nil || IsNil(o.Passphrase) {
		var ret string
		return ret
	}
	return *o.Passphrase
}

// GetPassphraseOk returns a tuple with the Passphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateKeyCredentials) GetPassphraseOk() (*string, bool) {
	if o == nil || IsNil(o.Passphrase) {
		return nil, false
	}
	return o.Passphrase, true
}

// HasPassphrase returns a boolean if a field has been set.
func (o *PrivateKeyCredentials) HasPassphrase() bool {
	if o != nil && !IsNil(o.Passphrase) {
		return true
	}

	return false
}

// SetPassphrase gets a reference to the given string and assigns it to the Passphrase field.
func (o *PrivateKeyCredentials) SetPassphrase(v string) {
	o.Passphrase = &v
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
	if o == nil {
		return nil, false
	}
	return &o.PrivateKey, true
}

// SetPrivateKey sets field value
func (o *PrivateKeyCredentials) SetPrivateKey(v string) {
	o.PrivateKey = v
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
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *PrivateKeyCredentials) SetUserId(v string) {
	o.UserId = v
}

func (o PrivateKeyCredentials) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrivateKeyCredentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Passphrase) {
		toSerialize["passphrase"] = o.Passphrase
	}
	toSerialize["privateKey"] = o.PrivateKey
	toSerialize["userId"] = o.UserId
	return toSerialize, nil
}

func (o *PrivateKeyCredentials) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"privateKey",
		"userId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPrivateKeyCredentials := _PrivateKeyCredentials{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPrivateKeyCredentials)

	if err != nil {
		return err
	}

	*o = PrivateKeyCredentials(varPrivateKeyCredentials)

	return err
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


