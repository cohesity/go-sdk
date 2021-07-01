# RestoreO365PublicFoldersParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RootPublicFolderVec** | Pointer to [**[]RestoreO365PublicFoldersParamsRootPublicFolder**](RestoreO365PublicFoldersParamsRootPublicFolder.md) | In a RestoreJob , user will provide the list of Root Public Folders to be restored. Provision is there for restoring full and partial Public Folder recovery. | [optional] 
**TargetFolderPath** | Pointer to **NullableString** |  | [optional] 
**TargetRootPublicFolder** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 

## Methods

### NewRestoreO365PublicFoldersParams

`func NewRestoreO365PublicFoldersParams() *RestoreO365PublicFoldersParams`

NewRestoreO365PublicFoldersParams instantiates a new RestoreO365PublicFoldersParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreO365PublicFoldersParamsWithDefaults

`func NewRestoreO365PublicFoldersParamsWithDefaults() *RestoreO365PublicFoldersParams`

NewRestoreO365PublicFoldersParamsWithDefaults instantiates a new RestoreO365PublicFoldersParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRootPublicFolderVec

`func (o *RestoreO365PublicFoldersParams) GetRootPublicFolderVec() []RestoreO365PublicFoldersParamsRootPublicFolder`

GetRootPublicFolderVec returns the RootPublicFolderVec field if non-nil, zero value otherwise.

### GetRootPublicFolderVecOk

`func (o *RestoreO365PublicFoldersParams) GetRootPublicFolderVecOk() (*[]RestoreO365PublicFoldersParamsRootPublicFolder, bool)`

GetRootPublicFolderVecOk returns a tuple with the RootPublicFolderVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPublicFolderVec

`func (o *RestoreO365PublicFoldersParams) SetRootPublicFolderVec(v []RestoreO365PublicFoldersParamsRootPublicFolder)`

SetRootPublicFolderVec sets RootPublicFolderVec field to given value.

### HasRootPublicFolderVec

`func (o *RestoreO365PublicFoldersParams) HasRootPublicFolderVec() bool`

HasRootPublicFolderVec returns a boolean if a field has been set.

### SetRootPublicFolderVecNil

`func (o *RestoreO365PublicFoldersParams) SetRootPublicFolderVecNil(b bool)`

 SetRootPublicFolderVecNil sets the value for RootPublicFolderVec to be an explicit nil

### UnsetRootPublicFolderVec
`func (o *RestoreO365PublicFoldersParams) UnsetRootPublicFolderVec()`

UnsetRootPublicFolderVec ensures that no value is present for RootPublicFolderVec, not even an explicit nil
### GetTargetFolderPath

`func (o *RestoreO365PublicFoldersParams) GetTargetFolderPath() string`

GetTargetFolderPath returns the TargetFolderPath field if non-nil, zero value otherwise.

### GetTargetFolderPathOk

`func (o *RestoreO365PublicFoldersParams) GetTargetFolderPathOk() (*string, bool)`

GetTargetFolderPathOk returns a tuple with the TargetFolderPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetFolderPath

`func (o *RestoreO365PublicFoldersParams) SetTargetFolderPath(v string)`

SetTargetFolderPath sets TargetFolderPath field to given value.

### HasTargetFolderPath

`func (o *RestoreO365PublicFoldersParams) HasTargetFolderPath() bool`

HasTargetFolderPath returns a boolean if a field has been set.

### SetTargetFolderPathNil

`func (o *RestoreO365PublicFoldersParams) SetTargetFolderPathNil(b bool)`

 SetTargetFolderPathNil sets the value for TargetFolderPath to be an explicit nil

### UnsetTargetFolderPath
`func (o *RestoreO365PublicFoldersParams) UnsetTargetFolderPath()`

UnsetTargetFolderPath ensures that no value is present for TargetFolderPath, not even an explicit nil
### GetTargetRootPublicFolder

`func (o *RestoreO365PublicFoldersParams) GetTargetRootPublicFolder() EntityProto`

GetTargetRootPublicFolder returns the TargetRootPublicFolder field if non-nil, zero value otherwise.

### GetTargetRootPublicFolderOk

`func (o *RestoreO365PublicFoldersParams) GetTargetRootPublicFolderOk() (*EntityProto, bool)`

GetTargetRootPublicFolderOk returns a tuple with the TargetRootPublicFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRootPublicFolder

`func (o *RestoreO365PublicFoldersParams) SetTargetRootPublicFolder(v EntityProto)`

SetTargetRootPublicFolder sets TargetRootPublicFolder field to given value.

### HasTargetRootPublicFolder

`func (o *RestoreO365PublicFoldersParams) HasTargetRootPublicFolder() bool`

HasTargetRootPublicFolder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


