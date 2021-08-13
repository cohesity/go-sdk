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

// ViewUserQuotasConfig Specifies the View user quotas config.
type ViewUserQuotasConfig struct {
	// Specifies whether user quota is enabled for the View.
	Enabled NullableBool `json:"enabled"`
	// Specifies the default user quota policy.
	DefaultQuotaPolicy *QuotaPolicy `json:"defaultQuotaPolicy,omitempty"`
}

// NewViewUserQuotasConfig instantiates a new ViewUserQuotasConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViewUserQuotasConfig(enabled NullableBool) *ViewUserQuotasConfig {
	this := ViewUserQuotasConfig{}
	this.Enabled = enabled
	return &this
}

// NewViewUserQuotasConfigWithDefaults instantiates a new ViewUserQuotasConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewUserQuotasConfigWithDefaults() *ViewUserQuotasConfig {
	this := ViewUserQuotasConfig{}
	return &this
}

// GetEnabled returns the Enabled field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *ViewUserQuotasConfig) GetEnabled() bool {
	if o == nil || o.Enabled.Get() == nil {
		var ret bool
		return ret
	}

	return *o.Enabled.Get()
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ViewUserQuotasConfig) GetEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Enabled.Get(), o.Enabled.IsSet()
}

// SetEnabled sets field value
func (o *ViewUserQuotasConfig) SetEnabled(v bool) {
	o.Enabled.Set(&v)
}

// GetDefaultQuotaPolicy returns the DefaultQuotaPolicy field value if set, zero value otherwise.
func (o *ViewUserQuotasConfig) GetDefaultQuotaPolicy() QuotaPolicy {
	if o == nil || o.DefaultQuotaPolicy == nil {
		var ret QuotaPolicy
		return ret
	}
	return *o.DefaultQuotaPolicy
}

// GetDefaultQuotaPolicyOk returns a tuple with the DefaultQuotaPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewUserQuotasConfig) GetDefaultQuotaPolicyOk() (*QuotaPolicy, bool) {
	if o == nil || o.DefaultQuotaPolicy == nil {
		return nil, false
	}
	return o.DefaultQuotaPolicy, true
}

// HasDefaultQuotaPolicy returns a boolean if a field has been set.
func (o *ViewUserQuotasConfig) HasDefaultQuotaPolicy() bool {
	if o != nil && o.DefaultQuotaPolicy != nil {
		return true
	}

	return false
}

// SetDefaultQuotaPolicy gets a reference to the given QuotaPolicy and assigns it to the DefaultQuotaPolicy field.
func (o *ViewUserQuotasConfig) SetDefaultQuotaPolicy(v QuotaPolicy) {
	o.DefaultQuotaPolicy = &v
}

func (o ViewUserQuotasConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled.Get()
	}
	if o.DefaultQuotaPolicy != nil {
		toSerialize["defaultQuotaPolicy"] = o.DefaultQuotaPolicy
	}
	return json.Marshal(toSerialize)
}

type NullableViewUserQuotasConfig struct {
	value *ViewUserQuotasConfig
	isSet bool
}

func (v NullableViewUserQuotasConfig) Get() *ViewUserQuotasConfig {
	return v.value
}

func (v *NullableViewUserQuotasConfig) Set(val *ViewUserQuotasConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableViewUserQuotasConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableViewUserQuotasConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViewUserQuotasConfig(val *ViewUserQuotasConfig) *NullableViewUserQuotasConfig {
	return &NullableViewUserQuotasConfig{value: val, isSet: true}
}

func (v NullableViewUserQuotasConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViewUserQuotasConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o ViewUserQuotasConfig) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}