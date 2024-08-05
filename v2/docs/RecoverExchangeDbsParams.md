# RecoverExchangeDbsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExchangeTargetParams** | Pointer to [**NullableRecoverExchangeDbsParamsExchangeTargetParams**](RecoverExchangeDbsParamsExchangeTargetParams.md) |  | [optional] 
**RecoverToNewSource** | **bool** | Specifies the parameter whether the recovery should be performed to a new or an existing target. | 
**TargetEnvironment** | **string** | Specifies the environment of the recovery target. The corresponding params below must be filled out. | 

## Methods

### NewRecoverExchangeDbsParams

`func NewRecoverExchangeDbsParams(recoverToNewSource bool, targetEnvironment string, ) *RecoverExchangeDbsParams`

NewRecoverExchangeDbsParams instantiates a new RecoverExchangeDbsParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverExchangeDbsParamsWithDefaults

`func NewRecoverExchangeDbsParamsWithDefaults() *RecoverExchangeDbsParams`

NewRecoverExchangeDbsParamsWithDefaults instantiates a new RecoverExchangeDbsParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExchangeTargetParams

`func (o *RecoverExchangeDbsParams) GetExchangeTargetParams() RecoverExchangeDbsParamsExchangeTargetParams`

GetExchangeTargetParams returns the ExchangeTargetParams field if non-nil, zero value otherwise.

### GetExchangeTargetParamsOk

`func (o *RecoverExchangeDbsParams) GetExchangeTargetParamsOk() (*RecoverExchangeDbsParamsExchangeTargetParams, bool)`

GetExchangeTargetParamsOk returns a tuple with the ExchangeTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeTargetParams

`func (o *RecoverExchangeDbsParams) SetExchangeTargetParams(v RecoverExchangeDbsParamsExchangeTargetParams)`

SetExchangeTargetParams sets ExchangeTargetParams field to given value.

### HasExchangeTargetParams

`func (o *RecoverExchangeDbsParams) HasExchangeTargetParams() bool`

HasExchangeTargetParams returns a boolean if a field has been set.

### SetExchangeTargetParamsNil

`func (o *RecoverExchangeDbsParams) SetExchangeTargetParamsNil(b bool)`

 SetExchangeTargetParamsNil sets the value for ExchangeTargetParams to be an explicit nil

### UnsetExchangeTargetParams
`func (o *RecoverExchangeDbsParams) UnsetExchangeTargetParams()`

UnsetExchangeTargetParams ensures that no value is present for ExchangeTargetParams, not even an explicit nil
### GetRecoverToNewSource

`func (o *RecoverExchangeDbsParams) GetRecoverToNewSource() bool`

GetRecoverToNewSource returns the RecoverToNewSource field if non-nil, zero value otherwise.

### GetRecoverToNewSourceOk

`func (o *RecoverExchangeDbsParams) GetRecoverToNewSourceOk() (*bool, bool)`

GetRecoverToNewSourceOk returns a tuple with the RecoverToNewSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverToNewSource

`func (o *RecoverExchangeDbsParams) SetRecoverToNewSource(v bool)`

SetRecoverToNewSource sets RecoverToNewSource field to given value.


### GetTargetEnvironment

`func (o *RecoverExchangeDbsParams) GetTargetEnvironment() string`

GetTargetEnvironment returns the TargetEnvironment field if non-nil, zero value otherwise.

### GetTargetEnvironmentOk

`func (o *RecoverExchangeDbsParams) GetTargetEnvironmentOk() (*string, bool)`

GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironment

`func (o *RecoverExchangeDbsParams) SetTargetEnvironment(v string)`

SetTargetEnvironment sets TargetEnvironment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


