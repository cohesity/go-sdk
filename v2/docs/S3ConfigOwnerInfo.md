# S3ConfigOwnerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DistinguishedName** | Pointer to **NullableString** | Specifies the distinguished name of the bucket owner for an ABAC enabled S3 Bucket. | [optional] 
**UserId** | Pointer to **NullableString** | Specifies the user id of the owner. | [optional] 

## Methods

### NewS3ConfigOwnerInfo

`func NewS3ConfigOwnerInfo() *S3ConfigOwnerInfo`

NewS3ConfigOwnerInfo instantiates a new S3ConfigOwnerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3ConfigOwnerInfoWithDefaults

`func NewS3ConfigOwnerInfoWithDefaults() *S3ConfigOwnerInfo`

NewS3ConfigOwnerInfoWithDefaults instantiates a new S3ConfigOwnerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDistinguishedName

`func (o *S3ConfigOwnerInfo) GetDistinguishedName() string`

GetDistinguishedName returns the DistinguishedName field if non-nil, zero value otherwise.

### GetDistinguishedNameOk

`func (o *S3ConfigOwnerInfo) GetDistinguishedNameOk() (*string, bool)`

GetDistinguishedNameOk returns a tuple with the DistinguishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinguishedName

`func (o *S3ConfigOwnerInfo) SetDistinguishedName(v string)`

SetDistinguishedName sets DistinguishedName field to given value.

### HasDistinguishedName

`func (o *S3ConfigOwnerInfo) HasDistinguishedName() bool`

HasDistinguishedName returns a boolean if a field has been set.

### SetDistinguishedNameNil

`func (o *S3ConfigOwnerInfo) SetDistinguishedNameNil(b bool)`

 SetDistinguishedNameNil sets the value for DistinguishedName to be an explicit nil

### UnsetDistinguishedName
`func (o *S3ConfigOwnerInfo) UnsetDistinguishedName()`

UnsetDistinguishedName ensures that no value is present for DistinguishedName, not even an explicit nil
### GetUserId

`func (o *S3ConfigOwnerInfo) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *S3ConfigOwnerInfo) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *S3ConfigOwnerInfo) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *S3ConfigOwnerInfo) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *S3ConfigOwnerInfo) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *S3ConfigOwnerInfo) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


