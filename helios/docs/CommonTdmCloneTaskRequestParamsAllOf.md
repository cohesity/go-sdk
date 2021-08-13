# CommonTdmCloneTaskRequestParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SnapshotId** | **NullableString** | Specifies the snapshot ID, from which the clone is to be created. | 
**TargetHostId** | **NullableInt64** | Specifies the ID of the host, where the clone needs to be created. | 
**PointInTimeUsecs** | Pointer to **NullableInt64** | Specifies the timestamp (in usecs from epoch) for creating the clone at a point in time in the past. | [optional] 

## Methods

### NewCommonTdmCloneTaskRequestParamsAllOf

`func NewCommonTdmCloneTaskRequestParamsAllOf(snapshotId NullableString, targetHostId NullableInt64, ) *CommonTdmCloneTaskRequestParamsAllOf`

NewCommonTdmCloneTaskRequestParamsAllOf instantiates a new CommonTdmCloneTaskRequestParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonTdmCloneTaskRequestParamsAllOfWithDefaults

`func NewCommonTdmCloneTaskRequestParamsAllOfWithDefaults() *CommonTdmCloneTaskRequestParamsAllOf`

NewCommonTdmCloneTaskRequestParamsAllOfWithDefaults instantiates a new CommonTdmCloneTaskRequestParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshotId

`func (o *CommonTdmCloneTaskRequestParamsAllOf) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *CommonTdmCloneTaskRequestParamsAllOf) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *CommonTdmCloneTaskRequestParamsAllOf) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.


### SetSnapshotIdNil

`func (o *CommonTdmCloneTaskRequestParamsAllOf) SetSnapshotIdNil(b bool)`

 SetSnapshotIdNil sets the value for SnapshotId to be an explicit nil

### UnsetSnapshotId
`func (o *CommonTdmCloneTaskRequestParamsAllOf) UnsetSnapshotId()`

UnsetSnapshotId ensures that no value is present for SnapshotId, not even an explicit nil
### GetTargetHostId

`func (o *CommonTdmCloneTaskRequestParamsAllOf) GetTargetHostId() int64`

GetTargetHostId returns the TargetHostId field if non-nil, zero value otherwise.

### GetTargetHostIdOk

`func (o *CommonTdmCloneTaskRequestParamsAllOf) GetTargetHostIdOk() (*int64, bool)`

GetTargetHostIdOk returns a tuple with the TargetHostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetHostId

`func (o *CommonTdmCloneTaskRequestParamsAllOf) SetTargetHostId(v int64)`

SetTargetHostId sets TargetHostId field to given value.


### SetTargetHostIdNil

`func (o *CommonTdmCloneTaskRequestParamsAllOf) SetTargetHostIdNil(b bool)`

 SetTargetHostIdNil sets the value for TargetHostId to be an explicit nil

### UnsetTargetHostId
`func (o *CommonTdmCloneTaskRequestParamsAllOf) UnsetTargetHostId()`

UnsetTargetHostId ensures that no value is present for TargetHostId, not even an explicit nil
### GetPointInTimeUsecs

`func (o *CommonTdmCloneTaskRequestParamsAllOf) GetPointInTimeUsecs() int64`

GetPointInTimeUsecs returns the PointInTimeUsecs field if non-nil, zero value otherwise.

### GetPointInTimeUsecsOk

`func (o *CommonTdmCloneTaskRequestParamsAllOf) GetPointInTimeUsecsOk() (*int64, bool)`

GetPointInTimeUsecsOk returns a tuple with the PointInTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointInTimeUsecs

`func (o *CommonTdmCloneTaskRequestParamsAllOf) SetPointInTimeUsecs(v int64)`

SetPointInTimeUsecs sets PointInTimeUsecs field to given value.

### HasPointInTimeUsecs

`func (o *CommonTdmCloneTaskRequestParamsAllOf) HasPointInTimeUsecs() bool`

HasPointInTimeUsecs returns a boolean if a field has been set.

### SetPointInTimeUsecsNil

`func (o *CommonTdmCloneTaskRequestParamsAllOf) SetPointInTimeUsecsNil(b bool)`

 SetPointInTimeUsecsNil sets the value for PointInTimeUsecs to be an explicit nil

### UnsetPointInTimeUsecs
`func (o *CommonTdmCloneTaskRequestParamsAllOf) UnsetPointInTimeUsecs()`

UnsetPointInTimeUsecs ensures that no value is present for PointInTimeUsecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


