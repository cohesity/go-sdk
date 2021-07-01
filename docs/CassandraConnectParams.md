# CassandraConnectParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CassandraPortsInfo** | Pointer to [**CassandraPortsInfo**](CassandraPortsInfo.md) |  | [optional] 
**CassandraSecurityInfo** | Pointer to [**CassandraSecurityInfo**](CassandraSecurityInfo.md) |  | [optional] 
**CassandraVersion** | Pointer to **NullableString** | Cassandra version | [optional] 
**ConfigDirectory** | Pointer to **NullableString** | Specifies the Directory path containing Config YAML for discovery. | [optional] 
**DataCenters** | Pointer to **[]string** | Specifies the List of all physical data center or virtual data center. In most cases, the data centers will be listed after discovery operation however, if they are not listed, you must manually type the data center names. Leaving the field blank will disallow data center-specific backup or restore. Entering a subset of all data centers may cause problems in data movement. | [optional] 
**DseConfigDirectory** | Pointer to **NullableString** | Specifies the Directory from where DSE specific configuration can be read. | [optional] 
**IsDseAuthenticator** | Pointer to **NullableBool** | Specifies whether this cluster has DSE Authenticator. | [optional] 
**IsDseTieredStorage** | Pointer to **NullableBool** | Specifies whether this cluster has DSE tiered storage. | [optional] 
**IsJmxAuthEnable** | Pointer to **NullableBool** | Specifies if JMX Authentication enabled in this cluster. | [optional] 
**KerberosPrincipal** | Pointer to **NullableString** | Specifies the Kerberos Principal for Kerberos connection | [optional] 
**PrimaryHost** | Pointer to **NullableString** | Specifies the Primary Host for the Cassandra cluster. | [optional] 
**Seeds** | Pointer to **[]string** | Specifies the Seed nodes of this Cassandra cluster. | [optional] 
**SolrNodes** | Pointer to **[]string** | Specifies the Solr node IP Addresses | [optional] 
**SolrPort** | Pointer to **NullableInt32** | Specifies the Solr node Port. | [optional] 

## Methods

### NewCassandraConnectParams

`func NewCassandraConnectParams() *CassandraConnectParams`

NewCassandraConnectParams instantiates a new CassandraConnectParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCassandraConnectParamsWithDefaults

`func NewCassandraConnectParamsWithDefaults() *CassandraConnectParams`

NewCassandraConnectParamsWithDefaults instantiates a new CassandraConnectParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCassandraPortsInfo

`func (o *CassandraConnectParams) GetCassandraPortsInfo() CassandraPortsInfo`

GetCassandraPortsInfo returns the CassandraPortsInfo field if non-nil, zero value otherwise.

### GetCassandraPortsInfoOk

`func (o *CassandraConnectParams) GetCassandraPortsInfoOk() (*CassandraPortsInfo, bool)`

GetCassandraPortsInfoOk returns a tuple with the CassandraPortsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraPortsInfo

`func (o *CassandraConnectParams) SetCassandraPortsInfo(v CassandraPortsInfo)`

SetCassandraPortsInfo sets CassandraPortsInfo field to given value.

### HasCassandraPortsInfo

`func (o *CassandraConnectParams) HasCassandraPortsInfo() bool`

HasCassandraPortsInfo returns a boolean if a field has been set.

### GetCassandraSecurityInfo

`func (o *CassandraConnectParams) GetCassandraSecurityInfo() CassandraSecurityInfo`

GetCassandraSecurityInfo returns the CassandraSecurityInfo field if non-nil, zero value otherwise.

### GetCassandraSecurityInfoOk

`func (o *CassandraConnectParams) GetCassandraSecurityInfoOk() (*CassandraSecurityInfo, bool)`

GetCassandraSecurityInfoOk returns a tuple with the CassandraSecurityInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraSecurityInfo

`func (o *CassandraConnectParams) SetCassandraSecurityInfo(v CassandraSecurityInfo)`

SetCassandraSecurityInfo sets CassandraSecurityInfo field to given value.

### HasCassandraSecurityInfo

`func (o *CassandraConnectParams) HasCassandraSecurityInfo() bool`

HasCassandraSecurityInfo returns a boolean if a field has been set.

### GetCassandraVersion

`func (o *CassandraConnectParams) GetCassandraVersion() string`

GetCassandraVersion returns the CassandraVersion field if non-nil, zero value otherwise.

