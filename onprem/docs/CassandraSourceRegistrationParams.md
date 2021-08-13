# CassandraSourceRegistrationParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SeedNode** | **string** | Any one seed node of the Cassandra cluster. | 
**ConfigDirectory** | **string** | Directory path containing Cassandra configuration YAML file. | 
**DseConfigurationDirectory** | Pointer to **NullableString** | Directory from where DSE specific configuration can be read. This should be set only when you are using the DSE distribution of Cassandra. | [optional] 
**IsDseTieredStorage** | **bool** | Set to true if this cluster has DSE tiered storage. | 
**IsDseAuthenticator** | **bool** | Set to true if this cluster has DSE Authenticator. | 
**SshPasswordCredentials** | Pointer to [**NullableCassandraConnectionParamsSshPasswordCredentials**](CassandraConnectionParamsSshPasswordCredentials.md) |  | [optional] 
**SshPrivateKeyCredentials** | Pointer to [**NullableCassandraConnectionParamsSshPrivateKeyCredentials**](CassandraConnectionParamsSshPrivateKeyCredentials.md) |  | [optional] 
**JmxCredentials** | Pointer to [**NullableCassandraSourceRegistrationParamsAllOfJmxCredentials**](CassandraSourceRegistrationParamsAllOfJmxCredentials.md) |  | [optional] 
**CassandraCredentials** | Pointer to [**NullableCassandraSourceRegistrationParamsAllOfCassandraCredentials**](CassandraSourceRegistrationParamsAllOfCassandraCredentials.md) |  | [optional] 
**DataCenterNames** | Pointer to **[]string** | Data centers for this cluster. | [optional] 
**KerberosPrincipal** | Pointer to **NullableString** | Principal for the kerberos connection. (This is required only if your Cassandra has Kerberos authentication. Please refer to the user guide.) | [optional] 
**DseSolrInfo** | Pointer to [**DSESolrInfo**](DSESolrInfo.md) |  | [optional] 

## Methods

### NewCassandraSourceRegistrationParams

`func NewCassandraSourceRegistrationParams(seedNode string, configDirectory string, isDseTieredStorage bool, isDseAuthenticator bool, ) *CassandraSourceRegistrationParams`

NewCassandraSourceRegistrationParams instantiates a new CassandraSourceRegistrationParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCassandraSourceRegistrationParamsWithDefaults

`func NewCassandraSourceRegistrationParamsWithDefaults() *CassandraSourceRegistrationParams`

NewCassandraSourceRegistrationParamsWithDefaults instantiates a new CassandraSourceRegistrationParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeedNode

`func (o *CassandraSourceRegistrationParams) GetSeedNode() string`

GetSeedNode returns the SeedNode field if non-nil, zero value otherwise.

### GetSeedNodeOk

`func (o *CassandraSourceRegistrationParams) GetSeedNodeOk() (*string, bool)`

GetSeedNodeOk returns a tuple with the SeedNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeedNode

`func (o *CassandraSourceRegistrationParams) SetSeedNode(v string)`

SetSeedNode sets SeedNode field to given value.


### GetConfigDirectory

`func (o *CassandraSourceRegistrationParams) GetConfigDirectory() string`

GetConfigDirectory returns the ConfigDirectory field if non-nil, zero value otherwise.

### GetConfigDirectoryOk

`func (o *CassandraSourceRegistrationParams) GetConfigDirectoryOk() (*string, bool)`

GetConfigDirectoryOk returns a tuple with the ConfigDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigDirectory

`func (o *CassandraSourceRegistrationParams) SetConfigDirectory(v string)`

SetConfigDirectory sets ConfigDirectory field to given value.


### GetDseConfigurationDirectory

`func (o *CassandraSourceRegistrationParams) GetDseConfigurationDirectory() string`

GetDseConfigurationDirectory returns the DseConfigurationDirectory field if non-nil, zero value otherwise.

### GetDseConfigurationDirectoryOk

`func (o *CassandraSourceRegistrationParams) GetDseConfigurationDirectoryOk() (*string, bool)`

GetDseConfigurationDirectoryOk returns a tuple with the DseConfigurationDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDseConfigurationDirectory

`func (o *CassandraSourceRegistrationParams) SetDseConfigurationDirectory(v string)`

SetDseConfigurationDirectory sets DseConfigurationDirectory field to given value.

### HasDseConfigurationDirectory

`func (o *CassandraSourceRegistrationParams) HasDseConfigurationDirectory() bool`

HasDseConfigurationDirectory returns a boolean if a field has been set.

### SetDseConfigurationDirectoryNil

`func (o *CassandraSourceRegistrationParams) SetDseConfigurationDirectoryNil(b bool)`

 SetDseConfigurationDirectoryNil sets the value for DseConfigurationDirectory to be an explicit nil

### UnsetDseConfigurationDirectory
`func (o *CassandraSourceRegistrationParams) UnsetDseConfigurationDirectory()`

UnsetDseConfigurationDirectory ensures that no value is present for DseConfigurationDirectory, not even an explicit nil
### GetIsDseTieredStorage

