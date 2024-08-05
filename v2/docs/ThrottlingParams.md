# ThrottlingParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveTaskLatencyThresholdMsecs** | Pointer to **NullableInt64** | If the latency of a datastore is above this value, then an existing backup task that uses the datastore will start getting throttled. | [optional] 
**DataStoreParams** | Pointer to [**[]DatastoreParams**](DatastoreParams.md) | Specifies datastore specific parameters. | [optional] 
**MaxConcurrentStreams** | Pointer to **NullableInt32** | If this value is &gt; 0 and the number of streams concurrently active on a datastore is equal to it, then any further requests to access the datastore would be denied until the number of active streams reduces. This applies for all the datastores in the specified host. | [optional] 
**NewTaskLatencyThresholdMsecs** | Pointer to **NullableInt64** | If the latency of a datastore is above this value, then a new backup task that uses the datastore won&#39;t be started. | [optional] 

## Methods

### NewThrottlingParams

`func NewThrottlingParams() *ThrottlingParams`

NewThrottlingParams instantiates a new ThrottlingParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThrottlingParamsWithDefaults

`func NewThrottlingParamsWithDefaults() *ThrottlingParams`

NewThrottlingParamsWithDefaults instantiates a new ThrottlingParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveTaskLatencyThresholdMsecs

`func (o *ThrottlingParams) GetActiveTaskLatencyThresholdMsecs() int64`

GetActiveTaskLatencyThresholdMsecs returns the ActiveTaskLatencyThresholdMsecs field if non-nil, zero value otherwise.

### GetActiveTaskLatencyThresholdMsecsOk

`func (o *ThrottlingParams) GetActiveTaskLatencyThresholdMsecsOk() (*int64, bool)`

GetActiveTaskLatencyThresholdMsecsOk returns a tuple with the ActiveTaskLatencyThresholdMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveTaskLatencyThresholdMsecs

`func (o *ThrottlingParams) SetActiveTaskLatencyThresholdMsecs(v int64)`

SetActiveTaskLatencyThresholdMsecs sets ActiveTaskLatencyThresholdMsecs field to given value.

### HasActiveTaskLatencyThresholdMsecs

`func (o *ThrottlingParams) HasActiveTaskLatencyThresholdMsecs() bool`

HasActiveTaskLatencyThresholdMsecs returns a boolean if a field has been set.

### SetActiveTaskLatencyThresholdMsecsNil

`func (o *ThrottlingParams) SetActiveTaskLatencyThresholdMsecsNil(b bool)`

 SetActiveTaskLatencyThresholdMsecsNil sets the value for ActiveTaskLatencyThresholdMsecs to be an explicit nil

### UnsetActiveTaskLatencyThresholdMsecs
`func (o *ThrottlingParams) UnsetActiveTaskLatencyThresholdMsecs()`

UnsetActiveTaskLatencyThresholdMsecs ensures that no value is present for ActiveTaskLatencyThresholdMsecs, not even an explicit nil
### GetDataStoreParams

`func (o *ThrottlingParams) GetDataStoreParams() []DatastoreParams`

GetDataStoreParams returns the DataStoreParams field if non-nil, zero value otherwise.

### GetDataStoreParamsOk

`func (o *ThrottlingParams) GetDataStoreParamsOk() (*[]DatastoreParams, bool)`

GetDataStoreParamsOk returns a tuple with the DataStoreParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataStoreParams

`func (o *ThrottlingParams) SetDataStoreParams(v []DatastoreParams)`

SetDataStoreParams sets DataStoreParams field to given value.

### HasDataStoreParams

`func (o *ThrottlingParams) HasDataStoreParams() bool`

HasDataStoreParams returns a boolean if a field has been set.

### SetDataStoreParamsNil

`func (o *ThrottlingParams) SetDataStoreParamsNil(b bool)`

 SetDataStoreParamsNil sets the value for DataStoreParams to be an explicit nil

### UnsetDataStoreParams
`func (o *ThrottlingParams) UnsetDataStoreParams()`

UnsetDataStoreParams ensures that no value is present for DataStoreParams, not even an explicit nil
### GetMaxConcurrentStreams

`func (o *ThrottlingParams) GetMaxConcurrentStreams() int32`

GetMaxConcurrentStreams returns the MaxConcurrentStreams field if non-nil, zero value otherwise.

### GetMaxConcurrentStreamsOk

`func (o *ThrottlingParams) GetMaxConcurrentStreamsOk() (*int32, bool)`

GetMaxConcurrentStreamsOk returns a tuple with the MaxConcurrentStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConcurrentStreams

`func (o *ThrottlingParams) SetMaxConcurrentStreams(v int32)`

SetMaxConcurrentStreams sets MaxConcurrentStreams field to given value.

### HasMaxConcurrentStreams

`func (o *ThrottlingParams) HasMaxConcurrentStreams() bool`

HasMaxConcurrentStreams returns a boolean if a field has been set.

### SetMaxConcurrentStreamsNil

`func (o *ThrottlingParams) SetMaxConcurrentStreamsNil(b bool)`

 SetMaxConcurrentStreamsNil sets the value for MaxConcurrentStreams to be an explicit nil

### UnsetMaxConcurrentStreams
`func (o *ThrottlingParams) UnsetMaxConcurrentStreams()`

UnsetMaxConcurrentStreams ensures that no value is present for MaxConcurrentStreams, not even an explicit nil
### GetNewTaskLatencyThresholdMsecs

`func (o *ThrottlingParams) GetNewTaskLatencyThresholdMsecs() int64`

GetNewTaskLatencyThresholdMsecs returns the NewTaskLatencyThresholdMsecs field if non-nil, zero value otherwise.

### GetNewTaskLatencyThresholdMsecsOk

`func (o *ThrottlingParams) GetNewTaskLatencyThresholdMsecsOk() (*int64, bool)`

GetNewTaskLatencyThresholdMsecsOk returns a tuple with the NewTaskLatencyThresholdMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewTaskLatencyThresholdMsecs

`func (o *ThrottlingParams) SetNewTaskLatencyThresholdMsecs(v int64)`

SetNewTaskLatencyThresholdMsecs sets NewTaskLatencyThresholdMsecs field to given value.

### HasNewTaskLatencyThresholdMsecs

`func (o *ThrottlingParams) HasNewTaskLatencyThresholdMsecs() bool`

HasNewTaskLatencyThresholdMsecs returns a boolean if a field has been set.

### SetNewTaskLatencyThresholdMsecsNil

`func (o *ThrottlingParams) SetNewTaskLatencyThresholdMsecsNil(b bool)`

 SetNewTaskLatencyThresholdMsecsNil sets the value for NewTaskLatencyThresholdMsecs to be an explicit nil

### UnsetNewTaskLatencyThresholdMsecs
`func (o *ThrottlingParams) UnsetNewTaskLatencyThresholdMsecs()`

UnsetNewTaskLatencyThresholdMsecs ensures that no value is present for NewTaskLatencyThresholdMsecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


