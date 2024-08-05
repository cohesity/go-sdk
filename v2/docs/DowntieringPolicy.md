# DowntieringPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableAuditLogging** | Pointer to **NullableBool** | Specifies whether to audit log the file tiering activity. | [optional] [default to false]
**FilePath** | Pointer to [**FileFilteringPolicy**](FileFilteringPolicy.md) |  | [optional] 
**FileSize** | Pointer to [**FileSizePolicy**](FileSizePolicy.md) |  | [optional] 
**AutoOrphanDataCleanup** | Pointer to **NullableBool** | Specifies whether to remove the orphan data from the target if the symlink is removed from the source. | [optional] [default to true]
**FileAge** | Pointer to [**DowntieringFileAgePolicy**](DowntieringFileAgePolicy.md) |  | [optional] 
**IndexingPolicy** | Pointer to [**IndexingPolicy**](IndexingPolicy.md) |  | [optional] 
**QosPolicy** | Pointer to **NullableString** | Specifies whether the data tiering task will be written to HDD or SSD. | [optional] 
**Retention** | Pointer to [**Retention**](Retention.md) |  | [optional] 
**SkipBackSymlink** | Pointer to **NullableBool** | Specifies whether to create a symlink for the migrated data from source to target. | [optional] [default to true]
**TagsInfo** | Pointer to [**[]DataTieringTagObject**](DataTieringTagObject.md) | Array of Tag objects used to represent different file based policies | [optional] 
**Target** | Pointer to [**NullableDowntieringTarget**](DowntieringTarget.md) |  | [optional] 
**TieringGoal** | Pointer to **NullableInt64** | Specifies the maximum amount of data that should be present on source after downtiering. | [optional] 

## Methods

### NewDowntieringPolicy

`func NewDowntieringPolicy() *DowntieringPolicy`

NewDowntieringPolicy instantiates a new DowntieringPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDowntieringPolicyWithDefaults

`func NewDowntieringPolicyWithDefaults() *DowntieringPolicy`

NewDowntieringPolicyWithDefaults instantiates a new DowntieringPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableAuditLogging

`func (o *DowntieringPolicy) GetEnableAuditLogging() bool`

GetEnableAuditLogging returns the EnableAuditLogging field if non-nil, zero value otherwise.

### GetEnableAuditLoggingOk

`func (o *DowntieringPolicy) GetEnableAuditLoggingOk() (*bool, bool)`

GetEnableAuditLoggingOk returns a tuple with the EnableAuditLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAuditLogging

`func (o *DowntieringPolicy) SetEnableAuditLogging(v bool)`

SetEnableAuditLogging sets EnableAuditLogging field to given value.

### HasEnableAuditLogging

`func (o *DowntieringPolicy) HasEnableAuditLogging() bool`

HasEnableAuditLogging returns a boolean if a field has been set.

### SetEnableAuditLoggingNil

`func (o *DowntieringPolicy) SetEnableAuditLoggingNil(b bool)`

 SetEnableAuditLoggingNil sets the value for EnableAuditLogging to be an explicit nil

### UnsetEnableAuditLogging
`func (o *DowntieringPolicy) UnsetEnableAuditLogging()`

UnsetEnableAuditLogging ensures that no value is present for EnableAuditLogging, not even an explicit nil
### GetFilePath

`func (o *DowntieringPolicy) GetFilePath() FileFilteringPolicy`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *DowntieringPolicy) GetFilePathOk() (*FileFilteringPolicy, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *DowntieringPolicy) SetFilePath(v FileFilteringPolicy)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *DowntieringPolicy) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### GetFileSize

`func (o *DowntieringPolicy) GetFileSize() FileSizePolicy`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *DowntieringPolicy) GetFileSizeOk() (*FileSizePolicy, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *DowntieringPolicy) SetFileSize(v FileSizePolicy)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *DowntieringPolicy) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetAutoOrphanDataCleanup

`func (o *DowntieringPolicy) GetAutoOrphanDataCleanup() bool`

