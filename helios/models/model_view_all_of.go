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

// ViewAllOf struct for ViewAllOf
type ViewAllOf struct {
	// Specifies an id of the View assigned by the Cohesity Cluster.
	ViewId NullableInt64 `json:"viewId,omitempty"`
	// If True, category in response is not set by user but inferred by Iris because none is set. Category can only be none when view was created by v1 API or cloned from a view created by v1 API.  Inference Logic is as follows: 1. Object Services if only S3 or Swift protocol is selected. 2. Backup Target only if one read-write protocol is selected and    QoS is \"Backup Target Commvault\" or \"Backup Target SSD\". 3. File Services if there are more than 1 read-write protocol or    it doesn't fit any other category.
	IsCategoryInferred NullableBool `json:"isCategoryInferred,omitempty"`
	// DataLock (Write Once Read Many) lock expiry epoch time in microseconds. If a view is marked as a DataLock view, only a Data Security Officer (a user having Data Security Privilege) can delete the view until the lock expiry time.
	DataLockExpiryUsecs NullableInt64 `json:"dataLockExpiryUsecs,omitempty"`
	// Specifies the Object Services key mapping config of the view. This parameter can only be set during create and cannot be changed. Configuration of Object Services key mapping. Specifies the type of Object Services key mapping config.
	ObjectServicesMappingConfig NullableString `json:"objectServicesMappingConfig,omitempty"`
	// Specifies the id of the Storage Domain (View Box) where the View is stored.
	StorageDomainId NullableInt64 `json:"storageDomainId,omitempty"`
	// Specifies the name of the Storage Domain (View Box) where the View is stored.
	StorageDomainName NullableString `json:"storageDomainName,omitempty"`
	// Specifies whether to support case insensitive file/folder names. This parameter can only be set during create and cannot be changed.
	CaseInsensitiveNamesEnabled NullableBool `json:"caseInsensitiveNamesEnabled,omitempty"`
	// Specifies the time that the View was created in milliseconds.
	CreateTimeMsecs NullableInt64 `json:"createTimeMsecs,omitempty"`
	// Specifies the NFS mount path of the View (without the hostname information). This path is used to support NFS mounting of the paths specified in the nfsExportPathList on Windows systems.
	BasicMountPath NullableString `json:"basicMountPath,omitempty"`
	// This field is currently deprecated. Please use NFS MountPaths which would be an array of strings.
	NfsMountPath NullableString `json:"nfsMountPath,omitempty"`
	// Array of NFS Paths. Specifies the path for mounting this View as an NFS share. If Kerberos Provider has multiple hostaliases, each host alias has  its own path.
	NfsMountPaths []string `json:"nfsMountPaths,omitempty"`
	// Array of SMB Paths. Specifies the possible paths that can be used to mount this View as a SMB share. If Active Directory has multiple account names; each machine account has its own path.
	SmbMountPaths []string `json:"smbMountPaths,omitempty"`
	ViewProtection *ViewProtection `json:"viewProtection,omitempty"`
	// Aliases created for the view. A view alias allows a directory path inside a view to be mounted using the alias name.
	Aliases []ViewAliasInfo `json:"aliases,omitempty"`
	// Specifies if a view contains migrated data.
	IsTargetForMigratedData NullableBool `json:"isTargetForMigratedData,omitempty"`
	// Specifies the information about the failover of the view.
	ViewFailover *ViewFailover `json:"viewFailover,omitempty"`
	// Specifies statistics about the View.
	Stats *ViewStats `json:"stats,omitempty"`
	// Specifies the file count by size for the View.
	FileCountBySize *[]FileCount `json:"fileCountBySize,omitempty"`
	// Specifies the sid of the view owner.
	OwnerSid NullableString `json:"ownerSid,omitempty"`
}

// NewViewAllOf instantiates a new ViewAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViewAllOf() *ViewAllOf {
	this := ViewAllOf{}
	return &this
}

// NewViewAllOfWithDefaults instantiates a new ViewAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewAllOfWithDefaults() *ViewAllOf {
	this := ViewAllOf{}
	return &this
}

// GetViewId returns the ViewId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ViewAllOf) GetViewId() int64 {
	if o == nil || o.ViewId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ViewId.Get()
}

