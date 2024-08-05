# UpdateStorageDomainParam

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

### NewUpdateStorageDomainParam

`func NewUpdateStorageDomainParam(clusterPartitionId NullableInt64, name NullableString, ) *UpdateStorageDomainParam`

NewUpdateStorageDomainParam instantiates a new UpdateStorageDomainParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateStorageDomainParamWithDefaults

`func NewUpdateStorageDomainParamWithDefaults() *UpdateStorageDomainParam`

NewUpdateStorageDomainParamWithDefaults instantiates a new UpdateStorageDomainParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdDomainName

`func (o *UpdateStorageDomainParam) GetAdDomainName() string`

GetAdDomainName returns the AdDomainName field if non-nil, zero value otherwise.

### GetAdDomainNameOk

`func (o *UpdateStorageDomainParam) GetAdDomainNameOk() (*string, bool)`

GetAdDomainNameOk returns a tuple with the AdDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdDomainName

`func (o *UpdateStorageDomainParam) SetAdDomainName(v string)`

SetAdDomainName sets AdDomainName field to given value.

### HasAdDomainName

`func (o *UpdateStorageDomainParam) HasAdDomainName() bool`

HasAdDomainName returns a boolean if a field has been set.

### SetAdDomainNameNil

`func (o *UpdateStorageDomainParam) SetAdDomainNameNil(b bool)`

 SetAdDomainNameNil sets the value for AdDomainName to be an explicit nil

### UnsetAdDomainName
`func (o *UpdateStorageDomainParam) UnsetAdDomainName()`

UnsetAdDomainName ensures that no value is present for AdDomainName, not even an explicit nil
### GetBlobBrickSizeBytes

`func (o *UpdateStorageDomainParam) GetBlobBrickSizeBytes() int32`

GetBlobBrickSizeBytes returns the BlobBrickSizeBytes field if non-nil, zero value otherwise.

### GetBlobBrickSizeBytesOk

`func (o *UpdateStorageDomainParam) GetBlobBrickSizeBytesOk() (*int32, bool)`

GetBlobBrickSizeBytesOk returns a tuple with the BlobBrickSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobBrickSizeBytes

`func (o *UpdateStorageDomainParam) SetBlobBrickSizeBytes(v int32)`

SetBlobBrickSizeBytes sets BlobBrickSizeBytes field to given value.

### HasBlobBrickSizeBytes

`func (o *UpdateStorageDomainParam) HasBlobBrickSizeBytes() bool`

HasBlobBrickSizeBytes returns a boolean if a field has been set.

### SetBlobBrickSizeBytesNil

`func (o *UpdateStorageDomainParam) SetBlobBrickSizeBytesNil(b bool)`

 SetBlobBrickSizeBytesNil sets the value for BlobBrickSizeBytes to be an explicit nil

### UnsetBlobBrickSizeBytes
`func (o *UpdateStorageDomainParam) UnsetBlobBrickSizeBytes()`

UnsetBlobBrickSizeBytes ensures that no value is present for BlobBrickSizeBytes, not even an explicit nil
### GetCloudDomainId

`func (o *UpdateStorageDomainParam) GetCloudDomainId() int64`

GetCloudDomainId returns the CloudDomainId field if non-nil, zero value otherwise.

### GetCloudDomainIdOk

`func (o *UpdateStorageDomainParam) GetCloudDomainIdOk() (*int64, bool)`

GetCloudDomainIdOk returns a tuple with the CloudDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudDomainId

`func (o *UpdateStorageDomainParam) SetCloudDomainId(v int64)`

SetCloudDomainId sets CloudDomainId field to given value.

### HasCloudDomainId

`func (o *UpdateStorageDomainParam) HasCloudDomainId() bool`

HasCloudDomainId returns a boolean if a field has been set.

### SetCloudDomainIdNil

`func (o *UpdateStorageDomainParam) SetCloudDomainIdNil(b bool)`

 SetCloudDomainIdNil sets the value for CloudDomainId to be an explicit nil

### UnsetCloudDomainId
`func (o *UpdateStorageDomainParam) UnsetCloudDomainId()`

UnsetCloudDomainId ensures that no value is present for CloudDomainId, not even an explicit nil
### GetCloudDownWaterFallParams

`func (o *UpdateStorageDomainParam) GetCloudDownWaterFallParams() StorageDomainCloudDownWaterFallParams`

