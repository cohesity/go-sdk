# UpdateFeatureFlagParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Name of the feature flag that is to be updated. | [optional] 
**IsUiFeature** | Pointer to **NullableBool** | Bool to specify if it&#39;s UI feature flag. | [optional] 
**IsApproved** | Pointer to **NullableBool** | New override status for the feature flag. | [optional] 
**Reason** | Pointer to **NullableString** | Reason for updating the feature flag override status. | [optional] 
**Clear** | Pointer to **NullableBool** | Clear any override status for the feature flag. | [optional] 
**Timestamp** | Pointer to **NullableInt64** | Specifies the timestamp of override operation. | [optional] 

## Methods

### NewUpdateFeatureFlagParams

`func NewUpdateFeatureFlagParams() *UpdateFeatureFlagParams`

NewUpdateFeatureFlagParams instantiates a new UpdateFeatureFlagParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFeatureFlagParamsWithDefaults

`func NewUpdateFeatureFlagParamsWithDefaults() *UpdateFeatureFlagParams`

NewUpdateFeatureFlagParamsWithDefaults instantiates a new UpdateFeatureFlagParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateFeatureFlagParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateFeatureFlagParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateFeatureFlagParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateFeatureFlagParams) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateFeatureFlagParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateFeatureFlagParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsUiFeature

`func (o *UpdateFeatureFlagParams) GetIsUiFeature() bool`

GetIsUiFeature returns the IsUiFeature field if non-nil, zero value otherwise.

### GetIsUiFeatureOk

`func (o *UpdateFeatureFlagParams) GetIsUiFeatureOk() (*bool, bool)`

GetIsUiFeatureOk returns a tuple with the IsUiFeature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUiFeature

`func (o *UpdateFeatureFlagParams) SetIsUiFeature(v bool)`

SetIsUiFeature sets IsUiFeature field to given value.

### HasIsUiFeature

`func (o *UpdateFeatureFlagParams) HasIsUiFeature() bool`

HasIsUiFeature returns a boolean if a field has been set.

### SetIsUiFeatureNil

`func (o *UpdateFeatureFlagParams) SetIsUiFeatureNil(b bool)`

 SetIsUiFeatureNil sets the value for IsUiFeature to be an explicit nil

### UnsetIsUiFeature
`func (o *UpdateFeatureFlagParams) UnsetIsUiFeature()`

UnsetIsUiFeature ensures that no value is present for IsUiFeature, not even an explicit nil
### GetIsApproved

`func (o *UpdateFeatureFlagParams) GetIsApproved() bool`

GetIsApproved returns the IsApproved field if non-nil, zero value otherwise.

### GetIsApprovedOk

`func (o *UpdateFeatureFlagParams) GetIsApprovedOk() (*bool, bool)`

GetIsApprovedOk returns a tuple with the IsApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApproved

`func (o *UpdateFeatureFlagParams) SetIsApproved(v bool)`

SetIsApproved sets IsApproved field to given value.

### HasIsApproved

`func (o *UpdateFeatureFlagParams) HasIsApproved() bool`

HasIsApproved returns a boolean if a field has been set.

### SetIsApprovedNil

`func (o *UpdateFeatureFlagParams) SetIsApprovedNil(b bool)`

 SetIsApprovedNil sets the value for IsApproved to be an explicit nil

### UnsetIsApproved
`func (o *UpdateFeatureFlagParams) UnsetIsApproved()`

UnsetIsApproved ensures that no value is present for IsApproved, not even an explicit nil
### GetReason

`func (o *UpdateFeatureFlagParams) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *UpdateFeatureFlagParams) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *UpdateFeatureFlagParams) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *UpdateFeatureFlagParams) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *UpdateFeatureFlagParams) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *UpdateFeatureFlagParams) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetClear

`func (o *UpdateFeatureFlagParams) GetClear() bool`

GetClear returns the Clear field if non-nil, zero value otherwise.

### GetClearOk

`func (o *UpdateFeatureFlagParams) GetClearOk() (*bool, bool)`

GetClearOk returns a tuple with the Clear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClear

`func (o *UpdateFeatureFlagParams) SetClear(v bool)`

SetClear sets Clear field to given value.

### HasClear

`func (o *UpdateFeatureFlagParams) HasClear() bool`

HasClear returns a boolean if a field has been set.

### SetClearNil

`func (o *UpdateFeatureFlagParams) SetClearNil(b bool)`

 SetClearNil sets the value for Clear to be an explicit nil

### UnsetClear
`func (o *UpdateFeatureFlagParams) UnsetClear()`

UnsetClear ensures that no value is present for Clear, not even an explicit nil
### GetTimestamp

`func (o *UpdateFeatureFlagParams) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *UpdateFeatureFlagParams) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *UpdateFeatureFlagParams) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *UpdateFeatureFlagParams) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *UpdateFeatureFlagParams) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *UpdateFeatureFlagParams) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


