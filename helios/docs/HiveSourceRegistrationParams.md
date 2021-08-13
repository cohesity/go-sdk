# HiveSourceRegistrationParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetastoreAddress** | Pointer to **string** | The MetastoreAddress for this Hive. | [optional] [readonly] 
**MetastorePort** | Pointer to **int32** | The MetastorePort for this Hive. | [optional] [readonly] 
**AuthType** | Pointer to **NullableString** | Authentication type. | [optional] [readonly] 
**Host** | **string** | IP or hostname of any host from which the Hive configuration file hive-site.xml can be read. | 
**ConfigurationDirectory** | **string** | The directory containing the hive-site.xml. | 
**SshPasswordCredentials** | Pointer to [**NullableHiveSourceRegistrationParamsAllOfSshPasswordCredentials**](HiveSourceRegistrationParamsAllOfSshPasswordCredentials.md) |  | [optional] 
**SshPrivateKeyCredentials** | Pointer to [**NullableHbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials**](HbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials.md) |  | [optional] 
**HdfsSourceRegistrationID** | **int64** | Protection Source registration id of the HDFS on which this Hive is running. | 
**KerberosPrincipal** | Pointer to **NullableString** | The kerberos principal to be used to connect to this Hive source. | [optional] 

## Methods

### NewHiveSourceRegistrationParams

`func NewHiveSourceRegistrationParams(host string, configurationDirectory string, hdfsSourceRegistrationID int64, ) *HiveSourceRegistrationParams`

NewHiveSourceRegistrationParams instantiates a new HiveSourceRegistrationParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHiveSourceRegistrationParamsWithDefaults

`func NewHiveSourceRegistrationParamsWithDefaults() *HiveSourceRegistrationParams`

NewHiveSourceRegistrationParamsWithDefaults instantiates a new HiveSourceRegistrationParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetastoreAddress

`func (o *HiveSourceRegistrationParams) GetMetastoreAddress() string`

GetMetastoreAddress returns the MetastoreAddress field if non-nil, zero value otherwise.

### GetMetastoreAddressOk

`func (o *HiveSourceRegistrationParams) GetMetastoreAddressOk() (*string, bool)`

GetMetastoreAddressOk returns a tuple with the MetastoreAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetastoreAddress

`func (o *HiveSourceRegistrationParams) SetMetastoreAddress(v string)`

SetMetastoreAddress sets MetastoreAddress field to given value.

### HasMetastoreAddress

`func (o *HiveSourceRegistrationParams) HasMetastoreAddress() bool`

HasMetastoreAddress returns a boolean if a field has been set.

### GetMetastorePort

`func (o *HiveSourceRegistrationParams) GetMetastorePort() int32`

GetMetastorePort returns the MetastorePort field if non-nil, zero value otherwise.

### GetMetastorePortOk

`func (o *HiveSourceRegistrationParams) GetMetastorePortOk() (*int32, bool)`

GetMetastorePortOk returns a tuple with the MetastorePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetastorePort

`func (o *HiveSourceRegistrationParams) SetMetastorePort(v int32)`

SetMetastorePort sets MetastorePort field to given value.

### HasMetastorePort

`func (o *HiveSourceRegistrationParams) HasMetastorePort() bool`

HasMetastorePort returns a boolean if a field has been set.

### GetAuthType

