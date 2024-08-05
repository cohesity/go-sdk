# ObjectSiteParamOwnerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchivalTargetInfo** | Pointer to [**NullableCommonRecoverObjectSnapshotParamsArchivalTargetInfo**](CommonRecoverObjectSnapshotParamsArchivalTargetInfo.md) |  | [optional] 
**BytesRestored** | Pointer to **NullableInt64** | Specify the total bytes restored. | [optional] [readonly] 
**EndTimeUsecs** | Pointer to **NullableInt64** | Specifies the end time of the Recovery in Unix timestamp epoch in microseconds. This field will be populated only after Recovery is finished. | [optional] [readonly] 
**Messages** | Pointer to **[]string** | Specify error messages about the object. | [optional] [readonly] 
**ObjectInfo** | Pointer to [**NullableCommonRecoverObjectSnapshotParamsObjectInfo**](CommonRecoverObjectSnapshotParamsObjectInfo.md) |  | [optional] 
**PointInTimeUsecs** | Pointer to **NullableInt64** | Specifies the timestamp (in microseconds. from epoch) for recovering to a point-in-time in the past. | [optional] 
**ProgressTaskId** | Pointer to **NullableString** | Progress monitor task id for Recovery of VM. | [optional] [readonly] 
**ProtectionGroupId** | Pointer to **NullableString** | Specifies the protection group id of the object snapshot. | [optional] 
**ProtectionGroupName** | Pointer to **NullableString** | Specifies the protection group name of the object snapshot. | [optional] 
**RecoverFromStandby** | Pointer to **NullableBool** | Specifies that user wants to perform standby restore if it is enabled for this object. | [optional] 
**SnapshotCreationTimeUsecs** | Pointer to **NullableInt64** | Specifies the time when the snapshot is created in Unix timestamp epoch in microseconds. | [optional] [readonly] 
**SnapshotId** | **string** | Specifies the snapshot id. | 
**SnapshotTargetType** | Pointer to **NullableString** | Specifies the snapshot target type. | [optional] [readonly] 
**StartTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time of the Recovery in Unix timestamp epoch in microseconds. | [optional] [readonly] 
**Status** | Pointer to **NullableString** | Status of the Recovery. &#39;Running&#39; indicates that the Recovery is still running. &#39;Canceled&#39; indicates that the Recovery has been cancelled. &#39;Canceling&#39; indicates that the Recovery is in the process of being cancelled. &#39;Failed&#39; indicates that the Recovery has failed. &#39;Succeeded&#39; indicates that the Recovery has finished successfully. &#39;SucceededWithWarning&#39; indicates that the Recovery finished successfully, but there were some warning messages. &#39;Skipped&#39; indicates that the Recovery task was skipped. | [optional] [readonly] 
**StorageDomainId** | Pointer to **NullableInt64** | Specifies the ID of the Storage Domain where this snapshot is stored. | [optional] [readonly] 

## Methods

### NewObjectSiteParamOwnerInfo

`func NewObjectSiteParamOwnerInfo(snapshotId string, ) *ObjectSiteParamOwnerInfo`

NewObjectSiteParamOwnerInfo instantiates a new ObjectSiteParamOwnerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectSiteParamOwnerInfoWithDefaults

`func NewObjectSiteParamOwnerInfoWithDefaults() *ObjectSiteParamOwnerInfo`

NewObjectSiteParamOwnerInfoWithDefaults instantiates a new ObjectSiteParamOwnerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchivalTargetInfo

`func (o *ObjectSiteParamOwnerInfo) GetArchivalTargetInfo() CommonRecoverObjectSnapshotParamsArchivalTargetInfo`

GetArchivalTargetInfo returns the ArchivalTargetInfo field if non-nil, zero value otherwise.

### GetArchivalTargetInfoOk

`func (o *ObjectSiteParamOwnerInfo) GetArchivalTargetInfoOk() (*CommonRecoverObjectSnapshotParamsArchivalTargetInfo, bool)`

