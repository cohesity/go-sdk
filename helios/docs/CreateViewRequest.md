# CreateViewRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageDomainId** | Pointer to **NullableInt64** | Specifies the id of the Storage Domain (View Box) where the View will be created. | [optional] 
**CaseInsensitiveNamesEnabled** | Pointer to **NullableBool** | Specifies whether to support case insensitive file/folder names. This parameter can only be set during create and cannot be changed. | [optional] 
**ObjectServicesMappingConfig** | Pointer to **NullableString** | Specifies the Object Services key mapping config of the view. This parameter can only be set during create and cannot be changed. Configuration of Object Services key mapping. Specifies the type of Object Services key mapping config. | [optional] 
**ViewProtectionConfig** | Pointer to [**ViewProtectionConfig**](ViewProtectionConfig.md) | Specifies the protection config of the View. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the View. | [optional] 
**Category** | Pointer to **NullableString** | Specifies the category of the View. | [optional] 
**ProtocolAccess** | Pointer to [**[]ViewProtocol**](ViewProtocol.md) | Specifies the supported Protocols for the View. | [optional] 
**Qos** | Pointer to [**QoS**](QoS.md) | Specifies the Quality of Service (QoS) Policy for the View. | [optional] 
**OverrideGlobalSubnetWhitelist** | Pointer to **NullableBool** | Specifies whether view level client subnet whitelist overrides cluster and global setting. | [optional] 
**SubnetWhitelist** | Pointer to [**[]Subnet**](Subnet.md) | Array of Subnets. Specifies a list of Subnets with IP addresses that have permissions to access the View. (Overrides or extends the Subnets specified at the global Cohesity Cluster level.) | [optional] 
**OverrideGlobalNetgroupWhitelist** | Pointer to **NullableBool** | Specifies whether view level client netgroup whitelist overrides cluster and global setting. | [optional] 
**NetgroupWhitelist** | Pointer to [**NisNetgroups**](NisNetgroups.md) | Array of Netgroups. Specifies a list of netgroups with domains that have permissions to access the View. (Overrides or extends the Netgroup specified at the global Cohesity Cluster level.) | [optional] 
**SecurityMode** | Pointer to **NullableString** | Specifies the security mode used for this view. Currently we support the following modes: Native, Unified and NTFS style. &#39;NativeMode&#39; indicates a native security mode. &#39;UnifiedMode&#39; indicates a unified security mode. &#39;NtfsMode&#39; indicates a NTFS style security mode. | [optional] 
**StoragePolicyOverride** | Pointer to [**StoragePolicyOverride**](StoragePolicyOverride.md) | Specifies if inline deduplication and compression settings inherited from the Storage Domain (View Box) should be disabled for this View. | [optional] 
**LogicalQuota** | Pointer to [**QuotaPolicy**](QuotaPolicy.md) | Specifies an optional logical quota limit (in bytes) for the usage allowed on this View. (Logical data is when the data is fully hydrated and expanded.) This limit overrides the limit inherited from the Storage Domain (View Box) (if set). If logicalQuota is nil, the limit is inherited from the Storage Domain (View Box) (if set). A new write is not allowed if the Storage Domain (View Box) will exceed the specified quota. However, it takes time for the Cohesity Cluster to calculate the usage across Nodes, so the limit may be exceeded by a small amount. In addition, if the limit is increased or data is removed, there may be a delay before the Cohesity Cluster allows more data to be written to the View, as the Cluster is calculating the usage across Nodes. | [optional] 
**FileLockConfig** | Pointer to [**FileLevelDataLockConfig**](FileLevelDataLockConfig.md) | Optional config that enables file locking for this view. It cannot be disabled during the edit of a view, if it has been enabled during the creation of the view. Also, it cannot be enabled if it was disabled during the creation of the view. | [optional] 
**FileExtensionFilter** | Pointer to [**FileExtensionFilter**](FileExtensionFilter.md) | Optional filtering criteria that should be satisfied by all the files created in this view. It does not affect existing files. | [optional] 
**AntivirusScanConfig** | Pointer to [**AntivirusScanConfig**](AntivirusScanConfig.md) | Specifies the antivirus scan config settings for this View. | [optional] 
**Description** | Pointer to **NullableString** | Specifies an optional text description about the View. | [optional] 
**AllowMountOnWindows** | Pointer to **NullableBool** | Specifies if this View can be mounted using the NFS protocol on Windows systems. If true, this View can be NFS mounted on Windows systems. | [optional] 
**EnableMinion** | Pointer to **NullableBool** | Specifies if this view should allow minion or not. If true, this will allow minion. | [optional] 
**EnableFilerAuditLogging** | Pointer to **NullableBool** | Specifies if Filer Audit Logging is enabled for this view. | [optional] 
**TenantId** | Pointer to **NullableString** | Optional tenant id who has access to this View. | [optional] 
**EnableLiveIndexing** | Pointer to **NullableBool** | Specifies whether to enable live indexing for the view. | [optional] 
**EnableOfflineCaching** | Pointer to **NullableBool** | Specifies whether to enable offline file caching of the view. | [optional] 
**AccessSids** | Pointer to **[]string** | Array of Security Identifiers (SIDs) Specifies the list of security identifiers (SIDs) for the restricted Principals who have access to this View. | [optional] 
**ViewLockEnabled** | Pointer to **NullableBool** | Specifies whether view lock is enabled. If enabled the view cannot be modified or deleted until unlock. By default it is disabled. | [optional] 
**IsReadOnly** | Pointer to **NullableBool** | Specifies if the view is a read only view. User will no longer be able to write to this view if this is set to true. | [optional] 
**ViewPinningConfig** | Pointer to [**ViewPinningConfig**](ViewPinningConfig.md) | Specifies the pinning config of this view. | [optional] 
**SelfServiceSnapshotConfig** | Pointer to [**SelfServiceSnapshotConfig**](SelfServiceSnapshotConfig.md) | Specifies self service config of this view. | [optional] 
**EnableMetadataAccelerator** | Pointer to **NullableBool** | Specifies if metadata accelerator is enabled for this view. Only supported while creating a view. | [optional] 
**IsExternallyTriggeredBackupTarget** | Pointer to **NullableBool** | Specifies whether the view is for externally triggered backup target. If so, Magneto will ignore the backup schedule for the view protection job of this view. By default it is disabled. | [optional] 
**EnableNfsViewDiscovery** | Pointer to **NullableBool** | If set, it enables discovery of view for NFS. | [optional] 
**NfsAllSquash** | Pointer to [**NfsSquash**](NfsSquash.md) | Specifies the NFS all squash config. | [optional] 
**NfsRootPermissions** | Pointer to [**NfsRootPermissions**](NfsRootPermissions.md) | Specifies the NFS root permission config of the view file system. | [optional] 
**NfsRootSquash** | Pointer to [**NfsSquash**](NfsSquash.md) | Specifies the NFS root squash config. | [optional] 
**EnableNfsUnixAuthentication** | Pointer to **NullableBool** | If set, it enables NFS UNIX Authentication | [optional] 
**EnableNfsKerberosAuthentication** | Pointer to **NullableBool** | If set, it enables NFS Kerberos Authentication | [optional] 
**EnableNfsKerberosIntegrity** | Pointer to **NullableBool** | If set, it enables NFS Kerberos Integrity | [optional] 
**EnableNfsKerberosPrivacy** | Pointer to **NullableBool** | If set, it enables NFS Kerberos Privacy | [optional] 
**EnableSmbViewDiscovery** | Pointer to **NullableBool** | If set, it enables discovery of view for SMB. | [optional] 
**EnableSmbAccessBasedEnumeration** | Pointer to **NullableBool** | Specifies if access-based enumeration should be enabled. If &#39;true&#39;, only files and folders that the user has permissions to access are visible on the SMB share for that user. | [optional] 
**EnableSmbEncryption** | Pointer to **NullableBool** | Specifies the SMB encryption for the View. If set, it enables the SMB encryption for the View. Encryption is supported only by SMB 3.x dialects. Dialects that do not support would still access data in unencrypted format. | [optional] 
**EnforceSmbEncryption** | Pointer to **NullableBool** | Specifies the SMB encryption for all the sessions for the View. If set, encryption is enforced for all the sessions for the View. When enabled all future and existing unencrypted sessions are disallowed. | [optional] 
**EnableFastDurableHandle** | Pointer to **NullableBool** | Specifies whether fast durable handle is enabled. If enabled, view open handle will be kept in memory, which results in a higher performance. But the handles cannot be recovered if node or service crashes. | [optional] 
**EnableSmbOplock** | Pointer to **NullableBool** | Specifies whether SMB opportunistic lock is enabled. | [optional] 
**SmbPermissionsInfo** | Pointer to [**SmbPermissionsInfo**](SmbPermissionsInfo.md) | Specifies the SMB permissions for the View. | [optional] 
**SharePermissions** | Pointer to [**ViewSharePermissions**](ViewSharePermissions.md) | Specifies share level permissions of the view. | [optional] 
**S3AccessPath** | Pointer to **NullableString** | Specifies the path to access this View as an S3 share. | [optional] [readonly] 
**AclConfig** | Pointer to [**AclConfig**](AclConfig.md) | Specifies the ACL config of the View as an S3 bucket. | [optional] 
**OwnerInfo** | Pointer to [**S3OwnerInfo**](S3OwnerInfo.md) | Specifies the owner info of the View as an S3 bucket. | [optional] 
**SwiftProjectDomain** | Pointer to **NullableString** | Specifies the Keystone project domain. | [optional] 
**SwiftProjectName** | Pointer to **NullableString** | Specifies the Keystone project name. | [optional] 
**SwiftUserDomain** | Pointer to **NullableString** | Specifies the Keystone user domain. | [optional] 
**SwiftUsername** | Pointer to **NullableString** | Specifies the Keystone username. | [optional] 

