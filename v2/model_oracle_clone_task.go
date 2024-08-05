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

// checks if the OracleCloneTask type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OracleCloneTask{}

// OracleCloneTask Specifies the information about an Oracle clone task.
type OracleCloneTask struct {
	PostScript *CommonPostBackupScriptParams `json:"postScript,omitempty"`
	PreScript *CommonPreBackupScriptParams `json:"preScript,omitempty"`
	// Specifies the base folder of Oracle installation on the target host.
	BaseFolder NullableString `json:"baseFolder"`
	// Specifies the name of the cloned database.
	DbName NullableString `json:"dbName"`
	// Specifies the home folder for the cloned database.
	HomeFolder NullableString `json:"homeFolder"`
	// Specifies the System Global Area (SGA) for the clone database.
	Sga NullableString `json:"sga,omitempty"`
}

type _OracleCloneTask OracleCloneTask

// NewOracleCloneTask instantiates a new OracleCloneTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOracleCloneTask(baseFolder NullableString, dbName NullableString, homeFolder NullableString) *OracleCloneTask {
	this := OracleCloneTask{}
	this.BaseFolder = baseFolder
	this.DbName = dbName
	this.HomeFolder = homeFolder
	return &this
}

// NewOracleCloneTaskWithDefaults instantiates a new OracleCloneTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOracleCloneTaskWithDefaults() *OracleCloneTask {
	this := OracleCloneTask{}
	return &this
}

// GetPostScript returns the PostScript field value if set, zero value otherwise.
func (o *OracleCloneTask) GetPostScript() CommonPostBackupScriptParams {
	if o == nil || IsNil(o.PostScript) {
		var ret CommonPostBackupScriptParams
		return ret
	}
	return *o.PostScript
}

// GetPostScriptOk returns a tuple with the PostScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleCloneTask) GetPostScriptOk() (*CommonPostBackupScriptParams, bool) {
	if o == nil || IsNil(o.PostScript) {
		return nil, false
	}
	return o.PostScript, true
}

// HasPostScript returns a boolean if a field has been set.
func (o *OracleCloneTask) HasPostScript() bool {
	if o != nil && !IsNil(o.PostScript) {
		return true
	}

	return false
}

// SetPostScript gets a reference to the given CommonPostBackupScriptParams and assigns it to the PostScript field.
func (o *OracleCloneTask) SetPostScript(v CommonPostBackupScriptParams) {
	o.PostScript = &v
}

// GetPreScript returns the PreScript field value if set, zero value otherwise.
func (o *OracleCloneTask) GetPreScript() CommonPreBackupScriptParams {
	if o == nil || IsNil(o.PreScript) {
		var ret CommonPreBackupScriptParams
		return ret
	}
	return *o.PreScript
}

// GetPreScriptOk returns a tuple with the PreScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleCloneTask) GetPreScriptOk() (*CommonPreBackupScriptParams, bool) {
	if o == nil || IsNil(o.PreScript) {
		return nil, false
	}
	return o.PreScript, true
}

// HasPreScript returns a boolean if a field has been set.
func (o *OracleCloneTask) HasPreScript() bool {
	if o != nil && !IsNil(o.PreScript) {
		return true
	}

	return false
}

// SetPreScript gets a reference to the given CommonPreBackupScriptParams and assigns it to the PreScript field.
func (o *OracleCloneTask) SetPreScript(v CommonPreBackupScriptParams) {
	o.PreScript = &v
}

// GetBaseFolder returns the BaseFolder field value
// If the value is explicit nil, the zero value for string will be returned
func (o *OracleCloneTask) GetBaseFolder() string {
	if o == nil || o.BaseFolder.Get() == nil {
		var ret string
		return ret
	}

	return *o.BaseFolder.Get()
}

// GetBaseFolderOk returns a tuple with the BaseFolder field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OracleCloneTask) GetBaseFolderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BaseFolder.Get(), o.BaseFolder.IsSet()
}

