# S3OwnerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DistinguishedName** | Pointer to **NullableString** | Specifies the distinguished name of the bucket owner for an ABAC enabled S3 Bucket. | [optional] 
**UserId** | Pointer to **NullableString** | Specifies the user id of the owner. | [optional] 

## Methods

### NewS3OwnerInfo

`func NewS3OwnerInfo() *S3OwnerInfo`

NewS3OwnerInfo instantiates a new S3OwnerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3OwnerInfoWithDefaults

`func NewS3OwnerInfoWithDefaults() *S3OwnerInfo`

NewS3OwnerInfoWithDefaults instantiates a new S3OwnerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDistinguishedName

`func (o *S3OwnerInfo) GetDistinguishedName() string`

GetDistinguishedName returns the DistinguishedName field if non-nil, zero value otherwise.

### GetDistinguishedNameOk

`func (o *S3OwnerInfo) GetDistinguishedNameOk() (*string, bool)`

GetDistinguishedNameOk returns a tuple with the DistinguishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinguishedName

`func (o *S3OwnerInfo) SetDistinguishedName(v string)`

SetDistinguishedName sets DistinguishedName field to given value.

### HasDistinguishedName

`func (o *S3OwnerInfo) HasDistinguishedName() bool`

HasDistinguishedName returns a boolean if a field has been set.

### SetDistinguishedNameNil

`func (o *S3OwnerInfo) SetDistinguishedNameNil(b bool)`

 SetDistinguishedNameNil sets the value for DistinguishedName to be an explicit nil

### UnsetDistinguishedName
`func (o *S3OwnerInfo) UnsetDistinguishedName()`

UnsetDistinguishedName ensures that no value is present for DistinguishedName, not even an explicit nil
### GetUserId

`func (o *S3OwnerInfo) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *S3OwnerInfo) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *S3OwnerInfo) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *S3OwnerInfo) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *S3OwnerInfo) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *S3OwnerInfo) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


