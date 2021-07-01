# Share

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllSmbMountPaths** | Pointer to **[]string** | Array of SMB Paths.  Specifies the possible paths that can be used to mount this Share as a SMB share. If Active Directory has multiple account names; each machine account has its own path. | [optional] 
**EnableSmbEncryption** | Pointer to **NullableBool** | Specifies the SMB encryption for the View Alias. If set, it enables the SMB encryption for the View Alias. Encryption is supported only by SMB 3.x dialects. Dialects that do not support would still access data in unencrypted format. | [optional] 
**EnableSmbViewDiscovery** | Pointer to **NullableBool** | If set, it enables discovery of view alias for SMB. | [optional] 
**EnforceSmbEncryption** | Pointer to **NullableBool** | Specifies the SMB encryption for all the sessions for the View Alias. If set, encryption is enforced for all the sessions for the View Alias. When enabled all future and existing unencrypted sessions are disallowed. | [optional] 
**NfsMountPath** | Pointer to **NullableString** | Specifies the path for mounting this Share as an NFS share. | [optional] 
**Path** | Pointer to **NullableString** | Specifies the path information for this share. | [optional] 
**S3AccessPath** | Pointer to **NullableString** | Specifies the path to access this View as an S3 share. | [optional] 
**ShareName** | Pointer to **NullableString** | The name of the share. | [optional] 
**SharePermissions** | Pointer to [**[]SmbPermission**](SmbPermission.md) | Specifies a list of share level permissions. | [optional] 
**SmbMountPath** | Pointer to **NullableString** | Specifies the main path for mounting this Share as an SMB share. | [optional] 
**SubnetWhitelist** | Pointer to [**[]Subnet**](Subnet.md) | Specifies a list of Subnets with IP addresses that have permissions to access the View Alias. (Overrides the Subnets specified at the global Cohesity Cluster level and View level.) | [optional] 
**TenantId** | Pointer to **NullableString** | Specifies the unique id of the tenant. | [optional] 
**ViewName** | Pointer to **NullableString** | Specifies the view name this share belongs to. | [optional] 

## Methods

### NewShare

`func NewShare() *Share`

NewShare instantiates a new Share object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShareWithDefaults

`func NewShareWithDefaults() *Share`

NewShareWithDefaults instantiates a new Share object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllSmbMountPaths

`func (o *Share) GetAllSmbMountPaths() []string`

GetAllSmbMountPaths returns the AllSmbMountPaths field if non-nil, zero value otherwise.

### GetAllSmbMountPathsOk

`func (o *Share) GetAllSmbMountPathsOk() (*[]string, bool)`

GetAllSmbMountPathsOk returns a tuple with the AllSmbMountPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllSmbMountPaths

`func (o *Share) SetAllSmbMountPaths(v []string)`

SetAllSmbMountPaths sets AllSmbMountPaths field to given value.

### HasAllSmbMountPaths

`func (o *Share) HasAllSmbMountPaths() bool`

HasAllSmbMountPaths returns a boolean if a field has been set.

### SetAllSmbMountPathsNil

`func (o *Share) SetAllSmbMountPathsNil(b bool)`

 SetAllSmbMountPathsNil sets the value for AllSmbMountPaths to be an explicit nil

### UnsetAllSmbMountPaths
`func (o *Share) UnsetAllSmbMountPaths()`

UnsetAllSmbMountPaths ensures that no value is present for AllSmbMountPaths, not even an explicit nil
### GetEnableSmbEncryption

`func (o *Share) GetEnableSmbEncryption() bool`

GetEnableSmbEncryption returns the EnableSmbEncryption field if non-nil, zero value otherwise.

### GetEnableSmbEncryptionOk

`func (o *Share) GetEnableSmbEncryptionOk() (*bool, bool)`

GetEnableSmbEncryptionOk returns a tuple with the EnableSmbEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSmbEncryption

`func (o *Share) SetEnableSmbEncryption(v bool)`

SetEnableSmbEncryption sets EnableSmbEncryption field to given value.

### HasEnableSmbEncryption

`func (o *Share) HasEnableSmbEncryption() bool`

HasEnableSmbEncryption returns a boolean if a field has been set.

### SetEnableSmbEncryptionNil

`func (o *Share) SetEnableSmbEncryptionNil(b bool)`

 SetEnableSmbEncryptionNil sets the value for EnableSmbEncryption to be an explicit nil

### UnsetEnableSmbEncryption
`func (o *Share) UnsetEnableSmbEncryption()`