// GetViewIdOk returns a tuple with the ViewId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ViewAllOf) GetViewIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ViewId.Get(), o.ViewId.IsSet()
}

// HasViewId returns a boolean if a field has been set.
func (o *ViewAllOf) HasViewId() bool {
	if o != nil && o.ViewId.IsSet() {
		return true
	}

	return false
}

// SetViewId gets a reference to the given NullableInt64 and assigns it to the ViewId field.
func (o *ViewAllOf) SetViewId(v int64) {
	o.ViewId.Set(&v)
}
// SetViewIdNil sets the value for ViewId to be an explicit nil
func (o *ViewAllOf) SetViewIdNil() {
	o.ViewId.Set(nil)
}

// UnsetViewId ensures that no value is present for ViewId, not even an explicit nil
func (o *ViewAllOf) UnsetViewId() {
	o.ViewId.Unset()
}

// GetIsCategoryInferred returns the IsCategoryInferred field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ViewAllOf) GetIsCategoryInferred() bool {
	if o == nil || o.IsCategoryInferred.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsCategoryInferred.Get()
}

// GetIsCategoryInferredOk returns a tuple with the IsCategoryInferred field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ViewAllOf) GetIsCategoryInferredOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsCategoryInferred.Get(), o.IsCategoryInferred.IsSet()
}

// HasIsCategoryInferred returns a boolean if a field has been set.
func (o *ViewAllOf) HasIsCategoryInferred() bool {
	if o != nil && o.IsCategoryInferred.IsSet() {
		return true
	}

	return false
}

// SetIsCategoryInferred gets a reference to the given NullableBool and assigns it to the IsCategoryInferred field.
func (o *ViewAllOf) SetIsCategoryInferred(v bool) {
	o.IsCategoryInferred.Set(&v)
}
// SetIsCategoryInferredNil sets the value for IsCategoryInferred to be an explicit nil
func (o *ViewAllOf) SetIsCategoryInferredNil() {
	o.IsCategoryInferred.Set(nil)
}

// UnsetIsCategoryInferred ensures that no value is present for IsCategoryInferred, not even an explicit nil
func (o *ViewAllOf) UnsetIsCategoryInferred() {
	o.IsCategoryInferred.Unset()
}

// GetDataLockExpiryUsecs returns the DataLockExpiryUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ViewAllOf) GetDataLockExpiryUsecs() int64 {
	if o == nil || o.DataLockExpiryUsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.DataLockExpiryUsecs.Get()
}

// GetDataLockExpiryUsecsOk returns a tuple with the DataLockExpiryUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ViewAllOf) GetDataLockExpiryUsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DataLockExpiryUsecs.Get(), o.DataLockExpiryUsecs.IsSet()
}

// HasDataLockExpiryUsecs returns a boolean if a field has been set.
func (o *ViewAllOf) HasDataLockExpiryUsecs() bool {
	if o != nil && o.DataLockExpiryUsecs.IsSet() {
		return true
	}

	return false
}

// SetDataLockExpiryUsecs gets a reference to the given NullableInt64 and assigns it to the DataLockExpiryUsecs field.
func (o *ViewAllOf) SetDataLockExpiryUsecs(v int64) {
	o.DataLockExpiryUsecs.Set(&v)
}
// SetDataLockExpiryUsecsNil sets the value for DataLockExpiryUsecs to be an explicit nil
func (o *ViewAllOf) SetDataLockExpiryUsecsNil() {
	o.DataLockExpiryUsecs.Set(nil)
}

// UnsetDataLockExpiryUsecs ensures that no value is present for DataLockExpiryUsecs, not even an explicit nil
func (o *ViewAllOf) UnsetDataLockExpiryUsecs() {
	o.DataLockExpiryUsecs.Unset()
}

// GetObjectServicesMappingConfig returns the ObjectServicesMappingConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ViewAllOf) GetObjectServicesMappingConfig() string {
	if o == nil || o.ObjectServicesMappingConfig.Get() == nil {
		var ret string
		return ret
	}
	return *o.ObjectServicesMappingConfig.Get()
}

// GetObjectServicesMappingConfigOk returns a tuple with the ObjectServicesMappingConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ViewAllOf) GetObjectServicesMappingConfigOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ObjectServicesMappingConfig.Get(), o.ObjectServicesMappingConfig.IsSet()
}

