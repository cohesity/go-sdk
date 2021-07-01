# RunUid

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobUid** | Pointer to [**UniversalId**](UniversalId.md) |  | [optional] 
**StartTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time of the Protection Job Run. | [optional] 

## Methods

### NewRunUid

`func NewRunUid() *RunUid`

NewRunUid instantiates a new RunUid object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunUidWithDefaults

`func NewRunUidWithDefaults() *RunUid`

NewRunUidWithDefaults instantiates a new RunUid object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobUid

`func (o *RunUid) GetJobUid() UniversalId`

GetJobUid returns the JobUid field if non-nil, zero value otherwise.

### GetJobUidOk

`func (o *RunUid) GetJobUidOk() (*UniversalId, bool)`

GetJobUidOk returns a tuple with the JobUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobUid

`func (o *RunUid) SetJobUid(v UniversalId)`

SetJobUid sets JobUid field to given value.

### HasJobUid

`func (o *RunUid) HasJobUid() bool`

HasJobUid returns a boolean if a field has been set.

### GetStartTimeUsecs

`func (o *RunUid) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *RunUid) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *RunUid) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *RunUid) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *RunUid) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *RunUid) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


