# AwsTargetParamsForRecoverRDSPostgresCustomServerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | **string** | Specifies the Ip in which to deploy the Rds objects. | 
**Port** | **NullableInt32** | Specifies the port to use to connect to the server. | 
**Region** | [**NullableRecoverRDSPostgresCustomServerConfigRegion**](RecoverRDSPostgresCustomServerConfigRegion.md) |  | 
**Source** | Pointer to [**NullableRecoverRDSPostgresCustomServerConfigSource**](RecoverRDSPostgresCustomServerConfigSource.md) |  | [optional] 
**StandardCredentials** | [**Credentials**](Credentials.md) |  | 

## Methods

### NewAwsTargetParamsForRecoverRDSPostgresCustomServerConfig

`func NewAwsTargetParamsForRecoverRDSPostgresCustomServerConfig(ip string, port NullableInt32, region NullableRecoverRDSPostgresCustomServerConfigRegion, standardCredentials Credentials, ) *AwsTargetParamsForRecoverRDSPostgresCustomServerConfig`

NewAwsTargetParamsForRecoverRDSPostgresCustomServerConfig instantiates a new AwsTargetParamsForRecoverRDSPostgresCustomServerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsTargetParamsForRecoverRDSPostgresCustomServerConfigWithDefaults

`func NewAwsTargetParamsForRecoverRDSPostgresCustomServerConfigWithDefaults() *AwsTargetParamsForRecoverRDSPostgresCustomServerConfig`

NewAwsTargetParamsForRecoverRDSPostgresCustomServerConfigWithDefaults instantiates a new AwsTargetParamsForRecoverRDSPostgresCustomServerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *AwsTargetParamsForRecoverRDSPostgresCustomServerConfig) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *AwsTargetParamsForRecoverRDSPostgresCustomServerConfig) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *AwsTargetParamsForRecoverRDSPostgresCustomServerConfig) SetIp(v string)`

SetIp sets Ip field to given value.


### GetPort

`func (o *AwsTargetParamsForRecoverRDSPostgresCustomServerConfig) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *AwsTargetParamsForRecoverRDSPostgresCustomServerConfig) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *AwsTargetParamsForRecoverRDSPostgresCustomServerConfig) SetPort(v int32)`

SetPort sets Port field to given value.


### SetPortNil

`func (o *AwsTargetParamsForRecoverRDSPostgresCustomServerConfig) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *AwsTargetParamsForRecoverRDSPostgresCustomServerConfig) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetRegion

`func (o *AwsTargetParamsForRecoverRDSPostgresCustomServerConfig) GetRegion() RecoverRDSPostgresCustomServerConfigRegion`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AwsTargetParamsForRecoverRDSPostgresCustomServerConfig) GetRegionOk() (*RecoverRDSPostgresCustomServerConfigRegion, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AwsTargetParamsForRecoverRDSPostgresCustomServerConfig) SetRegion(v RecoverRDSPostgresCustomServerConfigRegion)`

SetRegion sets Region field to given value.


### SetRegionNil

`func (o *AwsTargetParamsForRecoverRDSPostgresCustomServerConfig) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *AwsTargetParamsForRecoverRDSPostgresCustomServerConfig) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetSource

`func (o *AwsTargetParamsForRecoverRDSPostgresCustomServerConfig) GetSource() RecoverRDSPostgresCustomServerConfigSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AwsTargetParamsForRecoverRDSPostgresCustomServerConfig) GetSourceOk() (*RecoverRDSPostgresCustomServerConfigSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AwsTargetParamsForRecoverRDSPostgresCustomServerConfig) SetSource(v RecoverRDSPostgresCustomServerConfigSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *AwsTargetParamsForRecoverRDSPostgresCustomServerConfig) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *AwsTargetParamsForRecoverRDSPostgresCustomServerConfig) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *AwsTargetParamsForRecoverRDSPostgresCustomServerConfig) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetStandardCredentials

`func (o *AwsTargetParamsForRecoverRDSPostgresCustomServerConfig) GetStandardCredentials() Credentials`

GetStandardCredentials returns the StandardCredentials field if non-nil, zero value otherwise.

### GetStandardCredentialsOk

`func (o *AwsTargetParamsForRecoverRDSPostgresCustomServerConfig) GetStandardCredentialsOk() (*Credentials, bool)`

GetStandardCredentialsOk returns a tuple with the StandardCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardCredentials

`func (o *AwsTargetParamsForRecoverRDSPostgresCustomServerConfig) SetStandardCredentials(v Credentials)`

SetStandardCredentials sets StandardCredentials field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