## Methods

### NewCreateViewRequest

`func NewCreateViewRequest() *CreateViewRequest`

NewCreateViewRequest instantiates a new CreateViewRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateViewRequestWithDefaults

`func NewCreateViewRequestWithDefaults() *CreateViewRequest`

NewCreateViewRequestWithDefaults instantiates a new CreateViewRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageDomainId

`func (o *CreateViewRequest) GetStorageDomainId() int64`

GetStorageDomainId returns the StorageDomainId field if non-nil, zero value otherwise.

### GetStorageDomainIdOk

`func (o *CreateViewRequest) GetStorageDomainIdOk() (*int64, bool)`

GetStorageDomainIdOk returns a tuple with the StorageDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainId

`func (o *CreateViewRequest) SetStorageDomainId(v int64)`

SetStorageDomainId sets StorageDomainId field to given value.

### HasStorageDomainId

`func (o *CreateViewRequest) HasStorageDomainId() bool`

HasStorageDomainId returns a boolean if a field has been set.

### SetStorageDomainIdNil

`func (o *CreateViewRequest) SetStorageDomainIdNil(b bool)`

 SetStorageDomainIdNil sets the value for StorageDomainId to be an explicit nil

### UnsetStorageDomainId
`func (o *CreateViewRequest) UnsetStorageDomainId()`

UnsetStorageDomainId ensures that no value is present for StorageDomainId, not even an explicit nil
### GetCaseInsensitiveNamesEnabled

`func (o *CreateViewRequest) GetCaseInsensitiveNamesEnabled() bool`

GetCaseInsensitiveNamesEnabled returns the CaseInsensitiveNamesEnabled field if non-nil, zero value otherwise.

### GetCaseInsensitiveNamesEnabledOk

`func (o *CreateViewRequest) GetCaseInsensitiveNamesEnabledOk() (*bool, bool)`

GetCaseInsensitiveNamesEnabledOk returns a tuple with the CaseInsensitiveNamesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseInsensitiveNamesEnabled

`func (o *CreateViewRequest) SetCaseInsensitiveNamesEnabled(v bool)`

SetCaseInsensitiveNamesEnabled sets CaseInsensitiveNamesEnabled field to given value.

### HasCaseInsensitiveNamesEnabled

`func (o *CreateViewRequest) HasCaseInsensitiveNamesEnabled() bool`

HasCaseInsensitiveNamesEnabled returns a boolean if a field has been set.

### SetCaseInsensitiveNamesEnabledNil

`func (o *CreateViewRequest) SetCaseInsensitiveNamesEnabledNil(b bool)`

 SetCaseInsensitiveNamesEnabledNil sets the value for CaseInsensitiveNamesEnabled to be an explicit nil

### UnsetCaseInsensitiveNamesEnabled
`func (o *CreateViewRequest) UnsetCaseInsensitiveNamesEnabled()`

UnsetCaseInsensitiveNamesEnabled ensures that no value is present for CaseInsensitiveNamesEnabled, not even an explicit nil
### GetObjectServicesMappingConfig

`func (o *CreateViewRequest) GetObjectServicesMappingConfig() string`

GetObjectServicesMappingConfig returns the ObjectServicesMappingConfig field if non-nil, zero value otherwise.

### GetObjectServicesMappingConfigOk

`func (o *CreateViewRequest) GetObjectServicesMappingConfigOk() (*string, bool)`

GetObjectServicesMappingConfigOk returns a tuple with the ObjectServicesMappingConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectServicesMappingConfig

`func (o *CreateViewRequest) SetObjectServicesMappingConfig(v string)`

SetObjectServicesMappingConfig sets ObjectServicesMappingConfig field to given value.

### HasObjectServicesMappingConfig

`func (o *CreateViewRequest) HasObjectServicesMappingConfig() bool`

HasObjectServicesMappingConfig returns a boolean if a field has been set.

### SetObjectServicesMappingConfigNil

`func (o *CreateViewRequest) SetObjectServicesMappingConfigNil(b bool)`

 SetObjectServicesMappingConfigNil sets the value for ObjectServicesMappingConfig to be an explicit nil

### UnsetObjectServicesMappingConfig
`func (o *CreateViewRequest) UnsetObjectServicesMappingConfig()`

UnsetObjectServicesMappingConfig ensures that no value is present for ObjectServicesMappingConfig, not even an explicit nil
### GetViewProtectionConfig

`func (o *CreateViewRequest) GetViewProtectionConfig() ViewProtectionConfig`

GetViewProtectionConfig returns the ViewProtectionConfig field if non-nil, zero value otherwise.

### GetViewProtectionConfigOk

`func (o *CreateViewRequest) GetViewProtectionConfigOk() (*ViewProtectionConfig, bool)`

GetViewProtectionConfigOk returns a tuple with the ViewProtectionConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewProtectionConfig

`func (o *CreateViewRequest) SetViewProtectionConfig(v ViewProtectionConfig)`

SetViewProtectionConfig sets ViewProtectionConfig field to given value.

### HasViewProtectionConfig

`func (o *CreateViewRequest) HasViewProtectionConfig() bool`

HasViewProtectionConfig returns a boolean if a field has been set.

### GetName

`func (o *CreateViewRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateViewRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateViewRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateViewRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateViewRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateViewRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCategory

`func (o *CreateViewRequest) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CreateViewRequest) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CreateViewRequest) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CreateViewRequest) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *CreateViewRequest) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *CreateViewRequest) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetProtocolAccess

`func (o *CreateViewRequest) GetProtocolAccess() []ViewProtocol`

GetProtocolAccess returns the ProtocolAccess field if non-nil, zero value otherwise.

### GetProtocolAccessOk

`func (o *CreateViewRequest) GetProtocolAccessOk() (*[]ViewProtocol, bool)`

GetProtocolAccessOk returns a tuple with the ProtocolAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolAccess

`func (o *CreateViewRequest) SetProtocolAccess(v []ViewProtocol)`

SetProtocolAccess sets ProtocolAccess field to given value.

### HasProtocolAccess

`func (o *CreateViewRequest) HasProtocolAccess() bool`

HasProtocolAccess returns a boolean if a field has been set.

### SetProtocolAccessNil

`func (o *CreateViewRequest) SetProtocolAccessNil(b bool)`

 SetProtocolAccessNil sets the value for ProtocolAccess to be an explicit nil

### UnsetProtocolAccess
`func (o *CreateViewRequest) UnsetProtocolAccess()`

UnsetProtocolAccess ensures that no value is present for ProtocolAccess, not even an explicit nil
### GetQos

`func (o *CreateViewRequest) GetQos() QoS`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *CreateViewRequest) GetQosOk() (*QoS, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *CreateViewRequest) SetQos(v QoS)`

SetQos sets Qos field to given value.

### HasQos

`func (o *CreateViewRequest) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetOverrideGlobalSubnetWhitelist

`func (o *CreateViewRequest) GetOverrideGlobalSubnetWhitelist() bool`

