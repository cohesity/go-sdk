# ConnectionSubnet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **NullableString** | Specifies the IP address part of the CIDR notation. | [optional] 
**MaskBits** | Pointer to **NullableInt32** | Specifies the number of set bits in the subnet mask. | [optional] 

## Methods

### NewConnectionSubnet

`func NewConnectionSubnet() *ConnectionSubnet`

NewConnectionSubnet instantiates a new ConnectionSubnet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionSubnetWithDefaults

`func NewConnectionSubnetWithDefaults() *ConnectionSubnet`

NewConnectionSubnetWithDefaults instantiates a new ConnectionSubnet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *ConnectionSubnet) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ConnectionSubnet) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ConnectionSubnet) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *ConnectionSubnet) HasIp() bool`

HasIp returns a boolean if a field has been set.

### SetIpNil

`func (o *ConnectionSubnet) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *ConnectionSubnet) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
### GetMaskBits

`func (o *ConnectionSubnet) GetMaskBits() int32`

GetMaskBits returns the MaskBits field if non-nil, zero value otherwise.

### GetMaskBitsOk

`func (o *ConnectionSubnet) GetMaskBitsOk() (*int32, bool)`

GetMaskBitsOk returns a tuple with the MaskBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskBits

`func (o *ConnectionSubnet) SetMaskBits(v int32)`

SetMaskBits sets MaskBits field to given value.

### HasMaskBits

`func (o *ConnectionSubnet) HasMaskBits() bool`

HasMaskBits returns a boolean if a field has been set.

### SetMaskBitsNil

`func (o *ConnectionSubnet) SetMaskBitsNil(b bool)`

 SetMaskBitsNil sets the value for MaskBits to be an explicit nil

### UnsetMaskBits
`func (o *ConnectionSubnet) UnsetMaskBits()`

UnsetMaskBits ensures that no value is present for MaskBits, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


