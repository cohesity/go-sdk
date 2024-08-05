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

// checks if the KmsConfigurationAddUpdateParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KmsConfigurationAddUpdateParams{}

// KmsConfigurationAddUpdateParams Parameters to update or add key management system(KMS) on the cluster.
type KmsConfigurationAddUpdateParams struct {
	// Ids of external targets used to assign the KMS for encryption. Once an external KMS (AWS KMS or KIMP KMS) is assigned to an external target, it cannot be changed.
	ExternalTargetIds []int64 `json:"externalTargetIds,omitempty"`
	KmipKmsParams *KmipKmsConfiguration `json:"kmipKmsParams,omitempty"`
	// Name of the KMS.
	Name string `json:"name"`
	// Ids of storage domains used to assign the KMS for encryption. Once an external KMS (AWS KMS or KIMP KMS) is assigned to a storage domain, it cannot be changed.
	StorageDomainIds []int64 `json:"storageDomainIds,omitempty"`
}

type _KmsConfigurationAddUpdateParams KmsConfigurationAddUpdateParams

// NewKmsConfigurationAddUpdateParams instantiates a new KmsConfigurationAddUpdateParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKmsConfigurationAddUpdateParams(name string) *KmsConfigurationAddUpdateParams {
	this := KmsConfigurationAddUpdateParams{}
	this.Name = name
	return &this
}

// NewKmsConfigurationAddUpdateParamsWithDefaults instantiates a new KmsConfigurationAddUpdateParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKmsConfigurationAddUpdateParamsWithDefaults() *KmsConfigurationAddUpdateParams {
	this := KmsConfigurationAddUpdateParams{}
	return &this
}

// GetExternalTargetIds returns the ExternalTargetIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KmsConfigurationAddUpdateParams) GetExternalTargetIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.ExternalTargetIds
}

// GetExternalTargetIdsOk returns a tuple with the ExternalTargetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KmsConfigurationAddUpdateParams) GetExternalTargetIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.ExternalTargetIds) {
		return nil, false
	}
	return o.ExternalTargetIds, true
}

// HasExternalTargetIds returns a boolean if a field has been set.
func (o *KmsConfigurationAddUpdateParams) HasExternalTargetIds() bool {
	if o != nil && !IsNil(o.ExternalTargetIds) {
		return true
	}

	return false
}

// SetExternalTargetIds gets a reference to the given []int64 and assigns it to the ExternalTargetIds field.
func (o *KmsConfigurationAddUpdateParams) SetExternalTargetIds(v []int64) {
	o.ExternalTargetIds = v
}

// GetKmipKmsParams returns the KmipKmsParams field value if set, zero value otherwise.
func (o *KmsConfigurationAddUpdateParams) GetKmipKmsParams() KmipKmsConfiguration {
	if o == nil || IsNil(o.KmipKmsParams) {
		var ret KmipKmsConfiguration
		return ret
	}
	return *o.KmipKmsParams
}

// GetKmipKmsParamsOk returns a tuple with the KmipKmsParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KmsConfigurationAddUpdateParams) GetKmipKmsParamsOk() (*KmipKmsConfiguration, bool) {
	if o == nil || IsNil(o.KmipKmsParams) {
		return nil, false
	}
	return o.KmipKmsParams, true
}

// HasKmipKmsParams returns a boolean if a field has been set.
func (o *KmsConfigurationAddUpdateParams) HasKmipKmsParams() bool {
	if o != nil && !IsNil(o.KmipKmsParams) {
		return true
	}

	return false
}

// SetKmipKmsParams gets a reference to the given KmipKmsConfiguration and assigns it to the KmipKmsParams field.
func (o *KmsConfigurationAddUpdateParams) SetKmipKmsParams(v KmipKmsConfiguration) {
	o.KmipKmsParams = &v
}

// GetName returns the Name field value
func (o *KmsConfigurationAddUpdateParams) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *KmsConfigurationAddUpdateParams) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *KmsConfigurationAddUpdateParams) SetName(v string) {
	o.Name = v
}

// GetStorageDomainIds returns the StorageDomainIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KmsConfigurationAddUpdateParams) GetStorageDomainIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.StorageDomainIds
}

// GetStorageDomainIdsOk returns a tuple with the StorageDomainIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KmsConfigurationAddUpdateParams) GetStorageDomainIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.StorageDomainIds) {
		return nil, false
	}
	return o.StorageDomainIds, true
}

// HasStorageDomainIds returns a boolean if a field has been set.
func (o *KmsConfigurationAddUpdateParams) HasStorageDomainIds() bool {
	if o != nil && !IsNil(o.StorageDomainIds) {
		return true
	}

	return false
}

// SetStorageDomainIds gets a reference to the given []int64 and assigns it to the StorageDomainIds field.
func (o *KmsConfigurationAddUpdateParams) SetStorageDomainIds(v []int64) {
	o.StorageDomainIds = v
}

func (o KmsConfigurationAddUpdateParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KmsConfigurationAddUpdateParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ExternalTargetIds != nil {
		toSerialize["externalTargetIds"] = o.ExternalTargetIds
	}
	if !IsNil(o.KmipKmsParams) {
		toSerialize["kmipKmsParams"] = o.KmipKmsParams
	}
	toSerialize["name"] = o.Name
	if o.StorageDomainIds != nil {
		toSerialize["storageDomainIds"] = o.StorageDomainIds
	}
	return toSerialize, nil
}

func (o *KmsConfigurationAddUpdateParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varKmsConfigurationAddUpdateParams := _KmsConfigurationAddUpdateParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKmsConfigurationAddUpdateParams)

	if err != nil {
		return err
	}

	*o = KmsConfigurationAddUpdateParams(varKmsConfigurationAddUpdateParams)

	return err
}

type NullableKmsConfigurationAddUpdateParams struct {
	value *KmsConfigurationAddUpdateParams
	isSet bool
}

func (v NullableKmsConfigurationAddUpdateParams) Get() *KmsConfigurationAddUpdateParams {
	return v.value
}

func (v *NullableKmsConfigurationAddUpdateParams) Set(val *KmsConfigurationAddUpdateParams) {
	v.value = val
	v.isSet = true
}

func (v NullableKmsConfigurationAddUpdateParams) IsSet() bool {
	return v.isSet
}

func (v *NullableKmsConfigurationAddUpdateParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKmsConfigurationAddUpdateParams(val *KmsConfigurationAddUpdateParams) *NullableKmsConfigurationAddUpdateParams {
	return &NullableKmsConfigurationAddUpdateParams{value: val, isSet: true}
}

func (v NullableKmsConfigurationAddUpdateParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKmsConfigurationAddUpdateParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