GetArchivalTargetInfoOk returns a tuple with the ArchivalTargetInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivalTargetInfo

`func (o *ObjectSiteParamOwnerInfo) SetArchivalTargetInfo(v CommonRecoverObjectSnapshotParamsArchivalTargetInfo)`

SetArchivalTargetInfo sets ArchivalTargetInfo field to given value.

### HasArchivalTargetInfo

`func (o *ObjectSiteParamOwnerInfo) HasArchivalTargetInfo() bool`

HasArchivalTargetInfo returns a boolean if a field has been set.

### SetArchivalTargetInfoNil

`func (o *ObjectSiteParamOwnerInfo) SetArchivalTargetInfoNil(b bool)`

 SetArchivalTargetInfoNil sets the value for ArchivalTargetInfo to be an explicit nil

### UnsetArchivalTargetInfo
`func (o *ObjectSiteParamOwnerInfo) UnsetArchivalTargetInfo()`

UnsetArchivalTargetInfo ensures that no value is present for ArchivalTargetInfo, not even an explicit nil
### GetBytesRestored

`func (o *ObjectSiteParamOwnerInfo) GetBytesRestored() int64`

GetBytesRestored returns the BytesRestored field if non-nil, zero value otherwise.

### GetBytesRestoredOk

`func (o *ObjectSiteParamOwnerInfo) GetBytesRestoredOk() (*int64, bool)`

GetBytesRestoredOk returns a tuple with the BytesRestored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesRestored

`func (o *ObjectSiteParamOwnerInfo) SetBytesRestored(v int64)`

SetBytesRestored sets BytesRestored field to given value.

### HasBytesRestored

`func (o *ObjectSiteParamOwnerInfo) HasBytesRestored() bool`

HasBytesRestored returns a boolean if a field has been set.

### SetBytesRestoredNil

`func (o *ObjectSiteParamOwnerInfo) SetBytesRestoredNil(b bool)`

 SetBytesRestoredNil sets the value for BytesRestored to be an explicit nil

### UnsetBytesRestored
`func (o *ObjectSiteParamOwnerInfo) UnsetBytesRestored()`

UnsetBytesRestored ensures that no value is present for BytesRestored, not even an explicit nil
### GetEndTimeUsecs

`func (o *ObjectSiteParamOwnerInfo) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *ObjectSiteParamOwnerInfo) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *ObjectSiteParamOwnerInfo) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *ObjectSiteParamOwnerInfo) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *ObjectSiteParamOwnerInfo) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *ObjectSiteParamOwnerInfo) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetMessages

`func (o *ObjectSiteParamOwnerInfo) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ObjectSiteParamOwnerInfo) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ObjectSiteParamOwnerInfo) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *ObjectSiteParamOwnerInfo) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### SetMessagesNil

`func (o *ObjectSiteParamOwnerInfo) SetMessagesNil(b bool)`

 SetMessagesNil sets the value for Messages to be an explicit nil

### UnsetMessages
`func (o *ObjectSiteParamOwnerInfo) UnsetMessages()`

UnsetMessages ensures that no value is present for Messages, not even an explicit nil
### GetObjectInfo

`func (o *ObjectSiteParamOwnerInfo) GetObjectInfo() CommonRecoverObjectSnapshotParamsObjectInfo`

GetObjectInfo returns the ObjectInfo field if non-nil, zero value otherwise.

### GetObjectInfoOk

`func (o *ObjectSiteParamOwnerInfo) GetObjectInfoOk() (*CommonRecoverObjectSnapshotParamsObjectInfo, bool)`

GetObjectInfoOk returns a tuple with the ObjectInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectInfo

`func (o *ObjectSiteParamOwnerInfo) SetObjectInfo(v CommonRecoverObjectSnapshotParamsObjectInfo)`

SetObjectInfo sets ObjectInfo field to given value.

