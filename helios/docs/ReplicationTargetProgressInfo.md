# ReplicationTargetProgressInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **NullableInt64** | Specifies the id of the cluster. | [optional] 
**ClusterIncarnationId** | Pointer to **NullableInt64** | Specifies the incarnation id of the cluster. | [optional] 
**ClusterName** | Pointer to **NullableString** | Specifies the name of the cluster. | [optional] [readonly] 
**AwsTargetConfig** | Pointer to [**AWSTargetConfig**](AWSTargetConfig.md) |  | [optional] 
**AzureTargetConfig** | Pointer to [**AzureTargetConfig**](AzureTargetConfig.md) |  | [optional] 
**Status** | Pointer to **NullableString** | Specifies the current status of the progress task. | [optional] 
**PercentageCompleted** | Pointer to **NullableFloat32** | Specifies the current completed percentage of the progress task. | [optional] 
**StartTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time of the progress task in Unix epoch Timestamp(in microseconds). | [optional] 
**EndTimeUsecs** | Pointer to **NullableInt64** | Specifies the end time of the progress task in Unix epoch Timestamp(in microseconds). | [optional] 
**ExpectedRemainingTimeUsecs** | Pointer to **NullableInt64** | Specifies the expected remaining time of the progress task in Unix epoch Timestamp(in microseconds). | [optional] 
**Events** | Pointer to [**[]ProgressTaskEvent**](ProgressTaskEvent.md) | Specifies the event log created for progress Task. | [optional] 
**Stats** | Pointer to [**ProgressStats**](ProgressStats.md) |  | [optional] 
**Objects** | Pointer to [**[]ObjectProgressInfo**](ObjectProgressInfo.md) | Specifies progress for objects. | [optional] 

## Methods

### NewReplicationTargetProgressInfo

`func NewReplicationTargetProgressInfo() *ReplicationTargetProgressInfo`

NewReplicationTargetProgressInfo instantiates a new ReplicationTargetProgressInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicationTargetProgressInfoWithDefaults

`func NewReplicationTargetProgressInfoWithDefaults() *ReplicationTargetProgressInfo`

NewReplicationTargetProgressInfoWithDefaults instantiates a new ReplicationTargetProgressInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *ReplicationTargetProgressInfo) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ReplicationTargetProgressInfo) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ReplicationTargetProgressInfo) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *ReplicationTargetProgressInfo) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### SetClusterIdNil

`func (o *ReplicationTargetProgressInfo) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *ReplicationTargetProgressInfo) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetClusterIncarnationId

`func (o *ReplicationTargetProgressInfo) GetClusterIncarnationId() int64`

GetClusterIncarnationId returns the ClusterIncarnationId field if non-nil, zero value otherwise.

### GetClusterIncarnationIdOk

`func (o *ReplicationTargetProgressInfo) GetClusterIncarnationIdOk() (*int64, bool)`

GetClusterIncarnationIdOk returns a tuple with the ClusterIncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIncarnationId

`func (o *ReplicationTargetProgressInfo) SetClusterIncarnationId(v int64)`

SetClusterIncarnationId sets ClusterIncarnationId field to given value.

### HasClusterIncarnationId

`func (o *ReplicationTargetProgressInfo) HasClusterIncarnationId() bool`

HasClusterIncarnationId returns a boolean if a field has been set.

### SetClusterIncarnationIdNil

`func (o *ReplicationTargetProgressInfo) SetClusterIncarnationIdNil(b bool)`

 SetClusterIncarnationIdNil sets the value for ClusterIncarnationId to be an explicit nil

### UnsetClusterIncarnationId
`func (o *ReplicationTargetProgressInfo) UnsetClusterIncarnationId()`

UnsetClusterIncarnationId ensures that no value is present for ClusterIncarnationId, not even an explicit nil
### GetClusterName

`func (o *ReplicationTargetProgressInfo) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *ReplicationTargetProgressInfo) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *ReplicationTargetProgressInfo) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *ReplicationTargetProgressInfo) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### SetClusterNameNil

`func (o *ReplicationTargetProgressInfo) SetClusterNameNil(b bool)`

 SetClusterNameNil sets the value for ClusterName to be an explicit nil

### UnsetClusterName
`func (o *ReplicationTargetProgressInfo) UnsetClusterName()`

UnsetClusterName ensures that no value is present for ClusterName, not even an explicit nil
### GetAwsTargetConfig

`func (o *ReplicationTargetProgressInfo) GetAwsTargetConfig() AWSTargetConfig`

GetAwsTargetConfig returns the AwsTargetConfig field if non-nil, zero value otherwise.

### GetAwsTargetConfigOk

`func (o *ReplicationTargetProgressInfo) GetAwsTargetConfigOk() (*AWSTargetConfig, bool)`

GetAwsTargetConfigOk returns a tuple with the AwsTargetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsTargetConfig

`func (o *ReplicationTargetProgressInfo) SetAwsTargetConfig(v AWSTargetConfig)`

SetAwsTargetConfig sets AwsTargetConfig field to given value.

### HasAwsTargetConfig

`func (o *ReplicationTargetProgressInfo) HasAwsTargetConfig() bool`

HasAwsTargetConfig returns a boolean if a field has been set.

### GetAzureTargetConfig

`func (o *ReplicationTargetProgressInfo) GetAzureTargetConfig() AzureTargetConfig`

GetAzureTargetConfig returns the AzureTargetConfig field if non-nil, zero value otherwise.

