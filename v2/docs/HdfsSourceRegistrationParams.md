# HdfsSourceRegistrationParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthType** | Pointer to **NullableString** | Authentication type. | [optional] [readonly] 
**NamenodeAddress** | Pointer to **string** | The HDFS Namenode IP or hostname. | [optional] [readonly] 
**WebhdfsPort** | Pointer to **int32** | The HDFS WebHDFS port. | [optional] [readonly] 
**ConfigurationDirectory** | **string** | The directory containing the core-site.xml and hdfs-site.xml configuration files. | 
**ConnectionType** | Pointer to **NullableString** | HDFS Connection Type. | [optional] 
**HadoopDistribution** | **string** | The hadoop distribution for this cluster. This can be either &#39;CDH&#39; or &#39;HDP&#39; | 
**HadoopVersion** | **string** | The hadoop version for this cluster. | 
**Host** | **string** | IP or hostname of any host from which the HDFS configuration files core-site.xml and hdfs-site.xml can be read. | 
**KerberosPrincipal** | Pointer to **NullableString** | The kerberos principal to be used to connect to this HDFS source. | [optional] 
**SshPasswordCredentials** | Pointer to [**NullableHbaseSourceRegistrationParamsAllOfSshPasswordCredentials**](HbaseSourceRegistrationParamsAllOfSshPasswordCredentials.md) |  | [optional] 
**SshPrivateKeyCredentials** | Pointer to [**NullableHbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials**](HbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials.md) |  | [optional] 

## Methods

### NewHdfsSourceRegistrationParams

`func NewHdfsSourceRegistrationParams(configurationDirectory string, hadoopDistribution string, hadoopVersion string, host string, ) *HdfsSourceRegistrationParams`

NewHdfsSourceRegistrationParams instantiates a new HdfsSourceRegistrationParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHdfsSourceRegistrationParamsWithDefaults

`func NewHdfsSourceRegistrationParamsWithDefaults() *HdfsSourceRegistrationParams`

NewHdfsSourceRegistrationParamsWithDefaults instantiates a new HdfsSourceRegistrationParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthType

`func (o *HdfsSourceRegistrationParams) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *HdfsSourceRegistrationParams) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *HdfsSourceRegistrationParams) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *HdfsSourceRegistrationParams) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### SetAuthTypeNil

`func (o *HdfsSourceRegistrationParams) SetAuthTypeNil(b bool)`

 SetAuthTypeNil sets the value for AuthType to be an explicit nil

### UnsetAuthType
`func (o *HdfsSourceRegistrationParams) UnsetAuthType()`

UnsetAuthType ensures that no value is present for AuthType, not even an explicit nil
### GetNamenodeAddress

`func (o *HdfsSourceRegistrationParams) GetNamenodeAddress() string`

GetNamenodeAddress returns the NamenodeAddress field if non-nil, zero value otherwise.

### GetNamenodeAddressOk

`func (o *HdfsSourceRegistrationParams) GetNamenodeAddressOk() (*string, bool)`

GetNamenodeAddressOk returns a tuple with the NamenodeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamenodeAddress

`func (o *HdfsSourceRegistrationParams) SetNamenodeAddress(v string)`

SetNamenodeAddress sets NamenodeAddress field to given value.

### HasNamenodeAddress

`func (o *HdfsSourceRegistrationParams) HasNamenodeAddress() bool`

HasNamenodeAddress returns a boolean if a field has been set.

### GetWebhdfsPort

`func (o *HdfsSourceRegistrationParams) GetWebhdfsPort() int32`

GetWebhdfsPort returns the WebhdfsPort field if non-nil, zero value otherwise.

### GetWebhdfsPortOk

`func (o *HdfsSourceRegistrationParams) GetWebhdfsPortOk() (*int32, bool)`

GetWebhdfsPortOk returns a tuple with the WebhdfsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhdfsPort

`func (o *HdfsSourceRegistrationParams) SetWebhdfsPort(v int32)`

SetWebhdfsPort sets WebhdfsPort field to given value.

### HasWebhdfsPort

`func (o *HdfsSourceRegistrationParams) HasWebhdfsPort() bool`

HasWebhdfsPort returns a boolean if a field has been set.

### GetConfigurationDirectory

`func (o *HdfsSourceRegistrationParams) GetConfigurationDirectory() string`

GetConfigurationDirectory returns the ConfigurationDirectory field if non-nil, zero value otherwise.

### GetConfigurationDirectoryOk

`func (o *HdfsSourceRegistrationParams) GetConfigurationDirectoryOk() (*string, bool)`

GetConfigurationDirectoryOk returns a tuple with the ConfigurationDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationDirectory

`func (o *HdfsSourceRegistrationParams) SetConfigurationDirectory(v string)`

SetConfigurationDirectory sets ConfigurationDirectory field to given value.


### GetConnectionType

`func (o *HdfsSourceRegistrationParams) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *HdfsSourceRegistrationParams) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *HdfsSourceRegistrationParams) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *HdfsSourceRegistrationParams) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### SetConnectionTypeNil

`func (o *HdfsSourceRegistrationParams) SetConnectionTypeNil(b bool)`

 SetConnectionTypeNil sets the value for ConnectionType to be an explicit nil

### UnsetConnectionType
`func (o *HdfsSourceRegistrationParams) UnsetConnectionType()`

UnsetConnectionType ensures that no value is present for ConnectionType, not even an explicit nil
### GetHadoopDistribution

`func (o *HdfsSourceRegistrationParams) GetHadoopDistribution() string`

