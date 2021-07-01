# Tenant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveDirectories** | Pointer to [**[]ActiveDirectoryEntry**](ActiveDirectoryEntry.md) | Specifies the active directories this tenant is associated to. | [optional] 
**BifrostEnabled** | Pointer to **NullableBool** | Specifies whether bifrost (Ambassador proxy) is enabled for tenant. | [optional] 
**ClusterHostname** | Pointer to **NullableString** | The hostname for Cohesity cluster as seen by tenants and as is routable from the tenant&#39;s network. Tenant&#39;s VLAN&#39;s hostname, if available can be used instead but it is mandatory to provide this value if there&#39;s no VLAN hostname to use. Also, when set, this field would take precedence over VLAN hostname. | [optional] 
**ClusterIps** | Pointer to **[]string** | Set of IPs as seen from the tenant&#39;s network for the Cohesity cluster. Only one from &#39;ClusterHostname&#39; and &#39;ClusterIps&#39; is needed. | [optional] 
**CreatedTimeMsecs** | Pointer to **NullableInt64** | Specifies the epoch time in milliseconds when the tenant account was created on the Cohesity Cluster. | [optional] 
**Deleted** | Pointer to **NullableBool** | Specifies if the Tenant is deleted. | [optional] 
**DeletedTimeMsecs** | Pointer to **NullableInt64** | Specifies the timestamp at which the tenant was deleted. | [optional] 
**DeletionFinished** | Pointer to **NullableBool** | Specifies if the object collection is complete for the tenant. | [optional] 
**DeletionInfoVec** | Pointer to [**[]TenantDeletionInfo**](TenantDeletionInfo.md) | Specifies the current deletion state of object categories. | [optional] 
**Description** | Pointer to **NullableString** | Specifies the description of this tenant. | [optional] 
**EntityIds** | Pointer to **[]int64** | Specifies the EntityIds this tenant is associated to. | [optional] 
**LastUpdatedTimeMsecs** | Pointer to **NullableInt64** | Specifies the epoch time in milliseconds when the tenant account was last modified on the Cohesity Cluster. | [optional] 
**LdapProviders** | Pointer to [**[]LdapProviderResponse**](LdapProviderResponse.md) | Specifies the ldap providers this tenant is associated to. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the tenant. | [optional] 
**OrgSuffix** | Pointer to **NullableString** | Specifies the organization suffix needed to construct tenant id. Tenant id is not completely auto generated rather chosen by tenant/SP admin as we needed same tenant id on multiple clusters to support replication across clusters, etc. | [optional] 
**ParentTenantId** | Pointer to **NullableString** | Specifies the parent tenant of this tenant if available. | [optional] 
**PolicyIds** | Pointer to **[]string** | Specifies the PolicyIds this tenant is associated to. | [optional] 
**ProtectionJobs** | Pointer to [**[]BackupJobProto**](BackupJobProto.md) | Specifies the ProtectionJobs this tenant is associated to. | [optional] 
**SubscribeToAlertEmails** | Pointer to **NullableBool** | Service provider can optionally unsubscribe from the Tenant Alert Emails. | [optional] 
**SwiftConfig** | Pointer to [**SwiftParams**](SwiftParams.md) |  | [optional] 
**TenantId** | Pointer to **NullableString** | Specifies the unique id of the tenant. | [optional] 
**ViewBoxIds** | Pointer to **[]int64** | Specifies the ViewBoxIds this tenant is associated to. | [optional] 
**Views** | Pointer to [**[]View**](View.md) | Specifies the Views this tenant is associated to. | [optional] 
**VlanIfaceNames** | Pointer to **[]string** | Specifies the VlanIfaceNames this tenant is associated to, in the format of bond1.200. | [optional] 

## Methods

### NewTenant

`func NewTenant() *Tenant`

NewTenant instantiates a new Tenant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantWithDefaults

`func NewTenantWithDefaults() *Tenant`

NewTenantWithDefaults instantiates a new Tenant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveDirectories

`func (o *Tenant) GetActiveDirectories() []ActiveDirectoryEntry`

GetActiveDirectories returns the ActiveDirectories field if non-nil, zero value otherwise.

### GetActiveDirectoriesOk

`func (o *Tenant) GetActiveDirectoriesOk() (*[]ActiveDirectoryEntry, bool)`

