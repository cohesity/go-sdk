# RestoreAppObjectParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdRestoreParams** | Pointer to [**RestoreADAppObjectParams**](RestoreADAppObjectParams.md) |  | [optional] 
**CloneTaskId** | Pointer to **NullableInt64** | Id of finished clone task which has to be refreshed with different data. | [optional] 
**ExchangeRestoreParams** | Pointer to [**RestoreExchangeParams**](RestoreExchangeParams.md) |  | [optional] 
**OracleRestoreParams** | Pointer to [**RestoreOracleAppObjectParams**](RestoreOracleAppObjectParams.md) |  | [optional] 
**SqlRestoreParams** | Pointer to [**RestoreSqlAppObjectParams**](RestoreSqlAppObjectParams.md) |  | [optional] 
**TargetHost** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**TargetHostParentSource** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 

## Methods

### NewRestoreAppObjectParams

`func NewRestoreAppObjectParams() *RestoreAppObjectParams`

NewRestoreAppObjectParams instantiates a new RestoreAppObjectParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreAppObjectParamsWithDefaults

`func NewRestoreAppObjectParamsWithDefaults() *RestoreAppObjectParams`

NewRestoreAppObjectParamsWithDefaults instantiates a new RestoreAppObjectParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdRestoreParams

`func (o *RestoreAppObjectParams) GetAdRestoreParams() RestoreADAppObjectParams`

GetAdRestoreParams returns the AdRestoreParams field if non-nil, zero value otherwise.

### GetAdRestoreParamsOk

`func (o *RestoreAppObjectParams) GetAdRestoreParamsOk() (*RestoreADAppObjectParams, bool)`

GetAdRestoreParamsOk returns a tuple with the AdRestoreParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdRestoreParams

`func (o *RestoreAppObjectParams) SetAdRestoreParams(v RestoreADAppObjectParams)`

SetAdRestoreParams sets AdRestoreParams field to given value.

### HasAdRestoreParams

`func (o *RestoreAppObjectParams) HasAdRestoreParams() bool`

HasAdRestoreParams returns a boolean if a field has been set.

### GetCloneTaskId

`func (o *RestoreAppObjectParams) GetCloneTaskId() int64`

GetCloneTaskId returns the CloneTaskId field if non-nil, zero value otherwise.

### GetCloneTaskIdOk

`func (o *RestoreAppObjectParams) GetCloneTaskIdOk() (*int64, bool)`

GetCloneTaskIdOk returns a tuple with the CloneTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneTaskId

`func (o *RestoreAppObjectParams) SetCloneTaskId(v int64)`

SetCloneTaskId sets CloneTaskId field to given value.

### HasCloneTaskId

`func (o *RestoreAppObjectParams) HasCloneTaskId() bool`

HasCloneTaskId returns a boolean if a field has been set.

### SetCloneTaskIdNil

`func (o *RestoreAppObjectParams) SetCloneTaskIdNil(b bool)`

 SetCloneTaskIdNil sets the value for CloneTaskId to be an explicit nil

### UnsetCloneTaskId
`func (o *RestoreAppObjectParams) UnsetCloneTaskId()`

UnsetCloneTaskId ensures that no value is present for CloneTaskId, not even an explicit nil
### GetExchangeRestoreParams

`func (o *RestoreAppObjectParams) GetExchangeRestoreParams() RestoreExchangeParams`

GetExchangeRestoreParams returns the ExchangeRestoreParams field if non-nil, zero value otherwise.

### GetExchangeRestoreParamsOk

`func (o *RestoreAppObjectParams) GetExchangeRestoreParamsOk() (*RestoreExchangeParams, bool)`

GetExchangeRestoreParamsOk returns a tuple with the ExchangeRestoreParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRestoreParams

`func (o *RestoreAppObjectParams) SetExchangeRestoreParams(v RestoreExchangeParams)`

SetExchangeRestoreParams sets ExchangeRestoreParams field to given value.

### HasExchangeRestoreParams

`func (o *RestoreAppObjectParams) HasExchangeRestoreParams() bool`

HasExchangeRestoreParams returns a boolean if a field has been set.

### GetOracleRestoreParams

`func (o *RestoreAppObjectParams) GetOracleRestoreParams() RestoreOracleAppObjectParams`

GetOracleRestoreParams returns the OracleRestoreParams field if non-nil, zero value otherwise.

### GetOracleRestoreParamsOk

`func (o *RestoreAppObjectParams) GetOracleRestoreParamsOk() (*RestoreOracleAppObjectParams, bool)`

GetOracleRestoreParamsOk returns a tuple with the OracleRestoreParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleRestoreParams

`func (o *RestoreAppObjectParams) SetOracleRestoreParams(v RestoreOracleAppObjectParams)`

SetOracleRestoreParams sets OracleRestoreParams field to given value.

### HasOracleRestoreParams

`func (o *RestoreAppObjectParams) HasOracleRestoreParams() bool`

HasOracleRestoreParams returns a boolean if a field has been set.

### GetSqlRestoreParams

`func (o *RestoreAppObjectParams) GetSqlRestoreParams() RestoreSqlAppObjectParams`

GetSqlRestoreParams returns the SqlRestoreParams field if non-nil, zero value otherwise.

### GetSqlRestoreParamsOk

`func (o *RestoreAppObjectParams) GetSqlRestoreParamsOk() (*RestoreSqlAppObjectParams, bool)`

GetSqlRestoreParamsOk returns a tuple with the SqlRestoreParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlRestoreParams

`func (o *RestoreAppObjectParams) SetSqlRestoreParams(v RestoreSqlAppObjectParams)`

SetSqlRestoreParams sets SqlRestoreParams field to given value.

### HasSqlRestoreParams

`func (o *RestoreAppObjectParams) HasSqlRestoreParams() bool`

HasSqlRestoreParams returns a boolean if a field has been set.

### GetTargetHost

`func (o *RestoreAppObjectParams) GetTargetHost() EntityProto`

GetTargetHost returns the TargetHost field if non-nil, zero value otherwise.

### GetTargetHostOk

`func (o *RestoreAppObjectParams) GetTargetHostOk() (*EntityProto, bool)`

GetTargetHostOk returns a tuple with the TargetHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetHost

`func (o *RestoreAppObjectParams) SetTargetHost(v EntityProto)`

SetTargetHost sets TargetHost field to given value.

### HasTargetHost

`func (o *RestoreAppObjectParams) HasTargetHost() bool`

HasTargetHost returns a boolean if a field has been set.

### GetTargetHostParentSource

`func (o *RestoreAppObjectParams) GetTargetHostParentSource() EntityProto`

GetTargetHostParentSource returns the TargetHostParentSource field if non-nil, zero value otherwise.

### GetTargetHostParentSourceOk

`func (o *RestoreAppObjectParams) GetTargetHostParentSourceOk() (*EntityProto, bool)`

GetTargetHostParentSourceOk returns a tuple with the TargetHostParentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetHostParentSource

`func (o *RestoreAppObjectParams) SetTargetHostParentSource(v EntityProto)`

SetTargetHostParentSource sets TargetHostParentSource field to given value.

### HasTargetHostParentSource

`func (o *RestoreAppObjectParams) HasTargetHostParentSource() bool`

HasTargetHostParentSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


