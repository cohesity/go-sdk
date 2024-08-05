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

// checks if the McmRigelClaimRequestParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &McmRigelClaimRequestParams{}

// McmRigelClaimRequestParams Specifies the request params to claim Rigel to Helios.
type McmRigelClaimRequestParams struct {
	// Claim token used for authentication.
	ClaimToken string `json:"claimToken"`
	// Specifies the cluster id.
	ClusterId NullableInt64 `json:"clusterId,omitempty"`
	// Specifies the cluster incarnation id.
	ClusterIncarnationId NullableInt64 `json:"clusterIncarnationId,omitempty"`
	// Unique id for rigel instance.
	RigelGuid int64 `json:"rigelGuid"`
	// Specifies the Rigel IP.
	RigelIp NullableString `json:"rigelIp,omitempty"`
	// Specifies the Rigel name.
	RigelName NullableString `json:"rigelName,omitempty"`
	// Specifies the Rigel type that is being claimed.
	RigelType NullableString `json:"rigelType,omitempty"`
	// Specifies the Rigel Software version.
	SoftwareVersion NullableString `json:"softwareVersion,omitempty"`
}

type _McmRigelClaimRequestParams McmRigelClaimRequestParams

// NewMcmRigelClaimRequestParams instantiates a new McmRigelClaimRequestParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMcmRigelClaimRequestParams(claimToken string, rigelGuid int64) *McmRigelClaimRequestParams {
	this := McmRigelClaimRequestParams{}
	this.ClaimToken = claimToken
	this.RigelGuid = rigelGuid
	return &this
}

// NewMcmRigelClaimRequestParamsWithDefaults instantiates a new McmRigelClaimRequestParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMcmRigelClaimRequestParamsWithDefaults() *McmRigelClaimRequestParams {
	this := McmRigelClaimRequestParams{}
	return &this
}

// GetClaimToken returns the ClaimToken field value
func (o *McmRigelClaimRequestParams) GetClaimToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClaimToken
}

// GetClaimTokenOk returns a tuple with the ClaimToken field value
// and a boolean to check if the value has been set.
func (o *McmRigelClaimRequestParams) GetClaimTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClaimToken, true
}

// SetClaimToken sets field value
func (o *McmRigelClaimRequestParams) SetClaimToken(v string) {
	o.ClaimToken = v
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *McmRigelClaimRequestParams) GetClusterId() int64 {
	if o == nil || IsNil(o.ClusterId.Get()) {
		var ret int64
		return ret
	}
	return *o.ClusterId.Get()
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *McmRigelClaimRequestParams) GetClusterIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClusterId.Get(), o.ClusterId.IsSet()
}

// HasClusterId returns a boolean if a field has been set.
func (o *McmRigelClaimRequestParams) HasClusterId() bool {
	if o != nil && o.ClusterId.IsSet() {
		return true
	}

	return false
}

// SetClusterId gets a reference to the given NullableInt64 and assigns it to the ClusterId field.
func (o *McmRigelClaimRequestParams) SetClusterId(v int64) {
	o.ClusterId.Set(&v)
}
// SetClusterIdNil sets the value for ClusterId to be an explicit nil
func (o *McmRigelClaimRequestParams) SetClusterIdNil() {
	o.ClusterId.Set(nil)
}

// UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
func (o *McmRigelClaimRequestParams) UnsetClusterId() {
	o.ClusterId.Unset()
}

// GetClusterIncarnationId returns the ClusterIncarnationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *McmRigelClaimRequestParams) GetClusterIncarnationId() int64 {
	if o == nil || IsNil(o.ClusterIncarnationId.Get()) {
		var ret int64
		return ret
	}
	return *o.ClusterIncarnationId.Get()
}

// GetClusterIncarnationIdOk returns a tuple with the ClusterIncarnationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *McmRigelClaimRequestParams) GetClusterIncarnationIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClusterIncarnationId.Get(), o.ClusterIncarnationId.IsSet()
}

// HasClusterIncarnationId returns a boolean if a field has been set.
func (o *McmRigelClaimRequestParams) HasClusterIncarnationId() bool {
	if o != nil && o.ClusterIncarnationId.IsSet() {
		return true
	}

	return false
}

