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

// checks if the AlertList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertList{}

// AlertList Specifies the response of get alerts.
type AlertList struct {
	// Specifies the list of alerts.
	Alerts []AlertInfo `json:"alerts"`
}

type _AlertList AlertList

// NewAlertList instantiates a new AlertList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertList(alerts []AlertInfo) *AlertList {
	this := AlertList{}
	this.Alerts = alerts
	return &this
}

// NewAlertListWithDefaults instantiates a new AlertList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertListWithDefaults() *AlertList {
	this := AlertList{}
	return &this
}

// GetAlerts returns the Alerts field value
func (o *AlertList) GetAlerts() []AlertInfo {
	if o == nil {
		var ret []AlertInfo
		return ret
	}

	return o.Alerts
}

// GetAlertsOk returns a tuple with the Alerts field value
// and a boolean to check if the value has been set.
func (o *AlertList) GetAlertsOk() ([]AlertInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Alerts, true
}

// SetAlerts sets field value
func (o *AlertList) SetAlerts(v []AlertInfo) {
	o.Alerts = v
}

func (o AlertList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["alerts"] = o.Alerts
	return toSerialize, nil
}

func (o *AlertList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"alerts",
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

	varAlertList := _AlertList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAlertList)

	if err != nil {
		return err
	}

	*o = AlertList(varAlertList)

	return err
}

type NullableAlertList struct {
	value *AlertList
	isSet bool
}

func (v NullableAlertList) Get() *AlertList {
	return v.value
}

func (v *NullableAlertList) Set(val *AlertList) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertList) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertList(val *AlertList) *NullableAlertList {
	return &NullableAlertList{value: val, isSet: true}
}

func (v NullableAlertList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


