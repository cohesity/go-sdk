# ViewAlias

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AliasName** | Pointer to **NullableString** | Alias name. | [optional] 
**EnableSmbEncryption** | Pointer to **NullableBool** | Specifies the SMB encryption for the View Alias. If set, it enables the SMB encryption for the View Alias. Encryption is supported only by SMB 3.x dialects. Dialects that do not support would still access data in unencrypted format. | [optional] 
**EnableSmbViewDiscovery** | Pointer to **NullableBool** | If set, it enables discovery of view alias for SMB. | [optional] 
**EnforceSmbEncryption** | Pointer to **NullableBool** | Specifies the SMB encryption for all the sessions for the View Alias. If set, encryption is enforced for all the sessions for the View Alias. When enabled all future and existing unencrypted sessions are disallowed. | [optional] 
**SharePermissions** | Pointer to [**[]SmbPermission**](SmbPermission.md) | Specifies a list of share level permissions. | [optional] 
**SubnetWhitelist** | Pointer to [**[]Subnet**](Subnet.md) | Specifies a list of Subnets with IP addresses that have permissions to access the View Alias. (Overrides the Subnets specified at the global Cohesity Cluster level and View level.) | [optional] 
**ViewName** | Pointer to **NullableString** | View name. | [optional] 
**ViewPath** | Pointer to **NullableString** | View path for the alias. | [optional] 

## Methods

### NewViewAlias

`func NewViewAlias() *ViewAlias`

NewViewAlias instantiates a new ViewAlias object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewAliasWithDefaults

`func NewViewAliasWithDefaults() *ViewAlias`

NewViewAliasWithDefaults instantiates a new ViewAlias object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliasName

`func (o *ViewAlias) GetAliasName() string`

GetAliasName returns the AliasName field if non-nil, zero value otherwise.

### GetAliasNameOk

`func (o *ViewAlias) GetAliasNameOk() (*string, bool)`

GetAliasNameOk returns a tuple with the AliasName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasName

`func (o *ViewAlias) SetAliasName(v string)`

SetAliasName sets AliasName field to given value.

### HasAliasName

`func (o *ViewAlias) HasAliasName() bool`

HasAliasName returns a boolean if a field has been set.

### SetAliasNameNil

`func (o *ViewAlias) SetAliasNameNil(b bool)`

 SetAliasNameNil sets the value for AliasName to be an explicit nil

### UnsetAliasName
`func (o *ViewAlias) UnsetAliasName()`

UnsetAliasName ensures that no value is present for AliasName, not even an explicit nil
### GetEnableSmbEncryption

`func (o *ViewAlias) GetEnableSmbEncryption() bool`

GetEnableSmbEncryption returns the EnableSmbEncryption field if non-nil, zero value otherwise.

### GetEnableSmbEncryptionOk

`func (o *ViewAlias) GetEnableSmbEncryptionOk() (*bool, bool)`

GetEnableSmbEncryptionOk returns a tuple with the EnableSmbEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSmbEncryption

`func (o *ViewAlias) SetEnableSmbEncryption(v bool)`

SetEnableSmbEncryption sets EnableSmbEncryption field to given value.

### HasEnableSmbEncryption

`func (o *ViewAlias) HasEnableSmbEncryption() bool`

HasEnableSmbEncryption returns a boolean if a field has been set.

### SetEnableSmbEncryptionNil

`func (o *ViewAlias) SetEnableSmbEncryptionNil(b bool)`

 SetEnableSmbEncryptionNil sets the value for EnableSmbEncryption to be an explicit nil

### UnsetEnableSmbEncryption
`func (o *ViewAlias) UnsetEnableSmbEncryption()`

UnsetEnableSmbEncryption ensures that no value is present for EnableSmbEncryption, not even an explicit nil
### GetEnableSmbViewDiscovery

`func (o *ViewAlias) GetEnableSmbViewDiscovery() bool`

GetEnableSmbViewDiscovery returns the EnableSmbViewDiscovery field if non-nil, zero value otherwise.

### GetEnableSmbViewDiscoveryOk

`func (o *ViewAlias) GetEnableSmbViewDiscoveryOk() (*bool, bool)`

GetEnableSmbViewDiscoveryOk returns a tuple with the EnableSmbViewDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSmbViewDiscovery

`func (o *ViewAlias) SetEnableSmbViewDiscovery(v bool)`

SetEnableSmbViewDiscovery sets EnableSmbViewDiscovery field to given value.

### HasEnableSmbViewDiscovery

`func (o *ViewAlias) HasEnableSmbViewDiscovery() bool`

HasEnableSmbViewDiscovery returns a boolean if a field has been set.

### SetEnableSmbViewDiscoveryNil

`func (o *ViewAlias) SetEnableSmbViewDiscoveryNil(b bool)`

 SetEnableSmbViewDiscoveryNil sets the value for EnableSmbViewDiscovery to be an explicit nil

### UnsetEnableSmbViewDiscovery
`func (o *ViewAlias) UnsetEnableSmbViewDiscovery()`

