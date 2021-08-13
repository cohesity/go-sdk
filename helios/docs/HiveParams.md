# HiveParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecoveryAction** | **string** | Specifies the type of recover action to be performed. | 
**RecoverHiveParams** | [**RecoverHiveParams**](RecoverHiveParams.md) |  | 

## Methods

### NewHiveParams

`func NewHiveParams(recoveryAction string, recoverHiveParams RecoverHiveParams, ) *HiveParams`

NewHiveParams instantiates a new HiveParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHiveParamsWithDefaults

`func NewHiveParamsWithDefaults() *HiveParams`

NewHiveParamsWithDefaults instantiates a new HiveParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecoveryAction

`func (o *HiveParams) GetRecoveryAction() string`

GetRecoveryAction returns the RecoveryAction field if non-nil, zero value otherwise.

### GetRecoveryActionOk

`func (o *HiveParams) GetRecoveryActionOk() (*string, bool)`

GetRecoveryActionOk returns a tuple with the RecoveryAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryAction

`func (o *HiveParams) SetRecoveryAction(v string)`

SetRecoveryAction sets RecoveryAction field to given value.


### GetRecoverHiveParams

`func (o *HiveParams) GetRecoverHiveParams() RecoverHiveParams`

GetRecoverHiveParams returns the RecoverHiveParams field if non-nil, zero value otherwise.

### GetRecoverHiveParamsOk

`func (o *HiveParams) GetRecoverHiveParamsOk() (*RecoverHiveParams, bool)`

GetRecoverHiveParamsOk returns a tuple with the RecoverHiveParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverHiveParams

`func (o *HiveParams) SetRecoverHiveParams(v RecoverHiveParams)`

SetRecoverHiveParams sets RecoverHiveParams field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


