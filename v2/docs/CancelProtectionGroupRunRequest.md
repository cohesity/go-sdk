# CancelProtectionGroupRunRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchivalTaskId** | Pointer to **[]string** | Specifies the task id of the archival run. | [optional] 
**CloudSpinTaskId** | Pointer to **[]string** | Specifies the task id of the cloudSpin run. | [optional] 
**LocalTaskId** | Pointer to **NullableString** | Specifies the task id of the local run. | [optional] 
**ObjectIds** | Pointer to **[]int64** | List of entity ids for which we need to cancel the backup tasks. If this is provided it will not cancel the complete run but will cancel only subset of backup tasks (if backup tasks are cancelled correspoding copy task will also get cancelled). If the backup tasks are completed successfully it will not cancel those backup tasks. | [optional] 
**ReplicationTaskId** | Pointer to **[]string** | Specifies the task id of the replication run. | [optional] 
**RunId** | **NullableString** | Specifies a unique run id of the Protection Group run. | 

## Methods

### NewCancelProtectionGroupRunRequest

`func NewCancelProtectionGroupRunRequest(runId NullableString, ) *CancelProtectionGroupRunRequest`

NewCancelProtectionGroupRunRequest instantiates a new CancelProtectionGroupRunRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelProtectionGroupRunRequestWithDefaults

`func NewCancelProtectionGroupRunRequestWithDefaults() *CancelProtectionGroupRunRequest`

NewCancelProtectionGroupRunRequestWithDefaults instantiates a new CancelProtectionGroupRunRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchivalTaskId

`func (o *CancelProtectionGroupRunRequest) GetArchivalTaskId() []string`

GetArchivalTaskId returns the ArchivalTaskId field if non-nil, zero value otherwise.

### GetArchivalTaskIdOk

`func (o *CancelProtectionGroupRunRequest) GetArchivalTaskIdOk() (*[]string, bool)`

GetArchivalTaskIdOk returns a tuple with the ArchivalTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivalTaskId

`func (o *CancelProtectionGroupRunRequest) SetArchivalTaskId(v []string)`

SetArchivalTaskId sets ArchivalTaskId field to given value.

### HasArchivalTaskId

`func (o *CancelProtectionGroupRunRequest) HasArchivalTaskId() bool`

HasArchivalTaskId returns a boolean if a field has been set.

### SetArchivalTaskIdNil

`func (o *CancelProtectionGroupRunRequest) SetArchivalTaskIdNil(b bool)`

 SetArchivalTaskIdNil sets the value for ArchivalTaskId to be an explicit nil

### UnsetArchivalTaskId
`func (o *CancelProtectionGroupRunRequest) UnsetArchivalTaskId()`

UnsetArchivalTaskId ensures that no value is present for ArchivalTaskId, not even an explicit nil
### GetCloudSpinTaskId

`func (o *CancelProtectionGroupRunRequest) GetCloudSpinTaskId() []string`

GetCloudSpinTaskId returns the CloudSpinTaskId field if non-nil, zero value otherwise.

### GetCloudSpinTaskIdOk

`func (o *CancelProtectionGroupRunRequest) GetCloudSpinTaskIdOk() (*[]string, bool)`

GetCloudSpinTaskIdOk returns a tuple with the CloudSpinTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudSpinTaskId

`func (o *CancelProtectionGroupRunRequest) SetCloudSpinTaskId(v []string)`

SetCloudSpinTaskId sets CloudSpinTaskId field to given value.

### HasCloudSpinTaskId

`func (o *CancelProtectionGroupRunRequest) HasCloudSpinTaskId() bool`

HasCloudSpinTaskId returns a boolean if a field has been set.

### SetCloudSpinTaskIdNil

`func (o *CancelProtectionGroupRunRequest) SetCloudSpinTaskIdNil(b bool)`

 SetCloudSpinTaskIdNil sets the value for CloudSpinTaskId to be an explicit nil

### UnsetCloudSpinTaskId
`func (o *CancelProtectionGroupRunRequest) UnsetCloudSpinTaskId()`

UnsetCloudSpinTaskId ensures that no value is present for CloudSpinTaskId, not even an explicit nil
### GetLocalTaskId

