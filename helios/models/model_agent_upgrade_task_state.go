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

// AgentUpgradeTaskState Specifies the state of an agent upgrade task.
type AgentUpgradeTaskState struct {
	// Specifies the name of the task.
	Name NullableString `json:"name,omitempty"`
	// Specifies the description of the task.
	Description NullableString `json:"description,omitempty"`
	// Specifies the ID of the task.
	Id NullableInt64 `json:"id,omitempty"`
	// Specifies the time, as a Unix epoch timestamp in microseconds, when the task started execution.
	StartTimeUsecs NullableInt64 `json:"startTimeUsecs,omitempty"`
	// Specifies the time when the upgrade task completed execution as a Unix epoch Timestamp (in microseconds).
	EndTimeUsecs NullableInt64 `json:"endTimeUsecs,omitempty"`
	// Specifies the status of the task.<br> 'Scheduled' indicates that the upgrade task is yet to start.<br> 'Running' indicates that the upgrade task has started execution.<br> 'Succeeded' indicates that the upgrade task completed without an error.<br> 'Failed' indicates that upgrade has failed for all agents. 'PartiallyFailed' indicates that upgrade has failed for some agents.
	Status NullableString `json:"status,omitempty"`
	// Specifies the agents upgraded in the task.
	AgentIDs []int64 `json:"agentIDs,omitempty"`
	// Specifies the version to which agents are upgraded.
	ClusterVersion NullableString `json:"clusterVersion,omitempty"`
	// Specifes the type of task.<br> 'Auto' indicates an auto agent upgrade task which is started after a cluster upgrade.<br> 'Manual' indicates a schedule based agent upgrade task.<br> 'Retry' indicates an agent upgrade task which was retried.
	Type NullableString `json:"type,omitempty"`
	// Specifies the time when the task should start execution as a Unix epoch Timestamp (in microseconds). If no schedule is specified, the task will start immediately.
	ScheduleTimeUsecs NullableInt64 `json:"scheduleTimeUsecs,omitempty"`
	// Specifies the time before which the upgrade task should start execution as a Unix epoch Timestamp (in microseconds). If this is not specified the task will start anytime after scheduleTimeUsecs.
	ScheduleEndTimeUsecs NullableInt64 `json:"scheduleEndTimeUsecs,omitempty"`
	// Specifies ID of a task which was retried if type is 'Retry'.
	RetriedTaskID NullableInt64 `json:"retriedTaskID,omitempty"`
	// Specifies if a task can be retried.
	IsRetryable NullableBool `json:"isRetryable,omitempty"`
	Error *Error `json:"error,omitempty"`
	// Specifies the upgrade information for each agent.
	Agents []AgentUpgradeInfoObject `json:"agents,omitempty"`
}

// NewAgentUpgradeTaskState instantiates a new AgentUpgradeTaskState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentUpgradeTaskState() *AgentUpgradeTaskState {
	this := AgentUpgradeTaskState{}
	return &this
}

// NewAgentUpgradeTaskStateWithDefaults instantiates a new AgentUpgradeTaskState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentUpgradeTaskStateWithDefaults() *AgentUpgradeTaskState {
	this := AgentUpgradeTaskState{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgentUpgradeTaskState) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentUpgradeTaskState) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *AgentUpgradeTaskState) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *AgentUpgradeTaskState) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *AgentUpgradeTaskState) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *AgentUpgradeTaskState) UnsetName() {
	o.Name.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgentUpgradeTaskState) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentUpgradeTaskState) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *AgentUpgradeTaskState) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *AgentUpgradeTaskState) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *AgentUpgradeTaskState) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *AgentUpgradeTaskState) UnsetDescription() {
	o.Description.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgentUpgradeTaskState) GetId() int64 {
	if o == nil || o.Id.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentUpgradeTaskState) GetIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *AgentUpgradeTaskState) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *AgentUpgradeTaskState) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *AgentUpgradeTaskState) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *AgentUpgradeTaskState) UnsetId() {
	o.Id.Unset()
}

// GetStartTimeUsecs returns the StartTimeUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgentUpgradeTaskState) GetStartTimeUsecs() int64 {
	if o == nil || o.StartTimeUsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.StartTimeUsecs.Get()
}

// GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentUpgradeTaskState) GetStartTimeUsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartTimeUsecs.Get(), o.StartTimeUsecs.IsSet()
}

// HasStartTimeUsecs returns a boolean if a field has been set.
func (o *AgentUpgradeTaskState) HasStartTimeUsecs() bool {
	if o != nil && o.StartTimeUsecs.IsSet() {
		return true
	}

	return false
}

