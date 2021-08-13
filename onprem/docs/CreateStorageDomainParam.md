# CreateStorageDomainParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt64** | Specifies the Storage Domain id. | [optional] [readonly] 
**Name** | **NullableString** | Specifies the Storage Domain name. | 
**ClusterPartitionId** | **NullableInt64** | Specifies the cluster partition id of the Storage Domain. | 
**ClusterPartitionName** | Pointer to **NullableString** | Specifies the cluster partition name of the Storage Domain. | [optional] [readonly] 
**SubnetWhitelist** | Pointer to [**[]Subnet**](Subnet.md) | Specifies a list of Subnets with IP addresses that have permissions to access the Storage Domain. | [optional] 
**StoragePolicy** | Pointer to [**StoragePolicy**](StoragePolicy.md) | Specifies the storage policy for this Storage Domain. | [optional] 
**PhysicalQuota** | Pointer to [**QuotaPolicy**](QuotaPolicy.md) | Specifies a quota limit for physical usage of this Storage Domain. This quota defines a limit of data that can be physically (after data size is reduced by block tracking, compression and deduplication) stored on this storage domain. A new write will not be allowed when the storage domain usage will exceeds the specified quota. Due to the latency of calculating usage across all nodes, the actual storage domain usage may exceed the quota limit by a little bit. | [optional] 
**CloudDownWaterFallParams** | Pointer to [**CloudDownWaterFallParams**](CloudDownWaterFallParams.md) | Specifies the cloud down water fall parameters for this Storage Domain. | [optional] 
**DefaultViewQuota** | Pointer to [**QuotaPolicy**](QuotaPolicy.md) | Specifies a default logical quota limit for all views in this Storage Domain. This quota can be overwritten by a view level quota. | [optional] 
**DefaultUserQuota** | Pointer to [**QuotaPolicy**](QuotaPolicy.md) | Specifies a default user quota limit for users within views in this Storage Domain. | [optional] 
**S3BucketsEnabled** | Pointer to **NullableBool** | Specifies whether to enable creation of S3 bucket on this Storage Domain. | [optional] 
**AdDomainName** | Pointer to **NullableString** | Specifies the Active Directory domain name that this Storage Domain is mapped to. | [optional] 
**NisDomainNames** | Pointer to **[]string** | Specifies the NIS domain names that this Storage Domain is mapped to. | [optional] 
**KerberosRealmName** | Pointer to **NullableString** | Specifies the Kerberos realm name that this Storage Domain is mapped to. | [optional] 
**LdapProviderId** | Pointer to **NullableInt64** | Specifies the LDAP provider id that this Storage Domain is mapped to. | [optional] 
**TenantIds** | Pointer to **[]string** | Specifies a list of tenant ids that that Storage Domain belongs. There can only be one tenant id in this field unless Storage Domain sharing between tenants is allowed on this cluster. | [optional] 
**DirectArchiveEnabled** | Pointer to **NullableBool** | Specifies whether to enable driect archive on this Storage Domain. If enabled, this Storage Domain can be used as a staging area while copying a large dataset that can&#39;t fit on the cluster to an external target. | [optional] 
**BlobBrickSizeBytes** | Pointer to **NullableInt32** | Specifies the brick size used for blobs in this Storage Domain. | [optional] 
**KmsServerId** | Pointer to **NullableInt64** | Specifies the associated KMS server id. | [optional] 
**DekRotationEnabled** | Pointer to **NullableBool** | Specifies whether DEK(Data Encryption Key) rotation is enabled for this Storage Domain. This is applicable only when the Storage Domain uses AWS or similar KMS in which the KEK (Key Encryption Key) is not created and maintained by Cohesity. For Internal KMS and keys stored in Safenet servers, DEK rotation will not be performed. | [optional] 
**TreatFileSyncAsDataSync** | Pointer to **NullableBool** | If &#39;true&#39;, when the Cohesity Cluster is writing to a file, the file modification time is not persisted synchronously during the file write, so the modification time may not be accurate. (Typically the file modification time is off by 30 seconds but it can be longer.) | [optional] 
**Recommended** | Pointer to **NullableBool** | Specifies whether Storage Domain is recommended for the specified View template. | [optional] [readonly] 
**Stats** | Pointer to [**DataUsageStats**](DataUsageStats.md) | Specifies the Storage Domain stats. | [optional] 
**Schemas** | Pointer to [**[]Schema**](Schema.md) | Specifies the Storage Domain schemas. | [optional] [readonly] 
**FileCountBySize** | Pointer to [**[]FileCount**](FileCount.md) | Specifies the file count by size for the View. | [optional] [readonly] 
**CloudDomainId** | Pointer to **NullableInt64** | Specifies the cloud domain Id. | [optional] 
**VaultId** | Pointer to **NullableInt64** | Specifies the vault Id associated with cloud domain ID. | [optional] 

