# RecoverHyperVParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownloadFileAndFolderParams** | Pointer to [**NullableRecoverHyperVParamsDownloadFileAndFolderParams**](RecoverHyperVParamsDownloadFileAndFolderParams.md) |  | [optional] 
**MountVolumeParams** | Pointer to [**NullableRecoverHyperVParamsMountVolumeParams**](RecoverHyperVParamsMountVolumeParams.md) |  | [optional] 
**Objects** | Pointer to [**[]CommonRecoverObjectSnapshotParams**](CommonRecoverObjectSnapshotParams.md) | Specifies the list of recover Object parameters. This property is mandatory for all recovery action types except recover vms. While recovering VMs, a user can specify snapshots of VM&#39;s or a Protection Group Run details to recover all the VM&#39;s that are backed up by that Run. For recovering files, specifies the object contains the file to recover. | [optional] 
**RecoverFileAndFolderParams** | Pointer to [**NullableRecoverHyperVParamsRecoverFileAndFolderParams**](RecoverHyperVParamsRecoverFileAndFolderParams.md) |  | [optional] 
**RecoverVmParams** | Pointer to [**NullableRecoverHyperVParamsRecoverVmParams**](RecoverHyperVParamsRecoverVmParams.md) |  | [optional] 
**RecoveryAction** | **string** | Specifies the type of recovery action to be performed. | 

## Methods

### NewRecoverHyperVParams

`func NewRecoverHyperVParams(recoveryAction string, ) *RecoverHyperVParams`

NewRecoverHyperVParams instantiates a new RecoverHyperVParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverHyperVParamsWithDefaults

`func NewRecoverHyperVParamsWithDefaults() *RecoverHyperVParams`

NewRecoverHyperVParamsWithDefaults instantiates a new RecoverHyperVParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownloadFileAndFolderParams

`func (o *RecoverHyperVParams) GetDownloadFileAndFolderParams() RecoverHyperVParamsDownloadFileAndFolderParams`

GetDownloadFileAndFolderParams returns the DownloadFileAndFolderParams field if non-nil, zero value otherwise.

### GetDownloadFileAndFolderParamsOk

`func (o *RecoverHyperVParams) GetDownloadFileAndFolderParamsOk() (*RecoverHyperVParamsDownloadFileAndFolderParams, bool)`

GetDownloadFileAndFolderParamsOk returns a tuple with the DownloadFileAndFolderParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadFileAndFolderParams

`func (o *RecoverHyperVParams) SetDownloadFileAndFolderParams(v RecoverHyperVParamsDownloadFileAndFolderParams)`

SetDownloadFileAndFolderParams sets DownloadFileAndFolderParams field to given value.

### HasDownloadFileAndFolderParams

`func (o *RecoverHyperVParams) HasDownloadFileAndFolderParams() bool`

HasDownloadFileAndFolderParams returns a boolean if a field has been set.

### SetDownloadFileAndFolderParamsNil

`func (o *RecoverHyperVParams) SetDownloadFileAndFolderParamsNil(b bool)`

 SetDownloadFileAndFolderParamsNil sets the value for DownloadFileAndFolderParams to be an explicit nil

### UnsetDownloadFileAndFolderParams
`func (o *RecoverHyperVParams) UnsetDownloadFileAndFolderParams()`

UnsetDownloadFileAndFolderParams ensures that no value is present for DownloadFileAndFolderParams, not even an explicit nil
### GetMountVolumeParams

`func (o *RecoverHyperVParams) GetMountVolumeParams() RecoverHyperVParamsMountVolumeParams`

GetMountVolumeParams returns the MountVolumeParams field if non-nil, zero value otherwise.

### GetMountVolumeParamsOk

`func (o *RecoverHyperVParams) GetMountVolumeParamsOk() (*RecoverHyperVParamsMountVolumeParams, bool)`

GetMountVolumeParamsOk returns a tuple with the MountVolumeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountVolumeParams

`func (o *RecoverHyperVParams) SetMountVolumeParams(v RecoverHyperVParamsMountVolumeParams)`

SetMountVolumeParams sets MountVolumeParams field to given value.

### HasMountVolumeParams

`func (o *RecoverHyperVParams) HasMountVolumeParams() bool`

HasMountVolumeParams returns a boolean if a field has been set.

### SetMountVolumeParamsNil

