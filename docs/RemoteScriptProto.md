# RemoteScriptProto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteHostParams** | Pointer to [**RemoteHostConnectorParams**](RemoteHostConnectorParams.md) |  | [optional] 
**Script** | Pointer to [**ScriptPathAndParams**](ScriptPathAndParams.md) |  | [optional] 
**Status** | Pointer to [**ScriptExecutionStatus**](ScriptExecutionStatus.md) |  | [optional] 

## Methods

### NewRemoteScriptProto

`func NewRemoteScriptProto() *RemoteScriptProto`

NewRemoteScriptProto instantiates a new RemoteScriptProto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteScriptProtoWithDefaults

`func NewRemoteScriptProtoWithDefaults() *RemoteScriptProto`

NewRemoteScriptProtoWithDefaults instantiates a new RemoteScriptProto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteHostParams

`func (o *RemoteScriptProto) GetRemoteHostParams() RemoteHostConnectorParams`

GetRemoteHostParams returns the RemoteHostParams field if non-nil, zero value otherwise.

### GetRemoteHostParamsOk

`func (o *RemoteScriptProto) GetRemoteHostParamsOk() (*RemoteHostConnectorParams, bool)`

GetRemoteHostParamsOk returns a tuple with the RemoteHostParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteHostParams

`func (o *RemoteScriptProto) SetRemoteHostParams(v RemoteHostConnectorParams)`

SetRemoteHostParams sets RemoteHostParams field to given value.

### HasRemoteHostParams

`func (o *RemoteScriptProto) HasRemoteHostParams() bool`

HasRemoteHostParams returns a boolean if a field has been set.

### GetScript

`func (o *RemoteScriptProto) GetScript() ScriptPathAndParams`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *RemoteScriptProto) GetScriptOk() (*ScriptPathAndParams, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *RemoteScriptProto) SetScript(v ScriptPathAndParams)`

SetScript sets Script field to given value.

### HasScript

`func (o *RemoteScriptProto) HasScript() bool`

HasScript returns a boolean if a field has been set.

### GetStatus

`func (o *RemoteScriptProto) GetStatus() ScriptExecutionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RemoteScriptProto) GetStatusOk() (*ScriptExecutionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RemoteScriptProto) SetStatus(v ScriptExecutionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RemoteScriptProto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


