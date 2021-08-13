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

// ThrottlingParams Specifies throttling params.
type ThrottlingParams struct {
	// If the latency of a datastore is above this value, then a new backup task that uses the datastore won't be started.
	NewTaskLatencyThresholdMsecs NullableInt64 `json:"newTaskLatencyThresholdMsecs,omitempty"`
	// If the latency of a datastore is above this value, then an existing backup task that uses the datastore will start getting throttled.
	ActiveTaskLatencyThresholdMsecs NullableInt64 `json:"activeTaskLatencyThresholdMsecs,omitempty"`
	// If this value is > 0 and the number of streams concurrently active on a datastore is equal to it, then any further requests to access the datastore would be denied until the number of active streams reduces. This applies for all the datastores in the specified host.
	MaxConcurrentStreams NullableInt32 `json:"maxConcurrentStreams,omitempty"`
	// Specifies datastore specific parameters.
	DataStoreParams []DatastoreParams `json:"dataStoreParams,omitempty"`
}

// NewThrottlingParams instantiates a new ThrottlingParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThrottlingParams() *ThrottlingParams {
	this := ThrottlingParams{}
	return &this
}

// NewThrottlingParamsWithDefaults instantiates a new ThrottlingParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThrottlingParamsWithDefaults() *ThrottlingParams {
	this := ThrottlingParams{}
	return &this
}

// GetNewTaskLatencyThresholdMsecs returns the NewTaskLatencyThresholdMsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThrottlingParams) GetNewTaskLatencyThresholdMsecs() int64 {
	if o == nil || o.NewTaskLatencyThresholdMsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.NewTaskLatencyThresholdMsecs.Get()
}

// GetNewTaskLatencyThresholdMsecsOk returns a tuple with the NewTaskLatencyThresholdMsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThrottlingParams) GetNewTaskLatencyThresholdMsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NewTaskLatencyThresholdMsecs.Get(), o.NewTaskLatencyThresholdMsecs.IsSet()
}

// HasNewTaskLatencyThresholdMsecs returns a boolean if a field has been set.
func (o *ThrottlingParams) HasNewTaskLatencyThresholdMsecs() bool {
	if o != nil && o.NewTaskLatencyThresholdMsecs.IsSet() {
		return true
	}

	return false
}

// SetNewTaskLatencyThresholdMsecs gets a reference to the given NullableInt64 and assigns it to the NewTaskLatencyThresholdMsecs field.
func (o *ThrottlingParams) SetNewTaskLatencyThresholdMsecs(v int64) {
	o.NewTaskLatencyThresholdMsecs.Set(&v)
}
// SetNewTaskLatencyThresholdMsecsNil sets the value for NewTaskLatencyThresholdMsecs to be an explicit nil
func (o *ThrottlingParams) SetNewTaskLatencyThresholdMsecsNil() {
	o.NewTaskLatencyThresholdMsecs.Set(nil)
}

// UnsetNewTaskLatencyThresholdMsecs ensures that no value is present for NewTaskLatencyThresholdMsecs, not even an explicit nil
func (o *ThrottlingParams) UnsetNewTaskLatencyThresholdMsecs() {
	o.NewTaskLatencyThresholdMsecs.Unset()
}

// GetActiveTaskLatencyThresholdMsecs returns the ActiveTaskLatencyThresholdMsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThrottlingParams) GetActiveTaskLatencyThresholdMsecs() int64 {
	if o == nil || o.ActiveTaskLatencyThresholdMsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ActiveTaskLatencyThresholdMsecs.Get()
}

// GetActiveTaskLatencyThresholdMsecsOk returns a tuple with the ActiveTaskLatencyThresholdMsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThrottlingParams) GetActiveTaskLatencyThresholdMsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ActiveTaskLatencyThresholdMsecs.Get(), o.ActiveTaskLatencyThresholdMsecs.IsSet()
}

// HasActiveTaskLatencyThresholdMsecs returns a boolean if a field has been set.
func (o *ThrottlingParams) HasActiveTaskLatencyThresholdMsecs() bool {
	if o != nil && o.ActiveTaskLatencyThresholdMsecs.IsSet() {
		return true
	}

	return false
}

// SetActiveTaskLatencyThresholdMsecs gets a reference to the given NullableInt64 and assigns it to the ActiveTaskLatencyThresholdMsecs field.
func (o *ThrottlingParams) SetActiveTaskLatencyThresholdMsecs(v int64) {
	o.ActiveTaskLatencyThresholdMsecs.Set(&v)
}
// SetActiveTaskLatencyThresholdMsecsNil sets the value for ActiveTaskLatencyThresholdMsecs to be an explicit nil
func (o *ThrottlingParams) SetActiveTaskLatencyThresholdMsecsNil() {
	o.ActiveTaskLatencyThresholdMsecs.Set(nil)
}

