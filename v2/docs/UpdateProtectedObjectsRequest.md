# UpdateProtectedObjectsRequest

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
**AwsParams** | Pointer to [**AwsObjectProtectionUpdateRequestParams**](AwsObjectProtectionUpdateRequestParams.md) |  | [optional] 
**AzureParams** | Pointer to [**AzureObjectProtectionUpdateRequestParams**](AzureObjectProtectionUpdateRequestParams.md) |  | [optional] 
**ElastifileParams** | Pointer to [**ElastifileObjectProtectionUpdateRequestParams**](ElastifileObjectProtectionUpdateRequestParams.md) |  | [optional] 
**Environment** | Pointer to **NullableString** | Specifies the environment for current object. | [optional] 
**ExperimentalAdapterParams** | Pointer to [**ExperimentalAdapterObjectProtectionUpdateRequestParams**](ExperimentalAdapterObjectProtectionUpdateRequestParams.md) |  | [optional] 
**FlashbladeParams** | Pointer to [**FlashbladeObjectProtectionUpdateRequestParams**](FlashbladeObjectProtectionUpdateRequestParams.md) |  | [optional] 
**GenericNasParams** | Pointer to [**GenericNasObjectProtectionUpdateRequestParams**](GenericNasObjectProtectionUpdateRequestParams.md) |  | [optional] 
**GpfsParams** | Pointer to [**GpfsObjectProtectionUpdateRequestParams**](GpfsObjectProtectionUpdateRequestParams.md) |  | [optional] 
**HypervParams** | Pointer to [**HyperVObjectProtectionUpdateRequestParams**](HyperVObjectProtectionUpdateRequestParams.md) |  | [optional] 
**IsilonParams** | Pointer to [**IsilonObjectProtectionUpdateRequestParams**](IsilonObjectProtectionUpdateRequestParams.md) |  | [optional] 
**MssqlParams** | Pointer to [**MssqlObjectProtectionUpdateRequestParams**](MssqlObjectProtectionUpdateRequestParams.md) |  | [optional] 
**NetappParams** | Pointer to [**NetappObjectProtectionUpdateRequestParams**](NetappObjectProtectionUpdateRequestParams.md) |  | [optional] 
**Office365Params** | Pointer to [**Office365ObjectProtectionUpdateRequestParams**](Office365ObjectProtectionUpdateRequestParams.md) |  | [optional] 
**OracleParams** | Pointer to [**OracleObjectProtectionUpdateRequestParams**](OracleObjectProtectionUpdateRequestParams.md) |  | [optional] 
**PhysicalParams** | Pointer to [**PhysicalObjectProtectionUpdateRequestParams**](PhysicalObjectProtectionUpdateRequestParams.md) |  | [optional] 
**SapHanaParams** | Pointer to [**SapHanaObjectProtectionUpdateRequestParams**](SapHanaObjectProtectionUpdateRequestParams.md) |  | [optional] 
**SfdcParams** | Pointer to [**SfdcObjectProtectionUpdateRequestParams**](SfdcObjectProtectionUpdateRequestParams.md) |  | [optional] 
**UdaParams** | Pointer to [**UdaObjectProtectionUpdateRequestParams**](UdaObjectProtectionUpdateRequestParams.md) |  | [optional] 
**VmwareParams** | Pointer to [**VmwareObjectProtectionUpdateRequestParams**](VmwareObjectProtectionUpdateRequestParams.md) |  | [optional] 

## Methods

### NewUpdateProtectedObjectsRequest

`func NewUpdateProtectedObjectsRequest() *UpdateProtectedObjectsRequest`

NewUpdateProtectedObjectsRequest instantiates a new UpdateProtectedObjectsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProtectedObjectsRequestWithDefaults

`func NewUpdateProtectedObjectsRequestWithDefaults() *UpdateProtectedObjectsRequest`

NewUpdateProtectedObjectsRequestWithDefaults instantiates a new UpdateProtectedObjectsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbortInBlackouts

`func (o *UpdateProtectedObjectsRequest) GetAbortInBlackouts() bool`

