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

// CassandraProtectionGroupParamsAllOf struct for CassandraProtectionGroupParamsAllOf
type CassandraProtectionGroupParamsAllOf struct {
	// Only the specified data centers will be considered while taking backup. The keyspaces having replication strategy 'Simple' can be backed up only if all the datacenters for the cassandra cluster are specified. For any keyspace having replication strategy as 'Network', all the associated data centers should be specified.
	DataCenters *[]string `json:"dataCenters,omitempty"`
}

// NewCassandraProtectionGroupParamsAllOf instantiates a new CassandraProtectionGroupParamsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCassandraProtectionGroupParamsAllOf() *CassandraProtectionGroupParamsAllOf {
	this := CassandraProtectionGroupParamsAllOf{}
	return &this
}

// NewCassandraProtectionGroupParamsAllOfWithDefaults instantiates a new CassandraProtectionGroupParamsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCassandraProtectionGroupParamsAllOfWithDefaults() *CassandraProtectionGroupParamsAllOf {
	this := CassandraProtectionGroupParamsAllOf{}
	return &this
}

// GetDataCenters returns the DataCenters field value if set, zero value otherwise.
func (o *CassandraProtectionGroupParamsAllOf) GetDataCenters() []string {
	if o == nil || o.DataCenters == nil {
		var ret []string
		return ret
	}
	return *o.DataCenters
}

// GetDataCentersOk returns a tuple with the DataCenters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CassandraProtectionGroupParamsAllOf) GetDataCentersOk() (*[]string, bool) {
	if o == nil || o.DataCenters == nil {
		return nil, false
	}
	return o.DataCenters, true
}

// HasDataCenters returns a boolean if a field has been set.
func (o *CassandraProtectionGroupParamsAllOf) HasDataCenters() bool {
	if o != nil && o.DataCenters != nil {
		return true
	}

	return false
}

// SetDataCenters gets a reference to the given []string and assigns it to the DataCenters field.
func (o *CassandraProtectionGroupParamsAllOf) SetDataCenters(v []string) {
	o.DataCenters = &v
}

func (o CassandraProtectionGroupParamsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DataCenters != nil {
		toSerialize["dataCenters"] = o.DataCenters
	}
	return json.Marshal(toSerialize)
}

type NullableCassandraProtectionGroupParamsAllOf struct {
	value *CassandraProtectionGroupParamsAllOf
	isSet bool
}

func (v NullableCassandraProtectionGroupParamsAllOf) Get() *CassandraProtectionGroupParamsAllOf {
	return v.value
}

func (v *NullableCassandraProtectionGroupParamsAllOf) Set(val *CassandraProtectionGroupParamsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCassandraProtectionGroupParamsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCassandraProtectionGroupParamsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCassandraProtectionGroupParamsAllOf(val *CassandraProtectionGroupParamsAllOf) *NullableCassandraProtectionGroupParamsAllOf {
	return &NullableCassandraProtectionGroupParamsAllOf{value: val, isSet: true}
}

func (v NullableCassandraProtectionGroupParamsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCassandraProtectionGroupParamsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


