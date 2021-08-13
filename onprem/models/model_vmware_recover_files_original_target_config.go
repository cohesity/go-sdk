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

// VmwareRecoverFilesOriginalTargetConfig Specifies the configuration for recovering files and folders to the original target.
type VmwareRecoverFilesOriginalTargetConfig struct {
	// Specifies the method to recover files and folders.
	RecoverMethod string `json:"recoverMethod"`
	// Specifies the credentials for the target VM. This is mandatory if the recoverMethod is AutoDeploy or UseHypervisorApis.
	TargetVmCredentials NullableCredentials `json:"targetVmCredentials,omitempty"`
	// Specifies whether to recover files and folders to the original path location. If false, alternatePath must be specified.
	RecoverToOriginalPath NullableBool `json:"recoverToOriginalPath"`
	// Specifies the alternate path location to recover files to.
	AlternatePath NullableString `json:"alternatePath,omitempty"`
}

// NewVmwareRecoverFilesOriginalTargetConfig instantiates a new VmwareRecoverFilesOriginalTargetConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVmwareRecoverFilesOriginalTargetConfig(recoverMethod string, recoverToOriginalPath NullableBool) *VmwareRecoverFilesOriginalTargetConfig {
	this := VmwareRecoverFilesOriginalTargetConfig{}
	this.RecoverMethod = recoverMethod
	this.RecoverToOriginalPath = recoverToOriginalPath
	return &this
}

// NewVmwareRecoverFilesOriginalTargetConfigWithDefaults instantiates a new VmwareRecoverFilesOriginalTargetConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmwareRecoverFilesOriginalTargetConfigWithDefaults() *VmwareRecoverFilesOriginalTargetConfig {
	this := VmwareRecoverFilesOriginalTargetConfig{}
	return &this
}

// GetRecoverMethod returns the RecoverMethod field value
func (o *VmwareRecoverFilesOriginalTargetConfig) GetRecoverMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecoverMethod
}

// GetRecoverMethodOk returns a tuple with the RecoverMethod field value
// and a boolean to check if the value has been set.
func (o *VmwareRecoverFilesOriginalTargetConfig) GetRecoverMethodOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RecoverMethod, true
}

// SetRecoverMethod sets field value
func (o *VmwareRecoverFilesOriginalTargetConfig) SetRecoverMethod(v string) {
	o.RecoverMethod = v
}

// GetTargetVmCredentials returns the TargetVmCredentials field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VmwareRecoverFilesOriginalTargetConfig) GetTargetVmCredentials() Credentials {
	if o == nil || o.TargetVmCredentials.Get() == nil {
		var ret Credentials
		return ret
	}
	return *o.TargetVmCredentials.Get()
}

// GetTargetVmCredentialsOk returns a tuple with the TargetVmCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmwareRecoverFilesOriginalTargetConfig) GetTargetVmCredentialsOk() (*Credentials, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TargetVmCredentials.Get(), o.TargetVmCredentials.IsSet()
}

// HasTargetVmCredentials returns a boolean if a field has been set.
func (o *VmwareRecoverFilesOriginalTargetConfig) HasTargetVmCredentials() bool {
	if o != nil && o.TargetVmCredentials.IsSet() {
		return true
	}

	return false
}

// SetTargetVmCredentials gets a reference to the given NullableCredentials and assigns it to the TargetVmCredentials field.
func (o *VmwareRecoverFilesOriginalTargetConfig) SetTargetVmCredentials(v Credentials) {
	o.TargetVmCredentials.Set(&v)
}
// SetTargetVmCredentialsNil sets the value for TargetVmCredentials to be an explicit nil
func (o *VmwareRecoverFilesOriginalTargetConfig) SetTargetVmCredentialsNil() {
	o.TargetVmCredentials.Set(nil)
}

// UnsetTargetVmCredentials ensures that no value is present for TargetVmCredentials, not even an explicit nil
func (o *VmwareRecoverFilesOriginalTargetConfig) UnsetTargetVmCredentials() {
	o.TargetVmCredentials.Unset()
}