`func (o *RecoverHyperVParams) SetMountVolumeParamsNil(b bool)`

 SetMountVolumeParamsNil sets the value for MountVolumeParams to be an explicit nil

### UnsetMountVolumeParams
`func (o *RecoverHyperVParams) UnsetMountVolumeParams()`

UnsetMountVolumeParams ensures that no value is present for MountVolumeParams, not even an explicit nil
### GetObjects

`func (o *RecoverHyperVParams) GetObjects() []CommonRecoverObjectSnapshotParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *RecoverHyperVParams) GetObjectsOk() (*[]CommonRecoverObjectSnapshotParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *RecoverHyperVParams) SetObjects(v []CommonRecoverObjectSnapshotParams)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *RecoverHyperVParams) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### SetObjectsNil

`func (o *RecoverHyperVParams) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *RecoverHyperVParams) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetRecoverFileAndFolderParams

`func (o *RecoverHyperVParams) GetRecoverFileAndFolderParams() RecoverHyperVParamsRecoverFileAndFolderParams`

GetRecoverFileAndFolderParams returns the RecoverFileAndFolderParams field if non-nil, zero value otherwise.

### GetRecoverFileAndFolderParamsOk

`func (o *RecoverHyperVParams) GetRecoverFileAndFolderParamsOk() (*RecoverHyperVParamsRecoverFileAndFolderParams, bool)`

GetRecoverFileAndFolderParamsOk returns a tuple with the RecoverFileAndFolderParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverFileAndFolderParams

`func (o *RecoverHyperVParams) SetRecoverFileAndFolderParams(v RecoverHyperVParamsRecoverFileAndFolderParams)`

SetRecoverFileAndFolderParams sets RecoverFileAndFolderParams field to given value.

### HasRecoverFileAndFolderParams

`func (o *RecoverHyperVParams) HasRecoverFileAndFolderParams() bool`

HasRecoverFileAndFolderParams returns a boolean if a field has been set.

### SetRecoverFileAndFolderParamsNil

`func (o *RecoverHyperVParams) SetRecoverFileAndFolderParamsNil(b bool)`

 SetRecoverFileAndFolderParamsNil sets the value for RecoverFileAndFolderParams to be an explicit nil

### UnsetRecoverFileAndFolderParams
`func (o *RecoverHyperVParams) UnsetRecoverFileAndFolderParams()`

UnsetRecoverFileAndFolderParams ensures that no value is present for RecoverFileAndFolderParams, not even an explicit nil
### GetRecoverVmParams

`func (o *RecoverHyperVParams) GetRecoverVmParams() RecoverHyperVParamsRecoverVmParams`

GetRecoverVmParams returns the RecoverVmParams field if non-nil, zero value otherwise.

### GetRecoverVmParamsOk

`func (o *RecoverHyperVParams) GetRecoverVmParamsOk() (*RecoverHyperVParamsRecoverVmParams, bool)`

GetRecoverVmParamsOk returns a tuple with the RecoverVmParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverVmParams

`func (o *RecoverHyperVParams) SetRecoverVmParams(v RecoverHyperVParamsRecoverVmParams)`

SetRecoverVmParams sets RecoverVmParams field to given value.

### HasRecoverVmParams

`func (o *RecoverHyperVParams) HasRecoverVmParams() bool`

HasRecoverVmParams returns a boolean if a field has been set.

### SetRecoverVmParamsNil

`func (o *RecoverHyperVParams) SetRecoverVmParamsNil(b bool)`

 SetRecoverVmParamsNil sets the value for RecoverVmParams to be an explicit nil

### UnsetRecoverVmParams
`func (o *RecoverHyperVParams) UnsetRecoverVmParams()`

UnsetRecoverVmParams ensures that no value is present for RecoverVmParams, not even an explicit nil
### GetRecoveryAction

`func (o *RecoverHyperVParams) GetRecoveryAction() string`

GetRecoveryAction returns the RecoveryAction field if non-nil, zero value otherwise.

### GetRecoveryActionOk

`func (o *RecoverHyperVParams) GetRecoveryActionOk() (*string, bool)`

GetRecoveryActionOk returns a tuple with the RecoveryAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryAction

`func (o *RecoverHyperVParams) SetRecoveryAction(v string)`

SetRecoveryAction sets RecoveryAction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


