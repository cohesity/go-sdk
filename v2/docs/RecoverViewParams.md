# RecoverViewParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownloadFileAndFolderParams** | Pointer to [**NullableRecoverViewParamsDownloadFileAndFolderParams**](RecoverViewParamsDownloadFileAndFolderParams.md) |  | [optional] 
**Objects** | Pointer to [**[]CommonRecoverObjectSnapshotParams**](CommonRecoverObjectSnapshotParams.md) | Specifies the list of recover Object parameters. | [optional] 
**RecoverFileAndFolderParams** | Pointer to [**NullableRecoverViewParamsRecoverFileAndFolderParams**](RecoverViewParamsRecoverFileAndFolderParams.md) |  | [optional] 
**RecoveryAction** | **string** | Specifies the type of recovery action to be performed. | 

## Methods

### NewRecoverViewParams

`func NewRecoverViewParams(recoveryAction string, ) *RecoverViewParams`

NewRecoverViewParams instantiates a new RecoverViewParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverViewParamsWithDefaults

`func NewRecoverViewParamsWithDefaults() *RecoverViewParams`

NewRecoverViewParamsWithDefaults instantiates a new RecoverViewParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownloadFileAndFolderParams

`func (o *RecoverViewParams) GetDownloadFileAndFolderParams() RecoverViewParamsDownloadFileAndFolderParams`

GetDownloadFileAndFolderParams returns the DownloadFileAndFolderParams field if non-nil, zero value otherwise.

### GetDownloadFileAndFolderParamsOk

`func (o *RecoverViewParams) GetDownloadFileAndFolderParamsOk() (*RecoverViewParamsDownloadFileAndFolderParams, bool)`

GetDownloadFileAndFolderParamsOk returns a tuple with the DownloadFileAndFolderParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadFileAndFolderParams

`func (o *RecoverViewParams) SetDownloadFileAndFolderParams(v RecoverViewParamsDownloadFileAndFolderParams)`

SetDownloadFileAndFolderParams sets DownloadFileAndFolderParams field to given value.

### HasDownloadFileAndFolderParams

`func (o *RecoverViewParams) HasDownloadFileAndFolderParams() bool`

HasDownloadFileAndFolderParams returns a boolean if a field has been set.

### SetDownloadFileAndFolderParamsNil

`func (o *RecoverViewParams) SetDownloadFileAndFolderParamsNil(b bool)`

 SetDownloadFileAndFolderParamsNil sets the value for DownloadFileAndFolderParams to be an explicit nil

### UnsetDownloadFileAndFolderParams
`func (o *RecoverViewParams) UnsetDownloadFileAndFolderParams()`

UnsetDownloadFileAndFolderParams ensures that no value is present for DownloadFileAndFolderParams, not even an explicit nil
### GetObjects

`func (o *RecoverViewParams) GetObjects() []CommonRecoverObjectSnapshotParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *RecoverViewParams) GetObjectsOk() (*[]CommonRecoverObjectSnapshotParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *RecoverViewParams) SetObjects(v []CommonRecoverObjectSnapshotParams)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *RecoverViewParams) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### SetObjectsNil

`func (o *RecoverViewParams) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *RecoverViewParams) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetRecoverFileAndFolderParams

`func (o *RecoverViewParams) GetRecoverFileAndFolderParams() RecoverViewParamsRecoverFileAndFolderParams`

GetRecoverFileAndFolderParams returns the RecoverFileAndFolderParams field if non-nil, zero value otherwise.

### GetRecoverFileAndFolderParamsOk

`func (o *RecoverViewParams) GetRecoverFileAndFolderParamsOk() (*RecoverViewParamsRecoverFileAndFolderParams, bool)`

GetRecoverFileAndFolderParamsOk returns a tuple with the RecoverFileAndFolderParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverFileAndFolderParams

`func (o *RecoverViewParams) SetRecoverFileAndFolderParams(v RecoverViewParamsRecoverFileAndFolderParams)`

SetRecoverFileAndFolderParams sets RecoverFileAndFolderParams field to given value.

### HasRecoverFileAndFolderParams

`func (o *RecoverViewParams) HasRecoverFileAndFolderParams() bool`

HasRecoverFileAndFolderParams returns a boolean if a field has been set.

### SetRecoverFileAndFolderParamsNil

`func (o *RecoverViewParams) SetRecoverFileAndFolderParamsNil(b bool)`

 SetRecoverFileAndFolderParamsNil sets the value for RecoverFileAndFolderParams to be an explicit nil

### UnsetRecoverFileAndFolderParams
`func (o *RecoverViewParams) UnsetRecoverFileAndFolderParams()`

UnsetRecoverFileAndFolderParams ensures that no value is present for RecoverFileAndFolderParams, not even an explicit nil
### GetRecoveryAction

`func (o *RecoverViewParams) GetRecoveryAction() string`

GetRecoveryAction returns the RecoveryAction field if non-nil, zero value otherwise.

### GetRecoveryActionOk

`func (o *RecoverViewParams) GetRecoveryActionOk() (*string, bool)`

GetRecoveryActionOk returns a tuple with the RecoveryAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryAction

`func (o *RecoverViewParams) SetRecoveryAction(v string)`

SetRecoveryAction sets RecoveryAction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


