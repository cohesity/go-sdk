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

// CreateCadStorageDomainParam Sepcifies the parameters to create a CAD Storage Domain.
type CreateCadStorageDomainParam struct {
	// Specifies the external target id mapped to the CAD Storage Domain.
	ExternalTargetId NullableInt64 `json:"externalTargetId"`
}

// NewCreateCadStorageDomainParam instantiates a new CreateCadStorageDomainParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCadStorageDomainParam(externalTargetId NullableInt64) *CreateCadStorageDomainParam {
	this := CreateCadStorageDomainParam{}
	this.ExternalTargetId = externalTargetId
	return &this
}

// NewCreateCadStorageDomainParamWithDefaults instantiates a new CreateCadStorageDomainParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCadStorageDomainParamWithDefaults() *CreateCadStorageDomainParam {
	this := CreateCadStorageDomainParam{}
	return &this
}

// GetExternalTargetId returns the ExternalTargetId field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *CreateCadStorageDomainParam) GetExternalTargetId() int64 {
	if o == nil || o.ExternalTargetId.Get() == nil {
		var ret int64
		return ret
	}

	return *o.ExternalTargetId.Get()
}

// GetExternalTargetIdOk returns a tuple with the ExternalTargetId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCadStorageDomainParam) GetExternalTargetIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExternalTargetId.Get(), o.ExternalTargetId.IsSet()
}

// SetExternalTargetId sets field value
func (o *CreateCadStorageDomainParam) SetExternalTargetId(v int64) {
	o.ExternalTargetId.Set(&v)
}

func (o CreateCadStorageDomainParam) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["externalTargetId"] = o.ExternalTargetId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCreateCadStorageDomainParam struct {
	value *CreateCadStorageDomainParam
	isSet bool
}

func (v NullableCreateCadStorageDomainParam) Get() *CreateCadStorageDomainParam {
	return v.value
}

func (v *NullableCreateCadStorageDomainParam) Set(val *CreateCadStorageDomainParam) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCadStorageDomainParam) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCadStorageDomainParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCadStorageDomainParam(val *CreateCadStorageDomainParam) *NullableCreateCadStorageDomainParam {
	return &NullableCreateCadStorageDomainParam{value: val, isSet: true}
}

func (v NullableCreateCadStorageDomainParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCadStorageDomainParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o CreateCadStorageDomainParam) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}