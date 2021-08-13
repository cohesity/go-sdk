/*
 * Cohesity REST API
 *
 * Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

// package helios
package models

import (
	"encoding/json"
	. "github.com/cohesity/go-sdk/helios/utils"
)

var _ = NullableBool{}

// UpdateMFAResult struct for UpdateMFAResult
type UpdateMFAResult struct {
	// Specifies the TOTP secret key.
	TotpSecretKey *string `json:"totpSecretKey,omitempty"`
	// Specifies the TOTP key URI for generating MFA QR code.
	TotpUri *string `json:"totpUri,omitempty"`
	// Specifies the TOTP account name to be configured for support user.
	AccountName *string `json:"accountName,omitempty"`
	// Specifies whether MFA is enabled for support user.
	Enabled *bool `json:"enabled,omitempty"`
	// Specifies the mechanism to receive the OTP code.
	MfaType NullableString `json:"mfaType,omitempty"`
	// Specifies email address of the support user. Used when MFA mode is email.
	Email NullableString `json:"email,omitempty"`
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

// GetTotpSecretKey returns the TotpSecretKey field value if set, zero value otherwise.
func (o *UpdateMFAResult) GetTotpSecretKey() string {
	if o == nil || o.TotpSecretKey == nil {
		var ret string
		return ret
	}
	return *o.TotpSecretKey
}

// GetTotpSecretKeyOk returns a tuple with the TotpSecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMFAResult) GetTotpSecretKeyOk() (*string, bool) {
	if o == nil || o.TotpSecretKey == nil {
		return nil, false
	}
	return o.TotpSecretKey, true
}

// HasTotpSecretKey returns a boolean if a field has been set.
func (o *UpdateMFAResult) HasTotpSecretKey() bool {
	if o != nil && o.TotpSecretKey != nil {
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
	if o == nil || o.TotpUri == nil {
		var ret string
		return ret
	}
	return *o.TotpUri
}

// GetTotpUriOk returns a tuple with the TotpUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMFAResult) GetTotpUriOk() (*string, bool) {
	if o == nil || o.TotpUri == nil {
		return nil, false
	}
	return o.TotpUri, true
}

// HasTotpUri returns a boolean if a field has been set.
func (o *UpdateMFAResult) HasTotpUri() bool {
	if o != nil && o.TotpUri != nil {
		return true
	}

	return false
}

// SetTotpUri gets a reference to the given string and assigns it to the TotpUri field.
func (o *UpdateMFAResult) SetTotpUri(v string) {
	o.TotpUri = &v
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *UpdateMFAResult) GetAccountName() string {
	if o == nil || o.AccountName == nil {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMFAResult) GetAccountNameOk() (*string, bool) {
	if o == nil || o.AccountName == nil {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *UpdateMFAResult) HasAccountName() bool {
	if o != nil && o.AccountName != nil {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *UpdateMFAResult) SetAccountName(v string) {
	o.AccountName = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateMFAResult) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMFAResult) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateMFAResult) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateMFAResult) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetMfaType returns the MfaType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateMFAResult) GetMfaType() string {
	if o == nil || o.MfaType.Get() == nil {
		var ret string
		return ret
	}
	return *o.MfaType.Get()
}

// GetMfaTypeOk returns a tuple with the MfaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateMFAResult) GetMfaTypeOk() (*string, bool) {
	if o == nil  {
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

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateMFAResult) GetEmail() string {
	if o == nil || o.Email.Get() == nil {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateMFAResult) GetEmailOk() (*string, bool) {
	if o == nil  {
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

func (o UpdateMFAResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TotpSecretKey != nil {
		toSerialize["totpSecretKey"] = o.TotpSecretKey
	}
	if o.TotpUri != nil {
		toSerialize["totpUri"] = o.TotpUri
	}
	if o.AccountName != nil {
		toSerialize["accountName"] = o.AccountName
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.MfaType.IsSet() {
		toSerialize["mfaType"] = o.MfaType.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	return json.Marshal(toSerialize)
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


