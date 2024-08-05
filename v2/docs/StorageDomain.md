# StorageDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdDomainName** | Pointer to **NullableString** | Specifies the Active Directory domain name that this Storage Domain is mapped to. | [optional] 
**BlobBrickSizeBytes** | Pointer to **NullableInt32** | Specifies the brick size used for blobs in this Storage Domain. | [optional] 
**CloudDomainId** | Pointer to **NullableInt64** | Specifies the cloud domain Id. | [optional] 
**CloudDownWaterFallParams** | Pointer to [**StorageDomainCloudDownWaterFallParams**](StorageDomainCloudDownWaterFallParams.md) |  | [optional] 
**ClusterPartitionId** | **NullableInt64** | Specifies the cluster partition id of the Storage Domain. | 
**ClusterPartitionName** | Pointer to **NullableString** | Specifies the cluster partition name of the Storage Domain. | [optional] [readonly] 
**DefaultUserQuota** | Pointer to [**StorageDomainDefaultUserQuota**](StorageDomainDefaultUserQuota.md) |  | [optional] 
**DefaultViewQuota** | Pointer to [**StorageDomainDefaultViewQuota**](StorageDomainDefaultViewQuota.md) |  | [optional] 
**DekRotationEnabled** | Pointer to **NullableBool** | Specifies whether DEK(Data Encryption Key) rotation is enabled for this Storage Domain. This is applicable only when the Storage Domain uses AWS or similar KMS in which the KEK (Key Encryption Key) is not created and maintained by Cohesity. For Internal KMS and keys stored in Safenet servers, DEK rotation will not be performed. | [optional] 
**DirectArchiveEnabled** | Pointer to **NullableBool** | Specifies whether to enable driect archive on this Storage Domain. If enabled, this Storage Domain can be used as a staging area while copying a large dataset that can&#39;t fit on the cluster to an external target. | [optional] 
**FileCountBySize** | Pointer to [**[]FileCount**](FileCount.md) | Specifies the file count by size for the View. | [optional] [readonly] 
**Id** | Pointer to **NullableInt64** | Specifies the Storage Domain id. | [optional] [readonly] 
**KerberosRealmName** | Pointer to **NullableString** | Specifies the Kerberos realm name that this Storage Domain is mapped to. | [optional] 
**KmsServerId** | Pointer to **NullableInt64** | Specifies the associated KMS server id. | [optional] 
**LastKeyRotationTimestampMsecs** | Pointer to **NullableInt64** | Last key rotation timestamp in msecs for storage domain. | [optional] 
**LdapProviderId** | Pointer to **NullableInt64** | Specifies the LDAP provider id that this Storage Domain is mapped to. | [optional] 
**Name** | **NullableString** | Specifies the Storage Domain name. | 
**NisDomainNames** | Pointer to **[]string** | Specifies the NIS domain names that this Storage Domain is mapped to. | [optional] 
**PhysicalQuota** | Pointer to [**StorageDomainPhysicalQuota**](StorageDomainPhysicalQuota.md) |  | [optional] 
**Recommended** | Pointer to **NullableBool** | Specifies whether Storage Domain is recommended for the specified View template. | [optional] [readonly] 
**RemovalState** | Pointer to **NullableString** | Specifies the current removal state of the Storage Domain. &#39;DontRemove&#39; means the state of object is functional and it is not being removed. &#39;MarkedForRemoval&#39; means the object is being removed. &#39;OkToRemove&#39; means the object has been removed on the Cohesity Cluster and if the object is physical, it can be removed from the Cohesity Cluster. | [optional] [readonly] 
**S3BucketsEnabled** | Pointer to **NullableBool** | Specifies whether to enable creation of S3 bucket on this Storage Domain. | [optional] 
**Schemas** | Pointer to [**[]Schema**](Schema.md) | Specifies the Storage Domain schemas. | [optional] [readonly] 
**Stats** | Pointer to [**StorageDomainStats**](StorageDomainStats.md) |  | [optional] 
**StoragePolicy** | Pointer to [**StorageDomainStoragePolicy**](StorageDomainStoragePolicy.md) |  | [optional] 
**SubnetWhitelist** | Pointer to [**[]Subnet**](Subnet.md) | Specifies a list of Subnets with IP addresses that have permissions to access the Storage Domain. | [optional] 
**TenantIds** | Pointer to **[]string** | Specifies a list of tenant ids that that Storage Domain belongs. There can only be one tenant id in this field unless Storage Domain sharing between tenants is allowed on this cluster. | [optional] 
**TreatFileSyncAsDataSync** | Pointer to **NullableBool** | If &#39;true&#39;, when the Cohesity Cluster is writing to a file, the file modification time is not persisted synchronously during the file write, so the modification time may not be accurate. (Typically the file modification time is off by 30 seconds but it can be longer.) | [optional] 
**VaultId** | Pointer to **NullableInt64** | Specifies the vault Id associated with cloud domain ID. | [optional] 

