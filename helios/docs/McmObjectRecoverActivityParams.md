# McmObjectRecoverActivityParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Specifies the id of the Recovery. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the Recovery. | [optional] 
**StartTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time of the Recovery in Unix timestamp epoch in microseconds. | [optional] 
**EndTimeUsecs** | Pointer to **NullableInt64** | Specifies the end time of the Recovery in Unix timestamp epoch in microseconds. This field will be populated only after Recovery is finished. | [optional] 
**Status** | Pointer to **NullableString** | Status of the Recovery. &#39;Running&#39; indicates that the Recovery is still running. &#39;Canceled&#39; indicates that the Recovery has been cancelled. &#39;Canceling&#39; indicates that the Recovery is in the process of being cancelled. &#39;Failed&#39; indicates that the Recovery has failed. &#39;Succeeded&#39; indicates that the Recovery has finished successfully. &#39;SucceededWithWarning&#39; indicates that the Recovery finished successfully, but there were some warning messages. | [optional] 
**ProgressTaskId** | Pointer to **NullableString** | Progress monitor task id for Recovery. | [optional] 
**RecoveryType** | Pointer to **NullableString** | Specifies the recovery type. | [optional] 

## Methods

### NewMcmObjectRecoverActivityParams

`func NewMcmObjectRecoverActivityParams() *McmObjectRecoverActivityParams`

NewMcmObjectRecoverActivityParams instantiates a new McmObjectRecoverActivityParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcmObjectRecoverActivityParamsWithDefaults

`func NewMcmObjectRecoverActivityParamsWithDefaults() *McmObjectRecoverActivityParams`

NewMcmObjectRecoverActivityParamsWithDefaults instantiates a new McmObjectRecoverActivityParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *McmObjectRecoverActivityParams) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *McmObjectRecoverActivityParams) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *McmObjectRecoverActivityParams) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *McmObjectRecoverActivityParams) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *McmObjectRecoverActivityParams) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *McmObjectRecoverActivityParams) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *McmObjectRecoverActivityParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *McmObjectRecoverActivityParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *McmObjectRecoverActivityParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *McmObjectRecoverActivityParams) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *McmObjectRecoverActivityParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *McmObjectRecoverActivityParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStartTimeUsecs

`func (o *McmObjectRecoverActivityParams) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *McmObjectRecoverActivityParams) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *McmObjectRecoverActivityParams) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *McmObjectRecoverActivityParams) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *McmObjectRecoverActivityParams) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *McmObjectRecoverActivityParams) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetEndTimeUsecs

`func (o *McmObjectRecoverActivityParams) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *McmObjectRecoverActivityParams) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *McmObjectRecoverActivityParams) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *McmObjectRecoverActivityParams) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *McmObjectRecoverActivityParams) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *McmObjectRecoverActivityParams) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetStatus

`func (o *McmObjectRecoverActivityParams) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *McmObjectRecoverActivityParams) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *McmObjectRecoverActivityParams) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *McmObjectRecoverActivityParams) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *McmObjectRecoverActivityParams) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *McmObjectRecoverActivityParams) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetProgressTaskId

`func (o *McmObjectRecoverActivityParams) GetProgressTaskId() string`

GetProgressTaskId returns the ProgressTaskId field if non-nil, zero value otherwise.

### GetProgressTaskIdOk

`func (o *McmObjectRecoverActivityParams) GetProgressTaskIdOk() (*string, bool)`

GetProgressTaskIdOk returns a tuple with the ProgressTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressTaskId

`func (o *McmObjectRecoverActivityParams) SetProgressTaskId(v string)`

SetProgressTaskId sets ProgressTaskId field to given value.

### HasProgressTaskId

`func (o *McmObjectRecoverActivityParams) HasProgressTaskId() bool`

HasProgressTaskId returns a boolean if a field has been set.

### SetProgressTaskIdNil

`func (o *McmObjectRecoverActivityParams) SetProgressTaskIdNil(b bool)`

 SetProgressTaskIdNil sets the value for ProgressTaskId to be an explicit nil

### UnsetProgressTaskId
`func (o *McmObjectRecoverActivityParams) UnsetProgressTaskId()`

UnsetProgressTaskId ensures that no value is present for ProgressTaskId, not even an explicit nil
### GetRecoveryType

`func (o *McmObjectRecoverActivityParams) GetRecoveryType() string`

GetRecoveryType returns the RecoveryType field if non-nil, zero value otherwise.

### GetRecoveryTypeOk

`func (o *McmObjectRecoverActivityParams) GetRecoveryTypeOk() (*string, bool)`

GetRecoveryTypeOk returns a tuple with the RecoveryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryType

`func (o *McmObjectRecoverActivityParams) SetRecoveryType(v string)`

SetRecoveryType sets RecoveryType field to given value.

### HasRecoveryType

`func (o *McmObjectRecoverActivityParams) HasRecoveryType() bool`

HasRecoveryType returns a boolean if a field has been set.

### SetRecoveryTypeNil

`func (o *McmObjectRecoverActivityParams) SetRecoveryTypeNil(b bool)`

 SetRecoveryTypeNil sets the value for RecoveryType to be an explicit nil

### UnsetRecoveryType
`func (o *McmObjectRecoverActivityParams) UnsetRecoveryType()`

UnsetRecoveryType ensures that no value is present for RecoveryType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