// SetBaseFolder sets field value
func (o *OracleCloneTask) SetBaseFolder(v string) {
	o.BaseFolder.Set(&v)
}

// GetDbName returns the DbName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *OracleCloneTask) GetDbName() string {
	if o == nil || o.DbName.Get() == nil {
		var ret string
		return ret
	}

	return *o.DbName.Get()
}

// GetDbNameOk returns a tuple with the DbName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OracleCloneTask) GetDbNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DbName.Get(), o.DbName.IsSet()
}

// SetDbName sets field value
func (o *OracleCloneTask) SetDbName(v string) {
	o.DbName.Set(&v)
}

// GetHomeFolder returns the HomeFolder field value
// If the value is explicit nil, the zero value for string will be returned
func (o *OracleCloneTask) GetHomeFolder() string {
	if o == nil || o.HomeFolder.Get() == nil {
		var ret string
		return ret
	}

	return *o.HomeFolder.Get()
}

// GetHomeFolderOk returns a tuple with the HomeFolder field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OracleCloneTask) GetHomeFolderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HomeFolder.Get(), o.HomeFolder.IsSet()
}

// SetHomeFolder sets field value
func (o *OracleCloneTask) SetHomeFolder(v string) {
	o.HomeFolder.Set(&v)
}

// GetSga returns the Sga field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OracleCloneTask) GetSga() string {
	if o == nil || IsNil(o.Sga.Get()) {
		var ret string
		return ret
	}
	return *o.Sga.Get()
}

// GetSgaOk returns a tuple with the Sga field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OracleCloneTask) GetSgaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sga.Get(), o.Sga.IsSet()
}

// HasSga returns a boolean if a field has been set.
func (o *OracleCloneTask) HasSga() bool {
	if o != nil && o.Sga.IsSet() {
		return true
	}

	return false
}

// SetSga gets a reference to the given NullableString and assigns it to the Sga field.
func (o *OracleCloneTask) SetSga(v string) {
	o.Sga.Set(&v)
}
// SetSgaNil sets the value for Sga to be an explicit nil
func (o *OracleCloneTask) SetSgaNil() {
	o.Sga.Set(nil)
}

// UnsetSga ensures that no value is present for Sga, not even an explicit nil
func (o *OracleCloneTask) UnsetSga() {
	o.Sga.Unset()
}

func (o OracleCloneTask) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OracleCloneTask) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PostScript) {
		toSerialize["postScript"] = o.PostScript
	}
	if !IsNil(o.PreScript) {
		toSerialize["preScript"] = o.PreScript
	}
	toSerialize["baseFolder"] = o.BaseFolder.Get()
	toSerialize["dbName"] = o.DbName.Get()
	toSerialize["homeFolder"] = o.HomeFolder.Get()
	if o.Sga.IsSet() {
		toSerialize["sga"] = o.Sga.Get()
	}
	return toSerialize, nil
}

func (o *OracleCloneTask) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"baseFolder",
		"dbName",
		"homeFolder",
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

	varOracleCloneTask := _OracleCloneTask{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOracleCloneTask)

	if err != nil {
		return err
	}

	*o = OracleCloneTask(varOracleCloneTask)

	return err
}

type NullableOracleCloneTask struct {
	value *OracleCloneTask
	isSet bool
}

func (v NullableOracleCloneTask) Get() *OracleCloneTask {
	return v.value
}

func (v *NullableOracleCloneTask) Set(val *OracleCloneTask) {
	v.value = val
	v.isSet = true
}

func (v NullableOracleCloneTask) IsSet() bool {
	return v.isSet
}

func (v *NullableOracleCloneTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOracleCloneTask(val *OracleCloneTask) *NullableOracleCloneTask {
	return &NullableOracleCloneTask{value: val, isSet: true}
}

func (v NullableOracleCloneTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOracleCloneTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


