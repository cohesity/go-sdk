# MssqlConnectionResponseParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostIdentifier** | **string** | Specifies the unique identifier to locate the SQL node or cluster. The host identifier can be IP address or FQDN. | 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 
**SkipConnectionDiscovery** | Pointer to **NullableBool** | Specifies whether to skip the discovery phase of all SQL servers, AAG groups etc during registration process. | [optional] 
**Servers** | Pointer to [**[]SQLServer**](SQLServer.md) | Specifies the list of SQL servers. If SQL server is a part of avalibility group then it will be returned in aagServers field. This will include the list of all standalone SQL servers and servers belonging to any FCI enviournment. | [optional] 
**FciClusters** | Pointer to [**[]FCICluster**](FCICluster.md) | Specifies the list of FCI (Failover Cluster Instaces) Clusters. This will contain the list of all failover pools under a windows cluster. FCI clusters which are part of AAG, will be returned seperatly under aagServers field. | [optional] 
**AagGroups** | Pointer to [**[]AAGGroup**](AAGGroup.md) | Specifies the list of AAG (Always on Avalibility) groups. | [optional] 

## Methods

### NewMssqlConnectionResponseParams

`func NewMssqlConnectionResponseParams(hostIdentifier string, ) *MssqlConnectionResponseParams`

NewMssqlConnectionResponseParams instantiates a new MssqlConnectionResponseParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMssqlConnectionResponseParamsWithDefaults

`func NewMssqlConnectionResponseParamsWithDefaults() *MssqlConnectionResponseParams`

NewMssqlConnectionResponseParamsWithDefaults instantiates a new MssqlConnectionResponseParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostIdentifier

`func (o *MssqlConnectionResponseParams) GetHostIdentifier() string`

GetHostIdentifier returns the HostIdentifier field if non-nil, zero value otherwise.

### GetHostIdentifierOk

`func (o *MssqlConnectionResponseParams) GetHostIdentifierOk() (*string, bool)`

GetHostIdentifierOk returns a tuple with the HostIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostIdentifier

`func (o *MssqlConnectionResponseParams) SetHostIdentifier(v string)`

SetHostIdentifier sets HostIdentifier field to given value.


### GetError

`func (o *MssqlConnectionResponseParams) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MssqlConnectionResponseParams) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MssqlConnectionResponseParams) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *MssqlConnectionResponseParams) HasError() bool`

HasError returns a boolean if a field has been set.

### GetSkipConnectionDiscovery

`func (o *MssqlConnectionResponseParams) GetSkipConnectionDiscovery() bool`

GetSkipConnectionDiscovery returns the SkipConnectionDiscovery field if non-nil, zero value otherwise.

### GetSkipConnectionDiscoveryOk

`func (o *MssqlConnectionResponseParams) GetSkipConnectionDiscoveryOk() (*bool, bool)`

GetSkipConnectionDiscoveryOk returns a tuple with the SkipConnectionDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipConnectionDiscovery

`func (o *MssqlConnectionResponseParams) SetSkipConnectionDiscovery(v bool)`

SetSkipConnectionDiscovery sets SkipConnectionDiscovery field to given value.

### HasSkipConnectionDiscovery

`func (o *MssqlConnectionResponseParams) HasSkipConnectionDiscovery() bool`

HasSkipConnectionDiscovery returns a boolean if a field has been set.

### SetSkipConnectionDiscoveryNil

`func (o *MssqlConnectionResponseParams) SetSkipConnectionDiscoveryNil(b bool)`

 SetSkipConnectionDiscoveryNil sets the value for SkipConnectionDiscovery to be an explicit nil

### UnsetSkipConnectionDiscovery
`func (o *MssqlConnectionResponseParams) UnsetSkipConnectionDiscovery()`

UnsetSkipConnectionDiscovery ensures that no value is present for SkipConnectionDiscovery, not even an explicit nil
### GetServers

`func (o *MssqlConnectionResponseParams) GetServers() []SQLServer`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *MssqlConnectionResponseParams) GetServersOk() (*[]SQLServer, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *MssqlConnectionResponseParams) SetServers(v []SQLServer)`

SetServers sets Servers field to given value.

### HasServers

`func (o *MssqlConnectionResponseParams) HasServers() bool`

HasServers returns a boolean if a field has been set.

### SetServersNil

`func (o *MssqlConnectionResponseParams) SetServersNil(b bool)`

 SetServersNil sets the value for Servers to be an explicit nil

### UnsetServers
`func (o *MssqlConnectionResponseParams) UnsetServers()`

UnsetServers ensures that no value is present for Servers, not even an explicit nil
### GetFciClusters

`func (o *MssqlConnectionResponseParams) GetFciClusters() []FCICluster`

GetFciClusters returns the FciClusters field if non-nil, zero value otherwise.

### GetFciClustersOk

`func (o *MssqlConnectionResponseParams) GetFciClustersOk() (*[]FCICluster, bool)`

GetFciClustersOk returns a tuple with the FciClusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFciClusters

`func (o *MssqlConnectionResponseParams) SetFciClusters(v []FCICluster)`

SetFciClusters sets FciClusters field to given value.

### HasFciClusters

`func (o *MssqlConnectionResponseParams) HasFciClusters() bool`

HasFciClusters returns a boolean if a field has been set.

### SetFciClustersNil

`func (o *MssqlConnectionResponseParams) SetFciClustersNil(b bool)`

 SetFciClustersNil sets the value for FciClusters to be an explicit nil

### UnsetFciClusters
`func (o *MssqlConnectionResponseParams) UnsetFciClusters()`

UnsetFciClusters ensures that no value is present for FciClusters, not even an explicit nil
### GetAagGroups

`func (o *MssqlConnectionResponseParams) GetAagGroups() []AAGGroup`

GetAagGroups returns the AagGroups field if non-nil, zero value otherwise.

### GetAagGroupsOk

`func (o *MssqlConnectionResponseParams) GetAagGroupsOk() (*[]AAGGroup, bool)`

GetAagGroupsOk returns a tuple with the AagGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAagGroups

`func (o *MssqlConnectionResponseParams) SetAagGroups(v []AAGGroup)`

SetAagGroups sets AagGroups field to given value.

### HasAagGroups

`func (o *MssqlConnectionResponseParams) HasAagGroups() bool`

HasAagGroups returns a boolean if a field has been set.

### SetAagGroupsNil

`func (o *MssqlConnectionResponseParams) SetAagGroupsNil(b bool)`

 SetAagGroupsNil sets the value for AagGroups to be an explicit nil

### UnsetAagGroups
`func (o *MssqlConnectionResponseParams) UnsetAagGroups()`

UnsetAagGroups ensures that no value is present for AagGroups, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


