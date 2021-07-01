# JobRunsTile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastDayNumJobErrors** | Pointer to **NullableInt32** | Number of Error runs in the last 24 hours. | [optional] 
**LastDayNumJobRuns** | Pointer to **NullableInt32** | Number of Job Runs in the last 24 hours. | [optional] 
**LastDayNumJobSlaViolations** | Pointer to **NullableInt32** | Number of SLA Violations in the last 24 hours. | [optional] 
**NumJobRunning** | Pointer to **NullableInt32** | Number of Jobs currently running. | [optional] 
**ObjectsProtectedByPolicy** | Pointer to [**[]ObjectsProtectedByPolicy**](ObjectsProtectedByPolicy.md) | Objects Protected By Policy. | [optional] 

## Methods

### NewJobRunsTile

`func NewJobRunsTile() *JobRunsTile`

NewJobRunsTile instantiates a new JobRunsTile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobRunsTileWithDefaults

`func NewJobRunsTileWithDefaults() *JobRunsTile`

NewJobRunsTileWithDefaults instantiates a new JobRunsTile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastDayNumJobErrors

`func (o *JobRunsTile) GetLastDayNumJobErrors() int32`

GetLastDayNumJobErrors returns the LastDayNumJobErrors field if non-nil, zero value otherwise.

### GetLastDayNumJobErrorsOk

`func (o *JobRunsTile) GetLastDayNumJobErrorsOk() (*int32, bool)`

GetLastDayNumJobErrorsOk returns a tuple with the LastDayNumJobErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDayNumJobErrors

`func (o *JobRunsTile) SetLastDayNumJobErrors(v int32)`

SetLastDayNumJobErrors sets LastDayNumJobErrors field to given value.

### HasLastDayNumJobErrors

`func (o *JobRunsTile) HasLastDayNumJobErrors() bool`

HasLastDayNumJobErrors returns a boolean if a field has been set.

### SetLastDayNumJobErrorsNil

`func (o *JobRunsTile) SetLastDayNumJobErrorsNil(b bool)`

 SetLastDayNumJobErrorsNil sets the value for LastDayNumJobErrors to be an explicit nil

### UnsetLastDayNumJobErrors
`func (o *JobRunsTile) UnsetLastDayNumJobErrors()`

UnsetLastDayNumJobErrors ensures that no value is present for LastDayNumJobErrors, not even an explicit nil
### GetLastDayNumJobRuns

`func (o *JobRunsTile) GetLastDayNumJobRuns() int32`

GetLastDayNumJobRuns returns the LastDayNumJobRuns field if non-nil, zero value otherwise.

### GetLastDayNumJobRunsOk

`func (o *JobRunsTile) GetLastDayNumJobRunsOk() (*int32, bool)`

GetLastDayNumJobRunsOk returns a tuple with the LastDayNumJobRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDayNumJobRuns

`func (o *JobRunsTile) SetLastDayNumJobRuns(v int32)`

SetLastDayNumJobRuns sets LastDayNumJobRuns field to given value.

### HasLastDayNumJobRuns

`func (o *JobRunsTile) HasLastDayNumJobRuns() bool`

HasLastDayNumJobRuns returns a boolean if a field has been set.

### SetLastDayNumJobRunsNil

`func (o *JobRunsTile) SetLastDayNumJobRunsNil(b bool)`

 SetLastDayNumJobRunsNil sets the value for LastDayNumJobRuns to be an explicit nil

### UnsetLastDayNumJobRuns
`func (o *JobRunsTile) UnsetLastDayNumJobRuns()`

UnsetLastDayNumJobRuns ensures that no value is present for LastDayNumJobRuns, not even an explicit nil
### GetLastDayNumJobSlaViolations

`func (o *JobRunsTile) GetLastDayNumJobSlaViolations() int32`

GetLastDayNumJobSlaViolations returns the LastDayNumJobSlaViolations field if non-nil, zero value otherwise.

### GetLastDayNumJobSlaViolationsOk

`func (o *JobRunsTile) GetLastDayNumJobSlaViolationsOk() (*int32, bool)`

GetLastDayNumJobSlaViolationsOk returns a tuple with the LastDayNumJobSlaViolations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDayNumJobSlaViolations

`func (o *JobRunsTile) SetLastDayNumJobSlaViolations(v int32)`

SetLastDayNumJobSlaViolations sets LastDayNumJobSlaViolations field to given value.

### HasLastDayNumJobSlaViolations

`func (o *JobRunsTile) HasLastDayNumJobSlaViolations() bool`

HasLastDayNumJobSlaViolations returns a boolean if a field has been set.

### SetLastDayNumJobSlaViolationsNil

`func (o *JobRunsTile) SetLastDayNumJobSlaViolationsNil(b bool)`

 SetLastDayNumJobSlaViolationsNil sets the value for LastDayNumJobSlaViolations to be an explicit nil

### UnsetLastDayNumJobSlaViolations
`func (o *JobRunsTile) UnsetLastDayNumJobSlaViolations()`

UnsetLastDayNumJobSlaViolations ensures that no value is present for LastDayNumJobSlaViolations, not even an explicit nil
### GetNumJobRunning

`func (o *JobRunsTile) GetNumJobRunning() int32`

GetNumJobRunning returns the NumJobRunning field if non-nil, zero value otherwise.

### GetNumJobRunningOk

`func (o *JobRunsTile) GetNumJobRunningOk() (*int32, bool)`

GetNumJobRunningOk returns a tuple with the NumJobRunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumJobRunning

`func (o *JobRunsTile) SetNumJobRunning(v int32)`

SetNumJobRunning sets NumJobRunning field to given value.

### HasNumJobRunning

`func (o *JobRunsTile) HasNumJobRunning() bool`

HasNumJobRunning returns a boolean if a field has been set.

### SetNumJobRunningNil

`func (o *JobRunsTile) SetNumJobRunningNil(b bool)`

 SetNumJobRunningNil sets the value for NumJobRunning to be an explicit nil

### UnsetNumJobRunning
`func (o *JobRunsTile) UnsetNumJobRunning()`

UnsetNumJobRunning ensures that no value is present for NumJobRunning, not even an explicit nil
### GetObjectsProtectedByPolicy

`func (o *JobRunsTile) GetObjectsProtectedByPolicy() []ObjectsProtectedByPolicy`

GetObjectsProtectedByPolicy returns the ObjectsProtectedByPolicy field if non-nil, zero value otherwise.

### GetObjectsProtectedByPolicyOk

`func (o *JobRunsTile) GetObjectsProtectedByPolicyOk() (*[]ObjectsProtectedByPolicy, bool)`

GetObjectsProtectedByPolicyOk returns a tuple with the ObjectsProtectedByPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectsProtectedByPolicy

`func (o *JobRunsTile) SetObjectsProtectedByPolicy(v []ObjectsProtectedByPolicy)`

SetObjectsProtectedByPolicy sets ObjectsProtectedByPolicy field to given value.

### HasObjectsProtectedByPolicy

`func (o *JobRunsTile) HasObjectsProtectedByPolicy() bool`

HasObjectsProtectedByPolicy returns a boolean if a field has been set.

### SetObjectsProtectedByPolicyNil

`func (o *JobRunsTile) SetObjectsProtectedByPolicyNil(b bool)`

 SetObjectsProtectedByPolicyNil sets the value for ObjectsProtectedByPolicy to be an explicit nil

### UnsetObjectsProtectedByPolicy
`func (o *JobRunsTile) UnsetObjectsProtectedByPolicy()`

UnsetObjectsProtectedByPolicy ensures that no value is present for ObjectsProtectedByPolicy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


