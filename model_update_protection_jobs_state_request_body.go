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

// UpdateProtectionJobsStateRequestBody Specifies the parameters to perform an action of list of Protection Jobs.
type UpdateProtectionJobsStateRequestBody struct {
	// Specifies the action to be performed on all the specified Protection Jobs. Specifies the type of action to be performed on Protection Job. 'kActivate' specifies that Protection Job should be activated. 'kDeactivate' specifies that Protection Job should be deactivated. 'kPause' specifies that Protection Job should be paused. 'kResume' specifies that Protection Job should be resumed.
	Action NullableString `json:"action,omitempty"`
	// Specifies a list of Protection Job ids for which the state should change.
	JobIds []int64 `json:"jobIds,omitempty"`
}

// NewUpdateProtectionJobsStateRequestBody instantiates a new UpdateProtectionJobsStateRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateProtectionJobsStateRequestBody() *UpdateProtectionJobsStateRequestBody {
	this := UpdateProtectionJobsStateRequestBody{}
	return &this
}

// NewUpdateProtectionJobsStateRequestBodyWithDefaults instantiates a new UpdateProtectionJobsStateRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateProtectionJobsStateRequestBodyWithDefaults() *UpdateProtectionJobsStateRequestBody {
	this := UpdateProtectionJobsStateRequestBody{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateProtectionJobsStateRequestBody) GetAction() string {
	if o == nil || o.Action.Get() == nil {
		var ret string
		return ret
	}
	return *o.Action.Get()
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateProtectionJobsStateRequestBody) GetActionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Action.Get(), o.Action.IsSet()
}

// HasAction returns a boolean if a field has been set.
func (o *UpdateProtectionJobsStateRequestBody) HasAction() bool {
	if o != nil && o.Action.IsSet() {
		return true
	}

	return false
}

// SetAction gets a reference to the given NullableString and assigns it to the Action field.
func (o *UpdateProtectionJobsStateRequestBody) SetAction(v string) {
	o.Action.Set(&v)
}
// SetActionNil sets the value for Action to be an explicit nil
func (o *UpdateProtectionJobsStateRequestBody) SetActionNil() {
	o.Action.Set(nil)
}

// UnsetAction ensures that no value is present for Action, not even an explicit nil
func (o *UpdateProtectionJobsStateRequestBody) UnsetAction() {
	o.Action.Unset()
}

// GetJobIds returns the JobIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateProtectionJobsStateRequestBody) GetJobIds() []int64 {
	if o == nil  {
		var ret []int64
		return ret
	}
	return o.JobIds
}

// GetJobIdsOk returns a tuple with the JobIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateProtectionJobsStateRequestBody) GetJobIdsOk() (*[]int64, bool) {
	if o == nil || o.JobIds == nil {
		return nil, false
	}
	return &o.JobIds, true
}

// HasJobIds returns a boolean if a field has been set.
func (o *UpdateProtectionJobsStateRequestBody) HasJobIds() bool {
	if o != nil && o.JobIds != nil {
		return true
	}

	return false
}

// SetJobIds gets a reference to the given []int64 and assigns it to the JobIds field.
func (o *UpdateProtectionJobsStateRequestBody) SetJobIds(v []int64) {
	o.JobIds = v
}

func (o UpdateProtectionJobsStateRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Action.IsSet() {
		toSerialize["action"] = o.Action.Get()
	}
	if o.JobIds != nil {
		toSerialize["jobIds"] = o.JobIds
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateProtectionJobsStateRequestBody struct {
	value *UpdateProtectionJobsStateRequestBody
	isSet bool
}

func (v NullableUpdateProtectionJobsStateRequestBody) Get() *UpdateProtectionJobsStateRequestBody {
	return v.value
}

func (v *NullableUpdateProtectionJobsStateRequestBody) Set(val *UpdateProtectionJobsStateRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateProtectionJobsStateRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateProtectionJobsStateRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateProtectionJobsStateRequestBody(val *UpdateProtectionJobsStateRequestBody) *NullableUpdateProtectionJobsStateRequestBody {
	return &NullableUpdateProtectionJobsStateRequestBody{value: val, isSet: true}
}

func (v NullableUpdateProtectionJobsStateRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateProtectionJobsStateRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


