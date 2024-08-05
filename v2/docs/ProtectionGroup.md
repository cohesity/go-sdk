# ProtectionGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbortInBlackouts** | Pointer to **NullableBool** | Specifies whether currently executing jobs should abort if a blackout period specified by a policy starts. Available only if the selected policy has at least one blackout period. Default value is false. | [optional] 
**AdvancedConfigs** | Pointer to [**[]KeyValuePair**](KeyValuePair.md) | Specifies the advanced configuration for a protection job. | [optional] 
**AlertPolicy** | Pointer to [**ProtectionGroupAlertingPolicy**](ProtectionGroupAlertingPolicy.md) |  | [optional] 
**ClusterId** | Pointer to **NullableString** | Specifies the cluster ID. | [optional] 
**Description** | Pointer to **NullableString** | Specifies a description of the Protection Group. | [optional] 
**EndTimeUsecs** | Pointer to **NullableInt64** | Specifies the end time in micro seconds for this Protection Group. If this is not specified, the Protection Group won&#39;t be ended. | [optional] 
**Environment** | Pointer to **NullableString** | Specifies the environment of the Protection Group. | [optional] 
**Id** | Pointer to **NullableString** | Specifies the ID of the Protection Group. | [optional] 
**InvalidEntities** | Pointer to [**[]MissingEntityParams**](MissingEntityParams.md) | Specifies the Information about invalid entities. An entity will be considered invalid if it is part of an active protection group but has lost compatibility for the given backup type. | [optional] 
**IsActive** | Pointer to **NullableBool** | Specifies if the Protection Group is active or not. | [optional] 
**IsDeleted** | Pointer to **NullableBool** | Specifies if the Protection Group has been deleted. | [optional] 
**IsPaused** | Pointer to **NullableBool** | Specifies if the the Protection Group is paused. New runs are not scheduled for the paused Protection Groups. Active run if any is not impacted. | [optional] 
**IsProtectOnce** | Pointer to **NullableBool** | Specifies if the the Protection Group is using a protect once type of policy. This field is helpful to identify run happen for this group. | [optional] 
**LastModifiedTimestampUsecs** | Pointer to **NullableInt64** | Specifies the last time this protection group was updated. If this is passed into a PUT request, then the backend will validate that the timestamp passed in matches the time that the protection group was actually last modified. If the two timestamps do not match, then the request will be rejected with a stale error. | [optional] 
**LastRun** | Pointer to [**ProtectionGroupRun**](ProtectionGroupRun.md) |  | [optional] 
**MissingEntities** | Pointer to [**[]MissingEntityParams**](MissingEntityParams.md) | Specifies the Information about missing entities. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the Protection Group. | [optional] 
**NumProtectedObjects** | Pointer to **NullableInt64** | Specifies the number of protected objects of the Protection Group. | [optional] 
**PauseInBlackouts** | Pointer to **NullableBool** | Specifies whether currently executing jobs should be paused if a blackout period specified by a policy starts. Available only if the selected policy has at least one blackout period. Default value is false. This field should not be set to true if &#39;abortInBlackouts&#39; is sent as true. | [optional] 
**Permissions** | Pointer to [**[]Tenant**](Tenant.md) | Specifies the list of tenants that have permissions for this protection group. | [optional] 
**PolicyId** | Pointer to **NullableString** | Specifies the unique id of the Protection Policy associated with the Protection Group. The Policy provides retry settings Protection Schedules, Priority, SLA, etc. | [optional] 
**Priority** | Pointer to **NullableString** | Specifies the priority of the Protection Group. | [optional] 
**QosPolicy** | Pointer to **NullableString** | Specifies whether the Protection Group will be written to HDD or SSD. | [optional] 
**RegionId** | Pointer to **NullableString** | Specifies the region ID. | [optional] 
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

### NewProtectionGroup

`func NewProtectionGroup() *ProtectionGroup`

NewProtectionGroup instantiates a new ProtectionGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionGroupWithDefaults

`func NewProtectionGroupWithDefaults() *ProtectionGroup`

NewProtectionGroupWithDefaults instantiates a new ProtectionGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbortInBlackouts

`func (o *ProtectionGroup) GetAbortInBlackouts() bool`

GetAbortInBlackouts returns the AbortInBlackouts field if non-nil, zero value otherwise.

### GetAbortInBlackoutsOk

`func (o *ProtectionGroup) GetAbortInBlackoutsOk() (*bool, bool)`

GetAbortInBlackoutsOk returns a tuple with the AbortInBlackouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbortInBlackouts

`func (o *ProtectionGroup) SetAbortInBlackouts(v bool)`

SetAbortInBlackouts sets AbortInBlackouts field to given value.

### HasAbortInBlackouts

`func (o *ProtectionGroup) HasAbortInBlackouts() bool`

HasAbortInBlackouts returns a boolean if a field has been set.

### SetAbortInBlackoutsNil

`func (o *ProtectionGroup) SetAbortInBlackoutsNil(b bool)`

 SetAbortInBlackoutsNil sets the value for AbortInBlackouts to be an explicit nil

### UnsetAbortInBlackouts
`func (o *ProtectionGroup) UnsetAbortInBlackouts()`

UnsetAbortInBlackouts ensures that no value is present for AbortInBlackouts, not even an explicit nil
### GetAdvancedConfigs

`func (o *ProtectionGroup) GetAdvancedConfigs() []KeyValuePair`

GetAdvancedConfigs returns the AdvancedConfigs field if non-nil, zero value otherwise.

### GetAdvancedConfigsOk

`func (o *ProtectionGroup) GetAdvancedConfigsOk() (*[]KeyValuePair, bool)`

GetAdvancedConfigsOk returns a tuple with the AdvancedConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedConfigs

`func (o *ProtectionGroup) SetAdvancedConfigs(v []KeyValuePair)`

SetAdvancedConfigs sets AdvancedConfigs field to given value.

### HasAdvancedConfigs

`func (o *ProtectionGroup) HasAdvancedConfigs() bool`

HasAdvancedConfigs returns a boolean if a field has been set.

### SetAdvancedConfigsNil

`func (o *ProtectionGroup) SetAdvancedConfigsNil(b bool)`

 SetAdvancedConfigsNil sets the value for AdvancedConfigs to be an explicit nil

### UnsetAdvancedConfigs
`func (o *ProtectionGroup) UnsetAdvancedConfigs()`

UnsetAdvancedConfigs ensures that no value is present for AdvancedConfigs, not even an explicit nil
### GetAlertPolicy

