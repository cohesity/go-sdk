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

// ClusterConfigProtoQoSMapping If a new enum value is added to either QoSMapping.Type or QoSPrincipal.Priority in a future version, direct upgrades must be disallowed from a pre-2.5 version to that version (without upgrading to 2.5 first). Contact nexus team for getting an appropriate restriction into the upgrade compatibility list.
type ClusterConfigProtoQoSMapping struct {
	// Principal id of the QoS principal to which qos_context maps to.
	PrincipalId NullableInt64 `json:"principalId,omitempty"`
	QosContext *ClusterConfigProtoQoSMappingQoSContext `json:"qosContext,omitempty"`
}

// NewClusterConfigProtoQoSMapping instantiates a new ClusterConfigProtoQoSMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterConfigProtoQoSMapping() *ClusterConfigProtoQoSMapping {
	this := ClusterConfigProtoQoSMapping{}
	return &this
}

// NewClusterConfigProtoQoSMappingWithDefaults instantiates a new ClusterConfigProtoQoSMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterConfigProtoQoSMappingWithDefaults() *ClusterConfigProtoQoSMapping {
	this := ClusterConfigProtoQoSMapping{}
	return &this
}

// GetPrincipalId returns the PrincipalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterConfigProtoQoSMapping) GetPrincipalId() int64 {
	if o == nil || o.PrincipalId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.PrincipalId.Get()
}

// GetPrincipalIdOk returns a tuple with the PrincipalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterConfigProtoQoSMapping) GetPrincipalIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PrincipalId.Get(), o.PrincipalId.IsSet()
}

// HasPrincipalId returns a boolean if a field has been set.
func (o *ClusterConfigProtoQoSMapping) HasPrincipalId() bool {
	if o != nil && o.PrincipalId.IsSet() {
		return true
	}

	return false
}

// SetPrincipalId gets a reference to the given NullableInt64 and assigns it to the PrincipalId field.
func (o *ClusterConfigProtoQoSMapping) SetPrincipalId(v int64) {
	o.PrincipalId.Set(&v)
}
// SetPrincipalIdNil sets the value for PrincipalId to be an explicit nil
func (o *ClusterConfigProtoQoSMapping) SetPrincipalIdNil() {
	o.PrincipalId.Set(nil)
}

// UnsetPrincipalId ensures that no value is present for PrincipalId, not even an explicit nil
func (o *ClusterConfigProtoQoSMapping) UnsetPrincipalId() {
	o.PrincipalId.Unset()
}

// GetQosContext returns the QosContext field value if set, zero value otherwise.
func (o *ClusterConfigProtoQoSMapping) GetQosContext() ClusterConfigProtoQoSMappingQoSContext {
	if o == nil || o.QosContext == nil {
		var ret ClusterConfigProtoQoSMappingQoSContext
		return ret
	}
	return *o.QosContext
}

// GetQosContextOk returns a tuple with the QosContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterConfigProtoQoSMapping) GetQosContextOk() (*ClusterConfigProtoQoSMappingQoSContext, bool) {
	if o == nil || o.QosContext == nil {
		return nil, false
	}
	return o.QosContext, true
}

// HasQosContext returns a boolean if a field has been set.
func (o *ClusterConfigProtoQoSMapping) HasQosContext() bool {
	if o != nil && o.QosContext != nil {
		return true
	}

	return false
}

// SetQosContext gets a reference to the given ClusterConfigProtoQoSMappingQoSContext and assigns it to the QosContext field.
func (o *ClusterConfigProtoQoSMapping) SetQosContext(v ClusterConfigProtoQoSMappingQoSContext) {
	o.QosContext = &v
}

func (o ClusterConfigProtoQoSMapping) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PrincipalId.IsSet() {
		toSerialize["principalId"] = o.PrincipalId.Get()
	}
	if o.QosContext != nil {
		toSerialize["qosContext"] = o.QosContext
	}
	return json.Marshal(toSerialize)
}

type NullableClusterConfigProtoQoSMapping struct {
	value *ClusterConfigProtoQoSMapping
	isSet bool
}

func (v NullableClusterConfigProtoQoSMapping) Get() *ClusterConfigProtoQoSMapping {
	return v.value
}

func (v *NullableClusterConfigProtoQoSMapping) Set(val *ClusterConfigProtoQoSMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterConfigProtoQoSMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterConfigProtoQoSMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterConfigProtoQoSMapping(val *ClusterConfigProtoQoSMapping) *NullableClusterConfigProtoQoSMapping {
	return &NullableClusterConfigProtoQoSMapping{value: val, isSet: true}
}

func (v NullableClusterConfigProtoQoSMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterConfigProtoQoSMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


