# RecoverPublicFoldersParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinueOnError** | Pointer to **NullableBool** | Specifies whether to continue recovering other Public Folders if one of Public Folder failed to recover. Default value is false. | [optional] 
**RootPublicFolders** | [**[]RootPublicFolderParam**](RootPublicFolderParam.md) | Specifies a list of RootPublicFolder params associated with the objects to recover. | 
**TargetFolderPath** | Pointer to **NullableString** | Specifies the path to the target folder. | [optional] 
**TargetRootPublicFolder** | Pointer to [**NullableRecoverPublicFoldersParamsTargetRootPublicFolder**](RecoverPublicFoldersParamsTargetRootPublicFolder.md) |  | [optional] 

## Methods

### NewRecoverPublicFoldersParams

`func NewRecoverPublicFoldersParams(rootPublicFolders []RootPublicFolderParam, ) *RecoverPublicFoldersParams`

NewRecoverPublicFoldersParams instantiates a new RecoverPublicFoldersParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverPublicFoldersParamsWithDefaults

`func NewRecoverPublicFoldersParamsWithDefaults() *RecoverPublicFoldersParams`

NewRecoverPublicFoldersParamsWithDefaults instantiates a new RecoverPublicFoldersParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinueOnError

`func (o *RecoverPublicFoldersParams) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *RecoverPublicFoldersParams) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *RecoverPublicFoldersParams) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *RecoverPublicFoldersParams) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *RecoverPublicFoldersParams) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *RecoverPublicFoldersParams) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetRootPublicFolders

`func (o *RecoverPublicFoldersParams) GetRootPublicFolders() []RootPublicFolderParam`

GetRootPublicFolders returns the RootPublicFolders field if non-nil, zero value otherwise.

### GetRootPublicFoldersOk

`func (o *RecoverPublicFoldersParams) GetRootPublicFoldersOk() (*[]RootPublicFolderParam, bool)`

GetRootPublicFoldersOk returns a tuple with the RootPublicFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPublicFolders

`func (o *RecoverPublicFoldersParams) SetRootPublicFolders(v []RootPublicFolderParam)`

SetRootPublicFolders sets RootPublicFolders field to given value.


### SetRootPublicFoldersNil

`func (o *RecoverPublicFoldersParams) SetRootPublicFoldersNil(b bool)`

 SetRootPublicFoldersNil sets the value for RootPublicFolders to be an explicit nil

### UnsetRootPublicFolders
`func (o *RecoverPublicFoldersParams) UnsetRootPublicFolders()`

UnsetRootPublicFolders ensures that no value is present for RootPublicFolders, not even an explicit nil
### GetTargetFolderPath

`func (o *RecoverPublicFoldersParams) GetTargetFolderPath() string`

GetTargetFolderPath returns the TargetFolderPath field if non-nil, zero value otherwise.

### GetTargetFolderPathOk

`func (o *RecoverPublicFoldersParams) GetTargetFolderPathOk() (*string, bool)`

GetTargetFolderPathOk returns a tuple with the TargetFolderPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetFolderPath

`func (o *RecoverPublicFoldersParams) SetTargetFolderPath(v string)`

SetTargetFolderPath sets TargetFolderPath field to given value.

### HasTargetFolderPath

`func (o *RecoverPublicFoldersParams) HasTargetFolderPath() bool`

HasTargetFolderPath returns a boolean if a field has been set.

### SetTargetFolderPathNil

`func (o *RecoverPublicFoldersParams) SetTargetFolderPathNil(b bool)`

 SetTargetFolderPathNil sets the value for TargetFolderPath to be an explicit nil

### UnsetTargetFolderPath
`func (o *RecoverPublicFoldersParams) UnsetTargetFolderPath()`

UnsetTargetFolderPath ensures that no value is present for TargetFolderPath, not even an explicit nil
### GetTargetRootPublicFolder

`func (o *RecoverPublicFoldersParams) GetTargetRootPublicFolder() RecoverPublicFoldersParamsTargetRootPublicFolder`

GetTargetRootPublicFolder returns the TargetRootPublicFolder field if non-nil, zero value otherwise.

### GetTargetRootPublicFolderOk

`func (o *RecoverPublicFoldersParams) GetTargetRootPublicFolderOk() (*RecoverPublicFoldersParamsTargetRootPublicFolder, bool)`

GetTargetRootPublicFolderOk returns a tuple with the TargetRootPublicFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRootPublicFolder

`func (o *RecoverPublicFoldersParams) SetTargetRootPublicFolder(v RecoverPublicFoldersParamsTargetRootPublicFolder)`

SetTargetRootPublicFolder sets TargetRootPublicFolder field to given value.

### HasTargetRootPublicFolder

`func (o *RecoverPublicFoldersParams) HasTargetRootPublicFolder() bool`

HasTargetRootPublicFolder returns a boolean if a field has been set.

### SetTargetRootPublicFolderNil

`func (o *RecoverPublicFoldersParams) SetTargetRootPublicFolderNil(b bool)`

 SetTargetRootPublicFolderNil sets the value for TargetRootPublicFolder to be an explicit nil

### UnsetTargetRootPublicFolder
`func (o *RecoverPublicFoldersParams) UnsetTargetRootPublicFolder()`

UnsetTargetRootPublicFolder ensures that no value is present for TargetRootPublicFolder, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


