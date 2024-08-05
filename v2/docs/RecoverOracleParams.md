# RecoverOracleParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | [**[]RecoverOracleDbSnapshotParams**](RecoverOracleDbSnapshotParams.md) | Specifies the list of parameters for list of objects to be recovered. | 
**RecoverAppParams** | Pointer to [**NullableRecoverOracleParamsRecoverAppParams**](RecoverOracleParamsRecoverAppParams.md) |  | [optional] 
**RecoveryAction** | **string** | Specifies the type of recover action to be performed. | 

## Methods

### NewRecoverOracleParams

`func NewRecoverOracleParams(objects []RecoverOracleDbSnapshotParams, recoveryAction string, ) *RecoverOracleParams`

NewRecoverOracleParams instantiates a new RecoverOracleParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverOracleParamsWithDefaults

`func NewRecoverOracleParamsWithDefaults() *RecoverOracleParams`

NewRecoverOracleParamsWithDefaults instantiates a new RecoverOracleParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjects

`func (o *RecoverOracleParams) GetObjects() []RecoverOracleDbSnapshotParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *RecoverOracleParams) GetObjectsOk() (*[]RecoverOracleDbSnapshotParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *RecoverOracleParams) SetObjects(v []RecoverOracleDbSnapshotParams)`

SetObjects sets Objects field to given value.


### SetObjectsNil

`func (o *RecoverOracleParams) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *RecoverOracleParams) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetRecoverAppParams

`func (o *RecoverOracleParams) GetRecoverAppParams() RecoverOracleParamsRecoverAppParams`

GetRecoverAppParams returns the RecoverAppParams field if non-nil, zero value otherwise.

### GetRecoverAppParamsOk

`func (o *RecoverOracleParams) GetRecoverAppParamsOk() (*RecoverOracleParamsRecoverAppParams, bool)`

GetRecoverAppParamsOk returns a tuple with the RecoverAppParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverAppParams

`func (o *RecoverOracleParams) SetRecoverAppParams(v RecoverOracleParamsRecoverAppParams)`

SetRecoverAppParams sets RecoverAppParams field to given value.

### HasRecoverAppParams

`func (o *RecoverOracleParams) HasRecoverAppParams() bool`

HasRecoverAppParams returns a boolean if a field has been set.

### SetRecoverAppParamsNil

`func (o *RecoverOracleParams) SetRecoverAppParamsNil(b bool)`

 SetRecoverAppParamsNil sets the value for RecoverAppParams to be an explicit nil

### UnsetRecoverAppParams
`func (o *RecoverOracleParams) UnsetRecoverAppParams()`

UnsetRecoverAppParams ensures that no value is present for RecoverAppParams, not even an explicit nil
### GetRecoveryAction

`func (o *RecoverOracleParams) GetRecoveryAction() string`

GetRecoveryAction returns the RecoveryAction field if non-nil, zero value otherwise.

### GetRecoveryActionOk

`func (o *RecoverOracleParams) GetRecoveryActionOk() (*string, bool)`

GetRecoveryActionOk returns a tuple with the RecoveryAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryAction

`func (o *RecoverOracleParams) SetRecoveryAction(v string)`

SetRecoveryAction sets RecoveryAction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


