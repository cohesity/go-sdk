# TemplateViewParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaseInsensitiveNamesEnabled** | Pointer to **NullableBool** | Specifies whether to support case insensitive file/folder names. This parameter can only be set during create and cannot be changed. | [optional] 
**Intent** | Pointer to **map[string]interface{}** | Specifies the intent of the View. | [optional] 
**ObjectServicesMappingConfig** | Pointer to **NullableString** | Specifies the Object Services key mapping config of the view. This parameter can only be set during create and cannot be changed. Configuration of Object Services key mapping. Specifies the type of Object Services key mapping config. | [optional] 
**S3FolderSupportEnabled** | Pointer to **NullableBool** | Specifies whether to support s3 folder support feature. This parameter can only be set during create and cannot be changed. | [optional] 
**StorageDomainId** | **NullableInt64** | Specifies the id of the Storage Domain (View Box) where the View will be created. | 
**ViewProtectionConfig** | Pointer to **map[string]interface{}** | Specifies the protection config of the View. | [optional] 
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

## Methods

### NewTemplateViewParams

`func NewTemplateViewParams(storageDomainId NullableInt64, category NullableString, protocolAccess []ViewProtocol, qos map[string]interface{}, ) *TemplateViewParams`

NewTemplateViewParams instantiates a new TemplateViewParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateViewParamsWithDefaults

`func NewTemplateViewParamsWithDefaults() *TemplateViewParams`

NewTemplateViewParamsWithDefaults instantiates a new TemplateViewParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaseInsensitiveNamesEnabled

`func (o *TemplateViewParams) GetCaseInsensitiveNamesEnabled() bool`

GetCaseInsensitiveNamesEnabled returns the CaseInsensitiveNamesEnabled field if non-nil, zero value otherwise.

### GetCaseInsensitiveNamesEnabledOk

`func (o *TemplateViewParams) GetCaseInsensitiveNamesEnabledOk() (*bool, bool)`

GetCaseInsensitiveNamesEnabledOk returns a tuple with the CaseInsensitiveNamesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseInsensitiveNamesEnabled

`func (o *TemplateViewParams) SetCaseInsensitiveNamesEnabled(v bool)`

SetCaseInsensitiveNamesEnabled sets CaseInsensitiveNamesEnabled field to given value.

### HasCaseInsensitiveNamesEnabled

`func (o *TemplateViewParams) HasCaseInsensitiveNamesEnabled() bool`

HasCaseInsensitiveNamesEnabled returns a boolean if a field has been set.

### SetCaseInsensitiveNamesEnabledNil

`func (o *TemplateViewParams) SetCaseInsensitiveNamesEnabledNil(b bool)`

 SetCaseInsensitiveNamesEnabledNil sets the value for CaseInsensitiveNamesEnabled to be an explicit nil

### UnsetCaseInsensitiveNamesEnabled
`func (o *TemplateViewParams) UnsetCaseInsensitiveNamesEnabled()`

UnsetCaseInsensitiveNamesEnabled ensures that no value is present for CaseInsensitiveNamesEnabled, not even an explicit nil
### GetIntent

`func (o *TemplateViewParams) GetIntent() map[string]interface{}`

GetIntent returns the Intent field if non-nil, zero value otherwise.

### GetIntentOk

`func (o *TemplateViewParams) GetIntentOk() (*map[string]interface{}, bool)`

GetIntentOk returns a tuple with the Intent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntent

`func (o *TemplateViewParams) SetIntent(v map[string]interface{})`

SetIntent sets Intent field to given value.

### HasIntent

`func (o *TemplateViewParams) HasIntent() bool`

HasIntent returns a boolean if a field has been set.

### GetObjectServicesMappingConfig

`func (o *TemplateViewParams) GetObjectServicesMappingConfig() string`

GetObjectServicesMappingConfig returns the ObjectServicesMappingConfig field if non-nil, zero value otherwise.

### GetObjectServicesMappingConfigOk

`func (o *TemplateViewParams) GetObjectServicesMappingConfigOk() (*string, bool)`

GetObjectServicesMappingConfigOk returns a tuple with the ObjectServicesMappingConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectServicesMappingConfig

`func (o *TemplateViewParams) SetObjectServicesMappingConfig(v string)`

SetObjectServicesMappingConfig sets ObjectServicesMappingConfig field to given value.

### HasObjectServicesMappingConfig

`func (o *TemplateViewParams) HasObjectServicesMappingConfig() bool`

HasObjectServicesMappingConfig returns a boolean if a field has been set.

### SetObjectServicesMappingConfigNil

`func (o *TemplateViewParams) SetObjectServicesMappingConfigNil(b bool)`

 SetObjectServicesMappingConfigNil sets the value for ObjectServicesMappingConfig to be an explicit nil

### UnsetObjectServicesMappingConfig
`func (o *TemplateViewParams) UnsetObjectServicesMappingConfig()`

UnsetObjectServicesMappingConfig ensures that no value is present for ObjectServicesMappingConfig, not even an explicit nil
### GetS3FolderSupportEnabled

`func (o *TemplateViewParams) GetS3FolderSupportEnabled() bool`

GetS3FolderSupportEnabled returns the S3FolderSupportEnabled field if non-nil, zero value otherwise.

### GetS3FolderSupportEnabledOk

`func (o *TemplateViewParams) GetS3FolderSupportEnabledOk() (*bool, bool)`

GetS3FolderSupportEnabledOk returns a tuple with the S3FolderSupportEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3FolderSupportEnabled

`func (o *TemplateViewParams) SetS3FolderSupportEnabled(v bool)`

SetS3FolderSupportEnabled sets S3FolderSupportEnabled field to given value.

### HasS3FolderSupportEnabled

`func (o *TemplateViewParams) HasS3FolderSupportEnabled() bool`

HasS3FolderSupportEnabled returns a boolean if a field has been set.

### SetS3FolderSupportEnabledNil

`func (o *TemplateViewParams) SetS3FolderSupportEnabledNil(b bool)`

 SetS3FolderSupportEnabledNil sets the value for S3FolderSupportEnabled to be an explicit nil

### UnsetS3FolderSupportEnabled
`func (o *TemplateViewParams) UnsetS3FolderSupportEnabled()`

UnsetS3FolderSupportEnabled ensures that no value is present for S3FolderSupportEnabled, not even an explicit nil
### GetStorageDomainId

`func (o *TemplateViewParams) GetStorageDomainId() int64`

GetStorageDomainId returns the StorageDomainId field if non-nil, zero value otherwise.

### GetStorageDomainIdOk

`func (o *TemplateViewParams) GetStorageDomainIdOk() (*int64, bool)`

GetStorageDomainIdOk returns a tuple with the StorageDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainId

`func (o *TemplateViewParams) SetStorageDomainId(v int64)`

SetStorageDomainId sets StorageDomainId field to given value.


### SetStorageDomainIdNil

`func (o *TemplateViewParams) SetStorageDomainIdNil(b bool)`

 SetStorageDomainIdNil sets the value for StorageDomainId to be an explicit nil

### UnsetStorageDomainId
`func (o *TemplateViewParams) UnsetStorageDomainId()`

UnsetStorageDomainId ensures that no value is present for StorageDomainId, not even an explicit nil
### GetViewProtectionConfig

`func (o *TemplateViewParams) GetViewProtectionConfig() map[string]interface{}`

GetViewProtectionConfig returns the ViewProtectionConfig field if non-nil, zero value otherwise.

### GetViewProtectionConfigOk

`func (o *TemplateViewParams) GetViewProtectionConfigOk() (*map[string]interface{}, bool)`

GetViewProtectionConfigOk returns a tuple with the ViewProtectionConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewProtectionConfig

`func (o *TemplateViewParams) SetViewProtectionConfig(v map[string]interface{})`

SetViewProtectionConfig sets ViewProtectionConfig field to given value.

### HasViewProtectionConfig

`func (o *TemplateViewParams) HasViewProtectionConfig() bool`

HasViewProtectionConfig returns a boolean if a field has been set.

### GetAccessSids

`func (o *TemplateViewParams) GetAccessSids() []string`

GetAccessSids returns the AccessSids field if non-nil, zero value otherwise.

### GetAccessSidsOk

`func (o *TemplateViewParams) GetAccessSidsOk() (*[]string, bool)`

GetAccessSidsOk returns a tuple with the AccessSids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessSids

`func (o *TemplateViewParams) SetAccessSids(v []string)`

SetAccessSids sets AccessSids field to given value.

### HasAccessSids

`func (o *TemplateViewParams) HasAccessSids() bool`

HasAccessSids returns a boolean if a field has been set.

### SetAccessSidsNil

`func (o *TemplateViewParams) SetAccessSidsNil(b bool)`

 SetAccessSidsNil sets the value for AccessSids to be an explicit nil

### UnsetAccessSids
`func (o *TemplateViewParams) UnsetAccessSids()`

UnsetAccessSids ensures that no value is present for AccessSids, not even an explicit nil
### GetAllowMountOnWindows

