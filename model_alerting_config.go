/*
 * Cohesity REST API
 *
 * This API list provides operations for interfacing with the Cohesity Cluster.
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cohesitysdk

import (
	"encoding/json"
)

// AlertingConfig Specifies optional settings for alerting.
type AlertingConfig struct {
	// Exists to maintain backwards compatibility with versions before eff8198.
	EmailAddresses []string `json:"emailAddresses,omitempty"`
	// Specifies additional email addresses where alert notifications (configured in the AlertingPolicy) must be sent.
	EmailDeliveryTargets []EmailDeliveryTarget `json:"emailDeliveryTargets,omitempty"`
	// Specifies the boolean to raise per object alert for failures.
	RaiseObjectLevelFailureAlert NullableBool `json:"raiseObjectLevelFailureAlert,omitempty"`
}

// NewAlertingConfig instantiates a new AlertingConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertingConfig() *AlertingConfig {
	this := AlertingConfig{}
	return &this
}

// NewAlertingConfigWithDefaults instantiates a new AlertingConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertingConfigWithDefaults() *AlertingConfig {
	this := AlertingConfig{}
	return &this
}

// GetEmailAddresses returns the EmailAddresses field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlertingConfig) GetEmailAddresses() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.EmailAddresses
}

// GetEmailAddressesOk returns a tuple with the EmailAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlertingConfig) GetEmailAddressesOk() (*[]string, bool) {
	if o == nil || o.EmailAddresses == nil {
		return nil, false
	}
	return &o.EmailAddresses, true
}

// HasEmailAddresses returns a boolean if a field has been set.
func (o *AlertingConfig) HasEmailAddresses() bool {
	if o != nil && o.EmailAddresses != nil {
		return true
	}

	return false
}

// SetEmailAddresses gets a reference to the given []string and assigns it to the EmailAddresses field.
func (o *AlertingConfig) SetEmailAddresses(v []string) {
	o.EmailAddresses = v
}

// GetEmailDeliveryTargets returns the EmailDeliveryTargets field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlertingConfig) GetEmailDeliveryTargets() []EmailDeliveryTarget {
	if o == nil  {
		var ret []EmailDeliveryTarget
		return ret
	}
	return o.EmailDeliveryTargets
}

// GetEmailDeliveryTargetsOk returns a tuple with the EmailDeliveryTargets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlertingConfig) GetEmailDeliveryTargetsOk() (*[]EmailDeliveryTarget, bool) {
	if o == nil || o.EmailDeliveryTargets == nil {
		return nil, false
	}
	return &o.EmailDeliveryTargets, true
}

// HasEmailDeliveryTargets returns a boolean if a field has been set.
func (o *AlertingConfig) HasEmailDeliveryTargets() bool {
	if o != nil && o.EmailDeliveryTargets != nil {
		return true
	}

	return false
}

// SetEmailDeliveryTargets gets a reference to the given []EmailDeliveryTarget and assigns it to the EmailDeliveryTargets field.
func (o *AlertingConfig) SetEmailDeliveryTargets(v []EmailDeliveryTarget) {
	o.EmailDeliveryTargets = v
}

// GetRaiseObjectLevelFailureAlert returns the RaiseObjectLevelFailureAlert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlertingConfig) GetRaiseObjectLevelFailureAlert() bool {
	if o == nil || o.RaiseObjectLevelFailureAlert.Get() == nil {
		var ret bool
		return ret
	}
	return *o.RaiseObjectLevelFailureAlert.Get()
}

// GetRaiseObjectLevelFailureAlertOk returns a tuple with the RaiseObjectLevelFailureAlert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlertingConfig) GetRaiseObjectLevelFailureAlertOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RaiseObjectLevelFailureAlert.Get(), o.RaiseObjectLevelFailureAlert.IsSet()
}

// HasRaiseObjectLevelFailureAlert returns a boolean if a field has been set.
func (o *AlertingConfig) HasRaiseObjectLevelFailureAlert() bool {
	if o != nil && o.RaiseObjectLevelFailureAlert.IsSet() {
		return true
	}

	return false
}

// SetRaiseObjectLevelFailureAlert gets a reference to the given NullableBool and assigns it to the RaiseObjectLevelFailureAlert field.
func (o *AlertingConfig) SetRaiseObjectLevelFailureAlert(v bool) {
	o.RaiseObjectLevelFailureAlert.Set(&v)
}
// SetRaiseObjectLevelFailureAlertNil sets the value for RaiseObjectLevelFailureAlert to be an explicit nil
func (o *AlertingConfig) SetRaiseObjectLevelFailureAlertNil() {
	o.RaiseObjectLevelFailureAlert.Set(nil)
}

// UnsetRaiseObjectLevelFailureAlert ensures that no value is present for RaiseObjectLevelFailureAlert, not even an explicit nil
func (o *AlertingConfig) UnsetRaiseObjectLevelFailureAlert() {
	o.RaiseObjectLevelFailureAlert.Unset()
}

func (o AlertingConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EmailAddresses != nil {
		toSerialize["emailAddresses"] = o.EmailAddresses
	}
	if o.EmailDeliveryTargets != nil {
		toSerialize["emailDeliveryTargets"] = o.EmailDeliveryTargets
	}
	if o.RaiseObjectLevelFailureAlert.IsSet() {
		toSerialize["raiseObjectLevelFailureAlert"] = o.RaiseObjectLevelFailureAlert.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAlertingConfig struct {
	value *AlertingConfig
	isSet bool
}

func (v NullableAlertingConfig) Get() *AlertingConfig {
	return v.value
}

func (v *NullableAlertingConfig) Set(val *AlertingConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertingConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertingConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertingConfig(val *AlertingConfig) *NullableAlertingConfig {
	return &NullableAlertingConfig{value: val, isSet: true}
}

func (v NullableAlertingConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertingConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


