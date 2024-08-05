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

// checks if the DocumentLibraryItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentLibraryItem{}

// DocumentLibraryItem Specifies a Document Library indexed item.
type DocumentLibraryItem struct {
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
	// Specifies the Unix timestamp epoch in seconds at which this item is created.
	CreationTimeSecs NullableInt64 `json:"creationTimeSecs,omitempty"`
	// Specifies the file type.
	FileType NullableString `json:"fileType,omitempty"`
	// Specifies the id of the document library item.
	ItemId NullableString `json:"itemId,omitempty"`
	// Specifies the size in bytes for the indexed item.
	ItemSize NullableInt64 `json:"itemSize,omitempty"`
	// Specifies the email of the owner of the document library item.
	OwnerEmail NullableString `json:"ownerEmail,omitempty"`
	// Specifies the name of the owner of the document library item.
	OwnerName NullableString `json:"ownerName,omitempty"`
}

// NewDocumentLibraryItem instantiates a new DocumentLibraryItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentLibraryItem() *DocumentLibraryItem {
	this := DocumentLibraryItem{}
	return &this
}

// NewDocumentLibraryItemWithDefaults instantiates a new DocumentLibraryItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentLibraryItemWithDefaults() *DocumentLibraryItem {
	this := DocumentLibraryItem{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentLibraryItem) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentLibraryItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *DocumentLibraryItem) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *DocumentLibraryItem) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *DocumentLibraryItem) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *DocumentLibraryItem) UnsetName() {
	o.Name.Unset()
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentLibraryItem) GetPath() string {
	if o == nil || IsNil(o.Path.Get()) {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentLibraryItem) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *DocumentLibraryItem) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *DocumentLibraryItem) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *DocumentLibraryItem) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *DocumentLibraryItem) UnsetPath() {
	o.Path.Unset()
}

// GetPolicyId returns the PolicyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentLibraryItem) GetPolicyId() string {
	if o == nil || IsNil(o.PolicyId.Get()) {
		var ret string
		return ret
	}
	return *o.PolicyId.Get()
}

// GetPolicyIdOk returns a tuple with the PolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentLibraryItem) GetPolicyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PolicyId.Get(), o.PolicyId.IsSet()
}

// HasPolicyId returns a boolean if a field has been set.
func (o *DocumentLibraryItem) HasPolicyId() bool {
	if o != nil && o.PolicyId.IsSet() {
		return true
	}

	return false
}

// SetPolicyId gets a reference to the given NullableString and assigns it to the PolicyId field.
func (o *DocumentLibraryItem) SetPolicyId(v string) {
	o.PolicyId.Set(&v)
}
// SetPolicyIdNil sets the value for PolicyId to be an explicit nil
func (o *DocumentLibraryItem) SetPolicyIdNil() {
	o.PolicyId.Set(nil)
}

// UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
func (o *DocumentLibraryItem) UnsetPolicyId() {
	o.PolicyId.Unset()
}

// GetPolicyName returns the PolicyName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentLibraryItem) GetPolicyName() string {
	if o == nil || IsNil(o.PolicyName.Get()) {
		var ret string
		return ret
	}
	return *o.PolicyName.Get()
}

// GetPolicyNameOk returns a tuple with the PolicyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentLibraryItem) GetPolicyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PolicyName.Get(), o.PolicyName.IsSet()
}

// HasPolicyName returns a boolean if a field has been set.
func (o *DocumentLibraryItem) HasPolicyName() bool {
	if o != nil && o.PolicyName.IsSet() {
		return true
	}

	return false
}

// SetPolicyName gets a reference to the given NullableString and assigns it to the PolicyName field.
func (o *DocumentLibraryItem) SetPolicyName(v string) {
	o.PolicyName.Set(&v)
}
// SetPolicyNameNil sets the value for PolicyName to be an explicit nil
func (o *DocumentLibraryItem) SetPolicyNameNil() {
	o.PolicyName.Set(nil)
}

// UnsetPolicyName ensures that no value is present for PolicyName, not even an explicit nil
func (o *DocumentLibraryItem) UnsetPolicyName() {
	o.PolicyName.Unset()
}

// GetProtectionGroupId returns the ProtectionGroupId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentLibraryItem) GetProtectionGroupId() string {
	if o == nil || IsNil(o.ProtectionGroupId.Get()) {
		var ret string
		return ret
	}
	return *o.ProtectionGroupId.Get()
}

// GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentLibraryItem) GetProtectionGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProtectionGroupId.Get(), o.ProtectionGroupId.IsSet()
}

// HasProtectionGroupId returns a boolean if a field has been set.
func (o *DocumentLibraryItem) HasProtectionGroupId() bool {
	if o != nil && o.ProtectionGroupId.IsSet() {
		return true
	}

	return false
}

