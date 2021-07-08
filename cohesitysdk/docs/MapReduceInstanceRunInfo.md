# MapReduceInstanceRunInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTime** | Pointer to **NullableInt64** | Time when map redcue job completed. | [optional] 
**ErrorMessage** | Pointer to **NullableString** | If this run failed, then error message for failure. | [optional] 
**ExecutionStartTimeUsecs** | Pointer to **NullableInt64** | Time (in usecs) when job was picked up for execution. | [optional] 
**FilesProcessed** | Pointer to **NullableInt32** | Number of files processed in this run. | [optional] 
**MapDoneTimeUsecs** | Pointer to **NullableInt64** | Time (in usecs) when map tasks were done. | [optional] 
**MapInputBytes** | Pointer to **NullableInt64** | Total size of data processed by this run in bytes. | [optional] 
**MappersSpawned** | Pointer to **NullableInt32** | Number of mappers spawned till now. | [optional] 
**NumMapOutputs** | Pointer to **NullableInt64** | Number of outputs from mappers. | [optional] 
**NumReduceOutputs** | Pointer to **NullableInt64** | Number of outputs from reducers. | [optional] 
**PercentageCompletion** | Pointer to **NullableFloat32** | Percentage completion of this run so far. | [optional] 
**PercentageMapperCompletion** | Pointer to **NullableFloat32** | Percentage of mapper phase completed. | [optional] 
**PercentageReducerCompletion** | Pointer to **NullableFloat32** | Percentage of reducer phase completed. | [optional] 
**ReducersSpawned** | Pointer to **NullableInt32** | Number of reducers spawned till now. | [optional] 
**RemainingTimeMins** | Pointer to **NullableInt32** | Expected remaining time in minutes for completion of this run. | [optional] 
**StartTime** | Pointer to **NullableInt64** | Time when map reduce job was started by user. | [optional] 
**Status** | Pointer to **NullableInt32** | Status of this run. | [optional] 
**TotalNumMappers** | Pointer to **NullableInt32** | Total number of mappers to be spawned. | [optional] 
**TotalNumReducers** | Pointer to **NullableInt32** | Total number of reducers to be spawned. | [optional] 

## Methods

### NewMapReduceInstanceRunInfo

`func NewMapReduceInstanceRunInfo() *MapReduceInstanceRunInfo`

NewMapReduceInstanceRunInfo instantiates a new MapReduceInstanceRunInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMapReduceInstanceRunInfoWithDefaults

`func NewMapReduceInstanceRunInfoWithDefaults() *MapReduceInstanceRunInfo`

NewMapReduceInstanceRunInfoWithDefaults instantiates a new MapReduceInstanceRunInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTime

`func (o *MapReduceInstanceRunInfo) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *MapReduceInstanceRunInfo) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *MapReduceInstanceRunInfo) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *MapReduceInstanceRunInfo) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *MapReduceInstanceRunInfo) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *MapReduceInstanceRunInfo) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetErrorMessage

`func (o *MapReduceInstanceRunInfo) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *MapReduceInstanceRunInfo) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *MapReduceInstanceRunInfo) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *MapReduceInstanceRunInfo) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *MapReduceInstanceRunInfo) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *MapReduceInstanceRunInfo) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetExecutionStartTimeUsecs

`func (o *MapReduceInstanceRunInfo) GetExecutionStartTimeUsecs() int64`

GetExecutionStartTimeUsecs returns the ExecutionStartTimeUsecs field if non-nil, zero value otherwise.

### GetExecutionStartTimeUsecsOk

`func (o *MapReduceInstanceRunInfo) GetExecutionStartTimeUsecsOk() (*int64, bool)`

GetExecutionStartTimeUsecsOk returns a tuple with the ExecutionStartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionStartTimeUsecs

`func (o *MapReduceInstanceRunInfo) SetExecutionStartTimeUsecs(v int64)`

SetExecutionStartTimeUsecs sets ExecutionStartTimeUsecs field to given value.

### HasExecutionStartTimeUsecs

`func (o *MapReduceInstanceRunInfo) HasExecutionStartTimeUsecs() bool`

HasExecutionStartTimeUsecs returns a boolean if a field has been set.

### SetExecutionStartTimeUsecsNil

