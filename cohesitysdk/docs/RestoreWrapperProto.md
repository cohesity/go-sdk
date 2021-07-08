# RestoreWrapperProto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestroyClonedTaskStateVec** | Pointer to [**[]DestroyClonedTaskStateProto**](DestroyClonedTaskStateProto.md) | For a restore task of type &#39;Clone&#39;, this field contains the info of the destroy task(s). | [optional] 
**OwnerRestoreWrapperProto** | Pointer to [**RestoreWrapperProto**](RestoreWrapperProto.md) |  | [optional] 
**PerformRefreshTaskStateVec** | Pointer to [**[]PerformRestoreTaskStateProto**](PerformRestoreTaskStateProto.md) | Contains information of the refresh tasks for a clone | [optional] 
**PerformRestoreJobState** | Pointer to [**PerformRestoreJobStateProto**](PerformRestoreJobStateProto.md) |  | [optional] 
**PerformRestoreTaskState** | Pointer to [**PerformRestoreTaskStateProto**](PerformRestoreTaskStateProto.md) |  | [optional] 
**RestoreSubTaskWrapperProtoVec** | Pointer to **[]map[string]interface{}** | If this restore has sub tasks, the following field will get populated with the wrapper proto of all of its sub-tasks.  Note that this field is only populated for Iris in response to &#39;GetRestoreTasksArg&#39; RPC. It is not persisted in Magneto&#39;s WAL.  List of environments that use this field: kSQL : Used for multi-stage SQL restore that supports a hot-standy. | [optional] 

## Methods

### NewRestoreWrapperProto

`func NewRestoreWrapperProto() *RestoreWrapperProto`

NewRestoreWrapperProto instantiates a new RestoreWrapperProto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreWrapperProtoWithDefaults

`func NewRestoreWrapperProtoWithDefaults() *RestoreWrapperProto`

NewRestoreWrapperProtoWithDefaults instantiates a new RestoreWrapperProto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestroyClonedTaskStateVec

`func (o *RestoreWrapperProto) GetDestroyClonedTaskStateVec() []DestroyClonedTaskStateProto`

GetDestroyClonedTaskStateVec returns the DestroyClonedTaskStateVec field if non-nil, zero value otherwise.

### GetDestroyClonedTaskStateVecOk

`func (o *RestoreWrapperProto) GetDestroyClonedTaskStateVecOk() (*[]DestroyClonedTaskStateProto, bool)`

GetDestroyClonedTaskStateVecOk returns a tuple with the DestroyClonedTaskStateVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroyClonedTaskStateVec

`func (o *RestoreWrapperProto) SetDestroyClonedTaskStateVec(v []DestroyClonedTaskStateProto)`

SetDestroyClonedTaskStateVec sets DestroyClonedTaskStateVec field to given value.

### HasDestroyClonedTaskStateVec

`func (o *RestoreWrapperProto) HasDestroyClonedTaskStateVec() bool`

HasDestroyClonedTaskStateVec returns a boolean if a field has been set.

### SetDestroyClonedTaskStateVecNil

`func (o *RestoreWrapperProto) SetDestroyClonedTaskStateVecNil(b bool)`

 SetDestroyClonedTaskStateVecNil sets the value for DestroyClonedTaskStateVec to be an explicit nil

### UnsetDestroyClonedTaskStateVec
`func (o *RestoreWrapperProto) UnsetDestroyClonedTaskStateVec()`

UnsetDestroyClonedTaskStateVec ensures that no value is present for DestroyClonedTaskStateVec, not even an explicit nil
### GetOwnerRestoreWrapperProto

`func (o *RestoreWrapperProto) GetOwnerRestoreWrapperProto() RestoreWrapperProto`

GetOwnerRestoreWrapperProto returns the OwnerRestoreWrapperProto field if non-nil, zero value otherwise.

### GetOwnerRestoreWrapperProtoOk

`func (o *RestoreWrapperProto) GetOwnerRestoreWrapperProtoOk() (*RestoreWrapperProto, bool)`

GetOwnerRestoreWrapperProtoOk returns a tuple with the OwnerRestoreWrapperProto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerRestoreWrapperProto

`func (o *RestoreWrapperProto) SetOwnerRestoreWrapperProto(v RestoreWrapperProto)`

SetOwnerRestoreWrapperProto sets OwnerRestoreWrapperProto field to given value.

### HasOwnerRestoreWrapperProto

`func (o *RestoreWrapperProto) HasOwnerRestoreWrapperProto() bool`

HasOwnerRestoreWrapperProto returns a boolean if a field has been set.

### GetPerformRefreshTaskStateVec

`func (o *RestoreWrapperProto) GetPerformRefreshTaskStateVec() []PerformRestoreTaskStateProto`

GetPerformRefreshTaskStateVec returns the PerformRefreshTaskStateVec field if non-nil, zero value otherwise.

### GetPerformRefreshTaskStateVecOk

