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

// MssqlObjectProtectionResponseParams Specifies the response parameters specific to MSSQL object protection.
type MssqlObjectProtectionResponseParams struct {
	// Specifies the MSSQL Object Protection type.
	ObjectProtectionType string `json:"objectProtectionType"`
	CommonFileObjectProtectionTypeParams *CommonMssqlFileObjectProtectionParams `json:"commonFileObjectProtectionTypeParams,omitempty"`
	CommonNativeObjectProtectionTypeParams *CommonMssqlNativeObjectProtectionParams `json:"commonNativeObjectProtectionTypeParams,omitempty"`
}

// NewMssqlObjectProtectionResponseParams instantiates a new MssqlObjectProtectionResponseParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMssqlObjectProtectionResponseParams(objectProtectionType string) *MssqlObjectProtectionResponseParams {
	this := MssqlObjectProtectionResponseParams{}
	this.ObjectProtectionType = objectProtectionType
	return &this
}

// NewMssqlObjectProtectionResponseParamsWithDefaults instantiates a new MssqlObjectProtectionResponseParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMssqlObjectProtectionResponseParamsWithDefaults() *MssqlObjectProtectionResponseParams {
	this := MssqlObjectProtectionResponseParams{}
	return &this
}

// GetObjectProtectionType returns the ObjectProtectionType field value
func (o *MssqlObjectProtectionResponseParams) GetObjectProtectionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectProtectionType
}

// GetObjectProtectionTypeOk returns a tuple with the ObjectProtectionType field value
// and a boolean to check if the value has been set.
func (o *MssqlObjectProtectionResponseParams) GetObjectProtectionTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectProtectionType, true
}

// SetObjectProtectionType sets field value
func (o *MssqlObjectProtectionResponseParams) SetObjectProtectionType(v string) {
	o.ObjectProtectionType = v
}

// GetCommonFileObjectProtectionTypeParams returns the CommonFileObjectProtectionTypeParams field value if set, zero value otherwise.
func (o *MssqlObjectProtectionResponseParams) GetCommonFileObjectProtectionTypeParams() CommonMssqlFileObjectProtectionParams {
	if o == nil || o.CommonFileObjectProtectionTypeParams == nil {
		var ret CommonMssqlFileObjectProtectionParams
		return ret
	}
	return *o.CommonFileObjectProtectionTypeParams
}

// GetCommonFileObjectProtectionTypeParamsOk returns a tuple with the CommonFileObjectProtectionTypeParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MssqlObjectProtectionResponseParams) GetCommonFileObjectProtectionTypeParamsOk() (*CommonMssqlFileObjectProtectionParams, bool) {
	if o == nil || o.CommonFileObjectProtectionTypeParams == nil {
		return nil, false
	}
	return o.CommonFileObjectProtectionTypeParams, true
}

// HasCommonFileObjectProtectionTypeParams returns a boolean if a field has been set.
func (o *MssqlObjectProtectionResponseParams) HasCommonFileObjectProtectionTypeParams() bool {
	if o != nil && o.CommonFileObjectProtectionTypeParams != nil {
		return true
	}

	return false
}

// SetCommonFileObjectProtectionTypeParams gets a reference to the given CommonMssqlFileObjectProtectionParams and assigns it to the CommonFileObjectProtectionTypeParams field.
func (o *MssqlObjectProtectionResponseParams) SetCommonFileObjectProtectionTypeParams(v CommonMssqlFileObjectProtectionParams) {
	o.CommonFileObjectProtectionTypeParams = &v
}

// GetCommonNativeObjectProtectionTypeParams returns the CommonNativeObjectProtectionTypeParams field value if set, zero value otherwise.
func (o *MssqlObjectProtectionResponseParams) GetCommonNativeObjectProtectionTypeParams() CommonMssqlNativeObjectProtectionParams {
	if o == nil || o.CommonNativeObjectProtectionTypeParams == nil {
		var ret CommonMssqlNativeObjectProtectionParams
		return ret
	}
	return *o.CommonNativeObjectProtectionTypeParams
}

// GetCommonNativeObjectProtectionTypeParamsOk returns a tuple with the CommonNativeObjectProtectionTypeParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MssqlObjectProtectionResponseParams) GetCommonNativeObjectProtectionTypeParamsOk() (*CommonMssqlNativeObjectProtectionParams, bool) {
	if o == nil || o.CommonNativeObjectProtectionTypeParams == nil {
		return nil, false
	}
	return o.CommonNativeObjectProtectionTypeParams, true
}

// HasCommonNativeObjectProtectionTypeParams returns a boolean if a field has been set.
func (o *MssqlObjectProtectionResponseParams) HasCommonNativeObjectProtectionTypeParams() bool {
	if o != nil && o.CommonNativeObjectProtectionTypeParams != nil {
		return true
	}

	return false
}

// SetCommonNativeObjectProtectionTypeParams gets a reference to the given CommonMssqlNativeObjectProtectionParams and assigns it to the CommonNativeObjectProtectionTypeParams field.
func (o *MssqlObjectProtectionResponseParams) SetCommonNativeObjectProtectionTypeParams(v CommonMssqlNativeObjectProtectionParams) {
	o.CommonNativeObjectProtectionTypeParams = &v
}

func (o MssqlObjectProtectionResponseParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["objectProtectionType"] = o.ObjectProtectionType
	}
	if o.CommonFileObjectProtectionTypeParams != nil {
		toSerialize["commonFileObjectProtectionTypeParams"] = o.CommonFileObjectProtectionTypeParams
	}
	if o.CommonNativeObjectProtectionTypeParams != nil {
		toSerialize["commonNativeObjectProtectionTypeParams"] = o.CommonNativeObjectProtectionTypeParams
	}
	return json.Marshal(toSerialize)
}

type NullableMssqlObjectProtectionResponseParams struct {
	value *MssqlObjectProtectionResponseParams
	isSet bool
}

func (v NullableMssqlObjectProtectionResponseParams) Get() *MssqlObjectProtectionResponseParams {
	return v.value
}

func (v *NullableMssqlObjectProtectionResponseParams) Set(val *MssqlObjectProtectionResponseParams) {
	v.value = val
	v.isSet = true
}

func (v NullableMssqlObjectProtectionResponseParams) IsSet() bool {
	return v.isSet
}

func (v *NullableMssqlObjectProtectionResponseParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMssqlObjectProtectionResponseParams(val *MssqlObjectProtectionResponseParams) *NullableMssqlObjectProtectionResponseParams {
	return &NullableMssqlObjectProtectionResponseParams{value: val, isSet: true}
}

func (v NullableMssqlObjectProtectionResponseParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMssqlObjectProtectionResponseParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o MssqlObjectProtectionResponseParams) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}