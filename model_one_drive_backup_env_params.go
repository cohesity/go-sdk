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

// OneDriveBackupEnvParams Message to capture any additonal backup params for OneDrive within the Office365 environment.
type OneDriveBackupEnvParams struct {
	FilteringPolicy *FilteringPolicyProto `json:"filteringPolicy,omitempty"`
	// Specifies whether the OneDrive(s) for all the Office365 Users present in the protection job should be backed up.
	ShouldBackupOnedrive NullableBool `json:"shouldBackupOnedrive,omitempty"`
}

// NewOneDriveBackupEnvParams instantiates a new OneDriveBackupEnvParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOneDriveBackupEnvParams() *OneDriveBackupEnvParams {
	this := OneDriveBackupEnvParams{}
	return &this
}

// NewOneDriveBackupEnvParamsWithDefaults instantiates a new OneDriveBackupEnvParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOneDriveBackupEnvParamsWithDefaults() *OneDriveBackupEnvParams {
	this := OneDriveBackupEnvParams{}
	return &this
}

// GetFilteringPolicy returns the FilteringPolicy field value if set, zero value otherwise.
func (o *OneDriveBackupEnvParams) GetFilteringPolicy() FilteringPolicyProto {
	if o == nil || o.FilteringPolicy == nil {
		var ret FilteringPolicyProto
		return ret
	}
	return *o.FilteringPolicy
}

// GetFilteringPolicyOk returns a tuple with the FilteringPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OneDriveBackupEnvParams) GetFilteringPolicyOk() (*FilteringPolicyProto, bool) {
	if o == nil || o.FilteringPolicy == nil {
		return nil, false
	}
	return o.FilteringPolicy, true
}

// HasFilteringPolicy returns a boolean if a field has been set.
func (o *OneDriveBackupEnvParams) HasFilteringPolicy() bool {
	if o != nil && o.FilteringPolicy != nil {
		return true
	}

	return false
}

// SetFilteringPolicy gets a reference to the given FilteringPolicyProto and assigns it to the FilteringPolicy field.
func (o *OneDriveBackupEnvParams) SetFilteringPolicy(v FilteringPolicyProto) {
	o.FilteringPolicy = &v
}

// GetShouldBackupOnedrive returns the ShouldBackupOnedrive field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OneDriveBackupEnvParams) GetShouldBackupOnedrive() bool {
	if o == nil || o.ShouldBackupOnedrive.Get() == nil {
		var ret bool
		return ret
	}
	return *o.ShouldBackupOnedrive.Get()
}

// GetShouldBackupOnedriveOk returns a tuple with the ShouldBackupOnedrive field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OneDriveBackupEnvParams) GetShouldBackupOnedriveOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ShouldBackupOnedrive.Get(), o.ShouldBackupOnedrive.IsSet()
}

// HasShouldBackupOnedrive returns a boolean if a field has been set.
func (o *OneDriveBackupEnvParams) HasShouldBackupOnedrive() bool {
	if o != nil && o.ShouldBackupOnedrive.IsSet() {
		return true
	}

	return false
}

// SetShouldBackupOnedrive gets a reference to the given NullableBool and assigns it to the ShouldBackupOnedrive field.
func (o *OneDriveBackupEnvParams) SetShouldBackupOnedrive(v bool) {
	o.ShouldBackupOnedrive.Set(&v)
}
// SetShouldBackupOnedriveNil sets the value for ShouldBackupOnedrive to be an explicit nil
func (o *OneDriveBackupEnvParams) SetShouldBackupOnedriveNil() {
	o.ShouldBackupOnedrive.Set(nil)
}

// UnsetShouldBackupOnedrive ensures that no value is present for ShouldBackupOnedrive, not even an explicit nil
func (o *OneDriveBackupEnvParams) UnsetShouldBackupOnedrive() {
	o.ShouldBackupOnedrive.Unset()
}

func (o OneDriveBackupEnvParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FilteringPolicy != nil {
		toSerialize["filteringPolicy"] = o.FilteringPolicy
	}
	if o.ShouldBackupOnedrive.IsSet() {
		toSerialize["shouldBackupOnedrive"] = o.ShouldBackupOnedrive.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableOneDriveBackupEnvParams struct {
	value *OneDriveBackupEnvParams
	isSet bool
}

func (v NullableOneDriveBackupEnvParams) Get() *OneDriveBackupEnvParams {
	return v.value
}

func (v *NullableOneDriveBackupEnvParams) Set(val *OneDriveBackupEnvParams) {
	v.value = val
	v.isSet = true
}

func (v NullableOneDriveBackupEnvParams) IsSet() bool {
	return v.isSet
}

func (v *NullableOneDriveBackupEnvParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOneDriveBackupEnvParams(val *OneDriveBackupEnvParams) *NullableOneDriveBackupEnvParams {
	return &NullableOneDriveBackupEnvParams{value: val, isSet: true}
}

func (v NullableOneDriveBackupEnvParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOneDriveBackupEnvParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


