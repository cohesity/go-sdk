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

// PolicyConfig Specifies the policy configuration for protecting an object.
type PolicyConfig struct {
	// Specifies the list of protection policies for protecting an object with multiple policies.
	Policies []ObjectPolicy `json:"policies,omitempty"`
}

// NewPolicyConfig instantiates a new PolicyConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyConfig() *PolicyConfig {
	this := PolicyConfig{}
	return &this
}

// NewPolicyConfigWithDefaults instantiates a new PolicyConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyConfigWithDefaults() *PolicyConfig {
	this := PolicyConfig{}
	return &this
}

// GetPolicies returns the Policies field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyConfig) GetPolicies() []ObjectPolicy {
	if o == nil  {
		var ret []ObjectPolicy
		return ret
	}
	return o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyConfig) GetPoliciesOk() (*[]ObjectPolicy, bool) {
	if o == nil || o.Policies == nil {
		return nil, false
	}
	return &o.Policies, true
}

// HasPolicies returns a boolean if a field has been set.
func (o *PolicyConfig) HasPolicies() bool {
	if o != nil && o.Policies != nil {
		return true
	}

	return false
}

// SetPolicies gets a reference to the given []ObjectPolicy and assigns it to the Policies field.
func (o *PolicyConfig) SetPolicies(v []ObjectPolicy) {
	o.Policies = v
}

func (o PolicyConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Policies != nil {
		toSerialize["policies"] = o.Policies
	}
	return json.Marshal(toSerialize)
}

type NullablePolicyConfig struct {
	value *PolicyConfig
	isSet bool
}

func (v NullablePolicyConfig) Get() *PolicyConfig {
	return v.value
}

func (v *NullablePolicyConfig) Set(val *PolicyConfig) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyConfig) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyConfig(val *PolicyConfig) *NullablePolicyConfig {
	return &NullablePolicyConfig{value: val, isSet: true}
}

func (v NullablePolicyConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