UnsetEnableSmbViewDiscovery ensures that no value is present for EnableSmbViewDiscovery, not even an explicit nil
### GetEnforceSmbEncryption

`func (o *ViewAlias) GetEnforceSmbEncryption() bool`

GetEnforceSmbEncryption returns the EnforceSmbEncryption field if non-nil, zero value otherwise.

### GetEnforceSmbEncryptionOk

`func (o *ViewAlias) GetEnforceSmbEncryptionOk() (*bool, bool)`

GetEnforceSmbEncryptionOk returns a tuple with the EnforceSmbEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceSmbEncryption

`func (o *ViewAlias) SetEnforceSmbEncryption(v bool)`

SetEnforceSmbEncryption sets EnforceSmbEncryption field to given value.

### HasEnforceSmbEncryption

`func (o *ViewAlias) HasEnforceSmbEncryption() bool`

HasEnforceSmbEncryption returns a boolean if a field has been set.

### SetEnforceSmbEncryptionNil

`func (o *ViewAlias) SetEnforceSmbEncryptionNil(b bool)`

 SetEnforceSmbEncryptionNil sets the value for EnforceSmbEncryption to be an explicit nil

### UnsetEnforceSmbEncryption
`func (o *ViewAlias) UnsetEnforceSmbEncryption()`

UnsetEnforceSmbEncryption ensures that no value is present for EnforceSmbEncryption, not even an explicit nil
### GetSharePermissions

`func (o *ViewAlias) GetSharePermissions() []SmbPermission`

GetSharePermissions returns the SharePermissions field if non-nil, zero value otherwise.

### GetSharePermissionsOk

`func (o *ViewAlias) GetSharePermissionsOk() (*[]SmbPermission, bool)`

GetSharePermissionsOk returns a tuple with the SharePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharePermissions

`func (o *ViewAlias) SetSharePermissions(v []SmbPermission)`

SetSharePermissions sets SharePermissions field to given value.

### HasSharePermissions

`func (o *ViewAlias) HasSharePermissions() bool`

HasSharePermissions returns a boolean if a field has been set.

### SetSharePermissionsNil

`func (o *ViewAlias) SetSharePermissionsNil(b bool)`

 SetSharePermissionsNil sets the value for SharePermissions to be an explicit nil

### UnsetSharePermissions
`func (o *ViewAlias) UnsetSharePermissions()`

UnsetSharePermissions ensures that no value is present for SharePermissions, not even an explicit nil
### GetSubnetWhitelist

`func (o *ViewAlias) GetSubnetWhitelist() []Subnet`

GetSubnetWhitelist returns the SubnetWhitelist field if non-nil, zero value otherwise.

### GetSubnetWhitelistOk

`func (o *ViewAlias) GetSubnetWhitelistOk() (*[]Subnet, bool)`

GetSubnetWhitelistOk returns a tuple with the SubnetWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetWhitelist

`func (o *ViewAlias) SetSubnetWhitelist(v []Subnet)`

SetSubnetWhitelist sets SubnetWhitelist field to given value.

### HasSubnetWhitelist

`func (o *ViewAlias) HasSubnetWhitelist() bool`

HasSubnetWhitelist returns a boolean if a field has been set.

### SetSubnetWhitelistNil

`func (o *ViewAlias) SetSubnetWhitelistNil(b bool)`

 SetSubnetWhitelistNil sets the value for SubnetWhitelist to be an explicit nil

### UnsetSubnetWhitelist
`func (o *ViewAlias) UnsetSubnetWhitelist()`

UnsetSubnetWhitelist ensures that no value is present for SubnetWhitelist, not even an explicit nil
### GetViewName

`func (o *ViewAlias) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *ViewAlias) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *ViewAlias) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *ViewAlias) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### SetViewNameNil

`func (o *ViewAlias) SetViewNameNil(b bool)`

 SetViewNameNil sets the value for ViewName to be an explicit nil

### UnsetViewName
`func (o *ViewAlias) UnsetViewName()`

UnsetViewName ensures that no value is present for ViewName, not even an explicit nil
### GetViewPath

`func (o *ViewAlias) GetViewPath() string`

GetViewPath returns the ViewPath field if non-nil, zero value otherwise.

### GetViewPathOk

`func (o *ViewAlias) GetViewPathOk() (*string, bool)`

GetViewPathOk returns a tuple with the ViewPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewPath

`func (o *ViewAlias) SetViewPath(v string)`

SetViewPath sets ViewPath field to given value.

### HasViewPath

`func (o *ViewAlias) HasViewPath() bool`

HasViewPath returns a boolean if a field has been set.

### SetViewPathNil

`func (o *ViewAlias) SetViewPathNil(b bool)`

 SetViewPathNil sets the value for ViewPath to be an explicit nil

### UnsetViewPath
`func (o *ViewAlias) UnsetViewPath()`

UnsetViewPath ensures that no value is present for ViewPath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


