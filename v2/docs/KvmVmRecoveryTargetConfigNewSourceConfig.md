# KvmVmRecoveryTargetConfigNewSourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | [**RecoverKvmVmNewSourceConfigCluster**](RecoverKvmVmNewSourceConfigCluster.md) |  | 
**DataCenter** | [**RecoverKvmVmNewSourceConfigDataCenter**](RecoverKvmVmNewSourceConfigDataCenter.md) |  | 
**NetworkConfig** | Pointer to [**RecoverKvmVmNewSourceConfigNetworkConfig**](RecoverKvmVmNewSourceConfigNetworkConfig.md) |  | [optional] 
**Source** | [**RecoverKvmVmNewSourceConfigSource**](RecoverKvmVmNewSourceConfigSource.md) |  | 
**StorageDomain** | [**RecoverKvmVmNewSourceConfigStorageDomain**](RecoverKvmVmNewSourceConfigStorageDomain.md) |  | 

## Methods

### NewKvmVmRecoveryTargetConfigNewSourceConfig

`func NewKvmVmRecoveryTargetConfigNewSourceConfig(cluster RecoverKvmVmNewSourceConfigCluster, dataCenter RecoverKvmVmNewSourceConfigDataCenter, source RecoverKvmVmNewSourceConfigSource, storageDomain RecoverKvmVmNewSourceConfigStorageDomain, ) *KvmVmRecoveryTargetConfigNewSourceConfig`

NewKvmVmRecoveryTargetConfigNewSourceConfig instantiates a new KvmVmRecoveryTargetConfigNewSourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvmVmRecoveryTargetConfigNewSourceConfigWithDefaults

`func NewKvmVmRecoveryTargetConfigNewSourceConfigWithDefaults() *KvmVmRecoveryTargetConfigNewSourceConfig`

NewKvmVmRecoveryTargetConfigNewSourceConfigWithDefaults instantiates a new KvmVmRecoveryTargetConfigNewSourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *KvmVmRecoveryTargetConfigNewSourceConfig) GetCluster() RecoverKvmVmNewSourceConfigCluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *KvmVmRecoveryTargetConfigNewSourceConfig) GetClusterOk() (*RecoverKvmVmNewSourceConfigCluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *KvmVmRecoveryTargetConfigNewSourceConfig) SetCluster(v RecoverKvmVmNewSourceConfigCluster)`

SetCluster sets Cluster field to given value.


### GetDataCenter

`func (o *KvmVmRecoveryTargetConfigNewSourceConfig) GetDataCenter() RecoverKvmVmNewSourceConfigDataCenter`

GetDataCenter returns the DataCenter field if non-nil, zero value otherwise.

### GetDataCenterOk

`func (o *KvmVmRecoveryTargetConfigNewSourceConfig) GetDataCenterOk() (*RecoverKvmVmNewSourceConfigDataCenter, bool)`

GetDataCenterOk returns a tuple with the DataCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenter

`func (o *KvmVmRecoveryTargetConfigNewSourceConfig) SetDataCenter(v RecoverKvmVmNewSourceConfigDataCenter)`

SetDataCenter sets DataCenter field to given value.


### GetNetworkConfig

`func (o *KvmVmRecoveryTargetConfigNewSourceConfig) GetNetworkConfig() RecoverKvmVmNewSourceConfigNetworkConfig`

GetNetworkConfig returns the NetworkConfig field if non-nil, zero value otherwise.

### GetNetworkConfigOk

`func (o *KvmVmRecoveryTargetConfigNewSourceConfig) GetNetworkConfigOk() (*RecoverKvmVmNewSourceConfigNetworkConfig, bool)`

GetNetworkConfigOk returns a tuple with the NetworkConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConfig

`func (o *KvmVmRecoveryTargetConfigNewSourceConfig) SetNetworkConfig(v RecoverKvmVmNewSourceConfigNetworkConfig)`

SetNetworkConfig sets NetworkConfig field to given value.

### HasNetworkConfig

`func (o *KvmVmRecoveryTargetConfigNewSourceConfig) HasNetworkConfig() bool`

HasNetworkConfig returns a boolean if a field has been set.

### GetSource

`func (o *KvmVmRecoveryTargetConfigNewSourceConfig) GetSource() RecoverKvmVmNewSourceConfigSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *KvmVmRecoveryTargetConfigNewSourceConfig) GetSourceOk() (*RecoverKvmVmNewSourceConfigSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *KvmVmRecoveryTargetConfigNewSourceConfig) SetSource(v RecoverKvmVmNewSourceConfigSource)`

SetSource sets Source field to given value.


### GetStorageDomain

`func (o *KvmVmRecoveryTargetConfigNewSourceConfig) GetStorageDomain() RecoverKvmVmNewSourceConfigStorageDomain`

GetStorageDomain returns the StorageDomain field if non-nil, zero value otherwise.

### GetStorageDomainOk

`func (o *KvmVmRecoveryTargetConfigNewSourceConfig) GetStorageDomainOk() (*RecoverKvmVmNewSourceConfigStorageDomain, bool)`

GetStorageDomainOk returns a tuple with the StorageDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomain

`func (o *KvmVmRecoveryTargetConfigNewSourceConfig) SetStorageDomain(v RecoverKvmVmNewSourceConfigStorageDomain)`

SetStorageDomain sets StorageDomain field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


