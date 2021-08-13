# CloudSpinTargetResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt64** | Specifies the unique id of the cloud spin entity. | [optional] 
**AwsParams** | Pointer to [**AwsCloudSpinParams**](AwsCloudSpinParams.md) |  | [optional] 
**AzureParams** | Pointer to [**AzureCloudSpinParams**](AzureCloudSpinParams.md) |  | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the already added cloud spin target. | [optional] [readonly] 
**StartTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time of Cloud Spin in Unix epoch Timestamp(in microseconds) for a target. | [optional] 
**EndTimeUsecs** | Pointer to **NullableInt64** | Specifies the end time of Cloud Spin in Unix epoch Timestamp(in microseconds) for a target. | [optional] 
**Status** | Pointer to **NullableString** | Status of the Cloud Spin for a target. &#39;Running&#39; indicates that the run is still running. &#39;Canceled&#39; indicates that the run has been canceled. &#39;Canceling&#39; indicates that the run is in the process of being canceled. &#39;Paused&#39; indicates that the ongoing run has been paused. &#39;Failed&#39; indicates that the run has failed. &#39;Missed&#39; indicates that the run was unable to take place at the scheduled time because the previous run was still happening. &#39;Succeeded&#39; indicates that the run has finished successfully. &#39;SucceededWithWarning&#39; indicates that the run finished successfully, but there were some warning messages. | [optional] 
**Message** | Pointer to **NullableString** | Message about the Cloud Spin run. | [optional] 
**Stats** | Pointer to [**CloudSpinDataStats**](CloudSpinDataStats.md) |  | [optional] 
**IsManuallyDeleted** | Pointer to **NullableBool** | Specifies whether the snapshot is deleted manually. | [optional] 
**ExpiryTimeUsecs** | Pointer to **NullableInt64** | Specifies the expiry time of attempt in Unix epoch Timestamp (in microseconds) for an object. | [optional] 
**CloudspinTaskId** | Pointer to **NullableString** | Task ID for a CloudSpin protection run. | [optional] 
**ProgressTaskId** | Pointer to **NullableString** | Progress monitor task id for Cloud Spin run. | [optional] 
**DataLockConstraints** | Pointer to [**DataLockConstraints**](DataLockConstraints.md) |  | [optional] 
**OnLegalHold** | Pointer to **NullableBool** | Specifies the legal hold status for a cloud spin target. | [optional] 

## Methods

### NewCloudSpinTargetResult

`func NewCloudSpinTargetResult() *CloudSpinTargetResult`

NewCloudSpinTargetResult instantiates a new CloudSpinTargetResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSpinTargetResultWithDefaults

`func NewCloudSpinTargetResultWithDefaults() *CloudSpinTargetResult`

NewCloudSpinTargetResultWithDefaults instantiates a new CloudSpinTargetResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CloudSpinTargetResult) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudSpinTargetResult) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudSpinTargetResult) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CloudSpinTargetResult) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CloudSpinTargetResult) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CloudSpinTargetResult) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAwsParams

`func (o *CloudSpinTargetResult) GetAwsParams() AwsCloudSpinParams`

GetAwsParams returns the AwsParams field if non-nil, zero value otherwise.

### GetAwsParamsOk

`func (o *CloudSpinTargetResult) GetAwsParamsOk() (*AwsCloudSpinParams, bool)`

GetAwsParamsOk returns a tuple with the AwsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsParams

`func (o *CloudSpinTargetResult) SetAwsParams(v AwsCloudSpinParams)`

SetAwsParams sets AwsParams field to given value.

### HasAwsParams

`func (o *CloudSpinTargetResult) HasAwsParams() bool`

HasAwsParams returns a boolean if a field has been set.

### GetAzureParams

`func (o *CloudSpinTargetResult) GetAzureParams() AzureCloudSpinParams`

GetAzureParams returns the AzureParams field if non-nil, zero value otherwise.

### GetAzureParamsOk

`func (o *CloudSpinTargetResult) GetAzureParamsOk() (*AzureCloudSpinParams, bool)`

GetAzureParamsOk returns a tuple with the AzureParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureParams

`func (o *CloudSpinTargetResult) SetAzureParams(v AzureCloudSpinParams)`

SetAzureParams sets AzureParams field to given value.

