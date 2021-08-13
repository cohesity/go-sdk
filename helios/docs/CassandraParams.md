# CassandraParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecoveryAction** | **string** | Specifies the type of recover action to be performed. | 
**RecoverCassandraParams** | [**RecoverCassandraParams**](RecoverCassandraParams.md) |  | 

## Methods

### NewCassandraParams

`func NewCassandraParams(recoveryAction string, recoverCassandraParams RecoverCassandraParams, ) *CassandraParams`

NewCassandraParams instantiates a new CassandraParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCassandraParamsWithDefaults

`func NewCassandraParamsWithDefaults() *CassandraParams`

NewCassandraParamsWithDefaults instantiates a new CassandraParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


