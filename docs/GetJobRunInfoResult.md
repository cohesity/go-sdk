# GetJobRunInfoResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BytesTransferred** | Pointer to **NullableInt64** | Specifies bytes transferred in the run. | [optional] 
**EndTimeUsecs** | Pointer to **NullableInt64** | Specifies the end time of the run. | [optional] 
**FailureEntities** | Pointer to **NullableInt64** | Specifies the number of failed objects in the run. | [optional] 
**JobId** | Pointer to **NullableString** | Specifies the job id. | [optional] 
**JobRunId** | Pointer to **NullableString** | Specifies the job run id. | [optional] 
**JobType** | Pointer to **NullableString** | Specifies the job type, protection, replication, archival, apollo, indexing etc. | [optional] 
**SlaViolated** | Pointer to **NullableBool** | Specifies if the sla was violated the run. | [optional] 
**StartTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time of the run. | [optional] 
**Status** | Pointer to **NullableInt64** | Specifies status of the run | [optional] 
**SuccessEntities** | Pointer to **NullableInt64** | Specifies the number successful objects in the run. | [optional] 
**TotalEntities** | Pointer to **NullableInt64** | Specifies the number of objects in the run. | [optional] 

## Methods

### NewGetJobRunInfoResult

`func NewGetJobRunInfoResult() *GetJobRunInfoResult`

NewGetJobRunInfoResult instantiates a new GetJobRunInfoResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetJobRunInfoResultWithDefaults

`func NewGetJobRunInfoResultWithDefaults() *GetJobRunInfoResult`

NewGetJobRunInfoResultWithDefaults instantiates a new GetJobRunInfoResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBytesTransferred

`func (o *GetJobRunInfoResult) GetBytesTransferred() int64`

GetBytesTransferred returns the BytesTransferred field if non-nil, zero value otherwise.

### GetBytesTransferredOk

`func (o *GetJobRunInfoResult) GetBytesTransferredOk() (*int64, bool)`

GetBytesTransferredOk returns a tuple with the BytesTransferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesTransferred

`func (o *GetJobRunInfoResult) SetBytesTransferred(v int64)`

SetBytesTransferred sets BytesTransferred field to given value.

### HasBytesTransferred

`func (o *GetJobRunInfoResult) HasBytesTransferred() bool`

HasBytesTransferred returns a boolean if a field has been set.

### SetBytesTransferredNil

`func (o *GetJobRunInfoResult) SetBytesTransferredNil(b bool)`

 SetBytesTransferredNil sets the value for BytesTransferred to be an explicit nil

### UnsetBytesTransferred
`func (o *GetJobRunInfoResult) UnsetBytesTransferred()`

UnsetBytesTransferred ensures that no value is present for BytesTransferred, not even an explicit nil
### GetEndTimeUsecs

`func (o *GetJobRunInfoResult) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *GetJobRunInfoResult) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *GetJobRunInfoResult) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *GetJobRunInfoResult) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *GetJobRunInfoResult) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *GetJobRunInfoResult) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetFailureEntities

`func (o *GetJobRunInfoResult) GetFailureEntities() int64`

GetFailureEntities returns the FailureEntities field if non-nil, zero value otherwise.

### GetFailureEntitiesOk

`func (o *GetJobRunInfoResult) GetFailureEntitiesOk() (*int64, bool)`

GetFailureEntitiesOk returns a tuple with the FailureEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureEntities

`func (o *GetJobRunInfoResult) SetFailureEntities(v int64)`

SetFailureEntities sets FailureEntities field to given value.

### HasFailureEntities

`func (o *GetJobRunInfoResult) HasFailureEntities() bool`

HasFailureEntities returns a boolean if a field has been set.

### SetFailureEntitiesNil

`func (o *GetJobRunInfoResult) SetFailureEntitiesNil(b bool)`

 SetFailureEntitiesNil sets the value for FailureEntities to be an explicit nil

### UnsetFailureEntities
`func (o *GetJobRunInfoResult) UnsetFailureEntities()`

UnsetFailureEntities ensures that no value is present for FailureEntities, not even an explicit nil
### GetJobId

`func (o *GetJobRunInfoResult) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *GetJobRunInfoResult) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *GetJobRunInfoResult) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *GetJobRunInfoResult) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### SetJobIdNil

`func (o *GetJobRunInfoResult) SetJobIdNil(b bool)`

 SetJobIdNil sets the value for JobId to be an explicit nil

### UnsetJobId
`func (o *GetJobRunInfoResult) UnsetJobId()`

UnsetJobId ensures that no value is present for JobId, not even an explicit nil
### GetJobRunId

`func (o *GetJobRunInfoResult) GetJobRunId() string`

GetJobRunId returns the JobRunId field if non-nil, zero value otherwise.

### GetJobRunIdOk

`func (o *GetJobRunInfoResult) GetJobRunIdOk() (*string, bool)`

GetJobRunIdOk returns a tuple with the JobRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobRunId

