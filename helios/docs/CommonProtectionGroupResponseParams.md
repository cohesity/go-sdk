# CommonProtectionGroupResponseParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Specifies the ID of the Protection Group. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the Protection Group. | [optional] 
**PolicyId** | Pointer to **NullableString** | Specifies the unique id of the Protection Policy associated with the Protection Group. The Policy provides retry settings Protection Schedules, Priority, SLA, etc. | [optional] 
**Priority** | Pointer to **NullableString** | Specifies the priority of the Protection Group. | [optional] 
**StorageDomainId** | Pointer to **NullableInt64** | Specifies the Storage Domain (View Box) ID where this Protection Group writes data. | [optional] 
**Description** | Pointer to **NullableString** | Specifies a description of the Protection Group. | [optional] 
**StartTime** | Pointer to [**TimeOfDay**](TimeOfDay.md) |  | [optional] 
**EndTimeUsecs** | Pointer to **NullableInt64** | Specifies the end time in micro seconds for this Protection Group. If this is not specified, the Protection Group won&#39;t be ended. | [optional] 
**AlertPolicy** | Pointer to [**ProtectionGroupAlertingPolicy**](ProtectionGroupAlertingPolicy.md) |  | [optional] 
**Sla** | Pointer to [**[]SlaRule**](SlaRule.md) | Specifies the SLA parameters for this Protection Group. | [optional] 
**QosPolicy** | Pointer to **NullableString** | Specifies whether the Protection Group will be written to HDD or SSD. | [optional] 
**AbortInBlackouts** | Pointer to **NullableBool** | Specifies whether currently executing jobs should abort if a blackout period specified by a policy starts. Available only if the selected policy has at least one blackout period. Default value is false. | [optional] 
**PauseInBlackouts** | Pointer to **NullableBool** | Specifies whether currently executing jobs should be paused if a blackout period specified by a policy starts. Available only if the selected policy has at least one blackout period. Default value is false. This field should not be set to true if &#39;abortInBlackouts&#39; is sent as true. | [optional] 
**IsActive** | Pointer to **NullableBool** | Specifies if the Protection Group is active or not. | [optional] 
**IsDeleted** | Pointer to **NullableBool** | Specifies if the Protection Group has been deleted. | [optional] 
**IsPaused** | Pointer to **NullableBool** | Specifies if the the Protection Group is paused. New runs are not scheduled for the paused Protection Groups. Active run if any is not impacted. | [optional] 
**Environment** | Pointer to **NullableString** | Specifies the environment of the Protection Group. | [optional] 
**LastRun** | Pointer to [**ProtectionGroupRun**](ProtectionGroupRun.md) |  | [optional] 
**Permissions** | Pointer to [**[]Tenant**](Tenant.md) | Specifies the list of tenants that have permissions for this protection group. | [optional] 
**IsProtectOnce** | Pointer to **NullableBool** | Specifies if the the Protection Group is using a protect once type of policy. This field is helpful to identify run happen for this group. | [optional] 
**MissingEntities** | Pointer to [**[]MissingEntityParams**](MissingEntityParams.md) | Specifies the Information about missing entities. | [optional] 
**NumProtectedObjects** | Pointer to **NullableInt64** | Specifies the number of protected objects of the Protection Group. | [optional] 

## Methods

### NewCommonProtectionGroupResponseParams

`func NewCommonProtectionGroupResponseParams() *CommonProtectionGroupResponseParams`

NewCommonProtectionGroupResponseParams instantiates a new CommonProtectionGroupResponseParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonProtectionGroupResponseParamsWithDefaults

`func NewCommonProtectionGroupResponseParamsWithDefaults() *CommonProtectionGroupResponseParams`

