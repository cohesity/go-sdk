# RestoreInfoProto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PostScriptStatus** | Pointer to [**ScriptExecutionStatus**](ScriptExecutionStatus.md) |  | [optional] 
**PreScriptStatus** | Pointer to [**ScriptExecutionStatus**](ScriptExecutionStatus.md) |  | [optional] 
**RestoreEntityVec** | Pointer to [**[]RestoreInfoProtoRestoreEntity**](RestoreInfoProtoRestoreEntity.md) | Contains the file paths and the information of the restored entities. | [optional] 
**TargetType** | Pointer to **NullableInt32** | Specifies the target type for the task. The field is only valid if the task has got a permit. | [optional] 
**Type** | Pointer to **NullableInt32** | The type of environment this restore info pertains to. | [optional] 

## Methods

### NewRestoreInfoProto

`func NewRestoreInfoProto() *RestoreInfoProto`

NewRestoreInfoProto instantiates a new RestoreInfoProto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreInfoProtoWithDefaults

`func NewRestoreInfoProtoWithDefaults() *RestoreInfoProto`

NewRestoreInfoProtoWithDefaults instantiates a new RestoreInfoProto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPostScriptStatus

`func (o *RestoreInfoProto) GetPostScriptStatus() ScriptExecutionStatus`

GetPostScriptStatus returns the PostScriptStatus field if non-nil, zero value otherwise.

### GetPostScriptStatusOk

`func (o *RestoreInfoProto) GetPostScriptStatusOk() (*ScriptExecutionStatus, bool)`

GetPostScriptStatusOk returns a tuple with the PostScriptStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostScriptStatus

`func (o *RestoreInfoProto) SetPostScriptStatus(v ScriptExecutionStatus)`

SetPostScriptStatus sets PostScriptStatus field to given value.

### HasPostScriptStatus

`func (o *RestoreInfoProto) HasPostScriptStatus() bool`

HasPostScriptStatus returns a boolean if a field has been set.

### GetPreScriptStatus

`func (o *RestoreInfoProto) GetPreScriptStatus() ScriptExecutionStatus`

GetPreScriptStatus returns the PreScriptStatus field if non-nil, zero value otherwise.

### GetPreScriptStatusOk

`func (o *RestoreInfoProto) GetPreScriptStatusOk() (*ScriptExecutionStatus, bool)`

GetPreScriptStatusOk returns a tuple with the PreScriptStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreScriptStatus

`func (o *RestoreInfoProto) SetPreScriptStatus(v ScriptExecutionStatus)`

SetPreScriptStatus sets PreScriptStatus field to given value.

### HasPreScriptStatus

`func (o *RestoreInfoProto) HasPreScriptStatus() bool`

HasPreScriptStatus returns a boolean if a field has been set.

### GetRestoreEntityVec

`func (o *RestoreInfoProto) GetRestoreEntityVec() []RestoreInfoProtoRestoreEntity`

GetRestoreEntityVec returns the RestoreEntityVec field if non-nil, zero value otherwise.

### GetRestoreEntityVecOk

`func (o *RestoreInfoProto) GetRestoreEntityVecOk() (*[]RestoreInfoProtoRestoreEntity, bool)`

GetRestoreEntityVecOk returns a tuple with the RestoreEntityVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreEntityVec

`func (o *RestoreInfoProto) SetRestoreEntityVec(v []RestoreInfoProtoRestoreEntity)`

SetRestoreEntityVec sets RestoreEntityVec field to given value.

### HasRestoreEntityVec

`func (o *RestoreInfoProto) HasRestoreEntityVec() bool`

HasRestoreEntityVec returns a boolean if a field has been set.

### SetRestoreEntityVecNil

`func (o *RestoreInfoProto) SetRestoreEntityVecNil(b bool)`

 SetRestoreEntityVecNil sets the value for RestoreEntityVec to be an explicit nil

### UnsetRestoreEntityVec
`func (o *RestoreInfoProto) UnsetRestoreEntityVec()`

UnsetRestoreEntityVec ensures that no value is present for RestoreEntityVec, not even an explicit nil
### GetTargetType

`func (o *RestoreInfoProto) GetTargetType() int32`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *RestoreInfoProto) GetTargetTypeOk() (*int32, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *RestoreInfoProto) SetTargetType(v int32)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *RestoreInfoProto) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### SetTargetTypeNil

`func (o *RestoreInfoProto) SetTargetTypeNil(b bool)`

 SetTargetTypeNil sets the value for TargetType to be an explicit nil

### UnsetTargetType
`func (o *RestoreInfoProto) UnsetTargetType()`

UnsetTargetType ensures that no value is present for TargetType, not even an explicit nil
### GetType

`func (o *RestoreInfoProto) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RestoreInfoProto) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RestoreInfoProto) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *RestoreInfoProto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *RestoreInfoProto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *RestoreInfoProto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


