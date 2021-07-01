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

// ExchangeRestoreViewParameters struct for ExchangeRestoreViewParameters
type ExchangeRestoreViewParameters struct {
	// Specifies whether we should white-list the restore view for all the IP addresses. If this parameter is set to false, then only the machine on which the view is mounted will be white-listed.
	Endpoint NullableBool `json:"endpoint,omitempty"`
	// Specifies the path of the SMB share.
	MountPoint NullableString `json:"mountPoint,omitempty"`
	// Specifies the view to which the files of an Exchange database has to be cloned.
	ViewName NullableString `json:"viewName,omitempty"`
}

// NewExchangeRestoreViewParameters instantiates a new ExchangeRestoreViewParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangeRestoreViewParameters() *ExchangeRestoreViewParameters {
	this := ExchangeRestoreViewParameters{}
	return &this
}

// NewExchangeRestoreViewParametersWithDefaults instantiates a new ExchangeRestoreViewParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangeRestoreViewParametersWithDefaults() *ExchangeRestoreViewParameters {
	this := ExchangeRestoreViewParameters{}
	return &this
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExchangeRestoreViewParameters) GetEndpoint() bool {
	if o == nil || o.Endpoint.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Endpoint.Get()
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExchangeRestoreViewParameters) GetEndpointOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Endpoint.Get(), o.Endpoint.IsSet()
}

// HasEndpoint returns a boolean if a field has been set.
func (o *ExchangeRestoreViewParameters) HasEndpoint() bool {
	if o != nil && o.Endpoint.IsSet() {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given NullableBool and assigns it to the Endpoint field.
func (o *ExchangeRestoreViewParameters) SetEndpoint(v bool) {
	o.Endpoint.Set(&v)
}
// SetEndpointNil sets the value for Endpoint to be an explicit nil
func (o *ExchangeRestoreViewParameters) SetEndpointNil() {
	o.Endpoint.Set(nil)
}

// UnsetEndpoint ensures that no value is present for Endpoint, not even an explicit nil
func (o *ExchangeRestoreViewParameters) UnsetEndpoint() {
	o.Endpoint.Unset()
}

// GetMountPoint returns the MountPoint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExchangeRestoreViewParameters) GetMountPoint() string {
	if o == nil || o.MountPoint.Get() == nil {
		var ret string
		return ret
	}
	return *o.MountPoint.Get()
}

// GetMountPointOk returns a tuple with the MountPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExchangeRestoreViewParameters) GetMountPointOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MountPoint.Get(), o.MountPoint.IsSet()
}

// HasMountPoint returns a boolean if a field has been set.
func (o *ExchangeRestoreViewParameters) HasMountPoint() bool {
	if o != nil && o.MountPoint.IsSet() {
		return true
	}

	return false
}

// SetMountPoint gets a reference to the given NullableString and assigns it to the MountPoint field.
func (o *ExchangeRestoreViewParameters) SetMountPoint(v string) {
	o.MountPoint.Set(&v)
}
// SetMountPointNil sets the value for MountPoint to be an explicit nil
func (o *ExchangeRestoreViewParameters) SetMountPointNil() {
	o.MountPoint.Set(nil)
}

// UnsetMountPoint ensures that no value is present for MountPoint, not even an explicit nil
func (o *ExchangeRestoreViewParameters) UnsetMountPoint() {
	o.MountPoint.Unset()
}

// GetViewName returns the ViewName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExchangeRestoreViewParameters) GetViewName() string {
	if o == nil || o.ViewName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ViewName.Get()
}

// GetViewNameOk returns a tuple with the ViewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExchangeRestoreViewParameters) GetViewNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ViewName.Get(), o.ViewName.IsSet()
}

// HasViewName returns a boolean if a field has been set.
func (o *ExchangeRestoreViewParameters) HasViewName() bool {
	if o != nil && o.ViewName.IsSet() {
		return true
	}

	return false
}

// SetViewName gets a reference to the given NullableString and assigns it to the ViewName field.
func (o *ExchangeRestoreViewParameters) SetViewName(v string) {
	o.ViewName.Set(&v)
}
// SetViewNameNil sets the value for ViewName to be an explicit nil
func (o *ExchangeRestoreViewParameters) SetViewNameNil() {
	o.ViewName.Set(nil)
}

// UnsetViewName ensures that no value is present for ViewName, not even an explicit nil
func (o *ExchangeRestoreViewParameters) UnsetViewName() {
	o.ViewName.Unset()
}

func (o ExchangeRestoreViewParameters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Endpoint.IsSet() {
		toSerialize["endpoint"] = o.Endpoint.Get()
	}
	if o.MountPoint.IsSet() {
		toSerialize["mountPoint"] = o.MountPoint.Get()
	}
	if o.ViewName.IsSet() {
		toSerialize["viewName"] = o.ViewName.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableExchangeRestoreViewParameters struct {
	value *ExchangeRestoreViewParameters
	isSet bool
}

func (v NullableExchangeRestoreViewParameters) Get() *ExchangeRestoreViewParameters {
	return v.value
}

func (v *NullableExchangeRestoreViewParameters) Set(val *ExchangeRestoreViewParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeRestoreViewParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeRestoreViewParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeRestoreViewParameters(val *ExchangeRestoreViewParameters) *NullableExchangeRestoreViewParameters {
	return &NullableExchangeRestoreViewParameters{value: val, isSet: true}
}

func (v NullableExchangeRestoreViewParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeRestoreViewParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


