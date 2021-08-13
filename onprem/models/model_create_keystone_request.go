/*
 * Cohesity REST API
 *
 * Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

// package onprem
package models

import (
	"encoding/json"
	. "github.com/cohesity/go-sdk/onprem/utils"
	"fmt"
)

var _ = NullableBool{}

// CreateKeystoneRequest Specifies the parameters to create a Keystone configuration.
type CreateKeystoneRequest struct {
	// Specifies parameters related to Keystone administrator.
	AdminCreds KeystoneAdminParams `json:"adminCreds"`
	// Specifies parameters related to Keystone scope.
	Scope KeystoneScopeParams `json:"scope"`
	// Specifies the Keystone configuration name.
	Name NullableString `json:"name"`
	// Specifies the Keystone configuration id.
	Id NullableInt64 `json:"id,omitempty"`
	// Specifies the url points to the Keystone service.
	AuthUrl NullableString `json:"authUrl"`
}

// NewCreateKeystoneRequest instantiates a new CreateKeystoneRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateKeystoneRequest(adminCreds KeystoneAdminParams, scope KeystoneScopeParams, name NullableString, authUrl NullableString) *CreateKeystoneRequest {
	this := CreateKeystoneRequest{}
	this.AdminCreds = adminCreds
	this.Scope = scope
	this.Name = name
	this.AuthUrl = authUrl
	return &this
}

// NewCreateKeystoneRequestWithDefaults instantiates a new CreateKeystoneRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateKeystoneRequestWithDefaults() *CreateKeystoneRequest {
	this := CreateKeystoneRequest{}
	return &this
}

// GetAdminCreds returns the AdminCreds field value
func (o *CreateKeystoneRequest) GetAdminCreds() KeystoneAdminParams {
	if o == nil {
		var ret KeystoneAdminParams
		return ret
	}

	return o.AdminCreds
}

// GetAdminCredsOk returns a tuple with the AdminCreds field value
// and a boolean to check if the value has been set.
func (o *CreateKeystoneRequest) GetAdminCredsOk() (*KeystoneAdminParams, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AdminCreds, true
}

// SetAdminCreds sets field value
func (o *CreateKeystoneRequest) SetAdminCreds(v KeystoneAdminParams) {
	o.AdminCreds = v
}

// GetScope returns the Scope field value
func (o *CreateKeystoneRequest) GetScope() KeystoneScopeParams {
	if o == nil {
		var ret KeystoneScopeParams
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *CreateKeystoneRequest) GetScopeOk() (*KeystoneScopeParams, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *CreateKeystoneRequest) SetScope(v KeystoneScopeParams) {
	o.Scope = v
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreateKeystoneRequest) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateKeystoneRequest) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *CreateKeystoneRequest) SetName(v string) {
	o.Name.Set(&v)
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateKeystoneRequest) GetId() int64 {
	if o == nil || o.Id.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateKeystoneRequest) GetIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *CreateKeystoneRequest) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *CreateKeystoneRequest) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *CreateKeystoneRequest) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *CreateKeystoneRequest) UnsetId() {
	o.Id.Unset()
}

// GetAuthUrl returns the AuthUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreateKeystoneRequest) GetAuthUrl() string {
	if o == nil || o.AuthUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.AuthUrl.Get()
}

// GetAuthUrlOk returns a tuple with the AuthUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateKeystoneRequest) GetAuthUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AuthUrl.Get(), o.AuthUrl.IsSet()
}

// SetAuthUrl sets field value
func (o *CreateKeystoneRequest) SetAuthUrl(v string) {
	o.AuthUrl.Set(&v)
}

func (o CreateKeystoneRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["adminCreds"] = o.AdminCreds
	}
	if true {
		toSerialize["scope"] = o.Scope
	}
	if true {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if true {
		toSerialize["authUrl"] = o.AuthUrl.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCreateKeystoneRequest struct {
	value *CreateKeystoneRequest
	isSet bool
}

func (v NullableCreateKeystoneRequest) Get() *CreateKeystoneRequest {
	return v.value
}

func (v *NullableCreateKeystoneRequest) Set(val *CreateKeystoneRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateKeystoneRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateKeystoneRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateKeystoneRequest(val *CreateKeystoneRequest) *NullableCreateKeystoneRequest {
	return &NullableCreateKeystoneRequest{value: val, isSet: true}
}

func (v NullableCreateKeystoneRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateKeystoneRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o CreateKeystoneRequest) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}