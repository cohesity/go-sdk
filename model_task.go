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

// Task Specifies one task.
type Task struct {
	// The latest attributes reported for this task.
	Attributes []TaskAttribute `json:"attributes,omitempty"`
	// Specifies the end time of the task.
	EndTimeSeconds NullableInt64 `json:"endTimeSeconds,omitempty"`
	// Specifies an optional error message for this task.
	ErrorMessage NullableString `json:"errorMessage,omitempty"`
	// Specifies the events reported for this task.
	Events []TaskEvent `json:"events,omitempty"`
	// Specifies the estimated end time of the task.
	ExpectedEndTimeSeconds NullableInt64 `json:"expectedEndTimeSeconds,omitempty"`
	// Specifies the expected remaining time for this task in seconds.
	ExpectedSecondsRemaining NullableInt64 `json:"expectedSecondsRemaining,omitempty"`
	// The expected raw count of the total work remaining. This is the highest work count value reported by the client. This field can be set to let pulse compute percentFinished by looking at the currently reported remainingWorkCount and the expectedTotalWorkCount.
	ExpectedTotalWorkCount NullableInt64 `json:"expectedTotalWorkCount,omitempty"`
	// Specifies the timestamp when the last progress was reported.
	LastUpdateTimeSeconds NullableInt64 `json:"lastUpdateTimeSeconds,omitempty"`
	// Specifies the reported progress on the task.
	PercentFinished NullableFloat32 `json:"percentFinished,omitempty"`
	// Specifies the start time of the task.
	StartTimeSeconds NullableInt64 `json:"startTimeSeconds,omitempty"`
	// Specifies the status of the task being queried. 'kActive' indicates that the task is still active. 'kFinished' indicates that the task has finished without any errors. 'kFinishedWithError' indicates that the task has finished, but that there was an errror of some kind. 'kCancelled' indicates that the task was cancelled. 'kFinishedGarbageCollected' indicates that the task was garbage collected due to its subtasks not finishing within the allotted time.
	Status NullableString `json:"status,omitempty"`
	// Specifies a list of subtasks belonging to this task.
	SubTasks []map[string]interface{} `json:"subTasks,omitempty"`
	// Specifes the path of this task.
	TaskPath NullableString `json:"taskPath,omitempty"`
}

// NewTask instantiates a new Task object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTask() *Task {
	this := Task{}
	return &this
}

// NewTaskWithDefaults instantiates a new Task object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskWithDefaults() *Task {
	this := Task{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Task) GetAttributes() []TaskAttribute {
	if o == nil  {
		var ret []TaskAttribute
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Task) GetAttributesOk() (*[]TaskAttribute, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *Task) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []TaskAttribute and assigns it to the Attributes field.
func (o *Task) SetAttributes(v []TaskAttribute) {
	o.Attributes = v
}

// GetEndTimeSeconds returns the EndTimeSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Task) GetEndTimeSeconds() int64 {
	if o == nil || o.EndTimeSeconds.Get() == nil {
		var ret int64
		return ret
	}
	return *o.EndTimeSeconds.Get()
}

// GetEndTimeSecondsOk returns a tuple with the EndTimeSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Task) GetEndTimeSecondsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndTimeSeconds.Get(), o.EndTimeSeconds.IsSet()
}

// HasEndTimeSeconds returns a boolean if a field has been set.
func (o *Task) HasEndTimeSeconds() bool {
	if o != nil && o.EndTimeSeconds.IsSet() {
		return true
	}

	return false
}

// SetEndTimeSeconds gets a reference to the given NullableInt64 and assigns it to the EndTimeSeconds field.
func (o *Task) SetEndTimeSeconds(v int64) {
	o.EndTimeSeconds.Set(&v)
}
// SetEndTimeSecondsNil sets the value for EndTimeSeconds to be an explicit nil
func (o *Task) SetEndTimeSecondsNil() {
	o.EndTimeSeconds.Set(nil)
}

