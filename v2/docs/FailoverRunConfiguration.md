# FailoverRunConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CancelNonFailoverRuns** | Pointer to **NullableBool** | If set to true, other ongoing runs backing up the same set of entities being failed over will be initiated for cancellation. Non conflicting run operations such as replications to other clusters, archivals will not be cancelled. If set to false, then new run will wait for all the pending operations to finish normally before scheduling a new backup/replication. | [optional] 
**Objects** | [**[]FailoverObject**](FailoverObject.md) | Specifies the list of all local entity ids of all the objects being failed from the source cluster. | 
**PauseNextRuns** | Pointer to **NullableBool** | If this is set to true then unless failover operation is completed, all the next runs will be pasued. | [optional] 
**ProtectionGroupId** | Pointer to **NullableString** | Specifies the active protection group id on the source cluster from where the objects are being failed over. | [optional] 
**ReplicationClusterId** | **NullableInt64** | Specifies the replication cluster Id where planned run will replicate objects. | 
**RunType** | Pointer to **string** | Specifies the type of the backup run to be triggered by this request. If this is not set defaults to incremental backup. | [optional] 
**ViewId** | Pointer to **NullableInt64** | If failover is initiated by view based orchastrator, then this field specifies the local view id of source cluster which is being failed over. | [optional] 

## Methods

### NewFailoverRunConfiguration

`func NewFailoverRunConfiguration(objects []FailoverObject, replicationClusterId NullableInt64, ) *FailoverRunConfiguration`

NewFailoverRunConfiguration instantiates a new FailoverRunConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailoverRunConfigurationWithDefaults

`func NewFailoverRunConfigurationWithDefaults() *FailoverRunConfiguration`

NewFailoverRunConfigurationWithDefaults instantiates a new FailoverRunConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancelNonFailoverRuns

`func (o *FailoverRunConfiguration) GetCancelNonFailoverRuns() bool`

GetCancelNonFailoverRuns returns the CancelNonFailoverRuns field if non-nil, zero value otherwise.

### GetCancelNonFailoverRunsOk

`func (o *FailoverRunConfiguration) GetCancelNonFailoverRunsOk() (*bool, bool)`

GetCancelNonFailoverRunsOk returns a tuple with the CancelNonFailoverRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelNonFailoverRuns

`func (o *FailoverRunConfiguration) SetCancelNonFailoverRuns(v bool)`

SetCancelNonFailoverRuns sets CancelNonFailoverRuns field to given value.

### HasCancelNonFailoverRuns

`func (o *FailoverRunConfiguration) HasCancelNonFailoverRuns() bool`

HasCancelNonFailoverRuns returns a boolean if a field has been set.

### SetCancelNonFailoverRunsNil

`func (o *FailoverRunConfiguration) SetCancelNonFailoverRunsNil(b bool)`

 SetCancelNonFailoverRunsNil sets the value for CancelNonFailoverRuns to be an explicit nil

### UnsetCancelNonFailoverRuns
`func (o *FailoverRunConfiguration) UnsetCancelNonFailoverRuns()`

UnsetCancelNonFailoverRuns ensures that no value is present for CancelNonFailoverRuns, not even an explicit nil
### GetObjects

`func (o *FailoverRunConfiguration) GetObjects() []FailoverObject`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *FailoverRunConfiguration) GetObjectsOk() (*[]FailoverObject, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *FailoverRunConfiguration) SetObjects(v []FailoverObject)`

SetObjects sets Objects field to given value.


### SetObjectsNil

`func (o *FailoverRunConfiguration) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *FailoverRunConfiguration) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetPauseNextRuns

`func (o *FailoverRunConfiguration) GetPauseNextRuns() bool`

GetPauseNextRuns returns the PauseNextRuns field if non-nil, zero value otherwise.

### GetPauseNextRunsOk

`func (o *FailoverRunConfiguration) GetPauseNextRunsOk() (*bool, bool)`

GetPauseNextRunsOk returns a tuple with the PauseNextRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseNextRuns

`func (o *FailoverRunConfiguration) SetPauseNextRuns(v bool)`

