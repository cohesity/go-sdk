# ObjectMailboxParamMailboxParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecoverEntireMailbox** | Pointer to **NullableBool** | Specifies whether to recover the whole Mailbox. | [optional] 
**RecoverFolders** | Pointer to [**[]FolderItem**](FolderItem.md) | Specifies a list of email folders to recover. | [optional] 

## Methods

### NewObjectMailboxParamMailboxParams

`func NewObjectMailboxParamMailboxParams() *ObjectMailboxParamMailboxParams`

NewObjectMailboxParamMailboxParams instantiates a new ObjectMailboxParamMailboxParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectMailboxParamMailboxParamsWithDefaults

`func NewObjectMailboxParamMailboxParamsWithDefaults() *ObjectMailboxParamMailboxParams`

NewObjectMailboxParamMailboxParamsWithDefaults instantiates a new ObjectMailboxParamMailboxParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecoverEntireMailbox

`func (o *ObjectMailboxParamMailboxParams) GetRecoverEntireMailbox() bool`

GetRecoverEntireMailbox returns the RecoverEntireMailbox field if non-nil, zero value otherwise.

### GetRecoverEntireMailboxOk

`func (o *ObjectMailboxParamMailboxParams) GetRecoverEntireMailboxOk() (*bool, bool)`

GetRecoverEntireMailboxOk returns a tuple with the RecoverEntireMailbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverEntireMailbox

`func (o *ObjectMailboxParamMailboxParams) SetRecoverEntireMailbox(v bool)`

SetRecoverEntireMailbox sets RecoverEntireMailbox field to given value.

### HasRecoverEntireMailbox

`func (o *ObjectMailboxParamMailboxParams) HasRecoverEntireMailbox() bool`

HasRecoverEntireMailbox returns a boolean if a field has been set.

### SetRecoverEntireMailboxNil

`func (o *ObjectMailboxParamMailboxParams) SetRecoverEntireMailboxNil(b bool)`

 SetRecoverEntireMailboxNil sets the value for RecoverEntireMailbox to be an explicit nil

### UnsetRecoverEntireMailbox
`func (o *ObjectMailboxParamMailboxParams) UnsetRecoverEntireMailbox()`

UnsetRecoverEntireMailbox ensures that no value is present for RecoverEntireMailbox, not even an explicit nil
### GetRecoverFolders

`func (o *ObjectMailboxParamMailboxParams) GetRecoverFolders() []FolderItem`

GetRecoverFolders returns the RecoverFolders field if non-nil, zero value otherwise.

### GetRecoverFoldersOk

`func (o *ObjectMailboxParamMailboxParams) GetRecoverFoldersOk() (*[]FolderItem, bool)`

GetRecoverFoldersOk returns a tuple with the RecoverFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverFolders

`func (o *ObjectMailboxParamMailboxParams) SetRecoverFolders(v []FolderItem)`

SetRecoverFolders sets RecoverFolders field to given value.

### HasRecoverFolders

`func (o *ObjectMailboxParamMailboxParams) HasRecoverFolders() bool`

HasRecoverFolders returns a boolean if a field has been set.

### SetRecoverFoldersNil

`func (o *ObjectMailboxParamMailboxParams) SetRecoverFoldersNil(b bool)`

 SetRecoverFoldersNil sets the value for RecoverFolders to be an explicit nil

### UnsetRecoverFolders
`func (o *ObjectMailboxParamMailboxParams) UnsetRecoverFolders()`

UnsetRecoverFolders ensures that no value is present for RecoverFolders, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


