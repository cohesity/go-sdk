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

// checks if the AlertTarget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertTarget{}

// AlertTarget Specifies an alert target to receive an alert.
type AlertTarget struct {
	// Specifies an email address to receive an alert.
	EmailAddress NullableString `json:"emailAddress"`
	// Specifies the language of the delivery target. Default value is 'en-us'.
	Language NullableString `json:"language,omitempty"`
	// Specifies the recipient type of email recipient. Default value is 'kTo'.
	RecipientType NullableString `json:"recipientType,omitempty"`
}

type _AlertTarget AlertTarget

// NewAlertTarget instantiates a new AlertTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertTarget(emailAddress NullableString) *AlertTarget {
	this := AlertTarget{}
	this.EmailAddress = emailAddress
	return &this
}

// NewAlertTargetWithDefaults instantiates a new AlertTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertTargetWithDefaults() *AlertTarget {
	this := AlertTarget{}
	return &this
}

// GetEmailAddress returns the EmailAddress field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AlertTarget) GetEmailAddress() string {
	if o == nil || o.EmailAddress.Get() == nil {
		var ret string
		return ret
	}

	return *o.EmailAddress.Get()
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlertTarget) GetEmailAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmailAddress.Get(), o.EmailAddress.IsSet()
}

// SetEmailAddress sets field value
func (o *AlertTarget) SetEmailAddress(v string) {
	o.EmailAddress.Set(&v)
}

// GetLanguage returns the Language field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlertTarget) GetLanguage() string {
	if o == nil || IsNil(o.Language.Get()) {
		var ret string
		return ret
	}
	return *o.Language.Get()
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlertTarget) GetLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Language.Get(), o.Language.IsSet()
}

// HasLanguage returns a boolean if a field has been set.
func (o *AlertTarget) HasLanguage() bool {
	if o != nil && o.Language.IsSet() {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given NullableString and assigns it to the Language field.
func (o *AlertTarget) SetLanguage(v string) {
	o.Language.Set(&v)
}
// SetLanguageNil sets the value for Language to be an explicit nil
func (o *AlertTarget) SetLanguageNil() {
	o.Language.Set(nil)
}

// UnsetLanguage ensures that no value is present for Language, not even an explicit nil
func (o *AlertTarget) UnsetLanguage() {
	o.Language.Unset()
}

// GetRecipientType returns the RecipientType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlertTarget) GetRecipientType() string {
	if o == nil || IsNil(o.RecipientType.Get()) {
		var ret string
		return ret
	}
	return *o.RecipientType.Get()
}

// GetRecipientTypeOk returns a tuple with the RecipientType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlertTarget) GetRecipientTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecipientType.Get(), o.RecipientType.IsSet()
}

// HasRecipientType returns a boolean if a field has been set.
func (o *AlertTarget) HasRecipientType() bool {
	if o != nil && o.RecipientType.IsSet() {
		return true
	}

	return false
}

// SetRecipientType gets a reference to the given NullableString and assigns it to the RecipientType field.
func (o *AlertTarget) SetRecipientType(v string) {
	o.RecipientType.Set(&v)
}
// SetRecipientTypeNil sets the value for RecipientType to be an explicit nil
func (o *AlertTarget) SetRecipientTypeNil() {
	o.RecipientType.Set(nil)
}

// UnsetRecipientType ensures that no value is present for RecipientType, not even an explicit nil
func (o *AlertTarget) UnsetRecipientType() {
	o.RecipientType.Unset()
}

func (o AlertTarget) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertTarget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["emailAddress"] = o.EmailAddress.Get()
	if o.Language.IsSet() {
		toSerialize["language"] = o.Language.Get()
	}
	if o.RecipientType.IsSet() {
		toSerialize["recipientType"] = o.RecipientType.Get()
	}
	return toSerialize, nil
}

func (o *AlertTarget) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"emailAddress",
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

	varAlertTarget := _AlertTarget{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAlertTarget)

	if err != nil {
		return err
	}

	*o = AlertTarget(varAlertTarget)

	return err
}

type NullableAlertTarget struct {
	value *AlertTarget
	isSet bool
}

func (v NullableAlertTarget) Get() *AlertTarget {
	return v.value
}

func (v *NullableAlertTarget) Set(val *AlertTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertTarget(val *AlertTarget) *NullableAlertTarget {
	return &NullableAlertTarget{value: val, isSet: true}
}

func (v NullableAlertTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