// SetClusterIncarnationId gets a reference to the given NullableInt64 and assigns it to the ClusterIncarnationId field.
func (o *McmRigelClaimRequestParams) SetClusterIncarnationId(v int64) {
	o.ClusterIncarnationId.Set(&v)
}
// SetClusterIncarnationIdNil sets the value for ClusterIncarnationId to be an explicit nil
func (o *McmRigelClaimRequestParams) SetClusterIncarnationIdNil() {
	o.ClusterIncarnationId.Set(nil)
}

// UnsetClusterIncarnationId ensures that no value is present for ClusterIncarnationId, not even an explicit nil
func (o *McmRigelClaimRequestParams) UnsetClusterIncarnationId() {
	o.ClusterIncarnationId.Unset()
}

// GetRigelGuid returns the RigelGuid field value
func (o *McmRigelClaimRequestParams) GetRigelGuid() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RigelGuid
}

// GetRigelGuidOk returns a tuple with the RigelGuid field value
// and a boolean to check if the value has been set.
func (o *McmRigelClaimRequestParams) GetRigelGuidOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RigelGuid, true
}

// SetRigelGuid sets field value
func (o *McmRigelClaimRequestParams) SetRigelGuid(v int64) {
	o.RigelGuid = v
}

// GetRigelIp returns the RigelIp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *McmRigelClaimRequestParams) GetRigelIp() string {
	if o == nil || IsNil(o.RigelIp.Get()) {
		var ret string
		return ret
	}
	return *o.RigelIp.Get()
}

// GetRigelIpOk returns a tuple with the RigelIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *McmRigelClaimRequestParams) GetRigelIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RigelIp.Get(), o.RigelIp.IsSet()
}

// HasRigelIp returns a boolean if a field has been set.
func (o *McmRigelClaimRequestParams) HasRigelIp() bool {
	if o != nil && o.RigelIp.IsSet() {
		return true
	}

	return false
}

// SetRigelIp gets a reference to the given NullableString and assigns it to the RigelIp field.
func (o *McmRigelClaimRequestParams) SetRigelIp(v string) {
	o.RigelIp.Set(&v)
}
// SetRigelIpNil sets the value for RigelIp to be an explicit nil
func (o *McmRigelClaimRequestParams) SetRigelIpNil() {
	o.RigelIp.Set(nil)
}

// UnsetRigelIp ensures that no value is present for RigelIp, not even an explicit nil
func (o *McmRigelClaimRequestParams) UnsetRigelIp() {
	o.RigelIp.Unset()
}

// GetRigelName returns the RigelName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *McmRigelClaimRequestParams) GetRigelName() string {
	if o == nil || IsNil(o.RigelName.Get()) {
		var ret string
		return ret
	}
	return *o.RigelName.Get()
}

// GetRigelNameOk returns a tuple with the RigelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *McmRigelClaimRequestParams) GetRigelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RigelName.Get(), o.RigelName.IsSet()
}

// HasRigelName returns a boolean if a field has been set.
func (o *McmRigelClaimRequestParams) HasRigelName() bool {
	if o != nil && o.RigelName.IsSet() {
		return true
	}

	return false
}

// SetRigelName gets a reference to the given NullableString and assigns it to the RigelName field.
func (o *McmRigelClaimRequestParams) SetRigelName(v string) {
	o.RigelName.Set(&v)
}
// SetRigelNameNil sets the value for RigelName to be an explicit nil
func (o *McmRigelClaimRequestParams) SetRigelNameNil() {
	o.RigelName.Set(nil)
}

// UnsetRigelName ensures that no value is present for RigelName, not even an explicit nil
func (o *McmRigelClaimRequestParams) UnsetRigelName() {
	o.RigelName.Unset()
}

// GetRigelType returns the RigelType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *McmRigelClaimRequestParams) GetRigelType() string {
	if o == nil || IsNil(o.RigelType.Get()) {
		var ret string
		return ret
	}
	return *o.RigelType.Get()
}

// GetRigelTypeOk returns a tuple with the RigelType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *McmRigelClaimRequestParams) GetRigelTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RigelType.Get(), o.RigelType.IsSet()
}

// HasRigelType returns a boolean if a field has been set.
func (o *McmRigelClaimRequestParams) HasRigelType() bool {
	if o != nil && o.RigelType.IsSet() {
		return true
	}

	return false
}

