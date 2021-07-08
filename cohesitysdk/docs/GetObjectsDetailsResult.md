# GetObjectsDetailsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTimeMsecs** | Pointer to **NullableInt64** | Specifies the end time of the run. | [optional] 
**EntityEnv** | Pointer to **NullableInt64** | Specifies the entity environment of the object. | [optional] 
**EntityId** | Pointer to **NullableInt64** | Specifies the entity id of the object. | [optional] 
**EntityName** | Pointer to **NullableString** | Specifies the name of the entity. | [optional] 
**JobId** | Pointer to **NullableString** | Specifies the job id. | [optional] 
**JobRunId** | Pointer to **NullableString** | Specifies the job run id. | [optional] 
**JobType** | Pointer to **NullableString** | Specifies the job type, protection, replication, archival, apollo, indexing etc. | [optional] 
**StartTimeMsecs** | Pointer to **NullableInt64** | Specifies the start time of the run. | [optional] 
**Status** | Pointer to **NullableInt64** | Specifies status of the object run. | [optional] 

## Methods

### NewGetObjectsDetailsResult

`func NewGetObjectsDetailsResult() *GetObjectsDetailsResult`

NewGetObjectsDetailsResult instantiates a new GetObjectsDetailsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetObjectsDetailsResultWithDefaults

`func NewGetObjectsDetailsResultWithDefaults() *GetObjectsDetailsResult`

NewGetObjectsDetailsResultWithDefaults instantiates a new GetObjectsDetailsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTimeMsecs

`func (o *GetObjectsDetailsResult) GetEndTimeMsecs() int64`

GetEndTimeMsecs returns the EndTimeMsecs field if non-nil, zero value otherwise.

### GetEndTimeMsecsOk

`func (o *GetObjectsDetailsResult) GetEndTimeMsecsOk() (*int64, bool)`

GetEndTimeMsecsOk returns a tuple with the EndTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeMsecs

`func (o *GetObjectsDetailsResult) SetEndTimeMsecs(v int64)`

SetEndTimeMsecs sets EndTimeMsecs field to given value.

### HasEndTimeMsecs

`func (o *GetObjectsDetailsResult) HasEndTimeMsecs() bool`

HasEndTimeMsecs returns a boolean if a field has been set.

### SetEndTimeMsecsNil

`func (o *GetObjectsDetailsResult) SetEndTimeMsecsNil(b bool)`

 SetEndTimeMsecsNil sets the value for EndTimeMsecs to be an explicit nil

### UnsetEndTimeMsecs
`func (o *GetObjectsDetailsResult) UnsetEndTimeMsecs()`

UnsetEndTimeMsecs ensures that no value is present for EndTimeMsecs, not even an explicit nil
### GetEntityEnv

`func (o *GetObjectsDetailsResult) GetEntityEnv() int64`

GetEntityEnv returns the EntityEnv field if non-nil, zero value otherwise.

### GetEntityEnvOk

`func (o *GetObjectsDetailsResult) GetEntityEnvOk() (*int64, bool)`

GetEntityEnvOk returns a tuple with the EntityEnv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityEnv

`func (o *GetObjectsDetailsResult) SetEntityEnv(v int64)`

SetEntityEnv sets EntityEnv field to given value.

### HasEntityEnv

`func (o *GetObjectsDetailsResult) HasEntityEnv() bool`

HasEntityEnv returns a boolean if a field has been set.

### SetEntityEnvNil

`func (o *GetObjectsDetailsResult) SetEntityEnvNil(b bool)`

 SetEntityEnvNil sets the value for EntityEnv to be an explicit nil

### UnsetEntityEnv
`func (o *GetObjectsDetailsResult) UnsetEntityEnv()`

UnsetEntityEnv ensures that no value is present for EntityEnv, not even an explicit nil
### GetEntityId

`func (o *GetObjectsDetailsResult) GetEntityId() int64`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *GetObjectsDetailsResult) GetEntityIdOk() (*int64, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *GetObjectsDetailsResult) SetEntityId(v int64)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *GetObjectsDetailsResult) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### SetEntityIdNil

`func (o *GetObjectsDetailsResult) SetEntityIdNil(b bool)`

 SetEntityIdNil sets the value for EntityId to be an explicit nil

### UnsetEntityId
`func (o *GetObjectsDetailsResult) UnsetEntityId()`

UnsetEntityId ensures that no value is present for EntityId, not even an explicit nil
### GetEntityName

`func (o *GetObjectsDetailsResult) GetEntityName() string`

