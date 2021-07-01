# ChangeServiceStateResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **NullableString** | Specifies a description of the result of the operation. | [optional] 
**StatusUrl** | Pointer to **NullableString** | Specifies a URL which can be queried to check the status of the operation. | [optional] 

## Methods

### NewChangeServiceStateResult

`func NewChangeServiceStateResult() *ChangeServiceStateResult`

NewChangeServiceStateResult instantiates a new ChangeServiceStateResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeServiceStateResultWithDefaults

`func NewChangeServiceStateResultWithDefaults() *ChangeServiceStateResult`

NewChangeServiceStateResultWithDefaults instantiates a new ChangeServiceStateResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ChangeServiceStateResult) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ChangeServiceStateResult) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ChangeServiceStateResult) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ChangeServiceStateResult) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *ChangeServiceStateResult) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *ChangeServiceStateResult) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetStatusUrl

`func (o *ChangeServiceStateResult) GetStatusUrl() string`

GetStatusUrl returns the StatusUrl field if non-nil, zero value otherwise.

### GetStatusUrlOk

`func (o *ChangeServiceStateResult) GetStatusUrlOk() (*string, bool)`

GetStatusUrlOk returns a tuple with the StatusUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusUrl

`func (o *ChangeServiceStateResult) SetStatusUrl(v string)`

SetStatusUrl sets StatusUrl field to given value.

### HasStatusUrl

`func (o *ChangeServiceStateResult) HasStatusUrl() bool`

HasStatusUrl returns a boolean if a field has been set.

### SetStatusUrlNil

`func (o *ChangeServiceStateResult) SetStatusUrlNil(b bool)`

 SetStatusUrlNil sets the value for StatusUrl to be an explicit nil

### UnsetStatusUrl
`func (o *ChangeServiceStateResult) UnsetStatusUrl()`

UnsetStatusUrl ensures that no value is present for StatusUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


