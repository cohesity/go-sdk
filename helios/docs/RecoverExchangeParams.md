# RecoverExchangeParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecoveryAction** | **string** | Specifies the type of recover action to be performed. | 
**RecoverAppParams** | Pointer to [**NullableRecoverExchangeAppParams**](RecoverExchangeAppParams.md) | Specifies the parameters to recover Exchange databases. | [optional] 

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


### GetRecoverAppParams

`func (o *RecoverExchangeParams) GetRecoverAppParams() RecoverExchangeAppParams`

GetRecoverAppParams returns the RecoverAppParams field if non-nil, zero value otherwise.

### GetRecoverAppParamsOk

`func (o *RecoverExchangeParams) GetRecoverAppParamsOk() (*RecoverExchangeAppParams, bool)`

GetRecoverAppParamsOk returns a tuple with the RecoverAppParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverAppParams

`func (o *RecoverExchangeParams) SetRecoverAppParams(v RecoverExchangeAppParams)`

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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


