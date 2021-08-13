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

// UserUiConfig Specifies the params to update UI config.
type UserUiConfig struct {
	// Specifies the user's preferences for UI customization.
	Preferences NullableString `json:"preferences,omitempty"`
	// Specifies the locale.
	Locale NullableString `json:"locale,omitempty"`
}

// NewUserUiConfig instantiates a new UserUiConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserUiConfig() *UserUiConfig {
	this := UserUiConfig{}
	return &this
}

// NewUserUiConfigWithDefaults instantiates a new UserUiConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserUiConfigWithDefaults() *UserUiConfig {
	this := UserUiConfig{}
	return &this
}

// GetPreferences returns the Preferences field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserUiConfig) GetPreferences() string {
	if o == nil || o.Preferences.Get() == nil {
		var ret string
		return ret
	}
	return *o.Preferences.Get()
}

// GetPreferencesOk returns a tuple with the Preferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserUiConfig) GetPreferencesOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Preferences.Get(), o.Preferences.IsSet()
}

// HasPreferences returns a boolean if a field has been set.
func (o *UserUiConfig) HasPreferences() bool {
	if o != nil && o.Preferences.IsSet() {
		return true
	}

	return false
}

// SetPreferences gets a reference to the given NullableString and assigns it to the Preferences field.
func (o *UserUiConfig) SetPreferences(v string) {
	o.Preferences.Set(&v)
}
// SetPreferencesNil sets the value for Preferences to be an explicit nil
func (o *UserUiConfig) SetPreferencesNil() {
	o.Preferences.Set(nil)
}

// UnsetPreferences ensures that no value is present for Preferences, not even an explicit nil
func (o *UserUiConfig) UnsetPreferences() {
	o.Preferences.Unset()
}

// GetLocale returns the Locale field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserUiConfig) GetLocale() string {
	if o == nil || o.Locale.Get() == nil {
		var ret string
		return ret
	}
	return *o.Locale.Get()
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserUiConfig) GetLocaleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Locale.Get(), o.Locale.IsSet()
}

// HasLocale returns a boolean if a field has been set.
func (o *UserUiConfig) HasLocale() bool {
	if o != nil && o.Locale.IsSet() {
		return true
	}

	return false
}

// SetLocale gets a reference to the given NullableString and assigns it to the Locale field.
func (o *UserUiConfig) SetLocale(v string) {
	o.Locale.Set(&v)
}
// SetLocaleNil sets the value for Locale to be an explicit nil
func (o *UserUiConfig) SetLocaleNil() {
	o.Locale.Set(nil)
}

// UnsetLocale ensures that no value is present for Locale, not even an explicit nil
func (o *UserUiConfig) UnsetLocale() {
	o.Locale.Unset()
}

func (o UserUiConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Preferences.IsSet() {
		toSerialize["preferences"] = o.Preferences.Get()
	}
	if o.Locale.IsSet() {
		toSerialize["locale"] = o.Locale.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableUserUiConfig struct {
	value *UserUiConfig
	isSet bool
}

func (v NullableUserUiConfig) Get() *UserUiConfig {
	return v.value
}

func (v *NullableUserUiConfig) Set(val *UserUiConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableUserUiConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableUserUiConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserUiConfig(val *UserUiConfig) *NullableUserUiConfig {
	return &NullableUserUiConfig{value: val, isSet: true}
}

func (v NullableUserUiConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserUiConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o UserUiConfig) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}