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

// checks if the ObjectSnapshotGenericNasParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectSnapshotGenericNasParams{}

// ObjectSnapshotGenericNasParams Specifies the parameters specific to Generic NAS type snapshot.
type ObjectSnapshotGenericNasParams struct {
	// Specifies a list of NAS mount protocols supported by this object.
	SupportedNasMountProtocols []string `json:"supportedNasMountProtocols,omitempty"`
}

// NewObjectSnapshotGenericNasParams instantiates a new ObjectSnapshotGenericNasParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectSnapshotGenericNasParams() *ObjectSnapshotGenericNasParams {
	this := ObjectSnapshotGenericNasParams{}
	return &this
}

// NewObjectSnapshotGenericNasParamsWithDefaults instantiates a new ObjectSnapshotGenericNasParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectSnapshotGenericNasParamsWithDefaults() *ObjectSnapshotGenericNasParams {
	this := ObjectSnapshotGenericNasParams{}
	return &this
}

// GetSupportedNasMountProtocols returns the SupportedNasMountProtocols field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ObjectSnapshotGenericNasParams) GetSupportedNasMountProtocols() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SupportedNasMountProtocols
}

// GetSupportedNasMountProtocolsOk returns a tuple with the SupportedNasMountProtocols field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ObjectSnapshotGenericNasParams) GetSupportedNasMountProtocolsOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedNasMountProtocols) {
		return nil, false
	}
	return o.SupportedNasMountProtocols, true
}

// HasSupportedNasMountProtocols returns a boolean if a field has been set.
func (o *ObjectSnapshotGenericNasParams) HasSupportedNasMountProtocols() bool {
	if o != nil && !IsNil(o.SupportedNasMountProtocols) {
		return true
	}

	return false
}

// SetSupportedNasMountProtocols gets a reference to the given []string and assigns it to the SupportedNasMountProtocols field.
func (o *ObjectSnapshotGenericNasParams) SetSupportedNasMountProtocols(v []string) {
	o.SupportedNasMountProtocols = v
}

func (o ObjectSnapshotGenericNasParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectSnapshotGenericNasParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.SupportedNasMountProtocols != nil {
		toSerialize["supportedNasMountProtocols"] = o.SupportedNasMountProtocols
	}
	return toSerialize, nil
}

type NullableObjectSnapshotGenericNasParams struct {
	value *ObjectSnapshotGenericNasParams
	isSet bool
}

func (v NullableObjectSnapshotGenericNasParams) Get() *ObjectSnapshotGenericNasParams {
	return v.value
}

func (v *NullableObjectSnapshotGenericNasParams) Set(val *ObjectSnapshotGenericNasParams) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectSnapshotGenericNasParams) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectSnapshotGenericNasParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectSnapshotGenericNasParams(val *ObjectSnapshotGenericNasParams) *NullableObjectSnapshotGenericNasParams {
	return &NullableObjectSnapshotGenericNasParams{value: val, isSet: true}
}

func (v NullableObjectSnapshotGenericNasParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectSnapshotGenericNasParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


