# VmwareVAppRecoveryTargetConfigNewSourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | **NullableString** | Specifies the type of VMware source to which the VMs are being restored. | 
**VCloudDirectorParams** | Pointer to [**RecoverVmwareVAppVCDSourceConfig**](RecoverVmwareVAppVCDSourceConfig.md) |  | [optional] 

## Methods

### NewVmwareVAppRecoveryTargetConfigNewSourceConfig

`func NewVmwareVAppRecoveryTargetConfigNewSourceConfig(sourceType NullableString, ) *VmwareVAppRecoveryTargetConfigNewSourceConfig`

NewVmwareVAppRecoveryTargetConfigNewSourceConfig instantiates a new VmwareVAppRecoveryTargetConfigNewSourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareVAppRecoveryTargetConfigNewSourceConfigWithDefaults

`func NewVmwareVAppRecoveryTargetConfigNewSourceConfigWithDefaults() *VmwareVAppRecoveryTargetConfigNewSourceConfig`

NewVmwareVAppRecoveryTargetConfigNewSourceConfigWithDefaults instantiates a new VmwareVAppRecoveryTargetConfigNewSourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *VmwareVAppRecoveryTargetConfigNewSourceConfig) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *VmwareVAppRecoveryTargetConfigNewSourceConfig) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *VmwareVAppRecoveryTargetConfigNewSourceConfig) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### SetSourceTypeNil

`func (o *VmwareVAppRecoveryTargetConfigNewSourceConfig) SetSourceTypeNil(b bool)`

 SetSourceTypeNil sets the value for SourceType to be an explicit nil

### UnsetSourceType
`func (o *VmwareVAppRecoveryTargetConfigNewSourceConfig) UnsetSourceType()`

UnsetSourceType ensures that no value is present for SourceType, not even an explicit nil
### GetVCloudDirectorParams

`func (o *VmwareVAppRecoveryTargetConfigNewSourceConfig) GetVCloudDirectorParams() RecoverVmwareVAppVCDSourceConfig`

GetVCloudDirectorParams returns the VCloudDirectorParams field if non-nil, zero value otherwise.

### GetVCloudDirectorParamsOk

`func (o *VmwareVAppRecoveryTargetConfigNewSourceConfig) GetVCloudDirectorParamsOk() (*RecoverVmwareVAppVCDSourceConfig, bool)`

GetVCloudDirectorParamsOk returns a tuple with the VCloudDirectorParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCloudDirectorParams

`func (o *VmwareVAppRecoveryTargetConfigNewSourceConfig) SetVCloudDirectorParams(v RecoverVmwareVAppVCDSourceConfig)`

SetVCloudDirectorParams sets VCloudDirectorParams field to given value.

### HasVCloudDirectorParams

`func (o *VmwareVAppRecoveryTargetConfigNewSourceConfig) HasVCloudDirectorParams() bool`

HasVCloudDirectorParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


