# FeatureFlag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsApproved** | Pointer to **NullableBool** | Bool to approval status of feature flag. | [optional] 
**IsUiFeature** | Pointer to **NullableBool** | Bool to denote if it&#39;s a UI feature. | [optional] 
**Name** | Pointer to **NullableString** | Name of the feature flag. | [optional] 
**Reason** | Pointer to **NullableString** | Reason for the feature flag override status. | [optional] 
**Timestamp** | Pointer to **NullableInt64** | Timestamp in secs when the override is done. | [optional] 

## Methods

### NewFeatureFlag

`func NewFeatureFlag() *FeatureFlag`

NewFeatureFlag instantiates a new FeatureFlag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureFlagWithDefaults

`func NewFeatureFlagWithDefaults() *FeatureFlag`

NewFeatureFlagWithDefaults instantiates a new FeatureFlag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsApproved

`func (o *FeatureFlag) GetIsApproved() bool`

GetIsApproved returns the IsApproved field if non-nil, zero value otherwise.

### GetIsApprovedOk

`func (o *FeatureFlag) GetIsApprovedOk() (*bool, bool)`

GetIsApprovedOk returns a tuple with the IsApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApproved

`func (o *FeatureFlag) SetIsApproved(v bool)`

SetIsApproved sets IsApproved field to given value.

### HasIsApproved

`func (o *FeatureFlag) HasIsApproved() bool`

HasIsApproved returns a boolean if a field has been set.

### SetIsApprovedNil

`func (o *FeatureFlag) SetIsApprovedNil(b bool)`

 SetIsApprovedNil sets the value for IsApproved to be an explicit nil

### UnsetIsApproved
`func (o *FeatureFlag) UnsetIsApproved()`

UnsetIsApproved ensures that no value is present for IsApproved, not even an explicit nil
### GetIsUiFeature

`func (o *FeatureFlag) GetIsUiFeature() bool`

GetIsUiFeature returns the IsUiFeature field if non-nil, zero value otherwise.

### GetIsUiFeatureOk

`func (o *FeatureFlag) GetIsUiFeatureOk() (*bool, bool)`

GetIsUiFeatureOk returns a tuple with the IsUiFeature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUiFeature

`func (o *FeatureFlag) SetIsUiFeature(v bool)`

SetIsUiFeature sets IsUiFeature field to given value.

### HasIsUiFeature

`func (o *FeatureFlag) HasIsUiFeature() bool`

HasIsUiFeature returns a boolean if a field has been set.

### SetIsUiFeatureNil

`func (o *FeatureFlag) SetIsUiFeatureNil(b bool)`

 SetIsUiFeatureNil sets the value for IsUiFeature to be an explicit nil

### UnsetIsUiFeature
`func (o *FeatureFlag) UnsetIsUiFeature()`

UnsetIsUiFeature ensures that no value is present for IsUiFeature, not even an explicit nil
### GetName

`func (o *FeatureFlag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FeatureFlag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FeatureFlag) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FeatureFlag) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *FeatureFlag) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *FeatureFlag) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetReason

`func (o *FeatureFlag) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *FeatureFlag) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *FeatureFlag) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *FeatureFlag) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *FeatureFlag) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *FeatureFlag) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetTimestamp

`func (o *FeatureFlag) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *FeatureFlag) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *FeatureFlag) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *FeatureFlag) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *FeatureFlag) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *FeatureFlag) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


