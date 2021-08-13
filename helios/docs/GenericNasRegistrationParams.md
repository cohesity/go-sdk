# GenericNasRegistrationParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MountPoint** | **NullableString** | Specifies the MountPoint for Generic NAS Source. | 
**Mode** | **NullableString** | Specifies the mode of the source. &#39;kNfs3&#39; indicates NFS mode. &#39;kCifs1&#39; indicates SMB mode. | 
**Description** | Pointer to **NullableString** | Specifies the Description for Generic NAS Source. | [optional] 
**SkipValidation** | Pointer to **NullableBool** | Specifies if validation has to be skipped while registering the mount point. | [optional] 
**SmbMountCredentials** | Pointer to [**SmbMountCredentials**](SmbMountCredentials.md) |  | [optional] 
**ThrottlingConfig** | Pointer to [**NasThrottlingConfig**](NasThrottlingConfig.md) |  | [optional] 

## Methods

### NewGenericNasRegistrationParams

`func NewGenericNasRegistrationParams(mountPoint NullableString, mode NullableString, ) *GenericNasRegistrationParams`

NewGenericNasRegistrationParams instantiates a new GenericNasRegistrationParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericNasRegistrationParamsWithDefaults

`func NewGenericNasRegistrationParamsWithDefaults() *GenericNasRegistrationParams`

NewGenericNasRegistrationParamsWithDefaults instantiates a new GenericNasRegistrationParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMountPoint

`func (o *GenericNasRegistrationParams) GetMountPoint() string`

GetMountPoint returns the MountPoint field if non-nil, zero value otherwise.

### GetMountPointOk

`func (o *GenericNasRegistrationParams) GetMountPointOk() (*string, bool)`

GetMountPointOk returns a tuple with the MountPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPoint

`func (o *GenericNasRegistrationParams) SetMountPoint(v string)`

SetMountPoint sets MountPoint field to given value.


### SetMountPointNil

`func (o *GenericNasRegistrationParams) SetMountPointNil(b bool)`

 SetMountPointNil sets the value for MountPoint to be an explicit nil

### UnsetMountPoint
`func (o *GenericNasRegistrationParams) UnsetMountPoint()`

UnsetMountPoint ensures that no value is present for MountPoint, not even an explicit nil
### GetMode

`func (o *GenericNasRegistrationParams) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *GenericNasRegistrationParams) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *GenericNasRegistrationParams) SetMode(v string)`

SetMode sets Mode field to given value.


### SetModeNil

`func (o *GenericNasRegistrationParams) SetModeNil(b bool)`

 SetModeNil sets the value for Mode to be an explicit nil

### UnsetMode
`func (o *GenericNasRegistrationParams) UnsetMode()`

UnsetMode ensures that no value is present for Mode, not even an explicit nil
### GetDescription

`func (o *GenericNasRegistrationParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GenericNasRegistrationParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GenericNasRegistrationParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GenericNasRegistrationParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GenericNasRegistrationParams) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GenericNasRegistrationParams) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSkipValidation

`func (o *GenericNasRegistrationParams) GetSkipValidation() bool`

GetSkipValidation returns the SkipValidation field if non-nil, zero value otherwise.

### GetSkipValidationOk

`func (o *GenericNasRegistrationParams) GetSkipValidationOk() (*bool, bool)`

GetSkipValidationOk returns a tuple with the SkipValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipValidation

`func (o *GenericNasRegistrationParams) SetSkipValidation(v bool)`

SetSkipValidation sets SkipValidation field to given value.

### HasSkipValidation

`func (o *GenericNasRegistrationParams) HasSkipValidation() bool`

HasSkipValidation returns a boolean if a field has been set.

### SetSkipValidationNil

`func (o *GenericNasRegistrationParams) SetSkipValidationNil(b bool)`

 SetSkipValidationNil sets the value for SkipValidation to be an explicit nil

### UnsetSkipValidation
`func (o *GenericNasRegistrationParams) UnsetSkipValidation()`

UnsetSkipValidation ensures that no value is present for SkipValidation, not even an explicit nil
### GetSmbMountCredentials

`func (o *GenericNasRegistrationParams) GetSmbMountCredentials() SmbMountCredentials`

GetSmbMountCredentials returns the SmbMountCredentials field if non-nil, zero value otherwise.

### GetSmbMountCredentialsOk

`func (o *GenericNasRegistrationParams) GetSmbMountCredentialsOk() (*SmbMountCredentials, bool)`

GetSmbMountCredentialsOk returns a tuple with the SmbMountCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbMountCredentials

`func (o *GenericNasRegistrationParams) SetSmbMountCredentials(v SmbMountCredentials)`

SetSmbMountCredentials sets SmbMountCredentials field to given value.

### HasSmbMountCredentials

`func (o *GenericNasRegistrationParams) HasSmbMountCredentials() bool`

HasSmbMountCredentials returns a boolean if a field has been set.

### GetThrottlingConfig

`func (o *GenericNasRegistrationParams) GetThrottlingConfig() NasThrottlingConfig`

GetThrottlingConfig returns the ThrottlingConfig field if non-nil, zero value otherwise.

### GetThrottlingConfigOk

`func (o *GenericNasRegistrationParams) GetThrottlingConfigOk() (*NasThrottlingConfig, bool)`

GetThrottlingConfigOk returns a tuple with the ThrottlingConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottlingConfig

`func (o *GenericNasRegistrationParams) SetThrottlingConfig(v NasThrottlingConfig)`

SetThrottlingConfig sets ThrottlingConfig field to given value.

### HasThrottlingConfig

`func (o *GenericNasRegistrationParams) HasThrottlingConfig() bool`

HasThrottlingConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


