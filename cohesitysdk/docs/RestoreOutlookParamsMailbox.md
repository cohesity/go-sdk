# RestoreOutlookParamsMailbox

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FolderVec** | Pointer to [**[]RestoreOutlookParamsFolder**](RestoreOutlookParamsFolder.md) | If is_entire_mailbox_required is set to false, user will then specify which particular folders are to be restored. | [optional] 
**IsEntireMailboxRequired** | Pointer to **NullableBool** | Specify if the entire mailbox is to be restored. | [optional] 
**Object** | Pointer to [**RestoreObject**](RestoreObject.md) |  | [optional] 

## Methods

### NewRestoreOutlookParamsMailbox

`func NewRestoreOutlookParamsMailbox() *RestoreOutlookParamsMailbox`

NewRestoreOutlookParamsMailbox instantiates a new RestoreOutlookParamsMailbox object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreOutlookParamsMailboxWithDefaults

`func NewRestoreOutlookParamsMailboxWithDefaults() *RestoreOutlookParamsMailbox`

NewRestoreOutlookParamsMailboxWithDefaults instantiates a new RestoreOutlookParamsMailbox object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFolderVec

`func (o *RestoreOutlookParamsMailbox) GetFolderVec() []RestoreOutlookParamsFolder`

GetFolderVec returns the FolderVec field if non-nil, zero value otherwise.

### GetFolderVecOk

`func (o *RestoreOutlookParamsMailbox) GetFolderVecOk() (*[]RestoreOutlookParamsFolder, bool)`

GetFolderVecOk returns a tuple with the FolderVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderVec

`func (o *RestoreOutlookParamsMailbox) SetFolderVec(v []RestoreOutlookParamsFolder)`

SetFolderVec sets FolderVec field to given value.

### HasFolderVec

`func (o *RestoreOutlookParamsMailbox) HasFolderVec() bool`

HasFolderVec returns a boolean if a field has been set.

### SetFolderVecNil

`func (o *RestoreOutlookParamsMailbox) SetFolderVecNil(b bool)`

 SetFolderVecNil sets the value for FolderVec to be an explicit nil

### UnsetFolderVec
`func (o *RestoreOutlookParamsMailbox) UnsetFolderVec()`

UnsetFolderVec ensures that no value is present for FolderVec, not even an explicit nil
### GetIsEntireMailboxRequired

`func (o *RestoreOutlookParamsMailbox) GetIsEntireMailboxRequired() bool`

GetIsEntireMailboxRequired returns the IsEntireMailboxRequired field if non-nil, zero value otherwise.

### GetIsEntireMailboxRequiredOk

`func (o *RestoreOutlookParamsMailbox) GetIsEntireMailboxRequiredOk() (*bool, bool)`

GetIsEntireMailboxRequiredOk returns a tuple with the IsEntireMailboxRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEntireMailboxRequired

`func (o *RestoreOutlookParamsMailbox) SetIsEntireMailboxRequired(v bool)`

SetIsEntireMailboxRequired sets IsEntireMailboxRequired field to given value.

### HasIsEntireMailboxRequired

`func (o *RestoreOutlookParamsMailbox) HasIsEntireMailboxRequired() bool`

HasIsEntireMailboxRequired returns a boolean if a field has been set.

### SetIsEntireMailboxRequiredNil

`func (o *RestoreOutlookParamsMailbox) SetIsEntireMailboxRequiredNil(b bool)`

 SetIsEntireMailboxRequiredNil sets the value for IsEntireMailboxRequired to be an explicit nil

### UnsetIsEntireMailboxRequired
`func (o *RestoreOutlookParamsMailbox) UnsetIsEntireMailboxRequired()`

UnsetIsEntireMailboxRequired ensures that no value is present for IsEntireMailboxRequired, not even an explicit nil
### GetObject

`func (o *RestoreOutlookParamsMailbox) GetObject() RestoreObject`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *RestoreOutlookParamsMailbox) GetObjectOk() (*RestoreObject, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *RestoreOutlookParamsMailbox) SetObject(v RestoreObject)`

SetObject sets Object field to given value.

### HasObject

`func (o *RestoreOutlookParamsMailbox) HasObject() bool`

HasObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


