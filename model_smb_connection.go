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

// SmbConnection struct for SmbConnection
type SmbConnection struct {
	// Specifies the Client IP address of the connection.
	ClientIp NullableString `json:"clientIp,omitempty"`
	// Domain name of the corresponding user.
	DomainName NullableString `json:"domainName,omitempty"`
	// Specifies a Node IP address where the connection request is received.
	NodeIp NullableString `json:"nodeIp,omitempty"`
	// Mount path.
	Path NullableString `json:"path,omitempty"`
	// Specifies the Server IP address of the connection. This could be a VIP, VLAN IP, or node IP on the Cluster.
	ServerIp NullableString `json:"serverIp,omitempty"`
	// Session id.
	SessionId NullableInt64 `json:"sessionId,omitempty"`
	// List of SIDs in the SMB session token.
	Sids []string `json:"sids,omitempty"`
	// User name used to login for this session.
	UserName NullableString `json:"userName,omitempty"`
	// Specifies the id of the view.
	ViewId NullableInt64 `json:"viewId,omitempty"`
	// Specifies the name of the view.
	ViewName NullableString `json:"viewName,omitempty"`
}

// NewSmbConnection instantiates a new SmbConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmbConnection() *SmbConnection {
	this := SmbConnection{}
	return &this
}

// NewSmbConnectionWithDefaults instantiates a new SmbConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmbConnectionWithDefaults() *SmbConnection {
	this := SmbConnection{}
	return &this
}

// GetClientIp returns the ClientIp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmbConnection) GetClientIp() string {
	if o == nil || o.ClientIp.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClientIp.Get()
}

// GetClientIpOk returns a tuple with the ClientIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmbConnection) GetClientIpOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ClientIp.Get(), o.ClientIp.IsSet()
}

// HasClientIp returns a boolean if a field has been set.
func (o *SmbConnection) HasClientIp() bool {
	if o != nil && o.ClientIp.IsSet() {
		return true
	}

	return false
}

// SetClientIp gets a reference to the given NullableString and assigns it to the ClientIp field.
func (o *SmbConnection) SetClientIp(v string) {
	o.ClientIp.Set(&v)
}
// SetClientIpNil sets the value for ClientIp to be an explicit nil
func (o *SmbConnection) SetClientIpNil() {
	o.ClientIp.Set(nil)
}

// UnsetClientIp ensures that no value is present for ClientIp, not even an explicit nil
func (o *SmbConnection) UnsetClientIp() {
	o.ClientIp.Unset()
}

// GetDomainName returns the DomainName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmbConnection) GetDomainName() string {
	if o == nil || o.DomainName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DomainName.Get()
}

// GetDomainNameOk returns a tuple with the DomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmbConnection) GetDomainNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DomainName.Get(), o.DomainName.IsSet()
}

// HasDomainName returns a boolean if a field has been set.
func (o *SmbConnection) HasDomainName() bool {
	if o != nil && o.DomainName.IsSet() {
		return true
	}

	return false
}

// SetDomainName gets a reference to the given NullableString and assigns it to the DomainName field.
func (o *SmbConnection) SetDomainName(v string) {
	o.DomainName.Set(&v)
}
// SetDomainNameNil sets the value for DomainName to be an explicit nil
func (o *SmbConnection) SetDomainNameNil() {
	o.DomainName.Set(nil)
}

// UnsetDomainName ensures that no value is present for DomainName, not even an explicit nil
func (o *SmbConnection) UnsetDomainName() {
	o.DomainName.Unset()
}

// GetNodeIp returns the NodeIp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmbConnection) GetNodeIp() string {
	if o == nil || o.NodeIp.Get() == nil {
		var ret string
		return ret
	}
	return *o.NodeIp.Get()
}

// GetNodeIpOk returns a tuple with the NodeIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmbConnection) GetNodeIpOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NodeIp.Get(), o.NodeIp.IsSet()
}

