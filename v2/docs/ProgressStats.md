# ProgressStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupFileCount** | Pointer to **NullableInt64** | Specifies the total number of file and directory entities that are backed up in this run. Only applicable to file based backups. | [optional] 
**FileWalkDone** | Pointer to **NullableBool** | Specifies whether the file system walk is done. Only applicable to file based backups. | [optional] 
**TotalFileCount** | Pointer to **NullableInt64** | Specifies the total number of file and directory entities visited in this backup. Only applicable to file based backups. | [optional] 

## Methods

### NewProgressStats

`func NewProgressStats() *ProgressStats`

NewProgressStats instantiates a new ProgressStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProgressStatsWithDefaults

`func NewProgressStatsWithDefaults() *ProgressStats`

NewProgressStatsWithDefaults instantiates a new ProgressStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupFileCount

`func (o *ProgressStats) GetBackupFileCount() int64`

GetBackupFileCount returns the BackupFileCount field if non-nil, zero value otherwise.

### GetBackupFileCountOk

`func (o *ProgressStats) GetBackupFileCountOk() (*int64, bool)`

GetBackupFileCountOk returns a tuple with the BackupFileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupFileCount

`func (o *ProgressStats) SetBackupFileCount(v int64)`

SetBackupFileCount sets BackupFileCount field to given value.

### HasBackupFileCount

`func (o *ProgressStats) HasBackupFileCount() bool`

HasBackupFileCount returns a boolean if a field has been set.

### SetBackupFileCountNil

`func (o *ProgressStats) SetBackupFileCountNil(b bool)`

 SetBackupFileCountNil sets the value for BackupFileCount to be an explicit nil

### UnsetBackupFileCount
`func (o *ProgressStats) UnsetBackupFileCount()`

UnsetBackupFileCount ensures that no value is present for BackupFileCount, not even an explicit nil
### GetFileWalkDone

`func (o *ProgressStats) GetFileWalkDone() bool`

GetFileWalkDone returns the FileWalkDone field if non-nil, zero value otherwise.

### GetFileWalkDoneOk

`func (o *ProgressStats) GetFileWalkDoneOk() (*bool, bool)`

GetFileWalkDoneOk returns a tuple with the FileWalkDone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileWalkDone

`func (o *ProgressStats) SetFileWalkDone(v bool)`

SetFileWalkDone sets FileWalkDone field to given value.

### HasFileWalkDone

`func (o *ProgressStats) HasFileWalkDone() bool`

HasFileWalkDone returns a boolean if a field has been set.

### SetFileWalkDoneNil

`func (o *ProgressStats) SetFileWalkDoneNil(b bool)`

 SetFileWalkDoneNil sets the value for FileWalkDone to be an explicit nil

### UnsetFileWalkDone
`func (o *ProgressStats) UnsetFileWalkDone()`

UnsetFileWalkDone ensures that no value is present for FileWalkDone, not even an explicit nil
### GetTotalFileCount

`func (o *ProgressStats) GetTotalFileCount() int64`

GetTotalFileCount returns the TotalFileCount field if non-nil, zero value otherwise.

### GetTotalFileCountOk

`func (o *ProgressStats) GetTotalFileCountOk() (*int64, bool)`

GetTotalFileCountOk returns a tuple with the TotalFileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFileCount

`func (o *ProgressStats) SetTotalFileCount(v int64)`

SetTotalFileCount sets TotalFileCount field to given value.

### HasTotalFileCount

`func (o *ProgressStats) HasTotalFileCount() bool`

HasTotalFileCount returns a boolean if a field has been set.

### SetTotalFileCountNil

`func (o *ProgressStats) SetTotalFileCountNil(b bool)`

 SetTotalFileCountNil sets the value for TotalFileCount to be an explicit nil

### UnsetTotalFileCount
`func (o *ProgressStats) UnsetTotalFileCount()`

UnsetTotalFileCount ensures that no value is present for TotalFileCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


