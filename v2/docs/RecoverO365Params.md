# RecoverO365Params

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownloadChatsParams** | Pointer to [**NullableRecoverO365ParamsDownloadChatsParams**](RecoverO365ParamsDownloadChatsParams.md) |  | [optional] 
**DownloadFileAndFolderParams** | Pointer to [**NullableRecoverO365ParamsDownloadFileAndFolderParams**](RecoverO365ParamsDownloadFileAndFolderParams.md) |  | [optional] 
**Objects** | Pointer to [**[]CommonRecoverObjectSnapshotParams**](CommonRecoverObjectSnapshotParams.md) | Specifies the list of recover Object parameters. | [optional] 
**RecoverMailboxParams** | Pointer to [**RecoverO365ParamsRecoverMailboxParams**](RecoverO365ParamsRecoverMailboxParams.md) |  | [optional] 
**RecoverMsGroupParams** | Pointer to [**RecoverO365ParamsRecoverMsGroupParams**](RecoverO365ParamsRecoverMsGroupParams.md) |  | [optional] 
**RecoverMsTeamParams** | Pointer to [**RecoverO365ParamsRecoverMsTeamParams**](RecoverO365ParamsRecoverMsTeamParams.md) |  | [optional] 
**RecoverOneDriveParams** | Pointer to [**RecoverO365ParamsRecoverOneDriveParams**](RecoverO365ParamsRecoverOneDriveParams.md) |  | [optional] 
**RecoverPublicFoldersParams** | Pointer to [**RecoverO365ParamsRecoverPublicFoldersParams**](RecoverO365ParamsRecoverPublicFoldersParams.md) |  | [optional] 
**RecoverSiteParams** | Pointer to [**RecoverO365ParamsRecoverSiteParams**](RecoverO365ParamsRecoverSiteParams.md) |  | [optional] 
**RecoveryAction** | **string** | Specifies the type of recovery action to be performed. | 

## Methods

### NewRecoverO365Params

`func NewRecoverO365Params(recoveryAction string, ) *RecoverO365Params`

NewRecoverO365Params instantiates a new RecoverO365Params object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverO365ParamsWithDefaults

`func NewRecoverO365ParamsWithDefaults() *RecoverO365Params`

NewRecoverO365ParamsWithDefaults instantiates a new RecoverO365Params object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownloadChatsParams

`func (o *RecoverO365Params) GetDownloadChatsParams() RecoverO365ParamsDownloadChatsParams`

GetDownloadChatsParams returns the DownloadChatsParams field if non-nil, zero value otherwise.

### GetDownloadChatsParamsOk

`func (o *RecoverO365Params) GetDownloadChatsParamsOk() (*RecoverO365ParamsDownloadChatsParams, bool)`

GetDownloadChatsParamsOk returns a tuple with the DownloadChatsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadChatsParams

`func (o *RecoverO365Params) SetDownloadChatsParams(v RecoverO365ParamsDownloadChatsParams)`

SetDownloadChatsParams sets DownloadChatsParams field to given value.

### HasDownloadChatsParams

`func (o *RecoverO365Params) HasDownloadChatsParams() bool`

HasDownloadChatsParams returns a boolean if a field has been set.

### SetDownloadChatsParamsNil

`func (o *RecoverO365Params) SetDownloadChatsParamsNil(b bool)`

 SetDownloadChatsParamsNil sets the value for DownloadChatsParams to be an explicit nil

### UnsetDownloadChatsParams
`func (o *RecoverO365Params) UnsetDownloadChatsParams()`

UnsetDownloadChatsParams ensures that no value is present for DownloadChatsParams, not even an explicit nil
### GetDownloadFileAndFolderParams

`func (o *RecoverO365Params) GetDownloadFileAndFolderParams() RecoverO365ParamsDownloadFileAndFolderParams`

GetDownloadFileAndFolderParams returns the DownloadFileAndFolderParams field if non-nil, zero value otherwise.

### GetDownloadFileAndFolderParamsOk

