# RestoreFileCopyStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EstimationSkipped** | Pointer to **NullableBool** | This will be set to true if the estimation step was skipped. NOTE: If estimation is skipped, then progress info will not be available. | [optional] 
**NumBytesCopied** | Pointer to **NullableInt64** | Number of bytes copied so far. | [optional] 
**NumDirectoriesCopied** | Pointer to **NullableInt32** | Number of directories copied so far. NOTE: This just means the creation of directory (not the contents of the directory). | [optional] 
**NumFilesCopied** | Pointer to **NullableInt32** | Number of files copied so far. | [optional] 
**TotalBytesToCopy** | Pointer to **NullableInt64** | Total number of bytes to copy. | [optional] 
**TotalDirectoriesToCopy** | Pointer to **NullableInt32** | Total number of directories to copy. NOTE: This just means the creation of directory (not the contents of the directory). | [optional] 
**TotalFilesToCopy** | Pointer to **NullableInt32** | Total number of files to copy. | [optional] 

## Methods

### NewRestoreFileCopyStats

`func NewRestoreFileCopyStats() *RestoreFileCopyStats`

NewRestoreFileCopyStats instantiates a new RestoreFileCopyStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreFileCopyStatsWithDefaults

`func NewRestoreFileCopyStatsWithDefaults() *RestoreFileCopyStats`

NewRestoreFileCopyStatsWithDefaults instantiates a new RestoreFileCopyStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEstimationSkipped

`func (o *RestoreFileCopyStats) GetEstimationSkipped() bool`

GetEstimationSkipped returns the EstimationSkipped field if non-nil, zero value otherwise.

### GetEstimationSkippedOk

`func (o *RestoreFileCopyStats) GetEstimationSkippedOk() (*bool, bool)`

GetEstimationSkippedOk returns a tuple with the EstimationSkipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimationSkipped

`func (o *RestoreFileCopyStats) SetEstimationSkipped(v bool)`

SetEstimationSkipped sets EstimationSkipped field to given value.

### HasEstimationSkipped

`func (o *RestoreFileCopyStats) HasEstimationSkipped() bool`

HasEstimationSkipped returns a boolean if a field has been set.

### SetEstimationSkippedNil

`func (o *RestoreFileCopyStats) SetEstimationSkippedNil(b bool)`

 SetEstimationSkippedNil sets the value for EstimationSkipped to be an explicit nil

### UnsetEstimationSkipped
`func (o *RestoreFileCopyStats) UnsetEstimationSkipped()`

UnsetEstimationSkipped ensures that no value is present for EstimationSkipped, not even an explicit nil
### GetNumBytesCopied

`func (o *RestoreFileCopyStats) GetNumBytesCopied() int64`

GetNumBytesCopied returns the NumBytesCopied field if non-nil, zero value otherwise.

### GetNumBytesCopiedOk

`func (o *RestoreFileCopyStats) GetNumBytesCopiedOk() (*int64, bool)`

GetNumBytesCopiedOk returns a tuple with the NumBytesCopied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumBytesCopied

`func (o *RestoreFileCopyStats) SetNumBytesCopied(v int64)`

SetNumBytesCopied sets NumBytesCopied field to given value.

### HasNumBytesCopied

`func (o *RestoreFileCopyStats) HasNumBytesCopied() bool`

HasNumBytesCopied returns a boolean if a field has been set.

### SetNumBytesCopiedNil

`func (o *RestoreFileCopyStats) SetNumBytesCopiedNil(b bool)`

 SetNumBytesCopiedNil sets the value for NumBytesCopied to be an explicit nil

### UnsetNumBytesCopied
`func (o *RestoreFileCopyStats) UnsetNumBytesCopied()`

UnsetNumBytesCopied ensures that no value is present for NumBytesCopied, not even an explicit nil
### GetNumDirectoriesCopied

`func (o *RestoreFileCopyStats) GetNumDirectoriesCopied() int32`

GetNumDirectoriesCopied returns the NumDirectoriesCopied field if non-nil, zero value otherwise.

### GetNumDirectoriesCopiedOk

`func (o *RestoreFileCopyStats) GetNumDirectoriesCopiedOk() (*int32, bool)`

GetNumDirectoriesCopiedOk returns a tuple with the NumDirectoriesCopied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDirectoriesCopied

`func (o *RestoreFileCopyStats) SetNumDirectoriesCopied(v int32)`

SetNumDirectoriesCopied sets NumDirectoriesCopied field to given value.

### HasNumDirectoriesCopied

`func (o *RestoreFileCopyStats) HasNumDirectoriesCopied() bool`

HasNumDirectoriesCopied returns a boolean if a field has been set.

### SetNumDirectoriesCopiedNil

`func (o *RestoreFileCopyStats) SetNumDirectoriesCopiedNil(b bool)`

 SetNumDirectoriesCopiedNil sets the value for NumDirectoriesCopied to be an explicit nil

### UnsetNumDirectoriesCopied
`func (o *RestoreFileCopyStats) UnsetNumDirectoriesCopied()`

UnsetNumDirectoriesCopied ensures that no value is present for NumDirectoriesCopied, not even an explicit nil
### GetNumFilesCopied

