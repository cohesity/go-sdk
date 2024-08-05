# View

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessSids** | Pointer to **[]string** | Array of Security Identifiers (SIDs) Specifies the list of security identifiers (SIDs) for the restricted Principals who have access to this View. | [optional] 
**AllowMountOnWindows** | Pointer to **NullableBool** | Specifies if this View can be mounted using the NFS protocol on Windows systems. If true, this View can be NFS mounted on Windows systems. | [optional] 
**AntivirusScanConfig** | Pointer to **map[string]interface{}** | Specifies the antivirus scan config settings for this View. | [optional] 
**Category** | **NullableString** | Specifies the category of the View. | 
**Description** | Pointer to **NullableString** | Specifies an optional text description about the View. | [optional] 
**EnableFilerAuditLogging** | Pointer to **NullableBool** | Specifies if Filer Audit Logging is enabled for this view. | [optional] 
**EnableLiveIndexing** | Pointer to **NullableBool** | Specifies whether to enable live indexing for the view. | [optional] 
**EnableMetadataAccelerator** | Pointer to **NullableBool** | Specifies if metadata accelerator is enabled for this view. Only supported while creating a view. | [optional] 
**EnableMinion** | Pointer to **NullableBool** | Specifies if this view should allow minion or not. If true, this will allow minion. | [optional] 
**EnableOfflineCaching** | Pointer to **NullableBool** | Specifies whether to enable offline file caching of the view. | [optional] 
**FileExtensionFilter** | Pointer to **map[string]interface{}** | Optional filtering criteria that should be satisfied by all the files created in this view. It does not affect existing files. | [optional] 
**FileLockConfig** | Pointer to **map[string]interface{}** | Optional config that enables file locking for this view. It cannot be disabled during the edit of a view, if it has been enabled during the creation of the view. Also, it cannot be enabled if it was disabled during the creation of the view. | [optional] 
**FilerLifecycleManagement** | Pointer to **map[string]interface{}** | Specifies the Lifecycle policy of this filer (NFS/SMB) view. | [optional] 
**IsExternallyTriggeredBackupTarget** | Pointer to **NullableBool** | Specifies whether the view is for externally triggered backup target. If so, Magneto will ignore the backup schedule for the view protection job of this view. By default it is disabled. | [optional] 
**IsReadOnly** | Pointer to **NullableBool** | Specifies if the view is a read only view. User will no longer be able to write to this view if this is set to true. | [optional] 
**LexicographicPrefetch** | Pointer to **NullableBool** | If small files are accessed sequentially from a client, this specifies whether to detect and prefetch files based on the lexicographic index to improve file read performance. | [optional] 
**LogicalQuota** | Pointer to **map[string]interface{}** | Specifies an optional logical quota limit (in bytes) for the usage allowed on this View. (Logical data is when the data is fully hydrated and expanded.) This limit overrides the limit inherited from the Storage Domain (View Box) (if set). If logicalQuota is nil, the limit is inherited from the Storage Domain (View Box) (if set). A new write is not allowed if the Storage Domain (View Box) will exceed the specified quota. However, it takes time for the Cohesity Cluster to calculate the usage across Nodes, so the limit may be exceeded by a small amount. In addition, if the limit is increased or data is removed, there may be a delay before the Cohesity Cluster allows more data to be written to the View, as the Cluster is calculating the usage across Nodes. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the View. | [optional] 
**NetgroupWhitelist** | Pointer to **map[string]interface{}** | Array of Netgroups. Specifies a list of netgroups with domains that have permissions to access the View. (Overrides or extends the Netgroup specified at the global Cohesity Cluster level.) | [optional] 
**OverrideGlobalNetgroupWhitelist** | Pointer to **NullableBool** | Specifies whether view level client netgroup whitelist overrides cluster and global setting. | [optional] 
**OverrideGlobalSubnetWhitelist** | Pointer to **NullableBool** | Specifies whether view level client subnet whitelist overrides cluster and global setting. | [optional] 
**ProtocolAccess** | [**[]ViewProtocol**](ViewProtocol.md) | Specifies the supported Protocols for the View. | 
**Qos** | **map[string]interface{}** | Specifies the Quality of Service (QoS) Policy for the View. | 
**SecurityMode** | Pointer to **NullableString** | Specifies the security mode used for this view. Currently we support the following modes: Native, Unified and NTFS style. &#39;NativeMode&#39; indicates a native security mode. &#39;UnifiedMode&#39; indicates a unified security mode. &#39;NtfsMode&#39; indicates a NTFS style security mode. | [optional] 
**SelfServiceSnapshotConfig** | Pointer to **map[string]interface{}** | Specifies self service config of this view. | [optional] 
**StoragePolicyOverride** | Pointer to **map[string]interface{}** | Specifies if inline deduplication and compression settings inherited from the Storage Domain (View Box) should be disabled for this View. | [optional] 
**SubnetWhitelist** | Pointer to [**[]Subnet**](Subnet.md) | Array of Subnets. Specifies a list of Subnets with IP addresses that have permissions to access the View. (Overrides or extends the Subnets specified at the global Cohesity Cluster level.) | [optional] 
**TenantId** | Pointer to **NullableString** | Optional tenant id who has access to this View. | [optional] 
**ViewLockEnabled** | Pointer to **NullableBool** | Specifies whether view lock is enabled. If enabled the view cannot be modified or deleted until unlock. By default it is disabled. | [optional] 
**ViewPinningConfig** | Pointer to **map[string]interface{}** | Specifies the pinning config of this view. | [optional] 
**EnableNfsKerberosAuthentication** | Pointer to **NullableBool** | If set, it enables NFS Kerberos Authentication | [optional] 
**EnableNfsKerberosIntegrity** | Pointer to **NullableBool** | If set, it enables NFS Kerberos Integrity | [optional] 
**EnableNfsKerberosPrivacy** | Pointer to **NullableBool** | If set, it enables NFS Kerberos Privacy | [optional] 
**EnableNfsUnixAuthentication** | Pointer to **NullableBool** | If set, it enables NFS UNIX Authentication | [optional] 
**EnableNfsViewDiscovery** | Pointer to **NullableBool** | If set, it enables discovery of view for NFS. | [optional] 
**EnableNfsWcc** | Pointer to **NullableBool** | If set, it enables NFS weak cache consistency. | [optional] 
**NfsAllSquash** | Pointer to [**NfsConfigNfsAllSquash**](NfsConfigNfsAllSquash.md) |  | [optional] 
**NfsRootPermissions** | Pointer to [**NfsConfigNfsRootPermissions**](NfsConfigNfsRootPermissions.md) |  | [optional] 
**NfsRootSquash** | Pointer to [**NfsConfigNfsRootSquash**](NfsConfigNfsRootSquash.md) |  | [optional] 
**EnableFastDurableHandle** | Pointer to **NullableBool** | Specifies whether fast durable handle is enabled. If enabled, view open handle will be kept in memory, which results in a higher performance. But the handles cannot be recovered if node or service crashes. | [optional] 
**EnableSmbAccessBasedEnumeration** | Pointer to **NullableBool** | Specifies if access-based enumeration should be enabled. If &#39;true&#39;, only files and folders that the user has permissions to access are visible on the SMB share for that user. | [optional] 
**EnableSmbEncryption** | Pointer to **NullableBool** | Specifies the SMB encryption for the View. If set, it enables the SMB encryption for the View. Encryption is supported only by SMB 3.x dialects. Dialects that do not support would still access data in unencrypted format. | [optional] 
**EnableSmbOplock** | Pointer to **NullableBool** | Specifies whether SMB opportunistic lock is enabled. | [optional] 
**EnableSmbViewDiscovery** | Pointer to **NullableBool** | If set, it enables discovery of view for SMB. | [optional] 
**EnforceSmbEncryption** | Pointer to **NullableBool** | Specifies the SMB encryption for all the sessions for the View. If set, encryption is enforced for all the sessions for the View. When enabled all future and existing unencrypted sessions are disallowed. | [optional] 
**SharePermissions** | Pointer to [**SmbConfigSharePermissions**](SmbConfigSharePermissions.md) |  | [optional] 
**SmbPermissionsInfo** | Pointer to [**SmbConfigSmbPermissionsInfo**](SmbConfigSmbPermissionsInfo.md) |  | [optional] 
**AclConfig** | Pointer to [**S3ConfigAclConfig**](S3ConfigAclConfig.md) |  | [optional] 
**BucketPolicy** | Pointer to [**S3ConfigBucketPolicy**](S3ConfigBucketPolicy.md) |  | [optional] 
**EnableAbac** | Pointer to **NullableBool** | Specifies if this View has S3 ABAC enabled. This can only be set while creating a view. The ABAC server corresponding the tenant will be used for authentication and authorization checks.  | [optional] 
**LifecycleManagement** | Pointer to [**S3ConfigLifecycleManagement**](S3ConfigLifecycleManagement.md) |  | [optional] 
**OwnerInfo** | Pointer to [**S3ConfigOwnerInfo**](S3ConfigOwnerInfo.md) |  | [optional] 
**S3AccessPath** | Pointer to **NullableString** | Specifies the path to access this View as an S3 share. | [optional] [readonly] 
**S3EfficientMpuMaxSubfiles** | Pointer to **NullableInt32** | Specifies if this View has S3 MPU 2.0 enabled. This can set while editing a view.  | [optional] 
**S3EnableEfficientMpu** | Pointer to **NullableBool** | Specifies if this View has S3 MPU 2.0 enabled. This can set while editing a view.  | [optional] 
**S3MigrationAction** | Pointer to **NullableString** | Specifies the S3 migration action to be performed on this View. Supported migration actions are: [Enable, Cancel, Pause, Resume]. | [optional] 
**S3MigrationState** | Pointer to **NullableString** | Specifies the current S3 migration state for this View. A View can be under following migration states: [Eligible, Enable, Pause, Complete, UnderMigration]. | [optional] 
**Versioning** | Pointer to **NullableString** | Specifies the versioning state of S3 bucket. Buckets can be in one of three states: UnVersioned (default), VersioningEnabled, or VersioningSuspended. Once versioning is enabled for a bucket, it can never return to an UnVersioned state. However, versioning on the bucket can be suspended. | [optional] 
**SwiftProjectDomain** | Pointer to **NullableString** | Specifies the Keystone project domain. | [optional] 
**SwiftProjectName** | Pointer to **NullableString** | Specifies the Keystone project name. | [optional] 
**SwiftUserDomain** | Pointer to **NullableString** | Specifies the Keystone user domain. | [optional] 
**SwiftUsername** | Pointer to **NullableString** | Specifies the Keystone username. | [optional] 
**Aliases** | Pointer to [**[]ViewAliasInfo**](ViewAliasInfo.md) | Aliases created for the view. A view alias allows a directory path inside a view to be mounted using the alias name. | [optional] [readonly] 
**BasicMountPath** | Pointer to **NullableString** | Specifies the NFS mount path of the View (without the hostname information). This path is used to support NFS mounting of the paths specified in the nfsExportPathList on Windows systems. | [optional] [readonly] 
**CaseInsensitiveNamesEnabled** | Pointer to **NullableBool** | Specifies whether to support case insensitive file/folder names. This parameter can only be set during create and cannot be changed. | [optional] [readonly] 
**CreateTimeMsecs** | Pointer to **NullableInt64** | Specifies the time that the View was created in milliseconds. | [optional] [readonly] 
**DataLockExpiryUsecs** | Pointer to **NullableInt64** | DataLock (Write Once Read Many) lock expiry epoch time in microseconds. If a view is marked as a DataLock view, only a Data Security Officer (a user having Data Security Privilege) can delete the view until the lock expiry time. | [optional] 
**FileCountBySize** | Pointer to [**[]FileCount**](FileCount.md) | Specifies the file count by size for the View. | [optional] 
**Intent** | Pointer to **map[string]interface{}** | Specifies the intent of the View. | [optional] 
**IsCategoryInferred** | Pointer to **NullableBool** | If True, category in response is not set by user but inferred by Iris because none is set. Category can only be none when view was created by v1 API or cloned from a view created by v1 API.  Inference Logic is as follows: 1. Object Services if only S3 or Swift protocol is selected. 2. Backup Target only if one read-write protocol is selected and    QoS is \&quot;Backup Target Commvault\&quot; or \&quot;Backup Target SSD\&quot;. 3. File Services if there are more than 1 read-write protocol or    it doesn&#39;t fit any other category. | [optional] [readonly] 
**IsTargetForMigratedData** | Pointer to **NullableBool** | Specifies if a view contains migrated data. | [optional] [readonly] 
**NfsMountPaths** | Pointer to **[]string** | Array of NFS Paths. Specifies the path for mounting this View as an NFS share. If Kerberos Provider has multiple hostaliases, each host alias has  its own path. | [optional] [readonly] 
**ObjectServicesMappingConfig** | Pointer to **NullableString** | Specifies the Object Services key mapping config of the view. This parameter can only be set during create and cannot be changed. Configuration of Object Services key mapping. Specifies the type of Object Services key mapping config. | [optional] [readonly] 
**OwnerSid** | Pointer to **NullableString** | Specifies the sid of the view owner. | [optional] 
**S3FolderSupportEnabled** | Pointer to **NullableBool** | Specifies whether to support s3 folder support feature. This parameter can only be set during create and cannot be changed. | [optional] [readonly] 
**SmbMountPaths** | Pointer to **[]string** | Array of SMB Paths. Specifies the possible paths that can be used to mount this View as a SMB share. If Active Directory has multiple account names; each machine account has its own path. | [optional] [readonly] 
**Stats** | Pointer to **map[string]interface{}** | Specifies statistics about the View. | [optional] 
**StorageDomainId** | Pointer to **NullableInt64** | Specifies the id of the Storage Domain (View Box) where the View is stored. | [optional] [readonly] 
**StorageDomainName** | Pointer to **NullableString** | Specifies the name of the Storage Domain (View Box) where the View is stored. | [optional] [readonly] 
**ViewFailover** | Pointer to **map[string]interface{}** | Specifies the information about the failover of the view. | [optional] 
**ViewId** | Pointer to **NullableInt64** | Specifies an id of the View assigned by the Cohesity Cluster. | [optional] [readonly] 
**ViewProtection** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewView

