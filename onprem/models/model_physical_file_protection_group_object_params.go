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

// PhysicalFileProtectionGroupObjectParams struct for PhysicalFileProtectionGroupObjectParams
type PhysicalFileProtectionGroupObjectParams struct {
	// Specifies the ID of the object protected.
	Id int64 `json:"id"`
	// Specifies the name of the object protected.
	Name NullableString `json:"name,omitempty"`
	// Specifies a list of file paths to be protected by this Protection Group.
	FilePaths *[]PhysicalFileBackupPathParams `json:"filePaths,omitempty"`
	// Specifies whether path level or object level skip nested volume setting will be used.
	UsesPathLevelSkipNestedVolumeSetting NullableBool `json:"usesPathLevelSkipNestedVolumeSetting,omitempty"`
	// Specifies mount types of nested volumes to be skipped.
	NestedVolumeTypesToSkip *[]string `json:"nestedVolumeTypesToSkip,omitempty"`
	// Specifies whether to follow NAS target pointed by symlink for windows sources.
	FollowNasSymlinkTarget NullableBool `json:"followNasSymlinkTarget,omitempty"`
	// Specifies the path of metadatafile on source. This file contains absolute paths of files that needs to be backed up on the same source.
	MetadataFilePath NullableString `json:"metadataFilePath,omitempty"`
}

// NewPhysicalFileProtectionGroupObjectParams instantiates a new PhysicalFileProtectionGroupObjectParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhysicalFileProtectionGroupObjectParams(id int64) *PhysicalFileProtectionGroupObjectParams {
	this := PhysicalFileProtectionGroupObjectParams{}
	this.Id = id
	return &this
}

// NewPhysicalFileProtectionGroupObjectParamsWithDefaults instantiates a new PhysicalFileProtectionGroupObjectParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhysicalFileProtectionGroupObjectParamsWithDefaults() *PhysicalFileProtectionGroupObjectParams {
	this := PhysicalFileProtectionGroupObjectParams{}
	return &this
}

// GetId returns the Id field value
func (o *PhysicalFileProtectionGroupObjectParams) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PhysicalFileProtectionGroupObjectParams) GetIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PhysicalFileProtectionGroupObjectParams) SetId(v int64) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PhysicalFileProtectionGroupObjectParams) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PhysicalFileProtectionGroupObjectParams) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *PhysicalFileProtectionGroupObjectParams) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *PhysicalFileProtectionGroupObjectParams) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *PhysicalFileProtectionGroupObjectParams) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *PhysicalFileProtectionGroupObjectParams) UnsetName() {
	o.Name.Unset()
}

// GetFilePaths returns the FilePaths field value if set, zero value otherwise.
func (o *PhysicalFileProtectionGroupObjectParams) GetFilePaths() []PhysicalFileBackupPathParams {
	if o == nil || o.FilePaths == nil {
		var ret []PhysicalFileBackupPathParams
		return ret
	}
	return *o.FilePaths
}

// GetFilePathsOk returns a tuple with the FilePaths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalFileProtectionGroupObjectParams) GetFilePathsOk() (*[]PhysicalFileBackupPathParams, bool) {
	if o == nil || o.FilePaths == nil {
		return nil, false
	}
	return o.FilePaths, true
}

// HasFilePaths returns a boolean if a field has been set.
func (o *PhysicalFileProtectionGroupObjectParams) HasFilePaths() bool {
	if o != nil && o.FilePaths != nil {
		return true
	}

	return false
}

// SetFilePaths gets a reference to the given []PhysicalFileBackupPathParams and assigns it to the FilePaths field.
func (o *PhysicalFileProtectionGroupObjectParams) SetFilePaths(v []PhysicalFileBackupPathParams) {
	o.FilePaths = &v
}

// GetUsesPathLevelSkipNestedVolumeSetting returns the UsesPathLevelSkipNestedVolumeSetting field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PhysicalFileProtectionGroupObjectParams) GetUsesPathLevelSkipNestedVolumeSetting() bool {
	if o == nil || o.UsesPathLevelSkipNestedVolumeSetting.Get() == nil {
		var ret bool
		return ret
	}
	return *o.UsesPathLevelSkipNestedVolumeSetting.Get()
}

