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

// SourceFiltersSourceFilter Plain text filter: { source_filter: \"TestDatabase\", is_regex: false}. Wildcard filter: { source_filter: \"Test?Database*\", is_regex: false}. Regex filter: { source_filter: \"^Test.*Database$\", is_regex: true}.
type SourceFiltersSourceFilter struct {
	// If true, this implies 'source_filter' is a regex filter. If false, it will be treated as wildcard/plain text filter.
	IsRegex NullableBool `json:"isRegex,omitempty"`
	// This contains the filter string.
	SourceFilter NullableString `json:"sourceFilter,omitempty"`
}

// NewSourceFiltersSourceFilter instantiates a new SourceFiltersSourceFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceFiltersSourceFilter() *SourceFiltersSourceFilter {
	this := SourceFiltersSourceFilter{}
	return &this
}

// NewSourceFiltersSourceFilterWithDefaults instantiates a new SourceFiltersSourceFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceFiltersSourceFilterWithDefaults() *SourceFiltersSourceFilter {
	this := SourceFiltersSourceFilter{}
	return &this
}

// GetIsRegex returns the IsRegex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourceFiltersSourceFilter) GetIsRegex() bool {
	if o == nil || o.IsRegex.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsRegex.Get()
}

// GetIsRegexOk returns a tuple with the IsRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SourceFiltersSourceFilter) GetIsRegexOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsRegex.Get(), o.IsRegex.IsSet()
}

// HasIsRegex returns a boolean if a field has been set.
func (o *SourceFiltersSourceFilter) HasIsRegex() bool {
	if o != nil && o.IsRegex.IsSet() {
		return true
	}

	return false
}

// SetIsRegex gets a reference to the given NullableBool and assigns it to the IsRegex field.
func (o *SourceFiltersSourceFilter) SetIsRegex(v bool) {
	o.IsRegex.Set(&v)
}
// SetIsRegexNil sets the value for IsRegex to be an explicit nil
func (o *SourceFiltersSourceFilter) SetIsRegexNil() {
	o.IsRegex.Set(nil)
}

// UnsetIsRegex ensures that no value is present for IsRegex, not even an explicit nil
func (o *SourceFiltersSourceFilter) UnsetIsRegex() {
	o.IsRegex.Unset()
}

// GetSourceFilter returns the SourceFilter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourceFiltersSourceFilter) GetSourceFilter() string {
	if o == nil || o.SourceFilter.Get() == nil {
		var ret string
		return ret
	}
	return *o.SourceFilter.Get()
}

// GetSourceFilterOk returns a tuple with the SourceFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SourceFiltersSourceFilter) GetSourceFilterOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SourceFilter.Get(), o.SourceFilter.IsSet()
}

// HasSourceFilter returns a boolean if a field has been set.
func (o *SourceFiltersSourceFilter) HasSourceFilter() bool {
	if o != nil && o.SourceFilter.IsSet() {
		return true
	}

	return false
}

// SetSourceFilter gets a reference to the given NullableString and assigns it to the SourceFilter field.
func (o *SourceFiltersSourceFilter) SetSourceFilter(v string) {
	o.SourceFilter.Set(&v)
}
// SetSourceFilterNil sets the value for SourceFilter to be an explicit nil
func (o *SourceFiltersSourceFilter) SetSourceFilterNil() {
	o.SourceFilter.Set(nil)
}

// UnsetSourceFilter ensures that no value is present for SourceFilter, not even an explicit nil
func (o *SourceFiltersSourceFilter) UnsetSourceFilter() {
	o.SourceFilter.Unset()
}

func (o SourceFiltersSourceFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsRegex.IsSet() {
		toSerialize["isRegex"] = o.IsRegex.Get()
	}
	if o.SourceFilter.IsSet() {
		toSerialize["sourceFilter"] = o.SourceFilter.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSourceFiltersSourceFilter struct {
	value *SourceFiltersSourceFilter
	isSet bool
}

func (v NullableSourceFiltersSourceFilter) Get() *SourceFiltersSourceFilter {
	return v.value
}

func (v *NullableSourceFiltersSourceFilter) Set(val *SourceFiltersSourceFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceFiltersSourceFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceFiltersSourceFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceFiltersSourceFilter(val *SourceFiltersSourceFilter) *NullableSourceFiltersSourceFilter {
	return &NullableSourceFiltersSourceFilter{value: val, isSet: true}
}

func (v NullableSourceFiltersSourceFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceFiltersSourceFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


