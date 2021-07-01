# ViewBox

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdDomainName** | Pointer to **NullableString** | Specifies an active directory domain that this view box is mapped to. | [optional] 
**ClientSubnetWhiteList** | Pointer to [**[]Subnet**](Subnet.md) | Array of Subnets.  Specifies the Subnets from which this Storage Domain (View Box) accepts requests. | [optional] 
**CloudDownWaterfallThresholdPct** | Pointer to **NullableInt32** | Specifies the cloud down water-fall threshold percentage. This indicates how full should a viewbox at least be before we down water-fall its data to cloud tier. If this field is set, the physical quota limit must be set also and will be used as viewbox capacity. | [optional] 
**CloudDownWaterfallThresholdSecs** | Pointer to **NullableInt32** | Specifies the cloud down water-fall threshold seconds. This indicates what&#39;s the time threshold on water-falling data to cloud tier. | [optional] 
**ClusterPartitionId** | **NullableInt64** | Specifies the Cluster Partition id where the Storage Domain (View Box) is located. | 
**ClusterPartitionName** | Pointer to **NullableString** | Specifies the Cohesity Cluster name where the Storage Domain (View Box) is located. | [optional] [readonly] 
**DefaultUserQuotaPolicy** | Pointer to [**NullableQuotaPolicy**](QuotaPolicy.md) | Specifies an optional quota policy/limits that are inherited by all users within the views in this viewbox. | [optional] 
**DefaultViewQuotaPolicy** | Pointer to [**NullableQuotaPolicy**](QuotaPolicy.md) | Specifies an optional default logical quota limit (in bytes) for the Views in this Storage Domain (View Box). (Logical data is when the data is fully hydrated and expanded.) However, this inherited quota can be overwritten at the View level. A new write is not allowed if the Storage Domain (View Box) will exceed the specified quota. However, it takes time for the Cohesity Cluster to calculate the usage across Nodes, so the limit may be exceeded by a small amount. In addition, if the limit is increased or data is removed, there may be delay before the Cohesity Cluster allows more data to be written to the Storage Domain (View Box), as the Cluster is calculating the usage across Nodes. | [optional] 
**DirectArchiveEnabled** | Pointer to **NullableBool** | Specifies whether this viewbox can be used as a staging area while copying a largedataset that can&#39;t fit on the cluster to an external target. The amount of data that can be stored on the viewbox can be specified using &#39;physical_quota&#39;. | [optional] 
**Id** | Pointer to **NullableInt64** | Specifies the Id of the Storage Domain (View Box). | [optional] 
**IsRecommended** | Pointer to **NullableBool** | Recommendation for the view if the template id was passed during query. We say the view is to be recommended, if the dedup and inlinecompression both are same as the template&#39;s property. | [optional] 
**LdapProviderId** | Pointer to **NullableInt64** | When set, the following provides the LDAP provider the view box is mapped to. For any view from this view box, when accessed via NFS the following LDAP provider is looked up for getting Unix IDs of the corresponding user. Similarly, when a view is accessed via SMB and if the AD user&#39;s domain matches with the view box&#39;s AD, the following LDAP provider will be used to lookup Unix IDs for the corresponding AD user. Additionally there is also a mapping between LDAP provider and AD domain that is stored in AD provider config. It will be used if AD is not set on the view box. | [optional] 
**Name** | **NullableString** | Specifies the name of the Storage Domain (View Box). | 
**NisDomainNameVec** | Pointer to **[]string** | Specifies the NIS domain that this view box is mapped to. | [optional] 
**PhysicalQuota** | Pointer to [**NullableQuotaPolicy**](QuotaPolicy.md) | Specifies an optional quota limit (in bytes) for the physical usage of this Storage Domain (View Box). This quota limit defines a physical limit for size of the data that can be physically stored on the Storage Domain (View Box), after the data has been reduced by change block tracking, compression and deduplication. The physical usage is the aggregate sum of the data stored for this Storage Domain (View Box) on all disks in the Cluster. (The usage includes Cloud Tier data and user data.) A new write is not allowed if the Storage Domain (View Box) will exceed the specified quota. However, it takes time for the Cohesity Cluster to calculate the usage across Nodes, so the limit may be exceeded by a small amount. In addition, if the limit is increased or data is removed, there may be a delay before the Cohesity Cluster allows more data to be written to the Storage Domain (View Box), as the Cluster is calculating the usage across Nodes. | [optional] 
**RemovalState** | Pointer to **NullableString** | Specifies the current removal state of the Storage Domain (View Box). &#39;kDontRemove&#39; means the state of object is functional and it is not being removed. &#39;kMarkedForRemoval&#39; means the object is being removed. &#39;kOkToRemove&#39; means the object has been removed on the Cohesity Cluster and if the object is physical, it can be removed from the Cohesity Cluster. | [optional] 
**S3BucketsAllowed** | Pointer to **NullableBool** | Specifies whether creation of a S3 bucket is allowed in this Storage Domain (View Box). When a new S3 bucket creation request arrives, we&#39;ll look at all the View Boxes and the first Storage Domain (View Box) that allows creating S3 buckets in it will be the one where the bucket will be placed. | [optional] 
**SchemaInfoList** | Pointer to [**[]SchemaInfo**](SchemaInfo.md) | Specifies the time series schema info of the view box. | [optional] 
**Stats** | Pointer to [**ViewBoxStats**](ViewBoxStats.md) |  | [optional] 
**StoragePolicy** | Pointer to [**StoragePolicy**](StoragePolicy.md) |  | [optional] 
**TenantIdVec** | Pointer to **[]string** | Optional ids for the tenants that this view box belongs. This must be checked before granting access to users. Unless the cluster enables view box sharing between tenants is allowed, there shall be at most one item in this list. Note that if all tenant may be deleted - such viewboxes must be garbage collected. This is currently done by a background thread in iris. | [optional] 
**TreatFileSyncAsDataSync** | Pointer to **NullableBool** | If &#39;true&#39;, when the Cohesity Cluster is writing to a file, the file modification time is not persisted synchronously during the file write, so the modification time may not be accurate. (Typically the file modification time is off by 30 seconds but it can be longer.) Only set to &#39;false&#39; if your environment requires a very accurate modification time. The default value is &#39;true&#39; which provides the best Cohesity Cluster performance. | [optional] 

