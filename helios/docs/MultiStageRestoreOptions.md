# MultiStageRestoreOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableMultiStageRestore** | Pointer to **NullableBool** | Set this to true if you are creating a multi-stage Sql restore task needed for features such as Hot-Standby. | [optional] 
**EnableAutoSync** | Pointer to **NullableBool** | Set this to true if you want to enable auto sync for multi stage restore. | [optional] 

## Methods

### NewMultiStageRestoreOptions

`func NewMultiStageRestoreOptions() *MultiStageRestoreOptions`

NewMultiStageRestoreOptions instantiates a new MultiStageRestoreOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiStageRestoreOptionsWithDefaults

`func NewMultiStageRestoreOptionsWithDefaults() *MultiStageRestoreOptions`

NewMultiStageRestoreOptionsWithDefaults instantiates a new MultiStageRestoreOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableMultiStageRestore

`func (o *MultiStageRestoreOptions) GetEnableMultiStageRestore() bool`

GetEnableMultiStageRestore returns the EnableMultiStageRestore field if non-nil, zero value otherwise.

### GetEnableMultiStageRestoreOk

`func (o *MultiStageRestoreOptions) GetEnableMultiStageRestoreOk() (*bool, bool)`

GetEnableMultiStageRestoreOk returns a tuple with the EnableMultiStageRestore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMultiStageRestore

`func (o *MultiStageRestoreOptions) SetEnableMultiStageRestore(v bool)`

SetEnableMultiStageRestore sets EnableMultiStageRestore field to given value.

### HasEnableMultiStageRestore

`func (o *MultiStageRestoreOptions) HasEnableMultiStageRestore() bool`

HasEnableMultiStageRestore returns a boolean if a field has been set.

### SetEnableMultiStageRestoreNil

`func (o *MultiStageRestoreOptions) SetEnableMultiStageRestoreNil(b bool)`

 SetEnableMultiStageRestoreNil sets the value for EnableMultiStageRestore to be an explicit nil

### UnsetEnableMultiStageRestore
`func (o *MultiStageRestoreOptions) UnsetEnableMultiStageRestore()`

UnsetEnableMultiStageRestore ensures that no value is present for EnableMultiStageRestore, not even an explicit nil
### GetEnableAutoSync

`func (o *MultiStageRestoreOptions) GetEnableAutoSync() bool`

GetEnableAutoSync returns the EnableAutoSync field if non-nil, zero value otherwise.

### GetEnableAutoSyncOk

`func (o *MultiStageRestoreOptions) GetEnableAutoSyncOk() (*bool, bool)`

GetEnableAutoSyncOk returns a tuple with the EnableAutoSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutoSync

`func (o *MultiStageRestoreOptions) SetEnableAutoSync(v bool)`

SetEnableAutoSync sets EnableAutoSync field to given value.

### HasEnableAutoSync

`func (o *MultiStageRestoreOptions) HasEnableAutoSync() bool`

HasEnableAutoSync returns a boolean if a field has been set.

### SetEnableAutoSyncNil

`func (o *MultiStageRestoreOptions) SetEnableAutoSyncNil(b bool)`

 SetEnableAutoSyncNil sets the value for EnableAutoSync to be an explicit nil

### UnsetEnableAutoSync
`func (o *MultiStageRestoreOptions) UnsetEnableAutoSync()`

UnsetEnableAutoSync ensures that no value is present for EnableAutoSync, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