### HasAzureParams

`func (o *CloudSpinTargetResult) HasAzureParams() bool`

HasAzureParams returns a boolean if a field has been set.

### GetName

`func (o *CloudSpinTargetResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudSpinTargetResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudSpinTargetResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudSpinTargetResult) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CloudSpinTargetResult) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CloudSpinTargetResult) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStartTimeUsecs

`func (o *CloudSpinTargetResult) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *CloudSpinTargetResult) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *CloudSpinTargetResult) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *CloudSpinTargetResult) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *CloudSpinTargetResult) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *CloudSpinTargetResult) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetEndTimeUsecs

`func (o *CloudSpinTargetResult) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *CloudSpinTargetResult) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *CloudSpinTargetResult) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *CloudSpinTargetResult) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *CloudSpinTargetResult) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *CloudSpinTargetResult) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetStatus

`func (o *CloudSpinTargetResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudSpinTargetResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudSpinTargetResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudSpinTargetResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *CloudSpinTargetResult) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *CloudSpinTargetResult) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetMessage

`func (o *CloudSpinTargetResult) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CloudSpinTargetResult) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CloudSpinTargetResult) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CloudSpinTargetResult) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *CloudSpinTargetResult) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *CloudSpinTargetResult) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetStats

`func (o *CloudSpinTargetResult) GetStats() CloudSpinDataStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *CloudSpinTargetResult) GetStatsOk() (*CloudSpinDataStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *CloudSpinTargetResult) SetStats(v CloudSpinDataStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *CloudSpinTargetResult) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetIsManuallyDeleted

`func (o *CloudSpinTargetResult) GetIsManuallyDeleted() bool`

GetIsManuallyDeleted returns the IsManuallyDeleted field if non-nil, zero value otherwise.

### GetIsManuallyDeletedOk

`func (o *CloudSpinTargetResult) GetIsManuallyDeletedOk() (*bool, bool)`

GetIsManuallyDeletedOk returns a tuple with the IsManuallyDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManuallyDeleted

`func (o *CloudSpinTargetResult) SetIsManuallyDeleted(v bool)`

SetIsManuallyDeleted sets IsManuallyDeleted field to given value.

### HasIsManuallyDeleted

`func (o *CloudSpinTargetResult) HasIsManuallyDeleted() bool`

HasIsManuallyDeleted returns a boolean if a field has been set.

### SetIsManuallyDeletedNil

`func (o *CloudSpinTargetResult) SetIsManuallyDeletedNil(b bool)`

 SetIsManuallyDeletedNil sets the value for IsManuallyDeleted to be an explicit nil

### UnsetIsManuallyDeleted
`func (o *CloudSpinTargetResult) UnsetIsManuallyDeleted()`

UnsetIsManuallyDeleted ensures that no value is present for IsManuallyDeleted, not even an explicit nil
### GetExpiryTimeUsecs

`func (o *CloudSpinTargetResult) GetExpiryTimeUsecs() int64`

GetExpiryTimeUsecs returns the ExpiryTimeUsecs field if non-nil, zero value otherwise.

### GetExpiryTimeUsecsOk

`func (o *CloudSpinTargetResult) GetExpiryTimeUsecsOk() (*int64, bool)`

GetExpiryTimeUsecsOk returns a tuple with the ExpiryTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTimeUsecs

`func (o *CloudSpinTargetResult) SetExpiryTimeUsecs(v int64)`

SetExpiryTimeUsecs sets ExpiryTimeUsecs field to given value.

### HasExpiryTimeUsecs

`func (o *CloudSpinTargetResult) HasExpiryTimeUsecs() bool`

HasExpiryTimeUsecs returns a boolean if a field has been set.

### SetExpiryTimeUsecsNil

`func (o *CloudSpinTargetResult) SetExpiryTimeUsecsNil(b bool)`

 SetExpiryTimeUsecsNil sets the value for ExpiryTimeUsecs to be an explicit nil

### UnsetExpiryTimeUsecs
`func (o *CloudSpinTargetResult) UnsetExpiryTimeUsecs()`

UnsetExpiryTimeUsecs ensures that no value is present for ExpiryTimeUsecs, not even an explicit nil
### GetCloudspinTaskId