## Methods

### NewCreateStorageDomainParam

`func NewCreateStorageDomainParam(name NullableString, clusterPartitionId NullableInt64, ) *CreateStorageDomainParam`

NewCreateStorageDomainParam instantiates a new CreateStorageDomainParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStorageDomainParamWithDefaults

`func NewCreateStorageDomainParamWithDefaults() *CreateStorageDomainParam`

NewCreateStorageDomainParamWithDefaults instantiates a new CreateStorageDomainParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateStorageDomainParam) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateStorageDomainParam) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateStorageDomainParam) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CreateStorageDomainParam) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CreateStorageDomainParam) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CreateStorageDomainParam) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *CreateStorageDomainParam) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateStorageDomainParam) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateStorageDomainParam) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CreateStorageDomainParam) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateStorageDomainParam) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetClusterPartitionId

`func (o *CreateStorageDomainParam) GetClusterPartitionId() int64`

GetClusterPartitionId returns the ClusterPartitionId field if non-nil, zero value otherwise.

### GetClusterPartitionIdOk

`func (o *CreateStorageDomainParam) GetClusterPartitionIdOk() (*int64, bool)`

GetClusterPartitionIdOk returns a tuple with the ClusterPartitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterPartitionId

`func (o *CreateStorageDomainParam) SetClusterPartitionId(v int64)`

SetClusterPartitionId sets ClusterPartitionId field to given value.


### SetClusterPartitionIdNil

`func (o *CreateStorageDomainParam) SetClusterPartitionIdNil(b bool)`

 SetClusterPartitionIdNil sets the value for ClusterPartitionId to be an explicit nil

### UnsetClusterPartitionId
`func (o *CreateStorageDomainParam) UnsetClusterPartitionId()`

UnsetClusterPartitionId ensures that no value is present for ClusterPartitionId, not even an explicit nil
### GetClusterPartitionName

`func (o *CreateStorageDomainParam) GetClusterPartitionName() string`

GetClusterPartitionName returns the ClusterPartitionName field if non-nil, zero value otherwise.

### GetClusterPartitionNameOk

`func (o *CreateStorageDomainParam) GetClusterPartitionNameOk() (*string, bool)`

GetClusterPartitionNameOk returns a tuple with the ClusterPartitionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterPartitionName

`func (o *CreateStorageDomainParam) SetClusterPartitionName(v string)`

SetClusterPartitionName sets ClusterPartitionName field to given value.

### HasClusterPartitionName

`func (o *CreateStorageDomainParam) HasClusterPartitionName() bool`

HasClusterPartitionName returns a boolean if a field has been set.

### SetClusterPartitionNameNil

`func (o *CreateStorageDomainParam) SetClusterPartitionNameNil(b bool)`

 SetClusterPartitionNameNil sets the value for ClusterPartitionName to be an explicit nil

### UnsetClusterPartitionName
`func (o *CreateStorageDomainParam) UnsetClusterPartitionName()`

UnsetClusterPartitionName ensures that no value is present for ClusterPartitionName, not even an explicit nil
### GetSubnetWhitelist