## Methods

### NewViewBox

`func NewViewBox(clusterPartitionId NullableInt64, name NullableString, ) *ViewBox`

NewViewBox instantiates a new ViewBox object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewBoxWithDefaults

`func NewViewBoxWithDefaults() *ViewBox`

NewViewBoxWithDefaults instantiates a new ViewBox object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdDomainName

`func (o *ViewBox) GetAdDomainName() string`

GetAdDomainName returns the AdDomainName field if non-nil, zero value otherwise.

### GetAdDomainNameOk

`func (o *ViewBox) GetAdDomainNameOk() (*string, bool)`

GetAdDomainNameOk returns a tuple with the AdDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdDomainName

`func (o *ViewBox) SetAdDomainName(v string)`

SetAdDomainName sets AdDomainName field to given value.

### HasAdDomainName

`func (o *ViewBox) HasAdDomainName() bool`

HasAdDomainName returns a boolean if a field has been set.

### SetAdDomainNameNil

`func (o *ViewBox) SetAdDomainNameNil(b bool)`

 SetAdDomainNameNil sets the value for AdDomainName to be an explicit nil

### UnsetAdDomainName
`func (o *ViewBox) UnsetAdDomainName()`

UnsetAdDomainName ensures that no value is present for AdDomainName, not even an explicit nil
### GetClientSubnetWhiteList

`func (o *ViewBox) GetClientSubnetWhiteList() []Subnet`

GetClientSubnetWhiteList returns the ClientSubnetWhiteList field if non-nil, zero value otherwise.

### GetClientSubnetWhiteListOk

`func (o *ViewBox) GetClientSubnetWhiteListOk() (*[]Subnet, bool)`

GetClientSubnetWhiteListOk returns a tuple with the ClientSubnetWhiteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSubnetWhiteList

`func (o *ViewBox) SetClientSubnetWhiteList(v []Subnet)`

SetClientSubnetWhiteList sets ClientSubnetWhiteList field to given value.

### HasClientSubnetWhiteList

`func (o *ViewBox) HasClientSubnetWhiteList() bool`

