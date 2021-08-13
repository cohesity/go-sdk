# PrimaryBackupTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetType** | Pointer to **NullableString** | Specifies the primary backup location where backups will be stored. If not specified, then default is assumed as local backup on Cohesity cluster. | [optional] 
**ArchivalTargetSettings** | Pointer to [**PrimaryArchivalTarget**](PrimaryArchivalTarget.md) |  | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