`func (o *CreateStorageDomainParam) GetSubnetWhitelist() []Subnet`

GetSubnetWhitelist returns the SubnetWhitelist field if non-nil, zero value otherwise.

### GetSubnetWhitelistOk

`func (o *CreateStorageDomainParam) GetSubnetWhitelistOk() (*[]Subnet, bool)`

GetSubnetWhitelistOk returns a tuple with the SubnetWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetWhitelist

`func (o *CreateStorageDomainParam) SetSubnetWhitelist(v []Subnet)`

SetSubnetWhitelist sets SubnetWhitelist field to given value.

### HasSubnetWhitelist

`func (o *CreateStorageDomainParam) HasSubnetWhitelist() bool`

HasSubnetWhitelist returns a boolean if a field has been set.

### SetSubnetWhitelistNil

`func (o *CreateStorageDomainParam) SetSubnetWhitelistNil(b bool)`

 SetSubnetWhitelistNil sets the value for SubnetWhitelist to be an explicit nil

### UnsetSubnetWhitelist
`func (o *CreateStorageDomainParam) UnsetSubnetWhitelist()`

UnsetSubnetWhitelist ensures that no value is present for SubnetWhitelist, not even an explicit nil
### GetStoragePolicy

`func (o *CreateStorageDomainParam) GetStoragePolicy() StoragePolicy`

GetStoragePolicy returns the StoragePolicy field if non-nil, zero value otherwise.

### GetStoragePolicyOk

`func (o *CreateStorageDomainParam) GetStoragePolicyOk() (*StoragePolicy, bool)`

GetStoragePolicyOk returns a tuple with the StoragePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePolicy

`func (o *CreateStorageDomainParam) SetStoragePolicy(v StoragePolicy)`

SetStoragePolicy sets StoragePolicy field to given value.

### HasStoragePolicy

`func (o *CreateStorageDomainParam) HasStoragePolicy() bool`

HasStoragePolicy returns a boolean if a field has been set.

### GetPhysicalQuota

`func (o *CreateStorageDomainParam) GetPhysicalQuota() QuotaPolicy`

GetPhysicalQuota returns the PhysicalQuota field if non-nil, zero value otherwise.

### GetPhysicalQuotaOk

`func (o *CreateStorageDomainParam) GetPhysicalQuotaOk() (*QuotaPolicy, bool)`

GetPhysicalQuotaOk returns a tuple with the PhysicalQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalQuota

`func (o *CreateStorageDomainParam) SetPhysicalQuota(v QuotaPolicy)`

SetPhysicalQuota sets PhysicalQuota field to given value.

### HasPhysicalQuota

`func (o *CreateStorageDomainParam) HasPhysicalQuota() bool`

HasPhysicalQuota returns a boolean if a field has been set.

### GetCloudDownWaterFallParams

`func (o *CreateStorageDomainParam) GetCloudDownWaterFallParams() CloudDownWaterFallParams`

GetCloudDownWaterFallParams returns the CloudDownWaterFallParams field if non-nil, zero value otherwise.

### GetCloudDownWaterFallParamsOk

`func (o *CreateStorageDomainParam) GetCloudDownWaterFallParamsOk() (*CloudDownWaterFallParams, bool)`

GetCloudDownWaterFallParamsOk returns a tuple with the CloudDownWaterFallParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudDownWaterFallParams

`func (o *CreateStorageDomainParam) SetCloudDownWaterFallParams(v CloudDownWaterFallParams)`

SetCloudDownWaterFallParams sets CloudDownWaterFallParams field to given value.

### HasCloudDownWaterFallParams

`func (o *CreateStorageDomainParam) HasCloudDownWaterFallParams() bool`

HasCloudDownWaterFallParams returns a boolean if a field has been set.

### GetDefaultViewQuota

`func (o *CreateStorageDomainParam) GetDefaultViewQuota() QuotaPolicy`

GetDefaultViewQuota returns the DefaultViewQuota field if non-nil, zero value otherwise.

### GetDefaultViewQuotaOk

