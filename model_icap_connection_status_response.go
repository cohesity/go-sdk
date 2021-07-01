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

// IcapConnectionStatusResponse Specifies Icap server connection status response.
type IcapConnectionStatusResponse struct {
	// Specifies the failed connection status of Icap servers.
	FailedConnectionStatus []string `json:"failedConnectionStatus,omitempty"`
	// Specifies the success connection status of Icap servers.
	SucceededConnectionStatus []string `json:"succeededConnectionStatus,omitempty"`
}

// NewIcapConnectionStatusResponse instantiates a new IcapConnectionStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIcapConnectionStatusResponse() *IcapConnectionStatusResponse {
	this := IcapConnectionStatusResponse{}
	return &this
}

// NewIcapConnectionStatusResponseWithDefaults instantiates a new IcapConnectionStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIcapConnectionStatusResponseWithDefaults() *IcapConnectionStatusResponse {
	this := IcapConnectionStatusResponse{}
	return &this
}

// GetFailedConnectionStatus returns the FailedConnectionStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IcapConnectionStatusResponse) GetFailedConnectionStatus() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.FailedConnectionStatus
}

// GetFailedConnectionStatusOk returns a tuple with the FailedConnectionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IcapConnectionStatusResponse) GetFailedConnectionStatusOk() (*[]string, bool) {
	if o == nil || o.FailedConnectionStatus == nil {
		return nil, false
	}
	return &o.FailedConnectionStatus, true
}

// HasFailedConnectionStatus returns a boolean if a field has been set.
func (o *IcapConnectionStatusResponse) HasFailedConnectionStatus() bool {
	if o != nil && o.FailedConnectionStatus != nil {
		return true
	}

	return false
}

// SetFailedConnectionStatus gets a reference to the given []string and assigns it to the FailedConnectionStatus field.
func (o *IcapConnectionStatusResponse) SetFailedConnectionStatus(v []string) {
	o.FailedConnectionStatus = v
}

// GetSucceededConnectionStatus returns the SucceededConnectionStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IcapConnectionStatusResponse) GetSucceededConnectionStatus() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.SucceededConnectionStatus
}

// GetSucceededConnectionStatusOk returns a tuple with the SucceededConnectionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IcapConnectionStatusResponse) GetSucceededConnectionStatusOk() (*[]string, bool) {
	if o == nil || o.SucceededConnectionStatus == nil {
		return nil, false
	}
	return &o.SucceededConnectionStatus, true
}

// HasSucceededConnectionStatus returns a boolean if a field has been set.
func (o *IcapConnectionStatusResponse) HasSucceededConnectionStatus() bool {
	if o != nil && o.SucceededConnectionStatus != nil {
		return true
	}

	return false
}

// SetSucceededConnectionStatus gets a reference to the given []string and assigns it to the SucceededConnectionStatus field.
func (o *IcapConnectionStatusResponse) SetSucceededConnectionStatus(v []string) {
	o.SucceededConnectionStatus = v
}

func (o IcapConnectionStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FailedConnectionStatus != nil {
		toSerialize["failedConnectionStatus"] = o.FailedConnectionStatus
	}
	if o.SucceededConnectionStatus != nil {
		toSerialize["succeededConnectionStatus"] = o.SucceededConnectionStatus
	}
	return json.Marshal(toSerialize)
}

type NullableIcapConnectionStatusResponse struct {
	value *IcapConnectionStatusResponse
	isSet bool
}

func (v NullableIcapConnectionStatusResponse) Get() *IcapConnectionStatusResponse {
	return v.value
}

func (v *NullableIcapConnectionStatusResponse) Set(val *IcapConnectionStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIcapConnectionStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIcapConnectionStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIcapConnectionStatusResponse(val *IcapConnectionStatusResponse) *NullableIcapConnectionStatusResponse {
	return &NullableIcapConnectionStatusResponse{value: val, isSet: true}
}

func (v NullableIcapConnectionStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIcapConnectionStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


