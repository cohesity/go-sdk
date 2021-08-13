# PublicFolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FolderId** | **NullableString** | Specifies the Unique ID of the folder. | 
**RecoverEntireFolder** | Pointer to **NullableBool** | Specifies whether to recover the whole PublicFolder. | [optional] 
**ItemIds** | Pointer to **[]string** | Specifies a list of item ids to recover. This field is applicable only if &#39;recoverEntireFolder&#39; is false. | [optional] 

## Methods

### NewPublicFolder

`func NewPublicFolder(folderId NullableString, ) *PublicFolder`

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


### SetFolderIdNil

`func (o *PublicFolder) SetFolderIdNil(b bool)`

 SetFolderIdNil sets the value for FolderId to be an explicit nil

### UnsetFolderId
`func (o *PublicFolder) UnsetFolderId()`

UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
### GetRecoverEntireFolder

`func (o *PublicFolder) GetRecoverEntireFolder() bool`

GetRecoverEntireFolder returns the RecoverEntireFolder field if non-nil, zero value otherwise.

### GetRecoverEntireFolderOk

`func (o *PublicFolder) GetRecoverEntireFolderOk() (*bool, bool)`

GetRecoverEntireFolderOk returns a tuple with the RecoverEntireFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverEntireFolder

`func (o *PublicFolder) SetRecoverEntireFolder(v bool)`

SetRecoverEntireFolder sets RecoverEntireFolder field to given value.

### HasRecoverEntireFolder

`func (o *PublicFolder) HasRecoverEntireFolder() bool`

HasRecoverEntireFolder returns a boolean if a field has been set.

### SetRecoverEntireFolderNil

`func (o *PublicFolder) SetRecoverEntireFolderNil(b bool)`

 SetRecoverEntireFolderNil sets the value for RecoverEntireFolder to be an explicit nil

### UnsetRecoverEntireFolder
`func (o *PublicFolder) UnsetRecoverEntireFolder()`

UnsetRecoverEntireFolder ensures that no value is present for RecoverEntireFolder, not even an explicit nil
### GetItemIds

`func (o *PublicFolder) GetItemIds() []string`

GetItemIds returns the ItemIds field if non-nil, zero value otherwise.

### GetItemIdsOk

`func (o *PublicFolder) GetItemIdsOk() (*[]string, bool)`

GetItemIdsOk returns a tuple with the ItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemIds

`func (o *PublicFolder) SetItemIds(v []string)`

SetItemIds sets ItemIds field to given value.

### HasItemIds

`func (o *PublicFolder) HasItemIds() bool`

HasItemIds returns a boolean if a field has been set.

### SetItemIdsNil

`func (o *PublicFolder) SetItemIdsNil(b bool)`

 SetItemIdsNil sets the value for ItemIds to be an explicit nil

### UnsetItemIds
`func (o *PublicFolder) UnsetItemIds()`

UnsetItemIds ensures that no value is present for ItemIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


