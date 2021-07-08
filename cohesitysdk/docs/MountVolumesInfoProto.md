# MountVolumesInfoProto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CleanupError** | Pointer to [**ErrorProto**](ErrorProto.md) |  | [optional] 
**Error** | Pointer to [**ErrorProto**](ErrorProto.md) |  | [optional] 
**Finished** | Pointer to **NullableBool** | This will be set to true if the task is complete on the slave. | [optional] 
**MountVolumeResultVec** | Pointer to [**[]MountVolumeResult**](MountVolumeResult.md) | Contains the result information of the mounted volumes. | [optional] 
**RestoreDisksTaskInfoProto** | Pointer to [**SetupRestoreDiskTaskInfoProto**](SetupRestoreDiskTaskInfoProto.md) |  | [optional] 
**SlaveTaskStartTimeUsecs** | Pointer to **NullableInt64** | This is the timestamp at which the slave task started. | [optional] 
**Type** | Pointer to **NullableInt32** | The type of environment this mount volumes info pertains to. | [optional] 
**VmOnlineDisksError** | Pointer to [**ErrorProto**](ErrorProto.md) |  | [optional] 

## Methods

### NewMountVolumesInfoProto

`func NewMountVolumesInfoProto() *MountVolumesInfoProto`

NewMountVolumesInfoProto instantiates a new MountVolumesInfoProto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMountVolumesInfoProtoWithDefaults

`func NewMountVolumesInfoProtoWithDefaults() *MountVolumesInfoProto`

NewMountVolumesInfoProtoWithDefaults instantiates a new MountVolumesInfoProto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCleanupError

`func (o *MountVolumesInfoProto) GetCleanupError() ErrorProto`

GetCleanupError returns the CleanupError field if non-nil, zero value otherwise.

### GetCleanupErrorOk

`func (o *MountVolumesInfoProto) GetCleanupErrorOk() (*ErrorProto, bool)`

GetCleanupErrorOk returns a tuple with the CleanupError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanupError

`func (o *MountVolumesInfoProto) SetCleanupError(v ErrorProto)`

SetCleanupError sets CleanupError field to given value.

### HasCleanupError

`func (o *MountVolumesInfoProto) HasCleanupError() bool`

HasCleanupError returns a boolean if a field has been set.

### GetError

`func (o *MountVolumesInfoProto) GetError() ErrorProto`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MountVolumesInfoProto) GetErrorOk() (*ErrorProto, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MountVolumesInfoProto) SetError(v ErrorProto)`

SetError sets Error field to given value.

### HasError

`func (o *MountVolumesInfoProto) HasError() bool`

HasError returns a boolean if a field has been set.

### GetFinished

`func (o *MountVolumesInfoProto) GetFinished() bool`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *MountVolumesInfoProto) GetFinishedOk() (*bool, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *MountVolumesInfoProto) SetFinished(v bool)`

SetFinished sets Finished field to given value.

### HasFinished

`func (o *MountVolumesInfoProto) HasFinished() bool`

HasFinished returns a boolean if a field has been set.

### SetFinishedNil

`func (o *MountVolumesInfoProto) SetFinishedNil(b bool)`

 SetFinishedNil sets the value for Finished to be an explicit nil

### UnsetFinished
`func (o *MountVolumesInfoProto) UnsetFinished()`

UnsetFinished ensures that no value is present for Finished, not even an explicit nil
### GetMountVolumeResultVec

`func (o *MountVolumesInfoProto) GetMountVolumeResultVec() []MountVolumeResult`

GetMountVolumeResultVec returns the MountVolumeResultVec field if non-nil, zero value otherwise.

### GetMountVolumeResultVecOk

`func (o *MountVolumesInfoProto) GetMountVolumeResultVecOk() (*[]MountVolumeResult, bool)`

GetMountVolumeResultVecOk returns a tuple with the MountVolumeResultVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountVolumeResultVec

`func (o *MountVolumesInfoProto) SetMountVolumeResultVec(v []MountVolumeResult)`

SetMountVolumeResultVec sets MountVolumeResultVec field to given value.

### HasMountVolumeResultVec

`func (o *MountVolumesInfoProto) HasMountVolumeResultVec() bool`

HasMountVolumeResultVec returns a boolean if a field has been set.

