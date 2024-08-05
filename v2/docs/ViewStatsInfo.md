# ViewStatsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocols** | Pointer to **[]string** | Specifies the protocols of this view. | [optional] 
**Stats** | Pointer to [**[]ViewStatsInfoDetails**](ViewStatsInfoDetails.md) | Specifies the list of View stats. | [optional] 
**ViewId** | Pointer to **NullableInt64** | Specifies the view Id. | [optional] 
**ViewName** | Pointer to **NullableString** | Specifies the view name. | [optional] 

## Methods

### NewViewStatsInfo

`func NewViewStatsInfo() *ViewStatsInfo`

NewViewStatsInfo instantiates a new ViewStatsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewStatsInfoWithDefaults

`func NewViewStatsInfoWithDefaults() *ViewStatsInfo`

NewViewStatsInfoWithDefaults instantiates a new ViewStatsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocols

`func (o *ViewStatsInfo) GetProtocols() []string`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *ViewStatsInfo) GetProtocolsOk() (*[]string, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *ViewStatsInfo) SetProtocols(v []string)`

SetProtocols sets Protocols field to given value.

### HasProtocols

`func (o *ViewStatsInfo) HasProtocols() bool`

HasProtocols returns a boolean if a field has been set.

### SetProtocolsNil

`func (o *ViewStatsInfo) SetProtocolsNil(b bool)`

 SetProtocolsNil sets the value for Protocols to be an explicit nil

### UnsetProtocols
`func (o *ViewStatsInfo) UnsetProtocols()`

UnsetProtocols ensures that no value is present for Protocols, not even an explicit nil
### GetStats

`func (o *ViewStatsInfo) GetStats() []ViewStatsInfoDetails`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ViewStatsInfo) GetStatsOk() (*[]ViewStatsInfoDetails, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ViewStatsInfo) SetStats(v []ViewStatsInfoDetails)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ViewStatsInfo) HasStats() bool`

HasStats returns a boolean if a field has been set.

### SetStatsNil

`func (o *ViewStatsInfo) SetStatsNil(b bool)`

 SetStatsNil sets the value for Stats to be an explicit nil

### UnsetStats
`func (o *ViewStatsInfo) UnsetStats()`

UnsetStats ensures that no value is present for Stats, not even an explicit nil
### GetViewId

`func (o *ViewStatsInfo) GetViewId() int64`

GetViewId returns the ViewId field if non-nil, zero value otherwise.

### GetViewIdOk

`func (o *ViewStatsInfo) GetViewIdOk() (*int64, bool)`

GetViewIdOk returns a tuple with the ViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewId

`func (o *ViewStatsInfo) SetViewId(v int64)`

SetViewId sets ViewId field to given value.

### HasViewId

`func (o *ViewStatsInfo) HasViewId() bool`

HasViewId returns a boolean if a field has been set.

### SetViewIdNil

`func (o *ViewStatsInfo) SetViewIdNil(b bool)`

 SetViewIdNil sets the value for ViewId to be an explicit nil

### UnsetViewId
`func (o *ViewStatsInfo) UnsetViewId()`

UnsetViewId ensures that no value is present for ViewId, not even an explicit nil
### GetViewName

`func (o *ViewStatsInfo) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *ViewStatsInfo) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *ViewStatsInfo) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *ViewStatsInfo) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### SetViewNameNil

`func (o *ViewStatsInfo) SetViewNameNil(b bool)`

 SetViewNameNil sets the value for ViewName to be an explicit nil

### UnsetViewName
`func (o *ViewStatsInfo) UnsetViewName()`

UnsetViewName ensures that no value is present for ViewName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


