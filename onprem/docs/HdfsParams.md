# HdfsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecoveryAction** | **string** | Specifies the type of recover action to be performed. | 
**RecoverHdfsParams** | [**RecoverHdfsParams**](RecoverHdfsParams.md) |  | 

## Methods

### NewHdfsParams

`func NewHdfsParams(recoveryAction string, recoverHdfsParams RecoverHdfsParams, ) *HdfsParams`

NewHdfsParams instantiates a new HdfsParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHdfsParamsWithDefaults

`func NewHdfsParamsWithDefaults() *HdfsParams`

NewHdfsParamsWithDefaults instantiates a new HdfsParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecoveryAction

`func (o *HdfsParams) GetRecoveryAction() string`

GetRecoveryAction returns the RecoveryAction field if non-nil, zero value otherwise.

### GetRecoveryActionOk

`func (o *HdfsParams) GetRecoveryActionOk() (*string, bool)`

GetRecoveryActionOk returns a tuple with the RecoveryAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryAction

`func (o *HdfsParams) SetRecoveryAction(v string)`

SetRecoveryAction sets RecoveryAction field to given value.


### GetRecoverHdfsParams

`func (o *HdfsParams) GetRecoverHdfsParams() RecoverHdfsParams`

GetRecoverHdfsParams returns the RecoverHdfsParams field if non-nil, zero value otherwise.

### GetRecoverHdfsParamsOk

`func (o *HdfsParams) GetRecoverHdfsParamsOk() (*RecoverHdfsParams, bool)`

GetRecoverHdfsParamsOk returns a tuple with the RecoverHdfsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverHdfsParams

`func (o *HdfsParams) SetRecoverHdfsParams(v RecoverHdfsParams)`

SetRecoverHdfsParams sets RecoverHdfsParams field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