GetAbortInBlackouts returns the AbortInBlackouts field if non-nil, zero value otherwise.

### GetAbortInBlackoutsOk

`func (o *UpdateProtectedObjectsRequest) GetAbortInBlackoutsOk() (*bool, bool)`

GetAbortInBlackoutsOk returns a tuple with the AbortInBlackouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbortInBlackouts

`func (o *UpdateProtectedObjectsRequest) SetAbortInBlackouts(v bool)`

SetAbortInBlackouts sets AbortInBlackouts field to given value.

### HasAbortInBlackouts

`func (o *UpdateProtectedObjectsRequest) HasAbortInBlackouts() bool`

HasAbortInBlackouts returns a boolean if a field has been set.

### SetAbortInBlackoutsNil

`func (o *UpdateProtectedObjectsRequest) SetAbortInBlackoutsNil(b bool)`

 SetAbortInBlackoutsNil sets the value for AbortInBlackouts to be an explicit nil

### UnsetAbortInBlackouts
`func (o *UpdateProtectedObjectsRequest) UnsetAbortInBlackouts()`

UnsetAbortInBlackouts ensures that no value is present for AbortInBlackouts, not even an explicit nil
### GetEndTimeUsecs

`func (o *UpdateProtectedObjectsRequest) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *UpdateProtectedObjectsRequest) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *UpdateProtectedObjectsRequest) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *UpdateProtectedObjectsRequest) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *UpdateProtectedObjectsRequest) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *UpdateProtectedObjectsRequest) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetPolicyConfig

`func (o *UpdateProtectedObjectsRequest) GetPolicyConfig() PolicyConfig`

GetPolicyConfig returns the PolicyConfig field if non-nil, zero value otherwise.

### GetPolicyConfigOk

`func (o *UpdateProtectedObjectsRequest) GetPolicyConfigOk() (*PolicyConfig, bool)`

GetPolicyConfigOk returns a tuple with the PolicyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyConfig

`func (o *UpdateProtectedObjectsRequest) SetPolicyConfig(v PolicyConfig)`

SetPolicyConfig sets PolicyConfig field to given value.

### HasPolicyConfig

`func (o *UpdateProtectedObjectsRequest) HasPolicyConfig() bool`

HasPolicyConfig returns a boolean if a field has been set.

### GetPolicyId

`func (o *UpdateProtectedObjectsRequest) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *UpdateProtectedObjectsRequest) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *UpdateProtectedObjectsRequest) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *UpdateProtectedObjectsRequest) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### SetPolicyIdNil

`func (o *UpdateProtectedObjectsRequest) SetPolicyIdNil(b bool)`

 SetPolicyIdNil sets the value for PolicyId to be an explicit nil

### UnsetPolicyId
`func (o *UpdateProtectedObjectsRequest) UnsetPolicyId()`

UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
### GetPriority

