# FailoverRunsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailoverPlannedRuns** | Pointer to [**[]PlannedRunPollStatus**](PlannedRunPollStatus.md) | Specifies the list of planned runs created during various planeed failover workflows. Each planned run is uniqely identified by falioverId and runId. | [optional] 

## Methods

### NewFailoverRunsResponse

`func NewFailoverRunsResponse() *FailoverRunsResponse`

NewFailoverRunsResponse instantiates a new FailoverRunsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailoverRunsResponseWithDefaults

`func NewFailoverRunsResponseWithDefaults() *FailoverRunsResponse`

NewFailoverRunsResponseWithDefaults instantiates a new FailoverRunsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailoverPlannedRuns

`func (o *FailoverRunsResponse) GetFailoverPlannedRuns() []PlannedRunPollStatus`

GetFailoverPlannedRuns returns the FailoverPlannedRuns field if non-nil, zero value otherwise.

### GetFailoverPlannedRunsOk

`func (o *FailoverRunsResponse) GetFailoverPlannedRunsOk() (*[]PlannedRunPollStatus, bool)`

GetFailoverPlannedRunsOk returns a tuple with the FailoverPlannedRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverPlannedRuns

`func (o *FailoverRunsResponse) SetFailoverPlannedRuns(v []PlannedRunPollStatus)`

SetFailoverPlannedRuns sets FailoverPlannedRuns field to given value.

### HasFailoverPlannedRuns

`func (o *FailoverRunsResponse) HasFailoverPlannedRuns() bool`

HasFailoverPlannedRuns returns a boolean if a field has been set.

### SetFailoverPlannedRunsNil

`func (o *FailoverRunsResponse) SetFailoverPlannedRunsNil(b bool)`

 SetFailoverPlannedRunsNil sets the value for FailoverPlannedRuns to be an explicit nil

### UnsetFailoverPlannedRuns
`func (o *FailoverRunsResponse) UnsetFailoverPlannedRuns()`

UnsetFailoverPlannedRuns ensures that no value is present for FailoverPlannedRuns, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