GetActiveDirectoriesOk returns a tuple with the ActiveDirectories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDirectories

`func (o *Tenant) SetActiveDirectories(v []ActiveDirectoryEntry)`

SetActiveDirectories sets ActiveDirectories field to given value.

### HasActiveDirectories

`func (o *Tenant) HasActiveDirectories() bool`

HasActiveDirectories returns a boolean if a field has been set.

### SetActiveDirectoriesNil

`func (o *Tenant) SetActiveDirectoriesNil(b bool)`

 SetActiveDirectoriesNil sets the value for ActiveDirectories to be an explicit nil

### UnsetActiveDirectories
`func (o *Tenant) UnsetActiveDirectories()`

UnsetActiveDirectories ensures that no value is present for ActiveDirectories, not even an explicit nil
### GetBifrostEnabled

`func (o *Tenant) GetBifrostEnabled() bool`

GetBifrostEnabled returns the BifrostEnabled field if non-nil, zero value otherwise.

### GetBifrostEnabledOk

`func (o *Tenant) GetBifrostEnabledOk() (*bool, bool)`

GetBifrostEnabledOk returns a tuple with the BifrostEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBifrostEnabled

`func (o *Tenant) SetBifrostEnabled(v bool)`

SetBifrostEnabled sets BifrostEnabled field to given value.

### HasBifrostEnabled

`func (o *Tenant) HasBifrostEnabled() bool`

HasBifrostEnabled returns a boolean if a field has been set.

### SetBifrostEnabledNil

`func (o *Tenant) SetBifrostEnabledNil(b bool)`

 SetBifrostEnabledNil sets the value for BifrostEnabled to be an explicit nil

### UnsetBifrostEnabled
`func (o *Tenant) UnsetBifrostEnabled()`

UnsetBifrostEnabled ensures that no value is present for BifrostEnabled, not even an explicit nil
### GetClusterHostname

`func (o *Tenant) GetClusterHostname() string`

GetClusterHostname returns the ClusterHostname field if non-nil, zero value otherwise.

### GetClusterHostnameOk

`func (o *Tenant) GetClusterHostnameOk() (*string, bool)`

GetClusterHostnameOk returns a tuple with the ClusterHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterHostname

`func (o *Tenant) SetClusterHostname(v string)`

SetClusterHostname sets ClusterHostname field to given value.

### HasClusterHostname

`func (o *Tenant) HasClusterHostname() bool`

HasClusterHostname returns a boolean if a field has been set.

### SetClusterHostnameNil

`func (o *Tenant) SetClusterHostnameNil(b bool)`

 SetClusterHostnameNil sets the value for ClusterHostname to be an explicit nil

### UnsetClusterHostname
`func (o *Tenant) UnsetClusterHostname()`

UnsetClusterHostname ensures that no value is present for ClusterHostname, not even an explicit nil
### GetClusterIps

`func (o *Tenant) GetClusterIps() []string`

GetClusterIps returns the ClusterIps field if non-nil, zero value otherwise.

### GetClusterIpsOk

`func (o *Tenant) GetClusterIpsOk() (*[]string, bool)`

GetClusterIpsOk returns a tuple with the ClusterIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIps

`func (o *Tenant) SetClusterIps(v []string)`

SetClusterIps sets ClusterIps field to given value.

### HasClusterIps

`func (o *Tenant) HasClusterIps() bool`

HasClusterIps returns a boolean if a field has been set.

### SetClusterIpsNil

`func (o *Tenant) SetClusterIpsNil(b bool)`

 SetClusterIpsNil sets the value for ClusterIps to be an explicit nil

### UnsetClusterIps
`func (o *Tenant) UnsetClusterIps()`

UnsetClusterIps ensures that no value is present for ClusterIps, not even an explicit nil
### GetCreatedTimeMsecs

`func (o *Tenant) GetCreatedTimeMsecs() int64`

GetCreatedTimeMsecs returns the CreatedTimeMsecs field if non-nil, zero value otherwise.

### GetCreatedTimeMsecsOk

`func (o *Tenant) GetCreatedTimeMsecsOk() (*int64, bool)`

GetCreatedTimeMsecsOk returns a tuple with the CreatedTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimeMsecs

`func (o *Tenant) SetCreatedTimeMsecs(v int64)`

