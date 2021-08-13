# ConstructMetaInfoParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | **NullableString** | Specifies the environment type of the Protection group | 
**OracleParams** | Pointer to [**ConstructRestoreMetaInfoOracleParams**](ConstructRestoreMetaInfoOracleParams.md) | Oracle Params to construct meta info for alternate restore or clone. | [optional] 

## Methods

### NewConstructMetaInfoParams

`func NewConstructMetaInfoParams(environment NullableString, ) *ConstructMetaInfoParams`

NewConstructMetaInfoParams instantiates a new ConstructMetaInfoParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConstructMetaInfoParamsWithDefaults

`func NewConstructMetaInfoParamsWithDefaults() *ConstructMetaInfoParams`

NewConstructMetaInfoParamsWithDefaults instantiates a new ConstructMetaInfoParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *ConstructMetaInfoParams) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ConstructMetaInfoParams) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ConstructMetaInfoParams) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### SetEnvironmentNil

`func (o *ConstructMetaInfoParams) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *ConstructMetaInfoParams) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetOracleParams

`func (o *ConstructMetaInfoParams) GetOracleParams() ConstructRestoreMetaInfoOracleParams`

GetOracleParams returns the OracleParams field if non-nil, zero value otherwise.

### GetOracleParamsOk

`func (o *ConstructMetaInfoParams) GetOracleParamsOk() (*ConstructRestoreMetaInfoOracleParams, bool)`

GetOracleParamsOk returns a tuple with the OracleParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleParams

`func (o *ConstructMetaInfoParams) SetOracleParams(v ConstructRestoreMetaInfoOracleParams)`

SetOracleParams sets OracleParams field to given value.

### HasOracleParams

`func (o *ConstructMetaInfoParams) HasOracleParams() bool`

HasOracleParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


