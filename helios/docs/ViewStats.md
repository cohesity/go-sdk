# ViewStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt64** | Specifies the id of the View. | [optional] 
**DataUsageStats** | Pointer to [**DataUsageStats**](DataUsageStats.md) | Specifies the data usage metric of the data stored in this View. | [optional] 

## Methods

### NewViewStats

`func NewViewStats() *ViewStats`

NewViewStats instantiates a new ViewStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewStatsWithDefaults

`func NewViewStatsWithDefaults() *ViewStats`

NewViewStatsWithDefaults instantiates a new ViewStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ViewStats) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ViewStats) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ViewStats) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ViewStats) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ViewStats) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ViewStats) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetDataUsageStats

`func (o *ViewStats) GetDataUsageStats() DataUsageStats`

GetDataUsageStats returns the DataUsageStats field if non-nil, zero value otherwise.

### GetDataUsageStatsOk

`func (o *ViewStats) GetDataUsageStatsOk() (*DataUsageStats, bool)`

GetDataUsageStatsOk returns a tuple with the DataUsageStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataUsageStats

`func (o *ViewStats) SetDataUsageStats(v DataUsageStats)`

SetDataUsageStats sets DataUsageStats field to given value.

### HasDataUsageStats

`func (o *ViewStats) HasDataUsageStats() bool`

HasDataUsageStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


