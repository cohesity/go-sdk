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

// AgentInformation Specifies the agent details.
type AgentInformation struct {
	// Specifies the status of agent connection.
	ConnectionStatus NullableString `json:"connectionStatus,omitempty"`
	// Specifies the whether agent version is compatible with cluster version ro use various features.
	SupportStatus NullableString `json:"supportStatus,omitempty"`
	// Specifies the software version of the agent
	AgentSwVersion NullableString `json:"agentSwVersion,omitempty"`
	// Specifies the time in usecs when the last agent info was fetched.
	LastFetchedTimeInUsecs NullableInt64 `json:"lastFetchedTimeInUsecs,omitempty"`
	// Specifies the list of host checks and its results.
	HostSettingChecks []HostSettingCheck `json:"hostSettingChecks,omitempty"`
}

// NewAgentInformation instantiates a new AgentInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentInformation() *AgentInformation {
	this := AgentInformation{}
	return &this
}

// NewAgentInformationWithDefaults instantiates a new AgentInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentInformationWithDefaults() *AgentInformation {
	this := AgentInformation{}
	return &this
}

// GetConnectionStatus returns the ConnectionStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgentInformation) GetConnectionStatus() string {
	if o == nil || o.ConnectionStatus.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConnectionStatus.Get()
}

// GetConnectionStatusOk returns a tuple with the ConnectionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentInformation) GetConnectionStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ConnectionStatus.Get(), o.ConnectionStatus.IsSet()
}

// HasConnectionStatus returns a boolean if a field has been set.
func (o *AgentInformation) HasConnectionStatus() bool {
	if o != nil && o.ConnectionStatus.IsSet() {
		return true
	}

	return false
}

// SetConnectionStatus gets a reference to the given NullableString and assigns it to the ConnectionStatus field.
func (o *AgentInformation) SetConnectionStatus(v string) {
	o.ConnectionStatus.Set(&v)
}
// SetConnectionStatusNil sets the value for ConnectionStatus to be an explicit nil
func (o *AgentInformation) SetConnectionStatusNil() {
	o.ConnectionStatus.Set(nil)
}

// UnsetConnectionStatus ensures that no value is present for ConnectionStatus, not even an explicit nil
func (o *AgentInformation) UnsetConnectionStatus() {
	o.ConnectionStatus.Unset()
}

// GetSupportStatus returns the SupportStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgentInformation) GetSupportStatus() string {
	if o == nil || o.SupportStatus.Get() == nil {
		var ret string
		return ret
	}
	return *o.SupportStatus.Get()
}

// GetSupportStatusOk returns a tuple with the SupportStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentInformation) GetSupportStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SupportStatus.Get(), o.SupportStatus.IsSet()
}

// HasSupportStatus returns a boolean if a field has been set.
func (o *AgentInformation) HasSupportStatus() bool {
	if o != nil && o.SupportStatus.IsSet() {
		return true
	}

	return false
}

// SetSupportStatus gets a reference to the given NullableString and assigns it to the SupportStatus field.
func (o *AgentInformation) SetSupportStatus(v string) {
	o.SupportStatus.Set(&v)
}
// SetSupportStatusNil sets the value for SupportStatus to be an explicit nil
func (o *AgentInformation) SetSupportStatusNil() {
	o.SupportStatus.Set(nil)
}

// UnsetSupportStatus ensures that no value is present for SupportStatus, not even an explicit nil
func (o *AgentInformation) UnsetSupportStatus() {
	o.SupportStatus.Unset()
}

// GetAgentSwVersion returns the AgentSwVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgentInformation) GetAgentSwVersion() string {
	if o == nil || o.AgentSwVersion.Get() == nil {
		var ret string
		return ret
	}
	return *o.AgentSwVersion.Get()
}

// GetAgentSwVersionOk returns a tuple with the AgentSwVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentInformation) GetAgentSwVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AgentSwVersion.Get(), o.AgentSwVersion.IsSet()
}

// HasAgentSwVersion returns a boolean if a field has been set.
func (o *AgentInformation) HasAgentSwVersion() bool {
	if o != nil && o.AgentSwVersion.IsSet() {
		return true
	}

	return false
}

