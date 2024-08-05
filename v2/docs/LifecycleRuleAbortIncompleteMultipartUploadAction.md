# LifecycleRuleAbortIncompleteMultipartUploadAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Days** | Pointer to **NullableInt64** | Specifies the number of days after which to abort an incomplete multipart upload. | [optional] 

## Methods

### NewLifecycleRuleAbortIncompleteMultipartUploadAction

`func NewLifecycleRuleAbortIncompleteMultipartUploadAction() *LifecycleRuleAbortIncompleteMultipartUploadAction`

NewLifecycleRuleAbortIncompleteMultipartUploadAction instantiates a new LifecycleRuleAbortIncompleteMultipartUploadAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLifecycleRuleAbortIncompleteMultipartUploadActionWithDefaults

`func NewLifecycleRuleAbortIncompleteMultipartUploadActionWithDefaults() *LifecycleRuleAbortIncompleteMultipartUploadAction`

NewLifecycleRuleAbortIncompleteMultipartUploadActionWithDefaults instantiates a new LifecycleRuleAbortIncompleteMultipartUploadAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDays

`func (o *LifecycleRuleAbortIncompleteMultipartUploadAction) GetDays() int64`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *LifecycleRuleAbortIncompleteMultipartUploadAction) GetDaysOk() (*int64, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *LifecycleRuleAbortIncompleteMultipartUploadAction) SetDays(v int64)`

SetDays sets Days field to given value.

### HasDays

`func (o *LifecycleRuleAbortIncompleteMultipartUploadAction) HasDays() bool`

HasDays returns a boolean if a field has been set.

### SetDaysNil

`func (o *LifecycleRuleAbortIncompleteMultipartUploadAction) SetDaysNil(b bool)`

 SetDaysNil sets the value for Days to be an explicit nil

### UnsetDays
`func (o *LifecycleRuleAbortIncompleteMultipartUploadAction) UnsetDays()`

UnsetDays ensures that no value is present for Days, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


