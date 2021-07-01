# NasCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **NullableString** | Specifies the hostname or IP address of the NAS server. | [optional] 
**MountPath** | Pointer to **NullableString** | Specifies the mount path to the NAS server. | [optional] 
**Password** | Pointer to **NullableString** | If using the CIFS protocol and a Username was specified, specify the password for the username. | [optional] 
**ShareType** | Pointer to **NullableString** | Specifies the sharing protocol type used to mount the file system. Currently, only NFS is supported. &#39;kNFS&#39; indicates use the NFS protocol to mount the file system. &#39;kCIFS&#39; indicates use the CIFS protocol to mount the file system. | [optional] 
**Username** | Pointer to **NullableString** | If using the CIFS protocol, you can optional specify a username to use when mounting. | [optional] 

## Methods

### NewNasCredentials

`func NewNasCredentials() *NasCredentials`

NewNasCredentials instantiates a new NasCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNasCredentialsWithDefaults

`func NewNasCredentialsWithDefaults() *NasCredentials`

NewNasCredentialsWithDefaults instantiates a new NasCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *NasCredentials) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *NasCredentials) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *NasCredentials) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *NasCredentials) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *NasCredentials) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *NasCredentials) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetMountPath

`func (o *NasCredentials) GetMountPath() string`

GetMountPath returns the MountPath field if non-nil, zero value otherwise.

### GetMountPathOk

`func (o *NasCredentials) GetMountPathOk() (*string, bool)`

GetMountPathOk returns a tuple with the MountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPath

`func (o *NasCredentials) SetMountPath(v string)`

SetMountPath sets MountPath field to given value.

### HasMountPath

`func (o *NasCredentials) HasMountPath() bool`

HasMountPath returns a boolean if a field has been set.

### SetMountPathNil

`func (o *NasCredentials) SetMountPathNil(b bool)`

 SetMountPathNil sets the value for MountPath to be an explicit nil

### UnsetMountPath
`func (o *NasCredentials) UnsetMountPath()`

UnsetMountPath ensures that no value is present for MountPath, not even an explicit nil
### GetPassword

`func (o *NasCredentials) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *NasCredentials) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *NasCredentials) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *NasCredentials) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *NasCredentials) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *NasCredentials) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetShareType

`func (o *NasCredentials) GetShareType() string`

GetShareType returns the ShareType field if non-nil, zero value otherwise.

### GetShareTypeOk

`func (o *NasCredentials) GetShareTypeOk() (*string, bool)`

GetShareTypeOk returns a tuple with the ShareType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareType

`func (o *NasCredentials) SetShareType(v string)`

SetShareType sets ShareType field to given value.

### HasShareType

`func (o *NasCredentials) HasShareType() bool`

HasShareType returns a boolean if a field has been set.

### SetShareTypeNil

`func (o *NasCredentials) SetShareTypeNil(b bool)`

 SetShareTypeNil sets the value for ShareType to be an explicit nil

### UnsetShareType
`func (o *NasCredentials) UnsetShareType()`

UnsetShareType ensures that no value is present for ShareType, not even an explicit nil
### GetUsername

`func (o *NasCredentials) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *NasCredentials) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *NasCredentials) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *NasCredentials) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *NasCredentials) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *NasCredentials) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


