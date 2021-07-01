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

// ApiKey Specifies the parameters of an API key.
type ApiKey struct {
	// Specifies the API key created time in milli seconds.
	CreatedTimeMsecs NullableInt64 `json:"createdTimeMsecs,omitempty"`
	// Specifies the user sid who created this API key.
	CreatedUserSid NullableString `json:"createdUserSid,omitempty"`
	// Specifies the username who created this API key.
	CreatedUsername NullableString `json:"createdUsername,omitempty"`
	// Specifies a time stamp when the API key will expire in milli seconds.
	ExpiringTimeMsecs NullableInt64 `json:"expiringTimeMsecs,omitempty"`
	// Specifies the API key id.
	Id NullableString `json:"id,omitempty"`
	// Specifies if the API key is active. Only an active and not expired API key can be used for authentication.
	IsActive NullableBool `json:"isActive,omitempty"`
	// Specifies if the API key is expired. Expired API keys cannot be used for authentication.
	IsExpired NullableBool `json:"isExpired,omitempty"`
	// Specifies the API key name.
	Name NullableString `json:"name,omitempty"`
	// Specifies the user sid who owns this API key.
	OwnerUserSid NullableString `json:"ownerUserSid,omitempty"`
	// Specifies the username who owns this API key.
	OwnerUsername NullableString `json:"ownerUsername,omitempty"`
}

// NewApiKey instantiates a new ApiKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiKey() *ApiKey {
	this := ApiKey{}
	return &this
}

// NewApiKeyWithDefaults instantiates a new ApiKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiKeyWithDefaults() *ApiKey {
	this := ApiKey{}
	return &this
}

// GetCreatedTimeMsecs returns the CreatedTimeMsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiKey) GetCreatedTimeMsecs() int64 {
	if o == nil || o.CreatedTimeMsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CreatedTimeMsecs.Get()
}

// GetCreatedTimeMsecsOk returns a tuple with the CreatedTimeMsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiKey) GetCreatedTimeMsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedTimeMsecs.Get(), o.CreatedTimeMsecs.IsSet()
}

// HasCreatedTimeMsecs returns a boolean if a field has been set.
func (o *ApiKey) HasCreatedTimeMsecs() bool {
	if o != nil && o.CreatedTimeMsecs.IsSet() {
		return true
	}

	return false
}

// SetCreatedTimeMsecs gets a reference to the given NullableInt64 and assigns it to the CreatedTimeMsecs field.
func (o *ApiKey) SetCreatedTimeMsecs(v int64) {
	o.CreatedTimeMsecs.Set(&v)
}
// SetCreatedTimeMsecsNil sets the value for CreatedTimeMsecs to be an explicit nil
func (o *ApiKey) SetCreatedTimeMsecsNil() {
	o.CreatedTimeMsecs.Set(nil)
}

// UnsetCreatedTimeMsecs ensures that no value is present for CreatedTimeMsecs, not even an explicit nil
func (o *ApiKey) UnsetCreatedTimeMsecs() {
	o.CreatedTimeMsecs.Unset()
}

// GetCreatedUserSid returns the CreatedUserSid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiKey) GetCreatedUserSid() string {
	if o == nil || o.CreatedUserSid.Get() == nil {
		var ret string
		return ret
	}
	return *o.CreatedUserSid.Get()
}

// GetCreatedUserSidOk returns a tuple with the CreatedUserSid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiKey) GetCreatedUserSidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedUserSid.Get(), o.CreatedUserSid.IsSet()
}

// HasCreatedUserSid returns a boolean if a field has been set.
func (o *ApiKey) HasCreatedUserSid() bool {
	if o != nil && o.CreatedUserSid.IsSet() {
		return true
	}

	return false
}

// SetCreatedUserSid gets a reference to the given NullableString and assigns it to the CreatedUserSid field.
func (o *ApiKey) SetCreatedUserSid(v string) {
	o.CreatedUserSid.Set(&v)
}
// SetCreatedUserSidNil sets the value for CreatedUserSid to be an explicit nil
func (o *ApiKey) SetCreatedUserSidNil() {
	o.CreatedUserSid.Set(nil)
}

// UnsetCreatedUserSid ensures that no value is present for CreatedUserSid, not even an explicit nil
func (o *ApiKey) UnsetCreatedUserSid() {
	o.CreatedUserSid.Unset()
}

