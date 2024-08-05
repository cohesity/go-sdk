# UserS3KeysParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**S3AccessKeyId** | Pointer to **NullableString** | Specifies the S3 Account Access Key ID. Allowed characters are: AlphaNumeric(a-zA-z0-9), underscore(_) and hyphen(-). Key should contain exactly 43 characters. | [optional] 
**S3SecretKey** | Pointer to **NullableString** | Specifies the S3 Account Secret Key. Allowed characters are: AlphaNumeric(a-zA-z0-9), underscore(_) and hyphen(-). Key should contain exactly 43 characters. | [optional] 
**Sid** | **NullableString** | Specifies the SID of the User. | 

## Methods

### NewUserS3KeysParams

`func NewUserS3KeysParams(sid NullableString, ) *UserS3KeysParams`

NewUserS3KeysParams instantiates a new UserS3KeysParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserS3KeysParamsWithDefaults

`func NewUserS3KeysParamsWithDefaults() *UserS3KeysParams`

NewUserS3KeysParamsWithDefaults instantiates a new UserS3KeysParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetS3AccessKeyId

`func (o *UserS3KeysParams) GetS3AccessKeyId() string`

GetS3AccessKeyId returns the S3AccessKeyId field if non-nil, zero value otherwise.

### GetS3AccessKeyIdOk

`func (o *UserS3KeysParams) GetS3AccessKeyIdOk() (*string, bool)`

GetS3AccessKeyIdOk returns a tuple with the S3AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3AccessKeyId

`func (o *UserS3KeysParams) SetS3AccessKeyId(v string)`

SetS3AccessKeyId sets S3AccessKeyId field to given value.

### HasS3AccessKeyId

`func (o *UserS3KeysParams) HasS3AccessKeyId() bool`

HasS3AccessKeyId returns a boolean if a field has been set.

### SetS3AccessKeyIdNil

`func (o *UserS3KeysParams) SetS3AccessKeyIdNil(b bool)`

 SetS3AccessKeyIdNil sets the value for S3AccessKeyId to be an explicit nil

### UnsetS3AccessKeyId
`func (o *UserS3KeysParams) UnsetS3AccessKeyId()`

UnsetS3AccessKeyId ensures that no value is present for S3AccessKeyId, not even an explicit nil
### GetS3SecretKey

`func (o *UserS3KeysParams) GetS3SecretKey() string`

GetS3SecretKey returns the S3SecretKey field if non-nil, zero value otherwise.

### GetS3SecretKeyOk

`func (o *UserS3KeysParams) GetS3SecretKeyOk() (*string, bool)`

GetS3SecretKeyOk returns a tuple with the S3SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3SecretKey

`func (o *UserS3KeysParams) SetS3SecretKey(v string)`

SetS3SecretKey sets S3SecretKey field to given value.

### HasS3SecretKey

`func (o *UserS3KeysParams) HasS3SecretKey() bool`

HasS3SecretKey returns a boolean if a field has been set.

### SetS3SecretKeyNil

`func (o *UserS3KeysParams) SetS3SecretKeyNil(b bool)`

 SetS3SecretKeyNil sets the value for S3SecretKey to be an explicit nil

### UnsetS3SecretKey
`func (o *UserS3KeysParams) UnsetS3SecretKey()`

UnsetS3SecretKey ensures that no value is present for S3SecretKey, not even an explicit nil
### GetSid

`func (o *UserS3KeysParams) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *UserS3KeysParams) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *UserS3KeysParams) SetSid(v string)`

SetSid sets Sid field to given value.


### SetSidNil

`func (o *UserS3KeysParams) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *UserS3KeysParams) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