GetCloudDownWaterFallParams returns the CloudDownWaterFallParams field if non-nil, zero value otherwise.

### GetCloudDownWaterFallParamsOk

`func (o *UpdateStorageDomainParam) GetCloudDownWaterFallParamsOk() (*StorageDomainCloudDownWaterFallParams, bool)`

GetCloudDownWaterFallParamsOk returns a tuple with the CloudDownWaterFallParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudDownWaterFallParams

`func (o *UpdateStorageDomainParam) SetCloudDownWaterFallParams(v StorageDomainCloudDownWaterFallParams)`

SetCloudDownWaterFallParams sets CloudDownWaterFallParams field to given value.

### HasCloudDownWaterFallParams

`func (o *UpdateStorageDomainParam) HasCloudDownWaterFallParams() bool`

HasCloudDownWaterFallParams returns a boolean if a field has been set.

### GetClusterPartitionId

`func (o *UpdateStorageDomainParam) GetClusterPartitionId() int64`

GetClusterPartitionId returns the ClusterPartitionId field if non-nil, zero value otherwise.

### GetClusterPartitionIdOk

`func (o *UpdateStorageDomainParam) GetClusterPartitionIdOk() (*int64, bool)`

GetClusterPartitionIdOk returns a tuple with the ClusterPartitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterPartitionId

`func (o *UpdateStorageDomainParam) SetClusterPartitionId(v int64)`

SetClusterPartitionId sets ClusterPartitionId field to given value.


### SetClusterPartitionIdNil

`func (o *UpdateStorageDomainParam) SetClusterPartitionIdNil(b bool)`

 SetClusterPartitionIdNil sets the value for ClusterPartitionId to be an explicit nil

### UnsetClusterPartitionId
`func (o *UpdateStorageDomainParam) UnsetClusterPartitionId()`

UnsetClusterPartitionId ensures that no value is present for ClusterPartitionId, not even an explicit nil
### GetClusterPartitionName

`func (o *UpdateStorageDomainParam) GetClusterPartitionName() string`

GetClusterPartitionName returns the ClusterPartitionName field if non-nil, zero value otherwise.

### GetClusterPartitionNameOk

`func (o *UpdateStorageDomainParam) GetClusterPartitionNameOk() (*string, bool)`

GetClusterPartitionNameOk returns a tuple with the ClusterPartitionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterPartitionName

`func (o *UpdateStorageDomainParam) SetClusterPartitionName(v string)`

SetClusterPartitionName sets ClusterPartitionName field to given value.

### HasClusterPartitionName

`func (o *UpdateStorageDomainParam) HasClusterPartitionName() bool`

HasClusterPartitionName returns a boolean if a field has been set.

### SetClusterPartitionNameNil

`func (o *UpdateStorageDomainParam) SetClusterPartitionNameNil(b bool)`

 SetClusterPartitionNameNil sets the value for ClusterPartitionName to be an explicit nil

### UnsetClusterPartitionName
`func (o *UpdateStorageDomainParam) UnsetClusterPartitionName()`

UnsetClusterPartitionName ensures that no value is present for ClusterPartitionName, not even an explicit nil
### GetDefaultUserQuota

`func (o *UpdateStorageDomainParam) GetDefaultUserQuota() StorageDomainDefaultUserQuota`

GetDefaultUserQuota returns the DefaultUserQuota field if non-nil, zero value otherwise.

### GetDefaultUserQuotaOk

`func (o *UpdateStorageDomainParam) GetDefaultUserQuotaOk() (*StorageDomainDefaultUserQuota, bool)`

GetDefaultUserQuotaOk returns a tuple with the DefaultUserQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUserQuota

`func (o *UpdateStorageDomainParam) SetDefaultUserQuota(v StorageDomainDefaultUserQuota)`

SetDefaultUserQuota sets DefaultUserQuota field to given value.

### HasDefaultUserQuota

`func (o *UpdateStorageDomainParam) HasDefaultUserQuota() bool`

HasDefaultUserQuota returns a boolean if a field has been set.

### GetDefaultViewQuota

`func (o *UpdateStorageDomainParam) GetDefaultViewQuota() StorageDomainDefaultViewQuota`

GetDefaultViewQuota returns the DefaultViewQuota field if non-nil, zero value otherwise.

### GetDefaultViewQuotaOk

