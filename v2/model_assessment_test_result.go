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

// checks if the AssessmentTestResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssessmentTestResult{}

// AssessmentTestResult The result of a test run as part of assessment.
type AssessmentTestResult struct {
	// Specifies the test ID.
	TestId *string `json:"testId,omitempty"`
	// Specifies the kb link for diagnosing test failure.
	TestKbLink *string `json:"testKbLink,omitempty"`
	// Specifies the test name.
	TestName *string `json:"testName,omitempty"`
	// Specifies the test output message.
	TestOutput *string `json:"testOutput,omitempty"`
	// Specifies the test result.
	TestResult *string `json:"testResult,omitempty"`
}

// NewAssessmentTestResult instantiates a new AssessmentTestResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssessmentTestResult() *AssessmentTestResult {
	this := AssessmentTestResult{}
	return &this
}

// NewAssessmentTestResultWithDefaults instantiates a new AssessmentTestResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssessmentTestResultWithDefaults() *AssessmentTestResult {
	this := AssessmentTestResult{}
	return &this
}

// GetTestId returns the TestId field value if set, zero value otherwise.
func (o *AssessmentTestResult) GetTestId() string {
	if o == nil || IsNil(o.TestId) {
		var ret string
		return ret
	}
	return *o.TestId
}

// GetTestIdOk returns a tuple with the TestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentTestResult) GetTestIdOk() (*string, bool) {
	if o == nil || IsNil(o.TestId) {
		return nil, false
	}
	return o.TestId, true
}

// HasTestId returns a boolean if a field has been set.
func (o *AssessmentTestResult) HasTestId() bool {
	if o != nil && !IsNil(o.TestId) {
		return true
	}

	return false
}

// SetTestId gets a reference to the given string and assigns it to the TestId field.
func (o *AssessmentTestResult) SetTestId(v string) {
	o.TestId = &v
}

// GetTestKbLink returns the TestKbLink field value if set, zero value otherwise.
func (o *AssessmentTestResult) GetTestKbLink() string {
	if o == nil || IsNil(o.TestKbLink) {
		var ret string
		return ret
	}
	return *o.TestKbLink
}

// GetTestKbLinkOk returns a tuple with the TestKbLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentTestResult) GetTestKbLinkOk() (*string, bool) {
	if o == nil || IsNil(o.TestKbLink) {
		return nil, false
	}
	return o.TestKbLink, true
}

// HasTestKbLink returns a boolean if a field has been set.
func (o *AssessmentTestResult) HasTestKbLink() bool {
	if o != nil && !IsNil(o.TestKbLink) {
		return true
	}

	return false
}

// SetTestKbLink gets a reference to the given string and assigns it to the TestKbLink field.
func (o *AssessmentTestResult) SetTestKbLink(v string) {
	o.TestKbLink = &v
}

// GetTestName returns the TestName field value if set, zero value otherwise.
func (o *AssessmentTestResult) GetTestName() string {
	if o == nil || IsNil(o.TestName) {
		var ret string
		return ret
	}
	return *o.TestName
}

// GetTestNameOk returns a tuple with the TestName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentTestResult) GetTestNameOk() (*string, bool) {
	if o == nil || IsNil(o.TestName) {
		return nil, false
	}
	return o.TestName, true
}

// HasTestName returns a boolean if a field has been set.
func (o *AssessmentTestResult) HasTestName() bool {
	if o != nil && !IsNil(o.TestName) {
		return true
	}

	return false
}

// SetTestName gets a reference to the given string and assigns it to the TestName field.
func (o *AssessmentTestResult) SetTestName(v string) {
	o.TestName = &v
}

// GetTestOutput returns the TestOutput field value if set, zero value otherwise.
func (o *AssessmentTestResult) GetTestOutput() string {
	if o == nil || IsNil(o.TestOutput) {
		var ret string
		return ret
	}
	return *o.TestOutput
}

// GetTestOutputOk returns a tuple with the TestOutput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentTestResult) GetTestOutputOk() (*string, bool) {
	if o == nil || IsNil(o.TestOutput) {
		return nil, false
	}
	return o.TestOutput, true
}

// HasTestOutput returns a boolean if a field has been set.
func (o *AssessmentTestResult) HasTestOutput() bool {
	if o != nil && !IsNil(o.TestOutput) {
		return true
	}

	return false
}

// SetTestOutput gets a reference to the given string and assigns it to the TestOutput field.
func (o *AssessmentTestResult) SetTestOutput(v string) {
	o.TestOutput = &v
}

// GetTestResult returns the TestResult field value if set, zero value otherwise.
func (o *AssessmentTestResult) GetTestResult() string {
	if o == nil || IsNil(o.TestResult) {
		var ret string
		return ret
	}
	return *o.TestResult
}

// GetTestResultOk returns a tuple with the TestResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssessmentTestResult) GetTestResultOk() (*string, bool) {
	if o == nil || IsNil(o.TestResult) {
		return nil, false
	}
	return o.TestResult, true
}

// HasTestResult returns a boolean if a field has been set.
func (o *AssessmentTestResult) HasTestResult() bool {
	if o != nil && !IsNil(o.TestResult) {
		return true
	}

	return false
}

// SetTestResult gets a reference to the given string and assigns it to the TestResult field.
func (o *AssessmentTestResult) SetTestResult(v string) {
	o.TestResult = &v
}

func (o AssessmentTestResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssessmentTestResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TestId) {
		toSerialize["testId"] = o.TestId
	}
	if !IsNil(o.TestKbLink) {
		toSerialize["testKbLink"] = o.TestKbLink
	}
	if !IsNil(o.TestName) {
		toSerialize["testName"] = o.TestName
	}
	if !IsNil(o.TestOutput) {
		toSerialize["testOutput"] = o.TestOutput
	}
	if !IsNil(o.TestResult) {
		toSerialize["testResult"] = o.TestResult
	}
	return toSerialize, nil
}

type NullableAssessmentTestResult struct {
	value *AssessmentTestResult
	isSet bool
}

func (v NullableAssessmentTestResult) Get() *AssessmentTestResult {
	return v.value
}

func (v *NullableAssessmentTestResult) Set(val *AssessmentTestResult) {
	v.value = val
	v.isSet = true
}

func (v NullableAssessmentTestResult) IsSet() bool {
	return v.isSet
}

func (v *NullableAssessmentTestResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssessmentTestResult(val *AssessmentTestResult) *NullableAssessmentTestResult {
	return &NullableAssessmentTestResult{value: val, isSet: true}
}

func (v NullableAssessmentTestResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssessmentTestResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


