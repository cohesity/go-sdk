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

// AgentInfoObject Specifies the upgrade state of the agent.
type AgentInfoObject struct {
	// Specifies the name of the source where the agent is installed.
	Name NullableString `json:"name,omitempty"`
	// Specifies the software version of the agent before upgrade.
	PreviousSoftwareVersion NullableString `json:"previousSoftwareVersion,omitempty"`
	// Specifies the upgrade status of the agent.<br> 'Scheduled' indicates that upgrade for the agent is yet to start.<br> 'Started' indicates that upgrade for the agent is started.<br> 'Succeeded' indicates that agent was upgraded successfully.<br> 'Failed' indicates that upgrade for the agent has failed.<br> 'Skipped' indicates that upgrade for the agent was skipped.
	Status NullableString `json:"status,omitempty"`
	// Specifies the time when the upgrade for an agent started as a Unix epoch Timestamp (in microseconds).
	StartTimeUsecs NullableInt64 `json:"startTimeUsecs,omitempty"`
	// Specifies the time when the upgrade for an agent completed as a Unix epoch Timestamp (in microseconds).
	EndTimeUsecs NullableInt64 `json:"endTimeUsecs,omitempty"`
	Error *Error `json:"error,omitempty"`
}

// NewAgentInfoObject instantiates a new AgentInfoObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentInfoObject() *AgentInfoObject {
	this := AgentInfoObject{}
	return &this
}

// NewAgentInfoObjectWithDefaults instantiates a new AgentInfoObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentInfoObjectWithDefaults() *AgentInfoObject {
	this := AgentInfoObject{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgentInfoObject) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentInfoObject) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *AgentInfoObject) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *AgentInfoObject) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *AgentInfoObject) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *AgentInfoObject) UnsetName() {
	o.Name.Unset()
}

// GetPreviousSoftwareVersion returns the PreviousSoftwareVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgentInfoObject) GetPreviousSoftwareVersion() string {
	if o == nil || o.PreviousSoftwareVersion.Get() == nil {
		var ret string
		return ret
	}
	return *o.PreviousSoftwareVersion.Get()
}

// GetPreviousSoftwareVersionOk returns a tuple with the PreviousSoftwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentInfoObject) GetPreviousSoftwareVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PreviousSoftwareVersion.Get(), o.PreviousSoftwareVersion.IsSet()
}

// HasPreviousSoftwareVersion returns a boolean if a field has been set.
func (o *AgentInfoObject) HasPreviousSoftwareVersion() bool {
	if o != nil && o.PreviousSoftwareVersion.IsSet() {
		return true
	}

	return false
}

// SetPreviousSoftwareVersion gets a reference to the given NullableString and assigns it to the PreviousSoftwareVersion field.
func (o *AgentInfoObject) SetPreviousSoftwareVersion(v string) {
	o.PreviousSoftwareVersion.Set(&v)
}
// SetPreviousSoftwareVersionNil sets the value for PreviousSoftwareVersion to be an explicit nil
func (o *AgentInfoObject) SetPreviousSoftwareVersionNil() {
	o.PreviousSoftwareVersion.Set(nil)
}

// UnsetPreviousSoftwareVersion ensures that no value is present for PreviousSoftwareVersion, not even an explicit nil
func (o *AgentInfoObject) UnsetPreviousSoftwareVersion() {
	o.PreviousSoftwareVersion.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgentInfoObject) GetStatus() string {
	if o == nil || o.Status.Get() == nil {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentInfoObject) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *AgentInfoObject) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *AgentInfoObject) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *AgentInfoObject) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *AgentInfoObject) UnsetStatus() {
	o.Status.Unset()
}

// GetStartTimeUsecs returns the StartTimeUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgentInfoObject) GetStartTimeUsecs() int64 {
	if o == nil || o.StartTimeUsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.StartTimeUsecs.Get()
}

// GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentInfoObject) GetStartTimeUsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartTimeUsecs.Get(), o.StartTimeUsecs.IsSet()
}

// HasStartTimeUsecs returns a boolean if a field has been set.
func (o *AgentInfoObject) HasStartTimeUsecs() bool {
	if o != nil && o.StartTimeUsecs.IsSet() {
		return true
	}

	return false
}

// SetStartTimeUsecs gets a reference to the given NullableInt64 and assigns it to the StartTimeUsecs field.
func (o *AgentInfoObject) SetStartTimeUsecs(v int64) {
	o.StartTimeUsecs.Set(&v)
}
// SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil
func (o *AgentInfoObject) SetStartTimeUsecsNil() {
	o.StartTimeUsecs.Set(nil)
}

// UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
func (o *AgentInfoObject) UnsetStartTimeUsecs() {
	o.StartTimeUsecs.Unset()
}

// GetEndTimeUsecs returns the EndTimeUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgentInfoObject) GetEndTimeUsecs() int64 {
	if o == nil || o.EndTimeUsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.EndTimeUsecs.Get()
}

// GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentInfoObject) GetEndTimeUsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndTimeUsecs.Get(), o.EndTimeUsecs.IsSet()
}

// HasEndTimeUsecs returns a boolean if a field has been set.
func (o *AgentInfoObject) HasEndTimeUsecs() bool {
	if o != nil && o.EndTimeUsecs.IsSet() {
		return true
	}

	return false
}

// SetEndTimeUsecs gets a reference to the given NullableInt64 and assigns it to the EndTimeUsecs field.
func (o *AgentInfoObject) SetEndTimeUsecs(v int64) {
	o.EndTimeUsecs.Set(&v)
}
// SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil
func (o *AgentInfoObject) SetEndTimeUsecsNil() {
	o.EndTimeUsecs.Set(nil)
}

// UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
func (o *AgentInfoObject) UnsetEndTimeUsecs() {
	o.EndTimeUsecs.Unset()
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *AgentInfoObject) GetError() Error {
	if o == nil || o.Error == nil {
		var ret Error
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentInfoObject) GetErrorOk() (*Error, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *AgentInfoObject) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given Error and assigns it to the Error field.
func (o *AgentInfoObject) SetError(v Error) {
	o.Error = &v
}

func (o AgentInfoObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.PreviousSoftwareVersion.IsSet() {
		toSerialize["previousSoftwareVersion"] = o.PreviousSoftwareVersion.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.StartTimeUsecs.IsSet() {
		toSerialize["startTimeUsecs"] = o.StartTimeUsecs.Get()
	}
	if o.EndTimeUsecs.IsSet() {
		toSerialize["endTimeUsecs"] = o.EndTimeUsecs.Get()
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableAgentInfoObject struct {
	value *AgentInfoObject
	isSet bool
}

func (v NullableAgentInfoObject) Get() *AgentInfoObject {
	return v.value
}

func (v *NullableAgentInfoObject) Set(val *AgentInfoObject) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentInfoObject) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentInfoObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentInfoObject(val *AgentInfoObject) *NullableAgentInfoObject {
	return &NullableAgentInfoObject{value: val, isSet: true}
}

func (v NullableAgentInfoObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentInfoObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o AgentInfoObject) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}