`func (o *UpdateProtectedObjectsRequest) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *UpdateProtectedObjectsRequest) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *UpdateProtectedObjectsRequest) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *UpdateProtectedObjectsRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *UpdateProtectedObjectsRequest) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *UpdateProtectedObjectsRequest) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetQosPolicy

`func (o *UpdateProtectedObjectsRequest) GetQosPolicy() string`

GetQosPolicy returns the QosPolicy field if non-nil, zero value otherwise.

### GetQosPolicyOk

`func (o *UpdateProtectedObjectsRequest) GetQosPolicyOk() (*string, bool)`

GetQosPolicyOk returns a tuple with the QosPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosPolicy

`func (o *UpdateProtectedObjectsRequest) SetQosPolicy(v string)`

SetQosPolicy sets QosPolicy field to given value.

### HasQosPolicy

`func (o *UpdateProtectedObjectsRequest) HasQosPolicy() bool`

HasQosPolicy returns a boolean if a field has been set.

### SetQosPolicyNil

`func (o *UpdateProtectedObjectsRequest) SetQosPolicyNil(b bool)`

 SetQosPolicyNil sets the value for QosPolicy to be an explicit nil

### UnsetQosPolicy
`func (o *UpdateProtectedObjectsRequest) UnsetQosPolicy()`

UnsetQosPolicy ensures that no value is present for QosPolicy, not even an explicit nil
### GetSkipRigelForBackup

`func (o *UpdateProtectedObjectsRequest) GetSkipRigelForBackup() bool`

GetSkipRigelForBackup returns the SkipRigelForBackup field if non-nil, zero value otherwise.

### GetSkipRigelForBackupOk

`func (o *UpdateProtectedObjectsRequest) GetSkipRigelForBackupOk() (*bool, bool)`

GetSkipRigelForBackupOk returns a tuple with the SkipRigelForBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipRigelForBackup

`func (o *UpdateProtectedObjectsRequest) SetSkipRigelForBackup(v bool)`

SetSkipRigelForBackup sets SkipRigelForBackup field to given value.

### HasSkipRigelForBackup

`func (o *UpdateProtectedObjectsRequest) HasSkipRigelForBackup() bool`

HasSkipRigelForBackup returns a boolean if a field has been set.

### SetSkipRigelForBackupNil

`func (o *UpdateProtectedObjectsRequest) SetSkipRigelForBackupNil(b bool)`

 SetSkipRigelForBackupNil sets the value for SkipRigelForBackup to be an explicit nil

### UnsetSkipRigelForBackup
`func (o *UpdateProtectedObjectsRequest) UnsetSkipRigelForBackup()`

UnsetSkipRigelForBackup ensures that no value is present for SkipRigelForBackup, not even an explicit nil
### GetSla

`func (o *UpdateProtectedObjectsRequest) GetSla() []SlaRule`

GetSla returns the Sla field if non-nil, zero value otherwise.

### GetSlaOk

`func (o *UpdateProtectedObjectsRequest) GetSlaOk() (*[]SlaRule, bool)`

GetSlaOk returns a tuple with the Sla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSla

`func (o *UpdateProtectedObjectsRequest) SetSla(v []SlaRule)`

SetSla sets Sla field to given value.

### HasSla

`func (o *UpdateProtectedObjectsRequest) HasSla() bool`

HasSla returns a boolean if a field has been set.

### SetSlaNil

`func (o *UpdateProtectedObjectsRequest) SetSlaNil(b bool)`

 SetSlaNil sets the value for Sla to be an explicit nil

### UnsetSla
`func (o *UpdateProtectedObjectsRequest) UnsetSla()`

UnsetSla ensures that no value is present for Sla, not even an explicit nil
### GetStartTime

`func (o *UpdateProtectedObjectsRequest) GetStartTime() TimeOfDay`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *UpdateProtectedObjectsRequest) GetStartTimeOk() (*TimeOfDay, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *UpdateProtectedObjectsRequest) SetStartTime(v TimeOfDay)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *UpdateProtectedObjectsRequest) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStorageDomainId

`func (o *UpdateProtectedObjectsRequest) GetStorageDomainId() int64`

GetStorageDomainId returns the StorageDomainId field if non-nil, zero value otherwise.

### GetStorageDomainIdOk

`func (o *UpdateProtectedObjectsRequest) GetStorageDomainIdOk() (*int64, bool)`

GetStorageDomainIdOk returns a tuple with the StorageDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainId

`func (o *UpdateProtectedObjectsRequest) SetStorageDomainId(v int64)`

SetStorageDomainId sets StorageDomainId field to given value.

### HasStorageDomainId

`func (o *UpdateProtectedObjectsRequest) HasStorageDomainId() bool`

HasStorageDomainId returns a boolean if a field has been set.

### SetStorageDomainIdNil

`func (o *UpdateProtectedObjectsRequest) SetStorageDomainIdNil(b bool)`

 SetStorageDomainIdNil sets the value for StorageDomainId to be an explicit nil

### UnsetStorageDomainId
`func (o *UpdateProtectedObjectsRequest) UnsetStorageDomainId()`

UnsetStorageDomainId ensures that no value is present for StorageDomainId, not even an explicit nil
### GetAwsParams

`func (o *UpdateProtectedObjectsRequest) GetAwsParams() AwsObjectProtectionUpdateRequestParams`

GetAwsParams returns the AwsParams field if non-nil, zero value otherwise.

### GetAwsParamsOk

`func (o *UpdateProtectedObjectsRequest) GetAwsParamsOk() (*AwsObjectProtectionUpdateRequestParams, bool)`

GetAwsParamsOk returns a tuple with the AwsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsParams

`func (o *UpdateProtectedObjectsRequest) SetAwsParams(v AwsObjectProtectionUpdateRequestParams)`

SetAwsParams sets AwsParams field to given value.

### HasAwsParams

`func (o *UpdateProtectedObjectsRequest) HasAwsParams() bool`

HasAwsParams returns a boolean if a field has been set.

### GetAzureParams

`func (o *UpdateProtectedObjectsRequest) GetAzureParams() AzureObjectProtectionUpdateRequestParams`

GetAzureParams returns the AzureParams field if non-nil, zero value otherwise.

### GetAzureParamsOk

`func (o *UpdateProtectedObjectsRequest) GetAzureParamsOk() (*AzureObjectProtectionUpdateRequestParams, bool)`

GetAzureParamsOk returns a tuple with the AzureParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureParams

`func (o *UpdateProtectedObjectsRequest) SetAzureParams(v AzureObjectProtectionUpdateRequestParams)`

SetAzureParams sets AzureParams field to given value.

### HasAzureParams

`func (o *UpdateProtectedObjectsRequest) HasAzureParams() bool`

HasAzureParams returns a boolean if a field has been set.

### GetElastifileParams

`func (o *UpdateProtectedObjectsRequest) GetElastifileParams() ElastifileObjectProtectionUpdateRequestParams`

GetElastifileParams returns the ElastifileParams field if non-nil, zero value otherwise.

### GetElastifileParamsOk

`func (o *UpdateProtectedObjectsRequest) GetElastifileParamsOk() (*ElastifileObjectProtectionUpdateRequestParams, bool)`

GetElastifileParamsOk returns a tuple with the ElastifileParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElastifileParams

`func (o *UpdateProtectedObjectsRequest) SetElastifileParams(v ElastifileObjectProtectionUpdateRequestParams)`

SetElastifileParams sets ElastifileParams field to given value.

### HasElastifileParams

`func (o *UpdateProtectedObjectsRequest) HasElastifileParams() bool`

HasElastifileParams returns a boolean if a field has been set.

### GetEnvironment

`func (o *UpdateProtectedObjectsRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *UpdateProtectedObjectsRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *UpdateProtectedObjectsRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *UpdateProtectedObjectsRequest) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *UpdateProtectedObjectsRequest) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *UpdateProtectedObjectsRequest) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetExperimentalAdapterParams

