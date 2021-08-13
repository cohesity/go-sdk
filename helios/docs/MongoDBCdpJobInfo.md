# MongoDBCdpJobInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LatestRecoveryPointInTimeUsecs** | Pointer to **NullableInt64** | Specifies the latest available recovery point timestamp (in microseconds from epoch) | [optional] 

## Methods

### NewMongoDBCdpJobInfo

`func NewMongoDBCdpJobInfo() *MongoDBCdpJobInfo`

NewMongoDBCdpJobInfo instantiates a new MongoDBCdpJobInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMongoDBCdpJobInfoWithDefaults

`func NewMongoDBCdpJobInfoWithDefaults() *MongoDBCdpJobInfo`

NewMongoDBCdpJobInfoWithDefaults instantiates a new MongoDBCdpJobInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLatestRecoveryPointInTimeUsecs

`func (o *MongoDBCdpJobInfo) GetLatestRecoveryPointInTimeUsecs() int64`

GetLatestRecoveryPointInTimeUsecs returns the LatestRecoveryPointInTimeUsecs field if non-nil, zero value otherwise.

### GetLatestRecoveryPointInTimeUsecsOk

`func (o *MongoDBCdpJobInfo) GetLatestRecoveryPointInTimeUsecsOk() (*int64, bool)`

GetLatestRecoveryPointInTimeUsecsOk returns a tuple with the LatestRecoveryPointInTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestRecoveryPointInTimeUsecs

`func (o *MongoDBCdpJobInfo) SetLatestRecoveryPointInTimeUsecs(v int64)`

SetLatestRecoveryPointInTimeUsecs sets LatestRecoveryPointInTimeUsecs field to given value.

### HasLatestRecoveryPointInTimeUsecs

`func (o *MongoDBCdpJobInfo) HasLatestRecoveryPointInTimeUsecs() bool`

HasLatestRecoveryPointInTimeUsecs returns a boolean if a field has been set.

### SetLatestRecoveryPointInTimeUsecsNil

`func (o *MongoDBCdpJobInfo) SetLatestRecoveryPointInTimeUsecsNil(b bool)`

 SetLatestRecoveryPointInTimeUsecsNil sets the value for LatestRecoveryPointInTimeUsecs to be an explicit nil

### UnsetLatestRecoveryPointInTimeUsecs
`func (o *MongoDBCdpJobInfo) UnsetLatestRecoveryPointInTimeUsecs()`

UnsetLatestRecoveryPointInTimeUsecs ensures that no value is present for LatestRecoveryPointInTimeUsecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


