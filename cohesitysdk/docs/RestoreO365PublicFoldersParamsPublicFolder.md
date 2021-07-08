# RestoreO365PublicFoldersParamsPublicFolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbsoluteFolderPath** | Pointer to **NullableString** | The absolute path of the folder. | [optional] 
**FolderId** | Pointer to **NullableString** | The Unique ID of the folder. | [optional] 
**IsEntireFolderRequired** | Pointer to **NullableBool** | Specify if the entire Public Folder is to be restored. | [optional] 
**ItemIdVec** | Pointer to **[]string** | If is_entire_folder_required is set to false, users will then specify which particular items are to be restored. | [optional] 

## Methods

### NewRestoreO365PublicFoldersParamsPublicFolder

`func NewRestoreO365PublicFoldersParamsPublicFolder() *RestoreO365PublicFoldersParamsPublicFolder`

NewRestoreO365PublicFoldersParamsPublicFolder instantiates a new RestoreO365PublicFoldersParamsPublicFolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreO365PublicFoldersParamsPublicFolderWithDefaults

`func NewRestoreO365PublicFoldersParamsPublicFolderWithDefaults() *RestoreO365PublicFoldersParamsPublicFolder`

NewRestoreO365PublicFoldersParamsPublicFolderWithDefaults instantiates a new RestoreO365PublicFoldersParamsPublicFolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbsoluteFolderPath

`func (o *RestoreO365PublicFoldersParamsPublicFolder) GetAbsoluteFolderPath() string`

GetAbsoluteFolderPath returns the AbsoluteFolderPath field if non-nil, zero value otherwise.

### GetAbsoluteFolderPathOk

`func (o *RestoreO365PublicFoldersParamsPublicFolder) GetAbsoluteFolderPathOk() (*string, bool)`

GetAbsoluteFolderPathOk returns a tuple with the AbsoluteFolderPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteFolderPath

`func (o *RestoreO365PublicFoldersParamsPublicFolder) SetAbsoluteFolderPath(v string)`

SetAbsoluteFolderPath sets AbsoluteFolderPath field to given value.

### HasAbsoluteFolderPath

`func (o *RestoreO365PublicFoldersParamsPublicFolder) HasAbsoluteFolderPath() bool`

HasAbsoluteFolderPath returns a boolean if a field has been set.

### SetAbsoluteFolderPathNil

`func (o *RestoreO365PublicFoldersParamsPublicFolder) SetAbsoluteFolderPathNil(b bool)`

 SetAbsoluteFolderPathNil sets the value for AbsoluteFolderPath to be an explicit nil

### UnsetAbsoluteFolderPath
`func (o *RestoreO365PublicFoldersParamsPublicFolder) UnsetAbsoluteFolderPath()`

UnsetAbsoluteFolderPath ensures that no value is present for AbsoluteFolderPath, not even an explicit nil
### GetFolderId

`func (o *RestoreO365PublicFoldersParamsPublicFolder) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *RestoreO365PublicFoldersParamsPublicFolder) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *RestoreO365PublicFoldersParamsPublicFolder) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *RestoreO365PublicFoldersParamsPublicFolder) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### SetFolderIdNil

`func (o *RestoreO365PublicFoldersParamsPublicFolder) SetFolderIdNil(b bool)`

 SetFolderIdNil sets the value for FolderId to be an explicit nil

### UnsetFolderId
`func (o *RestoreO365PublicFoldersParamsPublicFolder) UnsetFolderId()`

UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
### GetIsEntireFolderRequired

`func (o *RestoreO365PublicFoldersParamsPublicFolder) GetIsEntireFolderRequired() bool`

GetIsEntireFolderRequired returns the IsEntireFolderRequired field if non-nil, zero value otherwise.

### GetIsEntireFolderRequiredOk

`func (o *RestoreO365PublicFoldersParamsPublicFolder) GetIsEntireFolderRequiredOk() (*bool, bool)`

GetIsEntireFolderRequiredOk returns a tuple with the IsEntireFolderRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEntireFolderRequired

`func (o *RestoreO365PublicFoldersParamsPublicFolder) SetIsEntireFolderRequired(v bool)`

SetIsEntireFolderRequired sets IsEntireFolderRequired field to given value.

### HasIsEntireFolderRequired

`func (o *RestoreO365PublicFoldersParamsPublicFolder) HasIsEntireFolderRequired() bool`

HasIsEntireFolderRequired returns a boolean if a field has been set.

### SetIsEntireFolderRequiredNil

`func (o *RestoreO365PublicFoldersParamsPublicFolder) SetIsEntireFolderRequiredNil(b bool)`

 SetIsEntireFolderRequiredNil sets the value for IsEntireFolderRequired to be an explicit nil

### UnsetIsEntireFolderRequired
`func (o *RestoreO365PublicFoldersParamsPublicFolder) UnsetIsEntireFolderRequired()`

UnsetIsEntireFolderRequired ensures that no value is present for IsEntireFolderRequired, not even an explicit nil
### GetItemIdVec

`func (o *RestoreO365PublicFoldersParamsPublicFolder) GetItemIdVec() []string`

GetItemIdVec returns the ItemIdVec field if non-nil, zero value otherwise.

### GetItemIdVecOk

`func (o *RestoreO365PublicFoldersParamsPublicFolder) GetItemIdVecOk() (*[]string, bool)`

GetItemIdVecOk returns a tuple with the ItemIdVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemIdVec

`func (o *RestoreO365PublicFoldersParamsPublicFolder) SetItemIdVec(v []string)`

SetItemIdVec sets ItemIdVec field to given value.

### HasItemIdVec

`func (o *RestoreO365PublicFoldersParamsPublicFolder) HasItemIdVec() bool`

HasItemIdVec returns a boolean if a field has been set.

### SetItemIdVecNil

`func (o *RestoreO365PublicFoldersParamsPublicFolder) SetItemIdVecNil(b bool)`

 SetItemIdVecNil sets the value for ItemIdVec to be an explicit nil

### UnsetItemIdVec
`func (o *RestoreO365PublicFoldersParamsPublicFolder) UnsetItemIdVec()`

UnsetItemIdVec ensures that no value is present for ItemIdVec, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


