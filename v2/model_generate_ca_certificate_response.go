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

// checks if the GenerateCaCertificateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenerateCaCertificateResponse{}

// GenerateCaCertificateResponse Specifies the response to generate ca certificate request
type GenerateCaCertificateResponse struct {
	// Certificate in pem format.
	Certificate NullableString `json:"certificate,omitempty"`
	// Private key of the certificate.
	PrivateKey NullableString `json:"privateKey,omitempty"`
}

// NewGenerateCaCertificateResponse instantiates a new GenerateCaCertificateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateCaCertificateResponse() *GenerateCaCertificateResponse {
	this := GenerateCaCertificateResponse{}
	return &this
}

// NewGenerateCaCertificateResponseWithDefaults instantiates a new GenerateCaCertificateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateCaCertificateResponseWithDefaults() *GenerateCaCertificateResponse {
	this := GenerateCaCertificateResponse{}
	return &this
}

// GetCertificate returns the Certificate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GenerateCaCertificateResponse) GetCertificate() string {
	if o == nil || IsNil(o.Certificate.Get()) {
		var ret string
		return ret
	}
	return *o.Certificate.Get()
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GenerateCaCertificateResponse) GetCertificateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Certificate.Get(), o.Certificate.IsSet()
}

// HasCertificate returns a boolean if a field has been set.
func (o *GenerateCaCertificateResponse) HasCertificate() bool {
	if o != nil && o.Certificate.IsSet() {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given NullableString and assigns it to the Certificate field.
func (o *GenerateCaCertificateResponse) SetCertificate(v string) {
	o.Certificate.Set(&v)
}
// SetCertificateNil sets the value for Certificate to be an explicit nil
func (o *GenerateCaCertificateResponse) SetCertificateNil() {
	o.Certificate.Set(nil)
}

// UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
func (o *GenerateCaCertificateResponse) UnsetCertificate() {
	o.Certificate.Unset()
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GenerateCaCertificateResponse) GetPrivateKey() string {
	if o == nil || IsNil(o.PrivateKey.Get()) {
		var ret string
		return ret
	}
	return *o.PrivateKey.Get()
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GenerateCaCertificateResponse) GetPrivateKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrivateKey.Get(), o.PrivateKey.IsSet()
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *GenerateCaCertificateResponse) HasPrivateKey() bool {
	if o != nil && o.PrivateKey.IsSet() {
		return true
	}

	return false
}

// SetPrivateKey gets a reference to the given NullableString and assigns it to the PrivateKey field.
func (o *GenerateCaCertificateResponse) SetPrivateKey(v string) {
	o.PrivateKey.Set(&v)
}
// SetPrivateKeyNil sets the value for PrivateKey to be an explicit nil
func (o *GenerateCaCertificateResponse) SetPrivateKeyNil() {
	o.PrivateKey.Set(nil)
}

// UnsetPrivateKey ensures that no value is present for PrivateKey, not even an explicit nil
func (o *GenerateCaCertificateResponse) UnsetPrivateKey() {
	o.PrivateKey.Unset()
}

func (o GenerateCaCertificateResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenerateCaCertificateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Certificate.IsSet() {
		toSerialize["certificate"] = o.Certificate.Get()
	}
	if o.PrivateKey.IsSet() {
		toSerialize["privateKey"] = o.PrivateKey.Get()
	}
	return toSerialize, nil
}

type NullableGenerateCaCertificateResponse struct {
	value *GenerateCaCertificateResponse
	isSet bool
}

func (v NullableGenerateCaCertificateResponse) Get() *GenerateCaCertificateResponse {
	return v.value
}

func (v *NullableGenerateCaCertificateResponse) Set(val *GenerateCaCertificateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateCaCertificateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateCaCertificateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateCaCertificateResponse(val *GenerateCaCertificateResponse) *NullableGenerateCaCertificateResponse {
	return &NullableGenerateCaCertificateResponse{value: val, isSet: true}
}

func (v NullableGenerateCaCertificateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenerateCaCertificateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


