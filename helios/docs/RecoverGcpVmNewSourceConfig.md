# RecoverGcpVmNewSourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | [**NullableRecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the id of the parent source to recover the VMs. | 
**Project** | [**NullableRecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the GCP project in which to deploy the VM. | 
**Region** | [**NullableRecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the GCP region in which to deploy the VM. | 
**AvailabilityZone** | [**NullableRecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the GCP zone in which to deploy the VM. | 
**NetworkConfig** | [**NullableRecoverGcpVmNewSourceNetworkConfig**](RecoverGcpVmNewSourceNetworkConfig.md) | Specifies the networking configuration to be applied to the recovered VMs. | 

## Methods

### NewRecoverGcpVmNewSourceConfig

`func NewRecoverGcpVmNewSourceConfig(source NullableRecoveryObjectIdentifier, project NullableRecoveryObjectIdentifier, region NullableRecoveryObjectIdentifier, availabilityZone NullableRecoveryObjectIdentifier, networkConfig NullableRecoverGcpVmNewSourceNetworkConfig, ) *RecoverGcpVmNewSourceConfig`

NewRecoverGcpVmNewSourceConfig instantiates a new RecoverGcpVmNewSourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverGcpVmNewSourceConfigWithDefaults

`func NewRecoverGcpVmNewSourceConfigWithDefaults() *RecoverGcpVmNewSourceConfig`

NewRecoverGcpVmNewSourceConfigWithDefaults instantiates a new RecoverGcpVmNewSourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *RecoverGcpVmNewSourceConfig) GetSource() RecoveryObjectIdentifier`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RecoverGcpVmNewSourceConfig) GetSourceOk() (*RecoveryObjectIdentifier, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RecoverGcpVmNewSourceConfig) SetSource(v RecoveryObjectIdentifier)`

SetSource sets Source field to given value.


### SetSourceNil

`func (o *RecoverGcpVmNewSourceConfig) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *RecoverGcpVmNewSourceConfig) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetProject

`func (o *RecoverGcpVmNewSourceConfig) GetProject() RecoveryObjectIdentifier`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *RecoverGcpVmNewSourceConfig) GetProjectOk() (*RecoveryObjectIdentifier, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *RecoverGcpVmNewSourceConfig) SetProject(v RecoveryObjectIdentifier)`

SetProject sets Project field to given value.


### SetProjectNil

`func (o *RecoverGcpVmNewSourceConfig) SetProjectNil(b bool)`

 SetProjectNil sets the value for Project to be an explicit nil

### UnsetProject
`func (o *RecoverGcpVmNewSourceConfig) UnsetProject()`

UnsetProject ensures that no value is present for Project, not even an explicit nil
### GetRegion

`func (o *RecoverGcpVmNewSourceConfig) GetRegion() RecoveryObjectIdentifier`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *RecoverGcpVmNewSourceConfig) GetRegionOk() (*RecoveryObjectIdentifier, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *RecoverGcpVmNewSourceConfig) SetRegion(v RecoveryObjectIdentifier)`

SetRegion sets Region field to given value.


### SetRegionNil

`func (o *RecoverGcpVmNewSourceConfig) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *RecoverGcpVmNewSourceConfig) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetAvailabilityZone

`func (o *RecoverGcpVmNewSourceConfig) GetAvailabilityZone() RecoveryObjectIdentifier`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *RecoverGcpVmNewSourceConfig) GetAvailabilityZoneOk() (*RecoveryObjectIdentifier, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *RecoverGcpVmNewSourceConfig) SetAvailabilityZone(v RecoveryObjectIdentifier)`

SetAvailabilityZone sets AvailabilityZone field to given value.


### SetAvailabilityZoneNil

`func (o *RecoverGcpVmNewSourceConfig) SetAvailabilityZoneNil(b bool)`

 SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil

### UnsetAvailabilityZone
`func (o *RecoverGcpVmNewSourceConfig) UnsetAvailabilityZone()`

UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil
### GetNetworkConfig

`func (o *RecoverGcpVmNewSourceConfig) GetNetworkConfig() RecoverGcpVmNewSourceNetworkConfig`

GetNetworkConfig returns the NetworkConfig field if non-nil, zero value otherwise.

### GetNetworkConfigOk

`func (o *RecoverGcpVmNewSourceConfig) GetNetworkConfigOk() (*RecoverGcpVmNewSourceNetworkConfig, bool)`

GetNetworkConfigOk returns a tuple with the NetworkConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConfig

`func (o *RecoverGcpVmNewSourceConfig) SetNetworkConfig(v RecoverGcpVmNewSourceNetworkConfig)`

SetNetworkConfig sets NetworkConfig field to given value.


### SetNetworkConfigNil

`func (o *RecoverGcpVmNewSourceConfig) SetNetworkConfigNil(b bool)`

 SetNetworkConfigNil sets the value for NetworkConfig to be an explicit nil

### UnsetNetworkConfig
`func (o *RecoverGcpVmNewSourceConfig) UnsetNetworkConfig()`

UnsetNetworkConfig ensures that no value is present for NetworkConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


