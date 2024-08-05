# CommonProtectionGroupRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbortInBlackouts** | Pointer to **NullableBool** | Specifies whether currently executing jobs should abort if a blackout period specified by a policy starts. Available only if the selected policy has at least one blackout period. Default value is false. This field should not be set to true if &#39;pauseInBlackouts&#39; is set to true. | [optional] 
**AdvancedConfigs** | Pointer to [**[]KeyValuePair**](KeyValuePair.md) | Specifies the advanced configuration for a protection job. | [optional] 
**AlertPolicy** | Pointer to [**ProtectionGroupAlertingPolicy**](ProtectionGroupAlertingPolicy.md) |  | [optional] 
**Description** | Pointer to **NullableString** | Specifies a description of the Protection Group. | [optional] 
**EndTimeUsecs** | Pointer to **NullableInt64** | Specifies the end time in micro seconds for this Protection Group. If this is not specified, the Protection Group won&#39;t be ended. | [optional] 
**Environment** | **NullableString** | Specifies the environment type of the Protection Group. | 
**IsPaused** | Pointer to **NullableBool** | Specifies if the the Protection Group is paused. New runs are not scheduled for the paused Protection Groups. Active run if any is not impacted. | [optional] 
**LastModifiedTimestampUsecs** | Pointer to **NullableInt64** | Specifies the last time this protection group was updated. If this is passed into a PUT request, then the backend will validate that the timestamp passed in matches the time that the protection group was actually last modified. If the two timestamps do not match, then the request will be rejected with a stale error. | [optional] 
**Name** | **NullableString** | Specifies the name of the Protection Group. | 
**PauseInBlackouts** | Pointer to **NullableBool** | Specifies whether currently executing jobs should be paused if a blackout period specified by a policy starts. Available only if the selected policy has at least one blackout period. Default value is false. This field should not be set to true if &#39;abortInBlackouts&#39; is sent as true. | [optional] 
**PolicyId** | **NullableString** | Specifies the unique id of the Protection Policy associated with the Protection Group. The Policy provides retry settings Protection Schedules, Priority, SLA, etc. | 
**Priority** | Pointer to **NullableString** | Specifies the priority of the Protection Group. | [optional] 
**QosPolicy** | Pointer to **NullableString** | Specifies whether the Protection Group will be written to HDD or SSD. | [optional] 
**Sla** | Pointer to [**[]SlaRule**](SlaRule.md) | Specifies the SLA parameters for this Protection Group. | [optional] 
**StartTime** | Pointer to [**TimeOfDay**](TimeOfDay.md) |  | [optional] 
**StorageDomainId** | Pointer to **NullableInt64** | Specifies the Storage Domain (View Box) ID where this Protection Group writes data. | [optional] 

## Methods

### NewCommonProtectionGroupRequestParams

`func NewCommonProtectionGroupRequestParams(environment NullableString, name NullableString, policyId NullableString, ) *CommonProtectionGroupRequestParams`

NewCommonProtectionGroupRequestParams instantiates a new CommonProtectionGroupRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonProtectionGroupRequestParamsWithDefaults

`func NewCommonProtectionGroupRequestParamsWithDefaults() *CommonProtectionGroupRequestParams`

NewCommonProtectionGroupRequestParamsWithDefaults instantiates a new CommonProtectionGroupRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbortInBlackouts

`func (o *CommonProtectionGroupRequestParams) GetAbortInBlackouts() bool`

GetAbortInBlackouts returns the AbortInBlackouts field if non-nil, zero value otherwise.

### GetAbortInBlackoutsOk

`func (o *CommonProtectionGroupRequestParams) GetAbortInBlackoutsOk() (*bool, bool)`

GetAbortInBlackoutsOk returns a tuple with the AbortInBlackouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbortInBlackouts

`func (o *CommonProtectionGroupRequestParams) SetAbortInBlackouts(v bool)`

SetAbortInBlackouts sets AbortInBlackouts field to given value.

### HasAbortInBlackouts

`func (o *CommonProtectionGroupRequestParams) HasAbortInBlackouts() bool`

