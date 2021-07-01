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

// NewS3SecretAccessKey struct for NewS3SecretAccessKey
type NewS3SecretAccessKey struct {
	// Specifies the new S3 Secret Access key.
	NewKey NullableString `json:"newKey,omitempty"`
}

// NewNewS3SecretAccessKey instantiates a new NewS3SecretAccessKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewS3SecretAccessKey() *NewS3SecretAccessKey {
	this := NewS3SecretAccessKey{}
	return &this
}

// NewNewS3SecretAccessKeyWithDefaults instantiates a new NewS3SecretAccessKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewS3SecretAccessKeyWithDefaults() *NewS3SecretAccessKey {
	this := NewS3SecretAccessKey{}
	return &this
}

// GetNewKey returns the NewKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NewS3SecretAccessKey) GetNewKey() string {
	if o == nil || o.NewKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.NewKey.Get()
}

// GetNewKeyOk returns a tuple with the NewKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NewS3SecretAccessKey) GetNewKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NewKey.Get(), o.NewKey.IsSet()
}

// HasNewKey returns a boolean if a field has been set.
func (o *NewS3SecretAccessKey) HasNewKey() bool {
	if o != nil && o.NewKey.IsSet() {
		return true
	}

	return false
}

// SetNewKey gets a reference to the given NullableString and assigns it to the NewKey field.
func (o *NewS3SecretAccessKey) SetNewKey(v string) {
	o.NewKey.Set(&v)
}
// SetNewKeyNil sets the value for NewKey to be an explicit nil
func (o *NewS3SecretAccessKey) SetNewKeyNil() {
	o.NewKey.Set(nil)
}

// UnsetNewKey ensures that no value is present for NewKey, not even an explicit nil
func (o *NewS3SecretAccessKey) UnsetNewKey() {
	o.NewKey.Unset()
}

func (o NewS3SecretAccessKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NewKey.IsSet() {
		toSerialize["newKey"] = o.NewKey.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableNewS3SecretAccessKey struct {
	value *NewS3SecretAccessKey
	isSet bool
}

func (v NullableNewS3SecretAccessKey) Get() *NewS3SecretAccessKey {
	return v.value
}

func (v *NullableNewS3SecretAccessKey) Set(val *NewS3SecretAccessKey) {
	v.value = val
	v.isSet = true
}

func (v NullableNewS3SecretAccessKey) IsSet() bool {
	return v.isSet
}

func (v *NullableNewS3SecretAccessKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewS3SecretAccessKey(val *NewS3SecretAccessKey) *NullableNewS3SecretAccessKey {
	return &NullableNewS3SecretAccessKey{value: val, isSet: true}
}

func (v NullableNewS3SecretAccessKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewS3SecretAccessKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


