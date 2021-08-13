# SubnetInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubnetIp** | Pointer to **NullableString** | Subnet IP. | [optional] 
**NetmaskBits** | Pointer to **NullableInt32** | Subnet netmask bits. | [optional] 
**Gateway** | Pointer to **NullableString** | Gateway. | [optional] 

## Methods

### NewSubnetInfo

`func NewSubnetInfo() *SubnetInfo`

NewSubnetInfo instantiates a new SubnetInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubnetInfoWithDefaults

`func NewSubnetInfoWithDefaults() *SubnetInfo`

NewSubnetInfoWithDefaults instantiates a new SubnetInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubnetIp

`func (o *SubnetInfo) GetSubnetIp() string`

GetSubnetIp returns the SubnetIp field if non-nil, zero value otherwise.

### GetSubnetIpOk

`func (o *SubnetInfo) GetSubnetIpOk() (*string, bool)`

GetSubnetIpOk returns a tuple with the SubnetIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetIp

`func (o *SubnetInfo) SetSubnetIp(v string)`

SetSubnetIp sets SubnetIp field to given value.

### HasSubnetIp

`func (o *SubnetInfo) HasSubnetIp() bool`

HasSubnetIp returns a boolean if a field has been set.

### SetSubnetIpNil

`func (o *SubnetInfo) SetSubnetIpNil(b bool)`

 SetSubnetIpNil sets the value for SubnetIp to be an explicit nil

### UnsetSubnetIp
`func (o *SubnetInfo) UnsetSubnetIp()`

UnsetSubnetIp ensures that no value is present for SubnetIp, not even an explicit nil
### GetNetmaskBits

`func (o *SubnetInfo) GetNetmaskBits() int32`

GetNetmaskBits returns the NetmaskBits field if non-nil, zero value otherwise.

### GetNetmaskBitsOk

`func (o *SubnetInfo) GetNetmaskBitsOk() (*int32, bool)`

GetNetmaskBitsOk returns a tuple with the NetmaskBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmaskBits

`func (o *SubnetInfo) SetNetmaskBits(v int32)`

SetNetmaskBits sets NetmaskBits field to given value.

### HasNetmaskBits

`func (o *SubnetInfo) HasNetmaskBits() bool`

HasNetmaskBits returns a boolean if a field has been set.

### SetNetmaskBitsNil

`func (o *SubnetInfo) SetNetmaskBitsNil(b bool)`

 SetNetmaskBitsNil sets the value for NetmaskBits to be an explicit nil

### UnsetNetmaskBits
`func (o *SubnetInfo) UnsetNetmaskBits()`

UnsetNetmaskBits ensures that no value is present for NetmaskBits, not even an explicit nil
### GetGateway

`func (o *SubnetInfo) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *SubnetInfo) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *SubnetInfo) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *SubnetInfo) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### SetGatewayNil

`func (o *SubnetInfo) SetGatewayNil(b bool)`

 SetGatewayNil sets the value for Gateway to be an explicit nil

### UnsetGateway
`func (o *SubnetInfo) UnsetGateway()`

UnsetGateway ensures that no value is present for Gateway, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


