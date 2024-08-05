/*
Cohesity REST API

Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
)

// checks if the GenerateSignCsrResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenerateSignCsrResponse{}

// GenerateSignCsrResponse Specifies the response to sign csr request
type GenerateSignCsrResponse struct {
	// Specifies ca cert in pem format
	CaCert []string `json:"caCert,omitempty"`
	// Specifies certificate in pem format.
	Certificate NullableString `json:"certificate,omitempty"`
}

// NewGenerateSignCsrResponse instantiates a new GenerateSignCsrResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateSignCsrResponse() *GenerateSignCsrResponse {
	this := GenerateSignCsrResponse{}
	return &this
}

// NewGenerateSignCsrResponseWithDefaults instantiates a new GenerateSignCsrResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateSignCsrResponseWithDefaults() *GenerateSignCsrResponse {
	this := GenerateSignCsrResponse{}
	return &this
}

// GetCaCert returns the CaCert field value if set, zero value otherwise.
func (o *GenerateSignCsrResponse) GetCaCert() []string {
	if o == nil || IsNil(o.CaCert) {
		var ret []string
		return ret
	}
	return o.CaCert
}

// GetCaCertOk returns a tuple with the CaCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateSignCsrResponse) GetCaCertOk() ([]string, bool) {
	if o == nil || IsNil(o.CaCert) {
		return nil, false
	}
	return o.CaCert, true
}

// HasCaCert returns a boolean if a field has been set.
func (o *GenerateSignCsrResponse) HasCaCert() bool {
	if o != nil && !IsNil(o.CaCert) {
		return true
	}

	return false
}

// SetCaCert gets a reference to the given []string and assigns it to the CaCert field.
func (o *GenerateSignCsrResponse) SetCaCert(v []string) {
	o.CaCert = v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GenerateSignCsrResponse) GetCertificate() string {
	if o == nil || IsNil(o.Certificate.Get()) {
		var ret string
		return ret
	}
	return *o.Certificate.Get()
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GenerateSignCsrResponse) GetCertificateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Certificate.Get(), o.Certificate.IsSet()
}

// HasCertificate returns a boolean if a field has been set.
func (o *GenerateSignCsrResponse) HasCertificate() bool {
	if o != nil && o.Certificate.IsSet() {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given NullableString and assigns it to the Certificate field.
func (o *GenerateSignCsrResponse) SetCertificate(v string) {
	o.Certificate.Set(&v)
}
// SetCertificateNil sets the value for Certificate to be an explicit nil
func (o *GenerateSignCsrResponse) SetCertificateNil() {
	o.Certificate.Set(nil)
}

// UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
func (o *GenerateSignCsrResponse) UnsetCertificate() {
	o.Certificate.Unset()
}

func (o GenerateSignCsrResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenerateSignCsrResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CaCert) {
		toSerialize["caCert"] = o.CaCert
	}
	if o.Certificate.IsSet() {
		toSerialize["certificate"] = o.Certificate.Get()
	}
	return toSerialize, nil
}

type NullableGenerateSignCsrResponse struct {
	value *GenerateSignCsrResponse
	isSet bool
}

func (v NullableGenerateSignCsrResponse) Get() *GenerateSignCsrResponse {
	return v.value
}

func (v *NullableGenerateSignCsrResponse) Set(val *GenerateSignCsrResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateSignCsrResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateSignCsrResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateSignCsrResponse(val *GenerateSignCsrResponse) *NullableGenerateSignCsrResponse {
	return &NullableGenerateSignCsrResponse{value: val, isSet: true}
}

func (v NullableGenerateSignCsrResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenerateSignCsrResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