`func (o *CreateStorageDomainParam) GetDefaultViewQuotaOk() (*QuotaPolicy, bool)`

GetDefaultViewQuotaOk returns a tuple with the DefaultViewQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultViewQuota

`func (o *CreateStorageDomainParam) SetDefaultViewQuota(v QuotaPolicy)`

SetDefaultViewQuota sets DefaultViewQuota field to given value.

### HasDefaultViewQuota

`func (o *CreateStorageDomainParam) HasDefaultViewQuota() bool`

HasDefaultViewQuota returns a boolean if a field has been set.

### GetDefaultUserQuota

`func (o *CreateStorageDomainParam) GetDefaultUserQuota() QuotaPolicy`

GetDefaultUserQuota returns the DefaultUserQuota field if non-nil, zero value otherwise.

### GetDefaultUserQuotaOk

`func (o *CreateStorageDomainParam) GetDefaultUserQuotaOk() (*QuotaPolicy, bool)`

GetDefaultUserQuotaOk returns a tuple with the DefaultUserQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUserQuota

`func (o *CreateStorageDomainParam) SetDefaultUserQuota(v QuotaPolicy)`

SetDefaultUserQuota sets DefaultUserQuota field to given value.

### HasDefaultUserQuota

`func (o *CreateStorageDomainParam) HasDefaultUserQuota() bool`

HasDefaultUserQuota returns a boolean if a field has been set.

### GetS3BucketsEnabled

`func (o *CreateStorageDomainParam) GetS3BucketsEnabled() bool`

GetS3BucketsEnabled returns the S3BucketsEnabled field if non-nil, zero value otherwise.

### GetS3BucketsEnabledOk

`func (o *CreateStorageDomainParam) GetS3BucketsEnabledOk() (*bool, bool)`

GetS3BucketsEnabledOk returns a tuple with the S3BucketsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3BucketsEnabled

`func (o *CreateStorageDomainParam) SetS3BucketsEnabled(v bool)`

SetS3BucketsEnabled sets S3BucketsEnabled field to given value.

### HasS3BucketsEnabled

`func (o *CreateStorageDomainParam) HasS3BucketsEnabled() bool`

HasS3BucketsEnabled returns a boolean if a field has been set.

### SetS3BucketsEnabledNil

`func (o *CreateStorageDomainParam) SetS3BucketsEnabledNil(b bool)`

 SetS3BucketsEnabledNil sets the value for S3BucketsEnabled to be an explicit nil

### UnsetS3BucketsEnabled
`func (o *CreateStorageDomainParam) UnsetS3BucketsEnabled()`

UnsetS3BucketsEnabled ensures that no value is present for S3BucketsEnabled, not even an explicit nil
### GetAdDomainName

`func (o *CreateStorageDomainParam) GetAdDomainName() string`

GetAdDomainName returns the AdDomainName field if non-nil, zero value otherwise.

### GetAdDomainNameOk

`func (o *CreateStorageDomainParam) GetAdDomainNameOk() (*string, bool)`

GetAdDomainNameOk returns a tuple with the AdDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdDomainName

`func (o *CreateStorageDomainParam) SetAdDomainName(v string)`

SetAdDomainName sets AdDomainName field to given value.

### HasAdDomainName

`func (o *CreateStorageDomainParam) HasAdDomainName() bool`

HasAdDomainName returns a boolean if a field has been set.

### SetAdDomainNameNil

`func (o *CreateStorageDomainParam) SetAdDomainNameNil(b bool)`

 SetAdDomainNameNil sets the value for AdDomainName to be an explicit nil

### UnsetAdDomainName
`func (o *CreateStorageDomainParam) UnsetAdDomainName()`

UnsetAdDomainName ensures that no value is present for AdDomainName, not even an explicit nil
### GetNisDomainNames

`func (o *CreateStorageDomainParam) GetNisDomainNames() []string`

GetNisDomainNames returns the NisDomainNames field if non-nil, zero value otherwise.

### GetNisDomainNamesOk

