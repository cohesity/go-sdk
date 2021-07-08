# InputSpecInputFilesSelector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileTimeFilter** | Pointer to [**InputSpecFileTimeFilter**](InputSpecFileTimeFilter.md) |  | [optional] 
**FilenameGlob** | Pointer to **[]string** | Glob patterns to match on file. e.g. {*.txt, *.cc} | [optional] 
**JobIds** | Pointer to **[]int64** |  | [optional] 
**MaxSnapshotTimestamp** | Pointer to **NullableInt64** | Exclusive end of snapshot_timestamp range. If missing, inf will be used as the timestamp range. | [optional] 
**MinSnapshotTimestamp** | Pointer to **NullableInt64** | Inclusive. If missing, 0 will the default lower end of timestamp range | [optional] 
**PartitionIds** | Pointer to **[]int64** |  | [optional] 
**ProcessLatestOnly** | Pointer to **NullableBool** | Boolean flag to indicate if only latest snapshot of each object should be processed. | [optional] 
**RootDir** | Pointer to **NullableString** | Within each volume, traversal will be rooted at this directory. A typical value here might be /home | [optional] 
**ViewBoxIds** | Pointer to **[]int64** |  | [optional] 
**ViewName** | Pointer to **NullableString** | This is the view name user enters manually. If this is set we will process this view only. partition_id and view_box_id will be populated only if view_name is present. | [optional] 

## Methods

### NewInputSpecInputFilesSelector

`func NewInputSpecInputFilesSelector() *InputSpecInputFilesSelector`

NewInputSpecInputFilesSelector instantiates a new InputSpecInputFilesSelector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputSpecInputFilesSelectorWithDefaults

`func NewInputSpecInputFilesSelectorWithDefaults() *InputSpecInputFilesSelector`

NewInputSpecInputFilesSelectorWithDefaults instantiates a new InputSpecInputFilesSelector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileTimeFilter

`func (o *InputSpecInputFilesSelector) GetFileTimeFilter() InputSpecFileTimeFilter`

GetFileTimeFilter returns the FileTimeFilter field if non-nil, zero value otherwise.

### GetFileTimeFilterOk

`func (o *InputSpecInputFilesSelector) GetFileTimeFilterOk() (*InputSpecFileTimeFilter, bool)`

GetFileTimeFilterOk returns a tuple with the FileTimeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileTimeFilter

`func (o *InputSpecInputFilesSelector) SetFileTimeFilter(v InputSpecFileTimeFilter)`

SetFileTimeFilter sets FileTimeFilter field to given value.

### HasFileTimeFilter

`func (o *InputSpecInputFilesSelector) HasFileTimeFilter() bool`

HasFileTimeFilter returns a boolean if a field has been set.

### GetFilenameGlob

`func (o *InputSpecInputFilesSelector) GetFilenameGlob() []string`

GetFilenameGlob returns the FilenameGlob field if non-nil, zero value otherwise.

### GetFilenameGlobOk

`func (o *InputSpecInputFilesSelector) GetFilenameGlobOk() (*[]string, bool)`

GetFilenameGlobOk returns a tuple with the FilenameGlob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilenameGlob

`func (o *InputSpecInputFilesSelector) SetFilenameGlob(v []string)`

SetFilenameGlob sets FilenameGlob field to given value.

### HasFilenameGlob

`func (o *InputSpecInputFilesSelector) HasFilenameGlob() bool`

HasFilenameGlob returns a boolean if a field has been set.

### SetFilenameGlobNil

`func (o *InputSpecInputFilesSelector) SetFilenameGlobNil(b bool)`

 SetFilenameGlobNil sets the value for FilenameGlob to be an explicit nil

### UnsetFilenameGlob
`func (o *InputSpecInputFilesSelector) UnsetFilenameGlob()`

UnsetFilenameGlob ensures that no value is present for FilenameGlob, not even an explicit nil
### GetJobIds

`func (o *InputSpecInputFilesSelector) GetJobIds() []int64`

GetJobIds returns the JobIds field if non-nil, zero value otherwise.

### GetJobIdsOk

`func (o *InputSpecInputFilesSelector) GetJobIdsOk() (*[]int64, bool)`

GetJobIdsOk returns a tuple with the JobIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobIds

