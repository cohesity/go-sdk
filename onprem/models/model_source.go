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

// Source Specifies the Protection Source.
type Source struct {
	// Specifies object id.
	Id NullableInt64 `json:"id,omitempty"`
	// Specifies the name of the object.
	Name NullableString `json:"name,omitempty"`
	// Specifies registered source id to which object belongs.
	SourceId NullableInt64 `json:"sourceId,omitempty"`
	// Specifies registered source name to which object belongs.
	SourceName NullableString `json:"sourceName,omitempty"`
	// Specifies the environment of the object.
	Environment NullableString `json:"environment,omitempty"`
	// Specifies the hash identifier of the object.
	ObjectHash NullableString `json:"objectHash,omitempty"`
	// Specifies the type of the object.
	ObjectType NullableString `json:"objectType,omitempty"`
	// Specifies the logical size of object in bytes.
	LogicalSizeBytes NullableInt64 `json:"logicalSizeBytes,omitempty"`
	// Specifies the uuid which is a unique identifier of the object.
	Uuid NullableString `json:"uuid,omitempty"`
	// Specifies the global id which is a unique identifier of the object.
	GlobalId NullableString `json:"globalId,omitempty"`
	// Specifies the protection type of the object if any.
	ProtectionType NullableString `json:"protectionType,omitempty"`
	// Specifies the operating system type of the object.
	OsType NullableString `json:"osType,omitempty"`
	VCenterSummary *ObjectTypeVCenterParams `json:"vCenterSummary,omitempty"`
	SharepointSiteSummary *SharepointObjectParams `json:"sharepointSiteSummary,omitempty"`
	// Specifies the count and size of protected and unprotected objects for the size.
	ProtectionStats []ObjectProtectionStatsSummary `json:"protectionStats,omitempty"`
	Permissions *PermissionInfo `json:"permissions,omitempty"`
	VmwareParams *VmwareObjectEntityParams `json:"vmwareParams,omitempty"`
	// Specifies the parameters for Isilon object.
	IsilonParams *IsilonObjectParams `json:"isilonParams,omitempty"`
	// Specifies the parameters for NetApp object.
	NetappParams *NetappObjectParams `json:"netappParams,omitempty"`
	// Specifies the parameters for GenericNas object.
	GenericNasParams *CommonNasObjectParams `json:"genericNasParams,omitempty"`
	// Specifies the parameters for Flashblade object.
	FlashbladeParams *FlashbladeObjectParams `json:"flashbladeParams,omitempty"`
	// Specifies the parameters for Elastifile object.
	ElastifileParams *CommonNasObjectParams `json:"elastifileParams,omitempty"`
	// Specifies the parameters for GPFS object.
	GpfsParams *CommonNasObjectParams `json:"gpfsParams,omitempty"`
	// Specifies the parameters for Msssql object.
	MssqlParams *MssqlObjectEntityParams `json:"mssqlParams,omitempty"`
	// Specifies the parameters for Oracle object.
	OracleParams *OracleObjectEntityParams `json:"oracleParams,omitempty"`
	// Specifies the parameters for Physical object.
	PhysicalParams *PhysicalObjectEntityParams `json:"physicalParams,omitempty"`
	// Specifies the parameters for Sharepoint object.
	SharepointParams *SharepointObjectEntityParams `json:"sharepointParams,omitempty"`
	// Time at which the data about this protection source was last refreshed.
	LastRefreshedTime NullableInt64 `json:"lastRefreshedTime,omitempty"`
	// Id of the registration as part of which this source was discovered.
	RegistrationId NullableInt64 `json:"registrationId,omitempty"`
}

// NewSource instantiates a new Source object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSource() *Source {
	this := Source{}
	return &this
}

// NewSourceWithDefaults instantiates a new Source object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceWithDefaults() *Source {
	this := Source{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Source) GetId() int64 {
	if o == nil || o.Id.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Source) GetIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Source) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *Source) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *Source) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Source) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Source) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Source) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Source) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Source) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *Source) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Source) UnsetName() {
	o.Name.Unset()
}

// GetSourceId returns the SourceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Source) GetSourceId() int64 {
	if o == nil || o.SourceId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.SourceId.Get()
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Source) GetSourceIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SourceId.Get(), o.SourceId.IsSet()
}

