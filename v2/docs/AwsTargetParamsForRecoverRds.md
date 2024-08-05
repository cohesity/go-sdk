# AwsTargetParamsForRecoverRds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RdsConfig** | Pointer to [**NullableAwsTargetParamsForRecoverRdsRdsConfig**](AwsTargetParamsForRecoverRdsRdsConfig.md) |  | [optional] 
**RecoveryTargetConfig** | Pointer to [**NullableAwsTargetParamsForRecoverRdsRecoveryTargetConfig**](AwsTargetParamsForRecoverRdsRecoveryTargetConfig.md) |  | [optional] 

## Methods

### NewAwsTargetParamsForRecoverRds

`func NewAwsTargetParamsForRecoverRds() *AwsTargetParamsForRecoverRds`

NewAwsTargetParamsForRecoverRds instantiates a new AwsTargetParamsForRecoverRds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsTargetParamsForRecoverRdsWithDefaults

`func NewAwsTargetParamsForRecoverRdsWithDefaults() *AwsTargetParamsForRecoverRds`

NewAwsTargetParamsForRecoverRdsWithDefaults instantiates a new AwsTargetParamsForRecoverRds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRdsConfig

`func (o *AwsTargetParamsForRecoverRds) GetRdsConfig() AwsTargetParamsForRecoverRdsRdsConfig`

GetRdsConfig returns the RdsConfig field if non-nil, zero value otherwise.

### GetRdsConfigOk

`func (o *AwsTargetParamsForRecoverRds) GetRdsConfigOk() (*AwsTargetParamsForRecoverRdsRdsConfig, bool)`

GetRdsConfigOk returns a tuple with the RdsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdsConfig

`func (o *AwsTargetParamsForRecoverRds) SetRdsConfig(v AwsTargetParamsForRecoverRdsRdsConfig)`

SetRdsConfig sets RdsConfig field to given value.

### HasRdsConfig

`func (o *AwsTargetParamsForRecoverRds) HasRdsConfig() bool`

HasRdsConfig returns a boolean if a field has been set.

### SetRdsConfigNil

`func (o *AwsTargetParamsForRecoverRds) SetRdsConfigNil(b bool)`

 SetRdsConfigNil sets the value for RdsConfig to be an explicit nil

### UnsetRdsConfig
`func (o *AwsTargetParamsForRecoverRds) UnsetRdsConfig()`

UnsetRdsConfig ensures that no value is present for RdsConfig, not even an explicit nil
### GetRecoveryTargetConfig

`func (o *AwsTargetParamsForRecoverRds) GetRecoveryTargetConfig() AwsTargetParamsForRecoverRdsRecoveryTargetConfig`

GetRecoveryTargetConfig returns the RecoveryTargetConfig field if non-nil, zero value otherwise.

### GetRecoveryTargetConfigOk

`func (o *AwsTargetParamsForRecoverRds) GetRecoveryTargetConfigOk() (*AwsTargetParamsForRecoverRdsRecoveryTargetConfig, bool)`

GetRecoveryTargetConfigOk returns a tuple with the RecoveryTargetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTargetConfig

`func (o *AwsTargetParamsForRecoverRds) SetRecoveryTargetConfig(v AwsTargetParamsForRecoverRdsRecoveryTargetConfig)`

SetRecoveryTargetConfig sets RecoveryTargetConfig field to given value.

### HasRecoveryTargetConfig

`func (o *AwsTargetParamsForRecoverRds) HasRecoveryTargetConfig() bool`

HasRecoveryTargetConfig returns a boolean if a field has been set.

### SetRecoveryTargetConfigNil

`func (o *AwsTargetParamsForRecoverRds) SetRecoveryTargetConfigNil(b bool)`

 SetRecoveryTargetConfigNil sets the value for RecoveryTargetConfig to be an explicit nil

### UnsetRecoveryTargetConfig
`func (o *AwsTargetParamsForRecoverRds) UnsetRecoveryTargetConfig()`

UnsetRecoveryTargetConfig ensures that no value is present for RecoveryTargetConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


