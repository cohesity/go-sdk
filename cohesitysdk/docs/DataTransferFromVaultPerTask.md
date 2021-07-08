# DataTransferFromVaultPerTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumLogicalBytesTransferred** | Pointer to **NullableInt64** | Specifies the total number of logical bytes that are transferred from this Vault to the Cohesity Cluster for this task. The logical size is when the data is fully hydrated or expanded. | [optional] 
**NumPhysicalBytesTransferred** | Pointer to **NullableInt64** | Specifies the total number of physical bytes that are transferred from this Vault to the Cohesity Cluster for this task. | [optional] 
**TaskName** | Pointer to **NullableString** | Specifies the task name. | [optional] 
**TaskType** | Pointer to **NullableString** | Specifies the task type. | [optional] 

## Methods

### NewDataTransferFromVaultPerTask

`func NewDataTransferFromVaultPerTask() *DataTransferFromVaultPerTask`

NewDataTransferFromVaultPerTask instantiates a new DataTransferFromVaultPerTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataTransferFromVaultPerTaskWithDefaults

`func NewDataTransferFromVaultPerTaskWithDefaults() *DataTransferFromVaultPerTask`

NewDataTransferFromVaultPerTaskWithDefaults instantiates a new DataTransferFromVaultPerTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumLogicalBytesTransferred

`func (o *DataTransferFromVaultPerTask) GetNumLogicalBytesTransferred() int64`

GetNumLogicalBytesTransferred returns the NumLogicalBytesTransferred field if non-nil, zero value otherwise.

### GetNumLogicalBytesTransferredOk

`func (o *DataTransferFromVaultPerTask) GetNumLogicalBytesTransferredOk() (*int64, bool)`

GetNumLogicalBytesTransferredOk returns a tuple with the NumLogicalBytesTransferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumLogicalBytesTransferred

`func (o *DataTransferFromVaultPerTask) SetNumLogicalBytesTransferred(v int64)`

SetNumLogicalBytesTransferred sets NumLogicalBytesTransferred field to given value.

### HasNumLogicalBytesTransferred

`func (o *DataTransferFromVaultPerTask) HasNumLogicalBytesTransferred() bool`

HasNumLogicalBytesTransferred returns a boolean if a field has been set.

### SetNumLogicalBytesTransferredNil

`func (o *DataTransferFromVaultPerTask) SetNumLogicalBytesTransferredNil(b bool)`

 SetNumLogicalBytesTransferredNil sets the value for NumLogicalBytesTransferred to be an explicit nil

### UnsetNumLogicalBytesTransferred
`func (o *DataTransferFromVaultPerTask) UnsetNumLogicalBytesTransferred()`

UnsetNumLogicalBytesTransferred ensures that no value is present for NumLogicalBytesTransferred, not even an explicit nil
### GetNumPhysicalBytesTransferred

`func (o *DataTransferFromVaultPerTask) GetNumPhysicalBytesTransferred() int64`

GetNumPhysicalBytesTransferred returns the NumPhysicalBytesTransferred field if non-nil, zero value otherwise.

### GetNumPhysicalBytesTransferredOk

`func (o *DataTransferFromVaultPerTask) GetNumPhysicalBytesTransferredOk() (*int64, bool)`

GetNumPhysicalBytesTransferredOk returns a tuple with the NumPhysicalBytesTransferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPhysicalBytesTransferred

`func (o *DataTransferFromVaultPerTask) SetNumPhysicalBytesTransferred(v int64)`

SetNumPhysicalBytesTransferred sets NumPhysicalBytesTransferred field to given value.

### HasNumPhysicalBytesTransferred

`func (o *DataTransferFromVaultPerTask) HasNumPhysicalBytesTransferred() bool`

HasNumPhysicalBytesTransferred returns a boolean if a field has been set.

### SetNumPhysicalBytesTransferredNil

`func (o *DataTransferFromVaultPerTask) SetNumPhysicalBytesTransferredNil(b bool)`

 SetNumPhysicalBytesTransferredNil sets the value for NumPhysicalBytesTransferred to be an explicit nil

### UnsetNumPhysicalBytesTransferred
`func (o *DataTransferFromVaultPerTask) UnsetNumPhysicalBytesTransferred()`

UnsetNumPhysicalBytesTransferred ensures that no value is present for NumPhysicalBytesTransferred, not even an explicit nil
### GetTaskName

`func (o *DataTransferFromVaultPerTask) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *DataTransferFromVaultPerTask) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *DataTransferFromVaultPerTask) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.

### HasTaskName

`func (o *DataTransferFromVaultPerTask) HasTaskName() bool`

HasTaskName returns a boolean if a field has been set.

### SetTaskNameNil

`func (o *DataTransferFromVaultPerTask) SetTaskNameNil(b bool)`

 SetTaskNameNil sets the value for TaskName to be an explicit nil

### UnsetTaskName
`func (o *DataTransferFromVaultPerTask) UnsetTaskName()`

UnsetTaskName ensures that no value is present for TaskName, not even an explicit nil
### GetTaskType

`func (o *DataTransferFromVaultPerTask) GetTaskType() string`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *DataTransferFromVaultPerTask) GetTaskTypeOk() (*string, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *DataTransferFromVaultPerTask) SetTaskType(v string)`

SetTaskType sets TaskType field to given value.

### HasTaskType

`func (o *DataTransferFromVaultPerTask) HasTaskType() bool`

HasTaskType returns a boolean if a field has been set.

### SetTaskTypeNil

`func (o *DataTransferFromVaultPerTask) SetTaskTypeNil(b bool)`

 SetTaskTypeNil sets the value for TaskType to be an explicit nil

### UnsetTaskType
`func (o *DataTransferFromVaultPerTask) UnsetTaskType()`

UnsetTaskType ensures that no value is present for TaskType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


