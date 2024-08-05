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

// checks if the RecoverPureSanVolumeNewSourceConfigSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecoverPureSanVolumeNewSourceConfigSource{}

// RecoverPureSanVolumeNewSourceConfigSource Specifies the id of the new target parent source to recover the Pure SAN Volume to. This field must be specified if recoverToNewSource is true.
type RecoverPureSanVolumeNewSourceConfigSource struct {
	// Specifies the id of the object.
	Id NullableInt64 `json:"id"`
	// Specifies the name of the object.
	Name NullableString `json:"name,omitempty"`
}

type _RecoverPureSanVolumeNewSourceConfigSource RecoverPureSanVolumeNewSourceConfigSource

// NewRecoverPureSanVolumeNewSourceConfigSource instantiates a new RecoverPureSanVolumeNewSourceConfigSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverPureSanVolumeNewSourceConfigSource(id NullableInt64) *RecoverPureSanVolumeNewSourceConfigSource {
	this := RecoverPureSanVolumeNewSourceConfigSource{}
	this.Id = id
	return &this
}

// NewRecoverPureSanVolumeNewSourceConfigSourceWithDefaults instantiates a new RecoverPureSanVolumeNewSourceConfigSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverPureSanVolumeNewSourceConfigSourceWithDefaults() *RecoverPureSanVolumeNewSourceConfigSource {
	this := RecoverPureSanVolumeNewSourceConfigSource{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *RecoverPureSanVolumeNewSourceConfigSource) GetId() int64 {
	if o == nil || o.Id.Get() == nil {
		var ret int64
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverPureSanVolumeNewSourceConfigSource) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *RecoverPureSanVolumeNewSourceConfigSource) SetId(v int64) {
	o.Id.Set(&v)
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoverPureSanVolumeNewSourceConfigSource) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoverPureSanVolumeNewSourceConfigSource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *RecoverPureSanVolumeNewSourceConfigSource) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *RecoverPureSanVolumeNewSourceConfigSource) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *RecoverPureSanVolumeNewSourceConfigSource) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *RecoverPureSanVolumeNewSourceConfigSource) UnsetName() {
	o.Name.Unset()
}

func (o RecoverPureSanVolumeNewSourceConfigSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecoverPureSanVolumeNewSourceConfigSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id.Get()
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return toSerialize, nil
}

func (o *RecoverPureSanVolumeNewSourceConfigSource) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varRecoverPureSanVolumeNewSourceConfigSource := _RecoverPureSanVolumeNewSourceConfigSource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRecoverPureSanVolumeNewSourceConfigSource)

	if err != nil {
		return err
	}

	*o = RecoverPureSanVolumeNewSourceConfigSource(varRecoverPureSanVolumeNewSourceConfigSource)

	return err
}

type NullableRecoverPureSanVolumeNewSourceConfigSource struct {
	value *RecoverPureSanVolumeNewSourceConfigSource
	isSet bool
}

func (v NullableRecoverPureSanVolumeNewSourceConfigSource) Get() *RecoverPureSanVolumeNewSourceConfigSource {
	return v.value
}

func (v *NullableRecoverPureSanVolumeNewSourceConfigSource) Set(val *RecoverPureSanVolumeNewSourceConfigSource) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverPureSanVolumeNewSourceConfigSource) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverPureSanVolumeNewSourceConfigSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverPureSanVolumeNewSourceConfigSource(val *RecoverPureSanVolumeNewSourceConfigSource) *NullableRecoverPureSanVolumeNewSourceConfigSource {
	return &NullableRecoverPureSanVolumeNewSourceConfigSource{value: val, isSet: true}
}

func (v NullableRecoverPureSanVolumeNewSourceConfigSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverPureSanVolumeNewSourceConfigSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


