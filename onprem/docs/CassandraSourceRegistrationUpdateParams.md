# CassandraSourceRegistrationUpdateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SeedNode** | Pointer to **string** | Any one seed node of the Cassandra cluster. | [optional] 
**ConfigDirectory** | Pointer to **string** | Directory path containing Cassandra configuration YAML file. | [optional] 
**DseConfigurationDirectory** | Pointer to **NullableString** | Directory from where DSE specific configuration can be read. This should be set only when you are using the DSE distribution of Cassandra. | [optional] 
**IsDseTieredStorage** | Pointer to **bool** | Set to true if this cluster has DSE tiered storage. | [optional] 
**IsDseAuthenticator** | Pointer to **bool** | Set to true if this cluster has DSE Authenticator. | [optional] 
**SshPasswordCredentials** | Pointer to [**NullableCassandraConnectionParamsSshPasswordCredentials**](CassandraConnectionParamsSshPasswordCredentials.md) |  | [optional] 
**SshPrivateKeyCredentials** | Pointer to [**NullableCassandraConnectionParamsSshPrivateKeyCredentials**](CassandraConnectionParamsSshPrivateKeyCredentials.md) |  | [optional] 
**JmxCredentials** | Pointer to [**NullableCassandraSourceRegistrationUpdateParamsJmxCredentials**](CassandraSourceRegistrationUpdateParamsJmxCredentials.md) |  | [optional] 
**CassandraCredentials** | Pointer to [**NullableCassandraSourceRegistrationUpdateParamsCassandraCredentials**](CassandraSourceRegistrationUpdateParamsCassandraCredentials.md) |  | [optional] 
**DataCenterNames** | Pointer to **[]string** | Data centers for this cluster. | [optional] 
**KerberosPrincipal** | Pointer to **NullableString** | Principal for the kerberos connection. (This is required only if your Cassandra has Kerberos authentication. Please refer to the user guide.) | [optional] 
**DseSolrInfo** | Pointer to [**DSESolrInfo**](DSESolrInfo.md) |  | [optional] 

## Methods

### NewCassandraSourceRegistrationUpdateParams

`func NewCassandraSourceRegistrationUpdateParams() *CassandraSourceRegistrationUpdateParams`

NewCassandraSourceRegistrationUpdateParams instantiates a new CassandraSourceRegistrationUpdateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCassandraSourceRegistrationUpdateParamsWithDefaults

`func NewCassandraSourceRegistrationUpdateParamsWithDefaults() *CassandraSourceRegistrationUpdateParams`

NewCassandraSourceRegistrationUpdateParamsWithDefaults instantiates a new CassandraSourceRegistrationUpdateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeedNode

`func (o *CassandraSourceRegistrationUpdateParams) GetSeedNode() string`

GetSeedNode returns the SeedNode field if non-nil, zero value otherwise.

### GetSeedNodeOk

`func (o *CassandraSourceRegistrationUpdateParams) GetSeedNodeOk() (*string, bool)`

GetSeedNodeOk returns a tuple with the SeedNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeedNode

`func (o *CassandraSourceRegistrationUpdateParams) SetSeedNode(v string)`

SetSeedNode sets SeedNode field to given value.

### HasSeedNode

`func (o *CassandraSourceRegistrationUpdateParams) HasSeedNode() bool`

HasSeedNode returns a boolean if a field has been set.

### GetConfigDirectory

`func (o *CassandraSourceRegistrationUpdateParams) GetConfigDirectory() string`

GetConfigDirectory returns the ConfigDirectory field if non-nil, zero value otherwise.

### GetConfigDirectoryOk

`func (o *CassandraSourceRegistrationUpdateParams) GetConfigDirectoryOk() (*string, bool)`

GetConfigDirectoryOk returns a tuple with the ConfigDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigDirectory

`func (o *CassandraSourceRegistrationUpdateParams) SetConfigDirectory(v string)`

SetConfigDirectory sets ConfigDirectory field to given value.

### HasConfigDirectory

`func (o *CassandraSourceRegistrationUpdateParams) HasConfigDirectory() bool`

HasConfigDirectory returns a boolean if a field has been set.

