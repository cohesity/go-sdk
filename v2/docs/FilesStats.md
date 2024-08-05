# FilesStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileSizeDistribution** | Pointer to [**[]FileCount**](FileCount.md) | Specifies the aggregated distribution by size of files stored in the Cohesity cluster. | [optional] 
**FilesStats** | Pointer to [**[]FilesStatsForEntity**](FilesStatsForEntity.md) | Specifies a list of file stats for entities. | [optional] 

## Methods

### NewFilesStats

`func NewFilesStats() *FilesStats`

NewFilesStats instantiates a new FilesStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesStatsWithDefaults

`func NewFilesStatsWithDefaults() *FilesStats`

NewFilesStatsWithDefaults instantiates a new FilesStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileSizeDistribution

`func (o *FilesStats) GetFileSizeDistribution() []FileCount`

GetFileSizeDistribution returns the FileSizeDistribution field if non-nil, zero value otherwise.

### GetFileSizeDistributionOk

`func (o *FilesStats) GetFileSizeDistributionOk() (*[]FileCount, bool)`

GetFileSizeDistributionOk returns a tuple with the FileSizeDistribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSizeDistribution

`func (o *FilesStats) SetFileSizeDistribution(v []FileCount)`

SetFileSizeDistribution sets FileSizeDistribution field to given value.

### HasFileSizeDistribution

`func (o *FilesStats) HasFileSizeDistribution() bool`

HasFileSizeDistribution returns a boolean if a field has been set.

### SetFileSizeDistributionNil

`func (o *FilesStats) SetFileSizeDistributionNil(b bool)`

 SetFileSizeDistributionNil sets the value for FileSizeDistribution to be an explicit nil

### UnsetFileSizeDistribution
`func (o *FilesStats) UnsetFileSizeDistribution()`

UnsetFileSizeDistribution ensures that no value is present for FileSizeDistribution, not even an explicit nil
### GetFilesStats

`func (o *FilesStats) GetFilesStats() []FilesStatsForEntity`

GetFilesStats returns the FilesStats field if non-nil, zero value otherwise.

### GetFilesStatsOk

`func (o *FilesStats) GetFilesStatsOk() (*[]FilesStatsForEntity, bool)`

GetFilesStatsOk returns a tuple with the FilesStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesStats

`func (o *FilesStats) SetFilesStats(v []FilesStatsForEntity)`

SetFilesStats sets FilesStats field to given value.

### HasFilesStats

`func (o *FilesStats) HasFilesStats() bool`

HasFilesStats returns a boolean if a field has been set.

### SetFilesStatsNil

`func (o *FilesStats) SetFilesStatsNil(b bool)`

 SetFilesStatsNil sets the value for FilesStats to be an explicit nil

### UnsetFilesStats
`func (o *FilesStats) UnsetFilesStats()`

UnsetFilesStats ensures that no value is present for FilesStats, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


