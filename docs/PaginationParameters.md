# PaginationParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AfterCursorEntityId** | Pointer to **NullableInt64** | Specifies the entity id starting from which the items are to be returned. | [optional] 
**BeforeCursorEntityId** | Pointer to **NullableInt64** | Specifies the entity id upto which the items are to be returned. | [optional] 
**NodeId** | Pointer to **NullableInt64** | Specifies the entity id for the Node at any level within the Source entity hierarchy whose children are to be paginated. | [optional] 
**PageSize** | Pointer to **NullableInt64** | Specifies the maximum number of entities to be returned within the page. | [optional] 

## Methods

### NewPaginationParameters

`func NewPaginationParameters() *PaginationParameters`

NewPaginationParameters instantiates a new PaginationParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationParametersWithDefaults

`func NewPaginationParametersWithDefaults() *PaginationParameters`

NewPaginationParametersWithDefaults instantiates a new PaginationParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfterCursorEntityId

`func (o *PaginationParameters) GetAfterCursorEntityId() int64`

GetAfterCursorEntityId returns the AfterCursorEntityId field if non-nil, zero value otherwise.

### GetAfterCursorEntityIdOk

`func (o *PaginationParameters) GetAfterCursorEntityIdOk() (*int64, bool)`

GetAfterCursorEntityIdOk returns a tuple with the AfterCursorEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterCursorEntityId

`func (o *PaginationParameters) SetAfterCursorEntityId(v int64)`

SetAfterCursorEntityId sets AfterCursorEntityId field to given value.

### HasAfterCursorEntityId

`func (o *PaginationParameters) HasAfterCursorEntityId() bool`

HasAfterCursorEntityId returns a boolean if a field has been set.

### SetAfterCursorEntityIdNil

`func (o *PaginationParameters) SetAfterCursorEntityIdNil(b bool)`

 SetAfterCursorEntityIdNil sets the value for AfterCursorEntityId to be an explicit nil

### UnsetAfterCursorEntityId
`func (o *PaginationParameters) UnsetAfterCursorEntityId()`

UnsetAfterCursorEntityId ensures that no value is present for AfterCursorEntityId, not even an explicit nil
### GetBeforeCursorEntityId

`func (o *PaginationParameters) GetBeforeCursorEntityId() int64`

GetBeforeCursorEntityId returns the BeforeCursorEntityId field if non-nil, zero value otherwise.

### GetBeforeCursorEntityIdOk

`func (o *PaginationParameters) GetBeforeCursorEntityIdOk() (*int64, bool)`

GetBeforeCursorEntityIdOk returns a tuple with the BeforeCursorEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeforeCursorEntityId

`func (o *PaginationParameters) SetBeforeCursorEntityId(v int64)`

SetBeforeCursorEntityId sets BeforeCursorEntityId field to given value.

### HasBeforeCursorEntityId

`func (o *PaginationParameters) HasBeforeCursorEntityId() bool`

HasBeforeCursorEntityId returns a boolean if a field has been set.

### SetBeforeCursorEntityIdNil

`func (o *PaginationParameters) SetBeforeCursorEntityIdNil(b bool)`

 SetBeforeCursorEntityIdNil sets the value for BeforeCursorEntityId to be an explicit nil

### UnsetBeforeCursorEntityId
`func (o *PaginationParameters) UnsetBeforeCursorEntityId()`

UnsetBeforeCursorEntityId ensures that no value is present for BeforeCursorEntityId, not even an explicit nil
### GetNodeId

`func (o *PaginationParameters) GetNodeId() int64`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *PaginationParameters) GetNodeIdOk() (*int64, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *PaginationParameters) SetNodeId(v int64)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *PaginationParameters) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### SetNodeIdNil

`func (o *PaginationParameters) SetNodeIdNil(b bool)`

 SetNodeIdNil sets the value for NodeId to be an explicit nil

### UnsetNodeId
`func (o *PaginationParameters) UnsetNodeId()`

UnsetNodeId ensures that no value is present for NodeId, not even an explicit nil
### GetPageSize

`func (o *PaginationParameters) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *PaginationParameters) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *PaginationParameters) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *PaginationParameters) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### SetPageSizeNil

`func (o *PaginationParameters) SetPageSizeNil(b bool)`

 SetPageSizeNil sets the value for PageSize to be an explicit nil

### UnsetPageSize
`func (o *PaginationParameters) UnsetPageSize()`

UnsetPageSize ensures that no value is present for PageSize, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


