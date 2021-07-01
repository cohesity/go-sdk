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

// SchemaInfo Specifies the metric data point where public data metric name as key and the schema defined metric name as a value.
type SchemaInfo struct {
	// Specifies the id of the entity represented as a string.
	EntityId NullableString `json:"entityId,omitempty"`
	// Specifies the key which is public facing name for metric name.
	Key NullableString `json:"key,omitempty"`
	// Specifies the Apollo schema metric name.
	MetricName NullableString `json:"metricName,omitempty"`
	// Specifies the name of entity schema such as 'ApolloViewBoxStats'.
	SchemaName NullableString `json:"schemaName,omitempty"`
}

// NewSchemaInfo instantiates a new SchemaInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemaInfo() *SchemaInfo {
	this := SchemaInfo{}
	return &this
}

// NewSchemaInfoWithDefaults instantiates a new SchemaInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemaInfoWithDefaults() *SchemaInfo {
	this := SchemaInfo{}
	return &this
}

// GetEntityId returns the EntityId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SchemaInfo) GetEntityId() string {
	if o == nil || o.EntityId.Get() == nil {
		var ret string
		return ret
	}
	return *o.EntityId.Get()
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SchemaInfo) GetEntityIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EntityId.Get(), o.EntityId.IsSet()
}

// HasEntityId returns a boolean if a field has been set.
func (o *SchemaInfo) HasEntityId() bool {
	if o != nil && o.EntityId.IsSet() {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given NullableString and assigns it to the EntityId field.
func (o *SchemaInfo) SetEntityId(v string) {
	o.EntityId.Set(&v)
}
// SetEntityIdNil sets the value for EntityId to be an explicit nil
func (o *SchemaInfo) SetEntityIdNil() {
	o.EntityId.Set(nil)
}

// UnsetEntityId ensures that no value is present for EntityId, not even an explicit nil
func (o *SchemaInfo) UnsetEntityId() {
	o.EntityId.Unset()
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SchemaInfo) GetKey() string {
	if o == nil || o.Key.Get() == nil {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SchemaInfo) GetKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *SchemaInfo) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *SchemaInfo) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *SchemaInfo) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *SchemaInfo) UnsetKey() {
	o.Key.Unset()
}

// GetMetricName returns the MetricName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SchemaInfo) GetMetricName() string {
	if o == nil || o.MetricName.Get() == nil {
		var ret string
		return ret
	}
	return *o.MetricName.Get()
}

// GetMetricNameOk returns a tuple with the MetricName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SchemaInfo) GetMetricNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MetricName.Get(), o.MetricName.IsSet()
}

// HasMetricName returns a boolean if a field has been set.
func (o *SchemaInfo) HasMetricName() bool {
	if o != nil && o.MetricName.IsSet() {
		return true
	}

	return false
}

// SetMetricName gets a reference to the given NullableString and assigns it to the MetricName field.
func (o *SchemaInfo) SetMetricName(v string) {
	o.MetricName.Set(&v)
}
// SetMetricNameNil sets the value for MetricName to be an explicit nil
func (o *SchemaInfo) SetMetricNameNil() {
	o.MetricName.Set(nil)
}

// UnsetMetricName ensures that no value is present for MetricName, not even an explicit nil
func (o *SchemaInfo) UnsetMetricName() {
	o.MetricName.Unset()
}

// GetSchemaName returns the SchemaName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SchemaInfo) GetSchemaName() string {
	if o == nil || o.SchemaName.Get() == nil {
		var ret string
		return ret
	}
	return *o.SchemaName.Get()
}

// GetSchemaNameOk returns a tuple with the SchemaName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SchemaInfo) GetSchemaNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SchemaName.Get(), o.SchemaName.IsSet()
}

// HasSchemaName returns a boolean if a field has been set.
func (o *SchemaInfo) HasSchemaName() bool {
	if o != nil && o.SchemaName.IsSet() {
		return true
	}

	return false
}

// SetSchemaName gets a reference to the given NullableString and assigns it to the SchemaName field.
func (o *SchemaInfo) SetSchemaName(v string) {
	o.SchemaName.Set(&v)
}
// SetSchemaNameNil sets the value for SchemaName to be an explicit nil
func (o *SchemaInfo) SetSchemaNameNil() {
	o.SchemaName.Set(nil)
}

// UnsetSchemaName ensures that no value is present for SchemaName, not even an explicit nil
func (o *SchemaInfo) UnsetSchemaName() {
	o.SchemaName.Unset()
}

func (o SchemaInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EntityId.IsSet() {
		toSerialize["entityId"] = o.EntityId.Get()
	}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	if o.MetricName.IsSet() {
		toSerialize["metricName"] = o.MetricName.Get()
	}
	if o.SchemaName.IsSet() {
		toSerialize["schemaName"] = o.SchemaName.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSchemaInfo struct {
	value *SchemaInfo
	isSet bool
}

func (v NullableSchemaInfo) Get() *SchemaInfo {
	return v.value
}

func (v *NullableSchemaInfo) Set(val *SchemaInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemaInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemaInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemaInfo(val *SchemaInfo) *NullableSchemaInfo {
	return &NullableSchemaInfo{value: val, isSet: true}
}

func (v NullableSchemaInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemaInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


