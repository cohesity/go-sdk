# RunDiagnosticsMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **NullableString** | Specifies the status message returned after initiating a run diagnostics request. | [optional] 

## Methods

### NewRunDiagnosticsMessage

`func NewRunDiagnosticsMessage() *RunDiagnosticsMessage`

NewRunDiagnosticsMessage instantiates a new RunDiagnosticsMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunDiagnosticsMessageWithDefaults

`func NewRunDiagnosticsMessageWithDefaults() *RunDiagnosticsMessage`

NewRunDiagnosticsMessageWithDefaults instantiates a new RunDiagnosticsMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *RunDiagnosticsMessage) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RunDiagnosticsMessage) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RunDiagnosticsMessage) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *RunDiagnosticsMessage) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *RunDiagnosticsMessage) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *RunDiagnosticsMessage) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


