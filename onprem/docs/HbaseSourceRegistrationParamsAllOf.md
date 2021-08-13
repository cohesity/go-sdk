# HbaseSourceRegistrationParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | IP or hostname of any host from which the HBase configuration file hbase-site.xml can be read. | 
**ConfigurationDirectory** | **string** | The directory containing the hbase-site.xml. | 
**SshPasswordCredentials** | Pointer to [**NullableHbaseSourceRegistrationParamsAllOfSshPasswordCredentials**](HbaseSourceRegistrationParamsAllOfSshPasswordCredentials.md) |  | [optional] 
**SshPrivateKeyCredentials** | Pointer to [**NullableHbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials**](HbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials.md) |  | [optional] 
**HdfsSourceRegistrationID** | **int64** | Protection Source registration id of the HDFS on which this HBase is running. | 
**KerberosPrincipal** | Pointer to **NullableString** | The kerberos principal to be used to connect to this Hbase source. | [optional] 

## Methods

### NewHbaseSourceRegistrationParamsAllOf

`func NewHbaseSourceRegistrationParamsAllOf(host string, configurationDirectory string, hdfsSourceRegistrationID int64, ) *HbaseSourceRegistrationParamsAllOf`

NewHbaseSourceRegistrationParamsAllOf instantiates a new HbaseSourceRegistrationParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHbaseSourceRegistrationParamsAllOfWithDefaults

`func NewHbaseSourceRegistrationParamsAllOfWithDefaults() *HbaseSourceRegistrationParamsAllOf`

