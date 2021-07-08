# RecoveryTaskInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Name of the recovery task. | [optional] 
**TaskId** | Pointer to **NullableString** | Id of the recovery task. | [optional] 
**Type** | Pointer to **NullableString** | Denotes if the recovery task has an archival target. This param is used to reflect if the recovery op has an archival target to work with. &#39;local&#39; indicates no archival target. &#39;archive&#39; indicates that objects restored using an archival target. | [optional] 

## Methods

### NewRecoveryTaskInfo

`func NewRecoveryTaskInfo() *RecoveryTaskInfo`

NewRecoveryTaskInfo instantiates a new RecoveryTaskInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryTaskInfoWithDefaults

`func NewRecoveryTaskInfoWithDefaults() *RecoveryTaskInfo`

NewRecoveryTaskInfoWithDefaults instantiates a new RecoveryTaskInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RecoveryTaskInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecoveryTaskInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecoveryTaskInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RecoveryTaskInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *RecoveryTaskInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *RecoveryTaskInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTaskId

`func (o *RecoveryTaskInfo) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *RecoveryTaskInfo) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *RecoveryTaskInfo) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *RecoveryTaskInfo) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### SetTaskIdNil

`func (o *RecoveryTaskInfo) SetTaskIdNil(b bool)`

 SetTaskIdNil sets the value for TaskId to be an explicit nil

### UnsetTaskId
`func (o *RecoveryTaskInfo) UnsetTaskId()`

UnsetTaskId ensures that no value is present for TaskId, not even an explicit nil
### GetType

`func (o *RecoveryTaskInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RecoveryTaskInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RecoveryTaskInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RecoveryTaskInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *RecoveryTaskInfo) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *RecoveryTaskInfo) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