// HasSourceId returns a boolean if a field has been set.
func (o *Source) HasSourceId() bool {
	if o != nil && o.SourceId.IsSet() {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given NullableInt64 and assigns it to the SourceId field.
func (o *Source) SetSourceId(v int64) {
	o.SourceId.Set(&v)
}
// SetSourceIdNil sets the value for SourceId to be an explicit nil
func (o *Source) SetSourceIdNil() {
	o.SourceId.Set(nil)
}

// UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
func (o *Source) UnsetSourceId() {
	o.SourceId.Unset()
}

// GetSourceName returns the SourceName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Source) GetSourceName() string {
	if o == nil || o.SourceName.Get() == nil {
		var ret string
		return ret
	}
	return *o.SourceName.Get()
}

// GetSourceNameOk returns a tuple with the SourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Source) GetSourceNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SourceName.Get(), o.SourceName.IsSet()
}

// HasSourceName returns a boolean if a field has been set.
func (o *Source) HasSourceName() bool {
	if o != nil && o.SourceName.IsSet() {
		return true
	}

	return false
}

// SetSourceName gets a reference to the given NullableString and assigns it to the SourceName field.
func (o *Source) SetSourceName(v string) {
	o.SourceName.Set(&v)
}
// SetSourceNameNil sets the value for SourceName to be an explicit nil
func (o *Source) SetSourceNameNil() {
	o.SourceName.Set(nil)
}

// UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil
func (o *Source) UnsetSourceName() {
	o.SourceName.Unset()
}

// GetEnvironment returns the Environment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Source) GetEnvironment() string {
	if o == nil || o.Environment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Environment.Get()
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Source) GetEnvironmentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Environment.Get(), o.Environment.IsSet()
}

// HasEnvironment returns a boolean if a field has been set.
func (o *Source) HasEnvironment() bool {
	if o != nil && o.Environment.IsSet() {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given NullableString and assigns it to the Environment field.
func (o *Source) SetEnvironment(v string) {
	o.Environment.Set(&v)
}
// SetEnvironmentNil sets the value for Environment to be an explicit nil
func (o *Source) SetEnvironmentNil() {
	o.Environment.Set(nil)
}

// UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
func (o *Source) UnsetEnvironment() {
	o.Environment.Unset()
}

// GetObjectHash returns the ObjectHash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Source) GetObjectHash() string {
	if o == nil || o.ObjectHash.Get() == nil {
		var ret string
		return ret
	}
	return *o.ObjectHash.Get()
}

// GetObjectHashOk returns a tuple with the ObjectHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Source) GetObjectHashOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ObjectHash.Get(), o.ObjectHash.IsSet()
}

// HasObjectHash returns a boolean if a field has been set.
func (o *Source) HasObjectHash() bool {
	if o != nil && o.ObjectHash.IsSet() {
		return true
	}

	return false
}

// SetObjectHash gets a reference to the given NullableString and assigns it to the ObjectHash field.
func (o *Source) SetObjectHash(v string) {
	o.ObjectHash.Set(&v)
}
// SetObjectHashNil sets the value for ObjectHash to be an explicit nil
func (o *Source) SetObjectHashNil() {
	o.ObjectHash.Set(nil)
}

// UnsetObjectHash ensures that no value is present for ObjectHash, not even an explicit nil
func (o *Source) UnsetObjectHash() {
	o.ObjectHash.Unset()
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Source) GetObjectType() string {
	if o == nil || o.ObjectType.Get() == nil {
		var ret string
		return ret
	}
	return *o.ObjectType.Get()
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Source) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ObjectType.Get(), o.ObjectType.IsSet()
}

