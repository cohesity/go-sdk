# ProtectedObjectBackupConfig

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
**Environment** | Pointer to **NullableString** | Specifies the environment for current object. | [optional] 
**VmwareParams** | Pointer to [**VmwareObjectProtectionResponseParams**](VmwareObjectProtectionResponseParams.md) |  | [optional] 
**GenericNasParams** | Pointer to [**GenericNasObjectProtectionResponseParams**](GenericNasObjectProtectionResponseParams.md) |  | [optional] 
**GpfsParams** | Pointer to [**GpfsObjectProtectionResponseParams**](GpfsObjectProtectionResponseParams.md) |  | [optional] 
**ElastifileParams** | Pointer to [**ElastifileObjectProtectionResponseParams**](ElastifileObjectProtectionResponseParams.md) |  | [optional] 
**NetappParams** | Pointer to [**NetappObjectProtectionResponseParams**](NetappObjectProtectionResponseParams.md) |  | [optional] 
**IsilonParams** | Pointer to [**IsilonObjectProtectionResponseParams**](IsilonObjectProtectionResponseParams.md) |  | [optional] 
**FlashbladeParams** | Pointer to [**FlashbladeObjectProtectionResponseParams**](FlashbladeObjectProtectionResponseParams.md) |  | [optional] 
**MssqlParams** | Pointer to [**MssqlObjectProtectionResponseParams**](MssqlObjectProtectionResponseParams.md) |  | [optional] 
**OracleParams** | Pointer to [**OracleObjectProtectionResponseParams**](OracleObjectProtectionResponseParams.md) |  | [optional] 
**Office365Params** | Pointer to [**Office365ObjectProtectionResponseParams**](Office365ObjectProtectionResponseParams.md) |  | [optional] 
**AwsParams** | Pointer to [**AwsObjectProtectionResponseParams**](AwsObjectProtectionResponseParams.md) |  | [optional] 
**HypervParams** | Pointer to [**HyperVObjectProtectionResponseParams**](HyperVObjectProtectionResponseParams.md) |  | [optional] 
**PhysicalParams** | Pointer to [**PhysicalObjectProtectionResponseParams**](PhysicalObjectProtectionResponseParams.md) |  | [optional] 
**IsAutoProtectConfig** | Pointer to **NullableBool** | Specifies whether or not this configuration is applied to an autoprotected object rather than this specific object. | [optional] 
**AutoProtectParentId** | Pointer to **NullableInt64** | Specifies the parent ID of the object which the backup configuration is applied to if this is an auto protect config. | [optional] 
**IsPaused** | Pointer to **NullableBool** | Specifies whether or not protection has been paused on this object. | [optional] 
**IsActive** | Pointer to **NullableBool** | Specifies whether or not protection has been deactivated on this object. | [optional] 

## Methods

### NewProtectedObjectBackupConfig

`func NewProtectedObjectBackupConfig() *ProtectedObjectBackupConfig`

NewProtectedObjectBackupConfig instantiates a new ProtectedObjectBackupConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectedObjectBackupConfigWithDefaults

`func NewProtectedObjectBackupConfigWithDefaults() *ProtectedObjectBackupConfig`

NewProtectedObjectBackupConfigWithDefaults instantiates a new ProtectedObjectBackupConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyId

`func (o *ProtectedObjectBackupConfig) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *ProtectedObjectBackupConfig) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *ProtectedObjectBackupConfig) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *ProtectedObjectBackupConfig) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### SetPolicyIdNil

`func (o *ProtectedObjectBackupConfig) SetPolicyIdNil(b bool)`

 SetPolicyIdNil sets the value for PolicyId to be an explicit nil

### UnsetPolicyId
`func (o *ProtectedObjectBackupConfig) UnsetPolicyId()`

UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
### GetPolicyConfig

`func (o *ProtectedObjectBackupConfig) GetPolicyConfig() PolicyConfig`

GetPolicyConfig returns the PolicyConfig field if non-nil, zero value otherwise.

### GetPolicyConfigOk

`func (o *ProtectedObjectBackupConfig) GetPolicyConfigOk() (*PolicyConfig, bool)`

GetPolicyConfigOk returns a tuple with the PolicyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyConfig

`func (o *ProtectedObjectBackupConfig) SetPolicyConfig(v PolicyConfig)`

SetPolicyConfig sets PolicyConfig field to given value.

### HasPolicyConfig

`func (o *ProtectedObjectBackupConfig) HasPolicyConfig() bool`

HasPolicyConfig returns a boolean if a field has been set.

### GetStorageDomainId

`func (o *ProtectedObjectBackupConfig) GetStorageDomainId() int64`

GetStorageDomainId returns the StorageDomainId field if non-nil, zero value otherwise.

### GetStorageDomainIdOk

`func (o *ProtectedObjectBackupConfig) GetStorageDomainIdOk() (*int64, bool)`

GetStorageDomainIdOk returns a tuple with the StorageDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainId

`func (o *ProtectedObjectBackupConfig) SetStorageDomainId(v int64)`

SetStorageDomainId sets StorageDomainId field to given value.

### HasStorageDomainId

`func (o *ProtectedObjectBackupConfig) HasStorageDomainId() bool`

HasStorageDomainId returns a boolean if a field has been set.

### SetStorageDomainIdNil

`func (o *ProtectedObjectBackupConfig) SetStorageDomainIdNil(b bool)`

 SetStorageDomainIdNil sets the value for StorageDomainId to be an explicit nil

### UnsetStorageDomainId
`func (o *ProtectedObjectBackupConfig) UnsetStorageDomainId()`

UnsetStorageDomainId ensures that no value is present for StorageDomainId, not even an explicit nil
### GetStartTime

`func (o *ProtectedObjectBackupConfig) GetStartTime() TimeOfDay`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ProtectedObjectBackupConfig) GetStartTimeOk() (*TimeOfDay, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ProtectedObjectBackupConfig) SetStartTime(v TimeOfDay)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ProtectedObjectBackupConfig) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetPriority

`func (o *ProtectedObjectBackupConfig) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ProtectedObjectBackupConfig) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ProtectedObjectBackupConfig) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ProtectedObjectBackupConfig) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *ProtectedObjectBackupConfig) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *ProtectedObjectBackupConfig) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetSla

`func (o *ProtectedObjectBackupConfig) GetSla() []SlaRule`

GetSla returns the Sla field if non-nil, zero value otherwise.

### GetSlaOk

`func (o *ProtectedObjectBackupConfig) GetSlaOk() (*[]SlaRule, bool)`

GetSlaOk returns a tuple with the Sla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSla

`func (o *ProtectedObjectBackupConfig) SetSla(v []SlaRule)`

SetSla sets Sla field to given value.

### HasSla

`func (o *ProtectedObjectBackupConfig) HasSla() bool`

HasSla returns a boolean if a field has been set.

### SetSlaNil

`func (o *ProtectedObjectBackupConfig) SetSlaNil(b bool)`

 SetSlaNil sets the value for Sla to be an explicit nil

### UnsetSla
`func (o *ProtectedObjectBackupConfig) UnsetSla()`

UnsetSla ensures that no value is present for Sla, not even an explicit nil
### GetQosPolicy

`func (o *ProtectedObjectBackupConfig) GetQosPolicy() string`

GetQosPolicy returns the QosPolicy field if non-nil, zero value otherwise.

### GetQosPolicyOk

`func (o *ProtectedObjectBackupConfig) GetQosPolicyOk() (*string, bool)`

GetQosPolicyOk returns a tuple with the QosPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosPolicy

`func (o *ProtectedObjectBackupConfig) SetQosPolicy(v string)`

SetQosPolicy sets QosPolicy field to given value.

### HasQosPolicy

`func (o *ProtectedObjectBackupConfig) HasQosPolicy() bool`

HasQosPolicy returns a boolean if a field has been set.

### SetQosPolicyNil

`func (o *ProtectedObjectBackupConfig) SetQosPolicyNil(b bool)`

 SetQosPolicyNil sets the value for QosPolicy to be an explicit nil

### UnsetQosPolicy
`func (o *ProtectedObjectBackupConfig) UnsetQosPolicy()`

