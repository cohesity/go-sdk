# StatsTaskInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupGenericStats** | Pointer to [**BackupGenericStats**](BackupGenericStats.md) |  | [optional] 
**NasStats** | Pointer to [**BackupNasStats**](BackupNasStats.md) |  | [optional] 

## Methods

### NewStatsTaskInfo

`func NewStatsTaskInfo() *StatsTaskInfo`

NewStatsTaskInfo instantiates a new StatsTaskInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatsTaskInfoWithDefaults

`func NewStatsTaskInfoWithDefaults() *StatsTaskInfo`

NewStatsTaskInfoWithDefaults instantiates a new StatsTaskInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupGenericStats

`func (o *StatsTaskInfo) GetBackupGenericStats() BackupGenericStats`

GetBackupGenericStats returns the BackupGenericStats field if non-nil, zero value otherwise.

### GetBackupGenericStatsOk

`func (o *StatsTaskInfo) GetBackupGenericStatsOk() (*BackupGenericStats, bool)`

GetBackupGenericStatsOk returns a tuple with the BackupGenericStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupGenericStats

`func (o *StatsTaskInfo) SetBackupGenericStats(v BackupGenericStats)`

SetBackupGenericStats sets BackupGenericStats field to given value.

### HasBackupGenericStats

`func (o *StatsTaskInfo) HasBackupGenericStats() bool`

HasBackupGenericStats returns a boolean if a field has been set.

### GetNasStats

`func (o *StatsTaskInfo) GetNasStats() BackupNasStats`

GetNasStats returns the NasStats field if non-nil, zero value otherwise.

### GetNasStatsOk

`func (o *StatsTaskInfo) GetNasStatsOk() (*BackupNasStats, bool)`

GetNasStatsOk returns a tuple with the NasStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasStats

`func (o *StatsTaskInfo) SetNasStats(v BackupNasStats)`

SetNasStats sets NasStats field to given value.

### HasNasStats

`func (o *StatsTaskInfo) HasNasStats() bool`

HasNasStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


