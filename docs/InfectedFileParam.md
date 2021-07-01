# InfectedFileParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | Pointer to **NullableInt64** | Specifies the entity id of the infected file. | [optional] 
**RemediationState** | Pointer to **NullableString** | Specifies the remediation state of the file. Remediation State. &#39;kQuarantine&#39; indicates &#39;Quarantine&#39; state of the file. This state blocks the client access. The administrator will have to manually delete, rescan or unquarantine the file. &#39;kUnquarantine&#39; indicates &#39;Unquarantine&#39; state of the file. The administrator has manually moved files from quarantined to the unquarantined state to allow client access. Unquarantined files are not scanned for virus until manually reset. | [optional] 
**RootInodeId** | Pointer to **NullableInt64** | Specifies the root inode id of the file system that infected file belongs to. | [optional] 
**ViewId** | Pointer to **NullableInt64** | Specifies the id of the View the infected file belongs to. | [optional] 

## Methods

### NewInfectedFileParam

`func NewInfectedFileParam() *InfectedFileParam`

NewInfectedFileParam instantiates a new InfectedFileParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfectedFileParamWithDefaults

`func NewInfectedFileParamWithDefaults() *InfectedFileParam`

NewInfectedFileParamWithDefaults instantiates a new InfectedFileParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *InfectedFileParam) GetEntityId() int64`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *InfectedFileParam) GetEntityIdOk() (*int64, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *InfectedFileParam) SetEntityId(v int64)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *InfectedFileParam) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### SetEntityIdNil

`func (o *InfectedFileParam) SetEntityIdNil(b bool)`

 SetEntityIdNil sets the value for EntityId to be an explicit nil

### UnsetEntityId
`func (o *InfectedFileParam) UnsetEntityId()`

UnsetEntityId ensures that no value is present for EntityId, not even an explicit nil
### GetRemediationState

`func (o *InfectedFileParam) GetRemediationState() string`

GetRemediationState returns the RemediationState field if non-nil, zero value otherwise.

### GetRemediationStateOk

`func (o *InfectedFileParam) GetRemediationStateOk() (*string, bool)`

GetRemediationStateOk returns a tuple with the RemediationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationState

`func (o *InfectedFileParam) SetRemediationState(v string)`

SetRemediationState sets RemediationState field to given value.

### HasRemediationState

`func (o *InfectedFileParam) HasRemediationState() bool`

HasRemediationState returns a boolean if a field has been set.

### SetRemediationStateNil

`func (o *InfectedFileParam) SetRemediationStateNil(b bool)`

 SetRemediationStateNil sets the value for RemediationState to be an explicit nil

### UnsetRemediationState
`func (o *InfectedFileParam) UnsetRemediationState()`

UnsetRemediationState ensures that no value is present for RemediationState, not even an explicit nil
### GetRootInodeId

`func (o *InfectedFileParam) GetRootInodeId() int64`

GetRootInodeId returns the RootInodeId field if non-nil, zero value otherwise.

### GetRootInodeIdOk

`func (o *InfectedFileParam) GetRootInodeIdOk() (*int64, bool)`

GetRootInodeIdOk returns a tuple with the RootInodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootInodeId

`func (o *InfectedFileParam) SetRootInodeId(v int64)`

SetRootInodeId sets RootInodeId field to given value.

### HasRootInodeId

`func (o *InfectedFileParam) HasRootInodeId() bool`

HasRootInodeId returns a boolean if a field has been set.

### SetRootInodeIdNil

`func (o *InfectedFileParam) SetRootInodeIdNil(b bool)`

 SetRootInodeIdNil sets the value for RootInodeId to be an explicit nil

### UnsetRootInodeId
`func (o *InfectedFileParam) UnsetRootInodeId()`

UnsetRootInodeId ensures that no value is present for RootInodeId, not even an explicit nil
### GetViewId

`func (o *InfectedFileParam) GetViewId() int64`

GetViewId returns the ViewId field if non-nil, zero value otherwise.

### GetViewIdOk

`func (o *InfectedFileParam) GetViewIdOk() (*int64, bool)`

GetViewIdOk returns a tuple with the ViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewId

`func (o *InfectedFileParam) SetViewId(v int64)`

SetViewId sets ViewId field to given value.

### HasViewId

`func (o *InfectedFileParam) HasViewId() bool`

HasViewId returns a boolean if a field has been set.

### SetViewIdNil

`func (o *InfectedFileParam) SetViewIdNil(b bool)`

 SetViewIdNil sets the value for ViewId to be an explicit nil

### UnsetViewId
`func (o *InfectedFileParam) UnsetViewId()`

UnsetViewId ensures that no value is present for ViewId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