HasAbortInBlackouts returns a boolean if a field has been set.

### SetAbortInBlackoutsNil

`func (o *CommonProtectionGroupRequestParams) SetAbortInBlackoutsNil(b bool)`

 SetAbortInBlackoutsNil sets the value for AbortInBlackouts to be an explicit nil

### UnsetAbortInBlackouts
`func (o *CommonProtectionGroupRequestParams) UnsetAbortInBlackouts()`

UnsetAbortInBlackouts ensures that no value is present for AbortInBlackouts, not even an explicit nil
### GetAdvancedConfigs

`func (o *CommonProtectionGroupRequestParams) GetAdvancedConfigs() []KeyValuePair`

GetAdvancedConfigs returns the AdvancedConfigs field if non-nil, zero value otherwise.

### GetAdvancedConfigsOk

`func (o *CommonProtectionGroupRequestParams) GetAdvancedConfigsOk() (*[]KeyValuePair, bool)`

GetAdvancedConfigsOk returns a tuple with the AdvancedConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedConfigs

`func (o *CommonProtectionGroupRequestParams) SetAdvancedConfigs(v []KeyValuePair)`

SetAdvancedConfigs sets AdvancedConfigs field to given value.

### HasAdvancedConfigs

`func (o *CommonProtectionGroupRequestParams) HasAdvancedConfigs() bool`

HasAdvancedConfigs returns a boolean if a field has been set.

### SetAdvancedConfigsNil

`func (o *CommonProtectionGroupRequestParams) SetAdvancedConfigsNil(b bool)`

 SetAdvancedConfigsNil sets the value for AdvancedConfigs to be an explicit nil

### UnsetAdvancedConfigs
`func (o *CommonProtectionGroupRequestParams) UnsetAdvancedConfigs()`

UnsetAdvancedConfigs ensures that no value is present for AdvancedConfigs, not even an explicit nil
### GetAlertPolicy

`func (o *CommonProtectionGroupRequestParams) GetAlertPolicy() ProtectionGroupAlertingPolicy`

GetAlertPolicy returns the AlertPolicy field if non-nil, zero value otherwise.

### GetAlertPolicyOk

`func (o *CommonProtectionGroupRequestParams) GetAlertPolicyOk() (*ProtectionGroupAlertingPolicy, bool)`

GetAlertPolicyOk returns a tuple with the AlertPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertPolicy

`func (o *CommonProtectionGroupRequestParams) SetAlertPolicy(v ProtectionGroupAlertingPolicy)`

SetAlertPolicy sets AlertPolicy field to given value.

### HasAlertPolicy

`func (o *CommonProtectionGroupRequestParams) HasAlertPolicy() bool`

HasAlertPolicy returns a boolean if a field has been set.

### GetDescription

`func (o *CommonProtectionGroupRequestParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CommonProtectionGroupRequestParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CommonProtectionGroupRequestParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CommonProtectionGroupRequestParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CommonProtectionGroupRequestParams) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CommonProtectionGroupRequestParams) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEndTimeUsecs

`func (o *CommonProtectionGroupRequestParams) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *CommonProtectionGroupRequestParams) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *CommonProtectionGroupRequestParams) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *CommonProtectionGroupRequestParams) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *CommonProtectionGroupRequestParams) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *CommonProtectionGroupRequestParams) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetEnvironment

`func (o *CommonProtectionGroupRequestParams) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CommonProtectionGroupRequestParams) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CommonProtectionGroupRequestParams) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### SetEnvironmentNil

`func (o *CommonProtectionGroupRequestParams) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *CommonProtectionGroupRequestParams) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetIsPaused

`func (o *CommonProtectionGroupRequestParams) GetIsPaused() bool`

GetIsPaused returns the IsPaused field if non-nil, zero value otherwise.

### GetIsPausedOk

`func (o *CommonProtectionGroupRequestParams) GetIsPausedOk() (*bool, bool)`

GetIsPausedOk returns a tuple with the IsPaused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPaused

`func (o *CommonProtectionGroupRequestParams) SetIsPaused(v bool)`

