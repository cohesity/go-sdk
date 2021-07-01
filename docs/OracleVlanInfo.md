# OracleVlanInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gateway** | Pointer to **NullableString** | Gateway for this VLAN. | [optional] 
**Id** | Pointer to **NullableInt32** | ID of this VLAN. | [optional] 
**IpVec** | Pointer to **[]string** | List of IPs in this VLAN. | [optional] 
**SubnetIp** | Pointer to **NullableString** | Subnet IP for this VLAN. | [optional] 

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

### SetIdNil

`func (o *OracleVlanInfo) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *OracleVlanInfo) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIpVec

`func (o *OracleVlanInfo) GetIpVec() []string`

GetIpVec returns the IpVec field if non-nil, zero value otherwise.

### GetIpVecOk

`func (o *OracleVlanInfo) GetIpVecOk() (*[]string, bool)`

GetIpVecOk returns a tuple with the IpVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVec

`func (o *OracleVlanInfo) SetIpVec(v []string)`

SetIpVec sets IpVec field to given value.

### HasIpVec

`func (o *OracleVlanInfo) HasIpVec() bool`

HasIpVec returns a boolean if a field has been set.

### SetIpVecNil

`func (o *OracleVlanInfo) SetIpVecNil(b bool)`

 SetIpVecNil sets the value for IpVec to be an explicit nil

### UnsetIpVec
`func (o *OracleVlanInfo) UnsetIpVec()`

UnsetIpVec ensures that no value is present for IpVec, not even an explicit nil
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


