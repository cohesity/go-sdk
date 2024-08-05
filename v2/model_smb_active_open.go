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

// checks if the SmbActiveOpen type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmbActiveOpen{}

// SmbActiveOpen Specifies an active open of an SMB file, its access and sharing information.
type SmbActiveOpen struct {
	// Specifies the File Access Type. Following documentation was taken from MSDN. https://msdn.microsoft.com/en-us/library/Cc246802.aspx  'FileReadData' indicates the right to read data from the file or named   pipe. 'FileWriteData' indicates the right to write data into the file or named   pipe beyond the end of the file. 'FileAppendData' indicates the right to append data into the file or named   pipe. 'FileReadEa' indicates the right to read the extended attributes of the   file or named pipe. 'FileWriteEa' indicates the right to write or change the extended   attributes to the file or named pipe. 'FileExecute' indicates the right to delete entries within a directory. 'FileDeleteChild' indicates the right to execute the file. 'FileReadAttributes' indicates the right to read the attributes of the   file. 'FileWriteAttributes' indicates the right to change the attributes of the   file. 'Delete' indicates the right to delete the file. 'ReadControl' indicates the right to read the security descriptor for the   file or named pipe. 'WriteDac' indicates the right to change the discretionary access control   list (DACL) in the security descriptor for the file or named pipe. For   the DACL data structure, see ACL in [MS-DTYP]. 'WriteOwner' indicates the right to change the owner in the security   descriptor for the file or named pipe. 'Synchronize' is used only by SMB2 clients. 'AccessSystemSecurity' indicates the right to read or change the system   access control list (SACL) in the security descriptor for the file or   named pipe. For the SACL data structure, see ACL in [MS-DTYP].<42> 'MaximumAllowed' indicates that the client is requesting an open to the   file with the highest level of access the client has on this file.   If no access is granted for the client on this file, the server MUST   fail the open with STATUS_ACCESS_DENIED. 'GenericAll' indicates a request for all the access flags that are   previously listed except MaximumAllowed and AccessSystemSecurity. 'GenericExecute' indicates a request for the following combination of   access flags listed above:   FileReadAttributes| FileExecute| Synchronize| ReadControl. 'GenericWrite' indicates a request for the following combination of   access flags listed above:   FileWriteData| FileAppendData| FileWriteAttributes| FileWriteEa|   Synchronize| ReadControl. 'GenericRead' indicates a request for the following combination of   access flags listed above:   FileReadData| FileReadAttributes| FileReadEa| Synchronize|   ReadControl.
	AccessInfoList []string `json:"accessInfoList,omitempty"`
	// Specifies whether access privilege of others if they're allowed to read/write/delete.
	AccessPrivilege []string `json:"accessPrivilege,omitempty"`
	// Specifies the id of the active open.
	OpenId NullableInt64 `json:"openId,omitempty"`
}

// NewSmbActiveOpen instantiates a new SmbActiveOpen object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmbActiveOpen() *SmbActiveOpen {
	this := SmbActiveOpen{}
	return &this
}

// NewSmbActiveOpenWithDefaults instantiates a new SmbActiveOpen object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmbActiveOpenWithDefaults() *SmbActiveOpen {
	this := SmbActiveOpen{}
	return &this
}

// GetAccessInfoList returns the AccessInfoList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmbActiveOpen) GetAccessInfoList() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AccessInfoList
}

// GetAccessInfoListOk returns a tuple with the AccessInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmbActiveOpen) GetAccessInfoListOk() ([]string, bool) {
	if o == nil || IsNil(o.AccessInfoList) {
		return nil, false
	}
	return o.AccessInfoList, true
}

// HasAccessInfoList returns a boolean if a field has been set.
func (o *SmbActiveOpen) HasAccessInfoList() bool {
	if o != nil && !IsNil(o.AccessInfoList) {
		return true
	}

	return false
}

// SetAccessInfoList gets a reference to the given []string and assigns it to the AccessInfoList field.
func (o *SmbActiveOpen) SetAccessInfoList(v []string) {
	o.AccessInfoList = v
}

// GetAccessPrivilege returns the AccessPrivilege field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmbActiveOpen) GetAccessPrivilege() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AccessPrivilege
}

// GetAccessPrivilegeOk returns a tuple with the AccessPrivilege field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmbActiveOpen) GetAccessPrivilegeOk() ([]string, bool) {
	if o == nil || IsNil(o.AccessPrivilege) {
		return nil, false
	}
	return o.AccessPrivilege, true
}

// HasAccessPrivilege returns a boolean if a field has been set.
func (o *SmbActiveOpen) HasAccessPrivilege() bool {
	if o != nil && !IsNil(o.AccessPrivilege) {
		return true
	}

	return false
}

// SetAccessPrivilege gets a reference to the given []string and assigns it to the AccessPrivilege field.
func (o *SmbActiveOpen) SetAccessPrivilege(v []string) {
	o.AccessPrivilege = v
}

// GetOpenId returns the OpenId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmbActiveOpen) GetOpenId() int64 {
	if o == nil || IsNil(o.OpenId.Get()) {
		var ret int64
		return ret
	}
	return *o.OpenId.Get()
}

// GetOpenIdOk returns a tuple with the OpenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmbActiveOpen) GetOpenIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenId.Get(), o.OpenId.IsSet()
}

// HasOpenId returns a boolean if a field has been set.
func (o *SmbActiveOpen) HasOpenId() bool {
	if o != nil && o.OpenId.IsSet() {
		return true
	}

	return false
}

// SetOpenId gets a reference to the given NullableInt64 and assigns it to the OpenId field.
func (o *SmbActiveOpen) SetOpenId(v int64) {
	o.OpenId.Set(&v)
}
// SetOpenIdNil sets the value for OpenId to be an explicit nil
func (o *SmbActiveOpen) SetOpenIdNil() {
	o.OpenId.Set(nil)
}

// UnsetOpenId ensures that no value is present for OpenId, not even an explicit nil
func (o *SmbActiveOpen) UnsetOpenId() {
	o.OpenId.Unset()
}

func (o SmbActiveOpen) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmbActiveOpen) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessInfoList != nil {
		toSerialize["accessInfoList"] = o.AccessInfoList
	}
	if o.AccessPrivilege != nil {
		toSerialize["accessPrivilege"] = o.AccessPrivilege
	}
	if o.OpenId.IsSet() {
		toSerialize["openId"] = o.OpenId.Get()
	}
	return toSerialize, nil
}

type NullableSmbActiveOpen struct {
	value *SmbActiveOpen
	isSet bool
}

func (v NullableSmbActiveOpen) Get() *SmbActiveOpen {
	return v.value
}

func (v *NullableSmbActiveOpen) Set(val *SmbActiveOpen) {
	v.value = val
	v.isSet = true
}

func (v NullableSmbActiveOpen) IsSet() bool {
	return v.isSet
}

func (v *NullableSmbActiveOpen) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmbActiveOpen(val *SmbActiveOpen) *NullableSmbActiveOpen {
	return &NullableSmbActiveOpen{value: val, isSet: true}
}

func (v NullableSmbActiveOpen) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmbActiveOpen) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


