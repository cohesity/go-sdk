# HdfsSourceRegistrationUpdateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **NullableString** | IP or hostname of any host from which the HDFS configuration files core-site.xml and hdfs-site.xml can be read. | [optional] 
**ConfigurationDirectory** | Pointer to **NullableString** | The directory containing the core-site.xml and hdfs-site.xml configuration files. | [optional] 
**SshPasswordCredentials** | Pointer to [**NullableHbaseSourceRegistrationParamsAllOfSshPasswordCredentials**](HbaseSourceRegistrationParamsAllOfSshPasswordCredentials.md) |  | [optional] 
**SshPrivateKeyCredentials** | Pointer to [**NullableHbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials**](HbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials.md) |  | [optional] 
**KerberosPrincipal** | Pointer to **NullableString** | The kerberos principal to be used to connect to this HDFS source. | [optional] 

## Methods

### NewHdfsSourceRegistrationUpdateParams

`func NewHdfsSourceRegistrationUpdateParams() *HdfsSourceRegistrationUpdateParams`

NewHdfsSourceRegistrationUpdateParams instantiates a new HdfsSourceRegistrationUpdateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHdfsSourceRegistrationUpdateParamsWithDefaults

`func NewHdfsSourceRegistrationUpdateParamsWithDefaults() *HdfsSourceRegistrationUpdateParams`

NewHdfsSourceRegistrationUpdateParamsWithDefaults instantiates a new HdfsSourceRegistrationUpdateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *HdfsSourceRegistrationUpdateParams) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *HdfsSourceRegistrationUpdateParams) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *HdfsSourceRegistrationUpdateParams) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *HdfsSourceRegistrationUpdateParams) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *HdfsSourceRegistrationUpdateParams) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *HdfsSourceRegistrationUpdateParams) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetConfigurationDirectory

`func (o *HdfsSourceRegistrationUpdateParams) GetConfigurationDirectory() string`

GetConfigurationDirectory returns the ConfigurationDirectory field if non-nil, zero value otherwise.

### GetConfigurationDirectoryOk

`func (o *HdfsSourceRegistrationUpdateParams) GetConfigurationDirectoryOk() (*string, bool)`

GetConfigurationDirectoryOk returns a tuple with the ConfigurationDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationDirectory

`func (o *HdfsSourceRegistrationUpdateParams) SetConfigurationDirectory(v string)`

SetConfigurationDirectory sets ConfigurationDirectory field to given value.

### HasConfigurationDirectory

`func (o *HdfsSourceRegistrationUpdateParams) HasConfigurationDirectory() bool`

HasConfigurationDirectory returns a boolean if a field has been set.

### SetConfigurationDirectoryNil

`func (o *HdfsSourceRegistrationUpdateParams) SetConfigurationDirectoryNil(b bool)`

 SetConfigurationDirectoryNil sets the value for ConfigurationDirectory to be an explicit nil

### UnsetConfigurationDirectory
`func (o *HdfsSourceRegistrationUpdateParams) UnsetConfigurationDirectory()`

UnsetConfigurationDirectory ensures that no value is present for ConfigurationDirectory, not even an explicit nil
### GetSshPasswordCredentials

`func (o *HdfsSourceRegistrationUpdateParams) GetSshPasswordCredentials() HbaseSourceRegistrationParamsAllOfSshPasswordCredentials`

GetSshPasswordCredentials returns the SshPasswordCredentials field if non-nil, zero value otherwise.

### GetSshPasswordCredentialsOk

`func (o *HdfsSourceRegistrationUpdateParams) GetSshPasswordCredentialsOk() (*HbaseSourceRegistrationParamsAllOfSshPasswordCredentials, bool)`

GetSshPasswordCredentialsOk returns a tuple with the SshPasswordCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPasswordCredentials

`func (o *HdfsSourceRegistrationUpdateParams) SetSshPasswordCredentials(v HbaseSourceRegistrationParamsAllOfSshPasswordCredentials)`

SetSshPasswordCredentials sets SshPasswordCredentials field to given value.

### HasSshPasswordCredentials

`func (o *HdfsSourceRegistrationUpdateParams) HasSshPasswordCredentials() bool`

HasSshPasswordCredentials returns a boolean if a field has been set.

### SetSshPasswordCredentialsNil

`func (o *HdfsSourceRegistrationUpdateParams) SetSshPasswordCredentialsNil(b bool)`

 SetSshPasswordCredentialsNil sets the value for SshPasswordCredentials to be an explicit nil

### UnsetSshPasswordCredentials
`func (o *HdfsSourceRegistrationUpdateParams) UnsetSshPasswordCredentials()`

UnsetSshPasswordCredentials ensures that no value is present for SshPasswordCredentials, not even an explicit nil
### GetSshPrivateKeyCredentials

`func (o *HdfsSourceRegistrationUpdateParams) GetSshPrivateKeyCredentials() HbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials`

GetSshPrivateKeyCredentials returns the SshPrivateKeyCredentials field if non-nil, zero value otherwise.

### GetSshPrivateKeyCredentialsOk

`func (o *HdfsSourceRegistrationUpdateParams) GetSshPrivateKeyCredentialsOk() (*HbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials, bool)`

GetSshPrivateKeyCredentialsOk returns a tuple with the SshPrivateKeyCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPrivateKeyCredentials

`func (o *HdfsSourceRegistrationUpdateParams) SetSshPrivateKeyCredentials(v HbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials)`

SetSshPrivateKeyCredentials sets SshPrivateKeyCredentials field to given value.

### HasSshPrivateKeyCredentials

`func (o *HdfsSourceRegistrationUpdateParams) HasSshPrivateKeyCredentials() bool`

HasSshPrivateKeyCredentials returns a boolean if a field has been set.

### SetSshPrivateKeyCredentialsNil

`func (o *HdfsSourceRegistrationUpdateParams) SetSshPrivateKeyCredentialsNil(b bool)`

 SetSshPrivateKeyCredentialsNil sets the value for SshPrivateKeyCredentials to be an explicit nil

### UnsetSshPrivateKeyCredentials
`func (o *HdfsSourceRegistrationUpdateParams) UnsetSshPrivateKeyCredentials()`

UnsetSshPrivateKeyCredentials ensures that no value is present for SshPrivateKeyCredentials, not even an explicit nil
### GetKerberosPrincipal

`func (o *HdfsSourceRegistrationUpdateParams) GetKerberosPrincipal() string`

GetKerberosPrincipal returns the KerberosPrincipal field if non-nil, zero value otherwise.

### GetKerberosPrincipalOk

`func (o *HdfsSourceRegistrationUpdateParams) GetKerberosPrincipalOk() (*string, bool)`

GetKerberosPrincipalOk returns a tuple with the KerberosPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosPrincipal

`func (o *HdfsSourceRegistrationUpdateParams) SetKerberosPrincipal(v string)`

SetKerberosPrincipal sets KerberosPrincipal field to given value.

### HasKerberosPrincipal

`func (o *HdfsSourceRegistrationUpdateParams) HasKerberosPrincipal() bool`

HasKerberosPrincipal returns a boolean if a field has been set.

### SetKerberosPrincipalNil

`func (o *HdfsSourceRegistrationUpdateParams) SetKerberosPrincipalNil(b bool)`

 SetKerberosPrincipalNil sets the value for KerberosPrincipal to be an explicit nil

### UnsetKerberosPrincipal
`func (o *HdfsSourceRegistrationUpdateParams) UnsetKerberosPrincipal()`

UnsetKerberosPrincipal ensures that no value is present for KerberosPrincipal, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


