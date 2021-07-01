# MagnetoInstanceId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttemptNum** | Pointer to **NullableInt64** | The attempt of the job instance that took the snapshot. | [optional] 
**JobInstanceId** | Pointer to **NullableInt64** | Instance of the job that took the snapshot. | [optional] 
**JobStartTimeUsecs** | Pointer to **NullableInt64** | Start time of the job instance. | [optional] 

## Methods

### NewMagnetoInstanceId

`func NewMagnetoInstanceId() *MagnetoInstanceId`

NewMagnetoInstanceId instantiates a new MagnetoInstanceId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMagnetoInstanceIdWithDefaults

`func NewMagnetoInstanceIdWithDefaults() *MagnetoInstanceId`

NewMagnetoInstanceIdWithDefaults instantiates a new MagnetoInstanceId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttemptNum

`func (o *MagnetoInstanceId) GetAttemptNum() int64`

GetAttemptNum returns the AttemptNum field if non-nil, zero value otherwise.

### GetAttemptNumOk

`func (o *MagnetoInstanceId) GetAttemptNumOk() (*int64, bool)`

GetAttemptNumOk returns a tuple with the AttemptNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptNum

`func (o *MagnetoInstanceId) SetAttemptNum(v int64)`

SetAttemptNum sets AttemptNum field to given value.

### HasAttemptNum

`func (o *MagnetoInstanceId) HasAttemptNum() bool`

HasAttemptNum returns a boolean if a field has been set.

### SetAttemptNumNil

`func (o *MagnetoInstanceId) SetAttemptNumNil(b bool)`

 SetAttemptNumNil sets the value for AttemptNum to be an explicit nil

### UnsetAttemptNum
`func (o *MagnetoInstanceId) UnsetAttemptNum()`

UnsetAttemptNum ensures that no value is present for AttemptNum, not even an explicit nil
### GetJobInstanceId

`func (o *MagnetoInstanceId) GetJobInstanceId() int64`

GetJobInstanceId returns the JobInstanceId field if non-nil, zero value otherwise.

### GetJobInstanceIdOk

`func (o *MagnetoInstanceId) GetJobInstanceIdOk() (*int64, bool)`

GetJobInstanceIdOk returns a tuple with the JobInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobInstanceId

`func (o *MagnetoInstanceId) SetJobInstanceId(v int64)`

SetJobInstanceId sets JobInstanceId field to given value.

### HasJobInstanceId

`func (o *MagnetoInstanceId) HasJobInstanceId() bool`

HasJobInstanceId returns a boolean if a field has been set.

### SetJobInstanceIdNil

`func (o *MagnetoInstanceId) SetJobInstanceIdNil(b bool)`

 SetJobInstanceIdNil sets the value for JobInstanceId to be an explicit nil

### UnsetJobInstanceId
`func (o *MagnetoInstanceId) UnsetJobInstanceId()`

UnsetJobInstanceId ensures that no value is present for JobInstanceId, not even an explicit nil
### GetJobStartTimeUsecs

`func (o *MagnetoInstanceId) GetJobStartTimeUsecs() int64`

GetJobStartTimeUsecs returns the JobStartTimeUsecs field if non-nil, zero value otherwise.

### GetJobStartTimeUsecsOk

`func (o *MagnetoInstanceId) GetJobStartTimeUsecsOk() (*int64, bool)`

GetJobStartTimeUsecsOk returns a tuple with the JobStartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobStartTimeUsecs

`func (o *MagnetoInstanceId) SetJobStartTimeUsecs(v int64)`

SetJobStartTimeUsecs sets JobStartTimeUsecs field to given value.

### HasJobStartTimeUsecs

`func (o *MagnetoInstanceId) HasJobStartTimeUsecs() bool`

HasJobStartTimeUsecs returns a boolean if a field has been set.

### SetJobStartTimeUsecsNil

`func (o *MagnetoInstanceId) SetJobStartTimeUsecsNil(b bool)`

 SetJobStartTimeUsecsNil sets the value for JobStartTimeUsecs to be an explicit nil

### UnsetJobStartTimeUsecs
`func (o *MagnetoInstanceId) UnsetJobStartTimeUsecs()`

UnsetJobStartTimeUsecs ensures that no value is present for JobStartTimeUsecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