// UnsetActiveTaskLatencyThresholdMsecs ensures that no value is present for ActiveTaskLatencyThresholdMsecs, not even an explicit nil
func (o *ThrottlingParams) UnsetActiveTaskLatencyThresholdMsecs() {
	o.ActiveTaskLatencyThresholdMsecs.Unset()
}

// GetMaxConcurrentStreams returns the MaxConcurrentStreams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThrottlingParams) GetMaxConcurrentStreams() int32 {
	if o == nil || o.MaxConcurrentStreams.Get() == nil {
		var ret int32
		return ret
	}
	return *o.MaxConcurrentStreams.Get()
}

// GetMaxConcurrentStreamsOk returns a tuple with the MaxConcurrentStreams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThrottlingParams) GetMaxConcurrentStreamsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MaxConcurrentStreams.Get(), o.MaxConcurrentStreams.IsSet()
}

// HasMaxConcurrentStreams returns a boolean if a field has been set.
func (o *ThrottlingParams) HasMaxConcurrentStreams() bool {
	if o != nil && o.MaxConcurrentStreams.IsSet() {
		return true
	}

	return false
}

// SetMaxConcurrentStreams gets a reference to the given NullableInt32 and assigns it to the MaxConcurrentStreams field.
func (o *ThrottlingParams) SetMaxConcurrentStreams(v int32) {
	o.MaxConcurrentStreams.Set(&v)
}
// SetMaxConcurrentStreamsNil sets the value for MaxConcurrentStreams to be an explicit nil
func (o *ThrottlingParams) SetMaxConcurrentStreamsNil() {
	o.MaxConcurrentStreams.Set(nil)
}

// UnsetMaxConcurrentStreams ensures that no value is present for MaxConcurrentStreams, not even an explicit nil
func (o *ThrottlingParams) UnsetMaxConcurrentStreams() {
	o.MaxConcurrentStreams.Unset()
}

// GetDataStoreParams returns the DataStoreParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThrottlingParams) GetDataStoreParams() []DatastoreParams {
	if o == nil  {
		var ret []DatastoreParams
		return ret
	}
	return o.DataStoreParams
}

// GetDataStoreParamsOk returns a tuple with the DataStoreParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThrottlingParams) GetDataStoreParamsOk() (*[]DatastoreParams, bool) {
	if o == nil || o.DataStoreParams == nil {
		return nil, false
	}
	return &o.DataStoreParams, true
}

// HasDataStoreParams returns a boolean if a field has been set.
func (o *ThrottlingParams) HasDataStoreParams() bool {
	if o != nil && o.DataStoreParams != nil {
		return true
	}

	return false
}

// SetDataStoreParams gets a reference to the given []DatastoreParams and assigns it to the DataStoreParams field.
func (o *ThrottlingParams) SetDataStoreParams(v []DatastoreParams) {
	o.DataStoreParams = v
}

func (o ThrottlingParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NewTaskLatencyThresholdMsecs.IsSet() {
		toSerialize["newTaskLatencyThresholdMsecs"] = o.NewTaskLatencyThresholdMsecs.Get()
	}
	if o.ActiveTaskLatencyThresholdMsecs.IsSet() {
		toSerialize["activeTaskLatencyThresholdMsecs"] = o.ActiveTaskLatencyThresholdMsecs.Get()
	}
	if o.MaxConcurrentStreams.IsSet() {
		toSerialize["maxConcurrentStreams"] = o.MaxConcurrentStreams.Get()
	}
	if o.DataStoreParams != nil {
		toSerialize["dataStoreParams"] = o.DataStoreParams
	}
	return json.Marshal(toSerialize)
}

type NullableThrottlingParams struct {
	value *ThrottlingParams
	isSet bool
}

func (v NullableThrottlingParams) Get() *ThrottlingParams {
	return v.value
}

func (v *NullableThrottlingParams) Set(val *ThrottlingParams) {
	v.value = val
	v.isSet = true
}

func (v NullableThrottlingParams) IsSet() bool {
	return v.isSet
}

func (v *NullableThrottlingParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThrottlingParams(val *ThrottlingParams) *NullableThrottlingParams {
	return &NullableThrottlingParams{value: val, isSet: true}
}

func (v NullableThrottlingParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThrottlingParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


