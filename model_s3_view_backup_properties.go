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

// S3ViewBackupProperties // -----------------------------------------------------------------------------
type S3ViewBackupProperties struct {
	// Access key for the buckets which will be created for the source initiated jobs. This needs to be passed to Netapp for it to for doing all s3 communications.
	AccessKey NullableString `json:"accessKey,omitempty"`
	S3Config *S3BucketConfigProto `json:"s3Config,omitempty"`
	// Secret key for the buckets will be created for the source initiated jobs. This secret key needed to be sent to Netapp for writing data to our S3 views.
	SecretKey NullableString `json:"secretKey,omitempty"`
	// The snapshot prefix which is needed at the time of incremental for backups. This is only needed for case of DP volume.
	SnapshotPrefixName NullableString `json:"snapshotPrefixName,omitempty"`
	// The id of the S3 view.
	ViewId NullableInt64 `json:"viewId,omitempty"`
}

// NewS3ViewBackupProperties instantiates a new S3ViewBackupProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewS3ViewBackupProperties() *S3ViewBackupProperties {
	this := S3ViewBackupProperties{}
	return &this
}

// NewS3ViewBackupPropertiesWithDefaults instantiates a new S3ViewBackupProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewS3ViewBackupPropertiesWithDefaults() *S3ViewBackupProperties {
	this := S3ViewBackupProperties{}
	return &this
}

// GetAccessKey returns the AccessKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *S3ViewBackupProperties) GetAccessKey() string {
	if o == nil || o.AccessKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.AccessKey.Get()
}

// GetAccessKeyOk returns a tuple with the AccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *S3ViewBackupProperties) GetAccessKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccessKey.Get(), o.AccessKey.IsSet()
}

// HasAccessKey returns a boolean if a field has been set.
func (o *S3ViewBackupProperties) HasAccessKey() bool {
	if o != nil && o.AccessKey.IsSet() {
		return true
	}

	return false
}

// SetAccessKey gets a reference to the given NullableString and assigns it to the AccessKey field.
func (o *S3ViewBackupProperties) SetAccessKey(v string) {
	o.AccessKey.Set(&v)
}
// SetAccessKeyNil sets the value for AccessKey to be an explicit nil
func (o *S3ViewBackupProperties) SetAccessKeyNil() {
	o.AccessKey.Set(nil)
}

// UnsetAccessKey ensures that no value is present for AccessKey, not even an explicit nil
func (o *S3ViewBackupProperties) UnsetAccessKey() {
	o.AccessKey.Unset()
}

// GetS3Config returns the S3Config field value if set, zero value otherwise.
func (o *S3ViewBackupProperties) GetS3Config() S3BucketConfigProto {
	if o == nil || o.S3Config == nil {
		var ret S3BucketConfigProto
		return ret
	}
	return *o.S3Config
}

// GetS3ConfigOk returns a tuple with the S3Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3ViewBackupProperties) GetS3ConfigOk() (*S3BucketConfigProto, bool) {
	if o == nil || o.S3Config == nil {
		return nil, false
	}
	return o.S3Config, true
}

// HasS3Config returns a boolean if a field has been set.
func (o *S3ViewBackupProperties) HasS3Config() bool {
	if o != nil && o.S3Config != nil {
		return true
	}

	return false
}

// SetS3Config gets a reference to the given S3BucketConfigProto and assigns it to the S3Config field.
func (o *S3ViewBackupProperties) SetS3Config(v S3BucketConfigProto) {
	o.S3Config = &v
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *S3ViewBackupProperties) GetSecretKey() string {
	if o == nil || o.SecretKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.SecretKey.Get()
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *S3ViewBackupProperties) GetSecretKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SecretKey.Get(), o.SecretKey.IsSet()
}

// HasSecretKey returns a boolean if a field has been set.
func (o *S3ViewBackupProperties) HasSecretKey() bool {
	if o != nil && o.SecretKey.IsSet() {
		return true
	}

	return false
}

// SetSecretKey gets a reference to the given NullableString and assigns it to the SecretKey field.
func (o *S3ViewBackupProperties) SetSecretKey(v string) {
	o.SecretKey.Set(&v)
}
// SetSecretKeyNil sets the value for SecretKey to be an explicit nil
func (o *S3ViewBackupProperties) SetSecretKeyNil() {
	o.SecretKey.Set(nil)
}

