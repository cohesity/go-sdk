# MongodbParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecoveryAction** | **string** | Specifies the type of recover action to be performed. | 
**RecoverMongodbParams** | [**RecoverMongodbParams**](RecoverMongodbParams.md) |  | 

## Methods

### NewMongodbParams

`func NewMongodbParams(recoveryAction string, recoverMongodbParams RecoverMongodbParams, ) *MongodbParams`

NewMongodbParams instantiates a new MongodbParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMongodbParamsWithDefaults

`func NewMongodbParamsWithDefaults() *MongodbParams`

NewMongodbParamsWithDefaults instantiates a new MongodbParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecoveryAction

`func (o *MongodbParams) GetRecoveryAction() string`

GetRecoveryAction returns the RecoveryAction field if non-nil, zero value otherwise.

### GetRecoveryActionOk

`func (o *MongodbParams) GetRecoveryActionOk() (*string, bool)`

GetRecoveryActionOk returns a tuple with the RecoveryAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryAction

`func (o *MongodbParams) SetRecoveryAction(v string)`

SetRecoveryAction sets RecoveryAction field to given value.


### GetRecoverMongodbParams

`func (o *MongodbParams) GetRecoverMongodbParams() RecoverMongodbParams`

GetRecoverMongodbParams returns the RecoverMongodbParams field if non-nil, zero value otherwise.

### GetRecoverMongodbParamsOk

`func (o *MongodbParams) GetRecoverMongodbParamsOk() (*RecoverMongodbParams, bool)`

GetRecoverMongodbParamsOk returns a tuple with the RecoverMongodbParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverMongodbParams

`func (o *MongodbParams) SetRecoverMongodbParams(v RecoverMongodbParams)`

SetRecoverMongodbParams sets RecoverMongodbParams field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


