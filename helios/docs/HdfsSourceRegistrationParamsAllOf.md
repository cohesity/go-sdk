# HdfsSourceRegistrationParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | IP or hostname of any host from which the HDFS configuration files core-site.xml and hdfs-site.xml can be read. | 
**ConfigurationDirectory** | **string** | The directory containing the core-site.xml and hdfs-site.xml configuration files. | 
**SshPasswordCredentials** | Pointer to [**NullableHbaseSourceRegistrationParamsAllOfSshPasswordCredentials**](HbaseSourceRegistrationParamsAllOfSshPasswordCredentials.md) |  | [optional] 
**SshPrivateKeyCredentials** | Pointer to [**NullableHbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials**](HbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials.md) |  | [optional] 
**KerberosPrincipal** | Pointer to **NullableString** | The kerberos principal to be used to connect to this HDFS source. | [optional] 
**HadoopDistribution** | **string** | The hadoop distribution for this cluster. This can be either &#39;CDH&#39; or &#39;HDP&#39; | 
**HadoopVersion** | **string** | The hadoop version for this cluster. | 

## Methods

### NewHdfsSourceRegistrationParamsAllOf

`func NewHdfsSourceRegistrationParamsAllOf(host string, configurationDirectory string, hadoopDistribution string, hadoopVersion string, ) *HdfsSourceRegistrationParamsAllOf`

NewHdfsSourceRegistrationParamsAllOf instantiates a new HdfsSourceRegistrationParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHdfsSourceRegistrationParamsAllOfWithDefaults

`func NewHdfsSourceRegistrationParamsAllOfWithDefaults() *HdfsSourceRegistrationParamsAllOf`

