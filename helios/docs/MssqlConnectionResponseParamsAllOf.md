# MssqlConnectionResponseParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 
**SkipConnectionDiscovery** | Pointer to **NullableBool** | Specifies whether to skip the discovery phase of all SQL servers, AAG groups etc during registration process. | [optional] 
**Servers** | Pointer to [**[]SQLServer**](SQLServer.md) | Specifies the list of SQL servers. If SQL server is a part of avalibility group then it will be returned in aagServers field. This will include the list of all standalone SQL servers and servers belonging to any FCI enviournment. | [optional] 
**FciClusters** | Pointer to [**[]FCICluster**](FCICluster.md) | Specifies the list of FCI (Failover Cluster Instaces) Clusters. This will contain the list of all failover pools under a windows cluster. FCI clusters which are part of AAG, will be returned seperatly under aagServers field. | [optional] 
**AagGroups** | Pointer to [**[]AAGGroup**](AAGGroup.md) | Specifies the list of AAG (Always on Avalibility) groups. | [optional] 

## Methods

### NewMssqlConnectionResponseParamsAllOf

`func NewMssqlConnectionResponseParamsAllOf() *MssqlConnectionResponseParamsAllOf`

NewMssqlConnectionResponseParamsAllOf instantiates a new MssqlConnectionResponseParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMssqlConnectionResponseParamsAllOfWithDefaults

`func NewMssqlConnectionResponseParamsAllOfWithDefaults() *MssqlConnectionResponseParamsAllOf`

NewMssqlConnectionResponseParamsAllOfWithDefaults instantiates a new MssqlConnectionResponseParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *MssqlConnectionResponseParamsAllOf) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MssqlConnectionResponseParamsAllOf) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MssqlConnectionResponseParamsAllOf) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *MssqlConnectionResponseParamsAllOf) HasError() bool`

HasError returns a boolean if a field has been set.

### GetSkipConnectionDiscovery

`func (o *MssqlConnectionResponseParamsAllOf) GetSkipConnectionDiscovery() bool`

GetSkipConnectionDiscovery returns the SkipConnectionDiscovery field if non-nil, zero value otherwise.

### GetSkipConnectionDiscoveryOk

`func (o *MssqlConnectionResponseParamsAllOf) GetSkipConnectionDiscoveryOk() (*bool, bool)`

GetSkipConnectionDiscoveryOk returns a tuple with the SkipConnectionDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipConnectionDiscovery

`func (o *MssqlConnectionResponseParamsAllOf) SetSkipConnectionDiscovery(v bool)`

SetSkipConnectionDiscovery sets SkipConnectionDiscovery field to given value.

### HasSkipConnectionDiscovery

`func (o *MssqlConnectionResponseParamsAllOf) HasSkipConnectionDiscovery() bool`

HasSkipConnectionDiscovery returns a boolean if a field has been set.

### SetSkipConnectionDiscoveryNil

`func (o *MssqlConnectionResponseParamsAllOf) SetSkipConnectionDiscoveryNil(b bool)`

 SetSkipConnectionDiscoveryNil sets the value for SkipConnectionDiscovery to be an explicit nil

### UnsetSkipConnectionDiscovery
`func (o *MssqlConnectionResponseParamsAllOf) UnsetSkipConnectionDiscovery()`

UnsetSkipConnectionDiscovery ensures that no value is present for SkipConnectionDiscovery, not even an explicit nil
### GetServers

`func (o *MssqlConnectionResponseParamsAllOf) GetServers() []SQLServer`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *MssqlConnectionResponseParamsAllOf) GetServersOk() (*[]SQLServer, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *MssqlConnectionResponseParamsAllOf) SetServers(v []SQLServer)`

SetServers sets Servers field to given value.

### HasServers

`func (o *MssqlConnectionResponseParamsAllOf) HasServers() bool`

HasServers returns a boolean if a field has been set.

### SetServersNil

`func (o *MssqlConnectionResponseParamsAllOf) SetServersNil(b bool)`

 SetServersNil sets the value for Servers to be an explicit nil

### UnsetServers
`func (o *MssqlConnectionResponseParamsAllOf) UnsetServers()`

UnsetServers ensures that no value is present for Servers, not even an explicit nil
### GetFciClusters

`func (o *MssqlConnectionResponseParamsAllOf) GetFciClusters() []FCICluster`

GetFciClusters returns the FciClusters field if non-nil, zero value otherwise.

### GetFciClustersOk

`func (o *MssqlConnectionResponseParamsAllOf) GetFciClustersOk() (*[]FCICluster, bool)`

GetFciClustersOk returns a tuple with the FciClusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFciClusters

`func (o *MssqlConnectionResponseParamsAllOf) SetFciClusters(v []FCICluster)`

SetFciClusters sets FciClusters field to given value.

### HasFciClusters

`func (o *MssqlConnectionResponseParamsAllOf) HasFciClusters() bool`

HasFciClusters returns a boolean if a field has been set.

### SetFciClustersNil

`func (o *MssqlConnectionResponseParamsAllOf) SetFciClustersNil(b bool)`

 SetFciClustersNil sets the value for FciClusters to be an explicit nil

### UnsetFciClusters
`func (o *MssqlConnectionResponseParamsAllOf) UnsetFciClusters()`

UnsetFciClusters ensures that no value is present for FciClusters, not even an explicit nil
### GetAagGroups

`func (o *MssqlConnectionResponseParamsAllOf) GetAagGroups() []AAGGroup`

GetAagGroups returns the AagGroups field if non-nil, zero value otherwise.

### GetAagGroupsOk

`func (o *MssqlConnectionResponseParamsAllOf) GetAagGroupsOk() (*[]AAGGroup, bool)`

GetAagGroupsOk returns a tuple with the AagGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAagGroups

`func (o *MssqlConnectionResponseParamsAllOf) SetAagGroups(v []AAGGroup)`

SetAagGroups sets AagGroups field to given value.

### HasAagGroups

`func (o *MssqlConnectionResponseParamsAllOf) HasAagGroups() bool`

HasAagGroups returns a boolean if a field has been set.

### SetAagGroupsNil

`func (o *MssqlConnectionResponseParamsAllOf) SetAagGroupsNil(b bool)`

 SetAagGroupsNil sets the value for AagGroups to be an explicit nil

### UnsetAagGroups
`func (o *MssqlConnectionResponseParamsAllOf) UnsetAagGroups()`

UnsetAagGroups ensures that no value is present for AagGroups, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