`func (o *RecoverO365Params) GetDownloadFileAndFolderParamsOk() (*RecoverO365ParamsDownloadFileAndFolderParams, bool)`

GetDownloadFileAndFolderParamsOk returns a tuple with the DownloadFileAndFolderParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadFileAndFolderParams

`func (o *RecoverO365Params) SetDownloadFileAndFolderParams(v RecoverO365ParamsDownloadFileAndFolderParams)`

SetDownloadFileAndFolderParams sets DownloadFileAndFolderParams field to given value.

### HasDownloadFileAndFolderParams

`func (o *RecoverO365Params) HasDownloadFileAndFolderParams() bool`

HasDownloadFileAndFolderParams returns a boolean if a field has been set.

### SetDownloadFileAndFolderParamsNil

`func (o *RecoverO365Params) SetDownloadFileAndFolderParamsNil(b bool)`

 SetDownloadFileAndFolderParamsNil sets the value for DownloadFileAndFolderParams to be an explicit nil

### UnsetDownloadFileAndFolderParams
`func (o *RecoverO365Params) UnsetDownloadFileAndFolderParams()`

UnsetDownloadFileAndFolderParams ensures that no value is present for DownloadFileAndFolderParams, not even an explicit nil
### GetObjects

`func (o *RecoverO365Params) GetObjects() []CommonRecoverObjectSnapshotParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *RecoverO365Params) GetObjectsOk() (*[]CommonRecoverObjectSnapshotParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *RecoverO365Params) SetObjects(v []CommonRecoverObjectSnapshotParams)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *RecoverO365Params) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### SetObjectsNil

`func (o *RecoverO365Params) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *RecoverO365Params) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetRecoverMailboxParams

`func (o *RecoverO365Params) GetRecoverMailboxParams() RecoverO365ParamsRecoverMailboxParams`

GetRecoverMailboxParams returns the RecoverMailboxParams field if non-nil, zero value otherwise.

### GetRecoverMailboxParamsOk

`func (o *RecoverO365Params) GetRecoverMailboxParamsOk() (*RecoverO365ParamsRecoverMailboxParams, bool)`

GetRecoverMailboxParamsOk returns a tuple with the RecoverMailboxParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverMailboxParams

`func (o *RecoverO365Params) SetRecoverMailboxParams(v RecoverO365ParamsRecoverMailboxParams)`

SetRecoverMailboxParams sets RecoverMailboxParams field to given value.

### HasRecoverMailboxParams

`func (o *RecoverO365Params) HasRecoverMailboxParams() bool`

HasRecoverMailboxParams returns a boolean if a field has been set.

### GetRecoverMsGroupParams

`func (o *RecoverO365Params) GetRecoverMsGroupParams() RecoverO365ParamsRecoverMsGroupParams`

GetRecoverMsGroupParams returns the RecoverMsGroupParams field if non-nil, zero value otherwise.

### GetRecoverMsGroupParamsOk

`func (o *RecoverO365Params) GetRecoverMsGroupParamsOk() (*RecoverO365ParamsRecoverMsGroupParams, bool)`

GetRecoverMsGroupParamsOk returns a tuple with the RecoverMsGroupParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverMsGroupParams

`func (o *RecoverO365Params) SetRecoverMsGroupParams(v RecoverO365ParamsRecoverMsGroupParams)`

SetRecoverMsGroupParams sets RecoverMsGroupParams field to given value.

### HasRecoverMsGroupParams

`func (o *RecoverO365Params) HasRecoverMsGroupParams() bool`

HasRecoverMsGroupParams returns a boolean if a field has been set.

### GetRecoverMsTeamParams

`func (o *RecoverO365Params) GetRecoverMsTeamParams() RecoverO365ParamsRecoverMsTeamParams`

GetRecoverMsTeamParams returns the RecoverMsTeamParams field if non-nil, zero value otherwise.

### GetRecoverMsTeamParamsOk

`func (o *RecoverO365Params) GetRecoverMsTeamParamsOk() (*RecoverO365ParamsRecoverMsTeamParams, bool)`

