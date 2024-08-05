# KvmTargetParamsForRecoverVmVlanConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisableVlan** | Pointer to **NullableBool** | If this is set to true, then even if VLANs are configured on the system, the partition VIPs will be used for the Recovery. | [optional] 
**Id** | Pointer to **NullableInt32** | If this is set, then the Cohesity host name or the IP address associated with this vlan is used for mounting Cohesity&#39;s view on the remote host. | [optional] 
**InterfaceName** | Pointer to **NullableString** | Interface group to use for Recovery. | [optional] [readonly] 

## Methods

### NewKvmTargetParamsForRecoverVmVlanConfig

`func NewKvmTargetParamsForRecoverVmVlanConfig() *KvmTargetParamsForRecoverVmVlanConfig`

NewKvmTargetParamsForRecoverVmVlanConfig instantiates a new KvmTargetParamsForRecoverVmVlanConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvmTargetParamsForRecoverVmVlanConfigWithDefaults

`func NewKvmTargetParamsForRecoverVmVlanConfigWithDefaults() *KvmTargetParamsForRecoverVmVlanConfig`

NewKvmTargetParamsForRecoverVmVlanConfigWithDefaults instantiates a new KvmTargetParamsForRecoverVmVlanConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisableVlan

`func (o *KvmTargetParamsForRecoverVmVlanConfig) GetDisableVlan() bool`

GetDisableVlan returns the DisableVlan field if non-nil, zero value otherwise.

### GetDisableVlanOk

`func (o *KvmTargetParamsForRecoverVmVlanConfig) GetDisableVlanOk() (*bool, bool)`

GetDisableVlanOk returns a tuple with the DisableVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableVlan

`func (o *KvmTargetParamsForRecoverVmVlanConfig) SetDisableVlan(v bool)`

SetDisableVlan sets DisableVlan field to given value.

### HasDisableVlan

`func (o *KvmTargetParamsForRecoverVmVlanConfig) HasDisableVlan() bool`

HasDisableVlan returns a boolean if a field has been set.

### SetDisableVlanNil

`func (o *KvmTargetParamsForRecoverVmVlanConfig) SetDisableVlanNil(b bool)`

 SetDisableVlanNil sets the value for DisableVlan to be an explicit nil

### UnsetDisableVlan
`func (o *KvmTargetParamsForRecoverVmVlanConfig) UnsetDisableVlan()`

UnsetDisableVlan ensures that no value is present for DisableVlan, not even an explicit nil
### GetId

`func (o *KvmTargetParamsForRecoverVmVlanConfig) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KvmTargetParamsForRecoverVmVlanConfig) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KvmTargetParamsForRecoverVmVlanConfig) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KvmTargetParamsForRecoverVmVlanConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *KvmTargetParamsForRecoverVmVlanConfig) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *KvmTargetParamsForRecoverVmVlanConfig) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetInterfaceName

`func (o *KvmTargetParamsForRecoverVmVlanConfig) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *KvmTargetParamsForRecoverVmVlanConfig) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *KvmTargetParamsForRecoverVmVlanConfig) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *KvmTargetParamsForRecoverVmVlanConfig) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### SetInterfaceNameNil

`func (o *KvmTargetParamsForRecoverVmVlanConfig) SetInterfaceNameNil(b bool)`

 SetInterfaceNameNil sets the value for InterfaceName to be an explicit nil

### UnsetInterfaceName
`func (o *KvmTargetParamsForRecoverVmVlanConfig) UnsetInterfaceName()`

UnsetInterfaceName ensures that no value is present for InterfaceName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