GetOverrideGlobalSubnetWhitelist returns the OverrideGlobalSubnetWhitelist field if non-nil, zero value otherwise.

### GetOverrideGlobalSubnetWhitelistOk

`func (o *CreateViewRequest) GetOverrideGlobalSubnetWhitelistOk() (*bool, bool)`

GetOverrideGlobalSubnetWhitelistOk returns a tuple with the OverrideGlobalSubnetWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideGlobalSubnetWhitelist

`func (o *CreateViewRequest) SetOverrideGlobalSubnetWhitelist(v bool)`

SetOverrideGlobalSubnetWhitelist sets OverrideGlobalSubnetWhitelist field to given value.

### HasOverrideGlobalSubnetWhitelist

`func (o *CreateViewRequest) HasOverrideGlobalSubnetWhitelist() bool`

HasOverrideGlobalSubnetWhitelist returns a boolean if a field has been set.

### SetOverrideGlobalSubnetWhitelistNil

`func (o *CreateViewRequest) SetOverrideGlobalSubnetWhitelistNil(b bool)`

 SetOverrideGlobalSubnetWhitelistNil sets the value for OverrideGlobalSubnetWhitelist to be an explicit nil

### UnsetOverrideGlobalSubnetWhitelist
`func (o *CreateViewRequest) UnsetOverrideGlobalSubnetWhitelist()`

UnsetOverrideGlobalSubnetWhitelist ensures that no value is present for OverrideGlobalSubnetWhitelist, not even an explicit nil
### GetSubnetWhitelist

`func (o *CreateViewRequest) GetSubnetWhitelist() []Subnet`

GetSubnetWhitelist returns the SubnetWhitelist field if non-nil, zero value otherwise.

### GetSubnetWhitelistOk

`func (o *CreateViewRequest) GetSubnetWhitelistOk() (*[]Subnet, bool)`

GetSubnetWhitelistOk returns a tuple with the SubnetWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetWhitelist

`func (o *CreateViewRequest) SetSubnetWhitelist(v []Subnet)`

SetSubnetWhitelist sets SubnetWhitelist field to given value.

### HasSubnetWhitelist

`func (o *CreateViewRequest) HasSubnetWhitelist() bool`

HasSubnetWhitelist returns a boolean if a field has been set.

### SetSubnetWhitelistNil

`func (o *CreateViewRequest) SetSubnetWhitelistNil(b bool)`

 SetSubnetWhitelistNil sets the value for SubnetWhitelist to be an explicit nil

### UnsetSubnetWhitelist
`func (o *CreateViewRequest) UnsetSubnetWhitelist()`

UnsetSubnetWhitelist ensures that no value is present for SubnetWhitelist, not even an explicit nil
### GetOverrideGlobalNetgroupWhitelist

`func (o *CreateViewRequest) GetOverrideGlobalNetgroupWhitelist() bool`

GetOverrideGlobalNetgroupWhitelist returns the OverrideGlobalNetgroupWhitelist field if non-nil, zero value otherwise.

### GetOverrideGlobalNetgroupWhitelistOk

`func (o *CreateViewRequest) GetOverrideGlobalNetgroupWhitelistOk() (*bool, bool)`

GetOverrideGlobalNetgroupWhitelistOk returns a tuple with the OverrideGlobalNetgroupWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideGlobalNetgroupWhitelist

`func (o *CreateViewRequest) SetOverrideGlobalNetgroupWhitelist(v bool)`

SetOverrideGlobalNetgroupWhitelist sets OverrideGlobalNetgroupWhitelist field to given value.

### HasOverrideGlobalNetgroupWhitelist

`func (o *CreateViewRequest) HasOverrideGlobalNetgroupWhitelist() bool`

HasOverrideGlobalNetgroupWhitelist returns a boolean if a field has been set.

### SetOverrideGlobalNetgroupWhitelistNil

`func (o *CreateViewRequest) SetOverrideGlobalNetgroupWhitelistNil(b bool)`

 SetOverrideGlobalNetgroupWhitelistNil sets the value for OverrideGlobalNetgroupWhitelist to be an explicit nil

### UnsetOverrideGlobalNetgroupWhitelist
`func (o *CreateViewRequest) UnsetOverrideGlobalNetgroupWhitelist()`

UnsetOverrideGlobalNetgroupWhitelist ensures that no value is present for OverrideGlobalNetgroupWhitelist, not even an explicit nil
### GetNetgroupWhitelist

`func (o *CreateViewRequest) GetNetgroupWhitelist() NisNetgroups`

GetNetgroupWhitelist returns the NetgroupWhitelist field if non-nil, zero value otherwise.

### GetNetgroupWhitelistOk

`func (o *CreateViewRequest) GetNetgroupWhitelistOk() (*NisNetgroups, bool)`

GetNetgroupWhitelistOk returns a tuple with the NetgroupWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetgroupWhitelist

`func (o *CreateViewRequest) SetNetgroupWhitelist(v NisNetgroups)`

SetNetgroupWhitelist sets NetgroupWhitelist field to given value.

### HasNetgroupWhitelist

`func (o *CreateViewRequest) HasNetgroupWhitelist() bool`

HasNetgroupWhitelist returns a boolean if a field has been set.

### GetSecurityMode

`func (o *CreateViewRequest) GetSecurityMode() string`

GetSecurityMode returns the SecurityMode field if non-nil, zero value otherwise.

### GetSecurityModeOk

`func (o *CreateViewRequest) GetSecurityModeOk() (*string, bool)`

GetSecurityModeOk returns a tuple with the SecurityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityMode

`func (o *CreateViewRequest) SetSecurityMode(v string)`

SetSecurityMode sets SecurityMode field to given value.

### HasSecurityMode

`func (o *CreateViewRequest) HasSecurityMode() bool`

HasSecurityMode returns a boolean if a field has been set.

### SetSecurityModeNil

`func (o *CreateViewRequest) SetSecurityModeNil(b bool)`

 SetSecurityModeNil sets the value for SecurityMode to be an explicit nil

### UnsetSecurityMode
`func (o *CreateViewRequest) UnsetSecurityMode()`

UnsetSecurityMode ensures that no value is present for SecurityMode, not even an explicit nil
### GetStoragePolicyOverride

`func (o *CreateViewRequest) GetStoragePolicyOverride() StoragePolicyOverride`

GetStoragePolicyOverride returns the StoragePolicyOverride field if non-nil, zero value otherwise.

### GetStoragePolicyOverrideOk

`func (o *CreateViewRequest) GetStoragePolicyOverrideOk() (*StoragePolicyOverride, bool)`

GetStoragePolicyOverrideOk returns a tuple with the StoragePolicyOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePolicyOverride

`func (o *CreateViewRequest) SetStoragePolicyOverride(v StoragePolicyOverride)`

SetStoragePolicyOverride sets StoragePolicyOverride field to given value.

### HasStoragePolicyOverride

`func (o *CreateViewRequest) HasStoragePolicyOverride() bool`

HasStoragePolicyOverride returns a boolean if a field has been set.

### GetLogicalQuota

`func (o *CreateViewRequest) GetLogicalQuota() QuotaPolicy`

GetLogicalQuota returns the LogicalQuota field if non-nil, zero value otherwise.

### GetLogicalQuotaOk

`func (o *CreateViewRequest) GetLogicalQuotaOk() (*QuotaPolicy, bool)`

GetLogicalQuotaOk returns a tuple with the LogicalQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalQuota

`func (o *CreateViewRequest) SetLogicalQuota(v QuotaPolicy)`

SetLogicalQuota sets LogicalQuota field to given value.

### HasLogicalQuota

`func (o *CreateViewRequest) HasLogicalQuota() bool`

HasLogicalQuota returns a boolean if a field has been set.

### GetFileLockConfig

`func (o *CreateViewRequest) GetFileLockConfig() FileLevelDataLockConfig`

GetFileLockConfig returns the FileLockConfig field if non-nil, zero value otherwise.

### GetFileLockConfigOk

`func (o *CreateViewRequest) GetFileLockConfigOk() (*FileLevelDataLockConfig, bool)`

GetFileLockConfigOk returns a tuple with the FileLockConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLockConfig

`func (o *CreateViewRequest) SetFileLockConfig(v FileLevelDataLockConfig)`

SetFileLockConfig sets FileLockConfig field to given value.

### HasFileLockConfig

`func (o *CreateViewRequest) HasFileLockConfig() bool`

