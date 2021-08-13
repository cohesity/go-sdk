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

// CommonTieringAwsExternalTargetParamsAllOf struct for CommonTieringAwsExternalTargetParamsAllOf
type CommonTieringAwsExternalTargetParamsAllOf struct {
	// Specifies the AWS External Target storage class.
	StorageClass NullableString `json:"storageClass"`
}

// NewCommonTieringAwsExternalTargetParamsAllOf instantiates a new CommonTieringAwsExternalTargetParamsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonTieringAwsExternalTargetParamsAllOf(storageClass NullableString) *CommonTieringAwsExternalTargetParamsAllOf {
	this := CommonTieringAwsExternalTargetParamsAllOf{}
	this.StorageClass = storageClass
	return &this
}

// NewCommonTieringAwsExternalTargetParamsAllOfWithDefaults instantiates a new CommonTieringAwsExternalTargetParamsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonTieringAwsExternalTargetParamsAllOfWithDefaults() *CommonTieringAwsExternalTargetParamsAllOf {
	this := CommonTieringAwsExternalTargetParamsAllOf{}
	return &this
}

// GetStorageClass returns the StorageClass field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CommonTieringAwsExternalTargetParamsAllOf) GetStorageClass() string {
	if o == nil || o.StorageClass.Get() == nil {
		var ret string
		return ret
	}

	return *o.StorageClass.Get()
}

// GetStorageClassOk returns a tuple with the StorageClass field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonTieringAwsExternalTargetParamsAllOf) GetStorageClassOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StorageClass.Get(), o.StorageClass.IsSet()
}

// SetStorageClass sets field value
func (o *CommonTieringAwsExternalTargetParamsAllOf) SetStorageClass(v string) {
	o.StorageClass.Set(&v)
}

func (o CommonTieringAwsExternalTargetParamsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["storageClass"] = o.StorageClass.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCommonTieringAwsExternalTargetParamsAllOf struct {
	value *CommonTieringAwsExternalTargetParamsAllOf
	isSet bool
}

func (v NullableCommonTieringAwsExternalTargetParamsAllOf) Get() *CommonTieringAwsExternalTargetParamsAllOf {
	return v.value
}

func (v *NullableCommonTieringAwsExternalTargetParamsAllOf) Set(val *CommonTieringAwsExternalTargetParamsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonTieringAwsExternalTargetParamsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonTieringAwsExternalTargetParamsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonTieringAwsExternalTargetParamsAllOf(val *CommonTieringAwsExternalTargetParamsAllOf) *NullableCommonTieringAwsExternalTargetParamsAllOf {
	return &NullableCommonTieringAwsExternalTargetParamsAllOf{value: val, isSet: true}
}

func (v NullableCommonTieringAwsExternalTargetParamsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonTieringAwsExternalTargetParamsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


