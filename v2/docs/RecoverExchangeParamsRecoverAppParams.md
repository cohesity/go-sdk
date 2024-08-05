# RecoverExchangeParamsRecoverAppParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExchangeTargetParams** | Pointer to [**NullableRecoverExchangeAppParamsExchangeTargetParams**](RecoverExchangeAppParamsExchangeTargetParams.md) |  | [optional] 
**TargetEnvironment** | **string** | Specifies the environment of the recovery target. The corresponding params below must be filled out. | 

## Methods

### NewRecoverExchangeParamsRecoverAppParams

`func NewRecoverExchangeParamsRecoverAppParams(targetEnvironment string, ) *RecoverExchangeParamsRecoverAppParams`

NewRecoverExchangeParamsRecoverAppParams instantiates a new RecoverExchangeParamsRecoverAppParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverExchangeParamsRecoverAppParamsWithDefaults

`func NewRecoverExchangeParamsRecoverAppParamsWithDefaults() *RecoverExchangeParamsRecoverAppParams`

NewRecoverExchangeParamsRecoverAppParamsWithDefaults instantiates a new RecoverExchangeParamsRecoverAppParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExchangeTargetParams

`func (o *RecoverExchangeParamsRecoverAppParams) GetExchangeTargetParams() RecoverExchangeAppParamsExchangeTargetParams`

GetExchangeTargetParams returns the ExchangeTargetParams field if non-nil, zero value otherwise.

### GetExchangeTargetParamsOk

`func (o *RecoverExchangeParamsRecoverAppParams) GetExchangeTargetParamsOk() (*RecoverExchangeAppParamsExchangeTargetParams, bool)`

GetExchangeTargetParamsOk returns a tuple with the ExchangeTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeTargetParams

`func (o *RecoverExchangeParamsRecoverAppParams) SetExchangeTargetParams(v RecoverExchangeAppParamsExchangeTargetParams)`

SetExchangeTargetParams sets ExchangeTargetParams field to given value.

### HasExchangeTargetParams

`func (o *RecoverExchangeParamsRecoverAppParams) HasExchangeTargetParams() bool`

HasExchangeTargetParams returns a boolean if a field has been set.

### SetExchangeTargetParamsNil

`func (o *RecoverExchangeParamsRecoverAppParams) SetExchangeTargetParamsNil(b bool)`

 SetExchangeTargetParamsNil sets the value for ExchangeTargetParams to be an explicit nil

### UnsetExchangeTargetParams
`func (o *RecoverExchangeParamsRecoverAppParams) UnsetExchangeTargetParams()`

UnsetExchangeTargetParams ensures that no value is present for ExchangeTargetParams, not even an explicit nil
### GetTargetEnvironment

`func (o *RecoverExchangeParamsRecoverAppParams) GetTargetEnvironment() string`

GetTargetEnvironment returns the TargetEnvironment field if non-nil, zero value otherwise.

### GetTargetEnvironmentOk

`func (o *RecoverExchangeParamsRecoverAppParams) GetTargetEnvironmentOk() (*string, bool)`

GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironment

`func (o *RecoverExchangeParamsRecoverAppParams) SetTargetEnvironment(v string)`

SetTargetEnvironment sets TargetEnvironment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