UnsetEnableSmbEncryption ensures that no value is present for EnableSmbEncryption, not even an explicit nil
### GetEnableSmbViewDiscovery

`func (o *Share) GetEnableSmbViewDiscovery() bool`

GetEnableSmbViewDiscovery returns the EnableSmbViewDiscovery field if non-nil, zero value otherwise.

### GetEnableSmbViewDiscoveryOk

`func (o *Share) GetEnableSmbViewDiscoveryOk() (*bool, bool)`

GetEnableSmbViewDiscoveryOk returns a tuple with the EnableSmbViewDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSmbViewDiscovery

`func (o *Share) SetEnableSmbViewDiscovery(v bool)`

SetEnableSmbViewDiscovery sets EnableSmbViewDiscovery field to given value.

### HasEnableSmbViewDiscovery

`func (o *Share) HasEnableSmbViewDiscovery() bool`

HasEnableSmbViewDiscovery returns a boolean if a field has been set.

### SetEnableSmbViewDiscoveryNil

`func (o *Share) SetEnableSmbViewDiscoveryNil(b bool)`

 SetEnableSmbViewDiscoveryNil sets the value for EnableSmbViewDiscovery to be an explicit nil

### UnsetEnableSmbViewDiscovery
`func (o *Share) UnsetEnableSmbViewDiscovery()`

UnsetEnableSmbViewDiscovery ensures that no value is present for EnableSmbViewDiscovery, not even an explicit nil
### GetEnforceSmbEncryption

`func (o *Share) GetEnforceSmbEncryption() bool`

GetEnforceSmbEncryption returns the EnforceSmbEncryption field if non-nil, zero value otherwise.

### GetEnforceSmbEncryptionOk

`func (o *Share) GetEnforceSmbEncryptionOk() (*bool, bool)`

GetEnforceSmbEncryptionOk returns a tuple with the EnforceSmbEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceSmbEncryption

`func (o *Share) SetEnforceSmbEncryption(v bool)`

SetEnforceSmbEncryption sets EnforceSmbEncryption field to given value.

### HasEnforceSmbEncryption

`func (o *Share) HasEnforceSmbEncryption() bool`

HasEnforceSmbEncryption returns a boolean if a field has been set.

### SetEnforceSmbEncryptionNil

`func (o *Share) SetEnforceSmbEncryptionNil(b bool)`

 SetEnforceSmbEncryptionNil sets the value for EnforceSmbEncryption to be an explicit nil

### UnsetEnforceSmbEncryption
`func (o *Share) UnsetEnforceSmbEncryption()`

UnsetEnforceSmbEncryption ensures that no value is present for EnforceSmbEncryption, not even an explicit nil
### GetNfsMountPath

`func (o *Share) GetNfsMountPath() string`

GetNfsMountPath returns the NfsMountPath field if non-nil, zero value otherwise.

### GetNfsMountPathOk

`func (o *Share) GetNfsMountPathOk() (*string, bool)`

GetNfsMountPathOk returns a tuple with the NfsMountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsMountPath

`func (o *Share) SetNfsMountPath(v string)`

SetNfsMountPath sets NfsMountPath field to given value.

### HasNfsMountPath

`func (o *Share) HasNfsMountPath() bool`

HasNfsMountPath returns a boolean if a field has been set.

### SetNfsMountPathNil

`func (o *Share) SetNfsMountPathNil(b bool)`

 SetNfsMountPathNil sets the value for NfsMountPath to be an explicit nil

### UnsetNfsMountPath
`func (o *Share) UnsetNfsMountPath()`

UnsetNfsMountPath ensures that no value is present for NfsMountPath, not even an explicit nil
### GetPath