`func (o *ProtectionGroup) GetAlertPolicy() ProtectionGroupAlertingPolicy`

GetAlertPolicy returns the AlertPolicy field if non-nil, zero value otherwise.

### GetAlertPolicyOk

`func (o *ProtectionGroup) GetAlertPolicyOk() (*ProtectionGroupAlertingPolicy, bool)`

GetAlertPolicyOk returns a tuple with the AlertPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertPolicy

`func (o *ProtectionGroup) SetAlertPolicy(v ProtectionGroupAlertingPolicy)`

SetAlertPolicy sets AlertPolicy field to given value.

### HasAlertPolicy

`func (o *ProtectionGroup) HasAlertPolicy() bool`

HasAlertPolicy returns a boolean if a field has been set.

### GetClusterId

`func (o *ProtectionGroup) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ProtectionGroup) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ProtectionGroup) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *ProtectionGroup) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### SetClusterIdNil

`func (o *ProtectionGroup) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *ProtectionGroup) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetDescription

`func (o *ProtectionGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProtectionGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProtectionGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProtectionGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ProtectionGroup) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ProtectionGroup) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEndTimeUsecs

`func (o *ProtectionGroup) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *ProtectionGroup) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *ProtectionGroup) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *ProtectionGroup) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *ProtectionGroup) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *ProtectionGroup) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetEnvironment

`func (o *ProtectionGroup) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ProtectionGroup) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ProtectionGroup) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ProtectionGroup) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *ProtectionGroup) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *ProtectionGroup) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetId

`func (o *ProtectionGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProtectionGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProtectionGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProtectionGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ProtectionGroup) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ProtectionGroup) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetInvalidEntities

`func (o *ProtectionGroup) GetInvalidEntities() []MissingEntityParams`

GetInvalidEntities returns the InvalidEntities field if non-nil, zero value otherwise.

### GetInvalidEntitiesOk

`func (o *ProtectionGroup) GetInvalidEntitiesOk() (*[]MissingEntityParams, bool)`

GetInvalidEntitiesOk returns a tuple with the InvalidEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidEntities

`func (o *ProtectionGroup) SetInvalidEntities(v []MissingEntityParams)`

SetInvalidEntities sets InvalidEntities field to given value.

### HasInvalidEntities

`func (o *ProtectionGroup) HasInvalidEntities() bool`

HasInvalidEntities returns a boolean if a field has been set.

### SetInvalidEntitiesNil

`func (o *ProtectionGroup) SetInvalidEntitiesNil(b bool)`

 SetInvalidEntitiesNil sets the value for InvalidEntities to be an explicit nil

### UnsetInvalidEntities
`func (o *ProtectionGroup) UnsetInvalidEntities()`

UnsetInvalidEntities ensures that no value is present for InvalidEntities, not even an explicit nil
### GetIsActive

`func (o *ProtectionGroup) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ProtectionGroup) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ProtectionGroup) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *ProtectionGroup) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *ProtectionGroup) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *ProtectionGroup) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetIsDeleted

`func (o *ProtectionGroup) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *ProtectionGroup) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *ProtectionGroup) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *ProtectionGroup) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### SetIsDeletedNil

`func (o *ProtectionGroup) SetIsDeletedNil(b bool)`

 SetIsDeletedNil sets the value for IsDeleted to be an explicit nil

### UnsetIsDeleted
`func (o *ProtectionGroup) UnsetIsDeleted()`

UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil
### GetIsPaused

`func (o *ProtectionGroup) GetIsPaused() bool`

GetIsPaused returns the IsPaused field if non-nil, zero value otherwise.

### GetIsPausedOk

`func (o *ProtectionGroup) GetIsPausedOk() (*bool, bool)`

GetIsPausedOk returns a tuple with the IsPaused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPaused

`func (o *ProtectionGroup) SetIsPaused(v bool)`

SetIsPaused sets IsPaused field to given value.

### HasIsPaused

`func (o *ProtectionGroup) HasIsPaused() bool`

HasIsPaused returns a boolean if a field has been set.

### SetIsPausedNil

`func (o *ProtectionGroup) SetIsPausedNil(b bool)`

 SetIsPausedNil sets the value for IsPaused to be an explicit nil

### UnsetIsPaused
`func (o *ProtectionGroup) UnsetIsPaused()`

UnsetIsPaused ensures that no value is present for IsPaused, not even an explicit nil
### GetIsProtectOnce

`func (o *ProtectionGroup) GetIsProtectOnce() bool`

GetIsProtectOnce returns the IsProtectOnce field if non-nil, zero value otherwise.

### GetIsProtectOnceOk

`func (o *ProtectionGroup) GetIsProtectOnceOk() (*bool, bool)`

GetIsProtectOnceOk returns a tuple with the IsProtectOnce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProtectOnce

`func (o *ProtectionGroup) SetIsProtectOnce(v bool)`

SetIsProtectOnce sets IsProtectOnce field to given value.

### HasIsProtectOnce

`func (o *ProtectionGroup) HasIsProtectOnce() bool`

HasIsProtectOnce returns a boolean if a field has been set.

### SetIsProtectOnceNil

`func (o *ProtectionGroup) SetIsProtectOnceNil(b bool)`

 SetIsProtectOnceNil sets the value for IsProtectOnce to be an explicit nil

### UnsetIsProtectOnce
`func (o *ProtectionGroup) UnsetIsProtectOnce()`

UnsetIsProtectOnce ensures that no value is present for IsProtectOnce, not even an explicit nil
### GetLastModifiedTimestampUsecs

`func (o *ProtectionGroup) GetLastModifiedTimestampUsecs() int64`

GetLastModifiedTimestampUsecs returns the LastModifiedTimestampUsecs field if non-nil, zero value otherwise.

### GetLastModifiedTimestampUsecsOk

`func (o *ProtectionGroup) GetLastModifiedTimestampUsecsOk() (*int64, bool)`

GetLastModifiedTimestampUsecsOk returns a tuple with the LastModifiedTimestampUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTimestampUsecs

`func (o *ProtectionGroup) SetLastModifiedTimestampUsecs(v int64)`

SetLastModifiedTimestampUsecs sets LastModifiedTimestampUsecs field to given value.

### HasLastModifiedTimestampUsecs

`func (o *ProtectionGroup) HasLastModifiedTimestampUsecs() bool`

HasLastModifiedTimestampUsecs returns a boolean if a field has been set.

### SetLastModifiedTimestampUsecsNil

`func (o *ProtectionGroup) SetLastModifiedTimestampUsecsNil(b bool)`

 SetLastModifiedTimestampUsecsNil sets the value for LastModifiedTimestampUsecs to be an explicit nil