`func NewView(category NullableString, protocolAccess []ViewProtocol, qos map[string]interface{}, ) *View`

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
### GetAllowMountOnWindows

`func (o *View) GetAllowMountOnWindows() bool`

GetAllowMountOnWindows returns the AllowMountOnWindows field if non-nil, zero value otherwise.

### GetAllowMountOnWindowsOk

`func (o *View) GetAllowMountOnWindowsOk() (*bool, bool)`

GetAllowMountOnWindowsOk returns a tuple with the AllowMountOnWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMountOnWindows

`func (o *View) SetAllowMountOnWindows(v bool)`

SetAllowMountOnWindows sets AllowMountOnWindows field to given value.

### HasAllowMountOnWindows

`func (o *View) HasAllowMountOnWindows() bool`

HasAllowMountOnWindows returns a boolean if a field has been set.

### SetAllowMountOnWindowsNil

`func (o *View) SetAllowMountOnWindowsNil(b bool)`

 SetAllowMountOnWindowsNil sets the value for AllowMountOnWindows to be an explicit nil

### UnsetAllowMountOnWindows
`func (o *View) UnsetAllowMountOnWindows()`

UnsetAllowMountOnWindows ensures that no value is present for AllowMountOnWindows, not even an explicit nil
### GetAntivirusScanConfig

