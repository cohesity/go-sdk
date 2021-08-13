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

// ObjectOneDriveParam Specifies OneDrive recovery parameters associated with a user.
type ObjectOneDriveParam struct {
	// Specifies the OneDrive owner info.
	OwnerInfo *CommonRecoverObjectSnapshotParams `json:"ownerInfo,omitempty"`
	// Specifies parameters to recover a OneDrive.
	OneDriveParams []OneDriveParam `json:"oneDriveParams,omitempty"`
}

// NewObjectOneDriveParam instantiates a new ObjectOneDriveParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectOneDriveParam() *ObjectOneDriveParam {
	this := ObjectOneDriveParam{}
	return &this
}

// NewObjectOneDriveParamWithDefaults instantiates a new ObjectOneDriveParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectOneDriveParamWithDefaults() *ObjectOneDriveParam {
	this := ObjectOneDriveParam{}
	return &this
}

// GetOwnerInfo returns the OwnerInfo field value if set, zero value otherwise.
func (o *ObjectOneDriveParam) GetOwnerInfo() CommonRecoverObjectSnapshotParams {
	if o == nil || o.OwnerInfo == nil {
		var ret CommonRecoverObjectSnapshotParams
		return ret
	}
	return *o.OwnerInfo
}

// GetOwnerInfoOk returns a tuple with the OwnerInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectOneDriveParam) GetOwnerInfoOk() (*CommonRecoverObjectSnapshotParams, bool) {
	if o == nil || o.OwnerInfo == nil {
		return nil, false
	}
	return o.OwnerInfo, true
}

// HasOwnerInfo returns a boolean if a field has been set.
func (o *ObjectOneDriveParam) HasOwnerInfo() bool {
	if o != nil && o.OwnerInfo != nil {
		return true
	}

	return false
}

// SetOwnerInfo gets a reference to the given CommonRecoverObjectSnapshotParams and assigns it to the OwnerInfo field.
func (o *ObjectOneDriveParam) SetOwnerInfo(v CommonRecoverObjectSnapshotParams) {
	o.OwnerInfo = &v
}

// GetOneDriveParams returns the OneDriveParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ObjectOneDriveParam) GetOneDriveParams() []OneDriveParam {
	if o == nil  {
		var ret []OneDriveParam
		return ret
	}
	return o.OneDriveParams
}

// GetOneDriveParamsOk returns a tuple with the OneDriveParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ObjectOneDriveParam) GetOneDriveParamsOk() (*[]OneDriveParam, bool) {
	if o == nil || o.OneDriveParams == nil {
		return nil, false
	}
	return &o.OneDriveParams, true
}

// HasOneDriveParams returns a boolean if a field has been set.
func (o *ObjectOneDriveParam) HasOneDriveParams() bool {
	if o != nil && o.OneDriveParams != nil {
		return true
	}

	return false
}

// SetOneDriveParams gets a reference to the given []OneDriveParam and assigns it to the OneDriveParams field.
func (o *ObjectOneDriveParam) SetOneDriveParams(v []OneDriveParam) {
	o.OneDriveParams = v
}

func (o ObjectOneDriveParam) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OwnerInfo != nil {
		toSerialize["ownerInfo"] = o.OwnerInfo
	}
	if o.OneDriveParams != nil {
		toSerialize["oneDriveParams"] = o.OneDriveParams
	}
	return json.Marshal(toSerialize)
}

type NullableObjectOneDriveParam struct {
	value *ObjectOneDriveParam
	isSet bool
}

func (v NullableObjectOneDriveParam) Get() *ObjectOneDriveParam {
	return v.value
}

func (v *NullableObjectOneDriveParam) Set(val *ObjectOneDriveParam) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectOneDriveParam) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectOneDriveParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectOneDriveParam(val *ObjectOneDriveParam) *NullableObjectOneDriveParam {
	return &NullableObjectOneDriveParam{value: val, isSet: true}
}

func (v NullableObjectOneDriveParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectOneDriveParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o ObjectOneDriveParam) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}