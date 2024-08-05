# UpdateOdpRemoteClusterParams

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

## Methods

### NewUpdateOdpRemoteClusterParams

`func NewUpdateOdpRemoteClusterParams(clusterName NullableString, ) *UpdateOdpRemoteClusterParams`

NewUpdateOdpRemoteClusterParams instantiates a new UpdateOdpRemoteClusterParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOdpRemoteClusterParamsWithDefaults

`func NewUpdateOdpRemoteClusterParamsWithDefaults() *UpdateOdpRemoteClusterParams`

NewUpdateOdpRemoteClusterParamsWithDefaults instantiates a new UpdateOdpRemoteClusterParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllEndpointsReachable

`func (o *UpdateOdpRemoteClusterParams) GetAllEndpointsReachable() bool`

GetAllEndpointsReachable returns the AllEndpointsReachable field if non-nil, zero value otherwise.

### GetAllEndpointsReachableOk

`func (o *UpdateOdpRemoteClusterParams) GetAllEndpointsReachableOk() (*bool, bool)`

GetAllEndpointsReachableOk returns a tuple with the AllEndpointsReachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllEndpointsReachable

`func (o *UpdateOdpRemoteClusterParams) SetAllEndpointsReachable(v bool)`

SetAllEndpointsReachable sets AllEndpointsReachable field to given value.

### HasAllEndpointsReachable

`func (o *UpdateOdpRemoteClusterParams) HasAllEndpointsReachable() bool`

HasAllEndpointsReachable returns a boolean if a field has been set.

### SetAllEndpointsReachableNil

`func (o *UpdateOdpRemoteClusterParams) SetAllEndpointsReachableNil(b bool)`

 SetAllEndpointsReachableNil sets the value for AllEndpointsReachable to be an explicit nil

### UnsetAllEndpointsReachable
`func (o *UpdateOdpRemoteClusterParams) UnsetAllEndpointsReachable()`

UnsetAllEndpointsReachable ensures that no value is present for AllEndpointsReachable, not even an explicit nil
### GetClusterIdStale

`func (o *UpdateOdpRemoteClusterParams) GetClusterIdStale() bool`

GetClusterIdStale returns the ClusterIdStale field if non-nil, zero value otherwise.

### GetClusterIdStaleOk

`func (o *UpdateOdpRemoteClusterParams) GetClusterIdStaleOk() (*bool, bool)`

GetClusterIdStaleOk returns a tuple with the ClusterIdStale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIdStale

`func (o *UpdateOdpRemoteClusterParams) SetClusterIdStale(v bool)`

SetClusterIdStale sets ClusterIdStale field to given value.

### HasClusterIdStale

`func (o *UpdateOdpRemoteClusterParams) HasClusterIdStale() bool`

HasClusterIdStale returns a boolean if a field has been set.

### SetClusterIdStaleNil

`func (o *UpdateOdpRemoteClusterParams) SetClusterIdStaleNil(b bool)`

 SetClusterIdStaleNil sets the value for ClusterIdStale to be an explicit nil

### UnsetClusterIdStale
`func (o *UpdateOdpRemoteClusterParams) UnsetClusterIdStale()`

UnsetClusterIdStale ensures that no value is present for ClusterIdStale, not even an explicit nil
### GetClusterName