HasFileLockConfig returns a boolean if a field has been set.

### GetFileExtensionFilter

`func (o *CreateViewRequest) GetFileExtensionFilter() FileExtensionFilter`

GetFileExtensionFilter returns the FileExtensionFilter field if non-nil, zero value otherwise.

### GetFileExtensionFilterOk

`func (o *CreateViewRequest) GetFileExtensionFilterOk() (*FileExtensionFilter, bool)`

GetFileExtensionFilterOk returns a tuple with the FileExtensionFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileExtensionFilter

`func (o *CreateViewRequest) SetFileExtensionFilter(v FileExtensionFilter)`

SetFileExtensionFilter sets FileExtensionFilter field to given value.

### HasFileExtensionFilter

`func (o *CreateViewRequest) HasFileExtensionFilter() bool`

HasFileExtensionFilter returns a boolean if a field has been set.

### GetAntivirusScanConfig

`func (o *CreateViewRequest) GetAntivirusScanConfig() AntivirusScanConfig`

GetAntivirusScanConfig returns the AntivirusScanConfig field if non-nil, zero value otherwise.

### GetAntivirusScanConfigOk

`func (o *CreateViewRequest) GetAntivirusScanConfigOk() (*AntivirusScanConfig, bool)`

GetAntivirusScanConfigOk returns a tuple with the AntivirusScanConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntivirusScanConfig

`func (o *CreateViewRequest) SetAntivirusScanConfig(v AntivirusScanConfig)`

SetAntivirusScanConfig sets AntivirusScanConfig field to given value.

### HasAntivirusScanConfig

`func (o *CreateViewRequest) HasAntivirusScanConfig() bool`

HasAntivirusScanConfig returns a boolean if a field has been set.

### GetDescription

`func (o *CreateViewRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateViewRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateViewRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateViewRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateViewRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateViewRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAllowMountOnWindows

`func (o *CreateViewRequest) GetAllowMountOnWindows() bool`

GetAllowMountOnWindows returns the AllowMountOnWindows field if non-nil, zero value otherwise.

### GetAllowMountOnWindowsOk

`func (o *CreateViewRequest) GetAllowMountOnWindowsOk() (*bool, bool)`

GetAllowMountOnWindowsOk returns a tuple with the AllowMountOnWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMountOnWindows

`func (o *CreateViewRequest) SetAllowMountOnWindows(v bool)`

SetAllowMountOnWindows sets AllowMountOnWindows field to given value.

### HasAllowMountOnWindows

`func (o *CreateViewRequest) HasAllowMountOnWindows() bool`

HasAllowMountOnWindows returns a boolean if a field has been set.

### SetAllowMountOnWindowsNil

`func (o *CreateViewRequest) SetAllowMountOnWindowsNil(b bool)`

 SetAllowMountOnWindowsNil sets the value for AllowMountOnWindows to be an explicit nil

### UnsetAllowMountOnWindows
`func (o *CreateViewRequest) UnsetAllowMountOnWindows()`

UnsetAllowMountOnWindows ensures that no value is present for AllowMountOnWindows, not even an explicit nil
### GetEnableMinion

`func (o *CreateViewRequest) GetEnableMinion() bool`

GetEnableMinion returns the EnableMinion field if non-nil, zero value otherwise.

### GetEnableMinionOk

`func (o *CreateViewRequest) GetEnableMinionOk() (*bool, bool)`

GetEnableMinionOk returns a tuple with the EnableMinion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMinion

`func (o *CreateViewRequest) SetEnableMinion(v bool)`

SetEnableMinion sets EnableMinion field to given value.

### HasEnableMinion

`func (o *CreateViewRequest) HasEnableMinion() bool`

HasEnableMinion returns a boolean if a field has been set.

### SetEnableMinionNil

`func (o *CreateViewRequest) SetEnableMinionNil(b bool)`

 SetEnableMinionNil sets the value for EnableMinion to be an explicit nil

### UnsetEnableMinion
`func (o *CreateViewRequest) UnsetEnableMinion()`

UnsetEnableMinion ensures that no value is present for EnableMinion, not even an explicit nil
### GetEnableFilerAuditLogging

`func (o *CreateViewRequest) GetEnableFilerAuditLogging() bool`

GetEnableFilerAuditLogging returns the EnableFilerAuditLogging field if non-nil, zero value otherwise.

### GetEnableFilerAuditLoggingOk

`func (o *CreateViewRequest) GetEnableFilerAuditLoggingOk() (*bool, bool)`

GetEnableFilerAuditLoggingOk returns a tuple with the EnableFilerAuditLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFilerAuditLogging

`func (o *CreateViewRequest) SetEnableFilerAuditLogging(v bool)`

SetEnableFilerAuditLogging sets EnableFilerAuditLogging field to given value.

### HasEnableFilerAuditLogging

`func (o *CreateViewRequest) HasEnableFilerAuditLogging() bool`

HasEnableFilerAuditLogging returns a boolean if a field has been set.

### SetEnableFilerAuditLoggingNil

`func (o *CreateViewRequest) SetEnableFilerAuditLoggingNil(b bool)`

 SetEnableFilerAuditLoggingNil sets the value for EnableFilerAuditLogging to be an explicit nil

### UnsetEnableFilerAuditLogging
`func (o *CreateViewRequest) UnsetEnableFilerAuditLogging()`

UnsetEnableFilerAuditLogging ensures that no value is present for EnableFilerAuditLogging, not even an explicit nil
### GetTenantId

`func (o *CreateViewRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CreateViewRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CreateViewRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CreateViewRequest) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *CreateViewRequest) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *CreateViewRequest) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnableLiveIndexing

`func (o *CreateViewRequest) GetEnableLiveIndexing() bool`

GetEnableLiveIndexing returns the EnableLiveIndexing field if non-nil, zero value otherwise.

### GetEnableLiveIndexingOk

`func (o *CreateViewRequest) GetEnableLiveIndexingOk() (*bool, bool)`

GetEnableLiveIndexingOk returns a tuple with the EnableLiveIndexing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLiveIndexing

`func (o *CreateViewRequest) SetEnableLiveIndexing(v bool)`

SetEnableLiveIndexing sets EnableLiveIndexing field to given value.

### HasEnableLiveIndexing

`func (o *CreateViewRequest) HasEnableLiveIndexing() bool`

HasEnableLiveIndexing returns a boolean if a field has been set.

### SetEnableLiveIndexingNil

`func (o *CreateViewRequest) SetEnableLiveIndexingNil(b bool)`

 SetEnableLiveIndexingNil sets the value for EnableLiveIndexing to be an explicit nil

### UnsetEnableLiveIndexing
`func (o *CreateViewRequest) UnsetEnableLiveIndexing()`

UnsetEnableLiveIndexing ensures that no value is present for EnableLiveIndexing, not even an explicit nil
### GetEnableOfflineCaching

`func (o *CreateViewRequest) GetEnableOfflineCaching() bool`

GetEnableOfflineCaching returns the EnableOfflineCaching field if non-nil, zero value otherwise.

### GetEnableOfflineCachingOk

`func (o *CreateViewRequest) GetEnableOfflineCachingOk() (*bool, bool)`

GetEnableOfflineCachingOk returns a tuple with the EnableOfflineCaching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableOfflineCaching

`func (o *CreateViewRequest) SetEnableOfflineCaching(v bool)`

SetEnableOfflineCaching sets EnableOfflineCaching field to given value.

### HasEnableOfflineCaching

`func (o *CreateViewRequest) HasEnableOfflineCaching() bool`

HasEnableOfflineCaching returns a boolean if a field has been set.

### SetEnableOfflineCachingNil

`func (o *CreateViewRequest) SetEnableOfflineCachingNil(b bool)`

 SetEnableOfflineCachingNil sets the value for EnableOfflineCaching to be an explicit nil

### UnsetEnableOfflineCaching
`func (o *CreateViewRequest) UnsetEnableOfflineCaching()`

UnsetEnableOfflineCaching ensures that no value is present for EnableOfflineCaching, not even an explicit nil
### GetAccessSids

`func (o *CreateViewRequest) GetAccessSids() []string`

GetAccessSids returns the AccessSids field if non-nil, zero value otherwise.

### GetAccessSidsOk

`func (o *CreateViewRequest) GetAccessSidsOk() (*[]string, bool)`

GetAccessSidsOk returns a tuple with the AccessSids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessSids

`func (o *CreateViewRequest) SetAccessSids(v []string)`

SetAccessSids sets AccessSids field to given value.

### HasAccessSids

`func (o *CreateViewRequest) HasAccessSids() bool`

HasAccessSids returns a boolean if a field has been set.

### SetAccessSidsNil

`func (o *CreateViewRequest) SetAccessSidsNil(b bool)`

 SetAccessSidsNil sets the value for AccessSids to be an explicit nil

### UnsetAccessSids
`func (o *CreateViewRequest) UnsetAccessSids()`

UnsetAccessSids ensures that no value is present for AccessSids, not even an explicit nil
### GetViewLockEnabled

`func (o *CreateViewRequest) GetViewLockEnabled() bool`

GetViewLockEnabled returns the ViewLockEnabled field if non-nil, zero value otherwise.

### GetViewLockEnabledOk

`func (o *CreateViewRequest) GetViewLockEnabledOk() (*bool, bool)`

GetViewLockEnabledOk returns a tuple with the ViewLockEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewLockEnabled

`func (o *CreateViewRequest) SetViewLockEnabled(v bool)`

SetViewLockEnabled sets ViewLockEnabled field to given value.

### HasViewLockEnabled

`func (o *CreateViewRequest) HasViewLockEnabled() bool`

HasViewLockEnabled returns a boolean if a field has been set.

### SetViewLockEnabledNil

`func (o *CreateViewRequest) SetViewLockEnabledNil(b bool)`

 SetViewLockEnabledNil sets the value for ViewLockEnabled to be an explicit nil

### UnsetViewLockEnabled
`func (o *CreateViewRequest) UnsetViewLockEnabled()`

UnsetViewLockEnabled ensures that no value is present for ViewLockEnabled, not even an explicit nil
### GetIsReadOnly

