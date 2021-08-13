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

// CreateClientcsrResponseBody Specifies the response to creating the CSRs.
type CreateClientcsrResponseBody struct {
	// Specifies the public key generated for this CSR for the server.
	PublicKeyServer NullableString `json:"publicKeyServer,omitempty"`
	// Specifies the CSR generated for the server.
	CsrServer NullableString `json:"csrServer,omitempty"`
	// Specifies the public key generated for this CSR for the client.
	PublicKeyClient NullableString `json:"publicKeyClient,omitempty"`
	// Specifies the CSR generated for the client.
	CsrClient NullableString `json:"csrClient,omitempty"`
	// Specifies the path to CSR generated for the server
	FileCsrServer NullableString `json:"fileCsrServer,omitempty"`
	// Specifies the path to CSR generated for the client
	FileCsrClient NullableString `json:"fileCsrClient,omitempty"`
}

// NewCreateClientcsrResponseBody instantiates a new CreateClientcsrResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateClientcsrResponseBody() *CreateClientcsrResponseBody {
	this := CreateClientcsrResponseBody{}
	return &this
}

// NewCreateClientcsrResponseBodyWithDefaults instantiates a new CreateClientcsrResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateClientcsrResponseBodyWithDefaults() *CreateClientcsrResponseBody {
	this := CreateClientcsrResponseBody{}
	return &this
}

// GetPublicKeyServer returns the PublicKeyServer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateClientcsrResponseBody) GetPublicKeyServer() string {
	if o == nil || o.PublicKeyServer.Get() == nil {
		var ret string
		return ret
	}
	return *o.PublicKeyServer.Get()
}

// GetPublicKeyServerOk returns a tuple with the PublicKeyServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateClientcsrResponseBody) GetPublicKeyServerOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PublicKeyServer.Get(), o.PublicKeyServer.IsSet()
}

// HasPublicKeyServer returns a boolean if a field has been set.
func (o *CreateClientcsrResponseBody) HasPublicKeyServer() bool {
	if o != nil && o.PublicKeyServer.IsSet() {
		return true
	}

	return false
}

// SetPublicKeyServer gets a reference to the given NullableString and assigns it to the PublicKeyServer field.
func (o *CreateClientcsrResponseBody) SetPublicKeyServer(v string) {
	o.PublicKeyServer.Set(&v)
}
// SetPublicKeyServerNil sets the value for PublicKeyServer to be an explicit nil
func (o *CreateClientcsrResponseBody) SetPublicKeyServerNil() {
	o.PublicKeyServer.Set(nil)
}

// UnsetPublicKeyServer ensures that no value is present for PublicKeyServer, not even an explicit nil
func (o *CreateClientcsrResponseBody) UnsetPublicKeyServer() {
	o.PublicKeyServer.Unset()
}

// GetCsrServer returns the CsrServer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateClientcsrResponseBody) GetCsrServer() string {
	if o == nil || o.CsrServer.Get() == nil {
		var ret string
		return ret
	}
	return *o.CsrServer.Get()
}

// GetCsrServerOk returns a tuple with the CsrServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateClientcsrResponseBody) GetCsrServerOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CsrServer.Get(), o.CsrServer.IsSet()
}

// HasCsrServer returns a boolean if a field has been set.
func (o *CreateClientcsrResponseBody) HasCsrServer() bool {
	if o != nil && o.CsrServer.IsSet() {
		return true
	}

	return false
}

// SetCsrServer gets a reference to the given NullableString and assigns it to the CsrServer field.
func (o *CreateClientcsrResponseBody) SetCsrServer(v string) {
	o.CsrServer.Set(&v)
}
// SetCsrServerNil sets the value for CsrServer to be an explicit nil
func (o *CreateClientcsrResponseBody) SetCsrServerNil() {
	o.CsrServer.Set(nil)
}

// UnsetCsrServer ensures that no value is present for CsrServer, not even an explicit nil
func (o *CreateClientcsrResponseBody) UnsetCsrServer() {
	o.CsrServer.Unset()
}

// GetPublicKeyClient returns the PublicKeyClient field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateClientcsrResponseBody) GetPublicKeyClient() string {
	if o == nil || o.PublicKeyClient.Get() == nil {
		var ret string
		return ret
	}
	return *o.PublicKeyClient.Get()
}

// GetPublicKeyClientOk returns a tuple with the PublicKeyClient field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateClientcsrResponseBody) GetPublicKeyClientOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PublicKeyClient.Get(), o.PublicKeyClient.IsSet()
}

// HasPublicKeyClient returns a boolean if a field has been set.
func (o *CreateClientcsrResponseBody) HasPublicKeyClient() bool {
	if o != nil && o.PublicKeyClient.IsSet() {
		return true
	}

	return false
}

// SetPublicKeyClient gets a reference to the given NullableString and assigns it to the PublicKeyClient field.
func (o *CreateClientcsrResponseBody) SetPublicKeyClient(v string) {
	o.PublicKeyClient.Set(&v)
}
// SetPublicKeyClientNil sets the value for PublicKeyClient to be an explicit nil
func (o *CreateClientcsrResponseBody) SetPublicKeyClientNil() {
	o.PublicKeyClient.Set(nil)
}

// UnsetPublicKeyClient ensures that no value is present for PublicKeyClient, not even an explicit nil
func (o *CreateClientcsrResponseBody) UnsetPublicKeyClient() {
	o.PublicKeyClient.Unset()
}

