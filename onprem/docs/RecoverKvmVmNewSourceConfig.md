# RecoverKvmVmNewSourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | [**RecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the id of the parent source to recover the VMs. | 
**Cluster** | [**RecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the resource (KVMH host) to which the restored VM will be attached | 
**DataCenter** | [**RecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the datacenter where the VM&#39;s files should be restored to. | 
**StorageDomain** | [**RecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the Storage Domain where the VM&#39;s disk should be restored to. | 
**NetworkConfig** | Pointer to [**RecoverKvmVmNewSourceNetworkConfig**](RecoverKvmVmNewSourceNetworkConfig.md) | Specifies the networking configuration to be applied to the recovered VMs. | [optional] 

## Methods

### NewRecoverKvmVmNewSourceConfig

`func NewRecoverKvmVmNewSourceConfig(source RecoveryObjectIdentifier, cluster RecoveryObjectIdentifier, dataCenter RecoveryObjectIdentifier, storageDomain RecoveryObjectIdentifier, ) *RecoverKvmVmNewSourceConfig`

NewRecoverKvmVmNewSourceConfig instantiates a new RecoverKvmVmNewSourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverKvmVmNewSourceConfigWithDefaults

`func NewRecoverKvmVmNewSourceConfigWithDefaults() *RecoverKvmVmNewSourceConfig`

NewRecoverKvmVmNewSourceConfigWithDefaults instantiates a new RecoverKvmVmNewSourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *RecoverKvmVmNewSourceConfig) GetSource() RecoveryObjectIdentifier`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RecoverKvmVmNewSourceConfig) GetSourceOk() (*RecoveryObjectIdentifier, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RecoverKvmVmNewSourceConfig) SetSource(v RecoveryObjectIdentifier)`

SetSource sets Source field to given value.


### GetCluster

`func (o *RecoverKvmVmNewSourceConfig) GetCluster() RecoveryObjectIdentifier`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *RecoverKvmVmNewSourceConfig) GetClusterOk() (*RecoveryObjectIdentifier, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *RecoverKvmVmNewSourceConfig) SetCluster(v RecoveryObjectIdentifier)`

SetCluster sets Cluster field to given value.


### GetDataCenter

`func (o *RecoverKvmVmNewSourceConfig) GetDataCenter() RecoveryObjectIdentifier`

GetDataCenter returns the DataCenter field if non-nil, zero value otherwise.

### GetDataCenterOk

`func (o *RecoverKvmVmNewSourceConfig) GetDataCenterOk() (*RecoveryObjectIdentifier, bool)`

GetDataCenterOk returns a tuple with the DataCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenter

`func (o *RecoverKvmVmNewSourceConfig) SetDataCenter(v RecoveryObjectIdentifier)`

SetDataCenter sets DataCenter field to given value.


### GetStorageDomain

`func (o *RecoverKvmVmNewSourceConfig) GetStorageDomain() RecoveryObjectIdentifier`

GetStorageDomain returns the StorageDomain field if non-nil, zero value otherwise.

### GetStorageDomainOk

`func (o *RecoverKvmVmNewSourceConfig) GetStorageDomainOk() (*RecoveryObjectIdentifier, bool)`

GetStorageDomainOk returns a tuple with the StorageDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomain

`func (o *RecoverKvmVmNewSourceConfig) SetStorageDomain(v RecoveryObjectIdentifier)`

SetStorageDomain sets StorageDomain field to given value.


### GetNetworkConfig

`func (o *RecoverKvmVmNewSourceConfig) GetNetworkConfig() RecoverKvmVmNewSourceNetworkConfig`

GetNetworkConfig returns the NetworkConfig field if non-nil, zero value otherwise.

### GetNetworkConfigOk

`func (o *RecoverKvmVmNewSourceConfig) GetNetworkConfigOk() (*RecoverKvmVmNewSourceNetworkConfig, bool)`

GetNetworkConfigOk returns a tuple with the NetworkConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConfig

`func (o *RecoverKvmVmNewSourceConfig) SetNetworkConfig(v RecoverKvmVmNewSourceNetworkConfig)`

SetNetworkConfig sets NetworkConfig field to given value.

### HasNetworkConfig

`func (o *RecoverKvmVmNewSourceConfig) HasNetworkConfig() bool`

HasNetworkConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


