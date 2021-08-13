# PublicFolderItemAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **NullableString** | Specifies the Public folder item type. | [optional] 
**Id** | Pointer to **NullableString** | Specifies the id of the indexed item. | [optional] 
**Subject** | Pointer to **NullableString** | Specifies the subject of the indexed item. | [optional] 
**HasAttachments** | Pointer to **NullableBool** | Specifies whether the item has any attachments | [optional] 
**ItemClass** | Pointer to **NullableString** | Specifies the item class of the indexed item. | [optional] 
**ReceivedTimeSecs** | Pointer to **NullableInt64** | Specifies the Unix timestamp epoch in seconds at which this item is received. | [optional] 
**ItemSize** | Pointer to **NullableInt64** | Specifies the size in bytes for the indexed item. | [optional] 
**ParentFolderId** | Pointer to **NullableString** | Specifies the id of parent folder the indexed item. | [optional] 

## Methods

### NewPublicFolderItemAllOf

`func NewPublicFolderItemAllOf() *PublicFolderItemAllOf`

NewPublicFolderItemAllOf instantiates a new PublicFolderItemAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicFolderItemAllOfWithDefaults

`func NewPublicFolderItemAllOfWithDefaults() *PublicFolderItemAllOf`

NewPublicFolderItemAllOfWithDefaults instantiates a new PublicFolderItemAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PublicFolderItemAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PublicFolderItemAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PublicFolderItemAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PublicFolderItemAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *PublicFolderItemAllOf) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *PublicFolderItemAllOf) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *PublicFolderItemAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicFolderItemAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicFolderItemAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PublicFolderItemAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *PublicFolderItemAllOf) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *PublicFolderItemAllOf) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetSubject

`func (o *PublicFolderItemAllOf) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *PublicFolderItemAllOf) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *PublicFolderItemAllOf) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *PublicFolderItemAllOf) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *PublicFolderItemAllOf) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *PublicFolderItemAllOf) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetHasAttachments

`func (o *PublicFolderItemAllOf) GetHasAttachments() bool`

GetHasAttachments returns the HasAttachments field if non-nil, zero value otherwise.

### GetHasAttachmentsOk

`func (o *PublicFolderItemAllOf) GetHasAttachmentsOk() (*bool, bool)`

GetHasAttachmentsOk returns a tuple with the HasAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAttachments

`func (o *PublicFolderItemAllOf) SetHasAttachments(v bool)`

SetHasAttachments sets HasAttachments field to given value.

### HasHasAttachments

`func (o *PublicFolderItemAllOf) HasHasAttachments() bool`

HasHasAttachments returns a boolean if a field has been set.

### SetHasAttachmentsNil

`func (o *PublicFolderItemAllOf) SetHasAttachmentsNil(b bool)`

 SetHasAttachmentsNil sets the value for HasAttachments to be an explicit nil

### UnsetHasAttachments
`func (o *PublicFolderItemAllOf) UnsetHasAttachments()`

UnsetHasAttachments ensures that no value is present for HasAttachments, not even an explicit nil
### GetItemClass

`func (o *PublicFolderItemAllOf) GetItemClass() string`

GetItemClass returns the ItemClass field if non-nil, zero value otherwise.

### GetItemClassOk

`func (o *PublicFolderItemAllOf) GetItemClassOk() (*string, bool)`

GetItemClassOk returns a tuple with the ItemClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemClass

`func (o *PublicFolderItemAllOf) SetItemClass(v string)`

SetItemClass sets ItemClass field to given value.

### HasItemClass

`func (o *PublicFolderItemAllOf) HasItemClass() bool`

HasItemClass returns a boolean if a field has been set.

### SetItemClassNil

`func (o *PublicFolderItemAllOf) SetItemClassNil(b bool)`

 SetItemClassNil sets the value for ItemClass to be an explicit nil

### UnsetItemClass
`func (o *PublicFolderItemAllOf) UnsetItemClass()`

UnsetItemClass ensures that no value is present for ItemClass, not even an explicit nil
### GetReceivedTimeSecs

`func (o *PublicFolderItemAllOf) GetReceivedTimeSecs() int64`

GetReceivedTimeSecs returns the ReceivedTimeSecs field if non-nil, zero value otherwise.

### GetReceivedTimeSecsOk

`func (o *PublicFolderItemAllOf) GetReceivedTimeSecsOk() (*int64, bool)`

GetReceivedTimeSecsOk returns a tuple with the ReceivedTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedTimeSecs

`func (o *PublicFolderItemAllOf) SetReceivedTimeSecs(v int64)`

SetReceivedTimeSecs sets ReceivedTimeSecs field to given value.

### HasReceivedTimeSecs

`func (o *PublicFolderItemAllOf) HasReceivedTimeSecs() bool`

HasReceivedTimeSecs returns a boolean if a field has been set.

### SetReceivedTimeSecsNil

`func (o *PublicFolderItemAllOf) SetReceivedTimeSecsNil(b bool)`

 SetReceivedTimeSecsNil sets the value for ReceivedTimeSecs to be an explicit nil

### UnsetReceivedTimeSecs
`func (o *PublicFolderItemAllOf) UnsetReceivedTimeSecs()`

UnsetReceivedTimeSecs ensures that no value is present for ReceivedTimeSecs, not even an explicit nil
### GetItemSize

`func (o *PublicFolderItemAllOf) GetItemSize() int64`

GetItemSize returns the ItemSize field if non-nil, zero value otherwise.

### GetItemSizeOk

`func (o *PublicFolderItemAllOf) GetItemSizeOk() (*int64, bool)`

GetItemSizeOk returns a tuple with the ItemSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemSize

`func (o *PublicFolderItemAllOf) SetItemSize(v int64)`

SetItemSize sets ItemSize field to given value.

### HasItemSize

`func (o *PublicFolderItemAllOf) HasItemSize() bool`

HasItemSize returns a boolean if a field has been set.

### SetItemSizeNil

`func (o *PublicFolderItemAllOf) SetItemSizeNil(b bool)`

 SetItemSizeNil sets the value for ItemSize to be an explicit nil

### UnsetItemSize
`func (o *PublicFolderItemAllOf) UnsetItemSize()`

UnsetItemSize ensures that no value is present for ItemSize, not even an explicit nil
### GetParentFolderId

`func (o *PublicFolderItemAllOf) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *PublicFolderItemAllOf) GetParentFolderIdOk() (*string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *PublicFolderItemAllOf) SetParentFolderId(v string)`

SetParentFolderId sets ParentFolderId field to given value.

### HasParentFolderId

`func (o *PublicFolderItemAllOf) HasParentFolderId() bool`

HasParentFolderId returns a boolean if a field has been set.

### SetParentFolderIdNil

`func (o *PublicFolderItemAllOf) SetParentFolderIdNil(b bool)`

 SetParentFolderIdNil sets the value for ParentFolderId to be an explicit nil

### UnsetParentFolderId
`func (o *PublicFolderItemAllOf) UnsetParentFolderId()`

UnsetParentFolderId ensures that no value is present for ParentFolderId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