## Methods

### NewStorageDomain

`func NewStorageDomain(clusterPartitionId NullableInt64, name NullableString, ) *StorageDomain`

NewStorageDomain instantiates a new StorageDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageDomainWithDefaults

`func NewStorageDomainWithDefaults() *StorageDomain`

NewStorageDomainWithDefaults instantiates a new StorageDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdDomainName

`func (o *StorageDomain) GetAdDomainName() string`

GetAdDomainName returns the AdDomainName field if non-nil, zero value otherwise.

### GetAdDomainNameOk

`func (o *StorageDomain) GetAdDomainNameOk() (*string, bool)`

GetAdDomainNameOk returns a tuple with the AdDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdDomainName

`func (o *StorageDomain) SetAdDomainName(v string)`

SetAdDomainName sets AdDomainName field to given value.

### HasAdDomainName

`func (o *StorageDomain) HasAdDomainName() bool`

HasAdDomainName returns a boolean if a field has been set.

### SetAdDomainNameNil

`func (o *StorageDomain) SetAdDomainNameNil(b bool)`

 SetAdDomainNameNil sets the value for AdDomainName to be an explicit nil

### UnsetAdDomainName
`func (o *StorageDomain) UnsetAdDomainName()`

UnsetAdDomainName ensures that no value is present for AdDomainName, not even an explicit nil
### GetBlobBrickSizeBytes

`func (o *StorageDomain) GetBlobBrickSizeBytes() int32`

GetBlobBrickSizeBytes returns the BlobBrickSizeBytes field if non-nil, zero value otherwise.

### GetBlobBrickSizeBytesOk

`func (o *StorageDomain) GetBlobBrickSizeBytesOk() (*int32, bool)`

GetBlobBrickSizeBytesOk returns a tuple with the BlobBrickSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobBrickSizeBytes

`func (o *StorageDomain) SetBlobBrickSizeBytes(v int32)`

SetBlobBrickSizeBytes sets BlobBrickSizeBytes field to given value.

### HasBlobBrickSizeBytes

`func (o *StorageDomain) HasBlobBrickSizeBytes() bool`

HasBlobBrickSizeBytes returns a boolean if a field has been set.

### SetBlobBrickSizeBytesNil

`func (o *StorageDomain) SetBlobBrickSizeBytesNil(b bool)`

 SetBlobBrickSizeBytesNil sets the value for BlobBrickSizeBytes to be an explicit nil

### UnsetBlobBrickSizeBytes
`func (o *StorageDomain) UnsetBlobBrickSizeBytes()`

UnsetBlobBrickSizeBytes ensures that no value is present for BlobBrickSizeBytes, not even an explicit nil
### GetCloudDomainId

`func (o *StorageDomain) GetCloudDomainId() int64`

GetCloudDomainId returns the CloudDomainId field if non-nil, zero value otherwise.

### GetCloudDomainIdOk

`func (o *StorageDomain) GetCloudDomainIdOk() (*int64, bool)`

GetCloudDomainIdOk returns a tuple with the CloudDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudDomainId

`func (o *StorageDomain) SetCloudDomainId(v int64)`

SetCloudDomainId sets CloudDomainId field to given value.

### HasCloudDomainId

`func (o *StorageDomain) HasCloudDomainId() bool`

