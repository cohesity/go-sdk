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

// EnvBackupParams Message to capture any additional environment specific backup params at the job level.
type EnvBackupParams struct {
	ExchangeBackupJobParams *ExchangeBackupJobParams `json:"exchangeBackupJobParams,omitempty"`
	FileStubbingParams *FileStubbingParams `json:"fileStubbingParams,omitempty"`
	FileUptieringParams *FileUptieringParams `json:"fileUptieringParams,omitempty"`
	HypervBackupParams *HyperVBackupEnvParams `json:"hypervBackupParams,omitempty"`
	NasBackupParams *NasBackupParams `json:"nasBackupParams,omitempty"`
	NosqlBackupJobParams *NoSqlBackupJobParams `json:"nosqlBackupJobParams,omitempty"`
	O365BackupParams *O365BackupEnvParams `json:"o365BackupParams,omitempty"`
	OracleBackupJobParams *OracleBackupJobParams `json:"oracleBackupJobParams,omitempty"`
	OutlookBackupParams *OutlookBackupEnvParams `json:"outlookBackupParams,omitempty"`
	PhysicalBackupParams *PhysicalBackupEnvParams `json:"physicalBackupParams,omitempty"`
	SnapshotManagerParams *SnapshotManagerParams `json:"snapshotManagerParams,omitempty"`
	SqlBackupJobParams *SqlBackupJobParams `json:"sqlBackupJobParams,omitempty"`
	VmwareBackupParams *VMwareBackupEnvParams `json:"vmwareBackupParams,omitempty"`
}

// NewEnvBackupParams instantiates a new EnvBackupParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvBackupParams() *EnvBackupParams {
	this := EnvBackupParams{}
	return &this
}

// NewEnvBackupParamsWithDefaults instantiates a new EnvBackupParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvBackupParamsWithDefaults() *EnvBackupParams {
	this := EnvBackupParams{}
	return &this
}

// GetExchangeBackupJobParams returns the ExchangeBackupJobParams field value if set, zero value otherwise.
func (o *EnvBackupParams) GetExchangeBackupJobParams() ExchangeBackupJobParams {
	if o == nil || o.ExchangeBackupJobParams == nil {
		var ret ExchangeBackupJobParams
		return ret
	}
	return *o.ExchangeBackupJobParams
}

// GetExchangeBackupJobParamsOk returns a tuple with the ExchangeBackupJobParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvBackupParams) GetExchangeBackupJobParamsOk() (*ExchangeBackupJobParams, bool) {
	if o == nil || o.ExchangeBackupJobParams == nil {
		return nil, false
	}
	return o.ExchangeBackupJobParams, true
}

// HasExchangeBackupJobParams returns a boolean if a field has been set.
func (o *EnvBackupParams) HasExchangeBackupJobParams() bool {
	if o != nil && o.ExchangeBackupJobParams != nil {
		return true
	}

	return false
}

// SetExchangeBackupJobParams gets a reference to the given ExchangeBackupJobParams and assigns it to the ExchangeBackupJobParams field.
func (o *EnvBackupParams) SetExchangeBackupJobParams(v ExchangeBackupJobParams) {
	o.ExchangeBackupJobParams = &v
}

// GetFileStubbingParams returns the FileStubbingParams field value if set, zero value otherwise.
func (o *EnvBackupParams) GetFileStubbingParams() FileStubbingParams {
	if o == nil || o.FileStubbingParams == nil {
		var ret FileStubbingParams
		return ret
	}
	return *o.FileStubbingParams
}

// GetFileStubbingParamsOk returns a tuple with the FileStubbingParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvBackupParams) GetFileStubbingParamsOk() (*FileStubbingParams, bool) {
	if o == nil || o.FileStubbingParams == nil {
		return nil, false
	}
	return o.FileStubbingParams, true
}

// HasFileStubbingParams returns a boolean if a field has been set.
func (o *EnvBackupParams) HasFileStubbingParams() bool {
	if o != nil && o.FileStubbingParams != nil {
		return true
	}

	return false
}

// SetFileStubbingParams gets a reference to the given FileStubbingParams and assigns it to the FileStubbingParams field.
func (o *EnvBackupParams) SetFileStubbingParams(v FileStubbingParams) {
	o.FileStubbingParams = &v
}

// GetFileUptieringParams returns the FileUptieringParams field value if set, zero value otherwise.
func (o *EnvBackupParams) GetFileUptieringParams() FileUptieringParams {
	if o == nil || o.FileUptieringParams == nil {
		var ret FileUptieringParams
		return ret
	}
	return *o.FileUptieringParams
}

// GetFileUptieringParamsOk returns a tuple with the FileUptieringParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvBackupParams) GetFileUptieringParamsOk() (*FileUptieringParams, bool) {
	if o == nil || o.FileUptieringParams == nil {
		return nil, false
	}
	return o.FileUptieringParams, true
}

