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

// AntivirusServiceGroupParams struct for AntivirusServiceGroupParams
type AntivirusServiceGroupParams struct {
	// Specifies the Antivirus services for this provider.
	AntivirusServices []AntivirusServiceConfigParams `json:"antivirusServices,omitempty"`
	// Specifies the description of the Antivirus service group.
	Description NullableString `json:"description,omitempty"`
	// Specifies the name of the Antivirus service group.
	Name NullableString `json:"name"`
}

// NewAntivirusServiceGroupParams instantiates a new AntivirusServiceGroupParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAntivirusServiceGroupParams(name NullableString) *AntivirusServiceGroupParams {
	this := AntivirusServiceGroupParams{}
	this.Name = name
	return &this
}

// NewAntivirusServiceGroupParamsWithDefaults instantiates a new AntivirusServiceGroupParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAntivirusServiceGroupParamsWithDefaults() *AntivirusServiceGroupParams {
	this := AntivirusServiceGroupParams{}
	return &this
}

// GetAntivirusServices returns the AntivirusServices field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AntivirusServiceGroupParams) GetAntivirusServices() []AntivirusServiceConfigParams {
	if o == nil  {
		var ret []AntivirusServiceConfigParams
		return ret
	}
	return o.AntivirusServices
}

// GetAntivirusServicesOk returns a tuple with the AntivirusServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AntivirusServiceGroupParams) GetAntivirusServicesOk() (*[]AntivirusServiceConfigParams, bool) {
	if o == nil || o.AntivirusServices == nil {
		return nil, false
	}
	return &o.AntivirusServices, true
}

// HasAntivirusServices returns a boolean if a field has been set.
func (o *AntivirusServiceGroupParams) HasAntivirusServices() bool {
	if o != nil && o.AntivirusServices != nil {
		return true
	}

	return false
}

// SetAntivirusServices gets a reference to the given []AntivirusServiceConfigParams and assigns it to the AntivirusServices field.
func (o *AntivirusServiceGroupParams) SetAntivirusServices(v []AntivirusServiceConfigParams) {
	o.AntivirusServices = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AntivirusServiceGroupParams) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AntivirusServiceGroupParams) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *AntivirusServiceGroupParams) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *AntivirusServiceGroupParams) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *AntivirusServiceGroupParams) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *AntivirusServiceGroupParams) UnsetDescription() {
	o.Description.Unset()
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AntivirusServiceGroupParams) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AntivirusServiceGroupParams) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *AntivirusServiceGroupParams) SetName(v string) {
	o.Name.Set(&v)
}

func (o AntivirusServiceGroupParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AntivirusServices != nil {
		toSerialize["antivirusServices"] = o.AntivirusServices
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if true {
		toSerialize["name"] = o.Name.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAntivirusServiceGroupParams struct {
	value *AntivirusServiceGroupParams
	isSet bool
}

func (v NullableAntivirusServiceGroupParams) Get() *AntivirusServiceGroupParams {
	return v.value
}

func (v *NullableAntivirusServiceGroupParams) Set(val *AntivirusServiceGroupParams) {
	v.value = val
	v.isSet = true
}

func (v NullableAntivirusServiceGroupParams) IsSet() bool {
	return v.isSet
}

func (v *NullableAntivirusServiceGroupParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAntivirusServiceGroupParams(val *AntivirusServiceGroupParams) *NullableAntivirusServiceGroupParams {
	return &NullableAntivirusServiceGroupParams{value: val, isSet: true}
}

func (v NullableAntivirusServiceGroupParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAntivirusServiceGroupParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