// HasObjectServicesMappingConfig returns a boolean if a field has been set.
func (o *ViewAllOf) HasObjectServicesMappingConfig() bool {
	if o != nil && o.ObjectServicesMappingConfig.IsSet() {
		return true
	}

	return false
}

// SetObjectServicesMappingConfig gets a reference to the given NullableString and assigns it to the ObjectServicesMappingConfig field.
func (o *ViewAllOf) SetObjectServicesMappingConfig(v string) {
	o.ObjectServicesMappingConfig.Set(&v)
}
// SetObjectServicesMappingConfigNil sets the value for ObjectServicesMappingConfig to be an explicit nil
func (o *ViewAllOf) SetObjectServicesMappingConfigNil() {
	o.ObjectServicesMappingConfig.Set(nil)
}

// UnsetObjectServicesMappingConfig ensures that no value is present for ObjectServicesMappingConfig, not even an explicit nil
func (o *ViewAllOf) UnsetObjectServicesMappingConfig() {
	o.ObjectServicesMappingConfig.Unset()
}

// GetStorageDomainId returns the StorageDomainId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ViewAllOf) GetStorageDomainId() int64 {
	if o == nil || o.StorageDomainId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.StorageDomainId.Get()
}

// GetStorageDomainIdOk returns a tuple with the StorageDomainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ViewAllOf) GetStorageDomainIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StorageDomainId.Get(), o.StorageDomainId.IsSet()
}

// HasStorageDomainId returns a boolean if a field has been set.
func (o *ViewAllOf) HasStorageDomainId() bool {
	if o != nil && o.StorageDomainId.IsSet() {
		return true
	}

	return false
}

// SetStorageDomainId gets a reference to the given NullableInt64 and assigns it to the StorageDomainId field.
func (o *ViewAllOf) SetStorageDomainId(v int64) {
	o.StorageDomainId.Set(&v)
}
// SetStorageDomainIdNil sets the value for StorageDomainId to be an explicit nil
func (o *ViewAllOf) SetStorageDomainIdNil() {
	o.StorageDomainId.Set(nil)
}

// UnsetStorageDomainId ensures that no value is present for StorageDomainId, not even an explicit nil
func (o *ViewAllOf) UnsetStorageDomainId() {
	o.StorageDomainId.Unset()
}

// GetStorageDomainName returns the StorageDomainName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ViewAllOf) GetStorageDomainName() string {
	if o == nil || o.StorageDomainName.Get() == nil {
		var ret string
		return ret
	}
	return *o.StorageDomainName.Get()
}

// GetStorageDomainNameOk returns a tuple with the StorageDomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ViewAllOf) GetStorageDomainNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StorageDomainName.Get(), o.StorageDomainName.IsSet()
}

// HasStorageDomainName returns a boolean if a field has been set.
func (o *ViewAllOf) HasStorageDomainName() bool {
	if o != nil && o.StorageDomainName.IsSet() {
		return true
	}

	return false
}

// SetStorageDomainName gets a reference to the given NullableString and assigns it to the StorageDomainName field.
func (o *ViewAllOf) SetStorageDomainName(v string) {
	o.StorageDomainName.Set(&v)
}
// SetStorageDomainNameNil sets the value for StorageDomainName to be an explicit nil
func (o *ViewAllOf) SetStorageDomainNameNil() {
	o.StorageDomainName.Set(nil)
}

// UnsetStorageDomainName ensures that no value is present for StorageDomainName, not even an explicit nil
func (o *ViewAllOf) UnsetStorageDomainName() {
	o.StorageDomainName.Unset()
}

// GetCaseInsensitiveNamesEnabled returns the CaseInsensitiveNamesEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ViewAllOf) GetCaseInsensitiveNamesEnabled() bool {
	if o == nil || o.CaseInsensitiveNamesEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.CaseInsensitiveNamesEnabled.Get()
}

// GetCaseInsensitiveNamesEnabledOk returns a tuple with the CaseInsensitiveNamesEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ViewAllOf) GetCaseInsensitiveNamesEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CaseInsensitiveNamesEnabled.Get(), o.CaseInsensitiveNamesEnabled.IsSet()
}

