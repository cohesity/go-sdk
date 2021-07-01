# ClusterInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **NullableInt64** | Specifies the id of the cluster. | [optional] 
**IncarnationId** | Pointer to **NullableInt64** | Specifies the incarnation id of the cluster. | [optional] 

## Methods

### NewClusterInfo

`func NewClusterInfo() *ClusterInfo`

NewClusterInfo instantiates a new ClusterInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterInfoWithDefaults

`func NewClusterInfoWithDefaults() *ClusterInfo`

NewClusterInfoWithDefaults instantiates a new ClusterInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *ClusterInfo) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ClusterInfo) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ClusterInfo) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *ClusterInfo) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### SetClusterIdNil

`func (o *ClusterInfo) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *ClusterInfo) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetIncarnationId

`func (o *ClusterInfo) GetIncarnationId() int64`

GetIncarnationId returns the IncarnationId field if non-nil, zero value otherwise.

### GetIncarnationIdOk

`func (o *ClusterInfo) GetIncarnationIdOk() (*int64, bool)`

GetIncarnationIdOk returns a tuple with the IncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncarnationId

`func (o *ClusterInfo) SetIncarnationId(v int64)`

SetIncarnationId sets IncarnationId field to given value.

### HasIncarnationId

`func (o *ClusterInfo) HasIncarnationId() bool`

HasIncarnationId returns a boolean if a field has been set.

### SetIncarnationIdNil

`func (o *ClusterInfo) SetIncarnationIdNil(b bool)`

 SetIncarnationIdNil sets the value for IncarnationId to be an explicit nil

### UnsetIncarnationId
`func (o *ClusterInfo) UnsetIncarnationId()`

UnsetIncarnationId ensures that no value is present for IncarnationId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


