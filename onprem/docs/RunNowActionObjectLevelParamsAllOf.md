# RunNowActionObjectLevelParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TakeLocalSnapshotOnly** | Pointer to **NullableBool** | If sepcified as true then runNow will only take local snapshot ignoring all other targets such as replication, archivals etc. If not sepcified or specified as false then runNow will follow the policy settings. | [optional] 
**BackupType** | Pointer to **string** | Specifies the backup type should be used for RunNow action. | [optional] 

## Methods

### NewRunNowActionObjectLevelParamsAllOf

`func NewRunNowActionObjectLevelParamsAllOf() *RunNowActionObjectLevelParamsAllOf`

NewRunNowActionObjectLevelParamsAllOf instantiates a new RunNowActionObjectLevelParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunNowActionObjectLevelParamsAllOfWithDefaults

`func NewRunNowActionObjectLevelParamsAllOfWithDefaults() *RunNowActionObjectLevelParamsAllOf`

NewRunNowActionObjectLevelParamsAllOfWithDefaults instantiates a new RunNowActionObjectLevelParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTakeLocalSnapshotOnly

`func (o *RunNowActionObjectLevelParamsAllOf) GetTakeLocalSnapshotOnly() bool`

GetTakeLocalSnapshotOnly returns the TakeLocalSnapshotOnly field if non-nil, zero value otherwise.

### GetTakeLocalSnapshotOnlyOk

`func (o *RunNowActionObjectLevelParamsAllOf) GetTakeLocalSnapshotOnlyOk() (*bool, bool)`

GetTakeLocalSnapshotOnlyOk returns a tuple with the TakeLocalSnapshotOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakeLocalSnapshotOnly

`func (o *RunNowActionObjectLevelParamsAllOf) SetTakeLocalSnapshotOnly(v bool)`

SetTakeLocalSnapshotOnly sets TakeLocalSnapshotOnly field to given value.

### HasTakeLocalSnapshotOnly

`func (o *RunNowActionObjectLevelParamsAllOf) HasTakeLocalSnapshotOnly() bool`

HasTakeLocalSnapshotOnly returns a boolean if a field has been set.

### SetTakeLocalSnapshotOnlyNil

`func (o *RunNowActionObjectLevelParamsAllOf) SetTakeLocalSnapshotOnlyNil(b bool)`

 SetTakeLocalSnapshotOnlyNil sets the value for TakeLocalSnapshotOnly to be an explicit nil

### UnsetTakeLocalSnapshotOnly
`func (o *RunNowActionObjectLevelParamsAllOf) UnsetTakeLocalSnapshotOnly()`

UnsetTakeLocalSnapshotOnly ensures that no value is present for TakeLocalSnapshotOnly, not even an explicit nil
### GetBackupType

`func (o *RunNowActionObjectLevelParamsAllOf) GetBackupType() string`

GetBackupType returns the BackupType field if non-nil, zero value otherwise.

### GetBackupTypeOk

`func (o *RunNowActionObjectLevelParamsAllOf) GetBackupTypeOk() (*string, bool)`

GetBackupTypeOk returns a tuple with the BackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupType

`func (o *RunNowActionObjectLevelParamsAllOf) SetBackupType(v string)`

SetBackupType sets BackupType field to given value.

### HasBackupType

`func (o *RunNowActionObjectLevelParamsAllOf) HasBackupType() bool`

HasBackupType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


