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

// checks if the Certificate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Certificate{}

// Certificate Certificate payload
type Certificate struct {
	// CA chain
	CaChainPem []string `json:"caChainPem,omitempty"`
	// Certificate (pem) to be imported
	CertPem []string `json:"certPem,omitempty"`
	// Private key
	PrivateKey *string `json:"privateKey,omitempty"`
	ProtoType *string `json:"protoType,omitempty"`
	ServiceTypes []string `json:"serviceTypes,omitempty"`
}

// NewCertificate instantiates a new Certificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificate() *Certificate {
	this := Certificate{}
	return &this
}

// NewCertificateWithDefaults instantiates a new Certificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateWithDefaults() *Certificate {
	this := Certificate{}
	return &this
}

// GetCaChainPem returns the CaChainPem field value if set, zero value otherwise.
func (o *Certificate) GetCaChainPem() []string {
	if o == nil || IsNil(o.CaChainPem) {
		var ret []string
		return ret
	}
	return o.CaChainPem
}

// GetCaChainPemOk returns a tuple with the CaChainPem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Certificate) GetCaChainPemOk() ([]string, bool) {
	if o == nil || IsNil(o.CaChainPem) {
		return nil, false
	}
	return o.CaChainPem, true
}

// HasCaChainPem returns a boolean if a field has been set.
func (o *Certificate) HasCaChainPem() bool {
	if o != nil && !IsNil(o.CaChainPem) {
		return true
	}

	return false
}

// SetCaChainPem gets a reference to the given []string and assigns it to the CaChainPem field.
func (o *Certificate) SetCaChainPem(v []string) {
	o.CaChainPem = v
}

// GetCertPem returns the CertPem field value if set, zero value otherwise.
func (o *Certificate) GetCertPem() []string {
	if o == nil || IsNil(o.CertPem) {
		var ret []string
		return ret
	}
	return o.CertPem
}

// GetCertPemOk returns a tuple with the CertPem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Certificate) GetCertPemOk() ([]string, bool) {
	if o == nil || IsNil(o.CertPem) {
		return nil, false
	}
	return o.CertPem, true
}

// HasCertPem returns a boolean if a field has been set.
func (o *Certificate) HasCertPem() bool {
	if o != nil && !IsNil(o.CertPem) {
		return true
	}

	return false
}

// SetCertPem gets a reference to the given []string and assigns it to the CertPem field.
func (o *Certificate) SetCertPem(v []string) {
	o.CertPem = v
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise.
func (o *Certificate) GetPrivateKey() string {
	if o == nil || IsNil(o.PrivateKey) {
		var ret string
		return ret
	}
	return *o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Certificate) GetPrivateKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateKey) {
		return nil, false
	}
	return o.PrivateKey, true
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *Certificate) HasPrivateKey() bool {
	if o != nil && !IsNil(o.PrivateKey) {
		return true
	}

	return false
}

// SetPrivateKey gets a reference to the given string and assigns it to the PrivateKey field.
func (o *Certificate) SetPrivateKey(v string) {
	o.PrivateKey = &v
}

// GetProtoType returns the ProtoType field value if set, zero value otherwise.
func (o *Certificate) GetProtoType() string {
	if o == nil || IsNil(o.ProtoType) {
		var ret string
		return ret
	}
	return *o.ProtoType
}

// GetProtoTypeOk returns a tuple with the ProtoType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Certificate) GetProtoTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ProtoType) {
		return nil, false
	}
	return o.ProtoType, true
}

// HasProtoType returns a boolean if a field has been set.
func (o *Certificate) HasProtoType() bool {
	if o != nil && !IsNil(o.ProtoType) {
		return true
	}

	return false
}

// SetProtoType gets a reference to the given string and assigns it to the ProtoType field.
func (o *Certificate) SetProtoType(v string) {
	o.ProtoType = &v
}

// GetServiceTypes returns the ServiceTypes field value if set, zero value otherwise.
func (o *Certificate) GetServiceTypes() []string {
	if o == nil || IsNil(o.ServiceTypes) {
		var ret []string
		return ret
	}
	return o.ServiceTypes
}

// GetServiceTypesOk returns a tuple with the ServiceTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Certificate) GetServiceTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.ServiceTypes) {
		return nil, false
	}
	return o.ServiceTypes, true
}

// HasServiceTypes returns a boolean if a field has been set.
func (o *Certificate) HasServiceTypes() bool {
	if o != nil && !IsNil(o.ServiceTypes) {
		return true
	}

	return false
}

// SetServiceTypes gets a reference to the given []string and assigns it to the ServiceTypes field.
func (o *Certificate) SetServiceTypes(v []string) {
	o.ServiceTypes = v
}

func (o Certificate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Certificate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CaChainPem) {
		toSerialize["caChainPem"] = o.CaChainPem
	}
	if !IsNil(o.CertPem) {
		toSerialize["certPem"] = o.CertPem
	}
	if !IsNil(o.PrivateKey) {
		toSerialize["privateKey"] = o.PrivateKey
	}
	if !IsNil(o.ProtoType) {
		toSerialize["protoType"] = o.ProtoType
	}
	if !IsNil(o.ServiceTypes) {
		toSerialize["serviceTypes"] = o.ServiceTypes
	}
	return toSerialize, nil
}

type NullableCertificate struct {
	value *Certificate
	isSet bool
}

func (v NullableCertificate) Get() *Certificate {
	return v.value
}

func (v *NullableCertificate) Set(val *Certificate) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificate) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificate(val *Certificate) *NullableCertificate {
	return &NullableCertificate{value: val, isSet: true}
}

func (v NullableCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


