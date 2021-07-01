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

// MapReduceInstance Information about a Map reduce instance. An instance can be run only once.
type MapReduceInstance struct {
	// System generated ID of map reduce instance.
	Id NullableInt64 `json:"id,omitempty"`
	InputParams []MapReduceInstanceInputParam `json:"inputParams,omitempty"`
	InputSpec *InputSpec `json:"inputSpec,omitempty"`
	// ID of Map reduce info.
	MapReduceInfoId NullableInt64 `json:"mapReduceInfoId,omitempty"`
	OutputSpec *OutputSpec `json:"outputSpec,omitempty"`
	RunInfo *MapReduceInstanceRunInfo `json:"runInfo,omitempty"`
}

// NewMapReduceInstance instantiates a new MapReduceInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMapReduceInstance() *MapReduceInstance {
	this := MapReduceInstance{}
	return &this
}

// NewMapReduceInstanceWithDefaults instantiates a new MapReduceInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMapReduceInstanceWithDefaults() *MapReduceInstance {
	this := MapReduceInstance{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MapReduceInstance) GetId() int64 {
	if o == nil || o.Id.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MapReduceInstance) GetIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *MapReduceInstance) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *MapReduceInstance) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *MapReduceInstance) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *MapReduceInstance) UnsetId() {
	o.Id.Unset()
}

// GetInputParams returns the InputParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MapReduceInstance) GetInputParams() []MapReduceInstanceInputParam {
	if o == nil  {
		var ret []MapReduceInstanceInputParam
		return ret
	}
	return o.InputParams
}

// GetInputParamsOk returns a tuple with the InputParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MapReduceInstance) GetInputParamsOk() (*[]MapReduceInstanceInputParam, bool) {
	if o == nil || o.InputParams == nil {
		return nil, false
	}
	return &o.InputParams, true
}

// HasInputParams returns a boolean if a field has been set.
func (o *MapReduceInstance) HasInputParams() bool {
	if o != nil && o.InputParams != nil {
		return true
	}

	return false
}

// SetInputParams gets a reference to the given []MapReduceInstanceInputParam and assigns it to the InputParams field.
func (o *MapReduceInstance) SetInputParams(v []MapReduceInstanceInputParam) {
	o.InputParams = v
}

// GetInputSpec returns the InputSpec field value if set, zero value otherwise.
func (o *MapReduceInstance) GetInputSpec() InputSpec {
	if o == nil || o.InputSpec == nil {
		var ret InputSpec
		return ret
	}
	return *o.InputSpec
}

// GetInputSpecOk returns a tuple with the InputSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MapReduceInstance) GetInputSpecOk() (*InputSpec, bool) {
	if o == nil || o.InputSpec == nil {
		return nil, false
	}
	return o.InputSpec, true
}

// HasInputSpec returns a boolean if a field has been set.
func (o *MapReduceInstance) HasInputSpec() bool {
	if o != nil && o.InputSpec != nil {
		return true
	}

	return false
}

// SetInputSpec gets a reference to the given InputSpec and assigns it to the InputSpec field.
func (o *MapReduceInstance) SetInputSpec(v InputSpec) {
	o.InputSpec = &v
}

// GetMapReduceInfoId returns the MapReduceInfoId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MapReduceInstance) GetMapReduceInfoId() int64 {
	if o == nil || o.MapReduceInfoId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.MapReduceInfoId.Get()
}

// GetMapReduceInfoIdOk returns a tuple with the MapReduceInfoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MapReduceInstance) GetMapReduceInfoIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MapReduceInfoId.Get(), o.MapReduceInfoId.IsSet()
}

// HasMapReduceInfoId returns a boolean if a field has been set.
func (o *MapReduceInstance) HasMapReduceInfoId() bool {
	if o != nil && o.MapReduceInfoId.IsSet() {
		return true
	}

	return false
}

// SetMapReduceInfoId gets a reference to the given NullableInt64 and assigns it to the MapReduceInfoId field.
func (o *MapReduceInstance) SetMapReduceInfoId(v int64) {
	o.MapReduceInfoId.Set(&v)
}
// SetMapReduceInfoIdNil sets the value for MapReduceInfoId to be an explicit nil
func (o *MapReduceInstance) SetMapReduceInfoIdNil() {
	o.MapReduceInfoId.Set(nil)
}

// UnsetMapReduceInfoId ensures that no value is present for MapReduceInfoId, not even an explicit nil
func (o *MapReduceInstance) UnsetMapReduceInfoId() {
	o.MapReduceInfoId.Unset()
}

// GetOutputSpec returns the OutputSpec field value if set, zero value otherwise.
func (o *MapReduceInstance) GetOutputSpec() OutputSpec {
	if o == nil || o.OutputSpec == nil {
		var ret OutputSpec
		return ret
	}
	return *o.OutputSpec
}

// GetOutputSpecOk returns a tuple with the OutputSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MapReduceInstance) GetOutputSpecOk() (*OutputSpec, bool) {
	if o == nil || o.OutputSpec == nil {
		return nil, false
	}
	return o.OutputSpec, true
}

// HasOutputSpec returns a boolean if a field has been set.
func (o *MapReduceInstance) HasOutputSpec() bool {
	if o != nil && o.OutputSpec != nil {
		return true
	}

	return false
}

// SetOutputSpec gets a reference to the given OutputSpec and assigns it to the OutputSpec field.
func (o *MapReduceInstance) SetOutputSpec(v OutputSpec) {
	o.OutputSpec = &v
}

// GetRunInfo returns the RunInfo field value if set, zero value otherwise.
func (o *MapReduceInstance) GetRunInfo() MapReduceInstanceRunInfo {
	if o == nil || o.RunInfo == nil {
		var ret MapReduceInstanceRunInfo
		return ret
	}
	return *o.RunInfo
}

// GetRunInfoOk returns a tuple with the RunInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MapReduceInstance) GetRunInfoOk() (*MapReduceInstanceRunInfo, bool) {
	if o == nil || o.RunInfo == nil {
		return nil, false
	}
	return o.RunInfo, true
}

// HasRunInfo returns a boolean if a field has been set.
func (o *MapReduceInstance) HasRunInfo() bool {
	if o != nil && o.RunInfo != nil {
		return true
	}

	return false
}

// SetRunInfo gets a reference to the given MapReduceInstanceRunInfo and assigns it to the RunInfo field.
func (o *MapReduceInstance) SetRunInfo(v MapReduceInstanceRunInfo) {
	o.RunInfo = &v
}

func (o MapReduceInstance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.InputParams != nil {
		toSerialize["inputParams"] = o.InputParams
	}
	if o.InputSpec != nil {
		toSerialize["inputSpec"] = o.InputSpec
	}
	if o.MapReduceInfoId.IsSet() {
		toSerialize["mapReduceInfoId"] = o.MapReduceInfoId.Get()
	}
	if o.OutputSpec != nil {
		toSerialize["outputSpec"] = o.OutputSpec
	}
	if o.RunInfo != nil {
		toSerialize["runInfo"] = o.RunInfo
	}
	return json.Marshal(toSerialize)
}

type NullableMapReduceInstance struct {
	value *MapReduceInstance
	isSet bool
}

func (v NullableMapReduceInstance) Get() *MapReduceInstance {
	return v.value
}

func (v *NullableMapReduceInstance) Set(val *MapReduceInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableMapReduceInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableMapReduceInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMapReduceInstance(val *MapReduceInstance) *NullableMapReduceInstance {
	return &NullableMapReduceInstance{value: val, isSet: true}
}

func (v NullableMapReduceInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMapReduceInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


