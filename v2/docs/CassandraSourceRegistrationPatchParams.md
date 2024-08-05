# CassandraSourceRegistrationPatchParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CassandraCredentials** | Pointer to [**NullableCassandraSourceRegistrationPatchParamsCassandraCredentials**](CassandraSourceRegistrationPatchParamsCassandraCredentials.md) |  | [optional] 
**CommitLogBackupLocation** | Pointer to **NullableString** | Commit Logs backup location on cassandra nodes | [optional] 
**ConfigDirectory** | Pointer to **NullableString** | Directory path containing Cassandra configuration YAML file. | [optional] 
**DataCenterNames** | Pointer to **[]string** | Data centers for this cluster. | [optional] 
**DseConfigurationDirectory** | Pointer to **NullableString** | Directory from where DSE specific configuration can be read. This should be set only when you are using the DSE distribution of Cassandra. | [optional] 
**DseSolrInfo** | Pointer to [**DSESolrInfo**](DSESolrInfo.md) |  | [optional] 
**IsDseAuthenticator** | Pointer to **NullableBool** | Set to true if this cluster has DSE Authenticator. | [optional] 
**IsDseTieredStorage** | Pointer to **NullableBool** | Set to true if this cluster has DSE tiered storage. | [optional] 
**JmxCredentials** | Pointer to [**NullableCassandraSourceRegistrationPatchParamsJmxCredentials**](CassandraSourceRegistrationPatchParamsJmxCredentials.md) |  | [optional] 
**KerberosPrincipal** | Pointer to **NullableString** | Principal for the kerberos connection. (This is required only if your Cassandra has Kerberos authentication. Please refer to the user guide.) | [optional] 
**SeedNode** | Pointer to **NullableString** | Any one seed node of the Cassandra cluster. | [optional] 
**SshPasswordCredentials** | Pointer to [**NullableSshPasswordCredentials**](SshPasswordCredentials.md) |  | [optional] 
**SshPrivateKeyCredentials** | Pointer to [**NullableSshPrivateKeyCredentials**](SshPrivateKeyCredentials.md) |  | [optional] 

## Methods

### NewCassandraSourceRegistrationPatchParams

`func NewCassandraSourceRegistrationPatchParams() *CassandraSourceRegistrationPatchParams`

NewCassandraSourceRegistrationPatchParams instantiates a new CassandraSourceRegistrationPatchParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCassandraSourceRegistrationPatchParamsWithDefaults

`func NewCassandraSourceRegistrationPatchParamsWithDefaults() *CassandraSourceRegistrationPatchParams`

NewCassandraSourceRegistrationPatchParamsWithDefaults instantiates a new CassandraSourceRegistrationPatchParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCassandraCredentials

`func (o *CassandraSourceRegistrationPatchParams) GetCassandraCredentials() CassandraSourceRegistrationPatchParamsCassandraCredentials`

GetCassandraCredentials returns the CassandraCredentials field if non-nil, zero value otherwise.

### GetCassandraCredentialsOk

`func (o *CassandraSourceRegistrationPatchParams) GetCassandraCredentialsOk() (*CassandraSourceRegistrationPatchParamsCassandraCredentials, bool)`

GetCassandraCredentialsOk returns a tuple with the CassandraCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraCredentials

`func (o *CassandraSourceRegistrationPatchParams) SetCassandraCredentials(v CassandraSourceRegistrationPatchParamsCassandraCredentials)`

SetCassandraCredentials sets CassandraCredentials field to given value.

### HasCassandraCredentials

`func (o *CassandraSourceRegistrationPatchParams) HasCassandraCredentials() bool`

HasCassandraCredentials returns a boolean if a field has been set.

### SetCassandraCredentialsNil

`func (o *CassandraSourceRegistrationPatchParams) SetCassandraCredentialsNil(b bool)`

 SetCassandraCredentialsNil sets the value for CassandraCredentials to be an explicit nil

