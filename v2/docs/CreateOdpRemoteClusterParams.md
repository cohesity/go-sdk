# CreateOdpRemoteClusterParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**ClusterId** | **NullableInt64** | Specifies the ODP Remote Cluster id. | 
**ClusterIncarnationId** | **NullableInt64** | Specifies the ODP Remote Cluster incarnation id. | 

## Methods

### NewCreateOdpRemoteClusterParams

`func NewCreateOdpRemoteClusterParams(clusterName NullableString, clusterId NullableInt64, clusterIncarnationId NullableInt64, ) *CreateOdpRemoteClusterParams`

NewCreateOdpRemoteClusterParams instantiates a new CreateOdpRemoteClusterParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOdpRemoteClusterParamsWithDefaults

`func NewCreateOdpRemoteClusterParamsWithDefaults() *CreateOdpRemoteClusterParams`

NewCreateOdpRemoteClusterParamsWithDefaults instantiates a new CreateOdpRemoteClusterParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllEndpointsReachable

`func (o *CreateOdpRemoteClusterParams) GetAllEndpointsReachable() bool`

GetAllEndpointsReachable returns the AllEndpointsReachable field if non-nil, zero value otherwise.

### GetAllEndpointsReachableOk

`func (o *CreateOdpRemoteClusterParams) GetAllEndpointsReachableOk() (*bool, bool)`

GetAllEndpointsReachableOk returns a tuple with the AllEndpointsReachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllEndpointsReachable

`func (o *CreateOdpRemoteClusterParams) SetAllEndpointsReachable(v bool)`

SetAllEndpointsReachable sets AllEndpointsReachable field to given value.

### HasAllEndpointsReachable

`func (o *CreateOdpRemoteClusterParams) HasAllEndpointsReachable() bool`

HasAllEndpointsReachable returns a boolean if a field has been set.

### SetAllEndpointsReachableNil

`func (o *CreateOdpRemoteClusterParams) SetAllEndpointsReachableNil(b bool)`

 SetAllEndpointsReachableNil sets the value for AllEndpointsReachable to be an explicit nil

### UnsetAllEndpointsReachable
`func (o *CreateOdpRemoteClusterParams) UnsetAllEndpointsReachable()`

UnsetAllEndpointsReachable ensures that no value is present for AllEndpointsReachable, not even an explicit nil
### GetClusterIdStale

`func (o *CreateOdpRemoteClusterParams) GetClusterIdStale() bool`

GetClusterIdStale returns the ClusterIdStale field if non-nil, zero value otherwise.

### GetClusterIdStaleOk

`func (o *CreateOdpRemoteClusterParams) GetClusterIdStaleOk() (*bool, bool)`

GetClusterIdStaleOk returns a tuple with the ClusterIdStale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIdStale

`func (o *CreateOdpRemoteClusterParams) SetClusterIdStale(v bool)`

SetClusterIdStale sets ClusterIdStale field to given value.

### HasClusterIdStale

`func (o *CreateOdpRemoteClusterParams) HasClusterIdStale() bool`

HasClusterIdStale returns a boolean if a field has been set.

### SetClusterIdStaleNil

`func (o *CreateOdpRemoteClusterParams) SetClusterIdStaleNil(b bool)`

 SetClusterIdStaleNil sets the value for ClusterIdStale to be an explicit nil

### UnsetClusterIdStale
`func (o *CreateOdpRemoteClusterParams) UnsetClusterIdStale()`

UnsetClusterIdStale ensures that no value is present for ClusterIdStale, not even an explicit nil
### GetClusterName

