# NetworkMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisableNetwork** | Pointer to **NullableBool** | Specifies if the network should be disabled. On restore or clone of the VM, if the network should be kept in disabled state, set this flag to true. The mapped network is enabled by default. | [optional] 
**PreserveMacAddress** | Pointer to **NullableBool** | Specifies if the source mac address should be preserved after restore or clone. In case of collision of mac address on target network the job won&#39;t fail. Address collision should be resolved manually. | [optional] 
**SourceNetworkId** | Pointer to **NullableInt64** | Specifies the id of the source network. | [optional] 
**TargetNetworkId** | Pointer to **NullableInt64** | Specifies the id of target network. | [optional] 

## Methods

### NewNetworkMapping

`func NewNetworkMapping() *NetworkMapping`

NewNetworkMapping instantiates a new NetworkMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkMappingWithDefaults

`func NewNetworkMappingWithDefaults() *NetworkMapping`

NewNetworkMappingWithDefaults instantiates a new NetworkMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisableNetwork

`func (o *NetworkMapping) GetDisableNetwork() bool`

GetDisableNetwork returns the DisableNetwork field if non-nil, zero value otherwise.

### GetDisableNetworkOk

`func (o *NetworkMapping) GetDisableNetworkOk() (*bool, bool)`

GetDisableNetworkOk returns a tuple with the DisableNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableNetwork

`func (o *NetworkMapping) SetDisableNetwork(v bool)`

SetDisableNetwork sets DisableNetwork field to given value.

### HasDisableNetwork

`func (o *NetworkMapping) HasDisableNetwork() bool`

HasDisableNetwork returns a boolean if a field has been set.

### SetDisableNetworkNil

`func (o *NetworkMapping) SetDisableNetworkNil(b bool)`

 SetDisableNetworkNil sets the value for DisableNetwork to be an explicit nil

### UnsetDisableNetwork
`func (o *NetworkMapping) UnsetDisableNetwork()`

UnsetDisableNetwork ensures that no value is present for DisableNetwork, not even an explicit nil
### GetPreserveMacAddress

`func (o *NetworkMapping) GetPreserveMacAddress() bool`

GetPreserveMacAddress returns the PreserveMacAddress field if non-nil, zero value otherwise.

### GetPreserveMacAddressOk

`func (o *NetworkMapping) GetPreserveMacAddressOk() (*bool, bool)`

GetPreserveMacAddressOk returns a tuple with the PreserveMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveMacAddress

`func (o *NetworkMapping) SetPreserveMacAddress(v bool)`

SetPreserveMacAddress sets PreserveMacAddress field to given value.

### HasPreserveMacAddress

`func (o *NetworkMapping) HasPreserveMacAddress() bool`

HasPreserveMacAddress returns a boolean if a field has been set.

### SetPreserveMacAddressNil

`func (o *NetworkMapping) SetPreserveMacAddressNil(b bool)`

 SetPreserveMacAddressNil sets the value for PreserveMacAddress to be an explicit nil

### UnsetPreserveMacAddress
`func (o *NetworkMapping) UnsetPreserveMacAddress()`

UnsetPreserveMacAddress ensures that no value is present for PreserveMacAddress, not even an explicit nil
### GetSourceNetworkId

`func (o *NetworkMapping) GetSourceNetworkId() int64`

GetSourceNetworkId returns the SourceNetworkId field if non-nil, zero value otherwise.

### GetSourceNetworkIdOk

`func (o *NetworkMapping) GetSourceNetworkIdOk() (*int64, bool)`

GetSourceNetworkIdOk returns a tuple with the SourceNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceNetworkId

`func (o *NetworkMapping) SetSourceNetworkId(v int64)`

SetSourceNetworkId sets SourceNetworkId field to given value.

### HasSourceNetworkId

`func (o *NetworkMapping) HasSourceNetworkId() bool`

HasSourceNetworkId returns a boolean if a field has been set.

### SetSourceNetworkIdNil

`func (o *NetworkMapping) SetSourceNetworkIdNil(b bool)`

 SetSourceNetworkIdNil sets the value for SourceNetworkId to be an explicit nil

### UnsetSourceNetworkId
`func (o *NetworkMapping) UnsetSourceNetworkId()`

UnsetSourceNetworkId ensures that no value is present for SourceNetworkId, not even an explicit nil
### GetTargetNetworkId

`func (o *NetworkMapping) GetTargetNetworkId() int64`

GetTargetNetworkId returns the TargetNetworkId field if non-nil, zero value otherwise.

### GetTargetNetworkIdOk

`func (o *NetworkMapping) GetTargetNetworkIdOk() (*int64, bool)`

GetTargetNetworkIdOk returns a tuple with the TargetNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetNetworkId

`func (o *NetworkMapping) SetTargetNetworkId(v int64)`

SetTargetNetworkId sets TargetNetworkId field to given value.

### HasTargetNetworkId

`func (o *NetworkMapping) HasTargetNetworkId() bool`

HasTargetNetworkId returns a boolean if a field has been set.

### SetTargetNetworkIdNil

`func (o *NetworkMapping) SetTargetNetworkIdNil(b bool)`

 SetTargetNetworkIdNil sets the value for TargetNetworkId to be an explicit nil

### UnsetTargetNetworkId
`func (o *NetworkMapping) UnsetTargetNetworkId()`

UnsetTargetNetworkId ensures that no value is present for TargetNetworkId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


