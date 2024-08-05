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

// checks if the AzureHotBlobParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureHotBlobParams{}

// AzureHotBlobParams Specifies the parameters which are specific to Azure related with tier type Hot Blob
type AzureHotBlobParams struct {
	// Specifies the category of the external target.
	Category NullableString `json:"category"`
	// Specifies the access key to deploy Azure function to function app
	FunctionAppDeploymentKey NullableString `json:"functionAppDeploymentKey,omitempty"`
	// Specifies the name of the Azure function app, which is the host of Azure functions.
	FunctionAppName NullableString `json:"functionAppName,omitempty"`
}

type _AzureHotBlobParams AzureHotBlobParams

// NewAzureHotBlobParams instantiates a new AzureHotBlobParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureHotBlobParams(category NullableString) *AzureHotBlobParams {
	this := AzureHotBlobParams{}
	this.Category = category
	return &this
}

// NewAzureHotBlobParamsWithDefaults instantiates a new AzureHotBlobParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureHotBlobParamsWithDefaults() *AzureHotBlobParams {
	this := AzureHotBlobParams{}
	return &this
}

// GetCategory returns the Category field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AzureHotBlobParams) GetCategory() string {
	if o == nil || o.Category.Get() == nil {
		var ret string
		return ret
	}

	return *o.Category.Get()
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureHotBlobParams) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Category.Get(), o.Category.IsSet()
}

// SetCategory sets field value
func (o *AzureHotBlobParams) SetCategory(v string) {
	o.Category.Set(&v)
}

// GetFunctionAppDeploymentKey returns the FunctionAppDeploymentKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureHotBlobParams) GetFunctionAppDeploymentKey() string {
	if o == nil || IsNil(o.FunctionAppDeploymentKey.Get()) {
		var ret string
		return ret
	}
	return *o.FunctionAppDeploymentKey.Get()
}

// GetFunctionAppDeploymentKeyOk returns a tuple with the FunctionAppDeploymentKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureHotBlobParams) GetFunctionAppDeploymentKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FunctionAppDeploymentKey.Get(), o.FunctionAppDeploymentKey.IsSet()
}

// HasFunctionAppDeploymentKey returns a boolean if a field has been set.
func (o *AzureHotBlobParams) HasFunctionAppDeploymentKey() bool {
	if o != nil && o.FunctionAppDeploymentKey.IsSet() {
		return true
	}

	return false
}

// SetFunctionAppDeploymentKey gets a reference to the given NullableString and assigns it to the FunctionAppDeploymentKey field.
func (o *AzureHotBlobParams) SetFunctionAppDeploymentKey(v string) {
	o.FunctionAppDeploymentKey.Set(&v)
}
// SetFunctionAppDeploymentKeyNil sets the value for FunctionAppDeploymentKey to be an explicit nil
func (o *AzureHotBlobParams) SetFunctionAppDeploymentKeyNil() {
	o.FunctionAppDeploymentKey.Set(nil)
}

// UnsetFunctionAppDeploymentKey ensures that no value is present for FunctionAppDeploymentKey, not even an explicit nil
func (o *AzureHotBlobParams) UnsetFunctionAppDeploymentKey() {
	o.FunctionAppDeploymentKey.Unset()
}

// GetFunctionAppName returns the FunctionAppName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureHotBlobParams) GetFunctionAppName() string {
	if o == nil || IsNil(o.FunctionAppName.Get()) {
		var ret string
		return ret
	}
	return *o.FunctionAppName.Get()
}

// GetFunctionAppNameOk returns a tuple with the FunctionAppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureHotBlobParams) GetFunctionAppNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FunctionAppName.Get(), o.FunctionAppName.IsSet()
}

// HasFunctionAppName returns a boolean if a field has been set.
func (o *AzureHotBlobParams) HasFunctionAppName() bool {
	if o != nil && o.FunctionAppName.IsSet() {
		return true
	}

	return false
}

// SetFunctionAppName gets a reference to the given NullableString and assigns it to the FunctionAppName field.
func (o *AzureHotBlobParams) SetFunctionAppName(v string) {
	o.FunctionAppName.Set(&v)
}
// SetFunctionAppNameNil sets the value for FunctionAppName to be an explicit nil
func (o *AzureHotBlobParams) SetFunctionAppNameNil() {
	o.FunctionAppName.Set(nil)
}

// UnsetFunctionAppName ensures that no value is present for FunctionAppName, not even an explicit nil
func (o *AzureHotBlobParams) UnsetFunctionAppName() {
	o.FunctionAppName.Unset()
}

func (o AzureHotBlobParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureHotBlobParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["category"] = o.Category.Get()
	if o.FunctionAppDeploymentKey.IsSet() {
		toSerialize["functionAppDeploymentKey"] = o.FunctionAppDeploymentKey.Get()
	}
	if o.FunctionAppName.IsSet() {
		toSerialize["functionAppName"] = o.FunctionAppName.Get()
	}
	return toSerialize, nil
}

func (o *AzureHotBlobParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"category",
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

	varAzureHotBlobParams := _AzureHotBlobParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAzureHotBlobParams)

	if err != nil {
		return err
	}

	*o = AzureHotBlobParams(varAzureHotBlobParams)

	return err
}

type NullableAzureHotBlobParams struct {
	value *AzureHotBlobParams
	isSet bool
}

func (v NullableAzureHotBlobParams) Get() *AzureHotBlobParams {
	return v.value
}

func (v *NullableAzureHotBlobParams) Set(val *AzureHotBlobParams) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureHotBlobParams) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureHotBlobParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureHotBlobParams(val *AzureHotBlobParams) *NullableAzureHotBlobParams {
	return &NullableAzureHotBlobParams{value: val, isSet: true}
}

func (v NullableAzureHotBlobParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureHotBlobParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


