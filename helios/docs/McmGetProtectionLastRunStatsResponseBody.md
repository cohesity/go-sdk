# McmGetProtectionLastRunStatsResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregatedStats** | Pointer to [**[]McmLastRunStatusStats**](McmLastRunStatusStats.md) | Specifies the aggregated status across all adapters for respective last run. | [optional] 
**EnvStats** | Pointer to [**[]McmLastRunEnvSpecificStatusStats**](McmLastRunEnvSpecificStatusStats.md) | Specifies the enviournment specific last run status stats. | [optional] 

## Methods

### NewMcmGetProtectionLastRunStatsResponseBody

`func NewMcmGetProtectionLastRunStatsResponseBody() *McmGetProtectionLastRunStatsResponseBody`

NewMcmGetProtectionLastRunStatsResponseBody instantiates a new McmGetProtectionLastRunStatsResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcmGetProtectionLastRunStatsResponseBodyWithDefaults

`func NewMcmGetProtectionLastRunStatsResponseBodyWithDefaults() *McmGetProtectionLastRunStatsResponseBody`

NewMcmGetProtectionLastRunStatsResponseBodyWithDefaults instantiates a new McmGetProtectionLastRunStatsResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregatedStats

`func (o *McmGetProtectionLastRunStatsResponseBody) GetAggregatedStats() []McmLastRunStatusStats`

GetAggregatedStats returns the AggregatedStats field if non-nil, zero value otherwise.

### GetAggregatedStatsOk

`func (o *McmGetProtectionLastRunStatsResponseBody) GetAggregatedStatsOk() (*[]McmLastRunStatusStats, bool)`

GetAggregatedStatsOk returns a tuple with the AggregatedStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregatedStats

`func (o *McmGetProtectionLastRunStatsResponseBody) SetAggregatedStats(v []McmLastRunStatusStats)`

SetAggregatedStats sets AggregatedStats field to given value.

### HasAggregatedStats

`func (o *McmGetProtectionLastRunStatsResponseBody) HasAggregatedStats() bool`

HasAggregatedStats returns a boolean if a field has been set.

### GetEnvStats

`func (o *McmGetProtectionLastRunStatsResponseBody) GetEnvStats() []McmLastRunEnvSpecificStatusStats`

GetEnvStats returns the EnvStats field if non-nil, zero value otherwise.

### GetEnvStatsOk

`func (o *McmGetProtectionLastRunStatsResponseBody) GetEnvStatsOk() (*[]McmLastRunEnvSpecificStatusStats, bool)`

GetEnvStatsOk returns a tuple with the EnvStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvStats

`func (o *McmGetProtectionLastRunStatsResponseBody) SetEnvStats(v []McmLastRunEnvSpecificStatusStats)`

SetEnvStats sets EnvStats field to given value.

### HasEnvStats

`func (o *McmGetProtectionLastRunStatsResponseBody) HasEnvStats() bool`

HasEnvStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


