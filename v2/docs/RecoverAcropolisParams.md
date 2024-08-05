# RecoverAcropolisParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownloadFileAndFolderParams** | Pointer to [**NullableRecoverAcropolisParamsDownloadFileAndFolderParams**](RecoverAcropolisParamsDownloadFileAndFolderParams.md) |  | [optional] 
**Objects** | Pointer to [**[]RecoverAcropolisSnapshotParams**](RecoverAcropolisSnapshotParams.md) | Specifies the list of recover Object parameters. This property is mandatory for all recovery action types except recover vms. While recovering VMs, a user can specify snapshots of VM&#39;s or a Protection Group Run details to recover all the VM&#39;s that are backed up by that Run. For recovering files, specifies the object contains the file to recover. | [optional] 
**RecoverFileAndFolderParams** | Pointer to [**NullableRecoverAcropolisParamsRecoverFileAndFolderParams**](RecoverAcropolisParamsRecoverFileAndFolderParams.md) |  | [optional] 
**RecoverVmParams** | Pointer to [**NullableRecoverAcropolisParamsRecoverVmParams**](RecoverAcropolisParamsRecoverVmParams.md) |  | [optional] 
**RecoveryAction** | **string** | Specifies the type of recovery action to be performed. | 

## Methods

### NewRecoverAcropolisParams

`func NewRecoverAcropolisParams(recoveryAction string, ) *RecoverAcropolisParams`

NewRecoverAcropolisParams instantiates a new RecoverAcropolisParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverAcropolisParamsWithDefaults

`func NewRecoverAcropolisParamsWithDefaults() *RecoverAcropolisParams`

NewRecoverAcropolisParamsWithDefaults instantiates a new RecoverAcropolisParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownloadFileAndFolderParams

`func (o *RecoverAcropolisParams) GetDownloadFileAndFolderParams() RecoverAcropolisParamsDownloadFileAndFolderParams`

GetDownloadFileAndFolderParams returns the DownloadFileAndFolderParams field if non-nil, zero value otherwise.

### GetDownloadFileAndFolderParamsOk

`func (o *RecoverAcropolisParams) GetDownloadFileAndFolderParamsOk() (*RecoverAcropolisParamsDownloadFileAndFolderParams, bool)`

GetDownloadFileAndFolderParamsOk returns a tuple with the DownloadFileAndFolderParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadFileAndFolderParams

`func (o *RecoverAcropolisParams) SetDownloadFileAndFolderParams(v RecoverAcropolisParamsDownloadFileAndFolderParams)`

SetDownloadFileAndFolderParams sets DownloadFileAndFolderParams field to given value.

### HasDownloadFileAndFolderParams

`func (o *RecoverAcropolisParams) HasDownloadFileAndFolderParams() bool`

HasDownloadFileAndFolderParams returns a boolean if a field has been set.

### SetDownloadFileAndFolderParamsNil

`func (o *RecoverAcropolisParams) SetDownloadFileAndFolderParamsNil(b bool)`

 SetDownloadFileAndFolderParamsNil sets the value for DownloadFileAndFolderParams to be an explicit nil

### UnsetDownloadFileAndFolderParams
`func (o *RecoverAcropolisParams) UnsetDownloadFileAndFolderParams()`

UnsetDownloadFileAndFolderParams ensures that no value is present for DownloadFileAndFolderParams, not even an explicit nil
### GetObjects

`func (o *RecoverAcropolisParams) GetObjects() []RecoverAcropolisSnapshotParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *RecoverAcropolisParams) GetObjectsOk() (*[]RecoverAcropolisSnapshotParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *RecoverAcropolisParams) SetObjects(v []RecoverAcropolisSnapshotParams)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *RecoverAcropolisParams) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### SetObjectsNil

`func (o *RecoverAcropolisParams) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *RecoverAcropolisParams) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetRecoverFileAndFolderParams

`func (o *RecoverAcropolisParams) GetRecoverFileAndFolderParams() RecoverAcropolisParamsRecoverFileAndFolderParams`

GetRecoverFileAndFolderParams returns the RecoverFileAndFolderParams field if non-nil, zero value otherwise.

### GetRecoverFileAndFolderParamsOk

`func (o *RecoverAcropolisParams) GetRecoverFileAndFolderParamsOk() (*RecoverAcropolisParamsRecoverFileAndFolderParams, bool)`

GetRecoverFileAndFolderParamsOk returns a tuple with the RecoverFileAndFolderParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverFileAndFolderParams

`func (o *RecoverAcropolisParams) SetRecoverFileAndFolderParams(v RecoverAcropolisParamsRecoverFileAndFolderParams)`

SetRecoverFileAndFolderParams sets RecoverFileAndFolderParams field to given value.

### HasRecoverFileAndFolderParams

`func (o *RecoverAcropolisParams) HasRecoverFileAndFolderParams() bool`

HasRecoverFileAndFolderParams returns a boolean if a field has been set.

### SetRecoverFileAndFolderParamsNil

`func (o *RecoverAcropolisParams) SetRecoverFileAndFolderParamsNil(b bool)`

 SetRecoverFileAndFolderParamsNil sets the value for RecoverFileAndFolderParams to be an explicit nil

### UnsetRecoverFileAndFolderParams
`func (o *RecoverAcropolisParams) UnsetRecoverFileAndFolderParams()`

UnsetRecoverFileAndFolderParams ensures that no value is present for RecoverFileAndFolderParams, not even an explicit nil
### GetRecoverVmParams

`func (o *RecoverAcropolisParams) GetRecoverVmParams() RecoverAcropolisParamsRecoverVmParams`

GetRecoverVmParams returns the RecoverVmParams field if non-nil, zero value otherwise.

### GetRecoverVmParamsOk

`func (o *RecoverAcropolisParams) GetRecoverVmParamsOk() (*RecoverAcropolisParamsRecoverVmParams, bool)`

GetRecoverVmParamsOk returns a tuple with the RecoverVmParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverVmParams

`func (o *RecoverAcropolisParams) SetRecoverVmParams(v RecoverAcropolisParamsRecoverVmParams)`

SetRecoverVmParams sets RecoverVmParams field to given value.

### HasRecoverVmParams

`func (o *RecoverAcropolisParams) HasRecoverVmParams() bool`

HasRecoverVmParams returns a boolean if a field has been set.

### SetRecoverVmParamsNil

`func (o *RecoverAcropolisParams) SetRecoverVmParamsNil(b bool)`

 SetRecoverVmParamsNil sets the value for RecoverVmParams to be an explicit nil

### UnsetRecoverVmParams
`func (o *RecoverAcropolisParams) UnsetRecoverVmParams()`

UnsetRecoverVmParams ensures that no value is present for RecoverVmParams, not even an explicit nil
### GetRecoveryAction

`func (o *RecoverAcropolisParams) GetRecoveryAction() string`

GetRecoveryAction returns the RecoveryAction field if non-nil, zero value otherwise.

### GetRecoveryActionOk

`func (o *RecoverAcropolisParams) GetRecoveryActionOk() (*string, bool)`

GetRecoveryActionOk returns a tuple with the RecoveryAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryAction

`func (o *RecoverAcropolisParams) SetRecoveryAction(v string)`

SetRecoveryAction sets RecoveryAction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


