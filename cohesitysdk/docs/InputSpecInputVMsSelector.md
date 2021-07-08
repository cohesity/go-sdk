# InputSpecInputVMsSelector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileTimeFilter** | Pointer to [**InputSpecFileTimeFilter**](InputSpecFileTimeFilter.md) |  | [optional] 
**FilenameGlob** | Pointer to **[]string** | After VMDKs are selected as above, the files within them can be selected by using these predicates. | [optional] 
**JobIds** | Pointer to **[]int64** |  | [optional] 
**MaxSnapshotTimestamp** | Pointer to **NullableInt64** | Exclusive end of snapshot_timestamp range. If missing, inf will be used as the timestamp range. | [optional] 
**MinSnapshotTimestamp** | Pointer to **NullableInt64** | Inclusive. If missing, 0 will the default lower end of timestamp range | [optional] 
**PartitionIds** | Pointer to **[]int64** | Filters are AND of ORs. | [optional] 
**ProcessLatestOnly** | Pointer to **NullableBool** | Boolean flag to indicate if only latest snapshot of each object should be processed. | [optional] 
**RootDir** | Pointer to **NullableString** | Within each volume, traversal will be rooted at this directory. A typical value here might be /home | [optional] 
**SourceEntityIds** | Pointer to **[]int64** |  | [optional] 
**ViewBoxIds** | Pointer to **[]int64** |  | [optional] 
**ViewNames** | Pointer to **[]string** |  | [optional] 

## Methods

### NewInputSpecInputVMsSelector

`func NewInputSpecInputVMsSelector() *InputSpecInputVMsSelector`

NewInputSpecInputVMsSelector instantiates a new InputSpecInputVMsSelector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputSpecInputVMsSelectorWithDefaults

`func NewInputSpecInputVMsSelectorWithDefaults() *InputSpecInputVMsSelector`

NewInputSpecInputVMsSelectorWithDefaults instantiates a new InputSpecInputVMsSelector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileTimeFilter

`func (o *InputSpecInputVMsSelector) GetFileTimeFilter() InputSpecFileTimeFilter`

GetFileTimeFilter returns the FileTimeFilter field if non-nil, zero value otherwise.

### GetFileTimeFilterOk

`func (o *InputSpecInputVMsSelector) GetFileTimeFilterOk() (*InputSpecFileTimeFilter, bool)`

GetFileTimeFilterOk returns a tuple with the FileTimeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileTimeFilter

`func (o *InputSpecInputVMsSelector) SetFileTimeFilter(v InputSpecFileTimeFilter)`

SetFileTimeFilter sets FileTimeFilter field to given value.

### HasFileTimeFilter

`func (o *InputSpecInputVMsSelector) HasFileTimeFilter() bool`

HasFileTimeFilter returns a boolean if a field has been set.

### GetFilenameGlob

`func (o *InputSpecInputVMsSelector) GetFilenameGlob() []string`

GetFilenameGlob returns the FilenameGlob field if non-nil, zero value otherwise.

### GetFilenameGlobOk

`func (o *InputSpecInputVMsSelector) GetFilenameGlobOk() (*[]string, bool)`

GetFilenameGlobOk returns a tuple with the FilenameGlob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilenameGlob

`func (o *InputSpecInputVMsSelector) SetFilenameGlob(v []string)`

SetFilenameGlob sets FilenameGlob field to given value.

### HasFilenameGlob

`func (o *InputSpecInputVMsSelector) HasFilenameGlob() bool`

HasFilenameGlob returns a boolean if a field has been set.

### SetFilenameGlobNil

`func (o *InputSpecInputVMsSelector) SetFilenameGlobNil(b bool)`

 SetFilenameGlobNil sets the value for FilenameGlob to be an explicit nil

### UnsetFilenameGlob
`func (o *InputSpecInputVMsSelector) UnsetFilenameGlob()`

UnsetFilenameGlob ensures that no value is present for FilenameGlob, not even an explicit nil
### GetJobIds

`func (o *InputSpecInputVMsSelector) GetJobIds() []int64`

GetJobIds returns the JobIds field if non-nil, zero value otherwise.

### GetJobIdsOk

`func (o *InputSpecInputVMsSelector) GetJobIdsOk() (*[]int64, bool)`

GetJobIdsOk returns a tuple with the JobIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobIds

`func (o *InputSpecInputVMsSelector) SetJobIds(v []int64)`

SetJobIds sets JobIds field to given value.

### HasJobIds

`func (o *InputSpecInputVMsSelector) HasJobIds() bool`

HasJobIds returns a boolean if a field has been set.

### SetJobIdsNil

`func (o *InputSpecInputVMsSelector) SetJobIdsNil(b bool)`

 SetJobIdsNil sets the value for JobIds to be an explicit nil

### UnsetJobIds
`func (o *InputSpecInputVMsSelector) UnsetJobIds()`