`func (o *MapReduceInstanceRunInfo) SetExecutionStartTimeUsecsNil(b bool)`

 SetExecutionStartTimeUsecsNil sets the value for ExecutionStartTimeUsecs to be an explicit nil

### UnsetExecutionStartTimeUsecs
`func (o *MapReduceInstanceRunInfo) UnsetExecutionStartTimeUsecs()`

UnsetExecutionStartTimeUsecs ensures that no value is present for ExecutionStartTimeUsecs, not even an explicit nil
### GetFilesProcessed

`func (o *MapReduceInstanceRunInfo) GetFilesProcessed() int32`

GetFilesProcessed returns the FilesProcessed field if non-nil, zero value otherwise.

### GetFilesProcessedOk

`func (o *MapReduceInstanceRunInfo) GetFilesProcessedOk() (*int32, bool)`

GetFilesProcessedOk returns a tuple with the FilesProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesProcessed

`func (o *MapReduceInstanceRunInfo) SetFilesProcessed(v int32)`

SetFilesProcessed sets FilesProcessed field to given value.

### HasFilesProcessed

`func (o *MapReduceInstanceRunInfo) HasFilesProcessed() bool`

HasFilesProcessed returns a boolean if a field has been set.

### SetFilesProcessedNil

`func (o *MapReduceInstanceRunInfo) SetFilesProcessedNil(b bool)`

 SetFilesProcessedNil sets the value for FilesProcessed to be an explicit nil

### UnsetFilesProcessed
`func (o *MapReduceInstanceRunInfo) UnsetFilesProcessed()`

UnsetFilesProcessed ensures that no value is present for FilesProcessed, not even an explicit nil
### GetMapDoneTimeUsecs

`func (o *MapReduceInstanceRunInfo) GetMapDoneTimeUsecs() int64`

GetMapDoneTimeUsecs returns the MapDoneTimeUsecs field if non-nil, zero value otherwise.

### GetMapDoneTimeUsecsOk

`func (o *MapReduceInstanceRunInfo) GetMapDoneTimeUsecsOk() (*int64, bool)`

GetMapDoneTimeUsecsOk returns a tuple with the MapDoneTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapDoneTimeUsecs

`func (o *MapReduceInstanceRunInfo) SetMapDoneTimeUsecs(v int64)`

SetMapDoneTimeUsecs sets MapDoneTimeUsecs field to given value.

### HasMapDoneTimeUsecs

`func (o *MapReduceInstanceRunInfo) HasMapDoneTimeUsecs() bool`

HasMapDoneTimeUsecs returns a boolean if a field has been set.

### SetMapDoneTimeUsecsNil

`func (o *MapReduceInstanceRunInfo) SetMapDoneTimeUsecsNil(b bool)`

 SetMapDoneTimeUsecsNil sets the value for MapDoneTimeUsecs to be an explicit nil

### UnsetMapDoneTimeUsecs
`func (o *MapReduceInstanceRunInfo) UnsetMapDoneTimeUsecs()`

UnsetMapDoneTimeUsecs ensures that no value is present for MapDoneTimeUsecs, not even an explicit nil
### GetMapInputBytes

`func (o *MapReduceInstanceRunInfo) GetMapInputBytes() int64`

GetMapInputBytes returns the MapInputBytes field if non-nil, zero value otherwise.

### GetMapInputBytesOk

`func (o *MapReduceInstanceRunInfo) GetMapInputBytesOk() (*int64, bool)`

GetMapInputBytesOk returns a tuple with the MapInputBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapInputBytes

`func (o *MapReduceInstanceRunInfo) SetMapInputBytes(v int64)`

SetMapInputBytes sets MapInputBytes field to given value.

### HasMapInputBytes

`func (o *MapReduceInstanceRunInfo) HasMapInputBytes() bool`

HasMapInputBytes returns a boolean if a field has been set.

### SetMapInputBytesNil

`func (o *MapReduceInstanceRunInfo) SetMapInputBytesNil(b bool)`

 SetMapInputBytesNil sets the value for MapInputBytes to be an explicit nil

### UnsetMapInputBytes
`func (o *MapReduceInstanceRunInfo) UnsetMapInputBytes()`

UnsetMapInputBytes ensures that no value is present for MapInputBytes, not even an explicit nil
### GetMappersSpawned

`func (o *MapReduceInstanceRunInfo) GetMappersSpawned() int32`

GetMappersSpawned returns the MappersSpawned field if non-nil, zero value otherwise.