// HasCaseInsensitiveNamesEnabled returns a boolean if a field has been set.
func (o *ViewAllOf) HasCaseInsensitiveNamesEnabled() bool {
	if o != nil && o.CaseInsensitiveNamesEnabled.IsSet() {
		return true
	}

	return false
}

// SetCaseInsensitiveNamesEnabled gets a reference to the given NullableBool and assigns it to the CaseInsensitiveNamesEnabled field.
func (o *ViewAllOf) SetCaseInsensitiveNamesEnabled(v bool) {
	o.CaseInsensitiveNamesEnabled.Set(&v)
}
// SetCaseInsensitiveNamesEnabledNil sets the value for CaseInsensitiveNamesEnabled to be an explicit nil
func (o *ViewAllOf) SetCaseInsensitiveNamesEnabledNil() {
	o.CaseInsensitiveNamesEnabled.Set(nil)
}

// UnsetCaseInsensitiveNamesEnabled ensures that no value is present for CaseInsensitiveNamesEnabled, not even an explicit nil
func (o *ViewAllOf) UnsetCaseInsensitiveNamesEnabled() {
	o.CaseInsensitiveNamesEnabled.Unset()
}

// GetCreateTimeMsecs returns the CreateTimeMsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ViewAllOf) GetCreateTimeMsecs() int64 {
	if o == nil || o.CreateTimeMsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CreateTimeMsecs.Get()
}

// GetCreateTimeMsecsOk returns a tuple with the CreateTimeMsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ViewAllOf) GetCreateTimeMsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreateTimeMsecs.Get(), o.CreateTimeMsecs.IsSet()
}

// HasCreateTimeMsecs returns a boolean if a field has been set.
func (o *ViewAllOf) HasCreateTimeMsecs() bool {
	if o != nil && o.CreateTimeMsecs.IsSet() {
		return true
	}

	return false
}

// SetCreateTimeMsecs gets a reference to the given NullableInt64 and assigns it to the CreateTimeMsecs field.
func (o *ViewAllOf) SetCreateTimeMsecs(v int64) {
	o.CreateTimeMsecs.Set(&v)
}
// SetCreateTimeMsecsNil sets the value for CreateTimeMsecs to be an explicit nil
func (o *ViewAllOf) SetCreateTimeMsecsNil() {
	o.CreateTimeMsecs.Set(nil)
}

// UnsetCreateTimeMsecs ensures that no value is present for CreateTimeMsecs, not even an explicit nil
func (o *ViewAllOf) UnsetCreateTimeMsecs() {
	o.CreateTimeMsecs.Unset()
}

// GetBasicMountPath returns the BasicMountPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ViewAllOf) GetBasicMountPath() string {
	if o == nil || o.BasicMountPath.Get() == nil {
		var ret string
		return ret
	}
	return *o.BasicMountPath.Get()
}

// GetBasicMountPathOk returns a tuple with the BasicMountPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ViewAllOf) GetBasicMountPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BasicMountPath.Get(), o.BasicMountPath.IsSet()
}

// HasBasicMountPath returns a boolean if a field has been set.
func (o *ViewAllOf) HasBasicMountPath() bool {
	if o != nil && o.BasicMountPath.IsSet() {
		return true
	}

	return false
}

// SetBasicMountPath gets a reference to the given NullableString and assigns it to the BasicMountPath field.
func (o *ViewAllOf) SetBasicMountPath(v string) {
	o.BasicMountPath.Set(&v)
}
// SetBasicMountPathNil sets the value for BasicMountPath to be an explicit nil
func (o *ViewAllOf) SetBasicMountPathNil() {
	o.BasicMountPath.Set(nil)
}

// UnsetBasicMountPath ensures that no value is present for BasicMountPath, not even an explicit nil
func (o *ViewAllOf) UnsetBasicMountPath() {
	o.BasicMountPath.Unset()
}

// GetNfsMountPath returns the NfsMountPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ViewAllOf) GetNfsMountPath() string {
	if o == nil || o.NfsMountPath.Get() == nil {
		var ret string
		return ret
	}
	return *o.NfsMountPath.Get()
}

// GetNfsMountPathOk returns a tuple with the NfsMountPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ViewAllOf) GetNfsMountPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NfsMountPath.Get(), o.NfsMountPath.IsSet()
}

