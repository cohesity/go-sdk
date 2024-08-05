# CassandraConnectionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigDirectory** | **string** | Directory path containing Cassandra configuration YAML file. | 
**DseConfigurationDirectory** | Pointer to **NullableString** | Directory from where DSE specific configuration can be read. This should be set only when you are using the DSE distribution of Cassandra. | [optional] 
**IsDseAuthenticator** | **bool** | Set to true if this cluster has DSE Authenticator. | 
**IsDseTieredStorage** | **bool** | Set to true if this cluster has DSE tiered storage. | 
**SeedNode** | **string** | Any one seed node of the Cassandra cluster. | 
**SshPasswordCredentials** | Pointer to [**NullableSshPasswordCredentials**](SshPasswordCredentials.md) |  | [optional] 
**SshPrivateKeyCredentials** | Pointer to [**NullableSshPrivateKeyCredentials**](SshPrivateKeyCredentials.md) |  | [optional] 

## Methods

### NewCassandraConnectionParams

`func NewCassandraConnectionParams(configDirectory string, isDseAuthenticator bool, isDseTieredStorage bool, seedNode string, ) *CassandraConnectionParams`

NewCassandraConnectionParams instantiates a new CassandraConnectionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCassandraConnectionParamsWithDefaults

`func NewCassandraConnectionParamsWithDefaults() *CassandraConnectionParams`

NewCassandraConnectionParamsWithDefaults instantiates a new CassandraConnectionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigDirectory

`func (o *CassandraConnectionParams) GetConfigDirectory() string`

GetConfigDirectory returns the ConfigDirectory field if non-nil, zero value otherwise.

### GetConfigDirectoryOk

`func (o *CassandraConnectionParams) GetConfigDirectoryOk() (*string, bool)`

GetConfigDirectoryOk returns a tuple with the ConfigDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigDirectory

`func (o *CassandraConnectionParams) SetConfigDirectory(v string)`

SetConfigDirectory sets ConfigDirectory field to given value.


### GetDseConfigurationDirectory

`func (o *CassandraConnectionParams) GetDseConfigurationDirectory() string`

GetDseConfigurationDirectory returns the DseConfigurationDirectory field if non-nil, zero value otherwise.

### GetDseConfigurationDirectoryOk

`func (o *CassandraConnectionParams) GetDseConfigurationDirectoryOk() (*string, bool)`

GetDseConfigurationDirectoryOk returns a tuple with the DseConfigurationDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDseConfigurationDirectory

`func (o *CassandraConnectionParams) SetDseConfigurationDirectory(v string)`

SetDseConfigurationDirectory sets DseConfigurationDirectory field to given value.

### HasDseConfigurationDirectory

`func (o *CassandraConnectionParams) HasDseConfigurationDirectory() bool`

HasDseConfigurationDirectory returns a boolean if a field has been set.

### SetDseConfigurationDirectoryNil

`func (o *CassandraConnectionParams) SetDseConfigurationDirectoryNil(b bool)`

 SetDseConfigurationDirectoryNil sets the value for DseConfigurationDirectory to be an explicit nil

### UnsetDseConfigurationDirectory
`func (o *CassandraConnectionParams) UnsetDseConfigurationDirectory()`

UnsetDseConfigurationDirectory ensures that no value is present for DseConfigurationDirectory, not even an explicit nil
### GetIsDseAuthenticator

`func (o *CassandraConnectionParams) GetIsDseAuthenticator() bool`

GetIsDseAuthenticator returns the IsDseAuthenticator field if non-nil, zero value otherwise.

### GetIsDseAuthenticatorOk

`func (o *CassandraConnectionParams) GetIsDseAuthenticatorOk() (*bool, bool)`

GetIsDseAuthenticatorOk returns a tuple with the IsDseAuthenticator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDseAuthenticator

`func (o *CassandraConnectionParams) SetIsDseAuthenticator(v bool)`

SetIsDseAuthenticator sets IsDseAuthenticator field to given value.


### GetIsDseTieredStorage

