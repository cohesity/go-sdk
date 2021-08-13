/*
 * Cohesity REST API
 *
 * Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

// package helios
package models

import (
	"encoding/json"
	. "github.com/cohesity/go-sdk/helios/utils"
)

var _ = NullableBool{}

// ArchivalExternalTargetParamsAllOf struct for ArchivalExternalTargetParamsAllOf
type ArchivalExternalTargetParamsAllOf struct {
	AzureParams *ArchivalAzureExternalTargetParams `json:"azureParams,omitempty"`
	GcpParams *ArchivalGcpExternalTargetParams `json:"gcpParams,omitempty"`
	AwsParams *ArchivalAwsExternalTargetParams `json:"awsParams,omitempty"`
	OracleParams *ArchivalOracleExternalTargetParams `json:"oracleParams,omitempty"`
	NasParams *ArchivalNasExternalTargetParams `json:"nasParams,omitempty"`
	QstarTapeParams *ArchivalQstarTapeExternalTargetParams `json:"qstarTapeParams,omitempty"`
	S3CompParams *ArchivalS3CompExternalTargetParams `json:"s3CompParams,omitempty"`
}

// NewArchivalExternalTargetParamsAllOf instantiates a new ArchivalExternalTargetParamsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArchivalExternalTargetParamsAllOf() *ArchivalExternalTargetParamsAllOf {
	this := ArchivalExternalTargetParamsAllOf{}
	return &this
}

// NewArchivalExternalTargetParamsAllOfWithDefaults instantiates a new ArchivalExternalTargetParamsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArchivalExternalTargetParamsAllOfWithDefaults() *ArchivalExternalTargetParamsAllOf {
	this := ArchivalExternalTargetParamsAllOf{}
	return &this
}

// GetAzureParams returns the AzureParams field value if set, zero value otherwise.
func (o *ArchivalExternalTargetParamsAllOf) GetAzureParams() ArchivalAzureExternalTargetParams {
	if o == nil || o.AzureParams == nil {
		var ret ArchivalAzureExternalTargetParams
		return ret
	}
	return *o.AzureParams
}

// GetAzureParamsOk returns a tuple with the AzureParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchivalExternalTargetParamsAllOf) GetAzureParamsOk() (*ArchivalAzureExternalTargetParams, bool) {
	if o == nil || o.AzureParams == nil {
		return nil, false
	}
	return o.AzureParams, true
}

// HasAzureParams returns a boolean if a field has been set.
func (o *ArchivalExternalTargetParamsAllOf) HasAzureParams() bool {
	if o != nil && o.AzureParams != nil {
		return true
	}

	return false
}

// SetAzureParams gets a reference to the given ArchivalAzureExternalTargetParams and assigns it to the AzureParams field.
func (o *ArchivalExternalTargetParamsAllOf) SetAzureParams(v ArchivalAzureExternalTargetParams) {
	o.AzureParams = &v
}

// GetGcpParams returns the GcpParams field value if set, zero value otherwise.
func (o *ArchivalExternalTargetParamsAllOf) GetGcpParams() ArchivalGcpExternalTargetParams {
	if o == nil || o.GcpParams == nil {
		var ret ArchivalGcpExternalTargetParams
		return ret
	}
	return *o.GcpParams
}

// GetGcpParamsOk returns a tuple with the GcpParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchivalExternalTargetParamsAllOf) GetGcpParamsOk() (*ArchivalGcpExternalTargetParams, bool) {
	if o == nil || o.GcpParams == nil {
		return nil, false
	}
	return o.GcpParams, true
}

// HasGcpParams returns a boolean if a field has been set.
func (o *ArchivalExternalTargetParamsAllOf) HasGcpParams() bool {
	if o != nil && o.GcpParams != nil {
		return true
	}

	return false
}

// SetGcpParams gets a reference to the given ArchivalGcpExternalTargetParams and assigns it to the GcpParams field.
func (o *ArchivalExternalTargetParamsAllOf) SetGcpParams(v ArchivalGcpExternalTargetParams) {
	o.GcpParams = &v
}

// GetAwsParams returns the AwsParams field value if set, zero value otherwise.
func (o *ArchivalExternalTargetParamsAllOf) GetAwsParams() ArchivalAwsExternalTargetParams {
	if o == nil || o.AwsParams == nil {
		var ret ArchivalAwsExternalTargetParams
		return ret
	}
	return *o.AwsParams
}

// GetAwsParamsOk returns a tuple with the AwsParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchivalExternalTargetParamsAllOf) GetAwsParamsOk() (*ArchivalAwsExternalTargetParams, bool) {
	if o == nil || o.AwsParams == nil {
		return nil, false
	}
	return o.AwsParams, true
}

// HasAwsParams returns a boolean if a field has been set.
func (o *ArchivalExternalTargetParamsAllOf) HasAwsParams() bool {
	if o != nil && o.AwsParams != nil {
		return true
	}

	return false
}

// SetAwsParams gets a reference to the given ArchivalAwsExternalTargetParams and assigns it to the AwsParams field.
func (o *ArchivalExternalTargetParamsAllOf) SetAwsParams(v ArchivalAwsExternalTargetParams) {
	o.AwsParams = &v
}

// GetOracleParams returns the OracleParams field value if set, zero value otherwise.
func (o *ArchivalExternalTargetParamsAllOf) GetOracleParams() ArchivalOracleExternalTargetParams {
	if o == nil || o.OracleParams == nil {
		var ret ArchivalOracleExternalTargetParams
		return ret
	}
	return *o.OracleParams
}

// GetOracleParamsOk returns a tuple with the OracleParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchivalExternalTargetParamsAllOf) GetOracleParamsOk() (*ArchivalOracleExternalTargetParams, bool) {
	if o == nil || o.OracleParams == nil {
		return nil, false
	}
	return o.OracleParams, true
}

// HasOracleParams returns a boolean if a field has been set.
func (o *ArchivalExternalTargetParamsAllOf) HasOracleParams() bool {
	if o != nil && o.OracleParams != nil {
		return true
	}

	return false
}

// SetOracleParams gets a reference to the given ArchivalOracleExternalTargetParams and assigns it to the OracleParams field.
func (o *ArchivalExternalTargetParamsAllOf) SetOracleParams(v ArchivalOracleExternalTargetParams) {
	o.OracleParams = &v
}

// GetNasParams returns the NasParams field value if set, zero value otherwise.
func (o *ArchivalExternalTargetParamsAllOf) GetNasParams() ArchivalNasExternalTargetParams {
	if o == nil || o.NasParams == nil {
		var ret ArchivalNasExternalTargetParams
		return ret
	}
	return *o.NasParams
}

// GetNasParamsOk returns a tuple with the NasParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchivalExternalTargetParamsAllOf) GetNasParamsOk() (*ArchivalNasExternalTargetParams, bool) {
	if o == nil || o.NasParams == nil {
		return nil, false
	}
	return o.NasParams, true
}

// HasNasParams returns a boolean if a field has been set.
func (o *ArchivalExternalTargetParamsAllOf) HasNasParams() bool {
	if o != nil && o.NasParams != nil {
		return true
	}

	return false
}

// SetNasParams gets a reference to the given ArchivalNasExternalTargetParams and assigns it to the NasParams field.
func (o *ArchivalExternalTargetParamsAllOf) SetNasParams(v ArchivalNasExternalTargetParams) {
	o.NasParams = &v
}

// GetQstarTapeParams returns the QstarTapeParams field value if set, zero value otherwise.
func (o *ArchivalExternalTargetParamsAllOf) GetQstarTapeParams() ArchivalQstarTapeExternalTargetParams {
	if o == nil || o.QstarTapeParams == nil {
		var ret ArchivalQstarTapeExternalTargetParams
		return ret
	}
	return *o.QstarTapeParams
}

// GetQstarTapeParamsOk returns a tuple with the QstarTapeParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchivalExternalTargetParamsAllOf) GetQstarTapeParamsOk() (*ArchivalQstarTapeExternalTargetParams, bool) {
	if o == nil || o.QstarTapeParams == nil {
		return nil, false
	}
	return o.QstarTapeParams, true
}

// HasQstarTapeParams returns a boolean if a field has been set.
func (o *ArchivalExternalTargetParamsAllOf) HasQstarTapeParams() bool {
	if o != nil && o.QstarTapeParams != nil {
		return true
	}

	return false
}

// SetQstarTapeParams gets a reference to the given ArchivalQstarTapeExternalTargetParams and assigns it to the QstarTapeParams field.
func (o *ArchivalExternalTargetParamsAllOf) SetQstarTapeParams(v ArchivalQstarTapeExternalTargetParams) {
	o.QstarTapeParams = &v
}

// GetS3CompParams returns the S3CompParams field value if set, zero value otherwise.
func (o *ArchivalExternalTargetParamsAllOf) GetS3CompParams() ArchivalS3CompExternalTargetParams {
	if o == nil || o.S3CompParams == nil {
		var ret ArchivalS3CompExternalTargetParams
		return ret
	}
	return *o.S3CompParams
}

// GetS3CompParamsOk returns a tuple with the S3CompParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchivalExternalTargetParamsAllOf) GetS3CompParamsOk() (*ArchivalS3CompExternalTargetParams, bool) {
	if o == nil || o.S3CompParams == nil {
		return nil, false
	}
	return o.S3CompParams, true
}

// HasS3CompParams returns a boolean if a field has been set.
func (o *ArchivalExternalTargetParamsAllOf) HasS3CompParams() bool {
	if o != nil && o.S3CompParams != nil {
		return true
	}

	return false
}

// SetS3CompParams gets a reference to the given ArchivalS3CompExternalTargetParams and assigns it to the S3CompParams field.
func (o *ArchivalExternalTargetParamsAllOf) SetS3CompParams(v ArchivalS3CompExternalTargetParams) {
	o.S3CompParams = &v
}

func (o ArchivalExternalTargetParamsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AzureParams != nil {
		toSerialize["azureParams"] = o.AzureParams
	}
	if o.GcpParams != nil {
		toSerialize["gcpParams"] = o.GcpParams
	}
	if o.AwsParams != nil {
		toSerialize["awsParams"] = o.AwsParams
	}
	if o.OracleParams != nil {
		toSerialize["oracleParams"] = o.OracleParams
	}
	if o.NasParams != nil {
		toSerialize["nasParams"] = o.NasParams
	}
	if o.QstarTapeParams != nil {
		toSerialize["qstarTapeParams"] = o.QstarTapeParams
	}
	if o.S3CompParams != nil {
		toSerialize["s3CompParams"] = o.S3CompParams
	}
	return json.Marshal(toSerialize)
}

type NullableArchivalExternalTargetParamsAllOf struct {
	value *ArchivalExternalTargetParamsAllOf
	isSet bool
}

func (v NullableArchivalExternalTargetParamsAllOf) Get() *ArchivalExternalTargetParamsAllOf {
	return v.value
}

func (v *NullableArchivalExternalTargetParamsAllOf) Set(val *ArchivalExternalTargetParamsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableArchivalExternalTargetParamsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableArchivalExternalTargetParamsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArchivalExternalTargetParamsAllOf(val *ArchivalExternalTargetParamsAllOf) *NullableArchivalExternalTargetParamsAllOf {
	return &NullableArchivalExternalTargetParamsAllOf{value: val, isSet: true}
}

func (v NullableArchivalExternalTargetParamsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArchivalExternalTargetParamsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


