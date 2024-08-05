# PatchOperationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InProgress** | Pointer to **NullableBool** | Specifies whether a operation is in progress now. | [optional] 
**Operation** | Pointer to **NullableString** | Specifies the patch operation. It is either apply or revert patch operation. | [optional] 
**OperationMessage** | Pointer to **NullableString** | Specifies a message about the patch operation. | [optional] 
**Percentage** | Pointer to **NullableInt64** | Specifies the percentage of completion of the current patch operation in progress or the last patch operation completed. | [optional] 
**ServicesProgress** | Pointer to [**[]ServiceUnitProgress**](ServiceUnitProgress.md) | Specifies the details of patch operation services at each patch level. | [optional] 
**TimeRemainingSeconds** | Pointer to **NullableInt64** | Specifies the time remaining to complete the patch operation. | [optional] 
**TimeTakenSeconds** | Pointer to **NullableInt64** | Specifies the time taken so far to complete the patch operation. | [optional] 

## Methods

### NewPatchOperationStatus

`func NewPatchOperationStatus() *PatchOperationStatus`

NewPatchOperationStatus instantiates a new PatchOperationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchOperationStatusWithDefaults

`func NewPatchOperationStatusWithDefaults() *PatchOperationStatus`

NewPatchOperationStatusWithDefaults instantiates a new PatchOperationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInProgress

`func (o *PatchOperationStatus) GetInProgress() bool`

GetInProgress returns the InProgress field if non-nil, zero value otherwise.

### GetInProgressOk

`func (o *PatchOperationStatus) GetInProgressOk() (*bool, bool)`

GetInProgressOk returns a tuple with the InProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInProgress

`func (o *PatchOperationStatus) SetInProgress(v bool)`

SetInProgress sets InProgress field to given value.

### HasInProgress

`func (o *PatchOperationStatus) HasInProgress() bool`

HasInProgress returns a boolean if a field has been set.

### SetInProgressNil

`func (o *PatchOperationStatus) SetInProgressNil(b bool)`

 SetInProgressNil sets the value for InProgress to be an explicit nil

### UnsetInProgress
`func (o *PatchOperationStatus) UnsetInProgress()`

UnsetInProgress ensures that no value is present for InProgress, not even an explicit nil
### GetOperation

`func (o *PatchOperationStatus) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *PatchOperationStatus) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *PatchOperationStatus) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *PatchOperationStatus) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### SetOperationNil

`func (o *PatchOperationStatus) SetOperationNil(b bool)`

 SetOperationNil sets the value for Operation to be an explicit nil

### UnsetOperation
`func (o *PatchOperationStatus) UnsetOperation()`

UnsetOperation ensures that no value is present for Operation, not even an explicit nil
### GetOperationMessage

`func (o *PatchOperationStatus) GetOperationMessage() string`

GetOperationMessage returns the OperationMessage field if non-nil, zero value otherwise.

### GetOperationMessageOk

`func (o *PatchOperationStatus) GetOperationMessageOk() (*string, bool)`

GetOperationMessageOk returns a tuple with the OperationMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationMessage

`func (o *PatchOperationStatus) SetOperationMessage(v string)`

SetOperationMessage sets OperationMessage field to given value.

### HasOperationMessage

`func (o *PatchOperationStatus) HasOperationMessage() bool`

HasOperationMessage returns a boolean if a field has been set.

### SetOperationMessageNil

`func (o *PatchOperationStatus) SetOperationMessageNil(b bool)`

 SetOperationMessageNil sets the value for OperationMessage to be an explicit nil

### UnsetOperationMessage
`func (o *PatchOperationStatus) UnsetOperationMessage()`

UnsetOperationMessage ensures that no value is present for OperationMessage, not even an explicit nil
### GetPercentage

`func (o *PatchOperationStatus) GetPercentage() int64`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *PatchOperationStatus) GetPercentageOk() (*int64, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *PatchOperationStatus) SetPercentage(v int64)`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *PatchOperationStatus) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.

### SetPercentageNil

