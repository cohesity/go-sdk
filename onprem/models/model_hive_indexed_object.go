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

// HiveIndexedObject Specifies a Hive indexed object.
type HiveIndexedObject struct {
	// Specifies the name of the object.
	Name NullableString `json:"name,omitempty"`
	// Specifies the path of the object.
	Path NullableString `json:"path,omitempty"`
	// \"Specifies the protection group id which contains this object.\"
	ProtectionGroupId NullableString `json:"protectionGroupId,omitempty"`
	// \"Specifies the protection group name which contains this object.\"
	ProtectionGroupName NullableString `json:"protectionGroupName,omitempty"`
	// \"Specifies the Storage Domain id where the backup data of Object is present.\"
	StorageDomainId NullableInt64 `json:"storageDomainId,omitempty"`
	// Specifies the Source Object information.
	SourceInfo *ObjectSummary `json:"sourceInfo,omitempty"`
	// Specifies tag applied to the object.
	Tags []TagInfo `json:"tags,omitempty"`
	// Specifies snapshot tags applied to the object.
	SnapshotTags []SnapshotTagInfo `json:"snapshotTags,omitempty"`
	// Specifies the Hive Object Type.
	Type NullableString `json:"type,omitempty"`
	// Specifies the id of the indexed object.
	Id NullableString `json:"id,omitempty"`
}

// NewHiveIndexedObject instantiates a new HiveIndexedObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHiveIndexedObject() *HiveIndexedObject {
	this := HiveIndexedObject{}
	return &this
}

// NewHiveIndexedObjectWithDefaults instantiates a new HiveIndexedObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHiveIndexedObjectWithDefaults() *HiveIndexedObject {
	this := HiveIndexedObject{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HiveIndexedObject) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HiveIndexedObject) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *HiveIndexedObject) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *HiveIndexedObject) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *HiveIndexedObject) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *HiveIndexedObject) UnsetName() {
	o.Name.Unset()
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HiveIndexedObject) GetPath() string {
	if o == nil || o.Path.Get() == nil {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HiveIndexedObject) GetPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *HiveIndexedObject) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *HiveIndexedObject) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *HiveIndexedObject) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *HiveIndexedObject) UnsetPath() {
	o.Path.Unset()
}

// GetProtectionGroupId returns the ProtectionGroupId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HiveIndexedObject) GetProtectionGroupId() string {
	if o == nil || o.ProtectionGroupId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProtectionGroupId.Get()
}

// GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HiveIndexedObject) GetProtectionGroupIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ProtectionGroupId.Get(), o.ProtectionGroupId.IsSet()
}

// HasProtectionGroupId returns a boolean if a field has been set.
func (o *HiveIndexedObject) HasProtectionGroupId() bool {
	if o != nil && o.ProtectionGroupId.IsSet() {
		return true
	}

	return false
}

// SetProtectionGroupId gets a reference to the given NullableString and assigns it to the ProtectionGroupId field.
func (o *HiveIndexedObject) SetProtectionGroupId(v string) {
	o.ProtectionGroupId.Set(&v)
}
// SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil
func (o *HiveIndexedObject) SetProtectionGroupIdNil() {
	o.ProtectionGroupId.Set(nil)
}

// UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil
func (o *HiveIndexedObject) UnsetProtectionGroupId() {
	o.ProtectionGroupId.Unset()
}

// GetProtectionGroupName returns the ProtectionGroupName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HiveIndexedObject) GetProtectionGroupName() string {
	if o == nil || o.ProtectionGroupName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProtectionGroupName.Get()
}

// GetProtectionGroupNameOk returns a tuple with the ProtectionGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HiveIndexedObject) GetProtectionGroupNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ProtectionGroupName.Get(), o.ProtectionGroupName.IsSet()
}

// HasProtectionGroupName returns a boolean if a field has been set.
func (o *HiveIndexedObject) HasProtectionGroupName() bool {
	if o != nil && o.ProtectionGroupName.IsSet() {
		return true
	}

	return false
}

// SetProtectionGroupName gets a reference to the given NullableString and assigns it to the ProtectionGroupName field.
func (o *HiveIndexedObject) SetProtectionGroupName(v string) {
	o.ProtectionGroupName.Set(&v)
}
// SetProtectionGroupNameNil sets the value for ProtectionGroupName to be an explicit nil
func (o *HiveIndexedObject) SetProtectionGroupNameNil() {
	o.ProtectionGroupName.Set(nil)
}

// UnsetProtectionGroupName ensures that no value is present for ProtectionGroupName, not even an explicit nil
func (o *HiveIndexedObject) UnsetProtectionGroupName() {
	o.ProtectionGroupName.Unset()
}

// GetStorageDomainId returns the StorageDomainId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HiveIndexedObject) GetStorageDomainId() int64 {
	if o == nil || o.StorageDomainId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.StorageDomainId.Get()
}

// GetStorageDomainIdOk returns a tuple with the StorageDomainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HiveIndexedObject) GetStorageDomainIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StorageDomainId.Get(), o.StorageDomainId.IsSet()
}

