# RunNowActionObjectLevelParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableInt64** | Specifies the ID of the object. | 
**Name** | Pointer to **NullableString** | Specifies the name of the object. | [optional] [readonly] 
**TakeLocalSnapshotOnly** | Pointer to **NullableBool** | If sepcified as true then runNow will only take local snapshot ignoring all other targets such as replication, archivals etc. If not sepcified or specified as false then runNow will follow the policy settings. | [optional] 
**BackupType** | Pointer to **string** | Specifies the backup type should be used for RunNow action. | [optional] 

## Methods

### NewRunNowActionObjectLevelParams

`func NewRunNowActionObjectLevelParams(id NullableInt64, ) *RunNowActionObjectLevelParams`

NewRunNowActionObjectLevelParams instantiates a new RunNowActionObjectLevelParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunNowActionObjectLevelParamsWithDefaults

`func NewRunNowActionObjectLevelParamsWithDefaults() *RunNowActionObjectLevelParams`

NewRunNowActionObjectLevelParamsWithDefaults instantiates a new RunNowActionObjectLevelParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RunNowActionObjectLevelParams) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RunNowActionObjectLevelParams) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RunNowActionObjectLevelParams) SetId(v int64)`

SetId sets Id field to given value.


### SetIdNil

`func (o *RunNowActionObjectLevelParams) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *RunNowActionObjectLevelParams) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *RunNowActionObjectLevelParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RunNowActionObjectLevelParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RunNowActionObjectLevelParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RunNowActionObjectLevelParams) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *RunNowActionObjectLevelParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *RunNowActionObjectLevelParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTakeLocalSnapshotOnly

`func (o *RunNowActionObjectLevelParams) GetTakeLocalSnapshotOnly() bool`

GetTakeLocalSnapshotOnly returns the TakeLocalSnapshotOnly field if non-nil, zero value otherwise.

### GetTakeLocalSnapshotOnlyOk

`func (o *RunNowActionObjectLevelParams) GetTakeLocalSnapshotOnlyOk() (*bool, bool)`

GetTakeLocalSnapshotOnlyOk returns a tuple with the TakeLocalSnapshotOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakeLocalSnapshotOnly

`func (o *RunNowActionObjectLevelParams) SetTakeLocalSnapshotOnly(v bool)`

SetTakeLocalSnapshotOnly sets TakeLocalSnapshotOnly field to given value.

### HasTakeLocalSnapshotOnly

`func (o *RunNowActionObjectLevelParams) HasTakeLocalSnapshotOnly() bool`

HasTakeLocalSnapshotOnly returns a boolean if a field has been set.

### SetTakeLocalSnapshotOnlyNil

`func (o *RunNowActionObjectLevelParams) SetTakeLocalSnapshotOnlyNil(b bool)`

 SetTakeLocalSnapshotOnlyNil sets the value for TakeLocalSnapshotOnly to be an explicit nil

### UnsetTakeLocalSnapshotOnly
`func (o *RunNowActionObjectLevelParams) UnsetTakeLocalSnapshotOnly()`

UnsetTakeLocalSnapshotOnly ensures that no value is present for TakeLocalSnapshotOnly, not even an explicit nil
### GetBackupType

`func (o *RunNowActionObjectLevelParams) GetBackupType() string`

GetBackupType returns the BackupType field if non-nil, zero value otherwise.

### GetBackupTypeOk

`func (o *RunNowActionObjectLevelParams) GetBackupTypeOk() (*string, bool)`

GetBackupTypeOk returns a tuple with the BackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupType

`func (o *RunNowActionObjectLevelParams) SetBackupType(v string)`

SetBackupType sets BackupType field to given value.

### HasBackupType

`func (o *RunNowActionObjectLevelParams) HasBackupType() bool`

HasBackupType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


