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

// ErasureCodingParams Specifies parameters for erasure coding.
type ErasureCodingParams struct {
	// Specifies whether to enable erasure coding on a Storage Domain.
	Enabled NullableBool `json:"enabled"`
	// Specifies the algorithm used for erasure coding.
	Algorithm NullableString `json:"algorithm,omitempty"`
	// Specifies the number of data stripes.
	NumDataStripes NullableInt32 `json:"numDataStripes"`
	// Specifies the number of coded stripes.
	NumCodedStripes NullableInt32 `json:"numCodedStripes"`
	// Specifies whether inline erasure coding is enabled. This field is appliciable only if enabled is set to true.
	InlineEnabled NullableBool `json:"inlineEnabled,omitempty"`
	// Specifies the time in seconds when erasure coding starts.
	DelaySecs NullableInt32 `json:"delaySecs,omitempty"`
	// Whether to encode chunk files as part of a container. This field is appliciable only if enabled is set to true.
	UseContainer NullableBool `json:"useContainer,omitempty"`
}

// NewErasureCodingParams instantiates a new ErasureCodingParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErasureCodingParams(enabled NullableBool, numDataStripes NullableInt32, numCodedStripes NullableInt32) *ErasureCodingParams {
	this := ErasureCodingParams{}
	this.Enabled = enabled
	this.NumDataStripes = numDataStripes
	this.NumCodedStripes = numCodedStripes
	return &this
}

// NewErasureCodingParamsWithDefaults instantiates a new ErasureCodingParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErasureCodingParamsWithDefaults() *ErasureCodingParams {
	this := ErasureCodingParams{}
	return &this
}

// GetEnabled returns the Enabled field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *ErasureCodingParams) GetEnabled() bool {
	if o == nil || o.Enabled.Get() == nil {
		var ret bool
		return ret
	}

	return *o.Enabled.Get()
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ErasureCodingParams) GetEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Enabled.Get(), o.Enabled.IsSet()
}

// SetEnabled sets field value
func (o *ErasureCodingParams) SetEnabled(v bool) {
	o.Enabled.Set(&v)
}

// GetAlgorithm returns the Algorithm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ErasureCodingParams) GetAlgorithm() string {
	if o == nil || o.Algorithm.Get() == nil {
		var ret string
		return ret
	}
	return *o.Algorithm.Get()
}

// GetAlgorithmOk returns a tuple with the Algorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ErasureCodingParams) GetAlgorithmOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Algorithm.Get(), o.Algorithm.IsSet()
}

// HasAlgorithm returns a boolean if a field has been set.
func (o *ErasureCodingParams) HasAlgorithm() bool {
	if o != nil && o.Algorithm.IsSet() {
		return true
	}

	return false
}

// SetAlgorithm gets a reference to the given NullableString and assigns it to the Algorithm field.
func (o *ErasureCodingParams) SetAlgorithm(v string) {
	o.Algorithm.Set(&v)
}
// SetAlgorithmNil sets the value for Algorithm to be an explicit nil
func (o *ErasureCodingParams) SetAlgorithmNil() {
	o.Algorithm.Set(nil)
}

// UnsetAlgorithm ensures that no value is present for Algorithm, not even an explicit nil
func (o *ErasureCodingParams) UnsetAlgorithm() {
	o.Algorithm.Unset()
}

// GetNumDataStripes returns the NumDataStripes field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *ErasureCodingParams) GetNumDataStripes() int32 {
	if o == nil || o.NumDataStripes.Get() == nil {
		var ret int32
		return ret
	}

	return *o.NumDataStripes.Get()
}

// GetNumDataStripesOk returns a tuple with the NumDataStripes field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ErasureCodingParams) GetNumDataStripesOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NumDataStripes.Get(), o.NumDataStripes.IsSet()
}

// SetNumDataStripes sets field value
func (o *ErasureCodingParams) SetNumDataStripes(v int32) {
	o.NumDataStripes.Set(&v)
}

// GetNumCodedStripes returns the NumCodedStripes field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *ErasureCodingParams) GetNumCodedStripes() int32 {
	if o == nil || o.NumCodedStripes.Get() == nil {
		var ret int32
		return ret
	}

	return *o.NumCodedStripes.Get()
}

// GetNumCodedStripesOk returns a tuple with the NumCodedStripes field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ErasureCodingParams) GetNumCodedStripesOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NumCodedStripes.Get(), o.NumCodedStripes.IsSet()
}

// SetNumCodedStripes sets field value
func (o *ErasureCodingParams) SetNumCodedStripes(v int32) {
	o.NumCodedStripes.Set(&v)
}

// GetInlineEnabled returns the InlineEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ErasureCodingParams) GetInlineEnabled() bool {
	if o == nil || o.InlineEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.InlineEnabled.Get()
}

