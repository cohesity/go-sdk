# NfsConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableNfsKerberosAuthentication** | Pointer to **NullableBool** | If set, it enables NFS Kerberos Authentication | [optional] 
**EnableNfsKerberosIntegrity** | Pointer to **NullableBool** | If set, it enables NFS Kerberos Integrity | [optional] 
**EnableNfsKerberosPrivacy** | Pointer to **NullableBool** | If set, it enables NFS Kerberos Privacy | [optional] 
**EnableNfsUnixAuthentication** | Pointer to **NullableBool** | If set, it enables NFS UNIX Authentication | [optional] 
**EnableNfsViewDiscovery** | Pointer to **NullableBool** | If set, it enables discovery of view for NFS. | [optional] 
**EnableNfsWcc** | Pointer to **NullableBool** | If set, it enables NFS weak cache consistency. | [optional] 
**NfsAllSquash** | Pointer to [**NfsConfigNfsAllSquash**](NfsConfigNfsAllSquash.md) |  | [optional] 
**NfsRootPermissions** | Pointer to [**NfsConfigNfsRootPermissions**](NfsConfigNfsRootPermissions.md) |  | [optional] 
**NfsRootSquash** | Pointer to [**NfsConfigNfsRootSquash**](NfsConfigNfsRootSquash.md) |  | [optional] 

## Methods

### NewNfsConfig

`func NewNfsConfig() *NfsConfig`

NewNfsConfig instantiates a new NfsConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNfsConfigWithDefaults

`func NewNfsConfigWithDefaults() *NfsConfig`

NewNfsConfigWithDefaults instantiates a new NfsConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableNfsKerberosAuthentication

`func (o *NfsConfig) GetEnableNfsKerberosAuthentication() bool`

GetEnableNfsKerberosAuthentication returns the EnableNfsKerberosAuthentication field if non-nil, zero value otherwise.

### GetEnableNfsKerberosAuthenticationOk

`func (o *NfsConfig) GetEnableNfsKerberosAuthenticationOk() (*bool, bool)`

GetEnableNfsKerberosAuthenticationOk returns a tuple with the EnableNfsKerberosAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNfsKerberosAuthentication

`func (o *NfsConfig) SetEnableNfsKerberosAuthentication(v bool)`

SetEnableNfsKerberosAuthentication sets EnableNfsKerberosAuthentication field to given value.

### HasEnableNfsKerberosAuthentication

`func (o *NfsConfig) HasEnableNfsKerberosAuthentication() bool`

HasEnableNfsKerberosAuthentication returns a boolean if a field has been set.

### SetEnableNfsKerberosAuthenticationNil

`func (o *NfsConfig) SetEnableNfsKerberosAuthenticationNil(b bool)`

 SetEnableNfsKerberosAuthenticationNil sets the value for EnableNfsKerberosAuthentication to be an explicit nil

### UnsetEnableNfsKerberosAuthentication
`func (o *NfsConfig) UnsetEnableNfsKerberosAuthentication()`

UnsetEnableNfsKerberosAuthentication ensures that no value is present for EnableNfsKerberosAuthentication, not even an explicit nil
### GetEnableNfsKerberosIntegrity

`func (o *NfsConfig) GetEnableNfsKerberosIntegrity() bool`

GetEnableNfsKerberosIntegrity returns the EnableNfsKerberosIntegrity field if non-nil, zero value otherwise.

### GetEnableNfsKerberosIntegrityOk

`func (o *NfsConfig) GetEnableNfsKerberosIntegrityOk() (*bool, bool)`

GetEnableNfsKerberosIntegrityOk returns a tuple with the EnableNfsKerberosIntegrity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNfsKerberosIntegrity

`func (o *NfsConfig) SetEnableNfsKerberosIntegrity(v bool)`

SetEnableNfsKerberosIntegrity sets EnableNfsKerberosIntegrity field to given value.

### HasEnableNfsKerberosIntegrity

`func (o *NfsConfig) HasEnableNfsKerberosIntegrity() bool`

HasEnableNfsKerberosIntegrity returns a boolean if a field has been set.

### SetEnableNfsKerberosIntegrityNil

