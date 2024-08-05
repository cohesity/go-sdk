# HyperVTargetParamsForRecoverVmVlanConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisableVlan** | Pointer to **NullableBool** | If this is set to true, then even if VLANs are configured on the system, the partition VIPs will be used for the Recovery. | [optional] 
**Id** | Pointer to **NullableInt32** | If this is set, then the Cohesity host name or the IP address associated with this vlan is used for mounting Cohesity&#39;s view on the remote host. | [optional] 
**InterfaceName** | Pointer to **NullableString** | Interface group to use for Recovery. | [optional] [readonly] 

## Methods

### NewHyperVTargetParamsForRecoverVmVlanConfig

`func NewHyperVTargetParamsForRecoverVmVlanConfig() *HyperVTargetParamsForRecoverVmVlanConfig`

NewHyperVTargetParamsForRecoverVmVlanConfig instantiates a new HyperVTargetParamsForRecoverVmVlanConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperVTargetParamsForRecoverVmVlanConfigWithDefaults

`func NewHyperVTargetParamsForRecoverVmVlanConfigWithDefaults() *HyperVTargetParamsForRecoverVmVlanConfig`

NewHyperVTargetParamsForRecoverVmVlanConfigWithDefaults instantiates a new HyperVTargetParamsForRecoverVmVlanConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisableVlan

`func (o *HyperVTargetParamsForRecoverVmVlanConfig) GetDisableVlan() bool`

GetDisableVlan returns the DisableVlan field if non-nil, zero value otherwise.

### GetDisableVlanOk

`func (o *HyperVTargetParamsForRecoverVmVlanConfig) GetDisableVlanOk() (*bool, bool)`

GetDisableVlanOk returns a tuple with the DisableVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableVlan

`func (o *HyperVTargetParamsForRecoverVmVlanConfig) SetDisableVlan(v bool)`

SetDisableVlan sets DisableVlan field to given value.

### HasDisableVlan

`func (o *HyperVTargetParamsForRecoverVmVlanConfig) HasDisableVlan() bool`

HasDisableVlan returns a boolean if a field has been set.

### SetDisableVlanNil

`func (o *HyperVTargetParamsForRecoverVmVlanConfig) SetDisableVlanNil(b bool)`

 SetDisableVlanNil sets the value for DisableVlan to be an explicit nil

### UnsetDisableVlan
`func (o *HyperVTargetParamsForRecoverVmVlanConfig) UnsetDisableVlan()`

UnsetDisableVlan ensures that no value is present for DisableVlan, not even an explicit nil
### GetId

`func (o *HyperVTargetParamsForRecoverVmVlanConfig) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HyperVTargetParamsForRecoverVmVlanConfig) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HyperVTargetParamsForRecoverVmVlanConfig) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *HyperVTargetParamsForRecoverVmVlanConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *HyperVTargetParamsForRecoverVmVlanConfig) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *HyperVTargetParamsForRecoverVmVlanConfig) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetInterfaceName

`func (o *HyperVTargetParamsForRecoverVmVlanConfig) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *HyperVTargetParamsForRecoverVmVlanConfig) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *HyperVTargetParamsForRecoverVmVlanConfig) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *HyperVTargetParamsForRecoverVmVlanConfig) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### SetInterfaceNameNil

`func (o *HyperVTargetParamsForRecoverVmVlanConfig) SetInterfaceNameNil(b bool)`

 SetInterfaceNameNil sets the value for InterfaceName to be an explicit nil

### UnsetInterfaceName
`func (o *HyperVTargetParamsForRecoverVmVlanConfig) UnsetInterfaceName()`

UnsetInterfaceName ensures that no value is present for InterfaceName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