`func (o *UpdateStorageDomainParam) GetDefaultViewQuotaOk() (*StorageDomainDefaultViewQuota, bool)`

GetDefaultViewQuotaOk returns a tuple with the DefaultViewQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultViewQuota

`func (o *UpdateStorageDomainParam) SetDefaultViewQuota(v StorageDomainDefaultViewQuota)`

SetDefaultViewQuota sets DefaultViewQuota field to given value.

### HasDefaultViewQuota

`func (o *UpdateStorageDomainParam) HasDefaultViewQuota() bool`

HasDefaultViewQuota returns a boolean if a field has been set.

### GetDekRotationEnabled

`func (o *UpdateStorageDomainParam) GetDekRotationEnabled() bool`

GetDekRotationEnabled returns the DekRotationEnabled field if non-nil, zero value otherwise.

### GetDekRotationEnabledOk

`func (o *UpdateStorageDomainParam) GetDekRotationEnabledOk() (*bool, bool)`

GetDekRotationEnabledOk returns a tuple with the DekRotationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDekRotationEnabled

`func (o *UpdateStorageDomainParam) SetDekRotationEnabled(v bool)`

SetDekRotationEnabled sets DekRotationEnabled field to given value.

### HasDekRotationEnabled

`func (o *UpdateStorageDomainParam) HasDekRotationEnabled() bool`

HasDekRotationEnabled returns a boolean if a field has been set.

### SetDekRotationEnabledNil

`func (o *UpdateStorageDomainParam) SetDekRotationEnabledNil(b bool)`

 SetDekRotationEnabledNil sets the value for DekRotationEnabled to be an explicit nil

### UnsetDekRotationEnabled
`func (o *UpdateStorageDomainParam) UnsetDekRotationEnabled()`

UnsetDekRotationEnabled ensures that no value is present for DekRotationEnabled, not even an explicit nil
### GetDirectArchiveEnabled

`func (o *UpdateStorageDomainParam) GetDirectArchiveEnabled() bool`

GetDirectArchiveEnabled returns the DirectArchiveEnabled field if non-nil, zero value otherwise.

### GetDirectArchiveEnabledOk

`func (o *UpdateStorageDomainParam) GetDirectArchiveEnabledOk() (*bool, bool)`

GetDirectArchiveEnabledOk returns a tuple with the DirectArchiveEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectArchiveEnabled

`func (o *UpdateStorageDomainParam) SetDirectArchiveEnabled(v bool)`

SetDirectArchiveEnabled sets DirectArchiveEnabled field to given value.

### HasDirectArchiveEnabled

`func (o *UpdateStorageDomainParam) HasDirectArchiveEnabled() bool`

HasDirectArchiveEnabled returns a boolean if a field has been set.

### SetDirectArchiveEnabledNil

`func (o *UpdateStorageDomainParam) SetDirectArchiveEnabledNil(b bool)`

 SetDirectArchiveEnabledNil sets the value for DirectArchiveEnabled to be an explicit nil

### UnsetDirectArchiveEnabled
`func (o *UpdateStorageDomainParam) UnsetDirectArchiveEnabled()`

UnsetDirectArchiveEnabled ensures that no value is present for DirectArchiveEnabled, not even an explicit nil
### GetFileCountBySize

`func (o *UpdateStorageDomainParam) GetFileCountBySize() []FileCount`

GetFileCountBySize returns the FileCountBySize field if non-nil, zero value otherwise.

### GetFileCountBySizeOk

`func (o *UpdateStorageDomainParam) GetFileCountBySizeOk() (*[]FileCount, bool)`

GetFileCountBySizeOk returns a tuple with the FileCountBySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileCountBySize

`func (o *UpdateStorageDomainParam) SetFileCountBySize(v []FileCount)`

SetFileCountBySize sets FileCountBySize field to given value.

### HasFileCountBySize

`func (o *UpdateStorageDomainParam) HasFileCountBySize() bool`

HasFileCountBySize returns a boolean if a field has been set.

### SetFileCountBySizeNil

`func (o *UpdateStorageDomainParam) SetFileCountBySizeNil(b bool)`

 SetFileCountBySizeNil sets the value for FileCountBySize to be an explicit nil

### UnsetFileCountBySize
`func (o *UpdateStorageDomainParam) UnsetFileCountBySize()`

