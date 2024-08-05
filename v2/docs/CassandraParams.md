# CassandraParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinueOnError** | Pointer to **NullableBool** | Specifies whether to continue recovering other objects if one of Object failed to recover. Default value is false. | [optional] 
**IsMultiStageRestore** | Pointer to **NullableBool** | Specifies whether the current recovery operation is a multi-stage restore operation. | [optional] 
**RecoverCassandraParams** | [**RecoverCassandraParams**](RecoverCassandraParams.md) |  | 
**RecoveryAction** | **string** | Specifies the type of recover action to be performed. | 

## Methods

### NewCassandraParams

`func NewCassandraParams(recoverCassandraParams RecoverCassandraParams, recoveryAction string, ) *CassandraParams`

NewCassandraParams instantiates a new CassandraParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCassandraParamsWithDefaults

`func NewCassandraParamsWithDefaults() *CassandraParams`

NewCassandraParamsWithDefaults instantiates a new CassandraParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinueOnError

`func (o *CassandraParams) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *CassandraParams) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *CassandraParams) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *CassandraParams) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *CassandraParams) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *CassandraParams) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetIsMultiStageRestore

`func (o *CassandraParams) GetIsMultiStageRestore() bool`

GetIsMultiStageRestore returns the IsMultiStageRestore field if non-nil, zero value otherwise.

### GetIsMultiStageRestoreOk

`func (o *CassandraParams) GetIsMultiStageRestoreOk() (*bool, bool)`

GetIsMultiStageRestoreOk returns a tuple with the IsMultiStageRestore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultiStageRestore

`func (o *CassandraParams) SetIsMultiStageRestore(v bool)`

SetIsMultiStageRestore sets IsMultiStageRestore field to given value.

### HasIsMultiStageRestore

`func (o *CassandraParams) HasIsMultiStageRestore() bool`

HasIsMultiStageRestore returns a boolean if a field has been set.

### SetIsMultiStageRestoreNil

`func (o *CassandraParams) SetIsMultiStageRestoreNil(b bool)`

 SetIsMultiStageRestoreNil sets the value for IsMultiStageRestore to be an explicit nil

### UnsetIsMultiStageRestore
`func (o *CassandraParams) UnsetIsMultiStageRestore()`

UnsetIsMultiStageRestore ensures that no value is present for IsMultiStageRestore, not even an explicit nil
### GetRecoverCassandraParams

`func (o *CassandraParams) GetRecoverCassandraParams() RecoverCassandraParams`

GetRecoverCassandraParams returns the RecoverCassandraParams field if non-nil, zero value otherwise.

### GetRecoverCassandraParamsOk

`func (o *CassandraParams) GetRecoverCassandraParamsOk() (*RecoverCassandraParams, bool)`

GetRecoverCassandraParamsOk returns a tuple with the RecoverCassandraParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverCassandraParams

`func (o *CassandraParams) SetRecoverCassandraParams(v RecoverCassandraParams)`

SetRecoverCassandraParams sets RecoverCassandraParams field to given value.


### GetRecoveryAction

`func (o *CassandraParams) GetRecoveryAction() string`

GetRecoveryAction returns the RecoveryAction field if non-nil, zero value otherwise.

### GetRecoveryActionOk

`func (o *CassandraParams) GetRecoveryActionOk() (*string, bool)`

GetRecoveryActionOk returns a tuple with the RecoveryAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryAction

`func (o *CassandraParams) SetRecoveryAction(v string)`

SetRecoveryAction sets RecoveryAction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