`func (o *NfsConfig) SetEnableNfsKerberosIntegrityNil(b bool)`

 SetEnableNfsKerberosIntegrityNil sets the value for EnableNfsKerberosIntegrity to be an explicit nil

### UnsetEnableNfsKerberosIntegrity
`func (o *NfsConfig) UnsetEnableNfsKerberosIntegrity()`

UnsetEnableNfsKerberosIntegrity ensures that no value is present for EnableNfsKerberosIntegrity, not even an explicit nil
### GetEnableNfsKerberosPrivacy

`func (o *NfsConfig) GetEnableNfsKerberosPrivacy() bool`

GetEnableNfsKerberosPrivacy returns the EnableNfsKerberosPrivacy field if non-nil, zero value otherwise.

### GetEnableNfsKerberosPrivacyOk

`func (o *NfsConfig) GetEnableNfsKerberosPrivacyOk() (*bool, bool)`

GetEnableNfsKerberosPrivacyOk returns a tuple with the EnableNfsKerberosPrivacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNfsKerberosPrivacy

`func (o *NfsConfig) SetEnableNfsKerberosPrivacy(v bool)`

SetEnableNfsKerberosPrivacy sets EnableNfsKerberosPrivacy field to given value.

### HasEnableNfsKerberosPrivacy

`func (o *NfsConfig) HasEnableNfsKerberosPrivacy() bool`

HasEnableNfsKerberosPrivacy returns a boolean if a field has been set.

### SetEnableNfsKerberosPrivacyNil

`func (o *NfsConfig) SetEnableNfsKerberosPrivacyNil(b bool)`

 SetEnableNfsKerberosPrivacyNil sets the value for EnableNfsKerberosPrivacy to be an explicit nil

### UnsetEnableNfsKerberosPrivacy
`func (o *NfsConfig) UnsetEnableNfsKerberosPrivacy()`

UnsetEnableNfsKerberosPrivacy ensures that no value is present for EnableNfsKerberosPrivacy, not even an explicit nil
### GetEnableNfsUnixAuthentication

`func (o *NfsConfig) GetEnableNfsUnixAuthentication() bool`

GetEnableNfsUnixAuthentication returns the EnableNfsUnixAuthentication field if non-nil, zero value otherwise.

### GetEnableNfsUnixAuthenticationOk

`func (o *NfsConfig) GetEnableNfsUnixAuthenticationOk() (*bool, bool)`

GetEnableNfsUnixAuthenticationOk returns a tuple with the EnableNfsUnixAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNfsUnixAuthentication

`func (o *NfsConfig) SetEnableNfsUnixAuthentication(v bool)`

SetEnableNfsUnixAuthentication sets EnableNfsUnixAuthentication field to given value.

### HasEnableNfsUnixAuthentication

`func (o *NfsConfig) HasEnableNfsUnixAuthentication() bool`

HasEnableNfsUnixAuthentication returns a boolean if a field has been set.

### SetEnableNfsUnixAuthenticationNil

`func (o *NfsConfig) SetEnableNfsUnixAuthenticationNil(b bool)`

 SetEnableNfsUnixAuthenticationNil sets the value for EnableNfsUnixAuthentication to be an explicit nil

### UnsetEnableNfsUnixAuthentication
`func (o *NfsConfig) UnsetEnableNfsUnixAuthentication()`

UnsetEnableNfsUnixAuthentication ensures that no value is present for EnableNfsUnixAuthentication, not even an explicit nil
### GetEnableNfsViewDiscovery

`func (o *NfsConfig) GetEnableNfsViewDiscovery() bool`

GetEnableNfsViewDiscovery returns the EnableNfsViewDiscovery field if non-nil, zero value otherwise.

### GetEnableNfsViewDiscoveryOk

`func (o *NfsConfig) GetEnableNfsViewDiscoveryOk() (*bool, bool)`

GetEnableNfsViewDiscoveryOk returns a tuple with the EnableNfsViewDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNfsViewDiscovery

`func (o *NfsConfig) SetEnableNfsViewDiscovery(v bool)`

SetEnableNfsViewDiscovery sets EnableNfsViewDiscovery field to given value.

### HasEnableNfsViewDiscovery

