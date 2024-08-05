# HbaseParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecoverHbaseParams** | [**RecoverHbaseParams**](RecoverHbaseParams.md) |  | 
**RecoveryAction** | **string** | Specifies the type of recover action to be performed. | 

## Methods

### NewHbaseParams

`func NewHbaseParams(recoverHbaseParams RecoverHbaseParams, recoveryAction string, ) *HbaseParams`

NewHbaseParams instantiates a new HbaseParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHbaseParamsWithDefaults

`func NewHbaseParamsWithDefaults() *HbaseParams`

NewHbaseParamsWithDefaults instantiates a new HbaseParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecoverHbaseParams

`func (o *HbaseParams) GetRecoverHbaseParams() RecoverHbaseParams`

GetRecoverHbaseParams returns the RecoverHbaseParams field if non-nil, zero value otherwise.

### GetRecoverHbaseParamsOk

`func (o *HbaseParams) GetRecoverHbaseParamsOk() (*RecoverHbaseParams, bool)`

GetRecoverHbaseParamsOk returns a tuple with the RecoverHbaseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverHbaseParams

`func (o *HbaseParams) SetRecoverHbaseParams(v RecoverHbaseParams)`

SetRecoverHbaseParams sets RecoverHbaseParams field to given value.


### GetRecoveryAction

`func (o *HbaseParams) GetRecoveryAction() string`

GetRecoveryAction returns the RecoveryAction field if non-nil, zero value otherwise.

### GetRecoveryActionOk

`func (o *HbaseParams) GetRecoveryActionOk() (*string, bool)`

GetRecoveryActionOk returns a tuple with the RecoveryAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryAction

`func (o *HbaseParams) SetRecoveryAction(v string)`

SetRecoveryAction sets RecoveryAction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


