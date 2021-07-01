# Subnet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Component** | Pointer to **NullableString** | Component that has reserved the subnet. | [optional] 
**Description** | Pointer to **NullableString** | Description of the subnet. | [optional] 
**Id** | Pointer to **NullableInt32** | ID of the subnet. | [optional] 
**Ip** | Pointer to **NullableString** | Specifies either an IPv6 address or an IPv4 address. | [optional] 
**NetmaskBits** | Pointer to **NullableInt32** | Specifies the netmask using bits. | [optional] 
**NetmaskIp4** | Pointer to **NullableString** | Specifies the netmask using an IP4 address. The netmask can only be set using netmaskIp4 if the IP address is an IPv4 address. | [optional] 
**NfsAccess** | Pointer to **NullableString** | Specifies whether clients from this subnet can mount using NFS protocol. Protocol access level. &#39;kDisabled&#39; indicates Protocol access level &#39;Disabled&#39; &#39;kReadOnly&#39; indicates Protocol access level &#39;ReadOnly&#39; &#39;kReadWrite&#39; indicates Protocol access level &#39;ReadWrite&#39; | [optional] 
**NfsAllSquash** | Pointer to **NullableBool** | Specifies whether all clients from this subnet can map view with view_all_squash_uid/view_all_squash_gid configured in the view. | [optional] 
**NfsRootSquash** | Pointer to **NullableBool** | Specifies whether clients from this subnet can mount as root on NFS. | [optional] 
**SmbAccess** | Pointer to **NullableString** | Specifies whether clients from this subnet can mount using SMB protocol. Protocol access level. &#39;kDisabled&#39; indicates Protocol access level &#39;Disabled&#39; &#39;kReadOnly&#39; indicates Protocol access level &#39;ReadOnly&#39; &#39;kReadWrite&#39; indicates Protocol access level &#39;ReadWrite&#39; | [optional] 

## Methods

### NewSubnet

`func NewSubnet() *Subnet`

NewSubnet instantiates a new Subnet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubnetWithDefaults

`func NewSubnetWithDefaults() *Subnet`

NewSubnetWithDefaults instantiates a new Subnet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponent

`func (o *Subnet) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *Subnet) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *Subnet) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *Subnet) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### SetComponentNil

`func (o *Subnet) SetComponentNil(b bool)`

 SetComponentNil sets the value for Component to be an explicit nil

### UnsetComponent
`func (o *Subnet) UnsetComponent()`

UnsetComponent ensures that no value is present for Component, not even an explicit nil
### GetDescription

`func (o *Subnet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Subnet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Subnet) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Subnet) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Subnet) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Subnet) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *Subnet) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Subnet) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Subnet) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Subnet) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Subnet) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Subnet) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIp

`func (o *Subnet) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *Subnet) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *Subnet) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *Subnet) HasIp() bool`

HasIp returns a boolean if a field has been set.

### SetIpNil