`func (o *UpdateProtectedObjectsRequest) GetExperimentalAdapterParams() ExperimentalAdapterObjectProtectionUpdateRequestParams`

GetExperimentalAdapterParams returns the ExperimentalAdapterParams field if non-nil, zero value otherwise.

### GetExperimentalAdapterParamsOk

`func (o *UpdateProtectedObjectsRequest) GetExperimentalAdapterParamsOk() (*ExperimentalAdapterObjectProtectionUpdateRequestParams, bool)`

GetExperimentalAdapterParamsOk returns a tuple with the ExperimentalAdapterParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentalAdapterParams

`func (o *UpdateProtectedObjectsRequest) SetExperimentalAdapterParams(v ExperimentalAdapterObjectProtectionUpdateRequestParams)`

SetExperimentalAdapterParams sets ExperimentalAdapterParams field to given value.

### HasExperimentalAdapterParams

`func (o *UpdateProtectedObjectsRequest) HasExperimentalAdapterParams() bool`

HasExperimentalAdapterParams returns a boolean if a field has been set.

### GetFlashbladeParams

`func (o *UpdateProtectedObjectsRequest) GetFlashbladeParams() FlashbladeObjectProtectionUpdateRequestParams`

GetFlashbladeParams returns the FlashbladeParams field if non-nil, zero value otherwise.

