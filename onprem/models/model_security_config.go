/*
 * Cohesity REST API
 *
 * Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

// package onprem
package models

import (
	"encoding/json"
	. "github.com/cohesity/go-sdk/onprem/utils"
	"fmt"
)

var _ = NullableBool{}

// SecurityConfig Specifies the fields of security settings.
type SecurityConfig struct {
	PasswordStrength *SecurityConfigPasswordStrength `json:"passwordStrength,omitempty"`
	PasswordReuse *SecurityConfigPasswordReuse `json:"passwordReuse,omitempty"`
	PasswordLifetime *SecurityConfigPasswordLifetime `json:"passwordLifetime,omitempty"`
	AccountLockout *SecurityConfigAccountLockout `json:"accountLockout,omitempty"`
	DataClassification *SecurityConfigDataClassification `json:"dataClassification,omitempty"`
	SessionConfiguration *SecurityConfigSessionConfiguration `json:"sessionConfiguration,omitempty"`
	CertificateBasedAuth *SecurityConfigCertificateBasedAuth `json:"certificateBasedAuth,omitempty"`
}

// NewSecurityConfig instantiates a new SecurityConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityConfig() *SecurityConfig {
	this := SecurityConfig{}
	return &this
}

// NewSecurityConfigWithDefaults instantiates a new SecurityConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityConfigWithDefaults() *SecurityConfig {
	this := SecurityConfig{}
	return &this
}

// GetPasswordStrength returns the PasswordStrength field value if set, zero value otherwise.
func (o *SecurityConfig) GetPasswordStrength() SecurityConfigPasswordStrength {
	if o == nil || o.PasswordStrength == nil {
		var ret SecurityConfigPasswordStrength
		return ret
	}
	return *o.PasswordStrength
}

// GetPasswordStrengthOk returns a tuple with the PasswordStrength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityConfig) GetPasswordStrengthOk() (*SecurityConfigPasswordStrength, bool) {
	if o == nil || o.PasswordStrength == nil {
		return nil, false
	}
	return o.PasswordStrength, true
}

// HasPasswordStrength returns a boolean if a field has been set.
func (o *SecurityConfig) HasPasswordStrength() bool {
	if o != nil && o.PasswordStrength != nil {
		return true
	}

	return false
}

// SetPasswordStrength gets a reference to the given SecurityConfigPasswordStrength and assigns it to the PasswordStrength field.
func (o *SecurityConfig) SetPasswordStrength(v SecurityConfigPasswordStrength) {
	o.PasswordStrength = &v
}

// GetPasswordReuse returns the PasswordReuse field value if set, zero value otherwise.
func (o *SecurityConfig) GetPasswordReuse() SecurityConfigPasswordReuse {
	if o == nil || o.PasswordReuse == nil {
		var ret SecurityConfigPasswordReuse
		return ret
	}
	return *o.PasswordReuse
}

// GetPasswordReuseOk returns a tuple with the PasswordReuse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityConfig) GetPasswordReuseOk() (*SecurityConfigPasswordReuse, bool) {
	if o == nil || o.PasswordReuse == nil {
		return nil, false
	}
	return o.PasswordReuse, true
}

// HasPasswordReuse returns a boolean if a field has been set.
func (o *SecurityConfig) HasPasswordReuse() bool {
	if o != nil && o.PasswordReuse != nil {
		return true
	}

	return false
}

// SetPasswordReuse gets a reference to the given SecurityConfigPasswordReuse and assigns it to the PasswordReuse field.
func (o *SecurityConfig) SetPasswordReuse(v SecurityConfigPasswordReuse) {
	o.PasswordReuse = &v
}

// GetPasswordLifetime returns the PasswordLifetime field value if set, zero value otherwise.
func (o *SecurityConfig) GetPasswordLifetime() SecurityConfigPasswordLifetime {
	if o == nil || o.PasswordLifetime == nil {
		var ret SecurityConfigPasswordLifetime
		return ret
	}
	return *o.PasswordLifetime
}

// GetPasswordLifetimeOk returns a tuple with the PasswordLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityConfig) GetPasswordLifetimeOk() (*SecurityConfigPasswordLifetime, bool) {
	if o == nil || o.PasswordLifetime == nil {
		return nil, false
	}
	return o.PasswordLifetime, true
}

// HasPasswordLifetime returns a boolean if a field has been set.
func (o *SecurityConfig) HasPasswordLifetime() bool {
	if o != nil && o.PasswordLifetime != nil {
		return true
	}

	return false
}

// SetPasswordLifetime gets a reference to the given SecurityConfigPasswordLifetime and assigns it to the PasswordLifetime field.
func (o *SecurityConfig) SetPasswordLifetime(v SecurityConfigPasswordLifetime) {
	o.PasswordLifetime = &v
}

// GetAccountLockout returns the AccountLockout field value if set, zero value otherwise.
func (o *SecurityConfig) GetAccountLockout() SecurityConfigAccountLockout {
	if o == nil || o.AccountLockout == nil {
		var ret SecurityConfigAccountLockout
		return ret
	}
	return *o.AccountLockout
}

// GetAccountLockoutOk returns a tuple with the AccountLockout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityConfig) GetAccountLockoutOk() (*SecurityConfigAccountLockout, bool) {
	if o == nil || o.AccountLockout == nil {
		return nil, false
	}
	return o.AccountLockout, true
}

// HasAccountLockout returns a boolean if a field has been set.
func (o *SecurityConfig) HasAccountLockout() bool {
	if o != nil && o.AccountLockout != nil {
		return true
	}

	return false
}

// SetAccountLockout gets a reference to the given SecurityConfigAccountLockout and assigns it to the AccountLockout field.
func (o *SecurityConfig) SetAccountLockout(v SecurityConfigAccountLockout) {
	o.AccountLockout = &v
}

// GetDataClassification returns the DataClassification field value if set, zero value otherwise.
func (o *SecurityConfig) GetDataClassification() SecurityConfigDataClassification {
	if o == nil || o.DataClassification == nil {
		var ret SecurityConfigDataClassification
		return ret
	}
	return *o.DataClassification
}

// GetDataClassificationOk returns a tuple with the DataClassification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityConfig) GetDataClassificationOk() (*SecurityConfigDataClassification, bool) {
	if o == nil || o.DataClassification == nil {
		return nil, false
	}
	return o.DataClassification, true
}

// HasDataClassification returns a boolean if a field has been set.
func (o *SecurityConfig) HasDataClassification() bool {
	if o != nil && o.DataClassification != nil {
		return true
	}

	return false
}

// SetDataClassification gets a reference to the given SecurityConfigDataClassification and assigns it to the DataClassification field.
func (o *SecurityConfig) SetDataClassification(v SecurityConfigDataClassification) {
	o.DataClassification = &v
}

// GetSessionConfiguration returns the SessionConfiguration field value if set, zero value otherwise.
func (o *SecurityConfig) GetSessionConfiguration() SecurityConfigSessionConfiguration {
	if o == nil || o.SessionConfiguration == nil {
		var ret SecurityConfigSessionConfiguration
		return ret
	}
	return *o.SessionConfiguration
}

// GetSessionConfigurationOk returns a tuple with the SessionConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityConfig) GetSessionConfigurationOk() (*SecurityConfigSessionConfiguration, bool) {
	if o == nil || o.SessionConfiguration == nil {
		return nil, false
	}
	return o.SessionConfiguration, true
}

// HasSessionConfiguration returns a boolean if a field has been set.
func (o *SecurityConfig) HasSessionConfiguration() bool {
	if o != nil && o.SessionConfiguration != nil {
		return true
	}

	return false
}

// SetSessionConfiguration gets a reference to the given SecurityConfigSessionConfiguration and assigns it to the SessionConfiguration field.
func (o *SecurityConfig) SetSessionConfiguration(v SecurityConfigSessionConfiguration) {
	o.SessionConfiguration = &v
}

// GetCertificateBasedAuth returns the CertificateBasedAuth field value if set, zero value otherwise.
func (o *SecurityConfig) GetCertificateBasedAuth() SecurityConfigCertificateBasedAuth {
	if o == nil || o.CertificateBasedAuth == nil {
		var ret SecurityConfigCertificateBasedAuth
		return ret
	}
	return *o.CertificateBasedAuth
}

// GetCertificateBasedAuthOk returns a tuple with the CertificateBasedAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityConfig) GetCertificateBasedAuthOk() (*SecurityConfigCertificateBasedAuth, bool) {
	if o == nil || o.CertificateBasedAuth == nil {
		return nil, false
	}
	return o.CertificateBasedAuth, true
}

// HasCertificateBasedAuth returns a boolean if a field has been set.
func (o *SecurityConfig) HasCertificateBasedAuth() bool {
	if o != nil && o.CertificateBasedAuth != nil {
		return true
	}

	return false
}

// SetCertificateBasedAuth gets a reference to the given SecurityConfigCertificateBasedAuth and assigns it to the CertificateBasedAuth field.
func (o *SecurityConfig) SetCertificateBasedAuth(v SecurityConfigCertificateBasedAuth) {
	o.CertificateBasedAuth = &v
}

func (o SecurityConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PasswordStrength != nil {
		toSerialize["passwordStrength"] = o.PasswordStrength
	}
	if o.PasswordReuse != nil {
		toSerialize["passwordReuse"] = o.PasswordReuse
	}
	if o.PasswordLifetime != nil {
		toSerialize["passwordLifetime"] = o.PasswordLifetime
	}
	if o.AccountLockout != nil {
		toSerialize["accountLockout"] = o.AccountLockout
	}
	if o.DataClassification != nil {
		toSerialize["dataClassification"] = o.DataClassification
	}
	if o.SessionConfiguration != nil {
		toSerialize["sessionConfiguration"] = o.SessionConfiguration
	}
	if o.CertificateBasedAuth != nil {
		toSerialize["certificateBasedAuth"] = o.CertificateBasedAuth
	}
	return json.Marshal(toSerialize)
}

type NullableSecurityConfig struct {
	value *SecurityConfig
	isSet bool
}

func (v NullableSecurityConfig) Get() *SecurityConfig {
	return v.value
}

func (v *NullableSecurityConfig) Set(val *SecurityConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityConfig(val *SecurityConfig) *NullableSecurityConfig {
	return &NullableSecurityConfig{value: val, isSet: true}
}

func (v NullableSecurityConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o SecurityConfig) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}