# ArchivalConfigAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetId** | **NullableInt64** | Specifies the Archival target to copy the Snapshots to. | 
**TargetName** | Pointer to **NullableString** | Specifies the Archival target name where Snapshots are copied. | [optional] [readonly] 
**TargetType** | Pointer to **NullableString** | Specifies the Archival target type where Snapshots are copied. | [optional] [readonly] 
**TierSettings** | Pointer to [**TierLevelSettings**](TierLevelSettings.md) |  | [optional] 
**ExtendedRetention** | Pointer to [**[]ExtendedRetentionPolicy**](ExtendedRetentionPolicy.md) | Specifies additional retention policies that should be applied to the archived backup. Archived backup snapshot will be retained up to a time that is the maximum of all retention policies that are applicable to it. | [optional] 

## Methods

### NewArchivalConfigAllOf

`func NewArchivalConfigAllOf(targetId NullableInt64, ) *ArchivalConfigAllOf`

NewArchivalConfigAllOf instantiates a new ArchivalConfigAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArchivalConfigAllOfWithDefaults

`func NewArchivalConfigAllOfWithDefaults() *ArchivalConfigAllOf`

NewArchivalConfigAllOfWithDefaults instantiates a new ArchivalConfigAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetId

`func (o *ArchivalConfigAllOf) GetTargetId() int64`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *ArchivalConfigAllOf) GetTargetIdOk() (*int64, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *ArchivalConfigAllOf) SetTargetId(v int64)`

SetTargetId sets TargetId field to given value.


### SetTargetIdNil

`func (o *ArchivalConfigAllOf) SetTargetIdNil(b bool)`

 SetTargetIdNil sets the value for TargetId to be an explicit nil

### UnsetTargetId
`func (o *ArchivalConfigAllOf) UnsetTargetId()`

UnsetTargetId ensures that no value is present for TargetId, not even an explicit nil
### GetTargetName

`func (o *ArchivalConfigAllOf) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *ArchivalConfigAllOf) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *ArchivalConfigAllOf) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *ArchivalConfigAllOf) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### SetTargetNameNil

`func (o *ArchivalConfigAllOf) SetTargetNameNil(b bool)`

 SetTargetNameNil sets the value for TargetName to be an explicit nil

### UnsetTargetName
`func (o *ArchivalConfigAllOf) UnsetTargetName()`

UnsetTargetName ensures that no value is present for TargetName, not even an explicit nil
### GetTargetType

`func (o *ArchivalConfigAllOf) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *ArchivalConfigAllOf) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *ArchivalConfigAllOf) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *ArchivalConfigAllOf) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### SetTargetTypeNil

`func (o *ArchivalConfigAllOf) SetTargetTypeNil(b bool)`

 SetTargetTypeNil sets the value for TargetType to be an explicit nil

### UnsetTargetType
`func (o *ArchivalConfigAllOf) UnsetTargetType()`

UnsetTargetType ensures that no value is present for TargetType, not even an explicit nil
### GetTierSettings

`func (o *ArchivalConfigAllOf) GetTierSettings() TierLevelSettings`

GetTierSettings returns the TierSettings field if non-nil, zero value otherwise.

### GetTierSettingsOk

`func (o *ArchivalConfigAllOf) GetTierSettingsOk() (*TierLevelSettings, bool)`

GetTierSettingsOk returns a tuple with the TierSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierSettings

`func (o *ArchivalConfigAllOf) SetTierSettings(v TierLevelSettings)`

SetTierSettings sets TierSettings field to given value.

### HasTierSettings

`func (o *ArchivalConfigAllOf) HasTierSettings() bool`

HasTierSettings returns a boolean if a field has been set.

### GetExtendedRetention

`func (o *ArchivalConfigAllOf) GetExtendedRetention() []ExtendedRetentionPolicy`

GetExtendedRetention returns the ExtendedRetention field if non-nil, zero value otherwise.

### GetExtendedRetentionOk

`func (o *ArchivalConfigAllOf) GetExtendedRetentionOk() (*[]ExtendedRetentionPolicy, bool)`

GetExtendedRetentionOk returns a tuple with the ExtendedRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedRetention

`func (o *ArchivalConfigAllOf) SetExtendedRetention(v []ExtendedRetentionPolicy)`

SetExtendedRetention sets ExtendedRetention field to given value.

### HasExtendedRetention

`func (o *ArchivalConfigAllOf) HasExtendedRetention() bool`

HasExtendedRetention returns a boolean if a field has been set.

### SetExtendedRetentionNil

`func (o *ArchivalConfigAllOf) SetExtendedRetentionNil(b bool)`

 SetExtendedRetentionNil sets the value for ExtendedRetention to be an explicit nil

### UnsetExtendedRetention
`func (o *ArchivalConfigAllOf) UnsetExtendedRetention()`

UnsetExtendedRetention ensures that no value is present for ExtendedRetention, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