`func (o *View) GetAntivirusScanConfig() map[string]interface{}`

GetAntivirusScanConfig returns the AntivirusScanConfig field if non-nil, zero value otherwise.

### GetAntivirusScanConfigOk

`func (o *View) GetAntivirusScanConfigOk() (*map[string]interface{}, bool)`

GetAntivirusScanConfigOk returns a tuple with the AntivirusScanConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntivirusScanConfig

`func (o *View) SetAntivirusScanConfig(v map[string]interface{})`

SetAntivirusScanConfig sets AntivirusScanConfig field to given value.

### HasAntivirusScanConfig

`func (o *View) HasAntivirusScanConfig() bool`

HasAntivirusScanConfig returns a boolean if a field has been set.

### GetCategory

`func (o *View) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *View) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *View) SetCategory(v string)`

SetCategory sets Category field to given value.


### SetCategoryNil

`func (o *View) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *View) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
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
### GetEnableMetadataAccelerator

`func (o *View) GetEnableMetadataAccelerator() bool`

GetEnableMetadataAccelerator returns the EnableMetadataAccelerator field if non-nil, zero value otherwise.

### GetEnableMetadataAcceleratorOk

`func (o *View) GetEnableMetadataAcceleratorOk() (*bool, bool)`

GetEnableMetadataAcceleratorOk returns a tuple with the EnableMetadataAccelerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMetadataAccelerator

`func (o *View) SetEnableMetadataAccelerator(v bool)`

SetEnableMetadataAccelerator sets EnableMetadataAccelerator field to given value.

### HasEnableMetadataAccelerator

`func (o *View) HasEnableMetadataAccelerator() bool`

HasEnableMetadataAccelerator returns a boolean if a field has been set.

### SetEnableMetadataAcceleratorNil

`func (o *View) SetEnableMetadataAcceleratorNil(b bool)`

 SetEnableMetadataAcceleratorNil sets the value for EnableMetadataAccelerator to be an explicit nil

### UnsetEnableMetadataAccelerator
`func (o *View) UnsetEnableMetadataAccelerator()`

UnsetEnableMetadataAccelerator ensures that no value is present for EnableMetadataAccelerator, not even an explicit nil
### GetEnableMinion

`func (o *View) GetEnableMinion() bool`

GetEnableMinion returns the EnableMinion field if non-nil, zero value otherwise.

### GetEnableMinionOk

`func (o *View) GetEnableMinionOk() (*bool, bool)`

GetEnableMinionOk returns a tuple with the EnableMinion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMinion

`func (o *View) SetEnableMinion(v bool)`

SetEnableMinion sets EnableMinion field to given value.

### HasEnableMinion

`func (o *View) HasEnableMinion() bool`

HasEnableMinion returns a boolean if a field has been set.

### SetEnableMinionNil

`func (o *View) SetEnableMinionNil(b bool)`

 SetEnableMinionNil sets the value for EnableMinion to be an explicit nil

### UnsetEnableMinion
`func (o *View) UnsetEnableMinion()`

UnsetEnableMinion ensures that no value is present for EnableMinion, not even an explicit nil
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
### GetFileExtensionFilter

`func (o *View) GetFileExtensionFilter() map[string]interface{}`

GetFileExtensionFilter returns the FileExtensionFilter field if non-nil, zero value otherwise.

### GetFileExtensionFilterOk

`func (o *View) GetFileExtensionFilterOk() (*map[string]interface{}, bool)`

GetFileExtensionFilterOk returns a tuple with the FileExtensionFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileExtensionFilter

`func (o *View) SetFileExtensionFilter(v map[string]interface{})`

SetFileExtensionFilter sets FileExtensionFilter field to given value.

### HasFileExtensionFilter

`func (o *View) HasFileExtensionFilter() bool`

HasFileExtensionFilter returns a boolean if a field has been set.

### GetFileLockConfig

`func (o *View) GetFileLockConfig() map[string]interface{}`

GetFileLockConfig returns the FileLockConfig field if non-nil, zero value otherwise.

### GetFileLockConfigOk

`func (o *View) GetFileLockConfigOk() (*map[string]interface{}, bool)`

GetFileLockConfigOk returns a tuple with the FileLockConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLockConfig

`func (o *View) SetFileLockConfig(v map[string]interface{})`

SetFileLockConfig sets FileLockConfig field to given value.

### HasFileLockConfig

`func (o *View) HasFileLockConfig() bool`

HasFileLockConfig returns a boolean if a field has been set.

### GetFilerLifecycleManagement

`func (o *View) GetFilerLifecycleManagement() map[string]interface{}`

GetFilerLifecycleManagement returns the FilerLifecycleManagement field if non-nil, zero value otherwise.

### GetFilerLifecycleManagementOk

`func (o *View) GetFilerLifecycleManagementOk() (*map[string]interface{}, bool)`

GetFilerLifecycleManagementOk returns a tuple with the FilerLifecycleManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilerLifecycleManagement

`func (o *View) SetFilerLifecycleManagement(v map[string]interface{})`

SetFilerLifecycleManagement sets FilerLifecycleManagement field to given value.

### HasFilerLifecycleManagement

`func (o *View) HasFilerLifecycleManagement() bool`

HasFilerLifecycleManagement returns a boolean if a field has been set.

### GetIsExternallyTriggeredBackupTarget

`func (o *View) GetIsExternallyTriggeredBackupTarget() bool`

GetIsExternallyTriggeredBackupTarget returns the IsExternallyTriggeredBackupTarget field if non-nil, zero value otherwise.

### GetIsExternallyTriggeredBackupTargetOk

`func (o *View) GetIsExternallyTriggeredBackupTargetOk() (*bool, bool)`

GetIsExternallyTriggeredBackupTargetOk returns a tuple with the IsExternallyTriggeredBackupTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternallyTriggeredBackupTarget

`func (o *View) SetIsExternallyTriggeredBackupTarget(v bool)`

SetIsExternallyTriggeredBackupTarget sets IsExternallyTriggeredBackupTarget field to given value.

### HasIsExternallyTriggeredBackupTarget

`func (o *View) HasIsExternallyTriggeredBackupTarget() bool`

HasIsExternallyTriggeredBackupTarget returns a boolean if a field has been set.

### SetIsExternallyTriggeredBackupTargetNil

`func (o *View) SetIsExternallyTriggeredBackupTargetNil(b bool)`

 SetIsExternallyTriggeredBackupTargetNil sets the value for IsExternallyTriggeredBackupTarget to be an explicit nil

### UnsetIsExternallyTriggeredBackupTarget
`func (o *View) UnsetIsExternallyTriggeredBackupTarget()`

UnsetIsExternallyTriggeredBackupTarget ensures that no value is present for IsExternallyTriggeredBackupTarget, not even an explicit nil
### GetIsReadOnly

`func (o *View) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *View) GetIsReadOnlyOk() (*bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadOnly

`func (o *View) SetIsReadOnly(v bool)`

SetIsReadOnly sets IsReadOnly field to given value.

### HasIsReadOnly

`func (o *View) HasIsReadOnly() bool`

HasIsReadOnly returns a boolean if a field has been set.

### SetIsReadOnlyNil

`func (o *View) SetIsReadOnlyNil(b bool)`

 SetIsReadOnlyNil sets the value for IsReadOnly to be an explicit nil

### UnsetIsReadOnly
`func (o *View) UnsetIsReadOnly()`

UnsetIsReadOnly ensures that no value is present for IsReadOnly, not even an explicit nil
### GetLexicographicPrefetch

`func (o *View) GetLexicographicPrefetch() bool`

GetLexicographicPrefetch returns the LexicographicPrefetch field if non-nil, zero value otherwise.

### GetLexicographicPrefetchOk

`func (o *View) GetLexicographicPrefetchOk() (*bool, bool)`

GetLexicographicPrefetchOk returns a tuple with the LexicographicPrefetch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLexicographicPrefetch

`func (o *View) SetLexicographicPrefetch(v bool)`

SetLexicographicPrefetch sets LexicographicPrefetch field to given value.

### HasLexicographicPrefetch

`func (o *View) HasLexicographicPrefetch() bool`

HasLexicographicPrefetch returns a boolean if a field has been set.

### SetLexicographicPrefetchNil

`func (o *View) SetLexicographicPrefetchNil(b bool)`

 SetLexicographicPrefetchNil sets the value for LexicographicPrefetch to be an explicit nil

### UnsetLexicographicPrefetch
`func (o *View) UnsetLexicographicPrefetch()`

UnsetLexicographicPrefetch ensures that no value is present for LexicographicPrefetch, not even an explicit nil
### GetLogicalQuota

`func (o *View) GetLogicalQuota() map[string]interface{}`

GetLogicalQuota returns the LogicalQuota field if non-nil, zero value otherwise.

### GetLogicalQuotaOk

`func (o *View) GetLogicalQuotaOk() (*map[string]interface{}, bool)`

GetLogicalQuotaOk returns a tuple with the LogicalQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalQuota

`func (o *View) SetLogicalQuota(v map[string]interface{})`

SetLogicalQuota sets LogicalQuota field to given value.

### HasLogicalQuota

`func (o *View) HasLogicalQuota() bool`

HasLogicalQuota returns a boolean if a field has been set.

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

`func (o *View) GetNetgroupWhitelist() map[string]interface{}`

GetNetgroupWhitelist returns the NetgroupWhitelist field if non-nil, zero value otherwise.

### GetNetgroupWhitelistOk

`func (o *View) GetNetgroupWhitelistOk() (*map[string]interface{}, bool)`

GetNetgroupWhitelistOk returns a tuple with the NetgroupWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetgroupWhitelist

`func (o *View) SetNetgroupWhitelist(v map[string]interface{})`

SetNetgroupWhitelist sets NetgroupWhitelist field to given value.

### HasNetgroupWhitelist

`func (o *View) HasNetgroupWhitelist() bool`

HasNetgroupWhitelist returns a boolean if a field has been set.

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
### GetOverrideGlobalSubnetWhitelist

`func (o *View) GetOverrideGlobalSubnetWhitelist() bool`

GetOverrideGlobalSubnetWhitelist returns the OverrideGlobalSubnetWhitelist field if non-nil, zero value otherwise.

### GetOverrideGlobalSubnetWhitelistOk

`func (o *View) GetOverrideGlobalSubnetWhitelistOk() (*bool, bool)`

GetOverrideGlobalSubnetWhitelistOk returns a tuple with the OverrideGlobalSubnetWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideGlobalSubnetWhitelist

`func (o *View) SetOverrideGlobalSubnetWhitelist(v bool)`

SetOverrideGlobalSubnetWhitelist sets OverrideGlobalSubnetWhitelist field to given value.

### HasOverrideGlobalSubnetWhitelist

`func (o *View) HasOverrideGlobalSubnetWhitelist() bool`

HasOverrideGlobalSubnetWhitelist returns a boolean if a field has been set.

### SetOverrideGlobalSubnetWhitelistNil

`func (o *View) SetOverrideGlobalSubnetWhitelistNil(b bool)`

 SetOverrideGlobalSubnetWhitelistNil sets the value for OverrideGlobalSubnetWhitelist to be an explicit nil

### UnsetOverrideGlobalSubnetWhitelist
`func (o *View) UnsetOverrideGlobalSubnetWhitelist()`

UnsetOverrideGlobalSubnetWhitelist ensures that no value is present for OverrideGlobalSubnetWhitelist, not even an explicit nil
### GetProtocolAccess

`func (o *View) GetProtocolAccess() []ViewProtocol`

GetProtocolAccess returns the ProtocolAccess field if non-nil, zero value otherwise.

### GetProtocolAccessOk

`func (o *View) GetProtocolAccessOk() (*[]ViewProtocol, bool)`

GetProtocolAccessOk returns a tuple with the ProtocolAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolAccess

`func (o *View) SetProtocolAccess(v []ViewProtocol)`

SetProtocolAccess sets ProtocolAccess field to given value.


### SetProtocolAccessNil

`func (o *View) SetProtocolAccessNil(b bool)`

 SetProtocolAccessNil sets the value for ProtocolAccess to be an explicit nil

### UnsetProtocolAccess
`func (o *View) UnsetProtocolAccess()`

UnsetProtocolAccess ensures that no value is present for ProtocolAccess, not even an explicit nil
### GetQos

`func (o *View) GetQos() map[string]interface{}`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *View) GetQosOk() (*map[string]interface{}, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *View) SetQos(v map[string]interface{})`

SetQos sets Qos field to given value.


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
### GetSelfServiceSnapshotConfig

`func (o *View) GetSelfServiceSnapshotConfig() map[string]interface{}`

GetSelfServiceSnapshotConfig returns the SelfServiceSnapshotConfig field if non-nil, zero value otherwise.

### GetSelfServiceSnapshotConfigOk

`func (o *View) GetSelfServiceSnapshotConfigOk() (*map[string]interface{}, bool)`

GetSelfServiceSnapshotConfigOk returns a tuple with the SelfServiceSnapshotConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfServiceSnapshotConfig

`func (o *View) SetSelfServiceSnapshotConfig(v map[string]interface{})`

SetSelfServiceSnapshotConfig sets SelfServiceSnapshotConfig field to given value.

### HasSelfServiceSnapshotConfig

`func (o *View) HasSelfServiceSnapshotConfig() bool`

HasSelfServiceSnapshotConfig returns a boolean if a field has been set.

### GetStoragePolicyOverride

`func (o *View) GetStoragePolicyOverride() map[string]interface{}`

GetStoragePolicyOverride returns the StoragePolicyOverride field if non-nil, zero value otherwise.

### GetStoragePolicyOverrideOk

`func (o *View) GetStoragePolicyOverrideOk() (*map[string]interface{}, bool)`

GetStoragePolicyOverrideOk returns a tuple with the StoragePolicyOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePolicyOverride

`func (o *View) SetStoragePolicyOverride(v map[string]interface{})`

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
### GetViewPinningConfig

`func (o *View) GetViewPinningConfig() map[string]interface{}`

GetViewPinningConfig returns the ViewPinningConfig field if non-nil, zero value otherwise.

### GetViewPinningConfigOk

`func (o *View) GetViewPinningConfigOk() (*map[string]interface{}, bool)`

GetViewPinningConfigOk returns a tuple with the ViewPinningConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewPinningConfig

`func (o *View) SetViewPinningConfig(v map[string]interface{})`

SetViewPinningConfig sets ViewPinningConfig field to given value.

### HasViewPinningConfig

`func (o *View) HasViewPinningConfig() bool`

HasViewPinningConfig returns a boolean if a field has been set.

### GetEnableNfsKerberosAuthentication

`func (o *View) GetEnableNfsKerberosAuthentication() bool`

GetEnableNfsKerberosAuthentication returns the EnableNfsKerberosAuthentication field if non-nil, zero value otherwise.

### GetEnableNfsKerberosAuthenticationOk

`func (o *View) GetEnableNfsKerberosAuthenticationOk() (*bool, bool)`

GetEnableNfsKerberosAuthenticationOk returns a tuple with the EnableNfsKerberosAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNfsKerberosAuthentication

`func (o *View) SetEnableNfsKerberosAuthentication(v bool)`

SetEnableNfsKerberosAuthentication sets EnableNfsKerberosAuthentication field to given value.

### HasEnableNfsKerberosAuthentication

`func (o *View) HasEnableNfsKerberosAuthentication() bool`

HasEnableNfsKerberosAuthentication returns a boolean if a field has been set.

### SetEnableNfsKerberosAuthenticationNil

`func (o *View) SetEnableNfsKerberosAuthenticationNil(b bool)`

 SetEnableNfsKerberosAuthenticationNil sets the value for EnableNfsKerberosAuthentication to be an explicit nil

### UnsetEnableNfsKerberosAuthentication
`func (o *View) UnsetEnableNfsKerberosAuthentication()`

UnsetEnableNfsKerberosAuthentication ensures that no value is present for EnableNfsKerberosAuthentication, not even an explicit nil
### GetEnableNfsKerberosIntegrity

`func (o *View) GetEnableNfsKerberosIntegrity() bool`

GetEnableNfsKerberosIntegrity returns the EnableNfsKerberosIntegrity field if non-nil, zero value otherwise.

### GetEnableNfsKerberosIntegrityOk

`func (o *View) GetEnableNfsKerberosIntegrityOk() (*bool, bool)`

GetEnableNfsKerberosIntegrityOk returns a tuple with the EnableNfsKerberosIntegrity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNfsKerberosIntegrity

`func (o *View) SetEnableNfsKerberosIntegrity(v bool)`

SetEnableNfsKerberosIntegrity sets EnableNfsKerberosIntegrity field to given value.

### HasEnableNfsKerberosIntegrity

`func (o *View) HasEnableNfsKerberosIntegrity() bool`

HasEnableNfsKerberosIntegrity returns a boolean if a field has been set.

### SetEnableNfsKerberosIntegrityNil

`func (o *View) SetEnableNfsKerberosIntegrityNil(b bool)`

 SetEnableNfsKerberosIntegrityNil sets the value for EnableNfsKerberosIntegrity to be an explicit nil

### UnsetEnableNfsKerberosIntegrity
`func (o *View) UnsetEnableNfsKerberosIntegrity()`

UnsetEnableNfsKerberosIntegrity ensures that no value is present for EnableNfsKerberosIntegrity, not even an explicit nil
### GetEnableNfsKerberosPrivacy

`func (o *View) GetEnableNfsKerberosPrivacy() bool`

GetEnableNfsKerberosPrivacy returns the EnableNfsKerberosPrivacy field if non-nil, zero value otherwise.

### GetEnableNfsKerberosPrivacyOk

`func (o *View) GetEnableNfsKerberosPrivacyOk() (*bool, bool)`

GetEnableNfsKerberosPrivacyOk returns a tuple with the EnableNfsKerberosPrivacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNfsKerberosPrivacy

`func (o *View) SetEnableNfsKerberosPrivacy(v bool)`

SetEnableNfsKerberosPrivacy sets EnableNfsKerberosPrivacy field to given value.

### HasEnableNfsKerberosPrivacy

`func (o *View) HasEnableNfsKerberosPrivacy() bool`

HasEnableNfsKerberosPrivacy returns a boolean if a field has been set.

### SetEnableNfsKerberosPrivacyNil

`func (o *View) SetEnableNfsKerberosPrivacyNil(b bool)`

 SetEnableNfsKerberosPrivacyNil sets the value for EnableNfsKerberosPrivacy to be an explicit nil

### UnsetEnableNfsKerberosPrivacy
`func (o *View) UnsetEnableNfsKerberosPrivacy()`

UnsetEnableNfsKerberosPrivacy ensures that no value is present for EnableNfsKerberosPrivacy, not even an explicit nil
### GetEnableNfsUnixAuthentication

`func (o *View) GetEnableNfsUnixAuthentication() bool`

GetEnableNfsUnixAuthentication returns the EnableNfsUnixAuthentication field if non-nil, zero value otherwise.

### GetEnableNfsUnixAuthenticationOk

`func (o *View) GetEnableNfsUnixAuthenticationOk() (*bool, bool)`

GetEnableNfsUnixAuthenticationOk returns a tuple with the EnableNfsUnixAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNfsUnixAuthentication

`func (o *View) SetEnableNfsUnixAuthentication(v bool)`

SetEnableNfsUnixAuthentication sets EnableNfsUnixAuthentication field to given value.

### HasEnableNfsUnixAuthentication

`func (o *View) HasEnableNfsUnixAuthentication() bool`

HasEnableNfsUnixAuthentication returns a boolean if a field has been set.

### SetEnableNfsUnixAuthenticationNil

`func (o *View) SetEnableNfsUnixAuthenticationNil(b bool)`

 SetEnableNfsUnixAuthenticationNil sets the value for EnableNfsUnixAuthentication to be an explicit nil

### UnsetEnableNfsUnixAuthentication
`func (o *View) UnsetEnableNfsUnixAuthentication()`

UnsetEnableNfsUnixAuthentication ensures that no value is present for EnableNfsUnixAuthentication, not even an explicit nil
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
### GetEnableNfsWcc

`func (o *View) GetEnableNfsWcc() bool`

GetEnableNfsWcc returns the EnableNfsWcc field if non-nil, zero value otherwise.

### GetEnableNfsWccOk

`func (o *View) GetEnableNfsWccOk() (*bool, bool)`

GetEnableNfsWccOk returns a tuple with the EnableNfsWcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNfsWcc

`func (o *View) SetEnableNfsWcc(v bool)`

SetEnableNfsWcc sets EnableNfsWcc field to given value.

### HasEnableNfsWcc

`func (o *View) HasEnableNfsWcc() bool`

HasEnableNfsWcc returns a boolean if a field has been set.

### SetEnableNfsWccNil

`func (o *View) SetEnableNfsWccNil(b bool)`

 SetEnableNfsWccNil sets the value for EnableNfsWcc to be an explicit nil

### UnsetEnableNfsWcc
`func (o *View) UnsetEnableNfsWcc()`

UnsetEnableNfsWcc ensures that no value is present for EnableNfsWcc, not even an explicit nil
### GetNfsAllSquash

`func (o *View) GetNfsAllSquash() NfsConfigNfsAllSquash`

GetNfsAllSquash returns the NfsAllSquash field if non-nil, zero value otherwise.

### GetNfsAllSquashOk

`func (o *View) GetNfsAllSquashOk() (*NfsConfigNfsAllSquash, bool)`

GetNfsAllSquashOk returns a tuple with the NfsAllSquash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsAllSquash

`func (o *View) SetNfsAllSquash(v NfsConfigNfsAllSquash)`

SetNfsAllSquash sets NfsAllSquash field to given value.

### HasNfsAllSquash

`func (o *View) HasNfsAllSquash() bool`

HasNfsAllSquash returns a boolean if a field has been set.

### GetNfsRootPermissions

`func (o *View) GetNfsRootPermissions() NfsConfigNfsRootPermissions`

GetNfsRootPermissions returns the NfsRootPermissions field if non-nil, zero value otherwise.

### GetNfsRootPermissionsOk

`func (o *View) GetNfsRootPermissionsOk() (*NfsConfigNfsRootPermissions, bool)`

GetNfsRootPermissionsOk returns a tuple with the NfsRootPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsRootPermissions

`func (o *View) SetNfsRootPermissions(v NfsConfigNfsRootPermissions)`

SetNfsRootPermissions sets NfsRootPermissions field to given value.

### HasNfsRootPermissions

`func (o *View) HasNfsRootPermissions() bool`

HasNfsRootPermissions returns a boolean if a field has been set.

### GetNfsRootSquash

`func (o *View) GetNfsRootSquash() NfsConfigNfsRootSquash`

GetNfsRootSquash returns the NfsRootSquash field if non-nil, zero value otherwise.

### GetNfsRootSquashOk

`func (o *View) GetNfsRootSquashOk() (*NfsConfigNfsRootSquash, bool)`

GetNfsRootSquashOk returns a tuple with the NfsRootSquash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsRootSquash

`func (o *View) SetNfsRootSquash(v NfsConfigNfsRootSquash)`

SetNfsRootSquash sets NfsRootSquash field to given value.

### HasNfsRootSquash

`func (o *View) HasNfsRootSquash() bool`

HasNfsRootSquash returns a boolean if a field has been set.

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
### GetSharePermissions

`func (o *View) GetSharePermissions() SmbConfigSharePermissions`

GetSharePermissions returns the SharePermissions field if non-nil, zero value otherwise.

### GetSharePermissionsOk

`func (o *View) GetSharePermissionsOk() (*SmbConfigSharePermissions, bool)`

GetSharePermissionsOk returns a tuple with the SharePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharePermissions

`func (o *View) SetSharePermissions(v SmbConfigSharePermissions)`

SetSharePermissions sets SharePermissions field to given value.

### HasSharePermissions

`func (o *View) HasSharePermissions() bool`

HasSharePermissions returns a boolean if a field has been set.

### GetSmbPermissionsInfo

`func (o *View) GetSmbPermissionsInfo() SmbConfigSmbPermissionsInfo`

GetSmbPermissionsInfo returns the SmbPermissionsInfo field if non-nil, zero value otherwise.

### GetSmbPermissionsInfoOk

`func (o *View) GetSmbPermissionsInfoOk() (*SmbConfigSmbPermissionsInfo, bool)`

GetSmbPermissionsInfoOk returns a tuple with the SmbPermissionsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbPermissionsInfo

`func (o *View) SetSmbPermissionsInfo(v SmbConfigSmbPermissionsInfo)`

SetSmbPermissionsInfo sets SmbPermissionsInfo field to given value.

### HasSmbPermissionsInfo

`func (o *View) HasSmbPermissionsInfo() bool`

HasSmbPermissionsInfo returns a boolean if a field has been set.

### GetAclConfig

`func (o *View) GetAclConfig() S3ConfigAclConfig`

GetAclConfig returns the AclConfig field if non-nil, zero value otherwise.

### GetAclConfigOk

`func (o *View) GetAclConfigOk() (*S3ConfigAclConfig, bool)`

GetAclConfigOk returns a tuple with the AclConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclConfig

`func (o *View) SetAclConfig(v S3ConfigAclConfig)`

SetAclConfig sets AclConfig field to given value.

### HasAclConfig

`func (o *View) HasAclConfig() bool`

HasAclConfig returns a boolean if a field has been set.

### GetBucketPolicy

`func (o *View) GetBucketPolicy() S3ConfigBucketPolicy`

GetBucketPolicy returns the BucketPolicy field if non-nil, zero value otherwise.

### GetBucketPolicyOk

`func (o *View) GetBucketPolicyOk() (*S3ConfigBucketPolicy, bool)`

GetBucketPolicyOk returns a tuple with the BucketPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketPolicy

`func (o *View) SetBucketPolicy(v S3ConfigBucketPolicy)`

SetBucketPolicy sets BucketPolicy field to given value.

### HasBucketPolicy

`func (o *View) HasBucketPolicy() bool`

HasBucketPolicy returns a boolean if a field has been set.

### GetEnableAbac

`func (o *View) GetEnableAbac() bool`

GetEnableAbac returns the EnableAbac field if non-nil, zero value otherwise.

### GetEnableAbacOk

`func (o *View) GetEnableAbacOk() (*bool, bool)`

GetEnableAbacOk returns a tuple with the EnableAbac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAbac

`func (o *View) SetEnableAbac(v bool)`

SetEnableAbac sets EnableAbac field to given value.

### HasEnableAbac

`func (o *View) HasEnableAbac() bool`

HasEnableAbac returns a boolean if a field has been set.

### SetEnableAbacNil

`func (o *View) SetEnableAbacNil(b bool)`

 SetEnableAbacNil sets the value for EnableAbac to be an explicit nil

### UnsetEnableAbac
`func (o *View) UnsetEnableAbac()`

UnsetEnableAbac ensures that no value is present for EnableAbac, not even an explicit nil
### GetLifecycleManagement

`func (o *View) GetLifecycleManagement() S3ConfigLifecycleManagement`

GetLifecycleManagement returns the LifecycleManagement field if non-nil, zero value otherwise.

### GetLifecycleManagementOk

`func (o *View) GetLifecycleManagementOk() (*S3ConfigLifecycleManagement, bool)`

GetLifecycleManagementOk returns a tuple with the LifecycleManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleManagement

`func (o *View) SetLifecycleManagement(v S3ConfigLifecycleManagement)`

SetLifecycleManagement sets LifecycleManagement field to given value.

### HasLifecycleManagement

`func (o *View) HasLifecycleManagement() bool`

HasLifecycleManagement returns a boolean if a field has been set.

### GetOwnerInfo

`func (o *View) GetOwnerInfo() S3ConfigOwnerInfo`

GetOwnerInfo returns the OwnerInfo field if non-nil, zero value otherwise.

### GetOwnerInfoOk

`func (o *View) GetOwnerInfoOk() (*S3ConfigOwnerInfo, bool)`

GetOwnerInfoOk returns a tuple with the OwnerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerInfo

`func (o *View) SetOwnerInfo(v S3ConfigOwnerInfo)`

SetOwnerInfo sets OwnerInfo field to given value.

### HasOwnerInfo

`func (o *View) HasOwnerInfo() bool`

HasOwnerInfo returns a boolean if a field has been set.

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
### GetS3EfficientMpuMaxSubfiles

`func (o *View) GetS3EfficientMpuMaxSubfiles() int32`

GetS3EfficientMpuMaxSubfiles returns the S3EfficientMpuMaxSubfiles field if non-nil, zero value otherwise.

### GetS3EfficientMpuMaxSubfilesOk

`func (o *View) GetS3EfficientMpuMaxSubfilesOk() (*int32, bool)`

GetS3EfficientMpuMaxSubfilesOk returns a tuple with the S3EfficientMpuMaxSubfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3EfficientMpuMaxSubfiles

`func (o *View) SetS3EfficientMpuMaxSubfiles(v int32)`

SetS3EfficientMpuMaxSubfiles sets S3EfficientMpuMaxSubfiles field to given value.

### HasS3EfficientMpuMaxSubfiles

`func (o *View) HasS3EfficientMpuMaxSubfiles() bool`

HasS3EfficientMpuMaxSubfiles returns a boolean if a field has been set.

### SetS3EfficientMpuMaxSubfilesNil

`func (o *View) SetS3EfficientMpuMaxSubfilesNil(b bool)`

 SetS3EfficientMpuMaxSubfilesNil sets the value for S3EfficientMpuMaxSubfiles to be an explicit nil

### UnsetS3EfficientMpuMaxSubfiles
`func (o *View) UnsetS3EfficientMpuMaxSubfiles()`

UnsetS3EfficientMpuMaxSubfiles ensures that no value is present for S3EfficientMpuMaxSubfiles, not even an explicit nil
### GetS3EnableEfficientMpu

`func (o *View) GetS3EnableEfficientMpu() bool`

GetS3EnableEfficientMpu returns the S3EnableEfficientMpu field if non-nil, zero value otherwise.

### GetS3EnableEfficientMpuOk

`func (o *View) GetS3EnableEfficientMpuOk() (*bool, bool)`

GetS3EnableEfficientMpuOk returns a tuple with the S3EnableEfficientMpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3EnableEfficientMpu

`func (o *View) SetS3EnableEfficientMpu(v bool)`

SetS3EnableEfficientMpu sets S3EnableEfficientMpu field to given value.

### HasS3EnableEfficientMpu

`func (o *View) HasS3EnableEfficientMpu() bool`

HasS3EnableEfficientMpu returns a boolean if a field has been set.

### SetS3EnableEfficientMpuNil

`func (o *View) SetS3EnableEfficientMpuNil(b bool)`

 SetS3EnableEfficientMpuNil sets the value for S3EnableEfficientMpu to be an explicit nil

### UnsetS3EnableEfficientMpu
`func (o *View) UnsetS3EnableEfficientMpu()`

UnsetS3EnableEfficientMpu ensures that no value is present for S3EnableEfficientMpu, not even an explicit nil
### GetS3MigrationAction

`func (o *View) GetS3MigrationAction() string`

GetS3MigrationAction returns the S3MigrationAction field if non-nil, zero value otherwise.

### GetS3MigrationActionOk

`func (o *View) GetS3MigrationActionOk() (*string, bool)`

GetS3MigrationActionOk returns a tuple with the S3MigrationAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3MigrationAction

`func (o *View) SetS3MigrationAction(v string)`

SetS3MigrationAction sets S3MigrationAction field to given value.

### HasS3MigrationAction

`func (o *View) HasS3MigrationAction() bool`

HasS3MigrationAction returns a boolean if a field has been set.

### SetS3MigrationActionNil

`func (o *View) SetS3MigrationActionNil(b bool)`

 SetS3MigrationActionNil sets the value for S3MigrationAction to be an explicit nil

### UnsetS3MigrationAction
`func (o *View) UnsetS3MigrationAction()`

UnsetS3MigrationAction ensures that no value is present for S3MigrationAction, not even an explicit nil
### GetS3MigrationState

`func (o *View) GetS3MigrationState() string`

GetS3MigrationState returns the S3MigrationState field if non-nil, zero value otherwise.

### GetS3MigrationStateOk

`func (o *View) GetS3MigrationStateOk() (*string, bool)`

GetS3MigrationStateOk returns a tuple with the S3MigrationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3MigrationState

`func (o *View) SetS3MigrationState(v string)`

SetS3MigrationState sets S3MigrationState field to given value.

### HasS3MigrationState

`func (o *View) HasS3MigrationState() bool`

HasS3MigrationState returns a boolean if a field has been set.

### SetS3MigrationStateNil

`func (o *View) SetS3MigrationStateNil(b bool)`

 SetS3MigrationStateNil sets the value for S3MigrationState to be an explicit nil

### UnsetS3MigrationState
`func (o *View) UnsetS3MigrationState()`

UnsetS3MigrationState ensures that no value is present for S3MigrationState, not even an explicit nil
### GetVersioning

`func (o *View) GetVersioning() string`

GetVersioning returns the Versioning field if non-nil, zero value otherwise.

### GetVersioningOk

`func (o *View) GetVersioningOk() (*string, bool)`

GetVersioningOk returns a tuple with the Versioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersioning

`func (o *View) SetVersioning(v string)`

SetVersioning sets Versioning field to given value.

### HasVersioning

`func (o *View) HasVersioning() bool`

HasVersioning returns a boolean if a field has been set.

### SetVersioningNil

`func (o *View) SetVersioningNil(b bool)`

 SetVersioningNil sets the value for Versioning to be an explicit nil

### UnsetVersioning
`func (o *View) UnsetVersioning()`

UnsetVersioning ensures that no value is present for Versioning, not even an explicit nil
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
### GetAliases

`func (o *View) GetAliases() []ViewAliasInfo`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *View) GetAliasesOk() (*[]ViewAliasInfo, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *View) SetAliases(v []ViewAliasInfo)`

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
### GetFileCountBySize

