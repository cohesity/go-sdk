# OracleVlanInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpList** | Pointer to **[]string** | Specifies the list of Ips in this VLAN. | [optional] 
**Gateway** | Pointer to **NullableString** | Specifies the gateway of this VLAN. | [optional] 
**Id** | Pointer to **int32** | Specifies the Id of this VLAN. | [optional] 
**SubnetIp** | Pointer to **NullableString** | Specifies the subnet Ip for this VLAN. | [optional] 

## Methods

### NewOracleVlanInfo

`func NewOracleVlanInfo() *OracleVlanInfo`

NewOracleVlanInfo instantiates a new OracleVlanInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOracleVlanInfoWithDefaults

`func NewOracleVlanInfoWithDefaults() *OracleVlanInfo`

NewOracleVlanInfoWithDefaults instantiates a new OracleVlanInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpList

`func (o *OracleVlanInfo) GetIpList() []string`

GetIpList returns the IpList field if non-nil, zero value otherwise.

### GetIpListOk

`func (o *OracleVlanInfo) GetIpListOk() (*[]string, bool)`

GetIpListOk returns a tuple with the IpList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpList

`func (o *OracleVlanInfo) SetIpList(v []string)`

SetIpList sets IpList field to given value.

### HasIpList

`func (o *OracleVlanInfo) HasIpList() bool`

HasIpList returns a boolean if a field has been set.

### SetIpListNil

`func (o *OracleVlanInfo) SetIpListNil(b bool)`

 SetIpListNil sets the value for IpList to be an explicit nil

### UnsetIpList
`func (o *OracleVlanInfo) UnsetIpList()`

UnsetIpList ensures that no value is present for IpList, not even an explicit nil
### GetGateway

`func (o *OracleVlanInfo) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *OracleVlanInfo) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *OracleVlanInfo) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *OracleVlanInfo) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### SetGatewayNil

`func (o *OracleVlanInfo) SetGatewayNil(b bool)`

 SetGatewayNil sets the value for Gateway to be an explicit nil

### UnsetGateway
`func (o *OracleVlanInfo) UnsetGateway()`

UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
### GetId

`func (o *OracleVlanInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OracleVlanInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OracleVlanInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OracleVlanInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSubnetIp

`func (o *OracleVlanInfo) GetSubnetIp() string`

GetSubnetIp returns the SubnetIp field if non-nil, zero value otherwise.

### GetSubnetIpOk

`func (o *OracleVlanInfo) GetSubnetIpOk() (*string, bool)`

GetSubnetIpOk returns a tuple with the SubnetIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetIp

`func (o *OracleVlanInfo) SetSubnetIp(v string)`

SetSubnetIp sets SubnetIp field to given value.

### HasSubnetIp

`func (o *OracleVlanInfo) HasSubnetIp() bool`

HasSubnetIp returns a boolean if a field has been set.

### SetSubnetIpNil

`func (o *OracleVlanInfo) SetSubnetIpNil(b bool)`

 SetSubnetIpNil sets the value for SubnetIp to be an explicit nil

### UnsetSubnetIp
`func (o *OracleVlanInfo) UnsetSubnetIp()`

UnsetSubnetIp ensures that no value is present for SubnetIp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


