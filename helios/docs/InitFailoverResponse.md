# InitFailoverResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReplicaToSourceObjects** | Pointer to [**[]SourceReplicaObject**](SourceReplicaObject.md) | Specifies the list of corrosponding source objects mapped with replica objects provided at the time of initiating failover request. | [optional] 
**SourceClusterInfo** | Pointer to [**FailoverSourceCluster**](FailoverSourceCluster.md) |  | [optional] 

## Methods

### NewInitFailoverResponse

`func NewInitFailoverResponse() *InitFailoverResponse`

NewInitFailoverResponse instantiates a new InitFailoverResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitFailoverResponseWithDefaults

`func NewInitFailoverResponseWithDefaults() *InitFailoverResponse`

NewInitFailoverResponseWithDefaults instantiates a new InitFailoverResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReplicaToSourceObjects

`func (o *InitFailoverResponse) GetReplicaToSourceObjects() []SourceReplicaObject`

GetReplicaToSourceObjects returns the ReplicaToSourceObjects field if non-nil, zero value otherwise.

### GetReplicaToSourceObjectsOk

`func (o *InitFailoverResponse) GetReplicaToSourceObjectsOk() (*[]SourceReplicaObject, bool)`

GetReplicaToSourceObjectsOk returns a tuple with the ReplicaToSourceObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaToSourceObjects

`func (o *InitFailoverResponse) SetReplicaToSourceObjects(v []SourceReplicaObject)`

SetReplicaToSourceObjects sets ReplicaToSourceObjects field to given value.

### HasReplicaToSourceObjects

`func (o *InitFailoverResponse) HasReplicaToSourceObjects() bool`

HasReplicaToSourceObjects returns a boolean if a field has been set.

### SetReplicaToSourceObjectsNil

`func (o *InitFailoverResponse) SetReplicaToSourceObjectsNil(b bool)`

 SetReplicaToSourceObjectsNil sets the value for ReplicaToSourceObjects to be an explicit nil

### UnsetReplicaToSourceObjects
`func (o *InitFailoverResponse) UnsetReplicaToSourceObjects()`

UnsetReplicaToSourceObjects ensures that no value is present for ReplicaToSourceObjects, not even an explicit nil
### GetSourceClusterInfo

`func (o *InitFailoverResponse) GetSourceClusterInfo() FailoverSourceCluster`

GetSourceClusterInfo returns the SourceClusterInfo field if non-nil, zero value otherwise.

### GetSourceClusterInfoOk

`func (o *InitFailoverResponse) GetSourceClusterInfoOk() (*FailoverSourceCluster, bool)`

GetSourceClusterInfoOk returns a tuple with the SourceClusterInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceClusterInfo

`func (o *InitFailoverResponse) SetSourceClusterInfo(v FailoverSourceCluster)`

SetSourceClusterInfo sets SourceClusterInfo field to given value.

### HasSourceClusterInfo

`func (o *InitFailoverResponse) HasSourceClusterInfo() bool`

HasSourceClusterInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


