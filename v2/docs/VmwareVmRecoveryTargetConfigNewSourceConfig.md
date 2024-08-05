# VmwareVmRecoveryTargetConfigNewSourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | **NullableString** | Specifies the type of VMware source to which the VMs are being restored. | 
**StandaloneHostParams** | Pointer to [**RecoverVmwareVmEsxiSourceConfig**](RecoverVmwareVmEsxiSourceConfig.md) |  | [optional] 
**VCenterParams** | Pointer to [**RecoverVmwareVmVCenterSourceConfig**](RecoverVmwareVmVCenterSourceConfig.md) |  | [optional] 
**VCloudDirectorParams** | Pointer to [**RecoverVmwareVmVCDSourceConfig**](RecoverVmwareVmVCDSourceConfig.md) |  | [optional] 

## Methods

### NewVmwareVmRecoveryTargetConfigNewSourceConfig

`func NewVmwareVmRecoveryTargetConfigNewSourceConfig(sourceType NullableString, ) *VmwareVmRecoveryTargetConfigNewSourceConfig`

NewVmwareVmRecoveryTargetConfigNewSourceConfig instantiates a new VmwareVmRecoveryTargetConfigNewSourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareVmRecoveryTargetConfigNewSourceConfigWithDefaults

`func NewVmwareVmRecoveryTargetConfigNewSourceConfigWithDefaults() *VmwareVmRecoveryTargetConfigNewSourceConfig`

NewVmwareVmRecoveryTargetConfigNewSourceConfigWithDefaults instantiates a new VmwareVmRecoveryTargetConfigNewSourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *VmwareVmRecoveryTargetConfigNewSourceConfig) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *VmwareVmRecoveryTargetConfigNewSourceConfig) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *VmwareVmRecoveryTargetConfigNewSourceConfig) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### SetSourceTypeNil

`func (o *VmwareVmRecoveryTargetConfigNewSourceConfig) SetSourceTypeNil(b bool)`

 SetSourceTypeNil sets the value for SourceType to be an explicit nil

### UnsetSourceType
`func (o *VmwareVmRecoveryTargetConfigNewSourceConfig) UnsetSourceType()`

UnsetSourceType ensures that no value is present for SourceType, not even an explicit nil
### GetStandaloneHostParams

`func (o *VmwareVmRecoveryTargetConfigNewSourceConfig) GetStandaloneHostParams() RecoverVmwareVmEsxiSourceConfig`

GetStandaloneHostParams returns the StandaloneHostParams field if non-nil, zero value otherwise.

### GetStandaloneHostParamsOk

`func (o *VmwareVmRecoveryTargetConfigNewSourceConfig) GetStandaloneHostParamsOk() (*RecoverVmwareVmEsxiSourceConfig, bool)`

GetStandaloneHostParamsOk returns a tuple with the StandaloneHostParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandaloneHostParams

`func (o *VmwareVmRecoveryTargetConfigNewSourceConfig) SetStandaloneHostParams(v RecoverVmwareVmEsxiSourceConfig)`

SetStandaloneHostParams sets StandaloneHostParams field to given value.

### HasStandaloneHostParams

`func (o *VmwareVmRecoveryTargetConfigNewSourceConfig) HasStandaloneHostParams() bool`

HasStandaloneHostParams returns a boolean if a field has been set.

### GetVCenterParams

`func (o *VmwareVmRecoveryTargetConfigNewSourceConfig) GetVCenterParams() RecoverVmwareVmVCenterSourceConfig`

GetVCenterParams returns the VCenterParams field if non-nil, zero value otherwise.

### GetVCenterParamsOk

`func (o *VmwareVmRecoveryTargetConfigNewSourceConfig) GetVCenterParamsOk() (*RecoverVmwareVmVCenterSourceConfig, bool)`

GetVCenterParamsOk returns a tuple with the VCenterParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCenterParams

`func (o *VmwareVmRecoveryTargetConfigNewSourceConfig) SetVCenterParams(v RecoverVmwareVmVCenterSourceConfig)`

SetVCenterParams sets VCenterParams field to given value.

### HasVCenterParams

`func (o *VmwareVmRecoveryTargetConfigNewSourceConfig) HasVCenterParams() bool`

HasVCenterParams returns a boolean if a field has been set.

### GetVCloudDirectorParams

`func (o *VmwareVmRecoveryTargetConfigNewSourceConfig) GetVCloudDirectorParams() RecoverVmwareVmVCDSourceConfig`

GetVCloudDirectorParams returns the VCloudDirectorParams field if non-nil, zero value otherwise.

### GetVCloudDirectorParamsOk

`func (o *VmwareVmRecoveryTargetConfigNewSourceConfig) GetVCloudDirectorParamsOk() (*RecoverVmwareVmVCDSourceConfig, bool)`

GetVCloudDirectorParamsOk returns a tuple with the VCloudDirectorParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCloudDirectorParams

`func (o *VmwareVmRecoveryTargetConfigNewSourceConfig) SetVCloudDirectorParams(v RecoverVmwareVmVCDSourceConfig)`

SetVCloudDirectorParams sets VCloudDirectorParams field to given value.

### HasVCloudDirectorParams

`func (o *VmwareVmRecoveryTargetConfigNewSourceConfig) HasVCloudDirectorParams() bool`

HasVCloudDirectorParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


