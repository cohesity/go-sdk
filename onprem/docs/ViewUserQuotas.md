# ViewUserQuotas

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserQuotas** | Pointer to [**[]ViewUserQuota**](ViewUserQuota.md) | Specifies the list of View user quotas. | [optional] 
**Cookie** | Pointer to **NullableString** | Specifies the pagination cookie. | [optional] 

## Methods

### NewViewUserQuotas

`func NewViewUserQuotas() *ViewUserQuotas`

NewViewUserQuotas instantiates a new ViewUserQuotas object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewUserQuotasWithDefaults

`func NewViewUserQuotasWithDefaults() *ViewUserQuotas`

NewViewUserQuotasWithDefaults instantiates a new ViewUserQuotas object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserQuotas

`func (o *ViewUserQuotas) GetUserQuotas() []ViewUserQuota`

GetUserQuotas returns the UserQuotas field if non-nil, zero value otherwise.

### GetUserQuotasOk

`func (o *ViewUserQuotas) GetUserQuotasOk() (*[]ViewUserQuota, bool)`

GetUserQuotasOk returns a tuple with the UserQuotas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserQuotas

`func (o *ViewUserQuotas) SetUserQuotas(v []ViewUserQuota)`

SetUserQuotas sets UserQuotas field to given value.

### HasUserQuotas

`func (o *ViewUserQuotas) HasUserQuotas() bool`

HasUserQuotas returns a boolean if a field has been set.

### SetUserQuotasNil

`func (o *ViewUserQuotas) SetUserQuotasNil(b bool)`

 SetUserQuotasNil sets the value for UserQuotas to be an explicit nil

### UnsetUserQuotas
`func (o *ViewUserQuotas) UnsetUserQuotas()`

UnsetUserQuotas ensures that no value is present for UserQuotas, not even an explicit nil
### GetCookie

`func (o *ViewUserQuotas) GetCookie() string`

GetCookie returns the Cookie field if non-nil, zero value otherwise.

### GetCookieOk

`func (o *ViewUserQuotas) GetCookieOk() (*string, bool)`

GetCookieOk returns a tuple with the Cookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookie

`func (o *ViewUserQuotas) SetCookie(v string)`

SetCookie sets Cookie field to given value.

### HasCookie

`func (o *ViewUserQuotas) HasCookie() bool`

HasCookie returns a boolean if a field has been set.

### SetCookieNil

`func (o *ViewUserQuotas) SetCookieNil(b bool)`

 SetCookieNil sets the value for Cookie to be an explicit nil

### UnsetCookie
`func (o *ViewUserQuotas) UnsetCookie()`

UnsetCookie ensures that no value is present for Cookie, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