GetAutoOrphanDataCleanup returns the AutoOrphanDataCleanup field if non-nil, zero value otherwise.

### GetAutoOrphanDataCleanupOk

`func (o *DowntieringPolicy) GetAutoOrphanDataCleanupOk() (*bool, bool)`

GetAutoOrphanDataCleanupOk returns a tuple with the AutoOrphanDataCleanup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoOrphanDataCleanup

`func (o *DowntieringPolicy) SetAutoOrphanDataCleanup(v bool)`

SetAutoOrphanDataCleanup sets AutoOrphanDataCleanup field to given value.

### HasAutoOrphanDataCleanup

`func (o *DowntieringPolicy) HasAutoOrphanDataCleanup() bool`

HasAutoOrphanDataCleanup returns a boolean if a field has been set.

### SetAutoOrphanDataCleanupNil

`func (o *DowntieringPolicy) SetAutoOrphanDataCleanupNil(b bool)`

 SetAutoOrphanDataCleanupNil sets the value for AutoOrphanDataCleanup to be an explicit nil

### UnsetAutoOrphanDataCleanup
`func (o *DowntieringPolicy) UnsetAutoOrphanDataCleanup()`

UnsetAutoOrphanDataCleanup ensures that no value is present for AutoOrphanDataCleanup, not even an explicit nil
### GetFileAge

`func (o *DowntieringPolicy) GetFileAge() DowntieringFileAgePolicy`

GetFileAge returns the FileAge field if non-nil, zero value otherwise.

### GetFileAgeOk

`func (o *DowntieringPolicy) GetFileAgeOk() (*DowntieringFileAgePolicy, bool)`

GetFileAgeOk returns a tuple with the FileAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileAge

`func (o *DowntieringPolicy) SetFileAge(v DowntieringFileAgePolicy)`

SetFileAge sets FileAge field to given value.

### HasFileAge

`func (o *DowntieringPolicy) HasFileAge() bool`

HasFileAge returns a boolean if a field has been set.

### GetIndexingPolicy

`func (o *DowntieringPolicy) GetIndexingPolicy() IndexingPolicy`

GetIndexingPolicy returns the IndexingPolicy field if non-nil, zero value otherwise.

### GetIndexingPolicyOk

`func (o *DowntieringPolicy) GetIndexingPolicyOk() (*IndexingPolicy, bool)`

GetIndexingPolicyOk returns a tuple with the IndexingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingPolicy

`func (o *DowntieringPolicy) SetIndexingPolicy(v IndexingPolicy)`

SetIndexingPolicy sets IndexingPolicy field to given value.

### HasIndexingPolicy

`func (o *DowntieringPolicy) HasIndexingPolicy() bool`

HasIndexingPolicy returns a boolean if a field has been set.

### GetQosPolicy

`func (o *DowntieringPolicy) GetQosPolicy() string`

GetQosPolicy returns the QosPolicy field if non-nil, zero value otherwise.

### GetQosPolicyOk

`func (o *DowntieringPolicy) GetQosPolicyOk() (*string, bool)`

GetQosPolicyOk returns a tuple with the QosPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosPolicy

`func (o *DowntieringPolicy) SetQosPolicy(v string)`

SetQosPolicy sets QosPolicy field to given value.

### HasQosPolicy

`func (o *DowntieringPolicy) HasQosPolicy() bool`

HasQosPolicy returns a boolean if a field has been set.

### SetQosPolicyNil

`func (o *DowntieringPolicy) SetQosPolicyNil(b bool)`

 SetQosPolicyNil sets the value for QosPolicy to be an explicit nil

### UnsetQosPolicy
`func (o *DowntieringPolicy) UnsetQosPolicy()`

UnsetQosPolicy ensures that no value is present for QosPolicy, not even an explicit nil
### GetRetention

