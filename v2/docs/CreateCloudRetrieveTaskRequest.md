# CreateCloudRetrieveTaskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **NullableInt64** | Specifies the cluster ID. | [optional] 
**EndTimeUsecs** | Pointer to **NullableInt64** | Specifies the end time in microseconds and filter the tasks by the end time. | [optional] 
**JobIds** | Pointer to **[]string** | Job ids to retrieve. | [optional] 
**RetrieveAllJobs** | Pointer to **NullableBool** | Specifies whether to retrieve all tasks. | [optional] 
**StartTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time in microseconds and filter the task by the start time. | [optional] 
**VaultIds** | Pointer to **[]int64** | Specifies the array of vault IDs. | [optional] 

## Methods

### NewCreateCloudRetrieveTaskRequest

`func NewCreateCloudRetrieveTaskRequest() *CreateCloudRetrieveTaskRequest`

NewCreateCloudRetrieveTaskRequest instantiates a new CreateCloudRetrieveTaskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCloudRetrieveTaskRequestWithDefaults

`func NewCreateCloudRetrieveTaskRequestWithDefaults() *CreateCloudRetrieveTaskRequest`

NewCreateCloudRetrieveTaskRequestWithDefaults instantiates a new CreateCloudRetrieveTaskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *CreateCloudRetrieveTaskRequest) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *CreateCloudRetrieveTaskRequest) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *CreateCloudRetrieveTaskRequest) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *CreateCloudRetrieveTaskRequest) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### SetClusterIdNil

`func (o *CreateCloudRetrieveTaskRequest) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *CreateCloudRetrieveTaskRequest) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetEndTimeUsecs

`func (o *CreateCloudRetrieveTaskRequest) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *CreateCloudRetrieveTaskRequest) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *CreateCloudRetrieveTaskRequest) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *CreateCloudRetrieveTaskRequest) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *CreateCloudRetrieveTaskRequest) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *CreateCloudRetrieveTaskRequest) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetJobIds

`func (o *CreateCloudRetrieveTaskRequest) GetJobIds() []string`

GetJobIds returns the JobIds field if non-nil, zero value otherwise.

### GetJobIdsOk

`func (o *CreateCloudRetrieveTaskRequest) GetJobIdsOk() (*[]string, bool)`

GetJobIdsOk returns a tuple with the JobIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobIds

`func (o *CreateCloudRetrieveTaskRequest) SetJobIds(v []string)`

SetJobIds sets JobIds field to given value.

### HasJobIds

`func (o *CreateCloudRetrieveTaskRequest) HasJobIds() bool`

HasJobIds returns a boolean if a field has been set.

### SetJobIdsNil

`func (o *CreateCloudRetrieveTaskRequest) SetJobIdsNil(b bool)`

 SetJobIdsNil sets the value for JobIds to be an explicit nil

### UnsetJobIds
`func (o *CreateCloudRetrieveTaskRequest) UnsetJobIds()`

UnsetJobIds ensures that no value is present for JobIds, not even an explicit nil
### GetRetrieveAllJobs

`func (o *CreateCloudRetrieveTaskRequest) GetRetrieveAllJobs() bool`

GetRetrieveAllJobs returns the RetrieveAllJobs field if non-nil, zero value otherwise.

### GetRetrieveAllJobsOk

`func (o *CreateCloudRetrieveTaskRequest) GetRetrieveAllJobsOk() (*bool, bool)`

GetRetrieveAllJobsOk returns a tuple with the RetrieveAllJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetrieveAllJobs

`func (o *CreateCloudRetrieveTaskRequest) SetRetrieveAllJobs(v bool)`

SetRetrieveAllJobs sets RetrieveAllJobs field to given value.

### HasRetrieveAllJobs

`func (o *CreateCloudRetrieveTaskRequest) HasRetrieveAllJobs() bool`

HasRetrieveAllJobs returns a boolean if a field has been set.

### SetRetrieveAllJobsNil

`func (o *CreateCloudRetrieveTaskRequest) SetRetrieveAllJobsNil(b bool)`

 SetRetrieveAllJobsNil sets the value for RetrieveAllJobs to be an explicit nil

### UnsetRetrieveAllJobs
`func (o *CreateCloudRetrieveTaskRequest) UnsetRetrieveAllJobs()`

UnsetRetrieveAllJobs ensures that no value is present for RetrieveAllJobs, not even an explicit nil
### GetStartTimeUsecs

`func (o *CreateCloudRetrieveTaskRequest) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *CreateCloudRetrieveTaskRequest) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *CreateCloudRetrieveTaskRequest) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *CreateCloudRetrieveTaskRequest) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *CreateCloudRetrieveTaskRequest) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *CreateCloudRetrieveTaskRequest) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetVaultIds

`func (o *CreateCloudRetrieveTaskRequest) GetVaultIds() []int64`

GetVaultIds returns the VaultIds field if non-nil, zero value otherwise.

### GetVaultIdsOk

`func (o *CreateCloudRetrieveTaskRequest) GetVaultIdsOk() (*[]int64, bool)`

GetVaultIdsOk returns a tuple with the VaultIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultIds

`func (o *CreateCloudRetrieveTaskRequest) SetVaultIds(v []int64)`

SetVaultIds sets VaultIds field to given value.

### HasVaultIds

`func (o *CreateCloudRetrieveTaskRequest) HasVaultIds() bool`

HasVaultIds returns a boolean if a field has been set.

### SetVaultIdsNil

`func (o *CreateCloudRetrieveTaskRequest) SetVaultIdsNil(b bool)`

 SetVaultIdsNil sets the value for VaultIds to be an explicit nil

### UnsetVaultIds
`func (o *CreateCloudRetrieveTaskRequest) UnsetVaultIds()`

UnsetVaultIds ensures that no value is present for VaultIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


