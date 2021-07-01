# ScriptExecutionStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**ErrorProto**](ErrorProto.md) |  | [optional] 
**Executing** | Pointer to **NullableBool** | Indicates if a script is executing. This is particularly useful when there is a cancellation request and Magneto crashes at that point before cleaning up the running script. | [optional] 
**ExitCode** | Pointer to **NullableInt32** | Exit code of the script. | [optional] 
**State** | Pointer to **NullableInt32** | Execution state of the script. | [optional] 

## Methods

### NewScriptExecutionStatus

`func NewScriptExecutionStatus() *ScriptExecutionStatus`

NewScriptExecutionStatus instantiates a new ScriptExecutionStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptExecutionStatusWithDefaults

`func NewScriptExecutionStatusWithDefaults() *ScriptExecutionStatus`

NewScriptExecutionStatusWithDefaults instantiates a new ScriptExecutionStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *ScriptExecutionStatus) GetError() ErrorProto`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ScriptExecutionStatus) GetErrorOk() (*ErrorProto, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ScriptExecutionStatus) SetError(v ErrorProto)`

SetError sets Error field to given value.

### HasError

`func (o *ScriptExecutionStatus) HasError() bool`

HasError returns a boolean if a field has been set.

### GetExecuting

`func (o *ScriptExecutionStatus) GetExecuting() bool`

GetExecuting returns the Executing field if non-nil, zero value otherwise.

### GetExecutingOk

`func (o *ScriptExecutionStatus) GetExecutingOk() (*bool, bool)`

GetExecutingOk returns a tuple with the Executing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecuting

`func (o *ScriptExecutionStatus) SetExecuting(v bool)`

SetExecuting sets Executing field to given value.

### HasExecuting

`func (o *ScriptExecutionStatus) HasExecuting() bool`

HasExecuting returns a boolean if a field has been set.

### SetExecutingNil

`func (o *ScriptExecutionStatus) SetExecutingNil(b bool)`

 SetExecutingNil sets the value for Executing to be an explicit nil

### UnsetExecuting
`func (o *ScriptExecutionStatus) UnsetExecuting()`

UnsetExecuting ensures that no value is present for Executing, not even an explicit nil
### GetExitCode

`func (o *ScriptExecutionStatus) GetExitCode() int32`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *ScriptExecutionStatus) GetExitCodeOk() (*int32, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *ScriptExecutionStatus) SetExitCode(v int32)`

SetExitCode sets ExitCode field to given value.

### HasExitCode

`func (o *ScriptExecutionStatus) HasExitCode() bool`

HasExitCode returns a boolean if a field has been set.

### SetExitCodeNil

`func (o *ScriptExecutionStatus) SetExitCodeNil(b bool)`

 SetExitCodeNil sets the value for ExitCode to be an explicit nil

### UnsetExitCode
`func (o *ScriptExecutionStatus) UnsetExitCode()`

UnsetExitCode ensures that no value is present for ExitCode, not even an explicit nil
### GetState

`func (o *ScriptExecutionStatus) GetState() int32`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ScriptExecutionStatus) GetStateOk() (*int32, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ScriptExecutionStatus) SetState(v int32)`

SetState sets State field to given value.

### HasState

`func (o *ScriptExecutionStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *ScriptExecutionStatus) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *ScriptExecutionStatus) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


