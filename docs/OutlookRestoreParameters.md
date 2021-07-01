# OutlookRestoreParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OutlookMailboxList** | Pointer to [**[]OutlookMailbox**](OutlookMailbox.md) | Specifies the list of mailboxes to be restored. | [optional] 
**TargetFolderPath** | Pointer to **NullableString** | Specifies the target folder path to restore the mailboxes. This will always be specified whether the target mailbox is the original or an alternate one. | [optional] 
**TargetMailbox** | Pointer to [**ProtectionSource**](ProtectionSource.md) |  | [optional] 

## Methods

### NewOutlookRestoreParameters

`func NewOutlookRestoreParameters() *OutlookRestoreParameters`

NewOutlookRestoreParameters instantiates a new OutlookRestoreParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutlookRestoreParametersWithDefaults

`func NewOutlookRestoreParametersWithDefaults() *OutlookRestoreParameters`

NewOutlookRestoreParametersWithDefaults instantiates a new OutlookRestoreParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOutlookMailboxList

`func (o *OutlookRestoreParameters) GetOutlookMailboxList() []OutlookMailbox`

GetOutlookMailboxList returns the OutlookMailboxList field if non-nil, zero value otherwise.

### GetOutlookMailboxListOk

`func (o *OutlookRestoreParameters) GetOutlookMailboxListOk() (*[]OutlookMailbox, bool)`

GetOutlookMailboxListOk returns a tuple with the OutlookMailboxList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutlookMailboxList

`func (o *OutlookRestoreParameters) SetOutlookMailboxList(v []OutlookMailbox)`

SetOutlookMailboxList sets OutlookMailboxList field to given value.

### HasOutlookMailboxList

`func (o *OutlookRestoreParameters) HasOutlookMailboxList() bool`

HasOutlookMailboxList returns a boolean if a field has been set.

### SetOutlookMailboxListNil

`func (o *OutlookRestoreParameters) SetOutlookMailboxListNil(b bool)`

 SetOutlookMailboxListNil sets the value for OutlookMailboxList to be an explicit nil

### UnsetOutlookMailboxList
`func (o *OutlookRestoreParameters) UnsetOutlookMailboxList()`

UnsetOutlookMailboxList ensures that no value is present for OutlookMailboxList, not even an explicit nil
### GetTargetFolderPath

`func (o *OutlookRestoreParameters) GetTargetFolderPath() string`

GetTargetFolderPath returns the TargetFolderPath field if non-nil, zero value otherwise.

### GetTargetFolderPathOk

`func (o *OutlookRestoreParameters) GetTargetFolderPathOk() (*string, bool)`

GetTargetFolderPathOk returns a tuple with the TargetFolderPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetFolderPath

`func (o *OutlookRestoreParameters) SetTargetFolderPath(v string)`

SetTargetFolderPath sets TargetFolderPath field to given value.

### HasTargetFolderPath

`func (o *OutlookRestoreParameters) HasTargetFolderPath() bool`

HasTargetFolderPath returns a boolean if a field has been set.

### SetTargetFolderPathNil

`func (o *OutlookRestoreParameters) SetTargetFolderPathNil(b bool)`

 SetTargetFolderPathNil sets the value for TargetFolderPath to be an explicit nil

### UnsetTargetFolderPath
`func (o *OutlookRestoreParameters) UnsetTargetFolderPath()`

UnsetTargetFolderPath ensures that no value is present for TargetFolderPath, not even an explicit nil
### GetTargetMailbox

`func (o *OutlookRestoreParameters) GetTargetMailbox() ProtectionSource`

GetTargetMailbox returns the TargetMailbox field if non-nil, zero value otherwise.

### GetTargetMailboxOk

`func (o *OutlookRestoreParameters) GetTargetMailboxOk() (*ProtectionSource, bool)`

GetTargetMailboxOk returns a tuple with the TargetMailbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMailbox

`func (o *OutlookRestoreParameters) SetTargetMailbox(v ProtectionSource)`

SetTargetMailbox sets TargetMailbox field to given value.

### HasTargetMailbox

`func (o *OutlookRestoreParameters) HasTargetMailbox() bool`

HasTargetMailbox returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


