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

// ConnectionConfig Specifies a connection associated with the source.
type ConnectionConfig struct {
	// Specifies the id of the connection.
	ConnectionId NullableInt64 `json:"connectionId,omitempty"`
	// Specifies the entity id of the source. The source can a non-root entity.
	EntityId NullableInt64 `json:"entityId,omitempty"`
}

// NewConnectionConfig instantiates a new ConnectionConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionConfig() *ConnectionConfig {
	this := ConnectionConfig{}
	return &this
}

// NewConnectionConfigWithDefaults instantiates a new ConnectionConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionConfigWithDefaults() *ConnectionConfig {
	this := ConnectionConfig{}
	return &this
}

// GetConnectionId returns the ConnectionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectionConfig) GetConnectionId() int64 {
	if o == nil || o.ConnectionId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ConnectionId.Get()
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectionConfig) GetConnectionIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ConnectionId.Get(), o.ConnectionId.IsSet()
}

// HasConnectionId returns a boolean if a field has been set.
func (o *ConnectionConfig) HasConnectionId() bool {
	if o != nil && o.ConnectionId.IsSet() {
		return true
	}

	return false
}

// SetConnectionId gets a reference to the given NullableInt64 and assigns it to the ConnectionId field.
func (o *ConnectionConfig) SetConnectionId(v int64) {
	o.ConnectionId.Set(&v)
}
// SetConnectionIdNil sets the value for ConnectionId to be an explicit nil
func (o *ConnectionConfig) SetConnectionIdNil() {
	o.ConnectionId.Set(nil)
}

// UnsetConnectionId ensures that no value is present for ConnectionId, not even an explicit nil
func (o *ConnectionConfig) UnsetConnectionId() {
	o.ConnectionId.Unset()
}

// GetEntityId returns the EntityId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectionConfig) GetEntityId() int64 {
	if o == nil || o.EntityId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.EntityId.Get()
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectionConfig) GetEntityIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EntityId.Get(), o.EntityId.IsSet()
}

// HasEntityId returns a boolean if a field has been set.
func (o *ConnectionConfig) HasEntityId() bool {
	if o != nil && o.EntityId.IsSet() {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given NullableInt64 and assigns it to the EntityId field.
func (o *ConnectionConfig) SetEntityId(v int64) {
	o.EntityId.Set(&v)
}
// SetEntityIdNil sets the value for EntityId to be an explicit nil
func (o *ConnectionConfig) SetEntityIdNil() {
	o.EntityId.Set(nil)
}

// UnsetEntityId ensures that no value is present for EntityId, not even an explicit nil
func (o *ConnectionConfig) UnsetEntityId() {
	o.EntityId.Unset()
}

func (o ConnectionConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ConnectionId.IsSet() {
		toSerialize["connectionId"] = o.ConnectionId.Get()
	}
	if o.EntityId.IsSet() {
		toSerialize["entityId"] = o.EntityId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableConnectionConfig struct {
	value *ConnectionConfig
	isSet bool
}

func (v NullableConnectionConfig) Get() *ConnectionConfig {
	return v.value
}

func (v *NullableConnectionConfig) Set(val *ConnectionConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionConfig(val *ConnectionConfig) *NullableConnectionConfig {
	return &NullableConnectionConfig{value: val, isSet: true}
}

func (v NullableConnectionConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


