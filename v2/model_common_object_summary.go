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

// checks if the CommonObjectSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonObjectSummary{}

// CommonObjectSummary Specifies an Common summary properties for an object.
type CommonObjectSummary struct {
	// Specifies the environment of the object.
	Environment NullableString `json:"environment,omitempty"`
	// Specifies object id.
	Id NullableInt64 `json:"id,omitempty"`
	// Specifies the name of the object.
	Name NullableString `json:"name,omitempty"`
	// Specifies registered source id to which object belongs.
	SourceId NullableInt64 `json:"sourceId,omitempty"`
	// Specifies registered source name to which object belongs.
	SourceName NullableString `json:"sourceName,omitempty"`
	// Specifies child object details.
	ChildObjects []ObjectSummary `json:"childObjects,omitempty"`
	// Specifies the global id which is a unique identifier of the object.
	GlobalId NullableString `json:"globalId,omitempty"`
	// Specifies the logical size of object in bytes.
	LogicalSizeBytes NullableInt64 `json:"logicalSizeBytes,omitempty"`
	// Specifies the hash identifier of the object.
	ObjectHash NullableString `json:"objectHash,omitempty"`
	// Specifies the type of the object.
	ObjectType NullableString `json:"objectType,omitempty"`
	// Specifies the operating system type of the object.
	OsType NullableString `json:"osType,omitempty"`
	// Specifies the protection type of the object if any.
	ProtectionType NullableString `json:"protectionType,omitempty"`
	SharepointSiteSummary *SharepointObjectParams `json:"sharepointSiteSummary,omitempty"`
	// Specifies the uuid which is a unique identifier of the object.
	Uuid NullableString `json:"uuid,omitempty"`
	VCenterSummary *ObjectTypeVCenterParams `json:"vCenterSummary,omitempty"`
	WindowsClusterSummary *ObjectTypeWindowsClusterParams `json:"windowsClusterSummary,omitempty"`
	Permissions *PermissionInfo `json:"permissions,omitempty"`
	// Specifies the count and size of protected and unprotected objects for the size.
	ProtectionStats []ObjectProtectionStatsSummary `json:"protectionStats,omitempty"`
}

// NewCommonObjectSummary instantiates a new CommonObjectSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonObjectSummary() *CommonObjectSummary {
	this := CommonObjectSummary{}
	return &this
}

// NewCommonObjectSummaryWithDefaults instantiates a new CommonObjectSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonObjectSummaryWithDefaults() *CommonObjectSummary {
	this := CommonObjectSummary{}
	return &this
}

// GetEnvironment returns the Environment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonObjectSummary) GetEnvironment() string {
	if o == nil || IsNil(o.Environment.Get()) {
		var ret string
		return ret
	}
	return *o.Environment.Get()
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonObjectSummary) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Environment.Get(), o.Environment.IsSet()
}

// HasEnvironment returns a boolean if a field has been set.
func (o *CommonObjectSummary) HasEnvironment() bool {
	if o != nil && o.Environment.IsSet() {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given NullableString and assigns it to the Environment field.
func (o *CommonObjectSummary) SetEnvironment(v string) {
	o.Environment.Set(&v)
}
// SetEnvironmentNil sets the value for Environment to be an explicit nil
func (o *CommonObjectSummary) SetEnvironmentNil() {
	o.Environment.Set(nil)
}

// UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
func (o *CommonObjectSummary) UnsetEnvironment() {
	o.Environment.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonObjectSummary) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonObjectSummary) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *CommonObjectSummary) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *CommonObjectSummary) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *CommonObjectSummary) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *CommonObjectSummary) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonObjectSummary) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonObjectSummary) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CommonObjectSummary) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CommonObjectSummary) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *CommonObjectSummary) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CommonObjectSummary) UnsetName() {
	o.Name.Unset()
}

// GetSourceId returns the SourceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonObjectSummary) GetSourceId() int64 {
	if o == nil || IsNil(o.SourceId.Get()) {
		var ret int64
		return ret
	}
	return *o.SourceId.Get()
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonObjectSummary) GetSourceIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceId.Get(), o.SourceId.IsSet()
}