`func (o *Share) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Share) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Share) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *Share) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *Share) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *Share) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetS3AccessPath

`func (o *Share) GetS3AccessPath() string`

GetS3AccessPath returns the S3AccessPath field if non-nil, zero value otherwise.

### GetS3AccessPathOk

`func (o *Share) GetS3AccessPathOk() (*string, bool)`

GetS3AccessPathOk returns a tuple with the S3AccessPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3AccessPath

`func (o *Share) SetS3AccessPath(v string)`

SetS3AccessPath sets S3AccessPath field to given value.

### HasS3AccessPath

`func (o *Share) HasS3AccessPath() bool`

HasS3AccessPath returns a boolean if a field has been set.

### SetS3AccessPathNil

`func (o *Share) SetS3AccessPathNil(b bool)`

 SetS3AccessPathNil sets the value for S3AccessPath to be an explicit nil

### UnsetS3AccessPath
`func (o *Share) UnsetS3AccessPath()`

UnsetS3AccessPath ensures that no value is present for S3AccessPath, not even an explicit nil
### GetShareName

`func (o *Share) GetShareName() string`

GetShareName returns the ShareName field if non-nil, zero value otherwise.

### GetShareNameOk

`func (o *Share) GetShareNameOk() (*string, bool)`

GetShareNameOk returns a tuple with the ShareName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareName

`func (o *Share) SetShareName(v string)`

SetShareName sets ShareName field to given value.

### HasShareName

`func (o *Share) HasShareName() bool`

HasShareName returns a boolean if a field has been set.

### SetShareNameNil

`func (o *Share) SetShareNameNil(b bool)`

 SetShareNameNil sets the value for ShareName to be an explicit nil

### UnsetShareName
`func (o *Share) UnsetShareName()`

UnsetShareName ensures that no value is present for ShareName, not even an explicit nil
### GetSharePermissions

`func (o *Share) GetSharePermissions() []SmbPermission`

GetSharePermissions returns the SharePermissions field if non-nil, zero value otherwise.

### GetSharePermissionsOk

`func (o *Share) GetSharePermissionsOk() (*[]SmbPermission, bool)`

GetSharePermissionsOk returns a tuple with the SharePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharePermissions

`func (o *Share) SetSharePermissions(v []SmbPermission)`

SetSharePermissions sets SharePermissions field to given value.

### HasSharePermissions

`func (o *Share) HasSharePermissions() bool`

HasSharePermissions returns a boolean if a field has been set.

### SetSharePermissionsNil

`func (o *Share) SetSharePermissionsNil(b bool)`

 SetSharePermissionsNil sets the value for SharePermissions to be an explicit nil

### UnsetSharePermissions
`func (o *Share) UnsetSharePermissions()`

UnsetSharePermissions ensures that no value is present for SharePermissions, not even an explicit nil
### GetSmbMountPath

`func (o *Share) GetSmbMountPath() string`

GetSmbMountPath returns the SmbMountPath field if non-nil, zero value otherwise.

### GetSmbMountPathOk

`func (o *Share) GetSmbMountPathOk() (*string, bool)`

GetSmbMountPathOk returns a tuple with the SmbMountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbMountPath

`func (o *Share) SetSmbMountPath(v string)`

SetSmbMountPath sets SmbMountPath field to given value.

### HasSmbMountPath

`func (o *Share) HasSmbMountPath() bool`

HasSmbMountPath returns a boolean if a field has been set.

### SetSmbMountPathNil

`func (o *Share) SetSmbMountPathNil(b bool)`

 SetSmbMountPathNil sets the value for SmbMountPath to be an explicit nil

### UnsetSmbMountPath
`func (o *Share) UnsetSmbMountPath()`

UnsetSmbMountPath ensures that no value is present for SmbMountPath, not even an explicit nil
### GetSubnetWhitelist

`func (o *Share) GetSubnetWhitelist() []Subnet`

GetSubnetWhitelist returns the SubnetWhitelist field if non-nil, zero value otherwise.

### GetSubnetWhitelistOk

`func (o *Share) GetSubnetWhitelistOk() (*[]Subnet, bool)`

GetSubnetWhitelistOk returns a tuple with the SubnetWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetWhitelist

`func (o *Share) SetSubnetWhitelist(v []Subnet)`

SetSubnetWhitelist sets SubnetWhitelist field to given value.

### HasSubnetWhitelist

`func (o *Share) HasSubnetWhitelist() bool`

HasSubnetWhitelist returns a boolean if a field has been set.

### SetSubnetWhitelistNil

`func (o *Share) SetSubnetWhitelistNil(b bool)`

 SetSubnetWhitelistNil sets the value for SubnetWhitelist to be an explicit nil

### UnsetSubnetWhitelist
`func (o *Share) UnsetSubnetWhitelist()`

UnsetSubnetWhitelist ensures that no value is present for SubnetWhitelist, not even an explicit nil
### GetTenantId

`func (o *Share) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Share) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Share) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *Share) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *Share) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *Share) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetViewName

`func (o *Share) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *Share) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *Share) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *Share) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### SetViewNameNil

`func (o *Share) SetViewNameNil(b bool)`

 SetViewNameNil sets the value for ViewName to be an explicit nil

### UnsetViewName
`func (o *Share) UnsetViewName()`

UnsetViewName ensures that no value is present for ViewName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


