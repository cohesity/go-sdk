# ReplicaInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiryTimeUsecs** | Pointer to **NullableInt64** | Specifies the expiration time of the snapshot within the target location. | [optional] 
**SnapshotTargetSettings** | Pointer to [**SnapshotTargetSettings**](SnapshotTargetSettings.md) |  | [optional] 
**Uid** | Pointer to [**UniversalId**](UniversalId.md) |  | [optional] 

## Methods

### NewReplicaInfo

`func NewReplicaInfo() *ReplicaInfo`

NewReplicaInfo instantiates a new ReplicaInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicaInfoWithDefaults

`func NewReplicaInfoWithDefaults() *ReplicaInfo`

NewReplicaInfoWithDefaults instantiates a new ReplicaInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiryTimeUsecs

`func (o *ReplicaInfo) GetExpiryTimeUsecs() int64`

GetExpiryTimeUsecs returns the ExpiryTimeUsecs field if non-nil, zero value otherwise.

### GetExpiryTimeUsecsOk

`func (o *ReplicaInfo) GetExpiryTimeUsecsOk() (*int64, bool)`

GetExpiryTimeUsecsOk returns a tuple with the ExpiryTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTimeUsecs

`func (o *ReplicaInfo) SetExpiryTimeUsecs(v int64)`

SetExpiryTimeUsecs sets ExpiryTimeUsecs field to given value.

### HasExpiryTimeUsecs

`func (o *ReplicaInfo) HasExpiryTimeUsecs() bool`

HasExpiryTimeUsecs returns a boolean if a field has been set.

### SetExpiryTimeUsecsNil

`func (o *ReplicaInfo) SetExpiryTimeUsecsNil(b bool)`

 SetExpiryTimeUsecsNil sets the value for ExpiryTimeUsecs to be an explicit nil

### UnsetExpiryTimeUsecs
`func (o *ReplicaInfo) UnsetExpiryTimeUsecs()`

UnsetExpiryTimeUsecs ensures that no value is present for ExpiryTimeUsecs, not even an explicit nil
### GetSnapshotTargetSettings

`func (o *ReplicaInfo) GetSnapshotTargetSettings() SnapshotTargetSettings`

GetSnapshotTargetSettings returns the SnapshotTargetSettings field if non-nil, zero value otherwise.

### GetSnapshotTargetSettingsOk

`func (o *ReplicaInfo) GetSnapshotTargetSettingsOk() (*SnapshotTargetSettings, bool)`

GetSnapshotTargetSettingsOk returns a tuple with the SnapshotTargetSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotTargetSettings

`func (o *ReplicaInfo) SetSnapshotTargetSettings(v SnapshotTargetSettings)`

SetSnapshotTargetSettings sets SnapshotTargetSettings field to given value.

### HasSnapshotTargetSettings

`func (o *ReplicaInfo) HasSnapshotTargetSettings() bool`

HasSnapshotTargetSettings returns a boolean if a field has been set.

### GetUid

`func (o *ReplicaInfo) GetUid() UniversalId`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ReplicaInfo) GetUidOk() (*UniversalId, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ReplicaInfo) SetUid(v UniversalId)`

SetUid sets Uid field to given value.

### HasUid

`func (o *ReplicaInfo) HasUid() bool`

HasUid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


