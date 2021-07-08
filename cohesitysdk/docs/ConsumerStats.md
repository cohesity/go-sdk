# ConsumerStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsumerType** | Pointer to **NullableString** | Specifies the type of the consumer. Type of the consumer can be one of the following three,  &#39;kViews&#39;, indicates the stats info of Views used per organization (tenant) per view box (storage domain). &#39;kProtectionRuns&#39;, indicates the stats info of Protection Runs used per organization (tenant) per view box (storage domain). &#39;kReplicationRuns&#39;, indicates the stats info of Replication In used per organization (tenant) per view box (storage domain). &#39;kViewProtectionRuns&#39;, indicates the stats info of View Protection Runs used per organization (tenant) per view box (storage domain). | [optional] 
**GroupList** | Pointer to [**[]StatsGroup**](StatsGroup.md) | Specifies a list of groups associated to this consumer. | [optional] 
**Id** | Pointer to **NullableInt64** | Specifies the id of the consumer. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the consumer. | [optional] 
**ProtectionEnvironment** | Pointer to **NullableString** | Specifies the source environment of the protection job. Supported environment types such as &#39;kView&#39;, &#39;kSQL&#39;, &#39;kVMware&#39;, etc. NOTE: &#39;kPuppeteer&#39; refers to Cohesity&#39;s Remote Adapter. &#39;kVMware&#39; indicates the VMware Protection Source environment. &#39;kHyperV&#39; indicates the HyperV Protection Source environment. &#39;kSQL&#39; indicates the SQL Protection Source environment. &#39;kView&#39; indicates the View Protection Source environment. &#39;kPuppeteer&#39; indicates the Cohesity&#39;s Remote Adapter. &#39;kPhysical&#39; indicates the physical Protection Source environment. &#39;kPure&#39; indicates the Pure Storage Protection Source environment. &#39;Nimble&#39; indicates the Nimble Storage Protection Source environment. &#39;kAzure&#39; indicates the Microsoft&#39;s Azure Protection Source environment. &#39;kNetapp&#39; indicates the Netapp Protection Source environment. &#39;kAgent&#39; indicates the Agent Protection Source environment. &#39;kGenericNas&#39; indicates the Generic Network Attached Storage Protection Source environment. &#39;kAcropolis&#39; indicates the Acropolis Protection Source environment. &#39;kPhsicalFiles&#39; indicates the Physical Files Protection Source environment. &#39;kIsilon&#39; indicates the Dell EMC&#39;s Isilon Protection Source environment. &#39;kGPFS&#39; indicates IBM&#39;s GPFS Protection Source environment. &#39;kKVM&#39; indicates the KVM Protection Source environment. &#39;kAWS&#39; indicates the AWS Protection Source environment. &#39;kExchange&#39; indicates the Exchange Protection Source environment. &#39;kHyperVVSS&#39; indicates the HyperV VSS Protection Source environment. &#39;kOracle&#39; indicates the Oracle Protection Source environment. &#39;kGCP&#39; indicates the Google Cloud Platform Protection Source environment. &#39;kFlashBlade&#39; indicates the Flash Blade Protection Source environment. &#39;kAWSNative&#39; indicates the AWS Native Protection Source environment. &#39;kO365&#39; indicates the Office 365 Protection Source environment. &#39;kO365Outlook&#39; indicates Office 365 outlook Protection Source environment. &#39;kHyperFlex&#39; indicates the Hyper Flex Protection Source environment. &#39;kGCPNative&#39; indicates the GCP Native Protection Source environment. &#39;kAzureNative&#39; indicates the Azure Native Protection Source environment. &#39;kKubernetes&#39; indicates a Kubernetes Protection Source environment. &#39;kElastifile&#39; indicates Elastifile Protection Source environment. &#39;kAD&#39; indicates Active Directory Protection Source environment. &#39;kRDSSnapshotManager&#39; indicates AWS RDS Protection Source environment. &#39;kCassandra&#39; indicates Cassandra Protection Source environment. &#39;kMongoDB&#39; indicates MongoDB Protection Source environment. &#39;kCouchbase&#39; indicates Couchbase Protection Source environment. &#39;kHdfs&#39; indicates Hdfs Protection Source environment. &#39;kHive&#39; indicates Hive Protection Source environment. &#39;kHBase&#39; indicates HBase Protection Source environment. | [optional] 
**ProtectionPolicyName** | Pointer to **NullableString** | Specifies the name of the protection policy for &#39;kProtectionRuns&#39; and &#39;kReplicationRuns&#39; consumer. | [optional] 
**QuotaHardLimitBytes** | Pointer to **NullableInt64** | Specifies the hard limit of logical quota of the consumer. This field will be returned only if consumer type is view. | [optional] 
**SchemaInfoList** | Pointer to [**[]UsageSchemaInfo**](UsageSchemaInfo.md) | Specifies a list of schemaInfos of the consumer. | [optional] 
**Stats** | Pointer to [**DataUsageStats**](DataUsageStats.md) |  | [optional] 

