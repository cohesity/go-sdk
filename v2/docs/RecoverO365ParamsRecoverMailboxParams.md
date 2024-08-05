# RecoverO365ParamsRecoverMailboxParams

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

### NewRecoverO365ParamsRecoverMailboxParams

`func NewRecoverO365ParamsRecoverMailboxParams(objects []ObjectMailboxParam, ) *RecoverO365ParamsRecoverMailboxParams`

NewRecoverO365ParamsRecoverMailboxParams instantiates a new RecoverO365ParamsRecoverMailboxParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverO365ParamsRecoverMailboxParamsWithDefaults

`func NewRecoverO365ParamsRecoverMailboxParamsWithDefaults() *RecoverO365ParamsRecoverMailboxParams`

NewRecoverO365ParamsRecoverMailboxParamsWithDefaults instantiates a new RecoverO365ParamsRecoverMailboxParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinueOnError

`func (o *RecoverO365ParamsRecoverMailboxParams) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *RecoverO365ParamsRecoverMailboxParams) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *RecoverO365ParamsRecoverMailboxParams) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *RecoverO365ParamsRecoverMailboxParams) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *RecoverO365ParamsRecoverMailboxParams) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *RecoverO365ParamsRecoverMailboxParams) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetObjects

`func (o *RecoverO365ParamsRecoverMailboxParams) GetObjects() []ObjectMailboxParam`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *RecoverO365ParamsRecoverMailboxParams) GetObjectsOk() (*[]ObjectMailboxParam, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *RecoverO365ParamsRecoverMailboxParams) SetObjects(v []ObjectMailboxParam)`

SetObjects sets Objects field to given value.


### SetObjectsNil

