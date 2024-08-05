# VmwareVAppTemplateRecoveryTargetConfigNewSourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | **NullableString** | Specifies the type of VMware source to which the vApp templatess are being restored. | 
**VCloudDirectorParams** | Pointer to [**RecoverVmwareVAppTemplateVCDSourceConfig**](RecoverVmwareVAppTemplateVCDSourceConfig.md) |  | [optional] 

## Methods

### NewVmwareVAppTemplateRecoveryTargetConfigNewSourceConfig

`func NewVmwareVAppTemplateRecoveryTargetConfigNewSourceConfig(sourceType NullableString, ) *VmwareVAppTemplateRecoveryTargetConfigNewSourceConfig`

NewVmwareVAppTemplateRecoveryTargetConfigNewSourceConfig instantiates a new VmwareVAppTemplateRecoveryTargetConfigNewSourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareVAppTemplateRecoveryTargetConfigNewSourceConfigWithDefaults

`func NewVmwareVAppTemplateRecoveryTargetConfigNewSourceConfigWithDefaults() *VmwareVAppTemplateRecoveryTargetConfigNewSourceConfig`

NewVmwareVAppTemplateRecoveryTargetConfigNewSourceConfigWithDefaults instantiates a new VmwareVAppTemplateRecoveryTargetConfigNewSourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *VmwareVAppTemplateRecoveryTargetConfigNewSourceConfig) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *VmwareVAppTemplateRecoveryTargetConfigNewSourceConfig) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *VmwareVAppTemplateRecoveryTargetConfigNewSourceConfig) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### SetSourceTypeNil

`func (o *VmwareVAppTemplateRecoveryTargetConfigNewSourceConfig) SetSourceTypeNil(b bool)`

 SetSourceTypeNil sets the value for SourceType to be an explicit nil

### UnsetSourceType
`func (o *VmwareVAppTemplateRecoveryTargetConfigNewSourceConfig) UnsetSourceType()`

UnsetSourceType ensures that no value is present for SourceType, not even an explicit nil
### GetVCloudDirectorParams

`func (o *VmwareVAppTemplateRecoveryTargetConfigNewSourceConfig) GetVCloudDirectorParams() RecoverVmwareVAppTemplateVCDSourceConfig`

GetVCloudDirectorParams returns the VCloudDirectorParams field if non-nil, zero value otherwise.

### GetVCloudDirectorParamsOk

`func (o *VmwareVAppTemplateRecoveryTargetConfigNewSourceConfig) GetVCloudDirectorParamsOk() (*RecoverVmwareVAppTemplateVCDSourceConfig, bool)`

GetVCloudDirectorParamsOk returns a tuple with the VCloudDirectorParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCloudDirectorParams

`func (o *VmwareVAppTemplateRecoveryTargetConfigNewSourceConfig) SetVCloudDirectorParams(v RecoverVmwareVAppTemplateVCDSourceConfig)`

SetVCloudDirectorParams sets VCloudDirectorParams field to given value.

### HasVCloudDirectorParams

`func (o *VmwareVAppTemplateRecoveryTargetConfigNewSourceConfig) HasVCloudDirectorParams() bool`

HasVCloudDirectorParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


