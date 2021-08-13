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

// GetProtectedObjectResponse Specifies the protected objects response.
type GetProtectedObjectResponse struct {
	Object *ProtectedObjectInfo `json:"object,omitempty"`
}

// NewGetProtectedObjectResponse instantiates a new GetProtectedObjectResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProtectedObjectResponse() *GetProtectedObjectResponse {
	this := GetProtectedObjectResponse{}
	return &this
}

// NewGetProtectedObjectResponseWithDefaults instantiates a new GetProtectedObjectResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProtectedObjectResponseWithDefaults() *GetProtectedObjectResponse {
	this := GetProtectedObjectResponse{}
	return &this
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *GetProtectedObjectResponse) GetObject() ProtectedObjectInfo {
	if o == nil || o.Object == nil {
		var ret ProtectedObjectInfo
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProtectedObjectResponse) GetObjectOk() (*ProtectedObjectInfo, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *GetProtectedObjectResponse) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given ProtectedObjectInfo and assigns it to the Object field.
func (o *GetProtectedObjectResponse) SetObject(v ProtectedObjectInfo) {
	o.Object = &v
}

func (o GetProtectedObjectResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	return json.Marshal(toSerialize)
}

type NullableGetProtectedObjectResponse struct {
	value *GetProtectedObjectResponse
	isSet bool
}

func (v NullableGetProtectedObjectResponse) Get() *GetProtectedObjectResponse {
	return v.value
}

func (v *NullableGetProtectedObjectResponse) Set(val *GetProtectedObjectResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProtectedObjectResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProtectedObjectResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProtectedObjectResponse(val *GetProtectedObjectResponse) *NullableGetProtectedObjectResponse {
	return &NullableGetProtectedObjectResponse{value: val, isSet: true}
}

func (v NullableGetProtectedObjectResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProtectedObjectResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o GetProtectedObjectResponse) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}