HasCloudDomainId returns a boolean if a field has been set.

### SetCloudDomainIdNil

`func (o *StorageDomain) SetCloudDomainIdNil(b bool)`

 SetCloudDomainIdNil sets the value for CloudDomainId to be an explicit nil

### UnsetCloudDomainId
`func (o *StorageDomain) UnsetCloudDomainId()`

UnsetCloudDomainId ensures that no value is present for CloudDomainId, not even an explicit nil
### GetCloudDownWaterFallParams

`func (o *StorageDomain) GetCloudDownWaterFallParams() StorageDomainCloudDownWaterFallParams`

GetCloudDownWaterFallParams returns the CloudDownWaterFallParams field if non-nil, zero value otherwise.

### GetCloudDownWaterFallParamsOk

`func (o *StorageDomain) GetCloudDownWaterFallParamsOk() (*StorageDomainCloudDownWaterFallParams, bool)`

GetCloudDownWaterFallParamsOk returns a tuple with the CloudDownWaterFallParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudDownWaterFallParams

`func (o *StorageDomain) SetCloudDownWaterFallParams(v StorageDomainCloudDownWaterFallParams)`

SetCloudDownWaterFallParams sets CloudDownWaterFallParams field to given value.

### HasCloudDownWaterFallParams

`func (o *StorageDomain) HasCloudDownWaterFallParams() bool`

HasCloudDownWaterFallParams returns a boolean if a field has been set.

### GetClusterPartitionId

`func (o *StorageDomain) GetClusterPartitionId() int64`

GetClusterPartitionId returns the ClusterPartitionId field if non-nil, zero value otherwise.

### GetClusterPartitionIdOk

`func (o *StorageDomain) GetClusterPartitionIdOk() (*int64, bool)`

GetClusterPartitionIdOk returns a tuple with the ClusterPartitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterPartitionId

`func (o *StorageDomain) SetClusterPartitionId(v int64)`

SetClusterPartitionId sets ClusterPartitionId field to given value.


### SetClusterPartitionIdNil

`func (o *StorageDomain) SetClusterPartitionIdNil(b bool)`

 SetClusterPartitionIdNil sets the value for ClusterPartitionId to be an explicit nil

### UnsetClusterPartitionId
`func (o *StorageDomain) UnsetClusterPartitionId()`

UnsetClusterPartitionId ensures that no value is present for ClusterPartitionId, not even an explicit nil
### GetClusterPartitionName

`func (o *StorageDomain) GetClusterPartitionName() string`

GetClusterPartitionName returns the ClusterPartitionName field if non-nil, zero value otherwise.

### GetClusterPartitionNameOk

`func (o *StorageDomain) GetClusterPartitionNameOk() (*string, bool)`

GetClusterPartitionNameOk returns a tuple with the ClusterPartitionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterPartitionName

`func (o *StorageDomain) SetClusterPartitionName(v string)`

SetClusterPartitionName sets ClusterPartitionName field to given value.

### HasClusterPartitionName

`func (o *StorageDomain) HasClusterPartitionName() bool`

HasClusterPartitionName returns a boolean if a field has been set.

### SetClusterPartitionNameNil

`func (o *StorageDomain) SetClusterPartitionNameNil(b bool)`

 SetClusterPartitionNameNil sets the value for ClusterPartitionName to be an explicit nil

### UnsetClusterPartitionName
`func (o *StorageDomain) UnsetClusterPartitionName()`

UnsetClusterPartitionName ensures that no value is present for ClusterPartitionName, not even an explicit nil
### GetDefaultUserQuota

`func (o *StorageDomain) GetDefaultUserQuota() StorageDomainDefaultUserQuota`

GetDefaultUserQuota returns the DefaultUserQuota field if non-nil, zero value otherwise.

### GetDefaultUserQuotaOk

`func (o *StorageDomain) GetDefaultUserQuotaOk() (*StorageDomainDefaultUserQuota, bool)`

GetDefaultUserQuotaOk returns a tuple with the DefaultUserQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUserQuota

`func (o *StorageDomain) SetDefaultUserQuota(v StorageDomainDefaultUserQuota)`

