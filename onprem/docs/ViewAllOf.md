# ViewAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ViewId** | Pointer to **NullableInt64** | Specifies an id of the View assigned by the Cohesity Cluster. | [optional] [readonly] 
**IsCategoryInferred** | Pointer to **NullableBool** | If True, category in response is not set by user but inferred by Iris because none is set. Category can only be none when view was created by v1 API or cloned from a view created by v1 API.  Inference Logic is as follows: 1. Object Services if only S3 or Swift protocol is selected. 2. Backup Target only if one read-write protocol is selected and    QoS is \&quot;Backup Target Commvault\&quot; or \&quot;Backup Target SSD\&quot;. 3. File Services if there are more than 1 read-write protocol or    it doesn&#39;t fit any other category. | [optional] [readonly] 
**DataLockExpiryUsecs** | Pointer to **NullableInt64** | DataLock (Write Once Read Many) lock expiry epoch time in microseconds. If a view is marked as a DataLock view, only a Data Security Officer (a user having Data Security Privilege) can delete the view until the lock expiry time. | [optional] 
**ObjectServicesMappingConfig** | Pointer to **NullableString** | Specifies the Object Services key mapping config of the view. This parameter can only be set during create and cannot be changed. Configuration of Object Services key mapping. Specifies the type of Object Services key mapping config. | [optional] [readonly] 
**StorageDomainId** | Pointer to **NullableInt64** | Specifies the id of the Storage Domain (View Box) where the View is stored. | [optional] [readonly] 
**StorageDomainName** | Pointer to **NullableString** | Specifies the name of the Storage Domain (View Box) where the View is stored. | [optional] [readonly] 
**CaseInsensitiveNamesEnabled** | Pointer to **NullableBool** | Specifies whether to support case insensitive file/folder names. This parameter can only be set during create and cannot be changed. | [optional] [readonly] 
**CreateTimeMsecs** | Pointer to **NullableInt64** | Specifies the time that the View was created in milliseconds. | [optional] [readonly] 
**BasicMountPath** | Pointer to **NullableString** | Specifies the NFS mount path of the View (without the hostname information). This path is used to support NFS mounting of the paths specified in the nfsExportPathList on Windows systems. | [optional] [readonly] 
**NfsMountPath** | Pointer to **NullableString** | This field is currently deprecated. Please use NFS MountPaths which would be an array of strings. | [optional] [readonly] 
**NfsMountPaths** | Pointer to **[]string** | Array of NFS Paths. Specifies the path for mounting this View as an NFS share. If Kerberos Provider has multiple hostaliases, each host alias has  its own path. | [optional] [readonly] 
**SmbMountPaths** | Pointer to **[]string** | Array of SMB Paths. Specifies the possible paths that can be used to mount this View as a SMB share. If Active Directory has multiple account names; each machine account has its own path. | [optional] [readonly] 
**ViewProtection** | Pointer to [**ViewProtection**](ViewProtection.md) |  | [optional] 
**Aliases** | Pointer to [**[]ViewAliasInfo**](ViewAliasInfo.md) | Aliases created for the view. A view alias allows a directory path inside a view to be mounted using the alias name. | [optional] [readonly] 
**IsTargetForMigratedData** | Pointer to **NullableBool** | Specifies if a view contains migrated data. | [optional] [readonly] 
**ViewFailover** | Pointer to [**ViewFailover**](ViewFailover.md) | Specifies the information about the failover of the view. | [optional] 
**Stats** | Pointer to [**ViewStats**](ViewStats.md) | Specifies statistics about the View. | [optional] 
**FileCountBySize** | Pointer to [**[]FileCount**](FileCount.md) | Specifies the file count by size for the View. | [optional] 
**OwnerSid** | Pointer to **NullableString** | Specifies the sid of the view owner. | [optional] 

## Methods

### NewViewAllOf

`func NewViewAllOf() *ViewAllOf`

NewViewAllOf instantiates a new ViewAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewAllOfWithDefaults

`func NewViewAllOfWithDefaults() *ViewAllOf`

NewViewAllOfWithDefaults instantiates a new ViewAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetViewId

`func (o *ViewAllOf) GetViewId() int64`

GetViewId returns the ViewId field if non-nil, zero value otherwise.

### GetViewIdOk

`func (o *ViewAllOf) GetViewIdOk() (*int64, bool)`

GetViewIdOk returns a tuple with the ViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewId

`func (o *ViewAllOf) SetViewId(v int64)`

SetViewId sets ViewId field to given value.

### HasViewId

`func (o *ViewAllOf) HasViewId() bool`

HasViewId returns a boolean if a field has been set.

