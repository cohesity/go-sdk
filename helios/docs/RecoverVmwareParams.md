# RecoverVmwareParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | Pointer to [**[]RecoverVmwareSnapshotParams**](RecoverVmwareSnapshotParams.md) | Specifies the list of recover Object parameters. This property is mandatory for all recovery action types except recover vms. While recovering VMs, a user can specify snapshots of VM&#39;s or a Protection Group Run details to recover all the VM&#39;s that are backed up by that Run. For recovering files, specifies the object contains the file to recover. | [optional] 
**RecoveryAction** | **string** | Specifies the type of recovery action to be performed. | 
**RecoverVmParams** | Pointer to [**NullableRecoverVmwareVmParams**](RecoverVmwareVmParams.md) | Specifies the parameters to recover VMware VM. | [optional] 
**RecoverVmDiskParams** | Pointer to [**NullableRecoverVmwareDiskParams**](RecoverVmwareDiskParams.md) | Specifies the parameters to recover VMware Disks. | [optional] 
**MountVolumeParams** | Pointer to [**NullableMountVmwareVolumeParams**](MountVmwareVolumeParams.md) | Specifies the parameters to mount VMware Volumes. | [optional] 
**RecoverVAppParams** | Pointer to [**NullableRecoverVmwareVAppParams**](RecoverVmwareVAppParams.md) | Specifies the parameters to recover a VMware vApp. | [optional] 
**RecoverVAppTemplateParams** | Pointer to [**NullableRecoverVmwareVAppTemplateParams**](RecoverVmwareVAppTemplateParams.md) | Specifies the parameters to recover a VMware vApp template. | [optional] 
**RecoverFileAndFolderParams** | Pointer to [**NullableRecoverVMwareFileAndFolderParams**](RecoverVMwareFileAndFolderParams.md) | Specifies the parameters to recover files and folders. | [optional] 
**DownloadFileAndFolderParams** | Pointer to [**NullableCommonDownloadFileAndFolderParams**](CommonDownloadFileAndFolderParams.md) | Specifies the parameters to download files and folders. | [optional] 

## Methods

### NewRecoverVmwareParams

`func NewRecoverVmwareParams(recoveryAction string, ) *RecoverVmwareParams`

NewRecoverVmwareParams instantiates a new RecoverVmwareParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverVmwareParamsWithDefaults

`func NewRecoverVmwareParamsWithDefaults() *RecoverVmwareParams`

NewRecoverVmwareParamsWithDefaults instantiates a new RecoverVmwareParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjects

`func (o *RecoverVmwareParams) GetObjects() []RecoverVmwareSnapshotParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *RecoverVmwareParams) GetObjectsOk() (*[]RecoverVmwareSnapshotParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *RecoverVmwareParams) SetObjects(v []RecoverVmwareSnapshotParams)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *RecoverVmwareParams) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### SetObjectsNil

`func (o *RecoverVmwareParams) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *RecoverVmwareParams) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetRecoveryAction

`func (o *RecoverVmwareParams) GetRecoveryAction() string`

GetRecoveryAction returns the RecoveryAction field if non-nil, zero value otherwise.

### GetRecoveryActionOk

`func (o *RecoverVmwareParams) GetRecoveryActionOk() (*string, bool)`

GetRecoveryActionOk returns a tuple with the RecoveryAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryAction

`func (o *RecoverVmwareParams) SetRecoveryAction(v string)`

SetRecoveryAction sets RecoveryAction field to given value.


### GetRecoverVmParams

`func (o *RecoverVmwareParams) GetRecoverVmParams() RecoverVmwareVmParams`

GetRecoverVmParams returns the RecoverVmParams field if non-nil, zero value otherwise.

### GetRecoverVmParamsOk

`func (o *RecoverVmwareParams) GetRecoverVmParamsOk() (*RecoverVmwareVmParams, bool)`

GetRecoverVmParamsOk returns a tuple with the RecoverVmParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverVmParams

`func (o *RecoverVmwareParams) SetRecoverVmParams(v RecoverVmwareVmParams)`

SetRecoverVmParams sets RecoverVmParams field to given value.

### HasRecoverVmParams

`func (o *RecoverVmwareParams) HasRecoverVmParams() bool`

HasRecoverVmParams returns a boolean if a field has been set.

### SetRecoverVmParamsNil

`func (o *RecoverVmwareParams) SetRecoverVmParamsNil(b bool)`

 SetRecoverVmParamsNil sets the value for RecoverVmParams to be an explicit nil

