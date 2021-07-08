# GetMapReduceAppRunsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **NullableInt64** | ApplicationId is the Id of the map reduce application. | [optional] 
**AppInstanceId** | Pointer to **NullableInt64** | ApplicationInstanceId is the Id of the map reduce application instance. | [optional] 
**IncludeDetails** | Pointer to **NullableBool** | If this flag is true, then send details of instance, else send only RunInfo. | [optional] 
**LastNumInstances** | Pointer to **NullableInt32** | Give last N instance of an app based on end time. | [optional] 
**MaxRunEndTimeInSecs** | Pointer to **NullableInt64** | MaxRunEndTimestampInSecs specifies the maximum job run end timestamp in seconds. App run instances with end time less than equal to MaxRunEndTimestampInSecs will be selected. Default is LONG_MAX (inf). | [optional] 
**MaxRunStartTimeInSecs** | Pointer to **NullableInt64** | MaxRunStartTimestampInSecs specifies the maximum job run start timestamp in seconds. App run instances with start time less than equal to MaxRunStartTimestampInSecs will be selected. Default is LONG_MAX (inf). | [optional] 
**MinRunEndTimeInSecs** | Pointer to **NullableInt64** | MinRunEndTimestampInSecs specifies the minimum job run end timestamp in seconds. App run instances with end time greater than equal to MinRunEndTimestampInSecs will be selected. Default is 0, i.e. beginning of time. | [optional] 
**MinRunStartTimeInSecs** | Pointer to **NullableInt64** | MinRunStartTimestampInSecs specifies the minimum job run start timestamp in seconds. App run instances with start time greater than equal to MinRunStartTimestampInSecs will be selected. Default is 0, i.e. beginning of time. | [optional] 
**PageSize** | Pointer to **NullableInt32** | Number of results to be displayed on a page. | [optional] 
**RunStatus** | Pointer to **NullableString** | Filter instances based on the map reduce application run status. | [optional] 
**StartOffset** | Pointer to **NullableInt32** | Start offset for pagination from where result needs to be fetched. | [optional] 

## Methods

### NewGetMapReduceAppRunsParams

`func NewGetMapReduceAppRunsParams() *GetMapReduceAppRunsParams`

NewGetMapReduceAppRunsParams instantiates a new GetMapReduceAppRunsParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMapReduceAppRunsParamsWithDefaults

`func NewGetMapReduceAppRunsParamsWithDefaults() *GetMapReduceAppRunsParams`

