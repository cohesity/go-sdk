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

// checks if the SourceHierarchyObjectSummaries type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourceHierarchyObjectSummaries{}

// SourceHierarchyObjectSummaries Specifies a list of Source Hierarchy Object Summaries.
type SourceHierarchyObjectSummaries struct {
	// Specifies a list of Source Hierarchy Object summaries.
	Objects []SourceHierarchyObjectSummary `json:"objects,omitempty"`
}

// NewSourceHierarchyObjectSummaries instantiates a new SourceHierarchyObjectSummaries object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceHierarchyObjectSummaries() *SourceHierarchyObjectSummaries {
	this := SourceHierarchyObjectSummaries{}
	return &this
}

// NewSourceHierarchyObjectSummariesWithDefaults instantiates a new SourceHierarchyObjectSummaries object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceHierarchyObjectSummariesWithDefaults() *SourceHierarchyObjectSummaries {
	this := SourceHierarchyObjectSummaries{}
	return &this
}

// GetObjects returns the Objects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourceHierarchyObjectSummaries) GetObjects() []SourceHierarchyObjectSummary {
	if o == nil {
		var ret []SourceHierarchyObjectSummary
		return ret
	}
	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SourceHierarchyObjectSummaries) GetObjectsOk() ([]SourceHierarchyObjectSummary, bool) {
	if o == nil || IsNil(o.Objects) {
		return nil, false
	}
	return o.Objects, true
}

// HasObjects returns a boolean if a field has been set.
func (o *SourceHierarchyObjectSummaries) HasObjects() bool {
	if o != nil && !IsNil(o.Objects) {
		return true
	}

	return false
}

// SetObjects gets a reference to the given []SourceHierarchyObjectSummary and assigns it to the Objects field.
func (o *SourceHierarchyObjectSummaries) SetObjects(v []SourceHierarchyObjectSummary) {
	o.Objects = v
}

func (o SourceHierarchyObjectSummaries) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceHierarchyObjectSummaries) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Objects != nil {
		toSerialize["objects"] = o.Objects
	}
	return toSerialize, nil
}

type NullableSourceHierarchyObjectSummaries struct {
	value *SourceHierarchyObjectSummaries
	isSet bool
}

func (v NullableSourceHierarchyObjectSummaries) Get() *SourceHierarchyObjectSummaries {
	return v.value
}

func (v *NullableSourceHierarchyObjectSummaries) Set(val *SourceHierarchyObjectSummaries) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceHierarchyObjectSummaries) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceHierarchyObjectSummaries) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceHierarchyObjectSummaries(val *SourceHierarchyObjectSummaries) *NullableSourceHierarchyObjectSummaries {
	return &NullableSourceHierarchyObjectSummaries{value: val, isSet: true}
}

func (v NullableSourceHierarchyObjectSummaries) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceHierarchyObjectSummaries) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


