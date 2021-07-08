# EntitySchemaProtoGranularity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RollupIntervalSecs** | Pointer to **NullableInt32** | Defines the rollup interval or a bucket size. All data points within one time bucket are rolled up to one summary data point using the defined rollup function. For example, say, raw metric is published at ~30 secs granularity. To generate a hourly or a daily summary time series, client can define rolled up metrics having interval 3600 secs and 86400 secs respectively. | [optional] 
**TimeToLiveSecs** | Pointer to **NullableInt64** | Defines the duration for which the rolled up data is to be stored. Once the lifespan has elapsed, expired data is garbage collected. | [optional] 

## Methods

### NewEntitySchemaProtoGranularity

`func NewEntitySchemaProtoGranularity() *EntitySchemaProtoGranularity`

NewEntitySchemaProtoGranularity instantiates a new EntitySchemaProtoGranularity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitySchemaProtoGranularityWithDefaults

`func NewEntitySchemaProtoGranularityWithDefaults() *EntitySchemaProtoGranularity`

NewEntitySchemaProtoGranularityWithDefaults instantiates a new EntitySchemaProtoGranularity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRollupIntervalSecs

`func (o *EntitySchemaProtoGranularity) GetRollupIntervalSecs() int32`

GetRollupIntervalSecs returns the RollupIntervalSecs field if non-nil, zero value otherwise.

### GetRollupIntervalSecsOk

`func (o *EntitySchemaProtoGranularity) GetRollupIntervalSecsOk() (*int32, bool)`

GetRollupIntervalSecsOk returns a tuple with the RollupIntervalSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollupIntervalSecs

`func (o *EntitySchemaProtoGranularity) SetRollupIntervalSecs(v int32)`

SetRollupIntervalSecs sets RollupIntervalSecs field to given value.

### HasRollupIntervalSecs

`func (o *EntitySchemaProtoGranularity) HasRollupIntervalSecs() bool`

HasRollupIntervalSecs returns a boolean if a field has been set.

### SetRollupIntervalSecsNil

`func (o *EntitySchemaProtoGranularity) SetRollupIntervalSecsNil(b bool)`

 SetRollupIntervalSecsNil sets the value for RollupIntervalSecs to be an explicit nil

### UnsetRollupIntervalSecs
`func (o *EntitySchemaProtoGranularity) UnsetRollupIntervalSecs()`

UnsetRollupIntervalSecs ensures that no value is present for RollupIntervalSecs, not even an explicit nil
### GetTimeToLiveSecs

`func (o *EntitySchemaProtoGranularity) GetTimeToLiveSecs() int64`

GetTimeToLiveSecs returns the TimeToLiveSecs field if non-nil, zero value otherwise.

### GetTimeToLiveSecsOk

`func (o *EntitySchemaProtoGranularity) GetTimeToLiveSecsOk() (*int64, bool)`

GetTimeToLiveSecsOk returns a tuple with the TimeToLiveSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToLiveSecs

`func (o *EntitySchemaProtoGranularity) SetTimeToLiveSecs(v int64)`

SetTimeToLiveSecs sets TimeToLiveSecs field to given value.

### HasTimeToLiveSecs

`func (o *EntitySchemaProtoGranularity) HasTimeToLiveSecs() bool`

HasTimeToLiveSecs returns a boolean if a field has been set.

### SetTimeToLiveSecsNil

`func (o *EntitySchemaProtoGranularity) SetTimeToLiveSecsNil(b bool)`

 SetTimeToLiveSecsNil sets the value for TimeToLiveSecs to be an explicit nil

### UnsetTimeToLiveSecs
`func (o *EntitySchemaProtoGranularity) UnsetTimeToLiveSecs()`

UnsetTimeToLiveSecs ensures that no value is present for TimeToLiveSecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


