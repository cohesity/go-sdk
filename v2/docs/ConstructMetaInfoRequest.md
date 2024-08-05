# ConstructMetaInfoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | **NullableString** | Specifies the environment type of the Protection group | 
**OracleParams** | Pointer to [**ConstructMetaInfoParamsOracleParams**](ConstructMetaInfoParamsOracleParams.md) |  | [optional] 
**SfdcParams** | Pointer to [**ConstructMetaInfoParamsSfdcParams**](ConstructMetaInfoParamsSfdcParams.md) |  | [optional] 

## Methods

### NewConstructMetaInfoRequest

`func NewConstructMetaInfoRequest(environment NullableString, ) *ConstructMetaInfoRequest`

NewConstructMetaInfoRequest instantiates a new ConstructMetaInfoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConstructMetaInfoRequestWithDefaults

`func NewConstructMetaInfoRequestWithDefaults() *ConstructMetaInfoRequest`

NewConstructMetaInfoRequestWithDefaults instantiates a new ConstructMetaInfoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *ConstructMetaInfoRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ConstructMetaInfoRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ConstructMetaInfoRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### SetEnvironmentNil

`func (o *ConstructMetaInfoRequest) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *ConstructMetaInfoRequest) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetOracleParams

`func (o *ConstructMetaInfoRequest) GetOracleParams() ConstructMetaInfoParamsOracleParams`

GetOracleParams returns the OracleParams field if non-nil, zero value otherwise.

### GetOracleParamsOk

`func (o *ConstructMetaInfoRequest) GetOracleParamsOk() (*ConstructMetaInfoParamsOracleParams, bool)`

GetOracleParamsOk returns a tuple with the OracleParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleParams

`func (o *ConstructMetaInfoRequest) SetOracleParams(v ConstructMetaInfoParamsOracleParams)`

SetOracleParams sets OracleParams field to given value.

### HasOracleParams

`func (o *ConstructMetaInfoRequest) HasOracleParams() bool`

HasOracleParams returns a boolean if a field has been set.

### GetSfdcParams

`func (o *ConstructMetaInfoRequest) GetSfdcParams() ConstructMetaInfoParamsSfdcParams`

GetSfdcParams returns the SfdcParams field if non-nil, zero value otherwise.

### GetSfdcParamsOk

`func (o *ConstructMetaInfoRequest) GetSfdcParamsOk() (*ConstructMetaInfoParamsSfdcParams, bool)`

GetSfdcParamsOk returns a tuple with the SfdcParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfdcParams

`func (o *ConstructMetaInfoRequest) SetSfdcParams(v ConstructMetaInfoParamsSfdcParams)`

SetSfdcParams sets SfdcParams field to given value.

### HasSfdcParams

`func (o *ConstructMetaInfoRequest) HasSfdcParams() bool`

HasSfdcParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