`func (o *View) GetFileCountBySize() []FileCount`

GetFileCountBySize returns the FileCountBySize field if non-nil, zero value otherwise.

### GetFileCountBySizeOk

`func (o *View) GetFileCountBySizeOk() (*[]FileCount, bool)`

GetFileCountBySizeOk returns a tuple with the FileCountBySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileCountBySize

`func (o *View) SetFileCountBySize(v []FileCount)`

SetFileCountBySize sets FileCountBySize field to given value.

### HasFileCountBySize

`func (o *View) HasFileCountBySize() bool`

HasFileCountBySize returns a boolean if a field has been set.

### SetFileCountBySizeNil

`func (o *View) SetFileCountBySizeNil(b bool)`

 SetFileCountBySizeNil sets the value for FileCountBySize to be an explicit nil

### UnsetFileCountBySize
`func (o *View) UnsetFileCountBySize()`

UnsetFileCountBySize ensures that no value is present for FileCountBySize, not even an explicit nil
### GetIntent

`func (o *View) GetIntent() map[string]interface{}`

GetIntent returns the Intent field if non-nil, zero value otherwise.

### GetIntentOk

`func (o *View) GetIntentOk() (*map[string]interface{}, bool)`

GetIntentOk returns a tuple with the Intent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntent

`func (o *View) SetIntent(v map[string]interface{})`

SetIntent sets Intent field to given value.

### HasIntent

`func (o *View) HasIntent() bool`

HasIntent returns a boolean if a field has been set.

### GetIsCategoryInferred

`func (o *View) GetIsCategoryInferred() bool`

GetIsCategoryInferred returns the IsCategoryInferred field if non-nil, zero value otherwise.

### GetIsCategoryInferredOk

`func (o *View) GetIsCategoryInferredOk() (*bool, bool)`

GetIsCategoryInferredOk returns a tuple with the IsCategoryInferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCategoryInferred

`func (o *View) SetIsCategoryInferred(v bool)`

SetIsCategoryInferred sets IsCategoryInferred field to given value.

### HasIsCategoryInferred

`func (o *View) HasIsCategoryInferred() bool`

HasIsCategoryInferred returns a boolean if a field has been set.

### SetIsCategoryInferredNil

`func (o *View) SetIsCategoryInferredNil(b bool)`

 SetIsCategoryInferredNil sets the value for IsCategoryInferred to be an explicit nil

### UnsetIsCategoryInferred
`func (o *View) UnsetIsCategoryInferred()`

UnsetIsCategoryInferred ensures that no value is present for IsCategoryInferred, not even an explicit nil
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
### GetNfsMountPaths