`func (o *PatchOperationStatus) SetPercentageNil(b bool)`

 SetPercentageNil sets the value for Percentage to be an explicit nil

### UnsetPercentage
`func (o *PatchOperationStatus) UnsetPercentage()`

UnsetPercentage ensures that no value is present for Percentage, not even an explicit nil
### GetServicesProgress

`func (o *PatchOperationStatus) GetServicesProgress() []ServiceUnitProgress`

GetServicesProgress returns the ServicesProgress field if non-nil, zero value otherwise.

### GetServicesProgressOk

`func (o *PatchOperationStatus) GetServicesProgressOk() (*[]ServiceUnitProgress, bool)`

GetServicesProgressOk returns a tuple with the ServicesProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicesProgress

`func (o *PatchOperationStatus) SetServicesProgress(v []ServiceUnitProgress)`

SetServicesProgress sets ServicesProgress field to given value.

### HasServicesProgress

`func (o *PatchOperationStatus) HasServicesProgress() bool`

HasServicesProgress returns a boolean if a field has been set.

### SetServicesProgressNil

`func (o *PatchOperationStatus) SetServicesProgressNil(b bool)`

 SetServicesProgressNil sets the value for ServicesProgress to be an explicit nil

### UnsetServicesProgress
`func (o *PatchOperationStatus) UnsetServicesProgress()`

UnsetServicesProgress ensures that no value is present for ServicesProgress, not even an explicit nil
### GetTimeRemainingSeconds

`func (o *PatchOperationStatus) GetTimeRemainingSeconds() int64`

GetTimeRemainingSeconds returns the TimeRemainingSeconds field if non-nil, zero value otherwise.

### GetTimeRemainingSecondsOk

`func (o *PatchOperationStatus) GetTimeRemainingSecondsOk() (*int64, bool)`

GetTimeRemainingSecondsOk returns a tuple with the TimeRemainingSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRemainingSeconds

`func (o *PatchOperationStatus) SetTimeRemainingSeconds(v int64)`

SetTimeRemainingSeconds sets TimeRemainingSeconds field to given value.

### HasTimeRemainingSeconds

`func (o *PatchOperationStatus) HasTimeRemainingSeconds() bool`

HasTimeRemainingSeconds returns a boolean if a field has been set.

### SetTimeRemainingSecondsNil

`func (o *PatchOperationStatus) SetTimeRemainingSecondsNil(b bool)`

 SetTimeRemainingSecondsNil sets the value for TimeRemainingSeconds to be an explicit nil

### UnsetTimeRemainingSeconds
`func (o *PatchOperationStatus) UnsetTimeRemainingSeconds()`

UnsetTimeRemainingSeconds ensures that no value is present for TimeRemainingSeconds, not even an explicit nil
### GetTimeTakenSeconds

`func (o *PatchOperationStatus) GetTimeTakenSeconds() int64`

GetTimeTakenSeconds returns the TimeTakenSeconds field if non-nil, zero value otherwise.

### GetTimeTakenSecondsOk

`func (o *PatchOperationStatus) GetTimeTakenSecondsOk() (*int64, bool)`

GetTimeTakenSecondsOk returns a tuple with the TimeTakenSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeTakenSeconds

`func (o *PatchOperationStatus) SetTimeTakenSeconds(v int64)`

SetTimeTakenSeconds sets TimeTakenSeconds field to given value.

### HasTimeTakenSeconds

`func (o *PatchOperationStatus) HasTimeTakenSeconds() bool`

HasTimeTakenSeconds returns a boolean if a field has been set.

### SetTimeTakenSecondsNil

`func (o *PatchOperationStatus) SetTimeTakenSecondsNil(b bool)`

 SetTimeTakenSecondsNil sets the value for TimeTakenSeconds to be an explicit nil

### UnsetTimeTakenSeconds
`func (o *PatchOperationStatus) UnsetTimeTakenSeconds()`

UnsetTimeTakenSeconds ensures that no value is present for TimeTakenSeconds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