`func (o *NfsConfig) HasEnableNfsViewDiscovery() bool`

HasEnableNfsViewDiscovery returns a boolean if a field has been set.

### SetEnableNfsViewDiscoveryNil

`func (o *NfsConfig) SetEnableNfsViewDiscoveryNil(b bool)`

 SetEnableNfsViewDiscoveryNil sets the value for EnableNfsViewDiscovery to be an explicit nil

### UnsetEnableNfsViewDiscovery
`func (o *NfsConfig) UnsetEnableNfsViewDiscovery()`

UnsetEnableNfsViewDiscovery ensures that no value is present for EnableNfsViewDiscovery, not even an explicit nil
### GetEnableNfsWcc

`func (o *NfsConfig) GetEnableNfsWcc() bool`

GetEnableNfsWcc returns the EnableNfsWcc field if non-nil, zero value otherwise.

### GetEnableNfsWccOk

`func (o *NfsConfig) GetEnableNfsWccOk() (*bool, bool)`

GetEnableNfsWccOk returns a tuple with the EnableNfsWcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNfsWcc

`func (o *NfsConfig) SetEnableNfsWcc(v bool)`

SetEnableNfsWcc sets EnableNfsWcc field to given value.

### HasEnableNfsWcc

`func (o *NfsConfig) HasEnableNfsWcc() bool`

HasEnableNfsWcc returns a boolean if a field has been set.

### SetEnableNfsWccNil

`func (o *NfsConfig) SetEnableNfsWccNil(b bool)`

 SetEnableNfsWccNil sets the value for EnableNfsWcc to be an explicit nil

### UnsetEnableNfsWcc
`func (o *NfsConfig) UnsetEnableNfsWcc()`

UnsetEnableNfsWcc ensures that no value is present for EnableNfsWcc, not even an explicit nil
### GetNfsAllSquash

`func (o *NfsConfig) GetNfsAllSquash() NfsConfigNfsAllSquash`

GetNfsAllSquash returns the NfsAllSquash field if non-nil, zero value otherwise.

### GetNfsAllSquashOk

`func (o *NfsConfig) GetNfsAllSquashOk() (*NfsConfigNfsAllSquash, bool)`

GetNfsAllSquashOk returns a tuple with the NfsAllSquash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsAllSquash

`func (o *NfsConfig) SetNfsAllSquash(v NfsConfigNfsAllSquash)`

SetNfsAllSquash sets NfsAllSquash field to given value.

### HasNfsAllSquash

`func (o *NfsConfig) HasNfsAllSquash() bool`

HasNfsAllSquash returns a boolean if a field has been set.

### GetNfsRootPermissions

`func (o *NfsConfig) GetNfsRootPermissions() NfsConfigNfsRootPermissions`

GetNfsRootPermissions returns the NfsRootPermissions field if non-nil, zero value otherwise.

### GetNfsRootPermissionsOk

`func (o *NfsConfig) GetNfsRootPermissionsOk() (*NfsConfigNfsRootPermissions, bool)`

GetNfsRootPermissionsOk returns a tuple with the NfsRootPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsRootPermissions

`func (o *NfsConfig) SetNfsRootPermissions(v NfsConfigNfsRootPermissions)`

SetNfsRootPermissions sets NfsRootPermissions field to given value.

### HasNfsRootPermissions

`func (o *NfsConfig) HasNfsRootPermissions() bool`

HasNfsRootPermissions returns a boolean if a field has been set.

### GetNfsRootSquash

`func (o *NfsConfig) GetNfsRootSquash() NfsConfigNfsRootSquash`

GetNfsRootSquash returns the NfsRootSquash field if non-nil, zero value otherwise.

### GetNfsRootSquashOk

`func (o *NfsConfig) GetNfsRootSquashOk() (*NfsConfigNfsRootSquash, bool)`

GetNfsRootSquashOk returns a tuple with the NfsRootSquash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsRootSquash

`func (o *NfsConfig) SetNfsRootSquash(v NfsConfigNfsRootSquash)`

SetNfsRootSquash sets NfsRootSquash field to given value.

### HasNfsRootSquash

`func (o *NfsConfig) HasNfsRootSquash() bool`

HasNfsRootSquash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