// GetUsesPathLevelSkipNestedVolumeSettingOk returns a tuple with the UsesPathLevelSkipNestedVolumeSetting field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PhysicalFileProtectionGroupObjectParams) GetUsesPathLevelSkipNestedVolumeSettingOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UsesPathLevelSkipNestedVolumeSetting.Get(), o.UsesPathLevelSkipNestedVolumeSetting.IsSet()
}

// HasUsesPathLevelSkipNestedVolumeSetting returns a boolean if a field has been set.
func (o *PhysicalFileProtectionGroupObjectParams) HasUsesPathLevelSkipNestedVolumeSetting() bool {
	if o != nil && o.UsesPathLevelSkipNestedVolumeSetting.IsSet() {
		return true
	}

	return false
}

// SetUsesPathLevelSkipNestedVolumeSetting gets a reference to the given NullableBool and assigns it to the UsesPathLevelSkipNestedVolumeSetting field.
func (o *PhysicalFileProtectionGroupObjectParams) SetUsesPathLevelSkipNestedVolumeSetting(v bool) {
	o.UsesPathLevelSkipNestedVolumeSetting.Set(&v)
}
// SetUsesPathLevelSkipNestedVolumeSettingNil sets the value for UsesPathLevelSkipNestedVolumeSetting to be an explicit nil
func (o *PhysicalFileProtectionGroupObjectParams) SetUsesPathLevelSkipNestedVolumeSettingNil() {
	o.UsesPathLevelSkipNestedVolumeSetting.Set(nil)
}

// UnsetUsesPathLevelSkipNestedVolumeSetting ensures that no value is present for UsesPathLevelSkipNestedVolumeSetting, not even an explicit nil
func (o *PhysicalFileProtectionGroupObjectParams) UnsetUsesPathLevelSkipNestedVolumeSetting() {
	o.UsesPathLevelSkipNestedVolumeSetting.Unset()
}

// GetNestedVolumeTypesToSkip returns the NestedVolumeTypesToSkip field value if set, zero value otherwise.
func (o *PhysicalFileProtectionGroupObjectParams) GetNestedVolumeTypesToSkip() []string {
	if o == nil || o.NestedVolumeTypesToSkip == nil {
		var ret []string
		return ret
	}
	return *o.NestedVolumeTypesToSkip
}

// GetNestedVolumeTypesToSkipOk returns a tuple with the NestedVolumeTypesToSkip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalFileProtectionGroupObjectParams) GetNestedVolumeTypesToSkipOk() (*[]string, bool) {
	if o == nil || o.NestedVolumeTypesToSkip == nil {
		return nil, false
	}
	return o.NestedVolumeTypesToSkip, true
}

// HasNestedVolumeTypesToSkip returns a boolean if a field has been set.
func (o *PhysicalFileProtectionGroupObjectParams) HasNestedVolumeTypesToSkip() bool {
	if o != nil && o.NestedVolumeTypesToSkip != nil {
		return true
	}

	return false
}

// SetNestedVolumeTypesToSkip gets a reference to the given []string and assigns it to the NestedVolumeTypesToSkip field.
func (o *PhysicalFileProtectionGroupObjectParams) SetNestedVolumeTypesToSkip(v []string) {
	o.NestedVolumeTypesToSkip = &v
}

// GetFollowNasSymlinkTarget returns the FollowNasSymlinkTarget field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PhysicalFileProtectionGroupObjectParams) GetFollowNasSymlinkTarget() bool {
	if o == nil || o.FollowNasSymlinkTarget.Get() == nil {
		var ret bool
		return ret
	}
	return *o.FollowNasSymlinkTarget.Get()
}

// GetFollowNasSymlinkTargetOk returns a tuple with the FollowNasSymlinkTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PhysicalFileProtectionGroupObjectParams) GetFollowNasSymlinkTargetOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FollowNasSymlinkTarget.Get(), o.FollowNasSymlinkTarget.IsSet()
}

// HasFollowNasSymlinkTarget returns a boolean if a field has been set.
func (o *PhysicalFileProtectionGroupObjectParams) HasFollowNasSymlinkTarget() bool {
	if o != nil && o.FollowNasSymlinkTarget.IsSet() {
		return true
	}

	return false
}