// HasSourceId returns a boolean if a field has been set.
func (o *CommonObjectSummary) HasSourceId() bool {
	if o != nil && o.SourceId.IsSet() {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given NullableInt64 and assigns it to the SourceId field.
func (o *CommonObjectSummary) SetSourceId(v int64) {
	o.SourceId.Set(&v)
}
// SetSourceIdNil sets the value for SourceId to be an explicit nil
func (o *CommonObjectSummary) SetSourceIdNil() {
	o.SourceId.Set(nil)
}

// UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
func (o *CommonObjectSummary) UnsetSourceId() {
	o.SourceId.Unset()
}

// GetSourceName returns the SourceName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonObjectSummary) GetSourceName() string {
	if o == nil || IsNil(o.SourceName.Get()) {
		var ret string
		return ret
	}
	return *o.SourceName.Get()
}

// GetSourceNameOk returns a tuple with the SourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonObjectSummary) GetSourceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceName.Get(), o.SourceName.IsSet()
}

// HasSourceName returns a boolean if a field has been set.
func (o *CommonObjectSummary) HasSourceName() bool {
	if o != nil && o.SourceName.IsSet() {
		return true
	}

	return false
}

// SetSourceName gets a reference to the given NullableString and assigns it to the SourceName field.
func (o *CommonObjectSummary) SetSourceName(v string) {
	o.SourceName.Set(&v)
}
// SetSourceNameNil sets the value for SourceName to be an explicit nil
func (o *CommonObjectSummary) SetSourceNameNil() {
	o.SourceName.Set(nil)
}

// UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil
func (o *CommonObjectSummary) UnsetSourceName() {
	o.SourceName.Unset()
}

// GetChildObjects returns the ChildObjects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonObjectSummary) GetChildObjects() []ObjectSummary {
	if o == nil {
		var ret []ObjectSummary
		return ret
	}
	return o.ChildObjects
}

// GetChildObjectsOk returns a tuple with the ChildObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonObjectSummary) GetChildObjectsOk() ([]ObjectSummary, bool) {
	if o == nil || IsNil(o.ChildObjects) {
		return nil, false
	}
	return o.ChildObjects, true
}

// HasChildObjects returns a boolean if a field has been set.
func (o *CommonObjectSummary) HasChildObjects() bool {
	if o != nil && !IsNil(o.ChildObjects) {
		return true
	}

	return false
}

// SetChildObjects gets a reference to the given []ObjectSummary and assigns it to the ChildObjects field.
func (o *CommonObjectSummary) SetChildObjects(v []ObjectSummary) {
	o.ChildObjects = v
}

// GetGlobalId returns the GlobalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonObjectSummary) GetGlobalId() string {
	if o == nil || IsNil(o.GlobalId.Get()) {
		var ret string
		return ret
	}
	return *o.GlobalId.Get()
}

// GetGlobalIdOk returns a tuple with the GlobalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonObjectSummary) GetGlobalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GlobalId.Get(), o.GlobalId.IsSet()
}

// HasGlobalId returns a boolean if a field has been set.
func (o *CommonObjectSummary) HasGlobalId() bool {
	if o != nil && o.GlobalId.IsSet() {
		return true
	}

	return false
}

// SetGlobalId gets a reference to the given NullableString and assigns it to the GlobalId field.
func (o *CommonObjectSummary) SetGlobalId(v string) {
	o.GlobalId.Set(&v)
}
// SetGlobalIdNil sets the value for GlobalId to be an explicit nil
func (o *CommonObjectSummary) SetGlobalIdNil() {
	o.GlobalId.Set(nil)
}

// UnsetGlobalId ensures that no value is present for GlobalId, not even an explicit nil
func (o *CommonObjectSummary) UnsetGlobalId() {
	o.GlobalId.Unset()
}

// GetLogicalSizeBytes returns the LogicalSizeBytes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonObjectSummary) GetLogicalSizeBytes() int64 {
	if o == nil || IsNil(o.LogicalSizeBytes.Get()) {
		var ret int64
		return ret
	}
	return *o.LogicalSizeBytes.Get()
}

// GetLogicalSizeBytesOk returns a tuple with the LogicalSizeBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonObjectSummary) GetLogicalSizeBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogicalSizeBytes.Get(), o.LogicalSizeBytes.IsSet()
}

// HasLogicalSizeBytes returns a boolean if a field has been set.
func (o *CommonObjectSummary) HasLogicalSizeBytes() bool {
	if o != nil && o.LogicalSizeBytes.IsSet() {
		return true
	}

	return false
}

