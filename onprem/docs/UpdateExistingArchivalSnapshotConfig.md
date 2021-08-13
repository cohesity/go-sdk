# UpdateExistingArchivalSnapshotConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Specifies the id of the archival target. | 
**ArchivalTargetType** | **NullableString** | Specifies the snapshot&#39;s archival target type from which recovery has been performed. | 
**EnableLegalHold** | Pointer to **NullableBool** | Specifies whether to retain the snapshot for legal purpose. If set to true, the snapshots cannot be deleted until the retention period. Note that using this option may cause the Cluster to run out of space. If set to false explicitly, the hold is removed, and the snapshots will expire as specified in the policy of the Protection Group. If this field is not specified, there is no change to the hold of the run. This field can be set only by a User having Data Security Role. | [optional] 
**DeleteSnapshot** | Pointer to **NullableBool** | Specifies whether to delete the snapshot. When this is set to true, all other params will be ignored. | [optional] 
**Resync** | Pointer to **NullableBool** | Specifies whether to retry the archival operation in case if earlier attempt failed. If not specified or set to false, archival is not retried. | [optional] 
**DataLock** | Pointer to **NullableString** | Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept until the maximum of the snapshot retention time. During that time, the snapshots cannot be deleted. &lt;br&gt;&#39;Compliance&#39; implies WORM retention is set for compliance reason. &lt;br&gt;&#39;Administrative&#39; implies WORM retention is set for administrative purposes. | [optional] 
**DaysToKeep** | Pointer to **NullableInt64** | Specifies number of days to retain the snapshots. If positive, then this value is added to exisiting expiry time thereby increasing  the retention period of the snapshot. Conversly, if this value is negative, then value is subtracted to existing expiry time thereby decreasing the retention period of the snaphot. Here, by this operation if expiry time goes below current time then snapshot is immediately deleted. | [optional] 

## Methods

### NewUpdateExistingArchivalSnapshotConfig

`func NewUpdateExistingArchivalSnapshotConfig(id int64, archivalTargetType NullableString, ) *UpdateExistingArchivalSnapshotConfig`

NewUpdateExistingArchivalSnapshotConfig instantiates a new UpdateExistingArchivalSnapshotConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateExistingArchivalSnapshotConfigWithDefaults

`func NewUpdateExistingArchivalSnapshotConfigWithDefaults() *UpdateExistingArchivalSnapshotConfig`