`func (o *CassandraConnectionParams) GetIsDseTieredStorage() bool`

GetIsDseTieredStorage returns the IsDseTieredStorage field if non-nil, zero value otherwise.

### GetIsDseTieredStorageOk

`func (o *CassandraConnectionParams) GetIsDseTieredStorageOk() (*bool, bool)`

GetIsDseTieredStorageOk returns a tuple with the IsDseTieredStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDseTieredStorage

`func (o *CassandraConnectionParams) SetIsDseTieredStorage(v bool)`

SetIsDseTieredStorage sets IsDseTieredStorage field to given value.


### GetSeedNode

`func (o *CassandraConnectionParams) GetSeedNode() string`

GetSeedNode returns the SeedNode field if non-nil, zero value otherwise.

### GetSeedNodeOk

`func (o *CassandraConnectionParams) GetSeedNodeOk() (*string, bool)`

GetSeedNodeOk returns a tuple with the SeedNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeedNode

`func (o *CassandraConnectionParams) SetSeedNode(v string)`

SetSeedNode sets SeedNode field to given value.


### GetSshPasswordCredentials

`func (o *CassandraConnectionParams) GetSshPasswordCredentials() SshPasswordCredentials`

GetSshPasswordCredentials returns the SshPasswordCredentials field if non-nil, zero value otherwise.

### GetSshPasswordCredentialsOk

`func (o *CassandraConnectionParams) GetSshPasswordCredentialsOk() (*SshPasswordCredentials, bool)`

GetSshPasswordCredentialsOk returns a tuple with the SshPasswordCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPasswordCredentials

`func (o *CassandraConnectionParams) SetSshPasswordCredentials(v SshPasswordCredentials)`

SetSshPasswordCredentials sets SshPasswordCredentials field to given value.

### HasSshPasswordCredentials

`func (o *CassandraConnectionParams) HasSshPasswordCredentials() bool`

HasSshPasswordCredentials returns a boolean if a field has been set.

### SetSshPasswordCredentialsNil

`func (o *CassandraConnectionParams) SetSshPasswordCredentialsNil(b bool)`

 SetSshPasswordCredentialsNil sets the value for SshPasswordCredentials to be an explicit nil

### UnsetSshPasswordCredentials
`func (o *CassandraConnectionParams) UnsetSshPasswordCredentials()`

UnsetSshPasswordCredentials ensures that no value is present for SshPasswordCredentials, not even an explicit nil
### GetSshPrivateKeyCredentials

`func (o *CassandraConnectionParams) GetSshPrivateKeyCredentials() SshPrivateKeyCredentials`

GetSshPrivateKeyCredentials returns the SshPrivateKeyCredentials field if non-nil, zero value otherwise.

### GetSshPrivateKeyCredentialsOk

`func (o *CassandraConnectionParams) GetSshPrivateKeyCredentialsOk() (*SshPrivateKeyCredentials, bool)`

GetSshPrivateKeyCredentialsOk returns a tuple with the SshPrivateKeyCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPrivateKeyCredentials

`func (o *CassandraConnectionParams) SetSshPrivateKeyCredentials(v SshPrivateKeyCredentials)`

SetSshPrivateKeyCredentials sets SshPrivateKeyCredentials field to given value.

### HasSshPrivateKeyCredentials

`func (o *CassandraConnectionParams) HasSshPrivateKeyCredentials() bool`

HasSshPrivateKeyCredentials returns a boolean if a field has been set.

### SetSshPrivateKeyCredentialsNil

`func (o *CassandraConnectionParams) SetSshPrivateKeyCredentialsNil(b bool)`

 SetSshPrivateKeyCredentialsNil sets the value for SshPrivateKeyCredentials to be an explicit nil

### UnsetSshPrivateKeyCredentials
`func (o *CassandraConnectionParams) UnsetSshPrivateKeyCredentials()`

UnsetSshPrivateKeyCredentials ensures that no value is present for SshPrivateKeyCredentials, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


