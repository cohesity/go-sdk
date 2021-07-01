# PublicFolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FolderId** | Pointer to **NullableString** | Specifies the unique ID of the folder. | [optional] 
**PublicFolderItemIdList** | Pointer to **[]string** | Specifies the PublicFolder items within the folder to be restored incase the user wishes not to restore the entire folder. | [optional] 
**RestoreEntireFolder** | Pointer to **NullableBool** | Specifies whether the entire folder is to be restored. | [optional] 

## Methods

### NewPublicFolder

`func NewPublicFolder() *PublicFolder`

NewPublicFolder instantiates a new PublicFolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicFolderWithDefaults

`func NewPublicFolderWithDefaults() *PublicFolder`

NewPublicFolderWithDefaults instantiates a new PublicFolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFolderId

`func (o *PublicFolder) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *PublicFolder) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *PublicFolder) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *PublicFolder) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### SetFolderIdNil

`func (o *PublicFolder) SetFolderIdNil(b bool)`

 SetFolderIdNil sets the value for FolderId to be an explicit nil

### UnsetFolderId
`func (o *PublicFolder) UnsetFolderId()`

UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
### GetPublicFolderItemIdList

`func (o *PublicFolder) GetPublicFolderItemIdList() []string`

GetPublicFolderItemIdList returns the PublicFolderItemIdList field if non-nil, zero value otherwise.

### GetPublicFolderItemIdListOk

`func (o *PublicFolder) GetPublicFolderItemIdListOk() (*[]string, bool)`

GetPublicFolderItemIdListOk returns a tuple with the PublicFolderItemIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicFolderItemIdList

`func (o *PublicFolder) SetPublicFolderItemIdList(v []string)`

SetPublicFolderItemIdList sets PublicFolderItemIdList field to given value.

### HasPublicFolderItemIdList

`func (o *PublicFolder) HasPublicFolderItemIdList() bool`

HasPublicFolderItemIdList returns a boolean if a field has been set.

### SetPublicFolderItemIdListNil

`func (o *PublicFolder) SetPublicFolderItemIdListNil(b bool)`

 SetPublicFolderItemIdListNil sets the value for PublicFolderItemIdList to be an explicit nil

### UnsetPublicFolderItemIdList
`func (o *PublicFolder) UnsetPublicFolderItemIdList()`

UnsetPublicFolderItemIdList ensures that no value is present for PublicFolderItemIdList, not even an explicit nil
### GetRestoreEntireFolder

`func (o *PublicFolder) GetRestoreEntireFolder() bool`

GetRestoreEntireFolder returns the RestoreEntireFolder field if non-nil, zero value otherwise.

### GetRestoreEntireFolderOk

`func (o *PublicFolder) GetRestoreEntireFolderOk() (*bool, bool)`

GetRestoreEntireFolderOk returns a tuple with the RestoreEntireFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreEntireFolder

`func (o *PublicFolder) SetRestoreEntireFolder(v bool)`

SetRestoreEntireFolder sets RestoreEntireFolder field to given value.

### HasRestoreEntireFolder

`func (o *PublicFolder) HasRestoreEntireFolder() bool`

HasRestoreEntireFolder returns a boolean if a field has been set.

### SetRestoreEntireFolderNil

`func (o *PublicFolder) SetRestoreEntireFolderNil(b bool)`

 SetRestoreEntireFolderNil sets the value for RestoreEntireFolder to be an explicit nil

### UnsetRestoreEntireFolder
`func (o *PublicFolder) UnsetRestoreEntireFolder()`

UnsetRestoreEntireFolder ensures that no value is present for RestoreEntireFolder, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


