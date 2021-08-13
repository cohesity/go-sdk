# UptieringFileAgePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Condition** | Pointer to **NullableString** | Specifies the condition for the file age. | [optional] 
**AgeMsecs** | Pointer to **NullableInt64** | Specifies the number of msecs used for file selection. | [optional] 
**NumFileAccess** | Pointer to **NullableInt32** | Specifies number of file access in last ageMsecs. | [optional] 

## Methods

### NewUptieringFileAgePolicy

`func NewUptieringFileAgePolicy() *UptieringFileAgePolicy`

NewUptieringFileAgePolicy instantiates a new UptieringFileAgePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUptieringFileAgePolicyWithDefaults

`func NewUptieringFileAgePolicyWithDefaults() *UptieringFileAgePolicy`

NewUptieringFileAgePolicyWithDefaults instantiates a new UptieringFileAgePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCondition

`func (o *UptieringFileAgePolicy) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *UptieringFileAgePolicy) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *UptieringFileAgePolicy) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *UptieringFileAgePolicy) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *UptieringFileAgePolicy) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *UptieringFileAgePolicy) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil
### GetAgeMsecs

`func (o *UptieringFileAgePolicy) GetAgeMsecs() int64`

GetAgeMsecs returns the AgeMsecs field if non-nil, zero value otherwise.

### GetAgeMsecsOk

`func (o *UptieringFileAgePolicy) GetAgeMsecsOk() (*int64, bool)`

GetAgeMsecsOk returns a tuple with the AgeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeMsecs

`func (o *UptieringFileAgePolicy) SetAgeMsecs(v int64)`

SetAgeMsecs sets AgeMsecs field to given value.

### HasAgeMsecs

`func (o *UptieringFileAgePolicy) HasAgeMsecs() bool`

HasAgeMsecs returns a boolean if a field has been set.

### SetAgeMsecsNil

`func (o *UptieringFileAgePolicy) SetAgeMsecsNil(b bool)`

 SetAgeMsecsNil sets the value for AgeMsecs to be an explicit nil

### UnsetAgeMsecs
`func (o *UptieringFileAgePolicy) UnsetAgeMsecs()`

UnsetAgeMsecs ensures that no value is present for AgeMsecs, not even an explicit nil
### GetNumFileAccess

`func (o *UptieringFileAgePolicy) GetNumFileAccess() int32`

GetNumFileAccess returns the NumFileAccess field if non-nil, zero value otherwise.

### GetNumFileAccessOk

`func (o *UptieringFileAgePolicy) GetNumFileAccessOk() (*int32, bool)`

GetNumFileAccessOk returns a tuple with the NumFileAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFileAccess

`func (o *UptieringFileAgePolicy) SetNumFileAccess(v int32)`

SetNumFileAccess sets NumFileAccess field to given value.

### HasNumFileAccess

`func (o *UptieringFileAgePolicy) HasNumFileAccess() bool`

HasNumFileAccess returns a boolean if a field has been set.

### SetNumFileAccessNil

`func (o *UptieringFileAgePolicy) SetNumFileAccessNil(b bool)`

 SetNumFileAccessNil sets the value for NumFileAccess to be an explicit nil

### UnsetNumFileAccess
`func (o *UptieringFileAgePolicy) UnsetNumFileAccess()`

UnsetNumFileAccess ensures that no value is present for NumFileAccess, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


