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

// BackupJobProtoExclusionTimeRange A proto to specify a time range within a single day when backups are not permitted to run.
type BackupJobProtoExclusionTimeRange struct {
	// If the day is not set, the time range applies to all days.
	Day NullableInt32 `json:"day,omitempty"`
	EndTime *Time `json:"endTime,omitempty"`
	StartTime *Time `json:"startTime,omitempty"`
}

// NewBackupJobProtoExclusionTimeRange instantiates a new BackupJobProtoExclusionTimeRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupJobProtoExclusionTimeRange() *BackupJobProtoExclusionTimeRange {
	this := BackupJobProtoExclusionTimeRange{}
	return &this
}

// NewBackupJobProtoExclusionTimeRangeWithDefaults instantiates a new BackupJobProtoExclusionTimeRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupJobProtoExclusionTimeRangeWithDefaults() *BackupJobProtoExclusionTimeRange {
	this := BackupJobProtoExclusionTimeRange{}
	return &this
}

// GetDay returns the Day field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupJobProtoExclusionTimeRange) GetDay() int32 {
	if o == nil || o.Day.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Day.Get()
}

// GetDayOk returns a tuple with the Day field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupJobProtoExclusionTimeRange) GetDayOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Day.Get(), o.Day.IsSet()
}

// HasDay returns a boolean if a field has been set.
func (o *BackupJobProtoExclusionTimeRange) HasDay() bool {
	if o != nil && o.Day.IsSet() {
		return true
	}

	return false
}

// SetDay gets a reference to the given NullableInt32 and assigns it to the Day field.
func (o *BackupJobProtoExclusionTimeRange) SetDay(v int32) {
	o.Day.Set(&v)
}
// SetDayNil sets the value for Day to be an explicit nil
func (o *BackupJobProtoExclusionTimeRange) SetDayNil() {
	o.Day.Set(nil)
}

// UnsetDay ensures that no value is present for Day, not even an explicit nil
func (o *BackupJobProtoExclusionTimeRange) UnsetDay() {
	o.Day.Unset()
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *BackupJobProtoExclusionTimeRange) GetEndTime() Time {
	if o == nil || o.EndTime == nil {
		var ret Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupJobProtoExclusionTimeRange) GetEndTimeOk() (*Time, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *BackupJobProtoExclusionTimeRange) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given Time and assigns it to the EndTime field.
func (o *BackupJobProtoExclusionTimeRange) SetEndTime(v Time) {
	o.EndTime = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *BackupJobProtoExclusionTimeRange) GetStartTime() Time {
	if o == nil || o.StartTime == nil {
		var ret Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupJobProtoExclusionTimeRange) GetStartTimeOk() (*Time, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *BackupJobProtoExclusionTimeRange) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given Time and assigns it to the StartTime field.
func (o *BackupJobProtoExclusionTimeRange) SetStartTime(v Time) {
	o.StartTime = &v
}

func (o BackupJobProtoExclusionTimeRange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Day.IsSet() {
		toSerialize["day"] = o.Day.Get()
	}
	if o.EndTime != nil {
		toSerialize["endTime"] = o.EndTime
	}
	if o.StartTime != nil {
		toSerialize["startTime"] = o.StartTime
	}
	return json.Marshal(toSerialize)
}

type NullableBackupJobProtoExclusionTimeRange struct {
	value *BackupJobProtoExclusionTimeRange
	isSet bool
}

func (v NullableBackupJobProtoExclusionTimeRange) Get() *BackupJobProtoExclusionTimeRange {
	return v.value
}

func (v *NullableBackupJobProtoExclusionTimeRange) Set(val *BackupJobProtoExclusionTimeRange) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupJobProtoExclusionTimeRange) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupJobProtoExclusionTimeRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupJobProtoExclusionTimeRange(val *BackupJobProtoExclusionTimeRange) *NullableBackupJobProtoExclusionTimeRange {
	return &NullableBackupJobProtoExclusionTimeRange{value: val, isSet: true}
}

func (v NullableBackupJobProtoExclusionTimeRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupJobProtoExclusionTimeRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


