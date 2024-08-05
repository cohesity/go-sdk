# RecoverPhysicalParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownloadFileAndFolderParams** | Pointer to [**NullableRecoverHyperVParamsDownloadFileAndFolderParams**](RecoverHyperVParamsDownloadFileAndFolderParams.md) |  | [optional] 
**MountVolumeParams** | Pointer to [**NullableRecoverPhysicalParamsMountVolumeParams**](RecoverPhysicalParamsMountVolumeParams.md) |  | [optional] 
**Objects** | [**[]CommonRecoverObjectSnapshotParams**](CommonRecoverObjectSnapshotParams.md) | Specifies the list of Recover Object parameters. For recovering files, specifies the object contains the file to recover. | 
**RecoverFileAndFolderParams** | Pointer to [**NullableRecoverPhysicalParamsRecoverFileAndFolderParams**](RecoverPhysicalParamsRecoverFileAndFolderParams.md) |  | [optional] 
**RecoverVolumeParams** | Pointer to [**NullableRecoverPhysicalParamsRecoverVolumeParams**](RecoverPhysicalParamsRecoverVolumeParams.md) |  | [optional] 
**RecoveryAction** | **string** | Specifies the type of recover action to be performed. | 
**SystemRecoveryParams** | Pointer to [**NullableRecoverPhysicalParamsSystemRecoveryParams**](RecoverPhysicalParamsSystemRecoveryParams.md) |  | [optional] 

## Methods

### NewRecoverPhysicalParams

`func NewRecoverPhysicalParams(objects []CommonRecoverObjectSnapshotParams, recoveryAction string, ) *RecoverPhysicalParams`

NewRecoverPhysicalParams instantiates a new RecoverPhysicalParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverPhysicalParamsWithDefaults

`func NewRecoverPhysicalParamsWithDefaults() *RecoverPhysicalParams`

NewRecoverPhysicalParamsWithDefaults instantiates a new RecoverPhysicalParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownloadFileAndFolderParams

`func (o *RecoverPhysicalParams) GetDownloadFileAndFolderParams() RecoverHyperVParamsDownloadFileAndFolderParams`

GetDownloadFileAndFolderParams returns the DownloadFileAndFolderParams field if non-nil, zero value otherwise.

### GetDownloadFileAndFolderParamsOk

`func (o *RecoverPhysicalParams) GetDownloadFileAndFolderParamsOk() (*RecoverHyperVParamsDownloadFileAndFolderParams, bool)`

GetDownloadFileAndFolderParamsOk returns a tuple with the DownloadFileAndFolderParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadFileAndFolderParams

`func (o *RecoverPhysicalParams) SetDownloadFileAndFolderParams(v RecoverHyperVParamsDownloadFileAndFolderParams)`

SetDownloadFileAndFolderParams sets DownloadFileAndFolderParams field to given value.

### HasDownloadFileAndFolderParams

`func (o *RecoverPhysicalParams) HasDownloadFileAndFolderParams() bool`

HasDownloadFileAndFolderParams returns a boolean if a field has been set.

### SetDownloadFileAndFolderParamsNil

`func (o *RecoverPhysicalParams) SetDownloadFileAndFolderParamsNil(b bool)`

 SetDownloadFileAndFolderParamsNil sets the value for DownloadFileAndFolderParams to be an explicit nil

### UnsetDownloadFileAndFolderParams
`func (o *RecoverPhysicalParams) UnsetDownloadFileAndFolderParams()`

UnsetDownloadFileAndFolderParams ensures that no value is present for DownloadFileAndFolderParams, not even an explicit nil
### GetMountVolumeParams

`func (o *RecoverPhysicalParams) GetMountVolumeParams() RecoverPhysicalParamsMountVolumeParams`

GetMountVolumeParams returns the MountVolumeParams field if non-nil, zero value otherwise.

### GetMountVolumeParamsOk

`func (o *RecoverPhysicalParams) GetMountVolumeParamsOk() (*RecoverPhysicalParamsMountVolumeParams, bool)`

GetMountVolumeParamsOk returns a tuple with the MountVolumeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountVolumeParams

`func (o *RecoverPhysicalParams) SetMountVolumeParams(v RecoverPhysicalParamsMountVolumeParams)`

SetMountVolumeParams sets MountVolumeParams field to given value.

### HasMountVolumeParams

`func (o *RecoverPhysicalParams) HasMountVolumeParams() bool`

HasMountVolumeParams returns a boolean if a field has been set.

### SetMountVolumeParamsNil

`func (o *RecoverPhysicalParams) SetMountVolumeParamsNil(b bool)`

 SetMountVolumeParamsNil sets the value for MountVolumeParams to be an explicit nil

### UnsetMountVolumeParams
`func (o *RecoverPhysicalParams) UnsetMountVolumeParams()`

UnsetMountVolumeParams ensures that no value is present for MountVolumeParams, not even an explicit nil
### GetObjects

`func (o *RecoverPhysicalParams) GetObjects() []CommonRecoverObjectSnapshotParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *RecoverPhysicalParams) GetObjectsOk() (*[]CommonRecoverObjectSnapshotParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *RecoverPhysicalParams) SetObjects(v []CommonRecoverObjectSnapshotParams)`

