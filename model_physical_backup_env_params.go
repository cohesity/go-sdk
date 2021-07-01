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

// PhysicalBackupEnvParams Message to capture any additional backup params for a Physical environment.
type PhysicalBackupEnvParams struct {
	// If this is set to true, then incremental backup will be performed after the server restarts, otherwise a full-backup will be done. NOTE: This is applicable to windows host environments.
	EnableIncrementalBackupAfterRestart NullableBool `json:"enableIncrementalBackupAfterRestart,omitempty"`
	FilteringPolicy *FilteringPolicyProto `json:"filteringPolicy,omitempty"`
}

// NewPhysicalBackupEnvParams instantiates a new PhysicalBackupEnvParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhysicalBackupEnvParams() *PhysicalBackupEnvParams {
	this := PhysicalBackupEnvParams{}
	return &this
}

// NewPhysicalBackupEnvParamsWithDefaults instantiates a new PhysicalBackupEnvParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhysicalBackupEnvParamsWithDefaults() *PhysicalBackupEnvParams {
	this := PhysicalBackupEnvParams{}
	return &this
}

// GetEnableIncrementalBackupAfterRestart returns the EnableIncrementalBackupAfterRestart field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PhysicalBackupEnvParams) GetEnableIncrementalBackupAfterRestart() bool {
	if o == nil || o.EnableIncrementalBackupAfterRestart.Get() == nil {
		var ret bool
		return ret
	}
	return *o.EnableIncrementalBackupAfterRestart.Get()
}

// GetEnableIncrementalBackupAfterRestartOk returns a tuple with the EnableIncrementalBackupAfterRestart field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PhysicalBackupEnvParams) GetEnableIncrementalBackupAfterRestartOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EnableIncrementalBackupAfterRestart.Get(), o.EnableIncrementalBackupAfterRestart.IsSet()
}

// HasEnableIncrementalBackupAfterRestart returns a boolean if a field has been set.
func (o *PhysicalBackupEnvParams) HasEnableIncrementalBackupAfterRestart() bool {
	if o != nil && o.EnableIncrementalBackupAfterRestart.IsSet() {
		return true
	}

	return false
}

// SetEnableIncrementalBackupAfterRestart gets a reference to the given NullableBool and assigns it to the EnableIncrementalBackupAfterRestart field.
func (o *PhysicalBackupEnvParams) SetEnableIncrementalBackupAfterRestart(v bool) {
	o.EnableIncrementalBackupAfterRestart.Set(&v)
}
// SetEnableIncrementalBackupAfterRestartNil sets the value for EnableIncrementalBackupAfterRestart to be an explicit nil
func (o *PhysicalBackupEnvParams) SetEnableIncrementalBackupAfterRestartNil() {
	o.EnableIncrementalBackupAfterRestart.Set(nil)
}

// UnsetEnableIncrementalBackupAfterRestart ensures that no value is present for EnableIncrementalBackupAfterRestart, not even an explicit nil
func (o *PhysicalBackupEnvParams) UnsetEnableIncrementalBackupAfterRestart() {
	o.EnableIncrementalBackupAfterRestart.Unset()
}

// GetFilteringPolicy returns the FilteringPolicy field value if set, zero value otherwise.
func (o *PhysicalBackupEnvParams) GetFilteringPolicy() FilteringPolicyProto {
	if o == nil || o.FilteringPolicy == nil {
		var ret FilteringPolicyProto
		return ret
	}
	return *o.FilteringPolicy
}

// GetFilteringPolicyOk returns a tuple with the FilteringPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalBackupEnvParams) GetFilteringPolicyOk() (*FilteringPolicyProto, bool) {
	if o == nil || o.FilteringPolicy == nil {
		return nil, false
	}
	return o.FilteringPolicy, true
}

// HasFilteringPolicy returns a boolean if a field has been set.
func (o *PhysicalBackupEnvParams) HasFilteringPolicy() bool {
	if o != nil && o.FilteringPolicy != nil {
		return true
	}

	return false
}

// SetFilteringPolicy gets a reference to the given FilteringPolicyProto and assigns it to the FilteringPolicy field.
func (o *PhysicalBackupEnvParams) SetFilteringPolicy(v FilteringPolicyProto) {
	o.FilteringPolicy = &v
}

func (o PhysicalBackupEnvParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EnableIncrementalBackupAfterRestart.IsSet() {
		toSerialize["enableIncrementalBackupAfterRestart"] = o.EnableIncrementalBackupAfterRestart.Get()
	}
	if o.FilteringPolicy != nil {
		toSerialize["filteringPolicy"] = o.FilteringPolicy
	}
	return json.Marshal(toSerialize)
}

type NullablePhysicalBackupEnvParams struct {
	value *PhysicalBackupEnvParams
	isSet bool
}

func (v NullablePhysicalBackupEnvParams) Get() *PhysicalBackupEnvParams {
	return v.value
}

func (v *NullablePhysicalBackupEnvParams) Set(val *PhysicalBackupEnvParams) {
	v.value = val
	v.isSet = true
}

func (v NullablePhysicalBackupEnvParams) IsSet() bool {
	return v.isSet
}

func (v *NullablePhysicalBackupEnvParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhysicalBackupEnvParams(val *PhysicalBackupEnvParams) *NullablePhysicalBackupEnvParams {
	return &NullablePhysicalBackupEnvParams{value: val, isSet: true}
}

func (v NullablePhysicalBackupEnvParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhysicalBackupEnvParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