HasClientSubnetWhiteList returns a boolean if a field has been set.

### SetClientSubnetWhiteListNil

`func (o *ViewBox) SetClientSubnetWhiteListNil(b bool)`

 SetClientSubnetWhiteListNil sets the value for ClientSubnetWhiteList to be an explicit nil

### UnsetClientSubnetWhiteList
`func (o *ViewBox) UnsetClientSubnetWhiteList()`

UnsetClientSubnetWhiteList ensures that no value is present for ClientSubnetWhiteList, not even an explicit nil
### GetCloudDownWaterfallThresholdPct

`func (o *ViewBox) GetCloudDownWaterfallThresholdPct() int32`

GetCloudDownWaterfallThresholdPct returns the CloudDownWaterfallThresholdPct field if non-nil, zero value otherwise.

### GetCloudDownWaterfallThresholdPctOk

`func (o *ViewBox) GetCloudDownWaterfallThresholdPctOk() (*int32, bool)`

GetCloudDownWaterfallThresholdPctOk returns a tuple with the CloudDownWaterfallThresholdPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudDownWaterfallThresholdPct

`func (o *ViewBox) SetCloudDownWaterfallThresholdPct(v int32)`

SetCloudDownWaterfallThresholdPct sets CloudDownWaterfallThresholdPct field to given value.

### HasCloudDownWaterfallThresholdPct

`func (o *ViewBox) HasCloudDownWaterfallThresholdPct() bool`

HasCloudDownWaterfallThresholdPct returns a boolean if a field has been set.

### SetCloudDownWaterfallThresholdPctNil

`func (o *ViewBox) SetCloudDownWaterfallThresholdPctNil(b bool)`

 SetCloudDownWaterfallThresholdPctNil sets the value for CloudDownWaterfallThresholdPct to be an explicit nil

### UnsetCloudDownWaterfallThresholdPct
`func (o *ViewBox) UnsetCloudDownWaterfallThresholdPct()`

UnsetCloudDownWaterfallThresholdPct ensures that no value is present for CloudDownWaterfallThresholdPct, not even an explicit nil
### GetCloudDownWaterfallThresholdSecs

`func (o *ViewBox) GetCloudDownWaterfallThresholdSecs() int32`

GetCloudDownWaterfallThresholdSecs returns the CloudDownWaterfallThresholdSecs field if non-nil, zero value otherwise.

### GetCloudDownWaterfallThresholdSecsOk

`func (o *ViewBox) GetCloudDownWaterfallThresholdSecsOk() (*int32, bool)`

GetCloudDownWaterfallThresholdSecsOk returns a tuple with the CloudDownWaterfallThresholdSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudDownWaterfallThresholdSecs

`func (o *ViewBox) SetCloudDownWaterfallThresholdSecs(v int32)`

SetCloudDownWaterfallThresholdSecs sets CloudDownWaterfallThresholdSecs field to given value.

### HasCloudDownWaterfallThresholdSecs

`func (o *ViewBox) HasCloudDownWaterfallThresholdSecs() bool`

HasCloudDownWaterfallThresholdSecs returns a boolean if a field has been set.

### SetCloudDownWaterfallThresholdSecsNil

`func (o *ViewBox) SetCloudDownWaterfallThresholdSecsNil(b bool)`

 SetCloudDownWaterfallThresholdSecsNil sets the value for CloudDownWaterfallThresholdSecs to be an explicit nil

### UnsetCloudDownWaterfallThresholdSecs
`func (o *ViewBox) UnsetCloudDownWaterfallThresholdSecs()`

UnsetCloudDownWaterfallThresholdSecs ensures that no value is present for CloudDownWaterfallThresholdSecs, not even an explicit nil
### GetClusterPartitionId

`func (o *ViewBox) GetClusterPartitionId() int64`

GetClusterPartitionId returns the ClusterPartitionId field if non-nil, zero value otherwise.

### GetClusterPartitionIdOk

`func (o *ViewBox) GetClusterPartitionIdOk() (*int64, bool)`

GetClusterPartitionIdOk returns a tuple with the ClusterPartitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterPartitionId

`func (o *ViewBox) SetClusterPartitionId(v int64)`

SetClusterPartitionId sets ClusterPartitionId field to given value.


