# RemoteScriptPathAndParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinueOnError** | Pointer to **NullableBool** | Specifies if the script needs to continue even if there is an occurence of an error. If this flag is set to true, then backup job will start even if the pre backup script fails. Applicable only for pre backup scripts. | [optional] 
**IsActive** | Pointer to **NullableBool** | Specifies if the script is active. If set to false, this script will not be executed even if it is part of the backup job. | [optional] 
**ScriptParams** | Pointer to **NullableString** | Specifies the parameters and values to pass into the remote script. For example if the script expects values for the &#39;database&#39; and &#39;user&#39; parameters, specify the parameters and values using the following string: \&quot;database&#x3D;myDatabase user&#x3D;me\&quot;. | [optional] 
**ScriptPath** | Pointer to **NullableString** | Specifies the path to the script on the remote host. | [optional] 
**TimeoutSecs** | Pointer to **NullableInt32** | Specifies the timeout of the script in seconds. The script will be killed if it exceeds this value. If the value of the field is &#39;-1&#39; then timeout is not set for the script. | [optional] 

## Methods

### NewRemoteScriptPathAndParams

`func NewRemoteScriptPathAndParams() *RemoteScriptPathAndParams`

NewRemoteScriptPathAndParams instantiates a new RemoteScriptPathAndParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteScriptPathAndParamsWithDefaults

`func NewRemoteScriptPathAndParamsWithDefaults() *RemoteScriptPathAndParams`

NewRemoteScriptPathAndParamsWithDefaults instantiates a new RemoteScriptPathAndParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinueOnError

`func (o *RemoteScriptPathAndParams) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *RemoteScriptPathAndParams) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *RemoteScriptPathAndParams) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *RemoteScriptPathAndParams) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *RemoteScriptPathAndParams) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *RemoteScriptPathAndParams) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetIsActive

`func (o *RemoteScriptPathAndParams) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *RemoteScriptPathAndParams) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *RemoteScriptPathAndParams) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *RemoteScriptPathAndParams) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *RemoteScriptPathAndParams) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *RemoteScriptPathAndParams) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetScriptParams

`func (o *RemoteScriptPathAndParams) GetScriptParams() string`

GetScriptParams returns the ScriptParams field if non-nil, zero value otherwise.

### GetScriptParamsOk

`func (o *RemoteScriptPathAndParams) GetScriptParamsOk() (*string, bool)`

GetScriptParamsOk returns a tuple with the ScriptParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptParams

`func (o *RemoteScriptPathAndParams) SetScriptParams(v string)`

SetScriptParams sets ScriptParams field to given value.

### HasScriptParams

`func (o *RemoteScriptPathAndParams) HasScriptParams() bool`

HasScriptParams returns a boolean if a field has been set.

### SetScriptParamsNil

`func (o *RemoteScriptPathAndParams) SetScriptParamsNil(b bool)`

 SetScriptParamsNil sets the value for ScriptParams to be an explicit nil

### UnsetScriptParams
`func (o *RemoteScriptPathAndParams) UnsetScriptParams()`

UnsetScriptParams ensures that no value is present for ScriptParams, not even an explicit nil
### GetScriptPath

`func (o *RemoteScriptPathAndParams) GetScriptPath() string`

GetScriptPath returns the ScriptPath field if non-nil, zero value otherwise.

### GetScriptPathOk

`func (o *RemoteScriptPathAndParams) GetScriptPathOk() (*string, bool)`

GetScriptPathOk returns a tuple with the ScriptPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptPath

`func (o *RemoteScriptPathAndParams) SetScriptPath(v string)`

SetScriptPath sets ScriptPath field to given value.

### HasScriptPath

`func (o *RemoteScriptPathAndParams) HasScriptPath() bool`

HasScriptPath returns a boolean if a field has been set.

### SetScriptPathNil

`func (o *RemoteScriptPathAndParams) SetScriptPathNil(b bool)`

 SetScriptPathNil sets the value for ScriptPath to be an explicit nil

### UnsetScriptPath
`func (o *RemoteScriptPathAndParams) UnsetScriptPath()`

UnsetScriptPath ensures that no value is present for ScriptPath, not even an explicit nil
### GetTimeoutSecs

`func (o *RemoteScriptPathAndParams) GetTimeoutSecs() int32`

GetTimeoutSecs returns the TimeoutSecs field if non-nil, zero value otherwise.

### GetTimeoutSecsOk

`func (o *RemoteScriptPathAndParams) GetTimeoutSecsOk() (*int32, bool)`

GetTimeoutSecsOk returns a tuple with the TimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSecs

`func (o *RemoteScriptPathAndParams) SetTimeoutSecs(v int32)`

SetTimeoutSecs sets TimeoutSecs field to given value.

### HasTimeoutSecs

`func (o *RemoteScriptPathAndParams) HasTimeoutSecs() bool`

HasTimeoutSecs returns a boolean if a field has been set.

### SetTimeoutSecsNil

`func (o *RemoteScriptPathAndParams) SetTimeoutSecsNil(b bool)`

 SetTimeoutSecsNil sets the value for TimeoutSecs to be an explicit nil

### UnsetTimeoutSecs
`func (o *RemoteScriptPathAndParams) UnsetTimeoutSecs()`

UnsetTimeoutSecs ensures that no value is present for TimeoutSecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