### UnsetRecoverVmParams
`func (o *RecoverVmwareParams) UnsetRecoverVmParams()`

UnsetRecoverVmParams ensures that no value is present for RecoverVmParams, not even an explicit nil
### GetRecoverVmDiskParams

`func (o *RecoverVmwareParams) GetRecoverVmDiskParams() RecoverVmwareDiskParams`

GetRecoverVmDiskParams returns the RecoverVmDiskParams field if non-nil, zero value otherwise.

### GetRecoverVmDiskParamsOk

`func (o *RecoverVmwareParams) GetRecoverVmDiskParamsOk() (*RecoverVmwareDiskParams, bool)`

GetRecoverVmDiskParamsOk returns a tuple with the RecoverVmDiskParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverVmDiskParams

`func (o *RecoverVmwareParams) SetRecoverVmDiskParams(v RecoverVmwareDiskParams)`

SetRecoverVmDiskParams sets RecoverVmDiskParams field to given value.

### HasRecoverVmDiskParams

`func (o *RecoverVmwareParams) HasRecoverVmDiskParams() bool`

HasRecoverVmDiskParams returns a boolean if a field has been set.

### SetRecoverVmDiskParamsNil

`func (o *RecoverVmwareParams) SetRecoverVmDiskParamsNil(b bool)`

 SetRecoverVmDiskParamsNil sets the value for RecoverVmDiskParams to be an explicit nil

### UnsetRecoverVmDiskParams
`func (o *RecoverVmwareParams) UnsetRecoverVmDiskParams()`

UnsetRecoverVmDiskParams ensures that no value is present for RecoverVmDiskParams, not even an explicit nil
### GetMountVolumeParams

`func (o *RecoverVmwareParams) GetMountVolumeParams() MountVmwareVolumeParams`

GetMountVolumeParams returns the MountVolumeParams field if non-nil, zero value otherwise.

### GetMountVolumeParamsOk

`func (o *RecoverVmwareParams) GetMountVolumeParamsOk() (*MountVmwareVolumeParams, bool)`

GetMountVolumeParamsOk returns a tuple with the MountVolumeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountVolumeParams

`func (o *RecoverVmwareParams) SetMountVolumeParams(v MountVmwareVolumeParams)`

SetMountVolumeParams sets MountVolumeParams field to given value.

### HasMountVolumeParams

`func (o *RecoverVmwareParams) HasMountVolumeParams() bool`

HasMountVolumeParams returns a boolean if a field has been set.

### SetMountVolumeParamsNil

`func (o *RecoverVmwareParams) SetMountVolumeParamsNil(b bool)`

 SetMountVolumeParamsNil sets the value for MountVolumeParams to be an explicit nil

### UnsetMountVolumeParams
`func (o *RecoverVmwareParams) UnsetMountVolumeParams()`

UnsetMountVolumeParams ensures that no value is present for MountVolumeParams, not even an explicit nil
### GetRecoverVAppParams

`func (o *RecoverVmwareParams) GetRecoverVAppParams() RecoverVmwareVAppParams`

GetRecoverVAppParams returns the RecoverVAppParams field if non-nil, zero value otherwise.

### GetRecoverVAppParamsOk

`func (o *RecoverVmwareParams) GetRecoverVAppParamsOk() (*RecoverVmwareVAppParams, bool)`

GetRecoverVAppParamsOk returns a tuple with the RecoverVAppParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverVAppParams

`func (o *RecoverVmwareParams) SetRecoverVAppParams(v RecoverVmwareVAppParams)`

SetRecoverVAppParams sets RecoverVAppParams field to given value.

### HasRecoverVAppParams

`func (o *RecoverVmwareParams) HasRecoverVAppParams() bool`

HasRecoverVAppParams returns a boolean if a field has been set.

### SetRecoverVAppParamsNil

`func (o *RecoverVmwareParams) SetRecoverVAppParamsNil(b bool)`

 SetRecoverVAppParamsNil sets the value for RecoverVAppParams to be an explicit nil

### UnsetRecoverVAppParams
`func (o *RecoverVmwareParams) UnsetRecoverVAppParams()`

UnsetRecoverVAppParams ensures that no value is present for RecoverVAppParams, not even an explicit nil
### GetRecoverVAppTemplateParams

`func (o *RecoverVmwareParams) GetRecoverVAppTemplateParams() RecoverVmwareVAppTemplateParams`

GetRecoverVAppTemplateParams returns the RecoverVAppTemplateParams field if non-nil, zero value otherwise.

