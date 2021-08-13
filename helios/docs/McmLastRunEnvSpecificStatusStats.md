# McmLastRunEnvSpecificStatusStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | Pointer to **NullableString** | Specifies the environment type of object. | [optional] 
**AggregatedStats** | Pointer to [**[]McmLastRunStatusStats**](McmLastRunStatusStats.md) | Specifies the aggregated status across all adapters for respective last run. | [optional] 

## Methods

### NewMcmLastRunEnvSpecificStatusStats

`func NewMcmLastRunEnvSpecificStatusStats() *McmLastRunEnvSpecificStatusStats`

NewMcmLastRunEnvSpecificStatusStats instantiates a new McmLastRunEnvSpecificStatusStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcmLastRunEnvSpecificStatusStatsWithDefaults

`func NewMcmLastRunEnvSpecificStatusStatsWithDefaults() *McmLastRunEnvSpecificStatusStats`

NewMcmLastRunEnvSpecificStatusStatsWithDefaults instantiates a new McmLastRunEnvSpecificStatusStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *McmLastRunEnvSpecificStatusStats) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *McmLastRunEnvSpecificStatusStats) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *McmLastRunEnvSpecificStatusStats) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *McmLastRunEnvSpecificStatusStats) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *McmLastRunEnvSpecificStatusStats) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *McmLastRunEnvSpecificStatusStats) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetAggregatedStats

`func (o *McmLastRunEnvSpecificStatusStats) GetAggregatedStats() []McmLastRunStatusStats`

GetAggregatedStats returns the AggregatedStats field if non-nil, zero value otherwise.

### GetAggregatedStatsOk

`func (o *McmLastRunEnvSpecificStatusStats) GetAggregatedStatsOk() (*[]McmLastRunStatusStats, bool)`

GetAggregatedStatsOk returns a tuple with the AggregatedStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregatedStats

`func (o *McmLastRunEnvSpecificStatusStats) SetAggregatedStats(v []McmLastRunStatusStats)`

SetAggregatedStats sets AggregatedStats field to given value.

### HasAggregatedStats

`func (o *McmLastRunEnvSpecificStatusStats) HasAggregatedStats() bool`

HasAggregatedStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


