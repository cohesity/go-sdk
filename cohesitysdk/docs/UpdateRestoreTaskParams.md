# UpdateRestoreTaskParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdOptions** | Pointer to [**AdRestoreOptions**](AdRestoreOptions.md) |  | [optional] 
**EnableAutoSync** | Pointer to **NullableBool** | Enables Auto Sync feature for SQL Multi-stage Restore task. | [optional] 
**RestoreTaskId** | Pointer to **NullableInt64** | Specifies the ID of the existing Restore Task to update. | [optional] 
**SqlOptions** | Pointer to **NullableString** | Specifies the sql options to update the Restore Task with. Specifies the action type of multi stage SQL restore.  &#39;kCreate&#39; specifies the create action for a restore. &#39;kUpdate&#39; specifies the user action to update an ongoing restore. &#39;kFinalize&#39; specifies the user action to finalize a restore. | [optional] 

## Methods

### NewUpdateRestoreTaskParams

`func NewUpdateRestoreTaskParams() *UpdateRestoreTaskParams`

NewUpdateRestoreTaskParams instantiates a new UpdateRestoreTaskParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRestoreTaskParamsWithDefaults

`func NewUpdateRestoreTaskParamsWithDefaults() *UpdateRestoreTaskParams`

NewUpdateRestoreTaskParamsWithDefaults instantiates a new UpdateRestoreTaskParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdOptions

`func (o *UpdateRestoreTaskParams) GetAdOptions() AdRestoreOptions`

GetAdOptions returns the AdOptions field if non-nil, zero value otherwise.

### GetAdOptionsOk

`func (o *UpdateRestoreTaskParams) GetAdOptionsOk() (*AdRestoreOptions, bool)`

GetAdOptionsOk returns a tuple with the AdOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdOptions

`func (o *UpdateRestoreTaskParams) SetAdOptions(v AdRestoreOptions)`

SetAdOptions sets AdOptions field to given value.

### HasAdOptions

`func (o *UpdateRestoreTaskParams) HasAdOptions() bool`

HasAdOptions returns a boolean if a field has been set.

### GetEnableAutoSync

`func (o *UpdateRestoreTaskParams) GetEnableAutoSync() bool`

GetEnableAutoSync returns the EnableAutoSync field if non-nil, zero value otherwise.

### GetEnableAutoSyncOk

`func (o *UpdateRestoreTaskParams) GetEnableAutoSyncOk() (*bool, bool)`

GetEnableAutoSyncOk returns a tuple with the EnableAutoSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutoSync

`func (o *UpdateRestoreTaskParams) SetEnableAutoSync(v bool)`

SetEnableAutoSync sets EnableAutoSync field to given value.

### HasEnableAutoSync

`func (o *UpdateRestoreTaskParams) HasEnableAutoSync() bool`

HasEnableAutoSync returns a boolean if a field has been set.

### SetEnableAutoSyncNil

`func (o *UpdateRestoreTaskParams) SetEnableAutoSyncNil(b bool)`

 SetEnableAutoSyncNil sets the value for EnableAutoSync to be an explicit nil

### UnsetEnableAutoSync
`func (o *UpdateRestoreTaskParams) UnsetEnableAutoSync()`

UnsetEnableAutoSync ensures that no value is present for EnableAutoSync, not even an explicit nil
### GetRestoreTaskId

`func (o *UpdateRestoreTaskParams) GetRestoreTaskId() int64`

GetRestoreTaskId returns the RestoreTaskId field if non-nil, zero value otherwise.

### GetRestoreTaskIdOk

`func (o *UpdateRestoreTaskParams) GetRestoreTaskIdOk() (*int64, bool)`

GetRestoreTaskIdOk returns a tuple with the RestoreTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreTaskId

`func (o *UpdateRestoreTaskParams) SetRestoreTaskId(v int64)`

SetRestoreTaskId sets RestoreTaskId field to given value.

### HasRestoreTaskId

`func (o *UpdateRestoreTaskParams) HasRestoreTaskId() bool`

HasRestoreTaskId returns a boolean if a field has been set.

### SetRestoreTaskIdNil

`func (o *UpdateRestoreTaskParams) SetRestoreTaskIdNil(b bool)`

 SetRestoreTaskIdNil sets the value for RestoreTaskId to be an explicit nil

### UnsetRestoreTaskId
`func (o *UpdateRestoreTaskParams) UnsetRestoreTaskId()`

UnsetRestoreTaskId ensures that no value is present for RestoreTaskId, not even an explicit nil
### GetSqlOptions

`func (o *UpdateRestoreTaskParams) GetSqlOptions() string`

GetSqlOptions returns the SqlOptions field if non-nil, zero value otherwise.

### GetSqlOptionsOk

`func (o *UpdateRestoreTaskParams) GetSqlOptionsOk() (*string, bool)`

GetSqlOptionsOk returns a tuple with the SqlOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlOptions

`func (o *UpdateRestoreTaskParams) SetSqlOptions(v string)`

SetSqlOptions sets SqlOptions field to given value.

### HasSqlOptions

`func (o *UpdateRestoreTaskParams) HasSqlOptions() bool`

HasSqlOptions returns a boolean if a field has been set.

### SetSqlOptionsNil

`func (o *UpdateRestoreTaskParams) SetSqlOptionsNil(b bool)`

 SetSqlOptionsNil sets the value for SqlOptions to be an explicit nil

### UnsetSqlOptions
`func (o *UpdateRestoreTaskParams) UnsetSqlOptions()`

UnsetSqlOptions ensures that no value is present for SqlOptions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


