# AcropolisRestoreParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisableNetwork** | Pointer to **NullableBool** | Specifies whether the network should be left in disabled state. Attached network is enabled by default. Set this flag to true to disable it. | [optional] 
**NetworkId** | Pointer to **NullableInt64** | Specifies a network configuration to be attached to the cloned or recovered object. For kCloneVMs and kRecoverVMs tasks, original network configuration is detached if the cloned or recovered object is kept under a different parent Protection Source or a different Resource Pool. By default, for kRecoverVMs task, original network configuration is preserved if the recovered object is kept under the same parent Protection Source and the same Resource Pool. Specify this field to override the preserved network configuration or to attach a new network configuration to the cloned or recovered objects. You can get the networkId of the kNetwork object by setting includeNetworks to &#39;true&#39; in the GET /public/protectionSources operation. In the response, get the id of the desired kNetwork object, the resource pool, and the registered parent Protection Source. | [optional] 
**PoweredOn** | Pointer to **NullableBool** | Specifies the power state of the cloned or recovered objects. By default, the cloned or recovered objects are powered off. | [optional] 
**Prefix** | Pointer to **NullableString** | Specifies a prefix to prepended to the source object name to derive a new name for the recovered or cloned object. By default, cloned or recovered objects retain their original name. Length of this field is limited to 8 characters. | [optional] 
**StorageContainerId** | Pointer to **NullableInt64** | A storage container where the VM&#39;s files should be restored to. This field is optional if the VM is being restored to its original parent source. If not specified, the VM&#39;s files will be restored to their original storage container(s). This field is mandatory if the VMs are being restored to a different parent source. | [optional] 
**Suffix** | Pointer to **NullableString** | Specifies a suffix to appended to the original source object name to derive a new name for the recovered or cloned object. By default, cloned or recovered objects retain their original name. Length of this field is limited to 8 characters. | [optional] 

## Methods

### NewAcropolisRestoreParameters

`func NewAcropolisRestoreParameters() *AcropolisRestoreParameters`

NewAcropolisRestoreParameters instantiates a new AcropolisRestoreParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcropolisRestoreParametersWithDefaults

`func NewAcropolisRestoreParametersWithDefaults() *AcropolisRestoreParameters`

NewAcropolisRestoreParametersWithDefaults instantiates a new AcropolisRestoreParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisableNetwork

`func (o *AcropolisRestoreParameters) GetDisableNetwork() bool`

GetDisableNetwork returns the DisableNetwork field if non-nil, zero value otherwise.

### GetDisableNetworkOk

`func (o *AcropolisRestoreParameters) GetDisableNetworkOk() (*bool, bool)`

GetDisableNetworkOk returns a tuple with the DisableNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableNetwork

`func (o *AcropolisRestoreParameters) SetDisableNetwork(v bool)`

SetDisableNetwork sets DisableNetwork field to given value.

### HasDisableNetwork

`func (o *AcropolisRestoreParameters) HasDisableNetwork() bool`

HasDisableNetwork returns a boolean if a field has been set.

### SetDisableNetworkNil

`func (o *AcropolisRestoreParameters) SetDisableNetworkNil(b bool)`

 SetDisableNetworkNil sets the value for DisableNetwork to be an explicit nil

### UnsetDisableNetwork
`func (o *AcropolisRestoreParameters) UnsetDisableNetwork()`

UnsetDisableNetwork ensures that no value is present for DisableNetwork, not even an explicit nil
### GetNetworkId

`func (o *AcropolisRestoreParameters) GetNetworkId() int64`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *AcropolisRestoreParameters) GetNetworkIdOk() (*int64, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *AcropolisRestoreParameters) SetNetworkId(v int64)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *AcropolisRestoreParameters) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### SetNetworkIdNil

`func (o *AcropolisRestoreParameters) SetNetworkIdNil(b bool)`

 SetNetworkIdNil sets the value for NetworkId to be an explicit nil

### UnsetNetworkId
`func (o *AcropolisRestoreParameters) UnsetNetworkId()`

UnsetNetworkId ensures that no value is present for NetworkId, not even an explicit nil
### GetPoweredOn

`func (o *AcropolisRestoreParameters) GetPoweredOn() bool`

GetPoweredOn returns the PoweredOn field if non-nil, zero value otherwise.

### GetPoweredOnOk

`func (o *AcropolisRestoreParameters) GetPoweredOnOk() (*bool, bool)`

GetPoweredOnOk returns a tuple with the PoweredOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoweredOn

`func (o *AcropolisRestoreParameters) SetPoweredOn(v bool)`

SetPoweredOn sets PoweredOn field to given value.

### HasPoweredOn

`func (o *AcropolisRestoreParameters) HasPoweredOn() bool`

HasPoweredOn returns a boolean if a field has been set.

### SetPoweredOnNil

`func (o *AcropolisRestoreParameters) SetPoweredOnNil(b bool)`

 SetPoweredOnNil sets the value for PoweredOn to be an explicit nil

### UnsetPoweredOn
`func (o *AcropolisRestoreParameters) UnsetPoweredOn()`

UnsetPoweredOn ensures that no value is present for PoweredOn, not even an explicit nil
### GetPrefix

`func (o *AcropolisRestoreParameters) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *AcropolisRestoreParameters) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *AcropolisRestoreParameters) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *AcropolisRestoreParameters) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### SetPrefixNil

`func (o *AcropolisRestoreParameters) SetPrefixNil(b bool)`

 SetPrefixNil sets the value for Prefix to be an explicit nil

### UnsetPrefix
`func (o *AcropolisRestoreParameters) UnsetPrefix()`

UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil
### GetStorageContainerId

`func (o *AcropolisRestoreParameters) GetStorageContainerId() int64`

GetStorageContainerId returns the StorageContainerId field if non-nil, zero value otherwise.

### GetStorageContainerIdOk

`func (o *AcropolisRestoreParameters) GetStorageContainerIdOk() (*int64, bool)`

GetStorageContainerIdOk returns a tuple with the StorageContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageContainerId

`func (o *AcropolisRestoreParameters) SetStorageContainerId(v int64)`

SetStorageContainerId sets StorageContainerId field to given value.

### HasStorageContainerId

`func (o *AcropolisRestoreParameters) HasStorageContainerId() bool`

HasStorageContainerId returns a boolean if a field has been set.

### SetStorageContainerIdNil

`func (o *AcropolisRestoreParameters) SetStorageContainerIdNil(b bool)`

 SetStorageContainerIdNil sets the value for StorageContainerId to be an explicit nil

### UnsetStorageContainerId
`func (o *AcropolisRestoreParameters) UnsetStorageContainerId()`

UnsetStorageContainerId ensures that no value is present for StorageContainerId, not even an explicit nil
### GetSuffix

`func (o *AcropolisRestoreParameters) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *AcropolisRestoreParameters) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *AcropolisRestoreParameters) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *AcropolisRestoreParameters) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.

### SetSuffixNil

`func (o *AcropolisRestoreParameters) SetSuffixNil(b bool)`

 SetSuffixNil sets the value for Suffix to be an explicit nil

### UnsetSuffix
`func (o *AcropolisRestoreParameters) UnsetSuffix()`

UnsetSuffix ensures that no value is present for Suffix, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


