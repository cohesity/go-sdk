# View

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessSids** | Pointer to **[]string** | Array of Security Identifiers (SIDs)  Specifies the list of security identifiers (SIDs) for the restricted Principals who have access to this View. | [optional] 
**Aliases** | Pointer to [**[]ViewAlias**](ViewAlias.md) | Aliases created for the view. A view alias allows a directory path inside a view to be mounted using the alias name. | [optional] 
**AllSmbMountPaths** | Pointer to **[]string** | Array of SMB Paths.  Specifies the possible paths that can be used to mount this View as a SMB share. If Active Directory has multiple account names; each machine account has its own path. | [optional] 
**AntivirusScanConfig** | Pointer to [**AntivirusScanConfig**](AntivirusScanConfig.md) |  | [optional] 
**BasicMountPath** | Pointer to **NullableString** | Specifies the NFS mount path of the View (without the hostname information). This path is used to support NFS mounting of the paths specified in the nfsExportPathList on Windows systems. | [optional] 
**CaseInsensitiveNamesEnabled** | Pointer to **NullableBool** | Specifies whether to support case insensitive file/folder names. This parameter can only be set during create and cannot be changed. | [optional] 
**CreateTimeMsecs** | Pointer to **NullableInt64** | Specifies the time that the View was created in milliseconds. | [optional] 
**DataLockExpiryUsecs** | Pointer to **NullableInt64** | DataLock (Write Once Read Many) lock expiry epoch time in microseconds. If a view is marked as a DataLock view, only a Data Security Officer (a user having Data Security Privilege) can delete the view until the lock expiry time. | [optional] 
**Description** | Pointer to **NullableString** | Specifies an optional text description about the View. | [optional] 
**EnableFastDurableHandle** | Pointer to **NullableBool** | Specifies whether fast durable handle is enabled. If enabled, view open handle will be kept in memory, which results in a higher performance. But the handles cannot be recovered if node or service crashes. | [optional] 
**EnableFilerAuditLogging** | Pointer to **NullableBool** | Specifies if Filer Audit Logging is enabled for this view. | [optional] 
**EnableLiveIndexing** | Pointer to **NullableBool** | Specifies whether to enable live indexing for the view. | [optional] 
**EnableMixedModePermissions** | Pointer to **NullableBool** | If set, mixed mode (NFS and SMB) access is enabled for this view. This field is deprecated. Use the field SecurityMode. deprecated: true | [optional] 
**EnableNfsViewDiscovery** | Pointer to **NullableBool** | If set, it enables discovery of view for NFS. | [optional] 
**EnableOfflineCaching** | Pointer to **NullableBool** | Specifies whether to enable offline file caching of the view. | [optional] 
**EnableSmbAccessBasedEnumeration** | Pointer to **NullableBool** | Specifies if access-based enumeration should be enabled. If &#39;true&#39;, only files and folders that the user has permissions to access are visible on the SMB share for that user. | [optional] 
**EnableSmbEncryption** | Pointer to **NullableBool** | Specifies the SMB encryption for the View. If set, it enables the SMB encryption for the View. Encryption is supported only by SMB 3.x dialects. Dialects that do not support would still access data in unencrypted format. | [optional] 
**EnableSmbOplock** | Pointer to **NullableBool** | Specifies whether SMB opportunistic lock is enabled. | [optional] 
**EnableSmbViewDiscovery** | Pointer to **NullableBool** | If set, it enables discovery of view for SMB. | [optional] 
**EnforceSmbEncryption** | Pointer to **NullableBool** | Specifies the SMB encryption for all the sessions for the View. If set, encryption is enforced for all the sessions for the View. When enabled all future and existing unencrypted sessions are disallowed. | [optional] 
**FileExtensionFilter** | Pointer to [**FileExtensionFilter**](FileExtensionFilter.md) |  | [optional] 
**FileLockConfig** | Pointer to [**FileLevelDataLockConfig**](FileLevelDataLockConfig.md) |  | [optional] 
**IsTargetForMigratedData** | Pointer to **NullableBool** | Specifies if a view contains migrated data. | [optional] 
**LogicalQuota** | Pointer to [**NullableQuotaPolicy**](QuotaPolicy.md) | Specifies an optional logical quota limit (in bytes) for the usage allowed on this View. (Logical data is when the data is fully hydrated and expanded.) This limit overrides the limit inherited from the Storage Domain (View Box) (if set). If logicalQuota is nil, the limit is inherited from the Storage Domain (View Box) (if set). A new write is not allowed if the Storage Domain (View Box) will exceed the specified quota. However, it takes time for the Cohesity Cluster to calculate the usage across Nodes, so the limit may be exceeded by a small amount. In addition, if the limit is increased or data is removed, there may be a delay before the Cohesity Cluster allows more data to be written to the View, as the Cluster is calculating the usage across Nodes. | [optional] 
**LogicalUsageBytes** | Pointer to **NullableInt64** | LogicalUsageBytes is the logical usage in bytes for the view. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the View. | [optional] 
**NetgroupWhitelist** | Pointer to [**[]NisNetgroup**](NisNetgroup.md) | Array of Netgroups.  Specifies a list of Netgroups that have permissions to access the View. (Overrides the Netgroups specified at the global Cohesity Cluster level.) | [optional] 
**NfsAllSquash** | Pointer to [**NfsSquash**](NfsSquash.md) |  | [optional] 
**NfsMountPath** | Pointer to **NullableString** | Specifies the path for mounting this View as an NFS share. | [optional] 
**NfsRootPermissions** | Pointer to [**NfsRootPermissions**](NfsRootPermissions.md) |  | [optional] 
**NfsRootSquash** | Pointer to [**NfsSquash**](NfsSquash.md) |  | [optional] 
**OverrideGlobalNetgroupWhitelist** | Pointer to **NullableBool** | Specifies whether view level client netgroup whitelist overrides cluster and global setting. | [optional] 
**OverrideGlobalWhitelist** | Pointer to **NullableBool** | Specifies whether view level client subnet whitelist overrides cluster and global setting. | [optional] 
**ProtocolAccess** | Pointer to **NullableString** | Specifies the supported Protocols for the View. &#39;kAll&#39; enables protocol access to following three views: NFS, SMB and S3. &#39;kNFSOnly&#39; enables protocol access to NFS only. &#39;kSMBOnly&#39; enables protocol access to SMB only. &#39;kS3Only&#39; enables protocol access to S3 only. &#39;kSwiftOnly&#39; enables protocol access to Swift only. | [optional] 
**Qos** | Pointer to [**QoS**](QoS.md) |  | [optional] 
**S3AccessPath** | Pointer to **NullableString** | Specifies the path to access this View as an S3 share. | [optional] 
**S3KeyMappingConfig** | Pointer to **NullableString** | Specifies the S3 key mapping config of the view. This parameter can only be set during create and cannot be changed. Configuration of S3 key mapping.  Specifies the type of S3 key mapping config. | [optional] 
**SecurityMode** | Pointer to **NullableString** | Specifies the security mode used for this view. Currently we support the following modes: Native, Unified and NTFS style. &#39;kNativeMode&#39; indicates a native security mode. &#39;kUnifiedMode&#39; indicates a unified security mode. &#39;kNtfsMode&#39; indicates a NTFS style security mode. | [optional] 
**SharePermissions** | Pointer to [**[]SmbPermission**](SmbPermission.md) | Specifies a list of share level permissions. | [optional] 
**SmbMountPath** | Pointer to **NullableString** | Specifies the main path for mounting this View as an SMB share. | [optional] 
**SmbPermissionsInfo** | Pointer to [**SmbPermissionsInfo**](SmbPermissionsInfo.md) |  | [optional] 
**Stats** | Pointer to [**ViewStats**](ViewStats.md) |  | [optional] 
**StoragePolicyOverride** | Pointer to [**StoragePolicyOverride**](StoragePolicyOverride.md) |  | [optional] 
**SubnetWhitelist** | Pointer to [**[]Subnet**](Subnet.md) | Array of Subnets.  Specifies a list of Subnets with IP addresses that have permissions to access the View. (Overrides the Subnets specified at the global Cohesity Cluster level.) | [optional] 
**SwiftProjectDomain** | Pointer to **NullableString** | Specifies the Keystone project domain. | [optional] 
**SwiftProjectName** | Pointer to **NullableString** | Specifies the Keystone project name. | [optional] 
**SwiftUserDomain** | Pointer to **NullableString** | Specifies the Keystone user domain. | [optional] 
**SwiftUsername** | Pointer to **NullableString** | Specifies the Keystone username. | [optional] 
**TenantId** | Pointer to **NullableString** | Optional tenant id who has access to this View. | [optional] 
**ViewBoxId** | Pointer to **NullableInt64** | Specifies the id of the Storage Domain (View Box) where the View is stored. | [optional] 
**ViewBoxName** | Pointer to **NullableString** | Specifies the name of the Storage Domain (View Box) where the View is stored. | [optional] 
**ViewId** | Pointer to **NullableInt64** | Specifies an id of the View assigned by the Cohesity Cluster. | [optional] 
**ViewLockEnabled** | Pointer to **NullableBool** | Specifies whether view lock is enabled. If enabled the view cannot be modified or deleted until unlock. By default it is disabled. | [optional] 
**ViewProtection** | Pointer to [**ViewProtection**](ViewProtection.md) |  | [optional] 

