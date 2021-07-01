# UserQuotaSummaryForUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumViewsAboveAlertThreshold** | Pointer to **NullableInt32** | Number of views in which user has exceeded alert threshold limit. | [optional] 
**NumViewsAboveHardLimit** | Pointer to **NullableInt32** | Number of views in which the user has exceeded hard limit. | [optional] 
**TotalNumViews** | Pointer to **NullableInt32** | Total number of views in which the user has a quota policy specified or has non-zero usage. | [optional] 

## Methods

### NewUserQuotaSummaryForUser

`func NewUserQuotaSummaryForUser() *UserQuotaSummaryForUser`

NewUserQuotaSummaryForUser instantiates a new UserQuotaSummaryForUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserQuotaSummaryForUserWithDefaults

`func NewUserQuotaSummaryForUserWithDefaults() *UserQuotaSummaryForUser`

NewUserQuotaSummaryForUserWithDefaults instantiates a new UserQuotaSummaryForUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumViewsAboveAlertThreshold

`func (o *UserQuotaSummaryForUser) GetNumViewsAboveAlertThreshold() int32`

GetNumViewsAboveAlertThreshold returns the NumViewsAboveAlertThreshold field if non-nil, zero value otherwise.

### GetNumViewsAboveAlertThresholdOk

`func (o *UserQuotaSummaryForUser) GetNumViewsAboveAlertThresholdOk() (*int32, bool)`

GetNumViewsAboveAlertThresholdOk returns a tuple with the NumViewsAboveAlertThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumViewsAboveAlertThreshold

`func (o *UserQuotaSummaryForUser) SetNumViewsAboveAlertThreshold(v int32)`

SetNumViewsAboveAlertThreshold sets NumViewsAboveAlertThreshold field to given value.

### HasNumViewsAboveAlertThreshold

`func (o *UserQuotaSummaryForUser) HasNumViewsAboveAlertThreshold() bool`

HasNumViewsAboveAlertThreshold returns a boolean if a field has been set.

### SetNumViewsAboveAlertThresholdNil

`func (o *UserQuotaSummaryForUser) SetNumViewsAboveAlertThresholdNil(b bool)`

 SetNumViewsAboveAlertThresholdNil sets the value for NumViewsAboveAlertThreshold to be an explicit nil

### UnsetNumViewsAboveAlertThreshold
`func (o *UserQuotaSummaryForUser) UnsetNumViewsAboveAlertThreshold()`

UnsetNumViewsAboveAlertThreshold ensures that no value is present for NumViewsAboveAlertThreshold, not even an explicit nil
### GetNumViewsAboveHardLimit

`func (o *UserQuotaSummaryForUser) GetNumViewsAboveHardLimit() int32`

GetNumViewsAboveHardLimit returns the NumViewsAboveHardLimit field if non-nil, zero value otherwise.

### GetNumViewsAboveHardLimitOk

`func (o *UserQuotaSummaryForUser) GetNumViewsAboveHardLimitOk() (*int32, bool)`

GetNumViewsAboveHardLimitOk returns a tuple with the NumViewsAboveHardLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumViewsAboveHardLimit

`func (o *UserQuotaSummaryForUser) SetNumViewsAboveHardLimit(v int32)`

SetNumViewsAboveHardLimit sets NumViewsAboveHardLimit field to given value.

### HasNumViewsAboveHardLimit

`func (o *UserQuotaSummaryForUser) HasNumViewsAboveHardLimit() bool`

HasNumViewsAboveHardLimit returns a boolean if a field has been set.

### SetNumViewsAboveHardLimitNil

`func (o *UserQuotaSummaryForUser) SetNumViewsAboveHardLimitNil(b bool)`

 SetNumViewsAboveHardLimitNil sets the value for NumViewsAboveHardLimit to be an explicit nil

### UnsetNumViewsAboveHardLimit
`func (o *UserQuotaSummaryForUser) UnsetNumViewsAboveHardLimit()`

UnsetNumViewsAboveHardLimit ensures that no value is present for NumViewsAboveHardLimit, not even an explicit nil
### GetTotalNumViews

`func (o *UserQuotaSummaryForUser) GetTotalNumViews() int32`

GetTotalNumViews returns the TotalNumViews field if non-nil, zero value otherwise.

### GetTotalNumViewsOk

`func (o *UserQuotaSummaryForUser) GetTotalNumViewsOk() (*int32, bool)`

GetTotalNumViewsOk returns a tuple with the TotalNumViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNumViews

`func (o *UserQuotaSummaryForUser) SetTotalNumViews(v int32)`

SetTotalNumViews sets TotalNumViews field to given value.

### HasTotalNumViews

`func (o *UserQuotaSummaryForUser) HasTotalNumViews() bool`

HasTotalNumViews returns a boolean if a field has been set.

### SetTotalNumViewsNil

`func (o *UserQuotaSummaryForUser) SetTotalNumViewsNil(b bool)`

 SetTotalNumViewsNil sets the value for TotalNumViews to be an explicit nil

### UnsetTotalNumViews
`func (o *UserQuotaSummaryForUser) UnsetTotalNumViews()`

UnsetTotalNumViews ensures that no value is present for TotalNumViews, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