// SetProtectionGroupId gets a reference to the given NullableString and assigns it to the ProtectionGroupId field.
func (o *DocumentLibraryItem) SetProtectionGroupId(v string) {
	o.ProtectionGroupId.Set(&v)
}
// SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil
func (o *DocumentLibraryItem) SetProtectionGroupIdNil() {
	o.ProtectionGroupId.Set(nil)
}

// UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil
func (o *DocumentLibraryItem) UnsetProtectionGroupId() {
	o.ProtectionGroupId.Unset()
}

// GetProtectionGroupName returns the ProtectionGroupName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentLibraryItem) GetProtectionGroupName() string {
	if o == nil || IsNil(o.ProtectionGroupName.Get()) {
		var ret string
		return ret
	}
	return *o.ProtectionGroupName.Get()
}

// GetProtectionGroupNameOk returns a tuple with the ProtectionGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentLibraryItem) GetProtectionGroupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProtectionGroupName.Get(), o.ProtectionGroupName.IsSet()
}

// HasProtectionGroupName returns a boolean if a field has been set.
func (o *DocumentLibraryItem) HasProtectionGroupName() bool {
	if o != nil && o.ProtectionGroupName.IsSet() {
		return true
	}

	return false
}

// SetProtectionGroupName gets a reference to the given NullableString and assigns it to the ProtectionGroupName field.
func (o *DocumentLibraryItem) SetProtectionGroupName(v string) {
	o.ProtectionGroupName.Set(&v)
}
// SetProtectionGroupNameNil sets the value for ProtectionGroupName to be an explicit nil
func (o *DocumentLibraryItem) SetProtectionGroupNameNil() {
	o.ProtectionGroupName.Set(nil)
}

// UnsetProtectionGroupName ensures that no value is present for ProtectionGroupName, not even an explicit nil
func (o *DocumentLibraryItem) UnsetProtectionGroupName() {
	o.ProtectionGroupName.Unset()
}

// GetSourceInfo returns the SourceInfo field value if set, zero value otherwise.
func (o *DocumentLibraryItem) GetSourceInfo() map[string]interface{} {
	if o == nil || IsNil(o.SourceInfo) {
		var ret map[string]interface{}
		return ret
	}
	return o.SourceInfo
}

// GetSourceInfoOk returns a tuple with the SourceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentLibraryItem) GetSourceInfoOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.SourceInfo) {
		return map[string]interface{}{}, false
	}
	return o.SourceInfo, true
}

// HasSourceInfo returns a boolean if a field has been set.
func (o *DocumentLibraryItem) HasSourceInfo() bool {
	if o != nil && !IsNil(o.SourceInfo) {
		return true
	}

	return false
}

// SetSourceInfo gets a reference to the given map[string]interface{} and assigns it to the SourceInfo field.
func (o *DocumentLibraryItem) SetSourceInfo(v map[string]interface{}) {
	o.SourceInfo = v
}

// GetStorageDomainId returns the StorageDomainId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentLibraryItem) GetStorageDomainId() int64 {
	if o == nil || IsNil(o.StorageDomainId.Get()) {
		var ret int64
		return ret
	}
	return *o.StorageDomainId.Get()
}

// GetStorageDomainIdOk returns a tuple with the StorageDomainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentLibraryItem) GetStorageDomainIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.StorageDomainId.Get(), o.StorageDomainId.IsSet()
}

// HasStorageDomainId returns a boolean if a field has been set.
func (o *DocumentLibraryItem) HasStorageDomainId() bool {
	if o != nil && o.StorageDomainId.IsSet() {
		return true
	}

	return false
}

// SetStorageDomainId gets a reference to the given NullableInt64 and assigns it to the StorageDomainId field.
func (o *DocumentLibraryItem) SetStorageDomainId(v int64) {
	o.StorageDomainId.Set(&v)
}
// SetStorageDomainIdNil sets the value for StorageDomainId to be an explicit nil
func (o *DocumentLibraryItem) SetStorageDomainIdNil() {
	o.StorageDomainId.Set(nil)
}

// UnsetStorageDomainId ensures that no value is present for StorageDomainId, not even an explicit nil
func (o *DocumentLibraryItem) UnsetStorageDomainId() {
	o.StorageDomainId.Unset()
}

// GetSnapshotTags returns the SnapshotTags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentLibraryItem) GetSnapshotTags() []SnapshotTagInfo {
	if o == nil {
		var ret []SnapshotTagInfo
		return ret
	}
	return o.SnapshotTags
}

// GetSnapshotTagsOk returns a tuple with the SnapshotTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentLibraryItem) GetSnapshotTagsOk() ([]SnapshotTagInfo, bool) {
	if o == nil || IsNil(o.SnapshotTags) {
		return nil, false
	}
	return o.SnapshotTags, true
}

