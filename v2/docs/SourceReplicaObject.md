# SourceReplicaObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReplicaObjectId** | Pointer to **NullableInt64** | Specifies the object Id existing on the replciation cluster. | [optional] 
**SourceObjectId** | Pointer to **NullableInt64** | Specifies the corrosponding object id existing on the source cluster. | [optional] 

## Methods

### NewSourceReplicaObject

`func NewSourceReplicaObject() *SourceReplicaObject`

NewSourceReplicaObject instantiates a new SourceReplicaObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceReplicaObjectWithDefaults

`func NewSourceReplicaObjectWithDefaults() *SourceReplicaObject`

NewSourceReplicaObjectWithDefaults instantiates a new SourceReplicaObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReplicaObjectId

`func (o *SourceReplicaObject) GetReplicaObjectId() int64`

GetReplicaObjectId returns the ReplicaObjectId field if non-nil, zero value otherwise.

### GetReplicaObjectIdOk

`func (o *SourceReplicaObject) GetReplicaObjectIdOk() (*int64, bool)`

GetReplicaObjectIdOk returns a tuple with the ReplicaObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaObjectId

`func (o *SourceReplicaObject) SetReplicaObjectId(v int64)`

SetReplicaObjectId sets ReplicaObjectId field to given value.

### HasReplicaObjectId

`func (o *SourceReplicaObject) HasReplicaObjectId() bool`

HasReplicaObjectId returns a boolean if a field has been set.

### SetReplicaObjectIdNil

`func (o *SourceReplicaObject) SetReplicaObjectIdNil(b bool)`

 SetReplicaObjectIdNil sets the value for ReplicaObjectId to be an explicit nil

### UnsetReplicaObjectId
`func (o *SourceReplicaObject) UnsetReplicaObjectId()`

UnsetReplicaObjectId ensures that no value is present for ReplicaObjectId, not even an explicit nil
### GetSourceObjectId

`func (o *SourceReplicaObject) GetSourceObjectId() int64`

GetSourceObjectId returns the SourceObjectId field if non-nil, zero value otherwise.

### GetSourceObjectIdOk

`func (o *SourceReplicaObject) GetSourceObjectIdOk() (*int64, bool)`

GetSourceObjectIdOk returns a tuple with the SourceObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceObjectId

`func (o *SourceReplicaObject) SetSourceObjectId(v int64)`

SetSourceObjectId sets SourceObjectId field to given value.

### HasSourceObjectId

`func (o *SourceReplicaObject) HasSourceObjectId() bool`

HasSourceObjectId returns a boolean if a field has been set.

### SetSourceObjectIdNil

`func (o *SourceReplicaObject) SetSourceObjectIdNil(b bool)`

 SetSourceObjectIdNil sets the value for SourceObjectId to be an explicit nil

### UnsetSourceObjectId
`func (o *SourceReplicaObject) UnsetSourceObjectId()`

UnsetSourceObjectId ensures that no value is present for SourceObjectId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