### GetDseConfigurationDirectory

`func (o *CassandraSourceRegistrationUpdateParams) GetDseConfigurationDirectory() string`

GetDseConfigurationDirectory returns the DseConfigurationDirectory field if non-nil, zero value otherwise.

### GetDseConfigurationDirectoryOk

`func (o *CassandraSourceRegistrationUpdateParams) GetDseConfigurationDirectoryOk() (*string, bool)`

GetDseConfigurationDirectoryOk returns a tuple with the DseConfigurationDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDseConfigurationDirectory

`func (o *CassandraSourceRegistrationUpdateParams) SetDseConfigurationDirectory(v string)`

SetDseConfigurationDirectory sets DseConfigurationDirectory field to given value.

### HasDseConfigurationDirectory

`func (o *CassandraSourceRegistrationUpdateParams) HasDseConfigurationDirectory() bool`

HasDseConfigurationDirectory returns a boolean if a field has been set.

### SetDseConfigurationDirectoryNil

`func (o *CassandraSourceRegistrationUpdateParams) SetDseConfigurationDirectoryNil(b bool)`

 SetDseConfigurationDirectoryNil sets the value for DseConfigurationDirectory to be an explicit nil

### UnsetDseConfigurationDirectory
`func (o *CassandraSourceRegistrationUpdateParams) UnsetDseConfigurationDirectory()`

UnsetDseConfigurationDirectory ensures that no value is present for DseConfigurationDirectory, not even an explicit nil
### GetIsDseTieredStorage

`func (o *CassandraSourceRegistrationUpdateParams) GetIsDseTieredStorage() bool`

GetIsDseTieredStorage returns the IsDseTieredStorage field if non-nil, zero value otherwise.

### GetIsDseTieredStorageOk

`func (o *CassandraSourceRegistrationUpdateParams) GetIsDseTieredStorageOk() (*bool, bool)`

GetIsDseTieredStorageOk returns a tuple with the IsDseTieredStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDseTieredStorage

`func (o *CassandraSourceRegistrationUpdateParams) SetIsDseTieredStorage(v bool)`

SetIsDseTieredStorage sets IsDseTieredStorage field to given value.

### HasIsDseTieredStorage

`func (o *CassandraSourceRegistrationUpdateParams) HasIsDseTieredStorage() bool`

HasIsDseTieredStorage returns a boolean if a field has been set.

### GetIsDseAuthenticator

`func (o *CassandraSourceRegistrationUpdateParams) GetIsDseAuthenticator() bool`

GetIsDseAuthenticator returns the IsDseAuthenticator field if non-nil, zero value otherwise.

### GetIsDseAuthenticatorOk

`func (o *CassandraSourceRegistrationUpdateParams) GetIsDseAuthenticatorOk() (*bool, bool)`

GetIsDseAuthenticatorOk returns a tuple with the IsDseAuthenticator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDseAuthenticator

`func (o *CassandraSourceRegistrationUpdateParams) SetIsDseAuthenticator(v bool)`

SetIsDseAuthenticator sets IsDseAuthenticator field to given value.

### HasIsDseAuthenticator

`func (o *CassandraSourceRegistrationUpdateParams) HasIsDseAuthenticator() bool`

HasIsDseAuthenticator returns a boolean if a field has been set.

### GetSshPasswordCredentials

`func (o *CassandraSourceRegistrationUpdateParams) GetSshPasswordCredentials() CassandraConnectionParamsSshPasswordCredentials`

GetSshPasswordCredentials returns the SshPasswordCredentials field if non-nil, zero value otherwise.

### GetSshPasswordCredentialsOk

`func (o *CassandraSourceRegistrationUpdateParams) GetSshPasswordCredentialsOk() (*CassandraConnectionParamsSshPasswordCredentials, bool)`

GetSshPasswordCredentialsOk returns a tuple with the SshPasswordCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPasswordCredentials

`func (o *CassandraSourceRegistrationUpdateParams) SetSshPasswordCredentials(v CassandraConnectionParamsSshPasswordCredentials)`

SetSshPasswordCredentials sets SshPasswordCredentials field to given value.

### HasSshPasswordCredentials

