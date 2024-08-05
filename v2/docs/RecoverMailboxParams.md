# RecoverMailboxParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinueOnError** | Pointer to **NullableBool** | Specifies whether to continue recovering other Mailboxes if one of Mailbox failed to recover. Default value is false. | [optional] 
**Objects** | [**[]ObjectMailboxParam**](ObjectMailboxParam.md) | Specifies a list of Mailbox params associated with the objects to recover. | 
**PstParams** | Pointer to [**RecoverMailboxParamsPstParams**](RecoverMailboxParamsPstParams.md) |  | [optional] 
**SkipRecoverArchiveMailbox** | Pointer to **NullableBool** | Specifies whether to skip the recovery of the archive mailbox and/or items present in the archive mailbox. Default value is true | [optional] 
**SkipRecoverArchiveRecoverableItems** | Pointer to **NullableBool** | Specifies whether to skip the recovery of the Archive Recoverable Items present in the selected snapshot. Default value is true | [optional] 
**SkipRecoverRecoverableItems** | Pointer to **NullableBool** | Specifies whether to skip the recovery of the Recoverable Items present in the selected snapshot. Default value is true | [optional] 
**TargetMailbox** | Pointer to [**RecoverMailboxParamsTargetMailbox**](RecoverMailboxParamsTargetMailbox.md) |  | [optional] 

## Methods

### NewRecoverMailboxParams

`func NewRecoverMailboxParams(objects []ObjectMailboxParam, ) *RecoverMailboxParams`

NewRecoverMailboxParams instantiates a new RecoverMailboxParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverMailboxParamsWithDefaults

`func NewRecoverMailboxParamsWithDefaults() *RecoverMailboxParams`

NewRecoverMailboxParamsWithDefaults instantiates a new RecoverMailboxParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinueOnError

`func (o *RecoverMailboxParams) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *RecoverMailboxParams) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *RecoverMailboxParams) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *RecoverMailboxParams) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *RecoverMailboxParams) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *RecoverMailboxParams) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetObjects

`func (o *RecoverMailboxParams) GetObjects() []ObjectMailboxParam`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *RecoverMailboxParams) GetObjectsOk() (*[]ObjectMailboxParam, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *RecoverMailboxParams) SetObjects(v []ObjectMailboxParam)`

SetObjects sets Objects field to given value.


### SetObjectsNil

