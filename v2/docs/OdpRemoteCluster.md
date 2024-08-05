# OdpRemoteCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **NullableInt64** | Specifies the ODP Remote Cluster id. | 
**ClusterIncarnationId** | **NullableInt64** | Specifies the ODP Remote Cluster incarnation id. | 
**AllEndpointsReachable** | Pointer to **NullableBool** | Specifies if all endpoints on ODP Remote Cluster are reachable. | [optional] 
**ClusterIdStale** | Pointer to **NullableBool** | Specifies if the cluster id is stale and needs to be refreshed. | [optional] 
**ClusterName** | **NullableString** | Specifies the ODP Remote Cluster name. | 
**CompressionEnabled** | Pointer to **NullableBool** | Specifies whether to compress the data transferred to ODP Remote Cluster. | [optional] 
**InterfaceGroupName** | Pointer to **NullableString** | Specifies the interface group name of the ODP Remote Cluster. | [optional] 
**KeyEncryptionKey** | Pointer to **NullableString** | Specifies the key used for encrypting the data transferred to ODP Remote Cluster. | [optional] 
**RemoteTenantId** | Pointer to **NullableString** | Specifies the tenant id for ODP Remote Cluster. | [optional] 
**StorageDomainPairs** | Pointer to [**[]StorageDomainPair**](StorageDomainPair.md) | Specifies a list of Storage Domain pairs. | [optional] 
**TenantId** | Pointer to **NullableString** | Specifies the tenant id. | [optional] 
**UseBifrostBrokerChannel** | Pointer to **NullableBool** | Specifies whether to use Bifrost Broker channel for remote connection. | [optional] 
**UsedForReplication** | Pointer to **NullableBool** | Specifies if the ODP Remote Cluster is used for replication. | [optional] 

## Methods

### NewOdpRemoteCluster

`func NewOdpRemoteCluster(clusterId NullableInt64, clusterIncarnationId NullableInt64, clusterName NullableString, ) *OdpRemoteCluster`

NewOdpRemoteCluster instantiates a new OdpRemoteCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOdpRemoteClusterWithDefaults

`func NewOdpRemoteClusterWithDefaults() *OdpRemoteCluster`

NewOdpRemoteClusterWithDefaults instantiates a new OdpRemoteCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *OdpRemoteCluster) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *OdpRemoteCluster) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *OdpRemoteCluster) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.


### SetClusterIdNil

`func (o *OdpRemoteCluster) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *OdpRemoteCluster) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetClusterIncarnationId

`func (o *OdpRemoteCluster) GetClusterIncarnationId() int64`

GetClusterIncarnationId returns the ClusterIncarnationId field if non-nil, zero value otherwise.

### GetClusterIncarnationIdOk

`func (o *OdpRemoteCluster) GetClusterIncarnationIdOk() (*int64, bool)`

GetClusterIncarnationIdOk returns a tuple with the ClusterIncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIncarnationId

`func (o *OdpRemoteCluster) SetClusterIncarnationId(v int64)`

SetClusterIncarnationId sets ClusterIncarnationId field to given value.


### SetClusterIncarnationIdNil

`func (o *OdpRemoteCluster) SetClusterIncarnationIdNil(b bool)`

 SetClusterIncarnationIdNil sets the value for ClusterIncarnationId to be an explicit nil

### UnsetClusterIncarnationId
`func (o *OdpRemoteCluster) UnsetClusterIncarnationId()`

UnsetClusterIncarnationId ensures that no value is present for ClusterIncarnationId, not even an explicit nil
### GetAllEndpointsReachable

`func (o *OdpRemoteCluster) GetAllEndpointsReachable() bool`

GetAllEndpointsReachable returns the AllEndpointsReachable field if non-nil, zero value otherwise.

### GetAllEndpointsReachableOk

`func (o *OdpRemoteCluster) GetAllEndpointsReachableOk() (*bool, bool)`

GetAllEndpointsReachableOk returns a tuple with the AllEndpointsReachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllEndpointsReachable

`func (o *OdpRemoteCluster) SetAllEndpointsReachable(v bool)`

SetAllEndpointsReachable sets AllEndpointsReachable field to given value.

### HasAllEndpointsReachable

`func (o *OdpRemoteCluster) HasAllEndpointsReachable() bool`

HasAllEndpointsReachable returns a boolean if a field has been set.

### SetAllEndpointsReachableNil

`func (o *OdpRemoteCluster) SetAllEndpointsReachableNil(b bool)`

 SetAllEndpointsReachableNil sets the value for AllEndpointsReachable to be an explicit nil

### UnsetAllEndpointsReachable
`func (o *OdpRemoteCluster) UnsetAllEndpointsReachable()`

UnsetAllEndpointsReachable ensures that no value is present for AllEndpointsReachable, not even an explicit nil
### GetClusterIdStale

`func (o *OdpRemoteCluster) GetClusterIdStale() bool`

GetClusterIdStale returns the ClusterIdStale field if non-nil, zero value otherwise.

### GetClusterIdStaleOk

`func (o *OdpRemoteCluster) GetClusterIdStaleOk() (*bool, bool)`

GetClusterIdStaleOk returns a tuple with the ClusterIdStale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIdStale

`func (o *OdpRemoteCluster) SetClusterIdStale(v bool)`

SetClusterIdStale sets ClusterIdStale field to given value.

### HasClusterIdStale

`func (o *OdpRemoteCluster) HasClusterIdStale() bool`

HasClusterIdStale returns a boolean if a field has been set.

### SetClusterIdStaleNil

`func (o *OdpRemoteCluster) SetClusterIdStaleNil(b bool)`

 SetClusterIdStaleNil sets the value for ClusterIdStale to be an explicit nil

### UnsetClusterIdStale
`func (o *OdpRemoteCluster) UnsetClusterIdStale()`

UnsetClusterIdStale ensures that no value is present for ClusterIdStale, not even an explicit nil
### GetClusterName

`func (o *OdpRemoteCluster) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *OdpRemoteCluster) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *OdpRemoteCluster) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### SetClusterNameNil