## Methods

### NewView

`func NewView() *View`

NewView instantiates a new View object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewWithDefaults

`func NewViewWithDefaults() *View`

NewViewWithDefaults instantiates a new View object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessSids

`func (o *View) GetAccessSids() []string`

GetAccessSids returns the AccessSids field if non-nil, zero value otherwise.

### GetAccessSidsOk

`func (o *View) GetAccessSidsOk() (*[]string, bool)`

GetAccessSidsOk returns a tuple with the AccessSids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessSids

`func (o *View) SetAccessSids(v []string)`

SetAccessSids sets AccessSids field to given value.

### HasAccessSids

`func (o *View) HasAccessSids() bool`

HasAccessSids returns a boolean if a field has been set.

### SetAccessSidsNil

`func (o *View) SetAccessSidsNil(b bool)`

 SetAccessSidsNil sets the value for AccessSids to be an explicit nil

### UnsetAccessSids
`func (o *View) UnsetAccessSids()`

UnsetAccessSids ensures that no value is present for AccessSids, not even an explicit nil
### GetAliases

`func (o *View) GetAliases() []ViewAlias`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *View) GetAliasesOk() (*[]ViewAlias, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *View) SetAliases(v []ViewAlias)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *View) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### SetAliasesNil

`func (o *View) SetAliasesNil(b bool)`

 SetAliasesNil sets the value for Aliases to be an explicit nil

### UnsetAliases
`func (o *View) UnsetAliases()`

UnsetAliases ensures that no value is present for Aliases, not even an explicit nil
### GetAllSmbMountPaths

`func (o *View) GetAllSmbMountPaths() []string`

GetAllSmbMountPaths returns the AllSmbMountPaths field if non-nil, zero value otherwise.

### GetAllSmbMountPathsOk

`func (o *View) GetAllSmbMountPathsOk() (*[]string, bool)`

GetAllSmbMountPathsOk returns a tuple with the AllSmbMountPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllSmbMountPaths

`func (o *View) SetAllSmbMountPaths(v []string)`

SetAllSmbMountPaths sets AllSmbMountPaths field to given value.

### HasAllSmbMountPaths

`func (o *View) HasAllSmbMountPaths() bool`

HasAllSmbMountPaths returns a boolean if a field has been set.

### SetAllSmbMountPathsNil

`func (o *View) SetAllSmbMountPathsNil(b bool)`

 SetAllSmbMountPathsNil sets the value for AllSmbMountPaths to be an explicit nil

### UnsetAllSmbMountPaths
`func (o *View) UnsetAllSmbMountPaths()`

UnsetAllSmbMountPaths ensures that no value is present for AllSmbMountPaths, not even an explicit nil
### GetAntivirusScanConfig

`func (o *View) GetAntivirusScanConfig() AntivirusScanConfig`

GetAntivirusScanConfig returns the AntivirusScanConfig field if non-nil, zero value otherwise.

### GetAntivirusScanConfigOk

`func (o *View) GetAntivirusScanConfigOk() (*AntivirusScanConfig, bool)`

GetAntivirusScanConfigOk returns a tuple with the AntivirusScanConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntivirusScanConfig

`func (o *View) SetAntivirusScanConfig(v AntivirusScanConfig)`

SetAntivirusScanConfig sets AntivirusScanConfig field to given value.

### HasAntivirusScanConfig

`func (o *View) HasAntivirusScanConfig() bool`

HasAntivirusScanConfig returns a boolean if a field has been set.

### GetBasicMountPath

`func (o *View) GetBasicMountPath() string`

GetBasicMountPath returns the BasicMountPath field if non-nil, zero value otherwise.

### GetBasicMountPathOk

`func (o *View) GetBasicMountPathOk() (*string, bool)`

GetBasicMountPathOk returns a tuple with the BasicMountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicMountPath

`func (o *View) SetBasicMountPath(v string)`

SetBasicMountPath sets BasicMountPath field to given value.

### HasBasicMountPath

`func (o *View) HasBasicMountPath() bool`

HasBasicMountPath returns a boolean if a field has been set.

### SetBasicMountPathNil

`func (o *View) SetBasicMountPathNil(b bool)`

 SetBasicMountPathNil sets the value for BasicMountPath to be an explicit nil

### UnsetBasicMountPath
`func (o *View) UnsetBasicMountPath()`

UnsetBasicMountPath ensures that no value is present for BasicMountPath, not even an explicit nil
### GetCaseInsensitiveNamesEnabled

`func (o *View) GetCaseInsensitiveNamesEnabled() bool`

GetCaseInsensitiveNamesEnabled returns the CaseInsensitiveNamesEnabled field if non-nil, zero value otherwise.

### GetCaseInsensitiveNamesEnabledOk

`func (o *View) GetCaseInsensitiveNamesEnabledOk() (*bool, bool)`

GetCaseInsensitiveNamesEnabledOk returns a tuple with the CaseInsensitiveNamesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseInsensitiveNamesEnabled

`func (o *View) SetCaseInsensitiveNamesEnabled(v bool)`

SetCaseInsensitiveNamesEnabled sets CaseInsensitiveNamesEnabled field to given value.

### HasCaseInsensitiveNamesEnabled

`func (o *View) HasCaseInsensitiveNamesEnabled() bool`

HasCaseInsensitiveNamesEnabled returns a boolean if a field has been set.

### SetCaseInsensitiveNamesEnabledNil

`func (o *View) SetCaseInsensitiveNamesEnabledNil(b bool)`

 SetCaseInsensitiveNamesEnabledNil sets the value for CaseInsensitiveNamesEnabled to be an explicit nil

### UnsetCaseInsensitiveNamesEnabled
`func (o *View) UnsetCaseInsensitiveNamesEnabled()`

