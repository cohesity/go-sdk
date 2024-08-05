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

// checks if the AwsS3StandardParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsS3StandardParams{}

// AwsS3StandardParams Specifies the parameters which are specific to AWS related External Targets with storage class S3 Standard.
type AwsS3StandardParams struct {
	// Specifies the AWS External Target type.
	CloudType NullableString `json:"cloudType"`
	AwsCloudC2SParams *AwsCloudC2SParams `json:"awsCloudC2SParams,omitempty"`
	AwsCloudGovParams *AwsCloudGovParams `json:"awsCloudGovParams,omitempty"`
	AwsCloudStandardParams *AwsCloudStandardParams `json:"awsCloudStandardParams,omitempty"`
}

type _AwsS3StandardParams AwsS3StandardParams

// NewAwsS3StandardParams instantiates a new AwsS3StandardParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsS3StandardParams(cloudType NullableString) *AwsS3StandardParams {
	this := AwsS3StandardParams{}
	this.CloudType = cloudType
	return &this
}

// NewAwsS3StandardParamsWithDefaults instantiates a new AwsS3StandardParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsS3StandardParamsWithDefaults() *AwsS3StandardParams {
	this := AwsS3StandardParams{}
	return &this
}

// GetCloudType returns the CloudType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AwsS3StandardParams) GetCloudType() string {
	if o == nil || o.CloudType.Get() == nil {
		var ret string
		return ret
	}

	return *o.CloudType.Get()
}

// GetCloudTypeOk returns a tuple with the CloudType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsS3StandardParams) GetCloudTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CloudType.Get(), o.CloudType.IsSet()
}

// SetCloudType sets field value
func (o *AwsS3StandardParams) SetCloudType(v string) {
	o.CloudType.Set(&v)
}

// GetAwsCloudC2SParams returns the AwsCloudC2SParams field value if set, zero value otherwise.
func (o *AwsS3StandardParams) GetAwsCloudC2SParams() AwsCloudC2SParams {
	if o == nil || IsNil(o.AwsCloudC2SParams) {
		var ret AwsCloudC2SParams
		return ret
	}
	return *o.AwsCloudC2SParams
}

// GetAwsCloudC2SParamsOk returns a tuple with the AwsCloudC2SParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsS3StandardParams) GetAwsCloudC2SParamsOk() (*AwsCloudC2SParams, bool) {
	if o == nil || IsNil(o.AwsCloudC2SParams) {
		return nil, false
	}
	return o.AwsCloudC2SParams, true
}

// HasAwsCloudC2SParams returns a boolean if a field has been set.
func (o *AwsS3StandardParams) HasAwsCloudC2SParams() bool {
	if o != nil && !IsNil(o.AwsCloudC2SParams) {
		return true
	}

	return false
}

// SetAwsCloudC2SParams gets a reference to the given AwsCloudC2SParams and assigns it to the AwsCloudC2SParams field.
func (o *AwsS3StandardParams) SetAwsCloudC2SParams(v AwsCloudC2SParams) {
	o.AwsCloudC2SParams = &v
}

// GetAwsCloudGovParams returns the AwsCloudGovParams field value if set, zero value otherwise.
func (o *AwsS3StandardParams) GetAwsCloudGovParams() AwsCloudGovParams {
	if o == nil || IsNil(o.AwsCloudGovParams) {
		var ret AwsCloudGovParams
		return ret
	}
	return *o.AwsCloudGovParams
}

// GetAwsCloudGovParamsOk returns a tuple with the AwsCloudGovParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsS3StandardParams) GetAwsCloudGovParamsOk() (*AwsCloudGovParams, bool) {
	if o == nil || IsNil(o.AwsCloudGovParams) {
		return nil, false
	}
	return o.AwsCloudGovParams, true
}

// HasAwsCloudGovParams returns a boolean if a field has been set.
func (o *AwsS3StandardParams) HasAwsCloudGovParams() bool {
	if o != nil && !IsNil(o.AwsCloudGovParams) {
		return true
	}

	return false
}

// SetAwsCloudGovParams gets a reference to the given AwsCloudGovParams and assigns it to the AwsCloudGovParams field.
func (o *AwsS3StandardParams) SetAwsCloudGovParams(v AwsCloudGovParams) {
	o.AwsCloudGovParams = &v
}

// GetAwsCloudStandardParams returns the AwsCloudStandardParams field value if set, zero value otherwise.
func (o *AwsS3StandardParams) GetAwsCloudStandardParams() AwsCloudStandardParams {
	if o == nil || IsNil(o.AwsCloudStandardParams) {
		var ret AwsCloudStandardParams
		return ret
	}
	return *o.AwsCloudStandardParams
}

// GetAwsCloudStandardParamsOk returns a tuple with the AwsCloudStandardParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsS3StandardParams) GetAwsCloudStandardParamsOk() (*AwsCloudStandardParams, bool) {
	if o == nil || IsNil(o.AwsCloudStandardParams) {
		return nil, false
	}
	return o.AwsCloudStandardParams, true
}

// HasAwsCloudStandardParams returns a boolean if a field has been set.
func (o *AwsS3StandardParams) HasAwsCloudStandardParams() bool {
	if o != nil && !IsNil(o.AwsCloudStandardParams) {
		return true
	}

	return false
}

// SetAwsCloudStandardParams gets a reference to the given AwsCloudStandardParams and assigns it to the AwsCloudStandardParams field.
func (o *AwsS3StandardParams) SetAwsCloudStandardParams(v AwsCloudStandardParams) {
	o.AwsCloudStandardParams = &v
}

func (o AwsS3StandardParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsS3StandardParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cloudType"] = o.CloudType.Get()
	if !IsNil(o.AwsCloudC2SParams) {
		toSerialize["awsCloudC2SParams"] = o.AwsCloudC2SParams
	}
	if !IsNil(o.AwsCloudGovParams) {
		toSerialize["awsCloudGovParams"] = o.AwsCloudGovParams
	}
	if !IsNil(o.AwsCloudStandardParams) {
		toSerialize["awsCloudStandardParams"] = o.AwsCloudStandardParams
	}
	return toSerialize, nil
}

func (o *AwsS3StandardParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cloudType",
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

	varAwsS3StandardParams := _AwsS3StandardParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAwsS3StandardParams)

	if err != nil {
		return err
	}

	*o = AwsS3StandardParams(varAwsS3StandardParams)

	return err
}

type NullableAwsS3StandardParams struct {
	value *AwsS3StandardParams
	isSet bool
}

func (v NullableAwsS3StandardParams) Get() *AwsS3StandardParams {
	return v.value
}

func (v *NullableAwsS3StandardParams) Set(val *AwsS3StandardParams) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsS3StandardParams) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsS3StandardParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsS3StandardParams(val *AwsS3StandardParams) *NullableAwsS3StandardParams {
	return &NullableAwsS3StandardParams{value: val, isSet: true}
}

func (v NullableAwsS3StandardParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsS3StandardParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


