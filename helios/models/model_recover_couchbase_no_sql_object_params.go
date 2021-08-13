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

// RecoverCouchbaseNoSqlObjectParams Specifies the fully qualified object name and other attributes of each object to be recovered.
type RecoverCouchbaseNoSqlObjectParams struct {
	// Specifies the fully qualified name of the object to be restored.
	ObjectName NullableString `json:"objectName"`
	// Specifies the new name to which the object should be renamed, at the time of recovery.
	RenameTo NullableString `json:"renameTo,omitempty"`
}

// NewRecoverCouchbaseNoSqlObjectParams instantiates a new RecoverCouchbaseNoSqlObjectParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverCouchbaseNoSqlObjectParams(objectName NullableString) *RecoverCouchbaseNoSqlObjectParams {
	this := RecoverCouchbaseNoSqlObjectParams{}
	this.ObjectName = objectName
	return &this
}

// NewRecoverCouchbaseNoSqlObjectParamsWithDefaults instantiates a new RecoverCouchbaseNoSqlObjectParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverCouchbaseNoSqlObjectParamsWithDefaults() *RecoverCouchbaseNoSqlObjectParams {
	this := RecoverCouchbaseNoSqlObjectParams{}
	return &this
}

// GetObjectName returns the ObjectName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RecoverCouchbaseNoSqlObjectParams) GetObjectName() string {
	if o == nil || o.ObjectName.Get() == nil {
		var ret string
		return ret
	}

	return *o.ObjectName.Get()
}

// GetObjectNameOk returns a tuple with the ObjectName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverCouchbaseNoSqlObjectParams) GetObjectNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ObjectName.Get(), o.ObjectName.IsSet()
}

// SetObjectName sets field value
func (o *RecoverCouchbaseNoSqlObjectParams) SetObjectName(v string) {
	o.ObjectName.Set(&v)
}

// GetRenameTo returns the RenameTo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverCouchbaseNoSqlObjectParams) GetRenameTo() string {
	if o == nil || o.RenameTo.Get() == nil {
		var ret string
		return ret
	}
	return *o.RenameTo.Get()
}

// GetRenameToOk returns a tuple with the RenameTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverCouchbaseNoSqlObjectParams) GetRenameToOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RenameTo.Get(), o.RenameTo.IsSet()
}

// HasRenameTo returns a boolean if a field has been set.
func (o *RecoverCouchbaseNoSqlObjectParams) HasRenameTo() bool {
	if o != nil && o.RenameTo.IsSet() {
		return true
	}

	return false
}

// SetRenameTo gets a reference to the given NullableString and assigns it to the RenameTo field.
func (o *RecoverCouchbaseNoSqlObjectParams) SetRenameTo(v string) {
	o.RenameTo.Set(&v)
}
// SetRenameToNil sets the value for RenameTo to be an explicit nil
func (o *RecoverCouchbaseNoSqlObjectParams) SetRenameToNil() {
	o.RenameTo.Set(nil)
}

// UnsetRenameTo ensures that no value is present for RenameTo, not even an explicit nil
func (o *RecoverCouchbaseNoSqlObjectParams) UnsetRenameTo() {
	o.RenameTo.Unset()
}

func (o RecoverCouchbaseNoSqlObjectParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["objectName"] = o.ObjectName.Get()
	}
	if o.RenameTo.IsSet() {
		toSerialize["renameTo"] = o.RenameTo.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRecoverCouchbaseNoSqlObjectParams struct {
	value *RecoverCouchbaseNoSqlObjectParams
	isSet bool
}

func (v NullableRecoverCouchbaseNoSqlObjectParams) Get() *RecoverCouchbaseNoSqlObjectParams {
	return v.value
}

func (v *NullableRecoverCouchbaseNoSqlObjectParams) Set(val *RecoverCouchbaseNoSqlObjectParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverCouchbaseNoSqlObjectParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverCouchbaseNoSqlObjectParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverCouchbaseNoSqlObjectParams(val *RecoverCouchbaseNoSqlObjectParams) *NullableRecoverCouchbaseNoSqlObjectParams {
	return &NullableRecoverCouchbaseNoSqlObjectParams{value: val, isSet: true}
}

func (v NullableRecoverCouchbaseNoSqlObjectParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverCouchbaseNoSqlObjectParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


