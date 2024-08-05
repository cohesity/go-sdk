# CassandraSourceConfigParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CassandraPartitioner** | Pointer to **NullableString** | Cassandra partitioner required in compaction. | [optional] 
**CassandraPortInfo** | Pointer to [**CassandraPortInfo**](CassandraPortInfo.md) |  | [optional] 
**CassandraSecurityInfo** | Pointer to [**CassandraSecurityInfo**](CassandraSecurityInfo.md) |  | [optional] 
**CassandraVersion** | Pointer to **NullableString** | Cassandra Version. | [optional] 
**CommitLogBackupLocation** | Pointer to **NullableString** | Commit Logs backup location on cassandra nodes | [optional] 
**DataCenterNames** | Pointer to **[]string** | Data centers for this cluster. | [optional] 
**DseVersion** | Pointer to **NullableString** | DSE Version | [optional] 
**EndpointSnitch** | Pointer to **NullableString** | Endpoint snitch used for this cluster. | [optional] 
**IsJmxAuthEnable** | Pointer to **NullableBool** | Is JMX Authentication enabled in this cluster ? | [optional] 
**KerberosSaslProtocol** | Pointer to **NullableString** | Populated if cassandraAuthType is Kerberos. | [optional] 
**Seeds** | Pointer to **[]string** | Seed nodes of this cluster. | [optional] 

## Methods

### NewCassandraSourceConfigParams

`func NewCassandraSourceConfigParams() *CassandraSourceConfigParams`

NewCassandraSourceConfigParams instantiates a new CassandraSourceConfigParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCassandraSourceConfigParamsWithDefaults

`func NewCassandraSourceConfigParamsWithDefaults() *CassandraSourceConfigParams`

NewCassandraSourceConfigParamsWithDefaults instantiates a new CassandraSourceConfigParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCassandraPartitioner

`func (o *CassandraSourceConfigParams) GetCassandraPartitioner() string`

GetCassandraPartitioner returns the CassandraPartitioner field if non-nil, zero value otherwise.

### GetCassandraPartitionerOk

`func (o *CassandraSourceConfigParams) GetCassandraPartitionerOk() (*string, bool)`

GetCassandraPartitionerOk returns a tuple with the CassandraPartitioner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraPartitioner

`func (o *CassandraSourceConfigParams) SetCassandraPartitioner(v string)`

SetCassandraPartitioner sets CassandraPartitioner field to given value.

### HasCassandraPartitioner

`func (o *CassandraSourceConfigParams) HasCassandraPartitioner() bool`

HasCassandraPartitioner returns a boolean if a field has been set.

### SetCassandraPartitionerNil

`func (o *CassandraSourceConfigParams) SetCassandraPartitionerNil(b bool)`

 SetCassandraPartitionerNil sets the value for CassandraPartitioner to be an explicit nil

### UnsetCassandraPartitioner
`func (o *CassandraSourceConfigParams) UnsetCassandraPartitioner()`

UnsetCassandraPartitioner ensures that no value is present for CassandraPartitioner, not even an explicit nil
### GetCassandraPortInfo

`func (o *CassandraSourceConfigParams) GetCassandraPortInfo() CassandraPortInfo`

GetCassandraPortInfo returns the CassandraPortInfo field if non-nil, zero value otherwise.

### GetCassandraPortInfoOk

`func (o *CassandraSourceConfigParams) GetCassandraPortInfoOk() (*CassandraPortInfo, bool)`

GetCassandraPortInfoOk returns a tuple with the CassandraPortInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraPortInfo

`func (o *CassandraSourceConfigParams) SetCassandraPortInfo(v CassandraPortInfo)`

SetCassandraPortInfo sets CassandraPortInfo field to given value.

### HasCassandraPortInfo

`func (o *CassandraSourceConfigParams) HasCassandraPortInfo() bool`

HasCassandraPortInfo returns a boolean if a field has been set.

### GetCassandraSecurityInfo

`func (o *CassandraSourceConfigParams) GetCassandraSecurityInfo() CassandraSecurityInfo`

GetCassandraSecurityInfo returns the CassandraSecurityInfo field if non-nil, zero value otherwise.

### GetCassandraSecurityInfoOk

`func (o *CassandraSourceConfigParams) GetCassandraSecurityInfoOk() (*CassandraSecurityInfo, bool)`

GetCassandraSecurityInfoOk returns a tuple with the CassandraSecurityInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraSecurityInfo

`func (o *CassandraSourceConfigParams) SetCassandraSecurityInfo(v CassandraSecurityInfo)`

