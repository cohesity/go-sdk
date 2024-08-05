# RecoverAzureSqlParamsAzureTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewSourceConfig** | Pointer to [**NullableAzureTargetParamsForRecoverAzureSqlNewSourceConfig**](AzureTargetParamsForRecoverAzureSqlNewSourceConfig.md) |  | [optional] 
**RecoverToNewSource** | **NullableBool** | Specifies the parameter whether the recovery should be performed to a new or an existing target. | 

## Methods

### NewRecoverAzureSqlParamsAzureTargetParams

`func NewRecoverAzureSqlParamsAzureTargetParams(recoverToNewSource NullableBool, ) *RecoverAzureSqlParamsAzureTargetParams`

NewRecoverAzureSqlParamsAzureTargetParams instantiates a new RecoverAzureSqlParamsAzureTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverAzureSqlParamsAzureTargetParamsWithDefaults

`func NewRecoverAzureSqlParamsAzureTargetParamsWithDefaults() *RecoverAzureSqlParamsAzureTargetParams`

NewRecoverAzureSqlParamsAzureTargetParamsWithDefaults instantiates a new RecoverAzureSqlParamsAzureTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewSourceConfig

`func (o *RecoverAzureSqlParamsAzureTargetParams) GetNewSourceConfig() AzureTargetParamsForRecoverAzureSqlNewSourceConfig`

GetNewSourceConfig returns the NewSourceConfig field if non-nil, zero value otherwise.

### GetNewSourceConfigOk

`func (o *RecoverAzureSqlParamsAzureTargetParams) GetNewSourceConfigOk() (*AzureTargetParamsForRecoverAzureSqlNewSourceConfig, bool)`

GetNewSourceConfigOk returns a tuple with the NewSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSourceConfig

`func (o *RecoverAzureSqlParamsAzureTargetParams) SetNewSourceConfig(v AzureTargetParamsForRecoverAzureSqlNewSourceConfig)`

SetNewSourceConfig sets NewSourceConfig field to given value.

### HasNewSourceConfig

`func (o *RecoverAzureSqlParamsAzureTargetParams) HasNewSourceConfig() bool`

HasNewSourceConfig returns a boolean if a field has been set.

### SetNewSourceConfigNil

`func (o *RecoverAzureSqlParamsAzureTargetParams) SetNewSourceConfigNil(b bool)`

 SetNewSourceConfigNil sets the value for NewSourceConfig to be an explicit nil

### UnsetNewSourceConfig
`func (o *RecoverAzureSqlParamsAzureTargetParams) UnsetNewSourceConfig()`

UnsetNewSourceConfig ensures that no value is present for NewSourceConfig, not even an explicit nil
### GetRecoverToNewSource

`func (o *RecoverAzureSqlParamsAzureTargetParams) GetRecoverToNewSource() bool`

GetRecoverToNewSource returns the RecoverToNewSource field if non-nil, zero value otherwise.

### GetRecoverToNewSourceOk

`func (o *RecoverAzureSqlParamsAzureTargetParams) GetRecoverToNewSourceOk() (*bool, bool)`

GetRecoverToNewSourceOk returns a tuple with the RecoverToNewSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverToNewSource

`func (o *RecoverAzureSqlParamsAzureTargetParams) SetRecoverToNewSource(v bool)`

SetRecoverToNewSource sets RecoverToNewSource field to given value.


### SetRecoverToNewSourceNil

`func (o *RecoverAzureSqlParamsAzureTargetParams) SetRecoverToNewSourceNil(b bool)`

 SetRecoverToNewSourceNil sets the value for RecoverToNewSource to be an explicit nil

### UnsetRecoverToNewSource
`func (o *RecoverAzureSqlParamsAzureTargetParams) UnsetRecoverToNewSource()`

UnsetRecoverToNewSource ensures that no value is present for RecoverToNewSource, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


