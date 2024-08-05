# SupportMfaConfigInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **NullableString** | Specifies email address of the support user. Used when MFA mode is email. | [optional] 
**Enabled** | Pointer to **bool** | Specifies whether MFA is enabled for support user. | [optional] [default to false]
**MfaCode** | Pointer to **NullableString** | MFA code that needs to be passed when disabling MFA or changing email address when email based MFA is configured. | [optional] 
**MfaType** | Pointer to **NullableString** | Specifies the mechanism to receive the OTP code. | [optional] 
**OtpVerificationState** | Pointer to **NullableString** | Specifies the status of otp verification. | [optional] 

## Methods

### NewSupportMfaConfigInfo

`func NewSupportMfaConfigInfo() *SupportMfaConfigInfo`

NewSupportMfaConfigInfo instantiates a new SupportMfaConfigInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportMfaConfigInfoWithDefaults

`func NewSupportMfaConfigInfoWithDefaults() *SupportMfaConfigInfo`

NewSupportMfaConfigInfoWithDefaults instantiates a new SupportMfaConfigInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *SupportMfaConfigInfo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SupportMfaConfigInfo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SupportMfaConfigInfo) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SupportMfaConfigInfo) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *SupportMfaConfigInfo) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *SupportMfaConfigInfo) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetEnabled

`func (o *SupportMfaConfigInfo) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SupportMfaConfigInfo) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SupportMfaConfigInfo) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SupportMfaConfigInfo) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMfaCode

`func (o *SupportMfaConfigInfo) GetMfaCode() string`

GetMfaCode returns the MfaCode field if non-nil, zero value otherwise.

### GetMfaCodeOk

`func (o *SupportMfaConfigInfo) GetMfaCodeOk() (*string, bool)`

GetMfaCodeOk returns a tuple with the MfaCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaCode

`func (o *SupportMfaConfigInfo) SetMfaCode(v string)`

SetMfaCode sets MfaCode field to given value.

### HasMfaCode

`func (o *SupportMfaConfigInfo) HasMfaCode() bool`

HasMfaCode returns a boolean if a field has been set.

### SetMfaCodeNil

`func (o *SupportMfaConfigInfo) SetMfaCodeNil(b bool)`

 SetMfaCodeNil sets the value for MfaCode to be an explicit nil

### UnsetMfaCode
`func (o *SupportMfaConfigInfo) UnsetMfaCode()`

UnsetMfaCode ensures that no value is present for MfaCode, not even an explicit nil
### GetMfaType

`func (o *SupportMfaConfigInfo) GetMfaType() string`

GetMfaType returns the MfaType field if non-nil, zero value otherwise.

### GetMfaTypeOk

`func (o *SupportMfaConfigInfo) GetMfaTypeOk() (*string, bool)`

GetMfaTypeOk returns a tuple with the MfaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaType

`func (o *SupportMfaConfigInfo) SetMfaType(v string)`

SetMfaType sets MfaType field to given value.

### HasMfaType

`func (o *SupportMfaConfigInfo) HasMfaType() bool`

HasMfaType returns a boolean if a field has been set.

### SetMfaTypeNil

`func (o *SupportMfaConfigInfo) SetMfaTypeNil(b bool)`

 SetMfaTypeNil sets the value for MfaType to be an explicit nil

### UnsetMfaType
`func (o *SupportMfaConfigInfo) UnsetMfaType()`

UnsetMfaType ensures that no value is present for MfaType, not even an explicit nil
### GetOtpVerificationState

`func (o *SupportMfaConfigInfo) GetOtpVerificationState() string`

GetOtpVerificationState returns the OtpVerificationState field if non-nil, zero value otherwise.

### GetOtpVerificationStateOk

`func (o *SupportMfaConfigInfo) GetOtpVerificationStateOk() (*string, bool)`

GetOtpVerificationStateOk returns a tuple with the OtpVerificationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpVerificationState

`func (o *SupportMfaConfigInfo) SetOtpVerificationState(v string)`

SetOtpVerificationState sets OtpVerificationState field to given value.

### HasOtpVerificationState

`func (o *SupportMfaConfigInfo) HasOtpVerificationState() bool`

HasOtpVerificationState returns a boolean if a field has been set.

### SetOtpVerificationStateNil

`func (o *SupportMfaConfigInfo) SetOtpVerificationStateNil(b bool)`

 SetOtpVerificationStateNil sets the value for OtpVerificationState to be an explicit nil

### UnsetOtpVerificationState
`func (o *SupportMfaConfigInfo) UnsetOtpVerificationState()`

UnsetOtpVerificationState ensures that no value is present for OtpVerificationState, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


