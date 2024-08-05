# RecoverAwsVmParamsAwsTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinueOnError** | Pointer to **NullableBool** | Specifies whether to continue recovering other vms if one of vms failed to recover. Default value is false. | [optional] 
**CustomTags** | Pointer to [**[]SimpleTags**](SimpleTags.md) | Specifies the custom tags that need to be present on on every temporary and permanent entity that this job creates. | [optional] 
**FleetConfig** | Pointer to [**NullableAwsTargetParamsForRecoverVmFleetConfig**](AwsTargetParamsForRecoverVmFleetConfig.md) |  | [optional] 
**PowerOnVms** | Pointer to **NullableBool** | Specifies whether to power on vms after recovery. If not specified, or false, recovered vms will be in powered off state. | [optional] 
**RecoveryTargetConfig** | Pointer to [**NullableAwsTargetParamsForRecoverVmRecoveryTargetConfig**](AwsTargetParamsForRecoverVmRecoveryTargetConfig.md) |  | [optional] 
**RenameRecoveredVmsParams** | Pointer to [**NullableAcropolisTargetParamsForRecoverVmRenameRecoveredVmsParams**](AcropolisTargetParamsForRecoverVmRenameRecoveredVmsParams.md) |  | [optional] 

## Methods

### NewRecoverAwsVmParamsAwsTargetParams

`func NewRecoverAwsVmParamsAwsTargetParams() *RecoverAwsVmParamsAwsTargetParams`

NewRecoverAwsVmParamsAwsTargetParams instantiates a new RecoverAwsVmParamsAwsTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverAwsVmParamsAwsTargetParamsWithDefaults

`func NewRecoverAwsVmParamsAwsTargetParamsWithDefaults() *RecoverAwsVmParamsAwsTargetParams`

NewRecoverAwsVmParamsAwsTargetParamsWithDefaults instantiates a new RecoverAwsVmParamsAwsTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinueOnError

`func (o *RecoverAwsVmParamsAwsTargetParams) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *RecoverAwsVmParamsAwsTargetParams) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *RecoverAwsVmParamsAwsTargetParams) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *RecoverAwsVmParamsAwsTargetParams) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *RecoverAwsVmParamsAwsTargetParams) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *RecoverAwsVmParamsAwsTargetParams) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetCustomTags

`func (o *RecoverAwsVmParamsAwsTargetParams) GetCustomTags() []SimpleTags`

GetCustomTags returns the CustomTags field if non-nil, zero value otherwise.

### GetCustomTagsOk

`func (o *RecoverAwsVmParamsAwsTargetParams) GetCustomTagsOk() (*[]SimpleTags, bool)`

GetCustomTagsOk returns a tuple with the CustomTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTags

`func (o *RecoverAwsVmParamsAwsTargetParams) SetCustomTags(v []SimpleTags)`

SetCustomTags sets CustomTags field to given value.

### HasCustomTags

`func (o *RecoverAwsVmParamsAwsTargetParams) HasCustomTags() bool`

HasCustomTags returns a boolean if a field has been set.

### SetCustomTagsNil

`func (o *RecoverAwsVmParamsAwsTargetParams) SetCustomTagsNil(b bool)`

 SetCustomTagsNil sets the value for CustomTags to be an explicit nil

### UnsetCustomTags
`func (o *RecoverAwsVmParamsAwsTargetParams) UnsetCustomTags()`

UnsetCustomTags ensures that no value is present for CustomTags, not even an explicit nil
### GetFleetConfig

`func (o *RecoverAwsVmParamsAwsTargetParams) GetFleetConfig() AwsTargetParamsForRecoverVmFleetConfig`

GetFleetConfig returns the FleetConfig field if non-nil, zero value otherwise.

### GetFleetConfigOk

`func (o *RecoverAwsVmParamsAwsTargetParams) GetFleetConfigOk() (*AwsTargetParamsForRecoverVmFleetConfig, bool)`

GetFleetConfigOk returns a tuple with the FleetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetConfig

`func (o *RecoverAwsVmParamsAwsTargetParams) SetFleetConfig(v AwsTargetParamsForRecoverVmFleetConfig)`

SetFleetConfig sets FleetConfig field to given value.

### HasFleetConfig

`func (o *RecoverAwsVmParamsAwsTargetParams) HasFleetConfig() bool`

HasFleetConfig returns a boolean if a field has been set.

### SetFleetConfigNil

`func (o *RecoverAwsVmParamsAwsTargetParams) SetFleetConfigNil(b bool)`

 SetFleetConfigNil sets the value for FleetConfig to be an explicit nil

### UnsetFleetConfig
`func (o *RecoverAwsVmParamsAwsTargetParams) UnsetFleetConfig()`

UnsetFleetConfig ensures that no value is present for FleetConfig, not even an explicit nil
### GetPowerOnVms

`func (o *RecoverAwsVmParamsAwsTargetParams) GetPowerOnVms() bool`

GetPowerOnVms returns the PowerOnVms field if non-nil, zero value otherwise.

### GetPowerOnVmsOk

`func (o *RecoverAwsVmParamsAwsTargetParams) GetPowerOnVmsOk() (*bool, bool)`

GetPowerOnVmsOk returns a tuple with the PowerOnVms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOnVms

`func (o *RecoverAwsVmParamsAwsTargetParams) SetPowerOnVms(v bool)`

SetPowerOnVms sets PowerOnVms field to given value.

### HasPowerOnVms

`func (o *RecoverAwsVmParamsAwsTargetParams) HasPowerOnVms() bool`

HasPowerOnVms returns a boolean if a field has been set.

### SetPowerOnVmsNil

`func (o *RecoverAwsVmParamsAwsTargetParams) SetPowerOnVmsNil(b bool)`

 SetPowerOnVmsNil sets the value for PowerOnVms to be an explicit nil

### UnsetPowerOnVms
`func (o *RecoverAwsVmParamsAwsTargetParams) UnsetPowerOnVms()`

UnsetPowerOnVms ensures that no value is present for PowerOnVms, not even an explicit nil
### GetRecoveryTargetConfig

`func (o *RecoverAwsVmParamsAwsTargetParams) GetRecoveryTargetConfig() AwsTargetParamsForRecoverVmRecoveryTargetConfig`

GetRecoveryTargetConfig returns the RecoveryTargetConfig field if non-nil, zero value otherwise.

### GetRecoveryTargetConfigOk

`func (o *RecoverAwsVmParamsAwsTargetParams) GetRecoveryTargetConfigOk() (*AwsTargetParamsForRecoverVmRecoveryTargetConfig, bool)`

GetRecoveryTargetConfigOk returns a tuple with the RecoveryTargetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTargetConfig

`func (o *RecoverAwsVmParamsAwsTargetParams) SetRecoveryTargetConfig(v AwsTargetParamsForRecoverVmRecoveryTargetConfig)`

SetRecoveryTargetConfig sets RecoveryTargetConfig field to given value.

### HasRecoveryTargetConfig

`func (o *RecoverAwsVmParamsAwsTargetParams) HasRecoveryTargetConfig() bool`

HasRecoveryTargetConfig returns a boolean if a field has been set.

### SetRecoveryTargetConfigNil

`func (o *RecoverAwsVmParamsAwsTargetParams) SetRecoveryTargetConfigNil(b bool)`

 SetRecoveryTargetConfigNil sets the value for RecoveryTargetConfig to be an explicit nil

### UnsetRecoveryTargetConfig
`func (o *RecoverAwsVmParamsAwsTargetParams) UnsetRecoveryTargetConfig()`

UnsetRecoveryTargetConfig ensures that no value is present for RecoveryTargetConfig, not even an explicit nil
### GetRenameRecoveredVmsParams

`func (o *RecoverAwsVmParamsAwsTargetParams) GetRenameRecoveredVmsParams() AcropolisTargetParamsForRecoverVmRenameRecoveredVmsParams`

GetRenameRecoveredVmsParams returns the RenameRecoveredVmsParams field if non-nil, zero value otherwise.

### GetRenameRecoveredVmsParamsOk

`func (o *RecoverAwsVmParamsAwsTargetParams) GetRenameRecoveredVmsParamsOk() (*AcropolisTargetParamsForRecoverVmRenameRecoveredVmsParams, bool)`

GetRenameRecoveredVmsParamsOk returns a tuple with the RenameRecoveredVmsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenameRecoveredVmsParams

`func (o *RecoverAwsVmParamsAwsTargetParams) SetRenameRecoveredVmsParams(v AcropolisTargetParamsForRecoverVmRenameRecoveredVmsParams)`

SetRenameRecoveredVmsParams sets RenameRecoveredVmsParams field to given value.

### HasRenameRecoveredVmsParams

`func (o *RecoverAwsVmParamsAwsTargetParams) HasRenameRecoveredVmsParams() bool`

HasRenameRecoveredVmsParams returns a boolean if a field has been set.

### SetRenameRecoveredVmsParamsNil

`func (o *RecoverAwsVmParamsAwsTargetParams) SetRenameRecoveredVmsParamsNil(b bool)`

 SetRenameRecoveredVmsParamsNil sets the value for RenameRecoveredVmsParams to be an explicit nil

### UnsetRenameRecoveredVmsParams
`func (o *RecoverAwsVmParamsAwsTargetParams) UnsetRenameRecoveredVmsParams()`

UnsetRenameRecoveredVmsParams ensures that no value is present for RenameRecoveredVmsParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