`func (o *View) GetNfsMountPaths() []string`

GetNfsMountPaths returns the NfsMountPaths field if non-nil, zero value otherwise.

### GetNfsMountPathsOk

`func (o *View) GetNfsMountPathsOk() (*[]string, bool)`

GetNfsMountPathsOk returns a tuple with the NfsMountPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsMountPaths

`func (o *View) SetNfsMountPaths(v []string)`

SetNfsMountPaths sets NfsMountPaths field to given value.

### HasNfsMountPaths

`func (o *View) HasNfsMountPaths() bool`

HasNfsMountPaths returns a boolean if a field has been set.

### SetNfsMountPathsNil

`func (o *View) SetNfsMountPathsNil(b bool)`

 SetNfsMountPathsNil sets the value for NfsMountPaths to be an explicit nil

### UnsetNfsMountPaths
`func (o *View) UnsetNfsMountPaths()`

UnsetNfsMountPaths ensures that no value is present for NfsMountPaths, not even an explicit nil
### GetObjectServicesMappingConfig

`func (o *View) GetObjectServicesMappingConfig() string`

GetObjectServicesMappingConfig returns the ObjectServicesMappingConfig field if non-nil, zero value otherwise.

### GetObjectServicesMappingConfigOk

`func (o *View) GetObjectServicesMappingConfigOk() (*string, bool)`

