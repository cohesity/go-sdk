# ShareAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** | Specifies the Share name. | 
**ViewName** | **NullableString** | Specifies the View name of this Share. | 
**ViewPath** | **NullableString** | Specifies the View path of this Share. | 
**NfsMountPaths** | Pointer to **[]string** | Specifies the path for mounting this Share as an NFS share. If Kerberos Provider has multiple hostaliases, each host alias has its own path. | [optional] [readonly] 
**SmbMountPaths** | Pointer to **[]string** | Specifies the possible paths that can be used to mount this Share as a SMB share. If Active Directory has multiple account names, each machine account has its own path. | [optional] [readonly] 
**S3AccessPath** | Pointer to **NullableString** | Specifies the path to access this Share as an S3 share. | [optional] [readonly] 
**TenantId** | Pointer to **NullableString** | Specifies the tenant id who has access to this Share. | [optional] [readonly] 

## Methods

### NewShareAllOf

`func NewShareAllOf(name NullableString, viewName NullableString, viewPath NullableString, ) *ShareAllOf`

NewShareAllOf instantiates a new ShareAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShareAllOfWithDefaults

`func NewShareAllOfWithDefaults() *ShareAllOf`

NewShareAllOfWithDefaults instantiates a new ShareAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ShareAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShareAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShareAllOf) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *ShareAllOf) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ShareAllOf) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetViewName

`func (o *ShareAllOf) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *ShareAllOf) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *ShareAllOf) SetViewName(v string)`

SetViewName sets ViewName field to given value.


### SetViewNameNil

`func (o *ShareAllOf) SetViewNameNil(b bool)`

 SetViewNameNil sets the value for ViewName to be an explicit nil

### UnsetViewName
`func (o *ShareAllOf) UnsetViewName()`

UnsetViewName ensures that no value is present for ViewName, not even an explicit nil
### GetViewPath

`func (o *ShareAllOf) GetViewPath() string`

GetViewPath returns the ViewPath field if non-nil, zero value otherwise.

### GetViewPathOk

`func (o *ShareAllOf) GetViewPathOk() (*string, bool)`

GetViewPathOk returns a tuple with the ViewPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewPath

`func (o *ShareAllOf) SetViewPath(v string)`

SetViewPath sets ViewPath field to given value.


### SetViewPathNil

`func (o *ShareAllOf) SetViewPathNil(b bool)`

 SetViewPathNil sets the value for ViewPath to be an explicit nil

### UnsetViewPath
`func (o *ShareAllOf) UnsetViewPath()`

UnsetViewPath ensures that no value is present for ViewPath, not even an explicit nil
### GetNfsMountPaths

`func (o *ShareAllOf) GetNfsMountPaths() []string`

GetNfsMountPaths returns the NfsMountPaths field if non-nil, zero value otherwise.

### GetNfsMountPathsOk

`func (o *ShareAllOf) GetNfsMountPathsOk() (*[]string, bool)`

GetNfsMountPathsOk returns a tuple with the NfsMountPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsMountPaths

`func (o *ShareAllOf) SetNfsMountPaths(v []string)`

SetNfsMountPaths sets NfsMountPaths field to given value.

### HasNfsMountPaths

`func (o *ShareAllOf) HasNfsMountPaths() bool`

HasNfsMountPaths returns a boolean if a field has been set.

### SetNfsMountPathsNil

`func (o *ShareAllOf) SetNfsMountPathsNil(b bool)`

 SetNfsMountPathsNil sets the value for NfsMountPaths to be an explicit nil

### UnsetNfsMountPaths
`func (o *ShareAllOf) UnsetNfsMountPaths()`

UnsetNfsMountPaths ensures that no value is present for NfsMountPaths, not even an explicit nil
### GetSmbMountPaths

`func (o *ShareAllOf) GetSmbMountPaths() []string`

GetSmbMountPaths returns the SmbMountPaths field if non-nil, zero value otherwise.

### GetSmbMountPathsOk

`func (o *ShareAllOf) GetSmbMountPathsOk() (*[]string, bool)`

GetSmbMountPathsOk returns a tuple with the SmbMountPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbMountPaths

`func (o *ShareAllOf) SetSmbMountPaths(v []string)`

SetSmbMountPaths sets SmbMountPaths field to given value.

### HasSmbMountPaths

`func (o *ShareAllOf) HasSmbMountPaths() bool`

HasSmbMountPaths returns a boolean if a field has been set.

### SetSmbMountPathsNil

`func (o *ShareAllOf) SetSmbMountPathsNil(b bool)`

 SetSmbMountPathsNil sets the value for SmbMountPaths to be an explicit nil

### UnsetSmbMountPaths
`func (o *ShareAllOf) UnsetSmbMountPaths()`

UnsetSmbMountPaths ensures that no value is present for SmbMountPaths, not even an explicit nil
### GetS3AccessPath

`func (o *ShareAllOf) GetS3AccessPath() string`

GetS3AccessPath returns the S3AccessPath field if non-nil, zero value otherwise.

### GetS3AccessPathOk

`func (o *ShareAllOf) GetS3AccessPathOk() (*string, bool)`

GetS3AccessPathOk returns a tuple with the S3AccessPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3AccessPath

`func (o *ShareAllOf) SetS3AccessPath(v string)`

SetS3AccessPath sets S3AccessPath field to given value.

### HasS3AccessPath

`func (o *ShareAllOf) HasS3AccessPath() bool`

HasS3AccessPath returns a boolean if a field has been set.

### SetS3AccessPathNil

`func (o *ShareAllOf) SetS3AccessPathNil(b bool)`

 SetS3AccessPathNil sets the value for S3AccessPath to be an explicit nil

### UnsetS3AccessPath
`func (o *ShareAllOf) UnsetS3AccessPath()`

UnsetS3AccessPath ensures that no value is present for S3AccessPath, not even an explicit nil
### GetTenantId

`func (o *ShareAllOf) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ShareAllOf) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ShareAllOf) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ShareAllOf) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ShareAllOf) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ShareAllOf) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


