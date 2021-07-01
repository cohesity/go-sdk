# GetViewBoxStatsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatsList** | Pointer to [**[]StorageDomainStats**](StorageDomainStats.md) | Specifies a list of view box stats. | [optional] 

## Methods

### NewGetViewBoxStatsResult

`func NewGetViewBoxStatsResult() *GetViewBoxStatsResult`

NewGetViewBoxStatsResult instantiates a new GetViewBoxStatsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetViewBoxStatsResultWithDefaults

`func NewGetViewBoxStatsResultWithDefaults() *GetViewBoxStatsResult`

NewGetViewBoxStatsResultWithDefaults instantiates a new GetViewBoxStatsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatsList

`func (o *GetViewBoxStatsResult) GetStatsList() []StorageDomainStats`

GetStatsList returns the StatsList field if non-nil, zero value otherwise.

### GetStatsListOk

`func (o *GetViewBoxStatsResult) GetStatsListOk() (*[]StorageDomainStats, bool)`

GetStatsListOk returns a tuple with the StatsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsList

`func (o *GetViewBoxStatsResult) SetStatsList(v []StorageDomainStats)`

SetStatsList sets StatsList field to given value.

### HasStatsList

`func (o *GetViewBoxStatsResult) HasStatsList() bool`

HasStatsList returns a boolean if a field has been set.

### SetStatsListNil

`func (o *GetViewBoxStatsResult) SetStatsListNil(b bool)`

 SetStatsListNil sets the value for StatsList to be an explicit nil

### UnsetStatsList
`func (o *GetViewBoxStatsResult) UnsetStatsList()`

UnsetStatsList ensures that no value is present for StatsList, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


