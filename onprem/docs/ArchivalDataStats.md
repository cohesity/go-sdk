# ArchivalDataStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogicalSizeBytes** | Pointer to **NullableInt64** | Specifies the logicalSizeBytes. | [optional] 
**BytesRead** | Pointer to **NullableInt64** | Specifies total logical bytes read for creating the snapshot. | [optional] 
**LogicalBytesTransferred** | Pointer to **NullableInt64** | Specifies the logical bytes transferred. | [optional] 
**PhysicalBytesTransferred** | Pointer to **NullableInt64** | Specifies the physical bytes transferred. | [optional] 
**AvgLogicalTransferRateBps** | Pointer to **NullableInt64** | Specifies the average rate of transfer in bytes per second. | [optional] 
**FileWalkDone** | Pointer to **NullableBool** | Specifies whether the file system walk is done. Only applicable to file based backups. | [optional] 
**TotalFileCount** | Pointer to **NullableInt64** | Specifies the total number of file and directory entities visited in this backup. Only applicable to file based backups. | [optional] 
**BackupFileCount** | Pointer to **NullableInt64** | Specifies the total number of file and directory entities that are backed up in this run. Only applicable to file based backups. | [optional] 

## Methods

### NewArchivalDataStats

`func NewArchivalDataStats() *ArchivalDataStats`

NewArchivalDataStats instantiates a new ArchivalDataStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArchivalDataStatsWithDefaults

`func NewArchivalDataStatsWithDefaults() *ArchivalDataStats`

NewArchivalDataStatsWithDefaults instantiates a new ArchivalDataStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogicalSizeBytes

`func (o *ArchivalDataStats) GetLogicalSizeBytes() int64`

GetLogicalSizeBytes returns the LogicalSizeBytes field if non-nil, zero value otherwise.

### GetLogicalSizeBytesOk

`func (o *ArchivalDataStats) GetLogicalSizeBytesOk() (*int64, bool)`

GetLogicalSizeBytesOk returns a tuple with the LogicalSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalSizeBytes

`func (o *ArchivalDataStats) SetLogicalSizeBytes(v int64)`

SetLogicalSizeBytes sets LogicalSizeBytes field to given value.

### HasLogicalSizeBytes

`func (o *ArchivalDataStats) HasLogicalSizeBytes() bool`

HasLogicalSizeBytes returns a boolean if a field has been set.

### SetLogicalSizeBytesNil

`func (o *ArchivalDataStats) SetLogicalSizeBytesNil(b bool)`

 SetLogicalSizeBytesNil sets the value for LogicalSizeBytes to be an explicit nil

### UnsetLogicalSizeBytes
`func (o *ArchivalDataStats) UnsetLogicalSizeBytes()`

UnsetLogicalSizeBytes ensures that no value is present for LogicalSizeBytes, not even an explicit nil
### GetBytesRead

`func (o *ArchivalDataStats) GetBytesRead() int64`

GetBytesRead returns the BytesRead field if non-nil, zero value otherwise.

### GetBytesReadOk

`func (o *ArchivalDataStats) GetBytesReadOk() (*int64, bool)`

GetBytesReadOk returns a tuple with the BytesRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesRead

`func (o *ArchivalDataStats) SetBytesRead(v int64)`

SetBytesRead sets BytesRead field to given value.

### HasBytesRead

`func (o *ArchivalDataStats) HasBytesRead() bool`

HasBytesRead returns a boolean if a field has been set.

### SetBytesReadNil

`func (o *ArchivalDataStats) SetBytesReadNil(b bool)`

 SetBytesReadNil sets the value for BytesRead to be an explicit nil

### UnsetBytesRead
`func (o *ArchivalDataStats) UnsetBytesRead()`

UnsetBytesRead ensures that no value is present for BytesRead, not even an explicit nil
### GetLogicalBytesTransferred

`func (o *ArchivalDataStats) GetLogicalBytesTransferred() int64`

GetLogicalBytesTransferred returns the LogicalBytesTransferred field if non-nil, zero value otherwise.

### GetLogicalBytesTransferredOk

`func (o *ArchivalDataStats) GetLogicalBytesTransferredOk() (*int64, bool)`

GetLogicalBytesTransferredOk returns a tuple with the LogicalBytesTransferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalBytesTransferred

`func (o *ArchivalDataStats) SetLogicalBytesTransferred(v int64)`

SetLogicalBytesTransferred sets LogicalBytesTransferred field to given value.

### HasLogicalBytesTransferred

`func (o *ArchivalDataStats) HasLogicalBytesTransferred() bool`

HasLogicalBytesTransferred returns a boolean if a field has been set.

### SetLogicalBytesTransferredNil

`func (o *ArchivalDataStats) SetLogicalBytesTransferredNil(b bool)`

 SetLogicalBytesTransferredNil sets the value for LogicalBytesTransferred to be an explicit nil

### UnsetLogicalBytesTransferred
`func (o *ArchivalDataStats) UnsetLogicalBytesTransferred()`

UnsetLogicalBytesTransferred ensures that no value is present for LogicalBytesTransferred, not even an explicit nil
### GetPhysicalBytesTransferred

`func (o *ArchivalDataStats) GetPhysicalBytesTransferred() int64`

GetPhysicalBytesTransferred returns the PhysicalBytesTransferred field if non-nil, zero value otherwise.

### GetPhysicalBytesTransferredOk

`func (o *ArchivalDataStats) GetPhysicalBytesTransferredOk() (*int64, bool)`

GetPhysicalBytesTransferredOk returns a tuple with the PhysicalBytesTransferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalBytesTransferred

`func (o *ArchivalDataStats) SetPhysicalBytesTransferred(v int64)`

SetPhysicalBytesTransferred sets PhysicalBytesTransferred field to given value.

### HasPhysicalBytesTransferred

`func (o *ArchivalDataStats) HasPhysicalBytesTransferred() bool`

HasPhysicalBytesTransferred returns a boolean if a field has been set.

### SetPhysicalBytesTransferredNil

`func (o *ArchivalDataStats) SetPhysicalBytesTransferredNil(b bool)`

 SetPhysicalBytesTransferredNil sets the value for PhysicalBytesTransferred to be an explicit nil

### UnsetPhysicalBytesTransferred
`func (o *ArchivalDataStats) UnsetPhysicalBytesTransferred()`

UnsetPhysicalBytesTransferred ensures that no value is present for PhysicalBytesTransferred, not even an explicit nil
### GetAvgLogicalTransferRateBps

`func (o *ArchivalDataStats) GetAvgLogicalTransferRateBps() int64`

GetAvgLogicalTransferRateBps returns the AvgLogicalTransferRateBps field if non-nil, zero value otherwise.

### GetAvgLogicalTransferRateBpsOk

`func (o *ArchivalDataStats) GetAvgLogicalTransferRateBpsOk() (*int64, bool)`

GetAvgLogicalTransferRateBpsOk returns a tuple with the AvgLogicalTransferRateBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgLogicalTransferRateBps

`func (o *ArchivalDataStats) SetAvgLogicalTransferRateBps(v int64)`

SetAvgLogicalTransferRateBps sets AvgLogicalTransferRateBps field to given value.

### HasAvgLogicalTransferRateBps

`func (o *ArchivalDataStats) HasAvgLogicalTransferRateBps() bool`

HasAvgLogicalTransferRateBps returns a boolean if a field has been set.

### SetAvgLogicalTransferRateBpsNil

`func (o *ArchivalDataStats) SetAvgLogicalTransferRateBpsNil(b bool)`

 SetAvgLogicalTransferRateBpsNil sets the value for AvgLogicalTransferRateBps to be an explicit nil

### UnsetAvgLogicalTransferRateBps
`func (o *ArchivalDataStats) UnsetAvgLogicalTransferRateBps()`

UnsetAvgLogicalTransferRateBps ensures that no value is present for AvgLogicalTransferRateBps, not even an explicit nil
### GetFileWalkDone

`func (o *ArchivalDataStats) GetFileWalkDone() bool`

GetFileWalkDone returns the FileWalkDone field if non-nil, zero value otherwise.

### GetFileWalkDoneOk

`func (o *ArchivalDataStats) GetFileWalkDoneOk() (*bool, bool)`

GetFileWalkDoneOk returns a tuple with the FileWalkDone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileWalkDone

`func (o *ArchivalDataStats) SetFileWalkDone(v bool)`

SetFileWalkDone sets FileWalkDone field to given value.

### HasFileWalkDone

`func (o *ArchivalDataStats) HasFileWalkDone() bool`

HasFileWalkDone returns a boolean if a field has been set.

### SetFileWalkDoneNil

`func (o *ArchivalDataStats) SetFileWalkDoneNil(b bool)`

 SetFileWalkDoneNil sets the value for FileWalkDone to be an explicit nil

### UnsetFileWalkDone
`func (o *ArchivalDataStats) UnsetFileWalkDone()`

UnsetFileWalkDone ensures that no value is present for FileWalkDone, not even an explicit nil
### GetTotalFileCount

`func (o *ArchivalDataStats) GetTotalFileCount() int64`

GetTotalFileCount returns the TotalFileCount field if non-nil, zero value otherwise.

### GetTotalFileCountOk

`func (o *ArchivalDataStats) GetTotalFileCountOk() (*int64, bool)`

GetTotalFileCountOk returns a tuple with the TotalFileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFileCount

`func (o *ArchivalDataStats) SetTotalFileCount(v int64)`

SetTotalFileCount sets TotalFileCount field to given value.

### HasTotalFileCount

`func (o *ArchivalDataStats) HasTotalFileCount() bool`

HasTotalFileCount returns a boolean if a field has been set.

### SetTotalFileCountNil

`func (o *ArchivalDataStats) SetTotalFileCountNil(b bool)`

 SetTotalFileCountNil sets the value for TotalFileCount to be an explicit nil

### UnsetTotalFileCount
`func (o *ArchivalDataStats) UnsetTotalFileCount()`

UnsetTotalFileCount ensures that no value is present for TotalFileCount, not even an explicit nil
### GetBackupFileCount

`func (o *ArchivalDataStats) GetBackupFileCount() int64`

GetBackupFileCount returns the BackupFileCount field if non-nil, zero value otherwise.

### GetBackupFileCountOk

`func (o *ArchivalDataStats) GetBackupFileCountOk() (*int64, bool)`

GetBackupFileCountOk returns a tuple with the BackupFileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupFileCount

`func (o *ArchivalDataStats) SetBackupFileCount(v int64)`

SetBackupFileCount sets BackupFileCount field to given value.

### HasBackupFileCount

`func (o *ArchivalDataStats) HasBackupFileCount() bool`

HasBackupFileCount returns a boolean if a field has been set.

### SetBackupFileCountNil

`func (o *ArchivalDataStats) SetBackupFileCountNil(b bool)`

 SetBackupFileCountNil sets the value for BackupFileCount to be an explicit nil

### UnsetBackupFileCount
`func (o *ArchivalDataStats) UnsetBackupFileCount()`

UnsetBackupFileCount ensures that no value is present for BackupFileCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