SetCreatedTimeMsecs sets CreatedTimeMsecs field to given value.

### HasCreatedTimeMsecs

`func (o *Tenant) HasCreatedTimeMsecs() bool`

HasCreatedTimeMsecs returns a boolean if a field has been set.

### SetCreatedTimeMsecsNil

`func (o *Tenant) SetCreatedTimeMsecsNil(b bool)`

 SetCreatedTimeMsecsNil sets the value for CreatedTimeMsecs to be an explicit nil

### UnsetCreatedTimeMsecs
`func (o *Tenant) UnsetCreatedTimeMsecs()`

UnsetCreatedTimeMsecs ensures that no value is present for CreatedTimeMsecs, not even an explicit nil
### GetDeleted

`func (o *Tenant) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Tenant) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Tenant) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Tenant) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### SetDeletedNil

`func (o *Tenant) SetDeletedNil(b bool)`

 SetDeletedNil sets the value for Deleted to be an explicit nil

### UnsetDeleted
`func (o *Tenant) UnsetDeleted()`

UnsetDeleted ensures that no value is present for Deleted, not even an explicit nil
### GetDeletedTimeMsecs

`func (o *Tenant) GetDeletedTimeMsecs() int64`

GetDeletedTimeMsecs returns the DeletedTimeMsecs field if non-nil, zero value otherwise.

### GetDeletedTimeMsecsOk

`func (o *Tenant) GetDeletedTimeMsecsOk() (*int64, bool)`

GetDeletedTimeMsecsOk returns a tuple with the DeletedTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedTimeMsecs

`func (o *Tenant) SetDeletedTimeMsecs(v int64)`

SetDeletedTimeMsecs sets DeletedTimeMsecs field to given value.

### HasDeletedTimeMsecs

`func (o *Tenant) HasDeletedTimeMsecs() bool`

HasDeletedTimeMsecs returns a boolean if a field has been set.

### SetDeletedTimeMsecsNil

`func (o *Tenant) SetDeletedTimeMsecsNil(b bool)`

 SetDeletedTimeMsecsNil sets the value for DeletedTimeMsecs to be an explicit nil

### UnsetDeletedTimeMsecs
`func (o *Tenant) UnsetDeletedTimeMsecs()`

UnsetDeletedTimeMsecs ensures that no value is present for DeletedTimeMsecs, not even an explicit nil
### GetDeletionFinished

`func (o *Tenant) GetDeletionFinished() bool`

GetDeletionFinished returns the DeletionFinished field if non-nil, zero value otherwise.

### GetDeletionFinishedOk

`func (o *Tenant) GetDeletionFinishedOk() (*bool, bool)`

GetDeletionFinishedOk returns a tuple with the DeletionFinished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionFinished

`func (o *Tenant) SetDeletionFinished(v bool)`

SetDeletionFinished sets DeletionFinished field to given value.

### HasDeletionFinished

`func (o *Tenant) HasDeletionFinished() bool`

HasDeletionFinished returns a boolean if a field has been set.

### SetDeletionFinishedNil

`func (o *Tenant) SetDeletionFinishedNil(b bool)`

 SetDeletionFinishedNil sets the value for DeletionFinished to be an explicit nil

### UnsetDeletionFinished
`func (o *Tenant) UnsetDeletionFinished()`

UnsetDeletionFinished ensures that no value is present for DeletionFinished, not even an explicit nil
### GetDeletionInfoVec

`func (o *Tenant) GetDeletionInfoVec() []TenantDeletionInfo`

GetDeletionInfoVec returns the DeletionInfoVec field if non-nil, zero value otherwise.

### GetDeletionInfoVecOk

`func (o *Tenant) GetDeletionInfoVecOk() (*[]TenantDeletionInfo, bool)`

GetDeletionInfoVecOk returns a tuple with the DeletionInfoVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionInfoVec

`func (o *Tenant) SetDeletionInfoVec(v []TenantDeletionInfo)`

SetDeletionInfoVec sets DeletionInfoVec field to given value.

### HasDeletionInfoVec

`func (o *Tenant) HasDeletionInfoVec() bool`

HasDeletionInfoVec returns a boolean if a field has been set.

### SetDeletionInfoVecNil

`func (o *Tenant) SetDeletionInfoVecNil(b bool)`

 SetDeletionInfoVecNil sets the value for DeletionInfoVec to be an explicit nil

### UnsetDeletionInfoVec
`func (o *Tenant) UnsetDeletionInfoVec()`

UnsetDeletionInfoVec ensures that no value is present for DeletionInfoVec, not even an explicit nil
### GetDescription

`func (o *Tenant) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Tenant) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Tenant) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Tenant) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Tenant) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Tenant) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEntityIds

`func (o *Tenant) GetEntityIds() []int64`

GetEntityIds returns the EntityIds field if non-nil, zero value otherwise.

### GetEntityIdsOk

`func (o *Tenant) GetEntityIdsOk() (*[]int64, bool)`

GetEntityIdsOk returns a tuple with the EntityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityIds

`func (o *Tenant) SetEntityIds(v []int64)`

SetEntityIds sets EntityIds field to given value.

### HasEntityIds

`func (o *Tenant) HasEntityIds() bool`

HasEntityIds returns a boolean if a field has been set.

### SetEntityIdsNil

`func (o *Tenant) SetEntityIdsNil(b bool)`

 SetEntityIdsNil sets the value for EntityIds to be an explicit nil

### UnsetEntityIds
`func (o *Tenant) UnsetEntityIds()`

UnsetEntityIds ensures that no value is present for EntityIds, not even an explicit nil
### GetLastUpdatedTimeMsecs

`func (o *Tenant) GetLastUpdatedTimeMsecs() int64`

GetLastUpdatedTimeMsecs returns the LastUpdatedTimeMsecs field if non-nil, zero value otherwise.

### GetLastUpdatedTimeMsecsOk

`func (o *Tenant) GetLastUpdatedTimeMsecsOk() (*int64, bool)`

GetLastUpdatedTimeMsecsOk returns a tuple with the LastUpdatedTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTimeMsecs

`func (o *Tenant) SetLastUpdatedTimeMsecs(v int64)`

SetLastUpdatedTimeMsecs sets LastUpdatedTimeMsecs field to given value.

### HasLastUpdatedTimeMsecs

`func (o *Tenant) HasLastUpdatedTimeMsecs() bool`

HasLastUpdatedTimeMsecs returns a boolean if a field has been set.

### SetLastUpdatedTimeMsecsNil

`func (o *Tenant) SetLastUpdatedTimeMsecsNil(b bool)`

 SetLastUpdatedTimeMsecsNil sets the value for LastUpdatedTimeMsecs to be an explicit nil

### UnsetLastUpdatedTimeMsecs
`func (o *Tenant) UnsetLastUpdatedTimeMsecs()`

UnsetLastUpdatedTimeMsecs ensures that no value is present for LastUpdatedTimeMsecs, not even an explicit nil
### GetLdapProviders

`func (o *Tenant) GetLdapProviders() []LdapProviderResponse`

GetLdapProviders returns the LdapProviders field if non-nil, zero value otherwise.

### GetLdapProvidersOk

`func (o *Tenant) GetLdapProvidersOk() (*[]LdapProviderResponse, bool)`

GetLdapProvidersOk returns a tuple with the LdapProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapProviders

`func (o *Tenant) SetLdapProviders(v []LdapProviderResponse)`

SetLdapProviders sets LdapProviders field to given value.

### HasLdapProviders

`func (o *Tenant) HasLdapProviders() bool`

HasLdapProviders returns a boolean if a field has been set.

### SetLdapProvidersNil

`func (o *Tenant) SetLdapProvidersNil(b bool)`

 SetLdapProvidersNil sets the value for LdapProviders to be an explicit nil

### UnsetLdapProviders
`func (o *Tenant) UnsetLdapProviders()`

UnsetLdapProviders ensures that no value is present for LdapProviders, not even an explicit nil
### GetName

