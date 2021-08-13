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

// GetViewsResult Specifies the list of Views returned that matched the specified filter criteria.
type GetViewsResult struct {
	// Array of Views. Specifies the list of Views returned in this response. The list is sorted by decreasing View ids.
	Views []View `json:"views,omitempty"`
	// If false, more Views are available to return. If the number of Views to return exceeds the number of Views specified in maxCount (default of 1000) of the original GET /public/views request, the first set of Views are returned and this field returns false. To get the next set of Views, in the next GET /public/views request send the last id from the previous viewList.
	LastResult NullableBool `json:"lastResult,omitempty"`
	// Number of views returned. This will only be returned if ViewCountOnly is set in arguments.
	Count NullableInt64 `json:"count,omitempty"`
}

// NewGetViewsResult instantiates a new GetViewsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetViewsResult() *GetViewsResult {
	this := GetViewsResult{}
	return &this
}

// NewGetViewsResultWithDefaults instantiates a new GetViewsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetViewsResultWithDefaults() *GetViewsResult {
	this := GetViewsResult{}
	return &this
}

// GetViews returns the Views field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetViewsResult) GetViews() []View {
	if o == nil  {
		var ret []View
		return ret
	}
	return o.Views
}

// GetViewsOk returns a tuple with the Views field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetViewsResult) GetViewsOk() (*[]View, bool) {
	if o == nil || o.Views == nil {
		return nil, false
	}
	return &o.Views, true
}

// HasViews returns a boolean if a field has been set.
func (o *GetViewsResult) HasViews() bool {
	if o != nil && o.Views != nil {
		return true
	}

	return false
}

// SetViews gets a reference to the given []View and assigns it to the Views field.
func (o *GetViewsResult) SetViews(v []View) {
	o.Views = v
}

// GetLastResult returns the LastResult field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetViewsResult) GetLastResult() bool {
	if o == nil || o.LastResult.Get() == nil {
		var ret bool
		return ret
	}
	return *o.LastResult.Get()
}

// GetLastResultOk returns a tuple with the LastResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetViewsResult) GetLastResultOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastResult.Get(), o.LastResult.IsSet()
}

// HasLastResult returns a boolean if a field has been set.
func (o *GetViewsResult) HasLastResult() bool {
	if o != nil && o.LastResult.IsSet() {
		return true
	}

	return false
}

// SetLastResult gets a reference to the given NullableBool and assigns it to the LastResult field.
func (o *GetViewsResult) SetLastResult(v bool) {
	o.LastResult.Set(&v)
}
// SetLastResultNil sets the value for LastResult to be an explicit nil
func (o *GetViewsResult) SetLastResultNil() {
	o.LastResult.Set(nil)
}

// UnsetLastResult ensures that no value is present for LastResult, not even an explicit nil
func (o *GetViewsResult) UnsetLastResult() {
	o.LastResult.Unset()
}

// GetCount returns the Count field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetViewsResult) GetCount() int64 {
	if o == nil || o.Count.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Count.Get()
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetViewsResult) GetCountOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Count.Get(), o.Count.IsSet()
}

// HasCount returns a boolean if a field has been set.
func (o *GetViewsResult) HasCount() bool {
	if o != nil && o.Count.IsSet() {
		return true
	}

	return false
}

// SetCount gets a reference to the given NullableInt64 and assigns it to the Count field.
func (o *GetViewsResult) SetCount(v int64) {
	o.Count.Set(&v)
}
// SetCountNil sets the value for Count to be an explicit nil
func (o *GetViewsResult) SetCountNil() {
	o.Count.Set(nil)
}

// UnsetCount ensures that no value is present for Count, not even an explicit nil
func (o *GetViewsResult) UnsetCount() {
	o.Count.Unset()
}

func (o GetViewsResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Views != nil {
		toSerialize["views"] = o.Views
	}
	if o.LastResult.IsSet() {
		toSerialize["lastResult"] = o.LastResult.Get()
	}
	if o.Count.IsSet() {
		toSerialize["count"] = o.Count.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableGetViewsResult struct {
	value *GetViewsResult
	isSet bool
}

func (v NullableGetViewsResult) Get() *GetViewsResult {
	return v.value
}

func (v *NullableGetViewsResult) Set(val *GetViewsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetViewsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetViewsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetViewsResult(val *GetViewsResult) *NullableGetViewsResult {
	return &NullableGetViewsResult{value: val, isSet: true}
}

func (v NullableGetViewsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetViewsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