UnsetJobIds ensures that no value is present for JobIds, not even an explicit nil
### GetMaxSnapshotTimestamp

`func (o *InputSpecInputVMsSelector) GetMaxSnapshotTimestamp() int64`

GetMaxSnapshotTimestamp returns the MaxSnapshotTimestamp field if non-nil, zero value otherwise.

### GetMaxSnapshotTimestampOk

`func (o *InputSpecInputVMsSelector) GetMaxSnapshotTimestampOk() (*int64, bool)`

GetMaxSnapshotTimestampOk returns a tuple with the MaxSnapshotTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSnapshotTimestamp

`func (o *InputSpecInputVMsSelector) SetMaxSnapshotTimestamp(v int64)`

SetMaxSnapshotTimestamp sets MaxSnapshotTimestamp field to given value.

### HasMaxSnapshotTimestamp

`func (o *InputSpecInputVMsSelector) HasMaxSnapshotTimestamp() bool`

HasMaxSnapshotTimestamp returns a boolean if a field has been set.

### SetMaxSnapshotTimestampNil

`func (o *InputSpecInputVMsSelector) SetMaxSnapshotTimestampNil(b bool)`

 SetMaxSnapshotTimestampNil sets the value for MaxSnapshotTimestamp to be an explicit nil

### UnsetMaxSnapshotTimestamp
`func (o *InputSpecInputVMsSelector) UnsetMaxSnapshotTimestamp()`

UnsetMaxSnapshotTimestamp ensures that no value is present for MaxSnapshotTimestamp, not even an explicit nil
### GetMinSnapshotTimestamp

`func (o *InputSpecInputVMsSelector) GetMinSnapshotTimestamp() int64`

GetMinSnapshotTimestamp returns the MinSnapshotTimestamp field if non-nil, zero value otherwise.

### GetMinSnapshotTimestampOk

`func (o *InputSpecInputVMsSelector) GetMinSnapshotTimestampOk() (*int64, bool)`

GetMinSnapshotTimestampOk returns a tuple with the MinSnapshotTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSnapshotTimestamp

`func (o *InputSpecInputVMsSelector) SetMinSnapshotTimestamp(v int64)`

SetMinSnapshotTimestamp sets MinSnapshotTimestamp field to given value.

### HasMinSnapshotTimestamp

`func (o *InputSpecInputVMsSelector) HasMinSnapshotTimestamp() bool`

HasMinSnapshotTimestamp returns a boolean if a field has been set.

### SetMinSnapshotTimestampNil

`func (o *InputSpecInputVMsSelector) SetMinSnapshotTimestampNil(b bool)`

 SetMinSnapshotTimestampNil sets the value for MinSnapshotTimestamp to be an explicit nil

### UnsetMinSnapshotTimestamp
`func (o *InputSpecInputVMsSelector) UnsetMinSnapshotTimestamp()`

UnsetMinSnapshotTimestamp ensures that no value is present for MinSnapshotTimestamp, not even an explicit nil
### GetPartitionIds

`func (o *InputSpecInputVMsSelector) GetPartitionIds() []int64`

GetPartitionIds returns the PartitionIds field if non-nil, zero value otherwise.

### GetPartitionIdsOk

`func (o *InputSpecInputVMsSelector) GetPartitionIdsOk() (*[]int64, bool)`

GetPartitionIdsOk returns a tuple with the PartitionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionIds

`func (o *InputSpecInputVMsSelector) SetPartitionIds(v []int64)`

SetPartitionIds sets PartitionIds field to given value.

### HasPartitionIds

`func (o *InputSpecInputVMsSelector) HasPartitionIds() bool`

HasPartitionIds returns a boolean if a field has been set.

### SetPartitionIdsNil

`func (o *InputSpecInputVMsSelector) SetPartitionIdsNil(b bool)`

 SetPartitionIdsNil sets the value for PartitionIds to be an explicit nil

### UnsetPartitionIds
`func (o *InputSpecInputVMsSelector) UnsetPartitionIds()`

UnsetPartitionIds ensures that no value is present for PartitionIds, not even an explicit nil
### GetProcessLatestOnly

`func (o *InputSpecInputVMsSelector) GetProcessLatestOnly() bool`

GetProcessLatestOnly returns the ProcessLatestOnly field if non-nil, zero value otherwise.

### GetProcessLatestOnlyOk

`func (o *InputSpecInputVMsSelector) GetProcessLatestOnlyOk() (*bool, bool)`

GetProcessLatestOnlyOk returns a tuple with the ProcessLatestOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessLatestOnly

`func (o *InputSpecInputVMsSelector) SetProcessLatestOnly(v bool)`

SetProcessLatestOnly sets ProcessLatestOnly field to given value.

### HasProcessLatestOnly

`func (o *InputSpecInputVMsSelector) HasProcessLatestOnly() bool`

HasProcessLatestOnly returns a boolean if a field has been set.

