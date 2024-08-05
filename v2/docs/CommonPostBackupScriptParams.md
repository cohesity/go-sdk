# CommonPostBackupScriptParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsActive** | Pointer to **NullableBool** | Specifies whether the script should be enabled, default value set to true. | [optional] 
**Params** | Pointer to **NullableString** | Specifies the arguments or parameters and values to pass into the remote script. For example if the script expects values for the &#39;database&#39; and &#39;user&#39; parameters, specify the parameters and values using the following string: \&quot;database&#x3D;myDatabase user&#x3D;me\&quot;. | [optional] 
**Path** | **string** | Specifies the absolute path to the script on the remote host. | 
**TimeoutSecs** | Pointer to **NullableInt32** | Specifies the timeout of the script in seconds. The script will be killed if it exceeds this value. By default, no timeout will occur if left empty. | [optional] 

## Methods

### NewCommonPostBackupScriptParams

`func NewCommonPostBackupScriptParams(path string, ) *CommonPostBackupScriptParams`

NewCommonPostBackupScriptParams instantiates a new CommonPostBackupScriptParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonPostBackupScriptParamsWithDefaults

`func NewCommonPostBackupScriptParamsWithDefaults() *CommonPostBackupScriptParams`

NewCommonPostBackupScriptParamsWithDefaults instantiates a new CommonPostBackupScriptParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsActive

`func (o *CommonPostBackupScriptParams) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *CommonPostBackupScriptParams) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *CommonPostBackupScriptParams) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *CommonPostBackupScriptParams) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *CommonPostBackupScriptParams) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *CommonPostBackupScriptParams) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetParams

`func (o *CommonPostBackupScriptParams) GetParams() string`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *CommonPostBackupScriptParams) GetParamsOk() (*string, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *CommonPostBackupScriptParams) SetParams(v string)`

SetParams sets Params field to given value.

### HasParams

`func (o *CommonPostBackupScriptParams) HasParams() bool`

HasParams returns a boolean if a field has been set.

### SetParamsNil

`func (o *CommonPostBackupScriptParams) SetParamsNil(b bool)`

 SetParamsNil sets the value for Params to be an explicit nil

### UnsetParams
`func (o *CommonPostBackupScriptParams) UnsetParams()`

UnsetParams ensures that no value is present for Params, not even an explicit nil
### GetPath

`func (o *CommonPostBackupScriptParams) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CommonPostBackupScriptParams) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CommonPostBackupScriptParams) SetPath(v string)`

SetPath sets Path field to given value.


### GetTimeoutSecs

`func (o *CommonPostBackupScriptParams) GetTimeoutSecs() int32`

GetTimeoutSecs returns the TimeoutSecs field if non-nil, zero value otherwise.

### GetTimeoutSecsOk

`func (o *CommonPostBackupScriptParams) GetTimeoutSecsOk() (*int32, bool)`

GetTimeoutSecsOk returns a tuple with the TimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSecs

`func (o *CommonPostBackupScriptParams) SetTimeoutSecs(v int32)`

SetTimeoutSecs sets TimeoutSecs field to given value.

### HasTimeoutSecs

`func (o *CommonPostBackupScriptParams) HasTimeoutSecs() bool`

HasTimeoutSecs returns a boolean if a field has been set.

### SetTimeoutSecsNil

`func (o *CommonPostBackupScriptParams) SetTimeoutSecsNil(b bool)`

 SetTimeoutSecsNil sets the value for TimeoutSecs to be an explicit nil

### UnsetTimeoutSecs
`func (o *CommonPostBackupScriptParams) UnsetTimeoutSecs()`

UnsetTimeoutSecs ensures that no value is present for TimeoutSecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


