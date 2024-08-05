# PatchCompatibilityStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Specifies patch compatibility check is required for apply or revert patch. | [optional] 
**CurrentVersion** | Pointer to **NullableString** | Specifies the current version. | [optional] 
**IsCompatible** | Pointer to **NullableBool** | Specifies whether the patch is compatible for the specified action. | [optional] 
**TargetVersion** | Pointer to **NullableString** | Specifies the target version. | [optional] 

## Methods

### NewPatchCompatibilityStatus

`func NewPatchCompatibilityStatus() *PatchCompatibilityStatus`

NewPatchCompatibilityStatus instantiates a new PatchCompatibilityStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchCompatibilityStatusWithDefaults

`func NewPatchCompatibilityStatusWithDefaults() *PatchCompatibilityStatus`

NewPatchCompatibilityStatusWithDefaults instantiates a new PatchCompatibilityStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *PatchCompatibilityStatus) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PatchCompatibilityStatus) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PatchCompatibilityStatus) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *PatchCompatibilityStatus) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCurrentVersion

`func (o *PatchCompatibilityStatus) GetCurrentVersion() string`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *PatchCompatibilityStatus) GetCurrentVersionOk() (*string, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *PatchCompatibilityStatus) SetCurrentVersion(v string)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *PatchCompatibilityStatus) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### SetCurrentVersionNil

`func (o *PatchCompatibilityStatus) SetCurrentVersionNil(b bool)`

 SetCurrentVersionNil sets the value for CurrentVersion to be an explicit nil

### UnsetCurrentVersion
`func (o *PatchCompatibilityStatus) UnsetCurrentVersion()`

UnsetCurrentVersion ensures that no value is present for CurrentVersion, not even an explicit nil
### GetIsCompatible

`func (o *PatchCompatibilityStatus) GetIsCompatible() bool`

GetIsCompatible returns the IsCompatible field if non-nil, zero value otherwise.

### GetIsCompatibleOk

`func (o *PatchCompatibilityStatus) GetIsCompatibleOk() (*bool, bool)`

GetIsCompatibleOk returns a tuple with the IsCompatible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCompatible

`func (o *PatchCompatibilityStatus) SetIsCompatible(v bool)`

SetIsCompatible sets IsCompatible field to given value.

### HasIsCompatible

`func (o *PatchCompatibilityStatus) HasIsCompatible() bool`

HasIsCompatible returns a boolean if a field has been set.

### SetIsCompatibleNil

`func (o *PatchCompatibilityStatus) SetIsCompatibleNil(b bool)`

 SetIsCompatibleNil sets the value for IsCompatible to be an explicit nil

### UnsetIsCompatible
`func (o *PatchCompatibilityStatus) UnsetIsCompatible()`

UnsetIsCompatible ensures that no value is present for IsCompatible, not even an explicit nil
### GetTargetVersion

`func (o *PatchCompatibilityStatus) GetTargetVersion() string`

GetTargetVersion returns the TargetVersion field if non-nil, zero value otherwise.

### GetTargetVersionOk

`func (o *PatchCompatibilityStatus) GetTargetVersionOk() (*string, bool)`

GetTargetVersionOk returns a tuple with the TargetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVersion

`func (o *PatchCompatibilityStatus) SetTargetVersion(v string)`

SetTargetVersion sets TargetVersion field to given value.

### HasTargetVersion

`func (o *PatchCompatibilityStatus) HasTargetVersion() bool`

HasTargetVersion returns a boolean if a field has been set.

### SetTargetVersionNil

`func (o *PatchCompatibilityStatus) SetTargetVersionNil(b bool)`

 SetTargetVersionNil sets the value for TargetVersion to be an explicit nil

### UnsetTargetVersion
`func (o *PatchCompatibilityStatus) UnsetTargetVersion()`

UnsetTargetVersion ensures that no value is present for TargetVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


