# ObjectMailboxParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OwnerInfo** | [**CommonRecoverObjectSnapshotParams**](CommonRecoverObjectSnapshotParams.md) | Specifies the Mailbox owner info. | 
**MailboxParams** | [**MailboxParam**](MailboxParam.md) | Specifies parameters to recover a Mailbox. | 

## Methods

### NewObjectMailboxParam

`func NewObjectMailboxParam(ownerInfo CommonRecoverObjectSnapshotParams, mailboxParams MailboxParam, ) *ObjectMailboxParam`

NewObjectMailboxParam instantiates a new ObjectMailboxParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectMailboxParamWithDefaults

`func NewObjectMailboxParamWithDefaults() *ObjectMailboxParam`

NewObjectMailboxParamWithDefaults instantiates a new ObjectMailboxParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwnerInfo

`func (o *ObjectMailboxParam) GetOwnerInfo() CommonRecoverObjectSnapshotParams`

GetOwnerInfo returns the OwnerInfo field if non-nil, zero value otherwise.

### GetOwnerInfoOk

`func (o *ObjectMailboxParam) GetOwnerInfoOk() (*CommonRecoverObjectSnapshotParams, bool)`

GetOwnerInfoOk returns a tuple with the OwnerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerInfo

`func (o *ObjectMailboxParam) SetOwnerInfo(v CommonRecoverObjectSnapshotParams)`

SetOwnerInfo sets OwnerInfo field to given value.


### GetMailboxParams

`func (o *ObjectMailboxParam) GetMailboxParams() MailboxParam`

GetMailboxParams returns the MailboxParams field if non-nil, zero value otherwise.

### GetMailboxParamsOk

`func (o *ObjectMailboxParam) GetMailboxParamsOk() (*MailboxParam, bool)`

GetMailboxParamsOk returns a tuple with the MailboxParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailboxParams

`func (o *ObjectMailboxParam) SetMailboxParams(v MailboxParam)`

SetMailboxParams sets MailboxParams field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


