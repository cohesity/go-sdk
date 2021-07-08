# RetrieveArchiveInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvgLogicalTransferRateBps** | Pointer to **NullableInt64** | Average logical bytes transfer rate in bytes per second as seen by Icebox. | [optional] 
**BytesTransferred** | Pointer to **NullableInt64** | Number of physical bytes transferred for this retrieval task so far. | [optional] 
**EndTimeUsecs** | Pointer to **NullableInt64** | Time when this retrieval task ended at Icebox side. If not set, then the retrieval has not ended yet. | [optional] 
**Error** | Pointer to [**ErrorProto**](ErrorProto.md) |  | [optional] 
**LogicalBytesTransferred** | Pointer to **NullableInt64** | Number of logical bytes transferred so far. | [optional] 
**LogicalSizeBytes** | Pointer to **NullableInt64** | Total logical size of the retrieval task. | [optional] 
**ProgressMonitorTaskPath** | Pointer to **NullableString** | The root path of the progress monitor for this task. | [optional] 
**RetrievedEntityVec** | Pointer to [**[]RetrieveArchiveInfoRetrievedEntity**](RetrieveArchiveInfoRetrievedEntity.md) | Contains info about all retrieved entities. | [optional] 
**StartTimeUsecs** | Pointer to **NullableInt64** | Time when this retrieval task was started by Icebox. If not set, then retrieval has not been started yet. | [optional] 
**StubViewName** | Pointer to **NullableString** | The stub view that Icebox created. Stub view can be used for selectively restoring or accessing files from an archive location. | [optional] 
**StubViewRelativeDirName** | Pointer to **NullableString** | Relative directory inside the stub view that corresponds with the archive. | [optional] 
**TargetViewName** | Pointer to **NullableString** | The name of the target view where Icebox has retrieved and staged the requested entities. | [optional] 
**UserActionRequiredMsg** | Pointer to **NullableString** | Message to display in the UI if any manual intervention is needed to make forward progress for the retrieve from archive task. This message is mainly relevant for tape based retrieve from archive tasks where a backup admin might be asked to load new media when the tape library does not have the relevant media to retrieve the archive from. | [optional] 

## Methods

### NewRetrieveArchiveInfo

`func NewRetrieveArchiveInfo() *RetrieveArchiveInfo`

NewRetrieveArchiveInfo instantiates a new RetrieveArchiveInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetrieveArchiveInfoWithDefaults

`func NewRetrieveArchiveInfoWithDefaults() *RetrieveArchiveInfo`

NewRetrieveArchiveInfoWithDefaults instantiates a new RetrieveArchiveInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvgLogicalTransferRateBps

`func (o *RetrieveArchiveInfo) GetAvgLogicalTransferRateBps() int64`

GetAvgLogicalTransferRateBps returns the AvgLogicalTransferRateBps field if non-nil, zero value otherwise.

### GetAvgLogicalTransferRateBpsOk

`func (o *RetrieveArchiveInfo) GetAvgLogicalTransferRateBpsOk() (*int64, bool)`

GetAvgLogicalTransferRateBpsOk returns a tuple with the AvgLogicalTransferRateBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgLogicalTransferRateBps

`func (o *RetrieveArchiveInfo) SetAvgLogicalTransferRateBps(v int64)`

SetAvgLogicalTransferRateBps sets AvgLogicalTransferRateBps field to given value.

### HasAvgLogicalTransferRateBps

`func (o *RetrieveArchiveInfo) HasAvgLogicalTransferRateBps() bool`

HasAvgLogicalTransferRateBps returns a boolean if a field has been set.

### SetAvgLogicalTransferRateBpsNil

`func (o *RetrieveArchiveInfo) SetAvgLogicalTransferRateBpsNil(b bool)`

 SetAvgLogicalTransferRateBpsNil sets the value for AvgLogicalTransferRateBps to be an explicit nil

### UnsetAvgLogicalTransferRateBps
`func (o *RetrieveArchiveInfo) UnsetAvgLogicalTransferRateBps()`

UnsetAvgLogicalTransferRateBps ensures that no value is present for AvgLogicalTransferRateBps, not even an explicit nil
### GetBytesTransferred