### GetMappersSpawnedOk

`func (o *MapReduceInstanceRunInfo) GetMappersSpawnedOk() (*int32, bool)`

GetMappersSpawnedOk returns a tuple with the MappersSpawned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappersSpawned

`func (o *MapReduceInstanceRunInfo) SetMappersSpawned(v int32)`

SetMappersSpawned sets MappersSpawned field to given value.

### HasMappersSpawned

`func (o *MapReduceInstanceRunInfo) HasMappersSpawned() bool`

HasMappersSpawned returns a boolean if a field has been set.

### SetMappersSpawnedNil

`func (o *MapReduceInstanceRunInfo) SetMappersSpawnedNil(b bool)`

 SetMappersSpawnedNil sets the value for MappersSpawned to be an explicit nil

### UnsetMappersSpawned
`func (o *MapReduceInstanceRunInfo) UnsetMappersSpawned()`

UnsetMappersSpawned ensures that no value is present for MappersSpawned, not even an explicit nil
### GetNumMapOutputs

`func (o *MapReduceInstanceRunInfo) GetNumMapOutputs() int64`

GetNumMapOutputs returns the NumMapOutputs field if non-nil, zero value otherwise.

### GetNumMapOutputsOk

`func (o *MapReduceInstanceRunInfo) GetNumMapOutputsOk() (*int64, bool)`

GetNumMapOutputsOk returns a tuple with the NumMapOutputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumMapOutputs

`func (o *MapReduceInstanceRunInfo) SetNumMapOutputs(v int64)`

SetNumMapOutputs sets NumMapOutputs field to given value.

### HasNumMapOutputs

`func (o *MapReduceInstanceRunInfo) HasNumMapOutputs() bool`

HasNumMapOutputs returns a boolean if a field has been set.

### SetNumMapOutputsNil

`func (o *MapReduceInstanceRunInfo) SetNumMapOutputsNil(b bool)`

 SetNumMapOutputsNil sets the value for NumMapOutputs to be an explicit nil

### UnsetNumMapOutputs
`func (o *MapReduceInstanceRunInfo) UnsetNumMapOutputs()`

UnsetNumMapOutputs ensures that no value is present for NumMapOutputs, not even an explicit nil
### GetNumReduceOutputs

`func (o *MapReduceInstanceRunInfo) GetNumReduceOutputs() int64`

GetNumReduceOutputs returns the NumReduceOutputs field if non-nil, zero value otherwise.

### GetNumReduceOutputsOk

`func (o *MapReduceInstanceRunInfo) GetNumReduceOutputsOk() (*int64, bool)`

GetNumReduceOutputsOk returns a tuple with the NumReduceOutputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumReduceOutputs

`func (o *MapReduceInstanceRunInfo) SetNumReduceOutputs(v int64)`

SetNumReduceOutputs sets NumReduceOutputs field to given value.

### HasNumReduceOutputs

`func (o *MapReduceInstanceRunInfo) HasNumReduceOutputs() bool`

HasNumReduceOutputs returns a boolean if a field has been set.

### SetNumReduceOutputsNil

`func (o *MapReduceInstanceRunInfo) SetNumReduceOutputsNil(b bool)`

 SetNumReduceOutputsNil sets the value for NumReduceOutputs to be an explicit nil

### UnsetNumReduceOutputs
`func (o *MapReduceInstanceRunInfo) UnsetNumReduceOutputs()`

UnsetNumReduceOutputs ensures that no value is present for NumReduceOutputs, not even an explicit nil
### GetPercentageCompletion

`func (o *MapReduceInstanceRunInfo) GetPercentageCompletion() float32`

GetPercentageCompletion returns the PercentageCompletion field if non-nil, zero value otherwise.

### GetPercentageCompletionOk

`func (o *MapReduceInstanceRunInfo) GetPercentageCompletionOk() (*float32, bool)`

GetPercentageCompletionOk returns a tuple with the PercentageCompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageCompletion

`func (o *MapReduceInstanceRunInfo) SetPercentageCompletion(v float32)`

SetPercentageCompletion sets PercentageCompletion field to given value.

### HasPercentageCompletion

`func (o *MapReduceInstanceRunInfo) HasPercentageCompletion() bool`

HasPercentageCompletion returns a boolean if a field has been set.

### SetPercentageCompletionNil