`func (o *DowntieringPolicy) GetRetention() Retention`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *DowntieringPolicy) GetRetentionOk() (*Retention, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *DowntieringPolicy) SetRetention(v Retention)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *DowntieringPolicy) HasRetention() bool`

HasRetention returns a boolean if a field has been set.

### GetSkipBackSymlink

`func (o *DowntieringPolicy) GetSkipBackSymlink() bool`

GetSkipBackSymlink returns the SkipBackSymlink field if non-nil, zero value otherwise.

### GetSkipBackSymlinkOk

`func (o *DowntieringPolicy) GetSkipBackSymlinkOk() (*bool, bool)`

GetSkipBackSymlinkOk returns a tuple with the SkipBackSymlink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipBackSymlink

`func (o *DowntieringPolicy) SetSkipBackSymlink(v bool)`

SetSkipBackSymlink sets SkipBackSymlink field to given value.

### HasSkipBackSymlink

`func (o *DowntieringPolicy) HasSkipBackSymlink() bool`

HasSkipBackSymlink returns a boolean if a field has been set.

### SetSkipBackSymlinkNil

`func (o *DowntieringPolicy) SetSkipBackSymlinkNil(b bool)`

 SetSkipBackSymlinkNil sets the value for SkipBackSymlink to be an explicit nil

### UnsetSkipBackSymlink
`func (o *DowntieringPolicy) UnsetSkipBackSymlink()`

UnsetSkipBackSymlink ensures that no value is present for SkipBackSymlink, not even an explicit nil
### GetTagsInfo

`func (o *DowntieringPolicy) GetTagsInfo() []DataTieringTagObject`

GetTagsInfo returns the TagsInfo field if non-nil, zero value otherwise.

### GetTagsInfoOk

`func (o *DowntieringPolicy) GetTagsInfoOk() (*[]DataTieringTagObject, bool)`

GetTagsInfoOk returns a tuple with the TagsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagsInfo

`func (o *DowntieringPolicy) SetTagsInfo(v []DataTieringTagObject)`

SetTagsInfo sets TagsInfo field to given value.

### HasTagsInfo

`func (o *DowntieringPolicy) HasTagsInfo() bool`

HasTagsInfo returns a boolean if a field has been set.

### SetTagsInfoNil

`func (o *DowntieringPolicy) SetTagsInfoNil(b bool)`

 SetTagsInfoNil sets the value for TagsInfo to be an explicit nil

### UnsetTagsInfo
`func (o *DowntieringPolicy) UnsetTagsInfo()`

UnsetTagsInfo ensures that no value is present for TagsInfo, not even an explicit nil
### GetTarget

`func (o *DowntieringPolicy) GetTarget() DowntieringTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *DowntieringPolicy) GetTargetOk() (*DowntieringTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *DowntieringPolicy) SetTarget(v DowntieringTarget)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *DowntieringPolicy) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTargetNil

`func (o *DowntieringPolicy) SetTargetNil(b bool)`

 SetTargetNil sets the value for Target to be an explicit nil

### UnsetTarget
`func (o *DowntieringPolicy) UnsetTarget()`

UnsetTarget ensures that no value is present for Target, not even an explicit nil
### GetTieringGoal

`func (o *DowntieringPolicy) GetTieringGoal() int64`

GetTieringGoal returns the TieringGoal field if non-nil, zero value otherwise.

### GetTieringGoalOk

`func (o *DowntieringPolicy) GetTieringGoalOk() (*int64, bool)`

GetTieringGoalOk returns a tuple with the TieringGoal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTieringGoal

`func (o *DowntieringPolicy) SetTieringGoal(v int64)`

SetTieringGoal sets TieringGoal field to given value.

### HasTieringGoal

`func (o *DowntieringPolicy) HasTieringGoal() bool`

HasTieringGoal returns a boolean if a field has been set.

### SetTieringGoalNil

`func (o *DowntieringPolicy) SetTieringGoalNil(b bool)`

 SetTieringGoalNil sets the value for TieringGoal to be an explicit nil

### UnsetTieringGoal
`func (o *DowntieringPolicy) UnsetTieringGoal()`

UnsetTieringGoal ensures that no value is present for TieringGoal, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


