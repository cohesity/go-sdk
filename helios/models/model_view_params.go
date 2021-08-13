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

// ViewParams Specifies the details about a view.
type ViewParams struct {
	// Specifies the ID of the view.
	Id NullableString `json:"id,omitempty"`
	// Specifies the name of the view.
	Name NullableString `json:"name,omitempty"`
}

// NewViewParams instantiates a new ViewParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViewParams() *ViewParams {
	this := ViewParams{}
	return &this
}

// NewViewParamsWithDefaults instantiates a new ViewParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewParamsWithDefaults() *ViewParams {
	this := ViewParams{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ViewParams) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ViewParams) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *ViewParams) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *ViewParams) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *ViewParams) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *ViewParams) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ViewParams) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ViewParams) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ViewParams) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ViewParams) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ViewParams) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ViewParams) UnsetName() {
	o.Name.Unset()
}

func (o ViewParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableViewParams struct {
	value *ViewParams
	isSet bool
}

func (v NullableViewParams) Get() *ViewParams {
	return v.value
}

func (v *NullableViewParams) Set(val *ViewParams) {
	v.value = val
	v.isSet = true
}

func (v NullableViewParams) IsSet() bool {
	return v.isSet
}

func (v *NullableViewParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViewParams(val *ViewParams) *NullableViewParams {
	return &NullableViewParams{value: val, isSet: true}
}

func (v NullableViewParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViewParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


