# AwsCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to **NullableString** | Specifies Access key of the AWS account. | [optional] 
**AmazonResourceName** | Pointer to **NullableString** | Specifies Amazon Resource Name (owner ID) of the IAM user, act as an unique identifier of as AWS entity. | [optional] 
**AuthMethod** | Pointer to **NullableString** | Specifies the authentication method to be used for API calls. Specifies the authentication method to be used for API calls. &#39;kUseIAMUser&#39; indicates a user based authentication. &#39;kUseIAMRole&#39; indicates a role based authentication, used only for AWS CE. | [optional] 
**AwsType** | Pointer to **NullableString** | Specifies the entity type such as &#39;kIAMUser&#39; if the environment is kAWS. Specifies the type of an AWS source entity. &#39;kIAMUser&#39; indicates a unique user within an AWS account. &#39;kRegion&#39; indicates a geographical region in the global infrastructure. &#39;kAvailabilityZone&#39; indicates an availability zone within a region. &#39;kEC2Instance&#39; indicates a Virtual Machine running in AWS environment. &#39;kVPC&#39; indicates a virtual private cloud (VPC) network within AWS. &#39;kSubnet&#39; indicates a subnet inside the VPC. &#39;kNetworkSecurityGroup&#39; represents a network security group. &#39;kInstanceType&#39; represents various machine types. &#39;kKeyPair&#39; represents a pair of public and private key used to login into a Virtual Machine. &#39;kTag&#39; represents a tag attached to EC2 instance. &#39;kRDSOptionGroup&#39; represents a RDS option group for configuring database features. &#39;kRDSParameterGroup&#39; represents a RDS parameter group. &#39;kRDSInstance&#39; represents a RDS DB instance. &#39;kRDSSubnet&#39; represents a RDS subnet. &#39;kRDSTag&#39; represents a tag attached to RDS instance. | [optional] 
**IamRoleArn** | Pointer to **NullableString** | Specifies the IAM role which will be used to access the security credentials required for API calls. | [optional] 
**SecretAccessKey** | Pointer to **NullableString** | Specifies Secret Access key of the AWS account. | [optional] 
**SubscriptionType** | Pointer to **NullableString** | Specifies the subscription type of AWS such as &#39;kAWSCommercial&#39; or &#39;kAWSGovCloud&#39;. Specifies the subscription type of an AWS source entity. &#39;kAWSCommercial&#39; indicates a standard AWS subscription. &#39;kAWSGovCloud&#39; indicates a govt AWS subscription. | [optional] 

## Methods

### NewAwsCredentials

`func NewAwsCredentials() *AwsCredentials`

NewAwsCredentials instantiates a new AwsCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsCredentialsWithDefaults

`func NewAwsCredentialsWithDefaults() *AwsCredentials`

NewAwsCredentialsWithDefaults instantiates a new AwsCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *AwsCredentials) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *AwsCredentials) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *AwsCredentials) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *AwsCredentials) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### SetAccessKeyNil

`func (o *AwsCredentials) SetAccessKeyNil(b bool)`

 SetAccessKeyNil sets the value for AccessKey to be an explicit nil

### UnsetAccessKey
`func (o *AwsCredentials) UnsetAccessKey()`

UnsetAccessKey ensures that no value is present for AccessKey, not even an explicit nil
### GetAmazonResourceName

`func (o *AwsCredentials) GetAmazonResourceName() string`

GetAmazonResourceName returns the AmazonResourceName field if non-nil, zero value otherwise.

### GetAmazonResourceNameOk

`func (o *AwsCredentials) GetAmazonResourceNameOk() (*string, bool)`

GetAmazonResourceNameOk returns a tuple with the AmazonResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonResourceName

`func (o *AwsCredentials) SetAmazonResourceName(v string)`

SetAmazonResourceName sets AmazonResourceName field to given value.

### HasAmazonResourceName

`func (o *AwsCredentials) HasAmazonResourceName() bool`

HasAmazonResourceName returns a boolean if a field has been set.

### SetAmazonResourceNameNil

`func (o *AwsCredentials) SetAmazonResourceNameNil(b bool)`

 SetAmazonResourceNameNil sets the value for AmazonResourceName to be an explicit nil

### UnsetAmazonResourceName
`func (o *AwsCredentials) UnsetAmazonResourceName()`

UnsetAmazonResourceName ensures that no value is present for AmazonResourceName, not even an explicit nil
### GetAuthMethod

`func (o *AwsCredentials) GetAuthMethod() string`

GetAuthMethod returns the AuthMethod field if non-nil, zero value otherwise.

