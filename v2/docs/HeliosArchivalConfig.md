# HeliosArchivalConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigId** | Pointer to **NullableString** | Specifies the unique identifier for the target getting added. This field need to be passed only when helios policies are updated. | [optional] 
**CopyOnRunSuccess** | Pointer to **NullableBool** | Specifies if Snapshots are copied from the first completely successful Protection Group Run or the first partially successful Protection Group Run occurring at the start of the replication schedule. &lt;br&gt; If true, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule that was completely successful i.e. Snapshots for all the Objects in the Protection Group were successfully captured. &lt;br&gt; If false, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule, even if first Protection Group Run was not completely successful i.e. Snapshots were not captured for all Objects in the Protection Group. | [optional] 
**Retention** | Pointer to [**HeliosRetention**](HeliosRetention.md) |  | [optional] 
**Schedule** | Pointer to [**HeliosTargetSchedule**](HeliosTargetSchedule.md) |  | [optional] 
**ExtendedRetention** | Pointer to [**[]HeliosExtendedRetentionPolicy**](HeliosExtendedRetentionPolicy.md) | Specifies additional retention policies that should be applied to the archived backup. Archived backup snapshot will be retained up to a time that is the maximum of all retention policies that are applicable to it. | [optional] 
**TargetId** | **NullableInt64** | Specifies the Archival target to copy the Snapshots to. | 
**TargetName** | Pointer to **NullableString** | Specifies the Archival target name where Snapshots are copied. | [optional] [readonly] 
**TargetType** | Pointer to **NullableString** | Specifies the Archival target type where Snapshots are copied. | [optional] [readonly] 
**TierSettings** | Pointer to [**HeliosTierLevelSettings**](HeliosTierLevelSettings.md) |  | [optional] 

## Methods

### NewHeliosArchivalConfig

`func NewHeliosArchivalConfig(targetId NullableInt64, ) *HeliosArchivalConfig`

NewHeliosArchivalConfig instantiates a new HeliosArchivalConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosArchivalConfigWithDefaults

`func NewHeliosArchivalConfigWithDefaults() *HeliosArchivalConfig`

NewHeliosArchivalConfigWithDefaults instantiates a new HeliosArchivalConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigId

`func (o *HeliosArchivalConfig) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *HeliosArchivalConfig) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *HeliosArchivalConfig) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.

### HasConfigId

`func (o *HeliosArchivalConfig) HasConfigId() bool`

HasConfigId returns a boolean if a field has been set.

### SetConfigIdNil

`func (o *HeliosArchivalConfig) SetConfigIdNil(b bool)`

 SetConfigIdNil sets the value for ConfigId to be an explicit nil

### UnsetConfigId
`func (o *HeliosArchivalConfig) UnsetConfigId()`

UnsetConfigId ensures that no value is present for ConfigId, not even an explicit nil
### GetCopyOnRunSuccess

`func (o *HeliosArchivalConfig) GetCopyOnRunSuccess() bool`

GetCopyOnRunSuccess returns the CopyOnRunSuccess field if non-nil, zero value otherwise.

### GetCopyOnRunSuccessOk

`func (o *HeliosArchivalConfig) GetCopyOnRunSuccessOk() (*bool, bool)`

GetCopyOnRunSuccessOk returns a tuple with the CopyOnRunSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyOnRunSuccess

`func (o *HeliosArchivalConfig) SetCopyOnRunSuccess(v bool)`

SetCopyOnRunSuccess sets CopyOnRunSuccess field to given value.

### HasCopyOnRunSuccess

`func (o *HeliosArchivalConfig) HasCopyOnRunSuccess() bool`

HasCopyOnRunSuccess returns a boolean if a field has been set.

### SetCopyOnRunSuccessNil

`func (o *HeliosArchivalConfig) SetCopyOnRunSuccessNil(b bool)`

 SetCopyOnRunSuccessNil sets the value for CopyOnRunSuccess to be an explicit nil

### UnsetCopyOnRunSuccess
`func (o *HeliosArchivalConfig) UnsetCopyOnRunSuccess()`

UnsetCopyOnRunSuccess ensures that no value is present for CopyOnRunSuccess, not even an explicit nil
### GetRetention

`func (o *HeliosArchivalConfig) GetRetention() HeliosRetention`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *HeliosArchivalConfig) GetRetentionOk() (*HeliosRetention, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *HeliosArchivalConfig) SetRetention(v HeliosRetention)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *HeliosArchivalConfig) HasRetention() bool`

