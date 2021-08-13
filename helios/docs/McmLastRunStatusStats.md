# McmLastRunStatusStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **NullableString** | Specifies the status of last run. | [optional] 
**Count** | Pointer to **NullableInt64** | Specifies the number of objects having this status. | [optional] 

## Methods

### NewMcmLastRunStatusStats

`func NewMcmLastRunStatusStats() *McmLastRunStatusStats`

NewMcmLastRunStatusStats instantiates a new McmLastRunStatusStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcmLastRunStatusStatsWithDefaults

`func NewMcmLastRunStatusStatsWithDefaults() *McmLastRunStatusStats`

NewMcmLastRunStatusStatsWithDefaults instantiates a new McmLastRunStatusStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *McmLastRunStatusStats) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *McmLastRunStatusStats) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *McmLastRunStatusStats) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *McmLastRunStatusStats) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *McmLastRunStatusStats) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *McmLastRunStatusStats) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetCount

`func (o *McmLastRunStatusStats) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *McmLastRunStatusStats) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *McmLastRunStatusStats) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *McmLastRunStatusStats) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *McmLastRunStatusStats) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *McmLastRunStatusStats) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


