# HiveSourceRegistrationUpdateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigurationDirectory** | Pointer to **NullableString** | The directory containing the hive-site.xml. | [optional] 
**Host** | Pointer to **NullableString** | IP or hostname of any host from which the Hive configuration file hive-site.xml can be read. | [optional] 
**KerberosPrincipal** | Pointer to **NullableString** | The kerberos principal to be used to connect to this Hive source. | [optional] 
**SshPasswordCredentials** | Pointer to [**NullableHbaseSourceRegistrationParamsAllOfSshPasswordCredentials**](HbaseSourceRegistrationParamsAllOfSshPasswordCredentials.md) |  | [optional] 
**SshPrivateKeyCredentials** | Pointer to [**NullableHbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials**](HbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials.md) |  | [optional] 

## Methods

### NewHiveSourceRegistrationUpdateParams

`func NewHiveSourceRegistrationUpdateParams() *HiveSourceRegistrationUpdateParams`

NewHiveSourceRegistrationUpdateParams instantiates a new HiveSourceRegistrationUpdateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHiveSourceRegistrationUpdateParamsWithDefaults

`func NewHiveSourceRegistrationUpdateParamsWithDefaults() *HiveSourceRegistrationUpdateParams`

NewHiveSourceRegistrationUpdateParamsWithDefaults instantiates a new HiveSourceRegistrationUpdateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurationDirectory

`func (o *HiveSourceRegistrationUpdateParams) GetConfigurationDirectory() string`

GetConfigurationDirectory returns the ConfigurationDirectory field if non-nil, zero value otherwise.

### GetConfigurationDirectoryOk

`func (o *HiveSourceRegistrationUpdateParams) GetConfigurationDirectoryOk() (*string, bool)`

GetConfigurationDirectoryOk returns a tuple with the ConfigurationDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationDirectory

`func (o *HiveSourceRegistrationUpdateParams) SetConfigurationDirectory(v string)`

SetConfigurationDirectory sets ConfigurationDirectory field to given value.

### HasConfigurationDirectory

`func (o *HiveSourceRegistrationUpdateParams) HasConfigurationDirectory() bool`

HasConfigurationDirectory returns a boolean if a field has been set.

### SetConfigurationDirectoryNil

`func (o *HiveSourceRegistrationUpdateParams) SetConfigurationDirectoryNil(b bool)`

 SetConfigurationDirectoryNil sets the value for ConfigurationDirectory to be an explicit nil

### UnsetConfigurationDirectory
`func (o *HiveSourceRegistrationUpdateParams) UnsetConfigurationDirectory()`

UnsetConfigurationDirectory ensures that no value is present for ConfigurationDirectory, not even an explicit nil
### GetHost

`func (o *HiveSourceRegistrationUpdateParams) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *HiveSourceRegistrationUpdateParams) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *HiveSourceRegistrationUpdateParams) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *HiveSourceRegistrationUpdateParams) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *HiveSourceRegistrationUpdateParams) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *HiveSourceRegistrationUpdateParams) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetKerberosPrincipal

`func (o *HiveSourceRegistrationUpdateParams) GetKerberosPrincipal() string`

GetKerberosPrincipal returns the KerberosPrincipal field if non-nil, zero value otherwise.

### GetKerberosPrincipalOk

`func (o *HiveSourceRegistrationUpdateParams) GetKerberosPrincipalOk() (*string, bool)`

GetKerberosPrincipalOk returns a tuple with the KerberosPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosPrincipal

`func (o *HiveSourceRegistrationUpdateParams) SetKerberosPrincipal(v string)`

SetKerberosPrincipal sets KerberosPrincipal field to given value.

### HasKerberosPrincipal

`func (o *HiveSourceRegistrationUpdateParams) HasKerberosPrincipal() bool`

HasKerberosPrincipal returns a boolean if a field has been set.

### SetKerberosPrincipalNil

`func (o *HiveSourceRegistrationUpdateParams) SetKerberosPrincipalNil(b bool)`

 SetKerberosPrincipalNil sets the value for KerberosPrincipal to be an explicit nil

### UnsetKerberosPrincipal
`func (o *HiveSourceRegistrationUpdateParams) UnsetKerberosPrincipal()`

UnsetKerberosPrincipal ensures that no value is present for KerberosPrincipal, not even an explicit nil
### GetSshPasswordCredentials

`func (o *HiveSourceRegistrationUpdateParams) GetSshPasswordCredentials() HbaseSourceRegistrationParamsAllOfSshPasswordCredentials`

GetSshPasswordCredentials returns the SshPasswordCredentials field if non-nil, zero value otherwise.

### GetSshPasswordCredentialsOk

`func (o *HiveSourceRegistrationUpdateParams) GetSshPasswordCredentialsOk() (*HbaseSourceRegistrationParamsAllOfSshPasswordCredentials, bool)`

GetSshPasswordCredentialsOk returns a tuple with the SshPasswordCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPasswordCredentials

`func (o *HiveSourceRegistrationUpdateParams) SetSshPasswordCredentials(v HbaseSourceRegistrationParamsAllOfSshPasswordCredentials)`

SetSshPasswordCredentials sets SshPasswordCredentials field to given value.

### HasSshPasswordCredentials

`func (o *HiveSourceRegistrationUpdateParams) HasSshPasswordCredentials() bool`

HasSshPasswordCredentials returns a boolean if a field has been set.

### SetSshPasswordCredentialsNil

`func (o *HiveSourceRegistrationUpdateParams) SetSshPasswordCredentialsNil(b bool)`

 SetSshPasswordCredentialsNil sets the value for SshPasswordCredentials to be an explicit nil

### UnsetSshPasswordCredentials
`func (o *HiveSourceRegistrationUpdateParams) UnsetSshPasswordCredentials()`

UnsetSshPasswordCredentials ensures that no value is present for SshPasswordCredentials, not even an explicit nil
### GetSshPrivateKeyCredentials

`func (o *HiveSourceRegistrationUpdateParams) GetSshPrivateKeyCredentials() HbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials`

GetSshPrivateKeyCredentials returns the SshPrivateKeyCredentials field if non-nil, zero value otherwise.

### GetSshPrivateKeyCredentialsOk

`func (o *HiveSourceRegistrationUpdateParams) GetSshPrivateKeyCredentialsOk() (*HbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials, bool)`

GetSshPrivateKeyCredentialsOk returns a tuple with the SshPrivateKeyCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPrivateKeyCredentials

`func (o *HiveSourceRegistrationUpdateParams) SetSshPrivateKeyCredentials(v HbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials)`

SetSshPrivateKeyCredentials sets SshPrivateKeyCredentials field to given value.

### HasSshPrivateKeyCredentials

`func (o *HiveSourceRegistrationUpdateParams) HasSshPrivateKeyCredentials() bool`

HasSshPrivateKeyCredentials returns a boolean if a field has been set.

### SetSshPrivateKeyCredentialsNil

`func (o *HiveSourceRegistrationUpdateParams) SetSshPrivateKeyCredentialsNil(b bool)`

 SetSshPrivateKeyCredentialsNil sets the value for SshPrivateKeyCredentials to be an explicit nil

### UnsetSshPrivateKeyCredentials
`func (o *HiveSourceRegistrationUpdateParams) UnsetSshPrivateKeyCredentials()`

UnsetSshPrivateKeyCredentials ensures that no value is present for SshPrivateKeyCredentials, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