### SetClusterPartitionIdNil

`func (o *ViewBox) SetClusterPartitionIdNil(b bool)`

 SetClusterPartitionIdNil sets the value for ClusterPartitionId to be an explicit nil

### UnsetClusterPartitionId
`func (o *ViewBox) UnsetClusterPartitionId()`

UnsetClusterPartitionId ensures that no value is present for ClusterPartitionId, not even an explicit nil
### GetClusterPartitionName

`func (o *ViewBox) GetClusterPartitionName() string`

GetClusterPartitionName returns the ClusterPartitionName field if non-nil, zero value otherwise.

### GetClusterPartitionNameOk

`func (o *ViewBox) GetClusterPartitionNameOk() (*string, bool)`

GetClusterPartitionNameOk returns a tuple with the ClusterPartitionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterPartitionName

`func (o *ViewBox) SetClusterPartitionName(v string)`

SetClusterPartitionName sets ClusterPartitionName field to given value.

### HasClusterPartitionName

`func (o *ViewBox) HasClusterPartitionName() bool`

HasClusterPartitionName returns a boolean if a field has been set.

### SetClusterPartitionNameNil

`func (o *ViewBox) SetClusterPartitionNameNil(b bool)`

 SetClusterPartitionNameNil sets the value for ClusterPartitionName to be an explicit nil

### UnsetClusterPartitionName
`func (o *ViewBox) UnsetClusterPartitionName()`

UnsetClusterPartitionName ensures that no value is present for ClusterPartitionName, not even an explicit nil
### GetDefaultUserQuotaPolicy

`func (o *ViewBox) GetDefaultUserQuotaPolicy() QuotaPolicy`

GetDefaultUserQuotaPolicy returns the DefaultUserQuotaPolicy field if non-nil, zero value otherwise.

### GetDefaultUserQuotaPolicyOk

`func (o *ViewBox) GetDefaultUserQuotaPolicyOk() (*QuotaPolicy, bool)`

GetDefaultUserQuotaPolicyOk returns a tuple with the DefaultUserQuotaPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUserQuotaPolicy

`func (o *ViewBox) SetDefaultUserQuotaPolicy(v QuotaPolicy)`

SetDefaultUserQuotaPolicy sets DefaultUserQuotaPolicy field to given value.

### HasDefaultUserQuotaPolicy

`func (o *ViewBox) HasDefaultUserQuotaPolicy() bool`

HasDefaultUserQuotaPolicy returns a boolean if a field has been set.

### SetDefaultUserQuotaPolicyNil

`func (o *ViewBox) SetDefaultUserQuotaPolicyNil(b bool)`

 SetDefaultUserQuotaPolicyNil sets the value for DefaultUserQuotaPolicy to be an explicit nil

### UnsetDefaultUserQuotaPolicy
`func (o *ViewBox) UnsetDefaultUserQuotaPolicy()`

UnsetDefaultUserQuotaPolicy ensures that no value is present for DefaultUserQuotaPolicy, not even an explicit nil
### GetDefaultViewQuotaPolicy

`func (o *ViewBox) GetDefaultViewQuotaPolicy() QuotaPolicy`

GetDefaultViewQuotaPolicy returns the DefaultViewQuotaPolicy field if non-nil, zero value otherwise.

### GetDefaultViewQuotaPolicyOk

`func (o *ViewBox) GetDefaultViewQuotaPolicyOk() (*QuotaPolicy, bool)`

GetDefaultViewQuotaPolicyOk returns a tuple with the DefaultViewQuotaPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultViewQuotaPolicy

`func (o *ViewBox) SetDefaultViewQuotaPolicy(v QuotaPolicy)`

SetDefaultViewQuotaPolicy sets DefaultViewQuotaPolicy field to given value.

### HasDefaultViewQuotaPolicy

`func (o *ViewBox) HasDefaultViewQuotaPolicy() bool`

HasDefaultViewQuotaPolicy returns a boolean if a field has been set.

### SetDefaultViewQuotaPolicyNil

`func (o *ViewBox) SetDefaultViewQuotaPolicyNil(b bool)`

 SetDefaultViewQuotaPolicyNil sets the value for DefaultViewQuotaPolicy to be an explicit nil

