# RecoverO365Params

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | Pointer to [**[]CommonRecoverObjectSnapshotParams**](CommonRecoverObjectSnapshotParams.md) | Specifies the list of recover Object parameters. | [optional] 
**RecoveryAction** | **string** | Specifies the type of recovery action to be performed. | 
**RecoverOneDriveParams** | Pointer to [**RecoverOneDriveParams**](RecoverOneDriveParams.md) | Specifies the parameters to recover Office 365 One Drive. | [optional] 
**RecoverMailboxParams** | Pointer to [**RecoverMailboxParams**](RecoverMailboxParams.md) | Specifies the parameters to recover Office 365 Mailbox. | [optional] 
**RecoverPublicFoldersParams** | Pointer to [**RecoverPublicFoldersParams**](RecoverPublicFoldersParams.md) | Specifies the parameters to recover Office 365 Public Folders. | [optional] 
**RecoverMsGroupParams** | Pointer to [**RecoverMsGroupParams**](RecoverMsGroupParams.md) | Specifies the parameters to recover Microsoft 365 Group. | [optional] 
**RecoverMsTeamParams** | Pointer to [**RecoverMsTeamParams**](RecoverMsTeamParams.md) | Specifies the parameters to recover Microsoft 365 Teams. | [optional] 
**RecoverSiteParams** | Pointer to [**RecoverSiteParams**](RecoverSiteParams.md) | Specifies the parameters to recover Microsoft Office 365 Sharepoint Site. | [optional] 

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


### GetRecoverOneDriveParams

`func (o *RecoverO365Params) GetRecoverOneDriveParams() RecoverOneDriveParams`

GetRecoverOneDriveParams returns the RecoverOneDriveParams field if non-nil, zero value otherwise.

### GetRecoverOneDriveParamsOk

`func (o *RecoverO365Params) GetRecoverOneDriveParamsOk() (*RecoverOneDriveParams, bool)`

GetRecoverOneDriveParamsOk returns a tuple with the RecoverOneDriveParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverOneDriveParams

`func (o *RecoverO365Params) SetRecoverOneDriveParams(v RecoverOneDriveParams)`

SetRecoverOneDriveParams sets RecoverOneDriveParams field to given value.

### HasRecoverOneDriveParams

`func (o *RecoverO365Params) HasRecoverOneDriveParams() bool`

HasRecoverOneDriveParams returns a boolean if a field has been set.

### GetRecoverMailboxParams

`func (o *RecoverO365Params) GetRecoverMailboxParams() RecoverMailboxParams`

GetRecoverMailboxParams returns the RecoverMailboxParams field if non-nil, zero value otherwise.

### GetRecoverMailboxParamsOk

`func (o *RecoverO365Params) GetRecoverMailboxParamsOk() (*RecoverMailboxParams, bool)`

GetRecoverMailboxParamsOk returns a tuple with the RecoverMailboxParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverMailboxParams

`func (o *RecoverO365Params) SetRecoverMailboxParams(v RecoverMailboxParams)`

SetRecoverMailboxParams sets RecoverMailboxParams field to given value.

### HasRecoverMailboxParams

`func (o *RecoverO365Params) HasRecoverMailboxParams() bool`

HasRecoverMailboxParams returns a boolean if a field has been set.

### GetRecoverPublicFoldersParams

`func (o *RecoverO365Params) GetRecoverPublicFoldersParams() RecoverPublicFoldersParams`

GetRecoverPublicFoldersParams returns the RecoverPublicFoldersParams field if non-nil, zero value otherwise.

### GetRecoverPublicFoldersParamsOk

`func (o *RecoverO365Params) GetRecoverPublicFoldersParamsOk() (*RecoverPublicFoldersParams, bool)`

GetRecoverPublicFoldersParamsOk returns a tuple with the RecoverPublicFoldersParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverPublicFoldersParams

`func (o *RecoverO365Params) SetRecoverPublicFoldersParams(v RecoverPublicFoldersParams)`

SetRecoverPublicFoldersParams sets RecoverPublicFoldersParams field to given value.

### HasRecoverPublicFoldersParams

`func (o *RecoverO365Params) HasRecoverPublicFoldersParams() bool`

HasRecoverPublicFoldersParams returns a boolean if a field has been set.

### GetRecoverMsGroupParams

`func (o *RecoverO365Params) GetRecoverMsGroupParams() RecoverMsGroupParams`

GetRecoverMsGroupParams returns the RecoverMsGroupParams field if non-nil, zero value otherwise.

### GetRecoverMsGroupParamsOk

`func (o *RecoverO365Params) GetRecoverMsGroupParamsOk() (*RecoverMsGroupParams, bool)`

GetRecoverMsGroupParamsOk returns a tuple with the RecoverMsGroupParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverMsGroupParams

`func (o *RecoverO365Params) SetRecoverMsGroupParams(v RecoverMsGroupParams)`

SetRecoverMsGroupParams sets RecoverMsGroupParams field to given value.

### HasRecoverMsGroupParams

`func (o *RecoverO365Params) HasRecoverMsGroupParams() bool`

HasRecoverMsGroupParams returns a boolean if a field has been set.

### GetRecoverMsTeamParams

`func (o *RecoverO365Params) GetRecoverMsTeamParams() RecoverMsTeamParams`

GetRecoverMsTeamParams returns the RecoverMsTeamParams field if non-nil, zero value otherwise.

### GetRecoverMsTeamParamsOk

`func (o *RecoverO365Params) GetRecoverMsTeamParamsOk() (*RecoverMsTeamParams, bool)`

GetRecoverMsTeamParamsOk returns a tuple with the RecoverMsTeamParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverMsTeamParams

`func (o *RecoverO365Params) SetRecoverMsTeamParams(v RecoverMsTeamParams)`

SetRecoverMsTeamParams sets RecoverMsTeamParams field to given value.

### HasRecoverMsTeamParams

`func (o *RecoverO365Params) HasRecoverMsTeamParams() bool`

HasRecoverMsTeamParams returns a boolean if a field has been set.

### GetRecoverSiteParams

`func (o *RecoverO365Params) GetRecoverSiteParams() RecoverSiteParams`

GetRecoverSiteParams returns the RecoverSiteParams field if non-nil, zero value otherwise.

### GetRecoverSiteParamsOk

`func (o *RecoverO365Params) GetRecoverSiteParamsOk() (*RecoverSiteParams, bool)`

GetRecoverSiteParamsOk returns a tuple with the RecoverSiteParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverSiteParams

`func (o *RecoverO365Params) SetRecoverSiteParams(v RecoverSiteParams)`

SetRecoverSiteParams sets RecoverSiteParams field to given value.

### HasRecoverSiteParams

`func (o *RecoverO365Params) HasRecoverSiteParams() bool`

HasRecoverSiteParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


