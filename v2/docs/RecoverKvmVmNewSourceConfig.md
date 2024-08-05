# RecoverKvmVmNewSourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | [**RecoverKvmVmNewSourceConfigCluster**](RecoverKvmVmNewSourceConfigCluster.md) |  | 
**DataCenter** | [**RecoverKvmVmNewSourceConfigDataCenter**](RecoverKvmVmNewSourceConfigDataCenter.md) |  | 
**NetworkConfig** | Pointer to [**RecoverKvmVmNewSourceConfigNetworkConfig**](RecoverKvmVmNewSourceConfigNetworkConfig.md) |  | [optional] 
**Source** | [**RecoverKvmVmNewSourceConfigSource**](RecoverKvmVmNewSourceConfigSource.md) |  | 
**StorageDomain** | [**RecoverKvmVmNewSourceConfigStorageDomain**](RecoverKvmVmNewSourceConfigStorageDomain.md) |  | 

## Methods

### NewRecoverKvmVmNewSourceConfig

`func NewRecoverKvmVmNewSourceConfig(cluster RecoverKvmVmNewSourceConfigCluster, dataCenter RecoverKvmVmNewSourceConfigDataCenter, source RecoverKvmVmNewSourceConfigSource, storageDomain RecoverKvmVmNewSourceConfigStorageDomain, ) *RecoverKvmVmNewSourceConfig`

NewRecoverKvmVmNewSourceConfig instantiates a new RecoverKvmVmNewSourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverKvmVmNewSourceConfigWithDefaults

`func NewRecoverKvmVmNewSourceConfigWithDefaults() *RecoverKvmVmNewSourceConfig`

NewRecoverKvmVmNewSourceConfigWithDefaults instantiates a new RecoverKvmVmNewSourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *RecoverKvmVmNewSourceConfig) GetCluster() RecoverKvmVmNewSourceConfigCluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *RecoverKvmVmNewSourceConfig) GetClusterOk() (*RecoverKvmVmNewSourceConfigCluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *RecoverKvmVmNewSourceConfig) SetCluster(v RecoverKvmVmNewSourceConfigCluster)`

SetCluster sets Cluster field to given value.


### GetDataCenter

`func (o *RecoverKvmVmNewSourceConfig) GetDataCenter() RecoverKvmVmNewSourceConfigDataCenter`

GetDataCenter returns the DataCenter field if non-nil, zero value otherwise.

### GetDataCenterOk

`func (o *RecoverKvmVmNewSourceConfig) GetDataCenterOk() (*RecoverKvmVmNewSourceConfigDataCenter, bool)`

GetDataCenterOk returns a tuple with the DataCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenter

`func (o *RecoverKvmVmNewSourceConfig) SetDataCenter(v RecoverKvmVmNewSourceConfigDataCenter)`

SetDataCenter sets DataCenter field to given value.


### GetNetworkConfig

`func (o *RecoverKvmVmNewSourceConfig) GetNetworkConfig() RecoverKvmVmNewSourceConfigNetworkConfig`

GetNetworkConfig returns the NetworkConfig field if non-nil, zero value otherwise.

### GetNetworkConfigOk

`func (o *RecoverKvmVmNewSourceConfig) GetNetworkConfigOk() (*RecoverKvmVmNewSourceConfigNetworkConfig, bool)`

GetNetworkConfigOk returns a tuple with the NetworkConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConfig

`func (o *RecoverKvmVmNewSourceConfig) SetNetworkConfig(v RecoverKvmVmNewSourceConfigNetworkConfig)`

SetNetworkConfig sets NetworkConfig field to given value.

### HasNetworkConfig

`func (o *RecoverKvmVmNewSourceConfig) HasNetworkConfig() bool`

HasNetworkConfig returns a boolean if a field has been set.

### GetSource

`func (o *RecoverKvmVmNewSourceConfig) GetSource() RecoverKvmVmNewSourceConfigSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RecoverKvmVmNewSourceConfig) GetSourceOk() (*RecoverKvmVmNewSourceConfigSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RecoverKvmVmNewSourceConfig) SetSource(v RecoverKvmVmNewSourceConfigSource)`

SetSource sets Source field to given value.


### GetStorageDomain

`func (o *RecoverKvmVmNewSourceConfig) GetStorageDomain() RecoverKvmVmNewSourceConfigStorageDomain`

GetStorageDomain returns the StorageDomain field if non-nil, zero value otherwise.

### GetStorageDomainOk

`func (o *RecoverKvmVmNewSourceConfig) GetStorageDomainOk() (*RecoverKvmVmNewSourceConfigStorageDomain, bool)`

GetStorageDomainOk returns a tuple with the StorageDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomain

`func (o *RecoverKvmVmNewSourceConfig) SetStorageDomain(v RecoverKvmVmNewSourceConfigStorageDomain)`

SetStorageDomain sets StorageDomain field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