// HasNfsMountPath returns a boolean if a field has been set.
func (o *ViewAllOf) HasNfsMountPath() bool {
	if o != nil && o.NfsMountPath.IsSet() {
		return true
	}

	return false
}

// SetNfsMountPath gets a reference to the given NullableString and assigns it to the NfsMountPath field.
func (o *ViewAllOf) SetNfsMountPath(v string) {
	o.NfsMountPath.Set(&v)
}
// SetNfsMountPathNil sets the value for NfsMountPath to be an explicit nil
func (o *ViewAllOf) SetNfsMountPathNil() {
	o.NfsMountPath.Set(nil)
}

// UnsetNfsMountPath ensures that no value is present for NfsMountPath, not even an explicit nil
func (o *ViewAllOf) UnsetNfsMountPath() {
	o.NfsMountPath.Unset()
}

// GetNfsMountPaths returns the NfsMountPaths field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ViewAllOf) GetNfsMountPaths() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.NfsMountPaths
}

// GetNfsMountPathsOk returns a tuple with the NfsMountPaths field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ViewAllOf) GetNfsMountPathsOk() (*[]string, bool) {
	if o == nil || o.NfsMountPaths == nil {
		return nil, false
	}
	return &o.NfsMountPaths, true
}

// HasNfsMountPaths returns a boolean if a field has been set.
func (o *ViewAllOf) HasNfsMountPaths() bool {
	if o != nil && o.NfsMountPaths != nil {
		return true
	}

	return false
}

// SetNfsMountPaths gets a reference to the given []string and assigns it to the NfsMountPaths field.
func (o *ViewAllOf) SetNfsMountPaths(v []string) {
	o.NfsMountPaths = v
}

// GetSmbMountPaths returns the SmbMountPaths field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ViewAllOf) GetSmbMountPaths() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.SmbMountPaths
}

// GetSmbMountPathsOk returns a tuple with the SmbMountPaths field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ViewAllOf) GetSmbMountPathsOk() (*[]string, bool) {
	if o == nil || o.SmbMountPaths == nil {
		return nil, false
	}
	return &o.SmbMountPaths, true
}

// HasSmbMountPaths returns a boolean if a field has been set.
func (o *ViewAllOf) HasSmbMountPaths() bool {
	if o != nil && o.SmbMountPaths != nil {
		return true
	}

	return false
}

// SetSmbMountPaths gets a reference to the given []string and assigns it to the SmbMountPaths field.
func (o *ViewAllOf) SetSmbMountPaths(v []string) {
	o.SmbMountPaths = v
}

// GetViewProtection returns the ViewProtection field value if set, zero value otherwise.
func (o *ViewAllOf) GetViewProtection() ViewProtection {
	if o == nil || o.ViewProtection == nil {
		var ret ViewProtection
		return ret
	}
	return *o.ViewProtection
}

// GetViewProtectionOk returns a tuple with the ViewProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewAllOf) GetViewProtectionOk() (*ViewProtection, bool) {
	if o == nil || o.ViewProtection == nil {
		return nil, false
	}
	return o.ViewProtection, true
}

// HasViewProtection returns a boolean if a field has been set.
func (o *ViewAllOf) HasViewProtection() bool {
	if o != nil && o.ViewProtection != nil {
		return true
	}

	return false
}

// SetViewProtection gets a reference to the given ViewProtection and assigns it to the ViewProtection field.
func (o *ViewAllOf) SetViewProtection(v ViewProtection) {
	o.ViewProtection = &v
}

// GetAliases returns the Aliases field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ViewAllOf) GetAliases() []ViewAliasInfo {
	if o == nil  {
		var ret []ViewAliasInfo
		return ret
	}
	return o.Aliases
}

// GetAliasesOk returns a tuple with the Aliases field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ViewAllOf) GetAliasesOk() (*[]ViewAliasInfo, bool) {
	if o == nil || o.Aliases == nil {
		return nil, false
	}
	return &o.Aliases, true
}

// HasAliases returns a boolean if a field has been set.
func (o *ViewAllOf) HasAliases() bool {
	if o != nil && o.Aliases != nil {
		return true
	}

	return false
}

// SetAliases gets a reference to the given []ViewAliasInfo and assigns it to the Aliases field.
func (o *ViewAllOf) SetAliases(v []ViewAliasInfo) {
	o.Aliases = v
}

