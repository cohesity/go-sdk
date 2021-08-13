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

// CreateUpgradeTaskRequest Specifies the params to create a schedule based agent upgrade task.
type CreateUpgradeTaskRequest struct {
	// Specifies the name of the task.
	Name NullableString `json:"name,omitempty"`
	// Specifies the description of the task.
	Description NullableString `json:"description,omitempty"`
	// Specifies agent IDs to be upgraded in the task.
	AgentIDs []int64 `json:"agentIDs,omitempty"`
	// Specifies the start time of the task specified by the user as a Unix epoch Timestamp (in microseconds). If no schedule is specified, the agents will be upgraded immediately.
	ScheduleTimeUsecs NullableInt64 `json:"scheduleTimeUsecs,omitempty"`
	// Specifies the time before which the upgrade task should start execution as a Unix epoch Timestamp (in microseconds). If this is not specified the task will start anytime after scheduleTimeUsecs.
	ScheduleEndTimeUsecs NullableInt64 `json:"scheduleEndTimeUsecs,omitempty"`
	// Specifies the task that needs to be retried.
	RetryTaskId NullableInt64 `json:"retryTaskId,omitempty"`
}

// NewCreateUpgradeTaskRequest instantiates a new CreateUpgradeTaskRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateUpgradeTaskRequest() *CreateUpgradeTaskRequest {
	this := CreateUpgradeTaskRequest{}
	return &this
}

// NewCreateUpgradeTaskRequestWithDefaults instantiates a new CreateUpgradeTaskRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateUpgradeTaskRequestWithDefaults() *CreateUpgradeTaskRequest {
	this := CreateUpgradeTaskRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateUpgradeTaskRequest) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateUpgradeTaskRequest) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CreateUpgradeTaskRequest) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CreateUpgradeTaskRequest) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *CreateUpgradeTaskRequest) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CreateUpgradeTaskRequest) UnsetName() {
	o.Name.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateUpgradeTaskRequest) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateUpgradeTaskRequest) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateUpgradeTaskRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *CreateUpgradeTaskRequest) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *CreateUpgradeTaskRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *CreateUpgradeTaskRequest) UnsetDescription() {
	o.Description.Unset()
}

// GetAgentIDs returns the AgentIDs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateUpgradeTaskRequest) GetAgentIDs() []int64 {
	if o == nil  {
		var ret []int64
		return ret
	}
	return o.AgentIDs
}

// GetAgentIDsOk returns a tuple with the AgentIDs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateUpgradeTaskRequest) GetAgentIDsOk() (*[]int64, bool) {
	if o == nil || o.AgentIDs == nil {
		return nil, false
	}
	return &o.AgentIDs, true
}

// HasAgentIDs returns a boolean if a field has been set.
func (o *CreateUpgradeTaskRequest) HasAgentIDs() bool {
	if o != nil && o.AgentIDs != nil {
		return true
	}

	return false
}

// SetAgentIDs gets a reference to the given []int64 and assigns it to the AgentIDs field.
func (o *CreateUpgradeTaskRequest) SetAgentIDs(v []int64) {
	o.AgentIDs = v
}

// GetScheduleTimeUsecs returns the ScheduleTimeUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateUpgradeTaskRequest) GetScheduleTimeUsecs() int64 {
	if o == nil || o.ScheduleTimeUsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ScheduleTimeUsecs.Get()
}

// GetScheduleTimeUsecsOk returns a tuple with the ScheduleTimeUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateUpgradeTaskRequest) GetScheduleTimeUsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ScheduleTimeUsecs.Get(), o.ScheduleTimeUsecs.IsSet()
}

// HasScheduleTimeUsecs returns a boolean if a field has been set.
func (o *CreateUpgradeTaskRequest) HasScheduleTimeUsecs() bool {
	if o != nil && o.ScheduleTimeUsecs.IsSet() {
		return true
	}

	return false
}

// SetScheduleTimeUsecs gets a reference to the given NullableInt64 and assigns it to the ScheduleTimeUsecs field.
func (o *CreateUpgradeTaskRequest) SetScheduleTimeUsecs(v int64) {
	o.ScheduleTimeUsecs.Set(&v)
}
// SetScheduleTimeUsecsNil sets the value for ScheduleTimeUsecs to be an explicit nil
func (o *CreateUpgradeTaskRequest) SetScheduleTimeUsecsNil() {
	o.ScheduleTimeUsecs.Set(nil)
}

