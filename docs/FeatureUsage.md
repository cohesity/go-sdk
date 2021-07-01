# FeatureUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentUsageGiB** | Pointer to **NullableInt64** | Feature usage by the cluster. | [optional] 
**FeatureName** | Pointer to **NullableString** | Name of feature. | [optional] 
**NumVm** | Pointer to **NullableInt64** | Number of VM spinned. | [optional] 

## Methods

### NewFeatureUsage

`func NewFeatureUsage() *FeatureUsage`

NewFeatureUsage instantiates a new FeatureUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureUsageWithDefaults

`func NewFeatureUsageWithDefaults() *FeatureUsage`

NewFeatureUsageWithDefaults instantiates a new FeatureUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentUsageGiB

`func (o *FeatureUsage) GetCurrentUsageGiB() int64`

GetCurrentUsageGiB returns the CurrentUsageGiB field if non-nil, zero value otherwise.

### GetCurrentUsageGiBOk

`func (o *FeatureUsage) GetCurrentUsageGiBOk() (*int64, bool)`

GetCurrentUsageGiBOk returns a tuple with the CurrentUsageGiB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUsageGiB

`func (o *FeatureUsage) SetCurrentUsageGiB(v int64)`

SetCurrentUsageGiB sets CurrentUsageGiB field to given value.

### HasCurrentUsageGiB

`func (o *FeatureUsage) HasCurrentUsageGiB() bool`

HasCurrentUsageGiB returns a boolean if a field has been set.

### SetCurrentUsageGiBNil

`func (o *FeatureUsage) SetCurrentUsageGiBNil(b bool)`

 SetCurrentUsageGiBNil sets the value for CurrentUsageGiB to be an explicit nil

### UnsetCurrentUsageGiB
`func (o *FeatureUsage) UnsetCurrentUsageGiB()`

UnsetCurrentUsageGiB ensures that no value is present for CurrentUsageGiB, not even an explicit nil
### GetFeatureName

`func (o *FeatureUsage) GetFeatureName() string`

GetFeatureName returns the FeatureName field if non-nil, zero value otherwise.

### GetFeatureNameOk

`func (o *FeatureUsage) GetFeatureNameOk() (*string, bool)`

GetFeatureNameOk returns a tuple with the FeatureName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureName

`func (o *FeatureUsage) SetFeatureName(v string)`

SetFeatureName sets FeatureName field to given value.

### HasFeatureName

`func (o *FeatureUsage) HasFeatureName() bool`

HasFeatureName returns a boolean if a field has been set.

### SetFeatureNameNil

`func (o *FeatureUsage) SetFeatureNameNil(b bool)`

 SetFeatureNameNil sets the value for FeatureName to be an explicit nil

### UnsetFeatureName
`func (o *FeatureUsage) UnsetFeatureName()`

UnsetFeatureName ensures that no value is present for FeatureName, not even an explicit nil
### GetNumVm

`func (o *FeatureUsage) GetNumVm() int64`

GetNumVm returns the NumVm field if non-nil, zero value otherwise.

### GetNumVmOk

`func (o *FeatureUsage) GetNumVmOk() (*int64, bool)`

GetNumVmOk returns a tuple with the NumVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumVm

`func (o *FeatureUsage) SetNumVm(v int64)`

SetNumVm sets NumVm field to given value.

### HasNumVm

`func (o *FeatureUsage) HasNumVm() bool`

HasNumVm returns a boolean if a field has been set.

### SetNumVmNil

`func (o *FeatureUsage) SetNumVmNil(b bool)`

 SetNumVmNil sets the value for NumVm to be an explicit nil

### UnsetNumVm
`func (o *FeatureUsage) UnsetNumVm()`

UnsetNumVm ensures that no value is present for NumVm, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