`func (o *CreateStorageDomainParam) GetNisDomainNamesOk() (*[]string, bool)`

GetNisDomainNamesOk returns a tuple with the NisDomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNisDomainNames

`func (o *CreateStorageDomainParam) SetNisDomainNames(v []string)`

SetNisDomainNames sets NisDomainNames field to given value.

### HasNisDomainNames

`func (o *CreateStorageDomainParam) HasNisDomainNames() bool`

HasNisDomainNames returns a boolean if a field has been set.

### SetNisDomainNamesNil

`func (o *CreateStorageDomainParam) SetNisDomainNamesNil(b bool)`

 SetNisDomainNamesNil sets the value for NisDomainNames to be an explicit nil

### UnsetNisDomainNames
`func (o *CreateStorageDomainParam) UnsetNisDomainNames()`

UnsetNisDomainNames ensures that no value is present for NisDomainNames, not even an explicit nil
### GetKerberosRealmName

`func (o *CreateStorageDomainParam) GetKerberosRealmName() string`

GetKerberosRealmName returns the KerberosRealmName field if non-nil, zero value otherwise.

### GetKerberosRealmNameOk

`func (o *CreateStorageDomainParam) GetKerberosRealmNameOk() (*string, bool)`

GetKerberosRealmNameOk returns a tuple with the KerberosRealmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosRealmName

`func (o *CreateStorageDomainParam) SetKerberosRealmName(v string)`

SetKerberosRealmName sets KerberosRealmName field to given value.

### HasKerberosRealmName

`func (o *CreateStorageDomainParam) HasKerberosRealmName() bool`

HasKerberosRealmName returns a boolean if a field has been set.

### SetKerberosRealmNameNil

`func (o *CreateStorageDomainParam) SetKerberosRealmNameNil(b bool)`

 SetKerberosRealmNameNil sets the value for KerberosRealmName to be an explicit nil

### UnsetKerberosRealmName
`func (o *CreateStorageDomainParam) UnsetKerberosRealmName()`

UnsetKerberosRealmName ensures that no value is present for KerberosRealmName, not even an explicit nil
### GetLdapProviderId

`func (o *CreateStorageDomainParam) GetLdapProviderId() int64`

GetLdapProviderId returns the LdapProviderId field if non-nil, zero value otherwise.

### GetLdapProviderIdOk

`func (o *CreateStorageDomainParam) GetLdapProviderIdOk() (*int64, bool)`

GetLdapProviderIdOk returns a tuple with the LdapProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapProviderId

`func (o *CreateStorageDomainParam) SetLdapProviderId(v int64)`

SetLdapProviderId sets LdapProviderId field to given value.

### HasLdapProviderId

`func (o *CreateStorageDomainParam) HasLdapProviderId() bool`

HasLdapProviderId returns a boolean if a field has been set.

### SetLdapProviderIdNil

`func (o *CreateStorageDomainParam) SetLdapProviderIdNil(b bool)`

 SetLdapProviderIdNil sets the value for LdapProviderId to be an explicit nil

### UnsetLdapProviderId
`func (o *CreateStorageDomainParam) UnsetLdapProviderId()`

UnsetLdapProviderId ensures that no value is present for LdapProviderId, not even an explicit nil
### GetTenantIds

`func (o *CreateStorageDomainParam) GetTenantIds() []string`

GetTenantIds returns the TenantIds field if non-nil, zero value otherwise.

### GetTenantIdsOk

`func (o *CreateStorageDomainParam) GetTenantIdsOk() (*[]string, bool)`

GetTenantIdsOk returns a tuple with the TenantIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIds

`func (o *CreateStorageDomainParam) SetTenantIds(v []string)`

SetTenantIds sets TenantIds field to given value.

### HasTenantIds

`func (o *CreateStorageDomainParam) HasTenantIds() bool`

HasTenantIds returns a boolean if a field has been set.

### SetTenantIdsNil

