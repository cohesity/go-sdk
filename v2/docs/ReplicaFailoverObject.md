# ReplicaFailoverObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailoverObjectId** | Pointer to **NullableInt64** | Specifies the corrosponding object id of the failover object. | [optional] 
**ReplicaObjectId** | Pointer to **NullableInt64** | Specifies the object Id existing on the replciation cluster. | [optional] 

## Methods

### NewReplicaFailoverObject

`func NewReplicaFailoverObject() *ReplicaFailoverObject`

NewReplicaFailoverObject instantiates a new ReplicaFailoverObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicaFailoverObjectWithDefaults

`func NewReplicaFailoverObjectWithDefaults() *ReplicaFailoverObject`

NewReplicaFailoverObjectWithDefaults instantiates a new ReplicaFailoverObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailoverObjectId

`func (o *ReplicaFailoverObject) GetFailoverObjectId() int64`

GetFailoverObjectId returns the FailoverObjectId field if non-nil, zero value otherwise.

### GetFailoverObjectIdOk

`func (o *ReplicaFailoverObject) GetFailoverObjectIdOk() (*int64, bool)`

GetFailoverObjectIdOk returns a tuple with the FailoverObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverObjectId

`func (o *ReplicaFailoverObject) SetFailoverObjectId(v int64)`

SetFailoverObjectId sets FailoverObjectId field to given value.

### HasFailoverObjectId

`func (o *ReplicaFailoverObject) HasFailoverObjectId() bool`

HasFailoverObjectId returns a boolean if a field has been set.

### SetFailoverObjectIdNil

`func (o *ReplicaFailoverObject) SetFailoverObjectIdNil(b bool)`

 SetFailoverObjectIdNil sets the value for FailoverObjectId to be an explicit nil

### UnsetFailoverObjectId
`func (o *ReplicaFailoverObject) UnsetFailoverObjectId()`

UnsetFailoverObjectId ensures that no value is present for FailoverObjectId, not even an explicit nil
### GetReplicaObjectId

`func (o *ReplicaFailoverObject) GetReplicaObjectId() int64`

GetReplicaObjectId returns the ReplicaObjectId field if non-nil, zero value otherwise.

### GetReplicaObjectIdOk

`func (o *ReplicaFailoverObject) GetReplicaObjectIdOk() (*int64, bool)`

GetReplicaObjectIdOk returns a tuple with the ReplicaObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaObjectId

`func (o *ReplicaFailoverObject) SetReplicaObjectId(v int64)`

SetReplicaObjectId sets ReplicaObjectId field to given value.

### HasReplicaObjectId

`func (o *ReplicaFailoverObject) HasReplicaObjectId() bool`

HasReplicaObjectId returns a boolean if a field has been set.

### SetReplicaObjectIdNil

`func (o *ReplicaFailoverObject) SetReplicaObjectIdNil(b bool)`

 SetReplicaObjectIdNil sets the value for ReplicaObjectId to be an explicit nil

### UnsetReplicaObjectId
`func (o *ReplicaFailoverObject) UnsetReplicaObjectId()`

UnsetReplicaObjectId ensures that no value is present for ReplicaObjectId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