UnsetQosPolicy ensures that no value is present for QosPolicy, not even an explicit nil
### GetAbortInBlackouts

`func (o *ProtectedObjectBackupConfig) GetAbortInBlackouts() bool`

GetAbortInBlackouts returns the AbortInBlackouts field if non-nil, zero value otherwise.

### GetAbortInBlackoutsOk

`func (o *ProtectedObjectBackupConfig) GetAbortInBlackoutsOk() (*bool, bool)`

GetAbortInBlackoutsOk returns a tuple with the AbortInBlackouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbortInBlackouts

`func (o *ProtectedObjectBackupConfig) SetAbortInBlackouts(v bool)`

SetAbortInBlackouts sets AbortInBlackouts field to given value.

### HasAbortInBlackouts

`func (o *ProtectedObjectBackupConfig) HasAbortInBlackouts() bool`

HasAbortInBlackouts returns a boolean if a field has been set.

### SetAbortInBlackoutsNil

`func (o *ProtectedObjectBackupConfig) SetAbortInBlackoutsNil(b bool)`

 SetAbortInBlackoutsNil sets the value for AbortInBlackouts to be an explicit nil

### UnsetAbortInBlackouts
`func (o *ProtectedObjectBackupConfig) UnsetAbortInBlackouts()`

UnsetAbortInBlackouts ensures that no value is present for AbortInBlackouts, not even an explicit nil
### GetEndTimeUsecs

`func (o *ProtectedObjectBackupConfig) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *ProtectedObjectBackupConfig) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *ProtectedObjectBackupConfig) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *ProtectedObjectBackupConfig) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *ProtectedObjectBackupConfig) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *ProtectedObjectBackupConfig) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetEnvironment

`func (o *ProtectedObjectBackupConfig) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ProtectedObjectBackupConfig) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ProtectedObjectBackupConfig) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ProtectedObjectBackupConfig) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *ProtectedObjectBackupConfig) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *ProtectedObjectBackupConfig) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetVmwareParams

`func (o *ProtectedObjectBackupConfig) GetVmwareParams() VmwareObjectProtectionResponseParams`

GetVmwareParams returns the VmwareParams field if non-nil, zero value otherwise.

### GetVmwareParamsOk

`func (o *ProtectedObjectBackupConfig) GetVmwareParamsOk() (*VmwareObjectProtectionResponseParams, bool)`

GetVmwareParamsOk returns a tuple with the VmwareParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareParams

`func (o *ProtectedObjectBackupConfig) SetVmwareParams(v VmwareObjectProtectionResponseParams)`

SetVmwareParams sets VmwareParams field to given value.

### HasVmwareParams

`func (o *ProtectedObjectBackupConfig) HasVmwareParams() bool`

HasVmwareParams returns a boolean if a field has been set.

### GetGenericNasParams

`func (o *ProtectedObjectBackupConfig) GetGenericNasParams() GenericNasObjectProtectionResponseParams`

GetGenericNasParams returns the GenericNasParams field if non-nil, zero value otherwise.

### GetGenericNasParamsOk

`func (o *ProtectedObjectBackupConfig) GetGenericNasParamsOk() (*GenericNasObjectProtectionResponseParams, bool)`

GetGenericNasParamsOk returns a tuple with the GenericNasParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericNasParams

`func (o *ProtectedObjectBackupConfig) SetGenericNasParams(v GenericNasObjectProtectionResponseParams)`

SetGenericNasParams sets GenericNasParams field to given value.

### HasGenericNasParams

`func (o *ProtectedObjectBackupConfig) HasGenericNasParams() bool`

HasGenericNasParams returns a boolean if a field has been set.

### GetGpfsParams

`func (o *ProtectedObjectBackupConfig) GetGpfsParams() GpfsObjectProtectionResponseParams`

GetGpfsParams returns the GpfsParams field if non-nil, zero value otherwise.

### GetGpfsParamsOk

`func (o *ProtectedObjectBackupConfig) GetGpfsParamsOk() (*GpfsObjectProtectionResponseParams, bool)`

GetGpfsParamsOk returns a tuple with the GpfsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpfsParams

`func (o *ProtectedObjectBackupConfig) SetGpfsParams(v GpfsObjectProtectionResponseParams)`

