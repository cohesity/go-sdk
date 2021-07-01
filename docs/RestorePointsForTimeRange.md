# RestorePointsForTimeRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullSnapshotInfo** | Pointer to [**[]FullSnapshotInfo**](FullSnapshotInfo.md) | Specifies the info related to the recovery object. | [optional] 
**TimeRanges** | Pointer to [**[]TimeRangeSettings**](TimeRangeSettings.md) | Specifies the time ranges of the restore object between full snapshots. | [optional] 

## Methods

### NewRestorePointsForTimeRange

`func NewRestorePointsForTimeRange() *RestorePointsForTimeRange`

NewRestorePointsForTimeRange instantiates a new RestorePointsForTimeRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestorePointsForTimeRangeWithDefaults

`func NewRestorePointsForTimeRangeWithDefaults() *RestorePointsForTimeRange`

NewRestorePointsForTimeRangeWithDefaults instantiates a new RestorePointsForTimeRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullSnapshotInfo

`func (o *RestorePointsForTimeRange) GetFullSnapshotInfo() []FullSnapshotInfo`

GetFullSnapshotInfo returns the FullSnapshotInfo field if non-nil, zero value otherwise.

### GetFullSnapshotInfoOk

`func (o *RestorePointsForTimeRange) GetFullSnapshotInfoOk() (*[]FullSnapshotInfo, bool)`

GetFullSnapshotInfoOk returns a tuple with the FullSnapshotInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullSnapshotInfo

`func (o *RestorePointsForTimeRange) SetFullSnapshotInfo(v []FullSnapshotInfo)`

SetFullSnapshotInfo sets FullSnapshotInfo field to given value.

### HasFullSnapshotInfo

`func (o *RestorePointsForTimeRange) HasFullSnapshotInfo() bool`

HasFullSnapshotInfo returns a boolean if a field has been set.

### SetFullSnapshotInfoNil

`func (o *RestorePointsForTimeRange) SetFullSnapshotInfoNil(b bool)`

 SetFullSnapshotInfoNil sets the value for FullSnapshotInfo to be an explicit nil

### UnsetFullSnapshotInfo
`func (o *RestorePointsForTimeRange) UnsetFullSnapshotInfo()`

UnsetFullSnapshotInfo ensures that no value is present for FullSnapshotInfo, not even an explicit nil
### GetTimeRanges

`func (o *RestorePointsForTimeRange) GetTimeRanges() []TimeRangeSettings`

GetTimeRanges returns the TimeRanges field if non-nil, zero value otherwise.

### GetTimeRangesOk

`func (o *RestorePointsForTimeRange) GetTimeRangesOk() (*[]TimeRangeSettings, bool)`

GetTimeRangesOk returns a tuple with the TimeRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRanges

`func (o *RestorePointsForTimeRange) SetTimeRanges(v []TimeRangeSettings)`

SetTimeRanges sets TimeRanges field to given value.

### HasTimeRanges

`func (o *RestorePointsForTimeRange) HasTimeRanges() bool`

HasTimeRanges returns a boolean if a field has been set.

### SetTimeRangesNil

`func (o *RestorePointsForTimeRange) SetTimeRangesNil(b bool)`

 SetTimeRangesNil sets the value for TimeRanges to be an explicit nil

### UnsetTimeRanges
`func (o *RestorePointsForTimeRange) UnsetTimeRanges()`

UnsetTimeRanges ensures that no value is present for TimeRanges, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