SetIsPaused sets IsPaused field to given value.

### HasIsPaused

`func (o *CommonProtectionGroupRequestParams) HasIsPaused() bool`

HasIsPaused returns a boolean if a field has been set.

### SetIsPausedNil

`func (o *CommonProtectionGroupRequestParams) SetIsPausedNil(b bool)`

 SetIsPausedNil sets the value for IsPaused to be an explicit nil

### UnsetIsPaused
`func (o *CommonProtectionGroupRequestParams) UnsetIsPaused()`

UnsetIsPaused ensures that no value is present for IsPaused, not even an explicit nil
### GetLastModifiedTimestampUsecs

`func (o *CommonProtectionGroupRequestParams) GetLastModifiedTimestampUsecs() int64`

GetLastModifiedTimestampUsecs returns the LastModifiedTimestampUsecs field if non-nil, zero value otherwise.

### GetLastModifiedTimestampUsecsOk

`func (o *CommonProtectionGroupRequestParams) GetLastModifiedTimestampUsecsOk() (*int64, bool)`

GetLastModifiedTimestampUsecsOk returns a tuple with the LastModifiedTimestampUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTimestampUsecs

`func (o *CommonProtectionGroupRequestParams) SetLastModifiedTimestampUsecs(v int64)`

SetLastModifiedTimestampUsecs sets LastModifiedTimestampUsecs field to given value.

### HasLastModifiedTimestampUsecs

`func (o *CommonProtectionGroupRequestParams) HasLastModifiedTimestampUsecs() bool`

HasLastModifiedTimestampUsecs returns a boolean if a field has been set.

### SetLastModifiedTimestampUsecsNil

`func (o *CommonProtectionGroupRequestParams) SetLastModifiedTimestampUsecsNil(b bool)`

 SetLastModifiedTimestampUsecsNil sets the value for LastModifiedTimestampUsecs to be an explicit nil

### UnsetLastModifiedTimestampUsecs
`func (o *CommonProtectionGroupRequestParams) UnsetLastModifiedTimestampUsecs()`

UnsetLastModifiedTimestampUsecs ensures that no value is present for LastModifiedTimestampUsecs, not even an explicit nil
### GetName

`func (o *CommonProtectionGroupRequestParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CommonProtectionGroupRequestParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CommonProtectionGroupRequestParams) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CommonProtectionGroupRequestParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CommonProtectionGroupRequestParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPauseInBlackouts

`func (o *CommonProtectionGroupRequestParams) GetPauseInBlackouts() bool`

GetPauseInBlackouts returns the PauseInBlackouts field if non-nil, zero value otherwise.

### GetPauseInBlackoutsOk

`func (o *CommonProtectionGroupRequestParams) GetPauseInBlackoutsOk() (*bool, bool)`

GetPauseInBlackoutsOk returns a tuple with the PauseInBlackouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseInBlackouts

`func (o *CommonProtectionGroupRequestParams) SetPauseInBlackouts(v bool)`

SetPauseInBlackouts sets PauseInBlackouts field to given value.

### HasPauseInBlackouts

`func (o *CommonProtectionGroupRequestParams) HasPauseInBlackouts() bool`

HasPauseInBlackouts returns a boolean if a field has been set.

### SetPauseInBlackoutsNil

`func (o *CommonProtectionGroupRequestParams) SetPauseInBlackoutsNil(b bool)`

 SetPauseInBlackoutsNil sets the value for PauseInBlackouts to be an explicit nil

### UnsetPauseInBlackouts
`func (o *CommonProtectionGroupRequestParams) UnsetPauseInBlackouts()`

UnsetPauseInBlackouts ensures that no value is present for PauseInBlackouts, not even an explicit nil
### GetPolicyId

`func (o *CommonProtectionGroupRequestParams) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *CommonProtectionGroupRequestParams) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *CommonProtectionGroupRequestParams) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.


### SetPolicyIdNil

`func (o *CommonProtectionGroupRequestParams) SetPolicyIdNil(b bool)`

 SetPolicyIdNil sets the value for PolicyId to be an explicit nil