`func (o *TemplateViewParams) GetAllowMountOnWindows() bool`

GetAllowMountOnWindows returns the AllowMountOnWindows field if non-nil, zero value otherwise.

### GetAllowMountOnWindowsOk

`func (o *TemplateViewParams) GetAllowMountOnWindowsOk() (*bool, bool)`

GetAllowMountOnWindowsOk returns a tuple with the AllowMountOnWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMountOnWindows

`func (o *TemplateViewParams) SetAllowMountOnWindows(v bool)`

SetAllowMountOnWindows sets AllowMountOnWindows field to given value.

### HasAllowMountOnWindows

`func (o *TemplateViewParams) HasAllowMountOnWindows() bool`

HasAllowMountOnWindows returns a boolean if a field has been set.

### SetAllowMountOnWindowsNil

`func (o *TemplateViewParams) SetAllowMountOnWindowsNil(b bool)`

 SetAllowMountOnWindowsNil sets the value for AllowMountOnWindows to be an explicit nil

### UnsetAllowMountOnWindows
`func (o *TemplateViewParams) UnsetAllowMountOnWindows()`

UnsetAllowMountOnWindows ensures that no value is present for AllowMountOnWindows, not even an explicit nil
### GetAntivirusScanConfig

`func (o *TemplateViewParams) GetAntivirusScanConfig() map[string]interface{}`

GetAntivirusScanConfig returns the AntivirusScanConfig field if non-nil, zero value otherwise.

### GetAntivirusScanConfigOk

`func (o *TemplateViewParams) GetAntivirusScanConfigOk() (*map[string]interface{}, bool)`

GetAntivirusScanConfigOk returns a tuple with the AntivirusScanConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntivirusScanConfig

`func (o *TemplateViewParams) SetAntivirusScanConfig(v map[string]interface{})`

SetAntivirusScanConfig sets AntivirusScanConfig field to given value.

### HasAntivirusScanConfig

`func (o *TemplateViewParams) HasAntivirusScanConfig() bool`

HasAntivirusScanConfig returns a boolean if a field has been set.

### GetCategory

`func (o *TemplateViewParams) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *TemplateViewParams) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *TemplateViewParams) SetCategory(v string)`

SetCategory sets Category field to given value.


### SetCategoryNil

`func (o *TemplateViewParams) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *TemplateViewParams) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetDescription

`func (o *TemplateViewParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TemplateViewParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TemplateViewParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TemplateViewParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TemplateViewParams) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TemplateViewParams) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnableFilerAuditLogging

`func (o *TemplateViewParams) GetEnableFilerAuditLogging() bool`

GetEnableFilerAuditLogging returns the EnableFilerAuditLogging field if non-nil, zero value otherwise.

### GetEnableFilerAuditLoggingOk

`func (o *TemplateViewParams) GetEnableFilerAuditLoggingOk() (*bool, bool)`

GetEnableFilerAuditLoggingOk returns a tuple with the EnableFilerAuditLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFilerAuditLogging

`func (o *TemplateViewParams) SetEnableFilerAuditLogging(v bool)`

SetEnableFilerAuditLogging sets EnableFilerAuditLogging field to given value.

### HasEnableFilerAuditLogging

`func (o *TemplateViewParams) HasEnableFilerAuditLogging() bool`

HasEnableFilerAuditLogging returns a boolean if a field has been set.

### SetEnableFilerAuditLoggingNil

`func (o *TemplateViewParams) SetEnableFilerAuditLoggingNil(b bool)`

 SetEnableFilerAuditLoggingNil sets the value for EnableFilerAuditLogging to be an explicit nil

### UnsetEnableFilerAuditLogging
`func (o *TemplateViewParams) UnsetEnableFilerAuditLogging()`

UnsetEnableFilerAuditLogging ensures that no value is present for EnableFilerAuditLogging, not even an explicit nil
### GetEnableLiveIndexing

`func (o *TemplateViewParams) GetEnableLiveIndexing() bool`

GetEnableLiveIndexing returns the EnableLiveIndexing field if non-nil, zero value otherwise.

### GetEnableLiveIndexingOk

`func (o *TemplateViewParams) GetEnableLiveIndexingOk() (*bool, bool)`

GetEnableLiveIndexingOk returns a tuple with the EnableLiveIndexing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLiveIndexing

`func (o *TemplateViewParams) SetEnableLiveIndexing(v bool)`

SetEnableLiveIndexing sets EnableLiveIndexing field to given value.

### HasEnableLiveIndexing

`func (o *TemplateViewParams) HasEnableLiveIndexing() bool`

HasEnableLiveIndexing returns a boolean if a field has been set.

### SetEnableLiveIndexingNil

`func (o *TemplateViewParams) SetEnableLiveIndexingNil(b bool)`

 SetEnableLiveIndexingNil sets the value for EnableLiveIndexing to be an explicit nil

### UnsetEnableLiveIndexing
`func (o *TemplateViewParams) UnsetEnableLiveIndexing()`

UnsetEnableLiveIndexing ensures that no value is present for EnableLiveIndexing, not even an explicit nil
### GetEnableMetadataAccelerator

`func (o *TemplateViewParams) GetEnableMetadataAccelerator() bool`

GetEnableMetadataAccelerator returns the EnableMetadataAccelerator field if non-nil, zero value otherwise.

### GetEnableMetadataAcceleratorOk

`func (o *TemplateViewParams) GetEnableMetadataAcceleratorOk() (*bool, bool)`

GetEnableMetadataAcceleratorOk returns a tuple with the EnableMetadataAccelerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMetadataAccelerator

`func (o *TemplateViewParams) SetEnableMetadataAccelerator(v bool)`

SetEnableMetadataAccelerator sets EnableMetadataAccelerator field to given value.

### HasEnableMetadataAccelerator

`func (o *TemplateViewParams) HasEnableMetadataAccelerator() bool`

HasEnableMetadataAccelerator returns a boolean if a field has been set.

### SetEnableMetadataAcceleratorNil

`func (o *TemplateViewParams) SetEnableMetadataAcceleratorNil(b bool)`

 SetEnableMetadataAcceleratorNil sets the value for EnableMetadataAccelerator to be an explicit nil

### UnsetEnableMetadataAccelerator
`func (o *TemplateViewParams) UnsetEnableMetadataAccelerator()`

UnsetEnableMetadataAccelerator ensures that no value is present for EnableMetadataAccelerator, not even an explicit nil
### GetEnableMinion

`func (o *TemplateViewParams) GetEnableMinion() bool`

GetEnableMinion returns the EnableMinion field if non-nil, zero value otherwise.

### GetEnableMinionOk

`func (o *TemplateViewParams) GetEnableMinionOk() (*bool, bool)`

GetEnableMinionOk returns a tuple with the EnableMinion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMinion

`func (o *TemplateViewParams) SetEnableMinion(v bool)`

SetEnableMinion sets EnableMinion field to given value.

### HasEnableMinion

`func (o *TemplateViewParams) HasEnableMinion() bool`

HasEnableMinion returns a boolean if a field has been set.

### SetEnableMinionNil

`func (o *TemplateViewParams) SetEnableMinionNil(b bool)`

 SetEnableMinionNil sets the value for EnableMinion to be an explicit nil

### UnsetEnableMinion
`func (o *TemplateViewParams) UnsetEnableMinion()`

UnsetEnableMinion ensures that no value is present for EnableMinion, not even an explicit nil
### GetEnableOfflineCaching

`func (o *TemplateViewParams) GetEnableOfflineCaching() bool`

GetEnableOfflineCaching returns the EnableOfflineCaching field if non-nil, zero value otherwise.

### GetEnableOfflineCachingOk

`func (o *TemplateViewParams) GetEnableOfflineCachingOk() (*bool, bool)`

GetEnableOfflineCachingOk returns a tuple with the EnableOfflineCaching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableOfflineCaching

`func (o *TemplateViewParams) SetEnableOfflineCaching(v bool)`

SetEnableOfflineCaching sets EnableOfflineCaching field to given value.

### HasEnableOfflineCaching

`func (o *TemplateViewParams) HasEnableOfflineCaching() bool`

HasEnableOfflineCaching returns a boolean if a field has been set.

### SetEnableOfflineCachingNil

`func (o *TemplateViewParams) SetEnableOfflineCachingNil(b bool)`

 SetEnableOfflineCachingNil sets the value for EnableOfflineCaching to be an explicit nil

### UnsetEnableOfflineCaching
`func (o *TemplateViewParams) UnsetEnableOfflineCaching()`

UnsetEnableOfflineCaching ensures that no value is present for EnableOfflineCaching, not even an explicit nil
### GetFileExtensionFilter

`func (o *TemplateViewParams) GetFileExtensionFilter() map[string]interface{}`

GetFileExtensionFilter returns the FileExtensionFilter field if non-nil, zero value otherwise.

### GetFileExtensionFilterOk

`func (o *TemplateViewParams) GetFileExtensionFilterOk() (*map[string]interface{}, bool)`

GetFileExtensionFilterOk returns a tuple with the FileExtensionFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileExtensionFilter

`func (o *TemplateViewParams) SetFileExtensionFilter(v map[string]interface{})`

SetFileExtensionFilter sets FileExtensionFilter field to given value.

### HasFileExtensionFilter

`func (o *TemplateViewParams) HasFileExtensionFilter() bool`

HasFileExtensionFilter returns a boolean if a field has been set.

### GetFileLockConfig

`func (o *TemplateViewParams) GetFileLockConfig() map[string]interface{}`

GetFileLockConfig returns the FileLockConfig field if non-nil, zero value otherwise.

### GetFileLockConfigOk

`func (o *TemplateViewParams) GetFileLockConfigOk() (*map[string]interface{}, bool)`

GetFileLockConfigOk returns a tuple with the FileLockConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLockConfig

`func (o *TemplateViewParams) SetFileLockConfig(v map[string]interface{})`

SetFileLockConfig sets FileLockConfig field to given value.

### HasFileLockConfig

`func (o *TemplateViewParams) HasFileLockConfig() bool`

HasFileLockConfig returns a boolean if a field has been set.

### GetFilerLifecycleManagement

`func (o *TemplateViewParams) GetFilerLifecycleManagement() map[string]interface{}`

GetFilerLifecycleManagement returns the FilerLifecycleManagement field if non-nil, zero value otherwise.

### GetFilerLifecycleManagementOk

`func (o *TemplateViewParams) GetFilerLifecycleManagementOk() (*map[string]interface{}, bool)`

GetFilerLifecycleManagementOk returns a tuple with the FilerLifecycleManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilerLifecycleManagement

`func (o *TemplateViewParams) SetFilerLifecycleManagement(v map[string]interface{})`

SetFilerLifecycleManagement sets FilerLifecycleManagement field to given value.

### HasFilerLifecycleManagement

`func (o *TemplateViewParams) HasFilerLifecycleManagement() bool`

HasFilerLifecycleManagement returns a boolean if a field has been set.

### GetIsExternallyTriggeredBackupTarget

`func (o *TemplateViewParams) GetIsExternallyTriggeredBackupTarget() bool`

GetIsExternallyTriggeredBackupTarget returns the IsExternallyTriggeredBackupTarget field if non-nil, zero value otherwise.

### GetIsExternallyTriggeredBackupTargetOk

`func (o *TemplateViewParams) GetIsExternallyTriggeredBackupTargetOk() (*bool, bool)`

GetIsExternallyTriggeredBackupTargetOk returns a tuple with the IsExternallyTriggeredBackupTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternallyTriggeredBackupTarget

`func (o *TemplateViewParams) SetIsExternallyTriggeredBackupTarget(v bool)`

SetIsExternallyTriggeredBackupTarget sets IsExternallyTriggeredBackupTarget field to given value.

### HasIsExternallyTriggeredBackupTarget

`func (o *TemplateViewParams) HasIsExternallyTriggeredBackupTarget() bool`

HasIsExternallyTriggeredBackupTarget returns a boolean if a field has been set.

### SetIsExternallyTriggeredBackupTargetNil

`func (o *TemplateViewParams) SetIsExternallyTriggeredBackupTargetNil(b bool)`

 SetIsExternallyTriggeredBackupTargetNil sets the value for IsExternallyTriggeredBackupTarget to be an explicit nil

### UnsetIsExternallyTriggeredBackupTarget
`func (o *TemplateViewParams) UnsetIsExternallyTriggeredBackupTarget()`

UnsetIsExternallyTriggeredBackupTarget ensures that no value is present for IsExternallyTriggeredBackupTarget, not even an explicit nil
### GetIsReadOnly

`func (o *TemplateViewParams) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *TemplateViewParams) GetIsReadOnlyOk() (*bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadOnly

`func (o *TemplateViewParams) SetIsReadOnly(v bool)`

SetIsReadOnly sets IsReadOnly field to given value.

### HasIsReadOnly

`func (o *TemplateViewParams) HasIsReadOnly() bool`

HasIsReadOnly returns a boolean if a field has been set.

### SetIsReadOnlyNil

`func (o *TemplateViewParams) SetIsReadOnlyNil(b bool)`

 SetIsReadOnlyNil sets the value for IsReadOnly to be an explicit nil

### UnsetIsReadOnly
`func (o *TemplateViewParams) UnsetIsReadOnly()`

UnsetIsReadOnly ensures that no value is present for IsReadOnly, not even an explicit nil
### GetLexicographicPrefetch

`func (o *TemplateViewParams) GetLexicographicPrefetch() bool`

GetLexicographicPrefetch returns the LexicographicPrefetch field if non-nil, zero value otherwise.

### GetLexicographicPrefetchOk

`func (o *TemplateViewParams) GetLexicographicPrefetchOk() (*bool, bool)`

GetLexicographicPrefetchOk returns a tuple with the LexicographicPrefetch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLexicographicPrefetch

`func (o *TemplateViewParams) SetLexicographicPrefetch(v bool)`

SetLexicographicPrefetch sets LexicographicPrefetch field to given value.

### HasLexicographicPrefetch

`func (o *TemplateViewParams) HasLexicographicPrefetch() bool`

HasLexicographicPrefetch returns a boolean if a field has been set.

### SetLexicographicPrefetchNil

`func (o *TemplateViewParams) SetLexicographicPrefetchNil(b bool)`

 SetLexicographicPrefetchNil sets the value for LexicographicPrefetch to be an explicit nil

### UnsetLexicographicPrefetch
`func (o *TemplateViewParams) UnsetLexicographicPrefetch()`

UnsetLexicographicPrefetch ensures that no value is present for LexicographicPrefetch, not even an explicit nil
### GetLogicalQuota

`func (o *TemplateViewParams) GetLogicalQuota() map[string]interface{}`

GetLogicalQuota returns the LogicalQuota field if non-nil, zero value otherwise.

### GetLogicalQuotaOk

`func (o *TemplateViewParams) GetLogicalQuotaOk() (*map[string]interface{}, bool)`

GetLogicalQuotaOk returns a tuple with the LogicalQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalQuota

`func (o *TemplateViewParams) SetLogicalQuota(v map[string]interface{})`

SetLogicalQuota sets LogicalQuota field to given value.

### HasLogicalQuota

`func (o *TemplateViewParams) HasLogicalQuota() bool`

HasLogicalQuota returns a boolean if a field has been set.

### GetName

`func (o *TemplateViewParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplateViewParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplateViewParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TemplateViewParams) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TemplateViewParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TemplateViewParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNetgroupWhitelist

`func (o *TemplateViewParams) GetNetgroupWhitelist() map[string]interface{}`

GetNetgroupWhitelist returns the NetgroupWhitelist field if non-nil, zero value otherwise.

### GetNetgroupWhitelistOk

`func (o *TemplateViewParams) GetNetgroupWhitelistOk() (*map[string]interface{}, bool)`

GetNetgroupWhitelistOk returns a tuple with the NetgroupWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetgroupWhitelist

`func (o *TemplateViewParams) SetNetgroupWhitelist(v map[string]interface{})`

SetNetgroupWhitelist sets NetgroupWhitelist field to given value.

### HasNetgroupWhitelist

`func (o *TemplateViewParams) HasNetgroupWhitelist() bool`

HasNetgroupWhitelist returns a boolean if a field has been set.

### GetOverrideGlobalNetgroupWhitelist

`func (o *TemplateViewParams) GetOverrideGlobalNetgroupWhitelist() bool`

GetOverrideGlobalNetgroupWhitelist returns the OverrideGlobalNetgroupWhitelist field if non-nil, zero value otherwise.

### GetOverrideGlobalNetgroupWhitelistOk

`func (o *TemplateViewParams) GetOverrideGlobalNetgroupWhitelistOk() (*bool, bool)`

GetOverrideGlobalNetgroupWhitelistOk returns a tuple with the OverrideGlobalNetgroupWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideGlobalNetgroupWhitelist

`func (o *TemplateViewParams) SetOverrideGlobalNetgroupWhitelist(v bool)`

SetOverrideGlobalNetgroupWhitelist sets OverrideGlobalNetgroupWhitelist field to given value.

### HasOverrideGlobalNetgroupWhitelist

`func (o *TemplateViewParams) HasOverrideGlobalNetgroupWhitelist() bool`

HasOverrideGlobalNetgroupWhitelist returns a boolean if a field has been set.

### SetOverrideGlobalNetgroupWhitelistNil

`func (o *TemplateViewParams) SetOverrideGlobalNetgroupWhitelistNil(b bool)`

 SetOverrideGlobalNetgroupWhitelistNil sets the value for OverrideGlobalNetgroupWhitelist to be an explicit nil

### UnsetOverrideGlobalNetgroupWhitelist
`func (o *TemplateViewParams) UnsetOverrideGlobalNetgroupWhitelist()`

UnsetOverrideGlobalNetgroupWhitelist ensures that no value is present for OverrideGlobalNetgroupWhitelist, not even an explicit nil
### GetOverrideGlobalSubnetWhitelist

