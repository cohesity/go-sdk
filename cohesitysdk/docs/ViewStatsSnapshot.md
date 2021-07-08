# ViewStatsSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **NullableInt64** | Specifies the unix time in milliseconds when these values were generated | [optional] 
**ViewStatsList** | Pointer to [**[]ViewStatInfo**](ViewStatInfo.md) | Specifies the list of Views and their statistics at the given timestamp. | [optional] 

## Methods

### NewViewStatsSnapshot

`func NewViewStatsSnapshot() *ViewStatsSnapshot`

NewViewStatsSnapshot instantiates a new ViewStatsSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewStatsSnapshotWithDefaults

`func NewViewStatsSnapshotWithDefaults() *ViewStatsSnapshot`

NewViewStatsSnapshotWithDefaults instantiates a new ViewStatsSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *ViewStatsSnapshot) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ViewStatsSnapshot) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ViewStatsSnapshot) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ViewStatsSnapshot) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ViewStatsSnapshot) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ViewStatsSnapshot) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetViewStatsList

`func (o *ViewStatsSnapshot) GetViewStatsList() []ViewStatInfo`

GetViewStatsList returns the ViewStatsList field if non-nil, zero value otherwise.

### GetViewStatsListOk

`func (o *ViewStatsSnapshot) GetViewStatsListOk() (*[]ViewStatInfo, bool)`

GetViewStatsListOk returns a tuple with the ViewStatsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewStatsList

`func (o *ViewStatsSnapshot) SetViewStatsList(v []ViewStatInfo)`

SetViewStatsList sets ViewStatsList field to given value.

### HasViewStatsList

`func (o *ViewStatsSnapshot) HasViewStatsList() bool`

HasViewStatsList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


