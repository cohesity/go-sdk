# GcpVmRecoveryTargetConfigNewSourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailabilityZone** | [**NullableRecoverGcpVmNewSourceConfigAvailabilityZone**](RecoverGcpVmNewSourceConfigAvailabilityZone.md) |  | 
**NetworkConfig** | [**NullableRecoverGcpVmNewSourceConfigNetworkConfig**](RecoverGcpVmNewSourceConfigNetworkConfig.md) |  | 
**Project** | [**NullableRecoverGcpVmNewSourceConfigProject**](RecoverGcpVmNewSourceConfigProject.md) |  | 
**Region** | [**NullableRecoverGcpVmNewSourceConfigRegion**](RecoverGcpVmNewSourceConfigRegion.md) |  | 
**Source** | [**NullableRecoverAcropolisVmNewSourceConfigSource**](RecoverAcropolisVmNewSourceConfigSource.md) |  | 

## Methods

### NewGcpVmRecoveryTargetConfigNewSourceConfig

`func NewGcpVmRecoveryTargetConfigNewSourceConfig(availabilityZone NullableRecoverGcpVmNewSourceConfigAvailabilityZone, networkConfig NullableRecoverGcpVmNewSourceConfigNetworkConfig, project NullableRecoverGcpVmNewSourceConfigProject, region NullableRecoverGcpVmNewSourceConfigRegion, source NullableRecoverAcropolisVmNewSourceConfigSource, ) *GcpVmRecoveryTargetConfigNewSourceConfig`

NewGcpVmRecoveryTargetConfigNewSourceConfig instantiates a new GcpVmRecoveryTargetConfigNewSourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpVmRecoveryTargetConfigNewSourceConfigWithDefaults

`func NewGcpVmRecoveryTargetConfigNewSourceConfigWithDefaults() *GcpVmRecoveryTargetConfigNewSourceConfig`

NewGcpVmRecoveryTargetConfigNewSourceConfigWithDefaults instantiates a new GcpVmRecoveryTargetConfigNewSourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailabilityZone

`func (o *GcpVmRecoveryTargetConfigNewSourceConfig) GetAvailabilityZone() RecoverGcpVmNewSourceConfigAvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *GcpVmRecoveryTargetConfigNewSourceConfig) GetAvailabilityZoneOk() (*RecoverGcpVmNewSourceConfigAvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *GcpVmRecoveryTargetConfigNewSourceConfig) SetAvailabilityZone(v RecoverGcpVmNewSourceConfigAvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.


### SetAvailabilityZoneNil

`func (o *GcpVmRecoveryTargetConfigNewSourceConfig) SetAvailabilityZoneNil(b bool)`

 SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil

### UnsetAvailabilityZone
`func (o *GcpVmRecoveryTargetConfigNewSourceConfig) UnsetAvailabilityZone()`

UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil
### GetNetworkConfig

`func (o *GcpVmRecoveryTargetConfigNewSourceConfig) GetNetworkConfig() RecoverGcpVmNewSourceConfigNetworkConfig`

GetNetworkConfig returns the NetworkConfig field if non-nil, zero value otherwise.

### GetNetworkConfigOk

`func (o *GcpVmRecoveryTargetConfigNewSourceConfig) GetNetworkConfigOk() (*RecoverGcpVmNewSourceConfigNetworkConfig, bool)`

GetNetworkConfigOk returns a tuple with the NetworkConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConfig

`func (o *GcpVmRecoveryTargetConfigNewSourceConfig) SetNetworkConfig(v RecoverGcpVmNewSourceConfigNetworkConfig)`

SetNetworkConfig sets NetworkConfig field to given value.


### SetNetworkConfigNil

`func (o *GcpVmRecoveryTargetConfigNewSourceConfig) SetNetworkConfigNil(b bool)`

 SetNetworkConfigNil sets the value for NetworkConfig to be an explicit nil

### UnsetNetworkConfig
`func (o *GcpVmRecoveryTargetConfigNewSourceConfig) UnsetNetworkConfig()`

UnsetNetworkConfig ensures that no value is present for NetworkConfig, not even an explicit nil
### GetProject

`func (o *GcpVmRecoveryTargetConfigNewSourceConfig) GetProject() RecoverGcpVmNewSourceConfigProject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *GcpVmRecoveryTargetConfigNewSourceConfig) GetProjectOk() (*RecoverGcpVmNewSourceConfigProject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *GcpVmRecoveryTargetConfigNewSourceConfig) SetProject(v RecoverGcpVmNewSourceConfigProject)`

SetProject sets Project field to given value.


### SetProjectNil

`func (o *GcpVmRecoveryTargetConfigNewSourceConfig) SetProjectNil(b bool)`

 SetProjectNil sets the value for Project to be an explicit nil

### UnsetProject
`func (o *GcpVmRecoveryTargetConfigNewSourceConfig) UnsetProject()`

UnsetProject ensures that no value is present for Project, not even an explicit nil
### GetRegion

`func (o *GcpVmRecoveryTargetConfigNewSourceConfig) GetRegion() RecoverGcpVmNewSourceConfigRegion`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GcpVmRecoveryTargetConfigNewSourceConfig) GetRegionOk() (*RecoverGcpVmNewSourceConfigRegion, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GcpVmRecoveryTargetConfigNewSourceConfig) SetRegion(v RecoverGcpVmNewSourceConfigRegion)`

SetRegion sets Region field to given value.


### SetRegionNil

`func (o *GcpVmRecoveryTargetConfigNewSourceConfig) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *GcpVmRecoveryTargetConfigNewSourceConfig) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetSource

`func (o *GcpVmRecoveryTargetConfigNewSourceConfig) GetSource() RecoverAcropolisVmNewSourceConfigSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GcpVmRecoveryTargetConfigNewSourceConfig) GetSourceOk() (*RecoverAcropolisVmNewSourceConfigSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GcpVmRecoveryTargetConfigNewSourceConfig) SetSource(v RecoverAcropolisVmNewSourceConfigSource)`

SetSource sets Source field to given value.


### SetSourceNil

`func (o *GcpVmRecoveryTargetConfigNewSourceConfig) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *GcpVmRecoveryTargetConfigNewSourceConfig) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


