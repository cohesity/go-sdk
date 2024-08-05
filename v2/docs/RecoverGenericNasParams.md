# RecoverGenericNasParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownloadFileAndFolderParams** | Pointer to [**NullableRecoverAcropolisParamsDownloadFileAndFolderParams**](RecoverAcropolisParamsDownloadFileAndFolderParams.md) |  | [optional] 
**Objects** | [**[]CommonRecoverObjectSnapshotParams**](CommonRecoverObjectSnapshotParams.md) | Specifies the list of recover Object parameters. | 
**RecoverFileAndFolderParams** | Pointer to [**NullableRecoverGenericNasParamsRecoverFileAndFolderParams**](RecoverGenericNasParamsRecoverFileAndFolderParams.md) |  | [optional] 
**RecoverNasVolumeParams** | Pointer to [**NullableRecoverGenericNasParamsRecoverNasVolumeParams**](RecoverGenericNasParamsRecoverNasVolumeParams.md) |  | [optional] 
**RecoveryAction** | **string** | Specifies the type of recover action to be performed. | 

## Methods

### NewRecoverGenericNasParams

`func NewRecoverGenericNasParams(objects []CommonRecoverObjectSnapshotParams, recoveryAction string, ) *RecoverGenericNasParams`

NewRecoverGenericNasParams instantiates a new RecoverGenericNasParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverGenericNasParamsWithDefaults

`func NewRecoverGenericNasParamsWithDefaults() *RecoverGenericNasParams`

NewRecoverGenericNasParamsWithDefaults instantiates a new RecoverGenericNasParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownloadFileAndFolderParams

`func (o *RecoverGenericNasParams) GetDownloadFileAndFolderParams() RecoverAcropolisParamsDownloadFileAndFolderParams`

GetDownloadFileAndFolderParams returns the DownloadFileAndFolderParams field if non-nil, zero value otherwise.

### GetDownloadFileAndFolderParamsOk

`func (o *RecoverGenericNasParams) GetDownloadFileAndFolderParamsOk() (*RecoverAcropolisParamsDownloadFileAndFolderParams, bool)`

GetDownloadFileAndFolderParamsOk returns a tuple with the DownloadFileAndFolderParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadFileAndFolderParams

`func (o *RecoverGenericNasParams) SetDownloadFileAndFolderParams(v RecoverAcropolisParamsDownloadFileAndFolderParams)`

SetDownloadFileAndFolderParams sets DownloadFileAndFolderParams field to given value.

### HasDownloadFileAndFolderParams

`func (o *RecoverGenericNasParams) HasDownloadFileAndFolderParams() bool`

HasDownloadFileAndFolderParams returns a boolean if a field has been set.

### SetDownloadFileAndFolderParamsNil

`func (o *RecoverGenericNasParams) SetDownloadFileAndFolderParamsNil(b bool)`

 SetDownloadFileAndFolderParamsNil sets the value for DownloadFileAndFolderParams to be an explicit nil

### UnsetDownloadFileAndFolderParams
`func (o *RecoverGenericNasParams) UnsetDownloadFileAndFolderParams()`

UnsetDownloadFileAndFolderParams ensures that no value is present for DownloadFileAndFolderParams, not even an explicit nil
### GetObjects

`func (o *RecoverGenericNasParams) GetObjects() []CommonRecoverObjectSnapshotParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *RecoverGenericNasParams) GetObjectsOk() (*[]CommonRecoverObjectSnapshotParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *RecoverGenericNasParams) SetObjects(v []CommonRecoverObjectSnapshotParams)`

SetObjects sets Objects field to given value.


### SetObjectsNil

`func (o *RecoverGenericNasParams) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *RecoverGenericNasParams) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetRecoverFileAndFolderParams

`func (o *RecoverGenericNasParams) GetRecoverFileAndFolderParams() RecoverGenericNasParamsRecoverFileAndFolderParams`

GetRecoverFileAndFolderParams returns the RecoverFileAndFolderParams field if non-nil, zero value otherwise.

### GetRecoverFileAndFolderParamsOk

`func (o *RecoverGenericNasParams) GetRecoverFileAndFolderParamsOk() (*RecoverGenericNasParamsRecoverFileAndFolderParams, bool)`

GetRecoverFileAndFolderParamsOk returns a tuple with the RecoverFileAndFolderParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverFileAndFolderParams

`func (o *RecoverGenericNasParams) SetRecoverFileAndFolderParams(v RecoverGenericNasParamsRecoverFileAndFolderParams)`

SetRecoverFileAndFolderParams sets RecoverFileAndFolderParams field to given value.

### HasRecoverFileAndFolderParams

`func (o *RecoverGenericNasParams) HasRecoverFileAndFolderParams() bool`

HasRecoverFileAndFolderParams returns a boolean if a field has been set.

### SetRecoverFileAndFolderParamsNil

`func (o *RecoverGenericNasParams) SetRecoverFileAndFolderParamsNil(b bool)`

 SetRecoverFileAndFolderParamsNil sets the value for RecoverFileAndFolderParams to be an explicit nil

### UnsetRecoverFileAndFolderParams
`func (o *RecoverGenericNasParams) UnsetRecoverFileAndFolderParams()`

UnsetRecoverFileAndFolderParams ensures that no value is present for RecoverFileAndFolderParams, not even an explicit nil
### GetRecoverNasVolumeParams

`func (o *RecoverGenericNasParams) GetRecoverNasVolumeParams() RecoverGenericNasParamsRecoverNasVolumeParams`

GetRecoverNasVolumeParams returns the RecoverNasVolumeParams field if non-nil, zero value otherwise.

### GetRecoverNasVolumeParamsOk

`func (o *RecoverGenericNasParams) GetRecoverNasVolumeParamsOk() (*RecoverGenericNasParamsRecoverNasVolumeParams, bool)`

GetRecoverNasVolumeParamsOk returns a tuple with the RecoverNasVolumeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverNasVolumeParams

`func (o *RecoverGenericNasParams) SetRecoverNasVolumeParams(v RecoverGenericNasParamsRecoverNasVolumeParams)`

SetRecoverNasVolumeParams sets RecoverNasVolumeParams field to given value.

### HasRecoverNasVolumeParams

`func (o *RecoverGenericNasParams) HasRecoverNasVolumeParams() bool`

HasRecoverNasVolumeParams returns a boolean if a field has been set.

### SetRecoverNasVolumeParamsNil

`func (o *RecoverGenericNasParams) SetRecoverNasVolumeParamsNil(b bool)`

 SetRecoverNasVolumeParamsNil sets the value for RecoverNasVolumeParams to be an explicit nil

### UnsetRecoverNasVolumeParams
`func (o *RecoverGenericNasParams) UnsetRecoverNasVolumeParams()`

UnsetRecoverNasVolumeParams ensures that no value is present for RecoverNasVolumeParams, not even an explicit nil
### GetRecoveryAction

`func (o *RecoverGenericNasParams) GetRecoveryAction() string`

GetRecoveryAction returns the RecoveryAction field if non-nil, zero value otherwise.

### GetRecoveryActionOk

`func (o *RecoverGenericNasParams) GetRecoveryActionOk() (*string, bool)`

GetRecoveryActionOk returns a tuple with the RecoveryAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryAction

`func (o *RecoverGenericNasParams) SetRecoveryAction(v string)`

SetRecoveryAction sets RecoveryAction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