### SetMountVolumeResultVecNil

`func (o *MountVolumesInfoProto) SetMountVolumeResultVecNil(b bool)`

 SetMountVolumeResultVecNil sets the value for MountVolumeResultVec to be an explicit nil

### UnsetMountVolumeResultVec
`func (o *MountVolumesInfoProto) UnsetMountVolumeResultVec()`

UnsetMountVolumeResultVec ensures that no value is present for MountVolumeResultVec, not even an explicit nil
### GetRestoreDisksTaskInfoProto

`func (o *MountVolumesInfoProto) GetRestoreDisksTaskInfoProto() SetupRestoreDiskTaskInfoProto`

GetRestoreDisksTaskInfoProto returns the RestoreDisksTaskInfoProto field if non-nil, zero value otherwise.

### GetRestoreDisksTaskInfoProtoOk

`func (o *MountVolumesInfoProto) GetRestoreDisksTaskInfoProtoOk() (*SetupRestoreDiskTaskInfoProto, bool)`

GetRestoreDisksTaskInfoProtoOk returns a tuple with the RestoreDisksTaskInfoProto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreDisksTaskInfoProto

`func (o *MountVolumesInfoProto) SetRestoreDisksTaskInfoProto(v SetupRestoreDiskTaskInfoProto)`

SetRestoreDisksTaskInfoProto sets RestoreDisksTaskInfoProto field to given value.

### HasRestoreDisksTaskInfoProto

`func (o *MountVolumesInfoProto) HasRestoreDisksTaskInfoProto() bool`

HasRestoreDisksTaskInfoProto returns a boolean if a field has been set.

### GetSlaveTaskStartTimeUsecs

`func (o *MountVolumesInfoProto) GetSlaveTaskStartTimeUsecs() int64`

GetSlaveTaskStartTimeUsecs returns the SlaveTaskStartTimeUsecs field if non-nil, zero value otherwise.

### GetSlaveTaskStartTimeUsecsOk

`func (o *MountVolumesInfoProto) GetSlaveTaskStartTimeUsecsOk() (*int64, bool)`

GetSlaveTaskStartTimeUsecsOk returns a tuple with the SlaveTaskStartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaveTaskStartTimeUsecs

`func (o *MountVolumesInfoProto) SetSlaveTaskStartTimeUsecs(v int64)`

SetSlaveTaskStartTimeUsecs sets SlaveTaskStartTimeUsecs field to given value.

### HasSlaveTaskStartTimeUsecs

`func (o *MountVolumesInfoProto) HasSlaveTaskStartTimeUsecs() bool`

HasSlaveTaskStartTimeUsecs returns a boolean if a field has been set.

### SetSlaveTaskStartTimeUsecsNil

`func (o *MountVolumesInfoProto) SetSlaveTaskStartTimeUsecsNil(b bool)`

 SetSlaveTaskStartTimeUsecsNil sets the value for SlaveTaskStartTimeUsecs to be an explicit nil

### UnsetSlaveTaskStartTimeUsecs
`func (o *MountVolumesInfoProto) UnsetSlaveTaskStartTimeUsecs()`

UnsetSlaveTaskStartTimeUsecs ensures that no value is present for SlaveTaskStartTimeUsecs, not even an explicit nil
### GetType

`func (o *MountVolumesInfoProto) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MountVolumesInfoProto) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MountVolumesInfoProto) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *MountVolumesInfoProto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *MountVolumesInfoProto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MountVolumesInfoProto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetVmOnlineDisksError

`func (o *MountVolumesInfoProto) GetVmOnlineDisksError() ErrorProto`

GetVmOnlineDisksError returns the VmOnlineDisksError field if non-nil, zero value otherwise.

### GetVmOnlineDisksErrorOk

`func (o *MountVolumesInfoProto) GetVmOnlineDisksErrorOk() (*ErrorProto, bool)`

GetVmOnlineDisksErrorOk returns a tuple with the VmOnlineDisksError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmOnlineDisksError

`func (o *MountVolumesInfoProto) SetVmOnlineDisksError(v ErrorProto)`

SetVmOnlineDisksError sets VmOnlineDisksError field to given value.

### HasVmOnlineDisksError

`func (o *MountVolumesInfoProto) HasVmOnlineDisksError() bool`

HasVmOnlineDisksError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


