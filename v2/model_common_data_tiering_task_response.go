/*
Cohesity REST API

Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CommonDataTieringTaskResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonDataTieringTaskResponse{}

// CommonDataTieringTaskResponse Specifies the data tiering task details.
type CommonDataTieringTaskResponse struct {
	AlertPolicy *ProtectionGroupAlertingPolicy `json:"alertPolicy,omitempty"`
	// Specifies a description of the data tiering task.
	Description NullableString `json:"description,omitempty"`
	// Specifies the id of the data tiering task.
	Id NullableString `json:"id,omitempty"`
	// Whether the data tiering task is active or not.
	IsActive NullableBool `json:"isActive,omitempty"`
	// Tracks whether the backup job has actually been deleted.
	IsDeleted NullableBool `json:"isDeleted,omitempty"`
	// Whether the data tiering task is paused. New runs are not scheduled for the paused tasks. Active run of the task (if any) is not impacted.
	IsPaused NullableBool `json:"isPaused,omitempty"`
	LastRun *DataTieringTaskRun `json:"lastRun,omitempty"`
	// Specifies the name of the data tiering task.
	Name NullableString `json:"name"`
	Schedule *DataTieringSchedule `json:"schedule,omitempty"`
	Source *DataTieringSource `json:"source,omitempty"`
	Target NullableDataTieringTarget `json:"target,omitempty"`
	// Type of data tiering task. 'Downtier' indicates downtiering task. 'Uptier' indicates uptiering task.
	Type NullableString `json:"type"`
}

type _CommonDataTieringTaskResponse CommonDataTieringTaskResponse

// NewCommonDataTieringTaskResponse instantiates a new CommonDataTieringTaskResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonDataTieringTaskResponse(name NullableString, type_ NullableString) *CommonDataTieringTaskResponse {
	this := CommonDataTieringTaskResponse{}
	var isActive bool = true
	this.IsActive = *NewNullableBool(&isActive)
	var isDeleted bool = true
	this.IsDeleted = *NewNullableBool(&isDeleted)
	var isPaused bool = true
	this.IsPaused = *NewNullableBool(&isPaused)
	this.Name = name
	this.Type = type_
	return &this
}

// NewCommonDataTieringTaskResponseWithDefaults instantiates a new CommonDataTieringTaskResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonDataTieringTaskResponseWithDefaults() *CommonDataTieringTaskResponse {
	this := CommonDataTieringTaskResponse{}
	var isActive bool = true
	this.IsActive = *NewNullableBool(&isActive)
	var isDeleted bool = true
	this.IsDeleted = *NewNullableBool(&isDeleted)
	var isPaused bool = true
	this.IsPaused = *NewNullableBool(&isPaused)
	return &this
}

// GetAlertPolicy returns the AlertPolicy field value if set, zero value otherwise.
func (o *CommonDataTieringTaskResponse) GetAlertPolicy() ProtectionGroupAlertingPolicy {
	if o == nil || IsNil(o.AlertPolicy) {
		var ret ProtectionGroupAlertingPolicy
		return ret
	}
	return *o.AlertPolicy
}

// GetAlertPolicyOk returns a tuple with the AlertPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonDataTieringTaskResponse) GetAlertPolicyOk() (*ProtectionGroupAlertingPolicy, bool) {
	if o == nil || IsNil(o.AlertPolicy) {
		return nil, false
	}
	return o.AlertPolicy, true
}

// HasAlertPolicy returns a boolean if a field has been set.
func (o *CommonDataTieringTaskResponse) HasAlertPolicy() bool {
	if o != nil && !IsNil(o.AlertPolicy) {
		return true
	}

	return false
}

// SetAlertPolicy gets a reference to the given ProtectionGroupAlertingPolicy and assigns it to the AlertPolicy field.
func (o *CommonDataTieringTaskResponse) SetAlertPolicy(v ProtectionGroupAlertingPolicy) {
	o.AlertPolicy = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonDataTieringTaskResponse) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonDataTieringTaskResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *CommonDataTieringTaskResponse) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *CommonDataTieringTaskResponse) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *CommonDataTieringTaskResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *CommonDataTieringTaskResponse) UnsetDescription() {
	o.Description.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonDataTieringTaskResponse) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonDataTieringTaskResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *CommonDataTieringTaskResponse) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *CommonDataTieringTaskResponse) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *CommonDataTieringTaskResponse) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *CommonDataTieringTaskResponse) UnsetId() {
	o.Id.Unset()
}

// GetIsActive returns the IsActive field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonDataTieringTaskResponse) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive.Get()) {
		var ret bool
		return ret
	}
	return *o.IsActive.Get()
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonDataTieringTaskResponse) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsActive.Get(), o.IsActive.IsSet()
}

// HasIsActive returns a boolean if a field has been set.
func (o *CommonDataTieringTaskResponse) HasIsActive() bool {
	if o != nil && o.IsActive.IsSet() {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given NullableBool and assigns it to the IsActive field.
func (o *CommonDataTieringTaskResponse) SetIsActive(v bool) {
	o.IsActive.Set(&v)
}
// SetIsActiveNil sets the value for IsActive to be an explicit nil
func (o *CommonDataTieringTaskResponse) SetIsActiveNil() {
	o.IsActive.Set(nil)
}

// UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
func (o *CommonDataTieringTaskResponse) UnsetIsActive() {
	o.IsActive.Unset()
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonDataTieringTaskResponse) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted.Get()) {
		var ret bool
		return ret
	}
	return *o.IsDeleted.Get()
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonDataTieringTaskResponse) GetIsDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsDeleted.Get(), o.IsDeleted.IsSet()
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *CommonDataTieringTaskResponse) HasIsDeleted() bool {
	if o != nil && o.IsDeleted.IsSet() {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given NullableBool and assigns it to the IsDeleted field.
func (o *CommonDataTieringTaskResponse) SetIsDeleted(v bool) {
	o.IsDeleted.Set(&v)
}
// SetIsDeletedNil sets the value for IsDeleted to be an explicit nil
func (o *CommonDataTieringTaskResponse) SetIsDeletedNil() {
	o.IsDeleted.Set(nil)
}

// UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil
func (o *CommonDataTieringTaskResponse) UnsetIsDeleted() {
	o.IsDeleted.Unset()
}

// GetIsPaused returns the IsPaused field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonDataTieringTaskResponse) GetIsPaused() bool {
	if o == nil || IsNil(o.IsPaused.Get()) {
		var ret bool
		return ret
	}
	return *o.IsPaused.Get()
}

// GetIsPausedOk returns a tuple with the IsPaused field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonDataTieringTaskResponse) GetIsPausedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsPaused.Get(), o.IsPaused.IsSet()
}

// HasIsPaused returns a boolean if a field has been set.
func (o *CommonDataTieringTaskResponse) HasIsPaused() bool {
	if o != nil && o.IsPaused.IsSet() {
		return true
	}

	return false
}

// SetIsPaused gets a reference to the given NullableBool and assigns it to the IsPaused field.
func (o *CommonDataTieringTaskResponse) SetIsPaused(v bool) {
	o.IsPaused.Set(&v)
}
// SetIsPausedNil sets the value for IsPaused to be an explicit nil
func (o *CommonDataTieringTaskResponse) SetIsPausedNil() {
	o.IsPaused.Set(nil)
}

// UnsetIsPaused ensures that no value is present for IsPaused, not even an explicit nil
func (o *CommonDataTieringTaskResponse) UnsetIsPaused() {
	o.IsPaused.Unset()
}

// GetLastRun returns the LastRun field value if set, zero value otherwise.
func (o *CommonDataTieringTaskResponse) GetLastRun() DataTieringTaskRun {
	if o == nil || IsNil(o.LastRun) {
		var ret DataTieringTaskRun
		return ret
	}
	return *o.LastRun
}

// GetLastRunOk returns a tuple with the LastRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonDataTieringTaskResponse) GetLastRunOk() (*DataTieringTaskRun, bool) {
	if o == nil || IsNil(o.LastRun) {
		return nil, false
	}
	return o.LastRun, true
}

// HasLastRun returns a boolean if a field has been set.
func (o *CommonDataTieringTaskResponse) HasLastRun() bool {
	if o != nil && !IsNil(o.LastRun) {
		return true
	}

	return false
}

// SetLastRun gets a reference to the given DataTieringTaskRun and assigns it to the LastRun field.
func (o *CommonDataTieringTaskResponse) SetLastRun(v DataTieringTaskRun) {
	o.LastRun = &v
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CommonDataTieringTaskResponse) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonDataTieringTaskResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *CommonDataTieringTaskResponse) SetName(v string) {
	o.Name.Set(&v)
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *CommonDataTieringTaskResponse) GetSchedule() DataTieringSchedule {
	if o == nil || IsNil(o.Schedule) {
		var ret DataTieringSchedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonDataTieringTaskResponse) GetScheduleOk() (*DataTieringSchedule, bool) {
	if o == nil || IsNil(o.Schedule) {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *CommonDataTieringTaskResponse) HasSchedule() bool {
	if o != nil && !IsNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given DataTieringSchedule and assigns it to the Schedule field.
func (o *CommonDataTieringTaskResponse) SetSchedule(v DataTieringSchedule) {
	o.Schedule = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *CommonDataTieringTaskResponse) GetSource() DataTieringSource {
	if o == nil || IsNil(o.Source) {
		var ret DataTieringSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonDataTieringTaskResponse) GetSourceOk() (*DataTieringSource, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *CommonDataTieringTaskResponse) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given DataTieringSource and assigns it to the Source field.
func (o *CommonDataTieringTaskResponse) SetSource(v DataTieringSource) {
	o.Source = &v
}

// GetTarget returns the Target field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonDataTieringTaskResponse) GetTarget() DataTieringTarget {
	if o == nil || IsNil(o.Target.Get()) {
		var ret DataTieringTarget
		return ret
	}
	return *o.Target.Get()
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonDataTieringTaskResponse) GetTargetOk() (*DataTieringTarget, bool) {
	if o == nil {
		return nil, false
	}
	return o.Target.Get(), o.Target.IsSet()
}

// HasTarget returns a boolean if a field has been set.
func (o *CommonDataTieringTaskResponse) HasTarget() bool {
	if o != nil && o.Target.IsSet() {
		return true
	}

	return false
}

// SetTarget gets a reference to the given NullableDataTieringTarget and assigns it to the Target field.
func (o *CommonDataTieringTaskResponse) SetTarget(v DataTieringTarget) {
	o.Target.Set(&v)
}
// SetTargetNil sets the value for Target to be an explicit nil
func (o *CommonDataTieringTaskResponse) SetTargetNil() {
	o.Target.Set(nil)
}

// UnsetTarget ensures that no value is present for Target, not even an explicit nil
func (o *CommonDataTieringTaskResponse) UnsetTarget() {
	o.Target.Unset()
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CommonDataTieringTaskResponse) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}

	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonDataTieringTaskResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// SetType sets field value
func (o *CommonDataTieringTaskResponse) SetType(v string) {
	o.Type.Set(&v)
}

func (o CommonDataTieringTaskResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonDataTieringTaskResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AlertPolicy) {
		toSerialize["alertPolicy"] = o.AlertPolicy
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.IsActive.IsSet() {
		toSerialize["isActive"] = o.IsActive.Get()
	}
	if o.IsDeleted.IsSet() {
		toSerialize["isDeleted"] = o.IsDeleted.Get()
	}
	if o.IsPaused.IsSet() {
		toSerialize["isPaused"] = o.IsPaused.Get()
	}
	if !IsNil(o.LastRun) {
		toSerialize["lastRun"] = o.LastRun
	}
	toSerialize["name"] = o.Name.Get()
	if !IsNil(o.Schedule) {
		toSerialize["schedule"] = o.Schedule
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if o.Target.IsSet() {
		toSerialize["target"] = o.Target.Get()
	}
	toSerialize["type"] = o.Type.Get()
	return toSerialize, nil
}

func (o *CommonDataTieringTaskResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCommonDataTieringTaskResponse := _CommonDataTieringTaskResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCommonDataTieringTaskResponse)

	if err != nil {
		return err
	}

	*o = CommonDataTieringTaskResponse(varCommonDataTieringTaskResponse)

	return err
}

type NullableCommonDataTieringTaskResponse struct {
	value *CommonDataTieringTaskResponse
	isSet bool
}

func (v NullableCommonDataTieringTaskResponse) Get() *CommonDataTieringTaskResponse {
	return v.value
}

func (v *NullableCommonDataTieringTaskResponse) Set(val *CommonDataTieringTaskResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonDataTieringTaskResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonDataTieringTaskResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonDataTieringTaskResponse(val *CommonDataTieringTaskResponse) *NullableCommonDataTieringTaskResponse {
	return &NullableCommonDataTieringTaskResponse{value: val, isSet: true}
}

func (v NullableCommonDataTieringTaskResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonDataTieringTaskResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


