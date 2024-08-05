# HadoopConnectionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigurationDirectory** | **string** | The directory containing the application specific config file. . | 
**HdfsConnectionType** | Pointer to **NullableString** | HDFS Connection Type. | [optional] 
**Host** | **string** | IP or hostname of any host from which the  configuration file can be read. | 
**SshPasswordCredentials** | Pointer to [**NullableHadoopConnectionParamsSshPasswordCredentials**](HadoopConnectionParamsSshPasswordCredentials.md) |  | [optional] 
**SshPrivateKeyCredentials** | Pointer to [**NullableHadoopConnectionParamsSshPrivateKeyCredentials**](HadoopConnectionParamsSshPrivateKeyCredentials.md) |  | [optional] 

## Methods

### NewHadoopConnectionParams

`func NewHadoopConnectionParams(configurationDirectory string, host string, ) *HadoopConnectionParams`

NewHadoopConnectionParams instantiates a new HadoopConnectionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHadoopConnectionParamsWithDefaults

`func NewHadoopConnectionParamsWithDefaults() *HadoopConnectionParams`

NewHadoopConnectionParamsWithDefaults instantiates a new HadoopConnectionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurationDirectory

`func (o *HadoopConnectionParams) GetConfigurationDirectory() string`

GetConfigurationDirectory returns the ConfigurationDirectory field if non-nil, zero value otherwise.

### GetConfigurationDirectoryOk

`func (o *HadoopConnectionParams) GetConfigurationDirectoryOk() (*string, bool)`

GetConfigurationDirectoryOk returns a tuple with the ConfigurationDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationDirectory

`func (o *HadoopConnectionParams) SetConfigurationDirectory(v string)`

SetConfigurationDirectory sets ConfigurationDirectory field to given value.


### GetHdfsConnectionType

`func (o *HadoopConnectionParams) GetHdfsConnectionType() string`

GetHdfsConnectionType returns the HdfsConnectionType field if non-nil, zero value otherwise.

### GetHdfsConnectionTypeOk

`func (o *HadoopConnectionParams) GetHdfsConnectionTypeOk() (*string, bool)`

GetHdfsConnectionTypeOk returns a tuple with the HdfsConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsConnectionType

`func (o *HadoopConnectionParams) SetHdfsConnectionType(v string)`

SetHdfsConnectionType sets HdfsConnectionType field to given value.

### HasHdfsConnectionType

`func (o *HadoopConnectionParams) HasHdfsConnectionType() bool`

HasHdfsConnectionType returns a boolean if a field has been set.

### SetHdfsConnectionTypeNil

`func (o *HadoopConnectionParams) SetHdfsConnectionTypeNil(b bool)`

 SetHdfsConnectionTypeNil sets the value for HdfsConnectionType to be an explicit nil

### UnsetHdfsConnectionType
`func (o *HadoopConnectionParams) UnsetHdfsConnectionType()`

UnsetHdfsConnectionType ensures that no value is present for HdfsConnectionType, not even an explicit nil
### GetHost

`func (o *HadoopConnectionParams) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *HadoopConnectionParams) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *HadoopConnectionParams) SetHost(v string)`

SetHost sets Host field to given value.


### GetSshPasswordCredentials

`func (o *HadoopConnectionParams) GetSshPasswordCredentials() HadoopConnectionParamsSshPasswordCredentials`

GetSshPasswordCredentials returns the SshPasswordCredentials field if non-nil, zero value otherwise.

### GetSshPasswordCredentialsOk

`func (o *HadoopConnectionParams) GetSshPasswordCredentialsOk() (*HadoopConnectionParamsSshPasswordCredentials, bool)`

GetSshPasswordCredentialsOk returns a tuple with the SshPasswordCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPasswordCredentials

`func (o *HadoopConnectionParams) SetSshPasswordCredentials(v HadoopConnectionParamsSshPasswordCredentials)`

SetSshPasswordCredentials sets SshPasswordCredentials field to given value.

### HasSshPasswordCredentials

`func (o *HadoopConnectionParams) HasSshPasswordCredentials() bool`

HasSshPasswordCredentials returns a boolean if a field has been set.

### SetSshPasswordCredentialsNil

`func (o *HadoopConnectionParams) SetSshPasswordCredentialsNil(b bool)`

 SetSshPasswordCredentialsNil sets the value for SshPasswordCredentials to be an explicit nil

### UnsetSshPasswordCredentials
`func (o *HadoopConnectionParams) UnsetSshPasswordCredentials()`

UnsetSshPasswordCredentials ensures that no value is present for SshPasswordCredentials, not even an explicit nil
### GetSshPrivateKeyCredentials

`func (o *HadoopConnectionParams) GetSshPrivateKeyCredentials() HadoopConnectionParamsSshPrivateKeyCredentials`

GetSshPrivateKeyCredentials returns the SshPrivateKeyCredentials field if non-nil, zero value otherwise.

### GetSshPrivateKeyCredentialsOk

`func (o *HadoopConnectionParams) GetSshPrivateKeyCredentialsOk() (*HadoopConnectionParamsSshPrivateKeyCredentials, bool)`

GetSshPrivateKeyCredentialsOk returns a tuple with the SshPrivateKeyCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPrivateKeyCredentials

`func (o *HadoopConnectionParams) SetSshPrivateKeyCredentials(v HadoopConnectionParamsSshPrivateKeyCredentials)`

SetSshPrivateKeyCredentials sets SshPrivateKeyCredentials field to given value.

### HasSshPrivateKeyCredentials

`func (o *HadoopConnectionParams) HasSshPrivateKeyCredentials() bool`

HasSshPrivateKeyCredentials returns a boolean if a field has been set.

### SetSshPrivateKeyCredentialsNil

`func (o *HadoopConnectionParams) SetSshPrivateKeyCredentialsNil(b bool)`

 SetSshPrivateKeyCredentialsNil sets the value for SshPrivateKeyCredentials to be an explicit nil

### UnsetSshPrivateKeyCredentials
`func (o *HadoopConnectionParams) UnsetSshPrivateKeyCredentials()`

UnsetSshPrivateKeyCredentials ensures that no value is present for SshPrivateKeyCredentials, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


