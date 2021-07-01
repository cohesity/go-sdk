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

// SqlAagHostAndDatabases Specifies AAGs and databases information on an SQL server. If AAGs exist on the server, specifies information about the AAG and databases in the group for each AAG found on the server.
type SqlAagHostAndDatabases struct {
	// Specifies a list of AAGs and database members in each AAG.
	AagDatabases []AagAndDatabases `json:"aagDatabases,omitempty"`
	ApplicationNode *ProtectionSourceNode `json:"applicationNode,omitempty"`
	// Specifies all database entities found on the server. Database may or may not be in an AAG.
	Databases []ProtectionSource `json:"databases,omitempty"`
	// Specifies an error message when the host is not registered as an SQL host.
	ErrorMessage NullableString `json:"errorMessage,omitempty"`
	// Specifies the name of the host that is not registered as an SQL server on Cohesity Cluser.
	UnknownHostName NullableString `json:"unknownHostName,omitempty"`
}

// NewSqlAagHostAndDatabases instantiates a new SqlAagHostAndDatabases object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSqlAagHostAndDatabases() *SqlAagHostAndDatabases {
	this := SqlAagHostAndDatabases{}
	return &this
}

// NewSqlAagHostAndDatabasesWithDefaults instantiates a new SqlAagHostAndDatabases object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSqlAagHostAndDatabasesWithDefaults() *SqlAagHostAndDatabases {
	this := SqlAagHostAndDatabases{}
	return &this
}

// GetAagDatabases returns the AagDatabases field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SqlAagHostAndDatabases) GetAagDatabases() []AagAndDatabases {
	if o == nil  {
		var ret []AagAndDatabases
		return ret
	}
	return o.AagDatabases
}

// GetAagDatabasesOk returns a tuple with the AagDatabases field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SqlAagHostAndDatabases) GetAagDatabasesOk() (*[]AagAndDatabases, bool) {
	if o == nil || o.AagDatabases == nil {
		return nil, false
	}
	return &o.AagDatabases, true
}

// HasAagDatabases returns a boolean if a field has been set.
func (o *SqlAagHostAndDatabases) HasAagDatabases() bool {
	if o != nil && o.AagDatabases != nil {
		return true
	}

	return false
}

// SetAagDatabases gets a reference to the given []AagAndDatabases and assigns it to the AagDatabases field.
func (o *SqlAagHostAndDatabases) SetAagDatabases(v []AagAndDatabases) {
	o.AagDatabases = v
}

// GetApplicationNode returns the ApplicationNode field value if set, zero value otherwise.
func (o *SqlAagHostAndDatabases) GetApplicationNode() ProtectionSourceNode {
	if o == nil || o.ApplicationNode == nil {
		var ret ProtectionSourceNode
		return ret
	}
	return *o.ApplicationNode
}

// GetApplicationNodeOk returns a tuple with the ApplicationNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SqlAagHostAndDatabases) GetApplicationNodeOk() (*ProtectionSourceNode, bool) {
	if o == nil || o.ApplicationNode == nil {
		return nil, false
	}
	return o.ApplicationNode, true
}

// HasApplicationNode returns a boolean if a field has been set.
func (o *SqlAagHostAndDatabases) HasApplicationNode() bool {
	if o != nil && o.ApplicationNode != nil {
		return true
	}

	return false
}

// SetApplicationNode gets a reference to the given ProtectionSourceNode and assigns it to the ApplicationNode field.
func (o *SqlAagHostAndDatabases) SetApplicationNode(v ProtectionSourceNode) {
	o.ApplicationNode = &v
}

// GetDatabases returns the Databases field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SqlAagHostAndDatabases) GetDatabases() []ProtectionSource {
	if o == nil  {
		var ret []ProtectionSource
		return ret
	}
	return o.Databases
}

// GetDatabasesOk returns a tuple with the Databases field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SqlAagHostAndDatabases) GetDatabasesOk() (*[]ProtectionSource, bool) {
	if o == nil || o.Databases == nil {
		return nil, false
	}
	return &o.Databases, true
}

