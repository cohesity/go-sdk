# DowntieringPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QosPolicy** | Pointer to **NullableString** | Specifies whether the data tiering task will be written to HDD or SSD. | [optional] 
**IndexingPolicy** | Pointer to [**IndexingPolicy**](IndexingPolicy.md) |  | [optional] 
**Retention** | Pointer to [**Retention**](Retention.md) |  | [optional] 
**SkipBackSymlink** | Pointer to **NullableBool** | Specifies whether to create a symlink for the migrated data from source to target. | [optional] [default to true]
**AutoOrphanDataCleanup** | Pointer to **NullableBool** | Specifies whether to remove the orphan data from the target if the symlink is removed from the source. | [optional] [default to true]
**TieringGoal** | Pointer to **NullableInt64** | Specifies the maximum amount of data that should be present on source after downtiering. | [optional] 
**FileAge** | Pointer to [**DowntieringFileAgePolicy**](DowntieringFileAgePolicy.md) |  | [optional] 
**Target** | Pointer to [**NullableDowntieringTarget**](DowntieringTarget.md) |  | [optional] 

## Methods

### NewDowntieringPolicyAllOf

`func NewDowntieringPolicyAllOf() *DowntieringPolicyAllOf`

NewDowntieringPolicyAllOf instantiates a new DowntieringPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDowntieringPolicyAllOfWithDefaults

`func NewDowntieringPolicyAllOfWithDefaults() *DowntieringPolicyAllOf`

NewDowntieringPolicyAllOfWithDefaults instantiates a new DowntieringPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQosPolicy

`func (o *DowntieringPolicyAllOf) GetQosPolicy() string`

GetQosPolicy returns the QosPolicy field if non-nil, zero value otherwise.

### GetQosPolicyOk

`func (o *DowntieringPolicyAllOf) GetQosPolicyOk() (*string, bool)`

GetQosPolicyOk returns a tuple with the QosPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosPolicy

`func (o *DowntieringPolicyAllOf) SetQosPolicy(v string)`

SetQosPolicy sets QosPolicy field to given value.

### HasQosPolicy

`func (o *DowntieringPolicyAllOf) HasQosPolicy() bool`

HasQosPolicy returns a boolean if a field has been set.

### SetQosPolicyNil

`func (o *DowntieringPolicyAllOf) SetQosPolicyNil(b bool)`

 SetQosPolicyNil sets the value for QosPolicy to be an explicit nil

### UnsetQosPolicy
`func (o *DowntieringPolicyAllOf) UnsetQosPolicy()`

UnsetQosPolicy ensures that no value is present for QosPolicy, not even an explicit nil
### GetIndexingPolicy

`func (o *DowntieringPolicyAllOf) GetIndexingPolicy() IndexingPolicy`

GetIndexingPolicy returns the IndexingPolicy field if non-nil, zero value otherwise.

### GetIndexingPolicyOk

`func (o *DowntieringPolicyAllOf) GetIndexingPolicyOk() (*IndexingPolicy, bool)`

GetIndexingPolicyOk returns a tuple with the IndexingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingPolicy

`func (o *DowntieringPolicyAllOf) SetIndexingPolicy(v IndexingPolicy)`

SetIndexingPolicy sets IndexingPolicy field to given value.

### HasIndexingPolicy

`func (o *DowntieringPolicyAllOf) HasIndexingPolicy() bool`

HasIndexingPolicy returns a boolean if a field has been set.

### GetRetention

