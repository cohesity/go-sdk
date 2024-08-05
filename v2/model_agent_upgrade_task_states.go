/*
Cohesity REST API

Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
)

// checks if the AgentUpgradeTaskStates type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentUpgradeTaskStates{}

// AgentUpgradeTaskStates List of agent upgrade tasks.
type AgentUpgradeTaskStates struct {
	// Specifies the list of agent upgrade tasks.
	Tasks []AgentUpgradeTaskState `json:"tasks,omitempty"`
}

// NewAgentUpgradeTaskStates instantiates a new AgentUpgradeTaskStates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentUpgradeTaskStates() *AgentUpgradeTaskStates {
	this := AgentUpgradeTaskStates{}
	return &this
}

// NewAgentUpgradeTaskStatesWithDefaults instantiates a new AgentUpgradeTaskStates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentUpgradeTaskStatesWithDefaults() *AgentUpgradeTaskStates {
	this := AgentUpgradeTaskStates{}
	return &this
}

// GetTasks returns the Tasks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgentUpgradeTaskStates) GetTasks() []AgentUpgradeTaskState {
	if o == nil {
		var ret []AgentUpgradeTaskState
		return ret
	}
	return o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentUpgradeTaskStates) GetTasksOk() ([]AgentUpgradeTaskState, bool) {
	if o == nil || IsNil(o.Tasks) {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *AgentUpgradeTaskStates) HasTasks() bool {
	if o != nil && !IsNil(o.Tasks) {
		return true
	}

	return false
}

// SetTasks gets a reference to the given []AgentUpgradeTaskState and assigns it to the Tasks field.
func (o *AgentUpgradeTaskStates) SetTasks(v []AgentUpgradeTaskState) {
	o.Tasks = v
}

func (o AgentUpgradeTaskStates) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentUpgradeTaskStates) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Tasks != nil {
		toSerialize["tasks"] = o.Tasks
	}
	return toSerialize, nil
}

type NullableAgentUpgradeTaskStates struct {
	value *AgentUpgradeTaskStates
	isSet bool
}

func (v NullableAgentUpgradeTaskStates) Get() *AgentUpgradeTaskStates {
	return v.value
}

func (v *NullableAgentUpgradeTaskStates) Set(val *AgentUpgradeTaskStates) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentUpgradeTaskStates) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentUpgradeTaskStates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentUpgradeTaskStates(val *AgentUpgradeTaskStates) *NullableAgentUpgradeTaskStates {
	return &NullableAgentUpgradeTaskStates{value: val, isSet: true}
}

func (v NullableAgentUpgradeTaskStates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentUpgradeTaskStates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


