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

// LegalHoldings Specifies the legal holding of a Protection Source.
type LegalHoldings struct {
	// Specifies whether the source is put on legal hold or not.
	HoldForLegalPurpose NullableBool `json:"holdForLegalPurpose,omitempty"`
	// Specifies an Protection Source Id in the snapshot.
	ProtectionSourceId NullableInt64 `json:"protectionSourceId,omitempty"`
}

// NewLegalHoldings instantiates a new LegalHoldings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLegalHoldings() *LegalHoldings {
	this := LegalHoldings{}
	return &this
}

// NewLegalHoldingsWithDefaults instantiates a new LegalHoldings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLegalHoldingsWithDefaults() *LegalHoldings {
	this := LegalHoldings{}
	return &this
}

// GetHoldForLegalPurpose returns the HoldForLegalPurpose field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LegalHoldings) GetHoldForLegalPurpose() bool {
	if o == nil || o.HoldForLegalPurpose.Get() == nil {
		var ret bool
		return ret
	}
	return *o.HoldForLegalPurpose.Get()
}

// GetHoldForLegalPurposeOk returns a tuple with the HoldForLegalPurpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LegalHoldings) GetHoldForLegalPurposeOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.HoldForLegalPurpose.Get(), o.HoldForLegalPurpose.IsSet()
}

// HasHoldForLegalPurpose returns a boolean if a field has been set.
func (o *LegalHoldings) HasHoldForLegalPurpose() bool {
	if o != nil && o.HoldForLegalPurpose.IsSet() {
		return true
	}

	return false
}

// SetHoldForLegalPurpose gets a reference to the given NullableBool and assigns it to the HoldForLegalPurpose field.
func (o *LegalHoldings) SetHoldForLegalPurpose(v bool) {
	o.HoldForLegalPurpose.Set(&v)
}
// SetHoldForLegalPurposeNil sets the value for HoldForLegalPurpose to be an explicit nil
func (o *LegalHoldings) SetHoldForLegalPurposeNil() {
	o.HoldForLegalPurpose.Set(nil)
}

// UnsetHoldForLegalPurpose ensures that no value is present for HoldForLegalPurpose, not even an explicit nil
func (o *LegalHoldings) UnsetHoldForLegalPurpose() {
	o.HoldForLegalPurpose.Unset()
}

// GetProtectionSourceId returns the ProtectionSourceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LegalHoldings) GetProtectionSourceId() int64 {
	if o == nil || o.ProtectionSourceId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ProtectionSourceId.Get()
}

// GetProtectionSourceIdOk returns a tuple with the ProtectionSourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LegalHoldings) GetProtectionSourceIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ProtectionSourceId.Get(), o.ProtectionSourceId.IsSet()
}

// HasProtectionSourceId returns a boolean if a field has been set.
func (o *LegalHoldings) HasProtectionSourceId() bool {
	if o != nil && o.ProtectionSourceId.IsSet() {
		return true
	}

	return false
}

// SetProtectionSourceId gets a reference to the given NullableInt64 and assigns it to the ProtectionSourceId field.
func (o *LegalHoldings) SetProtectionSourceId(v int64) {
	o.ProtectionSourceId.Set(&v)
}
// SetProtectionSourceIdNil sets the value for ProtectionSourceId to be an explicit nil
func (o *LegalHoldings) SetProtectionSourceIdNil() {
	o.ProtectionSourceId.Set(nil)
}

// UnsetProtectionSourceId ensures that no value is present for ProtectionSourceId, not even an explicit nil
func (o *LegalHoldings) UnsetProtectionSourceId() {
	o.ProtectionSourceId.Unset()
}

func (o LegalHoldings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.HoldForLegalPurpose.IsSet() {
		toSerialize["holdForLegalPurpose"] = o.HoldForLegalPurpose.Get()
	}
	if o.ProtectionSourceId.IsSet() {
		toSerialize["protectionSourceId"] = o.ProtectionSourceId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableLegalHoldings struct {
	value *LegalHoldings
	isSet bool
}

func (v NullableLegalHoldings) Get() *LegalHoldings {
	return v.value
}

func (v *NullableLegalHoldings) Set(val *LegalHoldings) {
	v.value = val
	v.isSet = true
}

func (v NullableLegalHoldings) IsSet() bool {
	return v.isSet
}

func (v *NullableLegalHoldings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLegalHoldings(val *LegalHoldings) *NullableLegalHoldings {
	return &NullableLegalHoldings{value: val, isSet: true}
}

func (v NullableLegalHoldings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLegalHoldings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


