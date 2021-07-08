# ProtectedSourceSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupRun** | Pointer to [**BackupRun**](BackupRun.md) |  | [optional] 
**CopyRuns** | Pointer to [**[]CopyRun**](CopyRun.md) | Specifies details about the Copy tasks of the Job Run. A Copy task copies the captured snapshots to an external target or a Remote Cohesity Cluster. | [optional] 
**IsPaused** | Pointer to **NullableBool** | Specifies the status of the backup job. | [optional] 
**NextProtectionRunTimeUsecs** | Pointer to **NullableInt64** | Specifies the time at which the next Protection Run is scheduled for the given Protection Source in Unix epoch Time (microseconds). | [optional] 
**ProtectedSourceUid** | Pointer to [**UniversalId**](UniversalId.md) |  | [optional] 
**ProtectionSource** | Pointer to [**ProtectionSource**](ProtectionSource.md) |  | [optional] 
**SourceParameters** | Pointer to [**[]SourceSpecialParameter**](SourceSpecialParameter.md) | Specifies additional special settings for a single Protected Source. | [optional] 

## Methods

### NewProtectedSourceSummary

`func NewProtectedSourceSummary() *ProtectedSourceSummary`

NewProtectedSourceSummary instantiates a new ProtectedSourceSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectedSourceSummaryWithDefaults

`func NewProtectedSourceSummaryWithDefaults() *ProtectedSourceSummary`

NewProtectedSourceSummaryWithDefaults instantiates a new ProtectedSourceSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupRun

`func (o *ProtectedSourceSummary) GetBackupRun() BackupRun`

GetBackupRun returns the BackupRun field if non-nil, zero value otherwise.

### GetBackupRunOk

`func (o *ProtectedSourceSummary) GetBackupRunOk() (*BackupRun, bool)`

GetBackupRunOk returns a tuple with the BackupRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRun

`func (o *ProtectedSourceSummary) SetBackupRun(v BackupRun)`

SetBackupRun sets BackupRun field to given value.

### HasBackupRun

`func (o *ProtectedSourceSummary) HasBackupRun() bool`

HasBackupRun returns a boolean if a field has been set.

### GetCopyRuns

`func (o *ProtectedSourceSummary) GetCopyRuns() []CopyRun`

GetCopyRuns returns the CopyRuns field if non-nil, zero value otherwise.

### GetCopyRunsOk

`func (o *ProtectedSourceSummary) GetCopyRunsOk() (*[]CopyRun, bool)`

GetCopyRunsOk returns a tuple with the CopyRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyRuns

`func (o *ProtectedSourceSummary) SetCopyRuns(v []CopyRun)`

SetCopyRuns sets CopyRuns field to given value.

### HasCopyRuns

`func (o *ProtectedSourceSummary) HasCopyRuns() bool`

HasCopyRuns returns a boolean if a field has been set.

### SetCopyRunsNil

`func (o *ProtectedSourceSummary) SetCopyRunsNil(b bool)`

 SetCopyRunsNil sets the value for CopyRuns to be an explicit nil

### UnsetCopyRuns
`func (o *ProtectedSourceSummary) UnsetCopyRuns()`

UnsetCopyRuns ensures that no value is present for CopyRuns, not even an explicit nil
### GetIsPaused

`func (o *ProtectedSourceSummary) GetIsPaused() bool`

GetIsPaused returns the IsPaused field if non-nil, zero value otherwise.

### GetIsPausedOk

`func (o *ProtectedSourceSummary) GetIsPausedOk() (*bool, bool)`

GetIsPausedOk returns a tuple with the IsPaused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPaused

`func (o *ProtectedSourceSummary) SetIsPaused(v bool)`

SetIsPaused sets IsPaused field to given value.

### HasIsPaused

`func (o *ProtectedSourceSummary) HasIsPaused() bool`

HasIsPaused returns a boolean if a field has been set.

### SetIsPausedNil

`func (o *ProtectedSourceSummary) SetIsPausedNil(b bool)`

 SetIsPausedNil sets the value for IsPaused to be an explicit nil

### UnsetIsPaused
`func (o *ProtectedSourceSummary) UnsetIsPaused()`

UnsetIsPaused ensures that no value is present for IsPaused, not even an explicit nil
### GetNextProtectionRunTimeUsecs

