# MonthlyUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeatureName** | Pointer to **NullableString** |  | [optional] 
**MonthlyAvgUsage** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewMonthlyUsage

`func NewMonthlyUsage() *MonthlyUsage`

NewMonthlyUsage instantiates a new MonthlyUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonthlyUsageWithDefaults

`func NewMonthlyUsageWithDefaults() *MonthlyUsage`

NewMonthlyUsageWithDefaults instantiates a new MonthlyUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatureName

`func (o *MonthlyUsage) GetFeatureName() string`

GetFeatureName returns the FeatureName field if non-nil, zero value otherwise.

### GetFeatureNameOk

`func (o *MonthlyUsage) GetFeatureNameOk() (*string, bool)`

GetFeatureNameOk returns a tuple with the FeatureName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureName

`func (o *MonthlyUsage) SetFeatureName(v string)`

SetFeatureName sets FeatureName field to given value.

### HasFeatureName

`func (o *MonthlyUsage) HasFeatureName() bool`

HasFeatureName returns a boolean if a field has been set.

### SetFeatureNameNil

`func (o *MonthlyUsage) SetFeatureNameNil(b bool)`

 SetFeatureNameNil sets the value for FeatureName to be an explicit nil

### UnsetFeatureName
`func (o *MonthlyUsage) UnsetFeatureName()`

UnsetFeatureName ensures that no value is present for FeatureName, not even an explicit nil
### GetMonthlyAvgUsage

`func (o *MonthlyUsage) GetMonthlyAvgUsage() []int64`

GetMonthlyAvgUsage returns the MonthlyAvgUsage field if non-nil, zero value otherwise.

### GetMonthlyAvgUsageOk

`func (o *MonthlyUsage) GetMonthlyAvgUsageOk() (*[]int64, bool)`

GetMonthlyAvgUsageOk returns a tuple with the MonthlyAvgUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyAvgUsage

`func (o *MonthlyUsage) SetMonthlyAvgUsage(v []int64)`

SetMonthlyAvgUsage sets MonthlyAvgUsage field to given value.

### HasMonthlyAvgUsage

`func (o *MonthlyUsage) HasMonthlyAvgUsage() bool`

HasMonthlyAvgUsage returns a boolean if a field has been set.

### SetMonthlyAvgUsageNil

`func (o *MonthlyUsage) SetMonthlyAvgUsageNil(b bool)`

 SetMonthlyAvgUsageNil sets the value for MonthlyAvgUsage to be an explicit nil

### UnsetMonthlyAvgUsage
`func (o *MonthlyUsage) UnsetMonthlyAvgUsage()`

UnsetMonthlyAvgUsage ensures that no value is present for MonthlyAvgUsage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