`func (o *HiveSourceRegistrationParams) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *HiveSourceRegistrationParams) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *HiveSourceRegistrationParams) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *HiveSourceRegistrationParams) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### SetAuthTypeNil

`func (o *HiveSourceRegistrationParams) SetAuthTypeNil(b bool)`

 SetAuthTypeNil sets the value for AuthType to be an explicit nil

### UnsetAuthType
`func (o *HiveSourceRegistrationParams) UnsetAuthType()`

UnsetAuthType ensures that no value is present for AuthType, not even an explicit nil
### GetHost

`func (o *HiveSourceRegistrationParams) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *HiveSourceRegistrationParams) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *HiveSourceRegistrationParams) SetHost(v string)`

SetHost sets Host field to given value.


### GetConfigurationDirectory

`func (o *HiveSourceRegistrationParams) GetConfigurationDirectory() string`

GetConfigurationDirectory returns the ConfigurationDirectory field if non-nil, zero value otherwise.

### GetConfigurationDirectoryOk

`func (o *HiveSourceRegistrationParams) GetConfigurationDirectoryOk() (*string, bool)`

GetConfigurationDirectoryOk returns a tuple with the ConfigurationDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationDirectory

`func (o *HiveSourceRegistrationParams) SetConfigurationDirectory(v string)`

SetConfigurationDirectory sets ConfigurationDirectory field to given value.


### GetSshPasswordCredentials

`func (o *HiveSourceRegistrationParams) GetSshPasswordCredentials() HiveSourceRegistrationParamsAllOfSshPasswordCredentials`

GetSshPasswordCredentials returns the SshPasswordCredentials field if non-nil, zero value otherwise.

### GetSshPasswordCredentialsOk

`func (o *HiveSourceRegistrationParams) GetSshPasswordCredentialsOk() (*HiveSourceRegistrationParamsAllOfSshPasswordCredentials, bool)`

GetSshPasswordCredentialsOk returns a tuple with the SshPasswordCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPasswordCredentials

`func (o *HiveSourceRegistrationParams) SetSshPasswordCredentials(v HiveSourceRegistrationParamsAllOfSshPasswordCredentials)`

SetSshPasswordCredentials sets SshPasswordCredentials field to given value.

### HasSshPasswordCredentials

`func (o *HiveSourceRegistrationParams) HasSshPasswordCredentials() bool`

HasSshPasswordCredentials returns a boolean if a field has been set.

### SetSshPasswordCredentialsNil

`func (o *HiveSourceRegistrationParams) SetSshPasswordCredentialsNil(b bool)`

 SetSshPasswordCredentialsNil sets the value for SshPasswordCredentials to be an explicit nil

### UnsetSshPasswordCredentials
`func (o *HiveSourceRegistrationParams) UnsetSshPasswordCredentials()`

UnsetSshPasswordCredentials ensures that no value is present for SshPasswordCredentials, not even an explicit nil
### GetSshPrivateKeyCredentials

`func (o *HiveSourceRegistrationParams) GetSshPrivateKeyCredentials() HbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials`

GetSshPrivateKeyCredentials returns the SshPrivateKeyCredentials field if non-nil, zero value otherwise.

### GetSshPrivateKeyCredentialsOk

`func (o *HiveSourceRegistrationParams) GetSshPrivateKeyCredentialsOk() (*HbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials, bool)`

GetSshPrivateKeyCredentialsOk returns a tuple with the SshPrivateKeyCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPrivateKeyCredentials

`func (o *HiveSourceRegistrationParams) SetSshPrivateKeyCredentials(v HbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials)`

SetSshPrivateKeyCredentials sets SshPrivateKeyCredentials field to given value.

### HasSshPrivateKeyCredentials

`func (o *HiveSourceRegistrationParams) HasSshPrivateKeyCredentials() bool`

HasSshPrivateKeyCredentials returns a boolean if a field has been set.

### SetSshPrivateKeyCredentialsNil

`func (o *HiveSourceRegistrationParams) SetSshPrivateKeyCredentialsNil(b bool)`

 SetSshPrivateKeyCredentialsNil sets the value for SshPrivateKeyCredentials to be an explicit nil

### UnsetSshPrivateKeyCredentials
`func (o *HiveSourceRegistrationParams) UnsetSshPrivateKeyCredentials()`

UnsetSshPrivateKeyCredentials ensures that no value is present for SshPrivateKeyCredentials, not even an explicit nil
### GetHdfsSourceRegistrationID

`func (o *HiveSourceRegistrationParams) GetHdfsSourceRegistrationID() int64`

GetHdfsSourceRegistrationID returns the HdfsSourceRegistrationID field if non-nil, zero value otherwise.

### GetHdfsSourceRegistrationIDOk

`func (o *HiveSourceRegistrationParams) GetHdfsSourceRegistrationIDOk() (*int64, bool)`

GetHdfsSourceRegistrationIDOk returns a tuple with the HdfsSourceRegistrationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsSourceRegistrationID

`func (o *HiveSourceRegistrationParams) SetHdfsSourceRegistrationID(v int64)`

SetHdfsSourceRegistrationID sets HdfsSourceRegistrationID field to given value.


### GetKerberosPrincipal

`func (o *HiveSourceRegistrationParams) GetKerberosPrincipal() string`

GetKerberosPrincipal returns the KerberosPrincipal field if non-nil, zero value otherwise.

### GetKerberosPrincipalOk

`func (o *HiveSourceRegistrationParams) GetKerberosPrincipalOk() (*string, bool)`

GetKerberosPrincipalOk returns a tuple with the KerberosPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosPrincipal

`func (o *HiveSourceRegistrationParams) SetKerberosPrincipal(v string)`

SetKerberosPrincipal sets KerberosPrincipal field to given value.

### HasKerberosPrincipal

`func (o *HiveSourceRegistrationParams) HasKerberosPrincipal() bool`

HasKerberosPrincipal returns a boolean if a field has been set.

### SetKerberosPrincipalNil

`func (o *HiveSourceRegistrationParams) SetKerberosPrincipalNil(b bool)`

 SetKerberosPrincipalNil sets the value for KerberosPrincipal to be an explicit nil

### UnsetKerberosPrincipal
`func (o *HiveSourceRegistrationParams) UnsetKerberosPrincipal()`

UnsetKerberosPrincipal ensures that no value is present for KerberosPrincipal, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


