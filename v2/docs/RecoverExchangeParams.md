# RecoverExchangeParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecoverAppParams** | Pointer to [**NullableRecoverExchangeParamsRecoverAppParams**](RecoverExchangeParamsRecoverAppParams.md) |  | [optional] 
**RecoverExchangeDbsParams** | Pointer to [**NullableRecoverExchangeParamsRecoverExchangeDbsParams**](RecoverExchangeParamsRecoverExchangeDbsParams.md) |  | [optional] 
**RecoveryAction** | **string** | Specifies the type of recover action to be performed. | 

## Methods

### NewRecoverExchangeParams

`func NewRecoverExchangeParams(recoveryAction string, ) *RecoverExchangeParams`

NewRecoverExchangeParams instantiates a new RecoverExchangeParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverExchangeParamsWithDefaults

`func NewRecoverExchangeParamsWithDefaults() *RecoverExchangeParams`

NewRecoverExchangeParamsWithDefaults instantiates a new RecoverExchangeParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecoverAppParams

`func (o *RecoverExchangeParams) GetRecoverAppParams() RecoverExchangeParamsRecoverAppParams`

GetRecoverAppParams returns the RecoverAppParams field if non-nil, zero value otherwise.

### GetRecoverAppParamsOk

`func (o *RecoverExchangeParams) GetRecoverAppParamsOk() (*RecoverExchangeParamsRecoverAppParams, bool)`

GetRecoverAppParamsOk returns a tuple with the RecoverAppParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverAppParams

`func (o *RecoverExchangeParams) SetRecoverAppParams(v RecoverExchangeParamsRecoverAppParams)`

SetRecoverAppParams sets RecoverAppParams field to given value.

### HasRecoverAppParams

`func (o *RecoverExchangeParams) HasRecoverAppParams() bool`

HasRecoverAppParams returns a boolean if a field has been set.

### SetRecoverAppParamsNil

`func (o *RecoverExchangeParams) SetRecoverAppParamsNil(b bool)`

 SetRecoverAppParamsNil sets the value for RecoverAppParams to be an explicit nil

### UnsetRecoverAppParams
`func (o *RecoverExchangeParams) UnsetRecoverAppParams()`

UnsetRecoverAppParams ensures that no value is present for RecoverAppParams, not even an explicit nil
### GetRecoverExchangeDbsParams

`func (o *RecoverExchangeParams) GetRecoverExchangeDbsParams() RecoverExchangeParamsRecoverExchangeDbsParams`

GetRecoverExchangeDbsParams returns the RecoverExchangeDbsParams field if non-nil, zero value otherwise.

### GetRecoverExchangeDbsParamsOk

`func (o *RecoverExchangeParams) GetRecoverExchangeDbsParamsOk() (*RecoverExchangeParamsRecoverExchangeDbsParams, bool)`

GetRecoverExchangeDbsParamsOk returns a tuple with the RecoverExchangeDbsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverExchangeDbsParams

`func (o *RecoverExchangeParams) SetRecoverExchangeDbsParams(v RecoverExchangeParamsRecoverExchangeDbsParams)`

SetRecoverExchangeDbsParams sets RecoverExchangeDbsParams field to given value.

### HasRecoverExchangeDbsParams

`func (o *RecoverExchangeParams) HasRecoverExchangeDbsParams() bool`

HasRecoverExchangeDbsParams returns a boolean if a field has been set.

### SetRecoverExchangeDbsParamsNil

`func (o *RecoverExchangeParams) SetRecoverExchangeDbsParamsNil(b bool)`

 SetRecoverExchangeDbsParamsNil sets the value for RecoverExchangeDbsParams to be an explicit nil

### UnsetRecoverExchangeDbsParams
`func (o *RecoverExchangeParams) UnsetRecoverExchangeDbsParams()`

UnsetRecoverExchangeDbsParams ensures that no value is present for RecoverExchangeDbsParams, not even an explicit nil
### GetRecoveryAction

`func (o *RecoverExchangeParams) GetRecoveryAction() string`

GetRecoveryAction returns the RecoveryAction field if non-nil, zero value otherwise.

### GetRecoveryActionOk

`func (o *RecoverExchangeParams) GetRecoveryActionOk() (*string, bool)`

GetRecoveryActionOk returns a tuple with the RecoveryAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryAction

`func (o *RecoverExchangeParams) SetRecoveryAction(v string)`

SetRecoveryAction sets RecoveryAction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


