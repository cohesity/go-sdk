# ProtectionRunsStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProtectionRunStatus** | Pointer to **NullableString** | Specifies the status of protection runs. | [optional] 
**ProtectionRunsCount** | Pointer to **NullableInt64** | Specifies the number of protection runs. | [optional] 

## Methods

### NewProtectionRunsStats

`func NewProtectionRunsStats() *ProtectionRunsStats`

NewProtectionRunsStats instantiates a new ProtectionRunsStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionRunsStatsWithDefaults

`func NewProtectionRunsStatsWithDefaults() *ProtectionRunsStats`

NewProtectionRunsStatsWithDefaults instantiates a new ProtectionRunsStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtectionRunStatus

`func (o *ProtectionRunsStats) GetProtectionRunStatus() string`

GetProtectionRunStatus returns the ProtectionRunStatus field if non-nil, zero value otherwise.

### GetProtectionRunStatusOk

`func (o *ProtectionRunsStats) GetProtectionRunStatusOk() (*string, bool)`

GetProtectionRunStatusOk returns a tuple with the ProtectionRunStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionRunStatus

`func (o *ProtectionRunsStats) SetProtectionRunStatus(v string)`

SetProtectionRunStatus sets ProtectionRunStatus field to given value.

### HasProtectionRunStatus

`func (o *ProtectionRunsStats) HasProtectionRunStatus() bool`

HasProtectionRunStatus returns a boolean if a field has been set.

### SetProtectionRunStatusNil

`func (o *ProtectionRunsStats) SetProtectionRunStatusNil(b bool)`

 SetProtectionRunStatusNil sets the value for ProtectionRunStatus to be an explicit nil

### UnsetProtectionRunStatus
`func (o *ProtectionRunsStats) UnsetProtectionRunStatus()`

UnsetProtectionRunStatus ensures that no value is present for ProtectionRunStatus, not even an explicit nil
### GetProtectionRunsCount

`func (o *ProtectionRunsStats) GetProtectionRunsCount() int64`

GetProtectionRunsCount returns the ProtectionRunsCount field if non-nil, zero value otherwise.

### GetProtectionRunsCountOk

`func (o *ProtectionRunsStats) GetProtectionRunsCountOk() (*int64, bool)`

GetProtectionRunsCountOk returns a tuple with the ProtectionRunsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionRunsCount

`func (o *ProtectionRunsStats) SetProtectionRunsCount(v int64)`

SetProtectionRunsCount sets ProtectionRunsCount field to given value.

### HasProtectionRunsCount

`func (o *ProtectionRunsStats) HasProtectionRunsCount() bool`

HasProtectionRunsCount returns a boolean if a field has been set.

### SetProtectionRunsCountNil

`func (o *ProtectionRunsStats) SetProtectionRunsCountNil(b bool)`

 SetProtectionRunsCountNil sets the value for ProtectionRunsCount to be an explicit nil

### UnsetProtectionRunsCount
`func (o *ProtectionRunsStats) UnsetProtectionRunsCount()`

UnsetProtectionRunsCount ensures that no value is present for ProtectionRunsCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


