# ViewClientStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientIp** | Pointer to **NullableString** | Specifies the client IP. | [optional] 
**Stats** | Pointer to [**[]ClientStats**](ClientStats.md) | Specifies the list of Client stats. | [optional] 

## Methods

### NewViewClientStats

`func NewViewClientStats() *ViewClientStats`

NewViewClientStats instantiates a new ViewClientStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewClientStatsWithDefaults

`func NewViewClientStatsWithDefaults() *ViewClientStats`

NewViewClientStatsWithDefaults instantiates a new ViewClientStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientIp

`func (o *ViewClientStats) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *ViewClientStats) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *ViewClientStats) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *ViewClientStats) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### SetClientIpNil

`func (o *ViewClientStats) SetClientIpNil(b bool)`

 SetClientIpNil sets the value for ClientIp to be an explicit nil

### UnsetClientIp
`func (o *ViewClientStats) UnsetClientIp()`

UnsetClientIp ensures that no value is present for ClientIp, not even an explicit nil
### GetStats

`func (o *ViewClientStats) GetStats() []ClientStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ViewClientStats) GetStatsOk() (*[]ClientStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ViewClientStats) SetStats(v []ClientStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ViewClientStats) HasStats() bool`

HasStats returns a boolean if a field has been set.

### SetStatsNil

`func (o *ViewClientStats) SetStatsNil(b bool)`

 SetStatsNil sets the value for Stats to be an explicit nil

### UnsetStats
`func (o *ViewClientStats) UnsetStats()`

UnsetStats ensures that no value is present for Stats, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


