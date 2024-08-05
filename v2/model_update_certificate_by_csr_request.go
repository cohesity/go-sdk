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

// checks if the UpdateCertificateByCsrRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateCertificateByCsrRequest{}

// UpdateCertificateByCsrRequest Specifies the request to update a certificate.
type UpdateCertificateByCsrRequest struct {
	// Specifies the certificate to be imported.
	Certificate NullableString `json:"certificate"`
	// Specifies the id of the csr corresponding to the certificate.
	CsrId NullableString `json:"csrId"`
}

type _UpdateCertificateByCsrRequest UpdateCertificateByCsrRequest

// NewUpdateCertificateByCsrRequest instantiates a new UpdateCertificateByCsrRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCertificateByCsrRequest(certificate NullableString, csrId NullableString) *UpdateCertificateByCsrRequest {
	this := UpdateCertificateByCsrRequest{}
	this.Certificate = certificate
	this.CsrId = csrId
	return &this
}

// NewUpdateCertificateByCsrRequestWithDefaults instantiates a new UpdateCertificateByCsrRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCertificateByCsrRequestWithDefaults() *UpdateCertificateByCsrRequest {
	this := UpdateCertificateByCsrRequest{}
	return &this
}

// GetCertificate returns the Certificate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UpdateCertificateByCsrRequest) GetCertificate() string {
	if o == nil || o.Certificate.Get() == nil {
		var ret string
		return ret
	}

	return *o.Certificate.Get()
}

// GetCertificateOk returns a tuple with the Certificate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateCertificateByCsrRequest) GetCertificateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Certificate.Get(), o.Certificate.IsSet()
}

// SetCertificate sets field value
func (o *UpdateCertificateByCsrRequest) SetCertificate(v string) {
	o.Certificate.Set(&v)
}

// GetCsrId returns the CsrId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UpdateCertificateByCsrRequest) GetCsrId() string {
	if o == nil || o.CsrId.Get() == nil {
		var ret string
		return ret
	}

	return *o.CsrId.Get()
}

// GetCsrIdOk returns a tuple with the CsrId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateCertificateByCsrRequest) GetCsrIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CsrId.Get(), o.CsrId.IsSet()
}

// SetCsrId sets field value
func (o *UpdateCertificateByCsrRequest) SetCsrId(v string) {
	o.CsrId.Set(&v)
}

func (o UpdateCertificateByCsrRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateCertificateByCsrRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["certificate"] = o.Certificate.Get()
	toSerialize["csrId"] = o.CsrId.Get()
	return toSerialize, nil
}

func (o *UpdateCertificateByCsrRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"certificate",
		"csrId",
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

	varUpdateCertificateByCsrRequest := _UpdateCertificateByCsrRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateCertificateByCsrRequest)

	if err != nil {
		return err
	}

	*o = UpdateCertificateByCsrRequest(varUpdateCertificateByCsrRequest)

	return err
}

type NullableUpdateCertificateByCsrRequest struct {
	value *UpdateCertificateByCsrRequest
	isSet bool
}

func (v NullableUpdateCertificateByCsrRequest) Get() *UpdateCertificateByCsrRequest {
	return v.value
}

func (v *NullableUpdateCertificateByCsrRequest) Set(val *UpdateCertificateByCsrRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCertificateByCsrRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCertificateByCsrRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCertificateByCsrRequest(val *UpdateCertificateByCsrRequest) *NullableUpdateCertificateByCsrRequest {
	return &NullableUpdateCertificateByCsrRequest{value: val, isSet: true}
}

func (v NullableUpdateCertificateByCsrRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCertificateByCsrRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


