# VlanParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisableVlan** | Pointer to **NullableBool** | Specifies whether to use the VIPs even when VLANs are configured on the Cluster. If configured, VLAN IP addresses are used by default. If VLANs are not configured, this flag is ignored. Set this flag to true to force using the partition VIPs when VLANs are configured on the Cluster. | [optional] 
**InterfaceName** | Pointer to **NullableString** | Specifies the physical interface group name to use for mounting Cohesity&#39;s view on the remote host. If specified, Cohesity hostname or the IP address on this VLAN is used. | [optional] 
**Vlan** | Pointer to **NullableInt32** | Specifies the VLAN to use for mounting Cohesity&#39;s view on the remote host. If specified, Cohesity hostname or the IP address on this VLAN is used. | [optional] 

## Methods

### NewVlanParameters

`func NewVlanParameters() *VlanParameters`

NewVlanParameters instantiates a new VlanParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVlanParametersWithDefaults

`func NewVlanParametersWithDefaults() *VlanParameters`

NewVlanParametersWithDefaults instantiates a new VlanParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisableVlan

`func (o *VlanParameters) GetDisableVlan() bool`

GetDisableVlan returns the DisableVlan field if non-nil, zero value otherwise.

### GetDisableVlanOk

`func (o *VlanParameters) GetDisableVlanOk() (*bool, bool)`

GetDisableVlanOk returns a tuple with the DisableVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableVlan

`func (o *VlanParameters) SetDisableVlan(v bool)`

SetDisableVlan sets DisableVlan field to given value.

### HasDisableVlan

`func (o *VlanParameters) HasDisableVlan() bool`

HasDisableVlan returns a boolean if a field has been set.

### SetDisableVlanNil

`func (o *VlanParameters) SetDisableVlanNil(b bool)`

 SetDisableVlanNil sets the value for DisableVlan to be an explicit nil

### UnsetDisableVlan
`func (o *VlanParameters) UnsetDisableVlan()`

UnsetDisableVlan ensures that no value is present for DisableVlan, not even an explicit nil
### GetInterfaceName

`func (o *VlanParameters) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *VlanParameters) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *VlanParameters) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *VlanParameters) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### SetInterfaceNameNil

`func (o *VlanParameters) SetInterfaceNameNil(b bool)`

 SetInterfaceNameNil sets the value for InterfaceName to be an explicit nil

### UnsetInterfaceName
`func (o *VlanParameters) UnsetInterfaceName()`

UnsetInterfaceName ensures that no value is present for InterfaceName, not even an explicit nil
### GetVlan

`func (o *VlanParameters) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *VlanParameters) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *VlanParameters) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *VlanParameters) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### SetVlanNil

`func (o *VlanParameters) SetVlanNil(b bool)`

 SetVlanNil sets the value for Vlan to be an explicit nil

### UnsetVlan
`func (o *VlanParameters) UnsetVlan()`

UnsetVlan ensures that no value is present for Vlan, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