### GetAuthMethodOk

`func (o *AwsCredentials) GetAuthMethodOk() (*string, bool)`

GetAuthMethodOk returns a tuple with the AuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethod

`func (o *AwsCredentials) SetAuthMethod(v string)`

SetAuthMethod sets AuthMethod field to given value.

### HasAuthMethod

`func (o *AwsCredentials) HasAuthMethod() bool`

HasAuthMethod returns a boolean if a field has been set.

### SetAuthMethodNil

`func (o *AwsCredentials) SetAuthMethodNil(b bool)`

 SetAuthMethodNil sets the value for AuthMethod to be an explicit nil

### UnsetAuthMethod
`func (o *AwsCredentials) UnsetAuthMethod()`

UnsetAuthMethod ensures that no value is present for AuthMethod, not even an explicit nil
### GetAwsType

`func (o *AwsCredentials) GetAwsType() string`

GetAwsType returns the AwsType field if non-nil, zero value otherwise.

### GetAwsTypeOk

`func (o *AwsCredentials) GetAwsTypeOk() (*string, bool)`

GetAwsTypeOk returns a tuple with the AwsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsType

`func (o *AwsCredentials) SetAwsType(v string)`

SetAwsType sets AwsType field to given value.

### HasAwsType

`func (o *AwsCredentials) HasAwsType() bool`

HasAwsType returns a boolean if a field has been set.

### SetAwsTypeNil

`func (o *AwsCredentials) SetAwsTypeNil(b bool)`

 SetAwsTypeNil sets the value for AwsType to be an explicit nil

### UnsetAwsType
`func (o *AwsCredentials) UnsetAwsType()`

UnsetAwsType ensures that no value is present for AwsType, not even an explicit nil
### GetIamRoleArn

`func (o *AwsCredentials) GetIamRoleArn() string`

GetIamRoleArn returns the IamRoleArn field if non-nil, zero value otherwise.

### GetIamRoleArnOk

`func (o *AwsCredentials) GetIamRoleArnOk() (*string, bool)`

GetIamRoleArnOk returns a tuple with the IamRoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamRoleArn

`func (o *AwsCredentials) SetIamRoleArn(v string)`

SetIamRoleArn sets IamRoleArn field to given value.

### HasIamRoleArn

`func (o *AwsCredentials) HasIamRoleArn() bool`

HasIamRoleArn returns a boolean if a field has been set.

### SetIamRoleArnNil

`func (o *AwsCredentials) SetIamRoleArnNil(b bool)`

 SetIamRoleArnNil sets the value for IamRoleArn to be an explicit nil

### UnsetIamRoleArn
`func (o *AwsCredentials) UnsetIamRoleArn()`

UnsetIamRoleArn ensures that no value is present for IamRoleArn, not even an explicit nil
### GetSecretAccessKey

`func (o *AwsCredentials) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *AwsCredentials) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *AwsCredentials) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.

### HasSecretAccessKey

`func (o *AwsCredentials) HasSecretAccessKey() bool`

HasSecretAccessKey returns a boolean if a field has been set.

### SetSecretAccessKeyNil

`func (o *AwsCredentials) SetSecretAccessKeyNil(b bool)`

 SetSecretAccessKeyNil sets the value for SecretAccessKey to be an explicit nil

### UnsetSecretAccessKey
`func (o *AwsCredentials) UnsetSecretAccessKey()`

UnsetSecretAccessKey ensures that no value is present for SecretAccessKey, not even an explicit nil
### GetSubscriptionType

`func (o *AwsCredentials) GetSubscriptionType() string`

GetSubscriptionType returns the SubscriptionType field if non-nil, zero value otherwise.

### GetSubscriptionTypeOk

`func (o *AwsCredentials) GetSubscriptionTypeOk() (*string, bool)`

GetSubscriptionTypeOk returns a tuple with the SubscriptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionType

`func (o *AwsCredentials) SetSubscriptionType(v string)`

SetSubscriptionType sets SubscriptionType field to given value.

### HasSubscriptionType

`func (o *AwsCredentials) HasSubscriptionType() bool`

HasSubscriptionType returns a boolean if a field has been set.

### SetSubscriptionTypeNil

`func (o *AwsCredentials) SetSubscriptionTypeNil(b bool)`

 SetSubscriptionTypeNil sets the value for SubscriptionType to be an explicit nil

### UnsetSubscriptionType
`func (o *AwsCredentials) UnsetSubscriptionType()`

UnsetSubscriptionType ensures that no value is present for SubscriptionType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