// UnsetSecretKey ensures that no value is present for SecretKey, not even an explicit nil
func (o *S3ViewBackupProperties) UnsetSecretKey() {
	o.SecretKey.Unset()
}

// GetSnapshotPrefixName returns the SnapshotPrefixName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *S3ViewBackupProperties) GetSnapshotPrefixName() string {
	if o == nil || o.SnapshotPrefixName.Get() == nil {
		var ret string
		return ret
	}
	return *o.SnapshotPrefixName.Get()
}

// GetSnapshotPrefixNameOk returns a tuple with the SnapshotPrefixName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *S3ViewBackupProperties) GetSnapshotPrefixNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SnapshotPrefixName.Get(), o.SnapshotPrefixName.IsSet()
}

// HasSnapshotPrefixName returns a boolean if a field has been set.
func (o *S3ViewBackupProperties) HasSnapshotPrefixName() bool {
	if o != nil && o.SnapshotPrefixName.IsSet() {
		return true
	}

	return false
}

// SetSnapshotPrefixName gets a reference to the given NullableString and assigns it to the SnapshotPrefixName field.
func (o *S3ViewBackupProperties) SetSnapshotPrefixName(v string) {
	o.SnapshotPrefixName.Set(&v)
}
// SetSnapshotPrefixNameNil sets the value for SnapshotPrefixName to be an explicit nil
func (o *S3ViewBackupProperties) SetSnapshotPrefixNameNil() {
	o.SnapshotPrefixName.Set(nil)
}

// UnsetSnapshotPrefixName ensures that no value is present for SnapshotPrefixName, not even an explicit nil
func (o *S3ViewBackupProperties) UnsetSnapshotPrefixName() {
	o.SnapshotPrefixName.Unset()
}

// GetViewId returns the ViewId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *S3ViewBackupProperties) GetViewId() int64 {
	if o == nil || o.ViewId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ViewId.Get()
}

// GetViewIdOk returns a tuple with the ViewId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *S3ViewBackupProperties) GetViewIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ViewId.Get(), o.ViewId.IsSet()
}

// HasViewId returns a boolean if a field has been set.
func (o *S3ViewBackupProperties) HasViewId() bool {
	if o != nil && o.ViewId.IsSet() {
		return true
	}

	return false
}

// SetViewId gets a reference to the given NullableInt64 and assigns it to the ViewId field.
func (o *S3ViewBackupProperties) SetViewId(v int64) {
	o.ViewId.Set(&v)
}
// SetViewIdNil sets the value for ViewId to be an explicit nil
func (o *S3ViewBackupProperties) SetViewIdNil() {
	o.ViewId.Set(nil)
}

// UnsetViewId ensures that no value is present for ViewId, not even an explicit nil
func (o *S3ViewBackupProperties) UnsetViewId() {
	o.ViewId.Unset()
}

func (o S3ViewBackupProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessKey.IsSet() {
		toSerialize["accessKey"] = o.AccessKey.Get()
	}
	if o.S3Config != nil {
		toSerialize["s3Config"] = o.S3Config
	}
	if o.SecretKey.IsSet() {
		toSerialize["secretKey"] = o.SecretKey.Get()
	}
	if o.SnapshotPrefixName.IsSet() {
		toSerialize["snapshotPrefixName"] = o.SnapshotPrefixName.Get()
	}
	if o.ViewId.IsSet() {
		toSerialize["viewId"] = o.ViewId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableS3ViewBackupProperties struct {
	value *S3ViewBackupProperties
	isSet bool
}

func (v NullableS3ViewBackupProperties) Get() *S3ViewBackupProperties {
	return v.value
}

func (v *NullableS3ViewBackupProperties) Set(val *S3ViewBackupProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableS3ViewBackupProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableS3ViewBackupProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableS3ViewBackupProperties(val *S3ViewBackupProperties) *NullableS3ViewBackupProperties {
	return &NullableS3ViewBackupProperties{value: val, isSet: true}
}

func (v NullableS3ViewBackupProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableS3ViewBackupProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


