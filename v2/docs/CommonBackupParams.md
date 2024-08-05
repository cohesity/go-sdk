# CommonBackupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbortInBlackouts** | Pointer to **NullableBool** | Specifies whether currently executing object backup should abort if a blackout period specified by a policy starts. Available only if the selected policy has at least one blackout period. Default value is false. | [optional] 
**EndTimeUsecs** | Pointer to **NullableInt64** | Specifies the end time in micro seconds for this Protection Group. If this is not specified, the Protection Group won&#39;t be ended. | [optional] 
**PolicyConfig** | Pointer to [**PolicyConfig**](PolicyConfig.md) |  | [optional] 
**PolicyId** | Pointer to **NullableString** | Specifies the unique id of the Protection Policy. The Policy settings will be attached with every object and will be used in backup. | [optional] 
**Priority** | Pointer to **NullableString** | Specifies the priority for the objects backup. | [optional] 
**QosPolicy** | Pointer to **NullableString** | Specifies whether object backup will be written to HDD or SSD. | [optional] 
**SkipRigelForBackup** | Pointer to **NullableBool** | Specifies whether to skip Rigel for backup or not. | [optional] 
**Sla** | Pointer to [**[]SlaRule**](SlaRule.md) | Specifies the SLA parameters for list of objects. | [optional] 
**StartTime** | Pointer to [**TimeOfDay**](TimeOfDay.md) |  | [optional] 
**StorageDomainId** | Pointer to **NullableInt64** | Specifies the Storage Domain (View Box) ID where the object backup will be taken. This is not required if Cloud archive direct is benig used. | [optional] 

## Methods

### NewCommonBackupParams

`func NewCommonBackupParams() *CommonBackupParams`

NewCommonBackupParams instantiates a new CommonBackupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonBackupParamsWithDefaults

`func NewCommonBackupParamsWithDefaults() *CommonBackupParams`

NewCommonBackupParamsWithDefaults instantiates a new CommonBackupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbortInBlackouts

`func (o *CommonBackupParams) GetAbortInBlackouts() bool`

GetAbortInBlackouts returns the AbortInBlackouts field if non-nil, zero value otherwise.

### GetAbortInBlackoutsOk

`func (o *CommonBackupParams) GetAbortInBlackoutsOk() (*bool, bool)`

GetAbortInBlackoutsOk returns a tuple with the AbortInBlackouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbortInBlackouts

`func (o *CommonBackupParams) SetAbortInBlackouts(v bool)`

SetAbortInBlackouts sets AbortInBlackouts field to given value.

### HasAbortInBlackouts

`func (o *CommonBackupParams) HasAbortInBlackouts() bool`

HasAbortInBlackouts returns a boolean if a field has been set.

### SetAbortInBlackoutsNil

`func (o *CommonBackupParams) SetAbortInBlackoutsNil(b bool)`

 SetAbortInBlackoutsNil sets the value for AbortInBlackouts to be an explicit nil

### UnsetAbortInBlackouts
`func (o *CommonBackupParams) UnsetAbortInBlackouts()`

UnsetAbortInBlackouts ensures that no value is present for AbortInBlackouts, not even an explicit nil
### GetEndTimeUsecs

`func (o *CommonBackupParams) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *CommonBackupParams) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *CommonBackupParams) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *CommonBackupParams) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *CommonBackupParams) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *CommonBackupParams) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetPolicyConfig

`func (o *CommonBackupParams) GetPolicyConfig() PolicyConfig`

GetPolicyConfig returns the PolicyConfig field if non-nil, zero value otherwise.

### GetPolicyConfigOk

`func (o *CommonBackupParams) GetPolicyConfigOk() (*PolicyConfig, bool)`

GetPolicyConfigOk returns a tuple with the PolicyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyConfig

`func (o *CommonBackupParams) SetPolicyConfig(v PolicyConfig)`

SetPolicyConfig sets PolicyConfig field to given value.

### HasPolicyConfig

`func (o *CommonBackupParams) HasPolicyConfig() bool`

HasPolicyConfig returns a boolean if a field has been set.

### GetPolicyId

`func (o *CommonBackupParams) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *CommonBackupParams) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *CommonBackupParams) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *CommonBackupParams) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### SetPolicyIdNil

`func (o *CommonBackupParams) SetPolicyIdNil(b bool)`

 SetPolicyIdNil sets the value for PolicyId to be an explicit nil

### UnsetPolicyId
`func (o *CommonBackupParams) UnsetPolicyId()`

UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
### GetPriority

