# VmwareThrottlingParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewTaskLatencyThresholdMsecs** | Pointer to **NullableInt64** | If the latency of a datastore is above this value, then a new backup task that uses the datastore won&#39;t be started. | [optional] 
**ActiveTaskLatencyThresholdMsecs** | Pointer to **NullableInt64** | If the latency of a datastore is above this value, then an existing backup task that uses the datastore will start getting throttled. | [optional] 
**MaxConcurrentStreams** | Pointer to **NullableInt32** | If this value is &gt; 0 and the number of streams concurrently active on a datastore is equal to it, then any further requests to access the datastore would be denied until the number of active streams reduces. This applies for all the datastores in the specified host. | [optional] 

## Methods

### NewVmwareThrottlingParams

`func NewVmwareThrottlingParams() *VmwareThrottlingParams`

NewVmwareThrottlingParams instantiates a new VmwareThrottlingParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareThrottlingParamsWithDefaults

`func NewVmwareThrottlingParamsWithDefaults() *VmwareThrottlingParams`

NewVmwareThrottlingParamsWithDefaults instantiates a new VmwareThrottlingParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewTaskLatencyThresholdMsecs

`func (o *VmwareThrottlingParams) GetNewTaskLatencyThresholdMsecs() int64`

GetNewTaskLatencyThresholdMsecs returns the NewTaskLatencyThresholdMsecs field if non-nil, zero value otherwise.

### GetNewTaskLatencyThresholdMsecsOk

`func (o *VmwareThrottlingParams) GetNewTaskLatencyThresholdMsecsOk() (*int64, bool)`

GetNewTaskLatencyThresholdMsecsOk returns a tuple with the NewTaskLatencyThresholdMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewTaskLatencyThresholdMsecs

`func (o *VmwareThrottlingParams) SetNewTaskLatencyThresholdMsecs(v int64)`

SetNewTaskLatencyThresholdMsecs sets NewTaskLatencyThresholdMsecs field to given value.

### HasNewTaskLatencyThresholdMsecs

`func (o *VmwareThrottlingParams) HasNewTaskLatencyThresholdMsecs() bool`

HasNewTaskLatencyThresholdMsecs returns a boolean if a field has been set.

### SetNewTaskLatencyThresholdMsecsNil

`func (o *VmwareThrottlingParams) SetNewTaskLatencyThresholdMsecsNil(b bool)`

 SetNewTaskLatencyThresholdMsecsNil sets the value for NewTaskLatencyThresholdMsecs to be an explicit nil

### UnsetNewTaskLatencyThresholdMsecs
`func (o *VmwareThrottlingParams) UnsetNewTaskLatencyThresholdMsecs()`

UnsetNewTaskLatencyThresholdMsecs ensures that no value is present for NewTaskLatencyThresholdMsecs, not even an explicit nil
### GetActiveTaskLatencyThresholdMsecs

`func (o *VmwareThrottlingParams) GetActiveTaskLatencyThresholdMsecs() int64`

GetActiveTaskLatencyThresholdMsecs returns the ActiveTaskLatencyThresholdMsecs field if non-nil, zero value otherwise.

### GetActiveTaskLatencyThresholdMsecsOk

`func (o *VmwareThrottlingParams) GetActiveTaskLatencyThresholdMsecsOk() (*int64, bool)`

GetActiveTaskLatencyThresholdMsecsOk returns a tuple with the ActiveTaskLatencyThresholdMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveTaskLatencyThresholdMsecs

`func (o *VmwareThrottlingParams) SetActiveTaskLatencyThresholdMsecs(v int64)`

SetActiveTaskLatencyThresholdMsecs sets ActiveTaskLatencyThresholdMsecs field to given value.

### HasActiveTaskLatencyThresholdMsecs

`func (o *VmwareThrottlingParams) HasActiveTaskLatencyThresholdMsecs() bool`

HasActiveTaskLatencyThresholdMsecs returns a boolean if a field has been set.

### SetActiveTaskLatencyThresholdMsecsNil

`func (o *VmwareThrottlingParams) SetActiveTaskLatencyThresholdMsecsNil(b bool)`

 SetActiveTaskLatencyThresholdMsecsNil sets the value for ActiveTaskLatencyThresholdMsecs to be an explicit nil

### UnsetActiveTaskLatencyThresholdMsecs
`func (o *VmwareThrottlingParams) UnsetActiveTaskLatencyThresholdMsecs()`

UnsetActiveTaskLatencyThresholdMsecs ensures that no value is present for ActiveTaskLatencyThresholdMsecs, not even an explicit nil
### GetMaxConcurrentStreams

`func (o *VmwareThrottlingParams) GetMaxConcurrentStreams() int32`

GetMaxConcurrentStreams returns the MaxConcurrentStreams field if non-nil, zero value otherwise.

### GetMaxConcurrentStreamsOk

`func (o *VmwareThrottlingParams) GetMaxConcurrentStreamsOk() (*int32, bool)`

GetMaxConcurrentStreamsOk returns a tuple with the MaxConcurrentStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConcurrentStreams

`func (o *VmwareThrottlingParams) SetMaxConcurrentStreams(v int32)`

SetMaxConcurrentStreams sets MaxConcurrentStreams field to given value.

### HasMaxConcurrentStreams

`func (o *VmwareThrottlingParams) HasMaxConcurrentStreams() bool`

HasMaxConcurrentStreams returns a boolean if a field has been set.

### SetMaxConcurrentStreamsNil

`func (o *VmwareThrottlingParams) SetMaxConcurrentStreamsNil(b bool)`

 SetMaxConcurrentStreamsNil sets the value for MaxConcurrentStreams to be an explicit nil

### UnsetMaxConcurrentStreams
`func (o *VmwareThrottlingParams) UnsetMaxConcurrentStreams()`

UnsetMaxConcurrentStreams ensures that no value is present for MaxConcurrentStreams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


