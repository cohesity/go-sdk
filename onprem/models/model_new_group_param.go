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

// NewGroupParam Specifies the parameters for using a new protection group.
type NewGroupParam struct {
	// Specifies the name of the new protection group.
	Name NullableString `json:"name"`
	// Specifies the policy id of the new protection group.
	PolicyId NullableString `json:"policyId"`
	// Specifies the storage domain id of the new protection group.
	StorageDomainId NullableInt64 `json:"storageDomainId"`
}

// NewNewGroupParam instantiates a new NewGroupParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewGroupParam(name NullableString, policyId NullableString, storageDomainId NullableInt64) *NewGroupParam {
	this := NewGroupParam{}
	this.Name = name
	this.PolicyId = policyId
	this.StorageDomainId = storageDomainId
	return &this
}

// NewNewGroupParamWithDefaults instantiates a new NewGroupParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewGroupParamWithDefaults() *NewGroupParam {
	this := NewGroupParam{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *NewGroupParam) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NewGroupParam) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *NewGroupParam) SetName(v string) {
	o.Name.Set(&v)
}

// GetPolicyId returns the PolicyId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *NewGroupParam) GetPolicyId() string {
	if o == nil || o.PolicyId.Get() == nil {
		var ret string
		return ret
	}

	return *o.PolicyId.Get()
}

// GetPolicyIdOk returns a tuple with the PolicyId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NewGroupParam) GetPolicyIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PolicyId.Get(), o.PolicyId.IsSet()
}

// SetPolicyId sets field value
func (o *NewGroupParam) SetPolicyId(v string) {
	o.PolicyId.Set(&v)
}

// GetStorageDomainId returns the StorageDomainId field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *NewGroupParam) GetStorageDomainId() int64 {
	if o == nil || o.StorageDomainId.Get() == nil {
		var ret int64
		return ret
	}

	return *o.StorageDomainId.Get()
}

// GetStorageDomainIdOk returns a tuple with the StorageDomainId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NewGroupParam) GetStorageDomainIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StorageDomainId.Get(), o.StorageDomainId.IsSet()
}

// SetStorageDomainId sets field value
func (o *NewGroupParam) SetStorageDomainId(v int64) {
	o.StorageDomainId.Set(&v)
}

func (o NewGroupParam) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name.Get()
	}
	if true {
		toSerialize["policyId"] = o.PolicyId.Get()
	}
	if true {
		toSerialize["storageDomainId"] = o.StorageDomainId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableNewGroupParam struct {
	value *NewGroupParam
	isSet bool
}

func (v NullableNewGroupParam) Get() *NewGroupParam {
	return v.value
}

func (v *NullableNewGroupParam) Set(val *NewGroupParam) {
	v.value = val
	v.isSet = true
}

func (v NullableNewGroupParam) IsSet() bool {
	return v.isSet
}

func (v *NullableNewGroupParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewGroupParam(val *NewGroupParam) *NullableNewGroupParam {
	return &NullableNewGroupParam{value: val, isSet: true}
}

func (v NullableNewGroupParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewGroupParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o NewGroupParam) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}