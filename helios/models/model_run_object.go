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

// RunObject Specifies the object details to create a protection run.
type RunObject struct {
	// Specifies the id of object.
	Id NullableInt64 `json:"id"`
	// Specifies a list of ids of applications.
	AppIds []int64 `json:"appIds,omitempty"`
	PhysicalParams *RunObjectPhysicalParams `json:"physicalParams,omitempty"`
}

// NewRunObject instantiates a new RunObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunObject(id NullableInt64) *RunObject {
	this := RunObject{}
	this.Id = id
	return &this
}

// NewRunObjectWithDefaults instantiates a new RunObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunObjectWithDefaults() *RunObject {
	this := RunObject{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *RunObject) GetId() int64 {
	if o == nil || o.Id.Get() == nil {
		var ret int64
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RunObject) GetIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *RunObject) SetId(v int64) {
	o.Id.Set(&v)
}

// GetAppIds returns the AppIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RunObject) GetAppIds() []int64 {
	if o == nil  {
		var ret []int64
		return ret
	}
	return o.AppIds
}

// GetAppIdsOk returns a tuple with the AppIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RunObject) GetAppIdsOk() (*[]int64, bool) {
	if o == nil || o.AppIds == nil {
		return nil, false
	}
	return &o.AppIds, true
}

// HasAppIds returns a boolean if a field has been set.
func (o *RunObject) HasAppIds() bool {
	if o != nil && o.AppIds != nil {
		return true
	}

	return false
}

// SetAppIds gets a reference to the given []int64 and assigns it to the AppIds field.
func (o *RunObject) SetAppIds(v []int64) {
	o.AppIds = v
}

// GetPhysicalParams returns the PhysicalParams field value if set, zero value otherwise.
func (o *RunObject) GetPhysicalParams() RunObjectPhysicalParams {
	if o == nil || o.PhysicalParams == nil {
		var ret RunObjectPhysicalParams
		return ret
	}
	return *o.PhysicalParams
}

// GetPhysicalParamsOk returns a tuple with the PhysicalParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunObject) GetPhysicalParamsOk() (*RunObjectPhysicalParams, bool) {
	if o == nil || o.PhysicalParams == nil {
		return nil, false
	}
	return o.PhysicalParams, true
}

// HasPhysicalParams returns a boolean if a field has been set.
func (o *RunObject) HasPhysicalParams() bool {
	if o != nil && o.PhysicalParams != nil {
		return true
	}

	return false
}

// SetPhysicalParams gets a reference to the given RunObjectPhysicalParams and assigns it to the PhysicalParams field.
func (o *RunObject) SetPhysicalParams(v RunObjectPhysicalParams) {
	o.PhysicalParams = &v
}

func (o RunObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id.Get()
	}
	if o.AppIds != nil {
		toSerialize["appIds"] = o.AppIds
	}
	if o.PhysicalParams != nil {
		toSerialize["physicalParams"] = o.PhysicalParams
	}
	return json.Marshal(toSerialize)
}

type NullableRunObject struct {
	value *RunObject
	isSet bool
}

func (v NullableRunObject) Get() *RunObject {
	return v.value
}

func (v *NullableRunObject) Set(val *RunObject) {
	v.value = val
	v.isSet = true
}

func (v NullableRunObject) IsSet() bool {
	return v.isSet
}

func (v *NullableRunObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunObject(val *RunObject) *NullableRunObject {
	return &NullableRunObject{value: val, isSet: true}
}

func (v NullableRunObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


