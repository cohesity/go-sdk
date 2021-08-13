/*
 * Cohesity REST API
 *
 * Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

// package onprem
package models

import (
	"encoding/json"
	. "github.com/cohesity/go-sdk/onprem/utils"
	"fmt"
)

var _ = NullableBool{}

// AuditLogsEntityTypes Specifies entity types of audit logs.
type AuditLogsEntityTypes struct {
	// Specifies a list of audit logs entity types.
	EntityTypes []string `json:"entityTypes,omitempty"`
}

// NewAuditLogsEntityTypes instantiates a new AuditLogsEntityTypes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditLogsEntityTypes() *AuditLogsEntityTypes {
	this := AuditLogsEntityTypes{}
	return &this
}

// NewAuditLogsEntityTypesWithDefaults instantiates a new AuditLogsEntityTypes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditLogsEntityTypesWithDefaults() *AuditLogsEntityTypes {
	this := AuditLogsEntityTypes{}
	return &this
}

// GetEntityTypes returns the EntityTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditLogsEntityTypes) GetEntityTypes() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.EntityTypes
}

// GetEntityTypesOk returns a tuple with the EntityTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditLogsEntityTypes) GetEntityTypesOk() (*[]string, bool) {
	if o == nil || o.EntityTypes == nil {
		return nil, false
	}
	return &o.EntityTypes, true
}

// HasEntityTypes returns a boolean if a field has been set.
func (o *AuditLogsEntityTypes) HasEntityTypes() bool {
	if o != nil && o.EntityTypes != nil {
		return true
	}

	return false
}

// SetEntityTypes gets a reference to the given []string and assigns it to the EntityTypes field.
func (o *AuditLogsEntityTypes) SetEntityTypes(v []string) {
	o.EntityTypes = v
}

func (o AuditLogsEntityTypes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EntityTypes != nil {
		toSerialize["entityTypes"] = o.EntityTypes
	}
	return json.Marshal(toSerialize)
}

type NullableAuditLogsEntityTypes struct {
	value *AuditLogsEntityTypes
	isSet bool
}

func (v NullableAuditLogsEntityTypes) Get() *AuditLogsEntityTypes {
	return v.value
}

func (v *NullableAuditLogsEntityTypes) Set(val *AuditLogsEntityTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditLogsEntityTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditLogsEntityTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditLogsEntityTypes(val *AuditLogsEntityTypes) *NullableAuditLogsEntityTypes {
	return &NullableAuditLogsEntityTypes{value: val, isSet: true}
}

func (v NullableAuditLogsEntityTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditLogsEntityTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o AuditLogsEntityTypes) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}