HasRetention returns a boolean if a field has been set.

### GetSchedule

`func (o *HeliosArchivalConfig) GetSchedule() HeliosTargetSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *HeliosArchivalConfig) GetScheduleOk() (*HeliosTargetSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *HeliosArchivalConfig) SetSchedule(v HeliosTargetSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *HeliosArchivalConfig) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetExtendedRetention

`func (o *HeliosArchivalConfig) GetExtendedRetention() []HeliosExtendedRetentionPolicy`

GetExtendedRetention returns the ExtendedRetention field if non-nil, zero value otherwise.

### GetExtendedRetentionOk

`func (o *HeliosArchivalConfig) GetExtendedRetentionOk() (*[]HeliosExtendedRetentionPolicy, bool)`

GetExtendedRetentionOk returns a tuple with the ExtendedRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedRetention

`func (o *HeliosArchivalConfig) SetExtendedRetention(v []HeliosExtendedRetentionPolicy)`

SetExtendedRetention sets ExtendedRetention field to given value.

### HasExtendedRetention

`func (o *HeliosArchivalConfig) HasExtendedRetention() bool`

HasExtendedRetention returns a boolean if a field has been set.

### SetExtendedRetentionNil

`func (o *HeliosArchivalConfig) SetExtendedRetentionNil(b bool)`

 SetExtendedRetentionNil sets the value for ExtendedRetention to be an explicit nil

### UnsetExtendedRetention
`func (o *HeliosArchivalConfig) UnsetExtendedRetention()`

UnsetExtendedRetention ensures that no value is present for ExtendedRetention, not even an explicit nil
### GetTargetId

`func (o *HeliosArchivalConfig) GetTargetId() int64`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *HeliosArchivalConfig) GetTargetIdOk() (*int64, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *HeliosArchivalConfig) SetTargetId(v int64)`

SetTargetId sets TargetId field to given value.


### SetTargetIdNil

`func (o *HeliosArchivalConfig) SetTargetIdNil(b bool)`

 SetTargetIdNil sets the value for TargetId to be an explicit nil

### UnsetTargetId
`func (o *HeliosArchivalConfig) UnsetTargetId()`

UnsetTargetId ensures that no value is present for TargetId, not even an explicit nil
### GetTargetName

`func (o *HeliosArchivalConfig) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *HeliosArchivalConfig) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *HeliosArchivalConfig) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *HeliosArchivalConfig) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### SetTargetNameNil

`func (o *HeliosArchivalConfig) SetTargetNameNil(b bool)`

 SetTargetNameNil sets the value for TargetName to be an explicit nil

### UnsetTargetName
`func (o *HeliosArchivalConfig) UnsetTargetName()`

UnsetTargetName ensures that no value is present for TargetName, not even an explicit nil
### GetTargetType

`func (o *HeliosArchivalConfig) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *HeliosArchivalConfig) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *HeliosArchivalConfig) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *HeliosArchivalConfig) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### SetTargetTypeNil

`func (o *HeliosArchivalConfig) SetTargetTypeNil(b bool)`

 SetTargetTypeNil sets the value for TargetType to be an explicit nil

### UnsetTargetType
`func (o *HeliosArchivalConfig) UnsetTargetType()`

UnsetTargetType ensures that no value is present for TargetType, not even an explicit nil
### GetTierSettings

`func (o *HeliosArchivalConfig) GetTierSettings() HeliosTierLevelSettings`

GetTierSettings returns the TierSettings field if non-nil, zero value otherwise.

### GetTierSettingsOk

`func (o *HeliosArchivalConfig) GetTierSettingsOk() (*HeliosTierLevelSettings, bool)`

GetTierSettingsOk returns a tuple with the TierSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierSettings

`func (o *HeliosArchivalConfig) SetTierSettings(v HeliosTierLevelSettings)`

SetTierSettings sets TierSettings field to given value.

### HasTierSettings

`func (o *HeliosArchivalConfig) HasTierSettings() bool`

HasTierSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


