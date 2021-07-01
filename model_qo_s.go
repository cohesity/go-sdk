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

// QoS Specifies the Quality of Service (QoS) Policy for the View.
type QoS struct {
	// Specifies the name of the QoS Policy used for the View.
	PrincipalId NullableInt64 `json:"principalId,omitempty"`
	// Specifies the name of the QoS Policy used for the View such as 'TestAndDev High', 'Backup Target SSD', 'Backup Target High' 'TestAndDev Low' and 'Backup Target Low'. For a complete list and descriptions, see the 'Create or Edit Views' topic in the documentation. If not specified, the default is 'Backup Target Low'.
	PrincipalName NullableString `json:"principalName,omitempty"`
}

// NewQoS instantiates a new QoS object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQoS() *QoS {
	this := QoS{}
	return &this
}

// NewQoSWithDefaults instantiates a new QoS object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQoSWithDefaults() *QoS {
	this := QoS{}
	return &this
}

// GetPrincipalId returns the PrincipalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QoS) GetPrincipalId() int64 {
	if o == nil || o.PrincipalId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.PrincipalId.Get()
}

// GetPrincipalIdOk returns a tuple with the PrincipalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QoS) GetPrincipalIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PrincipalId.Get(), o.PrincipalId.IsSet()
}

// HasPrincipalId returns a boolean if a field has been set.
func (o *QoS) HasPrincipalId() bool {
	if o != nil && o.PrincipalId.IsSet() {
		return true
	}

	return false
}

// SetPrincipalId gets a reference to the given NullableInt64 and assigns it to the PrincipalId field.
func (o *QoS) SetPrincipalId(v int64) {
	o.PrincipalId.Set(&v)
}
// SetPrincipalIdNil sets the value for PrincipalId to be an explicit nil
func (o *QoS) SetPrincipalIdNil() {
	o.PrincipalId.Set(nil)
}

// UnsetPrincipalId ensures that no value is present for PrincipalId, not even an explicit nil
func (o *QoS) UnsetPrincipalId() {
	o.PrincipalId.Unset()
}

// GetPrincipalName returns the PrincipalName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QoS) GetPrincipalName() string {
	if o == nil || o.PrincipalName.Get() == nil {
		var ret string
		return ret
	}
	return *o.PrincipalName.Get()
}

// GetPrincipalNameOk returns a tuple with the PrincipalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QoS) GetPrincipalNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PrincipalName.Get(), o.PrincipalName.IsSet()
}

// HasPrincipalName returns a boolean if a field has been set.
func (o *QoS) HasPrincipalName() bool {
	if o != nil && o.PrincipalName.IsSet() {
		return true
	}

	return false
}

// SetPrincipalName gets a reference to the given NullableString and assigns it to the PrincipalName field.
func (o *QoS) SetPrincipalName(v string) {
	o.PrincipalName.Set(&v)
}
// SetPrincipalNameNil sets the value for PrincipalName to be an explicit nil
func (o *QoS) SetPrincipalNameNil() {
	o.PrincipalName.Set(nil)
}

// UnsetPrincipalName ensures that no value is present for PrincipalName, not even an explicit nil
func (o *QoS) UnsetPrincipalName() {
	o.PrincipalName.Unset()
}

func (o QoS) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PrincipalId.IsSet() {
		toSerialize["principalId"] = o.PrincipalId.Get()
	}
	if o.PrincipalName.IsSet() {
		toSerialize["principalName"] = o.PrincipalName.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableQoS struct {
	value *QoS
	isSet bool
}

func (v NullableQoS) Get() *QoS {
	return v.value
}

func (v *NullableQoS) Set(val *QoS) {
	v.value = val
	v.isSet = true
}

func (v NullableQoS) IsSet() bool {
	return v.isSet
}

func (v *NullableQoS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQoS(val *QoS) *NullableQoS {
	return &NullableQoS{value: val, isSet: true}
}

func (v NullableQoS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQoS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


