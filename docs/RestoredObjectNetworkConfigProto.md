# RestoredObjectNetworkConfigProto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DetachNetwork** | Pointer to **NullableBool** | If this is set to true, then the network will be detached from the recovered or cloned VMs. NOTE: If this is set to true, then all the following fields will be ignored. | [optional] 
**DisableNetwork** | Pointer to **NullableBool** | This can be set to true to indicate that the attached network should be left in disabled state. | [optional] 
**Mappings** | Pointer to [**[]NetworkMappingProto**](NetworkMappingProto.md) | The network mappings to be applied to the target object. | [optional] 
**NetworkEntity** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**PreserveMacAddressOnNewNetwork** | Pointer to **NullableBool** | If this is true and we are attaching to a new network entity, then the VM&#39;s MAC address will be preserved on the new network. | [optional] 
**VnicEntity** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 

## Methods

### NewRestoredObjectNetworkConfigProto

`func NewRestoredObjectNetworkConfigProto() *RestoredObjectNetworkConfigProto`

NewRestoredObjectNetworkConfigProto instantiates a new RestoredObjectNetworkConfigProto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoredObjectNetworkConfigProtoWithDefaults

`func NewRestoredObjectNetworkConfigProtoWithDefaults() *RestoredObjectNetworkConfigProto`

NewRestoredObjectNetworkConfigProtoWithDefaults instantiates a new RestoredObjectNetworkConfigProto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetachNetwork

`func (o *RestoredObjectNetworkConfigProto) GetDetachNetwork() bool`

GetDetachNetwork returns the DetachNetwork field if non-nil, zero value otherwise.

### GetDetachNetworkOk

`func (o *RestoredObjectNetworkConfigProto) GetDetachNetworkOk() (*bool, bool)`

GetDetachNetworkOk returns a tuple with the DetachNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetachNetwork

`func (o *RestoredObjectNetworkConfigProto) SetDetachNetwork(v bool)`

SetDetachNetwork sets DetachNetwork field to given value.

### HasDetachNetwork

`func (o *RestoredObjectNetworkConfigProto) HasDetachNetwork() bool`

HasDetachNetwork returns a boolean if a field has been set.

### SetDetachNetworkNil

`func (o *RestoredObjectNetworkConfigProto) SetDetachNetworkNil(b bool)`

 SetDetachNetworkNil sets the value for DetachNetwork to be an explicit nil

### UnsetDetachNetwork
`func (o *RestoredObjectNetworkConfigProto) UnsetDetachNetwork()`

UnsetDetachNetwork ensures that no value is present for DetachNetwork, not even an explicit nil
### GetDisableNetwork

`func (o *RestoredObjectNetworkConfigProto) GetDisableNetwork() bool`

GetDisableNetwork returns the DisableNetwork field if non-nil, zero value otherwise.

### GetDisableNetworkOk

`func (o *RestoredObjectNetworkConfigProto) GetDisableNetworkOk() (*bool, bool)`

GetDisableNetworkOk returns a tuple with the DisableNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableNetwork

`func (o *RestoredObjectNetworkConfigProto) SetDisableNetwork(v bool)`

SetDisableNetwork sets DisableNetwork field to given value.

### HasDisableNetwork

`func (o *RestoredObjectNetworkConfigProto) HasDisableNetwork() bool`

HasDisableNetwork returns a boolean if a field has been set.

### SetDisableNetworkNil

`func (o *RestoredObjectNetworkConfigProto) SetDisableNetworkNil(b bool)`

 SetDisableNetworkNil sets the value for DisableNetwork to be an explicit nil

### UnsetDisableNetwork
`func (o *RestoredObjectNetworkConfigProto) UnsetDisableNetwork()`

UnsetDisableNetwork ensures that no value is present for DisableNetwork, not even an explicit nil
### GetMappings

