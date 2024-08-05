# ReplicationParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllEndpointsReachable** | Pointer to **NullableBool** | Specifies if all endpoints on Remote Cluster are reachable. | [optional] [default to false]
**BandwidthLimit** | Pointer to [**ReplicationParamsBandwidthLimit**](ReplicationParamsBandwidthLimit.md) |  | [optional] 
**CompressionEnabled** | Pointer to **NullableBool** | Specifies whether to compress the outbound data when transferring the replication data over the network to the Remote Cluster. | [optional] [default to true]
**EncryptionKey** | Pointer to **NullableString** | Specifies the encryption key used for encrypting the replication data from a local Cluster to a Remote Cluster. If a key is not specified, replication traffic encryption is disabled. When Snapshots are replicated from a local Cluster to a Remote Cluster, the encryption key specified on the local Cluster must be the same as the key specified on the Remote Cluster. | [optional] 
**StorageDomainPairs** | Pointer to [**[]StorageDomainPair**](StorageDomainPair.md) | Specifies a list of Storage Domain pairs. | [optional] 

## Methods

### NewReplicationParams

`func NewReplicationParams() *ReplicationParams`

NewReplicationParams instantiates a new ReplicationParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicationParamsWithDefaults

`func NewReplicationParamsWithDefaults() *ReplicationParams`

NewReplicationParamsWithDefaults instantiates a new ReplicationParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllEndpointsReachable

`func (o *ReplicationParams) GetAllEndpointsReachable() bool`

GetAllEndpointsReachable returns the AllEndpointsReachable field if non-nil, zero value otherwise.

### GetAllEndpointsReachableOk

`func (o *ReplicationParams) GetAllEndpointsReachableOk() (*bool, bool)`

GetAllEndpointsReachableOk returns a tuple with the AllEndpointsReachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllEndpointsReachable

`func (o *ReplicationParams) SetAllEndpointsReachable(v bool)`

SetAllEndpointsReachable sets AllEndpointsReachable field to given value.

### HasAllEndpointsReachable

`func (o *ReplicationParams) HasAllEndpointsReachable() bool`

HasAllEndpointsReachable returns a boolean if a field has been set.

### SetAllEndpointsReachableNil

`func (o *ReplicationParams) SetAllEndpointsReachableNil(b bool)`

 SetAllEndpointsReachableNil sets the value for AllEndpointsReachable to be an explicit nil

### UnsetAllEndpointsReachable
`func (o *ReplicationParams) UnsetAllEndpointsReachable()`

UnsetAllEndpointsReachable ensures that no value is present for AllEndpointsReachable, not even an explicit nil
### GetBandwidthLimit

`func (o *ReplicationParams) GetBandwidthLimit() ReplicationParamsBandwidthLimit`

GetBandwidthLimit returns the BandwidthLimit field if non-nil, zero value otherwise.

### GetBandwidthLimitOk

`func (o *ReplicationParams) GetBandwidthLimitOk() (*ReplicationParamsBandwidthLimit, bool)`

GetBandwidthLimitOk returns a tuple with the BandwidthLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthLimit

`func (o *ReplicationParams) SetBandwidthLimit(v ReplicationParamsBandwidthLimit)`

SetBandwidthLimit sets BandwidthLimit field to given value.

### HasBandwidthLimit

`func (o *ReplicationParams) HasBandwidthLimit() bool`

HasBandwidthLimit returns a boolean if a field has been set.

### GetCompressionEnabled

`func (o *ReplicationParams) GetCompressionEnabled() bool`

GetCompressionEnabled returns the CompressionEnabled field if non-nil, zero value otherwise.

### GetCompressionEnabledOk

`func (o *ReplicationParams) GetCompressionEnabledOk() (*bool, bool)`

GetCompressionEnabledOk returns a tuple with the CompressionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionEnabled

`func (o *ReplicationParams) SetCompressionEnabled(v bool)`

SetCompressionEnabled sets CompressionEnabled field to given value.

### HasCompressionEnabled

`func (o *ReplicationParams) HasCompressionEnabled() bool`

HasCompressionEnabled returns a boolean if a field has been set.

### SetCompressionEnabledNil

`func (o *ReplicationParams) SetCompressionEnabledNil(b bool)`

 SetCompressionEnabledNil sets the value for CompressionEnabled to be an explicit nil

### UnsetCompressionEnabled
`func (o *ReplicationParams) UnsetCompressionEnabled()`

UnsetCompressionEnabled ensures that no value is present for CompressionEnabled, not even an explicit nil
### GetEncryptionKey

`func (o *ReplicationParams) GetEncryptionKey() string`

GetEncryptionKey returns the EncryptionKey field if non-nil, zero value otherwise.

### GetEncryptionKeyOk

`func (o *ReplicationParams) GetEncryptionKeyOk() (*string, bool)`

GetEncryptionKeyOk returns a tuple with the EncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKey

`func (o *ReplicationParams) SetEncryptionKey(v string)`

SetEncryptionKey sets EncryptionKey field to given value.

### HasEncryptionKey

`func (o *ReplicationParams) HasEncryptionKey() bool`

HasEncryptionKey returns a boolean if a field has been set.

### SetEncryptionKeyNil

`func (o *ReplicationParams) SetEncryptionKeyNil(b bool)`

 SetEncryptionKeyNil sets the value for EncryptionKey to be an explicit nil

### UnsetEncryptionKey
`func (o *ReplicationParams) UnsetEncryptionKey()`

UnsetEncryptionKey ensures that no value is present for EncryptionKey, not even an explicit nil
### GetStorageDomainPairs

`func (o *ReplicationParams) GetStorageDomainPairs() []StorageDomainPair`

GetStorageDomainPairs returns the StorageDomainPairs field if non-nil, zero value otherwise.

### GetStorageDomainPairsOk

`func (o *ReplicationParams) GetStorageDomainPairsOk() (*[]StorageDomainPair, bool)`

GetStorageDomainPairsOk returns a tuple with the StorageDomainPairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainPairs

`func (o *ReplicationParams) SetStorageDomainPairs(v []StorageDomainPair)`

SetStorageDomainPairs sets StorageDomainPairs field to given value.

### HasStorageDomainPairs

`func (o *ReplicationParams) HasStorageDomainPairs() bool`

HasStorageDomainPairs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