// GetCreatedUsername returns the CreatedUsername field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiKey) GetCreatedUsername() string {
	if o == nil || o.CreatedUsername.Get() == nil {
		var ret string
		return ret
	}
	return *o.CreatedUsername.Get()
}

// GetCreatedUsernameOk returns a tuple with the CreatedUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiKey) GetCreatedUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedUsername.Get(), o.CreatedUsername.IsSet()
}

// HasCreatedUsername returns a boolean if a field has been set.
func (o *ApiKey) HasCreatedUsername() bool {
	if o != nil && o.CreatedUsername.IsSet() {
		return true
	}

	return false
}

// SetCreatedUsername gets a reference to the given NullableString and assigns it to the CreatedUsername field.
func (o *ApiKey) SetCreatedUsername(v string) {
	o.CreatedUsername.Set(&v)
}
// SetCreatedUsernameNil sets the value for CreatedUsername to be an explicit nil
func (o *ApiKey) SetCreatedUsernameNil() {
	o.CreatedUsername.Set(nil)
}

// UnsetCreatedUsername ensures that no value is present for CreatedUsername, not even an explicit nil
func (o *ApiKey) UnsetCreatedUsername() {
	o.CreatedUsername.Unset()
}

// GetExpiringTimeMsecs returns the ExpiringTimeMsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiKey) GetExpiringTimeMsecs() int64 {
	if o == nil || o.ExpiringTimeMsecs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ExpiringTimeMsecs.Get()
}

// GetExpiringTimeMsecsOk returns a tuple with the ExpiringTimeMsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiKey) GetExpiringTimeMsecsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExpiringTimeMsecs.Get(), o.ExpiringTimeMsecs.IsSet()
}

// HasExpiringTimeMsecs returns a boolean if a field has been set.
func (o *ApiKey) HasExpiringTimeMsecs() bool {
	if o != nil && o.ExpiringTimeMsecs.IsSet() {
		return true
	}

	return false
}

// SetExpiringTimeMsecs gets a reference to the given NullableInt64 and assigns it to the ExpiringTimeMsecs field.
func (o *ApiKey) SetExpiringTimeMsecs(v int64) {
	o.ExpiringTimeMsecs.Set(&v)
}
// SetExpiringTimeMsecsNil sets the value for ExpiringTimeMsecs to be an explicit nil
func (o *ApiKey) SetExpiringTimeMsecsNil() {
	o.ExpiringTimeMsecs.Set(nil)
}

// UnsetExpiringTimeMsecs ensures that no value is present for ExpiringTimeMsecs, not even an explicit nil
func (o *ApiKey) UnsetExpiringTimeMsecs() {
	o.ExpiringTimeMsecs.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiKey) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiKey) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *ApiKey) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *ApiKey) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *ApiKey) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *ApiKey) UnsetId() {
	o.Id.Unset()
}

// GetIsActive returns the IsActive field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiKey) GetIsActive() bool {
	if o == nil || o.IsActive.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsActive.Get()
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiKey) GetIsActiveOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsActive.Get(), o.IsActive.IsSet()
}

// HasIsActive returns a boolean if a field has been set.
func (o *ApiKey) HasIsActive() bool {
	if o != nil && o.IsActive.IsSet() {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given NullableBool and assigns it to the IsActive field.
func (o *ApiKey) SetIsActive(v bool) {
	o.IsActive.Set(&v)
}
// SetIsActiveNil sets the value for IsActive to be an explicit nil
func (o *ApiKey) SetIsActiveNil() {
	o.IsActive.Set(nil)
}

// UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
func (o *ApiKey) UnsetIsActive() {
	o.IsActive.Unset()
}

// GetIsExpired returns the IsExpired field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiKey) GetIsExpired() bool {
	if o == nil || o.IsExpired.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsExpired.Get()
}

// GetIsExpiredOk returns a tuple with the IsExpired field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiKey) GetIsExpiredOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsExpired.Get(), o.IsExpired.IsSet()
}

// HasIsExpired returns a boolean if a field has been set.
func (o *ApiKey) HasIsExpired() bool {
	if o != nil && o.IsExpired.IsSet() {
		return true
	}

	return false
}

// SetIsExpired gets a reference to the given NullableBool and assigns it to the IsExpired field.
func (o *ApiKey) SetIsExpired(v bool) {
	o.IsExpired.Set(&v)
}
// SetIsExpiredNil sets the value for IsExpired to be an explicit nil
func (o *ApiKey) SetIsExpiredNil() {
	o.IsExpired.Set(nil)
}

