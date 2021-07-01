# NetworkMappingProto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisableNetwork** | Pointer to **NullableBool** | This can be set to true to indicate that the attached network should be left in disabled state. This value takes priority over the value in RestoredObjectNetworkConfigProto. | [optional] 
**PreserveMacAddressOnNewNetwork** | Pointer to **NullableBool** | VM&#39;s MAC address will be preserved on the new network. This value takes priority over the value in RestoredObjectNetworkConfigProto. | [optional] 
**SourceNetworkEntity** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**TargetNetworkEntity** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 

## Methods

### NewNetworkMappingProto

`func NewNetworkMappingProto() *NetworkMappingProto`

NewNetworkMappingProto instantiates a new NetworkMappingProto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkMappingProtoWithDefaults

`func NewNetworkMappingProtoWithDefaults() *NetworkMappingProto`

NewNetworkMappingProtoWithDefaults instantiates a new NetworkMappingProto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisableNetwork

`func (o *NetworkMappingProto) GetDisableNetwork() bool`

GetDisableNetwork returns the DisableNetwork field if non-nil, zero value otherwise.

### GetDisableNetworkOk

`func (o *NetworkMappingProto) GetDisableNetworkOk() (*bool, bool)`

GetDisableNetworkOk returns a tuple with the DisableNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableNetwork

`func (o *NetworkMappingProto) SetDisableNetwork(v bool)`

SetDisableNetwork sets DisableNetwork field to given value.

### HasDisableNetwork

`func (o *NetworkMappingProto) HasDisableNetwork() bool`

HasDisableNetwork returns a boolean if a field has been set.

### SetDisableNetworkNil

`func (o *NetworkMappingProto) SetDisableNetworkNil(b bool)`

 SetDisableNetworkNil sets the value for DisableNetwork to be an explicit nil

### UnsetDisableNetwork
`func (o *NetworkMappingProto) UnsetDisableNetwork()`

UnsetDisableNetwork ensures that no value is present for DisableNetwork, not even an explicit nil
### GetPreserveMacAddressOnNewNetwork

`func (o *NetworkMappingProto) GetPreserveMacAddressOnNewNetwork() bool`

GetPreserveMacAddressOnNewNetwork returns the PreserveMacAddressOnNewNetwork field if non-nil, zero value otherwise.

### GetPreserveMacAddressOnNewNetworkOk

`func (o *NetworkMappingProto) GetPreserveMacAddressOnNewNetworkOk() (*bool, bool)`

GetPreserveMacAddressOnNewNetworkOk returns a tuple with the PreserveMacAddressOnNewNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveMacAddressOnNewNetwork

`func (o *NetworkMappingProto) SetPreserveMacAddressOnNewNetwork(v bool)`

SetPreserveMacAddressOnNewNetwork sets PreserveMacAddressOnNewNetwork field to given value.

### HasPreserveMacAddressOnNewNetwork

`func (o *NetworkMappingProto) HasPreserveMacAddressOnNewNetwork() bool`

HasPreserveMacAddressOnNewNetwork returns a boolean if a field has been set.

### SetPreserveMacAddressOnNewNetworkNil

`func (o *NetworkMappingProto) SetPreserveMacAddressOnNewNetworkNil(b bool)`

 SetPreserveMacAddressOnNewNetworkNil sets the value for PreserveMacAddressOnNewNetwork to be an explicit nil

### UnsetPreserveMacAddressOnNewNetwork
`func (o *NetworkMappingProto) UnsetPreserveMacAddressOnNewNetwork()`

UnsetPreserveMacAddressOnNewNetwork ensures that no value is present for PreserveMacAddressOnNewNetwork, not even an explicit nil
### GetSourceNetworkEntity

`func (o *NetworkMappingProto) GetSourceNetworkEntity() EntityProto`

GetSourceNetworkEntity returns the SourceNetworkEntity field if non-nil, zero value otherwise.

### GetSourceNetworkEntityOk

`func (o *NetworkMappingProto) GetSourceNetworkEntityOk() (*EntityProto, bool)`

GetSourceNetworkEntityOk returns a tuple with the SourceNetworkEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceNetworkEntity

`func (o *NetworkMappingProto) SetSourceNetworkEntity(v EntityProto)`

SetSourceNetworkEntity sets SourceNetworkEntity field to given value.

### HasSourceNetworkEntity

`func (o *NetworkMappingProto) HasSourceNetworkEntity() bool`

HasSourceNetworkEntity returns a boolean if a field has been set.

### GetTargetNetworkEntity

`func (o *NetworkMappingProto) GetTargetNetworkEntity() EntityProto`

GetTargetNetworkEntity returns the TargetNetworkEntity field if non-nil, zero value otherwise.

### GetTargetNetworkEntityOk

`func (o *NetworkMappingProto) GetTargetNetworkEntityOk() (*EntityProto, bool)`

GetTargetNetworkEntityOk returns a tuple with the TargetNetworkEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetNetworkEntity

`func (o *NetworkMappingProto) SetTargetNetworkEntity(v EntityProto)`

SetTargetNetworkEntity sets TargetNetworkEntity field to given value.

### HasTargetNetworkEntity

`func (o *NetworkMappingProto) HasTargetNetworkEntity() bool`

HasTargetNetworkEntity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


