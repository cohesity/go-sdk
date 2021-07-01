# BasicTaskInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Name of the recovery task. | [optional] 
**TaskId** | Pointer to **NullableString** | Id of the recovery task. | [optional] 

## Methods

### NewBasicTaskInfo

`func NewBasicTaskInfo() *BasicTaskInfo`

NewBasicTaskInfo instantiates a new BasicTaskInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasicTaskInfoWithDefaults

`func NewBasicTaskInfoWithDefaults() *BasicTaskInfo`

NewBasicTaskInfoWithDefaults instantiates a new BasicTaskInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BasicTaskInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BasicTaskInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BasicTaskInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BasicTaskInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BasicTaskInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BasicTaskInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTaskId

`func (o *BasicTaskInfo) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *BasicTaskInfo) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *BasicTaskInfo) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *BasicTaskInfo) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### SetTaskIdNil

`func (o *BasicTaskInfo) SetTaskIdNil(b bool)`

 SetTaskIdNil sets the value for TaskId to be an explicit nil

### UnsetTaskId
`func (o *BasicTaskInfo) UnsetTaskId()`

UnsetTaskId ensures that no value is present for TaskId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


