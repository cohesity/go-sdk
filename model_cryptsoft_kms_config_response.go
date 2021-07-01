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

// CryptsoftKmsConfigResponse struct for CryptsoftKmsConfigResponse
type CryptsoftKmsConfigResponse struct {
	// Specifies expiry date of client certificate.
	ClientCertificateExpiryDate NullableInt64 `json:"clientCertificateExpiryDate,omitempty"`
	// Specifies protocol version used to communicate with the KMS.
	KmipProtocolVersion NullableString `json:"kmipProtocolVersion,omitempty"`
	// Specifies the KMS IP address.
	ServerIp NullableString `json:"serverIp,omitempty"`
	// Specifies port on which the server is listening. Default port is 5696.
	ServerPort NullableInt32 `json:"serverPort,omitempty"`
}

// NewCryptsoftKmsConfigResponse instantiates a new CryptsoftKmsConfigResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCryptsoftKmsConfigResponse() *CryptsoftKmsConfigResponse {
	this := CryptsoftKmsConfigResponse{}
	return &this
}

// NewCryptsoftKmsConfigResponseWithDefaults instantiates a new CryptsoftKmsConfigResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCryptsoftKmsConfigResponseWithDefaults() *CryptsoftKmsConfigResponse {
	this := CryptsoftKmsConfigResponse{}
	return &this
}

// GetClientCertificateExpiryDate returns the ClientCertificateExpiryDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CryptsoftKmsConfigResponse) GetClientCertificateExpiryDate() int64 {
	if o == nil || o.ClientCertificateExpiryDate.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ClientCertificateExpiryDate.Get()
}

// GetClientCertificateExpiryDateOk returns a tuple with the ClientCertificateExpiryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CryptsoftKmsConfigResponse) GetClientCertificateExpiryDateOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ClientCertificateExpiryDate.Get(), o.ClientCertificateExpiryDate.IsSet()
}

// HasClientCertificateExpiryDate returns a boolean if a field has been set.
func (o *CryptsoftKmsConfigResponse) HasClientCertificateExpiryDate() bool {
	if o != nil && o.ClientCertificateExpiryDate.IsSet() {
		return true
	}

	return false
}

// SetClientCertificateExpiryDate gets a reference to the given NullableInt64 and assigns it to the ClientCertificateExpiryDate field.
func (o *CryptsoftKmsConfigResponse) SetClientCertificateExpiryDate(v int64) {
	o.ClientCertificateExpiryDate.Set(&v)
}
// SetClientCertificateExpiryDateNil sets the value for ClientCertificateExpiryDate to be an explicit nil
func (o *CryptsoftKmsConfigResponse) SetClientCertificateExpiryDateNil() {
	o.ClientCertificateExpiryDate.Set(nil)
}

// UnsetClientCertificateExpiryDate ensures that no value is present for ClientCertificateExpiryDate, not even an explicit nil
func (o *CryptsoftKmsConfigResponse) UnsetClientCertificateExpiryDate() {
	o.ClientCertificateExpiryDate.Unset()
}

// GetKmipProtocolVersion returns the KmipProtocolVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CryptsoftKmsConfigResponse) GetKmipProtocolVersion() string {
	if o == nil || o.KmipProtocolVersion.Get() == nil {
		var ret string
		return ret
	}
	return *o.KmipProtocolVersion.Get()
}

// GetKmipProtocolVersionOk returns a tuple with the KmipProtocolVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CryptsoftKmsConfigResponse) GetKmipProtocolVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.KmipProtocolVersion.Get(), o.KmipProtocolVersion.IsSet()
}

// HasKmipProtocolVersion returns a boolean if a field has been set.
func (o *CryptsoftKmsConfigResponse) HasKmipProtocolVersion() bool {
	if o != nil && o.KmipProtocolVersion.IsSet() {
		return true
	}

	return false
}

// SetKmipProtocolVersion gets a reference to the given NullableString and assigns it to the KmipProtocolVersion field.
func (o *CryptsoftKmsConfigResponse) SetKmipProtocolVersion(v string) {
	o.KmipProtocolVersion.Set(&v)
}
// SetKmipProtocolVersionNil sets the value for KmipProtocolVersion to be an explicit nil
func (o *CryptsoftKmsConfigResponse) SetKmipProtocolVersionNil() {
	o.KmipProtocolVersion.Set(nil)
}