`func (o *CassandraSourceRegistrationParams) GetIsDseTieredStorage() bool`

GetIsDseTieredStorage returns the IsDseTieredStorage field if non-nil, zero value otherwise.

### GetIsDseTieredStorageOk

`func (o *CassandraSourceRegistrationParams) GetIsDseTieredStorageOk() (*bool, bool)`

GetIsDseTieredStorageOk returns a tuple with the IsDseTieredStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDseTieredStorage

`func (o *CassandraSourceRegistrationParams) SetIsDseTieredStorage(v bool)`

SetIsDseTieredStorage sets IsDseTieredStorage field to given value.


### GetIsDseAuthenticator

`func (o *CassandraSourceRegistrationParams) GetIsDseAuthenticator() bool`

GetIsDseAuthenticator returns the IsDseAuthenticator field if non-nil, zero value otherwise.

### GetIsDseAuthenticatorOk

`func (o *CassandraSourceRegistrationParams) GetIsDseAuthenticatorOk() (*bool, bool)`

GetIsDseAuthenticatorOk returns a tuple with the IsDseAuthenticator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDseAuthenticator

`func (o *CassandraSourceRegistrationParams) SetIsDseAuthenticator(v bool)`

SetIsDseAuthenticator sets IsDseAuthenticator field to given value.


### GetSshPasswordCredentials

`func (o *CassandraSourceRegistrationParams) GetSshPasswordCredentials() CassandraConnectionParamsSshPasswordCredentials`

GetSshPasswordCredentials returns the SshPasswordCredentials field if non-nil, zero value otherwise.

### GetSshPasswordCredentialsOk

`func (o *CassandraSourceRegistrationParams) GetSshPasswordCredentialsOk() (*CassandraConnectionParamsSshPasswordCredentials, bool)`

GetSshPasswordCredentialsOk returns a tuple with the SshPasswordCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPasswordCredentials

`func (o *CassandraSourceRegistrationParams) SetSshPasswordCredentials(v CassandraConnectionParamsSshPasswordCredentials)`

SetSshPasswordCredentials sets SshPasswordCredentials field to given value.

### HasSshPasswordCredentials

`func (o *CassandraSourceRegistrationParams) HasSshPasswordCredentials() bool`

HasSshPasswordCredentials returns a boolean if a field has been set.

### SetSshPasswordCredentialsNil

`func (o *CassandraSourceRegistrationParams) SetSshPasswordCredentialsNil(b bool)`

 SetSshPasswordCredentialsNil sets the value for SshPasswordCredentials to be an explicit nil

### UnsetSshPasswordCredentials
`func (o *CassandraSourceRegistrationParams) UnsetSshPasswordCredentials()`

UnsetSshPasswordCredentials ensures that no value is present for SshPasswordCredentials, not even an explicit nil
### GetSshPrivateKeyCredentials

`func (o *CassandraSourceRegistrationParams) GetSshPrivateKeyCredentials() CassandraConnectionParamsSshPrivateKeyCredentials`

GetSshPrivateKeyCredentials returns the SshPrivateKeyCredentials field if non-nil, zero value otherwise.

### GetSshPrivateKeyCredentialsOk

`func (o *CassandraSourceRegistrationParams) GetSshPrivateKeyCredentialsOk() (*CassandraConnectionParamsSshPrivateKeyCredentials, bool)`

GetSshPrivateKeyCredentialsOk returns a tuple with the SshPrivateKeyCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPrivateKeyCredentials

`func (o *CassandraSourceRegistrationParams) SetSshPrivateKeyCredentials(v CassandraConnectionParamsSshPrivateKeyCredentials)`

SetSshPrivateKeyCredentials sets SshPrivateKeyCredentials field to given value.

### HasSshPrivateKeyCredentials

`func (o *CassandraSourceRegistrationParams) HasSshPrivateKeyCredentials() bool`

HasSshPrivateKeyCredentials returns a boolean if a field has been set.

### SetSshPrivateKeyCredentialsNil

`func (o *CassandraSourceRegistrationParams) SetSshPrivateKeyCredentialsNil(b bool)`

 SetSshPrivateKeyCredentialsNil sets the value for SshPrivateKeyCredentials to be an explicit nil

### UnsetSshPrivateKeyCredentials
`func (o *CassandraSourceRegistrationParams) UnsetSshPrivateKeyCredentials()`

UnsetSshPrivateKeyCredentials ensures that no value is present for SshPrivateKeyCredentials, not even an explicit nil
### GetJmxCredentials

`func (o *CassandraSourceRegistrationParams) GetJmxCredentials() CassandraSourceRegistrationParamsAllOfJmxCredentials`

GetJmxCredentials returns the JmxCredentials field if non-nil, zero value otherwise.

### GetJmxCredentialsOk

`func (o *CassandraSourceRegistrationParams) GetJmxCredentialsOk() (*CassandraSourceRegistrationParamsAllOfJmxCredentials, bool)`

GetJmxCredentialsOk returns a tuple with the JmxCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJmxCredentials

`func (o *CassandraSourceRegistrationParams) SetJmxCredentials(v CassandraSourceRegistrationParamsAllOfJmxCredentials)`

SetJmxCredentials sets JmxCredentials field to given value.

### HasJmxCredentials

`func (o *CassandraSourceRegistrationParams) HasJmxCredentials() bool`

HasJmxCredentials returns a boolean if a field has been set.

### SetJmxCredentialsNil

`func (o *CassandraSourceRegistrationParams) SetJmxCredentialsNil(b bool)`

 SetJmxCredentialsNil sets the value for JmxCredentials to be an explicit nil

### UnsetJmxCredentials
`func (o *CassandraSourceRegistrationParams) UnsetJmxCredentials()`

UnsetJmxCredentials ensures that no value is present for JmxCredentials, not even an explicit nil
### GetCassandraCredentials

`func (o *CassandraSourceRegistrationParams) GetCassandraCredentials() CassandraSourceRegistrationParamsAllOfCassandraCredentials`

GetCassandraCredentials returns the CassandraCredentials field if non-nil, zero value otherwise.

### GetCassandraCredentialsOk

`func (o *CassandraSourceRegistrationParams) GetCassandraCredentialsOk() (*CassandraSourceRegistrationParamsAllOfCassandraCredentials, bool)`

GetCassandraCredentialsOk returns a tuple with the CassandraCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraCredentials

`func (o *CassandraSourceRegistrationParams) SetCassandraCredentials(v CassandraSourceRegistrationParamsAllOfCassandraCredentials)`

SetCassandraCredentials sets CassandraCredentials field to given value.

### HasCassandraCredentials

`func (o *CassandraSourceRegistrationParams) HasCassandraCredentials() bool`

HasCassandraCredentials returns a boolean if a field has been set.

### SetCassandraCredentialsNil

`func (o *CassandraSourceRegistrationParams) SetCassandraCredentialsNil(b bool)`

 SetCassandraCredentialsNil sets the value for CassandraCredentials to be an explicit nil

### UnsetCassandraCredentials
`func (o *CassandraSourceRegistrationParams) UnsetCassandraCredentials()`

UnsetCassandraCredentials ensures that no value is present for CassandraCredentials, not even an explicit nil
### GetDataCenterNames

`func (o *CassandraSourceRegistrationParams) GetDataCenterNames() []string`

GetDataCenterNames returns the DataCenterNames field if non-nil, zero value otherwise.

### GetDataCenterNamesOk

`func (o *CassandraSourceRegistrationParams) GetDataCenterNamesOk() (*[]string, bool)`

GetDataCenterNamesOk returns a tuple with the DataCenterNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenterNames

`func (o *CassandraSourceRegistrationParams) SetDataCenterNames(v []string)`

SetDataCenterNames sets DataCenterNames field to given value.

### HasDataCenterNames

`func (o *CassandraSourceRegistrationParams) HasDataCenterNames() bool`

HasDataCenterNames returns a boolean if a field has been set.

### GetKerberosPrincipal

`func (o *CassandraSourceRegistrationParams) GetKerberosPrincipal() string`

GetKerberosPrincipal returns the KerberosPrincipal field if non-nil, zero value otherwise.

### GetKerberosPrincipalOk

`func (o *CassandraSourceRegistrationParams) GetKerberosPrincipalOk() (*string, bool)`

GetKerberosPrincipalOk returns a tuple with the KerberosPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosPrincipal

`func (o *CassandraSourceRegistrationParams) SetKerberosPrincipal(v string)`

SetKerberosPrincipal sets KerberosPrincipal field to given value.

### HasKerberosPrincipal

`func (o *CassandraSourceRegistrationParams) HasKerberosPrincipal() bool`

HasKerberosPrincipal returns a boolean if a field has been set.

### SetKerberosPrincipalNil

`func (o *CassandraSourceRegistrationParams) SetKerberosPrincipalNil(b bool)`

 SetKerberosPrincipalNil sets the value for KerberosPrincipal to be an explicit nil

### UnsetKerberosPrincipal
`func (o *CassandraSourceRegistrationParams) UnsetKerberosPrincipal()`

UnsetKerberosPrincipal ensures that no value is present for KerberosPrincipal, not even an explicit nil
### GetDseSolrInfo

`func (o *CassandraSourceRegistrationParams) GetDseSolrInfo() DSESolrInfo`

GetDseSolrInfo returns the DseSolrInfo field if non-nil, zero value otherwise.

### GetDseSolrInfoOk

`func (o *CassandraSourceRegistrationParams) GetDseSolrInfoOk() (*DSESolrInfo, bool)`

GetDseSolrInfoOk returns a tuple with the DseSolrInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDseSolrInfo

`func (o *CassandraSourceRegistrationParams) SetDseSolrInfo(v DSESolrInfo)`

SetDseSolrInfo sets DseSolrInfo field to given value.

### HasDseSolrInfo

`func (o *CassandraSourceRegistrationParams) HasDseSolrInfo() bool`

HasDseSolrInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


