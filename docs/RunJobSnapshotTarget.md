# RunJobSnapshotTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchivalTarget** | Pointer to [**ArchivalExternalTarget**](ArchivalExternalTarget.md) |  | [optional] 
**CloudReplicationTarget** | Pointer to [**CloudDeployTargetDetails**](CloudDeployTargetDetails.md) |  | [optional] 
**DaysToKeep** | Pointer to **NullableInt64** | Specifies the number of days to retain copied Snapshots on the target. | [optional] 
**HoldForLegalPurpose** | Pointer to **NullableBool** | Specifies optionally whether to retain the snapshot for legal purpose. If set to true, the run cannot be deleted until the retention period. Note that using this option may cause the Cluster to run out of space. If set to false explicitly, the hold is removed, and the run will expire as specified in the policy of the Protection Job. If this field is not specified, there is no change to the hold of the run. This field can be set only by a User having Data Security Role. | [optional] 
**ReplicationTarget** | Pointer to [**ReplicationTargetSettings**](ReplicationTargetSettings.md) |  | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of a Snapshot target such as &#39;kLocal&#39;, &#39;kRemote&#39; or &#39;kArchival&#39;. &#39;kLocal&#39; means the Snapshot is stored on a local Cohesity Cluster. &#39;kRemote&#39; means the Snapshot is stored on a Remote Cohesity Cluster. (It was copied to the Remote Cohesity Cluster using replication.) &#39;kArchival&#39; means the Snapshot is stored on a Archival External Target (such as Tape or AWS). &#39;kCloudDeploy&#39; means the Snapshot is stored on a Cloud platform. | [optional] 

## Methods

### NewRunJobSnapshotTarget

`func NewRunJobSnapshotTarget() *RunJobSnapshotTarget`

NewRunJobSnapshotTarget instantiates a new RunJobSnapshotTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunJobSnapshotTargetWithDefaults

`func NewRunJobSnapshotTargetWithDefaults() *RunJobSnapshotTarget`

NewRunJobSnapshotTargetWithDefaults instantiates a new RunJobSnapshotTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchivalTarget

`func (o *RunJobSnapshotTarget) GetArchivalTarget() ArchivalExternalTarget`

GetArchivalTarget returns the ArchivalTarget field if non-nil, zero value otherwise.

### GetArchivalTargetOk

`func (o *RunJobSnapshotTarget) GetArchivalTargetOk() (*ArchivalExternalTarget, bool)`

GetArchivalTargetOk returns a tuple with the ArchivalTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivalTarget

`func (o *RunJobSnapshotTarget) SetArchivalTarget(v ArchivalExternalTarget)`

SetArchivalTarget sets ArchivalTarget field to given value.

### HasArchivalTarget

`func (o *RunJobSnapshotTarget) HasArchivalTarget() bool`

HasArchivalTarget returns a boolean if a field has been set.

### GetCloudReplicationTarget

`func (o *RunJobSnapshotTarget) GetCloudReplicationTarget() CloudDeployTargetDetails`

GetCloudReplicationTarget returns the CloudReplicationTarget field if non-nil, zero value otherwise.

### GetCloudReplicationTargetOk

`func (o *RunJobSnapshotTarget) GetCloudReplicationTargetOk() (*CloudDeployTargetDetails, bool)`

GetCloudReplicationTargetOk returns a tuple with the CloudReplicationTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudReplicationTarget

`func (o *RunJobSnapshotTarget) SetCloudReplicationTarget(v CloudDeployTargetDetails)`

SetCloudReplicationTarget sets CloudReplicationTarget field to given value.

### HasCloudReplicationTarget

`func (o *RunJobSnapshotTarget) HasCloudReplicationTarget() bool`

HasCloudReplicationTarget returns a boolean if a field has been set.

### GetDaysToKeep

`func (o *RunJobSnapshotTarget) GetDaysToKeep() int64`

GetDaysToKeep returns the DaysToKeep field if non-nil, zero value otherwise.

### GetDaysToKeepOk

`func (o *RunJobSnapshotTarget) GetDaysToKeepOk() (*int64, bool)`

GetDaysToKeepOk returns a tuple with the DaysToKeep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysToKeep

`func (o *RunJobSnapshotTarget) SetDaysToKeep(v int64)`

SetDaysToKeep sets DaysToKeep field to given value.

### HasDaysToKeep

`func (o *RunJobSnapshotTarget) HasDaysToKeep() bool`

HasDaysToKeep returns a boolean if a field has been set.

### SetDaysToKeepNil

`func (o *RunJobSnapshotTarget) SetDaysToKeepNil(b bool)`

 SetDaysToKeepNil sets the value for DaysToKeep to be an explicit nil

### UnsetDaysToKeep
`func (o *RunJobSnapshotTarget) UnsetDaysToKeep()`

UnsetDaysToKeep ensures that no value is present for DaysToKeep, not even an explicit nil
### GetHoldForLegalPurpose

`func (o *RunJobSnapshotTarget) GetHoldForLegalPurpose() bool`

GetHoldForLegalPurpose returns the HoldForLegalPurpose field if non-nil, zero value otherwise.

### GetHoldForLegalPurposeOk

`func (o *RunJobSnapshotTarget) GetHoldForLegalPurposeOk() (*bool, bool)`

GetHoldForLegalPurposeOk returns a tuple with the HoldForLegalPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldForLegalPurpose

`func (o *RunJobSnapshotTarget) SetHoldForLegalPurpose(v bool)`

SetHoldForLegalPurpose sets HoldForLegalPurpose field to given value.

### HasHoldForLegalPurpose

`func (o *RunJobSnapshotTarget) HasHoldForLegalPurpose() bool`

HasHoldForLegalPurpose returns a boolean if a field has been set.

### SetHoldForLegalPurposeNil

`func (o *RunJobSnapshotTarget) SetHoldForLegalPurposeNil(b bool)`

 SetHoldForLegalPurposeNil sets the value for HoldForLegalPurpose to be an explicit nil

### UnsetHoldForLegalPurpose
`func (o *RunJobSnapshotTarget) UnsetHoldForLegalPurpose()`

UnsetHoldForLegalPurpose ensures that no value is present for HoldForLegalPurpose, not even an explicit nil
### GetReplicationTarget

`func (o *RunJobSnapshotTarget) GetReplicationTarget() ReplicationTargetSettings`

GetReplicationTarget returns the ReplicationTarget field if non-nil, zero value otherwise.

### GetReplicationTargetOk

`func (o *RunJobSnapshotTarget) GetReplicationTargetOk() (*ReplicationTargetSettings, bool)`

GetReplicationTargetOk returns a tuple with the ReplicationTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationTarget

`func (o *RunJobSnapshotTarget) SetReplicationTarget(v ReplicationTargetSettings)`

SetReplicationTarget sets ReplicationTarget field to given value.

### HasReplicationTarget

`func (o *RunJobSnapshotTarget) HasReplicationTarget() bool`

HasReplicationTarget returns a boolean if a field has been set.

### GetType

`func (o *RunJobSnapshotTarget) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RunJobSnapshotTarget) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RunJobSnapshotTarget) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RunJobSnapshotTarget) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *RunJobSnapshotTarget) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *RunJobSnapshotTarget) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