`func (o *TemplateViewParams) GetOverrideGlobalSubnetWhitelist() bool`

GetOverrideGlobalSubnetWhitelist returns the OverrideGlobalSubnetWhitelist field if non-nil, zero value otherwise.

### GetOverrideGlobalSubnetWhitelistOk

`func (o *TemplateViewParams) GetOverrideGlobalSubnetWhitelistOk() (*bool, bool)`

GetOverrideGlobalSubnetWhitelistOk returns a tuple with the OverrideGlobalSubnetWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideGlobalSubnetWhitelist

`func (o *TemplateViewParams) SetOverrideGlobalSubnetWhitelist(v bool)`

SetOverrideGlobalSubnetWhitelist sets OverrideGlobalSubnetWhitelist field to given value.

### HasOverrideGlobalSubnetWhitelist

`func (o *TemplateViewParams) HasOverrideGlobalSubnetWhitelist() bool`

HasOverrideGlobalSubnetWhitelist returns a boolean if a field has been set.

### SetOverrideGlobalSubnetWhitelistNil

`func (o *TemplateViewParams) SetOverrideGlobalSubnetWhitelistNil(b bool)`

 SetOverrideGlobalSubnetWhitelistNil sets the value for OverrideGlobalSubnetWhitelist to be an explicit nil

### UnsetOverrideGlobalSubnetWhitelist
`func (o *TemplateViewParams) UnsetOverrideGlobalSubnetWhitelist()`

UnsetOverrideGlobalSubnetWhitelist ensures that no value is present for OverrideGlobalSubnetWhitelist, not even an explicit nil
### GetProtocolAccess

`func (o *TemplateViewParams) GetProtocolAccess() []ViewProtocol`

GetProtocolAccess returns the ProtocolAccess field if non-nil, zero value otherwise.

### GetProtocolAccessOk

`func (o *TemplateViewParams) GetProtocolAccessOk() (*[]ViewProtocol, bool)`

GetProtocolAccessOk returns a tuple with the ProtocolAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolAccess

`func (o *TemplateViewParams) SetProtocolAccess(v []ViewProtocol)`

SetProtocolAccess sets ProtocolAccess field to given value.


### SetProtocolAccessNil

`func (o *TemplateViewParams) SetProtocolAccessNil(b bool)`

 SetProtocolAccessNil sets the value for ProtocolAccess to be an explicit nil

### UnsetProtocolAccess
`func (o *TemplateViewParams) UnsetProtocolAccess()`

UnsetProtocolAccess ensures that no value is present for ProtocolAccess, not even an explicit nil
### GetQos

`func (o *TemplateViewParams) GetQos() map[string]interface{}`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *TemplateViewParams) GetQosOk() (*map[string]interface{}, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *TemplateViewParams) SetQos(v map[string]interface{})`

SetQos sets Qos field to given value.


### GetSecurityMode

`func (o *TemplateViewParams) GetSecurityMode() string`

GetSecurityMode returns the SecurityMode field if non-nil, zero value otherwise.

### GetSecurityModeOk

`func (o *TemplateViewParams) GetSecurityModeOk() (*string, bool)`

GetSecurityModeOk returns a tuple with the SecurityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityMode

`func (o *TemplateViewParams) SetSecurityMode(v string)`

SetSecurityMode sets SecurityMode field to given value.

### HasSecurityMode

`func (o *TemplateViewParams) HasSecurityMode() bool`

HasSecurityMode returns a boolean if a field has been set.

### SetSecurityModeNil

`func (o *TemplateViewParams) SetSecurityModeNil(b bool)`

 SetSecurityModeNil sets the value for SecurityMode to be an explicit nil

### UnsetSecurityMode
`func (o *TemplateViewParams) UnsetSecurityMode()`

UnsetSecurityMode ensures that no value is present for SecurityMode, not even an explicit nil
### GetSelfServiceSnapshotConfig

`func (o *TemplateViewParams) GetSelfServiceSnapshotConfig() map[string]interface{}`

GetSelfServiceSnapshotConfig returns the SelfServiceSnapshotConfig field if non-nil, zero value otherwise.

### GetSelfServiceSnapshotConfigOk

`func (o *TemplateViewParams) GetSelfServiceSnapshotConfigOk() (*map[string]interface{}, bool)`

GetSelfServiceSnapshotConfigOk returns a tuple with the SelfServiceSnapshotConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfServiceSnapshotConfig

`func (o *TemplateViewParams) SetSelfServiceSnapshotConfig(v map[string]interface{})`

SetSelfServiceSnapshotConfig sets SelfServiceSnapshotConfig field to given value.

### HasSelfServiceSnapshotConfig

`func (o *TemplateViewParams) HasSelfServiceSnapshotConfig() bool`

HasSelfServiceSnapshotConfig returns a boolean if a field has been set.

### GetStoragePolicyOverride

`func (o *TemplateViewParams) GetStoragePolicyOverride() map[string]interface{}`

GetStoragePolicyOverride returns the StoragePolicyOverride field if non-nil, zero value otherwise.

### GetStoragePolicyOverrideOk

`func (o *TemplateViewParams) GetStoragePolicyOverrideOk() (*map[string]interface{}, bool)`

GetStoragePolicyOverrideOk returns a tuple with the StoragePolicyOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePolicyOverride

`func (o *TemplateViewParams) SetStoragePolicyOverride(v map[string]interface{})`

SetStoragePolicyOverride sets StoragePolicyOverride field to given value.

### HasStoragePolicyOverride

`func (o *TemplateViewParams) HasStoragePolicyOverride() bool`

HasStoragePolicyOverride returns a boolean if a field has been set.

### GetSubnetWhitelist

`func (o *TemplateViewParams) GetSubnetWhitelist() []Subnet`

GetSubnetWhitelist returns the SubnetWhitelist field if non-nil, zero value otherwise.

### GetSubnetWhitelistOk

`func (o *TemplateViewParams) GetSubnetWhitelistOk() (*[]Subnet, bool)`

GetSubnetWhitelistOk returns a tuple with the SubnetWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetWhitelist

`func (o *TemplateViewParams) SetSubnetWhitelist(v []Subnet)`

SetSubnetWhitelist sets SubnetWhitelist field to given value.

### HasSubnetWhitelist

`func (o *TemplateViewParams) HasSubnetWhitelist() bool`

HasSubnetWhitelist returns a boolean if a field has been set.

### SetSubnetWhitelistNil

`func (o *TemplateViewParams) SetSubnetWhitelistNil(b bool)`

 SetSubnetWhitelistNil sets the value for SubnetWhitelist to be an explicit nil

### UnsetSubnetWhitelist
`func (o *TemplateViewParams) UnsetSubnetWhitelist()`

UnsetSubnetWhitelist ensures that no value is present for SubnetWhitelist, not even an explicit nil
### GetTenantId

`func (o *TemplateViewParams) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TemplateViewParams) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TemplateViewParams) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *TemplateViewParams) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *TemplateViewParams) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *TemplateViewParams) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetViewLockEnabled

`func (o *TemplateViewParams) GetViewLockEnabled() bool`

GetViewLockEnabled returns the ViewLockEnabled field if non-nil, zero value otherwise.

### GetViewLockEnabledOk

`func (o *TemplateViewParams) GetViewLockEnabledOk() (*bool, bool)`

GetViewLockEnabledOk returns a tuple with the ViewLockEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewLockEnabled

`func (o *TemplateViewParams) SetViewLockEnabled(v bool)`

SetViewLockEnabled sets ViewLockEnabled field to given value.

### HasViewLockEnabled

`func (o *TemplateViewParams) HasViewLockEnabled() bool`

HasViewLockEnabled returns a boolean if a field has been set.

### SetViewLockEnabledNil

`func (o *TemplateViewParams) SetViewLockEnabledNil(b bool)`

 SetViewLockEnabledNil sets the value for ViewLockEnabled to be an explicit nil

### UnsetViewLockEnabled
`func (o *TemplateViewParams) UnsetViewLockEnabled()`

UnsetViewLockEnabled ensures that no value is present for ViewLockEnabled, not even an explicit nil
### GetViewPinningConfig

`func (o *TemplateViewParams) GetViewPinningConfig() map[string]interface{}`

GetViewPinningConfig returns the ViewPinningConfig field if non-nil, zero value otherwise.

### GetViewPinningConfigOk

`func (o *TemplateViewParams) GetViewPinningConfigOk() (*map[string]interface{}, bool)`

GetViewPinningConfigOk returns a tuple with the ViewPinningConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewPinningConfig

`func (o *TemplateViewParams) SetViewPinningConfig(v map[string]interface{})`

SetViewPinningConfig sets ViewPinningConfig field to given value.

### HasViewPinningConfig

`func (o *TemplateViewParams) HasViewPinningConfig() bool`

HasViewPinningConfig returns a boolean if a field has been set.

### GetEnableNfsKerberosAuthentication

`func (o *TemplateViewParams) GetEnableNfsKerberosAuthentication() bool`