// HasObjectType returns a boolean if a field has been set.
func (o *Source) HasObjectType() bool {
	if o != nil && o.ObjectType.IsSet() {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given NullableString and assigns it to the ObjectType field.
func (o *Source) SetObjectType(v string) {
	o.ObjectType.Set(&v)
}
// SetObjectTypeNil sets the value for ObjectType to be an explicit nil
func (o *Source) SetObjectTypeNil() {
	o.ObjectType.Set(nil)
}

// UnsetObjectType ensures that no value is present for ObjectType, not even an explicit nil
func (o *Source) UnsetObjectType() {
	o.ObjectType.Unset()
}

// GetLogicalSizeBytes returns the LogicalSizeBytes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Source) GetLogicalSizeBytes() int64 {
	if o == nil || o.LogicalSizeBytes.Get() == nil {
		var ret int64
		return ret
	}
	return *o.LogicalSizeBytes.Get()
}

// GetLogicalSizeBytesOk returns a tuple with the LogicalSizeBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Source) GetLogicalSizeBytesOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LogicalSizeBytes.Get(), o.LogicalSizeBytes.IsSet()
}

// HasLogicalSizeBytes returns a boolean if a field has been set.
func (o *Source) HasLogicalSizeBytes() bool {
	if o != nil && o.LogicalSizeBytes.IsSet() {
		return true
	}

	return false
}

// SetLogicalSizeBytes gets a reference to the given NullableInt64 and assigns it to the LogicalSizeBytes field.
func (o *Source) SetLogicalSizeBytes(v int64) {
	o.LogicalSizeBytes.Set(&v)
}
// SetLogicalSizeBytesNil sets the value for LogicalSizeBytes to be an explicit nil
func (o *Source) SetLogicalSizeBytesNil() {
	o.LogicalSizeBytes.Set(nil)
}

// UnsetLogicalSizeBytes ensures that no value is present for LogicalSizeBytes, not even an explicit nil
func (o *Source) UnsetLogicalSizeBytes() {
	o.LogicalSizeBytes.Unset()
}

// GetUuid returns the Uuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Source) GetUuid() string {
	if o == nil || o.Uuid.Get() == nil {
		var ret string
		return ret
	}
	return *o.Uuid.Get()
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Source) GetUuidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Uuid.Get(), o.Uuid.IsSet()
}

// HasUuid returns a boolean if a field has been set.
func (o *Source) HasUuid() bool {
	if o != nil && o.Uuid.IsSet() {
		return true
	}

	return false
}

// SetUuid gets a reference to the given NullableString and assigns it to the Uuid field.
func (o *Source) SetUuid(v string) {
	o.Uuid.Set(&v)
}
// SetUuidNil sets the value for Uuid to be an explicit nil
func (o *Source) SetUuidNil() {
	o.Uuid.Set(nil)
}

// UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
func (o *Source) UnsetUuid() {
	o.Uuid.Unset()
}

// GetGlobalId returns the GlobalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Source) GetGlobalId() string {
	if o == nil || o.GlobalId.Get() == nil {
		var ret string
		return ret
	}
	return *o.GlobalId.Get()
}

// GetGlobalIdOk returns a tuple with the GlobalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Source) GetGlobalIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.GlobalId.Get(), o.GlobalId.IsSet()
}

// HasGlobalId returns a boolean if a field has been set.
func (o *Source) HasGlobalId() bool {
	if o != nil && o.GlobalId.IsSet() {
		return true
	}

	return false
}

// SetGlobalId gets a reference to the given NullableString and assigns it to the GlobalId field.
func (o *Source) SetGlobalId(v string) {
	o.GlobalId.Set(&v)
}
// SetGlobalIdNil sets the value for GlobalId to be an explicit nil
func (o *Source) SetGlobalIdNil() {
	o.GlobalId.Set(nil)
}

// UnsetGlobalId ensures that no value is present for GlobalId, not even an explicit nil
func (o *Source) UnsetGlobalId() {
	o.GlobalId.Unset()
}

// GetProtectionType returns the ProtectionType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Source) GetProtectionType() string {
	if o == nil || o.ProtectionType.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProtectionType.Get()
}

// GetProtectionTypeOk returns a tuple with the ProtectionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Source) GetProtectionTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ProtectionType.Get(), o.ProtectionType.IsSet()
}

// HasProtectionType returns a boolean if a field has been set.
func (o *Source) HasProtectionType() bool {
	if o != nil && o.ProtectionType.IsSet() {
		return true
	}

	return false
}

// SetProtectionType gets a reference to the given NullableString and assigns it to the ProtectionType field.
func (o *Source) SetProtectionType(v string) {
	o.ProtectionType.Set(&v)
}
// SetProtectionTypeNil sets the value for ProtectionType to be an explicit nil
func (o *Source) SetProtectionTypeNil() {
	o.ProtectionType.Set(nil)
}