### GetCassandraVersionOk

`func (o *CassandraConnectParams) GetCassandraVersionOk() (*string, bool)`

GetCassandraVersionOk returns a tuple with the CassandraVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraVersion

`func (o *CassandraConnectParams) SetCassandraVersion(v string)`

SetCassandraVersion sets CassandraVersion field to given value.

### HasCassandraVersion

`func (o *CassandraConnectParams) HasCassandraVersion() bool`

HasCassandraVersion returns a boolean if a field has been set.

### SetCassandraVersionNil

`func (o *CassandraConnectParams) SetCassandraVersionNil(b bool)`

 SetCassandraVersionNil sets the value for CassandraVersion to be an explicit nil

### UnsetCassandraVersion
`func (o *CassandraConnectParams) UnsetCassandraVersion()`

UnsetCassandraVersion ensures that no value is present for CassandraVersion, not even an explicit nil
### GetConfigDirectory

`func (o *CassandraConnectParams) GetConfigDirectory() string`

GetConfigDirectory returns the ConfigDirectory field if non-nil, zero value otherwise.

### GetConfigDirectoryOk

`func (o *CassandraConnectParams) GetConfigDirectoryOk() (*string, bool)`

GetConfigDirectoryOk returns a tuple with the ConfigDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigDirectory

`func (o *CassandraConnectParams) SetConfigDirectory(v string)`

SetConfigDirectory sets ConfigDirectory field to given value.

### HasConfigDirectory

`func (o *CassandraConnectParams) HasConfigDirectory() bool`

HasConfigDirectory returns a boolean if a field has been set.

### SetConfigDirectoryNil

`func (o *CassandraConnectParams) SetConfigDirectoryNil(b bool)`

 SetConfigDirectoryNil sets the value for ConfigDirectory to be an explicit nil

### UnsetConfigDirectory
`func (o *CassandraConnectParams) UnsetConfigDirectory()`

UnsetConfigDirectory ensures that no value is present for ConfigDirectory, not even an explicit nil
### GetDataCenters

`func (o *CassandraConnectParams) GetDataCenters() []string`

GetDataCenters returns the DataCenters field if non-nil, zero value otherwise.

### GetDataCentersOk

`func (o *CassandraConnectParams) GetDataCentersOk() (*[]string, bool)`

GetDataCentersOk returns a tuple with the DataCenters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenters

`func (o *CassandraConnectParams) SetDataCenters(v []string)`

SetDataCenters sets DataCenters field to given value.

### HasDataCenters

`func (o *CassandraConnectParams) HasDataCenters() bool`

HasDataCenters returns a boolean if a field has been set.

### SetDataCentersNil

`func (o *CassandraConnectParams) SetDataCentersNil(b bool)`

 SetDataCentersNil sets the value for DataCenters to be an explicit nil

### UnsetDataCenters
`func (o *CassandraConnectParams) UnsetDataCenters()`

UnsetDataCenters ensures that no value is present for DataCenters, not even an explicit nil
### GetDseConfigDirectory

`func (o *CassandraConnectParams) GetDseConfigDirectory() string`

GetDseConfigDirectory returns the DseConfigDirectory field if non-nil, zero value otherwise.

### GetDseConfigDirectoryOk

`func (o *CassandraConnectParams) GetDseConfigDirectoryOk() (*string, bool)`

GetDseConfigDirectoryOk returns a tuple with the DseConfigDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDseConfigDirectory

`func (o *CassandraConnectParams) SetDseConfigDirectory(v string)`

SetDseConfigDirectory sets DseConfigDirectory field to given value.

### HasDseConfigDirectory

`func (o *CassandraConnectParams) HasDseConfigDirectory() bool`

HasDseConfigDirectory returns a boolean if a field has been set.

### SetDseConfigDirectoryNil

`func (o *CassandraConnectParams) SetDseConfigDirectoryNil(b bool)`

 SetDseConfigDirectoryNil sets the value for DseConfigDirectory to be an explicit nil

### UnsetDseConfigDirectory
`func (o *CassandraConnectParams) UnsetDseConfigDirectory()`

UnsetDseConfigDirectory ensures that no value is present for DseConfigDirectory, not even an explicit nil
### GetIsDseAuthenticator

`func (o *CassandraConnectParams) GetIsDseAuthenticator() bool`

GetIsDseAuthenticator returns the IsDseAuthenticator field if non-nil, zero value otherwise.

