# ViewUserQuotas

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cookie** | Pointer to **NullableString** | This cookie can be used in the succeeding call to list user quotas and usages to get the next set of user quota overrides. If set to nil, it means that there&#39;s no more results that the server could provide. | [optional] 
**QuotaAndUsageInAllViews** | Pointer to [**[]QuotaAndUsageInView**](QuotaAndUsageInView.md) | The quota and usage information for a user in all his views. | [optional] 
**SummaryForUser** | Pointer to [**UserQuotaSummaryForUser**](UserQuotaSummaryForUser.md) |  | [optional] 
**SummaryForView** | Pointer to [**UserQuotaSummaryForView**](UserQuotaSummaryForView.md) |  | [optional] 
**UserQuotaSettings** | Pointer to [**UserQuotaSettings**](UserQuotaSettings.md) |  | [optional] 
**UsersQuotaAndUsage** | Pointer to [**[]UserQuotaAndUsage**](UserQuotaAndUsage.md) | The list of user quota policies/overrides and usages. | [optional] 

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
### GetQuotaAndUsageInAllViews

`func (o *ViewUserQuotas) GetQuotaAndUsageInAllViews() []QuotaAndUsageInView`

GetQuotaAndUsageInAllViews returns the QuotaAndUsageInAllViews field if non-nil, zero value otherwise.

### GetQuotaAndUsageInAllViewsOk

`func (o *ViewUserQuotas) GetQuotaAndUsageInAllViewsOk() (*[]QuotaAndUsageInView, bool)`

GetQuotaAndUsageInAllViewsOk returns a tuple with the QuotaAndUsageInAllViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaAndUsageInAllViews

`func (o *ViewUserQuotas) SetQuotaAndUsageInAllViews(v []QuotaAndUsageInView)`

SetQuotaAndUsageInAllViews sets QuotaAndUsageInAllViews field to given value.

### HasQuotaAndUsageInAllViews

`func (o *ViewUserQuotas) HasQuotaAndUsageInAllViews() bool`

HasQuotaAndUsageInAllViews returns a boolean if a field has been set.

### SetQuotaAndUsageInAllViewsNil

`func (o *ViewUserQuotas) SetQuotaAndUsageInAllViewsNil(b bool)`

 SetQuotaAndUsageInAllViewsNil sets the value for QuotaAndUsageInAllViews to be an explicit nil

### UnsetQuotaAndUsageInAllViews
`func (o *ViewUserQuotas) UnsetQuotaAndUsageInAllViews()`

UnsetQuotaAndUsageInAllViews ensures that no value is present for QuotaAndUsageInAllViews, not even an explicit nil
### GetSummaryForUser

`func (o *ViewUserQuotas) GetSummaryForUser() UserQuotaSummaryForUser`

GetSummaryForUser returns the SummaryForUser field if non-nil, zero value otherwise.

### GetSummaryForUserOk

`func (o *ViewUserQuotas) GetSummaryForUserOk() (*UserQuotaSummaryForUser, bool)`

GetSummaryForUserOk returns a tuple with the SummaryForUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryForUser

`func (o *ViewUserQuotas) SetSummaryForUser(v UserQuotaSummaryForUser)`

SetSummaryForUser sets SummaryForUser field to given value.

### HasSummaryForUser

`func (o *ViewUserQuotas) HasSummaryForUser() bool`

HasSummaryForUser returns a boolean if a field has been set.

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

### GetUserQuotaSettings

`func (o *ViewUserQuotas) GetUserQuotaSettings() UserQuotaSettings`

GetUserQuotaSettings returns the UserQuotaSettings field if non-nil, zero value otherwise.

### GetUserQuotaSettingsOk

`func (o *ViewUserQuotas) GetUserQuotaSettingsOk() (*UserQuotaSettings, bool)`

GetUserQuotaSettingsOk returns a tuple with the UserQuotaSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserQuotaSettings

`func (o *ViewUserQuotas) SetUserQuotaSettings(v UserQuotaSettings)`

SetUserQuotaSettings sets UserQuotaSettings field to given value.

### HasUserQuotaSettings

`func (o *ViewUserQuotas) HasUserQuotaSettings() bool`

HasUserQuotaSettings returns a boolean if a field has been set.

### GetUsersQuotaAndUsage

`func (o *ViewUserQuotas) GetUsersQuotaAndUsage() []UserQuotaAndUsage`

GetUsersQuotaAndUsage returns the UsersQuotaAndUsage field if non-nil, zero value otherwise.

### GetUsersQuotaAndUsageOk

`func (o *ViewUserQuotas) GetUsersQuotaAndUsageOk() (*[]UserQuotaAndUsage, bool)`

GetUsersQuotaAndUsageOk returns a tuple with the UsersQuotaAndUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersQuotaAndUsage

`func (o *ViewUserQuotas) SetUsersQuotaAndUsage(v []UserQuotaAndUsage)`

SetUsersQuotaAndUsage sets UsersQuotaAndUsage field to given value.

### HasUsersQuotaAndUsage

`func (o *ViewUserQuotas) HasUsersQuotaAndUsage() bool`

HasUsersQuotaAndUsage returns a boolean if a field has been set.

### SetUsersQuotaAndUsageNil

`func (o *ViewUserQuotas) SetUsersQuotaAndUsageNil(b bool)`

 SetUsersQuotaAndUsageNil sets the value for UsersQuotaAndUsage to be an explicit nil

### UnsetUsersQuotaAndUsage
`func (o *ViewUserQuotas) UnsetUsersQuotaAndUsage()`

UnsetUsersQuotaAndUsage ensures that no value is present for UsersQuotaAndUsage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