SetGpfsParams sets GpfsParams field to given value.

### HasGpfsParams

`func (o *ProtectedObjectBackupConfig) HasGpfsParams() bool`

HasGpfsParams returns a boolean if a field has been set.

### GetElastifileParams

`func (o *ProtectedObjectBackupConfig) GetElastifileParams() ElastifileObjectProtectionResponseParams`

GetElastifileParams returns the ElastifileParams field if non-nil, zero value otherwise.

### GetElastifileParamsOk

`func (o *ProtectedObjectBackupConfig) GetElastifileParamsOk() (*ElastifileObjectProtectionResponseParams, bool)`

GetElastifileParamsOk returns a tuple with the ElastifileParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElastifileParams

`func (o *ProtectedObjectBackupConfig) SetElastifileParams(v ElastifileObjectProtectionResponseParams)`

SetElastifileParams sets ElastifileParams field to given value.

### HasElastifileParams

`func (o *ProtectedObjectBackupConfig) HasElastifileParams() bool`

HasElastifileParams returns a boolean if a field has been set.

### GetNetappParams

`func (o *ProtectedObjectBackupConfig) GetNetappParams() NetappObjectProtectionResponseParams`

GetNetappParams returns the NetappParams field if non-nil, zero value otherwise.

### GetNetappParamsOk

`func (o *ProtectedObjectBackupConfig) GetNetappParamsOk() (*NetappObjectProtectionResponseParams, bool)`

GetNetappParamsOk returns a tuple with the NetappParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetappParams

`func (o *ProtectedObjectBackupConfig) SetNetappParams(v NetappObjectProtectionResponseParams)`

SetNetappParams sets NetappParams field to given value.

### HasNetappParams

`func (o *ProtectedObjectBackupConfig) HasNetappParams() bool`

HasNetappParams returns a boolean if a field has been set.

### GetIsilonParams

`func (o *ProtectedObjectBackupConfig) GetIsilonParams() IsilonObjectProtectionResponseParams`

GetIsilonParams returns the IsilonParams field if non-nil, zero value otherwise.

### GetIsilonParamsOk

`func (o *ProtectedObjectBackupConfig) GetIsilonParamsOk() (*IsilonObjectProtectionResponseParams, bool)`

GetIsilonParamsOk returns a tuple with the IsilonParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsilonParams

`func (o *ProtectedObjectBackupConfig) SetIsilonParams(v IsilonObjectProtectionResponseParams)`

SetIsilonParams sets IsilonParams field to given value.

### HasIsilonParams

`func (o *ProtectedObjectBackupConfig) HasIsilonParams() bool`

HasIsilonParams returns a boolean if a field has been set.

### GetFlashbladeParams

`func (o *ProtectedObjectBackupConfig) GetFlashbladeParams() FlashbladeObjectProtectionResponseParams`

GetFlashbladeParams returns the FlashbladeParams field if non-nil, zero value otherwise.

### GetFlashbladeParamsOk

`func (o *ProtectedObjectBackupConfig) GetFlashbladeParamsOk() (*FlashbladeObjectProtectionResponseParams, bool)`

GetFlashbladeParamsOk returns a tuple with the FlashbladeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlashbladeParams

`func (o *ProtectedObjectBackupConfig) SetFlashbladeParams(v FlashbladeObjectProtectionResponseParams)`

SetFlashbladeParams sets FlashbladeParams field to given value.

### HasFlashbladeParams

`func (o *ProtectedObjectBackupConfig) HasFlashbladeParams() bool`

HasFlashbladeParams returns a boolean if a field has been set.

### GetMssqlParams

`func (o *ProtectedObjectBackupConfig) GetMssqlParams() MssqlObjectProtectionResponseParams`

GetMssqlParams returns the MssqlParams field if non-nil, zero value otherwise.

### GetMssqlParamsOk

`func (o *ProtectedObjectBackupConfig) GetMssqlParamsOk() (*MssqlObjectProtectionResponseParams, bool)`

GetMssqlParamsOk returns a tuple with the MssqlParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlParams

`func (o *ProtectedObjectBackupConfig) SetMssqlParams(v MssqlObjectProtectionResponseParams)`

