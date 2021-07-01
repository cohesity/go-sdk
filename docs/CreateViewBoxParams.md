# CreateViewBoxParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdDomainName** | Pointer to **NullableString** | Specifies an active directory domain that this view box is mapped to. | [optional] 
**ClientSubnetWhiteList** | Pointer to [**[]Subnet**](Subnet.md) | Array of Subnets.  Specifies the Subnets from which this Storage Domain (View Box) accepts requests. | [optional] 
**CloudDownWaterfallThresholdPct** | Pointer to **NullableInt32** | Specifies the cloud down water-fall threshold percentage. This indicates how full should a viewbox at least be before we down water-fall its data to cloud tier. If this field is set, the physical quota limit must be set also and will be used as viewbox capacity. | [optional] 
**CloudDownWaterfallThresholdSecs** | Pointer to **NullableInt32** | Specifies the cloud down water-fall threshold seconds. This indicates what&#39;s the time threshold on water-falling data to cloud tier. | [optional] 
**ClusterPartitionId** | **NullableInt64** | Specifies the Cluster Partition id where the Storage Domain (View Box) is located. | 
**DefaultUserQuotaPolicy** | Pointer to [**NullableQuotaPolicy**](QuotaPolicy.md) | Specifies an optional quota policy/limits that are inherited by all users within the views in this viewbox. | [optional] 
**DefaultViewQuotaPolicy** | Pointer to [**NullableQuotaPolicy**](QuotaPolicy.md) | Specifies an optional default logical quota limit (in bytes) for the Views in this Storage Domain (View Box). (Logical data is when the data is fully hydrated and expanded.) However, this inherited quota can be overwritten at the View level. A new write is not allowed if the Storage Domain (View Box) will exceed the specified quota. However, it takes time for the Cohesity Cluster to calculate the usage across Nodes, so the limit may be exceeded by a small amount. In addition, if the limit is increased or data is removed, there may be delay before the Cohesity Cluster allows more data to be written to the Storage Domain (View Box), as the Cluster is calculating the usage across Nodes. | [optional] 
**DirectArchiveEnabled** | Pointer to **NullableBool** | Specifies whether this viewbox can be used as a staging area while copying a largedataset that can&#39;t fit on the cluster to an external target. The amount of data that can be stored on the viewbox can be specified using &#39;physical_quota&#39;. | [optional] 
**LdapProviderId** | Pointer to **NullableInt64** | When set, the following provides the LDAP provider the view box is mapped to. For any view from this view box, when accessed via NFS the following LDAP provider is looked up for getting Unix IDs of the corresponding user. Similarly, when a view is accessed via SMB and if the AD user&#39;s domain matches with the view box&#39;s AD, the following LDAP provider will be used to lookup Unix IDs for the corresponding AD user. Additionally there is also a mapping between LDAP provider and AD domain that is stored in AD provider config. It will be used if AD is not set on the view box. | [optional] 
**Name** | **NullableString** | Specifies the name of the Storage Domain (View Box). | 
**NisDomainNameVec** | Pointer to **[]string** | Specifies the NIS domain that this view box is mapped to. | [optional] 
**PhysicalQuota** | Pointer to [**NullableQuotaPolicy**](QuotaPolicy.md) | Specifies an optional quota limit (in bytes) for the physical usage of this Storage Domain (View Box). This quota limit defines a physical limit for size of the data that can be physically stored on the Storage Domain (View Box), after the data has been reduced by change block tracking, compression and deduplication. The physical usage is the aggregate sum of the data stored for this Storage Domain (View Box) on all disks in the Cluster. (The usage includes Cloud Tier data and user data.) A new write is not allowed if the Storage Domain (View Box) will exceed the specified quota. However, it takes time for the Cohesity Cluster to calculate the usage across Nodes, so the limit may be exceeded by a small amount. In addition, if the limit is increased or data is removed, there may be a delay before the Cohesity Cluster allows more data to be written to the Storage Domain (View Box), as the Cluster is calculating the usage across Nodes. | [optional] 
**S3BucketsAllowed** | Pointer to **NullableBool** | Specifies whether creation of a S3 bucket is allowed in this Storage Domain (View Box). When a new S3 bucket creation request arrives, we&#39;ll look at all the View Boxes and the first Storage Domain (View Box) that allows creating S3 buckets in it will be the one where the bucket will be placed. | [optional] 
**StoragePolicy** | Pointer to [**StoragePolicy**](StoragePolicy.md) |  | [optional] 
**TenantIdVec** | Pointer to **[]string** | Optional ids for the tenants that this view box belongs. This must be checked before granting access to users. Unless the cluster enables view box sharing between tenants is allowed, there shall be at most one item in this list. Note that if all tenant may be deleted - such viewboxes must be garbage collected. This is currently done by a background thread in iris. | [optional] 

