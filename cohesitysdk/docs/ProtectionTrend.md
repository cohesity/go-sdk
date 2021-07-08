# ProtectionTrend

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cancelled** | Pointer to **NullableInt64** | Specifies number of cancelled runs across trends. | [optional] 
**Environment** | Pointer to **NullableString** | Specifies environment. Supported environment types such as &#39;kView&#39;, &#39;kSQL&#39;, &#39;kVMware&#39;, etc. NOTE: &#39;kPuppeteer&#39; refers to Cohesity&#39;s Remote Adapter. &#39;kVMware&#39; indicates the VMware Protection Source environment. &#39;kHyperV&#39; indicates the HyperV Protection Source environment. &#39;kSQL&#39; indicates the SQL Protection Source environment. &#39;kView&#39; indicates the View Protection Source environment. &#39;kPuppeteer&#39; indicates the Cohesity&#39;s Remote Adapter. &#39;kPhysical&#39; indicates the physical Protection Source environment. &#39;kPure&#39; indicates the Pure Storage Protection Source environment. &#39;Nimble&#39; indicates the Nimble Storage Protection Source environment. &#39;kAzure&#39; indicates the Microsoft&#39;s Azure Protection Source environment. &#39;kNetapp&#39; indicates the Netapp Protection Source environment. &#39;kAgent&#39; indicates the Agent Protection Source environment. &#39;kGenericNas&#39; indicates the Generic Network Attached Storage Protection Source environment. &#39;kAcropolis&#39; indicates the Acropolis Protection Source environment. &#39;kPhsicalFiles&#39; indicates the Physical Files Protection Source environment. &#39;kIsilon&#39; indicates the Dell EMC&#39;s Isilon Protection Source environment. &#39;kGPFS&#39; indicates IBM&#39;s GPFS Protection Source environment. &#39;kKVM&#39; indicates the KVM Protection Source environment. &#39;kAWS&#39; indicates the AWS Protection Source environment. &#39;kExchange&#39; indicates the Exchange Protection Source environment. &#39;kHyperVVSS&#39; indicates the HyperV VSS Protection Source environment. &#39;kOracle&#39; indicates the Oracle Protection Source environment. &#39;kGCP&#39; indicates the Google Cloud Platform Protection Source environment. &#39;kFlashBlade&#39; indicates the Flash Blade Protection Source environment. &#39;kAWSNative&#39; indicates the AWS Native Protection Source environment. &#39;kO365&#39; indicates the Office 365 Protection Source environment. &#39;kO365Outlook&#39; indicates Office 365 outlook Protection Source environment. &#39;kHyperFlex&#39; indicates the Hyper Flex Protection Source environment. &#39;kGCPNative&#39; indicates the GCP Native Protection Source environment. &#39;kAzureNative&#39; indicates the Azure Native Protection Source environment. &#39;kKubernetes&#39; indicates a Kubernetes Protection Source environment. &#39;kElastifile&#39; indicates Elastifile Protection Source environment. &#39;kAD&#39; indicates Active Directory Protection Source environment. &#39;kRDSSnapshotManager&#39; indicates AWS RDS Protection Source environment. &#39;kCassandra&#39; indicates Cassandra Protection Source environment. &#39;kMongoDB&#39; indicates MongoDB Protection Source environment. &#39;kCouchbase&#39; indicates Couchbase Protection Source environment. &#39;kHdfs&#39; indicates Hdfs Protection Source environment. &#39;kHive&#39; indicates Hive Protection Source environment. &#39;kHBase&#39; indicates HBase Protection Source environment. | [optional] 
**Failed** | Pointer to **NullableInt64** | Specifies number of failed runs across trends. | [optional] 
**Id** | Pointer to **NullableInt64** | Specifies protected object&#39;s Id. | [optional] 
**Name** | Pointer to **NullableString** | Specifies protected object&#39;s name. | [optional] 
**ParentSourceId** | Pointer to **NullableInt64** | Specifies protected object&#39;s parent id. | [optional] 
**ParentSourceName** | Pointer to **NullableString** | Specifies protected object&#39;s parent name. | [optional] 
**Running** | Pointer to **NullableInt64** | Specifies number of in-progress runs across trends. | [optional] 
**Successful** | Pointer to **NullableInt64** | Specifies number of successful runs across trends. | [optional] 
**Total** | Pointer to **NullableInt64** | Specifies total number of runs across trends. | [optional] 
**Trends** | Pointer to [**[]TrendingData**](TrendingData.md) | Aggregated protection runs information by days/weeks. | [optional] 

## Methods

### NewProtectionTrend

`func NewProtectionTrend() *ProtectionTrend`

NewProtectionTrend instantiates a new ProtectionTrend object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionTrendWithDefaults

`func NewProtectionTrendWithDefaults() *ProtectionTrend`

NewProtectionTrendWithDefaults instantiates a new ProtectionTrend object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancelled

`func (o *ProtectionTrend) GetCancelled() int64`

GetCancelled returns the Cancelled field if non-nil, zero value otherwise.

### GetCancelledOk

`func (o *ProtectionTrend) GetCancelledOk() (*int64, bool)`

GetCancelledOk returns a tuple with the Cancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelled

`func (o *ProtectionTrend) SetCancelled(v int64)`

SetCancelled sets Cancelled field to given value.

### HasCancelled

`func (o *ProtectionTrend) HasCancelled() bool`

HasCancelled returns a boolean if a field has been set.

### SetCancelledNil

`func (o *ProtectionTrend) SetCancelledNil(b bool)`

 SetCancelledNil sets the value for Cancelled to be an explicit nil

### UnsetCancelled
`func (o *ProtectionTrend) UnsetCancelled()`