NewHdfsSourceRegistrationParamsAllOfWithDefaults instantiates a new HdfsSourceRegistrationParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *HdfsSourceRegistrationParamsAllOf) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *HdfsSourceRegistrationParamsAllOf) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *HdfsSourceRegistrationParamsAllOf) SetHost(v string)`

SetHost sets Host field to given value.


### GetConfigurationDirectory

`func (o *HdfsSourceRegistrationParamsAllOf) GetConfigurationDirectory() string`

GetConfigurationDirectory returns the ConfigurationDirectory field if non-nil, zero value otherwise.

### GetConfigurationDirectoryOk

`func (o *HdfsSourceRegistrationParamsAllOf) GetConfigurationDirectoryOk() (*string, bool)`

GetConfigurationDirectoryOk returns a tuple with the ConfigurationDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationDirectory

`func (o *HdfsSourceRegistrationParamsAllOf) SetConfigurationDirectory(v string)`

SetConfigurationDirectory sets ConfigurationDirectory field to given value.


### GetSshPasswordCredentials

`func (o *HdfsSourceRegistrationParamsAllOf) GetSshPasswordCredentials() HbaseSourceRegistrationParamsAllOfSshPasswordCredentials`

GetSshPasswordCredentials returns the SshPasswordCredentials field if non-nil, zero value otherwise.

### GetSshPasswordCredentialsOk

`func (o *HdfsSourceRegistrationParamsAllOf) GetSshPasswordCredentialsOk() (*HbaseSourceRegistrationParamsAllOfSshPasswordCredentials, bool)`

GetSshPasswordCredentialsOk returns a tuple with the SshPasswordCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPasswordCredentials

`func (o *HdfsSourceRegistrationParamsAllOf) SetSshPasswordCredentials(v HbaseSourceRegistrationParamsAllOfSshPasswordCredentials)`

SetSshPasswordCredentials sets SshPasswordCredentials field to given value.

### HasSshPasswordCredentials

`func (o *HdfsSourceRegistrationParamsAllOf) HasSshPasswordCredentials() bool`

HasSshPasswordCredentials returns a boolean if a field has been set.

### SetSshPasswordCredentialsNil

`func (o *HdfsSourceRegistrationParamsAllOf) SetSshPasswordCredentialsNil(b bool)`

 SetSshPasswordCredentialsNil sets the value for SshPasswordCredentials to be an explicit nil

### UnsetSshPasswordCredentials
`func (o *HdfsSourceRegistrationParamsAllOf) UnsetSshPasswordCredentials()`

UnsetSshPasswordCredentials ensures that no value is present for SshPasswordCredentials, not even an explicit nil
### GetSshPrivateKeyCredentials

`func (o *HdfsSourceRegistrationParamsAllOf) GetSshPrivateKeyCredentials() HbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials`

GetSshPrivateKeyCredentials returns the SshPrivateKeyCredentials field if non-nil, zero value otherwise.

### GetSshPrivateKeyCredentialsOk

`func (o *HdfsSourceRegistrationParamsAllOf) GetSshPrivateKeyCredentialsOk() (*HbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials, bool)`

GetSshPrivateKeyCredentialsOk returns a tuple with the SshPrivateKeyCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPrivateKeyCredentials

`func (o *HdfsSourceRegistrationParamsAllOf) SetSshPrivateKeyCredentials(v HbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials)`

SetSshPrivateKeyCredentials sets SshPrivateKeyCredentials field to given value.

### HasSshPrivateKeyCredentials

`func (o *HdfsSourceRegistrationParamsAllOf) HasSshPrivateKeyCredentials() bool`

HasSshPrivateKeyCredentials returns a boolean if a field has been set.

### SetSshPrivateKeyCredentialsNil

`func (o *HdfsSourceRegistrationParamsAllOf) SetSshPrivateKeyCredentialsNil(b bool)`

 SetSshPrivateKeyCredentialsNil sets the value for SshPrivateKeyCredentials to be an explicit nil

### UnsetSshPrivateKeyCredentials
`func (o *HdfsSourceRegistrationParamsAllOf) UnsetSshPrivateKeyCredentials()`

UnsetSshPrivateKeyCredentials ensures that no value is present for SshPrivateKeyCredentials, not even an explicit nil
### GetKerberosPrincipal

`func (o *HdfsSourceRegistrationParamsAllOf) GetKerberosPrincipal() string`

GetKerberosPrincipal returns the KerberosPrincipal field if non-nil, zero value otherwise.

### GetKerberosPrincipalOk

`func (o *HdfsSourceRegistrationParamsAllOf) GetKerberosPrincipalOk() (*string, bool)`

GetKerberosPrincipalOk returns a tuple with the KerberosPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosPrincipal

`func (o *HdfsSourceRegistrationParamsAllOf) SetKerberosPrincipal(v string)`

SetKerberosPrincipal sets KerberosPrincipal field to given value.

### HasKerberosPrincipal

`func (o *HdfsSourceRegistrationParamsAllOf) HasKerberosPrincipal() bool`

HasKerberosPrincipal returns a boolean if a field has been set.

### SetKerberosPrincipalNil

`func (o *HdfsSourceRegistrationParamsAllOf) SetKerberosPrincipalNil(b bool)`

 SetKerberosPrincipalNil sets the value for KerberosPrincipal to be an explicit nil

### UnsetKerberosPrincipal
`func (o *HdfsSourceRegistrationParamsAllOf) UnsetKerberosPrincipal()`

UnsetKerberosPrincipal ensures that no value is present for KerberosPrincipal, not even an explicit nil
### GetHadoopDistribution

`func (o *HdfsSourceRegistrationParamsAllOf) GetHadoopDistribution() string`

GetHadoopDistribution returns the HadoopDistribution field if non-nil, zero value otherwise.

### GetHadoopDistributionOk

`func (o *HdfsSourceRegistrationParamsAllOf) GetHadoopDistributionOk() (*string, bool)`

GetHadoopDistributionOk returns a tuple with the HadoopDistribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHadoopDistribution

`func (o *HdfsSourceRegistrationParamsAllOf) SetHadoopDistribution(v string)`

SetHadoopDistribution sets HadoopDistribution field to given value.


### GetHadoopVersion

`func (o *HdfsSourceRegistrationParamsAllOf) GetHadoopVersion() string`

GetHadoopVersion returns the HadoopVersion field if non-nil, zero value otherwise.

### GetHadoopVersionOk

`func (o *HdfsSourceRegistrationParamsAllOf) GetHadoopVersionOk() (*string, bool)`

GetHadoopVersionOk returns a tuple with the HadoopVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHadoopVersion

`func (o *HdfsSourceRegistrationParamsAllOf) SetHadoopVersion(v string)`

SetHadoopVersion sets HadoopVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