// SetLogicalSizeBytes gets a reference to the given NullableInt64 and assigns it to the LogicalSizeBytes field.
func (o *CommonObjectSummary) SetLogicalSizeBytes(v int64) {
	o.LogicalSizeBytes.Set(&v)
}
// SetLogicalSizeBytesNil sets the value for LogicalSizeBytes to be an explicit nil
func (o *CommonObjectSummary) SetLogicalSizeBytesNil() {
	o.LogicalSizeBytes.Set(nil)
}

// UnsetLogicalSizeBytes ensures that no value is present for LogicalSizeBytes, not even an explicit nil
func (o *CommonObjectSummary) UnsetLogicalSizeBytes() {
	o.LogicalSizeBytes.Unset()
}

// GetObjectHash returns the ObjectHash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonObjectSummary) GetObjectHash() string {
	if o == nil || IsNil(o.ObjectHash.Get()) {
		var ret string
		return ret
	}
	return *o.ObjectHash.Get()
}

// GetObjectHashOk returns a tuple with the ObjectHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonObjectSummary) GetObjectHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ObjectHash.Get(), o.ObjectHash.IsSet()
}

// HasObjectHash returns a boolean if a field has been set.
func (o *CommonObjectSummary) HasObjectHash() bool {
	if o != nil && o.ObjectHash.IsSet() {
		return true
	}

	return false
}

// SetObjectHash gets a reference to the given NullableString and assigns it to the ObjectHash field.
func (o *CommonObjectSummary) SetObjectHash(v string) {
	o.ObjectHash.Set(&v)
}
// SetObjectHashNil sets the value for ObjectHash to be an explicit nil
func (o *CommonObjectSummary) SetObjectHashNil() {
	o.ObjectHash.Set(nil)
}

// UnsetObjectHash ensures that no value is present for ObjectHash, not even an explicit nil
func (o *CommonObjectSummary) UnsetObjectHash() {
	o.ObjectHash.Unset()
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonObjectSummary) GetObjectType() string {
	if o == nil || IsNil(o.ObjectType.Get()) {
		var ret string
		return ret
	}
	return *o.ObjectType.Get()
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonObjectSummary) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ObjectType.Get(), o.ObjectType.IsSet()
}

// HasObjectType returns a boolean if a field has been set.
func (o *CommonObjectSummary) HasObjectType() bool {
	if o != nil && o.ObjectType.IsSet() {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given NullableString and assigns it to the ObjectType field.
func (o *CommonObjectSummary) SetObjectType(v string) {
	o.ObjectType.Set(&v)
}
// SetObjectTypeNil sets the value for ObjectType to be an explicit nil
func (o *CommonObjectSummary) SetObjectTypeNil() {
	o.ObjectType.Set(nil)
}

// UnsetObjectType ensures that no value is present for ObjectType, not even an explicit nil
func (o *CommonObjectSummary) UnsetObjectType() {
	o.ObjectType.Unset()
}

// GetOsType returns the OsType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonObjectSummary) GetOsType() string {
	if o == nil || IsNil(o.OsType.Get()) {
		var ret string
		return ret
	}
	return *o.OsType.Get()
}

// GetOsTypeOk returns a tuple with the OsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonObjectSummary) GetOsTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OsType.Get(), o.OsType.IsSet()
}

// HasOsType returns a boolean if a field has been set.
func (o *CommonObjectSummary) HasOsType() bool {
	if o != nil && o.OsType.IsSet() {
		return true
	}

	return false
}

// SetOsType gets a reference to the given NullableString and assigns it to the OsType field.
func (o *CommonObjectSummary) SetOsType(v string) {
	o.OsType.Set(&v)
}
// SetOsTypeNil sets the value for OsType to be an explicit nil
func (o *CommonObjectSummary) SetOsTypeNil() {
	o.OsType.Set(nil)
}

// UnsetOsType ensures that no value is present for OsType, not even an explicit nil
func (o *CommonObjectSummary) UnsetOsType() {
	o.OsType.Unset()
}

// GetProtectionType returns the ProtectionType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonObjectSummary) GetProtectionType() string {
	if o == nil || IsNil(o.ProtectionType.Get()) {
		var ret string
		return ret
	}
	return *o.ProtectionType.Get()
}

// GetProtectionTypeOk returns a tuple with the ProtectionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonObjectSummary) GetProtectionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProtectionType.Get(), o.ProtectionType.IsSet()
}

// HasProtectionType returns a boolean if a field has been set.
func (o *CommonObjectSummary) HasProtectionType() bool {
	if o != nil && o.ProtectionType.IsSet() {
		return true
	}

	return false
}