## Methods

### NewCreateViewBoxParams

`func NewCreateViewBoxParams(clusterPartitionId NullableInt64, name NullableString, ) *CreateViewBoxParams`

NewCreateViewBoxParams instantiates a new CreateViewBoxParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateViewBoxParamsWithDefaults

`func NewCreateViewBoxParamsWithDefaults() *CreateViewBoxParams`

NewCreateViewBoxParamsWithDefaults instantiates a new CreateViewBoxParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdDomainName

`func (o *CreateViewBoxParams) GetAdDomainName() string`

GetAdDomainName returns the AdDomainName field if non-nil, zero value otherwise.

### GetAdDomainNameOk

`func (o *CreateViewBoxParams) GetAdDomainNameOk() (*string, bool)`

GetAdDomainNameOk returns a tuple with the AdDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdDomainName

`func (o *CreateViewBoxParams) SetAdDomainName(v string)`

SetAdDomainName sets AdDomainName field to given value.

### HasAdDomainName

`func (o *CreateViewBoxParams) HasAdDomainName() bool`

HasAdDomainName returns a boolean if a field has been set.

### SetAdDomainNameNil

`func (o *CreateViewBoxParams) SetAdDomainNameNil(b bool)`

 SetAdDomainNameNil sets the value for AdDomainName to be an explicit nil

### UnsetAdDomainName
`func (o *CreateViewBoxParams) UnsetAdDomainName()`

UnsetAdDomainName ensures that no value is present for AdDomainName, not even an explicit nil
### GetClientSubnetWhiteList

`func (o *CreateViewBoxParams) GetClientSubnetWhiteList() []Subnet`

GetClientSubnetWhiteList returns the ClientSubnetWhiteList field if non-nil, zero value otherwise.

### GetClientSubnetWhiteListOk

`func (o *CreateViewBoxParams) GetClientSubnetWhiteListOk() (*[]Subnet, bool)`

GetClientSubnetWhiteListOk returns a tuple with the ClientSubnetWhiteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSubnetWhiteList

`func (o *CreateViewBoxParams) SetClientSubnetWhiteList(v []Subnet)`

SetClientSubnetWhiteList sets ClientSubnetWhiteList field to given value.

### HasClientSubnetWhiteList

`func (o *CreateViewBoxParams) HasClientSubnetWhiteList() bool`

HasClientSubnetWhiteList returns a boolean if a field has been set.

### SetClientSubnetWhiteListNil

`func (o *CreateViewBoxParams) SetClientSubnetWhiteListNil(b bool)`

 SetClientSubnetWhiteListNil sets the value for ClientSubnetWhiteList to be an explicit nil

### UnsetClientSubnetWhiteList
`func (o *CreateViewBoxParams) UnsetClientSubnetWhiteList()`

UnsetClientSubnetWhiteList ensures that no value is present for ClientSubnetWhiteList, not even an explicit nil
### GetCloudDownWaterfallThresholdPct

