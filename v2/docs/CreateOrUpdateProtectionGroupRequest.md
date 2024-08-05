# CreateOrUpdateProtectionGroupRequest

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
**AcropolisParams** | Pointer to [**AcropolisProtectionGroupParams**](AcropolisProtectionGroupParams.md) |  | [optional] 
**AdParams** | Pointer to [**ADProtectionGroupParams**](ADProtectionGroupParams.md) |  | [optional] 
**AwsParams** | Pointer to [**AwsProtectionGroupParams**](AwsProtectionGroupParams.md) |  | [optional] 
**AzureParams** | Pointer to [**AzureProtectionGroupParams**](AzureProtectionGroupParams.md) |  | [optional] 
**CassandraParams** | Pointer to [**CassandraProtectionGroupParams**](CassandraProtectionGroupParams.md) |  | [optional] 
**CouchbaseParams** | Pointer to [**NoSqlProtectionGroupParams**](NoSqlProtectionGroupParams.md) |  | [optional] 
**ElastifileParams** | Pointer to [**NullableElastifileProtectionGroupParams**](ElastifileProtectionGroupParams.md) |  | [optional] 
**ExchangeParams** | Pointer to [**ExchangeProtectionGroupParams**](ExchangeProtectionGroupParams.md) |  | [optional] 
**ExperimentalAdapterParams** | Pointer to [**ExperimentalAdapterProtectionGroupParams**](ExperimentalAdapterProtectionGroupParams.md) |  | [optional] 
**FlashbladeParams** | Pointer to [**NullableFlashbladeProtectionGroupParams**](FlashbladeProtectionGroupParams.md) |  | [optional] 
**GcpParams** | Pointer to [**GcpProtectionGroupParams**](GcpProtectionGroupParams.md) |  | [optional] 
**GenericNasParams** | Pointer to [**GenericNasProtectionGroupParams**](GenericNasProtectionGroupParams.md) |  | [optional] 
**GpfsParams** | Pointer to [**NullableGpfsProtectionGroupParams**](GpfsProtectionGroupParams.md) |  | [optional] 
**HbaseParams** | Pointer to [**NoSqlProtectionGroupParams**](NoSqlProtectionGroupParams.md) |  | [optional] 
**HdfsParams** | Pointer to [**HdfsProtectionGroupParams**](HdfsProtectionGroupParams.md) |  | [optional] 
**HiveParams** | Pointer to [**NoSqlProtectionGroupParams**](NoSqlProtectionGroupParams.md) |  | [optional] 
**HypervParams** | Pointer to [**HyperVProtectionGroupParams**](HyperVProtectionGroupParams.md) |  | [optional] 
**IbmFlashSystemParams** | Pointer to [**NullableIbmFlashSystemProtectionGroupParams**](IbmFlashSystemProtectionGroupParams.md) |  | [optional] 
**IsilonParams** | Pointer to [**NullableIsilonProtectionGroupParams**](IsilonProtectionGroupParams.md) |  | [optional] 
**KubernetesParams** | Pointer to [**KubernetesProtectionGroupParams**](KubernetesProtectionGroupParams.md) |  | [optional] 
**KvmParams** | Pointer to [**NullableKvmProtectionGroupParams**](KvmProtectionGroupParams.md) |  | [optional] 
**MongodbParams** | Pointer to [**MongoDBProtectionGroupParams**](MongoDBProtectionGroupParams.md) |  | [optional] 
**MssqlParams** | Pointer to [**MSSQLProtectionGroupParams**](MSSQLProtectionGroupParams.md) |  | [optional] 
**NetappParams** | Pointer to [**NullableNetappProtectionGroupParams**](NetappProtectionGroupParams.md) |  | [optional] 
**NimbleParams** | Pointer to [**NullableNimbleProtectionGroupParams**](NimbleProtectionGroupParams.md) |  | [optional] 
**Office365Params** | Pointer to [**Office365ProtectionGroupParams**](Office365ProtectionGroupParams.md) |  | [optional] 
**OracleParams** | Pointer to [**OracleProtectionGroupParams**](OracleProtectionGroupParams.md) |  | [optional] 
**PhysicalParams** | Pointer to [**PhysicalProtectionGroupParams**](PhysicalProtectionGroupParams.md) |  | [optional] 
**PureParams** | Pointer to [**NullablePureProtectionGroupParams**](PureProtectionGroupParams.md) |  | [optional] 
**RemoteAdapterParams** | Pointer to [**NullableRemoteAdapterProtectionGroupParams**](RemoteAdapterProtectionGroupParams.md) |  | [optional] 
**SapHanaParams** | Pointer to [**SapHanaProtectionGroupParams**](SapHanaProtectionGroupParams.md) |  | [optional] 
**SfdcParams** | Pointer to [**SfdcProtectionGroupParams**](SfdcProtectionGroupParams.md) |  | [optional] 
**UdaParams** | Pointer to [**UdaProtectionGroupParams**](UdaProtectionGroupParams.md) |  | [optional] 
**ViewParams** | Pointer to [**NullableViewProtectionGroupParams**](ViewProtectionGroupParams.md) |  | [optional] 
**VmwareParams** | Pointer to [**VmwareProtectionGroupParams**](VmwareProtectionGroupParams.md) |  | [optional] 

## Methods

### NewCreateOrUpdateProtectionGroupRequest

`func NewCreateOrUpdateProtectionGroupRequest(environment NullableString, name NullableString, policyId NullableString, ) *CreateOrUpdateProtectionGroupRequest`

NewCreateOrUpdateProtectionGroupRequest instantiates a new CreateOrUpdateProtectionGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrUpdateProtectionGroupRequestWithDefaults

`func NewCreateOrUpdateProtectionGroupRequestWithDefaults() *CreateOrUpdateProtectionGroupRequest`

NewCreateOrUpdateProtectionGroupRequestWithDefaults instantiates a new CreateOrUpdateProtectionGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbortInBlackouts

`func (o *CreateOrUpdateProtectionGroupRequest) GetAbortInBlackouts() bool`

GetAbortInBlackouts returns the AbortInBlackouts field if non-nil, zero value otherwise.

### GetAbortInBlackoutsOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetAbortInBlackoutsOk() (*bool, bool)`

GetAbortInBlackoutsOk returns a tuple with the AbortInBlackouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbortInBlackouts

`func (o *CreateOrUpdateProtectionGroupRequest) SetAbortInBlackouts(v bool)`

