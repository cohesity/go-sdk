# CommonObjectProtectParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyId** | Pointer to **NullableString** | Specifies the unique id of the Protection Policy. The Policy settings will be attached with every object and will be used in backup. | [optional] 
**PolicyConfig** | Pointer to [**PolicyConfig**](PolicyConfig.md) |  | [optional] 
**StorageDomainId** | Pointer to **NullableInt64** | Specifies the Storage Domain (View Box) ID where the object backup will be taken. This is not required if Cloud archive direct is benig used. | [optional] 
**StartTime** | Pointer to [**TimeOfDay**](TimeOfDay.md) |  | [optional] 
**Priority** | Pointer to **NullableString** | Specifies the priority for the objects backup. | [optional] 
**Sla** | Pointer to [**[]SlaRule**](SlaRule.md) | Specifies the SLA parameters for list of objects. | [optional] 
**QosPolicy** | Pointer to **NullableString** | Specifies whether object backup will be written to HDD or SSD. | [optional] 
**AbortInBlackouts** | Pointer to **NullableBool** | Specifies whether currently executing object backup should abort if a blackout period specified by a policy starts. Available only if the selected policy has at least one blackout period. Default value is false. | [optional] 
**EndTimeUsecs** | Pointer to **NullableInt64** | Specifies the end time in micro seconds for this Protection Group. If this is not specified, the Protection Group won&#39;t be ended. | [optional] 

## Methods

### NewCommonObjectProtectParams

`func NewCommonObjectProtectParams() *CommonObjectProtectParams`

NewCommonObjectProtectParams instantiates a new CommonObjectProtectParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonObjectProtectParamsWithDefaults

`func NewCommonObjectProtectParamsWithDefaults() *CommonObjectProtectParams`

NewCommonObjectProtectParamsWithDefaults instantiates a new CommonObjectProtectParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyId

`func (o *CommonObjectProtectParams) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *CommonObjectProtectParams) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *CommonObjectProtectParams) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *CommonObjectProtectParams) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### SetPolicyIdNil

`func (o *CommonObjectProtectParams) SetPolicyIdNil(b bool)`

 SetPolicyIdNil sets the value for PolicyId to be an explicit nil

### UnsetPolicyId
`func (o *CommonObjectProtectParams) UnsetPolicyId()`

UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
### GetPolicyConfig

`func (o *CommonObjectProtectParams) GetPolicyConfig() PolicyConfig`

GetPolicyConfig returns the PolicyConfig field if non-nil, zero value otherwise.

### GetPolicyConfigOk

`func (o *CommonObjectProtectParams) GetPolicyConfigOk() (*PolicyConfig, bool)`

GetPolicyConfigOk returns a tuple with the PolicyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyConfig

`func (o *CommonObjectProtectParams) SetPolicyConfig(v PolicyConfig)`

SetPolicyConfig sets PolicyConfig field to given value.

### HasPolicyConfig

`func (o *CommonObjectProtectParams) HasPolicyConfig() bool`

HasPolicyConfig returns a boolean if a field has been set.

### GetStorageDomainId

`func (o *CommonObjectProtectParams) GetStorageDomainId() int64`

GetStorageDomainId returns the StorageDomainId field if non-nil, zero value otherwise.

### GetStorageDomainIdOk

`func (o *CommonObjectProtectParams) GetStorageDomainIdOk() (*int64, bool)`

GetStorageDomainIdOk returns a tuple with the StorageDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainId

`func (o *CommonObjectProtectParams) SetStorageDomainId(v int64)`

SetStorageDomainId sets StorageDomainId field to given value.

### HasStorageDomainId

`func (o *CommonObjectProtectParams) HasStorageDomainId() bool`

HasStorageDomainId returns a boolean if a field has been set.

### SetStorageDomainIdNil

`func (o *CommonObjectProtectParams) SetStorageDomainIdNil(b bool)`

 SetStorageDomainIdNil sets the value for StorageDomainId to be an explicit nil

### UnsetStorageDomainId
`func (o *CommonObjectProtectParams) UnsetStorageDomainId()`

UnsetStorageDomainId ensures that no value is present for StorageDomainId, not even an explicit nil
### GetStartTime

`func (o *CommonObjectProtectParams) GetStartTime() TimeOfDay`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *CommonObjectProtectParams) GetStartTimeOk() (*TimeOfDay, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *CommonObjectProtectParams) SetStartTime(v TimeOfDay)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *CommonObjectProtectParams) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetPriority

`func (o *CommonObjectProtectParams) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CommonObjectProtectParams) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CommonObjectProtectParams) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CommonObjectProtectParams) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *CommonObjectProtectParams) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *CommonObjectProtectParams) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetSla

`func (o *CommonObjectProtectParams) GetSla() []SlaRule`

GetSla returns the Sla field if non-nil, zero value otherwise.

### GetSlaOk

`func (o *CommonObjectProtectParams) GetSlaOk() (*[]SlaRule, bool)`

GetSlaOk returns a tuple with the Sla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSla

`func (o *CommonObjectProtectParams) SetSla(v []SlaRule)`

SetSla sets Sla field to given value.

### HasSla

`func (o *CommonObjectProtectParams) HasSla() bool`

HasSla returns a boolean if a field has been set.

### SetSlaNil

`func (o *CommonObjectProtectParams) SetSlaNil(b bool)`

 SetSlaNil sets the value for Sla to be an explicit nil

### UnsetSla
`func (o *CommonObjectProtectParams) UnsetSla()`

UnsetSla ensures that no value is present for Sla, not even an explicit nil
### GetQosPolicy

`func (o *CommonObjectProtectParams) GetQosPolicy() string`

GetQosPolicy returns the QosPolicy field if non-nil, zero value otherwise.

### GetQosPolicyOk

`func (o *CommonObjectProtectParams) GetQosPolicyOk() (*string, bool)`

GetQosPolicyOk returns a tuple with the QosPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosPolicy

`func (o *CommonObjectProtectParams) SetQosPolicy(v string)`

SetQosPolicy sets QosPolicy field to given value.

### HasQosPolicy

`func (o *CommonObjectProtectParams) HasQosPolicy() bool`

HasQosPolicy returns a boolean if a field has been set.

### SetQosPolicyNil

`func (o *CommonObjectProtectParams) SetQosPolicyNil(b bool)`

 SetQosPolicyNil sets the value for QosPolicy to be an explicit nil

### UnsetQosPolicy
`func (o *CommonObjectProtectParams) UnsetQosPolicy()`

UnsetQosPolicy ensures that no value is present for QosPolicy, not even an explicit nil
### GetAbortInBlackouts

`func (o *CommonObjectProtectParams) GetAbortInBlackouts() bool`

GetAbortInBlackouts returns the AbortInBlackouts field if non-nil, zero value otherwise.

### GetAbortInBlackoutsOk

`func (o *CommonObjectProtectParams) GetAbortInBlackoutsOk() (*bool, bool)`

GetAbortInBlackoutsOk returns a tuple with the AbortInBlackouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbortInBlackouts

`func (o *CommonObjectProtectParams) SetAbortInBlackouts(v bool)`

SetAbortInBlackouts sets AbortInBlackouts field to given value.

### HasAbortInBlackouts

`func (o *CommonObjectProtectParams) HasAbortInBlackouts() bool`

HasAbortInBlackouts returns a boolean if a field has been set.

### SetAbortInBlackoutsNil

`func (o *CommonObjectProtectParams) SetAbortInBlackoutsNil(b bool)`

 SetAbortInBlackoutsNil sets the value for AbortInBlackouts to be an explicit nil

### UnsetAbortInBlackouts
`func (o *CommonObjectProtectParams) UnsetAbortInBlackouts()`

UnsetAbortInBlackouts ensures that no value is present for AbortInBlackouts, not even an explicit nil
### GetEndTimeUsecs

`func (o *CommonObjectProtectParams) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *CommonObjectProtectParams) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *CommonObjectProtectParams) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *CommonObjectProtectParams) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *CommonObjectProtectParams) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *CommonObjectProtectParams) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