`func (o *GetJobRunInfoResult) SetJobRunId(v string)`

SetJobRunId sets JobRunId field to given value.

### HasJobRunId

`func (o *GetJobRunInfoResult) HasJobRunId() bool`

HasJobRunId returns a boolean if a field has been set.

### SetJobRunIdNil

`func (o *GetJobRunInfoResult) SetJobRunIdNil(b bool)`

 SetJobRunIdNil sets the value for JobRunId to be an explicit nil

### UnsetJobRunId
`func (o *GetJobRunInfoResult) UnsetJobRunId()`

UnsetJobRunId ensures that no value is present for JobRunId, not even an explicit nil
### GetJobType

`func (o *GetJobRunInfoResult) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *GetJobRunInfoResult) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *GetJobRunInfoResult) SetJobType(v string)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *GetJobRunInfoResult) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### SetJobTypeNil

`func (o *GetJobRunInfoResult) SetJobTypeNil(b bool)`

 SetJobTypeNil sets the value for JobType to be an explicit nil

### UnsetJobType
`func (o *GetJobRunInfoResult) UnsetJobType()`

UnsetJobType ensures that no value is present for JobType, not even an explicit nil
### GetSlaViolated

`func (o *GetJobRunInfoResult) GetSlaViolated() bool`

GetSlaViolated returns the SlaViolated field if non-nil, zero value otherwise.

### GetSlaViolatedOk

`func (o *GetJobRunInfoResult) GetSlaViolatedOk() (*bool, bool)`

GetSlaViolatedOk returns a tuple with the SlaViolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaViolated

`func (o *GetJobRunInfoResult) SetSlaViolated(v bool)`

SetSlaViolated sets SlaViolated field to given value.

### HasSlaViolated

`func (o *GetJobRunInfoResult) HasSlaViolated() bool`

HasSlaViolated returns a boolean if a field has been set.

### SetSlaViolatedNil

`func (o *GetJobRunInfoResult) SetSlaViolatedNil(b bool)`

 SetSlaViolatedNil sets the value for SlaViolated to be an explicit nil

### UnsetSlaViolated
`func (o *GetJobRunInfoResult) UnsetSlaViolated()`

UnsetSlaViolated ensures that no value is present for SlaViolated, not even an explicit nil
### GetStartTimeUsecs

`func (o *GetJobRunInfoResult) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *GetJobRunInfoResult) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *GetJobRunInfoResult) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *GetJobRunInfoResult) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *GetJobRunInfoResult) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *GetJobRunInfoResult) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetStatus

`func (o *GetJobRunInfoResult) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetJobRunInfoResult) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetJobRunInfoResult) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetJobRunInfoResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *GetJobRunInfoResult) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *GetJobRunInfoResult) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetSuccessEntities

`func (o *GetJobRunInfoResult) GetSuccessEntities() int64`

GetSuccessEntities returns the SuccessEntities field if non-nil, zero value otherwise.

### GetSuccessEntitiesOk

`func (o *GetJobRunInfoResult) GetSuccessEntitiesOk() (*int64, bool)`

GetSuccessEntitiesOk returns a tuple with the SuccessEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessEntities

`func (o *GetJobRunInfoResult) SetSuccessEntities(v int64)`

SetSuccessEntities sets SuccessEntities field to given value.

### HasSuccessEntities

`func (o *GetJobRunInfoResult) HasSuccessEntities() bool`

HasSuccessEntities returns a boolean if a field has been set.

### SetSuccessEntitiesNil

`func (o *GetJobRunInfoResult) SetSuccessEntitiesNil(b bool)`

 SetSuccessEntitiesNil sets the value for SuccessEntities to be an explicit nil

### UnsetSuccessEntities
`func (o *GetJobRunInfoResult) UnsetSuccessEntities()`

UnsetSuccessEntities ensures that no value is present for SuccessEntities, not even an explicit nil
### GetTotalEntities

`func (o *GetJobRunInfoResult) GetTotalEntities() int64`

GetTotalEntities returns the TotalEntities field if non-nil, zero value otherwise.

### GetTotalEntitiesOk

`func (o *GetJobRunInfoResult) GetTotalEntitiesOk() (*int64, bool)`

GetTotalEntitiesOk returns a tuple with the TotalEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEntities

`func (o *GetJobRunInfoResult) SetTotalEntities(v int64)`

SetTotalEntities sets TotalEntities field to given value.

### HasTotalEntities

`func (o *GetJobRunInfoResult) HasTotalEntities() bool`

HasTotalEntities returns a boolean if a field has been set.

### SetTotalEntitiesNil

`func (o *GetJobRunInfoResult) SetTotalEntitiesNil(b bool)`

 SetTotalEntitiesNil sets the value for TotalEntities to be an explicit nil

### UnsetTotalEntities
`func (o *GetJobRunInfoResult) UnsetTotalEntities()`

UnsetTotalEntities ensures that no value is present for TotalEntities, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