// SetStartTimeUsecs gets a reference to the given NullableInt64 and assigns it to the StartTimeUsecs field.
func (o *AgentUpgradeTaskState) SetStartTimeUsecs(v int64) {
	o.StartTimeUsecs.Set(&v)
}
// SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil
func (o *AgentUpgradeTaskState) SetStartTimeUsecsNil() {
	o.StartTimeUsecs.Set(nil)
}

// UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
func (o *AgentUpgradeTaskState) UnsetStartTimeUsecs() {
	o.StartTimeUsecs.Unset()
}

// GetEndTimeUsecs returns the EndTimeUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgentUpgradeTaskState) GetEndTimeUsecs() int64 {
	if o == nil || o.EndTimeUsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.EndTimeUsecs.Get()
}

// GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentUpgradeTaskState) GetEndTimeUsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndTimeUsecs.Get(), o.EndTimeUsecs.IsSet()
}

// HasEndTimeUsecs returns a boolean if a field has been set.
func (o *AgentUpgradeTaskState) HasEndTimeUsecs() bool {
	if o != nil && o.EndTimeUsecs.IsSet() {
		return true
	}

	return false
}

// SetEndTimeUsecs gets a reference to the given NullableInt64 and assigns it to the EndTimeUsecs field.
func (o *AgentUpgradeTaskState) SetEndTimeUsecs(v int64) {
	o.EndTimeUsecs.Set(&v)
}
// SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil
func (o *AgentUpgradeTaskState) SetEndTimeUsecsNil() {
	o.EndTimeUsecs.Set(nil)
}

// UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
func (o *AgentUpgradeTaskState) UnsetEndTimeUsecs() {
	o.EndTimeUsecs.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgentUpgradeTaskState) GetStatus() string {
	if o == nil || o.Status.Get() == nil {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentUpgradeTaskState) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *AgentUpgradeTaskState) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *AgentUpgradeTaskState) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *AgentUpgradeTaskState) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *AgentUpgradeTaskState) UnsetStatus() {
	o.Status.Unset()
}

// GetAgentIDs returns the AgentIDs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgentUpgradeTaskState) GetAgentIDs() []int64 {
	if o == nil  {
		var ret []int64
		return ret
	}
	return o.AgentIDs
}

// GetAgentIDsOk returns a tuple with the AgentIDs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentUpgradeTaskState) GetAgentIDsOk() (*[]int64, bool) {
	if o == nil || o.AgentIDs == nil {
		return nil, false
	}
	return &o.AgentIDs, true
}

// HasAgentIDs returns a boolean if a field has been set.
func (o *AgentUpgradeTaskState) HasAgentIDs() bool {
	if o != nil && o.AgentIDs != nil {
		return true
	}

	return false
}

// SetAgentIDs gets a reference to the given []int64 and assigns it to the AgentIDs field.
func (o *AgentUpgradeTaskState) SetAgentIDs(v []int64) {
	o.AgentIDs = v
}

// GetClusterVersion returns the ClusterVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgentUpgradeTaskState) GetClusterVersion() string {
	if o == nil || o.ClusterVersion.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClusterVersion.Get()
}

// GetClusterVersionOk returns a tuple with the ClusterVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentUpgradeTaskState) GetClusterVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ClusterVersion.Get(), o.ClusterVersion.IsSet()
}

// HasClusterVersion returns a boolean if a field has been set.
func (o *AgentUpgradeTaskState) HasClusterVersion() bool {
	if o != nil && o.ClusterVersion.IsSet() {
		return true
	}

	return false
}

// SetClusterVersion gets a reference to the given NullableString and assigns it to the ClusterVersion field.
func (o *AgentUpgradeTaskState) SetClusterVersion(v string) {
	o.ClusterVersion.Set(&v)
}
// SetClusterVersionNil sets the value for ClusterVersion to be an explicit nil
func (o *AgentUpgradeTaskState) SetClusterVersionNil() {
	o.ClusterVersion.Set(nil)
}

// UnsetClusterVersion ensures that no value is present for ClusterVersion, not even an explicit nil
func (o *AgentUpgradeTaskState) UnsetClusterVersion() {
	o.ClusterVersion.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgentUpgradeTaskState) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentUpgradeTaskState) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *AgentUpgradeTaskState) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *AgentUpgradeTaskState) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *AgentUpgradeTaskState) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *AgentUpgradeTaskState) UnsetType() {
	o.Type.Unset()
}

