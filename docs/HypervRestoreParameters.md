# HypervRestoreParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatastoreId** | Pointer to **NullableInt64** | A datastore entity where the object&#39;s files should be restored to. This field is optional if object is being restored to its original parent source. If not specified, the object&#39;s files will be restored to their original datastore locations. This field is mandatory if object is being restored to a different resource entity or to a different parent source. | [optional] 
**DisableNetwork** | Pointer to **NullableBool** | Specifies whether the network should be left in disabled state. Attached network is enabled by default. Set this flag to true to disable it. | [optional] 
**NetworkId** | Pointer to **NullableInt64** | Specifies a network configuration to be attached to the cloned or recovered object. For kCloneVMs and kRecoverVMs tasks, original network configuration is detached if the cloned or recovered object is kept under a different parent Protection Source or a different Resource Pool. By default, for kRecoverVMs task, original network configuration is preserved if the recovered object is kept under the same parent Protection Source and the same Resource Pool. Specify this field to override the preserved network configuration or to attach a new network configuration to the cloned or recovered objects. You can get the networkId of the kNetwork object by setting includeNetworks to &#39;true&#39; in the GET /public/protectionSources operation. In the response, get the id of the desired kNetwork object, the resource pool, and the registered parent Protection Source. | [optional] 
**PoweredOn** | Pointer to **NullableBool** | Specifies the power state of the cloned or recovered objects. By default, the cloned or recovered objects are powered off. | [optional] 
**Prefix** | Pointer to **NullableString** | Specifies a prefix to prepended to the source object name to derive a new name for the recovered or cloned object. By default, cloned or recovered objects retain their original name. Length of this field is limited to 8 characters. | [optional] 
**PreserveTags** | Pointer to **NullableBool** | Specifies whether or not to preserve tags during the operation. | [optional] 
**ResourceId** | Pointer to **NullableInt64** | The resource (HyperV host) to which the restored VM will be attached.  This field is optional for a kRecoverVMs task if the VMs are being restored to its original parent source. If not specified, restored VMs will be attached to its original host. This field is mandatory if the VMs are being restored to a different parent source. | [optional] 
**Suffix** | Pointer to **NullableString** | Specifies a suffix to appended to the original source object name to derive a new name for the recovered or cloned object. By default, cloned or recovered objects retain their original name. Length of this field is limited to 8 characters. | [optional] 

## Methods

### NewHypervRestoreParameters

`func NewHypervRestoreParameters() *HypervRestoreParameters`

NewHypervRestoreParameters instantiates a new HypervRestoreParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervRestoreParametersWithDefaults

`func NewHypervRestoreParametersWithDefaults() *HypervRestoreParameters`

NewHypervRestoreParametersWithDefaults instantiates a new HypervRestoreParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatastoreId

`func (o *HypervRestoreParameters) GetDatastoreId() int64`

GetDatastoreId returns the DatastoreId field if non-nil, zero value otherwise.

### GetDatastoreIdOk

`func (o *HypervRestoreParameters) GetDatastoreIdOk() (*int64, bool)`

GetDatastoreIdOk returns a tuple with the DatastoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreId

`func (o *HypervRestoreParameters) SetDatastoreId(v int64)`

SetDatastoreId sets DatastoreId field to given value.

### HasDatastoreId

`func (o *HypervRestoreParameters) HasDatastoreId() bool`

HasDatastoreId returns a boolean if a field has been set.

### SetDatastoreIdNil

`func (o *HypervRestoreParameters) SetDatastoreIdNil(b bool)`

 SetDatastoreIdNil sets the value for DatastoreId to be an explicit nil

### UnsetDatastoreId
`func (o *HypervRestoreParameters) UnsetDatastoreId()`

UnsetDatastoreId ensures that no value is present for DatastoreId, not even an explicit nil
### GetDisableNetwork

`func (o *HypervRestoreParameters) GetDisableNetwork() bool`

GetDisableNetwork returns the DisableNetwork field if non-nil, zero value otherwise.

### GetDisableNetworkOk

`func (o *HypervRestoreParameters) GetDisableNetworkOk() (*bool, bool)`

GetDisableNetworkOk returns a tuple with the DisableNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableNetwork

`func (o *HypervRestoreParameters) SetDisableNetwork(v bool)`

SetDisableNetwork sets DisableNetwork field to given value.

### HasDisableNetwork

`func (o *HypervRestoreParameters) HasDisableNetwork() bool`

HasDisableNetwork returns a boolean if a field has been set.

### SetDisableNetworkNil