// UnsetScheduleTimeUsecs ensures that no value is present for ScheduleTimeUsecs, not even an explicit nil
func (o *CreateUpgradeTaskRequest) UnsetScheduleTimeUsecs() {
	o.ScheduleTimeUsecs.Unset()
}

// GetScheduleEndTimeUsecs returns the ScheduleEndTimeUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateUpgradeTaskRequest) GetScheduleEndTimeUsecs() int64 {
	if o == nil || o.ScheduleEndTimeUsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ScheduleEndTimeUsecs.Get()
}

// GetScheduleEndTimeUsecsOk returns a tuple with the ScheduleEndTimeUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateUpgradeTaskRequest) GetScheduleEndTimeUsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ScheduleEndTimeUsecs.Get(), o.ScheduleEndTimeUsecs.IsSet()
}

// HasScheduleEndTimeUsecs returns a boolean if a field has been set.
func (o *CreateUpgradeTaskRequest) HasScheduleEndTimeUsecs() bool {
	if o != nil && o.ScheduleEndTimeUsecs.IsSet() {
		return true
	}

	return false
}

// SetScheduleEndTimeUsecs gets a reference to the given NullableInt64 and assigns it to the ScheduleEndTimeUsecs field.
func (o *CreateUpgradeTaskRequest) SetScheduleEndTimeUsecs(v int64) {
	o.ScheduleEndTimeUsecs.Set(&v)
}
// SetScheduleEndTimeUsecsNil sets the value for ScheduleEndTimeUsecs to be an explicit nil
func (o *CreateUpgradeTaskRequest) SetScheduleEndTimeUsecsNil() {
	o.ScheduleEndTimeUsecs.Set(nil)
}

// UnsetScheduleEndTimeUsecs ensures that no value is present for ScheduleEndTimeUsecs, not even an explicit nil
func (o *CreateUpgradeTaskRequest) UnsetScheduleEndTimeUsecs() {
	o.ScheduleEndTimeUsecs.Unset()
}

// GetRetryTaskId returns the RetryTaskId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateUpgradeTaskRequest) GetRetryTaskId() int64 {
	if o == nil || o.RetryTaskId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RetryTaskId.Get()
}

// GetRetryTaskIdOk returns a tuple with the RetryTaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateUpgradeTaskRequest) GetRetryTaskIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RetryTaskId.Get(), o.RetryTaskId.IsSet()
}

// HasRetryTaskId returns a boolean if a field has been set.
func (o *CreateUpgradeTaskRequest) HasRetryTaskId() bool {
	if o != nil && o.RetryTaskId.IsSet() {
		return true
	}

	return false
}

// SetRetryTaskId gets a reference to the given NullableInt64 and assigns it to the RetryTaskId field.
func (o *CreateUpgradeTaskRequest) SetRetryTaskId(v int64) {
	o.RetryTaskId.Set(&v)
}
// SetRetryTaskIdNil sets the value for RetryTaskId to be an explicit nil
func (o *CreateUpgradeTaskRequest) SetRetryTaskIdNil() {
	o.RetryTaskId.Set(nil)
}

// UnsetRetryTaskId ensures that no value is present for RetryTaskId, not even an explicit nil
func (o *CreateUpgradeTaskRequest) UnsetRetryTaskId() {
	o.RetryTaskId.Unset()
}

func (o CreateUpgradeTaskRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.AgentIDs != nil {
		toSerialize["agentIDs"] = o.AgentIDs
	}
	if o.ScheduleTimeUsecs.IsSet() {
		toSerialize["scheduleTimeUsecs"] = o.ScheduleTimeUsecs.Get()
	}
	if o.ScheduleEndTimeUsecs.IsSet() {
		toSerialize["scheduleEndTimeUsecs"] = o.ScheduleEndTimeUsecs.Get()
	}
	if o.RetryTaskId.IsSet() {
		toSerialize["retryTaskId"] = o.RetryTaskId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCreateUpgradeTaskRequest struct {
	value *CreateUpgradeTaskRequest
	isSet bool
}

func (v NullableCreateUpgradeTaskRequest) Get() *CreateUpgradeTaskRequest {
	return v.value
}

func (v *NullableCreateUpgradeTaskRequest) Set(val *CreateUpgradeTaskRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUpgradeTaskRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUpgradeTaskRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUpgradeTaskRequest(val *CreateUpgradeTaskRequest) *NullableCreateUpgradeTaskRequest {
	return &NullableCreateUpgradeTaskRequest{value: val, isSet: true}
}

func (v NullableCreateUpgradeTaskRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUpgradeTaskRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o CreateUpgradeTaskRequest) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}