# ProtectionRunInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupRun** | Pointer to [**BackupRun**](BackupRun.md) |  | [optional] 
**CopyRun** | Pointer to [**[]CopyRun**](CopyRun.md) | Array of Copy Run Tasks.  Specifies details about the Copy tasks of this Job Run. A Copy task copies the captured snapshots to an external target or a Remote Cohesity Cluster. | [optional] 
**JobId** | Pointer to **NullableInt64** | Specifies the id of the Protection Job that was run. | [optional] 
**JobName** | Pointer to **NullableString** | Specifies the name of the Protection Job name that was run. | [optional] 
**JobUid** | Pointer to [**NullableUniversalId**](UniversalId.md) | Specifies the globally unique id of the Protection Job that was run. | [optional] 
**ProtectionShellInfo** | Pointer to [**ProtectionShellInfo**](ProtectionShellInfo.md) |  | [optional] 
**ViewBoxId** | Pointer to **NullableInt64** | Specifies the Storage Domain (View Box) to store the backed up data. Specify the id of the Storage Domain (View Box). | [optional] 

## Methods

### NewProtectionRunInstance

`func NewProtectionRunInstance() *ProtectionRunInstance`

NewProtectionRunInstance instantiates a new ProtectionRunInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionRunInstanceWithDefaults

`func NewProtectionRunInstanceWithDefaults() *ProtectionRunInstance`

NewProtectionRunInstanceWithDefaults instantiates a new ProtectionRunInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupRun

`func (o *ProtectionRunInstance) GetBackupRun() BackupRun`

GetBackupRun returns the BackupRun field if non-nil, zero value otherwise.

### GetBackupRunOk

`func (o *ProtectionRunInstance) GetBackupRunOk() (*BackupRun, bool)`

GetBackupRunOk returns a tuple with the BackupRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRun

`func (o *ProtectionRunInstance) SetBackupRun(v BackupRun)`

SetBackupRun sets BackupRun field to given value.

### HasBackupRun

`func (o *ProtectionRunInstance) HasBackupRun() bool`

HasBackupRun returns a boolean if a field has been set.

### GetCopyRun

`func (o *ProtectionRunInstance) GetCopyRun() []CopyRun`

GetCopyRun returns the CopyRun field if non-nil, zero value otherwise.

### GetCopyRunOk

`func (o *ProtectionRunInstance) GetCopyRunOk() (*[]CopyRun, bool)`

GetCopyRunOk returns a tuple with the CopyRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyRun

`func (o *ProtectionRunInstance) SetCopyRun(v []CopyRun)`

SetCopyRun sets CopyRun field to given value.

### HasCopyRun

`func (o *ProtectionRunInstance) HasCopyRun() bool`

HasCopyRun returns a boolean if a field has been set.

### SetCopyRunNil

`func (o *ProtectionRunInstance) SetCopyRunNil(b bool)`

 SetCopyRunNil sets the value for CopyRun to be an explicit nil

### UnsetCopyRun
`func (o *ProtectionRunInstance) UnsetCopyRun()`

UnsetCopyRun ensures that no value is present for CopyRun, not even an explicit nil
### GetJobId

`func (o *ProtectionRunInstance) GetJobId() int64`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *ProtectionRunInstance) GetJobIdOk() (*int64, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *ProtectionRunInstance) SetJobId(v int64)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *ProtectionRunInstance) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### SetJobIdNil

`func (o *ProtectionRunInstance) SetJobIdNil(b bool)`

 SetJobIdNil sets the value for JobId to be an explicit nil

### UnsetJobId
`func (o *ProtectionRunInstance) UnsetJobId()`

UnsetJobId ensures that no value is present for JobId, not even an explicit nil
### GetJobName

`func (o *ProtectionRunInstance) GetJobName() string`

GetJobName returns the JobName field if non-nil, zero value otherwise.

### GetJobNameOk

`func (o *ProtectionRunInstance) GetJobNameOk() (*string, bool)`

GetJobNameOk returns a tuple with the JobName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobName

`func (o *ProtectionRunInstance) SetJobName(v string)`

SetJobName sets JobName field to given value.

### HasJobName

`func (o *ProtectionRunInstance) HasJobName() bool`

HasJobName returns a boolean if a field has been set.

### SetJobNameNil

`func (o *ProtectionRunInstance) SetJobNameNil(b bool)`

 SetJobNameNil sets the value for JobName to be an explicit nil

### UnsetJobName
`func (o *ProtectionRunInstance) UnsetJobName()`

UnsetJobName ensures that no value is present for JobName, not even an explicit nil
### GetJobUid

`func (o *ProtectionRunInstance) GetJobUid() UniversalId`

GetJobUid returns the JobUid field if non-nil, zero value otherwise.

### GetJobUidOk

`func (o *ProtectionRunInstance) GetJobUidOk() (*UniversalId, bool)`

GetJobUidOk returns a tuple with the JobUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobUid

`func (o *ProtectionRunInstance) SetJobUid(v UniversalId)`

SetJobUid sets JobUid field to given value.

### HasJobUid

`func (o *ProtectionRunInstance) HasJobUid() bool`

HasJobUid returns a boolean if a field has been set.

### SetJobUidNil

`func (o *ProtectionRunInstance) SetJobUidNil(b bool)`

 SetJobUidNil sets the value for JobUid to be an explicit nil

### UnsetJobUid
`func (o *ProtectionRunInstance) UnsetJobUid()`

UnsetJobUid ensures that no value is present for JobUid, not even an explicit nil
### GetProtectionShellInfo

`func (o *ProtectionRunInstance) GetProtectionShellInfo() ProtectionShellInfo`

GetProtectionShellInfo returns the ProtectionShellInfo field if non-nil, zero value otherwise.

### GetProtectionShellInfoOk

`func (o *ProtectionRunInstance) GetProtectionShellInfoOk() (*ProtectionShellInfo, bool)`

GetProtectionShellInfoOk returns a tuple with the ProtectionShellInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionShellInfo

`func (o *ProtectionRunInstance) SetProtectionShellInfo(v ProtectionShellInfo)`

SetProtectionShellInfo sets ProtectionShellInfo field to given value.

### HasProtectionShellInfo

`func (o *ProtectionRunInstance) HasProtectionShellInfo() bool`

HasProtectionShellInfo returns a boolean if a field has been set.

### GetViewBoxId

`func (o *ProtectionRunInstance) GetViewBoxId() int64`

GetViewBoxId returns the ViewBoxId field if non-nil, zero value otherwise.

### GetViewBoxIdOk

`func (o *ProtectionRunInstance) GetViewBoxIdOk() (*int64, bool)`

GetViewBoxIdOk returns a tuple with the ViewBoxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewBoxId

`func (o *ProtectionRunInstance) SetViewBoxId(v int64)`

SetViewBoxId sets ViewBoxId field to given value.

### HasViewBoxId

`func (o *ProtectionRunInstance) HasViewBoxId() bool`

HasViewBoxId returns a boolean if a field has been set.

### SetViewBoxIdNil

`func (o *ProtectionRunInstance) SetViewBoxIdNil(b bool)`

 SetViewBoxIdNil sets the value for ViewBoxId to be an explicit nil

### UnsetViewBoxId
`func (o *ProtectionRunInstance) UnsetViewBoxId()`

UnsetViewBoxId ensures that no value is present for ViewBoxId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


