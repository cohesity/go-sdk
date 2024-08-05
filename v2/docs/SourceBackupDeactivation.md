# SourceBackupDeactivation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeepFailoverObjects** | Pointer to **NullableBool** | If this is set to true then objects will not be removed from protection group. If this is set to false, then all objects which are being failed over will be removed from the protection group. If protection group left with zero entities then it will be paused automatically. | [optional] 
**Objects** | Pointer to [**[]FailoverObject**](FailoverObject.md) | Specifies the list of all local entity ids of all the objects being failed from the source cluster. Backup will be deactiaved for all given objects. | [optional] 
**ProtectionGroupId** | Pointer to **NullableString** | Specifies the protection group id of the source cluster from where the objects being failed over. If this is not specified then it will be infer from the list of objects being failed over. | [optional] 
**ReplicationClusterId** | Pointer to **NullableInt64** | Specifies the replication cluster Id involved in failover operation. | [optional] 
**ViewId** | Pointer to **NullableString** | If failover is initiated by view based orchastrator, then this field specifies the local view id of source cluster which is being failed over. Backup will be deactivated for view object. | [optional] [readonly] 

## Methods

### NewSourceBackupDeactivation

`func NewSourceBackupDeactivation() *SourceBackupDeactivation`

NewSourceBackupDeactivation instantiates a new SourceBackupDeactivation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceBackupDeactivationWithDefaults

`func NewSourceBackupDeactivationWithDefaults() *SourceBackupDeactivation`

NewSourceBackupDeactivationWithDefaults instantiates a new SourceBackupDeactivation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeepFailoverObjects

`func (o *SourceBackupDeactivation) GetKeepFailoverObjects() bool`

GetKeepFailoverObjects returns the KeepFailoverObjects field if non-nil, zero value otherwise.

### GetKeepFailoverObjectsOk

`func (o *SourceBackupDeactivation) GetKeepFailoverObjectsOk() (*bool, bool)`

GetKeepFailoverObjectsOk returns a tuple with the KeepFailoverObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepFailoverObjects

`func (o *SourceBackupDeactivation) SetKeepFailoverObjects(v bool)`

SetKeepFailoverObjects sets KeepFailoverObjects field to given value.

### HasKeepFailoverObjects

`func (o *SourceBackupDeactivation) HasKeepFailoverObjects() bool`

HasKeepFailoverObjects returns a boolean if a field has been set.

### SetKeepFailoverObjectsNil

`func (o *SourceBackupDeactivation) SetKeepFailoverObjectsNil(b bool)`

 SetKeepFailoverObjectsNil sets the value for KeepFailoverObjects to be an explicit nil

### UnsetKeepFailoverObjects
`func (o *SourceBackupDeactivation) UnsetKeepFailoverObjects()`

UnsetKeepFailoverObjects ensures that no value is present for KeepFailoverObjects, not even an explicit nil
### GetObjects

`func (o *SourceBackupDeactivation) GetObjects() []FailoverObject`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *SourceBackupDeactivation) GetObjectsOk() (*[]FailoverObject, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *SourceBackupDeactivation) SetObjects(v []FailoverObject)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *SourceBackupDeactivation) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### SetObjectsNil

`func (o *SourceBackupDeactivation) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *SourceBackupDeactivation) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetProtectionGroupId

`func (o *SourceBackupDeactivation) GetProtectionGroupId() string`

GetProtectionGroupId returns the ProtectionGroupId field if non-nil, zero value otherwise.

### GetProtectionGroupIdOk

`func (o *SourceBackupDeactivation) GetProtectionGroupIdOk() (*string, bool)`

GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupId

`func (o *SourceBackupDeactivation) SetProtectionGroupId(v string)`

SetProtectionGroupId sets ProtectionGroupId field to given value.

### HasProtectionGroupId

`func (o *SourceBackupDeactivation) HasProtectionGroupId() bool`

HasProtectionGroupId returns a boolean if a field has been set.

### SetProtectionGroupIdNil

`func (o *SourceBackupDeactivation) SetProtectionGroupIdNil(b bool)`

 SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil

### UnsetProtectionGroupId
`func (o *SourceBackupDeactivation) UnsetProtectionGroupId()`

UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil
### GetReplicationClusterId

`func (o *SourceBackupDeactivation) GetReplicationClusterId() int64`

GetReplicationClusterId returns the ReplicationClusterId field if non-nil, zero value otherwise.

### GetReplicationClusterIdOk

`func (o *SourceBackupDeactivation) GetReplicationClusterIdOk() (*int64, bool)`

GetReplicationClusterIdOk returns a tuple with the ReplicationClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationClusterId

`func (o *SourceBackupDeactivation) SetReplicationClusterId(v int64)`

SetReplicationClusterId sets ReplicationClusterId field to given value.

### HasReplicationClusterId

`func (o *SourceBackupDeactivation) HasReplicationClusterId() bool`

HasReplicationClusterId returns a boolean if a field has been set.

### SetReplicationClusterIdNil

`func (o *SourceBackupDeactivation) SetReplicationClusterIdNil(b bool)`

 SetReplicationClusterIdNil sets the value for ReplicationClusterId to be an explicit nil

### UnsetReplicationClusterId
`func (o *SourceBackupDeactivation) UnsetReplicationClusterId()`

UnsetReplicationClusterId ensures that no value is present for ReplicationClusterId, not even an explicit nil
### GetViewId

`func (o *SourceBackupDeactivation) GetViewId() string`

GetViewId returns the ViewId field if non-nil, zero value otherwise.

### GetViewIdOk

`func (o *SourceBackupDeactivation) GetViewIdOk() (*string, bool)`

GetViewIdOk returns a tuple with the ViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewId

`func (o *SourceBackupDeactivation) SetViewId(v string)`

SetViewId sets ViewId field to given value.

### HasViewId

`func (o *SourceBackupDeactivation) HasViewId() bool`

HasViewId returns a boolean if a field has been set.

### SetViewIdNil

`func (o *SourceBackupDeactivation) SetViewIdNil(b bool)`

 SetViewIdNil sets the value for ViewId to be an explicit nil

### UnsetViewId
`func (o *SourceBackupDeactivation) UnsetViewId()`

UnsetViewId ensures that no value is present for ViewId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