### SetViewIdNil

`func (o *ViewAllOf) SetViewIdNil(b bool)`

 SetViewIdNil sets the value for ViewId to be an explicit nil

### UnsetViewId
`func (o *ViewAllOf) UnsetViewId()`

UnsetViewId ensures that no value is present for ViewId, not even an explicit nil
### GetIsCategoryInferred

`func (o *ViewAllOf) GetIsCategoryInferred() bool`

GetIsCategoryInferred returns the IsCategoryInferred field if non-nil, zero value otherwise.

### GetIsCategoryInferredOk

`func (o *ViewAllOf) GetIsCategoryInferredOk() (*bool, bool)`

GetIsCategoryInferredOk returns a tuple with the IsCategoryInferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCategoryInferred

`func (o *ViewAllOf) SetIsCategoryInferred(v bool)`

SetIsCategoryInferred sets IsCategoryInferred field to given value.

### HasIsCategoryInferred

`func (o *ViewAllOf) HasIsCategoryInferred() bool`

HasIsCategoryInferred returns a boolean if a field has been set.

### SetIsCategoryInferredNil

`func (o *ViewAllOf) SetIsCategoryInferredNil(b bool)`

 SetIsCategoryInferredNil sets the value for IsCategoryInferred to be an explicit nil

### UnsetIsCategoryInferred
`func (o *ViewAllOf) UnsetIsCategoryInferred()`

UnsetIsCategoryInferred ensures that no value is present for IsCategoryInferred, not even an explicit nil
### GetDataLockExpiryUsecs

`func (o *ViewAllOf) GetDataLockExpiryUsecs() int64`

GetDataLockExpiryUsecs returns the DataLockExpiryUsecs field if non-nil, zero value otherwise.

### GetDataLockExpiryUsecsOk

`func (o *ViewAllOf) GetDataLockExpiryUsecsOk() (*int64, bool)`

GetDataLockExpiryUsecsOk returns a tuple with the DataLockExpiryUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLockExpiryUsecs

`func (o *ViewAllOf) SetDataLockExpiryUsecs(v int64)`

SetDataLockExpiryUsecs sets DataLockExpiryUsecs field to given value.

### HasDataLockExpiryUsecs

`func (o *ViewAllOf) HasDataLockExpiryUsecs() bool`

HasDataLockExpiryUsecs returns a boolean if a field has been set.

### SetDataLockExpiryUsecsNil

`func (o *ViewAllOf) SetDataLockExpiryUsecsNil(b bool)`

 SetDataLockExpiryUsecsNil sets the value for DataLockExpiryUsecs to be an explicit nil

### UnsetDataLockExpiryUsecs
`func (o *ViewAllOf) UnsetDataLockExpiryUsecs()`

UnsetDataLockExpiryUsecs ensures that no value is present for DataLockExpiryUsecs, not even an explicit nil
### GetObjectServicesMappingConfig

`func (o *ViewAllOf) GetObjectServicesMappingConfig() string`

GetObjectServicesMappingConfig returns the ObjectServicesMappingConfig field if non-nil, zero value otherwise.

### GetObjectServicesMappingConfigOk

`func (o *ViewAllOf) GetObjectServicesMappingConfigOk() (*string, bool)`

GetObjectServicesMappingConfigOk returns a tuple with the ObjectServicesMappingConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectServicesMappingConfig

`func (o *ViewAllOf) SetObjectServicesMappingConfig(v string)`

SetObjectServicesMappingConfig sets ObjectServicesMappingConfig field to given value.

### HasObjectServicesMappingConfig

`func (o *ViewAllOf) HasObjectServicesMappingConfig() bool`

HasObjectServicesMappingConfig returns a boolean if a field has been set.

### SetObjectServicesMappingConfigNil

`func (o *ViewAllOf) SetObjectServicesMappingConfigNil(b bool)`

 SetObjectServicesMappingConfigNil sets the value for ObjectServicesMappingConfig to be an explicit nil

### UnsetObjectServicesMappingConfig
`func (o *ViewAllOf) UnsetObjectServicesMappingConfig()`

UnsetObjectServicesMappingConfig ensures that no value is present for ObjectServicesMappingConfig, not even an explicit nil
### GetStorageDomainId

`func (o *ViewAllOf) GetStorageDomainId() int64`

GetStorageDomainId returns the StorageDomainId field if non-nil, zero value otherwise.

### GetStorageDomainIdOk

`func (o *ViewAllOf) GetStorageDomainIdOk() (*int64, bool)`

GetStorageDomainIdOk returns a tuple with the StorageDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainId

`func (o *ViewAllOf) SetStorageDomainId(v int64)`