`func (o *MapReduceInstanceRunInfo) SetPercentageCompletionNil(b bool)`

 SetPercentageCompletionNil sets the value for PercentageCompletion to be an explicit nil

### UnsetPercentageCompletion
`func (o *MapReduceInstanceRunInfo) UnsetPercentageCompletion()`

UnsetPercentageCompletion ensures that no value is present for PercentageCompletion, not even an explicit nil
### GetPercentageMapperCompletion

`func (o *MapReduceInstanceRunInfo) GetPercentageMapperCompletion() float32`

GetPercentageMapperCompletion returns the PercentageMapperCompletion field if non-nil, zero value otherwise.

### GetPercentageMapperCompletionOk

`func (o *MapReduceInstanceRunInfo) GetPercentageMapperCompletionOk() (*float32, bool)`

GetPercentageMapperCompletionOk returns a tuple with the PercentageMapperCompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageMapperCompletion

`func (o *MapReduceInstanceRunInfo) SetPercentageMapperCompletion(v float32)`

SetPercentageMapperCompletion sets PercentageMapperCompletion field to given value.

### HasPercentageMapperCompletion

`func (o *MapReduceInstanceRunInfo) HasPercentageMapperCompletion() bool`

HasPercentageMapperCompletion returns a boolean if a field has been set.

### SetPercentageMapperCompletionNil

`func (o *MapReduceInstanceRunInfo) SetPercentageMapperCompletionNil(b bool)`

 SetPercentageMapperCompletionNil sets the value for PercentageMapperCompletion to be an explicit nil

### UnsetPercentageMapperCompletion
`func (o *MapReduceInstanceRunInfo) UnsetPercentageMapperCompletion()`

UnsetPercentageMapperCompletion ensures that no value is present for PercentageMapperCompletion, not even an explicit nil
### GetPercentageReducerCompletion

`func (o *MapReduceInstanceRunInfo) GetPercentageReducerCompletion() float32`

GetPercentageReducerCompletion returns the PercentageReducerCompletion field if non-nil, zero value otherwise.

### GetPercentageReducerCompletionOk

`func (o *MapReduceInstanceRunInfo) GetPercentageReducerCompletionOk() (*float32, bool)`

GetPercentageReducerCompletionOk returns a tuple with the PercentageReducerCompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageReducerCompletion

`func (o *MapReduceInstanceRunInfo) SetPercentageReducerCompletion(v float32)`

SetPercentageReducerCompletion sets PercentageReducerCompletion field to given value.

### HasPercentageReducerCompletion

`func (o *MapReduceInstanceRunInfo) HasPercentageReducerCompletion() bool`

HasPercentageReducerCompletion returns a boolean if a field has been set.

### SetPercentageReducerCompletionNil

`func (o *MapReduceInstanceRunInfo) SetPercentageReducerCompletionNil(b bool)`

 SetPercentageReducerCompletionNil sets the value for PercentageReducerCompletion to be an explicit nil

### UnsetPercentageReducerCompletion
`func (o *MapReduceInstanceRunInfo) UnsetPercentageReducerCompletion()`

UnsetPercentageReducerCompletion ensures that no value is present for PercentageReducerCompletion, not even an explicit nil
### GetReducersSpawned

`func (o *MapReduceInstanceRunInfo) GetReducersSpawned() int32`

GetReducersSpawned returns the ReducersSpawned field if non-nil, zero value otherwise.

### GetReducersSpawnedOk

`func (o *MapReduceInstanceRunInfo) GetReducersSpawnedOk() (*int32, bool)`

GetReducersSpawnedOk returns a tuple with the ReducersSpawned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReducersSpawned

`func (o *MapReduceInstanceRunInfo) SetReducersSpawned(v int32)`

SetReducersSpawned sets ReducersSpawned field to given value.

### HasReducersSpawned

`func (o *MapReduceInstanceRunInfo) HasReducersSpawned() bool`

HasReducersSpawned returns a boolean if a field has been set.

### SetReducersSpawnedNil

`func (o *MapReduceInstanceRunInfo) SetReducersSpawnedNil(b bool)`

 SetReducersSpawnedNil sets the value for ReducersSpawned to be an explicit nil

### UnsetReducersSpawned
`func (o *MapReduceInstanceRunInfo) UnsetReducersSpawned()`

