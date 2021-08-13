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

// AwsObjectProtectionRequestParamsAllOf struct for AwsObjectProtectionRequestParamsAllOf
type AwsObjectProtectionRequestParamsAllOf struct {
	NativeProtectionTypeParams *AwsNativeObjectProtectionParams `json:"nativeProtectionTypeParams,omitempty"`
	SnapshotManagerProtectionTypeParams *AwsSnapshotManagerObjectProtectionParams `json:"snapshotManagerProtectionTypeParams,omitempty"`
}

// NewAwsObjectProtectionRequestParamsAllOf instantiates a new AwsObjectProtectionRequestParamsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsObjectProtectionRequestParamsAllOf() *AwsObjectProtectionRequestParamsAllOf {
	this := AwsObjectProtectionRequestParamsAllOf{}
	return &this
}

// NewAwsObjectProtectionRequestParamsAllOfWithDefaults instantiates a new AwsObjectProtectionRequestParamsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsObjectProtectionRequestParamsAllOfWithDefaults() *AwsObjectProtectionRequestParamsAllOf {
	this := AwsObjectProtectionRequestParamsAllOf{}
	return &this
}

// GetNativeProtectionTypeParams returns the NativeProtectionTypeParams field value if set, zero value otherwise.
func (o *AwsObjectProtectionRequestParamsAllOf) GetNativeProtectionTypeParams() AwsNativeObjectProtectionParams {
	if o == nil || o.NativeProtectionTypeParams == nil {
		var ret AwsNativeObjectProtectionParams
		return ret
	}
	return *o.NativeProtectionTypeParams
}

// GetNativeProtectionTypeParamsOk returns a tuple with the NativeProtectionTypeParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsObjectProtectionRequestParamsAllOf) GetNativeProtectionTypeParamsOk() (*AwsNativeObjectProtectionParams, bool) {
	if o == nil || o.NativeProtectionTypeParams == nil {
		return nil, false
	}
	return o.NativeProtectionTypeParams, true
}

// HasNativeProtectionTypeParams returns a boolean if a field has been set.
func (o *AwsObjectProtectionRequestParamsAllOf) HasNativeProtectionTypeParams() bool {
	if o != nil && o.NativeProtectionTypeParams != nil {
		return true
	}

	return false
}

// SetNativeProtectionTypeParams gets a reference to the given AwsNativeObjectProtectionParams and assigns it to the NativeProtectionTypeParams field.
func (o *AwsObjectProtectionRequestParamsAllOf) SetNativeProtectionTypeParams(v AwsNativeObjectProtectionParams) {
	o.NativeProtectionTypeParams = &v
}

// GetSnapshotManagerProtectionTypeParams returns the SnapshotManagerProtectionTypeParams field value if set, zero value otherwise.
func (o *AwsObjectProtectionRequestParamsAllOf) GetSnapshotManagerProtectionTypeParams() AwsSnapshotManagerObjectProtectionParams {
	if o == nil || o.SnapshotManagerProtectionTypeParams == nil {
		var ret AwsSnapshotManagerObjectProtectionParams
		return ret
	}
	return *o.SnapshotManagerProtectionTypeParams
}

// GetSnapshotManagerProtectionTypeParamsOk returns a tuple with the SnapshotManagerProtectionTypeParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsObjectProtectionRequestParamsAllOf) GetSnapshotManagerProtectionTypeParamsOk() (*AwsSnapshotManagerObjectProtectionParams, bool) {
	if o == nil || o.SnapshotManagerProtectionTypeParams == nil {
		return nil, false
	}
	return o.SnapshotManagerProtectionTypeParams, true
}

// HasSnapshotManagerProtectionTypeParams returns a boolean if a field has been set.
func (o *AwsObjectProtectionRequestParamsAllOf) HasSnapshotManagerProtectionTypeParams() bool {
	if o != nil && o.SnapshotManagerProtectionTypeParams != nil {
		return true
	}

	return false
}

// SetSnapshotManagerProtectionTypeParams gets a reference to the given AwsSnapshotManagerObjectProtectionParams and assigns it to the SnapshotManagerProtectionTypeParams field.
func (o *AwsObjectProtectionRequestParamsAllOf) SetSnapshotManagerProtectionTypeParams(v AwsSnapshotManagerObjectProtectionParams) {
	o.SnapshotManagerProtectionTypeParams = &v
}

func (o AwsObjectProtectionRequestParamsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NativeProtectionTypeParams != nil {
		toSerialize["nativeProtectionTypeParams"] = o.NativeProtectionTypeParams
	}
	if o.SnapshotManagerProtectionTypeParams != nil {
		toSerialize["snapshotManagerProtectionTypeParams"] = o.SnapshotManagerProtectionTypeParams
	}
	return json.Marshal(toSerialize)
}

type NullableAwsObjectProtectionRequestParamsAllOf struct {
	value *AwsObjectProtectionRequestParamsAllOf
	isSet bool
}

func (v NullableAwsObjectProtectionRequestParamsAllOf) Get() *AwsObjectProtectionRequestParamsAllOf {
	return v.value
}

func (v *NullableAwsObjectProtectionRequestParamsAllOf) Set(val *AwsObjectProtectionRequestParamsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsObjectProtectionRequestParamsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsObjectProtectionRequestParamsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsObjectProtectionRequestParamsAllOf(val *AwsObjectProtectionRequestParamsAllOf) *NullableAwsObjectProtectionRequestParamsAllOf {
	return &NullableAwsObjectProtectionRequestParamsAllOf{value: val, isSet: true}
}

func (v NullableAwsObjectProtectionRequestParamsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsObjectProtectionRequestParamsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o AwsObjectProtectionRequestParamsAllOf) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}