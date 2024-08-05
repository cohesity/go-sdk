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

// checks if the EntitySelector type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitySelector{}

// EntitySelector Selector for entities.
type EntitySelector struct {
	// If set to true all the entities ever known to this cluster will be exported. These include but are not limited to registered, unregistered, and replicated entities
	AllEntities NullableBool `json:"allEntities,omitempty"`
	// If this is true, the entities that are regisered will be qualified. Else, all the entities will be matched.
	RegisteredEntities NullableBool `json:"registeredEntities,omitempty"`
}

// NewEntitySelector instantiates a new EntitySelector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitySelector() *EntitySelector {
	this := EntitySelector{}
	return &this
}

// NewEntitySelectorWithDefaults instantiates a new EntitySelector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitySelectorWithDefaults() *EntitySelector {
	this := EntitySelector{}
	return &this
}

// GetAllEntities returns the AllEntities field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntitySelector) GetAllEntities() bool {
	if o == nil || IsNil(o.AllEntities.Get()) {
		var ret bool
		return ret
	}
	return *o.AllEntities.Get()
}

// GetAllEntitiesOk returns a tuple with the AllEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntitySelector) GetAllEntitiesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllEntities.Get(), o.AllEntities.IsSet()
}

// HasAllEntities returns a boolean if a field has been set.
func (o *EntitySelector) HasAllEntities() bool {
	if o != nil && o.AllEntities.IsSet() {
		return true
	}

	return false
}

// SetAllEntities gets a reference to the given NullableBool and assigns it to the AllEntities field.
func (o *EntitySelector) SetAllEntities(v bool) {
	o.AllEntities.Set(&v)
}
// SetAllEntitiesNil sets the value for AllEntities to be an explicit nil
func (o *EntitySelector) SetAllEntitiesNil() {
	o.AllEntities.Set(nil)
}

// UnsetAllEntities ensures that no value is present for AllEntities, not even an explicit nil
func (o *EntitySelector) UnsetAllEntities() {
	o.AllEntities.Unset()
}

// GetRegisteredEntities returns the RegisteredEntities field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntitySelector) GetRegisteredEntities() bool {
	if o == nil || IsNil(o.RegisteredEntities.Get()) {
		var ret bool
		return ret
	}
	return *o.RegisteredEntities.Get()
}

// GetRegisteredEntitiesOk returns a tuple with the RegisteredEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntitySelector) GetRegisteredEntitiesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredEntities.Get(), o.RegisteredEntities.IsSet()
}

// HasRegisteredEntities returns a boolean if a field has been set.
func (o *EntitySelector) HasRegisteredEntities() bool {
	if o != nil && o.RegisteredEntities.IsSet() {
		return true
	}

	return false
}

// SetRegisteredEntities gets a reference to the given NullableBool and assigns it to the RegisteredEntities field.
func (o *EntitySelector) SetRegisteredEntities(v bool) {
	o.RegisteredEntities.Set(&v)
}
// SetRegisteredEntitiesNil sets the value for RegisteredEntities to be an explicit nil
func (o *EntitySelector) SetRegisteredEntitiesNil() {
	o.RegisteredEntities.Set(nil)
}

// UnsetRegisteredEntities ensures that no value is present for RegisteredEntities, not even an explicit nil
func (o *EntitySelector) UnsetRegisteredEntities() {
	o.RegisteredEntities.Unset()
}

func (o EntitySelector) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitySelector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AllEntities.IsSet() {
		toSerialize["allEntities"] = o.AllEntities.Get()
	}
	if o.RegisteredEntities.IsSet() {
		toSerialize["registeredEntities"] = o.RegisteredEntities.Get()
	}
	return toSerialize, nil
}

type NullableEntitySelector struct {
	value *EntitySelector
	isSet bool
}

func (v NullableEntitySelector) Get() *EntitySelector {
	return v.value
}

func (v *NullableEntitySelector) Set(val *EntitySelector) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitySelector) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitySelector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitySelector(val *EntitySelector) *NullableEntitySelector {
	return &NullableEntitySelector{value: val, isSet: true}
}

func (v NullableEntitySelector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitySelector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