GetEnableNfsKerberosAuthentication returns the EnableNfsKerberosAuthentication field if non-nil, zero value otherwise.

### GetEnableNfsKerberosAuthenticationOk

`func (o *TemplateViewParams) GetEnableNfsKerberosAuthenticationOk() (*bool, bool)`

GetEnableNfsKerberosAuthenticationOk returns a tuple with the EnableNfsKerberosAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNfsKerberosAuthentication

`func (o *TemplateViewParams) SetEnableNfsKerberosAuthentication(v bool)`

SetEnableNfsKerberosAuthentication sets EnableNfsKerberosAuthentication field to given value.

### HasEnableNfsKerberosAuthentication

`func (o *TemplateViewParams) HasEnableNfsKerberosAuthentication() bool`

HasEnableNfsKerberosAuthentication returns a boolean if a field has been set.

### SetEnableNfsKerberosAuthenticationNil

`func (o *TemplateViewParams) SetEnableNfsKerberosAuthenticationNil(b bool)`

 SetEnableNfsKerberosAuthenticationNil sets the value for EnableNfsKerberosAuthentication to be an explicit nil

### UnsetEnableNfsKerberosAuthentication
`func (o *TemplateViewParams) UnsetEnableNfsKerberosAuthentication()`

UnsetEnableNfsKerberosAuthentication ensures that no value is present for EnableNfsKerberosAuthentication, not even an explicit nil
### GetEnableNfsKerberosIntegrity

`func (o *TemplateViewParams) GetEnableNfsKerberosIntegrity() bool`

GetEnableNfsKerberosIntegrity returns the EnableNfsKerberosIntegrity field if non-nil, zero value otherwise.

### GetEnableNfsKerberosIntegrityOk

`func (o *TemplateViewParams) GetEnableNfsKerberosIntegrityOk() (*bool, bool)`

GetEnableNfsKerberosIntegrityOk returns a tuple with the EnableNfsKerberosIntegrity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNfsKerberosIntegrity

`func (o *TemplateViewParams) SetEnableNfsKerberosIntegrity(v bool)`

SetEnableNfsKerberosIntegrity sets EnableNfsKerberosIntegrity field to given value.

### HasEnableNfsKerberosIntegrity

`func (o *TemplateViewParams) HasEnableNfsKerberosIntegrity() bool`

HasEnableNfsKerberosIntegrity returns a boolean if a field has been set.

### SetEnableNfsKerberosIntegrityNil

`func (o *TemplateViewParams) SetEnableNfsKerberosIntegrityNil(b bool)`

 SetEnableNfsKerberosIntegrityNil sets the value for EnableNfsKerberosIntegrity to be an explicit nil

### UnsetEnableNfsKerberosIntegrity
`func (o *TemplateViewParams) UnsetEnableNfsKerberosIntegrity()`

UnsetEnableNfsKerberosIntegrity ensures that no value is present for EnableNfsKerberosIntegrity, not even an explicit nil
### GetEnableNfsKerberosPrivacy

`func (o *TemplateViewParams) GetEnableNfsKerberosPrivacy() bool`

GetEnableNfsKerberosPrivacy returns the EnableNfsKerberosPrivacy field if non-nil, zero value otherwise.

### GetEnableNfsKerberosPrivacyOk

`func (o *TemplateViewParams) GetEnableNfsKerberosPrivacyOk() (*bool, bool)`

GetEnableNfsKerberosPrivacyOk returns a tuple with the EnableNfsKerberosPrivacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNfsKerberosPrivacy

`func (o *TemplateViewParams) SetEnableNfsKerberosPrivacy(v bool)`

SetEnableNfsKerberosPrivacy sets EnableNfsKerberosPrivacy field to given value.

### HasEnableNfsKerberosPrivacy

`func (o *TemplateViewParams) HasEnableNfsKerberosPrivacy() bool`

HasEnableNfsKerberosPrivacy returns a boolean if a field has been set.

### SetEnableNfsKerberosPrivacyNil

`func (o *TemplateViewParams) SetEnableNfsKerberosPrivacyNil(b bool)`

 SetEnableNfsKerberosPrivacyNil sets the value for EnableNfsKerberosPrivacy to be an explicit nil

### UnsetEnableNfsKerberosPrivacy
`func (o *TemplateViewParams) UnsetEnableNfsKerberosPrivacy()`

UnsetEnableNfsKerberosPrivacy ensures that no value is present for EnableNfsKerberosPrivacy, not even an explicit nil
### GetEnableNfsUnixAuthentication

`func (o *TemplateViewParams) GetEnableNfsUnixAuthentication() bool`

GetEnableNfsUnixAuthentication returns the EnableNfsUnixAuthentication field if non-nil, zero value otherwise.

### GetEnableNfsUnixAuthenticationOk

`func (o *TemplateViewParams) GetEnableNfsUnixAuthenticationOk() (*bool, bool)`

GetEnableNfsUnixAuthenticationOk returns a tuple with the EnableNfsUnixAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNfsUnixAuthentication

`func (o *TemplateViewParams) SetEnableNfsUnixAuthentication(v bool)`

SetEnableNfsUnixAuthentication sets EnableNfsUnixAuthentication field to given value.

### HasEnableNfsUnixAuthentication

`func (o *TemplateViewParams) HasEnableNfsUnixAuthentication() bool`

HasEnableNfsUnixAuthentication returns a boolean if a field has been set.

### SetEnableNfsUnixAuthenticationNil

`func (o *TemplateViewParams) SetEnableNfsUnixAuthenticationNil(b bool)`

 SetEnableNfsUnixAuthenticationNil sets the value for EnableNfsUnixAuthentication to be an explicit nil

### UnsetEnableNfsUnixAuthentication
`func (o *TemplateViewParams) UnsetEnableNfsUnixAuthentication()`

UnsetEnableNfsUnixAuthentication ensures that no value is present for EnableNfsUnixAuthentication, not even an explicit nil
### GetEnableNfsViewDiscovery

`func (o *TemplateViewParams) GetEnableNfsViewDiscovery() bool`

GetEnableNfsViewDiscovery returns the EnableNfsViewDiscovery field if non-nil, zero value otherwise.

### GetEnableNfsViewDiscoveryOk

`func (o *TemplateViewParams) GetEnableNfsViewDiscoveryOk() (*bool, bool)`

GetEnableNfsViewDiscoveryOk returns a tuple with the EnableNfsViewDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNfsViewDiscovery

`func (o *TemplateViewParams) SetEnableNfsViewDiscovery(v bool)`

SetEnableNfsViewDiscovery sets EnableNfsViewDiscovery field to given value.

### HasEnableNfsViewDiscovery

`func (o *TemplateViewParams) HasEnableNfsViewDiscovery() bool`

HasEnableNfsViewDiscovery returns a boolean if a field has been set.

### SetEnableNfsViewDiscoveryNil

`func (o *TemplateViewParams) SetEnableNfsViewDiscoveryNil(b bool)`

 SetEnableNfsViewDiscoveryNil sets the value for EnableNfsViewDiscovery to be an explicit nil

### UnsetEnableNfsViewDiscovery
`func (o *TemplateViewParams) UnsetEnableNfsViewDiscovery()`

UnsetEnableNfsViewDiscovery ensures that no value is present for EnableNfsViewDiscovery, not even an explicit nil
### GetEnableNfsWcc

`func (o *TemplateViewParams) GetEnableNfsWcc() bool`

GetEnableNfsWcc returns the EnableNfsWcc field if non-nil, zero value otherwise.

### GetEnableNfsWccOk

`func (o *TemplateViewParams) GetEnableNfsWccOk() (*bool, bool)`

GetEnableNfsWccOk returns a tuple with the EnableNfsWcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNfsWcc

`func (o *TemplateViewParams) SetEnableNfsWcc(v bool)`

SetEnableNfsWcc sets EnableNfsWcc field to given value.

### HasEnableNfsWcc

`func (o *TemplateViewParams) HasEnableNfsWcc() bool`

HasEnableNfsWcc returns a boolean if a field has been set.

### SetEnableNfsWccNil

`func (o *TemplateViewParams) SetEnableNfsWccNil(b bool)`

 SetEnableNfsWccNil sets the value for EnableNfsWcc to be an explicit nil

### UnsetEnableNfsWcc
`func (o *TemplateViewParams) UnsetEnableNfsWcc()`

UnsetEnableNfsWcc ensures that no value is present for EnableNfsWcc, not even an explicit nil
### GetNfsAllSquash

`func (o *TemplateViewParams) GetNfsAllSquash() NfsConfigNfsAllSquash`

GetNfsAllSquash returns the NfsAllSquash field if non-nil, zero value otherwise.

### GetNfsAllSquashOk

`func (o *TemplateViewParams) GetNfsAllSquashOk() (*NfsConfigNfsAllSquash, bool)`

GetNfsAllSquashOk returns a tuple with the NfsAllSquash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsAllSquash

`func (o *TemplateViewParams) SetNfsAllSquash(v NfsConfigNfsAllSquash)`

SetNfsAllSquash sets NfsAllSquash field to given value.

