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

// CommonCreateOrUpdateRigelConnectionRequest Specify the common params to create or update a connection of Rigel.
type CommonCreateOrUpdateRigelConnectionRequest struct {
	// Specifies the id of the tenant which the connection belongs to.
	TenantId NullableString `json:"tenantId"`
	// Specifies the name of the connection.
	Name NullableString `json:"name"`
}

// NewCommonCreateOrUpdateRigelConnectionRequest instantiates a new CommonCreateOrUpdateRigelConnectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonCreateOrUpdateRigelConnectionRequest(tenantId NullableString, name NullableString) *CommonCreateOrUpdateRigelConnectionRequest {
	this := CommonCreateOrUpdateRigelConnectionRequest{}
	this.TenantId = tenantId
	this.Name = name
	return &this
}

// NewCommonCreateOrUpdateRigelConnectionRequestWithDefaults instantiates a new CommonCreateOrUpdateRigelConnectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonCreateOrUpdateRigelConnectionRequestWithDefaults() *CommonCreateOrUpdateRigelConnectionRequest {
	this := CommonCreateOrUpdateRigelConnectionRequest{}
	return &this
}

// GetTenantId returns the TenantId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CommonCreateOrUpdateRigelConnectionRequest) GetTenantId() string {
	if o == nil || o.TenantId.Get() == nil {
		var ret string
		return ret
	}

	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonCreateOrUpdateRigelConnectionRequest) GetTenantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// SetTenantId sets field value
func (o *CommonCreateOrUpdateRigelConnectionRequest) SetTenantId(v string) {
	o.TenantId.Set(&v)
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CommonCreateOrUpdateRigelConnectionRequest) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonCreateOrUpdateRigelConnectionRequest) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *CommonCreateOrUpdateRigelConnectionRequest) SetName(v string) {
	o.Name.Set(&v)
}

func (o CommonCreateOrUpdateRigelConnectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	if true {
		toSerialize["name"] = o.Name.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCommonCreateOrUpdateRigelConnectionRequest struct {
	value *CommonCreateOrUpdateRigelConnectionRequest
	isSet bool
}

func (v NullableCommonCreateOrUpdateRigelConnectionRequest) Get() *CommonCreateOrUpdateRigelConnectionRequest {
	return v.value
}

func (v *NullableCommonCreateOrUpdateRigelConnectionRequest) Set(val *CommonCreateOrUpdateRigelConnectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonCreateOrUpdateRigelConnectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonCreateOrUpdateRigelConnectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonCreateOrUpdateRigelConnectionRequest(val *CommonCreateOrUpdateRigelConnectionRequest) *NullableCommonCreateOrUpdateRigelConnectionRequest {
	return &NullableCommonCreateOrUpdateRigelConnectionRequest{value: val, isSet: true}
}

func (v NullableCommonCreateOrUpdateRigelConnectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonCreateOrUpdateRigelConnectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o CommonCreateOrUpdateRigelConnectionRequest) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}