# RetrieveArchiveInfoRetrievedEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BytesTransferred** | Pointer to **NullableInt64** | Number of physical bytes transferred over wire for this entity. | [optional] 
**EndTimestampUsecs** | Pointer to **NullableInt64** | Time in microseconds when retrieve of this entity finished or failed. | [optional] 
**Entity** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**Error** | Pointer to [**ErrorProto**](ErrorProto.md) |  | [optional] 
**LogicalBytesTransferred** | Pointer to **NullableInt64** | Number of logical bytes transferred so far. | [optional] 
**LogicalSizeBytes** | Pointer to **NullableInt64** | Total logical size of this entity. | [optional] 
**ProgressMonitorTaskPath** | Pointer to **NullableString** | The path relative to the root path of the retrieval task progress monitor of this entity progress monitor. | [optional] 
**RelativeSnapshotDir** | Pointer to **NullableString** | The path relative to the root of the file system where the snapshot of this entity was retrieved/copied to. | [optional] 
**StartTimestampUsecs** | Pointer to **NullableInt64** | Time in microseconds when retrieve of this entity started. | [optional] 
**Status** | Pointer to **NullableInt32** | The retrieval status of this entity. | [optional] 

## Methods

### NewRetrieveArchiveInfoRetrievedEntity

`func NewRetrieveArchiveInfoRetrievedEntity() *RetrieveArchiveInfoRetrievedEntity`

NewRetrieveArchiveInfoRetrievedEntity instantiates a new RetrieveArchiveInfoRetrievedEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetrieveArchiveInfoRetrievedEntityWithDefaults

`func NewRetrieveArchiveInfoRetrievedEntityWithDefaults() *RetrieveArchiveInfoRetrievedEntity`

NewRetrieveArchiveInfoRetrievedEntityWithDefaults instantiates a new RetrieveArchiveInfoRetrievedEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBytesTransferred

`func (o *RetrieveArchiveInfoRetrievedEntity) GetBytesTransferred() int64`

GetBytesTransferred returns the BytesTransferred field if non-nil, zero value otherwise.

### GetBytesTransferredOk

`func (o *RetrieveArchiveInfoRetrievedEntity) GetBytesTransferredOk() (*int64, bool)`

GetBytesTransferredOk returns a tuple with the BytesTransferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesTransferred

`func (o *RetrieveArchiveInfoRetrievedEntity) SetBytesTransferred(v int64)`

SetBytesTransferred sets BytesTransferred field to given value.

### HasBytesTransferred

`func (o *RetrieveArchiveInfoRetrievedEntity) HasBytesTransferred() bool`

HasBytesTransferred returns a boolean if a field has been set.

### SetBytesTransferredNil

`func (o *RetrieveArchiveInfoRetrievedEntity) SetBytesTransferredNil(b bool)`

 SetBytesTransferredNil sets the value for BytesTransferred to be an explicit nil

### UnsetBytesTransferred
`func (o *RetrieveArchiveInfoRetrievedEntity) UnsetBytesTransferred()`

UnsetBytesTransferred ensures that no value is present for BytesTransferred, not even an explicit nil
### GetEndTimestampUsecs

`func (o *RetrieveArchiveInfoRetrievedEntity) GetEndTimestampUsecs() int64`

GetEndTimestampUsecs returns the EndTimestampUsecs field if non-nil, zero value otherwise.

### GetEndTimestampUsecsOk

`func (o *RetrieveArchiveInfoRetrievedEntity) GetEndTimestampUsecsOk() (*int64, bool)`

GetEndTimestampUsecsOk returns a tuple with the EndTimestampUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimestampUsecs

`func (o *RetrieveArchiveInfoRetrievedEntity) SetEndTimestampUsecs(v int64)`

SetEndTimestampUsecs sets EndTimestampUsecs field to given value.

### HasEndTimestampUsecs

`func (o *RetrieveArchiveInfoRetrievedEntity) HasEndTimestampUsecs() bool`

HasEndTimestampUsecs returns a boolean if a field has been set.

### SetEndTimestampUsecsNil

`func (o *RetrieveArchiveInfoRetrievedEntity) SetEndTimestampUsecsNil(b bool)`

 SetEndTimestampUsecsNil sets the value for EndTimestampUsecs to be an explicit nil

### UnsetEndTimestampUsecs
`func (o *RetrieveArchiveInfoRetrievedEntity) UnsetEndTimestampUsecs()`