`func (o *CloudSpinTargetResult) GetCloudspinTaskId() string`

GetCloudspinTaskId returns the CloudspinTaskId field if non-nil, zero value otherwise.

### GetCloudspinTaskIdOk

`func (o *CloudSpinTargetResult) GetCloudspinTaskIdOk() (*string, bool)`

GetCloudspinTaskIdOk returns a tuple with the CloudspinTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudspinTaskId

`func (o *CloudSpinTargetResult) SetCloudspinTaskId(v string)`

SetCloudspinTaskId sets CloudspinTaskId field to given value.

### HasCloudspinTaskId

`func (o *CloudSpinTargetResult) HasCloudspinTaskId() bool`

HasCloudspinTaskId returns a boolean if a field has been set.

### SetCloudspinTaskIdNil

`func (o *CloudSpinTargetResult) SetCloudspinTaskIdNil(b bool)`

 SetCloudspinTaskIdNil sets the value for CloudspinTaskId to be an explicit nil

### UnsetCloudspinTaskId
`func (o *CloudSpinTargetResult) UnsetCloudspinTaskId()`

UnsetCloudspinTaskId ensures that no value is present for CloudspinTaskId, not even an explicit nil
### GetProgressTaskId

`func (o *CloudSpinTargetResult) GetProgressTaskId() string`

GetProgressTaskId returns the ProgressTaskId field if non-nil, zero value otherwise.

### GetProgressTaskIdOk

`func (o *CloudSpinTargetResult) GetProgressTaskIdOk() (*string, bool)`

GetProgressTaskIdOk returns a tuple with the ProgressTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressTaskId

`func (o *CloudSpinTargetResult) SetProgressTaskId(v string)`

SetProgressTaskId sets ProgressTaskId field to given value.

### HasProgressTaskId

`func (o *CloudSpinTargetResult) HasProgressTaskId() bool`

HasProgressTaskId returns a boolean if a field has been set.

### SetProgressTaskIdNil

`func (o *CloudSpinTargetResult) SetProgressTaskIdNil(b bool)`

 SetProgressTaskIdNil sets the value for ProgressTaskId to be an explicit nil

### UnsetProgressTaskId
`func (o *CloudSpinTargetResult) UnsetProgressTaskId()`

UnsetProgressTaskId ensures that no value is present for ProgressTaskId, not even an explicit nil
### GetDataLockConstraints

`func (o *CloudSpinTargetResult) GetDataLockConstraints() DataLockConstraints`

GetDataLockConstraints returns the DataLockConstraints field if non-nil, zero value otherwise.

### GetDataLockConstraintsOk

`func (o *CloudSpinTargetResult) GetDataLockConstraintsOk() (*DataLockConstraints, bool)`

GetDataLockConstraintsOk returns a tuple with the DataLockConstraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLockConstraints

`func (o *CloudSpinTargetResult) SetDataLockConstraints(v DataLockConstraints)`

SetDataLockConstraints sets DataLockConstraints field to given value.

### HasDataLockConstraints

`func (o *CloudSpinTargetResult) HasDataLockConstraints() bool`

HasDataLockConstraints returns a boolean if a field has been set.

### GetOnLegalHold

`func (o *CloudSpinTargetResult) GetOnLegalHold() bool`

GetOnLegalHold returns the OnLegalHold field if non-nil, zero value otherwise.

### GetOnLegalHoldOk

`func (o *CloudSpinTargetResult) GetOnLegalHoldOk() (*bool, bool)`

GetOnLegalHoldOk returns a tuple with the OnLegalHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnLegalHold

`func (o *CloudSpinTargetResult) SetOnLegalHold(v bool)`

SetOnLegalHold sets OnLegalHold field to given value.

### HasOnLegalHold

`func (o *CloudSpinTargetResult) HasOnLegalHold() bool`

HasOnLegalHold returns a boolean if a field has been set.

### SetOnLegalHoldNil

`func (o *CloudSpinTargetResult) SetOnLegalHoldNil(b bool)`

 SetOnLegalHoldNil sets the value for OnLegalHold to be an explicit nil

### UnsetOnLegalHold
`func (o *CloudSpinTargetResult) UnsetOnLegalHold()`

UnsetOnLegalHold ensures that no value is present for OnLegalHold, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