// HasDatabases returns a boolean if a field has been set.
func (o *SqlAagHostAndDatabases) HasDatabases() bool {
	if o != nil && o.Databases != nil {
		return true
	}

	return false
}

// SetDatabases gets a reference to the given []ProtectionSource and assigns it to the Databases field.
func (o *SqlAagHostAndDatabases) SetDatabases(v []ProtectionSource) {
	o.Databases = v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SqlAagHostAndDatabases) GetErrorMessage() string {
	if o == nil || o.ErrorMessage.Get() == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage.Get()
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SqlAagHostAndDatabases) GetErrorMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ErrorMessage.Get(), o.ErrorMessage.IsSet()
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *SqlAagHostAndDatabases) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage.IsSet() {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given NullableString and assigns it to the ErrorMessage field.
func (o *SqlAagHostAndDatabases) SetErrorMessage(v string) {
	o.ErrorMessage.Set(&v)
}
// SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil
func (o *SqlAagHostAndDatabases) SetErrorMessageNil() {
	o.ErrorMessage.Set(nil)
}

// UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
func (o *SqlAagHostAndDatabases) UnsetErrorMessage() {
	o.ErrorMessage.Unset()
}

// GetUnknownHostName returns the UnknownHostName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SqlAagHostAndDatabases) GetUnknownHostName() string {
	if o == nil || o.UnknownHostName.Get() == nil {
		var ret string
		return ret
	}
	return *o.UnknownHostName.Get()
}

// GetUnknownHostNameOk returns a tuple with the UnknownHostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SqlAagHostAndDatabases) GetUnknownHostNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UnknownHostName.Get(), o.UnknownHostName.IsSet()
}

// HasUnknownHostName returns a boolean if a field has been set.
func (o *SqlAagHostAndDatabases) HasUnknownHostName() bool {
	if o != nil && o.UnknownHostName.IsSet() {
		return true
	}

	return false
}

// SetUnknownHostName gets a reference to the given NullableString and assigns it to the UnknownHostName field.
func (o *SqlAagHostAndDatabases) SetUnknownHostName(v string) {
	o.UnknownHostName.Set(&v)
}
// SetUnknownHostNameNil sets the value for UnknownHostName to be an explicit nil
func (o *SqlAagHostAndDatabases) SetUnknownHostNameNil() {
	o.UnknownHostName.Set(nil)
}

// UnsetUnknownHostName ensures that no value is present for UnknownHostName, not even an explicit nil
func (o *SqlAagHostAndDatabases) UnsetUnknownHostName() {
	o.UnknownHostName.Unset()
}

func (o SqlAagHostAndDatabases) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AagDatabases != nil {
		toSerialize["aagDatabases"] = o.AagDatabases
	}
	if o.ApplicationNode != nil {
		toSerialize["applicationNode"] = o.ApplicationNode
	}
	if o.Databases != nil {
		toSerialize["databases"] = o.Databases
	}
	if o.ErrorMessage.IsSet() {
		toSerialize["errorMessage"] = o.ErrorMessage.Get()
	}
	if o.UnknownHostName.IsSet() {
		toSerialize["unknownHostName"] = o.UnknownHostName.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSqlAagHostAndDatabases struct {
	value *SqlAagHostAndDatabases
	isSet bool
}

func (v NullableSqlAagHostAndDatabases) Get() *SqlAagHostAndDatabases {
	return v.value
}

func (v *NullableSqlAagHostAndDatabases) Set(val *SqlAagHostAndDatabases) {
	v.value = val
	v.isSet = true
}

func (v NullableSqlAagHostAndDatabases) IsSet() bool {
	return v.isSet
}

func (v *NullableSqlAagHostAndDatabases) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSqlAagHostAndDatabases(val *SqlAagHostAndDatabases) *NullableSqlAagHostAndDatabases {
	return &NullableSqlAagHostAndDatabases{value: val, isSet: true}
}

func (v NullableSqlAagHostAndDatabases) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSqlAagHostAndDatabases) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


