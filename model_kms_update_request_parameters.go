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

// KmsUpdateRequestParameters struct for KmsUpdateRequestParameters
type KmsUpdateRequestParameters struct {
	AwsKms *AwsKmsUpdateParams `json:"awsKms,omitempty"`
	CryptsoftKms *CryptsoftKmsUpdateParams `json:"cryptsoftKms,omitempty"`
	// The Id of a KMS server.
	Id NullableInt64 `json:"id,omitempty"`
	// Specifies the name given to the KMS Server.
	ServerName NullableString `json:"serverName,omitempty"`
}

// NewKmsUpdateRequestParameters instantiates a new KmsUpdateRequestParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKmsUpdateRequestParameters() *KmsUpdateRequestParameters {
	this := KmsUpdateRequestParameters{}
	return &this
}

// NewKmsUpdateRequestParametersWithDefaults instantiates a new KmsUpdateRequestParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKmsUpdateRequestParametersWithDefaults() *KmsUpdateRequestParameters {
	this := KmsUpdateRequestParameters{}
	return &this
}

// GetAwsKms returns the AwsKms field value if set, zero value otherwise.
func (o *KmsUpdateRequestParameters) GetAwsKms() AwsKmsUpdateParams {
	if o == nil || o.AwsKms == nil {
		var ret AwsKmsUpdateParams
		return ret
	}
	return *o.AwsKms
}

// GetAwsKmsOk returns a tuple with the AwsKms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KmsUpdateRequestParameters) GetAwsKmsOk() (*AwsKmsUpdateParams, bool) {
	if o == nil || o.AwsKms == nil {
		return nil, false
	}
	return o.AwsKms, true
}

// HasAwsKms returns a boolean if a field has been set.
func (o *KmsUpdateRequestParameters) HasAwsKms() bool {
	if o != nil && o.AwsKms != nil {
		return true
	}

	return false
}

// SetAwsKms gets a reference to the given AwsKmsUpdateParams and assigns it to the AwsKms field.
func (o *KmsUpdateRequestParameters) SetAwsKms(v AwsKmsUpdateParams) {
	o.AwsKms = &v
}

// GetCryptsoftKms returns the CryptsoftKms field value if set, zero value otherwise.
func (o *KmsUpdateRequestParameters) GetCryptsoftKms() CryptsoftKmsUpdateParams {
	if o == nil || o.CryptsoftKms == nil {
		var ret CryptsoftKmsUpdateParams
		return ret
	}
	return *o.CryptsoftKms
}

// GetCryptsoftKmsOk returns a tuple with the CryptsoftKms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KmsUpdateRequestParameters) GetCryptsoftKmsOk() (*CryptsoftKmsUpdateParams, bool) {
	if o == nil || o.CryptsoftKms == nil {
		return nil, false
	}
	return o.CryptsoftKms, true
}

// HasCryptsoftKms returns a boolean if a field has been set.
func (o *KmsUpdateRequestParameters) HasCryptsoftKms() bool {
	if o != nil && o.CryptsoftKms != nil {
		return true
	}

	return false
}

// SetCryptsoftKms gets a reference to the given CryptsoftKmsUpdateParams and assigns it to the CryptsoftKms field.
func (o *KmsUpdateRequestParameters) SetCryptsoftKms(v CryptsoftKmsUpdateParams) {
	o.CryptsoftKms = &v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KmsUpdateRequestParameters) GetId() int64 {
	if o == nil || o.Id.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KmsUpdateRequestParameters) GetIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *KmsUpdateRequestParameters) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *KmsUpdateRequestParameters) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *KmsUpdateRequestParameters) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *KmsUpdateRequestParameters) UnsetId() {
	o.Id.Unset()
}

// GetServerName returns the ServerName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KmsUpdateRequestParameters) GetServerName() string {
	if o == nil || o.ServerName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ServerName.Get()
}

// GetServerNameOk returns a tuple with the ServerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KmsUpdateRequestParameters) GetServerNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ServerName.Get(), o.ServerName.IsSet()
}

// HasServerName returns a boolean if a field has been set.
func (o *KmsUpdateRequestParameters) HasServerName() bool {
	if o != nil && o.ServerName.IsSet() {
		return true
	}

	return false
}

// SetServerName gets a reference to the given NullableString and assigns it to the ServerName field.
func (o *KmsUpdateRequestParameters) SetServerName(v string) {
	o.ServerName.Set(&v)
}
// SetServerNameNil sets the value for ServerName to be an explicit nil
func (o *KmsUpdateRequestParameters) SetServerNameNil() {
	o.ServerName.Set(nil)
}

// UnsetServerName ensures that no value is present for ServerName, not even an explicit nil
func (o *KmsUpdateRequestParameters) UnsetServerName() {
	o.ServerName.Unset()
}

func (o KmsUpdateRequestParameters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AwsKms != nil {
		toSerialize["awsKms"] = o.AwsKms
	}
	if o.CryptsoftKms != nil {
		toSerialize["cryptsoftKms"] = o.CryptsoftKms
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.ServerName.IsSet() {
		toSerialize["serverName"] = o.ServerName.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableKmsUpdateRequestParameters struct {
	value *KmsUpdateRequestParameters
	isSet bool
}

func (v NullableKmsUpdateRequestParameters) Get() *KmsUpdateRequestParameters {
	return v.value
}

func (v *NullableKmsUpdateRequestParameters) Set(val *KmsUpdateRequestParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableKmsUpdateRequestParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableKmsUpdateRequestParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKmsUpdateRequestParameters(val *KmsUpdateRequestParameters) *NullableKmsUpdateRequestParameters {
	return &NullableKmsUpdateRequestParameters{value: val, isSet: true}
}

func (v NullableKmsUpdateRequestParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKmsUpdateRequestParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