SetMssqlParams sets MssqlParams field to given value.

### HasMssqlParams

`func (o *ProtectedObjectBackupConfig) HasMssqlParams() bool`

HasMssqlParams returns a boolean if a field has been set.

### GetOracleParams

`func (o *ProtectedObjectBackupConfig) GetOracleParams() OracleObjectProtectionResponseParams`

GetOracleParams returns the OracleParams field if non-nil, zero value otherwise.

### GetOracleParamsOk

`func (o *ProtectedObjectBackupConfig) GetOracleParamsOk() (*OracleObjectProtectionResponseParams, bool)`

GetOracleParamsOk returns a tuple with the OracleParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleParams

`func (o *ProtectedObjectBackupConfig) SetOracleParams(v OracleObjectProtectionResponseParams)`

SetOracleParams sets OracleParams field to given value.

### HasOracleParams

`func (o *ProtectedObjectBackupConfig) HasOracleParams() bool`

HasOracleParams returns a boolean if a field has been set.

### GetOffice365Params

`func (o *ProtectedObjectBackupConfig) GetOffice365Params() Office365ObjectProtectionResponseParams`

GetOffice365Params returns the Office365Params field if non-nil, zero value otherwise.

### GetOffice365ParamsOk

`func (o *ProtectedObjectBackupConfig) GetOffice365ParamsOk() (*Office365ObjectProtectionResponseParams, bool)`

GetOffice365ParamsOk returns a tuple with the Office365Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffice365Params

`func (o *ProtectedObjectBackupConfig) SetOffice365Params(v Office365ObjectProtectionResponseParams)`

SetOffice365Params sets Office365Params field to given value.

### HasOffice365Params

`func (o *ProtectedObjectBackupConfig) HasOffice365Params() bool`

HasOffice365Params returns a boolean if a field has been set.

### GetAwsParams

`func (o *ProtectedObjectBackupConfig) GetAwsParams() AwsObjectProtectionResponseParams`

GetAwsParams returns the AwsParams field if non-nil, zero value otherwise.

### GetAwsParamsOk

`func (o *ProtectedObjectBackupConfig) GetAwsParamsOk() (*AwsObjectProtectionResponseParams, bool)`

GetAwsParamsOk returns a tuple with the AwsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsParams

`func (o *ProtectedObjectBackupConfig) SetAwsParams(v AwsObjectProtectionResponseParams)`

SetAwsParams sets AwsParams field to given value.

### HasAwsParams

`func (o *ProtectedObjectBackupConfig) HasAwsParams() bool`

HasAwsParams returns a boolean if a field has been set.

### GetHypervParams

`func (o *ProtectedObjectBackupConfig) GetHypervParams() HyperVObjectProtectionResponseParams`

GetHypervParams returns the HypervParams field if non-nil, zero value otherwise.

### GetHypervParamsOk

`func (o *ProtectedObjectBackupConfig) GetHypervParamsOk() (*HyperVObjectProtectionResponseParams, bool)`

GetHypervParamsOk returns a tuple with the HypervParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervParams

`func (o *ProtectedObjectBackupConfig) SetHypervParams(v HyperVObjectProtectionResponseParams)`

SetHypervParams sets HypervParams field to given value.

### HasHypervParams

`func (o *ProtectedObjectBackupConfig) HasHypervParams() bool`

HasHypervParams returns a boolean if a field has been set.

### GetPhysicalParams

`func (o *ProtectedObjectBackupConfig) GetPhysicalParams() PhysicalObjectProtectionResponseParams`

GetPhysicalParams returns the PhysicalParams field if non-nil, zero value otherwise.

### GetPhysicalParamsOk

`func (o *ProtectedObjectBackupConfig) GetPhysicalParamsOk() (*PhysicalObjectProtectionResponseParams, bool)`

GetPhysicalParamsOk returns a tuple with the PhysicalParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalParams

`func (o *ProtectedObjectBackupConfig) SetPhysicalParams(v PhysicalObjectProtectionResponseParams)`

SetPhysicalParams sets PhysicalParams field to given value.

### HasPhysicalParams