### HasObjectInfo

`func (o *ObjectSiteParamOwnerInfo) HasObjectInfo() bool`

HasObjectInfo returns a boolean if a field has been set.

### SetObjectInfoNil

`func (o *ObjectSiteParamOwnerInfo) SetObjectInfoNil(b bool)`

 SetObjectInfoNil sets the value for ObjectInfo to be an explicit nil

### UnsetObjectInfo
`func (o *ObjectSiteParamOwnerInfo) UnsetObjectInfo()`

UnsetObjectInfo ensures that no value is present for ObjectInfo, not even an explicit nil
### GetPointInTimeUsecs

`func (o *ObjectSiteParamOwnerInfo) GetPointInTimeUsecs() int64`

GetPointInTimeUsecs returns the PointInTimeUsecs field if non-nil, zero value otherwise.

### GetPointInTimeUsecsOk

`func (o *ObjectSiteParamOwnerInfo) GetPointInTimeUsecsOk() (*int64, bool)`

GetPointInTimeUsecsOk returns a tuple with the PointInTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointInTimeUsecs

`func (o *ObjectSiteParamOwnerInfo) SetPointInTimeUsecs(v int64)`

SetPointInTimeUsecs sets PointInTimeUsecs field to given value.

### HasPointInTimeUsecs

`func (o *ObjectSiteParamOwnerInfo) HasPointInTimeUsecs() bool`

HasPointInTimeUsecs returns a boolean if a field has been set.

### SetPointInTimeUsecsNil

`func (o *ObjectSiteParamOwnerInfo) SetPointInTimeUsecsNil(b bool)`

 SetPointInTimeUsecsNil sets the value for PointInTimeUsecs to be an explicit nil

### UnsetPointInTimeUsecs
`func (o *ObjectSiteParamOwnerInfo) UnsetPointInTimeUsecs()`

UnsetPointInTimeUsecs ensures that no value is present for PointInTimeUsecs, not even an explicit nil
### GetProgressTaskId

`func (o *ObjectSiteParamOwnerInfo) GetProgressTaskId() string`

GetProgressTaskId returns the ProgressTaskId field if non-nil, zero value otherwise.

### GetProgressTaskIdOk

`func (o *ObjectSiteParamOwnerInfo) GetProgressTaskIdOk() (*string, bool)`

GetProgressTaskIdOk returns a tuple with the ProgressTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressTaskId

`func (o *ObjectSiteParamOwnerInfo) SetProgressTaskId(v string)`

SetProgressTaskId sets ProgressTaskId field to given value.

### HasProgressTaskId

`func (o *ObjectSiteParamOwnerInfo) HasProgressTaskId() bool`

HasProgressTaskId returns a boolean if a field has been set.

### SetProgressTaskIdNil

`func (o *ObjectSiteParamOwnerInfo) SetProgressTaskIdNil(b bool)`

 SetProgressTaskIdNil sets the value for ProgressTaskId to be an explicit nil

### UnsetProgressTaskId
`func (o *ObjectSiteParamOwnerInfo) UnsetProgressTaskId()`

UnsetProgressTaskId ensures that no value is present for ProgressTaskId, not even an explicit nil
### GetProtectionGroupId

`func (o *ObjectSiteParamOwnerInfo) GetProtectionGroupId() string`

GetProtectionGroupId returns the ProtectionGroupId field if non-nil, zero value otherwise.

### GetProtectionGroupIdOk

`func (o *ObjectSiteParamOwnerInfo) GetProtectionGroupIdOk() (*string, bool)`

GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupId

`func (o *ObjectSiteParamOwnerInfo) SetProtectionGroupId(v string)`

SetProtectionGroupId sets ProtectionGroupId field to given value.

### HasProtectionGroupId

`func (o *ObjectSiteParamOwnerInfo) HasProtectionGroupId() bool`

HasProtectionGroupId returns a boolean if a field has been set.

