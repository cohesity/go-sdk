# AmazonCloudCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKeyId** | Pointer to **NullableString** | Specifies the access key for Amazon service account. See the Cohesity online help for the value to specify for this field based on the current S3 Compatible Vault (External Target) type. For example for Iron Mountain, specify the user name from Iron Mountain for this field. | [optional] 
**AuthMethod** | Pointer to **NullableString** | Specifies the iauth method used for the request See the Cohesity online help for the value to specify for this field based on the current S3-compatible Vault (External Target) type. Specifies the authentication method to be used for API calls. &#39;kUseIAMUser&#39; indicates a user based authentication. &#39;kUseIAMRole&#39; indicates a role based authentication, used only for AWS CE. | [optional] 
**C2sAccessPortal** | Pointer to [**C2SAccessPortal**](C2SAccessPortal.md) |  | [optional] 
**IamRoleArn** | Pointer to **NullableString** | Specifies the iam role arn Amazon service account. See the Cohesity online help for the value to specify for this field based on the current S3-compatible Vault (External Target) type. | [optional] 
**Region** | Pointer to **NullableString** | Specifies the region to use for the Amazon service account. | [optional] 
**SecretAccessKey** | Pointer to **NullableString** | Specifies the secret access key for Amazon service account. See the Cohesity online help for the value to specify for this field based on the current S3-compatible Vault (External Target) type. | [optional] 
**ServiceUrl** | Pointer to **NullableString** | Specifies the URL (Endpoint) for the service such as s3like.notamazon.com. This field is only significant for S3-compatible cloud services. | [optional] 
**SignatureVersion** | Pointer to **NullableInt32** | Specifies the version of the S3 Compliance. This field must be set to 2 or 4 and the default version is 2. This field is only significant for S3-compatible cloud services. See the Cohesity online help for the supported S3-compatible Vault (External Target) types and the value to specify for this field based on the current S3-compatible Vault (External Target) type. | [optional] 
**TierType** | Pointer to **NullableString** | Specifies the storage class of AWS. AmazonTierType specifies the storage class for AWS. &#39;kAmazonS3Standard&#39; indicates a tier type of Amazon properties that is accessed frequently. &#39;kAmazonS3StandardIA&#39; indicates a tier type of Amazon properties that is accessed less frequently, but requires rapid access when needed. &#39;kAmazonGlacier&#39; indicates a tier type of Amazon properties that is accessed rarely. &#39;kAmazonS3OneZoneIA&#39; indicates a tier type of Amazon properties for long-lived, but less frequently accessed data. &#39;kAmazonS3IntelligentTiering&#39; indicates a tier type of Amazon properties for data with unknown or changing access patterns. &#39;kAmazonS3Glacier&#39; indicates a tier type of Amazon properties for data that provides secure, durable object storage for long-term data retention and digital preservation. It provides three options for access to archives, from a few minutes to several hours. &#39;kAmazonS3GlacierDeepArchive&#39; indicates a tier type of Amazon properties for data that provides secure, durable object storage for long-term data retention and digital preservation. It provides two access options ranging from 12 to 48 hours. | [optional] 
**UseHttps** | Pointer to **NullableBool** | Specifies whether to use http or https to connect to the service. If true, a secure connection (https) is used. This field is only significant for S3-compatible cloud services. | [optional] 

## Methods

### NewAmazonCloudCredentials

`func NewAmazonCloudCredentials() *AmazonCloudCredentials`

NewAmazonCloudCredentials instantiates a new AmazonCloudCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonCloudCredentialsWithDefaults

`func NewAmazonCloudCredentialsWithDefaults() *AmazonCloudCredentials`

NewAmazonCloudCredentialsWithDefaults instantiates a new AmazonCloudCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKeyId

`func (o *AmazonCloudCredentials) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *AmazonCloudCredentials) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *AmazonCloudCredentials) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.

### HasAccessKeyId

`func (o *AmazonCloudCredentials) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### SetAccessKeyIdNil

`func (o *AmazonCloudCredentials) SetAccessKeyIdNil(b bool)`

 SetAccessKeyIdNil sets the value for AccessKeyId to be an explicit nil

