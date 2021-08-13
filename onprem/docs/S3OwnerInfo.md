# S3OwnerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **NullableString** | Specifies the user id of the owner. | 

## Methods

### NewS3OwnerInfo

`func NewS3OwnerInfo(userId NullableString, ) *S3OwnerInfo`

NewS3OwnerInfo instantiates a new S3OwnerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3OwnerInfoWithDefaults

`func NewS3OwnerInfoWithDefaults() *S3OwnerInfo`

NewS3OwnerInfoWithDefaults instantiates a new S3OwnerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


### SetUserIdNil

`func (o *S3OwnerInfo) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *S3OwnerInfo) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


