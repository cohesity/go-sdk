# UniversalIdProto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **NullableInt64** | The id of the cluster at which the object was created. | [optional] 
**ClusterIncarnationId** | Pointer to **NullableInt64** | The incarnation id of the above cluster. | [optional] 
**ObjectId** | Pointer to **NullableInt64** | The object id - this is unique within the above cluster. | [optional] 

## Methods

### NewUniversalIdProto

`func NewUniversalIdProto() *UniversalIdProto`

NewUniversalIdProto instantiates a new UniversalIdProto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniversalIdProtoWithDefaults

`func NewUniversalIdProtoWithDefaults() *UniversalIdProto`

NewUniversalIdProtoWithDefaults instantiates a new UniversalIdProto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *UniversalIdProto) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *UniversalIdProto) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *UniversalIdProto) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *UniversalIdProto) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### SetClusterIdNil

`func (o *UniversalIdProto) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *UniversalIdProto) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetClusterIncarnationId

`func (o *UniversalIdProto) GetClusterIncarnationId() int64`

GetClusterIncarnationId returns the ClusterIncarnationId field if non-nil, zero value otherwise.

### GetClusterIncarnationIdOk

`func (o *UniversalIdProto) GetClusterIncarnationIdOk() (*int64, bool)`

GetClusterIncarnationIdOk returns a tuple with the ClusterIncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIncarnationId

`func (o *UniversalIdProto) SetClusterIncarnationId(v int64)`

SetClusterIncarnationId sets ClusterIncarnationId field to given value.

### HasClusterIncarnationId

`func (o *UniversalIdProto) HasClusterIncarnationId() bool`

HasClusterIncarnationId returns a boolean if a field has been set.

### SetClusterIncarnationIdNil

`func (o *UniversalIdProto) SetClusterIncarnationIdNil(b bool)`

 SetClusterIncarnationIdNil sets the value for ClusterIncarnationId to be an explicit nil

### UnsetClusterIncarnationId
`func (o *UniversalIdProto) UnsetClusterIncarnationId()`

UnsetClusterIncarnationId ensures that no value is present for ClusterIncarnationId, not even an explicit nil
### GetObjectId

`func (o *UniversalIdProto) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *UniversalIdProto) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *UniversalIdProto) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *UniversalIdProto) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### SetObjectIdNil

`func (o *UniversalIdProto) SetObjectIdNil(b bool)`

 SetObjectIdNil sets the value for ObjectId to be an explicit nil

### UnsetObjectId
`func (o *UniversalIdProto) UnsetObjectId()`

UnsetObjectId ensures that no value is present for ObjectId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