### UnsetAccessKeyId
`func (o *AmazonCloudCredentials) UnsetAccessKeyId()`

UnsetAccessKeyId ensures that no value is present for AccessKeyId, not even an explicit nil
### GetAuthMethod

`func (o *AmazonCloudCredentials) GetAuthMethod() string`

GetAuthMethod returns the AuthMethod field if non-nil, zero value otherwise.

### GetAuthMethodOk

`func (o *AmazonCloudCredentials) GetAuthMethodOk() (*string, bool)`

GetAuthMethodOk returns a tuple with the AuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethod

`func (o *AmazonCloudCredentials) SetAuthMethod(v string)`

SetAuthMethod sets AuthMethod field to given value.

### HasAuthMethod

`func (o *AmazonCloudCredentials) HasAuthMethod() bool`

HasAuthMethod returns a boolean if a field has been set.

### SetAuthMethodNil

`func (o *AmazonCloudCredentials) SetAuthMethodNil(b bool)`

 SetAuthMethodNil sets the value for AuthMethod to be an explicit nil

### UnsetAuthMethod
`func (o *AmazonCloudCredentials) UnsetAuthMethod()`

UnsetAuthMethod ensures that no value is present for AuthMethod, not even an explicit nil
### GetC2sAccessPortal

`func (o *AmazonCloudCredentials) GetC2sAccessPortal() C2SAccessPortal`

GetC2sAccessPortal returns the C2sAccessPortal field if non-nil, zero value otherwise.

### GetC2sAccessPortalOk

`func (o *AmazonCloudCredentials) GetC2sAccessPortalOk() (*C2SAccessPortal, bool)`

GetC2sAccessPortalOk returns a tuple with the C2sAccessPortal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetC2sAccessPortal

`func (o *AmazonCloudCredentials) SetC2sAccessPortal(v C2SAccessPortal)`

SetC2sAccessPortal sets C2sAccessPortal field to given value.

### HasC2sAccessPortal

`func (o *AmazonCloudCredentials) HasC2sAccessPortal() bool`

HasC2sAccessPortal returns a boolean if a field has been set.

### GetIamRoleArn

`func (o *AmazonCloudCredentials) GetIamRoleArn() string`

GetIamRoleArn returns the IamRoleArn field if non-nil, zero value otherwise.

### GetIamRoleArnOk

`func (o *AmazonCloudCredentials) GetIamRoleArnOk() (*string, bool)`

GetIamRoleArnOk returns a tuple with the IamRoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamRoleArn

`func (o *AmazonCloudCredentials) SetIamRoleArn(v string)`

SetIamRoleArn sets IamRoleArn field to given value.

### HasIamRoleArn

`func (o *AmazonCloudCredentials) HasIamRoleArn() bool`

HasIamRoleArn returns a boolean if a field has been set.

### SetIamRoleArnNil

`func (o *AmazonCloudCredentials) SetIamRoleArnNil(b bool)`

 SetIamRoleArnNil sets the value for IamRoleArn to be an explicit nil

### UnsetIamRoleArn
`func (o *AmazonCloudCredentials) UnsetIamRoleArn()`

UnsetIamRoleArn ensures that no value is present for IamRoleArn, not even an explicit nil
### GetRegion

`func (o *AmazonCloudCredentials) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AmazonCloudCredentials) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AmazonCloudCredentials) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AmazonCloudCredentials) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *AmazonCloudCredentials) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *AmazonCloudCredentials) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetSecretAccessKey

`func (o *AmazonCloudCredentials) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *AmazonCloudCredentials) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *AmazonCloudCredentials) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.

### HasSecretAccessKey

`func (o *AmazonCloudCredentials) HasSecretAccessKey() bool`

HasSecretAccessKey returns a boolean if a field has been set.

### SetSecretAccessKeyNil

`func (o *AmazonCloudCredentials) SetSecretAccessKeyNil(b bool)`

 SetSecretAccessKeyNil sets the value for SecretAccessKey to be an explicit nil

### UnsetSecretAccessKey
`func (o *AmazonCloudCredentials) UnsetSecretAccessKey()`

