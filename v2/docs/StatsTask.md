# StatsTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupGenericStats** | Pointer to [**BackupGenericStats**](BackupGenericStats.md) |  | [optional] 
**NasStats** | Pointer to [**BackupNasStats**](BackupNasStats.md) |  | [optional] 
**Id** | Pointer to **NullableString** | Specifies the task id of the Stats task. | [optional] 

## Methods

### NewStatsTask

`func NewStatsTask() *StatsTask`

NewStatsTask instantiates a new StatsTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatsTaskWithDefaults

`func NewStatsTaskWithDefaults() *StatsTask`

NewStatsTaskWithDefaults instantiates a new StatsTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupGenericStats

`func (o *StatsTask) GetBackupGenericStats() BackupGenericStats`

GetBackupGenericStats returns the BackupGenericStats field if non-nil, zero value otherwise.

### GetBackupGenericStatsOk

`func (o *StatsTask) GetBackupGenericStatsOk() (*BackupGenericStats, bool)`

GetBackupGenericStatsOk returns a tuple with the BackupGenericStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupGenericStats

`func (o *StatsTask) SetBackupGenericStats(v BackupGenericStats)`

SetBackupGenericStats sets BackupGenericStats field to given value.

### HasBackupGenericStats

`func (o *StatsTask) HasBackupGenericStats() bool`

HasBackupGenericStats returns a boolean if a field has been set.

### GetNasStats

`func (o *StatsTask) GetNasStats() BackupNasStats`

GetNasStats returns the NasStats field if non-nil, zero value otherwise.

### GetNasStatsOk

`func (o *StatsTask) GetNasStatsOk() (*BackupNasStats, bool)`

GetNasStatsOk returns a tuple with the NasStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasStats

`func (o *StatsTask) SetNasStats(v BackupNasStats)`

SetNasStats sets NasStats field to given value.

### HasNasStats

`func (o *StatsTask) HasNasStats() bool`

HasNasStats returns a boolean if a field has been set.

### GetId

`func (o *StatsTask) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StatsTask) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StatsTask) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StatsTask) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *StatsTask) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *StatsTask) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


