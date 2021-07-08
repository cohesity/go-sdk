# JobPolicyProto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupPolicy** | Pointer to [**BackupPolicyProto**](BackupPolicyProto.md) |  | [optional] 
**SnapshotTargetPolicyVec** | Pointer to [**[]SnapshotTargetPolicyProto**](SnapshotTargetPolicyProto.md) | Specifies additional policies that can be used to copy snapshots created by a backup run to different targets (such as a remote replica, tape etc). Each policy also specifies the retention policy that should be applied to the copied snapshots at the respective target. | [optional] 

## Methods

### NewJobPolicyProto

`func NewJobPolicyProto() *JobPolicyProto`

NewJobPolicyProto instantiates a new JobPolicyProto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobPolicyProtoWithDefaults

`func NewJobPolicyProtoWithDefaults() *JobPolicyProto`

NewJobPolicyProtoWithDefaults instantiates a new JobPolicyProto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupPolicy

`func (o *JobPolicyProto) GetBackupPolicy() BackupPolicyProto`

GetBackupPolicy returns the BackupPolicy field if non-nil, zero value otherwise.

### GetBackupPolicyOk

`func (o *JobPolicyProto) GetBackupPolicyOk() (*BackupPolicyProto, bool)`

GetBackupPolicyOk returns a tuple with the BackupPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupPolicy

`func (o *JobPolicyProto) SetBackupPolicy(v BackupPolicyProto)`

SetBackupPolicy sets BackupPolicy field to given value.

### HasBackupPolicy

`func (o *JobPolicyProto) HasBackupPolicy() bool`

HasBackupPolicy returns a boolean if a field has been set.

### GetSnapshotTargetPolicyVec

`func (o *JobPolicyProto) GetSnapshotTargetPolicyVec() []SnapshotTargetPolicyProto`

GetSnapshotTargetPolicyVec returns the SnapshotTargetPolicyVec field if non-nil, zero value otherwise.

### GetSnapshotTargetPolicyVecOk

`func (o *JobPolicyProto) GetSnapshotTargetPolicyVecOk() (*[]SnapshotTargetPolicyProto, bool)`

GetSnapshotTargetPolicyVecOk returns a tuple with the SnapshotTargetPolicyVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotTargetPolicyVec

`func (o *JobPolicyProto) SetSnapshotTargetPolicyVec(v []SnapshotTargetPolicyProto)`

SetSnapshotTargetPolicyVec sets SnapshotTargetPolicyVec field to given value.

### HasSnapshotTargetPolicyVec

`func (o *JobPolicyProto) HasSnapshotTargetPolicyVec() bool`

HasSnapshotTargetPolicyVec returns a boolean if a field has been set.

### SetSnapshotTargetPolicyVecNil

`func (o *JobPolicyProto) SetSnapshotTargetPolicyVecNil(b bool)`

 SetSnapshotTargetPolicyVecNil sets the value for SnapshotTargetPolicyVec to be an explicit nil

### UnsetSnapshotTargetPolicyVec
`func (o *JobPolicyProto) UnsetSnapshotTargetPolicyVec()`

UnsetSnapshotTargetPolicyVec ensures that no value is present for SnapshotTargetPolicyVec, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