// HasFileUptieringParams returns a boolean if a field has been set.
func (o *EnvBackupParams) HasFileUptieringParams() bool {
	if o != nil && o.FileUptieringParams != nil {
		return true
	}

	return false
}

// SetFileUptieringParams gets a reference to the given FileUptieringParams and assigns it to the FileUptieringParams field.
func (o *EnvBackupParams) SetFileUptieringParams(v FileUptieringParams) {
	o.FileUptieringParams = &v
}

// GetHypervBackupParams returns the HypervBackupParams field value if set, zero value otherwise.
func (o *EnvBackupParams) GetHypervBackupParams() HyperVBackupEnvParams {
	if o == nil || o.HypervBackupParams == nil {
		var ret HyperVBackupEnvParams
		return ret
	}
	return *o.HypervBackupParams
}

// GetHypervBackupParamsOk returns a tuple with the HypervBackupParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvBackupParams) GetHypervBackupParamsOk() (*HyperVBackupEnvParams, bool) {
	if o == nil || o.HypervBackupParams == nil {
		return nil, false
	}
	return o.HypervBackupParams, true
}

// HasHypervBackupParams returns a boolean if a field has been set.
func (o *EnvBackupParams) HasHypervBackupParams() bool {
	if o != nil && o.HypervBackupParams != nil {
		return true
	}

	return false
}

// SetHypervBackupParams gets a reference to the given HyperVBackupEnvParams and assigns it to the HypervBackupParams field.
func (o *EnvBackupParams) SetHypervBackupParams(v HyperVBackupEnvParams) {
	o.HypervBackupParams = &v
}

// GetNasBackupParams returns the NasBackupParams field value if set, zero value otherwise.
func (o *EnvBackupParams) GetNasBackupParams() NasBackupParams {
	if o == nil || o.NasBackupParams == nil {
		var ret NasBackupParams
		return ret
	}
	return *o.NasBackupParams
}

// GetNasBackupParamsOk returns a tuple with the NasBackupParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvBackupParams) GetNasBackupParamsOk() (*NasBackupParams, bool) {
	if o == nil || o.NasBackupParams == nil {
		return nil, false
	}
	return o.NasBackupParams, true
}

// HasNasBackupParams returns a boolean if a field has been set.
func (o *EnvBackupParams) HasNasBackupParams() bool {
	if o != nil && o.NasBackupParams != nil {
		return true
	}

	return false
}

// SetNasBackupParams gets a reference to the given NasBackupParams and assigns it to the NasBackupParams field.
func (o *EnvBackupParams) SetNasBackupParams(v NasBackupParams) {
	o.NasBackupParams = &v
}

// GetNosqlBackupJobParams returns the NosqlBackupJobParams field value if set, zero value otherwise.
func (o *EnvBackupParams) GetNosqlBackupJobParams() NoSqlBackupJobParams {
	if o == nil || o.NosqlBackupJobParams == nil {
		var ret NoSqlBackupJobParams
		return ret
	}
	return *o.NosqlBackupJobParams
}

// GetNosqlBackupJobParamsOk returns a tuple with the NosqlBackupJobParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvBackupParams) GetNosqlBackupJobParamsOk() (*NoSqlBackupJobParams, bool) {
	if o == nil || o.NosqlBackupJobParams == nil {
		return nil, false
	}
	return o.NosqlBackupJobParams, true
}

// HasNosqlBackupJobParams returns a boolean if a field has been set.
func (o *EnvBackupParams) HasNosqlBackupJobParams() bool {
	if o != nil && o.NosqlBackupJobParams != nil {
		return true
	}

	return false
}

// SetNosqlBackupJobParams gets a reference to the given NoSqlBackupJobParams and assigns it to the NosqlBackupJobParams field.
func (o *EnvBackupParams) SetNosqlBackupJobParams(v NoSqlBackupJobParams) {
	o.NosqlBackupJobParams = &v
}

// GetO365BackupParams returns the O365BackupParams field value if set, zero value otherwise.
func (o *EnvBackupParams) GetO365BackupParams() O365BackupEnvParams {
	if o == nil || o.O365BackupParams == nil {
		var ret O365BackupEnvParams
		return ret
	}
	return *o.O365BackupParams
}

// GetO365BackupParamsOk returns a tuple with the O365BackupParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvBackupParams) GetO365BackupParamsOk() (*O365BackupEnvParams, bool) {
	if o == nil || o.O365BackupParams == nil {
		return nil, false
	}
	return o.O365BackupParams, true
}

// HasO365BackupParams returns a boolean if a field has been set.
func (o *EnvBackupParams) HasO365BackupParams() bool {
	if o != nil && o.O365BackupParams != nil {
		return true
	}

	return false
}

