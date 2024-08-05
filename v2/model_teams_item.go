/*
Cohesity REST API

Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
)

// checks if the TeamsItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TeamsItem{}

// TeamsItem Specifies the indexed M365 Teams item.
type TeamsItem struct {
	// Specifies the name of the object.
	Name NullableString `json:"name,omitempty"`
	// Specifies the path of the object.
	Path NullableString `json:"path,omitempty"`
	// Specifies the protection policy id for this file.
	PolicyId NullableString `json:"policyId,omitempty"`
	// Specifies the protection policy name for this file.
	PolicyName NullableString `json:"policyName,omitempty"`
	// \"Specifies the protection group id which contains this object.\"
	ProtectionGroupId NullableString `json:"protectionGroupId,omitempty"`
	// \"Specifies the protection group name which contains this object.\"
	ProtectionGroupName NullableString `json:"protectionGroupName,omitempty"`
	// Specifies the Source Object information.
	SourceInfo map[string]interface{} `json:"sourceInfo,omitempty"`
	// \"Specifies the Storage Domain id where the backup data of Object is present.\"
	StorageDomainId NullableInt64 `json:"storageDomainId,omitempty"`
	// Specifies snapshot tags applied to the object.
	SnapshotTags []SnapshotTagInfo `json:"snapshotTags,omitempty"`
	// Specifies tag applied to the object.
	Tags []TagInfo `json:"tags,omitempty"`
	ChannelItem *ChannelItem `json:"channelItem,omitempty"`
	FileItem *TeamsFileItem `json:"fileItem,omitempty"`
	// Specifies the M365 Teams item type.
	Type NullableString `json:"type,omitempty"`
}

// NewTeamsItem instantiates a new TeamsItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamsItem() *TeamsItem {
	this := TeamsItem{}
	return &this
}

// NewTeamsItemWithDefaults instantiates a new TeamsItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamsItemWithDefaults() *TeamsItem {
	this := TeamsItem{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TeamsItem) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TeamsItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *TeamsItem) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *TeamsItem) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *TeamsItem) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *TeamsItem) UnsetName() {
	o.Name.Unset()
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TeamsItem) GetPath() string {
	if o == nil || IsNil(o.Path.Get()) {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TeamsItem) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *TeamsItem) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *TeamsItem) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *TeamsItem) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *TeamsItem) UnsetPath() {
	o.Path.Unset()
}

// GetPolicyId returns the PolicyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TeamsItem) GetPolicyId() string {
	if o == nil || IsNil(o.PolicyId.Get()) {
		var ret string
		return ret
	}
	return *o.PolicyId.Get()
}

// GetPolicyIdOk returns a tuple with the PolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TeamsItem) GetPolicyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PolicyId.Get(), o.PolicyId.IsSet()
}

// HasPolicyId returns a boolean if a field has been set.
func (o *TeamsItem) HasPolicyId() bool {
	if o != nil && o.PolicyId.IsSet() {
		return true
	}

	return false
}

// SetPolicyId gets a reference to the given NullableString and assigns it to the PolicyId field.
func (o *TeamsItem) SetPolicyId(v string) {
	o.PolicyId.Set(&v)
}
// SetPolicyIdNil sets the value for PolicyId to be an explicit nil
func (o *TeamsItem) SetPolicyIdNil() {
	o.PolicyId.Set(nil)
}

// UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
func (o *TeamsItem) UnsetPolicyId() {
	o.PolicyId.Unset()
}

// GetPolicyName returns the PolicyName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TeamsItem) GetPolicyName() string {
	if o == nil || IsNil(o.PolicyName.Get()) {
		var ret string
		return ret
	}
	return *o.PolicyName.Get()
}

// GetPolicyNameOk returns a tuple with the PolicyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TeamsItem) GetPolicyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PolicyName.Get(), o.PolicyName.IsSet()
}

// HasPolicyName returns a boolean if a field has been set.
func (o *TeamsItem) HasPolicyName() bool {
	if o != nil && o.PolicyName.IsSet() {
		return true
	}

	return false
}

// SetPolicyName gets a reference to the given NullableString and assigns it to the PolicyName field.
func (o *TeamsItem) SetPolicyName(v string) {
	o.PolicyName.Set(&v)
}
// SetPolicyNameNil sets the value for PolicyName to be an explicit nil
func (o *TeamsItem) SetPolicyNameNil() {
	o.PolicyName.Set(nil)
}

// UnsetPolicyName ensures that no value is present for PolicyName, not even an explicit nil
func (o *TeamsItem) UnsetPolicyName() {
	o.PolicyName.Unset()
}

// GetProtectionGroupId returns the ProtectionGroupId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TeamsItem) GetProtectionGroupId() string {
	if o == nil || IsNil(o.ProtectionGroupId.Get()) {
		var ret string
		return ret
	}
	return *o.ProtectionGroupId.Get()
}

// GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TeamsItem) GetProtectionGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProtectionGroupId.Get(), o.ProtectionGroupId.IsSet()
}

// HasProtectionGroupId returns a boolean if a field has been set.
func (o *TeamsItem) HasProtectionGroupId() bool {
	if o != nil && o.ProtectionGroupId.IsSet() {
		return true
	}

	return false
}

// SetProtectionGroupId gets a reference to the given NullableString and assigns it to the ProtectionGroupId field.
func (o *TeamsItem) SetProtectionGroupId(v string) {
	o.ProtectionGroupId.Set(&v)
}
// SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil
func (o *TeamsItem) SetProtectionGroupIdNil() {
	o.ProtectionGroupId.Set(nil)
}

// UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil
func (o *TeamsItem) UnsetProtectionGroupId() {
	o.ProtectionGroupId.Unset()
}

// GetProtectionGroupName returns the ProtectionGroupName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TeamsItem) GetProtectionGroupName() string {
	if o == nil || IsNil(o.ProtectionGroupName.Get()) {
		var ret string
		return ret
	}
	return *o.ProtectionGroupName.Get()
}

// GetProtectionGroupNameOk returns a tuple with the ProtectionGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TeamsItem) GetProtectionGroupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProtectionGroupName.Get(), o.ProtectionGroupName.IsSet()
}

// HasProtectionGroupName returns a boolean if a field has been set.
func (o *TeamsItem) HasProtectionGroupName() bool {
	if o != nil && o.ProtectionGroupName.IsSet() {
		return true
	}

	return false
}

// SetProtectionGroupName gets a reference to the given NullableString and assigns it to the ProtectionGroupName field.
func (o *TeamsItem) SetProtectionGroupName(v string) {
	o.ProtectionGroupName.Set(&v)
}
// SetProtectionGroupNameNil sets the value for ProtectionGroupName to be an explicit nil
func (o *TeamsItem) SetProtectionGroupNameNil() {
	o.ProtectionGroupName.Set(nil)
}

// UnsetProtectionGroupName ensures that no value is present for ProtectionGroupName, not even an explicit nil
func (o *TeamsItem) UnsetProtectionGroupName() {
	o.ProtectionGroupName.Unset()
}

// GetSourceInfo returns the SourceInfo field value if set, zero value otherwise.
func (o *TeamsItem) GetSourceInfo() map[string]interface{} {
	if o == nil || IsNil(o.SourceInfo) {
		var ret map[string]interface{}
		return ret
	}
	return o.SourceInfo
}

// GetSourceInfoOk returns a tuple with the SourceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsItem) GetSourceInfoOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.SourceInfo) {
		return map[string]interface{}{}, false
	}
	return o.SourceInfo, true
}

// HasSourceInfo returns a boolean if a field has been set.
func (o *TeamsItem) HasSourceInfo() bool {
	if o != nil && !IsNil(o.SourceInfo) {
		return true
	}

	return false
}

// SetSourceInfo gets a reference to the given map[string]interface{} and assigns it to the SourceInfo field.
func (o *TeamsItem) SetSourceInfo(v map[string]interface{}) {
	o.SourceInfo = v
}

// GetStorageDomainId returns the StorageDomainId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TeamsItem) GetStorageDomainId() int64 {
	if o == nil || IsNil(o.StorageDomainId.Get()) {
		var ret int64
		return ret
	}
	return *o.StorageDomainId.Get()
}

// GetStorageDomainIdOk returns a tuple with the StorageDomainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TeamsItem) GetStorageDomainIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.StorageDomainId.Get(), o.StorageDomainId.IsSet()
}

// HasStorageDomainId returns a boolean if a field has been set.
func (o *TeamsItem) HasStorageDomainId() bool {
	if o != nil && o.StorageDomainId.IsSet() {
		return true
	}

	return false
}

// SetStorageDomainId gets a reference to the given NullableInt64 and assigns it to the StorageDomainId field.
func (o *TeamsItem) SetStorageDomainId(v int64) {
	o.StorageDomainId.Set(&v)
}
// SetStorageDomainIdNil sets the value for StorageDomainId to be an explicit nil
func (o *TeamsItem) SetStorageDomainIdNil() {
	o.StorageDomainId.Set(nil)
}

// UnsetStorageDomainId ensures that no value is present for StorageDomainId, not even an explicit nil
func (o *TeamsItem) UnsetStorageDomainId() {
	o.StorageDomainId.Unset()
}

// GetSnapshotTags returns the SnapshotTags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TeamsItem) GetSnapshotTags() []SnapshotTagInfo {
	if o == nil {
		var ret []SnapshotTagInfo
		return ret
	}
	return o.SnapshotTags
}

// GetSnapshotTagsOk returns a tuple with the SnapshotTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TeamsItem) GetSnapshotTagsOk() ([]SnapshotTagInfo, bool) {
	if o == nil || IsNil(o.SnapshotTags) {
		return nil, false
	}
	return o.SnapshotTags, true
}

// HasSnapshotTags returns a boolean if a field has been set.
func (o *TeamsItem) HasSnapshotTags() bool {
	if o != nil && !IsNil(o.SnapshotTags) {
		return true
	}

	return false
}

// SetSnapshotTags gets a reference to the given []SnapshotTagInfo and assigns it to the SnapshotTags field.
func (o *TeamsItem) SetSnapshotTags(v []SnapshotTagInfo) {
	o.SnapshotTags = v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TeamsItem) GetTags() []TagInfo {
	if o == nil {
		var ret []TagInfo
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TeamsItem) GetTagsOk() ([]TagInfo, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *TeamsItem) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagInfo and assigns it to the Tags field.
func (o *TeamsItem) SetTags(v []TagInfo) {
	o.Tags = v
}

// GetChannelItem returns the ChannelItem field value if set, zero value otherwise.
func (o *TeamsItem) GetChannelItem() ChannelItem {
	if o == nil || IsNil(o.ChannelItem) {
		var ret ChannelItem
		return ret
	}
	return *o.ChannelItem
}

// GetChannelItemOk returns a tuple with the ChannelItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsItem) GetChannelItemOk() (*ChannelItem, bool) {
	if o == nil || IsNil(o.ChannelItem) {
		return nil, false
	}
	return o.ChannelItem, true
}

// HasChannelItem returns a boolean if a field has been set.
func (o *TeamsItem) HasChannelItem() bool {
	if o != nil && !IsNil(o.ChannelItem) {
		return true
	}

	return false
}

// SetChannelItem gets a reference to the given ChannelItem and assigns it to the ChannelItem field.
func (o *TeamsItem) SetChannelItem(v ChannelItem) {
	o.ChannelItem = &v
}

// GetFileItem returns the FileItem field value if set, zero value otherwise.
func (o *TeamsItem) GetFileItem() TeamsFileItem {
	if o == nil || IsNil(o.FileItem) {
		var ret TeamsFileItem
		return ret
	}
	return *o.FileItem
}

// GetFileItemOk returns a tuple with the FileItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsItem) GetFileItemOk() (*TeamsFileItem, bool) {
	if o == nil || IsNil(o.FileItem) {
		return nil, false
	}
	return o.FileItem, true
}

// HasFileItem returns a boolean if a field has been set.
func (o *TeamsItem) HasFileItem() bool {
	if o != nil && !IsNil(o.FileItem) {
		return true
	}

	return false
}

// SetFileItem gets a reference to the given TeamsFileItem and assigns it to the FileItem field.
func (o *TeamsItem) SetFileItem(v TeamsFileItem) {
	o.FileItem = &v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TeamsItem) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TeamsItem) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *TeamsItem) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *TeamsItem) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *TeamsItem) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *TeamsItem) UnsetType() {
	o.Type.Unset()
}

func (o TeamsItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TeamsItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Path.IsSet() {
		toSerialize["path"] = o.Path.Get()
	}
	if o.PolicyId.IsSet() {
		toSerialize["policyId"] = o.PolicyId.Get()
	}
	if o.PolicyName.IsSet() {
		toSerialize["policyName"] = o.PolicyName.Get()
	}
	if o.ProtectionGroupId.IsSet() {
		toSerialize["protectionGroupId"] = o.ProtectionGroupId.Get()
	}
	if o.ProtectionGroupName.IsSet() {
		toSerialize["protectionGroupName"] = o.ProtectionGroupName.Get()
	}
	if !IsNil(o.SourceInfo) {
		toSerialize["sourceInfo"] = o.SourceInfo
	}
	if o.StorageDomainId.IsSet() {
		toSerialize["storageDomainId"] = o.StorageDomainId.Get()
	}
	if o.SnapshotTags != nil {
		toSerialize["snapshotTags"] = o.SnapshotTags
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.ChannelItem) {
		toSerialize["channelItem"] = o.ChannelItem
	}
	if !IsNil(o.FileItem) {
		toSerialize["fileItem"] = o.FileItem
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	return toSerialize, nil
}

type NullableTeamsItem struct {
	value *TeamsItem
	isSet bool
}

func (v NullableTeamsItem) Get() *TeamsItem {
	return v.value
}

func (v *NullableTeamsItem) Set(val *TeamsItem) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamsItem) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamsItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamsItem(val *TeamsItem) *NullableTeamsItem {
	return &NullableTeamsItem{value: val, isSet: true}
}

func (v NullableTeamsItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamsItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