SetAbortInBlackouts sets AbortInBlackouts field to given value.

### HasAbortInBlackouts

`func (o *CreateOrUpdateProtectionGroupRequest) HasAbortInBlackouts() bool`

HasAbortInBlackouts returns a boolean if a field has been set.

### SetAbortInBlackoutsNil

`func (o *CreateOrUpdateProtectionGroupRequest) SetAbortInBlackoutsNil(b bool)`

 SetAbortInBlackoutsNil sets the value for AbortInBlackouts to be an explicit nil

### UnsetAbortInBlackouts
`func (o *CreateOrUpdateProtectionGroupRequest) UnsetAbortInBlackouts()`

UnsetAbortInBlackouts ensures that no value is present for AbortInBlackouts, not even an explicit nil
### GetAdvancedConfigs

`func (o *CreateOrUpdateProtectionGroupRequest) GetAdvancedConfigs() []KeyValuePair`

GetAdvancedConfigs returns the AdvancedConfigs field if non-nil, zero value otherwise.

### GetAdvancedConfigsOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetAdvancedConfigsOk() (*[]KeyValuePair, bool)`

GetAdvancedConfigsOk returns a tuple with the AdvancedConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedConfigs

`func (o *CreateOrUpdateProtectionGroupRequest) SetAdvancedConfigs(v []KeyValuePair)`

SetAdvancedConfigs sets AdvancedConfigs field to given value.

### HasAdvancedConfigs

`func (o *CreateOrUpdateProtectionGroupRequest) HasAdvancedConfigs() bool`

HasAdvancedConfigs returns a boolean if a field has been set.

### SetAdvancedConfigsNil

`func (o *CreateOrUpdateProtectionGroupRequest) SetAdvancedConfigsNil(b bool)`

 SetAdvancedConfigsNil sets the value for AdvancedConfigs to be an explicit nil

### UnsetAdvancedConfigs
`func (o *CreateOrUpdateProtectionGroupRequest) UnsetAdvancedConfigs()`

UnsetAdvancedConfigs ensures that no value is present for AdvancedConfigs, not even an explicit nil
### GetAlertPolicy

`func (o *CreateOrUpdateProtectionGroupRequest) GetAlertPolicy() ProtectionGroupAlertingPolicy`

GetAlertPolicy returns the AlertPolicy field if non-nil, zero value otherwise.

### GetAlertPolicyOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetAlertPolicyOk() (*ProtectionGroupAlertingPolicy, bool)`

GetAlertPolicyOk returns a tuple with the AlertPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertPolicy

`func (o *CreateOrUpdateProtectionGroupRequest) SetAlertPolicy(v ProtectionGroupAlertingPolicy)`

SetAlertPolicy sets AlertPolicy field to given value.

### HasAlertPolicy

`func (o *CreateOrUpdateProtectionGroupRequest) HasAlertPolicy() bool`

HasAlertPolicy returns a boolean if a field has been set.

### GetDescription

`func (o *CreateOrUpdateProtectionGroupRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateOrUpdateProtectionGroupRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateOrUpdateProtectionGroupRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateOrUpdateProtectionGroupRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateOrUpdateProtectionGroupRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEndTimeUsecs

`func (o *CreateOrUpdateProtectionGroupRequest) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *CreateOrUpdateProtectionGroupRequest) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *CreateOrUpdateProtectionGroupRequest) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *CreateOrUpdateProtectionGroupRequest) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *CreateOrUpdateProtectionGroupRequest) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetEnvironment

`func (o *CreateOrUpdateProtectionGroupRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CreateOrUpdateProtectionGroupRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### SetEnvironmentNil

`func (o *CreateOrUpdateProtectionGroupRequest) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *CreateOrUpdateProtectionGroupRequest) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetIsPaused

`func (o *CreateOrUpdateProtectionGroupRequest) GetIsPaused() bool`

GetIsPaused returns the IsPaused field if non-nil, zero value otherwise.

### GetIsPausedOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetIsPausedOk() (*bool, bool)`

GetIsPausedOk returns a tuple with the IsPaused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPaused

`func (o *CreateOrUpdateProtectionGroupRequest) SetIsPaused(v bool)`

SetIsPaused sets IsPaused field to given value.

### HasIsPaused

`func (o *CreateOrUpdateProtectionGroupRequest) HasIsPaused() bool`

HasIsPaused returns a boolean if a field has been set.

### SetIsPausedNil

`func (o *CreateOrUpdateProtectionGroupRequest) SetIsPausedNil(b bool)`

 SetIsPausedNil sets the value for IsPaused to be an explicit nil

### UnsetIsPaused
`func (o *CreateOrUpdateProtectionGroupRequest) UnsetIsPaused()`

UnsetIsPaused ensures that no value is present for IsPaused, not even an explicit nil
### GetLastModifiedTimestampUsecs

`func (o *CreateOrUpdateProtectionGroupRequest) GetLastModifiedTimestampUsecs() int64`

GetLastModifiedTimestampUsecs returns the LastModifiedTimestampUsecs field if non-nil, zero value otherwise.

### GetLastModifiedTimestampUsecsOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetLastModifiedTimestampUsecsOk() (*int64, bool)`

GetLastModifiedTimestampUsecsOk returns a tuple with the LastModifiedTimestampUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTimestampUsecs

`func (o *CreateOrUpdateProtectionGroupRequest) SetLastModifiedTimestampUsecs(v int64)`

SetLastModifiedTimestampUsecs sets LastModifiedTimestampUsecs field to given value.

### HasLastModifiedTimestampUsecs

`func (o *CreateOrUpdateProtectionGroupRequest) HasLastModifiedTimestampUsecs() bool`

HasLastModifiedTimestampUsecs returns a boolean if a field has been set.

### SetLastModifiedTimestampUsecsNil

`func (o *CreateOrUpdateProtectionGroupRequest) SetLastModifiedTimestampUsecsNil(b bool)`

 SetLastModifiedTimestampUsecsNil sets the value for LastModifiedTimestampUsecs to be an explicit nil

### UnsetLastModifiedTimestampUsecs
`func (o *CreateOrUpdateProtectionGroupRequest) UnsetLastModifiedTimestampUsecs()`

UnsetLastModifiedTimestampUsecs ensures that no value is present for LastModifiedTimestampUsecs, not even an explicit nil
### GetName

`func (o *CreateOrUpdateProtectionGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOrUpdateProtectionGroupRequest) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CreateOrUpdateProtectionGroupRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateOrUpdateProtectionGroupRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPauseInBlackouts

`func (o *CreateOrUpdateProtectionGroupRequest) GetPauseInBlackouts() bool`

GetPauseInBlackouts returns the PauseInBlackouts field if non-nil, zero value otherwise.

### GetPauseInBlackoutsOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetPauseInBlackoutsOk() (*bool, bool)`

GetPauseInBlackoutsOk returns a tuple with the PauseInBlackouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseInBlackouts

`func (o *CreateOrUpdateProtectionGroupRequest) SetPauseInBlackouts(v bool)`

SetPauseInBlackouts sets PauseInBlackouts field to given value.

### HasPauseInBlackouts

`func (o *CreateOrUpdateProtectionGroupRequest) HasPauseInBlackouts() bool`

HasPauseInBlackouts returns a boolean if a field has been set.

### SetPauseInBlackoutsNil

`func (o *CreateOrUpdateProtectionGroupRequest) SetPauseInBlackoutsNil(b bool)`

 SetPauseInBlackoutsNil sets the value for PauseInBlackouts to be an explicit nil

### UnsetPauseInBlackouts
`func (o *CreateOrUpdateProtectionGroupRequest) UnsetPauseInBlackouts()`

UnsetPauseInBlackouts ensures that no value is present for PauseInBlackouts, not even an explicit nil
### GetPolicyId

`func (o *CreateOrUpdateProtectionGroupRequest) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *CreateOrUpdateProtectionGroupRequest) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.


### SetPolicyIdNil

`func (o *CreateOrUpdateProtectionGroupRequest) SetPolicyIdNil(b bool)`

 SetPolicyIdNil sets the value for PolicyId to be an explicit nil

### UnsetPolicyId
`func (o *CreateOrUpdateProtectionGroupRequest) UnsetPolicyId()`

UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
### GetPriority

`func (o *CreateOrUpdateProtectionGroupRequest) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CreateOrUpdateProtectionGroupRequest) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CreateOrUpdateProtectionGroupRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *CreateOrUpdateProtectionGroupRequest) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *CreateOrUpdateProtectionGroupRequest) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetQosPolicy

`func (o *CreateOrUpdateProtectionGroupRequest) GetQosPolicy() string`

GetQosPolicy returns the QosPolicy field if non-nil, zero value otherwise.

### GetQosPolicyOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetQosPolicyOk() (*string, bool)`

GetQosPolicyOk returns a tuple with the QosPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosPolicy

`func (o *CreateOrUpdateProtectionGroupRequest) SetQosPolicy(v string)`

SetQosPolicy sets QosPolicy field to given value.

### HasQosPolicy

`func (o *CreateOrUpdateProtectionGroupRequest) HasQosPolicy() bool`

HasQosPolicy returns a boolean if a field has been set.

### SetQosPolicyNil

`func (o *CreateOrUpdateProtectionGroupRequest) SetQosPolicyNil(b bool)`

 SetQosPolicyNil sets the value for QosPolicy to be an explicit nil

### UnsetQosPolicy
`func (o *CreateOrUpdateProtectionGroupRequest) UnsetQosPolicy()`

UnsetQosPolicy ensures that no value is present for QosPolicy, not even an explicit nil
### GetSla

`func (o *CreateOrUpdateProtectionGroupRequest) GetSla() []SlaRule`

GetSla returns the Sla field if non-nil, zero value otherwise.

### GetSlaOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetSlaOk() (*[]SlaRule, bool)`

GetSlaOk returns a tuple with the Sla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSla

`func (o *CreateOrUpdateProtectionGroupRequest) SetSla(v []SlaRule)`

SetSla sets Sla field to given value.

### HasSla

`func (o *CreateOrUpdateProtectionGroupRequest) HasSla() bool`

HasSla returns a boolean if a field has been set.

### SetSlaNil

`func (o *CreateOrUpdateProtectionGroupRequest) SetSlaNil(b bool)`

 SetSlaNil sets the value for Sla to be an explicit nil

### UnsetSla
`func (o *CreateOrUpdateProtectionGroupRequest) UnsetSla()`

UnsetSla ensures that no value is present for Sla, not even an explicit nil
### GetStartTime

`func (o *CreateOrUpdateProtectionGroupRequest) GetStartTime() TimeOfDay`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetStartTimeOk() (*TimeOfDay, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *CreateOrUpdateProtectionGroupRequest) SetStartTime(v TimeOfDay)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *CreateOrUpdateProtectionGroupRequest) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStorageDomainId

`func (o *CreateOrUpdateProtectionGroupRequest) GetStorageDomainId() int64`

GetStorageDomainId returns the StorageDomainId field if non-nil, zero value otherwise.

### GetStorageDomainIdOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetStorageDomainIdOk() (*int64, bool)`

GetStorageDomainIdOk returns a tuple with the StorageDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainId

`func (o *CreateOrUpdateProtectionGroupRequest) SetStorageDomainId(v int64)`

SetStorageDomainId sets StorageDomainId field to given value.

### HasStorageDomainId

`func (o *CreateOrUpdateProtectionGroupRequest) HasStorageDomainId() bool`

HasStorageDomainId returns a boolean if a field has been set.

### SetStorageDomainIdNil

`func (o *CreateOrUpdateProtectionGroupRequest) SetStorageDomainIdNil(b bool)`

 SetStorageDomainIdNil sets the value for StorageDomainId to be an explicit nil

### UnsetStorageDomainId
`func (o *CreateOrUpdateProtectionGroupRequest) UnsetStorageDomainId()`

UnsetStorageDomainId ensures that no value is present for StorageDomainId, not even an explicit nil
### GetAcropolisParams

`func (o *CreateOrUpdateProtectionGroupRequest) GetAcropolisParams() AcropolisProtectionGroupParams`

GetAcropolisParams returns the AcropolisParams field if non-nil, zero value otherwise.

### GetAcropolisParamsOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetAcropolisParamsOk() (*AcropolisProtectionGroupParams, bool)`

GetAcropolisParamsOk returns a tuple with the AcropolisParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcropolisParams

`func (o *CreateOrUpdateProtectionGroupRequest) SetAcropolisParams(v AcropolisProtectionGroupParams)`