SetDefaultUserQuota sets DefaultUserQuota field to given value.

### HasDefaultUserQuota

`func (o *StorageDomain) HasDefaultUserQuota() bool`

HasDefaultUserQuota returns a boolean if a field has been set.

### GetDefaultViewQuota

`func (o *StorageDomain) GetDefaultViewQuota() StorageDomainDefaultViewQuota`

GetDefaultViewQuota returns the DefaultViewQuota field if non-nil, zero value otherwise.

### GetDefaultViewQuotaOk

`func (o *StorageDomain) GetDefaultViewQuotaOk() (*StorageDomainDefaultViewQuota, bool)`

GetDefaultViewQuotaOk returns a tuple with the DefaultViewQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultViewQuota

`func (o *StorageDomain) SetDefaultViewQuota(v StorageDomainDefaultViewQuota)`

SetDefaultViewQuota sets DefaultViewQuota field to given value.

### HasDefaultViewQuota

`func (o *StorageDomain) HasDefaultViewQuota() bool`

HasDefaultViewQuota returns a boolean if a field has been set.

### GetDekRotationEnabled

`func (o *StorageDomain) GetDekRotationEnabled() bool`

GetDekRotationEnabled returns the DekRotationEnabled field if non-nil, zero value otherwise.

### GetDekRotationEnabledOk

`func (o *StorageDomain) GetDekRotationEnabledOk() (*bool, bool)`

GetDekRotationEnabledOk returns a tuple with the DekRotationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDekRotationEnabled

`func (o *StorageDomain) SetDekRotationEnabled(v bool)`

SetDekRotationEnabled sets DekRotationEnabled field to given value.

### HasDekRotationEnabled

`func (o *StorageDomain) HasDekRotationEnabled() bool`

HasDekRotationEnabled returns a boolean if a field has been set.

### SetDekRotationEnabledNil

`func (o *StorageDomain) SetDekRotationEnabledNil(b bool)`

 SetDekRotationEnabledNil sets the value for DekRotationEnabled to be an explicit nil

### UnsetDekRotationEnabled
`func (o *StorageDomain) UnsetDekRotationEnabled()`

UnsetDekRotationEnabled ensures that no value is present for DekRotationEnabled, not even an explicit nil
### GetDirectArchiveEnabled

`func (o *StorageDomain) GetDirectArchiveEnabled() bool`

GetDirectArchiveEnabled returns the DirectArchiveEnabled field if non-nil, zero value otherwise.

### GetDirectArchiveEnabledOk

`func (o *StorageDomain) GetDirectArchiveEnabledOk() (*bool, bool)`

GetDirectArchiveEnabledOk returns a tuple with the DirectArchiveEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectArchiveEnabled

`func (o *StorageDomain) SetDirectArchiveEnabled(v bool)`

SetDirectArchiveEnabled sets DirectArchiveEnabled field to given value.

### HasDirectArchiveEnabled

`func (o *StorageDomain) HasDirectArchiveEnabled() bool`

HasDirectArchiveEnabled returns a boolean if a field has been set.

### SetDirectArchiveEnabledNil

`func (o *StorageDomain) SetDirectArchiveEnabledNil(b bool)`

 SetDirectArchiveEnabledNil sets the value for DirectArchiveEnabled to be an explicit nil

### UnsetDirectArchiveEnabled
`func (o *StorageDomain) UnsetDirectArchiveEnabled()`

UnsetDirectArchiveEnabled ensures that no value is present for DirectArchiveEnabled, not even an explicit nil
### GetFileCountBySize

`func (o *StorageDomain) GetFileCountBySize() []FileCount`

GetFileCountBySize returns the FileCountBySize field if non-nil, zero value otherwise.

### GetFileCountBySizeOk

`func (o *StorageDomain) GetFileCountBySizeOk() (*[]FileCount, bool)`

GetFileCountBySizeOk returns a tuple with the FileCountBySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileCountBySize

`func (o *StorageDomain) SetFileCountBySize(v []FileCount)`

SetFileCountBySize sets FileCountBySize field to given value.

