# CommonPrePostScriptParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsActive** | Pointer to **NullableBool** | Specifies whether the script should be enabled, default value set to true. | [optional] 
**Params** | Pointer to **NullableString** | Specifies the arguments or parameters and values to pass into the remote script. For example if the script expects values for the &#39;database&#39; and &#39;user&#39; parameters, specify the parameters and values using the following string: \&quot;database&#x3D;myDatabase user&#x3D;me\&quot;. | [optional] 
**Path** | **string** | Specifies the absolute path to the script on the remote host. | 
**TimeoutSecs** | Pointer to **NullableInt32** | Specifies the timeout of the script in seconds. The script will be killed if it exceeds this value. By default, no timeout will occur if left empty. | [optional] 

## Methods

### NewCommonPrePostScriptParams

`func NewCommonPrePostScriptParams(path string, ) *CommonPrePostScriptParams`

NewCommonPrePostScriptParams instantiates a new CommonPrePostScriptParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonPrePostScriptParamsWithDefaults

`func NewCommonPrePostScriptParamsWithDefaults() *CommonPrePostScriptParams`

NewCommonPrePostScriptParamsWithDefaults instantiates a new CommonPrePostScriptParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsActive

`func (o *CommonPrePostScriptParams) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *CommonPrePostScriptParams) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *CommonPrePostScriptParams) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *CommonPrePostScriptParams) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *CommonPrePostScriptParams) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *CommonPrePostScriptParams) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetParams

`func (o *CommonPrePostScriptParams) GetParams() string`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *CommonPrePostScriptParams) GetParamsOk() (*string, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *CommonPrePostScriptParams) SetParams(v string)`

SetParams sets Params field to given value.

### HasParams

`func (o *CommonPrePostScriptParams) HasParams() bool`

HasParams returns a boolean if a field has been set.

### SetParamsNil

`func (o *CommonPrePostScriptParams) SetParamsNil(b bool)`

 SetParamsNil sets the value for Params to be an explicit nil

### UnsetParams
`func (o *CommonPrePostScriptParams) UnsetParams()`

UnsetParams ensures that no value is present for Params, not even an explicit nil
### GetPath

`func (o *CommonPrePostScriptParams) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CommonPrePostScriptParams) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CommonPrePostScriptParams) SetPath(v string)`

SetPath sets Path field to given value.


### GetTimeoutSecs

`func (o *CommonPrePostScriptParams) GetTimeoutSecs() int32`

GetTimeoutSecs returns the TimeoutSecs field if non-nil, zero value otherwise.

### GetTimeoutSecsOk

`func (o *CommonPrePostScriptParams) GetTimeoutSecsOk() (*int32, bool)`

GetTimeoutSecsOk returns a tuple with the TimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSecs

`func (o *CommonPrePostScriptParams) SetTimeoutSecs(v int32)`

SetTimeoutSecs sets TimeoutSecs field to given value.

### HasTimeoutSecs

`func (o *CommonPrePostScriptParams) HasTimeoutSecs() bool`

HasTimeoutSecs returns a boolean if a field has been set.

### SetTimeoutSecsNil

`func (o *CommonPrePostScriptParams) SetTimeoutSecsNil(b bool)`

 SetTimeoutSecsNil sets the value for TimeoutSecs to be an explicit nil

### UnsetTimeoutSecs
`func (o *CommonPrePostScriptParams) UnsetTimeoutSecs()`

UnsetTimeoutSecs ensures that no value is present for TimeoutSecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