`func (o *Tenant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Tenant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Tenant) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Tenant) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Tenant) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Tenant) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOrgSuffix

`func (o *Tenant) GetOrgSuffix() string`

GetOrgSuffix returns the OrgSuffix field if non-nil, zero value otherwise.

### GetOrgSuffixOk

`func (o *Tenant) GetOrgSuffixOk() (*string, bool)`

GetOrgSuffixOk returns a tuple with the OrgSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgSuffix

`func (o *Tenant) SetOrgSuffix(v string)`

SetOrgSuffix sets OrgSuffix field to given value.

### HasOrgSuffix

`func (o *Tenant) HasOrgSuffix() bool`

HasOrgSuffix returns a boolean if a field has been set.

### SetOrgSuffixNil

`func (o *Tenant) SetOrgSuffixNil(b bool)`

 SetOrgSuffixNil sets the value for OrgSuffix to be an explicit nil

### UnsetOrgSuffix
`func (o *Tenant) UnsetOrgSuffix()`

UnsetOrgSuffix ensures that no value is present for OrgSuffix, not even an explicit nil
### GetParentTenantId

`func (o *Tenant) GetParentTenantId() string`

GetParentTenantId returns the ParentTenantId field if non-nil, zero value otherwise.

### GetParentTenantIdOk

`func (o *Tenant) GetParentTenantIdOk() (*string, bool)`

GetParentTenantIdOk returns a tuple with the ParentTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTenantId

`func (o *Tenant) SetParentTenantId(v string)`

SetParentTenantId sets ParentTenantId field to given value.

### HasParentTenantId

`func (o *Tenant) HasParentTenantId() bool`

HasParentTenantId returns a boolean if a field has been set.

### SetParentTenantIdNil

`func (o *Tenant) SetParentTenantIdNil(b bool)`

 SetParentTenantIdNil sets the value for ParentTenantId to be an explicit nil

### UnsetParentTenantId
`func (o *Tenant) UnsetParentTenantId()`

UnsetParentTenantId ensures that no value is present for ParentTenantId, not even an explicit nil
### GetPolicyIds

`func (o *Tenant) GetPolicyIds() []string`

GetPolicyIds returns the PolicyIds field if non-nil, zero value otherwise.

### GetPolicyIdsOk

`func (o *Tenant) GetPolicyIdsOk() (*[]string, bool)`

GetPolicyIdsOk returns a tuple with the PolicyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyIds

`func (o *Tenant) SetPolicyIds(v []string)`

SetPolicyIds sets PolicyIds field to given value.

### HasPolicyIds

`func (o *Tenant) HasPolicyIds() bool`

HasPolicyIds returns a boolean if a field has been set.

### SetPolicyIdsNil

`func (o *Tenant) SetPolicyIdsNil(b bool)`

 SetPolicyIdsNil sets the value for PolicyIds to be an explicit nil

### UnsetPolicyIds
`func (o *Tenant) UnsetPolicyIds()`

UnsetPolicyIds ensures that no value is present for PolicyIds, not even an explicit nil
### GetProtectionJobs

`func (o *Tenant) GetProtectionJobs() []BackupJobProto`

GetProtectionJobs returns the ProtectionJobs field if non-nil, zero value otherwise.

### GetProtectionJobsOk

`func (o *Tenant) GetProtectionJobsOk() (*[]BackupJobProto, bool)`

GetProtectionJobsOk returns a tuple with the ProtectionJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionJobs

`func (o *Tenant) SetProtectionJobs(v []BackupJobProto)`

SetProtectionJobs sets ProtectionJobs field to given value.

### HasProtectionJobs

`func (o *Tenant) HasProtectionJobs() bool`

HasProtectionJobs returns a boolean if a field has been set.

### SetProtectionJobsNil

`func (o *Tenant) SetProtectionJobsNil(b bool)`

 SetProtectionJobsNil sets the value for ProtectionJobs to be an explicit nil

### UnsetProtectionJobs
`func (o *Tenant) UnsetProtectionJobs()`

UnsetProtectionJobs ensures that no value is present for ProtectionJobs, not even an explicit nil
### GetSubscribeToAlertEmails

`func (o *Tenant) GetSubscribeToAlertEmails() bool`

GetSubscribeToAlertEmails returns the SubscribeToAlertEmails field if non-nil, zero value otherwise.

### GetSubscribeToAlertEmailsOk

`func (o *Tenant) GetSubscribeToAlertEmailsOk() (*bool, bool)`

GetSubscribeToAlertEmailsOk returns a tuple with the SubscribeToAlertEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribeToAlertEmails

`func (o *Tenant) SetSubscribeToAlertEmails(v bool)`

SetSubscribeToAlertEmails sets SubscribeToAlertEmails field to given value.

### HasSubscribeToAlertEmails

`func (o *Tenant) HasSubscribeToAlertEmails() bool`

HasSubscribeToAlertEmails returns a boolean if a field has been set.

### SetSubscribeToAlertEmailsNil

`func (o *Tenant) SetSubscribeToAlertEmailsNil(b bool)`

 SetSubscribeToAlertEmailsNil sets the value for SubscribeToAlertEmails to be an explicit nil

### UnsetSubscribeToAlertEmails
`func (o *Tenant) UnsetSubscribeToAlertEmails()`

UnsetSubscribeToAlertEmails ensures that no value is present for SubscribeToAlertEmails, not even an explicit nil
### GetSwiftConfig

`func (o *Tenant) GetSwiftConfig() SwiftParams`

GetSwiftConfig returns the SwiftConfig field if non-nil, zero value otherwise.

### GetSwiftConfigOk

`func (o *Tenant) GetSwiftConfigOk() (*SwiftParams, bool)`

GetSwiftConfigOk returns a tuple with the SwiftConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftConfig

`func (o *Tenant) SetSwiftConfig(v SwiftParams)`

SetSwiftConfig sets SwiftConfig field to given value.

### HasSwiftConfig

`func (o *Tenant) HasSwiftConfig() bool`

HasSwiftConfig returns a boolean if a field has been set.

### GetTenantId

`func (o *Tenant) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Tenant) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Tenant) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *Tenant) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *Tenant) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *Tenant) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetViewBoxIds

`func (o *Tenant) GetViewBoxIds() []int64`

GetViewBoxIds returns the ViewBoxIds field if non-nil, zero value otherwise.

### GetViewBoxIdsOk

`func (o *Tenant) GetViewBoxIdsOk() (*[]int64, bool)`

GetViewBoxIdsOk returns a tuple with the ViewBoxIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewBoxIds

`func (o *Tenant) SetViewBoxIds(v []int64)`

SetViewBoxIds sets ViewBoxIds field to given value.

### HasViewBoxIds

`func (o *Tenant) HasViewBoxIds() bool`

HasViewBoxIds returns a boolean if a field has been set.

### SetViewBoxIdsNil

`func (o *Tenant) SetViewBoxIdsNil(b bool)`

 SetViewBoxIdsNil sets the value for ViewBoxIds to be an explicit nil

### UnsetViewBoxIds
`func (o *Tenant) UnsetViewBoxIds()`

UnsetViewBoxIds ensures that no value is present for ViewBoxIds, not even an explicit nil
### GetViews

`func (o *Tenant) GetViews() []View`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *Tenant) GetViewsOk() (*[]View, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *Tenant) SetViews(v []View)`

SetViews sets Views field to given value.

### HasViews

`func (o *Tenant) HasViews() bool`

HasViews returns a boolean if a field has been set.

### SetViewsNil

`func (o *Tenant) SetViewsNil(b bool)`

 SetViewsNil sets the value for Views to be an explicit nil

### UnsetViews
`func (o *Tenant) UnsetViews()`

UnsetViews ensures that no value is present for Views, not even an explicit nil
### GetVlanIfaceNames

`func (o *Tenant) GetVlanIfaceNames() []string`

GetVlanIfaceNames returns the VlanIfaceNames field if non-nil, zero value otherwise.

### GetVlanIfaceNamesOk

`func (o *Tenant) GetVlanIfaceNamesOk() (*[]string, bool)`

GetVlanIfaceNamesOk returns a tuple with the VlanIfaceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanIfaceNames

`func (o *Tenant) SetVlanIfaceNames(v []string)`

SetVlanIfaceNames sets VlanIfaceNames field to given value.

### HasVlanIfaceNames

`func (o *Tenant) HasVlanIfaceNames() bool`

HasVlanIfaceNames returns a boolean if a field has been set.

### SetVlanIfaceNamesNil

`func (o *Tenant) SetVlanIfaceNamesNil(b bool)`

 SetVlanIfaceNamesNil sets the value for VlanIfaceNames to be an explicit nil

### UnsetVlanIfaceNames
`func (o *Tenant) UnsetVlanIfaceNames()`

UnsetVlanIfaceNames ensures that no value is present for VlanIfaceNames, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