UnsetFileCountBySize ensures that no value is present for FileCountBySize, not even an explicit nil
### GetId

`func (o *UpdateStorageDomainParam) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateStorageDomainParam) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateStorageDomainParam) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateStorageDomainParam) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *UpdateStorageDomainParam) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *UpdateStorageDomainParam) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetKerberosRealmName

`func (o *UpdateStorageDomainParam) GetKerberosRealmName() string`

GetKerberosRealmName returns the KerberosRealmName field if non-nil, zero value otherwise.

### GetKerberosRealmNameOk

`func (o *UpdateStorageDomainParam) GetKerberosRealmNameOk() (*string, bool)`

GetKerberosRealmNameOk returns a tuple with the KerberosRealmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosRealmName

`func (o *UpdateStorageDomainParam) SetKerberosRealmName(v string)`

SetKerberosRealmName sets KerberosRealmName field to given value.

### HasKerberosRealmName

`func (o *UpdateStorageDomainParam) HasKerberosRealmName() bool`

HasKerberosRealmName returns a boolean if a field has been set.

### SetKerberosRealmNameNil

`func (o *UpdateStorageDomainParam) SetKerberosRealmNameNil(b bool)`

 SetKerberosRealmNameNil sets the value for KerberosRealmName to be an explicit nil

### UnsetKerberosRealmName
`func (o *UpdateStorageDomainParam) UnsetKerberosRealmName()`

UnsetKerberosRealmName ensures that no value is present for KerberosRealmName, not even an explicit nil
### GetKmsServerId

`func (o *UpdateStorageDomainParam) GetKmsServerId() int64`

GetKmsServerId returns the KmsServerId field if non-nil, zero value otherwise.

### GetKmsServerIdOk

`func (o *UpdateStorageDomainParam) GetKmsServerIdOk() (*int64, bool)`

GetKmsServerIdOk returns a tuple with the KmsServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsServerId

`func (o *UpdateStorageDomainParam) SetKmsServerId(v int64)`

SetKmsServerId sets KmsServerId field to given value.

### HasKmsServerId

`func (o *UpdateStorageDomainParam) HasKmsServerId() bool`

HasKmsServerId returns a boolean if a field has been set.

### SetKmsServerIdNil

`func (o *UpdateStorageDomainParam) SetKmsServerIdNil(b bool)`

 SetKmsServerIdNil sets the value for KmsServerId to be an explicit nil

### UnsetKmsServerId
`func (o *UpdateStorageDomainParam) UnsetKmsServerId()`

UnsetKmsServerId ensures that no value is present for KmsServerId, not even an explicit nil
### GetLastKeyRotationTimestampMsecs

`func (o *UpdateStorageDomainParam) GetLastKeyRotationTimestampMsecs() int64`

GetLastKeyRotationTimestampMsecs returns the LastKeyRotationTimestampMsecs field if non-nil, zero value otherwise.

### GetLastKeyRotationTimestampMsecsOk

`func (o *UpdateStorageDomainParam) GetLastKeyRotationTimestampMsecsOk() (*int64, bool)`

GetLastKeyRotationTimestampMsecsOk returns a tuple with the LastKeyRotationTimestampMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastKeyRotationTimestampMsecs

`func (o *UpdateStorageDomainParam) SetLastKeyRotationTimestampMsecs(v int64)`

SetLastKeyRotationTimestampMsecs sets LastKeyRotationTimestampMsecs field to given value.

### HasLastKeyRotationTimestampMsecs

`func (o *UpdateStorageDomainParam) HasLastKeyRotationTimestampMsecs() bool`

HasLastKeyRotationTimestampMsecs returns a boolean if a field has been set.

### SetLastKeyRotationTimestampMsecsNil

`func (o *UpdateStorageDomainParam) SetLastKeyRotationTimestampMsecsNil(b bool)`

 SetLastKeyRotationTimestampMsecsNil sets the value for LastKeyRotationTimestampMsecs to be an explicit nil

### UnsetLastKeyRotationTimestampMsecs
`func (o *UpdateStorageDomainParam) UnsetLastKeyRotationTimestampMsecs()`

UnsetLastKeyRotationTimestampMsecs ensures that no value is present for LastKeyRotationTimestampMsecs, not even an explicit nil
### GetLdapProviderId

`func (o *UpdateStorageDomainParam) GetLdapProviderId() int64`

GetLdapProviderId returns the LdapProviderId field if non-nil, zero value otherwise.

### GetLdapProviderIdOk

`func (o *UpdateStorageDomainParam) GetLdapProviderIdOk() (*int64, bool)`

GetLdapProviderIdOk returns a tuple with the LdapProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapProviderId

`func (o *UpdateStorageDomainParam) SetLdapProviderId(v int64)`

SetLdapProviderId sets LdapProviderId field to given value.

### HasLdapProviderId

`func (o *UpdateStorageDomainParam) HasLdapProviderId() bool`

HasLdapProviderId returns a boolean if a field has been set.

### SetLdapProviderIdNil

`func (o *UpdateStorageDomainParam) SetLdapProviderIdNil(b bool)`

 SetLdapProviderIdNil sets the value for LdapProviderId to be an explicit nil

### UnsetLdapProviderId
`func (o *UpdateStorageDomainParam) UnsetLdapProviderId()`

UnsetLdapProviderId ensures that no value is present for LdapProviderId, not even an explicit nil
### GetName

`func (o *UpdateStorageDomainParam) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateStorageDomainParam) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateStorageDomainParam) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *UpdateStorageDomainParam) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateStorageDomainParam) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNisDomainNames

`func (o *UpdateStorageDomainParam) GetNisDomainNames() []string`

GetNisDomainNames returns the NisDomainNames field if non-nil, zero value otherwise.

### GetNisDomainNamesOk

`func (o *UpdateStorageDomainParam) GetNisDomainNamesOk() (*[]string, bool)`

GetNisDomainNamesOk returns a tuple with the NisDomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNisDomainNames

`func (o *UpdateStorageDomainParam) SetNisDomainNames(v []string)`

SetNisDomainNames sets NisDomainNames field to given value.

### HasNisDomainNames

`func (o *UpdateStorageDomainParam) HasNisDomainNames() bool`

HasNisDomainNames returns a boolean if a field has been set.

### SetNisDomainNamesNil

`func (o *UpdateStorageDomainParam) SetNisDomainNamesNil(b bool)`

 SetNisDomainNamesNil sets the value for NisDomainNames to be an explicit nil

### UnsetNisDomainNames
`func (o *UpdateStorageDomainParam) UnsetNisDomainNames()`

UnsetNisDomainNames ensures that no value is present for NisDomainNames, not even an explicit nil
### GetPhysicalQuota

`func (o *UpdateStorageDomainParam) GetPhysicalQuota() StorageDomainPhysicalQuota`

GetPhysicalQuota returns the PhysicalQuota field if non-nil, zero value otherwise.

### GetPhysicalQuotaOk

`func (o *UpdateStorageDomainParam) GetPhysicalQuotaOk() (*StorageDomainPhysicalQuota, bool)`

GetPhysicalQuotaOk returns a tuple with the PhysicalQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalQuota

`func (o *UpdateStorageDomainParam) SetPhysicalQuota(v StorageDomainPhysicalQuota)`

SetPhysicalQuota sets PhysicalQuota field to given value.

### HasPhysicalQuota

`func (o *UpdateStorageDomainParam) HasPhysicalQuota() bool`

HasPhysicalQuota returns a boolean if a field has been set.

### GetRecommended

`func (o *UpdateStorageDomainParam) GetRecommended() bool`

GetRecommended returns the Recommended field if non-nil, zero value otherwise.

### GetRecommendedOk

`func (o *UpdateStorageDomainParam) GetRecommendedOk() (*bool, bool)`

GetRecommendedOk returns a tuple with the Recommended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommended

`func (o *UpdateStorageDomainParam) SetRecommended(v bool)`

SetRecommended sets Recommended field to given value.

### HasRecommended

`func (o *UpdateStorageDomainParam) HasRecommended() bool`

HasRecommended returns a boolean if a field has been set.

### SetRecommendedNil

`func (o *UpdateStorageDomainParam) SetRecommendedNil(b bool)`

 SetRecommendedNil sets the value for Recommended to be an explicit nil

### UnsetRecommended
`func (o *UpdateStorageDomainParam) UnsetRecommended()`

UnsetRecommended ensures that no value is present for Recommended, not even an explicit nil
### GetRemovalState

`func (o *UpdateStorageDomainParam) GetRemovalState() string`

GetRemovalState returns the RemovalState field if non-nil, zero value otherwise.

