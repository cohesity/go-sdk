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

// checks if the CreateLdapParamsSimpleAuthParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateLdapParamsSimpleAuthParams{}

// CreateLdapParamsSimpleAuthParams Specifies the parameters for LDAP with 'Simple' authentication type.
type CreateLdapParamsSimpleAuthParams struct {
	// Specifies whether to use SSL for LDAP connections.
	UseSsl NullableBool `json:"useSsl,omitempty"`
	// Specifies the user distinguished name that is used for LDAP authentication.
	UserDistinguishedName NullableString `json:"userDistinguishedName"`
	// Specifies the user password that is used for LDAP authentication.
	UserPassword NullableString `json:"userPassword,omitempty"`
}

type _CreateLdapParamsSimpleAuthParams CreateLdapParamsSimpleAuthParams

// NewCreateLdapParamsSimpleAuthParams instantiates a new CreateLdapParamsSimpleAuthParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLdapParamsSimpleAuthParams(userDistinguishedName NullableString) *CreateLdapParamsSimpleAuthParams {
	this := CreateLdapParamsSimpleAuthParams{}
	this.UserDistinguishedName = userDistinguishedName
	return &this
}

// NewCreateLdapParamsSimpleAuthParamsWithDefaults instantiates a new CreateLdapParamsSimpleAuthParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLdapParamsSimpleAuthParamsWithDefaults() *CreateLdapParamsSimpleAuthParams {
	this := CreateLdapParamsSimpleAuthParams{}
	return &this
}

// GetUseSsl returns the UseSsl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateLdapParamsSimpleAuthParams) GetUseSsl() bool {
	if o == nil || IsNil(o.UseSsl.Get()) {
		var ret bool
		return ret
	}
	return *o.UseSsl.Get()
}

// GetUseSslOk returns a tuple with the UseSsl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateLdapParamsSimpleAuthParams) GetUseSslOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.UseSsl.Get(), o.UseSsl.IsSet()
}

// HasUseSsl returns a boolean if a field has been set.
func (o *CreateLdapParamsSimpleAuthParams) HasUseSsl() bool {
	if o != nil && o.UseSsl.IsSet() {
		return true
	}

	return false
}

// SetUseSsl gets a reference to the given NullableBool and assigns it to the UseSsl field.
func (o *CreateLdapParamsSimpleAuthParams) SetUseSsl(v bool) {
	o.UseSsl.Set(&v)
}
// SetUseSslNil sets the value for UseSsl to be an explicit nil
func (o *CreateLdapParamsSimpleAuthParams) SetUseSslNil() {
	o.UseSsl.Set(nil)
}

// UnsetUseSsl ensures that no value is present for UseSsl, not even an explicit nil
func (o *CreateLdapParamsSimpleAuthParams) UnsetUseSsl() {
	o.UseSsl.Unset()
}

// GetUserDistinguishedName returns the UserDistinguishedName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreateLdapParamsSimpleAuthParams) GetUserDistinguishedName() string {
	if o == nil || o.UserDistinguishedName.Get() == nil {
		var ret string
		return ret
	}

	return *o.UserDistinguishedName.Get()
}

// GetUserDistinguishedNameOk returns a tuple with the UserDistinguishedName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateLdapParamsSimpleAuthParams) GetUserDistinguishedNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserDistinguishedName.Get(), o.UserDistinguishedName.IsSet()
}

// SetUserDistinguishedName sets field value
func (o *CreateLdapParamsSimpleAuthParams) SetUserDistinguishedName(v string) {
	o.UserDistinguishedName.Set(&v)
}

// GetUserPassword returns the UserPassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateLdapParamsSimpleAuthParams) GetUserPassword() string {
	if o == nil || IsNil(o.UserPassword.Get()) {
		var ret string
		return ret
	}
	return *o.UserPassword.Get()
}

// GetUserPasswordOk returns a tuple with the UserPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateLdapParamsSimpleAuthParams) GetUserPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserPassword.Get(), o.UserPassword.IsSet()
}

// HasUserPassword returns a boolean if a field has been set.
func (o *CreateLdapParamsSimpleAuthParams) HasUserPassword() bool {
	if o != nil && o.UserPassword.IsSet() {
		return true
	}

	return false
}

// SetUserPassword gets a reference to the given NullableString and assigns it to the UserPassword field.
func (o *CreateLdapParamsSimpleAuthParams) SetUserPassword(v string) {
	o.UserPassword.Set(&v)
}
// SetUserPasswordNil sets the value for UserPassword to be an explicit nil
func (o *CreateLdapParamsSimpleAuthParams) SetUserPasswordNil() {
	o.UserPassword.Set(nil)
}

// UnsetUserPassword ensures that no value is present for UserPassword, not even an explicit nil
func (o *CreateLdapParamsSimpleAuthParams) UnsetUserPassword() {
	o.UserPassword.Unset()
}

func (o CreateLdapParamsSimpleAuthParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateLdapParamsSimpleAuthParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.UseSsl.IsSet() {
		toSerialize["useSsl"] = o.UseSsl.Get()
	}
	toSerialize["userDistinguishedName"] = o.UserDistinguishedName.Get()
	if o.UserPassword.IsSet() {
		toSerialize["userPassword"] = o.UserPassword.Get()
	}
	return toSerialize, nil
}

func (o *CreateLdapParamsSimpleAuthParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"userDistinguishedName",
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

	varCreateLdapParamsSimpleAuthParams := _CreateLdapParamsSimpleAuthParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateLdapParamsSimpleAuthParams)

	if err != nil {
		return err
	}

	*o = CreateLdapParamsSimpleAuthParams(varCreateLdapParamsSimpleAuthParams)

	return err
}

type NullableCreateLdapParamsSimpleAuthParams struct {
	value *CreateLdapParamsSimpleAuthParams
	isSet bool
}

func (v NullableCreateLdapParamsSimpleAuthParams) Get() *CreateLdapParamsSimpleAuthParams {
	return v.value
}

func (v *NullableCreateLdapParamsSimpleAuthParams) Set(val *CreateLdapParamsSimpleAuthParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLdapParamsSimpleAuthParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLdapParamsSimpleAuthParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLdapParamsSimpleAuthParams(val *CreateLdapParamsSimpleAuthParams) *NullableCreateLdapParamsSimpleAuthParams {
	return &NullableCreateLdapParamsSimpleAuthParams{value: val, isSet: true}
}

func (v NullableCreateLdapParamsSimpleAuthParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateLdapParamsSimpleAuthParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