`func (o *CancelProtectionGroupRunRequest) GetLocalTaskId() string`

GetLocalTaskId returns the LocalTaskId field if non-nil, zero value otherwise.

### GetLocalTaskIdOk

`func (o *CancelProtectionGroupRunRequest) GetLocalTaskIdOk() (*string, bool)`

GetLocalTaskIdOk returns a tuple with the LocalTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalTaskId

`func (o *CancelProtectionGroupRunRequest) SetLocalTaskId(v string)`

SetLocalTaskId sets LocalTaskId field to given value.

### HasLocalTaskId

`func (o *CancelProtectionGroupRunRequest) HasLocalTaskId() bool`

HasLocalTaskId returns a boolean if a field has been set.

### SetLocalTaskIdNil

`func (o *CancelProtectionGroupRunRequest) SetLocalTaskIdNil(b bool)`

 SetLocalTaskIdNil sets the value for LocalTaskId to be an explicit nil

### UnsetLocalTaskId
`func (o *CancelProtectionGroupRunRequest) UnsetLocalTaskId()`

UnsetLocalTaskId ensures that no value is present for LocalTaskId, not even an explicit nil
### GetObjectIds

`func (o *CancelProtectionGroupRunRequest) GetObjectIds() []int64`

GetObjectIds returns the ObjectIds field if non-nil, zero value otherwise.

### GetObjectIdsOk

`func (o *CancelProtectionGroupRunRequest) GetObjectIdsOk() (*[]int64, bool)`

GetObjectIdsOk returns a tuple with the ObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectIds

`func (o *CancelProtectionGroupRunRequest) SetObjectIds(v []int64)`

SetObjectIds sets ObjectIds field to given value.

### HasObjectIds

`func (o *CancelProtectionGroupRunRequest) HasObjectIds() bool`

HasObjectIds returns a boolean if a field has been set.

### SetObjectIdsNil

`func (o *CancelProtectionGroupRunRequest) SetObjectIdsNil(b bool)`

 SetObjectIdsNil sets the value for ObjectIds to be an explicit nil

### UnsetObjectIds
`func (o *CancelProtectionGroupRunRequest) UnsetObjectIds()`

UnsetObjectIds ensures that no value is present for ObjectIds, not even an explicit nil
### GetReplicationTaskId

`func (o *CancelProtectionGroupRunRequest) GetReplicationTaskId() []string`

GetReplicationTaskId returns the ReplicationTaskId field if non-nil, zero value otherwise.

### GetReplicationTaskIdOk

`func (o *CancelProtectionGroupRunRequest) GetReplicationTaskIdOk() (*[]string, bool)`

GetReplicationTaskIdOk returns a tuple with the ReplicationTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationTaskId

`func (o *CancelProtectionGroupRunRequest) SetReplicationTaskId(v []string)`

SetReplicationTaskId sets ReplicationTaskId field to given value.

### HasReplicationTaskId

`func (o *CancelProtectionGroupRunRequest) HasReplicationTaskId() bool`

HasReplicationTaskId returns a boolean if a field has been set.

### SetReplicationTaskIdNil

`func (o *CancelProtectionGroupRunRequest) SetReplicationTaskIdNil(b bool)`

 SetReplicationTaskIdNil sets the value for ReplicationTaskId to be an explicit nil

### UnsetReplicationTaskId
`func (o *CancelProtectionGroupRunRequest) UnsetReplicationTaskId()`

UnsetReplicationTaskId ensures that no value is present for ReplicationTaskId, not even an explicit nil
### GetRunId

`func (o *CancelProtectionGroupRunRequest) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *CancelProtectionGroupRunRequest) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *CancelProtectionGroupRunRequest) SetRunId(v string)`

SetRunId sets RunId field to given value.


### SetRunIdNil

`func (o *CancelProtectionGroupRunRequest) SetRunIdNil(b bool)`

 SetRunIdNil sets the value for RunId to be an explicit nil

### UnsetRunId
`func (o *CancelProtectionGroupRunRequest) UnsetRunId()`

UnsetRunId ensures that no value is present for RunId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


