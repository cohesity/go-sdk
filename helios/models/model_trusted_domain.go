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

// TrustedDomain Specifies the details of a trusted domain.
type TrustedDomain struct {
	// Specifies a domain name.
	DomainName NullableString `json:"domainName,omitempty"`
	// Specifies a list of preferred domain controllers for this domain.
	PreferredDomainControllers *[]DomainController `json:"preferredDomainControllers,omitempty"`
}

// NewTrustedDomain instantiates a new TrustedDomain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrustedDomain() *TrustedDomain {
	this := TrustedDomain{}
	return &this
}

// NewTrustedDomainWithDefaults instantiates a new TrustedDomain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrustedDomainWithDefaults() *TrustedDomain {
	this := TrustedDomain{}
	return &this
}

// GetDomainName returns the DomainName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrustedDomain) GetDomainName() string {
	if o == nil || o.DomainName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DomainName.Get()
}

// GetDomainNameOk returns a tuple with the DomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrustedDomain) GetDomainNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DomainName.Get(), o.DomainName.IsSet()
}

// HasDomainName returns a boolean if a field has been set.
func (o *TrustedDomain) HasDomainName() bool {
	if o != nil && o.DomainName.IsSet() {
		return true
	}

	return false
}

// SetDomainName gets a reference to the given NullableString and assigns it to the DomainName field.
func (o *TrustedDomain) SetDomainName(v string) {
	o.DomainName.Set(&v)
}
// SetDomainNameNil sets the value for DomainName to be an explicit nil
func (o *TrustedDomain) SetDomainNameNil() {
	o.DomainName.Set(nil)
}

// UnsetDomainName ensures that no value is present for DomainName, not even an explicit nil
func (o *TrustedDomain) UnsetDomainName() {
	o.DomainName.Unset()
}

// GetPreferredDomainControllers returns the PreferredDomainControllers field value if set, zero value otherwise.
func (o *TrustedDomain) GetPreferredDomainControllers() []DomainController {
	if o == nil || o.PreferredDomainControllers == nil {
		var ret []DomainController
		return ret
	}
	return *o.PreferredDomainControllers
}

// GetPreferredDomainControllersOk returns a tuple with the PreferredDomainControllers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrustedDomain) GetPreferredDomainControllersOk() (*[]DomainController, bool) {
	if o == nil || o.PreferredDomainControllers == nil {
		return nil, false
	}
	return o.PreferredDomainControllers, true
}

// HasPreferredDomainControllers returns a boolean if a field has been set.
func (o *TrustedDomain) HasPreferredDomainControllers() bool {
	if o != nil && o.PreferredDomainControllers != nil {
		return true
	}

	return false
}

// SetPreferredDomainControllers gets a reference to the given []DomainController and assigns it to the PreferredDomainControllers field.
func (o *TrustedDomain) SetPreferredDomainControllers(v []DomainController) {
	o.PreferredDomainControllers = &v
}

func (o TrustedDomain) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DomainName.IsSet() {
		toSerialize["domainName"] = o.DomainName.Get()
	}
	if o.PreferredDomainControllers != nil {
		toSerialize["preferredDomainControllers"] = o.PreferredDomainControllers
	}
	return json.Marshal(toSerialize)
}

type NullableTrustedDomain struct {
	value *TrustedDomain
	isSet bool
}

func (v NullableTrustedDomain) Get() *TrustedDomain {
	return v.value
}

func (v *NullableTrustedDomain) Set(val *TrustedDomain) {
	v.value = val
	v.isSet = true
}

func (v NullableTrustedDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableTrustedDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrustedDomain(val *TrustedDomain) *NullableTrustedDomain {
	return &NullableTrustedDomain{value: val, isSet: true}
}

func (v NullableTrustedDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrustedDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


