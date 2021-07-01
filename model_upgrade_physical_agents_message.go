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

// UpgradePhysicalAgentsMessage Specifies the status of an upgrade request.
type UpgradePhysicalAgentsMessage struct {
	// Specifies the status message returned after initiating an upgrade request. Status of each agent upgrade can be obtained by listing Physical Servers using the GET /public/protectionSources operation.
	Message NullableString `json:"message,omitempty"`
}

// NewUpgradePhysicalAgentsMessage instantiates a new UpgradePhysicalAgentsMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpgradePhysicalAgentsMessage() *UpgradePhysicalAgentsMessage {
	this := UpgradePhysicalAgentsMessage{}
	return &this
}

// NewUpgradePhysicalAgentsMessageWithDefaults instantiates a new UpgradePhysicalAgentsMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpgradePhysicalAgentsMessageWithDefaults() *UpgradePhysicalAgentsMessage {
	this := UpgradePhysicalAgentsMessage{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpgradePhysicalAgentsMessage) GetMessage() string {
	if o == nil || o.Message.Get() == nil {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpgradePhysicalAgentsMessage) GetMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *UpgradePhysicalAgentsMessage) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *UpgradePhysicalAgentsMessage) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *UpgradePhysicalAgentsMessage) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *UpgradePhysicalAgentsMessage) UnsetMessage() {
	o.Message.Unset()
}

func (o UpgradePhysicalAgentsMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableUpgradePhysicalAgentsMessage struct {
	value *UpgradePhysicalAgentsMessage
	isSet bool
}

func (v NullableUpgradePhysicalAgentsMessage) Get() *UpgradePhysicalAgentsMessage {
	return v.value
}

func (v *NullableUpgradePhysicalAgentsMessage) Set(val *UpgradePhysicalAgentsMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableUpgradePhysicalAgentsMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableUpgradePhysicalAgentsMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpgradePhysicalAgentsMessage(val *UpgradePhysicalAgentsMessage) *NullableUpgradePhysicalAgentsMessage {
	return &NullableUpgradePhysicalAgentsMessage{value: val, isSet: true}
}

func (v NullableUpgradePhysicalAgentsMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpgradePhysicalAgentsMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