`func (o *OdpRemoteCluster) SetClusterNameNil(b bool)`

 SetClusterNameNil sets the value for ClusterName to be an explicit nil

### UnsetClusterName
`func (o *OdpRemoteCluster) UnsetClusterName()`

UnsetClusterName ensures that no value is present for ClusterName, not even an explicit nil
### GetCompressionEnabled

`func (o *OdpRemoteCluster) GetCompressionEnabled() bool`

GetCompressionEnabled returns the CompressionEnabled field if non-nil, zero value otherwise.

### GetCompressionEnabledOk

`func (o *OdpRemoteCluster) GetCompressionEnabledOk() (*bool, bool)`

GetCompressionEnabledOk returns a tuple with the CompressionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionEnabled

`func (o *OdpRemoteCluster) SetCompressionEnabled(v bool)`

SetCompressionEnabled sets CompressionEnabled field to given value.

### HasCompressionEnabled

`func (o *OdpRemoteCluster) HasCompressionEnabled() bool`

HasCompressionEnabled returns a boolean if a field has been set.

### SetCompressionEnabledNil

`func (o *OdpRemoteCluster) SetCompressionEnabledNil(b bool)`

 SetCompressionEnabledNil sets the value for CompressionEnabled to be an explicit nil

### UnsetCompressionEnabled
`func (o *OdpRemoteCluster) UnsetCompressionEnabled()`

UnsetCompressionEnabled ensures that no value is present for CompressionEnabled, not even an explicit nil
### GetInterfaceGroupName

`func (o *OdpRemoteCluster) GetInterfaceGroupName() string`

GetInterfaceGroupName returns the InterfaceGroupName field if non-nil, zero value otherwise.

### GetInterfaceGroupNameOk

`func (o *OdpRemoteCluster) GetInterfaceGroupNameOk() (*string, bool)`

GetInterfaceGroupNameOk returns a tuple with the InterfaceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceGroupName

`func (o *OdpRemoteCluster) SetInterfaceGroupName(v string)`

SetInterfaceGroupName sets InterfaceGroupName field to given value.

### HasInterfaceGroupName

`func (o *OdpRemoteCluster) HasInterfaceGroupName() bool`

HasInterfaceGroupName returns a boolean if a field has been set.

### SetInterfaceGroupNameNil

`func (o *OdpRemoteCluster) SetInterfaceGroupNameNil(b bool)`

 SetInterfaceGroupNameNil sets the value for InterfaceGroupName to be an explicit nil

### UnsetInterfaceGroupName
`func (o *OdpRemoteCluster) UnsetInterfaceGroupName()`

UnsetInterfaceGroupName ensures that no value is present for InterfaceGroupName, not even an explicit nil
### GetKeyEncryptionKey

`func (o *OdpRemoteCluster) GetKeyEncryptionKey() string`

GetKeyEncryptionKey returns the KeyEncryptionKey field if non-nil, zero value otherwise.

### GetKeyEncryptionKeyOk

`func (o *OdpRemoteCluster) GetKeyEncryptionKeyOk() (*string, bool)`

GetKeyEncryptionKeyOk returns a tuple with the KeyEncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyEncryptionKey

`func (o *OdpRemoteCluster) SetKeyEncryptionKey(v string)`

SetKeyEncryptionKey sets KeyEncryptionKey field to given value.

### HasKeyEncryptionKey

`func (o *OdpRemoteCluster) HasKeyEncryptionKey() bool`

HasKeyEncryptionKey returns a boolean if a field has been set.

### SetKeyEncryptionKeyNil

`func (o *OdpRemoteCluster) SetKeyEncryptionKeyNil(b bool)`

 SetKeyEncryptionKeyNil sets the value for KeyEncryptionKey to be an explicit nil

### UnsetKeyEncryptionKey
`func (o *OdpRemoteCluster) UnsetKeyEncryptionKey()`

UnsetKeyEncryptionKey ensures that no value is present for KeyEncryptionKey, not even an explicit nil
### GetRemoteTenantId

`func (o *OdpRemoteCluster) GetRemoteTenantId() string`