### SetProcessLatestOnlyNil

`func (o *InputSpecInputVMsSelector) SetProcessLatestOnlyNil(b bool)`

 SetProcessLatestOnlyNil sets the value for ProcessLatestOnly to be an explicit nil

### UnsetProcessLatestOnly
`func (o *InputSpecInputVMsSelector) UnsetProcessLatestOnly()`

UnsetProcessLatestOnly ensures that no value is present for ProcessLatestOnly, not even an explicit nil
### GetRootDir

`func (o *InputSpecInputVMsSelector) GetRootDir() string`

GetRootDir returns the RootDir field if non-nil, zero value otherwise.

### GetRootDirOk

`func (o *InputSpecInputVMsSelector) GetRootDirOk() (*string, bool)`

GetRootDirOk returns a tuple with the RootDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDir

`func (o *InputSpecInputVMsSelector) SetRootDir(v string)`

SetRootDir sets RootDir field to given value.

### HasRootDir

`func (o *InputSpecInputVMsSelector) HasRootDir() bool`

HasRootDir returns a boolean if a field has been set.

### SetRootDirNil

`func (o *InputSpecInputVMsSelector) SetRootDirNil(b bool)`

 SetRootDirNil sets the value for RootDir to be an explicit nil

### UnsetRootDir
`func (o *InputSpecInputVMsSelector) UnsetRootDir()`

UnsetRootDir ensures that no value is present for RootDir, not even an explicit nil
### GetSourceEntityIds

`func (o *InputSpecInputVMsSelector) GetSourceEntityIds() []int64`

GetSourceEntityIds returns the SourceEntityIds field if non-nil, zero value otherwise.

### GetSourceEntityIdsOk

`func (o *InputSpecInputVMsSelector) GetSourceEntityIdsOk() (*[]int64, bool)`

GetSourceEntityIdsOk returns a tuple with the SourceEntityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceEntityIds

`func (o *InputSpecInputVMsSelector) SetSourceEntityIds(v []int64)`

SetSourceEntityIds sets SourceEntityIds field to given value.

### HasSourceEntityIds

`func (o *InputSpecInputVMsSelector) HasSourceEntityIds() bool`

HasSourceEntityIds returns a boolean if a field has been set.

### SetSourceEntityIdsNil

`func (o *InputSpecInputVMsSelector) SetSourceEntityIdsNil(b bool)`

 SetSourceEntityIdsNil sets the value for SourceEntityIds to be an explicit nil

### UnsetSourceEntityIds
`func (o *InputSpecInputVMsSelector) UnsetSourceEntityIds()`

UnsetSourceEntityIds ensures that no value is present for SourceEntityIds, not even an explicit nil
### GetViewBoxIds

`func (o *InputSpecInputVMsSelector) GetViewBoxIds() []int64`

GetViewBoxIds returns the ViewBoxIds field if non-nil, zero value otherwise.

### GetViewBoxIdsOk

`func (o *InputSpecInputVMsSelector) GetViewBoxIdsOk() (*[]int64, bool)`

GetViewBoxIdsOk returns a tuple with the ViewBoxIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewBoxIds

`func (o *InputSpecInputVMsSelector) SetViewBoxIds(v []int64)`

SetViewBoxIds sets ViewBoxIds field to given value.

### HasViewBoxIds

`func (o *InputSpecInputVMsSelector) HasViewBoxIds() bool`

HasViewBoxIds returns a boolean if a field has been set.

### SetViewBoxIdsNil

`func (o *InputSpecInputVMsSelector) SetViewBoxIdsNil(b bool)`

 SetViewBoxIdsNil sets the value for ViewBoxIds to be an explicit nil

### UnsetViewBoxIds
`func (o *InputSpecInputVMsSelector) UnsetViewBoxIds()`

UnsetViewBoxIds ensures that no value is present for ViewBoxIds, not even an explicit nil
### GetViewNames

`func (o *InputSpecInputVMsSelector) GetViewNames() []string`

GetViewNames returns the ViewNames field if non-nil, zero value otherwise.

### GetViewNamesOk

`func (o *InputSpecInputVMsSelector) GetViewNamesOk() (*[]string, bool)`

GetViewNamesOk returns a tuple with the ViewNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewNames

`func (o *InputSpecInputVMsSelector) SetViewNames(v []string)`

SetViewNames sets ViewNames field to given value.

### HasViewNames

`func (o *InputSpecInputVMsSelector) HasViewNames() bool`

HasViewNames returns a boolean if a field has been set.

### SetViewNamesNil

`func (o *InputSpecInputVMsSelector) SetViewNamesNil(b bool)`

 SetViewNamesNil sets the value for ViewNames to be an explicit nil

### UnsetViewNames
`func (o *InputSpecInputVMsSelector) UnsetViewNames()`

UnsetViewNames ensures that no value is present for ViewNames, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


