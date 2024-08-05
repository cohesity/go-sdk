# RecoverNetappParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownloadFileAndFolderParams** | Pointer to [**NullableRecoverAcropolisParamsDownloadFileAndFolderParams**](RecoverAcropolisParamsDownloadFileAndFolderParams.md) |  | [optional] 
**Objects** | [**[]CommonRecoverObjectSnapshotParams**](CommonRecoverObjectSnapshotParams.md) | Specifies the list of recover Object parameters. | 
**RecoverFileAndFolderParams** | Pointer to [**NullableRecoverNetappParamsRecoverFileAndFolderParams**](RecoverNetappParamsRecoverFileAndFolderParams.md) |  | [optional] 
**RecoverNasVolumeParams** | Pointer to [**NullableRecoverNetappParamsRecoverNasVolumeParams**](RecoverNetappParamsRecoverNasVolumeParams.md) |  | [optional] 
**RecoveryAction** | **string** | Specifies the type of recover action to be performed. | 

## Methods

### NewRecoverNetappParams

`func NewRecoverNetappParams(objects []CommonRecoverObjectSnapshotParams, recoveryAction string, ) *RecoverNetappParams`

NewRecoverNetappParams instantiates a new RecoverNetappParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverNetappParamsWithDefaults

`func NewRecoverNetappParamsWithDefaults() *RecoverNetappParams`

NewRecoverNetappParamsWithDefaults instantiates a new RecoverNetappParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownloadFileAndFolderParams

`func (o *RecoverNetappParams) GetDownloadFileAndFolderParams() RecoverAcropolisParamsDownloadFileAndFolderParams`

GetDownloadFileAndFolderParams returns the DownloadFileAndFolderParams field if non-nil, zero value otherwise.

### GetDownloadFileAndFolderParamsOk

`func (o *RecoverNetappParams) GetDownloadFileAndFolderParamsOk() (*RecoverAcropolisParamsDownloadFileAndFolderParams, bool)`

GetDownloadFileAndFolderParamsOk returns a tuple with the DownloadFileAndFolderParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadFileAndFolderParams

`func (o *RecoverNetappParams) SetDownloadFileAndFolderParams(v RecoverAcropolisParamsDownloadFileAndFolderParams)`

SetDownloadFileAndFolderParams sets DownloadFileAndFolderParams field to given value.

### HasDownloadFileAndFolderParams

`func (o *RecoverNetappParams) HasDownloadFileAndFolderParams() bool`

HasDownloadFileAndFolderParams returns a boolean if a field has been set.

### SetDownloadFileAndFolderParamsNil

`func (o *RecoverNetappParams) SetDownloadFileAndFolderParamsNil(b bool)`

 SetDownloadFileAndFolderParamsNil sets the value for DownloadFileAndFolderParams to be an explicit nil

### UnsetDownloadFileAndFolderParams
`func (o *RecoverNetappParams) UnsetDownloadFileAndFolderParams()`

UnsetDownloadFileAndFolderParams ensures that no value is present for DownloadFileAndFolderParams, not even an explicit nil
### GetObjects

`func (o *RecoverNetappParams) GetObjects() []CommonRecoverObjectSnapshotParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *RecoverNetappParams) GetObjectsOk() (*[]CommonRecoverObjectSnapshotParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *RecoverNetappParams) SetObjects(v []CommonRecoverObjectSnapshotParams)`

SetObjects sets Objects field to given value.


### SetObjectsNil

`func (o *RecoverNetappParams) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *RecoverNetappParams) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetRecoverFileAndFolderParams

`func (o *RecoverNetappParams) GetRecoverFileAndFolderParams() RecoverNetappParamsRecoverFileAndFolderParams`

GetRecoverFileAndFolderParams returns the RecoverFileAndFolderParams field if non-nil, zero value otherwise.

### GetRecoverFileAndFolderParamsOk

`func (o *RecoverNetappParams) GetRecoverFileAndFolderParamsOk() (*RecoverNetappParamsRecoverFileAndFolderParams, bool)`

GetRecoverFileAndFolderParamsOk returns a tuple with the RecoverFileAndFolderParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverFileAndFolderParams

`func (o *RecoverNetappParams) SetRecoverFileAndFolderParams(v RecoverNetappParamsRecoverFileAndFolderParams)`

SetRecoverFileAndFolderParams sets RecoverFileAndFolderParams field to given value.

### HasRecoverFileAndFolderParams

`func (o *RecoverNetappParams) HasRecoverFileAndFolderParams() bool`

HasRecoverFileAndFolderParams returns a boolean if a field has been set.

### SetRecoverFileAndFolderParamsNil

`func (o *RecoverNetappParams) SetRecoverFileAndFolderParamsNil(b bool)`

 SetRecoverFileAndFolderParamsNil sets the value for RecoverFileAndFolderParams to be an explicit nil

### UnsetRecoverFileAndFolderParams
`func (o *RecoverNetappParams) UnsetRecoverFileAndFolderParams()`

UnsetRecoverFileAndFolderParams ensures that no value is present for RecoverFileAndFolderParams, not even an explicit nil
### GetRecoverNasVolumeParams

`func (o *RecoverNetappParams) GetRecoverNasVolumeParams() RecoverNetappParamsRecoverNasVolumeParams`

GetRecoverNasVolumeParams returns the RecoverNasVolumeParams field if non-nil, zero value otherwise.

### GetRecoverNasVolumeParamsOk

`func (o *RecoverNetappParams) GetRecoverNasVolumeParamsOk() (*RecoverNetappParamsRecoverNasVolumeParams, bool)`

GetRecoverNasVolumeParamsOk returns a tuple with the RecoverNasVolumeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverNasVolumeParams

`func (o *RecoverNetappParams) SetRecoverNasVolumeParams(v RecoverNetappParamsRecoverNasVolumeParams)`

SetRecoverNasVolumeParams sets RecoverNasVolumeParams field to given value.

### HasRecoverNasVolumeParams

`func (o *RecoverNetappParams) HasRecoverNasVolumeParams() bool`

HasRecoverNasVolumeParams returns a boolean if a field has been set.

### SetRecoverNasVolumeParamsNil

`func (o *RecoverNetappParams) SetRecoverNasVolumeParamsNil(b bool)`

 SetRecoverNasVolumeParamsNil sets the value for RecoverNasVolumeParams to be an explicit nil

### UnsetRecoverNasVolumeParams
`func (o *RecoverNetappParams) UnsetRecoverNasVolumeParams()`

UnsetRecoverNasVolumeParams ensures that no value is present for RecoverNasVolumeParams, not even an explicit nil
### GetRecoveryAction

`func (o *RecoverNetappParams) GetRecoveryAction() string`

GetRecoveryAction returns the RecoveryAction field if non-nil, zero value otherwise.

### GetRecoveryActionOk

`func (o *RecoverNetappParams) GetRecoveryActionOk() (*string, bool)`

GetRecoveryActionOk returns a tuple with the RecoveryAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryAction

`func (o *RecoverNetappParams) SetRecoveryAction(v string)`

SetRecoveryAction sets RecoveryAction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