SetAcropolisParams sets AcropolisParams field to given value.

### HasAcropolisParams

`func (o *CreateOrUpdateProtectionGroupRequest) HasAcropolisParams() bool`

HasAcropolisParams returns a boolean if a field has been set.

### GetAdParams

`func (o *CreateOrUpdateProtectionGroupRequest) GetAdParams() ADProtectionGroupParams`

GetAdParams returns the AdParams field if non-nil, zero value otherwise.

### GetAdParamsOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetAdParamsOk() (*ADProtectionGroupParams, bool)`

GetAdParamsOk returns a tuple with the AdParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdParams

`func (o *CreateOrUpdateProtectionGroupRequest) SetAdParams(v ADProtectionGroupParams)`

SetAdParams sets AdParams field to given value.

### HasAdParams

`func (o *CreateOrUpdateProtectionGroupRequest) HasAdParams() bool`

HasAdParams returns a boolean if a field has been set.

### GetAwsParams

`func (o *CreateOrUpdateProtectionGroupRequest) GetAwsParams() AwsProtectionGroupParams`

GetAwsParams returns the AwsParams field if non-nil, zero value otherwise.

### GetAwsParamsOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetAwsParamsOk() (*AwsProtectionGroupParams, bool)`

GetAwsParamsOk returns a tuple with the AwsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsParams

`func (o *CreateOrUpdateProtectionGroupRequest) SetAwsParams(v AwsProtectionGroupParams)`

SetAwsParams sets AwsParams field to given value.

### HasAwsParams

`func (o *CreateOrUpdateProtectionGroupRequest) HasAwsParams() bool`

HasAwsParams returns a boolean if a field has been set.

### GetAzureParams

`func (o *CreateOrUpdateProtectionGroupRequest) GetAzureParams() AzureProtectionGroupParams`

GetAzureParams returns the AzureParams field if non-nil, zero value otherwise.

### GetAzureParamsOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetAzureParamsOk() (*AzureProtectionGroupParams, bool)`

GetAzureParamsOk returns a tuple with the AzureParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureParams

`func (o *CreateOrUpdateProtectionGroupRequest) SetAzureParams(v AzureProtectionGroupParams)`

SetAzureParams sets AzureParams field to given value.

### HasAzureParams

`func (o *CreateOrUpdateProtectionGroupRequest) HasAzureParams() bool`

HasAzureParams returns a boolean if a field has been set.

### GetCassandraParams

`func (o *CreateOrUpdateProtectionGroupRequest) GetCassandraParams() CassandraProtectionGroupParams`

GetCassandraParams returns the CassandraParams field if non-nil, zero value otherwise.

### GetCassandraParamsOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetCassandraParamsOk() (*CassandraProtectionGroupParams, bool)`

GetCassandraParamsOk returns a tuple with the CassandraParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraParams

`func (o *CreateOrUpdateProtectionGroupRequest) SetCassandraParams(v CassandraProtectionGroupParams)`

SetCassandraParams sets CassandraParams field to given value.

### HasCassandraParams

`func (o *CreateOrUpdateProtectionGroupRequest) HasCassandraParams() bool`

HasCassandraParams returns a boolean if a field has been set.

### GetCouchbaseParams

`func (o *CreateOrUpdateProtectionGroupRequest) GetCouchbaseParams() NoSqlProtectionGroupParams`

GetCouchbaseParams returns the CouchbaseParams field if non-nil, zero value otherwise.

### GetCouchbaseParamsOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetCouchbaseParamsOk() (*NoSqlProtectionGroupParams, bool)`

GetCouchbaseParamsOk returns a tuple with the CouchbaseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouchbaseParams

`func (o *CreateOrUpdateProtectionGroupRequest) SetCouchbaseParams(v NoSqlProtectionGroupParams)`

SetCouchbaseParams sets CouchbaseParams field to given value.

### HasCouchbaseParams

`func (o *CreateOrUpdateProtectionGroupRequest) HasCouchbaseParams() bool`

HasCouchbaseParams returns a boolean if a field has been set.

### GetElastifileParams

`func (o *CreateOrUpdateProtectionGroupRequest) GetElastifileParams() ElastifileProtectionGroupParams`

GetElastifileParams returns the ElastifileParams field if non-nil, zero value otherwise.

### GetElastifileParamsOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetElastifileParamsOk() (*ElastifileProtectionGroupParams, bool)`

GetElastifileParamsOk returns a tuple with the ElastifileParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElastifileParams

`func (o *CreateOrUpdateProtectionGroupRequest) SetElastifileParams(v ElastifileProtectionGroupParams)`

SetElastifileParams sets ElastifileParams field to given value.

### HasElastifileParams

`func (o *CreateOrUpdateProtectionGroupRequest) HasElastifileParams() bool`

HasElastifileParams returns a boolean if a field has been set.

### SetElastifileParamsNil

`func (o *CreateOrUpdateProtectionGroupRequest) SetElastifileParamsNil(b bool)`

 SetElastifileParamsNil sets the value for ElastifileParams to be an explicit nil

### UnsetElastifileParams
`func (o *CreateOrUpdateProtectionGroupRequest) UnsetElastifileParams()`

UnsetElastifileParams ensures that no value is present for ElastifileParams, not even an explicit nil
### GetExchangeParams

`func (o *CreateOrUpdateProtectionGroupRequest) GetExchangeParams() ExchangeProtectionGroupParams`

GetExchangeParams returns the ExchangeParams field if non-nil, zero value otherwise.

### GetExchangeParamsOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetExchangeParamsOk() (*ExchangeProtectionGroupParams, bool)`

GetExchangeParamsOk returns a tuple with the ExchangeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeParams

`func (o *CreateOrUpdateProtectionGroupRequest) SetExchangeParams(v ExchangeProtectionGroupParams)`

SetExchangeParams sets ExchangeParams field to given value.

### HasExchangeParams

`func (o *CreateOrUpdateProtectionGroupRequest) HasExchangeParams() bool`

HasExchangeParams returns a boolean if a field has been set.

### GetExperimentalAdapterParams

`func (o *CreateOrUpdateProtectionGroupRequest) GetExperimentalAdapterParams() ExperimentalAdapterProtectionGroupParams`

GetExperimentalAdapterParams returns the ExperimentalAdapterParams field if non-nil, zero value otherwise.

### GetExperimentalAdapterParamsOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetExperimentalAdapterParamsOk() (*ExperimentalAdapterProtectionGroupParams, bool)`

GetExperimentalAdapterParamsOk returns a tuple with the ExperimentalAdapterParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentalAdapterParams

`func (o *CreateOrUpdateProtectionGroupRequest) SetExperimentalAdapterParams(v ExperimentalAdapterProtectionGroupParams)`

SetExperimentalAdapterParams sets ExperimentalAdapterParams field to given value.

### HasExperimentalAdapterParams

`func (o *CreateOrUpdateProtectionGroupRequest) HasExperimentalAdapterParams() bool`

HasExperimentalAdapterParams returns a boolean if a field has been set.

### GetFlashbladeParams

`func (o *CreateOrUpdateProtectionGroupRequest) GetFlashbladeParams() FlashbladeProtectionGroupParams`

GetFlashbladeParams returns the FlashbladeParams field if non-nil, zero value otherwise.

### GetFlashbladeParamsOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetFlashbladeParamsOk() (*FlashbladeProtectionGroupParams, bool)`

GetFlashbladeParamsOk returns a tuple with the FlashbladeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlashbladeParams

`func (o *CreateOrUpdateProtectionGroupRequest) SetFlashbladeParams(v FlashbladeProtectionGroupParams)`

SetFlashbladeParams sets FlashbladeParams field to given value.

### HasFlashbladeParams

`func (o *CreateOrUpdateProtectionGroupRequest) HasFlashbladeParams() bool`

HasFlashbladeParams returns a boolean if a field has been set.

### SetFlashbladeParamsNil

`func (o *CreateOrUpdateProtectionGroupRequest) SetFlashbladeParamsNil(b bool)`

 SetFlashbladeParamsNil sets the value for FlashbladeParams to be an explicit nil

### UnsetFlashbladeParams
`func (o *CreateOrUpdateProtectionGroupRequest) UnsetFlashbladeParams()`

UnsetFlashbladeParams ensures that no value is present for FlashbladeParams, not even an explicit nil
### GetGcpParams

`func (o *CreateOrUpdateProtectionGroupRequest) GetGcpParams() GcpProtectionGroupParams`

GetGcpParams returns the GcpParams field if non-nil, zero value otherwise.

### GetGcpParamsOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetGcpParamsOk() (*GcpProtectionGroupParams, bool)`

GetGcpParamsOk returns a tuple with the GcpParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpParams

`func (o *CreateOrUpdateProtectionGroupRequest) SetGcpParams(v GcpProtectionGroupParams)`

SetGcpParams sets GcpParams field to given value.

### HasGcpParams

`func (o *CreateOrUpdateProtectionGroupRequest) HasGcpParams() bool`

HasGcpParams returns a boolean if a field has been set.

### GetGenericNasParams

`func (o *CreateOrUpdateProtectionGroupRequest) GetGenericNasParams() GenericNasProtectionGroupParams`

GetGenericNasParams returns the GenericNasParams field if non-nil, zero value otherwise.

### GetGenericNasParamsOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetGenericNasParamsOk() (*GenericNasProtectionGroupParams, bool)`

GetGenericNasParamsOk returns a tuple with the GenericNasParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericNasParams

`func (o *CreateOrUpdateProtectionGroupRequest) SetGenericNasParams(v GenericNasProtectionGroupParams)`

SetGenericNasParams sets GenericNasParams field to given value.

### HasGenericNasParams

`func (o *CreateOrUpdateProtectionGroupRequest) HasGenericNasParams() bool`

HasGenericNasParams returns a boolean if a field has been set.

### GetGpfsParams

`func (o *CreateOrUpdateProtectionGroupRequest) GetGpfsParams() GpfsProtectionGroupParams`

GetGpfsParams returns the GpfsParams field if non-nil, zero value otherwise.

### GetGpfsParamsOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetGpfsParamsOk() (*GpfsProtectionGroupParams, bool)`

GetGpfsParamsOk returns a tuple with the GpfsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpfsParams

`func (o *CreateOrUpdateProtectionGroupRequest) SetGpfsParams(v GpfsProtectionGroupParams)`

SetGpfsParams sets GpfsParams field to given value.

### HasGpfsParams

`func (o *CreateOrUpdateProtectionGroupRequest) HasGpfsParams() bool`

HasGpfsParams returns a boolean if a field has been set.

### SetGpfsParamsNil

`func (o *CreateOrUpdateProtectionGroupRequest) SetGpfsParamsNil(b bool)`

 SetGpfsParamsNil sets the value for GpfsParams to be an explicit nil

### UnsetGpfsParams
`func (o *CreateOrUpdateProtectionGroupRequest) UnsetGpfsParams()`

UnsetGpfsParams ensures that no value is present for GpfsParams, not even an explicit nil
### GetHbaseParams

`func (o *CreateOrUpdateProtectionGroupRequest) GetHbaseParams() NoSqlProtectionGroupParams`

GetHbaseParams returns the HbaseParams field if non-nil, zero value otherwise.

### GetHbaseParamsOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetHbaseParamsOk() (*NoSqlProtectionGroupParams, bool)`

GetHbaseParamsOk returns a tuple with the HbaseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHbaseParams

`func (o *CreateOrUpdateProtectionGroupRequest) SetHbaseParams(v NoSqlProtectionGroupParams)`

SetHbaseParams sets HbaseParams field to given value.

### HasHbaseParams

`func (o *CreateOrUpdateProtectionGroupRequest) HasHbaseParams() bool`

HasHbaseParams returns a boolean if a field has been set.

### GetHdfsParams

`func (o *CreateOrUpdateProtectionGroupRequest) GetHdfsParams() HdfsProtectionGroupParams`

GetHdfsParams returns the HdfsParams field if non-nil, zero value otherwise.

### GetHdfsParamsOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetHdfsParamsOk() (*HdfsProtectionGroupParams, bool)`

GetHdfsParamsOk returns a tuple with the HdfsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsParams

`func (o *CreateOrUpdateProtectionGroupRequest) SetHdfsParams(v HdfsProtectionGroupParams)`

SetHdfsParams sets HdfsParams field to given value.

### HasHdfsParams

`func (o *CreateOrUpdateProtectionGroupRequest) HasHdfsParams() bool`

HasHdfsParams returns a boolean if a field has been set.

### GetHiveParams

