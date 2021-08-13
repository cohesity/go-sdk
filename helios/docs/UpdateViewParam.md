# UpdateViewParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewUpdateViewParam

`func NewUpdateViewParam() *UpdateViewParam`

NewUpdateViewParam instantiates a new UpdateViewParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateViewParamWithDefaults

`func NewUpdateViewParamWithDefaults() *UpdateViewParam`

NewUpdateViewParamWithDefaults instantiates a new UpdateViewParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateViewParam) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateViewParam) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateViewParam) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateViewParam) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateViewParam) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateViewParam) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCategory

`func (o *UpdateViewParam) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *UpdateViewParam) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *UpdateViewParam) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *UpdateViewParam) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *UpdateViewParam) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *UpdateViewParam) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetProtocolAccess

`func (o *UpdateViewParam) GetProtocolAccess() []ViewProtocol`

GetProtocolAccess returns the ProtocolAccess field if non-nil, zero value otherwise.

### GetProtocolAccessOk

`func (o *UpdateViewParam) GetProtocolAccessOk() (*[]ViewProtocol, bool)`

GetProtocolAccessOk returns a tuple with the ProtocolAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolAccess

`func (o *UpdateViewParam) SetProtocolAccess(v []ViewProtocol)`

SetProtocolAccess sets ProtocolAccess field to given value.

### HasProtocolAccess

`func (o *UpdateViewParam) HasProtocolAccess() bool`

HasProtocolAccess returns a boolean if a field has been set.

### SetProtocolAccessNil

`func (o *UpdateViewParam) SetProtocolAccessNil(b bool)`

 SetProtocolAccessNil sets the value for ProtocolAccess to be an explicit nil

### UnsetProtocolAccess
`func (o *UpdateViewParam) UnsetProtocolAccess()`

UnsetProtocolAccess ensures that no value is present for ProtocolAccess, not even an explicit nil
### GetQos

`func (o *UpdateViewParam) GetQos() QoS`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *UpdateViewParam) GetQosOk() (*QoS, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *UpdateViewParam) SetQos(v QoS)`

SetQos sets Qos field to given value.

### HasQos