### UnsetCassandraCredentials
`func (o *CassandraSourceRegistrationPatchParams) UnsetCassandraCredentials()`

UnsetCassandraCredentials ensures that no value is present for CassandraCredentials, not even an explicit nil
### GetCommitLogBackupLocation

`func (o *CassandraSourceRegistrationPatchParams) GetCommitLogBackupLocation() string`

GetCommitLogBackupLocation returns the CommitLogBackupLocation field if non-nil, zero value otherwise.

### GetCommitLogBackupLocationOk

`func (o *CassandraSourceRegistrationPatchParams) GetCommitLogBackupLocationOk() (*string, bool)`

GetCommitLogBackupLocationOk returns a tuple with the CommitLogBackupLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitLogBackupLocation

`func (o *CassandraSourceRegistrationPatchParams) SetCommitLogBackupLocation(v string)`

SetCommitLogBackupLocation sets CommitLogBackupLocation field to given value.

### HasCommitLogBackupLocation

`func (o *CassandraSourceRegistrationPatchParams) HasCommitLogBackupLocation() bool`

HasCommitLogBackupLocation returns a boolean if a field has been set.

### SetCommitLogBackupLocationNil

`func (o *CassandraSourceRegistrationPatchParams) SetCommitLogBackupLocationNil(b bool)`

 SetCommitLogBackupLocationNil sets the value for CommitLogBackupLocation to be an explicit nil

### UnsetCommitLogBackupLocation
`func (o *CassandraSourceRegistrationPatchParams) UnsetCommitLogBackupLocation()`

UnsetCommitLogBackupLocation ensures that no value is present for CommitLogBackupLocation, not even an explicit nil
### GetConfigDirectory

`func (o *CassandraSourceRegistrationPatchParams) GetConfigDirectory() string`

GetConfigDirectory returns the ConfigDirectory field if non-nil, zero value otherwise.

### GetConfigDirectoryOk

`func (o *CassandraSourceRegistrationPatchParams) GetConfigDirectoryOk() (*string, bool)`

GetConfigDirectoryOk returns a tuple with the ConfigDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigDirectory

`func (o *CassandraSourceRegistrationPatchParams) SetConfigDirectory(v string)`

SetConfigDirectory sets ConfigDirectory field to given value.

### HasConfigDirectory

`func (o *CassandraSourceRegistrationPatchParams) HasConfigDirectory() bool`

HasConfigDirectory returns a boolean if a field has been set.

### SetConfigDirectoryNil

`func (o *CassandraSourceRegistrationPatchParams) SetConfigDirectoryNil(b bool)`

 SetConfigDirectoryNil sets the value for ConfigDirectory to be an explicit nil

### UnsetConfigDirectory
`func (o *CassandraSourceRegistrationPatchParams) UnsetConfigDirectory()`

UnsetConfigDirectory ensures that no value is present for ConfigDirectory, not even an explicit nil
### GetDataCenterNames

`func (o *CassandraSourceRegistrationPatchParams) GetDataCenterNames() []string`

GetDataCenterNames returns the DataCenterNames field if non-nil, zero value otherwise.

### GetDataCenterNamesOk

`func (o *CassandraSourceRegistrationPatchParams) GetDataCenterNamesOk() (*[]string, bool)`

GetDataCenterNamesOk returns a tuple with the DataCenterNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenterNames

`func (o *CassandraSourceRegistrationPatchParams) SetDataCenterNames(v []string)`

SetDataCenterNames sets DataCenterNames field to given value.

### HasDataCenterNames

`func (o *CassandraSourceRegistrationPatchParams) HasDataCenterNames() bool`

HasDataCenterNames returns a boolean if a field has been set.

### GetDseConfigurationDirectory

`func (o *CassandraSourceRegistrationPatchParams) GetDseConfigurationDirectory() string`

GetDseConfigurationDirectory returns the DseConfigurationDirectory field if non-nil, zero value otherwise.

