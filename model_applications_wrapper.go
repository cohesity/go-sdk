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

// ApplicationsWrapper ApplicationsWrapper is the struct to define the list of map-reduce applications.
type ApplicationsWrapper struct {
	// Applications specifies the list of available map-reduce applications in analytics workbench.
	Applications []MapReduceInfo `json:"applications,omitempty"`
}

// NewApplicationsWrapper instantiates a new ApplicationsWrapper object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationsWrapper() *ApplicationsWrapper {
	this := ApplicationsWrapper{}
	return &this
}

// NewApplicationsWrapperWithDefaults instantiates a new ApplicationsWrapper object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationsWrapperWithDefaults() *ApplicationsWrapper {
	this := ApplicationsWrapper{}
	return &this
}

// GetApplications returns the Applications field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationsWrapper) GetApplications() []MapReduceInfo {
	if o == nil  {
		var ret []MapReduceInfo
		return ret
	}
	return o.Applications
}

// GetApplicationsOk returns a tuple with the Applications field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationsWrapper) GetApplicationsOk() (*[]MapReduceInfo, bool) {
	if o == nil || o.Applications == nil {
		return nil, false
	}
	return &o.Applications, true
}

// HasApplications returns a boolean if a field has been set.
func (o *ApplicationsWrapper) HasApplications() bool {
	if o != nil && o.Applications != nil {
		return true
	}

	return false
}

// SetApplications gets a reference to the given []MapReduceInfo and assigns it to the Applications field.
func (o *ApplicationsWrapper) SetApplications(v []MapReduceInfo) {
	o.Applications = v
}

func (o ApplicationsWrapper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Applications != nil {
		toSerialize["applications"] = o.Applications
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationsWrapper struct {
	value *ApplicationsWrapper
	isSet bool
}

func (v NullableApplicationsWrapper) Get() *ApplicationsWrapper {
	return v.value
}

func (v *NullableApplicationsWrapper) Set(val *ApplicationsWrapper) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationsWrapper) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationsWrapper) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationsWrapper(val *ApplicationsWrapper) *NullableApplicationsWrapper {
	return &NullableApplicationsWrapper{value: val, isSet: true}
}

func (v NullableApplicationsWrapper) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationsWrapper) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