`func (o *CreateOrUpdateProtectionGroupRequest) GetHiveParams() NoSqlProtectionGroupParams`

GetHiveParams returns the HiveParams field if non-nil, zero value otherwise.

### GetHiveParamsOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetHiveParamsOk() (*NoSqlProtectionGroupParams, bool)`

GetHiveParamsOk returns a tuple with the HiveParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiveParams

`func (o *CreateOrUpdateProtectionGroupRequest) SetHiveParams(v NoSqlProtectionGroupParams)`

SetHiveParams sets HiveParams field to given value.

### HasHiveParams

`func (o *CreateOrUpdateProtectionGroupRequest) HasHiveParams() bool`

HasHiveParams returns a boolean if a field has been set.

### GetHypervParams

`func (o *CreateOrUpdateProtectionGroupRequest) GetHypervParams() HyperVProtectionGroupParams`

GetHypervParams returns the HypervParams field if non-nil, zero value otherwise.

### GetHypervParamsOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetHypervParamsOk() (*HyperVProtectionGroupParams, bool)`

GetHypervParamsOk returns a tuple with the HypervParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervParams

`func (o *CreateOrUpdateProtectionGroupRequest) SetHypervParams(v HyperVProtectionGroupParams)`

SetHypervParams sets HypervParams field to given value.

### HasHypervParams

`func (o *CreateOrUpdateProtectionGroupRequest) HasHypervParams() bool`

HasHypervParams returns a boolean if a field has been set.

### GetIbmFlashSystemParams

`func (o *CreateOrUpdateProtectionGroupRequest) GetIbmFlashSystemParams() IbmFlashSystemProtectionGroupParams`

GetIbmFlashSystemParams returns the IbmFlashSystemParams field if non-nil, zero value otherwise.

### GetIbmFlashSystemParamsOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetIbmFlashSystemParamsOk() (*IbmFlashSystemProtectionGroupParams, bool)`

GetIbmFlashSystemParamsOk returns a tuple with the IbmFlashSystemParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbmFlashSystemParams

`func (o *CreateOrUpdateProtectionGroupRequest) SetIbmFlashSystemParams(v IbmFlashSystemProtectionGroupParams)`

SetIbmFlashSystemParams sets IbmFlashSystemParams field to given value.

### HasIbmFlashSystemParams

`func (o *CreateOrUpdateProtectionGroupRequest) HasIbmFlashSystemParams() bool`

HasIbmFlashSystemParams returns a boolean if a field has been set.

### SetIbmFlashSystemParamsNil

`func (o *CreateOrUpdateProtectionGroupRequest) SetIbmFlashSystemParamsNil(b bool)`

 SetIbmFlashSystemParamsNil sets the value for IbmFlashSystemParams to be an explicit nil

### UnsetIbmFlashSystemParams
`func (o *CreateOrUpdateProtectionGroupRequest) UnsetIbmFlashSystemParams()`

UnsetIbmFlashSystemParams ensures that no value is present for IbmFlashSystemParams, not even an explicit nil
### GetIsilonParams

`func (o *CreateOrUpdateProtectionGroupRequest) GetIsilonParams() IsilonProtectionGroupParams`

GetIsilonParams returns the IsilonParams field if non-nil, zero value otherwise.

### GetIsilonParamsOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetIsilonParamsOk() (*IsilonProtectionGroupParams, bool)`

GetIsilonParamsOk returns a tuple with the IsilonParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsilonParams

`func (o *CreateOrUpdateProtectionGroupRequest) SetIsilonParams(v IsilonProtectionGroupParams)`

SetIsilonParams sets IsilonParams field to given value.

### HasIsilonParams

`func (o *CreateOrUpdateProtectionGroupRequest) HasIsilonParams() bool`

HasIsilonParams returns a boolean if a field has been set.

### SetIsilonParamsNil

`func (o *CreateOrUpdateProtectionGroupRequest) SetIsilonParamsNil(b bool)`

 SetIsilonParamsNil sets the value for IsilonParams to be an explicit nil

### UnsetIsilonParams
`func (o *CreateOrUpdateProtectionGroupRequest) UnsetIsilonParams()`

UnsetIsilonParams ensures that no value is present for IsilonParams, not even an explicit nil
### GetKubernetesParams

`func (o *CreateOrUpdateProtectionGroupRequest) GetKubernetesParams() KubernetesProtectionGroupParams`

GetKubernetesParams returns the KubernetesParams field if non-nil, zero value otherwise.

### GetKubernetesParamsOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetKubernetesParamsOk() (*KubernetesProtectionGroupParams, bool)`

GetKubernetesParamsOk returns a tuple with the KubernetesParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesParams

`func (o *CreateOrUpdateProtectionGroupRequest) SetKubernetesParams(v KubernetesProtectionGroupParams)`

SetKubernetesParams sets KubernetesParams field to given value.

### HasKubernetesParams

`func (o *CreateOrUpdateProtectionGroupRequest) HasKubernetesParams() bool`

HasKubernetesParams returns a boolean if a field has been set.

### GetKvmParams

`func (o *CreateOrUpdateProtectionGroupRequest) GetKvmParams() KvmProtectionGroupParams`

GetKvmParams returns the KvmParams field if non-nil, zero value otherwise.

### GetKvmParamsOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetKvmParamsOk() (*KvmProtectionGroupParams, bool)`

GetKvmParamsOk returns a tuple with the KvmParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvmParams

`func (o *CreateOrUpdateProtectionGroupRequest) SetKvmParams(v KvmProtectionGroupParams)`

SetKvmParams sets KvmParams field to given value.

### HasKvmParams

`func (o *CreateOrUpdateProtectionGroupRequest) HasKvmParams() bool`

HasKvmParams returns a boolean if a field has been set.

### SetKvmParamsNil

`func (o *CreateOrUpdateProtectionGroupRequest) SetKvmParamsNil(b bool)`

 SetKvmParamsNil sets the value for KvmParams to be an explicit nil

### UnsetKvmParams
`func (o *CreateOrUpdateProtectionGroupRequest) UnsetKvmParams()`

UnsetKvmParams ensures that no value is present for KvmParams, not even an explicit nil
### GetMongodbParams

`func (o *CreateOrUpdateProtectionGroupRequest) GetMongodbParams() MongoDBProtectionGroupParams`

GetMongodbParams returns the MongodbParams field if non-nil, zero value otherwise.

### GetMongodbParamsOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetMongodbParamsOk() (*MongoDBProtectionGroupParams, bool)`

GetMongodbParamsOk returns a tuple with the MongodbParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbParams

`func (o *CreateOrUpdateProtectionGroupRequest) SetMongodbParams(v MongoDBProtectionGroupParams)`

SetMongodbParams sets MongodbParams field to given value.

### HasMongodbParams

`func (o *CreateOrUpdateProtectionGroupRequest) HasMongodbParams() bool`

HasMongodbParams returns a boolean if a field has been set.

### GetMssqlParams

`func (o *CreateOrUpdateProtectionGroupRequest) GetMssqlParams() MSSQLProtectionGroupParams`

GetMssqlParams returns the MssqlParams field if non-nil, zero value otherwise.

### GetMssqlParamsOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetMssqlParamsOk() (*MSSQLProtectionGroupParams, bool)`

GetMssqlParamsOk returns a tuple with the MssqlParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlParams

`func (o *CreateOrUpdateProtectionGroupRequest) SetMssqlParams(v MSSQLProtectionGroupParams)`

SetMssqlParams sets MssqlParams field to given value.

### HasMssqlParams

`func (o *CreateOrUpdateProtectionGroupRequest) HasMssqlParams() bool`

HasMssqlParams returns a boolean if a field has been set.

### GetNetappParams

`func (o *CreateOrUpdateProtectionGroupRequest) GetNetappParams() NetappProtectionGroupParams`

GetNetappParams returns the NetappParams field if non-nil, zero value otherwise.

### GetNetappParamsOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetNetappParamsOk() (*NetappProtectionGroupParams, bool)`

GetNetappParamsOk returns a tuple with the NetappParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetappParams

`func (o *CreateOrUpdateProtectionGroupRequest) SetNetappParams(v NetappProtectionGroupParams)`

SetNetappParams sets NetappParams field to given value.

### HasNetappParams

`func (o *CreateOrUpdateProtectionGroupRequest) HasNetappParams() bool`

HasNetappParams returns a boolean if a field has been set.

### SetNetappParamsNil

`func (o *CreateOrUpdateProtectionGroupRequest) SetNetappParamsNil(b bool)`

 SetNetappParamsNil sets the value for NetappParams to be an explicit nil

### UnsetNetappParams
`func (o *CreateOrUpdateProtectionGroupRequest) UnsetNetappParams()`

UnsetNetappParams ensures that no value is present for NetappParams, not even an explicit nil
### GetNimbleParams

`func (o *CreateOrUpdateProtectionGroupRequest) GetNimbleParams() NimbleProtectionGroupParams`

GetNimbleParams returns the NimbleParams field if non-nil, zero value otherwise.

### GetNimbleParamsOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetNimbleParamsOk() (*NimbleProtectionGroupParams, bool)`

GetNimbleParamsOk returns a tuple with the NimbleParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNimbleParams

`func (o *CreateOrUpdateProtectionGroupRequest) SetNimbleParams(v NimbleProtectionGroupParams)`

SetNimbleParams sets NimbleParams field to given value.

### HasNimbleParams

`func (o *CreateOrUpdateProtectionGroupRequest) HasNimbleParams() bool`

HasNimbleParams returns a boolean if a field has been set.

### SetNimbleParamsNil

`func (o *CreateOrUpdateProtectionGroupRequest) SetNimbleParamsNil(b bool)`

 SetNimbleParamsNil sets the value for NimbleParams to be an explicit nil

### UnsetNimbleParams
`func (o *CreateOrUpdateProtectionGroupRequest) UnsetNimbleParams()`

UnsetNimbleParams ensures that no value is present for NimbleParams, not even an explicit nil
### GetOffice365Params

`func (o *CreateOrUpdateProtectionGroupRequest) GetOffice365Params() Office365ProtectionGroupParams`

GetOffice365Params returns the Office365Params field if non-nil, zero value otherwise.

### GetOffice365ParamsOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetOffice365ParamsOk() (*Office365ProtectionGroupParams, bool)`

GetOffice365ParamsOk returns a tuple with the Office365Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffice365Params

`func (o *CreateOrUpdateProtectionGroupRequest) SetOffice365Params(v Office365ProtectionGroupParams)`

SetOffice365Params sets Office365Params field to given value.

### HasOffice365Params

`func (o *CreateOrUpdateProtectionGroupRequest) HasOffice365Params() bool`

HasOffice365Params returns a boolean if a field has been set.

### GetOracleParams

`func (o *CreateOrUpdateProtectionGroupRequest) GetOracleParams() OracleProtectionGroupParams`

GetOracleParams returns the OracleParams field if non-nil, zero value otherwise.

### GetOracleParamsOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetOracleParamsOk() (*OracleProtectionGroupParams, bool)`

GetOracleParamsOk returns a tuple with the OracleParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleParams

`func (o *CreateOrUpdateProtectionGroupRequest) SetOracleParams(v OracleProtectionGroupParams)`

SetOracleParams sets OracleParams field to given value.

### HasOracleParams

`func (o *CreateOrUpdateProtectionGroupRequest) HasOracleParams() bool`

HasOracleParams returns a boolean if a field has been set.

### GetPhysicalParams

`func (o *CreateOrUpdateProtectionGroupRequest) GetPhysicalParams() PhysicalProtectionGroupParams`

GetPhysicalParams returns the PhysicalParams field if non-nil, zero value otherwise.

### GetPhysicalParamsOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetPhysicalParamsOk() (*PhysicalProtectionGroupParams, bool)`

GetPhysicalParamsOk returns a tuple with the PhysicalParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalParams

`func (o *CreateOrUpdateProtectionGroupRequest) SetPhysicalParams(v PhysicalProtectionGroupParams)`

SetPhysicalParams sets PhysicalParams field to given value.

### HasPhysicalParams

`func (o *CreateOrUpdateProtectionGroupRequest) HasPhysicalParams() bool`

HasPhysicalParams returns a boolean if a field has been set.

### GetPureParams

`func (o *CreateOrUpdateProtectionGroupRequest) GetPureParams() PureProtectionGroupParams`

GetPureParams returns the PureParams field if non-nil, zero value otherwise.

### GetPureParamsOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetPureParamsOk() (*PureProtectionGroupParams, bool)`

GetPureParamsOk returns a tuple with the PureParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPureParams

`func (o *CreateOrUpdateProtectionGroupRequest) SetPureParams(v PureProtectionGroupParams)`

SetPureParams sets PureParams field to given value.

### HasPureParams

`func (o *CreateOrUpdateProtectionGroupRequest) HasPureParams() bool`

HasPureParams returns a boolean if a field has been set.

### SetPureParamsNil

`func (o *CreateOrUpdateProtectionGroupRequest) SetPureParamsNil(b bool)`

 SetPureParamsNil sets the value for PureParams to be an explicit nil

### UnsetPureParams
`func (o *CreateOrUpdateProtectionGroupRequest) UnsetPureParams()`

UnsetPureParams ensures that no value is present for PureParams, not even an explicit nil
### GetRemoteAdapterParams

`func (o *CreateOrUpdateProtectionGroupRequest) GetRemoteAdapterParams() RemoteAdapterProtectionGroupParams`

GetRemoteAdapterParams returns the RemoteAdapterParams field if non-nil, zero value otherwise.

### GetRemoteAdapterParamsOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetRemoteAdapterParamsOk() (*RemoteAdapterProtectionGroupParams, bool)`

GetRemoteAdapterParamsOk returns a tuple with the RemoteAdapterParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAdapterParams

`func (o *CreateOrUpdateProtectionGroupRequest) SetRemoteAdapterParams(v RemoteAdapterProtectionGroupParams)`

SetRemoteAdapterParams sets RemoteAdapterParams field to given value.

### HasRemoteAdapterParams

`func (o *CreateOrUpdateProtectionGroupRequest) HasRemoteAdapterParams() bool`

HasRemoteAdapterParams returns a boolean if a field has been set.

### SetRemoteAdapterParamsNil

`func (o *CreateOrUpdateProtectionGroupRequest) SetRemoteAdapterParamsNil(b bool)`

 SetRemoteAdapterParamsNil sets the value for RemoteAdapterParams to be an explicit nil

### UnsetRemoteAdapterParams
`func (o *CreateOrUpdateProtectionGroupRequest) UnsetRemoteAdapterParams()`

UnsetRemoteAdapterParams ensures that no value is present for RemoteAdapterParams, not even an explicit nil
### GetSapHanaParams

`func (o *CreateOrUpdateProtectionGroupRequest) GetSapHanaParams() SapHanaProtectionGroupParams`

GetSapHanaParams returns the SapHanaParams field if non-nil, zero value otherwise.

### GetSapHanaParamsOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetSapHanaParamsOk() (*SapHanaProtectionGroupParams, bool)`

GetSapHanaParamsOk returns a tuple with the SapHanaParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSapHanaParams

`func (o *CreateOrUpdateProtectionGroupRequest) SetSapHanaParams(v SapHanaProtectionGroupParams)`

SetSapHanaParams sets SapHanaParams field to given value.

### HasSapHanaParams

`func (o *CreateOrUpdateProtectionGroupRequest) HasSapHanaParams() bool`

HasSapHanaParams returns a boolean if a field has been set.

### GetSfdcParams

`func (o *CreateOrUpdateProtectionGroupRequest) GetSfdcParams() SfdcProtectionGroupParams`

GetSfdcParams returns the SfdcParams field if non-nil, zero value otherwise.

### GetSfdcParamsOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetSfdcParamsOk() (*SfdcProtectionGroupParams, bool)`

GetSfdcParamsOk returns a tuple with the SfdcParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfdcParams

`func (o *CreateOrUpdateProtectionGroupRequest) SetSfdcParams(v SfdcProtectionGroupParams)`

SetSfdcParams sets SfdcParams field to given value.

### HasSfdcParams

`func (o *CreateOrUpdateProtectionGroupRequest) HasSfdcParams() bool`

HasSfdcParams returns a boolean if a field has been set.

### GetUdaParams

`func (o *CreateOrUpdateProtectionGroupRequest) GetUdaParams() UdaProtectionGroupParams`

GetUdaParams returns the UdaParams field if non-nil, zero value otherwise.

### GetUdaParamsOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetUdaParamsOk() (*UdaProtectionGroupParams, bool)`

GetUdaParamsOk returns a tuple with the UdaParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdaParams

`func (o *CreateOrUpdateProtectionGroupRequest) SetUdaParams(v UdaProtectionGroupParams)`

SetUdaParams sets UdaParams field to given value.

### HasUdaParams

`func (o *CreateOrUpdateProtectionGroupRequest) HasUdaParams() bool`

HasUdaParams returns a boolean if a field has been set.

### GetViewParams

`func (o *CreateOrUpdateProtectionGroupRequest) GetViewParams() ViewProtectionGroupParams`

GetViewParams returns the ViewParams field if non-nil, zero value otherwise.

### GetViewParamsOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetViewParamsOk() (*ViewProtectionGroupParams, bool)`

GetViewParamsOk returns a tuple with the ViewParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewParams

`func (o *CreateOrUpdateProtectionGroupRequest) SetViewParams(v ViewProtectionGroupParams)`

SetViewParams sets ViewParams field to given value.

### HasViewParams

`func (o *CreateOrUpdateProtectionGroupRequest) HasViewParams() bool`

HasViewParams returns a boolean if a field has been set.

### SetViewParamsNil

`func (o *CreateOrUpdateProtectionGroupRequest) SetViewParamsNil(b bool)`

 SetViewParamsNil sets the value for ViewParams to be an explicit nil

### UnsetViewParams
`func (o *CreateOrUpdateProtectionGroupRequest) UnsetViewParams()`

UnsetViewParams ensures that no value is present for ViewParams, not even an explicit nil
### GetVmwareParams

`func (o *CreateOrUpdateProtectionGroupRequest) GetVmwareParams() VmwareProtectionGroupParams`

GetVmwareParams returns the VmwareParams field if non-nil, zero value otherwise.

### GetVmwareParamsOk

`func (o *CreateOrUpdateProtectionGroupRequest) GetVmwareParamsOk() (*VmwareProtectionGroupParams, bool)`

GetVmwareParamsOk returns a tuple with the VmwareParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareParams

`func (o *CreateOrUpdateProtectionGroupRequest) SetVmwareParams(v VmwareProtectionGroupParams)`

SetVmwareParams sets VmwareParams field to given value.

### HasVmwareParams

`func (o *CreateOrUpdateProtectionGroupRequest) HasVmwareParams() bool`

HasVmwareParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


