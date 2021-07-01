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

// HiveConnectParams Specifies an Object containing information about a registered Hive source.
type HiveConnectParams struct {
	// Specifies the entity id of the HDFS source for this Hive
	HdfsEntityId NullableInt64 `json:"hdfsEntityId,omitempty"`
	// Specifies the kerberos principal.
	KerberosPrincipal NullableString `json:"kerberosPrincipal,omitempty"`
	// Specifies the Hive metastore host.
	Metastore NullableString `json:"metastore,omitempty"`
	// Specifies the Hive metastore thrift Port
	ThriftPort NullableInt32 `json:"thriftPort,omitempty"`
}

// NewHiveConnectParams instantiates a new HiveConnectParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHiveConnectParams() *HiveConnectParams {
	this := HiveConnectParams{}
	return &this
}

// NewHiveConnectParamsWithDefaults instantiates a new HiveConnectParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHiveConnectParamsWithDefaults() *HiveConnectParams {
	this := HiveConnectParams{}
	return &this
}

// GetHdfsEntityId returns the HdfsEntityId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HiveConnectParams) GetHdfsEntityId() int64 {
	if o == nil || o.HdfsEntityId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.HdfsEntityId.Get()
}

// GetHdfsEntityIdOk returns a tuple with the HdfsEntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HiveConnectParams) GetHdfsEntityIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.HdfsEntityId.Get(), o.HdfsEntityId.IsSet()
}

// HasHdfsEntityId returns a boolean if a field has been set.
func (o *HiveConnectParams) HasHdfsEntityId() bool {
	if o != nil && o.HdfsEntityId.IsSet() {
		return true
	}

	return false
}

// SetHdfsEntityId gets a reference to the given NullableInt64 and assigns it to the HdfsEntityId field.
func (o *HiveConnectParams) SetHdfsEntityId(v int64) {
	o.HdfsEntityId.Set(&v)
}
// SetHdfsEntityIdNil sets the value for HdfsEntityId to be an explicit nil
func (o *HiveConnectParams) SetHdfsEntityIdNil() {
	o.HdfsEntityId.Set(nil)
}

// UnsetHdfsEntityId ensures that no value is present for HdfsEntityId, not even an explicit nil
func (o *HiveConnectParams) UnsetHdfsEntityId() {
	o.HdfsEntityId.Unset()
}

// GetKerberosPrincipal returns the KerberosPrincipal field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HiveConnectParams) GetKerberosPrincipal() string {
	if o == nil || o.KerberosPrincipal.Get() == nil {
		var ret string
		return ret
	}
	return *o.KerberosPrincipal.Get()
}

// GetKerberosPrincipalOk returns a tuple with the KerberosPrincipal field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HiveConnectParams) GetKerberosPrincipalOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.KerberosPrincipal.Get(), o.KerberosPrincipal.IsSet()
}

// HasKerberosPrincipal returns a boolean if a field has been set.
func (o *HiveConnectParams) HasKerberosPrincipal() bool {
	if o != nil && o.KerberosPrincipal.IsSet() {
		return true
	}

	return false
}

// SetKerberosPrincipal gets a reference to the given NullableString and assigns it to the KerberosPrincipal field.
func (o *HiveConnectParams) SetKerberosPrincipal(v string) {
	o.KerberosPrincipal.Set(&v)
}
// SetKerberosPrincipalNil sets the value for KerberosPrincipal to be an explicit nil
func (o *HiveConnectParams) SetKerberosPrincipalNil() {
	o.KerberosPrincipal.Set(nil)
}

// UnsetKerberosPrincipal ensures that no value is present for KerberosPrincipal, not even an explicit nil
func (o *HiveConnectParams) UnsetKerberosPrincipal() {
	o.KerberosPrincipal.Unset()
}

// GetMetastore returns the Metastore field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HiveConnectParams) GetMetastore() string {
	if o == nil || o.Metastore.Get() == nil {
		var ret string
		return ret
	}
	return *o.Metastore.Get()
}

// GetMetastoreOk returns a tuple with the Metastore field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HiveConnectParams) GetMetastoreOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Metastore.Get(), o.Metastore.IsSet()
}

// HasMetastore returns a boolean if a field has been set.
func (o *HiveConnectParams) HasMetastore() bool {
	if o != nil && o.Metastore.IsSet() {
		return true
	}

	return false
}

// SetMetastore gets a reference to the given NullableString and assigns it to the Metastore field.
func (o *HiveConnectParams) SetMetastore(v string) {
	o.Metastore.Set(&v)
}
// SetMetastoreNil sets the value for Metastore to be an explicit nil
func (o *HiveConnectParams) SetMetastoreNil() {
	o.Metastore.Set(nil)
}

// UnsetMetastore ensures that no value is present for Metastore, not even an explicit nil
func (o *HiveConnectParams) UnsetMetastore() {
	o.Metastore.Unset()
}

// GetThriftPort returns the ThriftPort field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HiveConnectParams) GetThriftPort() int32 {
	if o == nil || o.ThriftPort.Get() == nil {
		var ret int32
		return ret
	}
	return *o.ThriftPort.Get()
}

// GetThriftPortOk returns a tuple with the ThriftPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HiveConnectParams) GetThriftPortOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ThriftPort.Get(), o.ThriftPort.IsSet()
}

// HasThriftPort returns a boolean if a field has been set.
func (o *HiveConnectParams) HasThriftPort() bool {
	if o != nil && o.ThriftPort.IsSet() {
		return true
	}

	return false
}

// SetThriftPort gets a reference to the given NullableInt32 and assigns it to the ThriftPort field.
func (o *HiveConnectParams) SetThriftPort(v int32) {
	o.ThriftPort.Set(&v)
}
// SetThriftPortNil sets the value for ThriftPort to be an explicit nil
func (o *HiveConnectParams) SetThriftPortNil() {
	o.ThriftPort.Set(nil)
}

// UnsetThriftPort ensures that no value is present for ThriftPort, not even an explicit nil
func (o *HiveConnectParams) UnsetThriftPort() {
	o.ThriftPort.Unset()
}

func (o HiveConnectParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.HdfsEntityId.IsSet() {
		toSerialize["hdfsEntityId"] = o.HdfsEntityId.Get()
	}
	if o.KerberosPrincipal.IsSet() {
		toSerialize["kerberosPrincipal"] = o.KerberosPrincipal.Get()
	}
	if o.Metastore.IsSet() {
		toSerialize["metastore"] = o.Metastore.Get()
	}
	if o.ThriftPort.IsSet() {
		toSerialize["thriftPort"] = o.ThriftPort.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableHiveConnectParams struct {
	value *HiveConnectParams
	isSet bool
}

func (v NullableHiveConnectParams) Get() *HiveConnectParams {
	return v.value
}

func (v *NullableHiveConnectParams) Set(val *HiveConnectParams) {
	v.value = val
	v.isSet = true
}

func (v NullableHiveConnectParams) IsSet() bool {
	return v.isSet
}

func (v *NullableHiveConnectParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHiveConnectParams(val *HiveConnectParams) *NullableHiveConnectParams {
	return &NullableHiveConnectParams{value: val, isSet: true}
}

func (v NullableHiveConnectParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHiveConnectParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


