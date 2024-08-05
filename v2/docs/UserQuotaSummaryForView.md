# UserQuotaSummaryForView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultQuotaPolicy** | Pointer to [**QuotaPolicy**](QuotaPolicy.md) |  | [optional] 
**NumUsersAboveAlertThreshold** | Pointer to **NullableInt64** | Number of users who has exceeded their specified alert limit. | [optional] 
**NumUsersAboveHardLimit** | Pointer to **NullableInt64** | Number of users who has exceeded their specified quota hard limit. | [optional] 
**TotalNumUsers** | Pointer to **NullableInt64** | Total number of users who has either a user quota policy override specified or has non-zero logical usage. | [optional] 

## Methods

### NewUserQuotaSummaryForView

`func NewUserQuotaSummaryForView() *UserQuotaSummaryForView`

NewUserQuotaSummaryForView instantiates a new UserQuotaSummaryForView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserQuotaSummaryForViewWithDefaults

`func NewUserQuotaSummaryForViewWithDefaults() *UserQuotaSummaryForView`

NewUserQuotaSummaryForViewWithDefaults instantiates a new UserQuotaSummaryForView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultQuotaPolicy

`func (o *UserQuotaSummaryForView) GetDefaultQuotaPolicy() QuotaPolicy`

GetDefaultQuotaPolicy returns the DefaultQuotaPolicy field if non-nil, zero value otherwise.

### GetDefaultQuotaPolicyOk

`func (o *UserQuotaSummaryForView) GetDefaultQuotaPolicyOk() (*QuotaPolicy, bool)`

GetDefaultQuotaPolicyOk returns a tuple with the DefaultQuotaPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultQuotaPolicy

`func (o *UserQuotaSummaryForView) SetDefaultQuotaPolicy(v QuotaPolicy)`

SetDefaultQuotaPolicy sets DefaultQuotaPolicy field to given value.

### HasDefaultQuotaPolicy

`func (o *UserQuotaSummaryForView) HasDefaultQuotaPolicy() bool`

HasDefaultQuotaPolicy returns a boolean if a field has been set.

### GetNumUsersAboveAlertThreshold

`func (o *UserQuotaSummaryForView) GetNumUsersAboveAlertThreshold() int64`

GetNumUsersAboveAlertThreshold returns the NumUsersAboveAlertThreshold field if non-nil, zero value otherwise.

### GetNumUsersAboveAlertThresholdOk

`func (o *UserQuotaSummaryForView) GetNumUsersAboveAlertThresholdOk() (*int64, bool)`

GetNumUsersAboveAlertThresholdOk returns a tuple with the NumUsersAboveAlertThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUsersAboveAlertThreshold

`func (o *UserQuotaSummaryForView) SetNumUsersAboveAlertThreshold(v int64)`

SetNumUsersAboveAlertThreshold sets NumUsersAboveAlertThreshold field to given value.

### HasNumUsersAboveAlertThreshold

`func (o *UserQuotaSummaryForView) HasNumUsersAboveAlertThreshold() bool`

HasNumUsersAboveAlertThreshold returns a boolean if a field has been set.

### SetNumUsersAboveAlertThresholdNil

`func (o *UserQuotaSummaryForView) SetNumUsersAboveAlertThresholdNil(b bool)`

 SetNumUsersAboveAlertThresholdNil sets the value for NumUsersAboveAlertThreshold to be an explicit nil

### UnsetNumUsersAboveAlertThreshold
`func (o *UserQuotaSummaryForView) UnsetNumUsersAboveAlertThreshold()`

UnsetNumUsersAboveAlertThreshold ensures that no value is present for NumUsersAboveAlertThreshold, not even an explicit nil
### GetNumUsersAboveHardLimit

`func (o *UserQuotaSummaryForView) GetNumUsersAboveHardLimit() int64`

GetNumUsersAboveHardLimit returns the NumUsersAboveHardLimit field if non-nil, zero value otherwise.

### GetNumUsersAboveHardLimitOk

`func (o *UserQuotaSummaryForView) GetNumUsersAboveHardLimitOk() (*int64, bool)`

GetNumUsersAboveHardLimitOk returns a tuple with the NumUsersAboveHardLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUsersAboveHardLimit

`func (o *UserQuotaSummaryForView) SetNumUsersAboveHardLimit(v int64)`

SetNumUsersAboveHardLimit sets NumUsersAboveHardLimit field to given value.

### HasNumUsersAboveHardLimit

`func (o *UserQuotaSummaryForView) HasNumUsersAboveHardLimit() bool`

HasNumUsersAboveHardLimit returns a boolean if a field has been set.

### SetNumUsersAboveHardLimitNil

`func (o *UserQuotaSummaryForView) SetNumUsersAboveHardLimitNil(b bool)`

 SetNumUsersAboveHardLimitNil sets the value for NumUsersAboveHardLimit to be an explicit nil

### UnsetNumUsersAboveHardLimit
`func (o *UserQuotaSummaryForView) UnsetNumUsersAboveHardLimit()`

UnsetNumUsersAboveHardLimit ensures that no value is present for NumUsersAboveHardLimit, not even an explicit nil
### GetTotalNumUsers

`func (o *UserQuotaSummaryForView) GetTotalNumUsers() int64`

GetTotalNumUsers returns the TotalNumUsers field if non-nil, zero value otherwise.

### GetTotalNumUsersOk

`func (o *UserQuotaSummaryForView) GetTotalNumUsersOk() (*int64, bool)`

GetTotalNumUsersOk returns a tuple with the TotalNumUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNumUsers

`func (o *UserQuotaSummaryForView) SetTotalNumUsers(v int64)`

SetTotalNumUsers sets TotalNumUsers field to given value.

### HasTotalNumUsers

`func (o *UserQuotaSummaryForView) HasTotalNumUsers() bool`

HasTotalNumUsers returns a boolean if a field has been set.

### SetTotalNumUsersNil

`func (o *UserQuotaSummaryForView) SetTotalNumUsersNil(b bool)`

 SetTotalNumUsersNil sets the value for TotalNumUsers to be an explicit nil

### UnsetTotalNumUsers
`func (o *UserQuotaSummaryForView) UnsetTotalNumUsers()`

UnsetTotalNumUsers ensures that no value is present for TotalNumUsers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


