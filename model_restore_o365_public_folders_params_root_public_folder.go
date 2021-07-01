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

// RestoreO365PublicFoldersParamsRootPublicFolder struct for RestoreO365PublicFoldersParamsRootPublicFolder
type RestoreO365PublicFoldersParamsRootPublicFolder struct {
	// If is_entire_folder_required is set to false, user will then specify which particular sub-folders are to be restored.
	FolderVec []RestoreO365PublicFoldersParamsPublicFolder `json:"folderVec,omitempty"`
	// Specify if the entire Root Public Folder including the sub-folders is to be restored.
	IsEntireRootFolderRequired NullableBool `json:"isEntireRootFolderRequired,omitempty"`
	Object *RestoreObject `json:"object,omitempty"`
}

// NewRestoreO365PublicFoldersParamsRootPublicFolder instantiates a new RestoreO365PublicFoldersParamsRootPublicFolder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestoreO365PublicFoldersParamsRootPublicFolder() *RestoreO365PublicFoldersParamsRootPublicFolder {
	this := RestoreO365PublicFoldersParamsRootPublicFolder{}
	return &this
}

// NewRestoreO365PublicFoldersParamsRootPublicFolderWithDefaults instantiates a new RestoreO365PublicFoldersParamsRootPublicFolder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestoreO365PublicFoldersParamsRootPublicFolderWithDefaults() *RestoreO365PublicFoldersParamsRootPublicFolder {
	this := RestoreO365PublicFoldersParamsRootPublicFolder{}
	return &this
}

// GetFolderVec returns the FolderVec field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RestoreO365PublicFoldersParamsRootPublicFolder) GetFolderVec() []RestoreO365PublicFoldersParamsPublicFolder {
	if o == nil  {
		var ret []RestoreO365PublicFoldersParamsPublicFolder
		return ret
	}
	return o.FolderVec
}

// GetFolderVecOk returns a tuple with the FolderVec field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RestoreO365PublicFoldersParamsRootPublicFolder) GetFolderVecOk() (*[]RestoreO365PublicFoldersParamsPublicFolder, bool) {
	if o == nil || o.FolderVec == nil {
		return nil, false
	}
	return &o.FolderVec, true
}

// HasFolderVec returns a boolean if a field has been set.
func (o *RestoreO365PublicFoldersParamsRootPublicFolder) HasFolderVec() bool {
	if o != nil && o.FolderVec != nil {
		return true
	}

	return false
}

// SetFolderVec gets a reference to the given []RestoreO365PublicFoldersParamsPublicFolder and assigns it to the FolderVec field.
func (o *RestoreO365PublicFoldersParamsRootPublicFolder) SetFolderVec(v []RestoreO365PublicFoldersParamsPublicFolder) {
	o.FolderVec = v
}

// GetIsEntireRootFolderRequired returns the IsEntireRootFolderRequired field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RestoreO365PublicFoldersParamsRootPublicFolder) GetIsEntireRootFolderRequired() bool {
	if o == nil || o.IsEntireRootFolderRequired.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsEntireRootFolderRequired.Get()
}

// GetIsEntireRootFolderRequiredOk returns a tuple with the IsEntireRootFolderRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RestoreO365PublicFoldersParamsRootPublicFolder) GetIsEntireRootFolderRequiredOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsEntireRootFolderRequired.Get(), o.IsEntireRootFolderRequired.IsSet()
}

// HasIsEntireRootFolderRequired returns a boolean if a field has been set.
func (o *RestoreO365PublicFoldersParamsRootPublicFolder) HasIsEntireRootFolderRequired() bool {
	if o != nil && o.IsEntireRootFolderRequired.IsSet() {
		return true
	}

	return false
}

// SetIsEntireRootFolderRequired gets a reference to the given NullableBool and assigns it to the IsEntireRootFolderRequired field.
func (o *RestoreO365PublicFoldersParamsRootPublicFolder) SetIsEntireRootFolderRequired(v bool) {
	o.IsEntireRootFolderRequired.Set(&v)
}
// SetIsEntireRootFolderRequiredNil sets the value for IsEntireRootFolderRequired to be an explicit nil
func (o *RestoreO365PublicFoldersParamsRootPublicFolder) SetIsEntireRootFolderRequiredNil() {
	o.IsEntireRootFolderRequired.Set(nil)
}

// UnsetIsEntireRootFolderRequired ensures that no value is present for IsEntireRootFolderRequired, not even an explicit nil
func (o *RestoreO365PublicFoldersParamsRootPublicFolder) UnsetIsEntireRootFolderRequired() {
	o.IsEntireRootFolderRequired.Unset()
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *RestoreO365PublicFoldersParamsRootPublicFolder) GetObject() RestoreObject {
	if o == nil || o.Object == nil {
		var ret RestoreObject
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreO365PublicFoldersParamsRootPublicFolder) GetObjectOk() (*RestoreObject, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *RestoreO365PublicFoldersParamsRootPublicFolder) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given RestoreObject and assigns it to the Object field.
func (o *RestoreO365PublicFoldersParamsRootPublicFolder) SetObject(v RestoreObject) {
	o.Object = &v
}

func (o RestoreO365PublicFoldersParamsRootPublicFolder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FolderVec != nil {
		toSerialize["folderVec"] = o.FolderVec
	}
	if o.IsEntireRootFolderRequired.IsSet() {
		toSerialize["isEntireRootFolderRequired"] = o.IsEntireRootFolderRequired.Get()
	}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	return json.Marshal(toSerialize)
}

type NullableRestoreO365PublicFoldersParamsRootPublicFolder struct {
	value *RestoreO365PublicFoldersParamsRootPublicFolder
	isSet bool
}

func (v NullableRestoreO365PublicFoldersParamsRootPublicFolder) Get() *RestoreO365PublicFoldersParamsRootPublicFolder {
	return v.value
}

func (v *NullableRestoreO365PublicFoldersParamsRootPublicFolder) Set(val *RestoreO365PublicFoldersParamsRootPublicFolder) {
	v.value = val
	v.isSet = true
}

func (v NullableRestoreO365PublicFoldersParamsRootPublicFolder) IsSet() bool {
	return v.isSet
}

func (v *NullableRestoreO365PublicFoldersParamsRootPublicFolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestoreO365PublicFoldersParamsRootPublicFolder(val *RestoreO365PublicFoldersParamsRootPublicFolder) *NullableRestoreO365PublicFoldersParamsRootPublicFolder {
	return &NullableRestoreO365PublicFoldersParamsRootPublicFolder{value: val, isSet: true}
}

func (v NullableRestoreO365PublicFoldersParamsRootPublicFolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestoreO365PublicFoldersParamsRootPublicFolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