// GetIsTargetForMigratedData returns the IsTargetForMigratedData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ViewAllOf) GetIsTargetForMigratedData() bool {
	if o == nil || o.IsTargetForMigratedData.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsTargetForMigratedData.Get()
}

// GetIsTargetForMigratedDataOk returns a tuple with the IsTargetForMigratedData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ViewAllOf) GetIsTargetForMigratedDataOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsTargetForMigratedData.Get(), o.IsTargetForMigratedData.IsSet()
}

// HasIsTargetForMigratedData returns a boolean if a field has been set.
func (o *ViewAllOf) HasIsTargetForMigratedData() bool {
	if o != nil && o.IsTargetForMigratedData.IsSet() {
		return true
	}

	return false
}

// SetIsTargetForMigratedData gets a reference to the given NullableBool and assigns it to the IsTargetForMigratedData field.
func (o *ViewAllOf) SetIsTargetForMigratedData(v bool) {
	o.IsTargetForMigratedData.Set(&v)
}
// SetIsTargetForMigratedDataNil sets the value for IsTargetForMigratedData to be an explicit nil
func (o *ViewAllOf) SetIsTargetForMigratedDataNil() {
	o.IsTargetForMigratedData.Set(nil)
}

// UnsetIsTargetForMigratedData ensures that no value is present for IsTargetForMigratedData, not even an explicit nil
func (o *ViewAllOf) UnsetIsTargetForMigratedData() {
	o.IsTargetForMigratedData.Unset()
}

// GetViewFailover returns the ViewFailover field value if set, zero value otherwise.
func (o *ViewAllOf) GetViewFailover() ViewFailover {
	if o == nil || o.ViewFailover == nil {
		var ret ViewFailover
		return ret
	}
	return *o.ViewFailover
}

// GetViewFailoverOk returns a tuple with the ViewFailover field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewAllOf) GetViewFailoverOk() (*ViewFailover, bool) {
	if o == nil || o.ViewFailover == nil {
		return nil, false
	}
	return o.ViewFailover, true
}

// HasViewFailover returns a boolean if a field has been set.
func (o *ViewAllOf) HasViewFailover() bool {
	if o != nil && o.ViewFailover != nil {
		return true
	}

	return false
}

// SetViewFailover gets a reference to the given ViewFailover and assigns it to the ViewFailover field.
func (o *ViewAllOf) SetViewFailover(v ViewFailover) {
	o.ViewFailover = &v
}

// GetStats returns the Stats field value if set, zero value otherwise.
func (o *ViewAllOf) GetStats() ViewStats {
	if o == nil || o.Stats == nil {
		var ret ViewStats
		return ret
	}
	return *o.Stats
}

// GetStatsOk returns a tuple with the Stats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewAllOf) GetStatsOk() (*ViewStats, bool) {
	if o == nil || o.Stats == nil {
		return nil, false
	}
	return o.Stats, true
}

// HasStats returns a boolean if a field has been set.
func (o *ViewAllOf) HasStats() bool {
	if o != nil && o.Stats != nil {
		return true
	}

	return false
}

// SetStats gets a reference to the given ViewStats and assigns it to the Stats field.
func (o *ViewAllOf) SetStats(v ViewStats) {
	o.Stats = &v
}

// GetFileCountBySize returns the FileCountBySize field value if set, zero value otherwise.
func (o *ViewAllOf) GetFileCountBySize() []FileCount {
	if o == nil || o.FileCountBySize == nil {
		var ret []FileCount
		return ret
	}
	return *o.FileCountBySize
}

// GetFileCountBySizeOk returns a tuple with the FileCountBySize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewAllOf) GetFileCountBySizeOk() (*[]FileCount, bool) {
	if o == nil || o.FileCountBySize == nil {
		return nil, false
	}
	return o.FileCountBySize, true
}

// HasFileCountBySize returns a boolean if a field has been set.
func (o *ViewAllOf) HasFileCountBySize() bool {
	if o != nil && o.FileCountBySize != nil {
		return true
	}

	return false
}

// SetFileCountBySize gets a reference to the given []FileCount and assigns it to the FileCountBySize field.
func (o *ViewAllOf) SetFileCountBySize(v []FileCount) {
	o.FileCountBySize = &v
}

