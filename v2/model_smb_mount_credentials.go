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

// checks if the SmbMountCredentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmbMountCredentials{}

// SmbMountCredentials Specifies the credentials to mount a view.
type SmbMountCredentials struct {
	// Specifies the password to access target entity.
	Password string `json:"password"`
	// Specifies the username to access target entity.
	Username string `json:"username"`
}

type _SmbMountCredentials SmbMountCredentials

// NewSmbMountCredentials instantiates a new SmbMountCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmbMountCredentials(password string, username string) *SmbMountCredentials {
	this := SmbMountCredentials{}
	this.Password = password
	this.Username = username
	return &this
}

// NewSmbMountCredentialsWithDefaults instantiates a new SmbMountCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmbMountCredentialsWithDefaults() *SmbMountCredentials {
	this := SmbMountCredentials{}
	return &this
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
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *SmbMountCredentials) SetPassword(v string) {
	o.Password = v
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
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *SmbMountCredentials) SetUsername(v string) {
	o.Username = v
}

func (o SmbMountCredentials) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmbMountCredentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["password"] = o.Password
	toSerialize["username"] = o.Username
	return toSerialize, nil
}

func (o *SmbMountCredentials) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"password",
		"username",
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

	varSmbMountCredentials := _SmbMountCredentials{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSmbMountCredentials)

	if err != nil {
		return err
	}

	*o = SmbMountCredentials(varSmbMountCredentials)

	return err
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


