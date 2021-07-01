# SnapshotTargetSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchivalTarget** | Pointer to [**ArchivalExternalTarget**](ArchivalExternalTarget.md) |  | [optional] 
**CloudReplicationTarget** | Pointer to [**CloudDeployTargetDetails**](CloudDeployTargetDetails.md) |  | [optional] 
**ReplicationTarget** | Pointer to [**ReplicationTargetSettings**](ReplicationTargetSettings.md) |  | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of a Snapshot target such as &#39;kLocal&#39;, &#39;kRemote&#39; or &#39;kArchival&#39;. &#39;kLocal&#39; means the Snapshot is stored on a local Cohesity Cluster. &#39;kRemote&#39; means the Snapshot is stored on a Remote Cohesity Cluster. (It was copied to the Remote Cohesity Cluster using replication.) &#39;kArchival&#39; means the Snapshot is stored on a Archival External Target (such as Tape or AWS). &#39;kCloudDeploy&#39; means the Snapshot is stored on a Cloud platform. | [optional] 

## Methods

### NewSnapshotTargetSettings

`func NewSnapshotTargetSettings() *SnapshotTargetSettings`

NewSnapshotTargetSettings instantiates a new SnapshotTargetSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotTargetSettingsWithDefaults

`func NewSnapshotTargetSettingsWithDefaults() *SnapshotTargetSettings`

NewSnapshotTargetSettingsWithDefaults instantiates a new SnapshotTargetSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchivalTarget

`func (o *SnapshotTargetSettings) GetArchivalTarget() ArchivalExternalTarget`

GetArchivalTarget returns the ArchivalTarget field if non-nil, zero value otherwise.

### GetArchivalTargetOk

`func (o *SnapshotTargetSettings) GetArchivalTargetOk() (*ArchivalExternalTarget, bool)`

GetArchivalTargetOk returns a tuple with the ArchivalTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivalTarget

`func (o *SnapshotTargetSettings) SetArchivalTarget(v ArchivalExternalTarget)`

SetArchivalTarget sets ArchivalTarget field to given value.

### HasArchivalTarget

`func (o *SnapshotTargetSettings) HasArchivalTarget() bool`

HasArchivalTarget returns a boolean if a field has been set.

### GetCloudReplicationTarget

`func (o *SnapshotTargetSettings) GetCloudReplicationTarget() CloudDeployTargetDetails`

GetCloudReplicationTarget returns the CloudReplicationTarget field if non-nil, zero value otherwise.

### GetCloudReplicationTargetOk

`func (o *SnapshotTargetSettings) GetCloudReplicationTargetOk() (*CloudDeployTargetDetails, bool)`

GetCloudReplicationTargetOk returns a tuple with the CloudReplicationTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudReplicationTarget

`func (o *SnapshotTargetSettings) SetCloudReplicationTarget(v CloudDeployTargetDetails)`

SetCloudReplicationTarget sets CloudReplicationTarget field to given value.

### HasCloudReplicationTarget

`func (o *SnapshotTargetSettings) HasCloudReplicationTarget() bool`

HasCloudReplicationTarget returns a boolean if a field has been set.

### GetReplicationTarget

`func (o *SnapshotTargetSettings) GetReplicationTarget() ReplicationTargetSettings`

GetReplicationTarget returns the ReplicationTarget field if non-nil, zero value otherwise.

### GetReplicationTargetOk

`func (o *SnapshotTargetSettings) GetReplicationTargetOk() (*ReplicationTargetSettings, bool)`

GetReplicationTargetOk returns a tuple with the ReplicationTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationTarget

`func (o *SnapshotTargetSettings) SetReplicationTarget(v ReplicationTargetSettings)`

SetReplicationTarget sets ReplicationTarget field to given value.

### HasReplicationTarget

`func (o *SnapshotTargetSettings) HasReplicationTarget() bool`

HasReplicationTarget returns a boolean if a field has been set.

### GetType

`func (o *SnapshotTargetSettings) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SnapshotTargetSettings) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SnapshotTargetSettings) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SnapshotTargetSettings) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *SnapshotTargetSettings) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *SnapshotTargetSettings) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