`func (o *CreateOdpRemoteClusterParams) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *CreateOdpRemoteClusterParams) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *CreateOdpRemoteClusterParams) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### SetClusterNameNil

`func (o *CreateOdpRemoteClusterParams) SetClusterNameNil(b bool)`

 SetClusterNameNil sets the value for ClusterName to be an explicit nil

### UnsetClusterName
`func (o *CreateOdpRemoteClusterParams) UnsetClusterName()`

UnsetClusterName ensures that no value is present for ClusterName, not even an explicit nil
### GetCompressionEnabled

`func (o *CreateOdpRemoteClusterParams) GetCompressionEnabled() bool`

GetCompressionEnabled returns the CompressionEnabled field if non-nil, zero value otherwise.

### GetCompressionEnabledOk

`func (o *CreateOdpRemoteClusterParams) GetCompressionEnabledOk() (*bool, bool)`

GetCompressionEnabledOk returns a tuple with the CompressionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionEnabled

`func (o *CreateOdpRemoteClusterParams) SetCompressionEnabled(v bool)`

SetCompressionEnabled sets CompressionEnabled field to given value.

### HasCompressionEnabled

`func (o *CreateOdpRemoteClusterParams) HasCompressionEnabled() bool`

HasCompressionEnabled returns a boolean if a field has been set.

### SetCompressionEnabledNil

`func (o *CreateOdpRemoteClusterParams) SetCompressionEnabledNil(b bool)`

 SetCompressionEnabledNil sets the value for CompressionEnabled to be an explicit nil

### UnsetCompressionEnabled
`func (o *CreateOdpRemoteClusterParams) UnsetCompressionEnabled()`

UnsetCompressionEnabled ensures that no value is present for CompressionEnabled, not even an explicit nil
### GetInterfaceGroupName

`func (o *CreateOdpRemoteClusterParams) GetInterfaceGroupName() string`

GetInterfaceGroupName returns the InterfaceGroupName field if non-nil, zero value otherwise.

### GetInterfaceGroupNameOk

`func (o *CreateOdpRemoteClusterParams) GetInterfaceGroupNameOk() (*string, bool)`

GetInterfaceGroupNameOk returns a tuple with the InterfaceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceGroupName

`func (o *CreateOdpRemoteClusterParams) SetInterfaceGroupName(v string)`

SetInterfaceGroupName sets InterfaceGroupName field to given value.

### HasInterfaceGroupName

`func (o *CreateOdpRemoteClusterParams) HasInterfaceGroupName() bool`

HasInterfaceGroupName returns a boolean if a field has been set.

### SetInterfaceGroupNameNil

`func (o *CreateOdpRemoteClusterParams) SetInterfaceGroupNameNil(b bool)`

 SetInterfaceGroupNameNil sets the value for InterfaceGroupName to be an explicit nil

### UnsetInterfaceGroupName
`func (o *CreateOdpRemoteClusterParams) UnsetInterfaceGroupName()`

UnsetInterfaceGroupName ensures that no value is present for InterfaceGroupName, not even an explicit nil
### GetKeyEncryptionKey

`func (o *CreateOdpRemoteClusterParams) GetKeyEncryptionKey() string`

GetKeyEncryptionKey returns the KeyEncryptionKey field if non-nil, zero value otherwise.

### GetKeyEncryptionKeyOk

`func (o *CreateOdpRemoteClusterParams) GetKeyEncryptionKeyOk() (*string, bool)`

GetKeyEncryptionKeyOk returns a tuple with the KeyEncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyEncryptionKey

`func (o *CreateOdpRemoteClusterParams) SetKeyEncryptionKey(v string)`

SetKeyEncryptionKey sets KeyEncryptionKey field to given value.

### HasKeyEncryptionKey

`func (o *CreateOdpRemoteClusterParams) HasKeyEncryptionKey() bool`

HasKeyEncryptionKey returns a boolean if a field has been set.

### SetKeyEncryptionKeyNil

`func (o *CreateOdpRemoteClusterParams) SetKeyEncryptionKeyNil(b bool)`

 SetKeyEncryptionKeyNil sets the value for KeyEncryptionKey to be an explicit nil

### UnsetKeyEncryptionKey
`func (o *CreateOdpRemoteClusterParams) UnsetKeyEncryptionKey()`

UnsetKeyEncryptionKey ensures that no value is present for KeyEncryptionKey, not even an explicit nil
### GetRemoteTenantId

`func (o *CreateOdpRemoteClusterParams) GetRemoteTenantId() string`

GetRemoteTenantId returns the RemoteTenantId field if non-nil, zero value otherwise.

### GetRemoteTenantIdOk

`func (o *CreateOdpRemoteClusterParams) GetRemoteTenantIdOk() (*string, bool)`

GetRemoteTenantIdOk returns a tuple with the RemoteTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteTenantId

`func (o *CreateOdpRemoteClusterParams) SetRemoteTenantId(v string)`

SetRemoteTenantId sets RemoteTenantId field to given value.

### HasRemoteTenantId

`func (o *CreateOdpRemoteClusterParams) HasRemoteTenantId() bool`

HasRemoteTenantId returns a boolean if a field has been set.

### SetRemoteTenantIdNil

`func (o *CreateOdpRemoteClusterParams) SetRemoteTenantIdNil(b bool)`

 SetRemoteTenantIdNil sets the value for RemoteTenantId to be an explicit nil

### UnsetRemoteTenantId
`func (o *CreateOdpRemoteClusterParams) UnsetRemoteTenantId()`

UnsetRemoteTenantId ensures that no value is present for RemoteTenantId, not even an explicit nil
### GetStorageDomainPairs

`func (o *CreateOdpRemoteClusterParams) GetStorageDomainPairs() []StorageDomainPair`

GetStorageDomainPairs returns the StorageDomainPairs field if non-nil, zero value otherwise.

### GetStorageDomainPairsOk

`func (o *CreateOdpRemoteClusterParams) GetStorageDomainPairsOk() (*[]StorageDomainPair, bool)`

GetStorageDomainPairsOk returns a tuple with the StorageDomainPairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainPairs

`func (o *CreateOdpRemoteClusterParams) SetStorageDomainPairs(v []StorageDomainPair)`

SetStorageDomainPairs sets StorageDomainPairs field to given value.

### HasStorageDomainPairs

`func (o *CreateOdpRemoteClusterParams) HasStorageDomainPairs() bool`

HasStorageDomainPairs returns a boolean if a field has been set.

### SetStorageDomainPairsNil

`func (o *CreateOdpRemoteClusterParams) SetStorageDomainPairsNil(b bool)`

 SetStorageDomainPairsNil sets the value for StorageDomainPairs to be an explicit nil

### UnsetStorageDomainPairs
`func (o *CreateOdpRemoteClusterParams) UnsetStorageDomainPairs()`

UnsetStorageDomainPairs ensures that no value is present for StorageDomainPairs, not even an explicit nil
### GetTenantId

`func (o *CreateOdpRemoteClusterParams) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CreateOdpRemoteClusterParams) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CreateOdpRemoteClusterParams) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CreateOdpRemoteClusterParams) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *CreateOdpRemoteClusterParams) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *CreateOdpRemoteClusterParams) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetUseBifrostBrokerChannel

`func (o *CreateOdpRemoteClusterParams) GetUseBifrostBrokerChannel() bool`

GetUseBifrostBrokerChannel returns the UseBifrostBrokerChannel field if non-nil, zero value otherwise.

### GetUseBifrostBrokerChannelOk

`func (o *CreateOdpRemoteClusterParams) GetUseBifrostBrokerChannelOk() (*bool, bool)`

GetUseBifrostBrokerChannelOk returns a tuple with the UseBifrostBrokerChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBifrostBrokerChannel

`func (o *CreateOdpRemoteClusterParams) SetUseBifrostBrokerChannel(v bool)`

SetUseBifrostBrokerChannel sets UseBifrostBrokerChannel field to given value.

### HasUseBifrostBrokerChannel

`func (o *CreateOdpRemoteClusterParams) HasUseBifrostBrokerChannel() bool`

HasUseBifrostBrokerChannel returns a boolean if a field has been set.

### SetUseBifrostBrokerChannelNil

`func (o *CreateOdpRemoteClusterParams) SetUseBifrostBrokerChannelNil(b bool)`

 SetUseBifrostBrokerChannelNil sets the value for UseBifrostBrokerChannel to be an explicit nil

### UnsetUseBifrostBrokerChannel
`func (o *CreateOdpRemoteClusterParams) UnsetUseBifrostBrokerChannel()`

UnsetUseBifrostBrokerChannel ensures that no value is present for UseBifrostBrokerChannel, not even an explicit nil
### GetUsedForReplication

`func (o *CreateOdpRemoteClusterParams) GetUsedForReplication() bool`

GetUsedForReplication returns the UsedForReplication field if non-nil, zero value otherwise.

### GetUsedForReplicationOk

`func (o *CreateOdpRemoteClusterParams) GetUsedForReplicationOk() (*bool, bool)`

GetUsedForReplicationOk returns a tuple with the UsedForReplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedForReplication

`func (o *CreateOdpRemoteClusterParams) SetUsedForReplication(v bool)`

SetUsedForReplication sets UsedForReplication field to given value.

### HasUsedForReplication

`func (o *CreateOdpRemoteClusterParams) HasUsedForReplication() bool`

HasUsedForReplication returns a boolean if a field has been set.

### SetUsedForReplicationNil

`func (o *CreateOdpRemoteClusterParams) SetUsedForReplicationNil(b bool)`

 SetUsedForReplicationNil sets the value for UsedForReplication to be an explicit nil

### UnsetUsedForReplication
`func (o *CreateOdpRemoteClusterParams) UnsetUsedForReplication()`

UnsetUsedForReplication ensures that no value is present for UsedForReplication, not even an explicit nil
### GetClusterId

`func (o *CreateOdpRemoteClusterParams) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *CreateOdpRemoteClusterParams) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *CreateOdpRemoteClusterParams) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.


### SetClusterIdNil

`func (o *CreateOdpRemoteClusterParams) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *CreateOdpRemoteClusterParams) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetClusterIncarnationId

`func (o *CreateOdpRemoteClusterParams) GetClusterIncarnationId() int64`

GetClusterIncarnationId returns the ClusterIncarnationId field if non-nil, zero value otherwise.

### GetClusterIncarnationIdOk

`func (o *CreateOdpRemoteClusterParams) GetClusterIncarnationIdOk() (*int64, bool)`

GetClusterIncarnationIdOk returns a tuple with the ClusterIncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIncarnationId

`func (o *CreateOdpRemoteClusterParams) SetClusterIncarnationId(v int64)`

SetClusterIncarnationId sets ClusterIncarnationId field to given value.


### SetClusterIncarnationIdNil

`func (o *CreateOdpRemoteClusterParams) SetClusterIncarnationIdNil(b bool)`

 SetClusterIncarnationIdNil sets the value for ClusterIncarnationId to be an explicit nil

### UnsetClusterIncarnationId
`func (o *CreateOdpRemoteClusterParams) UnsetClusterIncarnationId()`

UnsetClusterIncarnationId ensures that no value is present for ClusterIncarnationId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