// GetRecoverToOriginalPath returns the RecoverToOriginalPath field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *VmwareRecoverFilesOriginalTargetConfig) GetRecoverToOriginalPath() bool {
	if o == nil || o.RecoverToOriginalPath.Get() == nil {
		var ret bool
		return ret
	}

	return *o.RecoverToOriginalPath.Get()
}

// GetRecoverToOriginalPathOk returns a tuple with the RecoverToOriginalPath field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmwareRecoverFilesOriginalTargetConfig) GetRecoverToOriginalPathOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RecoverToOriginalPath.Get(), o.RecoverToOriginalPath.IsSet()
}

// SetRecoverToOriginalPath sets field value
func (o *VmwareRecoverFilesOriginalTargetConfig) SetRecoverToOriginalPath(v bool) {
	o.RecoverToOriginalPath.Set(&v)
}

// GetAlternatePath returns the AlternatePath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VmwareRecoverFilesOriginalTargetConfig) GetAlternatePath() string {
	if o == nil || o.AlternatePath.Get() == nil {
		var ret string
		return ret
	}
	return *o.AlternatePath.Get()
}

// GetAlternatePathOk returns a tuple with the AlternatePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmwareRecoverFilesOriginalTargetConfig) GetAlternatePathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AlternatePath.Get(), o.AlternatePath.IsSet()
}

// HasAlternatePath returns a boolean if a field has been set.
func (o *VmwareRecoverFilesOriginalTargetConfig) HasAlternatePath() bool {
	if o != nil && o.AlternatePath.IsSet() {
		return true
	}

	return false
}

// SetAlternatePath gets a reference to the given NullableString and assigns it to the AlternatePath field.
func (o *VmwareRecoverFilesOriginalTargetConfig) SetAlternatePath(v string) {
	o.AlternatePath.Set(&v)
}
// SetAlternatePathNil sets the value for AlternatePath to be an explicit nil
func (o *VmwareRecoverFilesOriginalTargetConfig) SetAlternatePathNil() {
	o.AlternatePath.Set(nil)
}

// UnsetAlternatePath ensures that no value is present for AlternatePath, not even an explicit nil
func (o *VmwareRecoverFilesOriginalTargetConfig) UnsetAlternatePath() {
	o.AlternatePath.Unset()
}

func (o VmwareRecoverFilesOriginalTargetConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["recoverMethod"] = o.RecoverMethod
	}
	if o.TargetVmCredentials.IsSet() {
		toSerialize["targetVmCredentials"] = o.TargetVmCredentials.Get()
	}
	if true {
		toSerialize["recoverToOriginalPath"] = o.RecoverToOriginalPath.Get()
	}
	if o.AlternatePath.IsSet() {
		toSerialize["alternatePath"] = o.AlternatePath.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableVmwareRecoverFilesOriginalTargetConfig struct {
	value *VmwareRecoverFilesOriginalTargetConfig
	isSet bool
}

func (v NullableVmwareRecoverFilesOriginalTargetConfig) Get() *VmwareRecoverFilesOriginalTargetConfig {
	return v.value
}

func (v *NullableVmwareRecoverFilesOriginalTargetConfig) Set(val *VmwareRecoverFilesOriginalTargetConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableVmwareRecoverFilesOriginalTargetConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableVmwareRecoverFilesOriginalTargetConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVmwareRecoverFilesOriginalTargetConfig(val *VmwareRecoverFilesOriginalTargetConfig) *NullableVmwareRecoverFilesOriginalTargetConfig {
	return &NullableVmwareRecoverFilesOriginalTargetConfig{value: val, isSet: true}
}

func (v NullableVmwareRecoverFilesOriginalTargetConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVmwareRecoverFilesOriginalTargetConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o VmwareRecoverFilesOriginalTargetConfig) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}