NewCommonProtectionGroupResponseParamsWithDefaults instantiates a new CommonProtectionGroupResponseParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommonProtectionGroupResponseParams) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommonProtectionGroupResponseParams) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommonProtectionGroupResponseParams) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CommonProtectionGroupResponseParams) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CommonProtectionGroupResponseParams) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CommonProtectionGroupResponseParams) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *CommonProtectionGroupResponseParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CommonProtectionGroupResponseParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CommonProtectionGroupResponseParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CommonProtectionGroupResponseParams) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CommonProtectionGroupResponseParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CommonProtectionGroupResponseParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPolicyId

`func (o *CommonProtectionGroupResponseParams) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *CommonProtectionGroupResponseParams) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *CommonProtectionGroupResponseParams) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *CommonProtectionGroupResponseParams) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### SetPolicyIdNil

`func (o *CommonProtectionGroupResponseParams) SetPolicyIdNil(b bool)`

 SetPolicyIdNil sets the value for PolicyId to be an explicit nil

### UnsetPolicyId
`func (o *CommonProtectionGroupResponseParams) UnsetPolicyId()`

UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
### GetPriority

`func (o *CommonProtectionGroupResponseParams) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CommonProtectionGroupResponseParams) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CommonProtectionGroupResponseParams) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CommonProtectionGroupResponseParams) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *CommonProtectionGroupResponseParams) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *CommonProtectionGroupResponseParams) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetStorageDomainId

`func (o *CommonProtectionGroupResponseParams) GetStorageDomainId() int64`

GetStorageDomainId returns the StorageDomainId field if non-nil, zero value otherwise.

### GetStorageDomainIdOk

`func (o *CommonProtectionGroupResponseParams) GetStorageDomainIdOk() (*int64, bool)`

GetStorageDomainIdOk returns a tuple with the StorageDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainId

`func (o *CommonProtectionGroupResponseParams) SetStorageDomainId(v int64)`

SetStorageDomainId sets StorageDomainId field to given value.

### HasStorageDomainId

`func (o *CommonProtectionGroupResponseParams) HasStorageDomainId() bool`

HasStorageDomainId returns a boolean if a field has been set.

### SetStorageDomainIdNil

`func (o *CommonProtectionGroupResponseParams) SetStorageDomainIdNil(b bool)`

 SetStorageDomainIdNil sets the value for StorageDomainId to be an explicit nil

### UnsetStorageDomainId
`func (o *CommonProtectionGroupResponseParams) UnsetStorageDomainId()`

UnsetStorageDomainId ensures that no value is present for StorageDomainId, not even an explicit nil
### GetDescription

`func (o *CommonProtectionGroupResponseParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CommonProtectionGroupResponseParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CommonProtectionGroupResponseParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CommonProtectionGroupResponseParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CommonProtectionGroupResponseParams) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CommonProtectionGroupResponseParams) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetStartTime

`func (o *CommonProtectionGroupResponseParams) GetStartTime() TimeOfDay`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *CommonProtectionGroupResponseParams) GetStartTimeOk() (*TimeOfDay, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *CommonProtectionGroupResponseParams) SetStartTime(v TimeOfDay)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *CommonProtectionGroupResponseParams) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTimeUsecs

`func (o *CommonProtectionGroupResponseParams) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *CommonProtectionGroupResponseParams) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *CommonProtectionGroupResponseParams) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *CommonProtectionGroupResponseParams) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *CommonProtectionGroupResponseParams) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *CommonProtectionGroupResponseParams) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetAlertPolicy

`func (o *CommonProtectionGroupResponseParams) GetAlertPolicy() ProtectionGroupAlertingPolicy`

GetAlertPolicy returns the AlertPolicy field if non-nil, zero value otherwise.

### GetAlertPolicyOk

`func (o *CommonProtectionGroupResponseParams) GetAlertPolicyOk() (*ProtectionGroupAlertingPolicy, bool)`

GetAlertPolicyOk returns a tuple with the AlertPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertPolicy

`func (o *CommonProtectionGroupResponseParams) SetAlertPolicy(v ProtectionGroupAlertingPolicy)`

