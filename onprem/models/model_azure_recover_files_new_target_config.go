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

// AzureRecoverFilesNewTargetConfig Specifies the configuration for recovering files and folders to the new target.
type AzureRecoverFilesNewTargetConfig struct {
	// Specifies the target VM to recover files and folders to.
	TargetVm NullableRecoverTarget `json:"targetVm"`
	// Specifies the credentials for the target VM.
	TargetVmCredentials NullableCredentials `json:"targetVmCredentials"`
	// Specifies the absolute path location to recover files to.
	AbsolutePath NullableString `json:"absolutePath"`
}

// NewAzureRecoverFilesNewTargetConfig instantiates a new AzureRecoverFilesNewTargetConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureRecoverFilesNewTargetConfig(targetVm NullableRecoverTarget, targetVmCredentials NullableCredentials, absolutePath NullableString) *AzureRecoverFilesNewTargetConfig {
	this := AzureRecoverFilesNewTargetConfig{}
	this.TargetVm = targetVm
	this.TargetVmCredentials = targetVmCredentials
	this.AbsolutePath = absolutePath
	return &this
}

// NewAzureRecoverFilesNewTargetConfigWithDefaults instantiates a new AzureRecoverFilesNewTargetConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureRecoverFilesNewTargetConfigWithDefaults() *AzureRecoverFilesNewTargetConfig {
	this := AzureRecoverFilesNewTargetConfig{}
	return &this
}

// GetTargetVm returns the TargetVm field value
// If the value is explicit nil, the zero value for RecoverTarget will be returned
func (o *AzureRecoverFilesNewTargetConfig) GetTargetVm() RecoverTarget {
	if o == nil || o.TargetVm.Get() == nil {
		var ret RecoverTarget
		return ret
	}

	return *o.TargetVm.Get()
}

// GetTargetVmOk returns a tuple with the TargetVm field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureRecoverFilesNewTargetConfig) GetTargetVmOk() (*RecoverTarget, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TargetVm.Get(), o.TargetVm.IsSet()
}

// SetTargetVm sets field value
func (o *AzureRecoverFilesNewTargetConfig) SetTargetVm(v RecoverTarget) {
	o.TargetVm.Set(&v)
}

// GetTargetVmCredentials returns the TargetVmCredentials field value
// If the value is explicit nil, the zero value for Credentials will be returned
func (o *AzureRecoverFilesNewTargetConfig) GetTargetVmCredentials() Credentials {
	if o == nil || o.TargetVmCredentials.Get() == nil {
		var ret Credentials
		return ret
	}

	return *o.TargetVmCredentials.Get()
}

// GetTargetVmCredentialsOk returns a tuple with the TargetVmCredentials field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureRecoverFilesNewTargetConfig) GetTargetVmCredentialsOk() (*Credentials, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TargetVmCredentials.Get(), o.TargetVmCredentials.IsSet()
}

// SetTargetVmCredentials sets field value
func (o *AzureRecoverFilesNewTargetConfig) SetTargetVmCredentials(v Credentials) {
	o.TargetVmCredentials.Set(&v)
}

// GetAbsolutePath returns the AbsolutePath field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AzureRecoverFilesNewTargetConfig) GetAbsolutePath() string {
	if o == nil || o.AbsolutePath.Get() == nil {
		var ret string
		return ret
	}

	return *o.AbsolutePath.Get()
}

// GetAbsolutePathOk returns a tuple with the AbsolutePath field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureRecoverFilesNewTargetConfig) GetAbsolutePathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AbsolutePath.Get(), o.AbsolutePath.IsSet()
}

// SetAbsolutePath sets field value
func (o *AzureRecoverFilesNewTargetConfig) SetAbsolutePath(v string) {
	o.AbsolutePath.Set(&v)
}

func (o AzureRecoverFilesNewTargetConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["targetVm"] = o.TargetVm.Get()
	}
	if true {
		toSerialize["targetVmCredentials"] = o.TargetVmCredentials.Get()
	}
	if true {
		toSerialize["absolutePath"] = o.AbsolutePath.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAzureRecoverFilesNewTargetConfig struct {
	value *AzureRecoverFilesNewTargetConfig
	isSet bool
}

func (v NullableAzureRecoverFilesNewTargetConfig) Get() *AzureRecoverFilesNewTargetConfig {
	return v.value
}

func (v *NullableAzureRecoverFilesNewTargetConfig) Set(val *AzureRecoverFilesNewTargetConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureRecoverFilesNewTargetConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureRecoverFilesNewTargetConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureRecoverFilesNewTargetConfig(val *AzureRecoverFilesNewTargetConfig) *NullableAzureRecoverFilesNewTargetConfig {
	return &NullableAzureRecoverFilesNewTargetConfig{value: val, isSet: true}
}

func (v NullableAzureRecoverFilesNewTargetConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureRecoverFilesNewTargetConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o AzureRecoverFilesNewTargetConfig) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}