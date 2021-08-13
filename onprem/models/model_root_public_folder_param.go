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

// RootPublicFolderParam Specifies parameters to recover a RootPublicFolder.
type RootPublicFolderParam struct {
	// Specifies the RootPublicFolder recover Object info.
	RecoverObject CommonRecoverObjectSnapshotParams `json:"recoverObject"`
	// Specifies whether to recover the whole RootPublicFolder.
	RecoverEntireRootPublicFolder NullableBool `json:"recoverEntireRootPublicFolder,omitempty"`
	// Specifies a list of Public Folders to recover. This field is applicable only if 'recoverEntireRootPublicFolder' is false.
	RecoverFolders []PublicFolder `json:"recoverFolders,omitempty"`
}

// NewRootPublicFolderParam instantiates a new RootPublicFolderParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRootPublicFolderParam(recoverObject CommonRecoverObjectSnapshotParams) *RootPublicFolderParam {
	this := RootPublicFolderParam{}
	this.RecoverObject = recoverObject
	return &this
}

// NewRootPublicFolderParamWithDefaults instantiates a new RootPublicFolderParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRootPublicFolderParamWithDefaults() *RootPublicFolderParam {
	this := RootPublicFolderParam{}
	return &this
}

// GetRecoverObject returns the RecoverObject field value
func (o *RootPublicFolderParam) GetRecoverObject() CommonRecoverObjectSnapshotParams {
	if o == nil {
		var ret CommonRecoverObjectSnapshotParams
		return ret
	}

	return o.RecoverObject
}

// GetRecoverObjectOk returns a tuple with the RecoverObject field value
// and a boolean to check if the value has been set.
func (o *RootPublicFolderParam) GetRecoverObjectOk() (*CommonRecoverObjectSnapshotParams, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RecoverObject, true
}

// SetRecoverObject sets field value
func (o *RootPublicFolderParam) SetRecoverObject(v CommonRecoverObjectSnapshotParams) {
	o.RecoverObject = v
}

// GetRecoverEntireRootPublicFolder returns the RecoverEntireRootPublicFolder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RootPublicFolderParam) GetRecoverEntireRootPublicFolder() bool {
	if o == nil || o.RecoverEntireRootPublicFolder.Get() == nil {
		var ret bool
		return ret
	}
	return *o.RecoverEntireRootPublicFolder.Get()
}

// GetRecoverEntireRootPublicFolderOk returns a tuple with the RecoverEntireRootPublicFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RootPublicFolderParam) GetRecoverEntireRootPublicFolderOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RecoverEntireRootPublicFolder.Get(), o.RecoverEntireRootPublicFolder.IsSet()
}

// HasRecoverEntireRootPublicFolder returns a boolean if a field has been set.
func (o *RootPublicFolderParam) HasRecoverEntireRootPublicFolder() bool {
	if o != nil && o.RecoverEntireRootPublicFolder.IsSet() {
		return true
	}

	return false
}

// SetRecoverEntireRootPublicFolder gets a reference to the given NullableBool and assigns it to the RecoverEntireRootPublicFolder field.
func (o *RootPublicFolderParam) SetRecoverEntireRootPublicFolder(v bool) {
	o.RecoverEntireRootPublicFolder.Set(&v)
}
// SetRecoverEntireRootPublicFolderNil sets the value for RecoverEntireRootPublicFolder to be an explicit nil
func (o *RootPublicFolderParam) SetRecoverEntireRootPublicFolderNil() {
	o.RecoverEntireRootPublicFolder.Set(nil)
}

// UnsetRecoverEntireRootPublicFolder ensures that no value is present for RecoverEntireRootPublicFolder, not even an explicit nil
func (o *RootPublicFolderParam) UnsetRecoverEntireRootPublicFolder() {
	o.RecoverEntireRootPublicFolder.Unset()
}

// GetRecoverFolders returns the RecoverFolders field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RootPublicFolderParam) GetRecoverFolders() []PublicFolder {
	if o == nil  {
		var ret []PublicFolder
		return ret
	}
	return o.RecoverFolders
}

// GetRecoverFoldersOk returns a tuple with the RecoverFolders field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RootPublicFolderParam) GetRecoverFoldersOk() (*[]PublicFolder, bool) {
	if o == nil || o.RecoverFolders == nil {
		return nil, false
	}
	return &o.RecoverFolders, true
}

// HasRecoverFolders returns a boolean if a field has been set.
func (o *RootPublicFolderParam) HasRecoverFolders() bool {
	if o != nil && o.RecoverFolders != nil {
		return true
	}

	return false
}

// SetRecoverFolders gets a reference to the given []PublicFolder and assigns it to the RecoverFolders field.
func (o *RootPublicFolderParam) SetRecoverFolders(v []PublicFolder) {
	o.RecoverFolders = v
}

func (o RootPublicFolderParam) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["recoverObject"] = o.RecoverObject
	}
	if o.RecoverEntireRootPublicFolder.IsSet() {
		toSerialize["recoverEntireRootPublicFolder"] = o.RecoverEntireRootPublicFolder.Get()
	}
	if o.RecoverFolders != nil {
		toSerialize["recoverFolders"] = o.RecoverFolders
	}
	return json.Marshal(toSerialize)
}

type NullableRootPublicFolderParam struct {
	value *RootPublicFolderParam
	isSet bool
}

func (v NullableRootPublicFolderParam) Get() *RootPublicFolderParam {
	return v.value
}

func (v *NullableRootPublicFolderParam) Set(val *RootPublicFolderParam) {
	v.value = val
	v.isSet = true
}

func (v NullableRootPublicFolderParam) IsSet() bool {
	return v.isSet
}

func (v *NullableRootPublicFolderParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRootPublicFolderParam(val *RootPublicFolderParam) *NullableRootPublicFolderParam {
	return &NullableRootPublicFolderParam{value: val, isSet: true}
}

func (v NullableRootPublicFolderParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRootPublicFolderParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o RootPublicFolderParam) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}