### UnsetDefaultViewQuotaPolicy
`func (o *ViewBox) UnsetDefaultViewQuotaPolicy()`

UnsetDefaultViewQuotaPolicy ensures that no value is present for DefaultViewQuotaPolicy, not even an explicit nil
### GetDirectArchiveEnabled

`func (o *ViewBox) GetDirectArchiveEnabled() bool`

GetDirectArchiveEnabled returns the DirectArchiveEnabled field if non-nil, zero value otherwise.

### GetDirectArchiveEnabledOk

`func (o *ViewBox) GetDirectArchiveEnabledOk() (*bool, bool)`

GetDirectArchiveEnabledOk returns a tuple with the DirectArchiveEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectArchiveEnabled

`func (o *ViewBox) SetDirectArchiveEnabled(v bool)`

SetDirectArchiveEnabled sets DirectArchiveEnabled field to given value.

### HasDirectArchiveEnabled

`func (o *ViewBox) HasDirectArchiveEnabled() bool`

HasDirectArchiveEnabled returns a boolean if a field has been set.

### SetDirectArchiveEnabledNil

`func (o *ViewBox) SetDirectArchiveEnabledNil(b bool)`

 SetDirectArchiveEnabledNil sets the value for DirectArchiveEnabled to be an explicit nil

### UnsetDirectArchiveEnabled
`func (o *ViewBox) UnsetDirectArchiveEnabled()`

UnsetDirectArchiveEnabled ensures that no value is present for DirectArchiveEnabled, not even an explicit nil
### GetId

`func (o *ViewBox) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ViewBox) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ViewBox) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ViewBox) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ViewBox) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ViewBox) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIsRecommended

`func (o *ViewBox) GetIsRecommended() bool`

GetIsRecommended returns the IsRecommended field if non-nil, zero value otherwise.

### GetIsRecommendedOk

`func (o *ViewBox) GetIsRecommendedOk() (*bool, bool)`

GetIsRecommendedOk returns a tuple with the IsRecommended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRecommended

`func (o *ViewBox) SetIsRecommended(v bool)`

SetIsRecommended sets IsRecommended field to given value.

### HasIsRecommended

`func (o *ViewBox) HasIsRecommended() bool`

HasIsRecommended returns a boolean if a field has been set.

### SetIsRecommendedNil

`func (o *ViewBox) SetIsRecommendedNil(b bool)`

 SetIsRecommendedNil sets the value for IsRecommended to be an explicit nil

### UnsetIsRecommended
`func (o *ViewBox) UnsetIsRecommended()`

UnsetIsRecommended ensures that no value is present for IsRecommended, not even an explicit nil
### GetLdapProviderId

`func (o *ViewBox) GetLdapProviderId() int64`

GetLdapProviderId returns the LdapProviderId field if non-nil, zero value otherwise.

### GetLdapProviderIdOk

`func (o *ViewBox) GetLdapProviderIdOk() (*int64, bool)`

GetLdapProviderIdOk returns a tuple with the LdapProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapProviderId

`func (o *ViewBox) SetLdapProviderId(v int64)`

SetLdapProviderId sets LdapProviderId field to given value.

### HasLdapProviderId

`func (o *ViewBox) HasLdapProviderId() bool`

HasLdapProviderId returns a boolean if a field has been set.

### SetLdapProviderIdNil

`func (o *ViewBox) SetLdapProviderIdNil(b bool)`

 SetLdapProviderIdNil sets the value for LdapProviderId to be an explicit nil

### UnsetLdapProviderId
`func (o *ViewBox) UnsetLdapProviderId()`

UnsetLdapProviderId ensures that no value is present for LdapProviderId, not even an explicit nil
### GetName

`func (o *ViewBox) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ViewBox) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ViewBox) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *ViewBox) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ViewBox) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNisDomainNameVec

`func (o *ViewBox) GetNisDomainNameVec() []string`

GetNisDomainNameVec returns the NisDomainNameVec field if non-nil, zero value otherwise.

### GetNisDomainNameVecOk

`func (o *ViewBox) GetNisDomainNameVecOk() (*[]string, bool)`

GetNisDomainNameVecOk returns a tuple with the NisDomainNameVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNisDomainNameVec

`func (o *ViewBox) SetNisDomainNameVec(v []string)`

SetNisDomainNameVec sets NisDomainNameVec field to given value.

### HasNisDomainNameVec

`func (o *ViewBox) HasNisDomainNameVec() bool`

HasNisDomainNameVec returns a boolean if a field has been set.

### SetNisDomainNameVecNil

`func (o *ViewBox) SetNisDomainNameVecNil(b bool)`

 SetNisDomainNameVecNil sets the value for NisDomainNameVec to be an explicit nil

### UnsetNisDomainNameVec
`func (o *ViewBox) UnsetNisDomainNameVec()`

UnsetNisDomainNameVec ensures that no value is present for NisDomainNameVec, not even an explicit nil
### GetPhysicalQuota

`func (o *ViewBox) GetPhysicalQuota() QuotaPolicy`

GetPhysicalQuota returns the PhysicalQuota field if non-nil, zero value otherwise.

### GetPhysicalQuotaOk

`func (o *ViewBox) GetPhysicalQuotaOk() (*QuotaPolicy, bool)`

GetPhysicalQuotaOk returns a tuple with the PhysicalQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalQuota

`func (o *ViewBox) SetPhysicalQuota(v QuotaPolicy)`

SetPhysicalQuota sets PhysicalQuota field to given value.

### HasPhysicalQuota

`func (o *ViewBox) HasPhysicalQuota() bool`

HasPhysicalQuota returns a boolean if a field has been set.

### SetPhysicalQuotaNil

`func (o *ViewBox) SetPhysicalQuotaNil(b bool)`

 SetPhysicalQuotaNil sets the value for PhysicalQuota to be an explicit nil

### UnsetPhysicalQuota
`func (o *ViewBox) UnsetPhysicalQuota()`

UnsetPhysicalQuota ensures that no value is present for PhysicalQuota, not even an explicit nil
### GetRemovalState

`func (o *ViewBox) GetRemovalState() string`

GetRemovalState returns the RemovalState field if non-nil, zero value otherwise.

### GetRemovalStateOk

`func (o *ViewBox) GetRemovalStateOk() (*string, bool)`

GetRemovalStateOk returns a tuple with the RemovalState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalState

`func (o *ViewBox) SetRemovalState(v string)`

SetRemovalState sets RemovalState field to given value.

### HasRemovalState

`func (o *ViewBox) HasRemovalState() bool`

HasRemovalState returns a boolean if a field has been set.

### SetRemovalStateNil

`func (o *ViewBox) SetRemovalStateNil(b bool)`

 SetRemovalStateNil sets the value for RemovalState to be an explicit nil

### UnsetRemovalState
`func (o *ViewBox) UnsetRemovalState()`

UnsetRemovalState ensures that no value is present for RemovalState, not even an explicit nil
### GetS3BucketsAllowed

`func (o *ViewBox) GetS3BucketsAllowed() bool`

GetS3BucketsAllowed returns the S3BucketsAllowed field if non-nil, zero value otherwise.

### GetS3BucketsAllowedOk

`func (o *ViewBox) GetS3BucketsAllowedOk() (*bool, bool)`

GetS3BucketsAllowedOk returns a tuple with the S3BucketsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3BucketsAllowed

`func (o *ViewBox) SetS3BucketsAllowed(v bool)`

SetS3BucketsAllowed sets S3BucketsAllowed field to given value.

### HasS3BucketsAllowed

`func (o *ViewBox) HasS3BucketsAllowed() bool`

HasS3BucketsAllowed returns a boolean if a field has been set.

### SetS3BucketsAllowedNil

`func (o *ViewBox) SetS3BucketsAllowedNil(b bool)`

 SetS3BucketsAllowedNil sets the value for S3BucketsAllowed to be an explicit nil

### UnsetS3BucketsAllowed
`func (o *ViewBox) UnsetS3BucketsAllowed()`

UnsetS3BucketsAllowed ensures that no value is present for S3BucketsAllowed, not even an explicit nil
### GetSchemaInfoList

`func (o *ViewBox) GetSchemaInfoList() []SchemaInfo`

GetSchemaInfoList returns the SchemaInfoList field if non-nil, zero value otherwise.

### GetSchemaInfoListOk

