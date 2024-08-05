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

// checks if the HeliosAzureTiers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HeliosAzureTiers{}

// HeliosAzureTiers Specifies Azure tiers.
type HeliosAzureTiers struct {
	// Specifies the tiers that are used to move the archived backup from current tier to next tier. The order of the tiers determines which tier will be used next for moving the archived backup. The first tier input should always be default tier where backup will be acrhived. Each tier specifies how much time after the backup will be moved to next tier from the current tier.
	Tiers []HeliosAzureTier `json:"tiers"`
}

type _HeliosAzureTiers HeliosAzureTiers

// NewHeliosAzureTiers instantiates a new HeliosAzureTiers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeliosAzureTiers(tiers []HeliosAzureTier) *HeliosAzureTiers {
	this := HeliosAzureTiers{}
	this.Tiers = tiers
	return &this
}

// NewHeliosAzureTiersWithDefaults instantiates a new HeliosAzureTiers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeliosAzureTiersWithDefaults() *HeliosAzureTiers {
	this := HeliosAzureTiers{}
	return &this
}

// GetTiers returns the Tiers field value
// If the value is explicit nil, the zero value for []HeliosAzureTier will be returned
func (o *HeliosAzureTiers) GetTiers() []HeliosAzureTier {
	if o == nil {
		var ret []HeliosAzureTier
		return ret
	}

	return o.Tiers
}

// GetTiersOk returns a tuple with the Tiers field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeliosAzureTiers) GetTiersOk() ([]HeliosAzureTier, bool) {
	if o == nil || IsNil(o.Tiers) {
		return nil, false
	}
	return o.Tiers, true
}

// SetTiers sets field value
func (o *HeliosAzureTiers) SetTiers(v []HeliosAzureTier) {
	o.Tiers = v
}

func (o HeliosAzureTiers) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HeliosAzureTiers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Tiers != nil {
		toSerialize["tiers"] = o.Tiers
	}
	return toSerialize, nil
}

func (o *HeliosAzureTiers) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tiers",
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

	varHeliosAzureTiers := _HeliosAzureTiers{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHeliosAzureTiers)

	if err != nil {
		return err
	}

	*o = HeliosAzureTiers(varHeliosAzureTiers)

	return err
}

type NullableHeliosAzureTiers struct {
	value *HeliosAzureTiers
	isSet bool
}

func (v NullableHeliosAzureTiers) Get() *HeliosAzureTiers {
	return v.value
}

func (v *NullableHeliosAzureTiers) Set(val *HeliosAzureTiers) {
	v.value = val
	v.isSet = true
}

func (v NullableHeliosAzureTiers) IsSet() bool {
	return v.isSet
}

func (v *NullableHeliosAzureTiers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHeliosAzureTiers(val *HeliosAzureTiers) *NullableHeliosAzureTiers {
	return &NullableHeliosAzureTiers{value: val, isSet: true}
}

func (v NullableHeliosAzureTiers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHeliosAzureTiers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