`func (o *UpdateViewParam) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetOverrideGlobalSubnetWhitelist

`func (o *UpdateViewParam) GetOverrideGlobalSubnetWhitelist() bool`

GetOverrideGlobalSubnetWhitelist returns the OverrideGlobalSubnetWhitelist field if non-nil, zero value otherwise.

### GetOverrideGlobalSubnetWhitelistOk

`func (o *UpdateViewParam) GetOverrideGlobalSubnetWhitelistOk() (*bool, bool)`

GetOverrideGlobalSubnetWhitelistOk returns a tuple with the OverrideGlobalSubnetWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideGlobalSubnetWhitelist

`func (o *UpdateViewParam) SetOverrideGlobalSubnetWhitelist(v bool)`

SetOverrideGlobalSubnetWhitelist sets OverrideGlobalSubnetWhitelist field to given value.

### HasOverrideGlobalSubnetWhitelist

`func (o *UpdateViewParam) HasOverrideGlobalSubnetWhitelist() bool`

HasOverrideGlobalSubnetWhitelist returns a boolean if a field has been set.

### SetOverrideGlobalSubnetWhitelistNil

`func (o *UpdateViewParam) SetOverrideGlobalSubnetWhitelistNil(b bool)`

 SetOverrideGlobalSubnetWhitelistNil sets the value for OverrideGlobalSubnetWhitelist to be an explicit nil

### UnsetOverrideGlobalSubnetWhitelist
`func (o *UpdateViewParam) UnsetOverrideGlobalSubnetWhitelist()`

UnsetOverrideGlobalSubnetWhitelist ensures that no value is present for OverrideGlobalSubnetWhitelist, not even an explicit nil
### GetSubnetWhitelist

`func (o *UpdateViewParam) GetSubnetWhitelist() []Subnet`

GetSubnetWhitelist returns the SubnetWhitelist field if non-nil, zero value otherwise.

### GetSubnetWhitelistOk

`func (o *UpdateViewParam) GetSubnetWhitelistOk() (*[]Subnet, bool)`

GetSubnetWhitelistOk returns a tuple with the SubnetWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetWhitelist

`func (o *UpdateViewParam) SetSubnetWhitelist(v []Subnet)`

SetSubnetWhitelist sets SubnetWhitelist field to given value.

### HasSubnetWhitelist

`func (o *UpdateViewParam) HasSubnetWhitelist() bool`

HasSubnetWhitelist returns a boolean if a field has been set.

### SetSubnetWhitelistNil

`func (o *UpdateViewParam) SetSubnetWhitelistNil(b bool)`

 SetSubnetWhitelistNil sets the value for SubnetWhitelist to be an explicit nil

### UnsetSubnetWhitelist
`func (o *UpdateViewParam) UnsetSubnetWhitelist()`

UnsetSubnetWhitelist ensures that no value is present for SubnetWhitelist, not even an explicit nil
### GetOverrideGlobalNetgroupWhitelist

`func (o *UpdateViewParam) GetOverrideGlobalNetgroupWhitelist() bool`

GetOverrideGlobalNetgroupWhitelist returns the OverrideGlobalNetgroupWhitelist field if non-nil, zero value otherwise.

### GetOverrideGlobalNetgroupWhitelistOk

`func (o *UpdateViewParam) GetOverrideGlobalNetgroupWhitelistOk() (*bool, bool)`

GetOverrideGlobalNetgroupWhitelistOk returns a tuple with the OverrideGlobalNetgroupWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideGlobalNetgroupWhitelist

`func (o *UpdateViewParam) SetOverrideGlobalNetgroupWhitelist(v bool)`

SetOverrideGlobalNetgroupWhitelist sets OverrideGlobalNetgroupWhitelist field to given value.

### HasOverrideGlobalNetgroupWhitelist

`func (o *UpdateViewParam) HasOverrideGlobalNetgroupWhitelist() bool`

HasOverrideGlobalNetgroupWhitelist returns a boolean if a field has been set.

### SetOverrideGlobalNetgroupWhitelistNil

`func (o *UpdateViewParam) SetOverrideGlobalNetgroupWhitelistNil(b bool)`

 SetOverrideGlobalNetgroupWhitelistNil sets the value for OverrideGlobalNetgroupWhitelist to be an explicit nil

### UnsetOverrideGlobalNetgroupWhitelist
`func (o *UpdateViewParam) UnsetOverrideGlobalNetgroupWhitelist()`

UnsetOverrideGlobalNetgroupWhitelist ensures that no value is present for OverrideGlobalNetgroupWhitelist, not even an explicit nil
### GetNetgroupWhitelist

`func (o *UpdateViewParam) GetNetgroupWhitelist() NisNetgroups`

GetNetgroupWhitelist returns the NetgroupWhitelist field if non-nil, zero value otherwise.

### GetNetgroupWhitelistOk

`func (o *UpdateViewParam) GetNetgroupWhitelistOk() (*NisNetgroups, bool)`

GetNetgroupWhitelistOk returns a tuple with the NetgroupWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetgroupWhitelist

`func (o *UpdateViewParam) SetNetgroupWhitelist(v NisNetgroups)`

SetNetgroupWhitelist sets NetgroupWhitelist field to given value.

### HasNetgroupWhitelist

`func (o *UpdateViewParam) HasNetgroupWhitelist() bool`

HasNetgroupWhitelist returns a boolean if a field has been set.

### GetSecurityMode

`func (o *UpdateViewParam) GetSecurityMode() string`

GetSecurityMode returns the SecurityMode field if non-nil, zero value otherwise.

### GetSecurityModeOk

`func (o *UpdateViewParam) GetSecurityModeOk() (*string, bool)`

GetSecurityModeOk returns a tuple with the SecurityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityMode

`func (o *UpdateViewParam) SetSecurityMode(v string)`

SetSecurityMode sets SecurityMode field to given value.

### HasSecurityMode

`func (o *UpdateViewParam) HasSecurityMode() bool`

HasSecurityMode returns a boolean if a field has been set.

### SetSecurityModeNil

`func (o *UpdateViewParam) SetSecurityModeNil(b bool)`

 SetSecurityModeNil sets the value for SecurityMode to be an explicit nil

### UnsetSecurityMode
`func (o *UpdateViewParam) UnsetSecurityMode()`

UnsetSecurityMode ensures that no value is present for SecurityMode, not even an explicit nil
### GetStoragePolicyOverride

`func (o *UpdateViewParam) GetStoragePolicyOverride() StoragePolicyOverride`

GetStoragePolicyOverride returns the StoragePolicyOverride field if non-nil, zero value otherwise.

### GetStoragePolicyOverrideOk

`func (o *UpdateViewParam) GetStoragePolicyOverrideOk() (*StoragePolicyOverride, bool)`

GetStoragePolicyOverrideOk returns a tuple with the StoragePolicyOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePolicyOverride

`func (o *UpdateViewParam) SetStoragePolicyOverride(v StoragePolicyOverride)`

SetStoragePolicyOverride sets StoragePolicyOverride field to given value.

### HasStoragePolicyOverride

`func (o *UpdateViewParam) HasStoragePolicyOverride() bool`

HasStoragePolicyOverride returns a boolean if a field has been set.

### GetLogicalQuota

`func (o *UpdateViewParam) GetLogicalQuota() QuotaPolicy`

GetLogicalQuota returns the LogicalQuota field if non-nil, zero value otherwise.

### GetLogicalQuotaOk

`func (o *UpdateViewParam) GetLogicalQuotaOk() (*QuotaPolicy, bool)`

GetLogicalQuotaOk returns a tuple with the LogicalQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalQuota

`func (o *UpdateViewParam) SetLogicalQuota(v QuotaPolicy)`

SetLogicalQuota sets LogicalQuota field to given value.

### HasLogicalQuota

`func (o *UpdateViewParam) HasLogicalQuota() bool`

HasLogicalQuota returns a boolean if a field has been set.

### GetFileLockConfig

`func (o *UpdateViewParam) GetFileLockConfig() FileLevelDataLockConfig`

GetFileLockConfig returns the FileLockConfig field if non-nil, zero value otherwise.

### GetFileLockConfigOk

`func (o *UpdateViewParam) GetFileLockConfigOk() (*FileLevelDataLockConfig, bool)`

GetFileLockConfigOk returns a tuple with the FileLockConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLockConfig

`func (o *UpdateViewParam) SetFileLockConfig(v FileLevelDataLockConfig)`

SetFileLockConfig sets FileLockConfig field to given value.

### HasFileLockConfig

`func (o *UpdateViewParam) HasFileLockConfig() bool`

HasFileLockConfig returns a boolean if a field has been set.

### GetFileExtensionFilter

`func (o *UpdateViewParam) GetFileExtensionFilter() FileExtensionFilter`

GetFileExtensionFilter returns the FileExtensionFilter field if non-nil, zero value otherwise.

### GetFileExtensionFilterOk

`func (o *UpdateViewParam) GetFileExtensionFilterOk() (*FileExtensionFilter, bool)`

GetFileExtensionFilterOk returns a tuple with the FileExtensionFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileExtensionFilter

`func (o *UpdateViewParam) SetFileExtensionFilter(v FileExtensionFilter)`

SetFileExtensionFilter sets FileExtensionFilter field to given value.

### HasFileExtensionFilter

`func (o *UpdateViewParam) HasFileExtensionFilter() bool`

HasFileExtensionFilter returns a boolean if a field has been set.

### GetAntivirusScanConfig

`func (o *UpdateViewParam) GetAntivirusScanConfig() AntivirusScanConfig`

GetAntivirusScanConfig returns the AntivirusScanConfig field if non-nil, zero value otherwise.

### GetAntivirusScanConfigOk

`func (o *UpdateViewParam) GetAntivirusScanConfigOk() (*AntivirusScanConfig, bool)`

GetAntivirusScanConfigOk returns a tuple with the AntivirusScanConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntivirusScanConfig

`func (o *UpdateViewParam) SetAntivirusScanConfig(v AntivirusScanConfig)`

SetAntivirusScanConfig sets AntivirusScanConfig field to given value.

### HasAntivirusScanConfig

`func (o *UpdateViewParam) HasAntivirusScanConfig() bool`

HasAntivirusScanConfig returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateViewParam) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateViewParam) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateViewParam) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateViewParam) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateViewParam) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateViewParam) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAllowMountOnWindows

`func (o *UpdateViewParam) GetAllowMountOnWindows() bool`

GetAllowMountOnWindows returns the AllowMountOnWindows field if non-nil, zero value otherwise.

### GetAllowMountOnWindowsOk

`func (o *UpdateViewParam) GetAllowMountOnWindowsOk() (*bool, bool)`

GetAllowMountOnWindowsOk returns a tuple with the AllowMountOnWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMountOnWindows

`func (o *UpdateViewParam) SetAllowMountOnWindows(v bool)`

SetAllowMountOnWindows sets AllowMountOnWindows field to given value.

### HasAllowMountOnWindows

`func (o *UpdateViewParam) HasAllowMountOnWindows() bool`

HasAllowMountOnWindows returns a boolean if a field has been set.

### SetAllowMountOnWindowsNil

`func (o *UpdateViewParam) SetAllowMountOnWindowsNil(b bool)`

 SetAllowMountOnWindowsNil sets the value for AllowMountOnWindows to be an explicit nil

### UnsetAllowMountOnWindows
`func (o *UpdateViewParam) UnsetAllowMountOnWindows()`

UnsetAllowMountOnWindows ensures that no value is present for AllowMountOnWindows, not even an explicit nil
### GetEnableMinion

`func (o *UpdateViewParam) GetEnableMinion() bool`

GetEnableMinion returns the EnableMinion field if non-nil, zero value otherwise.

### GetEnableMinionOk

`func (o *UpdateViewParam) GetEnableMinionOk() (*bool, bool)`

GetEnableMinionOk returns a tuple with the EnableMinion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMinion

`func (o *UpdateViewParam) SetEnableMinion(v bool)`

SetEnableMinion sets EnableMinion field to given value.

### HasEnableMinion

`func (o *UpdateViewParam) HasEnableMinion() bool`

HasEnableMinion returns a boolean if a field has been set.

### SetEnableMinionNil

`func (o *UpdateViewParam) SetEnableMinionNil(b bool)`

 SetEnableMinionNil sets the value for EnableMinion to be an explicit nil

### UnsetEnableMinion
`func (o *UpdateViewParam) UnsetEnableMinion()`

UnsetEnableMinion ensures that no value is present for EnableMinion, not even an explicit nil
### GetEnableFilerAuditLogging

`func (o *UpdateViewParam) GetEnableFilerAuditLogging() bool`

GetEnableFilerAuditLogging returns the EnableFilerAuditLogging field if non-nil, zero value otherwise.

### GetEnableFilerAuditLoggingOk

`func (o *UpdateViewParam) GetEnableFilerAuditLoggingOk() (*bool, bool)`

GetEnableFilerAuditLoggingOk returns a tuple with the EnableFilerAuditLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFilerAuditLogging

`func (o *UpdateViewParam) SetEnableFilerAuditLogging(v bool)`

SetEnableFilerAuditLogging sets EnableFilerAuditLogging field to given value.

### HasEnableFilerAuditLogging

`func (o *UpdateViewParam) HasEnableFilerAuditLogging() bool`

HasEnableFilerAuditLogging returns a boolean if a field has been set.

### SetEnableFilerAuditLoggingNil

`func (o *UpdateViewParam) SetEnableFilerAuditLoggingNil(b bool)`

 SetEnableFilerAuditLoggingNil sets the value for EnableFilerAuditLogging to be an explicit nil

### UnsetEnableFilerAuditLogging
`func (o *UpdateViewParam) UnsetEnableFilerAuditLogging()`

UnsetEnableFilerAuditLogging ensures that no value is present for EnableFilerAuditLogging, not even an explicit nil
### GetTenantId

`func (o *UpdateViewParam) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *UpdateViewParam) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *UpdateViewParam) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *UpdateViewParam) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *UpdateViewParam) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *UpdateViewParam) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnableLiveIndexing

`func (o *UpdateViewParam) GetEnableLiveIndexing() bool`

GetEnableLiveIndexing returns the EnableLiveIndexing field if non-nil, zero value otherwise.

### GetEnableLiveIndexingOk

`func (o *UpdateViewParam) GetEnableLiveIndexingOk() (*bool, bool)`

GetEnableLiveIndexingOk returns a tuple with the EnableLiveIndexing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLiveIndexing

`func (o *UpdateViewParam) SetEnableLiveIndexing(v bool)`

SetEnableLiveIndexing sets EnableLiveIndexing field to given value.

### HasEnableLiveIndexing

`func (o *UpdateViewParam) HasEnableLiveIndexing() bool`

HasEnableLiveIndexing returns a boolean if a field has been set.

### SetEnableLiveIndexingNil

`func (o *UpdateViewParam) SetEnableLiveIndexingNil(b bool)`

 SetEnableLiveIndexingNil sets the value for EnableLiveIndexing to be an explicit nil

### UnsetEnableLiveIndexing
`func (o *UpdateViewParam) UnsetEnableLiveIndexing()`

UnsetEnableLiveIndexing ensures that no value is present for EnableLiveIndexing, not even an explicit nil
### GetEnableOfflineCaching

`func (o *UpdateViewParam) GetEnableOfflineCaching() bool`

GetEnableOfflineCaching returns the EnableOfflineCaching field if non-nil, zero value otherwise.

### GetEnableOfflineCachingOk

`func (o *UpdateViewParam) GetEnableOfflineCachingOk() (*bool, bool)`

GetEnableOfflineCachingOk returns a tuple with the EnableOfflineCaching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableOfflineCaching

`func (o *UpdateViewParam) SetEnableOfflineCaching(v bool)`

SetEnableOfflineCaching sets EnableOfflineCaching field to given value.

### HasEnableOfflineCaching

`func (o *UpdateViewParam) HasEnableOfflineCaching() bool`

HasEnableOfflineCaching returns a boolean if a field has been set.

### SetEnableOfflineCachingNil

`func (o *UpdateViewParam) SetEnableOfflineCachingNil(b bool)`

 SetEnableOfflineCachingNil sets the value for EnableOfflineCaching to be an explicit nil

### UnsetEnableOfflineCaching
`func (o *UpdateViewParam) UnsetEnableOfflineCaching()`

UnsetEnableOfflineCaching ensures that no value is present for EnableOfflineCaching, not even an explicit nil
### GetAccessSids

`func (o *UpdateViewParam) GetAccessSids() []string`

GetAccessSids returns the AccessSids field if non-nil, zero value otherwise.

### GetAccessSidsOk

`func (o *UpdateViewParam) GetAccessSidsOk() (*[]string, bool)`

GetAccessSidsOk returns a tuple with the AccessSids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessSids

`func (o *UpdateViewParam) SetAccessSids(v []string)`

SetAccessSids sets AccessSids field to given value.

### HasAccessSids

`func (o *UpdateViewParam) HasAccessSids() bool`

HasAccessSids returns a boolean if a field has been set.

### SetAccessSidsNil

`func (o *UpdateViewParam) SetAccessSidsNil(b bool)`

 SetAccessSidsNil sets the value for AccessSids to be an explicit nil

### UnsetAccessSids
`func (o *UpdateViewParam) UnsetAccessSids()`

UnsetAccessSids ensures that no value is present for AccessSids, not even an explicit nil
### GetViewLockEnabled