`func (o *HypervRestoreParameters) SetDisableNetworkNil(b bool)`

 SetDisableNetworkNil sets the value for DisableNetwork to be an explicit nil

### UnsetDisableNetwork
`func (o *HypervRestoreParameters) UnsetDisableNetwork()`

UnsetDisableNetwork ensures that no value is present for DisableNetwork, not even an explicit nil
### GetNetworkId

`func (o *HypervRestoreParameters) GetNetworkId() int64`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *HypervRestoreParameters) GetNetworkIdOk() (*int64, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *HypervRestoreParameters) SetNetworkId(v int64)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *HypervRestoreParameters) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### SetNetworkIdNil

`func (o *HypervRestoreParameters) SetNetworkIdNil(b bool)`

 SetNetworkIdNil sets the value for NetworkId to be an explicit nil

### UnsetNetworkId
`func (o *HypervRestoreParameters) UnsetNetworkId()`

UnsetNetworkId ensures that no value is present for NetworkId, not even an explicit nil
### GetPoweredOn

`func (o *HypervRestoreParameters) GetPoweredOn() bool`

GetPoweredOn returns the PoweredOn field if non-nil, zero value otherwise.

### GetPoweredOnOk

`func (o *HypervRestoreParameters) GetPoweredOnOk() (*bool, bool)`

GetPoweredOnOk returns a tuple with the PoweredOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoweredOn

`func (o *HypervRestoreParameters) SetPoweredOn(v bool)`

SetPoweredOn sets PoweredOn field to given value.

### HasPoweredOn

`func (o *HypervRestoreParameters) HasPoweredOn() bool`

HasPoweredOn returns a boolean if a field has been set.

### SetPoweredOnNil

`func (o *HypervRestoreParameters) SetPoweredOnNil(b bool)`

 SetPoweredOnNil sets the value for PoweredOn to be an explicit nil

### UnsetPoweredOn
`func (o *HypervRestoreParameters) UnsetPoweredOn()`

UnsetPoweredOn ensures that no value is present for PoweredOn, not even an explicit nil
### GetPrefix

`func (o *HypervRestoreParameters) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *HypervRestoreParameters) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *HypervRestoreParameters) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *HypervRestoreParameters) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### SetPrefixNil

`func (o *HypervRestoreParameters) SetPrefixNil(b bool)`

 SetPrefixNil sets the value for Prefix to be an explicit nil

### UnsetPrefix
`func (o *HypervRestoreParameters) UnsetPrefix()`

UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil
### GetPreserveTags

`func (o *HypervRestoreParameters) GetPreserveTags() bool`

GetPreserveTags returns the PreserveTags field if non-nil, zero value otherwise.

### GetPreserveTagsOk

`func (o *HypervRestoreParameters) GetPreserveTagsOk() (*bool, bool)`

GetPreserveTagsOk returns a tuple with the PreserveTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveTags

`func (o *HypervRestoreParameters) SetPreserveTags(v bool)`

SetPreserveTags sets PreserveTags field to given value.

### HasPreserveTags

`func (o *HypervRestoreParameters) HasPreserveTags() bool`

HasPreserveTags returns a boolean if a field has been set.

### SetPreserveTagsNil

`func (o *HypervRestoreParameters) SetPreserveTagsNil(b bool)`

 SetPreserveTagsNil sets the value for PreserveTags to be an explicit nil

### UnsetPreserveTags
`func (o *HypervRestoreParameters) UnsetPreserveTags()`

UnsetPreserveTags ensures that no value is present for PreserveTags, not even an explicit nil
### GetResourceId

`func (o *HypervRestoreParameters) GetResourceId() int64`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *HypervRestoreParameters) GetResourceIdOk() (*int64, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *HypervRestoreParameters) SetResourceId(v int64)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *HypervRestoreParameters) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### SetResourceIdNil

`func (o *HypervRestoreParameters) SetResourceIdNil(b bool)`

 SetResourceIdNil sets the value for ResourceId to be an explicit nil

### UnsetResourceId
`func (o *HypervRestoreParameters) UnsetResourceId()`

UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil
### GetSuffix

`func (o *HypervRestoreParameters) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *HypervRestoreParameters) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *HypervRestoreParameters) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *HypervRestoreParameters) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.

### SetSuffixNil

`func (o *HypervRestoreParameters) SetSuffixNil(b bool)`

 SetSuffixNil sets the value for Suffix to be an explicit nil

### UnsetSuffix
`func (o *HypervRestoreParameters) UnsetSuffix()`

UnsetSuffix ensures that no value is present for Suffix, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


