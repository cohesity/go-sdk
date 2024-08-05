# PrimaryArchivalTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetId** | **NullableInt64** | Specifies the Archival target id to take primary backup. | 
**TargetName** | Pointer to **NullableString** | Specifies the Archival target name where Snapshots are copied. | [optional] [readonly] 
**TierSettings** | Pointer to [**TierLevelSettings**](TierLevelSettings.md) |  | [optional] 

## Methods

### NewPrimaryArchivalTarget

`func NewPrimaryArchivalTarget(targetId NullableInt64, ) *PrimaryArchivalTarget`

NewPrimaryArchivalTarget instantiates a new PrimaryArchivalTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrimaryArchivalTargetWithDefaults

`func NewPrimaryArchivalTargetWithDefaults() *PrimaryArchivalTarget`

NewPrimaryArchivalTargetWithDefaults instantiates a new PrimaryArchivalTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetId

`func (o *PrimaryArchivalTarget) GetTargetId() int64`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *PrimaryArchivalTarget) GetTargetIdOk() (*int64, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *PrimaryArchivalTarget) SetTargetId(v int64)`

SetTargetId sets TargetId field to given value.


### SetTargetIdNil

`func (o *PrimaryArchivalTarget) SetTargetIdNil(b bool)`

 SetTargetIdNil sets the value for TargetId to be an explicit nil

### UnsetTargetId
`func (o *PrimaryArchivalTarget) UnsetTargetId()`

UnsetTargetId ensures that no value is present for TargetId, not even an explicit nil
### GetTargetName

`func (o *PrimaryArchivalTarget) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *PrimaryArchivalTarget) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *PrimaryArchivalTarget) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *PrimaryArchivalTarget) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### SetTargetNameNil

`func (o *PrimaryArchivalTarget) SetTargetNameNil(b bool)`

 SetTargetNameNil sets the value for TargetName to be an explicit nil

### UnsetTargetName
`func (o *PrimaryArchivalTarget) UnsetTargetName()`

UnsetTargetName ensures that no value is present for TargetName, not even an explicit nil
### GetTierSettings

`func (o *PrimaryArchivalTarget) GetTierSettings() TierLevelSettings`

GetTierSettings returns the TierSettings field if non-nil, zero value otherwise.

### GetTierSettingsOk

`func (o *PrimaryArchivalTarget) GetTierSettingsOk() (*TierLevelSettings, bool)`

GetTierSettingsOk returns a tuple with the TierSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierSettings

`func (o *PrimaryArchivalTarget) SetTierSettings(v TierLevelSettings)`

SetTierSettings sets TierSettings field to given value.

### HasTierSettings

`func (o *PrimaryArchivalTarget) HasTierSettings() bool`

HasTierSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


