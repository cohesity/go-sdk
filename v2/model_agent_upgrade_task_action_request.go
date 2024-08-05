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

// checks if the AgentUpgradeTaskActionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentUpgradeTaskActionRequest{}

// AgentUpgradeTaskActionRequest Specifies the params to perform an action on an agent upgrade task.
type AgentUpgradeTaskActionRequest struct {
	// Specifies the action type.<br> 'Cancel' indicates that the upgrade task needs to be cancelled.
	Action NullableString `json:"action,omitempty"`
	// Specifies the ID of the task.
	Id NullableInt64 `json:"id,omitempty"`
}

// NewAgentUpgradeTaskActionRequest instantiates a new AgentUpgradeTaskActionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentUpgradeTaskActionRequest() *AgentUpgradeTaskActionRequest {
	this := AgentUpgradeTaskActionRequest{}
	return &this
}

// NewAgentUpgradeTaskActionRequestWithDefaults instantiates a new AgentUpgradeTaskActionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentUpgradeTaskActionRequestWithDefaults() *AgentUpgradeTaskActionRequest {
	this := AgentUpgradeTaskActionRequest{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgentUpgradeTaskActionRequest) GetAction() string {
	if o == nil || IsNil(o.Action.Get()) {
		var ret string
		return ret
	}
	return *o.Action.Get()
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentUpgradeTaskActionRequest) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Action.Get(), o.Action.IsSet()
}

// HasAction returns a boolean if a field has been set.
func (o *AgentUpgradeTaskActionRequest) HasAction() bool {
	if o != nil && o.Action.IsSet() {
		return true
	}

	return false
}

// SetAction gets a reference to the given NullableString and assigns it to the Action field.
func (o *AgentUpgradeTaskActionRequest) SetAction(v string) {
	o.Action.Set(&v)
}
// SetActionNil sets the value for Action to be an explicit nil
func (o *AgentUpgradeTaskActionRequest) SetActionNil() {
	o.Action.Set(nil)
}

// UnsetAction ensures that no value is present for Action, not even an explicit nil
func (o *AgentUpgradeTaskActionRequest) UnsetAction() {
	o.Action.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgentUpgradeTaskActionRequest) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentUpgradeTaskActionRequest) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *AgentUpgradeTaskActionRequest) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *AgentUpgradeTaskActionRequest) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *AgentUpgradeTaskActionRequest) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *AgentUpgradeTaskActionRequest) UnsetId() {
	o.Id.Unset()
}

func (o AgentUpgradeTaskActionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentUpgradeTaskActionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Action.IsSet() {
		toSerialize["action"] = o.Action.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	return toSerialize, nil
}

type NullableAgentUpgradeTaskActionRequest struct {
	value *AgentUpgradeTaskActionRequest
	isSet bool
}

func (v NullableAgentUpgradeTaskActionRequest) Get() *AgentUpgradeTaskActionRequest {
	return v.value
}

func (v *NullableAgentUpgradeTaskActionRequest) Set(val *AgentUpgradeTaskActionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentUpgradeTaskActionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentUpgradeTaskActionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentUpgradeTaskActionRequest(val *AgentUpgradeTaskActionRequest) *NullableAgentUpgradeTaskActionRequest {
	return &NullableAgentUpgradeTaskActionRequest{value: val, isSet: true}
}

func (v NullableAgentUpgradeTaskActionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentUpgradeTaskActionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


