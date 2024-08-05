# RecoverVmwareVmEsxiSourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datastores** | Pointer to [**[]RecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the datastore objects where the object&#39;s files should be recovered to. | [optional] 
**NetworkConfig** | Pointer to [**NullableRecoverVmwareVAppVCDSourceConfigNetworkConfig**](RecoverVmwareVAppVCDSourceConfigNetworkConfig.md) |  | [optional] 
**ResourcePool** | [**NullableRecoverVmwareVmEsxiSourceConfigResourcePool**](RecoverVmwareVmEsxiSourceConfigResourcePool.md) |  | 
**Source** | [**NullableRecoverAcropolisVmNewSourceConfigSource**](RecoverAcropolisVmNewSourceConfigSource.md) |  | 
**VmFolder** | Pointer to [**NullableRecoverVmwareVmEsxiSourceConfigVmFolder**](RecoverVmwareVmEsxiSourceConfigVmFolder.md) |  | [optional] 

## Methods

### NewRecoverVmwareVmEsxiSourceConfig

`func NewRecoverVmwareVmEsxiSourceConfig(resourcePool NullableRecoverVmwareVmEsxiSourceConfigResourcePool, source NullableRecoverAcropolisVmNewSourceConfigSource, ) *RecoverVmwareVmEsxiSourceConfig`

NewRecoverVmwareVmEsxiSourceConfig instantiates a new RecoverVmwareVmEsxiSourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverVmwareVmEsxiSourceConfigWithDefaults

`func NewRecoverVmwareVmEsxiSourceConfigWithDefaults() *RecoverVmwareVmEsxiSourceConfig`

NewRecoverVmwareVmEsxiSourceConfigWithDefaults instantiates a new RecoverVmwareVmEsxiSourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatastores

`func (o *RecoverVmwareVmEsxiSourceConfig) GetDatastores() []RecoveryObjectIdentifier`

GetDatastores returns the Datastores field if non-nil, zero value otherwise.

### GetDatastoresOk

`func (o *RecoverVmwareVmEsxiSourceConfig) GetDatastoresOk() (*[]RecoveryObjectIdentifier, bool)`

GetDatastoresOk returns a tuple with the Datastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastores

`func (o *RecoverVmwareVmEsxiSourceConfig) SetDatastores(v []RecoveryObjectIdentifier)`

SetDatastores sets Datastores field to given value.

### HasDatastores

`func (o *RecoverVmwareVmEsxiSourceConfig) HasDatastores() bool`

HasDatastores returns a boolean if a field has been set.

### SetDatastoresNil

`func (o *RecoverVmwareVmEsxiSourceConfig) SetDatastoresNil(b bool)`

 SetDatastoresNil sets the value for Datastores to be an explicit nil

### UnsetDatastores
`func (o *RecoverVmwareVmEsxiSourceConfig) UnsetDatastores()`

UnsetDatastores ensures that no value is present for Datastores, not even an explicit nil
### GetNetworkConfig

`func (o *RecoverVmwareVmEsxiSourceConfig) GetNetworkConfig() RecoverVmwareVAppVCDSourceConfigNetworkConfig`

GetNetworkConfig returns the NetworkConfig field if non-nil, zero value otherwise.

### GetNetworkConfigOk

`func (o *RecoverVmwareVmEsxiSourceConfig) GetNetworkConfigOk() (*RecoverVmwareVAppVCDSourceConfigNetworkConfig, bool)`

GetNetworkConfigOk returns a tuple with the NetworkConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConfig

`func (o *RecoverVmwareVmEsxiSourceConfig) SetNetworkConfig(v RecoverVmwareVAppVCDSourceConfigNetworkConfig)`

SetNetworkConfig sets NetworkConfig field to given value.

### HasNetworkConfig

`func (o *RecoverVmwareVmEsxiSourceConfig) HasNetworkConfig() bool`

HasNetworkConfig returns a boolean if a field has been set.

### SetNetworkConfigNil

`func (o *RecoverVmwareVmEsxiSourceConfig) SetNetworkConfigNil(b bool)`

 SetNetworkConfigNil sets the value for NetworkConfig to be an explicit nil

### UnsetNetworkConfig
`func (o *RecoverVmwareVmEsxiSourceConfig) UnsetNetworkConfig()`

UnsetNetworkConfig ensures that no value is present for NetworkConfig, not even an explicit nil
### GetResourcePool

`func (o *RecoverVmwareVmEsxiSourceConfig) GetResourcePool() RecoverVmwareVmEsxiSourceConfigResourcePool`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *RecoverVmwareVmEsxiSourceConfig) GetResourcePoolOk() (*RecoverVmwareVmEsxiSourceConfigResourcePool, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *RecoverVmwareVmEsxiSourceConfig) SetResourcePool(v RecoverVmwareVmEsxiSourceConfigResourcePool)`

SetResourcePool sets ResourcePool field to given value.


### SetResourcePoolNil

`func (o *RecoverVmwareVmEsxiSourceConfig) SetResourcePoolNil(b bool)`

 SetResourcePoolNil sets the value for ResourcePool to be an explicit nil

### UnsetResourcePool
`func (o *RecoverVmwareVmEsxiSourceConfig) UnsetResourcePool()`

UnsetResourcePool ensures that no value is present for ResourcePool, not even an explicit nil
### GetSource

`func (o *RecoverVmwareVmEsxiSourceConfig) GetSource() RecoverAcropolisVmNewSourceConfigSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RecoverVmwareVmEsxiSourceConfig) GetSourceOk() (*RecoverAcropolisVmNewSourceConfigSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RecoverVmwareVmEsxiSourceConfig) SetSource(v RecoverAcropolisVmNewSourceConfigSource)`

SetSource sets Source field to given value.


### SetSourceNil

`func (o *RecoverVmwareVmEsxiSourceConfig) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *RecoverVmwareVmEsxiSourceConfig) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetVmFolder

`func (o *RecoverVmwareVmEsxiSourceConfig) GetVmFolder() RecoverVmwareVmEsxiSourceConfigVmFolder`

GetVmFolder returns the VmFolder field if non-nil, zero value otherwise.

### GetVmFolderOk

`func (o *RecoverVmwareVmEsxiSourceConfig) GetVmFolderOk() (*RecoverVmwareVmEsxiSourceConfigVmFolder, bool)`

GetVmFolderOk returns a tuple with the VmFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmFolder

`func (o *RecoverVmwareVmEsxiSourceConfig) SetVmFolder(v RecoverVmwareVmEsxiSourceConfigVmFolder)`

SetVmFolder sets VmFolder field to given value.

### HasVmFolder

`func (o *RecoverVmwareVmEsxiSourceConfig) HasVmFolder() bool`

HasVmFolder returns a boolean if a field has been set.

### SetVmFolderNil

`func (o *RecoverVmwareVmEsxiSourceConfig) SetVmFolderNil(b bool)`

 SetVmFolderNil sets the value for VmFolder to be an explicit nil

### UnsetVmFolder
`func (o *RecoverVmwareVmEsxiSourceConfig) UnsetVmFolder()`

UnsetVmFolder ensures that no value is present for VmFolder, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