### GetRemovalStateOk

`func (o *UpdateStorageDomainParam) GetRemovalStateOk() (*string, bool)`

GetRemovalStateOk returns a tuple with the RemovalState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalState

`func (o *UpdateStorageDomainParam) SetRemovalState(v string)`

SetRemovalState sets RemovalState field to given value.

### HasRemovalState

`func (o *UpdateStorageDomainParam) HasRemovalState() bool`

HasRemovalState returns a boolean if a field has been set.

### SetRemovalStateNil

`func (o *UpdateStorageDomainParam) SetRemovalStateNil(b bool)`

 SetRemovalStateNil sets the value for RemovalState to be an explicit nil

### UnsetRemovalState
`func (o *UpdateStorageDomainParam) UnsetRemovalState()`

UnsetRemovalState ensures that no value is present for RemovalState, not even an explicit nil
### GetS3BucketsEnabled

`func (o *UpdateStorageDomainParam) GetS3BucketsEnabled() bool`

GetS3BucketsEnabled returns the S3BucketsEnabled field if non-nil, zero value otherwise.

### GetS3BucketsEnabledOk

`func (o *UpdateStorageDomainParam) GetS3BucketsEnabledOk() (*bool, bool)`

GetS3BucketsEnabledOk returns a tuple with the S3BucketsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3BucketsEnabled

`func (o *UpdateStorageDomainParam) SetS3BucketsEnabled(v bool)`

SetS3BucketsEnabled sets S3BucketsEnabled field to given value.

### HasS3BucketsEnabled

`func (o *UpdateStorageDomainParam) HasS3BucketsEnabled() bool`

HasS3BucketsEnabled returns a boolean if a field has been set.

### SetS3BucketsEnabledNil

`func (o *UpdateStorageDomainParam) SetS3BucketsEnabledNil(b bool)`

 SetS3BucketsEnabledNil sets the value for S3BucketsEnabled to be an explicit nil

### UnsetS3BucketsEnabled
`func (o *UpdateStorageDomainParam) UnsetS3BucketsEnabled()`

UnsetS3BucketsEnabled ensures that no value is present for S3BucketsEnabled, not even an explicit nil
### GetSchemas

`func (o *UpdateStorageDomainParam) GetSchemas() []Schema`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *UpdateStorageDomainParam) GetSchemasOk() (*[]Schema, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *UpdateStorageDomainParam) SetSchemas(v []Schema)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *UpdateStorageDomainParam) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### SetSchemasNil

`func (o *UpdateStorageDomainParam) SetSchemasNil(b bool)`

 SetSchemasNil sets the value for Schemas to be an explicit nil

### UnsetSchemas
`func (o *UpdateStorageDomainParam) UnsetSchemas()`

UnsetSchemas ensures that no value is present for Schemas, not even an explicit nil
### GetStats

