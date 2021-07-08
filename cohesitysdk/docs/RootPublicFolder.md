# RootPublicFolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicFolderList** | Pointer to [**[]PublicFolder**](PublicFolder.md) | Specifies the list of Public Folders to be restored incase user wishes not to restore entire RootPublicFolder. | [optional] 
**RestoreEntireRootPublicFolder** | Pointer to **NullableBool** | Specifies whether the entire RootPublicFolder is to be restored. | [optional] 
**RootPublicFolderObject** | Pointer to [**RestoreObjectDetails**](RestoreObjectDetails.md) |  | [optional] 

## Methods

### NewRootPublicFolder

`func NewRootPublicFolder() *RootPublicFolder`

NewRootPublicFolder instantiates a new RootPublicFolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRootPublicFolderWithDefaults

`func NewRootPublicFolderWithDefaults() *RootPublicFolder`

NewRootPublicFolderWithDefaults instantiates a new RootPublicFolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicFolderList

`func (o *RootPublicFolder) GetPublicFolderList() []PublicFolder`

GetPublicFolderList returns the PublicFolderList field if non-nil, zero value otherwise.

### GetPublicFolderListOk

`func (o *RootPublicFolder) GetPublicFolderListOk() (*[]PublicFolder, bool)`

GetPublicFolderListOk returns a tuple with the PublicFolderList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicFolderList

`func (o *RootPublicFolder) SetPublicFolderList(v []PublicFolder)`

SetPublicFolderList sets PublicFolderList field to given value.

### HasPublicFolderList

`func (o *RootPublicFolder) HasPublicFolderList() bool`

HasPublicFolderList returns a boolean if a field has been set.

### SetPublicFolderListNil

`func (o *RootPublicFolder) SetPublicFolderListNil(b bool)`

 SetPublicFolderListNil sets the value for PublicFolderList to be an explicit nil

### UnsetPublicFolderList
`func (o *RootPublicFolder) UnsetPublicFolderList()`

UnsetPublicFolderList ensures that no value is present for PublicFolderList, not even an explicit nil
### GetRestoreEntireRootPublicFolder

`func (o *RootPublicFolder) GetRestoreEntireRootPublicFolder() bool`

GetRestoreEntireRootPublicFolder returns the RestoreEntireRootPublicFolder field if non-nil, zero value otherwise.

### GetRestoreEntireRootPublicFolderOk

`func (o *RootPublicFolder) GetRestoreEntireRootPublicFolderOk() (*bool, bool)`

GetRestoreEntireRootPublicFolderOk returns a tuple with the RestoreEntireRootPublicFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreEntireRootPublicFolder

`func (o *RootPublicFolder) SetRestoreEntireRootPublicFolder(v bool)`

SetRestoreEntireRootPublicFolder sets RestoreEntireRootPublicFolder field to given value.

### HasRestoreEntireRootPublicFolder

`func (o *RootPublicFolder) HasRestoreEntireRootPublicFolder() bool`

HasRestoreEntireRootPublicFolder returns a boolean if a field has been set.

### SetRestoreEntireRootPublicFolderNil

`func (o *RootPublicFolder) SetRestoreEntireRootPublicFolderNil(b bool)`

 SetRestoreEntireRootPublicFolderNil sets the value for RestoreEntireRootPublicFolder to be an explicit nil

### UnsetRestoreEntireRootPublicFolder
`func (o *RootPublicFolder) UnsetRestoreEntireRootPublicFolder()`

UnsetRestoreEntireRootPublicFolder ensures that no value is present for RestoreEntireRootPublicFolder, not even an explicit nil
### GetRootPublicFolderObject

`func (o *RootPublicFolder) GetRootPublicFolderObject() RestoreObjectDetails`

GetRootPublicFolderObject returns the RootPublicFolderObject field if non-nil, zero value otherwise.

### GetRootPublicFolderObjectOk

`func (o *RootPublicFolder) GetRootPublicFolderObjectOk() (*RestoreObjectDetails, bool)`

GetRootPublicFolderObjectOk returns a tuple with the RootPublicFolderObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPublicFolderObject

`func (o *RootPublicFolder) SetRootPublicFolderObject(v RestoreObjectDetails)`

SetRootPublicFolderObject sets RootPublicFolderObject field to given value.

### HasRootPublicFolderObject

`func (o *RootPublicFolder) HasRootPublicFolderObject() bool`

HasRootPublicFolderObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


