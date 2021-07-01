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

// SslCertificateConfig SslCertificateConfig represents the SSL certificate object exposed to the user.
type SslCertificateConfig struct {
	// Certificate is a SSL certificate used by Iris HTTPS webserver.
	Certificate NullableString `json:"certificate,omitempty"`
	// LastUpdateTimeMsecs is a time in milliseconds at which certificate was last updated.
	LastUpdateTimeMsecs NullableInt64 `json:"lastUpdateTimeMsecs,omitempty"`
	// PrivateKey is a matching private key of the above certificate.
	PrivateKey NullableString `json:"privateKey,omitempty"`
}

// NewSslCertificateConfig instantiates a new SslCertificateConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSslCertificateConfig() *SslCertificateConfig {
	this := SslCertificateConfig{}
	return &this
}

// NewSslCertificateConfigWithDefaults instantiates a new SslCertificateConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSslCertificateConfigWithDefaults() *SslCertificateConfig {
	this := SslCertificateConfig{}
	return &this
}

// GetCertificate returns the Certificate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SslCertificateConfig) GetCertificate() string {
	if o == nil || o.Certificate.Get() == nil {
		var ret string
		return ret
	}
	return *o.Certificate.Get()
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SslCertificateConfig) GetCertificateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Certificate.Get(), o.Certificate.IsSet()
}

// HasCertificate returns a boolean if a field has been set.
func (o *SslCertificateConfig) HasCertificate() bool {
	if o != nil && o.Certificate.IsSet() {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given NullableString and assigns it to the Certificate field.
func (o *SslCertificateConfig) SetCertificate(v string) {
	o.Certificate.Set(&v)
}
// SetCertificateNil sets the value for Certificate to be an explicit nil
func (o *SslCertificateConfig) SetCertificateNil() {
	o.Certificate.Set(nil)
}

// UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
func (o *SslCertificateConfig) UnsetCertificate() {
	o.Certificate.Unset()
}

// GetLastUpdateTimeMsecs returns the LastUpdateTimeMsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SslCertificateConfig) GetLastUpdateTimeMsecs() int64 {
	if o == nil || o.LastUpdateTimeMsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.LastUpdateTimeMsecs.Get()
}

// GetLastUpdateTimeMsecsOk returns a tuple with the LastUpdateTimeMsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SslCertificateConfig) GetLastUpdateTimeMsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastUpdateTimeMsecs.Get(), o.LastUpdateTimeMsecs.IsSet()
}

// HasLastUpdateTimeMsecs returns a boolean if a field has been set.
func (o *SslCertificateConfig) HasLastUpdateTimeMsecs() bool {
	if o != nil && o.LastUpdateTimeMsecs.IsSet() {
		return true
	}

	return false
}

// SetLastUpdateTimeMsecs gets a reference to the given NullableInt64 and assigns it to the LastUpdateTimeMsecs field.
func (o *SslCertificateConfig) SetLastUpdateTimeMsecs(v int64) {
	o.LastUpdateTimeMsecs.Set(&v)
}
// SetLastUpdateTimeMsecsNil sets the value for LastUpdateTimeMsecs to be an explicit nil
func (o *SslCertificateConfig) SetLastUpdateTimeMsecsNil() {
	o.LastUpdateTimeMsecs.Set(nil)
}

// UnsetLastUpdateTimeMsecs ensures that no value is present for LastUpdateTimeMsecs, not even an explicit nil
func (o *SslCertificateConfig) UnsetLastUpdateTimeMsecs() {
	o.LastUpdateTimeMsecs.Unset()
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SslCertificateConfig) GetPrivateKey() string {
	if o == nil || o.PrivateKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.PrivateKey.Get()
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SslCertificateConfig) GetPrivateKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PrivateKey.Get(), o.PrivateKey.IsSet()
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *SslCertificateConfig) HasPrivateKey() bool {
	if o != nil && o.PrivateKey.IsSet() {
		return true
	}

	return false
}

// SetPrivateKey gets a reference to the given NullableString and assigns it to the PrivateKey field.
func (o *SslCertificateConfig) SetPrivateKey(v string) {
	o.PrivateKey.Set(&v)
}
// SetPrivateKeyNil sets the value for PrivateKey to be an explicit nil
func (o *SslCertificateConfig) SetPrivateKeyNil() {
	o.PrivateKey.Set(nil)
}

// UnsetPrivateKey ensures that no value is present for PrivateKey, not even an explicit nil
func (o *SslCertificateConfig) UnsetPrivateKey() {
	o.PrivateKey.Unset()
}

func (o SslCertificateConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Certificate.IsSet() {
		toSerialize["certificate"] = o.Certificate.Get()
	}
	if o.LastUpdateTimeMsecs.IsSet() {
		toSerialize["lastUpdateTimeMsecs"] = o.LastUpdateTimeMsecs.Get()
	}
	if o.PrivateKey.IsSet() {
		toSerialize["privateKey"] = o.PrivateKey.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSslCertificateConfig struct {
	value *SslCertificateConfig
	isSet bool
}

func (v NullableSslCertificateConfig) Get() *SslCertificateConfig {
	return v.value
}

func (v *NullableSslCertificateConfig) Set(val *SslCertificateConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableSslCertificateConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableSslCertificateConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSslCertificateConfig(val *SslCertificateConfig) *NullableSslCertificateConfig {
	return &NullableSslCertificateConfig{value: val, isSet: true}
}

func (v NullableSslCertificateConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSslCertificateConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