// SetProtectionType gets a reference to the given NullableString and assigns it to the ProtectionType field.
func (o *CommonObjectSummary) SetProtectionType(v string) {
	o.ProtectionType.Set(&v)
}
// SetProtectionTypeNil sets the value for ProtectionType to be an explicit nil
func (o *CommonObjectSummary) SetProtectionTypeNil() {
	o.ProtectionType.Set(nil)
}

// UnsetProtectionType ensures that no value is present for ProtectionType, not even an explicit nil
func (o *CommonObjectSummary) UnsetProtectionType() {
	o.ProtectionType.Unset()
}

// GetSharepointSiteSummary returns the SharepointSiteSummary field value if set, zero value otherwise.
func (o *CommonObjectSummary) GetSharepointSiteSummary() SharepointObjectParams {
	if o == nil || IsNil(o.SharepointSiteSummary) {
		var ret SharepointObjectParams
		return ret
	}
	return *o.SharepointSiteSummary
}

// GetSharepointSiteSummaryOk returns a tuple with the SharepointSiteSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonObjectSummary) GetSharepointSiteSummaryOk() (*SharepointObjectParams, bool) {
	if o == nil || IsNil(o.SharepointSiteSummary) {
		return nil, false
	}
	return o.SharepointSiteSummary, true
}

// HasSharepointSiteSummary returns a boolean if a field has been set.
func (o *CommonObjectSummary) HasSharepointSiteSummary() bool {
	if o != nil && !IsNil(o.SharepointSiteSummary) {
		return true
	}

	return false
}

// SetSharepointSiteSummary gets a reference to the given SharepointObjectParams and assigns it to the SharepointSiteSummary field.
func (o *CommonObjectSummary) SetSharepointSiteSummary(v SharepointObjectParams) {
	o.SharepointSiteSummary = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonObjectSummary) GetUuid() string {
	if o == nil || IsNil(o.Uuid.Get()) {
		var ret string
		return ret
	}
	return *o.Uuid.Get()
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonObjectSummary) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uuid.Get(), o.Uuid.IsSet()
}

// HasUuid returns a boolean if a field has been set.
func (o *CommonObjectSummary) HasUuid() bool {
	if o != nil && o.Uuid.IsSet() {
		return true
	}

	return false
}

// SetUuid gets a reference to the given NullableString and assigns it to the Uuid field.
func (o *CommonObjectSummary) SetUuid(v string) {
	o.Uuid.Set(&v)
}
// SetUuidNil sets the value for Uuid to be an explicit nil
func (o *CommonObjectSummary) SetUuidNil() {
	o.Uuid.Set(nil)
}

// UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
func (o *CommonObjectSummary) UnsetUuid() {
	o.Uuid.Unset()
}

// GetVCenterSummary returns the VCenterSummary field value if set, zero value otherwise.
func (o *CommonObjectSummary) GetVCenterSummary() ObjectTypeVCenterParams {
	if o == nil || IsNil(o.VCenterSummary) {
		var ret ObjectTypeVCenterParams
		return ret
	}
	return *o.VCenterSummary
}

// GetVCenterSummaryOk returns a tuple with the VCenterSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonObjectSummary) GetVCenterSummaryOk() (*ObjectTypeVCenterParams, bool) {
	if o == nil || IsNil(o.VCenterSummary) {
		return nil, false
	}
	return o.VCenterSummary, true
}

// HasVCenterSummary returns a boolean if a field has been set.
func (o *CommonObjectSummary) HasVCenterSummary() bool {
	if o != nil && !IsNil(o.VCenterSummary) {
		return true
	}

	return false
}

// SetVCenterSummary gets a reference to the given ObjectTypeVCenterParams and assigns it to the VCenterSummary field.
func (o *CommonObjectSummary) SetVCenterSummary(v ObjectTypeVCenterParams) {
	o.VCenterSummary = &v
}

// GetWindowsClusterSummary returns the WindowsClusterSummary field value if set, zero value otherwise.
func (o *CommonObjectSummary) GetWindowsClusterSummary() ObjectTypeWindowsClusterParams {
	if o == nil || IsNil(o.WindowsClusterSummary) {
		var ret ObjectTypeWindowsClusterParams
		return ret
	}
	return *o.WindowsClusterSummary
}

// GetWindowsClusterSummaryOk returns a tuple with the WindowsClusterSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonObjectSummary) GetWindowsClusterSummaryOk() (*ObjectTypeWindowsClusterParams, bool) {
	if o == nil || IsNil(o.WindowsClusterSummary) {
		return nil, false
	}
	return o.WindowsClusterSummary, true
}