`func (o *CreateViewBoxParams) GetCloudDownWaterfallThresholdPct() int32`

GetCloudDownWaterfallThresholdPct returns the CloudDownWaterfallThresholdPct field if non-nil, zero value otherwise.

### GetCloudDownWaterfallThresholdPctOk

`func (o *CreateViewBoxParams) GetCloudDownWaterfallThresholdPctOk() (*int32, bool)`

GetCloudDownWaterfallThresholdPctOk returns a tuple with the CloudDownWaterfallThresholdPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudDownWaterfallThresholdPct

`func (o *CreateViewBoxParams) SetCloudDownWaterfallThresholdPct(v int32)`

SetCloudDownWaterfallThresholdPct sets CloudDownWaterfallThresholdPct field to given value.

### HasCloudDownWaterfallThresholdPct

`func (o *CreateViewBoxParams) HasCloudDownWaterfallThresholdPct() bool`

HasCloudDownWaterfallThresholdPct returns a boolean if a field has been set.

### SetCloudDownWaterfallThresholdPctNil

`func (o *CreateViewBoxParams) SetCloudDownWaterfallThresholdPctNil(b bool)`

 SetCloudDownWaterfallThresholdPctNil sets the value for CloudDownWaterfallThresholdPct to be an explicit nil

### UnsetCloudDownWaterfallThresholdPct
`func (o *CreateViewBoxParams) UnsetCloudDownWaterfallThresholdPct()`

UnsetCloudDownWaterfallThresholdPct ensures that no value is present for CloudDownWaterfallThresholdPct, not even an explicit nil
### GetCloudDownWaterfallThresholdSecs

`func (o *CreateViewBoxParams) GetCloudDownWaterfallThresholdSecs() int32`

GetCloudDownWaterfallThresholdSecs returns the CloudDownWaterfallThresholdSecs field if non-nil, zero value otherwise.

### GetCloudDownWaterfallThresholdSecsOk

`func (o *CreateViewBoxParams) GetCloudDownWaterfallThresholdSecsOk() (*int32, bool)`

GetCloudDownWaterfallThresholdSecsOk returns a tuple with the CloudDownWaterfallThresholdSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudDownWaterfallThresholdSecs

`func (o *CreateViewBoxParams) SetCloudDownWaterfallThresholdSecs(v int32)`

SetCloudDownWaterfallThresholdSecs sets CloudDownWaterfallThresholdSecs field to given value.

### HasCloudDownWaterfallThresholdSecs

`func (o *CreateViewBoxParams) HasCloudDownWaterfallThresholdSecs() bool`

HasCloudDownWaterfallThresholdSecs returns a boolean if a field has been set.

### SetCloudDownWaterfallThresholdSecsNil

`func (o *CreateViewBoxParams) SetCloudDownWaterfallThresholdSecsNil(b bool)`

 SetCloudDownWaterfallThresholdSecsNil sets the value for CloudDownWaterfallThresholdSecs to be an explicit nil

### UnsetCloudDownWaterfallThresholdSecs
`func (o *CreateViewBoxParams) UnsetCloudDownWaterfallThresholdSecs()`

UnsetCloudDownWaterfallThresholdSecs ensures that no value is present for CloudDownWaterfallThresholdSecs, not even an explicit nil
### GetClusterPartitionId

`func (o *CreateViewBoxParams) GetClusterPartitionId() int64`

GetClusterPartitionId returns the ClusterPartitionId field if non-nil, zero value otherwise.

### GetClusterPartitionIdOk

`func (o *CreateViewBoxParams) GetClusterPartitionIdOk() (*int64, bool)`

GetClusterPartitionIdOk returns a tuple with the ClusterPartitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterPartitionId

`func (o *CreateViewBoxParams) SetClusterPartitionId(v int64)`

SetClusterPartitionId sets ClusterPartitionId field to given value.


### SetClusterPartitionIdNil

