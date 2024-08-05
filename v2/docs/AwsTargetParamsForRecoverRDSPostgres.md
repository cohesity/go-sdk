# AwsTargetParamsForRecoverRDSPostgres

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomServerConfig** | Pointer to [**NullableAwsTargetParamsForRecoverRDSPostgresCustomServerConfig**](AwsTargetParamsForRecoverRDSPostgresCustomServerConfig.md) |  | [optional] 
**KnownSourceConfig** | Pointer to [**NullableAwsTargetParamsForRecoverRDSPostgresKnownSourceConfig**](AwsTargetParamsForRecoverRDSPostgresKnownSourceConfig.md) |  | [optional] 
**RecoverToKnownSource** | **NullableBool** | Specifies whether the recovery should be performed to a known or a custom target. | 

## Methods

### NewAwsTargetParamsForRecoverRDSPostgres

`func NewAwsTargetParamsForRecoverRDSPostgres(recoverToKnownSource NullableBool, ) *AwsTargetParamsForRecoverRDSPostgres`

NewAwsTargetParamsForRecoverRDSPostgres instantiates a new AwsTargetParamsForRecoverRDSPostgres object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsTargetParamsForRecoverRDSPostgresWithDefaults

`func NewAwsTargetParamsForRecoverRDSPostgresWithDefaults() *AwsTargetParamsForRecoverRDSPostgres`

NewAwsTargetParamsForRecoverRDSPostgresWithDefaults instantiates a new AwsTargetParamsForRecoverRDSPostgres object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomServerConfig

`func (o *AwsTargetParamsForRecoverRDSPostgres) GetCustomServerConfig() AwsTargetParamsForRecoverRDSPostgresCustomServerConfig`

GetCustomServerConfig returns the CustomServerConfig field if non-nil, zero value otherwise.

### GetCustomServerConfigOk

`func (o *AwsTargetParamsForRecoverRDSPostgres) GetCustomServerConfigOk() (*AwsTargetParamsForRecoverRDSPostgresCustomServerConfig, bool)`

GetCustomServerConfigOk returns a tuple with the CustomServerConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomServerConfig

`func (o *AwsTargetParamsForRecoverRDSPostgres) SetCustomServerConfig(v AwsTargetParamsForRecoverRDSPostgresCustomServerConfig)`

SetCustomServerConfig sets CustomServerConfig field to given value.

### HasCustomServerConfig

`func (o *AwsTargetParamsForRecoverRDSPostgres) HasCustomServerConfig() bool`

HasCustomServerConfig returns a boolean if a field has been set.

### SetCustomServerConfigNil

`func (o *AwsTargetParamsForRecoverRDSPostgres) SetCustomServerConfigNil(b bool)`

 SetCustomServerConfigNil sets the value for CustomServerConfig to be an explicit nil

### UnsetCustomServerConfig
`func (o *AwsTargetParamsForRecoverRDSPostgres) UnsetCustomServerConfig()`

UnsetCustomServerConfig ensures that no value is present for CustomServerConfig, not even an explicit nil
### GetKnownSourceConfig

`func (o *AwsTargetParamsForRecoverRDSPostgres) GetKnownSourceConfig() AwsTargetParamsForRecoverRDSPostgresKnownSourceConfig`

GetKnownSourceConfig returns the KnownSourceConfig field if non-nil, zero value otherwise.

### GetKnownSourceConfigOk

`func (o *AwsTargetParamsForRecoverRDSPostgres) GetKnownSourceConfigOk() (*AwsTargetParamsForRecoverRDSPostgresKnownSourceConfig, bool)`

GetKnownSourceConfigOk returns a tuple with the KnownSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnownSourceConfig

`func (o *AwsTargetParamsForRecoverRDSPostgres) SetKnownSourceConfig(v AwsTargetParamsForRecoverRDSPostgresKnownSourceConfig)`

SetKnownSourceConfig sets KnownSourceConfig field to given value.

### HasKnownSourceConfig

`func (o *AwsTargetParamsForRecoverRDSPostgres) HasKnownSourceConfig() bool`

HasKnownSourceConfig returns a boolean if a field has been set.

### SetKnownSourceConfigNil

`func (o *AwsTargetParamsForRecoverRDSPostgres) SetKnownSourceConfigNil(b bool)`

 SetKnownSourceConfigNil sets the value for KnownSourceConfig to be an explicit nil

### UnsetKnownSourceConfig
`func (o *AwsTargetParamsForRecoverRDSPostgres) UnsetKnownSourceConfig()`

UnsetKnownSourceConfig ensures that no value is present for KnownSourceConfig, not even an explicit nil
### GetRecoverToKnownSource

`func (o *AwsTargetParamsForRecoverRDSPostgres) GetRecoverToKnownSource() bool`

GetRecoverToKnownSource returns the RecoverToKnownSource field if non-nil, zero value otherwise.

### GetRecoverToKnownSourceOk

`func (o *AwsTargetParamsForRecoverRDSPostgres) GetRecoverToKnownSourceOk() (*bool, bool)`

GetRecoverToKnownSourceOk returns a tuple with the RecoverToKnownSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverToKnownSource

`func (o *AwsTargetParamsForRecoverRDSPostgres) SetRecoverToKnownSource(v bool)`

SetRecoverToKnownSource sets RecoverToKnownSource field to given value.


### SetRecoverToKnownSourceNil

`func (o *AwsTargetParamsForRecoverRDSPostgres) SetRecoverToKnownSourceNil(b bool)`

 SetRecoverToKnownSourceNil sets the value for RecoverToKnownSource to be an explicit nil

### UnsetRecoverToKnownSource
`func (o *AwsTargetParamsForRecoverRDSPostgres) UnsetRecoverToKnownSource()`

UnsetRecoverToKnownSource ensures that no value is present for RecoverToKnownSource, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


