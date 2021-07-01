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

// FileDistributionMetrics Specifies the File distribution metrics.
type FileDistributionMetrics struct {
	// Specifies the name of the metric.
	MetricName NullableString `json:"metricName,omitempty"`
	// Specifies the value of specified metric name.
	Value NullableInt64 `json:"value,omitempty"`
}

// NewFileDistributionMetrics instantiates a new FileDistributionMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileDistributionMetrics() *FileDistributionMetrics {
	this := FileDistributionMetrics{}
	return &this
}

// NewFileDistributionMetricsWithDefaults instantiates a new FileDistributionMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileDistributionMetricsWithDefaults() *FileDistributionMetrics {
	this := FileDistributionMetrics{}
	return &this
}

// GetMetricName returns the MetricName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileDistributionMetrics) GetMetricName() string {
	if o == nil || o.MetricName.Get() == nil {
		var ret string
		return ret
	}
	return *o.MetricName.Get()
}

// GetMetricNameOk returns a tuple with the MetricName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileDistributionMetrics) GetMetricNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MetricName.Get(), o.MetricName.IsSet()
}

// HasMetricName returns a boolean if a field has been set.
func (o *FileDistributionMetrics) HasMetricName() bool {
	if o != nil && o.MetricName.IsSet() {
		return true
	}

	return false
}

// SetMetricName gets a reference to the given NullableString and assigns it to the MetricName field.
func (o *FileDistributionMetrics) SetMetricName(v string) {
	o.MetricName.Set(&v)
}
// SetMetricNameNil sets the value for MetricName to be an explicit nil
func (o *FileDistributionMetrics) SetMetricNameNil() {
	o.MetricName.Set(nil)
}

// UnsetMetricName ensures that no value is present for MetricName, not even an explicit nil
func (o *FileDistributionMetrics) UnsetMetricName() {
	o.MetricName.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileDistributionMetrics) GetValue() int64 {
	if o == nil || o.Value.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileDistributionMetrics) GetValueOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *FileDistributionMetrics) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableInt64 and assigns it to the Value field.
func (o *FileDistributionMetrics) SetValue(v int64) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *FileDistributionMetrics) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *FileDistributionMetrics) UnsetValue() {
	o.Value.Unset()
}

func (o FileDistributionMetrics) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MetricName.IsSet() {
		toSerialize["metricName"] = o.MetricName.Get()
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableFileDistributionMetrics struct {
	value *FileDistributionMetrics
	isSet bool
}

func (v NullableFileDistributionMetrics) Get() *FileDistributionMetrics {
	return v.value
}

func (v *NullableFileDistributionMetrics) Set(val *FileDistributionMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableFileDistributionMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableFileDistributionMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileDistributionMetrics(val *FileDistributionMetrics) *NullableFileDistributionMetrics {
	return &NullableFileDistributionMetrics{value: val, isSet: true}
}

func (v NullableFileDistributionMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileDistributionMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


