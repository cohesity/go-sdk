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

// DataTieringAnalysisGroupRun Specifies the data tiering analysis group run.
type DataTieringAnalysisGroupRun struct {
	// Specifies the ID of the data tiering analysis group run.
	Id NullableString `json:"id,omitempty"`
	// Specifies the objects details analyzed during data tiering analysis group run.
	Objects []DataTieringObjectInfo `json:"objects,omitempty"`
	// Specifies the start time of analysis run in Unix epoch Timestamp(in microseconds).
	StartTimeUsecs NullableInt64 `json:"startTimeUsecs,omitempty"`
	// Specifies the end time of analysis run in Unix epoch Timestamp(in microseconds).
	EndTimeUsecs NullableInt64 `json:"endTimeUsecs,omitempty"`
	AnalysisInfo *DataTieringAnalysisInfo `json:"analysisInfo,omitempty"`
}

// NewDataTieringAnalysisGroupRun instantiates a new DataTieringAnalysisGroupRun object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataTieringAnalysisGroupRun() *DataTieringAnalysisGroupRun {
	this := DataTieringAnalysisGroupRun{}
	return &this
}

// NewDataTieringAnalysisGroupRunWithDefaults instantiates a new DataTieringAnalysisGroupRun object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataTieringAnalysisGroupRunWithDefaults() *DataTieringAnalysisGroupRun {
	this := DataTieringAnalysisGroupRun{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataTieringAnalysisGroupRun) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataTieringAnalysisGroupRun) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *DataTieringAnalysisGroupRun) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *DataTieringAnalysisGroupRun) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *DataTieringAnalysisGroupRun) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *DataTieringAnalysisGroupRun) UnsetId() {
	o.Id.Unset()
}

// GetObjects returns the Objects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataTieringAnalysisGroupRun) GetObjects() []DataTieringObjectInfo {
	if o == nil  {
		var ret []DataTieringObjectInfo
		return ret
	}
	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataTieringAnalysisGroupRun) GetObjectsOk() (*[]DataTieringObjectInfo, bool) {
	if o == nil || o.Objects == nil {
		return nil, false
	}
	return &o.Objects, true
}

// HasObjects returns a boolean if a field has been set.
func (o *DataTieringAnalysisGroupRun) HasObjects() bool {
	if o != nil && o.Objects != nil {
		return true
	}

	return false
}

// SetObjects gets a reference to the given []DataTieringObjectInfo and assigns it to the Objects field.
func (o *DataTieringAnalysisGroupRun) SetObjects(v []DataTieringObjectInfo) {
	o.Objects = v
}

// GetStartTimeUsecs returns the StartTimeUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataTieringAnalysisGroupRun) GetStartTimeUsecs() int64 {
	if o == nil || o.StartTimeUsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.StartTimeUsecs.Get()
}

// GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataTieringAnalysisGroupRun) GetStartTimeUsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartTimeUsecs.Get(), o.StartTimeUsecs.IsSet()
}

// HasStartTimeUsecs returns a boolean if a field has been set.
func (o *DataTieringAnalysisGroupRun) HasStartTimeUsecs() bool {
	if o != nil && o.StartTimeUsecs.IsSet() {
		return true
	}

	return false
}

// SetStartTimeUsecs gets a reference to the given NullableInt64 and assigns it to the StartTimeUsecs field.
func (o *DataTieringAnalysisGroupRun) SetStartTimeUsecs(v int64) {
	o.StartTimeUsecs.Set(&v)
}
// SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil
func (o *DataTieringAnalysisGroupRun) SetStartTimeUsecsNil() {
	o.StartTimeUsecs.Set(nil)
}

// UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
func (o *DataTieringAnalysisGroupRun) UnsetStartTimeUsecs() {
	o.StartTimeUsecs.Unset()
}

