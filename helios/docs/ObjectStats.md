# ObjectStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SnapshotsSummary** | Pointer to [**[]SnapshotsSummary**](SnapshotsSummary.md) | Specifies a summary of the object snapshots. | [optional] 

## Methods

### NewObjectStats

`func NewObjectStats() *ObjectStats`

NewObjectStats instantiates a new ObjectStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStatsWithDefaults

`func NewObjectStatsWithDefaults() *ObjectStats`

NewObjectStatsWithDefaults instantiates a new ObjectStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshotsSummary

`func (o *ObjectStats) GetSnapshotsSummary() []SnapshotsSummary`

GetSnapshotsSummary returns the SnapshotsSummary field if non-nil, zero value otherwise.

### GetSnapshotsSummaryOk

`func (o *ObjectStats) GetSnapshotsSummaryOk() (*[]SnapshotsSummary, bool)`

GetSnapshotsSummaryOk returns a tuple with the SnapshotsSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotsSummary

`func (o *ObjectStats) SetSnapshotsSummary(v []SnapshotsSummary)`

SetSnapshotsSummary sets SnapshotsSummary field to given value.

### HasSnapshotsSummary

`func (o *ObjectStats) HasSnapshotsSummary() bool`

HasSnapshotsSummary returns a boolean if a field has been set.

### SetSnapshotsSummaryNil

`func (o *ObjectStats) SetSnapshotsSummaryNil(b bool)`

 SetSnapshotsSummaryNil sets the value for SnapshotsSummary to be an explicit nil

### UnsetSnapshotsSummary
`func (o *ObjectStats) UnsetSnapshotsSummary()`

UnsetSnapshotsSummary ensures that no value is present for SnapshotsSummary, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