### GetAzureTargetConfigOk

`func (o *ReplicationTargetProgressInfo) GetAzureTargetConfigOk() (*AzureTargetConfig, bool)`

GetAzureTargetConfigOk returns a tuple with the AzureTargetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTargetConfig

`func (o *ReplicationTargetProgressInfo) SetAzureTargetConfig(v AzureTargetConfig)`

SetAzureTargetConfig sets AzureTargetConfig field to given value.

### HasAzureTargetConfig

`func (o *ReplicationTargetProgressInfo) HasAzureTargetConfig() bool`

HasAzureTargetConfig returns a boolean if a field has been set.

### GetStatus

`func (o *ReplicationTargetProgressInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReplicationTargetProgressInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReplicationTargetProgressInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReplicationTargetProgressInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ReplicationTargetProgressInfo) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ReplicationTargetProgressInfo) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetPercentageCompleted

`func (o *ReplicationTargetProgressInfo) GetPercentageCompleted() float32`

GetPercentageCompleted returns the PercentageCompleted field if non-nil, zero value otherwise.

### GetPercentageCompletedOk

`func (o *ReplicationTargetProgressInfo) GetPercentageCompletedOk() (*float32, bool)`

GetPercentageCompletedOk returns a tuple with the PercentageCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageCompleted

`func (o *ReplicationTargetProgressInfo) SetPercentageCompleted(v float32)`

SetPercentageCompleted sets PercentageCompleted field to given value.

### HasPercentageCompleted

`func (o *ReplicationTargetProgressInfo) HasPercentageCompleted() bool`

HasPercentageCompleted returns a boolean if a field has been set.

### SetPercentageCompletedNil

`func (o *ReplicationTargetProgressInfo) SetPercentageCompletedNil(b bool)`

 SetPercentageCompletedNil sets the value for PercentageCompleted to be an explicit nil

### UnsetPercentageCompleted
`func (o *ReplicationTargetProgressInfo) UnsetPercentageCompleted()`

UnsetPercentageCompleted ensures that no value is present for PercentageCompleted, not even an explicit nil
### GetStartTimeUsecs

`func (o *ReplicationTargetProgressInfo) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *ReplicationTargetProgressInfo) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *ReplicationTargetProgressInfo) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *ReplicationTargetProgressInfo) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *ReplicationTargetProgressInfo) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *ReplicationTargetProgressInfo) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetEndTimeUsecs

`func (o *ReplicationTargetProgressInfo) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *ReplicationTargetProgressInfo) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *ReplicationTargetProgressInfo) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *ReplicationTargetProgressInfo) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *ReplicationTargetProgressInfo) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *ReplicationTargetProgressInfo) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetExpectedRemainingTimeUsecs

`func (o *ReplicationTargetProgressInfo) GetExpectedRemainingTimeUsecs() int64`

GetExpectedRemainingTimeUsecs returns the ExpectedRemainingTimeUsecs field if non-nil, zero value otherwise.

### GetExpectedRemainingTimeUsecsOk

`func (o *ReplicationTargetProgressInfo) GetExpectedRemainingTimeUsecsOk() (*int64, bool)`

GetExpectedRemainingTimeUsecsOk returns a tuple with the ExpectedRemainingTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedRemainingTimeUsecs

`func (o *ReplicationTargetProgressInfo) SetExpectedRemainingTimeUsecs(v int64)`

SetExpectedRemainingTimeUsecs sets ExpectedRemainingTimeUsecs field to given value.

### HasExpectedRemainingTimeUsecs

`func (o *ReplicationTargetProgressInfo) HasExpectedRemainingTimeUsecs() bool`

HasExpectedRemainingTimeUsecs returns a boolean if a field has been set.

### SetExpectedRemainingTimeUsecsNil

`func (o *ReplicationTargetProgressInfo) SetExpectedRemainingTimeUsecsNil(b bool)`

 SetExpectedRemainingTimeUsecsNil sets the value for ExpectedRemainingTimeUsecs to be an explicit nil

### UnsetExpectedRemainingTimeUsecs
`func (o *ReplicationTargetProgressInfo) UnsetExpectedRemainingTimeUsecs()`

UnsetExpectedRemainingTimeUsecs ensures that no value is present for ExpectedRemainingTimeUsecs, not even an explicit nil
### GetEvents

`func (o *ReplicationTargetProgressInfo) GetEvents() []ProgressTaskEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ReplicationTargetProgressInfo) GetEventsOk() (*[]ProgressTaskEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ReplicationTargetProgressInfo) SetEvents(v []ProgressTaskEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *ReplicationTargetProgressInfo) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetStats

`func (o *ReplicationTargetProgressInfo) GetStats() ProgressStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ReplicationTargetProgressInfo) GetStatsOk() (*ProgressStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ReplicationTargetProgressInfo) SetStats(v ProgressStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ReplicationTargetProgressInfo) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetObjects

`func (o *ReplicationTargetProgressInfo) GetObjects() []ObjectProgressInfo`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *ReplicationTargetProgressInfo) GetObjectsOk() (*[]ObjectProgressInfo, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *ReplicationTargetProgressInfo) SetObjects(v []ObjectProgressInfo)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *ReplicationTargetProgressInfo) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### SetObjectsNil

`func (o *ReplicationTargetProgressInfo) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *ReplicationTargetProgressInfo) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