UnsetCancelled ensures that no value is present for Cancelled, not even an explicit nil
### GetEnvironment

`func (o *ProtectionTrend) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ProtectionTrend) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ProtectionTrend) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ProtectionTrend) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *ProtectionTrend) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *ProtectionTrend) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetFailed

`func (o *ProtectionTrend) GetFailed() int64`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *ProtectionTrend) GetFailedOk() (*int64, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *ProtectionTrend) SetFailed(v int64)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *ProtectionTrend) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### SetFailedNil

`func (o *ProtectionTrend) SetFailedNil(b bool)`

 SetFailedNil sets the value for Failed to be an explicit nil

### UnsetFailed
`func (o *ProtectionTrend) UnsetFailed()`

UnsetFailed ensures that no value is present for Failed, not even an explicit nil
### GetId

`func (o *ProtectionTrend) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProtectionTrend) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProtectionTrend) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ProtectionTrend) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ProtectionTrend) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ProtectionTrend) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *ProtectionTrend) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProtectionTrend) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProtectionTrend) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProtectionTrend) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ProtectionTrend) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProtectionTrend) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetParentSourceId

`func (o *ProtectionTrend) GetParentSourceId() int64`

GetParentSourceId returns the ParentSourceId field if non-nil, zero value otherwise.

### GetParentSourceIdOk

`func (o *ProtectionTrend) GetParentSourceIdOk() (*int64, bool)`

GetParentSourceIdOk returns a tuple with the ParentSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSourceId

`func (o *ProtectionTrend) SetParentSourceId(v int64)`

SetParentSourceId sets ParentSourceId field to given value.

### HasParentSourceId

`func (o *ProtectionTrend) HasParentSourceId() bool`

HasParentSourceId returns a boolean if a field has been set.

### SetParentSourceIdNil

`func (o *ProtectionTrend) SetParentSourceIdNil(b bool)`

 SetParentSourceIdNil sets the value for ParentSourceId to be an explicit nil

### UnsetParentSourceId
`func (o *ProtectionTrend) UnsetParentSourceId()`

UnsetParentSourceId ensures that no value is present for ParentSourceId, not even an explicit nil
### GetParentSourceName

`func (o *ProtectionTrend) GetParentSourceName() string`

GetParentSourceName returns the ParentSourceName field if non-nil, zero value otherwise.

### GetParentSourceNameOk

`func (o *ProtectionTrend) GetParentSourceNameOk() (*string, bool)`

GetParentSourceNameOk returns a tuple with the ParentSourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSourceName

`func (o *ProtectionTrend) SetParentSourceName(v string)`

SetParentSourceName sets ParentSourceName field to given value.

### HasParentSourceName

`func (o *ProtectionTrend) HasParentSourceName() bool`

HasParentSourceName returns a boolean if a field has been set.

### SetParentSourceNameNil

`func (o *ProtectionTrend) SetParentSourceNameNil(b bool)`

 SetParentSourceNameNil sets the value for ParentSourceName to be an explicit nil

### UnsetParentSourceName
`func (o *ProtectionTrend) UnsetParentSourceName()`

UnsetParentSourceName ensures that no value is present for ParentSourceName, not even an explicit nil
### GetRunning

`func (o *ProtectionTrend) GetRunning() int64`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *ProtectionTrend) GetRunningOk() (*int64, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *ProtectionTrend) SetRunning(v int64)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *ProtectionTrend) HasRunning() bool`

HasRunning returns a boolean if a field has been set.

### SetRunningNil

`func (o *ProtectionTrend) SetRunningNil(b bool)`

 SetRunningNil sets the value for Running to be an explicit nil

### UnsetRunning
`func (o *ProtectionTrend) UnsetRunning()`

UnsetRunning ensures that no value is present for Running, not even an explicit nil
### GetSuccessful

`func (o *ProtectionTrend) GetSuccessful() int64`

GetSuccessful returns the Successful field if non-nil, zero value otherwise.

### GetSuccessfulOk

`func (o *ProtectionTrend) GetSuccessfulOk() (*int64, bool)`

GetSuccessfulOk returns a tuple with the Successful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessful

`func (o *ProtectionTrend) SetSuccessful(v int64)`

SetSuccessful sets Successful field to given value.

### HasSuccessful

`func (o *ProtectionTrend) HasSuccessful() bool`

HasSuccessful returns a boolean if a field has been set.

### SetSuccessfulNil

`func (o *ProtectionTrend) SetSuccessfulNil(b bool)`

 SetSuccessfulNil sets the value for Successful to be an explicit nil

### UnsetSuccessful
`func (o *ProtectionTrend) UnsetSuccessful()`

UnsetSuccessful ensures that no value is present for Successful, not even an explicit nil
### GetTotal

`func (o *ProtectionTrend) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ProtectionTrend) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ProtectionTrend) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ProtectionTrend) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *ProtectionTrend) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *ProtectionTrend) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil
### GetTrends

`func (o *ProtectionTrend) GetTrends() []TrendingData`

GetTrends returns the Trends field if non-nil, zero value otherwise.

### GetTrendsOk

`func (o *ProtectionTrend) GetTrendsOk() (*[]TrendingData, bool)`

GetTrendsOk returns a tuple with the Trends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrends

`func (o *ProtectionTrend) SetTrends(v []TrendingData)`

SetTrends sets Trends field to given value.

### HasTrends

`func (o *ProtectionTrend) HasTrends() bool`

HasTrends returns a boolean if a field has been set.

### SetTrendsNil

`func (o *ProtectionTrend) SetTrendsNil(b bool)`

 SetTrendsNil sets the value for Trends to be an explicit nil

### UnsetTrends
`func (o *ProtectionTrend) UnsetTrends()`

UnsetTrends ensures that no value is present for Trends, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