### HasNfsAllSquash

`func (o *TemplateViewParams) HasNfsAllSquash() bool`

HasNfsAllSquash returns a boolean if a field has been set.

### GetNfsRootPermissions

`func (o *TemplateViewParams) GetNfsRootPermissions() NfsConfigNfsRootPermissions`

GetNfsRootPermissions returns the NfsRootPermissions field if non-nil, zero value otherwise.

### GetNfsRootPermissionsOk

`func (o *TemplateViewParams) GetNfsRootPermissionsOk() (*NfsConfigNfsRootPermissions, bool)`

GetNfsRootPermissionsOk returns a tuple with the NfsRootPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsRootPermissions

`func (o *TemplateViewParams) SetNfsRootPermissions(v NfsConfigNfsRootPermissions)`

SetNfsRootPermissions sets NfsRootPermissions field to given value.

### HasNfsRootPermissions

`func (o *TemplateViewParams) HasNfsRootPermissions() bool`

HasNfsRootPermissions returns a boolean if a field has been set.

### GetNfsRootSquash

`func (o *TemplateViewParams) GetNfsRootSquash() NfsConfigNfsRootSquash`

GetNfsRootSquash returns the NfsRootSquash field if non-nil, zero value otherwise.

### GetNfsRootSquashOk

`func (o *TemplateViewParams) GetNfsRootSquashOk() (*NfsConfigNfsRootSquash, bool)`

GetNfsRootSquashOk returns a tuple with the NfsRootSquash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsRootSquash

`func (o *TemplateViewParams) SetNfsRootSquash(v NfsConfigNfsRootSquash)`

SetNfsRootSquash sets NfsRootSquash field to given value.

### HasNfsRootSquash

`func (o *TemplateViewParams) HasNfsRootSquash() bool`

HasNfsRootSquash returns a boolean if a field has been set.

### GetEnableFastDurableHandle

`func (o *TemplateViewParams) GetEnableFastDurableHandle() bool`

GetEnableFastDurableHandle returns the EnableFastDurableHandle field if non-nil, zero value otherwise.

### GetEnableFastDurableHandleOk

`func (o *TemplateViewParams) GetEnableFastDurableHandleOk() (*bool, bool)`

GetEnableFastDurableHandleOk returns a tuple with the EnableFastDurableHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFastDurableHandle

`func (o *TemplateViewParams) SetEnableFastDurableHandle(v bool)`

SetEnableFastDurableHandle sets EnableFastDurableHandle field to given value.

### HasEnableFastDurableHandle

`func (o *TemplateViewParams) HasEnableFastDurableHandle() bool`

HasEnableFastDurableHandle returns a boolean if a field has been set.

### SetEnableFastDurableHandleNil

`func (o *TemplateViewParams) SetEnableFastDurableHandleNil(b bool)`

 SetEnableFastDurableHandleNil sets the value for EnableFastDurableHandle to be an explicit nil

### UnsetEnableFastDurableHandle
`func (o *TemplateViewParams) UnsetEnableFastDurableHandle()`

UnsetEnableFastDurableHandle ensures that no value is present for EnableFastDurableHandle, not even an explicit nil
### GetEnableSmbAccessBasedEnumeration

`func (o *TemplateViewParams) GetEnableSmbAccessBasedEnumeration() bool`

GetEnableSmbAccessBasedEnumeration returns the EnableSmbAccessBasedEnumeration field if non-nil, zero value otherwise.

### GetEnableSmbAccessBasedEnumerationOk

`func (o *TemplateViewParams) GetEnableSmbAccessBasedEnumerationOk() (*bool, bool)`

GetEnableSmbAccessBasedEnumerationOk returns a tuple with the EnableSmbAccessBasedEnumeration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSmbAccessBasedEnumeration

`func (o *TemplateViewParams) SetEnableSmbAccessBasedEnumeration(v bool)`

SetEnableSmbAccessBasedEnumeration sets EnableSmbAccessBasedEnumeration field to given value.

### HasEnableSmbAccessBasedEnumeration

`func (o *TemplateViewParams) HasEnableSmbAccessBasedEnumeration() bool`

HasEnableSmbAccessBasedEnumeration returns a boolean if a field has been set.

### SetEnableSmbAccessBasedEnumerationNil

`func (o *TemplateViewParams) SetEnableSmbAccessBasedEnumerationNil(b bool)`

 SetEnableSmbAccessBasedEnumerationNil sets the value for EnableSmbAccessBasedEnumeration to be an explicit nil

### UnsetEnableSmbAccessBasedEnumeration
`func (o *TemplateViewParams) UnsetEnableSmbAccessBasedEnumeration()`

UnsetEnableSmbAccessBasedEnumeration ensures that no value is present for EnableSmbAccessBasedEnumeration, not even an explicit nil
### GetEnableSmbEncryption

`func (o *TemplateViewParams) GetEnableSmbEncryption() bool`

GetEnableSmbEncryption returns the EnableSmbEncryption field if non-nil, zero value otherwise.

### GetEnableSmbEncryptionOk

`func (o *TemplateViewParams) GetEnableSmbEncryptionOk() (*bool, bool)`

GetEnableSmbEncryptionOk returns a tuple with the EnableSmbEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSmbEncryption

`func (o *TemplateViewParams) SetEnableSmbEncryption(v bool)`

SetEnableSmbEncryption sets EnableSmbEncryption field to given value.

### HasEnableSmbEncryption

`func (o *TemplateViewParams) HasEnableSmbEncryption() bool`

HasEnableSmbEncryption returns a boolean if a field has been set.

### SetEnableSmbEncryptionNil

`func (o *TemplateViewParams) SetEnableSmbEncryptionNil(b bool)`

 SetEnableSmbEncryptionNil sets the value for EnableSmbEncryption to be an explicit nil

### UnsetEnableSmbEncryption
`func (o *TemplateViewParams) UnsetEnableSmbEncryption()`

UnsetEnableSmbEncryption ensures that no value is present for EnableSmbEncryption, not even an explicit nil
### GetEnableSmbOplock

`func (o *TemplateViewParams) GetEnableSmbOplock() bool`

GetEnableSmbOplock returns the EnableSmbOplock field if non-nil, zero value otherwise.

### GetEnableSmbOplockOk

`func (o *TemplateViewParams) GetEnableSmbOplockOk() (*bool, bool)`

GetEnableSmbOplockOk returns a tuple with the EnableSmbOplock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSmbOplock

`func (o *TemplateViewParams) SetEnableSmbOplock(v bool)`

SetEnableSmbOplock sets EnableSmbOplock field to given value.

### HasEnableSmbOplock

`func (o *TemplateViewParams) HasEnableSmbOplock() bool`

HasEnableSmbOplock returns a boolean if a field has been set.

### SetEnableSmbOplockNil

`func (o *TemplateViewParams) SetEnableSmbOplockNil(b bool)`

 SetEnableSmbOplockNil sets the value for EnableSmbOplock to be an explicit nil

### UnsetEnableSmbOplock
`func (o *TemplateViewParams) UnsetEnableSmbOplock()`

UnsetEnableSmbOplock ensures that no value is present for EnableSmbOplock, not even an explicit nil
### GetEnableSmbViewDiscovery

`func (o *TemplateViewParams) GetEnableSmbViewDiscovery() bool`

GetEnableSmbViewDiscovery returns the EnableSmbViewDiscovery field if non-nil, zero value otherwise.

### GetEnableSmbViewDiscoveryOk

`func (o *TemplateViewParams) GetEnableSmbViewDiscoveryOk() (*bool, bool)`

GetEnableSmbViewDiscoveryOk returns a tuple with the EnableSmbViewDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSmbViewDiscovery

`func (o *TemplateViewParams) SetEnableSmbViewDiscovery(v bool)`

SetEnableSmbViewDiscovery sets EnableSmbViewDiscovery field to given value.

### HasEnableSmbViewDiscovery

`func (o *TemplateViewParams) HasEnableSmbViewDiscovery() bool`

HasEnableSmbViewDiscovery returns a boolean if a field has been set.

### SetEnableSmbViewDiscoveryNil

`func (o *TemplateViewParams) SetEnableSmbViewDiscoveryNil(b bool)`

 SetEnableSmbViewDiscoveryNil sets the value for EnableSmbViewDiscovery to be an explicit nil

### UnsetEnableSmbViewDiscovery
`func (o *TemplateViewParams) UnsetEnableSmbViewDiscovery()`

UnsetEnableSmbViewDiscovery ensures that no value is present for EnableSmbViewDiscovery, not even an explicit nil
### GetEnforceSmbEncryption

`func (o *TemplateViewParams) GetEnforceSmbEncryption() bool`

GetEnforceSmbEncryption returns the EnforceSmbEncryption field if non-nil, zero value otherwise.

### GetEnforceSmbEncryptionOk

`func (o *TemplateViewParams) GetEnforceSmbEncryptionOk() (*bool, bool)`

GetEnforceSmbEncryptionOk returns a tuple with the EnforceSmbEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceSmbEncryption

`func (o *TemplateViewParams) SetEnforceSmbEncryption(v bool)`

SetEnforceSmbEncryption sets EnforceSmbEncryption field to given value.

### HasEnforceSmbEncryption

`func (o *TemplateViewParams) HasEnforceSmbEncryption() bool`

HasEnforceSmbEncryption returns a boolean if a field has been set.

### SetEnforceSmbEncryptionNil

`func (o *TemplateViewParams) SetEnforceSmbEncryptionNil(b bool)`

 SetEnforceSmbEncryptionNil sets the value for EnforceSmbEncryption to be an explicit nil

### UnsetEnforceSmbEncryption
`func (o *TemplateViewParams) UnsetEnforceSmbEncryption()`

UnsetEnforceSmbEncryption ensures that no value is present for EnforceSmbEncryption, not even an explicit nil
### GetSharePermissions

`func (o *TemplateViewParams) GetSharePermissions() SmbConfigSharePermissions`

GetSharePermissions returns the SharePermissions field if non-nil, zero value otherwise.

### GetSharePermissionsOk

`func (o *TemplateViewParams) GetSharePermissionsOk() (*SmbConfigSharePermissions, bool)`

GetSharePermissionsOk returns a tuple with the SharePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharePermissions

`func (o *TemplateViewParams) SetSharePermissions(v SmbConfigSharePermissions)`

SetSharePermissions sets SharePermissions field to given value.

### HasSharePermissions

`func (o *TemplateViewParams) HasSharePermissions() bool`

HasSharePermissions returns a boolean if a field has been set.

### GetSmbPermissionsInfo

`func (o *TemplateViewParams) GetSmbPermissionsInfo() SmbConfigSmbPermissionsInfo`

GetSmbPermissionsInfo returns the SmbPermissionsInfo field if non-nil, zero value otherwise.

### GetSmbPermissionsInfoOk

`func (o *TemplateViewParams) GetSmbPermissionsInfoOk() (*SmbConfigSmbPermissionsInfo, bool)`

GetSmbPermissionsInfoOk returns a tuple with the SmbPermissionsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbPermissionsInfo

`func (o *TemplateViewParams) SetSmbPermissionsInfo(v SmbConfigSmbPermissionsInfo)`

SetSmbPermissionsInfo sets SmbPermissionsInfo field to given value.

### HasSmbPermissionsInfo

`func (o *TemplateViewParams) HasSmbPermissionsInfo() bool`

HasSmbPermissionsInfo returns a boolean if a field has been set.

### GetAclConfig

`func (o *TemplateViewParams) GetAclConfig() S3ConfigAclConfig`

GetAclConfig returns the AclConfig field if non-nil, zero value otherwise.

### GetAclConfigOk

`func (o *TemplateViewParams) GetAclConfigOk() (*S3ConfigAclConfig, bool)`

GetAclConfigOk returns a tuple with the AclConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclConfig

`func (o *TemplateViewParams) SetAclConfig(v S3ConfigAclConfig)`

SetAclConfig sets AclConfig field to given value.

### HasAclConfig

`func (o *TemplateViewParams) HasAclConfig() bool`

HasAclConfig returns a boolean if a field has been set.

### GetBucketPolicy

`func (o *TemplateViewParams) GetBucketPolicy() S3ConfigBucketPolicy`

GetBucketPolicy returns the BucketPolicy field if non-nil, zero value otherwise.

### GetBucketPolicyOk

`func (o *TemplateViewParams) GetBucketPolicyOk() (*S3ConfigBucketPolicy, bool)`

GetBucketPolicyOk returns a tuple with the BucketPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketPolicy

`func (o *TemplateViewParams) SetBucketPolicy(v S3ConfigBucketPolicy)`

SetBucketPolicy sets BucketPolicy field to given value.

### HasBucketPolicy

`func (o *TemplateViewParams) HasBucketPolicy() bool`

HasBucketPolicy returns a boolean if a field has been set.

### GetEnableAbac

`func (o *TemplateViewParams) GetEnableAbac() bool`

GetEnableAbac returns the EnableAbac field if non-nil, zero value otherwise.

### GetEnableAbacOk

`func (o *TemplateViewParams) GetEnableAbacOk() (*bool, bool)`

GetEnableAbacOk returns a tuple with the EnableAbac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAbac

`func (o *TemplateViewParams) SetEnableAbac(v bool)`

SetEnableAbac sets EnableAbac field to given value.

### HasEnableAbac

`func (o *TemplateViewParams) HasEnableAbac() bool`

HasEnableAbac returns a boolean if a field has been set.

### SetEnableAbacNil

`func (o *TemplateViewParams) SetEnableAbacNil(b bool)`

 SetEnableAbacNil sets the value for EnableAbac to be an explicit nil

### UnsetEnableAbac
`func (o *TemplateViewParams) UnsetEnableAbac()`

UnsetEnableAbac ensures that no value is present for EnableAbac, not even an explicit nil
### GetLifecycleManagement

`func (o *TemplateViewParams) GetLifecycleManagement() S3ConfigLifecycleManagement`

GetLifecycleManagement returns the LifecycleManagement field if non-nil, zero value otherwise.

### GetLifecycleManagementOk

`func (o *TemplateViewParams) GetLifecycleManagementOk() (*S3ConfigLifecycleManagement, bool)`

GetLifecycleManagementOk returns a tuple with the LifecycleManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleManagement

`func (o *TemplateViewParams) SetLifecycleManagement(v S3ConfigLifecycleManagement)`

SetLifecycleManagement sets LifecycleManagement field to given value.

### HasLifecycleManagement

`func (o *TemplateViewParams) HasLifecycleManagement() bool`

HasLifecycleManagement returns a boolean if a field has been set.

### GetOwnerInfo

`func (o *TemplateViewParams) GetOwnerInfo() S3ConfigOwnerInfo`

GetOwnerInfo returns the OwnerInfo field if non-nil, zero value otherwise.

### GetOwnerInfoOk

`func (o *TemplateViewParams) GetOwnerInfoOk() (*S3ConfigOwnerInfo, bool)`

GetOwnerInfoOk returns a tuple with the OwnerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerInfo

`func (o *TemplateViewParams) SetOwnerInfo(v S3ConfigOwnerInfo)`

SetOwnerInfo sets OwnerInfo field to given value.

### HasOwnerInfo

`func (o *TemplateViewParams) HasOwnerInfo() bool`

HasOwnerInfo returns a boolean if a field has been set.

### GetS3AccessPath

`func (o *TemplateViewParams) GetS3AccessPath() string`

GetS3AccessPath returns the S3AccessPath field if non-nil, zero value otherwise.

### GetS3AccessPathOk

`func (o *TemplateViewParams) GetS3AccessPathOk() (*string, bool)`

GetS3AccessPathOk returns a tuple with the S3AccessPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3AccessPath

`func (o *TemplateViewParams) SetS3AccessPath(v string)`

SetS3AccessPath sets S3AccessPath field to given value.

### HasS3AccessPath

`func (o *TemplateViewParams) HasS3AccessPath() bool`

HasS3AccessPath returns a boolean if a field has been set.

### SetS3AccessPathNil

`func (o *TemplateViewParams) SetS3AccessPathNil(b bool)`

 SetS3AccessPathNil sets the value for S3AccessPath to be an explicit nil

### UnsetS3AccessPath
`func (o *TemplateViewParams) UnsetS3AccessPath()`

UnsetS3AccessPath ensures that no value is present for S3AccessPath, not even an explicit nil
### GetS3EfficientMpuMaxSubfiles

`func (o *TemplateViewParams) GetS3EfficientMpuMaxSubfiles() int32`

GetS3EfficientMpuMaxSubfiles returns the S3EfficientMpuMaxSubfiles field if non-nil, zero value otherwise.

### GetS3EfficientMpuMaxSubfilesOk

`func (o *TemplateViewParams) GetS3EfficientMpuMaxSubfilesOk() (*int32, bool)`

GetS3EfficientMpuMaxSubfilesOk returns a tuple with the S3EfficientMpuMaxSubfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3EfficientMpuMaxSubfiles

`func (o *TemplateViewParams) SetS3EfficientMpuMaxSubfiles(v int32)`

SetS3EfficientMpuMaxSubfiles sets S3EfficientMpuMaxSubfiles field to given value.

### HasS3EfficientMpuMaxSubfiles

`func (o *TemplateViewParams) HasS3EfficientMpuMaxSubfiles() bool`

HasS3EfficientMpuMaxSubfiles returns a boolean if a field has been set.

### SetS3EfficientMpuMaxSubfilesNil

`func (o *TemplateViewParams) SetS3EfficientMpuMaxSubfilesNil(b bool)`

 SetS3EfficientMpuMaxSubfilesNil sets the value for S3EfficientMpuMaxSubfiles to be an explicit nil

### UnsetS3EfficientMpuMaxSubfiles
`func (o *TemplateViewParams) UnsetS3EfficientMpuMaxSubfiles()`