### GetIsDseAuthenticatorOk

`func (o *CassandraConnectParams) GetIsDseAuthenticatorOk() (*bool, bool)`

GetIsDseAuthenticatorOk returns a tuple with the IsDseAuthenticator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDseAuthenticator

`func (o *CassandraConnectParams) SetIsDseAuthenticator(v bool)`

SetIsDseAuthenticator sets IsDseAuthenticator field to given value.

### HasIsDseAuthenticator

`func (o *CassandraConnectParams) HasIsDseAuthenticator() bool`

HasIsDseAuthenticator returns a boolean if a field has been set.

### SetIsDseAuthenticatorNil

`func (o *CassandraConnectParams) SetIsDseAuthenticatorNil(b bool)`

 SetIsDseAuthenticatorNil sets the value for IsDseAuthenticator to be an explicit nil

### UnsetIsDseAuthenticator
`func (o *CassandraConnectParams) UnsetIsDseAuthenticator()`

UnsetIsDseAuthenticator ensures that no value is present for IsDseAuthenticator, not even an explicit nil
### GetIsDseTieredStorage

`func (o *CassandraConnectParams) GetIsDseTieredStorage() bool`

GetIsDseTieredStorage returns the IsDseTieredStorage field if non-nil, zero value otherwise.

### GetIsDseTieredStorageOk

`func (o *CassandraConnectParams) GetIsDseTieredStorageOk() (*bool, bool)`

GetIsDseTieredStorageOk returns a tuple with the IsDseTieredStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDseTieredStorage

`func (o *CassandraConnectParams) SetIsDseTieredStorage(v bool)`

SetIsDseTieredStorage sets IsDseTieredStorage field to given value.

### HasIsDseTieredStorage

`func (o *CassandraConnectParams) HasIsDseTieredStorage() bool`

HasIsDseTieredStorage returns a boolean if a field has been set.

### SetIsDseTieredStorageNil

`func (o *CassandraConnectParams) SetIsDseTieredStorageNil(b bool)`

 SetIsDseTieredStorageNil sets the value for IsDseTieredStorage to be an explicit nil

### UnsetIsDseTieredStorage
`func (o *CassandraConnectParams) UnsetIsDseTieredStorage()`

UnsetIsDseTieredStorage ensures that no value is present for IsDseTieredStorage, not even an explicit nil
### GetIsJmxAuthEnable

`func (o *CassandraConnectParams) GetIsJmxAuthEnable() bool`

GetIsJmxAuthEnable returns the IsJmxAuthEnable field if non-nil, zero value otherwise.

### GetIsJmxAuthEnableOk

`func (o *CassandraConnectParams) GetIsJmxAuthEnableOk() (*bool, bool)`

GetIsJmxAuthEnableOk returns a tuple with the IsJmxAuthEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsJmxAuthEnable

`func (o *CassandraConnectParams) SetIsJmxAuthEnable(v bool)`

SetIsJmxAuthEnable sets IsJmxAuthEnable field to given value.

### HasIsJmxAuthEnable

`func (o *CassandraConnectParams) HasIsJmxAuthEnable() bool`

HasIsJmxAuthEnable returns a boolean if a field has been set.

### SetIsJmxAuthEnableNil

`func (o *CassandraConnectParams) SetIsJmxAuthEnableNil(b bool)`

 SetIsJmxAuthEnableNil sets the value for IsJmxAuthEnable to be an explicit nil

### UnsetIsJmxAuthEnable
`func (o *CassandraConnectParams) UnsetIsJmxAuthEnable()`

UnsetIsJmxAuthEnable ensures that no value is present for IsJmxAuthEnable, not even an explicit nil
### GetKerberosPrincipal

`func (o *CassandraConnectParams) GetKerberosPrincipal() string`

GetKerberosPrincipal returns the KerberosPrincipal field if non-nil, zero value otherwise.

### GetKerberosPrincipalOk

`func (o *CassandraConnectParams) GetKerberosPrincipalOk() (*string, bool)`

GetKerberosPrincipalOk returns a tuple with the KerberosPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosPrincipal

`func (o *CassandraConnectParams) SetKerberosPrincipal(v string)`

SetKerberosPrincipal sets KerberosPrincipal field to given value.

### HasKerberosPrincipal

`func (o *CassandraConnectParams) HasKerberosPrincipal() bool`

HasKerberosPrincipal returns a boolean if a field has been set.

