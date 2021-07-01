# RestoreOutlookParamsFolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FolderId** | Pointer to **NullableString** | The Unique ID of the folder. | [optional] 
**FolderKey** | Pointer to **NullableInt64** | The Unique key of the folder. | [optional] 
**IsEntireFolderRequired** | Pointer to **NullableBool** | Specify if the entire folder is to be restored. | [optional] 
**ItemIdVec** | Pointer to **[]string** | If is_entire_folder_required is set to false, user will then specify which particular items are to be restored. | [optional] 

## Methods

### NewRestoreOutlookParamsFolder

`func NewRestoreOutlookParamsFolder() *RestoreOutlookParamsFolder`

NewRestoreOutlookParamsFolder instantiates a new RestoreOutlookParamsFolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreOutlookParamsFolderWithDefaults

`func NewRestoreOutlookParamsFolderWithDefaults() *RestoreOutlookParamsFolder`

NewRestoreOutlookParamsFolderWithDefaults instantiates a new RestoreOutlookParamsFolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFolderId

`func (o *RestoreOutlookParamsFolder) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *RestoreOutlookParamsFolder) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *RestoreOutlookParamsFolder) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *RestoreOutlookParamsFolder) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### SetFolderIdNil

`func (o *RestoreOutlookParamsFolder) SetFolderIdNil(b bool)`

 SetFolderIdNil sets the value for FolderId to be an explicit nil

### UnsetFolderId
`func (o *RestoreOutlookParamsFolder) UnsetFolderId()`

UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
### GetFolderKey

`func (o *RestoreOutlookParamsFolder) GetFolderKey() int64`

GetFolderKey returns the FolderKey field if non-nil, zero value otherwise.

### GetFolderKeyOk

`func (o *RestoreOutlookParamsFolder) GetFolderKeyOk() (*int64, bool)`

GetFolderKeyOk returns a tuple with the FolderKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderKey

`func (o *RestoreOutlookParamsFolder) SetFolderKey(v int64)`

SetFolderKey sets FolderKey field to given value.

### HasFolderKey

`func (o *RestoreOutlookParamsFolder) HasFolderKey() bool`

HasFolderKey returns a boolean if a field has been set.

### SetFolderKeyNil

`func (o *RestoreOutlookParamsFolder) SetFolderKeyNil(b bool)`

 SetFolderKeyNil sets the value for FolderKey to be an explicit nil

### UnsetFolderKey
`func (o *RestoreOutlookParamsFolder) UnsetFolderKey()`

UnsetFolderKey ensures that no value is present for FolderKey, not even an explicit nil
### GetIsEntireFolderRequired

`func (o *RestoreOutlookParamsFolder) GetIsEntireFolderRequired() bool`

GetIsEntireFolderRequired returns the IsEntireFolderRequired field if non-nil, zero value otherwise.

### GetIsEntireFolderRequiredOk

`func (o *RestoreOutlookParamsFolder) GetIsEntireFolderRequiredOk() (*bool, bool)`

GetIsEntireFolderRequiredOk returns a tuple with the IsEntireFolderRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEntireFolderRequired

`func (o *RestoreOutlookParamsFolder) SetIsEntireFolderRequired(v bool)`

SetIsEntireFolderRequired sets IsEntireFolderRequired field to given value.

### HasIsEntireFolderRequired

`func (o *RestoreOutlookParamsFolder) HasIsEntireFolderRequired() bool`

HasIsEntireFolderRequired returns a boolean if a field has been set.

### SetIsEntireFolderRequiredNil

`func (o *RestoreOutlookParamsFolder) SetIsEntireFolderRequiredNil(b bool)`

 SetIsEntireFolderRequiredNil sets the value for IsEntireFolderRequired to be an explicit nil

### UnsetIsEntireFolderRequired
`func (o *RestoreOutlookParamsFolder) UnsetIsEntireFolderRequired()`

UnsetIsEntireFolderRequired ensures that no value is present for IsEntireFolderRequired, not even an explicit nil
### GetItemIdVec

`func (o *RestoreOutlookParamsFolder) GetItemIdVec() []string`

GetItemIdVec returns the ItemIdVec field if non-nil, zero value otherwise.

### GetItemIdVecOk

`func (o *RestoreOutlookParamsFolder) GetItemIdVecOk() (*[]string, bool)`

GetItemIdVecOk returns a tuple with the ItemIdVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemIdVec

`func (o *RestoreOutlookParamsFolder) SetItemIdVec(v []string)`

SetItemIdVec sets ItemIdVec field to given value.

### HasItemIdVec

`func (o *RestoreOutlookParamsFolder) HasItemIdVec() bool`

HasItemIdVec returns a boolean if a field has been set.

### SetItemIdVecNil

`func (o *RestoreOutlookParamsFolder) SetItemIdVecNil(b bool)`

 SetItemIdVecNil sets the value for ItemIdVec to be an explicit nil

### UnsetItemIdVec
`func (o *RestoreOutlookParamsFolder) UnsetItemIdVec()`

UnsetItemIdVec ensures that no value is present for ItemIdVec, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


