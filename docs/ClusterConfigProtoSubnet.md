# ClusterConfigProtoSubnet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Component** | Pointer to **NullableInt32** | The component that has claimed this subnet. | [optional] 
**Description** | Pointer to **NullableString** | Description of the subnet. | [optional] 
**Gateway** | Pointer to **NullableString** | Gateway for the subnet. | [optional] 
**Id** | Pointer to **NullableInt32** | ID for this subnet. | [optional] 
**Ip** | Pointer to **NullableString** | ip is subnet IP address given either in v4 or v6. Netmask is specified by giving CIDR length in netmask_bits for ipv6. For IPv4 addresses, netmask_ip4 field is set in dotted decimal. | [optional] 
**NetmaskBits** | Pointer to **NullableInt32** |  | [optional] 
**NetmaskIp4** | Pointer to **NullableString** |  | [optional] 
**NfsAccess** | Pointer to **NullableInt32** | Whether clients from this subnet can mount using NFS protocol. | [optional] 
**NfsAllSquash** | Pointer to **NullableBool** | Whether all clients from this subnet can map view with view_all_squash_uid/view_all_squash_gid configured in the view. | [optional] 
**NfsRootSquash** | Pointer to **NullableBool** | Whether clients from this subnet can mount as root on NFS. | [optional] 
**SmbAccess** | Pointer to **NullableInt32** | Whether clients from this subnet can mount using SMB protocol. | [optional] 

## Methods

### NewClusterConfigProtoSubnet

`func NewClusterConfigProtoSubnet() *ClusterConfigProtoSubnet`

NewClusterConfigProtoSubnet instantiates a new ClusterConfigProtoSubnet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterConfigProtoSubnetWithDefaults

`func NewClusterConfigProtoSubnetWithDefaults() *ClusterConfigProtoSubnet`

NewClusterConfigProtoSubnetWithDefaults instantiates a new ClusterConfigProtoSubnet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponent

`func (o *ClusterConfigProtoSubnet) GetComponent() int32`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *ClusterConfigProtoSubnet) GetComponentOk() (*int32, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *ClusterConfigProtoSubnet) SetComponent(v int32)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *ClusterConfigProtoSubnet) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### SetComponentNil

`func (o *ClusterConfigProtoSubnet) SetComponentNil(b bool)`

 SetComponentNil sets the value for Component to be an explicit nil

### UnsetComponent
`func (o *ClusterConfigProtoSubnet) UnsetComponent()`

UnsetComponent ensures that no value is present for Component, not even an explicit nil
### GetDescription

`func (o *ClusterConfigProtoSubnet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClusterConfigProtoSubnet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClusterConfigProtoSubnet) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClusterConfigProtoSubnet) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ClusterConfigProtoSubnet) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ClusterConfigProtoSubnet) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetGateway

`func (o *ClusterConfigProtoSubnet) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *ClusterConfigProtoSubnet) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *ClusterConfigProtoSubnet) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *ClusterConfigProtoSubnet) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### SetGatewayNil

`func (o *ClusterConfigProtoSubnet) SetGatewayNil(b bool)`

 SetGatewayNil sets the value for Gateway to be an explicit nil

### UnsetGateway
`func (o *ClusterConfigProtoSubnet) UnsetGateway()`

UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
### GetId

`func (o *ClusterConfigProtoSubnet) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterConfigProtoSubnet) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterConfigProtoSubnet) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterConfigProtoSubnet) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ClusterConfigProtoSubnet) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ClusterConfigProtoSubnet) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIp

`func (o *ClusterConfigProtoSubnet) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ClusterConfigProtoSubnet) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ClusterConfigProtoSubnet) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *ClusterConfigProtoSubnet) HasIp() bool`

HasIp returns a boolean if a field has been set.

### SetIpNil

`func (o *ClusterConfigProtoSubnet) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *ClusterConfigProtoSubnet) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
### GetNetmaskBits

`func (o *ClusterConfigProtoSubnet) GetNetmaskBits() int32`

GetNetmaskBits returns the NetmaskBits field if non-nil, zero value otherwise.

### GetNetmaskBitsOk

`func (o *ClusterConfigProtoSubnet) GetNetmaskBitsOk() (*int32, bool)`

GetNetmaskBitsOk returns a tuple with the NetmaskBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmaskBits

`func (o *ClusterConfigProtoSubnet) SetNetmaskBits(v int32)`

SetNetmaskBits sets NetmaskBits field to given value.

### HasNetmaskBits

`func (o *ClusterConfigProtoSubnet) HasNetmaskBits() bool`

HasNetmaskBits returns a boolean if a field has been set.

### SetNetmaskBitsNil

`func (o *ClusterConfigProtoSubnet) SetNetmaskBitsNil(b bool)`

 SetNetmaskBitsNil sets the value for NetmaskBits to be an explicit nil

### UnsetNetmaskBits
`func (o *ClusterConfigProtoSubnet) UnsetNetmaskBits()`

UnsetNetmaskBits ensures that no value is present for NetmaskBits, not even an explicit nil
### GetNetmaskIp4

`func (o *ClusterConfigProtoSubnet) GetNetmaskIp4() string`

GetNetmaskIp4 returns the NetmaskIp4 field if non-nil, zero value otherwise.

### GetNetmaskIp4Ok

`func (o *ClusterConfigProtoSubnet) GetNetmaskIp4Ok() (*string, bool)`

GetNetmaskIp4Ok returns a tuple with the NetmaskIp4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmaskIp4