NewGetMapReduceAppRunsParamsWithDefaults instantiates a new GetMapReduceAppRunsParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *GetMapReduceAppRunsParams) GetAppId() int64`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *GetMapReduceAppRunsParams) GetAppIdOk() (*int64, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *GetMapReduceAppRunsParams) SetAppId(v int64)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *GetMapReduceAppRunsParams) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### SetAppIdNil

`func (o *GetMapReduceAppRunsParams) SetAppIdNil(b bool)`

 SetAppIdNil sets the value for AppId to be an explicit nil

### UnsetAppId
`func (o *GetMapReduceAppRunsParams) UnsetAppId()`

UnsetAppId ensures that no value is present for AppId, not even an explicit nil
### GetAppInstanceId

`func (o *GetMapReduceAppRunsParams) GetAppInstanceId() int64`

GetAppInstanceId returns the AppInstanceId field if non-nil, zero value otherwise.

### GetAppInstanceIdOk

`func (o *GetMapReduceAppRunsParams) GetAppInstanceIdOk() (*int64, bool)`

GetAppInstanceIdOk returns a tuple with the AppInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInstanceId

`func (o *GetMapReduceAppRunsParams) SetAppInstanceId(v int64)`

SetAppInstanceId sets AppInstanceId field to given value.

### HasAppInstanceId

`func (o *GetMapReduceAppRunsParams) HasAppInstanceId() bool`

HasAppInstanceId returns a boolean if a field has been set.

### SetAppInstanceIdNil

`func (o *GetMapReduceAppRunsParams) SetAppInstanceIdNil(b bool)`

 SetAppInstanceIdNil sets the value for AppInstanceId to be an explicit nil

### UnsetAppInstanceId
`func (o *GetMapReduceAppRunsParams) UnsetAppInstanceId()`

UnsetAppInstanceId ensures that no value is present for AppInstanceId, not even an explicit nil
### GetIncludeDetails

`func (o *GetMapReduceAppRunsParams) GetIncludeDetails() bool`

GetIncludeDetails returns the IncludeDetails field if non-nil, zero value otherwise.

### GetIncludeDetailsOk

`func (o *GetMapReduceAppRunsParams) GetIncludeDetailsOk() (*bool, bool)`

GetIncludeDetailsOk returns a tuple with the IncludeDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDetails

`func (o *GetMapReduceAppRunsParams) SetIncludeDetails(v bool)`

SetIncludeDetails sets IncludeDetails field to given value.

### HasIncludeDetails

`func (o *GetMapReduceAppRunsParams) HasIncludeDetails() bool`

HasIncludeDetails returns a boolean if a field has been set.

### SetIncludeDetailsNil

`func (o *GetMapReduceAppRunsParams) SetIncludeDetailsNil(b bool)`

 SetIncludeDetailsNil sets the value for IncludeDetails to be an explicit nil

### UnsetIncludeDetails
`func (o *GetMapReduceAppRunsParams) UnsetIncludeDetails()`

UnsetIncludeDetails ensures that no value is present for IncludeDetails, not even an explicit nil
### GetLastNumInstances

`func (o *GetMapReduceAppRunsParams) GetLastNumInstances() int32`

GetLastNumInstances returns the LastNumInstances field if non-nil, zero value otherwise.

### GetLastNumInstancesOk

`func (o *GetMapReduceAppRunsParams) GetLastNumInstancesOk() (*int32, bool)`

GetLastNumInstancesOk returns a tuple with the LastNumInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastNumInstances

`func (o *GetMapReduceAppRunsParams) SetLastNumInstances(v int32)`

SetLastNumInstances sets LastNumInstances field to given value.

### HasLastNumInstances

`func (o *GetMapReduceAppRunsParams) HasLastNumInstances() bool`

HasLastNumInstances returns a boolean if a field has been set.

### SetLastNumInstancesNil

`func (o *GetMapReduceAppRunsParams) SetLastNumInstancesNil(b bool)`

 SetLastNumInstancesNil sets the value for LastNumInstances to be an explicit nil

### UnsetLastNumInstances
`func (o *GetMapReduceAppRunsParams) UnsetLastNumInstances()`

UnsetLastNumInstances ensures that no value is present for LastNumInstances, not even an explicit nil
### GetMaxRunEndTimeInSecs

`func (o *GetMapReduceAppRunsParams) GetMaxRunEndTimeInSecs() int64`

GetMaxRunEndTimeInSecs returns the MaxRunEndTimeInSecs field if non-nil, zero value otherwise.

### GetMaxRunEndTimeInSecsOk

`func (o *GetMapReduceAppRunsParams) GetMaxRunEndTimeInSecsOk() (*int64, bool)`

GetMaxRunEndTimeInSecsOk returns a tuple with the MaxRunEndTimeInSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRunEndTimeInSecs

`func (o *GetMapReduceAppRunsParams) SetMaxRunEndTimeInSecs(v int64)`

SetMaxRunEndTimeInSecs sets MaxRunEndTimeInSecs field to given value.

### HasMaxRunEndTimeInSecs

`func (o *GetMapReduceAppRunsParams) HasMaxRunEndTimeInSecs() bool`

HasMaxRunEndTimeInSecs returns a boolean if a field has been set.

### SetMaxRunEndTimeInSecsNil

`func (o *GetMapReduceAppRunsParams) SetMaxRunEndTimeInSecsNil(b bool)`

 SetMaxRunEndTimeInSecsNil sets the value for MaxRunEndTimeInSecs to be an explicit nil

### UnsetMaxRunEndTimeInSecs
`func (o *GetMapReduceAppRunsParams) UnsetMaxRunEndTimeInSecs()`

UnsetMaxRunEndTimeInSecs ensures that no value is present for MaxRunEndTimeInSecs, not even an explicit nil
### GetMaxRunStartTimeInSecs

`func (o *GetMapReduceAppRunsParams) GetMaxRunStartTimeInSecs() int64`

GetMaxRunStartTimeInSecs returns the MaxRunStartTimeInSecs field if non-nil, zero value otherwise.

### GetMaxRunStartTimeInSecsOk

`func (o *GetMapReduceAppRunsParams) GetMaxRunStartTimeInSecsOk() (*int64, bool)`

GetMaxRunStartTimeInSecsOk returns a tuple with the MaxRunStartTimeInSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRunStartTimeInSecs

`func (o *GetMapReduceAppRunsParams) SetMaxRunStartTimeInSecs(v int64)`

SetMaxRunStartTimeInSecs sets MaxRunStartTimeInSecs field to given value.

### HasMaxRunStartTimeInSecs

`func (o *GetMapReduceAppRunsParams) HasMaxRunStartTimeInSecs() bool`

HasMaxRunStartTimeInSecs returns a boolean if a field has been set.

### SetMaxRunStartTimeInSecsNil

`func (o *GetMapReduceAppRunsParams) SetMaxRunStartTimeInSecsNil(b bool)`

 SetMaxRunStartTimeInSecsNil sets the value for MaxRunStartTimeInSecs to be an explicit nil

### UnsetMaxRunStartTimeInSecs
`func (o *GetMapReduceAppRunsParams) UnsetMaxRunStartTimeInSecs()`

UnsetMaxRunStartTimeInSecs ensures that no value is present for MaxRunStartTimeInSecs, not even an explicit nil
### GetMinRunEndTimeInSecs

`func (o *GetMapReduceAppRunsParams) GetMinRunEndTimeInSecs() int64`

GetMinRunEndTimeInSecs returns the MinRunEndTimeInSecs field if non-nil, zero value otherwise.

### GetMinRunEndTimeInSecsOk

`func (o *GetMapReduceAppRunsParams) GetMinRunEndTimeInSecsOk() (*int64, bool)`

GetMinRunEndTimeInSecsOk returns a tuple with the MinRunEndTimeInSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRunEndTimeInSecs

`func (o *GetMapReduceAppRunsParams) SetMinRunEndTimeInSecs(v int64)`

SetMinRunEndTimeInSecs sets MinRunEndTimeInSecs field to given value.

### HasMinRunEndTimeInSecs

`func (o *GetMapReduceAppRunsParams) HasMinRunEndTimeInSecs() bool`

HasMinRunEndTimeInSecs returns a boolean if a field has been set.

### SetMinRunEndTimeInSecsNil

`func (o *GetMapReduceAppRunsParams) SetMinRunEndTimeInSecsNil(b bool)`

 SetMinRunEndTimeInSecsNil sets the value for MinRunEndTimeInSecs to be an explicit nil

### UnsetMinRunEndTimeInSecs
`func (o *GetMapReduceAppRunsParams) UnsetMinRunEndTimeInSecs()`

UnsetMinRunEndTimeInSecs ensures that no value is present for MinRunEndTimeInSecs, not even an explicit nil
### GetMinRunStartTimeInSecs

`func (o *GetMapReduceAppRunsParams) GetMinRunStartTimeInSecs() int64`

GetMinRunStartTimeInSecs returns the MinRunStartTimeInSecs field if non-nil, zero value otherwise.

### GetMinRunStartTimeInSecsOk

`func (o *GetMapReduceAppRunsParams) GetMinRunStartTimeInSecsOk() (*int64, bool)`

GetMinRunStartTimeInSecsOk returns a tuple with the MinRunStartTimeInSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRunStartTimeInSecs

`func (o *GetMapReduceAppRunsParams) SetMinRunStartTimeInSecs(v int64)`

SetMinRunStartTimeInSecs sets MinRunStartTimeInSecs field to given value.

### HasMinRunStartTimeInSecs

`func (o *GetMapReduceAppRunsParams) HasMinRunStartTimeInSecs() bool`

HasMinRunStartTimeInSecs returns a boolean if a field has been set.

### SetMinRunStartTimeInSecsNil

`func (o *GetMapReduceAppRunsParams) SetMinRunStartTimeInSecsNil(b bool)`

 SetMinRunStartTimeInSecsNil sets the value for MinRunStartTimeInSecs to be an explicit nil

### UnsetMinRunStartTimeInSecs
`func (o *GetMapReduceAppRunsParams) UnsetMinRunStartTimeInSecs()`

UnsetMinRunStartTimeInSecs ensures that no value is present for MinRunStartTimeInSecs, not even an explicit nil
### GetPageSize

`func (o *GetMapReduceAppRunsParams) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *GetMapReduceAppRunsParams) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *GetMapReduceAppRunsParams) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *GetMapReduceAppRunsParams) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### SetPageSizeNil

