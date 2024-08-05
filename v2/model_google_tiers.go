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

// checks if the GoogleTiers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GoogleTiers{}

// GoogleTiers Specifies Google tiers.
type GoogleTiers struct {
	// Specifies the tiers that are used to move the archived backup from current tier to next tier. The order of the tiers determines which tier will be used next for moving the archived backup. The first tier input should always be default tier where backup will be acrhived. Each tier specifies how much time after the backup will be moved to next tier from the current tier.
	Tiers []GoogleTier `json:"tiers"`
}

type _GoogleTiers GoogleTiers

// NewGoogleTiers instantiates a new GoogleTiers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoogleTiers(tiers []GoogleTier) *GoogleTiers {
	this := GoogleTiers{}
	this.Tiers = tiers
	return &this
}

// NewGoogleTiersWithDefaults instantiates a new GoogleTiers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleTiersWithDefaults() *GoogleTiers {
	this := GoogleTiers{}
	return &this
}

// GetTiers returns the Tiers field value
// If the value is explicit nil, the zero value for []GoogleTier will be returned
func (o *GoogleTiers) GetTiers() []GoogleTier {
	if o == nil {
		var ret []GoogleTier
		return ret
	}

	return o.Tiers
}

// GetTiersOk returns a tuple with the Tiers field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GoogleTiers) GetTiersOk() ([]GoogleTier, bool) {
	if o == nil || IsNil(o.Tiers) {
		return nil, false
	}
	return o.Tiers, true
}

// SetTiers sets field value
func (o *GoogleTiers) SetTiers(v []GoogleTier) {
	o.Tiers = v
}

func (o GoogleTiers) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GoogleTiers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Tiers != nil {
		toSerialize["tiers"] = o.Tiers
	}
	return toSerialize, nil
}

func (o *GoogleTiers) UnmarshalJSON(data []byte) (err error) {
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

	varGoogleTiers := _GoogleTiers{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGoogleTiers)

	if err != nil {
		return err
	}

	*o = GoogleTiers(varGoogleTiers)

	return err
}

type NullableGoogleTiers struct {
	value *GoogleTiers
	isSet bool
}

func (v NullableGoogleTiers) Get() *GoogleTiers {
	return v.value
}

func (v *NullableGoogleTiers) Set(val *GoogleTiers) {
	v.value = val
	v.isSet = true
}

func (v NullableGoogleTiers) IsSet() bool {
	return v.isSet
}

func (v *NullableGoogleTiers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoogleTiers(val *GoogleTiers) *NullableGoogleTiers {
	return &NullableGoogleTiers{value: val, isSet: true}
}

func (v NullableGoogleTiers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoogleTiers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