### HasFileCountBySize

`func (o *StorageDomain) HasFileCountBySize() bool`

HasFileCountBySize returns a boolean if a field has been set.

### SetFileCountBySizeNil

`func (o *StorageDomain) SetFileCountBySizeNil(b bool)`

 SetFileCountBySizeNil sets the value for FileCountBySize to be an explicit nil

### UnsetFileCountBySize
`func (o *StorageDomain) UnsetFileCountBySize()`

UnsetFileCountBySize ensures that no value is present for FileCountBySize, not even an explicit nil
### GetId

`func (o *StorageDomain) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StorageDomain) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StorageDomain) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *StorageDomain) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *StorageDomain) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *StorageDomain) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetKerberosRealmName

`func (o *StorageDomain) GetKerberosRealmName() string`

GetKerberosRealmName returns the KerberosRealmName field if non-nil, zero value otherwise.

### GetKerberosRealmNameOk

`func (o *StorageDomain) GetKerberosRealmNameOk() (*string, bool)`

GetKerberosRealmNameOk returns a tuple with the KerberosRealmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosRealmName

`func (o *StorageDomain) SetKerberosRealmName(v string)`

SetKerberosRealmName sets KerberosRealmName field to given value.

### HasKerberosRealmName

`func (o *StorageDomain) HasKerberosRealmName() bool`

HasKerberosRealmName returns a boolean if a field has been set.

### SetKerberosRealmNameNil

`func (o *StorageDomain) SetKerberosRealmNameNil(b bool)`

 SetKerberosRealmNameNil sets the value for KerberosRealmName to be an explicit nil

### UnsetKerberosRealmName
`func (o *StorageDomain) UnsetKerberosRealmName()`

UnsetKerberosRealmName ensures that no value is present for KerberosRealmName, not even an explicit nil
### GetKmsServerId

`func (o *StorageDomain) GetKmsServerId() int64`

GetKmsServerId returns the KmsServerId field if non-nil, zero value otherwise.

### GetKmsServerIdOk

`func (o *StorageDomain) GetKmsServerIdOk() (*int64, bool)`

GetKmsServerIdOk returns a tuple with the KmsServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsServerId

`func (o *StorageDomain) SetKmsServerId(v int64)`

SetKmsServerId sets KmsServerId field to given value.

### HasKmsServerId

`func (o *StorageDomain) HasKmsServerId() bool`

HasKmsServerId returns a boolean if a field has been set.

### SetKmsServerIdNil

`func (o *StorageDomain) SetKmsServerIdNil(b bool)`

 SetKmsServerIdNil sets the value for KmsServerId to be an explicit nil

### UnsetKmsServerId
`func (o *StorageDomain) UnsetKmsServerId()`

UnsetKmsServerId ensures that no value is present for KmsServerId, not even an explicit nil
### GetLastKeyRotationTimestampMsecs

`func (o *StorageDomain) GetLastKeyRotationTimestampMsecs() int64`

GetLastKeyRotationTimestampMsecs returns the LastKeyRotationTimestampMsecs field if non-nil, zero value otherwise.

### GetLastKeyRotationTimestampMsecsOk

`func (o *StorageDomain) GetLastKeyRotationTimestampMsecsOk() (*int64, bool)`

GetLastKeyRotationTimestampMsecsOk returns a tuple with the LastKeyRotationTimestampMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastKeyRotationTimestampMsecs

`func (o *StorageDomain) SetLastKeyRotationTimestampMsecs(v int64)`

SetLastKeyRotationTimestampMsecs sets LastKeyRotationTimestampMsecs field to given value.

### HasLastKeyRotationTimestampMsecs

`func (o *StorageDomain) HasLastKeyRotationTimestampMsecs() bool`

HasLastKeyRotationTimestampMsecs returns a boolean if a field has been set.

### SetLastKeyRotationTimestampMsecsNil

`func (o *StorageDomain) SetLastKeyRotationTimestampMsecsNil(b bool)`

 SetLastKeyRotationTimestampMsecsNil sets the value for LastKeyRotationTimestampMsecs to be an explicit nil

### UnsetLastKeyRotationTimestampMsecs
`func (o *StorageDomain) UnsetLastKeyRotationTimestampMsecs()`

UnsetLastKeyRotationTimestampMsecs ensures that no value is present for LastKeyRotationTimestampMsecs, not even an explicit nil
### GetLdapProviderId

`func (o *StorageDomain) GetLdapProviderId() int64`

GetLdapProviderId returns the LdapProviderId field if non-nil, zero value otherwise.

### GetLdapProviderIdOk

`func (o *StorageDomain) GetLdapProviderIdOk() (*int64, bool)`

GetLdapProviderIdOk returns a tuple with the LdapProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapProviderId

`func (o *StorageDomain) SetLdapProviderId(v int64)`

SetLdapProviderId sets LdapProviderId field to given value.

### HasLdapProviderId

`func (o *StorageDomain) HasLdapProviderId() bool`

HasLdapProviderId returns a boolean if a field has been set.

### SetLdapProviderIdNil

`func (o *StorageDomain) SetLdapProviderIdNil(b bool)`

 SetLdapProviderIdNil sets the value for LdapProviderId to be an explicit nil

### UnsetLdapProviderId
`func (o *StorageDomain) UnsetLdapProviderId()`

UnsetLdapProviderId ensures that no value is present for LdapProviderId, not even an explicit nil
### GetName

`func (o *StorageDomain) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageDomain) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageDomain) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *StorageDomain) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *StorageDomain) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNisDomainNames

`func (o *StorageDomain) GetNisDomainNames() []string`

GetNisDomainNames returns the NisDomainNames field if non-nil, zero value otherwise.

### GetNisDomainNamesOk

`func (o *StorageDomain) GetNisDomainNamesOk() (*[]string, bool)`

GetNisDomainNamesOk returns a tuple with the NisDomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNisDomainNames

`func (o *StorageDomain) SetNisDomainNames(v []string)`

SetNisDomainNames sets NisDomainNames field to given value.

### HasNisDomainNames

`func (o *StorageDomain) HasNisDomainNames() bool`

HasNisDomainNames returns a boolean if a field has been set.

### SetNisDomainNamesNil

`func (o *StorageDomain) SetNisDomainNamesNil(b bool)`

 SetNisDomainNamesNil sets the value for NisDomainNames to be an explicit nil

### UnsetNisDomainNames
`func (o *StorageDomain) UnsetNisDomainNames()`

UnsetNisDomainNames ensures that no value is present for NisDomainNames, not even an explicit nil
### GetPhysicalQuota

`func (o *StorageDomain) GetPhysicalQuota() StorageDomainPhysicalQuota`

GetPhysicalQuota returns the PhysicalQuota field if non-nil, zero value otherwise.

### GetPhysicalQuotaOk

`func (o *StorageDomain) GetPhysicalQuotaOk() (*StorageDomainPhysicalQuota, bool)`

GetPhysicalQuotaOk returns a tuple with the PhysicalQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalQuota

`func (o *StorageDomain) SetPhysicalQuota(v StorageDomainPhysicalQuota)`

SetPhysicalQuota sets PhysicalQuota field to given value.

### HasPhysicalQuota

`func (o *StorageDomain) HasPhysicalQuota() bool`

HasPhysicalQuota returns a boolean if a field has been set.

### GetRecommended

`func (o *StorageDomain) GetRecommended() bool`

GetRecommended returns the Recommended field if non-nil, zero value otherwise.

### GetRecommendedOk

`func (o *StorageDomain) GetRecommendedOk() (*bool, bool)`

GetRecommendedOk returns a tuple with the Recommended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommended

`func (o *StorageDomain) SetRecommended(v bool)`

SetRecommended sets Recommended field to given value.

### HasRecommended

`func (o *StorageDomain) HasRecommended() bool`

HasRecommended returns a boolean if a field has been set.

### SetRecommendedNil

`func (o *StorageDomain) SetRecommendedNil(b bool)`

 SetRecommendedNil sets the value for Recommended to be an explicit nil

### UnsetRecommended
`func (o *StorageDomain) UnsetRecommended()`

