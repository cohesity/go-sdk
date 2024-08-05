# DocumentObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsDirectory** | Pointer to **NullableBool** | Specifies whether the document is a directory. Since currently only files are supported this should always be false. | [optional] 
**ItemId** | **NullableString** | Specifies the item id of the document. | 

## Methods

### NewDocumentObject

`func NewDocumentObject(itemId NullableString, ) *DocumentObject`

NewDocumentObject instantiates a new DocumentObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentObjectWithDefaults

`func NewDocumentObjectWithDefaults() *DocumentObject`

NewDocumentObjectWithDefaults instantiates a new DocumentObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsDirectory

`func (o *DocumentObject) GetIsDirectory() bool`

GetIsDirectory returns the IsDirectory field if non-nil, zero value otherwise.

### GetIsDirectoryOk

`func (o *DocumentObject) GetIsDirectoryOk() (*bool, bool)`

GetIsDirectoryOk returns a tuple with the IsDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDirectory

`func (o *DocumentObject) SetIsDirectory(v bool)`

SetIsDirectory sets IsDirectory field to given value.

### HasIsDirectory

`func (o *DocumentObject) HasIsDirectory() bool`

HasIsDirectory returns a boolean if a field has been set.

### SetIsDirectoryNil

`func (o *DocumentObject) SetIsDirectoryNil(b bool)`

 SetIsDirectoryNil sets the value for IsDirectory to be an explicit nil

### UnsetIsDirectory
`func (o *DocumentObject) UnsetIsDirectory()`

UnsetIsDirectory ensures that no value is present for IsDirectory, not even an explicit nil
### GetItemId

`func (o *DocumentObject) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *DocumentObject) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *DocumentObject) SetItemId(v string)`

SetItemId sets ItemId field to given value.


### SetItemIdNil

`func (o *DocumentObject) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *DocumentObject) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