`func (o *CreateViewBoxParams) SetClusterPartitionIdNil(b bool)`

 SetClusterPartitionIdNil sets the value for ClusterPartitionId to be an explicit nil

### UnsetClusterPartitionId
`func (o *CreateViewBoxParams) UnsetClusterPartitionId()`

UnsetClusterPartitionId ensures that no value is present for ClusterPartitionId, not even an explicit nil
### GetDefaultUserQuotaPolicy

`func (o *CreateViewBoxParams) GetDefaultUserQuotaPolicy() QuotaPolicy`

GetDefaultUserQuotaPolicy returns the DefaultUserQuotaPolicy field if non-nil, zero value otherwise.

### GetDefaultUserQuotaPolicyOk

`func (o *CreateViewBoxParams) GetDefaultUserQuotaPolicyOk() (*QuotaPolicy, bool)`

GetDefaultUserQuotaPolicyOk returns a tuple with the DefaultUserQuotaPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUserQuotaPolicy

`func (o *CreateViewBoxParams) SetDefaultUserQuotaPolicy(v QuotaPolicy)`

SetDefaultUserQuotaPolicy sets DefaultUserQuotaPolicy field to given value.

### HasDefaultUserQuotaPolicy

`func (o *CreateViewBoxParams) HasDefaultUserQuotaPolicy() bool`

HasDefaultUserQuotaPolicy returns a boolean if a field has been set.

### SetDefaultUserQuotaPolicyNil

`func (o *CreateViewBoxParams) SetDefaultUserQuotaPolicyNil(b bool)`

 SetDefaultUserQuotaPolicyNil sets the value for DefaultUserQuotaPolicy to be an explicit nil

### UnsetDefaultUserQuotaPolicy
`func (o *CreateViewBoxParams) UnsetDefaultUserQuotaPolicy()`

UnsetDefaultUserQuotaPolicy ensures that no value is present for DefaultUserQuotaPolicy, not even an explicit nil
### GetDefaultViewQuotaPolicy

`func (o *CreateViewBoxParams) GetDefaultViewQuotaPolicy() QuotaPolicy`

GetDefaultViewQuotaPolicy returns the DefaultViewQuotaPolicy field if non-nil, zero value otherwise.

### GetDefaultViewQuotaPolicyOk

`func (o *CreateViewBoxParams) GetDefaultViewQuotaPolicyOk() (*QuotaPolicy, bool)`

GetDefaultViewQuotaPolicyOk returns a tuple with the DefaultViewQuotaPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultViewQuotaPolicy

`func (o *CreateViewBoxParams) SetDefaultViewQuotaPolicy(v QuotaPolicy)`

SetDefaultViewQuotaPolicy sets DefaultViewQuotaPolicy field to given value.

### HasDefaultViewQuotaPolicy

`func (o *CreateViewBoxParams) HasDefaultViewQuotaPolicy() bool`

HasDefaultViewQuotaPolicy returns a boolean if a field has been set.

### SetDefaultViewQuotaPolicyNil

`func (o *CreateViewBoxParams) SetDefaultViewQuotaPolicyNil(b bool)`

 SetDefaultViewQuotaPolicyNil sets the value for DefaultViewQuotaPolicy to be an explicit nil

### UnsetDefaultViewQuotaPolicy
`func (o *CreateViewBoxParams) UnsetDefaultViewQuotaPolicy()`

UnsetDefaultViewQuotaPolicy ensures that no value is present for DefaultViewQuotaPolicy, not even an explicit nil
### GetDirectArchiveEnabled

`func (o *CreateViewBoxParams) GetDirectArchiveEnabled() bool`

GetDirectArchiveEnabled returns the DirectArchiveEnabled field if non-nil, zero value otherwise.

### GetDirectArchiveEnabledOk

`func (o *CreateViewBoxParams) GetDirectArchiveEnabledOk() (*bool, bool)`

