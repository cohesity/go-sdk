# PhysicalBackupSourceParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableSystemBackup** | Pointer to **NullableBool** | Allows Magneto to drive a \&quot;system\&quot; backup using a 3rd-party tool installed on the Agent host. | [optional] 
**FileBackupParams** | Pointer to [**PhysicalFileBackupParams**](PhysicalFileBackupParams.md) |  | [optional] 
**SnapshotParams** | Pointer to [**PhysicalSnapshotParams**](PhysicalSnapshotParams.md) |  | [optional] 
**SourceAppParams** | Pointer to [**SourceAppParams**](SourceAppParams.md) |  | [optional] 
**VolumeGuidVec** | Pointer to **[]string** | If this list is non-empty, then only volumes in this will be protected, otherwise all volumes belonging to the host will be protected. | [optional] 

## Methods

### NewPhysicalBackupSourceParams

`func NewPhysicalBackupSourceParams() *PhysicalBackupSourceParams`

NewPhysicalBackupSourceParams instantiates a new PhysicalBackupSourceParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhysicalBackupSourceParamsWithDefaults

`func NewPhysicalBackupSourceParamsWithDefaults() *PhysicalBackupSourceParams`

NewPhysicalBackupSourceParamsWithDefaults instantiates a new PhysicalBackupSourceParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableSystemBackup

`func (o *PhysicalBackupSourceParams) GetEnableSystemBackup() bool`

GetEnableSystemBackup returns the EnableSystemBackup field if non-nil, zero value otherwise.

### GetEnableSystemBackupOk

`func (o *PhysicalBackupSourceParams) GetEnableSystemBackupOk() (*bool, bool)`

GetEnableSystemBackupOk returns a tuple with the EnableSystemBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSystemBackup

`func (o *PhysicalBackupSourceParams) SetEnableSystemBackup(v bool)`

SetEnableSystemBackup sets EnableSystemBackup field to given value.

### HasEnableSystemBackup

`func (o *PhysicalBackupSourceParams) HasEnableSystemBackup() bool`

HasEnableSystemBackup returns a boolean if a field has been set.

### SetEnableSystemBackupNil

`func (o *PhysicalBackupSourceParams) SetEnableSystemBackupNil(b bool)`

 SetEnableSystemBackupNil sets the value for EnableSystemBackup to be an explicit nil

### UnsetEnableSystemBackup
`func (o *PhysicalBackupSourceParams) UnsetEnableSystemBackup()`

UnsetEnableSystemBackup ensures that no value is present for EnableSystemBackup, not even an explicit nil
### GetFileBackupParams

`func (o *PhysicalBackupSourceParams) GetFileBackupParams() PhysicalFileBackupParams`

GetFileBackupParams returns the FileBackupParams field if non-nil, zero value otherwise.

### GetFileBackupParamsOk

`func (o *PhysicalBackupSourceParams) GetFileBackupParamsOk() (*PhysicalFileBackupParams, bool)`

GetFileBackupParamsOk returns a tuple with the FileBackupParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileBackupParams

`func (o *PhysicalBackupSourceParams) SetFileBackupParams(v PhysicalFileBackupParams)`

SetFileBackupParams sets FileBackupParams field to given value.

### HasFileBackupParams

`func (o *PhysicalBackupSourceParams) HasFileBackupParams() bool`

HasFileBackupParams returns a boolean if a field has been set.

### GetSnapshotParams

`func (o *PhysicalBackupSourceParams) GetSnapshotParams() PhysicalSnapshotParams`

GetSnapshotParams returns the SnapshotParams field if non-nil, zero value otherwise.

### GetSnapshotParamsOk

`func (o *PhysicalBackupSourceParams) GetSnapshotParamsOk() (*PhysicalSnapshotParams, bool)`

GetSnapshotParamsOk returns a tuple with the SnapshotParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotParams

`func (o *PhysicalBackupSourceParams) SetSnapshotParams(v PhysicalSnapshotParams)`

SetSnapshotParams sets SnapshotParams field to given value.

### HasSnapshotParams

`func (o *PhysicalBackupSourceParams) HasSnapshotParams() bool`

HasSnapshotParams returns a boolean if a field has been set.

### GetSourceAppParams

`func (o *PhysicalBackupSourceParams) GetSourceAppParams() SourceAppParams`

GetSourceAppParams returns the SourceAppParams field if non-nil, zero value otherwise.

### GetSourceAppParamsOk

`func (o *PhysicalBackupSourceParams) GetSourceAppParamsOk() (*SourceAppParams, bool)`

GetSourceAppParamsOk returns a tuple with the SourceAppParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAppParams

`func (o *PhysicalBackupSourceParams) SetSourceAppParams(v SourceAppParams)`

SetSourceAppParams sets SourceAppParams field to given value.

### HasSourceAppParams

`func (o *PhysicalBackupSourceParams) HasSourceAppParams() bool`

HasSourceAppParams returns a boolean if a field has been set.

### GetVolumeGuidVec

`func (o *PhysicalBackupSourceParams) GetVolumeGuidVec() []string`

GetVolumeGuidVec returns the VolumeGuidVec field if non-nil, zero value otherwise.

### GetVolumeGuidVecOk

`func (o *PhysicalBackupSourceParams) GetVolumeGuidVecOk() (*[]string, bool)`

GetVolumeGuidVecOk returns a tuple with the VolumeGuidVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeGuidVec

`func (o *PhysicalBackupSourceParams) SetVolumeGuidVec(v []string)`

SetVolumeGuidVec sets VolumeGuidVec field to given value.

### HasVolumeGuidVec

`func (o *PhysicalBackupSourceParams) HasVolumeGuidVec() bool`

HasVolumeGuidVec returns a boolean if a field has been set.

### SetVolumeGuidVecNil

`func (o *PhysicalBackupSourceParams) SetVolumeGuidVecNil(b bool)`

 SetVolumeGuidVecNil sets the value for VolumeGuidVec to be an explicit nil

### UnsetVolumeGuidVec
`func (o *PhysicalBackupSourceParams) UnsetVolumeGuidVec()`

UnsetVolumeGuidVec ensures that no value is present for VolumeGuidVec, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


