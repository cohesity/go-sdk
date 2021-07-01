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

// OverwriteViewParam Specifies the overwrite parameters for the view.
type OverwriteViewParam struct {
	// Specifies the source view name.
	SourceViewName NullableString `json:"sourceViewName"`
	// Specifies the target view name.
	TargetViewName NullableString `json:"targetViewName"`
}

// NewOverwriteViewParam instantiates a new OverwriteViewParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOverwriteViewParam(sourceViewName NullableString, targetViewName NullableString) *OverwriteViewParam {
	this := OverwriteViewParam{}
	this.SourceViewName = sourceViewName
	this.TargetViewName = targetViewName
	return &this
}

// NewOverwriteViewParamWithDefaults instantiates a new OverwriteViewParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOverwriteViewParamWithDefaults() *OverwriteViewParam {
	this := OverwriteViewParam{}
	return &this
}

// GetSourceViewName returns the SourceViewName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *OverwriteViewParam) GetSourceViewName() string {
	if o == nil || o.SourceViewName.Get() == nil {
		var ret string
		return ret
	}

	return *o.SourceViewName.Get()
}

// GetSourceViewNameOk returns a tuple with the SourceViewName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OverwriteViewParam) GetSourceViewNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SourceViewName.Get(), o.SourceViewName.IsSet()
}

// SetSourceViewName sets field value
func (o *OverwriteViewParam) SetSourceViewName(v string) {
	o.SourceViewName.Set(&v)
}

// GetTargetViewName returns the TargetViewName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *OverwriteViewParam) GetTargetViewName() string {
	if o == nil || o.TargetViewName.Get() == nil {
		var ret string
		return ret
	}

	return *o.TargetViewName.Get()
}

// GetTargetViewNameOk returns a tuple with the TargetViewName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OverwriteViewParam) GetTargetViewNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TargetViewName.Get(), o.TargetViewName.IsSet()
}

// SetTargetViewName sets field value
func (o *OverwriteViewParam) SetTargetViewName(v string) {
	o.TargetViewName.Set(&v)
}

func (o OverwriteViewParam) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sourceViewName"] = o.SourceViewName.Get()
	}
	if true {
		toSerialize["targetViewName"] = o.TargetViewName.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableOverwriteViewParam struct {
	value *OverwriteViewParam
	isSet bool
}

func (v NullableOverwriteViewParam) Get() *OverwriteViewParam {
	return v.value
}

func (v *NullableOverwriteViewParam) Set(val *OverwriteViewParam) {
	v.value = val
	v.isSet = true
}

func (v NullableOverwriteViewParam) IsSet() bool {
	return v.isSet
}

func (v *NullableOverwriteViewParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOverwriteViewParam(val *OverwriteViewParam) *NullableOverwriteViewParam {
	return &NullableOverwriteViewParam{value: val, isSet: true}
}

func (v NullableOverwriteViewParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOverwriteViewParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