`func (o *RetrieveArchiveInfo) GetBytesTransferred() int64`

GetBytesTransferred returns the BytesTransferred field if non-nil, zero value otherwise.

### GetBytesTransferredOk

`func (o *RetrieveArchiveInfo) GetBytesTransferredOk() (*int64, bool)`

GetBytesTransferredOk returns a tuple with the BytesTransferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesTransferred

`func (o *RetrieveArchiveInfo) SetBytesTransferred(v int64)`

SetBytesTransferred sets BytesTransferred field to given value.

### HasBytesTransferred

`func (o *RetrieveArchiveInfo) HasBytesTransferred() bool`

HasBytesTransferred returns a boolean if a field has been set.

### SetBytesTransferredNil

`func (o *RetrieveArchiveInfo) SetBytesTransferredNil(b bool)`

 SetBytesTransferredNil sets the value for BytesTransferred to be an explicit nil

### UnsetBytesTransferred
`func (o *RetrieveArchiveInfo) UnsetBytesTransferred()`

UnsetBytesTransferred ensures that no value is present for BytesTransferred, not even an explicit nil
### GetEndTimeUsecs

`func (o *RetrieveArchiveInfo) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *RetrieveArchiveInfo) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *RetrieveArchiveInfo) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *RetrieveArchiveInfo) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *RetrieveArchiveInfo) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *RetrieveArchiveInfo) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetError

`func (o *RetrieveArchiveInfo) GetError() ErrorProto`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RetrieveArchiveInfo) GetErrorOk() (*ErrorProto, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RetrieveArchiveInfo) SetError(v ErrorProto)`

SetError sets Error field to given value.

### HasError

`func (o *RetrieveArchiveInfo) HasError() bool`

HasError returns a boolean if a field has been set.

### GetLogicalBytesTransferred

`func (o *RetrieveArchiveInfo) GetLogicalBytesTransferred() int64`

GetLogicalBytesTransferred returns the LogicalBytesTransferred field if non-nil, zero value otherwise.

### GetLogicalBytesTransferredOk

`func (o *RetrieveArchiveInfo) GetLogicalBytesTransferredOk() (*int64, bool)`

GetLogicalBytesTransferredOk returns a tuple with the LogicalBytesTransferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalBytesTransferred

`func (o *RetrieveArchiveInfo) SetLogicalBytesTransferred(v int64)`

SetLogicalBytesTransferred sets LogicalBytesTransferred field to given value.

### HasLogicalBytesTransferred

`func (o *RetrieveArchiveInfo) HasLogicalBytesTransferred() bool`

HasLogicalBytesTransferred returns a boolean if a field has been set.

### SetLogicalBytesTransferredNil

`func (o *RetrieveArchiveInfo) SetLogicalBytesTransferredNil(b bool)`

 SetLogicalBytesTransferredNil sets the value for LogicalBytesTransferred to be an explicit nil

### UnsetLogicalBytesTransferred
`func (o *RetrieveArchiveInfo) UnsetLogicalBytesTransferred()`

UnsetLogicalBytesTransferred ensures that no value is present for LogicalBytesTransferred, not even an explicit nil
### GetLogicalSizeBytes

`func (o *RetrieveArchiveInfo) GetLogicalSizeBytes() int64`

GetLogicalSizeBytes returns the LogicalSizeBytes field if non-nil, zero value otherwise.

### GetLogicalSizeBytesOk

`func (o *RetrieveArchiveInfo) GetLogicalSizeBytesOk() (*int64, bool)`

GetLogicalSizeBytesOk returns a tuple with the LogicalSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalSizeBytes

`func (o *RetrieveArchiveInfo) SetLogicalSizeBytes(v int64)`

SetLogicalSizeBytes sets LogicalSizeBytes field to given value.

### HasLogicalSizeBytes

`func (o *RetrieveArchiveInfo) HasLogicalSizeBytes() bool`

HasLogicalSizeBytes returns a boolean if a field has been set.

### SetLogicalSizeBytesNil

`func (o *RetrieveArchiveInfo) SetLogicalSizeBytesNil(b bool)`

 SetLogicalSizeBytesNil sets the value for LogicalSizeBytes to be an explicit nil