// UnsetEndTimeSeconds ensures that no value is present for EndTimeSeconds, not even an explicit nil
func (o *Task) UnsetEndTimeSeconds() {
	o.EndTimeSeconds.Unset()
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Task) GetErrorMessage() string {
	if o == nil || o.ErrorMessage.Get() == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage.Get()
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Task) GetErrorMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ErrorMessage.Get(), o.ErrorMessage.IsSet()
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *Task) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage.IsSet() {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given NullableString and assigns it to the ErrorMessage field.
func (o *Task) SetErrorMessage(v string) {
	o.ErrorMessage.Set(&v)
}
// SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil
func (o *Task) SetErrorMessageNil() {
	o.ErrorMessage.Set(nil)
}

// UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
func (o *Task) UnsetErrorMessage() {
	o.ErrorMessage.Unset()
}

// GetEvents returns the Events field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Task) GetEvents() []TaskEvent {
	if o == nil  {
		var ret []TaskEvent
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Task) GetEventsOk() (*[]TaskEvent, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return &o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *Task) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []TaskEvent and assigns it to the Events field.
func (o *Task) SetEvents(v []TaskEvent) {
	o.Events = v
}

// GetExpectedEndTimeSeconds returns the ExpectedEndTimeSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Task) GetExpectedEndTimeSeconds() int64 {
	if o == nil || o.ExpectedEndTimeSeconds.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ExpectedEndTimeSeconds.Get()
}

// GetExpectedEndTimeSecondsOk returns a tuple with the ExpectedEndTimeSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Task) GetExpectedEndTimeSecondsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExpectedEndTimeSeconds.Get(), o.ExpectedEndTimeSeconds.IsSet()
}

// HasExpectedEndTimeSeconds returns a boolean if a field has been set.
func (o *Task) HasExpectedEndTimeSeconds() bool {
	if o != nil && o.ExpectedEndTimeSeconds.IsSet() {
		return true
	}

	return false
}

// SetExpectedEndTimeSeconds gets a reference to the given NullableInt64 and assigns it to the ExpectedEndTimeSeconds field.
func (o *Task) SetExpectedEndTimeSeconds(v int64) {
	o.ExpectedEndTimeSeconds.Set(&v)
}
// SetExpectedEndTimeSecondsNil sets the value for ExpectedEndTimeSeconds to be an explicit nil
func (o *Task) SetExpectedEndTimeSecondsNil() {
	o.ExpectedEndTimeSeconds.Set(nil)
}

// UnsetExpectedEndTimeSeconds ensures that no value is present for ExpectedEndTimeSeconds, not even an explicit nil
func (o *Task) UnsetExpectedEndTimeSeconds() {
	o.ExpectedEndTimeSeconds.Unset()
}

// GetExpectedSecondsRemaining returns the ExpectedSecondsRemaining field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Task) GetExpectedSecondsRemaining() int64 {
	if o == nil || o.ExpectedSecondsRemaining.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ExpectedSecondsRemaining.Get()
}

// GetExpectedSecondsRemainingOk returns a tuple with the ExpectedSecondsRemaining field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Task) GetExpectedSecondsRemainingOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExpectedSecondsRemaining.Get(), o.ExpectedSecondsRemaining.IsSet()
}

// HasExpectedSecondsRemaining returns a boolean if a field has been set.
func (o *Task) HasExpectedSecondsRemaining() bool {
	if o != nil && o.ExpectedSecondsRemaining.IsSet() {
		return true
	}

	return false
}

// SetExpectedSecondsRemaining gets a reference to the given NullableInt64 and assigns it to the ExpectedSecondsRemaining field.
func (o *Task) SetExpectedSecondsRemaining(v int64) {
	o.ExpectedSecondsRemaining.Set(&v)
}
// SetExpectedSecondsRemainingNil sets the value for ExpectedSecondsRemaining to be an explicit nil
func (o *Task) SetExpectedSecondsRemainingNil() {
	o.ExpectedSecondsRemaining.Set(nil)
}

// UnsetExpectedSecondsRemaining ensures that no value is present for ExpectedSecondsRemaining, not even an explicit nil
func (o *Task) UnsetExpectedSecondsRemaining() {
	o.ExpectedSecondsRemaining.Unset()
}

// GetExpectedTotalWorkCount returns the ExpectedTotalWorkCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Task) GetExpectedTotalWorkCount() int64 {
	if o == nil || o.ExpectedTotalWorkCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ExpectedTotalWorkCount.Get()
}

// GetExpectedTotalWorkCountOk returns a tuple with the ExpectedTotalWorkCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Task) GetExpectedTotalWorkCountOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExpectedTotalWorkCount.Get(), o.ExpectedTotalWorkCount.IsSet()
}

// HasExpectedTotalWorkCount returns a boolean if a field has been set.
func (o *Task) HasExpectedTotalWorkCount() bool {
	if o != nil && o.ExpectedTotalWorkCount.IsSet() {
		return true
	}

	return false
}