// SetAgentSwVersion gets a reference to the given NullableString and assigns it to the AgentSwVersion field.
func (o *AgentInformation) SetAgentSwVersion(v string) {
	o.AgentSwVersion.Set(&v)
}
// SetAgentSwVersionNil sets the value for AgentSwVersion to be an explicit nil
func (o *AgentInformation) SetAgentSwVersionNil() {
	o.AgentSwVersion.Set(nil)
}

// UnsetAgentSwVersion ensures that no value is present for AgentSwVersion, not even an explicit nil
func (o *AgentInformation) UnsetAgentSwVersion() {
	o.AgentSwVersion.Unset()
}

// GetLastFetchedTimeInUsecs returns the LastFetchedTimeInUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgentInformation) GetLastFetchedTimeInUsecs() int64 {
	if o == nil || o.LastFetchedTimeInUsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.LastFetchedTimeInUsecs.Get()
}

// GetLastFetchedTimeInUsecsOk returns a tuple with the LastFetchedTimeInUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentInformation) GetLastFetchedTimeInUsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastFetchedTimeInUsecs.Get(), o.LastFetchedTimeInUsecs.IsSet()
}

// HasLastFetchedTimeInUsecs returns a boolean if a field has been set.
func (o *AgentInformation) HasLastFetchedTimeInUsecs() bool {
	if o != nil && o.LastFetchedTimeInUsecs.IsSet() {
		return true
	}

	return false
}

// SetLastFetchedTimeInUsecs gets a reference to the given NullableInt64 and assigns it to the LastFetchedTimeInUsecs field.
func (o *AgentInformation) SetLastFetchedTimeInUsecs(v int64) {
	o.LastFetchedTimeInUsecs.Set(&v)
}
// SetLastFetchedTimeInUsecsNil sets the value for LastFetchedTimeInUsecs to be an explicit nil
func (o *AgentInformation) SetLastFetchedTimeInUsecsNil() {
	o.LastFetchedTimeInUsecs.Set(nil)
}

// UnsetLastFetchedTimeInUsecs ensures that no value is present for LastFetchedTimeInUsecs, not even an explicit nil
func (o *AgentInformation) UnsetLastFetchedTimeInUsecs() {
	o.LastFetchedTimeInUsecs.Unset()
}

// GetHostSettingChecks returns the HostSettingChecks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgentInformation) GetHostSettingChecks() []HostSettingCheck {
	if o == nil  {
		var ret []HostSettingCheck
		return ret
	}
	return o.HostSettingChecks
}

// GetHostSettingChecksOk returns a tuple with the HostSettingChecks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentInformation) GetHostSettingChecksOk() (*[]HostSettingCheck, bool) {
	if o == nil || o.HostSettingChecks == nil {
		return nil, false
	}
	return &o.HostSettingChecks, true
}

// HasHostSettingChecks returns a boolean if a field has been set.
func (o *AgentInformation) HasHostSettingChecks() bool {
	if o != nil && o.HostSettingChecks != nil {
		return true
	}

	return false
}

// SetHostSettingChecks gets a reference to the given []HostSettingCheck and assigns it to the HostSettingChecks field.
func (o *AgentInformation) SetHostSettingChecks(v []HostSettingCheck) {
	o.HostSettingChecks = v
}

func (o AgentInformation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ConnectionStatus.IsSet() {
		toSerialize["connectionStatus"] = o.ConnectionStatus.Get()
	}
	if o.SupportStatus.IsSet() {
		toSerialize["supportStatus"] = o.SupportStatus.Get()
	}
	if o.AgentSwVersion.IsSet() {
		toSerialize["agentSwVersion"] = o.AgentSwVersion.Get()
	}
	if o.LastFetchedTimeInUsecs.IsSet() {
		toSerialize["lastFetchedTimeInUsecs"] = o.LastFetchedTimeInUsecs.Get()
	}
	if o.HostSettingChecks != nil {
		toSerialize["hostSettingChecks"] = o.HostSettingChecks
	}
	return json.Marshal(toSerialize)
}

type NullableAgentInformation struct {
	value *AgentInformation
	isSet bool
}

func (v NullableAgentInformation) Get() *AgentInformation {
	return v.value
}

func (v *NullableAgentInformation) Set(val *AgentInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentInformation(val *AgentInformation) *NullableAgentInformation {
	return &NullableAgentInformation{value: val, isSet: true}
}

func (v NullableAgentInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o AgentInformation) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}