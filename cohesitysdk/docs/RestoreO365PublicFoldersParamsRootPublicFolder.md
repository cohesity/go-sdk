# RestoreO365PublicFoldersParamsRootPublicFolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FolderVec** | Pointer to [**[]RestoreO365PublicFoldersParamsPublicFolder**](RestoreO365PublicFoldersParamsPublicFolder.md) | If is_entire_folder_required is set to false, user will then specify which particular sub-folders are to be restored. | [optional] 
**IsEntireRootFolderRequired** | Pointer to **NullableBool** | Specify if the entire Root Public Folder including the sub-folders is to be restored. | [optional] 
**Object** | Pointer to [**RestoreObject**](RestoreObject.md) |  | [optional] 

## Methods

### NewRestoreO365PublicFoldersParamsRootPublicFolder

`func NewRestoreO365PublicFoldersParamsRootPublicFolder() *RestoreO365PublicFoldersParamsRootPublicFolder`

NewRestoreO365PublicFoldersParamsRootPublicFolder instantiates a new RestoreO365PublicFoldersParamsRootPublicFolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreO365PublicFoldersParamsRootPublicFolderWithDefaults

`func NewRestoreO365PublicFoldersParamsRootPublicFolderWithDefaults() *RestoreO365PublicFoldersParamsRootPublicFolder`

NewRestoreO365PublicFoldersParamsRootPublicFolderWithDefaults instantiates a new RestoreO365PublicFoldersParamsRootPublicFolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFolderVec

`func (o *RestoreO365PublicFoldersParamsRootPublicFolder) GetFolderVec() []RestoreO365PublicFoldersParamsPublicFolder`

GetFolderVec returns the FolderVec field if non-nil, zero value otherwise.

### GetFolderVecOk

`func (o *RestoreO365PublicFoldersParamsRootPublicFolder) GetFolderVecOk() (*[]RestoreO365PublicFoldersParamsPublicFolder, bool)`

GetFolderVecOk returns a tuple with the FolderVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderVec

`func (o *RestoreO365PublicFoldersParamsRootPublicFolder) SetFolderVec(v []RestoreO365PublicFoldersParamsPublicFolder)`

SetFolderVec sets FolderVec field to given value.

### HasFolderVec

`func (o *RestoreO365PublicFoldersParamsRootPublicFolder) HasFolderVec() bool`

HasFolderVec returns a boolean if a field has been set.

### SetFolderVecNil

`func (o *RestoreO365PublicFoldersParamsRootPublicFolder) SetFolderVecNil(b bool)`

 SetFolderVecNil sets the value for FolderVec to be an explicit nil

### UnsetFolderVec
`func (o *RestoreO365PublicFoldersParamsRootPublicFolder) UnsetFolderVec()`

UnsetFolderVec ensures that no value is present for FolderVec, not even an explicit nil
### GetIsEntireRootFolderRequired

`func (o *RestoreO365PublicFoldersParamsRootPublicFolder) GetIsEntireRootFolderRequired() bool`

GetIsEntireRootFolderRequired returns the IsEntireRootFolderRequired field if non-nil, zero value otherwise.

### GetIsEntireRootFolderRequiredOk

`func (o *RestoreO365PublicFoldersParamsRootPublicFolder) GetIsEntireRootFolderRequiredOk() (*bool, bool)`

GetIsEntireRootFolderRequiredOk returns a tuple with the IsEntireRootFolderRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEntireRootFolderRequired

`func (o *RestoreO365PublicFoldersParamsRootPublicFolder) SetIsEntireRootFolderRequired(v bool)`

SetIsEntireRootFolderRequired sets IsEntireRootFolderRequired field to given value.

### HasIsEntireRootFolderRequired

`func (o *RestoreO365PublicFoldersParamsRootPublicFolder) HasIsEntireRootFolderRequired() bool`

HasIsEntireRootFolderRequired returns a boolean if a field has been set.

### SetIsEntireRootFolderRequiredNil

`func (o *RestoreO365PublicFoldersParamsRootPublicFolder) SetIsEntireRootFolderRequiredNil(b bool)`

 SetIsEntireRootFolderRequiredNil sets the value for IsEntireRootFolderRequired to be an explicit nil

### UnsetIsEntireRootFolderRequired
`func (o *RestoreO365PublicFoldersParamsRootPublicFolder) UnsetIsEntireRootFolderRequired()`

UnsetIsEntireRootFolderRequired ensures that no value is present for IsEntireRootFolderRequired, not even an explicit nil
### GetObject

`func (o *RestoreO365PublicFoldersParamsRootPublicFolder) GetObject() RestoreObject`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *RestoreO365PublicFoldersParamsRootPublicFolder) GetObjectOk() (*RestoreObject, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *RestoreO365PublicFoldersParamsRootPublicFolder) SetObject(v RestoreObject)`

SetObject sets Object field to given value.

### HasObject

`func (o *RestoreO365PublicFoldersParamsRootPublicFolder) HasObject() bool`

HasObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


