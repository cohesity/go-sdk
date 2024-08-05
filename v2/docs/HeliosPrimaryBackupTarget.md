# HeliosPrimaryBackupTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchivalTargetSettings** | Pointer to [**PrimaryArchivalTarget**](PrimaryArchivalTarget.md) |  | [optional] 
**TargetType** | Pointer to **NullableString** | Specifies the primary backup location where backups will be stored. If not specified, then default is assumed as local backup on Cohesity cluster. | [optional] [default to "Local"]

## Methods

### NewHeliosPrimaryBackupTarget

`func NewHeliosPrimaryBackupTarget() *HeliosPrimaryBackupTarget`

NewHeliosPrimaryBackupTarget instantiates a new HeliosPrimaryBackupTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosPrimaryBackupTargetWithDefaults

`func NewHeliosPrimaryBackupTargetWithDefaults() *HeliosPrimaryBackupTarget`

NewHeliosPrimaryBackupTargetWithDefaults instantiates a new HeliosPrimaryBackupTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchivalTargetSettings

`func (o *HeliosPrimaryBackupTarget) GetArchivalTargetSettings() PrimaryArchivalTarget`

GetArchivalTargetSettings returns the ArchivalTargetSettings field if non-nil, zero value otherwise.

### GetArchivalTargetSettingsOk

`func (o *HeliosPrimaryBackupTarget) GetArchivalTargetSettingsOk() (*PrimaryArchivalTarget, bool)`

GetArchivalTargetSettingsOk returns a tuple with the ArchivalTargetSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivalTargetSettings

`func (o *HeliosPrimaryBackupTarget) SetArchivalTargetSettings(v PrimaryArchivalTarget)`

SetArchivalTargetSettings sets ArchivalTargetSettings field to given value.

### HasArchivalTargetSettings

`func (o *HeliosPrimaryBackupTarget) HasArchivalTargetSettings() bool`

HasArchivalTargetSettings returns a boolean if a field has been set.

### GetTargetType

`func (o *HeliosPrimaryBackupTarget) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *HeliosPrimaryBackupTarget) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *HeliosPrimaryBackupTarget) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *HeliosPrimaryBackupTarget) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### SetTargetTypeNil

`func (o *HeliosPrimaryBackupTarget) SetTargetTypeNil(b bool)`

 SetTargetTypeNil sets the value for TargetType to be an explicit nil

### UnsetTargetType
`func (o *HeliosPrimaryBackupTarget) UnsetTargetType()`

UnsetTargetType ensures that no value is present for TargetType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


