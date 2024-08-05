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

// checks if the LifecycleRuleFilterAnd type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LifecycleRuleFilterAnd{}

// LifecycleRuleFilterAnd Specifies the Lifecycle Rule Filter to apply a logical AND to two or more predicates. The Lifecycle Rule will apply to any object matching all of the predicates configured inside the AND operator.
type LifecycleRuleFilterAnd struct {
	// Specifies a Prefix identifying one or more objects to which the rule applies.
	Prefix NullableString `json:"prefix,omitempty"`
	// Specifies the tag in the object's tag set to which the rule applies. All of these tags must exist in the object's tag set in order for the rule to apply.
	Tags []LifecycleRuleFilterTag `json:"tags,omitempty"`
}

// NewLifecycleRuleFilterAnd instantiates a new LifecycleRuleFilterAnd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLifecycleRuleFilterAnd() *LifecycleRuleFilterAnd {
	this := LifecycleRuleFilterAnd{}
	return &this
}

// NewLifecycleRuleFilterAndWithDefaults instantiates a new LifecycleRuleFilterAnd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLifecycleRuleFilterAndWithDefaults() *LifecycleRuleFilterAnd {
	this := LifecycleRuleFilterAnd{}
	return &this
}

// GetPrefix returns the Prefix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LifecycleRuleFilterAnd) GetPrefix() string {
	if o == nil || IsNil(o.Prefix.Get()) {
		var ret string
		return ret
	}
	return *o.Prefix.Get()
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LifecycleRuleFilterAnd) GetPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Prefix.Get(), o.Prefix.IsSet()
}

// HasPrefix returns a boolean if a field has been set.
func (o *LifecycleRuleFilterAnd) HasPrefix() bool {
	if o != nil && o.Prefix.IsSet() {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given NullableString and assigns it to the Prefix field.
func (o *LifecycleRuleFilterAnd) SetPrefix(v string) {
	o.Prefix.Set(&v)
}
// SetPrefixNil sets the value for Prefix to be an explicit nil
func (o *LifecycleRuleFilterAnd) SetPrefixNil() {
	o.Prefix.Set(nil)
}

// UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil
func (o *LifecycleRuleFilterAnd) UnsetPrefix() {
	o.Prefix.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LifecycleRuleFilterAnd) GetTags() []LifecycleRuleFilterTag {
	if o == nil {
		var ret []LifecycleRuleFilterTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LifecycleRuleFilterAnd) GetTagsOk() ([]LifecycleRuleFilterTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *LifecycleRuleFilterAnd) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []LifecycleRuleFilterTag and assigns it to the Tags field.
func (o *LifecycleRuleFilterAnd) SetTags(v []LifecycleRuleFilterTag) {
	o.Tags = v
}

func (o LifecycleRuleFilterAnd) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LifecycleRuleFilterAnd) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Prefix.IsSet() {
		toSerialize["prefix"] = o.Prefix.Get()
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableLifecycleRuleFilterAnd struct {
	value *LifecycleRuleFilterAnd
	isSet bool
}

func (v NullableLifecycleRuleFilterAnd) Get() *LifecycleRuleFilterAnd {
	return v.value
}

func (v *NullableLifecycleRuleFilterAnd) Set(val *LifecycleRuleFilterAnd) {
	v.value = val
	v.isSet = true
}

func (v NullableLifecycleRuleFilterAnd) IsSet() bool {
	return v.isSet
}

func (v *NullableLifecycleRuleFilterAnd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLifecycleRuleFilterAnd(val *LifecycleRuleFilterAnd) *NullableLifecycleRuleFilterAnd {
	return &NullableLifecycleRuleFilterAnd{value: val, isSet: true}
}

func (v NullableLifecycleRuleFilterAnd) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLifecycleRuleFilterAnd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