// UnsetIsExpired ensures that no value is present for IsExpired, not even an explicit nil
func (o *ApiKey) UnsetIsExpired() {
	o.IsExpired.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiKey) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiKey) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ApiKey) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ApiKey) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ApiKey) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ApiKey) UnsetName() {
	o.Name.Unset()
}

// GetOwnerUserSid returns the OwnerUserSid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiKey) GetOwnerUserSid() string {
	if o == nil || o.OwnerUserSid.Get() == nil {
		var ret string
		return ret
	}
	return *o.OwnerUserSid.Get()
}

// GetOwnerUserSidOk returns a tuple with the OwnerUserSid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiKey) GetOwnerUserSidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OwnerUserSid.Get(), o.OwnerUserSid.IsSet()
}

// HasOwnerUserSid returns a boolean if a field has been set.
func (o *ApiKey) HasOwnerUserSid() bool {
	if o != nil && o.OwnerUserSid.IsSet() {
		return true
	}

	return false
}

// SetOwnerUserSid gets a reference to the given NullableString and assigns it to the OwnerUserSid field.
func (o *ApiKey) SetOwnerUserSid(v string) {
	o.OwnerUserSid.Set(&v)
}
// SetOwnerUserSidNil sets the value for OwnerUserSid to be an explicit nil
func (o *ApiKey) SetOwnerUserSidNil() {
	o.OwnerUserSid.Set(nil)
}

// UnsetOwnerUserSid ensures that no value is present for OwnerUserSid, not even an explicit nil
func (o *ApiKey) UnsetOwnerUserSid() {
	o.OwnerUserSid.Unset()
}

// GetOwnerUsername returns the OwnerUsername field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiKey) GetOwnerUsername() string {
	if o == nil || o.OwnerUsername.Get() == nil {
		var ret string
		return ret
	}
	return *o.OwnerUsername.Get()
}

// GetOwnerUsernameOk returns a tuple with the OwnerUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiKey) GetOwnerUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OwnerUsername.Get(), o.OwnerUsername.IsSet()
}

// HasOwnerUsername returns a boolean if a field has been set.
func (o *ApiKey) HasOwnerUsername() bool {
	if o != nil && o.OwnerUsername.IsSet() {
		return true
	}

	return false
}

// SetOwnerUsername gets a reference to the given NullableString and assigns it to the OwnerUsername field.
func (o *ApiKey) SetOwnerUsername(v string) {
	o.OwnerUsername.Set(&v)
}
// SetOwnerUsernameNil sets the value for OwnerUsername to be an explicit nil
func (o *ApiKey) SetOwnerUsernameNil() {
	o.OwnerUsername.Set(nil)
}

// UnsetOwnerUsername ensures that no value is present for OwnerUsername, not even an explicit nil
func (o *ApiKey) UnsetOwnerUsername() {
	o.OwnerUsername.Unset()
}

func (o ApiKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedTimeMsecs.IsSet() {
		toSerialize["createdTimeMsecs"] = o.CreatedTimeMsecs.Get()
	}
	if o.CreatedUserSid.IsSet() {
		toSerialize["createdUserSid"] = o.CreatedUserSid.Get()
	}
	if o.CreatedUsername.IsSet() {
		toSerialize["createdUsername"] = o.CreatedUsername.Get()
	}
	if o.ExpiringTimeMsecs.IsSet() {
		toSerialize["expiringTimeMsecs"] = o.ExpiringTimeMsecs.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.IsActive.IsSet() {
		toSerialize["isActive"] = o.IsActive.Get()
	}
	if o.IsExpired.IsSet() {
		toSerialize["isExpired"] = o.IsExpired.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.OwnerUserSid.IsSet() {
		toSerialize["ownerUserSid"] = o.OwnerUserSid.Get()
	}
	if o.OwnerUsername.IsSet() {
		toSerialize["ownerUsername"] = o.OwnerUsername.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableApiKey struct {
	value *ApiKey
	isSet bool
}

func (v NullableApiKey) Get() *ApiKey {
	return v.value
}

func (v *NullableApiKey) Set(val *ApiKey) {
	v.value = val
	v.isSet = true
}

func (v NullableApiKey) IsSet() bool {
	return v.isSet
}

func (v *NullableApiKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiKey(val *ApiKey) *NullableApiKey {
	return &NullableApiKey{value: val, isSet: true}
}

func (v NullableApiKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


