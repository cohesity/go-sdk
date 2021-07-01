# UniversalId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **NullableInt64** | Specifies the Cohesity Cluster id where the object was created. | [optional] 
**ClusterIncarnationId** | Pointer to **NullableInt64** | Specifies an id for the Cohesity Cluster that is generated when a Cohesity Cluster is initially created. | [optional] 
**Id** | Pointer to **NullableInt64** | Specifies a unique id assigned to an object (such as a Job) by the Cohesity Cluster. | [optional] 

## Methods

### NewUniversalId

`func NewUniversalId() *UniversalId`

NewUniversalId instantiates a new UniversalId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniversalIdWithDefaults

`func NewUniversalIdWithDefaults() *UniversalId`

NewUniversalIdWithDefaults instantiates a new UniversalId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *UniversalId) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *UniversalId) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *UniversalId) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *UniversalId) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### SetClusterIdNil

`func (o *UniversalId) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *UniversalId) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetClusterIncarnationId

`func (o *UniversalId) GetClusterIncarnationId() int64`

GetClusterIncarnationId returns the ClusterIncarnationId field if non-nil, zero value otherwise.

### GetClusterIncarnationIdOk

`func (o *UniversalId) GetClusterIncarnationIdOk() (*int64, bool)`

GetClusterIncarnationIdOk returns a tuple with the ClusterIncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIncarnationId

`func (o *UniversalId) SetClusterIncarnationId(v int64)`

SetClusterIncarnationId sets ClusterIncarnationId field to given value.

### HasClusterIncarnationId

`func (o *UniversalId) HasClusterIncarnationId() bool`

HasClusterIncarnationId returns a boolean if a field has been set.

### SetClusterIncarnationIdNil

`func (o *UniversalId) SetClusterIncarnationIdNil(b bool)`

 SetClusterIncarnationIdNil sets the value for ClusterIncarnationId to be an explicit nil

### UnsetClusterIncarnationId
`func (o *UniversalId) UnsetClusterIncarnationId()`

UnsetClusterIncarnationId ensures that no value is present for ClusterIncarnationId, not even an explicit nil
### GetId

`func (o *UniversalId) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UniversalId) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UniversalId) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UniversalId) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *UniversalId) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *UniversalId) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


