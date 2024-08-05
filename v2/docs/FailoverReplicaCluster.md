# FailoverReplicaCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | [**[]FailoverObject**](FailoverObject.md) | Specifies the details about the objects being failed over. In case if view based orchastrator is calling this then they should pass a object id for replicated view entity which belongs to the live tracking view on replication cluster. | 
**ProtectionGroupId** | Pointer to **NullableString** | Specifies the protection group id from the replication cluster from where the objects being failed over. If this is not specified then it will be infer from the list of objects being failed over. The protection group id must be specified in this format &lt;cluster_id&gt;:&lt;cluster_incarnation_id:jobid&gt; | [optional] 

## Methods

### NewFailoverReplicaCluster

`func NewFailoverReplicaCluster(objects []FailoverObject, ) *FailoverReplicaCluster`

NewFailoverReplicaCluster instantiates a new FailoverReplicaCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailoverReplicaClusterWithDefaults

`func NewFailoverReplicaClusterWithDefaults() *FailoverReplicaCluster`

NewFailoverReplicaClusterWithDefaults instantiates a new FailoverReplicaCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjects

`func (o *FailoverReplicaCluster) GetObjects() []FailoverObject`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *FailoverReplicaCluster) GetObjectsOk() (*[]FailoverObject, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *FailoverReplicaCluster) SetObjects(v []FailoverObject)`

SetObjects sets Objects field to given value.


### SetObjectsNil

`func (o *FailoverReplicaCluster) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *FailoverReplicaCluster) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetProtectionGroupId

`func (o *FailoverReplicaCluster) GetProtectionGroupId() string`

GetProtectionGroupId returns the ProtectionGroupId field if non-nil, zero value otherwise.

### GetProtectionGroupIdOk

`func (o *FailoverReplicaCluster) GetProtectionGroupIdOk() (*string, bool)`

GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupId

`func (o *FailoverReplicaCluster) SetProtectionGroupId(v string)`

SetProtectionGroupId sets ProtectionGroupId field to given value.

### HasProtectionGroupId

`func (o *FailoverReplicaCluster) HasProtectionGroupId() bool`

HasProtectionGroupId returns a boolean if a field has been set.

### SetProtectionGroupIdNil

`func (o *FailoverReplicaCluster) SetProtectionGroupIdNil(b bool)`

 SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil

### UnsetProtectionGroupId
`func (o *FailoverReplicaCluster) UnsetProtectionGroupId()`

UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