GetDirectArchiveEnabledOk returns a tuple with the DirectArchiveEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectArchiveEnabled

`func (o *CreateViewBoxParams) SetDirectArchiveEnabled(v bool)`

SetDirectArchiveEnabled sets DirectArchiveEnabled field to given value.

### HasDirectArchiveEnabled

`func (o *CreateViewBoxParams) HasDirectArchiveEnabled() bool`

HasDirectArchiveEnabled returns a boolean if a field has been set.

### SetDirectArchiveEnabledNil

`func (o *CreateViewBoxParams) SetDirectArchiveEnabledNil(b bool)`

 SetDirectArchiveEnabledNil sets the value for DirectArchiveEnabled to be an explicit nil

### UnsetDirectArchiveEnabled
`func (o *CreateViewBoxParams) UnsetDirectArchiveEnabled()`

UnsetDirectArchiveEnabled ensures that no value is present for DirectArchiveEnabled, not even an explicit nil
### GetLdapProviderId

`func (o *CreateViewBoxParams) GetLdapProviderId() int64`

GetLdapProviderId returns the LdapProviderId field if non-nil, zero value otherwise.

### GetLdapProviderIdOk

`func (o *CreateViewBoxParams) GetLdapProviderIdOk() (*int64, bool)`

GetLdapProviderIdOk returns a tuple with the LdapProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapProviderId

`func (o *CreateViewBoxParams) SetLdapProviderId(v int64)`

SetLdapProviderId sets LdapProviderId field to given value.

### HasLdapProviderId

`func (o *CreateViewBoxParams) HasLdapProviderId() bool`

HasLdapProviderId returns a boolean if a field has been set.

### SetLdapProviderIdNil

`func (o *CreateViewBoxParams) SetLdapProviderIdNil(b bool)`

 SetLdapProviderIdNil sets the value for LdapProviderId to be an explicit nil

### UnsetLdapProviderId
`func (o *CreateViewBoxParams) UnsetLdapProviderId()`

UnsetLdapProviderId ensures that no value is present for LdapProviderId, not even an explicit nil
### GetName