// GetEndTimeUsecs returns the EndTimeUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataTieringAnalysisGroupRun) GetEndTimeUsecs() int64 {
	if o == nil || o.EndTimeUsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.EndTimeUsecs.Get()
}

// GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataTieringAnalysisGroupRun) GetEndTimeUsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndTimeUsecs.Get(), o.EndTimeUsecs.IsSet()
}

// HasEndTimeUsecs returns a boolean if a field has been set.
func (o *DataTieringAnalysisGroupRun) HasEndTimeUsecs() bool {
	if o != nil && o.EndTimeUsecs.IsSet() {
		return true
	}

	return false
}

// SetEndTimeUsecs gets a reference to the given NullableInt64 and assigns it to the EndTimeUsecs field.
func (o *DataTieringAnalysisGroupRun) SetEndTimeUsecs(v int64) {
	o.EndTimeUsecs.Set(&v)
}
// SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil
func (o *DataTieringAnalysisGroupRun) SetEndTimeUsecsNil() {
	o.EndTimeUsecs.Set(nil)
}

// UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
func (o *DataTieringAnalysisGroupRun) UnsetEndTimeUsecs() {
	o.EndTimeUsecs.Unset()
}

// GetAnalysisInfo returns the AnalysisInfo field value if set, zero value otherwise.
func (o *DataTieringAnalysisGroupRun) GetAnalysisInfo() DataTieringAnalysisInfo {
	if o == nil || o.AnalysisInfo == nil {
		var ret DataTieringAnalysisInfo
		return ret
	}
	return *o.AnalysisInfo
}

// GetAnalysisInfoOk returns a tuple with the AnalysisInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataTieringAnalysisGroupRun) GetAnalysisInfoOk() (*DataTieringAnalysisInfo, bool) {
	if o == nil || o.AnalysisInfo == nil {
		return nil, false
	}
	return o.AnalysisInfo, true
}

// HasAnalysisInfo returns a boolean if a field has been set.
func (o *DataTieringAnalysisGroupRun) HasAnalysisInfo() bool {
	if o != nil && o.AnalysisInfo != nil {
		return true
	}

	return false
}

// SetAnalysisInfo gets a reference to the given DataTieringAnalysisInfo and assigns it to the AnalysisInfo field.
func (o *DataTieringAnalysisGroupRun) SetAnalysisInfo(v DataTieringAnalysisInfo) {
	o.AnalysisInfo = &v
}

func (o DataTieringAnalysisGroupRun) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Objects != nil {
		toSerialize["objects"] = o.Objects
	}
	if o.StartTimeUsecs.IsSet() {
		toSerialize["startTimeUsecs"] = o.StartTimeUsecs.Get()
	}
	if o.EndTimeUsecs.IsSet() {
		toSerialize["endTimeUsecs"] = o.EndTimeUsecs.Get()
	}
	if o.AnalysisInfo != nil {
		toSerialize["analysisInfo"] = o.AnalysisInfo
	}
	return json.Marshal(toSerialize)
}

type NullableDataTieringAnalysisGroupRun struct {
	value *DataTieringAnalysisGroupRun
	isSet bool
}

func (v NullableDataTieringAnalysisGroupRun) Get() *DataTieringAnalysisGroupRun {
	return v.value
}

func (v *NullableDataTieringAnalysisGroupRun) Set(val *DataTieringAnalysisGroupRun) {
	v.value = val
	v.isSet = true
}

func (v NullableDataTieringAnalysisGroupRun) IsSet() bool {
	return v.isSet
}

func (v *NullableDataTieringAnalysisGroupRun) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataTieringAnalysisGroupRun(val *DataTieringAnalysisGroupRun) *NullableDataTieringAnalysisGroupRun {
	return &NullableDataTieringAnalysisGroupRun{value: val, isSet: true}
}

func (v NullableDataTieringAnalysisGroupRun) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataTieringAnalysisGroupRun) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o DataTieringAnalysisGroupRun) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}