`func (o *CreateStorageDomainParam) SetTenantIdsNil(b bool)`

 SetTenantIdsNil sets the value for TenantIds to be an explicit nil

### UnsetTenantIds
`func (o *CreateStorageDomainParam) UnsetTenantIds()`

UnsetTenantIds ensures that no value is present for TenantIds, not even an explicit nil
### GetDirectArchiveEnabled

`func (o *CreateStorageDomainParam) GetDirectArchiveEnabled() bool`

GetDirectArchiveEnabled returns the DirectArchiveEnabled field if non-nil, zero value otherwise.

### GetDirectArchiveEnabledOk

`func (o *CreateStorageDomainParam) GetDirectArchiveEnabledOk() (*bool, bool)`

GetDirectArchiveEnabledOk returns a tuple with the DirectArchiveEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectArchiveEnabled

`func (o *CreateStorageDomainParam) SetDirectArchiveEnabled(v bool)`

SetDirectArchiveEnabled sets DirectArchiveEnabled field to given value.

### HasDirectArchiveEnabled

`func (o *CreateStorageDomainParam) HasDirectArchiveEnabled() bool`

HasDirectArchiveEnabled returns a boolean if a field has been set.

### SetDirectArchiveEnabledNil

`func (o *CreateStorageDomainParam) SetDirectArchiveEnabledNil(b bool)`

 SetDirectArchiveEnabledNil sets the value for DirectArchiveEnabled to be an explicit nil

### UnsetDirectArchiveEnabled
`func (o *CreateStorageDomainParam) UnsetDirectArchiveEnabled()`

UnsetDirectArchiveEnabled ensures that no value is present for DirectArchiveEnabled, not even an explicit nil
### GetBlobBrickSizeBytes

`func (o *CreateStorageDomainParam) GetBlobBrickSizeBytes() int32`

GetBlobBrickSizeBytes returns the BlobBrickSizeBytes field if non-nil, zero value otherwise.

### GetBlobBrickSizeBytesOk

`func (o *CreateStorageDomainParam) GetBlobBrickSizeBytesOk() (*int32, bool)`

GetBlobBrickSizeBytesOk returns a tuple with the BlobBrickSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobBrickSizeBytes

`func (o *CreateStorageDomainParam) SetBlobBrickSizeBytes(v int32)`

SetBlobBrickSizeBytes sets BlobBrickSizeBytes field to given value.

### HasBlobBrickSizeBytes

`func (o *CreateStorageDomainParam) HasBlobBrickSizeBytes() bool`

HasBlobBrickSizeBytes returns a boolean if a field has been set.

### SetBlobBrickSizeBytesNil

`func (o *CreateStorageDomainParam) SetBlobBrickSizeBytesNil(b bool)`

 SetBlobBrickSizeBytesNil sets the value for BlobBrickSizeBytes to be an explicit nil

### UnsetBlobBrickSizeBytes
`func (o *CreateStorageDomainParam) UnsetBlobBrickSizeBytes()`

UnsetBlobBrickSizeBytes ensures that no value is present for BlobBrickSizeBytes, not even an explicit nil
### GetKmsServerId

`func (o *CreateStorageDomainParam) GetKmsServerId() int64`

GetKmsServerId returns the KmsServerId field if non-nil, zero value otherwise.

### GetKmsServerIdOk

`func (o *CreateStorageDomainParam) GetKmsServerIdOk() (*int64, bool)`

GetKmsServerIdOk returns a tuple with the KmsServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsServerId

`func (o *CreateStorageDomainParam) SetKmsServerId(v int64)`

SetKmsServerId sets KmsServerId field to given value.

### HasKmsServerId

`func (o *CreateStorageDomainParam) HasKmsServerId() bool`

HasKmsServerId returns a boolean if a field has been set.

### SetKmsServerIdNil

`func (o *CreateStorageDomainParam) SetKmsServerIdNil(b bool)`

 SetKmsServerIdNil sets the value for KmsServerId to be an explicit nil

### UnsetKmsServerId
`func (o *CreateStorageDomainParam) UnsetKmsServerId()`

