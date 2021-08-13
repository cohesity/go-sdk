# ConstructMetaInfoResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | Pointer to **NullableString** | Specifies the environment type of the Protection group | [optional] 
**OracleParams** | Pointer to [**OracleRestoreMetaInfoResult**](OracleRestoreMetaInfoResult.md) | Specifies 3 Maps required to fill pfile text box. | [optional] 

## Methods

### NewConstructMetaInfoResult

`func NewConstructMetaInfoResult() *ConstructMetaInfoResult`

NewConstructMetaInfoResult instantiates a new ConstructMetaInfoResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConstructMetaInfoResultWithDefaults

`func NewConstructMetaInfoResultWithDefaults() *ConstructMetaInfoResult`

NewConstructMetaInfoResultWithDefaults instantiates a new ConstructMetaInfoResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *ConstructMetaInfoResult) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ConstructMetaInfoResult) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ConstructMetaInfoResult) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ConstructMetaInfoResult) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *ConstructMetaInfoResult) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *ConstructMetaInfoResult) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetOracleParams

`func (o *ConstructMetaInfoResult) GetOracleParams() OracleRestoreMetaInfoResult`

GetOracleParams returns the OracleParams field if non-nil, zero value otherwise.

### GetOracleParamsOk

`func (o *ConstructMetaInfoResult) GetOracleParamsOk() (*OracleRestoreMetaInfoResult, bool)`

GetOracleParamsOk returns a tuple with the OracleParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleParams

`func (o *ConstructMetaInfoResult) SetOracleParams(v OracleRestoreMetaInfoResult)`

SetOracleParams sets OracleParams field to given value.

### HasOracleParams

`func (o *ConstructMetaInfoResult) HasOracleParams() bool`

HasOracleParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


