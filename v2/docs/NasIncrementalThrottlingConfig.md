# NasIncrementalThrottlingConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxMetadataFetchPercentage** | Pointer to **NullableInt32** | Specifies the percentage value of maximum concurrent metadata to be fetched during incremental backup of the source. | [optional] 
**MaxReadWritePercentage** | Pointer to **NullableInt32** | Specifies the percentage value of maximum concurrent read/write during incremental backup of the source. | [optional] 

## Methods

### NewNasIncrementalThrottlingConfig

`func NewNasIncrementalThrottlingConfig() *NasIncrementalThrottlingConfig`

NewNasIncrementalThrottlingConfig instantiates a new NasIncrementalThrottlingConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNasIncrementalThrottlingConfigWithDefaults

`func NewNasIncrementalThrottlingConfigWithDefaults() *NasIncrementalThrottlingConfig`

NewNasIncrementalThrottlingConfigWithDefaults instantiates a new NasIncrementalThrottlingConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxMetadataFetchPercentage

`func (o *NasIncrementalThrottlingConfig) GetMaxMetadataFetchPercentage() int32`

GetMaxMetadataFetchPercentage returns the MaxMetadataFetchPercentage field if non-nil, zero value otherwise.

### GetMaxMetadataFetchPercentageOk

`func (o *NasIncrementalThrottlingConfig) GetMaxMetadataFetchPercentageOk() (*int32, bool)`

GetMaxMetadataFetchPercentageOk returns a tuple with the MaxMetadataFetchPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMetadataFetchPercentage

`func (o *NasIncrementalThrottlingConfig) SetMaxMetadataFetchPercentage(v int32)`

SetMaxMetadataFetchPercentage sets MaxMetadataFetchPercentage field to given value.

### HasMaxMetadataFetchPercentage

`func (o *NasIncrementalThrottlingConfig) HasMaxMetadataFetchPercentage() bool`

HasMaxMetadataFetchPercentage returns a boolean if a field has been set.

### SetMaxMetadataFetchPercentageNil

`func (o *NasIncrementalThrottlingConfig) SetMaxMetadataFetchPercentageNil(b bool)`

 SetMaxMetadataFetchPercentageNil sets the value for MaxMetadataFetchPercentage to be an explicit nil

### UnsetMaxMetadataFetchPercentage
`func (o *NasIncrementalThrottlingConfig) UnsetMaxMetadataFetchPercentage()`

UnsetMaxMetadataFetchPercentage ensures that no value is present for MaxMetadataFetchPercentage, not even an explicit nil
### GetMaxReadWritePercentage

`func (o *NasIncrementalThrottlingConfig) GetMaxReadWritePercentage() int32`

GetMaxReadWritePercentage returns the MaxReadWritePercentage field if non-nil, zero value otherwise.

### GetMaxReadWritePercentageOk

`func (o *NasIncrementalThrottlingConfig) GetMaxReadWritePercentageOk() (*int32, bool)`

GetMaxReadWritePercentageOk returns a tuple with the MaxReadWritePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReadWritePercentage

`func (o *NasIncrementalThrottlingConfig) SetMaxReadWritePercentage(v int32)`

SetMaxReadWritePercentage sets MaxReadWritePercentage field to given value.

### HasMaxReadWritePercentage

`func (o *NasIncrementalThrottlingConfig) HasMaxReadWritePercentage() bool`

HasMaxReadWritePercentage returns a boolean if a field has been set.

### SetMaxReadWritePercentageNil

`func (o *NasIncrementalThrottlingConfig) SetMaxReadWritePercentageNil(b bool)`

 SetMaxReadWritePercentageNil sets the value for MaxReadWritePercentage to be an explicit nil

### UnsetMaxReadWritePercentage
`func (o *NasIncrementalThrottlingConfig) UnsetMaxReadWritePercentage()`

UnsetMaxReadWritePercentage ensures that no value is present for MaxReadWritePercentage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