### GetRecoverVAppTemplateParamsOk

`func (o *RecoverVmwareParams) GetRecoverVAppTemplateParamsOk() (*RecoverVmwareVAppTemplateParams, bool)`

GetRecoverVAppTemplateParamsOk returns a tuple with the RecoverVAppTemplateParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverVAppTemplateParams

`func (o *RecoverVmwareParams) SetRecoverVAppTemplateParams(v RecoverVmwareVAppTemplateParams)`

SetRecoverVAppTemplateParams sets RecoverVAppTemplateParams field to given value.

### HasRecoverVAppTemplateParams

`func (o *RecoverVmwareParams) HasRecoverVAppTemplateParams() bool`

HasRecoverVAppTemplateParams returns a boolean if a field has been set.

### SetRecoverVAppTemplateParamsNil

`func (o *RecoverVmwareParams) SetRecoverVAppTemplateParamsNil(b bool)`

 SetRecoverVAppTemplateParamsNil sets the value for RecoverVAppTemplateParams to be an explicit nil

### UnsetRecoverVAppTemplateParams
`func (o *RecoverVmwareParams) UnsetRecoverVAppTemplateParams()`

UnsetRecoverVAppTemplateParams ensures that no value is present for RecoverVAppTemplateParams, not even an explicit nil
### GetRecoverFileAndFolderParams

`func (o *RecoverVmwareParams) GetRecoverFileAndFolderParams() RecoverVMwareFileAndFolderParams`

GetRecoverFileAndFolderParams returns the RecoverFileAndFolderParams field if non-nil, zero value otherwise.

### GetRecoverFileAndFolderParamsOk

`func (o *RecoverVmwareParams) GetRecoverFileAndFolderParamsOk() (*RecoverVMwareFileAndFolderParams, bool)`

GetRecoverFileAndFolderParamsOk returns a tuple with the RecoverFileAndFolderParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverFileAndFolderParams

`func (o *RecoverVmwareParams) SetRecoverFileAndFolderParams(v RecoverVMwareFileAndFolderParams)`

SetRecoverFileAndFolderParams sets RecoverFileAndFolderParams field to given value.

### HasRecoverFileAndFolderParams

`func (o *RecoverVmwareParams) HasRecoverFileAndFolderParams() bool`

HasRecoverFileAndFolderParams returns a boolean if a field has been set.

### SetRecoverFileAndFolderParamsNil

`func (o *RecoverVmwareParams) SetRecoverFileAndFolderParamsNil(b bool)`

 SetRecoverFileAndFolderParamsNil sets the value for RecoverFileAndFolderParams to be an explicit nil

### UnsetRecoverFileAndFolderParams
`func (o *RecoverVmwareParams) UnsetRecoverFileAndFolderParams()`

UnsetRecoverFileAndFolderParams ensures that no value is present for RecoverFileAndFolderParams, not even an explicit nil
### GetDownloadFileAndFolderParams

`func (o *RecoverVmwareParams) GetDownloadFileAndFolderParams() CommonDownloadFileAndFolderParams`

GetDownloadFileAndFolderParams returns the DownloadFileAndFolderParams field if non-nil, zero value otherwise.

### GetDownloadFileAndFolderParamsOk

`func (o *RecoverVmwareParams) GetDownloadFileAndFolderParamsOk() (*CommonDownloadFileAndFolderParams, bool)`

GetDownloadFileAndFolderParamsOk returns a tuple with the DownloadFileAndFolderParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadFileAndFolderParams

`func (o *RecoverVmwareParams) SetDownloadFileAndFolderParams(v CommonDownloadFileAndFolderParams)`

SetDownloadFileAndFolderParams sets DownloadFileAndFolderParams field to given value.

### HasDownloadFileAndFolderParams

`func (o *RecoverVmwareParams) HasDownloadFileAndFolderParams() bool`

HasDownloadFileAndFolderParams returns a boolean if a field has been set.

### SetDownloadFileAndFolderParamsNil

`func (o *RecoverVmwareParams) SetDownloadFileAndFolderParamsNil(b bool)`

 SetDownloadFileAndFolderParamsNil sets the value for DownloadFileAndFolderParams to be an explicit nil

### UnsetDownloadFileAndFolderParams
`func (o *RecoverVmwareParams) UnsetDownloadFileAndFolderParams()`

UnsetDownloadFileAndFolderParams ensures that no value is present for DownloadFileAndFolderParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