// UnsetProtectionType ensures that no value is present for ProtectionType, not even an explicit nil
func (o *Source) UnsetProtectionType() {
	o.ProtectionType.Unset()
}

// GetOsType returns the OsType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Source) GetOsType() string {
	if o == nil || o.OsType.Get() == nil {
		var ret string
		return ret
	}
	return *o.OsType.Get()
}

// GetOsTypeOk returns a tuple with the OsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Source) GetOsTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OsType.Get(), o.OsType.IsSet()
}

// HasOsType returns a boolean if a field has been set.
func (o *Source) HasOsType() bool {
	if o != nil && o.OsType.IsSet() {
		return true
	}

	return false
}

// SetOsType gets a reference to the given NullableString and assigns it to the OsType field.
func (o *Source) SetOsType(v string) {
	o.OsType.Set(&v)
}
// SetOsTypeNil sets the value for OsType to be an explicit nil
func (o *Source) SetOsTypeNil() {
	o.OsType.Set(nil)
}

// UnsetOsType ensures that no value is present for OsType, not even an explicit nil
func (o *Source) UnsetOsType() {
	o.OsType.Unset()
}

// GetVCenterSummary returns the VCenterSummary field value if set, zero value otherwise.
func (o *Source) GetVCenterSummary() ObjectTypeVCenterParams {
	if o == nil || o.VCenterSummary == nil {
		var ret ObjectTypeVCenterParams
		return ret
	}
	return *o.VCenterSummary
}

// GetVCenterSummaryOk returns a tuple with the VCenterSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetVCenterSummaryOk() (*ObjectTypeVCenterParams, bool) {
	if o == nil || o.VCenterSummary == nil {
		return nil, false
	}
	return o.VCenterSummary, true
}

// HasVCenterSummary returns a boolean if a field has been set.
func (o *Source) HasVCenterSummary() bool {
	if o != nil && o.VCenterSummary != nil {
		return true
	}

	return false
}

// SetVCenterSummary gets a reference to the given ObjectTypeVCenterParams and assigns it to the VCenterSummary field.
func (o *Source) SetVCenterSummary(v ObjectTypeVCenterParams) {
	o.VCenterSummary = &v
}

// GetSharepointSiteSummary returns the SharepointSiteSummary field value if set, zero value otherwise.
func (o *Source) GetSharepointSiteSummary() SharepointObjectParams {
	if o == nil || o.SharepointSiteSummary == nil {
		var ret SharepointObjectParams
		return ret
	}
	return *o.SharepointSiteSummary
}

// GetSharepointSiteSummaryOk returns a tuple with the SharepointSiteSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetSharepointSiteSummaryOk() (*SharepointObjectParams, bool) {
	if o == nil || o.SharepointSiteSummary == nil {
		return nil, false
	}
	return o.SharepointSiteSummary, true
}

// HasSharepointSiteSummary returns a boolean if a field has been set.
func (o *Source) HasSharepointSiteSummary() bool {
	if o != nil && o.SharepointSiteSummary != nil {
		return true
	}

	return false
}

// SetSharepointSiteSummary gets a reference to the given SharepointObjectParams and assigns it to the SharepointSiteSummary field.
func (o *Source) SetSharepointSiteSummary(v SharepointObjectParams) {
	o.SharepointSiteSummary = &v
}

// GetProtectionStats returns the ProtectionStats field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Source) GetProtectionStats() []ObjectProtectionStatsSummary {
	if o == nil  {
		var ret []ObjectProtectionStatsSummary
		return ret
	}
	return o.ProtectionStats
}

// GetProtectionStatsOk returns a tuple with the ProtectionStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Source) GetProtectionStatsOk() (*[]ObjectProtectionStatsSummary, bool) {
	if o == nil || o.ProtectionStats == nil {
		return nil, false
	}
	return &o.ProtectionStats, true
}

// HasProtectionStats returns a boolean if a field has been set.
func (o *Source) HasProtectionStats() bool {
	if o != nil && o.ProtectionStats != nil {
		return true
	}

	return false
}