`func (o *CassandraSourceRegistrationUpdateParams) HasSshPasswordCredentials() bool`

HasSshPasswordCredentials returns a boolean if a field has been set.

### SetSshPasswordCredentialsNil

`func (o *CassandraSourceRegistrationUpdateParams) SetSshPasswordCredentialsNil(b bool)`

 SetSshPasswordCredentialsNil sets the value for SshPasswordCredentials to be an explicit nil

### UnsetSshPasswordCredentials
`func (o *CassandraSourceRegistrationUpdateParams) UnsetSshPasswordCredentials()`

UnsetSshPasswordCredentials ensures that no value is present for SshPasswordCredentials, not even an explicit nil
### GetSshPrivateKeyCredentials

`func (o *CassandraSourceRegistrationUpdateParams) GetSshPrivateKeyCredentials() CassandraConnectionParamsSshPrivateKeyCredentials`

GetSshPrivateKeyCredentials returns the SshPrivateKeyCredentials field if non-nil, zero value otherwise.

### GetSshPrivateKeyCredentialsOk

`func (o *CassandraSourceRegistrationUpdateParams) GetSshPrivateKeyCredentialsOk() (*CassandraConnectionParamsSshPrivateKeyCredentials, bool)`

GetSshPrivateKeyCredentialsOk returns a tuple with the SshPrivateKeyCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPrivateKeyCredentials

`func (o *CassandraSourceRegistrationUpdateParams) SetSshPrivateKeyCredentials(v CassandraConnectionParamsSshPrivateKeyCredentials)`

SetSshPrivateKeyCredentials sets SshPrivateKeyCredentials field to given value.

### HasSshPrivateKeyCredentials

`func (o *CassandraSourceRegistrationUpdateParams) HasSshPrivateKeyCredentials() bool`

HasSshPrivateKeyCredentials returns a boolean if a field has been set.

### SetSshPrivateKeyCredentialsNil

`func (o *CassandraSourceRegistrationUpdateParams) SetSshPrivateKeyCredentialsNil(b bool)`

 SetSshPrivateKeyCredentialsNil sets the value for SshPrivateKeyCredentials to be an explicit nil

### UnsetSshPrivateKeyCredentials
`func (o *CassandraSourceRegistrationUpdateParams) UnsetSshPrivateKeyCredentials()`

UnsetSshPrivateKeyCredentials ensures that no value is present for SshPrivateKeyCredentials, not even an explicit nil
### GetJmxCredentials

`func (o *CassandraSourceRegistrationUpdateParams) GetJmxCredentials() CassandraSourceRegistrationUpdateParamsJmxCredentials`

GetJmxCredentials returns the JmxCredentials field if non-nil, zero value otherwise.

### GetJmxCredentialsOk

`func (o *CassandraSourceRegistrationUpdateParams) GetJmxCredentialsOk() (*CassandraSourceRegistrationUpdateParamsJmxCredentials, bool)`

GetJmxCredentialsOk returns a tuple with the JmxCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJmxCredentials

`func (o *CassandraSourceRegistrationUpdateParams) SetJmxCredentials(v CassandraSourceRegistrationUpdateParamsJmxCredentials)`

SetJmxCredentials sets JmxCredentials field to given value.

### HasJmxCredentials

`func (o *CassandraSourceRegistrationUpdateParams) HasJmxCredentials() bool`

HasJmxCredentials returns a boolean if a field has been set.

### SetJmxCredentialsNil

`func (o *CassandraSourceRegistrationUpdateParams) SetJmxCredentialsNil(b bool)`

 SetJmxCredentialsNil sets the value for JmxCredentials to be an explicit nil

### UnsetJmxCredentials
`func (o *CassandraSourceRegistrationUpdateParams) UnsetJmxCredentials()`

UnsetJmxCredentials ensures that no value is present for JmxCredentials, not even an explicit nil
### GetCassandraCredentials

`func (o *CassandraSourceRegistrationUpdateParams) GetCassandraCredentials() CassandraSourceRegistrationUpdateParamsCassandraCredentials`

GetCassandraCredentials returns the CassandraCredentials field if non-nil, zero value otherwise.