// SetO365BackupParams gets a reference to the given O365BackupEnvParams and assigns it to the O365BackupParams field.
func (o *EnvBackupParams) SetO365BackupParams(v O365BackupEnvParams) {
	o.O365BackupParams = &v
}

// GetOracleBackupJobParams returns the OracleBackupJobParams field value if set, zero value otherwise.
func (o *EnvBackupParams) GetOracleBackupJobParams() OracleBackupJobParams {
	if o == nil || o.OracleBackupJobParams == nil {
		var ret OracleBackupJobParams
		return ret
	}
	return *o.OracleBackupJobParams
}

// GetOracleBackupJobParamsOk returns a tuple with the OracleBackupJobParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvBackupParams) GetOracleBackupJobParamsOk() (*OracleBackupJobParams, bool) {
	if o == nil || o.OracleBackupJobParams == nil {
		return nil, false
	}
	return o.OracleBackupJobParams, true
}

// HasOracleBackupJobParams returns a boolean if a field has been set.
func (o *EnvBackupParams) HasOracleBackupJobParams() bool {
	if o != nil && o.OracleBackupJobParams != nil {
		return true
	}

	return false
}

// SetOracleBackupJobParams gets a reference to the given OracleBackupJobParams and assigns it to the OracleBackupJobParams field.
func (o *EnvBackupParams) SetOracleBackupJobParams(v OracleBackupJobParams) {
	o.OracleBackupJobParams = &v
}

// GetOutlookBackupParams returns the OutlookBackupParams field value if set, zero value otherwise.
func (o *EnvBackupParams) GetOutlookBackupParams() OutlookBackupEnvParams {
	if o == nil || o.OutlookBackupParams == nil {
		var ret OutlookBackupEnvParams
		return ret
	}
	return *o.OutlookBackupParams
}

// GetOutlookBackupParamsOk returns a tuple with the OutlookBackupParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvBackupParams) GetOutlookBackupParamsOk() (*OutlookBackupEnvParams, bool) {
	if o == nil || o.OutlookBackupParams == nil {
		return nil, false
	}
	return o.OutlookBackupParams, true
}

// HasOutlookBackupParams returns a boolean if a field has been set.
func (o *EnvBackupParams) HasOutlookBackupParams() bool {
	if o != nil && o.OutlookBackupParams != nil {
		return true
	}

	return false
}

// SetOutlookBackupParams gets a reference to the given OutlookBackupEnvParams and assigns it to the OutlookBackupParams field.
func (o *EnvBackupParams) SetOutlookBackupParams(v OutlookBackupEnvParams) {
	o.OutlookBackupParams = &v
}

// GetPhysicalBackupParams returns the PhysicalBackupParams field value if set, zero value otherwise.
func (o *EnvBackupParams) GetPhysicalBackupParams() PhysicalBackupEnvParams {
	if o == nil || o.PhysicalBackupParams == nil {
		var ret PhysicalBackupEnvParams
		return ret
	}
	return *o.PhysicalBackupParams
}

// GetPhysicalBackupParamsOk returns a tuple with the PhysicalBackupParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvBackupParams) GetPhysicalBackupParamsOk() (*PhysicalBackupEnvParams, bool) {
	if o == nil || o.PhysicalBackupParams == nil {
		return nil, false
	}
	return o.PhysicalBackupParams, true
}

// HasPhysicalBackupParams returns a boolean if a field has been set.
func (o *EnvBackupParams) HasPhysicalBackupParams() bool {
	if o != nil && o.PhysicalBackupParams != nil {
		return true
	}

	return false
}

// SetPhysicalBackupParams gets a reference to the given PhysicalBackupEnvParams and assigns it to the PhysicalBackupParams field.
func (o *EnvBackupParams) SetPhysicalBackupParams(v PhysicalBackupEnvParams) {
	o.PhysicalBackupParams = &v
}

// GetSnapshotManagerParams returns the SnapshotManagerParams field value if set, zero value otherwise.
func (o *EnvBackupParams) GetSnapshotManagerParams() SnapshotManagerParams {
	if o == nil || o.SnapshotManagerParams == nil {
		var ret SnapshotManagerParams
		return ret
	}
	return *o.SnapshotManagerParams
}

// GetSnapshotManagerParamsOk returns a tuple with the SnapshotManagerParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvBackupParams) GetSnapshotManagerParamsOk() (*SnapshotManagerParams, bool) {
	if o == nil || o.SnapshotManagerParams == nil {
		return nil, false
	}
	return o.SnapshotManagerParams, true
}

// HasSnapshotManagerParams returns a boolean if a field has been set.
func (o *EnvBackupParams) HasSnapshotManagerParams() bool {
	if o != nil && o.SnapshotManagerParams != nil {
		return true
	}

	return false
}