// SetProtectionStats gets a reference to the given []ObjectProtectionStatsSummary and assigns it to the ProtectionStats field.
func (o *Source) SetProtectionStats(v []ObjectProtectionStatsSummary) {
	o.ProtectionStats = v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *Source) GetPermissions() PermissionInfo {
	if o == nil || o.Permissions == nil {
		var ret PermissionInfo
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetPermissionsOk() (*PermissionInfo, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *Source) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given PermissionInfo and assigns it to the Permissions field.
func (o *Source) SetPermissions(v PermissionInfo) {
	o.Permissions = &v
}

// GetVmwareParams returns the VmwareParams field value if set, zero value otherwise.
func (o *Source) GetVmwareParams() VmwareObjectEntityParams {
	if o == nil || o.VmwareParams == nil {
		var ret VmwareObjectEntityParams
		return ret
	}
	return *o.VmwareParams
}

// GetVmwareParamsOk returns a tuple with the VmwareParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetVmwareParamsOk() (*VmwareObjectEntityParams, bool) {
	if o == nil || o.VmwareParams == nil {
		return nil, false
	}
	return o.VmwareParams, true
}

// HasVmwareParams returns a boolean if a field has been set.
func (o *Source) HasVmwareParams() bool {
	if o != nil && o.VmwareParams != nil {
		return true
	}

	return false
}

// SetVmwareParams gets a reference to the given VmwareObjectEntityParams and assigns it to the VmwareParams field.
func (o *Source) SetVmwareParams(v VmwareObjectEntityParams) {
	o.VmwareParams = &v
}

// GetIsilonParams returns the IsilonParams field value if set, zero value otherwise.
func (o *Source) GetIsilonParams() IsilonObjectParams {
	if o == nil || o.IsilonParams == nil {
		var ret IsilonObjectParams
		return ret
	}
	return *o.IsilonParams
}

// GetIsilonParamsOk returns a tuple with the IsilonParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetIsilonParamsOk() (*IsilonObjectParams, bool) {
	if o == nil || o.IsilonParams == nil {
		return nil, false
	}
	return o.IsilonParams, true
}

// HasIsilonParams returns a boolean if a field has been set.
func (o *Source) HasIsilonParams() bool {
	if o != nil && o.IsilonParams != nil {
		return true
	}

	return false
}

// SetIsilonParams gets a reference to the given IsilonObjectParams and assigns it to the IsilonParams field.
func (o *Source) SetIsilonParams(v IsilonObjectParams) {
	o.IsilonParams = &v
}

// GetNetappParams returns the NetappParams field value if set, zero value otherwise.
func (o *Source) GetNetappParams() NetappObjectParams {
	if o == nil || o.NetappParams == nil {
		var ret NetappObjectParams
		return ret
	}
	return *o.NetappParams
}

// GetNetappParamsOk returns a tuple with the NetappParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetNetappParamsOk() (*NetappObjectParams, bool) {
	if o == nil || o.NetappParams == nil {
		return nil, false
	}
	return o.NetappParams, true
}

// HasNetappParams returns a boolean if a field has been set.
func (o *Source) HasNetappParams() bool {
	if o != nil && o.NetappParams != nil {
		return true
	}

	return false
}

// SetNetappParams gets a reference to the given NetappObjectParams and assigns it to the NetappParams field.
func (o *Source) SetNetappParams(v NetappObjectParams) {
	o.NetappParams = &v
}

// GetGenericNasParams returns the GenericNasParams field value if set, zero value otherwise.
func (o *Source) GetGenericNasParams() CommonNasObjectParams {
	if o == nil || o.GenericNasParams == nil {
		var ret CommonNasObjectParams
		return ret
	}
	return *o.GenericNasParams
}

// GetGenericNasParamsOk returns a tuple with the GenericNasParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetGenericNasParamsOk() (*CommonNasObjectParams, bool) {
	if o == nil || o.GenericNasParams == nil {
		return nil, false
	}
	return o.GenericNasParams, true
}

// HasGenericNasParams returns a boolean if a field has been set.
func (o *Source) HasGenericNasParams() bool {
	if o != nil && o.GenericNasParams != nil {
		return true
	}

	return false
}

// SetGenericNasParams gets a reference to the given CommonNasObjectParams and assigns it to the GenericNasParams field.
func (o *Source) SetGenericNasParams(v CommonNasObjectParams) {
	o.GenericNasParams = &v
}