`func (o *ViewBox) GetSchemaInfoListOk() (*[]SchemaInfo, bool)`

GetSchemaInfoListOk returns a tuple with the SchemaInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaInfoList

`func (o *ViewBox) SetSchemaInfoList(v []SchemaInfo)`

SetSchemaInfoList sets SchemaInfoList field to given value.

### HasSchemaInfoList

`func (o *ViewBox) HasSchemaInfoList() bool`

HasSchemaInfoList returns a boolean if a field has been set.

### SetSchemaInfoListNil

`func (o *ViewBox) SetSchemaInfoListNil(b bool)`

 SetSchemaInfoListNil sets the value for SchemaInfoList to be an explicit nil

### UnsetSchemaInfoList
`func (o *ViewBox) UnsetSchemaInfoList()`

UnsetSchemaInfoList ensures that no value is present for SchemaInfoList, not even an explicit nil
### GetStats

`func (o *ViewBox) GetStats() ViewBoxStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ViewBox) GetStatsOk() (*ViewBoxStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ViewBox) SetStats(v ViewBoxStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ViewBox) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetStoragePolicy

`func (o *ViewBox) GetStoragePolicy() StoragePolicy`

GetStoragePolicy returns the StoragePolicy field if non-nil, zero value otherwise.

### GetStoragePolicyOk

`func (o *ViewBox) GetStoragePolicyOk() (*StoragePolicy, bool)`

GetStoragePolicyOk returns a tuple with the StoragePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePolicy

`func (o *ViewBox) SetStoragePolicy(v StoragePolicy)`

SetStoragePolicy sets StoragePolicy field to given value.

### HasStoragePolicy

`func (o *ViewBox) HasStoragePolicy() bool`

HasStoragePolicy returns a boolean if a field has been set.

### GetTenantIdVec

`func (o *ViewBox) GetTenantIdVec() []string`

GetTenantIdVec returns the TenantIdVec field if non-nil, zero value otherwise.

### GetTenantIdVecOk

`func (o *ViewBox) GetTenantIdVecOk() (*[]string, bool)`

GetTenantIdVecOk returns a tuple with the TenantIdVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIdVec

`func (o *ViewBox) SetTenantIdVec(v []string)`

SetTenantIdVec sets TenantIdVec field to given value.

### HasTenantIdVec

`func (o *ViewBox) HasTenantIdVec() bool`

HasTenantIdVec returns a boolean if a field has been set.

### SetTenantIdVecNil

`func (o *ViewBox) SetTenantIdVecNil(b bool)`

 SetTenantIdVecNil sets the value for TenantIdVec to be an explicit nil

### UnsetTenantIdVec
`func (o *ViewBox) UnsetTenantIdVec()`

UnsetTenantIdVec ensures that no value is present for TenantIdVec, not even an explicit nil
### GetTreatFileSyncAsDataSync

`func (o *ViewBox) GetTreatFileSyncAsDataSync() bool`

GetTreatFileSyncAsDataSync returns the TreatFileSyncAsDataSync field if non-nil, zero value otherwise.

### GetTreatFileSyncAsDataSyncOk

`func (o *ViewBox) GetTreatFileSyncAsDataSyncOk() (*bool, bool)`

GetTreatFileSyncAsDataSyncOk returns a tuple with the TreatFileSyncAsDataSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreatFileSyncAsDataSync

`func (o *ViewBox) SetTreatFileSyncAsDataSync(v bool)`

SetTreatFileSyncAsDataSync sets TreatFileSyncAsDataSync field to given value.

### HasTreatFileSyncAsDataSync

`func (o *ViewBox) HasTreatFileSyncAsDataSync() bool`

HasTreatFileSyncAsDataSync returns a boolean if a field has been set.

### SetTreatFileSyncAsDataSyncNil

`func (o *ViewBox) SetTreatFileSyncAsDataSyncNil(b bool)`

 SetTreatFileSyncAsDataSyncNil sets the value for TreatFileSyncAsDataSync to be an explicit nil

### UnsetTreatFileSyncAsDataSync
`func (o *ViewBox) UnsetTreatFileSyncAsDataSync()`

UnsetTreatFileSyncAsDataSync ensures that no value is present for TreatFileSyncAsDataSync, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