// HasSnapshotTags returns a boolean if a field has been set.
func (o *DocumentLibraryItem) HasSnapshotTags() bool {
	if o != nil && !IsNil(o.SnapshotTags) {
		return true
	}

	return false
}

// SetSnapshotTags gets a reference to the given []SnapshotTagInfo and assigns it to the SnapshotTags field.
func (o *DocumentLibraryItem) SetSnapshotTags(v []SnapshotTagInfo) {
	o.SnapshotTags = v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentLibraryItem) GetTags() []TagInfo {
	if o == nil {
		var ret []TagInfo
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentLibraryItem) GetTagsOk() ([]TagInfo, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DocumentLibraryItem) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagInfo and assigns it to the Tags field.
func (o *DocumentLibraryItem) SetTags(v []TagInfo) {
	o.Tags = v
}

// GetCreationTimeSecs returns the CreationTimeSecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentLibraryItem) GetCreationTimeSecs() int64 {
	if o == nil || IsNil(o.CreationTimeSecs.Get()) {
		var ret int64
		return ret
	}
	return *o.CreationTimeSecs.Get()
}

// GetCreationTimeSecsOk returns a tuple with the CreationTimeSecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentLibraryItem) GetCreationTimeSecsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreationTimeSecs.Get(), o.CreationTimeSecs.IsSet()
}

// HasCreationTimeSecs returns a boolean if a field has been set.
func (o *DocumentLibraryItem) HasCreationTimeSecs() bool {
	if o != nil && o.CreationTimeSecs.IsSet() {
		return true
	}

	return false
}

// SetCreationTimeSecs gets a reference to the given NullableInt64 and assigns it to the CreationTimeSecs field.
func (o *DocumentLibraryItem) SetCreationTimeSecs(v int64) {
	o.CreationTimeSecs.Set(&v)
}
// SetCreationTimeSecsNil sets the value for CreationTimeSecs to be an explicit nil
func (o *DocumentLibraryItem) SetCreationTimeSecsNil() {
	o.CreationTimeSecs.Set(nil)
}

// UnsetCreationTimeSecs ensures that no value is present for CreationTimeSecs, not even an explicit nil
func (o *DocumentLibraryItem) UnsetCreationTimeSecs() {
	o.CreationTimeSecs.Unset()
}

// GetFileType returns the FileType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentLibraryItem) GetFileType() string {
	if o == nil || IsNil(o.FileType.Get()) {
		var ret string
		return ret
	}
	return *o.FileType.Get()
}

// GetFileTypeOk returns a tuple with the FileType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentLibraryItem) GetFileTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileType.Get(), o.FileType.IsSet()
}

// HasFileType returns a boolean if a field has been set.
func (o *DocumentLibraryItem) HasFileType() bool {
	if o != nil && o.FileType.IsSet() {
		return true
	}

	return false
}

// SetFileType gets a reference to the given NullableString and assigns it to the FileType field.
func (o *DocumentLibraryItem) SetFileType(v string) {
	o.FileType.Set(&v)
}
// SetFileTypeNil sets the value for FileType to be an explicit nil
func (o *DocumentLibraryItem) SetFileTypeNil() {
	o.FileType.Set(nil)
}

// UnsetFileType ensures that no value is present for FileType, not even an explicit nil
func (o *DocumentLibraryItem) UnsetFileType() {
	o.FileType.Unset()
}

// GetItemId returns the ItemId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentLibraryItem) GetItemId() string {
	if o == nil || IsNil(o.ItemId.Get()) {
		var ret string
		return ret
	}
	return *o.ItemId.Get()
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentLibraryItem) GetItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ItemId.Get(), o.ItemId.IsSet()
}

// HasItemId returns a boolean if a field has been set.
func (o *DocumentLibraryItem) HasItemId() bool {
	if o != nil && o.ItemId.IsSet() {
		return true
	}

	return false
}

// SetItemId gets a reference to the given NullableString and assigns it to the ItemId field.
func (o *DocumentLibraryItem) SetItemId(v string) {
	o.ItemId.Set(&v)
}
// SetItemIdNil sets the value for ItemId to be an explicit nil
func (o *DocumentLibraryItem) SetItemIdNil() {
	o.ItemId.Set(nil)
}

// UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
func (o *DocumentLibraryItem) UnsetItemId() {
	o.ItemId.Unset()
}

// GetItemSize returns the ItemSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentLibraryItem) GetItemSize() int64 {
	if o == nil || IsNil(o.ItemSize.Get()) {
		var ret int64
		return ret
	}
	return *o.ItemSize.Get()
}

// GetItemSizeOk returns a tuple with the ItemSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentLibraryItem) GetItemSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ItemSize.Get(), o.ItemSize.IsSet()
}

