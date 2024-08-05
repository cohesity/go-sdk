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

// checks if the AzureESConfigForIndexing type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureESConfigForIndexing{}

// AzureESConfigForIndexing Elasticsearch config for indexing in Azure CE.
type AzureESConfigForIndexing struct {
	// Client Id for the KeyVault.
	ClientId NullableString `json:"clientId"`
	// Fully qualified ES domain name.
	EsDomain NullableString `json:"esDomain"`
	// Name of the secret corresponding to tenant's ES creds.
	SecretName NullableString `json:"secretName"`
	// URL of the KeyVault where ES creds will be stored.
	VaultURL NullableString `json:"vaultURL"`
}

type _AzureESConfigForIndexing AzureESConfigForIndexing

// NewAzureESConfigForIndexing instantiates a new AzureESConfigForIndexing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureESConfigForIndexing(clientId NullableString, esDomain NullableString, secretName NullableString, vaultURL NullableString) *AzureESConfigForIndexing {
	this := AzureESConfigForIndexing{}
	this.ClientId = clientId
	this.EsDomain = esDomain
	this.SecretName = secretName
	this.VaultURL = vaultURL
	return &this
}

// NewAzureESConfigForIndexingWithDefaults instantiates a new AzureESConfigForIndexing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureESConfigForIndexingWithDefaults() *AzureESConfigForIndexing {
	this := AzureESConfigForIndexing{}
	return &this
}

// GetClientId returns the ClientId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AzureESConfigForIndexing) GetClientId() string {
	if o == nil || o.ClientId.Get() == nil {
		var ret string
		return ret
	}

	return *o.ClientId.Get()
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureESConfigForIndexing) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientId.Get(), o.ClientId.IsSet()
}

// SetClientId sets field value
func (o *AzureESConfigForIndexing) SetClientId(v string) {
	o.ClientId.Set(&v)
}

// GetEsDomain returns the EsDomain field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AzureESConfigForIndexing) GetEsDomain() string {
	if o == nil || o.EsDomain.Get() == nil {
		var ret string
		return ret
	}

	return *o.EsDomain.Get()
}

// GetEsDomainOk returns a tuple with the EsDomain field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureESConfigForIndexing) GetEsDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EsDomain.Get(), o.EsDomain.IsSet()
}

// SetEsDomain sets field value
func (o *AzureESConfigForIndexing) SetEsDomain(v string) {
	o.EsDomain.Set(&v)
}

// GetSecretName returns the SecretName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AzureESConfigForIndexing) GetSecretName() string {
	if o == nil || o.SecretName.Get() == nil {
		var ret string
		return ret
	}

	return *o.SecretName.Get()
}

// GetSecretNameOk returns a tuple with the SecretName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureESConfigForIndexing) GetSecretNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecretName.Get(), o.SecretName.IsSet()
}

// SetSecretName sets field value
func (o *AzureESConfigForIndexing) SetSecretName(v string) {
	o.SecretName.Set(&v)
}

// GetVaultURL returns the VaultURL field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AzureESConfigForIndexing) GetVaultURL() string {
	if o == nil || o.VaultURL.Get() == nil {
		var ret string
		return ret
	}

	return *o.VaultURL.Get()
}

// GetVaultURLOk returns a tuple with the VaultURL field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureESConfigForIndexing) GetVaultURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VaultURL.Get(), o.VaultURL.IsSet()
}

// SetVaultURL sets field value
func (o *AzureESConfigForIndexing) SetVaultURL(v string) {
	o.VaultURL.Set(&v)
}

func (o AzureESConfigForIndexing) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureESConfigForIndexing) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clientId"] = o.ClientId.Get()
	toSerialize["esDomain"] = o.EsDomain.Get()
	toSerialize["secretName"] = o.SecretName.Get()
	toSerialize["vaultURL"] = o.VaultURL.Get()
	return toSerialize, nil
}

func (o *AzureESConfigForIndexing) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"clientId",
		"esDomain",
		"secretName",
		"vaultURL",
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

	varAzureESConfigForIndexing := _AzureESConfigForIndexing{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAzureESConfigForIndexing)

	if err != nil {
		return err
	}

	*o = AzureESConfigForIndexing(varAzureESConfigForIndexing)

	return err
}

type NullableAzureESConfigForIndexing struct {
	value *AzureESConfigForIndexing
	isSet bool
}

func (v NullableAzureESConfigForIndexing) Get() *AzureESConfigForIndexing {
	return v.value
}

func (v *NullableAzureESConfigForIndexing) Set(val *AzureESConfigForIndexing) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureESConfigForIndexing) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureESConfigForIndexing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureESConfigForIndexing(val *AzureESConfigForIndexing) *NullableAzureESConfigForIndexing {
	return &NullableAzureESConfigForIndexing{value: val, isSet: true}
}

func (v NullableAzureESConfigForIndexing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureESConfigForIndexing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