### GetCassandraCredentialsOk

`func (o *CassandraSourceRegistrationUpdateParams) GetCassandraCredentialsOk() (*CassandraSourceRegistrationUpdateParamsCassandraCredentials, bool)`

GetCassandraCredentialsOk returns a tuple with the CassandraCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraCredentials

`func (o *CassandraSourceRegistrationUpdateParams) SetCassandraCredentials(v CassandraSourceRegistrationUpdateParamsCassandraCredentials)`

SetCassandraCredentials sets CassandraCredentials field to given value.

### HasCassandraCredentials

`func (o *CassandraSourceRegistrationUpdateParams) HasCassandraCredentials() bool`

HasCassandraCredentials returns a boolean if a field has been set.

### SetCassandraCredentialsNil

`func (o *CassandraSourceRegistrationUpdateParams) SetCassandraCredentialsNil(b bool)`

 SetCassandraCredentialsNil sets the value for CassandraCredentials to be an explicit nil

### UnsetCassandraCredentials
`func (o *CassandraSourceRegistrationUpdateParams) UnsetCassandraCredentials()`

UnsetCassandraCredentials ensures that no value is present for CassandraCredentials, not even an explicit nil
### GetDataCenterNames

`func (o *CassandraSourceRegistrationUpdateParams) GetDataCenterNames() []string`

GetDataCenterNames returns the DataCenterNames field if non-nil, zero value otherwise.

### GetDataCenterNamesOk

`func (o *CassandraSourceRegistrationUpdateParams) GetDataCenterNamesOk() (*[]string, bool)`

GetDataCenterNamesOk returns a tuple with the DataCenterNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenterNames

`func (o *CassandraSourceRegistrationUpdateParams) SetDataCenterNames(v []string)`

SetDataCenterNames sets DataCenterNames field to given value.

### HasDataCenterNames

`func (o *CassandraSourceRegistrationUpdateParams) HasDataCenterNames() bool`

HasDataCenterNames returns a boolean if a field has been set.

### GetKerberosPrincipal

`func (o *CassandraSourceRegistrationUpdateParams) GetKerberosPrincipal() string`

GetKerberosPrincipal returns the KerberosPrincipal field if non-nil, zero value otherwise.

### GetKerberosPrincipalOk

`func (o *CassandraSourceRegistrationUpdateParams) GetKerberosPrincipalOk() (*string, bool)`

GetKerberosPrincipalOk returns a tuple with the KerberosPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosPrincipal

`func (o *CassandraSourceRegistrationUpdateParams) SetKerberosPrincipal(v string)`

SetKerberosPrincipal sets KerberosPrincipal field to given value.

### HasKerberosPrincipal

`func (o *CassandraSourceRegistrationUpdateParams) HasKerberosPrincipal() bool`

HasKerberosPrincipal returns a boolean if a field has been set.

### SetKerberosPrincipalNil

`func (o *CassandraSourceRegistrationUpdateParams) SetKerberosPrincipalNil(b bool)`

 SetKerberosPrincipalNil sets the value for KerberosPrincipal to be an explicit nil

### UnsetKerberosPrincipal
`func (o *CassandraSourceRegistrationUpdateParams) UnsetKerberosPrincipal()`

UnsetKerberosPrincipal ensures that no value is present for KerberosPrincipal, not even an explicit nil
### GetDseSolrInfo

`func (o *CassandraSourceRegistrationUpdateParams) GetDseSolrInfo() DSESolrInfo`

GetDseSolrInfo returns the DseSolrInfo field if non-nil, zero value otherwise.

### GetDseSolrInfoOk

`func (o *CassandraSourceRegistrationUpdateParams) GetDseSolrInfoOk() (*DSESolrInfo, bool)`

GetDseSolrInfoOk returns a tuple with the DseSolrInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDseSolrInfo

`func (o *CassandraSourceRegistrationUpdateParams) SetDseSolrInfo(v DSESolrInfo)`

SetDseSolrInfo sets DseSolrInfo field to given value.

### HasDseSolrInfo

`func (o *CassandraSourceRegistrationUpdateParams) HasDseSolrInfo() bool`

HasDseSolrInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


