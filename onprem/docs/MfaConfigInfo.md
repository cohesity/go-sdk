# MfaConfigInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Specifies whether MFA is enabled on a cluster level. | [optional] [default to false]
**AuthenticationTypes** | Pointer to **[]string** | Specifies the list of mechanism to receive the OTP code. | [optional] 
**RetainUserMfaSettings** | Pointer to **NullableBool** | Specifies whether user MFA setting needs to be retained. | [optional] 

## Methods

### NewMfaConfigInfo

`func NewMfaConfigInfo() *MfaConfigInfo`

NewMfaConfigInfo instantiates a new MfaConfigInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMfaConfigInfoWithDefaults

`func NewMfaConfigInfoWithDefaults() *MfaConfigInfo`

NewMfaConfigInfoWithDefaults instantiates a new MfaConfigInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *MfaConfigInfo) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MfaConfigInfo) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MfaConfigInfo) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MfaConfigInfo) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAuthenticationTypes

`func (o *MfaConfigInfo) GetAuthenticationTypes() []string`

GetAuthenticationTypes returns the AuthenticationTypes field if non-nil, zero value otherwise.

### GetAuthenticationTypesOk

`func (o *MfaConfigInfo) GetAuthenticationTypesOk() (*[]string, bool)`

GetAuthenticationTypesOk returns a tuple with the AuthenticationTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationTypes

`func (o *MfaConfigInfo) SetAuthenticationTypes(v []string)`

SetAuthenticationTypes sets AuthenticationTypes field to given value.

### HasAuthenticationTypes

`func (o *MfaConfigInfo) HasAuthenticationTypes() bool`

HasAuthenticationTypes returns a boolean if a field has been set.

### SetAuthenticationTypesNil

`func (o *MfaConfigInfo) SetAuthenticationTypesNil(b bool)`

 SetAuthenticationTypesNil sets the value for AuthenticationTypes to be an explicit nil

### UnsetAuthenticationTypes
`func (o *MfaConfigInfo) UnsetAuthenticationTypes()`

UnsetAuthenticationTypes ensures that no value is present for AuthenticationTypes, not even an explicit nil
### GetRetainUserMfaSettings

`func (o *MfaConfigInfo) GetRetainUserMfaSettings() bool`

GetRetainUserMfaSettings returns the RetainUserMfaSettings field if non-nil, zero value otherwise.

### GetRetainUserMfaSettingsOk

`func (o *MfaConfigInfo) GetRetainUserMfaSettingsOk() (*bool, bool)`

GetRetainUserMfaSettingsOk returns a tuple with the RetainUserMfaSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainUserMfaSettings

`func (o *MfaConfigInfo) SetRetainUserMfaSettings(v bool)`

SetRetainUserMfaSettings sets RetainUserMfaSettings field to given value.

### HasRetainUserMfaSettings

`func (o *MfaConfigInfo) HasRetainUserMfaSettings() bool`

HasRetainUserMfaSettings returns a boolean if a field has been set.

### SetRetainUserMfaSettingsNil

`func (o *MfaConfigInfo) SetRetainUserMfaSettingsNil(b bool)`

 SetRetainUserMfaSettingsNil sets the value for RetainUserMfaSettings to be an explicit nil

### UnsetRetainUserMfaSettings
`func (o *MfaConfigInfo) UnsetRetainUserMfaSettings()`

UnsetRetainUserMfaSettings ensures that no value is present for RetainUserMfaSettings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