SetCassandraSecurityInfo sets CassandraSecurityInfo field to given value.

### HasCassandraSecurityInfo

`func (o *CassandraSourceConfigParams) HasCassandraSecurityInfo() bool`

HasCassandraSecurityInfo returns a boolean if a field has been set.

### GetCassandraVersion

`func (o *CassandraSourceConfigParams) GetCassandraVersion() string`

GetCassandraVersion returns the CassandraVersion field if non-nil, zero value otherwise.

### GetCassandraVersionOk

`func (o *CassandraSourceConfigParams) GetCassandraVersionOk() (*string, bool)`

GetCassandraVersionOk returns a tuple with the CassandraVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraVersion

`func (o *CassandraSourceConfigParams) SetCassandraVersion(v string)`

SetCassandraVersion sets CassandraVersion field to given value.

### HasCassandraVersion

`func (o *CassandraSourceConfigParams) HasCassandraVersion() bool`

HasCassandraVersion returns a boolean if a field has been set.

### SetCassandraVersionNil

`func (o *CassandraSourceConfigParams) SetCassandraVersionNil(b bool)`

 SetCassandraVersionNil sets the value for CassandraVersion to be an explicit nil

### UnsetCassandraVersion
`func (o *CassandraSourceConfigParams) UnsetCassandraVersion()`

UnsetCassandraVersion ensures that no value is present for CassandraVersion, not even an explicit nil
### GetCommitLogBackupLocation

`func (o *CassandraSourceConfigParams) GetCommitLogBackupLocation() string`

GetCommitLogBackupLocation returns the CommitLogBackupLocation field if non-nil, zero value otherwise.

### GetCommitLogBackupLocationOk

`func (o *CassandraSourceConfigParams) GetCommitLogBackupLocationOk() (*string, bool)`

GetCommitLogBackupLocationOk returns a tuple with the CommitLogBackupLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitLogBackupLocation

`func (o *CassandraSourceConfigParams) SetCommitLogBackupLocation(v string)`

SetCommitLogBackupLocation sets CommitLogBackupLocation field to given value.

### HasCommitLogBackupLocation

`func (o *CassandraSourceConfigParams) HasCommitLogBackupLocation() bool`

HasCommitLogBackupLocation returns a boolean if a field has been set.

### SetCommitLogBackupLocationNil

`func (o *CassandraSourceConfigParams) SetCommitLogBackupLocationNil(b bool)`

 SetCommitLogBackupLocationNil sets the value for CommitLogBackupLocation to be an explicit nil

### UnsetCommitLogBackupLocation
`func (o *CassandraSourceConfigParams) UnsetCommitLogBackupLocation()`

UnsetCommitLogBackupLocation ensures that no value is present for CommitLogBackupLocation, not even an explicit nil
### GetDataCenterNames

`func (o *CassandraSourceConfigParams) GetDataCenterNames() []string`

GetDataCenterNames returns the DataCenterNames field if non-nil, zero value otherwise.

### GetDataCenterNamesOk

`func (o *CassandraSourceConfigParams) GetDataCenterNamesOk() (*[]string, bool)`

GetDataCenterNamesOk returns a tuple with the DataCenterNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenterNames

`func (o *CassandraSourceConfigParams) SetDataCenterNames(v []string)`

SetDataCenterNames sets DataCenterNames field to given value.

### HasDataCenterNames

`func (o *CassandraSourceConfigParams) HasDataCenterNames() bool`

HasDataCenterNames returns a boolean if a field has been set.

### GetDseVersion

`func (o *CassandraSourceConfigParams) GetDseVersion() string`

GetDseVersion returns the DseVersion field if non-nil, zero value otherwise.

### GetDseVersionOk

`func (o *CassandraSourceConfigParams) GetDseVersionOk() (*string, bool)`

GetDseVersionOk returns a tuple with the DseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDseVersion

`func (o *CassandraSourceConfigParams) SetDseVersion(v string)`

SetDseVersion sets DseVersion field to given value.

### HasDseVersion

`func (o *CassandraSourceConfigParams) HasDseVersion() bool`

HasDseVersion returns a boolean if a field has been set.

### SetDseVersionNil

`func (o *CassandraSourceConfigParams) SetDseVersionNil(b bool)`

 SetDseVersionNil sets the value for DseVersion to be an explicit nil

### UnsetDseVersion
`func (o *CassandraSourceConfigParams) UnsetDseVersion()`

