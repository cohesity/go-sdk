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

// checks if the AntivirusServiceGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AntivirusServiceGroup{}

// AntivirusServiceGroup Specifies an Antivirus Service group.
type AntivirusServiceGroup struct {
	// Specifies a list of Antivirus Services for this group.
	AntivirusServices []AntivirusService `json:"antivirusServices"`
	// Specifies the description for the Antivirus Service group.
	Description NullableString `json:"description,omitempty"`
	// This field is currently deprecated. Specifies whether the Antivirus Group is enabled.
	Enabled NullableBool `json:"enabled,omitempty"`
	// Specifies the Antivirus Service group name.
	Name NullableString `json:"name"`
	// Specifies the state[Enable, Disable] of the group.
	State NullableString `json:"state,omitempty"`
	// Specifies the Antivirus Service group id.
	Id NullableInt64 `json:"id,omitempty"`
}

type _AntivirusServiceGroup AntivirusServiceGroup

// NewAntivirusServiceGroup instantiates a new AntivirusServiceGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAntivirusServiceGroup(antivirusServices []AntivirusService, name NullableString) *AntivirusServiceGroup {
	this := AntivirusServiceGroup{}
	this.AntivirusServices = antivirusServices
	this.Name = name
	return &this
}

// NewAntivirusServiceGroupWithDefaults instantiates a new AntivirusServiceGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAntivirusServiceGroupWithDefaults() *AntivirusServiceGroup {
	this := AntivirusServiceGroup{}
	return &this
}

// GetAntivirusServices returns the AntivirusServices field value
// If the value is explicit nil, the zero value for []AntivirusService will be returned
func (o *AntivirusServiceGroup) GetAntivirusServices() []AntivirusService {
	if o == nil {
		var ret []AntivirusService
		return ret
	}

	return o.AntivirusServices
}

// GetAntivirusServicesOk returns a tuple with the AntivirusServices field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AntivirusServiceGroup) GetAntivirusServicesOk() ([]AntivirusService, bool) {
	if o == nil || IsNil(o.AntivirusServices) {
		return nil, false
	}
	return o.AntivirusServices, true
}

// SetAntivirusServices sets field value
func (o *AntivirusServiceGroup) SetAntivirusServices(v []AntivirusService) {
	o.AntivirusServices = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AntivirusServiceGroup) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AntivirusServiceGroup) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *AntivirusServiceGroup) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *AntivirusServiceGroup) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *AntivirusServiceGroup) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *AntivirusServiceGroup) UnsetDescription() {
	o.Description.Unset()
}

// GetEnabled returns the Enabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AntivirusServiceGroup) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled.Get()) {
		var ret bool
		return ret
	}
	return *o.Enabled.Get()
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AntivirusServiceGroup) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enabled.Get(), o.Enabled.IsSet()
}

// HasEnabled returns a boolean if a field has been set.
func (o *AntivirusServiceGroup) HasEnabled() bool {
	if o != nil && o.Enabled.IsSet() {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given NullableBool and assigns it to the Enabled field.
func (o *AntivirusServiceGroup) SetEnabled(v bool) {
	o.Enabled.Set(&v)
}
// SetEnabledNil sets the value for Enabled to be an explicit nil
func (o *AntivirusServiceGroup) SetEnabledNil() {
	o.Enabled.Set(nil)
}

// UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
func (o *AntivirusServiceGroup) UnsetEnabled() {
	o.Enabled.Unset()
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AntivirusServiceGroup) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AntivirusServiceGroup) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *AntivirusServiceGroup) SetName(v string) {
	o.Name.Set(&v)
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AntivirusServiceGroup) GetState() string {
	if o == nil || IsNil(o.State.Get()) {
		var ret string
		return ret
	}
	return *o.State.Get()
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AntivirusServiceGroup) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.State.Get(), o.State.IsSet()
}

// HasState returns a boolean if a field has been set.
func (o *AntivirusServiceGroup) HasState() bool {
	if o != nil && o.State.IsSet() {
		return true
	}

	return false
}

// SetState gets a reference to the given NullableString and assigns it to the State field.
func (o *AntivirusServiceGroup) SetState(v string) {
	o.State.Set(&v)
}
// SetStateNil sets the value for State to be an explicit nil
func (o *AntivirusServiceGroup) SetStateNil() {
	o.State.Set(nil)
}

// UnsetState ensures that no value is present for State, not even an explicit nil
func (o *AntivirusServiceGroup) UnsetState() {
	o.State.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AntivirusServiceGroup) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AntivirusServiceGroup) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *AntivirusServiceGroup) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *AntivirusServiceGroup) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *AntivirusServiceGroup) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *AntivirusServiceGroup) UnsetId() {
	o.Id.Unset()
}

func (o AntivirusServiceGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AntivirusServiceGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AntivirusServices != nil {
		toSerialize["antivirusServices"] = o.AntivirusServices
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Enabled.IsSet() {
		toSerialize["enabled"] = o.Enabled.Get()
	}
	toSerialize["name"] = o.Name.Get()
	if o.State.IsSet() {
		toSerialize["state"] = o.State.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	return toSerialize, nil
}

func (o *AntivirusServiceGroup) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"antivirusServices",
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

	varAntivirusServiceGroup := _AntivirusServiceGroup{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAntivirusServiceGroup)

	if err != nil {
		return err
	}

	*o = AntivirusServiceGroup(varAntivirusServiceGroup)

	return err
}

type NullableAntivirusServiceGroup struct {
	value *AntivirusServiceGroup
	isSet bool
}

func (v NullableAntivirusServiceGroup) Get() *AntivirusServiceGroup {
	return v.value
}

func (v *NullableAntivirusServiceGroup) Set(val *AntivirusServiceGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableAntivirusServiceGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableAntivirusServiceGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAntivirusServiceGroup(val *AntivirusServiceGroup) *NullableAntivirusServiceGroup {
	return &NullableAntivirusServiceGroup{value: val, isSet: true}
}

func (v NullableAntivirusServiceGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAntivirusServiceGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


