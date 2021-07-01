# InfectedFileId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | Pointer to **NullableInt64** | Specifies the entity id of the infected file. | [optional] 
**RootInodeId** | Pointer to **NullableInt64** | Specifies the root inode id of the file system that infected file belongs to. | [optional] 
**ViewId** | Pointer to **NullableInt64** | Specifies the id of the View the infected file belongs to. | [optional] 

## Methods

### NewInfectedFileId

`func NewInfectedFileId() *InfectedFileId`

NewInfectedFileId instantiates a new InfectedFileId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfectedFileIdWithDefaults

`func NewInfectedFileIdWithDefaults() *InfectedFileId`

NewInfectedFileIdWithDefaults instantiates a new InfectedFileId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *InfectedFileId) GetEntityId() int64`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *InfectedFileId) GetEntityIdOk() (*int64, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *InfectedFileId) SetEntityId(v int64)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *InfectedFileId) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### SetEntityIdNil

`func (o *InfectedFileId) SetEntityIdNil(b bool)`

 SetEntityIdNil sets the value for EntityId to be an explicit nil

### UnsetEntityId
`func (o *InfectedFileId) UnsetEntityId()`

UnsetEntityId ensures that no value is present for EntityId, not even an explicit nil
### GetRootInodeId

`func (o *InfectedFileId) GetRootInodeId() int64`

GetRootInodeId returns the RootInodeId field if non-nil, zero value otherwise.

### GetRootInodeIdOk

`func (o *InfectedFileId) GetRootInodeIdOk() (*int64, bool)`

GetRootInodeIdOk returns a tuple with the RootInodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootInodeId

`func (o *InfectedFileId) SetRootInodeId(v int64)`

SetRootInodeId sets RootInodeId field to given value.

### HasRootInodeId

`func (o *InfectedFileId) HasRootInodeId() bool`

HasRootInodeId returns a boolean if a field has been set.

### SetRootInodeIdNil

`func (o *InfectedFileId) SetRootInodeIdNil(b bool)`

 SetRootInodeIdNil sets the value for RootInodeId to be an explicit nil

### UnsetRootInodeId
`func (o *InfectedFileId) UnsetRootInodeId()`

UnsetRootInodeId ensures that no value is present for RootInodeId, not even an explicit nil
### GetViewId

`func (o *InfectedFileId) GetViewId() int64`

GetViewId returns the ViewId field if non-nil, zero value otherwise.

### GetViewIdOk

`func (o *InfectedFileId) GetViewIdOk() (*int64, bool)`

GetViewIdOk returns a tuple with the ViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewId

`func (o *InfectedFileId) SetViewId(v int64)`

SetViewId sets ViewId field to given value.

### HasViewId

`func (o *InfectedFileId) HasViewId() bool`

HasViewId returns a boolean if a field has been set.

### SetViewIdNil

`func (o *InfectedFileId) SetViewIdNil(b bool)`

 SetViewIdNil sets the value for ViewId to be an explicit nil

### UnsetViewId
`func (o *InfectedFileId) UnsetViewId()`

UnsetViewId ensures that no value is present for ViewId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


