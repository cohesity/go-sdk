# ObjectStatsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | Pointer to **NullableString** | Specifies the environment of the object. | [optional] 
**Id** | Pointer to **NullableInt64** | Specifies object id. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the object. | [optional] 
**SourceId** | Pointer to **NullableInt64** | Specifies registered source id to which object belongs. | [optional] 
**SourceName** | Pointer to **NullableString** | Specifies registered source name to which object belongs. | [optional] 
**BackupGenericStats** | Pointer to [**BackupGenericStats**](BackupGenericStats.md) |  | [optional] 
**NasStats** | Pointer to [**BackupNasStats**](BackupNasStats.md) |  | [optional] 
**FailedAttempts** | Pointer to [**[]StatsTaskInfo**](StatsTaskInfo.md) | Specifies stats for failed attempts of this object. | [optional] 

## Methods

### NewObjectStatsInfo

`func NewObjectStatsInfo() *ObjectStatsInfo`

NewObjectStatsInfo instantiates a new ObjectStatsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStatsInfoWithDefaults

`func NewObjectStatsInfoWithDefaults() *ObjectStatsInfo`

NewObjectStatsInfoWithDefaults instantiates a new ObjectStatsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *ObjectStatsInfo) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ObjectStatsInfo) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ObjectStatsInfo) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ObjectStatsInfo) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *ObjectStatsInfo) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *ObjectStatsInfo) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetId

`func (o *ObjectStatsInfo) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObjectStatsInfo) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObjectStatsInfo) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ObjectStatsInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ObjectStatsInfo) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ObjectStatsInfo) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *ObjectStatsInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ObjectStatsInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ObjectStatsInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ObjectStatsInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ObjectStatsInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ObjectStatsInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSourceId

`func (o *ObjectStatsInfo) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *ObjectStatsInfo) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *ObjectStatsInfo) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *ObjectStatsInfo) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *ObjectStatsInfo) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *ObjectStatsInfo) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetSourceName

`func (o *ObjectStatsInfo) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *ObjectStatsInfo) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *ObjectStatsInfo) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *ObjectStatsInfo) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *ObjectStatsInfo) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *ObjectStatsInfo) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil
### GetBackupGenericStats

`func (o *ObjectStatsInfo) GetBackupGenericStats() BackupGenericStats`

GetBackupGenericStats returns the BackupGenericStats field if non-nil, zero value otherwise.

### GetBackupGenericStatsOk

`func (o *ObjectStatsInfo) GetBackupGenericStatsOk() (*BackupGenericStats, bool)`

GetBackupGenericStatsOk returns a tuple with the BackupGenericStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupGenericStats

`func (o *ObjectStatsInfo) SetBackupGenericStats(v BackupGenericStats)`

SetBackupGenericStats sets BackupGenericStats field to given value.

### HasBackupGenericStats

`func (o *ObjectStatsInfo) HasBackupGenericStats() bool`

HasBackupGenericStats returns a boolean if a field has been set.

### GetNasStats

`func (o *ObjectStatsInfo) GetNasStats() BackupNasStats`

GetNasStats returns the NasStats field if non-nil, zero value otherwise.

### GetNasStatsOk

`func (o *ObjectStatsInfo) GetNasStatsOk() (*BackupNasStats, bool)`

GetNasStatsOk returns a tuple with the NasStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasStats

`func (o *ObjectStatsInfo) SetNasStats(v BackupNasStats)`

SetNasStats sets NasStats field to given value.

### HasNasStats

`func (o *ObjectStatsInfo) HasNasStats() bool`

HasNasStats returns a boolean if a field has been set.

### GetFailedAttempts

`func (o *ObjectStatsInfo) GetFailedAttempts() []StatsTaskInfo`

GetFailedAttempts returns the FailedAttempts field if non-nil, zero value otherwise.

### GetFailedAttemptsOk

`func (o *ObjectStatsInfo) GetFailedAttemptsOk() (*[]StatsTaskInfo, bool)`

GetFailedAttemptsOk returns a tuple with the FailedAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAttempts

`func (o *ObjectStatsInfo) SetFailedAttempts(v []StatsTaskInfo)`

SetFailedAttempts sets FailedAttempts field to given value.

### HasFailedAttempts

`func (o *ObjectStatsInfo) HasFailedAttempts() bool`

HasFailedAttempts returns a boolean if a field has been set.

### SetFailedAttemptsNil

`func (o *ObjectStatsInfo) SetFailedAttemptsNil(b bool)`

 SetFailedAttemptsNil sets the value for FailedAttempts to be an explicit nil

### UnsetFailedAttempts
`func (o *ObjectStatsInfo) UnsetFailedAttempts()`

UnsetFailedAttempts ensures that no value is present for FailedAttempts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


