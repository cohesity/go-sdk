# LastProtectionRunStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumRunsFailed** | Pointer to **NullableInt64** | Specifies the number of Protection Jobs for which specified Protection Run failed. | [optional] 
**NumRunsFailedSla** | Pointer to **NullableInt64** | Specifies the number of Protection Jobs for which specified Protection Run failed SLA. | [optional] 
**NumRunsMetSla** | Pointer to **NullableInt64** | Specifies the number of Protection Jobs for which specified Protection Run met SLA. | [optional] 
**StatsByEnv** | Pointer to [**[]LastProtectionRunStatsByEnv**](LastProtectionRunStatsByEnv.md) | Specifies the last Protection Run stats by environment. | [optional] 

## Methods

### NewLastProtectionRunStats

`func NewLastProtectionRunStats() *LastProtectionRunStats`

NewLastProtectionRunStats instantiates a new LastProtectionRunStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLastProtectionRunStatsWithDefaults

`func NewLastProtectionRunStatsWithDefaults() *LastProtectionRunStats`

NewLastProtectionRunStatsWithDefaults instantiates a new LastProtectionRunStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumRunsFailed

`func (o *LastProtectionRunStats) GetNumRunsFailed() int64`

GetNumRunsFailed returns the NumRunsFailed field if non-nil, zero value otherwise.

### GetNumRunsFailedOk

`func (o *LastProtectionRunStats) GetNumRunsFailedOk() (*int64, bool)`

GetNumRunsFailedOk returns a tuple with the NumRunsFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumRunsFailed

`func (o *LastProtectionRunStats) SetNumRunsFailed(v int64)`

SetNumRunsFailed sets NumRunsFailed field to given value.

### HasNumRunsFailed

`func (o *LastProtectionRunStats) HasNumRunsFailed() bool`

HasNumRunsFailed returns a boolean if a field has been set.

### SetNumRunsFailedNil

`func (o *LastProtectionRunStats) SetNumRunsFailedNil(b bool)`

 SetNumRunsFailedNil sets the value for NumRunsFailed to be an explicit nil

### UnsetNumRunsFailed
`func (o *LastProtectionRunStats) UnsetNumRunsFailed()`

UnsetNumRunsFailed ensures that no value is present for NumRunsFailed, not even an explicit nil
### GetNumRunsFailedSla

`func (o *LastProtectionRunStats) GetNumRunsFailedSla() int64`

GetNumRunsFailedSla returns the NumRunsFailedSla field if non-nil, zero value otherwise.

### GetNumRunsFailedSlaOk

`func (o *LastProtectionRunStats) GetNumRunsFailedSlaOk() (*int64, bool)`

GetNumRunsFailedSlaOk returns a tuple with the NumRunsFailedSla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumRunsFailedSla

`func (o *LastProtectionRunStats) SetNumRunsFailedSla(v int64)`

SetNumRunsFailedSla sets NumRunsFailedSla field to given value.

### HasNumRunsFailedSla

`func (o *LastProtectionRunStats) HasNumRunsFailedSla() bool`

HasNumRunsFailedSla returns a boolean if a field has been set.

### SetNumRunsFailedSlaNil

`func (o *LastProtectionRunStats) SetNumRunsFailedSlaNil(b bool)`

 SetNumRunsFailedSlaNil sets the value for NumRunsFailedSla to be an explicit nil

### UnsetNumRunsFailedSla
`func (o *LastProtectionRunStats) UnsetNumRunsFailedSla()`

UnsetNumRunsFailedSla ensures that no value is present for NumRunsFailedSla, not even an explicit nil
### GetNumRunsMetSla

`func (o *LastProtectionRunStats) GetNumRunsMetSla() int64`

GetNumRunsMetSla returns the NumRunsMetSla field if non-nil, zero value otherwise.

### GetNumRunsMetSlaOk

`func (o *LastProtectionRunStats) GetNumRunsMetSlaOk() (*int64, bool)`

GetNumRunsMetSlaOk returns a tuple with the NumRunsMetSla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumRunsMetSla

`func (o *LastProtectionRunStats) SetNumRunsMetSla(v int64)`

SetNumRunsMetSla sets NumRunsMetSla field to given value.

### HasNumRunsMetSla

`func (o *LastProtectionRunStats) HasNumRunsMetSla() bool`

HasNumRunsMetSla returns a boolean if a field has been set.

### SetNumRunsMetSlaNil

`func (o *LastProtectionRunStats) SetNumRunsMetSlaNil(b bool)`

 SetNumRunsMetSlaNil sets the value for NumRunsMetSla to be an explicit nil

### UnsetNumRunsMetSla
`func (o *LastProtectionRunStats) UnsetNumRunsMetSla()`

UnsetNumRunsMetSla ensures that no value is present for NumRunsMetSla, not even an explicit nil
### GetStatsByEnv

`func (o *LastProtectionRunStats) GetStatsByEnv() []LastProtectionRunStatsByEnv`

GetStatsByEnv returns the StatsByEnv field if non-nil, zero value otherwise.

### GetStatsByEnvOk

`func (o *LastProtectionRunStats) GetStatsByEnvOk() (*[]LastProtectionRunStatsByEnv, bool)`

GetStatsByEnvOk returns a tuple with the StatsByEnv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsByEnv

`func (o *LastProtectionRunStats) SetStatsByEnv(v []LastProtectionRunStatsByEnv)`

SetStatsByEnv sets StatsByEnv field to given value.

### HasStatsByEnv

`func (o *LastProtectionRunStats) HasStatsByEnv() bool`

HasStatsByEnv returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