`func (o *ClusterConfigProtoSubnet) SetNetmaskIp4(v string)`

SetNetmaskIp4 sets NetmaskIp4 field to given value.

### HasNetmaskIp4

`func (o *ClusterConfigProtoSubnet) HasNetmaskIp4() bool`

HasNetmaskIp4 returns a boolean if a field has been set.

### SetNetmaskIp4Nil

`func (o *ClusterConfigProtoSubnet) SetNetmaskIp4Nil(b bool)`

 SetNetmaskIp4Nil sets the value for NetmaskIp4 to be an explicit nil

### UnsetNetmaskIp4
`func (o *ClusterConfigProtoSubnet) UnsetNetmaskIp4()`

UnsetNetmaskIp4 ensures that no value is present for NetmaskIp4, not even an explicit nil
### GetNfsAccess

`func (o *ClusterConfigProtoSubnet) GetNfsAccess() int32`

GetNfsAccess returns the NfsAccess field if non-nil, zero value otherwise.

### GetNfsAccessOk

`func (o *ClusterConfigProtoSubnet) GetNfsAccessOk() (*int32, bool)`

GetNfsAccessOk returns a tuple with the NfsAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsAccess

`func (o *ClusterConfigProtoSubnet) SetNfsAccess(v int32)`

SetNfsAccess sets NfsAccess field to given value.

### HasNfsAccess

`func (o *ClusterConfigProtoSubnet) HasNfsAccess() bool`

HasNfsAccess returns a boolean if a field has been set.

### SetNfsAccessNil

`func (o *ClusterConfigProtoSubnet) SetNfsAccessNil(b bool)`

 SetNfsAccessNil sets the value for NfsAccess to be an explicit nil

### UnsetNfsAccess
`func (o *ClusterConfigProtoSubnet) UnsetNfsAccess()`

UnsetNfsAccess ensures that no value is present for NfsAccess, not even an explicit nil
### GetNfsAllSquash

`func (o *ClusterConfigProtoSubnet) GetNfsAllSquash() bool`

GetNfsAllSquash returns the NfsAllSquash field if non-nil, zero value otherwise.

### GetNfsAllSquashOk

`func (o *ClusterConfigProtoSubnet) GetNfsAllSquashOk() (*bool, bool)`

GetNfsAllSquashOk returns a tuple with the NfsAllSquash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsAllSquash

`func (o *ClusterConfigProtoSubnet) SetNfsAllSquash(v bool)`

SetNfsAllSquash sets NfsAllSquash field to given value.

### HasNfsAllSquash

`func (o *ClusterConfigProtoSubnet) HasNfsAllSquash() bool`

HasNfsAllSquash returns a boolean if a field has been set.

### SetNfsAllSquashNil

`func (o *ClusterConfigProtoSubnet) SetNfsAllSquashNil(b bool)`

 SetNfsAllSquashNil sets the value for NfsAllSquash to be an explicit nil

### UnsetNfsAllSquash
`func (o *ClusterConfigProtoSubnet) UnsetNfsAllSquash()`

UnsetNfsAllSquash ensures that no value is present for NfsAllSquash, not even an explicit nil
### GetNfsRootSquash

`func (o *ClusterConfigProtoSubnet) GetNfsRootSquash() bool`

GetNfsRootSquash returns the NfsRootSquash field if non-nil, zero value otherwise.

### GetNfsRootSquashOk

`func (o *ClusterConfigProtoSubnet) GetNfsRootSquashOk() (*bool, bool)`

GetNfsRootSquashOk returns a tuple with the NfsRootSquash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsRootSquash

`func (o *ClusterConfigProtoSubnet) SetNfsRootSquash(v bool)`

SetNfsRootSquash sets NfsRootSquash field to given value.

### HasNfsRootSquash

`func (o *ClusterConfigProtoSubnet) HasNfsRootSquash() bool`

HasNfsRootSquash returns a boolean if a field has been set.

### SetNfsRootSquashNil

`func (o *ClusterConfigProtoSubnet) SetNfsRootSquashNil(b bool)`

 SetNfsRootSquashNil sets the value for NfsRootSquash to be an explicit nil

### UnsetNfsRootSquash
`func (o *ClusterConfigProtoSubnet) UnsetNfsRootSquash()`

UnsetNfsRootSquash ensures that no value is present for NfsRootSquash, not even an explicit nil
### GetSmbAccess

`func (o *ClusterConfigProtoSubnet) GetSmbAccess() int32`

GetSmbAccess returns the SmbAccess field if non-nil, zero value otherwise.

### GetSmbAccessOk

`func (o *ClusterConfigProtoSubnet) GetSmbAccessOk() (*int32, bool)`

GetSmbAccessOk returns a tuple with the SmbAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbAccess

`func (o *ClusterConfigProtoSubnet) SetSmbAccess(v int32)`

SetSmbAccess sets SmbAccess field to given value.

### HasSmbAccess

`func (o *ClusterConfigProtoSubnet) HasSmbAccess() bool`

HasSmbAccess returns a boolean if a field has been set.

### SetSmbAccessNil

`func (o *ClusterConfigProtoSubnet) SetSmbAccessNil(b bool)`

 SetSmbAccessNil sets the value for SmbAccess to be an explicit nil

### UnsetSmbAccess
`func (o *ClusterConfigProtoSubnet) UnsetSmbAccess()`

UnsetSmbAccess ensures that no value is present for SmbAccess, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


