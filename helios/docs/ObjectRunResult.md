# ObjectRunResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to [**ObjectSummary**](ObjectSummary.md) |  | [optional] 
**LocalSnapshotInfo** | Pointer to [**BackupRun**](BackupRun.md) |  | [optional] 
**OriginalBackupInfo** | Pointer to [**BackupRun**](BackupRun.md) |  | [optional] 
**ReplicationInfo** | Pointer to [**ReplicationRun**](ReplicationRun.md) |  | [optional] 
**ArchivalInfo** | Pointer to [**ArchivalRun**](ArchivalRun.md) |  | [optional] 
**CloudSpinInfo** | Pointer to [**CloudSpinRun**](CloudSpinRun.md) |  | [optional] 
**OnLegalHold** | Pointer to **NullableBool** | Specifies if object&#39;s snapshot is on legal hold. | [optional] 

## Methods

### NewObjectRunResult

`func NewObjectRunResult() *ObjectRunResult`

NewObjectRunResult instantiates a new ObjectRunResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectRunResultWithDefaults

`func NewObjectRunResultWithDefaults() *ObjectRunResult`

NewObjectRunResultWithDefaults instantiates a new ObjectRunResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *ObjectRunResult) GetObject() ObjectSummary`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ObjectRunResult) GetObjectOk() (*ObjectSummary, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ObjectRunResult) SetObject(v ObjectSummary)`

SetObject sets Object field to given value.

### HasObject

`func (o *ObjectRunResult) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetLocalSnapshotInfo

`func (o *ObjectRunResult) GetLocalSnapshotInfo() BackupRun`

GetLocalSnapshotInfo returns the LocalSnapshotInfo field if non-nil, zero value otherwise.

### GetLocalSnapshotInfoOk

`func (o *ObjectRunResult) GetLocalSnapshotInfoOk() (*BackupRun, bool)`

GetLocalSnapshotInfoOk returns a tuple with the LocalSnapshotInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSnapshotInfo

`func (o *ObjectRunResult) SetLocalSnapshotInfo(v BackupRun)`

SetLocalSnapshotInfo sets LocalSnapshotInfo field to given value.

### HasLocalSnapshotInfo

`func (o *ObjectRunResult) HasLocalSnapshotInfo() bool`

HasLocalSnapshotInfo returns a boolean if a field has been set.

### GetOriginalBackupInfo

`func (o *ObjectRunResult) GetOriginalBackupInfo() BackupRun`

GetOriginalBackupInfo returns the OriginalBackupInfo field if non-nil, zero value otherwise.

### GetOriginalBackupInfoOk

`func (o *ObjectRunResult) GetOriginalBackupInfoOk() (*BackupRun, bool)`

GetOriginalBackupInfoOk returns a tuple with the OriginalBackupInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalBackupInfo

`func (o *ObjectRunResult) SetOriginalBackupInfo(v BackupRun)`

SetOriginalBackupInfo sets OriginalBackupInfo field to given value.

### HasOriginalBackupInfo

`func (o *ObjectRunResult) HasOriginalBackupInfo() bool`

HasOriginalBackupInfo returns a boolean if a field has been set.

### GetReplicationInfo

`func (o *ObjectRunResult) GetReplicationInfo() ReplicationRun`

GetReplicationInfo returns the ReplicationInfo field if non-nil, zero value otherwise.

### GetReplicationInfoOk

`func (o *ObjectRunResult) GetReplicationInfoOk() (*ReplicationRun, bool)`

GetReplicationInfoOk returns a tuple with the ReplicationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationInfo

`func (o *ObjectRunResult) SetReplicationInfo(v ReplicationRun)`

SetReplicationInfo sets ReplicationInfo field to given value.

### HasReplicationInfo

`func (o *ObjectRunResult) HasReplicationInfo() bool`

HasReplicationInfo returns a boolean if a field has been set.

### GetArchivalInfo

`func (o *ObjectRunResult) GetArchivalInfo() ArchivalRun`

GetArchivalInfo returns the ArchivalInfo field if non-nil, zero value otherwise.

### GetArchivalInfoOk

`func (o *ObjectRunResult) GetArchivalInfoOk() (*ArchivalRun, bool)`

GetArchivalInfoOk returns a tuple with the ArchivalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivalInfo

`func (o *ObjectRunResult) SetArchivalInfo(v ArchivalRun)`

SetArchivalInfo sets ArchivalInfo field to given value.

### HasArchivalInfo

`func (o *ObjectRunResult) HasArchivalInfo() bool`

HasArchivalInfo returns a boolean if a field has been set.

### GetCloudSpinInfo

`func (o *ObjectRunResult) GetCloudSpinInfo() CloudSpinRun`

GetCloudSpinInfo returns the CloudSpinInfo field if non-nil, zero value otherwise.

### GetCloudSpinInfoOk

`func (o *ObjectRunResult) GetCloudSpinInfoOk() (*CloudSpinRun, bool)`

GetCloudSpinInfoOk returns a tuple with the CloudSpinInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudSpinInfo

`func (o *ObjectRunResult) SetCloudSpinInfo(v CloudSpinRun)`

SetCloudSpinInfo sets CloudSpinInfo field to given value.

### HasCloudSpinInfo

`func (o *ObjectRunResult) HasCloudSpinInfo() bool`

HasCloudSpinInfo returns a boolean if a field has been set.

### GetOnLegalHold

`func (o *ObjectRunResult) GetOnLegalHold() bool`

GetOnLegalHold returns the OnLegalHold field if non-nil, zero value otherwise.

### GetOnLegalHoldOk

`func (o *ObjectRunResult) GetOnLegalHoldOk() (*bool, bool)`

GetOnLegalHoldOk returns a tuple with the OnLegalHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnLegalHold

`func (o *ObjectRunResult) SetOnLegalHold(v bool)`

SetOnLegalHold sets OnLegalHold field to given value.

### HasOnLegalHold

`func (o *ObjectRunResult) HasOnLegalHold() bool`

HasOnLegalHold returns a boolean if a field has been set.

### SetOnLegalHoldNil

`func (o *ObjectRunResult) SetOnLegalHoldNil(b bool)`

 SetOnLegalHoldNil sets the value for OnLegalHold to be an explicit nil

### UnsetOnLegalHold
`func (o *ObjectRunResult) UnsetOnLegalHold()`

UnsetOnLegalHold ensures that no value is present for OnLegalHold, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


