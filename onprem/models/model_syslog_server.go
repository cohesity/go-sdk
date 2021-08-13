/*
 * Cohesity REST API
 *
 * Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

// package onprem
package models

import (
	"encoding/json"
	. "github.com/cohesity/go-sdk/onprem/utils"
	"fmt"
)

var _ = NullableBool{}

// SyslogServer Specifies information about syslog server.
type SyslogServer struct {
	// The id of the syslog server.
	Id NullableInt32 `json:"id,omitempty"`
	// Specifies the IP address or hostname of the syslog server.
	Ip NullableString `json:"ip,omitempty"`
	// Specifies the port where the syslog server listens.
	Port NullableInt32 `json:"port,omitempty"`
	// Specifies the protocol used to send the logs.
	Protocol NullableString `json:"protocol,omitempty"`
	// Specifies a unique name for the syslog server on the Cluster.
	Name NullableString `json:"name,omitempty"`
	// Specifies whether to enable the syslog server on the Cluster.
	Enabled NullableBool `json:"enabled,omitempty"`
	// Send enabled syslog facilities related logs to logging server.
	FacilityList *[]string `json:"facilityList,omitempty"`
	// Send programes related logs to logging server.
	ProgramNameList *[]string `json:"programNameList,omitempty"`
	// Send logs including the msg patterns to logging server.
	MsgPatternList *[]string `json:"msgPatternList,omitempty"`
	// Send logs including the msg patterns to logging server.
	RawMsgPatternList *[]string `json:"rawMsgPatternList,omitempty"`
	// Specify whether to enable tls support.
	IsTlsEnabled NullableBool `json:"isTlsEnabled,omitempty"`
	// Syslog server CA certificate.
	CaCertificate NullableString `json:"caCertificate,omitempty"`
}

// NewSyslogServer instantiates a new SyslogServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyslogServer() *SyslogServer {
	this := SyslogServer{}
	return &this
}

// NewSyslogServerWithDefaults instantiates a new SyslogServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyslogServerWithDefaults() *SyslogServer {
	this := SyslogServer{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SyslogServer) GetId() int32 {
	if o == nil || o.Id.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SyslogServer) GetIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *SyslogServer) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *SyslogServer) SetId(v int32) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *SyslogServer) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *SyslogServer) UnsetId() {
	o.Id.Unset()
}

// GetIp returns the Ip field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SyslogServer) GetIp() string {
	if o == nil || o.Ip.Get() == nil {
		var ret string
		return ret
	}
	return *o.Ip.Get()
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SyslogServer) GetIpOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Ip.Get(), o.Ip.IsSet()
}

// HasIp returns a boolean if a field has been set.
func (o *SyslogServer) HasIp() bool {
	if o != nil && o.Ip.IsSet() {
		return true
	}

	return false
}

// SetIp gets a reference to the given NullableString and assigns it to the Ip field.
func (o *SyslogServer) SetIp(v string) {
	o.Ip.Set(&v)
}
// SetIpNil sets the value for Ip to be an explicit nil
func (o *SyslogServer) SetIpNil() {
	o.Ip.Set(nil)
}

// UnsetIp ensures that no value is present for Ip, not even an explicit nil
func (o *SyslogServer) UnsetIp() {
	o.Ip.Unset()
}

// GetPort returns the Port field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SyslogServer) GetPort() int32 {
	if o == nil || o.Port.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Port.Get()
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SyslogServer) GetPortOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Port.Get(), o.Port.IsSet()
}

// HasPort returns a boolean if a field has been set.
func (o *SyslogServer) HasPort() bool {
	if o != nil && o.Port.IsSet() {
		return true
	}

	return false
}

// SetPort gets a reference to the given NullableInt32 and assigns it to the Port field.
func (o *SyslogServer) SetPort(v int32) {
	o.Port.Set(&v)
}
// SetPortNil sets the value for Port to be an explicit nil
func (o *SyslogServer) SetPortNil() {
	o.Port.Set(nil)
}

// UnsetPort ensures that no value is present for Port, not even an explicit nil
func (o *SyslogServer) UnsetPort() {
	o.Port.Unset()
}

// GetProtocol returns the Protocol field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SyslogServer) GetProtocol() string {
	if o == nil || o.Protocol.Get() == nil {
		var ret string
		return ret
	}
	return *o.Protocol.Get()
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SyslogServer) GetProtocolOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Protocol.Get(), o.Protocol.IsSet()
}

// HasProtocol returns a boolean if a field has been set.
func (o *SyslogServer) HasProtocol() bool {
	if o != nil && o.Protocol.IsSet() {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given NullableString and assigns it to the Protocol field.
func (o *SyslogServer) SetProtocol(v string) {
	o.Protocol.Set(&v)
}
// SetProtocolNil sets the value for Protocol to be an explicit nil
func (o *SyslogServer) SetProtocolNil() {
	o.Protocol.Set(nil)
}

// UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
func (o *SyslogServer) UnsetProtocol() {
	o.Protocol.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SyslogServer) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SyslogServer) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *SyslogServer) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *SyslogServer) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *SyslogServer) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *SyslogServer) UnsetName() {
	o.Name.Unset()
}

// GetEnabled returns the Enabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SyslogServer) GetEnabled() bool {
	if o == nil || o.Enabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Enabled.Get()
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SyslogServer) GetEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Enabled.Get(), o.Enabled.IsSet()
}

// HasEnabled returns a boolean if a field has been set.
func (o *SyslogServer) HasEnabled() bool {
	if o != nil && o.Enabled.IsSet() {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given NullableBool and assigns it to the Enabled field.
func (o *SyslogServer) SetEnabled(v bool) {
	o.Enabled.Set(&v)
}
// SetEnabledNil sets the value for Enabled to be an explicit nil
func (o *SyslogServer) SetEnabledNil() {
	o.Enabled.Set(nil)
}

// UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
func (o *SyslogServer) UnsetEnabled() {
	o.Enabled.Unset()
}

// GetFacilityList returns the FacilityList field value if set, zero value otherwise.
func (o *SyslogServer) GetFacilityList() []string {
	if o == nil || o.FacilityList == nil {
		var ret []string
		return ret
	}
	return *o.FacilityList
}

// GetFacilityListOk returns a tuple with the FacilityList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogServer) GetFacilityListOk() (*[]string, bool) {
	if o == nil || o.FacilityList == nil {
		return nil, false
	}
	return o.FacilityList, true
}

// HasFacilityList returns a boolean if a field has been set.
func (o *SyslogServer) HasFacilityList() bool {
	if o != nil && o.FacilityList != nil {
		return true
	}

	return false
}

// SetFacilityList gets a reference to the given []string and assigns it to the FacilityList field.
func (o *SyslogServer) SetFacilityList(v []string) {
	o.FacilityList = &v
}

// GetProgramNameList returns the ProgramNameList field value if set, zero value otherwise.
func (o *SyslogServer) GetProgramNameList() []string {
	if o == nil || o.ProgramNameList == nil {
		var ret []string
		return ret
	}
	return *o.ProgramNameList
}

// GetProgramNameListOk returns a tuple with the ProgramNameList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogServer) GetProgramNameListOk() (*[]string, bool) {
	if o == nil || o.ProgramNameList == nil {
		return nil, false
	}
	return o.ProgramNameList, true
}

// HasProgramNameList returns a boolean if a field has been set.
func (o *SyslogServer) HasProgramNameList() bool {
	if o != nil && o.ProgramNameList != nil {
		return true
	}

	return false
}

// SetProgramNameList gets a reference to the given []string and assigns it to the ProgramNameList field.
func (o *SyslogServer) SetProgramNameList(v []string) {
	o.ProgramNameList = &v
}

// GetMsgPatternList returns the MsgPatternList field value if set, zero value otherwise.
func (o *SyslogServer) GetMsgPatternList() []string {
	if o == nil || o.MsgPatternList == nil {
		var ret []string
		return ret
	}
	return *o.MsgPatternList
}

// GetMsgPatternListOk returns a tuple with the MsgPatternList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogServer) GetMsgPatternListOk() (*[]string, bool) {
	if o == nil || o.MsgPatternList == nil {
		return nil, false
	}
	return o.MsgPatternList, true
}

// HasMsgPatternList returns a boolean if a field has been set.
func (o *SyslogServer) HasMsgPatternList() bool {
	if o != nil && o.MsgPatternList != nil {
		return true
	}

	return false
}

// SetMsgPatternList gets a reference to the given []string and assigns it to the MsgPatternList field.
func (o *SyslogServer) SetMsgPatternList(v []string) {
	o.MsgPatternList = &v
}

// GetRawMsgPatternList returns the RawMsgPatternList field value if set, zero value otherwise.
func (o *SyslogServer) GetRawMsgPatternList() []string {
	if o == nil || o.RawMsgPatternList == nil {
		var ret []string
		return ret
	}
	return *o.RawMsgPatternList
}

// GetRawMsgPatternListOk returns a tuple with the RawMsgPatternList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogServer) GetRawMsgPatternListOk() (*[]string, bool) {
	if o == nil || o.RawMsgPatternList == nil {
		return nil, false
	}
	return o.RawMsgPatternList, true
}

// HasRawMsgPatternList returns a boolean if a field has been set.
func (o *SyslogServer) HasRawMsgPatternList() bool {
	if o != nil && o.RawMsgPatternList != nil {
		return true
	}

	return false
}

// SetRawMsgPatternList gets a reference to the given []string and assigns it to the RawMsgPatternList field.
func (o *SyslogServer) SetRawMsgPatternList(v []string) {
	o.RawMsgPatternList = &v
}

// GetIsTlsEnabled returns the IsTlsEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SyslogServer) GetIsTlsEnabled() bool {
	if o == nil || o.IsTlsEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsTlsEnabled.Get()
}

// GetIsTlsEnabledOk returns a tuple with the IsTlsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SyslogServer) GetIsTlsEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsTlsEnabled.Get(), o.IsTlsEnabled.IsSet()
}

// HasIsTlsEnabled returns a boolean if a field has been set.
func (o *SyslogServer) HasIsTlsEnabled() bool {
	if o != nil && o.IsTlsEnabled.IsSet() {
		return true
	}

	return false
}

// SetIsTlsEnabled gets a reference to the given NullableBool and assigns it to the IsTlsEnabled field.
func (o *SyslogServer) SetIsTlsEnabled(v bool) {
	o.IsTlsEnabled.Set(&v)
}
// SetIsTlsEnabledNil sets the value for IsTlsEnabled to be an explicit nil
func (o *SyslogServer) SetIsTlsEnabledNil() {
	o.IsTlsEnabled.Set(nil)
}

// UnsetIsTlsEnabled ensures that no value is present for IsTlsEnabled, not even an explicit nil
func (o *SyslogServer) UnsetIsTlsEnabled() {
	o.IsTlsEnabled.Unset()
}

// GetCaCertificate returns the CaCertificate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SyslogServer) GetCaCertificate() string {
	if o == nil || o.CaCertificate.Get() == nil {
		var ret string
		return ret
	}
	return *o.CaCertificate.Get()
}

// GetCaCertificateOk returns a tuple with the CaCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SyslogServer) GetCaCertificateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CaCertificate.Get(), o.CaCertificate.IsSet()
}

// HasCaCertificate returns a boolean if a field has been set.
func (o *SyslogServer) HasCaCertificate() bool {
	if o != nil && o.CaCertificate.IsSet() {
		return true
	}

	return false
}

// SetCaCertificate gets a reference to the given NullableString and assigns it to the CaCertificate field.
func (o *SyslogServer) SetCaCertificate(v string) {
	o.CaCertificate.Set(&v)
}
// SetCaCertificateNil sets the value for CaCertificate to be an explicit nil
func (o *SyslogServer) SetCaCertificateNil() {
	o.CaCertificate.Set(nil)
}

// UnsetCaCertificate ensures that no value is present for CaCertificate, not even an explicit nil
func (o *SyslogServer) UnsetCaCertificate() {
	o.CaCertificate.Unset()
}

func (o SyslogServer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Ip.IsSet() {
		toSerialize["ip"] = o.Ip.Get()
	}
	if o.Port.IsSet() {
		toSerialize["port"] = o.Port.Get()
	}
	if o.Protocol.IsSet() {
		toSerialize["protocol"] = o.Protocol.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Enabled.IsSet() {
		toSerialize["enabled"] = o.Enabled.Get()
	}
	if o.FacilityList != nil {
		toSerialize["facilityList"] = o.FacilityList
	}
	if o.ProgramNameList != nil {
		toSerialize["programNameList"] = o.ProgramNameList
	}
	if o.MsgPatternList != nil {
		toSerialize["msgPatternList"] = o.MsgPatternList
	}
	if o.RawMsgPatternList != nil {
		toSerialize["rawMsgPatternList"] = o.RawMsgPatternList
	}
	if o.IsTlsEnabled.IsSet() {
		toSerialize["isTlsEnabled"] = o.IsTlsEnabled.Get()
	}
	if o.CaCertificate.IsSet() {
		toSerialize["caCertificate"] = o.CaCertificate.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSyslogServer struct {
	value *SyslogServer
	isSet bool
}

func (v NullableSyslogServer) Get() *SyslogServer {
	return v.value
}

func (v *NullableSyslogServer) Set(val *SyslogServer) {
	v.value = val
	v.isSet = true
}

func (v NullableSyslogServer) IsSet() bool {
	return v.isSet
}

func (v *NullableSyslogServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyslogServer(val *SyslogServer) *NullableSyslogServer {
	return &NullableSyslogServer{value: val, isSet: true}
}

func (v NullableSyslogServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyslogServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o SyslogServer) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}