### UnsetLastModifiedTimestampUsecs
`func (o *ProtectionGroup) UnsetLastModifiedTimestampUsecs()`

UnsetLastModifiedTimestampUsecs ensures that no value is present for LastModifiedTimestampUsecs, not even an explicit nil
### GetLastRun

`func (o *ProtectionGroup) GetLastRun() ProtectionGroupRun`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *ProtectionGroup) GetLastRunOk() (*ProtectionGroupRun, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *ProtectionGroup) SetLastRun(v ProtectionGroupRun)`

SetLastRun sets LastRun field to given value.

### HasLastRun

`func (o *ProtectionGroup) HasLastRun() bool`

HasLastRun returns a boolean if a field has been set.

### GetMissingEntities

`func (o *ProtectionGroup) GetMissingEntities() []MissingEntityParams`

GetMissingEntities returns the MissingEntities field if non-nil, zero value otherwise.

### GetMissingEntitiesOk

`func (o *ProtectionGroup) GetMissingEntitiesOk() (*[]MissingEntityParams, bool)`

GetMissingEntitiesOk returns a tuple with the MissingEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingEntities

`func (o *ProtectionGroup) SetMissingEntities(v []MissingEntityParams)`

SetMissingEntities sets MissingEntities field to given value.

### HasMissingEntities

`func (o *ProtectionGroup) HasMissingEntities() bool`

HasMissingEntities returns a boolean if a field has been set.

### SetMissingEntitiesNil

`func (o *ProtectionGroup) SetMissingEntitiesNil(b bool)`

 SetMissingEntitiesNil sets the value for MissingEntities to be an explicit nil

### UnsetMissingEntities
`func (o *ProtectionGroup) UnsetMissingEntities()`

UnsetMissingEntities ensures that no value is present for MissingEntities, not even an explicit nil
### GetName

`func (o *ProtectionGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProtectionGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProtectionGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProtectionGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ProtectionGroup) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProtectionGroup) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNumProtectedObjects

`func (o *ProtectionGroup) GetNumProtectedObjects() int64`

GetNumProtectedObjects returns the NumProtectedObjects field if non-nil, zero value otherwise.

### GetNumProtectedObjectsOk

`func (o *ProtectionGroup) GetNumProtectedObjectsOk() (*int64, bool)`

GetNumProtectedObjectsOk returns a tuple with the NumProtectedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumProtectedObjects

`func (o *ProtectionGroup) SetNumProtectedObjects(v int64)`

SetNumProtectedObjects sets NumProtectedObjects field to given value.

### HasNumProtectedObjects

`func (o *ProtectionGroup) HasNumProtectedObjects() bool`

HasNumProtectedObjects returns a boolean if a field has been set.

### SetNumProtectedObjectsNil

`func (o *ProtectionGroup) SetNumProtectedObjectsNil(b bool)`

 SetNumProtectedObjectsNil sets the value for NumProtectedObjects to be an explicit nil

### UnsetNumProtectedObjects
`func (o *ProtectionGroup) UnsetNumProtectedObjects()`

UnsetNumProtectedObjects ensures that no value is present for NumProtectedObjects, not even an explicit nil
### GetPauseInBlackouts

`func (o *ProtectionGroup) GetPauseInBlackouts() bool`

GetPauseInBlackouts returns the PauseInBlackouts field if non-nil, zero value otherwise.

### GetPauseInBlackoutsOk

`func (o *ProtectionGroup) GetPauseInBlackoutsOk() (*bool, bool)`

GetPauseInBlackoutsOk returns a tuple with the PauseInBlackouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseInBlackouts

`func (o *ProtectionGroup) SetPauseInBlackouts(v bool)`

SetPauseInBlackouts sets PauseInBlackouts field to given value.

### HasPauseInBlackouts

`func (o *ProtectionGroup) HasPauseInBlackouts() bool`

HasPauseInBlackouts returns a boolean if a field has been set.

### SetPauseInBlackoutsNil

`func (o *ProtectionGroup) SetPauseInBlackoutsNil(b bool)`

 SetPauseInBlackoutsNil sets the value for PauseInBlackouts to be an explicit nil

### UnsetPauseInBlackouts
`func (o *ProtectionGroup) UnsetPauseInBlackouts()`

UnsetPauseInBlackouts ensures that no value is present for PauseInBlackouts, not even an explicit nil
### GetPermissions

`func (o *ProtectionGroup) GetPermissions() []Tenant`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ProtectionGroup) GetPermissionsOk() (*[]Tenant, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ProtectionGroup) SetPermissions(v []Tenant)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ProtectionGroup) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *ProtectionGroup) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *ProtectionGroup) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetPolicyId

`func (o *ProtectionGroup) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *ProtectionGroup) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *ProtectionGroup) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *ProtectionGroup) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### SetPolicyIdNil

`func (o *ProtectionGroup) SetPolicyIdNil(b bool)`

 SetPolicyIdNil sets the value for PolicyId to be an explicit nil

### UnsetPolicyId
`func (o *ProtectionGroup) UnsetPolicyId()`

UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
### GetPriority

`func (o *ProtectionGroup) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ProtectionGroup) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ProtectionGroup) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ProtectionGroup) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *ProtectionGroup) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *ProtectionGroup) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetQosPolicy

`func (o *ProtectionGroup) GetQosPolicy() string`

GetQosPolicy returns the QosPolicy field if non-nil, zero value otherwise.

### GetQosPolicyOk

`func (o *ProtectionGroup) GetQosPolicyOk() (*string, bool)`

GetQosPolicyOk returns a tuple with the QosPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosPolicy

`func (o *ProtectionGroup) SetQosPolicy(v string)`

SetQosPolicy sets QosPolicy field to given value.

### HasQosPolicy

`func (o *ProtectionGroup) HasQosPolicy() bool`

HasQosPolicy returns a boolean if a field has been set.

### SetQosPolicyNil

`func (o *ProtectionGroup) SetQosPolicyNil(b bool)`

 SetQosPolicyNil sets the value for QosPolicy to be an explicit nil

### UnsetQosPolicy
`func (o *ProtectionGroup) UnsetQosPolicy()`

UnsetQosPolicy ensures that no value is present for QosPolicy, not even an explicit nil
### GetRegionId

