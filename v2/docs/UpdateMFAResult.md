# UpdateMFAResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountName** | Pointer to **string** | Specifies the TOTP account name to be configured for support user. | [optional] 
**TotpSecretKey** | Pointer to **string** | Specifies the TOTP secret key. | [optional] 
**TotpUri** | Pointer to **string** | Specifies the TOTP key URI for generating MFA QR code. | [optional] 
**Email** | Pointer to **NullableString** | Specifies email address of the support user. Used when MFA mode is email. | [optional] 
**Enabled** | Pointer to **bool** | Specifies whether MFA is enabled for support user. | [optional] [default to false]
**MfaCode** | Pointer to **NullableString** | MFA code that needs to be passed when disabling MFA or changing email address when email based MFA is configured. | [optional] 
**MfaType** | Pointer to **NullableString** | Specifies the mechanism to receive the OTP code. | [optional] 
**OtpVerificationState** | Pointer to **NullableString** | Specifies the status of otp verification. | [optional] 

## Methods

### NewUpdateMFAResult

`func NewUpdateMFAResult() *UpdateMFAResult`

NewUpdateMFAResult instantiates a new UpdateMFAResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMFAResultWithDefaults

`func NewUpdateMFAResultWithDefaults() *UpdateMFAResult`

NewUpdateMFAResultWithDefaults instantiates a new UpdateMFAResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountName

`func (o *UpdateMFAResult) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *UpdateMFAResult) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *UpdateMFAResult) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *UpdateMFAResult) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetTotpSecretKey

`func (o *UpdateMFAResult) GetTotpSecretKey() string`

GetTotpSecretKey returns the TotpSecretKey field if non-nil, zero value otherwise.

### GetTotpSecretKeyOk

`func (o *UpdateMFAResult) GetTotpSecretKeyOk() (*string, bool)`

GetTotpSecretKeyOk returns a tuple with the TotpSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpSecretKey

`func (o *UpdateMFAResult) SetTotpSecretKey(v string)`

SetTotpSecretKey sets TotpSecretKey field to given value.

### HasTotpSecretKey

`func (o *UpdateMFAResult) HasTotpSecretKey() bool`

HasTotpSecretKey returns a boolean if a field has been set.

### GetTotpUri

`func (o *UpdateMFAResult) GetTotpUri() string`

GetTotpUri returns the TotpUri field if non-nil, zero value otherwise.

### GetTotpUriOk

`func (o *UpdateMFAResult) GetTotpUriOk() (*string, bool)`

GetTotpUriOk returns a tuple with the TotpUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpUri

`func (o *UpdateMFAResult) SetTotpUri(v string)`

SetTotpUri sets TotpUri field to given value.

### HasTotpUri

`func (o *UpdateMFAResult) HasTotpUri() bool`

HasTotpUri returns a boolean if a field has been set.

### GetEmail

`func (o *UpdateMFAResult) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateMFAResult) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateMFAResult) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateMFAResult) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *UpdateMFAResult) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *UpdateMFAResult) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetEnabled

`func (o *UpdateMFAResult) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateMFAResult) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateMFAResult) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateMFAResult) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMfaCode

`func (o *UpdateMFAResult) GetMfaCode() string`

GetMfaCode returns the MfaCode field if non-nil, zero value otherwise.

### GetMfaCodeOk

`func (o *UpdateMFAResult) GetMfaCodeOk() (*string, bool)`

GetMfaCodeOk returns a tuple with the MfaCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaCode

`func (o *UpdateMFAResult) SetMfaCode(v string)`

SetMfaCode sets MfaCode field to given value.

### HasMfaCode

`func (o *UpdateMFAResult) HasMfaCode() bool`

HasMfaCode returns a boolean if a field has been set.

### SetMfaCodeNil

`func (o *UpdateMFAResult) SetMfaCodeNil(b bool)`

 SetMfaCodeNil sets the value for MfaCode to be an explicit nil

### UnsetMfaCode
`func (o *UpdateMFAResult) UnsetMfaCode()`

UnsetMfaCode ensures that no value is present for MfaCode, not even an explicit nil
### GetMfaType

`func (o *UpdateMFAResult) GetMfaType() string`

GetMfaType returns the MfaType field if non-nil, zero value otherwise.

### GetMfaTypeOk

`func (o *UpdateMFAResult) GetMfaTypeOk() (*string, bool)`

GetMfaTypeOk returns a tuple with the MfaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaType

`func (o *UpdateMFAResult) SetMfaType(v string)`

SetMfaType sets MfaType field to given value.

### HasMfaType

`func (o *UpdateMFAResult) HasMfaType() bool`

HasMfaType returns a boolean if a field has been set.

### SetMfaTypeNil

`func (o *UpdateMFAResult) SetMfaTypeNil(b bool)`

 SetMfaTypeNil sets the value for MfaType to be an explicit nil

### UnsetMfaType
`func (o *UpdateMFAResult) UnsetMfaType()`

UnsetMfaType ensures that no value is present for MfaType, not even an explicit nil
### GetOtpVerificationState

`func (o *UpdateMFAResult) GetOtpVerificationState() string`

GetOtpVerificationState returns the OtpVerificationState field if non-nil, zero value otherwise.

### GetOtpVerificationStateOk

`func (o *UpdateMFAResult) GetOtpVerificationStateOk() (*string, bool)`

GetOtpVerificationStateOk returns a tuple with the OtpVerificationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpVerificationState

`func (o *UpdateMFAResult) SetOtpVerificationState(v string)`

SetOtpVerificationState sets OtpVerificationState field to given value.

### HasOtpVerificationState

`func (o *UpdateMFAResult) HasOtpVerificationState() bool`

HasOtpVerificationState returns a boolean if a field has been set.

### SetOtpVerificationStateNil

`func (o *UpdateMFAResult) SetOtpVerificationStateNil(b bool)`

 SetOtpVerificationStateNil sets the value for OtpVerificationState to be an explicit nil

### UnsetOtpVerificationState
`func (o *UpdateMFAResult) UnsetOtpVerificationState()`

UnsetOtpVerificationState ensures that no value is present for OtpVerificationState, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


