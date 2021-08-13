# TdmTasks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tasks** | Pointer to [**[]TdmTask**](TdmTask.md) | Specifies the collection of TDM tasks, filtered by the specified criteria. | [optional] 
**PaginationInfo** | Pointer to [**PaginationInfo**](PaginationInfo.md) |  | [optional] 

## Methods

### NewTdmTasks

`func NewTdmTasks() *TdmTasks`

NewTdmTasks instantiates a new TdmTasks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTdmTasksWithDefaults

`func NewTdmTasksWithDefaults() *TdmTasks`

NewTdmTasksWithDefaults instantiates a new TdmTasks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTasks

`func (o *TdmTasks) GetTasks() []TdmTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *TdmTasks) GetTasksOk() (*[]TdmTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *TdmTasks) SetTasks(v []TdmTask)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *TdmTasks) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### SetTasksNil

`func (o *TdmTasks) SetTasksNil(b bool)`

 SetTasksNil sets the value for Tasks to be an explicit nil

### UnsetTasks
`func (o *TdmTasks) UnsetTasks()`

UnsetTasks ensures that no value is present for Tasks, not even an explicit nil
### GetPaginationInfo

`func (o *TdmTasks) GetPaginationInfo() PaginationInfo`

GetPaginationInfo returns the PaginationInfo field if non-nil, zero value otherwise.

### GetPaginationInfoOk

`func (o *TdmTasks) GetPaginationInfoOk() (*PaginationInfo, bool)`

GetPaginationInfoOk returns a tuple with the PaginationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaginationInfo

`func (o *TdmTasks) SetPaginationInfo(v PaginationInfo)`

SetPaginationInfo sets PaginationInfo field to given value.

### HasPaginationInfo

`func (o *TdmTasks) HasPaginationInfo() bool`

HasPaginationInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


