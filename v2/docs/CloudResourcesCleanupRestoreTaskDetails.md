# CloudResourcesCleanupRestoreTaskDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmiInstanceIds** | Pointer to [**[]CloudResourcesCleanupAmiInstanceId**](CloudResourcesCleanupAmiInstanceId.md) | Specifies the mapping between AMI ID and the EC2 instance created. | [optional] 
**TaskId** | **NullableFloat32** | Specifies the ID of the restore task which was used to spawn the clones. | 

## Methods

### NewCloudResourcesCleanupRestoreTaskDetails

`func NewCloudResourcesCleanupRestoreTaskDetails(taskId NullableFloat32, ) *CloudResourcesCleanupRestoreTaskDetails`

NewCloudResourcesCleanupRestoreTaskDetails instantiates a new CloudResourcesCleanupRestoreTaskDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudResourcesCleanupRestoreTaskDetailsWithDefaults

`func NewCloudResourcesCleanupRestoreTaskDetailsWithDefaults() *CloudResourcesCleanupRestoreTaskDetails`

NewCloudResourcesCleanupRestoreTaskDetailsWithDefaults instantiates a new CloudResourcesCleanupRestoreTaskDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmiInstanceIds

`func (o *CloudResourcesCleanupRestoreTaskDetails) GetAmiInstanceIds() []CloudResourcesCleanupAmiInstanceId`

GetAmiInstanceIds returns the AmiInstanceIds field if non-nil, zero value otherwise.

### GetAmiInstanceIdsOk

`func (o *CloudResourcesCleanupRestoreTaskDetails) GetAmiInstanceIdsOk() (*[]CloudResourcesCleanupAmiInstanceId, bool)`

GetAmiInstanceIdsOk returns a tuple with the AmiInstanceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmiInstanceIds

`func (o *CloudResourcesCleanupRestoreTaskDetails) SetAmiInstanceIds(v []CloudResourcesCleanupAmiInstanceId)`

SetAmiInstanceIds sets AmiInstanceIds field to given value.

### HasAmiInstanceIds

`func (o *CloudResourcesCleanupRestoreTaskDetails) HasAmiInstanceIds() bool`

HasAmiInstanceIds returns a boolean if a field has been set.

### GetTaskId

`func (o *CloudResourcesCleanupRestoreTaskDetails) GetTaskId() float32`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *CloudResourcesCleanupRestoreTaskDetails) GetTaskIdOk() (*float32, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *CloudResourcesCleanupRestoreTaskDetails) SetTaskId(v float32)`

SetTaskId sets TaskId field to given value.


### SetTaskIdNil

`func (o *CloudResourcesCleanupRestoreTaskDetails) SetTaskIdNil(b bool)`

 SetTaskIdNil sets the value for TaskId to be an explicit nil

### UnsetTaskId
`func (o *CloudResourcesCleanupRestoreTaskDetails) UnsetTaskId()`

UnsetTaskId ensures that no value is present for TaskId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