// SetExpectedTotalWorkCount gets a reference to the given NullableInt64 and assigns it to the ExpectedTotalWorkCount field.
func (o *Task) SetExpectedTotalWorkCount(v int64) {
	o.ExpectedTotalWorkCount.Set(&v)
}
// SetExpectedTotalWorkCountNil sets the value for ExpectedTotalWorkCount to be an explicit nil
func (o *Task) SetExpectedTotalWorkCountNil() {
	o.ExpectedTotalWorkCount.Set(nil)
}

// UnsetExpectedTotalWorkCount ensures that no value is present for ExpectedTotalWorkCount, not even an explicit nil
func (o *Task) UnsetExpectedTotalWorkCount() {
	o.ExpectedTotalWorkCount.Unset()
}

// GetLastUpdateTimeSeconds returns the LastUpdateTimeSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Task) GetLastUpdateTimeSeconds() int64 {
	if o == nil || o.LastUpdateTimeSeconds.Get() == nil {
		var ret int64
		return ret
	}
	return *o.LastUpdateTimeSeconds.Get()
}

// GetLastUpdateTimeSecondsOk returns a tuple with the LastUpdateTimeSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Task) GetLastUpdateTimeSecondsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastUpdateTimeSeconds.Get(), o.LastUpdateTimeSeconds.IsSet()
}

// HasLastUpdateTimeSeconds returns a boolean if a field has been set.
func (o *Task) HasLastUpdateTimeSeconds() bool {
	if o != nil && o.LastUpdateTimeSeconds.IsSet() {
		return true
	}

	return false
}

// SetLastUpdateTimeSeconds gets a reference to the given NullableInt64 and assigns it to the LastUpdateTimeSeconds field.
func (o *Task) SetLastUpdateTimeSeconds(v int64) {
	o.LastUpdateTimeSeconds.Set(&v)
}
// SetLastUpdateTimeSecondsNil sets the value for LastUpdateTimeSeconds to be an explicit nil
func (o *Task) SetLastUpdateTimeSecondsNil() {
	o.LastUpdateTimeSeconds.Set(nil)
}

// UnsetLastUpdateTimeSeconds ensures that no value is present for LastUpdateTimeSeconds, not even an explicit nil
func (o *Task) UnsetLastUpdateTimeSeconds() {
	o.LastUpdateTimeSeconds.Unset()
}

// GetPercentFinished returns the PercentFinished field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Task) GetPercentFinished() float32 {
	if o == nil || o.PercentFinished.Get() == nil {
		var ret float32
		return ret
	}
	return *o.PercentFinished.Get()
}

// GetPercentFinishedOk returns a tuple with the PercentFinished field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Task) GetPercentFinishedOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PercentFinished.Get(), o.PercentFinished.IsSet()
}

// HasPercentFinished returns a boolean if a field has been set.
func (o *Task) HasPercentFinished() bool {
	if o != nil && o.PercentFinished.IsSet() {
		return true
	}

	return false
}

// SetPercentFinished gets a reference to the given NullableFloat32 and assigns it to the PercentFinished field.
func (o *Task) SetPercentFinished(v float32) {
	o.PercentFinished.Set(&v)
}
// SetPercentFinishedNil sets the value for PercentFinished to be an explicit nil
func (o *Task) SetPercentFinishedNil() {
	o.PercentFinished.Set(nil)
}

// UnsetPercentFinished ensures that no value is present for PercentFinished, not even an explicit nil
func (o *Task) UnsetPercentFinished() {
	o.PercentFinished.Unset()
}

// GetStartTimeSeconds returns the StartTimeSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Task) GetStartTimeSeconds() int64 {
	if o == nil || o.StartTimeSeconds.Get() == nil {
		var ret int64
		return ret
	}
	return *o.StartTimeSeconds.Get()
}

// GetStartTimeSecondsOk returns a tuple with the StartTimeSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Task) GetStartTimeSecondsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartTimeSeconds.Get(), o.StartTimeSeconds.IsSet()
}

// HasStartTimeSeconds returns a boolean if a field has been set.
func (o *Task) HasStartTimeSeconds() bool {
	if o != nil && o.StartTimeSeconds.IsSet() {
		return true
	}

	return false
}

