# HbaseSourceRegistrationParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthType** | Pointer to **NullableString** | Authentication type. | [optional] [readonly] 
**DataRootDirectory** | Pointer to **NullableString** | The &#39;Data root directory&#39; for this HBase. | [optional] [readonly] 
**ZookeeperQuorum** | Pointer to **[]string** | The &#39;Zookeeper Quorum&#39; for this HBase. | [optional] [readonly] 
**ConfigurationDirectory** | **string** | The directory containing the hbase-site.xml. | 
**HdfsSourceRegistrationID** | **int64** | Protection Source registration id of the HDFS on which this HBase is running. | 
**Host** | **string** | IP or hostname of any host from which the HBase configuration file hbase-site.xml can be read. | 
**KerberosPrincipal** | Pointer to **NullableString** | The kerberos principal to be used to connect to this Hbase source. | [optional] 
**SshPasswordCredentials** | Pointer to [**NullableHbaseSourceRegistrationParamsAllOfSshPasswordCredentials**](HbaseSourceRegistrationParamsAllOfSshPasswordCredentials.md) |  | [optional] 
**SshPrivateKeyCredentials** | Pointer to [**NullableHbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials**](HbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials.md) |  | [optional] 

## Methods

### NewHbaseSourceRegistrationParams

`func NewHbaseSourceRegistrationParams(configurationDirectory string, hdfsSourceRegistrationID int64, host string, ) *HbaseSourceRegistrationParams`

NewHbaseSourceRegistrationParams instantiates a new HbaseSourceRegistrationParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHbaseSourceRegistrationParamsWithDefaults

`func NewHbaseSourceRegistrationParamsWithDefaults() *HbaseSourceRegistrationParams`

NewHbaseSourceRegistrationParamsWithDefaults instantiates a new HbaseSourceRegistrationParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthType

`func (o *HbaseSourceRegistrationParams) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *HbaseSourceRegistrationParams) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *HbaseSourceRegistrationParams) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *HbaseSourceRegistrationParams) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### SetAuthTypeNil

`func (o *HbaseSourceRegistrationParams) SetAuthTypeNil(b bool)`

 SetAuthTypeNil sets the value for AuthType to be an explicit nil

### UnsetAuthType
`func (o *HbaseSourceRegistrationParams) UnsetAuthType()`

UnsetAuthType ensures that no value is present for AuthType, not even an explicit nil
### GetDataRootDirectory

`func (o *HbaseSourceRegistrationParams) GetDataRootDirectory() string`

GetDataRootDirectory returns the DataRootDirectory field if non-nil, zero value otherwise.

### GetDataRootDirectoryOk

`func (o *HbaseSourceRegistrationParams) GetDataRootDirectoryOk() (*string, bool)`

GetDataRootDirectoryOk returns a tuple with the DataRootDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRootDirectory

`func (o *HbaseSourceRegistrationParams) SetDataRootDirectory(v string)`

SetDataRootDirectory sets DataRootDirectory field to given value.

### HasDataRootDirectory

`func (o *HbaseSourceRegistrationParams) HasDataRootDirectory() bool`

HasDataRootDirectory returns a boolean if a field has been set.

### SetDataRootDirectoryNil

`func (o *HbaseSourceRegistrationParams) SetDataRootDirectoryNil(b bool)`

 SetDataRootDirectoryNil sets the value for DataRootDirectory to be an explicit nil

### UnsetDataRootDirectory
`func (o *HbaseSourceRegistrationParams) UnsetDataRootDirectory()`

UnsetDataRootDirectory ensures that no value is present for DataRootDirectory, not even an explicit nil
### GetZookeeperQuorum

`func (o *HbaseSourceRegistrationParams) GetZookeeperQuorum() []string`

GetZookeeperQuorum returns the ZookeeperQuorum field if non-nil, zero value otherwise.

### GetZookeeperQuorumOk

`func (o *HbaseSourceRegistrationParams) GetZookeeperQuorumOk() (*[]string, bool)`

GetZookeeperQuorumOk returns a tuple with the ZookeeperQuorum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZookeeperQuorum

`func (o *HbaseSourceRegistrationParams) SetZookeeperQuorum(v []string)`

SetZookeeperQuorum sets ZookeeperQuorum field to given value.

### HasZookeeperQuorum

`func (o *HbaseSourceRegistrationParams) HasZookeeperQuorum() bool`

HasZookeeperQuorum returns a boolean if a field has been set.

### SetZookeeperQuorumNil

`func (o *HbaseSourceRegistrationParams) SetZookeeperQuorumNil(b bool)`

 SetZookeeperQuorumNil sets the value for ZookeeperQuorum to be an explicit nil

### UnsetZookeeperQuorum
`func (o *HbaseSourceRegistrationParams) UnsetZookeeperQuorum()`

UnsetZookeeperQuorum ensures that no value is present for ZookeeperQuorum, not even an explicit nil
### GetConfigurationDirectory

`func (o *HbaseSourceRegistrationParams) GetConfigurationDirectory() string`

GetConfigurationDirectory returns the ConfigurationDirectory field if non-nil, zero value otherwise.

### GetConfigurationDirectoryOk

`func (o *HbaseSourceRegistrationParams) GetConfigurationDirectoryOk() (*string, bool)`

GetConfigurationDirectoryOk returns a tuple with the ConfigurationDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationDirectory

`func (o *HbaseSourceRegistrationParams) SetConfigurationDirectory(v string)`

SetConfigurationDirectory sets ConfigurationDirectory field to given value.


### GetHdfsSourceRegistrationID

`func (o *HbaseSourceRegistrationParams) GetHdfsSourceRegistrationID() int64`

GetHdfsSourceRegistrationID returns the HdfsSourceRegistrationID field if non-nil, zero value otherwise.

### GetHdfsSourceRegistrationIDOk

`func (o *HbaseSourceRegistrationParams) GetHdfsSourceRegistrationIDOk() (*int64, bool)`

GetHdfsSourceRegistrationIDOk returns a tuple with the HdfsSourceRegistrationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsSourceRegistrationID

`func (o *HbaseSourceRegistrationParams) SetHdfsSourceRegistrationID(v int64)`

SetHdfsSourceRegistrationID sets HdfsSourceRegistrationID field to given value.


### GetHost

`func (o *HbaseSourceRegistrationParams) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *HbaseSourceRegistrationParams) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *HbaseSourceRegistrationParams) SetHost(v string)`

SetHost sets Host field to given value.


### GetKerberosPrincipal

`func (o *HbaseSourceRegistrationParams) GetKerberosPrincipal() string`

GetKerberosPrincipal returns the KerberosPrincipal field if non-nil, zero value otherwise.

### GetKerberosPrincipalOk

`func (o *HbaseSourceRegistrationParams) GetKerberosPrincipalOk() (*string, bool)`

GetKerberosPrincipalOk returns a tuple with the KerberosPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosPrincipal

`func (o *HbaseSourceRegistrationParams) SetKerberosPrincipal(v string)`

SetKerberosPrincipal sets KerberosPrincipal field to given value.

### HasKerberosPrincipal

`func (o *HbaseSourceRegistrationParams) HasKerberosPrincipal() bool`

HasKerberosPrincipal returns a boolean if a field has been set.

### SetKerberosPrincipalNil

`func (o *HbaseSourceRegistrationParams) SetKerberosPrincipalNil(b bool)`

 SetKerberosPrincipalNil sets the value for KerberosPrincipal to be an explicit nil

### UnsetKerberosPrincipal
`func (o *HbaseSourceRegistrationParams) UnsetKerberosPrincipal()`

UnsetKerberosPrincipal ensures that no value is present for KerberosPrincipal, not even an explicit nil
### GetSshPasswordCredentials

`func (o *HbaseSourceRegistrationParams) GetSshPasswordCredentials() HbaseSourceRegistrationParamsAllOfSshPasswordCredentials`

GetSshPasswordCredentials returns the SshPasswordCredentials field if non-nil, zero value otherwise.

### GetSshPasswordCredentialsOk

`func (o *HbaseSourceRegistrationParams) GetSshPasswordCredentialsOk() (*HbaseSourceRegistrationParamsAllOfSshPasswordCredentials, bool)`

GetSshPasswordCredentialsOk returns a tuple with the SshPasswordCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPasswordCredentials

`func (o *HbaseSourceRegistrationParams) SetSshPasswordCredentials(v HbaseSourceRegistrationParamsAllOfSshPasswordCredentials)`

SetSshPasswordCredentials sets SshPasswordCredentials field to given value.

### HasSshPasswordCredentials

`func (o *HbaseSourceRegistrationParams) HasSshPasswordCredentials() bool`

HasSshPasswordCredentials returns a boolean if a field has been set.

### SetSshPasswordCredentialsNil

`func (o *HbaseSourceRegistrationParams) SetSshPasswordCredentialsNil(b bool)`

 SetSshPasswordCredentialsNil sets the value for SshPasswordCredentials to be an explicit nil

### UnsetSshPasswordCredentials
`func (o *HbaseSourceRegistrationParams) UnsetSshPasswordCredentials()`

UnsetSshPasswordCredentials ensures that no value is present for SshPasswordCredentials, not even an explicit nil
### GetSshPrivateKeyCredentials

`func (o *HbaseSourceRegistrationParams) GetSshPrivateKeyCredentials() HbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials`

GetSshPrivateKeyCredentials returns the SshPrivateKeyCredentials field if non-nil, zero value otherwise.

### GetSshPrivateKeyCredentialsOk

`func (o *HbaseSourceRegistrationParams) GetSshPrivateKeyCredentialsOk() (*HbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials, bool)`

GetSshPrivateKeyCredentialsOk returns a tuple with the SshPrivateKeyCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPrivateKeyCredentials

`func (o *HbaseSourceRegistrationParams) SetSshPrivateKeyCredentials(v HbaseSourceRegistrationParamsAllOfSshPrivateKeyCredentials)`

SetSshPrivateKeyCredentials sets SshPrivateKeyCredentials field to given value.

### HasSshPrivateKeyCredentials

`func (o *HbaseSourceRegistrationParams) HasSshPrivateKeyCredentials() bool`

HasSshPrivateKeyCredentials returns a boolean if a field has been set.

### SetSshPrivateKeyCredentialsNil

`func (o *HbaseSourceRegistrationParams) SetSshPrivateKeyCredentialsNil(b bool)`

 SetSshPrivateKeyCredentialsNil sets the value for SshPrivateKeyCredentials to be an explicit nil

### UnsetSshPrivateKeyCredentials
`func (o *HbaseSourceRegistrationParams) UnsetSshPrivateKeyCredentials()`

UnsetSshPrivateKeyCredentials ensures that no value is present for SshPrivateKeyCredentials, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