`func (o *Subnet) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *Subnet) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
### GetNetmaskBits

`func (o *Subnet) GetNetmaskBits() int32`

GetNetmaskBits returns the NetmaskBits field if non-nil, zero value otherwise.

### GetNetmaskBitsOk

`func (o *Subnet) GetNetmaskBitsOk() (*int32, bool)`

GetNetmaskBitsOk returns a tuple with the NetmaskBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmaskBits

`func (o *Subnet) SetNetmaskBits(v int32)`

SetNetmaskBits sets NetmaskBits field to given value.

### HasNetmaskBits

`func (o *Subnet) HasNetmaskBits() bool`

HasNetmaskBits returns a boolean if a field has been set.

### SetNetmaskBitsNil

`func (o *Subnet) SetNetmaskBitsNil(b bool)`

 SetNetmaskBitsNil sets the value for NetmaskBits to be an explicit nil

### UnsetNetmaskBits
`func (o *Subnet) UnsetNetmaskBits()`

UnsetNetmaskBits ensures that no value is present for NetmaskBits, not even an explicit nil
### GetNetmaskIp4

`func (o *Subnet) GetNetmaskIp4() string`

GetNetmaskIp4 returns the NetmaskIp4 field if non-nil, zero value otherwise.

### GetNetmaskIp4Ok

`func (o *Subnet) GetNetmaskIp4Ok() (*string, bool)`

GetNetmaskIp4Ok returns a tuple with the NetmaskIp4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmaskIp4

`func (o *Subnet) SetNetmaskIp4(v string)`

SetNetmaskIp4 sets NetmaskIp4 field to given value.

### HasNetmaskIp4

`func (o *Subnet) HasNetmaskIp4() bool`

HasNetmaskIp4 returns a boolean if a field has been set.

### SetNetmaskIp4Nil

`func (o *Subnet) SetNetmaskIp4Nil(b bool)`

 SetNetmaskIp4Nil sets the value for NetmaskIp4 to be an explicit nil

### UnsetNetmaskIp4
`func (o *Subnet) UnsetNetmaskIp4()`

UnsetNetmaskIp4 ensures that no value is present for NetmaskIp4, not even an explicit nil
### GetNfsAccess

`func (o *Subnet) GetNfsAccess() string`

GetNfsAccess returns the NfsAccess field if non-nil, zero value otherwise.

### GetNfsAccessOk

`func (o *Subnet) GetNfsAccessOk() (*string, bool)`

GetNfsAccessOk returns a tuple with the NfsAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsAccess

`func (o *Subnet) SetNfsAccess(v string)`

SetNfsAccess sets NfsAccess field to given value.

### HasNfsAccess

`func (o *Subnet) HasNfsAccess() bool`

HasNfsAccess returns a boolean if a field has been set.

### SetNfsAccessNil

`func (o *Subnet) SetNfsAccessNil(b bool)`

 SetNfsAccessNil sets the value for NfsAccess to be an explicit nil

### UnsetNfsAccess
`func (o *Subnet) UnsetNfsAccess()`

UnsetNfsAccess ensures that no value is present for NfsAccess, not even an explicit nil
### GetNfsAllSquash

`func (o *Subnet) GetNfsAllSquash() bool`

GetNfsAllSquash returns the NfsAllSquash field if non-nil, zero value otherwise.

### GetNfsAllSquashOk

`func (o *Subnet) GetNfsAllSquashOk() (*bool, bool)`

GetNfsAllSquashOk returns a tuple with the NfsAllSquash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsAllSquash

`func (o *Subnet) SetNfsAllSquash(v bool)`

SetNfsAllSquash sets NfsAllSquash field to given value.

### HasNfsAllSquash

`func (o *Subnet) HasNfsAllSquash() bool`

HasNfsAllSquash returns a boolean if a field has been set.

### SetNfsAllSquashNil

`func (o *Subnet) SetNfsAllSquashNil(b bool)`

 SetNfsAllSquashNil sets the value for NfsAllSquash to be an explicit nil

### UnsetNfsAllSquash
`func (o *Subnet) UnsetNfsAllSquash()`

UnsetNfsAllSquash ensures that no value is present for NfsAllSquash, not even an explicit nil
### GetNfsRootSquash

`func (o *Subnet) GetNfsRootSquash() bool`

GetNfsRootSquash returns the NfsRootSquash field if non-nil, zero value otherwise.

### GetNfsRootSquashOk

`func (o *Subnet) GetNfsRootSquashOk() (*bool, bool)`

GetNfsRootSquashOk returns a tuple with the NfsRootSquash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsRootSquash

`func (o *Subnet) SetNfsRootSquash(v bool)`

SetNfsRootSquash sets NfsRootSquash field to given value.

### HasNfsRootSquash

`func (o *Subnet) HasNfsRootSquash() bool`

HasNfsRootSquash returns a boolean if a field has been set.

### SetNfsRootSquashNil

`func (o *Subnet) SetNfsRootSquashNil(b bool)`

 SetNfsRootSquashNil sets the value for NfsRootSquash to be an explicit nil

### UnsetNfsRootSquash
`func (o *Subnet) UnsetNfsRootSquash()`

UnsetNfsRootSquash ensures that no value is present for NfsRootSquash, not even an explicit nil
### GetSmbAccess

`func (o *Subnet) GetSmbAccess() string`

GetSmbAccess returns the SmbAccess field if non-nil, zero value otherwise.

### GetSmbAccessOk

`func (o *Subnet) GetSmbAccessOk() (*string, bool)`

GetSmbAccessOk returns a tuple with the SmbAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbAccess

`func (o *Subnet) SetSmbAccess(v string)`

SetSmbAccess sets SmbAccess field to given value.

### HasSmbAccess

`func (o *Subnet) HasSmbAccess() bool`

HasSmbAccess returns a boolean if a field has been set.

### SetSmbAccessNil

`func (o *Subnet) SetSmbAccessNil(b bool)`

 SetSmbAccessNil sets the value for SmbAccess to be an explicit nil

### UnsetSmbAccess
`func (o *Subnet) UnsetSmbAccess()`

UnsetSmbAccess ensures that no value is present for SmbAccess, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