## Methods

### NewConsumerStats

`func NewConsumerStats() *ConsumerStats`

NewConsumerStats instantiates a new ConsumerStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerStatsWithDefaults

`func NewConsumerStatsWithDefaults() *ConsumerStats`

NewConsumerStatsWithDefaults instantiates a new ConsumerStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsumerType

`func (o *ConsumerStats) GetConsumerType() string`

GetConsumerType returns the ConsumerType field if non-nil, zero value otherwise.

### GetConsumerTypeOk

`func (o *ConsumerStats) GetConsumerTypeOk() (*string, bool)`

GetConsumerTypeOk returns a tuple with the ConsumerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerType

`func (o *ConsumerStats) SetConsumerType(v string)`

SetConsumerType sets ConsumerType field to given value.

### HasConsumerType

`func (o *ConsumerStats) HasConsumerType() bool`

HasConsumerType returns a boolean if a field has been set.

### SetConsumerTypeNil

`func (o *ConsumerStats) SetConsumerTypeNil(b bool)`

 SetConsumerTypeNil sets the value for ConsumerType to be an explicit nil

### UnsetConsumerType
`func (o *ConsumerStats) UnsetConsumerType()`

UnsetConsumerType ensures that no value is present for ConsumerType, not even an explicit nil
### GetGroupList

`func (o *ConsumerStats) GetGroupList() []StatsGroup`

GetGroupList returns the GroupList field if non-nil, zero value otherwise.

### GetGroupListOk

`func (o *ConsumerStats) GetGroupListOk() (*[]StatsGroup, bool)`

GetGroupListOk returns a tuple with the GroupList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupList

`func (o *ConsumerStats) SetGroupList(v []StatsGroup)`

SetGroupList sets GroupList field to given value.

### HasGroupList

`func (o *ConsumerStats) HasGroupList() bool`

HasGroupList returns a boolean if a field has been set.

### SetGroupListNil

`func (o *ConsumerStats) SetGroupListNil(b bool)`

 SetGroupListNil sets the value for GroupList to be an explicit nil

### UnsetGroupList
`func (o *ConsumerStats) UnsetGroupList()`

UnsetGroupList ensures that no value is present for GroupList, not even an explicit nil
### GetId