SetStorageDomainId sets StorageDomainId field to given value.

### HasStorageDomainId

`func (o *ViewAllOf) HasStorageDomainId() bool`

HasStorageDomainId returns a boolean if a field has been set.

### SetStorageDomainIdNil

`func (o *ViewAllOf) SetStorageDomainIdNil(b bool)`

 SetStorageDomainIdNil sets the value for StorageDomainId to be an explicit nil

### UnsetStorageDomainId
`func (o *ViewAllOf) UnsetStorageDomainId()`

UnsetStorageDomainId ensures that no value is present for StorageDomainId, not even an explicit nil
### GetStorageDomainName

`func (o *ViewAllOf) GetStorageDomainName() string`

GetStorageDomainName returns the StorageDomainName field if non-nil, zero value otherwise.

### GetStorageDomainNameOk

`func (o *ViewAllOf) GetStorageDomainNameOk() (*string, bool)`

GetStorageDomainNameOk returns a tuple with the StorageDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainName

`func (o *ViewAllOf) SetStorageDomainName(v string)`

SetStorageDomainName sets StorageDomainName field to given value.

### HasStorageDomainName

`func (o *ViewAllOf) HasStorageDomainName() bool`

HasStorageDomainName returns a boolean if a field has been set.

### SetStorageDomainNameNil

`func (o *ViewAllOf) SetStorageDomainNameNil(b bool)`

 SetStorageDomainNameNil sets the value for StorageDomainName to be an explicit nil

### UnsetStorageDomainName
`func (o *ViewAllOf) UnsetStorageDomainName()`

UnsetStorageDomainName ensures that no value is present for StorageDomainName, not even an explicit nil
### GetCaseInsensitiveNamesEnabled

`func (o *ViewAllOf) GetCaseInsensitiveNamesEnabled() bool`

GetCaseInsensitiveNamesEnabled returns the CaseInsensitiveNamesEnabled field if non-nil, zero value otherwise.

### GetCaseInsensitiveNamesEnabledOk

`func (o *ViewAllOf) GetCaseInsensitiveNamesEnabledOk() (*bool, bool)`

GetCaseInsensitiveNamesEnabledOk returns a tuple with the CaseInsensitiveNamesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseInsensitiveNamesEnabled

`func (o *ViewAllOf) SetCaseInsensitiveNamesEnabled(v bool)`

SetCaseInsensitiveNamesEnabled sets CaseInsensitiveNamesEnabled field to given value.

### HasCaseInsensitiveNamesEnabled

`func (o *ViewAllOf) HasCaseInsensitiveNamesEnabled() bool`

HasCaseInsensitiveNamesEnabled returns a boolean if a field has been set.

### SetCaseInsensitiveNamesEnabledNil

`func (o *ViewAllOf) SetCaseInsensitiveNamesEnabledNil(b bool)`

 SetCaseInsensitiveNamesEnabledNil sets the value for CaseInsensitiveNamesEnabled to be an explicit nil

### UnsetCaseInsensitiveNamesEnabled
`func (o *ViewAllOf) UnsetCaseInsensitiveNamesEnabled()`

UnsetCaseInsensitiveNamesEnabled ensures that no value is present for CaseInsensitiveNamesEnabled, not even an explicit nil
### GetCreateTimeMsecs

`func (o *ViewAllOf) GetCreateTimeMsecs() int64`

GetCreateTimeMsecs returns the CreateTimeMsecs field if non-nil, zero value otherwise.

### GetCreateTimeMsecsOk

`func (o *ViewAllOf) GetCreateTimeMsecsOk() (*int64, bool)`

GetCreateTimeMsecsOk returns a tuple with the CreateTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimeMsecs

`func (o *ViewAllOf) SetCreateTimeMsecs(v int64)`

SetCreateTimeMsecs sets CreateTimeMsecs field to given value.

### HasCreateTimeMsecs

`func (o *ViewAllOf) HasCreateTimeMsecs() bool`

HasCreateTimeMsecs returns a boolean if a field has been set.

### SetCreateTimeMsecsNil

`func (o *ViewAllOf) SetCreateTimeMsecsNil(b bool)`

 SetCreateTimeMsecsNil sets the value for CreateTimeMsecs to be an explicit nil

### UnsetCreateTimeMsecs
`func (o *ViewAllOf) UnsetCreateTimeMsecs()`

UnsetCreateTimeMsecs ensures that no value is present for CreateTimeMsecs, not even an explicit nil
### GetBasicMountPath

`func (o *ViewAllOf) GetBasicMountPath() string`

GetBasicMountPath returns the BasicMountPath field if non-nil, zero value otherwise.

