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

// checks if the CommonSourceRegistrationParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonSourceRegistrationParams{}

// CommonSourceRegistrationParams Specifies common parameters to register a source.
type CommonSourceRegistrationParams struct {
	// Specifies the password to access target entity.
	Password string `json:"password"`
	// Specifies the username to access target entity.
	Username string `json:"username"`
	// Specifies the description of the source being registered.
	Description NullableString `json:"description,omitempty"`
	// Specifies the endpoint IPaddress, URL or hostname of the host.
	Endpoint string `json:"endpoint"`
}

type _CommonSourceRegistrationParams CommonSourceRegistrationParams

// NewCommonSourceRegistrationParams instantiates a new CommonSourceRegistrationParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonSourceRegistrationParams(password string, username string, endpoint string) *CommonSourceRegistrationParams {
	this := CommonSourceRegistrationParams{}
	this.Password = password
	this.Username = username
	this.Endpoint = endpoint
	return &this
}

// NewCommonSourceRegistrationParamsWithDefaults instantiates a new CommonSourceRegistrationParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonSourceRegistrationParamsWithDefaults() *CommonSourceRegistrationParams {
	this := CommonSourceRegistrationParams{}
	return &this
}

// GetPassword returns the Password field value
func (o *CommonSourceRegistrationParams) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *CommonSourceRegistrationParams) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *CommonSourceRegistrationParams) SetPassword(v string) {
	o.Password = v
}

// GetUsername returns the Username field value
func (o *CommonSourceRegistrationParams) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *CommonSourceRegistrationParams) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *CommonSourceRegistrationParams) SetUsername(v string) {
	o.Username = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonSourceRegistrationParams) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonSourceRegistrationParams) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *CommonSourceRegistrationParams) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *CommonSourceRegistrationParams) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *CommonSourceRegistrationParams) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *CommonSourceRegistrationParams) UnsetDescription() {
	o.Description.Unset()
}

// GetEndpoint returns the Endpoint field value
func (o *CommonSourceRegistrationParams) GetEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value
// and a boolean to check if the value has been set.
func (o *CommonSourceRegistrationParams) GetEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Endpoint, true
}

// SetEndpoint sets field value
func (o *CommonSourceRegistrationParams) SetEndpoint(v string) {
	o.Endpoint = v
}

func (o CommonSourceRegistrationParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonSourceRegistrationParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["password"] = o.Password
	toSerialize["username"] = o.Username
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["endpoint"] = o.Endpoint
	return toSerialize, nil
}

func (o *CommonSourceRegistrationParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"password",
		"username",
		"endpoint",
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

	varCommonSourceRegistrationParams := _CommonSourceRegistrationParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCommonSourceRegistrationParams)

	if err != nil {
		return err
	}

	*o = CommonSourceRegistrationParams(varCommonSourceRegistrationParams)

	return err
}

type NullableCommonSourceRegistrationParams struct {
	value *CommonSourceRegistrationParams
	isSet bool
}

func (v NullableCommonSourceRegistrationParams) Get() *CommonSourceRegistrationParams {
	return v.value
}

func (v *NullableCommonSourceRegistrationParams) Set(val *CommonSourceRegistrationParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonSourceRegistrationParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonSourceRegistrationParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonSourceRegistrationParams(val *CommonSourceRegistrationParams) *NullableCommonSourceRegistrationParams {
	return &NullableCommonSourceRegistrationParams{value: val, isSet: true}
}

func (v NullableCommonSourceRegistrationParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonSourceRegistrationParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