`func (o *GetMapReduceAppRunsParams) SetPageSizeNil(b bool)`

 SetPageSizeNil sets the value for PageSize to be an explicit nil

### UnsetPageSize
`func (o *GetMapReduceAppRunsParams) UnsetPageSize()`

UnsetPageSize ensures that no value is present for PageSize, not even an explicit nil
### GetRunStatus

`func (o *GetMapReduceAppRunsParams) GetRunStatus() string`

GetRunStatus returns the RunStatus field if non-nil, zero value otherwise.

### GetRunStatusOk

`func (o *GetMapReduceAppRunsParams) GetRunStatusOk() (*string, bool)`

GetRunStatusOk returns a tuple with the RunStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunStatus

`func (o *GetMapReduceAppRunsParams) SetRunStatus(v string)`

SetRunStatus sets RunStatus field to given value.

### HasRunStatus

`func (o *GetMapReduceAppRunsParams) HasRunStatus() bool`

HasRunStatus returns a boolean if a field has been set.

### SetRunStatusNil

`func (o *GetMapReduceAppRunsParams) SetRunStatusNil(b bool)`

 SetRunStatusNil sets the value for RunStatus to be an explicit nil

### UnsetRunStatus
`func (o *GetMapReduceAppRunsParams) UnsetRunStatus()`

UnsetRunStatus ensures that no value is present for RunStatus, not even an explicit nil
### GetStartOffset

`func (o *GetMapReduceAppRunsParams) GetStartOffset() int32`

GetStartOffset returns the StartOffset field if non-nil, zero value otherwise.

### GetStartOffsetOk

`func (o *GetMapReduceAppRunsParams) GetStartOffsetOk() (*int32, bool)`

GetStartOffsetOk returns a tuple with the StartOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartOffset

`func (o *GetMapReduceAppRunsParams) SetStartOffset(v int32)`

SetStartOffset sets StartOffset field to given value.

### HasStartOffset

`func (o *GetMapReduceAppRunsParams) HasStartOffset() bool`

HasStartOffset returns a boolean if a field has been set.

### SetStartOffsetNil

`func (o *GetMapReduceAppRunsParams) SetStartOffsetNil(b bool)`

 SetStartOffsetNil sets the value for StartOffset to be an explicit nil

### UnsetStartOffset
`func (o *GetMapReduceAppRunsParams) UnsetStartOffset()`

UnsetStartOffset ensures that no value is present for StartOffset, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


