# UpdateViewAliasParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AliasName** | Pointer to **NullableString** | Name of the alias to be updated. | [optional] 
**EnableSmbEncryption** | Pointer to **NullableBool** | Specifies the SMB encryption for the View Alias. If set, it enables the SMB encryption for the View Alias. Encryption is supported only by SMB 3.x dialects. Dialects that do not support would still access data in unencrypted format. | [optional] 
**EnableSmbViewDiscovery** | Pointer to **NullableBool** | If set, it enables discovery of view alias for SMB. | [optional] 
**EnforceSmbEncryption** | Pointer to **NullableBool** | Specifies the SMB encryption for all the sessions for the View Alias. If set, encryption is enforced for all the sessions for the View Alias. When enabled all future and existing unencrypted sessions are disallowed. | [optional] 
**SharePermissions** | Pointer to [**[]SmbPermission**](SmbPermission.md) | Specifies a list of share level permissions. | [optional] 
**SubnetWhitelist** | Pointer to [**[]Subnet**](Subnet.md) | Specifies a list of Subnets with IP addresses that have permissions to access the View Alias. (Overrides the Subnets specified at the global Cohesity Cluster level and View level.) | [optional] 

## Methods

### NewUpdateViewAliasParam

`func NewUpdateViewAliasParam() *UpdateViewAliasParam`

NewUpdateViewAliasParam instantiates a new UpdateViewAliasParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateViewAliasParamWithDefaults

`func NewUpdateViewAliasParamWithDefaults() *UpdateViewAliasParam`

NewUpdateViewAliasParamWithDefaults instantiates a new UpdateViewAliasParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliasName

`func (o *UpdateViewAliasParam) GetAliasName() string`

GetAliasName returns the AliasName field if non-nil, zero value otherwise.

### GetAliasNameOk

`func (o *UpdateViewAliasParam) GetAliasNameOk() (*string, bool)`

GetAliasNameOk returns a tuple with the AliasName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasName

`func (o *UpdateViewAliasParam) SetAliasName(v string)`

SetAliasName sets AliasName field to given value.

### HasAliasName

`func (o *UpdateViewAliasParam) HasAliasName() bool`

HasAliasName returns a boolean if a field has been set.

### SetAliasNameNil

`func (o *UpdateViewAliasParam) SetAliasNameNil(b bool)`

 SetAliasNameNil sets the value for AliasName to be an explicit nil

### UnsetAliasName
`func (o *UpdateViewAliasParam) UnsetAliasName()`

UnsetAliasName ensures that no value is present for AliasName, not even an explicit nil
### GetEnableSmbEncryption

`func (o *UpdateViewAliasParam) GetEnableSmbEncryption() bool`

GetEnableSmbEncryption returns the EnableSmbEncryption field if non-nil, zero value otherwise.

### GetEnableSmbEncryptionOk

`func (o *UpdateViewAliasParam) GetEnableSmbEncryptionOk() (*bool, bool)`

GetEnableSmbEncryptionOk returns a tuple with the EnableSmbEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSmbEncryption

`func (o *UpdateViewAliasParam) SetEnableSmbEncryption(v bool)`

SetEnableSmbEncryption sets EnableSmbEncryption field to given value.

### HasEnableSmbEncryption

`func (o *UpdateViewAliasParam) HasEnableSmbEncryption() bool`

HasEnableSmbEncryption returns a boolean if a field has been set.

### SetEnableSmbEncryptionNil

`func (o *UpdateViewAliasParam) SetEnableSmbEncryptionNil(b bool)`

 SetEnableSmbEncryptionNil sets the value for EnableSmbEncryption to be an explicit nil

### UnsetEnableSmbEncryption
`func (o *UpdateViewAliasParam) UnsetEnableSmbEncryption()`

UnsetEnableSmbEncryption ensures that no value is present for EnableSmbEncryption, not even an explicit nil
### GetEnableSmbViewDiscovery

`func (o *UpdateViewAliasParam) GetEnableSmbViewDiscovery() bool`

GetEnableSmbViewDiscovery returns the EnableSmbViewDiscovery field if non-nil, zero value otherwise.

### GetEnableSmbViewDiscoveryOk

`func (o *UpdateViewAliasParam) GetEnableSmbViewDiscoveryOk() (*bool, bool)`

GetEnableSmbViewDiscoveryOk returns a tuple with the EnableSmbViewDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSmbViewDiscovery

`func (o *UpdateViewAliasParam) SetEnableSmbViewDiscovery(v bool)`