SetPauseNextRuns sets PauseNextRuns field to given value.

### HasPauseNextRuns

`func (o *FailoverRunConfiguration) HasPauseNextRuns() bool`

HasPauseNextRuns returns a boolean if a field has been set.

### SetPauseNextRunsNil

`func (o *FailoverRunConfiguration) SetPauseNextRunsNil(b bool)`

 SetPauseNextRunsNil sets the value for PauseNextRuns to be an explicit nil

### UnsetPauseNextRuns
`func (o *FailoverRunConfiguration) UnsetPauseNextRuns()`

UnsetPauseNextRuns ensures that no value is present for PauseNextRuns, not even an explicit nil
### GetProtectionGroupId

`func (o *FailoverRunConfiguration) GetProtectionGroupId() string`

GetProtectionGroupId returns the ProtectionGroupId field if non-nil, zero value otherwise.

### GetProtectionGroupIdOk

`func (o *FailoverRunConfiguration) GetProtectionGroupIdOk() (*string, bool)`

GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupId

`func (o *FailoverRunConfiguration) SetProtectionGroupId(v string)`

SetProtectionGroupId sets ProtectionGroupId field to given value.

### HasProtectionGroupId

`func (o *FailoverRunConfiguration) HasProtectionGroupId() bool`

HasProtectionGroupId returns a boolean if a field has been set.

### SetProtectionGroupIdNil

`func (o *FailoverRunConfiguration) SetProtectionGroupIdNil(b bool)`

 SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil

### UnsetProtectionGroupId
`func (o *FailoverRunConfiguration) UnsetProtectionGroupId()`

UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil
### GetReplicationClusterId

`func (o *FailoverRunConfiguration) GetReplicationClusterId() int64`

GetReplicationClusterId returns the ReplicationClusterId field if non-nil, zero value otherwise.

### GetReplicationClusterIdOk

`func (o *FailoverRunConfiguration) GetReplicationClusterIdOk() (*int64, bool)`

GetReplicationClusterIdOk returns a tuple with the ReplicationClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationClusterId

`func (o *FailoverRunConfiguration) SetReplicationClusterId(v int64)`

SetReplicationClusterId sets ReplicationClusterId field to given value.


### SetReplicationClusterIdNil

`func (o *FailoverRunConfiguration) SetReplicationClusterIdNil(b bool)`

 SetReplicationClusterIdNil sets the value for ReplicationClusterId to be an explicit nil

### UnsetReplicationClusterId
`func (o *FailoverRunConfiguration) UnsetReplicationClusterId()`

UnsetReplicationClusterId ensures that no value is present for ReplicationClusterId, not even an explicit nil
### GetRunType

`func (o *FailoverRunConfiguration) GetRunType() string`

GetRunType returns the RunType field if non-nil, zero value otherwise.

### GetRunTypeOk

`func (o *FailoverRunConfiguration) GetRunTypeOk() (*string, bool)`

GetRunTypeOk returns a tuple with the RunType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunType

`func (o *FailoverRunConfiguration) SetRunType(v string)`

SetRunType sets RunType field to given value.

### HasRunType

`func (o *FailoverRunConfiguration) HasRunType() bool`

HasRunType returns a boolean if a field has been set.

### GetViewId

`func (o *FailoverRunConfiguration) GetViewId() int64`

GetViewId returns the ViewId field if non-nil, zero value otherwise.

### GetViewIdOk

`func (o *FailoverRunConfiguration) GetViewIdOk() (*int64, bool)`

GetViewIdOk returns a tuple with the ViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewId

`func (o *FailoverRunConfiguration) SetViewId(v int64)`

SetViewId sets ViewId field to given value.

### HasViewId

`func (o *FailoverRunConfiguration) HasViewId() bool`

HasViewId returns a boolean if a field has been set.

### SetViewIdNil

`func (o *FailoverRunConfiguration) SetViewIdNil(b bool)`

 SetViewIdNil sets the value for ViewId to be an explicit nil

### UnsetViewId
`func (o *FailoverRunConfiguration) UnsetViewId()`

UnsetViewId ensures that no value is present for ViewId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