UnsetCaseInsensitiveNamesEnabled ensures that no value is present for CaseInsensitiveNamesEnabled, not even an explicit nil
### GetCreateTimeMsecs

`func (o *View) GetCreateTimeMsecs() int64`

GetCreateTimeMsecs returns the CreateTimeMsecs field if non-nil, zero value otherwise.

### GetCreateTimeMsecsOk

`func (o *View) GetCreateTimeMsecsOk() (*int64, bool)`

GetCreateTimeMsecsOk returns a tuple with the CreateTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimeMsecs

`func (o *View) SetCreateTimeMsecs(v int64)`

SetCreateTimeMsecs sets CreateTimeMsecs field to given value.

### HasCreateTimeMsecs

`func (o *View) HasCreateTimeMsecs() bool`

HasCreateTimeMsecs returns a boolean if a field has been set.

### SetCreateTimeMsecsNil

`func (o *View) SetCreateTimeMsecsNil(b bool)`

 SetCreateTimeMsecsNil sets the value for CreateTimeMsecs to be an explicit nil

### UnsetCreateTimeMsecs
`func (o *View) UnsetCreateTimeMsecs()`

UnsetCreateTimeMsecs ensures that no value is present for CreateTimeMsecs, not even an explicit nil
### GetDataLockExpiryUsecs

`func (o *View) GetDataLockExpiryUsecs() int64`

GetDataLockExpiryUsecs returns the DataLockExpiryUsecs field if non-nil, zero value otherwise.

### GetDataLockExpiryUsecsOk

`func (o *View) GetDataLockExpiryUsecsOk() (*int64, bool)`

GetDataLockExpiryUsecsOk returns a tuple with the DataLockExpiryUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLockExpiryUsecs

`func (o *View) SetDataLockExpiryUsecs(v int64)`

SetDataLockExpiryUsecs sets DataLockExpiryUsecs field to given value.

### HasDataLockExpiryUsecs

`func (o *View) HasDataLockExpiryUsecs() bool`

HasDataLockExpiryUsecs returns a boolean if a field has been set.

### SetDataLockExpiryUsecsNil

`func (o *View) SetDataLockExpiryUsecsNil(b bool)`

 SetDataLockExpiryUsecsNil sets the value for DataLockExpiryUsecs to be an explicit nil

### UnsetDataLockExpiryUsecs
`func (o *View) UnsetDataLockExpiryUsecs()`

UnsetDataLockExpiryUsecs ensures that no value is present for DataLockExpiryUsecs, not even an explicit nil
### GetDescription