`func (o *ConsumerStats) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConsumerStats) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConsumerStats) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ConsumerStats) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ConsumerStats) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ConsumerStats) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *ConsumerStats) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConsumerStats) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConsumerStats) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConsumerStats) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ConsumerStats) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ConsumerStats) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProtectionEnvironment

`func (o *ConsumerStats) GetProtectionEnvironment() string`

GetProtectionEnvironment returns the ProtectionEnvironment field if non-nil, zero value otherwise.

### GetProtectionEnvironmentOk

`func (o *ConsumerStats) GetProtectionEnvironmentOk() (*string, bool)`

GetProtectionEnvironmentOk returns a tuple with the ProtectionEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionEnvironment

`func (o *ConsumerStats) SetProtectionEnvironment(v string)`

SetProtectionEnvironment sets ProtectionEnvironment field to given value.

### HasProtectionEnvironment

`func (o *ConsumerStats) HasProtectionEnvironment() bool`

HasProtectionEnvironment returns a boolean if a field has been set.

### SetProtectionEnvironmentNil

`func (o *ConsumerStats) SetProtectionEnvironmentNil(b bool)`

 SetProtectionEnvironmentNil sets the value for ProtectionEnvironment to be an explicit nil

### UnsetProtectionEnvironment
`func (o *ConsumerStats) UnsetProtectionEnvironment()`

UnsetProtectionEnvironment ensures that no value is present for ProtectionEnvironment, not even an explicit nil
### GetProtectionPolicyName

`func (o *ConsumerStats) GetProtectionPolicyName() string`

GetProtectionPolicyName returns the ProtectionPolicyName field if non-nil, zero value otherwise.

### GetProtectionPolicyNameOk

`func (o *ConsumerStats) GetProtectionPolicyNameOk() (*string, bool)`

GetProtectionPolicyNameOk returns a tuple with the ProtectionPolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionPolicyName

`func (o *ConsumerStats) SetProtectionPolicyName(v string)`

SetProtectionPolicyName sets ProtectionPolicyName field to given value.

### HasProtectionPolicyName

`func (o *ConsumerStats) HasProtectionPolicyName() bool`

HasProtectionPolicyName returns a boolean if a field has been set.

### SetProtectionPolicyNameNil

`func (o *ConsumerStats) SetProtectionPolicyNameNil(b bool)`

 SetProtectionPolicyNameNil sets the value for ProtectionPolicyName to be an explicit nil

### UnsetProtectionPolicyName
`func (o *ConsumerStats) UnsetProtectionPolicyName()`

UnsetProtectionPolicyName ensures that no value is present for ProtectionPolicyName, not even an explicit nil
### GetQuotaHardLimitBytes

`func (o *ConsumerStats) GetQuotaHardLimitBytes() int64`

GetQuotaHardLimitBytes returns the QuotaHardLimitBytes field if non-nil, zero value otherwise.

### GetQuotaHardLimitBytesOk

`func (o *ConsumerStats) GetQuotaHardLimitBytesOk() (*int64, bool)`

GetQuotaHardLimitBytesOk returns a tuple with the QuotaHardLimitBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaHardLimitBytes

`func (o *ConsumerStats) SetQuotaHardLimitBytes(v int64)`

SetQuotaHardLimitBytes sets QuotaHardLimitBytes field to given value.

### HasQuotaHardLimitBytes

`func (o *ConsumerStats) HasQuotaHardLimitBytes() bool`

HasQuotaHardLimitBytes returns a boolean if a field has been set.

### SetQuotaHardLimitBytesNil

`func (o *ConsumerStats) SetQuotaHardLimitBytesNil(b bool)`

 SetQuotaHardLimitBytesNil sets the value for QuotaHardLimitBytes to be an explicit nil

### UnsetQuotaHardLimitBytes
`func (o *ConsumerStats) UnsetQuotaHardLimitBytes()`

UnsetQuotaHardLimitBytes ensures that no value is present for QuotaHardLimitBytes, not even an explicit nil
### GetSchemaInfoList

`func (o *ConsumerStats) GetSchemaInfoList() []UsageSchemaInfo`

GetSchemaInfoList returns the SchemaInfoList field if non-nil, zero value otherwise.

### GetSchemaInfoListOk

`func (o *ConsumerStats) GetSchemaInfoListOk() (*[]UsageSchemaInfo, bool)`

GetSchemaInfoListOk returns a tuple with the SchemaInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaInfoList

`func (o *ConsumerStats) SetSchemaInfoList(v []UsageSchemaInfo)`

SetSchemaInfoList sets SchemaInfoList field to given value.

### HasSchemaInfoList

`func (o *ConsumerStats) HasSchemaInfoList() bool`

HasSchemaInfoList returns a boolean if a field has been set.

### SetSchemaInfoListNil

`func (o *ConsumerStats) SetSchemaInfoListNil(b bool)`

 SetSchemaInfoListNil sets the value for SchemaInfoList to be an explicit nil

### UnsetSchemaInfoList
`func (o *ConsumerStats) UnsetSchemaInfoList()`

UnsetSchemaInfoList ensures that no value is present for SchemaInfoList, not even an explicit nil
### GetStats

`func (o *ConsumerStats) GetStats() DataUsageStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ConsumerStats) GetStatsOk() (*DataUsageStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ConsumerStats) SetStats(v DataUsageStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ConsumerStats) HasStats() bool`

HasStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


