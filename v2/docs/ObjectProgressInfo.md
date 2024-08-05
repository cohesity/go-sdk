# ObjectProgressInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | Pointer to **NullableString** | Specifies the environment of the object. | [optional] 
**Id** | Pointer to **NullableInt64** | Specifies object id. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the object. | [optional] 
**SourceId** | Pointer to **NullableInt64** | Specifies registered source id to which object belongs. | [optional] 
**SourceName** | Pointer to **NullableString** | Specifies registered source name to which object belongs. | [optional] 
**EndTimeUsecs** | Pointer to **NullableInt64** | Specifies the end time of the progress task in Unix epoch Timestamp(in microseconds). | [optional] 
**Events** | Pointer to [**[]ProgressTaskEvent**](ProgressTaskEvent.md) | Specifies the event log created for progress Task. | [optional] 
**ExpectedRemainingTimeUsecs** | Pointer to **NullableInt64** | Specifies the expected remaining time of the progress task in Unix epoch Timestamp(in microseconds). | [optional] 
**PercentageCompleted** | Pointer to **NullableFloat32** | Specifies the current completed percentage of the progress task. | [optional] 
**StartTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time of the progress task in Unix epoch Timestamp(in microseconds). | [optional] 
**Stats** | Pointer to [**ProgressStats**](ProgressStats.md) |  | [optional] 
**Status** | Pointer to **NullableString** | Specifies the current status of the progress task. | [optional] 
**FailedAttempts** | Pointer to [**[]ProgressTaskInfo**](ProgressTaskInfo.md) | Specifies progress for failed attempts of this object. | [optional] 

## Methods

### NewObjectProgressInfo

`func NewObjectProgressInfo() *ObjectProgressInfo`

NewObjectProgressInfo instantiates a new ObjectProgressInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectProgressInfoWithDefaults

`func NewObjectProgressInfoWithDefaults() *ObjectProgressInfo`

NewObjectProgressInfoWithDefaults instantiates a new ObjectProgressInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *ObjectProgressInfo) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ObjectProgressInfo) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ObjectProgressInfo) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ObjectProgressInfo) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *ObjectProgressInfo) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *ObjectProgressInfo) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetId

`func (o *ObjectProgressInfo) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObjectProgressInfo) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObjectProgressInfo) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ObjectProgressInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ObjectProgressInfo) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ObjectProgressInfo) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *ObjectProgressInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ObjectProgressInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ObjectProgressInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ObjectProgressInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ObjectProgressInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ObjectProgressInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSourceId

`func (o *ObjectProgressInfo) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *ObjectProgressInfo) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *ObjectProgressInfo) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *ObjectProgressInfo) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *ObjectProgressInfo) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *ObjectProgressInfo) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetSourceName

`func (o *ObjectProgressInfo) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *ObjectProgressInfo) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *ObjectProgressInfo) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *ObjectProgressInfo) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *ObjectProgressInfo) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *ObjectProgressInfo) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil
### GetEndTimeUsecs

`func (o *ObjectProgressInfo) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *ObjectProgressInfo) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *ObjectProgressInfo) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *ObjectProgressInfo) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *ObjectProgressInfo) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *ObjectProgressInfo) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetEvents

`func (o *ObjectProgressInfo) GetEvents() []ProgressTaskEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ObjectProgressInfo) GetEventsOk() (*[]ProgressTaskEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ObjectProgressInfo) SetEvents(v []ProgressTaskEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *ObjectProgressInfo) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetExpectedRemainingTimeUsecs

`func (o *ObjectProgressInfo) GetExpectedRemainingTimeUsecs() int64`

GetExpectedRemainingTimeUsecs returns the ExpectedRemainingTimeUsecs field if non-nil, zero value otherwise.

### GetExpectedRemainingTimeUsecsOk

`func (o *ObjectProgressInfo) GetExpectedRemainingTimeUsecsOk() (*int64, bool)`

GetExpectedRemainingTimeUsecsOk returns a tuple with the ExpectedRemainingTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedRemainingTimeUsecs

`func (o *ObjectProgressInfo) SetExpectedRemainingTimeUsecs(v int64)`

SetExpectedRemainingTimeUsecs sets ExpectedRemainingTimeUsecs field to given value.

### HasExpectedRemainingTimeUsecs

`func (o *ObjectProgressInfo) HasExpectedRemainingTimeUsecs() bool`

HasExpectedRemainingTimeUsecs returns a boolean if a field has been set.

### SetExpectedRemainingTimeUsecsNil

`func (o *ObjectProgressInfo) SetExpectedRemainingTimeUsecsNil(b bool)`

 SetExpectedRemainingTimeUsecsNil sets the value for ExpectedRemainingTimeUsecs to be an explicit nil

### UnsetExpectedRemainingTimeUsecs
`func (o *ObjectProgressInfo) UnsetExpectedRemainingTimeUsecs()`

UnsetExpectedRemainingTimeUsecs ensures that no value is present for ExpectedRemainingTimeUsecs, not even an explicit nil
### GetPercentageCompleted

`func (o *ObjectProgressInfo) GetPercentageCompleted() float32`

GetPercentageCompleted returns the PercentageCompleted field if non-nil, zero value otherwise.

### GetPercentageCompletedOk

`func (o *ObjectProgressInfo) GetPercentageCompletedOk() (*float32, bool)`

GetPercentageCompletedOk returns a tuple with the PercentageCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageCompleted

`func (o *ObjectProgressInfo) SetPercentageCompleted(v float32)`

SetPercentageCompleted sets PercentageCompleted field to given value.

### HasPercentageCompleted

`func (o *ObjectProgressInfo) HasPercentageCompleted() bool`

HasPercentageCompleted returns a boolean if a field has been set.

### SetPercentageCompletedNil

`func (o *ObjectProgressInfo) SetPercentageCompletedNil(b bool)`

 SetPercentageCompletedNil sets the value for PercentageCompleted to be an explicit nil

### UnsetPercentageCompleted
`func (o *ObjectProgressInfo) UnsetPercentageCompleted()`

UnsetPercentageCompleted ensures that no value is present for PercentageCompleted, not even an explicit nil
### GetStartTimeUsecs

`func (o *ObjectProgressInfo) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *ObjectProgressInfo) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *ObjectProgressInfo) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *ObjectProgressInfo) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *ObjectProgressInfo) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *ObjectProgressInfo) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetStats

`func (o *ObjectProgressInfo) GetStats() ProgressStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ObjectProgressInfo) GetStatsOk() (*ProgressStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ObjectProgressInfo) SetStats(v ProgressStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ObjectProgressInfo) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetStatus

`func (o *ObjectProgressInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ObjectProgressInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ObjectProgressInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ObjectProgressInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ObjectProgressInfo) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ObjectProgressInfo) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetFailedAttempts

`func (o *ObjectProgressInfo) GetFailedAttempts() []ProgressTaskInfo`

GetFailedAttempts returns the FailedAttempts field if non-nil, zero value otherwise.

### GetFailedAttemptsOk

`func (o *ObjectProgressInfo) GetFailedAttemptsOk() (*[]ProgressTaskInfo, bool)`

GetFailedAttemptsOk returns a tuple with the FailedAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAttempts

`func (o *ObjectProgressInfo) SetFailedAttempts(v []ProgressTaskInfo)`

SetFailedAttempts sets FailedAttempts field to given value.

### HasFailedAttempts

`func (o *ObjectProgressInfo) HasFailedAttempts() bool`

HasFailedAttempts returns a boolean if a field has been set.

### SetFailedAttemptsNil

`func (o *ObjectProgressInfo) SetFailedAttemptsNil(b bool)`

 SetFailedAttemptsNil sets the value for FailedAttempts to be an explicit nil

### UnsetFailedAttempts
`func (o *ObjectProgressInfo) UnsetFailedAttempts()`

UnsetFailedAttempts ensures that no value is present for FailedAttempts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


