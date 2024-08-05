# ViewUserQuotas

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultQuotaPolicy** | Pointer to [**ViewUserQuotaSettingsDefaultQuotaPolicy**](ViewUserQuotaSettingsDefaultQuotaPolicy.md) |  | [optional] 
**Enabled** | **bool** | Specifies whether user quota is enabled for the View. | 
**Cookie** | Pointer to **NullableString** | Specifies the pagination cookie. | [optional] 
**OverrideExistingPerUserQuotas** | Pointer to **NullableBool** | By default, the overrides specified in userQuotas is treated as delta and the existing overrides will be left untouched. Set this to true, if the existing overrides should be cleared before applying overrides specified in userQuotas. | [optional] 
**UserQuotas** | [**[]UserQuota**](UserQuota.md) | Array of UserQuota. Specifies the list of UserQuota for each user. | 
**SummaryForView** | Pointer to [**UserQuotaSummaryForView**](UserQuotaSummaryForView.md) |  | [optional] 

## Methods

### NewViewUserQuotas

`func NewViewUserQuotas(enabled bool, userQuotas []UserQuota, ) *ViewUserQuotas`

NewViewUserQuotas instantiates a new ViewUserQuotas object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewUserQuotasWithDefaults

`func NewViewUserQuotasWithDefaults() *ViewUserQuotas`

NewViewUserQuotasWithDefaults instantiates a new ViewUserQuotas object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultQuotaPolicy

`func (o *ViewUserQuotas) GetDefaultQuotaPolicy() ViewUserQuotaSettingsDefaultQuotaPolicy`

GetDefaultQuotaPolicy returns the DefaultQuotaPolicy field if non-nil, zero value otherwise.

### GetDefaultQuotaPolicyOk

`func (o *ViewUserQuotas) GetDefaultQuotaPolicyOk() (*ViewUserQuotaSettingsDefaultQuotaPolicy, bool)`

GetDefaultQuotaPolicyOk returns a tuple with the DefaultQuotaPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultQuotaPolicy

`func (o *ViewUserQuotas) SetDefaultQuotaPolicy(v ViewUserQuotaSettingsDefaultQuotaPolicy)`

SetDefaultQuotaPolicy sets DefaultQuotaPolicy field to given value.

### HasDefaultQuotaPolicy

`func (o *ViewUserQuotas) HasDefaultQuotaPolicy() bool`

HasDefaultQuotaPolicy returns a boolean if a field has been set.

### GetEnabled

`func (o *ViewUserQuotas) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ViewUserQuotas) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ViewUserQuotas) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


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
### GetOverrideExistingPerUserQuotas

`func (o *ViewUserQuotas) GetOverrideExistingPerUserQuotas() bool`

GetOverrideExistingPerUserQuotas returns the OverrideExistingPerUserQuotas field if non-nil, zero value otherwise.

### GetOverrideExistingPerUserQuotasOk

`func (o *ViewUserQuotas) GetOverrideExistingPerUserQuotasOk() (*bool, bool)`

GetOverrideExistingPerUserQuotasOk returns a tuple with the OverrideExistingPerUserQuotas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideExistingPerUserQuotas

`func (o *ViewUserQuotas) SetOverrideExistingPerUserQuotas(v bool)`

SetOverrideExistingPerUserQuotas sets OverrideExistingPerUserQuotas field to given value.

### HasOverrideExistingPerUserQuotas

`func (o *ViewUserQuotas) HasOverrideExistingPerUserQuotas() bool`

HasOverrideExistingPerUserQuotas returns a boolean if a field has been set.

### SetOverrideExistingPerUserQuotasNil

`func (o *ViewUserQuotas) SetOverrideExistingPerUserQuotasNil(b bool)`

 SetOverrideExistingPerUserQuotasNil sets the value for OverrideExistingPerUserQuotas to be an explicit nil

### UnsetOverrideExistingPerUserQuotas
`func (o *ViewUserQuotas) UnsetOverrideExistingPerUserQuotas()`

UnsetOverrideExistingPerUserQuotas ensures that no value is present for OverrideExistingPerUserQuotas, not even an explicit nil
### GetUserQuotas

`func (o *ViewUserQuotas) GetUserQuotas() []UserQuota`

GetUserQuotas returns the UserQuotas field if non-nil, zero value otherwise.

### GetUserQuotasOk

`func (o *ViewUserQuotas) GetUserQuotasOk() (*[]UserQuota, bool)`

GetUserQuotasOk returns a tuple with the UserQuotas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserQuotas

`func (o *ViewUserQuotas) SetUserQuotas(v []UserQuota)`

SetUserQuotas sets UserQuotas field to given value.


### SetUserQuotasNil

`func (o *ViewUserQuotas) SetUserQuotasNil(b bool)`

 SetUserQuotasNil sets the value for UserQuotas to be an explicit nil

### UnsetUserQuotas
`func (o *ViewUserQuotas) UnsetUserQuotas()`

UnsetUserQuotas ensures that no value is present for UserQuotas, not even an explicit nil
### GetSummaryForView

`func (o *ViewUserQuotas) GetSummaryForView() UserQuotaSummaryForView`

GetSummaryForView returns the SummaryForView field if non-nil, zero value otherwise.

### GetSummaryForViewOk

`func (o *ViewUserQuotas) GetSummaryForViewOk() (*UserQuotaSummaryForView, bool)`

GetSummaryForViewOk returns a tuple with the SummaryForView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryForView

`func (o *ViewUserQuotas) SetSummaryForView(v UserQuotaSummaryForView)`

SetSummaryForView sets SummaryForView field to given value.

### HasSummaryForView

`func (o *ViewUserQuotas) HasSummaryForView() bool`

HasSummaryForView returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


