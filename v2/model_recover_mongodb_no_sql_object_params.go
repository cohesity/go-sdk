/*
Cohesity REST API

Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the RecoverMongodbNoSqlObjectParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecoverMongodbNoSqlObjectParams{}

// RecoverMongodbNoSqlObjectParams Specifies the fully qualified object name and other attributes of each object to be recovered.
type RecoverMongodbNoSqlObjectParams struct {
	// Specifies the fully qualified name of the object to be restored.
	ObjectName NullableString `json:"objectName"`
	// Specifies the properties to be applied to the object at the time of recovery.
	ObjectProperties []NoSqlObjectProperty `json:"objectProperties,omitempty"`
	// Specifies the new name to which the object should be renamed. at the time of recovery.
	RenameTo NullableString `json:"renameTo,omitempty"`
}

type _RecoverMongodbNoSqlObjectParams RecoverMongodbNoSqlObjectParams

// NewRecoverMongodbNoSqlObjectParams instantiates a new RecoverMongodbNoSqlObjectParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverMongodbNoSqlObjectParams(objectName NullableString) *RecoverMongodbNoSqlObjectParams {
	this := RecoverMongodbNoSqlObjectParams{}
	this.ObjectName = objectName
	return &this
}

// NewRecoverMongodbNoSqlObjectParamsWithDefaults instantiates a new RecoverMongodbNoSqlObjectParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverMongodbNoSqlObjectParamsWithDefaults() *RecoverMongodbNoSqlObjectParams {
	this := RecoverMongodbNoSqlObjectParams{}
	return &this
}

// GetObjectName returns the ObjectName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RecoverMongodbNoSqlObjectParams) GetObjectName() string {
	if o == nil || o.ObjectName.Get() == nil {
		var ret string
		return ret
	}

	return *o.ObjectName.Get()
}

// GetObjectNameOk returns a tuple with the ObjectName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverMongodbNoSqlObjectParams) GetObjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ObjectName.Get(), o.ObjectName.IsSet()
}

// SetObjectName sets field value
func (o *RecoverMongodbNoSqlObjectParams) SetObjectName(v string) {
	o.ObjectName.Set(&v)
}

// GetObjectProperties returns the ObjectProperties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverMongodbNoSqlObjectParams) GetObjectProperties() []NoSqlObjectProperty {
	if o == nil {
		var ret []NoSqlObjectProperty
		return ret
	}
	return o.ObjectProperties
}

// GetObjectPropertiesOk returns a tuple with the ObjectProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverMongodbNoSqlObjectParams) GetObjectPropertiesOk() ([]NoSqlObjectProperty, bool) {
	if o == nil || IsNil(o.ObjectProperties) {
		return nil, false
	}
	return o.ObjectProperties, true
}

// HasObjectProperties returns a boolean if a field has been set.
func (o *RecoverMongodbNoSqlObjectParams) HasObjectProperties() bool {
	if o != nil && !IsNil(o.ObjectProperties) {
		return true
	}

	return false
}

// SetObjectProperties gets a reference to the given []NoSqlObjectProperty and assigns it to the ObjectProperties field.
func (o *RecoverMongodbNoSqlObjectParams) SetObjectProperties(v []NoSqlObjectProperty) {
	o.ObjectProperties = v
}

// GetRenameTo returns the RenameTo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverMongodbNoSqlObjectParams) GetRenameTo() string {
	if o == nil || IsNil(o.RenameTo.Get()) {
		var ret string
		return ret
	}
	return *o.RenameTo.Get()
}

// GetRenameToOk returns a tuple with the RenameTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverMongodbNoSqlObjectParams) GetRenameToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RenameTo.Get(), o.RenameTo.IsSet()
}

// HasRenameTo returns a boolean if a field has been set.
func (o *RecoverMongodbNoSqlObjectParams) HasRenameTo() bool {
	if o != nil && o.RenameTo.IsSet() {
		return true
	}

	return false
}

// SetRenameTo gets a reference to the given NullableString and assigns it to the RenameTo field.
func (o *RecoverMongodbNoSqlObjectParams) SetRenameTo(v string) {
	o.RenameTo.Set(&v)
}
// SetRenameToNil sets the value for RenameTo to be an explicit nil
func (o *RecoverMongodbNoSqlObjectParams) SetRenameToNil() {
	o.RenameTo.Set(nil)
}

// UnsetRenameTo ensures that no value is present for RenameTo, not even an explicit nil
func (o *RecoverMongodbNoSqlObjectParams) UnsetRenameTo() {
	o.RenameTo.Unset()
}

func (o RecoverMongodbNoSqlObjectParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecoverMongodbNoSqlObjectParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objectName"] = o.ObjectName.Get()
	if o.ObjectProperties != nil {
		toSerialize["objectProperties"] = o.ObjectProperties
	}
	if o.RenameTo.IsSet() {
		toSerialize["renameTo"] = o.RenameTo.Get()
	}
	return toSerialize, nil
}

func (o *RecoverMongodbNoSqlObjectParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objectName",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRecoverMongodbNoSqlObjectParams := _RecoverMongodbNoSqlObjectParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRecoverMongodbNoSqlObjectParams)

	if err != nil {
		return err
	}

	*o = RecoverMongodbNoSqlObjectParams(varRecoverMongodbNoSqlObjectParams)

	return err
}

type NullableRecoverMongodbNoSqlObjectParams struct {
	value *RecoverMongodbNoSqlObjectParams
	isSet bool
}

func (v NullableRecoverMongodbNoSqlObjectParams) Get() *RecoverMongodbNoSqlObjectParams {
	return v.value
}

func (v *NullableRecoverMongodbNoSqlObjectParams) Set(val *RecoverMongodbNoSqlObjectParams) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverMongodbNoSqlObjectParams) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverMongodbNoSqlObjectParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverMongodbNoSqlObjectParams(val *RecoverMongodbNoSqlObjectParams) *NullableRecoverMongodbNoSqlObjectParams {
	return &NullableRecoverMongodbNoSqlObjectParams{value: val, isSet: true}
}

func (v NullableRecoverMongodbNoSqlObjectParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverMongodbNoSqlObjectParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


