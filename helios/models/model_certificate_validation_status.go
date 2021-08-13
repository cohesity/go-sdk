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

// CertificateValidationStatus Describes the certificate validation status.
type CertificateValidationStatus struct {
	// Specifies the certificate validation status
	Type *string `json:"type,omitempty"`
}

// NewCertificateValidationStatus instantiates a new CertificateValidationStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateValidationStatus() *CertificateValidationStatus {
	this := CertificateValidationStatus{}
	return &this
}

// NewCertificateValidationStatusWithDefaults instantiates a new CertificateValidationStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateValidationStatusWithDefaults() *CertificateValidationStatus {
	this := CertificateValidationStatus{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CertificateValidationStatus) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateValidationStatus) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CertificateValidationStatus) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CertificateValidationStatus) SetType(v string) {
	o.Type = &v
}

func (o CertificateValidationStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableCertificateValidationStatus struct {
	value *CertificateValidationStatus
	isSet bool
}

func (v NullableCertificateValidationStatus) Get() *CertificateValidationStatus {
	return v.value
}

func (v *NullableCertificateValidationStatus) Set(val *CertificateValidationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateValidationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateValidationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateValidationStatus(val *CertificateValidationStatus) *NullableCertificateValidationStatus {
	return &NullableCertificateValidationStatus{value: val, isSet: true}
}

func (v NullableCertificateValidationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateValidationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