// SetRigelType gets a reference to the given NullableString and assigns it to the RigelType field.
func (o *McmRigelClaimRequestParams) SetRigelType(v string) {
	o.RigelType.Set(&v)
}
// SetRigelTypeNil sets the value for RigelType to be an explicit nil
func (o *McmRigelClaimRequestParams) SetRigelTypeNil() {
	o.RigelType.Set(nil)
}

// UnsetRigelType ensures that no value is present for RigelType, not even an explicit nil
func (o *McmRigelClaimRequestParams) UnsetRigelType() {
	o.RigelType.Unset()
}

// GetSoftwareVersion returns the SoftwareVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *McmRigelClaimRequestParams) GetSoftwareVersion() string {
	if o == nil || IsNil(o.SoftwareVersion.Get()) {
		var ret string
		return ret
	}
	return *o.SoftwareVersion.Get()
}

// GetSoftwareVersionOk returns a tuple with the SoftwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *McmRigelClaimRequestParams) GetSoftwareVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SoftwareVersion.Get(), o.SoftwareVersion.IsSet()
}

// HasSoftwareVersion returns a boolean if a field has been set.
func (o *McmRigelClaimRequestParams) HasSoftwareVersion() bool {
	if o != nil && o.SoftwareVersion.IsSet() {
		return true
	}

	return false
}

// SetSoftwareVersion gets a reference to the given NullableString and assigns it to the SoftwareVersion field.
func (o *McmRigelClaimRequestParams) SetSoftwareVersion(v string) {
	o.SoftwareVersion.Set(&v)
}
// SetSoftwareVersionNil sets the value for SoftwareVersion to be an explicit nil
func (o *McmRigelClaimRequestParams) SetSoftwareVersionNil() {
	o.SoftwareVersion.Set(nil)
}

// UnsetSoftwareVersion ensures that no value is present for SoftwareVersion, not even an explicit nil
func (o *McmRigelClaimRequestParams) UnsetSoftwareVersion() {
	o.SoftwareVersion.Unset()
}

func (o McmRigelClaimRequestParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o McmRigelClaimRequestParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["claimToken"] = o.ClaimToken
	if o.ClusterId.IsSet() {
		toSerialize["clusterId"] = o.ClusterId.Get()
	}
	if o.ClusterIncarnationId.IsSet() {
		toSerialize["clusterIncarnationId"] = o.ClusterIncarnationId.Get()
	}
	toSerialize["rigelGuid"] = o.RigelGuid
	if o.RigelIp.IsSet() {
		toSerialize["rigelIp"] = o.RigelIp.Get()
	}
	if o.RigelName.IsSet() {
		toSerialize["rigelName"] = o.RigelName.Get()
	}
	if o.RigelType.IsSet() {
		toSerialize["rigelType"] = o.RigelType.Get()
	}
	if o.SoftwareVersion.IsSet() {
		toSerialize["softwareVersion"] = o.SoftwareVersion.Get()
	}
	return toSerialize, nil
}

func (o *McmRigelClaimRequestParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"claimToken",
		"rigelGuid",
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

	varMcmRigelClaimRequestParams := _McmRigelClaimRequestParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMcmRigelClaimRequestParams)

	if err != nil {
		return err
	}

	*o = McmRigelClaimRequestParams(varMcmRigelClaimRequestParams)

	return err
}

type NullableMcmRigelClaimRequestParams struct {
	value *McmRigelClaimRequestParams
	isSet bool
}

func (v NullableMcmRigelClaimRequestParams) Get() *McmRigelClaimRequestParams {
	return v.value
}

func (v *NullableMcmRigelClaimRequestParams) Set(val *McmRigelClaimRequestParams) {
	v.value = val
	v.isSet = true
}

func (v NullableMcmRigelClaimRequestParams) IsSet() bool {
	return v.isSet
}

func (v *NullableMcmRigelClaimRequestParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMcmRigelClaimRequestParams(val *McmRigelClaimRequestParams) *NullableMcmRigelClaimRequestParams {
	return &NullableMcmRigelClaimRequestParams{value: val, isSet: true}
}

func (v NullableMcmRigelClaimRequestParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMcmRigelClaimRequestParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


