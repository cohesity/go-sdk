# VmwareVAppRecoveryTargetConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecoverToNewSource** | **bool** | Specifies the parameter whether the recovery should be performed to a new or an existing Source Target. | 
**NewSourceConfig** | Pointer to [**NullableRecoverVmwareVAppNewSourceConfig**](RecoverVmwareVAppNewSourceConfig.md) | Specifies the new destination Source configuration parameters where the VMs will be recovered. This is mandatory if recoverToNewSource is set to true. | [optional] 
**OriginalSourceConfig** | Pointer to [**NullableRecoverVmwareVAppOriginalSourceConfig**](RecoverVmwareVAppOriginalSourceConfig.md) | Specifies the Source configuration if VM&#39;s are being recovered to Original Source. If not specified, all the configuration parameters will be retained. | [optional] 

## Methods

### NewVmwareVAppRecoveryTargetConfig

`func NewVmwareVAppRecoveryTargetConfig(recoverToNewSource bool, ) *VmwareVAppRecoveryTargetConfig`

NewVmwareVAppRecoveryTargetConfig instantiates a new VmwareVAppRecoveryTargetConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareVAppRecoveryTargetConfigWithDefaults

`func NewVmwareVAppRecoveryTargetConfigWithDefaults() *VmwareVAppRecoveryTargetConfig`

NewVmwareVAppRecoveryTargetConfigWithDefaults instantiates a new VmwareVAppRecoveryTargetConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecoverToNewSource

`func (o *VmwareVAppRecoveryTargetConfig) GetRecoverToNewSource() bool`

GetRecoverToNewSource returns the RecoverToNewSource field if non-nil, zero value otherwise.

### GetRecoverToNewSourceOk

`func (o *VmwareVAppRecoveryTargetConfig) GetRecoverToNewSourceOk() (*bool, bool)`

GetRecoverToNewSourceOk returns a tuple with the RecoverToNewSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverToNewSource

`func (o *VmwareVAppRecoveryTargetConfig) SetRecoverToNewSource(v bool)`

SetRecoverToNewSource sets RecoverToNewSource field to given value.


### GetNewSourceConfig

`func (o *VmwareVAppRecoveryTargetConfig) GetNewSourceConfig() RecoverVmwareVAppNewSourceConfig`

GetNewSourceConfig returns the NewSourceConfig field if non-nil, zero value otherwise.

### GetNewSourceConfigOk

`func (o *VmwareVAppRecoveryTargetConfig) GetNewSourceConfigOk() (*RecoverVmwareVAppNewSourceConfig, bool)`

GetNewSourceConfigOk returns a tuple with the NewSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSourceConfig

`func (o *VmwareVAppRecoveryTargetConfig) SetNewSourceConfig(v RecoverVmwareVAppNewSourceConfig)`

SetNewSourceConfig sets NewSourceConfig field to given value.

### HasNewSourceConfig

`func (o *VmwareVAppRecoveryTargetConfig) HasNewSourceConfig() bool`

HasNewSourceConfig returns a boolean if a field has been set.

### SetNewSourceConfigNil

`func (o *VmwareVAppRecoveryTargetConfig) SetNewSourceConfigNil(b bool)`

 SetNewSourceConfigNil sets the value for NewSourceConfig to be an explicit nil

### UnsetNewSourceConfig
`func (o *VmwareVAppRecoveryTargetConfig) UnsetNewSourceConfig()`

UnsetNewSourceConfig ensures that no value is present for NewSourceConfig, not even an explicit nil
### GetOriginalSourceConfig

`func (o *VmwareVAppRecoveryTargetConfig) GetOriginalSourceConfig() RecoverVmwareVAppOriginalSourceConfig`

GetOriginalSourceConfig returns the OriginalSourceConfig field if non-nil, zero value otherwise.

### GetOriginalSourceConfigOk

`func (o *VmwareVAppRecoveryTargetConfig) GetOriginalSourceConfigOk() (*RecoverVmwareVAppOriginalSourceConfig, bool)`

GetOriginalSourceConfigOk returns a tuple with the OriginalSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalSourceConfig

`func (o *VmwareVAppRecoveryTargetConfig) SetOriginalSourceConfig(v RecoverVmwareVAppOriginalSourceConfig)`

SetOriginalSourceConfig sets OriginalSourceConfig field to given value.

### HasOriginalSourceConfig

`func (o *VmwareVAppRecoveryTargetConfig) HasOriginalSourceConfig() bool`

HasOriginalSourceConfig returns a boolean if a field has been set.

### SetOriginalSourceConfigNil

`func (o *VmwareVAppRecoveryTargetConfig) SetOriginalSourceConfigNil(b bool)`

 SetOriginalSourceConfigNil sets the value for OriginalSourceConfig to be an explicit nil

### UnsetOriginalSourceConfig
`func (o *VmwareVAppRecoveryTargetConfig) UnsetOriginalSourceConfig()`

UnsetOriginalSourceConfig ensures that no value is present for OriginalSourceConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


