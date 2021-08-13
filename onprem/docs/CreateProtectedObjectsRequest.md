# CreateProtectedObjectsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | [**[]EnvSpecificObjectProtectionRequestParams**](EnvSpecificObjectProtectionRequestParams.md) | Specifies the list of objects to be protected. Multiple objects from different adapters can be provided as input. | 
**ActivateRemoteObjectProtection** | Pointer to **NullableBool** | If set to true, it will look for the remote backup of the given user and object, and activates it. Creates a new backup if the remote backup is not found. After activation, this object cannot get snapshots from remote clusters. | [optional] 
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

### NewCreateProtectedObjectsRequest

`func NewCreateProtectedObjectsRequest(objects []EnvSpecificObjectProtectionRequestParams, ) *CreateProtectedObjectsRequest`

NewCreateProtectedObjectsRequest instantiates a new CreateProtectedObjectsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProtectedObjectsRequestWithDefaults

`func NewCreateProtectedObjectsRequestWithDefaults() *CreateProtectedObjectsRequest`

NewCreateProtectedObjectsRequestWithDefaults instantiates a new CreateProtectedObjectsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjects

`func (o *CreateProtectedObjectsRequest) GetObjects() []EnvSpecificObjectProtectionRequestParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *CreateProtectedObjectsRequest) GetObjectsOk() (*[]EnvSpecificObjectProtectionRequestParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *CreateProtectedObjectsRequest) SetObjects(v []EnvSpecificObjectProtectionRequestParams)`

SetObjects sets Objects field to given value.


### SetObjectsNil

`func (o *CreateProtectedObjectsRequest) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *CreateProtectedObjectsRequest) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetActivateRemoteObjectProtection

`func (o *CreateProtectedObjectsRequest) GetActivateRemoteObjectProtection() bool`

GetActivateRemoteObjectProtection returns the ActivateRemoteObjectProtection field if non-nil, zero value otherwise.

### GetActivateRemoteObjectProtectionOk

`func (o *CreateProtectedObjectsRequest) GetActivateRemoteObjectProtectionOk() (*bool, bool)`

GetActivateRemoteObjectProtectionOk returns a tuple with the ActivateRemoteObjectProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivateRemoteObjectProtection

`func (o *CreateProtectedObjectsRequest) SetActivateRemoteObjectProtection(v bool)`

SetActivateRemoteObjectProtection sets ActivateRemoteObjectProtection field to given value.

### HasActivateRemoteObjectProtection

`func (o *CreateProtectedObjectsRequest) HasActivateRemoteObjectProtection() bool`

HasActivateRemoteObjectProtection returns a boolean if a field has been set.

### SetActivateRemoteObjectProtectionNil

`func (o *CreateProtectedObjectsRequest) SetActivateRemoteObjectProtectionNil(b bool)`

 SetActivateRemoteObjectProtectionNil sets the value for ActivateRemoteObjectProtection to be an explicit nil

### UnsetActivateRemoteObjectProtection
`func (o *CreateProtectedObjectsRequest) UnsetActivateRemoteObjectProtection()`

UnsetActivateRemoteObjectProtection ensures that no value is present for ActivateRemoteObjectProtection, not even an explicit nil
### GetPolicyId

`func (o *CreateProtectedObjectsRequest) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *CreateProtectedObjectsRequest) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *CreateProtectedObjectsRequest) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *CreateProtectedObjectsRequest) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### SetPolicyIdNil

`func (o *CreateProtectedObjectsRequest) SetPolicyIdNil(b bool)`

 SetPolicyIdNil sets the value for PolicyId to be an explicit nil

### UnsetPolicyId
`func (o *CreateProtectedObjectsRequest) UnsetPolicyId()`

UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
### GetPolicyConfig

`func (o *CreateProtectedObjectsRequest) GetPolicyConfig() PolicyConfig`

GetPolicyConfig returns the PolicyConfig field if non-nil, zero value otherwise.

### GetPolicyConfigOk

`func (o *CreateProtectedObjectsRequest) GetPolicyConfigOk() (*PolicyConfig, bool)`

GetPolicyConfigOk returns a tuple with the PolicyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyConfig

`func (o *CreateProtectedObjectsRequest) SetPolicyConfig(v PolicyConfig)`

SetPolicyConfig sets PolicyConfig field to given value.

### HasPolicyConfig

`func (o *CreateProtectedObjectsRequest) HasPolicyConfig() bool`

HasPolicyConfig returns a boolean if a field has been set.

### GetStorageDomainId

`func (o *CreateProtectedObjectsRequest) GetStorageDomainId() int64`

GetStorageDomainId returns the StorageDomainId field if non-nil, zero value otherwise.

### GetStorageDomainIdOk

`func (o *CreateProtectedObjectsRequest) GetStorageDomainIdOk() (*int64, bool)`

GetStorageDomainIdOk returns a tuple with the StorageDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainId

`func (o *CreateProtectedObjectsRequest) SetStorageDomainId(v int64)`

SetStorageDomainId sets StorageDomainId field to given value.

### HasStorageDomainId

`func (o *CreateProtectedObjectsRequest) HasStorageDomainId() bool`

HasStorageDomainId returns a boolean if a field has been set.

### SetStorageDomainIdNil

`func (o *CreateProtectedObjectsRequest) SetStorageDomainIdNil(b bool)`

 SetStorageDomainIdNil sets the value for StorageDomainId to be an explicit nil

### UnsetStorageDomainId
`func (o *CreateProtectedObjectsRequest) UnsetStorageDomainId()`

UnsetStorageDomainId ensures that no value is present for StorageDomainId, not even an explicit nil
### GetStartTime

`func (o *CreateProtectedObjectsRequest) GetStartTime() TimeOfDay`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *CreateProtectedObjectsRequest) GetStartTimeOk() (*TimeOfDay, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *CreateProtectedObjectsRequest) SetStartTime(v TimeOfDay)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *CreateProtectedObjectsRequest) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetPriority

`func (o *CreateProtectedObjectsRequest) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CreateProtectedObjectsRequest) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CreateProtectedObjectsRequest) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CreateProtectedObjectsRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *CreateProtectedObjectsRequest) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *CreateProtectedObjectsRequest) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetSla

`func (o *CreateProtectedObjectsRequest) GetSla() []SlaRule`

GetSla returns the Sla field if non-nil, zero value otherwise.

### GetSlaOk

`func (o *CreateProtectedObjectsRequest) GetSlaOk() (*[]SlaRule, bool)`

GetSlaOk returns a tuple with the Sla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSla

`func (o *CreateProtectedObjectsRequest) SetSla(v []SlaRule)`

SetSla sets Sla field to given value.

### HasSla

`func (o *CreateProtectedObjectsRequest) HasSla() bool`

HasSla returns a boolean if a field has been set.

### SetSlaNil

`func (o *CreateProtectedObjectsRequest) SetSlaNil(b bool)`

 SetSlaNil sets the value for Sla to be an explicit nil

### UnsetSla
`func (o *CreateProtectedObjectsRequest) UnsetSla()`

UnsetSla ensures that no value is present for Sla, not even an explicit nil
### GetQosPolicy

`func (o *CreateProtectedObjectsRequest) GetQosPolicy() string`

GetQosPolicy returns the QosPolicy field if non-nil, zero value otherwise.

### GetQosPolicyOk

`func (o *CreateProtectedObjectsRequest) GetQosPolicyOk() (*string, bool)`

GetQosPolicyOk returns a tuple with the QosPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosPolicy

`func (o *CreateProtectedObjectsRequest) SetQosPolicy(v string)`

SetQosPolicy sets QosPolicy field to given value.

### HasQosPolicy

`func (o *CreateProtectedObjectsRequest) HasQosPolicy() bool`

HasQosPolicy returns a boolean if a field has been set.

### SetQosPolicyNil

`func (o *CreateProtectedObjectsRequest) SetQosPolicyNil(b bool)`

 SetQosPolicyNil sets the value for QosPolicy to be an explicit nil

### UnsetQosPolicy
`func (o *CreateProtectedObjectsRequest) UnsetQosPolicy()`

UnsetQosPolicy ensures that no value is present for QosPolicy, not even an explicit nil
### GetAbortInBlackouts

`func (o *CreateProtectedObjectsRequest) GetAbortInBlackouts() bool`

GetAbortInBlackouts returns the AbortInBlackouts field if non-nil, zero value otherwise.

### GetAbortInBlackoutsOk

`func (o *CreateProtectedObjectsRequest) GetAbortInBlackoutsOk() (*bool, bool)`

GetAbortInBlackoutsOk returns a tuple with the AbortInBlackouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbortInBlackouts

`func (o *CreateProtectedObjectsRequest) SetAbortInBlackouts(v bool)`

SetAbortInBlackouts sets AbortInBlackouts field to given value.

### HasAbortInBlackouts

`func (o *CreateProtectedObjectsRequest) HasAbortInBlackouts() bool`

HasAbortInBlackouts returns a boolean if a field has been set.

### SetAbortInBlackoutsNil

`func (o *CreateProtectedObjectsRequest) SetAbortInBlackoutsNil(b bool)`

 SetAbortInBlackoutsNil sets the value for AbortInBlackouts to be an explicit nil

### UnsetAbortInBlackouts
`func (o *CreateProtectedObjectsRequest) UnsetAbortInBlackouts()`

UnsetAbortInBlackouts ensures that no value is present for AbortInBlackouts, not even an explicit nil
### GetEndTimeUsecs

`func (o *CreateProtectedObjectsRequest) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *CreateProtectedObjectsRequest) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *CreateProtectedObjectsRequest) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *CreateProtectedObjectsRequest) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *CreateProtectedObjectsRequest) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *CreateProtectedObjectsRequest) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


