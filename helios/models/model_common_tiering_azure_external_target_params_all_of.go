/*
 * Cohesity REST API
 *
 * Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

// package helios
package models

import (
	"encoding/json"
	. "github.com/cohesity/go-sdk/helios/utils"
)

var _ = NullableBool{}

// CommonTieringAzureExternalTargetParamsAllOf struct for CommonTieringAzureExternalTargetParamsAllOf
type CommonTieringAzureExternalTargetParamsAllOf struct {
	// Specifies the Azure External Target class.
	StorageClass NullableString `json:"storageClass"`
}

// NewCommonTieringAzureExternalTargetParamsAllOf instantiates a new CommonTieringAzureExternalTargetParamsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonTieringAzureExternalTargetParamsAllOf(storageClass NullableString) *CommonTieringAzureExternalTargetParamsAllOf {
	this := CommonTieringAzureExternalTargetParamsAllOf{}
	this.StorageClass = storageClass
	return &this
}

// NewCommonTieringAzureExternalTargetParamsAllOfWithDefaults instantiates a new CommonTieringAzureExternalTargetParamsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonTieringAzureExternalTargetParamsAllOfWithDefaults() *CommonTieringAzureExternalTargetParamsAllOf {
	this := CommonTieringAzureExternalTargetParamsAllOf{}
	return &this
}

// GetStorageClass returns the StorageClass field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CommonTieringAzureExternalTargetParamsAllOf) GetStorageClass() string {
	if o == nil || o.StorageClass.Get() == nil {
		var ret string
		return ret
	}

	return *o.StorageClass.Get()
}

// GetStorageClassOk returns a tuple with the StorageClass field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonTieringAzureExternalTargetParamsAllOf) GetStorageClassOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StorageClass.Get(), o.StorageClass.IsSet()
}

// SetStorageClass sets field value
func (o *CommonTieringAzureExternalTargetParamsAllOf) SetStorageClass(v string) {
	o.StorageClass.Set(&v)
}

func (o CommonTieringAzureExternalTargetParamsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["storageClass"] = o.StorageClass.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCommonTieringAzureExternalTargetParamsAllOf struct {
	value *CommonTieringAzureExternalTargetParamsAllOf
	isSet bool
}

func (v NullableCommonTieringAzureExternalTargetParamsAllOf) Get() *CommonTieringAzureExternalTargetParamsAllOf {
	return v.value
}

func (v *NullableCommonTieringAzureExternalTargetParamsAllOf) Set(val *CommonTieringAzureExternalTargetParamsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonTieringAzureExternalTargetParamsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonTieringAzureExternalTargetParamsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonTieringAzureExternalTargetParamsAllOf(val *CommonTieringAzureExternalTargetParamsAllOf) *NullableCommonTieringAzureExternalTargetParamsAllOf {
	return &NullableCommonTieringAzureExternalTargetParamsAllOf{value: val, isSet: true}
}

func (v NullableCommonTieringAzureExternalTargetParamsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonTieringAzureExternalTargetParamsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