`func (o *View) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *View) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *View) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *View) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *View) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *View) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnableFastDurableHandle

`func (o *View) GetEnableFastDurableHandle() bool`

GetEnableFastDurableHandle returns the EnableFastDurableHandle field if non-nil, zero value otherwise.

### GetEnableFastDurableHandleOk

`func (o *View) GetEnableFastDurableHandleOk() (*bool, bool)`

GetEnableFastDurableHandleOk returns a tuple with the EnableFastDurableHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFastDurableHandle

`func (o *View) SetEnableFastDurableHandle(v bool)`

SetEnableFastDurableHandle sets EnableFastDurableHandle field to given value.

### HasEnableFastDurableHandle

`func (o *View) HasEnableFastDurableHandle() bool`

HasEnableFastDurableHandle returns a boolean if a field has been set.

### SetEnableFastDurableHandleNil

`func (o *View) SetEnableFastDurableHandleNil(b bool)`

 SetEnableFastDurableHandleNil sets the value for EnableFastDurableHandle to be an explicit nil

### UnsetEnableFastDurableHandle
`func (o *View) UnsetEnableFastDurableHandle()`

UnsetEnableFastDurableHandle ensures that no value is present for EnableFastDurableHandle, not even an explicit nil
### GetEnableFilerAuditLogging

`func (o *View) GetEnableFilerAuditLogging() bool`

GetEnableFilerAuditLogging returns the EnableFilerAuditLogging field if non-nil, zero value otherwise.

### GetEnableFilerAuditLoggingOk

`func (o *View) GetEnableFilerAuditLoggingOk() (*bool, bool)`

GetEnableFilerAuditLoggingOk returns a tuple with the EnableFilerAuditLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFilerAuditLogging

`func (o *View) SetEnableFilerAuditLogging(v bool)`

SetEnableFilerAuditLogging sets EnableFilerAuditLogging field to given value.

### HasEnableFilerAuditLogging

`func (o *View) HasEnableFilerAuditLogging() bool`

HasEnableFilerAuditLogging returns a boolean if a field has been set.

### SetEnableFilerAuditLoggingNil

`func (o *View) SetEnableFilerAuditLoggingNil(b bool)`

 SetEnableFilerAuditLoggingNil sets the value for EnableFilerAuditLogging to be an explicit nil

### UnsetEnableFilerAuditLogging
`func (o *View) UnsetEnableFilerAuditLogging()`

UnsetEnableFilerAuditLogging ensures that no value is present for EnableFilerAuditLogging, not even an explicit nil
### GetEnableLiveIndexing

`func (o *View) GetEnableLiveIndexing() bool`

GetEnableLiveIndexing returns the EnableLiveIndexing field if non-nil, zero value otherwise.

### GetEnableLiveIndexingOk

`func (o *View) GetEnableLiveIndexingOk() (*bool, bool)`

GetEnableLiveIndexingOk returns a tuple with the EnableLiveIndexing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLiveIndexing

`func (o *View) SetEnableLiveIndexing(v bool)`

SetEnableLiveIndexing sets EnableLiveIndexing field to given value.

### HasEnableLiveIndexing

`func (o *View) HasEnableLiveIndexing() bool`

HasEnableLiveIndexing returns a boolean if a field has been set.

### SetEnableLiveIndexingNil

`func (o *View) SetEnableLiveIndexingNil(b bool)`

 SetEnableLiveIndexingNil sets the value for EnableLiveIndexing to be an explicit nil

### UnsetEnableLiveIndexing
`func (o *View) UnsetEnableLiveIndexing()`

UnsetEnableLiveIndexing ensures that no value is present for EnableLiveIndexing, not even an explicit nil
### GetEnableMixedModePermissions

`func (o *View) GetEnableMixedModePermissions() bool`

GetEnableMixedModePermissions returns the EnableMixedModePermissions field if non-nil, zero value otherwise.

### GetEnableMixedModePermissionsOk

`func (o *View) GetEnableMixedModePermissionsOk() (*bool, bool)`

GetEnableMixedModePermissionsOk returns a tuple with the EnableMixedModePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMixedModePermissions

`func (o *View) SetEnableMixedModePermissions(v bool)`

SetEnableMixedModePermissions sets EnableMixedModePermissions field to given value.

### HasEnableMixedModePermissions

`func (o *View) HasEnableMixedModePermissions() bool`

HasEnableMixedModePermissions returns a boolean if a field has been set.

### SetEnableMixedModePermissionsNil

`func (o *View) SetEnableMixedModePermissionsNil(b bool)`

 SetEnableMixedModePermissionsNil sets the value for EnableMixedModePermissions to be an explicit nil

### UnsetEnableMixedModePermissions
`func (o *View) UnsetEnableMixedModePermissions()`

UnsetEnableMixedModePermissions ensures that no value is present for EnableMixedModePermissions, not even an explicit nil
### GetEnableNfsViewDiscovery

`func (o *View) GetEnableNfsViewDiscovery() bool`

GetEnableNfsViewDiscovery returns the EnableNfsViewDiscovery field if non-nil, zero value otherwise.

### GetEnableNfsViewDiscoveryOk

`func (o *View) GetEnableNfsViewDiscoveryOk() (*bool, bool)`

GetEnableNfsViewDiscoveryOk returns a tuple with the EnableNfsViewDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNfsViewDiscovery

`func (o *View) SetEnableNfsViewDiscovery(v bool)`

SetEnableNfsViewDiscovery sets EnableNfsViewDiscovery field to given value.

### HasEnableNfsViewDiscovery

`func (o *View) HasEnableNfsViewDiscovery() bool`

HasEnableNfsViewDiscovery returns a boolean if a field has been set.

### SetEnableNfsViewDiscoveryNil

`func (o *View) SetEnableNfsViewDiscoveryNil(b bool)`

 SetEnableNfsViewDiscoveryNil sets the value for EnableNfsViewDiscovery to be an explicit nil

### UnsetEnableNfsViewDiscovery
`func (o *View) UnsetEnableNfsViewDiscovery()`

UnsetEnableNfsViewDiscovery ensures that no value is present for EnableNfsViewDiscovery, not even an explicit nil
### GetEnableOfflineCaching

`func (o *View) GetEnableOfflineCaching() bool`

GetEnableOfflineCaching returns the EnableOfflineCaching field if non-nil, zero value otherwise.

### GetEnableOfflineCachingOk

`func (o *View) GetEnableOfflineCachingOk() (*bool, bool)`

GetEnableOfflineCachingOk returns a tuple with the EnableOfflineCaching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableOfflineCaching

`func (o *View) SetEnableOfflineCaching(v bool)`

SetEnableOfflineCaching sets EnableOfflineCaching field to given value.

### HasEnableOfflineCaching

`func (o *View) HasEnableOfflineCaching() bool`

HasEnableOfflineCaching returns a boolean if a field has been set.

### SetEnableOfflineCachingNil

`func (o *View) SetEnableOfflineCachingNil(b bool)`

 SetEnableOfflineCachingNil sets the value for EnableOfflineCaching to be an explicit nil

### UnsetEnableOfflineCaching
`func (o *View) UnsetEnableOfflineCaching()`

UnsetEnableOfflineCaching ensures that no value is present for EnableOfflineCaching, not even an explicit nil
### GetEnableSmbAccessBasedEnumeration

`func (o *View) GetEnableSmbAccessBasedEnumeration() bool`

GetEnableSmbAccessBasedEnumeration returns the EnableSmbAccessBasedEnumeration field if non-nil, zero value otherwise.

### GetEnableSmbAccessBasedEnumerationOk

`func (o *View) GetEnableSmbAccessBasedEnumerationOk() (*bool, bool)`

GetEnableSmbAccessBasedEnumerationOk returns a tuple with the EnableSmbAccessBasedEnumeration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSmbAccessBasedEnumeration

`func (o *View) SetEnableSmbAccessBasedEnumeration(v bool)`

SetEnableSmbAccessBasedEnumeration sets EnableSmbAccessBasedEnumeration field to given value.

### HasEnableSmbAccessBasedEnumeration

`func (o *View) HasEnableSmbAccessBasedEnumeration() bool`

HasEnableSmbAccessBasedEnumeration returns a boolean if a field has been set.

### SetEnableSmbAccessBasedEnumerationNil

`func (o *View) SetEnableSmbAccessBasedEnumerationNil(b bool)`

 SetEnableSmbAccessBasedEnumerationNil sets the value for EnableSmbAccessBasedEnumeration to be an explicit nil

### UnsetEnableSmbAccessBasedEnumeration
`func (o *View) UnsetEnableSmbAccessBasedEnumeration()`

UnsetEnableSmbAccessBasedEnumeration ensures that no value is present for EnableSmbAccessBasedEnumeration, not even an explicit nil
### GetEnableSmbEncryption

`func (o *View) GetEnableSmbEncryption() bool`

GetEnableSmbEncryption returns the EnableSmbEncryption field if non-nil, zero value otherwise.

### GetEnableSmbEncryptionOk

`func (o *View) GetEnableSmbEncryptionOk() (*bool, bool)`

GetEnableSmbEncryptionOk returns a tuple with the EnableSmbEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSmbEncryption

`func (o *View) SetEnableSmbEncryption(v bool)`

SetEnableSmbEncryption sets EnableSmbEncryption field to given value.

### HasEnableSmbEncryption

`func (o *View) HasEnableSmbEncryption() bool`

HasEnableSmbEncryption returns a boolean if a field has been set.

### SetEnableSmbEncryptionNil

`func (o *View) SetEnableSmbEncryptionNil(b bool)`

 SetEnableSmbEncryptionNil sets the value for EnableSmbEncryption to be an explicit nil

### UnsetEnableSmbEncryption
`func (o *View) UnsetEnableSmbEncryption()`

UnsetEnableSmbEncryption ensures that no value is present for EnableSmbEncryption, not even an explicit nil
### GetEnableSmbOplock

`func (o *View) GetEnableSmbOplock() bool`

GetEnableSmbOplock returns the EnableSmbOplock field if non-nil, zero value otherwise.

### GetEnableSmbOplockOk

`func (o *View) GetEnableSmbOplockOk() (*bool, bool)`

GetEnableSmbOplockOk returns a tuple with the EnableSmbOplock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSmbOplock

`func (o *View) SetEnableSmbOplock(v bool)`

SetEnableSmbOplock sets EnableSmbOplock field to given value.

### HasEnableSmbOplock

`func (o *View) HasEnableSmbOplock() bool`

HasEnableSmbOplock returns a boolean if a field has been set.

### SetEnableSmbOplockNil

`func (o *View) SetEnableSmbOplockNil(b bool)`

 SetEnableSmbOplockNil sets the value for EnableSmbOplock to be an explicit nil

### UnsetEnableSmbOplock
`func (o *View) UnsetEnableSmbOplock()`

UnsetEnableSmbOplock ensures that no value is present for EnableSmbOplock, not even an explicit nil
### GetEnableSmbViewDiscovery

`func (o *View) GetEnableSmbViewDiscovery() bool`

GetEnableSmbViewDiscovery returns the EnableSmbViewDiscovery field if non-nil, zero value otherwise.

### GetEnableSmbViewDiscoveryOk

`func (o *View) GetEnableSmbViewDiscoveryOk() (*bool, bool)`

GetEnableSmbViewDiscoveryOk returns a tuple with the EnableSmbViewDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSmbViewDiscovery

`func (o *View) SetEnableSmbViewDiscovery(v bool)`

SetEnableSmbViewDiscovery sets EnableSmbViewDiscovery field to given value.

### HasEnableSmbViewDiscovery

`func (o *View) HasEnableSmbViewDiscovery() bool`

HasEnableSmbViewDiscovery returns a boolean if a field has been set.

### SetEnableSmbViewDiscoveryNil

`func (o *View) SetEnableSmbViewDiscoveryNil(b bool)`

 SetEnableSmbViewDiscoveryNil sets the value for EnableSmbViewDiscovery to be an explicit nil

### UnsetEnableSmbViewDiscovery
`func (o *View) UnsetEnableSmbViewDiscovery()`

UnsetEnableSmbViewDiscovery ensures that no value is present for EnableSmbViewDiscovery, not even an explicit nil
### GetEnforceSmbEncryption

`func (o *View) GetEnforceSmbEncryption() bool`

GetEnforceSmbEncryption returns the EnforceSmbEncryption field if non-nil, zero value otherwise.

### GetEnforceSmbEncryptionOk

`func (o *View) GetEnforceSmbEncryptionOk() (*bool, bool)`

GetEnforceSmbEncryptionOk returns a tuple with the EnforceSmbEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceSmbEncryption

`func (o *View) SetEnforceSmbEncryption(v bool)`

SetEnforceSmbEncryption sets EnforceSmbEncryption field to given value.

### HasEnforceSmbEncryption

`func (o *View) HasEnforceSmbEncryption() bool`

HasEnforceSmbEncryption returns a boolean if a field has been set.

### SetEnforceSmbEncryptionNil

`func (o *View) SetEnforceSmbEncryptionNil(b bool)`

 SetEnforceSmbEncryptionNil sets the value for EnforceSmbEncryption to be an explicit nil

### UnsetEnforceSmbEncryption
`func (o *View) UnsetEnforceSmbEncryption()`

UnsetEnforceSmbEncryption ensures that no value is present for EnforceSmbEncryption, not even an explicit nil
### GetFileExtensionFilter

`func (o *View) GetFileExtensionFilter() FileExtensionFilter`

GetFileExtensionFilter returns the FileExtensionFilter field if non-nil, zero value otherwise.

### GetFileExtensionFilterOk

`func (o *View) GetFileExtensionFilterOk() (*FileExtensionFilter, bool)`

GetFileExtensionFilterOk returns a tuple with the FileExtensionFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileExtensionFilter

`func (o *View) SetFileExtensionFilter(v FileExtensionFilter)`

SetFileExtensionFilter sets FileExtensionFilter field to given value.

### HasFileExtensionFilter

`func (o *View) HasFileExtensionFilter() bool`

HasFileExtensionFilter returns a boolean if a field has been set.

### GetFileLockConfig

`func (o *View) GetFileLockConfig() FileLevelDataLockConfig`

GetFileLockConfig returns the FileLockConfig field if non-nil, zero value otherwise.

### GetFileLockConfigOk

`func (o *View) GetFileLockConfigOk() (*FileLevelDataLockConfig, bool)`

GetFileLockConfigOk returns a tuple with the FileLockConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLockConfig

`func (o *View) SetFileLockConfig(v FileLevelDataLockConfig)`

SetFileLockConfig sets FileLockConfig field to given value.

### HasFileLockConfig

`func (o *View) HasFileLockConfig() bool`

HasFileLockConfig returns a boolean if a field has been set.

### GetIsTargetForMigratedData

`func (o *View) GetIsTargetForMigratedData() bool`

GetIsTargetForMigratedData returns the IsTargetForMigratedData field if non-nil, zero value otherwise.

### GetIsTargetForMigratedDataOk

`func (o *View) GetIsTargetForMigratedDataOk() (*bool, bool)`

GetIsTargetForMigratedDataOk returns a tuple with the IsTargetForMigratedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTargetForMigratedData

`func (o *View) SetIsTargetForMigratedData(v bool)`

SetIsTargetForMigratedData sets IsTargetForMigratedData field to given value.

### HasIsTargetForMigratedData

`func (o *View) HasIsTargetForMigratedData() bool`

HasIsTargetForMigratedData returns a boolean if a field has been set.

### SetIsTargetForMigratedDataNil

`func (o *View) SetIsTargetForMigratedDataNil(b bool)`

 SetIsTargetForMigratedDataNil sets the value for IsTargetForMigratedData to be an explicit nil

### UnsetIsTargetForMigratedData
`func (o *View) UnsetIsTargetForMigratedData()`

UnsetIsTargetForMigratedData ensures that no value is present for IsTargetForMigratedData, not even an explicit nil
### GetLogicalQuota

`func (o *View) GetLogicalQuota() QuotaPolicy`

GetLogicalQuota returns the LogicalQuota field if non-nil, zero value otherwise.

### GetLogicalQuotaOk

`func (o *View) GetLogicalQuotaOk() (*QuotaPolicy, bool)`

GetLogicalQuotaOk returns a tuple with the LogicalQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalQuota

`func (o *View) SetLogicalQuota(v QuotaPolicy)`

SetLogicalQuota sets LogicalQuota field to given value.

### HasLogicalQuota

`func (o *View) HasLogicalQuota() bool`

HasLogicalQuota returns a boolean if a field has been set.

### SetLogicalQuotaNil

`func (o *View) SetLogicalQuotaNil(b bool)`

 SetLogicalQuotaNil sets the value for LogicalQuota to be an explicit nil

### UnsetLogicalQuota
`func (o *View) UnsetLogicalQuota()`

UnsetLogicalQuota ensures that no value is present for LogicalQuota, not even an explicit nil
### GetLogicalUsageBytes

`func (o *View) GetLogicalUsageBytes() int64`

GetLogicalUsageBytes returns the LogicalUsageBytes field if non-nil, zero value otherwise.

### GetLogicalUsageBytesOk

`func (o *View) GetLogicalUsageBytesOk() (*int64, bool)`

GetLogicalUsageBytesOk returns a tuple with the LogicalUsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalUsageBytes

`func (o *View) SetLogicalUsageBytes(v int64)`

SetLogicalUsageBytes sets LogicalUsageBytes field to given value.

### HasLogicalUsageBytes

`func (o *View) HasLogicalUsageBytes() bool`

HasLogicalUsageBytes returns a boolean if a field has been set.

### SetLogicalUsageBytesNil

`func (o *View) SetLogicalUsageBytesNil(b bool)`

 SetLogicalUsageBytesNil sets the value for LogicalUsageBytes to be an explicit nil

### UnsetLogicalUsageBytes
`func (o *View) UnsetLogicalUsageBytes()`

UnsetLogicalUsageBytes ensures that no value is present for LogicalUsageBytes, not even an explicit nil
### GetName

`func (o *View) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *View) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *View) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *View) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *View) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *View) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNetgroupWhitelist

`func (o *View) GetNetgroupWhitelist() []NisNetgroup`

GetNetgroupWhitelist returns the NetgroupWhitelist field if non-nil, zero value otherwise.

### GetNetgroupWhitelistOk

`func (o *View) GetNetgroupWhitelistOk() (*[]NisNetgroup, bool)`

GetNetgroupWhitelistOk returns a tuple with the NetgroupWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetgroupWhitelist

`func (o *View) SetNetgroupWhitelist(v []NisNetgroup)`

SetNetgroupWhitelist sets NetgroupWhitelist field to given value.

### HasNetgroupWhitelist

`func (o *View) HasNetgroupWhitelist() bool`

HasNetgroupWhitelist returns a boolean if a field has been set.

### SetNetgroupWhitelistNil

`func (o *View) SetNetgroupWhitelistNil(b bool)`

 SetNetgroupWhitelistNil sets the value for NetgroupWhitelist to be an explicit nil

### UnsetNetgroupWhitelist
`func (o *View) UnsetNetgroupWhitelist()`

UnsetNetgroupWhitelist ensures that no value is present for NetgroupWhitelist, not even an explicit nil
### GetNfsAllSquash

`func (o *View) GetNfsAllSquash() NfsSquash`

GetNfsAllSquash returns the NfsAllSquash field if non-nil, zero value otherwise.

### GetNfsAllSquashOk

`func (o *View) GetNfsAllSquashOk() (*NfsSquash, bool)`

GetNfsAllSquashOk returns a tuple with the NfsAllSquash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsAllSquash

`func (o *View) SetNfsAllSquash(v NfsSquash)`

SetNfsAllSquash sets NfsAllSquash field to given value.

### HasNfsAllSquash

`func (o *View) HasNfsAllSquash() bool`