### SetProtectionGroupIdNil

`func (o *ObjectSiteParamOwnerInfo) SetProtectionGroupIdNil(b bool)`

 SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil

### UnsetProtectionGroupId
`func (o *ObjectSiteParamOwnerInfo) UnsetProtectionGroupId()`

UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil
### GetProtectionGroupName

`func (o *ObjectSiteParamOwnerInfo) GetProtectionGroupName() string`

GetProtectionGroupName returns the ProtectionGroupName field if non-nil, zero value otherwise.

### GetProtectionGroupNameOk

`func (o *ObjectSiteParamOwnerInfo) GetProtectionGroupNameOk() (*string, bool)`

GetProtectionGroupNameOk returns a tuple with the ProtectionGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupName

`func (o *ObjectSiteParamOwnerInfo) SetProtectionGroupName(v string)`

SetProtectionGroupName sets ProtectionGroupName field to given value.

### HasProtectionGroupName

`func (o *ObjectSiteParamOwnerInfo) HasProtectionGroupName() bool`

HasProtectionGroupName returns a boolean if a field has been set.

### SetProtectionGroupNameNil

`func (o *ObjectSiteParamOwnerInfo) SetProtectionGroupNameNil(b bool)`

 SetProtectionGroupNameNil sets the value for ProtectionGroupName to be an explicit nil

### UnsetProtectionGroupName
`func (o *ObjectSiteParamOwnerInfo) UnsetProtectionGroupName()`

UnsetProtectionGroupName ensures that no value is present for ProtectionGroupName, not even an explicit nil
### GetRecoverFromStandby

`func (o *ObjectSiteParamOwnerInfo) GetRecoverFromStandby() bool`

GetRecoverFromStandby returns the RecoverFromStandby field if non-nil, zero value otherwise.

### GetRecoverFromStandbyOk

`func (o *ObjectSiteParamOwnerInfo) GetRecoverFromStandbyOk() (*bool, bool)`

GetRecoverFromStandbyOk returns a tuple with the RecoverFromStandby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverFromStandby

`func (o *ObjectSiteParamOwnerInfo) SetRecoverFromStandby(v bool)`

SetRecoverFromStandby sets RecoverFromStandby field to given value.

### HasRecoverFromStandby

`func (o *ObjectSiteParamOwnerInfo) HasRecoverFromStandby() bool`

HasRecoverFromStandby returns a boolean if a field has been set.

### SetRecoverFromStandbyNil

`func (o *ObjectSiteParamOwnerInfo) SetRecoverFromStandbyNil(b bool)`

 SetRecoverFromStandbyNil sets the value for RecoverFromStandby to be an explicit nil

### UnsetRecoverFromStandby
`func (o *ObjectSiteParamOwnerInfo) UnsetRecoverFromStandby()`

UnsetRecoverFromStandby ensures that no value is present for RecoverFromStandby, not even an explicit nil
### GetSnapshotCreationTimeUsecs

`func (o *ObjectSiteParamOwnerInfo) GetSnapshotCreationTimeUsecs() int64`

GetSnapshotCreationTimeUsecs returns the SnapshotCreationTimeUsecs field if non-nil, zero value otherwise.

### GetSnapshotCreationTimeUsecsOk

`func (o *ObjectSiteParamOwnerInfo) GetSnapshotCreationTimeUsecsOk() (*int64, bool)`

GetSnapshotCreationTimeUsecsOk returns a tuple with the SnapshotCreationTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotCreationTimeUsecs

`func (o *ObjectSiteParamOwnerInfo) SetSnapshotCreationTimeUsecs(v int64)`

SetSnapshotCreationTimeUsecs sets SnapshotCreationTimeUsecs field to given value.

### HasSnapshotCreationTimeUsecs

`func (o *ObjectSiteParamOwnerInfo) HasSnapshotCreationTimeUsecs() bool`

HasSnapshotCreationTimeUsecs returns a boolean if a field has been set.