`func (o *RestoreFileCopyStats) GetNumFilesCopied() int32`

GetNumFilesCopied returns the NumFilesCopied field if non-nil, zero value otherwise.

### GetNumFilesCopiedOk

`func (o *RestoreFileCopyStats) GetNumFilesCopiedOk() (*int32, bool)`

GetNumFilesCopiedOk returns a tuple with the NumFilesCopied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFilesCopied

`func (o *RestoreFileCopyStats) SetNumFilesCopied(v int32)`

SetNumFilesCopied sets NumFilesCopied field to given value.

### HasNumFilesCopied

`func (o *RestoreFileCopyStats) HasNumFilesCopied() bool`

HasNumFilesCopied returns a boolean if a field has been set.

### SetNumFilesCopiedNil

`func (o *RestoreFileCopyStats) SetNumFilesCopiedNil(b bool)`

 SetNumFilesCopiedNil sets the value for NumFilesCopied to be an explicit nil

### UnsetNumFilesCopied
`func (o *RestoreFileCopyStats) UnsetNumFilesCopied()`

UnsetNumFilesCopied ensures that no value is present for NumFilesCopied, not even an explicit nil
### GetTotalBytesToCopy

`func (o *RestoreFileCopyStats) GetTotalBytesToCopy() int64`

GetTotalBytesToCopy returns the TotalBytesToCopy field if non-nil, zero value otherwise.

### GetTotalBytesToCopyOk

`func (o *RestoreFileCopyStats) GetTotalBytesToCopyOk() (*int64, bool)`

GetTotalBytesToCopyOk returns a tuple with the TotalBytesToCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBytesToCopy

`func (o *RestoreFileCopyStats) SetTotalBytesToCopy(v int64)`

SetTotalBytesToCopy sets TotalBytesToCopy field to given value.

### HasTotalBytesToCopy

`func (o *RestoreFileCopyStats) HasTotalBytesToCopy() bool`

HasTotalBytesToCopy returns a boolean if a field has been set.

### SetTotalBytesToCopyNil

`func (o *RestoreFileCopyStats) SetTotalBytesToCopyNil(b bool)`

 SetTotalBytesToCopyNil sets the value for TotalBytesToCopy to be an explicit nil

### UnsetTotalBytesToCopy
`func (o *RestoreFileCopyStats) UnsetTotalBytesToCopy()`

UnsetTotalBytesToCopy ensures that no value is present for TotalBytesToCopy, not even an explicit nil
### GetTotalDirectoriesToCopy

`func (o *RestoreFileCopyStats) GetTotalDirectoriesToCopy() int32`

GetTotalDirectoriesToCopy returns the TotalDirectoriesToCopy field if non-nil, zero value otherwise.

### GetTotalDirectoriesToCopyOk

`func (o *RestoreFileCopyStats) GetTotalDirectoriesToCopyOk() (*int32, bool)`

GetTotalDirectoriesToCopyOk returns a tuple with the TotalDirectoriesToCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDirectoriesToCopy

`func (o *RestoreFileCopyStats) SetTotalDirectoriesToCopy(v int32)`

SetTotalDirectoriesToCopy sets TotalDirectoriesToCopy field to given value.

### HasTotalDirectoriesToCopy

`func (o *RestoreFileCopyStats) HasTotalDirectoriesToCopy() bool`

HasTotalDirectoriesToCopy returns a boolean if a field has been set.

### SetTotalDirectoriesToCopyNil

`func (o *RestoreFileCopyStats) SetTotalDirectoriesToCopyNil(b bool)`

 SetTotalDirectoriesToCopyNil sets the value for TotalDirectoriesToCopy to be an explicit nil

### UnsetTotalDirectoriesToCopy
`func (o *RestoreFileCopyStats) UnsetTotalDirectoriesToCopy()`

UnsetTotalDirectoriesToCopy ensures that no value is present for TotalDirectoriesToCopy, not even an explicit nil
### GetTotalFilesToCopy

`func (o *RestoreFileCopyStats) GetTotalFilesToCopy() int32`

GetTotalFilesToCopy returns the TotalFilesToCopy field if non-nil, zero value otherwise.

### GetTotalFilesToCopyOk

`func (o *RestoreFileCopyStats) GetTotalFilesToCopyOk() (*int32, bool)`

GetTotalFilesToCopyOk returns a tuple with the TotalFilesToCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFilesToCopy

`func (o *RestoreFileCopyStats) SetTotalFilesToCopy(v int32)`

SetTotalFilesToCopy sets TotalFilesToCopy field to given value.

### HasTotalFilesToCopy

`func (o *RestoreFileCopyStats) HasTotalFilesToCopy() bool`

HasTotalFilesToCopy returns a boolean if a field has been set.

### SetTotalFilesToCopyNil

`func (o *RestoreFileCopyStats) SetTotalFilesToCopyNil(b bool)`

 SetTotalFilesToCopyNil sets the value for TotalFilesToCopy to be an explicit nil

### UnsetTotalFilesToCopy
`func (o *RestoreFileCopyStats) UnsetTotalFilesToCopy()`

UnsetTotalFilesToCopy ensures that no value is present for TotalFilesToCopy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


