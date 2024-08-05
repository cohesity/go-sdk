# AzureTargetParamsForRecoverAzureSql

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewSourceConfig** | Pointer to [**NullableAzureTargetParamsForRecoverAzureSqlNewSourceConfig**](AzureTargetParamsForRecoverAzureSqlNewSourceConfig.md) |  | [optional] 
**RecoverToNewSource** | **NullableBool** | Specifies the parameter whether the recovery should be performed to a new or an existing target. | 

## Methods

### NewAzureTargetParamsForRecoverAzureSql

`func NewAzureTargetParamsForRecoverAzureSql(recoverToNewSource NullableBool, ) *AzureTargetParamsForRecoverAzureSql`

NewAzureTargetParamsForRecoverAzureSql instantiates a new AzureTargetParamsForRecoverAzureSql object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureTargetParamsForRecoverAzureSqlWithDefaults

`func NewAzureTargetParamsForRecoverAzureSqlWithDefaults() *AzureTargetParamsForRecoverAzureSql`

NewAzureTargetParamsForRecoverAzureSqlWithDefaults instantiates a new AzureTargetParamsForRecoverAzureSql object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewSourceConfig

`func (o *AzureTargetParamsForRecoverAzureSql) GetNewSourceConfig() AzureTargetParamsForRecoverAzureSqlNewSourceConfig`

GetNewSourceConfig returns the NewSourceConfig field if non-nil, zero value otherwise.

### GetNewSourceConfigOk

`func (o *AzureTargetParamsForRecoverAzureSql) GetNewSourceConfigOk() (*AzureTargetParamsForRecoverAzureSqlNewSourceConfig, bool)`

GetNewSourceConfigOk returns a tuple with the NewSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSourceConfig

`func (o *AzureTargetParamsForRecoverAzureSql) SetNewSourceConfig(v AzureTargetParamsForRecoverAzureSqlNewSourceConfig)`

SetNewSourceConfig sets NewSourceConfig field to given value.

### HasNewSourceConfig

`func (o *AzureTargetParamsForRecoverAzureSql) HasNewSourceConfig() bool`

HasNewSourceConfig returns a boolean if a field has been set.

### SetNewSourceConfigNil

`func (o *AzureTargetParamsForRecoverAzureSql) SetNewSourceConfigNil(b bool)`

 SetNewSourceConfigNil sets the value for NewSourceConfig to be an explicit nil

### UnsetNewSourceConfig
`func (o *AzureTargetParamsForRecoverAzureSql) UnsetNewSourceConfig()`

UnsetNewSourceConfig ensures that no value is present for NewSourceConfig, not even an explicit nil
### GetRecoverToNewSource

`func (o *AzureTargetParamsForRecoverAzureSql) GetRecoverToNewSource() bool`

GetRecoverToNewSource returns the RecoverToNewSource field if non-nil, zero value otherwise.

### GetRecoverToNewSourceOk

`func (o *AzureTargetParamsForRecoverAzureSql) GetRecoverToNewSourceOk() (*bool, bool)`

GetRecoverToNewSourceOk returns a tuple with the RecoverToNewSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverToNewSource

`func (o *AzureTargetParamsForRecoverAzureSql) SetRecoverToNewSource(v bool)`

SetRecoverToNewSource sets RecoverToNewSource field to given value.


### SetRecoverToNewSourceNil

`func (o *AzureTargetParamsForRecoverAzureSql) SetRecoverToNewSourceNil(b bool)`

 SetRecoverToNewSourceNil sets the value for RecoverToNewSource to be an explicit nil

### UnsetRecoverToNewSource
`func (o *AzureTargetParamsForRecoverAzureSql) UnsetRecoverToNewSource()`

UnsetRecoverToNewSource ensures that no value is present for RecoverToNewSource, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