// GetInlineEnabledOk returns a tuple with the InlineEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ErasureCodingParams) GetInlineEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.InlineEnabled.Get(), o.InlineEnabled.IsSet()
}

// HasInlineEnabled returns a boolean if a field has been set.
func (o *ErasureCodingParams) HasInlineEnabled() bool {
	if o != nil && o.InlineEnabled.IsSet() {
		return true
	}

	return false
}

// SetInlineEnabled gets a reference to the given NullableBool and assigns it to the InlineEnabled field.
func (o *ErasureCodingParams) SetInlineEnabled(v bool) {
	o.InlineEnabled.Set(&v)
}
// SetInlineEnabledNil sets the value for InlineEnabled to be an explicit nil
func (o *ErasureCodingParams) SetInlineEnabledNil() {
	o.InlineEnabled.Set(nil)
}

// UnsetInlineEnabled ensures that no value is present for InlineEnabled, not even an explicit nil
func (o *ErasureCodingParams) UnsetInlineEnabled() {
	o.InlineEnabled.Unset()
}

// GetDelaySecs returns the DelaySecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ErasureCodingParams) GetDelaySecs() int32 {
	if o == nil || o.DelaySecs.Get() == nil {
		var ret int32
		return ret
	}
	return *o.DelaySecs.Get()
}

// GetDelaySecsOk returns a tuple with the DelaySecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ErasureCodingParams) GetDelaySecsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DelaySecs.Get(), o.DelaySecs.IsSet()
}

// HasDelaySecs returns a boolean if a field has been set.
func (o *ErasureCodingParams) HasDelaySecs() bool {
	if o != nil && o.DelaySecs.IsSet() {
		return true
	}

	return false
}

// SetDelaySecs gets a reference to the given NullableInt32 and assigns it to the DelaySecs field.
func (o *ErasureCodingParams) SetDelaySecs(v int32) {
	o.DelaySecs.Set(&v)
}
// SetDelaySecsNil sets the value for DelaySecs to be an explicit nil
func (o *ErasureCodingParams) SetDelaySecsNil() {
	o.DelaySecs.Set(nil)
}

// UnsetDelaySecs ensures that no value is present for DelaySecs, not even an explicit nil
func (o *ErasureCodingParams) UnsetDelaySecs() {
	o.DelaySecs.Unset()
}

// GetUseContainer returns the UseContainer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ErasureCodingParams) GetUseContainer() bool {
	if o == nil || o.UseContainer.Get() == nil {
		var ret bool
		return ret
	}
	return *o.UseContainer.Get()
}

// GetUseContainerOk returns a tuple with the UseContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ErasureCodingParams) GetUseContainerOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UseContainer.Get(), o.UseContainer.IsSet()
}

// HasUseContainer returns a boolean if a field has been set.
func (o *ErasureCodingParams) HasUseContainer() bool {
	if o != nil && o.UseContainer.IsSet() {
		return true
	}

	return false
}

// SetUseContainer gets a reference to the given NullableBool and assigns it to the UseContainer field.
func (o *ErasureCodingParams) SetUseContainer(v bool) {
	o.UseContainer.Set(&v)
}
// SetUseContainerNil sets the value for UseContainer to be an explicit nil
func (o *ErasureCodingParams) SetUseContainerNil() {
	o.UseContainer.Set(nil)
}

// UnsetUseContainer ensures that no value is present for UseContainer, not even an explicit nil
func (o *ErasureCodingParams) UnsetUseContainer() {
	o.UseContainer.Unset()
}

func (o ErasureCodingParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled.Get()
	}
	if o.Algorithm.IsSet() {
		toSerialize["algorithm"] = o.Algorithm.Get()
	}
	if true {
		toSerialize["numDataStripes"] = o.NumDataStripes.Get()
	}
	if true {
		toSerialize["numCodedStripes"] = o.NumCodedStripes.Get()
	}
	if o.InlineEnabled.IsSet() {
		toSerialize["inlineEnabled"] = o.InlineEnabled.Get()
	}
	if o.DelaySecs.IsSet() {
		toSerialize["delaySecs"] = o.DelaySecs.Get()
	}
	if o.UseContainer.IsSet() {
		toSerialize["useContainer"] = o.UseContainer.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableErasureCodingParams struct {
	value *ErasureCodingParams
	isSet bool
}

func (v NullableErasureCodingParams) Get() *ErasureCodingParams {
	return v.value
}

func (v *NullableErasureCodingParams) Set(val *ErasureCodingParams) {
	v.value = val
	v.isSet = true
}

func (v NullableErasureCodingParams) IsSet() bool {
	return v.isSet
}

func (v *NullableErasureCodingParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErasureCodingParams(val *ErasureCodingParams) *NullableErasureCodingParams {
	return &NullableErasureCodingParams{value: val, isSet: true}
}

func (v NullableErasureCodingParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErasureCodingParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