`func (o *UpdateOdpRemoteClusterParams) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *UpdateOdpRemoteClusterParams) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *UpdateOdpRemoteClusterParams) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### SetClusterNameNil

`func (o *UpdateOdpRemoteClusterParams) SetClusterNameNil(b bool)`

 SetClusterNameNil sets the value for ClusterName to be an explicit nil

### UnsetClusterName
`func (o *UpdateOdpRemoteClusterParams) UnsetClusterName()`

UnsetClusterName ensures that no value is present for ClusterName, not even an explicit nil
### GetCompressionEnabled

`func (o *UpdateOdpRemoteClusterParams) GetCompressionEnabled() bool`

GetCompressionEnabled returns the CompressionEnabled field if non-nil, zero value otherwise.

### GetCompressionEnabledOk

`func (o *UpdateOdpRemoteClusterParams) GetCompressionEnabledOk() (*bool, bool)`

GetCompressionEnabledOk returns a tuple with the CompressionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionEnabled

`func (o *UpdateOdpRemoteClusterParams) SetCompressionEnabled(v bool)`

SetCompressionEnabled sets CompressionEnabled field to given value.

### HasCompressionEnabled

`func (o *UpdateOdpRemoteClusterParams) HasCompressionEnabled() bool`

HasCompressionEnabled returns a boolean if a field has been set.

### SetCompressionEnabledNil

`func (o *UpdateOdpRemoteClusterParams) SetCompressionEnabledNil(b bool)`

 SetCompressionEnabledNil sets the value for CompressionEnabled to be an explicit nil

### UnsetCompressionEnabled
`func (o *UpdateOdpRemoteClusterParams) UnsetCompressionEnabled()`

UnsetCompressionEnabled ensures that no value is present for CompressionEnabled, not even an explicit nil
### GetInterfaceGroupName

`func (o *UpdateOdpRemoteClusterParams) GetInterfaceGroupName() string`

GetInterfaceGroupName returns the InterfaceGroupName field if non-nil, zero value otherwise.

### GetInterfaceGroupNameOk

`func (o *UpdateOdpRemoteClusterParams) GetInterfaceGroupNameOk() (*string, bool)`

GetInterfaceGroupNameOk returns a tuple with the InterfaceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceGroupName

`func (o *UpdateOdpRemoteClusterParams) SetInterfaceGroupName(v string)`

SetInterfaceGroupName sets InterfaceGroupName field to given value.

### HasInterfaceGroupName

`func (o *UpdateOdpRemoteClusterParams) HasInterfaceGroupName() bool`

HasInterfaceGroupName returns a boolean if a field has been set.

### SetInterfaceGroupNameNil

`func (o *UpdateOdpRemoteClusterParams) SetInterfaceGroupNameNil(b bool)`

 SetInterfaceGroupNameNil sets the value for InterfaceGroupName to be an explicit nil

### UnsetInterfaceGroupName
`func (o *UpdateOdpRemoteClusterParams) UnsetInterfaceGroupName()`

UnsetInterfaceGroupName ensures that no value is present for InterfaceGroupName, not even an explicit nil
### GetKeyEncryptionKey

`func (o *UpdateOdpRemoteClusterParams) GetKeyEncryptionKey() string`

GetKeyEncryptionKey returns the KeyEncryptionKey field if non-nil, zero value otherwise.

### GetKeyEncryptionKeyOk

`func (o *UpdateOdpRemoteClusterParams) GetKeyEncryptionKeyOk() (*string, bool)`

GetKeyEncryptionKeyOk returns a tuple with the KeyEncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyEncryptionKey

`func (o *UpdateOdpRemoteClusterParams) SetKeyEncryptionKey(v string)`

SetKeyEncryptionKey sets KeyEncryptionKey field to given value.

### HasKeyEncryptionKey

`func (o *UpdateOdpRemoteClusterParams) HasKeyEncryptionKey() bool`

HasKeyEncryptionKey returns a boolean if a field has been set.

### SetKeyEncryptionKeyNil

`func (o *UpdateOdpRemoteClusterParams) SetKeyEncryptionKeyNil(b bool)`

 SetKeyEncryptionKeyNil sets the value for KeyEncryptionKey to be an explicit nil

### UnsetKeyEncryptionKey
`func (o *UpdateOdpRemoteClusterParams) UnsetKeyEncryptionKey()`

UnsetKeyEncryptionKey ensures that no value is present for KeyEncryptionKey, not even an explicit nil
### GetRemoteTenantId

`func (o *UpdateOdpRemoteClusterParams) GetRemoteTenantId() string`

GetRemoteTenantId returns the RemoteTenantId field if non-nil, zero value otherwise.

### GetRemoteTenantIdOk

`func (o *UpdateOdpRemoteClusterParams) GetRemoteTenantIdOk() (*string, bool)`

GetRemoteTenantIdOk returns a tuple with the RemoteTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteTenantId

`func (o *UpdateOdpRemoteClusterParams) SetRemoteTenantId(v string)`

SetRemoteTenantId sets RemoteTenantId field to given value.

### HasRemoteTenantId

`func (o *UpdateOdpRemoteClusterParams) HasRemoteTenantId() bool`

HasRemoteTenantId returns a boolean if a field has been set.

### SetRemoteTenantIdNil

`func (o *UpdateOdpRemoteClusterParams) SetRemoteTenantIdNil(b bool)`

 SetRemoteTenantIdNil sets the value for RemoteTenantId to be an explicit nil

### UnsetRemoteTenantId
`func (o *UpdateOdpRemoteClusterParams) UnsetRemoteTenantId()`

UnsetRemoteTenantId ensures that no value is present for RemoteTenantId, not even an explicit nil
### GetStorageDomainPairs

`func (o *UpdateOdpRemoteClusterParams) GetStorageDomainPairs() []StorageDomainPair`

GetStorageDomainPairs returns the StorageDomainPairs field if non-nil, zero value otherwise.

### GetStorageDomainPairsOk

`func (o *UpdateOdpRemoteClusterParams) GetStorageDomainPairsOk() (*[]StorageDomainPair, bool)`

GetStorageDomainPairsOk returns a tuple with the StorageDomainPairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainPairs

`func (o *UpdateOdpRemoteClusterParams) SetStorageDomainPairs(v []StorageDomainPair)`

SetStorageDomainPairs sets StorageDomainPairs field to given value.

### HasStorageDomainPairs

`func (o *UpdateOdpRemoteClusterParams) HasStorageDomainPairs() bool`

HasStorageDomainPairs returns a boolean if a field has been set.

### SetStorageDomainPairsNil

`func (o *UpdateOdpRemoteClusterParams) SetStorageDomainPairsNil(b bool)`

 SetStorageDomainPairsNil sets the value for StorageDomainPairs to be an explicit nil

### UnsetStorageDomainPairs
`func (o *UpdateOdpRemoteClusterParams) UnsetStorageDomainPairs()`

UnsetStorageDomainPairs ensures that no value is present for StorageDomainPairs, not even an explicit nil
### GetTenantId

`func (o *UpdateOdpRemoteClusterParams) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *UpdateOdpRemoteClusterParams) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *UpdateOdpRemoteClusterParams) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *UpdateOdpRemoteClusterParams) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *UpdateOdpRemoteClusterParams) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *UpdateOdpRemoteClusterParams) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetUseBifrostBrokerChannel

`func (o *UpdateOdpRemoteClusterParams) GetUseBifrostBrokerChannel() bool`

GetUseBifrostBrokerChannel returns the UseBifrostBrokerChannel field if non-nil, zero value otherwise.

### GetUseBifrostBrokerChannelOk

`func (o *UpdateOdpRemoteClusterParams) GetUseBifrostBrokerChannelOk() (*bool, bool)`

GetUseBifrostBrokerChannelOk returns a tuple with the UseBifrostBrokerChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBifrostBrokerChannel

`func (o *UpdateOdpRemoteClusterParams) SetUseBifrostBrokerChannel(v bool)`

SetUseBifrostBrokerChannel sets UseBifrostBrokerChannel field to given value.

### HasUseBifrostBrokerChannel

`func (o *UpdateOdpRemoteClusterParams) HasUseBifrostBrokerChannel() bool`

HasUseBifrostBrokerChannel returns a boolean if a field has been set.

### SetUseBifrostBrokerChannelNil

`func (o *UpdateOdpRemoteClusterParams) SetUseBifrostBrokerChannelNil(b bool)`

 SetUseBifrostBrokerChannelNil sets the value for UseBifrostBrokerChannel to be an explicit nil

### UnsetUseBifrostBrokerChannel
`func (o *UpdateOdpRemoteClusterParams) UnsetUseBifrostBrokerChannel()`

UnsetUseBifrostBrokerChannel ensures that no value is present for UseBifrostBrokerChannel, not even an explicit nil
### GetUsedForReplication

`func (o *UpdateOdpRemoteClusterParams) GetUsedForReplication() bool`

GetUsedForReplication returns the UsedForReplication field if non-nil, zero value otherwise.

### GetUsedForReplicationOk

`func (o *UpdateOdpRemoteClusterParams) GetUsedForReplicationOk() (*bool, bool)`

GetUsedForReplicationOk returns a tuple with the UsedForReplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedForReplication

`func (o *UpdateOdpRemoteClusterParams) SetUsedForReplication(v bool)`

SetUsedForReplication sets UsedForReplication field to given value.

### HasUsedForReplication

`func (o *UpdateOdpRemoteClusterParams) HasUsedForReplication() bool`

HasUsedForReplication returns a boolean if a field has been set.

### SetUsedForReplicationNil

`func (o *UpdateOdpRemoteClusterParams) SetUsedForReplicationNil(b bool)`

 SetUsedForReplicationNil sets the value for UsedForReplication to be an explicit nil

### UnsetUsedForReplication
`func (o *UpdateOdpRemoteClusterParams) UnsetUsedForReplication()`

UnsetUsedForReplication ensures that no value is present for UsedForReplication, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


