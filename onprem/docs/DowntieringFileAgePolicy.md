# DowntieringFileAgePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Condition** | Pointer to **NullableString** | Specifies the condition for the file age. | [optional] 
**AgeMsecs** | Pointer to **NullableInt64** | Specifies the number of msecs used for file selection. | [optional] 

## Methods

### NewDowntieringFileAgePolicy

`func NewDowntieringFileAgePolicy() *DowntieringFileAgePolicy`

NewDowntieringFileAgePolicy instantiates a new DowntieringFileAgePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDowntieringFileAgePolicyWithDefaults

`func NewDowntieringFileAgePolicyWithDefaults() *DowntieringFileAgePolicy`

NewDowntieringFileAgePolicyWithDefaults instantiates a new DowntieringFileAgePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCondition

`func (o *DowntieringFileAgePolicy) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *DowntieringFileAgePolicy) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *DowntieringFileAgePolicy) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *DowntieringFileAgePolicy) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *DowntieringFileAgePolicy) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *DowntieringFileAgePolicy) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil
### GetAgeMsecs

`func (o *DowntieringFileAgePolicy) GetAgeMsecs() int64`

GetAgeMsecs returns the AgeMsecs field if non-nil, zero value otherwise.

### GetAgeMsecsOk

`func (o *DowntieringFileAgePolicy) GetAgeMsecsOk() (*int64, bool)`

GetAgeMsecsOk returns a tuple with the AgeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeMsecs

`func (o *DowntieringFileAgePolicy) SetAgeMsecs(v int64)`

SetAgeMsecs sets AgeMsecs field to given value.

### HasAgeMsecs

`func (o *DowntieringFileAgePolicy) HasAgeMsecs() bool`

HasAgeMsecs returns a boolean if a field has been set.

### SetAgeMsecsNil

`func (o *DowntieringFileAgePolicy) SetAgeMsecsNil(b bool)`

 SetAgeMsecsNil sets the value for AgeMsecs to be an explicit nil

### UnsetAgeMsecs
`func (o *DowntieringFileAgePolicy) UnsetAgeMsecs()`

UnsetAgeMsecs ensures that no value is present for AgeMsecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


