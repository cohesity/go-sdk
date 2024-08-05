# SnapshotHandle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobUid** | Pointer to **NullableInt64** | Specifies a distinct value that&#39;s unique to a source. | [optional] 
**RunStartTimeUsecs** | Pointer to **NullableInt64** | Run start time of the Magneto job run which has taken this snapshot. | [optional] 

## Methods

### NewSnapshotHandle

`func NewSnapshotHandle() *SnapshotHandle`

NewSnapshotHandle instantiates a new SnapshotHandle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotHandleWithDefaults

`func NewSnapshotHandleWithDefaults() *SnapshotHandle`

NewSnapshotHandleWithDefaults instantiates a new SnapshotHandle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobUid

`func (o *SnapshotHandle) GetJobUid() int64`

GetJobUid returns the JobUid field if non-nil, zero value otherwise.

### GetJobUidOk

`func (o *SnapshotHandle) GetJobUidOk() (*int64, bool)`

GetJobUidOk returns a tuple with the JobUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobUid

`func (o *SnapshotHandle) SetJobUid(v int64)`

SetJobUid sets JobUid field to given value.

### HasJobUid

`func (o *SnapshotHandle) HasJobUid() bool`

HasJobUid returns a boolean if a field has been set.

### SetJobUidNil

`func (o *SnapshotHandle) SetJobUidNil(b bool)`

 SetJobUidNil sets the value for JobUid to be an explicit nil

### UnsetJobUid
`func (o *SnapshotHandle) UnsetJobUid()`

UnsetJobUid ensures that no value is present for JobUid, not even an explicit nil
### GetRunStartTimeUsecs

`func (o *SnapshotHandle) GetRunStartTimeUsecs() int64`

GetRunStartTimeUsecs returns the RunStartTimeUsecs field if non-nil, zero value otherwise.

### GetRunStartTimeUsecsOk

`func (o *SnapshotHandle) GetRunStartTimeUsecsOk() (*int64, bool)`

GetRunStartTimeUsecsOk returns a tuple with the RunStartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunStartTimeUsecs

`func (o *SnapshotHandle) SetRunStartTimeUsecs(v int64)`

SetRunStartTimeUsecs sets RunStartTimeUsecs field to given value.

### HasRunStartTimeUsecs

`func (o *SnapshotHandle) HasRunStartTimeUsecs() bool`

HasRunStartTimeUsecs returns a boolean if a field has been set.

### SetRunStartTimeUsecsNil

`func (o *SnapshotHandle) SetRunStartTimeUsecsNil(b bool)`

 SetRunStartTimeUsecsNil sets the value for RunStartTimeUsecs to be an explicit nil

### UnsetRunStartTimeUsecs
`func (o *SnapshotHandle) UnsetRunStartTimeUsecs()`

UnsetRunStartTimeUsecs ensures that no value is present for RunStartTimeUsecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


