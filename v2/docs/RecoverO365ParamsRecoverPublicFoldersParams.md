# RecoverO365ParamsRecoverPublicFoldersParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinueOnError** | Pointer to **NullableBool** | Specifies whether to continue recovering other Public Folders if one of Public Folder failed to recover. Default value is false. | [optional] 
**RootPublicFolders** | [**[]RootPublicFolderParam**](RootPublicFolderParam.md) | Specifies a list of RootPublicFolder params associated with the objects to recover. | 
**TargetFolderPath** | Pointer to **NullableString** | Specifies the path to the target folder. | [optional] 
**TargetRootPublicFolder** | Pointer to [**NullableRecoverPublicFoldersParamsTargetRootPublicFolder**](RecoverPublicFoldersParamsTargetRootPublicFolder.md) |  | [optional] 

## Methods

### NewRecoverO365ParamsRecoverPublicFoldersParams

`func NewRecoverO365ParamsRecoverPublicFoldersParams(rootPublicFolders []RootPublicFolderParam, ) *RecoverO365ParamsRecoverPublicFoldersParams`

NewRecoverO365ParamsRecoverPublicFoldersParams instantiates a new RecoverO365ParamsRecoverPublicFoldersParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverO365ParamsRecoverPublicFoldersParamsWithDefaults

`func NewRecoverO365ParamsRecoverPublicFoldersParamsWithDefaults() *RecoverO365ParamsRecoverPublicFoldersParams`

NewRecoverO365ParamsRecoverPublicFoldersParamsWithDefaults instantiates a new RecoverO365ParamsRecoverPublicFoldersParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinueOnError

`func (o *RecoverO365ParamsRecoverPublicFoldersParams) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *RecoverO365ParamsRecoverPublicFoldersParams) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *RecoverO365ParamsRecoverPublicFoldersParams) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *RecoverO365ParamsRecoverPublicFoldersParams) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *RecoverO365ParamsRecoverPublicFoldersParams) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *RecoverO365ParamsRecoverPublicFoldersParams) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetRootPublicFolders

`func (o *RecoverO365ParamsRecoverPublicFoldersParams) GetRootPublicFolders() []RootPublicFolderParam`

GetRootPublicFolders returns the RootPublicFolders field if non-nil, zero value otherwise.

### GetRootPublicFoldersOk

`func (o *RecoverO365ParamsRecoverPublicFoldersParams) GetRootPublicFoldersOk() (*[]RootPublicFolderParam, bool)`

GetRootPublicFoldersOk returns a tuple with the RootPublicFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPublicFolders

`func (o *RecoverO365ParamsRecoverPublicFoldersParams) SetRootPublicFolders(v []RootPublicFolderParam)`

SetRootPublicFolders sets RootPublicFolders field to given value.


### SetRootPublicFoldersNil

`func (o *RecoverO365ParamsRecoverPublicFoldersParams) SetRootPublicFoldersNil(b bool)`

 SetRootPublicFoldersNil sets the value for RootPublicFolders to be an explicit nil

### UnsetRootPublicFolders
`func (o *RecoverO365ParamsRecoverPublicFoldersParams) UnsetRootPublicFolders()`

UnsetRootPublicFolders ensures that no value is present for RootPublicFolders, not even an explicit nil
### GetTargetFolderPath

`func (o *RecoverO365ParamsRecoverPublicFoldersParams) GetTargetFolderPath() string`

GetTargetFolderPath returns the TargetFolderPath field if non-nil, zero value otherwise.

### GetTargetFolderPathOk

`func (o *RecoverO365ParamsRecoverPublicFoldersParams) GetTargetFolderPathOk() (*string, bool)`

GetTargetFolderPathOk returns a tuple with the TargetFolderPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetFolderPath

`func (o *RecoverO365ParamsRecoverPublicFoldersParams) SetTargetFolderPath(v string)`

SetTargetFolderPath sets TargetFolderPath field to given value.

### HasTargetFolderPath

`func (o *RecoverO365ParamsRecoverPublicFoldersParams) HasTargetFolderPath() bool`

HasTargetFolderPath returns a boolean if a field has been set.

### SetTargetFolderPathNil

`func (o *RecoverO365ParamsRecoverPublicFoldersParams) SetTargetFolderPathNil(b bool)`

 SetTargetFolderPathNil sets the value for TargetFolderPath to be an explicit nil

### UnsetTargetFolderPath
`func (o *RecoverO365ParamsRecoverPublicFoldersParams) UnsetTargetFolderPath()`

UnsetTargetFolderPath ensures that no value is present for TargetFolderPath, not even an explicit nil
### GetTargetRootPublicFolder

`func (o *RecoverO365ParamsRecoverPublicFoldersParams) GetTargetRootPublicFolder() RecoverPublicFoldersParamsTargetRootPublicFolder`

GetTargetRootPublicFolder returns the TargetRootPublicFolder field if non-nil, zero value otherwise.

### GetTargetRootPublicFolderOk

`func (o *RecoverO365ParamsRecoverPublicFoldersParams) GetTargetRootPublicFolderOk() (*RecoverPublicFoldersParamsTargetRootPublicFolder, bool)`

GetTargetRootPublicFolderOk returns a tuple with the TargetRootPublicFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRootPublicFolder

`func (o *RecoverO365ParamsRecoverPublicFoldersParams) SetTargetRootPublicFolder(v RecoverPublicFoldersParamsTargetRootPublicFolder)`

SetTargetRootPublicFolder sets TargetRootPublicFolder field to given value.

### HasTargetRootPublicFolder

`func (o *RecoverO365ParamsRecoverPublicFoldersParams) HasTargetRootPublicFolder() bool`

HasTargetRootPublicFolder returns a boolean if a field has been set.

### SetTargetRootPublicFolderNil

`func (o *RecoverO365ParamsRecoverPublicFoldersParams) SetTargetRootPublicFolderNil(b bool)`

 SetTargetRootPublicFolderNil sets the value for TargetRootPublicFolder to be an explicit nil

### UnsetTargetRootPublicFolder
`func (o *RecoverO365ParamsRecoverPublicFoldersParams) UnsetTargetRootPublicFolder()`

UnsetTargetRootPublicFolder ensures that no value is present for TargetRootPublicFolder, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


