# IamUserAwsCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | **NullableString** | Specifies Access key of the AWS account. | 
**SecretAccessKey** | **NullableString** | Specifies Secret Access key of the AWS account. | 
**Arn** | **NullableString** | Specifies Amazon Resource Name (owner ID) of the IAM user, acts as an unique identifier of as AWS entity. | 

## Methods

### NewIamUserAwsCredentials

`func NewIamUserAwsCredentials(accessKey NullableString, secretAccessKey NullableString, arn NullableString, ) *IamUserAwsCredentials`

NewIamUserAwsCredentials instantiates a new IamUserAwsCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamUserAwsCredentialsWithDefaults

`func NewIamUserAwsCredentialsWithDefaults() *IamUserAwsCredentials`

NewIamUserAwsCredentialsWithDefaults instantiates a new IamUserAwsCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *IamUserAwsCredentials) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *IamUserAwsCredentials) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *IamUserAwsCredentials) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.


### SetAccessKeyNil

`func (o *IamUserAwsCredentials) SetAccessKeyNil(b bool)`

 SetAccessKeyNil sets the value for AccessKey to be an explicit nil

### UnsetAccessKey
`func (o *IamUserAwsCredentials) UnsetAccessKey()`

UnsetAccessKey ensures that no value is present for AccessKey, not even an explicit nil
### GetSecretAccessKey

`func (o *IamUserAwsCredentials) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *IamUserAwsCredentials) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *IamUserAwsCredentials) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.


### SetSecretAccessKeyNil

`func (o *IamUserAwsCredentials) SetSecretAccessKeyNil(b bool)`

 SetSecretAccessKeyNil sets the value for SecretAccessKey to be an explicit nil

### UnsetSecretAccessKey
`func (o *IamUserAwsCredentials) UnsetSecretAccessKey()`

UnsetSecretAccessKey ensures that no value is present for SecretAccessKey, not even an explicit nil
### GetArn

`func (o *IamUserAwsCredentials) GetArn() string`

GetArn returns the Arn field if non-nil, zero value otherwise.

### GetArnOk

`func (o *IamUserAwsCredentials) GetArnOk() (*string, bool)`

GetArnOk returns a tuple with the Arn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArn

`func (o *IamUserAwsCredentials) SetArn(v string)`

SetArn sets Arn field to given value.


### SetArnNil

`func (o *IamUserAwsCredentials) SetArnNil(b bool)`

 SetArnNil sets the value for Arn to be an explicit nil

### UnsetArn
`func (o *IamUserAwsCredentials) UnsetArn()`

UnsetArn ensures that no value is present for Arn, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


