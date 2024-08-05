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

// checks if the AzureApplicationCredentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureApplicationCredentials{}

// AzureApplicationCredentials Specifies the credentials of an application from Azure active directory.
type AzureApplicationCredentials struct {
	// Specifies the application id of an Azure active directory application.
	ApplicationId NullableString `json:"applicationId"`
	// Specifies the application object id of an Azure active directory application.
	ApplicationObjectId NullableString `json:"applicationObjectId,omitempty"`
	// Specifies the encrypted application key of an Azure active directory application.
	EncryptedApplicationKey NullableString `json:"encryptedApplicationKey,omitempty"`
}

type _AzureApplicationCredentials AzureApplicationCredentials

// NewAzureApplicationCredentials instantiates a new AzureApplicationCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureApplicationCredentials(applicationId NullableString) *AzureApplicationCredentials {
	this := AzureApplicationCredentials{}
	this.ApplicationId = applicationId
	return &this
}

// NewAzureApplicationCredentialsWithDefaults instantiates a new AzureApplicationCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureApplicationCredentialsWithDefaults() *AzureApplicationCredentials {
	this := AzureApplicationCredentials{}
	return &this
}

// GetApplicationId returns the ApplicationId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AzureApplicationCredentials) GetApplicationId() string {
	if o == nil || o.ApplicationId.Get() == nil {
		var ret string
		return ret
	}

	return *o.ApplicationId.Get()
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureApplicationCredentials) GetApplicationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApplicationId.Get(), o.ApplicationId.IsSet()
}

// SetApplicationId sets field value
func (o *AzureApplicationCredentials) SetApplicationId(v string) {
	o.ApplicationId.Set(&v)
}

// GetApplicationObjectId returns the ApplicationObjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureApplicationCredentials) GetApplicationObjectId() string {
	if o == nil || IsNil(o.ApplicationObjectId.Get()) {
		var ret string
		return ret
	}
	return *o.ApplicationObjectId.Get()
}

// GetApplicationObjectIdOk returns a tuple with the ApplicationObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureApplicationCredentials) GetApplicationObjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApplicationObjectId.Get(), o.ApplicationObjectId.IsSet()
}

// HasApplicationObjectId returns a boolean if a field has been set.
func (o *AzureApplicationCredentials) HasApplicationObjectId() bool {
	if o != nil && o.ApplicationObjectId.IsSet() {
		return true
	}

	return false
}

// SetApplicationObjectId gets a reference to the given NullableString and assigns it to the ApplicationObjectId field.
func (o *AzureApplicationCredentials) SetApplicationObjectId(v string) {
	o.ApplicationObjectId.Set(&v)
}
// SetApplicationObjectIdNil sets the value for ApplicationObjectId to be an explicit nil
func (o *AzureApplicationCredentials) SetApplicationObjectIdNil() {
	o.ApplicationObjectId.Set(nil)
}

// UnsetApplicationObjectId ensures that no value is present for ApplicationObjectId, not even an explicit nil
func (o *AzureApplicationCredentials) UnsetApplicationObjectId() {
	o.ApplicationObjectId.Unset()
}

// GetEncryptedApplicationKey returns the EncryptedApplicationKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureApplicationCredentials) GetEncryptedApplicationKey() string {
	if o == nil || IsNil(o.EncryptedApplicationKey.Get()) {
		var ret string
		return ret
	}
	return *o.EncryptedApplicationKey.Get()
}

// GetEncryptedApplicationKeyOk returns a tuple with the EncryptedApplicationKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureApplicationCredentials) GetEncryptedApplicationKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EncryptedApplicationKey.Get(), o.EncryptedApplicationKey.IsSet()
}

// HasEncryptedApplicationKey returns a boolean if a field has been set.
func (o *AzureApplicationCredentials) HasEncryptedApplicationKey() bool {
	if o != nil && o.EncryptedApplicationKey.IsSet() {
		return true
	}

	return false
}

// SetEncryptedApplicationKey gets a reference to the given NullableString and assigns it to the EncryptedApplicationKey field.
func (o *AzureApplicationCredentials) SetEncryptedApplicationKey(v string) {
	o.EncryptedApplicationKey.Set(&v)
}
// SetEncryptedApplicationKeyNil sets the value for EncryptedApplicationKey to be an explicit nil
func (o *AzureApplicationCredentials) SetEncryptedApplicationKeyNil() {
	o.EncryptedApplicationKey.Set(nil)
}

// UnsetEncryptedApplicationKey ensures that no value is present for EncryptedApplicationKey, not even an explicit nil
func (o *AzureApplicationCredentials) UnsetEncryptedApplicationKey() {
	o.EncryptedApplicationKey.Unset()
}

func (o AzureApplicationCredentials) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureApplicationCredentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["applicationId"] = o.ApplicationId.Get()
	if o.ApplicationObjectId.IsSet() {
		toSerialize["applicationObjectId"] = o.ApplicationObjectId.Get()
	}
	if o.EncryptedApplicationKey.IsSet() {
		toSerialize["encryptedApplicationKey"] = o.EncryptedApplicationKey.Get()
	}
	return toSerialize, nil
}

func (o *AzureApplicationCredentials) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"applicationId",
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

	varAzureApplicationCredentials := _AzureApplicationCredentials{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAzureApplicationCredentials)

	if err != nil {
		return err
	}

	*o = AzureApplicationCredentials(varAzureApplicationCredentials)

	return err
}

type NullableAzureApplicationCredentials struct {
	value *AzureApplicationCredentials
	isSet bool
}

func (v NullableAzureApplicationCredentials) Get() *AzureApplicationCredentials {
	return v.value
}

func (v *NullableAzureApplicationCredentials) Set(val *AzureApplicationCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureApplicationCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureApplicationCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureApplicationCredentials(val *AzureApplicationCredentials) *NullableAzureApplicationCredentials {
	return &NullableAzureApplicationCredentials{value: val, isSet: true}
}

func (v NullableAzureApplicationCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureApplicationCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


