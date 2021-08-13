# FilerAuditLogConfigs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SharePermissions** | Pointer to [**[]SmbPermission**](SmbPermission.md) | Specifies a list of share level permissions. | [optional] 
**SubnetWhitelist** | Pointer to [**[]Subnet**](Subnet.md) | Specifies a list of Subnets with IP addresses that have permissions to access a Cohesity View containing filer audit logs. | [optional] 
**OverrideGlobalSubnetWhitelist** | Pointer to **NullableBool** | Specifies whether view level client subnet whitelist overrides cluster and global setting. | [optional] 
**SmbMountPaths** | Pointer to **[]string** | Specifies a list of SMB mount paths of a Cohesity View containing filer audit logs. | [optional] [readonly] 
**NfsMountPath** | Pointer to **NullableString** | This field is currently deprecated. Please use NFS MountPaths which would be an array of strings. | [optional] [readonly] 
**NfsMountPaths** | Pointer to **[]string** | Specifies a list of NFS mount paths of a Cohesity View containing filer audit logs. | [optional] [readonly] 

## Methods

### NewFilerAuditLogConfigs

`func NewFilerAuditLogConfigs() *FilerAuditLogConfigs`

NewFilerAuditLogConfigs instantiates a new FilerAuditLogConfigs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilerAuditLogConfigsWithDefaults

`func NewFilerAuditLogConfigsWithDefaults() *FilerAuditLogConfigs`

NewFilerAuditLogConfigsWithDefaults instantiates a new FilerAuditLogConfigs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSharePermissions

`func (o *FilerAuditLogConfigs) GetSharePermissions() []SmbPermission`

GetSharePermissions returns the SharePermissions field if non-nil, zero value otherwise.

### GetSharePermissionsOk

`func (o *FilerAuditLogConfigs) GetSharePermissionsOk() (*[]SmbPermission, bool)`

GetSharePermissionsOk returns a tuple with the SharePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharePermissions

`func (o *FilerAuditLogConfigs) SetSharePermissions(v []SmbPermission)`

SetSharePermissions sets SharePermissions field to given value.

### HasSharePermissions

`func (o *FilerAuditLogConfigs) HasSharePermissions() bool`

HasSharePermissions returns a boolean if a field has been set.

### SetSharePermissionsNil

`func (o *FilerAuditLogConfigs) SetSharePermissionsNil(b bool)`

 SetSharePermissionsNil sets the value for SharePermissions to be an explicit nil

### UnsetSharePermissions
`func (o *FilerAuditLogConfigs) UnsetSharePermissions()`

UnsetSharePermissions ensures that no value is present for SharePermissions, not even an explicit nil
### GetSubnetWhitelist

`func (o *FilerAuditLogConfigs) GetSubnetWhitelist() []Subnet`

GetSubnetWhitelist returns the SubnetWhitelist field if non-nil, zero value otherwise.

### GetSubnetWhitelistOk

`func (o *FilerAuditLogConfigs) GetSubnetWhitelistOk() (*[]Subnet, bool)`

GetSubnetWhitelistOk returns a tuple with the SubnetWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetWhitelist

`func (o *FilerAuditLogConfigs) SetSubnetWhitelist(v []Subnet)`

SetSubnetWhitelist sets SubnetWhitelist field to given value.

### HasSubnetWhitelist

`func (o *FilerAuditLogConfigs) HasSubnetWhitelist() bool`

HasSubnetWhitelist returns a boolean if a field has been set.

### SetSubnetWhitelistNil

`func (o *FilerAuditLogConfigs) SetSubnetWhitelistNil(b bool)`

 SetSubnetWhitelistNil sets the value for SubnetWhitelist to be an explicit nil

### UnsetSubnetWhitelist
`func (o *FilerAuditLogConfigs) UnsetSubnetWhitelist()`

UnsetSubnetWhitelist ensures that no value is present for SubnetWhitelist, not even an explicit nil
### GetOverrideGlobalSubnetWhitelist

`func (o *FilerAuditLogConfigs) GetOverrideGlobalSubnetWhitelist() bool`

GetOverrideGlobalSubnetWhitelist returns the OverrideGlobalSubnetWhitelist field if non-nil, zero value otherwise.

### GetOverrideGlobalSubnetWhitelistOk

`func (o *FilerAuditLogConfigs) GetOverrideGlobalSubnetWhitelistOk() (*bool, bool)`

GetOverrideGlobalSubnetWhitelistOk returns a tuple with the OverrideGlobalSubnetWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideGlobalSubnetWhitelist

`func (o *FilerAuditLogConfigs) SetOverrideGlobalSubnetWhitelist(v bool)`

SetOverrideGlobalSubnetWhitelist sets OverrideGlobalSubnetWhitelist field to given value.

### HasOverrideGlobalSubnetWhitelist

`func (o *FilerAuditLogConfigs) HasOverrideGlobalSubnetWhitelist() bool`

HasOverrideGlobalSubnetWhitelist returns a boolean if a field has been set.

### SetOverrideGlobalSubnetWhitelistNil

`func (o *FilerAuditLogConfigs) SetOverrideGlobalSubnetWhitelistNil(b bool)`

 SetOverrideGlobalSubnetWhitelistNil sets the value for OverrideGlobalSubnetWhitelist to be an explicit nil

### UnsetOverrideGlobalSubnetWhitelist
`func (o *FilerAuditLogConfigs) UnsetOverrideGlobalSubnetWhitelist()`

UnsetOverrideGlobalSubnetWhitelist ensures that no value is present for OverrideGlobalSubnetWhitelist, not even an explicit nil
### GetSmbMountPaths

`func (o *FilerAuditLogConfigs) GetSmbMountPaths() []string`

GetSmbMountPaths returns the SmbMountPaths field if non-nil, zero value otherwise.

### GetSmbMountPathsOk

`func (o *FilerAuditLogConfigs) GetSmbMountPathsOk() (*[]string, bool)`

GetSmbMountPathsOk returns a tuple with the SmbMountPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbMountPaths

`func (o *FilerAuditLogConfigs) SetSmbMountPaths(v []string)`

SetSmbMountPaths sets SmbMountPaths field to given value.

### HasSmbMountPaths

`func (o *FilerAuditLogConfigs) HasSmbMountPaths() bool`

HasSmbMountPaths returns a boolean if a field has been set.

### SetSmbMountPathsNil

`func (o *FilerAuditLogConfigs) SetSmbMountPathsNil(b bool)`

 SetSmbMountPathsNil sets the value for SmbMountPaths to be an explicit nil

### UnsetSmbMountPaths
`func (o *FilerAuditLogConfigs) UnsetSmbMountPaths()`

UnsetSmbMountPaths ensures that no value is present for SmbMountPaths, not even an explicit nil
### GetNfsMountPath

`func (o *FilerAuditLogConfigs) GetNfsMountPath() string`

GetNfsMountPath returns the NfsMountPath field if non-nil, zero value otherwise.

### GetNfsMountPathOk

`func (o *FilerAuditLogConfigs) GetNfsMountPathOk() (*string, bool)`

GetNfsMountPathOk returns a tuple with the NfsMountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsMountPath

`func (o *FilerAuditLogConfigs) SetNfsMountPath(v string)`

SetNfsMountPath sets NfsMountPath field to given value.

### HasNfsMountPath

`func (o *FilerAuditLogConfigs) HasNfsMountPath() bool`

HasNfsMountPath returns a boolean if a field has been set.

### SetNfsMountPathNil

`func (o *FilerAuditLogConfigs) SetNfsMountPathNil(b bool)`

 SetNfsMountPathNil sets the value for NfsMountPath to be an explicit nil

### UnsetNfsMountPath
`func (o *FilerAuditLogConfigs) UnsetNfsMountPath()`

UnsetNfsMountPath ensures that no value is present for NfsMountPath, not even an explicit nil
### GetNfsMountPaths

`func (o *FilerAuditLogConfigs) GetNfsMountPaths() []string`

GetNfsMountPaths returns the NfsMountPaths field if non-nil, zero value otherwise.

### GetNfsMountPathsOk

`func (o *FilerAuditLogConfigs) GetNfsMountPathsOk() (*[]string, bool)`

GetNfsMountPathsOk returns a tuple with the NfsMountPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsMountPaths

`func (o *FilerAuditLogConfigs) SetNfsMountPaths(v []string)`

SetNfsMountPaths sets NfsMountPaths field to given value.

### HasNfsMountPaths

`func (o *FilerAuditLogConfigs) HasNfsMountPaths() bool`

HasNfsMountPaths returns a boolean if a field has been set.

### SetNfsMountPathsNil

`func (o *FilerAuditLogConfigs) SetNfsMountPathsNil(b bool)`

 SetNfsMountPathsNil sets the value for NfsMountPaths to be an explicit nil

### UnsetNfsMountPaths
`func (o *FilerAuditLogConfigs) UnsetNfsMountPaths()`

UnsetNfsMountPaths ensures that no value is present for NfsMountPaths, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


