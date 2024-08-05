# PatchCompatibilityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Specifies patch compatibility check is required for apply or revert patch. | 
**CurrentVersion** | **NullableString** | Specifies the current version. | 
**TargetVersion** | **NullableString** | Specifies the target version. | 

## Methods

### NewPatchCompatibilityRequest

`func NewPatchCompatibilityRequest(action string, currentVersion NullableString, targetVersion NullableString, ) *PatchCompatibilityRequest`

NewPatchCompatibilityRequest instantiates a new PatchCompatibilityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchCompatibilityRequestWithDefaults

`func NewPatchCompatibilityRequestWithDefaults() *PatchCompatibilityRequest`

NewPatchCompatibilityRequestWithDefaults instantiates a new PatchCompatibilityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *PatchCompatibilityRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PatchCompatibilityRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PatchCompatibilityRequest) SetAction(v string)`

SetAction sets Action field to given value.


### GetCurrentVersion

`func (o *PatchCompatibilityRequest) GetCurrentVersion() string`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *PatchCompatibilityRequest) GetCurrentVersionOk() (*string, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *PatchCompatibilityRequest) SetCurrentVersion(v string)`

SetCurrentVersion sets CurrentVersion field to given value.


### SetCurrentVersionNil

`func (o *PatchCompatibilityRequest) SetCurrentVersionNil(b bool)`

 SetCurrentVersionNil sets the value for CurrentVersion to be an explicit nil

### UnsetCurrentVersion
`func (o *PatchCompatibilityRequest) UnsetCurrentVersion()`

UnsetCurrentVersion ensures that no value is present for CurrentVersion, not even an explicit nil
### GetTargetVersion

`func (o *PatchCompatibilityRequest) GetTargetVersion() string`

GetTargetVersion returns the TargetVersion field if non-nil, zero value otherwise.

### GetTargetVersionOk

`func (o *PatchCompatibilityRequest) GetTargetVersionOk() (*string, bool)`

GetTargetVersionOk returns a tuple with the TargetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVersion

`func (o *PatchCompatibilityRequest) SetTargetVersion(v string)`

SetTargetVersion sets TargetVersion field to given value.


### SetTargetVersionNil

`func (o *PatchCompatibilityRequest) SetTargetVersionNil(b bool)`

 SetTargetVersionNil sets the value for TargetVersion to be an explicit nil

### UnsetTargetVersion
`func (o *PatchCompatibilityRequest) UnsetTargetVersion()`

UnsetTargetVersion ensures that no value is present for TargetVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


