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

// checks if the DomainController type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DomainController{}

// DomainController Specifies a domain controller.
type DomainController struct {
	// Specifies the domain controller name.
	Name NullableString `json:"name"`
	// Specifies the connection status.
	Status NullableString `json:"status,omitempty"`
}

type _DomainController DomainController

// NewDomainController instantiates a new DomainController object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainController(name NullableString) *DomainController {
	this := DomainController{}
	this.Name = name
	return &this
}

// NewDomainControllerWithDefaults instantiates a new DomainController object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainControllerWithDefaults() *DomainController {
	this := DomainController{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DomainController) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DomainController) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *DomainController) SetName(v string) {
	o.Name.Set(&v)
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DomainController) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DomainController) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *DomainController) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *DomainController) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *DomainController) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *DomainController) UnsetStatus() {
	o.Status.Unset()
}

func (o DomainController) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DomainController) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name.Get()
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	return toSerialize, nil
}

func (o *DomainController) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varDomainController := _DomainController{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDomainController)

	if err != nil {
		return err
	}

	*o = DomainController(varDomainController)

	return err
}

type NullableDomainController struct {
	value *DomainController
	isSet bool
}

func (v NullableDomainController) Get() *DomainController {
	return v.value
}

func (v *NullableDomainController) Set(val *DomainController) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainController) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainController) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainController(val *DomainController) *NullableDomainController {
	return &NullableDomainController{value: val, isSet: true}
}

func (v NullableDomainController) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainController) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


