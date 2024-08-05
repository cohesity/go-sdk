# ArchivalTargetStatsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchivalTaskId** | Pointer to **NullableString** | Specifies the archival task id. This is a protection group UID which only applies when archival type is &#39;Tape&#39;. | [optional] 
**OwnershipContext** | Pointer to **NullableString** | Specifies the ownership context for the target. | [optional] 
**TargetId** | Pointer to **NullableInt64** | Specifies the archival target ID. | [optional] 
**TargetName** | Pointer to **NullableString** | Specifies the archival target name. | [optional] 
**TargetType** | Pointer to **NullableString** | Specifies the archival target type. | [optional] 
**TierSettings** | Pointer to [**ArchivalTargetTierInfo**](ArchivalTargetTierInfo.md) |  | [optional] 
**UsageType** | Pointer to **NullableString** | Specifies the usage type for the target. | [optional] 
**BackupGenericStats** | Pointer to [**BackupGenericStats**](BackupGenericStats.md) |  | [optional] 
**NasStats** | Pointer to [**BackupNasStats**](BackupNasStats.md) |  | [optional] 
**Objects** | Pointer to [**[]ObjectStatsInfo**](ObjectStatsInfo.md) | Specifies stats for objects. | [optional] 

## Methods

### NewArchivalTargetStatsInfo

`func NewArchivalTargetStatsInfo() *ArchivalTargetStatsInfo`

NewArchivalTargetStatsInfo instantiates a new ArchivalTargetStatsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArchivalTargetStatsInfoWithDefaults

`func NewArchivalTargetStatsInfoWithDefaults() *ArchivalTargetStatsInfo`

NewArchivalTargetStatsInfoWithDefaults instantiates a new ArchivalTargetStatsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchivalTaskId

`func (o *ArchivalTargetStatsInfo) GetArchivalTaskId() string`

GetArchivalTaskId returns the ArchivalTaskId field if non-nil, zero value otherwise.

### GetArchivalTaskIdOk

`func (o *ArchivalTargetStatsInfo) GetArchivalTaskIdOk() (*string, bool)`

GetArchivalTaskIdOk returns a tuple with the ArchivalTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivalTaskId

`func (o *ArchivalTargetStatsInfo) SetArchivalTaskId(v string)`

SetArchivalTaskId sets ArchivalTaskId field to given value.

### HasArchivalTaskId

`func (o *ArchivalTargetStatsInfo) HasArchivalTaskId() bool`

HasArchivalTaskId returns a boolean if a field has been set.

### SetArchivalTaskIdNil

`func (o *ArchivalTargetStatsInfo) SetArchivalTaskIdNil(b bool)`

 SetArchivalTaskIdNil sets the value for ArchivalTaskId to be an explicit nil

### UnsetArchivalTaskId
`func (o *ArchivalTargetStatsInfo) UnsetArchivalTaskId()`

UnsetArchivalTaskId ensures that no value is present for ArchivalTaskId, not even an explicit nil
### GetOwnershipContext

`func (o *ArchivalTargetStatsInfo) GetOwnershipContext() string`

GetOwnershipContext returns the OwnershipContext field if non-nil, zero value otherwise.

### GetOwnershipContextOk

`func (o *ArchivalTargetStatsInfo) GetOwnershipContextOk() (*string, bool)`

GetOwnershipContextOk returns a tuple with the OwnershipContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnershipContext

`func (o *ArchivalTargetStatsInfo) SetOwnershipContext(v string)`

SetOwnershipContext sets OwnershipContext field to given value.

### HasOwnershipContext

`func (o *ArchivalTargetStatsInfo) HasOwnershipContext() bool`

HasOwnershipContext returns a boolean if a field has been set.

### SetOwnershipContextNil

`func (o *ArchivalTargetStatsInfo) SetOwnershipContextNil(b bool)`

 SetOwnershipContextNil sets the value for OwnershipContext to be an explicit nil

### UnsetOwnershipContext
`func (o *ArchivalTargetStatsInfo) UnsetOwnershipContext()`

UnsetOwnershipContext ensures that no value is present for OwnershipContext, not even an explicit nil
### GetTargetId

`func (o *ArchivalTargetStatsInfo) GetTargetId() int64`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *ArchivalTargetStatsInfo) GetTargetIdOk() (*int64, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *ArchivalTargetStatsInfo) SetTargetId(v int64)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *ArchivalTargetStatsInfo) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### SetTargetIdNil

`func (o *ArchivalTargetStatsInfo) SetTargetIdNil(b bool)`

 SetTargetIdNil sets the value for TargetId to be an explicit nil

### UnsetTargetId
`func (o *ArchivalTargetStatsInfo) UnsetTargetId()`

