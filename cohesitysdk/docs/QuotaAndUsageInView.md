# QuotaAndUsageInView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quota** | Pointer to [**QuotaPolicy**](QuotaPolicy.md) |  | [optional] 
**UsageBytes** | Pointer to **NullableInt64** | Usage in bytes of this user in this view. | [optional] 
**ViewId** | Pointer to **NullableInt64** | The usage and quota policy information of this user for this view. | [optional] 
**ViewName** | Pointer to **NullableString** | View name. | [optional] 

## Methods

### NewQuotaAndUsageInView

`func NewQuotaAndUsageInView() *QuotaAndUsageInView`

NewQuotaAndUsageInView instantiates a new QuotaAndUsageInView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotaAndUsageInViewWithDefaults

`func NewQuotaAndUsageInViewWithDefaults() *QuotaAndUsageInView`

NewQuotaAndUsageInViewWithDefaults instantiates a new QuotaAndUsageInView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuota

`func (o *QuotaAndUsageInView) GetQuota() QuotaPolicy`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *QuotaAndUsageInView) GetQuotaOk() (*QuotaPolicy, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *QuotaAndUsageInView) SetQuota(v QuotaPolicy)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *QuotaAndUsageInView) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### GetUsageBytes

`func (o *QuotaAndUsageInView) GetUsageBytes() int64`

GetUsageBytes returns the UsageBytes field if non-nil, zero value otherwise.

### GetUsageBytesOk

`func (o *QuotaAndUsageInView) GetUsageBytesOk() (*int64, bool)`

GetUsageBytesOk returns a tuple with the UsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageBytes

`func (o *QuotaAndUsageInView) SetUsageBytes(v int64)`

SetUsageBytes sets UsageBytes field to given value.

### HasUsageBytes

`func (o *QuotaAndUsageInView) HasUsageBytes() bool`

HasUsageBytes returns a boolean if a field has been set.

### SetUsageBytesNil

`func (o *QuotaAndUsageInView) SetUsageBytesNil(b bool)`

 SetUsageBytesNil sets the value for UsageBytes to be an explicit nil

### UnsetUsageBytes
`func (o *QuotaAndUsageInView) UnsetUsageBytes()`

UnsetUsageBytes ensures that no value is present for UsageBytes, not even an explicit nil
### GetViewId

`func (o *QuotaAndUsageInView) GetViewId() int64`

GetViewId returns the ViewId field if non-nil, zero value otherwise.

### GetViewIdOk

`func (o *QuotaAndUsageInView) GetViewIdOk() (*int64, bool)`

GetViewIdOk returns a tuple with the ViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewId

`func (o *QuotaAndUsageInView) SetViewId(v int64)`

SetViewId sets ViewId field to given value.

### HasViewId

`func (o *QuotaAndUsageInView) HasViewId() bool`

HasViewId returns a boolean if a field has been set.

### SetViewIdNil

`func (o *QuotaAndUsageInView) SetViewIdNil(b bool)`

 SetViewIdNil sets the value for ViewId to be an explicit nil

### UnsetViewId
`func (o *QuotaAndUsageInView) UnsetViewId()`

UnsetViewId ensures that no value is present for ViewId, not even an explicit nil
### GetViewName

`func (o *QuotaAndUsageInView) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *QuotaAndUsageInView) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *QuotaAndUsageInView) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *QuotaAndUsageInView) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### SetViewNameNil

`func (o *QuotaAndUsageInView) SetViewNameNil(b bool)`

 SetViewNameNil sets the value for ViewName to be an explicit nil

### UnsetViewName
`func (o *QuotaAndUsageInView) UnsetViewName()`

UnsetViewName ensures that no value is present for ViewName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


