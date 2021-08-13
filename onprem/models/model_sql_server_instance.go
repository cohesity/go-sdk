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

// SQLServerInstance Specifies the details of a SQL server.
type SQLServerInstance struct {
	// Specifies the unique id of the SQL server instance.
	Id NullableString `json:"id,omitempty"`
	// Specifies the name of the SQL server instance.
	Name NullableString `json:"name,omitempty"`
	// Specifies the wehther the SQL server instance is online or not.
	IsOnline NullableString `json:"isOnline,omitempty"`
	// Specifies the information about endpoint associated with this SQL server instance.
	Endpoints []ResourceEndpoint `json:"endpoints,omitempty"`
	// Specifies whether this SQL server instance is a part of Failover cluster or not.
	IsPartofFCI NullableBool `json:"isPartofFCI,omitempty"`
}

// NewSQLServerInstance instantiates a new SQLServerInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSQLServerInstance() *SQLServerInstance {
	this := SQLServerInstance{}
	return &this
}

// NewSQLServerInstanceWithDefaults instantiates a new SQLServerInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSQLServerInstanceWithDefaults() *SQLServerInstance {
	this := SQLServerInstance{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SQLServerInstance) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SQLServerInstance) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *SQLServerInstance) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *SQLServerInstance) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *SQLServerInstance) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *SQLServerInstance) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SQLServerInstance) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SQLServerInstance) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *SQLServerInstance) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *SQLServerInstance) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *SQLServerInstance) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *SQLServerInstance) UnsetName() {
	o.Name.Unset()
}

// GetIsOnline returns the IsOnline field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SQLServerInstance) GetIsOnline() string {
	if o == nil || o.IsOnline.Get() == nil {
		var ret string
		return ret
	}
	return *o.IsOnline.Get()
}

// GetIsOnlineOk returns a tuple with the IsOnline field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SQLServerInstance) GetIsOnlineOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsOnline.Get(), o.IsOnline.IsSet()
}

// HasIsOnline returns a boolean if a field has been set.
func (o *SQLServerInstance) HasIsOnline() bool {
	if o != nil && o.IsOnline.IsSet() {
		return true
	}

	return false
}

// SetIsOnline gets a reference to the given NullableString and assigns it to the IsOnline field.
func (o *SQLServerInstance) SetIsOnline(v string) {
	o.IsOnline.Set(&v)
}
// SetIsOnlineNil sets the value for IsOnline to be an explicit nil
func (o *SQLServerInstance) SetIsOnlineNil() {
	o.IsOnline.Set(nil)
}

// UnsetIsOnline ensures that no value is present for IsOnline, not even an explicit nil
func (o *SQLServerInstance) UnsetIsOnline() {
	o.IsOnline.Unset()
}

// GetEndpoints returns the Endpoints field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SQLServerInstance) GetEndpoints() []ResourceEndpoint {
	if o == nil  {
		var ret []ResourceEndpoint
		return ret
	}
	return o.Endpoints
}

// GetEndpointsOk returns a tuple with the Endpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SQLServerInstance) GetEndpointsOk() (*[]ResourceEndpoint, bool) {
	if o == nil || o.Endpoints == nil {
		return nil, false
	}
	return &o.Endpoints, true
}

// HasEndpoints returns a boolean if a field has been set.
func (o *SQLServerInstance) HasEndpoints() bool {
	if o != nil && o.Endpoints != nil {
		return true
	}

	return false
}

// SetEndpoints gets a reference to the given []ResourceEndpoint and assigns it to the Endpoints field.
func (o *SQLServerInstance) SetEndpoints(v []ResourceEndpoint) {
	o.Endpoints = v
}

// GetIsPartofFCI returns the IsPartofFCI field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SQLServerInstance) GetIsPartofFCI() bool {
	if o == nil || o.IsPartofFCI.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsPartofFCI.Get()
}

// GetIsPartofFCIOk returns a tuple with the IsPartofFCI field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SQLServerInstance) GetIsPartofFCIOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsPartofFCI.Get(), o.IsPartofFCI.IsSet()
}

// HasIsPartofFCI returns a boolean if a field has been set.
func (o *SQLServerInstance) HasIsPartofFCI() bool {
	if o != nil && o.IsPartofFCI.IsSet() {
		return true
	}

	return false
}

// SetIsPartofFCI gets a reference to the given NullableBool and assigns it to the IsPartofFCI field.
func (o *SQLServerInstance) SetIsPartofFCI(v bool) {
	o.IsPartofFCI.Set(&v)
}
// SetIsPartofFCINil sets the value for IsPartofFCI to be an explicit nil
func (o *SQLServerInstance) SetIsPartofFCINil() {
	o.IsPartofFCI.Set(nil)
}

// UnsetIsPartofFCI ensures that no value is present for IsPartofFCI, not even an explicit nil
func (o *SQLServerInstance) UnsetIsPartofFCI() {
	o.IsPartofFCI.Unset()
}

func (o SQLServerInstance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.IsOnline.IsSet() {
		toSerialize["isOnline"] = o.IsOnline.Get()
	}
	if o.Endpoints != nil {
		toSerialize["endpoints"] = o.Endpoints
	}
	if o.IsPartofFCI.IsSet() {
		toSerialize["isPartofFCI"] = o.IsPartofFCI.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSQLServerInstance struct {
	value *SQLServerInstance
	isSet bool
}

func (v NullableSQLServerInstance) Get() *SQLServerInstance {
	return v.value
}

func (v *NullableSQLServerInstance) Set(val *SQLServerInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableSQLServerInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableSQLServerInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSQLServerInstance(val *SQLServerInstance) *NullableSQLServerInstance {
	return &NullableSQLServerInstance{value: val, isSet: true}
}

func (v NullableSQLServerInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSQLServerInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o SQLServerInstance) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}