`func (o *InputSpecInputFilesSelector) SetJobIds(v []int64)`

SetJobIds sets JobIds field to given value.

### HasJobIds

`func (o *InputSpecInputFilesSelector) HasJobIds() bool`

HasJobIds returns a boolean if a field has been set.

### SetJobIdsNil

`func (o *InputSpecInputFilesSelector) SetJobIdsNil(b bool)`

 SetJobIdsNil sets the value for JobIds to be an explicit nil

### UnsetJobIds
`func (o *InputSpecInputFilesSelector) UnsetJobIds()`

UnsetJobIds ensures that no value is present for JobIds, not even an explicit nil
### GetMaxSnapshotTimestamp

`func (o *InputSpecInputFilesSelector) GetMaxSnapshotTimestamp() int64`

GetMaxSnapshotTimestamp returns the MaxSnapshotTimestamp field if non-nil, zero value otherwise.

### GetMaxSnapshotTimestampOk

`func (o *InputSpecInputFilesSelector) GetMaxSnapshotTimestampOk() (*int64, bool)`

GetMaxSnapshotTimestampOk returns a tuple with the MaxSnapshotTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSnapshotTimestamp

`func (o *InputSpecInputFilesSelector) SetMaxSnapshotTimestamp(v int64)`

SetMaxSnapshotTimestamp sets MaxSnapshotTimestamp field to given value.

### HasMaxSnapshotTimestamp

`func (o *InputSpecInputFilesSelector) HasMaxSnapshotTimestamp() bool`

HasMaxSnapshotTimestamp returns a boolean if a field has been set.

### SetMaxSnapshotTimestampNil

`func (o *InputSpecInputFilesSelector) SetMaxSnapshotTimestampNil(b bool)`

 SetMaxSnapshotTimestampNil sets the value for MaxSnapshotTimestamp to be an explicit nil

### UnsetMaxSnapshotTimestamp
`func (o *InputSpecInputFilesSelector) UnsetMaxSnapshotTimestamp()`

UnsetMaxSnapshotTimestamp ensures that no value is present for MaxSnapshotTimestamp, not even an explicit nil
### GetMinSnapshotTimestamp

`func (o *InputSpecInputFilesSelector) GetMinSnapshotTimestamp() int64`

GetMinSnapshotTimestamp returns the MinSnapshotTimestamp field if non-nil, zero value otherwise.

### GetMinSnapshotTimestampOk

`func (o *InputSpecInputFilesSelector) GetMinSnapshotTimestampOk() (*int64, bool)`

GetMinSnapshotTimestampOk returns a tuple with the MinSnapshotTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSnapshotTimestamp

`func (o *InputSpecInputFilesSelector) SetMinSnapshotTimestamp(v int64)`

SetMinSnapshotTimestamp sets MinSnapshotTimestamp field to given value.

### HasMinSnapshotTimestamp

`func (o *InputSpecInputFilesSelector) HasMinSnapshotTimestamp() bool`

HasMinSnapshotTimestamp returns a boolean if a field has been set.

### SetMinSnapshotTimestampNil

`func (o *InputSpecInputFilesSelector) SetMinSnapshotTimestampNil(b bool)`

 SetMinSnapshotTimestampNil sets the value for MinSnapshotTimestamp to be an explicit nil

### UnsetMinSnapshotTimestamp
`func (o *InputSpecInputFilesSelector) UnsetMinSnapshotTimestamp()`

UnsetMinSnapshotTimestamp ensures that no value is present for MinSnapshotTimestamp, not even an explicit nil
### GetPartitionIds

`func (o *InputSpecInputFilesSelector) GetPartitionIds() []int64`

GetPartitionIds returns the PartitionIds field if non-nil, zero value otherwise.

### GetPartitionIdsOk

`func (o *InputSpecInputFilesSelector) GetPartitionIdsOk() (*[]int64, bool)`

GetPartitionIdsOk returns a tuple with the PartitionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionIds

`func (o *InputSpecInputFilesSelector) SetPartitionIds(v []int64)`

SetPartitionIds sets PartitionIds field to given value.

### HasPartitionIds

`func (o *InputSpecInputFilesSelector) HasPartitionIds() bool`

HasPartitionIds returns a boolean if a field has been set.

### SetPartitionIdsNil

