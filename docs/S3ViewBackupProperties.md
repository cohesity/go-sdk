# S3ViewBackupProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to **NullableString** | Access key for the buckets which will be created for the source initiated jobs. This needs to be passed to Netapp for it to for doing all s3 communications. | [optional] 
**S3Config** | Pointer to [**S3BucketConfigProto**](S3BucketConfigProto.md) |  | [optional] 
**SecretKey** | Pointer to **NullableString** | Secret key for the buckets will be created for the source initiated jobs. This secret key needed to be sent to Netapp for writing data to our S3 views. | [optional] 
**SnapshotPrefixName** | Pointer to **NullableString** | The snapshot prefix which is needed at the time of incremental for backups. This is only needed for case of DP volume. | [optional] 
**ViewId** | Pointer to **NullableInt64** | The id of the S3 view. | [optional] 

## Methods

### NewS3ViewBackupProperties

`func NewS3ViewBackupProperties() *S3ViewBackupProperties`

NewS3ViewBackupProperties instantiates a new S3ViewBackupProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3ViewBackupPropertiesWithDefaults

`func NewS3ViewBackupPropertiesWithDefaults() *S3ViewBackupProperties`

NewS3ViewBackupPropertiesWithDefaults instantiates a new S3ViewBackupProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *S3ViewBackupProperties) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *S3ViewBackupProperties) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *S3ViewBackupProperties) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *S3ViewBackupProperties) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### SetAccessKeyNil

`func (o *S3ViewBackupProperties) SetAccessKeyNil(b bool)`

 SetAccessKeyNil sets the value for AccessKey to be an explicit nil

### UnsetAccessKey
`func (o *S3ViewBackupProperties) UnsetAccessKey()`

UnsetAccessKey ensures that no value is present for AccessKey, not even an explicit nil
### GetS3Config

`func (o *S3ViewBackupProperties) GetS3Config() S3BucketConfigProto`

GetS3Config returns the S3Config field if non-nil, zero value otherwise.

### GetS3ConfigOk

`func (o *S3ViewBackupProperties) GetS3ConfigOk() (*S3BucketConfigProto, bool)`

GetS3ConfigOk returns a tuple with the S3Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Config

`func (o *S3ViewBackupProperties) SetS3Config(v S3BucketConfigProto)`

SetS3Config sets S3Config field to given value.

### HasS3Config

`func (o *S3ViewBackupProperties) HasS3Config() bool`

HasS3Config returns a boolean if a field has been set.

### GetSecretKey

`func (o *S3ViewBackupProperties) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *S3ViewBackupProperties) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *S3ViewBackupProperties) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *S3ViewBackupProperties) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### SetSecretKeyNil

`func (o *S3ViewBackupProperties) SetSecretKeyNil(b bool)`

 SetSecretKeyNil sets the value for SecretKey to be an explicit nil

### UnsetSecretKey
`func (o *S3ViewBackupProperties) UnsetSecretKey()`

UnsetSecretKey ensures that no value is present for SecretKey, not even an explicit nil
### GetSnapshotPrefixName

`func (o *S3ViewBackupProperties) GetSnapshotPrefixName() string`

GetSnapshotPrefixName returns the SnapshotPrefixName field if non-nil, zero value otherwise.

### GetSnapshotPrefixNameOk

`func (o *S3ViewBackupProperties) GetSnapshotPrefixNameOk() (*string, bool)`

GetSnapshotPrefixNameOk returns a tuple with the SnapshotPrefixName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotPrefixName

`func (o *S3ViewBackupProperties) SetSnapshotPrefixName(v string)`

SetSnapshotPrefixName sets SnapshotPrefixName field to given value.

### HasSnapshotPrefixName

`func (o *S3ViewBackupProperties) HasSnapshotPrefixName() bool`

HasSnapshotPrefixName returns a boolean if a field has been set.

### SetSnapshotPrefixNameNil

`func (o *S3ViewBackupProperties) SetSnapshotPrefixNameNil(b bool)`

 SetSnapshotPrefixNameNil sets the value for SnapshotPrefixName to be an explicit nil

### UnsetSnapshotPrefixName
`func (o *S3ViewBackupProperties) UnsetSnapshotPrefixName()`

UnsetSnapshotPrefixName ensures that no value is present for SnapshotPrefixName, not even an explicit nil
### GetViewId

`func (o *S3ViewBackupProperties) GetViewId() int64`

GetViewId returns the ViewId field if non-nil, zero value otherwise.

### GetViewIdOk

`func (o *S3ViewBackupProperties) GetViewIdOk() (*int64, bool)`

GetViewIdOk returns a tuple with the ViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewId

`func (o *S3ViewBackupProperties) SetViewId(v int64)`

SetViewId sets ViewId field to given value.

### HasViewId

`func (o *S3ViewBackupProperties) HasViewId() bool`

HasViewId returns a boolean if a field has been set.

### SetViewIdNil

`func (o *S3ViewBackupProperties) SetViewIdNil(b bool)`

 SetViewIdNil sets the value for ViewId to be an explicit nil

### UnsetViewId
`func (o *S3ViewBackupProperties) UnsetViewId()`

UnsetViewId ensures that no value is present for ViewId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


