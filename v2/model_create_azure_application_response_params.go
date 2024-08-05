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

// checks if the CreateAzureApplicationResponseParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAzureApplicationResponseParams{}

// CreateAzureApplicationResponseParams Specifies the response parameters containing the Azure apps created within the Microsoft365 domain.
type CreateAzureApplicationResponseParams struct {
	// Specifies a list of Microsoft365 azure application credentials needed to authenticate & authorize users for Office 365/Azure Workflows.
	Microsoft365AppCredentialsList []Office365AppCredentials `json:"microsoft365AppCredentialsList,omitempty"`
}

// NewCreateAzureApplicationResponseParams instantiates a new CreateAzureApplicationResponseParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAzureApplicationResponseParams() *CreateAzureApplicationResponseParams {
	this := CreateAzureApplicationResponseParams{}
	return &this
}

// NewCreateAzureApplicationResponseParamsWithDefaults instantiates a new CreateAzureApplicationResponseParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAzureApplicationResponseParamsWithDefaults() *CreateAzureApplicationResponseParams {
	this := CreateAzureApplicationResponseParams{}
	return &this
}

// GetMicrosoft365AppCredentialsList returns the Microsoft365AppCredentialsList field value if set, zero value otherwise.
func (o *CreateAzureApplicationResponseParams) GetMicrosoft365AppCredentialsList() []Office365AppCredentials {
	if o == nil || IsNil(o.Microsoft365AppCredentialsList) {
		var ret []Office365AppCredentials
		return ret
	}
	return o.Microsoft365AppCredentialsList
}

// GetMicrosoft365AppCredentialsListOk returns a tuple with the Microsoft365AppCredentialsList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAzureApplicationResponseParams) GetMicrosoft365AppCredentialsListOk() ([]Office365AppCredentials, bool) {
	if o == nil || IsNil(o.Microsoft365AppCredentialsList) {
		return nil, false
	}
	return o.Microsoft365AppCredentialsList, true
}

// HasMicrosoft365AppCredentialsList returns a boolean if a field has been set.
func (o *CreateAzureApplicationResponseParams) HasMicrosoft365AppCredentialsList() bool {
	if o != nil && !IsNil(o.Microsoft365AppCredentialsList) {
		return true
	}

	return false
}

// SetMicrosoft365AppCredentialsList gets a reference to the given []Office365AppCredentials and assigns it to the Microsoft365AppCredentialsList field.
func (o *CreateAzureApplicationResponseParams) SetMicrosoft365AppCredentialsList(v []Office365AppCredentials) {
	o.Microsoft365AppCredentialsList = v
}

func (o CreateAzureApplicationResponseParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAzureApplicationResponseParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Microsoft365AppCredentialsList) {
		toSerialize["microsoft365AppCredentialsList"] = o.Microsoft365AppCredentialsList
	}
	return toSerialize, nil
}

type NullableCreateAzureApplicationResponseParams struct {
	value *CreateAzureApplicationResponseParams
	isSet bool
}

func (v NullableCreateAzureApplicationResponseParams) Get() *CreateAzureApplicationResponseParams {
	return v.value
}

func (v *NullableCreateAzureApplicationResponseParams) Set(val *CreateAzureApplicationResponseParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAzureApplicationResponseParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAzureApplicationResponseParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAzureApplicationResponseParams(val *CreateAzureApplicationResponseParams) *NullableCreateAzureApplicationResponseParams {
	return &NullableCreateAzureApplicationResponseParams{value: val, isSet: true}
}

func (v NullableCreateAzureApplicationResponseParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAzureApplicationResponseParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