`func (o *RestoredObjectNetworkConfigProto) GetMappings() []NetworkMappingProto`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *RestoredObjectNetworkConfigProto) GetMappingsOk() (*[]NetworkMappingProto, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *RestoredObjectNetworkConfigProto) SetMappings(v []NetworkMappingProto)`

SetMappings sets Mappings field to given value.

### HasMappings

`func (o *RestoredObjectNetworkConfigProto) HasMappings() bool`

HasMappings returns a boolean if a field has been set.

### SetMappingsNil

`func (o *RestoredObjectNetworkConfigProto) SetMappingsNil(b bool)`

 SetMappingsNil sets the value for Mappings to be an explicit nil

### UnsetMappings
`func (o *RestoredObjectNetworkConfigProto) UnsetMappings()`

UnsetMappings ensures that no value is present for Mappings, not even an explicit nil
### GetNetworkEntity

`func (o *RestoredObjectNetworkConfigProto) GetNetworkEntity() EntityProto`

GetNetworkEntity returns the NetworkEntity field if non-nil, zero value otherwise.

### GetNetworkEntityOk

`func (o *RestoredObjectNetworkConfigProto) GetNetworkEntityOk() (*EntityProto, bool)`

GetNetworkEntityOk returns a tuple with the NetworkEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkEntity

`func (o *RestoredObjectNetworkConfigProto) SetNetworkEntity(v EntityProto)`

SetNetworkEntity sets NetworkEntity field to given value.

### HasNetworkEntity

`func (o *RestoredObjectNetworkConfigProto) HasNetworkEntity() bool`

HasNetworkEntity returns a boolean if a field has been set.

### GetPreserveMacAddressOnNewNetwork

`func (o *RestoredObjectNetworkConfigProto) GetPreserveMacAddressOnNewNetwork() bool`

GetPreserveMacAddressOnNewNetwork returns the PreserveMacAddressOnNewNetwork field if non-nil, zero value otherwise.

### GetPreserveMacAddressOnNewNetworkOk

`func (o *RestoredObjectNetworkConfigProto) GetPreserveMacAddressOnNewNetworkOk() (*bool, bool)`

GetPreserveMacAddressOnNewNetworkOk returns a tuple with the PreserveMacAddressOnNewNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveMacAddressOnNewNetwork

`func (o *RestoredObjectNetworkConfigProto) SetPreserveMacAddressOnNewNetwork(v bool)`

SetPreserveMacAddressOnNewNetwork sets PreserveMacAddressOnNewNetwork field to given value.

### HasPreserveMacAddressOnNewNetwork

`func (o *RestoredObjectNetworkConfigProto) HasPreserveMacAddressOnNewNetwork() bool`

HasPreserveMacAddressOnNewNetwork returns a boolean if a field has been set.

### SetPreserveMacAddressOnNewNetworkNil

`func (o *RestoredObjectNetworkConfigProto) SetPreserveMacAddressOnNewNetworkNil(b bool)`

 SetPreserveMacAddressOnNewNetworkNil sets the value for PreserveMacAddressOnNewNetwork to be an explicit nil

### UnsetPreserveMacAddressOnNewNetwork
`func (o *RestoredObjectNetworkConfigProto) UnsetPreserveMacAddressOnNewNetwork()`

UnsetPreserveMacAddressOnNewNetwork ensures that no value is present for PreserveMacAddressOnNewNetwork, not even an explicit nil
### GetVnicEntity

`func (o *RestoredObjectNetworkConfigProto) GetVnicEntity() EntityProto`

GetVnicEntity returns the VnicEntity field if non-nil, zero value otherwise.

### GetVnicEntityOk

`func (o *RestoredObjectNetworkConfigProto) GetVnicEntityOk() (*EntityProto, bool)`

GetVnicEntityOk returns a tuple with the VnicEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnicEntity

`func (o *RestoredObjectNetworkConfigProto) SetVnicEntity(v EntityProto)`

SetVnicEntity sets VnicEntity field to given value.

### HasVnicEntity

`func (o *RestoredObjectNetworkConfigProto) HasVnicEntity() bool`

HasVnicEntity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