// HasNodeIp returns a boolean if a field has been set.
func (o *SmbConnection) HasNodeIp() bool {
	if o != nil && o.NodeIp.IsSet() {
		return true
	}

	return false
}

// SetNodeIp gets a reference to the given NullableString and assigns it to the NodeIp field.
func (o *SmbConnection) SetNodeIp(v string) {
	o.NodeIp.Set(&v)
}
// SetNodeIpNil sets the value for NodeIp to be an explicit nil
func (o *SmbConnection) SetNodeIpNil() {
	o.NodeIp.Set(nil)
}

// UnsetNodeIp ensures that no value is present for NodeIp, not even an explicit nil
func (o *SmbConnection) UnsetNodeIp() {
	o.NodeIp.Unset()
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmbConnection) GetPath() string {
	if o == nil || o.Path.Get() == nil {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmbConnection) GetPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *SmbConnection) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *SmbConnection) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *SmbConnection) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *SmbConnection) UnsetPath() {
	o.Path.Unset()
}

// GetServerIp returns the ServerIp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmbConnection) GetServerIp() string {
	if o == nil || o.ServerIp.Get() == nil {
		var ret string
		return ret
	}
	return *o.ServerIp.Get()
}

// GetServerIpOk returns a tuple with the ServerIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmbConnection) GetServerIpOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ServerIp.Get(), o.ServerIp.IsSet()
}

// HasServerIp returns a boolean if a field has been set.
func (o *SmbConnection) HasServerIp() bool {
	if o != nil && o.ServerIp.IsSet() {
		return true
	}

	return false
}

// SetServerIp gets a reference to the given NullableString and assigns it to the ServerIp field.
func (o *SmbConnection) SetServerIp(v string) {
	o.ServerIp.Set(&v)
}
// SetServerIpNil sets the value for ServerIp to be an explicit nil
func (o *SmbConnection) SetServerIpNil() {
	o.ServerIp.Set(nil)
}

// UnsetServerIp ensures that no value is present for ServerIp, not even an explicit nil
func (o *SmbConnection) UnsetServerIp() {
	o.ServerIp.Unset()
}

// GetSessionId returns the SessionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmbConnection) GetSessionId() int64 {
	if o == nil || o.SessionId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.SessionId.Get()
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmbConnection) GetSessionIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SessionId.Get(), o.SessionId.IsSet()
}

// HasSessionId returns a boolean if a field has been set.
func (o *SmbConnection) HasSessionId() bool {
	if o != nil && o.SessionId.IsSet() {
		return true
	}

	return false
}

// SetSessionId gets a reference to the given NullableInt64 and assigns it to the SessionId field.
func (o *SmbConnection) SetSessionId(v int64) {
	o.SessionId.Set(&v)
}
// SetSessionIdNil sets the value for SessionId to be an explicit nil
func (o *SmbConnection) SetSessionIdNil() {
	o.SessionId.Set(nil)
}

// UnsetSessionId ensures that no value is present for SessionId, not even an explicit nil
func (o *SmbConnection) UnsetSessionId() {
	o.SessionId.Unset()
}

// GetSids returns the Sids field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmbConnection) GetSids() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.Sids
}

// GetSidsOk returns a tuple with the Sids field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmbConnection) GetSidsOk() (*[]string, bool) {
	if o == nil || o.Sids == nil {
		return nil, false
	}
	return &o.Sids, true
}

// HasSids returns a boolean if a field has been set.
func (o *SmbConnection) HasSids() bool {
	if o != nil && o.Sids != nil {
		return true
	}

	return false
}

// SetSids gets a reference to the given []string and assigns it to the Sids field.
func (o *SmbConnection) SetSids(v []string) {
	o.Sids = v
}

// GetUserName returns the UserName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmbConnection) GetUserName() string {
	if o == nil || o.UserName.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserName.Get()
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmbConnection) GetUserNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserName.Get(), o.UserName.IsSet()
}

// HasUserName returns a boolean if a field has been set.
func (o *SmbConnection) HasUserName() bool {
	if o != nil && o.UserName.IsSet() {
		return true
	}

	return false
}