`func (o *ProtectedObjectBackupConfig) HasPhysicalParams() bool`

HasPhysicalParams returns a boolean if a field has been set.

### GetIsAutoProtectConfig

`func (o *ProtectedObjectBackupConfig) GetIsAutoProtectConfig() bool`

GetIsAutoProtectConfig returns the IsAutoProtectConfig field if non-nil, zero value otherwise.

### GetIsAutoProtectConfigOk

`func (o *ProtectedObjectBackupConfig) GetIsAutoProtectConfigOk() (*bool, bool)`

GetIsAutoProtectConfigOk returns a tuple with the IsAutoProtectConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoProtectConfig

`func (o *ProtectedObjectBackupConfig) SetIsAutoProtectConfig(v bool)`

SetIsAutoProtectConfig sets IsAutoProtectConfig field to given value.

### HasIsAutoProtectConfig

`func (o *ProtectedObjectBackupConfig) HasIsAutoProtectConfig() bool`

HasIsAutoProtectConfig returns a boolean if a field has been set.

### SetIsAutoProtectConfigNil

`func (o *ProtectedObjectBackupConfig) SetIsAutoProtectConfigNil(b bool)`

 SetIsAutoProtectConfigNil sets the value for IsAutoProtectConfig to be an explicit nil

### UnsetIsAutoProtectConfig
`func (o *ProtectedObjectBackupConfig) UnsetIsAutoProtectConfig()`

UnsetIsAutoProtectConfig ensures that no value is present for IsAutoProtectConfig, not even an explicit nil
### GetAutoProtectParentId

`func (o *ProtectedObjectBackupConfig) GetAutoProtectParentId() int64`

GetAutoProtectParentId returns the AutoProtectParentId field if non-nil, zero value otherwise.

### GetAutoProtectParentIdOk

`func (o *ProtectedObjectBackupConfig) GetAutoProtectParentIdOk() (*int64, bool)`

GetAutoProtectParentIdOk returns a tuple with the AutoProtectParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoProtectParentId

`func (o *ProtectedObjectBackupConfig) SetAutoProtectParentId(v int64)`

SetAutoProtectParentId sets AutoProtectParentId field to given value.

### HasAutoProtectParentId

`func (o *ProtectedObjectBackupConfig) HasAutoProtectParentId() bool`

HasAutoProtectParentId returns a boolean if a field has been set.

### SetAutoProtectParentIdNil

`func (o *ProtectedObjectBackupConfig) SetAutoProtectParentIdNil(b bool)`

 SetAutoProtectParentIdNil sets the value for AutoProtectParentId to be an explicit nil

### UnsetAutoProtectParentId
`func (o *ProtectedObjectBackupConfig) UnsetAutoProtectParentId()`

UnsetAutoProtectParentId ensures that no value is present for AutoProtectParentId, not even an explicit nil
### GetIsPaused

`func (o *ProtectedObjectBackupConfig) GetIsPaused() bool`

GetIsPaused returns the IsPaused field if non-nil, zero value otherwise.

### GetIsPausedOk

`func (o *ProtectedObjectBackupConfig) GetIsPausedOk() (*bool, bool)`

GetIsPausedOk returns a tuple with the IsPaused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPaused

`func (o *ProtectedObjectBackupConfig) SetIsPaused(v bool)`

SetIsPaused sets IsPaused field to given value.

### HasIsPaused

`func (o *ProtectedObjectBackupConfig) HasIsPaused() bool`

HasIsPaused returns a boolean if a field has been set.

### SetIsPausedNil

`func (o *ProtectedObjectBackupConfig) SetIsPausedNil(b bool)`

 SetIsPausedNil sets the value for IsPaused to be an explicit nil

### UnsetIsPaused
`func (o *ProtectedObjectBackupConfig) UnsetIsPaused()`

UnsetIsPaused ensures that no value is present for IsPaused, not even an explicit nil
### GetIsActive

`func (o *ProtectedObjectBackupConfig) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ProtectedObjectBackupConfig) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ProtectedObjectBackupConfig) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *ProtectedObjectBackupConfig) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *ProtectedObjectBackupConfig) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *ProtectedObjectBackupConfig) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


