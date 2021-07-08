# ScriptPathAndParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinueOnError** | Pointer to **NullableBool** | Applicable only for pre backup scripts. If this flag is set to true, then backup job will start even if the pre backup script fails. | [optional] 
**IsActive** | Pointer to **NullableBool** | Indicates if the script is active. If &#39;is_active&#39; is set to false, this script will not be executed even if it is part of the backup job. | [optional] 
**ScriptParams** | Pointer to **NullableString** | Custom parameters that users want to pass to the script. For example, if user wants to pass following params: 1. foo&#x3D;bar 2. v&#x3D;10. User can construct the param string as \&quot;far&#x3D;bar v&#x3D;10\&quot;. | [optional] 
**ScriptPath** | Pointer to **NullableString** | For backup jobs of type &#39;kPuppeteer&#39;, &#39;script_path&#39; is full path of location of the script within the host. For Pre/Post scripts of agent-based backup jobs, &#39;script_path&#39; is just name of the script, not full path. | [optional] 
**TimeoutSecs** | Pointer to **NullableInt32** | Timeout of the script. The script will be killed if it exceeds this value. &#39;-1&#39; indicates that the timeout is not set for the script. | [optional] 

## Methods

### NewScriptPathAndParams

`func NewScriptPathAndParams() *ScriptPathAndParams`

NewScriptPathAndParams instantiates a new ScriptPathAndParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptPathAndParamsWithDefaults

`func NewScriptPathAndParamsWithDefaults() *ScriptPathAndParams`

NewScriptPathAndParamsWithDefaults instantiates a new ScriptPathAndParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinueOnError

`func (o *ScriptPathAndParams) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *ScriptPathAndParams) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *ScriptPathAndParams) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *ScriptPathAndParams) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *ScriptPathAndParams) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *ScriptPathAndParams) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetIsActive

`func (o *ScriptPathAndParams) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ScriptPathAndParams) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ScriptPathAndParams) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *ScriptPathAndParams) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *ScriptPathAndParams) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *ScriptPathAndParams) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetScriptParams

`func (o *ScriptPathAndParams) GetScriptParams() string`

GetScriptParams returns the ScriptParams field if non-nil, zero value otherwise.

### GetScriptParamsOk

`func (o *ScriptPathAndParams) GetScriptParamsOk() (*string, bool)`

GetScriptParamsOk returns a tuple with the ScriptParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptParams

`func (o *ScriptPathAndParams) SetScriptParams(v string)`

SetScriptParams sets ScriptParams field to given value.

### HasScriptParams

`func (o *ScriptPathAndParams) HasScriptParams() bool`

HasScriptParams returns a boolean if a field has been set.

### SetScriptParamsNil

`func (o *ScriptPathAndParams) SetScriptParamsNil(b bool)`

 SetScriptParamsNil sets the value for ScriptParams to be an explicit nil

### UnsetScriptParams
`func (o *ScriptPathAndParams) UnsetScriptParams()`

UnsetScriptParams ensures that no value is present for ScriptParams, not even an explicit nil
### GetScriptPath

`func (o *ScriptPathAndParams) GetScriptPath() string`

GetScriptPath returns the ScriptPath field if non-nil, zero value otherwise.

### GetScriptPathOk

`func (o *ScriptPathAndParams) GetScriptPathOk() (*string, bool)`

GetScriptPathOk returns a tuple with the ScriptPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptPath

`func (o *ScriptPathAndParams) SetScriptPath(v string)`

SetScriptPath sets ScriptPath field to given value.

### HasScriptPath

`func (o *ScriptPathAndParams) HasScriptPath() bool`

HasScriptPath returns a boolean if a field has been set.

### SetScriptPathNil

`func (o *ScriptPathAndParams) SetScriptPathNil(b bool)`

 SetScriptPathNil sets the value for ScriptPath to be an explicit nil

### UnsetScriptPath
`func (o *ScriptPathAndParams) UnsetScriptPath()`

UnsetScriptPath ensures that no value is present for ScriptPath, not even an explicit nil
### GetTimeoutSecs

`func (o *ScriptPathAndParams) GetTimeoutSecs() int32`

GetTimeoutSecs returns the TimeoutSecs field if non-nil, zero value otherwise.

### GetTimeoutSecsOk

`func (o *ScriptPathAndParams) GetTimeoutSecsOk() (*int32, bool)`

GetTimeoutSecsOk returns a tuple with the TimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSecs

`func (o *ScriptPathAndParams) SetTimeoutSecs(v int32)`

SetTimeoutSecs sets TimeoutSecs field to given value.

### HasTimeoutSecs

`func (o *ScriptPathAndParams) HasTimeoutSecs() bool`

HasTimeoutSecs returns a boolean if a field has been set.

### SetTimeoutSecsNil

`func (o *ScriptPathAndParams) SetTimeoutSecsNil(b bool)`

 SetTimeoutSecsNil sets the value for TimeoutSecs to be an explicit nil

### UnsetTimeoutSecs
`func (o *ScriptPathAndParams) UnsetTimeoutSecs()`

UnsetTimeoutSecs ensures that no value is present for TimeoutSecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


