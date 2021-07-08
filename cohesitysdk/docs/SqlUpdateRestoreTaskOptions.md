# SqlUpdateRestoreTaskOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableAutoSync** | Pointer to **NullableBool** | Enable/Disable auto_sync for db migration | [optional] 
**MultiStageRestoreAction** | Pointer to **NullableInt32** | This field is set if we are performing an action on a multi-stage SQL restore. | [optional] 

## Methods

### NewSqlUpdateRestoreTaskOptions

`func NewSqlUpdateRestoreTaskOptions() *SqlUpdateRestoreTaskOptions`

NewSqlUpdateRestoreTaskOptions instantiates a new SqlUpdateRestoreTaskOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlUpdateRestoreTaskOptionsWithDefaults

`func NewSqlUpdateRestoreTaskOptionsWithDefaults() *SqlUpdateRestoreTaskOptions`

NewSqlUpdateRestoreTaskOptionsWithDefaults instantiates a new SqlUpdateRestoreTaskOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableAutoSync

`func (o *SqlUpdateRestoreTaskOptions) GetEnableAutoSync() bool`

GetEnableAutoSync returns the EnableAutoSync field if non-nil, zero value otherwise.

### GetEnableAutoSyncOk

`func (o *SqlUpdateRestoreTaskOptions) GetEnableAutoSyncOk() (*bool, bool)`

GetEnableAutoSyncOk returns a tuple with the EnableAutoSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutoSync

`func (o *SqlUpdateRestoreTaskOptions) SetEnableAutoSync(v bool)`

SetEnableAutoSync sets EnableAutoSync field to given value.

### HasEnableAutoSync

`func (o *SqlUpdateRestoreTaskOptions) HasEnableAutoSync() bool`

HasEnableAutoSync returns a boolean if a field has been set.

### SetEnableAutoSyncNil

`func (o *SqlUpdateRestoreTaskOptions) SetEnableAutoSyncNil(b bool)`

 SetEnableAutoSyncNil sets the value for EnableAutoSync to be an explicit nil

### UnsetEnableAutoSync
`func (o *SqlUpdateRestoreTaskOptions) UnsetEnableAutoSync()`

UnsetEnableAutoSync ensures that no value is present for EnableAutoSync, not even an explicit nil
### GetMultiStageRestoreAction

`func (o *SqlUpdateRestoreTaskOptions) GetMultiStageRestoreAction() int32`

GetMultiStageRestoreAction returns the MultiStageRestoreAction field if non-nil, zero value otherwise.

### GetMultiStageRestoreActionOk

`func (o *SqlUpdateRestoreTaskOptions) GetMultiStageRestoreActionOk() (*int32, bool)`

GetMultiStageRestoreActionOk returns a tuple with the MultiStageRestoreAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiStageRestoreAction

`func (o *SqlUpdateRestoreTaskOptions) SetMultiStageRestoreAction(v int32)`

SetMultiStageRestoreAction sets MultiStageRestoreAction field to given value.

### HasMultiStageRestoreAction

`func (o *SqlUpdateRestoreTaskOptions) HasMultiStageRestoreAction() bool`

HasMultiStageRestoreAction returns a boolean if a field has been set.

### SetMultiStageRestoreActionNil

`func (o *SqlUpdateRestoreTaskOptions) SetMultiStageRestoreActionNil(b bool)`

 SetMultiStageRestoreActionNil sets the value for MultiStageRestoreAction to be an explicit nil

### UnsetMultiStageRestoreAction
`func (o *SqlUpdateRestoreTaskOptions) UnsetMultiStageRestoreAction()`

UnsetMultiStageRestoreAction ensures that no value is present for MultiStageRestoreAction, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