HasNfsAllSquash returns a boolean if a field has been set.

### GetNfsMountPath

`func (o *View) GetNfsMountPath() string`

GetNfsMountPath returns the NfsMountPath field if non-nil, zero value otherwise.

### GetNfsMountPathOk

`func (o *View) GetNfsMountPathOk() (*string, bool)`

GetNfsMountPathOk returns a tuple with the NfsMountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsMountPath

`func (o *View) SetNfsMountPath(v string)`

SetNfsMountPath sets NfsMountPath field to given value.

### HasNfsMountPath

`func (o *View) HasNfsMountPath() bool`

HasNfsMountPath returns a boolean if a field has been set.

### SetNfsMountPathNil

`func (o *View) SetNfsMountPathNil(b bool)`

 SetNfsMountPathNil sets the value for NfsMountPath to be an explicit nil

### UnsetNfsMountPath
`func (o *View) UnsetNfsMountPath()`

UnsetNfsMountPath ensures that no value is present for NfsMountPath, not even an explicit nil
### GetNfsRootPermissions

`func (o *View) GetNfsRootPermissions() NfsRootPermissions`

GetNfsRootPermissions returns the NfsRootPermissions field if non-nil, zero value otherwise.

### GetNfsRootPermissionsOk

`func (o *View) GetNfsRootPermissionsOk() (*NfsRootPermissions, bool)`

GetNfsRootPermissionsOk returns a tuple with the NfsRootPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsRootPermissions

`func (o *View) SetNfsRootPermissions(v NfsRootPermissions)`

SetNfsRootPermissions sets NfsRootPermissions field to given value.

### HasNfsRootPermissions

`func (o *View) HasNfsRootPermissions() bool`

HasNfsRootPermissions returns a boolean if a field has been set.

### GetNfsRootSquash

`func (o *View) GetNfsRootSquash() NfsSquash`

GetNfsRootSquash returns the NfsRootSquash field if non-nil, zero value otherwise.

### GetNfsRootSquashOk

`func (o *View) GetNfsRootSquashOk() (*NfsSquash, bool)`

GetNfsRootSquashOk returns a tuple with the NfsRootSquash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsRootSquash

`func (o *View) SetNfsRootSquash(v NfsSquash)`

SetNfsRootSquash sets NfsRootSquash field to given value.

### HasNfsRootSquash

`func (o *View) HasNfsRootSquash() bool`

HasNfsRootSquash returns a boolean if a field has been set.

### GetOverrideGlobalNetgroupWhitelist

`func (o *View) GetOverrideGlobalNetgroupWhitelist() bool`

GetOverrideGlobalNetgroupWhitelist returns the OverrideGlobalNetgroupWhitelist field if non-nil, zero value otherwise.

### GetOverrideGlobalNetgroupWhitelistOk

`func (o *View) GetOverrideGlobalNetgroupWhitelistOk() (*bool, bool)`

GetOverrideGlobalNetgroupWhitelistOk returns a tuple with the OverrideGlobalNetgroupWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideGlobalNetgroupWhitelist

`func (o *View) SetOverrideGlobalNetgroupWhitelist(v bool)`

SetOverrideGlobalNetgroupWhitelist sets OverrideGlobalNetgroupWhitelist field to given value.

### HasOverrideGlobalNetgroupWhitelist

`func (o *View) HasOverrideGlobalNetgroupWhitelist() bool`

HasOverrideGlobalNetgroupWhitelist returns a boolean if a field has been set.

### SetOverrideGlobalNetgroupWhitelistNil

`func (o *View) SetOverrideGlobalNetgroupWhitelistNil(b bool)`

 SetOverrideGlobalNetgroupWhitelistNil sets the value for OverrideGlobalNetgroupWhitelist to be an explicit nil

### UnsetOverrideGlobalNetgroupWhitelist
`func (o *View) UnsetOverrideGlobalNetgroupWhitelist()`

UnsetOverrideGlobalNetgroupWhitelist ensures that no value is present for OverrideGlobalNetgroupWhitelist, not even an explicit nil
### GetOverrideGlobalWhitelist

`func (o *View) GetOverrideGlobalWhitelist() bool`

GetOverrideGlobalWhitelist returns the OverrideGlobalWhitelist field if non-nil, zero value otherwise.

### GetOverrideGlobalWhitelistOk

`func (o *View) GetOverrideGlobalWhitelistOk() (*bool, bool)`

GetOverrideGlobalWhitelistOk returns a tuple with the OverrideGlobalWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideGlobalWhitelist

`func (o *View) SetOverrideGlobalWhitelist(v bool)`

SetOverrideGlobalWhitelist sets OverrideGlobalWhitelist field to given value.

### HasOverrideGlobalWhitelist

`func (o *View) HasOverrideGlobalWhitelist() bool`

HasOverrideGlobalWhitelist returns a boolean if a field has been set.

### SetOverrideGlobalWhitelistNil

`func (o *View) SetOverrideGlobalWhitelistNil(b bool)`

 SetOverrideGlobalWhitelistNil sets the value for OverrideGlobalWhitelist to be an explicit nil

### UnsetOverrideGlobalWhitelist
`func (o *View) UnsetOverrideGlobalWhitelist()`

UnsetOverrideGlobalWhitelist ensures that no value is present for OverrideGlobalWhitelist, not even an explicit nil
### GetProtocolAccess

`func (o *View) GetProtocolAccess() string`

GetProtocolAccess returns the ProtocolAccess field if non-nil, zero value otherwise.

### GetProtocolAccessOk

`func (o *View) GetProtocolAccessOk() (*string, bool)`

GetProtocolAccessOk returns a tuple with the ProtocolAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolAccess

`func (o *View) SetProtocolAccess(v string)`

SetProtocolAccess sets ProtocolAccess field to given value.

### HasProtocolAccess

`func (o *View) HasProtocolAccess() bool`

HasProtocolAccess returns a boolean if a field has been set.

### SetProtocolAccessNil

`func (o *View) SetProtocolAccessNil(b bool)`

 SetProtocolAccessNil sets the value for ProtocolAccess to be an explicit nil

### UnsetProtocolAccess
`func (o *View) UnsetProtocolAccess()`

UnsetProtocolAccess ensures that no value is present for ProtocolAccess, not even an explicit nil
### GetQos

`func (o *View) GetQos() QoS`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *View) GetQosOk() (*QoS, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *View) SetQos(v QoS)`

SetQos sets Qos field to given value.

### HasQos

`func (o *View) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetS3AccessPath

`func (o *View) GetS3AccessPath() string`

GetS3AccessPath returns the S3AccessPath field if non-nil, zero value otherwise.

### GetS3AccessPathOk

`func (o *View) GetS3AccessPathOk() (*string, bool)`

GetS3AccessPathOk returns a tuple with the S3AccessPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3AccessPath

`func (o *View) SetS3AccessPath(v string)`

SetS3AccessPath sets S3AccessPath field to given value.

### HasS3AccessPath

`func (o *View) HasS3AccessPath() bool`

HasS3AccessPath returns a boolean if a field has been set.

### SetS3AccessPathNil

`func (o *View) SetS3AccessPathNil(b bool)`

 SetS3AccessPathNil sets the value for S3AccessPath to be an explicit nil

### UnsetS3AccessPath
`func (o *View) UnsetS3AccessPath()`

UnsetS3AccessPath ensures that no value is present for S3AccessPath, not even an explicit nil
### GetS3KeyMappingConfig

`func (o *View) GetS3KeyMappingConfig() string`

GetS3KeyMappingConfig returns the S3KeyMappingConfig field if non-nil, zero value otherwise.

