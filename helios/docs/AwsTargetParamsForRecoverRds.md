# AwsTargetParamsForRecoverRds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecoveryTargetConfig** | Pointer to [**NullableAwsRdsRecoveryTargetConfig**](AwsRdsRecoveryTargetConfig.md) | Specifies the recovery target configuration if recovery has to be done to a different location which is different from original source or to original Source with different configuration. If not specified, then the recovery of the vms will be performed to original location with all configuration parameters retained. | [optional] 
**RdsConfig** | Pointer to [**NullableRdsConfig**](RdsConfig.md) | Specifies the RDS params. | [optional] 

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

### GetRecoveryTargetConfig

`func (o *AwsTargetParamsForRecoverRds) GetRecoveryTargetConfig() AwsRdsRecoveryTargetConfig`

GetRecoveryTargetConfig returns the RecoveryTargetConfig field if non-nil, zero value otherwise.

### GetRecoveryTargetConfigOk

`func (o *AwsTargetParamsForRecoverRds) GetRecoveryTargetConfigOk() (*AwsRdsRecoveryTargetConfig, bool)`

GetRecoveryTargetConfigOk returns a tuple with the RecoveryTargetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTargetConfig

`func (o *AwsTargetParamsForRecoverRds) SetRecoveryTargetConfig(v AwsRdsRecoveryTargetConfig)`

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
### GetRdsConfig

`func (o *AwsTargetParamsForRecoverRds) GetRdsConfig() RdsConfig`

GetRdsConfig returns the RdsConfig field if non-nil, zero value otherwise.

### GetRdsConfigOk

`func (o *AwsTargetParamsForRecoverRds) GetRdsConfigOk() (*RdsConfig, bool)`

GetRdsConfigOk returns a tuple with the RdsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdsConfig

`func (o *AwsTargetParamsForRecoverRds) SetRdsConfig(v RdsConfig)`

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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


