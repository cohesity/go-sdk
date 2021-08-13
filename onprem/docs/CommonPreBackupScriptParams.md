# CommonPreBackupScriptParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | Specifies the absolute path to the script on the remote host. | 
**Params** | Pointer to **NullableString** | Specifies the arguments or parameters and values to pass into the remote script. For example if the script expects values for the &#39;database&#39; and &#39;user&#39; parameters, specify the parameters and values using the following string: \&quot;database&#x3D;myDatabase user&#x3D;me\&quot;. | [optional] 
**TimeoutSecs** | Pointer to **NullableInt32** | Specifies the timeout of the script in seconds. The script will be killed if it exceeds this value. By default, no timeout will occur if left empty. | [optional] 
**IsActive** | Pointer to **NullableBool** | Specifies whether the script should be enabled, default value set to true. | [optional] 
**ContinueOnError** | Pointer to **NullableBool** | Specifies if the script needs to continue even if there is an occurence of an error. If this flag is set to true, then Backup Run will start even if the pre backup script fails. If not specified or false, then backup run will not start when script fails. | [optional] 

## Methods

### NewCommonPreBackupScriptParams

`func NewCommonPreBackupScriptParams(path string, ) *CommonPreBackupScriptParams`

NewCommonPreBackupScriptParams instantiates a new CommonPreBackupScriptParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonPreBackupScriptParamsWithDefaults

`func NewCommonPreBackupScriptParamsWithDefaults() *CommonPreBackupScriptParams`

NewCommonPreBackupScriptParamsWithDefaults instantiates a new CommonPreBackupScriptParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *CommonPreBackupScriptParams) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CommonPreBackupScriptParams) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CommonPreBackupScriptParams) SetPath(v string)`

SetPath sets Path field to given value.


### GetParams

`func (o *CommonPreBackupScriptParams) GetParams() string`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *CommonPreBackupScriptParams) GetParamsOk() (*string, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *CommonPreBackupScriptParams) SetParams(v string)`

SetParams sets Params field to given value.

### HasParams

`func (o *CommonPreBackupScriptParams) HasParams() bool`

HasParams returns a boolean if a field has been set.

### SetParamsNil

`func (o *CommonPreBackupScriptParams) SetParamsNil(b bool)`

 SetParamsNil sets the value for Params to be an explicit nil

### UnsetParams
`func (o *CommonPreBackupScriptParams) UnsetParams()`

UnsetParams ensures that no value is present for Params, not even an explicit nil
### GetTimeoutSecs

`func (o *CommonPreBackupScriptParams) GetTimeoutSecs() int32`

GetTimeoutSecs returns the TimeoutSecs field if non-nil, zero value otherwise.

### GetTimeoutSecsOk

`func (o *CommonPreBackupScriptParams) GetTimeoutSecsOk() (*int32, bool)`

GetTimeoutSecsOk returns a tuple with the TimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSecs

`func (o *CommonPreBackupScriptParams) SetTimeoutSecs(v int32)`

SetTimeoutSecs sets TimeoutSecs field to given value.

### HasTimeoutSecs

`func (o *CommonPreBackupScriptParams) HasTimeoutSecs() bool`

HasTimeoutSecs returns a boolean if a field has been set.

### SetTimeoutSecsNil

`func (o *CommonPreBackupScriptParams) SetTimeoutSecsNil(b bool)`

 SetTimeoutSecsNil sets the value for TimeoutSecs to be an explicit nil

### UnsetTimeoutSecs
`func (o *CommonPreBackupScriptParams) UnsetTimeoutSecs()`

UnsetTimeoutSecs ensures that no value is present for TimeoutSecs, not even an explicit nil
### GetIsActive

`func (o *CommonPreBackupScriptParams) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *CommonPreBackupScriptParams) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *CommonPreBackupScriptParams) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *CommonPreBackupScriptParams) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *CommonPreBackupScriptParams) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *CommonPreBackupScriptParams) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetContinueOnError

`func (o *CommonPreBackupScriptParams) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *CommonPreBackupScriptParams) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *CommonPreBackupScriptParams) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *CommonPreBackupScriptParams) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *CommonPreBackupScriptParams) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *CommonPreBackupScriptParams) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


