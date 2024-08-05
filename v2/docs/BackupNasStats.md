# BackupNasStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileAnalysisRate** | Pointer to **NullableInt64** | Specifies the rate at which files are being analyzed in files per minute. | [optional] 
**FileDiscoveryRate** | Pointer to **NullableInt64** | Specifies the rate at which files are being discovered in files per minute. | [optional] 
**FileDiscoveryTime** | Pointer to **NullableInt64** | Specifies the time taken for file discovery. | [optional] 
**FileIngestionRate** | Pointer to **NullableInt64** | Specifies the rate at which files are being ingested in files per minute. | [optional] 
**FilesAnalyzed** | Pointer to **NullableInt64** | Specifies the number of files which have been analyzed. | [optional] 
**FilesDiscovered** | Pointer to **NullableInt64** | Specifies the number of files which have already been discovered. | [optional] 
**FilesIngested** | Pointer to **NullableInt64** | Specifies the number of files which have been ingested. | [optional] 
**WalkerRunTime** | Pointer to **NullableInt64** | Specifies the run time for directory walker in seconds. | [optional] 

## Methods

### NewBackupNasStats

`func NewBackupNasStats() *BackupNasStats`

NewBackupNasStats instantiates a new BackupNasStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupNasStatsWithDefaults

`func NewBackupNasStatsWithDefaults() *BackupNasStats`

NewBackupNasStatsWithDefaults instantiates a new BackupNasStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileAnalysisRate

`func (o *BackupNasStats) GetFileAnalysisRate() int64`

GetFileAnalysisRate returns the FileAnalysisRate field if non-nil, zero value otherwise.

### GetFileAnalysisRateOk

`func (o *BackupNasStats) GetFileAnalysisRateOk() (*int64, bool)`

GetFileAnalysisRateOk returns a tuple with the FileAnalysisRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileAnalysisRate

`func (o *BackupNasStats) SetFileAnalysisRate(v int64)`

SetFileAnalysisRate sets FileAnalysisRate field to given value.

### HasFileAnalysisRate

`func (o *BackupNasStats) HasFileAnalysisRate() bool`

HasFileAnalysisRate returns a boolean if a field has been set.

### SetFileAnalysisRateNil

`func (o *BackupNasStats) SetFileAnalysisRateNil(b bool)`

 SetFileAnalysisRateNil sets the value for FileAnalysisRate to be an explicit nil

### UnsetFileAnalysisRate
`func (o *BackupNasStats) UnsetFileAnalysisRate()`

UnsetFileAnalysisRate ensures that no value is present for FileAnalysisRate, not even an explicit nil
### GetFileDiscoveryRate

`func (o *BackupNasStats) GetFileDiscoveryRate() int64`

GetFileDiscoveryRate returns the FileDiscoveryRate field if non-nil, zero value otherwise.

### GetFileDiscoveryRateOk

`func (o *BackupNasStats) GetFileDiscoveryRateOk() (*int64, bool)`

GetFileDiscoveryRateOk returns a tuple with the FileDiscoveryRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileDiscoveryRate

`func (o *BackupNasStats) SetFileDiscoveryRate(v int64)`

SetFileDiscoveryRate sets FileDiscoveryRate field to given value.

### HasFileDiscoveryRate

`func (o *BackupNasStats) HasFileDiscoveryRate() bool`

HasFileDiscoveryRate returns a boolean if a field has been set.

### SetFileDiscoveryRateNil

`func (o *BackupNasStats) SetFileDiscoveryRateNil(b bool)`

 SetFileDiscoveryRateNil sets the value for FileDiscoveryRate to be an explicit nil

### UnsetFileDiscoveryRate
`func (o *BackupNasStats) UnsetFileDiscoveryRate()`

UnsetFileDiscoveryRate ensures that no value is present for FileDiscoveryRate, not even an explicit nil
### GetFileDiscoveryTime

`func (o *BackupNasStats) GetFileDiscoveryTime() int64`

GetFileDiscoveryTime returns the FileDiscoveryTime field if non-nil, zero value otherwise.

### GetFileDiscoveryTimeOk

`func (o *BackupNasStats) GetFileDiscoveryTimeOk() (*int64, bool)`

GetFileDiscoveryTimeOk returns a tuple with the FileDiscoveryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileDiscoveryTime

`func (o *BackupNasStats) SetFileDiscoveryTime(v int64)`

SetFileDiscoveryTime sets FileDiscoveryTime field to given value.

### HasFileDiscoveryTime

`func (o *BackupNasStats) HasFileDiscoveryTime() bool`

HasFileDiscoveryTime returns a boolean if a field has been set.

### SetFileDiscoveryTimeNil

`func (o *BackupNasStats) SetFileDiscoveryTimeNil(b bool)`

 SetFileDiscoveryTimeNil sets the value for FileDiscoveryTime to be an explicit nil

### UnsetFileDiscoveryTime
`func (o *BackupNasStats) UnsetFileDiscoveryTime()`

UnsetFileDiscoveryTime ensures that no value is present for FileDiscoveryTime, not even an explicit nil
### GetFileIngestionRate

`func (o *BackupNasStats) GetFileIngestionRate() int64`

GetFileIngestionRate returns the FileIngestionRate field if non-nil, zero value otherwise.

### GetFileIngestionRateOk