// GetFlashbladeParams returns the FlashbladeParams field value if set, zero value otherwise.
func (o *Source) GetFlashbladeParams() FlashbladeObjectParams {
	if o == nil || o.FlashbladeParams == nil {
		var ret FlashbladeObjectParams
		return ret
	}
	return *o.FlashbladeParams
}

// GetFlashbladeParamsOk returns a tuple with the FlashbladeParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetFlashbladeParamsOk() (*FlashbladeObjectParams, bool) {
	if o == nil || o.FlashbladeParams == nil {
		return nil, false
	}
	return o.FlashbladeParams, true
}

// HasFlashbladeParams returns a boolean if a field has been set.
func (o *Source) HasFlashbladeParams() bool {
	if o != nil && o.FlashbladeParams != nil {
		return true
	}

	return false
}

// SetFlashbladeParams gets a reference to the given FlashbladeObjectParams and assigns it to the FlashbladeParams field.
func (o *Source) SetFlashbladeParams(v FlashbladeObjectParams) {
	o.FlashbladeParams = &v
}

// GetElastifileParams returns the ElastifileParams field value if set, zero value otherwise.
func (o *Source) GetElastifileParams() CommonNasObjectParams {
	if o == nil || o.ElastifileParams == nil {
		var ret CommonNasObjectParams
		return ret
	}
	return *o.ElastifileParams
}

// GetElastifileParamsOk returns a tuple with the ElastifileParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetElastifileParamsOk() (*CommonNasObjectParams, bool) {
	if o == nil || o.ElastifileParams == nil {
		return nil, false
	}
	return o.ElastifileParams, true
}

// HasElastifileParams returns a boolean if a field has been set.
func (o *Source) HasElastifileParams() bool {
	if o != nil && o.ElastifileParams != nil {
		return true
	}

	return false
}

// SetElastifileParams gets a reference to the given CommonNasObjectParams and assigns it to the ElastifileParams field.
func (o *Source) SetElastifileParams(v CommonNasObjectParams) {
	o.ElastifileParams = &v
}

// GetGpfsParams returns the GpfsParams field value if set, zero value otherwise.
func (o *Source) GetGpfsParams() CommonNasObjectParams {
	if o == nil || o.GpfsParams == nil {
		var ret CommonNasObjectParams
		return ret
	}
	return *o.GpfsParams
}

// GetGpfsParamsOk returns a tuple with the GpfsParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetGpfsParamsOk() (*CommonNasObjectParams, bool) {
	if o == nil || o.GpfsParams == nil {
		return nil, false
	}
	return o.GpfsParams, true
}

// HasGpfsParams returns a boolean if a field has been set.
func (o *Source) HasGpfsParams() bool {
	if o != nil && o.GpfsParams != nil {
		return true
	}

	return false
}

// SetGpfsParams gets a reference to the given CommonNasObjectParams and assigns it to the GpfsParams field.
func (o *Source) SetGpfsParams(v CommonNasObjectParams) {
	o.GpfsParams = &v
}

// GetMssqlParams returns the MssqlParams field value if set, zero value otherwise.
func (o *Source) GetMssqlParams() MssqlObjectEntityParams {
	if o == nil || o.MssqlParams == nil {
		var ret MssqlObjectEntityParams
		return ret
	}
	return *o.MssqlParams
}

// GetMssqlParamsOk returns a tuple with the MssqlParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetMssqlParamsOk() (*MssqlObjectEntityParams, bool) {
	if o == nil || o.MssqlParams == nil {
		return nil, false
	}
	return o.MssqlParams, true
}

// HasMssqlParams returns a boolean if a field has been set.
func (o *Source) HasMssqlParams() bool {
	if o != nil && o.MssqlParams != nil {
		return true
	}

	return false
}

// SetMssqlParams gets a reference to the given MssqlObjectEntityParams and assigns it to the MssqlParams field.
func (o *Source) SetMssqlParams(v MssqlObjectEntityParams) {
	o.MssqlParams = &v
}

