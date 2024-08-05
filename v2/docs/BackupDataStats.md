# BackupDataStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BytesRead** | Pointer to **NullableInt64** | Specifies total logical bytes read for creating the snapshot. | [optional] 
**BytesWritten** | Pointer to **NullableInt64** | Specifies total size of data in bytes written after taking backup. | [optional] 
**LogicalSizeBytes** | Pointer to **NullableInt64** | Specifies total logical size of object(s) in bytes. | [optional] 

## Methods

### NewBackupDataStats

`func NewBackupDataStats() *BackupDataStats`

NewBackupDataStats instantiates a new BackupDataStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupDataStatsWithDefaults

`func NewBackupDataStatsWithDefaults() *BackupDataStats`

NewBackupDataStatsWithDefaults instantiates a new BackupDataStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBytesRead

`func (o *BackupDataStats) GetBytesRead() int64`

GetBytesRead returns the BytesRead field if non-nil, zero value otherwise.

### GetBytesReadOk

`func (o *BackupDataStats) GetBytesReadOk() (*int64, bool)`

GetBytesReadOk returns a tuple with the BytesRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesRead

`func (o *BackupDataStats) SetBytesRead(v int64)`

SetBytesRead sets BytesRead field to given value.

### HasBytesRead

`func (o *BackupDataStats) HasBytesRead() bool`

HasBytesRead returns a boolean if a field has been set.

### SetBytesReadNil

`func (o *BackupDataStats) SetBytesReadNil(b bool)`

 SetBytesReadNil sets the value for BytesRead to be an explicit nil

### UnsetBytesRead
`func (o *BackupDataStats) UnsetBytesRead()`

UnsetBytesRead ensures that no value is present for BytesRead, not even an explicit nil
### GetBytesWritten

`func (o *BackupDataStats) GetBytesWritten() int64`

GetBytesWritten returns the BytesWritten field if non-nil, zero value otherwise.

### GetBytesWrittenOk

`func (o *BackupDataStats) GetBytesWrittenOk() (*int64, bool)`

GetBytesWrittenOk returns a tuple with the BytesWritten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesWritten

`func (o *BackupDataStats) SetBytesWritten(v int64)`

SetBytesWritten sets BytesWritten field to given value.

### HasBytesWritten

`func (o *BackupDataStats) HasBytesWritten() bool`

HasBytesWritten returns a boolean if a field has been set.

### SetBytesWrittenNil

`func (o *BackupDataStats) SetBytesWrittenNil(b bool)`

 SetBytesWrittenNil sets the value for BytesWritten to be an explicit nil

### UnsetBytesWritten
`func (o *BackupDataStats) UnsetBytesWritten()`

UnsetBytesWritten ensures that no value is present for BytesWritten, not even an explicit nil
### GetLogicalSizeBytes

`func (o *BackupDataStats) GetLogicalSizeBytes() int64`

GetLogicalSizeBytes returns the LogicalSizeBytes field if non-nil, zero value otherwise.

### GetLogicalSizeBytesOk

`func (o *BackupDataStats) GetLogicalSizeBytesOk() (*int64, bool)`

GetLogicalSizeBytesOk returns a tuple with the LogicalSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalSizeBytes

`func (o *BackupDataStats) SetLogicalSizeBytes(v int64)`

SetLogicalSizeBytes sets LogicalSizeBytes field to given value.

### HasLogicalSizeBytes

`func (o *BackupDataStats) HasLogicalSizeBytes() bool`

HasLogicalSizeBytes returns a boolean if a field has been set.

### SetLogicalSizeBytesNil

`func (o *BackupDataStats) SetLogicalSizeBytesNil(b bool)`

 SetLogicalSizeBytesNil sets the value for LogicalSizeBytes to be an explicit nil

### UnsetLogicalSizeBytes
`func (o *BackupDataStats) UnsetLogicalSizeBytes()`

UnsetLogicalSizeBytes ensures that no value is present for LogicalSizeBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


