# RecoverAcropolisVmNewSourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | [**NullableRecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the id of the parent source to recover the VMs. | 
**StorageContainer** | Pointer to [**NullableRecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | A storage container where the VM&#39;s files should be restored to. | [optional] 
**NetworkConfig** | Pointer to [**NullableRecoverAcropolisVmNewSourceNetworkConfig**](RecoverAcropolisVmNewSourceNetworkConfig.md) | Specifies the networking configuration to be applied to the recovered VMs. | [optional] 

## Methods

### NewRecoverAcropolisVmNewSourceConfig

`func NewRecoverAcropolisVmNewSourceConfig(source NullableRecoveryObjectIdentifier, ) *RecoverAcropolisVmNewSourceConfig`

NewRecoverAcropolisVmNewSourceConfig instantiates a new RecoverAcropolisVmNewSourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverAcropolisVmNewSourceConfigWithDefaults

`func NewRecoverAcropolisVmNewSourceConfigWithDefaults() *RecoverAcropolisVmNewSourceConfig`

NewRecoverAcropolisVmNewSourceConfigWithDefaults instantiates a new RecoverAcropolisVmNewSourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *RecoverAcropolisVmNewSourceConfig) GetSource() RecoveryObjectIdentifier`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RecoverAcropolisVmNewSourceConfig) GetSourceOk() (*RecoveryObjectIdentifier, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RecoverAcropolisVmNewSourceConfig) SetSource(v RecoveryObjectIdentifier)`

SetSource sets Source field to given value.


### SetSourceNil

`func (o *RecoverAcropolisVmNewSourceConfig) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *RecoverAcropolisVmNewSourceConfig) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetStorageContainer

`func (o *RecoverAcropolisVmNewSourceConfig) GetStorageContainer() RecoveryObjectIdentifier`

GetStorageContainer returns the StorageContainer field if non-nil, zero value otherwise.

### GetStorageContainerOk

`func (o *RecoverAcropolisVmNewSourceConfig) GetStorageContainerOk() (*RecoveryObjectIdentifier, bool)`

GetStorageContainerOk returns a tuple with the StorageContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageContainer

`func (o *RecoverAcropolisVmNewSourceConfig) SetStorageContainer(v RecoveryObjectIdentifier)`

SetStorageContainer sets StorageContainer field to given value.

### HasStorageContainer

`func (o *RecoverAcropolisVmNewSourceConfig) HasStorageContainer() bool`

HasStorageContainer returns a boolean if a field has been set.

### SetStorageContainerNil

`func (o *RecoverAcropolisVmNewSourceConfig) SetStorageContainerNil(b bool)`

 SetStorageContainerNil sets the value for StorageContainer to be an explicit nil

### UnsetStorageContainer
`func (o *RecoverAcropolisVmNewSourceConfig) UnsetStorageContainer()`

UnsetStorageContainer ensures that no value is present for StorageContainer, not even an explicit nil
### GetNetworkConfig

`func (o *RecoverAcropolisVmNewSourceConfig) GetNetworkConfig() RecoverAcropolisVmNewSourceNetworkConfig`

GetNetworkConfig returns the NetworkConfig field if non-nil, zero value otherwise.

### GetNetworkConfigOk

`func (o *RecoverAcropolisVmNewSourceConfig) GetNetworkConfigOk() (*RecoverAcropolisVmNewSourceNetworkConfig, bool)`

GetNetworkConfigOk returns a tuple with the NetworkConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConfig

`func (o *RecoverAcropolisVmNewSourceConfig) SetNetworkConfig(v RecoverAcropolisVmNewSourceNetworkConfig)`

SetNetworkConfig sets NetworkConfig field to given value.

### HasNetworkConfig

`func (o *RecoverAcropolisVmNewSourceConfig) HasNetworkConfig() bool`

HasNetworkConfig returns a boolean if a field has been set.

### SetNetworkConfigNil

`func (o *RecoverAcropolisVmNewSourceConfig) SetNetworkConfigNil(b bool)`

 SetNetworkConfigNil sets the value for NetworkConfig to be an explicit nil

### UnsetNetworkConfig
`func (o *RecoverAcropolisVmNewSourceConfig) UnsetNetworkConfig()`

UnsetNetworkConfig ensures that no value is present for NetworkConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


