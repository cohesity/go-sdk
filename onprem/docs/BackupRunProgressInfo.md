# BackupRunProgressInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **NullableString** | Specifies the current status of the progress task. | [optional] 
**PercentageCompleted** | Pointer to **NullableFloat32** | Specifies the current completed percentage of the progress task. | [optional] 
**StartTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time of the progress task in Unix epoch Timestamp(in microseconds). | [optional] 
**EndTimeUsecs** | Pointer to **NullableInt64** | Specifies the end time of the progress task in Unix epoch Timestamp(in microseconds). | [optional] 
**ExpectedRemainingTimeUsecs** | Pointer to **NullableInt64** | Specifies the expected remaining time of the progress task in Unix epoch Timestamp(in microseconds). | [optional] 
**Events** | Pointer to [**[]ProgressTaskEvent**](ProgressTaskEvent.md) | Specifies the event log created for progress Task. | [optional] 
**Stats** | Pointer to [**ProgressStats**](ProgressStats.md) |  | [optional] 
**Objects** | Pointer to [**[]ObjectProgressInfo**](ObjectProgressInfo.md) | Specifies progress for objects. | [optional] 

## Methods

### NewBackupRunProgressInfo

`func NewBackupRunProgressInfo() *BackupRunProgressInfo`

NewBackupRunProgressInfo instantiates a new BackupRunProgressInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRunProgressInfoWithDefaults

`func NewBackupRunProgressInfoWithDefaults() *BackupRunProgressInfo`

NewBackupRunProgressInfoWithDefaults instantiates a new BackupRunProgressInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *BackupRunProgressInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BackupRunProgressInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BackupRunProgressInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BackupRunProgressInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *BackupRunProgressInfo) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *BackupRunProgressInfo) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetPercentageCompleted

`func (o *BackupRunProgressInfo) GetPercentageCompleted() float32`

GetPercentageCompleted returns the PercentageCompleted field if non-nil, zero value otherwise.

### GetPercentageCompletedOk

`func (o *BackupRunProgressInfo) GetPercentageCompletedOk() (*float32, bool)`

GetPercentageCompletedOk returns a tuple with the PercentageCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageCompleted

`func (o *BackupRunProgressInfo) SetPercentageCompleted(v float32)`

SetPercentageCompleted sets PercentageCompleted field to given value.

### HasPercentageCompleted

`func (o *BackupRunProgressInfo) HasPercentageCompleted() bool`

HasPercentageCompleted returns a boolean if a field has been set.

### SetPercentageCompletedNil

`func (o *BackupRunProgressInfo) SetPercentageCompletedNil(b bool)`

 SetPercentageCompletedNil sets the value for PercentageCompleted to be an explicit nil

### UnsetPercentageCompleted
`func (o *BackupRunProgressInfo) UnsetPercentageCompleted()`

UnsetPercentageCompleted ensures that no value is present for PercentageCompleted, not even an explicit nil
### GetStartTimeUsecs

`func (o *BackupRunProgressInfo) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *BackupRunProgressInfo) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *BackupRunProgressInfo) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *BackupRunProgressInfo) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *BackupRunProgressInfo) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *BackupRunProgressInfo) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetEndTimeUsecs

`func (o *BackupRunProgressInfo) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *BackupRunProgressInfo) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *BackupRunProgressInfo) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *BackupRunProgressInfo) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *BackupRunProgressInfo) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *BackupRunProgressInfo) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetExpectedRemainingTimeUsecs

`func (o *BackupRunProgressInfo) GetExpectedRemainingTimeUsecs() int64`

GetExpectedRemainingTimeUsecs returns the ExpectedRemainingTimeUsecs field if non-nil, zero value otherwise.

### GetExpectedRemainingTimeUsecsOk

`func (o *BackupRunProgressInfo) GetExpectedRemainingTimeUsecsOk() (*int64, bool)`

GetExpectedRemainingTimeUsecsOk returns a tuple with the ExpectedRemainingTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedRemainingTimeUsecs

`func (o *BackupRunProgressInfo) SetExpectedRemainingTimeUsecs(v int64)`

SetExpectedRemainingTimeUsecs sets ExpectedRemainingTimeUsecs field to given value.

### HasExpectedRemainingTimeUsecs

`func (o *BackupRunProgressInfo) HasExpectedRemainingTimeUsecs() bool`

HasExpectedRemainingTimeUsecs returns a boolean if a field has been set.

### SetExpectedRemainingTimeUsecsNil

`func (o *BackupRunProgressInfo) SetExpectedRemainingTimeUsecsNil(b bool)`

 SetExpectedRemainingTimeUsecsNil sets the value for ExpectedRemainingTimeUsecs to be an explicit nil

### UnsetExpectedRemainingTimeUsecs
`func (o *BackupRunProgressInfo) UnsetExpectedRemainingTimeUsecs()`

UnsetExpectedRemainingTimeUsecs ensures that no value is present for ExpectedRemainingTimeUsecs, not even an explicit nil
### GetEvents

`func (o *BackupRunProgressInfo) GetEvents() []ProgressTaskEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *BackupRunProgressInfo) GetEventsOk() (*[]ProgressTaskEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *BackupRunProgressInfo) SetEvents(v []ProgressTaskEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *BackupRunProgressInfo) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetStats

`func (o *BackupRunProgressInfo) GetStats() ProgressStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *BackupRunProgressInfo) GetStatsOk() (*ProgressStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *BackupRunProgressInfo) SetStats(v ProgressStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *BackupRunProgressInfo) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetObjects

`func (o *BackupRunProgressInfo) GetObjects() []ObjectProgressInfo`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *BackupRunProgressInfo) GetObjectsOk() (*[]ObjectProgressInfo, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *BackupRunProgressInfo) SetObjects(v []ObjectProgressInfo)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *BackupRunProgressInfo) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### SetObjectsNil

`func (o *BackupRunProgressInfo) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *BackupRunProgressInfo) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