### SetKerberosPrincipalNil

`func (o *CassandraConnectParams) SetKerberosPrincipalNil(b bool)`

 SetKerberosPrincipalNil sets the value for KerberosPrincipal to be an explicit nil

### UnsetKerberosPrincipal
`func (o *CassandraConnectParams) UnsetKerberosPrincipal()`

UnsetKerberosPrincipal ensures that no value is present for KerberosPrincipal, not even an explicit nil
### GetPrimaryHost

`func (o *CassandraConnectParams) GetPrimaryHost() string`

GetPrimaryHost returns the PrimaryHost field if non-nil, zero value otherwise.

### GetPrimaryHostOk

`func (o *CassandraConnectParams) GetPrimaryHostOk() (*string, bool)`

GetPrimaryHostOk returns a tuple with the PrimaryHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryHost

`func (o *CassandraConnectParams) SetPrimaryHost(v string)`

SetPrimaryHost sets PrimaryHost field to given value.

### HasPrimaryHost

`func (o *CassandraConnectParams) HasPrimaryHost() bool`

HasPrimaryHost returns a boolean if a field has been set.

### SetPrimaryHostNil

`func (o *CassandraConnectParams) SetPrimaryHostNil(b bool)`

 SetPrimaryHostNil sets the value for PrimaryHost to be an explicit nil

### UnsetPrimaryHost
`func (o *CassandraConnectParams) UnsetPrimaryHost()`

UnsetPrimaryHost ensures that no value is present for PrimaryHost, not even an explicit nil
### GetSeeds

`func (o *CassandraConnectParams) GetSeeds() []string`

GetSeeds returns the Seeds field if non-nil, zero value otherwise.

### GetSeedsOk

`func (o *CassandraConnectParams) GetSeedsOk() (*[]string, bool)`

GetSeedsOk returns a tuple with the Seeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeeds

`func (o *CassandraConnectParams) SetSeeds(v []string)`

SetSeeds sets Seeds field to given value.

### HasSeeds

`func (o *CassandraConnectParams) HasSeeds() bool`

HasSeeds returns a boolean if a field has been set.

### SetSeedsNil

`func (o *CassandraConnectParams) SetSeedsNil(b bool)`

 SetSeedsNil sets the value for Seeds to be an explicit nil

### UnsetSeeds
`func (o *CassandraConnectParams) UnsetSeeds()`

UnsetSeeds ensures that no value is present for Seeds, not even an explicit nil
### GetSolrNodes

`func (o *CassandraConnectParams) GetSolrNodes() []string`

GetSolrNodes returns the SolrNodes field if non-nil, zero value otherwise.

### GetSolrNodesOk

`func (o *CassandraConnectParams) GetSolrNodesOk() (*[]string, bool)`

GetSolrNodesOk returns a tuple with the SolrNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolrNodes

`func (o *CassandraConnectParams) SetSolrNodes(v []string)`

SetSolrNodes sets SolrNodes field to given value.

### HasSolrNodes

`func (o *CassandraConnectParams) HasSolrNodes() bool`

HasSolrNodes returns a boolean if a field has been set.

### SetSolrNodesNil

`func (o *CassandraConnectParams) SetSolrNodesNil(b bool)`

 SetSolrNodesNil sets the value for SolrNodes to be an explicit nil

### UnsetSolrNodes
`func (o *CassandraConnectParams) UnsetSolrNodes()`

UnsetSolrNodes ensures that no value is present for SolrNodes, not even an explicit nil
### GetSolrPort

`func (o *CassandraConnectParams) GetSolrPort() int32`

GetSolrPort returns the SolrPort field if non-nil, zero value otherwise.

### GetSolrPortOk

`func (o *CassandraConnectParams) GetSolrPortOk() (*int32, bool)`

GetSolrPortOk returns a tuple with the SolrPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolrPort

`func (o *CassandraConnectParams) SetSolrPort(v int32)`

SetSolrPort sets SolrPort field to given value.

### HasSolrPort

`func (o *CassandraConnectParams) HasSolrPort() bool`

HasSolrPort returns a boolean if a field has been set.

### SetSolrPortNil

`func (o *CassandraConnectParams) SetSolrPortNil(b bool)`

 SetSolrPortNil sets the value for SolrPort to be an explicit nil

### UnsetSolrPort
`func (o *CassandraConnectParams) UnsetSolrPort()`

UnsetSolrPort ensures that no value is present for SolrPort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