### GetBasicMountPathOk

`func (o *ViewAllOf) GetBasicMountPathOk() (*string, bool)`

GetBasicMountPathOk returns a tuple with the BasicMountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicMountPath

`func (o *ViewAllOf) SetBasicMountPath(v string)`

SetBasicMountPath sets BasicMountPath field to given value.

### HasBasicMountPath

`func (o *ViewAllOf) HasBasicMountPath() bool`

HasBasicMountPath returns a boolean if a field has been set.

### SetBasicMountPathNil

`func (o *ViewAllOf) SetBasicMountPathNil(b bool)`

 SetBasicMountPathNil sets the value for BasicMountPath to be an explicit nil

### UnsetBasicMountPath
`func (o *ViewAllOf) UnsetBasicMountPath()`

UnsetBasicMountPath ensures that no value is present for BasicMountPath, not even an explicit nil
### GetNfsMountPath

`func (o *ViewAllOf) GetNfsMountPath() string`

GetNfsMountPath returns the NfsMountPath field if non-nil, zero value otherwise.

### GetNfsMountPathOk

`func (o *ViewAllOf) GetNfsMountPathOk() (*string, bool)`

GetNfsMountPathOk returns a tuple with the NfsMountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsMountPath

`func (o *ViewAllOf) SetNfsMountPath(v string)`

SetNfsMountPath sets NfsMountPath field to given value.

### HasNfsMountPath

`func (o *ViewAllOf) HasNfsMountPath() bool`

HasNfsMountPath returns a boolean if a field has been set.

### SetNfsMountPathNil

`func (o *ViewAllOf) SetNfsMountPathNil(b bool)`

 SetNfsMountPathNil sets the value for NfsMountPath to be an explicit nil

### UnsetNfsMountPath
`func (o *ViewAllOf) UnsetNfsMountPath()`

UnsetNfsMountPath ensures that no value is present for NfsMountPath, not even an explicit nil
### GetNfsMountPaths

`func (o *ViewAllOf) GetNfsMountPaths() []string`

GetNfsMountPaths returns the NfsMountPaths field if non-nil, zero value otherwise.

### GetNfsMountPathsOk

`func (o *ViewAllOf) GetNfsMountPathsOk() (*[]string, bool)`

GetNfsMountPathsOk returns a tuple with the NfsMountPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsMountPaths

`func (o *ViewAllOf) SetNfsMountPaths(v []string)`

SetNfsMountPaths sets NfsMountPaths field to given value.

### HasNfsMountPaths

`func (o *ViewAllOf) HasNfsMountPaths() bool`

HasNfsMountPaths returns a boolean if a field has been set.

### SetNfsMountPathsNil

`func (o *ViewAllOf) SetNfsMountPathsNil(b bool)`

 SetNfsMountPathsNil sets the value for NfsMountPaths to be an explicit nil

### UnsetNfsMountPaths
`func (o *ViewAllOf) UnsetNfsMountPaths()`

UnsetNfsMountPaths ensures that no value is present for NfsMountPaths, not even an explicit nil
### GetSmbMountPaths

`func (o *ViewAllOf) GetSmbMountPaths() []string`

GetSmbMountPaths returns the SmbMountPaths field if non-nil, zero value otherwise.

### GetSmbMountPathsOk

`func (o *ViewAllOf) GetSmbMountPathsOk() (*[]string, bool)`

GetSmbMountPathsOk returns a tuple with the SmbMountPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbMountPaths

`func (o *ViewAllOf) SetSmbMountPaths(v []string)`

SetSmbMountPaths sets SmbMountPaths field to given value.

### HasSmbMountPaths

`func (o *ViewAllOf) HasSmbMountPaths() bool`

HasSmbMountPaths returns a boolean if a field has been set.

### SetSmbMountPathsNil

`func (o *ViewAllOf) SetSmbMountPathsNil(b bool)`

 SetSmbMountPathsNil sets the value for SmbMountPaths to be an explicit nil

### UnsetSmbMountPaths
`func (o *ViewAllOf) UnsetSmbMountPaths()`

UnsetSmbMountPaths ensures that no value is present for SmbMountPaths, not even an explicit nil
### GetViewProtection

`func (o *ViewAllOf) GetViewProtection() ViewProtection`

GetViewProtection returns the ViewProtection field if non-nil, zero value otherwise.

### GetViewProtectionOk

`func (o *ViewAllOf) GetViewProtectionOk() (*ViewProtection, bool)`

GetViewProtectionOk returns a tuple with the ViewProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewProtection

`func (o *ViewAllOf) SetViewProtection(v ViewProtection)`

