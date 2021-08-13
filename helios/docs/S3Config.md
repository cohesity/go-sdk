# S3Config

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**S3AccessPath** | Pointer to **NullableString** | Specifies the path to access this View as an S3 share. | [optional] [readonly] 
**AclConfig** | Pointer to [**AclConfig**](AclConfig.md) | Specifies the ACL config of the View as an S3 bucket. | [optional] 
**OwnerInfo** | Pointer to [**S3OwnerInfo**](S3OwnerInfo.md) | Specifies the owner info of the View as an S3 bucket. | [optional] 

## Methods

### NewS3Config

`func NewS3Config() *S3Config`

NewS3Config instantiates a new S3Config object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3ConfigWithDefaults

`func NewS3ConfigWithDefaults() *S3Config`

NewS3ConfigWithDefaults instantiates a new S3Config object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetS3AccessPath

`func (o *S3Config) GetS3AccessPath() string`

GetS3AccessPath returns the S3AccessPath field if non-nil, zero value otherwise.

### GetS3AccessPathOk

`func (o *S3Config) GetS3AccessPathOk() (*string, bool)`

GetS3AccessPathOk returns a tuple with the S3AccessPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3AccessPath

`func (o *S3Config) SetS3AccessPath(v string)`

SetS3AccessPath sets S3AccessPath field to given value.

### HasS3AccessPath

`func (o *S3Config) HasS3AccessPath() bool`

HasS3AccessPath returns a boolean if a field has been set.

### SetS3AccessPathNil

`func (o *S3Config) SetS3AccessPathNil(b bool)`

 SetS3AccessPathNil sets the value for S3AccessPath to be an explicit nil

### UnsetS3AccessPath
`func (o *S3Config) UnsetS3AccessPath()`

UnsetS3AccessPath ensures that no value is present for S3AccessPath, not even an explicit nil
### GetAclConfig

`func (o *S3Config) GetAclConfig() AclConfig`

GetAclConfig returns the AclConfig field if non-nil, zero value otherwise.

### GetAclConfigOk

`func (o *S3Config) GetAclConfigOk() (*AclConfig, bool)`

GetAclConfigOk returns a tuple with the AclConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclConfig

`func (o *S3Config) SetAclConfig(v AclConfig)`

SetAclConfig sets AclConfig field to given value.

### HasAclConfig

`func (o *S3Config) HasAclConfig() bool`

HasAclConfig returns a boolean if a field has been set.

### GetOwnerInfo

`func (o *S3Config) GetOwnerInfo() S3OwnerInfo`

GetOwnerInfo returns the OwnerInfo field if non-nil, zero value otherwise.

### GetOwnerInfoOk

`func (o *S3Config) GetOwnerInfoOk() (*S3OwnerInfo, bool)`

GetOwnerInfoOk returns a tuple with the OwnerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerInfo

`func (o *S3Config) SetOwnerInfo(v S3OwnerInfo)`

SetOwnerInfo sets OwnerInfo field to given value.

### HasOwnerInfo

`func (o *S3Config) HasOwnerInfo() bool`

HasOwnerInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