### UnsetLogicalSizeBytes
`func (o *RetrieveArchiveInfo) UnsetLogicalSizeBytes()`

UnsetLogicalSizeBytes ensures that no value is present for LogicalSizeBytes, not even an explicit nil
### GetProgressMonitorTaskPath

`func (o *RetrieveArchiveInfo) GetProgressMonitorTaskPath() string`

GetProgressMonitorTaskPath returns the ProgressMonitorTaskPath field if non-nil, zero value otherwise.

### GetProgressMonitorTaskPathOk

`func (o *RetrieveArchiveInfo) GetProgressMonitorTaskPathOk() (*string, bool)`

GetProgressMonitorTaskPathOk returns a tuple with the ProgressMonitorTaskPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressMonitorTaskPath

`func (o *RetrieveArchiveInfo) SetProgressMonitorTaskPath(v string)`

SetProgressMonitorTaskPath sets ProgressMonitorTaskPath field to given value.

### HasProgressMonitorTaskPath

`func (o *RetrieveArchiveInfo) HasProgressMonitorTaskPath() bool`

HasProgressMonitorTaskPath returns a boolean if a field has been set.

### SetProgressMonitorTaskPathNil

`func (o *RetrieveArchiveInfo) SetProgressMonitorTaskPathNil(b bool)`

 SetProgressMonitorTaskPathNil sets the value for ProgressMonitorTaskPath to be an explicit nil

### UnsetProgressMonitorTaskPath
`func (o *RetrieveArchiveInfo) UnsetProgressMonitorTaskPath()`

UnsetProgressMonitorTaskPath ensures that no value is present for ProgressMonitorTaskPath, not even an explicit nil
### GetRetrievedEntityVec

`func (o *RetrieveArchiveInfo) GetRetrievedEntityVec() []RetrieveArchiveInfoRetrievedEntity`

GetRetrievedEntityVec returns the RetrievedEntityVec field if non-nil, zero value otherwise.

### GetRetrievedEntityVecOk

`func (o *RetrieveArchiveInfo) GetRetrievedEntityVecOk() (*[]RetrieveArchiveInfoRetrievedEntity, bool)`

GetRetrievedEntityVecOk returns a tuple with the RetrievedEntityVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetrievedEntityVec

`func (o *RetrieveArchiveInfo) SetRetrievedEntityVec(v []RetrieveArchiveInfoRetrievedEntity)`

SetRetrievedEntityVec sets RetrievedEntityVec field to given value.

### HasRetrievedEntityVec

`func (o *RetrieveArchiveInfo) HasRetrievedEntityVec() bool`

HasRetrievedEntityVec returns a boolean if a field has been set.

### SetRetrievedEntityVecNil

`func (o *RetrieveArchiveInfo) SetRetrievedEntityVecNil(b bool)`

 SetRetrievedEntityVecNil sets the value for RetrievedEntityVec to be an explicit nil

### UnsetRetrievedEntityVec
`func (o *RetrieveArchiveInfo) UnsetRetrievedEntityVec()`

UnsetRetrievedEntityVec ensures that no value is present for RetrievedEntityVec, not even an explicit nil
### GetStartTimeUsecs

