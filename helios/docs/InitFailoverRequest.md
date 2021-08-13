# InitFailoverRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceCluster** | Pointer to [**FailoverSourceCluster**](FailoverSourceCluster.md) |  | [optional] 
**ReplicationCluster** | Pointer to [**FailoverReplicaCluster**](FailoverReplicaCluster.md) |  | [optional] 

## Methods

### NewInitFailoverRequest

`func NewInitFailoverRequest() *InitFailoverRequest`

NewInitFailoverRequest instantiates a new InitFailoverRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitFailoverRequestWithDefaults

`func NewInitFailoverRequestWithDefaults() *InitFailoverRequest`

NewInitFailoverRequestWithDefaults instantiates a new InitFailoverRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceCluster

`func (o *InitFailoverRequest) GetSourceCluster() FailoverSourceCluster`

GetSourceCluster returns the SourceCluster field if non-nil, zero value otherwise.

### GetSourceClusterOk

`func (o *InitFailoverRequest) GetSourceClusterOk() (*FailoverSourceCluster, bool)`

GetSourceClusterOk returns a tuple with the SourceCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCluster

`func (o *InitFailoverRequest) SetSourceCluster(v FailoverSourceCluster)`

SetSourceCluster sets SourceCluster field to given value.

### HasSourceCluster

`func (o *InitFailoverRequest) HasSourceCluster() bool`

HasSourceCluster returns a boolean if a field has been set.

### GetReplicationCluster

`func (o *InitFailoverRequest) GetReplicationCluster() FailoverReplicaCluster`

GetReplicationCluster returns the ReplicationCluster field if non-nil, zero value otherwise.

### GetReplicationClusterOk

`func (o *InitFailoverRequest) GetReplicationClusterOk() (*FailoverReplicaCluster, bool)`

GetReplicationClusterOk returns a tuple with the ReplicationCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationCluster

`func (o *InitFailoverRequest) SetReplicationCluster(v FailoverReplicaCluster)`

SetReplicationCluster sets ReplicationCluster field to given value.

### HasReplicationCluster

`func (o *InitFailoverRequest) HasReplicationCluster() bool`

HasReplicationCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


