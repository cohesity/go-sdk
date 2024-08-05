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

// checks if the AcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials{}

// AcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials Specifies the credentials for the target VM.
type AcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials struct {
	// Specifies the password to access target entity.
	Password string `json:"password"`
	// Specifies the username to access target entity.
	Username string `json:"username"`
}

type _AcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials AcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials

// NewAcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials instantiates a new AcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials(password string, username string) *AcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials {
	this := AcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials{}
	this.Password = password
	this.Username = username
	return &this
}

// NewAcropolisRecoverFilesOriginalTargetConfigTargetVmCredentialsWithDefaults instantiates a new AcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcropolisRecoverFilesOriginalTargetConfigTargetVmCredentialsWithDefaults() *AcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials {
	this := AcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials{}
	return &this
}

// GetPassword returns the Password field value
func (o *AcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *AcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *AcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials) SetPassword(v string) {
	o.Password = v
}

// GetUsername returns the Username field value
func (o *AcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *AcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *AcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials) SetUsername(v string) {
	o.Username = v
}

func (o AcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["password"] = o.Password
	toSerialize["username"] = o.Username
	return toSerialize, nil
}

func (o *AcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials) UnmarshalJSON(data []byte) (err error) {
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

	varAcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials := _AcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials)

	if err != nil {
		return err
	}

	*o = AcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials(varAcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials)

	return err
}

type NullableAcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials struct {
	value *AcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials
	isSet bool
}

func (v NullableAcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials) Get() *AcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials {
	return v.value
}

func (v *NullableAcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials) Set(val *AcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableAcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableAcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials(val *AcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials) *NullableAcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials {
	return &NullableAcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials{value: val, isSet: true}
}

func (v NullableAcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