GetRemoteTenantId returns the RemoteTenantId field if non-nil, zero value otherwise.

### GetRemoteTenantIdOk

`func (o *OdpRemoteCluster) GetRemoteTenantIdOk() (*string, bool)`

GetRemoteTenantIdOk returns a tuple with the RemoteTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteTenantId

`func (o *OdpRemoteCluster) SetRemoteTenantId(v string)`

SetRemoteTenantId sets RemoteTenantId field to given value.

### HasRemoteTenantId

`func (o *OdpRemoteCluster) HasRemoteTenantId() bool`

HasRemoteTenantId returns a boolean if a field has been set.

### SetRemoteTenantIdNil

`func (o *OdpRemoteCluster) SetRemoteTenantIdNil(b bool)`

 SetRemoteTenantIdNil sets the value for RemoteTenantId to be an explicit nil

### UnsetRemoteTenantId
`func (o *OdpRemoteCluster) UnsetRemoteTenantId()`

UnsetRemoteTenantId ensures that no value is present for RemoteTenantId, not even an explicit nil
### GetStorageDomainPairs

`func (o *OdpRemoteCluster) GetStorageDomainPairs() []StorageDomainPair`

GetStorageDomainPairs returns the StorageDomainPairs field if non-nil, zero value otherwise.

### GetStorageDomainPairsOk

`func (o *OdpRemoteCluster) GetStorageDomainPairsOk() (*[]StorageDomainPair, bool)`

GetStorageDomainPairsOk returns a tuple with the StorageDomainPairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainPairs

`func (o *OdpRemoteCluster) SetStorageDomainPairs(v []StorageDomainPair)`

SetStorageDomainPairs sets StorageDomainPairs field to given value.

### HasStorageDomainPairs

`func (o *OdpRemoteCluster) HasStorageDomainPairs() bool`

HasStorageDomainPairs returns a boolean if a field has been set.

### SetStorageDomainPairsNil

`func (o *OdpRemoteCluster) SetStorageDomainPairsNil(b bool)`

 SetStorageDomainPairsNil sets the value for StorageDomainPairs to be an explicit nil

### UnsetStorageDomainPairs
`func (o *OdpRemoteCluster) UnsetStorageDomainPairs()`

UnsetStorageDomainPairs ensures that no value is present for StorageDomainPairs, not even an explicit nil
### GetTenantId

`func (o *OdpRemoteCluster) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *OdpRemoteCluster) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *OdpRemoteCluster) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *OdpRemoteCluster) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *OdpRemoteCluster) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *OdpRemoteCluster) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetUseBifrostBrokerChannel

`func (o *OdpRemoteCluster) GetUseBifrostBrokerChannel() bool`

GetUseBifrostBrokerChannel returns the UseBifrostBrokerChannel field if non-nil, zero value otherwise.

### GetUseBifrostBrokerChannelOk

`func (o *OdpRemoteCluster) GetUseBifrostBrokerChannelOk() (*bool, bool)`

GetUseBifrostBrokerChannelOk returns a tuple with the UseBifrostBrokerChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBifrostBrokerChannel

`func (o *OdpRemoteCluster) SetUseBifrostBrokerChannel(v bool)`

SetUseBifrostBrokerChannel sets UseBifrostBrokerChannel field to given value.

### HasUseBifrostBrokerChannel

`func (o *OdpRemoteCluster) HasUseBifrostBrokerChannel() bool`

HasUseBifrostBrokerChannel returns a boolean if a field has been set.

### SetUseBifrostBrokerChannelNil

`func (o *OdpRemoteCluster) SetUseBifrostBrokerChannelNil(b bool)`

 SetUseBifrostBrokerChannelNil sets the value for UseBifrostBrokerChannel to be an explicit nil

### UnsetUseBifrostBrokerChannel
`func (o *OdpRemoteCluster) UnsetUseBifrostBrokerChannel()`

UnsetUseBifrostBrokerChannel ensures that no value is present for UseBifrostBrokerChannel, not even an explicit nil
### GetUsedForReplication

`func (o *OdpRemoteCluster) GetUsedForReplication() bool`

GetUsedForReplication returns the UsedForReplication field if non-nil, zero value otherwise.

### GetUsedForReplicationOk

`func (o *OdpRemoteCluster) GetUsedForReplicationOk() (*bool, bool)`

GetUsedForReplicationOk returns a tuple with the UsedForReplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedForReplication

`func (o *OdpRemoteCluster) SetUsedForReplication(v bool)`

SetUsedForReplication sets UsedForReplication field to given value.

### HasUsedForReplication

`func (o *OdpRemoteCluster) HasUsedForReplication() bool`

HasUsedForReplication returns a boolean if a field has been set.

### SetUsedForReplicationNil

`func (o *OdpRemoteCluster) SetUsedForReplicationNil(b bool)`

 SetUsedForReplicationNil sets the value for UsedForReplication to be an explicit nil

### UnsetUsedForReplication
`func (o *OdpRemoteCluster) UnsetUsedForReplication()`

UnsetUsedForReplication ensures that no value is present for UsedForReplication, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


