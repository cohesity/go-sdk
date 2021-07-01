# SnapshotTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchivalTarget** | Pointer to [**ArchivalTarget**](ArchivalTarget.md) |  | [optional] 
**CloudDeployTarget** | Pointer to [**CloudDeployTarget**](CloudDeployTarget.md) |  | [optional] 
**ReplicationTarget** | Pointer to [**ReplicationTarget**](ReplicationTarget.md) |  | [optional] 
**Type** | Pointer to **NullableInt32** | The type of snapshot target this proto represents. | [optional] 

## Methods

### NewSnapshotTarget

`func NewSnapshotTarget() *SnapshotTarget`

NewSnapshotTarget instantiates a new SnapshotTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotTargetWithDefaults

`func NewSnapshotTargetWithDefaults() *SnapshotTarget`

NewSnapshotTargetWithDefaults instantiates a new SnapshotTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchivalTarget

`func (o *SnapshotTarget) GetArchivalTarget() ArchivalTarget`

GetArchivalTarget returns the ArchivalTarget field if non-nil, zero value otherwise.

### GetArchivalTargetOk

`func (o *SnapshotTarget) GetArchivalTargetOk() (*ArchivalTarget, bool)`

GetArchivalTargetOk returns a tuple with the ArchivalTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivalTarget

`func (o *SnapshotTarget) SetArchivalTarget(v ArchivalTarget)`

SetArchivalTarget sets ArchivalTarget field to given value.

### HasArchivalTarget

`func (o *SnapshotTarget) HasArchivalTarget() bool`

HasArchivalTarget returns a boolean if a field has been set.

### GetCloudDeployTarget

`func (o *SnapshotTarget) GetCloudDeployTarget() CloudDeployTarget`

GetCloudDeployTarget returns the CloudDeployTarget field if non-nil, zero value otherwise.

### GetCloudDeployTargetOk

`func (o *SnapshotTarget) GetCloudDeployTargetOk() (*CloudDeployTarget, bool)`

GetCloudDeployTargetOk returns a tuple with the CloudDeployTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudDeployTarget

`func (o *SnapshotTarget) SetCloudDeployTarget(v CloudDeployTarget)`

SetCloudDeployTarget sets CloudDeployTarget field to given value.

### HasCloudDeployTarget

`func (o *SnapshotTarget) HasCloudDeployTarget() bool`

HasCloudDeployTarget returns a boolean if a field has been set.

### GetReplicationTarget

`func (o *SnapshotTarget) GetReplicationTarget() ReplicationTarget`

GetReplicationTarget returns the ReplicationTarget field if non-nil, zero value otherwise.

### GetReplicationTargetOk

`func (o *SnapshotTarget) GetReplicationTargetOk() (*ReplicationTarget, bool)`

GetReplicationTargetOk returns a tuple with the ReplicationTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationTarget

`func (o *SnapshotTarget) SetReplicationTarget(v ReplicationTarget)`

SetReplicationTarget sets ReplicationTarget field to given value.

### HasReplicationTarget

`func (o *SnapshotTarget) HasReplicationTarget() bool`

HasReplicationTarget returns a boolean if a field has been set.

### GetType

`func (o *SnapshotTarget) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SnapshotTarget) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SnapshotTarget) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *SnapshotTarget) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *SnapshotTarget) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *SnapshotTarget) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