`func (o *RecoverMailboxParams) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *RecoverMailboxParams) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetPstParams

`func (o *RecoverMailboxParams) GetPstParams() RecoverMailboxParamsPstParams`

GetPstParams returns the PstParams field if non-nil, zero value otherwise.

### GetPstParamsOk

`func (o *RecoverMailboxParams) GetPstParamsOk() (*RecoverMailboxParamsPstParams, bool)`

GetPstParamsOk returns a tuple with the PstParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPstParams

`func (o *RecoverMailboxParams) SetPstParams(v RecoverMailboxParamsPstParams)`

SetPstParams sets PstParams field to given value.

### HasPstParams

`func (o *RecoverMailboxParams) HasPstParams() bool`

HasPstParams returns a boolean if a field has been set.

### GetSkipRecoverArchiveMailbox

`func (o *RecoverMailboxParams) GetSkipRecoverArchiveMailbox() bool`

GetSkipRecoverArchiveMailbox returns the SkipRecoverArchiveMailbox field if non-nil, zero value otherwise.

### GetSkipRecoverArchiveMailboxOk

`func (o *RecoverMailboxParams) GetSkipRecoverArchiveMailboxOk() (*bool, bool)`

GetSkipRecoverArchiveMailboxOk returns a tuple with the SkipRecoverArchiveMailbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipRecoverArchiveMailbox

`func (o *RecoverMailboxParams) SetSkipRecoverArchiveMailbox(v bool)`

SetSkipRecoverArchiveMailbox sets SkipRecoverArchiveMailbox field to given value.

### HasSkipRecoverArchiveMailbox

`func (o *RecoverMailboxParams) HasSkipRecoverArchiveMailbox() bool`

HasSkipRecoverArchiveMailbox returns a boolean if a field has been set.

### SetSkipRecoverArchiveMailboxNil

`func (o *RecoverMailboxParams) SetSkipRecoverArchiveMailboxNil(b bool)`

 SetSkipRecoverArchiveMailboxNil sets the value for SkipRecoverArchiveMailbox to be an explicit nil

### UnsetSkipRecoverArchiveMailbox
`func (o *RecoverMailboxParams) UnsetSkipRecoverArchiveMailbox()`

UnsetSkipRecoverArchiveMailbox ensures that no value is present for SkipRecoverArchiveMailbox, not even an explicit nil
### GetSkipRecoverArchiveRecoverableItems

`func (o *RecoverMailboxParams) GetSkipRecoverArchiveRecoverableItems() bool`

GetSkipRecoverArchiveRecoverableItems returns the SkipRecoverArchiveRecoverableItems field if non-nil, zero value otherwise.

### GetSkipRecoverArchiveRecoverableItemsOk

`func (o *RecoverMailboxParams) GetSkipRecoverArchiveRecoverableItemsOk() (*bool, bool)`

GetSkipRecoverArchiveRecoverableItemsOk returns a tuple with the SkipRecoverArchiveRecoverableItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipRecoverArchiveRecoverableItems

`func (o *RecoverMailboxParams) SetSkipRecoverArchiveRecoverableItems(v bool)`

SetSkipRecoverArchiveRecoverableItems sets SkipRecoverArchiveRecoverableItems field to given value.

### HasSkipRecoverArchiveRecoverableItems

`func (o *RecoverMailboxParams) HasSkipRecoverArchiveRecoverableItems() bool`

HasSkipRecoverArchiveRecoverableItems returns a boolean if a field has been set.

### SetSkipRecoverArchiveRecoverableItemsNil

`func (o *RecoverMailboxParams) SetSkipRecoverArchiveRecoverableItemsNil(b bool)`

 SetSkipRecoverArchiveRecoverableItemsNil sets the value for SkipRecoverArchiveRecoverableItems to be an explicit nil

### UnsetSkipRecoverArchiveRecoverableItems
`func (o *RecoverMailboxParams) UnsetSkipRecoverArchiveRecoverableItems()`

UnsetSkipRecoverArchiveRecoverableItems ensures that no value is present for SkipRecoverArchiveRecoverableItems, not even an explicit nil
### GetSkipRecoverRecoverableItems

`func (o *RecoverMailboxParams) GetSkipRecoverRecoverableItems() bool`

GetSkipRecoverRecoverableItems returns the SkipRecoverRecoverableItems field if non-nil, zero value otherwise.

### GetSkipRecoverRecoverableItemsOk

`func (o *RecoverMailboxParams) GetSkipRecoverRecoverableItemsOk() (*bool, bool)`

GetSkipRecoverRecoverableItemsOk returns a tuple with the SkipRecoverRecoverableItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipRecoverRecoverableItems

`func (o *RecoverMailboxParams) SetSkipRecoverRecoverableItems(v bool)`

SetSkipRecoverRecoverableItems sets SkipRecoverRecoverableItems field to given value.

### HasSkipRecoverRecoverableItems

`func (o *RecoverMailboxParams) HasSkipRecoverRecoverableItems() bool`

HasSkipRecoverRecoverableItems returns a boolean if a field has been set.

### SetSkipRecoverRecoverableItemsNil

`func (o *RecoverMailboxParams) SetSkipRecoverRecoverableItemsNil(b bool)`

 SetSkipRecoverRecoverableItemsNil sets the value for SkipRecoverRecoverableItems to be an explicit nil

### UnsetSkipRecoverRecoverableItems
`func (o *RecoverMailboxParams) UnsetSkipRecoverRecoverableItems()`

UnsetSkipRecoverRecoverableItems ensures that no value is present for SkipRecoverRecoverableItems, not even an explicit nil
### GetTargetMailbox

`func (o *RecoverMailboxParams) GetTargetMailbox() RecoverMailboxParamsTargetMailbox`

GetTargetMailbox returns the TargetMailbox field if non-nil, zero value otherwise.

### GetTargetMailboxOk

`func (o *RecoverMailboxParams) GetTargetMailboxOk() (*RecoverMailboxParamsTargetMailbox, bool)`

GetTargetMailboxOk returns a tuple with the TargetMailbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMailbox

`func (o *RecoverMailboxParams) SetTargetMailbox(v RecoverMailboxParamsTargetMailbox)`

SetTargetMailbox sets TargetMailbox field to given value.

### HasTargetMailbox

`func (o *RecoverMailboxParams) HasTargetMailbox() bool`

HasTargetMailbox returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


