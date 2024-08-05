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

// checks if the EndpointOnlyRegistrationParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EndpointOnlyRegistrationParams{}

// EndpointOnlyRegistrationParams Specifies an endpoint to register a source.
type EndpointOnlyRegistrationParams struct {
	// Specifies the description of the source being registered.
	Description NullableString `json:"description,omitempty"`
	// Specifies the endpoint IPaddress, URL or hostname of the host.
	Endpoint string `json:"endpoint"`
}

type _EndpointOnlyRegistrationParams EndpointOnlyRegistrationParams

// NewEndpointOnlyRegistrationParams instantiates a new EndpointOnlyRegistrationParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointOnlyRegistrationParams(endpoint string) *EndpointOnlyRegistrationParams {
	this := EndpointOnlyRegistrationParams{}
	this.Endpoint = endpoint
	return &this
}

// NewEndpointOnlyRegistrationParamsWithDefaults instantiates a new EndpointOnlyRegistrationParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointOnlyRegistrationParamsWithDefaults() *EndpointOnlyRegistrationParams {
	this := EndpointOnlyRegistrationParams{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EndpointOnlyRegistrationParams) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EndpointOnlyRegistrationParams) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *EndpointOnlyRegistrationParams) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *EndpointOnlyRegistrationParams) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *EndpointOnlyRegistrationParams) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *EndpointOnlyRegistrationParams) UnsetDescription() {
	o.Description.Unset()
}

// GetEndpoint returns the Endpoint field value
func (o *EndpointOnlyRegistrationParams) GetEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value
// and a boolean to check if the value has been set.
func (o *EndpointOnlyRegistrationParams) GetEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Endpoint, true
}

// SetEndpoint sets field value
func (o *EndpointOnlyRegistrationParams) SetEndpoint(v string) {
	o.Endpoint = v
}

func (o EndpointOnlyRegistrationParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointOnlyRegistrationParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["endpoint"] = o.Endpoint
	return toSerialize, nil
}

func (o *EndpointOnlyRegistrationParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varEndpointOnlyRegistrationParams := _EndpointOnlyRegistrationParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEndpointOnlyRegistrationParams)

	if err != nil {
		return err
	}

	*o = EndpointOnlyRegistrationParams(varEndpointOnlyRegistrationParams)

	return err
}

type NullableEndpointOnlyRegistrationParams struct {
	value *EndpointOnlyRegistrationParams
	isSet bool
}

func (v NullableEndpointOnlyRegistrationParams) Get() *EndpointOnlyRegistrationParams {
	return v.value
}

func (v *NullableEndpointOnlyRegistrationParams) Set(val *EndpointOnlyRegistrationParams) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointOnlyRegistrationParams) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointOnlyRegistrationParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointOnlyRegistrationParams(val *EndpointOnlyRegistrationParams) *NullableEndpointOnlyRegistrationParams {
	return &NullableEndpointOnlyRegistrationParams{value: val, isSet: true}
}

func (v NullableEndpointOnlyRegistrationParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointOnlyRegistrationParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


