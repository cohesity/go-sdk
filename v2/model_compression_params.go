/*
Cohesity REST API

Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
)

// checks if the CompressionParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompressionParams{}

// CompressionParams Specifies parameters for compression.
type CompressionParams struct {
	// Specifies whether inline compression is enabled. This field is appliciable only if inlineDeduplicationEnabled is set to true and compression is enabled.
	InlineEnabled NullableBool `json:"inlineEnabled,omitempty"`
	// Specifies copmpression type for a Storage Domain.
	Type NullableString `json:"type,omitempty"`
}

// NewCompressionParams instantiates a new CompressionParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompressionParams() *CompressionParams {
	this := CompressionParams{}
	return &this
}

// NewCompressionParamsWithDefaults instantiates a new CompressionParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompressionParamsWithDefaults() *CompressionParams {
	this := CompressionParams{}
	return &this
}

// GetInlineEnabled returns the InlineEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompressionParams) GetInlineEnabled() bool {
	if o == nil || IsNil(o.InlineEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.InlineEnabled.Get()
}

// GetInlineEnabledOk returns a tuple with the InlineEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompressionParams) GetInlineEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.InlineEnabled.Get(), o.InlineEnabled.IsSet()
}

// HasInlineEnabled returns a boolean if a field has been set.
func (o *CompressionParams) HasInlineEnabled() bool {
	if o != nil && o.InlineEnabled.IsSet() {
		return true
	}

	return false
}

// SetInlineEnabled gets a reference to the given NullableBool and assigns it to the InlineEnabled field.
func (o *CompressionParams) SetInlineEnabled(v bool) {
	o.InlineEnabled.Set(&v)
}
// SetInlineEnabledNil sets the value for InlineEnabled to be an explicit nil
func (o *CompressionParams) SetInlineEnabledNil() {
	o.InlineEnabled.Set(nil)
}

// UnsetInlineEnabled ensures that no value is present for InlineEnabled, not even an explicit nil
func (o *CompressionParams) UnsetInlineEnabled() {
	o.InlineEnabled.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompressionParams) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompressionParams) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *CompressionParams) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *CompressionParams) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *CompressionParams) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *CompressionParams) UnsetType() {
	o.Type.Unset()
}

func (o CompressionParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompressionParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.InlineEnabled.IsSet() {
		toSerialize["inlineEnabled"] = o.InlineEnabled.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	return toSerialize, nil
}

type NullableCompressionParams struct {
	value *CompressionParams
	isSet bool
}

func (v NullableCompressionParams) Get() *CompressionParams {
	return v.value
}

func (v *NullableCompressionParams) Set(val *CompressionParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCompressionParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCompressionParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompressionParams(val *CompressionParams) *NullableCompressionParams {
	return &NullableCompressionParams{value: val, isSet: true}
}

func (v NullableCompressionParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompressionParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