### UnsetPolicyId
`func (o *CommonProtectionGroupRequestParams) UnsetPolicyId()`

UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
### GetPriority

`func (o *CommonProtectionGroupRequestParams) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CommonProtectionGroupRequestParams) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CommonProtectionGroupRequestParams) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CommonProtectionGroupRequestParams) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *CommonProtectionGroupRequestParams) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *CommonProtectionGroupRequestParams) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetQosPolicy

`func (o *CommonProtectionGroupRequestParams) GetQosPolicy() string`

GetQosPolicy returns the QosPolicy field if non-nil, zero value otherwise.

### GetQosPolicyOk

`func (o *CommonProtectionGroupRequestParams) GetQosPolicyOk() (*string, bool)`

GetQosPolicyOk returns a tuple with the QosPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosPolicy

`func (o *CommonProtectionGroupRequestParams) SetQosPolicy(v string)`

SetQosPolicy sets QosPolicy field to given value.

### HasQosPolicy

`func (o *CommonProtectionGroupRequestParams) HasQosPolicy() bool`

HasQosPolicy returns a boolean if a field has been set.

### SetQosPolicyNil

`func (o *CommonProtectionGroupRequestParams) SetQosPolicyNil(b bool)`

 SetQosPolicyNil sets the value for QosPolicy to be an explicit nil

### UnsetQosPolicy
`func (o *CommonProtectionGroupRequestParams) UnsetQosPolicy()`

UnsetQosPolicy ensures that no value is present for QosPolicy, not even an explicit nil
### GetSla

`func (o *CommonProtectionGroupRequestParams) GetSla() []SlaRule`

GetSla returns the Sla field if non-nil, zero value otherwise.

### GetSlaOk

`func (o *CommonProtectionGroupRequestParams) GetSlaOk() (*[]SlaRule, bool)`

GetSlaOk returns a tuple with the Sla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSla

`func (o *CommonProtectionGroupRequestParams) SetSla(v []SlaRule)`

SetSla sets Sla field to given value.

### HasSla

`func (o *CommonProtectionGroupRequestParams) HasSla() bool`

HasSla returns a boolean if a field has been set.

### SetSlaNil

`func (o *CommonProtectionGroupRequestParams) SetSlaNil(b bool)`

 SetSlaNil sets the value for Sla to be an explicit nil

### UnsetSla
`func (o *CommonProtectionGroupRequestParams) UnsetSla()`

UnsetSla ensures that no value is present for Sla, not even an explicit nil
### GetStartTime

`func (o *CommonProtectionGroupRequestParams) GetStartTime() TimeOfDay`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *CommonProtectionGroupRequestParams) GetStartTimeOk() (*TimeOfDay, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *CommonProtectionGroupRequestParams) SetStartTime(v TimeOfDay)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *CommonProtectionGroupRequestParams) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStorageDomainId

`func (o *CommonProtectionGroupRequestParams) GetStorageDomainId() int64`

GetStorageDomainId returns the StorageDomainId field if non-nil, zero value otherwise.

### GetStorageDomainIdOk

`func (o *CommonProtectionGroupRequestParams) GetStorageDomainIdOk() (*int64, bool)`

GetStorageDomainIdOk returns a tuple with the StorageDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainId

`func (o *CommonProtectionGroupRequestParams) SetStorageDomainId(v int64)`

SetStorageDomainId sets StorageDomainId field to given value.

### HasStorageDomainId

`func (o *CommonProtectionGroupRequestParams) HasStorageDomainId() bool`

HasStorageDomainId returns a boolean if a field has been set.

### SetStorageDomainIdNil

`func (o *CommonProtectionGroupRequestParams) SetStorageDomainIdNil(b bool)`

 SetStorageDomainIdNil sets the value for StorageDomainId to be an explicit nil

### UnsetStorageDomainId
`func (o *CommonProtectionGroupRequestParams) UnsetStorageDomainId()`

UnsetStorageDomainId ensures that no value is present for StorageDomainId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