### SetSnapshotCreationTimeUsecsNil

`func (o *ObjectSiteParamOwnerInfo) SetSnapshotCreationTimeUsecsNil(b bool)`

 SetSnapshotCreationTimeUsecsNil sets the value for SnapshotCreationTimeUsecs to be an explicit nil

### UnsetSnapshotCreationTimeUsecs
`func (o *ObjectSiteParamOwnerInfo) UnsetSnapshotCreationTimeUsecs()`

UnsetSnapshotCreationTimeUsecs ensures that no value is present for SnapshotCreationTimeUsecs, not even an explicit nil
### GetSnapshotId

`func (o *ObjectSiteParamOwnerInfo) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *ObjectSiteParamOwnerInfo) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *ObjectSiteParamOwnerInfo) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.


### GetSnapshotTargetType

`func (o *ObjectSiteParamOwnerInfo) GetSnapshotTargetType() string`

GetSnapshotTargetType returns the SnapshotTargetType field if non-nil, zero value otherwise.

### GetSnapshotTargetTypeOk

`func (o *ObjectSiteParamOwnerInfo) GetSnapshotTargetTypeOk() (*string, bool)`

GetSnapshotTargetTypeOk returns a tuple with the SnapshotTargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotTargetType

`func (o *ObjectSiteParamOwnerInfo) SetSnapshotTargetType(v string)`

SetSnapshotTargetType sets SnapshotTargetType field to given value.

### HasSnapshotTargetType

`func (o *ObjectSiteParamOwnerInfo) HasSnapshotTargetType() bool`

HasSnapshotTargetType returns a boolean if a field has been set.

### SetSnapshotTargetTypeNil

`func (o *ObjectSiteParamOwnerInfo) SetSnapshotTargetTypeNil(b bool)`

 SetSnapshotTargetTypeNil sets the value for SnapshotTargetType to be an explicit nil

### UnsetSnapshotTargetType
`func (o *ObjectSiteParamOwnerInfo) UnsetSnapshotTargetType()`

UnsetSnapshotTargetType ensures that no value is present for SnapshotTargetType, not even an explicit nil
### GetStartTimeUsecs

`func (o *ObjectSiteParamOwnerInfo) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *ObjectSiteParamOwnerInfo) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *ObjectSiteParamOwnerInfo) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *ObjectSiteParamOwnerInfo) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *ObjectSiteParamOwnerInfo) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *ObjectSiteParamOwnerInfo) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetStatus

`func (o *ObjectSiteParamOwnerInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ObjectSiteParamOwnerInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ObjectSiteParamOwnerInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ObjectSiteParamOwnerInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ObjectSiteParamOwnerInfo) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ObjectSiteParamOwnerInfo) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetStorageDomainId

`func (o *ObjectSiteParamOwnerInfo) GetStorageDomainId() int64`

GetStorageDomainId returns the StorageDomainId field if non-nil, zero value otherwise.

### GetStorageDomainIdOk

`func (o *ObjectSiteParamOwnerInfo) GetStorageDomainIdOk() (*int64, bool)`

GetStorageDomainIdOk returns a tuple with the StorageDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainId

`func (o *ObjectSiteParamOwnerInfo) SetStorageDomainId(v int64)`

SetStorageDomainId sets StorageDomainId field to given value.

### HasStorageDomainId

`func (o *ObjectSiteParamOwnerInfo) HasStorageDomainId() bool`

HasStorageDomainId returns a boolean if a field has been set.

### SetStorageDomainIdNil

`func (o *ObjectSiteParamOwnerInfo) SetStorageDomainIdNil(b bool)`

 SetStorageDomainIdNil sets the value for StorageDomainId to be an explicit nil

### UnsetStorageDomainId
`func (o *ObjectSiteParamOwnerInfo) UnsetStorageDomainId()`

UnsetStorageDomainId ensures that no value is present for StorageDomainId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