// SetStartTimeSeconds gets a reference to the given NullableInt64 and assigns it to the StartTimeSeconds field.
func (o *Task) SetStartTimeSeconds(v int64) {
	o.StartTimeSeconds.Set(&v)
}
// SetStartTimeSecondsNil sets the value for StartTimeSeconds to be an explicit nil
func (o *Task) SetStartTimeSecondsNil() {
	o.StartTimeSeconds.Set(nil)
}

// UnsetStartTimeSeconds ensures that no value is present for StartTimeSeconds, not even an explicit nil
func (o *Task) UnsetStartTimeSeconds() {
	o.StartTimeSeconds.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Task) GetStatus() string {
	if o == nil || o.Status.Get() == nil {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Task) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *Task) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *Task) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *Task) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *Task) UnsetStatus() {
	o.Status.Unset()
}

// GetSubTasks returns the SubTasks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Task) GetSubTasks() []map[string]interface{} {
	if o == nil  {
		var ret []map[string]interface{}
		return ret
	}
	return o.SubTasks
}

// GetSubTasksOk returns a tuple with the SubTasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Task) GetSubTasksOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.SubTasks == nil {
		return nil, false
	}
	return &o.SubTasks, true
}

// HasSubTasks returns a boolean if a field has been set.
func (o *Task) HasSubTasks() bool {
	if o != nil && o.SubTasks != nil {
		return true
	}

	return false
}

// SetSubTasks gets a reference to the given []map[string]interface{} and assigns it to the SubTasks field.
func (o *Task) SetSubTasks(v []map[string]interface{}) {
	o.SubTasks = v
}

// GetTaskPath returns the TaskPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Task) GetTaskPath() string {
	if o == nil || o.TaskPath.Get() == nil {
		var ret string
		return ret
	}
	return *o.TaskPath.Get()
}

// GetTaskPathOk returns a tuple with the TaskPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Task) GetTaskPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TaskPath.Get(), o.TaskPath.IsSet()
}

// HasTaskPath returns a boolean if a field has been set.
func (o *Task) HasTaskPath() bool {
	if o != nil && o.TaskPath.IsSet() {
		return true
	}

	return false
}

// SetTaskPath gets a reference to the given NullableString and assigns it to the TaskPath field.
func (o *Task) SetTaskPath(v string) {
	o.TaskPath.Set(&v)
}
// SetTaskPathNil sets the value for TaskPath to be an explicit nil
func (o *Task) SetTaskPathNil() {
	o.TaskPath.Set(nil)
}

// UnsetTaskPath ensures that no value is present for TaskPath, not even an explicit nil
func (o *Task) UnsetTaskPath() {
	o.TaskPath.Unset()
}

func (o Task) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.EndTimeSeconds.IsSet() {
		toSerialize["endTimeSeconds"] = o.EndTimeSeconds.Get()
	}
	if o.ErrorMessage.IsSet() {
		toSerialize["errorMessage"] = o.ErrorMessage.Get()
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	if o.ExpectedEndTimeSeconds.IsSet() {
		toSerialize["expectedEndTimeSeconds"] = o.ExpectedEndTimeSeconds.Get()
	}
	if o.ExpectedSecondsRemaining.IsSet() {
		toSerialize["expectedSecondsRemaining"] = o.ExpectedSecondsRemaining.Get()
	}
	if o.ExpectedTotalWorkCount.IsSet() {
		toSerialize["expectedTotalWorkCount"] = o.ExpectedTotalWorkCount.Get()
	}
	if o.LastUpdateTimeSeconds.IsSet() {
		toSerialize["lastUpdateTimeSeconds"] = o.LastUpdateTimeSeconds.Get()
	}
	if o.PercentFinished.IsSet() {
		toSerialize["percentFinished"] = o.PercentFinished.Get()
	}
	if o.StartTimeSeconds.IsSet() {
		toSerialize["startTimeSeconds"] = o.StartTimeSeconds.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.SubTasks != nil {
		toSerialize["subTasks"] = o.SubTasks
	}
	if o.TaskPath.IsSet() {
		toSerialize["taskPath"] = o.TaskPath.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableTask struct {
	value *Task
	isSet bool
}

func (v NullableTask) Get() *Task {
	return v.value
}

func (v *NullableTask) Set(val *Task) {
	v.value = val
	v.isSet = true
}

func (v NullableTask) IsSet() bool {
	return v.isSet
}

func (v *NullableTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTask(val *Task) *NullableTask {
	return &NullableTask{value: val, isSet: true}
}

func (v NullableTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


