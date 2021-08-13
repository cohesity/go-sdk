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

// PublicFolder Specifies an PublicFolder item to recover.
type PublicFolder struct {
	// Specifies the Unique ID of the folder.
	FolderId NullableString `json:"folderId"`
	// Specifies whether to recover the whole PublicFolder.
	RecoverEntireFolder NullableBool `json:"recoverEntireFolder,omitempty"`
	// Specifies a list of item ids to recover. This field is applicable only if 'recoverEntireFolder' is false.
	ItemIds []string `json:"itemIds,omitempty"`
}

// NewPublicFolder instantiates a new PublicFolder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicFolder(folderId NullableString) *PublicFolder {
	this := PublicFolder{}
	this.FolderId = folderId
	return &this
}

// NewPublicFolderWithDefaults instantiates a new PublicFolder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicFolderWithDefaults() *PublicFolder {
	this := PublicFolder{}
	return &this
}

// GetFolderId returns the FolderId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PublicFolder) GetFolderId() string {
	if o == nil || o.FolderId.Get() == nil {
		var ret string
		return ret
	}

	return *o.FolderId.Get()
}

// GetFolderIdOk returns a tuple with the FolderId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PublicFolder) GetFolderIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FolderId.Get(), o.FolderId.IsSet()
}

// SetFolderId sets field value
func (o *PublicFolder) SetFolderId(v string) {
	o.FolderId.Set(&v)
}

// GetRecoverEntireFolder returns the RecoverEntireFolder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PublicFolder) GetRecoverEntireFolder() bool {
	if o == nil || o.RecoverEntireFolder.Get() == nil {
		var ret bool
		return ret
	}
	return *o.RecoverEntireFolder.Get()
}

// GetRecoverEntireFolderOk returns a tuple with the RecoverEntireFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PublicFolder) GetRecoverEntireFolderOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RecoverEntireFolder.Get(), o.RecoverEntireFolder.IsSet()
}

// HasRecoverEntireFolder returns a boolean if a field has been set.
func (o *PublicFolder) HasRecoverEntireFolder() bool {
	if o != nil && o.RecoverEntireFolder.IsSet() {
		return true
	}

	return false
}

// SetRecoverEntireFolder gets a reference to the given NullableBool and assigns it to the RecoverEntireFolder field.
func (o *PublicFolder) SetRecoverEntireFolder(v bool) {
	o.RecoverEntireFolder.Set(&v)
}
// SetRecoverEntireFolderNil sets the value for RecoverEntireFolder to be an explicit nil
func (o *PublicFolder) SetRecoverEntireFolderNil() {
	o.RecoverEntireFolder.Set(nil)
}

// UnsetRecoverEntireFolder ensures that no value is present for RecoverEntireFolder, not even an explicit nil
func (o *PublicFolder) UnsetRecoverEntireFolder() {
	o.RecoverEntireFolder.Unset()
}

// GetItemIds returns the ItemIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PublicFolder) GetItemIds() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.ItemIds
}

// GetItemIdsOk returns a tuple with the ItemIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PublicFolder) GetItemIdsOk() (*[]string, bool) {
	if o == nil || o.ItemIds == nil {
		return nil, false
	}
	return &o.ItemIds, true
}

// HasItemIds returns a boolean if a field has been set.
func (o *PublicFolder) HasItemIds() bool {
	if o != nil && o.ItemIds != nil {
		return true
	}

	return false
}

// SetItemIds gets a reference to the given []string and assigns it to the ItemIds field.
func (o *PublicFolder) SetItemIds(v []string) {
	o.ItemIds = v
}

func (o PublicFolder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["folderId"] = o.FolderId.Get()
	}
	if o.RecoverEntireFolder.IsSet() {
		toSerialize["recoverEntireFolder"] = o.RecoverEntireFolder.Get()
	}
	if o.ItemIds != nil {
		toSerialize["itemIds"] = o.ItemIds
	}
	return json.Marshal(toSerialize)
}

type NullablePublicFolder struct {
	value *PublicFolder
	isSet bool
}

func (v NullablePublicFolder) Get() *PublicFolder {
	return v.value
}

func (v *NullablePublicFolder) Set(val *PublicFolder) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicFolder) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicFolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicFolder(val *PublicFolder) *NullablePublicFolder {
	return &NullablePublicFolder{value: val, isSet: true}
}

func (v NullablePublicFolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicFolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