UnsetKmsServerId ensures that no value is present for KmsServerId, not even an explicit nil
### GetDekRotationEnabled

`func (o *CreateStorageDomainParam) GetDekRotationEnabled() bool`

GetDekRotationEnabled returns the DekRotationEnabled field if non-nil, zero value otherwise.

### GetDekRotationEnabledOk

`func (o *CreateStorageDomainParam) GetDekRotationEnabledOk() (*bool, bool)`

GetDekRotationEnabledOk returns a tuple with the DekRotationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDekRotationEnabled

`func (o *CreateStorageDomainParam) SetDekRotationEnabled(v bool)`

SetDekRotationEnabled sets DekRotationEnabled field to given value.

### HasDekRotationEnabled

`func (o *CreateStorageDomainParam) HasDekRotationEnabled() bool`

HasDekRotationEnabled returns a boolean if a field has been set.

### SetDekRotationEnabledNil

`func (o *CreateStorageDomainParam) SetDekRotationEnabledNil(b bool)`

 SetDekRotationEnabledNil sets the value for DekRotationEnabled to be an explicit nil

### UnsetDekRotationEnabled
`func (o *CreateStorageDomainParam) UnsetDekRotationEnabled()`

UnsetDekRotationEnabled ensures that no value is present for DekRotationEnabled, not even an explicit nil
### GetTreatFileSyncAsDataSync

`func (o *CreateStorageDomainParam) GetTreatFileSyncAsDataSync() bool`

GetTreatFileSyncAsDataSync returns the TreatFileSyncAsDataSync field if non-nil, zero value otherwise.

### GetTreatFileSyncAsDataSyncOk

`func (o *CreateStorageDomainParam) GetTreatFileSyncAsDataSyncOk() (*bool, bool)`

GetTreatFileSyncAsDataSyncOk returns a tuple with the TreatFileSyncAsDataSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreatFileSyncAsDataSync

`func (o *CreateStorageDomainParam) SetTreatFileSyncAsDataSync(v bool)`

SetTreatFileSyncAsDataSync sets TreatFileSyncAsDataSync field to given value.

### HasTreatFileSyncAsDataSync

`func (o *CreateStorageDomainParam) HasTreatFileSyncAsDataSync() bool`

HasTreatFileSyncAsDataSync returns a boolean if a field has been set.

### SetTreatFileSyncAsDataSyncNil

`func (o *CreateStorageDomainParam) SetTreatFileSyncAsDataSyncNil(b bool)`

 SetTreatFileSyncAsDataSyncNil sets the value for TreatFileSyncAsDataSync to be an explicit nil

### UnsetTreatFileSyncAsDataSync
`func (o *CreateStorageDomainParam) UnsetTreatFileSyncAsDataSync()`

UnsetTreatFileSyncAsDataSync ensures that no value is present for TreatFileSyncAsDataSync, not even an explicit nil
### GetRecommended

`func (o *CreateStorageDomainParam) GetRecommended() bool`

GetRecommended returns the Recommended field if non-nil, zero value otherwise.

### GetRecommendedOk

`func (o *CreateStorageDomainParam) GetRecommendedOk() (*bool, bool)`

GetRecommendedOk returns a tuple with the Recommended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommended

`func (o *CreateStorageDomainParam) SetRecommended(v bool)`

SetRecommended sets Recommended field to given value.

### HasRecommended

`func (o *CreateStorageDomainParam) HasRecommended() bool`

HasRecommended returns a boolean if a field has been set.

### SetRecommendedNil

`func (o *CreateStorageDomainParam) SetRecommendedNil(b bool)`

 SetRecommendedNil sets the value for Recommended to be an explicit nil

### UnsetRecommended
`func (o *CreateStorageDomainParam) UnsetRecommended()`

UnsetRecommended ensures that no value is present for Recommended, not even an explicit nil
### GetStats