NewUpdateExistingArchivalSnapshotConfigWithDefaults instantiates a new UpdateExistingArchivalSnapshotConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateExistingArchivalSnapshotConfig) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateExistingArchivalSnapshotConfig) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateExistingArchivalSnapshotConfig) SetId(v int64)`

SetId sets Id field to given value.


### GetArchivalTargetType

`func (o *UpdateExistingArchivalSnapshotConfig) GetArchivalTargetType() string`

GetArchivalTargetType returns the ArchivalTargetType field if non-nil, zero value otherwise.

### GetArchivalTargetTypeOk

`func (o *UpdateExistingArchivalSnapshotConfig) GetArchivalTargetTypeOk() (*string, bool)`

GetArchivalTargetTypeOk returns a tuple with the ArchivalTargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivalTargetType

`func (o *UpdateExistingArchivalSnapshotConfig) SetArchivalTargetType(v string)`

SetArchivalTargetType sets ArchivalTargetType field to given value.


### SetArchivalTargetTypeNil

`func (o *UpdateExistingArchivalSnapshotConfig) SetArchivalTargetTypeNil(b bool)`

 SetArchivalTargetTypeNil sets the value for ArchivalTargetType to be an explicit nil

### UnsetArchivalTargetType
`func (o *UpdateExistingArchivalSnapshotConfig) UnsetArchivalTargetType()`

UnsetArchivalTargetType ensures that no value is present for ArchivalTargetType, not even an explicit nil
### GetEnableLegalHold

`func (o *UpdateExistingArchivalSnapshotConfig) GetEnableLegalHold() bool`

GetEnableLegalHold returns the EnableLegalHold field if non-nil, zero value otherwise.

### GetEnableLegalHoldOk

`func (o *UpdateExistingArchivalSnapshotConfig) GetEnableLegalHoldOk() (*bool, bool)`

GetEnableLegalHoldOk returns a tuple with the EnableLegalHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLegalHold

`func (o *UpdateExistingArchivalSnapshotConfig) SetEnableLegalHold(v bool)`

SetEnableLegalHold sets EnableLegalHold field to given value.

### HasEnableLegalHold

`func (o *UpdateExistingArchivalSnapshotConfig) HasEnableLegalHold() bool`

HasEnableLegalHold returns a boolean if a field has been set.

### SetEnableLegalHoldNil

`func (o *UpdateExistingArchivalSnapshotConfig) SetEnableLegalHoldNil(b bool)`

 SetEnableLegalHoldNil sets the value for EnableLegalHold to be an explicit nil

### UnsetEnableLegalHold
`func (o *UpdateExistingArchivalSnapshotConfig) UnsetEnableLegalHold()`

UnsetEnableLegalHold ensures that no value is present for EnableLegalHold, not even an explicit nil
### GetDeleteSnapshot

`func (o *UpdateExistingArchivalSnapshotConfig) GetDeleteSnapshot() bool`

GetDeleteSnapshot returns the DeleteSnapshot field if non-nil, zero value otherwise.

### GetDeleteSnapshotOk

`func (o *UpdateExistingArchivalSnapshotConfig) GetDeleteSnapshotOk() (*bool, bool)`

GetDeleteSnapshotOk returns a tuple with the DeleteSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteSnapshot

`func (o *UpdateExistingArchivalSnapshotConfig) SetDeleteSnapshot(v bool)`

SetDeleteSnapshot sets DeleteSnapshot field to given value.

### HasDeleteSnapshot

`func (o *UpdateExistingArchivalSnapshotConfig) HasDeleteSnapshot() bool`

HasDeleteSnapshot returns a boolean if a field has been set.

### SetDeleteSnapshotNil

`func (o *UpdateExistingArchivalSnapshotConfig) SetDeleteSnapshotNil(b bool)`

 SetDeleteSnapshotNil sets the value for DeleteSnapshot to be an explicit nil

### UnsetDeleteSnapshot
`func (o *UpdateExistingArchivalSnapshotConfig) UnsetDeleteSnapshot()`

UnsetDeleteSnapshot ensures that no value is present for DeleteSnapshot, not even an explicit nil
### GetResync

`func (o *UpdateExistingArchivalSnapshotConfig) GetResync() bool`

GetResync returns the Resync field if non-nil, zero value otherwise.

### GetResyncOk

`func (o *UpdateExistingArchivalSnapshotConfig) GetResyncOk() (*bool, bool)`

GetResyncOk returns a tuple with the Resync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResync

`func (o *UpdateExistingArchivalSnapshotConfig) SetResync(v bool)`

SetResync sets Resync field to given value.

### HasResync

`func (o *UpdateExistingArchivalSnapshotConfig) HasResync() bool`

HasResync returns a boolean if a field has been set.

### SetResyncNil

`func (o *UpdateExistingArchivalSnapshotConfig) SetResyncNil(b bool)`

 SetResyncNil sets the value for Resync to be an explicit nil

### UnsetResync
`func (o *UpdateExistingArchivalSnapshotConfig) UnsetResync()`

UnsetResync ensures that no value is present for Resync, not even an explicit nil
### GetDataLock

`func (o *UpdateExistingArchivalSnapshotConfig) GetDataLock() string`

GetDataLock returns the DataLock field if non-nil, zero value otherwise.

### GetDataLockOk

`func (o *UpdateExistingArchivalSnapshotConfig) GetDataLockOk() (*string, bool)`

GetDataLockOk returns a tuple with the DataLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLock

`func (o *UpdateExistingArchivalSnapshotConfig) SetDataLock(v string)`

SetDataLock sets DataLock field to given value.

### HasDataLock

`func (o *UpdateExistingArchivalSnapshotConfig) HasDataLock() bool`

HasDataLock returns a boolean if a field has been set.

### SetDataLockNil

`func (o *UpdateExistingArchivalSnapshotConfig) SetDataLockNil(b bool)`

 SetDataLockNil sets the value for DataLock to be an explicit nil

### UnsetDataLock
`func (o *UpdateExistingArchivalSnapshotConfig) UnsetDataLock()`

UnsetDataLock ensures that no value is present for DataLock, not even an explicit nil
### GetDaysToKeep

`func (o *UpdateExistingArchivalSnapshotConfig) GetDaysToKeep() int64`

GetDaysToKeep returns the DaysToKeep field if non-nil, zero value otherwise.

### GetDaysToKeepOk

`func (o *UpdateExistingArchivalSnapshotConfig) GetDaysToKeepOk() (*int64, bool)`

GetDaysToKeepOk returns a tuple with the DaysToKeep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysToKeep

`func (o *UpdateExistingArchivalSnapshotConfig) SetDaysToKeep(v int64)`

SetDaysToKeep sets DaysToKeep field to given value.

### HasDaysToKeep

`func (o *UpdateExistingArchivalSnapshotConfig) HasDaysToKeep() bool`

HasDaysToKeep returns a boolean if a field has been set.

### SetDaysToKeepNil

`func (o *UpdateExistingArchivalSnapshotConfig) SetDaysToKeepNil(b bool)`

 SetDaysToKeepNil sets the value for DaysToKeep to be an explicit nil

### UnsetDaysToKeep
`func (o *UpdateExistingArchivalSnapshotConfig) UnsetDaysToKeep()`

UnsetDaysToKeep ensures that no value is present for DaysToKeep, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


