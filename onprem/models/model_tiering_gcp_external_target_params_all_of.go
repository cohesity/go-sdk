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

// TieringGcpExternalTargetParamsAllOf struct for TieringGcpExternalTargetParamsAllOf
type TieringGcpExternalTargetParamsAllOf struct {
	// Specifies the GCP External Target class.
	StorageClass NullableString `json:"storageClass"`
}

// NewTieringGcpExternalTargetParamsAllOf instantiates a new TieringGcpExternalTargetParamsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTieringGcpExternalTargetParamsAllOf(storageClass NullableString) *TieringGcpExternalTargetParamsAllOf {
	this := TieringGcpExternalTargetParamsAllOf{}
	this.StorageClass = storageClass
	return &this
}

// NewTieringGcpExternalTargetParamsAllOfWithDefaults instantiates a new TieringGcpExternalTargetParamsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTieringGcpExternalTargetParamsAllOfWithDefaults() *TieringGcpExternalTargetParamsAllOf {
	this := TieringGcpExternalTargetParamsAllOf{}
	return &this
}

// GetStorageClass returns the StorageClass field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TieringGcpExternalTargetParamsAllOf) GetStorageClass() string {
	if o == nil || o.StorageClass.Get() == nil {
		var ret string
		return ret
	}

	return *o.StorageClass.Get()
}

// GetStorageClassOk returns a tuple with the StorageClass field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TieringGcpExternalTargetParamsAllOf) GetStorageClassOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StorageClass.Get(), o.StorageClass.IsSet()
}

// SetStorageClass sets field value
func (o *TieringGcpExternalTargetParamsAllOf) SetStorageClass(v string) {
	o.StorageClass.Set(&v)
}

func (o TieringGcpExternalTargetParamsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["storageClass"] = o.StorageClass.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableTieringGcpExternalTargetParamsAllOf struct {
	value *TieringGcpExternalTargetParamsAllOf
	isSet bool
}

func (v NullableTieringGcpExternalTargetParamsAllOf) Get() *TieringGcpExternalTargetParamsAllOf {
	return v.value
}

func (v *NullableTieringGcpExternalTargetParamsAllOf) Set(val *TieringGcpExternalTargetParamsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTieringGcpExternalTargetParamsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTieringGcpExternalTargetParamsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTieringGcpExternalTargetParamsAllOf(val *TieringGcpExternalTargetParamsAllOf) *NullableTieringGcpExternalTargetParamsAllOf {
	return &NullableTieringGcpExternalTargetParamsAllOf{value: val, isSet: true}
}

func (v NullableTieringGcpExternalTargetParamsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTieringGcpExternalTargetParamsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o TieringGcpExternalTargetParamsAllOf) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}