// HasItemSize returns a boolean if a field has been set.
func (o *DocumentLibraryItem) HasItemSize() bool {
	if o != nil && o.ItemSize.IsSet() {
		return true
	}

	return false
}

// SetItemSize gets a reference to the given NullableInt64 and assigns it to the ItemSize field.
func (o *DocumentLibraryItem) SetItemSize(v int64) {
	o.ItemSize.Set(&v)
}
// SetItemSizeNil sets the value for ItemSize to be an explicit nil
func (o *DocumentLibraryItem) SetItemSizeNil() {
	o.ItemSize.Set(nil)
}

// UnsetItemSize ensures that no value is present for ItemSize, not even an explicit nil
func (o *DocumentLibraryItem) UnsetItemSize() {
	o.ItemSize.Unset()
}

// GetOwnerEmail returns the OwnerEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentLibraryItem) GetOwnerEmail() string {
	if o == nil || IsNil(o.OwnerEmail.Get()) {
		var ret string
		return ret
	}
	return *o.OwnerEmail.Get()
}

// GetOwnerEmailOk returns a tuple with the OwnerEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentLibraryItem) GetOwnerEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OwnerEmail.Get(), o.OwnerEmail.IsSet()
}

// HasOwnerEmail returns a boolean if a field has been set.
func (o *DocumentLibraryItem) HasOwnerEmail() bool {
	if o != nil && o.OwnerEmail.IsSet() {
		return true
	}

	return false
}

// SetOwnerEmail gets a reference to the given NullableString and assigns it to the OwnerEmail field.
func (o *DocumentLibraryItem) SetOwnerEmail(v string) {
	o.OwnerEmail.Set(&v)
}
// SetOwnerEmailNil sets the value for OwnerEmail to be an explicit nil
func (o *DocumentLibraryItem) SetOwnerEmailNil() {
	o.OwnerEmail.Set(nil)
}

// UnsetOwnerEmail ensures that no value is present for OwnerEmail, not even an explicit nil
func (o *DocumentLibraryItem) UnsetOwnerEmail() {
	o.OwnerEmail.Unset()
}

// GetOwnerName returns the OwnerName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentLibraryItem) GetOwnerName() string {
	if o == nil || IsNil(o.OwnerName.Get()) {
		var ret string
		return ret
	}
	return *o.OwnerName.Get()
}

// GetOwnerNameOk returns a tuple with the OwnerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentLibraryItem) GetOwnerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OwnerName.Get(), o.OwnerName.IsSet()
}

// HasOwnerName returns a boolean if a field has been set.
func (o *DocumentLibraryItem) HasOwnerName() bool {
	if o != nil && o.OwnerName.IsSet() {
		return true
	}

	return false
}

// SetOwnerName gets a reference to the given NullableString and assigns it to the OwnerName field.
func (o *DocumentLibraryItem) SetOwnerName(v string) {
	o.OwnerName.Set(&v)
}
// SetOwnerNameNil sets the value for OwnerName to be an explicit nil
func (o *DocumentLibraryItem) SetOwnerNameNil() {
	o.OwnerName.Set(nil)
}

// UnsetOwnerName ensures that no value is present for OwnerName, not even an explicit nil
func (o *DocumentLibraryItem) UnsetOwnerName() {
	o.OwnerName.Unset()
}

func (o DocumentLibraryItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentLibraryItem) ToMap() (map[string]interface{}, error) {
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
	if o.CreationTimeSecs.IsSet() {
		toSerialize["creationTimeSecs"] = o.CreationTimeSecs.Get()
	}
	if o.FileType.IsSet() {
		toSerialize["fileType"] = o.FileType.Get()
	}
	if o.ItemId.IsSet() {
		toSerialize["itemId"] = o.ItemId.Get()
	}
	if o.ItemSize.IsSet() {
		toSerialize["itemSize"] = o.ItemSize.Get()
	}
	if o.OwnerEmail.IsSet() {
		toSerialize["ownerEmail"] = o.OwnerEmail.Get()
	}
	if o.OwnerName.IsSet() {
		toSerialize["ownerName"] = o.OwnerName.Get()
	}
	return toSerialize, nil
}

type NullableDocumentLibraryItem struct {
	value *DocumentLibraryItem
	isSet bool
}

func (v NullableDocumentLibraryItem) Get() *DocumentLibraryItem {
	return v.value
}

func (v *NullableDocumentLibraryItem) Set(val *DocumentLibraryItem) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentLibraryItem) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentLibraryItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentLibraryItem(val *DocumentLibraryItem) *NullableDocumentLibraryItem {
	return &NullableDocumentLibraryItem{value: val, isSet: true}
}

func (v NullableDocumentLibraryItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentLibraryItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


