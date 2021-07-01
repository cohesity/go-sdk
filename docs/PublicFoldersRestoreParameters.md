# PublicFoldersRestoreParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RootPublicFolderList** | Pointer to [**[]RootPublicFolder**](RootPublicFolder.md) | Specifies the list of Public Folders to be restored. | [optional] 
**TargetFolderPath** | Pointer to **NullableString** | Specifies the target folder path to restore the Public Folders. | [optional] 
**TargetRootPublicFolder** | Pointer to [**ProtectionSource**](ProtectionSource.md) |  | [optional] 

## Methods

### NewPublicFoldersRestoreParameters

`func NewPublicFoldersRestoreParameters() *PublicFoldersRestoreParameters`

NewPublicFoldersRestoreParameters instantiates a new PublicFoldersRestoreParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicFoldersRestoreParametersWithDefaults

`func NewPublicFoldersRestoreParametersWithDefaults() *PublicFoldersRestoreParameters`

NewPublicFoldersRestoreParametersWithDefaults instantiates a new PublicFoldersRestoreParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRootPublicFolderList

`func (o *PublicFoldersRestoreParameters) GetRootPublicFolderList() []RootPublicFolder`

GetRootPublicFolderList returns the RootPublicFolderList field if non-nil, zero value otherwise.

### GetRootPublicFolderListOk

`func (o *PublicFoldersRestoreParameters) GetRootPublicFolderListOk() (*[]RootPublicFolder, bool)`

GetRootPublicFolderListOk returns a tuple with the RootPublicFolderList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPublicFolderList

`func (o *PublicFoldersRestoreParameters) SetRootPublicFolderList(v []RootPublicFolder)`

SetRootPublicFolderList sets RootPublicFolderList field to given value.

### HasRootPublicFolderList

`func (o *PublicFoldersRestoreParameters) HasRootPublicFolderList() bool`

HasRootPublicFolderList returns a boolean if a field has been set.

### SetRootPublicFolderListNil

`func (o *PublicFoldersRestoreParameters) SetRootPublicFolderListNil(b bool)`

 SetRootPublicFolderListNil sets the value for RootPublicFolderList to be an explicit nil

### UnsetRootPublicFolderList
`func (o *PublicFoldersRestoreParameters) UnsetRootPublicFolderList()`

UnsetRootPublicFolderList ensures that no value is present for RootPublicFolderList, not even an explicit nil
### GetTargetFolderPath

`func (o *PublicFoldersRestoreParameters) GetTargetFolderPath() string`

GetTargetFolderPath returns the TargetFolderPath field if non-nil, zero value otherwise.

### GetTargetFolderPathOk

`func (o *PublicFoldersRestoreParameters) GetTargetFolderPathOk() (*string, bool)`

GetTargetFolderPathOk returns a tuple with the TargetFolderPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetFolderPath

`func (o *PublicFoldersRestoreParameters) SetTargetFolderPath(v string)`

SetTargetFolderPath sets TargetFolderPath field to given value.

### HasTargetFolderPath

`func (o *PublicFoldersRestoreParameters) HasTargetFolderPath() bool`

HasTargetFolderPath returns a boolean if a field has been set.

### SetTargetFolderPathNil

`func (o *PublicFoldersRestoreParameters) SetTargetFolderPathNil(b bool)`

 SetTargetFolderPathNil sets the value for TargetFolderPath to be an explicit nil

### UnsetTargetFolderPath
`func (o *PublicFoldersRestoreParameters) UnsetTargetFolderPath()`

UnsetTargetFolderPath ensures that no value is present for TargetFolderPath, not even an explicit nil
### GetTargetRootPublicFolder

`func (o *PublicFoldersRestoreParameters) GetTargetRootPublicFolder() ProtectionSource`

GetTargetRootPublicFolder returns the TargetRootPublicFolder field if non-nil, zero value otherwise.

### GetTargetRootPublicFolderOk

`func (o *PublicFoldersRestoreParameters) GetTargetRootPublicFolderOk() (*ProtectionSource, bool)`

GetTargetRootPublicFolderOk returns a tuple with the TargetRootPublicFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRootPublicFolder

`func (o *PublicFoldersRestoreParameters) SetTargetRootPublicFolder(v ProtectionSource)`

SetTargetRootPublicFolder sets TargetRootPublicFolder field to given value.

### HasTargetRootPublicFolder

`func (o *PublicFoldersRestoreParameters) HasTargetRootPublicFolder() bool`

HasTargetRootPublicFolder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


