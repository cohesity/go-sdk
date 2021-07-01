# OutlookMailbox

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MailboxObject** | Pointer to [**RestoreObjectDetails**](RestoreObjectDetails.md) |  | [optional] 
**OutlookFolderList** | Pointer to [**[]OutlookFolder**](OutlookFolder.md) | Specifies the list of folders to be restored incase user wishes not to restore entire mailbox. | [optional] 
**RestoreEntireMailbox** | Pointer to **NullableBool** | Specifies whether the entire mailbox is to be restored. | [optional] 

## Methods

### NewOutlookMailbox

`func NewOutlookMailbox() *OutlookMailbox`

NewOutlookMailbox instantiates a new OutlookMailbox object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutlookMailboxWithDefaults

`func NewOutlookMailboxWithDefaults() *OutlookMailbox`

NewOutlookMailboxWithDefaults instantiates a new OutlookMailbox object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMailboxObject

`func (o *OutlookMailbox) GetMailboxObject() RestoreObjectDetails`

GetMailboxObject returns the MailboxObject field if non-nil, zero value otherwise.

### GetMailboxObjectOk

`func (o *OutlookMailbox) GetMailboxObjectOk() (*RestoreObjectDetails, bool)`

GetMailboxObjectOk returns a tuple with the MailboxObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailboxObject

`func (o *OutlookMailbox) SetMailboxObject(v RestoreObjectDetails)`

SetMailboxObject sets MailboxObject field to given value.

### HasMailboxObject

`func (o *OutlookMailbox) HasMailboxObject() bool`

HasMailboxObject returns a boolean if a field has been set.

### GetOutlookFolderList

`func (o *OutlookMailbox) GetOutlookFolderList() []OutlookFolder`

GetOutlookFolderList returns the OutlookFolderList field if non-nil, zero value otherwise.

### GetOutlookFolderListOk

`func (o *OutlookMailbox) GetOutlookFolderListOk() (*[]OutlookFolder, bool)`

GetOutlookFolderListOk returns a tuple with the OutlookFolderList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutlookFolderList

`func (o *OutlookMailbox) SetOutlookFolderList(v []OutlookFolder)`

SetOutlookFolderList sets OutlookFolderList field to given value.

### HasOutlookFolderList

`func (o *OutlookMailbox) HasOutlookFolderList() bool`

HasOutlookFolderList returns a boolean if a field has been set.

### SetOutlookFolderListNil

`func (o *OutlookMailbox) SetOutlookFolderListNil(b bool)`

 SetOutlookFolderListNil sets the value for OutlookFolderList to be an explicit nil

### UnsetOutlookFolderList
`func (o *OutlookMailbox) UnsetOutlookFolderList()`

UnsetOutlookFolderList ensures that no value is present for OutlookFolderList, not even an explicit nil
### GetRestoreEntireMailbox

`func (o *OutlookMailbox) GetRestoreEntireMailbox() bool`

GetRestoreEntireMailbox returns the RestoreEntireMailbox field if non-nil, zero value otherwise.

### GetRestoreEntireMailboxOk

`func (o *OutlookMailbox) GetRestoreEntireMailboxOk() (*bool, bool)`

GetRestoreEntireMailboxOk returns a tuple with the RestoreEntireMailbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreEntireMailbox

`func (o *OutlookMailbox) SetRestoreEntireMailbox(v bool)`

SetRestoreEntireMailbox sets RestoreEntireMailbox field to given value.

### HasRestoreEntireMailbox

`func (o *OutlookMailbox) HasRestoreEntireMailbox() bool`

HasRestoreEntireMailbox returns a boolean if a field has been set.

### SetRestoreEntireMailboxNil

`func (o *OutlookMailbox) SetRestoreEntireMailboxNil(b bool)`

 SetRestoreEntireMailboxNil sets the value for RestoreEntireMailbox to be an explicit nil

### UnsetRestoreEntireMailbox
`func (o *OutlookMailbox) UnsetRestoreEntireMailbox()`

UnsetRestoreEntireMailbox ensures that no value is present for RestoreEntireMailbox, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