// GetScheduleTimeUsecs returns the ScheduleTimeUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgentUpgradeTaskState) GetScheduleTimeUsecs() int64 {
	if o == nil || o.ScheduleTimeUsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ScheduleTimeUsecs.Get()
}

// GetScheduleTimeUsecsOk returns a tuple with the ScheduleTimeUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentUpgradeTaskState) GetScheduleTimeUsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ScheduleTimeUsecs.Get(), o.ScheduleTimeUsecs.IsSet()
}

// HasScheduleTimeUsecs returns a boolean if a field has been set.
func (o *AgentUpgradeTaskState) HasScheduleTimeUsecs() bool {
	if o != nil && o.ScheduleTimeUsecs.IsSet() {
		return true
	}

	return false
}

// SetScheduleTimeUsecs gets a reference to the given NullableInt64 and assigns it to the ScheduleTimeUsecs field.
func (o *AgentUpgradeTaskState) SetScheduleTimeUsecs(v int64) {
	o.ScheduleTimeUsecs.Set(&v)
}
// SetScheduleTimeUsecsNil sets the value for ScheduleTimeUsecs to be an explicit nil
func (o *AgentUpgradeTaskState) SetScheduleTimeUsecsNil() {
	o.ScheduleTimeUsecs.Set(nil)
}

// UnsetScheduleTimeUsecs ensures that no value is present for ScheduleTimeUsecs, not even an explicit nil
func (o *AgentUpgradeTaskState) UnsetScheduleTimeUsecs() {
	o.ScheduleTimeUsecs.Unset()
}

// GetScheduleEndTimeUsecs returns the ScheduleEndTimeUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgentUpgradeTaskState) GetScheduleEndTimeUsecs() int64 {
	if o == nil || o.ScheduleEndTimeUsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ScheduleEndTimeUsecs.Get()
}

// GetScheduleEndTimeUsecsOk returns a tuple with the ScheduleEndTimeUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentUpgradeTaskState) GetScheduleEndTimeUsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ScheduleEndTimeUsecs.Get(), o.ScheduleEndTimeUsecs.IsSet()
}

// HasScheduleEndTimeUsecs returns a boolean if a field has been set.
func (o *AgentUpgradeTaskState) HasScheduleEndTimeUsecs() bool {
	if o != nil && o.ScheduleEndTimeUsecs.IsSet() {
		return true
	}

	return false
}

// SetScheduleEndTimeUsecs gets a reference to the given NullableInt64 and assigns it to the ScheduleEndTimeUsecs field.
func (o *AgentUpgradeTaskState) SetScheduleEndTimeUsecs(v int64) {
	o.ScheduleEndTimeUsecs.Set(&v)
}
// SetScheduleEndTimeUsecsNil sets the value for ScheduleEndTimeUsecs to be an explicit nil
func (o *AgentUpgradeTaskState) SetScheduleEndTimeUsecsNil() {
	o.ScheduleEndTimeUsecs.Set(nil)
}

// UnsetScheduleEndTimeUsecs ensures that no value is present for ScheduleEndTimeUsecs, not even an explicit nil
func (o *AgentUpgradeTaskState) UnsetScheduleEndTimeUsecs() {
	o.ScheduleEndTimeUsecs.Unset()
}

// GetRetriedTaskID returns the RetriedTaskID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgentUpgradeTaskState) GetRetriedTaskID() int64 {
	if o == nil || o.RetriedTaskID.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RetriedTaskID.Get()
}

// GetRetriedTaskIDOk returns a tuple with the RetriedTaskID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentUpgradeTaskState) GetRetriedTaskIDOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RetriedTaskID.Get(), o.RetriedTaskID.IsSet()
}

// HasRetriedTaskID returns a boolean if a field has been set.
func (o *AgentUpgradeTaskState) HasRetriedTaskID() bool {
	if o != nil && o.RetriedTaskID.IsSet() {
		return true
	}

	return false
}

// SetRetriedTaskID gets a reference to the given NullableInt64 and assigns it to the RetriedTaskID field.
func (o *AgentUpgradeTaskState) SetRetriedTaskID(v int64) {
	o.RetriedTaskID.Set(&v)
}
// SetRetriedTaskIDNil sets the value for RetriedTaskID to be an explicit nil
func (o *AgentUpgradeTaskState) SetRetriedTaskIDNil() {
	o.RetriedTaskID.Set(nil)
}

