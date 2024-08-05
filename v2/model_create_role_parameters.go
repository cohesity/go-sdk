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

// checks if the CreateRoleParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRoleParameters{}

// CreateRoleParameters Specifies the parameters to create a Role.
type CreateRoleParameters struct {
	// Specifies the description message for the Role.
	Description NullableString `json:"description,omitempty"`
	// Specifies the list of Privileges of the Role.
	Privileges []string `json:"privileges"`
	// Specifies the Role name.
	Name NullableString `json:"name"`
}

type _CreateRoleParameters CreateRoleParameters

// NewCreateRoleParameters instantiates a new CreateRoleParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRoleParameters(privileges []string, name NullableString) *CreateRoleParameters {
	this := CreateRoleParameters{}
	this.Privileges = privileges
	this.Name = name
	return &this
}

// NewCreateRoleParametersWithDefaults instantiates a new CreateRoleParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRoleParametersWithDefaults() *CreateRoleParameters {
	this := CreateRoleParameters{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateRoleParameters) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateRoleParameters) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateRoleParameters) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *CreateRoleParameters) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *CreateRoleParameters) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *CreateRoleParameters) UnsetDescription() {
	o.Description.Unset()
}

// GetPrivileges returns the Privileges field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *CreateRoleParameters) GetPrivileges() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Privileges
}

// GetPrivilegesOk returns a tuple with the Privileges field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateRoleParameters) GetPrivilegesOk() ([]string, bool) {
	if o == nil || IsNil(o.Privileges) {
		return nil, false
	}
	return o.Privileges, true
}

// SetPrivileges sets field value
func (o *CreateRoleParameters) SetPrivileges(v []string) {
	o.Privileges = v
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreateRoleParameters) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateRoleParameters) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *CreateRoleParameters) SetName(v string) {
	o.Name.Set(&v)
}

func (o CreateRoleParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRoleParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Privileges != nil {
		toSerialize["privileges"] = o.Privileges
	}
	toSerialize["name"] = o.Name.Get()
	return toSerialize, nil
}

func (o *CreateRoleParameters) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"privileges",
		"name",
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

	varCreateRoleParameters := _CreateRoleParameters{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateRoleParameters)

	if err != nil {
		return err
	}

	*o = CreateRoleParameters(varCreateRoleParameters)

	return err
}

type NullableCreateRoleParameters struct {
	value *CreateRoleParameters
	isSet bool
}

func (v NullableCreateRoleParameters) Get() *CreateRoleParameters {
	return v.value
}

func (v *NullableCreateRoleParameters) Set(val *CreateRoleParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRoleParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRoleParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRoleParameters(val *CreateRoleParameters) *NullableCreateRoleParameters {
	return &NullableCreateRoleParameters{value: val, isSet: true}
}

func (v NullableCreateRoleParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRoleParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