// GetOracleParams returns the OracleParams field value if set, zero value otherwise.
func (o *Source) GetOracleParams() OracleObjectEntityParams {
	if o == nil || o.OracleParams == nil {
		var ret OracleObjectEntityParams
		return ret
	}
	return *o.OracleParams
}

// GetOracleParamsOk returns a tuple with the OracleParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetOracleParamsOk() (*OracleObjectEntityParams, bool) {
	if o == nil || o.OracleParams == nil {
		return nil, false
	}
	return o.OracleParams, true
}

// HasOracleParams returns a boolean if a field has been set.
func (o *Source) HasOracleParams() bool {
	if o != nil && o.OracleParams != nil {
		return true
	}

	return false
}

// SetOracleParams gets a reference to the given OracleObjectEntityParams and assigns it to the OracleParams field.
func (o *Source) SetOracleParams(v OracleObjectEntityParams) {
	o.OracleParams = &v
}

// GetPhysicalParams returns the PhysicalParams field value if set, zero value otherwise.
func (o *Source) GetPhysicalParams() PhysicalObjectEntityParams {
	if o == nil || o.PhysicalParams == nil {
		var ret PhysicalObjectEntityParams
		return ret
	}
	return *o.PhysicalParams
}

// GetPhysicalParamsOk returns a tuple with the PhysicalParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetPhysicalParamsOk() (*PhysicalObjectEntityParams, bool) {
	if o == nil || o.PhysicalParams == nil {
		return nil, false
	}
	return o.PhysicalParams, true
}

// HasPhysicalParams returns a boolean if a field has been set.
func (o *Source) HasPhysicalParams() bool {
	if o != nil && o.PhysicalParams != nil {
		return true
	}

	return false
}

// SetPhysicalParams gets a reference to the given PhysicalObjectEntityParams and assigns it to the PhysicalParams field.
func (o *Source) SetPhysicalParams(v PhysicalObjectEntityParams) {
	o.PhysicalParams = &v
}

// GetSharepointParams returns the SharepointParams field value if set, zero value otherwise.
func (o *Source) GetSharepointParams() SharepointObjectEntityParams {
	if o == nil || o.SharepointParams == nil {
		var ret SharepointObjectEntityParams
		return ret
	}
	return *o.SharepointParams
}

// GetSharepointParamsOk returns a tuple with the SharepointParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetSharepointParamsOk() (*SharepointObjectEntityParams, bool) {
	if o == nil || o.SharepointParams == nil {
		return nil, false
	}
	return o.SharepointParams, true
}

// HasSharepointParams returns a boolean if a field has been set.
func (o *Source) HasSharepointParams() bool {
	if o != nil && o.SharepointParams != nil {
		return true
	}

	return false
}

// SetSharepointParams gets a reference to the given SharepointObjectEntityParams and assigns it to the SharepointParams field.
func (o *Source) SetSharepointParams(v SharepointObjectEntityParams) {
	o.SharepointParams = &v
}

// GetLastRefreshedTime returns the LastRefreshedTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Source) GetLastRefreshedTime() int64 {
	if o == nil || o.LastRefreshedTime.Get() == nil {
		var ret int64
		return ret
	}
	return *o.LastRefreshedTime.Get()
}

// GetLastRefreshedTimeOk returns a tuple with the LastRefreshedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Source) GetLastRefreshedTimeOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastRefreshedTime.Get(), o.LastRefreshedTime.IsSet()
}

// HasLastRefreshedTime returns a boolean if a field has been set.
func (o *Source) HasLastRefreshedTime() bool {
	if o != nil && o.LastRefreshedTime.IsSet() {
		return true
	}

	return false
}

// SetLastRefreshedTime gets a reference to the given NullableInt64 and assigns it to the LastRefreshedTime field.
func (o *Source) SetLastRefreshedTime(v int64) {
	o.LastRefreshedTime.Set(&v)
}
// SetLastRefreshedTimeNil sets the value for LastRefreshedTime to be an explicit nil
func (o *Source) SetLastRefreshedTimeNil() {
	o.LastRefreshedTime.Set(nil)
}

// UnsetLastRefreshedTime ensures that no value is present for LastRefreshedTime, not even an explicit nil
func (o *Source) UnsetLastRefreshedTime() {
	o.LastRefreshedTime.Unset()
}

