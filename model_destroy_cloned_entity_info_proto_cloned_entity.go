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

// DestroyClonedEntityInfoProtoClonedEntity struct for DestroyClonedEntityInfoProtoClonedEntity
type DestroyClonedEntityInfoProtoClonedEntity struct {
	Entity *EntityProto `json:"entity,omitempty"`
	// Path of all files created by the clone operation. Each path is relative to the clone view.
	RelativeRestorePathVec []string `json:"relativeRestorePathVec,omitempty"`
}

// NewDestroyClonedEntityInfoProtoClonedEntity instantiates a new DestroyClonedEntityInfoProtoClonedEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDestroyClonedEntityInfoProtoClonedEntity() *DestroyClonedEntityInfoProtoClonedEntity {
	this := DestroyClonedEntityInfoProtoClonedEntity{}
	return &this
}

// NewDestroyClonedEntityInfoProtoClonedEntityWithDefaults instantiates a new DestroyClonedEntityInfoProtoClonedEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDestroyClonedEntityInfoProtoClonedEntityWithDefaults() *DestroyClonedEntityInfoProtoClonedEntity {
	this := DestroyClonedEntityInfoProtoClonedEntity{}
	return &this
}

// GetEntity returns the Entity field value if set, zero value otherwise.
func (o *DestroyClonedEntityInfoProtoClonedEntity) GetEntity() EntityProto {
	if o == nil || o.Entity == nil {
		var ret EntityProto
		return ret
	}
	return *o.Entity
}

// GetEntityOk returns a tuple with the Entity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestroyClonedEntityInfoProtoClonedEntity) GetEntityOk() (*EntityProto, bool) {
	if o == nil || o.Entity == nil {
		return nil, false
	}
	return o.Entity, true
}

// HasEntity returns a boolean if a field has been set.
func (o *DestroyClonedEntityInfoProtoClonedEntity) HasEntity() bool {
	if o != nil && o.Entity != nil {
		return true
	}

	return false
}

// SetEntity gets a reference to the given EntityProto and assigns it to the Entity field.
func (o *DestroyClonedEntityInfoProtoClonedEntity) SetEntity(v EntityProto) {
	o.Entity = &v
}

// GetRelativeRestorePathVec returns the RelativeRestorePathVec field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DestroyClonedEntityInfoProtoClonedEntity) GetRelativeRestorePathVec() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.RelativeRestorePathVec
}

// GetRelativeRestorePathVecOk returns a tuple with the RelativeRestorePathVec field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DestroyClonedEntityInfoProtoClonedEntity) GetRelativeRestorePathVecOk() (*[]string, bool) {
	if o == nil || o.RelativeRestorePathVec == nil {
		return nil, false
	}
	return &o.RelativeRestorePathVec, true
}

// HasRelativeRestorePathVec returns a boolean if a field has been set.
func (o *DestroyClonedEntityInfoProtoClonedEntity) HasRelativeRestorePathVec() bool {
	if o != nil && o.RelativeRestorePathVec != nil {
		return true
	}

	return false
}

// SetRelativeRestorePathVec gets a reference to the given []string and assigns it to the RelativeRestorePathVec field.
func (o *DestroyClonedEntityInfoProtoClonedEntity) SetRelativeRestorePathVec(v []string) {
	o.RelativeRestorePathVec = v
}

func (o DestroyClonedEntityInfoProtoClonedEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Entity != nil {
		toSerialize["entity"] = o.Entity
	}
	if o.RelativeRestorePathVec != nil {
		toSerialize["relativeRestorePathVec"] = o.RelativeRestorePathVec
	}
	return json.Marshal(toSerialize)
}

type NullableDestroyClonedEntityInfoProtoClonedEntity struct {
	value *DestroyClonedEntityInfoProtoClonedEntity
	isSet bool
}

func (v NullableDestroyClonedEntityInfoProtoClonedEntity) Get() *DestroyClonedEntityInfoProtoClonedEntity {
	return v.value
}

func (v *NullableDestroyClonedEntityInfoProtoClonedEntity) Set(val *DestroyClonedEntityInfoProtoClonedEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableDestroyClonedEntityInfoProtoClonedEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableDestroyClonedEntityInfoProtoClonedEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDestroyClonedEntityInfoProtoClonedEntity(val *DestroyClonedEntityInfoProtoClonedEntity) *NullableDestroyClonedEntityInfoProtoClonedEntity {
	return &NullableDestroyClonedEntityInfoProtoClonedEntity{value: val, isSet: true}
}

func (v NullableDestroyClonedEntityInfoProtoClonedEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDestroyClonedEntityInfoProtoClonedEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