// HasStorageDomainId returns a boolean if a field has been set.
func (o *HiveIndexedObject) HasStorageDomainId() bool {
	if o != nil && o.StorageDomainId.IsSet() {
		return true
	}

	return false
}

// SetStorageDomainId gets a reference to the given NullableInt64 and assigns it to the StorageDomainId field.
func (o *HiveIndexedObject) SetStorageDomainId(v int64) {
	o.StorageDomainId.Set(&v)
}
// SetStorageDomainIdNil sets the value for StorageDomainId to be an explicit nil
func (o *HiveIndexedObject) SetStorageDomainIdNil() {
	o.StorageDomainId.Set(nil)
}

// UnsetStorageDomainId ensures that no value is present for StorageDomainId, not even an explicit nil
func (o *HiveIndexedObject) UnsetStorageDomainId() {
	o.StorageDomainId.Unset()
}

// GetSourceInfo returns the SourceInfo field value if set, zero value otherwise.
func (o *HiveIndexedObject) GetSourceInfo() ObjectSummary {
	if o == nil || o.SourceInfo == nil {
		var ret ObjectSummary
		return ret
	}
	return *o.SourceInfo
}

// GetSourceInfoOk returns a tuple with the SourceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HiveIndexedObject) GetSourceInfoOk() (*ObjectSummary, bool) {
	if o == nil || o.SourceInfo == nil {
		return nil, false
	}
	return o.SourceInfo, true
}

// HasSourceInfo returns a boolean if a field has been set.
func (o *HiveIndexedObject) HasSourceInfo() bool {
	if o != nil && o.SourceInfo != nil {
		return true
	}

	return false
}

// SetSourceInfo gets a reference to the given ObjectSummary and assigns it to the SourceInfo field.
func (o *HiveIndexedObject) SetSourceInfo(v ObjectSummary) {
	o.SourceInfo = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HiveIndexedObject) GetTags() []TagInfo {
	if o == nil  {
		var ret []TagInfo
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HiveIndexedObject) GetTagsOk() (*[]TagInfo, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *HiveIndexedObject) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagInfo and assigns it to the Tags field.
func (o *HiveIndexedObject) SetTags(v []TagInfo) {
	o.Tags = v
}

// GetSnapshotTags returns the SnapshotTags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HiveIndexedObject) GetSnapshotTags() []SnapshotTagInfo {
	if o == nil  {
		var ret []SnapshotTagInfo
		return ret
	}
	return o.SnapshotTags
}

// GetSnapshotTagsOk returns a tuple with the SnapshotTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HiveIndexedObject) GetSnapshotTagsOk() (*[]SnapshotTagInfo, bool) {
	if o == nil || o.SnapshotTags == nil {
		return nil, false
	}
	return &o.SnapshotTags, true
}

// HasSnapshotTags returns a boolean if a field has been set.
func (o *HiveIndexedObject) HasSnapshotTags() bool {
	if o != nil && o.SnapshotTags != nil {
		return true
	}

	return false
}

// SetSnapshotTags gets a reference to the given []SnapshotTagInfo and assigns it to the SnapshotTags field.
func (o *HiveIndexedObject) SetSnapshotTags(v []SnapshotTagInfo) {
	o.SnapshotTags = v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HiveIndexedObject) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HiveIndexedObject) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *HiveIndexedObject) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *HiveIndexedObject) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *HiveIndexedObject) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *HiveIndexedObject) UnsetType() {
	o.Type.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HiveIndexedObject) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HiveIndexedObject) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *HiveIndexedObject) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *HiveIndexedObject) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *HiveIndexedObject) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *HiveIndexedObject) UnsetId() {
	o.Id.Unset()
}

func (o HiveIndexedObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Path.IsSet() {
		toSerialize["path"] = o.Path.Get()
	}
	if o.ProtectionGroupId.IsSet() {
		toSerialize["protectionGroupId"] = o.ProtectionGroupId.Get()
	}
	if o.ProtectionGroupName.IsSet() {
		toSerialize["protectionGroupName"] = o.ProtectionGroupName.Get()
	}
	if o.StorageDomainId.IsSet() {
		toSerialize["storageDomainId"] = o.StorageDomainId.Get()
	}
	if o.SourceInfo != nil {
		toSerialize["sourceInfo"] = o.SourceInfo
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.SnapshotTags != nil {
		toSerialize["snapshotTags"] = o.SnapshotTags
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableHiveIndexedObject struct {
	value *HiveIndexedObject
	isSet bool
}

func (v NullableHiveIndexedObject) Get() *HiveIndexedObject {
	return v.value
}

func (v *NullableHiveIndexedObject) Set(val *HiveIndexedObject) {
	v.value = val
	v.isSet = true
}

func (v NullableHiveIndexedObject) IsSet() bool {
	return v.isSet
}

func (v *NullableHiveIndexedObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHiveIndexedObject(val *HiveIndexedObject) *NullableHiveIndexedObject {
	return &NullableHiveIndexedObject{value: val, isSet: true}
}

func (v NullableHiveIndexedObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHiveIndexedObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o HiveIndexedObject) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}