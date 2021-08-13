# HdfsProtectionGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludePaths** | Pointer to **[]string** | Specifies the paths to be included in the Protection Group. | [optional] 
**ExcludePaths** | Pointer to **[]string** | Specifies the paths to be excluded in the Protection Group. excludePaths will ovrride includePaths. | [optional] 
**Concurrency** | Pointer to **NullableInt32** | Specifies the maximum number of concurrent IO Streams that will be created to exchange data with the cluster. | [optional] 
**BandwidthMBPS** | Pointer to **NullableInt64** | Specifies the maximum network bandwidth that each concurrent IO Stream can use for exchanging data with the cluster. | [optional] 
**HdfsSourceId** | **NullableInt64** | The object ID of the HDFS source for this protection group. | 
**SourceId** | Pointer to **NullableInt64** | Object ID of the Source on which this protection was run . | [optional] [readonly] 
**IndexingPolicy** | Pointer to [**IndexingPolicy**](IndexingPolicy.md) |  | [optional] 
**SourceName** | Pointer to **NullableString** | Specifies the name of the Source on which this protection was run. | [optional] [readonly] 

## Methods

### NewHdfsProtectionGroupParams

`func NewHdfsProtectionGroupParams(hdfsSourceId NullableInt64, ) *HdfsProtectionGroupParams`

NewHdfsProtectionGroupParams instantiates a new HdfsProtectionGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHdfsProtectionGroupParamsWithDefaults

`func NewHdfsProtectionGroupParamsWithDefaults() *HdfsProtectionGroupParams`

NewHdfsProtectionGroupParamsWithDefaults instantiates a new HdfsProtectionGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludePaths

`func (o *HdfsProtectionGroupParams) GetIncludePaths() []string`

GetIncludePaths returns the IncludePaths field if non-nil, zero value otherwise.

### GetIncludePathsOk

`func (o *HdfsProtectionGroupParams) GetIncludePathsOk() (*[]string, bool)`

GetIncludePathsOk returns a tuple with the IncludePaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePaths

`func (o *HdfsProtectionGroupParams) SetIncludePaths(v []string)`

SetIncludePaths sets IncludePaths field to given value.

### HasIncludePaths

`func (o *HdfsProtectionGroupParams) HasIncludePaths() bool`

HasIncludePaths returns a boolean if a field has been set.

### GetExcludePaths

`func (o *HdfsProtectionGroupParams) GetExcludePaths() []string`

GetExcludePaths returns the ExcludePaths field if non-nil, zero value otherwise.

### GetExcludePathsOk

`func (o *HdfsProtectionGroupParams) GetExcludePathsOk() (*[]string, bool)`

GetExcludePathsOk returns a tuple with the ExcludePaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludePaths

`func (o *HdfsProtectionGroupParams) SetExcludePaths(v []string)`

SetExcludePaths sets ExcludePaths field to given value.

### HasExcludePaths

`func (o *HdfsProtectionGroupParams) HasExcludePaths() bool`

HasExcludePaths returns a boolean if a field has been set.

### GetConcurrency

`func (o *HdfsProtectionGroupParams) GetConcurrency() int32`

GetConcurrency returns the Concurrency field if non-nil, zero value otherwise.

### GetConcurrencyOk

`func (o *HdfsProtectionGroupParams) GetConcurrencyOk() (*int32, bool)`

GetConcurrencyOk returns a tuple with the Concurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrency

`func (o *HdfsProtectionGroupParams) SetConcurrency(v int32)`

SetConcurrency sets Concurrency field to given value.

### HasConcurrency

`func (o *HdfsProtectionGroupParams) HasConcurrency() bool`

HasConcurrency returns a boolean if a field has been set.

### SetConcurrencyNil

`func (o *HdfsProtectionGroupParams) SetConcurrencyNil(b bool)`

 SetConcurrencyNil sets the value for Concurrency to be an explicit nil

### UnsetConcurrency
`func (o *HdfsProtectionGroupParams) UnsetConcurrency()`

UnsetConcurrency ensures that no value is present for Concurrency, not even an explicit nil
### GetBandwidthMBPS

`func (o *HdfsProtectionGroupParams) GetBandwidthMBPS() int64`

GetBandwidthMBPS returns the BandwidthMBPS field if non-nil, zero value otherwise.

### GetBandwidthMBPSOk

`func (o *HdfsProtectionGroupParams) GetBandwidthMBPSOk() (*int64, bool)`

GetBandwidthMBPSOk returns a tuple with the BandwidthMBPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthMBPS

`func (o *HdfsProtectionGroupParams) SetBandwidthMBPS(v int64)`

SetBandwidthMBPS sets BandwidthMBPS field to given value.

### HasBandwidthMBPS

`func (o *HdfsProtectionGroupParams) HasBandwidthMBPS() bool`

HasBandwidthMBPS returns a boolean if a field has been set.

### SetBandwidthMBPSNil

`func (o *HdfsProtectionGroupParams) SetBandwidthMBPSNil(b bool)`

 SetBandwidthMBPSNil sets the value for BandwidthMBPS to be an explicit nil

### UnsetBandwidthMBPS
`func (o *HdfsProtectionGroupParams) UnsetBandwidthMBPS()`

UnsetBandwidthMBPS ensures that no value is present for BandwidthMBPS, not even an explicit nil
### GetHdfsSourceId

`func (o *HdfsProtectionGroupParams) GetHdfsSourceId() int64`

GetHdfsSourceId returns the HdfsSourceId field if non-nil, zero value otherwise.

### GetHdfsSourceIdOk

`func (o *HdfsProtectionGroupParams) GetHdfsSourceIdOk() (*int64, bool)`

GetHdfsSourceIdOk returns a tuple with the HdfsSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsSourceId

`func (o *HdfsProtectionGroupParams) SetHdfsSourceId(v int64)`

SetHdfsSourceId sets HdfsSourceId field to given value.


### SetHdfsSourceIdNil

`func (o *HdfsProtectionGroupParams) SetHdfsSourceIdNil(b bool)`

 SetHdfsSourceIdNil sets the value for HdfsSourceId to be an explicit nil

### UnsetHdfsSourceId
`func (o *HdfsProtectionGroupParams) UnsetHdfsSourceId()`

UnsetHdfsSourceId ensures that no value is present for HdfsSourceId, not even an explicit nil
### GetSourceId

`func (o *HdfsProtectionGroupParams) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *HdfsProtectionGroupParams) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *HdfsProtectionGroupParams) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *HdfsProtectionGroupParams) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *HdfsProtectionGroupParams) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *HdfsProtectionGroupParams) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetIndexingPolicy

`func (o *HdfsProtectionGroupParams) GetIndexingPolicy() IndexingPolicy`

GetIndexingPolicy returns the IndexingPolicy field if non-nil, zero value otherwise.

### GetIndexingPolicyOk

`func (o *HdfsProtectionGroupParams) GetIndexingPolicyOk() (*IndexingPolicy, bool)`

GetIndexingPolicyOk returns a tuple with the IndexingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingPolicy

`func (o *HdfsProtectionGroupParams) SetIndexingPolicy(v IndexingPolicy)`

SetIndexingPolicy sets IndexingPolicy field to given value.

### HasIndexingPolicy

`func (o *HdfsProtectionGroupParams) HasIndexingPolicy() bool`

HasIndexingPolicy returns a boolean if a field has been set.

### GetSourceName

`func (o *HdfsProtectionGroupParams) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *HdfsProtectionGroupParams) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *HdfsProtectionGroupParams) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *HdfsProtectionGroupParams) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *HdfsProtectionGroupParams) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *HdfsProtectionGroupParams) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