`func (o *DowntieringPolicyAllOf) GetRetention() Retention`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *DowntieringPolicyAllOf) GetRetentionOk() (*Retention, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *DowntieringPolicyAllOf) SetRetention(v Retention)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *DowntieringPolicyAllOf) HasRetention() bool`

HasRetention returns a boolean if a field has been set.

### GetSkipBackSymlink

`func (o *DowntieringPolicyAllOf) GetSkipBackSymlink() bool`

GetSkipBackSymlink returns the SkipBackSymlink field if non-nil, zero value otherwise.

### GetSkipBackSymlinkOk

`func (o *DowntieringPolicyAllOf) GetSkipBackSymlinkOk() (*bool, bool)`

GetSkipBackSymlinkOk returns a tuple with the SkipBackSymlink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipBackSymlink

`func (o *DowntieringPolicyAllOf) SetSkipBackSymlink(v bool)`

SetSkipBackSymlink sets SkipBackSymlink field to given value.

### HasSkipBackSymlink

`func (o *DowntieringPolicyAllOf) HasSkipBackSymlink() bool`

HasSkipBackSymlink returns a boolean if a field has been set.

### SetSkipBackSymlinkNil

`func (o *DowntieringPolicyAllOf) SetSkipBackSymlinkNil(b bool)`

 SetSkipBackSymlinkNil sets the value for SkipBackSymlink to be an explicit nil

### UnsetSkipBackSymlink
`func (o *DowntieringPolicyAllOf) UnsetSkipBackSymlink()`

UnsetSkipBackSymlink ensures that no value is present for SkipBackSymlink, not even an explicit nil
### GetAutoOrphanDataCleanup

`func (o *DowntieringPolicyAllOf) GetAutoOrphanDataCleanup() bool`

GetAutoOrphanDataCleanup returns the AutoOrphanDataCleanup field if non-nil, zero value otherwise.

### GetAutoOrphanDataCleanupOk

`func (o *DowntieringPolicyAllOf) GetAutoOrphanDataCleanupOk() (*bool, bool)`

GetAutoOrphanDataCleanupOk returns a tuple with the AutoOrphanDataCleanup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoOrphanDataCleanup

`func (o *DowntieringPolicyAllOf) SetAutoOrphanDataCleanup(v bool)`

SetAutoOrphanDataCleanup sets AutoOrphanDataCleanup field to given value.

### HasAutoOrphanDataCleanup

`func (o *DowntieringPolicyAllOf) HasAutoOrphanDataCleanup() bool`

HasAutoOrphanDataCleanup returns a boolean if a field has been set.

### SetAutoOrphanDataCleanupNil

`func (o *DowntieringPolicyAllOf) SetAutoOrphanDataCleanupNil(b bool)`

 SetAutoOrphanDataCleanupNil sets the value for AutoOrphanDataCleanup to be an explicit nil

### UnsetAutoOrphanDataCleanup
`func (o *DowntieringPolicyAllOf) UnsetAutoOrphanDataCleanup()`

UnsetAutoOrphanDataCleanup ensures that no value is present for AutoOrphanDataCleanup, not even an explicit nil
### GetTieringGoal

`func (o *DowntieringPolicyAllOf) GetTieringGoal() int64`

GetTieringGoal returns the TieringGoal field if non-nil, zero value otherwise.

### GetTieringGoalOk

`func (o *DowntieringPolicyAllOf) GetTieringGoalOk() (*int64, bool)`

GetTieringGoalOk returns a tuple with the TieringGoal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTieringGoal

`func (o *DowntieringPolicyAllOf) SetTieringGoal(v int64)`

SetTieringGoal sets TieringGoal field to given value.

### HasTieringGoal

`func (o *DowntieringPolicyAllOf) HasTieringGoal() bool`

HasTieringGoal returns a boolean if a field has been set.

### SetTieringGoalNil

`func (o *DowntieringPolicyAllOf) SetTieringGoalNil(b bool)`

 SetTieringGoalNil sets the value for TieringGoal to be an explicit nil

### UnsetTieringGoal
`func (o *DowntieringPolicyAllOf) UnsetTieringGoal()`

UnsetTieringGoal ensures that no value is present for TieringGoal, not even an explicit nil
### GetFileAge

`func (o *DowntieringPolicyAllOf) GetFileAge() DowntieringFileAgePolicy`

GetFileAge returns the FileAge field if non-nil, zero value otherwise.

### GetFileAgeOk

`func (o *DowntieringPolicyAllOf) GetFileAgeOk() (*DowntieringFileAgePolicy, bool)`

GetFileAgeOk returns a tuple with the FileAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileAge

`func (o *DowntieringPolicyAllOf) SetFileAge(v DowntieringFileAgePolicy)`

SetFileAge sets FileAge field to given value.

### HasFileAge

`func (o *DowntieringPolicyAllOf) HasFileAge() bool`

HasFileAge returns a boolean if a field has been set.

### GetTarget

`func (o *DowntieringPolicyAllOf) GetTarget() DowntieringTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *DowntieringPolicyAllOf) GetTargetOk() (*DowntieringTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *DowntieringPolicyAllOf) SetTarget(v DowntieringTarget)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *DowntieringPolicyAllOf) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTargetNil

`func (o *DowntieringPolicyAllOf) SetTargetNil(b bool)`

 SetTargetNil sets the value for Target to be an explicit nil

### UnsetTarget
`func (o *DowntieringPolicyAllOf) UnsetTarget()`

UnsetTarget ensures that no value is present for Target, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


