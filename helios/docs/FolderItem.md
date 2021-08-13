# FolderItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **NullableInt64** | Specifies the email folder key. | 
**RecoverEntireFolder** | Pointer to **NullableBool** | Specifies whether to recover the whole email folder. | [optional] 
**ItemIds** | Pointer to **[]string** | Specifies a list of item ids to recover. This field is applicable only if &#39;recoverEntireFolder&#39; is false. | [optional] 

## Methods

### NewFolderItem

`func NewFolderItem(key NullableInt64, ) *FolderItem`

NewFolderItem instantiates a new FolderItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFolderItemWithDefaults

`func NewFolderItemWithDefaults() *FolderItem`

NewFolderItemWithDefaults instantiates a new FolderItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *FolderItem) GetKey() int64`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FolderItem) GetKeyOk() (*int64, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FolderItem) SetKey(v int64)`

SetKey sets Key field to given value.


### SetKeyNil

`func (o *FolderItem) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *FolderItem) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetRecoverEntireFolder

`func (o *FolderItem) GetRecoverEntireFolder() bool`

GetRecoverEntireFolder returns the RecoverEntireFolder field if non-nil, zero value otherwise.

### GetRecoverEntireFolderOk

`func (o *FolderItem) GetRecoverEntireFolderOk() (*bool, bool)`

GetRecoverEntireFolderOk returns a tuple with the RecoverEntireFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverEntireFolder

`func (o *FolderItem) SetRecoverEntireFolder(v bool)`

SetRecoverEntireFolder sets RecoverEntireFolder field to given value.

### HasRecoverEntireFolder

`func (o *FolderItem) HasRecoverEntireFolder() bool`

HasRecoverEntireFolder returns a boolean if a field has been set.

### SetRecoverEntireFolderNil

`func (o *FolderItem) SetRecoverEntireFolderNil(b bool)`

 SetRecoverEntireFolderNil sets the value for RecoverEntireFolder to be an explicit nil

### UnsetRecoverEntireFolder
`func (o *FolderItem) UnsetRecoverEntireFolder()`

UnsetRecoverEntireFolder ensures that no value is present for RecoverEntireFolder, not even an explicit nil
### GetItemIds

`func (o *FolderItem) GetItemIds() []string`

GetItemIds returns the ItemIds field if non-nil, zero value otherwise.

### GetItemIdsOk

`func (o *FolderItem) GetItemIdsOk() (*[]string, bool)`

GetItemIdsOk returns a tuple with the ItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemIds

`func (o *FolderItem) SetItemIds(v []string)`

SetItemIds sets ItemIds field to given value.

### HasItemIds

`func (o *FolderItem) HasItemIds() bool`

HasItemIds returns a boolean if a field has been set.

### SetItemIdsNil

`func (o *FolderItem) SetItemIdsNil(b bool)`

 SetItemIdsNil sets the value for ItemIds to be an explicit nil

### UnsetItemIds
`func (o *FolderItem) UnsetItemIds()`

UnsetItemIds ensures that no value is present for ItemIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