`func (o *UpdateStorageDomainParam) GetStats() StorageDomainStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *UpdateStorageDomainParam) GetStatsOk() (*StorageDomainStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *UpdateStorageDomainParam) SetStats(v StorageDomainStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *UpdateStorageDomainParam) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetStoragePolicy

`func (o *UpdateStorageDomainParam) GetStoragePolicy() StorageDomainStoragePolicy`

GetStoragePolicy returns the StoragePolicy field if non-nil, zero value otherwise.

### GetStoragePolicyOk

`func (o *UpdateStorageDomainParam) GetStoragePolicyOk() (*StorageDomainStoragePolicy, bool)`

GetStoragePolicyOk returns a tuple with the StoragePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePolicy

`func (o *UpdateStorageDomainParam) SetStoragePolicy(v StorageDomainStoragePolicy)`

SetStoragePolicy sets StoragePolicy field to given value.

### HasStoragePolicy

`func (o *UpdateStorageDomainParam) HasStoragePolicy() bool`

HasStoragePolicy returns a boolean if a field has been set.

### GetSubnetWhitelist

`func (o *UpdateStorageDomainParam) GetSubnetWhitelist() []Subnet`

GetSubnetWhitelist returns the SubnetWhitelist field if non-nil, zero value otherwise.

### GetSubnetWhitelistOk

`func (o *UpdateStorageDomainParam) GetSubnetWhitelistOk() (*[]Subnet, bool)`

GetSubnetWhitelistOk returns a tuple with the SubnetWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetWhitelist

`func (o *UpdateStorageDomainParam) SetSubnetWhitelist(v []Subnet)`

SetSubnetWhitelist sets SubnetWhitelist field to given value.

### HasSubnetWhitelist

`func (o *UpdateStorageDomainParam) HasSubnetWhitelist() bool`

HasSubnetWhitelist returns a boolean if a field has been set.

### SetSubnetWhitelistNil

`func (o *UpdateStorageDomainParam) SetSubnetWhitelistNil(b bool)`

 SetSubnetWhitelistNil sets the value for SubnetWhitelist to be an explicit nil

### UnsetSubnetWhitelist
`func (o *UpdateStorageDomainParam) UnsetSubnetWhitelist()`

UnsetSubnetWhitelist ensures that no value is present for SubnetWhitelist, not even an explicit nil
### GetTenantIds

`func (o *UpdateStorageDomainParam) GetTenantIds() []string`

GetTenantIds returns the TenantIds field if non-nil, zero value otherwise.

### GetTenantIdsOk

`func (o *UpdateStorageDomainParam) GetTenantIdsOk() (*[]string, bool)`

GetTenantIdsOk returns a tuple with the TenantIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIds

`func (o *UpdateStorageDomainParam) SetTenantIds(v []string)`

SetTenantIds sets TenantIds field to given value.

### HasTenantIds

`func (o *UpdateStorageDomainParam) HasTenantIds() bool`

HasTenantIds returns a boolean if a field has been set.

### SetTenantIdsNil

`func (o *UpdateStorageDomainParam) SetTenantIdsNil(b bool)`

 SetTenantIdsNil sets the value for TenantIds to be an explicit nil

### UnsetTenantIds
`func (o *UpdateStorageDomainParam) UnsetTenantIds()`

UnsetTenantIds ensures that no value is present for TenantIds, not even an explicit nil
### GetTreatFileSyncAsDataSync

`func (o *UpdateStorageDomainParam) GetTreatFileSyncAsDataSync() bool`

GetTreatFileSyncAsDataSync returns the TreatFileSyncAsDataSync field if non-nil, zero value otherwise.

### GetTreatFileSyncAsDataSyncOk

`func (o *UpdateStorageDomainParam) GetTreatFileSyncAsDataSyncOk() (*bool, bool)`

GetTreatFileSyncAsDataSyncOk returns a tuple with the TreatFileSyncAsDataSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreatFileSyncAsDataSync

`func (o *UpdateStorageDomainParam) SetTreatFileSyncAsDataSync(v bool)`

SetTreatFileSyncAsDataSync sets TreatFileSyncAsDataSync field to given value.

### HasTreatFileSyncAsDataSync

`func (o *UpdateStorageDomainParam) HasTreatFileSyncAsDataSync() bool`

HasTreatFileSyncAsDataSync returns a boolean if a field has been set.

### SetTreatFileSyncAsDataSyncNil

`func (o *UpdateStorageDomainParam) SetTreatFileSyncAsDataSyncNil(b bool)`

 SetTreatFileSyncAsDataSyncNil sets the value for TreatFileSyncAsDataSync to be an explicit nil

### UnsetTreatFileSyncAsDataSync
`func (o *UpdateStorageDomainParam) UnsetTreatFileSyncAsDataSync()`

UnsetTreatFileSyncAsDataSync ensures that no value is present for TreatFileSyncAsDataSync, not even an explicit nil
### GetVaultId

`func (o *UpdateStorageDomainParam) GetVaultId() int64`

GetVaultId returns the VaultId field if non-nil, zero value otherwise.

### GetVaultIdOk

`func (o *UpdateStorageDomainParam) GetVaultIdOk() (*int64, bool)`

GetVaultIdOk returns a tuple with the VaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultId

`func (o *UpdateStorageDomainParam) SetVaultId(v int64)`

SetVaultId sets VaultId field to given value.

### HasVaultId

`func (o *UpdateStorageDomainParam) HasVaultId() bool`

HasVaultId returns a boolean if a field has been set.

### SetVaultIdNil

`func (o *UpdateStorageDomainParam) SetVaultIdNil(b bool)`

 SetVaultIdNil sets the value for VaultId to be an explicit nil

### UnsetVaultId
`func (o *UpdateStorageDomainParam) UnsetVaultId()`

UnsetVaultId ensures that no value is present for VaultId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