SetObjects sets Objects field to given value.


### SetObjectsNil

`func (o *RecoverPhysicalParams) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *RecoverPhysicalParams) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetRecoverFileAndFolderParams

`func (o *RecoverPhysicalParams) GetRecoverFileAndFolderParams() RecoverPhysicalParamsRecoverFileAndFolderParams`

GetRecoverFileAndFolderParams returns the RecoverFileAndFolderParams field if non-nil, zero value otherwise.

### GetRecoverFileAndFolderParamsOk

`func (o *RecoverPhysicalParams) GetRecoverFileAndFolderParamsOk() (*RecoverPhysicalParamsRecoverFileAndFolderParams, bool)`

GetRecoverFileAndFolderParamsOk returns a tuple with the RecoverFileAndFolderParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverFileAndFolderParams

`func (o *RecoverPhysicalParams) SetRecoverFileAndFolderParams(v RecoverPhysicalParamsRecoverFileAndFolderParams)`

SetRecoverFileAndFolderParams sets RecoverFileAndFolderParams field to given value.

### HasRecoverFileAndFolderParams

`func (o *RecoverPhysicalParams) HasRecoverFileAndFolderParams() bool`

HasRecoverFileAndFolderParams returns a boolean if a field has been set.

### SetRecoverFileAndFolderParamsNil

`func (o *RecoverPhysicalParams) SetRecoverFileAndFolderParamsNil(b bool)`

 SetRecoverFileAndFolderParamsNil sets the value for RecoverFileAndFolderParams to be an explicit nil

### UnsetRecoverFileAndFolderParams
`func (o *RecoverPhysicalParams) UnsetRecoverFileAndFolderParams()`

UnsetRecoverFileAndFolderParams ensures that no value is present for RecoverFileAndFolderParams, not even an explicit nil
### GetRecoverVolumeParams

`func (o *RecoverPhysicalParams) GetRecoverVolumeParams() RecoverPhysicalParamsRecoverVolumeParams`

GetRecoverVolumeParams returns the RecoverVolumeParams field if non-nil, zero value otherwise.

### GetRecoverVolumeParamsOk

`func (o *RecoverPhysicalParams) GetRecoverVolumeParamsOk() (*RecoverPhysicalParamsRecoverVolumeParams, bool)`

GetRecoverVolumeParamsOk returns a tuple with the RecoverVolumeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverVolumeParams

`func (o *RecoverPhysicalParams) SetRecoverVolumeParams(v RecoverPhysicalParamsRecoverVolumeParams)`

SetRecoverVolumeParams sets RecoverVolumeParams field to given value.

### HasRecoverVolumeParams

`func (o *RecoverPhysicalParams) HasRecoverVolumeParams() bool`

HasRecoverVolumeParams returns a boolean if a field has been set.

### SetRecoverVolumeParamsNil

`func (o *RecoverPhysicalParams) SetRecoverVolumeParamsNil(b bool)`

 SetRecoverVolumeParamsNil sets the value for RecoverVolumeParams to be an explicit nil

### UnsetRecoverVolumeParams
`func (o *RecoverPhysicalParams) UnsetRecoverVolumeParams()`

UnsetRecoverVolumeParams ensures that no value is present for RecoverVolumeParams, not even an explicit nil
### GetRecoveryAction

`func (o *RecoverPhysicalParams) GetRecoveryAction() string`

GetRecoveryAction returns the RecoveryAction field if non-nil, zero value otherwise.

### GetRecoveryActionOk

`func (o *RecoverPhysicalParams) GetRecoveryActionOk() (*string, bool)`

GetRecoveryActionOk returns a tuple with the RecoveryAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryAction

`func (o *RecoverPhysicalParams) SetRecoveryAction(v string)`

SetRecoveryAction sets RecoveryAction field to given value.


### GetSystemRecoveryParams

`func (o *RecoverPhysicalParams) GetSystemRecoveryParams() RecoverPhysicalParamsSystemRecoveryParams`

GetSystemRecoveryParams returns the SystemRecoveryParams field if non-nil, zero value otherwise.

### GetSystemRecoveryParamsOk

`func (o *RecoverPhysicalParams) GetSystemRecoveryParamsOk() (*RecoverPhysicalParamsSystemRecoveryParams, bool)`

GetSystemRecoveryParamsOk returns a tuple with the SystemRecoveryParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemRecoveryParams

`func (o *RecoverPhysicalParams) SetSystemRecoveryParams(v RecoverPhysicalParamsSystemRecoveryParams)`

SetSystemRecoveryParams sets SystemRecoveryParams field to given value.

### HasSystemRecoveryParams

`func (o *RecoverPhysicalParams) HasSystemRecoveryParams() bool`

HasSystemRecoveryParams returns a boolean if a field has been set.

### SetSystemRecoveryParamsNil

`func (o *RecoverPhysicalParams) SetSystemRecoveryParamsNil(b bool)`

 SetSystemRecoveryParamsNil sets the value for SystemRecoveryParams to be an explicit nil

### UnsetSystemRecoveryParams
`func (o *RecoverPhysicalParams) UnsetSystemRecoveryParams()`

UnsetSystemRecoveryParams ensures that no value is present for SystemRecoveryParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