`func (o *CreateViewBoxParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateViewBoxParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateViewBoxParams) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CreateViewBoxParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateViewBoxParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNisDomainNameVec

`func (o *CreateViewBoxParams) GetNisDomainNameVec() []string`

GetNisDomainNameVec returns the NisDomainNameVec field if non-nil, zero value otherwise.

### GetNisDomainNameVecOk

`func (o *CreateViewBoxParams) GetNisDomainNameVecOk() (*[]string, bool)`

GetNisDomainNameVecOk returns a tuple with the NisDomainNameVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNisDomainNameVec

`func (o *CreateViewBoxParams) SetNisDomainNameVec(v []string)`

SetNisDomainNameVec sets NisDomainNameVec field to given value.

### HasNisDomainNameVec

`func (o *CreateViewBoxParams) HasNisDomainNameVec() bool`

HasNisDomainNameVec returns a boolean if a field has been set.

### SetNisDomainNameVecNil

`func (o *CreateViewBoxParams) SetNisDomainNameVecNil(b bool)`

 SetNisDomainNameVecNil sets the value for NisDomainNameVec to be an explicit nil

### UnsetNisDomainNameVec
`func (o *CreateViewBoxParams) UnsetNisDomainNameVec()`

UnsetNisDomainNameVec ensures that no value is present for NisDomainNameVec, not even an explicit nil
### GetPhysicalQuota

`func (o *CreateViewBoxParams) GetPhysicalQuota() QuotaPolicy`

GetPhysicalQuota returns the PhysicalQuota field if non-nil, zero value otherwise.

### GetPhysicalQuotaOk

`func (o *CreateViewBoxParams) GetPhysicalQuotaOk() (*QuotaPolicy, bool)`

GetPhysicalQuotaOk returns a tuple with the PhysicalQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalQuota

`func (o *CreateViewBoxParams) SetPhysicalQuota(v QuotaPolicy)`

SetPhysicalQuota sets PhysicalQuota field to given value.

### HasPhysicalQuota

`func (o *CreateViewBoxParams) HasPhysicalQuota() bool`

HasPhysicalQuota returns a boolean if a field has been set.

### SetPhysicalQuotaNil

`func (o *CreateViewBoxParams) SetPhysicalQuotaNil(b bool)`

 SetPhysicalQuotaNil sets the value for PhysicalQuota to be an explicit nil

### UnsetPhysicalQuota
`func (o *CreateViewBoxParams) UnsetPhysicalQuota()`

UnsetPhysicalQuota ensures that no value is present for PhysicalQuota, not even an explicit nil
### GetS3BucketsAllowed

`func (o *CreateViewBoxParams) GetS3BucketsAllowed() bool`

GetS3BucketsAllowed returns the S3BucketsAllowed field if non-nil, zero value otherwise.

### GetS3BucketsAllowedOk

`func (o *CreateViewBoxParams) GetS3BucketsAllowedOk() (*bool, bool)`

GetS3BucketsAllowedOk returns a tuple with the S3BucketsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3BucketsAllowed

`func (o *CreateViewBoxParams) SetS3BucketsAllowed(v bool)`

SetS3BucketsAllowed sets S3BucketsAllowed field to given value.

### HasS3BucketsAllowed

`func (o *CreateViewBoxParams) HasS3BucketsAllowed() bool`

HasS3BucketsAllowed returns a boolean if a field has been set.

### SetS3BucketsAllowedNil

`func (o *CreateViewBoxParams) SetS3BucketsAllowedNil(b bool)`

 SetS3BucketsAllowedNil sets the value for S3BucketsAllowed to be an explicit nil

### UnsetS3BucketsAllowed
`func (o *CreateViewBoxParams) UnsetS3BucketsAllowed()`

UnsetS3BucketsAllowed ensures that no value is present for S3BucketsAllowed, not even an explicit nil
### GetStoragePolicy

`func (o *CreateViewBoxParams) GetStoragePolicy() StoragePolicy`

GetStoragePolicy returns the StoragePolicy field if non-nil, zero value otherwise.

### GetStoragePolicyOk

`func (o *CreateViewBoxParams) GetStoragePolicyOk() (*StoragePolicy, bool)`

GetStoragePolicyOk returns a tuple with the StoragePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePolicy

`func (o *CreateViewBoxParams) SetStoragePolicy(v StoragePolicy)`

SetStoragePolicy sets StoragePolicy field to given value.

### HasStoragePolicy

`func (o *CreateViewBoxParams) HasStoragePolicy() bool`

HasStoragePolicy returns a boolean if a field has been set.

### GetTenantIdVec

`func (o *CreateViewBoxParams) GetTenantIdVec() []string`

GetTenantIdVec returns the TenantIdVec field if non-nil, zero value otherwise.

### GetTenantIdVecOk

`func (o *CreateViewBoxParams) GetTenantIdVecOk() (*[]string, bool)`

GetTenantIdVecOk returns a tuple with the TenantIdVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIdVec

`func (o *CreateViewBoxParams) SetTenantIdVec(v []string)`

SetTenantIdVec sets TenantIdVec field to given value.

### HasTenantIdVec

`func (o *CreateViewBoxParams) HasTenantIdVec() bool`

HasTenantIdVec returns a boolean if a field has been set.

### SetTenantIdVecNil

`func (o *CreateViewBoxParams) SetTenantIdVecNil(b bool)`

 SetTenantIdVecNil sets the value for TenantIdVec to be an explicit nil

### UnsetTenantIdVec
`func (o *CreateViewBoxParams) UnsetTenantIdVec()`

UnsetTenantIdVec ensures that no value is present for TenantIdVec, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


