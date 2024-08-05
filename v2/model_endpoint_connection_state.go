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

// checks if the EndpointConnectionState type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EndpointConnectionState{}

// EndpointConnectionState Specify the connection state to each endpoint.
type EndpointConnectionState struct {
	// Specifies the name of the endpoint
	Name NullableString `json:"name,omitempty"`
	// Specifies the port of the endpoint
	Port NullableString `json:"port,omitempty"`
	// Specifies the results of the endpoints.
	Results []EndpointCheckResult `json:"results,omitempty"`
}

// NewEndpointConnectionState instantiates a new EndpointConnectionState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointConnectionState() *EndpointConnectionState {
	this := EndpointConnectionState{}
	return &this
}

// NewEndpointConnectionStateWithDefaults instantiates a new EndpointConnectionState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointConnectionStateWithDefaults() *EndpointConnectionState {
	this := EndpointConnectionState{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EndpointConnectionState) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EndpointConnectionState) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *EndpointConnectionState) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *EndpointConnectionState) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *EndpointConnectionState) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *EndpointConnectionState) UnsetName() {
	o.Name.Unset()
}

// GetPort returns the Port field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EndpointConnectionState) GetPort() string {
	if o == nil || IsNil(o.Port.Get()) {
		var ret string
		return ret
	}
	return *o.Port.Get()
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EndpointConnectionState) GetPortOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Port.Get(), o.Port.IsSet()
}

// HasPort returns a boolean if a field has been set.
func (o *EndpointConnectionState) HasPort() bool {
	if o != nil && o.Port.IsSet() {
		return true
	}

	return false
}

// SetPort gets a reference to the given NullableString and assigns it to the Port field.
func (o *EndpointConnectionState) SetPort(v string) {
	o.Port.Set(&v)
}
// SetPortNil sets the value for Port to be an explicit nil
func (o *EndpointConnectionState) SetPortNil() {
	o.Port.Set(nil)
}

// UnsetPort ensures that no value is present for Port, not even an explicit nil
func (o *EndpointConnectionState) UnsetPort() {
	o.Port.Unset()
}

// GetResults returns the Results field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EndpointConnectionState) GetResults() []EndpointCheckResult {
	if o == nil {
		var ret []EndpointCheckResult
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EndpointConnectionState) GetResultsOk() ([]EndpointCheckResult, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *EndpointConnectionState) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []EndpointCheckResult and assigns it to the Results field.
func (o *EndpointConnectionState) SetResults(v []EndpointCheckResult) {
	o.Results = v
}

func (o EndpointConnectionState) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointConnectionState) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Port.IsSet() {
		toSerialize["port"] = o.Port.Get()
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableEndpointConnectionState struct {
	value *EndpointConnectionState
	isSet bool
}

func (v NullableEndpointConnectionState) Get() *EndpointConnectionState {
	return v.value
}

func (v *NullableEndpointConnectionState) Set(val *EndpointConnectionState) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointConnectionState) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointConnectionState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointConnectionState(val *EndpointConnectionState) *NullableEndpointConnectionState {
	return &NullableEndpointConnectionState{value: val, isSet: true}
}

func (v NullableEndpointConnectionState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointConnectionState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