SetEnableSmbViewDiscovery sets EnableSmbViewDiscovery field to given value.

### HasEnableSmbViewDiscovery

`func (o *UpdateViewAliasParam) HasEnableSmbViewDiscovery() bool`

HasEnableSmbViewDiscovery returns a boolean if a field has been set.

### SetEnableSmbViewDiscoveryNil

`func (o *UpdateViewAliasParam) SetEnableSmbViewDiscoveryNil(b bool)`

 SetEnableSmbViewDiscoveryNil sets the value for EnableSmbViewDiscovery to be an explicit nil

### UnsetEnableSmbViewDiscovery
`func (o *UpdateViewAliasParam) UnsetEnableSmbViewDiscovery()`

UnsetEnableSmbViewDiscovery ensures that no value is present for EnableSmbViewDiscovery, not even an explicit nil
### GetEnforceSmbEncryption

`func (o *UpdateViewAliasParam) GetEnforceSmbEncryption() bool`

GetEnforceSmbEncryption returns the EnforceSmbEncryption field if non-nil, zero value otherwise.

### GetEnforceSmbEncryptionOk

`func (o *UpdateViewAliasParam) GetEnforceSmbEncryptionOk() (*bool, bool)`

GetEnforceSmbEncryptionOk returns a tuple with the EnforceSmbEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceSmbEncryption

`func (o *UpdateViewAliasParam) SetEnforceSmbEncryption(v bool)`

SetEnforceSmbEncryption sets EnforceSmbEncryption field to given value.

### HasEnforceSmbEncryption

`func (o *UpdateViewAliasParam) HasEnforceSmbEncryption() bool`

HasEnforceSmbEncryption returns a boolean if a field has been set.

### SetEnforceSmbEncryptionNil

`func (o *UpdateViewAliasParam) SetEnforceSmbEncryptionNil(b bool)`

 SetEnforceSmbEncryptionNil sets the value for EnforceSmbEncryption to be an explicit nil

### UnsetEnforceSmbEncryption
`func (o *UpdateViewAliasParam) UnsetEnforceSmbEncryption()`

UnsetEnforceSmbEncryption ensures that no value is present for EnforceSmbEncryption, not even an explicit nil
### GetSharePermissions

`func (o *UpdateViewAliasParam) GetSharePermissions() []SmbPermission`

GetSharePermissions returns the SharePermissions field if non-nil, zero value otherwise.

### GetSharePermissionsOk

`func (o *UpdateViewAliasParam) GetSharePermissionsOk() (*[]SmbPermission, bool)`

GetSharePermissionsOk returns a tuple with the SharePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharePermissions

`func (o *UpdateViewAliasParam) SetSharePermissions(v []SmbPermission)`

SetSharePermissions sets SharePermissions field to given value.

### HasSharePermissions

`func (o *UpdateViewAliasParam) HasSharePermissions() bool`

HasSharePermissions returns a boolean if a field has been set.

### SetSharePermissionsNil

`func (o *UpdateViewAliasParam) SetSharePermissionsNil(b bool)`

 SetSharePermissionsNil sets the value for SharePermissions to be an explicit nil

### UnsetSharePermissions
`func (o *UpdateViewAliasParam) UnsetSharePermissions()`

UnsetSharePermissions ensures that no value is present for SharePermissions, not even an explicit nil
### GetSubnetWhitelist

`func (o *UpdateViewAliasParam) GetSubnetWhitelist() []Subnet`

GetSubnetWhitelist returns the SubnetWhitelist field if non-nil, zero value otherwise.

### GetSubnetWhitelistOk

`func (o *UpdateViewAliasParam) GetSubnetWhitelistOk() (*[]Subnet, bool)`

GetSubnetWhitelistOk returns a tuple with the SubnetWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetWhitelist

`func (o *UpdateViewAliasParam) SetSubnetWhitelist(v []Subnet)`

SetSubnetWhitelist sets SubnetWhitelist field to given value.

### HasSubnetWhitelist

`func (o *UpdateViewAliasParam) HasSubnetWhitelist() bool`

HasSubnetWhitelist returns a boolean if a field has been set.

### SetSubnetWhitelistNil

`func (o *UpdateViewAliasParam) SetSubnetWhitelistNil(b bool)`

 SetSubnetWhitelistNil sets the value for SubnetWhitelist to be an explicit nil

### UnsetSubnetWhitelist
`func (o *UpdateViewAliasParam) UnsetSubnetWhitelist()`

UnsetSubnetWhitelist ensures that no value is present for SubnetWhitelist, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