`func (o *RestoreWrapperProto) GetPerformRefreshTaskStateVecOk() (*[]PerformRestoreTaskStateProto, bool)`

GetPerformRefreshTaskStateVecOk returns a tuple with the PerformRefreshTaskStateVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformRefreshTaskStateVec

`func (o *RestoreWrapperProto) SetPerformRefreshTaskStateVec(v []PerformRestoreTaskStateProto)`

SetPerformRefreshTaskStateVec sets PerformRefreshTaskStateVec field to given value.

### HasPerformRefreshTaskStateVec

`func (o *RestoreWrapperProto) HasPerformRefreshTaskStateVec() bool`

HasPerformRefreshTaskStateVec returns a boolean if a field has been set.

### SetPerformRefreshTaskStateVecNil

`func (o *RestoreWrapperProto) SetPerformRefreshTaskStateVecNil(b bool)`

 SetPerformRefreshTaskStateVecNil sets the value for PerformRefreshTaskStateVec to be an explicit nil

### UnsetPerformRefreshTaskStateVec
`func (o *RestoreWrapperProto) UnsetPerformRefreshTaskStateVec()`

UnsetPerformRefreshTaskStateVec ensures that no value is present for PerformRefreshTaskStateVec, not even an explicit nil
### GetPerformRestoreJobState

`func (o *RestoreWrapperProto) GetPerformRestoreJobState() PerformRestoreJobStateProto`

GetPerformRestoreJobState returns the PerformRestoreJobState field if non-nil, zero value otherwise.

### GetPerformRestoreJobStateOk

`func (o *RestoreWrapperProto) GetPerformRestoreJobStateOk() (*PerformRestoreJobStateProto, bool)`

GetPerformRestoreJobStateOk returns a tuple with the PerformRestoreJobState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformRestoreJobState

`func (o *RestoreWrapperProto) SetPerformRestoreJobState(v PerformRestoreJobStateProto)`

SetPerformRestoreJobState sets PerformRestoreJobState field to given value.

### HasPerformRestoreJobState

`func (o *RestoreWrapperProto) HasPerformRestoreJobState() bool`

HasPerformRestoreJobState returns a boolean if a field has been set.

### GetPerformRestoreTaskState

`func (o *RestoreWrapperProto) GetPerformRestoreTaskState() PerformRestoreTaskStateProto`

GetPerformRestoreTaskState returns the PerformRestoreTaskState field if non-nil, zero value otherwise.

### GetPerformRestoreTaskStateOk

`func (o *RestoreWrapperProto) GetPerformRestoreTaskStateOk() (*PerformRestoreTaskStateProto, bool)`

GetPerformRestoreTaskStateOk returns a tuple with the PerformRestoreTaskState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformRestoreTaskState

`func (o *RestoreWrapperProto) SetPerformRestoreTaskState(v PerformRestoreTaskStateProto)`

SetPerformRestoreTaskState sets PerformRestoreTaskState field to given value.

### HasPerformRestoreTaskState

`func (o *RestoreWrapperProto) HasPerformRestoreTaskState() bool`

HasPerformRestoreTaskState returns a boolean if a field has been set.

### GetRestoreSubTaskWrapperProtoVec

`func (o *RestoreWrapperProto) GetRestoreSubTaskWrapperProtoVec() []map[string]interface{}`

GetRestoreSubTaskWrapperProtoVec returns the RestoreSubTaskWrapperProtoVec field if non-nil, zero value otherwise.

### GetRestoreSubTaskWrapperProtoVecOk

`func (o *RestoreWrapperProto) GetRestoreSubTaskWrapperProtoVecOk() (*[]map[string]interface{}, bool)`

GetRestoreSubTaskWrapperProtoVecOk returns a tuple with the RestoreSubTaskWrapperProtoVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreSubTaskWrapperProtoVec

`func (o *RestoreWrapperProto) SetRestoreSubTaskWrapperProtoVec(v []map[string]interface{})`

SetRestoreSubTaskWrapperProtoVec sets RestoreSubTaskWrapperProtoVec field to given value.

### HasRestoreSubTaskWrapperProtoVec

`func (o *RestoreWrapperProto) HasRestoreSubTaskWrapperProtoVec() bool`

HasRestoreSubTaskWrapperProtoVec returns a boolean if a field has been set.

### SetRestoreSubTaskWrapperProtoVecNil

`func (o *RestoreWrapperProto) SetRestoreSubTaskWrapperProtoVecNil(b bool)`

 SetRestoreSubTaskWrapperProtoVecNil sets the value for RestoreSubTaskWrapperProtoVec to be an explicit nil

### UnsetRestoreSubTaskWrapperProtoVec
`func (o *RestoreWrapperProto) UnsetRestoreSubTaskWrapperProtoVec()`

UnsetRestoreSubTaskWrapperProtoVec ensures that no value is present for RestoreSubTaskWrapperProtoVec, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


