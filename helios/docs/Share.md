# Share

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableFilerAuditLogging** | Pointer to **NullableBool** | Specifies if Filer Audit Logging is enabled for this view. | [optional] 
**SmbConfig** | Pointer to [**AliasSmbConfig**](AliasSmbConfig.md) | SMB config for the alias (share). | [optional] 
**ClientSubnetWhitelist** | Pointer to [**[]Subnet**](Subnet.md) | List of external client subnet IPs that are allowed to access the share. | [optional] 
**Name** | **NullableString** | Specifies the Share name. | 
**ViewName** | **NullableString** | Specifies the View name of this Share. | 
**ViewPath** | **NullableString** | Specifies the View path of this Share. | 
**NfsMountPaths** | Pointer to **[]string** | Specifies the path for mounting this Share as an NFS share. If Kerberos Provider has multiple hostaliases, each host alias has its own path. | [optional] [readonly] 
**SmbMountPaths** | Pointer to **[]string** | Specifies the possible paths that can be used to mount this Share as a SMB share. If Active Directory has multiple account names, each machine account has its own path. | [optional] [readonly] 
**S3AccessPath** | Pointer to **NullableString** | Specifies the path to access this Share as an S3 share. | [optional] [readonly] 
**TenantId** | Pointer to **NullableString** | Specifies the tenant id who has access to this Share. | [optional] [readonly] 

## Methods

### NewShare

`func NewShare(name NullableString, viewName NullableString, viewPath NullableString, ) *Share`

NewShare instantiates a new Share object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShareWithDefaults

`func NewShareWithDefaults() *Share`

NewShareWithDefaults instantiates a new Share object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableFilerAuditLogging

`func (o *Share) GetEnableFilerAuditLogging() bool`

GetEnableFilerAuditLogging returns the EnableFilerAuditLogging field if non-nil, zero value otherwise.

### GetEnableFilerAuditLoggingOk

`func (o *Share) GetEnableFilerAuditLoggingOk() (*bool, bool)`

GetEnableFilerAuditLoggingOk returns a tuple with the EnableFilerAuditLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFilerAuditLogging

`func (o *Share) SetEnableFilerAuditLogging(v bool)`

SetEnableFilerAuditLogging sets EnableFilerAuditLogging field to given value.

### HasEnableFilerAuditLogging

`func (o *Share) HasEnableFilerAuditLogging() bool`

HasEnableFilerAuditLogging returns a boolean if a field has been set.

### SetEnableFilerAuditLoggingNil

`func (o *Share) SetEnableFilerAuditLoggingNil(b bool)`

 SetEnableFilerAuditLoggingNil sets the value for EnableFilerAuditLogging to be an explicit nil

### UnsetEnableFilerAuditLogging
`func (o *Share) UnsetEnableFilerAuditLogging()`

UnsetEnableFilerAuditLogging ensures that no value is present for EnableFilerAuditLogging, not even an explicit nil
### GetSmbConfig

`func (o *Share) GetSmbConfig() AliasSmbConfig`

GetSmbConfig returns the SmbConfig field if non-nil, zero value otherwise.

### GetSmbConfigOk

`func (o *Share) GetSmbConfigOk() (*AliasSmbConfig, bool)`

GetSmbConfigOk returns a tuple with the SmbConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbConfig

`func (o *Share) SetSmbConfig(v AliasSmbConfig)`

SetSmbConfig sets SmbConfig field to given value.

### HasSmbConfig

`func (o *Share) HasSmbConfig() bool`

HasSmbConfig returns a boolean if a field has been set.

### GetClientSubnetWhitelist

`func (o *Share) GetClientSubnetWhitelist() []Subnet`

GetClientSubnetWhitelist returns the ClientSubnetWhitelist field if non-nil, zero value otherwise.

### GetClientSubnetWhitelistOk

`func (o *Share) GetClientSubnetWhitelistOk() (*[]Subnet, bool)`

GetClientSubnetWhitelistOk returns a tuple with the ClientSubnetWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSubnetWhitelist

`func (o *Share) SetClientSubnetWhitelist(v []Subnet)`

SetClientSubnetWhitelist sets ClientSubnetWhitelist field to given value.

### HasClientSubnetWhitelist

`func (o *Share) HasClientSubnetWhitelist() bool`

HasClientSubnetWhitelist returns a boolean if a field has been set.

### SetClientSubnetWhitelistNil

`func (o *Share) SetClientSubnetWhitelistNil(b bool)`

 SetClientSubnetWhitelistNil sets the value for ClientSubnetWhitelist to be an explicit nil

### UnsetClientSubnetWhitelist
`func (o *Share) UnsetClientSubnetWhitelist()`

UnsetClientSubnetWhitelist ensures that no value is present for ClientSubnetWhitelist, not even an explicit nil
### GetName

`func (o *Share) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Share) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Share) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *Share) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Share) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
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


### SetViewNameNil

`func (o *Share) SetViewNameNil(b bool)`

 SetViewNameNil sets the value for ViewName to be an explicit nil

### UnsetViewName
`func (o *Share) UnsetViewName()`

UnsetViewName ensures that no value is present for ViewName, not even an explicit nil
### GetViewPath

`func (o *Share) GetViewPath() string`

GetViewPath returns the ViewPath field if non-nil, zero value otherwise.

### GetViewPathOk

`func (o *Share) GetViewPathOk() (*string, bool)`

GetViewPathOk returns a tuple with the ViewPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewPath

`func (o *Share) SetViewPath(v string)`

SetViewPath sets ViewPath field to given value.


### SetViewPathNil

`func (o *Share) SetViewPathNil(b bool)`

 SetViewPathNil sets the value for ViewPath to be an explicit nil

### UnsetViewPath
`func (o *Share) UnsetViewPath()`

UnsetViewPath ensures that no value is present for ViewPath, not even an explicit nil
### GetNfsMountPaths

`func (o *Share) GetNfsMountPaths() []string`

GetNfsMountPaths returns the NfsMountPaths field if non-nil, zero value otherwise.

### GetNfsMountPathsOk

`func (o *Share) GetNfsMountPathsOk() (*[]string, bool)`

GetNfsMountPathsOk returns a tuple with the NfsMountPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsMountPaths

`func (o *Share) SetNfsMountPaths(v []string)`

SetNfsMountPaths sets NfsMountPaths field to given value.

### HasNfsMountPaths

`func (o *Share) HasNfsMountPaths() bool`

HasNfsMountPaths returns a boolean if a field has been set.

### SetNfsMountPathsNil

`func (o *Share) SetNfsMountPathsNil(b bool)`

 SetNfsMountPathsNil sets the value for NfsMountPaths to be an explicit nil

### UnsetNfsMountPaths
`func (o *Share) UnsetNfsMountPaths()`

UnsetNfsMountPaths ensures that no value is present for NfsMountPaths, not even an explicit nil
### GetSmbMountPaths

`func (o *Share) GetSmbMountPaths() []string`

GetSmbMountPaths returns the SmbMountPaths field if non-nil, zero value otherwise.

### GetSmbMountPathsOk

`func (o *Share) GetSmbMountPathsOk() (*[]string, bool)`

GetSmbMountPathsOk returns a tuple with the SmbMountPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbMountPaths

`func (o *Share) SetSmbMountPaths(v []string)`

SetSmbMountPaths sets SmbMountPaths field to given value.

### HasSmbMountPaths

`func (o *Share) HasSmbMountPaths() bool`

HasSmbMountPaths returns a boolean if a field has been set.

### SetSmbMountPathsNil

`func (o *Share) SetSmbMountPathsNil(b bool)`

 SetSmbMountPathsNil sets the value for SmbMountPaths to be an explicit nil

### UnsetSmbMountPaths
`func (o *Share) UnsetSmbMountPaths()`

UnsetSmbMountPaths ensures that no value is present for SmbMountPaths, not even an explicit nil
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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


