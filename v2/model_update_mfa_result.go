/*
Cohesity REST API

Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
)

// checks if the UpdateMFAResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateMFAResult{}

// UpdateMFAResult struct for UpdateMFAResult
type UpdateMFAResult struct {
	// Specifies the TOTP account name to be configured for support user.
	AccountName *string `json:"accountName,omitempty"`
	// Specifies the TOTP secret key.
	TotpSecretKey *string `json:"totpSecretKey,omitempty"`
	// Specifies the TOTP key URI for generating MFA QR code.
	TotpUri *string `json:"totpUri,omitempty"`
	// Specifies email address of the support user. Used when MFA mode is email.
	Email NullableString `json:"email,omitempty"`
	// Specifies whether MFA is enabled for support user.
	Enabled *bool `json:"enabled,omitempty"`
	// MFA code that needs to be passed when disabling MFA or changing email address when email based MFA is configured.
	MfaCode NullableString `json:"mfaCode,omitempty"`
	// Specifies the mechanism to receive the OTP code.
	MfaType NullableString `json:"mfaType,omitempty"`
	// Specifies the status of otp verification.
	OtpVerificationState NullableString `json:"otpVerificationState,omitempty"`
}

// NewUpdateMFAResult instantiates a new UpdateMFAResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateMFAResult() *UpdateMFAResult {
	this := UpdateMFAResult{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// NewUpdateMFAResultWithDefaults instantiates a new UpdateMFAResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateMFAResultWithDefaults() *UpdateMFAResult {
	this := UpdateMFAResult{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *UpdateMFAResult) GetAccountName() string {
	if o == nil || IsNil(o.AccountName) {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMFAResult) GetAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.AccountName) {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *UpdateMFAResult) HasAccountName() bool {
	if o != nil && !IsNil(o.AccountName) {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *UpdateMFAResult) SetAccountName(v string) {
	o.AccountName = &v
}

// GetTotpSecretKey returns the TotpSecretKey field value if set, zero value otherwise.
func (o *UpdateMFAResult) GetTotpSecretKey() string {
	if o == nil || IsNil(o.TotpSecretKey) {
		var ret string
		return ret
	}
	return *o.TotpSecretKey
}

// GetTotpSecretKeyOk returns a tuple with the TotpSecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMFAResult) GetTotpSecretKeyOk() (*string, bool) {
	if o == nil || IsNil(o.TotpSecretKey) {
		return nil, false
	}
	return o.TotpSecretKey, true
}

// HasTotpSecretKey returns a boolean if a field has been set.
func (o *UpdateMFAResult) HasTotpSecretKey() bool {
	if o != nil && !IsNil(o.TotpSecretKey) {
		return true
	}

	return false
}

// SetTotpSecretKey gets a reference to the given string and assigns it to the TotpSecretKey field.
func (o *UpdateMFAResult) SetTotpSecretKey(v string) {
	o.TotpSecretKey = &v
}

// GetTotpUri returns the TotpUri field value if set, zero value otherwise.
func (o *UpdateMFAResult) GetTotpUri() string {
	if o == nil || IsNil(o.TotpUri) {
		var ret string
		return ret
	}
	return *o.TotpUri
}

// GetTotpUriOk returns a tuple with the TotpUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMFAResult) GetTotpUriOk() (*string, bool) {
	if o == nil || IsNil(o.TotpUri) {
		return nil, false
	}
	return o.TotpUri, true
}

// HasTotpUri returns a boolean if a field has been set.
func (o *UpdateMFAResult) HasTotpUri() bool {
	if o != nil && !IsNil(o.TotpUri) {
		return true
	}

	return false
}

// SetTotpUri gets a reference to the given string and assigns it to the TotpUri field.
func (o *UpdateMFAResult) SetTotpUri(v string) {
	o.TotpUri = &v
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateMFAResult) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateMFAResult) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *UpdateMFAResult) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *UpdateMFAResult) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *UpdateMFAResult) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *UpdateMFAResult) UnsetEmail() {
	o.Email.Unset()
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateMFAResult) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMFAResult) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateMFAResult) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateMFAResult) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetMfaCode returns the MfaCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateMFAResult) GetMfaCode() string {
	if o == nil || IsNil(o.MfaCode.Get()) {
		var ret string
		return ret
	}
	return *o.MfaCode.Get()
}

// GetMfaCodeOk returns a tuple with the MfaCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateMFAResult) GetMfaCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MfaCode.Get(), o.MfaCode.IsSet()
}

// HasMfaCode returns a boolean if a field has been set.
func (o *UpdateMFAResult) HasMfaCode() bool {
	if o != nil && o.MfaCode.IsSet() {
		return true
	}

	return false
}

// SetMfaCode gets a reference to the given NullableString and assigns it to the MfaCode field.
func (o *UpdateMFAResult) SetMfaCode(v string) {
	o.MfaCode.Set(&v)
}
// SetMfaCodeNil sets the value for MfaCode to be an explicit nil
func (o *UpdateMFAResult) SetMfaCodeNil() {
	o.MfaCode.Set(nil)
}

// UnsetMfaCode ensures that no value is present for MfaCode, not even an explicit nil
func (o *UpdateMFAResult) UnsetMfaCode() {
	o.MfaCode.Unset()
}

// GetMfaType returns the MfaType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateMFAResult) GetMfaType() string {
	if o == nil || IsNil(o.MfaType.Get()) {
		var ret string
		return ret
	}
	return *o.MfaType.Get()
}

// GetMfaTypeOk returns a tuple with the MfaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateMFAResult) GetMfaTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MfaType.Get(), o.MfaType.IsSet()
}

// HasMfaType returns a boolean if a field has been set.
func (o *UpdateMFAResult) HasMfaType() bool {
	if o != nil && o.MfaType.IsSet() {
		return true
	}

	return false
}

// SetMfaType gets a reference to the given NullableString and assigns it to the MfaType field.
func (o *UpdateMFAResult) SetMfaType(v string) {
	o.MfaType.Set(&v)
}
// SetMfaTypeNil sets the value for MfaType to be an explicit nil
func (o *UpdateMFAResult) SetMfaTypeNil() {
	o.MfaType.Set(nil)
}

// UnsetMfaType ensures that no value is present for MfaType, not even an explicit nil
func (o *UpdateMFAResult) UnsetMfaType() {
	o.MfaType.Unset()
}

// GetOtpVerificationState returns the OtpVerificationState field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateMFAResult) GetOtpVerificationState() string {
	if o == nil || IsNil(o.OtpVerificationState.Get()) {
		var ret string
		return ret
	}
	return *o.OtpVerificationState.Get()
}

// GetOtpVerificationStateOk returns a tuple with the OtpVerificationState field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateMFAResult) GetOtpVerificationStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OtpVerificationState.Get(), o.OtpVerificationState.IsSet()
}

// HasOtpVerificationState returns a boolean if a field has been set.
func (o *UpdateMFAResult) HasOtpVerificationState() bool {
	if o != nil && o.OtpVerificationState.IsSet() {
		return true
	}

	return false
}

// SetOtpVerificationState gets a reference to the given NullableString and assigns it to the OtpVerificationState field.
func (o *UpdateMFAResult) SetOtpVerificationState(v string) {
	o.OtpVerificationState.Set(&v)
}
// SetOtpVerificationStateNil sets the value for OtpVerificationState to be an explicit nil
func (o *UpdateMFAResult) SetOtpVerificationStateNil() {
	o.OtpVerificationState.Set(nil)
}

// UnsetOtpVerificationState ensures that no value is present for OtpVerificationState, not even an explicit nil
func (o *UpdateMFAResult) UnsetOtpVerificationState() {
	o.OtpVerificationState.Unset()
}

func (o UpdateMFAResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateMFAResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountName) {
		toSerialize["accountName"] = o.AccountName
	}
	if !IsNil(o.TotpSecretKey) {
		toSerialize["totpSecretKey"] = o.TotpSecretKey
	}
	if !IsNil(o.TotpUri) {
		toSerialize["totpUri"] = o.TotpUri
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if o.MfaCode.IsSet() {
		toSerialize["mfaCode"] = o.MfaCode.Get()
	}
	if o.MfaType.IsSet() {
		toSerialize["mfaType"] = o.MfaType.Get()
	}
	if o.OtpVerificationState.IsSet() {
		toSerialize["otpVerificationState"] = o.OtpVerificationState.Get()
	}
	return toSerialize, nil
}

type NullableUpdateMFAResult struct {
	value *UpdateMFAResult
	isSet bool
}

func (v NullableUpdateMFAResult) Get() *UpdateMFAResult {
	return v.value
}

func (v *NullableUpdateMFAResult) Set(val *UpdateMFAResult) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMFAResult) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMFAResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMFAResult(val *UpdateMFAResult) *NullableUpdateMFAResult {
	return &NullableUpdateMFAResult{value: val, isSet: true}
}

func (v NullableUpdateMFAResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMFAResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


