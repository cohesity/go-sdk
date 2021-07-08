# Overusage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeatureName** | Pointer to **NullableString** | Name of feature. | [optional] 
**OverusedGiB** | Pointer to **NullableInt64** | Feature overusage by the cluster. | [optional] 
**OverusedVm** | Pointer to **NullableInt64** | Number of overused VM spinned. | [optional] 

## Methods

### NewOverusage

`func NewOverusage() *Overusage`

NewOverusage instantiates a new Overusage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverusageWithDefaults

`func NewOverusageWithDefaults() *Overusage`

NewOverusageWithDefaults instantiates a new Overusage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatureName

`func (o *Overusage) GetFeatureName() string`

GetFeatureName returns the FeatureName field if non-nil, zero value otherwise.

### GetFeatureNameOk

`func (o *Overusage) GetFeatureNameOk() (*string, bool)`

GetFeatureNameOk returns a tuple with the FeatureName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureName

`func (o *Overusage) SetFeatureName(v string)`

SetFeatureName sets FeatureName field to given value.

### HasFeatureName

`func (o *Overusage) HasFeatureName() bool`

HasFeatureName returns a boolean if a field has been set.

### SetFeatureNameNil

`func (o *Overusage) SetFeatureNameNil(b bool)`

 SetFeatureNameNil sets the value for FeatureName to be an explicit nil

### UnsetFeatureName
`func (o *Overusage) UnsetFeatureName()`

UnsetFeatureName ensures that no value is present for FeatureName, not even an explicit nil
### GetOverusedGiB

`func (o *Overusage) GetOverusedGiB() int64`

GetOverusedGiB returns the OverusedGiB field if non-nil, zero value otherwise.

### GetOverusedGiBOk

`func (o *Overusage) GetOverusedGiBOk() (*int64, bool)`

GetOverusedGiBOk returns a tuple with the OverusedGiB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverusedGiB

`func (o *Overusage) SetOverusedGiB(v int64)`

SetOverusedGiB sets OverusedGiB field to given value.

### HasOverusedGiB

`func (o *Overusage) HasOverusedGiB() bool`

HasOverusedGiB returns a boolean if a field has been set.

### SetOverusedGiBNil

`func (o *Overusage) SetOverusedGiBNil(b bool)`

 SetOverusedGiBNil sets the value for OverusedGiB to be an explicit nil

### UnsetOverusedGiB
`func (o *Overusage) UnsetOverusedGiB()`

UnsetOverusedGiB ensures that no value is present for OverusedGiB, not even an explicit nil
### GetOverusedVm

`func (o *Overusage) GetOverusedVm() int64`

GetOverusedVm returns the OverusedVm field if non-nil, zero value otherwise.

### GetOverusedVmOk

`func (o *Overusage) GetOverusedVmOk() (*int64, bool)`

GetOverusedVmOk returns a tuple with the OverusedVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverusedVm

`func (o *Overusage) SetOverusedVm(v int64)`

SetOverusedVm sets OverusedVm field to given value.

### HasOverusedVm

`func (o *Overusage) HasOverusedVm() bool`

HasOverusedVm returns a boolean if a field has been set.

### SetOverusedVmNil

`func (o *Overusage) SetOverusedVmNil(b bool)`

 SetOverusedVmNil sets the value for OverusedVm to be an explicit nil

### UnsetOverusedVm
`func (o *Overusage) UnsetOverusedVm()`

UnsetOverusedVm ensures that no value is present for OverusedVm, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