UnsetRecommended ensures that no value is present for Recommended, not even an explicit nil
### GetRemovalState

`func (o *StorageDomain) GetRemovalState() string`

GetRemovalState returns the RemovalState field if non-nil, zero value otherwise.

### GetRemovalStateOk

`func (o *StorageDomain) GetRemovalStateOk() (*string, bool)`

GetRemovalStateOk returns a tuple with the RemovalState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalState

`func (o *StorageDomain) SetRemovalState(v string)`

SetRemovalState sets RemovalState field to given value.

### HasRemovalState

`func (o *StorageDomain) HasRemovalState() bool`

HasRemovalState returns a boolean if a field has been set.

### SetRemovalStateNil

`func (o *StorageDomain) SetRemovalStateNil(b bool)`

 SetRemovalStateNil sets the value for RemovalState to be an explicit nil

### UnsetRemovalState
`func (o *StorageDomain) UnsetRemovalState()`

UnsetRemovalState ensures that no value is present for RemovalState, not even an explicit nil
### GetS3BucketsEnabled

`func (o *StorageDomain) GetS3BucketsEnabled() bool`

GetS3BucketsEnabled returns the S3BucketsEnabled field if non-nil, zero value otherwise.

### GetS3BucketsEnabledOk

`func (o *StorageDomain) GetS3BucketsEnabledOk() (*bool, bool)`

GetS3BucketsEnabledOk returns a tuple with the S3BucketsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3BucketsEnabled

`func (o *StorageDomain) SetS3BucketsEnabled(v bool)`

SetS3BucketsEnabled sets S3BucketsEnabled field to given value.

### HasS3BucketsEnabled

`func (o *StorageDomain) HasS3BucketsEnabled() bool`

HasS3BucketsEnabled returns a boolean if a field has been set.

### SetS3BucketsEnabledNil

`func (o *StorageDomain) SetS3BucketsEnabledNil(b bool)`

 SetS3BucketsEnabledNil sets the value for S3BucketsEnabled to be an explicit nil

### UnsetS3BucketsEnabled
`func (o *StorageDomain) UnsetS3BucketsEnabled()`

UnsetS3BucketsEnabled ensures that no value is present for S3BucketsEnabled, not even an explicit nil
### GetSchemas

`func (o *StorageDomain) GetSchemas() []Schema`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *StorageDomain) GetSchemasOk() (*[]Schema, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *StorageDomain) SetSchemas(v []Schema)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *StorageDomain) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### SetSchemasNil

`func (o *StorageDomain) SetSchemasNil(b bool)`

 SetSchemasNil sets the value for Schemas to be an explicit nil

### UnsetSchemas
`func (o *StorageDomain) UnsetSchemas()`

UnsetSchemas ensures that no value is present for Schemas, not even an explicit nil
### GetStats

