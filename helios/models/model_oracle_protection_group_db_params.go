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

// OracleProtectionGroupDbParams Specifies the parameters of individual databases to create Oracle Protection Group.
type OracleProtectionGroupDbParams struct {
	// Specifies the id of the Oracle database.
	DatabaseId NullableInt64 `json:"databaseId,omitempty"`
	// Specifies the name of the Oracle database.
	DatabaseName NullableString `json:"databaseName,omitempty"`
	// Specifies the Oracle database node channels info. If not specified, the default values assigned by the server are applied to all the databases.
	DbChannels []OracleDbChannel `json:"dbChannels,omitempty"`
}

// NewOracleProtectionGroupDbParams instantiates a new OracleProtectionGroupDbParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOracleProtectionGroupDbParams() *OracleProtectionGroupDbParams {
	this := OracleProtectionGroupDbParams{}
	return &this
}

// NewOracleProtectionGroupDbParamsWithDefaults instantiates a new OracleProtectionGroupDbParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOracleProtectionGroupDbParamsWithDefaults() *OracleProtectionGroupDbParams {
	this := OracleProtectionGroupDbParams{}
	return &this
}

// GetDatabaseId returns the DatabaseId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OracleProtectionGroupDbParams) GetDatabaseId() int64 {
	if o == nil || o.DatabaseId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.DatabaseId.Get()
}

// GetDatabaseIdOk returns a tuple with the DatabaseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OracleProtectionGroupDbParams) GetDatabaseIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DatabaseId.Get(), o.DatabaseId.IsSet()
}

// HasDatabaseId returns a boolean if a field has been set.
func (o *OracleProtectionGroupDbParams) HasDatabaseId() bool {
	if o != nil && o.DatabaseId.IsSet() {
		return true
	}

	return false
}

// SetDatabaseId gets a reference to the given NullableInt64 and assigns it to the DatabaseId field.
func (o *OracleProtectionGroupDbParams) SetDatabaseId(v int64) {
	o.DatabaseId.Set(&v)
}
// SetDatabaseIdNil sets the value for DatabaseId to be an explicit nil
func (o *OracleProtectionGroupDbParams) SetDatabaseIdNil() {
	o.DatabaseId.Set(nil)
}

// UnsetDatabaseId ensures that no value is present for DatabaseId, not even an explicit nil
func (o *OracleProtectionGroupDbParams) UnsetDatabaseId() {
	o.DatabaseId.Unset()
}

// GetDatabaseName returns the DatabaseName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OracleProtectionGroupDbParams) GetDatabaseName() string {
	if o == nil || o.DatabaseName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DatabaseName.Get()
}

// GetDatabaseNameOk returns a tuple with the DatabaseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OracleProtectionGroupDbParams) GetDatabaseNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DatabaseName.Get(), o.DatabaseName.IsSet()
}

// HasDatabaseName returns a boolean if a field has been set.
func (o *OracleProtectionGroupDbParams) HasDatabaseName() bool {
	if o != nil && o.DatabaseName.IsSet() {
		return true
	}

	return false
}

// SetDatabaseName gets a reference to the given NullableString and assigns it to the DatabaseName field.
func (o *OracleProtectionGroupDbParams) SetDatabaseName(v string) {
	o.DatabaseName.Set(&v)
}
// SetDatabaseNameNil sets the value for DatabaseName to be an explicit nil
func (o *OracleProtectionGroupDbParams) SetDatabaseNameNil() {
	o.DatabaseName.Set(nil)
}

// UnsetDatabaseName ensures that no value is present for DatabaseName, not even an explicit nil
func (o *OracleProtectionGroupDbParams) UnsetDatabaseName() {
	o.DatabaseName.Unset()
}

// GetDbChannels returns the DbChannels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OracleProtectionGroupDbParams) GetDbChannels() []OracleDbChannel {
	if o == nil  {
		var ret []OracleDbChannel
		return ret
	}
	return o.DbChannels
}

// GetDbChannelsOk returns a tuple with the DbChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OracleProtectionGroupDbParams) GetDbChannelsOk() (*[]OracleDbChannel, bool) {
	if o == nil || o.DbChannels == nil {
		return nil, false
	}
	return &o.DbChannels, true
}

// HasDbChannels returns a boolean if a field has been set.
func (o *OracleProtectionGroupDbParams) HasDbChannels() bool {
	if o != nil && o.DbChannels != nil {
		return true
	}

	return false
}

// SetDbChannels gets a reference to the given []OracleDbChannel and assigns it to the DbChannels field.
func (o *OracleProtectionGroupDbParams) SetDbChannels(v []OracleDbChannel) {
	o.DbChannels = v
}

func (o OracleProtectionGroupDbParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DatabaseId.IsSet() {
		toSerialize["databaseId"] = o.DatabaseId.Get()
	}
	if o.DatabaseName.IsSet() {
		toSerialize["databaseName"] = o.DatabaseName.Get()
	}
	if o.DbChannels != nil {
		toSerialize["dbChannels"] = o.DbChannels
	}
	return json.Marshal(toSerialize)
}

type NullableOracleProtectionGroupDbParams struct {
	value *OracleProtectionGroupDbParams
	isSet bool
}

func (v NullableOracleProtectionGroupDbParams) Get() *OracleProtectionGroupDbParams {
	return v.value
}

func (v *NullableOracleProtectionGroupDbParams) Set(val *OracleProtectionGroupDbParams) {
	v.value = val
	v.isSet = true
}

func (v NullableOracleProtectionGroupDbParams) IsSet() bool {
	return v.isSet
}

func (v *NullableOracleProtectionGroupDbParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOracleProtectionGroupDbParams(val *OracleProtectionGroupDbParams) *NullableOracleProtectionGroupDbParams {
	return &NullableOracleProtectionGroupDbParams{value: val, isSet: true}
}

func (v NullableOracleProtectionGroupDbParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOracleProtectionGroupDbParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


