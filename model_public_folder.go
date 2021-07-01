/*
 * Cohesity REST API
 *
 * This API list provides operations for interfacing with the Cohesity Cluster.
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cohesitysdk

import (
	"encoding/json"
)

// PublicFolder Specifies the O365 PublicFolder details.
type PublicFolder struct {
	// Specifies the unique ID of the folder.
	FolderId NullableString `json:"folderId,omitempty"`
	// Specifies the PublicFolder items within the folder to be restored incase the user wishes not to restore the entire folder.
	PublicFolderItemIdList []string `json:"publicFolderItemIdList,omitempty"`
	// Specifies whether the entire folder is to be restored.
	RestoreEntireFolder NullableBool `json:"restoreEntireFolder,omitempty"`
}

// NewPublicFolder instantiates a new PublicFolder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicFolder() *PublicFolder {
	this := PublicFolder{}
	return &this
}

// NewPublicFolderWithDefaults instantiates a new PublicFolder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicFolderWithDefaults() *PublicFolder {
	this := PublicFolder{}
	return &this
}

// GetFolderId returns the FolderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PublicFolder) GetFolderId() string {
	if o == nil || o.FolderId.Get() == nil {
		var ret string
		return ret
	}
	return *o.FolderId.Get()
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PublicFolder) GetFolderIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FolderId.Get(), o.FolderId.IsSet()
}

// HasFolderId returns a boolean if a field has been set.
func (o *PublicFolder) HasFolderId() bool {
	if o != nil && o.FolderId.IsSet() {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given NullableString and assigns it to the FolderId field.
func (o *PublicFolder) SetFolderId(v string) {
	o.FolderId.Set(&v)
}
// SetFolderIdNil sets the value for FolderId to be an explicit nil
func (o *PublicFolder) SetFolderIdNil() {
	o.FolderId.Set(nil)
}

// UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
func (o *PublicFolder) UnsetFolderId() {
	o.FolderId.Unset()
}

// GetPublicFolderItemIdList returns the PublicFolderItemIdList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PublicFolder) GetPublicFolderItemIdList() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.PublicFolderItemIdList
}

// GetPublicFolderItemIdListOk returns a tuple with the PublicFolderItemIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PublicFolder) GetPublicFolderItemIdListOk() (*[]string, bool) {
	if o == nil || o.PublicFolderItemIdList == nil {
		return nil, false
	}
	return &o.PublicFolderItemIdList, true
}

// HasPublicFolderItemIdList returns a boolean if a field has been set.
func (o *PublicFolder) HasPublicFolderItemIdList() bool {
	if o != nil && o.PublicFolderItemIdList != nil {
		return true
	}

	return false
}

// SetPublicFolderItemIdList gets a reference to the given []string and assigns it to the PublicFolderItemIdList field.
func (o *PublicFolder) SetPublicFolderItemIdList(v []string) {
	o.PublicFolderItemIdList = v
}

// GetRestoreEntireFolder returns the RestoreEntireFolder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PublicFolder) GetRestoreEntireFolder() bool {
	if o == nil || o.RestoreEntireFolder.Get() == nil {
		var ret bool
		return ret
	}
	return *o.RestoreEntireFolder.Get()
}

// GetRestoreEntireFolderOk returns a tuple with the RestoreEntireFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PublicFolder) GetRestoreEntireFolderOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RestoreEntireFolder.Get(), o.RestoreEntireFolder.IsSet()
}

// HasRestoreEntireFolder returns a boolean if a field has been set.
func (o *PublicFolder) HasRestoreEntireFolder() bool {
	if o != nil && o.RestoreEntireFolder.IsSet() {
		return true
	}

	return false
}

// SetRestoreEntireFolder gets a reference to the given NullableBool and assigns it to the RestoreEntireFolder field.
func (o *PublicFolder) SetRestoreEntireFolder(v bool) {
	o.RestoreEntireFolder.Set(&v)
}
// SetRestoreEntireFolderNil sets the value for RestoreEntireFolder to be an explicit nil
func (o *PublicFolder) SetRestoreEntireFolderNil() {
	o.RestoreEntireFolder.Set(nil)
}

// UnsetRestoreEntireFolder ensures that no value is present for RestoreEntireFolder, not even an explicit nil
func (o *PublicFolder) UnsetRestoreEntireFolder() {
	o.RestoreEntireFolder.Unset()
}

func (o PublicFolder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FolderId.IsSet() {
		toSerialize["folderId"] = o.FolderId.Get()
	}
	if o.PublicFolderItemIdList != nil {
		toSerialize["publicFolderItemIdList"] = o.PublicFolderItemIdList
	}
	if o.RestoreEntireFolder.IsSet() {
		toSerialize["restoreEntireFolder"] = o.RestoreEntireFolder.Get()
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


