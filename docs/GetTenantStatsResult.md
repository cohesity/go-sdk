# GetTenantStatsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cookie** | Pointer to **NullableString** | Specifies an opaque string to pass to get the next set of active opens. If null is returned, this response is the last set of active opens. | [optional] 
**StatsList** | Pointer to [**[]TenantStats**](TenantStats.md) | Specifies a list of tenant stats. | [optional] 

## Methods

### NewGetTenantStatsResult

`func NewGetTenantStatsResult() *GetTenantStatsResult`

NewGetTenantStatsResult instantiates a new GetTenantStatsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTenantStatsResultWithDefaults

`func NewGetTenantStatsResultWithDefaults() *GetTenantStatsResult`

NewGetTenantStatsResultWithDefaults instantiates a new GetTenantStatsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCookie

`func (o *GetTenantStatsResult) GetCookie() string`

GetCookie returns the Cookie field if non-nil, zero value otherwise.

### GetCookieOk

`func (o *GetTenantStatsResult) GetCookieOk() (*string, bool)`

GetCookieOk returns a tuple with the Cookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookie

`func (o *GetTenantStatsResult) SetCookie(v string)`

SetCookie sets Cookie field to given value.

### HasCookie

`func (o *GetTenantStatsResult) HasCookie() bool`

HasCookie returns a boolean if a field has been set.

### SetCookieNil

`func (o *GetTenantStatsResult) SetCookieNil(b bool)`

 SetCookieNil sets the value for Cookie to be an explicit nil

### UnsetCookie
`func (o *GetTenantStatsResult) UnsetCookie()`

UnsetCookie ensures that no value is present for Cookie, not even an explicit nil
### GetStatsList

`func (o *GetTenantStatsResult) GetStatsList() []TenantStats`

GetStatsList returns the StatsList field if non-nil, zero value otherwise.

### GetStatsListOk

`func (o *GetTenantStatsResult) GetStatsListOk() (*[]TenantStats, bool)`

GetStatsListOk returns a tuple with the StatsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsList

`func (o *GetTenantStatsResult) SetStatsList(v []TenantStats)`

SetStatsList sets StatsList field to given value.

### HasStatsList

`func (o *GetTenantStatsResult) HasStatsList() bool`

HasStatsList returns a boolean if a field has been set.

### SetStatsListNil

`func (o *GetTenantStatsResult) SetStatsListNil(b bool)`

 SetStatsListNil sets the value for StatsList to be an explicit nil

### UnsetStatsList
`func (o *GetTenantStatsResult) UnsetStatsList()`

UnsetStatsList ensures that no value is present for StatsList, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


