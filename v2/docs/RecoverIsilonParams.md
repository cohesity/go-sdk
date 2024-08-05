# RecoverIsilonParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownloadFileAndFolderParams** | Pointer to [**NullableRecoverAcropolisParamsDownloadFileAndFolderParams**](RecoverAcropolisParamsDownloadFileAndFolderParams.md) |  | [optional] 
**Objects** | [**[]CommonRecoverObjectSnapshotParams**](CommonRecoverObjectSnapshotParams.md) | Specifies the list of recover Object parameters. | 
**RecoverFileAndFolderParams** | Pointer to [**NullableRecoverIsilonParamsRecoverFileAndFolderParams**](RecoverIsilonParamsRecoverFileAndFolderParams.md) |  | [optional] 
**RecoverNasVolumeParams** | Pointer to [**NullableRecoverIsilonParamsRecoverNasVolumeParams**](RecoverIsilonParamsRecoverNasVolumeParams.md) |  | [optional] 
**RecoveryAction** | **string** | Specifies the type of recover action to be performed. | 

## Methods

### NewRecoverIsilonParams

`func NewRecoverIsilonParams(objects []CommonRecoverObjectSnapshotParams, recoveryAction string, ) *RecoverIsilonParams`

NewRecoverIsilonParams instantiates a new RecoverIsilonParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverIsilonParamsWithDefaults

`func NewRecoverIsilonParamsWithDefaults() *RecoverIsilonParams`

NewRecoverIsilonParamsWithDefaults instantiates a new RecoverIsilonParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownloadFileAndFolderParams

`func (o *RecoverIsilonParams) GetDownloadFileAndFolderParams() RecoverAcropolisParamsDownloadFileAndFolderParams`

GetDownloadFileAndFolderParams returns the DownloadFileAndFolderParams field if non-nil, zero value otherwise.

### GetDownloadFileAndFolderParamsOk

`func (o *RecoverIsilonParams) GetDownloadFileAndFolderParamsOk() (*RecoverAcropolisParamsDownloadFileAndFolderParams, bool)`

GetDownloadFileAndFolderParamsOk returns a tuple with the DownloadFileAndFolderParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadFileAndFolderParams

`func (o *RecoverIsilonParams) SetDownloadFileAndFolderParams(v RecoverAcropolisParamsDownloadFileAndFolderParams)`

SetDownloadFileAndFolderParams sets DownloadFileAndFolderParams field to given value.

### HasDownloadFileAndFolderParams

`func (o *RecoverIsilonParams) HasDownloadFileAndFolderParams() bool`

HasDownloadFileAndFolderParams returns a boolean if a field has been set.

### SetDownloadFileAndFolderParamsNil

`func (o *RecoverIsilonParams) SetDownloadFileAndFolderParamsNil(b bool)`

 SetDownloadFileAndFolderParamsNil sets the value for DownloadFileAndFolderParams to be an explicit nil

### UnsetDownloadFileAndFolderParams
`func (o *RecoverIsilonParams) UnsetDownloadFileAndFolderParams()`

UnsetDownloadFileAndFolderParams ensures that no value is present for DownloadFileAndFolderParams, not even an explicit nil
### GetObjects

`func (o *RecoverIsilonParams) GetObjects() []CommonRecoverObjectSnapshotParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *RecoverIsilonParams) GetObjectsOk() (*[]CommonRecoverObjectSnapshotParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *RecoverIsilonParams) SetObjects(v []CommonRecoverObjectSnapshotParams)`

SetObjects sets Objects field to given value.


### SetObjectsNil

`func (o *RecoverIsilonParams) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *RecoverIsilonParams) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetRecoverFileAndFolderParams

`func (o *RecoverIsilonParams) GetRecoverFileAndFolderParams() RecoverIsilonParamsRecoverFileAndFolderParams`

GetRecoverFileAndFolderParams returns the RecoverFileAndFolderParams field if non-nil, zero value otherwise.

### GetRecoverFileAndFolderParamsOk

`func (o *RecoverIsilonParams) GetRecoverFileAndFolderParamsOk() (*RecoverIsilonParamsRecoverFileAndFolderParams, bool)`

GetRecoverFileAndFolderParamsOk returns a tuple with the RecoverFileAndFolderParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverFileAndFolderParams

`func (o *RecoverIsilonParams) SetRecoverFileAndFolderParams(v RecoverIsilonParamsRecoverFileAndFolderParams)`

SetRecoverFileAndFolderParams sets RecoverFileAndFolderParams field to given value.

### HasRecoverFileAndFolderParams

`func (o *RecoverIsilonParams) HasRecoverFileAndFolderParams() bool`

HasRecoverFileAndFolderParams returns a boolean if a field has been set.

### SetRecoverFileAndFolderParamsNil

`func (o *RecoverIsilonParams) SetRecoverFileAndFolderParamsNil(b bool)`

 SetRecoverFileAndFolderParamsNil sets the value for RecoverFileAndFolderParams to be an explicit nil

### UnsetRecoverFileAndFolderParams
`func (o *RecoverIsilonParams) UnsetRecoverFileAndFolderParams()`

UnsetRecoverFileAndFolderParams ensures that no value is present for RecoverFileAndFolderParams, not even an explicit nil
### GetRecoverNasVolumeParams

`func (o *RecoverIsilonParams) GetRecoverNasVolumeParams() RecoverIsilonParamsRecoverNasVolumeParams`

GetRecoverNasVolumeParams returns the RecoverNasVolumeParams field if non-nil, zero value otherwise.

### GetRecoverNasVolumeParamsOk

`func (o *RecoverIsilonParams) GetRecoverNasVolumeParamsOk() (*RecoverIsilonParamsRecoverNasVolumeParams, bool)`

GetRecoverNasVolumeParamsOk returns a tuple with the RecoverNasVolumeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverNasVolumeParams

`func (o *RecoverIsilonParams) SetRecoverNasVolumeParams(v RecoverIsilonParamsRecoverNasVolumeParams)`

SetRecoverNasVolumeParams sets RecoverNasVolumeParams field to given value.

### HasRecoverNasVolumeParams

`func (o *RecoverIsilonParams) HasRecoverNasVolumeParams() bool`

HasRecoverNasVolumeParams returns a boolean if a field has been set.

### SetRecoverNasVolumeParamsNil

`func (o *RecoverIsilonParams) SetRecoverNasVolumeParamsNil(b bool)`

 SetRecoverNasVolumeParamsNil sets the value for RecoverNasVolumeParams to be an explicit nil

### UnsetRecoverNasVolumeParams
`func (o *RecoverIsilonParams) UnsetRecoverNasVolumeParams()`

UnsetRecoverNasVolumeParams ensures that no value is present for RecoverNasVolumeParams, not even an explicit nil
### GetRecoveryAction

`func (o *RecoverIsilonParams) GetRecoveryAction() string`

GetRecoveryAction returns the RecoveryAction field if non-nil, zero value otherwise.

### GetRecoveryActionOk

`func (o *RecoverIsilonParams) GetRecoveryActionOk() (*string, bool)`

GetRecoveryActionOk returns a tuple with the RecoveryAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryAction

`func (o *RecoverIsilonParams) SetRecoveryAction(v string)`

SetRecoveryAction sets RecoveryAction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