`func (o *CommonBackupParams) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CommonBackupParams) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CommonBackupParams) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CommonBackupParams) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *CommonBackupParams) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *CommonBackupParams) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetQosPolicy

`func (o *CommonBackupParams) GetQosPolicy() string`

GetQosPolicy returns the QosPolicy field if non-nil, zero value otherwise.

### GetQosPolicyOk

`func (o *CommonBackupParams) GetQosPolicyOk() (*string, bool)`

GetQosPolicyOk returns a tuple with the QosPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosPolicy

`func (o *CommonBackupParams) SetQosPolicy(v string)`

SetQosPolicy sets QosPolicy field to given value.

### HasQosPolicy

`func (o *CommonBackupParams) HasQosPolicy() bool`

HasQosPolicy returns a boolean if a field has been set.

### SetQosPolicyNil

`func (o *CommonBackupParams) SetQosPolicyNil(b bool)`

 SetQosPolicyNil sets the value for QosPolicy to be an explicit nil

### UnsetQosPolicy
`func (o *CommonBackupParams) UnsetQosPolicy()`

UnsetQosPolicy ensures that no value is present for QosPolicy, not even an explicit nil
### GetSkipRigelForBackup

`func (o *CommonBackupParams) GetSkipRigelForBackup() bool`

GetSkipRigelForBackup returns the SkipRigelForBackup field if non-nil, zero value otherwise.

### GetSkipRigelForBackupOk

`func (o *CommonBackupParams) GetSkipRigelForBackupOk() (*bool, bool)`

GetSkipRigelForBackupOk returns a tuple with the SkipRigelForBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipRigelForBackup

`func (o *CommonBackupParams) SetSkipRigelForBackup(v bool)`

SetSkipRigelForBackup sets SkipRigelForBackup field to given value.

### HasSkipRigelForBackup

`func (o *CommonBackupParams) HasSkipRigelForBackup() bool`

HasSkipRigelForBackup returns a boolean if a field has been set.

### SetSkipRigelForBackupNil

`func (o *CommonBackupParams) SetSkipRigelForBackupNil(b bool)`

 SetSkipRigelForBackupNil sets the value for SkipRigelForBackup to be an explicit nil

### UnsetSkipRigelForBackup
`func (o *CommonBackupParams) UnsetSkipRigelForBackup()`

UnsetSkipRigelForBackup ensures that no value is present for SkipRigelForBackup, not even an explicit nil
### GetSla

`func (o *CommonBackupParams) GetSla() []SlaRule`

GetSla returns the Sla field if non-nil, zero value otherwise.

### GetSlaOk

`func (o *CommonBackupParams) GetSlaOk() (*[]SlaRule, bool)`

GetSlaOk returns a tuple with the Sla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSla

`func (o *CommonBackupParams) SetSla(v []SlaRule)`

SetSla sets Sla field to given value.

### HasSla

`func (o *CommonBackupParams) HasSla() bool`

HasSla returns a boolean if a field has been set.

### SetSlaNil

`func (o *CommonBackupParams) SetSlaNil(b bool)`

 SetSlaNil sets the value for Sla to be an explicit nil

### UnsetSla
`func (o *CommonBackupParams) UnsetSla()`

UnsetSla ensures that no value is present for Sla, not even an explicit nil
### GetStartTime

`func (o *CommonBackupParams) GetStartTime() TimeOfDay`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *CommonBackupParams) GetStartTimeOk() (*TimeOfDay, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *CommonBackupParams) SetStartTime(v TimeOfDay)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *CommonBackupParams) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStorageDomainId

`func (o *CommonBackupParams) GetStorageDomainId() int64`

GetStorageDomainId returns the StorageDomainId field if non-nil, zero value otherwise.

### GetStorageDomainIdOk

`func (o *CommonBackupParams) GetStorageDomainIdOk() (*int64, bool)`

GetStorageDomainIdOk returns a tuple with the StorageDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainId

`func (o *CommonBackupParams) SetStorageDomainId(v int64)`

SetStorageDomainId sets StorageDomainId field to given value.

### HasStorageDomainId

`func (o *CommonBackupParams) HasStorageDomainId() bool`

HasStorageDomainId returns a boolean if a field has been set.

### SetStorageDomainIdNil

`func (o *CommonBackupParams) SetStorageDomainIdNil(b bool)`

 SetStorageDomainIdNil sets the value for StorageDomainId to be an explicit nil

### UnsetStorageDomainId
`func (o *CommonBackupParams) UnsetStorageDomainId()`

UnsetStorageDomainId ensures that no value is present for StorageDomainId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