`func (o *BackupNasStats) GetFileIngestionRateOk() (*int64, bool)`

GetFileIngestionRateOk returns a tuple with the FileIngestionRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileIngestionRate

`func (o *BackupNasStats) SetFileIngestionRate(v int64)`

SetFileIngestionRate sets FileIngestionRate field to given value.

### HasFileIngestionRate

`func (o *BackupNasStats) HasFileIngestionRate() bool`

HasFileIngestionRate returns a boolean if a field has been set.

### SetFileIngestionRateNil

`func (o *BackupNasStats) SetFileIngestionRateNil(b bool)`

 SetFileIngestionRateNil sets the value for FileIngestionRate to be an explicit nil

### UnsetFileIngestionRate
`func (o *BackupNasStats) UnsetFileIngestionRate()`

UnsetFileIngestionRate ensures that no value is present for FileIngestionRate, not even an explicit nil
### GetFilesAnalyzed

`func (o *BackupNasStats) GetFilesAnalyzed() int64`

GetFilesAnalyzed returns the FilesAnalyzed field if non-nil, zero value otherwise.

### GetFilesAnalyzedOk

`func (o *BackupNasStats) GetFilesAnalyzedOk() (*int64, bool)`

GetFilesAnalyzedOk returns a tuple with the FilesAnalyzed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesAnalyzed

`func (o *BackupNasStats) SetFilesAnalyzed(v int64)`

SetFilesAnalyzed sets FilesAnalyzed field to given value.

### HasFilesAnalyzed

`func (o *BackupNasStats) HasFilesAnalyzed() bool`

HasFilesAnalyzed returns a boolean if a field has been set.

### SetFilesAnalyzedNil

`func (o *BackupNasStats) SetFilesAnalyzedNil(b bool)`

 SetFilesAnalyzedNil sets the value for FilesAnalyzed to be an explicit nil

### UnsetFilesAnalyzed
`func (o *BackupNasStats) UnsetFilesAnalyzed()`

UnsetFilesAnalyzed ensures that no value is present for FilesAnalyzed, not even an explicit nil
### GetFilesDiscovered

`func (o *BackupNasStats) GetFilesDiscovered() int64`

GetFilesDiscovered returns the FilesDiscovered field if non-nil, zero value otherwise.

### GetFilesDiscoveredOk

`func (o *BackupNasStats) GetFilesDiscoveredOk() (*int64, bool)`

GetFilesDiscoveredOk returns a tuple with the FilesDiscovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesDiscovered

`func (o *BackupNasStats) SetFilesDiscovered(v int64)`

SetFilesDiscovered sets FilesDiscovered field to given value.

### HasFilesDiscovered

`func (o *BackupNasStats) HasFilesDiscovered() bool`

HasFilesDiscovered returns a boolean if a field has been set.

### SetFilesDiscoveredNil

`func (o *BackupNasStats) SetFilesDiscoveredNil(b bool)`

 SetFilesDiscoveredNil sets the value for FilesDiscovered to be an explicit nil

### UnsetFilesDiscovered
`func (o *BackupNasStats) UnsetFilesDiscovered()`

UnsetFilesDiscovered ensures that no value is present for FilesDiscovered, not even an explicit nil
### GetFilesIngested

`func (o *BackupNasStats) GetFilesIngested() int64`

GetFilesIngested returns the FilesIngested field if non-nil, zero value otherwise.

### GetFilesIngestedOk

`func (o *BackupNasStats) GetFilesIngestedOk() (*int64, bool)`

GetFilesIngestedOk returns a tuple with the FilesIngested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesIngested

`func (o *BackupNasStats) SetFilesIngested(v int64)`

SetFilesIngested sets FilesIngested field to given value.

### HasFilesIngested

`func (o *BackupNasStats) HasFilesIngested() bool`

HasFilesIngested returns a boolean if a field has been set.

### SetFilesIngestedNil

`func (o *BackupNasStats) SetFilesIngestedNil(b bool)`

 SetFilesIngestedNil sets the value for FilesIngested to be an explicit nil

### UnsetFilesIngested
`func (o *BackupNasStats) UnsetFilesIngested()`

UnsetFilesIngested ensures that no value is present for FilesIngested, not even an explicit nil
### GetWalkerRunTime

`func (o *BackupNasStats) GetWalkerRunTime() int64`

GetWalkerRunTime returns the WalkerRunTime field if non-nil, zero value otherwise.

### GetWalkerRunTimeOk

`func (o *BackupNasStats) GetWalkerRunTimeOk() (*int64, bool)`

GetWalkerRunTimeOk returns a tuple with the WalkerRunTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalkerRunTime

`func (o *BackupNasStats) SetWalkerRunTime(v int64)`

SetWalkerRunTime sets WalkerRunTime field to given value.

### HasWalkerRunTime

`func (o *BackupNasStats) HasWalkerRunTime() bool`

HasWalkerRunTime returns a boolean if a field has been set.

### SetWalkerRunTimeNil

`func (o *BackupNasStats) SetWalkerRunTimeNil(b bool)`

 SetWalkerRunTimeNil sets the value for WalkerRunTime to be an explicit nil

### UnsetWalkerRunTime
`func (o *BackupNasStats) UnsetWalkerRunTime()`

UnsetWalkerRunTime ensures that no value is present for WalkerRunTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