`func (o *CreateStorageDomainParam) GetStats() DataUsageStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *CreateStorageDomainParam) GetStatsOk() (*DataUsageStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *CreateStorageDomainParam) SetStats(v DataUsageStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *CreateStorageDomainParam) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetSchemas

`func (o *CreateStorageDomainParam) GetSchemas() []Schema`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *CreateStorageDomainParam) GetSchemasOk() (*[]Schema, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *CreateStorageDomainParam) SetSchemas(v []Schema)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *CreateStorageDomainParam) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### SetSchemasNil

`func (o *CreateStorageDomainParam) SetSchemasNil(b bool)`

 SetSchemasNil sets the value for Schemas to be an explicit nil

### UnsetSchemas
`func (o *CreateStorageDomainParam) UnsetSchemas()`

UnsetSchemas ensures that no value is present for Schemas, not even an explicit nil
### GetFileCountBySize

`func (o *CreateStorageDomainParam) GetFileCountBySize() []FileCount`

GetFileCountBySize returns the FileCountBySize field if non-nil, zero value otherwise.

### GetFileCountBySizeOk

`func (o *CreateStorageDomainParam) GetFileCountBySizeOk() (*[]FileCount, bool)`

GetFileCountBySizeOk returns a tuple with the FileCountBySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileCountBySize

`func (o *CreateStorageDomainParam) SetFileCountBySize(v []FileCount)`

SetFileCountBySize sets FileCountBySize field to given value.

### HasFileCountBySize

`func (o *CreateStorageDomainParam) HasFileCountBySize() bool`

HasFileCountBySize returns a boolean if a field has been set.

### SetFileCountBySizeNil

`func (o *CreateStorageDomainParam) SetFileCountBySizeNil(b bool)`

 SetFileCountBySizeNil sets the value for FileCountBySize to be an explicit nil

### UnsetFileCountBySize
`func (o *CreateStorageDomainParam) UnsetFileCountBySize()`

UnsetFileCountBySize ensures that no value is present for FileCountBySize, not even an explicit nil
### GetCloudDomainId

`func (o *CreateStorageDomainParam) GetCloudDomainId() int64`

GetCloudDomainId returns the CloudDomainId field if non-nil, zero value otherwise.

### GetCloudDomainIdOk

`func (o *CreateStorageDomainParam) GetCloudDomainIdOk() (*int64, bool)`

GetCloudDomainIdOk returns a tuple with the CloudDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudDomainId

`func (o *CreateStorageDomainParam) SetCloudDomainId(v int64)`

SetCloudDomainId sets CloudDomainId field to given value.

### HasCloudDomainId

`func (o *CreateStorageDomainParam) HasCloudDomainId() bool`

HasCloudDomainId returns a boolean if a field has been set.

### SetCloudDomainIdNil

`func (o *CreateStorageDomainParam) SetCloudDomainIdNil(b bool)`

 SetCloudDomainIdNil sets the value for CloudDomainId to be an explicit nil

### UnsetCloudDomainId
`func (o *CreateStorageDomainParam) UnsetCloudDomainId()`

UnsetCloudDomainId ensures that no value is present for CloudDomainId, not even an explicit nil
### GetVaultId

`func (o *CreateStorageDomainParam) GetVaultId() int64`

GetVaultId returns the VaultId field if non-nil, zero value otherwise.

### GetVaultIdOk

`func (o *CreateStorageDomainParam) GetVaultIdOk() (*int64, bool)`

GetVaultIdOk returns a tuple with the VaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultId

`func (o *CreateStorageDomainParam) SetVaultId(v int64)`

SetVaultId sets VaultId field to given value.

### HasVaultId

`func (o *CreateStorageDomainParam) HasVaultId() bool`

HasVaultId returns a boolean if a field has been set.

### SetVaultIdNil

`func (o *CreateStorageDomainParam) SetVaultIdNil(b bool)`

 SetVaultIdNil sets the value for VaultId to be an explicit nil

### UnsetVaultId
`func (o *CreateStorageDomainParam) UnsetVaultId()`

UnsetVaultId ensures that no value is present for VaultId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