### GetDseConfigurationDirectoryOk

`func (o *CassandraSourceRegistrationPatchParams) GetDseConfigurationDirectoryOk() (*string, bool)`

GetDseConfigurationDirectoryOk returns a tuple with the DseConfigurationDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDseConfigurationDirectory

`func (o *CassandraSourceRegistrationPatchParams) SetDseConfigurationDirectory(v string)`

SetDseConfigurationDirectory sets DseConfigurationDirectory field to given value.

### HasDseConfigurationDirectory

`func (o *CassandraSourceRegistrationPatchParams) HasDseConfigurationDirectory() bool`

HasDseConfigurationDirectory returns a boolean if a field has been set.

### SetDseConfigurationDirectoryNil

`func (o *CassandraSourceRegistrationPatchParams) SetDseConfigurationDirectoryNil(b bool)`

 SetDseConfigurationDirectoryNil sets the value for DseConfigurationDirectory to be an explicit nil

### UnsetDseConfigurationDirectory
`func (o *CassandraSourceRegistrationPatchParams) UnsetDseConfigurationDirectory()`

UnsetDseConfigurationDirectory ensures that no value is present for DseConfigurationDirectory, not even an explicit nil
### GetDseSolrInfo

`func (o *CassandraSourceRegistrationPatchParams) GetDseSolrInfo() DSESolrInfo`

GetDseSolrInfo returns the DseSolrInfo field if non-nil, zero value otherwise.

### GetDseSolrInfoOk

`func (o *CassandraSourceRegistrationPatchParams) GetDseSolrInfoOk() (*DSESolrInfo, bool)`

GetDseSolrInfoOk returns a tuple with the DseSolrInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDseSolrInfo

`func (o *CassandraSourceRegistrationPatchParams) SetDseSolrInfo(v DSESolrInfo)`

SetDseSolrInfo sets DseSolrInfo field to given value.

### HasDseSolrInfo

`func (o *CassandraSourceRegistrationPatchParams) HasDseSolrInfo() bool`

HasDseSolrInfo returns a boolean if a field has been set.

### GetIsDseAuthenticator

`func (o *CassandraSourceRegistrationPatchParams) GetIsDseAuthenticator() bool`

GetIsDseAuthenticator returns the IsDseAuthenticator field if non-nil, zero value otherwise.

### GetIsDseAuthenticatorOk

`func (o *CassandraSourceRegistrationPatchParams) GetIsDseAuthenticatorOk() (*bool, bool)`

GetIsDseAuthenticatorOk returns a tuple with the IsDseAuthenticator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDseAuthenticator

`func (o *CassandraSourceRegistrationPatchParams) SetIsDseAuthenticator(v bool)`

SetIsDseAuthenticator sets IsDseAuthenticator field to given value.

### HasIsDseAuthenticator

`func (o *CassandraSourceRegistrationPatchParams) HasIsDseAuthenticator() bool`

HasIsDseAuthenticator returns a boolean if a field has been set.

### SetIsDseAuthenticatorNil

`func (o *CassandraSourceRegistrationPatchParams) SetIsDseAuthenticatorNil(b bool)`

 SetIsDseAuthenticatorNil sets the value for IsDseAuthenticator to be an explicit nil

### UnsetIsDseAuthenticator
`func (o *CassandraSourceRegistrationPatchParams) UnsetIsDseAuthenticator()`

UnsetIsDseAuthenticator ensures that no value is present for IsDseAuthenticator, not even an explicit nil
### GetIsDseTieredStorage

`func (o *CassandraSourceRegistrationPatchParams) GetIsDseTieredStorage() bool`

GetIsDseTieredStorage returns the IsDseTieredStorage field if non-nil, zero value otherwise.

### GetIsDseTieredStorageOk

`func (o *CassandraSourceRegistrationPatchParams) GetIsDseTieredStorageOk() (*bool, bool)`

GetIsDseTieredStorageOk returns a tuple with the IsDseTieredStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDseTieredStorage

`func (o *CassandraSourceRegistrationPatchParams) SetIsDseTieredStorage(v bool)`

SetIsDseTieredStorage sets IsDseTieredStorage field to given value.

### HasIsDseTieredStorage

`func (o *CassandraSourceRegistrationPatchParams) HasIsDseTieredStorage() bool`

HasIsDseTieredStorage returns a boolean if a field has been set.

### SetIsDseTieredStorageNil

`func (o *CassandraSourceRegistrationPatchParams) SetIsDseTieredStorageNil(b bool)`

 SetIsDseTieredStorageNil sets the value for IsDseTieredStorage to be an explicit nil

### UnsetIsDseTieredStorage
`func (o *CassandraSourceRegistrationPatchParams) UnsetIsDseTieredStorage()`

UnsetIsDseTieredStorage ensures that no value is present for IsDseTieredStorage, not even an explicit nil
### GetJmxCredentials

`func (o *CassandraSourceRegistrationPatchParams) GetJmxCredentials() CassandraSourceRegistrationPatchParamsJmxCredentials`

GetJmxCredentials returns the JmxCredentials field if non-nil, zero value otherwise.

### GetJmxCredentialsOk

`func (o *CassandraSourceRegistrationPatchParams) GetJmxCredentialsOk() (*CassandraSourceRegistrationPatchParamsJmxCredentials, bool)`

GetJmxCredentialsOk returns a tuple with the JmxCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJmxCredentials

`func (o *CassandraSourceRegistrationPatchParams) SetJmxCredentials(v CassandraSourceRegistrationPatchParamsJmxCredentials)`

SetJmxCredentials sets JmxCredentials field to given value.

### HasJmxCredentials

`func (o *CassandraSourceRegistrationPatchParams) HasJmxCredentials() bool`

HasJmxCredentials returns a boolean if a field has been set.

### SetJmxCredentialsNil

`func (o *CassandraSourceRegistrationPatchParams) SetJmxCredentialsNil(b bool)`

 SetJmxCredentialsNil sets the value for JmxCredentials to be an explicit nil

### UnsetJmxCredentials
`func (o *CassandraSourceRegistrationPatchParams) UnsetJmxCredentials()`

UnsetJmxCredentials ensures that no value is present for JmxCredentials, not even an explicit nil
### GetKerberosPrincipal

`func (o *CassandraSourceRegistrationPatchParams) GetKerberosPrincipal() string`

GetKerberosPrincipal returns the KerberosPrincipal field if non-nil, zero value otherwise.

### GetKerberosPrincipalOk

`func (o *CassandraSourceRegistrationPatchParams) GetKerberosPrincipalOk() (*string, bool)`

GetKerberosPrincipalOk returns a tuple with the KerberosPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosPrincipal

`func (o *CassandraSourceRegistrationPatchParams) SetKerberosPrincipal(v string)`

SetKerberosPrincipal sets KerberosPrincipal field to given value.

### HasKerberosPrincipal

`func (o *CassandraSourceRegistrationPatchParams) HasKerberosPrincipal() bool`

HasKerberosPrincipal returns a boolean if a field has been set.

### SetKerberosPrincipalNil

`func (o *CassandraSourceRegistrationPatchParams) SetKerberosPrincipalNil(b bool)`

 SetKerberosPrincipalNil sets the value for KerberosPrincipal to be an explicit nil

### UnsetKerberosPrincipal
`func (o *CassandraSourceRegistrationPatchParams) UnsetKerberosPrincipal()`

UnsetKerberosPrincipal ensures that no value is present for KerberosPrincipal, not even an explicit nil
### GetSeedNode

`func (o *CassandraSourceRegistrationPatchParams) GetSeedNode() string`

GetSeedNode returns the SeedNode field if non-nil, zero value otherwise.

### GetSeedNodeOk

`func (o *CassandraSourceRegistrationPatchParams) GetSeedNodeOk() (*string, bool)`