GetHadoopDistribution returns the HadoopDistribution field if non-nil, zero value otherwise.

### GetHadoopDistributionOk

`func (o *HdfsSourceRegistrationParams) GetHadoopDistributionOk() (*string, bool)`

GetHadoopDistributionOk returns a tuple with the HadoopDistribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHadoopDistribution

`func (o *HdfsSourceRegistrationParams) SetHadoopDistribution(v string)`

SetHadoopDistribution sets HadoopDistribution field to given value.


### GetHadoopVersion

`func (o *HdfsSourceRegistrationParams) GetHadoopVersion() string`

GetHadoopVersion returns the HadoopVersion field if non-nil, zero value otherwise.

### GetHadoopVersionOk

`func (o *HdfsSourceRegistrationParams) GetHadoopVersionOk() (*string, bool)`

GetHadoopVersionOk returns a tuple with the HadoopVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHadoopVersion

`func (o *HdfsSourceRegistrationParams) SetHadoopVersion(v string)`

SetHadoopVersion sets HadoopVersion field to given value.


### GetHost

`func (o *HdfsSourceRegistrationParams) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *HdfsSourceRegistrationParams) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *HdfsSourceRegistrationParams) SetHost(v string)`

SetHost sets Host field to given value.


### GetKerberosPrincipal

`func (o *HdfsSourceRegistrationParams) GetKerberosPrincipal() string`

GetKerberosPrincipal returns the KerberosPrincipal field if non-nil, zero value otherwise.

### GetKerberosPrincipalOk

`func (o *HdfsSourceRegistrationParams) GetKerberosPrincipalOk() (*string, bool)`

GetKerberosPrincipalOk returns a tuple with the KerberosPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosPrincipal

`func (o *HdfsSourceRegistrationParams) SetKerberosPrincipal(v string)`

SetKerberosPrincipal sets KerberosPrincipal field to given value.

### HasKerberosPrincipal

`func (o *HdfsSourceRegistrationParams) HasKerberosPrincipal() bool`

HasKerberosPrincipal returns a boolean if a field has been set.

### SetKerberosPrincipalNil

`func (o *HdfsSourceRegistrationParams) SetKerberosPrincipalNil(b bool)`

 SetKerberosPrincipalNil sets the value for KerberosPrincipal to be an explicit nil

### UnsetKerberosPrincipal
`func (o *HdfsSourceRegistrationParams) UnsetKerberosPrincipal()`

UnsetKerberosPrincipal ensures that no value is present for KerberosPrincipal, not even an explicit nil
### GetSshPasswordCredentials

`func (o *HdfsSourceRegistrationParams) GetSshPasswordCredentials() HbaseSourceRegistrationParamsAllOfSshPasswordCredentials`

GetSshPasswordCredentials returns the SshPasswordCredentials field if non-nil, zero value otherwise.

### GetSshPasswordCredentialsOk

`func (o *HdfsSourceRegistrationParams) GetSshPasswordCredentialsOk() (*HbaseSourceRegistrationParamsAllOfSshPasswordCredentials, bool)`

GetSshPasswordCredentialsOk returns a tuple with the SshPasswordCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPasswordCredentials

`func (o *HdfsSourceRegistrationParams) SetSshPasswordCredentials(v HbaseSourceRegistrationParamsAllOfSshPasswordCredentials)`

SetSshPasswordCredentials sets SshPasswordCredentials field to given value.

### HasSshPasswordCredentials

`func (o *HdfsSourceRegistrationParams) HasSshPasswordCredentials() bool`

HasSshPasswordCredentials returns a boolean if a field has been set.

### SetSshPasswordCredentialsNil

`func (o *HdfsSourceRegistrationParams) SetSshPasswordCredentialsNil(b bool)`

 SetSshPasswordCredentialsNil sets the value for SshPasswordCredentials to be an explicit nil

### UnsetSshPasswordCredentials
`func (o *HdfsSourceRegistrationParams) UnsetSshPasswordCredentials()`

UnsetSshPasswordCredentials ensures that no value is present for SshPasswordCredentials, not even an explicit nil
### GetSshPrivateKeyCredentials

`func (o *HdfsSourceRegistrationParams) GetSshPrivateKeyCredentials() HbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials`

GetSshPrivateKeyCredentials returns the SshPrivateKeyCredentials field if non-nil, zero value otherwise.

### GetSshPrivateKeyCredentialsOk

`func (o *HdfsSourceRegistrationParams) GetSshPrivateKeyCredentialsOk() (*HbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials, bool)`

GetSshPrivateKeyCredentialsOk returns a tuple with the SshPrivateKeyCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPrivateKeyCredentials

`func (o *HdfsSourceRegistrationParams) SetSshPrivateKeyCredentials(v HbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials)`

SetSshPrivateKeyCredentials sets SshPrivateKeyCredentials field to given value.

### HasSshPrivateKeyCredentials

`func (o *HdfsSourceRegistrationParams) HasSshPrivateKeyCredentials() bool`

HasSshPrivateKeyCredentials returns a boolean if a field has been set.

### SetSshPrivateKeyCredentialsNil

`func (o *HdfsSourceRegistrationParams) SetSshPrivateKeyCredentialsNil(b bool)`

 SetSshPrivateKeyCredentialsNil sets the value for SshPrivateKeyCredentials to be an explicit nil

### UnsetSshPrivateKeyCredentials
`func (o *HdfsSourceRegistrationParams) UnsetSshPrivateKeyCredentials()`

UnsetSshPrivateKeyCredentials ensures that no value is present for SshPrivateKeyCredentials, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