`func (o *RecoverO365ParamsRecoverMailboxParams) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *RecoverO365ParamsRecoverMailboxParams) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetPstParams

`func (o *RecoverO365ParamsRecoverMailboxParams) GetPstParams() RecoverMailboxParamsPstParams`

GetPstParams returns the PstParams field if non-nil, zero value otherwise.

### GetPstParamsOk

`func (o *RecoverO365ParamsRecoverMailboxParams) GetPstParamsOk() (*RecoverMailboxParamsPstParams, bool)`

GetPstParamsOk returns a tuple with the PstParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPstParams

`func (o *RecoverO365ParamsRecoverMailboxParams) SetPstParams(v RecoverMailboxParamsPstParams)`

SetPstParams sets PstParams field to given value.

### HasPstParams

`func (o *RecoverO365ParamsRecoverMailboxParams) HasPstParams() bool`

HasPstParams returns a boolean if a field has been set.

### GetSkipRecoverArchiveMailbox

`func (o *RecoverO365ParamsRecoverMailboxParams) GetSkipRecoverArchiveMailbox() bool`

GetSkipRecoverArchiveMailbox returns the SkipRecoverArchiveMailbox field if non-nil, zero value otherwise.

### GetSkipRecoverArchiveMailboxOk

`func (o *RecoverO365ParamsRecoverMailboxParams) GetSkipRecoverArchiveMailboxOk() (*bool, bool)`

GetSkipRecoverArchiveMailboxOk returns a tuple with the SkipRecoverArchiveMailbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipRecoverArchiveMailbox

`func (o *RecoverO365ParamsRecoverMailboxParams) SetSkipRecoverArchiveMailbox(v bool)`

SetSkipRecoverArchiveMailbox sets SkipRecoverArchiveMailbox field to given value.

### HasSkipRecoverArchiveMailbox

`func (o *RecoverO365ParamsRecoverMailboxParams) HasSkipRecoverArchiveMailbox() bool`

HasSkipRecoverArchiveMailbox returns a boolean if a field has been set.

### SetSkipRecoverArchiveMailboxNil

`func (o *RecoverO365ParamsRecoverMailboxParams) SetSkipRecoverArchiveMailboxNil(b bool)`

 SetSkipRecoverArchiveMailboxNil sets the value for SkipRecoverArchiveMailbox to be an explicit nil

### UnsetSkipRecoverArchiveMailbox
`func (o *RecoverO365ParamsRecoverMailboxParams) UnsetSkipRecoverArchiveMailbox()`

UnsetSkipRecoverArchiveMailbox ensures that no value is present for SkipRecoverArchiveMailbox, not even an explicit nil
### GetSkipRecoverArchiveRecoverableItems

`func (o *RecoverO365ParamsRecoverMailboxParams) GetSkipRecoverArchiveRecoverableItems() bool`

GetSkipRecoverArchiveRecoverableItems returns the SkipRecoverArchiveRecoverableItems field if non-nil, zero value otherwise.

### GetSkipRecoverArchiveRecoverableItemsOk

`func (o *RecoverO365ParamsRecoverMailboxParams) GetSkipRecoverArchiveRecoverableItemsOk() (*bool, bool)`

GetSkipRecoverArchiveRecoverableItemsOk returns a tuple with the SkipRecoverArchiveRecoverableItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipRecoverArchiveRecoverableItems

`func (o *RecoverO365ParamsRecoverMailboxParams) SetSkipRecoverArchiveRecoverableItems(v bool)`

SetSkipRecoverArchiveRecoverableItems sets SkipRecoverArchiveRecoverableItems field to given value.

### HasSkipRecoverArchiveRecoverableItems

`func (o *RecoverO365ParamsRecoverMailboxParams) HasSkipRecoverArchiveRecoverableItems() bool`

HasSkipRecoverArchiveRecoverableItems returns a boolean if a field has been set.

### SetSkipRecoverArchiveRecoverableItemsNil

`func (o *RecoverO365ParamsRecoverMailboxParams) SetSkipRecoverArchiveRecoverableItemsNil(b bool)`

 SetSkipRecoverArchiveRecoverableItemsNil sets the value for SkipRecoverArchiveRecoverableItems to be an explicit nil

### UnsetSkipRecoverArchiveRecoverableItems
`func (o *RecoverO365ParamsRecoverMailboxParams) UnsetSkipRecoverArchiveRecoverableItems()`

UnsetSkipRecoverArchiveRecoverableItems ensures that no value is present for SkipRecoverArchiveRecoverableItems, not even an explicit nil
### GetSkipRecoverRecoverableItems

`func (o *RecoverO365ParamsRecoverMailboxParams) GetSkipRecoverRecoverableItems() bool`

GetSkipRecoverRecoverableItems returns the SkipRecoverRecoverableItems field if non-nil, zero value otherwise.

### GetSkipRecoverRecoverableItemsOk

`func (o *RecoverO365ParamsRecoverMailboxParams) GetSkipRecoverRecoverableItemsOk() (*bool, bool)`

GetSkipRecoverRecoverableItemsOk returns a tuple with the SkipRecoverRecoverableItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipRecoverRecoverableItems

`func (o *RecoverO365ParamsRecoverMailboxParams) SetSkipRecoverRecoverableItems(v bool)`

SetSkipRecoverRecoverableItems sets SkipRecoverRecoverableItems field to given value.

### HasSkipRecoverRecoverableItems

`func (o *RecoverO365ParamsRecoverMailboxParams) HasSkipRecoverRecoverableItems() bool`

HasSkipRecoverRecoverableItems returns a boolean if a field has been set.

### SetSkipRecoverRecoverableItemsNil

`func (o *RecoverO365ParamsRecoverMailboxParams) SetSkipRecoverRecoverableItemsNil(b bool)`

 SetSkipRecoverRecoverableItemsNil sets the value for SkipRecoverRecoverableItems to be an explicit nil

### UnsetSkipRecoverRecoverableItems
`func (o *RecoverO365ParamsRecoverMailboxParams) UnsetSkipRecoverRecoverableItems()`

UnsetSkipRecoverRecoverableItems ensures that no value is present for SkipRecoverRecoverableItems, not even an explicit nil
### GetTargetMailbox

`func (o *RecoverO365ParamsRecoverMailboxParams) GetTargetMailbox() RecoverMailboxParamsTargetMailbox`

GetTargetMailbox returns the TargetMailbox field if non-nil, zero value otherwise.

### GetTargetMailboxOk

`func (o *RecoverO365ParamsRecoverMailboxParams) GetTargetMailboxOk() (*RecoverMailboxParamsTargetMailbox, bool)`

GetTargetMailboxOk returns a tuple with the TargetMailbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMailbox

`func (o *RecoverO365ParamsRecoverMailboxParams) SetTargetMailbox(v RecoverMailboxParamsTargetMailbox)`

SetTargetMailbox sets TargetMailbox field to given value.

### HasTargetMailbox

`func (o *RecoverO365ParamsRecoverMailboxParams) HasTargetMailbox() bool`

HasTargetMailbox returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


