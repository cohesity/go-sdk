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

// checks if the UpdateKeystoneRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateKeystoneRequest{}

// UpdateKeystoneRequest Specifies the parameters to update a Keystone configuration.
type UpdateKeystoneRequest struct {
	AdminCreds KeystoneCredentialsAdminCreds `json:"adminCreds"`
	Scope KeystoneCredentialsScope `json:"scope"`
	// Specifies the url points to the Keystone service.
	AuthUrl NullableString `json:"authUrl"`
	// Specifies the Keystone configuration id.
	Id NullableInt64 `json:"id,omitempty"`
	// Specifies the Keystone configuration name.
	Name NullableString `json:"name"`
}

type _UpdateKeystoneRequest UpdateKeystoneRequest

// NewUpdateKeystoneRequest instantiates a new UpdateKeystoneRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateKeystoneRequest(adminCreds KeystoneCredentialsAdminCreds, scope KeystoneCredentialsScope, authUrl NullableString, name NullableString) *UpdateKeystoneRequest {
	this := UpdateKeystoneRequest{}
	this.AdminCreds = adminCreds
	this.Scope = scope
	this.AuthUrl = authUrl
	this.Name = name
	return &this
}

// NewUpdateKeystoneRequestWithDefaults instantiates a new UpdateKeystoneRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateKeystoneRequestWithDefaults() *UpdateKeystoneRequest {
	this := UpdateKeystoneRequest{}
	return &this
}

// GetAdminCreds returns the AdminCreds field value
func (o *UpdateKeystoneRequest) GetAdminCreds() KeystoneCredentialsAdminCreds {
	if o == nil {
		var ret KeystoneCredentialsAdminCreds
		return ret
	}

	return o.AdminCreds
}

// GetAdminCredsOk returns a tuple with the AdminCreds field value
// and a boolean to check if the value has been set.
func (o *UpdateKeystoneRequest) GetAdminCredsOk() (*KeystoneCredentialsAdminCreds, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdminCreds, true
}

// SetAdminCreds sets field value
func (o *UpdateKeystoneRequest) SetAdminCreds(v KeystoneCredentialsAdminCreds) {
	o.AdminCreds = v
}

// GetScope returns the Scope field value
func (o *UpdateKeystoneRequest) GetScope() KeystoneCredentialsScope {
	if o == nil {
		var ret KeystoneCredentialsScope
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *UpdateKeystoneRequest) GetScopeOk() (*KeystoneCredentialsScope, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *UpdateKeystoneRequest) SetScope(v KeystoneCredentialsScope) {
	o.Scope = v
}

// GetAuthUrl returns the AuthUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UpdateKeystoneRequest) GetAuthUrl() string {
	if o == nil || o.AuthUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.AuthUrl.Get()
}

// GetAuthUrlOk returns a tuple with the AuthUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateKeystoneRequest) GetAuthUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthUrl.Get(), o.AuthUrl.IsSet()
}

// SetAuthUrl sets field value
func (o *UpdateKeystoneRequest) SetAuthUrl(v string) {
	o.AuthUrl.Set(&v)
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateKeystoneRequest) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateKeystoneRequest) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *UpdateKeystoneRequest) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *UpdateKeystoneRequest) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *UpdateKeystoneRequest) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *UpdateKeystoneRequest) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UpdateKeystoneRequest) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateKeystoneRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *UpdateKeystoneRequest) SetName(v string) {
	o.Name.Set(&v)
}

func (o UpdateKeystoneRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateKeystoneRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["adminCreds"] = o.AdminCreds
	toSerialize["scope"] = o.Scope
	toSerialize["authUrl"] = o.AuthUrl.Get()
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	toSerialize["name"] = o.Name.Get()
	return toSerialize, nil
}

func (o *UpdateKeystoneRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"adminCreds",
		"scope",
		"authUrl",
		"name",
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

	varUpdateKeystoneRequest := _UpdateKeystoneRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateKeystoneRequest)

	if err != nil {
		return err
	}

	*o = UpdateKeystoneRequest(varUpdateKeystoneRequest)

	return err
}

type NullableUpdateKeystoneRequest struct {
	value *UpdateKeystoneRequest
	isSet bool
}

func (v NullableUpdateKeystoneRequest) Get() *UpdateKeystoneRequest {
	return v.value
}

func (v *NullableUpdateKeystoneRequest) Set(val *UpdateKeystoneRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateKeystoneRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateKeystoneRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateKeystoneRequest(val *UpdateKeystoneRequest) *NullableUpdateKeystoneRequest {
	return &NullableUpdateKeystoneRequest{value: val, isSet: true}
}

func (v NullableUpdateKeystoneRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateKeystoneRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