### GetS3KeyMappingConfigOk

`func (o *View) GetS3KeyMappingConfigOk() (*string, bool)`

GetS3KeyMappingConfigOk returns a tuple with the S3KeyMappingConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3KeyMappingConfig

`func (o *View) SetS3KeyMappingConfig(v string)`

SetS3KeyMappingConfig sets S3KeyMappingConfig field to given value.

### HasS3KeyMappingConfig

`func (o *View) HasS3KeyMappingConfig() bool`

HasS3KeyMappingConfig returns a boolean if a field has been set.

### SetS3KeyMappingConfigNil

`func (o *View) SetS3KeyMappingConfigNil(b bool)`

 SetS3KeyMappingConfigNil sets the value for S3KeyMappingConfig to be an explicit nil

### UnsetS3KeyMappingConfig
`func (o *View) UnsetS3KeyMappingConfig()`

UnsetS3KeyMappingConfig ensures that no value is present for S3KeyMappingConfig, not even an explicit nil
### GetSecurityMode

`func (o *View) GetSecurityMode() string`

GetSecurityMode returns the SecurityMode field if non-nil, zero value otherwise.

### GetSecurityModeOk

`func (o *View) GetSecurityModeOk() (*string, bool)`

GetSecurityModeOk returns a tuple with the SecurityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityMode

`func (o *View) SetSecurityMode(v string)`

SetSecurityMode sets SecurityMode field to given value.

### HasSecurityMode

`func (o *View) HasSecurityMode() bool`

HasSecurityMode returns a boolean if a field has been set.

### SetSecurityModeNil

`func (o *View) SetSecurityModeNil(b bool)`

 SetSecurityModeNil sets the value for SecurityMode to be an explicit nil

### UnsetSecurityMode
`func (o *View) UnsetSecurityMode()`

UnsetSecurityMode ensures that no value is present for SecurityMode, not even an explicit nil
### GetSharePermissions

`func (o *View) GetSharePermissions() []SmbPermission`

GetSharePermissions returns the SharePermissions field if non-nil, zero value otherwise.

### GetSharePermissionsOk

`func (o *View) GetSharePermissionsOk() (*[]SmbPermission, bool)`

GetSharePermissionsOk returns a tuple with the SharePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharePermissions

`func (o *View) SetSharePermissions(v []SmbPermission)`

SetSharePermissions sets SharePermissions field to given value.

### HasSharePermissions

`func (o *View) HasSharePermissions() bool`

HasSharePermissions returns a boolean if a field has been set.

### SetSharePermissionsNil

`func (o *View) SetSharePermissionsNil(b bool)`

 SetSharePermissionsNil sets the value for SharePermissions to be an explicit nil

### UnsetSharePermissions
`func (o *View) UnsetSharePermissions()`

UnsetSharePermissions ensures that no value is present for SharePermissions, not even an explicit nil
### GetSmbMountPath

`func (o *View) GetSmbMountPath() string`

GetSmbMountPath returns the SmbMountPath field if non-nil, zero value otherwise.

### GetSmbMountPathOk

`func (o *View) GetSmbMountPathOk() (*string, bool)`

GetSmbMountPathOk returns a tuple with the SmbMountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbMountPath

`func (o *View) SetSmbMountPath(v string)`

SetSmbMountPath sets SmbMountPath field to given value.

### HasSmbMountPath

`func (o *View) HasSmbMountPath() bool`

HasSmbMountPath returns a boolean if a field has been set.

### SetSmbMountPathNil

`func (o *View) SetSmbMountPathNil(b bool)`

 SetSmbMountPathNil sets the value for SmbMountPath to be an explicit nil

### UnsetSmbMountPath
`func (o *View) UnsetSmbMountPath()`

UnsetSmbMountPath ensures that no value is present for SmbMountPath, not even an explicit nil
### GetSmbPermissionsInfo

`func (o *View) GetSmbPermissionsInfo() SmbPermissionsInfo`

GetSmbPermissionsInfo returns the SmbPermissionsInfo field if non-nil, zero value otherwise.

### GetSmbPermissionsInfoOk

`func (o *View) GetSmbPermissionsInfoOk() (*SmbPermissionsInfo, bool)`

GetSmbPermissionsInfoOk returns a tuple with the SmbPermissionsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbPermissionsInfo

`func (o *View) SetSmbPermissionsInfo(v SmbPermissionsInfo)`

SetSmbPermissionsInfo sets SmbPermissionsInfo field to given value.

### HasSmbPermissionsInfo

`func (o *View) HasSmbPermissionsInfo() bool`

HasSmbPermissionsInfo returns a boolean if a field has been set.

### GetStats

