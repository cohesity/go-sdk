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

// AwsIAmUserParams Specifies the parameters which are specific to IAmUSer Authentication Method for AWS External Target.
type AwsIAmUserParams struct {
	// Specifies the Access Key Id of the external target.
	AccessKeyId NullableString `json:"accessKeyId"`
	// Specifies the Secret Access Key of the external target.
	SecretAccessKey NullableString `json:"secretAccessKey"`
}

// NewAwsIAmUserParams instantiates a new AwsIAmUserParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsIAmUserParams(accessKeyId NullableString, secretAccessKey NullableString) *AwsIAmUserParams {
	this := AwsIAmUserParams{}
	this.AccessKeyId = accessKeyId
	this.SecretAccessKey = secretAccessKey
	return &this
}

// NewAwsIAmUserParamsWithDefaults instantiates a new AwsIAmUserParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsIAmUserParamsWithDefaults() *AwsIAmUserParams {
	this := AwsIAmUserParams{}
	return &this
}

// GetAccessKeyId returns the AccessKeyId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AwsIAmUserParams) GetAccessKeyId() string {
	if o == nil || o.AccessKeyId.Get() == nil {
		var ret string
		return ret
	}

	return *o.AccessKeyId.Get()
}

// GetAccessKeyIdOk returns a tuple with the AccessKeyId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsIAmUserParams) GetAccessKeyIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccessKeyId.Get(), o.AccessKeyId.IsSet()
}

// SetAccessKeyId sets field value
func (o *AwsIAmUserParams) SetAccessKeyId(v string) {
	o.AccessKeyId.Set(&v)
}

// GetSecretAccessKey returns the SecretAccessKey field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AwsIAmUserParams) GetSecretAccessKey() string {
	if o == nil || o.SecretAccessKey.Get() == nil {
		var ret string
		return ret
	}

	return *o.SecretAccessKey.Get()
}

// GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsIAmUserParams) GetSecretAccessKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SecretAccessKey.Get(), o.SecretAccessKey.IsSet()
}

// SetSecretAccessKey sets field value
func (o *AwsIAmUserParams) SetSecretAccessKey(v string) {
	o.SecretAccessKey.Set(&v)
}

func (o AwsIAmUserParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["accessKeyId"] = o.AccessKeyId.Get()
	}
	if true {
		toSerialize["secretAccessKey"] = o.SecretAccessKey.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAwsIAmUserParams struct {
	value *AwsIAmUserParams
	isSet bool
}

func (v NullableAwsIAmUserParams) Get() *AwsIAmUserParams {
	return v.value
}

func (v *NullableAwsIAmUserParams) Set(val *AwsIAmUserParams) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsIAmUserParams) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsIAmUserParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsIAmUserParams(val *AwsIAmUserParams) *NullableAwsIAmUserParams {
	return &NullableAwsIAmUserParams{value: val, isSet: true}
}

func (v NullableAwsIAmUserParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsIAmUserParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o AwsIAmUserParams) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}