SetAlertPolicy sets AlertPolicy field to given value.

### HasAlertPolicy

`func (o *CommonProtectionGroupResponseParams) HasAlertPolicy() bool`

HasAlertPolicy returns a boolean if a field has been set.

### GetSla

`func (o *CommonProtectionGroupResponseParams) GetSla() []SlaRule`

GetSla returns the Sla field if non-nil, zero value otherwise.

### GetSlaOk

`func (o *CommonProtectionGroupResponseParams) GetSlaOk() (*[]SlaRule, bool)`

GetSlaOk returns a tuple with the Sla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSla

`func (o *CommonProtectionGroupResponseParams) SetSla(v []SlaRule)`

SetSla sets Sla field to given value.

### HasSla

`func (o *CommonProtectionGroupResponseParams) HasSla() bool`

HasSla returns a boolean if a field has been set.

### SetSlaNil

`func (o *CommonProtectionGroupResponseParams) SetSlaNil(b bool)`

 SetSlaNil sets the value for Sla to be an explicit nil

### UnsetSla
`func (o *CommonProtectionGroupResponseParams) UnsetSla()`

UnsetSla ensures that no value is present for Sla, not even an explicit nil
### GetQosPolicy

`func (o *CommonProtectionGroupResponseParams) GetQosPolicy() string`

GetQosPolicy returns the QosPolicy field if non-nil, zero value otherwise.

### GetQosPolicyOk

`func (o *CommonProtectionGroupResponseParams) GetQosPolicyOk() (*string, bool)`

GetQosPolicyOk returns a tuple with the QosPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosPolicy

`func (o *CommonProtectionGroupResponseParams) SetQosPolicy(v string)`

SetQosPolicy sets QosPolicy field to given value.

### HasQosPolicy

`func (o *CommonProtectionGroupResponseParams) HasQosPolicy() bool`

HasQosPolicy returns a boolean if a field has been set.

### SetQosPolicyNil

`func (o *CommonProtectionGroupResponseParams) SetQosPolicyNil(b bool)`

 SetQosPolicyNil sets the value for QosPolicy to be an explicit nil

### UnsetQosPolicy
`func (o *CommonProtectionGroupResponseParams) UnsetQosPolicy()`

UnsetQosPolicy ensures that no value is present for QosPolicy, not even an explicit nil
### GetAbortInBlackouts

`func (o *CommonProtectionGroupResponseParams) GetAbortInBlackouts() bool`

GetAbortInBlackouts returns the AbortInBlackouts field if non-nil, zero value otherwise.

### GetAbortInBlackoutsOk

`func (o *CommonProtectionGroupResponseParams) GetAbortInBlackoutsOk() (*bool, bool)`

GetAbortInBlackoutsOk returns a tuple with the AbortInBlackouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbortInBlackouts

`func (o *CommonProtectionGroupResponseParams) SetAbortInBlackouts(v bool)`

SetAbortInBlackouts sets AbortInBlackouts field to given value.

### HasAbortInBlackouts

`func (o *CommonProtectionGroupResponseParams) HasAbortInBlackouts() bool`

HasAbortInBlackouts returns a boolean if a field has been set.

### SetAbortInBlackoutsNil

`func (o *CommonProtectionGroupResponseParams) SetAbortInBlackoutsNil(b bool)`

 SetAbortInBlackoutsNil sets the value for AbortInBlackouts to be an explicit nil

### UnsetAbortInBlackouts
`func (o *CommonProtectionGroupResponseParams) UnsetAbortInBlackouts()`

UnsetAbortInBlackouts ensures that no value is present for AbortInBlackouts, not even an explicit nil
### GetPauseInBlackouts

`func (o *CommonProtectionGroupResponseParams) GetPauseInBlackouts() bool`

GetPauseInBlackouts returns the PauseInBlackouts field if non-nil, zero value otherwise.

### GetPauseInBlackoutsOk

`func (o *CommonProtectionGroupResponseParams) GetPauseInBlackoutsOk() (*bool, bool)`

