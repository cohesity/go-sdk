# FilerLifecycleRuleAgingPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgingCriteria** | **NullableString** | Specifies the criteria for aging | 
**DateInUsecs** | Pointer to **NullableInt64** | Files that possess timestamps exceeding the specified value will be eligible for selection. | [optional] 
**Days** | Pointer to **NullableInt32** | Files that possess timestamps older than the specified value in days will be eligible for selection. | [optional] 

## Methods

### NewFilerLifecycleRuleAgingPolicy

`func NewFilerLifecycleRuleAgingPolicy(agingCriteria NullableString, ) *FilerLifecycleRuleAgingPolicy`

NewFilerLifecycleRuleAgingPolicy instantiates a new FilerLifecycleRuleAgingPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilerLifecycleRuleAgingPolicyWithDefaults

`func NewFilerLifecycleRuleAgingPolicyWithDefaults() *FilerLifecycleRuleAgingPolicy`

NewFilerLifecycleRuleAgingPolicyWithDefaults instantiates a new FilerLifecycleRuleAgingPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgingCriteria

`func (o *FilerLifecycleRuleAgingPolicy) GetAgingCriteria() string`

GetAgingCriteria returns the AgingCriteria field if non-nil, zero value otherwise.

### GetAgingCriteriaOk

`func (o *FilerLifecycleRuleAgingPolicy) GetAgingCriteriaOk() (*string, bool)`

GetAgingCriteriaOk returns a tuple with the AgingCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgingCriteria

`func (o *FilerLifecycleRuleAgingPolicy) SetAgingCriteria(v string)`

SetAgingCriteria sets AgingCriteria field to given value.


### SetAgingCriteriaNil

`func (o *FilerLifecycleRuleAgingPolicy) SetAgingCriteriaNil(b bool)`

 SetAgingCriteriaNil sets the value for AgingCriteria to be an explicit nil

### UnsetAgingCriteria
`func (o *FilerLifecycleRuleAgingPolicy) UnsetAgingCriteria()`

UnsetAgingCriteria ensures that no value is present for AgingCriteria, not even an explicit nil
### GetDateInUsecs

`func (o *FilerLifecycleRuleAgingPolicy) GetDateInUsecs() int64`

GetDateInUsecs returns the DateInUsecs field if non-nil, zero value otherwise.

### GetDateInUsecsOk

`func (o *FilerLifecycleRuleAgingPolicy) GetDateInUsecsOk() (*int64, bool)`

GetDateInUsecsOk returns a tuple with the DateInUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateInUsecs

`func (o *FilerLifecycleRuleAgingPolicy) SetDateInUsecs(v int64)`

SetDateInUsecs sets DateInUsecs field to given value.

### HasDateInUsecs

`func (o *FilerLifecycleRuleAgingPolicy) HasDateInUsecs() bool`

HasDateInUsecs returns a boolean if a field has been set.

### SetDateInUsecsNil

`func (o *FilerLifecycleRuleAgingPolicy) SetDateInUsecsNil(b bool)`

 SetDateInUsecsNil sets the value for DateInUsecs to be an explicit nil

### UnsetDateInUsecs
`func (o *FilerLifecycleRuleAgingPolicy) UnsetDateInUsecs()`

UnsetDateInUsecs ensures that no value is present for DateInUsecs, not even an explicit nil
### GetDays

`func (o *FilerLifecycleRuleAgingPolicy) GetDays() int32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *FilerLifecycleRuleAgingPolicy) GetDaysOk() (*int32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *FilerLifecycleRuleAgingPolicy) SetDays(v int32)`

SetDays sets Days field to given value.

### HasDays

`func (o *FilerLifecycleRuleAgingPolicy) HasDays() bool`

HasDays returns a boolean if a field has been set.

### SetDaysNil

`func (o *FilerLifecycleRuleAgingPolicy) SetDaysNil(b bool)`

 SetDaysNil sets the value for Days to be an explicit nil

### UnsetDays
`func (o *FilerLifecycleRuleAgingPolicy) UnsetDays()`

UnsetDays ensures that no value is present for Days, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