`func (o *RetrieveArchiveInfo) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *RetrieveArchiveInfo) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *RetrieveArchiveInfo) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *RetrieveArchiveInfo) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *RetrieveArchiveInfo) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *RetrieveArchiveInfo) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetStubViewName

`func (o *RetrieveArchiveInfo) GetStubViewName() string`

GetStubViewName returns the StubViewName field if non-nil, zero value otherwise.

### GetStubViewNameOk

`func (o *RetrieveArchiveInfo) GetStubViewNameOk() (*string, bool)`

GetStubViewNameOk returns a tuple with the StubViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStubViewName

`func (o *RetrieveArchiveInfo) SetStubViewName(v string)`

SetStubViewName sets StubViewName field to given value.

### HasStubViewName

`func (o *RetrieveArchiveInfo) HasStubViewName() bool`

HasStubViewName returns a boolean if a field has been set.

### SetStubViewNameNil

`func (o *RetrieveArchiveInfo) SetStubViewNameNil(b bool)`

 SetStubViewNameNil sets the value for StubViewName to be an explicit nil

### UnsetStubViewName
`func (o *RetrieveArchiveInfo) UnsetStubViewName()`

UnsetStubViewName ensures that no value is present for StubViewName, not even an explicit nil
### GetStubViewRelativeDirName

`func (o *RetrieveArchiveInfo) GetStubViewRelativeDirName() string`

GetStubViewRelativeDirName returns the StubViewRelativeDirName field if non-nil, zero value otherwise.

### GetStubViewRelativeDirNameOk

`func (o *RetrieveArchiveInfo) GetStubViewRelativeDirNameOk() (*string, bool)`

GetStubViewRelativeDirNameOk returns a tuple with the StubViewRelativeDirName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStubViewRelativeDirName

`func (o *RetrieveArchiveInfo) SetStubViewRelativeDirName(v string)`

SetStubViewRelativeDirName sets StubViewRelativeDirName field to given value.

### HasStubViewRelativeDirName

`func (o *RetrieveArchiveInfo) HasStubViewRelativeDirName() bool`

HasStubViewRelativeDirName returns a boolean if a field has been set.

### SetStubViewRelativeDirNameNil

`func (o *RetrieveArchiveInfo) SetStubViewRelativeDirNameNil(b bool)`

 SetStubViewRelativeDirNameNil sets the value for StubViewRelativeDirName to be an explicit nil

### UnsetStubViewRelativeDirName
`func (o *RetrieveArchiveInfo) UnsetStubViewRelativeDirName()`

UnsetStubViewRelativeDirName ensures that no value is present for StubViewRelativeDirName, not even an explicit nil
### GetTargetViewName

`func (o *RetrieveArchiveInfo) GetTargetViewName() string`

GetTargetViewName returns the TargetViewName field if non-nil, zero value otherwise.

### GetTargetViewNameOk

`func (o *RetrieveArchiveInfo) GetTargetViewNameOk() (*string, bool)`

GetTargetViewNameOk returns a tuple with the TargetViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetViewName

`func (o *RetrieveArchiveInfo) SetTargetViewName(v string)`

SetTargetViewName sets TargetViewName field to given value.

### HasTargetViewName

`func (o *RetrieveArchiveInfo) HasTargetViewName() bool`

HasTargetViewName returns a boolean if a field has been set.

### SetTargetViewNameNil

`func (o *RetrieveArchiveInfo) SetTargetViewNameNil(b bool)`

 SetTargetViewNameNil sets the value for TargetViewName to be an explicit nil

### UnsetTargetViewName
`func (o *RetrieveArchiveInfo) UnsetTargetViewName()`

UnsetTargetViewName ensures that no value is present for TargetViewName, not even an explicit nil
### GetUserActionRequiredMsg

`func (o *RetrieveArchiveInfo) GetUserActionRequiredMsg() string`

GetUserActionRequiredMsg returns the UserActionRequiredMsg field if non-nil, zero value otherwise.

### GetUserActionRequiredMsgOk

`func (o *RetrieveArchiveInfo) GetUserActionRequiredMsgOk() (*string, bool)`

GetUserActionRequiredMsgOk returns a tuple with the UserActionRequiredMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserActionRequiredMsg

`func (o *RetrieveArchiveInfo) SetUserActionRequiredMsg(v string)`

SetUserActionRequiredMsg sets UserActionRequiredMsg field to given value.

### HasUserActionRequiredMsg

`func (o *RetrieveArchiveInfo) HasUserActionRequiredMsg() bool`

HasUserActionRequiredMsg returns a boolean if a field has been set.

### SetUserActionRequiredMsgNil

`func (o *RetrieveArchiveInfo) SetUserActionRequiredMsgNil(b bool)`

 SetUserActionRequiredMsgNil sets the value for UserActionRequiredMsg to be an explicit nil

### UnsetUserActionRequiredMsg
`func (o *RetrieveArchiveInfo) UnsetUserActionRequiredMsg()`

UnsetUserActionRequiredMsg ensures that no value is present for UserActionRequiredMsg, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