// GetRegistrationId returns the RegistrationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Source) GetRegistrationId() int64 {
	if o == nil || o.RegistrationId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RegistrationId.Get()
}

// GetRegistrationIdOk returns a tuple with the RegistrationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Source) GetRegistrationIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RegistrationId.Get(), o.RegistrationId.IsSet()
}

// HasRegistrationId returns a boolean if a field has been set.
func (o *Source) HasRegistrationId() bool {
	if o != nil && o.RegistrationId.IsSet() {
		return true
	}

	return false
}

// SetRegistrationId gets a reference to the given NullableInt64 and assigns it to the RegistrationId field.
func (o *Source) SetRegistrationId(v int64) {
	o.RegistrationId.Set(&v)
}
// SetRegistrationIdNil sets the value for RegistrationId to be an explicit nil
func (o *Source) SetRegistrationIdNil() {
	o.RegistrationId.Set(nil)
}

// UnsetRegistrationId ensures that no value is present for RegistrationId, not even an explicit nil
func (o *Source) UnsetRegistrationId() {
	o.RegistrationId.Unset()
}

func (o Source) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	if o.Environment.IsSet() {
		toSerialize["environment"] = o.Environment.Get()
	}
	if o.ObjectHash.IsSet() {
		toSerialize["objectHash"] = o.ObjectHash.Get()
	}
	if o.ObjectType.IsSet() {
		toSerialize["objectType"] = o.ObjectType.Get()
	}
	if o.LogicalSizeBytes.IsSet() {
		toSerialize["logicalSizeBytes"] = o.LogicalSizeBytes.Get()
	}
	if o.Uuid.IsSet() {
		toSerialize["uuid"] = o.Uuid.Get()
	}
	if o.GlobalId.IsSet() {
		toSerialize["globalId"] = o.GlobalId.Get()
	}
	if o.ProtectionType.IsSet() {
		toSerialize["protectionType"] = o.ProtectionType.Get()
	}
	if o.OsType.IsSet() {
		toSerialize["osType"] = o.OsType.Get()
	}
	if o.VCenterSummary != nil {
		toSerialize["vCenterSummary"] = o.VCenterSummary
	}
	if o.SharepointSiteSummary != nil {
		toSerialize["sharepointSiteSummary"] = o.SharepointSiteSummary
	}
	if o.ProtectionStats != nil {
		toSerialize["protectionStats"] = o.ProtectionStats
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	if o.VmwareParams != nil {
		toSerialize["vmwareParams"] = o.VmwareParams
	}
	if o.IsilonParams != nil {
		toSerialize["isilonParams"] = o.IsilonParams
	}
	if o.NetappParams != nil {
		toSerialize["netappParams"] = o.NetappParams
	}
	if o.GenericNasParams != nil {
		toSerialize["genericNasParams"] = o.GenericNasParams
	}
	if o.FlashbladeParams != nil {
		toSerialize["flashbladeParams"] = o.FlashbladeParams
	}
	if o.ElastifileParams != nil {
		toSerialize["elastifileParams"] = o.ElastifileParams
	}
	if o.GpfsParams != nil {
		toSerialize["gpfsParams"] = o.GpfsParams
	}
	if o.MssqlParams != nil {
		toSerialize["mssqlParams"] = o.MssqlParams
	}
	if o.OracleParams != nil {
		toSerialize["oracleParams"] = o.OracleParams
	}
	if o.PhysicalParams != nil {
		toSerialize["physicalParams"] = o.PhysicalParams
	}
	if o.SharepointParams != nil {
		toSerialize["sharepointParams"] = o.SharepointParams
	}
	if o.LastRefreshedTime.IsSet() {
		toSerialize["lastRefreshedTime"] = o.LastRefreshedTime.Get()
	}
	if o.RegistrationId.IsSet() {
		toSerialize["registrationId"] = o.RegistrationId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSource struct {
	value *Source
	isSet bool
}

func (v NullableSource) Get() *Source {
	return v.value
}

func (v *NullableSource) Set(val *Source) {
	v.value = val
	v.isSet = true
}

func (v NullableSource) IsSet() bool {
	return v.isSet
}

func (v *NullableSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSource(val *Source) *NullableSource {
	return &NullableSource{value: val, isSet: true}
}

func (v NullableSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o Source) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}