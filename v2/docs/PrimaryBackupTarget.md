# PrimaryBackupTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchivalTargetSettings** | Pointer to [**PrimaryArchivalTarget**](PrimaryArchivalTarget.md) |  | [optional] 
**TargetType** | Pointer to **NullableString** | Specifies the primary backup location where backups will be stored. If not specified, then default is assumed as local backup on Cohesity cluster. | [optional] [default to "Local"]
**UseDefaultBackupTarget** | Pointer to **NullableBool** | Specifies if the default primary backup target must be used for backups. If this is not specified or set to false, then targets specified in &#39;archivalTargetSettings&#39; will be used for backups. If the value is specified as true, then default backup target is used internally. This field should only be set in the environment where tenant policy management is enabled and external targets are assigned to tenant when provisioning tenants. | [optional] 

## Methods

### NewPrimaryBackupTarget

`func NewPrimaryBackupTarget() *PrimaryBackupTarget`

NewPrimaryBackupTarget instantiates a new PrimaryBackupTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrimaryBackupTargetWithDefaults

`func NewPrimaryBackupTargetWithDefaults() *PrimaryBackupTarget`

NewPrimaryBackupTargetWithDefaults instantiates a new PrimaryBackupTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchivalTargetSettings

`func (o *PrimaryBackupTarget) GetArchivalTargetSettings() PrimaryArchivalTarget`

GetArchivalTargetSettings returns the ArchivalTargetSettings field if non-nil, zero value otherwise.

### GetArchivalTargetSettingsOk

`func (o *PrimaryBackupTarget) GetArchivalTargetSettingsOk() (*PrimaryArchivalTarget, bool)`

GetArchivalTargetSettingsOk returns a tuple with the ArchivalTargetSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivalTargetSettings

`func (o *PrimaryBackupTarget) SetArchivalTargetSettings(v PrimaryArchivalTarget)`

SetArchivalTargetSettings sets ArchivalTargetSettings field to given value.

### HasArchivalTargetSettings

`func (o *PrimaryBackupTarget) HasArchivalTargetSettings() bool`

HasArchivalTargetSettings returns a boolean if a field has been set.

### GetTargetType

`func (o *PrimaryBackupTarget) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *PrimaryBackupTarget) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *PrimaryBackupTarget) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *PrimaryBackupTarget) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### SetTargetTypeNil

`func (o *PrimaryBackupTarget) SetTargetTypeNil(b bool)`

 SetTargetTypeNil sets the value for TargetType to be an explicit nil

### UnsetTargetType
`func (o *PrimaryBackupTarget) UnsetTargetType()`

UnsetTargetType ensures that no value is present for TargetType, not even an explicit nil
### GetUseDefaultBackupTarget

`func (o *PrimaryBackupTarget) GetUseDefaultBackupTarget() bool`

GetUseDefaultBackupTarget returns the UseDefaultBackupTarget field if non-nil, zero value otherwise.

### GetUseDefaultBackupTargetOk

`func (o *PrimaryBackupTarget) GetUseDefaultBackupTargetOk() (*bool, bool)`

GetUseDefaultBackupTargetOk returns a tuple with the UseDefaultBackupTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDefaultBackupTarget

`func (o *PrimaryBackupTarget) SetUseDefaultBackupTarget(v bool)`

SetUseDefaultBackupTarget sets UseDefaultBackupTarget field to given value.

### HasUseDefaultBackupTarget

`func (o *PrimaryBackupTarget) HasUseDefaultBackupTarget() bool`

HasUseDefaultBackupTarget returns a boolean if a field has been set.

### SetUseDefaultBackupTargetNil

`func (o *PrimaryBackupTarget) SetUseDefaultBackupTargetNil(b bool)`

 SetUseDefaultBackupTargetNil sets the value for UseDefaultBackupTarget to be an explicit nil

### UnsetUseDefaultBackupTarget
`func (o *PrimaryBackupTarget) UnsetUseDefaultBackupTarget()`

UnsetUseDefaultBackupTarget ensures that no value is present for UseDefaultBackupTarget, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