GetObjectServicesMappingConfigOk returns a tuple with the ObjectServicesMappingConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectServicesMappingConfig

`func (o *View) SetObjectServicesMappingConfig(v string)`

SetObjectServicesMappingConfig sets ObjectServicesMappingConfig field to given value.

### HasObjectServicesMappingConfig

`func (o *View) HasObjectServicesMappingConfig() bool`

HasObjectServicesMappingConfig returns a boolean if a field has been set.

### SetObjectServicesMappingConfigNil

`func (o *View) SetObjectServicesMappingConfigNil(b bool)`

 SetObjectServicesMappingConfigNil sets the value for ObjectServicesMappingConfig to be an explicit nil

### UnsetObjectServicesMappingConfig
`func (o *View) UnsetObjectServicesMappingConfig()`

UnsetObjectServicesMappingConfig ensures that no value is present for ObjectServicesMappingConfig, not even an explicit nil
### GetOwnerSid

`func (o *View) GetOwnerSid() string`

GetOwnerSid returns the OwnerSid field if non-nil, zero value otherwise.

### GetOwnerSidOk

`func (o *View) GetOwnerSidOk() (*string, bool)`

GetOwnerSidOk returns a tuple with the OwnerSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerSid

`func (o *View) SetOwnerSid(v string)`