// UnsetRetriedTaskID ensures that no value is present for RetriedTaskID, not even an explicit nil
func (o *AgentUpgradeTaskState) UnsetRetriedTaskID() {
	o.RetriedTaskID.Unset()
}

// GetIsRetryable returns the IsRetryable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgentUpgradeTaskState) GetIsRetryable() bool {
	if o == nil || o.IsRetryable.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsRetryable.Get()
}

// GetIsRetryableOk returns a tuple with the IsRetryable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentUpgradeTaskState) GetIsRetryableOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsRetryable.Get(), o.IsRetryable.IsSet()
}

// HasIsRetryable returns a boolean if a field has been set.
func (o *AgentUpgradeTaskState) HasIsRetryable() bool {
	if o != nil && o.IsRetryable.IsSet() {
		return true
	}

	return false
}

// SetIsRetryable gets a reference to the given NullableBool and assigns it to the IsRetryable field.
func (o *AgentUpgradeTaskState) SetIsRetryable(v bool) {
	o.IsRetryable.Set(&v)
}
// SetIsRetryableNil sets the value for IsRetryable to be an explicit nil
func (o *AgentUpgradeTaskState) SetIsRetryableNil() {
	o.IsRetryable.Set(nil)
}

// UnsetIsRetryable ensures that no value is present for IsRetryable, not even an explicit nil
func (o *AgentUpgradeTaskState) UnsetIsRetryable() {
	o.IsRetryable.Unset()
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *AgentUpgradeTaskState) GetError() Error {
	if o == nil || o.Error == nil {
		var ret Error
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentUpgradeTaskState) GetErrorOk() (*Error, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *AgentUpgradeTaskState) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given Error and assigns it to the Error field.
func (o *AgentUpgradeTaskState) SetError(v Error) {
	o.Error = &v
}

// GetAgents returns the Agents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgentUpgradeTaskState) GetAgents() []AgentUpgradeInfoObject {
	if o == nil  {
		var ret []AgentUpgradeInfoObject
		return ret
	}
	return o.Agents
}

// GetAgentsOk returns a tuple with the Agents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentUpgradeTaskState) GetAgentsOk() (*[]AgentUpgradeInfoObject, bool) {
	if o == nil || o.Agents == nil {
		return nil, false
	}
	return &o.Agents, true
}

// HasAgents returns a boolean if a field has been set.
func (o *AgentUpgradeTaskState) HasAgents() bool {
	if o != nil && o.Agents != nil {
		return true
	}

	return false
}

// SetAgents gets a reference to the given []AgentUpgradeInfoObject and assigns it to the Agents field.
func (o *AgentUpgradeTaskState) SetAgents(v []AgentUpgradeInfoObject) {
	o.Agents = v
}

func (o AgentUpgradeTaskState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.StartTimeUsecs.IsSet() {
		toSerialize["startTimeUsecs"] = o.StartTimeUsecs.Get()
	}
	if o.EndTimeUsecs.IsSet() {
		toSerialize["endTimeUsecs"] = o.EndTimeUsecs.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.AgentIDs != nil {
		toSerialize["agentIDs"] = o.AgentIDs
	}
	if o.ClusterVersion.IsSet() {
		toSerialize["clusterVersion"] = o.ClusterVersion.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.ScheduleTimeUsecs.IsSet() {
		toSerialize["scheduleTimeUsecs"] = o.ScheduleTimeUsecs.Get()
	}
	if o.ScheduleEndTimeUsecs.IsSet() {
		toSerialize["scheduleEndTimeUsecs"] = o.ScheduleEndTimeUsecs.Get()
	}
	if o.RetriedTaskID.IsSet() {
		toSerialize["retriedTaskID"] = o.RetriedTaskID.Get()
	}
	if o.IsRetryable.IsSet() {
		toSerialize["isRetryable"] = o.IsRetryable.Get()
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.Agents != nil {
		toSerialize["agents"] = o.Agents
	}
	return json.Marshal(toSerialize)
}

type NullableAgentUpgradeTaskState struct {
	value *AgentUpgradeTaskState
	isSet bool
}

func (v NullableAgentUpgradeTaskState) Get() *AgentUpgradeTaskState {
	return v.value
}

func (v *NullableAgentUpgradeTaskState) Set(val *AgentUpgradeTaskState) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentUpgradeTaskState) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentUpgradeTaskState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentUpgradeTaskState(val *AgentUpgradeTaskState) *NullableAgentUpgradeTaskState {
	return &NullableAgentUpgradeTaskState{value: val, isSet: true}
}

func (v NullableAgentUpgradeTaskState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentUpgradeTaskState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