NewHbaseSourceRegistrationParamsAllOfWithDefaults instantiates a new HbaseSourceRegistrationParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *HbaseSourceRegistrationParamsAllOf) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *HbaseSourceRegistrationParamsAllOf) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *HbaseSourceRegistrationParamsAllOf) SetHost(v string)`

SetHost sets Host field to given value.


### GetConfigurationDirectory

`func (o *HbaseSourceRegistrationParamsAllOf) GetConfigurationDirectory() string`

GetConfigurationDirectory returns the ConfigurationDirectory field if non-nil, zero value otherwise.

### GetConfigurationDirectoryOk

`func (o *HbaseSourceRegistrationParamsAllOf) GetConfigurationDirectoryOk() (*string, bool)`

GetConfigurationDirectoryOk returns a tuple with the ConfigurationDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationDirectory

`func (o *HbaseSourceRegistrationParamsAllOf) SetConfigurationDirectory(v string)`

SetConfigurationDirectory sets ConfigurationDirectory field to given value.


### GetSshPasswordCredentials

`func (o *HbaseSourceRegistrationParamsAllOf) GetSshPasswordCredentials() HbaseSourceRegistrationParamsAllOfSshPasswordCredentials`

GetSshPasswordCredentials returns the SshPasswordCredentials field if non-nil, zero value otherwise.

### GetSshPasswordCredentialsOk

`func (o *HbaseSourceRegistrationParamsAllOf) GetSshPasswordCredentialsOk() (*HbaseSourceRegistrationParamsAllOfSshPasswordCredentials, bool)`

GetSshPasswordCredentialsOk returns a tuple with the SshPasswordCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPasswordCredentials

`func (o *HbaseSourceRegistrationParamsAllOf) SetSshPasswordCredentials(v HbaseSourceRegistrationParamsAllOfSshPasswordCredentials)`

SetSshPasswordCredentials sets SshPasswordCredentials field to given value.

### HasSshPasswordCredentials

`func (o *HbaseSourceRegistrationParamsAllOf) HasSshPasswordCredentials() bool`

HasSshPasswordCredentials returns a boolean if a field has been set.

### SetSshPasswordCredentialsNil

`func (o *HbaseSourceRegistrationParamsAllOf) SetSshPasswordCredentialsNil(b bool)`

 SetSshPasswordCredentialsNil sets the value for SshPasswordCredentials to be an explicit nil

### UnsetSshPasswordCredentials
`func (o *HbaseSourceRegistrationParamsAllOf) UnsetSshPasswordCredentials()`

UnsetSshPasswordCredentials ensures that no value is present for SshPasswordCredentials, not even an explicit nil
### GetSshPrivateKeyCredentials

`func (o *HbaseSourceRegistrationParamsAllOf) GetSshPrivateKeyCredentials() HbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials`

GetSshPrivateKeyCredentials returns the SshPrivateKeyCredentials field if non-nil, zero value otherwise.

### GetSshPrivateKeyCredentialsOk

`func (o *HbaseSourceRegistrationParamsAllOf) GetSshPrivateKeyCredentialsOk() (*HbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials, bool)`

GetSshPrivateKeyCredentialsOk returns a tuple with the SshPrivateKeyCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPrivateKeyCredentials

`func (o *HbaseSourceRegistrationParamsAllOf) SetSshPrivateKeyCredentials(v HbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials)`

SetSshPrivateKeyCredentials sets SshPrivateKeyCredentials field to given value.

### HasSshPrivateKeyCredentials

`func (o *HbaseSourceRegistrationParamsAllOf) HasSshPrivateKeyCredentials() bool`

HasSshPrivateKeyCredentials returns a boolean if a field has been set.

### SetSshPrivateKeyCredentialsNil

`func (o *HbaseSourceRegistrationParamsAllOf) SetSshPrivateKeyCredentialsNil(b bool)`

 SetSshPrivateKeyCredentialsNil sets the value for SshPrivateKeyCredentials to be an explicit nil

### UnsetSshPrivateKeyCredentials
`func (o *HbaseSourceRegistrationParamsAllOf) UnsetSshPrivateKeyCredentials()`

UnsetSshPrivateKeyCredentials ensures that no value is present for SshPrivateKeyCredentials, not even an explicit nil
### GetHdfsSourceRegistrationID

`func (o *HbaseSourceRegistrationParamsAllOf) GetHdfsSourceRegistrationID() int64`

GetHdfsSourceRegistrationID returns the HdfsSourceRegistrationID field if non-nil, zero value otherwise.

### GetHdfsSourceRegistrationIDOk

`func (o *HbaseSourceRegistrationParamsAllOf) GetHdfsSourceRegistrationIDOk() (*int64, bool)`

GetHdfsSourceRegistrationIDOk returns a tuple with the HdfsSourceRegistrationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsSourceRegistrationID

`func (o *HbaseSourceRegistrationParamsAllOf) SetHdfsSourceRegistrationID(v int64)`

SetHdfsSourceRegistrationID sets HdfsSourceRegistrationID field to given value.


### GetKerberosPrincipal

`func (o *HbaseSourceRegistrationParamsAllOf) GetKerberosPrincipal() string`

GetKerberosPrincipal returns the KerberosPrincipal field if non-nil, zero value otherwise.

### GetKerberosPrincipalOk

`func (o *HbaseSourceRegistrationParamsAllOf) GetKerberosPrincipalOk() (*string, bool)`

GetKerberosPrincipalOk returns a tuple with the KerberosPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosPrincipal

`func (o *HbaseSourceRegistrationParamsAllOf) SetKerberosPrincipal(v string)`

SetKerberosPrincipal sets KerberosPrincipal field to given value.

### HasKerberosPrincipal

`func (o *HbaseSourceRegistrationParamsAllOf) HasKerberosPrincipal() bool`

HasKerberosPrincipal returns a boolean if a field has been set.

### SetKerberosPrincipalNil

`func (o *HbaseSourceRegistrationParamsAllOf) SetKerberosPrincipalNil(b bool)`

 SetKerberosPrincipalNil sets the value for KerberosPrincipal to be an explicit nil

### UnsetKerberosPrincipal
`func (o *HbaseSourceRegistrationParamsAllOf) UnsetKerberosPrincipal()`

UnsetKerberosPrincipal ensures that no value is present for KerberosPrincipal, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


