# BackupRunStatsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupGenericStats** | Pointer to [**BackupGenericStats**](BackupGenericStats.md) |  | [optional] 
**NasStats** | Pointer to [**BackupNasStats**](BackupNasStats.md) |  | [optional] 
**Objects** | Pointer to [**[]ObjectStatsInfo**](ObjectStatsInfo.md) | Specifies stats for objects. | [optional] 

## Methods

### NewBackupRunStatsInfo

`func NewBackupRunStatsInfo() *BackupRunStatsInfo`

NewBackupRunStatsInfo instantiates a new BackupRunStatsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRunStatsInfoWithDefaults

`func NewBackupRunStatsInfoWithDefaults() *BackupRunStatsInfo`

NewBackupRunStatsInfoWithDefaults instantiates a new BackupRunStatsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupGenericStats

`func (o *BackupRunStatsInfo) GetBackupGenericStats() BackupGenericStats`

GetBackupGenericStats returns the BackupGenericStats field if non-nil, zero value otherwise.

### GetBackupGenericStatsOk

`func (o *BackupRunStatsInfo) GetBackupGenericStatsOk() (*BackupGenericStats, bool)`

GetBackupGenericStatsOk returns a tuple with the BackupGenericStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupGenericStats

`func (o *BackupRunStatsInfo) SetBackupGenericStats(v BackupGenericStats)`

SetBackupGenericStats sets BackupGenericStats field to given value.

### HasBackupGenericStats

`func (o *BackupRunStatsInfo) HasBackupGenericStats() bool`

HasBackupGenericStats returns a boolean if a field has been set.

### GetNasStats

`func (o *BackupRunStatsInfo) GetNasStats() BackupNasStats`

GetNasStats returns the NasStats field if non-nil, zero value otherwise.

### GetNasStatsOk

`func (o *BackupRunStatsInfo) GetNasStatsOk() (*BackupNasStats, bool)`

GetNasStatsOk returns a tuple with the NasStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasStats

`func (o *BackupRunStatsInfo) SetNasStats(v BackupNasStats)`

SetNasStats sets NasStats field to given value.

### HasNasStats

`func (o *BackupRunStatsInfo) HasNasStats() bool`

HasNasStats returns a boolean if a field has been set.

### GetObjects

`func (o *BackupRunStatsInfo) GetObjects() []ObjectStatsInfo`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *BackupRunStatsInfo) GetObjectsOk() (*[]ObjectStatsInfo, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *BackupRunStatsInfo) SetObjects(v []ObjectStatsInfo)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *BackupRunStatsInfo) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### SetObjectsNil

`func (o *BackupRunStatsInfo) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *BackupRunStatsInfo) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


