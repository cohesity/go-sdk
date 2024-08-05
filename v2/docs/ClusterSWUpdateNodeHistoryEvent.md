# ClusterSWUpdateNodeHistoryEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventTimestampSecs** | Pointer to **NullableInt64** | Unix epoch timestamp (in seconds) when the upgrade/patch was done on the node.  | [optional] 
**NodeId** | Pointer to **int64** | Node id of the node. | [optional] 

## Methods

### NewClusterSWUpdateNodeHistoryEvent

`func NewClusterSWUpdateNodeHistoryEvent() *ClusterSWUpdateNodeHistoryEvent`

NewClusterSWUpdateNodeHistoryEvent instantiates a new ClusterSWUpdateNodeHistoryEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterSWUpdateNodeHistoryEventWithDefaults

`func NewClusterSWUpdateNodeHistoryEventWithDefaults() *ClusterSWUpdateNodeHistoryEvent`

NewClusterSWUpdateNodeHistoryEventWithDefaults instantiates a new ClusterSWUpdateNodeHistoryEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventTimestampSecs

`func (o *ClusterSWUpdateNodeHistoryEvent) GetEventTimestampSecs() int64`

GetEventTimestampSecs returns the EventTimestampSecs field if non-nil, zero value otherwise.

### GetEventTimestampSecsOk

`func (o *ClusterSWUpdateNodeHistoryEvent) GetEventTimestampSecsOk() (*int64, bool)`

GetEventTimestampSecsOk returns a tuple with the EventTimestampSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTimestampSecs

`func (o *ClusterSWUpdateNodeHistoryEvent) SetEventTimestampSecs(v int64)`

SetEventTimestampSecs sets EventTimestampSecs field to given value.

### HasEventTimestampSecs

`func (o *ClusterSWUpdateNodeHistoryEvent) HasEventTimestampSecs() bool`

HasEventTimestampSecs returns a boolean if a field has been set.

### SetEventTimestampSecsNil

`func (o *ClusterSWUpdateNodeHistoryEvent) SetEventTimestampSecsNil(b bool)`

 SetEventTimestampSecsNil sets the value for EventTimestampSecs to be an explicit nil

### UnsetEventTimestampSecs
`func (o *ClusterSWUpdateNodeHistoryEvent) UnsetEventTimestampSecs()`

UnsetEventTimestampSecs ensures that no value is present for EventTimestampSecs, not even an explicit nil
### GetNodeId

`func (o *ClusterSWUpdateNodeHistoryEvent) GetNodeId() int64`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ClusterSWUpdateNodeHistoryEvent) GetNodeIdOk() (*int64, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ClusterSWUpdateNodeHistoryEvent) SetNodeId(v int64)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *ClusterSWUpdateNodeHistoryEvent) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