`func (o *CreateViewRequest) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *CreateViewRequest) GetIsReadOnlyOk() (*bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadOnly

`func (o *CreateViewRequest) SetIsReadOnly(v bool)`

SetIsReadOnly sets IsReadOnly field to given value.

### HasIsReadOnly

`func (o *CreateViewRequest) HasIsReadOnly() bool`

HasIsReadOnly returns a boolean if a field has been set.

### SetIsReadOnlyNil

`func (o *CreateViewRequest) SetIsReadOnlyNil(b bool)`

 SetIsReadOnlyNil sets the value for IsReadOnly to be an explicit nil

### UnsetIsReadOnly
`func (o *CreateViewRequest) UnsetIsReadOnly()`

UnsetIsReadOnly ensures that no value is present for IsReadOnly, not even an explicit nil
### GetViewPinningConfig

`func (o *CreateViewRequest) GetViewPinningConfig() ViewPinningConfig`

GetViewPinningConfig returns the ViewPinningConfig field if non-nil, zero value otherwise.

### GetViewPinningConfigOk

`func (o *CreateViewRequest) GetViewPinningConfigOk() (*ViewPinningConfig, bool)`

GetViewPinningConfigOk returns a tuple with the ViewPinningConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewPinningConfig

`func (o *CreateViewRequest) SetViewPinningConfig(v ViewPinningConfig)`

SetViewPinningConfig sets ViewPinningConfig field to given value.

### HasViewPinningConfig

`func (o *CreateViewRequest) HasViewPinningConfig() bool`

HasViewPinningConfig returns a boolean if a field has been set.

### GetSelfServiceSnapshotConfig

`func (o *CreateViewRequest) GetSelfServiceSnapshotConfig() SelfServiceSnapshotConfig`

GetSelfServiceSnapshotConfig returns the SelfServiceSnapshotConfig field if non-nil, zero value otherwise.

### GetSelfServiceSnapshotConfigOk

`func (o *CreateViewRequest) GetSelfServiceSnapshotConfigOk() (*SelfServiceSnapshotConfig, bool)`

GetSelfServiceSnapshotConfigOk returns a tuple with the SelfServiceSnapshotConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfServiceSnapshotConfig

`func (o *CreateViewRequest) SetSelfServiceSnapshotConfig(v SelfServiceSnapshotConfig)`

SetSelfServiceSnapshotConfig sets SelfServiceSnapshotConfig field to given value.

### HasSelfServiceSnapshotConfig

`func (o *CreateViewRequest) HasSelfServiceSnapshotConfig() bool`

HasSelfServiceSnapshotConfig returns a boolean if a field has been set.

### GetEnableMetadataAccelerator

`func (o *CreateViewRequest) GetEnableMetadataAccelerator() bool`

GetEnableMetadataAccelerator returns the EnableMetadataAccelerator field if non-nil, zero value otherwise.

### GetEnableMetadataAcceleratorOk

`func (o *CreateViewRequest) GetEnableMetadataAcceleratorOk() (*bool, bool)`

GetEnableMetadataAcceleratorOk returns a tuple with the EnableMetadataAccelerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMetadataAccelerator

`func (o *CreateViewRequest) SetEnableMetadataAccelerator(v bool)`

SetEnableMetadataAccelerator sets EnableMetadataAccelerator field to given value.

### HasEnableMetadataAccelerator

`func (o *CreateViewRequest) HasEnableMetadataAccelerator() bool`

HasEnableMetadataAccelerator returns a boolean if a field has been set.

### SetEnableMetadataAcceleratorNil

`func (o *CreateViewRequest) SetEnableMetadataAcceleratorNil(b bool)`

 SetEnableMetadataAcceleratorNil sets the value for EnableMetadataAccelerator to be an explicit nil

### UnsetEnableMetadataAccelerator
`func (o *CreateViewRequest) UnsetEnableMetadataAccelerator()`

UnsetEnableMetadataAccelerator ensures that no value is present for EnableMetadataAccelerator, not even an explicit nil
### GetIsExternallyTriggeredBackupTarget

`func (o *CreateViewRequest) GetIsExternallyTriggeredBackupTarget() bool`

GetIsExternallyTriggeredBackupTarget returns the IsExternallyTriggeredBackupTarget field if non-nil, zero value otherwise.

### GetIsExternallyTriggeredBackupTargetOk

`func (o *CreateViewRequest) GetIsExternallyTriggeredBackupTargetOk() (*bool, bool)`

GetIsExternallyTriggeredBackupTargetOk returns a tuple with the IsExternallyTriggeredBackupTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternallyTriggeredBackupTarget

`func (o *CreateViewRequest) SetIsExternallyTriggeredBackupTarget(v bool)`

SetIsExternallyTriggeredBackupTarget sets IsExternallyTriggeredBackupTarget field to given value.

### HasIsExternallyTriggeredBackupTarget

`func (o *CreateViewRequest) HasIsExternallyTriggeredBackupTarget() bool`

HasIsExternallyTriggeredBackupTarget returns a boolean if a field has been set.

### SetIsExternallyTriggeredBackupTargetNil

`func (o *CreateViewRequest) SetIsExternallyTriggeredBackupTargetNil(b bool)`

 SetIsExternallyTriggeredBackupTargetNil sets the value for IsExternallyTriggeredBackupTarget to be an explicit nil

### UnsetIsExternallyTriggeredBackupTarget
`func (o *CreateViewRequest) UnsetIsExternallyTriggeredBackupTarget()`

UnsetIsExternallyTriggeredBackupTarget ensures that no value is present for IsExternallyTriggeredBackupTarget, not even an explicit nil
### GetEnableNfsViewDiscovery

`func (o *CreateViewRequest) GetEnableNfsViewDiscovery() bool`

GetEnableNfsViewDiscovery returns the EnableNfsViewDiscovery field if non-nil, zero value otherwise.

### GetEnableNfsViewDiscoveryOk

`func (o *CreateViewRequest) GetEnableNfsViewDiscoveryOk() (*bool, bool)`

GetEnableNfsViewDiscoveryOk returns a tuple with the EnableNfsViewDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNfsViewDiscovery

`func (o *CreateViewRequest) SetEnableNfsViewDiscovery(v bool)`

SetEnableNfsViewDiscovery sets EnableNfsViewDiscovery field to given value.

### HasEnableNfsViewDiscovery

`func (o *CreateViewRequest) HasEnableNfsViewDiscovery() bool`

HasEnableNfsViewDiscovery returns a boolean if a field has been set.

### SetEnableNfsViewDiscoveryNil

`func (o *CreateViewRequest) SetEnableNfsViewDiscoveryNil(b bool)`

 SetEnableNfsViewDiscoveryNil sets the value for EnableNfsViewDiscovery to be an explicit nil

### UnsetEnableNfsViewDiscovery
`func (o *CreateViewRequest) UnsetEnableNfsViewDiscovery()`

UnsetEnableNfsViewDiscovery ensures that no value is present for EnableNfsViewDiscovery, not even an explicit nil
### GetNfsAllSquash

`func (o *CreateViewRequest) GetNfsAllSquash() NfsSquash`

GetNfsAllSquash returns the NfsAllSquash field if non-nil, zero value otherwise.

### GetNfsAllSquashOk

`func (o *CreateViewRequest) GetNfsAllSquashOk() (*NfsSquash, bool)`

GetNfsAllSquashOk returns a tuple with the NfsAllSquash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsAllSquash

`func (o *CreateViewRequest) SetNfsAllSquash(v NfsSquash)`

SetNfsAllSquash sets NfsAllSquash field to given value.

### HasNfsAllSquash

`func (o *CreateViewRequest) HasNfsAllSquash() bool`

HasNfsAllSquash returns a boolean if a field has been set.

### GetNfsRootPermissions

`func (o *CreateViewRequest) GetNfsRootPermissions() NfsRootPermissions`

GetNfsRootPermissions returns the NfsRootPermissions field if non-nil, zero value otherwise.

### GetNfsRootPermissionsOk

`func (o *CreateViewRequest) GetNfsRootPermissionsOk() (*NfsRootPermissions, bool)`

GetNfsRootPermissionsOk returns a tuple with the NfsRootPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsRootPermissions

`func (o *CreateViewRequest) SetNfsRootPermissions(v NfsRootPermissions)`

SetNfsRootPermissions sets NfsRootPermissions field to given value.

### HasNfsRootPermissions

`func (o *CreateViewRequest) HasNfsRootPermissions() bool`

HasNfsRootPermissions returns a boolean if a field has been set.

### GetNfsRootSquash

`func (o *CreateViewRequest) GetNfsRootSquash() NfsSquash`

GetNfsRootSquash returns the NfsRootSquash field if non-nil, zero value otherwise.

### GetNfsRootSquashOk

`func (o *CreateViewRequest) GetNfsRootSquashOk() (*NfsSquash, bool)`

GetNfsRootSquashOk returns a tuple with the NfsRootSquash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsRootSquash

`func (o *CreateViewRequest) SetNfsRootSquash(v NfsSquash)`

SetNfsRootSquash sets NfsRootSquash field to given value.

### HasNfsRootSquash

`func (o *CreateViewRequest) HasNfsRootSquash() bool`

HasNfsRootSquash returns a boolean if a field has been set.

### GetEnableNfsUnixAuthentication

`func (o *CreateViewRequest) GetEnableNfsUnixAuthentication() bool`

GetEnableNfsUnixAuthentication returns the EnableNfsUnixAuthentication field if non-nil, zero value otherwise.

### GetEnableNfsUnixAuthenticationOk

`func (o *CreateViewRequest) GetEnableNfsUnixAuthenticationOk() (*bool, bool)`

GetEnableNfsUnixAuthenticationOk returns a tuple with the EnableNfsUnixAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNfsUnixAuthentication

`func (o *CreateViewRequest) SetEnableNfsUnixAuthentication(v bool)`

SetEnableNfsUnixAuthentication sets EnableNfsUnixAuthentication field to given value.

### HasEnableNfsUnixAuthentication

`func (o *CreateViewRequest) HasEnableNfsUnixAuthentication() bool`

HasEnableNfsUnixAuthentication returns a boolean if a field has been set.

### SetEnableNfsUnixAuthenticationNil

`func (o *CreateViewRequest) SetEnableNfsUnixAuthenticationNil(b bool)`

 SetEnableNfsUnixAuthenticationNil sets the value for EnableNfsUnixAuthentication to be an explicit nil

### UnsetEnableNfsUnixAuthentication
`func (o *CreateViewRequest) UnsetEnableNfsUnixAuthentication()`

UnsetEnableNfsUnixAuthentication ensures that no value is present for EnableNfsUnixAuthentication, not even an explicit nil
### GetEnableNfsKerberosAuthentication

`func (o *CreateViewRequest) GetEnableNfsKerberosAuthentication() bool`

GetEnableNfsKerberosAuthentication returns the EnableNfsKerberosAuthentication field if non-nil, zero value otherwise.

### GetEnableNfsKerberosAuthenticationOk

`func (o *CreateViewRequest) GetEnableNfsKerberosAuthenticationOk() (*bool, bool)`

GetEnableNfsKerberosAuthenticationOk returns a tuple with the EnableNfsKerberosAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNfsKerberosAuthentication

`func (o *CreateViewRequest) SetEnableNfsKerberosAuthentication(v bool)`

SetEnableNfsKerberosAuthentication sets EnableNfsKerberosAuthentication field to given value.

### HasEnableNfsKerberosAuthentication

`func (o *CreateViewRequest) HasEnableNfsKerberosAuthentication() bool`

HasEnableNfsKerberosAuthentication returns a boolean if a field has been set.

### SetEnableNfsKerberosAuthenticationNil

`func (o *CreateViewRequest) SetEnableNfsKerberosAuthenticationNil(b bool)`

 SetEnableNfsKerberosAuthenticationNil sets the value for EnableNfsKerberosAuthentication to be an explicit nil

### UnsetEnableNfsKerberosAuthentication
`func (o *CreateViewRequest) UnsetEnableNfsKerberosAuthentication()`

UnsetEnableNfsKerberosAuthentication ensures that no value is present for EnableNfsKerberosAuthentication, not even an explicit nil
### GetEnableNfsKerberosIntegrity

`func (o *CreateViewRequest) GetEnableNfsKerberosIntegrity() bool`

GetEnableNfsKerberosIntegrity returns the EnableNfsKerberosIntegrity field if non-nil, zero value otherwise.

### GetEnableNfsKerberosIntegrityOk

`func (o *CreateViewRequest) GetEnableNfsKerberosIntegrityOk() (*bool, bool)`

GetEnableNfsKerberosIntegrityOk returns a tuple with the EnableNfsKerberosIntegrity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNfsKerberosIntegrity

`func (o *CreateViewRequest) SetEnableNfsKerberosIntegrity(v bool)`

SetEnableNfsKerberosIntegrity sets EnableNfsKerberosIntegrity field to given value.

### HasEnableNfsKerberosIntegrity

`func (o *CreateViewRequest) HasEnableNfsKerberosIntegrity() bool`

HasEnableNfsKerberosIntegrity returns a boolean if a field has been set.

### SetEnableNfsKerberosIntegrityNil

`func (o *CreateViewRequest) SetEnableNfsKerberosIntegrityNil(b bool)`

 SetEnableNfsKerberosIntegrityNil sets the value for EnableNfsKerberosIntegrity to be an explicit nil

### UnsetEnableNfsKerberosIntegrity
`func (o *CreateViewRequest) UnsetEnableNfsKerberosIntegrity()`

UnsetEnableNfsKerberosIntegrity ensures that no value is present for EnableNfsKerberosIntegrity, not even an explicit nil
### GetEnableNfsKerberosPrivacy

`func (o *CreateViewRequest) GetEnableNfsKerberosPrivacy() bool`

GetEnableNfsKerberosPrivacy returns the EnableNfsKerberosPrivacy field if non-nil, zero value otherwise.

### GetEnableNfsKerberosPrivacyOk

`func (o *CreateViewRequest) GetEnableNfsKerberosPrivacyOk() (*bool, bool)`

GetEnableNfsKerberosPrivacyOk returns a tuple with the EnableNfsKerberosPrivacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNfsKerberosPrivacy

`func (o *CreateViewRequest) SetEnableNfsKerberosPrivacy(v bool)`

SetEnableNfsKerberosPrivacy sets EnableNfsKerberosPrivacy field to given value.

### HasEnableNfsKerberosPrivacy

`func (o *CreateViewRequest) HasEnableNfsKerberosPrivacy() bool`

HasEnableNfsKerberosPrivacy returns a boolean if a field has been set.

### SetEnableNfsKerberosPrivacyNil

`func (o *CreateViewRequest) SetEnableNfsKerberosPrivacyNil(b bool)`

 SetEnableNfsKerberosPrivacyNil sets the value for EnableNfsKerberosPrivacy to be an explicit nil

### UnsetEnableNfsKerberosPrivacy
`func (o *CreateViewRequest) UnsetEnableNfsKerberosPrivacy()`

UnsetEnableNfsKerberosPrivacy ensures that no value is present for EnableNfsKerberosPrivacy, not even an explicit nil
### GetEnableSmbViewDiscovery

`func (o *CreateViewRequest) GetEnableSmbViewDiscovery() bool`

GetEnableSmbViewDiscovery returns the EnableSmbViewDiscovery field if non-nil, zero value otherwise.

### GetEnableSmbViewDiscoveryOk

`func (o *CreateViewRequest) GetEnableSmbViewDiscoveryOk() (*bool, bool)`

GetEnableSmbViewDiscoveryOk returns a tuple with the EnableSmbViewDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSmbViewDiscovery

`func (o *CreateViewRequest) SetEnableSmbViewDiscovery(v bool)`

SetEnableSmbViewDiscovery sets EnableSmbViewDiscovery field to given value.

### HasEnableSmbViewDiscovery

`func (o *CreateViewRequest) HasEnableSmbViewDiscovery() bool`

HasEnableSmbViewDiscovery returns a boolean if a field has been set.

### SetEnableSmbViewDiscoveryNil

`func (o *CreateViewRequest) SetEnableSmbViewDiscoveryNil(b bool)`

 SetEnableSmbViewDiscoveryNil sets the value for EnableSmbViewDiscovery to be an explicit nil

### UnsetEnableSmbViewDiscovery
`func (o *CreateViewRequest) UnsetEnableSmbViewDiscovery()`

UnsetEnableSmbViewDiscovery ensures that no value is present for EnableSmbViewDiscovery, not even an explicit nil
### GetEnableSmbAccessBasedEnumeration

`func (o *CreateViewRequest) GetEnableSmbAccessBasedEnumeration() bool`

GetEnableSmbAccessBasedEnumeration returns the EnableSmbAccessBasedEnumeration field if non-nil, zero value otherwise.

### GetEnableSmbAccessBasedEnumerationOk

`func (o *CreateViewRequest) GetEnableSmbAccessBasedEnumerationOk() (*bool, bool)`

GetEnableSmbAccessBasedEnumerationOk returns a tuple with the EnableSmbAccessBasedEnumeration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSmbAccessBasedEnumeration

`func (o *CreateViewRequest) SetEnableSmbAccessBasedEnumeration(v bool)`

SetEnableSmbAccessBasedEnumeration sets EnableSmbAccessBasedEnumeration field to given value.

### HasEnableSmbAccessBasedEnumeration

`func (o *CreateViewRequest) HasEnableSmbAccessBasedEnumeration() bool`

HasEnableSmbAccessBasedEnumeration returns a boolean if a field has been set.

### SetEnableSmbAccessBasedEnumerationNil

`func (o *CreateViewRequest) SetEnableSmbAccessBasedEnumerationNil(b bool)`

 SetEnableSmbAccessBasedEnumerationNil sets the value for EnableSmbAccessBasedEnumeration to be an explicit nil

### UnsetEnableSmbAccessBasedEnumeration
`func (o *CreateViewRequest) UnsetEnableSmbAccessBasedEnumeration()`

UnsetEnableSmbAccessBasedEnumeration ensures that no value is present for EnableSmbAccessBasedEnumeration, not even an explicit nil
### GetEnableSmbEncryption

`func (o *CreateViewRequest) GetEnableSmbEncryption() bool`

GetEnableSmbEncryption returns the EnableSmbEncryption field if non-nil, zero value otherwise.

### GetEnableSmbEncryptionOk

`func (o *CreateViewRequest) GetEnableSmbEncryptionOk() (*bool, bool)`

GetEnableSmbEncryptionOk returns a tuple with the EnableSmbEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSmbEncryption

`func (o *CreateViewRequest) SetEnableSmbEncryption(v bool)`

SetEnableSmbEncryption sets EnableSmbEncryption field to given value.

### HasEnableSmbEncryption

`func (o *CreateViewRequest) HasEnableSmbEncryption() bool`

HasEnableSmbEncryption returns a boolean if a field has been set.

### SetEnableSmbEncryptionNil

`func (o *CreateViewRequest) SetEnableSmbEncryptionNil(b bool)`

 SetEnableSmbEncryptionNil sets the value for EnableSmbEncryption to be an explicit nil

### UnsetEnableSmbEncryption
`func (o *CreateViewRequest) UnsetEnableSmbEncryption()`

UnsetEnableSmbEncryption ensures that no value is present for EnableSmbEncryption, not even an explicit nil
### GetEnforceSmbEncryption

`func (o *CreateViewRequest) GetEnforceSmbEncryption() bool`

GetEnforceSmbEncryption returns the EnforceSmbEncryption field if non-nil, zero value otherwise.

### GetEnforceSmbEncryptionOk

`func (o *CreateViewRequest) GetEnforceSmbEncryptionOk() (*bool, bool)`

GetEnforceSmbEncryptionOk returns a tuple with the EnforceSmbEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceSmbEncryption

`func (o *CreateViewRequest) SetEnforceSmbEncryption(v bool)`

SetEnforceSmbEncryption sets EnforceSmbEncryption field to given value.

### HasEnforceSmbEncryption

`func (o *CreateViewRequest) HasEnforceSmbEncryption() bool`

HasEnforceSmbEncryption returns a boolean if a field has been set.

### SetEnforceSmbEncryptionNil

`func (o *CreateViewRequest) SetEnforceSmbEncryptionNil(b bool)`

 SetEnforceSmbEncryptionNil sets the value for EnforceSmbEncryption to be an explicit nil

### UnsetEnforceSmbEncryption
`func (o *CreateViewRequest) UnsetEnforceSmbEncryption()`

UnsetEnforceSmbEncryption ensures that no value is present for EnforceSmbEncryption, not even an explicit nil
### GetEnableFastDurableHandle

`func (o *CreateViewRequest) GetEnableFastDurableHandle() bool`

GetEnableFastDurableHandle returns the EnableFastDurableHandle field if non-nil, zero value otherwise.

### GetEnableFastDurableHandleOk

`func (o *CreateViewRequest) GetEnableFastDurableHandleOk() (*bool, bool)`

GetEnableFastDurableHandleOk returns a tuple with the EnableFastDurableHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFastDurableHandle

`func (o *CreateViewRequest) SetEnableFastDurableHandle(v bool)`

SetEnableFastDurableHandle sets EnableFastDurableHandle field to given value.

### HasEnableFastDurableHandle

`func (o *CreateViewRequest) HasEnableFastDurableHandle() bool`

HasEnableFastDurableHandle returns a boolean if a field has been set.

### SetEnableFastDurableHandleNil

`func (o *CreateViewRequest) SetEnableFastDurableHandleNil(b bool)`

 SetEnableFastDurableHandleNil sets the value for EnableFastDurableHandle to be an explicit nil

### UnsetEnableFastDurableHandle
`func (o *CreateViewRequest) UnsetEnableFastDurableHandle()`

UnsetEnableFastDurableHandle ensures that no value is present for EnableFastDurableHandle, not even an explicit nil
### GetEnableSmbOplock

`func (o *CreateViewRequest) GetEnableSmbOplock() bool`

GetEnableSmbOplock returns the EnableSmbOplock field if non-nil, zero value otherwise.

### GetEnableSmbOplockOk

`func (o *CreateViewRequest) GetEnableSmbOplockOk() (*bool, bool)`

GetEnableSmbOplockOk returns a tuple with the EnableSmbOplock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSmbOplock

`func (o *CreateViewRequest) SetEnableSmbOplock(v bool)`

SetEnableSmbOplock sets EnableSmbOplock field to given value.

### HasEnableSmbOplock

`func (o *CreateViewRequest) HasEnableSmbOplock() bool`

HasEnableSmbOplock returns a boolean if a field has been set.

### SetEnableSmbOplockNil

`func (o *CreateViewRequest) SetEnableSmbOplockNil(b bool)`

 SetEnableSmbOplockNil sets the value for EnableSmbOplock to be an explicit nil

### UnsetEnableSmbOplock
`func (o *CreateViewRequest) UnsetEnableSmbOplock()`

UnsetEnableSmbOplock ensures that no value is present for EnableSmbOplock, not even an explicit nil
### GetSmbPermissionsInfo

`func (o *CreateViewRequest) GetSmbPermissionsInfo() SmbPermissionsInfo`

GetSmbPermissionsInfo returns the SmbPermissionsInfo field if non-nil, zero value otherwise.

### GetSmbPermissionsInfoOk

`func (o *CreateViewRequest) GetSmbPermissionsInfoOk() (*SmbPermissionsInfo, bool)`

GetSmbPermissionsInfoOk returns a tuple with the SmbPermissionsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbPermissionsInfo

`func (o *CreateViewRequest) SetSmbPermissionsInfo(v SmbPermissionsInfo)`

SetSmbPermissionsInfo sets SmbPermissionsInfo field to given value.

### HasSmbPermissionsInfo

`func (o *CreateViewRequest) HasSmbPermissionsInfo() bool`

HasSmbPermissionsInfo returns a boolean if a field has been set.

### GetSharePermissions

`func (o *CreateViewRequest) GetSharePermissions() ViewSharePermissions`

GetSharePermissions returns the SharePermissions field if non-nil, zero value otherwise.

### GetSharePermissionsOk

`func (o *CreateViewRequest) GetSharePermissionsOk() (*ViewSharePermissions, bool)`

GetSharePermissionsOk returns a tuple with the SharePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharePermissions

`func (o *CreateViewRequest) SetSharePermissions(v ViewSharePermissions)`

SetSharePermissions sets SharePermissions field to given value.

### HasSharePermissions

`func (o *CreateViewRequest) HasSharePermissions() bool`

HasSharePermissions returns a boolean if a field has been set.

### GetS3AccessPath

`func (o *CreateViewRequest) GetS3AccessPath() string`

GetS3AccessPath returns the S3AccessPath field if non-nil, zero value otherwise.

### GetS3AccessPathOk

`func (o *CreateViewRequest) GetS3AccessPathOk() (*string, bool)`

GetS3AccessPathOk returns a tuple with the S3AccessPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3AccessPath

`func (o *CreateViewRequest) SetS3AccessPath(v string)`

SetS3AccessPath sets S3AccessPath field to given value.

### HasS3AccessPath

`func (o *CreateViewRequest) HasS3AccessPath() bool`

HasS3AccessPath returns a boolean if a field has been set.

### SetS3AccessPathNil

`func (o *CreateViewRequest) SetS3AccessPathNil(b bool)`

 SetS3AccessPathNil sets the value for S3AccessPath to be an explicit nil

### UnsetS3AccessPath
`func (o *CreateViewRequest) UnsetS3AccessPath()`

UnsetS3AccessPath ensures that no value is present for S3AccessPath, not even an explicit nil
### GetAclConfig

`func (o *CreateViewRequest) GetAclConfig() AclConfig`

GetAclConfig returns the AclConfig field if non-nil, zero value otherwise.

### GetAclConfigOk

`func (o *CreateViewRequest) GetAclConfigOk() (*AclConfig, bool)`

GetAclConfigOk returns a tuple with the AclConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclConfig

`func (o *CreateViewRequest) SetAclConfig(v AclConfig)`

SetAclConfig sets AclConfig field to given value.

### HasAclConfig

`func (o *CreateViewRequest) HasAclConfig() bool`

HasAclConfig returns a boolean if a field has been set.

### GetOwnerInfo

`func (o *CreateViewRequest) GetOwnerInfo() S3OwnerInfo`

GetOwnerInfo returns the OwnerInfo field if non-nil, zero value otherwise.

### GetOwnerInfoOk

`func (o *CreateViewRequest) GetOwnerInfoOk() (*S3OwnerInfo, bool)`

GetOwnerInfoOk returns a tuple with the OwnerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerInfo

`func (o *CreateViewRequest) SetOwnerInfo(v S3OwnerInfo)`

SetOwnerInfo sets OwnerInfo field to given value.

### HasOwnerInfo

`func (o *CreateViewRequest) HasOwnerInfo() bool`

HasOwnerInfo returns a boolean if a field has been set.

### GetSwiftProjectDomain

`func (o *CreateViewRequest) GetSwiftProjectDomain() string`

GetSwiftProjectDomain returns the SwiftProjectDomain field if non-nil, zero value otherwise.

### GetSwiftProjectDomainOk

`func (o *CreateViewRequest) GetSwiftProjectDomainOk() (*string, bool)`

GetSwiftProjectDomainOk returns a tuple with the SwiftProjectDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftProjectDomain

`func (o *CreateViewRequest) SetSwiftProjectDomain(v string)`

SetSwiftProjectDomain sets SwiftProjectDomain field to given value.

### HasSwiftProjectDomain

`func (o *CreateViewRequest) HasSwiftProjectDomain() bool`

HasSwiftProjectDomain returns a boolean if a field has been set.

### SetSwiftProjectDomainNil

`func (o *CreateViewRequest) SetSwiftProjectDomainNil(b bool)`

 SetSwiftProjectDomainNil sets the value for SwiftProjectDomain to be an explicit nil

### UnsetSwiftProjectDomain
`func (o *CreateViewRequest) UnsetSwiftProjectDomain()`

UnsetSwiftProjectDomain ensures that no value is present for SwiftProjectDomain, not even an explicit nil
### GetSwiftProjectName

`func (o *CreateViewRequest) GetSwiftProjectName() string`

GetSwiftProjectName returns the SwiftProjectName field if non-nil, zero value otherwise.

### GetSwiftProjectNameOk

`func (o *CreateViewRequest) GetSwiftProjectNameOk() (*string, bool)`

GetSwiftProjectNameOk returns a tuple with the SwiftProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftProjectName

`func (o *CreateViewRequest) SetSwiftProjectName(v string)`

SetSwiftProjectName sets SwiftProjectName field to given value.

### HasSwiftProjectName

`func (o *CreateViewRequest) HasSwiftProjectName() bool`

HasSwiftProjectName returns a boolean if a field has been set.

### SetSwiftProjectNameNil

`func (o *CreateViewRequest) SetSwiftProjectNameNil(b bool)`

 SetSwiftProjectNameNil sets the value for SwiftProjectName to be an explicit nil

### UnsetSwiftProjectName
`func (o *CreateViewRequest) UnsetSwiftProjectName()`

UnsetSwiftProjectName ensures that no value is present for SwiftProjectName, not even an explicit nil
### GetSwiftUserDomain

`func (o *CreateViewRequest) GetSwiftUserDomain() string`

GetSwiftUserDomain returns the SwiftUserDomain field if non-nil, zero value otherwise.

### GetSwiftUserDomainOk

`func (o *CreateViewRequest) GetSwiftUserDomainOk() (*string, bool)`

GetSwiftUserDomainOk returns a tuple with the SwiftUserDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftUserDomain

`func (o *CreateViewRequest) SetSwiftUserDomain(v string)`

SetSwiftUserDomain sets SwiftUserDomain field to given value.

### HasSwiftUserDomain

`func (o *CreateViewRequest) HasSwiftUserDomain() bool`

HasSwiftUserDomain returns a boolean if a field has been set.

### SetSwiftUserDomainNil

`func (o *CreateViewRequest) SetSwiftUserDomainNil(b bool)`

 SetSwiftUserDomainNil sets the value for SwiftUserDomain to be an explicit nil

### UnsetSwiftUserDomain
`func (o *CreateViewRequest) UnsetSwiftUserDomain()`

UnsetSwiftUserDomain ensures that no value is present for SwiftUserDomain, not even an explicit nil
### GetSwiftUsername

`func (o *CreateViewRequest) GetSwiftUsername() string`

GetSwiftUsername returns the SwiftUsername field if non-nil, zero value otherwise.

### GetSwiftUsernameOk

`func (o *CreateViewRequest) GetSwiftUsernameOk() (*string, bool)`

GetSwiftUsernameOk returns a tuple with the SwiftUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftUsername

`func (o *CreateViewRequest) SetSwiftUsername(v string)`

SetSwiftUsername sets SwiftUsername field to given value.

### HasSwiftUsername

`func (o *CreateViewRequest) HasSwiftUsername() bool`

HasSwiftUsername returns a boolean if a field has been set.

### SetSwiftUsernameNil

`func (o *CreateViewRequest) SetSwiftUsernameNil(b bool)`

 SetSwiftUsernameNil sets the value for SwiftUsername to be an explicit nil

### UnsetSwiftUsername
`func (o *CreateViewRequest) UnsetSwiftUsername()`

UnsetSwiftUsername ensures that no value is present for SwiftUsername, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