// SetUserName gets a reference to the given NullableString and assigns it to the UserName field.
func (o *SmbConnection) SetUserName(v string) {
	o.UserName.Set(&v)
}
// SetUserNameNil sets the value for UserName to be an explicit nil
func (o *SmbConnection) SetUserNameNil() {
	o.UserName.Set(nil)
}

// UnsetUserName ensures that no value is present for UserName, not even an explicit nil
func (o *SmbConnection) UnsetUserName() {
	o.UserName.Unset()
}

// GetViewId returns the ViewId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmbConnection) GetViewId() int64 {
	if o == nil || o.ViewId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ViewId.Get()
}

// GetViewIdOk returns a tuple with the ViewId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmbConnection) GetViewIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ViewId.Get(), o.ViewId.IsSet()
}

// HasViewId returns a boolean if a field has been set.
func (o *SmbConnection) HasViewId() bool {
	if o != nil && o.ViewId.IsSet() {
		return true
	}

	return false
}

// SetViewId gets a reference to the given NullableInt64 and assigns it to the ViewId field.
func (o *SmbConnection) SetViewId(v int64) {
	o.ViewId.Set(&v)
}
// SetViewIdNil sets the value for ViewId to be an explicit nil
func (o *SmbConnection) SetViewIdNil() {
	o.ViewId.Set(nil)
}

// UnsetViewId ensures that no value is present for ViewId, not even an explicit nil
func (o *SmbConnection) UnsetViewId() {
	o.ViewId.Unset()
}

// GetViewName returns the ViewName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmbConnection) GetViewName() string {
	if o == nil || o.ViewName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ViewName.Get()
}

// GetViewNameOk returns a tuple with the ViewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmbConnection) GetViewNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ViewName.Get(), o.ViewName.IsSet()
}

// HasViewName returns a boolean if a field has been set.
func (o *SmbConnection) HasViewName() bool {
	if o != nil && o.ViewName.IsSet() {
		return true
	}

	return false
}

// SetViewName gets a reference to the given NullableString and assigns it to the ViewName field.
func (o *SmbConnection) SetViewName(v string) {
	o.ViewName.Set(&v)
}
// SetViewNameNil sets the value for ViewName to be an explicit nil
func (o *SmbConnection) SetViewNameNil() {
	o.ViewName.Set(nil)
}

// UnsetViewName ensures that no value is present for ViewName, not even an explicit nil
func (o *SmbConnection) UnsetViewName() {
	o.ViewName.Unset()
}

func (o SmbConnection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientIp.IsSet() {
		toSerialize["clientIp"] = o.ClientIp.Get()
	}
	if o.DomainName.IsSet() {
		toSerialize["domainName"] = o.DomainName.Get()
	}
	if o.NodeIp.IsSet() {
		toSerialize["nodeIp"] = o.NodeIp.Get()
	}
	if o.Path.IsSet() {
		toSerialize["path"] = o.Path.Get()
	}
	if o.ServerIp.IsSet() {
		toSerialize["serverIp"] = o.ServerIp.Get()
	}
	if o.SessionId.IsSet() {
		toSerialize["sessionId"] = o.SessionId.Get()
	}
	if o.Sids != nil {
		toSerialize["sids"] = o.Sids
	}
	if o.UserName.IsSet() {
		toSerialize["userName"] = o.UserName.Get()
	}
	if o.ViewId.IsSet() {
		toSerialize["viewId"] = o.ViewId.Get()
	}
	if o.ViewName.IsSet() {
		toSerialize["viewName"] = o.ViewName.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSmbConnection struct {
	value *SmbConnection
	isSet bool
}

func (v NullableSmbConnection) Get() *SmbConnection {
	return v.value
}

func (v *NullableSmbConnection) Set(val *SmbConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableSmbConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableSmbConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmbConnection(val *SmbConnection) *NullableSmbConnection {
	return &NullableSmbConnection{value: val, isSet: true}
}

func (v NullableSmbConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmbConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