GetRecoverMsTeamParamsOk returns a tuple with the RecoverMsTeamParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverMsTeamParams

`func (o *RecoverO365Params) SetRecoverMsTeamParams(v RecoverO365ParamsRecoverMsTeamParams)`

SetRecoverMsTeamParams sets RecoverMsTeamParams field to given value.

### HasRecoverMsTeamParams

`func (o *RecoverO365Params) HasRecoverMsTeamParams() bool`

HasRecoverMsTeamParams returns a boolean if a field has been set.

### GetRecoverOneDriveParams

`func (o *RecoverO365Params) GetRecoverOneDriveParams() RecoverO365ParamsRecoverOneDriveParams`

GetRecoverOneDriveParams returns the RecoverOneDriveParams field if non-nil, zero value otherwise.

### GetRecoverOneDriveParamsOk

`func (o *RecoverO365Params) GetRecoverOneDriveParamsOk() (*RecoverO365ParamsRecoverOneDriveParams, bool)`

GetRecoverOneDriveParamsOk returns a tuple with the RecoverOneDriveParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverOneDriveParams

`func (o *RecoverO365Params) SetRecoverOneDriveParams(v RecoverO365ParamsRecoverOneDriveParams)`

SetRecoverOneDriveParams sets RecoverOneDriveParams field to given value.

### HasRecoverOneDriveParams

`func (o *RecoverO365Params) HasRecoverOneDriveParams() bool`

HasRecoverOneDriveParams returns a boolean if a field has been set.

### GetRecoverPublicFoldersParams

`func (o *RecoverO365Params) GetRecoverPublicFoldersParams() RecoverO365ParamsRecoverPublicFoldersParams`

GetRecoverPublicFoldersParams returns the RecoverPublicFoldersParams field if non-nil, zero value otherwise.

### GetRecoverPublicFoldersParamsOk

`func (o *RecoverO365Params) GetRecoverPublicFoldersParamsOk() (*RecoverO365ParamsRecoverPublicFoldersParams, bool)`

GetRecoverPublicFoldersParamsOk returns a tuple with the RecoverPublicFoldersParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverPublicFoldersParams

`func (o *RecoverO365Params) SetRecoverPublicFoldersParams(v RecoverO365ParamsRecoverPublicFoldersParams)`

SetRecoverPublicFoldersParams sets RecoverPublicFoldersParams field to given value.

### HasRecoverPublicFoldersParams

`func (o *RecoverO365Params) HasRecoverPublicFoldersParams() bool`

HasRecoverPublicFoldersParams returns a boolean if a field has been set.

### GetRecoverSiteParams

`func (o *RecoverO365Params) GetRecoverSiteParams() RecoverO365ParamsRecoverSiteParams`

GetRecoverSiteParams returns the RecoverSiteParams field if non-nil, zero value otherwise.

### GetRecoverSiteParamsOk

`func (o *RecoverO365Params) GetRecoverSiteParamsOk() (*RecoverO365ParamsRecoverSiteParams, bool)`

GetRecoverSiteParamsOk returns a tuple with the RecoverSiteParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverSiteParams

`func (o *RecoverO365Params) SetRecoverSiteParams(v RecoverO365ParamsRecoverSiteParams)`

SetRecoverSiteParams sets RecoverSiteParams field to given value.

### HasRecoverSiteParams

`func (o *RecoverO365Params) HasRecoverSiteParams() bool`

HasRecoverSiteParams returns a boolean if a field has been set.

### GetRecoveryAction

`func (o *RecoverO365Params) GetRecoveryAction() string`

GetRecoveryAction returns the RecoveryAction field if non-nil, zero value otherwise.

### GetRecoveryActionOk

`func (o *RecoverO365Params) GetRecoveryActionOk() (*string, bool)`

GetRecoveryActionOk returns a tuple with the RecoveryAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryAction

`func (o *RecoverO365Params) SetRecoveryAction(v string)`

SetRecoveryAction sets RecoveryAction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