UnsetEndTimestampUsecs ensures that no value is present for EndTimestampUsecs, not even an explicit nil
### GetEntity

`func (o *RetrieveArchiveInfoRetrievedEntity) GetEntity() EntityProto`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *RetrieveArchiveInfoRetrievedEntity) GetEntityOk() (*EntityProto, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *RetrieveArchiveInfoRetrievedEntity) SetEntity(v EntityProto)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *RetrieveArchiveInfoRetrievedEntity) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetError

`func (o *RetrieveArchiveInfoRetrievedEntity) GetError() ErrorProto`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RetrieveArchiveInfoRetrievedEntity) GetErrorOk() (*ErrorProto, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RetrieveArchiveInfoRetrievedEntity) SetError(v ErrorProto)`

SetError sets Error field to given value.

### HasError

`func (o *RetrieveArchiveInfoRetrievedEntity) HasError() bool`

HasError returns a boolean if a field has been set.

### GetLogicalBytesTransferred

`func (o *RetrieveArchiveInfoRetrievedEntity) GetLogicalBytesTransferred() int64`

GetLogicalBytesTransferred returns the LogicalBytesTransferred field if non-nil, zero value otherwise.

### GetLogicalBytesTransferredOk

`func (o *RetrieveArchiveInfoRetrievedEntity) GetLogicalBytesTransferredOk() (*int64, bool)`

GetLogicalBytesTransferredOk returns a tuple with the LogicalBytesTransferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalBytesTransferred

`func (o *RetrieveArchiveInfoRetrievedEntity) SetLogicalBytesTransferred(v int64)`

SetLogicalBytesTransferred sets LogicalBytesTransferred field to given value.

### HasLogicalBytesTransferred

`func (o *RetrieveArchiveInfoRetrievedEntity) HasLogicalBytesTransferred() bool`

HasLogicalBytesTransferred returns a boolean if a field has been set.

### SetLogicalBytesTransferredNil

`func (o *RetrieveArchiveInfoRetrievedEntity) SetLogicalBytesTransferredNil(b bool)`

 SetLogicalBytesTransferredNil sets the value for LogicalBytesTransferred to be an explicit nil

### UnsetLogicalBytesTransferred
`func (o *RetrieveArchiveInfoRetrievedEntity) UnsetLogicalBytesTransferred()`

UnsetLogicalBytesTransferred ensures that no value is present for LogicalBytesTransferred, not even an explicit nil
### GetLogicalSizeBytes

`func (o *RetrieveArchiveInfoRetrievedEntity) GetLogicalSizeBytes() int64`

GetLogicalSizeBytes returns the LogicalSizeBytes field if non-nil, zero value otherwise.

### GetLogicalSizeBytesOk

`func (o *RetrieveArchiveInfoRetrievedEntity) GetLogicalSizeBytesOk() (*int64, bool)`

GetLogicalSizeBytesOk returns a tuple with the LogicalSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalSizeBytes

`func (o *RetrieveArchiveInfoRetrievedEntity) SetLogicalSizeBytes(v int64)`

SetLogicalSizeBytes sets LogicalSizeBytes field to given value.

### HasLogicalSizeBytes

`func (o *RetrieveArchiveInfoRetrievedEntity) HasLogicalSizeBytes() bool`

HasLogicalSizeBytes returns a boolean if a field has been set.

### SetLogicalSizeBytesNil

`func (o *RetrieveArchiveInfoRetrievedEntity) SetLogicalSizeBytesNil(b bool)`

 SetLogicalSizeBytesNil sets the value for LogicalSizeBytes to be an explicit nil

### UnsetLogicalSizeBytes
`func (o *RetrieveArchiveInfoRetrievedEntity) UnsetLogicalSizeBytes()`

UnsetLogicalSizeBytes ensures that no value is present for LogicalSizeBytes, not even an explicit nil
### GetProgressMonitorTaskPath

`func (o *RetrieveArchiveInfoRetrievedEntity) GetProgressMonitorTaskPath() string`

GetProgressMonitorTaskPath returns the ProgressMonitorTaskPath field if non-nil, zero value otherwise.

### GetProgressMonitorTaskPathOk

`func (o *RetrieveArchiveInfoRetrievedEntity) GetProgressMonitorTaskPathOk() (*string, bool)`

GetProgressMonitorTaskPathOk returns a tuple with the ProgressMonitorTaskPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressMonitorTaskPath

`func (o *RetrieveArchiveInfoRetrievedEntity) SetProgressMonitorTaskPath(v string)`

SetProgressMonitorTaskPath sets ProgressMonitorTaskPath field to given value.

### HasProgressMonitorTaskPath

`func (o *RetrieveArchiveInfoRetrievedEntity) HasProgressMonitorTaskPath() bool`

HasProgressMonitorTaskPath returns a boolean if a field has been set.

### SetProgressMonitorTaskPathNil

`func (o *RetrieveArchiveInfoRetrievedEntity) SetProgressMonitorTaskPathNil(b bool)`

 SetProgressMonitorTaskPathNil sets the value for ProgressMonitorTaskPath to be an explicit nil

### UnsetProgressMonitorTaskPath
`func (o *RetrieveArchiveInfoRetrievedEntity) UnsetProgressMonitorTaskPath()`

UnsetProgressMonitorTaskPath ensures that no value is present for ProgressMonitorTaskPath, not even an explicit nil
### GetRelativeSnapshotDir

`func (o *RetrieveArchiveInfoRetrievedEntity) GetRelativeSnapshotDir() string`

GetRelativeSnapshotDir returns the RelativeSnapshotDir field if non-nil, zero value otherwise.

### GetRelativeSnapshotDirOk

`func (o *RetrieveArchiveInfoRetrievedEntity) GetRelativeSnapshotDirOk() (*string, bool)`

GetRelativeSnapshotDirOk returns a tuple with the RelativeSnapshotDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeSnapshotDir

`func (o *RetrieveArchiveInfoRetrievedEntity) SetRelativeSnapshotDir(v string)`

SetRelativeSnapshotDir sets RelativeSnapshotDir field to given value.

### HasRelativeSnapshotDir

`func (o *RetrieveArchiveInfoRetrievedEntity) HasRelativeSnapshotDir() bool`

HasRelativeSnapshotDir returns a boolean if a field has been set.

### SetRelativeSnapshotDirNil

`func (o *RetrieveArchiveInfoRetrievedEntity) SetRelativeSnapshotDirNil(b bool)`

 SetRelativeSnapshotDirNil sets the value for RelativeSnapshotDir to be an explicit nil

### UnsetRelativeSnapshotDir
`func (o *RetrieveArchiveInfoRetrievedEntity) UnsetRelativeSnapshotDir()`

UnsetRelativeSnapshotDir ensures that no value is present for RelativeSnapshotDir, not even an explicit nil
### GetStartTimestampUsecs

`func (o *RetrieveArchiveInfoRetrievedEntity) GetStartTimestampUsecs() int64`

GetStartTimestampUsecs returns the StartTimestampUsecs field if non-nil, zero value otherwise.

### GetStartTimestampUsecsOk

`func (o *RetrieveArchiveInfoRetrievedEntity) GetStartTimestampUsecsOk() (*int64, bool)`

GetStartTimestampUsecsOk returns a tuple with the StartTimestampUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestampUsecs

`func (o *RetrieveArchiveInfoRetrievedEntity) SetStartTimestampUsecs(v int64)`

SetStartTimestampUsecs sets StartTimestampUsecs field to given value.

### HasStartTimestampUsecs

`func (o *RetrieveArchiveInfoRetrievedEntity) HasStartTimestampUsecs() bool`

HasStartTimestampUsecs returns a boolean if a field has been set.

### SetStartTimestampUsecsNil

`func (o *RetrieveArchiveInfoRetrievedEntity) SetStartTimestampUsecsNil(b bool)`

 SetStartTimestampUsecsNil sets the value for StartTimestampUsecs to be an explicit nil

### UnsetStartTimestampUsecs
`func (o *RetrieveArchiveInfoRetrievedEntity) UnsetStartTimestampUsecs()`

UnsetStartTimestampUsecs ensures that no value is present for StartTimestampUsecs, not even an explicit nil
### GetStatus

`func (o *RetrieveArchiveInfoRetrievedEntity) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RetrieveArchiveInfoRetrievedEntity) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RetrieveArchiveInfoRetrievedEntity) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RetrieveArchiveInfoRetrievedEntity) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *RetrieveArchiveInfoRetrievedEntity) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *RetrieveArchiveInfoRetrievedEntity) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