SetViewProtection sets ViewProtection field to given value.

### HasViewProtection

`func (o *ViewAllOf) HasViewProtection() bool`

HasViewProtection returns a boolean if a field has been set.

### GetAliases

`func (o *ViewAllOf) GetAliases() []ViewAliasInfo`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *ViewAllOf) GetAliasesOk() (*[]ViewAliasInfo, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *ViewAllOf) SetAliases(v []ViewAliasInfo)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *ViewAllOf) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### SetAliasesNil

`func (o *ViewAllOf) SetAliasesNil(b bool)`

 SetAliasesNil sets the value for Aliases to be an explicit nil

### UnsetAliases
`func (o *ViewAllOf) UnsetAliases()`

UnsetAliases ensures that no value is present for Aliases, not even an explicit nil
### GetIsTargetForMigratedData

`func (o *ViewAllOf) GetIsTargetForMigratedData() bool`

GetIsTargetForMigratedData returns the IsTargetForMigratedData field if non-nil, zero value otherwise.

### GetIsTargetForMigratedDataOk

`func (o *ViewAllOf) GetIsTargetForMigratedDataOk() (*bool, bool)`

GetIsTargetForMigratedDataOk returns a tuple with the IsTargetForMigratedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTargetForMigratedData

`func (o *ViewAllOf) SetIsTargetForMigratedData(v bool)`

SetIsTargetForMigratedData sets IsTargetForMigratedData field to given value.

### HasIsTargetForMigratedData

`func (o *ViewAllOf) HasIsTargetForMigratedData() bool`

HasIsTargetForMigratedData returns a boolean if a field has been set.

### SetIsTargetForMigratedDataNil

`func (o *ViewAllOf) SetIsTargetForMigratedDataNil(b bool)`

 SetIsTargetForMigratedDataNil sets the value for IsTargetForMigratedData to be an explicit nil

### UnsetIsTargetForMigratedData
`func (o *ViewAllOf) UnsetIsTargetForMigratedData()`

UnsetIsTargetForMigratedData ensures that no value is present for IsTargetForMigratedData, not even an explicit nil
### GetViewFailover

`func (o *ViewAllOf) GetViewFailover() ViewFailover`

GetViewFailover returns the ViewFailover field if non-nil, zero value otherwise.

### GetViewFailoverOk

`func (o *ViewAllOf) GetViewFailoverOk() (*ViewFailover, bool)`

GetViewFailoverOk returns a tuple with the ViewFailover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewFailover

`func (o *ViewAllOf) SetViewFailover(v ViewFailover)`

SetViewFailover sets ViewFailover field to given value.

### HasViewFailover

`func (o *ViewAllOf) HasViewFailover() bool`

HasViewFailover returns a boolean if a field has been set.

### GetStats

`func (o *ViewAllOf) GetStats() ViewStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ViewAllOf) GetStatsOk() (*ViewStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ViewAllOf) SetStats(v ViewStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ViewAllOf) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetFileCountBySize

`func (o *ViewAllOf) GetFileCountBySize() []FileCount`

GetFileCountBySize returns the FileCountBySize field if non-nil, zero value otherwise.

### GetFileCountBySizeOk

`func (o *ViewAllOf) GetFileCountBySizeOk() (*[]FileCount, bool)`

GetFileCountBySizeOk returns a tuple with the FileCountBySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileCountBySize

`func (o *ViewAllOf) SetFileCountBySize(v []FileCount)`

SetFileCountBySize sets FileCountBySize field to given value.

### HasFileCountBySize

`func (o *ViewAllOf) HasFileCountBySize() bool`

HasFileCountBySize returns a boolean if a field has been set.

### GetOwnerSid

`func (o *ViewAllOf) GetOwnerSid() string`

GetOwnerSid returns the OwnerSid field if non-nil, zero value otherwise.

### GetOwnerSidOk

`func (o *ViewAllOf) GetOwnerSidOk() (*string, bool)`

GetOwnerSidOk returns a tuple with the OwnerSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerSid

`func (o *ViewAllOf) SetOwnerSid(v string)`

SetOwnerSid sets OwnerSid field to given value.

### HasOwnerSid

`func (o *ViewAllOf) HasOwnerSid() bool`

HasOwnerSid returns a boolean if a field has been set.

### SetOwnerSidNil

`func (o *ViewAllOf) SetOwnerSidNil(b bool)`

 SetOwnerSidNil sets the value for OwnerSid to be an explicit nil

### UnsetOwnerSid
`func (o *ViewAllOf) UnsetOwnerSid()`

UnsetOwnerSid ensures that no value is present for OwnerSid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