SetOwnerSid sets OwnerSid field to given value.

### HasOwnerSid

`func (o *View) HasOwnerSid() bool`

HasOwnerSid returns a boolean if a field has been set.

### SetOwnerSidNil

`func (o *View) SetOwnerSidNil(b bool)`

 SetOwnerSidNil sets the value for OwnerSid to be an explicit nil

### UnsetOwnerSid
`func (o *View) UnsetOwnerSid()`

UnsetOwnerSid ensures that no value is present for OwnerSid, not even an explicit nil
### GetS3FolderSupportEnabled

`func (o *View) GetS3FolderSupportEnabled() bool`

GetS3FolderSupportEnabled returns the S3FolderSupportEnabled field if non-nil, zero value otherwise.

### GetS3FolderSupportEnabledOk

`func (o *View) GetS3FolderSupportEnabledOk() (*bool, bool)`

GetS3FolderSupportEnabledOk returns a tuple with the S3FolderSupportEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3FolderSupportEnabled

`func (o *View) SetS3FolderSupportEnabled(v bool)`

SetS3FolderSupportEnabled sets S3FolderSupportEnabled field to given value.

### HasS3FolderSupportEnabled

`func (o *View) HasS3FolderSupportEnabled() bool`

HasS3FolderSupportEnabled returns a boolean if a field has been set.

### SetS3FolderSupportEnabledNil

`func (o *View) SetS3FolderSupportEnabledNil(b bool)`

 SetS3FolderSupportEnabledNil sets the value for S3FolderSupportEnabled to be an explicit nil

### UnsetS3FolderSupportEnabled
`func (o *View) UnsetS3FolderSupportEnabled()`

UnsetS3FolderSupportEnabled ensures that no value is present for S3FolderSupportEnabled, not even an explicit nil
### GetSmbMountPaths

`func (o *View) GetSmbMountPaths() []string`

GetSmbMountPaths returns the SmbMountPaths field if non-nil, zero value otherwise.

### GetSmbMountPathsOk

`func (o *View) GetSmbMountPathsOk() (*[]string, bool)`

GetSmbMountPathsOk returns a tuple with the SmbMountPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbMountPaths

`func (o *View) SetSmbMountPaths(v []string)`

SetSmbMountPaths sets SmbMountPaths field to given value.

### HasSmbMountPaths

`func (o *View) HasSmbMountPaths() bool`

HasSmbMountPaths returns a boolean if a field has been set.

### SetSmbMountPathsNil

`func (o *View) SetSmbMountPathsNil(b bool)`

 SetSmbMountPathsNil sets the value for SmbMountPaths to be an explicit nil

### UnsetSmbMountPaths
`func (o *View) UnsetSmbMountPaths()`

UnsetSmbMountPaths ensures that no value is present for SmbMountPaths, not even an explicit nil
### GetStats

`func (o *View) GetStats() map[string]interface{}`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *View) GetStatsOk() (*map[string]interface{}, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *View) SetStats(v map[string]interface{})`

SetStats sets Stats field to given value.

### HasStats

`func (o *View) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetStorageDomainId

`func (o *View) GetStorageDomainId() int64`

GetStorageDomainId returns the StorageDomainId field if non-nil, zero value otherwise.

### GetStorageDomainIdOk

`func (o *View) GetStorageDomainIdOk() (*int64, bool)`

GetStorageDomainIdOk returns a tuple with the StorageDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainId

`func (o *View) SetStorageDomainId(v int64)`

SetStorageDomainId sets StorageDomainId field to given value.

### HasStorageDomainId

`func (o *View) HasStorageDomainId() bool`

HasStorageDomainId returns a boolean if a field has been set.

### SetStorageDomainIdNil

`func (o *View) SetStorageDomainIdNil(b bool)`

 SetStorageDomainIdNil sets the value for StorageDomainId to be an explicit nil

### UnsetStorageDomainId
`func (o *View) UnsetStorageDomainId()`

UnsetStorageDomainId ensures that no value is present for StorageDomainId, not even an explicit nil
### GetStorageDomainName

`func (o *View) GetStorageDomainName() string`

GetStorageDomainName returns the StorageDomainName field if non-nil, zero value otherwise.

### GetStorageDomainNameOk

`func (o *View) GetStorageDomainNameOk() (*string, bool)`

GetStorageDomainNameOk returns a tuple with the StorageDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainName

`func (o *View) SetStorageDomainName(v string)`

SetStorageDomainName sets StorageDomainName field to given value.

### HasStorageDomainName

`func (o *View) HasStorageDomainName() bool`

HasStorageDomainName returns a boolean if a field has been set.

### SetStorageDomainNameNil

`func (o *View) SetStorageDomainNameNil(b bool)`

 SetStorageDomainNameNil sets the value for StorageDomainName to be an explicit nil

### UnsetStorageDomainName
`func (o *View) UnsetStorageDomainName()`

UnsetStorageDomainName ensures that no value is present for StorageDomainName, not even an explicit nil
### GetViewFailover

`func (o *View) GetViewFailover() map[string]interface{}`

GetViewFailover returns the ViewFailover field if non-nil, zero value otherwise.

### GetViewFailoverOk

`func (o *View) GetViewFailoverOk() (*map[string]interface{}, bool)`

GetViewFailoverOk returns a tuple with the ViewFailover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewFailover

`func (o *View) SetViewFailover(v map[string]interface{})`

SetViewFailover sets ViewFailover field to given value.

### HasViewFailover

`func (o *View) HasViewFailover() bool`

HasViewFailover returns a boolean if a field has been set.

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
### GetViewProtection

`func (o *View) GetViewProtection() map[string]interface{}`

GetViewProtection returns the ViewProtection field if non-nil, zero value otherwise.

### GetViewProtectionOk

`func (o *View) GetViewProtectionOk() (*map[string]interface{}, bool)`

GetViewProtectionOk returns a tuple with the ViewProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewProtection

`func (o *View) SetViewProtection(v map[string]interface{})`

SetViewProtection sets ViewProtection field to given value.

### HasViewProtection

`func (o *View) HasViewProtection() bool`

HasViewProtection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