`func (o *ProtectedSourceSummary) GetNextProtectionRunTimeUsecs() int64`

GetNextProtectionRunTimeUsecs returns the NextProtectionRunTimeUsecs field if non-nil, zero value otherwise.

### GetNextProtectionRunTimeUsecsOk

`func (o *ProtectedSourceSummary) GetNextProtectionRunTimeUsecsOk() (*int64, bool)`

GetNextProtectionRunTimeUsecsOk returns a tuple with the NextProtectionRunTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextProtectionRunTimeUsecs

`func (o *ProtectedSourceSummary) SetNextProtectionRunTimeUsecs(v int64)`

SetNextProtectionRunTimeUsecs sets NextProtectionRunTimeUsecs field to given value.

### HasNextProtectionRunTimeUsecs

`func (o *ProtectedSourceSummary) HasNextProtectionRunTimeUsecs() bool`

HasNextProtectionRunTimeUsecs returns a boolean if a field has been set.

### SetNextProtectionRunTimeUsecsNil

`func (o *ProtectedSourceSummary) SetNextProtectionRunTimeUsecsNil(b bool)`

 SetNextProtectionRunTimeUsecsNil sets the value for NextProtectionRunTimeUsecs to be an explicit nil

### UnsetNextProtectionRunTimeUsecs
`func (o *ProtectedSourceSummary) UnsetNextProtectionRunTimeUsecs()`

UnsetNextProtectionRunTimeUsecs ensures that no value is present for NextProtectionRunTimeUsecs, not even an explicit nil
### GetProtectedSourceUid

`func (o *ProtectedSourceSummary) GetProtectedSourceUid() UniversalId`

GetProtectedSourceUid returns the ProtectedSourceUid field if non-nil, zero value otherwise.

### GetProtectedSourceUidOk

`func (o *ProtectedSourceSummary) GetProtectedSourceUidOk() (*UniversalId, bool)`

GetProtectedSourceUidOk returns a tuple with the ProtectedSourceUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedSourceUid

`func (o *ProtectedSourceSummary) SetProtectedSourceUid(v UniversalId)`

SetProtectedSourceUid sets ProtectedSourceUid field to given value.

### HasProtectedSourceUid

`func (o *ProtectedSourceSummary) HasProtectedSourceUid() bool`

HasProtectedSourceUid returns a boolean if a field has been set.

### GetProtectionSource

`func (o *ProtectedSourceSummary) GetProtectionSource() ProtectionSource`

GetProtectionSource returns the ProtectionSource field if non-nil, zero value otherwise.

### GetProtectionSourceOk

`func (o *ProtectedSourceSummary) GetProtectionSourceOk() (*ProtectionSource, bool)`

GetProtectionSourceOk returns a tuple with the ProtectionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionSource

`func (o *ProtectedSourceSummary) SetProtectionSource(v ProtectionSource)`

SetProtectionSource sets ProtectionSource field to given value.

### HasProtectionSource

`func (o *ProtectedSourceSummary) HasProtectionSource() bool`

HasProtectionSource returns a boolean if a field has been set.

### GetSourceParameters

`func (o *ProtectedSourceSummary) GetSourceParameters() []SourceSpecialParameter`

GetSourceParameters returns the SourceParameters field if non-nil, zero value otherwise.

### GetSourceParametersOk

`func (o *ProtectedSourceSummary) GetSourceParametersOk() (*[]SourceSpecialParameter, bool)`

GetSourceParametersOk returns a tuple with the SourceParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceParameters

`func (o *ProtectedSourceSummary) SetSourceParameters(v []SourceSpecialParameter)`

SetSourceParameters sets SourceParameters field to given value.

### HasSourceParameters

`func (o *ProtectedSourceSummary) HasSourceParameters() bool`

HasSourceParameters returns a boolean if a field has been set.

### SetSourceParametersNil

`func (o *ProtectedSourceSummary) SetSourceParametersNil(b bool)`

 SetSourceParametersNil sets the value for SourceParameters to be an explicit nil

### UnsetSourceParameters
`func (o *ProtectedSourceSummary) UnsetSourceParameters()`

UnsetSourceParameters ensures that no value is present for SourceParameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