### GetFlashbladeParamsOk

`func (o *UpdateProtectedObjectsRequest) GetFlashbladeParamsOk() (*FlashbladeObjectProtectionUpdateRequestParams, bool)`

GetFlashbladeParamsOk returns a tuple with the FlashbladeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlashbladeParams

`func (o *UpdateProtectedObjectsRequest) SetFlashbladeParams(v FlashbladeObjectProtectionUpdateRequestParams)`

SetFlashbladeParams sets FlashbladeParams field to given value.

### HasFlashbladeParams

`func (o *UpdateProtectedObjectsRequest) HasFlashbladeParams() bool`

HasFlashbladeParams returns a boolean if a field has been set.

### GetGenericNasParams

`func (o *UpdateProtectedObjectsRequest) GetGenericNasParams() GenericNasObjectProtectionUpdateRequestParams`

GetGenericNasParams returns the GenericNasParams field if non-nil, zero value otherwise.

### GetGenericNasParamsOk

`func (o *UpdateProtectedObjectsRequest) GetGenericNasParamsOk() (*GenericNasObjectProtectionUpdateRequestParams, bool)`

GetGenericNasParamsOk returns a tuple with the GenericNasParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericNasParams

`func (o *UpdateProtectedObjectsRequest) SetGenericNasParams(v GenericNasObjectProtectionUpdateRequestParams)`

SetGenericNasParams sets GenericNasParams field to given value.

### HasGenericNasParams

`func (o *UpdateProtectedObjectsRequest) HasGenericNasParams() bool`

HasGenericNasParams returns a boolean if a field has been set.

### GetGpfsParams

`func (o *UpdateProtectedObjectsRequest) GetGpfsParams() GpfsObjectProtectionUpdateRequestParams`

GetGpfsParams returns the GpfsParams field if non-nil, zero value otherwise.

### GetGpfsParamsOk

`func (o *UpdateProtectedObjectsRequest) GetGpfsParamsOk() (*GpfsObjectProtectionUpdateRequestParams, bool)`

GetGpfsParamsOk returns a tuple with the GpfsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpfsParams

`func (o *UpdateProtectedObjectsRequest) SetGpfsParams(v GpfsObjectProtectionUpdateRequestParams)`

SetGpfsParams sets GpfsParams field to given value.

### HasGpfsParams

`func (o *UpdateProtectedObjectsRequest) HasGpfsParams() bool`

HasGpfsParams returns a boolean if a field has been set.

### GetHypervParams

`func (o *UpdateProtectedObjectsRequest) GetHypervParams() HyperVObjectProtectionUpdateRequestParams`

GetHypervParams returns the HypervParams field if non-nil, zero value otherwise.

### GetHypervParamsOk

`func (o *UpdateProtectedObjectsRequest) GetHypervParamsOk() (*HyperVObjectProtectionUpdateRequestParams, bool)`

GetHypervParamsOk returns a tuple with the HypervParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervParams

`func (o *UpdateProtectedObjectsRequest) SetHypervParams(v HyperVObjectProtectionUpdateRequestParams)`

SetHypervParams sets HypervParams field to given value.

### HasHypervParams

`func (o *UpdateProtectedObjectsRequest) HasHypervParams() bool`

HasHypervParams returns a boolean if a field has been set.

### GetIsilonParams

`func (o *UpdateProtectedObjectsRequest) GetIsilonParams() IsilonObjectProtectionUpdateRequestParams`

GetIsilonParams returns the IsilonParams field if non-nil, zero value otherwise.

### GetIsilonParamsOk

`func (o *UpdateProtectedObjectsRequest) GetIsilonParamsOk() (*IsilonObjectProtectionUpdateRequestParams, bool)`

GetIsilonParamsOk returns a tuple with the IsilonParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsilonParams

`func (o *UpdateProtectedObjectsRequest) SetIsilonParams(v IsilonObjectProtectionUpdateRequestParams)`

SetIsilonParams sets IsilonParams field to given value.

### HasIsilonParams

`func (o *UpdateProtectedObjectsRequest) HasIsilonParams() bool`

HasIsilonParams returns a boolean if a field has been set.

### GetMssqlParams

`func (o *UpdateProtectedObjectsRequest) GetMssqlParams() MssqlObjectProtectionUpdateRequestParams`

GetMssqlParams returns the MssqlParams field if non-nil, zero value otherwise.

### GetMssqlParamsOk

`func (o *UpdateProtectedObjectsRequest) GetMssqlParamsOk() (*MssqlObjectProtectionUpdateRequestParams, bool)`

GetMssqlParamsOk returns a tuple with the MssqlParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlParams

`func (o *UpdateProtectedObjectsRequest) SetMssqlParams(v MssqlObjectProtectionUpdateRequestParams)`

SetMssqlParams sets MssqlParams field to given value.

### HasMssqlParams

`func (o *UpdateProtectedObjectsRequest) HasMssqlParams() bool`

HasMssqlParams returns a boolean if a field has been set.

### GetNetappParams

`func (o *UpdateProtectedObjectsRequest) GetNetappParams() NetappObjectProtectionUpdateRequestParams`

GetNetappParams returns the NetappParams field if non-nil, zero value otherwise.

### GetNetappParamsOk

`func (o *UpdateProtectedObjectsRequest) GetNetappParamsOk() (*NetappObjectProtectionUpdateRequestParams, bool)`

GetNetappParamsOk returns a tuple with the NetappParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetappParams

`func (o *UpdateProtectedObjectsRequest) SetNetappParams(v NetappObjectProtectionUpdateRequestParams)`

SetNetappParams sets NetappParams field to given value.

### HasNetappParams

`func (o *UpdateProtectedObjectsRequest) HasNetappParams() bool`

HasNetappParams returns a boolean if a field has been set.

### GetOffice365Params

`func (o *UpdateProtectedObjectsRequest) GetOffice365Params() Office365ObjectProtectionUpdateRequestParams`

GetOffice365Params returns the Office365Params field if non-nil, zero value otherwise.

### GetOffice365ParamsOk

`func (o *UpdateProtectedObjectsRequest) GetOffice365ParamsOk() (*Office365ObjectProtectionUpdateRequestParams, bool)`

GetOffice365ParamsOk returns a tuple with the Office365Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffice365Params

`func (o *UpdateProtectedObjectsRequest) SetOffice365Params(v Office365ObjectProtectionUpdateRequestParams)`

SetOffice365Params sets Office365Params field to given value.

### HasOffice365Params

`func (o *UpdateProtectedObjectsRequest) HasOffice365Params() bool`

HasOffice365Params returns a boolean if a field has been set.

### GetOracleParams

`func (o *UpdateProtectedObjectsRequest) GetOracleParams() OracleObjectProtectionUpdateRequestParams`

GetOracleParams returns the OracleParams field if non-nil, zero value otherwise.

### GetOracleParamsOk

`func (o *UpdateProtectedObjectsRequest) GetOracleParamsOk() (*OracleObjectProtectionUpdateRequestParams, bool)`

GetOracleParamsOk returns a tuple with the OracleParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleParams

`func (o *UpdateProtectedObjectsRequest) SetOracleParams(v OracleObjectProtectionUpdateRequestParams)`

SetOracleParams sets OracleParams field to given value.

### HasOracleParams

`func (o *UpdateProtectedObjectsRequest) HasOracleParams() bool`

HasOracleParams returns a boolean if a field has been set.

### GetPhysicalParams

`func (o *UpdateProtectedObjectsRequest) GetPhysicalParams() PhysicalObjectProtectionUpdateRequestParams`

GetPhysicalParams returns the PhysicalParams field if non-nil, zero value otherwise.

### GetPhysicalParamsOk