UnsetS3EfficientMpuMaxSubfiles ensures that no value is present for S3EfficientMpuMaxSubfiles, not even an explicit nil
### GetS3EnableEfficientMpu

`func (o *TemplateViewParams) GetS3EnableEfficientMpu() bool`

GetS3EnableEfficientMpu returns the S3EnableEfficientMpu field if non-nil, zero value otherwise.

### GetS3EnableEfficientMpuOk

`func (o *TemplateViewParams) GetS3EnableEfficientMpuOk() (*bool, bool)`

GetS3EnableEfficientMpuOk returns a tuple with the S3EnableEfficientMpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3EnableEfficientMpu

`func (o *TemplateViewParams) SetS3EnableEfficientMpu(v bool)`

SetS3EnableEfficientMpu sets S3EnableEfficientMpu field to given value.

### HasS3EnableEfficientMpu

`func (o *TemplateViewParams) HasS3EnableEfficientMpu() bool`

HasS3EnableEfficientMpu returns a boolean if a field has been set.

### SetS3EnableEfficientMpuNil

`func (o *TemplateViewParams) SetS3EnableEfficientMpuNil(b bool)`

 SetS3EnableEfficientMpuNil sets the value for S3EnableEfficientMpu to be an explicit nil

### UnsetS3EnableEfficientMpu
`func (o *TemplateViewParams) UnsetS3EnableEfficientMpu()`

UnsetS3EnableEfficientMpu ensures that no value is present for S3EnableEfficientMpu, not even an explicit nil
### GetS3MigrationAction

`func (o *TemplateViewParams) GetS3MigrationAction() string`

GetS3MigrationAction returns the S3MigrationAction field if non-nil, zero value otherwise.

### GetS3MigrationActionOk

`func (o *TemplateViewParams) GetS3MigrationActionOk() (*string, bool)`

GetS3MigrationActionOk returns a tuple with the S3MigrationAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3MigrationAction

`func (o *TemplateViewParams) SetS3MigrationAction(v string)`

SetS3MigrationAction sets S3MigrationAction field to given value.

### HasS3MigrationAction

`func (o *TemplateViewParams) HasS3MigrationAction() bool`

HasS3MigrationAction returns a boolean if a field has been set.

### SetS3MigrationActionNil

`func (o *TemplateViewParams) SetS3MigrationActionNil(b bool)`

 SetS3MigrationActionNil sets the value for S3MigrationAction to be an explicit nil

### UnsetS3MigrationAction
`func (o *TemplateViewParams) UnsetS3MigrationAction()`

UnsetS3MigrationAction ensures that no value is present for S3MigrationAction, not even an explicit nil
### GetS3MigrationState

`func (o *TemplateViewParams) GetS3MigrationState() string`

GetS3MigrationState returns the S3MigrationState field if non-nil, zero value otherwise.

### GetS3MigrationStateOk

`func (o *TemplateViewParams) GetS3MigrationStateOk() (*string, bool)`

GetS3MigrationStateOk returns a tuple with the S3MigrationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3MigrationState

`func (o *TemplateViewParams) SetS3MigrationState(v string)`

SetS3MigrationState sets S3MigrationState field to given value.

### HasS3MigrationState

`func (o *TemplateViewParams) HasS3MigrationState() bool`

HasS3MigrationState returns a boolean if a field has been set.

### SetS3MigrationStateNil

`func (o *TemplateViewParams) SetS3MigrationStateNil(b bool)`

 SetS3MigrationStateNil sets the value for S3MigrationState to be an explicit nil

### UnsetS3MigrationState
`func (o *TemplateViewParams) UnsetS3MigrationState()`

UnsetS3MigrationState ensures that no value is present for S3MigrationState, not even an explicit nil
### GetVersioning

`func (o *TemplateViewParams) GetVersioning() string`

GetVersioning returns the Versioning field if non-nil, zero value otherwise.

### GetVersioningOk

`func (o *TemplateViewParams) GetVersioningOk() (*string, bool)`

GetVersioningOk returns a tuple with the Versioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersioning

`func (o *TemplateViewParams) SetVersioning(v string)`

SetVersioning sets Versioning field to given value.

### HasVersioning

`func (o *TemplateViewParams) HasVersioning() bool`

HasVersioning returns a boolean if a field has been set.

### SetVersioningNil

`func (o *TemplateViewParams) SetVersioningNil(b bool)`

 SetVersioningNil sets the value for Versioning to be an explicit nil

### UnsetVersioning
`func (o *TemplateViewParams) UnsetVersioning()`

UnsetVersioning ensures that no value is present for Versioning, not even an explicit nil
### GetSwiftProjectDomain

`func (o *TemplateViewParams) GetSwiftProjectDomain() string`

GetSwiftProjectDomain returns the SwiftProjectDomain field if non-nil, zero value otherwise.

### GetSwiftProjectDomainOk

`func (o *TemplateViewParams) GetSwiftProjectDomainOk() (*string, bool)`

GetSwiftProjectDomainOk returns a tuple with the SwiftProjectDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftProjectDomain

`func (o *TemplateViewParams) SetSwiftProjectDomain(v string)`

SetSwiftProjectDomain sets SwiftProjectDomain field to given value.

### HasSwiftProjectDomain

`func (o *TemplateViewParams) HasSwiftProjectDomain() bool`

HasSwiftProjectDomain returns a boolean if a field has been set.

### SetSwiftProjectDomainNil

`func (o *TemplateViewParams) SetSwiftProjectDomainNil(b bool)`

 SetSwiftProjectDomainNil sets the value for SwiftProjectDomain to be an explicit nil

### UnsetSwiftProjectDomain
`func (o *TemplateViewParams) UnsetSwiftProjectDomain()`

UnsetSwiftProjectDomain ensures that no value is present for SwiftProjectDomain, not even an explicit nil
### GetSwiftProjectName

`func (o *TemplateViewParams) GetSwiftProjectName() string`

GetSwiftProjectName returns the SwiftProjectName field if non-nil, zero value otherwise.

### GetSwiftProjectNameOk

`func (o *TemplateViewParams) GetSwiftProjectNameOk() (*string, bool)`

GetSwiftProjectNameOk returns a tuple with the SwiftProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftProjectName

`func (o *TemplateViewParams) SetSwiftProjectName(v string)`

SetSwiftProjectName sets SwiftProjectName field to given value.

### HasSwiftProjectName

`func (o *TemplateViewParams) HasSwiftProjectName() bool`

HasSwiftProjectName returns a boolean if a field has been set.

### SetSwiftProjectNameNil

`func (o *TemplateViewParams) SetSwiftProjectNameNil(b bool)`

 SetSwiftProjectNameNil sets the value for SwiftProjectName to be an explicit nil

### UnsetSwiftProjectName
`func (o *TemplateViewParams) UnsetSwiftProjectName()`

UnsetSwiftProjectName ensures that no value is present for SwiftProjectName, not even an explicit nil
### GetSwiftUserDomain

`func (o *TemplateViewParams) GetSwiftUserDomain() string`

GetSwiftUserDomain returns the SwiftUserDomain field if non-nil, zero value otherwise.

### GetSwiftUserDomainOk

`func (o *TemplateViewParams) GetSwiftUserDomainOk() (*string, bool)`

GetSwiftUserDomainOk returns a tuple with the SwiftUserDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftUserDomain

`func (o *TemplateViewParams) SetSwiftUserDomain(v string)`

SetSwiftUserDomain sets SwiftUserDomain field to given value.

### HasSwiftUserDomain

`func (o *TemplateViewParams) HasSwiftUserDomain() bool`

HasSwiftUserDomain returns a boolean if a field has been set.

### SetSwiftUserDomainNil

`func (o *TemplateViewParams) SetSwiftUserDomainNil(b bool)`

 SetSwiftUserDomainNil sets the value for SwiftUserDomain to be an explicit nil

### UnsetSwiftUserDomain
`func (o *TemplateViewParams) UnsetSwiftUserDomain()`

UnsetSwiftUserDomain ensures that no value is present for SwiftUserDomain, not even an explicit nil
### GetSwiftUsername

`func (o *TemplateViewParams) GetSwiftUsername() string`

GetSwiftUsername returns the SwiftUsername field if non-nil, zero value otherwise.

### GetSwiftUsernameOk

`func (o *TemplateViewParams) GetSwiftUsernameOk() (*string, bool)`

GetSwiftUsernameOk returns a tuple with the SwiftUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftUsername

`func (o *TemplateViewParams) SetSwiftUsername(v string)`

SetSwiftUsername sets SwiftUsername field to given value.

### HasSwiftUsername

`func (o *TemplateViewParams) HasSwiftUsername() bool`

HasSwiftUsername returns a boolean if a field has been set.

### SetSwiftUsernameNil

`func (o *TemplateViewParams) SetSwiftUsernameNil(b bool)`

 SetSwiftUsernameNil sets the value for SwiftUsername to be an explicit nil

### UnsetSwiftUsername
`func (o *TemplateViewParams) UnsetSwiftUsername()`

UnsetSwiftUsername ensures that no value is present for SwiftUsername, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


