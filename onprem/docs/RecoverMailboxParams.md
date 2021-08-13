# RecoverMailboxParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | [**[]ObjectMailboxParam**](ObjectMailboxParam.md) | Specifies a list of Mailbox params associated with the objects to recover. | 
**TargetMailbox** | Pointer to [**TargetMailboxParam**](TargetMailboxParam.md) | Specifies the target Mailbox to recover to. If not specified, the objects will be recovered to original location. | [optional] 
**ContinueOnError** | Pointer to **NullableBool** | Specifies whether to continue recovering other Mailboxes if one of Mailbox failed to recover. Default value is false. | [optional] 

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
### GetTargetMailbox

`func (o *RecoverMailboxParams) GetTargetMailbox() TargetMailboxParam`

GetTargetMailbox returns the TargetMailbox field if non-nil, zero value otherwise.

### GetTargetMailboxOk

`func (o *RecoverMailboxParams) GetTargetMailboxOk() (*TargetMailboxParam, bool)`

GetTargetMailboxOk returns a tuple with the TargetMailbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMailbox

`func (o *RecoverMailboxParams) SetTargetMailbox(v TargetMailboxParam)`

SetTargetMailbox sets TargetMailbox field to given value.

### HasTargetMailbox

`func (o *RecoverMailboxParams) HasTargetMailbox() bool`

HasTargetMailbox returns a boolean if a field has been set.

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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


