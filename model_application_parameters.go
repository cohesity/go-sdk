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

// ApplicationParameters struct for ApplicationParameters
type ApplicationParameters struct {
	// If true, after the Cohesity Cluster successfully captures a Snapshot during a Job Run, the Cluster truncates the Exchange transaction logs on a Microsoft Exchange Server. The default value is false.
	TruncateExchangeLog NullableBool `json:"truncateExchangeLog,omitempty"`
}

// NewApplicationParameters instantiates a new ApplicationParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationParameters() *ApplicationParameters {
	this := ApplicationParameters{}
	return &this
}

// NewApplicationParametersWithDefaults instantiates a new ApplicationParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationParametersWithDefaults() *ApplicationParameters {
	this := ApplicationParameters{}
	return &this
}

// GetTruncateExchangeLog returns the TruncateExchangeLog field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationParameters) GetTruncateExchangeLog() bool {
	if o == nil || o.TruncateExchangeLog.Get() == nil {
		var ret bool
		return ret
	}
	return *o.TruncateExchangeLog.Get()
}

// GetTruncateExchangeLogOk returns a tuple with the TruncateExchangeLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationParameters) GetTruncateExchangeLogOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TruncateExchangeLog.Get(), o.TruncateExchangeLog.IsSet()
}

// HasTruncateExchangeLog returns a boolean if a field has been set.
func (o *ApplicationParameters) HasTruncateExchangeLog() bool {
	if o != nil && o.TruncateExchangeLog.IsSet() {
		return true
	}

	return false
}

// SetTruncateExchangeLog gets a reference to the given NullableBool and assigns it to the TruncateExchangeLog field.
func (o *ApplicationParameters) SetTruncateExchangeLog(v bool) {
	o.TruncateExchangeLog.Set(&v)
}
// SetTruncateExchangeLogNil sets the value for TruncateExchangeLog to be an explicit nil
func (o *ApplicationParameters) SetTruncateExchangeLogNil() {
	o.TruncateExchangeLog.Set(nil)
}

// UnsetTruncateExchangeLog ensures that no value is present for TruncateExchangeLog, not even an explicit nil
func (o *ApplicationParameters) UnsetTruncateExchangeLog() {
	o.TruncateExchangeLog.Unset()
}

func (o ApplicationParameters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TruncateExchangeLog.IsSet() {
		toSerialize["truncateExchangeLog"] = o.TruncateExchangeLog.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationParameters struct {
	value *ApplicationParameters
	isSet bool
}

func (v NullableApplicationParameters) Get() *ApplicationParameters {
	return v.value
}

func (v *NullableApplicationParameters) Set(val *ApplicationParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationParameters(val *ApplicationParameters) *NullableApplicationParameters {
	return &NullableApplicationParameters{value: val, isSet: true}
}

func (v NullableApplicationParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