`func (o *ProtectionGroup) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *ProtectionGroup) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *ProtectionGroup) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *ProtectionGroup) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### SetRegionIdNil

`func (o *ProtectionGroup) SetRegionIdNil(b bool)`

 SetRegionIdNil sets the value for RegionId to be an explicit nil

### UnsetRegionId
`func (o *ProtectionGroup) UnsetRegionId()`

UnsetRegionId ensures that no value is present for RegionId, not even an explicit nil
### GetSla

`func (o *ProtectionGroup) GetSla() []SlaRule`

GetSla returns the Sla field if non-nil, zero value otherwise.

### GetSlaOk

`func (o *ProtectionGroup) GetSlaOk() (*[]SlaRule, bool)`

GetSlaOk returns a tuple with the Sla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSla

`func (o *ProtectionGroup) SetSla(v []SlaRule)`

SetSla sets Sla field to given value.

### HasSla

`func (o *ProtectionGroup) HasSla() bool`

HasSla returns a boolean if a field has been set.

### SetSlaNil

`func (o *ProtectionGroup) SetSlaNil(b bool)`

 SetSlaNil sets the value for Sla to be an explicit nil

### UnsetSla
`func (o *ProtectionGroup) UnsetSla()`

UnsetSla ensures that no value is present for Sla, not even an explicit nil
### GetStartTime

`func (o *ProtectionGroup) GetStartTime() TimeOfDay`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ProtectionGroup) GetStartTimeOk() (*TimeOfDay, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ProtectionGroup) SetStartTime(v TimeOfDay)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ProtectionGroup) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStorageDomainId

`func (o *ProtectionGroup) GetStorageDomainId() int64`

GetStorageDomainId returns the StorageDomainId field if non-nil, zero value otherwise.

### GetStorageDomainIdOk

`func (o *ProtectionGroup) GetStorageDomainIdOk() (*int64, bool)`

GetStorageDomainIdOk returns a tuple with the StorageDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainId

`func (o *ProtectionGroup) SetStorageDomainId(v int64)`

SetStorageDomainId sets StorageDomainId field to given value.

### HasStorageDomainId

`func (o *ProtectionGroup) HasStorageDomainId() bool`

HasStorageDomainId returns a boolean if a field has been set.

### SetStorageDomainIdNil

`func (o *ProtectionGroup) SetStorageDomainIdNil(b bool)`

 SetStorageDomainIdNil sets the value for StorageDomainId to be an explicit nil

### UnsetStorageDomainId
`func (o *ProtectionGroup) UnsetStorageDomainId()`

UnsetStorageDomainId ensures that no value is present for StorageDomainId, not even an explicit nil
### GetAcropolisParams

`func (o *ProtectionGroup) GetAcropolisParams() AcropolisProtectionGroupParams`

GetAcropolisParams returns the AcropolisParams field if non-nil, zero value otherwise.

### GetAcropolisParamsOk

`func (o *ProtectionGroup) GetAcropolisParamsOk() (*AcropolisProtectionGroupParams, bool)`

GetAcropolisParamsOk returns a tuple with the AcropolisParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcropolisParams

`func (o *ProtectionGroup) SetAcropolisParams(v AcropolisProtectionGroupParams)`

SetAcropolisParams sets AcropolisParams field to given value.

### HasAcropolisParams

`func (o *ProtectionGroup) HasAcropolisParams() bool`

HasAcropolisParams returns a boolean if a field has been set.

### GetAdParams

`func (o *ProtectionGroup) GetAdParams() ADProtectionGroupParams`

GetAdParams returns the AdParams field if non-nil, zero value otherwise.

### GetAdParamsOk

`func (o *ProtectionGroup) GetAdParamsOk() (*ADProtectionGroupParams, bool)`

GetAdParamsOk returns a tuple with the AdParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdParams

`func (o *ProtectionGroup) SetAdParams(v ADProtectionGroupParams)`

SetAdParams sets AdParams field to given value.

### HasAdParams

`func (o *ProtectionGroup) HasAdParams() bool`

HasAdParams returns a boolean if a field has been set.

### GetAwsParams

`func (o *ProtectionGroup) GetAwsParams() AwsProtectionGroupParams`

GetAwsParams returns the AwsParams field if non-nil, zero value otherwise.

### GetAwsParamsOk

`func (o *ProtectionGroup) GetAwsParamsOk() (*AwsProtectionGroupParams, bool)`

GetAwsParamsOk returns a tuple with the AwsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsParams

`func (o *ProtectionGroup) SetAwsParams(v AwsProtectionGroupParams)`

SetAwsParams sets AwsParams field to given value.

### HasAwsParams

`func (o *ProtectionGroup) HasAwsParams() bool`

HasAwsParams returns a boolean if a field has been set.

### GetAzureParams

`func (o *ProtectionGroup) GetAzureParams() AzureProtectionGroupParams`

GetAzureParams returns the AzureParams field if non-nil, zero value otherwise.

### GetAzureParamsOk

`func (o *ProtectionGroup) GetAzureParamsOk() (*AzureProtectionGroupParams, bool)`

GetAzureParamsOk returns a tuple with the AzureParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureParams

`func (o *ProtectionGroup) SetAzureParams(v AzureProtectionGroupParams)`

SetAzureParams sets AzureParams field to given value.

### HasAzureParams

`func (o *ProtectionGroup) HasAzureParams() bool`

HasAzureParams returns a boolean if a field has been set.

### GetCassandraParams

`func (o *ProtectionGroup) GetCassandraParams() CassandraProtectionGroupParams`

GetCassandraParams returns the CassandraParams field if non-nil, zero value otherwise.

### GetCassandraParamsOk

`func (o *ProtectionGroup) GetCassandraParamsOk() (*CassandraProtectionGroupParams, bool)`

GetCassandraParamsOk returns a tuple with the CassandraParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraParams

`func (o *ProtectionGroup) SetCassandraParams(v CassandraProtectionGroupParams)`

SetCassandraParams sets CassandraParams field to given value.

### HasCassandraParams

`func (o *ProtectionGroup) HasCassandraParams() bool`

HasCassandraParams returns a boolean if a field has been set.

### GetCouchbaseParams

`func (o *ProtectionGroup) GetCouchbaseParams() NoSqlProtectionGroupParams`

GetCouchbaseParams returns the CouchbaseParams field if non-nil, zero value otherwise.

### GetCouchbaseParamsOk

`func (o *ProtectionGroup) GetCouchbaseParamsOk() (*NoSqlProtectionGroupParams, bool)`

GetCouchbaseParamsOk returns a tuple with the CouchbaseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouchbaseParams

`func (o *ProtectionGroup) SetCouchbaseParams(v NoSqlProtectionGroupParams)`

SetCouchbaseParams sets CouchbaseParams field to given value.

### HasCouchbaseParams

`func (o *ProtectionGroup) HasCouchbaseParams() bool`

HasCouchbaseParams returns a boolean if a field has been set.

### GetElastifileParams

`func (o *ProtectionGroup) GetElastifileParams() ElastifileProtectionGroupParams`

GetElastifileParams returns the ElastifileParams field if non-nil, zero value otherwise.

### GetElastifileParamsOk

`func (o *ProtectionGroup) GetElastifileParamsOk() (*ElastifileProtectionGroupParams, bool)`

GetElastifileParamsOk returns a tuple with the ElastifileParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElastifileParams

`func (o *ProtectionGroup) SetElastifileParams(v ElastifileProtectionGroupParams)`

SetElastifileParams sets ElastifileParams field to given value.

### HasElastifileParams

`func (o *ProtectionGroup) HasElastifileParams() bool`

HasElastifileParams returns a boolean if a field has been set.

### SetElastifileParamsNil

`func (o *ProtectionGroup) SetElastifileParamsNil(b bool)`

 SetElastifileParamsNil sets the value for ElastifileParams to be an explicit nil

### UnsetElastifileParams
`func (o *ProtectionGroup) UnsetElastifileParams()`

UnsetElastifileParams ensures that no value is present for ElastifileParams, not even an explicit nil
### GetExchangeParams

`func (o *ProtectionGroup) GetExchangeParams() ExchangeProtectionGroupParams`

GetExchangeParams returns the ExchangeParams field if non-nil, zero value otherwise.

### GetExchangeParamsOk

`func (o *ProtectionGroup) GetExchangeParamsOk() (*ExchangeProtectionGroupParams, bool)`

GetExchangeParamsOk returns a tuple with the ExchangeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeParams

`func (o *ProtectionGroup) SetExchangeParams(v ExchangeProtectionGroupParams)`

SetExchangeParams sets ExchangeParams field to given value.

### HasExchangeParams

`func (o *ProtectionGroup) HasExchangeParams() bool`

HasExchangeParams returns a boolean if a field has been set.

### GetExperimentalAdapterParams

`func (o *ProtectionGroup) GetExperimentalAdapterParams() ExperimentalAdapterProtectionGroupParams`

GetExperimentalAdapterParams returns the ExperimentalAdapterParams field if non-nil, zero value otherwise.

### GetExperimentalAdapterParamsOk

`func (o *ProtectionGroup) GetExperimentalAdapterParamsOk() (*ExperimentalAdapterProtectionGroupParams, bool)`

GetExperimentalAdapterParamsOk returns a tuple with the ExperimentalAdapterParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentalAdapterParams

`func (o *ProtectionGroup) SetExperimentalAdapterParams(v ExperimentalAdapterProtectionGroupParams)`

SetExperimentalAdapterParams sets ExperimentalAdapterParams field to given value.

### HasExperimentalAdapterParams

`func (o *ProtectionGroup) HasExperimentalAdapterParams() bool`

HasExperimentalAdapterParams returns a boolean if a field has been set.

### GetFlashbladeParams

`func (o *ProtectionGroup) GetFlashbladeParams() FlashbladeProtectionGroupParams`

GetFlashbladeParams returns the FlashbladeParams field if non-nil, zero value otherwise.

### GetFlashbladeParamsOk

`func (o *ProtectionGroup) GetFlashbladeParamsOk() (*FlashbladeProtectionGroupParams, bool)`

GetFlashbladeParamsOk returns a tuple with the FlashbladeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlashbladeParams

`func (o *ProtectionGroup) SetFlashbladeParams(v FlashbladeProtectionGroupParams)`

SetFlashbladeParams sets FlashbladeParams field to given value.

### HasFlashbladeParams

`func (o *ProtectionGroup) HasFlashbladeParams() bool`

HasFlashbladeParams returns a boolean if a field has been set.

### SetFlashbladeParamsNil

`func (o *ProtectionGroup) SetFlashbladeParamsNil(b bool)`

 SetFlashbladeParamsNil sets the value for FlashbladeParams to be an explicit nil

### UnsetFlashbladeParams
`func (o *ProtectionGroup) UnsetFlashbladeParams()`

UnsetFlashbladeParams ensures that no value is present for FlashbladeParams, not even an explicit nil
### GetGcpParams

`func (o *ProtectionGroup) GetGcpParams() GcpProtectionGroupParams`

GetGcpParams returns the GcpParams field if non-nil, zero value otherwise.

### GetGcpParamsOk

`func (o *ProtectionGroup) GetGcpParamsOk() (*GcpProtectionGroupParams, bool)`

GetGcpParamsOk returns a tuple with the GcpParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpParams

`func (o *ProtectionGroup) SetGcpParams(v GcpProtectionGroupParams)`

SetGcpParams sets GcpParams field to given value.

### HasGcpParams

`func (o *ProtectionGroup) HasGcpParams() bool`

HasGcpParams returns a boolean if a field has been set.

### GetGenericNasParams

`func (o *ProtectionGroup) GetGenericNasParams() GenericNasProtectionGroupParams`

GetGenericNasParams returns the GenericNasParams field if non-nil, zero value otherwise.

### GetGenericNasParamsOk

`func (o *ProtectionGroup) GetGenericNasParamsOk() (*GenericNasProtectionGroupParams, bool)`

GetGenericNasParamsOk returns a tuple with the GenericNasParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericNasParams

`func (o *ProtectionGroup) SetGenericNasParams(v GenericNasProtectionGroupParams)`

SetGenericNasParams sets GenericNasParams field to given value.

### HasGenericNasParams

`func (o *ProtectionGroup) HasGenericNasParams() bool`

HasGenericNasParams returns a boolean if a field has been set.

### GetGpfsParams

`func (o *ProtectionGroup) GetGpfsParams() GpfsProtectionGroupParams`

GetGpfsParams returns the GpfsParams field if non-nil, zero value otherwise.

### GetGpfsParamsOk

`func (o *ProtectionGroup) GetGpfsParamsOk() (*GpfsProtectionGroupParams, bool)`

GetGpfsParamsOk returns a tuple with the GpfsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpfsParams

`func (o *ProtectionGroup) SetGpfsParams(v GpfsProtectionGroupParams)`

SetGpfsParams sets GpfsParams field to given value.

### HasGpfsParams

`func (o *ProtectionGroup) HasGpfsParams() bool`

HasGpfsParams returns a boolean if a field has been set.

### SetGpfsParamsNil

`func (o *ProtectionGroup) SetGpfsParamsNil(b bool)`

 SetGpfsParamsNil sets the value for GpfsParams to be an explicit nil

### UnsetGpfsParams
`func (o *ProtectionGroup) UnsetGpfsParams()`

UnsetGpfsParams ensures that no value is present for GpfsParams, not even an explicit nil
### GetHbaseParams

`func (o *ProtectionGroup) GetHbaseParams() NoSqlProtectionGroupParams`

GetHbaseParams returns the HbaseParams field if non-nil, zero value otherwise.

### GetHbaseParamsOk

`func (o *ProtectionGroup) GetHbaseParamsOk() (*NoSqlProtectionGroupParams, bool)`

GetHbaseParamsOk returns a tuple with the HbaseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHbaseParams

`func (o *ProtectionGroup) SetHbaseParams(v NoSqlProtectionGroupParams)`

SetHbaseParams sets HbaseParams field to given value.

### HasHbaseParams

`func (o *ProtectionGroup) HasHbaseParams() bool`

HasHbaseParams returns a boolean if a field has been set.

### GetHdfsParams

`func (o *ProtectionGroup) GetHdfsParams() HdfsProtectionGroupParams`

GetHdfsParams returns the HdfsParams field if non-nil, zero value otherwise.

### GetHdfsParamsOk

`func (o *ProtectionGroup) GetHdfsParamsOk() (*HdfsProtectionGroupParams, bool)`

GetHdfsParamsOk returns a tuple with the HdfsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsParams

`func (o *ProtectionGroup) SetHdfsParams(v HdfsProtectionGroupParams)`

SetHdfsParams sets HdfsParams field to given value.

### HasHdfsParams

`func (o *ProtectionGroup) HasHdfsParams() bool`

HasHdfsParams returns a boolean if a field has been set.

### GetHiveParams

`func (o *ProtectionGroup) GetHiveParams() NoSqlProtectionGroupParams`

GetHiveParams returns the HiveParams field if non-nil, zero value otherwise.

### GetHiveParamsOk

`func (o *ProtectionGroup) GetHiveParamsOk() (*NoSqlProtectionGroupParams, bool)`

GetHiveParamsOk returns a tuple with the HiveParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiveParams

`func (o *ProtectionGroup) SetHiveParams(v NoSqlProtectionGroupParams)`

SetHiveParams sets HiveParams field to given value.

### HasHiveParams

`func (o *ProtectionGroup) HasHiveParams() bool`

HasHiveParams returns a boolean if a field has been set.

### GetHypervParams

`func (o *ProtectionGroup) GetHypervParams() HyperVProtectionGroupParams`

GetHypervParams returns the HypervParams field if non-nil, zero value otherwise.

### GetHypervParamsOk

`func (o *ProtectionGroup) GetHypervParamsOk() (*HyperVProtectionGroupParams, bool)`

GetHypervParamsOk returns a tuple with the HypervParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervParams

`func (o *ProtectionGroup) SetHypervParams(v HyperVProtectionGroupParams)`

SetHypervParams sets HypervParams field to given value.

### HasHypervParams

`func (o *ProtectionGroup) HasHypervParams() bool`

HasHypervParams returns a boolean if a field has been set.

### GetIbmFlashSystemParams

`func (o *ProtectionGroup) GetIbmFlashSystemParams() IbmFlashSystemProtectionGroupParams`

GetIbmFlashSystemParams returns the IbmFlashSystemParams field if non-nil, zero value otherwise.

### GetIbmFlashSystemParamsOk

`func (o *ProtectionGroup) GetIbmFlashSystemParamsOk() (*IbmFlashSystemProtectionGroupParams, bool)`

GetIbmFlashSystemParamsOk returns a tuple with the IbmFlashSystemParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbmFlashSystemParams

`func (o *ProtectionGroup) SetIbmFlashSystemParams(v IbmFlashSystemProtectionGroupParams)`

SetIbmFlashSystemParams sets IbmFlashSystemParams field to given value.

### HasIbmFlashSystemParams

`func (o *ProtectionGroup) HasIbmFlashSystemParams() bool`

HasIbmFlashSystemParams returns a boolean if a field has been set.

### SetIbmFlashSystemParamsNil

`func (o *ProtectionGroup) SetIbmFlashSystemParamsNil(b bool)`

 SetIbmFlashSystemParamsNil sets the value for IbmFlashSystemParams to be an explicit nil

### UnsetIbmFlashSystemParams
`func (o *ProtectionGroup) UnsetIbmFlashSystemParams()`

UnsetIbmFlashSystemParams ensures that no value is present for IbmFlashSystemParams, not even an explicit nil
### GetIsilonParams

`func (o *ProtectionGroup) GetIsilonParams() IsilonProtectionGroupParams`

GetIsilonParams returns the IsilonParams field if non-nil, zero value otherwise.

### GetIsilonParamsOk

`func (o *ProtectionGroup) GetIsilonParamsOk() (*IsilonProtectionGroupParams, bool)`

GetIsilonParamsOk returns a tuple with the IsilonParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsilonParams

`func (o *ProtectionGroup) SetIsilonParams(v IsilonProtectionGroupParams)`

SetIsilonParams sets IsilonParams field to given value.

### HasIsilonParams

`func (o *ProtectionGroup) HasIsilonParams() bool`

HasIsilonParams returns a boolean if a field has been set.

### SetIsilonParamsNil

`func (o *ProtectionGroup) SetIsilonParamsNil(b bool)`

 SetIsilonParamsNil sets the value for IsilonParams to be an explicit nil

### UnsetIsilonParams
`func (o *ProtectionGroup) UnsetIsilonParams()`

UnsetIsilonParams ensures that no value is present for IsilonParams, not even an explicit nil
### GetKubernetesParams

`func (o *ProtectionGroup) GetKubernetesParams() KubernetesProtectionGroupParams`

GetKubernetesParams returns the KubernetesParams field if non-nil, zero value otherwise.

### GetKubernetesParamsOk

`func (o *ProtectionGroup) GetKubernetesParamsOk() (*KubernetesProtectionGroupParams, bool)`

GetKubernetesParamsOk returns a tuple with the KubernetesParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesParams

`func (o *ProtectionGroup) SetKubernetesParams(v KubernetesProtectionGroupParams)`

SetKubernetesParams sets KubernetesParams field to given value.

### HasKubernetesParams

`func (o *ProtectionGroup) HasKubernetesParams() bool`

HasKubernetesParams returns a boolean if a field has been set.

### GetKvmParams

`func (o *ProtectionGroup) GetKvmParams() KvmProtectionGroupParams`

GetKvmParams returns the KvmParams field if non-nil, zero value otherwise.

### GetKvmParamsOk

`func (o *ProtectionGroup) GetKvmParamsOk() (*KvmProtectionGroupParams, bool)`

GetKvmParamsOk returns a tuple with the KvmParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvmParams

`func (o *ProtectionGroup) SetKvmParams(v KvmProtectionGroupParams)`

SetKvmParams sets KvmParams field to given value.

### HasKvmParams

`func (o *ProtectionGroup) HasKvmParams() bool`

HasKvmParams returns a boolean if a field has been set.

### SetKvmParamsNil

`func (o *ProtectionGroup) SetKvmParamsNil(b bool)`

 SetKvmParamsNil sets the value for KvmParams to be an explicit nil

### UnsetKvmParams
`func (o *ProtectionGroup) UnsetKvmParams()`

UnsetKvmParams ensures that no value is present for KvmParams, not even an explicit nil
### GetMongodbParams

`func (o *ProtectionGroup) GetMongodbParams() MongoDBProtectionGroupParams`

GetMongodbParams returns the MongodbParams field if non-nil, zero value otherwise.

### GetMongodbParamsOk

`func (o *ProtectionGroup) GetMongodbParamsOk() (*MongoDBProtectionGroupParams, bool)`

GetMongodbParamsOk returns a tuple with the MongodbParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbParams

`func (o *ProtectionGroup) SetMongodbParams(v MongoDBProtectionGroupParams)`

SetMongodbParams sets MongodbParams field to given value.

### HasMongodbParams

`func (o *ProtectionGroup) HasMongodbParams() bool`

HasMongodbParams returns a boolean if a field has been set.

### GetMssqlParams

`func (o *ProtectionGroup) GetMssqlParams() MSSQLProtectionGroupParams`

GetMssqlParams returns the MssqlParams field if non-nil, zero value otherwise.

### GetMssqlParamsOk

`func (o *ProtectionGroup) GetMssqlParamsOk() (*MSSQLProtectionGroupParams, bool)`

GetMssqlParamsOk returns a tuple with the MssqlParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlParams

`func (o *ProtectionGroup) SetMssqlParams(v MSSQLProtectionGroupParams)`

SetMssqlParams sets MssqlParams field to given value.

### HasMssqlParams

`func (o *ProtectionGroup) HasMssqlParams() bool`

HasMssqlParams returns a boolean if a field has been set.

### GetNetappParams

`func (o *ProtectionGroup) GetNetappParams() NetappProtectionGroupParams`

GetNetappParams returns the NetappParams field if non-nil, zero value otherwise.

### GetNetappParamsOk

`func (o *ProtectionGroup) GetNetappParamsOk() (*NetappProtectionGroupParams, bool)`

GetNetappParamsOk returns a tuple with the NetappParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetappParams

`func (o *ProtectionGroup) SetNetappParams(v NetappProtectionGroupParams)`

SetNetappParams sets NetappParams field to given value.

### HasNetappParams

`func (o *ProtectionGroup) HasNetappParams() bool`

HasNetappParams returns a boolean if a field has been set.

### SetNetappParamsNil

`func (o *ProtectionGroup) SetNetappParamsNil(b bool)`

 SetNetappParamsNil sets the value for NetappParams to be an explicit nil

### UnsetNetappParams
`func (o *ProtectionGroup) UnsetNetappParams()`

UnsetNetappParams ensures that no value is present for NetappParams, not even an explicit nil
### GetNimbleParams

`func (o *ProtectionGroup) GetNimbleParams() NimbleProtectionGroupParams`

GetNimbleParams returns the NimbleParams field if non-nil, zero value otherwise.

### GetNimbleParamsOk

`func (o *ProtectionGroup) GetNimbleParamsOk() (*NimbleProtectionGroupParams, bool)`

GetNimbleParamsOk returns a tuple with the NimbleParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNimbleParams

`func (o *ProtectionGroup) SetNimbleParams(v NimbleProtectionGroupParams)`

SetNimbleParams sets NimbleParams field to given value.

### HasNimbleParams

`func (o *ProtectionGroup) HasNimbleParams() bool`

HasNimbleParams returns a boolean if a field has been set.

### SetNimbleParamsNil

`func (o *ProtectionGroup) SetNimbleParamsNil(b bool)`

 SetNimbleParamsNil sets the value for NimbleParams to be an explicit nil

### UnsetNimbleParams
`func (o *ProtectionGroup) UnsetNimbleParams()`

UnsetNimbleParams ensures that no value is present for NimbleParams, not even an explicit nil
### GetOffice365Params

`func (o *ProtectionGroup) GetOffice365Params() Office365ProtectionGroupParams`

GetOffice365Params returns the Office365Params field if non-nil, zero value otherwise.

### GetOffice365ParamsOk

`func (o *ProtectionGroup) GetOffice365ParamsOk() (*Office365ProtectionGroupParams, bool)`

GetOffice365ParamsOk returns a tuple with the Office365Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffice365Params

`func (o *ProtectionGroup) SetOffice365Params(v Office365ProtectionGroupParams)`

SetOffice365Params sets Office365Params field to given value.

### HasOffice365Params

`func (o *ProtectionGroup) HasOffice365Params() bool`

HasOffice365Params returns a boolean if a field has been set.

### GetOracleParams

`func (o *ProtectionGroup) GetOracleParams() OracleProtectionGroupParams`

GetOracleParams returns the OracleParams field if non-nil, zero value otherwise.

### GetOracleParamsOk

`func (o *ProtectionGroup) GetOracleParamsOk() (*OracleProtectionGroupParams, bool)`

GetOracleParamsOk returns a tuple with the OracleParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleParams

`func (o *ProtectionGroup) SetOracleParams(v OracleProtectionGroupParams)`

SetOracleParams sets OracleParams field to given value.

### HasOracleParams

`func (o *ProtectionGroup) HasOracleParams() bool`

HasOracleParams returns a boolean if a field has been set.

### GetPhysicalParams

`func (o *ProtectionGroup) GetPhysicalParams() PhysicalProtectionGroupParams`

GetPhysicalParams returns the PhysicalParams field if non-nil, zero value otherwise.

### GetPhysicalParamsOk

`func (o *ProtectionGroup) GetPhysicalParamsOk() (*PhysicalProtectionGroupParams, bool)`

GetPhysicalParamsOk returns a tuple with the PhysicalParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalParams

`func (o *ProtectionGroup) SetPhysicalParams(v PhysicalProtectionGroupParams)`

SetPhysicalParams sets PhysicalParams field to given value.

### HasPhysicalParams

`func (o *ProtectionGroup) HasPhysicalParams() bool`

HasPhysicalParams returns a boolean if a field has been set.

### GetPureParams

`func (o *ProtectionGroup) GetPureParams() PureProtectionGroupParams`

GetPureParams returns the PureParams field if non-nil, zero value otherwise.

### GetPureParamsOk

`func (o *ProtectionGroup) GetPureParamsOk() (*PureProtectionGroupParams, bool)`

GetPureParamsOk returns a tuple with the PureParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPureParams

`func (o *ProtectionGroup) SetPureParams(v PureProtectionGroupParams)`

SetPureParams sets PureParams field to given value.

### HasPureParams

`func (o *ProtectionGroup) HasPureParams() bool`

HasPureParams returns a boolean if a field has been set.

### SetPureParamsNil

`func (o *ProtectionGroup) SetPureParamsNil(b bool)`

 SetPureParamsNil sets the value for PureParams to be an explicit nil

### UnsetPureParams
`func (o *ProtectionGroup) UnsetPureParams()`

UnsetPureParams ensures that no value is present for PureParams, not even an explicit nil
### GetRemoteAdapterParams

`func (o *ProtectionGroup) GetRemoteAdapterParams() RemoteAdapterProtectionGroupParams`

GetRemoteAdapterParams returns the RemoteAdapterParams field if non-nil, zero value otherwise.

### GetRemoteAdapterParamsOk

`func (o *ProtectionGroup) GetRemoteAdapterParamsOk() (*RemoteAdapterProtectionGroupParams, bool)`

GetRemoteAdapterParamsOk returns a tuple with the RemoteAdapterParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAdapterParams

`func (o *ProtectionGroup) SetRemoteAdapterParams(v RemoteAdapterProtectionGroupParams)`

SetRemoteAdapterParams sets RemoteAdapterParams field to given value.

### HasRemoteAdapterParams

`func (o *ProtectionGroup) HasRemoteAdapterParams() bool`

HasRemoteAdapterParams returns a boolean if a field has been set.

### SetRemoteAdapterParamsNil

`func (o *ProtectionGroup) SetRemoteAdapterParamsNil(b bool)`

 SetRemoteAdapterParamsNil sets the value for RemoteAdapterParams to be an explicit nil

### UnsetRemoteAdapterParams
`func (o *ProtectionGroup) UnsetRemoteAdapterParams()`

UnsetRemoteAdapterParams ensures that no value is present for RemoteAdapterParams, not even an explicit nil
### GetSapHanaParams

`func (o *ProtectionGroup) GetSapHanaParams() SapHanaProtectionGroupParams`

GetSapHanaParams returns the SapHanaParams field if non-nil, zero value otherwise.

### GetSapHanaParamsOk

`func (o *ProtectionGroup) GetSapHanaParamsOk() (*SapHanaProtectionGroupParams, bool)`

GetSapHanaParamsOk returns a tuple with the SapHanaParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSapHanaParams

`func (o *ProtectionGroup) SetSapHanaParams(v SapHanaProtectionGroupParams)`

SetSapHanaParams sets SapHanaParams field to given value.

### HasSapHanaParams

`func (o *ProtectionGroup) HasSapHanaParams() bool`

HasSapHanaParams returns a boolean if a field has been set.

### GetSfdcParams

`func (o *ProtectionGroup) GetSfdcParams() SfdcProtectionGroupParams`

GetSfdcParams returns the SfdcParams field if non-nil, zero value otherwise.

### GetSfdcParamsOk

`func (o *ProtectionGroup) GetSfdcParamsOk() (*SfdcProtectionGroupParams, bool)`

GetSfdcParamsOk returns a tuple with the SfdcParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfdcParams

`func (o *ProtectionGroup) SetSfdcParams(v SfdcProtectionGroupParams)`

SetSfdcParams sets SfdcParams field to given value.

### HasSfdcParams

`func (o *ProtectionGroup) HasSfdcParams() bool`

HasSfdcParams returns a boolean if a field has been set.

### GetUdaParams

`func (o *ProtectionGroup) GetUdaParams() UdaProtectionGroupParams`

GetUdaParams returns the UdaParams field if non-nil, zero value otherwise.

### GetUdaParamsOk

`func (o *ProtectionGroup) GetUdaParamsOk() (*UdaProtectionGroupParams, bool)`

GetUdaParamsOk returns a tuple with the UdaParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdaParams

`func (o *ProtectionGroup) SetUdaParams(v UdaProtectionGroupParams)`

SetUdaParams sets UdaParams field to given value.

### HasUdaParams

`func (o *ProtectionGroup) HasUdaParams() bool`

HasUdaParams returns a boolean if a field has been set.

### GetViewParams

`func (o *ProtectionGroup) GetViewParams() ViewProtectionGroupParams`

GetViewParams returns the ViewParams field if non-nil, zero value otherwise.

### GetViewParamsOk

`func (o *ProtectionGroup) GetViewParamsOk() (*ViewProtectionGroupParams, bool)`

GetViewParamsOk returns a tuple with the ViewParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewParams

`func (o *ProtectionGroup) SetViewParams(v ViewProtectionGroupParams)`

SetViewParams sets ViewParams field to given value.

### HasViewParams

`func (o *ProtectionGroup) HasViewParams() bool`

HasViewParams returns a boolean if a field has been set.

### SetViewParamsNil

`func (o *ProtectionGroup) SetViewParamsNil(b bool)`

 SetViewParamsNil sets the value for ViewParams to be an explicit nil

### UnsetViewParams
`func (o *ProtectionGroup) UnsetViewParams()`

UnsetViewParams ensures that no value is present for ViewParams, not even an explicit nil
### GetVmwareParams

`func (o *ProtectionGroup) GetVmwareParams() VmwareProtectionGroupParams`

GetVmwareParams returns the VmwareParams field if non-nil, zero value otherwise.

### GetVmwareParamsOk

`func (o *ProtectionGroup) GetVmwareParamsOk() (*VmwareProtectionGroupParams, bool)`

GetVmwareParamsOk returns a tuple with the VmwareParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareParams

`func (o *ProtectionGroup) SetVmwareParams(v VmwareProtectionGroupParams)`

SetVmwareParams sets VmwareParams field to given value.

### HasVmwareParams

`func (o *ProtectionGroup) HasVmwareParams() bool`

HasVmwareParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


