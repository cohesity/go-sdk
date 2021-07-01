# FileVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModifiedTimeUsecs** | Pointer to **NullableInt64** | Specifies the time when the file or folder was last modified. Specified as a Unix epoch Timestamp (in microseconds). | [optional] 
**SizeBytes** | Pointer to **NullableInt64** | Specifies the size of the file or folder (in bytes) from the most recent snapshot. | [optional] 
**Snapshots** | Pointer to [**[]SnapshotAttempt**](SnapshotAttempt.md) | Array of Snapshots.  Specifies the available snapshots of the file or folder. When a Job Run executes, it captures a snapshot of object (such as a VM) that contains the file or folder. | [optional] 

## Methods

### NewFileVersion

`func NewFileVersion() *FileVersion`

NewFileVersion instantiates a new FileVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileVersionWithDefaults

`func NewFileVersionWithDefaults() *FileVersion`

NewFileVersionWithDefaults instantiates a new FileVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModifiedTimeUsecs

`func (o *FileVersion) GetModifiedTimeUsecs() int64`

GetModifiedTimeUsecs returns the ModifiedTimeUsecs field if non-nil, zero value otherwise.

### GetModifiedTimeUsecsOk

`func (o *FileVersion) GetModifiedTimeUsecsOk() (*int64, bool)`

GetModifiedTimeUsecsOk returns a tuple with the ModifiedTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTimeUsecs

`func (o *FileVersion) SetModifiedTimeUsecs(v int64)`

SetModifiedTimeUsecs sets ModifiedTimeUsecs field to given value.

### HasModifiedTimeUsecs

`func (o *FileVersion) HasModifiedTimeUsecs() bool`

HasModifiedTimeUsecs returns a boolean if a field has been set.

### SetModifiedTimeUsecsNil

`func (o *FileVersion) SetModifiedTimeUsecsNil(b bool)`

 SetModifiedTimeUsecsNil sets the value for ModifiedTimeUsecs to be an explicit nil

### UnsetModifiedTimeUsecs
`func (o *FileVersion) UnsetModifiedTimeUsecs()`

UnsetModifiedTimeUsecs ensures that no value is present for ModifiedTimeUsecs, not even an explicit nil
### GetSizeBytes

`func (o *FileVersion) GetSizeBytes() int64`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *FileVersion) GetSizeBytesOk() (*int64, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *FileVersion) SetSizeBytes(v int64)`

SetSizeBytes sets SizeBytes field to given value.

### HasSizeBytes

`func (o *FileVersion) HasSizeBytes() bool`

HasSizeBytes returns a boolean if a field has been set.

### SetSizeBytesNil

`func (o *FileVersion) SetSizeBytesNil(b bool)`

 SetSizeBytesNil sets the value for SizeBytes to be an explicit nil

### UnsetSizeBytes
`func (o *FileVersion) UnsetSizeBytes()`

UnsetSizeBytes ensures that no value is present for SizeBytes, not even an explicit nil
### GetSnapshots

`func (o *FileVersion) GetSnapshots() []SnapshotAttempt`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *FileVersion) GetSnapshotsOk() (*[]SnapshotAttempt, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *FileVersion) SetSnapshots(v []SnapshotAttempt)`

SetSnapshots sets Snapshots field to given value.

### HasSnapshots

`func (o *FileVersion) HasSnapshots() bool`

HasSnapshots returns a boolean if a field has been set.

### SetSnapshotsNil

`func (o *FileVersion) SetSnapshotsNil(b bool)`

 SetSnapshotsNil sets the value for Snapshots to be an explicit nil

### UnsetSnapshots
`func (o *FileVersion) UnsetSnapshots()`

UnsetSnapshots ensures that no value is present for Snapshots, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