UnsetReducersSpawned ensures that no value is present for ReducersSpawned, not even an explicit nil
### GetRemainingTimeMins

`func (o *MapReduceInstanceRunInfo) GetRemainingTimeMins() int32`

GetRemainingTimeMins returns the RemainingTimeMins field if non-nil, zero value otherwise.

### GetRemainingTimeMinsOk

`func (o *MapReduceInstanceRunInfo) GetRemainingTimeMinsOk() (*int32, bool)`

GetRemainingTimeMinsOk returns a tuple with the RemainingTimeMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingTimeMins

`func (o *MapReduceInstanceRunInfo) SetRemainingTimeMins(v int32)`

SetRemainingTimeMins sets RemainingTimeMins field to given value.

### HasRemainingTimeMins

`func (o *MapReduceInstanceRunInfo) HasRemainingTimeMins() bool`

HasRemainingTimeMins returns a boolean if a field has been set.

### SetRemainingTimeMinsNil

`func (o *MapReduceInstanceRunInfo) SetRemainingTimeMinsNil(b bool)`

 SetRemainingTimeMinsNil sets the value for RemainingTimeMins to be an explicit nil

### UnsetRemainingTimeMins
`func (o *MapReduceInstanceRunInfo) UnsetRemainingTimeMins()`

UnsetRemainingTimeMins ensures that no value is present for RemainingTimeMins, not even an explicit nil
### GetStartTime

`func (o *MapReduceInstanceRunInfo) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *MapReduceInstanceRunInfo) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *MapReduceInstanceRunInfo) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *MapReduceInstanceRunInfo) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *MapReduceInstanceRunInfo) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *MapReduceInstanceRunInfo) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetStatus

`func (o *MapReduceInstanceRunInfo) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MapReduceInstanceRunInfo) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MapReduceInstanceRunInfo) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MapReduceInstanceRunInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *MapReduceInstanceRunInfo) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *MapReduceInstanceRunInfo) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTotalNumMappers

`func (o *MapReduceInstanceRunInfo) GetTotalNumMappers() int32`

GetTotalNumMappers returns the TotalNumMappers field if non-nil, zero value otherwise.

### GetTotalNumMappersOk

`func (o *MapReduceInstanceRunInfo) GetTotalNumMappersOk() (*int32, bool)`

GetTotalNumMappersOk returns a tuple with the TotalNumMappers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNumMappers

`func (o *MapReduceInstanceRunInfo) SetTotalNumMappers(v int32)`

SetTotalNumMappers sets TotalNumMappers field to given value.

### HasTotalNumMappers

`func (o *MapReduceInstanceRunInfo) HasTotalNumMappers() bool`

HasTotalNumMappers returns a boolean if a field has been set.

### SetTotalNumMappersNil

`func (o *MapReduceInstanceRunInfo) SetTotalNumMappersNil(b bool)`

 SetTotalNumMappersNil sets the value for TotalNumMappers to be an explicit nil

### UnsetTotalNumMappers
`func (o *MapReduceInstanceRunInfo) UnsetTotalNumMappers()`

UnsetTotalNumMappers ensures that no value is present for TotalNumMappers, not even an explicit nil
### GetTotalNumReducers

`func (o *MapReduceInstanceRunInfo) GetTotalNumReducers() int32`

GetTotalNumReducers returns the TotalNumReducers field if non-nil, zero value otherwise.

### GetTotalNumReducersOk

`func (o *MapReduceInstanceRunInfo) GetTotalNumReducersOk() (*int32, bool)`

GetTotalNumReducersOk returns a tuple with the TotalNumReducers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNumReducers

`func (o *MapReduceInstanceRunInfo) SetTotalNumReducers(v int32)`

SetTotalNumReducers sets TotalNumReducers field to given value.

### HasTotalNumReducers

`func (o *MapReduceInstanceRunInfo) HasTotalNumReducers() bool`

HasTotalNumReducers returns a boolean if a field has been set.

### SetTotalNumReducersNil

`func (o *MapReduceInstanceRunInfo) SetTotalNumReducersNil(b bool)`

 SetTotalNumReducersNil sets the value for TotalNumReducers to be an explicit nil

### UnsetTotalNumReducers
`func (o *MapReduceInstanceRunInfo) UnsetTotalNumReducers()`

UnsetTotalNumReducers ensures that no value is present for TotalNumReducers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