// GetOwnerSid returns the OwnerSid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ViewAllOf) GetOwnerSid() string {
	if o == nil || o.OwnerSid.Get() == nil {
		var ret string
		return ret
	}
	return *o.OwnerSid.Get()
}

// GetOwnerSidOk returns a tuple with the OwnerSid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ViewAllOf) GetOwnerSidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OwnerSid.Get(), o.OwnerSid.IsSet()
}

// HasOwnerSid returns a boolean if a field has been set.
func (o *ViewAllOf) HasOwnerSid() bool {
	if o != nil && o.OwnerSid.IsSet() {
		return true
	}

	return false
}

// SetOwnerSid gets a reference to the given NullableString and assigns it to the OwnerSid field.
func (o *ViewAllOf) SetOwnerSid(v string) {
	o.OwnerSid.Set(&v)
}
// SetOwnerSidNil sets the value for OwnerSid to be an explicit nil
func (o *ViewAllOf) SetOwnerSidNil() {
	o.OwnerSid.Set(nil)
}

// UnsetOwnerSid ensures that no value is present for OwnerSid, not even an explicit nil
func (o *ViewAllOf) UnsetOwnerSid() {
	o.OwnerSid.Unset()
}

func (o ViewAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ViewId.IsSet() {
		toSerialize["viewId"] = o.ViewId.Get()
	}
	if o.IsCategoryInferred.IsSet() {
		toSerialize["isCategoryInferred"] = o.IsCategoryInferred.Get()
	}
	if o.DataLockExpiryUsecs.IsSet() {
		toSerialize["dataLockExpiryUsecs"] = o.DataLockExpiryUsecs.Get()
	}
	if o.ObjectServicesMappingConfig.IsSet() {
		toSerialize["objectServicesMappingConfig"] = o.ObjectServicesMappingConfig.Get()
	}
	if o.StorageDomainId.IsSet() {
		toSerialize["storageDomainId"] = o.StorageDomainId.Get()
	}
	if o.StorageDomainName.IsSet() {
		toSerialize["storageDomainName"] = o.StorageDomainName.Get()
	}
	if o.CaseInsensitiveNamesEnabled.IsSet() {
		toSerialize["caseInsensitiveNamesEnabled"] = o.CaseInsensitiveNamesEnabled.Get()
	}
	if o.CreateTimeMsecs.IsSet() {
		toSerialize["createTimeMsecs"] = o.CreateTimeMsecs.Get()
	}
	if o.BasicMountPath.IsSet() {
		toSerialize["basicMountPath"] = o.BasicMountPath.Get()
	}
	if o.NfsMountPath.IsSet() {
		toSerialize["nfsMountPath"] = o.NfsMountPath.Get()
	}
	if o.NfsMountPaths != nil {
		toSerialize["nfsMountPaths"] = o.NfsMountPaths
	}
	if o.SmbMountPaths != nil {
		toSerialize["smbMountPaths"] = o.SmbMountPaths
	}
	if o.ViewProtection != nil {
		toSerialize["viewProtection"] = o.ViewProtection
	}
	if o.Aliases != nil {
		toSerialize["aliases"] = o.Aliases
	}
	if o.IsTargetForMigratedData.IsSet() {
		toSerialize["isTargetForMigratedData"] = o.IsTargetForMigratedData.Get()
	}
	if o.ViewFailover != nil {
		toSerialize["viewFailover"] = o.ViewFailover
	}
	if o.Stats != nil {
		toSerialize["stats"] = o.Stats
	}
	if o.FileCountBySize != nil {
		toSerialize["fileCountBySize"] = o.FileCountBySize
	}
	if o.OwnerSid.IsSet() {
		toSerialize["ownerSid"] = o.OwnerSid.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableViewAllOf struct {
	value *ViewAllOf
	isSet bool
}

func (v NullableViewAllOf) Get() *ViewAllOf {
	return v.value
}

func (v *NullableViewAllOf) Set(val *ViewAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableViewAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableViewAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViewAllOf(val *ViewAllOf) *NullableViewAllOf {
	return &NullableViewAllOf{value: val, isSet: true}
}

func (v NullableViewAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViewAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


