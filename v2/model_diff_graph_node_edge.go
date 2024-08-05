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

// checks if the DiffGraphNodeEdge type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiffGraphNodeEdge{}

// DiffGraphNodeEdge Represents the pair of edges and destination node which are either modified, added, deleted or unmodified compared to base snapshot.
type DiffGraphNodeEdge struct {
	// This will be present if either the pair of edge relation and destination node is modified or deleted. If diff type is unmodified this field will not be present.
	BaseRelation *GraphEdge `json:"baseRelation,omitempty"`
	// This will be present if either the pair of edge relation and destination node is added, modified or unmodified.
	CurrentRelation *GraphEdge `json:"currentRelation,omitempty"`
	// Specifies the diff type for the base node.
	DiffType NullableString `json:"diffType,omitempty"`
}

// NewDiffGraphNodeEdge instantiates a new DiffGraphNodeEdge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiffGraphNodeEdge() *DiffGraphNodeEdge {
	this := DiffGraphNodeEdge{}
	return &this
}

// NewDiffGraphNodeEdgeWithDefaults instantiates a new DiffGraphNodeEdge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiffGraphNodeEdgeWithDefaults() *DiffGraphNodeEdge {
	this := DiffGraphNodeEdge{}
	return &this
}

// GetBaseRelation returns the BaseRelation field value if set, zero value otherwise.
func (o *DiffGraphNodeEdge) GetBaseRelation() GraphEdge {
	if o == nil || IsNil(o.BaseRelation) {
		var ret GraphEdge
		return ret
	}
	return *o.BaseRelation
}

// GetBaseRelationOk returns a tuple with the BaseRelation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiffGraphNodeEdge) GetBaseRelationOk() (*GraphEdge, bool) {
	if o == nil || IsNil(o.BaseRelation) {
		return nil, false
	}
	return o.BaseRelation, true
}

// HasBaseRelation returns a boolean if a field has been set.
func (o *DiffGraphNodeEdge) HasBaseRelation() bool {
	if o != nil && !IsNil(o.BaseRelation) {
		return true
	}

	return false
}

// SetBaseRelation gets a reference to the given GraphEdge and assigns it to the BaseRelation field.
func (o *DiffGraphNodeEdge) SetBaseRelation(v GraphEdge) {
	o.BaseRelation = &v
}

// GetCurrentRelation returns the CurrentRelation field value if set, zero value otherwise.
func (o *DiffGraphNodeEdge) GetCurrentRelation() GraphEdge {
	if o == nil || IsNil(o.CurrentRelation) {
		var ret GraphEdge
		return ret
	}
	return *o.CurrentRelation
}

// GetCurrentRelationOk returns a tuple with the CurrentRelation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiffGraphNodeEdge) GetCurrentRelationOk() (*GraphEdge, bool) {
	if o == nil || IsNil(o.CurrentRelation) {
		return nil, false
	}
	return o.CurrentRelation, true
}

// HasCurrentRelation returns a boolean if a field has been set.
func (o *DiffGraphNodeEdge) HasCurrentRelation() bool {
	if o != nil && !IsNil(o.CurrentRelation) {
		return true
	}

	return false
}

// SetCurrentRelation gets a reference to the given GraphEdge and assigns it to the CurrentRelation field.
func (o *DiffGraphNodeEdge) SetCurrentRelation(v GraphEdge) {
	o.CurrentRelation = &v
}

// GetDiffType returns the DiffType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiffGraphNodeEdge) GetDiffType() string {
	if o == nil || IsNil(o.DiffType.Get()) {
		var ret string
		return ret
	}
	return *o.DiffType.Get()
}

// GetDiffTypeOk returns a tuple with the DiffType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiffGraphNodeEdge) GetDiffTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiffType.Get(), o.DiffType.IsSet()
}

// HasDiffType returns a boolean if a field has been set.
func (o *DiffGraphNodeEdge) HasDiffType() bool {
	if o != nil && o.DiffType.IsSet() {
		return true
	}

	return false
}

// SetDiffType gets a reference to the given NullableString and assigns it to the DiffType field.
func (o *DiffGraphNodeEdge) SetDiffType(v string) {
	o.DiffType.Set(&v)
}
// SetDiffTypeNil sets the value for DiffType to be an explicit nil
func (o *DiffGraphNodeEdge) SetDiffTypeNil() {
	o.DiffType.Set(nil)
}

// UnsetDiffType ensures that no value is present for DiffType, not even an explicit nil
func (o *DiffGraphNodeEdge) UnsetDiffType() {
	o.DiffType.Unset()
}

func (o DiffGraphNodeEdge) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiffGraphNodeEdge) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BaseRelation) {
		toSerialize["baseRelation"] = o.BaseRelation
	}
	if !IsNil(o.CurrentRelation) {
		toSerialize["currentRelation"] = o.CurrentRelation
	}
	if o.DiffType.IsSet() {
		toSerialize["diffType"] = o.DiffType.Get()
	}
	return toSerialize, nil
}

type NullableDiffGraphNodeEdge struct {
	value *DiffGraphNodeEdge
	isSet bool
}

func (v NullableDiffGraphNodeEdge) Get() *DiffGraphNodeEdge {
	return v.value
}

func (v *NullableDiffGraphNodeEdge) Set(val *DiffGraphNodeEdge) {
	v.value = val
	v.isSet = true
}

func (v NullableDiffGraphNodeEdge) IsSet() bool {
	return v.isSet
}

func (v *NullableDiffGraphNodeEdge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiffGraphNodeEdge(val *DiffGraphNodeEdge) *NullableDiffGraphNodeEdge {
	return &NullableDiffGraphNodeEdge{value: val, isSet: true}
}

func (v NullableDiffGraphNodeEdge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiffGraphNodeEdge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