`func (o *StorageDomain) GetStats() StorageDomainStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *StorageDomain) GetStatsOk() (*StorageDomainStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *StorageDomain) SetStats(v StorageDomainStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *StorageDomain) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetStoragePolicy

`func (o *StorageDomain) GetStoragePolicy() StorageDomainStoragePolicy`

GetStoragePolicy returns the StoragePolicy field if non-nil, zero value otherwise.

### GetStoragePolicyOk

`func (o *StorageDomain) GetStoragePolicyOk() (*StorageDomainStoragePolicy, bool)`

GetStoragePolicyOk returns a tuple with the StoragePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePolicy

`func (o *StorageDomain) SetStoragePolicy(v StorageDomainStoragePolicy)`

SetStoragePolicy sets StoragePolicy field to given value.

### HasStoragePolicy

`func (o *StorageDomain) HasStoragePolicy() bool`

HasStoragePolicy returns a boolean if a field has been set.

### GetSubnetWhitelist

`func (o *StorageDomain) GetSubnetWhitelist() []Subnet`

GetSubnetWhitelist returns the SubnetWhitelist field if non-nil, zero value otherwise.

### GetSubnetWhitelistOk

`func (o *StorageDomain) GetSubnetWhitelistOk() (*[]Subnet, bool)`

GetSubnetWhitelistOk returns a tuple with the SubnetWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetWhitelist

`func (o *StorageDomain) SetSubnetWhitelist(v []Subnet)`

SetSubnetWhitelist sets SubnetWhitelist field to given value.

### HasSubnetWhitelist

`func (o *StorageDomain) HasSubnetWhitelist() bool`

HasSubnetWhitelist returns a boolean if a field has been set.

### SetSubnetWhitelistNil

`func (o *StorageDomain) SetSubnetWhitelistNil(b bool)`

 SetSubnetWhitelistNil sets the value for SubnetWhitelist to be an explicit nil

### UnsetSubnetWhitelist
`func (o *StorageDomain) UnsetSubnetWhitelist()`

UnsetSubnetWhitelist ensures that no value is present for SubnetWhitelist, not even an explicit nil
### GetTenantIds

`func (o *StorageDomain) GetTenantIds() []string`

GetTenantIds returns the TenantIds field if non-nil, zero value otherwise.

### GetTenantIdsOk

`func (o *StorageDomain) GetTenantIdsOk() (*[]string, bool)`

GetTenantIdsOk returns a tuple with the TenantIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIds

`func (o *StorageDomain) SetTenantIds(v []string)`

SetTenantIds sets TenantIds field to given value.

### HasTenantIds

`func (o *StorageDomain) HasTenantIds() bool`

HasTenantIds returns a boolean if a field has been set.

### SetTenantIdsNil

`func (o *StorageDomain) SetTenantIdsNil(b bool)`

 SetTenantIdsNil sets the value for TenantIds to be an explicit nil

### UnsetTenantIds
`func (o *StorageDomain) UnsetTenantIds()`

UnsetTenantIds ensures that no value is present for TenantIds, not even an explicit nil
### GetTreatFileSyncAsDataSync

`func (o *StorageDomain) GetTreatFileSyncAsDataSync() bool`

GetTreatFileSyncAsDataSync returns the TreatFileSyncAsDataSync field if non-nil, zero value otherwise.

### GetTreatFileSyncAsDataSyncOk

`func (o *StorageDomain) GetTreatFileSyncAsDataSyncOk() (*bool, bool)`

GetTreatFileSyncAsDataSyncOk returns a tuple with the TreatFileSyncAsDataSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreatFileSyncAsDataSync

`func (o *StorageDomain) SetTreatFileSyncAsDataSync(v bool)`

SetTreatFileSyncAsDataSync sets TreatFileSyncAsDataSync field to given value.

### HasTreatFileSyncAsDataSync

`func (o *StorageDomain) HasTreatFileSyncAsDataSync() bool`

HasTreatFileSyncAsDataSync returns a boolean if a field has been set.

### SetTreatFileSyncAsDataSyncNil

`func (o *StorageDomain) SetTreatFileSyncAsDataSyncNil(b bool)`

 SetTreatFileSyncAsDataSyncNil sets the value for TreatFileSyncAsDataSync to be an explicit nil

### UnsetTreatFileSyncAsDataSync
`func (o *StorageDomain) UnsetTreatFileSyncAsDataSync()`

UnsetTreatFileSyncAsDataSync ensures that no value is present for TreatFileSyncAsDataSync, not even an explicit nil
### GetVaultId

`func (o *StorageDomain) GetVaultId() int64`

GetVaultId returns the VaultId field if non-nil, zero value otherwise.

### GetVaultIdOk

`func (o *StorageDomain) GetVaultIdOk() (*int64, bool)`

GetVaultIdOk returns a tuple with the VaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultId

`func (o *StorageDomain) SetVaultId(v int64)`

SetVaultId sets VaultId field to given value.

### HasVaultId

`func (o *StorageDomain) HasVaultId() bool`

HasVaultId returns a boolean if a field has been set.

### SetVaultIdNil

`func (o *StorageDomain) SetVaultIdNil(b bool)`

 SetVaultIdNil sets the value for VaultId to be an explicit nil

### UnsetVaultId
`func (o *StorageDomain) UnsetVaultId()`

UnsetVaultId ensures that no value is present for VaultId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


