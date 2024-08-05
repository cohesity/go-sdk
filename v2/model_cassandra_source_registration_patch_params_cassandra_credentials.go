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

// checks if the CassandraSourceRegistrationPatchParamsCassandraCredentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CassandraSourceRegistrationPatchParamsCassandraCredentials{}

// CassandraSourceRegistrationPatchParamsCassandraCredentials Cassandra Credentials for this cluster.
type CassandraSourceRegistrationPatchParamsCassandraCredentials struct {
	// Cassandra password.
	Password string `json:"password"`
	// Cassandra username.
	Username string `json:"username"`
}

type _CassandraSourceRegistrationPatchParamsCassandraCredentials CassandraSourceRegistrationPatchParamsCassandraCredentials

// NewCassandraSourceRegistrationPatchParamsCassandraCredentials instantiates a new CassandraSourceRegistrationPatchParamsCassandraCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCassandraSourceRegistrationPatchParamsCassandraCredentials(password string, username string) *CassandraSourceRegistrationPatchParamsCassandraCredentials {
	this := CassandraSourceRegistrationPatchParamsCassandraCredentials{}
	this.Password = password
	this.Username = username
	return &this
}

// NewCassandraSourceRegistrationPatchParamsCassandraCredentialsWithDefaults instantiates a new CassandraSourceRegistrationPatchParamsCassandraCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCassandraSourceRegistrationPatchParamsCassandraCredentialsWithDefaults() *CassandraSourceRegistrationPatchParamsCassandraCredentials {
	this := CassandraSourceRegistrationPatchParamsCassandraCredentials{}
	return &this
}

// GetPassword returns the Password field value
func (o *CassandraSourceRegistrationPatchParamsCassandraCredentials) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *CassandraSourceRegistrationPatchParamsCassandraCredentials) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *CassandraSourceRegistrationPatchParamsCassandraCredentials) SetPassword(v string) {
	o.Password = v
}

// GetUsername returns the Username field value
func (o *CassandraSourceRegistrationPatchParamsCassandraCredentials) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *CassandraSourceRegistrationPatchParamsCassandraCredentials) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *CassandraSourceRegistrationPatchParamsCassandraCredentials) SetUsername(v string) {
	o.Username = v
}

func (o CassandraSourceRegistrationPatchParamsCassandraCredentials) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CassandraSourceRegistrationPatchParamsCassandraCredentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["password"] = o.Password
	toSerialize["username"] = o.Username
	return toSerialize, nil
}

func (o *CassandraSourceRegistrationPatchParamsCassandraCredentials) UnmarshalJSON(data []byte) (err error) {
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

	varCassandraSourceRegistrationPatchParamsCassandraCredentials := _CassandraSourceRegistrationPatchParamsCassandraCredentials{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCassandraSourceRegistrationPatchParamsCassandraCredentials)

	if err != nil {
		return err
	}

	*o = CassandraSourceRegistrationPatchParamsCassandraCredentials(varCassandraSourceRegistrationPatchParamsCassandraCredentials)

	return err
}

type NullableCassandraSourceRegistrationPatchParamsCassandraCredentials struct {
	value *CassandraSourceRegistrationPatchParamsCassandraCredentials
	isSet bool
}

func (v NullableCassandraSourceRegistrationPatchParamsCassandraCredentials) Get() *CassandraSourceRegistrationPatchParamsCassandraCredentials {
	return v.value
}

func (v *NullableCassandraSourceRegistrationPatchParamsCassandraCredentials) Set(val *CassandraSourceRegistrationPatchParamsCassandraCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableCassandraSourceRegistrationPatchParamsCassandraCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableCassandraSourceRegistrationPatchParamsCassandraCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCassandraSourceRegistrationPatchParamsCassandraCredentials(val *CassandraSourceRegistrationPatchParamsCassandraCredentials) *NullableCassandraSourceRegistrationPatchParamsCassandraCredentials {
	return &NullableCassandraSourceRegistrationPatchParamsCassandraCredentials{value: val, isSet: true}
}

func (v NullableCassandraSourceRegistrationPatchParamsCassandraCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCassandraSourceRegistrationPatchParamsCassandraCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