UnsetDseVersion ensures that no value is present for DseVersion, not even an explicit nil
### GetEndpointSnitch

`func (o *CassandraSourceConfigParams) GetEndpointSnitch() string`

GetEndpointSnitch returns the EndpointSnitch field if non-nil, zero value otherwise.

### GetEndpointSnitchOk

`func (o *CassandraSourceConfigParams) GetEndpointSnitchOk() (*string, bool)`

GetEndpointSnitchOk returns a tuple with the EndpointSnitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointSnitch

`func (o *CassandraSourceConfigParams) SetEndpointSnitch(v string)`

SetEndpointSnitch sets EndpointSnitch field to given value.

### HasEndpointSnitch

`func (o *CassandraSourceConfigParams) HasEndpointSnitch() bool`

HasEndpointSnitch returns a boolean if a field has been set.

### SetEndpointSnitchNil

`func (o *CassandraSourceConfigParams) SetEndpointSnitchNil(b bool)`

 SetEndpointSnitchNil sets the value for EndpointSnitch to be an explicit nil

### UnsetEndpointSnitch
`func (o *CassandraSourceConfigParams) UnsetEndpointSnitch()`

UnsetEndpointSnitch ensures that no value is present for EndpointSnitch, not even an explicit nil
### GetIsJmxAuthEnable

`func (o *CassandraSourceConfigParams) GetIsJmxAuthEnable() bool`

GetIsJmxAuthEnable returns the IsJmxAuthEnable field if non-nil, zero value otherwise.

### GetIsJmxAuthEnableOk

`func (o *CassandraSourceConfigParams) GetIsJmxAuthEnableOk() (*bool, bool)`

GetIsJmxAuthEnableOk returns a tuple with the IsJmxAuthEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsJmxAuthEnable

`func (o *CassandraSourceConfigParams) SetIsJmxAuthEnable(v bool)`

SetIsJmxAuthEnable sets IsJmxAuthEnable field to given value.

### HasIsJmxAuthEnable

`func (o *CassandraSourceConfigParams) HasIsJmxAuthEnable() bool`

HasIsJmxAuthEnable returns a boolean if a field has been set.

### SetIsJmxAuthEnableNil

`func (o *CassandraSourceConfigParams) SetIsJmxAuthEnableNil(b bool)`

 SetIsJmxAuthEnableNil sets the value for IsJmxAuthEnable to be an explicit nil

### UnsetIsJmxAuthEnable
`func (o *CassandraSourceConfigParams) UnsetIsJmxAuthEnable()`

UnsetIsJmxAuthEnable ensures that no value is present for IsJmxAuthEnable, not even an explicit nil
### GetKerberosSaslProtocol

`func (o *CassandraSourceConfigParams) GetKerberosSaslProtocol() string`

GetKerberosSaslProtocol returns the KerberosSaslProtocol field if non-nil, zero value otherwise.

### GetKerberosSaslProtocolOk

`func (o *CassandraSourceConfigParams) GetKerberosSaslProtocolOk() (*string, bool)`

GetKerberosSaslProtocolOk returns a tuple with the KerberosSaslProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosSaslProtocol

`func (o *CassandraSourceConfigParams) SetKerberosSaslProtocol(v string)`

SetKerberosSaslProtocol sets KerberosSaslProtocol field to given value.

### HasKerberosSaslProtocol

`func (o *CassandraSourceConfigParams) HasKerberosSaslProtocol() bool`

HasKerberosSaslProtocol returns a boolean if a field has been set.

### SetKerberosSaslProtocolNil

`func (o *CassandraSourceConfigParams) SetKerberosSaslProtocolNil(b bool)`

 SetKerberosSaslProtocolNil sets the value for KerberosSaslProtocol to be an explicit nil

### UnsetKerberosSaslProtocol
`func (o *CassandraSourceConfigParams) UnsetKerberosSaslProtocol()`

UnsetKerberosSaslProtocol ensures that no value is present for KerberosSaslProtocol, not even an explicit nil
### GetSeeds

`func (o *CassandraSourceConfigParams) GetSeeds() []string`

GetSeeds returns the Seeds field if non-nil, zero value otherwise.

### GetSeedsOk

`func (o *CassandraSourceConfigParams) GetSeedsOk() (*[]string, bool)`

GetSeedsOk returns a tuple with the Seeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeeds

`func (o *CassandraSourceConfigParams) SetSeeds(v []string)`

SetSeeds sets Seeds field to given value.

### HasSeeds

`func (o *CassandraSourceConfigParams) HasSeeds() bool`

HasSeeds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