GetSeedNodeOk returns a tuple with the SeedNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeedNode

`func (o *CassandraSourceRegistrationPatchParams) SetSeedNode(v string)`

SetSeedNode sets SeedNode field to given value.

### HasSeedNode

`func (o *CassandraSourceRegistrationPatchParams) HasSeedNode() bool`

HasSeedNode returns a boolean if a field has been set.

### SetSeedNodeNil

`func (o *CassandraSourceRegistrationPatchParams) SetSeedNodeNil(b bool)`

 SetSeedNodeNil sets the value for SeedNode to be an explicit nil

### UnsetSeedNode
`func (o *CassandraSourceRegistrationPatchParams) UnsetSeedNode()`

UnsetSeedNode ensures that no value is present for SeedNode, not even an explicit nil
### GetSshPasswordCredentials

`func (o *CassandraSourceRegistrationPatchParams) GetSshPasswordCredentials() SshPasswordCredentials`

GetSshPasswordCredentials returns the SshPasswordCredentials field if non-nil, zero value otherwise.

### GetSshPasswordCredentialsOk

`func (o *CassandraSourceRegistrationPatchParams) GetSshPasswordCredentialsOk() (*SshPasswordCredentials, bool)`

GetSshPasswordCredentialsOk returns a tuple with the SshPasswordCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPasswordCredentials

`func (o *CassandraSourceRegistrationPatchParams) SetSshPasswordCredentials(v SshPasswordCredentials)`

SetSshPasswordCredentials sets SshPasswordCredentials field to given value.

### HasSshPasswordCredentials

`func (o *CassandraSourceRegistrationPatchParams) HasSshPasswordCredentials() bool`

HasSshPasswordCredentials returns a boolean if a field has been set.

### SetSshPasswordCredentialsNil

`func (o *CassandraSourceRegistrationPatchParams) SetSshPasswordCredentialsNil(b bool)`

 SetSshPasswordCredentialsNil sets the value for SshPasswordCredentials to be an explicit nil

### UnsetSshPasswordCredentials
`func (o *CassandraSourceRegistrationPatchParams) UnsetSshPasswordCredentials()`

UnsetSshPasswordCredentials ensures that no value is present for SshPasswordCredentials, not even an explicit nil
### GetSshPrivateKeyCredentials

`func (o *CassandraSourceRegistrationPatchParams) GetSshPrivateKeyCredentials() SshPrivateKeyCredentials`

GetSshPrivateKeyCredentials returns the SshPrivateKeyCredentials field if non-nil, zero value otherwise.

### GetSshPrivateKeyCredentialsOk

`func (o *CassandraSourceRegistrationPatchParams) GetSshPrivateKeyCredentialsOk() (*SshPrivateKeyCredentials, bool)`

GetSshPrivateKeyCredentialsOk returns a tuple with the SshPrivateKeyCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPrivateKeyCredentials

`func (o *CassandraSourceRegistrationPatchParams) SetSshPrivateKeyCredentials(v SshPrivateKeyCredentials)`

SetSshPrivateKeyCredentials sets SshPrivateKeyCredentials field to given value.

### HasSshPrivateKeyCredentials

`func (o *CassandraSourceRegistrationPatchParams) HasSshPrivateKeyCredentials() bool`

HasSshPrivateKeyCredentials returns a boolean if a field has been set.

### SetSshPrivateKeyCredentialsNil

`func (o *CassandraSourceRegistrationPatchParams) SetSshPrivateKeyCredentialsNil(b bool)`

 SetSshPrivateKeyCredentialsNil sets the value for SshPrivateKeyCredentials to be an explicit nil

### UnsetSshPrivateKeyCredentials
`func (o *CassandraSourceRegistrationPatchParams) UnsetSshPrivateKeyCredentials()`

UnsetSshPrivateKeyCredentials ensures that no value is present for SshPrivateKeyCredentials, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