GetPauseInBlackoutsOk returns a tuple with the PauseInBlackouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseInBlackouts

`func (o *CommonProtectionGroupResponseParams) SetPauseInBlackouts(v bool)`

SetPauseInBlackouts sets PauseInBlackouts field to given value.

### HasPauseInBlackouts

`func (o *CommonProtectionGroupResponseParams) HasPauseInBlackouts() bool`

HasPauseInBlackouts returns a boolean if a field has been set.

### SetPauseInBlackoutsNil

`func (o *CommonProtectionGroupResponseParams) SetPauseInBlackoutsNil(b bool)`

 SetPauseInBlackoutsNil sets the value for PauseInBlackouts to be an explicit nil

### UnsetPauseInBlackouts
`func (o *CommonProtectionGroupResponseParams) UnsetPauseInBlackouts()`

UnsetPauseInBlackouts ensures that no value is present for PauseInBlackouts, not even an explicit nil
### GetIsActive

`func (o *CommonProtectionGroupResponseParams) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *CommonProtectionGroupResponseParams) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *CommonProtectionGroupResponseParams) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *CommonProtectionGroupResponseParams) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *CommonProtectionGroupResponseParams) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *CommonProtectionGroupResponseParams) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetIsDeleted

`func (o *CommonProtectionGroupResponseParams) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *CommonProtectionGroupResponseParams) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *CommonProtectionGroupResponseParams) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *CommonProtectionGroupResponseParams) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### SetIsDeletedNil

`func (o *CommonProtectionGroupResponseParams) SetIsDeletedNil(b bool)`

 SetIsDeletedNil sets the value for IsDeleted to be an explicit nil

### UnsetIsDeleted
`func (o *CommonProtectionGroupResponseParams) UnsetIsDeleted()`

UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil
### GetIsPaused

`func (o *CommonProtectionGroupResponseParams) GetIsPaused() bool`

GetIsPaused returns the IsPaused field if non-nil, zero value otherwise.

### GetIsPausedOk

`func (o *CommonProtectionGroupResponseParams) GetIsPausedOk() (*bool, bool)`

GetIsPausedOk returns a tuple with the IsPaused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPaused

`func (o *CommonProtectionGroupResponseParams) SetIsPaused(v bool)`

SetIsPaused sets IsPaused field to given value.

### HasIsPaused

`func (o *CommonProtectionGroupResponseParams) HasIsPaused() bool`

HasIsPaused returns a boolean if a field has been set.

### SetIsPausedNil

`func (o *CommonProtectionGroupResponseParams) SetIsPausedNil(b bool)`

 SetIsPausedNil sets the value for IsPaused to be an explicit nil

### UnsetIsPaused
`func (o *CommonProtectionGroupResponseParams) UnsetIsPaused()`

UnsetIsPaused ensures that no value is present for IsPaused, not even an explicit nil
### GetEnvironment

`func (o *CommonProtectionGroupResponseParams) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CommonProtectionGroupResponseParams) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CommonProtectionGroupResponseParams) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CommonProtectionGroupResponseParams) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *CommonProtectionGroupResponseParams) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *CommonProtectionGroupResponseParams) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetLastRun

`func (o *CommonProtectionGroupResponseParams) GetLastRun() ProtectionGroupRun`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *CommonProtectionGroupResponseParams) GetLastRunOk() (*ProtectionGroupRun, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *CommonProtectionGroupResponseParams) SetLastRun(v ProtectionGroupRun)`

SetLastRun sets LastRun field to given value.

### HasLastRun

`func (o *CommonProtectionGroupResponseParams) HasLastRun() bool`

HasLastRun returns a boolean if a field has been set.

### GetPermissions

