# OnPremDeployTargetResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | Pointer to **NullableString** | Target environment of the onprem deploy task. | [optional] 
**ErrorMessage** | Pointer to **NullableString** | Specifies the error message for onprem task failure. | [optional] 
**Message** | Pointer to **NullableString** | Message about the onprem deploy run. | [optional] 
**Status** | Pointer to **NullableString** | Status of the OnPrem deploy for a target. &#39;Running&#39; indicates that the run is still running. &#39;Canceled&#39; indicates that the run has been canceled. &#39;Canceling&#39; indicates that the run is in the process of being canceled. &#39;Paused&#39; indicates that the ongoing run has been paused. &#39;Failed&#39; indicates that the run has failed. &#39;Missed&#39; indicates that the run was unable to take place at the scheduled time because the previous run was still happening. &#39;Succeeded&#39; indicates that the run has finished successfully. &#39;SucceededWithWarning&#39; indicates that the run finished successfully, but there were some warning messages. &#39;Skipped&#39; indicates that the run was skipped. | [optional] 
**VmwareParams** | Pointer to [**OnPremDeployTargetResultVmwareParams**](OnPremDeployTargetResultVmwareParams.md) |  | [optional] 

## Methods

### NewOnPremDeployTargetResult

`func NewOnPremDeployTargetResult() *OnPremDeployTargetResult`

NewOnPremDeployTargetResult instantiates a new OnPremDeployTargetResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnPremDeployTargetResultWithDefaults

`func NewOnPremDeployTargetResultWithDefaults() *OnPremDeployTargetResult`

NewOnPremDeployTargetResultWithDefaults instantiates a new OnPremDeployTargetResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *OnPremDeployTargetResult) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *OnPremDeployTargetResult) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *OnPremDeployTargetResult) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *OnPremDeployTargetResult) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *OnPremDeployTargetResult) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *OnPremDeployTargetResult) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetErrorMessage

`func (o *OnPremDeployTargetResult) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *OnPremDeployTargetResult) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *OnPremDeployTargetResult) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *OnPremDeployTargetResult) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *OnPremDeployTargetResult) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *OnPremDeployTargetResult) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetMessage

`func (o *OnPremDeployTargetResult) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *OnPremDeployTargetResult) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *OnPremDeployTargetResult) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *OnPremDeployTargetResult) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *OnPremDeployTargetResult) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *OnPremDeployTargetResult) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetStatus

`func (o *OnPremDeployTargetResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OnPremDeployTargetResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OnPremDeployTargetResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OnPremDeployTargetResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *OnPremDeployTargetResult) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *OnPremDeployTargetResult) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetVmwareParams

`func (o *OnPremDeployTargetResult) GetVmwareParams() OnPremDeployTargetResultVmwareParams`

GetVmwareParams returns the VmwareParams field if non-nil, zero value otherwise.

### GetVmwareParamsOk

`func (o *OnPremDeployTargetResult) GetVmwareParamsOk() (*OnPremDeployTargetResultVmwareParams, bool)`

GetVmwareParamsOk returns a tuple with the VmwareParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareParams

`func (o *OnPremDeployTargetResult) SetVmwareParams(v OnPremDeployTargetResultVmwareParams)`

SetVmwareParams sets VmwareParams field to given value.

### HasVmwareParams

`func (o *OnPremDeployTargetResult) HasVmwareParams() bool`

HasVmwareParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


