# CreateBondResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **NullableString** | Specifies a message describing the result of the operation. | [optional] 

## Methods

### NewCreateBondResult

`func NewCreateBondResult() *CreateBondResult`

NewCreateBondResult instantiates a new CreateBondResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBondResultWithDefaults

`func NewCreateBondResultWithDefaults() *CreateBondResult`

NewCreateBondResultWithDefaults instantiates a new CreateBondResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *CreateBondResult) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateBondResult) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateBondResult) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CreateBondResult) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *CreateBondResult) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *CreateBondResult) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


