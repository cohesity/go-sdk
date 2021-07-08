# SiteDriveItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsFileItem** | Pointer to **NullableBool** | Specifies whether the current item is a file or not. | [optional] 
**ItemId** | Pointer to **NullableString** | Specifies the Id of the Drive item. | [optional] 
**ItemPath** | Pointer to **NullableString** | Specifies the path of the Drive item within the drive. | [optional] 

## Methods

### NewSiteDriveItem

`func NewSiteDriveItem() *SiteDriveItem`

NewSiteDriveItem instantiates a new SiteDriveItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteDriveItemWithDefaults

`func NewSiteDriveItemWithDefaults() *SiteDriveItem`

NewSiteDriveItemWithDefaults instantiates a new SiteDriveItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsFileItem

`func (o *SiteDriveItem) GetIsFileItem() bool`

GetIsFileItem returns the IsFileItem field if non-nil, zero value otherwise.

### GetIsFileItemOk

`func (o *SiteDriveItem) GetIsFileItemOk() (*bool, bool)`

GetIsFileItemOk returns a tuple with the IsFileItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFileItem

`func (o *SiteDriveItem) SetIsFileItem(v bool)`

SetIsFileItem sets IsFileItem field to given value.

### HasIsFileItem

`func (o *SiteDriveItem) HasIsFileItem() bool`

HasIsFileItem returns a boolean if a field has been set.

### SetIsFileItemNil

`func (o *SiteDriveItem) SetIsFileItemNil(b bool)`

 SetIsFileItemNil sets the value for IsFileItem to be an explicit nil

### UnsetIsFileItem
`func (o *SiteDriveItem) UnsetIsFileItem()`

UnsetIsFileItem ensures that no value is present for IsFileItem, not even an explicit nil
### GetItemId

`func (o *SiteDriveItem) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *SiteDriveItem) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *SiteDriveItem) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *SiteDriveItem) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### SetItemIdNil

`func (o *SiteDriveItem) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *SiteDriveItem) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
### GetItemPath

`func (o *SiteDriveItem) GetItemPath() string`

GetItemPath returns the ItemPath field if non-nil, zero value otherwise.

### GetItemPathOk

`func (o *SiteDriveItem) GetItemPathOk() (*string, bool)`

GetItemPathOk returns a tuple with the ItemPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemPath

`func (o *SiteDriveItem) SetItemPath(v string)`

SetItemPath sets ItemPath field to given value.

### HasItemPath

`func (o *SiteDriveItem) HasItemPath() bool`

HasItemPath returns a boolean if a field has been set.

### SetItemPathNil

`func (o *SiteDriveItem) SetItemPathNil(b bool)`

 SetItemPathNil sets the value for ItemPath to be an explicit nil

### UnsetItemPath
`func (o *SiteDriveItem) UnsetItemPath()`

UnsetItemPath ensures that no value is present for ItemPath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