`func (o *View) GetStats() ViewStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *View) GetStatsOk() (*ViewStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *View) SetStats(v ViewStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *View) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetStoragePolicyOverride

`func (o *View) GetStoragePolicyOverride() StoragePolicyOverride`

GetStoragePolicyOverride returns the StoragePolicyOverride field if non-nil, zero value otherwise.

### GetStoragePolicyOverrideOk

`func (o *View) GetStoragePolicyOverrideOk() (*StoragePolicyOverride, bool)`

GetStoragePolicyOverrideOk returns a tuple with the StoragePolicyOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePolicyOverride

`func (o *View) SetStoragePolicyOverride(v StoragePolicyOverride)`

SetStoragePolicyOverride sets StoragePolicyOverride field to given value.

### HasStoragePolicyOverride

`func (o *View) HasStoragePolicyOverride() bool`

HasStoragePolicyOverride returns a boolean if a field has been set.

### GetSubnetWhitelist

`func (o *View) GetSubnetWhitelist() []Subnet`

GetSubnetWhitelist returns the SubnetWhitelist field if non-nil, zero value otherwise.

### GetSubnetWhitelistOk

`func (o *View) GetSubnetWhitelistOk() (*[]Subnet, bool)`

GetSubnetWhitelistOk returns a tuple with the SubnetWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetWhitelist

`func (o *View) SetSubnetWhitelist(v []Subnet)`

SetSubnetWhitelist sets SubnetWhitelist field to given value.

### HasSubnetWhitelist

`func (o *View) HasSubnetWhitelist() bool`

HasSubnetWhitelist returns a boolean if a field has been set.

### SetSubnetWhitelistNil

`func (o *View) SetSubnetWhitelistNil(b bool)`

 SetSubnetWhitelistNil sets the value for SubnetWhitelist to be an explicit nil

### UnsetSubnetWhitelist
`func (o *View) UnsetSubnetWhitelist()`

UnsetSubnetWhitelist ensures that no value is present for SubnetWhitelist, not even an explicit nil
### GetSwiftProjectDomain

`func (o *View) GetSwiftProjectDomain() string`

GetSwiftProjectDomain returns the SwiftProjectDomain field if non-nil, zero value otherwise.

### GetSwiftProjectDomainOk

`func (o *View) GetSwiftProjectDomainOk() (*string, bool)`

GetSwiftProjectDomainOk returns a tuple with the SwiftProjectDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftProjectDomain

`func (o *View) SetSwiftProjectDomain(v string)`

SetSwiftProjectDomain sets SwiftProjectDomain field to given value.

### HasSwiftProjectDomain

`func (o *View) HasSwiftProjectDomain() bool`

HasSwiftProjectDomain returns a boolean if a field has been set.

### SetSwiftProjectDomainNil

`func (o *View) SetSwiftProjectDomainNil(b bool)`

 SetSwiftProjectDomainNil sets the value for SwiftProjectDomain to be an explicit nil

### UnsetSwiftProjectDomain
`func (o *View) UnsetSwiftProjectDomain()`

UnsetSwiftProjectDomain ensures that no value is present for SwiftProjectDomain, not even an explicit nil
### GetSwiftProjectName

`func (o *View) GetSwiftProjectName() string`

GetSwiftProjectName returns the SwiftProjectName field if non-nil, zero value otherwise.

### GetSwiftProjectNameOk

`func (o *View) GetSwiftProjectNameOk() (*string, bool)`

GetSwiftProjectNameOk returns a tuple with the SwiftProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftProjectName

`func (o *View) SetSwiftProjectName(v string)`

SetSwiftProjectName sets SwiftProjectName field to given value.

### HasSwiftProjectName

`func (o *View) HasSwiftProjectName() bool`

HasSwiftProjectName returns a boolean if a field has been set.

### SetSwiftProjectNameNil

`func (o *View) SetSwiftProjectNameNil(b bool)`

 SetSwiftProjectNameNil sets the value for SwiftProjectName to be an explicit nil

### UnsetSwiftProjectName
`func (o *View) UnsetSwiftProjectName()`

UnsetSwiftProjectName ensures that no value is present for SwiftProjectName, not even an explicit nil
### GetSwiftUserDomain

`func (o *View) GetSwiftUserDomain() string`

GetSwiftUserDomain returns the SwiftUserDomain field if non-nil, zero value otherwise.

### GetSwiftUserDomainOk

`func (o *View) GetSwiftUserDomainOk() (*string, bool)`

GetSwiftUserDomainOk returns a tuple with the SwiftUserDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftUserDomain

`func (o *View) SetSwiftUserDomain(v string)`

SetSwiftUserDomain sets SwiftUserDomain field to given value.

### HasSwiftUserDomain

`func (o *View) HasSwiftUserDomain() bool`

HasSwiftUserDomain returns a boolean if a field has been set.

### SetSwiftUserDomainNil

`func (o *View) SetSwiftUserDomainNil(b bool)`

 SetSwiftUserDomainNil sets the value for SwiftUserDomain to be an explicit nil

### UnsetSwiftUserDomain
`func (o *View) UnsetSwiftUserDomain()`

UnsetSwiftUserDomain ensures that no value is present for SwiftUserDomain, not even an explicit nil
### GetSwiftUsername

`func (o *View) GetSwiftUsername() string`

GetSwiftUsername returns the SwiftUsername field if non-nil, zero value otherwise.

### GetSwiftUsernameOk

`func (o *View) GetSwiftUsernameOk() (*string, bool)`

GetSwiftUsernameOk returns a tuple with the SwiftUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftUsername

`func (o *View) SetSwiftUsername(v string)`

SetSwiftUsername sets SwiftUsername field to given value.

### HasSwiftUsername

`func (o *View) HasSwiftUsername() bool`

HasSwiftUsername returns a boolean if a field has been set.

### SetSwiftUsernameNil

`func (o *View) SetSwiftUsernameNil(b bool)`

 SetSwiftUsernameNil sets the value for SwiftUsername to be an explicit nil

### UnsetSwiftUsername
`func (o *View) UnsetSwiftUsername()`

UnsetSwiftUsername ensures that no value is present for SwiftUsername, not even an explicit nil
### GetTenantId

`func (o *View) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *View) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *View) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *View) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *View) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *View) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetViewBoxId

`func (o *View) GetViewBoxId() int64`

GetViewBoxId returns the ViewBoxId field if non-nil, zero value otherwise.

### GetViewBoxIdOk

`func (o *View) GetViewBoxIdOk() (*int64, bool)`

GetViewBoxIdOk returns a tuple with the ViewBoxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewBoxId

`func (o *View) SetViewBoxId(v int64)`

SetViewBoxId sets ViewBoxId field to given value.

### HasViewBoxId

`func (o *View) HasViewBoxId() bool`

HasViewBoxId returns a boolean if a field has been set.

### SetViewBoxIdNil

`func (o *View) SetViewBoxIdNil(b bool)`

 SetViewBoxIdNil sets the value for ViewBoxId to be an explicit nil

### UnsetViewBoxId
`func (o *View) UnsetViewBoxId()`

UnsetViewBoxId ensures that no value is present for ViewBoxId, not even an explicit nil
### GetViewBoxName

`func (o *View) GetViewBoxName() string`

GetViewBoxName returns the ViewBoxName field if non-nil, zero value otherwise.

### GetViewBoxNameOk

`func (o *View) GetViewBoxNameOk() (*string, bool)`

GetViewBoxNameOk returns a tuple with the ViewBoxName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewBoxName

`func (o *View) SetViewBoxName(v string)`

SetViewBoxName sets ViewBoxName field to given value.

### HasViewBoxName

`func (o *View) HasViewBoxName() bool`

HasViewBoxName returns a boolean if a field has been set.

### SetViewBoxNameNil

`func (o *View) SetViewBoxNameNil(b bool)`

 SetViewBoxNameNil sets the value for ViewBoxName to be an explicit nil

### UnsetViewBoxName
`func (o *View) UnsetViewBoxName()`

UnsetViewBoxName ensures that no value is present for ViewBoxName, not even an explicit nil
### GetViewId

`func (o *View) GetViewId() int64`

GetViewId returns the ViewId field if non-nil, zero value otherwise.

### GetViewIdOk

`func (o *View) GetViewIdOk() (*int64, bool)`

GetViewIdOk returns a tuple with the ViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewId

`func (o *View) SetViewId(v int64)`

SetViewId sets ViewId field to given value.

### HasViewId

`func (o *View) HasViewId() bool`

HasViewId returns a boolean if a field has been set.

### SetViewIdNil

`func (o *View) SetViewIdNil(b bool)`

 SetViewIdNil sets the value for ViewId to be an explicit nil

### UnsetViewId
`func (o *View) UnsetViewId()`

UnsetViewId ensures that no value is present for ViewId, not even an explicit nil
### GetViewLockEnabled

`func (o *View) GetViewLockEnabled() bool`

GetViewLockEnabled returns the ViewLockEnabled field if non-nil, zero value otherwise.

### GetViewLockEnabledOk

`func (o *View) GetViewLockEnabledOk() (*bool, bool)`

GetViewLockEnabledOk returns a tuple with the ViewLockEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewLockEnabled

`func (o *View) SetViewLockEnabled(v bool)`

SetViewLockEnabled sets ViewLockEnabled field to given value.

### HasViewLockEnabled

`func (o *View) HasViewLockEnabled() bool`

HasViewLockEnabled returns a boolean if a field has been set.

### SetViewLockEnabledNil

`func (o *View) SetViewLockEnabledNil(b bool)`

 SetViewLockEnabledNil sets the value for ViewLockEnabled to be an explicit nil

### UnsetViewLockEnabled
`func (o *View) UnsetViewLockEnabled()`

UnsetViewLockEnabled ensures that no value is present for ViewLockEnabled, not even an explicit nil
### GetViewProtection

`func (o *View) GetViewProtection() ViewProtection`

GetViewProtection returns the ViewProtection field if non-nil, zero value otherwise.

### GetViewProtectionOk

`func (o *View) GetViewProtectionOk() (*ViewProtection, bool)`

GetViewProtectionOk returns a tuple with the ViewProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewProtection

`func (o *View) SetViewProtection(v ViewProtection)`

SetViewProtection sets ViewProtection field to given value.

### HasViewProtection

`func (o *View) HasViewProtection() bool`

HasViewProtection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