// UnsetKmipProtocolVersion ensures that no value is present for KmipProtocolVersion, not even an explicit nil
func (o *CryptsoftKmsConfigResponse) UnsetKmipProtocolVersion() {
	o.KmipProtocolVersion.Unset()
}

// GetServerIp returns the ServerIp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CryptsoftKmsConfigResponse) GetServerIp() string {
	if o == nil || o.ServerIp.Get() == nil {
		var ret string
		return ret
	}
	return *o.ServerIp.Get()
}

// GetServerIpOk returns a tuple with the ServerIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CryptsoftKmsConfigResponse) GetServerIpOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ServerIp.Get(), o.ServerIp.IsSet()
}

// HasServerIp returns a boolean if a field has been set.
func (o *CryptsoftKmsConfigResponse) HasServerIp() bool {
	if o != nil && o.ServerIp.IsSet() {
		return true
	}

	return false
}

// SetServerIp gets a reference to the given NullableString and assigns it to the ServerIp field.
func (o *CryptsoftKmsConfigResponse) SetServerIp(v string) {
	o.ServerIp.Set(&v)
}
// SetServerIpNil sets the value for ServerIp to be an explicit nil
func (o *CryptsoftKmsConfigResponse) SetServerIpNil() {
	o.ServerIp.Set(nil)
}

// UnsetServerIp ensures that no value is present for ServerIp, not even an explicit nil
func (o *CryptsoftKmsConfigResponse) UnsetServerIp() {
	o.ServerIp.Unset()
}

// GetServerPort returns the ServerPort field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CryptsoftKmsConfigResponse) GetServerPort() int32 {
	if o == nil || o.ServerPort.Get() == nil {
		var ret int32
		return ret
	}
	return *o.ServerPort.Get()
}

// GetServerPortOk returns a tuple with the ServerPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CryptsoftKmsConfigResponse) GetServerPortOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ServerPort.Get(), o.ServerPort.IsSet()
}

// HasServerPort returns a boolean if a field has been set.
func (o *CryptsoftKmsConfigResponse) HasServerPort() bool {
	if o != nil && o.ServerPort.IsSet() {
		return true
	}

	return false
}

// SetServerPort gets a reference to the given NullableInt32 and assigns it to the ServerPort field.
func (o *CryptsoftKmsConfigResponse) SetServerPort(v int32) {
	o.ServerPort.Set(&v)
}
// SetServerPortNil sets the value for ServerPort to be an explicit nil
func (o *CryptsoftKmsConfigResponse) SetServerPortNil() {
	o.ServerPort.Set(nil)
}

// UnsetServerPort ensures that no value is present for ServerPort, not even an explicit nil
func (o *CryptsoftKmsConfigResponse) UnsetServerPort() {
	o.ServerPort.Unset()
}

func (o CryptsoftKmsConfigResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientCertificateExpiryDate.IsSet() {
		toSerialize["clientCertificateExpiryDate"] = o.ClientCertificateExpiryDate.Get()
	}
	if o.KmipProtocolVersion.IsSet() {
		toSerialize["kmipProtocolVersion"] = o.KmipProtocolVersion.Get()
	}
	if o.ServerIp.IsSet() {
		toSerialize["serverIp"] = o.ServerIp.Get()
	}
	if o.ServerPort.IsSet() {
		toSerialize["serverPort"] = o.ServerPort.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCryptsoftKmsConfigResponse struct {
	value *CryptsoftKmsConfigResponse
	isSet bool
}

func (v NullableCryptsoftKmsConfigResponse) Get() *CryptsoftKmsConfigResponse {
	return v.value
}

func (v *NullableCryptsoftKmsConfigResponse) Set(val *CryptsoftKmsConfigResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCryptsoftKmsConfigResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCryptsoftKmsConfigResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCryptsoftKmsConfigResponse(val *CryptsoftKmsConfigResponse) *NullableCryptsoftKmsConfigResponse {
	return &NullableCryptsoftKmsConfigResponse{value: val, isSet: true}
}

func (v NullableCryptsoftKmsConfigResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCryptsoftKmsConfigResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


