# OneDriveItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsFileItem** | Pointer to **NullableBool** | Specifies whether the current item is a file or not. | [optional] 
**ItemId** | Pointer to **NullableString** | Specifies the Id of the Drive item. | [optional] 
**ItemPath** | Pointer to **NullableString** | Specifies the path of the Drive item within the drive. | [optional] 

## Methods

### NewOneDriveItem

`func NewOneDriveItem() *OneDriveItem`

NewOneDriveItem instantiates a new OneDriveItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOneDriveItemWithDefaults

`func NewOneDriveItemWithDefaults() *OneDriveItem`

NewOneDriveItemWithDefaults instantiates a new OneDriveItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsFileItem

`func (o *OneDriveItem) GetIsFileItem() bool`

GetIsFileItem returns the IsFileItem field if non-nil, zero value otherwise.

### GetIsFileItemOk

`func (o *OneDriveItem) GetIsFileItemOk() (*bool, bool)`

GetIsFileItemOk returns a tuple with the IsFileItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFileItem

`func (o *OneDriveItem) SetIsFileItem(v bool)`

SetIsFileItem sets IsFileItem field to given value.

### HasIsFileItem

`func (o *OneDriveItem) HasIsFileItem() bool`

HasIsFileItem returns a boolean if a field has been set.

### SetIsFileItemNil

`func (o *OneDriveItem) SetIsFileItemNil(b bool)`

 SetIsFileItemNil sets the value for IsFileItem to be an explicit nil

### UnsetIsFileItem
`func (o *OneDriveItem) UnsetIsFileItem()`

UnsetIsFileItem ensures that no value is present for IsFileItem, not even an explicit nil
### GetItemId

`func (o *OneDriveItem) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *OneDriveItem) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *OneDriveItem) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *OneDriveItem) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### SetItemIdNil

`func (o *OneDriveItem) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *OneDriveItem) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
### GetItemPath

`func (o *OneDriveItem) GetItemPath() string`

GetItemPath returns the ItemPath field if non-nil, zero value otherwise.

### GetItemPathOk

`func (o *OneDriveItem) GetItemPathOk() (*string, bool)`

GetItemPathOk returns a tuple with the ItemPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemPath

`func (o *OneDriveItem) SetItemPath(v string)`

SetItemPath sets ItemPath field to given value.

### HasItemPath

`func (o *OneDriveItem) HasItemPath() bool`

HasItemPath returns a boolean if a field has been set.

### SetItemPathNil

`func (o *OneDriveItem) SetItemPathNil(b bool)`

 SetItemPathNil sets the value for ItemPath to be an explicit nil

### UnsetItemPath
`func (o *OneDriveItem) UnsetItemPath()`

UnsetItemPath ensures that no value is present for ItemPath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


