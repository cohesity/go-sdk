# OutlookFolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FolderId** | Pointer to **NullableString** | Specifies the unique ID of the folder. | [optional] 
**FolderKey** | Pointer to **NullableInt64** | Specifies the key unique within the mailbox of the folder. | [optional] 
**OutlookItemIdList** | Pointer to **[]string** | Specifies the outlook items within the folder to be restored incase the user wishes not to restore the entire folder. | [optional] 
**RestoreEntireFolder** | Pointer to **NullableBool** | Specifies whether the entire folder is to be restored. | [optional] 

## Methods

### NewOutlookFolder

`func NewOutlookFolder() *OutlookFolder`

NewOutlookFolder instantiates a new OutlookFolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutlookFolderWithDefaults

`func NewOutlookFolderWithDefaults() *OutlookFolder`

NewOutlookFolderWithDefaults instantiates a new OutlookFolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFolderId

`func (o *OutlookFolder) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *OutlookFolder) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *OutlookFolder) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *OutlookFolder) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### SetFolderIdNil

`func (o *OutlookFolder) SetFolderIdNil(b bool)`

 SetFolderIdNil sets the value for FolderId to be an explicit nil

### UnsetFolderId
`func (o *OutlookFolder) UnsetFolderId()`

UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
### GetFolderKey

`func (o *OutlookFolder) GetFolderKey() int64`

GetFolderKey returns the FolderKey field if non-nil, zero value otherwise.

### GetFolderKeyOk

`func (o *OutlookFolder) GetFolderKeyOk() (*int64, bool)`

GetFolderKeyOk returns a tuple with the FolderKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderKey

`func (o *OutlookFolder) SetFolderKey(v int64)`

SetFolderKey sets FolderKey field to given value.

### HasFolderKey

`func (o *OutlookFolder) HasFolderKey() bool`

HasFolderKey returns a boolean if a field has been set.

### SetFolderKeyNil

`func (o *OutlookFolder) SetFolderKeyNil(b bool)`

 SetFolderKeyNil sets the value for FolderKey to be an explicit nil

### UnsetFolderKey
`func (o *OutlookFolder) UnsetFolderKey()`

UnsetFolderKey ensures that no value is present for FolderKey, not even an explicit nil
### GetOutlookItemIdList

`func (o *OutlookFolder) GetOutlookItemIdList() []string`

GetOutlookItemIdList returns the OutlookItemIdList field if non-nil, zero value otherwise.

### GetOutlookItemIdListOk

`func (o *OutlookFolder) GetOutlookItemIdListOk() (*[]string, bool)`

GetOutlookItemIdListOk returns a tuple with the OutlookItemIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutlookItemIdList

`func (o *OutlookFolder) SetOutlookItemIdList(v []string)`

SetOutlookItemIdList sets OutlookItemIdList field to given value.

### HasOutlookItemIdList

`func (o *OutlookFolder) HasOutlookItemIdList() bool`

HasOutlookItemIdList returns a boolean if a field has been set.

### SetOutlookItemIdListNil

`func (o *OutlookFolder) SetOutlookItemIdListNil(b bool)`

 SetOutlookItemIdListNil sets the value for OutlookItemIdList to be an explicit nil

### UnsetOutlookItemIdList
`func (o *OutlookFolder) UnsetOutlookItemIdList()`

UnsetOutlookItemIdList ensures that no value is present for OutlookItemIdList, not even an explicit nil
### GetRestoreEntireFolder

`func (o *OutlookFolder) GetRestoreEntireFolder() bool`

GetRestoreEntireFolder returns the RestoreEntireFolder field if non-nil, zero value otherwise.

### GetRestoreEntireFolderOk

`func (o *OutlookFolder) GetRestoreEntireFolderOk() (*bool, bool)`

GetRestoreEntireFolderOk returns a tuple with the RestoreEntireFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreEntireFolder

`func (o *OutlookFolder) SetRestoreEntireFolder(v bool)`

SetRestoreEntireFolder sets RestoreEntireFolder field to given value.

### HasRestoreEntireFolder

`func (o *OutlookFolder) HasRestoreEntireFolder() bool`

HasRestoreEntireFolder returns a boolean if a field has been set.

### SetRestoreEntireFolderNil

`func (o *OutlookFolder) SetRestoreEntireFolderNil(b bool)`

 SetRestoreEntireFolderNil sets the value for RestoreEntireFolder to be an explicit nil

### UnsetRestoreEntireFolder
`func (o *OutlookFolder) UnsetRestoreEntireFolder()`

UnsetRestoreEntireFolder ensures that no value is present for RestoreEntireFolder, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


