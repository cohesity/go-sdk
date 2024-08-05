# UpdatePasswordResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **NullableString** | Specifies the response message of password update. | [optional] 

## Methods

### NewUpdatePasswordResult

`func NewUpdatePasswordResult() *UpdatePasswordResult`

NewUpdatePasswordResult instantiates a new UpdatePasswordResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePasswordResultWithDefaults

`func NewUpdatePasswordResultWithDefaults() *UpdatePasswordResult`

NewUpdatePasswordResultWithDefaults instantiates a new UpdatePasswordResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *UpdatePasswordResult) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UpdatePasswordResult) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UpdatePasswordResult) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UpdatePasswordResult) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *UpdatePasswordResult) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *UpdatePasswordResult) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


