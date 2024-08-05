# RecoverVmwareVmNewNetworkConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisableNetwork** | Pointer to **NullableBool** | Specifies whether the attached network should be left in disabled state. Default is false | [optional] 
**Mappings** | Pointer to [**[]RecoverVmwareVmNewNetworkConfigMapping**](RecoverVmwareVmNewNetworkConfigMapping.md) | Specifies the target network mapping for each VM&#39;s network adapter. | [optional] 
**NetworkPortGroup** | Pointer to [**NullableRecoverVmwareVmNewNetworkConfigNetworkPortGroup**](RecoverVmwareVmNewNetworkConfigNetworkPortGroup.md) |  | [optional] 
**PreserveMacAddress** | Pointer to **NullableBool** | If this is true and we are attaching to a new network entity, then the VM&#39;s MAC address will be preserved on the new network. Default value is false. | [optional] 

## Methods

### NewRecoverVmwareVmNewNetworkConfig

`func NewRecoverVmwareVmNewNetworkConfig() *RecoverVmwareVmNewNetworkConfig`

NewRecoverVmwareVmNewNetworkConfig instantiates a new RecoverVmwareVmNewNetworkConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverVmwareVmNewNetworkConfigWithDefaults

`func NewRecoverVmwareVmNewNetworkConfigWithDefaults() *RecoverVmwareVmNewNetworkConfig`

NewRecoverVmwareVmNewNetworkConfigWithDefaults instantiates a new RecoverVmwareVmNewNetworkConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisableNetwork

`func (o *RecoverVmwareVmNewNetworkConfig) GetDisableNetwork() bool`

GetDisableNetwork returns the DisableNetwork field if non-nil, zero value otherwise.

### GetDisableNetworkOk

`func (o *RecoverVmwareVmNewNetworkConfig) GetDisableNetworkOk() (*bool, bool)`

GetDisableNetworkOk returns a tuple with the DisableNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableNetwork

`func (o *RecoverVmwareVmNewNetworkConfig) SetDisableNetwork(v bool)`

SetDisableNetwork sets DisableNetwork field to given value.

### HasDisableNetwork

`func (o *RecoverVmwareVmNewNetworkConfig) HasDisableNetwork() bool`

HasDisableNetwork returns a boolean if a field has been set.

### SetDisableNetworkNil

`func (o *RecoverVmwareVmNewNetworkConfig) SetDisableNetworkNil(b bool)`

 SetDisableNetworkNil sets the value for DisableNetwork to be an explicit nil

### UnsetDisableNetwork
`func (o *RecoverVmwareVmNewNetworkConfig) UnsetDisableNetwork()`

UnsetDisableNetwork ensures that no value is present for DisableNetwork, not even an explicit nil
### GetMappings

`func (o *RecoverVmwareVmNewNetworkConfig) GetMappings() []RecoverVmwareVmNewNetworkConfigMapping`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *RecoverVmwareVmNewNetworkConfig) GetMappingsOk() (*[]RecoverVmwareVmNewNetworkConfigMapping, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *RecoverVmwareVmNewNetworkConfig) SetMappings(v []RecoverVmwareVmNewNetworkConfigMapping)`

SetMappings sets Mappings field to given value.

### HasMappings

`func (o *RecoverVmwareVmNewNetworkConfig) HasMappings() bool`

HasMappings returns a boolean if a field has been set.

### SetMappingsNil

`func (o *RecoverVmwareVmNewNetworkConfig) SetMappingsNil(b bool)`

 SetMappingsNil sets the value for Mappings to be an explicit nil

### UnsetMappings
`func (o *RecoverVmwareVmNewNetworkConfig) UnsetMappings()`

UnsetMappings ensures that no value is present for Mappings, not even an explicit nil
### GetNetworkPortGroup

`func (o *RecoverVmwareVmNewNetworkConfig) GetNetworkPortGroup() RecoverVmwareVmNewNetworkConfigNetworkPortGroup`

GetNetworkPortGroup returns the NetworkPortGroup field if non-nil, zero value otherwise.

### GetNetworkPortGroupOk

`func (o *RecoverVmwareVmNewNetworkConfig) GetNetworkPortGroupOk() (*RecoverVmwareVmNewNetworkConfigNetworkPortGroup, bool)`

GetNetworkPortGroupOk returns a tuple with the NetworkPortGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPortGroup

`func (o *RecoverVmwareVmNewNetworkConfig) SetNetworkPortGroup(v RecoverVmwareVmNewNetworkConfigNetworkPortGroup)`

SetNetworkPortGroup sets NetworkPortGroup field to given value.

### HasNetworkPortGroup

`func (o *RecoverVmwareVmNewNetworkConfig) HasNetworkPortGroup() bool`

HasNetworkPortGroup returns a boolean if a field has been set.

### SetNetworkPortGroupNil

`func (o *RecoverVmwareVmNewNetworkConfig) SetNetworkPortGroupNil(b bool)`

 SetNetworkPortGroupNil sets the value for NetworkPortGroup to be an explicit nil

### UnsetNetworkPortGroup
`func (o *RecoverVmwareVmNewNetworkConfig) UnsetNetworkPortGroup()`

UnsetNetworkPortGroup ensures that no value is present for NetworkPortGroup, not even an explicit nil
### GetPreserveMacAddress

`func (o *RecoverVmwareVmNewNetworkConfig) GetPreserveMacAddress() bool`

GetPreserveMacAddress returns the PreserveMacAddress field if non-nil, zero value otherwise.

### GetPreserveMacAddressOk

`func (o *RecoverVmwareVmNewNetworkConfig) GetPreserveMacAddressOk() (*bool, bool)`

GetPreserveMacAddressOk returns a tuple with the PreserveMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveMacAddress

`func (o *RecoverVmwareVmNewNetworkConfig) SetPreserveMacAddress(v bool)`

SetPreserveMacAddress sets PreserveMacAddress field to given value.

### HasPreserveMacAddress

`func (o *RecoverVmwareVmNewNetworkConfig) HasPreserveMacAddress() bool`

HasPreserveMacAddress returns a boolean if a field has been set.

### SetPreserveMacAddressNil

`func (o *RecoverVmwareVmNewNetworkConfig) SetPreserveMacAddressNil(b bool)`

 SetPreserveMacAddressNil sets the value for PreserveMacAddress to be an explicit nil

### UnsetPreserveMacAddress
`func (o *RecoverVmwareVmNewNetworkConfig) UnsetPreserveMacAddress()`

UnsetPreserveMacAddress ensures that no value is present for PreserveMacAddress, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