`func (o *InputSpecInputFilesSelector) SetPartitionIdsNil(b bool)`

 SetPartitionIdsNil sets the value for PartitionIds to be an explicit nil

### UnsetPartitionIds
`func (o *InputSpecInputFilesSelector) UnsetPartitionIds()`

UnsetPartitionIds ensures that no value is present for PartitionIds, not even an explicit nil
### GetProcessLatestOnly

`func (o *InputSpecInputFilesSelector) GetProcessLatestOnly() bool`

GetProcessLatestOnly returns the ProcessLatestOnly field if non-nil, zero value otherwise.

### GetProcessLatestOnlyOk

`func (o *InputSpecInputFilesSelector) GetProcessLatestOnlyOk() (*bool, bool)`

GetProcessLatestOnlyOk returns a tuple with the ProcessLatestOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessLatestOnly

`func (o *InputSpecInputFilesSelector) SetProcessLatestOnly(v bool)`

SetProcessLatestOnly sets ProcessLatestOnly field to given value.

### HasProcessLatestOnly

`func (o *InputSpecInputFilesSelector) HasProcessLatestOnly() bool`

HasProcessLatestOnly returns a boolean if a field has been set.

### SetProcessLatestOnlyNil

`func (o *InputSpecInputFilesSelector) SetProcessLatestOnlyNil(b bool)`

 SetProcessLatestOnlyNil sets the value for ProcessLatestOnly to be an explicit nil

### UnsetProcessLatestOnly
`func (o *InputSpecInputFilesSelector) UnsetProcessLatestOnly()`

UnsetProcessLatestOnly ensures that no value is present for ProcessLatestOnly, not even an explicit nil
### GetRootDir

`func (o *InputSpecInputFilesSelector) GetRootDir() string`

GetRootDir returns the RootDir field if non-nil, zero value otherwise.

### GetRootDirOk

`func (o *InputSpecInputFilesSelector) GetRootDirOk() (*string, bool)`

GetRootDirOk returns a tuple with the RootDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDir

`func (o *InputSpecInputFilesSelector) SetRootDir(v string)`

SetRootDir sets RootDir field to given value.

### HasRootDir

`func (o *InputSpecInputFilesSelector) HasRootDir() bool`

HasRootDir returns a boolean if a field has been set.

### SetRootDirNil

`func (o *InputSpecInputFilesSelector) SetRootDirNil(b bool)`

 SetRootDirNil sets the value for RootDir to be an explicit nil

### UnsetRootDir
`func (o *InputSpecInputFilesSelector) UnsetRootDir()`

UnsetRootDir ensures that no value is present for RootDir, not even an explicit nil
### GetViewBoxIds

`func (o *InputSpecInputFilesSelector) GetViewBoxIds() []int64`

GetViewBoxIds returns the ViewBoxIds field if non-nil, zero value otherwise.

### GetViewBoxIdsOk

`func (o *InputSpecInputFilesSelector) GetViewBoxIdsOk() (*[]int64, bool)`

GetViewBoxIdsOk returns a tuple with the ViewBoxIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewBoxIds

`func (o *InputSpecInputFilesSelector) SetViewBoxIds(v []int64)`

SetViewBoxIds sets ViewBoxIds field to given value.

### HasViewBoxIds

`func (o *InputSpecInputFilesSelector) HasViewBoxIds() bool`

HasViewBoxIds returns a boolean if a field has been set.

### SetViewBoxIdsNil

`func (o *InputSpecInputFilesSelector) SetViewBoxIdsNil(b bool)`

 SetViewBoxIdsNil sets the value for ViewBoxIds to be an explicit nil

### UnsetViewBoxIds
`func (o *InputSpecInputFilesSelector) UnsetViewBoxIds()`

UnsetViewBoxIds ensures that no value is present for ViewBoxIds, not even an explicit nil
### GetViewName

`func (o *InputSpecInputFilesSelector) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *InputSpecInputFilesSelector) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *InputSpecInputFilesSelector) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *InputSpecInputFilesSelector) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### SetViewNameNil

`func (o *InputSpecInputFilesSelector) SetViewNameNil(b bool)`

 SetViewNameNil sets the value for ViewName to be an explicit nil

### UnsetViewName
`func (o *InputSpecInputFilesSelector) UnsetViewName()`

UnsetViewName ensures that no value is present for ViewName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