// SetSnapshotManagerParams gets a reference to the given SnapshotManagerParams and assigns it to the SnapshotManagerParams field.
func (o *EnvBackupParams) SetSnapshotManagerParams(v SnapshotManagerParams) {
	o.SnapshotManagerParams = &v
}

// GetSqlBackupJobParams returns the SqlBackupJobParams field value if set, zero value otherwise.
func (o *EnvBackupParams) GetSqlBackupJobParams() SqlBackupJobParams {
	if o == nil || o.SqlBackupJobParams == nil {
		var ret SqlBackupJobParams
		return ret
	}
	return *o.SqlBackupJobParams
}

// GetSqlBackupJobParamsOk returns a tuple with the SqlBackupJobParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvBackupParams) GetSqlBackupJobParamsOk() (*SqlBackupJobParams, bool) {
	if o == nil || o.SqlBackupJobParams == nil {
		return nil, false
	}
	return o.SqlBackupJobParams, true
}

// HasSqlBackupJobParams returns a boolean if a field has been set.
func (o *EnvBackupParams) HasSqlBackupJobParams() bool {
	if o != nil && o.SqlBackupJobParams != nil {
		return true
	}

	return false
}

// SetSqlBackupJobParams gets a reference to the given SqlBackupJobParams and assigns it to the SqlBackupJobParams field.
func (o *EnvBackupParams) SetSqlBackupJobParams(v SqlBackupJobParams) {
	o.SqlBackupJobParams = &v
}

// GetVmwareBackupParams returns the VmwareBackupParams field value if set, zero value otherwise.
func (o *EnvBackupParams) GetVmwareBackupParams() VMwareBackupEnvParams {
	if o == nil || o.VmwareBackupParams == nil {
		var ret VMwareBackupEnvParams
		return ret
	}
	return *o.VmwareBackupParams
}

// GetVmwareBackupParamsOk returns a tuple with the VmwareBackupParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvBackupParams) GetVmwareBackupParamsOk() (*VMwareBackupEnvParams, bool) {
	if o == nil || o.VmwareBackupParams == nil {
		return nil, false
	}
	return o.VmwareBackupParams, true
}

// HasVmwareBackupParams returns a boolean if a field has been set.
func (o *EnvBackupParams) HasVmwareBackupParams() bool {
	if o != nil && o.VmwareBackupParams != nil {
		return true
	}

	return false
}

// SetVmwareBackupParams gets a reference to the given VMwareBackupEnvParams and assigns it to the VmwareBackupParams field.
func (o *EnvBackupParams) SetVmwareBackupParams(v VMwareBackupEnvParams) {
	o.VmwareBackupParams = &v
}

func (o EnvBackupParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExchangeBackupJobParams != nil {
		toSerialize["exchangeBackupJobParams"] = o.ExchangeBackupJobParams
	}
	if o.FileStubbingParams != nil {
		toSerialize["fileStubbingParams"] = o.FileStubbingParams
	}
	if o.FileUptieringParams != nil {
		toSerialize["fileUptieringParams"] = o.FileUptieringParams
	}
	if o.HypervBackupParams != nil {
		toSerialize["hypervBackupParams"] = o.HypervBackupParams
	}
	if o.NasBackupParams != nil {
		toSerialize["nasBackupParams"] = o.NasBackupParams
	}
	if o.NosqlBackupJobParams != nil {
		toSerialize["nosqlBackupJobParams"] = o.NosqlBackupJobParams
	}
	if o.O365BackupParams != nil {
		toSerialize["o365BackupParams"] = o.O365BackupParams
	}
	if o.OracleBackupJobParams != nil {
		toSerialize["oracleBackupJobParams"] = o.OracleBackupJobParams
	}
	if o.OutlookBackupParams != nil {
		toSerialize["outlookBackupParams"] = o.OutlookBackupParams
	}
	if o.PhysicalBackupParams != nil {
		toSerialize["physicalBackupParams"] = o.PhysicalBackupParams
	}
	if o.SnapshotManagerParams != nil {
		toSerialize["snapshotManagerParams"] = o.SnapshotManagerParams
	}
	if o.SqlBackupJobParams != nil {
		toSerialize["sqlBackupJobParams"] = o.SqlBackupJobParams
	}
	if o.VmwareBackupParams != nil {
		toSerialize["vmwareBackupParams"] = o.VmwareBackupParams
	}
	return json.Marshal(toSerialize)
}

type NullableEnvBackupParams struct {
	value *EnvBackupParams
	isSet bool
}

func (v NullableEnvBackupParams) Get() *EnvBackupParams {
	return v.value
}

func (v *NullableEnvBackupParams) Set(val *EnvBackupParams) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvBackupParams) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvBackupParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvBackupParams(val *EnvBackupParams) *NullableEnvBackupParams {
	return &NullableEnvBackupParams{value: val, isSet: true}
}

func (v NullableEnvBackupParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvBackupParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


