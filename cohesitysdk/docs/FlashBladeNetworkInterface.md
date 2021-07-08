# FlashBladeNetworkInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddress** | Pointer to **NullableString** | Specifies the IP address of the Pure Storage FlashBlade Array. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the network interface. | [optional] 
**Vlan** | Pointer to **NullableInt32** | Specifies the id of the VLAN network of the Pure Storage FlashBlade Array. | [optional] 

## Methods

### NewFlashBladeNetworkInterface

`func NewFlashBladeNetworkInterface() *FlashBladeNetworkInterface`

NewFlashBladeNetworkInterface instantiates a new FlashBladeNetworkInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlashBladeNetworkInterfaceWithDefaults

`func NewFlashBladeNetworkInterfaceWithDefaults() *FlashBladeNetworkInterface`

NewFlashBladeNetworkInterfaceWithDefaults instantiates a new FlashBladeNetworkInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAddress

`func (o *FlashBladeNetworkInterface) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *FlashBladeNetworkInterface) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *FlashBladeNetworkInterface) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *FlashBladeNetworkInterface) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *FlashBladeNetworkInterface) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *FlashBladeNetworkInterface) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetName

`func (o *FlashBladeNetworkInterface) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlashBladeNetworkInterface) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlashBladeNetworkInterface) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FlashBladeNetworkInterface) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *FlashBladeNetworkInterface) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *FlashBladeNetworkInterface) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetVlan

`func (o *FlashBladeNetworkInterface) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *FlashBladeNetworkInterface) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *FlashBladeNetworkInterface) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *FlashBladeNetworkInterface) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### SetVlanNil

`func (o *FlashBladeNetworkInterface) SetVlanNil(b bool)`

 SetVlanNil sets the value for Vlan to be an explicit nil

### UnsetVlan
`func (o *FlashBladeNetworkInterface) UnsetVlan()`

UnsetVlan ensures that no value is present for Vlan, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


