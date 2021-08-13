# RecoverKvmVmNewNetworkConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkPortGroup** | Pointer to [**RecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the network port group (i.e, either a standard switch port group or a distributed port group) that will attached to the recovered Object. This parameter is mandatory if detach network is specified as false. | [optional] 
**VnicProfile** | Pointer to [**RecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies VNic profile that will be attached to the restored object. | [optional] 

## Methods

### NewRecoverKvmVmNewNetworkConfig

`func NewRecoverKvmVmNewNetworkConfig() *RecoverKvmVmNewNetworkConfig`

NewRecoverKvmVmNewNetworkConfig instantiates a new RecoverKvmVmNewNetworkConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverKvmVmNewNetworkConfigWithDefaults

`func NewRecoverKvmVmNewNetworkConfigWithDefaults() *RecoverKvmVmNewNetworkConfig`

NewRecoverKvmVmNewNetworkConfigWithDefaults instantiates a new RecoverKvmVmNewNetworkConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkPortGroup

`func (o *RecoverKvmVmNewNetworkConfig) GetNetworkPortGroup() RecoveryObjectIdentifier`

GetNetworkPortGroup returns the NetworkPortGroup field if non-nil, zero value otherwise.

### GetNetworkPortGroupOk

`func (o *RecoverKvmVmNewNetworkConfig) GetNetworkPortGroupOk() (*RecoveryObjectIdentifier, bool)`

GetNetworkPortGroupOk returns a tuple with the NetworkPortGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPortGroup

`func (o *RecoverKvmVmNewNetworkConfig) SetNetworkPortGroup(v RecoveryObjectIdentifier)`

SetNetworkPortGroup sets NetworkPortGroup field to given value.

### HasNetworkPortGroup

`func (o *RecoverKvmVmNewNetworkConfig) HasNetworkPortGroup() bool`

HasNetworkPortGroup returns a boolean if a field has been set.

### GetVnicProfile

`func (o *RecoverKvmVmNewNetworkConfig) GetVnicProfile() RecoveryObjectIdentifier`

GetVnicProfile returns the VnicProfile field if non-nil, zero value otherwise.

### GetVnicProfileOk

`func (o *RecoverKvmVmNewNetworkConfig) GetVnicProfileOk() (*RecoveryObjectIdentifier, bool)`

GetVnicProfileOk returns a tuple with the VnicProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnicProfile

`func (o *RecoverKvmVmNewNetworkConfig) SetVnicProfile(v RecoveryObjectIdentifier)`

SetVnicProfile sets VnicProfile field to given value.

### HasVnicProfile

`func (o *RecoverKvmVmNewNetworkConfig) HasVnicProfile() bool`

HasVnicProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