UnsetSecretAccessKey ensures that no value is present for SecretAccessKey, not even an explicit nil
### GetServiceUrl

`func (o *AmazonCloudCredentials) GetServiceUrl() string`

GetServiceUrl returns the ServiceUrl field if non-nil, zero value otherwise.

### GetServiceUrlOk

`func (o *AmazonCloudCredentials) GetServiceUrlOk() (*string, bool)`

GetServiceUrlOk returns a tuple with the ServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrl

`func (o *AmazonCloudCredentials) SetServiceUrl(v string)`

SetServiceUrl sets ServiceUrl field to given value.

### HasServiceUrl

`func (o *AmazonCloudCredentials) HasServiceUrl() bool`

HasServiceUrl returns a boolean if a field has been set.

### SetServiceUrlNil

`func (o *AmazonCloudCredentials) SetServiceUrlNil(b bool)`

 SetServiceUrlNil sets the value for ServiceUrl to be an explicit nil

### UnsetServiceUrl
`func (o *AmazonCloudCredentials) UnsetServiceUrl()`

UnsetServiceUrl ensures that no value is present for ServiceUrl, not even an explicit nil
### GetSignatureVersion

`func (o *AmazonCloudCredentials) GetSignatureVersion() int32`

GetSignatureVersion returns the SignatureVersion field if non-nil, zero value otherwise.

### GetSignatureVersionOk

`func (o *AmazonCloudCredentials) GetSignatureVersionOk() (*int32, bool)`

GetSignatureVersionOk returns a tuple with the SignatureVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureVersion

`func (o *AmazonCloudCredentials) SetSignatureVersion(v int32)`

SetSignatureVersion sets SignatureVersion field to given value.

### HasSignatureVersion

`func (o *AmazonCloudCredentials) HasSignatureVersion() bool`

HasSignatureVersion returns a boolean if a field has been set.

### SetSignatureVersionNil

`func (o *AmazonCloudCredentials) SetSignatureVersionNil(b bool)`

 SetSignatureVersionNil sets the value for SignatureVersion to be an explicit nil

### UnsetSignatureVersion
`func (o *AmazonCloudCredentials) UnsetSignatureVersion()`

UnsetSignatureVersion ensures that no value is present for SignatureVersion, not even an explicit nil
### GetTierType

`func (o *AmazonCloudCredentials) GetTierType() string`

GetTierType returns the TierType field if non-nil, zero value otherwise.

### GetTierTypeOk

`func (o *AmazonCloudCredentials) GetTierTypeOk() (*string, bool)`

GetTierTypeOk returns a tuple with the TierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierType

`func (o *AmazonCloudCredentials) SetTierType(v string)`

SetTierType sets TierType field to given value.

### HasTierType

`func (o *AmazonCloudCredentials) HasTierType() bool`

HasTierType returns a boolean if a field has been set.

### SetTierTypeNil

`func (o *AmazonCloudCredentials) SetTierTypeNil(b bool)`

 SetTierTypeNil sets the value for TierType to be an explicit nil

### UnsetTierType
`func (o *AmazonCloudCredentials) UnsetTierType()`

UnsetTierType ensures that no value is present for TierType, not even an explicit nil
### GetUseHttps

`func (o *AmazonCloudCredentials) GetUseHttps() bool`

GetUseHttps returns the UseHttps field if non-nil, zero value otherwise.

### GetUseHttpsOk

`func (o *AmazonCloudCredentials) GetUseHttpsOk() (*bool, bool)`

GetUseHttpsOk returns a tuple with the UseHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseHttps

`func (o *AmazonCloudCredentials) SetUseHttps(v bool)`

SetUseHttps sets UseHttps field to given value.

### HasUseHttps

`func (o *AmazonCloudCredentials) HasUseHttps() bool`

HasUseHttps returns a boolean if a field has been set.

### SetUseHttpsNil

`func (o *AmazonCloudCredentials) SetUseHttpsNil(b bool)`

 SetUseHttpsNil sets the value for UseHttps to be an explicit nil

### UnsetUseHttps
`func (o *AmazonCloudCredentials) UnsetUseHttps()`

UnsetUseHttps ensures that no value is present for UseHttps, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


