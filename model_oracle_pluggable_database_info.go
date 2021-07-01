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

// OraclePluggableDatabaseInfo Specifies the informatiomn about the pluggable database. A Pluggabele Database(PDB) is a portable collection of schemas, schema objects, and nonschema objects that appears to an Oracle Net client as a non-CDB.
type OraclePluggableDatabaseInfo struct {
	// Specifies the ID of the Pluggable Database.
	DatabaseId NullableString `json:"databaseId,omitempty"`
	// Speicifes the name of the Pluggable Database.
	Name NullableString `json:"name,omitempty"`
	// Specifies the OPEN_MODE information. Specifies the OPEN_MODE type for the Oracle database. 'kMounted' indicates that the database is open in Mounted mode. 'kReadWrite' indicates that the database is open in R/W mode. 'kReadOnly' indicates that the database is open in ReadOnly mode. 'kMigrate' inidcates that the database is open in Migrate mode.
	OpenMode NullableString `json:"openMode,omitempty"`
	// Specifies the Size in Bytes of the Pluggable Database.
	SizeBytes NullableInt64 `json:"sizeBytes,omitempty"`
}

// NewOraclePluggableDatabaseInfo instantiates a new OraclePluggableDatabaseInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOraclePluggableDatabaseInfo() *OraclePluggableDatabaseInfo {
	this := OraclePluggableDatabaseInfo{}
	return &this
}

// NewOraclePluggableDatabaseInfoWithDefaults instantiates a new OraclePluggableDatabaseInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOraclePluggableDatabaseInfoWithDefaults() *OraclePluggableDatabaseInfo {
	this := OraclePluggableDatabaseInfo{}
	return &this
}

// GetDatabaseId returns the DatabaseId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OraclePluggableDatabaseInfo) GetDatabaseId() string {
	if o == nil || o.DatabaseId.Get() == nil {
		var ret string
		return ret
	}
	return *o.DatabaseId.Get()
}

// GetDatabaseIdOk returns a tuple with the DatabaseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OraclePluggableDatabaseInfo) GetDatabaseIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DatabaseId.Get(), o.DatabaseId.IsSet()
}

// HasDatabaseId returns a boolean if a field has been set.
func (o *OraclePluggableDatabaseInfo) HasDatabaseId() bool {
	if o != nil && o.DatabaseId.IsSet() {
		return true
	}

	return false
}

// SetDatabaseId gets a reference to the given NullableString and assigns it to the DatabaseId field.
func (o *OraclePluggableDatabaseInfo) SetDatabaseId(v string) {
	o.DatabaseId.Set(&v)
}
// SetDatabaseIdNil sets the value for DatabaseId to be an explicit nil
func (o *OraclePluggableDatabaseInfo) SetDatabaseIdNil() {
	o.DatabaseId.Set(nil)
}

// UnsetDatabaseId ensures that no value is present for DatabaseId, not even an explicit nil
func (o *OraclePluggableDatabaseInfo) UnsetDatabaseId() {
	o.DatabaseId.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OraclePluggableDatabaseInfo) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OraclePluggableDatabaseInfo) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *OraclePluggableDatabaseInfo) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *OraclePluggableDatabaseInfo) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *OraclePluggableDatabaseInfo) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *OraclePluggableDatabaseInfo) UnsetName() {
	o.Name.Unset()
}

// GetOpenMode returns the OpenMode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OraclePluggableDatabaseInfo) GetOpenMode() string {
	if o == nil || o.OpenMode.Get() == nil {
		var ret string
		return ret
	}
	return *o.OpenMode.Get()
}

// GetOpenModeOk returns a tuple with the OpenMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OraclePluggableDatabaseInfo) GetOpenModeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OpenMode.Get(), o.OpenMode.IsSet()
}

// HasOpenMode returns a boolean if a field has been set.
func (o *OraclePluggableDatabaseInfo) HasOpenMode() bool {
	if o != nil && o.OpenMode.IsSet() {
		return true
	}

	return false
}

// SetOpenMode gets a reference to the given NullableString and assigns it to the OpenMode field.
func (o *OraclePluggableDatabaseInfo) SetOpenMode(v string) {
	o.OpenMode.Set(&v)
}
// SetOpenModeNil sets the value for OpenMode to be an explicit nil
func (o *OraclePluggableDatabaseInfo) SetOpenModeNil() {
	o.OpenMode.Set(nil)
}

// UnsetOpenMode ensures that no value is present for OpenMode, not even an explicit nil
func (o *OraclePluggableDatabaseInfo) UnsetOpenMode() {
	o.OpenMode.Unset()
}

// GetSizeBytes returns the SizeBytes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OraclePluggableDatabaseInfo) GetSizeBytes() int64 {
	if o == nil || o.SizeBytes.Get() == nil {
		var ret int64
		return ret
	}
	return *o.SizeBytes.Get()
}

// GetSizeBytesOk returns a tuple with the SizeBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OraclePluggableDatabaseInfo) GetSizeBytesOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SizeBytes.Get(), o.SizeBytes.IsSet()
}

// HasSizeBytes returns a boolean if a field has been set.
func (o *OraclePluggableDatabaseInfo) HasSizeBytes() bool {
	if o != nil && o.SizeBytes.IsSet() {
		return true
	}

	return false
}

// SetSizeBytes gets a reference to the given NullableInt64 and assigns it to the SizeBytes field.
func (o *OraclePluggableDatabaseInfo) SetSizeBytes(v int64) {
	o.SizeBytes.Set(&v)
}
// SetSizeBytesNil sets the value for SizeBytes to be an explicit nil
func (o *OraclePluggableDatabaseInfo) SetSizeBytesNil() {
	o.SizeBytes.Set(nil)
}

// UnsetSizeBytes ensures that no value is present for SizeBytes, not even an explicit nil
func (o *OraclePluggableDatabaseInfo) UnsetSizeBytes() {
	o.SizeBytes.Unset()
}

func (o OraclePluggableDatabaseInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DatabaseId.IsSet() {
		toSerialize["databaseId"] = o.DatabaseId.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.OpenMode.IsSet() {
		toSerialize["openMode"] = o.OpenMode.Get()
	}
	if o.SizeBytes.IsSet() {
		toSerialize["sizeBytes"] = o.SizeBytes.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableOraclePluggableDatabaseInfo struct {
	value *OraclePluggableDatabaseInfo
	isSet bool
}

func (v NullableOraclePluggableDatabaseInfo) Get() *OraclePluggableDatabaseInfo {
	return v.value
}

func (v *NullableOraclePluggableDatabaseInfo) Set(val *OraclePluggableDatabaseInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableOraclePluggableDatabaseInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableOraclePluggableDatabaseInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOraclePluggableDatabaseInfo(val *OraclePluggableDatabaseInfo) *NullableOraclePluggableDatabaseInfo {
	return &NullableOraclePluggableDatabaseInfo{value: val, isSet: true}
}

func (v NullableOraclePluggableDatabaseInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOraclePluggableDatabaseInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


