# UserQuotaOverrides

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cookie** | Pointer to **NullableString** | Specifies the pagination cookie. | [optional] 
**OverrideExistingPerUserQuotas** | Pointer to **NullableBool** | By default, the overrides specified in userQuotas is treated as delta and the existing overrides will be left untouched. Set this to true, if the existing overrides should be cleared before applying overrides specified in userQuotas. | [optional] 
**UserQuotas** | [**[]UserQuota**](UserQuota.md) | Array of UserQuota. Specifies the list of UserQuota for each user. | 

## Methods

### NewUserQuotaOverrides

`func NewUserQuotaOverrides(userQuotas []UserQuota, ) *UserQuotaOverrides`

NewUserQuotaOverrides instantiates a new UserQuotaOverrides object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserQuotaOverridesWithDefaults

`func NewUserQuotaOverridesWithDefaults() *UserQuotaOverrides`

NewUserQuotaOverridesWithDefaults instantiates a new UserQuotaOverrides object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCookie

`func (o *UserQuotaOverrides) GetCookie() string`

GetCookie returns the Cookie field if non-nil, zero value otherwise.

### GetCookieOk

`func (o *UserQuotaOverrides) GetCookieOk() (*string, bool)`

GetCookieOk returns a tuple with the Cookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookie

`func (o *UserQuotaOverrides) SetCookie(v string)`

SetCookie sets Cookie field to given value.

### HasCookie

`func (o *UserQuotaOverrides) HasCookie() bool`

HasCookie returns a boolean if a field has been set.

### SetCookieNil

`func (o *UserQuotaOverrides) SetCookieNil(b bool)`

 SetCookieNil sets the value for Cookie to be an explicit nil

### UnsetCookie
`func (o *UserQuotaOverrides) UnsetCookie()`

UnsetCookie ensures that no value is present for Cookie, not even an explicit nil
### GetOverrideExistingPerUserQuotas

`func (o *UserQuotaOverrides) GetOverrideExistingPerUserQuotas() bool`

GetOverrideExistingPerUserQuotas returns the OverrideExistingPerUserQuotas field if non-nil, zero value otherwise.

### GetOverrideExistingPerUserQuotasOk

`func (o *UserQuotaOverrides) GetOverrideExistingPerUserQuotasOk() (*bool, bool)`

GetOverrideExistingPerUserQuotasOk returns a tuple with the OverrideExistingPerUserQuotas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideExistingPerUserQuotas

`func (o *UserQuotaOverrides) SetOverrideExistingPerUserQuotas(v bool)`

SetOverrideExistingPerUserQuotas sets OverrideExistingPerUserQuotas field to given value.

### HasOverrideExistingPerUserQuotas

`func (o *UserQuotaOverrides) HasOverrideExistingPerUserQuotas() bool`

HasOverrideExistingPerUserQuotas returns a boolean if a field has been set.

### SetOverrideExistingPerUserQuotasNil

`func (o *UserQuotaOverrides) SetOverrideExistingPerUserQuotasNil(b bool)`

 SetOverrideExistingPerUserQuotasNil sets the value for OverrideExistingPerUserQuotas to be an explicit nil

### UnsetOverrideExistingPerUserQuotas
`func (o *UserQuotaOverrides) UnsetOverrideExistingPerUserQuotas()`

UnsetOverrideExistingPerUserQuotas ensures that no value is present for OverrideExistingPerUserQuotas, not even an explicit nil
### GetUserQuotas

`func (o *UserQuotaOverrides) GetUserQuotas() []UserQuota`

GetUserQuotas returns the UserQuotas field if non-nil, zero value otherwise.

### GetUserQuotasOk

`func (o *UserQuotaOverrides) GetUserQuotasOk() (*[]UserQuota, bool)`

GetUserQuotasOk returns a tuple with the UserQuotas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserQuotas

`func (o *UserQuotaOverrides) SetUserQuotas(v []UserQuota)`

SetUserQuotas sets UserQuotas field to given value.


### SetUserQuotasNil

`func (o *UserQuotaOverrides) SetUserQuotasNil(b bool)`

 SetUserQuotasNil sets the value for UserQuotas to be an explicit nil

### UnsetUserQuotas
`func (o *UserQuotaOverrides) UnsetUserQuotas()`

UnsetUserQuotas ensures that no value is present for UserQuotas, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


