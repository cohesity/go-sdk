# HiveSourceRegistrationParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | IP or hostname of any host from which the Hive configuration file hive-site.xml can be read. | 
**ConfigurationDirectory** | **string** | The directory containing the hive-site.xml. | 
**SshPasswordCredentials** | Pointer to [**NullableHiveSourceRegistrationParamsAllOfSshPasswordCredentials**](HiveSourceRegistrationParamsAllOfSshPasswordCredentials.md) |  | [optional] 
**SshPrivateKeyCredentials** | Pointer to [**NullableHbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials**](HbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials.md) |  | [optional] 
**HdfsSourceRegistrationID** | **int64** | Protection Source registration id of the HDFS on which this Hive is running. | 
**KerberosPrincipal** | Pointer to **NullableString** | The kerberos principal to be used to connect to this Hive source. | [optional] 

## Methods

### NewHiveSourceRegistrationParamsAllOf

`func NewHiveSourceRegistrationParamsAllOf(host string, configurationDirectory string, hdfsSourceRegistrationID int64, ) *HiveSourceRegistrationParamsAllOf`

NewHiveSourceRegistrationParamsAllOf instantiates a new HiveSourceRegistrationParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHiveSourceRegistrationParamsAllOfWithDefaults

`func NewHiveSourceRegistrationParamsAllOfWithDefaults() *HiveSourceRegistrationParamsAllOf`

NewHiveSourceRegistrationParamsAllOfWithDefaults instantiates a new HiveSourceRegistrationParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *HiveSourceRegistrationParamsAllOf) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *HiveSourceRegistrationParamsAllOf) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *HiveSourceRegistrationParamsAllOf) SetHost(v string)`

SetHost sets Host field to given value.


### GetConfigurationDirectory

`func (o *HiveSourceRegistrationParamsAllOf) GetConfigurationDirectory() string`

GetConfigurationDirectory returns the ConfigurationDirectory field if non-nil, zero value otherwise.

### GetConfigurationDirectoryOk

`func (o *HiveSourceRegistrationParamsAllOf) GetConfigurationDirectoryOk() (*string, bool)`

GetConfigurationDirectoryOk returns a tuple with the ConfigurationDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationDirectory

`func (o *HiveSourceRegistrationParamsAllOf) SetConfigurationDirectory(v string)`

SetConfigurationDirectory sets ConfigurationDirectory field to given value.


### GetSshPasswordCredentials

`func (o *HiveSourceRegistrationParamsAllOf) GetSshPasswordCredentials() HiveSourceRegistrationParamsAllOfSshPasswordCredentials`

GetSshPasswordCredentials returns the SshPasswordCredentials field if non-nil, zero value otherwise.

### GetSshPasswordCredentialsOk

`func (o *HiveSourceRegistrationParamsAllOf) GetSshPasswordCredentialsOk() (*HiveSourceRegistrationParamsAllOfSshPasswordCredentials, bool)`

GetSshPasswordCredentialsOk returns a tuple with the SshPasswordCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPasswordCredentials

`func (o *HiveSourceRegistrationParamsAllOf) SetSshPasswordCredentials(v HiveSourceRegistrationParamsAllOfSshPasswordCredentials)`

SetSshPasswordCredentials sets SshPasswordCredentials field to given value.

### HasSshPasswordCredentials

`func (o *HiveSourceRegistrationParamsAllOf) HasSshPasswordCredentials() bool`

HasSshPasswordCredentials returns a boolean if a field has been set.

### SetSshPasswordCredentialsNil

`func (o *HiveSourceRegistrationParamsAllOf) SetSshPasswordCredentialsNil(b bool)`

 SetSshPasswordCredentialsNil sets the value for SshPasswordCredentials to be an explicit nil

### UnsetSshPasswordCredentials
`func (o *HiveSourceRegistrationParamsAllOf) UnsetSshPasswordCredentials()`

UnsetSshPasswordCredentials ensures that no value is present for SshPasswordCredentials, not even an explicit nil
### GetSshPrivateKeyCredentials

`func (o *HiveSourceRegistrationParamsAllOf) GetSshPrivateKeyCredentials() HbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials`

GetSshPrivateKeyCredentials returns the SshPrivateKeyCredentials field if non-nil, zero value otherwise.

### GetSshPrivateKeyCredentialsOk

`func (o *HiveSourceRegistrationParamsAllOf) GetSshPrivateKeyCredentialsOk() (*HbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials, bool)`

GetSshPrivateKeyCredentialsOk returns a tuple with the SshPrivateKeyCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPrivateKeyCredentials

`func (o *HiveSourceRegistrationParamsAllOf) SetSshPrivateKeyCredentials(v HbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials)`

SetSshPrivateKeyCredentials sets SshPrivateKeyCredentials field to given value.

### HasSshPrivateKeyCredentials

`func (o *HiveSourceRegistrationParamsAllOf) HasSshPrivateKeyCredentials() bool`

HasSshPrivateKeyCredentials returns a boolean if a field has been set.

### SetSshPrivateKeyCredentialsNil

`func (o *HiveSourceRegistrationParamsAllOf) SetSshPrivateKeyCredentialsNil(b bool)`

 SetSshPrivateKeyCredentialsNil sets the value for SshPrivateKeyCredentials to be an explicit nil

### UnsetSshPrivateKeyCredentials
`func (o *HiveSourceRegistrationParamsAllOf) UnsetSshPrivateKeyCredentials()`

UnsetSshPrivateKeyCredentials ensures that no value is present for SshPrivateKeyCredentials, not even an explicit nil
### GetHdfsSourceRegistrationID

`func (o *HiveSourceRegistrationParamsAllOf) GetHdfsSourceRegistrationID() int64`

GetHdfsSourceRegistrationID returns the HdfsSourceRegistrationID field if non-nil, zero value otherwise.

### GetHdfsSourceRegistrationIDOk

`func (o *HiveSourceRegistrationParamsAllOf) GetHdfsSourceRegistrationIDOk() (*int64, bool)`

GetHdfsSourceRegistrationIDOk returns a tuple with the HdfsSourceRegistrationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsSourceRegistrationID

`func (o *HiveSourceRegistrationParamsAllOf) SetHdfsSourceRegistrationID(v int64)`

SetHdfsSourceRegistrationID sets HdfsSourceRegistrationID field to given value.


### GetKerberosPrincipal

`func (o *HiveSourceRegistrationParamsAllOf) GetKerberosPrincipal() string`

GetKerberosPrincipal returns the KerberosPrincipal field if non-nil, zero value otherwise.

### GetKerberosPrincipalOk

`func (o *HiveSourceRegistrationParamsAllOf) GetKerberosPrincipalOk() (*string, bool)`

GetKerberosPrincipalOk returns a tuple with the KerberosPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosPrincipal

`func (o *HiveSourceRegistrationParamsAllOf) SetKerberosPrincipal(v string)`

SetKerberosPrincipal sets KerberosPrincipal field to given value.

### HasKerberosPrincipal

`func (o *HiveSourceRegistrationParamsAllOf) HasKerberosPrincipal() bool`

HasKerberosPrincipal returns a boolean if a field has been set.

### SetKerberosPrincipalNil

`func (o *HiveSourceRegistrationParamsAllOf) SetKerberosPrincipalNil(b bool)`

 SetKerberosPrincipalNil sets the value for KerberosPrincipal to be an explicit nil

### UnsetKerberosPrincipal
`func (o *HiveSourceRegistrationParamsAllOf) UnsetKerberosPrincipal()`

UnsetKerberosPrincipal ensures that no value is present for KerberosPrincipal, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


