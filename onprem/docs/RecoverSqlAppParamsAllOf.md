# RecoverSqlAppParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetEnvironment** | **string** | Specifies the environment of the recovery target. The corresponding params below must be filled out. | 
**SqlTargetParams** | Pointer to [**SqlTargetParamsForRecoverSqlApp**](SqlTargetParamsForRecoverSqlApp.md) |  | [optional] 

## Methods

### NewRecoverSqlAppParamsAllOf

`func NewRecoverSqlAppParamsAllOf(targetEnvironment string, ) *RecoverSqlAppParamsAllOf`

NewRecoverSqlAppParamsAllOf instantiates a new RecoverSqlAppParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverSqlAppParamsAllOfWithDefaults

`func NewRecoverSqlAppParamsAllOfWithDefaults() *RecoverSqlAppParamsAllOf`

NewRecoverSqlAppParamsAllOfWithDefaults instantiates a new RecoverSqlAppParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetEnvironment

`func (o *RecoverSqlAppParamsAllOf) GetTargetEnvironment() string`

GetTargetEnvironment returns the TargetEnvironment field if non-nil, zero value otherwise.

### GetTargetEnvironmentOk

`func (o *RecoverSqlAppParamsAllOf) GetTargetEnvironmentOk() (*string, bool)`

GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironment

`func (o *RecoverSqlAppParamsAllOf) SetTargetEnvironment(v string)`

SetTargetEnvironment sets TargetEnvironment field to given value.


### GetSqlTargetParams

`func (o *RecoverSqlAppParamsAllOf) GetSqlTargetParams() SqlTargetParamsForRecoverSqlApp`

GetSqlTargetParams returns the SqlTargetParams field if non-nil, zero value otherwise.

### GetSqlTargetParamsOk

`func (o *RecoverSqlAppParamsAllOf) GetSqlTargetParamsOk() (*SqlTargetParamsForRecoverSqlApp, bool)`

GetSqlTargetParamsOk returns a tuple with the SqlTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlTargetParams

`func (o *RecoverSqlAppParamsAllOf) SetSqlTargetParams(v SqlTargetParamsForRecoverSqlApp)`

SetSqlTargetParams sets SqlTargetParams field to given value.

### HasSqlTargetParams

`func (o *RecoverSqlAppParamsAllOf) HasSqlTargetParams() bool`

HasSqlTargetParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