UnsetTargetId ensures that no value is present for TargetId, not even an explicit nil
### GetTargetName

`func (o *ArchivalTargetStatsInfo) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *ArchivalTargetStatsInfo) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *ArchivalTargetStatsInfo) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *ArchivalTargetStatsInfo) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### SetTargetNameNil

`func (o *ArchivalTargetStatsInfo) SetTargetNameNil(b bool)`

 SetTargetNameNil sets the value for TargetName to be an explicit nil

### UnsetTargetName
`func (o *ArchivalTargetStatsInfo) UnsetTargetName()`

UnsetTargetName ensures that no value is present for TargetName, not even an explicit nil
### GetTargetType

`func (o *ArchivalTargetStatsInfo) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *ArchivalTargetStatsInfo) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *ArchivalTargetStatsInfo) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *ArchivalTargetStatsInfo) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### SetTargetTypeNil

`func (o *ArchivalTargetStatsInfo) SetTargetTypeNil(b bool)`

 SetTargetTypeNil sets the value for TargetType to be an explicit nil

### UnsetTargetType
`func (o *ArchivalTargetStatsInfo) UnsetTargetType()`

UnsetTargetType ensures that no value is present for TargetType, not even an explicit nil
### GetTierSettings

`func (o *ArchivalTargetStatsInfo) GetTierSettings() ArchivalTargetTierInfo`

GetTierSettings returns the TierSettings field if non-nil, zero value otherwise.

### GetTierSettingsOk

`func (o *ArchivalTargetStatsInfo) GetTierSettingsOk() (*ArchivalTargetTierInfo, bool)`

GetTierSettingsOk returns a tuple with the TierSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierSettings

`func (o *ArchivalTargetStatsInfo) SetTierSettings(v ArchivalTargetTierInfo)`

SetTierSettings sets TierSettings field to given value.

### HasTierSettings

`func (o *ArchivalTargetStatsInfo) HasTierSettings() bool`

HasTierSettings returns a boolean if a field has been set.

### GetUsageType

`func (o *ArchivalTargetStatsInfo) GetUsageType() string`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *ArchivalTargetStatsInfo) GetUsageTypeOk() (*string, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *ArchivalTargetStatsInfo) SetUsageType(v string)`

SetUsageType sets UsageType field to given value.

### HasUsageType

`func (o *ArchivalTargetStatsInfo) HasUsageType() bool`

HasUsageType returns a boolean if a field has been set.

### SetUsageTypeNil

`func (o *ArchivalTargetStatsInfo) SetUsageTypeNil(b bool)`

 SetUsageTypeNil sets the value for UsageType to be an explicit nil

### UnsetUsageType
`func (o *ArchivalTargetStatsInfo) UnsetUsageType()`

UnsetUsageType ensures that no value is present for UsageType, not even an explicit nil
### GetBackupGenericStats

`func (o *ArchivalTargetStatsInfo) GetBackupGenericStats() BackupGenericStats`

GetBackupGenericStats returns the BackupGenericStats field if non-nil, zero value otherwise.

### GetBackupGenericStatsOk

`func (o *ArchivalTargetStatsInfo) GetBackupGenericStatsOk() (*BackupGenericStats, bool)`

GetBackupGenericStatsOk returns a tuple with the BackupGenericStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupGenericStats

`func (o *ArchivalTargetStatsInfo) SetBackupGenericStats(v BackupGenericStats)`

SetBackupGenericStats sets BackupGenericStats field to given value.

### HasBackupGenericStats

`func (o *ArchivalTargetStatsInfo) HasBackupGenericStats() bool`

HasBackupGenericStats returns a boolean if a field has been set.

### GetNasStats

`func (o *ArchivalTargetStatsInfo) GetNasStats() BackupNasStats`

GetNasStats returns the NasStats field if non-nil, zero value otherwise.

### GetNasStatsOk

`func (o *ArchivalTargetStatsInfo) GetNasStatsOk() (*BackupNasStats, bool)`

GetNasStatsOk returns a tuple with the NasStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasStats

`func (o *ArchivalTargetStatsInfo) SetNasStats(v BackupNasStats)`

SetNasStats sets NasStats field to given value.

### HasNasStats

`func (o *ArchivalTargetStatsInfo) HasNasStats() bool`

HasNasStats returns a boolean if a field has been set.

### GetObjects

`func (o *ArchivalTargetStatsInfo) GetObjects() []ObjectStatsInfo`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *ArchivalTargetStatsInfo) GetObjectsOk() (*[]ObjectStatsInfo, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *ArchivalTargetStatsInfo) SetObjects(v []ObjectStatsInfo)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *ArchivalTargetStatsInfo) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### SetObjectsNil

`func (o *ArchivalTargetStatsInfo) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *ArchivalTargetStatsInfo) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


