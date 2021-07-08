# VirtualDiskRecoverTaskState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**RequestError**](RequestError.md) |  | [optional] 
**IsInstantRecoveryFinished** | Pointer to **NullableBool** | Specifies if instant recovery of the virtual disk is complete. | [optional] 
**TaskState** | Pointer to **NullableString** | Specifies the current state of the restore virtual disks task. Specifies the current state of the restore virtual disks task. &#39;kDetachDisksDone&#39; indicates the detached state of disks. &#39;kSetupDisksDone&#39; indicates that disks setup is completed. &#39;kMigrateDisksStarted&#39; indicates that disks are being migrated. &#39;kMigrateDisksDone&#39; indicates that disk migration is completed. &#39;kUnMountDatastoreDone&#39; indicates that disk has unmounted the datastore. | [optional] 
**VirtualDiskRestoreResponse** | Pointer to [**VirtualDiskRestoreResponse**](VirtualDiskRestoreResponse.md) |  | [optional] 

## Methods

### NewVirtualDiskRecoverTaskState

`func NewVirtualDiskRecoverTaskState() *VirtualDiskRecoverTaskState`

NewVirtualDiskRecoverTaskState instantiates a new VirtualDiskRecoverTaskState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualDiskRecoverTaskStateWithDefaults

`func NewVirtualDiskRecoverTaskStateWithDefaults() *VirtualDiskRecoverTaskState`

NewVirtualDiskRecoverTaskStateWithDefaults instantiates a new VirtualDiskRecoverTaskState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *VirtualDiskRecoverTaskState) GetError() RequestError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *VirtualDiskRecoverTaskState) GetErrorOk() (*RequestError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *VirtualDiskRecoverTaskState) SetError(v RequestError)`

SetError sets Error field to given value.

### HasError

`func (o *VirtualDiskRecoverTaskState) HasError() bool`

HasError returns a boolean if a field has been set.

### GetIsInstantRecoveryFinished

`func (o *VirtualDiskRecoverTaskState) GetIsInstantRecoveryFinished() bool`

GetIsInstantRecoveryFinished returns the IsInstantRecoveryFinished field if non-nil, zero value otherwise.

### GetIsInstantRecoveryFinishedOk

`func (o *VirtualDiskRecoverTaskState) GetIsInstantRecoveryFinishedOk() (*bool, bool)`

GetIsInstantRecoveryFinishedOk returns a tuple with the IsInstantRecoveryFinished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInstantRecoveryFinished

`func (o *VirtualDiskRecoverTaskState) SetIsInstantRecoveryFinished(v bool)`

SetIsInstantRecoveryFinished sets IsInstantRecoveryFinished field to given value.

### HasIsInstantRecoveryFinished

`func (o *VirtualDiskRecoverTaskState) HasIsInstantRecoveryFinished() bool`

HasIsInstantRecoveryFinished returns a boolean if a field has been set.

### SetIsInstantRecoveryFinishedNil

`func (o *VirtualDiskRecoverTaskState) SetIsInstantRecoveryFinishedNil(b bool)`

 SetIsInstantRecoveryFinishedNil sets the value for IsInstantRecoveryFinished to be an explicit nil

### UnsetIsInstantRecoveryFinished
`func (o *VirtualDiskRecoverTaskState) UnsetIsInstantRecoveryFinished()`

UnsetIsInstantRecoveryFinished ensures that no value is present for IsInstantRecoveryFinished, not even an explicit nil
### GetTaskState

`func (o *VirtualDiskRecoverTaskState) GetTaskState() string`

GetTaskState returns the TaskState field if non-nil, zero value otherwise.

### GetTaskStateOk

`func (o *VirtualDiskRecoverTaskState) GetTaskStateOk() (*string, bool)`

GetTaskStateOk returns a tuple with the TaskState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskState

`func (o *VirtualDiskRecoverTaskState) SetTaskState(v string)`

SetTaskState sets TaskState field to given value.

### HasTaskState

`func (o *VirtualDiskRecoverTaskState) HasTaskState() bool`

HasTaskState returns a boolean if a field has been set.

### SetTaskStateNil

`func (o *VirtualDiskRecoverTaskState) SetTaskStateNil(b bool)`

 SetTaskStateNil sets the value for TaskState to be an explicit nil

### UnsetTaskState
`func (o *VirtualDiskRecoverTaskState) UnsetTaskState()`

UnsetTaskState ensures that no value is present for TaskState, not even an explicit nil
### GetVirtualDiskRestoreResponse

`func (o *VirtualDiskRecoverTaskState) GetVirtualDiskRestoreResponse() VirtualDiskRestoreResponse`

GetVirtualDiskRestoreResponse returns the VirtualDiskRestoreResponse field if non-nil, zero value otherwise.

### GetVirtualDiskRestoreResponseOk

`func (o *VirtualDiskRecoverTaskState) GetVirtualDiskRestoreResponseOk() (*VirtualDiskRestoreResponse, bool)`

GetVirtualDiskRestoreResponseOk returns a tuple with the VirtualDiskRestoreResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDiskRestoreResponse

`func (o *VirtualDiskRecoverTaskState) SetVirtualDiskRestoreResponse(v VirtualDiskRestoreResponse)`

SetVirtualDiskRestoreResponse sets VirtualDiskRestoreResponse field to given value.

### HasVirtualDiskRestoreResponse

`func (o *VirtualDiskRecoverTaskState) HasVirtualDiskRestoreResponse() bool`

HasVirtualDiskRestoreResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


