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

// checks if the VmwareRecoverFilesNewTargetConfigTargetVmCredentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VmwareRecoverFilesNewTargetConfigTargetVmCredentials{}

// VmwareRecoverFilesNewTargetConfigTargetVmCredentials Specifies the credentials for the target VM. This is mandatory if the recoverMethod is AutoDeploy or UseHypervisorApis.
type VmwareRecoverFilesNewTargetConfigTargetVmCredentials struct {
	// Specifies the password to access target entity.
	Password string `json:"password"`
	// Specifies the username to access target entity.
	Username string `json:"username"`
}

type _VmwareRecoverFilesNewTargetConfigTargetVmCredentials VmwareRecoverFilesNewTargetConfigTargetVmCredentials

// NewVmwareRecoverFilesNewTargetConfigTargetVmCredentials instantiates a new VmwareRecoverFilesNewTargetConfigTargetVmCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVmwareRecoverFilesNewTargetConfigTargetVmCredentials(password string, username string) *VmwareRecoverFilesNewTargetConfigTargetVmCredentials {
	this := VmwareRecoverFilesNewTargetConfigTargetVmCredentials{}
	this.Password = password
	this.Username = username
	return &this
}

// NewVmwareRecoverFilesNewTargetConfigTargetVmCredentialsWithDefaults instantiates a new VmwareRecoverFilesNewTargetConfigTargetVmCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmwareRecoverFilesNewTargetConfigTargetVmCredentialsWithDefaults() *VmwareRecoverFilesNewTargetConfigTargetVmCredentials {
	this := VmwareRecoverFilesNewTargetConfigTargetVmCredentials{}
	return &this
}

// GetPassword returns the Password field value
func (o *VmwareRecoverFilesNewTargetConfigTargetVmCredentials) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *VmwareRecoverFilesNewTargetConfigTargetVmCredentials) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *VmwareRecoverFilesNewTargetConfigTargetVmCredentials) SetPassword(v string) {
	o.Password = v
}

// GetUsername returns the Username field value
func (o *VmwareRecoverFilesNewTargetConfigTargetVmCredentials) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *VmwareRecoverFilesNewTargetConfigTargetVmCredentials) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *VmwareRecoverFilesNewTargetConfigTargetVmCredentials) SetUsername(v string) {
	o.Username = v
}

func (o VmwareRecoverFilesNewTargetConfigTargetVmCredentials) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VmwareRecoverFilesNewTargetConfigTargetVmCredentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["password"] = o.Password
	toSerialize["username"] = o.Username
	return toSerialize, nil
}

func (o *VmwareRecoverFilesNewTargetConfigTargetVmCredentials) UnmarshalJSON(data []byte) (err error) {
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

	varVmwareRecoverFilesNewTargetConfigTargetVmCredentials := _VmwareRecoverFilesNewTargetConfigTargetVmCredentials{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVmwareRecoverFilesNewTargetConfigTargetVmCredentials)

	if err != nil {
		return err
	}

	*o = VmwareRecoverFilesNewTargetConfigTargetVmCredentials(varVmwareRecoverFilesNewTargetConfigTargetVmCredentials)

	return err
}

type NullableVmwareRecoverFilesNewTargetConfigTargetVmCredentials struct {
	value *VmwareRecoverFilesNewTargetConfigTargetVmCredentials
	isSet bool
}

func (v NullableVmwareRecoverFilesNewTargetConfigTargetVmCredentials) Get() *VmwareRecoverFilesNewTargetConfigTargetVmCredentials {
	return v.value
}

func (v *NullableVmwareRecoverFilesNewTargetConfigTargetVmCredentials) Set(val *VmwareRecoverFilesNewTargetConfigTargetVmCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableVmwareRecoverFilesNewTargetConfigTargetVmCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableVmwareRecoverFilesNewTargetConfigTargetVmCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVmwareRecoverFilesNewTargetConfigTargetVmCredentials(val *VmwareRecoverFilesNewTargetConfigTargetVmCredentials) *NullableVmwareRecoverFilesNewTargetConfigTargetVmCredentials {
	return &NullableVmwareRecoverFilesNewTargetConfigTargetVmCredentials{value: val, isSet: true}
}

func (v NullableVmwareRecoverFilesNewTargetConfigTargetVmCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVmwareRecoverFilesNewTargetConfigTargetVmCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