`func (o *UpdateViewParam) GetViewLockEnabled() bool`

GetViewLockEnabled returns the ViewLockEnabled field if non-nil, zero value otherwise.

### GetViewLockEnabledOk

`func (o *UpdateViewParam) GetViewLockEnabledOk() (*bool, bool)`

GetViewLockEnabledOk returns a tuple with the ViewLockEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewLockEnabled

`func (o *UpdateViewParam) SetViewLockEnabled(v bool)`

SetViewLockEnabled sets ViewLockEnabled field to given value.

### HasViewLockEnabled

`func (o *UpdateViewParam) HasViewLockEnabled() bool`

HasViewLockEnabled returns a boolean if a field has been set.

### SetViewLockEnabledNil

`func (o *UpdateViewParam) SetViewLockEnabledNil(b bool)`

 SetViewLockEnabledNil sets the value for ViewLockEnabled to be an explicit nil

### UnsetViewLockEnabled
`func (o *UpdateViewParam) UnsetViewLockEnabled()`

UnsetViewLockEnabled ensures that no value is present for ViewLockEnabled, not even an explicit nil
### GetIsReadOnly

`func (o *UpdateViewParam) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *UpdateViewParam) GetIsReadOnlyOk() (*bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadOnly

`func (o *UpdateViewParam) SetIsReadOnly(v bool)`

SetIsReadOnly sets IsReadOnly field to given value.

### HasIsReadOnly

`func (o *UpdateViewParam) HasIsReadOnly() bool`

HasIsReadOnly returns a boolean if a field has been set.

### SetIsReadOnlyNil

`func (o *UpdateViewParam) SetIsReadOnlyNil(b bool)`

 SetIsReadOnlyNil sets the value for IsReadOnly to be an explicit nil

### UnsetIsReadOnly
`func (o *UpdateViewParam) UnsetIsReadOnly()`

UnsetIsReadOnly ensures that no value is present for IsReadOnly, not even an explicit nil
### GetViewPinningConfig

`func (o *UpdateViewParam) GetViewPinningConfig() ViewPinningConfig`

GetViewPinningConfig returns the ViewPinningConfig field if non-nil, zero value otherwise.

### GetViewPinningConfigOk

`func (o *UpdateViewParam) GetViewPinningConfigOk() (*ViewPinningConfig, bool)`

GetViewPinningConfigOk returns a tuple with the ViewPinningConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewPinningConfig

`func (o *UpdateViewParam) SetViewPinningConfig(v ViewPinningConfig)`

SetViewPinningConfig sets ViewPinningConfig field to given value.

### HasViewPinningConfig

`func (o *UpdateViewParam) HasViewPinningConfig() bool`

HasViewPinningConfig returns a boolean if a field has been set.

### GetSelfServiceSnapshotConfig

`func (o *UpdateViewParam) GetSelfServiceSnapshotConfig() SelfServiceSnapshotConfig`

GetSelfServiceSnapshotConfig returns the SelfServiceSnapshotConfig field if non-nil, zero value otherwise.

### GetSelfServiceSnapshotConfigOk

`func (o *UpdateViewParam) GetSelfServiceSnapshotConfigOk() (*SelfServiceSnapshotConfig, bool)`

GetSelfServiceSnapshotConfigOk returns a tuple with the SelfServiceSnapshotConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfServiceSnapshotConfig

`func (o *UpdateViewParam) SetSelfServiceSnapshotConfig(v SelfServiceSnapshotConfig)`

SetSelfServiceSnapshotConfig sets SelfServiceSnapshotConfig field to given value.

### HasSelfServiceSnapshotConfig

`func (o *UpdateViewParam) HasSelfServiceSnapshotConfig() bool`

HasSelfServiceSnapshotConfig returns a boolean if a field has been set.

### GetEnableMetadataAccelerator

`func (o *UpdateViewParam) GetEnableMetadataAccelerator() bool`

GetEnableMetadataAccelerator returns the EnableMetadataAccelerator field if non-nil, zero value otherwise.

### GetEnableMetadataAcceleratorOk

`func (o *UpdateViewParam) GetEnableMetadataAcceleratorOk() (*bool, bool)`

GetEnableMetadataAcceleratorOk returns a tuple with the EnableMetadataAccelerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMetadataAccelerator

`func (o *UpdateViewParam) SetEnableMetadataAccelerator(v bool)`

SetEnableMetadataAccelerator sets EnableMetadataAccelerator field to given value.

### HasEnableMetadataAccelerator

`func (o *UpdateViewParam) HasEnableMetadataAccelerator() bool`

HasEnableMetadataAccelerator returns a boolean if a field has been set.

### SetEnableMetadataAcceleratorNil

`func (o *UpdateViewParam) SetEnableMetadataAcceleratorNil(b bool)`

 SetEnableMetadataAcceleratorNil sets the value for EnableMetadataAccelerator to be an explicit nil

### UnsetEnableMetadataAccelerator
`func (o *UpdateViewParam) UnsetEnableMetadataAccelerator()`

UnsetEnableMetadataAccelerator ensures that no value is present for EnableMetadataAccelerator, not even an explicit nil
### GetIsExternallyTriggeredBackupTarget

`func (o *UpdateViewParam) GetIsExternallyTriggeredBackupTarget() bool`

GetIsExternallyTriggeredBackupTarget returns the IsExternallyTriggeredBackupTarget field if non-nil, zero value otherwise.

### GetIsExternallyTriggeredBackupTargetOk

`func (o *UpdateViewParam) GetIsExternallyTriggeredBackupTargetOk() (*bool, bool)`

GetIsExternallyTriggeredBackupTargetOk returns a tuple with the IsExternallyTriggeredBackupTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternallyTriggeredBackupTarget

`func (o *UpdateViewParam) SetIsExternallyTriggeredBackupTarget(v bool)`

SetIsExternallyTriggeredBackupTarget sets IsExternallyTriggeredBackupTarget field to given value.

### HasIsExternallyTriggeredBackupTarget

`func (o *UpdateViewParam) HasIsExternallyTriggeredBackupTarget() bool`

HasIsExternallyTriggeredBackupTarget returns a boolean if a field has been set.

### SetIsExternallyTriggeredBackupTargetNil

`func (o *UpdateViewParam) SetIsExternallyTriggeredBackupTargetNil(b bool)`

 SetIsExternallyTriggeredBackupTargetNil sets the value for IsExternallyTriggeredBackupTarget to be an explicit nil

### UnsetIsExternallyTriggeredBackupTarget
`func (o *UpdateViewParam) UnsetIsExternallyTriggeredBackupTarget()`

UnsetIsExternallyTriggeredBackupTarget ensures that no value is present for IsExternallyTriggeredBackupTarget, not even an explicit nil
### GetEnableNfsViewDiscovery

`func (o *UpdateViewParam) GetEnableNfsViewDiscovery() bool`

GetEnableNfsViewDiscovery returns the EnableNfsViewDiscovery field if non-nil, zero value otherwise.

### GetEnableNfsViewDiscoveryOk

`func (o *UpdateViewParam) GetEnableNfsViewDiscoveryOk() (*bool, bool)`

GetEnableNfsViewDiscoveryOk returns a tuple with the EnableNfsViewDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNfsViewDiscovery

`func (o *UpdateViewParam) SetEnableNfsViewDiscovery(v bool)`

SetEnableNfsViewDiscovery sets EnableNfsViewDiscovery field to given value.

### HasEnableNfsViewDiscovery

`func (o *UpdateViewParam) HasEnableNfsViewDiscovery() bool`

HasEnableNfsViewDiscovery returns a boolean if a field has been set.

### SetEnableNfsViewDiscoveryNil

`func (o *UpdateViewParam) SetEnableNfsViewDiscoveryNil(b bool)`

 SetEnableNfsViewDiscoveryNil sets the value for EnableNfsViewDiscovery to be an explicit nil

### UnsetEnableNfsViewDiscovery
`func (o *UpdateViewParam) UnsetEnableNfsViewDiscovery()`

UnsetEnableNfsViewDiscovery ensures that no value is present for EnableNfsViewDiscovery, not even an explicit nil
### GetNfsAllSquash

`func (o *UpdateViewParam) GetNfsAllSquash() NfsSquash`

GetNfsAllSquash returns the NfsAllSquash field if non-nil, zero value otherwise.

### GetNfsAllSquashOk

`func (o *UpdateViewParam) GetNfsAllSquashOk() (*NfsSquash, bool)`

GetNfsAllSquashOk returns a tuple with the NfsAllSquash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsAllSquash

`func (o *UpdateViewParam) SetNfsAllSquash(v NfsSquash)`

SetNfsAllSquash sets NfsAllSquash field to given value.

### HasNfsAllSquash

`func (o *UpdateViewParam) HasNfsAllSquash() bool`

HasNfsAllSquash returns a boolean if a field has been set.

### GetNfsRootPermissions

`func (o *UpdateViewParam) GetNfsRootPermissions() NfsRootPermissions`

GetNfsRootPermissions returns the NfsRootPermissions field if non-nil, zero value otherwise.

### GetNfsRootPermissionsOk

`func (o *UpdateViewParam) GetNfsRootPermissionsOk() (*NfsRootPermissions, bool)`

GetNfsRootPermissionsOk returns a tuple with the NfsRootPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsRootPermissions

`func (o *UpdateViewParam) SetNfsRootPermissions(v NfsRootPermissions)`

SetNfsRootPermissions sets NfsRootPermissions field to given value.

### HasNfsRootPermissions

`func (o *UpdateViewParam) HasNfsRootPermissions() bool`

HasNfsRootPermissions returns a boolean if a field has been set.

### GetNfsRootSquash

`func (o *UpdateViewParam) GetNfsRootSquash() NfsSquash`

GetNfsRootSquash returns the NfsRootSquash field if non-nil, zero value otherwise.

### GetNfsRootSquashOk

`func (o *UpdateViewParam) GetNfsRootSquashOk() (*NfsSquash, bool)`

GetNfsRootSquashOk returns a tuple with the NfsRootSquash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsRootSquash

`func (o *UpdateViewParam) SetNfsRootSquash(v NfsSquash)`

SetNfsRootSquash sets NfsRootSquash field to given value.

### HasNfsRootSquash

`func (o *UpdateViewParam) HasNfsRootSquash() bool`

HasNfsRootSquash returns a boolean if a field has been set.

### GetEnableNfsUnixAuthentication

`func (o *UpdateViewParam) GetEnableNfsUnixAuthentication() bool`

GetEnableNfsUnixAuthentication returns the EnableNfsUnixAuthentication field if non-nil, zero value otherwise.

### GetEnableNfsUnixAuthenticationOk

`func (o *UpdateViewParam) GetEnableNfsUnixAuthenticationOk() (*bool, bool)`

GetEnableNfsUnixAuthenticationOk returns a tuple with the EnableNfsUnixAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNfsUnixAuthentication

`func (o *UpdateViewParam) SetEnableNfsUnixAuthentication(v bool)`

SetEnableNfsUnixAuthentication sets EnableNfsUnixAuthentication field to given value.

### HasEnableNfsUnixAuthentication

`func (o *UpdateViewParam) HasEnableNfsUnixAuthentication() bool`

HasEnableNfsUnixAuthentication returns a boolean if a field has been set.

### SetEnableNfsUnixAuthenticationNil

`func (o *UpdateViewParam) SetEnableNfsUnixAuthenticationNil(b bool)`

 SetEnableNfsUnixAuthenticationNil sets the value for EnableNfsUnixAuthentication to be an explicit nil

### UnsetEnableNfsUnixAuthentication
`func (o *UpdateViewParam) UnsetEnableNfsUnixAuthentication()`

UnsetEnableNfsUnixAuthentication ensures that no value is present for EnableNfsUnixAuthentication, not even an explicit nil
### GetEnableNfsKerberosAuthentication

`func (o *UpdateViewParam) GetEnableNfsKerberosAuthentication() bool`

GetEnableNfsKerberosAuthentication returns the EnableNfsKerberosAuthentication field if non-nil, zero value otherwise.

### GetEnableNfsKerberosAuthenticationOk

`func (o *UpdateViewParam) GetEnableNfsKerberosAuthenticationOk() (*bool, bool)`

GetEnableNfsKerberosAuthenticationOk returns a tuple with the EnableNfsKerberosAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNfsKerberosAuthentication

`func (o *UpdateViewParam) SetEnableNfsKerberosAuthentication(v bool)`

SetEnableNfsKerberosAuthentication sets EnableNfsKerberosAuthentication field to given value.

### HasEnableNfsKerberosAuthentication

`func (o *UpdateViewParam) HasEnableNfsKerberosAuthentication() bool`

HasEnableNfsKerberosAuthentication returns a boolean if a field has been set.

### SetEnableNfsKerberosAuthenticationNil

`func (o *UpdateViewParam) SetEnableNfsKerberosAuthenticationNil(b bool)`

 SetEnableNfsKerberosAuthenticationNil sets the value for EnableNfsKerberosAuthentication to be an explicit nil

### UnsetEnableNfsKerberosAuthentication
`func (o *UpdateViewParam) UnsetEnableNfsKerberosAuthentication()`

UnsetEnableNfsKerberosAuthentication ensures that no value is present for EnableNfsKerberosAuthentication, not even an explicit nil
### GetEnableNfsKerberosIntegrity

`func (o *UpdateViewParam) GetEnableNfsKerberosIntegrity() bool`

GetEnableNfsKerberosIntegrity returns the EnableNfsKerberosIntegrity field if non-nil, zero value otherwise.

### GetEnableNfsKerberosIntegrityOk

`func (o *UpdateViewParam) GetEnableNfsKerberosIntegrityOk() (*bool, bool)`

GetEnableNfsKerberosIntegrityOk returns a tuple with the EnableNfsKerberosIntegrity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNfsKerberosIntegrity

`func (o *UpdateViewParam) SetEnableNfsKerberosIntegrity(v bool)`

SetEnableNfsKerberosIntegrity sets EnableNfsKerberosIntegrity field to given value.

### HasEnableNfsKerberosIntegrity

`func (o *UpdateViewParam) HasEnableNfsKerberosIntegrity() bool`

HasEnableNfsKerberosIntegrity returns a boolean if a field has been set.

### SetEnableNfsKerberosIntegrityNil

`func (o *UpdateViewParam) SetEnableNfsKerberosIntegrityNil(b bool)`

 SetEnableNfsKerberosIntegrityNil sets the value for EnableNfsKerberosIntegrity to be an explicit nil

### UnsetEnableNfsKerberosIntegrity
`func (o *UpdateViewParam) UnsetEnableNfsKerberosIntegrity()`

UnsetEnableNfsKerberosIntegrity ensures that no value is present for EnableNfsKerberosIntegrity, not even an explicit nil
### GetEnableNfsKerberosPrivacy

`func (o *UpdateViewParam) GetEnableNfsKerberosPrivacy() bool`

GetEnableNfsKerberosPrivacy returns the EnableNfsKerberosPrivacy field if non-nil, zero value otherwise.

### GetEnableNfsKerberosPrivacyOk

`func (o *UpdateViewParam) GetEnableNfsKerberosPrivacyOk() (*bool, bool)`

GetEnableNfsKerberosPrivacyOk returns a tuple with the EnableNfsKerberosPrivacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNfsKerberosPrivacy

`func (o *UpdateViewParam) SetEnableNfsKerberosPrivacy(v bool)`

SetEnableNfsKerberosPrivacy sets EnableNfsKerberosPrivacy field to given value.

### HasEnableNfsKerberosPrivacy

`func (o *UpdateViewParam) HasEnableNfsKerberosPrivacy() bool`

HasEnableNfsKerberosPrivacy returns a boolean if a field has been set.

### SetEnableNfsKerberosPrivacyNil

`func (o *UpdateViewParam) SetEnableNfsKerberosPrivacyNil(b bool)`

 SetEnableNfsKerberosPrivacyNil sets the value for EnableNfsKerberosPrivacy to be an explicit nil

### UnsetEnableNfsKerberosPrivacy
`func (o *UpdateViewParam) UnsetEnableNfsKerberosPrivacy()`

UnsetEnableNfsKerberosPrivacy ensures that no value is present for EnableNfsKerberosPrivacy, not even an explicit nil
### GetEnableSmbViewDiscovery

`func (o *UpdateViewParam) GetEnableSmbViewDiscovery() bool`

GetEnableSmbViewDiscovery returns the EnableSmbViewDiscovery field if non-nil, zero value otherwise.

### GetEnableSmbViewDiscoveryOk

`func (o *UpdateViewParam) GetEnableSmbViewDiscoveryOk() (*bool, bool)`

GetEnableSmbViewDiscoveryOk returns a tuple with the EnableSmbViewDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSmbViewDiscovery

`func (o *UpdateViewParam) SetEnableSmbViewDiscovery(v bool)`

SetEnableSmbViewDiscovery sets EnableSmbViewDiscovery field to given value.

### HasEnableSmbViewDiscovery

`func (o *UpdateViewParam) HasEnableSmbViewDiscovery() bool`

HasEnableSmbViewDiscovery returns a boolean if a field has been set.

### SetEnableSmbViewDiscoveryNil

`func (o *UpdateViewParam) SetEnableSmbViewDiscoveryNil(b bool)`

 SetEnableSmbViewDiscoveryNil sets the value for EnableSmbViewDiscovery to be an explicit nil

### UnsetEnableSmbViewDiscovery
`func (o *UpdateViewParam) UnsetEnableSmbViewDiscovery()`

UnsetEnableSmbViewDiscovery ensures that no value is present for EnableSmbViewDiscovery, not even an explicit nil
### GetEnableSmbAccessBasedEnumeration

`func (o *UpdateViewParam) GetEnableSmbAccessBasedEnumeration() bool`

GetEnableSmbAccessBasedEnumeration returns the EnableSmbAccessBasedEnumeration field if non-nil, zero value otherwise.

### GetEnableSmbAccessBasedEnumerationOk

`func (o *UpdateViewParam) GetEnableSmbAccessBasedEnumerationOk() (*bool, bool)`

GetEnableSmbAccessBasedEnumerationOk returns a tuple with the EnableSmbAccessBasedEnumeration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSmbAccessBasedEnumeration

`func (o *UpdateViewParam) SetEnableSmbAccessBasedEnumeration(v bool)`

SetEnableSmbAccessBasedEnumeration sets EnableSmbAccessBasedEnumeration field to given value.

### HasEnableSmbAccessBasedEnumeration

`func (o *UpdateViewParam) HasEnableSmbAccessBasedEnumeration() bool`

HasEnableSmbAccessBasedEnumeration returns a boolean if a field has been set.

### SetEnableSmbAccessBasedEnumerationNil

`func (o *UpdateViewParam) SetEnableSmbAccessBasedEnumerationNil(b bool)`

 SetEnableSmbAccessBasedEnumerationNil sets the value for EnableSmbAccessBasedEnumeration to be an explicit nil

### UnsetEnableSmbAccessBasedEnumeration
`func (o *UpdateViewParam) UnsetEnableSmbAccessBasedEnumeration()`

UnsetEnableSmbAccessBasedEnumeration ensures that no value is present for EnableSmbAccessBasedEnumeration, not even an explicit nil
### GetEnableSmbEncryption

`func (o *UpdateViewParam) GetEnableSmbEncryption() bool`

GetEnableSmbEncryption returns the EnableSmbEncryption field if non-nil, zero value otherwise.

### GetEnableSmbEncryptionOk

`func (o *UpdateViewParam) GetEnableSmbEncryptionOk() (*bool, bool)`

GetEnableSmbEncryptionOk returns a tuple with the EnableSmbEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSmbEncryption

`func (o *UpdateViewParam) SetEnableSmbEncryption(v bool)`

SetEnableSmbEncryption sets EnableSmbEncryption field to given value.

### HasEnableSmbEncryption

`func (o *UpdateViewParam) HasEnableSmbEncryption() bool`

HasEnableSmbEncryption returns a boolean if a field has been set.

### SetEnableSmbEncryptionNil

`func (o *UpdateViewParam) SetEnableSmbEncryptionNil(b bool)`

 SetEnableSmbEncryptionNil sets the value for EnableSmbEncryption to be an explicit nil

### UnsetEnableSmbEncryption
`func (o *UpdateViewParam) UnsetEnableSmbEncryption()`

UnsetEnableSmbEncryption ensures that no value is present for EnableSmbEncryption, not even an explicit nil
### GetEnforceSmbEncryption

`func (o *UpdateViewParam) GetEnforceSmbEncryption() bool`

GetEnforceSmbEncryption returns the EnforceSmbEncryption field if non-nil, zero value otherwise.

### GetEnforceSmbEncryptionOk

`func (o *UpdateViewParam) GetEnforceSmbEncryptionOk() (*bool, bool)`

GetEnforceSmbEncryptionOk returns a tuple with the EnforceSmbEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceSmbEncryption

`func (o *UpdateViewParam) SetEnforceSmbEncryption(v bool)`

SetEnforceSmbEncryption sets EnforceSmbEncryption field to given value.

### HasEnforceSmbEncryption

`func (o *UpdateViewParam) HasEnforceSmbEncryption() bool`

HasEnforceSmbEncryption returns a boolean if a field has been set.

### SetEnforceSmbEncryptionNil

`func (o *UpdateViewParam) SetEnforceSmbEncryptionNil(b bool)`

 SetEnforceSmbEncryptionNil sets the value for EnforceSmbEncryption to be an explicit nil

### UnsetEnforceSmbEncryption
`func (o *UpdateViewParam) UnsetEnforceSmbEncryption()`

UnsetEnforceSmbEncryption ensures that no value is present for EnforceSmbEncryption, not even an explicit nil
### GetEnableFastDurableHandle

`func (o *UpdateViewParam) GetEnableFastDurableHandle() bool`

GetEnableFastDurableHandle returns the EnableFastDurableHandle field if non-nil, zero value otherwise.

### GetEnableFastDurableHandleOk

`func (o *UpdateViewParam) GetEnableFastDurableHandleOk() (*bool, bool)`

GetEnableFastDurableHandleOk returns a tuple with the EnableFastDurableHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFastDurableHandle

`func (o *UpdateViewParam) SetEnableFastDurableHandle(v bool)`

SetEnableFastDurableHandle sets EnableFastDurableHandle field to given value.

### HasEnableFastDurableHandle

`func (o *UpdateViewParam) HasEnableFastDurableHandle() bool`

HasEnableFastDurableHandle returns a boolean if a field has been set.

### SetEnableFastDurableHandleNil

`func (o *UpdateViewParam) SetEnableFastDurableHandleNil(b bool)`

 SetEnableFastDurableHandleNil sets the value for EnableFastDurableHandle to be an explicit nil

### UnsetEnableFastDurableHandle
`func (o *UpdateViewParam) UnsetEnableFastDurableHandle()`

UnsetEnableFastDurableHandle ensures that no value is present for EnableFastDurableHandle, not even an explicit nil
### GetEnableSmbOplock

`func (o *UpdateViewParam) GetEnableSmbOplock() bool`

GetEnableSmbOplock returns the EnableSmbOplock field if non-nil, zero value otherwise.

### GetEnableSmbOplockOk

`func (o *UpdateViewParam) GetEnableSmbOplockOk() (*bool, bool)`

GetEnableSmbOplockOk returns a tuple with the EnableSmbOplock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSmbOplock

`func (o *UpdateViewParam) SetEnableSmbOplock(v bool)`

SetEnableSmbOplock sets EnableSmbOplock field to given value.

### HasEnableSmbOplock

`func (o *UpdateViewParam) HasEnableSmbOplock() bool`

HasEnableSmbOplock returns a boolean if a field has been set.

### SetEnableSmbOplockNil

`func (o *UpdateViewParam) SetEnableSmbOplockNil(b bool)`

 SetEnableSmbOplockNil sets the value for EnableSmbOplock to be an explicit nil

### UnsetEnableSmbOplock
`func (o *UpdateViewParam) UnsetEnableSmbOplock()`

UnsetEnableSmbOplock ensures that no value is present for EnableSmbOplock, not even an explicit nil
### GetSmbPermissionsInfo

`func (o *UpdateViewParam) GetSmbPermissionsInfo() SmbPermissionsInfo`

GetSmbPermissionsInfo returns the SmbPermissionsInfo field if non-nil, zero value otherwise.

### GetSmbPermissionsInfoOk

`func (o *UpdateViewParam) GetSmbPermissionsInfoOk() (*SmbPermissionsInfo, bool)`

GetSmbPermissionsInfoOk returns a tuple with the SmbPermissionsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbPermissionsInfo

`func (o *UpdateViewParam) SetSmbPermissionsInfo(v SmbPermissionsInfo)`

SetSmbPermissionsInfo sets SmbPermissionsInfo field to given value.

### HasSmbPermissionsInfo

`func (o *UpdateViewParam) HasSmbPermissionsInfo() bool`

HasSmbPermissionsInfo returns a boolean if a field has been set.

### GetSharePermissions

`func (o *UpdateViewParam) GetSharePermissions() ViewSharePermissions`

GetSharePermissions returns the SharePermissions field if non-nil, zero value otherwise.

### GetSharePermissionsOk

`func (o *UpdateViewParam) GetSharePermissionsOk() (*ViewSharePermissions, bool)`

GetSharePermissionsOk returns a tuple with the SharePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharePermissions

`func (o *UpdateViewParam) SetSharePermissions(v ViewSharePermissions)`

SetSharePermissions sets SharePermissions field to given value.

### HasSharePermissions

`func (o *UpdateViewParam) HasSharePermissions() bool`

HasSharePermissions returns a boolean if a field has been set.

### GetS3AccessPath

`func (o *UpdateViewParam) GetS3AccessPath() string`

GetS3AccessPath returns the S3AccessPath field if non-nil, zero value otherwise.

### GetS3AccessPathOk

`func (o *UpdateViewParam) GetS3AccessPathOk() (*string, bool)`

GetS3AccessPathOk returns a tuple with the S3AccessPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3AccessPath

`func (o *UpdateViewParam) SetS3AccessPath(v string)`

SetS3AccessPath sets S3AccessPath field to given value.

### HasS3AccessPath

`func (o *UpdateViewParam) HasS3AccessPath() bool`

HasS3AccessPath returns a boolean if a field has been set.

### SetS3AccessPathNil

`func (o *UpdateViewParam) SetS3AccessPathNil(b bool)`

 SetS3AccessPathNil sets the value for S3AccessPath to be an explicit nil

### UnsetS3AccessPath
`func (o *UpdateViewParam) UnsetS3AccessPath()`

UnsetS3AccessPath ensures that no value is present for S3AccessPath, not even an explicit nil
### GetAclConfig

`func (o *UpdateViewParam) GetAclConfig() AclConfig`

GetAclConfig returns the AclConfig field if non-nil, zero value otherwise.

### GetAclConfigOk

`func (o *UpdateViewParam) GetAclConfigOk() (*AclConfig, bool)`

GetAclConfigOk returns a tuple with the AclConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclConfig

`func (o *UpdateViewParam) SetAclConfig(v AclConfig)`

SetAclConfig sets AclConfig field to given value.

### HasAclConfig

`func (o *UpdateViewParam) HasAclConfig() bool`

HasAclConfig returns a boolean if a field has been set.

### GetOwnerInfo

`func (o *UpdateViewParam) GetOwnerInfo() S3OwnerInfo`

GetOwnerInfo returns the OwnerInfo field if non-nil, zero value otherwise.

### GetOwnerInfoOk

`func (o *UpdateViewParam) GetOwnerInfoOk() (*S3OwnerInfo, bool)`

GetOwnerInfoOk returns a tuple with the OwnerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerInfo

`func (o *UpdateViewParam) SetOwnerInfo(v S3OwnerInfo)`

SetOwnerInfo sets OwnerInfo field to given value.

### HasOwnerInfo

`func (o *UpdateViewParam) HasOwnerInfo() bool`

HasOwnerInfo returns a boolean if a field has been set.

### GetSwiftProjectDomain

`func (o *UpdateViewParam) GetSwiftProjectDomain() string`

GetSwiftProjectDomain returns the SwiftProjectDomain field if non-nil, zero value otherwise.

### GetSwiftProjectDomainOk

`func (o *UpdateViewParam) GetSwiftProjectDomainOk() (*string, bool)`

GetSwiftProjectDomainOk returns a tuple with the SwiftProjectDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftProjectDomain

`func (o *UpdateViewParam) SetSwiftProjectDomain(v string)`

SetSwiftProjectDomain sets SwiftProjectDomain field to given value.

### HasSwiftProjectDomain

`func (o *UpdateViewParam) HasSwiftProjectDomain() bool`

HasSwiftProjectDomain returns a boolean if a field has been set.

### SetSwiftProjectDomainNil

`func (o *UpdateViewParam) SetSwiftProjectDomainNil(b bool)`

 SetSwiftProjectDomainNil sets the value for SwiftProjectDomain to be an explicit nil

### UnsetSwiftProjectDomain
`func (o *UpdateViewParam) UnsetSwiftProjectDomain()`

UnsetSwiftProjectDomain ensures that no value is present for SwiftProjectDomain, not even an explicit nil
### GetSwiftProjectName

`func (o *UpdateViewParam) GetSwiftProjectName() string`

GetSwiftProjectName returns the SwiftProjectName field if non-nil, zero value otherwise.

### GetSwiftProjectNameOk

`func (o *UpdateViewParam) GetSwiftProjectNameOk() (*string, bool)`

GetSwiftProjectNameOk returns a tuple with the SwiftProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftProjectName

`func (o *UpdateViewParam) SetSwiftProjectName(v string)`

SetSwiftProjectName sets SwiftProjectName field to given value.

### HasSwiftProjectName

`func (o *UpdateViewParam) HasSwiftProjectName() bool`

HasSwiftProjectName returns a boolean if a field has been set.

### SetSwiftProjectNameNil

`func (o *UpdateViewParam) SetSwiftProjectNameNil(b bool)`

 SetSwiftProjectNameNil sets the value for SwiftProjectName to be an explicit nil

### UnsetSwiftProjectName
`func (o *UpdateViewParam) UnsetSwiftProjectName()`

UnsetSwiftProjectName ensures that no value is present for SwiftProjectName, not even an explicit nil
### GetSwiftUserDomain

`func (o *UpdateViewParam) GetSwiftUserDomain() string`

GetSwiftUserDomain returns the SwiftUserDomain field if non-nil, zero value otherwise.

### GetSwiftUserDomainOk

`func (o *UpdateViewParam) GetSwiftUserDomainOk() (*string, bool)`

GetSwiftUserDomainOk returns a tuple with the SwiftUserDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftUserDomain

`func (o *UpdateViewParam) SetSwiftUserDomain(v string)`

SetSwiftUserDomain sets SwiftUserDomain field to given value.

### HasSwiftUserDomain

`func (o *UpdateViewParam) HasSwiftUserDomain() bool`

HasSwiftUserDomain returns a boolean if a field has been set.

### SetSwiftUserDomainNil

`func (o *UpdateViewParam) SetSwiftUserDomainNil(b bool)`

 SetSwiftUserDomainNil sets the value for SwiftUserDomain to be an explicit nil

### UnsetSwiftUserDomain
`func (o *UpdateViewParam) UnsetSwiftUserDomain()`

UnsetSwiftUserDomain ensures that no value is present for SwiftUserDomain, not even an explicit nil
### GetSwiftUsername

`func (o *UpdateViewParam) GetSwiftUsername() string`

GetSwiftUsername returns the SwiftUsername field if non-nil, zero value otherwise.

### GetSwiftUsernameOk

`func (o *UpdateViewParam) GetSwiftUsernameOk() (*string, bool)`

GetSwiftUsernameOk returns a tuple with the SwiftUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftUsername

`func (o *UpdateViewParam) SetSwiftUsername(v string)`

SetSwiftUsername sets SwiftUsername field to given value.

### HasSwiftUsername

`func (o *UpdateViewParam) HasSwiftUsername() bool`

HasSwiftUsername returns a boolean if a field has been set.

### SetSwiftUsernameNil

`func (o *UpdateViewParam) SetSwiftUsernameNil(b bool)`

 SetSwiftUsernameNil sets the value for SwiftUsername to be an explicit nil

### UnsetSwiftUsername
`func (o *UpdateViewParam) UnsetSwiftUsername()`

UnsetSwiftUsername ensures that no value is present for SwiftUsername, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