// SetFollowNasSymlinkTarget gets a reference to the given NullableBool and assigns it to the FollowNasSymlinkTarget field.
func (o *PhysicalFileProtectionGroupObjectParams) SetFollowNasSymlinkTarget(v bool) {
	o.FollowNasSymlinkTarget.Set(&v)
}
// SetFollowNasSymlinkTargetNil sets the value for FollowNasSymlinkTarget to be an explicit nil
func (o *PhysicalFileProtectionGroupObjectParams) SetFollowNasSymlinkTargetNil() {
	o.FollowNasSymlinkTarget.Set(nil)
}

// UnsetFollowNasSymlinkTarget ensures that no value is present for FollowNasSymlinkTarget, not even an explicit nil
func (o *PhysicalFileProtectionGroupObjectParams) UnsetFollowNasSymlinkTarget() {
	o.FollowNasSymlinkTarget.Unset()
}

// GetMetadataFilePath returns the MetadataFilePath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PhysicalFileProtectionGroupObjectParams) GetMetadataFilePath() string {
	if o == nil || o.MetadataFilePath.Get() == nil {
		var ret string
		return ret
	}
	return *o.MetadataFilePath.Get()
}

// GetMetadataFilePathOk returns a tuple with the MetadataFilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PhysicalFileProtectionGroupObjectParams) GetMetadataFilePathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MetadataFilePath.Get(), o.MetadataFilePath.IsSet()
}

// HasMetadataFilePath returns a boolean if a field has been set.
func (o *PhysicalFileProtectionGroupObjectParams) HasMetadataFilePath() bool {
	if o != nil && o.MetadataFilePath.IsSet() {
		return true
	}

	return false
}

// SetMetadataFilePath gets a reference to the given NullableString and assigns it to the MetadataFilePath field.
func (o *PhysicalFileProtectionGroupObjectParams) SetMetadataFilePath(v string) {
	o.MetadataFilePath.Set(&v)
}
// SetMetadataFilePathNil sets the value for MetadataFilePath to be an explicit nil
func (o *PhysicalFileProtectionGroupObjectParams) SetMetadataFilePathNil() {
	o.MetadataFilePath.Set(nil)
}

// UnsetMetadataFilePath ensures that no value is present for MetadataFilePath, not even an explicit nil
func (o *PhysicalFileProtectionGroupObjectParams) UnsetMetadataFilePath() {
	o.MetadataFilePath.Unset()
}

func (o PhysicalFileProtectionGroupObjectParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.FilePaths != nil {
		toSerialize["filePaths"] = o.FilePaths
	}
	if o.UsesPathLevelSkipNestedVolumeSetting.IsSet() {
		toSerialize["usesPathLevelSkipNestedVolumeSetting"] = o.UsesPathLevelSkipNestedVolumeSetting.Get()
	}
	if o.NestedVolumeTypesToSkip != nil {
		toSerialize["nestedVolumeTypesToSkip"] = o.NestedVolumeTypesToSkip
	}
	if o.FollowNasSymlinkTarget.IsSet() {
		toSerialize["followNasSymlinkTarget"] = o.FollowNasSymlinkTarget.Get()
	}
	if o.MetadataFilePath.IsSet() {
		toSerialize["metadataFilePath"] = o.MetadataFilePath.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePhysicalFileProtectionGroupObjectParams struct {
	value *PhysicalFileProtectionGroupObjectParams
	isSet bool
}

func (v NullablePhysicalFileProtectionGroupObjectParams) Get() *PhysicalFileProtectionGroupObjectParams {
	return v.value
}

func (v *NullablePhysicalFileProtectionGroupObjectParams) Set(val *PhysicalFileProtectionGroupObjectParams) {
	v.value = val
	v.isSet = true
}

func (v NullablePhysicalFileProtectionGroupObjectParams) IsSet() bool {
	return v.isSet
}

func (v *NullablePhysicalFileProtectionGroupObjectParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhysicalFileProtectionGroupObjectParams(val *PhysicalFileProtectionGroupObjectParams) *NullablePhysicalFileProtectionGroupObjectParams {
	return &NullablePhysicalFileProtectionGroupObjectParams{value: val, isSet: true}
}

func (v NullablePhysicalFileProtectionGroupObjectParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhysicalFileProtectionGroupObjectParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o PhysicalFileProtectionGroupObjectParams) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}