GetEntityName returns the EntityName field if non-nil, zero value otherwise.

### GetEntityNameOk

`func (o *GetObjectsDetailsResult) GetEntityNameOk() (*string, bool)`

GetEntityNameOk returns a tuple with the EntityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityName

`func (o *GetObjectsDetailsResult) SetEntityName(v string)`

SetEntityName sets EntityName field to given value.

### HasEntityName

`func (o *GetObjectsDetailsResult) HasEntityName() bool`

HasEntityName returns a boolean if a field has been set.

### SetEntityNameNil

`func (o *GetObjectsDetailsResult) SetEntityNameNil(b bool)`

 SetEntityNameNil sets the value for EntityName to be an explicit nil

### UnsetEntityName
`func (o *GetObjectsDetailsResult) UnsetEntityName()`

UnsetEntityName ensures that no value is present for EntityName, not even an explicit nil
### GetJobId

`func (o *GetObjectsDetailsResult) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *GetObjectsDetailsResult) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *GetObjectsDetailsResult) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *GetObjectsDetailsResult) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### SetJobIdNil

`func (o *GetObjectsDetailsResult) SetJobIdNil(b bool)`

 SetJobIdNil sets the value for JobId to be an explicit nil

### UnsetJobId
`func (o *GetObjectsDetailsResult) UnsetJobId()`

UnsetJobId ensures that no value is present for JobId, not even an explicit nil
### GetJobRunId

`func (o *GetObjectsDetailsResult) GetJobRunId() string`

GetJobRunId returns the JobRunId field if non-nil, zero value otherwise.

### GetJobRunIdOk

`func (o *GetObjectsDetailsResult) GetJobRunIdOk() (*string, bool)`

GetJobRunIdOk returns a tuple with the JobRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobRunId

`func (o *GetObjectsDetailsResult) SetJobRunId(v string)`

SetJobRunId sets JobRunId field to given value.

### HasJobRunId

`func (o *GetObjectsDetailsResult) HasJobRunId() bool`

HasJobRunId returns a boolean if a field has been set.

### SetJobRunIdNil

`func (o *GetObjectsDetailsResult) SetJobRunIdNil(b bool)`

 SetJobRunIdNil sets the value for JobRunId to be an explicit nil

### UnsetJobRunId
`func (o *GetObjectsDetailsResult) UnsetJobRunId()`

UnsetJobRunId ensures that no value is present for JobRunId, not even an explicit nil
### GetJobType

`func (o *GetObjectsDetailsResult) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *GetObjectsDetailsResult) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *GetObjectsDetailsResult) SetJobType(v string)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *GetObjectsDetailsResult) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### SetJobTypeNil

`func (o *GetObjectsDetailsResult) SetJobTypeNil(b bool)`

 SetJobTypeNil sets the value for JobType to be an explicit nil

### UnsetJobType
`func (o *GetObjectsDetailsResult) UnsetJobType()`

UnsetJobType ensures that no value is present for JobType, not even an explicit nil
### GetStartTimeMsecs

`func (o *GetObjectsDetailsResult) GetStartTimeMsecs() int64`

GetStartTimeMsecs returns the StartTimeMsecs field if non-nil, zero value otherwise.

### GetStartTimeMsecsOk

`func (o *GetObjectsDetailsResult) GetStartTimeMsecsOk() (*int64, bool)`

GetStartTimeMsecsOk returns a tuple with the StartTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeMsecs

`func (o *GetObjectsDetailsResult) SetStartTimeMsecs(v int64)`

SetStartTimeMsecs sets StartTimeMsecs field to given value.

### HasStartTimeMsecs

`func (o *GetObjectsDetailsResult) HasStartTimeMsecs() bool`

HasStartTimeMsecs returns a boolean if a field has been set.

### SetStartTimeMsecsNil

`func (o *GetObjectsDetailsResult) SetStartTimeMsecsNil(b bool)`

 SetStartTimeMsecsNil sets the value for StartTimeMsecs to be an explicit nil

### UnsetStartTimeMsecs
`func (o *GetObjectsDetailsResult) UnsetStartTimeMsecs()`

UnsetStartTimeMsecs ensures that no value is present for StartTimeMsecs, not even an explicit nil
### GetStatus

`func (o *GetObjectsDetailsResult) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetObjectsDetailsResult) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetObjectsDetailsResult) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetObjectsDetailsResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *GetObjectsDetailsResult) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *GetObjectsDetailsResult) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


