# AggregatedSubtreeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | Pointer to **NullableString** | Specifies the environment such as &#39;kSQL&#39; or &#39;kVMware&#39;, where the Protection Source exists. Supported environment types such as &#39;kView&#39;, &#39;kSQL&#39;, &#39;kVMware&#39;, etc. NOTE: &#39;kPuppeteer&#39; refers to Cohesity&#39;s Remote Adapter. &#39;kVMware&#39; indicates the VMware Protection Source environment. &#39;kHyperV&#39; indicates the HyperV Protection Source environment. &#39;kSQL&#39; indicates the SQL Protection Source environment. &#39;kView&#39; indicates the View Protection Source environment. &#39;kPuppeteer&#39; indicates the Cohesity&#39;s Remote Adapter. &#39;kPhysical&#39; indicates the physical Protection Source environment. &#39;kPure&#39; indicates the Pure Storage Protection Source environment. &#39;Nimble&#39; indicates the Nimble Storage Protection Source environment. &#39;kAzure&#39; indicates the Microsoft&#39;s Azure Protection Source environment. &#39;kNetapp&#39; indicates the Netapp Protection Source environment. &#39;kAgent&#39; indicates the Agent Protection Source environment. &#39;kGenericNas&#39; indicates the Generic Network Attached Storage Protection Source environment. &#39;kAcropolis&#39; indicates the Acropolis Protection Source environment. &#39;kPhsicalFiles&#39; indicates the Physical Files Protection Source environment. &#39;kIsilon&#39; indicates the Dell EMC&#39;s Isilon Protection Source environment. &#39;kGPFS&#39; indicates IBM&#39;s GPFS Protection Source environment. &#39;kKVM&#39; indicates the KVM Protection Source environment. &#39;kAWS&#39; indicates the AWS Protection Source environment. &#39;kExchange&#39; indicates the Exchange Protection Source environment. &#39;kHyperVVSS&#39; indicates the HyperV VSS Protection Source environment. &#39;kOracle&#39; indicates the Oracle Protection Source environment. &#39;kGCP&#39; indicates the Google Cloud Platform Protection Source environment. &#39;kFlashBlade&#39; indicates the Flash Blade Protection Source environment. &#39;kAWSNative&#39; indicates the AWS Native Protection Source environment. &#39;kO365&#39; indicates the Office 365 Protection Source environment. &#39;kO365Outlook&#39; indicates Office 365 outlook Protection Source environment. &#39;kHyperFlex&#39; indicates the Hyper Flex Protection Source environment. &#39;kGCPNative&#39; indicates the GCP Native Protection Source environment. &#39;kAzureNative&#39; indicates the Azure Native Protection Source environment. &#39;kKubernetes&#39; indicates a Kubernetes Protection Source environment. &#39;kElastifile&#39; indicates Elastifile Protection Source environment. &#39;kAD&#39; indicates Active Directory Protection Source environment. &#39;kRDSSnapshotManager&#39; indicates AWS RDS Protection Source environment. &#39;kCassandra&#39; indicates Cassandra Protection Source environment. &#39;kMongoDB&#39; indicates MongoDB Protection Source environment. &#39;kCouchbase&#39; indicates Couchbase Protection Source environment. &#39;kHdfs&#39; indicates Hdfs Protection Source environment. &#39;kHive&#39; indicates Hive Protection Source environment. &#39;kHBase&#39; indicates HBase Protection Source environment. | [optional] 
**LeavesCount** | Pointer to **NullableInt64** | Specifies the number of leaf nodes under the subtree of this node. | [optional] 
**TotalLogicalSize** | Pointer to **NullableInt64** | Specifies the total logical size of the data under the subtree of this node. | [optional] 

## Methods

### NewAggregatedSubtreeInfo

`func NewAggregatedSubtreeInfo() *AggregatedSubtreeInfo`

NewAggregatedSubtreeInfo instantiates a new AggregatedSubtreeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregatedSubtreeInfoWithDefaults

`func NewAggregatedSubtreeInfoWithDefaults() *AggregatedSubtreeInfo`

NewAggregatedSubtreeInfoWithDefaults instantiates a new AggregatedSubtreeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *AggregatedSubtreeInfo) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *AggregatedSubtreeInfo) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *AggregatedSubtreeInfo) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *AggregatedSubtreeInfo) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *AggregatedSubtreeInfo) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *AggregatedSubtreeInfo) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetLeavesCount

`func (o *AggregatedSubtreeInfo) GetLeavesCount() int64`

GetLeavesCount returns the LeavesCount field if non-nil, zero value otherwise.

### GetLeavesCountOk

`func (o *AggregatedSubtreeInfo) GetLeavesCountOk() (*int64, bool)`

GetLeavesCountOk returns a tuple with the LeavesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeavesCount

`func (o *AggregatedSubtreeInfo) SetLeavesCount(v int64)`

SetLeavesCount sets LeavesCount field to given value.

### HasLeavesCount

`func (o *AggregatedSubtreeInfo) HasLeavesCount() bool`

HasLeavesCount returns a boolean if a field has been set.

### SetLeavesCountNil

`func (o *AggregatedSubtreeInfo) SetLeavesCountNil(b bool)`

 SetLeavesCountNil sets the value for LeavesCount to be an explicit nil

### UnsetLeavesCount
`func (o *AggregatedSubtreeInfo) UnsetLeavesCount()`

UnsetLeavesCount ensures that no value is present for LeavesCount, not even an explicit nil
### GetTotalLogicalSize

`func (o *AggregatedSubtreeInfo) GetTotalLogicalSize() int64`

GetTotalLogicalSize returns the TotalLogicalSize field if non-nil, zero value otherwise.

### GetTotalLogicalSizeOk

`func (o *AggregatedSubtreeInfo) GetTotalLogicalSizeOk() (*int64, bool)`

GetTotalLogicalSizeOk returns a tuple with the TotalLogicalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLogicalSize

`func (o *AggregatedSubtreeInfo) SetTotalLogicalSize(v int64)`

SetTotalLogicalSize sets TotalLogicalSize field to given value.

### HasTotalLogicalSize

`func (o *AggregatedSubtreeInfo) HasTotalLogicalSize() bool`

HasTotalLogicalSize returns a boolean if a field has been set.

### SetTotalLogicalSizeNil

`func (o *AggregatedSubtreeInfo) SetTotalLogicalSizeNil(b bool)`

 SetTotalLogicalSizeNil sets the value for TotalLogicalSize to be an explicit nil

### UnsetTotalLogicalSize
`func (o *AggregatedSubtreeInfo) UnsetTotalLogicalSize()`

UnsetTotalLogicalSize ensures that no value is present for TotalLogicalSize, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


