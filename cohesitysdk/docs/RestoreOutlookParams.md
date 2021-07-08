# RestoreOutlookParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MailboxVec** | Pointer to [**[]RestoreOutlookParamsMailbox**](RestoreOutlookParamsMailbox.md) | In a RestoreJob , user will provide the list of mailboxes to be restored. Provision is there for restoring full AND partial mailbox recovery. | [optional] 
**TargetFolderPath** | Pointer to **NullableString** |  | [optional] 
**TargetMailbox** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 

## Methods

### NewRestoreOutlookParams

`func NewRestoreOutlookParams() *RestoreOutlookParams`

NewRestoreOutlookParams instantiates a new RestoreOutlookParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreOutlookParamsWithDefaults

`func NewRestoreOutlookParamsWithDefaults() *RestoreOutlookParams`

NewRestoreOutlookParamsWithDefaults instantiates a new RestoreOutlookParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMailboxVec

`func (o *RestoreOutlookParams) GetMailboxVec() []RestoreOutlookParamsMailbox`

GetMailboxVec returns the MailboxVec field if non-nil, zero value otherwise.

### GetMailboxVecOk

`func (o *RestoreOutlookParams) GetMailboxVecOk() (*[]RestoreOutlookParamsMailbox, bool)`

GetMailboxVecOk returns a tuple with the MailboxVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailboxVec

`func (o *RestoreOutlookParams) SetMailboxVec(v []RestoreOutlookParamsMailbox)`

SetMailboxVec sets MailboxVec field to given value.

### HasMailboxVec

`func (o *RestoreOutlookParams) HasMailboxVec() bool`

HasMailboxVec returns a boolean if a field has been set.

### SetMailboxVecNil

`func (o *RestoreOutlookParams) SetMailboxVecNil(b bool)`

 SetMailboxVecNil sets the value for MailboxVec to be an explicit nil

### UnsetMailboxVec
`func (o *RestoreOutlookParams) UnsetMailboxVec()`

UnsetMailboxVec ensures that no value is present for MailboxVec, not even an explicit nil
### GetTargetFolderPath

`func (o *RestoreOutlookParams) GetTargetFolderPath() string`

GetTargetFolderPath returns the TargetFolderPath field if non-nil, zero value otherwise.

### GetTargetFolderPathOk

`func (o *RestoreOutlookParams) GetTargetFolderPathOk() (*string, bool)`

GetTargetFolderPathOk returns a tuple with the TargetFolderPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetFolderPath

`func (o *RestoreOutlookParams) SetTargetFolderPath(v string)`

SetTargetFolderPath sets TargetFolderPath field to given value.

### HasTargetFolderPath

`func (o *RestoreOutlookParams) HasTargetFolderPath() bool`

HasTargetFolderPath returns a boolean if a field has been set.

### SetTargetFolderPathNil

`func (o *RestoreOutlookParams) SetTargetFolderPathNil(b bool)`

 SetTargetFolderPathNil sets the value for TargetFolderPath to be an explicit nil

### UnsetTargetFolderPath
`func (o *RestoreOutlookParams) UnsetTargetFolderPath()`

UnsetTargetFolderPath ensures that no value is present for TargetFolderPath, not even an explicit nil
### GetTargetMailbox

`func (o *RestoreOutlookParams) GetTargetMailbox() EntityProto`

GetTargetMailbox returns the TargetMailbox field if non-nil, zero value otherwise.

### GetTargetMailboxOk

`func (o *RestoreOutlookParams) GetTargetMailboxOk() (*EntityProto, bool)`

GetTargetMailboxOk returns a tuple with the TargetMailbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMailbox

`func (o *RestoreOutlookParams) SetTargetMailbox(v EntityProto)`

SetTargetMailbox sets TargetMailbox field to given value.

### HasTargetMailbox

`func (o *RestoreOutlookParams) HasTargetMailbox() bool`

HasTargetMailbox returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