// HasWindowsClusterSummary returns a boolean if a field has been set.
func (o *CommonObjectSummary) HasWindowsClusterSummary() bool {
	if o != nil && !IsNil(o.WindowsClusterSummary) {
		return true
	}

	return false
}

// SetWindowsClusterSummary gets a reference to the given ObjectTypeWindowsClusterParams and assigns it to the WindowsClusterSummary field.
func (o *CommonObjectSummary) SetWindowsClusterSummary(v ObjectTypeWindowsClusterParams) {
	o.WindowsClusterSummary = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *CommonObjectSummary) GetPermissions() PermissionInfo {
	if o == nil || IsNil(o.Permissions) {
		var ret PermissionInfo
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonObjectSummary) GetPermissionsOk() (*PermissionInfo, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *CommonObjectSummary) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given PermissionInfo and assigns it to the Permissions field.
func (o *CommonObjectSummary) SetPermissions(v PermissionInfo) {
	o.Permissions = &v
}

// GetProtectionStats returns the ProtectionStats field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonObjectSummary) GetProtectionStats() []ObjectProtectionStatsSummary {
	if o == nil {
		var ret []ObjectProtectionStatsSummary
		return ret
	}
	return o.ProtectionStats
}

// GetProtectionStatsOk returns a tuple with the ProtectionStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonObjectSummary) GetProtectionStatsOk() ([]ObjectProtectionStatsSummary, bool) {
	if o == nil || IsNil(o.ProtectionStats) {
		return nil, false
	}
	return o.ProtectionStats, true
}

// HasProtectionStats returns a boolean if a field has been set.
func (o *CommonObjectSummary) HasProtectionStats() bool {
	if o != nil && !IsNil(o.ProtectionStats) {
		return true
	}

	return false
}

// SetProtectionStats gets a reference to the given []ObjectProtectionStatsSummary and assigns it to the ProtectionStats field.
func (o *CommonObjectSummary) SetProtectionStats(v []ObjectProtectionStatsSummary) {
	o.ProtectionStats = v
}

func (o CommonObjectSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonObjectSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Environment.IsSet() {
		toSerialize["environment"] = o.Environment.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.SourceId.IsSet() {
		toSerialize["sourceId"] = o.SourceId.Get()
	}
	if o.SourceName.IsSet() {
		toSerialize["sourceName"] = o.SourceName.Get()
	}
	if o.ChildObjects != nil {
		toSerialize["childObjects"] = o.ChildObjects
	}
	if o.GlobalId.IsSet() {
		toSerialize["globalId"] = o.GlobalId.Get()
	}
	if o.LogicalSizeBytes.IsSet() {
		toSerialize["logicalSizeBytes"] = o.LogicalSizeBytes.Get()
	}
	if o.ObjectHash.IsSet() {
		toSerialize["objectHash"] = o.ObjectHash.Get()
	}
	if o.ObjectType.IsSet() {
		toSerialize["objectType"] = o.ObjectType.Get()
	}
	if o.OsType.IsSet() {
		toSerialize["osType"] = o.OsType.Get()
	}
	if o.ProtectionType.IsSet() {
		toSerialize["protectionType"] = o.ProtectionType.Get()
	}
	if !IsNil(o.SharepointSiteSummary) {
		toSerialize["sharepointSiteSummary"] = o.SharepointSiteSummary
	}
	if o.Uuid.IsSet() {
		toSerialize["uuid"] = o.Uuid.Get()
	}
	if !IsNil(o.VCenterSummary) {
		toSerialize["vCenterSummary"] = o.VCenterSummary
	}
	if !IsNil(o.WindowsClusterSummary) {
		toSerialize["windowsClusterSummary"] = o.WindowsClusterSummary
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	if o.ProtectionStats != nil {
		toSerialize["protectionStats"] = o.ProtectionStats
	}
	return toSerialize, nil
}

type NullableCommonObjectSummary struct {
	value *CommonObjectSummary
	isSet bool
}

func (v NullableCommonObjectSummary) Get() *CommonObjectSummary {
	return v.value
}

func (v *NullableCommonObjectSummary) Set(val *CommonObjectSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonObjectSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonObjectSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonObjectSummary(val *CommonObjectSummary) *NullableCommonObjectSummary {
	return &NullableCommonObjectSummary{value: val, isSet: true}
}

func (v NullableCommonObjectSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonObjectSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