// GetCsrClient returns the CsrClient field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateClientcsrResponseBody) GetCsrClient() string {
	if o == nil || o.CsrClient.Get() == nil {
		var ret string
		return ret
	}
	return *o.CsrClient.Get()
}

// GetCsrClientOk returns a tuple with the CsrClient field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateClientcsrResponseBody) GetCsrClientOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CsrClient.Get(), o.CsrClient.IsSet()
}

// HasCsrClient returns a boolean if a field has been set.
func (o *CreateClientcsrResponseBody) HasCsrClient() bool {
	if o != nil && o.CsrClient.IsSet() {
		return true
	}

	return false
}

// SetCsrClient gets a reference to the given NullableString and assigns it to the CsrClient field.
func (o *CreateClientcsrResponseBody) SetCsrClient(v string) {
	o.CsrClient.Set(&v)
}
// SetCsrClientNil sets the value for CsrClient to be an explicit nil
func (o *CreateClientcsrResponseBody) SetCsrClientNil() {
	o.CsrClient.Set(nil)
}

// UnsetCsrClient ensures that no value is present for CsrClient, not even an explicit nil
func (o *CreateClientcsrResponseBody) UnsetCsrClient() {
	o.CsrClient.Unset()
}

// GetFileCsrServer returns the FileCsrServer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateClientcsrResponseBody) GetFileCsrServer() string {
	if o == nil || o.FileCsrServer.Get() == nil {
		var ret string
		return ret
	}
	return *o.FileCsrServer.Get()
}

// GetFileCsrServerOk returns a tuple with the FileCsrServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateClientcsrResponseBody) GetFileCsrServerOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FileCsrServer.Get(), o.FileCsrServer.IsSet()
}

// HasFileCsrServer returns a boolean if a field has been set.
func (o *CreateClientcsrResponseBody) HasFileCsrServer() bool {
	if o != nil && o.FileCsrServer.IsSet() {
		return true
	}

	return false
}

// SetFileCsrServer gets a reference to the given NullableString and assigns it to the FileCsrServer field.
func (o *CreateClientcsrResponseBody) SetFileCsrServer(v string) {
	o.FileCsrServer.Set(&v)
}
// SetFileCsrServerNil sets the value for FileCsrServer to be an explicit nil
func (o *CreateClientcsrResponseBody) SetFileCsrServerNil() {
	o.FileCsrServer.Set(nil)
}

// UnsetFileCsrServer ensures that no value is present for FileCsrServer, not even an explicit nil
func (o *CreateClientcsrResponseBody) UnsetFileCsrServer() {
	o.FileCsrServer.Unset()
}

// GetFileCsrClient returns the FileCsrClient field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateClientcsrResponseBody) GetFileCsrClient() string {
	if o == nil || o.FileCsrClient.Get() == nil {
		var ret string
		return ret
	}
	return *o.FileCsrClient.Get()
}

// GetFileCsrClientOk returns a tuple with the FileCsrClient field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateClientcsrResponseBody) GetFileCsrClientOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FileCsrClient.Get(), o.FileCsrClient.IsSet()
}

// HasFileCsrClient returns a boolean if a field has been set.
func (o *CreateClientcsrResponseBody) HasFileCsrClient() bool {
	if o != nil && o.FileCsrClient.IsSet() {
		return true
	}

	return false
}

// SetFileCsrClient gets a reference to the given NullableString and assigns it to the FileCsrClient field.
func (o *CreateClientcsrResponseBody) SetFileCsrClient(v string) {
	o.FileCsrClient.Set(&v)
}
// SetFileCsrClientNil sets the value for FileCsrClient to be an explicit nil
func (o *CreateClientcsrResponseBody) SetFileCsrClientNil() {
	o.FileCsrClient.Set(nil)
}

// UnsetFileCsrClient ensures that no value is present for FileCsrClient, not even an explicit nil
func (o *CreateClientcsrResponseBody) UnsetFileCsrClient() {
	o.FileCsrClient.Unset()
}

func (o CreateClientcsrResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PublicKeyServer.IsSet() {
		toSerialize["publicKeyServer"] = o.PublicKeyServer.Get()
	}
	if o.CsrServer.IsSet() {
		toSerialize["csrServer"] = o.CsrServer.Get()
	}
	if o.PublicKeyClient.IsSet() {
		toSerialize["publicKeyClient"] = o.PublicKeyClient.Get()
	}
	if o.CsrClient.IsSet() {
		toSerialize["csrClient"] = o.CsrClient.Get()
	}
	if o.FileCsrServer.IsSet() {
		toSerialize["fileCsrServer"] = o.FileCsrServer.Get()
	}
	if o.FileCsrClient.IsSet() {
		toSerialize["fileCsrClient"] = o.FileCsrClient.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCreateClientcsrResponseBody struct {
	value *CreateClientcsrResponseBody
	isSet bool
}

func (v NullableCreateClientcsrResponseBody) Get() *CreateClientcsrResponseBody {
	return v.value
}

func (v *NullableCreateClientcsrResponseBody) Set(val *CreateClientcsrResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateClientcsrResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateClientcsrResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateClientcsrResponseBody(val *CreateClientcsrResponseBody) *NullableCreateClientcsrResponseBody {
	return &NullableCreateClientcsrResponseBody{value: val, isSet: true}
}

func (v NullableCreateClientcsrResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateClientcsrResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o CreateClientcsrResponseBody) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}