`func (o *UpdateProtectedObjectsRequest) GetPhysicalParamsOk() (*PhysicalObjectProtectionUpdateRequestParams, bool)`

GetPhysicalParamsOk returns a tuple with the PhysicalParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalParams

`func (o *UpdateProtectedObjectsRequest) SetPhysicalParams(v PhysicalObjectProtectionUpdateRequestParams)`

SetPhysicalParams sets PhysicalParams field to given value.

### HasPhysicalParams

`func (o *UpdateProtectedObjectsRequest) HasPhysicalParams() bool`

HasPhysicalParams returns a boolean if a field has been set.

### GetSapHanaParams

`func (o *UpdateProtectedObjectsRequest) GetSapHanaParams() SapHanaObjectProtectionUpdateRequestParams`

GetSapHanaParams returns the SapHanaParams field if non-nil, zero value otherwise.

### GetSapHanaParamsOk

`func (o *UpdateProtectedObjectsRequest) GetSapHanaParamsOk() (*SapHanaObjectProtectionUpdateRequestParams, bool)`

GetSapHanaParamsOk returns a tuple with the SapHanaParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSapHanaParams

`func (o *UpdateProtectedObjectsRequest) SetSapHanaParams(v SapHanaObjectProtectionUpdateRequestParams)`

SetSapHanaParams sets SapHanaParams field to given value.

### HasSapHanaParams

`func (o *UpdateProtectedObjectsRequest) HasSapHanaParams() bool`

HasSapHanaParams returns a boolean if a field has been set.

### GetSfdcParams

`func (o *UpdateProtectedObjectsRequest) GetSfdcParams() SfdcObjectProtectionUpdateRequestParams`

GetSfdcParams returns the SfdcParams field if non-nil, zero value otherwise.

### GetSfdcParamsOk

`func (o *UpdateProtectedObjectsRequest) GetSfdcParamsOk() (*SfdcObjectProtectionUpdateRequestParams, bool)`

GetSfdcParamsOk returns a tuple with the SfdcParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfdcParams

`func (o *UpdateProtectedObjectsRequest) SetSfdcParams(v SfdcObjectProtectionUpdateRequestParams)`

SetSfdcParams sets SfdcParams field to given value.

### HasSfdcParams

`func (o *UpdateProtectedObjectsRequest) HasSfdcParams() bool`

HasSfdcParams returns a boolean if a field has been set.

### GetUdaParams

`func (o *UpdateProtectedObjectsRequest) GetUdaParams() UdaObjectProtectionUpdateRequestParams`

GetUdaParams returns the UdaParams field if non-nil, zero value otherwise.

### GetUdaParamsOk

`func (o *UpdateProtectedObjectsRequest) GetUdaParamsOk() (*UdaObjectProtectionUpdateRequestParams, bool)`

GetUdaParamsOk returns a tuple with the UdaParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdaParams

`func (o *UpdateProtectedObjectsRequest) SetUdaParams(v UdaObjectProtectionUpdateRequestParams)`

SetUdaParams sets UdaParams field to given value.

### HasUdaParams

`func (o *UpdateProtectedObjectsRequest) HasUdaParams() bool`

HasUdaParams returns a boolean if a field has been set.

### GetVmwareParams

`func (o *UpdateProtectedObjectsRequest) GetVmwareParams() VmwareObjectProtectionUpdateRequestParams`

GetVmwareParams returns the VmwareParams field if non-nil, zero value otherwise.

### GetVmwareParamsOk

`func (o *UpdateProtectedObjectsRequest) GetVmwareParamsOk() (*VmwareObjectProtectionUpdateRequestParams, bool)`

GetVmwareParamsOk returns a tuple with the VmwareParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareParams

`func (o *UpdateProtectedObjectsRequest) SetVmwareParams(v VmwareObjectProtectionUpdateRequestParams)`

SetVmwareParams sets VmwareParams field to given value.

### HasVmwareParams

`func (o *UpdateProtectedObjectsRequest) HasVmwareParams() bool`

HasVmwareParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