`func (o *CommonProtectionGroupResponseParams) GetPermissions() []Tenant`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *CommonProtectionGroupResponseParams) GetPermissionsOk() (*[]Tenant, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *CommonProtectionGroupResponseParams) SetPermissions(v []Tenant)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *CommonProtectionGroupResponseParams) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *CommonProtectionGroupResponseParams) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *CommonProtectionGroupResponseParams) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetIsProtectOnce

`func (o *CommonProtectionGroupResponseParams) GetIsProtectOnce() bool`

GetIsProtectOnce returns the IsProtectOnce field if non-nil, zero value otherwise.

### GetIsProtectOnceOk

`func (o *CommonProtectionGroupResponseParams) GetIsProtectOnceOk() (*bool, bool)`

GetIsProtectOnceOk returns a tuple with the IsProtectOnce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProtectOnce

`func (o *CommonProtectionGroupResponseParams) SetIsProtectOnce(v bool)`

SetIsProtectOnce sets IsProtectOnce field to given value.

### HasIsProtectOnce

`func (o *CommonProtectionGroupResponseParams) HasIsProtectOnce() bool`

HasIsProtectOnce returns a boolean if a field has been set.

### SetIsProtectOnceNil

`func (o *CommonProtectionGroupResponseParams) SetIsProtectOnceNil(b bool)`

 SetIsProtectOnceNil sets the value for IsProtectOnce to be an explicit nil

### UnsetIsProtectOnce
`func (o *CommonProtectionGroupResponseParams) UnsetIsProtectOnce()`

UnsetIsProtectOnce ensures that no value is present for IsProtectOnce, not even an explicit nil
### GetMissingEntities

`func (o *CommonProtectionGroupResponseParams) GetMissingEntities() []MissingEntityParams`

GetMissingEntities returns the MissingEntities field if non-nil, zero value otherwise.

### GetMissingEntitiesOk

`func (o *CommonProtectionGroupResponseParams) GetMissingEntitiesOk() (*[]MissingEntityParams, bool)`

GetMissingEntitiesOk returns a tuple with the MissingEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingEntities

`func (o *CommonProtectionGroupResponseParams) SetMissingEntities(v []MissingEntityParams)`

SetMissingEntities sets MissingEntities field to given value.

### HasMissingEntities

`func (o *CommonProtectionGroupResponseParams) HasMissingEntities() bool`

HasMissingEntities returns a boolean if a field has been set.

### SetMissingEntitiesNil

`func (o *CommonProtectionGroupResponseParams) SetMissingEntitiesNil(b bool)`

 SetMissingEntitiesNil sets the value for MissingEntities to be an explicit nil

### UnsetMissingEntities
`func (o *CommonProtectionGroupResponseParams) UnsetMissingEntities()`

UnsetMissingEntities ensures that no value is present for MissingEntities, not even an explicit nil
### GetNumProtectedObjects

`func (o *CommonProtectionGroupResponseParams) GetNumProtectedObjects() int64`

GetNumProtectedObjects returns the NumProtectedObjects field if non-nil, zero value otherwise.

### GetNumProtectedObjectsOk

`func (o *CommonProtectionGroupResponseParams) GetNumProtectedObjectsOk() (*int64, bool)`

GetNumProtectedObjectsOk returns a tuple with the NumProtectedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumProtectedObjects

`func (o *CommonProtectionGroupResponseParams) SetNumProtectedObjects(v int64)`

SetNumProtectedObjects sets NumProtectedObjects field to given value.

### HasNumProtectedObjects

`func (o *CommonProtectionGroupResponseParams) HasNumProtectedObjects() bool`

HasNumProtectedObjects returns a boolean if a field has been set.

### SetNumProtectedObjectsNil

`func (o *CommonProtectionGroupResponseParams) SetNumProtectedObjectsNil(b bool)`

 SetNumProtectedObjectsNil sets the value for NumProtectedObjects to be an explicit nil

### UnsetNumProtectedObjects
`func (o *CommonProtectionGroupResponseParams) UnsetNumProtectedObjects()`

UnsetNumProtectedObjects ensures that no value is present for NumProtectedObjects, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


