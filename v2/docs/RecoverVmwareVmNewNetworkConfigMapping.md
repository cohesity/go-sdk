# RecoverVmwareVmNewNetworkConfigMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisableNetwork** | Pointer to **NullableBool** | Specifies whether the attached network should be left in disabled state for this mapping. Default is false. | [optional] 
**NetworkAdapterName** | Pointer to **NullableString** | Name of the VM&#39;s network adapter name. | [optional] 
**OrgVdcNetwork** | Pointer to [**OrgVDCNetwork**](OrgVDCNetwork.md) |  | [optional] 
**PreserveMacAddress** | Pointer to **NullableBool** | Specifies whether to preserve the MAC address of the source network entity while attaching to the new target network. Default is false. | [optional] 
**SourceNetworkEntity** | Pointer to [**NullableRecoverVmwareVmNewNetworkConfigMappingSourceNetworkEntity**](RecoverVmwareVmNewNetworkConfigMappingSourceNetworkEntity.md) |  | [optional] 
**TargetNetworkEntity** | Pointer to [**NullableRecoverVmwareVmNewNetworkConfigMappingTargetNetworkEntity**](RecoverVmwareVmNewNetworkConfigMappingTargetNetworkEntity.md) |  | [optional] 

## Methods

### NewRecoverVmwareVmNewNetworkConfigMapping

`func NewRecoverVmwareVmNewNetworkConfigMapping() *RecoverVmwareVmNewNetworkConfigMapping`

NewRecoverVmwareVmNewNetworkConfigMapping instantiates a new RecoverVmwareVmNewNetworkConfigMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverVmwareVmNewNetworkConfigMappingWithDefaults

`func NewRecoverVmwareVmNewNetworkConfigMappingWithDefaults() *RecoverVmwareVmNewNetworkConfigMapping`

NewRecoverVmwareVmNewNetworkConfigMappingWithDefaults instantiates a new RecoverVmwareVmNewNetworkConfigMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisableNetwork

`func (o *RecoverVmwareVmNewNetworkConfigMapping) GetDisableNetwork() bool`

GetDisableNetwork returns the DisableNetwork field if non-nil, zero value otherwise.

### GetDisableNetworkOk

`func (o *RecoverVmwareVmNewNetworkConfigMapping) GetDisableNetworkOk() (*bool, bool)`

GetDisableNetworkOk returns a tuple with the DisableNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableNetwork

`func (o *RecoverVmwareVmNewNetworkConfigMapping) SetDisableNetwork(v bool)`

SetDisableNetwork sets DisableNetwork field to given value.

### HasDisableNetwork

`func (o *RecoverVmwareVmNewNetworkConfigMapping) HasDisableNetwork() bool`

HasDisableNetwork returns a boolean if a field has been set.

### SetDisableNetworkNil

`func (o *RecoverVmwareVmNewNetworkConfigMapping) SetDisableNetworkNil(b bool)`

 SetDisableNetworkNil sets the value for DisableNetwork to be an explicit nil

### UnsetDisableNetwork
`func (o *RecoverVmwareVmNewNetworkConfigMapping) UnsetDisableNetwork()`

UnsetDisableNetwork ensures that no value is present for DisableNetwork, not even an explicit nil
### GetNetworkAdapterName

`func (o *RecoverVmwareVmNewNetworkConfigMapping) GetNetworkAdapterName() string`

GetNetworkAdapterName returns the NetworkAdapterName field if non-nil, zero value otherwise.

### GetNetworkAdapterNameOk

`func (o *RecoverVmwareVmNewNetworkConfigMapping) GetNetworkAdapterNameOk() (*string, bool)`

GetNetworkAdapterNameOk returns a tuple with the NetworkAdapterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkAdapterName

`func (o *RecoverVmwareVmNewNetworkConfigMapping) SetNetworkAdapterName(v string)`

SetNetworkAdapterName sets NetworkAdapterName field to given value.

### HasNetworkAdapterName

`func (o *RecoverVmwareVmNewNetworkConfigMapping) HasNetworkAdapterName() bool`

HasNetworkAdapterName returns a boolean if a field has been set.

### SetNetworkAdapterNameNil

`func (o *RecoverVmwareVmNewNetworkConfigMapping) SetNetworkAdapterNameNil(b bool)`

 SetNetworkAdapterNameNil sets the value for NetworkAdapterName to be an explicit nil

### UnsetNetworkAdapterName
`func (o *RecoverVmwareVmNewNetworkConfigMapping) UnsetNetworkAdapterName()`

UnsetNetworkAdapterName ensures that no value is present for NetworkAdapterName, not even an explicit nil
### GetOrgVdcNetwork

`func (o *RecoverVmwareVmNewNetworkConfigMapping) GetOrgVdcNetwork() OrgVDCNetwork`

GetOrgVdcNetwork returns the OrgVdcNetwork field if non-nil, zero value otherwise.

### GetOrgVdcNetworkOk

`func (o *RecoverVmwareVmNewNetworkConfigMapping) GetOrgVdcNetworkOk() (*OrgVDCNetwork, bool)`

GetOrgVdcNetworkOk returns a tuple with the OrgVdcNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgVdcNetwork

`func (o *RecoverVmwareVmNewNetworkConfigMapping) SetOrgVdcNetwork(v OrgVDCNetwork)`

SetOrgVdcNetwork sets OrgVdcNetwork field to given value.

### HasOrgVdcNetwork

`func (o *RecoverVmwareVmNewNetworkConfigMapping) HasOrgVdcNetwork() bool`

HasOrgVdcNetwork returns a boolean if a field has been set.

### GetPreserveMacAddress

`func (o *RecoverVmwareVmNewNetworkConfigMapping) GetPreserveMacAddress() bool`

GetPreserveMacAddress returns the PreserveMacAddress field if non-nil, zero value otherwise.

### GetPreserveMacAddressOk

`func (o *RecoverVmwareVmNewNetworkConfigMapping) GetPreserveMacAddressOk() (*bool, bool)`

GetPreserveMacAddressOk returns a tuple with the PreserveMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveMacAddress

`func (o *RecoverVmwareVmNewNetworkConfigMapping) SetPreserveMacAddress(v bool)`

SetPreserveMacAddress sets PreserveMacAddress field to given value.

### HasPreserveMacAddress

`func (o *RecoverVmwareVmNewNetworkConfigMapping) HasPreserveMacAddress() bool`

HasPreserveMacAddress returns a boolean if a field has been set.

### SetPreserveMacAddressNil

`func (o *RecoverVmwareVmNewNetworkConfigMapping) SetPreserveMacAddressNil(b bool)`

 SetPreserveMacAddressNil sets the value for PreserveMacAddress to be an explicit nil

### UnsetPreserveMacAddress
`func (o *RecoverVmwareVmNewNetworkConfigMapping) UnsetPreserveMacAddress()`

UnsetPreserveMacAddress ensures that no value is present for PreserveMacAddress, not even an explicit nil
### GetSourceNetworkEntity

`func (o *RecoverVmwareVmNewNetworkConfigMapping) GetSourceNetworkEntity() RecoverVmwareVmNewNetworkConfigMappingSourceNetworkEntity`

GetSourceNetworkEntity returns the SourceNetworkEntity field if non-nil, zero value otherwise.

### GetSourceNetworkEntityOk

`func (o *RecoverVmwareVmNewNetworkConfigMapping) GetSourceNetworkEntityOk() (*RecoverVmwareVmNewNetworkConfigMappingSourceNetworkEntity, bool)`

GetSourceNetworkEntityOk returns a tuple with the SourceNetworkEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceNetworkEntity

`func (o *RecoverVmwareVmNewNetworkConfigMapping) SetSourceNetworkEntity(v RecoverVmwareVmNewNetworkConfigMappingSourceNetworkEntity)`

SetSourceNetworkEntity sets SourceNetworkEntity field to given value.

### HasSourceNetworkEntity

`func (o *RecoverVmwareVmNewNetworkConfigMapping) HasSourceNetworkEntity() bool`

HasSourceNetworkEntity returns a boolean if a field has been set.

### SetSourceNetworkEntityNil

`func (o *RecoverVmwareVmNewNetworkConfigMapping) SetSourceNetworkEntityNil(b bool)`

 SetSourceNetworkEntityNil sets the value for SourceNetworkEntity to be an explicit nil

### UnsetSourceNetworkEntity
`func (o *RecoverVmwareVmNewNetworkConfigMapping) UnsetSourceNetworkEntity()`

UnsetSourceNetworkEntity ensures that no value is present for SourceNetworkEntity, not even an explicit nil
### GetTargetNetworkEntity

`func (o *RecoverVmwareVmNewNetworkConfigMapping) GetTargetNetworkEntity() RecoverVmwareVmNewNetworkConfigMappingTargetNetworkEntity`

GetTargetNetworkEntity returns the TargetNetworkEntity field if non-nil, zero value otherwise.

### GetTargetNetworkEntityOk

`func (o *RecoverVmwareVmNewNetworkConfigMapping) GetTargetNetworkEntityOk() (*RecoverVmwareVmNewNetworkConfigMappingTargetNetworkEntity, bool)`

GetTargetNetworkEntityOk returns a tuple with the TargetNetworkEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetNetworkEntity

`func (o *RecoverVmwareVmNewNetworkConfigMapping) SetTargetNetworkEntity(v RecoverVmwareVmNewNetworkConfigMappingTargetNetworkEntity)`

SetTargetNetworkEntity sets TargetNetworkEntity field to given value.

### HasTargetNetworkEntity

`func (o *RecoverVmwareVmNewNetworkConfigMapping) HasTargetNetworkEntity() bool`

HasTargetNetworkEntity returns a boolean if a field has been set.

### SetTargetNetworkEntityNil

`func (o *RecoverVmwareVmNewNetworkConfigMapping) SetTargetNetworkEntityNil(b bool)`

 SetTargetNetworkEntityNil sets the value for TargetNetworkEntity to be an explicit nil

### UnsetTargetNetworkEntity
`func (o *RecoverVmwareVmNewNetworkConfigMapping) UnsetTargetNetworkEntity()`

UnsetTargetNetworkEntity ensures that no value is present for TargetNetworkEntity, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


