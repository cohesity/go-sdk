# InitFailoverRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProtectionGroupEnvironment** | Pointer to **NullableString** | If this field is specified then protection groups will be looked up only for this specific environment | [optional] 
**ReplicationCluster** | Pointer to [**FailoverReplicaCluster**](FailoverReplicaCluster.md) |  | [optional] 
**SourceCluster** | Pointer to [**FailoverSourceCluster**](FailoverSourceCluster.md) |  | [optional] 

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

### GetProtectionGroupEnvironment

`func (o *InitFailoverRequest) GetProtectionGroupEnvironment() string`

GetProtectionGroupEnvironment returns the ProtectionGroupEnvironment field if non-nil, zero value otherwise.

### GetProtectionGroupEnvironmentOk

`func (o *InitFailoverRequest) GetProtectionGroupEnvironmentOk() (*string, bool)`

GetProtectionGroupEnvironmentOk returns a tuple with the ProtectionGroupEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupEnvironment

`func (o *InitFailoverRequest) SetProtectionGroupEnvironment(v string)`

SetProtectionGroupEnvironment sets ProtectionGroupEnvironment field to given value.

### HasProtectionGroupEnvironment

`func (o *InitFailoverRequest) HasProtectionGroupEnvironment() bool`

HasProtectionGroupEnvironment returns a boolean if a field has been set.

### SetProtectionGroupEnvironmentNil

`func (o *InitFailoverRequest) SetProtectionGroupEnvironmentNil(b bool)`

 SetProtectionGroupEnvironmentNil sets the value for ProtectionGroupEnvironment to be an explicit nil

### UnsetProtectionGroupEnvironment
`func (o *InitFailoverRequest) UnsetProtectionGroupEnvironment()`

UnsetProtectionGroupEnvironment ensures that no value is present for ProtectionGroupEnvironment, not even an explicit nil
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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


