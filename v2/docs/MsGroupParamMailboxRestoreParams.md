# MsGroupParamMailboxRestoreParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecoverEntireMailbox** | Pointer to **NullableBool** | Specifies whether to recover the whole Mailbox. | [optional] 
**RecoverFolders** | Pointer to [**[]FolderItem**](FolderItem.md) | Specifies a list of email folders to recover. | [optional] 

## Methods

### NewMsGroupParamMailboxRestoreParams

`func NewMsGroupParamMailboxRestoreParams() *MsGroupParamMailboxRestoreParams`

NewMsGroupParamMailboxRestoreParams instantiates a new MsGroupParamMailboxRestoreParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsGroupParamMailboxRestoreParamsWithDefaults

`func NewMsGroupParamMailboxRestoreParamsWithDefaults() *MsGroupParamMailboxRestoreParams`

NewMsGroupParamMailboxRestoreParamsWithDefaults instantiates a new MsGroupParamMailboxRestoreParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecoverEntireMailbox

`func (o *MsGroupParamMailboxRestoreParams) GetRecoverEntireMailbox() bool`

GetRecoverEntireMailbox returns the RecoverEntireMailbox field if non-nil, zero value otherwise.

### GetRecoverEntireMailboxOk

`func (o *MsGroupParamMailboxRestoreParams) GetRecoverEntireMailboxOk() (*bool, bool)`

GetRecoverEntireMailboxOk returns a tuple with the RecoverEntireMailbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverEntireMailbox

`func (o *MsGroupParamMailboxRestoreParams) SetRecoverEntireMailbox(v bool)`

SetRecoverEntireMailbox sets RecoverEntireMailbox field to given value.

### HasRecoverEntireMailbox

`func (o *MsGroupParamMailboxRestoreParams) HasRecoverEntireMailbox() bool`

HasRecoverEntireMailbox returns a boolean if a field has been set.

### SetRecoverEntireMailboxNil

`func (o *MsGroupParamMailboxRestoreParams) SetRecoverEntireMailboxNil(b bool)`

 SetRecoverEntireMailboxNil sets the value for RecoverEntireMailbox to be an explicit nil

### UnsetRecoverEntireMailbox
`func (o *MsGroupParamMailboxRestoreParams) UnsetRecoverEntireMailbox()`

UnsetRecoverEntireMailbox ensures that no value is present for RecoverEntireMailbox, not even an explicit nil
### GetRecoverFolders

`func (o *MsGroupParamMailboxRestoreParams) GetRecoverFolders() []FolderItem`

GetRecoverFolders returns the RecoverFolders field if non-nil, zero value otherwise.

### GetRecoverFoldersOk

`func (o *MsGroupParamMailboxRestoreParams) GetRecoverFoldersOk() (*[]FolderItem, bool)`

GetRecoverFoldersOk returns a tuple with the RecoverFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverFolders

`func (o *MsGroupParamMailboxRestoreParams) SetRecoverFolders(v []FolderItem)`

SetRecoverFolders sets RecoverFolders field to given value.

### HasRecoverFolders

`func (o *MsGroupParamMailboxRestoreParams) HasRecoverFolders() bool`

HasRecoverFolders returns a boolean if a field has been set.

### SetRecoverFoldersNil

`func (o *MsGroupParamMailboxRestoreParams) SetRecoverFoldersNil(b bool)`

 SetRecoverFoldersNil sets the value for RecoverFolders to be an explicit nil

### UnsetRecoverFolders
`func (o *MsGroupParamMailboxRestoreParams) UnsetRecoverFolders()`

UnsetRecoverFolders ensures that no value is present for RecoverFolders, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


