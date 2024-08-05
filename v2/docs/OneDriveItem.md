# OneDriveItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Specifies the item id. | [optional] 
**IsFile** | Pointer to **NullableBool** | Specifies if the item is a file. | [optional] 
**ItemPath** | Pointer to **NullableString** | Specifies the path to the OneDrive item. | [optional] 

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

### GetId

`func (o *OneDriveItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OneDriveItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OneDriveItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OneDriveItem) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *OneDriveItem) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *OneDriveItem) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIsFile

`func (o *OneDriveItem) GetIsFile() bool`

GetIsFile returns the IsFile field if non-nil, zero value otherwise.

### GetIsFileOk

`func (o *OneDriveItem) GetIsFileOk() (*bool, bool)`

GetIsFileOk returns a tuple with the IsFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFile

`func (o *OneDriveItem) SetIsFile(v bool)`

SetIsFile sets IsFile field to given value.

### HasIsFile

`func (o *OneDriveItem) HasIsFile() bool`

HasIsFile returns a boolean if a field has been set.

### SetIsFileNil

`func (o *OneDriveItem) SetIsFileNil(b bool)`

 SetIsFileNil sets the value for IsFile to be an explicit nil

### UnsetIsFile
`func (o *OneDriveItem) UnsetIsFile()`

UnsetIsFile ensures that no value is present for IsFile, not even an explicit nil
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


