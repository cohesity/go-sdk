# EndpointCheckResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckName** | Pointer to **NullableString** | Specifies the name of the check. | [optional] 
**Status** | Pointer to **NullableString** | Specifies the status of the check. | [optional] 
**ErrorMsg** | Pointer to **NullableString** | Specifies the error message to help troubleshoot. | [optional] 

## Methods

### NewEndpointCheckResult

`func NewEndpointCheckResult() *EndpointCheckResult`

NewEndpointCheckResult instantiates a new EndpointCheckResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointCheckResultWithDefaults

`func NewEndpointCheckResultWithDefaults() *EndpointCheckResult`

NewEndpointCheckResultWithDefaults instantiates a new EndpointCheckResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckName

`func (o *EndpointCheckResult) GetCheckName() string`

GetCheckName returns the CheckName field if non-nil, zero value otherwise.

### GetCheckNameOk

`func (o *EndpointCheckResult) GetCheckNameOk() (*string, bool)`

GetCheckNameOk returns a tuple with the CheckName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckName

`func (o *EndpointCheckResult) SetCheckName(v string)`

SetCheckName sets CheckName field to given value.

### HasCheckName

`func (o *EndpointCheckResult) HasCheckName() bool`

HasCheckName returns a boolean if a field has been set.

### SetCheckNameNil

`func (o *EndpointCheckResult) SetCheckNameNil(b bool)`

 SetCheckNameNil sets the value for CheckName to be an explicit nil

### UnsetCheckName
`func (o *EndpointCheckResult) UnsetCheckName()`

UnsetCheckName ensures that no value is present for CheckName, not even an explicit nil
### GetStatus

`func (o *EndpointCheckResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EndpointCheckResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EndpointCheckResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EndpointCheckResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *EndpointCheckResult) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *EndpointCheckResult) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetErrorMsg

`func (o *EndpointCheckResult) GetErrorMsg() string`

GetErrorMsg returns the ErrorMsg field if non-nil, zero value otherwise.

### GetErrorMsgOk

`func (o *EndpointCheckResult) GetErrorMsgOk() (*string, bool)`

GetErrorMsgOk returns a tuple with the ErrorMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMsg

`func (o *EndpointCheckResult) SetErrorMsg(v string)`

SetErrorMsg sets ErrorMsg field to given value.

### HasErrorMsg

`func (o *EndpointCheckResult) HasErrorMsg() bool`

HasErrorMsg returns a boolean if a field has been set.

### SetErrorMsgNil

`func (o *EndpointCheckResult) SetErrorMsgNil(b bool)`

 SetErrorMsgNil sets the value for ErrorMsg to be an explicit nil

### UnsetErrorMsg
`func (o *EndpointCheckResult) UnsetErrorMsg()`

UnsetErrorMsg ensures that no value is present for ErrorMsg, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


